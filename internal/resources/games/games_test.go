package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var games IGames = Games{}

func TestGetGeneration(t *testing.T) {
	rById, _ := games.GetGeneration("1")
	rByName, _ := games.GetGeneration("generation-i")
	rFail, _ := games.GetGeneration("test")
	assert.Equal(t, "generation-i", rById.Name, "Expected to receive 'generation-i'.")
	assert.Equal(t, "generation-i", rByName.Name, "Expected to receive 'generation-i'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetGenerationList(t *testing.T) {
	r := games.GetGenerationList()
	assert.Equal(t, "generation-i", r.Results[0].Name, "Expected to have a list of 'generation' resource.")
}

func TestGetPokedex(t *testing.T) {
	rById, _ := games.GetPokedex("1")
	rByName, _ := games.GetPokedex("national")
	rFail, _ := games.GetPokedex("test")
	assert.Equal(t, "national", rById.Name, "Expected to receive 'national'.")
	assert.Equal(t, "national", rByName.Name, "Expected to receive 'national'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokedexList(t *testing.T) {
	r := games.GetPokedexList()
	assert.Equal(t, "national", r.Results[0].Name, "Expected to have a list of 'pokedex' resource.")
}

func TestGetVersion(t *testing.T) {
	rById, _ := games.GetVersion("1")
	rByName, _ := games.GetVersion("red")
	rFail, _ := games.GetVersion("test")
	assert.Equal(t, "red", rById.Name, "Expected to receive 'red'.")
	assert.Equal(t, "red", rByName.Name, "Expected to receive 'red'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetVersionList(t *testing.T) {
	r := games.GetVersionList()
	assert.Equal(t, "red", r.Results[0].Name, "Expected to have a list of 'version' resource.")
}

func TestGetVersionGroup(t *testing.T) {
	rById, _ := games.GetVersionGroup("1")
	rByName, _ := games.GetVersionGroup("red-blue")
	rFail, _ := games.GetVersionGroup("test")
	assert.Equal(t, "red-blue", rById.Name, "Expected to receive 'red-blue'.")
	assert.Equal(t, "red-blue", rByName.Name, "Expected to receive 'red-blue'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetVersionGroupList(t *testing.T) {
	r := games.GetVersionGroupList()
	assert.Equal(t, "red-blue", r.Results[0].Name, "Expected to have a list of 'version group' resource.")
}
