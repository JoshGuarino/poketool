package items

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var items IItems = Items{}

func TestGetItem(t *testing.T) {
	rById, _ := items.GetItem("1")
	rByName, _ := items.GetItem("master-ball")
	rFail, _ := items.GetItem("test")
	assert.Equal(t, "master-ball", rById.Name, "Expected to receive 'master-ball'.")
	assert.Equal(t, "master-ball", rByName.Name, "Expected to receive 'master-ball'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemList(t *testing.T) {
	r := items.GetItemList()
	assert.Equal(t, "master-ball", r.Results[0].Name, "Expected to have a list of 'item' resource.")
}

func TestGetItemAttribute(t *testing.T) {
	rById, _ := items.GetItemAttribute("1")
	rByName, _ := items.GetItemAttribute("countable")
	rFail, _ := items.GetItemAttribute("test")
	assert.Equal(t, "countable", rById.Name, "Expected to receive 'countable'.")
	assert.Equal(t, "countable", rByName.Name, "Expected to receive 'countable'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemAttributeList(t *testing.T) {
	r := items.GetItemAttributeList()
	assert.Equal(t, "countable", r.Results[0].Name, "Expected to have a list of 'item attribute' resource.")
}

func TestGetItemCategory(t *testing.T) {
	rById, _ := items.GetItemCategory("1")
	rByName, _ := items.GetItemCategory("stat-boosts")
	rFail, _ := items.GetItemCategory("test")
	assert.Equal(t, "stat-boosts", rById.Name, "Expected to receive 'stat-boosts'.")
	assert.Equal(t, "stat-boosts", rByName.Name, "Expected to receive 'stat-boosts'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemCategoryList(t *testing.T) {
	r := items.GetItemCategoryList()
	assert.Equal(t, "stat-boosts", r.Results[0].Name, "Expected to have a list of 'item category' resource.")
}

func TestGetItemFlingEffect(t *testing.T) {
	rById, _ := items.GetItemFlingEffect("1")
	rByName, _ := items.GetItemFlingEffect("badly-poison")
	rFail, _ := items.GetItemFlingEffect("test")
	assert.Equal(t, "badly-poison", rById.Name, "Expected to receive 'badly-poison'.")
	assert.Equal(t, "badly-poison", rByName.Name, "Expected to receive 'badly-poison'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemFlingEffectList(t *testing.T) {
	r := items.GetItemFlingEffectList()
	assert.Equal(t, "badly-poison", r.Results[0].Name, "Expected to have a list of 'item fling effect' resource.")
}

func TestGetItemPocket(t *testing.T) {
	rById, _ := items.GetItemPocket("1")
	rByName, _ := items.GetItemPocket("misc")
	rFail, _ := items.GetItemPocket("test")
	assert.Equal(t, "misc", rById.Name, "Expected to receive 'misc'.")
	assert.Equal(t, "misc", rByName.Name, "Expected to receive 'misc'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemPocketList(t *testing.T) {
	r := items.GetItemPocketList()
	assert.Equal(t, "misc", r.Results[0].Name, "Expected to have a list of 'item pocket' resource.")
}
