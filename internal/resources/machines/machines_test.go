package machines

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMachine(t *testing.T) {
	rById, _ := GetMachine("1")
	rFail, _ := GetMachine("test")
	assert.Equal(t, 1, rById.ID, "Expected to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expected to receive an empty result.")
}

func TestGetEvolutionChainList(t *testing.T) {
	r := GetMachineList()
	assert.Equal(t, "https://pokeapi.co/api/v2/machine/1/", r.Results[0].URL, "Expected to have a list of 'machine' resource.")
}
