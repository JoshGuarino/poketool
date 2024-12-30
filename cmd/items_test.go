package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemsCmd(t *testing.T) {
	err := itemsCmd.Execute()
	assert.Equal(t, nil, err, "Expected items command to execute successfully")
}
