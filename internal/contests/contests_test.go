package contests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetContestType(t *testing.T) {
	rById, _ := GetContestType("1")
	rByName, _ := GetContestType("cool")
	rFail, _ := GetContestType("test")
	assert.Equal(t, "cool", rById.Name, "Expected to receive Cool.")
	assert.Equal(t, "cool", rByName.Name, "Expected to receive Cool.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetContestTypeList(t *testing.T) {
	r := GetContestTypeList()
	assert.Equal(t, "cool", r.Results[0].Name, "Expected to have a list of contest type resource.")
}

func TestGetContestEffect(t *testing.T) {
	rById, _ := GetContestEffect("1")
	rFail, _ := GetContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive contest type with id of 1.")
	assert.Equal(t, 0, rFail.ID, "Expect to receive an empty result.")
}

func TestGetContestEffectList(t *testing.T) {
	r := GetContestEffectList()
	assert.Equal(t, "https://pokeapi.co/api/v2/contest-effect/1/", r.Results[0].URL, "Expected to have a list of contest effect resource.")
}

func TestGetSuperContestEffect(t *testing.T) {
	rById, _ := GetSuperContestEffect("1")
	rFail, _ := GetSuperContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive Cool.")
	assert.Equal(t, 0, rFail.ID, "Expect to receive Cool.")
}

func TestGetSuperContestEffectList(t *testing.T) {
	r := GetSuperContestEffectList()
	assert.Equal(t, "https://pokeapi.co/api/v2/super-contest-effect/1/", r.Results[0].URL, "Expected to have a list of super contest effect resource.")
}
