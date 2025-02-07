package berries

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{berries: Berries{}}

func TestGetList(t *testing.T) {
	rBerries := controller.GetList("Berries")
	rFirmnesses := controller.GetList("Berry Firmnesses")
	rFlavors := controller.GetList("Berry Flavors")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rBerries.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rBerries.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rFirmnesses.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rFirmnesses.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rFlavors.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rFlavors.Results, "Expected to not have an empty array.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rBerry, _ := controller.GetSpecific("Berries", "1")
	rFirmness, _ := controller.GetSpecific("Berry Firmnesses", "1")
	rFlavor, _ := controller.GetSpecific("Berry Flavors", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Berry{}, rBerry, "Expected to have type 'Berry' struct.")
	assert.IsType(t, structs.BerryFirmness{}, rFirmness, "Expected to have type 'BerryFirmness' struct.")
	assert.IsType(t, structs.BerryFlavor{}, rFlavor, "Expected to have type 'BerryFlavor' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
