package evolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var evolution IEvolution = Evolution{}

func TestGetEvolutionChain(t *testing.T) {
	rById, _ := evolution.GetEvolutionChain("1")
	rFail, _ := evolution.GetEvolutionChain("test")
	assert.Equal(t, 1, rById.ID, "Expected to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expected to receive an empty result.")
}

func TestGetEvolutionChainList(t *testing.T) {
	r := evolution.GetEvolutionChainList()
	assert.Equal(t, "https://pokeapi.co/api/v2/evolution-chain/1/", r.Results[0].URL, "Expected to have a list of 'evolution chain' resource.")
}

func TestGetEvolutionTrigger(t *testing.T) {
	rById, _ := evolution.GetEvolutionTrigger("1")
	rByName, _ := evolution.GetEvolutionTrigger("level-up")
	rFail, _ := evolution.GetEvolutionTrigger("test")
	assert.Equal(t, "level-up", rById.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "level-up", rByName.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEvolutionTriggerList(t *testing.T) {
	r := evolution.GetEvolutionTriggerList()
	assert.Equal(t, "level-up", r.Results[0].Name, "Expected to have a list of 'evolution trigger' resource.")
}
