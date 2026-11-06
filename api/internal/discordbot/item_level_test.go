package discordbot

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

func wornAt(slot string, level *int) rankings.GearSlot {
	return rankings.GearSlot{Slot: slot, Worn: &rankings.GearWorn{ID: 1, ItemLevel: level}}
}

func ilvl(n int) *int { return &n }

// The one item-level rule (upgrades.EquippedItemLevel): every worn slot counts, an empty slot
// (the off hand under a two-hander included) does not, one unknown item hides the figure, and
// the mean rounds to a whole level as the roster and readiness board print it.
func TestAverageItemLevelRule(t *testing.T) {
	twoHander := []rankings.GearSlot{
		wornAt("head", ilvl(60)), wornAt("chest", ilvl(63)), wornAt("main_hand", ilvl(66)),
		{Slot: "off_hand"}, {Slot: "ranged"},
	}
	if got, ok := equippedItemLevel(twoHander); !ok || got != "63" {
		t.Fatalf("two-hander with an empty off hand: %q %v, want 63", got, ok)
	}
	withUnknown := append(twoHander, wornAt("waist", nil))
	if got, ok := equippedItemLevel(withUnknown); ok || got != "" {
		t.Fatalf("an unknown worn item: %q %v, want no figure", got, ok)
	}
	if _, ok := equippedItemLevel([]rankings.GearSlot{{Slot: "head"}}); ok {
		t.Fatal("nothing worn has no figure")
	}
}

// A guild member's /char reads the guild's own figure, so a healer whose export carries no gear
// still shows the item level the roster and readiness board show (guild round 3, items 5 and 6).
func TestCharItemLevelReadsTheGuildsFigureFirst(t *testing.T) {
	gear := []rankings.GearSlot{wornAt("head", ilvl(60)), wornAt("chest", ilvl(59)), wornAt("legs", ilvl(59)),
		wornAt("feet", ilvl(60)), wornAt("hands", ilvl(59))}
	if got, ok := charItemLevel(charData{card: rankings.Character{Gear: gear}}); !ok || got != "59" {
		t.Fatalf("card only: %q %v, want 59 (59.4 rounded)", got, ok)
	}
	healer := charData{facts: &guilds.ReadinessFacts{ItemLevel: ilvl(63)}}
	if got, ok := charItemLevel(healer); !ok || got != "63" {
		t.Fatalf("guild figure: %q %v, want 63", got, ok)
	}
}
