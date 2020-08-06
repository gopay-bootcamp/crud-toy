package grpcStart

import (
	// "fmt"
	"crud-toy/internal/server/grpc"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "grpcStart",
	Aliases: []string{"gs"},
	Short:   "Start the grpc server",
	Long:    `Start the grpc server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

//GetCmd allows the command to be accessed from rootCmd
func GetCmd() *cobra.Command {
	return startCmd
}
