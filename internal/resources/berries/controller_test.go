package berries

import (
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller = Controller{berries: Berries{}}

func TestGetList(t *testing.T) {
	rBerries := controller.GetList("Berries")
	rFirmnesses := controller.GetList("Berry Firmnesses")
	rFlavors := controller.GetList("Berry Flavors")
	rFail := controller.GetList("test")
	assert.IsType(t, structs.Resource{}, rBerries, "Expected to have type 'Resource' struct.")
	assert.IsType(t, structs.Resource{}, rFirmnesses, "Expected to have type 'Resource' struct.")
	assert.IsType(t, structs.Resource{}, rFlavors, "Expected to have type 'Resource' struct.")
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
