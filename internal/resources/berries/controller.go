package berries

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

func (c Controller) GetList(result string) structs.Resource {
	switch result {
	case "Berries":
		return c.berries.GetBerryList()
	case "Berry Firmnesses":
		return c.berries.GetBerryFirmnessList()
	case "Berry Flavors":
		return c.berries.GetBerryFlavorList()
	}

	return structs.Resource{}
}

func (c Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Berries":
		return c.berries.GetBerry(search)
	case "Berry Firmnesses":
		return c.berries.GetBerryFirmness(search)
	case "Berry Flavors":
		return c.berries.GetBerryFlavor(search)
	}

	return nil, nil
}
