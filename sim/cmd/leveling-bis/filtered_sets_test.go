package main

import (
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"testing"
)

func TestFilterSetAllowsByKindAndStanding(t *testing.T) {
	rep := fixtureItem(1, "Honored Helm", fixtureHead, 5, itemSource{Kind: "rep", Label: "Stormwind", Standing: "honored"})
	friendly := fixtureItem(2, "Friendly Helm", fixtureHead, 4, itemSource{Kind: "rep", Label: "Stormwind", Standing: "friendly"})
	both := fixtureItem(3, "Both Helm", fixtureHead, 4, itemSource{Kind: "vendor", Label: "Vendor", Standing: "honored"}, src("dungeon", "Deadmines"))
	leveling := publishedFilterSets[2]
	if leveling.Key != "leveling_default" || leveling.allows(rep) || !leveling.allows(friendly) {
		t.Errorf("leveling_default: honored allowed=%v, friendly allowed=%v; want false, true", leveling.allows(rep), leveling.allows(friendly))
	}
	if !leveling.allows(both) {
		t.Error("an honored vendor item that a dungeon also drops was dropped; its dungeon source needs no standing")
	}
	if dungeon := publishedFilterSets[1]; dungeon.allows(rep) || !dungeon.allows(both) {
		t.Error("dungeon filter set let a rep item through or dropped a dungeon drop")
	}
}

// Every distinct filtered set is measured once under the guard harness;
// keys whose set equals another reuse its measurement.
func TestMeasureFilteredSetsMeasuresEachDistinctSetOnce(t *testing.T) {
	bySlot := candidatesBySlot(sourceFilterFixturePool(false))
	picks := pick("warrior-arms", bySlot)
	engine := &fakeEngine{DPSFunc: func(req api.SimRequest) (float64, error) {
		dps := 100.0
		for _, g := range req.Character.Gear {
			dps += float64(g.ItemID % 100)
		}
		return dps, nil
	}, DefaultStdErr: 1.234}
	in := filteredSetInputs{Spec: specInfo{Spec: "warrior-arms", ClassSlug: "warrior"}, Race: "human", Band: 30, Picks: picks, PickBySlot: bySlot, SetDPS: 321.04, SetDPSError: 0.66}
	got, notes := measureFilteredSets(engine, in)
	if len(notes) != 0 {
		t.Fatalf("notes = %v", notes)
	}
	for _, key := range []string{"all", "dungeon", "leveling_default"} {
		if _, ok := got[key]; !ok {
			t.Errorf("filtered_set_dps has no %q", key)
		}
	}
	if got["all"].ChangedSlots != 0 || got["dungeon"].ChangedSlots == 0 {
		t.Errorf("changed slots: all=%d dungeon=%d; want 0 and some", got["all"].ChangedSlots, got["dungeon"].ChangedSlots)
	}
	if got["all"].DPS != 321 || got["all"].Error != 0.7 {
		t.Errorf("all = %+v, want set_dps itself (321.0 +/- 0.7)", got["all"])
	}
	if got["dungeon"].Error != 1.2 {
		t.Errorf("dungeon error = %v, want 1.2 (rounded to 0.1)", got["dungeon"].Error)
	}
	// "all" is the published set (no sim); dungeon and the leveling
	// default (only the Revered trinket dropped) are two distinct sets.
	if len(engine.Calls) != 2 {
		t.Errorf("%d sims for 2 new distinct sets: %v", len(engine.Calls), engine.Calls)
	}
	in.Band = 60
	at60, _ := measureFilteredSets(engine, in)
	if _, ok := at60["leveling_default"]; ok {
		t.Error("leveling_default measured at band 60")
	}

	dungeonSet := repickFiltered("warrior-arms", picks, bySlot, publishedFilterSets[1].allows)
	fail := &fakeEngine{FailGear: gearKey(buildGear(dungeonSet)), DefaultDPS: 50}
	partial, notes := measureFilteredSets(fail, filteredSetInputs{Spec: in.Spec, Race: "human", Band: 30, Picks: picks, PickBySlot: bySlot})
	if _, ok := partial["dungeon"]; ok || len(notes) == 0 || !strings.Contains(notes[0], "dungeon") {
		t.Errorf("a failed sim published %v with notes %v; want the key left out and a note", partial, notes)
	}
}
