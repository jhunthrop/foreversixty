package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

func TestLoadBisGearReadsRealAllianceBand60(t *testing.T) {
	dir, err := bisDir()
	if err != nil {
		t.Fatal(err)
	}
	gear, err := loadBisGear(dir, "warrior-fury.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(gear) == 0 {
		t.Fatal("warrior-fury.json band 60 alliance: got no gear")
	}
	for _, slot := range []string{"head", "chest", "main_hand"} {
		if _, ok := gear[slot]; !ok {
			t.Errorf("missing slot %q in %v", slot, gear)
		}
	}
}

func TestLoadBisGearEmptyFileNameIsEmptyGear(t *testing.T) {
	dir, err := bisDir()
	if err != nil {
		t.Fatal(err)
	}
	gear, err := loadBisGear(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(gear) != 0 {
		t.Fatalf("gear = %v, want empty (no BiS file means no invented item ids)", gear)
	}
}

func TestEveryRosterBisFileLoadsRealItemIDs(t *testing.T) {
	dir, err := bisDir()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range mockRoster() {
		if c.BisFile == "" {
			continue
		}
		gear, err := loadBisGear(dir, c.BisFile)
		if err != nil {
			t.Fatalf("%s (%s): %v", c.Name, c.BisFile, err)
		}
		if len(gear) == 0 {
			t.Errorf("%s (%s): no gear loaded", c.Name, c.BisFile)
		}
		for slot, item := range gear {
			if item <= 0 {
				t.Errorf("%s (%s): slot %s has non-positive item id %d", c.Name, c.BisFile, slot, item)
			}
		}
	}
}

func TestBuildFS1ProducesAValidDecodableExport(t *testing.T) {
	gear := map[string]gearPiece{
		"head":  {ItemID: 250528, Enchant: 0},
		"chest": {ItemID: 12345, Enchant: 41},
	}
	export := buildFS1("warrior", "human", talentString(roleDPS), gear, []string{"mining", "blacksmithing"}, nil)
	d, ok := fs1.Decode(export)
	if !ok {
		t.Fatalf("fs1.Decode(%q) ok = false", export)
	}
	if d.ClassSlug != "warrior" || d.RaceSlug != "human" {
		t.Fatalf("class/race = %s/%s, want warrior/human", d.ClassSlug, d.RaceSlug)
	}
	if !d.HasLevel || d.Level != fs1Level {
		t.Fatalf("level = %d hasLevel=%v, want %d/true", d.Level, d.HasLevel, fs1Level)
	}
	if d.Gear["head"] != 250528 || d.Gear["chest"] != 12345 {
		t.Fatalf("gear = %v, want head=250528 chest=12345", d.Gear)
	}
}

func TestBuildFS1WithNoGearStillDecodes(t *testing.T) {
	export := buildFS1("priest", "gnome", talentString(roleHealer), map[string]gearPiece{}, nil, nil)
	d, ok := fs1.Decode(export)
	if !ok {
		t.Fatalf("fs1.Decode(%q) ok = false", export)
	}
	if len(d.Gear) != 0 {
		t.Fatalf("gear = %v, want empty", d.Gear)
	}
}

func TestEveryRosterCharacterProducesADecodableExport(t *testing.T) {
	dir, err := bisDir()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range mockRoster() {
		slots, err := loadBisSlots(dir, c.BisFile)
		if err != nil {
			t.Fatalf("%s: %v", c.Name, err)
		}
		export := buildFS1(c.Class, c.RaceSlug, talentStringFor(c), gearWithReadiness(c, slots),
			professionsFor(c.Class), bagsFor(c))
		if _, ok := fs1.Decode(export); !ok {
			t.Fatalf("%s: export %q did not decode", c.Name, export)
		}
	}
}

func TestGearWithReadinessAppliesNonBisSlotsAndMissingEnchant(t *testing.T) {
	dir, err := bisDir()
	if err != nil {
		t.Fatal(err)
	}
	dunmaro := mustFind(t, mockRoster(), "Dunmaro") // NonBisSlots: 2
	slots, err := loadBisSlots(dir, dunmaro.BisFile)
	if err != nil {
		t.Fatal(err)
	}
	gear := gearWithReadiness(dunmaro, slots)
	bis, err := loadBisGear(dir, dunmaro.BisFile)
	if err != nil {
		t.Fatal(err)
	}
	swapped := 0
	for slot, piece := range gear {
		if piece.ItemID != bis[slot] {
			swapped++
		}
	}
	if swapped != dunmaro.NonBisSlots {
		t.Fatalf("swapped %d slots, want %d", swapped, dunmaro.NonBisSlots)
	}

	shiverknife := mustFind(t, mockRoster(), "Shiverknife") // MissingEnchantSlot: "wrist"
	slots, err = loadBisSlots(dir, shiverknife.BisFile)
	if err != nil {
		t.Fatal(err)
	}
	gear = gearWithReadiness(shiverknife, slots)
	if piece, ok := gear["wrist"]; ok && piece.Enchant != 0 {
		t.Fatalf("wrist enchant = %d, want 0 (Shiverknife's own missing-enchant slot)", piece.Enchant)
	}
	if piece, ok := gear["chest"]; ok && piece.Enchant == 0 {
		t.Fatalf("chest enchant = 0, want a real enchant (only wrist is missing one)")
	}
}

func mustFind(t *testing.T, roster []mockCharacter, name string) mockCharacter {
	t.Helper()
	for _, c := range roster {
		if c.Name == name {
			return c
		}
	}
	t.Fatalf("no roster character named %q", name)
	return mockCharacter{}
}
