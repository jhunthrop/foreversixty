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
