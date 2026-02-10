package encounters

import (
	"fmt"

	encountersGroup "github.com/JoshGuarino/PokeGo/pkg/resources/encounters"
)

type IController interface {
	GetList(result string) (any, error)
	GetSpecific(result string, search string) (any, error)
}

type Controller struct {
	encounters encountersGroup.Encounters
}

func NewController() Controller {
	return Controller{
		encounters: encountersGroup.NewEncountersGroup(),
	}
}

func (c Controller) GetList(result string) (any, error) {
	switch result {
	case "Encounter Method":
		return c.encounters.GetEncounterMethodList(20, 0)
	case "Encounter Condition":
		return c.encounters.GetEncounterConditionList(20, 0)
	case "Encounter Condition Value":
		return c.encounters.GetEncounterConditionValueList(20, 0)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c Controller) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Encounter Method":
		return c.encounters.GetEncounterMethod(search)
	case "Encounter Condition":
		return c.encounters.GetEncounterCondition(search)
	case "Encounter Condition Value":
		return c.encounters.GetEncounterConditionValue(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
