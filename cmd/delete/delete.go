package delete

import (
	"strconv"
	"todo_json/utils"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the specified task in some state",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), utils.CheckNum(1)),
	Run:   DeleteFunc,
}

func DeleteFunc(cmd *cobra.Command, args []string) {
	id, _ := strconv.Atoi(args[0])
	utils.DeleteTask(id)
}
