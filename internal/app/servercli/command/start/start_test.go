package start

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCmdUsage(t *testing.T) {
	assert.Equal(t, "start", startCmd.Use)
}

func TestListCmdHelp(t *testing.T) {
	assert.Equal(t, "Start the proctor server", startCmd.Short)
	assert.Equal(t, "Start the proctor server", startCmd.Long)
}

func TestListCmdRun(t *testing.T) {
	//TODO write test when Run method is implemented
}
