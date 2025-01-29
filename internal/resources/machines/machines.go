package machines

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (machines Machines) GetMachine(id string) (structs.Machine, error) {
	Machine, err := pokeapi.Machine(id)
	if err != nil {
		fmt.Printf(internal.ErrorStringGetByNameOrId, "machine", id)
		return structs.Machine{}, err
	}
	return Machine, nil
}

func (machines Machines) GetMachineList() structs.Resource {
	machineList := internal.GetResourceList(machineEndpoint)
	return machineList
}
