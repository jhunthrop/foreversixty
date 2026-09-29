package leveling

import "testing"

func TestKitConsumesIsRoguePoisonsFromTwenty(t *testing.T) {
	if got := KitConsumes("rogue", 19); got != nil {
		t.Fatalf("rogue at 19 = %v, want nothing (the poison quest is level 20)", got)
	}
	if got := KitConsumes("rogue", 20); len(got) != 2 {
		t.Fatalf("rogue at 20 = %v, want Instant Poison on both weapons", got)
	}
	if got := KitConsumes("rogue-assassination", 20); len(got) != 2 {
		t.Fatalf("rogue-assassination at 20 = %v, want Instant Poison on both weapons", got)
	}
	if got := KitConsumes("rogue-combat", 19); got != nil {
		t.Fatalf("rogue-combat at 19 = %v, want nothing", got)
	}
	if got := KitConsumes("warrior", 60); got != nil {
		t.Fatalf("warrior = %v, want nothing", got)
	}
}

func TestKitConsumesIsShamanEnhancementWeaponImbues(t *testing.T) {
	got := KitConsumes("shaman-enhancement", 10)
	want := []string{"main_hand_imbue:rockbiter_weapon", "off_hand_imbue:rockbiter_weapon"}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("shaman-enhancement at 10 = %v, want %v (Rockbiter Weapon on both weapons; Windfury Weapon is not learnable until 30)", got, want)
	}

	got = KitConsumes("shaman-enhancement", 29)
	if len(got) != 2 || got[0] != "main_hand_imbue:rockbiter_weapon" {
		t.Fatalf("shaman-enhancement at 29 = %v, want Rockbiter Weapon still on both weapons", got)
	}

	got = KitConsumes("shaman-enhancement", 30)
	want = []string{"main_hand_imbue:windfury_weapon", "off_hand_imbue:rockbiter_weapon"}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("shaman-enhancement at 30 = %v, want %v (Windfury Weapon on the main hand once learned)", got, want)
	}

	got = KitConsumes("shaman-enhancement", 60)
	if len(got) != 2 || got[0] != "main_hand_imbue:windfury_weapon" || got[1] != "off_hand_imbue:rockbiter_weapon" {
		t.Fatalf("shaman-enhancement at 60 = %v, want Windfury Weapon MH / Rockbiter Weapon OH still", got)
	}

	if got := KitConsumes("shaman-elemental", 60); got != nil {
		t.Fatalf("shaman-elemental = %v, want nothing (not a weapon-imbue spec)", got)
	}
	if got := KitConsumes("shaman-restoration", 60); got != nil {
		t.Fatalf("shaman-restoration = %v, want nothing (not a weapon-imbue spec)", got)
	}
	if got := KitConsumes("shaman", 60); got != nil {
		t.Fatalf("bare shaman class = %v, want nothing (the kit is a spec property, not a class one)", got)
	}
}
