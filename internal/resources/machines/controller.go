package machines

import "github.com/mtslzr/pokeapi-go/structs"

func GetList(result string) structs.Resource {
	switch result {
	case "Machine":
		return GetMachineList()
	}

	return structs.Resource{}
}

func GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Machine":
		return GetMachine(search)
	}

	return "", nil
}
