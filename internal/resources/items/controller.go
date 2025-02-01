package items

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Item":
		return controller.items.GetItemList()
	case "Item Attribute":
		return controller.items.GetItemAttributeList()
	case "Item Category":
		return controller.items.GetItemCategoryList()
	case "Item Fling Effect":
		return controller.items.GetItemFlingEffectList()
	case "Item Pocket":
		return controller.items.GetItemPocketList()
	}

	return structs.Resource{}
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

	return nil, nil
}
