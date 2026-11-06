package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// The published set figure is one finalSetIterations run on verifySeed,
// the count the site's ranking-sim link reruns at; a healer asks for the
// quotient its runner scales back up.
func TestFinalSetIsMeasuredAtThreeThousandIterationsOnTheVerifySeed(t *testing.T) {
	bySlot := candidatesBySlot(sourceFilterFixturePool(false))
	picks := pick("warrior-arms", bySlot)
	var seen []api.SimRequest
	engine := &fakeEngine{DPSFunc: func(req api.SimRequest) (float64, error) {
		seen = append(seen, req)
		return 250, nil
	}, DefaultStdErr: 0.4}
	dps, stdErr, err := measureFinalSet(engine, specInfo{Spec: "warrior-arms", ClassSlug: "warrior"}, "human", 40, "", picks)
	if err != nil || dps != 250 || stdErr != 0.4 {
		t.Fatalf("measureFinalSet = %v +/- %v, %v", dps, stdErr, err)
	}
	if len(seen) != 1 || seen[0].Iterations != 3000 || seen[0].RandomSeed != verifySeed {
		t.Fatalf("requests = %d, iterations %d seed %d; want one at 3000 on seed %d", len(seen), seen[0].Iterations, seen[0].RandomSeed, verifySeed)
	}
	if got := gearKey(seen[0].Character.Gear); got != gearKey(buildGear(picks)) {
		t.Errorf("measured gear %s, want the final set %s", got, gearKey(buildGear(picks)))
	}
	healer := specInfo{Spec: "priest-holy", ClassSlug: "priest", Role: healerRole}
	if got := finalSetRequestIterations(healer) * healIterationScale; got < finalSetIterations || got >= finalSetIterations+healIterationScale {
		t.Errorf("a healer's final set runs at %d, want the first multiple of %d at or above %d", got, healIterationScale, finalSetIterations)
	}
}
