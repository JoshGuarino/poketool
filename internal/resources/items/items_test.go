package items

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {
	rById, _ := GetItem("1")
	rByName, _ := GetItem("master-ball")
	rFail, _ := GetItem("test")
	assert.Equal(t, "master-ball", rById.Name, "Expected to receive 'master-ball'.")
	assert.Equal(t, "master-ball", rByName.Name, "Expected to receive 'master-ball'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemList(t *testing.T) {
	r := GetItemList()
	assert.Equal(t, "master-ball", r.Results[0].Name, "Expected to have a list of 'item' resource.")
}

func TestGetItemAttribute(t *testing.T) {
	rById, _ := GetItemAttribute("1")
	rByName, _ := GetItemAttribute("countable")
	rFail, _ := GetItemAttribute("test")
	assert.Equal(t, "countable", rById.Name, "Expected to receive 'countable'.")
	assert.Equal(t, "countable", rByName.Name, "Expected to receive 'countable'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemAttributeList(t *testing.T) {
	r := GetItemAttributeList()
	assert.Equal(t, "countable", r.Results[0].Name, "Expected to have a list of 'item attribute' resource.")
}

func TestGetItemCategory(t *testing.T) {
	rById, _ := GetItemCategory("1")
	rByName, _ := GetItemCategory("stat-boosts")
	rFail, _ := GetItemCategory("test")
	assert.Equal(t, "stat-boosts", rById.Name, "Expected to receive 'stat-boosts'.")
	assert.Equal(t, "stat-boosts", rByName.Name, "Expected to receive 'stat-boosts'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemCategoryList(t *testing.T) {
	r := GetItemCategoryList()
	assert.Equal(t, "stat-boosts", r.Results[0].Name, "Expected to have a list of 'item category' resource.")
}

func TestGetItemFlingEffect(t *testing.T) {
	rById, _ := GetItemFlingEffect("1")
	rByName, _ := GetItemFlingEffect("badly-poison")
	rFail, _ := GetItemFlingEffect("test")
	assert.Equal(t, "badly-poison", rById.Name, "Expected to receive 'badly-poison'.")
	assert.Equal(t, "badly-poison", rByName.Name, "Expected to receive 'badly-poison'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemFlingEffectList(t *testing.T) {
	r := GetItemFlingEffectList()
	assert.Equal(t, "badly-poison", r.Results[0].Name, "Expected to have a list of 'item fling effect' resource.")
}

func TestGetItemPocket(t *testing.T) {
	rById, _ := GetItemPocket("1")
	rByName, _ := GetItemPocket("misc")
	rFail, _ := GetItemPocket("test")
	assert.Equal(t, "misc", rById.Name, "Expected to receive 'misc'.")
	assert.Equal(t, "misc", rByName.Name, "Expected to receive 'misc'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetItemPocketList(t *testing.T) {
	r := GetItemPocketList()
	assert.Equal(t, "misc", r.Results[0].Name, "Expected to have a list of 'item pocket' resource.")
}
