package encounters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEncounterMethod(t *testing.T) {
	rById, _ := GetEncounterMethod("1")
	rByName, _ := GetEncounterMethod("walk")
	rFail, _ := GetEncounterMethod("test")
	assert.Equal(t, "walk", rById.Name, "Expected to receive 'walk'.")
	assert.Equal(t, "walk", rByName.Name, "Expected to receive 'walk'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEncounterMethodList(t *testing.T) {
	r := GetEncounterMethodList()
	assert.Equal(t, "walk", r.Results[0].Name, "Expected to have a list of 'encounter method' resource.")
}

func TestGetEncounterCondition(t *testing.T) {
	rById, _ := GetEncounterCondition("1")
	rByName, _ := GetEncounterCondition("swarm")
	rFail, _ := GetEncounterCondition("test")
	assert.Equal(t, "swarm", rById.Name, "Expected to receive 'swarm'.")
	assert.Equal(t, "swarm", rByName.Name, "Expected to receive 'swarm'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEncounterConditionList(t *testing.T) {
	r := GetEncounterConditionList()
	assert.Equal(t, "swarm", r.Results[0].Name, "Expected to have a list of 'encounter condition' resource.")
}

func TestGetEncounterConditionValue(t *testing.T) {
	rById, _ := GetEncounterConditionValue("1")
	rByName, _ := GetEncounterConditionValue("swarm-yes")
	rFail, _ := GetEncounterConditionValue("test")
	assert.Equal(t, "swarm-yes", rById.Name, "Expected to receive 'swarm-yes'.")
	assert.Equal(t, "swarm-yes", rByName.Name, "Expected to receive 'swarm-yes'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEncounterConditionValueList(t *testing.T) {
	r := GetEncounterConditionValueList()
	assert.Equal(t, "swarm-yes", r.Results[0].Name, "Expected to have a list of 'encounter condition' resource.")
}
