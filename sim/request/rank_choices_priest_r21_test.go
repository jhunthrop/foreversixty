package request

import "testing"

// Round 21 (priest lane): the filler Heal before Greater Heal. A new Heal rank
// costs more mana than its extra size returns on band gear in its first levels:
// rank 3 (learned at 34 as rank 4 arrives) keeps the whole healing bonus and
// heals more for the fight at 34 to 37 on Holy and 34 to 39 on Discipline,
// and rank 2 holds at 28 for a Holy Human or Undead, at 29 for a Holy Undead
// and at 30 to 33 for a Discipline Undead. Every other level and race casts the
// top rank (design/reviews/2026-11-05-sim-audit-round21/priest.md, P-HEAL).
func TestRound21PriestHealRankRows(t *testing.T) {
	const r2, r3, r4 = 2055, 6063, 6064
	for _, c := range []struct {
		name, spec, race string
		level            int
		want, not        int32
	}{
		{"Holy Human at 28", "priest-holy", "human", 28, r2, r3},
		{"Holy Undead at 29", "priest-holy", "undead", 29, r2, r3},
		{"Holy Human at 29 is the top rank", "priest-holy", "human", 29, r3, r2},
		{"Holy Gnome at 28 is the top rank", "priest-holy", "gnome", 28, r3, r2},
		{"Holy Human at 30 is the top rank", "priest-holy", "human", 30, r3, r2},
		{"Holy Human at 34", "priest-holy", "human", 34, r3, r4},
		{"Holy Dwarf at 37", "priest-holy", "dwarf", 37, r3, r4},
		{"Holy Troll at 38", "priest-holy", "troll", 38, r3, r4},
		{"Holy Gnome at 38 is the top rank", "priest-holy", "gnome", 38, r4, r3},
		{"Holy Undead at 39", "priest-holy", "undead", 39, r3, r4},
		{"Holy Troll at 39 is the top rank", "priest-holy", "troll", 39, r4, r3},
		{"Discipline Undead at 30", "priest-discipline", "undead", 30, r2, r3},
		{"Discipline Undead at 33", "priest-discipline", "undead", 33, r2, r3},
		{"Discipline Gnome at 30 is the top rank", "priest-discipline", "gnome", 30, r3, r2},
		{"Discipline Gnome at 28 is the top rank", "priest-discipline", "gnome", 28, r3, r2},
		{"Discipline Gnome at 34", "priest-discipline", "gnome", 34, r3, r4},
		{"Discipline Troll at 39", "priest-discipline", "troll", 39, r3, r4},
	} {
		ids := rotationIDs(t, c.spec, "priest", c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
