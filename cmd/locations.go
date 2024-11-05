package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/locations"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// locationsCmd represents the locations command
var locationsCmd = &cobra.Command{
	Use:   "locations",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#locations-section",
	Run: func(cmd *cobra.Command, args []string) {
		// select prompt
		prompt := promptui.Select{
			Label: "Select locations group resource",
			Items: []string{"Location", "Location Area", "Pal Park Area", "Region"},
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

			s, err := locations.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		l := locations.GetList(result)
		fmt.Println(l)
	},
}

func init() {
	rootCmd.AddCommand(locationsCmd)
	locationsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
