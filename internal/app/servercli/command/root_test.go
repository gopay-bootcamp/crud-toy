package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmdUsage(t *testing.T) {
	assert.Equal(t, "proctord", rootCmd.Use)
	assert.Equal(t, "proctord - Handle executing jobs and maintaining their configuration", rootCmd.Short)
	assert.Equal(t, "proctord - Handle executing jobs and maintaining their configuration", rootCmd.Long)
}
