package games

import (
	"fmt"

	gamesGroup "github.com/JoshGuarino/PokeGo/pkg/resources/games"
)

type GamesController struct {
	games gamesGroup.Games
}

func NewController() GamesController {
	return GamesController{
		games: gamesGroup.NewGamesGroup(),
	}
}

func (controller GamesController) GetList(result string, limit int, offset int) (any, error) {
	switch result {
	case "Generation":
		return controller.games.GetGenerationList(limit, offset)
	case "Pokedex":
		return controller.games.GetPokedexList(limit, offset)
	case "Version":
		return controller.games.GetVersionList(limit, offset)
	case "Version Group":
		return controller.games.GetVersionGroupList(limit, offset)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}

func (controller GamesController) GetSpecific(result string, search string) (any, error) {
	switch result {
	case "Generation":
		return controller.games.GetGeneration(search)
	case "Pokedex":
		return controller.games.GetPokedex(search)
	case "Version":
		return controller.games.GetVersion(search)
	case "Version Group":
		return controller.games.GetVersionGroup(search)
	}

	return nil, fmt.Errorf("Unable to find resource %s in group", result)
}
