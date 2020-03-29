package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/farzamalam/gopher-exercises/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add is used to add new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong.", err.Error())
			os.Exit(1)
		}
		fmt.Printf("New task \"%s\" has been added.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
