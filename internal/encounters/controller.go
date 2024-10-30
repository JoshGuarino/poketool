package encounters

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Encounter Method":
		return GetEncounterMethodList()
	case "Encounter Condition":
		return GetEncounterConditionList()
	case "Encounter Condtiion Value":
		return GetEncounterConditionValueList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Encounter Method":
		return GetEncounterMethod(search)
	case "Encounter Condition":
		return GetEncounterCondition(search)
	case "Encounter Condtiion Value":
		return GetEncounterConditionValue(search)
	}

	return "", nil
}
