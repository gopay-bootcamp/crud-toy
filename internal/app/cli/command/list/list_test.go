package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCmdUsage(t *testing.T) {
	assert.Equal(t, "list", listCmd.Use)
}

func TestListCmdHelp(t *testing.T) {
	assert.Equal(t, "List of all procs", listCmd.Short)
	assert.Equal(t, "List of all procs available for execution", listCmd.Long)
}

func TestListCmdRun(t *testing.T) {
	//TODO write test when Run method is implemented
}
