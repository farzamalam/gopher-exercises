package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add is used to add new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called.")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
