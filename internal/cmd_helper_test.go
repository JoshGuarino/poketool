package internal

import (
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
)

func TestCreateListPrompt(t *testing.T) {
	prompt := CreateListPrompt("test", []string{"1", "2"})
	assert.Equal(t, "test", prompt.Label, "Expected prompt to have label of 'test'.")
	assert.Equal(t, []string{"1", "2"}, prompt.Items, "Expected prompt to have select items of '1,2'.")
	assert.IsType(t, promptui.Select{}, prompt, "Expected to be instance of 'promptui.Select'.")
}

func TestCreateSearchPrompt(t *testing.T) {
	prompt := CreateSearchPrompt()
	assert.Equal(t, "Search", prompt.Label, "Expected prompt to have label of 'Search'.")
	assert.IsType(t, promptui.Prompt{}, prompt, "Expected to be instance of 'promptui.Prompt'.")
}

func TestFileOutputPrompt(t *testing.T) {
	prompt := CreateFileOutputPrompt()
	assert.Equal(t, "Select output file type", prompt.Label, "Expected prompt to have label of 'Select output file type'.")
	assert.Equal(t, []string{"JSON", "YAML", "XML"}, prompt.Items, "Expected prompt to have select items of 'JSON,YAML,XML'.")
	assert.IsType(t, promptui.Select{}, prompt, "Expected to be instance of 'promptui.Select'.")
}
