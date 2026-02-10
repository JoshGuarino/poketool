package contests

import (
	"github.com/JoshGuarino/PokeGo/pkg/models"
	contestsGroup "github.com/JoshGuarino/PokeGo/pkg/resources/contests"
)

type IContests interface {
	GetContestType(nameOrId string) (*models.ContestType, error)
	GetContestTypeList(limit int, offset int) (*models.NamedResourceList, error)
	GetContestEffect(id string) (*models.ContestEffect, error)
	GetContestEffectList(limit int, offset int) (*models.ResourceList, error)
	GetSuperContestEffect(id string) (*models.SuperContestEffect, error)
	GetSuperContestEffectList(limit int, offset int) (*models.ResourceList, error)
}

type Contests struct {
	contestsGroup contestsGroup.Contests
}

func NewContests() Contests {
	return Contests{
		contestsGroup: contestsGroup.NewContestsGroup(),
	}
}

func (c Contests) GetContestType(nameOrId string) (*models.ContestType, error) {
	contestType, err := c.contestsGroup.GetContestType(nameOrId)
	if err != nil {
		return nil, err
	}
	return contestType, nil
}

func (c Contests) GetContestTypeList(limit int, offset int) (*models.NamedResourceList, error) {
	contestTypeList, err := c.contestsGroup.GetContestTypeList(limit, offset)
	if err != nil {
		return nil, err
	}
	return contestTypeList, nil
}

func (c Contests) GetContestEffect(id string) (*models.ContestEffect, error) {
	contestEffect, err := c.contestsGroup.GetContestEffect(id)
	if err != nil {
		return nil, err
	}
	return contestEffect, nil
}

func (c Contests) GetContestEffectList(limit int, offset int) (*models.ResourceList, error) {
	contestEffectList, err := c.contestsGroup.GetContestEffectList(limit, offset)
	if err != nil {
		return nil, err
	}
	return contestEffectList, nil
}

func (c Contests) GetSuperContestEffect(id string) (*models.SuperContestEffect, error) {
	superContestEffect, err := c.contestsGroup.GetSuperContestEffect(id)
	if err != nil {
		return nil, err
	}
	return superContestEffect, nil
}

func (c Contests) GetSuperContestEffectList(limit int, offset int) (*models.ResourceList, error) {
	superContestEffectList, err := c.contestsGroup.GetSuperContestEffectList(limit, offset)
	if err != nil {
		return nil, err
	}
	return superContestEffectList, nil
}
