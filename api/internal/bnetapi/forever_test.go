// api/internal/bnetapi/forever_test.go
package bnetapi

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

// Fixtures under testdata/forever/ are live responses: the realm index (trimmed
// to three realms), realm details and the static class index, probed
// 2026-11-03, and account-character-live.json, one entry from a real sign-in
// on 2026-11-04 with its identity scrubbed. That entry has a surname, no realm,
// and the realm id only in the protected-character link; its public profile
// answered 404. account-profile-modeled.json was written before any user
// token was available and matches that live shape, with fictional characters.

func fixtureBytes(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile("testdata/forever/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func (fs *fixtureServer) raw(method, path string, body []byte) {
	fs.handlers[method+" "+path] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}
}

func foreverClient(fs *fixtureServer) *Client {
	c := newTestClient(fs)
	c.Game = "forever"
	return c
}

func TestAccountCharactersReadsForeverSurnamesAndRealmIDs(t *testing.T) {
	fs := newFixtureServer(t)
	fs.raw(http.MethodGet, "/profile/user/wow?namespace=profile-forever-us", fixtureBytes(t, "account-profile-modeled.json"))
	chars, err := foreverClient(fs).AccountCharacters(context.Background(), "us", "user-tok")
	if err != nil {
		t.Fatal(err)
	}
	if len(chars) != 3 {
		t.Fatalf("chars = %d, want 3", len(chars))
	}
	first := chars[0]
	if first.Name != "Elyra Duskvale" || first.ID != 7001 || first.RealmID != 6798 || first.RealmSlug != "" {
		t.Fatalf("first = %+v, want the full name, the id and the protected link's realm id, no realm slug", first)
	}
	if chars[1].RealmID != 6928 || chars[1].Faction != "horde" || chars[1].ClassSlug != "warrior" {
		t.Fatalf("second = %+v", chars[1])
	}
}

// TestAccountCharactersReadsTheLiveForeverEntry parses the live entry shape
// (testdata/forever/account-character-live.json, a real sign-in, scrubbed):
// a Skyborne race, level 1, the realm only in the protected-character link.
func TestAccountCharactersReadsTheLiveForeverEntry(t *testing.T) {
	var live struct {
		Entry json.RawMessage `json:"entry"`
	}
	if err := json.Unmarshal(fixtureBytes(t, "account-character-live.json"), &live); err != nil {
		t.Fatal(err)
	}
	fs := newFixtureServer(t)
	fs.raw(http.MethodGet, "/profile/user/wow?namespace=profile-forever-us",
		[]byte(`{"wow_accounts":[{"id":1,"characters":[`+string(live.Entry)+`]}]}`))
	chars, err := foreverClient(fs).AccountCharacters(context.Background(), "us", "user-tok")
	if err != nil {
		t.Fatal(err)
	}
	if len(chars) != 1 {
		t.Fatalf("chars = %d", len(chars))
	}
	c := chars[0]
	if c.Name != "Elyra Duskvale" || c.ID != 7001 || c.RealmID != 6800 || c.RealmSlug != "" ||
		c.RaceName != "High Order Skyborne" || c.ClassSlug != "hunter" || c.Level != 1 || c.Faction != "alliance" {
		t.Fatalf("char = %+v", c)
	}
}

func TestForeverCharacterResourcesUseTheNameSurnamePath(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.raw(http.MethodGet, "/profile/wow/character/elyra-duskvale?namespace=profile-forever-us",
		[]byte(`{"name":"Elyra","surname":"Duskvale","level":60,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Mage"},"race":{"name":"Human"},"guild":{"name":"Quiet Hours","id":3}}`))
	fs.raw(http.MethodGet, "/profile/wow/character/elyra-duskvale/equipment?namespace=profile-forever-us", []byte(`{}`))
	fs.raw(http.MethodGet, "/profile/wow/character/elyra-duskvale/specializations?namespace=profile-forever-us", []byte(`{}`))
	fs.raw(http.MethodGet, "/profile/wow/character/elyra-duskvale/character-media?namespace=profile-forever-us", []byte(`{"assets":[]}`))
	c := foreverClient(fs)
	ctx := context.Background()

	p, _, err := c.Character(ctx, "us", "", "Elyra Duskvale")
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != "Elyra Duskvale" || !p.HasGuild || p.GuildRealmSlug != "" || p.RealmSlug != "" {
		t.Fatalf("profile = %+v", p)
	}
	if _, err := c.Equipment(ctx, "us", "", "Elyra Duskvale"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.Specializations(ctx, "us", "", "Elyra Duskvale"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := c.CharacterMedia(ctx, "us", "", "Elyra Duskvale"); err != nil {
		t.Fatal(err)
	}
}

func TestForeverGuildRosterIsRealmlessAndCarriesIDs(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.raw(http.MethodGet, "/data/wow/guild/quiet-hours/roster?namespace=profile-forever-us",
		[]byte(`{"guild":{"name":"Quiet Hours","id":3},"members":[
			{"character":{"id":7001,"name":"Elyra","surname":"Duskvale","level":60,"playable_class":{"id":8}},"rank":0}]}`))
	roster, err := foreverClient(fs).GuildRoster(context.Background(), "us", "", "Quiet Hours")
	if err != nil {
		t.Fatal(err)
	}
	if len(roster.Members) != 1 || roster.Members[0].ID != 7001 || roster.Members[0].Name != "Elyra Duskvale" {
		t.Fatalf("roster = %+v", roster)
	}
}

func TestForeverRealmsResolveRulesets(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.raw(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-forever-us", fixtureBytes(t, "realm-index-us.json"))
	fs.raw(http.MethodGet, "/data/wow/realm/camwow-us1-realm-1-pve?namespace=dynamic-forever-us",
		[]byte(`{"id":6794,"name":"CAMWOW US1 Realm 1 PvE","category":"United States","type":{"type":"NORMAL"}}`))
	fs.raw(http.MethodGet, "/data/wow/realm/camwow-us1-realm-2-pvp?namespace=dynamic-forever-us", fixtureBytes(t, "realm-pvp-us.json"))
	fs.raw(http.MethodGet, "/data/wow/realm/camwow-us1-realm-10-rp?namespace=dynamic-forever-us", fixtureBytes(t, "realm-rp-us.json"))
	realms, err := foreverClient(fs).Realms(context.Background(), "us")
	if err != nil {
		t.Fatal(err)
	}
	want := map[int64]string{6794: "normal", 6798: "pvp", 6928: "rp"}
	for _, r := range realms {
		if got := RulesetOf(r); got != want[r.ID] {
			t.Errorf("realm %d (%s) ruleset = %q, want %q", r.ID, r.Slug, got, want[r.ID])
		}
	}
	if len(realms) != len(want) {
		t.Fatalf("realms = %d, want %d", len(realms), len(want))
	}
}

func TestStaticNamespaceIsReadFromLinks(t *testing.T) {
	if got := StaticNamespace(fixtureBytes(t, "playable-class-index-us.json")); got != "static-1.60.1_70399-forever-us" {
		t.Fatalf("StaticNamespace = %q", got)
	}
	if got := StaticNamespace(fixtureBytes(t, "realm-pvp-us.json")); got != "" {
		t.Fatalf("a dynamic-only body yielded %q", got)
	}
}

func TestCharacterSlugAndFullName(t *testing.T) {
	if got := FullName(" Bow ", "Jackzon"); got != "Bow Jackzon" {
		t.Fatalf("FullName = %q", got)
	}
	if got := FullName("Thoradin", ""); got != "Thoradin" {
		t.Fatalf("FullName without surname = %q", got)
	}
	if got := CharacterSlug("Bow  Jackzon"); got != "bow-jackzon" {
		t.Fatalf("CharacterSlug = %q", got)
	}
}
