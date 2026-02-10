package encounters

import (
	"fmt"

	encountersGroup "github.com/JoshGuarino/PokeGo/pkg/resources/encounters"
)

type EncountersController struct {
	encounters encountersGroup.Encounters
}

func NewController() EncountersController {
	return EncountersController{
		encounters: encountersGroup.NewEncountersGroup(),
	}
}

func (c EncountersController) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Encounter Method":
		return c.encounters.GetEncounterMethodList(limit, offset)
	case "Encounter Condition":
		return c.encounters.GetEncounterConditionList(limit, offset)
	case "Encounter Condition Value":
		return c.encounters.GetEncounterConditionValueList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c EncountersController) GetSpecific(result string, search string) (any, error) {
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
