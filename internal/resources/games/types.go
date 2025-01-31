package games

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	games Games
}

type IGames interface {
	GetGeneration(nameOrId string) (structs.Generation, error)
	GetGenerationList() structs.Resource
	GetPokedex(nameOrId string) (structs.Pokedex, error)
	GetPokedexList() structs.Resource
	GetVersion(nameOrId string) (structs.Version, error)
	GetVersionList() structs.Resource
	GetVersionGroup(nameOrId string) (structs.VersionGroup, error)
	GetVersionGroupList() structs.Resource
}

type Games struct{}
