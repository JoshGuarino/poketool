package berries

import (
	"github.com/JoshGuarino/PokeGo/pkg/models"
	berriesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/berries"
)

type IBerries interface {
	GetBerry(nameOrId string) (*models.Berry, error)
	GetBerryList(limit int, offset int) (*models.NamedResourceList, error)
	GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error)
	GetBerryFirmnessList(limit, offset int) (*models.NamedResourceList, error)
	GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error)
	GetBerryFlavorList(limit, offset int) (*models.NamedResourceList, error)
}

type Berries struct {
	berriesGroup berriesGroup.Berries
}

func NewBerries() Berries {
	return Berries{
		berriesGroup: berriesGroup.NewBerriesGroup(),
	}
}

func (b Berries) GetBerry(nameOrId string) (*models.Berry, error) {
	berry, err := b.berriesGroup.GetBerry(nameOrId)
	if err != nil {
		return nil, err
	}
	return berry, nil
}

func (b Berries) GetBerryList(limit int, offset int) (*models.NamedResourceList, error) {
	berryList, err := b.berriesGroup.GetBerryList(limit, offset)
	if err != nil {
		return nil, err
	}
	return berryList, nil
}

func (b Berries) GetBerryFirmness(nameOrId string) (*models.BerryFirmness, error) {
	berryFirmness, err := b.berriesGroup.GetBerryFirmness(nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFirmness, nil
}

func (b Berries) GetBerryFirmnessList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFirmnessList, err := b.berriesGroup.GetBerryFirmnessList(limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFirmnessList, nil
}

func (b Berries) GetBerryFlavor(nameOrId string) (*models.BerryFlavor, error) {
	berryFlavor, err := b.berriesGroup.GetBerryFlavor(nameOrId)
	if err != nil {
		return nil, err
	}
	return berryFlavor, nil
}

func (b Berries) GetBerryFlavorList(limit int, offset int) (*models.NamedResourceList, error) {
	berryFlavorList, err := b.berriesGroup.GetBerryFlavorList(limit, offset)
	if err != nil {
		return nil, err
	}
	return berryFlavorList, nil
}
