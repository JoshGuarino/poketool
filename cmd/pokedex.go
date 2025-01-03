package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pokedexCmd represents the pokedex command
var pokedexCmd = &cobra.Command{
	Use:   "pokedex",
	Short: "Search and display pokemon data like a real pokedex",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pokedex called")
	},
}

func init() {
	rootCmd.AddCommand(pokedexCmd)
}
