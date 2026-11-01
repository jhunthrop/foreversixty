// api/internal/guilds/members.go
package guilds

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// The four states a roster entry takes. The page writes them as "Approved", "Waiting for
// approval", "No account yet" and "Not seen in game yet"; the API sends the codes so the
// words stay the page's.
const (
	// EntryApproved is a verified character, whether an officer approved it, Battle.net
	// confirmed it or an invite link brought its account in.
	EntryApproved = "approved"
	// EntryWaiting is an account's character that no officer has approved yet.
	EntryWaiting = "waiting"
	// EntryNoAccount is a character the addon's snapshot names that no account has claimed.
	EntryNoAccount = "no_account"
	// EntryNotSeen is a member verified by invite whom the snapshot does not name: they
	// joined, but the addon has not seen them in the guild yet.
	EntryNotSeen = "not_seen"
)

// RosterEntry is one person on the guild's roster as the Roster tab lists them: one entry
// per character the newest addon snapshot names and per verified character, merged by name.
// CharacterKey is set when the entry has a guild_characters row (the page joins it to the
// detailed roster[] row with it) and empty for an in-game-only name or an invited member.
// Rank is the guild rank word: the addon's for an account's character, "member" for an entry
// the snapshot carries without a rank. Level, Spec and SeenAt are the snapshot's line, as the
// member's addon announced them (Level and Spec are nil when the member withholds them).
type RosterEntry struct {
	Name         string  `json:"name"`
	CharacterKey string  `json:"character_key,omitempty"`
	State        string  `json:"state"`
	Rank         string  `json:"rank"`
	Class        *string `json:"class,omitempty"`
	Level        *int    `json:"level,omitempty"`
	Spec         *string `json:"spec,omitempty"`
	SeenAt       string  `json:"seen_at,omitempty"`
	// Role, SpecName ("Holy Paladin") and SpecIcon are the catalogue's words for Spec, else
	// for the entry's roster row's spec (describeMembers); null when neither resolves.
	Role     *string `json:"role"`
	SpecName *string `json:"spec_name"`
	SpecIcon *string `json:"spec_icon"`
}

// RosterCounts are the numbers the count line above the roster reads: everyone on it, and
// how many of them are in game without an account, waiting for approval, or not seen in
// game yet. The approved are the rest.
type RosterCounts struct {
	OnRoster        int `json:"on_roster"`
	InGameNoAccount int `json:"in_game_no_account"`
	Waiting         int `json:"waiting"`
	NotSeen         int `json:"not_seen"`
}

// invitedMember is an account that joined by invite link: a verified synthetic row and no
// character of the guild's own yet. Name is the account's main (or first) character, or its
// public name when it has none; Class is that character's class when the account shares gear.
type invitedMember struct {
	UserID int64
	Name   string
	Class  *string
	// Gear is the account's consent to show gear: without it the snapshot's level and spec
	// stay off its entry, the way a character row's do.
	Gear bool
}

// mergeRoster builds the roster entries and their counts from the three sources: the
// guild's character rows (verified and waiting both), the invited accounts, and the newest
// addon snapshot's member lines (nil when the guild has none). Names match case-insensitively
// and a snapshot name's realm half is dropped, the way homeAddonSnapshot matches them. A row
// the snapshot also names carries the snapshot's line (its Addon); an invited member the
// snapshot names is approved and seen, one it does not is "not seen"; a snapshot name matching
// nobody is "no account", listed by name after the rest. Pure, so the table test needs no
// database.
func mergeRoster(rows []RosterRow, invited []invitedMember, snapshot []snapshotMember) ([]RosterEntry, RosterCounts) {
	lines := make(map[string]snapshotMember, len(snapshot))
	order := make([]string, 0, len(snapshot))
	for _, m := range snapshot {
		name, _ := character.SplitUnit(m.Name)
		key := strings.ToLower(name)
		if key == "" {
			continue
		}
		if _, seen := lines[key]; seen {
			continue
		}
		lines[key] = m
		order = append(order, key)
	}
	sort.Strings(order)

	named := make(map[string]bool, len(rows)+len(invited)+len(order))
	entries := make([]RosterEntry, 0, len(rows)+len(invited)+len(order))
	var counts RosterCounts
	for _, row := range rows {
		named[strings.ToLower(row.Name)] = true
		e := RosterEntry{Name: row.Name, CharacterKey: row.CharacterKey, State: EntryApproved, Rank: row.Rank, Class: row.Class}
		if !row.Verified {
			e.State = EntryWaiting
			counts.Waiting++
		}
		if row.Addon != nil {
			e.Level, e.Spec, e.SeenAt = row.Addon.Level, row.Addon.Spec, row.Addon.SeenAt
		}
		entries = append(entries, e)
	}
	for _, inv := range invited {
		key := strings.ToLower(inv.Name)
		if named[key] {
			continue // the same person already has a character row
		}
		named[key] = true
		e := RosterEntry{Name: inv.Name, State: EntryNotSeen, Rank: "member", Class: inv.Class}
		if m, ok := lines[key]; ok {
			e.State = EntryApproved
			e.Level, e.Spec, e.SeenAt = snapshotFacts(m)
			if !inv.Gear {
				e.Level, e.Spec = nil, nil
			}
		} else {
			counts.NotSeen++
		}
		entries = append(entries, e)
	}
	for _, key := range order {
		if named[key] {
			continue
		}
		m := lines[key]
		name, _ := character.SplitUnit(m.Name)
		e := RosterEntry{Name: name, State: EntryNoAccount, Rank: "member"}
		e.Level, e.Spec, e.SeenAt = snapshotFacts(m)
		entries = append(entries, e)
		counts.InGameNoAccount++
	}
	counts.OnRoster = len(entries)
	return entries, counts
}

// snapshotFacts is a snapshot line as an entry carries it: the level and the spec slug (nil
// when the line announced none) and the sender's local seen_at.
func snapshotFacts(m snapshotMember) (level *int, spec *string, seenAt string) {
	if m.Level > 0 {
		l := m.Level
		level = &l
	}
	if m.Spec != "" {
		sp := m.Spec
		spec = &sp
	}
	return level, spec, m.SeenAt
}

// invitedMembers lists the accounts that joined guildID by invite and whose own character is
// not already a row of the guild's: verified synthetic rows, each named by the account's main
// character (else its first by name), else by public name.
func (s *Store) invitedMembers(ctx context.Context, guildID int64) ([]invitedMember, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.user_id, c.name, case when gm.consent in ('gear', 'gear_bags') then c.class end,
		       gm.consent in ('gear', 'gear_bags')
		from guild_characters gc
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		join users u on u.id = gc.user_id
		left join lateral (
		  select c.key, c.name, c.class from characters c where c.user_id = gc.user_id
		  order by (c.key = u.main_character_key) desc, c.name limit 1
		) c on true
		where gc.guild_id = $1 and gc.character_key like 'account:%' and gc.verified_at is not null
		  and not exists (select 1 from guild_characters x where x.guild_id = gc.guild_id and x.character_key = c.key)
		order by gc.refreshed_at, gc.user_id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: invited members: %w", err)
	}
	defer rows.Close()
	out := []invitedMember{}
	for rows.Next() {
		var m invitedMember
		var name *string
		if err := rows.Scan(&m.UserID, &name, &m.Class, &m.Gear); err != nil {
			return nil, fmt.Errorf("guilds: invited members: %w", err)
		}
		m.Name = s.inviteeName(ctx, m.UserID, name)
		out = append(out, m)
	}
	return out, rows.Err()
}

// inviteeName names an invited account: its main (or first) character's name, else its public
// name, else a stand-in. The one rule the home's roster and the /v1/me roster count share.
func (s *Store) inviteeName(ctx context.Context, userID int64, characterName *string) string {
	if characterName != nil && *characterName != "" {
		return *characterName
	}
	if s.Accounts != nil {
		if u, err := s.Accounts.User(ctx, userID); err == nil && u.PublicName() != "" {
			return u.PublicName()
		}
	}
	return fmt.Sprintf("user-%d", userID)
}

// invitedRosterRows are the invited accounts' picked characters as full roster rows (gear,
// class, attendance), for the loot council, which ranks every member with gear on the site:
// an account that joined by invite is a candidate through its main (else its newest) character
// that has an export, though no guild_characters row of the guild's names it. The row's
// CharacterKey is that character's own key (attendance and fight evidence are keyed by it);
// awardKey is the synthetic account key, the guild_characters row an award has to name.
func (s *Store) invitedRosterRows(ctx context.Context, guildID int64, g Guild) ([]RosterRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select ae.character_key, gc.user_id, ae.region, ae.ruleset, ae.name, gc.rank, gc.verified_at is not null,
		       ae.updated_at >= now() - interval '24 hours', ae.updated_at, gm.consent,
		       case when gm.consent in ('gear', 'gear_bags') then fm.class end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.spec end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.role end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.ilvl end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.faction end,
		       case when gm.consent in ('gear', 'gear_bags') then ae.export end,
		       gc.character_key
		from guild_characters gc
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		join users u on u.id = gc.user_id
		join lateral (
		  select x.character_key from characters c join addon_exports x on x.character_key = c.key
		  where c.user_id = gc.user_id
		  order by (c.key = u.main_character_key) desc, x.updated_at desc limit 1
		) pick on true
		join addon_exports ae on ae.character_key = pick.character_key
		left join lateral (
		  select m.class, m.spec, m.role, m.ilvl, m.faction
		  from fight_metrics m join reports r on r.id = m.report_id
		  where m.player_key = ae.character_key and `+presence.EvidenceReport("r")+`
		  order by m.fought_at desc limit 1
		) fm on true
		where gc.guild_id = $1 and gc.character_key like 'account:%' and gc.verified_at is not null
		  and not exists (select 1 from guild_characters x
		                  where x.guild_id = gc.guild_id and x.character_key = pick.character_key)
		order by gc.user_id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: invited roster rows: %w", err)
	}
	defer rows.Close()
	out := []RosterRow{}
	for rows.Next() {
		var awardKey string
		row, err := s.scanRosterRow(rows, g, 0, false, false, &awardKey)
		if err != nil {
			return nil, err
		}
		row.awardKey = awardKey
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	source, err := s.attendanceSource(ctx, guildID)
	if err != nil {
		return nil, err
	}
	if err := s.enrichRoster(ctx, guildID, source, out); err != nil {
		return nil, err
	}
	return out, nil
}
