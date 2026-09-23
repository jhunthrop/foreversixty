// api/internal/bnetapi/profile.go
package bnetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// AccountCharacter is one character from GET /profile/user/wow — the
// account's own roster, readable only with the user's OAuth token.
type AccountCharacter struct {
	ID        int64
	Name      string
	RealmSlug string
	RealmName string
	ClassSlug string
	// RaceName and GenderType are the account-profile entry's own
	// playable_race.name and lowercased gender.type (spec §B); the
	// importer stores them as characters.race/.gender verbatim (race is
	// not lowercased — "Night Elf", not "night elf").
	RaceName   string
	GenderType string
	Level      int
	Faction    string // lowercase: "alliance" | "horde"
	// Raw is this character's own account-profile entry, exactly as
	// Blizzard answered it — stored verbatim as characters.bnet_account
	// (spec §B).
	Raw json.RawMessage
}

type accountCharactersResponse struct {
	WowAccounts []struct {
		Characters []json.RawMessage `json:"characters"`
	} `json:"wow_accounts"`
}

// accountCharacterFields is the typed shape read out of each account-
// profile entry's raw bytes (spec §0's "name, id, realm, playable_class,
// playable_race, gender, faction, level").
type accountCharacterFields struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	Realm struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"realm"`
	PlayableClass struct {
		Name string `json:"name"`
	} `json:"playable_class"`
	PlayableRace struct {
		Name string `json:"name"`
	} `json:"playable_race"`
	Gender struct {
		Type string `json:"type"`
	} `json:"gender"`
	Faction struct {
		Type string `json:"type"`
	} `json:"faction"`
	Level int `json:"level"`
}

// AccountCharacters lists every character on the account behind
// userToken, in region. A region with no WoW account under this
// namespace answers 403/404, which the caller (bnetimport) treats as
// "no account in this region" rather than an error.
func (c *Client) AccountCharacters(ctx context.Context, region, userToken string) ([]AccountCharacter, error) {
	u := c.APIHost(region) + "/profile/user/wow?namespace=" + c.ProfileNamespace(region)
	var res accountCharactersResponse
	if err := c.getJSONWithToken(ctx, "account_characters", u, userToken, &res); err != nil {
		return nil, err
	}
	var out []AccountCharacter
	for _, acct := range res.WowAccounts {
		for _, raw := range acct.Characters {
			var fields accountCharacterFields
			if err := json.Unmarshal(raw, &fields); err != nil {
				return nil, fmt.Errorf("bnetapi: account_characters: decode character: %w", err)
			}
			out = append(out, AccountCharacter{
				ID: fields.ID, Name: fields.Name,
				RealmSlug: fields.Realm.Slug, RealmName: fields.Realm.Name,
				ClassSlug:  classSlug(fields.PlayableClass.Name),
				RaceName:   fields.PlayableRace.Name,
				GenderType: strings.ToLower(fields.Gender.Type),
				Level:      fields.Level, Faction: strings.ToLower(fields.Faction.Type),
				Raw: raw,
			})
		}
	}
	return out, nil
}

// CharacterProfile is GET /profile/wow/character/{realm}/{name} — public
// data, readable with an app token, used both by the initial import and
// by the nightly refresh.
type CharacterProfile struct {
	Name               string
	Level              int
	Faction            string
	ClassSlug          string
	// RaceName is Blizzard's race.name, verbatim (title case, e.g. "Night Elf") — the same
	// string data/builds/<build>/races.json keys its own race rows by (spec
	// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2).
	RaceName           string
	RealmSlug          string
	RealmName          string
	GuildName          string
	HasGuild           bool
	LastLoginTimestamp int64
	// AverageItemLevel and EquippedItemLevel are pointers because the
	// field is simply absent from some responses (a fresh, ungeared
	// character) — nil means "Blizzard didn't say", not "zero" (spec §B,
	// §C: "item_level (equipped) omitted when unknown").
	AverageItemLevel  *int
	EquippedItemLevel *int
}

type characterProfileResponse struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Faction struct {
		Type string `json:"type"`
	} `json:"faction"`
	CharacterClass struct {
		Name string `json:"name"`
	} `json:"character_class"`
	Race struct {
		Name string `json:"name"`
	} `json:"race"`
	Realm struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"realm"`
	Guild *struct {
		Name string `json:"name"`
	} `json:"guild"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
	AverageItemLevel   *int  `json:"average_item_level"`
	EquippedItemLevel  *int  `json:"equipped_item_level"`
}

// Character reads one character's public profile and returns both the
// typed struct and the response's raw body (spec §B: "stored verbatim as
// bnet_profile"), so a caller that wants to capture it never has to
// refetch or re-marshal what it already has. name is lowercased before
// it reaches the URL, per Blizzard's own rule for this endpoint.
func (c *Client) Character(ctx context.Context, region, realmSlug, name string) (CharacterProfile, json.RawMessage, error) {
	token, err := c.AppToken(ctx)
	if err != nil {
		return CharacterProfile{}, nil, fmt.Errorf("bnetapi: character: %w", err)
	}
	u := c.APIHost(region) + "/profile/wow/character/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(strings.ToLower(name)) + "?namespace=" + c.ProfileNamespace(region)
	body, err := c.getBytes(ctx, "character", u, token)
	if err != nil {
		return CharacterProfile{}, nil, err
	}
	var res characterProfileResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return CharacterProfile{}, nil, fmt.Errorf("bnetapi: character: decode: %w", err)
	}
	p := CharacterProfile{
		Name: res.Name, Level: res.Level, Faction: strings.ToLower(res.Faction.Type),
		ClassSlug: classSlug(res.CharacterClass.Name), RaceName: res.Race.Name,
		RealmSlug: res.Realm.Slug, RealmName: res.Realm.Name,
		LastLoginTimestamp: res.LastLoginTimestamp,
		AverageItemLevel:   res.AverageItemLevel, EquippedItemLevel: res.EquippedItemLevel,
	}
	if res.Guild != nil {
		p.GuildName = res.Guild.Name
		p.HasGuild = true
	}
	return p, json.RawMessage(body), nil
}

// Equipment reads a character's equipped-items snapshot verbatim. Not
// parsed here — the planner's Blizzard-to-FS1 gear mapping is a later
// lane (spec §B) — so the caller stores the raw body directly.
func (c *Client) Equipment(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	token, err := c.AppToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("bnetapi: equipment: %w", err)
	}
	u := c.APIHost(region) + "/profile/wow/character/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(strings.ToLower(name)) + "/equipment?namespace=" + c.ProfileNamespace(region)
	body, err := c.getBytes(ctx, "equipment", u, token)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(body), nil
}

// Specializations reads a character's talent groups verbatim. Not parsed here — the
// Blizzard-to-FS1 talent mapping is bnetbuild's job (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2) — so the caller stores
// the raw body directly, the same shape Equipment already returns.
func (c *Client) Specializations(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	token, err := c.AppToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("bnetapi: specializations: %w", err)
	}
	u := c.APIHost(region) + "/profile/wow/character/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(strings.ToLower(name)) + "/specializations?namespace=" + c.ProfileNamespace(region)
	body, err := c.getBytes(ctx, "specializations", u, token)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(body), nil
}
