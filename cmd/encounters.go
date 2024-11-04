package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/encounters"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// encountersCmd represents the encounters command
var encountersCmd = &cobra.Command{
	Use:   "encounters",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#encounters-section",
	Run: func(cmd *cobra.Command, args []string) {
		// select prompt
		prompt := promptui.Select{
			Label: "Select encounters group resource",
			Items: []string{"Encounter Method", "Encounter Condition", "Encounter Condtiion Value"},
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

			s, err := encounters.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		e := encounters.GetList(result)
		fmt.Println(e)
	},
}

func init() {
	rootCmd.AddCommand(encountersCmd)
	encountersCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
