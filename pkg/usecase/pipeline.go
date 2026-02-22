package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/manisbindra/sbi/pkg/domain"
	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	"github.com/manisbindra/sbi/pkg/infrastructure/report"
	"github.com/manisbindra/sbi/pkg/infrastructure/scanner"
	log "github.com/sirupsen/logrus"
)

// Pipeline orchestrates the full scan and report generation workflow.
type Pipeline struct {
	config   domain.ScanConfig
	repoCfg  domain.RepositoryConfig
	repo     *database.Repository
	analyzer *scanner.ImageAnalyzer
	registry *scanner.RegistryScanner
}

// NewPipeline creates a new Pipeline.
func NewPipeline(config domain.ScanConfig, repo *database.Repository) (*Pipeline, error) {
	repoCfg, err := LoadRepositoryConfig(config.ConfigDir)
	if err != nil {
		return nil, fmt.Errorf("loading repository config: %w", err)
	}

	// CLI flag overrides config default if explicitly set
	maxTags := config.MaxTagsPerRepo
	if maxTags == 0 && repoCfg.Defaults.MaxTags > 0 {
		maxTags = repoCfg.Defaults.MaxTags
	}

	config.MaxTagsPerRepo = maxTags

	return &Pipeline{
		config:   config,
		repoCfg:  repoCfg,
		repo:     repo,
		analyzer: scanner.NewImageAnalyzer(config.ComprehensiveScan, config.CleanupImages),
		registry: scanner.NewRegistryScanner(repoCfg.Defaults.Registry),
	}, nil
}

// ScanAll loads repositories from config, discovers tags, and scans all images.
func (p *Pipeline) ScanAll() error {
	var allImages []string
	for _, group := range p.repoCfg.Repositories {
		allImages = append(allImages, group.Images...)
	}

	repos, singleImages := scanner.ParseImagePatterns(allImages)
	log.Infof("Found %d repositories and %d single images to scan", len(repos), len(singleImages))

	for _, repo := range repos {
		if err := p.scanRepository(repo); err != nil {
			log.Errorf("Error scanning repository %s: %v", repo, err)
		}
	}

	for _, img := range singleImages {
		if err := p.scanSingleImage(img); err != nil {
			log.Errorf("Error scanning image %s: %v", img, err)
		}
	}

	return nil
}

// ScanImage scans a single image by name.
func (p *Pipeline) ScanImage(imageName string) error {
	return p.scanSingleImage(imageName)
}

// GenerateReport generates both the markdown and JSON recommendations reports.
func (p *Pipeline) GenerateReport() error {
	if err := report.GenerateReport(p.repo, p.config.OutputPath, p.config.TopN, &p.repoCfg); err != nil {
		return err
	}

	jsonPath := strings.TrimSuffix(p.config.OutputPath, ".md") + ".json"

	return report.GenerateJSONReport(p.repo, jsonPath, p.config.TopN, &p.repoCfg)
}

func (p *Pipeline) scanRepository(repo string) error {
	log.Infof("Scanning repository: %s", repo)

	tags, err := p.registry.GetTags(repo)
	if err != nil {
		return fmt.Errorf("getting tags for %s: %w", repo, err)
	}

	filtered := scanner.FilterTags(tags, p.repoCfg.TagFilter)
	limited := scanner.LimitTags(filtered, p.config.MaxTagsPerRepo)
	log.Infof("Repository %s: %d tags found, %d after filtering, %d to scan",
		repo, len(tags), len(filtered), len(limited))

	defaultRegistry := p.repoCfg.Defaults.Registry
	for _, tag := range limited {
		imageName := scanner.BuildFullImageName(defaultRegistry, repo, tag)

		if err := p.scanSingleImage(imageName); err != nil {
			log.Errorf("Error scanning %s: %v", imageName, err)
		}
	}

	return nil
}

func (p *Pipeline) scanSingleImage(imageName string) error {
	if !p.config.UpdateExisting {
		exists, err := p.repo.ImageExists(imageName)
		if err != nil {
			return fmt.Errorf("checking image existence: %w", err)
		}

		if exists {
			log.Infof("Skipping existing image: %s", imageName)
			return nil
		}
	}

	analysis, err := p.analyzer.Analyze(imageName)
	if err != nil {
		return fmt.Errorf("analyzing %s: %w", imageName, err)
	}

	if err := p.repo.InsertImage(&analysis.Image); err != nil {
		return fmt.Errorf("inserting %s: %w", imageName, err)
	}

	log.Infof("Successfully scanned and stored: %s", imageName)

	return nil
}

// LoadRepositoryConfig reads the JSON config from the config directory.
func LoadRepositoryConfig(configDir string) (domain.RepositoryConfig, error) {
	path := configDir + "/repositories.json"

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn("No repositories.json found, using defaults")
			return defaultRepositoryConfig(), nil
		}

		return domain.RepositoryConfig{}, fmt.Errorf("reading %s: %w", path, err)
	}

	var cfg domain.RepositoryConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return domain.RepositoryConfig{}, fmt.Errorf("parsing %s: %w", path, err)
	}

	if cfg.Defaults.Registry == "" {
		cfg.Defaults.Registry = "mcr.microsoft.com"
	}

	if len(cfg.TagFilter.SkipExact) == 0 && len(cfg.TagFilter.ExcludeKeywords) == 0 &&
		len(cfg.TagFilter.ExcludePatterns) == 0 && !cfg.TagFilter.RequireDigit {
		cfg.TagFilter = scanner.DefaultTagFilter()
	}

	return cfg, nil
}

func defaultRepositoryConfig() domain.RepositoryConfig {
	return domain.RepositoryConfig{
		Defaults: domain.ConfigDefaults{
			Registry: "mcr.microsoft.com",
			MaxTags:  0,
		},
		TagFilter: scanner.DefaultTagFilter(),
		Repositories: []domain.RepositoryGroup{
			{
				Description: "Azure Linux base images",
				Images: []string{
					"azurelinux/base/python",
					"azurelinux/base/nodejs",
				},
			},
			{
				Description: "Azure Linux distroless images",
				Images: []string{
					"azurelinux/distroless/base",
					"azurelinux/distroless/python",
					"azurelinux/distroless/nodejs",
				},
			},
		},
	}
}
