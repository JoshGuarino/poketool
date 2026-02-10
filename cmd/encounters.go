package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/encounters"
	"github.com/spf13/cobra"
)

var encountersGroups = []string{"Encounter Method", "Encounter Condition", "Encounter Condtiion Value"}

// encountersCmd represents the encounters command
var encountersCmd = &cobra.Command{
	Use:   "encounters",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#encounters-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[any]{}

		// declare instance of encounters Controller
		controller := encounters.NewController()

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select encounters resource group", encountersGroups)
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
			list, _ := controller.GetList(resource)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, resource)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(encountersCmd)
	encountersCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	encountersCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
