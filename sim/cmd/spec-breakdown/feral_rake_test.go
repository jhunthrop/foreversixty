package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// The cat keeps Rake up (sim-accuracy 2026-11-12 feral C1): its hit is the
// paw plus the rank's amount and its ticks take attack power, so keeping it
// on the boss measured +19 to +20% on the bare bands from 30 and +10.6% at
// raid 60 against the Shred-only list. On the raid entry Rake is cast once
// every 9 seconds or so and Shred still fills the other globals.
func TestFeralCatKeepsRakeUp(t *testing.T) {
	o := options{repoRoot: "../../..", spec: "druid-feral", preset: "raid", band: 60, faction: "horde", iterations: 30, seed: 1}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	r := newReport(req, fightSecondsOf(req), est, player, nil)
	rake, shred := 0.0, 0.0
	for _, a := range r.Actions {
		switch {
		case strings.HasSuffix(a.Name, "(9904)"):
			rake += a.CastsPerIter
		case strings.HasSuffix(a.Name, "(9830)"):
			shred += a.CastsPerIter
		}
	}
	secs := fightSecondsOf(req)
	if rake < secs/12 || rake > secs/8 {
		t.Errorf("Rake %.1f casts in a %.0f s fight; want one about every 9 to 12 s", rake, secs)
	}
	if shred < rake {
		t.Errorf("Shred %.1f casts against Rake %.1f; want Shred as the filler", shred, rake)
	}
}
