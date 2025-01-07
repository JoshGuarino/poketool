package evolution

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Evolution Chain":
		return GetEvolutionChainList()
	case "Evolution Trigger":
		return GetEvolutionTriggerList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Evolution Chain":
		return GetEvolutionChain(search)
	case "Evolution Trigger":
    return GetEvolutionTrigger(search)
	}

	return "", nil
}
