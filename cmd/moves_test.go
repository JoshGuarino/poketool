package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesCmd(t *testing.T) {
	err := movesCmd.Execute()
	assert.Equal(t, nil, err, "Expected moves command to execute successfully")
}
