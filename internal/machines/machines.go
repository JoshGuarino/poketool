package machines

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetMachine(nameOrId string) (structs.Machine, error) {
	Machine, err := pokeapi.Machine(nameOrId)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "machine", nameOrId)
		return structs.Machine{}, err
	}
	return Machine, nil
}

func GetMachineList() structs.Resource {
	machineList := internal.GetResourceList(machineEndpoint)
	return machineList
}
