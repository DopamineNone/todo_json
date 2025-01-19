package add

import (
	"todo_json/utils"

	"github.com/spf13/cobra"
)

// AddCmd, add a new task into task json list
var AddCmd = &cobra.Command{
	Use:   "add {task}",
	Short: "add a new task",
	Args:  cobra.ExactArgs(1),
	Run:   addFunc,
}

func addFunc(cmd *cobra.Command, args []string) {
	// TO DO:
	utils.AddTasks(args[0])
}
