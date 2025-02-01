package items

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	items Items
}

type IItems interface {
	GetItem(nameOrId string) (structs.Item, error)
	GetItemList() structs.Resource
	GetItemAttribute(nameOrId string) (structs.ItemAttribute, error)
	GetItemAttributeList() structs.Resource
	GetItemCategory(nameOrId string) (structs.ItemCategory, error)
	GetItemCategoryList() structs.Resource
	GetItemFlingEffect(nameOrId string) (structs.ItemFlingEffect, error)
	GetItemFlingEffectList() structs.Resource
	GetItemPocket(nameOrId string) (structs.ItemPocket, error)
	GetItemPocketList() structs.Resource
}

type Items struct{}
