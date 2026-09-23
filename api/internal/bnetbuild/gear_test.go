// api/internal/bnetbuild/gear_test.go
package bnetbuild

import (
	"strings"
	"testing"
)

func TestEncodeGearMapsEveryFixtureSlotAndSkipsShirtAndTabard(t *testing.T) {
	raw := readEraKiloz(t, "equipment.json")
	gear, skipped, noSuffix, err := encodeGear(raw, SuffixTable{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(gear, "head=21329:2583") {
		t.Fatalf("gear = %q, want head=21329:2583 (Conqueror's Crown, Presence of Might)", gear)
	}
	if !strings.Contains(gear, "main_hand=22816:1900") || !strings.Contains(gear, "off_hand=18828:1900") {
		t.Fatalf("gear = %q, want both weapon hands with their Crusader enchant", gear)
	}
	if !strings.Contains(gear, "trinket1=11815") {
		t.Fatalf("gear = %q, want trinket1=11815 (Hand of Justice, no enchant)", gear)
	}
	wantSkipped := map[string]bool{"SHIRT": false, "TABARD": false}
	for _, s := range skipped {
		if _, ok := wantSkipped[s]; ok {
			wantSkipped[s] = true
		}
	}
	if !wantSkipped["SHIRT"] || !wantSkipped["TABARD"] {
		t.Fatalf("skippedSlots = %v, want SHIRT and TABARD both reported", skipped)
	}
	if len(noSuffix) != 0 {
		t.Fatalf("noSuffix = %v, want none: no fixture item's name matches a suffix", noSuffix)
	}
}

func TestEncodeGearMatchesASuffixedItemName(t *testing.T) {
	raw := []byte(`{"equipped_items":[{"item":{"id":1234},"slot":{"type":"WAIST"},"name":"Girdle of the Falcon","enchantments":[]}]}`)
	gear, _, noSuffix, err := encodeGear(raw, SuffixTable{"of the Falcon": 14})
	if err != nil {
		t.Fatal(err)
	}
	if gear != "waist=1234::14" {
		t.Fatalf("gear = %q, want waist=1234::14 (item id, no enchant, suffix 14)", gear)
	}
	if len(noSuffix) != 0 {
		t.Fatalf("noSuffix = %v, want none: this item matched", noSuffix)
	}
}

func TestEncodeGearOnlyEmitsThePermanentEnchant(t *testing.T) {
	raw := []byte(`{"equipped_items":[{"item":{"id":999},"slot":{"type":"FEET"},"name":"Test Boots",
		"enchantments":[{"enchantment_id":1,"enchantment_slot":{"type":"TEMPORARY"}},
		                {"enchantment_id":911,"enchantment_slot":{"type":"PERMANENT"}}]}]}`)
	gear, _, _, err := encodeGear(raw, SuffixTable{})
	if err != nil {
		t.Fatal(err)
	}
	if gear != "feet=999:911" {
		t.Fatalf("gear = %q, want feet=999:911 (the PERMANENT enchant only, not the TEMPORARY one)", gear)
	}
}
