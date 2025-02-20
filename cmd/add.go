package cmd

import (
	"log"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use:   "add",
	Short: "short add",
	Long:  "long for add",
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.AddExecute(args); err != nil {
			log.Fatalf("Failed to add files: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(&addCmd)
}
