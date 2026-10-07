package leveling

import "testing"

func TestKitConsumesIsRoguePoisonsFromTwenty(t *testing.T) {
	if got := KitConsumes("rogue", 19); got != nil {
		t.Fatalf("rogue at 19 = %v, want nothing (the poison quest is level 20)", got)
	}
	if got := KitConsumes("rogue", 20); len(got) != 2 {
		t.Fatalf("rogue at 20 = %v, want Instant Poison on both weapons", got)
	}
	if got := KitConsumes("rogue-assassination", 20); len(got) != 2 || got[0] != "main_hand_imbue:instant_poison" {
		t.Fatalf("rogue-assassination at 20 = %v, want Instant Poison on both weapons (Deadly is not learned yet)", got)
	}
	// Mutilate's bonus needs a lingering poison, so Assassination carries
	// Deadly Poison on the main hand once it is learned at 30.
	if got := KitConsumes("rogue-assassination", 30); len(got) != 2 || got[0] != "main_hand_imbue:deadly_poison" || got[1] != "off_hand_imbue:instant_poison" {
		t.Fatalf("rogue-assassination at 30 = %v, want Deadly Poison main hand and Instant Poison off hand", got)
	}
	if got := KitConsumes("rogue-combat", 60); got[0] != "main_hand_imbue:instant_poison" {
		t.Fatalf("rogue-combat at 60 = %v, want Instant Poison on both weapons", got)
	}
	if got := KitConsumes("rogue-combat", 19); got != nil {
		t.Fatalf("rogue-combat at 19 = %v, want nothing", got)
	}
	if got := KitConsumes("warrior", 60); got != nil {
		t.Fatalf("warrior = %v, want nothing", got)
	}
}

// TestKitConsumesIsShamanEnhancementWeaponImbues: shamans cannot dual
// wield in Forever (owner rule, 2026-09-30, bis-ranker-integrity-13's
// brief), so enhancement's off hand is a shield or held item, never a
// second imbued weapon - one imbue, always on the main hand.
func TestKitConsumesIsShamanEnhancementWeaponImbues(t *testing.T) {
	got := KitConsumes("shaman-enhancement", 10)
	want := []string{"main_hand_imbue:rockbiter_weapon"}
	if len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("shaman-enhancement at 10 = %v, want %v (Rockbiter Weapon on the main hand; Windfury Weapon is not learnable until 30)", got, want)
	}

	got = KitConsumes("shaman-enhancement", 29)
	if len(got) != 1 || got[0] != "main_hand_imbue:rockbiter_weapon" {
		t.Fatalf("shaman-enhancement at 29 = %v, want Rockbiter Weapon still on the main hand", got)
	}

	got = KitConsumes("shaman-enhancement", 30)
	want = []string{"main_hand_imbue:windfury_weapon"}
	if len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("shaman-enhancement at 30 = %v, want %v (Windfury Weapon on the main hand once learned)", got, want)
	}

	got = KitConsumes("shaman-enhancement", 60)
	if len(got) != 1 || got[0] != "main_hand_imbue:windfury_weapon" {
		t.Fatalf("shaman-enhancement at 60 = %v, want Windfury Weapon on the main hand, nothing on the off hand", got)
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

// TestDualWieldSpecsExcludesShaman: enhancement's off hand is a
// shield/held-item pool, not a weapon pool - it must not be a
// DualWieldSpecs member. Dual wield stays for rogue's three specs,
// warrior-fury and hunter's three specs, and no other spec is a member.
func TestDualWieldSpecsExcludesShaman(t *testing.T) {
	if DualWieldSpecs["shaman-enhancement"] {
		t.Fatal("shaman-enhancement must not be a DualWieldSpecs member: shamans cannot dual wield in Forever")
	}
	want := map[string]bool{
		"rogue-assassination":  true,
		"rogue-combat":         true,
		"rogue-subtlety":       true,
		"warrior-fury":         true,
		"hunter-beast-mastery": true,
		"hunter-marksmanship":  true,
		"hunter-survival":      true,
	}
	if len(DualWieldSpecs) != len(want) {
		t.Fatalf("DualWieldSpecs = %v, want exactly %v", DualWieldSpecs, want)
	}
	for spec := range want {
		if !DualWieldSpecs[spec] {
			t.Fatalf("DualWieldSpecs missing %q", spec)
		}
	}
}

func TestKitBuffsMageCarriesArcaneIntellectFromLevelOne(t *testing.T) {
	for _, spec := range []string{"mage", "mage-arcane", "mage-fire", "mage-frost"} {
		for _, level := range []int{1, 30, 60} {
			got := KitBuffs(spec, level)
			if len(got) != 1 || got[0] != "arcane_brilliance" {
				t.Fatalf("%s at %d = %v, want arcane_brilliance", spec, level, got)
			}
		}
	}
}

// Improved Mark of the Wild is a baseline passive, not a tree talent, in
// the live build: no druid guide build takes it, so no ":improved".
func TestKitBuffsDruidCarriesPlainMarkOfTheWild(t *testing.T) {
	for _, spec := range []string{"druid-balance", "druid-feral", "druid-restoration"} {
		got := KitBuffs(spec, 1)
		if len(got) != 1 || got[0] != "gift_of_the_wild" {
			t.Fatalf("%s at 1 = %v, want gift_of_the_wild", spec, got)
		}
	}
}

func TestKitBuffsPaladinCarriesBlessingOfMightFromFour(t *testing.T) {
	for _, spec := range []string{"paladin-holy", "paladin-protection", "paladin-retribution"} {
		if got := KitBuffs(spec, 3); got != nil {
			t.Fatalf("%s at 3 = %v, want nothing (rank 1 is learned at 4)", spec, got)
		}
		for _, level := range []int{4, 60} {
			got := KitBuffs(spec, level)
			if len(got) != 1 || got[0] != "blessing_of_might" {
				t.Fatalf("%s at %d = %v, want blessing_of_might", spec, level, got)
			}
		}
	}
}

// Survival's Expose Prey rolls only against a Hunter's Mark target, and
// the engine models the mark as a target debuff rather than a castable
// spell, so the survival kit carries it from the level the spell is
// learned (6, client spell 1130).
func TestKitBuffsSurvivalCarriesHuntersMarkFromSix(t *testing.T) {
	if got := KitBuffs("hunter-survival", 5); got != nil {
		t.Fatalf("hunter-survival at 5 = %v, want nothing (rank 1 is learned at 6)", got)
	}
	for _, level := range []int{6, 60} {
		got := KitBuffs("hunter-survival", level)
		if len(got) != 1 || got[0] != "hunters_mark" {
			t.Fatalf("hunter-survival at %d = %v, want hunters_mark", level, got)
		}
	}
}

func TestKitBuffsOtherClassesCarryNothing(t *testing.T) {
	for _, spec := range []string{"warrior-arms", "rogue-combat", "priest-shadow", "warlock-affliction", "shaman-enhancement", "hunter-beast-mastery", "hunter-marksmanship", "magefoo"} {
		if got := KitBuffs(spec, 60); got != nil {
			t.Fatalf("%s = %v, want nothing", spec, got)
		}
	}
}
