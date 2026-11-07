package request

import "testing"

// Sim-accuracy round 28 (casters): the leveling orders moved Destruction, Demonology and
// Frost talents ahead, so the rank rows of those levels were judged again on the new talents
// (spec-breakdown -level N -rank, bare, band gear, 10,000 iterations, seeds 1 and 50001,
// every race; a gain above 0.15% at both seeds).
func TestRound28CasterRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, preset, race string
		level                           int
		want, not                       int32
	}{
		// Destruction's Shadow Bolt rank 4 (1088) at 28 for every race, not only a Troll.
		{"Orc Destruction Shadow Bolt at 28", "warlock-destruction", "warlock", BarePreset, "orc", 28, 1088, 1106},
		{"Human Destruction Shadow Bolt at 28", "warlock-destruction", "warlock", BarePreset, "human", 28, 1088, 1106},
		// Destruction's Immolate rank 5 (11665) at 40: with Shadow and Flame 4 the top rank wins.
		{"Orc Destruction Immolate at 40", "warlock-destruction", "warlock", BarePreset, "orc", 40, 11665, 2941},
		{"Human Destruction Immolate at 40", "warlock-destruction", "warlock", BarePreset, "human", 40, 11665, 2941},
		// Demonology's Shadow Bolt rank 3 (705) at 22 and 23 for a Gnome; rank 4 (1088) for the rest.
		{"Gnome Demonology Shadow Bolt at 22", "warlock-demonology", "warlock", BarePreset, "gnome", 22, 705, 1088},
		{"Gnome Demonology Shadow Bolt at 23", "warlock-demonology", "warlock", BarePreset, "gnome", 23, 705, 1088},
		{"Undead Demonology Shadow Bolt at 22", "warlock-demonology", "warlock", BarePreset, "undead", 22, 1088, 705},
		// Frost's Frostbolt with Frost Channeling from 20 and Fingers of Frost from 30:
		// rank 4 (7322) at 29 and 30 for a Human, Troll or Skyborne, at 31 and 32 for every
		// race but a Gnome, at 33 for a Human, Troll or Skyborne; rank 5 (8406) otherwise
		// until rank 6 (8407) wins from 34.
		{"Skyborne Frost Frostbolt at 29", "mage-frost", "mage", BarePreset, "high-order-skyborne", 29, 7322, 8406},
		{"Troll Frost Frostbolt at 30", "mage-frost", "mage", BarePreset, "troll", 30, 7322, 8406},
		{"Orc Frost Frostbolt at 30", "mage-frost", "mage", BarePreset, "orc", 30, 8406, 7322},
		{"Undead Frost Frostbolt at 31", "mage-frost", "mage", BarePreset, "undead", 31, 7322, 8406},
		{"Gnome Frost Frostbolt at 31", "mage-frost", "mage", BarePreset, "gnome", 31, 8406, 7322},
		{"Orc Frost Frostbolt at 32", "mage-frost", "mage", BarePreset, "orc", 32, 7322, 8406},
		{"Gnome Frost Frostbolt at 32", "mage-frost", "mage", BarePreset, "gnome", 32, 8406, 8407},
		{"Human Frost Frostbolt at 33", "mage-frost", "mage", BarePreset, "human", 33, 7322, 8406},
		{"Undead Frost Frostbolt at 33", "mage-frost", "mage", BarePreset, "undead", 33, 8406, 7322},
		{"Troll Frost Frostbolt at 34", "mage-frost", "mage", BarePreset, "troll", 34, 8407, 7322},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, c.preset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
