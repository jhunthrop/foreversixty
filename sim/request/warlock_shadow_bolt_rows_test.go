package request

import "testing"

// Shadow Bolt rows re-measured level by level (sim-accuracy round 16 wave 2,
// spec-breakdown -level, band gear and talents, bare, 10,000 iterations, seeds 1
// and 50001). Under the MaxLevel + 12 penalty rank 3 keeps its full bonus to 29 and
// rank 4 to 37, so the lower rank only wins while mana is short at a band's floor:
//   - Destruction (round 24: rank 2 at 20 to 23 on the player's talents): rank 3 at 20 to 23; rank 4 from 24 (it ties at 24 and 25), rank 5
//     at 29 and 33, rank 6 at 39 (+2.8% over the old rank 4 row);
//   - Demonology: rank 3 at 20 and 21; rank 4 at 22 to 30, where it beats rank 5 by
//     up to 1.1% at 28 and 29; rank 5 from 31.
//
// Ranks: 2 = 695, 3 = 705, 4 = 1088, 5 = 1106, 6 = 7641.
func TestWarlockShadowBoltRowsFollowTheLevel(t *testing.T) {
	for _, c := range []struct {
		spec      string
		level     int
		want, not int32
	}{
		{"warlock-destruction", 23, 695, 705},
		{"warlock-destruction", 26, 1088, 705},
		{"warlock-destruction", 29, 1106, 705},
		{"warlock-destruction", 33, 1106, 1088},
		{"warlock-destruction", 39, 7641, 1088},
		{"warlock-demonology", 21, 705, 1088},
		{"warlock-demonology", 26, 1088, 705},
		{"warlock-demonology", 29, 1088, 1106},
		{"warlock-demonology", 30, 1088, 1106},
		{"warlock-demonology", 31, 1106, 1088},
	} {
		for _, race := range []string{"human", "orc"} {
			ids := rotationIDs(t, c.spec, "warlock", c.level, specRankChoices(c.spec, BarePreset, race))
			if !ids[c.want] || ids[c.not] {
				t.Errorf("%s %s at %d: want Shadow Bolt %d cast and %d not", c.spec, race, c.level, c.want, c.not)
			}
		}
	}
}
