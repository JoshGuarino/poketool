package encounters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var encounters IEncounters = NewEncounters()

func TestGetEncounterMethod(t *testing.T) {
	rById, _ := encounters.GetEncounterMethod("1")
	rByName, _ := encounters.GetEncounterMethod("walk")
	_, err := encounters.GetEncounterMethod("test")
	assert.Equal(t, "walk", rById.Name, "Expected to receive 'walk'.")
	assert.Equal(t, "walk", rByName.Name, "Expected to receive 'walk'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterMethodList(t *testing.T) {
	r, _ := encounters.GetEncounterMethodList(20, 0)
	assert.Equal(t, "walk", r.Results[0].Name, "Expected to have a list of 'encounter method' resource.")
}

func TestGetEncounterCondition(t *testing.T) {
	rById, _ := encounters.GetEncounterCondition("1")
	rByName, _ := encounters.GetEncounterCondition("swarm")
	_, err := encounters.GetEncounterCondition("test")
	assert.Equal(t, "swarm", rById.Name, "Expected to receive 'swarm'.")
	assert.Equal(t, "swarm", rByName.Name, "Expected to receive 'swarm'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionList(t *testing.T) {
	r, _ := encounters.GetEncounterConditionList(20, 0)
	assert.Equal(t, "swarm", r.Results[0].Name, "Expected to have a list of 'encounter condition' resource.")
}

func TestGetEncounterConditionValue(t *testing.T) {
	rById, _ := encounters.GetEncounterConditionValue("1")
	rByName, _ := encounters.GetEncounterConditionValue("swarm-yes")
	_, err := encounters.GetEncounterConditionValue("test")
	assert.Equal(t, "swarm-yes", rById.Name, "Expected to receive 'swarm-yes'.")
	assert.Equal(t, "swarm-yes", rByName.Name, "Expected to receive 'swarm-yes'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetEncounterConditionValueList(t *testing.T) {
	r, _ := encounters.GetEncounterConditionValueList(20, 0)
	assert.Equal(t, "swarm-yes", r.Results[0].Name, "Expected to have a list of 'encounter condition value' resource.")
}
