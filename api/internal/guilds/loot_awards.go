// api/internal/guilds/loot_awards.go
//
// Reading award rows, shared by the Loot tab, the council's nights and the decision record: each
// award with its night, its claimant and its decision record (ms/os, reason, the votes at award
// time), and the roster the council ranks.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// lootAward is one award row as the Loot tab reads it: the item, the raid (or, for a logged
// night with no planned raid, the report) it was filed under (nil for none yet) and when, with
// the line the page shows and the decision record.
type lootAward struct {
	itemID    int
	raidID    *int64
	reportID  *string
	awardedAt time.Time
	to        *AwardedTo
	class     string
	decision  awardDecisionRecord
}

// awardDecisionRecord is what the council kept with an award. Every field is nil on an award
// made before the council recorded it.
type awardDecisionRecord struct {
	msOs                        *string
	reason                      *string
	votesYes, votesNo, officers *int
}

// raidlessAwardNight is how long an award filed under no raid still marks its item awarded on
// the Loot tab: one raid night (the longest raid's match window, an hour before to six after a
// four-hour raid, is eleven hours).
const raidlessAwardNight = 12 * time.Hour

// lootAwards reads every award row for guildID's encounterID, oldest first, with the claimant's
// and the awarding officer's own names resolved. Every row is kept: the tier's count reads them
// all, and tonightsAwards picks the ones that mark an item awarded now.
func (s *Store) lootAwards(ctx context.Context, guildID, encounterID int64) ([]lootAward, error) {
	rows, err := s.Pool.Query(ctx, `
		select la.item_id, la.raid_id, la.report_id, la.id, la.character_key, coalesce(nullif(ae.name, ''), c.name, ''), la.awarded_at,
		       la.awarded_by, coalesce(c.class, ''), la.ms_os, la.reason, la.votes_yes, la.votes_no, la.votes_officers
		from loot_awards la
		left join addon_exports ae on ae.character_key = la.character_key
		left join characters c on c.key = la.character_key
		where la.guild_id = $1 and la.encounter_id = $2
		order by la.awarded_at, la.id`, guildID, encounterID)
	if err != nil {
		return nil, fmt.Errorf("guilds: loot awards: %w", err)
	}
	defer rows.Close()
	var out []lootAward
	var awarders []int64
	for rows.Next() {
		var row lootAward
		var awardedBy *int64
		a := &AwardedTo{}
		d := &row.decision
		if err := rows.Scan(&row.itemID, &row.raidID, &row.reportID, &a.AwardID, &a.CharacterKey, &a.Name,
			&row.awardedAt, &awardedBy, &row.class, &d.msOs, &d.reason, &d.votesYes, &d.votesNo, &d.officers); err != nil {
			return nil, fmt.Errorf("guilds: loot awards: %w", err)
		}
		a.At = row.awardedAt.UTC().Format(time.RFC3339)
		row.to = a
		out = append(out, row)
		awarders = append(awarders, awardedByOrZero(awardedBy))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: loot awards: %w", err)
	}
	return out, s.nameAwarders(ctx, guildID, out, awarders)
}

func awardedByOrZero(id *int64) int64 {
	if id == nil {
		return 0
	}
	return *id
}

// nameAwarders fills each award's by_name in one query. An award with no recorded officer keeps "".
func (s *Store) nameAwarders(ctx context.Context, guildID int64, awards []lootAward, awarders []int64) error {
	if s.Accounts == nil {
		return nil
	}
	names, err := s.awarderNamesOf(ctx, guildID, awarders)
	if err != nil {
		return err
	}
	for i, id := range awarders {
		if id != 0 {
			awards[i].to.ByName = names.of(id)
		}
	}
	return nil
}

// awarderNamesOf is awarderNames for a list of awarders that may hold 0 (no recorded officer).
func (s *Store) awarderNamesOf(ctx context.Context, guildID int64, awarders []int64) (awarderNameSet, error) {
	ids := make([]int64, 0, len(awarders))
	for _, id := range awarders {
		if id != 0 {
			ids = append(ids, id)
		}
	}
	return s.awarderNames(ctx, guildID, ids)
}

// tonightsAwards is each item's award on the night being played, keyed by item id: an award
// filed under night (resolveAwardNight with no named raid: the raid in play, else the latest
// night that killed the boss this week), or one filed under no night within the last
// raidlessAwardNight. Last week's award leaves the item free to award again; the latest award of
// an item wins.
func tonightsAwards(awards []lootAward, night awardNight, now time.Time) map[int]*AwardedTo {
	out := map[int]*AwardedTo{}
	for _, a := range awards {
		if awardOnNight(a, night, now) {
			out[a.itemID] = a.to
		}
	}
	return out
}

// awardOnNight is whether a marks its item awarded on night, the Loot tab's rule: the same raid,
// else (both raid-less) the same logged night, else a raid-less, report-less award made within
// raidlessAwardNight, whatever the night.
func awardOnNight(a lootAward, night awardNight, now time.Time) bool {
	if a.raidID == nil && a.reportID == nil {
		return now.Sub(a.awardedAt) <= raidlessAwardNight
	}
	return onNight(a.raidID, a.reportID, a.awardedAt, night, now)
}

// onNight is whether a row filed under raidID and reportID at `at` belongs to night, the
// council's strict rule: the same raid, else (both raid-less) the same logged night, else, for
// the live night only, a raid-less, report-less row made within raidlessAwardNight. nightMatchSQL
// is the same rule in SQL.
func onNight(raidID *int64, reportID *string, at time.Time, night awardNight, now time.Time) bool {
	switch {
	case raidID != nil:
		return night.raidID != nil && *raidID == *night.raidID
	case reportID != nil:
		return night.raidID == nil && night.reportID != nil && *reportID == *night.reportID
	default:
		return night.raidID == nil && night.reportID == nil && now.Sub(at) <= raidlessAwardNight
	}
}

// lootRosterSet is the roster the council ranks: the guild, its roster rows and the invited
// accounts' picked characters, which an award names by the account's synthetic key.
type lootRosterSet struct {
	guild   Guild
	roster  []RosterRow
	invited map[string]RosterRow // by award key
}

// lootRoster reads the roster the council ranks. An account that joined by invite has no
// character row of the guild's, but its gear is on the site and it takes part in the raid, so it
// is a candidate like any member.
func (s *Store) lootRoster(ctx context.Context, guildID int64) (lootRosterSet, error) {
	set, err := s.invitedSet(ctx, guildID)
	if err != nil {
		return lootRosterSet{}, err
	}
	roster, err := s.HomeRoster(ctx, guildID, set.guild, 0, false, false)
	if err != nil {
		return lootRosterSet{}, err
	}
	set.roster = append(roster, set.roster...)
	return set, nil
}

// invitedSet is the guild and its invited accounts' rows only (roster holds just those), enough to
// name awards (nameInvited) without reading the whole roster.
func (s *Store) invitedSet(ctx context.Context, guildID int64) (lootRosterSet, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return lootRosterSet{}, err
	}
	invited, err := s.invitedRosterRows(ctx, guildID, g)
	if err != nil {
		return lootRosterSet{}, err
	}
	set := lootRosterSet{guild: g, roster: invited, invited: make(map[string]RosterRow, len(invited))}
	for _, row := range invited {
		set.invited[row.awardKey] = row
	}
	return set, nil
}

// nameInvited names each award to an invited member: the account's synthetic row has no export
// carrying a name or class, so the picked character's stand in for them.
func (lr lootRosterSet) nameInvited(awards []lootAward) {
	for i := range awards {
		a := &awards[i]
		row, ok := lr.invited[a.to.CharacterKey]
		if !ok {
			continue
		}
		if a.to.Name == "" {
			a.to.Name = row.Name
		}
		if a.class == "" && row.Class != nil {
			a.class = *row.Class
		}
	}
}
