package main

import (
	"cmp"
	"context"
	"slices"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Fight is one stored fight the evaluation reads: its summary and the facts the pool
// bracket needs. ReportID and PlayerKeys are identifiers; they stay in this process and
// in fights.private.json, never in the report or the JSONL.
type Fight struct {
	ReportID   string
	FightIndex int
	Size       int
	FoughtAt   time.Time
	Summary    summary.Summary
	// PlayerKeys maps a roster name to its character key; empty offline, where the pool is skipped.
	PlayerKeys map[string]string
}

// Source yields up to n fights, read only.
type Source interface {
	Fights(ctx context.Context, n int) ([]Fight, error)
}

// Target is one fight and the player whose review is evaluated.
type Target struct {
	Label string
	Fight Fight
	Row   summary.RosterRow
}

// planTargets picks one player per fight, spreading the picks over specs: each fight
// takes its roster's least-used spec so far, the busiest player of it on a tie. Fights
// with no spec'd player are dropped.
func planTargets(fights []Fight, labelFor func(i int) string) []Target {
	used := map[string]int{}
	var out []Target
	for _, f := range fights {
		row, ok := pickPlayer(f.Summary.Roster, used)
		if !ok {
			continue
		}
		used[row.Spec]++
		out = append(out, Target{Label: labelFor(len(out)), Fight: f, Row: row})
	}
	return out
}

func pickPlayer(roster []summary.RosterRow, used map[string]int) (summary.RosterRow, bool) {
	var cands []summary.RosterRow
	for _, r := range roster {
		if r.Spec != "" && r.Class != "" {
			cands = append(cands, r)
		}
	}
	if len(cands) == 0 {
		return summary.RosterRow{}, false
	}
	slices.SortFunc(cands, func(a, b summary.RosterRow) int {
		return cmp.Or(
			cmp.Compare(used[a.Spec], used[b.Spec]),
			cmp.Compare(b.ActiveMS, a.ActiveMS),
			cmp.Compare(a.Name, b.Name),
		)
	})
	return cands[0], true
}
