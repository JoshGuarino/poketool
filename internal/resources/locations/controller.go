package locations

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Location":
		return GetLocationList()
	case "Location Area":
		return GetLocationAreaList()
	case "Pal Park Area":
		return GetPalParkAreaList()
	case "Region":
		return GetRegionList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Location":
		return GetLocation(search)
	case "Location Area":
		return GetLocationArea(search)
	case "Pal Park Area":
		return GetPalParkArea(search)
	case "Region":
		return GetRegion(search)
	}

	return "", nil
}
