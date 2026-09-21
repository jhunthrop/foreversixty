// logs/engine/rating/fixtures_test.go
// Shared test fixtures and a fake PercentileSource, used by every _test.go
// file in this package.
package rating

import "github.com/jhunthrop/foreversixty/logs/engine/summary"

const fixturePlayer = "Player-1-00000001"

// fixtureSummary is the shared minimal fixture most tests build on and
// mutate via their own helpers — a one-player, 180-second fight on
// Shazzrah (encounter_id 667), matching spec §1.6's worked examples' own
// setup (a 3-minute Molten Core kill, typical band).
func fixtureSummary() summary.Summary {
	return summary.Summary{
		EngineVersion: "test", DurationMS: 180000, EncounterID: 667, Difficulty: 0, Kill: true,
		Roster: []summary.RosterRow{
			{GUID: fixturePlayer, Name: "Fixture", Class: "Warrior", Spec: "Fury", Role: RoleDPS,
				DPS: 500, ActiveMS: 163800, ActivityPct: 91},
		},
	}
}

// fakePercentiles is a hand-fed, RECORDING PercentileSource: every
// component test supplies exactly the (bracket.Component -> pct/n/ok)
// triples its scenario needs, so §1.2's percentile-vs-absolute rule is
// exercised without any database -- but every value it is asked to place
// is also recorded, keyed by bracket.Component, in call order. A fake that
// only returns canned answers and ignores what it was asked can hide a
// real formula bug behind a still-correct-looking final score (the
// whole-branch review caught exactly this on the Survival avoidable-hit
// sub-part: the fake's canned 35th-percentile answer passed regardless of
// whether the real computed value was the excused 0 or the fully-counted
// 22.22); a test that cares about the underlying computation, not just the
// final display number, asserts against src.recorded[...] instead of only
// the Card's own Score field. Pointer receiver: the recording needs
// mutable state, so every caller constructs *fakePercentiles.
type fakePercentiles struct {
	placements map[string]fakePlacement
	band       string
	bandOK     bool
	recorded   map[string][]float64
}

type fakePlacement struct {
	pct float64
	n   int64
	ok  bool
}

func (f *fakePercentiles) Placement(bracket Bracket, value float64) (float64, int64, bool) {
	if f.recorded == nil {
		f.recorded = map[string][]float64{}
	}
	f.recorded[bracket.Component] = append(f.recorded[bracket.Component], value)
	p, ok := f.placements[bracket.Component]
	if !ok {
		return 0, 0, false
	}
	return p.pct, p.n, p.ok
}

func (f *fakePercentiles) KillTimeBand(encounterID, difficulty int64, durationMS int64) (string, bool) {
	return f.band, f.bandOK
}

// roundTo1 is roundTo(x, 1), used throughout the tests to compare against
// the spec's own worked-example precision without a flaky float equality
// check.
func roundTo1(x float64) float64 { return roundTo(x, 1) }
