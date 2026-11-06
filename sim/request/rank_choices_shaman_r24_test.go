package request

import "testing"

// Round 24 (shaman lane): the newest rank was cast the level it arrived.
// Elemental's Lightning Bolt filler keeps rank 3 from 20 (not only from 26);
// an orc's Flame Shock keeps rank 1 at 19; an orc's or troll's Earth Shock
// keeps rank 3 at 25; Restoration's tank Healing Wave keeps rank 3 at 18 and
// 19 (design/reviews/2026-11-05-sim-audit-round24/shaman.md).
func TestRound24ShamanRankRows(t *testing.T) {
	const lb3, lb4 = 548, 915
	const fs1, fs2 = 8050, 8052
	const es3, es4 = 8045, 8046
	const hw3 = 547
	for _, c := range []struct {
		name, spec, race string
		level            int
		want, not        int32
	}{
		{"bolt filler at 20", "shaman-elemental", "orc", 20, lb3, 0},
		{"bolt filler at 25", "shaman-elemental", "dwarf", 25, lb3, 0},
		{"orc Flame Shock at 19", "shaman-elemental", "orc", 19, fs1, fs2},
		{"dwarf Flame Shock at 19 is the top rank", "shaman-elemental", "dwarf", 19, fs2, fs1},
		{"orc Earth Shock at 25", "shaman-enhancement", "orc", 25, es3, es4},
		{"troll Earth Shock at 25", "shaman-enhancement", "troll", 25, es3, es4},
		{"tauren Earth Shock at 25 is the top rank", "shaman-enhancement", "tauren", 25, es4, es3},
		{"orc Earth Shock at 26 is the top rank", "shaman-enhancement", "orc", 26, es4, es3},
		{"tank Healing Wave at 18", "shaman-restoration", "tauren", 18, hw3, 0},
		{"tank Healing Wave at 19", "shaman-restoration", "dwarf", 19, hw3, 0},
	} {
		ids := rotationIDs(t, c.spec, "shaman", c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || (c.not != 0 && ids[c.not]) {
			t.Errorf("%s (%s): want %d cast and %d not", c.name, c.race, c.want, c.not)
		}
	}
	// The top-rank bolt line keeps the newest rank.
	if ids := rotationIDs(t, "shaman-elemental", "shaman", 21, specRankChoices("shaman-elemental", BarePreset, "orc")); !ids[lb4] {
		t.Errorf("at 21 the top-rank Lightning Bolt line no longer casts rank 4")
	}
}
