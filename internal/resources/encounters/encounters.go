package encounters

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (encounters Encounters) GetEncounterMethod(nameOrId string) (structs.EncounterMethod, error) {
	encounterMethod, err := pokeapi.EncounterMethod(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "encounter method", nameOrId)
		return structs.EncounterMethod{}, err
	}
	return encounterMethod, nil
}

func (encounters Encounters) GetEncounterMethodList() structs.Resource {
	encounterMethodList := internal.GetResourceList(encounterMethodEndpoint)
	return encounterMethodList
}

func (encounters Encounters) GetEncounterCondition(nameOrId string) (structs.EncounterCondition, error) {
	encounterCondition, err := pokeapi.EncounterCondition(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "encounter condition", nameOrId)
		return structs.EncounterCondition{}, err
	}
	return encounterCondition, nil
}

func (encounters Encounters) GetEncounterConditionList() structs.Resource {
	encounterConditionList := internal.GetResourceList(encounterConditionEndpoint)
	return encounterConditionList
}

func (encounters Encounters) GetEncounterConditionValue(nameOrId string) (structs.EncounterConditionValue, error) {
	encounterConditionValue, err := pokeapi.EncounterConditionValue(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "encounter condition value", nameOrId)
		return structs.EncounterConditionValue{}, err
	}
	return encounterConditionValue, nil
}

func (encounters Encounters) GetEncounterConditionValueList() structs.Resource {
	encounterConditionValueList := internal.GetResourceList(encounterConditionValueEndpoint)
	return encounterConditionValueList
}
