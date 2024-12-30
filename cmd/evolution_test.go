package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvolutionCmd(t *testing.T) {
	err := evolutionCmd.Execute()
	assert.Equal(t, nil, err, "Expected evolution command to execute successfully")
}
