package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// The page's trust line reads these from the file; they must be what the
// ranker's requests really carry.
func TestRunConditionsAreTheRankerRequests(t *testing.T) {
	dps := specInfo{Spec: "warrior-arms", ClassSlug: "warrior"}
	rc := rankerRunConditions(dps, 100)
	req := plainRequest(dps, bandCharacter("x", "human", "warrior", dps.Spec, 30, "", nil), verifyIterations, verifySeed)
	e := req.Encounter
	if rc.TargetType != e.TargetType || rc.TargetType != api.TargetTypeUnknown || rc.DurationSec != e.DurationSec || rc.Variation != e.Variation || rc.Targets != e.Targets || rc.ExecuteRatio != e.ExecuteRatio {
		t.Errorf("run = %+v, request encounter = %+v", rc, e)
	}
	if rc.DurationSec != 180 || rc.TargetLevelOffset != 3 || rc.Seed != req.RandomSeed {
		t.Errorf("run = %+v, want 180 s, +3 levels, the verify seed", rc)
	}
	if rc.Iterations != (runIterations{Set: 3000, Verify: 300, Guard: verifyIterations * healIterationScale, TrinketRank: 100, WeightsPerDirection: 100 * api.WeightsIterationsFactor}) || rc.HealIterationScale != 0 {
		t.Errorf("iterations = %+v scale %d", rc.Iterations, rc.HealIterationScale)
	}
	healer := rankerRunConditions(specInfo{Spec: "priest-holy", Role: healerRole}, 100)
	if healer.HealIterationScale != healIterationScale || healer.Iterations.Guard != verifyIterations || healer.Iterations.Set != 3003 {
		t.Errorf("healer run = %+v", healer)
	}
	for _, band := range []int{20, 40, 60} {
		if got := bandTargetLevel(band); got != band+3 || leveling.RankerEncounter().TargetLevel != 0 {
			t.Errorf("band %d target level = %d, want %d", band, got, band+3)
		}
	}
}
