package evolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var evolution IEvolution = NewEvolution()

func TestGetEvolutionChain(t *testing.T) {
	rById, _ := evolution.GetEvolutionChain("1")
	_, err := evolution.GetEvolutionChain("test")
	assert.Equal(t, 1, rById.ID, "Expected to receive id of '1'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionChainList(t *testing.T) {
	r, _ := evolution.GetEvolutionChainList(20, 0)
	assert.Equal(t, "https://staging.pokeapi.co/api/v2/evolution-chain/1/", r.Results[0].URL, "Expected to have a list of 'evolution chain' resource.")
}

func TestGetEvolutionTrigger(t *testing.T) {
	rById, _ := evolution.GetEvolutionTrigger("1")
	rByName, _ := evolution.GetEvolutionTrigger("level-up")
	_, err := evolution.GetEvolutionTrigger("test")
	assert.Equal(t, "level-up", rById.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "level-up", rByName.Name, "Expected to receive 'level-up'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEvolutionTriggerList(t *testing.T) {
	r, _ := evolution.GetEvolutionTriggerList(20, 0)
	assert.Equal(t, "level-up", r.Results[0].Name, "Expected to have a list of 'evolution trigger' resource.")
}
