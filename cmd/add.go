package cmd

import (
	"fmt"
	"log"

	"github.com/Wexlersolk/wexlgit/internal/handlers"
	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use:   "add",
	Short: "short add",
	Long:  "long for add",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files or directories specified. Usage: wgit add <file1> <dir1> ...")
			return
		}

		if err := handlers.AddExecute(args); err != nil {
			log.Fatalf("Failed to add files or directories: %v", err)
		}
	}}

func init() {
	rootCmd.AddCommand(&addCmd)
}
