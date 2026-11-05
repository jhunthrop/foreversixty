package request

import (
	"encoding/json"
	"testing"
)

// Round 22 (priest lane, guides round 24 sim 3): before Heal at 16 a
// Discipline priest's tank had only Power Word: Shield and Renew, while the
// party already had Lesser Heal rank 3. Two tank lines now cast Lesser Heal
// rank 3 below 16, mirroring the tank's two Heal lines (paced under 80%, and
// under 50% whatever the mana).
func TestRound22DisciplineTankLesserHealBeforeHeal(t *testing.T) {
	var list struct {
		Rotation struct {
			PriorityList []json.RawMessage `json:"priorityList"`
		} `json:"rotation"`
	}
	readJSON(t, "../../data/curated/apl/priest-discipline.json", &list)
	const below = `{"cmp":{"op":"OpLt","lhs":{"currentLevel":{}},"rhs":{"const":{"val":"16"}}}}`
	tankLines := 0
	for index, raw := range list.Rotation.PriorityList {
		var l struct {
			Action struct {
				CastSpell *struct {
					SpellID struct {
						SpellID int `json:"spellId"`
					} `json:"spellId"`
					Target struct {
						Index int `json:"index"`
					} `json:"target"`
				} `json:"castSpell"`
				Condition json.RawMessage `json:"condition"`
			} `json:"action"`
		}
		if err := json.Unmarshal(raw, &l); err != nil {
			t.Fatal(err)
		}
		c := l.Action.CastSpell
		if c == nil || c.SpellID.SpellID != 2053 || c.Target.Index != 5 {
			continue
		}
		tankLines++
		var and struct {
			And struct {
				Vals []json.RawMessage `json:"vals"`
			} `json:"and"`
		}
		if err := json.Unmarshal(l.Action.Condition, &and); err != nil || len(and.And.Vals) != 2 ||
			compactJSON(t, and.And.Vals[1]) != below {
			t.Errorf("Discipline tank Lesser Heal line %d is not scoped below level 16", index)
		}
	}
	if tankLines != 2 {
		t.Errorf("Discipline casts Lesser Heal on the tank on %d lines, want 2", tankLines)
	}
	ids := rotationIDs(t, "priest-discipline", "priest", 12, specRankChoices("priest-discipline", BarePreset, "human"))
	if !ids[2053] || ids[2054] {
		t.Error("Discipline at 12: want Lesser Heal rank 3 cast and Heal rank 1 not")
	}
}

// Round 22 (priest lane, guides round 24 sim 1): Heal rank 1 was never tried
// against the later ranks, although Forever has no sub-20 cut and rank 1
// keeps its whole coefficient to 33. Level by level it heals more for the
// fight on both healing specs through the 20s and early 30s, race by race on
// Holy (design/reviews/2026-11-05-sim-audit-round22/priest.md).
func TestRound22PriestHealRankRows(t *testing.T) {
	const r1, r2, r3, r4 = 2054, 2055, 6063, 6064
	for _, c := range []struct {
		name, spec, race string
		level            int
		want, not        int32
	}{
		{"Holy Human at 22", "priest-holy", "human", 22, r1, r2},
		{"Holy Gnome at 26", "priest-holy", "gnome", 26, r1, r2},
		{"Holy Troll at 27", "priest-holy", "troll", 27, r1, r2},
		{"Holy Gnome at 27 keeps rank 2", "priest-holy", "gnome", 27, r2, r1},
		{"Holy Undead at 28", "priest-holy", "undead", 28, r1, r2},
		{"Holy Human at 28 keeps rank 2", "priest-holy", "human", 28, r2, r1},
		{"Holy Dwarf at 28 is the top rank", "priest-holy", "dwarf", 28, r3, r1},
		{"Holy Night Elf at 29", "priest-holy", "night-elf", 29, r1, r3},
		{"Holy Gnome at 31", "priest-holy", "gnome", 31, r1, r3},
		{"Holy Undead at 31 is the top rank", "priest-holy", "undead", 31, r3, r1},
		{"Holy Dwarf at 32", "priest-holy", "dwarf", 32, r2, r3},
		{"Holy Troll at 32", "priest-holy", "troll", 32, r1, r3},
		{"Holy Human at 32 is the top rank", "priest-holy", "human", 32, r3, r1},
		{"Holy Undead at 33", "priest-holy", "undead", 33, r1, r3},
		{"Holy Troll at 34", "priest-holy", "troll", 34, r1, r3},
		{"Holy Troll at 35", "priest-holy", "troll", 35, r3, r1},
		{"Holy Human at 34", "priest-holy", "human", 34, r3, r1},
		{"Holy Undead at 38", "priest-holy", "undead", 38, r3, r4},
		{"Holy Troll at 38", "priest-holy", "troll", 38, r2, r3},
		{"Holy Undead at 39", "priest-holy", "undead", 39, r2, r4},
		{"Holy Troll at 39", "priest-holy", "troll", 39, r2, r4},
		{"Holy Human at 39 is the top rank", "priest-holy", "human", 39, r4, r2},
		{"Discipline Gnome at 22", "priest-discipline", "gnome", 22, r1, r2},
		{"Discipline Undead at 30", "priest-discipline", "undead", 30, r1, r2},
		{"Discipline Human at 35", "priest-discipline", "human", 35, r1, r3},
		{"Discipline Dwarf at 36", "priest-discipline", "dwarf", 36, r3, r1},
	} {
		ids := rotationIDs(t, c.spec, "priest", c.level, specRankChoices(c.spec, BarePreset, c.race))
		if !ids[c.want] || ids[c.not] {
			t.Errorf("%s: want %d cast and %d not", c.name, c.want, c.not)
		}
	}
}
