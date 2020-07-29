package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all procs",
	Long:  `List of all procs available for execution`,
	Run: func(cmd *cobra.Command, args []string) {
		procsList := "list called"
		fmt.Println(procsList)
	},
}

// execute the listCmd
func GetCmd() *cobra.Command {
	return listCmd
}

func init() {
	listCmd.Flags().BoolP("list-search", "s", false, "Search a particular command in the list")
}
