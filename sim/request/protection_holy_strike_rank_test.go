package request

import "testing"

// Sim-accuracy round 24: the rank sweep after Holy Strike rank 2 (678) arrives
// at 12 found rank 1 (679) ahead by 2.6 to 5.6% threat at 12 to 14 for the
// Dwarf and the Undead, and by 2.4% at 15 for the Undead (the Dwarf loses 2%
// there); from 16 rank 2 wins again for both.
func TestProtectionHolyStrikeRankOneWhileLeveling(t *testing.T) {
	for _, c := range []struct {
		race      string
		level     int
		want, not int32
	}{
		{"dwarf", 12, 679, 678},
		{"undead", 14, 679, 678},
		{"undead", 15, 679, 678},
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
