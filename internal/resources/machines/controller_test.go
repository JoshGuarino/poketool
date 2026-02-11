package machines

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rMachines, _ := controller.GetList("Machine", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.ResourceList{}, rMachines, "Expected to have type 'ResourceList' struct.")
	assert.NotEmpty(t, rMachines, "Expected to not have an empty struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rMachines, _ := controller.GetSpecific("Machine", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.Machine{}, rMachines, "Expected to have type 'Machine' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
