package internal

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type TestData struct {
	Number int    `json:"number" yaml:"number" xml:"number"`
	Name   string `json:"name" yaml:"name" xml:"name"`
}

type TestHelper struct {
	TestDir     string
	TestData    TestData
	DecodedData TestData
}

var th = &TestHelper{
	TestDir:  "test_dir",
	TestData: TestData{Number: 4, Name: "charmander"},
}

func (th *TestHelper) setup(t *testing.T) {
	err := os.Mkdir(th.TestDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
}

func (th *TestHelper) teardown() {
	os.RemoveAll(th.TestDir)
}

func TestWriteFile(t *testing.T) {
	// implement this
}

func TestEnsureDirectoryExists(t *testing.T) {
	// implement this
}

func TestWtiteJSON(t *testing.T) {
	th.setup(t)
	defer th.teardown()
	WriteJSON(th.TestData, th.TestDir, "test")

	// verify file exists
	outputPath := filepath.Join(th.TestDir, "test.json")
	_, err := os.Stat(outputPath)
	assert.NoError(t, err)

	// verify content in file
	file, err := os.ReadFile(outputPath)
	assert.NoError(t, err)
	err = json.Unmarshal(file, &th.DecodedData)
	assert.NoError(t, err)
	assert.Equal(t, th.TestData, th.DecodedData)
}

func TestWtiteYAML(t *testing.T) {
	th.setup(t)
	defer th.teardown()
	WriteYAML(th.TestData, th.TestDir, "test")

	// verify file exists
	outputPath := filepath.Join(th.TestDir, "test.yaml")
	_, err := os.Stat(outputPath)
	assert.NoError(t, err)

	// verify content in file
	file, err := os.ReadFile(outputPath)
	assert.NoError(t, err)
	err = yaml.Unmarshal(file, &th.DecodedData)
	assert.NoError(t, err)
	assert.Equal(t, th.TestData, th.DecodedData)
}

func TestWtiteXML(t *testing.T) {
	th.setup(t)
	defer th.teardown()
	WriteXML(th.TestData, th.TestDir, "test")

	// verify file exists
	outputPath := filepath.Join(th.TestDir, "test.xml")
	_, err := os.Stat(outputPath)
	assert.NoError(t, err)

	// verify content in file
	file, err := os.ReadFile(outputPath)
	assert.NoError(t, err)
	err = xml.Unmarshal(file, &th.DecodedData)
	assert.NoError(t, err)
	assert.Equal(t, th.TestData, th.DecodedData)
}
