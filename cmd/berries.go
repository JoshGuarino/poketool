package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/berries"
	"github.com/spf13/cobra"
)

var berriesGroups = []string{"Berries", "Berry Firmnesses", "Berry Flavors"}

// berriesCmd represents the berries command
var berriesCmd = &cobra.Command{
	Use:   "berries",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#berries-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[any]{}

		// declare instance of berries Controller
		var controller = berries.NewController()

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select berries resource group", berriesGroups)
		berriesGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			resource, err := controller.GetSpecific(berriesGroup, search)
			if err != nil {
				return
			}
			data.Data = resource
		} else {
			list, _ := controller.GetList(berriesGroup)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, berriesGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(berriesCmd)
	berriesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	berriesCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
