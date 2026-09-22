// api/internal/bnetapi/profile.go
package bnetapi

import (
	"context"
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
	Level     int
	Faction   string // lowercase: "alliance" | "horde"
}

type accountCharactersResponse struct {
	WowAccounts []struct {
		Characters []struct {
			Name  string `json:"name"`
			ID    int64  `json:"id"`
			Realm struct {
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"realm"`
			PlayableClass struct {
				Name string `json:"name"`
			} `json:"playable_class"`
			Faction struct {
				Type string `json:"type"`
			} `json:"faction"`
			Level int `json:"level"`
		} `json:"characters"`
	} `json:"wow_accounts"`
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
		for _, ch := range acct.Characters {
			out = append(out, AccountCharacter{
				ID: ch.ID, Name: ch.Name,
				RealmSlug: ch.Realm.Slug, RealmName: ch.Realm.Name,
				ClassSlug: classSlug(ch.PlayableClass.Name),
				Level:     ch.Level, Faction: strings.ToLower(ch.Faction.Type),
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
	RealmSlug          string
	RealmName          string
	GuildName          string
	HasGuild           bool
	LastLoginTimestamp int64
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
	Realm struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"realm"`
	Guild *struct {
		Name string `json:"name"`
	} `json:"guild"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
}

// Character reads one character's public profile. name is lowercased
// before it reaches the URL, per Blizzard's own rule for this endpoint.
func (c *Client) Character(ctx context.Context, region, realmSlug, name string) (CharacterProfile, error) {
	u := c.APIHost(region) + "/profile/wow/character/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(strings.ToLower(name)) + "?namespace=" + c.ProfileNamespace(region)
	var res characterProfileResponse
	if err := c.getJSON(ctx, "character", u, &res); err != nil {
		return CharacterProfile{}, err
	}
	p := CharacterProfile{
		Name: res.Name, Level: res.Level, Faction: strings.ToLower(res.Faction.Type),
		ClassSlug: classSlug(res.CharacterClass.Name),
		RealmSlug: res.Realm.Slug, RealmName: res.Realm.Name,
		LastLoginTimestamp: res.LastLoginTimestamp,
	}
	if res.Guild != nil {
		p.GuildName = res.Guild.Name
		p.HasGuild = true
	}
	return p, nil
}
