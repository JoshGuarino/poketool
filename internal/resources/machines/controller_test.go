package machines

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{machines: Machines{}}

func TestGetList(t *testing.T) {
	rMachines := controller.GetList("Machine")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rMachines.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rMachines.Results, "Expected to not have an empty array.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rMachines, _ := controller.GetSpecific("Machine", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Machine{}, rMachines, "Expected to have type 'Machine' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
