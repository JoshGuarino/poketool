package contests

import (
	"fmt"

	contestsGroup "github.com/JoshGuarino/PokeGo/pkg/resources/contests"
)

type ContestsController struct {
	contests contestsGroup.Contests
}

func NewController() ContestsController {
	return ContestsController{
		contests: contestsGroup.NewContestsGroup(),
	}
}

func (c ContestsController) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Contest Type":
		return c.contests.GetContestTypeList(limit, offset)
	case "Contest Effect":
		return c.contests.GetContestEffectList(limit, offset)
	case "Super Contest Effect":
		return c.contests.GetSuperContestEffectList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c ContestsController) GetSpecific(result string, search string) (any, error) {
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
