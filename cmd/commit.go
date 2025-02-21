package cmd

import (
	"github.com/spf13/cobra"
)

var commitCmd = cobra.Command{
	Use:   "commit",
	Short: "short explanation of commit",
	Long:  "long explanation of commit",
	Run: func(cmd *cobra.Command, args []string) {
	}}

func init() {
	rootCmd.AddCommand(&commitCmd)
}
