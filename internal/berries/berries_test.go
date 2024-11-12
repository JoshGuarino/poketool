package berries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBerry(t *testing.T) {
	rById, _ := GetBerry("1")
	rByName, _ := GetBerry("cheri")
	rFail, _ := GetBerry("test")
	assert.Equal(t, "cheri", rById.Name, "Expected to receive Cheri.")
	assert.Equal(t, "cheri", rByName.Name, "Expected to receive Cheri.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryList(t *testing.T) {
	result := GetBerryList()
	assert.Equal(t, "cheri", result.Results[0].Name, "Expected to have a list berries resource.")
}

func TestGetBerryFirmness(t *testing.T) {
	rById, _ := GetBerryFirmness("1")
	rByName, _ := GetBerryFirmness("very-soft")
	rFail, _ := GetBerryFirmness("test")
	assert.Equal(t, "very-soft", rById.Name, "Expected to receive Very Soft.")
	assert.Equal(t, "very-soft", rByName.Name, "Expected to receive Very Soft.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	result := GetBerryFirmnessList()
	assert.Equal(t, "very-soft", result.Results[0].Name, "Expected to have a list of berry firmness resource.")
}

func TestGetBerryFlavor(t *testing.T) {
	rById, _ := GetBerryFlavor("1")
	rByName, _ := GetBerryFlavor("spicy")
	rFail, _ := GetBerryFlavor("test")
	assert.Equal(t, "spicy", rById.Name, "Expected to receive Spicy.")
	assert.Equal(t, "spicy", rByName.Name, "Expected to receive Spicy.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetBerryFlavorList(t *testing.T) {
	result := GetBerryFlavorList()
	assert.Equal(t, "spicy", result.Results[0].Name, "Expected to have a list of berry flavor resource.")
}
