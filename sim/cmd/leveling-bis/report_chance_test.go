package main

import (
	"encoding/json"
	"testing"
)

// Task L1.2 (addon v2, contract 1): every published pick carries its
// drop chance as a fraction and its next two candidates as runners_up,
// so the addon's tooltip can say "The Deadmines · Defias Overseer · 14%"
// and "Also: X, Y" without a second data source.

func chanceRow(t *testing.T, picks map[string]slotPick, bySlot map[string][]scored, slot string) slotRow {
	t.Helper()
	r := buildReport(reportSpec(), 20, "alliance", "human", "", 0, nil, nil, picks, 0, nil, nil, nil, 0, 0, nil, nil, bySlot, 0.1, "")
	for _, row := range r.Slots {
		if row.Slot == slot {
			return row
		}
	}
	t.Fatalf("no %s row in the report", slot)
	return slotRow{}
}

func deadminesPick() *scored {
	return &scored{
		candidate: candidate{ID: 7230, Name: "Smite's Mighty Hammer"},
		Score:     20,
		HasSource: true,
		Source:    itemSource{Kind: "dungeon", Label: "The Deadmines: Defias Overseer", Chance: 14},
	}
}

func TestBuildReportPublishesADropPicksChanceAsAFraction(t *testing.T) {
	row := chanceRow(t, map[string]slotPick{"main_hand": {Item: deadminesPick()}}, nil, "main_hand")
	if row.Chance == nil || *row.Chance != 0.14 {
		t.Fatalf("Chance = %v, want 0.14 (loot.json's 14 percent as a fraction)", row.Chance)
	}
}

func TestBuildReportRoundsAChanceToTheLootTablesOwnPrecision(t *testing.T) {
	pick := deadminesPick()
	pick.Source.Chance = 51.3302
	row := chanceRow(t, map[string]slotPick{"main_hand": {Item: pick}}, nil, "main_hand")
	if row.Chance == nil || *row.Chance != 0.513302 {
		t.Fatalf("Chance = %v, want exactly 0.513302", row.Chance)
	}
}

func TestBuildReportLeavesChanceNilForAPickThatIsNotADrop(t *testing.T) {
	quest := &scored{
		candidate: candidate{ID: 6087, Name: "Chausses of Westfall"},
		Score:     12,
		HasSource: true,
		Source:    itemSource{Kind: "quest", Label: "The Defias Brotherhood"},
	}
	row := chanceRow(t, map[string]slotPick{"legs": {Item: quest}}, nil, "legs")
	if row.Chance != nil {
		t.Fatalf("Chance = %v, want nil: a quest reward is not a drop", *row.Chance)
	}
	encoded, err := json.Marshal(row)
	if err != nil {
		t.Fatal(err)
	}
	var keys map[string]any
	if err := json.Unmarshal(encoded, &keys); err != nil {
		t.Fatal(err)
	}
	if _, ok := keys["chance"]; ok {
		t.Fatalf("JSON %s carries a chance key; a pick that is not a drop must leave it out", encoded)
	}
}

func TestBuildReportPublishesTheNextTwoCandidatesAsRunnersUp(t *testing.T) {
	pick := deadminesPick()
	bySlot := map[string][]scored{
		"main_hand": {
			*pick,
			{candidate: candidate{ID: 270018, Name: "Hammerbone"}, Score: 19, HasSource: true, Source: itemSource{Kind: "quest", Label: "Leaders of the Fang"}},
			{candidate: candidate{ID: 1458, Name: "Large Axe"}, Score: 18, HasSource: true, Source: itemSource{Kind: "world_drop", Label: "World drop", Chance: 0.02}},
			{candidate: candidate{ID: 999, Name: "Third Best"}, Score: 17, HasSource: true, Source: itemSource{Kind: "vendor", Label: "Somebody"}},
		},
	}
	row := chanceRow(t, map[string]slotPick{"main_hand": {Item: pick}}, bySlot, "main_hand")
	want := []runnerUpRow{
		{ItemID: 270018, SourceKind: "quest", Source: "Leaders of the Fang"},
		{ItemID: 1458, SourceKind: "world_drop", Source: "World drop"},
	}
	if len(row.RunnersUp) != len(want) || row.RunnersUp[0] != want[0] || row.RunnersUp[1] != want[1] {
		t.Fatalf("RunnersUp = %+v, want %+v", row.RunnersUp, want)
	}

	encoded, err := json.Marshal(row.RunnersUp[0])
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{"item_id":270018,"source_kind":"quest","source":"Leaders of the Fang"}` {
		t.Fatalf("runner-up JSON = %s, want the pipeline's item_id/source_kind/source keys", encoded)
	}
}

func TestBuildReportPublishesNoRunnersUpForASlotWithNoOtherCandidate(t *testing.T) {
	row := chanceRow(t, map[string]slotPick{"main_hand": {Item: deadminesPick()}}, nil, "main_hand")
	if len(row.RunnersUp) != 0 {
		t.Fatalf("RunnersUp = %+v, want none", row.RunnersUp)
	}
}
