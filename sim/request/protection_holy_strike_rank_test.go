package request

import "testing"

// Sim-accuracy round 24 put Holy Strike rank 1 (679) on 12 to 14 and an
// Undead's 15; round 26 re-judged it on the corrected below-band gear and rank
// 2 (678) wins there by 2.2 to 2.8% threat (TestRound26ProtectionHolyStrikeTakesRankTwo).
func TestProtectionHolyStrikeRankRows(t *testing.T) {
	for _, c := range []struct {
		race      string
		level     int
		want, not int32
	}{
		{"dwarf", 15, 678, 679},
		{"undead", 16, 678, 679},
		// Round 24, second pass: rank 1 holds at 40; at 41 and 42 the top rank,
		// 5 (2495), wins by 0.5 to 1.1% for both races.
		{"dwarf", 40, 679, 2495},
		{"dwarf", 41, 2495, 679},
		{"undead", 42, 2495, 679},
	} {
		ids := rotationIDs(t, "paladin-protection", "paladin", c.level, specRankChoices("paladin-protection", BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s %d: want Holy Strike %d cast and %d not", c.race, c.level, c.want, c.not)
		}
	}
}
