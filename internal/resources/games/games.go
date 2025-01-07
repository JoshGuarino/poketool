package games

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetGeneration(nameOrId string) (structs.Generation, error) {
	generation, err := pokeapi.Generation(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "generation", nameOrId)
		return structs.Generation{}, err
	}
	return generation, nil
}

func GetGenerationList() structs.Resource {
	generationList := internal.GetResourceList(generationEndpoint)
	return generationList
}

func GetPokedex(nameOrId string) (structs.Pokedex, error) {
	pokedex, err := pokeapi.Pokedex(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pokedex", nameOrId)
    return structs.Pokedex{}, err
	}
	return pokedex, nil
}

func GetPokedexList() structs.Resource {
	pokedexList := internal.GetResourceList(pokedexEndpoint)
	return pokedexList
}

func GetVersion(nameOrId string) (structs.Version, error) {
	version, err := pokeapi.Version(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "version", nameOrId)
    return structs.Version{}, err
	}
	return version, nil
}

func GetVersionList() structs.Resource {
	versionList := internal.GetResourceList(versionEndpoint)
	return versionList
}

func GetVersionGroup(nameOrId string) (structs.VersionGroup, error) {
	versionGroup, err := pokeapi.VersionGroup(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "version group", nameOrId)
    return structs.VersionGroup{}, err
  }
	return versionGroup, nil
}

func GetVersionGroupList() structs.Resource {
	versionGroupList := internal.GetResourceList(versionGroupEndpoint)
	return versionGroupList
}
