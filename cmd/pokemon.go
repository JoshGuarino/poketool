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

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#pokemon-section",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Select pokemon group resource",
			Items: []string{"abilities", "characteristics", "egg group", "gender", "growth rate", "nature",
				"pokeathlon stat", "pokemon", "pokemon color", "pokemon form", "pokemon habitat", "pokemon shape",
				"pokemon species", "stat", "type"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose %q\n", result)

		// fmt.Println(prompt)
		pokemon.GetPokemonLocationAreas("1")
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	pokemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pokemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
