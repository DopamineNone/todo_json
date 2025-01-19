package list

import (
	"todo_json/utils"

	"github.com/spf13/cobra"
)

const (
	TODO        = "todo"
	IN_PROGRESS = "inprogress"
	DONE        = "done"
	ALL         = "all"
)

var ListCmd = &cobra.Command{
	Use:       "list {todo|in-progress|done|all}",
	Short:     "list the tasks of specified type",
	ValidArgs: []string{TODO, IN_PROGRESS, DONE, ALL},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:       ListFunc,
}

func ListFunc(cmd *cobra.Command, args []string) {
	utils.ListTasks(args[0])
}
