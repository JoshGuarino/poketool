package encounters

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetEncounterMethod(nameOrId string) structs.EncounterMethod {
	encounterMethod, err := pokeapi.EncounterMethod(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return encounterMethod
}

func GetEncounterMethodList() structs.Resource {
	encounterMethodList := internal.GetResourceList(encounterMethodEndpoint)
	return encounterMethodList
}

func GetEncounterCondition(nameOrId string) structs.EncounterCondition {
	encounterCondition, err := pokeapi.EncounterCondition(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return encounterCondition
}

func GetEncounterConditionList() structs.Resource {
	encounterConditionList := internal.GetResourceList(encounterConditionEndpoint)
	return encounterConditionList
}

func GetEncounterConditionValue(nameOrId string) structs.EncounterConditionValue {
	encounterConditionValue, err := pokeapi.EncounterConditionValue(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return encounterConditionValue
}

func GetEncounterConditionValueList() structs.Resource {
	encounterConditionValueList := internal.GetResourceList(encounterConditionValueEndpoint)
	return encounterConditionValueList
}
