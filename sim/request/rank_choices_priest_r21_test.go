package request

import (
	"encoding/json"
	"testing"
)

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

// Round 21 (priest lane): before Heal at 16 a Holy priest's only filler is
// Lesser Heal rank 3 (learned at 10). The list had no Lesser Heal line, so
// the band-10 card and the engine cast Renew alone at 10 to 15. Lesser Heal
// holds the filler until 16 and leaves the list once Heal is trained.
func TestRound21HolyLesserHealBeforeHeal(t *testing.T) {
	const lesserHeal3, heal1 = 2053, 2054
	for _, c := range []struct {
		name      string
		level     int
		want, not int32
	}{
		{"Holy at 10", 10, lesserHeal3, heal1},
		{"Holy at 15", 15, lesserHeal3, heal1},
	} {
		ids := rotationIDs(t, "priest-holy", "priest", c.level, specRankChoices("priest-holy", BarePreset, "human"))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}

// TestRound21HolyLesserHealStopsAtSixteen: every Holy Lesser Heal line is
// scoped below 16, where Heal takes the filler.
func TestRound21HolyLesserHealStopsAtSixteen(t *testing.T) {
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/priest-holy.json", &list)
	const below = `{"cmp":{"op":"OpLt","lhs":{"currentLevel":{}},"rhs":{"const":{"val":"16"}}}}`
	lessers := 0
	for index, raw := range list.Rotation.PriorityList {
		var l struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
				} `json:"castSpell"`
				Condition json.RawMessage `json:"condition"`
			} `json:"action"`
		}
		if err := json.Unmarshal(raw, &l); err != nil {
			t.Fatal(err)
		}
		if l.Action.CastSpell == nil || l.Action.CastSpell.SpellID.SpellID != 2053 {
			continue
		}
		lessers++
		var and struct {
			And struct {
				Vals []json.RawMessage `json:"vals"`
			} `json:"and"`
		}
		if err := json.Unmarshal(l.Action.Condition, &and); err != nil || len(and.And.Vals) != 2 ||
			compactJSON(t, and.And.Vals[1]) != below {
			t.Errorf("Holy Lesser Heal line %d is not scoped below level 16", index)
		}
	}
	if lessers != 10 {
		t.Errorf("Holy casts Lesser Heal rank 3 on %d lines, want 10 (the Heal lines' targets)", lessers)
	}
}
