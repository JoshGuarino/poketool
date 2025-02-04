package moves

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (moves Moves) GetMove(nameOrId string) (structs.Move, error) {
	move, err := pokeapi.Move(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move", nameOrId)
		return structs.Move{}, err
	}
	return move, err
}

func (moves Moves) GetMoveList() structs.Resource {
	moveList := internal.GetResourceList(moveEndpoint)
	return moveList
}

func (moves Moves) GetMoveAilment(nameOrId string) (structs.MoveAilment, error) {
	moveAilment, err := pokeapi.MoveAilment(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move ailment", nameOrId)
		return structs.MoveAilment{}, err
	}
	return moveAilment, nil
}

func (moves Moves) GetMoveAilmentList() structs.Resource {
	moveAilmentList := internal.GetResourceList(moveAilmentEndpoint)
	return moveAilmentList
}

func (moves Moves) GetMoveBattleStyle(nameOrId string) (structs.MoveBattleStyle, error) {
	moveBattleStyle, err := pokeapi.MoveBattleStyle(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move battle style", nameOrId)
		return structs.MoveBattleStyle{}, err
	}
	return moveBattleStyle, nil
}

func (moves Moves) GetMoveBattleStyleList() structs.Resource {
	moveBattleStyleList := internal.GetResourceList(moveBattleStyleEndpoint)
	return moveBattleStyleList
}

func (moves Moves) GetMoveCategory(nameOrId string) (structs.MoveCategory, error) {
	moveCategory, err := pokeapi.MoveCategory(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move category", nameOrId)
		return structs.MoveCategory{}, err
	}
	return moveCategory, nil
}

func (moves Moves) GetMoveCategoryList() structs.Resource {
	moveCategoryList := internal.GetResourceList(moveCategoryEndpoint)
	return moveCategoryList
}

func (moves Moves) GetMoveDamageClass(nameOrId string) (structs.MoveDamageClass, error) {
	moveDamageClass, err := pokeapi.MoveDamageClass(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move damage class", nameOrId)
		return structs.MoveDamageClass{}, err
	}
	return moveDamageClass, nil
}

func (moves Moves) GetMoveDamageClassList() structs.Resource {
	moveDamageClassList := internal.GetResourceList(moveDamageClassEndpoint)
	return moveDamageClassList
}

func (moves Moves) GetMoveLearnMethod(nameOrId string) (structs.MoveLearnMethod, error) {
	moveLearnMethod, err := pokeapi.MoveLearnMethod(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move learn method", nameOrId)
		return structs.MoveLearnMethod{}, err
	}
	return moveLearnMethod, nil
}

func (moves Moves) GetMoveLearnMethodList() structs.Resource {
	moveLearnMethodList := internal.GetResourceList(moveLearnMethodEndpoint)
	return moveLearnMethodList
}

func (moves Moves) GetMoveTarget(nameOrId string) (structs.MoveTarget, error) {
	moveTarget, err := pokeapi.MoveTarget(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "move target", nameOrId)
		return structs.MoveTarget{}, err
	}
	return moveTarget, nil
}

func (moves Moves) GetMoveTargetList() structs.Resource {
	moveTargetList := internal.GetResourceList(moveTargetEndpoint)
	return moveTargetList
}
