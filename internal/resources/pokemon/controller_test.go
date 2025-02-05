package pokemon

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{pokemon: Pokemon{}}

func TestGetList(t *testing.T) {
	rAbilities := controller.GetList("Abilities")
	rCharacteristics := controller.GetList("Characteristics")
	rEgg := controller.GetList("Egg Groups")
	rGender := controller.GetList("Genders")
	rGrowth := controller.GetList("Growth Rates")
	rNatures := controller.GetList("Natures")
	rPokeathlon := controller.GetList("Pokeathlon Stats")
	rPokemon := controller.GetList("Pokemon")
	rColors := controller.GetList("Pokemon Colors")
	rForms := controller.GetList("Pokemon Forms")
	rHabitats := controller.GetList("Pokemon Habitats")
	rShapes := controller.GetList("Pokemon Shapes")
	rSpecies := controller.GetList("Pokemon Species")
	rStats := controller.GetList("Stats")
	rTypes := controller.GetList("Types")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rAbilities.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rCharacteristics.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rEgg.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rGender.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rGrowth.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rNatures.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rPokeathlon.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rPokemon.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rColors.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rForms.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rHabitats.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rShapes.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rSpecies.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rStats.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rTypes.Results, "Expected to have array of type 'Result' struct.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rAbilities, _ := controller.GetSpecific("Abilities", "1")
	rCharacteristics, _ := controller.GetSpecific("Characteristics", "1")
	rEgg, _ := controller.GetSpecific("Egg Groups", "1")
	rGender, _ := controller.GetSpecific("Genders", "1")
	rGrowth, _ := controller.GetSpecific("Growth Rates", "1")
	rNatures, _ := controller.GetSpecific("Natures", "1")
	rPokeathlon, _ := controller.GetSpecific("Pokeathlon Stats", "1")
	rPokemon, _ := controller.GetSpecific("Pokemon", "1")
	rColors, _ := controller.GetSpecific("Pokemon Colors", "1")
	rForms, _ := controller.GetSpecific("Pokemon Forms", "1")
	rHabitats, _ := controller.GetSpecific("Pokemon Habitats", "1")
	rShapes, _ := controller.GetSpecific("Pokemon Shapes", "1")
	rSpecies, _ := controller.GetSpecific("Pokemon Species", "1")
	rStats, _ := controller.GetSpecific("Stats", "1")
	rTypes, _ := controller.GetSpecific("Types", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Ability{}, rAbilities, "Expected to have type 'Ability' struct.")
	assert.IsType(t, structs.Characteristic{}, rCharacteristics, "Expected to have type 'Characteristic' struct.")
	assert.IsType(t, structs.EggGroup{}, rEgg, "Expected to have type 'EggGroup' struct.")
	assert.IsType(t, structs.Gender{}, rGender, "Expected to have type 'Gender' struct.")
	assert.IsType(t, structs.GrowthRate{}, rGrowth, "Expected to have type 'GrowthRate' struct.")
	assert.IsType(t, structs.Nature{}, rNatures, "Expected to have type 'Nature' struct.")
	assert.IsType(t, structs.PokeathlonStat{}, rPokeathlon, "Expected to have type 'PokeathlonStat' struct.")
	assert.IsType(t, structs.Pokemon{}, rPokemon, "Expected to have type 'Pokemon' struct.")
	assert.IsType(t, structs.PokemonColor{}, rColors, "Expected to have type 'PokemonColor' struct.")
	assert.IsType(t, structs.PokemonForm{}, rForms, "Expected to have type 'PokemonForm' struct.")
	assert.IsType(t, structs.PokemonHabitat{}, rHabitats, "Expected to have type 'PokemonHabitat' struct.")
	assert.IsType(t, structs.PokemonShape{}, rShapes, "Expected to have type 'PokemonShape' struct.")
	assert.IsType(t, structs.PokemonSpecies{}, rSpecies, "Expected to have type 'PokemonSpecies' struct.")
	assert.IsType(t, structs.Stat{}, rStats, "Expected to have type 'Stat' struct.")
	assert.IsType(t, structs.Type{}, rTypes, "Expected to have type 'Type' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
