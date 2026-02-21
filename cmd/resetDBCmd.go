package main

import (
	"fmt"

	"github.com/manisbindra/sbi/pkg/infrastructure/database"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewResetDBCommand creates the `reset-db` subcommand.
func NewResetDBCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "reset-db",
		Short: "Clear all data from the database",
		Long:  `Deletes all image, vulnerability, language, and related data from the database.`,
		RunE:  runResetDB,
	}
}

func runResetDB(_ *cobra.Command, _ []string) error {
	setLogLevel()

	db, err := database.OpenDB(flgDBPath)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	repo := database.NewRepository(db)

	if err := repo.ClearDatabase(); err != nil {
		return fmt.Errorf("clearing database: %w", err)
	}

	log.Info("Database cleared successfully")
	fmt.Println("Database cleared successfully")

	return nil
}
