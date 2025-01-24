package berries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var berries = Berries{}

func TestGetBerry(t *testing.T) {
	rById, _ := berries.GetBerry("1")
	rByName, _ := berries.GetBerry("cheri")
	rFail, _ := berries.GetBerry("test")
	assert.Equal(t, "cheri", rById.Name, "Expected to receive 'cheri'.")
	assert.Equal(t, "cheri", rByName.Name, "Expected to receive 'cheri'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryList(t *testing.T) {
	r := berries.GetBerryList()
	assert.Equal(t, "cheri", r.Results[0].Name, "Expected to have a list of 'berry' resource.")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := berries.GetBerryFirmness("1")
	rByName, _ := berries.GetBerryFirmness("very-soft")
	rFail, _ := berries.GetBerryFirmness("test")
	assert.Equal(t, "very-soft", rById.Name, "Expected to receive 'very-soft'.")
	assert.Equal(t, "very-soft", rByName.Name, "Expected to receive 'very-soft'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	r := berries.GetBerryFirmnessList()
	assert.Equal(t, "very-soft", r.Results[0].Name, "Expected to have a list of 'berry firmness' resource.")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := berries.GetBerryFlavor("1")
	rByName, _ := berries.GetBerryFlavor("spicy")
	rFail, _ := berries.GetBerryFlavor("test")
	assert.Equal(t, "spicy", rById.Name, "Expected to receive 'spicy'.")
	assert.Equal(t, "spicy", rByName.Name, "Expected to receive 'spicy'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryFlavorList(t *testing.T) {
	r := berries.GetBerryFlavorList()
	assert.Equal(t, "spicy", r.Results[0].Name, "Expected to have a list of 'berry flavor' resource.")
}
