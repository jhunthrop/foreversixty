package request

import "testing"

// Destruction at 40 and 41 (sim-accuracy round 19, spec-breakdown -level, band 40
// gear and talents, bare, 10,000 iterations, seeds 1 and 50001, human and orc).
// The new Immolate rank 5 and Conflagrate rank 3 arrive together at 40 on band-40
// gear, the fight is short of mana, and a cheaper rank wins for a level each:
//   - 40: Immolate rank 4 beats rank 5 by 0.1 to 0.3% (orc) and 0.3% (human);
//   - 41: Corruption rank 3 beats rank 4 by 0.1 to 0.2% (orc) and 0.4% (human).
//
// Both together are no better than either alone, and the newest ranks win again
// from 42. Ranks: Immolate 4 = 2941, 5 = 11665; Corruption 3 = 6223, 4 = 7648.
func TestWarlockDestructionCheaperDoTsAt40And41(t *testing.T) {
	for _, c := range []struct {
		level     int
		want, not []int32
	}{
		{40, []int32{2941, 7648}, []int32{11665, 6223}},
		{41, []int32{11665, 6223}, []int32{2941, 7648}},
		{42, []int32{11665, 7648}, []int32{2941, 6223}},
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
