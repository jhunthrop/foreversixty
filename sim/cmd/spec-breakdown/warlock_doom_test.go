package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// Sim-accuracy round 15 (warlock lane, guide-accuracy round 1 S3): Bane of
// Doom (603: 1742 plus 4.0 times spell power after 60 s, one global, 300 mana)
// beats Bane of Agony for Affliction too, as it does for Demonology and
// Destruction. Doom goes up while a minute or more of the fight is left, Agony
// covers the last minute, and Amplify Curse waits for that Agony: the bane slot
// holds one bane, so a pre-pull Amplify Curse ran out on Doom unused.
func TestAfflictionRaidRotationKeepsBaneOfDoomAndAmplifiesTheLastAgony(t *testing.T) {
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
	r := newReport(req, est, player, nil)
	casts := map[string]float64{}
	for _, a := range r.Actions {
		for _, id := range []string{"(603)", "(11713)", "(18288)"} {
			if strings.HasSuffix(a.Name, id) {
				casts[id] = a.CastsPerIter
			}
		}
	}
	if casts["(603)"] < 1.9 {
		t.Errorf("Bane of Doom cast %.2f times a %.0f s fight, want 2 (at the pull and a minute later)", casts["(603)"], r.FightSeconds)
	}
	if casts["(11713)"] < 1.5 || casts["(11713)"] > 3.5 {
		t.Errorf("Bane of Agony cast %.2f times, want two or three 24 s casts covering the last minute", casts["(11713)"])
	}
	amplify := 0.0
	for _, a := range r.Auras {
		if strings.HasSuffix(a.Name, "(18288)") {
			amplify = a.UptimePct
		}
	}
	if casts["(18288)"] < 0.9 || amplify > 5 {
		t.Errorf("Amplify Curse cast %.2f times with %.2f%% uptime, want it cast once and spent at once on the last Agony", casts["(18288)"], amplify)
	}
}
