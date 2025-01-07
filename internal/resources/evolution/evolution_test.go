package evolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvolutionChain(t *testing.T) {
	rById, _ := GetEvolutionChain("1")
	rFail, _ := GetEvolutionChain("test")
	assert.Equal(t, 1, rById.ID, "Expected to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expected to receive an empty result.")
}

func TestGetEvolutionChainList(t *testing.T) {
	r := GetEvolutionChainList()
	assert.Equal(t, "https://pokeapi.co/api/v2/evolution-chain/1/", r.Results[0].URL, "Expected to have a list of 'evolution chain' resource.")
}

func TestGetEvolutionTrigger(t *testing.T) {
	rById, _ := GetEvolutionTrigger("1")
	rByName, _ := GetEvolutionTrigger("level-up")
	rFail, _ := GetEvolutionTrigger("test")
	assert.Equal(t, "level-up", rById.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "level-up", rByName.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEvolutionTriggerList(t *testing.T) {
	r := GetEvolutionTriggerList()
	assert.Equal(t, "level-up", r.Results[0].Name, "Expected to have a list of 'evolution trigger' resource.")
}
