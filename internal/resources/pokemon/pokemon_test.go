package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pokemon IPokemon = Pokemon{}

func TestGetAbility(t *testing.T) {
	rById, _ := pokemon.GetAbility("1")
	rByName, _ := pokemon.GetAbility("stench")
	rFail, _ := pokemon.GetAbility("test")
	assert.Equal(t, "stench", rById.Name, "Expected to receive 'stench'.")
	assert.Equal(t, "stench", rByName.Name, "Expected to receive 'stench'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetAbilityList(t *testing.T) {
	r := pokemon.GetAbilityList()
	assert.Equal(t, "stench", r.Results[0].Name, "Expected to have a list of 'ability' resource.")
}

func TestGetCharacteristic(t *testing.T) {
	rById, _ := pokemon.GetCharacteristic("1")
	rFail, _ := pokemon.GetCharacteristic("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expect to receive an empty result.")
}

func TestGetCharacteristicList(t *testing.T) {
	r := pokemon.GetCharacteristicList()
	assert.Equal(t, "https://pokeapi.co/api/v2/characteristic/1/", r.Results[0].URL, "Expected to have a list of 'characteristic' resource.")
}

func TestGetEggGroup(t *testing.T) {
	rById, _ := pokemon.GetEggGroup("1")
	rByName, _ := pokemon.GetEggGroup("monster")
	rFail, _ := pokemon.GetEggGroup("test")
	assert.Equal(t, "monster", rById.Name, "Expected to receive 'monster'.")
	assert.Equal(t, "monster", rByName.Name, "Expected to receive 'monster'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEggGroupList(t *testing.T) {
	r := pokemon.GetEggGroupList()
	assert.Equal(t, "monster", r.Results[0].Name, "Expected to have a list of 'egg group' resource.")
}

func TestGetGender(t *testing.T) {
	rById, _ := pokemon.GetGender("1")
	rByName, _ := pokemon.GetGender("female")
	rFail, _ := pokemon.GetGender("test")
	assert.Equal(t, "female", rById.Name, "Expected to receive 'female'.")
	assert.Equal(t, "female", rByName.Name, "Expected to receive 'female'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetGenderList(t *testing.T) {
	r := pokemon.GetGenderList()
	assert.Equal(t, "female", r.Results[0].Name, "Expected to have a list of 'gender' resource.")
}

func TestGetGrowthRate(t *testing.T) {
	rById, _ := pokemon.GetGrowthRate("1")
	rByName, _ := pokemon.GetGrowthRate("slow")
	rFail, _ := pokemon.GetGrowthRate("test")
	assert.Equal(t, "slow", rById.Name, "Expected to receive 'slow'.")
	assert.Equal(t, "slow", rByName.Name, "Expected to receive 'slow.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetGrowthRateList(t *testing.T) {
	r := pokemon.GetGrowthRateList()
	assert.Equal(t, "slow", r.Results[0].Name, "Expected to have a list of 'growth rate' resource.")
}

func TestGetNature(t *testing.T) {
	rById, _ := pokemon.GetNature("1")
	rByName, _ := pokemon.GetNature("hardy")
	rFail, _ := pokemon.GetNature("test")
	assert.Equal(t, "hardy", rById.Name, "Expected to receive 'hardy'.")
	assert.Equal(t, "hardy", rByName.Name, "Expected to receive 'hardy'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetNatureList(t *testing.T) {
	r := pokemon.GetNatureList()
	assert.Equal(t, "hardy", r.Results[0].Name, "Expected to have a list of 'nature' resource.")
}

func TestGetPokeathlonStat(t *testing.T) {
	rById, _ := pokemon.GetPokeathlonStat("1")
	rByName, _ := pokemon.GetPokeathlonStat("speed")
	rFail, _ := pokemon.GetPokeathlonStat("test")
	assert.Equal(t, "speed", rById.Name, "Expected to receive 'speed'.")
	assert.Equal(t, "speed", rByName.Name, "Expected to receive 'speed'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokeathlonStatList(t *testing.T) {
	r := pokemon.GetPokeathlonStatList()
	assert.Equal(t, "speed", r.Results[0].Name, "Expected to have a list of 'pokeathlon stat' resource.")
}

func TestGetPokemon(t *testing.T) {
	rById, _ := pokemon.GetPokemon("1")
	rByName, _ := pokemon.GetPokemon("bulbasaur")
	rFail, _ := pokemon.GetPokemon("test")
	assert.Equal(t, "bulbasaur", rById.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "bulbasaur", rByName.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonList(t *testing.T) {
	r := pokemon.GetPokemonList()
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to have a list of 'pokemon' resource.")
}

func TestGetPokemonLocationAreas(t *testing.T) {
	rById := pokemon.GetPokemonLocationAreas("1")
	rByName := pokemon.GetPokemonLocationAreas("bulbasaur")
	assert.Equal(t, "cerulean-city-area", rById[0].LocationArea.Name, "Expected to have a list of 'pokemon location area' resource.")
	assert.Equal(t, "cerulean-city-area", rByName[0].LocationArea.Name, "Expected to have a list of 'pokemon location area' resource.")
}

func TestGetPokemonColor(t *testing.T) {
	rById, _ := pokemon.GetPokemonColor("1")
	rByName, _ := pokemon.GetPokemonColor("black")
	rFail, _ := pokemon.GetPokemonColor("test")
	assert.Equal(t, "black", rById.Name, "Expected to receive 'black'.")
	assert.Equal(t, "black", rByName.Name, "Expected to receive 'black'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonColorList(t *testing.T) {
	r := pokemon.GetPokemonColorList()
	assert.Equal(t, "black", r.Results[0].Name, "Expected to have a list of 'pokemon color' resource.")
}

func TestGetPokemonForm(t *testing.T) {
	rById, _ := pokemon.GetPokemonForm("1")
	rByName, _ := pokemon.GetPokemonForm("bulbasaur")
	rFail, _ := pokemon.GetPokemonForm("test")
	assert.Equal(t, "bulbasaur", rById.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "bulbasaur", rByName.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonFormList(t *testing.T) {
	r := pokemon.GetPokemonFormList()
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to have a list of 'pokemon form' resource.")
}

func TestGetPokemonHabitat(t *testing.T) {
	rById, _ := pokemon.GetPokemonHabitat("1")
	rByName, _ := pokemon.GetPokemonHabitat("cave")
	rFail, _ := pokemon.GetPokemonHabitat("test")
	assert.Equal(t, "cave", rById.Name, "Expected to receive 'cave'.")
	assert.Equal(t, "cave", rByName.Name, "Expected to receive 'cave'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonHabitatList(t *testing.T) {
	r := pokemon.GetPokemonHabitatList()
	assert.Equal(t, "cave", r.Results[0].Name, "Expected to have a list of 'pokemon habitat' resource.")
}

func TestGetPokemonShape(t *testing.T) {
	rById, _ := pokemon.GetPokemonShape("1")
	rByName, _ := pokemon.GetPokemonShape("ball")
	rFail, _ := pokemon.GetPokemonShape("test")
	assert.Equal(t, "ball", rById.Name, "Expected to receive 'ball'.")
	assert.Equal(t, "ball", rByName.Name, "Expected to receive 'ball'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonShapeList(t *testing.T) {
	r := pokemon.GetPokemonShapeList()
	assert.Equal(t, "ball", r.Results[0].Name, "Expected to have a list of 'pokemon shape' resource.")
}

func TestGetPokemonSpecies(t *testing.T) {
	rById, _ := pokemon.GetPokemonSpecies("1")
	rByName, _ := pokemon.GetPokemonSpecies("bulbasaur")
	rFail, _ := pokemon.GetPokemonSpecies("test")
	assert.Equal(t, "bulbasaur", rById.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "bulbasaur", rByName.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonSpeciesList(t *testing.T) {
	r := pokemon.GetPokemonSpeciesList()
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to have a list of 'pokemon species' resource.")
}

func TestGetStat(t *testing.T) {
	rById, _ := pokemon.GetStat("1")
	rByName, _ := pokemon.GetStat("hp")
	rFail, _ := pokemon.GetStat("test")
	assert.Equal(t, "hp", rById.Name, "Expected to receive 'hp'.")
	assert.Equal(t, "hp", rByName.Name, "Expected to receive 'hp'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetStatList(t *testing.T) {
	r := pokemon.GetStatList()
	assert.Equal(t, "hp", r.Results[0].Name, "Expected to have a list of 'stat' resource.")
}

func TestGetType(t *testing.T) {
	rById, _ := pokemon.GetType("1")
	rByName, _ := pokemon.GetType("normal")
	rFail, _ := pokemon.GetType("test")
	assert.Equal(t, "normal", rById.Name, "Expected to receive 'normal'.")
	assert.Equal(t, "normal", rByName.Name, "Expected to receive 'normal'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetTypeList(t *testing.T) {
	r := pokemon.GetTypeList()
	assert.Equal(t, "normal", r.Results[0].Name, "Expected to have a list of 'type' resource.")
}
