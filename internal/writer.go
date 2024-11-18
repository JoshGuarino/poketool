package internal

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func WriteJSON(data interface{}, outputDir string, outputName string) {
	file, err := os.Create(outputName + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
}

func WriteYAML(data interface{}, outputDir string, outputName string) {
	file, err := os.Create(outputName + ".yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	yaml, _ := yaml.Marshal(data)
	file.Write(yaml)
}

func WriteXML(data interface{}, outputDir string, outputName string) {
	file, err := os.Create(outputName + ".xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func WriteCSV(data interface{}, outputDir string, outputName string) {
	file, err := os.Create(outputName + ".csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
