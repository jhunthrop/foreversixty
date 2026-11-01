package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// A rotation is searched against the ranker's type-neutral target, the one
// the published numbers fight, not the page's default Humanoid, where
// Murder and Improved Tracking add a flat few percent the published number
// does not have.
func TestRotationSearchFightsTheRankersTypeNeutralTarget(t *testing.T) {
	s := simSetup{spec: specInfo{Spec: "rogue-combat"}, seed: 1}
	if got := s.request(10).Encounter.TargetType; got != api.TargetTypeUnknown {
		t.Errorf("target type = %q, want %q", got, api.TargetTypeUnknown)
	}
}
