package machines

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var machines IMachines = Machines{}

func TestGetMachine(t *testing.T) {
	rById, _ := machines.GetMachine("1")
	rFail, _ := machines.GetMachine("test")
	assert.Equal(t, 1, rById.ID, "Expected to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expected to receive an empty result.")
}

func TestGetEvolutionChainList(t *testing.T) {
	r := machines.GetMachineList()
	assert.Equal(t, "https://pokeapi.co/api/v2/machine/1/", r.Results[0].URL, "Expected to have a list of 'machine' resource.")
}
