package berries

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetBerry(nameOrId string) (structs.Berry, error) {
	berry, err := pokeapi.Berry(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "berry", nameOrId)
		return structs.Berry{}, err
	}
	return berry, nil
}

func GetBerryList() structs.Resource {
	berryList := internal.GetResourceList(BerryEndpoint)
	return berryList
}

func GetBerryFirmness(nameOrId string) (structs.BerryFirmness, error) {
	BerryFirmness, err := pokeapi.BerryFirmness(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "berry firmness", nameOrId)
		return structs.BerryFirmness{}, nil
	}
	return BerryFirmness, nil
}

func GetBerryFirmnessList() structs.Resource {
	berryFirmnessList := internal.GetResourceList(BerryFirmnessEndpoint)
	return berryFirmnessList
}

func GetBerryFlavor(nameOrId string) (structs.BerryFlavor, error) {
	berryFlavor, err := pokeapi.BerryFlavor(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "berry flavor", nameOrId)
		return structs.BerryFlavor{}, err
	}
	return berryFlavor, nil
}

func GetBerryFlavorList() structs.Resource {
	berryFlavorList := internal.GetResourceList(BerryFlavorEndpoint)
	return berryFlavorList
}
