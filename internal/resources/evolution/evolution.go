package evolution

import (
	"github.com/JoshGuarino/PokeGo/pkg/models"
	evolutionGroup "github.com/JoshGuarino/PokeGo/pkg/resources/evolution"
)

type IEvolution interface {
	GetEvolutionChain(id string) (*models.EvolutionChain, error)
	GetEvolutionChainList(limit int, offset int) (*models.ResourceList, error)
	GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error)
	GetEvolutionTriggerList(limit int, offset int) (*models.NamedResourceList, error)
}

type Evolution struct {
	evolutionGroup evolutionGroup.Evolution
}

func NewEvolution() Evolution {
	return Evolution{
		evolutionGroup: evolutionGroup.NewEvolutionGroup(),
	}
}

func (e Evolution) GetEvolutionChain(id string) (*models.EvolutionChain, error) {
	evolutionChain, err := e.evolutionGroup.GetEvolutionChain(id)
	if err != nil {
		return nil, err
	}
	return evolutionChain, nil
}

func (e Evolution) GetEvolutionChainList(limit int, offset int) (*models.ResourceList, error) {
	evolutionChainList, err := e.evolutionGroup.GetEvolutionChainList(limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionChainList, nil
}

func (e Evolution) GetEvolutionTrigger(nameOrId string) (*models.EvolutionTrigger, error) {
	evolutionTrigger, err := e.evolutionGroup.GetEvolutionTrigger(nameOrId)
	if err != nil {
		return nil, err
	}
	return evolutionTrigger, nil
}

func (e Evolution) GetEvolutionTriggerList(limit int, offset int) (*models.NamedResourceList, error) {
	evolutionTriggerList, err := e.evolutionGroup.GetEvolutionTriggerList(limit, offset)
	if err != nil {
		return nil, err
	}
	return evolutionTriggerList, nil
}
