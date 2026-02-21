package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.ErrorLevel
	}

	log.SetLevel(logLevel)

	rootCmd := NewRootCommand()
	rootCmd.Version = fmt.Sprintf(": %s, commit: %s, date: %s", version, commit, date)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
