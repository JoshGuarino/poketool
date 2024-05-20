/*
Copyright © 2024 Josh Guarino <jguarino722@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/pokemon"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var search bool

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#pokemon-section",
	Run: func(cmd *cobra.Command, args []string) {

		// select prompt
		prompt := promptui.Select{
			Label: "Select pokemon group resource",
			Items: []string{"Abilities", "Characteristics", "Egg group", "Gender", "Growth Rate", "Nature",
				"Pokeathlon Stat", "Pokemon", "Pokemon Color", "Pokemon Form", "Pokemon Habitat", "Pokemon Shape",
				"Pokemon Species", "Stat", "Type"},
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
			result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			s := pokemon.GetSpecfic(result)
			fmt.Println(s)
			return
		}

		l := pokemon.GetList(result)
		fmt.Println(l)

		// p := pokemon.GetPokemon("1")
		// s, _ := json.MarshalIndent(p, "", "\t")
		// fmt.Printf("%+v\n", s)
		// fmt.Println(string(s))
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
	pokemonCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
