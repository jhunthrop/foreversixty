package request

import "testing"

// Sim-accuracy round 24, second pass, gave Hammer of Wrath rank 1 (24275) a
// row at 53 for every race and at 54 for an Undead. On 2026-11-12, with
// Judgement a ranged attack, rank 2 (24274) beats rank 1 by 0.4 to 0.9% at
// 53 and by 0.65% at 54 for an Undead (seeds 1 and 50001), so the rows are
// gone: every level casts the trained top rank (rank-choices.json carries
// no Hammer of Wrath row).
func TestRetributionHammerOfWrathRankFollowsTheSweep(t *testing.T) {
	for _, c := range []struct {
		race  string
		level int
		want  int32
	}{
		{"human", 52, 24274},
		{"undead", 52, 24274},
		{"human", 53, 24274},
		{"dwarf", 53, 24274},
		{"undead", 53, 24274},
		{"human", 54, 24274},
		{"undead", 54, 24274},
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
