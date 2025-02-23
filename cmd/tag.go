package cmd

import (
	"log"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
	"github.com/spf13/cobra"
)

var tagCmd = cobra.Command{
	Use:   "tag",
	Short: "short explanation of tag",
	Long:  "long explanation of tag",
	Run: func(cmd *cobra.Command, args []string) {
		if err := handlers.TagExecute("test_tag"); err != nil {
			log.Fatalf("Failed to commit: %v", err)
		}
	}}

func init() {
	rootCmd.AddCommand(&tagCmd)
}
