package reports

import (
	"math"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// TrashGroup is one zone's trash folded into one entry, the way Warcraft Logs collapses it: the
// Go port of web/src/lib/report/fights.ts trashGroup, so the apps draw the rows the web draws.
type TrashGroup struct {
	// Key is the first pull's index: the group's name in a meters call (scope=trash).
	Key        int    `json:"key"`
	Zone       string `json:"zone"`
	Fights     []int  `json:"fights"`
	DurationMS int64  `json:"duration_ms"`
}

// TrashGroups is every zone's trash entry, zone by zone in the order the night reached them
// (fights.ts fightRows). Trash the engine gave no zone is never grouped: it keeps its own rows.
func TrashGroups(fights []store.FightEntry) []TrashGroup {
	var zones []string
	byZone := map[string]*TrashGroup{}
	for _, f := range fights {
		if f.Zone == "" || f.Kind == string(fight.Encounter) {
			continue
		}
		g, ok := byZone[f.Zone]
		if !ok {
			g = &TrashGroup{Key: f.Index, Zone: f.Zone, Fights: []int{}}
			byZone[f.Zone] = g
			zones = append(zones, f.Zone)
		}
		g.Fights = append(g.Fights, f.Index)
		g.DurationMS += f.DurationMS
	}
	out := make([]TrashGroup, 0, len(zones))
	for _, z := range zones {
		out = append(out, *byZone[z])
	}
	return out
}

// trashGroupKeyed is the group whose key is key.
func trashGroupKeyed(groups []TrashGroup, key int) (TrashGroup, bool) {
	for _, g := range groups {
		if g.Key == key {
			return g, true
		}
	}
	return TrashGroup{}, false
}

// meterRow is one player's numbers across the pulls folded into a view: the Go port of
// web/src/lib/report/night.ts aggregateNight's player rows. One pull is a fold of one, so a
// single fight and a combined trash view read their figures through the same arithmetic.
type meterRow struct {
	GUID, Name, Class, Spec, Role string
	TimeMS, ActiveMS              int64
	Damage, Healing, Taken        int64
	Deaths                        int
}

// foldRosters folds the rosters of sums, in order. As on the web, a player's role is the one
// they first fought as, and the latest class and spec win (a mid-night respec files the player
// under what they finished as).
func foldRosters(sums []summary.Summary) []meterRow {
	var order []string
	rows := map[string]*meterRow{}
	for _, s := range sums {
		for _, r := range s.Roster {
			m, ok := rows[r.GUID]
			if !ok {
				m = &meterRow{GUID: r.GUID, Name: r.Name, Role: r.Role}
				rows[r.GUID] = m
				order = append(order, r.GUID)
			}
			m.TimeMS += s.DurationMS
			m.ActiveMS += r.ActiveMS
			m.Damage += r.DamageDone
			m.Healing += r.HealingDone
			m.Taken += r.DamageTaken
			m.Deaths += r.Deaths
			if r.Class != "" {
				m.Class = r.Class
			}
			if r.Spec != "" {
				m.Spec = r.Spec
			}
		}
	}
	out := make([]meterRow, 0, len(order))
	for _, g := range order {
		out = append(out, *rows[g])
	}
	return out
}

// Figure is a per-second figure as every app surface prints it: rounded to hundredths, the
// rounding the web's percentile query applies. Your parses and the meters both read their
// numbers through it, so a player's DPS is the same number on both.
func Figure(total int64, ms int64) float64 {
	if ms <= 0 {
		return 0
	}
	return math.Round(float64(total)/(float64(ms)/1000)*100) / 100
}

func (m meterRow) dps() float64  { return Figure(m.Damage, m.TimeMS) }
func (m meterRow) hps() float64  { return Figure(m.Healing, m.TimeMS) }
func (m meterRow) dtps() float64 { return Figure(m.Taken, m.TimeMS) }

// roleFigure is a row's headline parse metric and value, as the web's Summary tab picks it
// (ReportView roleMetric): healers on healing, everyone else, tanks too, on damage.
func (m meterRow) roleFigure() (metric string, value float64) {
	if m.Role == "healer" {
		return "hps", m.hps()
	}
	return "dps", m.dps()
}

// classSlug is the log's class word as the apps key it ("Warrior" is "warrior").
func classSlug(class string) string { return strings.ToLower(strings.TrimSpace(class)) }

// specSlug is the site's spec slug for a log's class and spec words ("Warrior", "Fury" is
// "warrior-fury"); "" when the data names no such spec.
func specSlug(class, spec string) string {
	c := classSlug(class)
	for _, s := range specs.All {
		if s.ClassSlug == c && s.Name == spec {
			return s.Spec
		}
	}
	return ""
}
