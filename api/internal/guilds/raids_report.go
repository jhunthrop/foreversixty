// api/internal/guilds/raids_report.go
//
// Lane G2: the raid-night report on a raid view (design/specs/2026-10-27-guild-attendance.md
// §3.4 and the contract paragraph). The store builds every row for every viewer; rows that name
// a character carry an unexported own flag, and Service.scopeAttendance narrows a member's
// report to their own characters with ownOnly.
package guilds

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

// KillOrderView is one named boss of the night. For a boss never killed, Killed is false,
// FightIndex is its last pull and WipesBefore counts every pull.
type KillOrderView struct {
	EncounterID int    `json:"encounter_id"`
	Name        string `json:"name"`
	FightIndex  int    `json:"fight_index"`
	Killed      bool   `json:"killed"`
	KillMS      *int64 `json:"kill_ms"`
	WipesBefore int    `json:"wipes_before"`
	Deaths      int    `json:"deaths"`
}

// UnnamedView counts the pulls with no published encounter.
type UnnamedView struct {
	Pulls int `json:"pulls"`
	Kills int `json:"kills"`
	Wipes int `json:"wipes"`
}

// ReportDeathView is one raider's deaths, split on wipes and kills.
type ReportDeathView struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	OnWipes      int    `json:"on_wipes"`
	OnKills      int    `json:"on_kills"`
	own          bool
}

// ConsumableRow is one roster character's Before pill. The log carries no buff events yet,
// so flask, food and potions are null.
type ConsumableRow struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Before       string `json:"before"`
	Flask        *bool  `json:"flask"`
	Food         *bool  `json:"food"`
	Potions      *int   `json:"potions"`
	own          bool
}

// ReportConsumablesView is the night's consumables: LogHasBuffs is false until buff events are stored.
type ReportConsumablesView struct {
	LogHasBuffs bool            `json:"log_has_buffs"`
	Rows        []ConsumableRow `json:"rows"`
}

// ReportLootAward is one award made for this raid.
type ReportLootAward struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
	Quality  int    `json:"quality"`
	// Icon is the item's icon file name from the loot table, "" when the build has none (the
	// page then draws the client's question mark).
	Icon         string `json:"icon"`
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	ByName       string `json:"by_name"`
	At           string `json:"at"`
}

// ReportLootBoss groups a boss's awards.
type ReportLootBoss struct {
	EncounterID int64             `json:"encounter_id"`
	Boss        string            `json:"boss"`
	Awards      []ReportLootAward `json:"awards"`
}

// MoveRow is one raider's change on a boss against the previous raid of the zone.
type MoveRow struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Metric       string `json:"metric"`
	Encounter    string `json:"encounter"`
	Before       int    `json:"before"`
	After        int    `json:"after"`
	own          bool
	delta        float64
}

// MovesView is the biggest changes since the previous done raid of the zone.
type MovesView struct {
	Previous *RaidRefView `json:"previous"`
	Rows     []MoveRow    `json:"rows"`
}

// ReportView is the raid-night report.
type ReportView struct {
	KillOrder   []KillOrderView       `json:"kill_order"`
	Unnamed     *UnnamedView          `json:"unnamed"`
	Deaths      []ReportDeathView     `json:"deaths"`
	Consumables ReportConsumablesView `json:"consumables"`
	Loot        []ReportLootBoss      `json:"loot"`
	Moves       MovesView             `json:"moves"`
	PostedAt    *time.Time            `json:"posted_at"`
}

// ownOnly is the report a member sees: every kill, unnamed and loot row, and only their own
// characters' deaths, consumables and moves. Nil stays nil.
func (r *ReportView) ownOnly() *ReportView {
	if r == nil {
		return nil
	}
	out := *r
	out.Deaths = keepOwn(r.Deaths, func(d ReportDeathView) bool { return d.own })
	out.Consumables.Rows = keepOwn(r.Consumables.Rows, func(c ConsumableRow) bool { return c.own })
	out.Moves.Rows = keepOwn(r.Moves.Rows, func(m MoveRow) bool { return m.own })
	return &out
}

func keepOwn[T any](rows []T, own func(T) bool) []T {
	out := []T{}
	for _, r := range rows {
		if own(r) {
			out = append(out, r)
		}
	}
	return out
}

// buildReport assembles the whole report of one matched raid view.
func (s *Store) buildReport(ctx context.Context, vc raidViewContext, v RaidView, meta raidMatchMeta) (*ReportView, error) {
	reportID := *meta.reportID
	fights, err := s.reportFights(ctx, reportID)
	if err != nil {
		return nil, err
	}
	kills, unnamed := killOrderOf(fights)
	rep := &ReportView{KillOrder: kills, Unnamed: unnamed, PostedAt: meta.postedAt}
	if rep.Deaths, err = s.reportDeaths(ctx, vc, reportID); err != nil {
		return nil, err
	}
	if rep.Consumables, err = s.reportConsumables(ctx, vc, reportID); err != nil {
		return nil, err
	}
	if rep.Loot, err = s.reportLoot(ctx, vc, v.ID); err != nil {
		return nil, err
	}
	if rep.Moves, err = s.reportMoves(ctx, vc, v, reportID, kills); err != nil {
		return nil, err
	}
	return rep, nil
}

type reportFight struct {
	index       int
	encounterID *int
	name        string
	kill        bool
	durationMS  int64
	deaths      int
}

func (s *Store) reportFights(ctx context.Context, reportID string) ([]reportFight, error) {
	rows, err := s.Pool.Query(ctx, `
		select f.fight_index, f.encounter_id, f.name, f.kill, f.duration_ms, f.deaths
		from fights f join reports r on r.id = f.report_id
		where f.report_id = $1 and `+presence.EvidenceReport("r")+` order by f.fight_index`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report fights: %w", err)
	}
	defer rows.Close()
	var out []reportFight
	for rows.Next() {
		var f reportFight
		if err := rows.Scan(&f.index, &f.encounterID, &f.name, &f.kill, &f.durationMS, &f.deaths); err != nil {
			return nil, fmt.Errorf("guilds: report fights: %w", err)
		}
		out = append(out, f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: report fights: %w", err)
	}
	return out, nil
}

// killOrderOf folds the fights, in order, into one entry per named boss (a boss's pulls after
// its first kill are ignored) sorted by their deciding fight, and counts the unnamed pulls.
func killOrderOf(fights []reportFight) ([]KillOrderView, *UnnamedView) {
	byEncounter := map[int]*KillOrderView{}
	order := []*KillOrderView{}
	var unnamed UnnamedView
	for _, f := range fights {
		if f.encounterID == nil {
			unnamed.Pulls++
			if f.kill {
				unnamed.Kills++
			}
			continue
		}
		e := byEncounter[*f.encounterID]
		if e == nil {
			e = &KillOrderView{EncounterID: *f.encounterID, Name: f.name}
			byEncounter[*f.encounterID] = e
			order = append(order, e)
		}
		if e.Killed {
			continue
		}
		e.FightIndex = f.index
		if f.kill {
			// A killed boss shows the deaths of its kill; an unkilled one the sum of its pulls.
			ms := f.durationMS
			e.Killed, e.KillMS, e.Deaths = true, &ms, f.deaths
		} else {
			e.WipesBefore++
			e.Deaths += f.deaths
		}
	}
	out := make([]KillOrderView, len(order))
	for i, e := range order {
		out[i] = *e
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].FightIndex < out[j].FightIndex })
	if unnamed.Pulls == 0 {
		return out, nil
	}
	unnamed.Wipes = unnamed.Pulls - unnamed.Kills
	return out, &unnamed
}

func (s *Store) reportDeaths(ctx context.Context, vc raidViewContext, reportID string) ([]ReportDeathView, error) {
	rows, err := s.Pool.Query(ctx, `
		select m.player_key, coalesce(max(m.player_name), ''), coalesce(lower(max(m.class)), ''),
		       coalesce(sum(m.deaths) filter (where not m.kill), 0), coalesce(sum(m.deaths) filter (where m.kill), 0)
		from fight_metrics m join reports r on r.id = m.report_id
		where m.report_id = $1 and m.state = 'ok' and `+presence.EvidenceReport("r")+`
		group by m.player_key having coalesce(sum(m.deaths), 0) > 0`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report deaths: %w", err)
	}
	defer rows.Close()
	out := []ReportDeathView{}
	for rows.Next() {
		var d ReportDeathView
		if err := rows.Scan(&d.CharacterKey, &d.Name, &d.Class, &d.OnWipes, &d.OnKills); err != nil {
			return nil, fmt.Errorf("guilds: report deaths: %w", err)
		}
		d.own = vc.ownsKey(d.CharacterKey)
		out = append(out, d)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: report deaths: %w", err)
	}
	sortDeaths(out)
	return out, nil
}

// sortDeaths orders raiders most deaths first, ties by class order then name.
func sortDeaths(out []ReportDeathView) {
	sort.SliceStable(out, func(i, j int) bool {
		if a, b := out[i].OnWipes+out[i].OnKills, out[j].OnWipes+out[j].OnKills; a != b {
			return a > b
		}
		if a, b := rank(classOrder, out[i].Class), rank(classOrder, out[j].Class); a != b {
			return a < b
		}
		return out[i].Name < out[j].Name
	})
}

// ownsKey reports whether the viewer's account holds the roster character key.
func (vc raidViewContext) ownsKey(key string) bool {
	owner, ok := vc.ownerOf[key]
	return ok && owner == vc.viewerID
}
