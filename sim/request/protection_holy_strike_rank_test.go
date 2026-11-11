package request

import "testing"

// Sim-accuracy round 24 put Holy Strike rank 1 (679) on 12 to 14 and an
// Undead's 15; round 26 re-judged it on the corrected below-band gear and rank
// 2 (678) wins there by 2.2 to 2.8% threat (TestRound26ProtectionHolyStrikeTakesRankTwo).
// The 2026-11-11 paladins lane fixed Holy Strike's formula (the flat Holy bonus
// is no longer cut by the weapon percent), and rank 1 lost its last row: at 40
// the top rank, 5 (2495), beats it by 4.2 to 4.7% threat for a Dwarf and an
// Undead at seeds 1 and 50001 (design/reviews/2026-11-11-sim-accuracy/paladins.md).
func TestProtectionHolyStrikeRankRows(t *testing.T) {
	for _, c := range []struct {
		race      string
		level     int
		want, not int32
	}{
		{"dwarf", 15, 678, 679},
		{"undead", 16, 678, 679},
		// The top rank at 40 to 42 (2026-11-11), as at every other level.
		{"dwarf", 40, 2495, 679},
		{"undead", 40, 2495, 679},
		{"dwarf", 41, 2495, 679},
		{"undead", 42, 2495, 679},
	} {
		ids := rotationIDs(t, "paladin-protection", "paladin", c.level, specRankChoices("paladin-protection", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s %d: want Holy Strike %d cast and %d not", c.race, c.level, c.want, c.not)
		}
	}
}
