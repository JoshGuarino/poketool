package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/evolution"
	"github.com/spf13/cobra"
)

var evolutionGroups = []string{"Evolution Chain", "Evolution Trigger"}

// evolutionCmd represents the evolution command
var evolutionCmd = &cobra.Command{
	Use:   "evolution",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#evolution-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// declare instance of evolution Controller
		var controller internal.IController = evolution.Controller{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select evolution resource group", evolutionGroups)
		evolutionGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			resource, err := controller.GetSpecific(evolutionGroup, search)
			if err != nil {
				return
			}
			data.Data = resource
		} else {
			list := controller.GetList(evolutionGroup)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, evolutionGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(evolutionCmd)
	evolutionCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	evolutionCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
