package main

import (
	"reflect"
	"testing"
)

// Every loot.json kind (web/src/lib/sim/loot.ts LOOT_KINDS) maps to one
// filter kind; the web twin (source-filter.test.ts) pins the same table.
func TestEverySourceKindMapsToOneFilterKind(t *testing.T) {
	want := map[string]string{
		"raid": "raid", "dungeon": "dungeon",
		"world": "world", "world_drop": "world", "zone": "world",
		"vendor": "vendor", "crafted": "crafted", "rep": "rep", "pvp": "pvp", "quest": "quest",
	}
	for kind, filter := range want {
		got, ok := filterKindOfSource(kind)
		if !ok || got != filter {
			t.Errorf("filterKindOfSource(%q) = %q, %v; want %q", kind, got, ok, filter)
		}
	}
	if len(sourceKindToFilterKind) != len(want) {
		t.Errorf("sourceKindToFilterKind has %d kinds, want %d", len(sourceKindToFilterKind), len(want))
	}
	if _, ok := filterKindOfSource("event"); ok {
		t.Error("an unknown kind mapped to a filter kind")
	}
	seen := map[string]bool{}
	for _, filter := range sourceKindToFilterKind {
		seen[filter] = true
	}
	if len(seen) != len(filterKindOrder) {
		t.Errorf("filter kinds in the mapping = %v, want exactly filterKindOrder %v", seen, filterKindOrder)
	}
}

// An item belongs to every kind it can be had from; per kind the source
// named is the one sourceFor would prefer among that kind's sources.
func TestSourcesByFilterKindKeepsEveryObtainableKind(t *testing.T) {
	obtainable := []itemSource{
		{Kind: "quest", Label: "The Quest"},
		{Kind: "dungeon", Label: "Deadmines: Smite", Chance: 10},
		{Kind: "dungeon", Label: "Deadmines: Cookie", Chance: 20},
		{Kind: "world_drop", Label: "World drop"},
		{Kind: "world", Label: "Some Mob"},
		{Kind: "zone", Label: "Westfall"},
		{Kind: "quest", Label: "Mage Only", Classes: []string{"mage"}},
	}
	got := sourcesByFilterKind(obtainable, "warrior")
	if kinds := orderedFilterKinds(got); !reflect.DeepEqual(kinds, []string{"dungeon", "world", "quest"}) {
		t.Fatalf("kinds = %v, want dungeon, world, quest", kinds)
	}
	if got["dungeon"].Label != "Deadmines: Cookie" {
		t.Errorf("dungeon source = %q, want the higher-chance boss", got["dungeon"].Label)
	}
	if got["world"].Kind != "world_drop" {
		t.Errorf("world source kind = %q, want world_drop (sourceKindPriority before world)", got["world"].Kind)
	}
	if got["quest"].Label != "The Quest" {
		t.Errorf("quest source = %q, want the quest the class can take", got["quest"].Label)
	}
	if only := sourcesByFilterKind([]itemSource{{Kind: "quest", Label: "Mage Only", Classes: []string{"mage"}}}, "warrior"); len(only) != 0 {
		t.Errorf("a class-restricted quest gave a warrior kinds %v", only)
	}
}

func TestFilterKindSourcesFallsBackToThePublishedSource(t *testing.T) {
	c := scored{Source: itemSource{Kind: "rep", Label: "Stormwind"}, HasSource: true}
	if got := orderedFilterKinds(filterKindSources(c)); !reflect.DeepEqual(got, []string{"rep"}) {
		t.Errorf("kinds = %v, want [rep]", got)
	}
	if got := filterKindSources(scored{}); got != nil {
		t.Errorf("an unsourced candidate has kinds %v", got)
	}
}
