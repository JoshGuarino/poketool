package games

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Generation":
		return controller.games.GetGenerationList()
	case "Pokedex":
		return controller.games.GetPokedexList()
	case "Version":
		return controller.games.GetVersionList()
	case "Version Group":
		return controller.games.GetVersionGroupList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Generation":
		return controller.games.GetGeneration(search)
	case "Pokedex":
		return controller.games.GetPokedex(search)
	case "Version":
		return controller.games.GetVersion(search)
	case "Version Group":
		return controller.games.GetVersionGroup(search)
	}

	return nil, nil
}
