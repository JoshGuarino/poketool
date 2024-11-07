package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/moves"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// movesCmd represents the moves command
var movesCmd = &cobra.Command{
	Use:   "moves",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#moves-section",
	Run: func(cmd *cobra.Command, args []string) {
		// select prompt
		prompt := promptui.Select{
			Label: "Select moves group resource",
			Items: []string{"Move", "Move Ailment", "Move Battle Style", "Move Category", "Move Damage Class", "Move Learn Method", "Move Target"},
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

			s, err := moves.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		m := moves.GetList(result)
		fmt.Println(m)
	},
}

func init() {
	rootCmd.AddCommand(movesCmd)
	movesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
