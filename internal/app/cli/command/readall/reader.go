package readall

import (
	"crud-toy/internal/app/cli/daemon"

	"github.com/spf13/cobra"
)

var readAllCmd = &cobra.Command{
	Use:   "readall",
	Short: "Show all Procs",
	Long:  `See all procs in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon.ReadAllProcs()
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return readAllCmd
}

func init() {

}
