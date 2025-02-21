package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/pokemon"
	"github.com/spf13/cobra"
)

var pokemonGroups = []string{
	"Abilities", "Characteristics", "Egg Groups", "Genders", "Growth Rates",
	"Natures", "Pokeathlon Stats", "Pokemon", "Pokemon Colors", "Pokemon Forms",
	"Pokemon Habitats", "Pokemon Shapes", "Pokemon Species", "Stats", "Types",
}

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#pokemon-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// declare instance of pokemon Controller
		var controller internal.IController = pokemon.Controller{}

		// select prompt
		prompt := internal.CreateListPrompt("Select pokemon resource group", pokemonGroups)
		pokemonGroup := internal.RunSelectPrompt(prompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			resource, err := controller.GetSpecific(pokemonGroup, search)
			if err != nil {
				return
			}
			data.Data = resource
		} else {
			list := controller.GetList(pokemonGroup)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, pokemonGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
	pokemonCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	pokemonCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
