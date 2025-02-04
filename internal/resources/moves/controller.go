package moves

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Move":
		return controller.moves.GetMoveList()
	case "Move Ailment":
		return controller.moves.GetMoveAilmentList()
	case "Move Battle Style":
		return controller.moves.GetMoveBattleStyleList()
	case "Move Category":
		return controller.moves.GetMoveCategoryList()
	case "Move Damage Class":
		return controller.moves.GetMoveDamageClassList()
	case "Move Learn Method":
		return controller.moves.GetMoveLearnMethodList()
	case "Move Target":
		return controller.moves.GetMoveTargetList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Move":
		return controller.moves.GetMove(search)
	case "Move Ailment":
		return controller.moves.GetMoveAilment(search)
	case "Move Battle Style":
		return controller.moves.GetMoveBattleStyle(search)
	case "Move Category":
		return controller.moves.GetMoveCategory(search)
	case "Move Damage Class":
		return controller.moves.GetMoveDamageClass(search)
	case "Move Learn Method":
		return controller.moves.GetMoveLearnMethod(search)
	case "Move Target":
		return controller.moves.GetMoveTarget(search)
	}

	return nil, nil
}
