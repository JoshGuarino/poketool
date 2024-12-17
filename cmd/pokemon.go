package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/pokemon"
	"github.com/spf13/cobra"
)

var resourceGroups = []string{
	"Abilities", "Characteristics", "Egg Groups", "Genders", "Growth Rates",
	"Natures", "Pokeathlon Stats", "Pokemon", "Pokemon Colors", "Pokemon Forms",
	"Pokemon Habitats", "Pokemon Shapes", "Pokemon Species", "Stats", "Types",
}

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#pokemon-section",
	Run: func(cmd *cobra.Command, args []string) {

		// select prompt
		prompt := internal.CreateListPrompt("Select pokemon resource group", resourceGroups)
		resourceGroup := internal.RunSelectPrompt(prompt)

		// flag to search for specific resource
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			s, err := pokemon.GetSpecific(resourceGroup, search)
			if err != nil {
				return
			}
			internal.DecideToOutputFileOrNot(outputToFile, s, resourceGroup)
			fmt.Println(s)
			return
		}

		// if not searching for we return a paginated list
		p := pokemon.GetList(resourceGroup)
		internal.DecideToOutputFileOrNot(outputToFile, p, resourceGroup)
		fmt.Println(p)
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
	pokemonCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	pokemonCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
