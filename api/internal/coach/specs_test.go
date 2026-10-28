package coach

import (
	"encoding/json"
	"os"
	"testing"
)

const specsJSONPath = "../../../data/curated/specs.json"

// specInfo is the part of a specs.json row the coach tests need.
type specInfo struct {
	Spec  string `json:"spec"`
	Class string `json:"class_slug"`
	Role  string `json:"role"`
}

func readTestFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

// loadSpecs reads every spec from data/curated/specs.json.
func loadSpecs(t *testing.T) []specInfo {
	t.Helper()
	data, err := os.ReadFile(specsJSONPath)
	if err != nil {
		t.Fatalf("read specs.json: %v", err)
	}
	var specs []specInfo
	if err := json.Unmarshal(data, &specs); err != nil {
		t.Fatalf("decode specs.json: %v", err)
	}
	return specs
}
