package machines

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetMachine(nameOrId string) structs.Machine {
	Machine, err := pokeapi.Machine(nameOrId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return Machine
}

func GetMachineList() structs.Resource {
	machineList := internal.GetResourceList(machineEndpoint)
	return machineList
}
