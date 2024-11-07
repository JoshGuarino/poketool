package moves

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Move":
		return GetMoveList()
	case "Move Ailment":
		return GetMoveAilmentList()
	case "Move Battle Style":
		return GetMoveBattleStyleList()
	case "Move Category":
		return GetMoveCategoryList()
	case "Move Damage Class":
		return GetMoveDamageClassList()
	case "Move Learn Method":
		return GetMoveLearnMethodList()
	case "Move Target":
		return GetMoveTargetList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Move":
		return GetMove(search)
	case "Move Ailment":
		return GetMoveAilment(search)
	case "Move Battle Style":
		return GetMoveBattleStyle(search)
	case "Move Category":
		return GetMoveCategory(search)
	case "Move Damage Class":
		return GetMoveDamageClass(search)
	case "Move Learn Method":
		return GetMoveLearnMethod(search)
	case "Move Target":
		return GetMoveTarget(search)
	}

	return "", nil
}
