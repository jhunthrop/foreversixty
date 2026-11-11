package request

import "testing"

// Sim-accuracy round 30 (open items): rank rows on the levels round 29's leveling orders moved,
// re-judged on the new talents level by level (design/reviews/2026-11-11-sim-accuracy/round30-open.md).
// Arcane: with Missile Barrage, Arcane Impact and Arcane Meditation at 21 to 28, Frostbolt rank 1
// (116) beats rank 4 (7322) at 24 to 27. Fire: with Precision at 41 to 43, Scorch rank 3 (8445) wins
// at 50 and 51 and rank 4 (8446) at 52 to 55, and a Human casts Pyroblast rank 4 (12523) at 41.
// Balance: Wrath rank 7 (8905) beats rank 6 (6780) from 48.
func TestRound30DamageRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, race string
		level                   int
		want, not               int32
	}{
		{"Troll Arcane Frostbolt at 24", "mage-arcane", "mage", "troll", 24, 116, 7322},
		{"Human Arcane Frostbolt at 27", "mage-arcane", "mage", "human", 27, 116, 7322},
		{"Gnome Arcane Frostbolt at 28", "mage-arcane", "mage", "gnome", 28, 7322, 116},
		{"Undead Fire Scorch at 50", "mage-fire", "mage", "undead", 50, 8445, 10205},
		{"Human Fire Scorch at 51", "mage-fire", "mage", "human", 51, 8445, 10205},
		{"Gnome Fire Scorch at 52", "mage-fire", "mage", "gnome", 52, 8446, 10206},
		{"Troll Fire Scorch at 55", "mage-fire", "mage", "troll", 55, 8446, 10206},
		{"Orc Fire Scorch at 56", "mage-fire", "mage", "orc", 56, 10206, 8446},
		{"Human Fire Pyroblast at 41", "mage-fire", "mage", "human", 41, 12523, 12522},
		{"Undead Fire Pyroblast at 41", "mage-fire", "mage", "undead", 41, 12522, 12523},
		{"Night Elf Balance Wrath at 47", "druid-balance", "druid", "night-elf", 47, 6780, 8905},
		{"Tauren Balance Wrath at 48", "druid-balance", "druid", "tauren", 48, 8905, 6780},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
