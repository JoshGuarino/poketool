package berries

import (
	"fmt"

	berriesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/berries"
)

type BerriesController struct {
	berries berriesGroup.Berries
}

func NewController() BerriesController {
	return BerriesController{
		berries: berriesGroup.NewBerriesGroup(),
	}
}

func (c BerriesController) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Berries":
		return c.berries.GetBerryList(limit, offset)
	case "Berry Firmnesses":
		return c.berries.GetBerryFirmnessList(limit, offset)
	case "Berry Flavors":
		return c.berries.GetBerryFlavorList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c BerriesController) GetSpecific(result string, search string) (any, error) {
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
