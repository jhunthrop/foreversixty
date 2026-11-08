// api/internal/guilds/raids_not_signed.go
//
// Who has not answered a raid, as an officer reads it (spec 2026-11-05-web-layout §7.1,
// D16): one row per account, the account's raider for this raid (its main, else its highest
// character), at or above the zone's level, with whether the account has linked Discord. The
// raid view's not_signed (every verified character of an account that has not answered) stays
// as it is for members; this list is the officer's Overview count and the ping's mentions.
package guilds

import (
	"context"
	"fmt"
	"sort"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// NotSignedRaider is one account that has not answered, as the officer sees it.
type NotSignedRaider struct {
	CharacterKey  string `json:"character_key"`
	Name          string `json:"name"`
	Class         string `json:"class"`
	AccountKey    string `json:"account_key"`
	DiscordLinked bool   `json:"discord_linked"`
	// DiscordDMURL opens a Discord DM with the member, only when they opted into officer DMs
	// (notification settings officer_dms) and linked Discord; null otherwise, so no id shows for
	// anyone who did not opt in.
	DiscordDMURL *string `json:"discord_dm_url"`
	// Role is the row's sign-up role (tank, healer, dps), from its roster group.
	Role string `json:"role"`
	// LastWeek is the account's answer for the raid the week before (raids_last_week.go): in,
	// late, bench, out or missed; null with no such raid. Officers' rows only fill it.
	LastWeek *string `json:"last_week"`
	// discordUserID is the linked Discord user the ping mentions; userID the account; never
	// marshaled.
	discordUserID string
	userID        int64
}

// rosterFact is what the officer lists read of one verified character beyond the raid roster.
type rosterFact struct {
	userID        int64
	isMain        bool
	level         *int
	discordUserID string
	officerDMs    bool
}

// officerRosterFacts reads every verified character of the guild: its account, whether it is
// the account's main, its level (from the export, at gear consent; nil otherwise or when the
// export carries none) and the account's linked Discord user ("" when none).
func (s *Store) officerRosterFacts(ctx context.Context, guildID int64) (map[string]rosterFact, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, gc.user_id, coalesce(u.main_character_key = gc.character_key, false),
		       case when gm.consent in ('gear', 'gear_bags') then ae.export end,
		       coalesce(l.discord_user_id, ''), coalesce(ns.officer_dms, false)
		from guild_characters gc
		join users u on u.id = gc.user_id
		left join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		left join addon_exports ae on ae.character_key = gc.character_key
		left join discord_links l on l.user_id = gc.user_id
		left join notification_settings ns on ns.user_id = gc.user_id
		where gc.guild_id = $1 and gc.verified_at is not null and gc.character_key not like 'account:%'`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: officer roster facts: %w", err)
	}
	defer rows.Close()
	out := map[string]rosterFact{}
	for rows.Next() {
		var key string
		var f rosterFact
		var export *string
		if err := rows.Scan(&key, &f.userID, &f.isMain, &export, &f.discordUserID, &f.officerDMs); err != nil {
			return nil, fmt.Errorf("guilds: officer roster facts: %w", err)
		}
		f.level = exportLevel(export)
		out[key] = f
	}
	return out, rows.Err()
}

// exportLevel is the export's level= section, nil without one.
func exportLevel(export *string) *int {
	if export == nil {
		return nil
	}
	d, ok := fs1.Decode(*export)
	if !ok || !d.HasLevel {
		return nil
	}
	level := d.Level
	return &level
}

// zoneMinLevel is the level a raider needs for the zone: its own minimum, else raidLevel for a
// raid; 0 (anyone) for a guild event.
func zoneMinLevel(z Zone) int {
	switch {
	case z.MinLevel != nil:
		return *z.MinLevel
	case z.IsRaid():
		return raidLevel
	}
	return 0
}

// notSignedRaiders is one row per account in v.NotSigned: the account's main when it is among
// them, else its highest known level, else the first by name; an account whose every
// character has a known level under the zone's is left out (an unknown level counts it in).
func notSignedRaiders(v RaidView, facts map[string]rosterFact) []NotSignedRaider {
	minLevel := zoneMinLevel(zoneOf(v.Zone))
	best := map[int64]RaidCharacter{}
	for _, c := range v.NotSigned {
		f, ok := facts[c.CharacterKey]
		if !ok || (f.level != nil && *f.level < minLevel) {
			continue
		}
		cur, seen := best[f.userID]
		if !seen || betterRaider(c, f, cur, facts[cur.CharacterKey]) {
			best[f.userID] = c
		}
	}
	out := make([]NotSignedRaider, 0, len(best))
	for uid, c := range best {
		f := facts[c.CharacterKey]
		out = append(out, NotSignedRaider{
			CharacterKey: c.CharacterKey, Name: c.Name, Class: c.Class, AccountKey: accountKeyFor(uid),
			DiscordLinked: f.discordUserID != "", DiscordDMURL: discordDMURL(f), Role: groupRoles[c.Group],
			discordUserID: f.discordUserID, userID: uid,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// discordDMURLPrefix opens a Discord user's profile, where Message starts a DM.
const discordDMURLPrefix = "https://discord.com/users/"

// discordDMURL is the member's DM link when they opted in and linked a valid Discord id.
func discordDMURL(f rosterFact) *string {
	if !f.officerDMs || !discordSnowflake.MatchString(f.discordUserID) {
		return nil
	}
	u := discordDMURLPrefix + f.discordUserID
	return &u
}

// betterRaider reports whether a (with facts fa) is a better pick for its account than b.
func betterRaider(a RaidCharacter, fa rosterFact, b RaidCharacter, fb rosterFact) bool {
	if fa.isMain != fb.isMain {
		return fa.isMain
	}
	la, lb := levelOr0(fa.level), levelOr0(fb.level)
	if la != lb {
		return la > lb
	}
	return a.Name < b.Name
}

func levelOr0(level *int) int {
	if level == nil {
		return 0
	}
	return *level
}

// raiderNames is the rows' names, in order.
func raiderNames(rows []NotSignedRaider) []string {
	out := make([]string, len(rows))
	for i, r := range rows {
		out[i] = r.Name
	}
	return out
}
