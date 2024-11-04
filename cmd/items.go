package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/items"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// itemsCmd represents the items command
var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#items-section",
	Run: func(cmd *cobra.Command, args []string) {
// select prompt
		prompt := promptui.Select{
			Label: "Select items group resource",
			Items: []string{"Item", "Item Attribute", "Item Category", "Item Fling Effect", "Item Pocket"},
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

			s, err := items.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		i := items.GetList(result)
		fmt.Println(i)
	},
}

func init() {
	rootCmd.AddCommand(itemsCmd) 
	itemsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
