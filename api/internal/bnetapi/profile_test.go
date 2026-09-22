// api/internal/bnetapi/profile_test.go
package bnetapi

import (
	"context"
	"net/http"
	"testing"
)

func TestAccountCharactersUsesTheUsersOwnToken(t *testing.T) {
	fs := newFixtureServer(t)
	var sawAuth string
	fs.handlers[http.MethodGet+" /profile/user/wow?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		sawAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"wow_accounts":[{"id":1,"characters":[
			{"name":"Thoradin","id":501,"realm":{"slug":"whitemane","name":"Whitemane"},
			 "playable_class":{"name":"Warrior"},"faction":{"type":"ALLIANCE"},"level":42}
		]}]}`))
	}
	c := newTestClient(fs)
	chars, err := c.AccountCharacters(context.Background(), "us", "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if sawAuth != "Bearer user-oauth-token" {
		t.Fatalf("Authorization = %q, want the user token, not an app token", sawAuth)
	}
	if len(chars) != 1 || chars[0].Name != "Thoradin" || chars[0].ClassSlug != "warrior" ||
		chars[0].Faction != "alliance" || chars[0].Level != 42 || chars[0].RealmSlug != "whitemane" {
		t.Fatalf("chars = %+v", chars)
	}
}

func TestCharacterReadsGuildedAndUnguilded(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.handlers[http.MethodGet+" /profile/wow/character/whitemane/thoradin?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Thoradin","level":42,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Warrior"},"realm":{"slug":"whitemane","name":"Whitemane"},
			"guild":{"name":"Iron Vanguard","id":12},"last_login_timestamp":1700000000000}`))
	}
	fs.handlers[http.MethodGet+" /profile/wow/character/whitemane/loneling?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Loneling","level":10,"faction":{"type":"HORDE"},
			"character_class":{"name":"Mage"},"realm":{"slug":"whitemane","name":"Whitemane"},
			"last_login_timestamp":1700000000000}`))
	}
	c := newTestClient(fs)

	guilded, err := c.Character(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if !guilded.HasGuild || guilded.GuildName != "Iron Vanguard" || guilded.ClassSlug != "warrior" {
		t.Fatalf("guilded = %+v", guilded)
	}

	unguilded, err := c.Character(context.Background(), "us", "whitemane", "Loneling")
	if err != nil {
		t.Fatal(err)
	}
	if unguilded.HasGuild {
		t.Fatalf("unguilded = %+v, want HasGuild false", unguilded)
	}
}
