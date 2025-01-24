package contests

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Contest Type":
		return controller.contests.GetContestTypeList()
	case "Contest Effect":
		return controller.contests.GetContestEffectList()
	case "Super Contest Effect":
		return controller.contests.GetSuperContestEffectList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Contest Type":
		return controller.contests.GetContestType(search)
	case "Contest Effect":
		return controller.contests.GetContestEffect(search)
	case "Super Contest Effect":
		return controller.contests.GetSuperContestEffect(search)
	}

	return nil, nil
}
