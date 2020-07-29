package find

import (
	"crud-toy/internal/app/cli/daemon"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find procs by id",
	Long:  `Search procs by unique id`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon.FindProcs(args[0])
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return findCmd
}

func init() {

}
