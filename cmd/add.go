package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use:   "add",
	Short: "short add",
	Long:  "long for add",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(&addCmd)
}
