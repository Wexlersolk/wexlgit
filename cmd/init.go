package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = cobra.Command{
	Use:   "init",
	Short: "init it is",
	Long:  "it is init",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(".wgit")
		if os.IsNotExist(err) {
			fmt.Println("it doesnt exist")
			// create it
		} else {
			fmt.Println("it already exists")
		}
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(&initCmd)
}
