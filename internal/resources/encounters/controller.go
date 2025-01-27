package encounters

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Encounter Method":
		return controller.encounters.GetEncounterMethodList()
	case "Encounter Condition":
		return controller.encounters.GetEncounterConditionList()
	case "Encounter Condition Value":
		return controller.encounters.GetEncounterConditionValueList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Encounter Method":
		return controller.encounters.GetEncounterMethod(search)
	case "Encounter Condition":
		return controller.encounters.GetEncounterCondition(search)
	case "Encounter Condition Value":
		return controller.encounters.GetEncounterConditionValue(search)
	}

	return nil, nil
}
