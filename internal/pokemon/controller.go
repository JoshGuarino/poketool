package pokemon

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetList(result string) structs.Resource {
	switch result {
	case "Abilities":
		return GetAbilityList()
	case "Characteristics":
		return GetCharacteristicList()
	case "Egg Groups":
		return GetEggGroupList()
	case "Genders":
		return GetGenderList()
	case "Growth Rates":
		return GetGrowthRateList()
	case "Natures":
		return GetNatureList()
	case "Pokeathlon Stats":
		return GetPokeathlonStatList()
	case "Pokemon":
		return GetPokemonList()
	case "Pokemon Colors":
		return GetPokemonColorList()
	case "Pokemon Forms":
		return GetPokemonFormList()
	case "pokemon Habitats":
		return GetPokemonHabitatList()
	case "Pokemon Shapes":
		return GetPokemonShapeList()
	case "Pokemon Species":
		return GetPokemonSpeciesList()
	case "Stats":
		return GetStatList()
	case "Types":
		return GetTypeList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) interface{} {
	switch result {
	case "Abilities":
		return GetAbility(search)
	case "Characteristics":
		return GetCharacteristic(search)
	case "Egg Group":
		return GetEggGroup(search)
	case "Genders":
		return GetGender(search)
	case "Growth Rates":
		return GetGrowthRate(search)
	case "Natures":
		return GetNature(search)
	case "Pokeathlon Stats":
		return GetPokeathlonStat(search)
	case "Pokemon":
		return GetPokemon(search)
	case "Pokemon Colors":
		return GetPokemonColor(search)
	case "Pokemon Forms":
		return GetPokemonForm(search)
	case "pokemon Habitats":
		return GetPokemonHabitat(search)
	case "Pokemon Shapes":
		return GetPokemonShape(search)
	case "Pokemon Species":
		return GetPokemonSpecies(search)
	case "Stats":
		return GetStat(search)
	case "Types":
		return GetType(search)
	}

	return ""
}
