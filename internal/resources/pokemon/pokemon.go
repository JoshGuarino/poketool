package pokemon

import (
	"encoding/json"
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (pokemon Pokemon) GetAbility(nameOrId string) (structs.Ability, error) {
	ability, err := pokeapi.Ability(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "ability", nameOrId)
		return structs.Ability{}, err
	}
	return ability, nil
}

func (pokemon Pokemon) GetAbilityList() structs.Resource {
	abilityList := internal.GetResourceList(abilityEndpoint)
	return abilityList
}

func (pokemon Pokemon) GetCharacteristic(id string) (structs.Characteristic, error) {
	characteristic, err := pokeapi.Characteristic(id)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetById, "characteristic", id)
		return structs.Characteristic{}, err
	}
	return characteristic, nil
}

func (pokemon Pokemon) GetCharacteristicList() structs.Resource {
	characteristicList := internal.GetResourceList(characteristicEndpoint)
	return characteristicList
}

func (pokemon Pokemon) GetEggGroup(nameOrId string) (structs.EggGroup, error) {
	eggGroup, err := pokeapi.EggGroup(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "egg group", nameOrId)
		return structs.EggGroup{}, err
	}
	return eggGroup, nil
}

func (pokemon Pokemon) GetEggGroupList() structs.Resource {
	eggGroupList := internal.GetResourceList(eggGroupEndpoint)
	return eggGroupList
}

func (pokemon Pokemon) GetGender(nameOrId string) (structs.Gender, error) {
	gender, err := pokeapi.Gender(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "gender", nameOrId)
		return structs.Gender{}, err
	}
	return gender, nil
}

func (pokemon Pokemon) GetGenderList() structs.Resource {
	genderList := internal.GetResourceList(genderEndpoint)
	return genderList
}

func (pokemon Pokemon) GetGrowthRate(nameOrId string) (structs.GrowthRate, error) {
	growthRate, err := pokeapi.GrowthRate(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "growth rate", nameOrId)
		return structs.GrowthRate{}, err
	}
	return growthRate, nil
}

func (pokemon Pokemon) GetGrowthRateList() structs.Resource {
	growthRateList := internal.GetResourceList(growthRateEndpoint)
	return growthRateList
}

func (pokemon Pokemon) GetNature(nameOrId string) (structs.Nature, error) {
	nature, err := pokeapi.Nature(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "nature", nameOrId)
		return structs.Nature{}, err
	}
	return nature, nil
}

func (pokemon Pokemon) GetNatureList() structs.Resource {
	natureList := internal.GetResourceList(natureEndpoint)
	return natureList
}

func (pokemon Pokemon) GetPokeathlonStat(nameOrId string) (structs.PokeathlonStat, error) {
	pokeathlonStat, err := pokeapi.PokeathlonStat(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokeathlon stat", nameOrId)
		return structs.PokeathlonStat{}, err
	}
	return pokeathlonStat, nil
}

func (pokemon Pokemon) GetPokeathlonStatList() structs.Resource {
	pokeathlonStatList := internal.GetResourceList(pokeathlonStatEndpoint)
	return pokeathlonStatList
}

func (pokemon Pokemon) GetPokemon(nameOrId string) (structs.Pokemon, error) {
	pokemonResource, err := pokeapi.Pokemon(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokemon", nameOrId)
		return structs.Pokemon{}, err
	}
	return pokemonResource, nil
}

func (pokemon Pokemon) GetPokemonList() structs.Resource {
	pokemonList := internal.GetResourceList(pokemonEndpoint)
	return pokemonList
}

func (pokemon Pokemon) GetPokemonLocationAreas(nameOrId string) []internal.LocationAreaEncounter {
	url := pokemonEndpoint + "/" + nameOrId + "/encounters"
	body := internal.Get(url)
	pokemonLocationAreas := []internal.LocationAreaEncounter{}
	json.Unmarshal(body, &pokemonLocationAreas)
	return pokemonLocationAreas
}

func (pokemon Pokemon) GetPokemonColor(nameOrId string) (structs.PokemonColor, error) {
	pokemonColor, err := pokeapi.PokemonColor(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokemon color", nameOrId)
		return structs.PokemonColor{}, err
	}
	return pokemonColor, nil
}

func (pokemon Pokemon) GetPokemonColorList() structs.Resource {
	pokemonColorList := internal.GetResourceList(pokemonColorEndpoint)
	return pokemonColorList
}

func (pokemon Pokemon) GetPokemonForm(nameOrId string) (structs.PokemonForm, error) {
	pokemonForm, err := pokeapi.PokemonForm(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokemon form", nameOrId)
		return structs.PokemonForm{}, err
	}
	return pokemonForm, nil
}

func (pokemon Pokemon) GetPokemonFormList() structs.Resource {
	pokemonFormList := internal.GetResourceList(pokemonFormEndpoint)
	return pokemonFormList
}

func (pokemon Pokemon) GetPokemonHabitat(nameOrId string) (structs.PokemonHabitat, error) {
	pokemonHabitat, err := pokeapi.PokemonHabitat(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokemon habitat", nameOrId)
		return structs.PokemonHabitat{}, err
	}
	return pokemonHabitat, nil
}

func (pokemon Pokemon) GetPokemonHabitatList() structs.Resource {
	pokemonHabitatList := internal.GetResourceList(pokemonHabitatEndpoint)
	return pokemonHabitatList
}

func (pokemon Pokemon) GetPokemonShape(nameOrId string) (structs.PokemonShape, error) {
	pokemonShape, err := pokeapi.PokemonShape(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokemon shape", nameOrId)
		return structs.PokemonShape{}, err
	}
	return pokemonShape, nil
}

func (pokemon Pokemon) GetPokemonShapeList() structs.Resource {
	pokemonShapeList := internal.GetResourceList(pokemonShapeEndpoint)
	return pokemonShapeList
}

func (pokemon Pokemon) GetPokemonSpecies(nameOrId string) (structs.PokemonSpecies, error) {
	pokemonSpecies, err := pokeapi.PokemonSpecies(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokemon species", nameOrId)
		return structs.PokemonSpecies{}, err
	}
	return pokemonSpecies, nil
}

func (pokemon Pokemon) GetPokemonSpeciesList() structs.Resource {
	pokemonSpeciesList := internal.GetResourceList(pokemonSpeciesEndpoint)
	return pokemonSpeciesList
}

func (pokemon Pokemon) GetStat(nameOrId string) (structs.Stat, error) {
	stat, err := pokeapi.Stat(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "stat", nameOrId)
		return structs.Stat{}, err
	}
	return stat, nil
}

func (pokemon Pokemon) GetStatList() structs.Resource {
	statList := internal.GetResourceList(statEndpoint)
	return statList
}

func (pokemon Pokemon) GetType(nameOrId string) (structs.Type, error) {
	pokemonType, err := pokeapi.Type(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "type", nameOrId)
		return structs.Type{}, err
	}
	return pokemonType, nil
}

func (pokemon Pokemon) GetTypeList() structs.Resource {
	typeList := internal.GetResourceList(typeEndpoint)
	return typeList
}
