package command

import (
	"crud-toy/internal/app/service/infra/logger"
	"crud-toy/internal/app/servercli/command/start"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "proctord",
	Short: "proctord - Handle executing jobs and maintaining their configuration",
	Long:  `proctord - Handle executing jobs and maintaining their configuration`,
}

//Execute allows the rootCmd to be executed from outside
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logger.Setup()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(start.GetCmd())
}
