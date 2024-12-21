package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/games"
	"github.com/spf13/cobra"
)

var gamesGroups = []string{"Generation", "Pokedex", "Version", "Version Group"}

// gamesCmd represents the games command
var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#games-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select games resource group", gamesGroups)
		gamesGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			s, err := games.GetSpecific(gamesGroup, search)
			if err != nil {
				return
			}
			data.Data = s
		} else {
			g := games.GetList(gamesGroup)
			data.Data = g
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, gamesGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(gamesCmd)
	gamesCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	gamesCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
