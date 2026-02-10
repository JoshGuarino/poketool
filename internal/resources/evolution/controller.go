package evolution

import (
	"fmt"

	evolutionGroup "github.com/JoshGuarino/PokeGo/pkg/resources/evolution"
)

type EvolutionController struct {
	evolution evolutionGroup.Evolution
}

func NewController() EvolutionController {
	return EvolutionController{
		evolution: evolutionGroup.NewEvolutionGroup(),
	}
}

func (c EvolutionController) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Evolution Chain":
		return c.evolution.GetEvolutionChainList(limit, offset)
	case "Evolution Trigger":
		return c.evolution.GetEvolutionTriggerList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c EvolutionController) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Evolution Chain":
		return c.evolution.GetEvolutionChain(search)
	case "Evolution Trigger":
		return c.evolution.GetEvolutionTrigger(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
