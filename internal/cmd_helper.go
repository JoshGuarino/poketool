package internal

import (
	"github.com/manifoldco/promptui"
)

func CreateListPrompt(label string, items []string) promptui.Select {
	return promptui.Select{
		Label: label,
		Items: items,
	}
}

func CreateSearchPrompt() promptui.Prompt {
	return promptui.Prompt{
		Label: "Search",
	}
}

func CreateFileOutputPrompt() promptui.Select {
	return promptui.Select{
		Label: "Select output file type",
		Items: []string{"JSON", "YAML", "XML"},
	}
}

func RunOutputFileProtocol(prompt promptui.Select) string {
	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}
