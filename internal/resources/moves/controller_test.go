package moves

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{moves: Moves{}}

func TestGetList(t *testing.T) {
	rMove := controller.GetList("Move")
	rAilment := controller.GetList("Move Ailment")
	rBattle := controller.GetList("Move Battle Style")
	rCategory := controller.GetList("Move Category")
	rDamage := controller.GetList("Move Damage Class")
	rLearn := controller.GetList("Move Learn Method")
	rTarget := controller.GetList("Move Target")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rMove.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rAilment.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rBattle.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rCategory.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rDamage.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rLearn.Results, "Expected to have array of type 'Result' struct.")
	assert.IsType(t, []structs.Result{}, rTarget.Results, "Expected to have array of type 'Result' struct.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rMove, _ := controller.GetSpecific("Move", "1")
	rAilment, _ := controller.GetSpecific("Move Ailment", "1")
	rBattle, _ := controller.GetSpecific("Move Battle Style", "1")
	rCategory, _ := controller.GetSpecific("Move Category", "1")
	rDamage, _ := controller.GetSpecific("Move Damage Class", "1")
	rLearn, _ := controller.GetSpecific("Move Learn Method", "1")
	rTarget, _ := controller.GetSpecific("Move Target", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Move{}, rMove, "Expected to have type 'Move' struct.")
	assert.IsType(t, structs.MoveAilment{}, rAilment, "Expected to have type 'MoveAilment' struct.")
	assert.IsType(t, structs.MoveBattleStyle{}, rBattle, "Expected to have type 'MoveBattleStyle' struct.")
	assert.IsType(t, structs.MoveCategory{}, rCategory, "Expected to have type 'MoveCategory' struct.")
	assert.IsType(t, structs.MoveDamageClass{}, rDamage, "Expected to have type 'MoveDamageClass' struct.")
	assert.IsType(t, structs.MoveLearnMethod{}, rLearn, "Expected to have type 'MoveLearnMethod' struct.")
	assert.IsType(t, structs.MoveTarget{}, rTarget, "Expected to have type 'MoveTarget' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
