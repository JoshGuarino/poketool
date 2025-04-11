package berries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var berries IBerries = Berries{}

func TestGetBerry(t *testing.T) {
	rById, _ := berries.GetBerry("1")
	rByName, _ := berries.GetBerry("cheri")
	_, err := berries.GetBerry("test")
	assert.Equal(t, "cheri", rById.Name, "Expected to receive 'cheri'.")
	assert.Equal(t, "cheri", rByName.Name, "Expected to receive 'cheri'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryList(t *testing.T) {
	r, _ := berries.GetBerryList(20, 0)
	assert.Equal(t, "cheri", r.Results[0].Name, "Expected to have a list of 'berry' resource.")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := berries.GetBerryFirmness("1")
	rByName, _ := berries.GetBerryFirmness("very-soft")
	_, err := berries.GetBerryFirmness("test")
	assert.Equal(t, "very-soft", rById.Name, "Expected to receive 'very-soft'.")
	assert.Equal(t, "very-soft", rByName.Name, "Expected to receive 'very-soft'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	r, _ := berries.GetBerryFirmnessList(20, 0)
	assert.Equal(t, "very-soft", r.Results[0].Name, "Expected to have a list of 'berry firmness' resource.")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := berries.GetBerryFlavor("1")
	rByName, _ := berries.GetBerryFlavor("spicy")
	_, err := berries.GetBerryFlavor("test")
	assert.Equal(t, "spicy", rById.Name, "Expected to receive 'spicy'.")
	assert.Equal(t, "spicy", rByName.Name, "Expected to receive 'spicy'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetBerryFlavorList(t *testing.T) {
	r, _ := berries.GetBerryFlavorList(20, 0)
	assert.Equal(t, "spicy", r.Results[0].Name, "Expected to have a list of 'berry flavor' resource.")
}
