package main

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// castsOf runs one spec-breakdown request and returns the casts a fight of
// the player's own row whose name starts with spell (a rank's id follows the
// name, "Conflagrate (1293818)").
func castsOf(t *testing.T, o options, spell string) float64 {
	t.Helper()
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	est, player, err := inproc.PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	names, err := loadSpellNames(o.repoRoot)
	if err != nil {
		t.Fatal(err)
	}
	casts := 0.0
	for _, a := range newReport(req, est, player, names).Actions {
		if strings.HasPrefix(a.Name, spell+" (") {
			casts += a.CastsPerIter
		}
	}
	return casts
}

// Sim-accuracy round 21 (warlock lane, guide-accuracy round 19 WARLOCK-R19-3):
// before Shadow and Flame, Conflagrate always eats the Immolate it needs, so
// the Destruction card holds it until Immolate has 6 seconds or less left
// (1.7% more at 34 and 35, 0.9% at 36, the same from 38). Once Shadow and
// Flame is taken (leveling 39, raid 37) Conflagrate goes back on cooldown:
// its buff is worth more than the burn it may eat.
func TestWarlockDestructionHoldsConflagrateUntilShadowAndFlame(t *testing.T) {
	for _, tc := range []struct {
		level         int
		talents       string
		atLeast, most float64
	}{
		// Level 36, leveling order (Conflagrate at 34, no Shadow and Flame):
		// on cooldown is about 15 casts in 180 s, held about 8.
		{level: 36, talents: "0-0-20532351031011", atLeast: 6, most: 10},
		// Level 44, Shadow and Flame full: on cooldown, about 14.
		{level: 44, talents: "0-0-2053235103101351", atLeast: 12, most: 18},
	} {
		o := options{
			repoRoot: "../../..", spec: "warlock-destruction", preset: "bare",
			band: tc.level / 10 * 10, level: tc.level, faction: "horde",
			talents: tc.talents, iterations: 200, seed: 1,
		}
		if got := castsOf(t, o, "Conflagrate"); got < tc.atLeast || got > tc.most {
			t.Errorf("level %d: Conflagrate cast %.2f times a fight, want %.0f to %.0f", tc.level, got, tc.atLeast, tc.most)
		}
	}
}

// Sim-accuracy round 21 (warlock lane, guide-accuracy round 19 "Sims
// requested"): below 60 there is no Bane of Doom, so Affliction presses
// Amplify Curse before each Bane of Agony it is ready for. Since the
// 2026-11-11 warlock lane (Agony all fight, no Doom) the same holds at 60:
// a 300 s fight has room for two at either level.
func TestWarlockAfflictionAmplifyCurseBeforeEachAgony(t *testing.T) {
	below := options{
		repoRoot: "../../..", spec: "warlock-affliction", preset: "bare",
		band: 50, level: 50, faction: "horde", talents: "25522300132001351-2004-0",
		iterations: 200, seed: 1, duration: 300,
	}
	if got := castsOf(t, below, "Amplify Curse"); got < 1.9 {
		t.Errorf("level 50, 300 s: Amplify Curse cast %.2f times, want 2 (once each cooldown)", got)
	}
	at60 := options{
		repoRoot: "../../..", spec: "warlock-affliction", preset: "raid",
		band: 60, faction: "horde", iterations: 200, seed: 1, duration: 300,
	}
	if got := castsOf(t, at60, "Amplify Curse"); got < 1.9 {
		t.Errorf("level 60, 300 s: Amplify Curse cast %.2f times, want 2 (once each cooldown)", got)
	}
}
