package contests

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetContestType(nameOrId string) structs.ContestType {
	contestType, err := pokeapi.ContestType(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return contestType
}

func GetContetTypeList() structs.Resource {
	contestTypeList := internal.GetResourceList(contestTypeEndpoint)
	return contestTypeList
}

func GetContestEffect(nameOrId string) structs.ContestEffect {
	contestEffect, err := pokeapi.ContestEffect(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return contestEffect
}

func GetContestEffectList() structs.Resource {
	contestEffectList := internal.GetResourceList(contestEffectEndpoint)
	return contestEffectList
}

func GetSuperContestEffect(nameOrId string) structs.SuperContestEffect {
	superContestEffect, err := pokeapi.SuperContestEffect(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return superContestEffect
}

func GetSuperContestEffectList() structs.Resource {
	superContestEffectList := internal.GetResourceList(superContestEffectEndpoint)
	return superContestEffectList
}
