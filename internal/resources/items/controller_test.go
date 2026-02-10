package items

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rItem, _ := controller.GetList("Item", 20, 0)
	rAttribute, _ := controller.GetList("Item Attribute", 20, 0)
	rCatergory, _ := controller.GetList("Item Category", 20, 0)
	rFling, _ := controller.GetList("Item Fling Effect", 20, 0)
	rPocket, _ := controller.GetList("Item Pocket", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.NamedResourceList{}, rItem, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rItem, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rAttribute, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rAttribute, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rCatergory, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rCatergory, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rFling, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rFling, "Expected to not have an empty array.")
	assert.IsType(t, &models.NamedResourceList{}, rPocket, "Expected to have array of type 'Result' struct.")
	assert.NotEmpty(t, rPocket, "Expected to not have an empty array.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rItem, _ := controller.GetSpecific("Item", "1")
	rAttribute, _ := controller.GetSpecific("Item Attribute", "1")
	rCatergory, _ := controller.GetSpecific("Item Category", "1")
	rFling, _ := controller.GetSpecific("Item Fling Effect", "1")
	rPocket, _ := controller.GetSpecific("Item Pocket", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.Item{}, rItem, "Expected to have type 'Item' struct.")
	assert.IsType(t, &models.ItemAttribute{}, rAttribute, "Expected to have type 'ItemAttribute' struct.")
	assert.IsType(t, &models.ItemCategory{}, rCatergory, "Expected to have type 'ItemCategory' struct.")
	assert.IsType(t, &models.ItemFlingEffect{}, rFling, "Expected to have type 'ItemFlingEffect' struct.")
	assert.IsType(t, &models.ItemPocket{}, rPocket, "Expected to have type 'ItemPocket' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
