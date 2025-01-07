package items

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
  switch result {
	case "Item":
    return GetItemList()
	case "Item Attribute":
		return GetItemAttributeList()
	case "Item Category":
		return GetItemCategoryList()
	case "Item Fling Effect":
    return GetItemFlingEffectList()
  case "Item Pocket":
    return GetItemPocketList()
  }

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
  switch result {
	case "Item":
    return GetItem(search)
	case "Item Attribute":
		return GetItemAttribute(search)
	case "Item Category":
		return GetItemCategory(search)
	case "Item Fling Effect":
    return GetItemFlingEffect(search)
  case "Item Pocket":
    return GetItemPocket(search)
  }

	return "", nil
} 
