package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add is used to add new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("New task \"%s\" has been added.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
