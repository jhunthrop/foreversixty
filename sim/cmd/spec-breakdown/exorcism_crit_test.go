package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// Exorcism never crits in Forever: the public beta logs on foreverlogs.gg
// (39 paladins at levels 20 to 30, fetched 2026-11-02) show 0 crits in 641
// Exorcism hits, where the same paladins' Judgement of Righteousness crit
// 18 times in 367 and their spell crit would have given about 35. The
// client row (879 to 10314) states no "cannot crit" attribute, so this is
// server behaviour. Against an Undead target, where the rotation casts it,
// the raid Retribution paladin's Exorcism must never crit.
func TestExorcismNeverCrits(t *testing.T) {
	o := options{
		repoRoot: "../../..", spec: "paladin-retribution", preset: "raid", band: 60,
		faction: "horde", iterations: 200, seed: 1, targetType: "undead",
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, a := range newReport(req, fightSecondsOf(req), est, player, nil).Actions {
		if !strings.Contains(a.Name, "(10314)") {
			continue
		}
		found = true
		if a.CritPct != 0 {
			t.Errorf("Exorcism crit %.1f%% of its casts, want 0", a.CritPct)
		}
	}
	if !found {
		t.Fatal("Exorcism (10314) was never cast against an Undead target")
	}
}
