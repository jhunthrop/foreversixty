package discordbot

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

func wornAt(slot string, level *int) rankings.GearSlot {
	return rankings.GearSlot{Slot: slot, Worn: &rankings.GearWorn{ID: 1, ItemLevel: level}}
}

func ilvl(n int) *int { return &n }

// The paperdoll's item-level rule (design/specs/2026-10-31-paperdoll.md §3.5): every worn slot
// counts, an empty slot (the off hand under a two-hander included) does not, and one unknown
// item hides the figure.
func TestAverageItemLevelRule(t *testing.T) {
	twoHander := []rankings.GearSlot{
		wornAt("head", ilvl(60)), wornAt("chest", ilvl(63)), wornAt("main_hand", ilvl(66)),
		{Slot: "off_hand"}, {Slot: "ranged"},
	}
	if got, ok := averageItemLevel(twoHander); !ok || got != "63.0" {
		t.Fatalf("two-hander with an empty off hand: %q %v, want 63.0", got, ok)
	}
	withUnknown := append(twoHander, wornAt("waist", nil))
	if got, ok := averageItemLevel(withUnknown); ok || got != "" {
		t.Fatalf("an unknown worn item: %q %v, want no figure", got, ok)
	}
	if _, ok := averageItemLevel([]rankings.GearSlot{{Slot: "head"}}); ok {
		t.Fatal("nothing worn has no figure")
	}
}
