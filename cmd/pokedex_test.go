package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokedexCmd(t *testing.T) {
	err := pokedexCmd.Execute()
	assert.Equal(t, nil, err, "Expected pokedex command to execute successfully")
}
