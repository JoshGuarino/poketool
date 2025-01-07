package locations

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetLocation(nameOrId string) (structs.Location, error) {
	location, err := pokeapi.Location(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "location", nameOrId)
		return structs.Location{}, err
	}
	return location, nil
}

func GetLocationList() structs.Resource {
	locationList := internal.GetResourceList(locationEndpoint)
	return locationList
}

func GetLocationArea(nameOrId string) (structs.LocationArea, error) {
	locationArea, err := pokeapi.LocationArea(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "location area", nameOrId)
		return structs.LocationArea{}, err
	}
	return locationArea, nil
}

func GetLocationAreaList() structs.Resource {
	locationAreaList := internal.GetResourceList(locationAreaEndpoint)
	return locationAreaList
}

func GetPalParkArea(nameOrId string) (structs.PalParkArea, error) {
	palParkArea, err := pokeapi.PalParkArea(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "pal park area", nameOrId)
		return structs.PalParkArea{}, err
	}
	return palParkArea, nil
}

func GetPalParkAreaList() structs.Resource {
	palParkAreaList := internal.GetResourceList(palParkAreaEndpoint)
	return palParkAreaList
}

func GetRegion(nameOrId string) (structs.Region, error) {
	region, err := pokeapi.Region(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "region", nameOrId)
		return structs.Region{}, err
	}
	return region, err
}

func GetRegionList() structs.Resource {
	regionList := internal.GetResourceList(regionEndpoint)
	return regionList
}
