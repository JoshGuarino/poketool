package internal

import (
	"os"
	"testing"
)

type TestHelper struct {
	TestDir  string
	TestData map[string]string
}

var th = &TestHelper{
	TestDir:  "test_dir",
	TestData: map[string]string{"key": "value"},
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

func TestWtiteJSON(t *testing.T) {
	th.setup(t)
	defer th.teardown()
	WriteJSON(th.TestData, th.TestDir, "test")
}

func TestWtiteYAML(t *testing.T) {
	th.setup(t)
	defer th.teardown()
	WriteYAML(th.TestData, th.TestDir, "test")
}

func TestWtiteXML(t *testing.T) {
	th.setup(t)
	defer th.teardown()
	WriteXML(th.TestData, th.TestDir, "test")
}
