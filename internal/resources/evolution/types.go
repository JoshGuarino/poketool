package evolution

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	evolution Evolution
}

type IEvolution interface {
	GetEvolutionChain(id string) (structs.EvolutionChain, error)
	GetEvolutionChainList() structs.Resource
	GetEvolutionTrigger(nameOrId string) (structs.EvolutionTrigger, error)
	GetEvolutionTriggerList() structs.Resource
}

type Evolution struct{}
