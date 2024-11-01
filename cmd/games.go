package cmd

import (
	"fmt"
  "github.com/joshguarino/poketool/internal/games"
  "github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// gamesCmd represents the games command
var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
	// select prompt
		prompt := promptui.Select{
			Label: "Select berries group resource",
			Items: []string{"Generation", "Pokedex", "Version", "Version Group"},
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

			s, err := games.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		g := games.GetList(result)
		fmt.Println(g)
	},
}

func init() {
	rootCmd.AddCommand(gamesCmd)
	gamesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
