package create

import (
	"crud-toy/internal/cli/daemon"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create procs",
	Long:  `Create procs by giving name, author`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon.CreateProcs(args)
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return createCmd
}

func init() {

}
