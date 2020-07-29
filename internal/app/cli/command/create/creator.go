package create

import (
	"crud-toy/internal/app/cli/daemon"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create procs",
	Long:  `Create procs by giving name, author`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon.CreateProcs(args[0], args[1])
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return createCmd
}

func init() {

}
