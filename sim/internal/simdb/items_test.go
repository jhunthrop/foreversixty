package simdb

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// Every slot name a row can produce must be a slot the envelope and
// sim/request know, or an expansion would place an item somewhere
// nothing can equip it.
func TestEverySlotNameIsInTheVocabulary(t *testing.T) {
	n, err := Len()
	if err != nil {
		t.Fatal(err)
	}
	if n == 0 {
		t.Fatal("the embedded table is empty")
	}
	db, err := load()
	if err != nil {
		t.Fatal(err)
	}
	var placeable int
	for _, row := range db.Items {
		item, ok := Lookup(int(row.Id))
		if !ok {
			t.Fatalf("item %d is in the table and Lookup does not find it", row.Id)
		}
		if len(item.Slots) > 0 {
			placeable++
		}
		for _, slot := range item.Slots {
			if !slices.Contains(api.GearSlots, slot) {
				t.Errorf("item %d (%s) is placeable in %q, which is not a gear slot", item.ID, item.Name, slot)
			}
		}
	}
	if placeable == 0 {
		t.Error("no item in the build is placeable in any slot")
	}
	t.Logf("items=%d placeable=%d", n, placeable)
}

// The engine's own eligibility rules, re-expressed: a ring and a
// trinket go in either of two slots, a one-hander in either hand, a
// two-hander and a main-hand-only weapon in the main hand, an off-hand
// in the off hand.
func TestSlotsFollowTheEnginesEligibilityRules(t *testing.T) {
	cases := []struct {
		name  string
		row   *proto.SimItem
		slots []string
	}{
		{"a helm", &proto.SimItem{Type: proto.ItemType_ItemTypeHead}, []string{"head"}},
		{"a ring", &proto.SimItem{Type: proto.ItemType_ItemTypeFinger}, []string{"finger1", "finger2"}},
		{"a trinket", &proto.SimItem{Type: proto.ItemType_ItemTypeTrinket}, []string{"trinket1", "trinket2"}},
		{"a bow", &proto.SimItem{Type: proto.ItemType_ItemTypeRanged}, []string{"ranged"}},
		{"a two-hander", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeTwoHand}, []string{"main_hand"}},
		{"a main-hand", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeMainHand}, []string{"main_hand"}},
		{"an off-hand", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeOffHand}, []string{"off_hand"}},
		{"a one-hander", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeOneHand}, []string{"main_hand", "off_hand"}},
		{"an unknown type", &proto.SimItem{Type: proto.ItemType_ItemTypeUnknown}, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := itemFrom(c.row)
			if !slices.Equal(got.Slots, c.slots) {
				t.Errorf("slots = %v, want %v", got.Slots, c.slots)
			}
		})
	}
}

func TestRowFieldsCross(t *testing.T) {
	got := itemFrom(&proto.SimItem{
		Id:                  19352,
		Name:                "Vis'kag the Bloodletter",
		Type:                proto.ItemType_ItemTypeWeapon,
		HandType:            proto.HandType_HandTypeOneHand,
		WeaponType:          proto.WeaponType_WeaponTypeSword,
		ArmorType:           proto.ArmorType_ArmorTypeUnknown,
		ClassAllowlist:      []proto.Class{proto.Class_ClassWarrior, proto.Class_ClassRogue},
		RequiredLevel:       60,
		Unique:              true,
		FactionRestriction:  proto.SimItem_FACTION_RESTRICTION_HORDE_ONLY,
		RandomSuffixOptions: []int32{1805, 1806},
	})
	if got.ID != 19352 || got.Name != "Vis'kag the Bloodletter" {
		t.Errorf("identity: %+v", got)
	}
	if got.HandType != HandOne || got.WeaponType != "sword" {
		t.Errorf("weapon: %+v", got)
	}
	if !slices.Equal(got.Classes, []string{"warrior", "rogue"}) {
		t.Errorf("classes = %v", got.Classes)
	}
	if got.RequiredLevel != 60 || !got.Unique || got.Faction != FactionHorde {
		t.Errorf("restrictions: %+v", got)
	}
	if !slices.Equal(got.SuffixOptions, []int{1805, 1806}) {
		t.Errorf("suffixes = %v", got.SuffixOptions)
	}
}

func TestLookupMissesCleanly(t *testing.T) {
	if _, ok := Lookup(0); ok {
		t.Error("item 0 resolved")
	}
	if _, ok := Lookup(-1); ok {
		t.Error("a negative id resolved")
	}
}
