package contests

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = Controller{contests: Contests{}}

func TestGetList(t *testing.T) {
	rType := controller.GetList("Contest Type")
	rEffect := controller.GetList("Contest Effect")
	rSuper := controller.GetList("Super Contest Effect")
	rFail := controller.GetList("")
	assert.IsType(t, structs.Resource{}, rType, "Expected to have type 'Resource' struct.")
	assert.IsType(t, structs.Resource{}, rEffect, "Expected to have type 'Resource' struct.")
	assert.IsType(t, structs.Resource{}, rSuper, "Expected to have type 'Resource' struct.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rType, _ := controller.GetSpecific("Contest Type", "1")
	rEffect, _ := controller.GetSpecific("Contest Effect", "1")
	rSuper, _ := controller.GetSpecific("Super Contest Effect", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.ContestType{}, rType, "Expected to have type 'ContestType' struct.")
	assert.IsType(t, structs.ContestEffect{}, rEffect, "Expected to have type 'ContestEffect' struct.")
	assert.IsType(t, structs.SuperContestEffect{}, rSuper, "Expected to have type 'SuperContestEffect' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
