package report

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/manisbindra/sbi/pkg/domain"
	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) (*sql.DB, *database.Repository) {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)

	_, err = db.Exec("PRAGMA foreign_keys=ON")
	require.NoError(t, err)

	require.NoError(t, database.CreateTables(db))

	return db, database.NewRepository(db)
}

func TestGenerateJSONReport_MultipleOSes(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			Digest: "sha256:abc123", SizeBytes: 85000000,
			CriticalVulnerabilities: 0, HighVulnerabilities: 0, TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12.1"}},
		},
		{
			Name: "ubuntu-python:3.12", Registry: "r", Repository: "repo2", Tag: "3.12",
			BaseOSName: "ubuntu", BaseOSVersion: "22.04",
			Digest: "sha256:def456", SizeBytes: 120000000,
			CriticalVulnerabilities: 1, HighVulnerabilities: 2, TotalVulnerabilities: 10,
			Languages: []domain.Language{{Language: "python", Version: "3.12.1"}},
		},
		{
			Name: "azl-go:1.21", Registry: "r", Repository: "repo3", Tag: "1.21",
			BaseOSName: "azurelinux", BaseOSVersion: "3.0",
			CriticalVulnerabilities: 0, TotalVulnerabilities: 0,
			Languages: []domain.Language{{Language: "go", Version: "1.21.0"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Equal(t, 10, report.TopN)
	require.Len(t, report.Images, 3)

	// Verify flat structure: each entry has language and baseOS
	for _, img := range report.Images {
		assert.NotEmpty(t, img.Language)
		assert.NotEmpty(t, img.BaseOS)
		assert.NotEmpty(t, img.Name)
		assert.Greater(t, img.Rank, 0)
	}

	// Verify python entries have both OSes
	var pythonOSes []string
	for _, img := range report.Images {
		if img.Language == "python" {
			pythonOSes = append(pythonOSes, img.BaseOS)
		}
	}
	assert.Contains(t, pythonOSes, "Azure Linux")
	assert.Contains(t, pythonOSes, "Ubuntu")
}

func TestGenerateJSONReport_RankResetsPerOSGroup(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	images := []domain.ImageRecord{
		{
			Name: "azl-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
			BaseOSName: "azurelinux", TotalVulnerabilities: 2,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
		{
			Name: "azl-python:3.11", Registry: "r", Repository: "repo2", Tag: "3.11",
			BaseOSName: "azurelinux", TotalVulnerabilities: 5,
			Languages: []domain.Language{{Language: "python", Version: "3.11"}},
		},
		{
			Name: "ubuntu-python:3.12", Registry: "r", Repository: "repo3", Tag: "3.12",
			BaseOSName: "ubuntu", TotalVulnerabilities: 3,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		},
	}

	for i := range images {
		require.NoError(t, repo.InsertImage(&images[i]))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 3)

	// Azure Linux images: rank 1, 2
	assert.Equal(t, 1, report.Images[0].Rank)
	assert.Equal(t, "Azure Linux", report.Images[0].BaseOS)
	assert.Equal(t, 2, report.Images[1].Rank)
	assert.Equal(t, "Azure Linux", report.Images[1].BaseOS)

	// Ubuntu: rank resets to 1
	assert.Equal(t, 1, report.Images[2].Rank)
	assert.Equal(t, "Ubuntu", report.Images[2].BaseOS)
}

func TestGenerateJSONReport_TopNLimit(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	for i := 0; i < 5; i++ {
		img := &domain.ImageRecord{
			Name: "azl-py:" + string(rune('a'+i)), Registry: "r", Repository: "repo", Tag: string(rune('a' + i)),
			BaseOSName: "azurelinux", TotalVulnerabilities: i,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		}
		require.NoError(t, repo.InsertImage(img))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 3, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Len(t, report.Images, 3)
	assert.Equal(t, 3, report.TopN)
}

func TestGenerateJSONReport_ZeroTopNReturnsAll(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	for i := 0; i < 5; i++ {
		img := &domain.ImageRecord{
			Name: "azl-py:" + string(rune('a'+i)), Registry: "r", Repository: "repo", Tag: string(rune('a' + i)),
			BaseOSName: "azurelinux", TotalVulnerabilities: i,
			Languages: []domain.Language{{Language: "python", Version: "3.12"}},
		}
		require.NoError(t, repo.InsertImage(img))
	}

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 0, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	assert.Len(t, report.Images, 5)
	assert.Equal(t, 0, report.TopN)
}

func TestGenerateJSONReport_EmptyDB(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	_, err := os.Stat(outPath)
	assert.True(t, os.IsNotExist(err), "empty DB should not produce a report file")
}

func TestGenerateJSONReport_IncludesHelperFields(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "mcr.microsoft.com/azurelinux/base/python:3.12", Registry: "mcr.microsoft.com",
		Repository: "azurelinux/base/python", Tag: "3.12",
		BaseOSName: "azurelinux", Digest: "sha256:abcdef1234567890",
		SizeBytes: 85000000, TotalVulnerabilities: 1,
		Languages: []domain.Language{{Language: "python", Version: "3.12.1"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 1)
	entry := report.Images[0]
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12@sha256:abcdef1234567890", entry.PinnedReference)
	assert.Equal(t, "mcr.microsoft.com/azurelinux/base/python:3.12", entry.StableTag)
	assert.Equal(t, "FROM mcr.microsoft.com/azurelinux/base/python:3.12@sha256:abcdef1234567890", entry.DockerfileFrom)
	assert.Equal(t, "81.1 MB", entry.SizeHuman)
	assert.Equal(t, "Azure Linux", entry.BaseOS)
	assert.Equal(t, "python", entry.Language)
}

func TestGenerateJSONReport_UnknownOSAsOther(t *testing.T) {
	db, repo := setupTestDB(t)
	defer func() { _ = db.Close() }()

	img := &domain.ImageRecord{
		Name: "unknown-python:3.12", Registry: "r", Repository: "repo", Tag: "3.12",
		TotalVulnerabilities: 1,
		Languages:            []domain.Language{{Language: "python", Version: "3.12"}},
	}
	require.NoError(t, repo.InsertImage(img))

	outPath := filepath.Join(t.TempDir(), "report.json")
	require.NoError(t, GenerateJSONReport(repo, outPath, 10, nil))

	data, err := os.ReadFile(outPath)
	require.NoError(t, err)

	var report JSONReport
	require.NoError(t, json.Unmarshal(data, &report))

	require.Len(t, report.Images, 1)
	assert.Equal(t, "Other", report.Images[0].BaseOS)
}
