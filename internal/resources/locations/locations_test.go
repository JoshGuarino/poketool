package locations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var locations ILocations = Locations{}

func TestGetLocation(t *testing.T) {
	rById, _ := locations.GetLocation("1")
	rByName, _ := locations.GetLocation("canalave-city")
	rFail, _ := locations.GetLocation("test")
	assert.Equal(t, "canalave-city", rById.Name, "Expected to receive 'canalave-city'.")
	assert.Equal(t, "canalave-city", rByName.Name, "Expected to receive 'canalave-city'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetLocationList(t *testing.T) {
	r := locations.GetLocationList()
	assert.Equal(t, "canalave-city", r.Results[0].Name, "Expected to have a list of 'location' resource.")
}

func TestGetLocationArea(t *testing.T) {
	rById, _ := locations.GetLocationArea("1")
	rByName, _ := locations.GetLocationArea("canalave-city-area")
	rFail, _ := locations.GetLocationArea("test")
	assert.Equal(t, "canalave-city-area", rById.Name, "Expected to receive 'canalave-city-area'.")
	assert.Equal(t, "canalave-city-area", rByName.Name, "Expected to receive 'canalave-city-area'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetLocationAreaList(t *testing.T) {
	r := locations.GetLocationAreaList()
	assert.Equal(t, "canalave-city-area", r.Results[0].Name, "Expected to have a list of 'location-area' resource.")
}

func TestGetPalParkArea(t *testing.T) {
	rById, _ := locations.GetPalParkArea("1")
	rByName, _ := locations.GetPalParkArea("forest")
	rFail, _ := locations.GetPalParkArea("test")
	assert.Equal(t, "forest", rById.Name, "Expected to receive 'forest'.")
	assert.Equal(t, "forest", rByName.Name, "Expected to receive 'forest'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPalParkAreaList(t *testing.T) {
	r := locations.GetPalParkAreaList()
	assert.Equal(t, "forest", r.Results[0].Name, "Expected to have a list of 'pal park area' resource.")
}

func TestGetRegion(t *testing.T) {
	rById, _ := locations.GetRegion("1")
	rByName, _ := locations.GetRegion("kanto")
	rFail, _ := locations.GetRegion("test")
	assert.Equal(t, "kanto", rById.Name, "Expected to receive 'kanto'.")
	assert.Equal(t, "kanto", rByName.Name, "Expected to receive 'kanto'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetRegionList(t *testing.T) {
	r := locations.GetRegionList()
	assert.Equal(t, "kanto", r.Results[0].Name, "Expected to have a list of 'region' resource.")
}
