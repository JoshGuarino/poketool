package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBerriesCmd(t *testing.T) {
	err := berriesCmd.Execute()
	assert.Equal(t, nil, err, "Expected berries command to execute successfully")
}
