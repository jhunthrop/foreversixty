package main

import "testing"

// Sim-accuracy round 22 (warlock lane, MythicSim lead): the Affliction raid
// build takes Demonic Sacrifice and Nightfall. It sacrifices the Imp before
// the pull (Burning Wish, 15% Shadow damage), drops Immolate once the Imp is
// gone (no Bane, no Aftermath, and it is Fire), and spends each Shadow
// Trance on an instant Shadow Bolt ahead of Wrack. A build without those
// talents (the leveling build, and the old raid build) keeps the old card:
// Immolate up, no Shadow Bolt at 60.
func TestWarlockAfflictionRaidSacrificesTheImpAndSpendsShadowTrance(t *testing.T) {
	raid := func(talents string) options {
		return options{
			repoRoot: "../../..", spec: "warlock-affliction", preset: "raid",
			band: 60, faction: "horde", talents: talents, iterations: 200, seed: 1,
		}
	}
	sac := raid("05520020135201351-0050203001-05")
	if got := castsOf(t, sac, "Demonic Sacrifice"); got < 0.99 || got > 1.01 {
		t.Errorf("sacrifice build: Demonic Sacrifice cast %.2f times, want 1 (pre-pull)", got)
	}
	if got := castsOf(t, sac, "Immolate"); got > 0 {
		t.Errorf("sacrifice build: Immolate cast %.2f times, want 0 once the Imp is sacrificed", got)
	}
	if got := castsOf(t, sac, "Shadow Bolt"); got < 3 || got > 12 {
		t.Errorf("sacrifice build: Shadow Bolt cast %.2f times, want 3 to 12 (one per Shadow Trance)", got)
	}

	old := raid("25520020135001351-00052-005004")
	if got := castsOf(t, old, "Demonic Sacrifice"); got > 0 {
		t.Errorf("no-sacrifice build: Demonic Sacrifice cast %.2f times, want 0", got)
	}
	if got := castsOf(t, old, "Immolate"); got < 5 {
		t.Errorf("no-sacrifice build: Immolate cast %.2f times, want it kept up", got)
	}
	if got := castsOf(t, old, "Shadow Bolt"); got > 0 {
		t.Errorf("no-Nightfall build: Shadow Bolt cast %.2f times, want 0 (Wrack is the filler)", got)
	}
}

// Sim-accuracy round 22 (warlock lane, guide-accuracy round 23 WARLOCK-R23-4):
// the Demonology raid build skips Soul Fire inside Decimation at raid spell
// power, where its Shadow bonuses make Shadow Bolt stronger, but from Soul
// Fire rank 2 at 56 to 59, on band gear, the Soul Fires win (+0.2 to +0.5%).
func TestWarlockDemonologyRaidSoulFireInsideDecimationAt56To59(t *testing.T) {
	at := func(level int, talents string) options {
		return options{
			repoRoot: "../../..", spec: "warlock-demonology", preset: "bare",
			band: 50, level: level, faction: "horde", talents: talents,
			iterations: 200, seed: 1,
		}
	}
	if got := castsOf(t, at(57, "05-2035103021220001351-0550002"), "Soul Fire"); got < 1 {
		t.Errorf("raid order at 57: Soul Fire cast %.2f times, want it inside Decimation", got)
	}
	if got := castsOf(t, at(55, "05-2035103021220001351-055"), "Soul Fire"); got > 0 {
		t.Errorf("raid order at 55: Soul Fire cast %.2f times, want 0 (rank 1 loses)", got)
	}
	raid := options{
		repoRoot: "../../..", spec: "warlock-demonology", preset: "raid",
		band: 60, faction: "horde", talents: "05-2035103021220001351-0550005",
		iterations: 200, seed: 1,
	}
	if got := castsOf(t, raid, "Soul Fire"); got > 0 {
		t.Errorf("raid 60: Soul Fire cast %.2f times, want 0", got)
	}
}
