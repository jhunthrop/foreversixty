// api/internal/guilds/readiness_core_test.go
package guilds

import (
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

func testBand() bis.Band {
	return bis.Band{Slots: []bis.Slot{
		{Slot: "chest", ItemID: 100, Alternatives: []bis.Alternative{{ItemID: 200, DpsDelta: -5}}},
		{Slot: "feet", ItemID: 300},
	}}
}

func TestComputeReadinessWithoutConsentChecksNothingGear(t *testing.T) {
	cr := computeReadiness("roster", testBand(), true,
		map[string]int{"chest": 999}, nil, nil, true, 10, true, nil, nil)
	if cr.GearChecked || cr.EnchantsChecked {
		t.Fatalf("checked = gear:%v enchants:%v, want both false without gear consent", cr.GearChecked, cr.EnchantsChecked)
	}
	if cr.ConsumablesState != "unknown" {
		t.Fatalf("consumables = %q, want unknown", cr.ConsumablesState)
	}
}

func TestComputeReadinessGearUpgradeAndEnchantGap(t *testing.T) {
	gear := map[string]int{"chest": 200, "feet": 300} // chest is the listed alt (+5), feet is the pick (no gap)
	enchants := map[string]int{"feet": 42}            // chest equipped, no enchant; feet enchanted
	cr := computeReadiness("gear", testBand(), true, gear, enchants, nil, true, 10, true, nil, nil)
	if !cr.GearChecked || cr.GearGap.Upgrades != 1 || cr.GearGap.GainDps != 5 {
		t.Fatalf("gearGap = %+v, want 1 upgrade worth +5", cr.GearGap)
	}
	if !cr.EnchantsChecked || !reflect.DeepEqual(cr.MissingEnchantSlots, []string{"chest"}) {
		t.Fatalf("missingEnchants = %v, want [chest]", cr.MissingEnchantSlots)
	}
}

func TestComputeReadinessConsumablesGatedOnGearBagsConsent(t *testing.T) {
	short := computeReadiness("gear_bags", bis.Band{}, false, nil, nil, []int{}, true, 0, false, nil, nil)
	if short.ConsumablesState != "short" {
		t.Fatalf("consumables = %q, want short for an empty bags section", short.ConsumablesState)
	}
	stocked := computeReadiness("gear_bags", bis.Band{}, false, nil, nil, []int{13510}, true, 0, false, nil, nil)
	if stocked.ConsumablesState != "stocked" {
		t.Fatalf("consumables = %q, want stocked", stocked.ConsumablesState)
	}
	unknown := computeReadiness("gear", bis.Band{}, false, nil, nil, []int{13510}, true, 0, false, nil, nil)
	if unknown.ConsumablesState != "unknown" {
		t.Fatalf("consumables = %q, want unknown without gear_bags consent even with bag data present", unknown.ConsumablesState)
	}
}

func TestComputeReadinessTalentPointsUnspent(t *testing.T) {
	cr := computeReadiness("roster", bis.Band{}, false, nil, nil, nil, false, 45, true, nil, nil)
	if cr.TalentPointsUnspent != 6 {
		t.Fatalf("unspent = %d, want 6 (51-45)", cr.TalentPointsUnspent)
	}
	full := computeReadiness("roster", bis.Band{}, false, nil, nil, nil, false, 51, true, nil, nil)
	if full.TalentPointsUnspent != 0 {
		t.Fatalf("unspent = %d, want 0 at full spend", full.TalentPointsUnspent)
	}
}

func TestComputeReadinessItemLevelDelta(t *testing.T) {
	lvl, median := 58, 63
	cr := computeReadiness("roster", bis.Band{}, false, nil, nil, nil, false, 0, false, &lvl, &median)
	if cr.ItemLevel == nil || *cr.ItemLevel != 58 {
		t.Fatalf("itemLevel = %v, want 58", cr.ItemLevel)
	}
	if cr.ItemLevelDelta == nil || *cr.ItemLevelDelta != -5 {
		t.Fatalf("itemLevelDelta = %v, want -5", cr.ItemLevelDelta)
	}
}

func TestFailTextOrderAndNeedsCap(t *testing.T) {
	cr := CharacterReadiness{
		GearChecked: true, GearGap: bis.GearGap{Upgrades: 3, GainDps: 23.8},
		EnchantsChecked: true, MissingEnchantSlots: []string{"chest", "feet"},
		ConsumablesState:    "short",
		TalentPointsUnspent: 1,
	}
	fails := cr.failText()
	want := []string{
		"3 gear upgrades waiting (+23.8 DPS)",
		"no enchant: Chest, Boots",
		"consumables short",
		"1 unspent talent point",
	}
	if !reflect.DeepEqual(fails, want) {
		t.Fatalf("failText = %v, want %v", fails, want)
	}
	if cr.FailingCount() != 4 {
		t.Fatalf("failingCount = %d, want 4", cr.FailingCount())
	}
	if got := cr.Needs(3); len(got) != 3 {
		t.Fatalf("Needs(3) = %v, want 3 entries", got)
	}
	if got := cr.Needs(0); len(got) != 4 {
		t.Fatalf("Needs(0) = %v, want uncapped (4)", got)
	}
}

func TestFailTextCleanRowIsEmpty(t *testing.T) {
	cr := CharacterReadiness{GearChecked: true, EnchantsChecked: true, ConsumablesState: "stocked"}
	if got := cr.failText(); len(got) != 0 {
		t.Fatalf("failText = %v, want empty for a clean row", got)
	}
}
