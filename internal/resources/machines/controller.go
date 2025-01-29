package machines

import "github.com/mtslzr/pokeapi-go/structs"

func (controller Controller) GetList(result string) structs.Resource {
	switch result {
	case "Machine":
		return controller.machines.GetMachineList()
	}

	return structs.Resource{}
}

func (controller Controller) GetSpecific(result string, search string) (interface{}, error) {
	switch result {
	case "Machine":
		return controller.machines.GetMachine(search)
	}

	return nil, nil
}
