package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func Version() string {
	return version
}

var rootCmd = &cobra.Command{
	Use:          "relx",
	Short:        "Automated release and changelog tool",
	Long:         "relx analyzes your Conventional Commits, determines the next version, updates CHANGELOG.md, tags the release and publishes to GitHub.",
	Version:      version,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
