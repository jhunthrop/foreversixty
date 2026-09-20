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

	requiredPaths := []string{
		"/health", "/version",
		"/v1/subscribe", "/v1/subscribe/confirm", "/v1/subscribe/unsubscribe",
		"/v1/builds", "/v1/builds/{id}", "/b/{id}", "/b/{id}/card.png",
		"/v1/auth/battlenet/start", "/v1/auth/battlenet/callback",
		"/v1/auth/email", "/v1/auth/email/callback", "/v1/auth/logout", "/v1/sessions", "/v1/me",
		"/v1/devices/pair", "/v1/devices/claim", "/v1/devices", "/v1/devices/{id}",
		"/v1/reports", "/v1/reports/{id}", "/v1/reports/{id}/visibility",
		"/v1/reports/{id}/access", "/v1/reports/{id}/files/{path}",
		"/v1/reports/{id}/fights/{n}", "/v1/reports/{id}/fights/{n}/live",
		"/v1/reports/{id}/raw", "/v1/reports/{id}/complete",
		"/v1/uploads", "/v1/uploads/{upload_id}/complete",
		"/v1/rankings", "/v1/rankings/percentile", "/v1/rankings/guilds",
		"/v1/characters/{region}/{ruleset}/{name}", "/v1/guilds/{region}/{ruleset}/{name}",
		"/v1/addon/exports", "/v1/addon/inbox", "/reports/{id}/card.png",
		"/v1/sims", "/v1/sims/{id}", "/v1/sims/{id}/progress", "/v1/sims/run",
		"/v1/specs", "/v1/characters/{region}/{ruleset}/{name}/sim-input",
		"/v1/phases",
	}
	for _, p := range requiredPaths {
		if _, ok := paths[p]; !ok {
			t.Errorf("openapi.yaml missing path %s", p)
		}
	}

	components, ok := doc["components"].(map[string]any)
	if !ok {
		t.Fatal("openapi.yaml missing 'components'")
	}
	schemas, ok := components["schemas"].(map[string]any)
	if !ok {
		t.Fatal("openapi.yaml missing 'components.schemas'")
	}
	for _, s := range []string{"Envelope", "Build", "BuildInput", "Report", "FightEntry",
		"MetricsRow", "RankingRow", "SimResult", "SimRequest", "SimRow", "SpecFidelity", "User"} {
		if _, ok := schemas[s]; !ok {
			t.Errorf("openapi.yaml missing schema %s", s)
		}
	}
}
