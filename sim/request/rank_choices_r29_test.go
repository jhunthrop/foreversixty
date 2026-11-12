package request

import "testing"

// Sim-accuracy round 29 (walks-heal): rank rows the new healer leveling orders moved
// (design/reviews/2026-11-07-sim-audit-round29/walksheal.md). Holy paladin: with Conviction
// and Divine Intellect in place of Holy Shock and Holy Power at 40 to 42, Flash of Light
// rank 2 (19939) on the tank beats rank 1 (19750) by 1.6 to 14%, so the tank's rank 2 row
// starts at 40. Restoration shaman: with Tidal Focus ahead of Improved Healing Wave, rank 1
// (331) heals 11% more than rank 3 (547) at 18 and 19, rank 2 (332) beats rank 1 at 20 to
// 22, and a dwarf or Windshaper takes rank 3 at 22.
func TestRound29HealerRankRows(t *testing.T) {
	for _, c := range []struct {
		name, spec, class, race string
		level                   int
		want, not               int32
	}{
		// Sim-accuracy 2026-11-12 (Holy paladin lane): rank 1 to 42 on the
		// engine without Illumination's casting regen; rank 2 from 43.
		{"Human Holy tank Flash of Light at 40", "paladin-holy", "paladin", "human", 40, 19750, 19939},
		{"Undead Holy tank Flash of Light at 42", "paladin-holy", "paladin", "undead", 42, 19750, 19939},
		{"Human Holy Flash of Light at 43", "paladin-holy", "paladin", "human", 43, 19939, 19750},
		{"Tauren Restoration Healing Wave at 18", "shaman-restoration", "shaman", "tauren", 18, 331, 547},
		{"Dwarf Restoration Healing Wave at 19", "shaman-restoration", "shaman", "dwarf", 19, 331, 547},
		{"Tauren Restoration Healing Wave at 20", "shaman-restoration", "shaman", "tauren", 20, 332, 331},
		{"Dwarf Restoration Healing Wave at 21", "shaman-restoration", "shaman", "dwarf", 21, 332, 331},
		{"Troll Restoration Healing Wave at 22", "shaman-restoration", "shaman", "troll", 22, 332, 331},
		{"Dwarf Restoration Healing Wave at 22", "shaman-restoration", "shaman", "dwarf", 22, 547, 332},
	} {
		ids := rotationIDs(t, c.spec, c.class, c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
