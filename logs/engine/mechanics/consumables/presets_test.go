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
