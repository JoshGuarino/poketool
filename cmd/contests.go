package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal"
	"github.com/joshguarino/poketool/internal/resources/contests"
	"github.com/spf13/cobra"
)

var contestsGroups = []string{"Contest Type", "Contest Effect", "Super Contest Effect"}

// contestsCmd represents the contests command
var contestsCmd = &cobra.Command{
	Use:   "contests",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#contests-section",
	Run: func(cmd *cobra.Command, args []string) {

		// generic data holder struct
		data := internal.Data[interface{}]{}

		// select prompt
		selectPrompt := internal.CreateListPrompt("Select contests resource group", contestsGroups)
		contestsGroup := internal.RunSelectPrompt(selectPrompt)

		// flag to search for specific resource else return paginated list
		if search {
			search := internal.RunSearchPrompt(internal.CreateSearchPrompt())
			s, err := contests.GetSpecific(contestsGroup, search)
			if err != nil {
				return
			}
			data.Data = s
		} else {
			c := contests.GetList(contestsGroup)
			data.Data = c
		}

		// create file if output flag exists
		if outputToFile {
			internal.OutputToFile(data.Data, contestsGroup)
		}
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(contestsCmd)
	contestsCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
	contestsCmd.Flags().BoolVarP(&outputToFile, "output", "o", false, "Output data to file")
}
