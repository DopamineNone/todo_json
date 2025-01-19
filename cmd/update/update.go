package update

import (
	"strconv"
	"todo_json/utils"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update {id} {task}",
	Short: "update specified task content",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), utils.CheckNum(1)),
	Run:   UpdateFunc,
}

func UpdateFunc(cmd *cobra.Command, args []string) {
	id, _ := strconv.Atoi(args[0])
	utils.UpdateTask(id, args[1])
}

func init() {
	UpdateCmd.Flags().IntP("test", "d", 0, "dont known")
	UpdateCmd.PersistentFlags()
}
