package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/contests"
	"github.com/spf13/cobra"
)

// contestsCmd represents the contests command
var contestsCmd = &cobra.Command{
	Use:   "contests",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#contests-section",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(contests.GetSuperContestEffect("1"))
	},
}

func init() {
	rootCmd.AddCommand(contestsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contestsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contestsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
