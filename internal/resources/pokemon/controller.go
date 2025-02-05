package pokemon

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Abilities":
		return controller.pokemon.GetAbilityList()
	case "Characteristics":
		return controller.pokemon.GetCharacteristicList()
	case "Egg Groups":
		return controller.pokemon.GetEggGroupList()
	case "Genders":
		return controller.pokemon.GetGenderList()
	case "Growth Rates":
		return controller.pokemon.GetGrowthRateList()
	case "Natures":
		return controller.pokemon.GetNatureList()
	case "Pokeathlon Stats":
		return controller.pokemon.GetPokeathlonStatList()
	case "Pokemon":
		return controller.pokemon.GetPokemonList()
	case "Pokemon Colors":
		return controller.pokemon.GetPokemonColorList()
	case "Pokemon Forms":
		return controller.pokemon.GetPokemonFormList()
	case "Pokemon Habitats":
		return controller.pokemon.GetPokemonHabitatList()
	case "Pokemon Shapes":
		return controller.pokemon.GetPokemonShapeList()
	case "Pokemon Species":
		return controller.pokemon.GetPokemonSpeciesList()
	case "Stats":
		return controller.pokemon.GetStatList()
	case "Types":
		return controller.pokemon.GetTypeList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Abilities":
		return controller.pokemon.GetAbility(search)
	case "Characteristics":
		return controller.pokemon.GetCharacteristic(search)
	case "Egg Groups":
		return controller.pokemon.GetEggGroup(search)
	case "Genders":
		return controller.pokemon.GetGender(search)
	case "Growth Rates":
		return controller.pokemon.GetGrowthRate(search)
	case "Natures":
		return controller.pokemon.GetNature(search)
	case "Pokeathlon Stats":
		return controller.pokemon.GetPokeathlonStat(search)
	case "Pokemon":
		return controller.pokemon.GetPokemon(search)
	case "Pokemon Colors":
		return controller.pokemon.GetPokemonColor(search)
	case "Pokemon Forms":
		return controller.pokemon.GetPokemonForm(search)
	case "Pokemon Habitats":
		return controller.pokemon.GetPokemonHabitat(search)
	case "Pokemon Shapes":
		return controller.pokemon.GetPokemonShape(search)
	case "Pokemon Species":
		return controller.pokemon.GetPokemonSpecies(search)
	case "Stats":
		return controller.pokemon.GetStat(search)
	case "Types":
		return controller.pokemon.GetType(search)
	}

	return nil, nil
}
