package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncountersCmd(t *testing.T) {
	err := encountersCmd.Execute()
	assert.Equal(t, nil, err, "Expected encounters command to execute successfully")
}
