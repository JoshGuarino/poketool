package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamesCmd(t *testing.T) {
	err := gamesCmd.Execute()
	assert.Equal(t, nil, err, "Expected games command to execute successfully")
}
