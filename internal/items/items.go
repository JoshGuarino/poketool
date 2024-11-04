package items

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetItem(nameOrId string) (structs.Item, error) {
	item, err := pokeapi.Item(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "generation", nameOrId)
		return structs.Item{}, err
	}
	return item, nil
}

func GetItemList() structs.Resource {
	itemList := internal.GetResourceList(itemEndpoint)
	return itemList
}

func GetItemAttribute(nameOrId string) (structs.ItemAttribute, error) {
	itemAttribute, err := pokeapi.ItemAttribute(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "generation", nameOrId)
		return structs.ItemAttribute{}, err
	}
	return itemAttribute, nil
}

func GetItemAttributeList() structs.Resource {
	itemAttributeList := internal.GetResourceList(itemAttributeEndpoint)
	return itemAttributeList
}

func GetItemCategory(nameOrId string) (structs.ItemCategory, error) {
	itemCategory, err := pokeapi.ItemCategory(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "generation", nameOrId)
		return structs.ItemCategory{}, err
	}
	return itemCategory, nil
}

func GetItemCategoryList() structs.Resource {
	itemCategoryList := internal.GetResourceList(itemCategoryEndpoint)
	return itemCategoryList
}

func GetItemFlingEffect(nameOrId string) (structs.ItemFlingEffect, error) {
	itemFlingEffect, err := pokeapi.ItemFlingEffect(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "generation", nameOrId)
		return structs.ItemFlingEffect{}, err
	}
	return itemFlingEffect, nil
}

func GetItemFlingEffectList() structs.Resource {
	itemFlingEffectList := internal.GetResourceList(itemFlingEffectEndpoint)
	return itemFlingEffectList
}

func GetItemPocket(nameOrId string) (structs.ItemPocket, error) {
	itemPocket, err := pokeapi.ItemPocket(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "generation", nameOrId)
		return structs.ItemPocket{}, err
	}
	return itemPocket, nil
}

func GetItemPocketList() structs.Resource {
	itemPocketList := internal.GetResourceList(itemPocketEndpoint)
	return itemPocketList
}
