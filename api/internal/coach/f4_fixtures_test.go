package coach

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// F4 fixtures. Rage is in tenths, as the log's rage series is: the bar holds 1000.
const (
	f4RageMax     = 1000
	f4EnergyMax   = 100
	f4ManaMax     = 10_000
	f4FirstCastMS = 1_000
	f4WindowSecs  = 1_000
)

// f4WithoutResources is s with every resource track of guid removed, so a fixture's
// tracks are exactly the ones it sets.
func f4WithoutResources(s summary.Summary, guid string) summary.Summary {
	s.Resources = slices.DeleteFunc(slices.Clone(s.Resources), func(r summary.ResourceTrack) bool { return r.GUID == guid })
	return s
}

// f4WithoutDeaths is s with every death of guid removed.
func f4WithoutDeaths(s summary.Summary, guid string) summary.Summary {
	s.Deaths = slices.DeleteFunc(slices.Clone(s.Deaths), func(d summary.Death) bool { return d.GUID == guid })
	return s
}

// f4Prepared is in's fight lasting durationMS with its player casting from the first
// second, never dying, and holding no resource track.
func f4Prepared(in Input, durationMS int64) Input {
	guid := in.Row.GUID
	s := withDuration(f4WithoutDeaths(f4WithoutResources(in.Fight, guid), guid), durationMS)
	in.Fight = withCasts(s, guid, 23881, "Bloodthirst", f4FirstCastMS)
	return in
}

// f4Series is one reading per second over durationMS: base everywhere, and special in
// each window.
func f4Series(durationMS, base, special int64, windows ...summary.Segment) []int64 {
	out := repeat(base, int(durationMS/f4WindowSecs))
	for _, w := range windows {
		for ms := w.StartMS; ms < w.EndMS; ms += f4WindowSecs {
			out[ms/f4WindowSecs] = special
		}
	}
	return out
}

// f4WithTrack is in with the player holding one track of powerType.
func f4WithTrack(in Input, powerType, max int64, series []int64) Input {
	in.Fight = withResource(in.Fight, in.Row.GUID, powerType, max, series...)
	return in
}

// furyRageCapped is the Fury warrior with rage at its cap in each window.
func furyRageCapped(t *testing.T, durationMS int64, windows ...summary.Segment) Input {
	t.Helper()
	return f4Capped(t, PowerRage, f4RageMax, durationMS, windows...)
}

// f4Capped is the Fury warrior with the given power type held at its cap in each window
// and at half of it elsewhere.
func f4Capped(t *testing.T, powerType, max, durationMS int64, windows ...summary.Segment) Input {
	t.Helper()
	in := f4Prepared(furyInput(t), durationMS)
	return f4WithTrack(in, powerType, max, f4Series(durationMS, max/2, max, windows...))
}

// holyManaLow is the Holy priest with mana under 5% from fromMS to toMS.
func holyManaLow(t *testing.T, durationMS, fromMS, toMS int64) Input {
	t.Helper()
	in := f4Prepared(holyInput(t), durationMS)
	return f4WithTrack(in, PowerMana, f4ManaMax, f4Series(durationMS, f4ManaMax*4/5, f4ManaMax/100, span(fromMS, toMS)))
}
