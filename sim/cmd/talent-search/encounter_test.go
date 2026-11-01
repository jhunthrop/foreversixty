package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// The search's baseline must be the number spec-breakdown and the ranker
// publish for the same build. It fought the page's default target (a
// Humanoid), where Improved Tracking and Murder count, while they fight a
// type-neutral one: Marksmanship's guide build read 647.0 here and 619.5
// in spec-breakdown (round-2 physical audit, M3).
func TestSearchFightsTheRankersTypeNeutralTarget(t *testing.T) {
	s := simSetup{spec: specInfo{Spec: "hunter-marksmanship", ClassSlug: "hunter"}, band: bisBand{Race: "dwarf"}, level: 60, seed: 1}
	req := s.request("0", 10)
	if req.Encounter.TargetType != api.TargetTypeUnknown {
		t.Errorf("search target type = %q, want the ranker's %q", req.Encounter.TargetType, api.TargetTypeUnknown)
	}
	if req.Iterations != 10 || req.RandomSeed != 1 || req.Character.Talents != "0" {
		t.Errorf("request = %+v", req)
	}
}
