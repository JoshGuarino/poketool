package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/contests"
	"github.com/spf13/cobra"
)

var contestsResources = []string{"Contest Type", "Contest Effect", "Super Contest Effect"}

// contestsCmd represents the contests command
var contestsCmd = &cobra.Command{
	Use:   "contests",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#contests-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[any]{}

		// declare instance of contests Controller
		controller := contests.NewController()

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select contests resource group", contestsResources)
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
	rootCmd.AddCommand(contestsCmd)
	contestsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	contestsCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
