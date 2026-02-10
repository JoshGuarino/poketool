package evolution

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/joshguarino/poketool/internal"
	"github.com/stretchr/testify/assert"
)

var controller internal.IController = NewController()

func TestGetList(t *testing.T) {
	rChain, _ := controller.GetList("Evolution Chain", 20, 0)
	rTrigger, _ := controller.GetList("Evolution Trigger", 20, 0)
	_, err := controller.GetList("test", 20, 0)
	assert.IsType(t, &models.ResourceList{}, rChain, "Expected to have array of type 'ResourceList' struct.")
	assert.NotEmpty(t, rChain, "Expected to not have an empty struct.")
	assert.IsType(t, &models.NamedResourceList{}, rTrigger, "Expected to have array of type 'NamedResourceList' struct.")
	assert.NotEmpty(t, rTrigger, "Expected to not have an empty struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSpecific(t *testing.T) {
	rChain, _ := controller.GetSpecific("Evolution Chain", "1")
	rTrigger, _ := controller.GetSpecific("Evolution Trigger", "1")
	_, err := controller.GetSpecific("test", "test")
	assert.IsType(t, &models.EvolutionChain{}, rChain, "Expected to have type 'EvolutionChain' struct.")
	assert.IsType(t, &models.EvolutionTrigger{}, rTrigger, "Expected to have type 'EvolutionTrigger' struct.")
	assert.Error(t, err, "Expected an error to be thrown.")
}
