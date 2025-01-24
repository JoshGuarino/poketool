package berries

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (berries Berries) GetBerry(nameOrId string) (structs.Berry, error) {
	berry, err := pokeapi.Berry(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "berry", nameOrId)
		return structs.Berry{}, err
	}
	return berry, nil
}

func (berries Berries) GetBerryList() structs.Resource {
	berryList := internal.GetResourceList(BerryEndpoint)
	return berryList
}

func (berries Berries) GetBerryFirmness(nameOrId string) (structs.BerryFirmness, error) {
	BerryFirmness, err := pokeapi.BerryFirmness(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "berry firmness", nameOrId)
		return structs.BerryFirmness{}, err
	}
	return BerryFirmness, nil
}

func (berries Berries) GetBerryFirmnessList() structs.Resource {
	berryFirmnessList := internal.GetResourceList(BerryFirmnessEndpoint)
	return berryFirmnessList
}

func (berries Berries) GetBerryFlavor(nameOrId string) (structs.BerryFlavor, error) {
	berryFlavor, err := pokeapi.BerryFlavor(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "berry flavor", nameOrId)
		return structs.BerryFlavor{}, err
	}
	return berryFlavor, nil
}

func (berries Berries) GetBerryFlavorList() structs.Resource {
	berryFlavorList := internal.GetResourceList(BerryFlavorEndpoint)
	return berryFlavorList
}
