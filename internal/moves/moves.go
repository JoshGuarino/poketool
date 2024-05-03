package moves

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetMove(nameOrId string) structs.Move {
	move, err := pokeapi.Move(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return move
}

func GetMoveList() structs.Resource {
	moveList := internal.GetResourceList(moveEndpoint)
	return moveList
}

func GetMoveAilment(nameOrId string) structs.MoveAilment {
	moveAilment, err := pokeapi.MoveAilment(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return moveAilment
}

func GetMoveAilmentList() structs.Resource {
	moveAilmentList := internal.GetResourceList(moveAilmentEndpoint)
	return moveAilmentList
}

func GetMoveBattleStyle(nameOrId string) structs.MoveBattleStyle {
	moveBattleStyle, err := pokeapi.MoveBattleStyle(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return moveBattleStyle
}

func GetMoveBattleStyleList() structs.Resource {
	moveBattleStyleList := internal.GetResourceList(moveBattleStyleEndpoint)
	return moveBattleStyleList
}

func GetMoveCategory(nameOrId string) structs.MoveCategory {
	moveCategory, err := pokeapi.MoveCategory(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return moveCategory
}

func GetMoveCategoryList() structs.Resource {
	moveCategoryList := internal.GetResourceList(moveCategoryEndpoint)
	return moveCategoryList
}

func GetMoveDamageClass(nameOrId string) structs.MoveDamageClass {
	moveDamageClass, err := pokeapi.MoveDamageClass(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return moveDamageClass
}

func GetMoveDamageClassList() structs.Resource {
	moveDamageClassList := internal.GetResourceList(moveDamageClassEndpoint)
	return moveDamageClassList
}

func GetMoveLearnMethod(nameOrId string) structs.MoveLearnMethod {
	moveLearnMethod, err := pokeapi.MoveLearnMethod(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return moveLearnMethod
}

func GetMoveLearnMethodList() structs.Resource {
	moveLearnMethodList := internal.GetResourceList(moveLearnMethodEndpoint)
	return moveLearnMethodList
}

func GetMoveTarget(nameOrId string) structs.MoveTarget {
	moveTarget, err := pokeapi.MoveTarget(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return moveTarget
}

func GetMoveTargetList() structs.Resource {
	moveTargetList := internal.GetResourceList(moveTargetEndpoint)
	return moveTargetList
}
