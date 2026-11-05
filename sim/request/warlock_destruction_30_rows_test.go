package request

import "testing"

// Destruction at 30 and 31 (sim-accuracy round 20, spec-breakdown -level, band 30
// gear and talents, bare, 10,000 iterations, seeds 1 and 50001, human and orc).
// Shadow Bolt rank 5 (1106) arrives at 28 and the band-30 fight is short of
// mana (2.7 s out of mana, 30 Life Taps): rank 4 (1088, 110 mana against 160)
// beats it by 0.2 to 0.3% at 30 and 0.1 to 0.2% at 31 on both races. At 32 the
// two tie on human, and rank 5 wins from 33.
func TestWarlockDestructionShadowBoltRank4At30And31(t *testing.T) {
	for _, c := range []struct {
		level     int
		want, not []int32
	}{
		{29, []int32{1106}, []int32{1088}},
		{30, []int32{1088}, []int32{1106}},
		{31, []int32{1088}, []int32{1106}},
		{32, []int32{1106}, []int32{1088}},
	} {
		for _, race := range []string{"human", "orc"} {
			ids := rotationIDs(t, "warlock-destruction", "warlock", c.level, specRankChoices("warlock-destruction", BarePreset, race))
			for _, id := range c.want {
				if !ids[id] {
					t.Errorf("destruction %s at %d: want %d cast", race, c.level, id)
				}
			}
			for _, id := range c.not {
				if ids[id] {
					t.Errorf("destruction %s at %d: want %d not cast", race, c.level, id)
				}
			}
		}
	}
}
