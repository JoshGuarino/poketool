package games

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Generation":
    return GetGenerationList()
	case "Pokedex":
		return GetPokedexList()
	case "Version":
		return GetVersionList()
	case "Version Group":
    return GetVersionGroupList()
  }

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
  case "Generation":
    return GetGeneration(search)
	case "Pokedex":
    return GetPokedex(search)
	case "Version":
    return GetVersion(search)
	case "Version Group":
    return GetVersionGroup(search)
  }

	return "", nil
}
