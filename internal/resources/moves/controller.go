package moves

import (
	"fmt"

	movesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/moves"
)

type Controller struct {
	moves movesGroup.Moves
}

func NewController() Controller {
	return Controller{
		moves: movesGroup.NewMovesGroup(),
	}
}

func (c Controller) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Move":
		return c.moves.GetMoveList(limit, offset)
	case "Move Ailment":
		return c.moves.GetMoveAilmentList(limit, offset)
	case "Move Battle Style":
		return c.moves.GetMoveBattleStyleList(limit, offset)
	case "Move Category":
		return c.moves.GetMoveCategoryList(limit, offset)
	case "Move Damage Class":
		return c.moves.GetMoveDamageClassList(limit, offset)
	case "Move Learn Method":
		return c.moves.GetMoveLearnMethodList(limit, offset)
	case "Move Target":
		return c.moves.GetMoveTargetList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
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

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
