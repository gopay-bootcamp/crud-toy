package command

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRootCmdUsage(t *testing.T) {
	assert.Equal(t, "proctor-command", rootCmd.Use)
	assert.Equal(t, "proctor is a automation orchestrator", rootCmd.Short)
	assert.Equal(t, "You can give commands to orchestrate your automation here.", rootCmd.Long)
}

func contains(commands []*cobra.Command, commandName string) bool {
	for _, command := range commands {
		if commandName == command.Name() {
			return true
		}
	}
	return false
}

func TestRootCmdSubCommands(t *testing.T) {

	assert.True(t, contains(rootCmd.Commands(), "list"))
}
