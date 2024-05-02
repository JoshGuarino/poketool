package evolution

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetEvolutionChain(nameOrId string) structs.EvolutionChain {
	evolutionChain, err := pokeapi.EvolutionChain(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return evolutionChain
}

func GetEvolutionChainList() structs.Resource {
	evolutionChainList := internal.GetResourceList(evolutionChainEndpoint)
	return evolutionChainList
}

func GetEvolutionTrigger(nameOrId string) structs.EvolutionTrigger {
	evolutionTrigger, err := pokeapi.EvolutionTrigger(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return evolutionTrigger
}

func GetEvolutionTriggerList() structs.Resource {
	evolutionTriggerList := internal.GetResourceList(evolutionTriggerEndpoint)
	return evolutionTriggerList
}
