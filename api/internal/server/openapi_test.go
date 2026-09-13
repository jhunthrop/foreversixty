// api/internal/server/openapi_test.go
package server

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOpenAPIListsEveryRoute(t *testing.T) {
	b, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}

	var doc map[string]any
	err = yaml.Unmarshal(b, &doc)
	if err != nil {
		t.Fatalf("openapi.yaml is not valid YAML: %v", err)
	}

	paths, ok := doc["paths"].(map[string]any)
	if !ok {
		t.Fatal("openapi.yaml missing or invalid 'paths' object")
	}

	requiredPaths := []string{"/health", "/version", "/v1/subscribe", "/v1/subscribe/confirm", "/v1/subscribe/unsubscribe"}
	for _, p := range requiredPaths {
		if _, ok := paths[p]; !ok {
			t.Errorf("openapi.yaml missing path %s", p)
		}
	}
}
