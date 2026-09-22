// api/internal/bnetapi/probe_test.go
package bnetapi

import (
	"context"
	"net/http"
	"testing"
)

func TestProbeReportsEachGamesStatus(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"realms": []any{}})
	fs.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic-us", http.StatusForbidden, nil)

	c := newTestClient(fs)
	results := c.Probe(context.Background(), "us", []string{"classic1x", "classic"})
	if len(results) != 2 {
		t.Fatalf("results = %+v", results)
	}
	if results[0].Game != "classic1x" || results[0].Status != http.StatusOK {
		t.Errorf("results[0] = %+v", results[0])
	}
	if results[1].Game != "classic" || results[1].Status != http.StatusForbidden {
		t.Errorf("results[1] = %+v", results[1])
	}
}
