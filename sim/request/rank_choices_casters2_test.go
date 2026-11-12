package request

import "testing"

// Sim-accuracy 2026-11-12 (casters lane, design/reviews/2026-11-12-sim-accuracy/casters.md): the
// mage walks at 23 to 27 re-judged on the engine that gives a mage its first mana gem at 28
// (ledger MR67). Fire's Fireball rank 4 (3140) keeps +0.2% on rank 5 (8400) for a Human at 25
// and for a Human or a Gnome at 27, at both seeds; every other race, and every race at 26, casts
// rank 5. Fire Blast rank 3 and Frost's Frostbolt rank 4 win at every level there, so no row.
func TestCastersTwoFireballRowsAtTwentyFiveAndTwentySeven(t *testing.T) {
	for _, c := range []struct {
		name, race string
		level      int
		want, not  int32
	}{
		{"Human Fireball at 25", "human", 25, 3140, 8400},
		{"Gnome Fireball at 25", "gnome", 25, 8400, 3140},
		{"Undead Fireball at 25", "undead", 25, 8400, 3140},
		{"Human Fireball at 26", "human", 26, 8400, 3140},
		{"Human Fireball at 27", "human", 27, 3140, 8400},
		{"Gnome Fireball at 27", "gnome", 27, 3140, 8400},
		{"Troll Fireball at 27", "troll", 27, 8400, 3140},
		{"Human Fireball at 28", "human", 28, 8400, 3140},
	} {
		ids := rotationIDs(t, "mage-fire", "mage", c.level, specRankChoices("mage-fire", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
