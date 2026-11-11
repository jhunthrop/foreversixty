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

// Sim-accuracy round 30 (open items): healer rank rows on round 29's healer orders, re-judged on
// each band's heal profile, race by race where the rows are race-keyed. Holy priest: Renew rank 1
// (139) at 16 and for every race at 28 and 29, rank 2 (6074) at 31 and 33; Discipline: Renew rank 3
// (6075) for every race at 35 and rank 4 (6076) for a Troll at 42; Restoration shaman: Healing Wave
// rank 6 (959) at 37 to 39, rank 5 (939) at 40 to 43, rank 6 from 44; Restoration druid: the party
// Rejuvenation rank 4 (2090) at 44 and rank 5 (2091) at 45 to 48, the tank's rank 9 (9840) at 56.
func TestRound30HealerRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, race string
		level                   int
		want, not               int32
	}{
		{"Human Holy Renew at 16", "priest-holy", "priest", "human", 16, 139, 6074},
		{"Undead Holy Renew at 28", "priest-holy", "priest", "undead", 28, 139, 6074},
		{"Dwarf Holy Renew at 29", "priest-holy", "priest", "dwarf", 29, 139, 6074},
		{"Gnome Holy Renew at 31", "priest-holy", "priest", "gnome", 31, 6074, 139},
		{"Human Holy Renew at 33", "priest-holy", "priest", "human", 33, 6074, 6075},
		{"Gnome Discipline Renew at 35", "priest-discipline", "priest", "gnome", 35, 6075, 6074},
		{"Troll Discipline Renew at 42", "priest-discipline", "priest", "troll", 42, 6076, 6078},
		{"Orc Restoration Healing Wave at 38", "shaman-restoration", "shaman", "orc", 38, 959, 913},
		{"Tauren Restoration Healing Wave at 41", "shaman-restoration", "shaman", "tauren", 41, 939, 8005},
		{"Dwarf Restoration Healing Wave at 44", "shaman-restoration", "shaman", "dwarf", 44, 959, 939},
		{"Night Elf Restoration Rejuvenation at 44", "druid-restoration", "druid", "night-elf", 44, 2090, 774},
		{"Tauren Restoration Rejuvenation at 47", "druid-restoration", "druid", "tauren", 47, 2091, 774},
		{"Night Elf Restoration tank Rejuvenation at 56", "druid-restoration", "druid", "night-elf", 56, 9840, 3627},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
