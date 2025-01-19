package mark

import (
	"strconv"
	"todo_json/utils"

	"github.com/spf13/cobra"
)

const (
	IN_PROGRESS = "inprogress"
	DONE        = "done"
)

var (
	in   bool
	done bool
)

var MarkCmd = &cobra.Command{
	Use:   "mark {id}",
	Short: "mark the specified task in some state",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), utils.CheckNum(1)),
	Run:   MarkFunc,
}

func MarkFunc(cmd *cobra.Command, args []string) {
	id, _ := strconv.Atoi(args[0])
	if in {
		utils.MarkTasks(id, "inprogress")
	} else if done {
		utils.MarkTasks(id, "done")
	}
}

func init() {
	MarkCmd.Flags().BoolVar(&in, IN_PROGRESS, false, "mark task with in-progress status")
	MarkCmd.Flags().BoolVar(&done, DONE, false, "mark task with done status")
	MarkCmd.MarkFlagsMutuallyExclusive(IN_PROGRESS, DONE)
	MarkCmd.MarkFlagsOneRequired(IN_PROGRESS, DONE)
}
