package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/items"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var itemsGroups = []string{"Item", "Item Attribute", "Item Category", "Item Fling Effect", "Item Pocket"}

// itemsCmd represents the items command
var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#items-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select itmes resource group", itemsGroups)
		itemsGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			s, err := items.GetSpecific(itemsGroup, search)
			if err != nil {
				return
			}
			data.Data = s
		} else {
			i := items.GetList(itemsGroup)
			data.Data = i
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, itemsGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(itemsCmd)
	itemsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
