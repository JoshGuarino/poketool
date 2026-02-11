package pokemon

import (
	"fmt"

	pokemonGroup "github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

type Controller struct {
	pokemon pokemonGroup.Pokemon
}

func NewController() Controller {
	return Controller{
		pokemon: pokemonGroup.NewPokemonGroup(),
	}
}

func (c Controller) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Abilities":
		return c.pokemon.GetAbilityList(limit, offset)
	case "Characteristics":
		return c.pokemon.GetCharacteristicList(limit, offset)
	case "Egg Groups":
		return c.pokemon.GetEggGroupList(limit, offset)
	case "Genders":
		return c.pokemon.GetGenderList(limit, offset)
	case "Growth Rates":
		return c.pokemon.GetGrowthRateList(limit, offset)
	case "Natures":
		return c.pokemon.GetNatureList(limit, offset)
	case "Pokeathlon Stats":
		return c.pokemon.GetPokeathlonStatList(limit, offset)
	case "Pokemon":
		return c.pokemon.GetPokemonList(limit, offset)
	case "Pokemon Colors":
		return c.pokemon.GetPokemonColorList(limit, offset)
	case "Pokemon Forms":
		return c.pokemon.GetPokemonFormList(limit, offset)
	case "Pokemon Habitats":
		return c.pokemon.GetPokemonHabitatList(limit, offset)
	case "Pokemon Shapes":
		return c.pokemon.GetPokemonShapeList(limit, offset)
	case "Pokemon Species":
		return c.pokemon.GetPokemonSpeciesList(limit, offset)
	case "Stats":
		return c.pokemon.GetStatList(limit, offset)
	case "Types":
		return c.pokemon.GetTypeList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c Controller) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Abilities":
		return c.pokemon.GetAbility(search)
	case "Characteristics":
		return c.pokemon.GetCharacteristic(search)
	case "Egg Groups":
		return c.pokemon.GetEggGroup(search)
	case "Genders":
		return c.pokemon.GetGender(search)
	case "Growth Rates":
		return c.pokemon.GetGrowthRate(search)
	case "Natures":
		return c.pokemon.GetNature(search)
	case "Pokeathlon Stats":
		return c.pokemon.GetPokeathlonStat(search)
	case "Pokemon":
		return c.pokemon.GetPokemon(search)
	case "Pokemon Colors":
		return c.pokemon.GetPokemonColor(search)
	case "Pokemon Forms":
		return c.pokemon.GetPokemonForm(search)
	case "Pokemon Habitats":
		return c.pokemon.GetPokemonHabitat(search)
	case "Pokemon Shapes":
		return c.pokemon.GetPokemonShape(search)
	case "Pokemon Species":
		return c.pokemon.GetPokemonSpecies(search)
	case "Stats":
		return c.pokemon.GetStat(search)
	case "Types":
		return c.pokemon.GetType(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
