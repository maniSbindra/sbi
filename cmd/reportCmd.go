package main

import (
	"strings"

	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	"github.com/manisbindra/sbi/pkg/infrastructure/report"
	"github.com/manisbindra/sbi/pkg/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewReportCommand creates the `report` subcommand.
func NewReportCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "report",
		Short: "Generate the daily recommendations markdown report",
		Long:  `Reads the existing database and generates a markdown report ranking images by language.`,
		RunE:  runReport,
	}
}

func runReport(_ *cobra.Command, _ []string) error {
	setLogLevel()

	db, err := database.OpenDB(flgDBPath)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	repo := database.NewRepository(db)

	// Load repo config for scanned repositories section
	repoCfg, err := usecase.LoadRepositoryConfig(flgConfigDir)
	if err != nil {
		log.Warnf("Could not load repository config: %v", err)
	}

	log.Info("Generating reports...")

	if err := report.GenerateReport(repo, flgOutputPath, flgTopN, &repoCfg); err != nil {
		return err
	}

	jsonPath := strings.TrimSuffix(flgOutputPath, ".md") + ".json"

	return report.GenerateJSONReport(repo, jsonPath, flgTopN, &repoCfg)
}
