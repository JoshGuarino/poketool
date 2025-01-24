package berries

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Berries":
		return controller.berries.GetBerryList()
	case "Berry Firmnesses":
		return controller.berries.GetBerryFirmnessList()
	case "Berry Flavors":
		return controller.berries.GetBerryFlavorList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Berries":
		return controller.berries.GetBerry(search)
	case "Berry Firmnesses":
		return controller.berries.GetBerryFirmness(search)
	case "Berry Flavors":
		return controller.berries.GetBerryFlavor(search)
	}

	return nil, nil
}
