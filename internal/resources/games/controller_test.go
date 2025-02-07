package games

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{games: Games{}}

func TestGetList(t *testing.T) {
	rGen := controller.GetList("Generation")
	rPokedex := controller.GetList("Pokedex")
	rVersion := controller.GetList("Version")
	rGroup := controller.GetList("Version Group")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rGen.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rGen.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rPokedex.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rPokedex.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rVersion.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rVersion.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rGroup.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rGroup.Results, "Expected to not have an empty array.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rGen, _ := controller.GetSpecific("Generation", "1")
	rPokedex, _ := controller.GetSpecific("Pokedex", "1")
	rVersion, _ := controller.GetSpecific("Version", "1")
	rGroup, _ := controller.GetSpecific("Version Group", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Generation{}, rGen, "Expected to have type 'Generation' struct.")
	assert.IsType(t, structs.Pokedex{}, rPokedex, "Expected to have type 'Pokedex' struct.")
	assert.IsType(t, structs.Version{}, rVersion, "Expected to have type 'Version' struct.")
	assert.IsType(t, structs.VersionGroup{}, rGroup, "Expected to have type 'VersionGroup' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
