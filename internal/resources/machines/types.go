package machines

import "github.com/mtslzr/pokeapi-go/structs"

type Controller struct {
	machines Machines
}

type IMachines interface {
	GetMachine(id string) (structs.Machine, error)
	GetMachineList() structs.Resource
}

type Machines struct{}
