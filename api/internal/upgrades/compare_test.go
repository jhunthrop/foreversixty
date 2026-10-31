package upgrades

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func compareBand() bis.Band {
	per := 0.5
	return bis.Band{ReferenceDpsPerPoint: &per, Weights: []bis.Weight{{Stat: "agility", Weight: 2}}, Slots: []bis.Slot{
		{Slot: "head", ItemID: 1, ItemName: "Pick Helm", Score: ptr(40.0), Alternatives: []bis.Alternative{{ItemID: 2, DpsDelta: -3}}},
		{Slot: "neck", ItemID: 20, ItemName: "Pick Amulet", Score: ptr(30.0)},
		{Slot: "chest", ItemID: 30, ItemName: "Pick Tunic", Score: ptr(10.0)},
		{Slot: "feet", ItemID: 40, ItemName: "Pick Boots", Score: ptr(20.0), EffectUnmodelled: true},
		{Slot: "trinket1", ItemID: 50, ItemName: "Pick Trinket"},
		{Slot: "off_hand", ItemID: 0},
		{Slot: "ranged", ItemID: 60, ItemName: "Pick Bow"},
		{Slot: "legs", ItemID: 70, ItemName: "Pick Legs", Score: ptr(5.0)},
	}}
}

func compareItems(id int) (trees.Item, bool) {
	items := map[int]trees.Item{
		1: {ID: 1, Name: "Pick Helm"}, 2: {ID: 2, Name: "Worn Helm"}, 21: {ID: 21, Name: "Worn Amulet", Stats: map[string]int{"agility": 5}},
		70: {ID: 70, Name: "Pick Legs"}, 61: {ID: 61, Name: "Worn Bow"}, 31: {ID: 31, Name: "Strong Tunic", Stats: map[string]int{"agility": 15}},
	}
	it, ok := items[id]
	return it, ok
}

// Compare is web/src/lib/home/upgrades.ts's upgradesFor: the same rows, gains, already-best
// slots, total and not-sim-checked count, in the bot's order.
func TestCompare(t *testing.T) {
	gear := map[string]int{"head": 2, "neck": 21, "chest": 31, "feet": 99, "ranged": 61, "legs": 70}
	got := Compare(compareBand(), gear, compareItems)
	want := []struct {
		slot      string
		gain      *float64
		unchecked bool
		known     bool
		group     Group
	}{
		{"trinket1", nil, false, false, GroupEmpty},
		{"neck", ptr(10.0), false, true, GroupMeasured},
		{"head", ptr(3.0), false, true, GroupMeasured},
		{"chest", ptr(-10.0), false, true, GroupMeasured},
		{"feet", nil, false, false, GroupUnmodelled},
		{"ranged", nil, true, true, GroupUnmodelled},
	}
	if len(got.Upgrades) != len(want) {
		t.Fatalf("upgrades = %+v", got.Upgrades)
	}
	for i, w := range want {
		u := got.Upgrades[i]
		if u.Slot != w.slot || !sameGain(u.Gain, w.gain) || u.NotSimChecked != w.unchecked || u.WornKnown != w.known || u.Group != w.group {
			t.Errorf("row %d = %+v, want %+v", i, u, w)
		}
	}
	if len(got.AlreadyBest) != 1 || got.AlreadyBest[0].Slot != "legs" || got.AlreadyBest[0].ItemName != "Pick Legs" {
		t.Fatalf("already best = %+v", got.AlreadyBest)
	}
	if !sameGain(&got.TotalGain, ptr(3.0)) || got.NotSimChecked != 1 || got.Unit != UnitDPS {
		t.Fatalf("total %v unchecked %d unit %q", got.TotalGain, got.NotSimChecked, got.Unit)
	}
}
