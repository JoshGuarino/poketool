package contests

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rType, _ := controller.GetList("Contest Type", 20, 0)
	rEffect, _ := controller.GetList("Contest Effect", 20, 0)
	rSuper, _ := controller.GetList("Super Contest Effect", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.NamedResourceList{}, rType, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rType, "Expected to not have an empty struct.")
	assert.IsType(t, &models.ResourceList{}, rEffect, "Expected to have array of type 'ResourceList' struct.")
	assert.NotEmpty(t, rEffect, "Expected to not have an empty struct.")
	assert.IsType(t, &models.ResourceList{}, rSuper, "Expected to have array of type 'ResourceList' struct.")
	assert.NotEmpty(t, rSuper, "Expected to not have an empty struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rType, _ := controller.GetSpecific("Contest Type", "1")
	rEffect, _ := controller.GetSpecific("Contest Effect", "1")
	rSuper, _ := controller.GetSpecific("Super Contest Effect", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.ContestType{}, rType, "Expected to have type 'ContestType' struct.")
	assert.IsType(t, &models.ContestEffect{}, rEffect, "Expected to have type 'ContestEffect' struct.")
	assert.IsType(t, &models.SuperContestEffect{}, rSuper, "Expected to have type 'SuperContestEffect' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
