package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/contests"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// contestsCmd represents the contests command
var contestsCmd = &cobra.Command{
	Use:   "contests",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#contests-section",
	Run: func(cmd *cobra.Command, args []string) {
		// select prompt
		prompt := promptui.Select{
			Label: "Select contests group resource",
			Items: []string{"Contest Type", "Contest Effect", "Super Contest Effect"},
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

			s, err := contests.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		c := contests.GetList(result)
		fmt.Println(c)
	},
}

func init() {
	rootCmd.AddCommand(contestsCmd)
	contestsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
