package evolution

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Evolution Chain":
		return controller.evolution.GetEvolutionChainList()
	case "Evolution Trigger":
		return controller.evolution.GetEvolutionTriggerList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Evolution Chain":
		return controller.evolution.GetEvolutionChain(search)
	case "Evolution Trigger":
		return controller.evolution.GetEvolutionTrigger(search)
	}

	return nil, nil
}
