package request

import "testing"

// Sim-accuracy round 24: the rank sweep after each new rank found Hammer of
// Wrath rank 1 (24275) ahead of rank 2 (24274) beyond error at 52 at both seeds
// and factions, and for Alliance at 56 with Horde level; every other level
// keeps the trained top rank (rank-choices.json notes carry the numbers).
func TestRetributionHammerOfWrathRankFollowsTheSweep(t *testing.T) {
	for _, c := range []struct {
		level int
		want  int32
	}{
		{52, 24275},
		{53, 24274},
		{56, 24275},
		{57, 24274},
		{60, 24239},
	} {
		ids := rotationIDs(t, "paladin-retribution", "paladin", c.level, specRankChoices("paladin-retribution", BarePreset, "human"))
		if !ids[c.want] {
			t.Errorf("bare %d: Hammer of Wrath does not cast %d", c.level, c.want)
		}
	}
}
