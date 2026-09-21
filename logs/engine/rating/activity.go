// logs/engine/rating/activity.go
package rating

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// window is one closed-open millisecond span, used for both curated
// downtime windows and an officer's assignment windows.
type window struct{ fromMS, toMS int64 }

// scoreActivity implements spec §1.3's Activity component. With no curated
// downtime windows (and no assignment windows) it falls back to the
// whole-fight RosterRow.ActivityPct, flagged approximate but never
// excluded ("Activity is the one component that is always computable").
func scoreActivity(fight summary.Summary, player string, bracket Bracket, table *mechanics.Table, assignments []Assignment, src PercentileSource) Component {
	bracket.Component = ComponentNameActivity
	c := Component{Name: ComponentNameActivity}

	row, ok := rosterRow(fight, player)
	if !ok {
		return placeBoundedOrSelf(c, src, bracket, 0)
	}

	aliveMS := timeAliveMS(fight, player)
	windows := downtimeWindows(fight, table, aliveMS)
	windows = append(windows, assignmentWindows(assignments, player, aliveMS)...)

	if len(windows) == 0 {
		return placeBoundedOrSelf(c, src, bracket, clamp(row.ActivityPct, 0, 100))
	}

	activeMS, denomMS := activeExcludingDowntime(fight, player, aliveMS, windows)
	if denomMS <= 0 {
		return placeBoundedOrSelf(c, src, bracket, clamp(row.ActivityPct, 0, 100))
	}
	share := clamp(float64(activeMS)/float64(denomMS)*100, 0, 100)
	return placeBoundedOrSelf(c, src, bracket, share)
}

// downtimeWindows expands the encounter's curated Downtime entries into
// concrete [from, to) spans, clipped to [0, aliveMS).
func downtimeWindows(fight summary.Summary, table *mechanics.Table, aliveMS int64) []window {
	if table == nil {
		return nil
	}
	var out []window
	for _, d := range table.Downtime {
		for _, at := range triggerInstantsMS(fight, d.Trigger) {
			w := window{fromMS: at, toMS: at + d.DurationMS}
			out = append(out, w)
		}
	}
	return clipWindows(out, aliveMS)
}

func assignmentWindows(assignments []Assignment, player string, aliveMS int64) []window {
	var out []window
	for _, a := range assignments {
		if a.PlayerKey != player {
			continue
		}
		out = append(out, window{fromMS: a.FromMS, toMS: a.ToMS})
	}
	return clipWindows(out, aliveMS)
}

func clipWindows(windows []window, aliveMS int64) []window {
	out := make([]window, 0, len(windows))
	for _, w := range windows {
		if w.fromMS >= aliveMS || w.toMS <= 0 {
			continue
		}
		if w.fromMS < 0 {
			w.fromMS = 0
		}
		if w.toMS > aliveMS {
			w.toMS = aliveMS
		}
		if w.toMS > w.fromMS {
			out = append(out, w)
		}
	}
	return out
}

func overlapsMS(aFrom, aTo, bFrom, bTo int64) bool {
	return aFrom < bTo && bFrom < aTo
}

// activeExcludingDowntime is RULING R8's per-second-bucket approximation:
// markActive's exact sub-second logic does not survive into Summary, so
// each whole second the player did damage or healing (Actor.Series is
// already one entry per second) counts as active, excluding any second
// that overlaps a downtime or assignment window.
func activeExcludingDowntime(fight summary.Summary, player string, aliveMS int64, windows []window) (activeMS, denomMS int64) {
	activeSeconds := activeSecondSet(fight, player)
	totalSeconds := int(aliveMS / 1000)
	if aliveMS%1000 != 0 {
		totalSeconds++
	}
	var overlapMS int64
	for sec := 0; sec < totalSeconds; sec++ {
		secStartMS := int64(sec) * 1000
		secEndMS := secStartMS + 1000
		if secEndMS > aliveMS {
			secEndMS = aliveMS
		}
		inDowntime := false
		for _, w := range windows {
			if overlapsMS(w.fromMS, w.toMS, secStartMS, secEndMS) {
				inDowntime = true
				break
			}
		}
		if inDowntime {
			overlapMS += secEndMS - secStartMS
			continue
		}
		if activeSeconds[sec] {
			activeMS += secEndMS - secStartMS
		}
	}
	return activeMS, aliveMS - overlapMS
}

func activeSecondSet(fight summary.Summary, player string) map[int]bool {
	set := map[int]bool{}
	for _, a := range fight.DamageDone {
		if a.GUID == player {
			markNonZeroSeconds(set, a.Series)
		}
	}
	for _, a := range fight.Healing {
		if a.GUID == player {
			markNonZeroSeconds(set, a.Series)
		}
	}
	return set
}

func markNonZeroSeconds(set map[int]bool, series []int64) {
	for i, v := range series {
		if v != 0 {
			set[i] = true
		}
	}
}

// triggerInstantsMS finds every instant a PhaseStart-shaped trigger fired,
// from data already in Summary (RULING R7): "aura_applied"/"aura_removed"
// read AuraTrack.Segments' start/end instants for the triggering spell id
// (on any track -- a downtime trigger is usually an enemy's own aura, and
// Summary does not tag a track as enemy-vs-player by a dedicated field),
// "cast_success" reads CastRow.Sequence. "cast_start" and a health_pct
// trigger have no equivalent data surviving Snapshot and produce no
// instants -- documented, not guessed at.
func triggerInstantsMS(fight summary.Summary, trigger mechanics.PhaseStart) []int64 {
	if trigger.SpellID <= 0 {
		return nil
	}
	var out []int64
	switch trigger.On {
	case mechanics.OnAuraApplied:
		for _, tr := range fight.Auras {
			if tr.SpellID != trigger.SpellID {
				continue
			}
			for _, seg := range tr.Segments {
				out = append(out, seg.StartMS)
			}
		}
	case mechanics.OnAuraRemoved:
		for _, tr := range fight.Auras {
			if tr.SpellID != trigger.SpellID {
				continue
			}
			for _, seg := range tr.Segments {
				out = append(out, seg.EndMS)
			}
		}
	case mechanics.OnCastSuccess:
		for _, cr := range fight.Casts {
			if cr.SpellID != trigger.SpellID {
				continue
			}
			out = append(out, cr.Sequence...)
		}
	}
	return out
}
