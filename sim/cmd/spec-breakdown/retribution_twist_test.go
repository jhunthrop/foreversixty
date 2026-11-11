package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// Unsealed Wrath (1310735): "Your melee attacks store a copy of your active
// Seal ... Landing an attack with a different Seal consumes all stacks". The
// raid rotation stores three Seals of Command, switches to Seal of
// Righteousness for one swing and goes back. The way back must wait for
// that swing: while the three Command copies are still stored, a switch
// back to Seal of Command releases nothing, so the rotation would trade
// the two Seals every global cooldown until a swing happened to land
// under Righteousness, spending a global and about 200 mana each time.
//
// One cycle is three Command swings and one Righteousness swing, so a
// rotation that waits casts Seal of Righteousness about once in four
// landed swings; the swap on every global cast it about once in two.
func TestRetributionTwistWaitsForTheRighteousnessSwing(t *testing.T) {
	o := options{
		repoRoot: "../../..", spec: "paladin-retribution", preset: "raid", band: 60,
		faction: "horde", iterations: 200, seed: 1,
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	var righteousness, swings float64
	for _, a := range newReport(req, fightSecondsOf(req), est, player, nil).Actions {
		switch {
		case strings.Contains(a.Name, "(20293)"):
			righteousness += a.CastsPerIter
		case a.Name == "OtherActionAttack#1":
			swings += a.HitsPerIter
		}
	}
	if swings == 0 {
		t.Fatal("no main-hand swings landed")
	}
	if ratio := righteousness / swings; ratio > 0.4 {
		t.Errorf("Seal of Righteousness cast %.1f times a fight against %.1f landed swings (%.2f a swing), want at most 0.4: "+
			"the switch back to Seal of Command does not wait for the Righteousness swing", righteousness, swings, ratio)
	}
}
