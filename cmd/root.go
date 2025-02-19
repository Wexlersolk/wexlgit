package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wgit",
	Short: "a brief description",
	Long:  "a longer description",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to execute execute", "error", err)
	}
}
