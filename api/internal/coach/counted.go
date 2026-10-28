package coach

import (
	"sort"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Counted is the time a fight counts against a player (§6.1): alive, outside the
// encounter's curated downtime, and from their first cast on.
type Counted struct {
	AliveMS     int64 // ratingengine.TimeAliveMS: to the first death, else the whole fight
	CountedMS   int64 // AliveMS less every Excluded window
	FirstCastMS int64 // the player's (or their pet's) first successful cast; 0 with no cast
	// Excluded is the time set aside inside [0, AliveMS): the curated downtime and the
	// time before the first cast, merged, in time order, never overlapping.
	Excluded []ratingengine.Window
}

// CountedTime is the player's counted time: time alive (ratingengine.TimeAliveMS) less
// curated downtime (ratingengine.DowntimeWindows) less the time before the first cast. A
// player with no cast at all has no first cast to measure from, so nothing is set aside
// for it.
func CountedTime(in Input) Counted {
	guid := in.Row.GUID
	alive := ratingengine.TimeAliveMS(in.Fight, guid)
	excluded := ratingengine.DowntimeWindows(in.Fight, in.Tables.Mechanics, alive)
	first := firstCastMS(in.Fight, guid)
	if first > 0 {
		excluded = append(excluded, ratingengine.Window{FromMS: 0, ToMS: min(first, alive)})
	}
	excluded = mergeWindows(excluded)
	counted := alive
	for _, w := range excluded {
		counted -= w.ToMS - w.FromMS
	}
	return Counted{AliveMS: alive, CountedMS: max(counted, 0), FirstCastMS: first, Excluded: excluded}
}

// ReviewableDeath reports whether d is reviewed: on a wipe, a death in the last
// WipeTailMS before the end is part of the wipe and is not (§6.1).
func ReviewableDeath(in Input, d summary.Death) bool {
	if in.Fight.Kill {
		return true
	}
	return d.AtMS < in.Fight.DurationMS-WipeTailMS
}

// firstCastMS is the earliest successful cast by guid or anything guid owns (a pet folds
// into its owner through CastRow.OwnerGUID); 0 when there is none.
func firstCastMS(fight summary.Summary, guid string) int64 {
	first := int64(-1)
	for _, c := range fight.Casts {
		if c.OwnerGUID != guid && c.GUID != guid {
			continue
		}
		for _, at := range c.Sequence {
			if first < 0 || at < first {
				first = at
			}
		}
	}
	return max(first, 0)
}

// mergeWindows is ws in time order with overlapping or touching windows joined.
func mergeWindows(ws []ratingengine.Window) []ratingengine.Window {
	sorted := append([]ratingengine.Window(nil), ws...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].FromMS < sorted[j].FromMS })
	var out []ratingengine.Window
	for _, w := range sorted {
		if n := len(out); n > 0 && w.FromMS <= out[n-1].ToMS {
			out[n-1].ToMS = max(out[n-1].ToMS, w.ToMS)
			continue
		}
		out = append(out, w)
	}
	return out
}
