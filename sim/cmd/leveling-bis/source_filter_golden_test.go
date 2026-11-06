package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

// sourceFilterGoldenPath is the band web/src/lib/bis/source-filter.test.ts
// resolves too: the published rows of the dual-wield fixture, and what
// resolveFilteredSlots gives under a set of filters. The two resolvers
// must agree on every case. FOREVER_UPDATE_GOLDEN=1 rewrites it.
const sourceFilterGoldenPath = "testdata/source-filter-band.json"

type sourceFilterGoldenCase struct {
	Kinds []string       `json:"kinds"`
	Want  map[string]int `json:"want"`
}

type sourceFilterGolden struct {
	Slots []slotRow                `json:"slots"`
	Cases []sourceFilterGoldenCase `json:"cases"`
}

func sourceFilterGoldenCases() [][]string {
	cases := [][]string{{}, filterKindOrder}
	for _, k := range filterKindOrder {
		cases = append(cases, []string{k})
	}
	return append(cases, []string{"vendor", "rep"}, []string{"quest", "crafted", "pvp"}, []string{"world", "raid"})
}

func buildSourceFilterGolden(spec string) sourceFilterGolden {
	bySlot := candidatesBySlot(sourceFilterFixturePool(spec == "rogue-combat"))
	picks := pick(spec, bySlot)
	rows := withSourceKindPicks(fixtureRows(picks), sourceFilterInputs{Spec: spec, Picks: picks, PickBySlot: bySlot, ReferenceDPSPerPoint: 0.5})
	var kept []slotRow
	for _, r := range rows {
		if r.ItemID != 0 || len(r.ByKind) > 0 {
			kept = append(kept, r)
		}
	}
	golden := sourceFilterGolden{Slots: kept}
	for _, kinds := range sourceFilterGoldenCases() {
		allowed := make(map[string]bool, len(kinds))
		for _, k := range kinds {
			allowed[k] = true
		}
		want := map[string]int{}
		for slot, got := range resolveFilteredSlots(kept, allowed) {
			if got != nil {
				want[slot] = got.ItemID
			}
		}
		sorted := append([]string{}, kinds...)
		sort.Strings(sorted)
		golden.Cases = append(golden.Cases, sourceFilterGoldenCase{Kinds: sorted, Want: want})
	}
	return golden
}

func TestSourceFilterGoldenSharedWithTheWeb(t *testing.T) {
	got := map[string]sourceFilterGolden{
		"two_hand":   buildSourceFilterGolden("warrior-arms"),
		"dual_wield": buildSourceFilterGolden("rogue-combat"),
	}
	if os.Getenv("FOREVER_UPDATE_GOLDEN") != "" {
		b, err := json.MarshalIndent(got, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.FromSlash(sourceFilterGoldenPath), append(b, '\n'), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	b, err := os.ReadFile(filepath.FromSlash(sourceFilterGoldenPath))
	if err != nil {
		t.Fatalf("%v (FOREVER_UPDATE_GOLDEN=1 writes it)", err)
	}
	var want map[string]sourceFilterGolden
	if err := json.Unmarshal(b, &want); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(roundTrip(t, got), want) {
		t.Errorf("%s is stale; FOREVER_UPDATE_GOLDEN=1 rewrites it", sourceFilterGoldenPath)
	}
}

func roundTrip(t *testing.T, v map[string]sourceFilterGolden) map[string]sourceFilterGolden {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]sourceFilterGolden
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	return out
}
