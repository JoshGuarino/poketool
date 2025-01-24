package berries

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	berries Berries
}

type IBerries interface {
	GetBerry(nameOrId string) (structs.Berry, error)
	GetBerryList() structs.Resource
	GetBerryFirmness(nameOrId string) (structs.BerryFirmness, error)
	GetBerryFirmnessList() structs.Resource
	GetBerryFlavor(nameOrId string) (structs.BerryFlavor, error)
	GetBerryFlavorList() structs.Resource
}

type Berries struct{}
