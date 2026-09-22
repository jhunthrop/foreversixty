// api/internal/bnetapi/guild_test.go
package bnetapi

import (
	"context"
	"net/http"
	"testing"
)

func TestGuildSlug(t *testing.T) {
	cases := map[string]string{
		"Iron Vanguard":  "iron-vanguard",
		"Ashes & Embers": "ashes--embers",
		"légion":         "lgion",
	}
	for in, want := range cases {
		if got := GuildSlug(in); got != want {
			t.Errorf("GuildSlug(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestGuildRoster(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.handlers[http.MethodGet+" /data/wow/guild/whitemane/iron-vanguard/roster?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"guild":{"name":"Iron Vanguard","id":12,"realm":{"slug":"whitemane"}},
			"members":[
				{"character":{"name":"Guildmaster","realm":{"slug":"whitemane"},"level":60,"playable_class":{"id":1}},"rank":0},
				{"character":{"name":"Thoradin","realm":{"slug":"whitemane"},"level":42,"playable_class":{"id":1}},"rank":1}
			]}`))
	}
	c := newTestClient(fs)
	roster, err := c.GuildRoster(context.Background(), "us", "whitemane", "Iron Vanguard")
	if err != nil {
		t.Fatal(err)
	}
	if roster.GuildID != 12 || roster.GuildName != "Iron Vanguard" || len(roster.Members) != 2 {
		t.Fatalf("roster = %+v", roster)
	}
	if roster.Members[0].Rank != 0 || roster.Members[1].Name != "Thoradin" {
		t.Fatalf("members = %+v", roster.Members)
	}
}
