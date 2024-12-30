package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonCmd(t *testing.T) {
	err := pokedexCmd.Execute()
	assert.Equal(t, nil, err, "Expected pokemon command to execute successfully")
}
