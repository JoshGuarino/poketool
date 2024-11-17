package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	body := Get(PokeApiBaseUrl)
	var data interface{}
	json.Unmarshal(body, &data)
	dataMap, _ := data.(map[string]interface{})
	pokemonUrl, _ := dataMap["pokemon"].(string)
	assert.Equal(t, "https://pokeapi.co/api/v2/pokemon/", pokemonUrl, "Expected to receive 'pokemon' url.")
}

func TestGetResourceList(t *testing.T) {
	r := GetResourceList("https://pokeapi.co/api/v2/pokemon/")
	assert.Equal(t, "bulbasaur", r.Results[0].Name, "Expected to receive 'bulbasaur'.")
}
