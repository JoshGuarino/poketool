package items

import (
	"testing"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = 

func TestGetList(t *testing.T) {
	rItem := controller.GetList("Item")
	rAttribute := controller.GetList("Item Attribute")
	rCatergory := controller.GetList("Item Category")
	rFling := controller.GetList("Item Fling Effect")
	rPocket := controller.GetList("Item Pocket")
	rFail := controller.GetList("test")
	assert.IsType(t, []structs.Result{}, rItem.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rItem.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rAttribute.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rAttribute.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rCatergory.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rCatergory.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rFling.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rFling.Results, "Expected to not have an empty array.")
	assert.IsType(t, []structs.Result{}, rPocket.Results, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rPocket.Results, "Expected to not have an empty array.")
	assert.Equal(t, structs.Resource{}, rFail, "Expected to have empty struct of type Resource{}.")
}

func TestGetSpecific(t *testing.T) {
	rItem, _ := controller.GetSpecific("Item", "1")
	rAttribute, _ := controller.GetSpecific("Item Attribute", "1")
	rCatergory, _ := controller.GetSpecific("Item Category", "1")
	rFling, _ := controller.GetSpecific("Item Fling Effect", "1")
	rPocket, _ := controller.GetSpecific("Item Pocket", "1")
	rFail, _ := controller.GetSpecific("test", "test")
	assert.IsType(t, structs.Item{}, rItem, "Expected to have type 'Item' struct.")
	assert.IsType(t, structs.ItemAttribute{}, rAttribute, "Expected to have type 'ItemAttribute' struct.")
	assert.IsType(t, structs.ItemCategory{}, rCatergory, "Expected to have type 'ItemCategory' struct.")
	assert.IsType(t, structs.ItemFlingEffect{}, rFling, "Expected to have type 'ItemFlingEffect' struct.")
	assert.IsType(t, structs.ItemPocket{}, rPocket, "Expected to have type 'ItemPocket' struct.")
	assert.Equal(t, nil, rFail, "Expected to have 'nil' value.")
}
