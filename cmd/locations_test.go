package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationsCmd(t *testing.T) {
	err := locationsCmd.Execute()
	assert.Equal(t, nil, err, "Expected locations command to execute successfully")
}
