package start

import (
	// "fmt"
	"crud-toy/internal/app/service/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Start the CRUD server",
	Long:    `Start the CRUD server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

//GetCmd allows the command to be accessed from rootCmd
func GetCmd() *cobra.Command {
	return startCmd
}
