package internal

import (
	"github.com/manifoldco/promptui"
)

type Data[T any] struct {
	Data T
}

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

func RunSelectPrompt(prompt promptui.Select) string {
	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}

func RunSearchPrompt(prompt promptui.Prompt) string {
	search, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return search
}

func OutputToFile(data interface{}, filename string) {
	fileType := RunSelectPrompt(CreateFileOutputPrompt())
	WriteFile(fileType, data, "test", filename)
}
