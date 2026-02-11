package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/berries"
	"github.com/joshguarino/poketool/internal/resources/contests"
	"github.com/joshguarino/poketool/internal/resources/encounters"
	"github.com/joshguarino/poketool/internal/resources/evolution"
	"github.com/joshguarino/poketool/internal/resources/games"
	"github.com/joshguarino/poketool/internal/resources/items"
	"github.com/joshguarino/poketool/internal/resources/locations"
	"github.com/joshguarino/poketool/internal/resources/machines"
	"github.com/joshguarino/poketool/internal/resources/moves"
	// "github.com/joshguarino/poketool/internal/resources/pokemon"
	"github.com/spf13/cobra"
)

type Command struct {
	Name        string
	Description string
	Resources   []string
	controller  internal.IController
}

var commands = []Command{
	{
		Name:        "berries",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#berries-section",
		Resources: []string{
			"Berries",
			"Berry Firmnesses",
			"Berry Flavors",
		},
		controller: berries.NewController(),
	},
	{
		Name:        "contests",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#contests-section",
		Resources: []string{
			"Contest Type",
			"Contest Effect",
			"Super Contest Effect",
		},
		controller: contests.NewController(),
	},
	{
		Name:        "encounters",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#encounters-section",
		Resources: []string{
			"Encounter Method",
			"Encounter Condition",
			"Encounter Condtition Value",
		},
		controller: encounters.NewController(),
	},
	{
		Name:        "evolution",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#evolution-section",
		Resources: []string{
			"Evolution Chain",
			"Evolution Trigger",
		},
		controller: evolution.NewController(),
	},
	{
		Name:        "games",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#games-section",
		Resources: []string{
			"Generation",
			"Pokedex",
			"Version",
			"Version Group",
		},
		controller: games.NewController(),
	},
	{
		Name:        "items",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#items-section",
		Resources: []string{
			"Item",
			"Item Attribute",
			"Item Category",
			"Item Fling Effect",
			"Item Pocket",
		},
		controller: items.NewController(),
	},
	{
		Name:        "locations",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#locations-section",
		Resources: []string{
			"Location",
			"Location Area",
			"Pal Park Area",
			"Region",
		},
		controller: locations.NewController(),
	},
	{
		Name:        "machines",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#machines-section",
		Resources: []string{
			"Machine",
		},
		controller: machines.NewController(),
	},
	{
		Name:        "moves",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#moves-section",
		Resources: []string{
			"Move",
			"Move Ailment",
			"Move Battle Style",
			"Move Category",
			"Move Damage Class",
			"Move Learn Method",
			"Move Target",
		},
		controller: moves.NewController(),
	},
	{
		Name:        "pokemon",
		Description: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#pokemon-section",
		Resources: []string{
			"Abilities",
			"Characteristics",
			"Egg Groups",
			"Genders",
			"Growth Rates",
			"Natures",
			"Pokeathlon Stats",
			"Pokemon",
			"Pokemon Colors",
			"Pokemon Forms",
			"Pokemon Habitats",
			"Pokemon Shapes",
			"Pokemon Species",
			"Stats",
			"Types",
		},
	},
}

func createSubCommand(command Command) *cobra.Command {
	// TODO: clean this mess up and think of switching to bubbletea
	return &cobra.Command{
		Use:   command.Name,
		Short: command.Description,
		Run: func(cmd *cobra.Command, args []string) {
			// generic data holder struct
			data := internal.Data[any]{}

			// declare instance of pokemon Controller
			var controller internal.IController = command.controller

			// select prompt
			selectPrompt := internal.CreateListPrompt("Select pokemon resource group", command.Resources)
			resource := internal.RunSelectPrompt(selectPrompt)

			// flag to search for specific resource else return paginated list
			if search {
				search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
				resource, err := controller.GetSpecific(resource, search)
				if err != nil {
					return
				}
				data.Data = resource
			} else {
				list, _ := controller.GetList(resource, 20, 0) // TODO: handle pagination
				data.Data = list
			}

			// create file if output flag exists
			if outputToFile {
				internal.OutputToFile(data.Data, resource)
			}
			fmt.Println(data.Data)
		},
	}
}

func init() {
	for _, command := range commands {
		subCmd := createSubCommand(command)
		rootCmd.AddCommand(subCmd)
		subCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
		subCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
	}
}
