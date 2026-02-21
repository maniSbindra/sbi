package main

import (
	"github.com/spf13/cobra"
)

var (
	flgDBPath         string
	flgConfigDir      string
	flgOutputPath     string
	flgTopN           int
	flgMaxTags        int
	flgComprehensive  bool
	flgNoCleanup      bool
	flgUpdateExisting bool
	flgVerbose        bool
	flgDebug          bool
)

// NewRootCommand creates the root cobra command.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "daily-recommendations",
		Short: "Scan container images and generate daily security recommendations by language",
		Long: `Scans container registries (MCR, Docker Hub, etc.) for base images,
analyzes them with Syft (SBOM) and Trivy (vulnerabilities), stores results in SQLite,
and generates a markdown report ranking images by security posture per language.`,
	}

	rootCmd.PersistentFlags().StringVar(&flgDBPath, "database", "azure_linux_images.db", "Path to SQLite database")
	rootCmd.PersistentFlags().StringVar(&flgConfigDir, "config-dir", "config", "Path to configuration directory")
	rootCmd.PersistentFlags().StringVar(&flgOutputPath, "output", "docs/daily_recommendations.md", "Path to output report file")
	rootCmd.PersistentFlags().IntVar(&flgTopN, "top-n", 10, "Number of top images per language")
	rootCmd.PersistentFlags().BoolVarP(&flgVerbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&flgDebug, "debug", "d", false, "Enable debug output")

	rootCmd.AddCommand(NewScanCommand())
	rootCmd.AddCommand(NewReportCommand())
	rootCmd.AddCommand(NewResetDBCommand())

	return rootCmd
}
