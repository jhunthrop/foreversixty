// api/internal/bnetapi/profile.go
package bnetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// AccountCharacter is one character from GET /profile/user/wow — the
// account's own roster, readable only with the user's OAuth token.
type AccountCharacter struct {
	ID int64
	// Name is the character's full display name. A Forever character answers
	// a first name and a separate surname, joined here as "Name Surname" (the
	// form the addon export and the combat log already carry); a character
	// from a realmed game has no surname and Name is the bare name.
	Name string
	// RealmSlug and RealmName are empty for a Forever character: Forever's
	// account profile names no realm. RealmID is the realm the protected-
	// character link carries (".../protected-character/{realmId}-{characterId}"),
	// or realm.id when Blizzard sends one; it is what resolves the ruleset.
	RealmSlug string
	RealmName string
	RealmID   int64
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
	Name    string `json:"name"`
	Surname string `json:"surname"`
	ID      int64  `json:"id"`
	Realm   struct {
		ID   int64  `json:"id"`
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"realm"`
	ProtectedCharacter struct {
		Href string `json:"href"`
	} `json:"protected_character"`
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
			realmID := fields.Realm.ID
			if realmID == 0 {
				realmID = protectedRealmID(fields.ProtectedCharacter.Href)
			}
			out = append(out, AccountCharacter{
				ID: fields.ID, Name: FullName(fields.Name, fields.Surname),
				RealmSlug: fields.Realm.Slug, RealmName: fields.Realm.Name, RealmID: realmID,
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
	// ID is Blizzard's character id, the key the character directory holds it by.
	ID        int64
	Name      string
	Level     int
	Faction   string
	ClassSlug string
	// RaceName is Blizzard's race.name, verbatim (title case, e.g. "Night Elf") — the same
	// string data/builds/<build>/races.json keys its own race rows by (spec
	// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2).
	RaceName  string
	RealmSlug string
	RealmName string
	// Ruleset is the site ruleset the profile's own ruleset names ("PvP" is
	// pvp), empty when it names none. A Forever profile carries it; it is
	// the one source left now Forever's realm index answers 404.
	Ruleset   string
	GuildName string
	// GuildRealmSlug is the guild's own realm, empty when Blizzard names
	// none (a Forever guild, like a Forever character, may be realmless).
	GuildRealmSlug string
	// GuildID and GuildSlug are the guild's Blizzard id and URL slug, when it has one.
	GuildID            int64
	GuildSlug          string
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
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
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
	Ruleset struct {
		Name string `json:"name"`
	} `json:"ruleset"`
	Guild *struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Slug  string `json:"slug"`
		Realm struct {
			Slug string `json:"slug"`
		} `json:"realm"`
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
	u := c.characterURL(region, realmSlug, name, "")
	body, err := c.getBytes(ctx, "character", u, token)
	if err != nil {
		return CharacterProfile{}, nil, err
	}
	var res characterProfileResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return CharacterProfile{}, nil, fmt.Errorf("bnetapi: character: decode: %w", err)
	}
	p := CharacterProfile{
		ID: res.ID, Name: FullName(res.Name, res.Surname), Level: res.Level, Faction: strings.ToLower(res.Faction.Type),
		ClassSlug: classSlug(res.CharacterClass.Name), RaceName: res.Race.Name,
		RealmSlug: res.Realm.Slug, RealmName: res.Realm.Name,
		LastLoginTimestamp: res.LastLoginTimestamp,
		AverageItemLevel:   res.AverageItemLevel, EquippedItemLevel: res.EquippedItemLevel,
	}
	if ruleset, ok := character.RulesetNamed(res.Ruleset.Name); ok {
		p.Ruleset = ruleset
	}
	if res.Guild != nil {
		p.GuildName = res.Guild.Name
		p.GuildRealmSlug = res.Guild.Realm.Slug
		p.GuildID, p.GuildSlug = res.Guild.ID, res.Guild.Slug
		p.HasGuild = true
	}
	return p, json.RawMessage(body), nil
}

// Equipment reads a character's equipped-items snapshot verbatim. Not
// parsed here — the planner's Blizzard-to-FS1 gear mapping is a later
// lane (spec §B) — so the caller stores the raw body directly.
func (c *Client) Equipment(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	return c.characterResource(ctx, "equipment", region, realmSlug, name)
}

// Specializations reads a character's talent groups verbatim. Not parsed here — the
// Blizzard-to-FS1 talent mapping is bnetbuild's job (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.2) — so the caller stores
// the raw body directly, the same shape Equipment already returns.
func (c *Client) Specializations(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	return c.characterResource(ctx, "specializations", region, realmSlug, name)
}

// Statistics reads a character's statistics verbatim: each primary stat's base and
// effective value, health, power, and the rest. The only place a Forever race's base stats
// (Skyborne) can be read from the game.
func (c *Client) Statistics(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	return c.characterResource(ctx, "statistics", region, realmSlug, name)
}

// characterResource reads one character sub-resource ("equipment", ...) verbatim.
func (c *Client) characterResource(ctx context.Context, resource, region, realmSlug, name string) (json.RawMessage, error) {
	token, err := c.AppToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("bnetapi: %s: %w", resource, err)
	}
	u := c.characterURL(region, realmSlug, name, "/"+resource)
	body, err := c.getBytes(ctx, resource, u, token)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(body), nil
}

// FullName joins a Forever character's first name and surname into the one
// display name the site uses everywhere ("Bow Jackzon"). A character with no
// surname keeps its bare name.
func FullName(name, surname string) string {
	name, surname = strings.TrimSpace(name), strings.TrimSpace(surname)
	if surname == "" {
		return name
	}
	return name + " " + surname
}

// CharacterSlug is the profile path's character segment: the full name
// lowercased with its words joined by hyphens, so "Bow Jackzon" reads
// "bow-jackzon", the {name}-{surname} form Forever's profile links use.
func CharacterSlug(name string) string {
	return strings.Join(strings.Fields(strings.ToLower(name)), "-")
}

// characterURL is one character's profile resource: a realmed game answers
// /profile/wow/character/{realm}/{name}, Forever (no realm) answers
// /profile/wow/character/{name}-{surname}. sub is "" for the summary or a
// sub-resource such as "/equipment".
func (c *Client) characterURL(region, realmSlug, name, sub string) string {
	path := "/profile/wow/character/"
	if realmSlug != "" {
		path += url.PathEscape(realmSlug) + "/"
	}
	return c.APIHost(region) + path + url.PathEscape(CharacterSlug(name)) + sub +
		"?namespace=" + c.ProfileNamespace(region)
}

// protectedCharacterRe reads the realm id out of an account profile's
// protected-character link, ".../protected-character/{realmId}-{characterId}".
var protectedCharacterRe = regexp.MustCompile(`/protected-character/(\d+)-\d+`)

// protectedRealmID is the realm id a protected-character link carries, 0
// when the link is absent or unreadable.
func protectedRealmID(href string) int64 {
	m := protectedCharacterRe.FindStringSubmatch(href)
	if m == nil {
		return 0
	}
	id, err := strconv.ParseInt(m[1], 10, 64)
	if err != nil {
		return 0
	}
	return id
}
