package berries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBerryId(t *testing.T) {
	result, _ := GetBerry("1")
	assert.Equal(t, "cheri", result.Name, "Expected to receive Cheri.")
}

func TestGetBerryByName(t *testing.T) {
	result, _ := GetBerry("cheri")
	assert.Equal(t, "cheri", result.Name, "Expected to receive Cheri.")
}

func TestGetBerryFail(t *testing.T) {
	result, _ := GetBerry("test")
	assert.Equal(t, "", result.Name, "Expected to receive an empty result.")
}

func TestGetBerryList(t *testing.T) {
	result := GetBerryList()
	assert.Equal(t, "cheri", result.Results[0].Name, "Expected to have a list berries resource.")
}

func TestGetBerryFirmness(t *testing.T) {
	result, _ := GetBerryFirmness("1")
	assert.Equal(t, "very-soft", result.Name, "Expected to receive Very Soft.")
}

func TestGetBerryFirmnessByName(t *testing.T) {
	result, _ := GetBerryFirmness("very-soft")
	assert.Equal(t, "very-soft", result.Name, "Expected to receive Very Soft.")
}

func TestGetBerryFirmnessFail(t *testing.T) {
	result, _ := GetBerryFirmness("test")
	assert.Equal(t, "", result.Name, "Expected to receive an empty result.")
}

func TestGetBerryFirmnessList(t *testing.T) {
	result := GetBerryFirmnessList()
	assert.Equal(t, "very-soft", result.Results[0].Name, "Expected to have a list of berry firmness resource.")
}

func TestGetBerryFlavor(t *testing.T) {
	result, _ := GetBerryFlavor("1")
	assert.Equal(t, "spicy", result.Name, "Expected to receive Spicy.")
}

func TestGetBerryFlavorByName(t *testing.T) {
	result, _ := GetBerryFlavor("spicy")
	assert.Equal(t, "spicy", result.Name, "Expected to receive Spicy.")
}

func TestGetBerryFlavorFail(t *testing.T) {
	result, _ := GetBerryFlavor("test")
	assert.Equal(t, "", result.Name, "Expected to receive an empty result.")
}

func TestGetBerryFlavorList(t *testing.T) {
	result := GetBerryFlavorList()
	assert.Equal(t, "spicy", result.Results[0].Name, "Expected to have a list of berry flavor resource.")
}
