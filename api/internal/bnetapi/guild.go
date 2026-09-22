// api/internal/bnetapi/guild.go
package bnetapi

import (
	"context"
	"net/url"
	"strings"
)

// Roster is GET /data/wow/guild/{realm}/{guild}/roster.
type Roster struct {
	GuildName string
	GuildID   int64
	RealmSlug string
	Members   []RosterMember
}

// RosterMember is one row of a guild roster. Rank 0 is the guild master
// (spec §0).
type RosterMember struct {
	Name      string
	RealmSlug string
	Level     int
	ClassID   int64
	Rank      int
}

type guildRosterResponse struct {
	Guild struct {
		Name  string `json:"name"`
		ID    int64  `json:"id"`
		Realm struct {
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"guild"`
	Members []struct {
		Character struct {
			Name  string `json:"name"`
			Realm struct {
				Slug string `json:"slug"`
			} `json:"realm"`
			Level         int `json:"level"`
			PlayableClass struct {
				ID int64 `json:"id"`
			} `json:"playable_class"`
		} `json:"character"`
		Rank int `json:"rank"`
	} `json:"members"`
}

// GuildRoster reads a guild's public roster.
func (c *Client) GuildRoster(ctx context.Context, region, realmSlug, guildName string) (Roster, error) {
	u := c.APIHost(region) + "/data/wow/guild/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(GuildSlug(guildName)) + "/roster?namespace=" + c.ProfileNamespace(region)
	var res guildRosterResponse
	if err := c.getJSON(ctx, "guild_roster", u, &res); err != nil {
		return Roster{}, err
	}
	roster := Roster{GuildName: res.Guild.Name, GuildID: res.Guild.ID, RealmSlug: res.Guild.Realm.Slug}
	for _, m := range res.Members {
		roster.Members = append(roster.Members, RosterMember{
			Name: m.Character.Name, RealmSlug: m.Character.Realm.Slug,
			Level: m.Character.Level, ClassID: m.Character.PlayableClass.ID, Rank: m.Rank,
		})
	}
	return roster, nil
}

// GuildSlug is Blizzard's own guild-name-to-URL-slug rule (spec §3):
// lowercase, spaces to hyphens, everything outside [a-z0-9-] dropped.
func GuildSlug(name string) string {
	lower := strings.ReplaceAll(strings.ToLower(name), " ", "-")
	var b strings.Builder
	for _, r := range lower {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
