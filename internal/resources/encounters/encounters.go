package encounters

import (
	"github.com/JoshGuarino/PokeGo/pkg/models"
	encountersGroup "github.com/JoshGuarino/PokeGo/pkg/resources/encounters"
)

type IEncounters interface {
	GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error)
	GetEncounterMethodList(limit int, offset int) (*models.NamedResourceList, error)
	GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error)
	GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error)
	GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error)
	GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error)
}

type Encounters struct {
	encountersGroup encountersGroup.Encounters
}

func NewEncounters() Encounters {
	return Encounters{
		encountersGroup: encountersGroup.NewEncountersGroup(),
	}
}

func (encounters Encounters) GetEncounterMethod(nameOrId string) (*models.EncounterMethod, error) {
	encounterMethod, err := encounters.encountersGroup.GetEncounterMethod(nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterMethod, nil
}

func (encounters Encounters) GetEncounterMethodList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterMethodList, err := encounters.encountersGroup.GetEncounterMethodList(limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterMethodList, nil
}

func (encounters Encounters) GetEncounterCondition(nameOrId string) (*models.EncounterCondition, error) {
	encounterCondition, err := encounters.encountersGroup.GetEncounterCondition(nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterCondition, nil
}

func (encounters Encounters) GetEncounterConditionList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionList, err := encounters.encountersGroup.GetEncounterConditionList(limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionList, nil
}

func (encounters Encounters) GetEncounterConditionValue(nameOrId string) (*models.EncounterConditionValue, error) {
	encounterConditionValue, err := encounters.encountersGroup.GetEncounterConditionValue(nameOrId)
	if err != nil {
		return nil, err
	}
	return encounterConditionValue, nil
}

func (encounters Encounters) GetEncounterConditionValueList(limit int, offset int) (*models.NamedResourceList, error) {
	encounterConditionValueList, err := encounters.encountersGroup.GetEncounterConditionValueList(limit, offset)
	if err != nil {
		return nil, err
	}
	return encounterConditionValueList, nil
}
