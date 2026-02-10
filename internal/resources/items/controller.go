package items

import (
	"fmt"

	itemsGroup "github.com/JoshGuarino/PokeGo/pkg/resources/items"
)

type Controller struct {
	items itemsGroup.Items
}

func (controller Controller) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Item":
		return controller.items.GetItemList(limit, offset)
	case "Item Attribute":
		return controller.items.GetItemAttributeList(limit, offset)
	case "Item Category":
		return controller.items.GetItemCategoryList(limit, offset)
	case "Item Fling Effect":
		return controller.items.GetItemFlingEffectList(limit, offset)
	case "Item Pocket":
		return controller.items.GetItemPocketList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Item":
		return controller.items.GetItem(search)
	case "Item Attribute":
		return controller.items.GetItemAttribute(search)
	case "Item Category":
		return controller.items.GetItemCategory(search)
	case "Item Fling Effect":
		return controller.items.GetItemFlingEffect(search)
	case "Item Pocket":
		return controller.items.GetItemPocket(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
