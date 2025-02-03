package locations

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Location":
		return controller.locations.GetLocationList()
	case "Location Area":
		return controller.locations.GetLocationAreaList()
	case "Pal Park Area":
		return controller.locations.GetPalParkAreaList()
	case "Region":
		return controller.locations.GetRegionList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Location":
		return controller.locations.GetLocation(search)
	case "Location Area":
		return controller.locations.GetLocationArea(search)
	case "Pal Park Area":
		return controller.locations.GetPalParkArea(search)
	case "Region":
		return controller.locations.GetRegion(search)
	}

	return nil, nil
}
