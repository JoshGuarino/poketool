package games

import (
	"github.com/JoshGuarino/PokeGo/pkg/models"
	gamesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/games"
)

type IGames interface {
	GetGeneration(nameOrId string) (*models.Generation, error)
	GetGenerationList(limit int, offest int) (*models.NamedResourceList, error)
	GetPokedex(nameOrId string) (*models.Pokedex, error)
	GetPokedexList(limit int, offest int) (*models.NamedResourceList, error)
	GetVersion(nameOrId string) (*models.Version, error)
	GetVersionList(limit int, offest int) (*models.NamedResourceList, error)
	GetVersionGroup(nameOrId string) (*models.VersionGroup, error)
	GetVersionGroupList(limit int, offest int) (*models.NamedResourceList, error)
}

type Games struct {
	gamesGroup gamesGroup.Games
}

func NewGames() Games {
	return Games{
		gamesGroup: gamesGroup.NewGamesGroup(),
	}
}

func (g Games) GetGeneration(nameOrId string) (*models.Generation, error) {
	generation, err := g.gamesGroup.GetGeneration(nameOrId)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

func (g Games) GetGenerationList(limit int, offest int) (*models.NamedResourceList, error) {
	generationList, err := g.gamesGroup.GetGenerationList(limit, offest)
	if err != nil {
		return nil, err
	}
	return generationList, nil
}

func (g Games) GetPokedex(nameOrId string) (*models.Pokedex, error) {
	pokedex, err := g.gamesGroup.GetPokedex(nameOrId)
	if err != nil {
		return nil, err
	}
	return pokedex, nil
}

func (g Games) GetPokedexList(limit int, offest int) (*models.NamedResourceList, error) {
	pokedexList, err := g.gamesGroup.GetPokedexList(limit, offest)
	if err != nil {
		return nil, err
	}
	return pokedexList, nil
}

func (g Games) GetVersion(nameOrId string) (*models.Version, error) {
	version, err := g.gamesGroup.GetVersion(nameOrId)
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (g Games) GetVersionList(limit int, offest int) (*models.NamedResourceList, error) {
	versionList, err := g.gamesGroup.GetVersionList(limit, offest)
	if err != nil {
		return nil, err
	}
	return versionList, nil
}

func (g Games) GetVersionGroup(nameOrId string) (*models.VersionGroup, error) {
	versionGroup, err := g.gamesGroup.GetVersionGroup(nameOrId)
	if err != nil {
		return nil, err
	}
	return versionGroup, nil
}

func (g Games) GetVersionGroupList(limit int, offest int) (*models.NamedResourceList, error) {
	versionGroupList, err := g.gamesGroup.GetVersionGroupList(limit, offest)
	if err != nil {
		return nil, err
	}
	return versionGroupList, nil
}
