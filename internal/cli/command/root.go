package command

import (
	"crud-toy/internal/cli/command/create"
	"crud-toy/internal/cli/command/delete"
	"crud-toy/internal/cli/command/find"
	"crud-toy/internal/cli/command/readall"
	"crud-toy/internal/cli/command/update"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "Store",
	Short: "Welcome to the toy store",
	Long:  `You can update your store details here.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(find.GetCmd())
	rootCmd.AddCommand(create.GetCmd())
	rootCmd.AddCommand(update.GetCmd())
	rootCmd.AddCommand(delete.GetCmd())
	rootCmd.AddCommand(readall.GetCmd())

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.proctor-command.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".proctor-command")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
