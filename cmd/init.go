package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = cobra.Command{
	Use:   "init",
	Short: "init it is",
	Long:  "it is init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(&initCmd)
}
