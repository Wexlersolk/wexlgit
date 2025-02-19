package cmd

import (
	"log"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
	"github.com/spf13/cobra"
)

var initCmd = cobra.Command{
	Use:   "init",
	Short: "init it is",
	Long:  "it is init",
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.InitRepository(); err != nil {
			log.Fatalf("Failed to initialize repository: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(&initCmd)
}
