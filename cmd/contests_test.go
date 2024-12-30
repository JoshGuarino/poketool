package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContestsCmd(t *testing.T) {
	err := contestsCmd.Execute()
	assert.Equal(t, nil, err, "Expected contests command to execute successfully")
}
