package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/moves"
	"github.com/spf13/cobra"
)

var movesGroups = []string{"Move", "Move Ailment", "Move Battle Style", "Move Category", "Move Damage Class", "Move Learn Method", "Move Target"}

// movesCmd represents the moves command
var movesCmd = &cobra.Command{
	Use:   "moves",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#moves-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// declare instance of moves Controller
		var controller internal.IController = moves.Controller{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select machines resource group", movesGroups)
		movesGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			resource, err := controller.GetSpecific(movesGroup, search)
			if err != nil {
				return
			}
			data.Data = resource
		} else {
			list := controller.GetList(movesGroup)
			data.Data = list
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, movesGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(movesCmd)
	movesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	movesCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
