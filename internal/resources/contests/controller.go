package contests

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Contest Type":
		return GetContestTypeList()
	case "Contest Effect":
		return GetContestEffectList()
	case "Super Contest Effect":
		return GetSuperContestEffectList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Contest Type":
		return GetContestType(search)
	case "Contest Effect":
		return GetContestEffect(search)
	case "Super Contest Effect":
		return GetSuperContestEffect(search)
	}

	return "", nil
}
