package request

import "testing"

// Sim-accuracy round 24, second pass (on the talents at each level, judged
// with the new Consecration rows): rank 2 (24274) beats rank 1 (24275) at 52
// and at 56, so those rows are gone; rank 1 wins at 53 for every race and at
// 54 for an Undead; every other level keeps the trained top rank
// (rank-choices.json notes carry the numbers).
func TestRetributionHammerOfWrathRankFollowsTheSweep(t *testing.T) {
	for _, c := range []struct {
		race  string
		level int
		want  int32
	}{
		{"human", 52, 24274},
		{"undead", 52, 24274},
		{"human", 53, 24275},
		{"dwarf", 53, 24275},
		{"undead", 53, 24275},
		{"human", 54, 24274},
		{"undead", 54, 24275},
		{"human", 56, 24274},
		{"undead", 56, 24274},
		{"human", 57, 24274},
		{"human", 60, 24239},
	} {
		ids := rotationIDs(t, "paladin-retribution", "paladin", c.level, specRankChoices("paladin-retribution", BarePreset, c.race))
		if !ids[c.want] {
			t.Errorf("%s bare %d: Hammer of Wrath does not cast %d", c.race, c.level, c.want)
		}
	}
}
