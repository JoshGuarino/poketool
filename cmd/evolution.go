package cmd

import (
	"fmt"

	"github.com/joshguarino/poketool/internal/evolution"
  "github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// evolutionCmd represents the evolution command
var evolutionCmd = &cobra.Command{
	Use:   "evolution",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    // select prompt
		prompt := promptui.Select{
			Label: "Select berries group resource",
			Items: []string{"Evolution Chain", "Evolution Trigger"},
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

			s, err := evolution.GetSpecific(result, search)
			if err != nil {
				return
			}
			fmt.Println(s)
			return
		}

		e := evolution.GetList(result)
		fmt.Println(e)
	},
}

func init() {
	rootCmd.AddCommand(evolutionCmd)
	evolutionCmd.Flags().BoolVarP(&search, "search", "s", false, "Find specific resource")
}
