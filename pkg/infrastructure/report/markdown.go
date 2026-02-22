package report

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/manisbindra/sbi/pkg/domain"
	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	log "github.com/sirupsen/logrus"
)

// GenerateReport produces a markdown recommendations report from the database.
func GenerateReport(repo *database.Repository, outputPath string, topN int, repoCfg *domain.RepositoryConfig) error {
	languages, err := repo.QueryLanguages()
	if err != nil {
		return fmt.Errorf("querying languages: %w", err)
	}

	if len(languages) == 0 {
		log.Warn("No languages found in database; report not generated.")
		return nil
	}

	ts := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	var sb strings.Builder

	sb.WriteString("# Daily Recommended Images by Language\n\n")
	fmt.Fprintf(&sb, "_Generated: %s. Criteria: lowest critical → high → total vulnerabilities → size. Top %d per language._\n\n", ts, topN)
	sb.WriteString("**Note:** Image sizes are based on Linux amd64 platform as reported by `docker images` on GitHub runners. Actual sizes may vary significantly on other platforms (macOS, Windows, etc.).\n\n")

	if repoCfg != nil {
		writeScannedRepos(&sb, repoCfg)
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

		sb.WriteString(fmt.Sprintf("## %s\n\n", strings.Title(lang))) //nolint:staticcheck
		sb.WriteString("| Rank | Image | Version | Crit | High | Total | Size | Digest |\n")
		sb.WriteString("|------|-------|---------|------|------|-------|------|--------|\n")

		for idx, img := range images {
			version := img.Version
			if version == "" {
				version = "-"
			}

			fmt.Fprintf(&sb, "| %d | `%s` | %s | %d | %d | %d | %s | `%s` |\n",
				idx+1, img.Name, version,
				img.CriticalVulnerabilities, img.HighVulnerabilities, img.TotalVulnerabilities,
				HumanSize(img.SizeBytes), FormatDigest(img.Digest))
		}

		sb.WriteString("\n")
	}

	// Ensure output directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(sb.String()), 0o644); err != nil {
		return fmt.Errorf("writing report: %w", err)
	}

	log.Infof("Wrote daily recommendations to %s", outputPath)

	return nil
}

// writeScannedRepos appends the scanned repositories section to the report.
func writeScannedRepos(sb *strings.Builder, cfg *domain.RepositoryConfig) {
	sb.WriteString("## Scanned Repositories and Images\n\n")

	var totalImages int
	for _, group := range cfg.Repositories {
		totalImages += len(group.Images)
	}

	fmt.Fprintf(sb, "This report includes analysis from **%d configured sources** across %d groups:\n\n",
		totalImages, len(cfg.Repositories))

	for _, group := range cfg.Repositories {
		if group.Description != "" {
			fmt.Fprintf(sb, "**%s:**\n\n", group.Description)
		}

		for _, img := range group.Images {
			fmt.Fprintf(sb, "- `%s`\n", img)
		}

		sb.WriteString("\n")
	}
}

// HumanSize converts bytes to a human-readable size string.
func HumanSize(numBytes int64) string {
	if numBytes <= 0 {
		return "-"
	}

	mb := float64(numBytes) / (1024 * 1024)
	if mb < 1024 {
		return fmt.Sprintf("%.1f MB", mb)
	}

	gb := mb / 1024

	return fmt.Sprintf("%.2f GB", gb)
}

// FormatDigest returns a shortened digest string for display.
func FormatDigest(digest string) string {
	if digest == "" {
		return ""
	}

	if strings.HasPrefix(digest, "sha256:") {
		hashPart := digest[7:]
		if len(hashPart) > 12 {
			return fmt.Sprintf("sha256:%s", hashPart[:12])
		}

		return digest
	}

	if len(digest) > 19 {
		return digest[:19]
	}

	return digest
}

// FormatRecommendedImages formats a list of recommended images for a given language (used externally).
func FormatRecommendedImages(lang string, images []domain.RecommendedImage) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("## %s\n\n", strings.Title(lang))) //nolint:staticcheck
	sb.WriteString("| Rank | Image | Version | Crit | High | Total | Size | Digest |\n")
	sb.WriteString("|------|-------|---------|------|------|-------|------|--------|\n")

	for idx, img := range images {
		version := img.Version
		if version == "" {
			version = "-"
		}

		fmt.Fprintf(&sb, "| %d | `%s` | %s | %d | %d | %d | %s | `%s` |\n",
			idx+1, img.Name, version,
			img.CriticalVulnerabilities, img.HighVulnerabilities, img.TotalVulnerabilities,
			HumanSize(img.SizeBytes), FormatDigest(img.Digest))
	}

	return sb.String()
}
