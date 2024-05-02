package games

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetGeneration(nameOrId string) structs.Generation {
	generation, err := pokeapi.Generation(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return generation
}

func GetGenerationList() structs.Resource {
	generationList := internal.GetResourceList(generationEndpoint)
	return generationList
}

func GetPokedex(nameOrId string) structs.Pokedex {
	pokedex, err := pokeapi.Pokedex(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pokedex
}

func GetPokedexList() structs.Resource {
	pokedexList := internal.GetResourceList(pokedexEndpoint)
	return pokedexList
}

func GetVersion(nameOrId string) structs.Version {
	version, err := pokeapi.Version(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return version
}

func GetVersionList() structs.Resource {
	versionList := internal.GetResourceList(versionEndpoint)
	return versionList
}

func GetVersionGroup(nameOrId string) structs.VersionGroup {
	versionGroup, err := pokeapi.VersionGroup(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return versionGroup
}

func GetVersionGroupList() structs.Resource {
	versionGroupList := internal.GetResourceList(versionGroupEndpoint)
	return versionGroupList
}
