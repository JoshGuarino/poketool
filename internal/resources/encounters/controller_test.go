package encounters

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rMethod, _ := controller.GetList("Encounter Method", 20, 0)
	rCondition, _ := controller.GetList("Encounter Condition", 20, 0)
	rValue, _ := controller.GetList("Encounter Condition Value", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.NamedResourceList{}, rMethod, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rMethod, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rCondition, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rCondition, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rValue, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rValue, "Expected to not have an empty struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rMethod, _ := controller.GetSpecific("Encounter Method", "1")
	rCondition, _ := controller.GetSpecific("Encounter Condition", "1")
	rValue, _ := controller.GetSpecific("Encounter Condition Value", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.EncounterMethod{}, rMethod, "Expected to have type 'EncounterMethod' struct.")
	assert.IsType(t, &models.EncounterCondition{}, rCondition, "Expected to have type 'EncounterCondition' struct.")
	assert.IsType(t, &models.EncounterConditionValue{}, rValue, "Expected to have type 'EncounterConditionValue' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
