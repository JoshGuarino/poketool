package items

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (items Items) GetItem(nameOrId string) (structs.Item, error) {
	item, err := pokeapi.Item(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "item", nameOrId)
		return structs.Item{}, err
	}
	return item, nil
}

func (items Items) GetItemList() structs.Resource {
	itemList := internal.GetResourceList(itemEndpoint)
	return itemList
}

func (items Items) GetItemAttribute(nameOrId string) (structs.ItemAttribute, error) {
	itemAttribute, err := pokeapi.ItemAttribute(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "item attribute", nameOrId)
		return structs.ItemAttribute{}, err
	}
	return itemAttribute, nil
}

func (items Items) GetItemAttributeList() structs.Resource {
	itemAttributeList := internal.GetResourceList(itemAttributeEndpoint)
	return itemAttributeList
}

func (items Items) GetItemCategory(nameOrId string) (structs.ItemCategory, error) {
	itemCategory, err := pokeapi.ItemCategory(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "item category", nameOrId)
		return structs.ItemCategory{}, err
	}
	return itemCategory, nil
}

func (items Items) GetItemCategoryList() structs.Resource {
	itemCategoryList := internal.GetResourceList(itemCategoryEndpoint)
	return itemCategoryList
}

func (items Items) GetItemFlingEffect(nameOrId string) (structs.ItemFlingEffect, error) {
	itemFlingEffect, err := pokeapi.ItemFlingEffect(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "item fling effect", nameOrId)
		return structs.ItemFlingEffect{}, err
	}
	return itemFlingEffect, nil
}

func (items Items) GetItemFlingEffectList() structs.Resource {
	itemFlingEffectList := internal.GetResourceList(itemFlingEffectEndpoint)
	return itemFlingEffectList
}

func (items Items) GetItemPocket(nameOrId string) (structs.ItemPocket, error) {
	itemPocket, err := pokeapi.ItemPocket(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "item pocket", nameOrId)
		return structs.ItemPocket{}, err
	}
	return itemPocket, nil
}

func (items Items) GetItemPocketList() structs.Resource {
	itemPocketList := internal.GetResourceList(itemPocketEndpoint)
	return itemPocketList
}
