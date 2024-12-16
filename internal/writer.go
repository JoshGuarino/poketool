package internal

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"gopkg.in/yaml.v3"
)

func WriteFile(fileType string, data interface{}, outputDir string, outputName string) {
	switch fileType {
	case "JSON":
		WriteJSON(data, outputDir, outputName)
		return
	case "YAML":
		WriteYAML(data, outputDir, outputName)
		return
	case "XML":
		WriteXML(data, outputDir, outputName)
		return
	}
}

func ensureDirectoryExists(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func WriteJSON(data interface{}, outputDir string, outputName string) {
	ensureDirectoryExists(outputDir)
	file, err := os.Create(outputDir + "/" + outputName + ".json")
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
	ensureDirectoryExists(outputDir)
	file, err := os.Create(outputDir + "/" + outputName + ".yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	yaml, _ := yaml.Marshal(data)
	file.Write(yaml)
}

func WriteXML(data interface{}, outputDir string, outputName string) {
	ensureDirectoryExists(outputDir)
	file, err := os.Create(outputDir + "/" + outputName + ".xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := xml.NewEncoder(file)
	encoder.Indent("", " ")
	encoder.Encode(data)
}
