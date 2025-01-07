package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAbility(t *testing.T) {
	rById, _ := GetAbility("1")
	rByName, _ := GetAbility("stench")
	rFail, _ := GetAbility("test")
	assert.Equal(t, "stench", rById.Name, "Expected to receive 'stench'.")
	assert.Equal(t, "stench", rByName.Name, "Expected to receive 'stench'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetAbilityList(t *testing.T) {
	r := GetAbilityList()
	assert.Equal(t, "stench", r.Results[0].Name, "Expected to have a list of 'ability' resource.")
}

func TestGetCharacteristic(t *testing.T) {
	rById, _ := GetCharacteristic("1")
	rFail, _ := GetCharacteristic("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expect to receive an empty result.")
}

func TestGetCharacteristicList(t *testing.T) {
	r := GetCharacteristicList()
	assert.Equal(t, "https://pokeapi.co/api/v2/characteristic/1/", r.Results[0].URL, "Expected to have a list of 'characteristic' resource.")
}

func TestGetEggGroup(t *testing.T) {
	rById, _ := GetEggGroup("1")
	rByName, _ := GetEggGroup("monster")
	rFail, _ := GetEggGroup("test")
	assert.Equal(t, "monster", rById.Name, "Expected to receive 'monster'.")
	assert.Equal(t, "monster", rByName.Name, "Expected to receive 'monster'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEggGroupList(t *testing.T) {
	r := GetEggGroupList()
	assert.Equal(t, "monster", r.Results[0].Name, "Expected to have a list of 'egg group' resource.")
}

func TestGetGender(t *testing.T) {
	rById, _ := GetGender("1")
	rByName, _ := GetGender("female")
	rFail, _ := GetGender("test")
	assert.Equal(t, "female", rById.Name, "Expected to receive 'female'.")
	assert.Equal(t, "female", rByName.Name, "Expected to receive 'female'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetGenderList(t *testing.T) {
	r := GetGenderList()
	assert.Equal(t, "female", r.Results[0].Name, "Expected to have a list of 'gender' resource.")
}

func TestGetGrowthRate(t *testing.T) {
	rById, _ := GetGrowthRate("1")
	rByName, _ := GetGrowthRate("slow")
	rFail, _ := GetGrowthRate("test")
	assert.Equal(t, "slow", rById.Name, "Expected to receive 'slow'.")
	assert.Equal(t, "slow", rByName.Name, "Expected to receive 'slow.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetGrowthRateList(t *testing.T) {
	r := GetGrowthRateList()
	assert.Equal(t, "slow", r.Results[0].Name, "Expected to have a list of 'growth rate' resource.")
}

func TestGetNature(t *testing.T) {
	rById, _ := GetNature("1")
	rByName, _ := GetNature("hardy")
	rFail, _ := GetNature("test")
	assert.Equal(t, "hardy", rById.Name, "Expected to receive 'hardy'.")
	assert.Equal(t, "hardy", rByName.Name, "Expected to receive 'hardy'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetNatureList(t *testing.T) {
	r := GetNatureList()
	assert.Equal(t, "hardy", r.Results[0].Name, "Expected to have a list of 'nature' resource.")
}

func TestGetPokeathlonStat(t *testing.T) {
	rById, _ := GetPokeathlonStat("1")
	rByName, _ := GetPokeathlonStat("speed")
	rFail, _ := GetPokeathlonStat("test")
	assert.Equal(t, "speed", rById.Name, "Expected to receive 'speed'.")
	assert.Equal(t, "speed", rByName.Name, "Expected to receive 'speed'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokeathlonStatList(t *testing.T) {
	r := GetPokeathlonStatList()
	assert.Equal(t, "speed", r.Results[0].Name, "Expected to have a list of 'pokeathlon stat' resource.")
}

func TestGetPokemon(t *testing.T) {
	rById, _ := GetPokemon("1")
	rByName, _ := GetPokemon("bulbasaur")
	rFail, _ := GetPokemon("test")
	assert.Equal(t, "bulbasaur", rById.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "bulbasaur", rByName.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonList(t *testing.T) {
	r := GetPokemonList()
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to have a list of 'pokemon' resource.")
}

func TestGetPokemonLocationAreas(t *testing.T) {
	rById := GetPokemonLocationAreas("1")
	rByName := GetPokemonLocationAreas("bulbasaur")
	assert.Equal(t, "cerulean-city-area", rById[0].LocationArea.Name, "Expected to have a list of 'pokemon location area' resource.")
	assert.Equal(t, "cerulean-city-area", rByName[0].LocationArea.Name, "Expected to have a list of 'pokemon location area' resource.")
}

func TestGetPokemonColor(t *testing.T) {
	rById, _ := GetPokemonColor("1")
	rByName, _ := GetPokemonColor("black")
	rFail, _ := GetPokemonColor("test")
	assert.Equal(t, "black", rById.Name, "Expected to receive 'black'.")
	assert.Equal(t, "black", rByName.Name, "Expected to receive 'black'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonColorList(t *testing.T) {
	r := GetPokemonColorList()
	assert.Equal(t, "black", r.Results[0].Name, "Expected to have a list of 'pokemon color' resource.")
}

func TestGetPokemonForm(t *testing.T) {
	rById, _ := GetPokemonForm("1")
	rByName, _ := GetPokemonForm("bulbasaur")
	rFail, _ := GetPokemonForm("test")
	assert.Equal(t, "bulbasaur", rById.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "bulbasaur", rByName.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonFormList(t *testing.T) {
	r := GetPokemonFormList()
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to have a list of 'pokemon form' resource.")
}

func TestGetPokemonHabitat(t *testing.T) {
	rById, _ := GetPokemonHabitat("1")
	rByName, _ := GetPokemonHabitat("cave")
	rFail, _ := GetPokemonHabitat("test")
	assert.Equal(t, "cave", rById.Name, "Expected to receive 'cave'.")
	assert.Equal(t, "cave", rByName.Name, "Expected to receive 'cave'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonHabitatList(t *testing.T) {
	r := GetPokemonHabitatList()
	assert.Equal(t, "cave", r.Results[0].Name, "Expected to have a list of 'pokemon habitat' resource.")
}

func TestGetPokemonShape(t *testing.T) {
	rById, _ := GetPokemonShape("1")
	rByName, _ := GetPokemonShape("ball")
	rFail, _ := GetPokemonShape("test")
	assert.Equal(t, "ball", rById.Name, "Expected to receive 'ball'.")
	assert.Equal(t, "ball", rByName.Name, "Expected to receive 'ball'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonShapeList(t *testing.T) {
	r := GetPokemonShapeList()
	assert.Equal(t, "ball", r.Results[0].Name, "Expected to have a list of 'pokemon shape' resource.")
}

func TestGetPokemonSpecies(t *testing.T) {
	rById, _ := GetPokemonSpecies("1")
	rByName, _ := GetPokemonSpecies("bulbasaur")
	rFail, _ := GetPokemonSpecies("test")
	assert.Equal(t, "bulbasaur", rById.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "bulbasaur", rByName.Name, "Expected to receive 'bulbasaur'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetPokemonSpeciesList(t *testing.T) {
	r := GetPokemonSpeciesList()
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to have a list of 'pokemon species' resource.")
}

func TestGetStat(t *testing.T) {
	rById, _ := GetStat("1")
	rByName, _ := GetStat("hp")
	rFail, _ := GetStat("test")
	assert.Equal(t, "hp", rById.Name, "Expected to receive 'hp'.")
	assert.Equal(t, "hp", rByName.Name, "Expected to receive 'hp'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetStatList(t *testing.T) {
	r := GetStatList()
	assert.Equal(t, "hp", r.Results[0].Name, "Expected to have a list of 'stat' resource.")
}

func TestGetType(t *testing.T) {
	rById, _ := GetType("1")
	rByName, _ := GetType("normal")
	rFail, _ := GetType("test")
	assert.Equal(t, "normal", rById.Name, "Expected to receive 'normal'.")
	assert.Equal(t, "normal", rByName.Name, "Expected to receive 'normal'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetTypeList(t *testing.T) {
	r := GetTypeList()
	assert.Equal(t, "normal", r.Results[0].Name, "Expected to have a list of 'type' resource.")
}
