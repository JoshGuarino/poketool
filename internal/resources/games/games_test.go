package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var games IGames = NewGames()

func TestGetGeneration(t *testing.T) {
	rById, _ := games.GetGeneration("1")
	rByName, _ := games.GetGeneration("generation-i")
	_, err := games.GetGeneration("test")
	assert.Equal(t, "generation-i", rById.Name, "Expected to receive 'generation-i'.")
	assert.Equal(t, "generation-i", rByName.Name, "Expected to receive 'generation-i'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetGenerationList(t *testing.T) {
	r, _ := games.GetGenerationList(20, 0)
	assert.Equal(t, "generation-i", r.Results[0].Name, "Expected to have a list of 'generation' resource.")
}

func TestGetPokedex(t *testing.T) {
	rById, _ := games.GetPokedex("1")
	rByName, _ := games.GetPokedex("national")
	_, err := games.GetPokedex("test")
	assert.Equal(t, "national", rById.Name, "Expected to receive 'national'.")
	assert.Equal(t, "national", rByName.Name, "Expected to receive 'national'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetPokedexList(t *testing.T) {
	r, _ := games.GetPokedexList(20, 0)
	assert.Equal(t, "national", r.Results[0].Name, "Expected to have a list of 'pokedex' resource.")
}

func TestGetVersion(t *testing.T) {
	rById, _ := games.GetVersion("1")
	rByName, _ := games.GetVersion("red")
	_, err := games.GetVersion("test")
	assert.Equal(t, "red", rById.Name, "Expected to receive 'red'.")
	assert.Equal(t, "red", rByName.Name, "Expected to receive 'red'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetVersionList(t *testing.T) {
	r, _ := games.GetVersionList(20, 0)
	assert.Equal(t, "red", r.Results[0].Name, "Expected to have a list of 'version' resource.")
}

func TestGetVersionGroup(t *testing.T) {
	rById, _ := games.GetVersionGroup("1")
	rByName, _ := games.GetVersionGroup("red-blue")
	_, err := games.GetVersionGroup("test")
	assert.Equal(t, "red-blue", rById.Name, "Expected to receive 'red-blue'.")
	assert.Equal(t, "red-blue", rByName.Name, "Expected to receive 'red-blue'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetVersionGroupList(t *testing.T) {
	r, _ := games.GetVersionGroupList(20, 0)
	assert.Equal(t, "red-blue", r.Results[0].Name, "Expected to have a list of 'version group' resource.")
}
