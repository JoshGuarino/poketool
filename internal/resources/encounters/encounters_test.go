package encounters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var encounters IEncounters = Encounters{}

func TestGetEncounterMethod(t *testing.T) {
	rById, _ := encounters.GetEncounterMethod("1")
	rByName, _ := encounters.GetEncounterMethod("walk")
	rFail, _ := encounters.GetEncounterMethod("test")
	assert.Equal(t, "walk", rById.Name, "Expected to receive 'walk'.")
	assert.Equal(t, "walk", rByName.Name, "Expected to receive 'walk'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEncounterMethodList(t *testing.T) {
	r := encounters.GetEncounterMethodList()
	assert.Equal(t, "walk", r.Results[0].Name, "Expected to have a list of 'encounter method' resource.")
}

func TestGetEncounterCondition(t *testing.T) {
	rById, _ := encounters.GetEncounterCondition("1")
	rByName, _ := encounters.GetEncounterCondition("swarm")
	rFail, _ := encounters.GetEncounterCondition("test")
	assert.Equal(t, "swarm", rById.Name, "Expected to receive 'swarm'.")
	assert.Equal(t, "swarm", rByName.Name, "Expected to receive 'swarm'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEncounterConditionList(t *testing.T) {
	r := encounters.GetEncounterConditionList()
	assert.Equal(t, "swarm", r.Results[0].Name, "Expected to have a list of 'encounter condition' resource.")
}

func TestGetEncounterConditionValue(t *testing.T) {
	rById, _ := encounters.GetEncounterConditionValue("1")
	rByName, _ := encounters.GetEncounterConditionValue("swarm-yes")
	rFail, _ := encounters.GetEncounterConditionValue("test")
	assert.Equal(t, "swarm-yes", rById.Name, "Expected to receive 'swarm-yes'.")
	assert.Equal(t, "swarm-yes", rByName.Name, "Expected to receive 'swarm-yes'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetEncounterConditionValueList(t *testing.T) {
	r := encounters.GetEncounterConditionValueList()
	assert.Equal(t, "swarm-yes", r.Results[0].Name, "Expected to have a list of 'encounter condition value' resource.")
}
