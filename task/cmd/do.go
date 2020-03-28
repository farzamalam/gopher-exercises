package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do is used to mark task that are completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Invalid argument :", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println("ids :", ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
