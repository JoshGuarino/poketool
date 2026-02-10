package berries

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var controller IController = NewController()

func TestGetList(t *testing.T) {
	rBerries, _ := controller.GetList("Berries")
	rFirmnesses, _ := controller.GetList("Berry Firmnesses")
	rFlavors, _ := controller.GetList("Berry Flavors")
	_, err := controller.GetList("test")
	assert.IsType(t, &models.NamedResourceList{}, rBerries, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rBerries, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rFirmnesses, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rFirmnesses, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rFlavors, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rFlavors, "Expected to not have an empty struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rBerry, _ := controller.GetSpecific("Berries", "1")
	rFirmness, _ := controller.GetSpecific("Berry Firmnesses", "1")
	rFlavor, _ := controller.GetSpecific("Berry Flavors", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.Berry{}, rBerry, "Expected to have type 'Berry' struct.")
	assert.IsType(t, &models.BerryFirmness{}, rFirmness, "Expected to have type 'BerryFirmness' struct.")
	assert.IsType(t, &models.BerryFlavor{}, rFlavor, "Expected to have type 'BerryFlavor' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
