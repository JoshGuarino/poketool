package pokemon

import (
	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
)

type Controller struct {
	pokemon Pokemon
}

type IPokemon interface {
	GetAbility(nameOrId string) (structs.Ability, error)
	GetAbilityList() structs.Resource
	GetCharacteristic(id string) (structs.Characteristic, error)
	GetCharacteristicList() structs.Resource
	GetEggGroup(nameOrId string) (structs.EggGroup, error)
	GetEggGroupList() structs.Resource
	GetGender(nameOrId string) (structs.Gender, error)
	GetGenderList() structs.Resource
	GetGrowthRate(nameOrId string) (structs.GrowthRate, error)
	GetGrowthRateList() structs.Resource
	GetNature(nameOrId string) (structs.Nature, error)
	GetNatureList() structs.Resource
	GetPokeathlonStat(nameOrId string) (structs.PokeathlonStat, error)
	GetPokeathlonStatList() structs.Resource
	GetPokemon(nameOrId string) (structs.Pokemon, error)
	GetPokemonList() structs.Resource
	GetPokemonLocationAreas(nameOrId string) []internal.LocationAreaEncounter
	GetPokemonColor(nameOrId string) (structs.PokemonColor, error)
	GetPokemonColorList() structs.Resource
	GetPokemonForm(nameOrId string) (structs.PokemonForm, error)
	GetPokemonFormList() structs.Resource
	GetPokemonHabitat(nameOrId string) (structs.PokemonHabitat, error)
	GetPokemonHabitatList() structs.Resource
	GetPokemonShape(nameOrId string) (structs.PokemonShape, error)
	GetPokemonShapeList() structs.Resource
	GetPokemonSpecies(nameOrId string) (structs.PokemonSpecies, error)
	GetPokemonSpeciesList() structs.Resource
	GetStat(nameOrId string) (structs.Stat, error)
	GetStatList() structs.Resource
	GetType(nameOrId string) (structs.Type, error)
	GetTypeList() structs.Resource
}

type Pokemon struct{}
