package cmd

import (
	"fmt"
	"os"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/pokemon"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#pokemon-section",
	Run: func(cmd *cobra.Command, args []string) {

		// select prompt
		prompt := promptui.Select{
			Label: "Select pokemon group resource",
			Items: []string{"Abilities", "Characteristics", "Egg Groups", "Genders", "Growth Rates", "Natures",
				"Pokeathlon Stats", "Pokemon", "Pokemon Colors", "Pokemon Forms", "Pokemon Habitats", "Pokemon Shapes",
				"Pokemon Species", "Stats", "Types"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// flag to search for specific resource
		if search {
			// search prompt
			prompt := promptui.Prompt{
				Label: "Search",
			}
			search, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			s, err := pokemon.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		p := pokemon.GetList(result)
		// fmt.Println(p)
		dir, _ := os.Getwd()
		internal.WriteJSON(p, dir, "test")
		internal.WriteYAML(p, dir, "test")
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
	pokemonCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
