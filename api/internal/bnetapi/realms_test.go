// api/internal/bnetapi/realms_test.go
package bnetapi

import (
	"context"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

func TestRulesetOf(t *testing.T) {
	cases := []struct {
		name string
		r    Realm
		want string
	}{
		{"normal", Realm{Type: "NORMAL"}, character.RulesetNormal},
		{"pvp", Realm{Type: "PVP"}, character.RulesetPvP},
		{"rp", Realm{Type: "RP"}, character.RulesetRP},
		{"rp_pvp", Realm{Type: "RP_PVP"}, character.RulesetRP},
		{"hardcore by category", Realm{Type: "PVP", Category: "Hardcore"}, character.RulesetHardcore},
		{"hardcore by name", Realm{Type: "NORMAL", Name: "Doomhowl Hardcore"}, character.RulesetHardcore},
		{"unknown type falls back to normal", Realm{Type: "SEASONAL"}, character.RulesetNormal},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := RulesetOf(c.r); got != c.want {
				t.Errorf("RulesetOf(%+v) = %q, want %q", c.r, got, c.want)
			}
		})
	}
}

func TestRealmsFetchesIndexAndPerRealmDetailAndCaches(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us", http.StatusOK, map[string]any{
		"realms": []map[string]any{
			{"id": 1, "name": "Whitemane", "slug": "whitemane"},
			{"id": 2, "name": "Faerlina", "slug": "faerlina"},
		},
	})
	fs.json(http.MethodGet, "/data/wow/realm/whitemane?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "PVP"}, "category": "PvP"})
	fs.json(http.MethodGet, "/data/wow/realm/faerlina?namespace=dynamic-classic1x-us", http.StatusOK,
		map[string]any{"type": map[string]string{"type": "NORMAL"}, "category": "PvE"})

	c := newTestClient(fs)
	realms, err := c.Realms(context.Background(), "us")
	if err != nil {
		t.Fatal(err)
	}
	if len(realms) != 2 || realms[0].Slug != "whitemane" || realms[0].Type != "PVP" {
		t.Fatalf("realms = %+v", realms)
	}

	// A second call must hit the cache, not the index endpoint again.
	if _, err := c.Realms(context.Background(), "us"); err != nil {
		t.Fatal(err)
	}
	if n := fs.count(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-us"); n != 1 {
		t.Fatalf("realm index called %d times, want 1 (cached)", n)
	}
}
