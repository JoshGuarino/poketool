package contests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var contests IContests = Contests{}

func TestGetContestType(t *testing.T) {
	rById, _ := contests.GetContestType("1")
	rByName, _ := contests.GetContestType("cool")
	rFail, _ := contests.GetContestType("test")
	assert.Equal(t, "cool", rById.Name, "Expected to receive 'cool'.")
	assert.Equal(t, "cool", rByName.Name, "Expected to receive 'cool'.")
	assert.Equal(t, "", rFail.Name, "Expected to receive an empty result.")
}

func TestGetContestTypeList(t *testing.T) {
	r := contests.GetContestTypeList()
	assert.Equal(t, "cool", r.Results[0].Name, "Expected to have a list of 'contest type' resource.")
}

func TestGetContestEffect(t *testing.T) {
	rById, _ := contests.GetContestEffect("1")
	rFail, _ := contests.GetContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive id of '1'.")
	assert.Equal(t, 0, rFail.ID, "Expect to receive an empty result.")
}

func TestGetContestEffectList(t *testing.T) {
	r := contests.GetContestEffectList()
	assert.Equal(t, "https://pokeapi.co/api/v2/contest-effect/1/", r.Results[0].URL, "Expected to have a list of 'contest effect' resource.")
}

func TestGetSuperContestEffect(t *testing.T) {
	rById, _ := contests.GetSuperContestEffect("1")
	rFail, _ := contests.GetSuperContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive id of 1.")
	assert.Equal(t, 0, rFail.ID, "Expect to receive and empty result.")
}

func TestGetSuperContestEffectList(t *testing.T) {
	r := contests.GetSuperContestEffectList()
	assert.Equal(t, "https://pokeapi.co/api/v2/super-contest-effect/1/", r.Results[0].URL, "Expected to have a list of 'super contest effect' resource.")
}
