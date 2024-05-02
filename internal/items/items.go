package items

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetItem(nameOrId string) structs.Item {
	item, err := pokeapi.Item(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return item
}

func GetItemList() structs.Resource {
	itemList := internal.GetResourceList(itemEndpoint)
	return itemList
}

func GetItemAttribute(nameOrId string) structs.ItemAttribute {
	itemAttribute, err := pokeapi.ItemAttribute(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return itemAttribute
}

func GetItemAttributeList() structs.Resource {
	itemAttributeList := internal.GetResourceList(itemAttributeEndpoint)
	return itemAttributeList
}

func GetItemCategory(nameOrId string) structs.ItemCategory {
	itemCategory, err := pokeapi.ItemCategory(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return itemCategory
}

func GetItemCategoryList() structs.Resource {
	itemCategoryList := internal.GetResourceList(itemCategoryEndpoint)
	return itemCategoryList
}

func GetItemFlingEffect(nameOrId string) structs.ItemFlingEffect {
	itemFlingEffect, err := pokeapi.ItemFlingEffect(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return itemFlingEffect
}

func GetItemFlingEffectList() structs.Resource {
	itemFlingEffectList := internal.GetResourceList(itemFlingEffectEndpoint)
	return itemFlingEffectList
}

func GetItemPocket(nameOrId string) structs.ItemPocket {
	itemPocket, err := pokeapi.ItemPocket(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return itemPocket
}

func GetItemPocketList() structs.Resource {
	itemPocketList := internal.GetResourceList(itemPocketEndpoint)
	return itemPocketList
}
