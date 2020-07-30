package delete

import (
	"crud-toy/internal/app/cli/daemon"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete procs by id",
	Long:  `Delete procs by unique id`,
	Run: func(cmd *cobra.Command, args []string) {
		err := daemon.DeleteProcs(args)
		if err != nil {
			fmt.Print(err)
		}
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return deleteCmd
}

func init() {

}
