package request

import "testing"

// Round 23 (priest lane): Shadow Word: Pain's rows started at 30, so at 26
// to 29 Shadow cast rank 4, the newest, which loses to rank 3 at 26 and 27
// and to rank 2 at 28 (and 29 for most races), where Mind Flay arrives and
// the fight runs dry; Devouring Plague's rank 1 rows follow the cheaper
// Pain (design/reviews/2026-11-05-sim-audit-round23/priest.md).
func TestRound23ShadowPainAndPlagueRankRows(t *testing.T) {
	const swp2, swp3, swp4 = 594, 970, 992
	const dp1, dp2 = 2944, 19276
	for _, c := range []struct {
		name, race string
		level      int
		want, not  int32
	}{
		{"Pain at 26", "human", 26, swp3, swp4},
		{"Pain at 27", "troll", 27, swp3, swp4},
		{"Pain at 28", "undead", 28, swp2, swp4},
		{"Pain Dwarf at 29", "dwarf", 29, swp2, swp4},
		{"Pain Human at 29", "human", 29, swp3, swp2},
		{"Pain at 30 holds rank 3", "gnome", 30, swp3, swp2},
		{"Plague Night Elf at 28", "night-elf", 28, dp1, dp2},
		{"Plague Dwarf at 29", "dwarf", 29, dp1, dp2},
		{"Plague Dwarf at 28 is the top rank", "dwarf", 28, dp2, dp1},
		{"Plague Human at 36", "human", 36, dp2, dp1},
		{"Plague Troll at 37", "troll", 37, dp2, dp1},
		{"Plague Gnome at 36 holds rank 1", "gnome", 36, dp1, dp2},
	} {
		ids := rotationIDs(t, "priest-shadow", "priest", c.level, specRankChoices("priest-shadow", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s (%s): want %d cast and %d not", c.name, c.race, c.want, c.not)
		}
	}
}

// Round 23 (priest lane): the newest Mind Blast and Shadow Word: Death were
// cast the level they arrived. Rank 4 Mind Blast beats rank 5 at 34 (and
// for some races at 35 and 36); Shadow Word: Death rank 2 beats rank 3 at 48
// and 49 for every race, as it already did from 50, and rank 1 beats rank 2
// at 40 for a Dwarf, Human or Night Elf.
func TestRound23ShadowMindBlastAndDeathRankRows(t *testing.T) {
	const mb4, mb5 = 8104, 8105
	const swd1, swd2, swd3 = 1309595, 1309633, 1309635
	for _, c := range []struct {
		name, race string
		level      int
		want, not  int32
	}{
		{"Mind Blast at 34", "human", 34, mb4, mb5},
		{"Mind Blast Undead at 35", "undead", 35, mb4, mb5},
		{"Mind Blast Human at 35 is the top rank", "human", 35, mb5, mb4},
		{"Mind Blast Night Elf at 36", "night-elf", 36, mb4, mb5},
		// Sim-accuracy 2026-11-12 (shadow): on the wand's first-shot windup a Human casts rank 2 at 40.
		{"Death Human at 40 is the top rank", "human", 40, swd2, swd1},
		{"Death Dwarf at 40", "dwarf", 40, swd1, swd2},
		{"Death Gnome at 40 is the top rank", "gnome", 40, swd2, swd1},
		{"Death Night Elf at 41", "night-elf", 41, swd1, swd2},
		{"Death Gnome at 43", "gnome", 43, swd1, swd2},
		{"Death at 48", "troll", 48, swd2, swd3},
		{"Death at 49", "dwarf", 49, swd2, swd3},
	} {
		ids := rotationIDs(t, "priest-shadow", "priest", c.level, specRankChoices("priest-shadow", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s (%s): want %d cast and %d not", c.name, c.race, c.want, c.not)
		}
	}
}

// Round 23 (priest lane): a new shield rank costs more than the extra it
// soaks for a few levels on band gear, race by race (Discipline rank 5 at 36
// to 40, rank 4 at 30 for an Undead, rank 8 at 54 and 55 for a Human), and a
// Holy Undead keeps Flash Heal rank 2 at 32.
func TestRound23PriestShieldAndFlashHealRankRows(t *testing.T) {
	const pws4, pws5, pws6, pws8, pws9 = 3747, 6065, 6066, 10899, 10900
	const fh2, fh3 = 9472, 9473
	for _, c := range []struct {
		name, spec, race string
		level            int
		want, not        int32
	}{
		{"Undead shield at 30", "priest-discipline", "undead", 30, pws4, pws5},
		{"Human shield at 30 is the top rank", "priest-discipline", "human", 30, pws5, pws4},
		{"Troll shield at 36", "priest-discipline", "troll", 36, pws5, pws6},
		{"Dwarf shield at 37", "priest-discipline", "dwarf", 37, pws5, pws6},
		{"Gnome shield at 38 is the top rank", "priest-discipline", "gnome", 38, pws6, pws5},
		{"Human shield at 40", "priest-discipline", "human", 40, pws5, pws6},
		{"Human shield at 55", "priest-discipline", "human", 55, pws8, pws9},
		{"Holy Undead Flash Heal at 32", "priest-holy", "undead", 32, fh2, fh3},
		{"Holy Human Flash Heal at 32 is the top rank", "priest-holy", "human", 32, fh3, fh2},
	} {
		ids := rotationIDs(t, c.spec, "priest", c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
