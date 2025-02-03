package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/locations"
	"github.com/spf13/cobra"
)

var locationsGroups = []string{"Location", "Location Area", "Pal Park Area", "Region"}

// locationsCmd represents the locations command
var locationsCmd = &cobra.Command{
	Use:   "locations",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#locations-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// declare instance of locations Controller
		var controller internal.IController = locations.Controller{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select locations resource group", locationsGroups)
		locationsGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			resource, err := controller.GetSpecific(locationsGroup, search)
			if err != nil {
				return
			}
			data.Data = resource
		} else {
			list := controller.GetList(locationsGroup)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, locationsGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(locationsCmd)
	locationsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	locationsCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
