package evolution

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{evolution: Evolution{}}

func TestGetList(t *testing.T) {
	rChain := controller.GetList("Evolution Chain")
	rTrigger := controller.GetList("Evolution Trigger")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rChain.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rTrigger.Results, "Expected to have array of type 'Result' struct.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rChain, _ := controller.GetSpecific("Evolution Chain", "1")
	rTrigger, _ := controller.GetSpecific("Evolution Trigger", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.EvolutionChain{}, rChain, "Expected to have type 'EvolutionChain' struct.")
	assert.IsType(t, structs.EvolutionTrigger{}, rTrigger, "Expected to have type 'EvolutionTrigger' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
