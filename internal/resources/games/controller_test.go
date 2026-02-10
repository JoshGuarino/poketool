package games

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rGen, _ := controller.GetList("Generation", 20, 0)
	rPokedex, _ := controller.GetList("Pokedex", 20, 0)
	rVersion, _ := controller.GetList("Version", 20, 0)
	rGroup, _ := controller.GetList("Version Group", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.NamedResourceList{}, rGen, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rGen, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rPokedex, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rPokedex, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rVersion, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rVersion, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rGroup, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rGroup, "Expected to not have an empty array.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rGen, _ := controller.GetSpecific("Generation", "1")
	rPokedex, _ := controller.GetSpecific("Pokedex", "1")
	rVersion, _ := controller.GetSpecific("Version", "1")
	rGroup, _ := controller.GetSpecific("Version Group", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.Generation{}, rGen, "Expected to have type 'Generation' struct.")
	assert.IsType(t, &models.Pokedex{}, rPokedex, "Expected to have type 'Pokedex' struct.")
	assert.IsType(t, &models.Version{}, rVersion, "Expected to have type 'Version' struct.")
	assert.IsType(t, &models.VersionGroup{}, rGroup, "Expected to have type 'VersionGroup' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
