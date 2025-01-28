package evolution

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (evolution Evolution) GetEvolutionChain(id string) (structs.EvolutionChain, error) {
	evolutionChain, err := pokeapi.EvolutionChain(id)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "evolution chain", id)
		return structs.EvolutionChain{}, err
	}
	return evolutionChain, nil
}

func (evolution Evolution) GetEvolutionChainList() structs.Resource {
	evolutionChainList := internal.GetResourceList(evolutionChainEndpoint)
	return evolutionChainList
}

func (evolution Evolution) GetEvolutionTrigger(nameOrId string) (structs.EvolutionTrigger, error) {
	evolutionTrigger, err := pokeapi.EvolutionTrigger(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "evolution trigger", nameOrId)
		return structs.EvolutionTrigger{}, err
	}
	return evolutionTrigger, nil
}

func (evolution Evolution) GetEvolutionTriggerList() structs.Resource {
	evolutionTriggerList := internal.GetResourceList(evolutionTriggerEndpoint)
	return evolutionTriggerList
}
