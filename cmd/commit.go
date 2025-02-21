package cmd

import (
	"log"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
	"github.com/spf13/cobra"
)

var commitCmd = cobra.Command{
	Use:   "commit",
	Short: "short explanation of commit",
	Long:  "long explanation of commit",
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.CommitExecute("test message"); err != nil {
			log.Fatalf("Failed to commit: %v", err)
		}
	}}

func init() {
	rootCmd.AddCommand(&commitCmd)
}
