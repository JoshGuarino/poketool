package moves

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var moves IMoves = Moves{}

func TestGetMove(t *testing.T) {
	rById, _ := moves.GetMove("1")
	rByName, _ := moves.GetMove("pound")
	rFail, _ := moves.GetMove("test")
	assert.Equal(t, "pound", rById.Name, "Expected to receive 'pound'.")
	assert.Equal(t, "pound", rByName.Name, "Expected to receive 'pound'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveList(t *testing.T) {
	r := moves.GetMoveList()
	assert.Equal(t, "pound", r.Results[0].Name, "Expected to have a list of 'move' resource.")
}

func TestGetMoveAilment(t *testing.T) {
	rById, _ := moves.GetMoveAilment("1")
	rByName, _ := moves.GetMoveAilment("paralysis")
	rFail, _ := moves.GetMoveAilment("test")
	assert.Equal(t, "paralysis", rById.Name, "Expected to receive 'paralysis'.")
	assert.Equal(t, "paralysis", rByName.Name, "Expected to receive 'paralysis'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveAilmentList(t *testing.T) {
	r := moves.GetMoveAilmentList()
	assert.Equal(t, "unknown", r.Results[0].Name, "Expected to have a list of 'move ailment' resource.")
}

func TestGetMoveBattleStyle(t *testing.T) {
	rById, _ := moves.GetMoveBattleStyle("1")
	rByName, _ := moves.GetMoveBattleStyle("attack")
	rFail, _ := moves.GetMoveBattleStyle("test")
	assert.Equal(t, "attack", rById.Name, "Expected to receive 'attack'.")
	assert.Equal(t, "attack", rByName.Name, "Expected to receive 'attack'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveBattleStyleList(t *testing.T) {
	r := moves.GetMoveBattleStyleList()
	assert.Equal(t, "attack", r.Results[0].Name, "Expected to have a list of 'move' resource.")
}

func TestGetMoveCategory(t *testing.T) {
	rById, _ := moves.GetMoveCategory("1")
	rByName, _ := moves.GetMoveCategory("ailment")
	rFail, _ := moves.GetMoveCategory("test")
	assert.Equal(t, "ailment", rById.Name, "Expected to receive 'pound'.")
	assert.Equal(t, "ailment", rByName.Name, "Expected to receive 'pound'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveCategoryList(t *testing.T) {
	r := moves.GetMoveCategoryList()
	assert.Equal(t, "damage", r.Results[0].Name, "Expected to have a list of 'move' resource.")
}

func TestGetMoveDamageClass(t *testing.T) {
	rById, _ := moves.GetMoveDamageClass("1")
	rByName, _ := moves.GetMoveDamageClass("status")
	rFail, _ := moves.GetMoveDamageClass("test")
	assert.Equal(t, "status", rById.Name, "Expected to receive 'status'.")
	assert.Equal(t, "status", rByName.Name, "Expected to receive 'status'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveDamageClassList(t *testing.T) {
	r := moves.GetMoveDamageClassList()
	assert.Equal(t, "status", r.Results[0].Name, "Expected to have a list of 'move damaage class' resource.")
}
func TestGetMoveLearnMethod(t *testing.T) {
	rById, _ := moves.GetMoveLearnMethod("1")
	rByName, _ := moves.GetMoveLearnMethod("level-up")
	rFail, _ := moves.GetMoveLearnMethod("test")
	assert.Equal(t, "level-up", rById.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "level-up", rByName.Name, "Expected to receive 'level-up'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveLearnMethodList(t *testing.T) {
	r := moves.GetMoveLearnMethodList()
	assert.Equal(t, "level-up", r.Results[0].Name, "Expected to have a list of 'move learn method' resource.")
}

func TestGetMoveTarget(t *testing.T) {
	rById, _ := moves.GetMoveTarget("1")
	rByName, _ := moves.GetMoveTarget("specific-move")
	rFail, _ := moves.GetMoveTarget("test")
	assert.Equal(t, "specific-move", rById.Name, "Expected to receive 'specific-move'.")
	assert.Equal(t, "specific-move", rByName.Name, "Expected to receive 'specific-move'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetMoveTargetList(t *testing.T) {
	r := moves.GetMoveTargetList()
	assert.Equal(t, "specific-move", r.Results[0].Name, "Expected to have a list of 'move target' resource.")
}
