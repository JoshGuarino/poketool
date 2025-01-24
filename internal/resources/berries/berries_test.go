package berries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var b = Berries{}

func TestGetBerry(t *testing.T) {
	rById, _ := b.GetBerry("1")
	rByName, _ := b.GetBerry("cheri")
	rFail, _ := b.GetBerry("test")
	assert.Equal(t, "cheri", rById.Name, "Expected to receive 'cheri'.")
	assert.Equal(t, "cheri", rByName.Name, "Expected to receive 'cheri'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryList(t *testing.T) {
	r := b.GetBerryList()
	assert.Equal(t, "cheri", r.Results[0].Name, "Expected to have a list of 'berry' resource.")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := b.GetBerryFirmness("1")
	rByName, _ := b.GetBerryFirmness("very-soft")
	rFail, _ := b.GetBerryFirmness("test")
	assert.Equal(t, "very-soft", rById.Name, "Expected to receive 'very-soft'.")
	assert.Equal(t, "very-soft", rByName.Name, "Expected to receive 'very-soft'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	r := b.GetBerryFirmnessList()
	assert.Equal(t, "very-soft", r.Results[0].Name, "Expected to have a list of 'berry firmness' resource.")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := b.GetBerryFlavor("1")
	rByName, _ := b.GetBerryFlavor("spicy")
	rFail, _ := b.GetBerryFlavor("test")
	assert.Equal(t, "spicy", rById.Name, "Expected to receive 'spicy'.")
	assert.Equal(t, "spicy", rByName.Name, "Expected to receive 'spicy'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryFlavorList(t *testing.T) {
	r := b.GetBerryFlavorList()
	assert.Equal(t, "spicy", r.Results[0].Name, "Expected to have a list of 'berry flavor' resource.")
}
