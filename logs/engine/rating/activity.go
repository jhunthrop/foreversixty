// logs/engine/rating/activity.go
package rating

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Window is one closed-open millisecond span [FromMS, ToMS), used for curated
// downtime windows, an officer's assignment windows and idle spans.
type Window struct{ FromMS, ToMS int64 }

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

	aliveMS, windows := excludedWindows(fight, player, table, assignments)
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

// ActivityIdleMS is the time the Activity component counts against player, in its own
// measure: the time alive outside the curated downtime windows and the player's assignment
// windows, less the seconds in it the player did damage or healing (activeExcludingDowntime),
// which is the sum of IdleWindows' spans. With no window at all it is the time alive less
// the roster's ActiveMS, the whole-fight figure scoreActivity falls back to. Never
// negative; 0 for a player not on the roster.
func ActivityIdleMS(fight summary.Summary, player string, table *mechanics.Table, assignments []Assignment) int64 {
	row, ok := rosterRow(fight, player)
	if !ok {
		return 0
	}
	aliveMS, windows := excludedWindows(fight, player, table, assignments)
	if len(windows) == 0 {
		return max(aliveMS-row.ActiveMS, 0)
	}
	var idle int64
	for _, w := range idleSpans(fight, player, aliveMS, windows) {
		idle += w.ToMS - w.FromMS
	}
	return idle
}

// IdleWindows is the windows variant of ActivityIdleMS, on the same four parameters: the
// idle spans it counts (whole seconds alive, outside every downtime and assignment window,
// with no damage or healing done), joined into runs, longest first, ties earliest first.
// With windows in force they sum to ActivityIdleMS. With none, ActivityIdleMS falls back to
// the roster's ActiveMS, which the per-second spans approximate. nil for a player not on
// the roster.
func IdleWindows(fight summary.Summary, player string, table *mechanics.Table, assignments []Assignment) []Window {
	if _, ok := rosterRow(fight, player); !ok {
		return nil
	}
	aliveMS, windows := excludedWindows(fight, player, table, assignments)
	spans := idleSpans(fight, player, aliveMS, windows)
	sort.SliceStable(spans, func(i, j int) bool {
		return spans[i].ToMS-spans[i].FromMS > spans[j].ToMS-spans[j].FromMS
	})
	return spans
}

// excludedWindows is player's time alive and the downtime and assignment windows inside it.
func excludedWindows(fight summary.Summary, player string, table *mechanics.Table, assignments []Assignment) (int64, []Window) {
	aliveMS := TimeAliveMS(fight, player)
	windows := DowntimeWindows(fight, table, aliveMS)
	return aliveMS, append(windows, assignmentWindows(assignments, player, aliveMS)...)
}

// idleSpans walks the same per-second buckets as activeExcludingDowntime and joins every
// idle second (alive, outside windows, no damage or healing) into runs, in time order.
func idleSpans(fight summary.Summary, player string, aliveMS int64, windows []Window) []Window {
	active := activeSecondSet(fight, player)
	var out []Window
	for secStart := int64(0); secStart < aliveMS; secStart += 1000 {
		secEnd := min(secStart+1000, aliveMS)
		if active[int(secStart/1000)] || overlapsAny(windows, secStart, secEnd) {
			continue
		}
		if n := len(out); n > 0 && out[n-1].ToMS == secStart {
			out[n-1].ToMS = secEnd
			continue
		}
		out = append(out, Window{FromMS: secStart, ToMS: secEnd})
	}
	return out
}

func overlapsAny(windows []Window, from, to int64) bool {
	for _, w := range windows {
		if overlapsMS(w.FromMS, w.ToMS, from, to) {
			return true
		}
	}
	return false
}

// DowntimeWindows expands the encounter's curated Downtime entries into
// concrete [from, to) spans, clipped to [0, aliveMS).
func DowntimeWindows(fight summary.Summary, table *mechanics.Table, aliveMS int64) []Window {
	if table == nil {
		return nil
	}
	var out []Window
	for _, d := range table.Downtime {
		for _, at := range triggerInstantsMS(fight, d.Trigger) {
			w := Window{FromMS: at, ToMS: at + d.DurationMS}
			out = append(out, w)
		}
	}
	return clipWindows(out, aliveMS)
}

func assignmentWindows(assignments []Assignment, player string, aliveMS int64) []Window {
	var out []Window
	for _, a := range assignments {
		if a.PlayerKey != player {
			continue
		}
		out = append(out, Window{FromMS: a.FromMS, ToMS: a.ToMS})
	}
	return clipWindows(out, aliveMS)
}

func clipWindows(windows []Window, aliveMS int64) []Window {
	out := make([]Window, 0, len(windows))
	for _, w := range windows {
		if w.FromMS >= aliveMS || w.ToMS <= 0 {
			continue
		}
		if w.FromMS < 0 {
			w.FromMS = 0
		}
		if w.ToMS > aliveMS {
			w.ToMS = aliveMS
		}
		if w.ToMS > w.FromMS {
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
func activeExcludingDowntime(fight summary.Summary, player string, aliveMS int64, windows []Window) (activeMS, denomMS int64) {
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
		if overlapsAny(windows, secStartMS, secEndMS) {
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
