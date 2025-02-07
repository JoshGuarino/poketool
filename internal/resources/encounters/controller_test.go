package encounters

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{encounters: Encounters{}}

func TestGetList(t *testing.T) {
	rMethod := controller.GetList("Encounter Method")
	rCondition := controller.GetList("Encounter Condition")
	rValue := controller.GetList("Encounter Condition Value")
	rFail := controller.GetList("")
	assert.IsType(t, []structs.Result{}, rMethod.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rMethod.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rCondition.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rCondition.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rValue.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rValue.Results, "Expected to not have an empty array.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rMethod, _ := controller.GetSpecific("Encounter Method", "1")
	rCondition, _ := controller.GetSpecific("Encounter Condition", "1")
	rValue, _ := controller.GetSpecific("Encounter Condition Value", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.EncounterMethod{}, rMethod, "Expected to have type 'EncounterMethod' struct.")
	assert.IsType(t, structs.EncounterCondition{}, rCondition, "Expected to have type 'EncounterCondition' struct.")
	assert.IsType(t, structs.EncounterConditionValue{}, rValue, "Expected to have type 'EncounterConditionValue' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
