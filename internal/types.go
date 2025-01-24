package internal

import (
	"github.com/mtslzr/pokeapi-go/structs"
)

type Data[T any] struct {
	Data T
}

type IController interface {
	GetList(result string) structs.Resource
	GetSpecific(result string) (interface{}, error)
}

type LocationAreaEncounter struct {
	LocationArea  structs.Result           `json:"location_area"`
	VersionDetail []VersionEncounterDetail `json:"version_details"`
}

type VersionEncounterDetail struct {
	Version          structs.Result `json:"version"`
	MaxChance        int            `json:"max_chance"`
	EncounterDetails []Encounter    `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int              `json:"min_level"`
	MaxLevel        int              `json:"max_level"`
	ConditionValues []structs.Result `json:"condition_values"`
	Chance          int              `json:"chance"`
	Method          structs.Result   `json:"method"`
}
