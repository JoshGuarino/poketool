package moves

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rMove, _ := controller.GetList("Move", 20, 0)
	rAilment, _ := controller.GetList("Move Ailment", 20, 0)
	rBattle, _ := controller.GetList("Move Battle Style", 20, 0)
	rCategory, _ := controller.GetList("Move Category", 20, 0)
	rDamage, _ := controller.GetList("Move Damage Class", 20, 0)
	rLearn, _ := controller.GetList("Move Learn Method", 20, 0)
	rTarget, _ := controller.GetList("Move Target", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.NamedResourceList{}, rMove, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rMove, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rAilment, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rAilment, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rBattle, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rBattle, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rCategory, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rCategory, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rDamage, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rDamage, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rLearn, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rLearn, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rTarget, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rTarget, "Expected to not have an empty struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rMove, _ := controller.GetSpecific("Move", "1")
	rAilment, _ := controller.GetSpecific("Move Ailment", "1")
	rBattle, _ := controller.GetSpecific("Move Battle Style", "1")
	rCategory, _ := controller.GetSpecific("Move Category", "1")
	rDamage, _ := controller.GetSpecific("Move Damage Class", "1")
	rLearn, _ := controller.GetSpecific("Move Learn Method", "1")
	rTarget, _ := controller.GetSpecific("Move Target", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.Move{}, rMove, "Expected to have type 'Move' struct.")
	assert.IsType(t, &models.MoveAilment{}, rAilment, "Expected to have type 'MoveAilment' struct.")
	assert.IsType(t, &models.MoveBattleStyle{}, rBattle, "Expected to have type 'MoveBattleStyle' struct.")
	assert.IsType(t, &models.MoveCategory{}, rCategory, "Expected to have type 'MoveCategory' struct.")
	assert.IsType(t, &models.MoveDamageClass{}, rDamage, "Expected to have type 'MoveDamageClass' struct.")
	assert.IsType(t, &models.MoveLearnMethod{}, rLearn, "Expected to have type 'MoveLearnMethod' struct.")
	assert.IsType(t, &models.MoveTarget{}, rTarget, "Expected to have type 'MoveTarget' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
