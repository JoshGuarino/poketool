package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// movesCmd represents the moves command
var movesCmd = &cobra.Command{
	Use:   "moves",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("moves called")
	},
}

func init() {
	rootCmd.AddCommand(movesCmd)
}
