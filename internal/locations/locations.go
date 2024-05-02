package locations

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetLocation(nameOrId string) structs.Location {
	location, err := pokeapi.Location(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return location
}

func GetLocationList() structs.Resource {
	locationList := internal.GetResourceList(locationEndpoint)
	return locationList
}

func GetLocationArea(nameOrId string) structs.LocationArea {
	locationArea, err := pokeapi.LocationArea(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return locationArea
}

func GetLocationAreaList() structs.Resource {
	locationAreaList := internal.GetResourceList(locationAreaEndpoint)
	return locationAreaList
}

func GetPalParkArea(nameOrId string) structs.PalParkArea {
	palParkArea, err := pokeapi.PalParkArea(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return palParkArea
}

func GetPalParkAreaList() structs.Resource {
	palParkAreaList := internal.GetResourceList(palParkAreaEndpoint)
	return palParkAreaList
}

func GetRegion(nameOrId string) structs.Region {
	region, err := pokeapi.Region(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return region
}

func GetRegionList() structs.Resource {
	regionList := internal.GetResourceList(regionEndpoint)
	return regionList
}
