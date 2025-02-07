package locations

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{locations: Locations{}}

func TestGetList(t *testing.T) {
	rLocation := controller.GetList("Location")
	rArea := controller.GetList("Location Area")
	rPalPark := controller.GetList("Pal Park Area")
	rRegion := controller.GetList("Region")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rLocation.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rLocation.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rArea.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rArea.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rPalPark.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rPalPark.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rRegion.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rRegion.Results, "Expected to not have an empty array.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rLocation, _ := controller.GetSpecific("Location", "1")
	rArea, _ := controller.GetSpecific("Location Area", "1")
	rPalPark, _ := controller.GetSpecific("Pal Park Area", "1")
	rRegion, _ := controller.GetSpecific("Region", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Location{}, rLocation, "Expected to have type 'Location' struct.")
	assert.IsType(t, structs.LocationArea{}, rArea, "Expected to have type 'LocationArea' struct.")
	assert.IsType(t, structs.PalParkArea{}, rPalPark, "Expected to have type 'PalParkArea' struct.")
	assert.IsType(t, structs.Region{}, rRegion, "Expected to have type 'Region' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
