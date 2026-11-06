// api/internal/guilds/loot_council_nights.go
//
// The council's nights (design/specs/2026-11-05-web-layout.md §7.2): the guild's loot grouped per
// raid night, newest first, each drop open or awarded, with the decision record. There is no
// record of what dropped beyond what the council acted on, so a night's drops are the items it
// awarded, voted on or flagged; the current night is listed even before anything happens on it,
// and the page offers the whole loot table there (GET .../loot).
package guilds

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// Drop states.
const (
	dropOpen    = "open"
	dropAwarded = "awarded"
)

// Night list bounds.
const (
	councilNightsDefault = 8
	councilNightsMax     = 30
)

// LootDecision is an award's decision record. MsOs, Reason and the vote counts are null on an
// award made before the council recorded them.
type LootDecision struct {
	AwardID      int64   `json:"award_id"`
	ItemID       int     `json:"item_id"`
	ItemName     string  `json:"item_name"`
	Icon         string  `json:"icon"`
	Quality      int     `json:"quality"`
	CharacterKey string  `json:"character_key"`
	Name         string  `json:"name"`
	Class        string  `json:"class"`
	MsOs         *string `json:"ms_os"`
	Reason       *string `json:"reason"`
	VotesYes     *int    `json:"votes_yes"`
	VotesNo      *int    `json:"votes_no"`
	// VotesOfficers is how many officers could vote when the drop was awarded.
	VotesOfficers *int   `json:"votes_officers"`
	AwardedBy     string `json:"awarded_by"`
	At            string `json:"at"`
}

// LootNightDrop is one drop of a night.
type LootNightDrop struct {
	ItemID      int           `json:"item_id"`
	Name        string        `json:"name"`
	Icon        string        `json:"icon"`
	Quality     int           `json:"quality"`
	Kind        string        `json:"kind"`
	EncounterID int64         `json:"encounter_id"`
	Boss        string        `json:"boss"`
	State       string        `json:"state"`
	Award       *LootDecision `json:"award"`
}

// LootNight is one raid night of the council.
type LootNight struct {
	// Night is the key every council call takes: "raid:<id>", "report:<id>" or "live".
	Night    string  `json:"night"`
	RaidID   *int64  `json:"raid_id"`
	ReportID *string `json:"report_id"`
	// Date is the night's day in the guild's raid zone (UTC when it has none), the CSV's rule.
	Date    string          `json:"date"`
	At      string          `json:"at"`
	Zone    string          `json:"zone"`
	Current bool            `json:"current"`
	Drops   []LootNightDrop `json:"drops"`
	// Decisions is the night's decision record: its awards, oldest first.
	Decisions []LootDecision `json:"decisions"`
}

// LootNightsView is the body of GET /v1/guilds/{id}/loot/nights.
type LootNightsView struct {
	Officers int         `json:"officers"`
	Nights   []LootNight `json:"nights"`
}

// toDecision is a's decision record.
func toDecision(a lootAward) LootDecision {
	it, _ := lootItemByID(a.itemID)
	return LootDecision{AwardID: a.to.AwardID, ItemID: a.itemID, ItemName: it.Name, Icon: it.Icon, Quality: it.Quality,
		CharacterKey: a.to.CharacterKey, Name: a.to.Name, Class: a.class, MsOs: a.decision.msOs, Reason: a.decision.reason,
		VotesYes: a.decision.votesYes, VotesNo: a.decision.votesNo, VotesOfficers: a.decision.officers,
		AwardedBy: a.to.ByName, At: a.to.At}
}

// nightAward is itemID's latest award filed under night (onNight), nil when it is open.
func nightAward(awards []lootAward, night awardNight, itemID int, now time.Time) *LootDecision {
	var out *LootDecision
	for _, a := range awards {
		if a.itemID == itemID && onNight(a.raidID, a.reportID, a.awardedAt, night, now) {
			d := toDecision(a)
			out = &d
		}
	}
	return out
}

// councilNightRow is one night the listing found, before its drops are read.
type councilNightRow struct {
	night awardNight
	at    time.Time
	// zone is a planned raid's zone key; nil for a logged night with no planned raid.
	zone *string
}

// CouncilNights lists the guild's last limit nights, newest first, the current night included.
func (s *Store) CouncilNights(ctx context.Context, guildID int64, limit int, now time.Time) (LootNightsView, error) {
	current, err := s.resolveAwardNight(ctx, guildID, allLoot, nil, now)
	if err != nil {
		return LootNightsView{}, err
	}
	found, err := s.councilNightRows(ctx, guildID, limit, now)
	if err != nil {
		return LootNightsView{}, err
	}
	found = withCurrentNight(found, current, now)
	loc, err := s.exportLocation(ctx, guildID)
	if err != nil {
		return LootNightsView{}, err
	}
	officers, err := s.officerCount(ctx, guildID)
	if err != nil {
		return LootNightsView{}, err
	}
	act, err := s.councilActivity(ctx, guildID)
	if err != nil {
		return LootNightsView{}, err
	}
	view := LootNightsView{Officers: officers, Nights: make([]LootNight, 0, len(found))}
	for _, n := range found {
		view.Nights = append(view.Nights, act.night(n, current, loc, now))
	}
	return view, nil
}

// withCurrentNight puts the current night first when the listing does not hold it (the live
// night, or a night with nothing on it yet).
func withCurrentNight(found []councilNightRow, current awardNight, now time.Time) []councilNightRow {
	for _, n := range found {
		if n.night.key() == current.key() {
			return found
		}
	}
	return append([]councilNightRow{{night: current, at: now}}, found...)
}

// councilNightRows is the guild's loot nights, newest first: every planned raid of a loot table
// zone (not cancelled) that is done, in play, or holds council activity, and every unflagged log of
// the guild no such planned raid holds that killed a table boss or holds council activity. A live night with activity
// in the last raid night leads.
func (s *Store) councilNightRows(ctx context.Context, guildID int64, limit int, now time.Time) ([]councilNightRow, error) {
	rows, err := s.Pool.Query(ctx, `
		with act as (
		  select raid_id, report_id, awarded_at as at from loot_awards where guild_id = $1
		  union all select raid_id, report_id, voted_at from loot_votes where guild_id = $1
		  union all select raid_id, report_id, set_at from loot_interest where guild_id = $1
		)
		select * from (
		  select g.id as raid_id, null::text as report_id, g.starts_at as at, g.zone from guild_raids g
		  where g.guild_id = $1 and g.zone = any($2) and g.status <> 'cancelled'
		    and g.starts_at - make_interval(hours => $5) <= $4
		    and (g.status = 'done'
		         or $4 <= g.starts_at + make_interval(mins => g.duration_minutes) + make_interval(hours => $6)
		         or exists (select 1 from act where act.raid_id = g.id))
		  union all
		  select null, r.id, `+logAnchorSQL("r")+`, null from reports r
		  where r.guild_id = $1 and `+presence.NotFlagged("r")+`
		    and not exists (select 1 from guild_raids gr
		                    where gr.report_id = r.id and gr.guild_id = $1 and gr.zone = any($2) and gr.status <> 'cancelled')
		    and (exists (select 1 from fights f where f.report_id = r.id and f.encounter_id = any($3) and f.kill)
		         or exists (select 1 from act where act.raid_id is null and act.report_id = r.id))
		  union all
		  select null, null, max(act.at), null from act
		  where act.raid_id is null and act.report_id is null and act.at >= $7
		  having count(*) > 0
		) n order by at desc limit $8`,
		guildID, allLoot.zones, allLoot.encounters, now, matchLeadHours, matchTailHours, liveCutoff(now), limit)
	if err != nil {
		return nil, fmt.Errorf("guilds: council nights: %w", err)
	}
	defer rows.Close()
	var out []councilNightRow
	for rows.Next() {
		var n councilNightRow
		if err := rows.Scan(&n.night.raidID, &n.night.reportID, &n.at, &n.zone); err != nil {
			return nil, fmt.Errorf("guilds: council nights: %w", err)
		}
		out = append(out, n)
	}
	return out, rows.Err()
}

// councilActivityRow is a vote or a flag as the night listing places it.
type councilActivityRow struct {
	raidID   *int64
	reportID *string
	itemID   int
	at       time.Time
}

// councilActivitySet is everything the council did, for placing on nights.
type councilActivitySet struct {
	awards []lootAward
	acted  []councilActivityRow
}

// councilActivity reads the guild's awards, votes and flags.
func (s *Store) councilActivity(ctx context.Context, guildID int64) (councilActivitySet, error) {
	lr, err := s.invitedSet(ctx, guildID)
	if err != nil {
		return councilActivitySet{}, err
	}
	awards, err := s.lootAwards(ctx, guildID, allLoot.encounters)
	if err != nil {
		return councilActivitySet{}, err
	}
	lr.nameInvited(awards)
	rows, err := s.Pool.Query(ctx, `
		select raid_id, report_id, item_id, voted_at from loot_votes where guild_id = $1 and encounter_id = any($2)
		union all
		select raid_id, report_id, item_id, set_at from loot_interest where guild_id = $1 and encounter_id = any($2)`,
		guildID, allLoot.encounters)
	if err != nil {
		return councilActivitySet{}, fmt.Errorf("guilds: council activity: %w", err)
	}
	defer rows.Close()
	set := councilActivitySet{awards: awards}
	for rows.Next() {
		var r councilActivityRow
		if err := rows.Scan(&r.raidID, &r.reportID, &r.itemID, &r.at); err != nil {
			return councilActivitySet{}, fmt.Errorf("guilds: council activity: %w", err)
		}
		set.acted = append(set.acted, r)
	}
	return set, rows.Err()
}

// night is the night n with its drops (in the loot table's order) and decision record.
func (set councilActivitySet) night(n councilNightRow, current awardNight, loc *time.Location, now time.Time) LootNight {
	out := LootNight{Night: n.night.key(), RaidID: n.night.raidID, ReportID: n.night.reportID,
		Date: n.at.In(loc).Format(exportDateLayout), At: n.at.UTC().Format(time.RFC3339),
		Current: n.night.key() == current.key(), Drops: []LootNightDrop{}, Decisions: []LootDecision{}}
	touched := map[int]bool{}
	for _, a := range set.awards {
		if onNight(a.raidID, a.reportID, a.awardedAt, n.night, now) {
			touched[a.itemID] = true
			out.Decisions = append(out.Decisions, toDecision(a))
		}
	}
	for _, r := range set.acted {
		if onNight(r.raidID, r.reportID, r.at, n.night, now) {
			touched[r.itemID] = true
		}
	}
	for _, it := range lootTable {
		if !touched[it.ItemID] {
			continue
		}
		drop := LootNightDrop{ItemID: it.ItemID, Name: it.Name, Icon: it.Icon, Quality: it.Quality, Kind: lootKind(it),
			EncounterID: it.EncounterID, Boss: bossName(it.EncounterID), State: dropOpen,
			Award: nightAward(set.awards, n.night, it.ItemID, now)}
		if drop.Award != nil {
			drop.State = dropAwarded
		}
		out.Drops = append(out.Drops, drop)
	}
	sort.SliceStable(out.Decisions, func(i, j int) bool { return out.Decisions[i].At < out.Decisions[j].At })
	out.Zone = nightZone(n.zone, out.Drops)
	return out
}

// nightZone is the raid a night is in: a planned raid's zone; else, for a logged night, the raid
// of its first drop; else the loot table's first raid.
func nightZone(zone *string, drops []LootNightDrop) string {
	if zone != nil {
		return lootRaidName(*zone)
	}
	if len(drops) > 0 {
		return bossRaidName(drops[0].EncounterID)
	}
	if len(lootBosses) > 0 {
		return lootBosses[0].RaidName
	}
	return ""
}

// withoutVotes is the member's copy of the nights: every decision keeps who, MS/OS, reason, by whom
// and when, never the officers' vote counts.
func (v LootNightsView) withoutVotes() LootNightsView {
	out := LootNightsView{Officers: v.Officers, Nights: make([]LootNight, len(v.Nights))}
	for i, n := range v.Nights {
		n.Drops = append([]LootNightDrop(nil), n.Drops...)
		n.Decisions = append([]LootDecision(nil), n.Decisions...)
		for j := range n.Drops {
			if n.Drops[j].Award != nil {
				d := n.Drops[j].Award.withoutVotes()
				n.Drops[j].Award = &d
			}
		}
		for j := range n.Decisions {
			n.Decisions[j] = n.Decisions[j].withoutVotes()
		}
		out.Nights[i] = n
	}
	return out
}

// withoutVotes is d with its vote counts cleared.
func (d LootDecision) withoutVotes() LootDecision {
	d.VotesYes, d.VotesNo, d.VotesOfficers = nil, nil, nil
	return d
}
