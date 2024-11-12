package contests

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetContestType(nameOrId string) (structs.ContestType, error) {
	contestType, err := pokeapi.ContestType(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "contest type", nameOrId)
		return structs.ContestType{}, err
	}
	return contestType, nil
}

func GetContestTypeList() structs.Resource {
	contestTypeList := internal.GetResourceList(contestTypeEndpoint)
	return contestTypeList
}

func GetContestEffect(id string) (structs.ContestEffect, error) {
	contestEffect, err := pokeapi.ContestEffect(id)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "contest effect", id)
		return structs.ContestEffect{}, err
	}
	return contestEffect, nil
}

func GetContestEffectList() structs.Resource {
	contestEffectList := internal.GetResourceList(contestEffectEndpoint)
	return contestEffectList
}

func GetSuperContestEffect(id string) (structs.SuperContestEffect, error) {
	superContestEffect, err := pokeapi.SuperContestEffect(id)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "super contest effect", id)
		return structs.SuperContestEffect{}, err
	}
	return superContestEffect, nil
}

func GetSuperContestEffectList() structs.Resource {
	superContestEffectList := internal.GetResourceList(superContestEffectEndpoint)
	return superContestEffectList
}
