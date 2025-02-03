package locations

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	locations Locations
}

type ILocations interface {
	GetLocation(nameOrId string) (structs.Location, error)
	GetLocationList() structs.Resource
	GetLocationArea(nameOrId string) (structs.LocationArea, error)
	GetLocationAreaList() structs.Resource
	GetPalParkArea(nameOrId string) (structs.PalParkArea, error)
	GetPalParkAreaList() structs.Resource
	GetRegion(nameOrId string) (structs.Region, error)
	GetRegionList() structs.Resource
}

type Locations struct{}
