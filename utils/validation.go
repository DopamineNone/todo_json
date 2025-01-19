package utils

import (
	"strconv"

	"github.com/spf13/cobra"
)

func CheckNum(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if n > len(args) {
			return nil
		}
		_, err := strconv.Atoi(args[n-1])
		return err
	}
}
