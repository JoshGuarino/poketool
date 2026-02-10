package contests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var contests IContests = NewContests()

func TestGetContestType(t *testing.T) {
	rById, _ := contests.GetContestType("1")
	rByName, _ := contests.GetContestType("cool")
	_, err := contests.GetContestType("test")
	assert.Equal(t, "cool", rById.Name, "Expected to receive 'cool'.")
	assert.Equal(t, "cool", rByName.Name, "Expected to receive 'cool'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestTypeList(t *testing.T) {
	r, _ := contests.GetContestTypeList(20, 0)
	assert.Equal(t, "cool", r.Results[0].Name, "Expected to have a list of 'contest type' resource.")
}

func TestGetContestEffect(t *testing.T) {
	rById, _ := contests.GetContestEffect("1")
	_, err := contests.GetContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive id of '1'.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetContestEffectList(t *testing.T) {
	r, _ := contests.GetContestEffectList(20, 0)
	assert.Equal(t, "https://staging.pokeapi.co/api/v2/contest-effect/1/", r.Results[0].URL, "Expected to have a list of 'contest effect' resource.")
}

func TestGetSuperContestEffect(t *testing.T) {
	rById, _ := contests.GetSuperContestEffect("1")
	_, err := contests.GetSuperContestEffect("test")
	assert.Equal(t, 1, rById.ID, "Expect to receive id of 1.")
	assert.Error(t, err, "Expected an error to be thrown.")
}

func TestGetSuperContestEffectList(t *testing.T) {
	r, _ := contests.GetSuperContestEffectList(20, 0)
	assert.Equal(t, "https://staging.pokeapi.co/api/v2/super-contest-effect/1/", r.Results[0].URL, "Expected to have a list of 'super contest effect' resource.")
}
