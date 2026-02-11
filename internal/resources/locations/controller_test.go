package locations

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rLocation, _ := controller.GetList("Location", 20, 0)
	rArea, _ := controller.GetList("Location Area", 20, 0)
	rPalPark, _ := controller.GetList("Pal Park Area", 20, 0)
	rRegion, _ := controller.GetList("Region", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.NamedResourceList{}, rLocation, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rLocation, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rArea, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rArea, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rPalPark, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rPalPark, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rRegion, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rRegion, "Expected to not have an empty array.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rLocation, _ := controller.GetSpecific("Location", "1")
	rArea, _ := controller.GetSpecific("Location Area", "1")
	rPalPark, _ := controller.GetSpecific("Pal Park Area", "1")
	rRegion, _ := controller.GetSpecific("Region", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.Location{}, rLocation, "Expected to have type 'Location' struct.")
	assert.IsType(t, &models.LocationArea{}, rArea, "Expected to have type 'LocationArea' struct.")
	assert.IsType(t, &models.PalParkArea{}, rPalPark, "Expected to have type 'PalParkArea' struct.")
	assert.IsType(t, &models.Region{}, rRegion, "Expected to have type 'Region' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
