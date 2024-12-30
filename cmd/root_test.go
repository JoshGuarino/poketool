package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	err := rootCmd.Execute()
	assert.Equal(t, nil, err, "Expected root command to execute successfully")
}
