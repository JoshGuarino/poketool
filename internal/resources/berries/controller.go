package berries

import (
	"fmt"

	berriesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/berries"
)

type IController interface {
	GetList(result string) (any, error)
	GetSpecific(result string, search string) (any, error)
}

type Controller struct {
	berries berriesGroup.Berries
}

func NewController() Controller {
	return Controller{
		berries: berriesGroup.NewBerriesGroup(),
	}
}

func (c Controller) GetList(result string) (any, error) {
	switch result {
	case "Berries":
		return c.berries.GetBerryList(20, 0)
	case "Berry Firmnesses":
		return c.berries.GetBerryFirmnessList(20, 0)
	case "Berry Flavors":
		return c.berries.GetBerryFlavorList(20, 0)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c Controller) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Berries":
		return c.berries.GetBerry(search)
	case "Berry Firmnesses":
		return c.berries.GetBerryFirmness(search)
	case "Berry Flavors":
		return c.berries.GetBerryFlavor(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
