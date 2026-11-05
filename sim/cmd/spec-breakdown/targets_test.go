package main

import "testing"

// -targets sets a pack fight's target count; without it the encounter keeps
// the ranker's single target.
func TestTargetsFlagSetsThePackSize(t *testing.T) {
	for _, c := range []struct{ flag, want int }{{0, 1}, {3, 3}} {
		o := options{
			repoRoot: "../../..", spec: "druid-feral-bear", preset: "raid", band: 60,
			faction: "horde", iterations: 10, seed: 1, targets: c.flag,
		}
		req, err := buildRequest(o)
		if err != nil {
			t.Fatal(err)
		}
		if req.Encounter.Targets != c.want {
			t.Errorf("-targets %d gives %d targets, want %d", c.flag, req.Encounter.Targets, c.want)
		}
	}
}
