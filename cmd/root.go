package cmd

import (
	"fmt"
	"os"
	"todo_json/cmd/add"
	"todo_json/cmd/delete"
	"todo_json/cmd/list"
	"todo_json/cmd/mark"
	"todo_json/cmd/update"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Todo App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello ToDo! v1.0.0")
		str, _ := cmd.NonInheritedFlags().GetString("name")
		fmt.Println(str)
	},
	TraverseChildren: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(add.AddCmd, delete.DeleteCmd, update.UpdateCmd, list.ListCmd, mark.MarkCmd)
}
