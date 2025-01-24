package berries

import (
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var c = Controller{berries: Berries{}}

func TestGetList(t *testing.T) {
	rBerries := c.GetList("Berries")
	rFirmnesses := c.GetList("Berry Firmnesses")
	rFlavors := c.GetList("Berry Flavors")
	rFail := c.GetList("")
	assert.IsType(t, structs.Resource{}, rBerries, "Expected to have type 'Resource' struct.")
	assert.IsType(t, structs.Resource{}, rFirmnesses, "Expected to have type 'Resource' struct.")
	assert.IsType(t, structs.Resource{}, rFlavors, "Expected to have type 'Resource' struct.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rBerry, _ := c.GetSpecific("Berries", "1")
	rFirmness, _ := c.GetSpecific("Berry Firmnesses", "1")
	rFlavor, _ := c.GetSpecific("Berry Flavors", "1")
	rFail, _ := c.GetSpecific("test", "test")
	assert.IsType(t, structs.Berry{}, rBerry, "Expected to have type 'Berry' struct.")
	assert.IsType(t, structs.BerryFirmness{}, rFirmness, "Expected to have type 'BerryFirmness' struct.")
	assert.IsType(t, structs.BerryFlavor{}, rFlavor, "Expected to have type 'BerryFlavor' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
