package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/machines"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// machinesCmd represents the machines command
var machinesCmd = &cobra.Command{
	Use:   "machines",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#locations-section",
	Run: func(cmd *cobra.Command, args []string) {
		// select prompt
		prompt := promptui.Select{
			Label: "Select machines group resource",
			Items: []string{"Machine"},
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

			s, err := machines.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		l := machines.GetList(result)
		fmt.Println(l)
	},
}

func init() {
	rootCmd.AddCommand(machinesCmd)
	machinesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
