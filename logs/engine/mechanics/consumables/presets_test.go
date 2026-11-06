package consumables

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestPresetExpects_FeralCatAt60CarriesNoFlask(t *testing.T) {
	got, ok := PresetExpects("druid-feral", 60)
	if !ok || got[CategoryFlask] || !got[CategoryFood] || !got[CategoryCombatPotion] || !got[CategoryWeaponEnchant] {
		t.Fatalf("druid-feral at 60: %v, %v", got, ok)
	}
}

func TestPresetExpects_NoFlaskBelow50(t *testing.T) {
	got, ok := PresetExpects("druid-balance", 30)
	if !ok || got[CategoryFlask] || got[CategoryFood] {
		t.Fatalf("druid-balance at 30: %v, %v", got, ok)
	}
	if got, _ := PresetExpects("druid-balance", 60); !got[CategoryFlask] {
		t.Fatalf("druid-balance at 60 carries Flask of Supreme Power: %v", got)
	}
}

func TestPresetExpects_PicksTheHighestBandAtOrBelowTheLevel(t *testing.T) {
	// druid-balance's bands start at 1, 35, 45, ...: Nightfin Soup arrives at 35.
	if got, _ := PresetExpects("druid-balance", 34); got[CategoryFood] {
		t.Fatalf("34 is the level 1 band: %v", got)
	}
	if got, _ := PresetExpects("druid-balance", 35); !got[CategoryFood] {
		t.Fatalf("35 is the level 35 band: %v", got)
	}
}

func TestPresetExpects_UnknownSpecOrLevelIsNotOK(t *testing.T) {
	for _, c := range []struct {
		spec  string
		level int64
	}{{"druid-feral", 0}, {"", 60}, {"monk-windwalker", 60}} {
		if got, ok := PresetExpects(c.spec, c.level); ok {
			t.Errorf("%q at %d: %v, want no preset", c.spec, c.level, got)
		}
	}
}

func TestMissedAtPull_ACategoryNotExpectedIsNeverMissing(t *testing.T) {
	rc, _ := Load().For("dps")
	fight := summary.Summary{Combatants: []summary.CombatantRow{{GUID: "p", Gear: make([]event.Item, 1)}}}
	pull := AtPull(fight, "p", rc)
	if !rc.MissedAtPull(pull, CategoryFlask) {
		t.Fatal("role behaviour: the dps flask list is complete, an empty sheet misses it")
	}
	narrowed := rc.Expect(map[string]bool{CategoryFood: true})
	if narrowed.MissedAtPull(pull, CategoryFlask) || narrowed.Expects(CategoryFlask) {
		t.Fatal("a flask the preset does not carry is never missing")
	}
	if !narrowed.Expects(CategoryFood) || !rc.Expects(CategoryFlask) {
		t.Fatal("Expects: the narrowed food, and every category without a narrowing")
	}
}

func TestForPlayer_NarrowsTheRoleToTheSpecPresetAtLevel(t *testing.T) {
	rc, ok := Load().ForPlayer("dps", "druid-feral", 60)
	if !ok {
		t.Fatal("the dps role has a catalogue")
	}
	if rc.Expects(CategoryFlask) || rc.IsComplete(CategoryFlask) {
		t.Fatal("a feral cat's preset carries no flask at 60: never expected, never proven absent")
	}
	if !rc.Expects(CategoryCombatPotion) || !rc.Expects(CategoryFood) {
		t.Fatal("the feral cat's preset carries a potion and food at 60")
	}
	if !rc.Expects(CategoryWorldBuffs) || !rc.IsComplete(CategoryWorldBuffs) {
		t.Fatal("no preset carries world buffs, so the narrowing leaves them to the role")
	}
	fire, _ := Load().ForPlayer("dps", "mage-fire", 60)
	if !fire.Expects(CategoryFlask) || !fire.IsComplete(CategoryFlask) {
		t.Fatal("a Fire mage's preset carries a flask at 60")
	}
}

func TestForPlayer_NoLevelOrNoPresetKeepsTheRole(t *testing.T) {
	role, _ := Load().For("dps")
	for _, c := range []struct {
		spec  string
		level int64
	}{{"druid-feral", 0}, {"monk-windwalker", 60}, {"", 60}} {
		rc, ok := Load().ForPlayer("dps", c.spec, c.level)
		if !ok || !rc.Expects(CategoryFlask) || rc.IsComplete(CategoryFlask) != role.IsComplete(CategoryFlask) {
			t.Errorf("%q at %d: want the role's own expectation", c.spec, c.level)
		}
	}
	if _, ok := Load().ForPlayer("nobody", "druid-feral", 60); ok {
		t.Error("an unknown role has no catalogue")
	}
}
