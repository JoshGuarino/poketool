package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/machines"
	"github.com/spf13/cobra"
)

var machinesGroups = []string{"Machine"}

// machinesCmd represents the machines command
var machinesCmd = &cobra.Command{
	Use:   "machines",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#machines-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// declare instance of machines Controller
		var controller internal.IController = machines.Controller{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select machines resource group", machinesGroups)
		machinesGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			resource, err := controller.GetSpecific(machinesGroup, search)
			if err != nil {
				return
			}
			data.Data = resource
		} else {
			list := controller.GetList(machinesGroup)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, machinesGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(machinesCmd)
	machinesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	machinesCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
