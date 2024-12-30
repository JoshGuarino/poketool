package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachinesCmd(t *testing.T) {
	err := machinesCmd.Execute()
	assert.Equal(t, nil, err, "Expected machines command to execute successfully")
}
