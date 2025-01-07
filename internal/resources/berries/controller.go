package berries

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Berries":
		return GetBerryList()
	case "Berry Firmnesses":
		return GetBerryFirmnessList()
	case "Berry Flavors":
		return GetBerryFlavorList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Berries":
		return GetBerry(search)
	case "Berry Firmnesses":
		return GetBerryFirmness(search)
	case "Berry Flavors":
		return GetBerryFlavor(search)
	}

	return "", nil
}
