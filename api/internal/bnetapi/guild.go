// api/internal/bnetapi/guild.go
package bnetapi

import (
	"context"
	"net/url"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// Roster is GET /data/wow/guild/{realm}/{guild}/roster.
type Roster struct {
	GuildName string
	GuildID   int64
	RealmSlug string
	Members   []RosterMember
}

// RosterMember is one row of a guild roster. Rank 0 is the guild master
// (spec §0). ID is Blizzard's character id, the match a realmless Forever
// roster is read by; Name is the full "Name Surname" when a surname is sent.
// A Forever roster lists every member, public or not, with race and faction.
type RosterMember struct {
	ID        int64
	Name      string
	RealmSlug string
	Level     int
	ClassID   int64
	RaceID    int64
	// Faction is lowercase, "alliance" or "horde"; empty when the roster names none.
	Faction string
	Rank    int
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
			ID      int64  `json:"id"`
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Realm   struct {
				Slug string `json:"slug"`
			} `json:"realm"`
			Level         int `json:"level"`
			PlayableClass struct {
				ID int64 `json:"id"`
			} `json:"playable_class"`
			PlayableRace struct {
				ID int64 `json:"id"`
			} `json:"playable_race"`
			Faction struct {
				Type string `json:"type"`
			} `json:"faction"`
		} `json:"character"`
		Rank int `json:"rank"`
	} `json:"members"`
}

// GuildRoster reads a guild's public roster. A realmed guild answers
// /data/wow/guild/{realm}/{guild}/roster; with realmSlug empty (a Forever
// guild whose profile names no realm) the realm segment is left out. That
// realmless form is unconfirmed on the live API (2026-11-03): a 404 is soft
// for every caller, so the import simply writes no membership.
func (c *Client) GuildRoster(ctx context.Context, region, realmSlug, guildName string) (Roster, error) {
	return c.GuildRosterBySlug(ctx, region, realmSlug, GuildSlug(guildName))
}

// GuildRosterBySlug is GuildRoster for a guild already named by its slug (the
// character directory keeps guilds by the slug Blizzard answers).
func (c *Client) GuildRosterBySlug(ctx context.Context, region, realmSlug, guildSlug string) (Roster, error) {
	u := c.guildURL(region, realmSlug, guildSlug, "/roster")
	var res guildRosterResponse
	if err := c.getJSON(ctx, "guild_roster", u, &res); err != nil {
		return Roster{}, err
	}
	roster := Roster{GuildName: res.Guild.Name, GuildID: res.Guild.ID, RealmSlug: res.Guild.Realm.Slug}
	for _, m := range res.Members {
		roster.Members = append(roster.Members, RosterMember{
			ID: m.Character.ID, Name: FullName(m.Character.Name, m.Character.Surname), RealmSlug: m.Character.Realm.Slug,
			Level: m.Character.Level, ClassID: m.Character.PlayableClass.ID, RaceID: m.Character.PlayableRace.ID,
			Faction: strings.ToLower(m.Character.Faction.Type), Rank: m.Rank,
		})
	}
	return roster, nil
}

// guildURL is one guild resource: /data/wow/guild/{realm}/{guild}{sub} for a
// realmed game, /data/wow/guild/{guild}{sub} for Forever (realmSlug empty).
func (c *Client) guildURL(region, realmSlug, guildSlug, sub string) string {
	path := "/data/wow/guild/"
	if realmSlug != "" {
		path += url.PathEscape(realmSlug) + "/"
	}
	return c.APIHost(region) + path + url.PathEscape(guildSlug) + sub + "?namespace=" + c.ProfileNamespace(region)
}

// Guild is GET /data/wow/guild/{slug}: a guild's own summary. Ruleset is the
// site ruleset its ruleset names ("PvP" is pvp), empty when it names none; a
// Forever guild's ruleset is every member's, public or not.
type Guild struct {
	ID          int64
	Name        string
	Slug        string
	Faction     string
	Ruleset     string
	MemberCount int
}

type guildResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Faction struct {
		Type string `json:"type"`
	} `json:"faction"`
	Ruleset struct {
		Name string `json:"name"`
	} `json:"ruleset"`
	MemberCount int `json:"member_count"`
}

// GuildSummary reads a guild's public summary by slug (realmSlug empty for Forever).
func (c *Client) GuildSummary(ctx context.Context, region, realmSlug, guildSlug string) (Guild, error) {
	var res guildResponse
	if err := c.getJSON(ctx, "guild", c.guildURL(region, realmSlug, guildSlug, ""), &res); err != nil {
		return Guild{}, err
	}
	g := Guild{ID: res.ID, Name: res.Name, Slug: res.Slug, Faction: strings.ToLower(res.Faction.Type), MemberCount: res.MemberCount}
	if g.Slug == "" {
		g.Slug = guildSlug
	}
	if ruleset, ok := character.RulesetNamed(res.Ruleset.Name); ok {
		g.Ruleset = ruleset
	}
	return g, nil
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
