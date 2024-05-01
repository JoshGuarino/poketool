package berries

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetBerry(nameOrId string) structs.Berry {
	berry, err := pokeapi.Berry(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return berry
}

func GetBerryList() structs.Resource {
	berryList := internal.GetResourceList(BerryEndpoint)
	return berryList
}

func GetBerryFirmness(nameOrId string) structs.BerryFirmness {
	BerryFirmness, err := pokeapi.BerryFirmness(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return BerryFirmness
}

func GetBerryFirmnessList() structs.Resource {
	berryFirmnessList := internal.GetResourceList(BerryFirmnessEndpoint)
	return berryFirmnessList
}

func GetBerryFlavor(nameOrId string) structs.BerryFlavor {
	berryFlavor, err := pokeapi.BerryFlavor(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return berryFlavor
}

func GetBerryFlavorList() structs.Resource {
	berryFlavorList := internal.GetResourceList(BerryFlavorEndpoint)
	return berryFlavorList
}
