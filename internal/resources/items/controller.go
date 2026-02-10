package items

import (
	"fmt"

	itemsGroup "github.com/JoshGuarino/PokeGo/pkg/resources/items"
)

type ItemsController struct {
	items itemsGroup.Items
}

func NewController() ItemsController {
	return ItemsController{
		items: itemsGroup.NewItemsGroup(),
	}
}

func (c ItemsController) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Item":
		return c.items.GetItemList(limit, offset)
	case "Item Attribute":
		return c.items.GetItemAttributeList(limit, offset)
	case "Item Category":
		return c.items.GetItemCategoryList(limit, offset)
	case "Item Fling Effect":
		return c.items.GetItemFlingEffectList(limit, offset)
	case "Item Pocket":
		return c.items.GetItemPocketList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c ItemsController) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Item":
		return c.items.GetItem(search)
	case "Item Attribute":
		return c.items.GetItemAttribute(search)
	case "Item Category":
		return c.items.GetItemCategory(search)
	case "Item Fling Effect":
		return c.items.GetItemFlingEffect(search)
	case "Item Pocket":
		return c.items.GetItemPocket(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
