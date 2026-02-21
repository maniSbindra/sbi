package main

import (
	"github.com/manisbindra/sbi/pkg/domain"
	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	"github.com/manisbindra/sbi/pkg/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewScanCommand creates the `scan` subcommand.
func NewScanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan container registries and analyze images",
		Long: `Discover image tags from configured registries, pull images,
analyze with Syft and Trivy, and store results in the database.`,
		RunE: runScan,
	}

	cmd.Flags().IntVar(&flgMaxTags, "max-tags", 5, "Maximum tags per repository (0 = all)")
	cmd.Flags().BoolVar(&flgComprehensive, "comprehensive", false, "Enable comprehensive scanning (secrets + misconfigs)")
	cmd.Flags().BoolVar(&flgNoCleanup, "no-cleanup", false, "Keep Docker images after scanning")
	cmd.Flags().BoolVar(&flgUpdateExisting, "update-existing", false, "Rescan existing images")

	return cmd
}

func runScan(cmd *cobra.Command, _ []string) error {
	setLogLevel()

	db, err := database.OpenDB(flgDBPath)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	if err := database.CreateTables(db); err != nil {
		return err
	}

	config := domain.ScanConfig{
		DBPath:            flgDBPath,
		ConfigDir:         flgConfigDir,
		OutputPath:        flgOutputPath,
		TopN:              flgTopN,
		MaxTagsPerRepo:    flgMaxTags,
		ComprehensiveScan: flgComprehensive,
		CleanupImages:     !flgNoCleanup,
		UpdateExisting:    flgUpdateExisting,
		Verbose:           flgVerbose,
	}

	repo := database.NewRepository(db)
	pipeline, err := usecase.NewPipeline(config, repo)
	if err != nil {
		return err
	}

	log.Info("Starting scan...")

	if err := pipeline.ScanAll(); err != nil {
		return err
	}

	log.Info("Scan complete. Generating report...")

	return pipeline.GenerateReport()
}

func setLogLevel() {
	if flgDebug {
		log.SetLevel(log.DebugLevel)
	} else if flgVerbose {
		log.SetLevel(log.InfoLevel)
	}
}
