package encounters

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	encounters Encounters
}

type IEncounters interface {
	GetEncounterMethod(nameOrId string) (structs.EncounterMethod, error)
	GetEncounterMethodList() structs.Resource
	GetEncounterCondition(nameOrId string) (structs.EncounterCondition, error)
	GetEncounterConditionList() structs.Resource
	GetEncounterConditionValue(nameOrId string) (structs.EncounterConditionValue, error)
	GetEncounterConditionValueList() structs.Resource
}

type Encounters struct{}
