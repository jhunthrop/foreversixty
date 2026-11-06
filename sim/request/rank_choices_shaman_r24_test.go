package request

import "testing"

// Round 24 (shaman lane): the newest rank was cast the level it arrived, and
// the second pass re-judged every row on the talents a player has at that
// level. Elemental's Lightning Bolt filler keeps rank 3 from 20 to 33 and,
// at 39, rank 4 for every race but an orc; Flame Shock rank 1 at 18 and 19
// (round 26: a dwarf only); Earth Shock keeps rank 3 at 25 (an orc casts
// rank 4) and at 26; Restoration's tank Healing Wave keeps rank 3 at 18 and
// 19, rank 1 at 20 and 21, rank 1 at 22 for a tauren and rank 2 for a dwarf
// (design/reviews/2026-11-05-sim-audit-round24/shaman.md).
func TestRound24ShamanRankRows(t *testing.T) {
	const lb3, lb4, lb5, lb7 = 548, 915, 943, 10391
	const fs1, fs2 = 8050, 8052
	const es3, es4 = 8045, 8046
	const hw1, hw2, hw3 = 331, 332, 547
	for _, c := range []struct {
		name, spec, race string
		level            int
		want, not        int32
	}{
		{"bolt filler at 20", "shaman-elemental", "orc", 20, lb3, 0},
		{"bolt filler at 25", "shaman-elemental", "dwarf", 25, lb3, 0},
		{"bolt filler at 28", "shaman-elemental", "orc", 28, lb3, 0},
		{"bolt filler at 29", "shaman-elemental", "dwarf", 29, lb3, 0},
		{"troll bolt filler at 39", "shaman-elemental", "troll", 39, lb4, lb5},
		{"dwarf bolt filler at 39", "shaman-elemental", "dwarf", 39, lb4, lb5},
		{"orc bolt filler at 39", "shaman-elemental", "orc", 39, lb5, lb4},
		// Round 26 re-judged Flame Shock at 18 and 19 on the corrected below-band
		// gear: a dwarf keeps rank 1 at both; a troll, an orc, a tauren and a
		// Windshaper Skyborne take the new rank 2.
		{"troll Flame Shock at 18 is the top rank", "shaman-elemental", "troll", 18, fs2, fs1},
		{"orc Flame Shock at 18 is the top rank", "shaman-elemental", "orc", 18, fs2, fs1},
		{"dwarf Flame Shock at 18", "shaman-elemental", "dwarf", 18, fs1, fs2},
		{"orc Flame Shock at 19 is the top rank", "shaman-elemental", "orc", 19, fs2, fs1},
		{"troll Flame Shock at 19 is the top rank", "shaman-elemental", "troll", 19, fs2, fs1},
		{"dwarf Flame Shock at 19", "shaman-elemental", "dwarf", 19, fs1, fs2},
		{"orc Earth Shock at 25 is the top rank", "shaman-enhancement", "orc", 25, es4, es3},
		{"troll Earth Shock at 25", "shaman-enhancement", "troll", 25, es3, es4},
		{"tauren Earth Shock at 25", "shaman-enhancement", "tauren", 25, es3, es4},
		{"dwarf Earth Shock at 25", "shaman-enhancement", "dwarf", 25, es3, es4},
		{"orc Earth Shock at 26", "shaman-enhancement", "orc", 26, es3, es4},
		{"Earth Shock at 27 is the top rank", "shaman-enhancement", "dwarf", 27, es4, es3},
		{"tank Healing Wave at 18", "shaman-restoration", "tauren", 18, hw3, 0},
		{"tank Healing Wave at 19", "shaman-restoration", "dwarf", 19, hw3, 0},
		{"tank Healing Wave at 20", "shaman-restoration", "dwarf", 20, hw1, hw2},
		{"tank Healing Wave at 21", "shaman-restoration", "tauren", 21, hw1, hw2},
		{"tauren tank Healing Wave at 22", "shaman-restoration", "tauren", 22, hw1, hw2},
		{"dwarf tank Healing Wave at 22", "shaman-restoration", "dwarf", 22, hw2, hw1},
		{"dwarf tank Healing Wave at 25", "shaman-restoration", "dwarf", 25, hw3, hw2},
	} {
		ids := rotationIDs(t, c.spec, "shaman", c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || (c.not != 0 && ids[c.not]) {
			t.Errorf("%s (%s): want %d cast and %d not", c.name, c.race, c.want, c.not)
		}
	}
	// The top-rank bolt lines keep the newest rank.
	if ids := rotationIDs(t, "shaman-elemental", "shaman", 21, specRankChoices("shaman-elemental", BarePreset, "orc")); !ids[lb4] {
		t.Errorf("at 21 the top-rank Lightning Bolt line no longer casts rank 4")
	}
	if ids := rotationIDs(t, "shaman-elemental", "shaman", 39, specRankChoices("shaman-elemental", BarePreset, "troll")); !ids[lb7] {
		t.Errorf("at 39 the top-rank Lightning Bolt line no longer casts rank 7")
	}
}
