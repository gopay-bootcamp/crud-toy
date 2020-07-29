package update

import (
	"crud-toy/internal/app/cli/daemon"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "update",
	Short: "Update procs by id",
	Long:  `Update procs by unique id`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon.UpdateProcs(args[0], args[1], args[2])
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return findCmd
}

func init() {

}
