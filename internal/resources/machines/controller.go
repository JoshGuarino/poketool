package machines

import (
	"fmt"

	machinesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/machines"
)

type Controller struct {
	machines machinesGroup.Machines
}

func NewController() Controller {
	return Controller{
		machines: machinesGroup.NewMachinesGroup(),
	}
}

func (c Controller) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Machine":
		return c.machines.GetMachineList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (c Controller) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Machine":
		return c.machines.GetMachine(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
