package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	log "github.com/sirupsen/logrus"
)

// JSONReport is the top-level structure for the JSON report.
type JSONReport struct {
	GeneratedAt string                  `json:"generatedAt"`
	TopN        int                     `json:"topN"`
	Languages   []JSONLanguageSection   `json:"languages"`
}

// JSONLanguageSection holds recommended images for a single language.
type JSONLanguageSection struct {
	Language string           `json:"language"`
	Images   []JSONImageEntry `json:"images"`
}

// JSONImageEntry represents a single recommended image in the JSON report.
type JSONImageEntry struct {
	Rank                    int    `json:"rank"`
	Name                    string `json:"name"`
	Version                 string `json:"version"`
	CriticalVulnerabilities int    `json:"criticalVulnerabilities"`
	HighVulnerabilities     int    `json:"highVulnerabilities"`
	TotalVulnerabilities    int    `json:"totalVulnerabilities"`
	SizeBytes               int64  `json:"sizeBytes"`
	SizeHuman               string `json:"sizeHuman"`
	Digest                  string `json:"digest"`
}

// GenerateJSONReport produces a JSON recommendations report from the database.
func GenerateJSONReport(repo *database.Repository, outputPath string, topN int) error {
	languages, err := repo.QueryLanguages()
	if err != nil {
		return fmt.Errorf("querying languages: %w", err)
	}

	if len(languages) == 0 {
		log.Warn("No languages found in database; JSON report not generated.")
		return nil
	}

	report := JSONReport{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		TopN:        topN,
	}

	for _, lang := range languages {
		images, err := repo.QueryTopImages(lang, topN)
		if err != nil {
			log.Warnf("Failed to query images for %s: %v", lang, err)
			continue
		}

		if len(images) == 0 {
			continue
		}

		section := JSONLanguageSection{
			Language: strings.ToLower(lang),
		}

		for idx, img := range images {
			section.Images = append(section.Images, JSONImageEntry{
				Rank:                    idx + 1,
				Name:                    img.Name,
				Version:                 img.Version,
				CriticalVulnerabilities: img.CriticalVulnerabilities,
				HighVulnerabilities:     img.HighVulnerabilities,
				TotalVulnerabilities:    img.TotalVulnerabilities,
				SizeBytes:               img.SizeBytes,
				SizeHuman:               HumanSize(img.SizeBytes),
				Digest:                  img.Digest,
			})
		}

		report.Languages = append(report.Languages, section)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling JSON report: %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return fmt.Errorf("writing JSON report: %w", err)
	}

	log.Infof("Wrote JSON recommendations to %s", outputPath)

	return nil
}
