package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// Sim-accuracy round 15 kept Bane of Doom for Affliction and held Amplify
// Curse for the last-minute Agony. The 2026-11-11 warlock lane (10 Nov
// hotfix: Soul Siphon counts only Corruption, Siphon Life and Bane of Agony,
// and Wrack speeds those three up) moved the card to Agony all fight with no
// Doom, +5.1 to +5.4% at raid 60 (design/reviews/2026-11-11-sim-accuracy/
// warlock.md). Amplify Curse now goes on the next Agony whenever it is ready
// and is spent at once.
func TestAfflictionRaidRotationRunsAgonyAllFightWithoutDoom(t *testing.T) {
	_, trees, err := leveling.GuideRaidBuildTalents("../../..", "warlock", "affliction")
	if err != nil {
		t.Fatal(err)
	}
	o := options{
		repoRoot: "../../..", spec: "warlock-affliction", preset: "raid", band: 60,
		faction: "horde", iterations: 100, seed: 1, talents: strings.Join(trees[:], "-"),
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	r := newReport(req, fightSecondsOf(req), est, player, nil)
	casts := map[string]float64{}
	for _, a := range r.Actions {
		for _, id := range []string{"(603)", "(11713)", "(18288)"} {
			if strings.HasSuffix(a.Name, id) {
				casts[id] = a.CastsPerIter
			}
		}
	}
	if casts["(603)"] > 0 {
		t.Errorf("Bane of Doom cast %.2f times a %.0f s fight, want none (Agony holds the bane slot)", casts["(603)"], r.FightSeconds)
	}
	// A 24 s Agony kept up from the pull is at least one cast per 24 s; the
	// under-3 s recast before a Wrack channel adds more.
	if minimum := r.FightSeconds / 24; casts["(11713)"] < minimum {
		t.Errorf("Bane of Agony cast %.2f times a %.0f s fight, want at least %.1f (up all fight)", casts["(11713)"], r.FightSeconds, minimum)
	}
	amplify := 0.0
	for _, a := range r.Auras {
		if strings.HasSuffix(a.Name, "(18288)") {
			amplify = a.UptimePct
		}
	}
	if casts["(18288)"] < 0.9 || amplify > 5 {
		t.Errorf("Amplify Curse cast %.2f times with %.2f%% uptime, want it cast and spent at once on the next Agony", casts["(18288)"], amplify)
	}
}
