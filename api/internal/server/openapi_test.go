// api/internal/server/openapi_test.go
package server

import (
	"os"
	"strings"
	"testing"
)

func TestOpenAPIListsEveryRoute(t *testing.T) {
	b, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc := string(b)
	for _, p := range []string{"/healthz:", "/version:", "/v1/subscribe:", "/v1/subscribe/confirm:", "/v1/subscribe/unsubscribe:"} {
		if !strings.Contains(doc, p) {
			t.Errorf("openapi.yaml missing path %s", p)
		}
	}
}
