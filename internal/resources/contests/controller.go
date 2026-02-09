package contests

import (
	"fmt"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	contestsGroup "github.com/JoshGuarino/PokeGo/pkg/resources/contests"
)

type IController interface {
	GetList(result string) (*models.NamedResourceList, error)
	GetSpecific(result string, search string) (any, error)
}

type Controller struct {
	contests contestsGroup.Contests
}

func NewController() Controller {
	return Controller{
		contests: contestsGroup.NewContestsGroup(),
	}
}

func (c Controller) GetList(result string) (any, error) {
	switch result {
	case "Contest Type":
		return c.contests.GetContestTypeList(20, 0)
	case "Contest Effect":
		return c.contests.GetContestEffectList(20, 0)
	case "Super Contest Effect":
		return c.contests.GetSuperContestEffectList(20, 0)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c Controller) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Contest Type":
		return c.contests.GetContestType(search)
	case "Contest Effect":
		return c.contests.GetContestEffect(search)
	case "Super Contest Effect":
		return c.contests.GetSuperContestEffect(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
