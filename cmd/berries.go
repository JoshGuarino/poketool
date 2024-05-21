/*
Copyright © 2024 Josh Guarino <jguarino722@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// berriesCmd represents the berries command
var berriesCmd = &cobra.Command{
	Use:   "berries",
	Short: "Access pokemon resource group data from pokeapi: https://pokeapi.co/docs/v2#berries-section",
	Run: func(cmd *cobra.Command, args []string) {

		// select prompt
		prompt := promptui.Select{
			Label: "Select berries group resource",
			Items: []string{"Berries", "Berry Firmnesses", "Berry Flavors"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(berriesCmd)
}
