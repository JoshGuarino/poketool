package contests

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	contests Contests
}

type IContests interface {
	GetContestType(nameOrId string) (structs.ContestType, error)
	GetContestTypeList() structs.Resource
	GetContestEffect(id string) (structs.ContestEffect, error)
	GetContestEffectList() structs.Resource
	GetSuperContestEffect(id string) (structs.SuperContestEffect, error)
	GetSuperContestEffectList() structs.Resource
}

type Contests struct{}
