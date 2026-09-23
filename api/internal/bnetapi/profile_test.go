// api/internal/bnetapi/profile_test.go
package bnetapi

import (
	"context"
	"net/http"
	"strings"
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
			 "playable_class":{"name":"Warrior"},"playable_race":{"name":"Dwarf"},
			 "gender":{"type":"MALE"},"faction":{"type":"ALLIANCE"},"level":42}
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
	if chars[0].RaceName != "Dwarf" {
		t.Fatalf("RaceName = %q, want the verbatim playable_race.name", chars[0].RaceName)
	}
	if chars[0].GenderType != "male" {
		t.Fatalf("GenderType = %q, want gender.type lowercased", chars[0].GenderType)
	}
	if len(chars[0].Raw) == 0 || !strings.Contains(string(chars[0].Raw), "Thoradin") {
		t.Fatalf("Raw = %s, want the character's own verbatim entry", chars[0].Raw)
	}
}

func TestCharacterReadsGuildedAndUnguilded(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.handlers[http.MethodGet+" /profile/wow/character/whitemane/thoradin?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Thoradin","level":42,"faction":{"type":"ALLIANCE"},
			"character_class":{"name":"Warrior"},"race":{"name":"Orc"},
			"realm":{"slug":"whitemane","name":"Whitemane"},
			"guild":{"name":"Iron Vanguard","id":12},"last_login_timestamp":1700000000000,
			"average_item_level":55,"equipped_item_level":54}`))
	}
	fs.handlers[http.MethodGet+" /profile/wow/character/whitemane/loneling?namespace=profile-classic1x-us"] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"Loneling","level":10,"faction":{"type":"HORDE"},
			"character_class":{"name":"Mage"},"realm":{"slug":"whitemane","name":"Whitemane"},
			"last_login_timestamp":1700000000000}`))
	}
	c := newTestClient(fs)

	guilded, rawGuilded, err := c.Character(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if !guilded.HasGuild || guilded.GuildName != "Iron Vanguard" || guilded.ClassSlug != "warrior" {
		t.Fatalf("guilded = %+v", guilded)
	}
	if guilded.AverageItemLevel == nil || *guilded.AverageItemLevel != 55 {
		t.Fatalf("AverageItemLevel = %v, want 55", guilded.AverageItemLevel)
	}
	if guilded.EquippedItemLevel == nil || *guilded.EquippedItemLevel != 54 {
		t.Fatalf("EquippedItemLevel = %v, want 54", guilded.EquippedItemLevel)
	}
	if guilded.RaceName != "Orc" {
		t.Fatalf("RaceName = %q, want the verbatim race.name", guilded.RaceName)
	}
	if len(rawGuilded) == 0 || !strings.Contains(string(rawGuilded), "Iron Vanguard") {
		t.Fatalf("raw = %s, want the verbatim response body", rawGuilded)
	}

	unguilded, _, err := c.Character(context.Background(), "us", "whitemane", "Loneling")
	if err != nil {
		t.Fatal(err)
	}
	if unguilded.HasGuild {
		t.Fatalf("unguilded = %+v, want HasGuild false", unguilded)
	}
	if unguilded.AverageItemLevel != nil || unguilded.EquippedItemLevel != nil {
		t.Fatalf("unguilded item levels = %+v, want both nil (absent from the response)", unguilded)
	}
}

func TestEquipmentReturnsTheRawBody(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/equipment?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"equipped_items": []map[string]any{{"slot": map[string]string{"type": "HEAD"}}}})

	c := newTestClient(fs)
	raw, err := c.Equipment(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "equipped_items") || !strings.Contains(string(raw), "HEAD") {
		t.Fatalf("raw = %s, want the verbatim equipment body", raw)
	}
}

func TestSpecializationsReturnsTheRawBody(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/specializations?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"specialization_groups": []map[string]any{{"is_active": true}}})

	c := newTestClient(fs)
	raw, err := c.Specializations(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "specialization_groups") {
		t.Fatalf("raw = %s, want the verbatim specializations body", raw)
	}
}
