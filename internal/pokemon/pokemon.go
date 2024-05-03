package pokemon

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetAbility(nameOrId string) structs.Ability {
	ability, err := pokeapi.Ability(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ability
}

func GetAbilityList() structs.Resource {
	abilityList := internal.GetResourceList(abilityEndpoint)
	return abilityList
}

func GetCharacteristic(id string) structs.Characteristic {
	characteristic, err := pokeapi.Characteristic(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return characteristic
}

func GetCharacteristicList() structs.Resource {
	characteristicList := internal.GetResourceList(characteristicEndpoint)
	return characteristicList
}

func GetEggGroup(nameOrId string) structs.EggGroup {
	eggGroup, err := pokeapi.EggGroup(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return eggGroup
}

func GetEggGroupList() structs.Resource {
	eggGroupList := internal.GetResourceList(eggGroupEndpoint)
	return eggGroupList
}

func GetGender(nameOrId string) structs.Gender {
	gender, err := pokeapi.Gender(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return gender
}

func GetGenderList() structs.Resource {
	genderList := internal.GetResourceList(genderEndpoint)
	return genderList
}

func GetGrowthRate(nameOrId string) structs.GrowthRate {
	growthRate, err := pokeapi.GrowthRate(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return growthRate
}

func GetGrowthRateList() structs.Resource {
	growthRateList := internal.GetResourceList(growthRateEndpoint)
	return growthRateList
}

func GetNature(nameOrId string) structs.Nature {
	nature, err := pokeapi.Nature(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nature
}

func GetNatureList() structs.Resource {
	natureList := internal.GetResourceList(natureEndpoint)
	return natureList
}

func GetPokeathlonStat(nameOrId string) structs.PokeathlonStat {
	pokeathlonStat, err := pokeapi.PokeathlonStat(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokeathlonStat
}

func GetPokeathlonStatList() structs.Resource {
	pokeathlonStatList := internal.GetResourceList(pokeathlonStatEndpoint)
	return pokeathlonStatList
}

func GetPokemon(nameOrId string) structs.Pokemon {
	pokemon, err := pokeapi.Pokemon(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemon
}

func GetPokemonList() structs.Resource {
	pokemonList := internal.GetResourceList(pokemonEndpoint)
	return pokemonList
}

func GetPokemonLocationAreas(nameOrId string) []internal.LocationAreaEncounter {
	pokemonLocationAreas := internal.GetPokemonLocationAreas(nameOrId)
	return pokemonLocationAreas
}

func GetPokemonColor(nameOrId string) structs.PokemonColor {
	pokemonColor, err := pokeapi.PokemonColor(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemonColor
}

func GetPokemonColorList() structs.Resource {
	pokemonColorList := internal.GetResourceList(pokemonColorEndpoint)
	return pokemonColorList
}

func GetPokemonForm(nameOrId string) structs.PokemonForm {
	pokemonForm, err := pokeapi.PokemonForm(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemonForm
}

func GetPokemonFormList() structs.Resource {
	pokemonFormList := internal.GetResourceList(pokemonFormEndpoint)
	return pokemonFormList
}

func GetPokemonHabitat(nameOrId string) structs.PokemonHabitat {
	pokemonHabitat, err := pokeapi.PokemonHabitat(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemonHabitat
}

func GetPokemonHabitatList() structs.Resource {
	pokemonHabitatList := internal.GetResourceList(pokemonHabitatEndpoint)
	return pokemonHabitatList
}

func GetPokemonShape(nameOrId string) structs.PokemonShape {
	pokemonShape, err := pokeapi.PokemonShape(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemonShape
}

func GetPokemonShapeList() structs.Resource {
	pokemonShapeList := internal.GetResourceList(pokemonShapeEndpoint)
	return pokemonShapeList
}

func GetPokemonSpecies(nameOrId string) structs.PokemonSpecies {
	pokemonSpecies, err := pokeapi.PokemonSpecies(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemonSpecies
}

func GetPokemonSpeciesList() structs.Resource {
	pokemonSpeciesList := internal.GetResourceList(pokemonSpeciesEndpoint)
	return pokemonSpeciesList
}

func GetStat(nameOrId string) structs.Stat {
	stat, err := pokeapi.Stat(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return stat
}

func GetStatList() structs.Resource {
	statList := internal.GetResourceList(statEndpoint)
	return statList
}

func GetType(nameOrId string) structs.Type {
	pokemonType, err := pokeapi.Type(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokemonType
}

func GetTypeList() structs.Resource {
	typeList := internal.GetResourceList(typeEndpoint)
	return typeList
}
