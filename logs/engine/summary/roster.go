// logs/engine/summary/roster.go
package summary

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
)

// CombatantRow is one player's gear, talents and consumables at the pull.
type CombatantRow struct {
	GUID         string       `json:"guid"`
	Name         string       `json:"name"`
	SpecID       int64        `json:"spec_id,omitempty"`
	Spec         string       `json:"spec,omitempty"`
	ItemLevel    int64        `json:"item_level,omitempty"`
	Gear         []event.Item `json:"gear"`
	Talents      []int64      `json:"talents"`
	Consumables  []AuraRef    `json:"consumables"`
	RaidBuffs    []AuraRef    `json:"raid_buffs"`
	MissingBuffs []int64      `json:"missing_buffs"`
}

// RosterRow is one player's line in the fight's roster.
type RosterRow struct {
	GUID        string  `json:"guid"`
	Name        string  `json:"name"`
	Class       string  `json:"class,omitempty"`
	ClassSource string  `json:"class_source,omitempty"`
	SpecID      int64   `json:"spec_id,omitempty"`
	Spec        string  `json:"spec,omitempty"`
	Role        string  `json:"role"`
	ItemLevel   int64   `json:"item_level,omitempty"`
	ActiveMS    int64   `json:"active_ms"`
	ActivityPct float64 `json:"activity_pct"`
	Deaths      int     `json:"deaths"`
	DamageDone  int64   `json:"damage_done"`
	HealingDone int64   `json:"healing_done"`
	DamageTaken int64   `json:"damage_taken"`
	DPS         float64 `json:"dps"`
	HPS         float64 `json:"hps"`
	DTPS        float64 `json:"dtps"`
}

// MetricRow is one ranking metric row: one player, one fight.
type MetricRow struct {
	ReportID      string    `json:"report_id"`
	FightIndex    int       `json:"fight_index"`
	PlayerGUID    string    `json:"player_guid"`
	PlayerName    string    `json:"player_name"`
	Class         string    `json:"class,omitempty"`
	SpecID        int64     `json:"spec_id,omitempty"`
	Spec          string    `json:"spec,omitempty"`
	Role          string    `json:"role"`
	Metric        string    `json:"metric"`
	Value         float64   `json:"value"`
	ActiveMS      int64     `json:"active_ms"`
	ItemLevel     int64     `json:"item_level,omitempty"`
	DurationMS    int64     `json:"duration_ms"`
	EncounterID   int64     `json:"encounter_id"`
	Difficulty    int64     `json:"difficulty"`
	Size          int64     `json:"size"`
	Kill          bool      `json:"kill"`
	Date          time.Time `json:"date"`
	EngineVersion string    `json:"engine_version"`
}

func (a *Accumulator) threatRows() []ThreatRow {
	out := make([]ThreatRow, 0, len(a.threat))
	for guid, v := range a.threat {
		out = append(out, ThreatRow{
			GUID: guid, Name: a.name(guid), Threat: v,
			ModelVersion: a.opt.Threat.Version(), Complete: a.opt.Threat.Complete(),
		})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Threat != out[j].Threat {
			return out[i].Threat > out[j].Threat
		}
		return out[i].GUID < out[j].GUID
	})
	return out
}

func (a *Accumulator) combatantRows() []CombatantRow {
	out := make([]CombatantRow, 0, len(a.combatants))
	for guid, c := range a.combatants {
		row := CombatantRow{
			GUID: guid, Name: a.name(guid),
			SpecID: c.SpecID, Spec: a.opt.SpecNames[c.SpecID],
			ItemLevel: c.ItemLevel,
			Gear:      copySlice(c.Gear), Talents: copySlice(c.Talents),
		}
		if row.Gear == nil {
			row.Gear = []event.Item{}
		}
		if row.Talents == nil {
			row.Talents = []int64{}
		}
		seen := map[int64]bool{}
		for _, au := range c.Auras {
			seen[au.SpellID] = true
			ref := AuraRef{SpellID: au.SpellID, SourceGUID: au.SourceGUID}
			if name, ok := a.opt.ConsumableSpells[au.SpellID]; ok {
				ref.Name = name
				row.Consumables = append(row.Consumables, ref)
			}
			if name, ok := a.opt.RaidBuffSpells[au.SpellID]; ok {
				ref.Name = name
				row.RaidBuffs = append(row.RaidBuffs, ref)
			}
		}
		for id := range a.opt.RaidBuffSpells {
			if !seen[id] {
				row.MissingBuffs = append(row.MissingBuffs, id)
			}
		}
		sortAuraRefs(row.Consumables)
		sortAuraRefs(row.RaidBuffs)
		sort.Slice(row.MissingBuffs, func(i, j int) bool { return row.MissingBuffs[i] < row.MissingBuffs[j] })
		if row.Consumables == nil {
			row.Consumables = []AuraRef{}
		}
		if row.RaidBuffs == nil {
			row.RaidBuffs = []AuraRef{}
		}
		if row.MissingBuffs == nil {
			row.MissingBuffs = []int64{}
		}
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GUID < out[j].GUID })
	return out
}

// rosterRows builds one line per player present, from the tables already
// computed so nothing is scanned twice.
func (a *Accumulator) rosterRows(f fight.Fight, s Summary) []RosterRow {
	index := func(rows []Actor) map[string]Actor {
		m := make(map[string]Actor, len(rows))
		for _, r := range rows {
			m[r.GUID] = r
		}
		return m
	}
	dd, hd, dt := index(s.DamageDone), index(s.Healing), index(s.DamageTaken)
	deaths := map[string]int{}
	for _, d := range s.Deaths {
		deaths[d.GUID]++
	}
	specs := map[string]*event.Combatant{}
	for g, c := range a.combatants {
		specs[g] = c
	}

	seconds := float64(s.DurationMS) / 1000
	out := make([]RosterRow, 0, len(f.Players))
	for _, guid := range f.Players {
		row := RosterRow{GUID: guid, Name: a.name(guid), Deaths: deaths[guid]}
		if a.opt.Registry != nil {
			if u, ok := a.opt.Registry.Get(guid); ok {
				row.Class, row.ClassSource = u.Class, u.ClassSource
				row.SpecID, row.ItemLevel = u.SpecID, u.ItemLevel
			}
		}
		if c, ok := specs[guid]; ok {
			row.SpecID, row.ItemLevel = c.SpecID, c.ItemLevel
		}
		row.Spec = a.opt.SpecNames[row.SpecID]
		row.DamageDone = dd[guid].Effective
		row.HealingDone = hd[guid].Effective
		row.DamageTaken = dt[guid].Effective
		row.ActiveMS = dd[guid].ActiveMS
		if row.ActiveMS == 0 {
			row.ActiveMS = hd[guid].ActiveMS
		}
		if seconds > 0 {
			row.DPS = float64(row.DamageDone) / seconds
			row.HPS = float64(row.HealingDone) / seconds
			row.DTPS = float64(row.DamageTaken) / seconds
			// markActive credits a full ActiveGap for an actor's first
			// action, so on a fight shorter than that gap the ratio can
			// exceed one. Clamp what the table shows; the underlying
			// ActiveMS is what the ranking metric divides by and is left
			// alone.
			row.ActivityPct = min(float64(row.ActiveMS)/float64(s.DurationMS)*100, 100)
		}
		row.Role = role(row)
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GUID < out[j].GUID })
	return out
}

// role picks the metric a player is ranked on. It is derived from what the
// player actually did, so it works with or without COMBATANT_INFO.
func role(r RosterRow) string {
	switch {
	case r.HealingDone > r.DamageDone:
		return "healer"
	case r.DamageTaken > r.DamageDone:
		return "tank"
	default:
		return "dps"
	}
}

// Metrics builds the ranking rows for a closed fight. Only encounters are
// ranked: trash has nothing to compare against.
func (a *Accumulator) Metrics(reportID string, f fight.Fight, s Summary, engineVersion string) []MetricRow {
	if f.Kind != fight.Encounter {
		return nil
	}
	out := make([]MetricRow, 0, len(s.Roster))
	for _, r := range s.Roster {
		row := MetricRow{
			ReportID: reportID, FightIndex: f.Index,
			PlayerGUID: r.GUID, PlayerName: r.Name,
			Class: r.Class, SpecID: r.SpecID, Spec: r.Spec, Role: r.Role,
			ActiveMS: r.ActiveMS, ItemLevel: r.ItemLevel,
			DurationMS: s.DurationMS, EncounterID: f.EncounterID,
			Difficulty: f.Difficulty, Size: f.Size, Kill: f.Kill,
			Date: f.Start, EngineVersion: engineVersion,
		}
		switch r.Role {
		case "healer":
			row.Metric, row.Value = "hps", r.HPS
		case "tank":
			row.Metric, row.Value = "dtps", r.DTPS
		default:
			row.Metric, row.Value = "dps", r.DPS
		}
		out = append(out, row)
	}
	return out
}
