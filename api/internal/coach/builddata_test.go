package coach

import (
	"slices"
	"testing"
)

// The fixtures' hand-written spell constants are the newest build's spellconst values:
// Bloodthirst's 6 s cooldown is its category cooldown, and its cost is 300 tenths of rage.
func TestFixtureSpellConstMatchesData(t *testing.T) {
	for class, fixture := range map[string]map[int64]SpellConst{"warrior": warriorSpells, "priest": priestSpells} {
		data, err := ParseSpellConst(buildFile(t, "spellconst", class+".json"))
		if err != nil {
			t.Fatal(err)
		}
		for id, want := range fixture {
			if got, ok := data[id]; !ok || got != want {
				t.Errorf("%s %d: fixture %+v, data %+v", class, id, want, got)
			}
		}
	}
	bt := warriorSpells[23881]
	if bt.CooldownMS != 6000 || bt.Cost != 300 || !bt.CostKnown {
		t.Errorf("Bloodthirst fixture %+v: want the 6 s category cooldown and cost 300", bt)
	}
}

func TestFixtureTrainablesMatchData(t *testing.T) {
	for class, fixture := range map[string]map[string][]TrainableRank{"warrior": warriorTrainable, "priest": priestTrainable} {
		data, err := ParseTrainables(buildFile(t, "trainables", class+".json"))
		if err != nil {
			t.Fatal(err)
		}
		for family, want := range fixture {
			if got := data[family]; !slices.Equal(got, want) {
				t.Errorf("%s %s: fixture %v, data %v", class, family, want, got)
			}
		}
	}
}

func TestParseSpellConst_CostKnown(t *testing.T) {
	got, err := ParseSpellConst([]byte(`{"spells": {
		"1": {"name": "Flat", "cost": 50, "cost_pct": 0},
		"2": {"name": "Share", "cost": 0, "cost_pct": 12.5},
		"3": {"name": "Free", "cost": 0, "cost_pct": 0},
		"4": {"name": "Unstated"},
		"5": {"name": "Category", "cooldown_ms": 1000, "category_cooldown_ms": 6000}
	}}`))
	if err != nil {
		t.Fatal(err)
	}
	want := map[int64]bool{1: true, 2: false, 3: true, 4: false}
	for id, known := range want {
		if got[id].CostKnown != known {
			t.Errorf("spell %d (%s): CostKnown = %v, want %v", id, got[id].Name, got[id].CostKnown, known)
		}
	}
	if got[5].CooldownMS != 6000 {
		t.Errorf("the cooldown is the larger of the two: %d", got[5].CooldownMS)
	}
	if _, err := ParseSpellConst([]byte(`{"spells": {"x": {}}}`)); err == nil {
		t.Error("a non-numeric spell key must be refused")
	}
}

// The fixture inputs carry what the F tasks read: the role's row at level 60, the
// shipped rotation, the role's consumables and the Onyxia table.
func TestFixtureInputs(t *testing.T) {
	cases := []struct {
		name             string
		in               Input
		guid, spec, role string
	}{
		{"fury", furyInput(t), furyGUID, "Fury", RoleDPS},
		{"holy", holyInput(t), priestGUID, "Holy", RoleHealer},
		{"prot", protInput(t), furyGUID, "Protection", RoleTank},
	}
	for _, c := range cases {
		in := c.in
		if in.Row.GUID != c.guid || in.Row.Spec != c.spec || in.Row.Role != c.role || in.Row.Level != fixtureLevel {
			t.Errorf("%s row %+v", c.name, in.Row)
		}
		if in.Fight.EncounterID != onyxiaEncounterID || !in.Fight.Kill || in.Size != 40 || in.BossName != onyxiaName {
			t.Errorf("%s fight: encounter %d kill %v size %d boss %q", c.name, in.Fight.EncounterID, in.Fight.Kill, in.Size, in.BossName)
		}
		if in.Tables.Mechanics == nil || len(in.Tables.Consumables.Flask) == 0 || len(in.Tables.RaidBuffs) == 0 {
			t.Errorf("%s tables %+v", c.name, in.Tables)
		}
		if len(in.Fight.Combatants) != len(in.Fight.Roster) {
			t.Errorf("%s: %d combatant rows for %d players", c.name, len(in.Fight.Combatants), len(in.Fight.Roster))
		}
	}
	if fury := furyInput(t); !fury.Tables.RotationOK || !fury.Tables.Rotation.Lists(23894) {
		t.Error("the fury fixture carries the shipped rotation, which lists Bloodthirst")
	}
}

// The mutators return a new summary and never change their argument.
func TestFixtureMutatorsDoNotChangeTheirArgument(t *testing.T) {
	s := goldenFight(t)
	casts, auras, deaths, roster := len(s.Casts), len(s.Auras), len(s.Deaths), s.Roster[0]
	_ = withCasts(s, furyGUID, 23881, "Bloodthirst", 0, 6000)
	_ = withoutCasts(s, furyGUID)
	_ = withAura(s, furyGUID, 6673, "Battle Shout", nil, span(0, 1000))
	_ = withDeath(s, s.Roster[0].GUID, 1000, 0, "Melee")
	_ = withRole(s, s.Roster[0].GUID, "Mage", "Fire", RoleDPS)
	_ = withSeeds(s, furyGUID, seededAt(flask("Flask of the Titans"), food("Grilled Squid")))
	if len(s.Casts) != casts || len(s.Auras) != auras || len(s.Deaths) != deaths || s.Roster[0] != roster {
		t.Fatal("a mutator changed its argument")
	}
	got := withSeeds(withDuration(s, 60_000), furyGUID, seededAt(flask("Flask of the Titans")))
	if tr := got.Auras[len(got.Auras)-1]; tr.TargetGUID != furyGUID || tr.Segments[0] != span(0, 60_000) {
		t.Fatalf("a seed is one segment from the pull to the end: %+v", tr)
	}
	if seq := seqEvery(0, 6000, 3); !slices.Equal(seq, []int64{0, 6000, 12000}) {
		t.Fatalf("seqEvery = %v", seq)
	}
	if r := repeat(500, 3); !slices.Equal(r, []int64{500, 500, 500}) {
		t.Fatalf("repeat = %v", r)
	}
	if p := poolCPM("Bloodthirst", 9.1, 19); p.Pool != nil || p.PoolCount != 19 {
		t.Fatalf("a pool under PoolMinPlayers is no pool: %+v", p)
	}
}
