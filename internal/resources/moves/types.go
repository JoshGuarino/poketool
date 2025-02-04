package moves

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	moves Moves
}

type IMoves interface {
	GetMove(nameOrId string) (structs.Move, error)
	GetMoveList() structs.Resource
	GetMoveAilment(nameOrId string) (structs.MoveAilment, error)
	GetMoveAilmentList() structs.Resource
	GetMoveBattleStyle(nameOrId string) (structs.MoveBattleStyle, error)
	GetMoveBattleStyleList() structs.Resource
	GetMoveCategory(nameOrId string) (structs.MoveCategory, error)
	GetMoveCategoryList() structs.Resource
	GetMoveDamageClass(nameOrId string) (structs.MoveDamageClass, error)
	GetMoveDamageClassList() structs.Resource
	GetMoveLearnMethod(nameOrId string) (structs.MoveLearnMethod, error)
	GetMoveLearnMethodList() structs.Resource
	GetMoveTarget(nameOrId string) (structs.MoveTarget, error)
	GetMoveTargetList() structs.Resource
}

type Moves struct{}
