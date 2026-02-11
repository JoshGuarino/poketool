package locations

import (
	"fmt"

	locationsGroup "github.com/JoshGuarino/PokeGo/pkg/resources/locations"
)

type LocationsController struct {
	locations locationsGroup.Locations
}

func NewController() LocationsController {
	return LocationsController{
		locations: locationsGroup.NewLocationsGroup(),
	}
}

func (c LocationsController) GetList(result string, limit int, offset int) (interface{}, error) {
	switch result {
	case "Location":
		return c.locations.GetLocationList(limit, offset)
	case "Location Area":
		return c.locations.GetLocationAreaList(limit, offset)
	case "Pal Park Area":
		return c.locations.GetPalParkAreaList(limit, offset)
	case "Region":
		return c.locations.GetRegionList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c LocationsController) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Location":
		return c.locations.GetLocation(search)
	case "Location Area":
		return c.locations.GetLocationArea(search)
	case "Pal Park Area":
		return c.locations.GetPalParkArea(search)
	case "Region":
		return c.locations.GetRegion(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
