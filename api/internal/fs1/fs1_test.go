package fs1

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestDecodeReadsTheFullHeadAndSections(t *testing.T) {
	export := "FS1:1.60.1.70009:hunter:troll:503200000/0/0:head=12640,chest=11726|level=24|who=Bow:Skyborne"
	d, ok := Decode(export)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if d.DataBuild != "1.60.1.70009" || d.ClassSlug != "hunter" || d.RaceSlug != "troll" {
		t.Fatalf("head = %+v", d)
	}
	wantGear := map[string]int{"head": 12640, "chest": 11726}
	if !reflect.DeepEqual(d.Gear, wantGear) {
		t.Fatalf("gear = %v, want %v", d.Gear, wantGear)
	}
	wantTrees := []string{"503200000", "0", "0"}
	if !reflect.DeepEqual(d.Talents.Trees, wantTrees) {
		t.Fatalf("trees = %v, want %v", d.Talents.Trees, wantTrees)
	}
	// "503200000" base-36 digit sum: 5+0+3+2+0+0+0+0+0 = 10.
	wantPoints := []int{10, 0, 0}
	if !reflect.DeepEqual(d.Talents.Points, wantPoints) {
		t.Fatalf("points = %v, want %v", d.Talents.Points, wantPoints)
	}
	if !d.HasLevel || d.Level != 24 {
		t.Fatalf("level = %v hasLevel=%v, want 24/true", d.Level, d.HasLevel)
	}
	if !d.HasWho || d.CharacterName != "Bow" || d.Realm != "Skyborne" {
		t.Fatalf("who = %q/%q hasWho=%v, want Bow/Skyborne/true", d.CharacterName, d.Realm, d.HasWho)
	}
}

func TestDecodeOmitsAbsentSections(t *testing.T) {
	d, ok := Decode("FS1:1.60.1.69893:warrior:orc:0/0/0:")
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if len(d.Gear) != 0 {
		t.Fatalf("gear = %v, want empty", d.Gear)
	}
	if len(d.Enchants) != 0 {
		t.Fatalf("enchants = %v, want empty", d.Enchants)
	}
	if len(d.Professions) != 0 {
		t.Fatalf("professions = %v, want empty", d.Professions)
	}
	if len(d.Bags) != 0 {
		t.Fatalf("bags = %v, want empty", d.Bags)
	}
	if d.HasLevel || d.HasWho {
		t.Fatalf("hasLevel=%v hasWho=%v, want both false", d.HasLevel, d.HasWho)
	}
}

func TestDecodeReadsEnchantsProfessionsAndBags(t *testing.T) {
	// chest carries an enchant (41); wrist carries an enchant plus a suffix (724:1234);
	// head carries no enchant at all - the three real shapes Codec.lua's encodeItemParts
	// writes.
	export := "FS1:1.60.1.70009:hunter:troll:503200000/0/0:head=12640,chest=11726:41,wrist=1234:724:1234" +
		"|professions=skinning,leatherworking,first-aid|bags=13510,13444:0,232433"
	d, ok := Decode(export)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	wantGear := map[string]int{"head": 12640, "chest": 11726, "wrist": 1234}
	if !reflect.DeepEqual(d.Gear, wantGear) {
		t.Fatalf("gear = %v, want %v", d.Gear, wantGear)
	}
	wantEnchants := map[string]int{"chest": 41, "wrist": 724}
	if !reflect.DeepEqual(d.Enchants, wantEnchants) {
		t.Fatalf("enchants = %v, want %v", d.Enchants, wantEnchants)
	}
	wantProfessions := []string{"skinning", "leatherworking", "first-aid"}
	if !reflect.DeepEqual(d.Professions, wantProfessions) {
		t.Fatalf("professions = %v, want %v", d.Professions, wantProfessions)
	}
	wantBags := []int{13510, 13444, 232433}
	if !reflect.DeepEqual(d.Bags, wantBags) {
		t.Fatalf("bags = %v, want %v", d.Bags, wantBags)
	}
}

func TestDecodeGearWithZeroEnchantRecordsNoEnchant(t *testing.T) {
	d, ok := Decode("FS1:1.60.1.69893:warrior:orc:0/0/0:chest=11726:0")
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if d.Gear["chest"] != 11726 {
		t.Fatalf("gear[chest] = %v, want 11726", d.Gear["chest"])
	}
	if _, present := d.Enchants["chest"]; present {
		t.Fatalf("enchants[chest] present = true, want absent for an explicit enchant=0")
	}
}

func TestDecodeMalformedProfessionsAndBagsEntriesAreSkipped(t *testing.T) {
	d, ok := Decode("FS1:1.60.1.69893:warrior:orc:0/0/0:|professions=,skinning,|bags=notanumber,13510")
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if want := []string{"skinning"}; !reflect.DeepEqual(d.Professions, want) {
		t.Fatalf("professions = %v, want %v", d.Professions, want)
	}
	if want := []int{13510}; !reflect.DeepEqual(d.Bags, want) {
		t.Fatalf("bags = %v, want %v", d.Bags, want)
	}
}

func TestDecodeMalformedSectionsAreSkippedNotFatal(t *testing.T) {
	// A bad level (out of 1..60) and a bad gear entry (non-numeric item id) are each
	// simply absent from the result; the head still decodes.
	d, ok := Decode("FS1:1.60.1.69893:warrior:orc:0/0/0:head=notanumber,chest=11726|level=999|who=Bad")
	if !ok {
		t.Fatal("ok = false, want true (only the head's own grammar can fail the decode)")
	}
	if _, present := d.Gear["head"]; present {
		t.Fatalf("gear[head] = %v, want the malformed entry skipped", d.Gear["head"])
	}
	if d.Gear["chest"] != 11726 {
		t.Fatalf("gear[chest] = %v, want 11726", d.Gear["chest"])
	}
	if d.HasLevel {
		t.Fatal("hasLevel = true, want the out-of-range level treated as absent")
	}
	// "who=Bad" with no colon: realm is simply empty, name is the whole field.
	if !d.HasWho || d.CharacterName != "Bad" || d.Realm != "" {
		t.Fatalf("who = %q/%q hasWho=%v", d.CharacterName, d.Realm, d.HasWho)
	}
}

func TestDecodeRefusesAHeadThatIsNotReadable(t *testing.T) {
	cases := map[string]string{
		"wrong prefix":        "FS2:1.60.1.69893:warrior:orc:0/0/0:",
		"too few head fields": "FS1:1.60.1.69893:warrior:orc",
		"wrong tree count":    "FS1:1.60.1.69893:warrior:orc:0/0:",
		"unreadable digit":    "FS1:1.60.1.69893:warrior:orc:0/!/0:",
		"empty string":        "",
	}
	for name, export := range cases {
		t.Run(name, func(t *testing.T) {
			if _, ok := Decode(export); ok {
				t.Fatalf("%s: ok = true, want false", name)
			}
		})
	}
}

// inventorySlots is Export.INVENTORY_SLOTS (addon/ForeverSixty/Export.lua) and
// bnetbuild.SLOTS (api/internal/bnetbuild/slots.go) — the site's slot vocabulary an FS1
// gear entry's name is always one of, copied here verbatim rather than re-derived, so a
// drift in any of the three lists shows up as a failure in this test.
var inventorySlots = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
	"finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
}

func TestDecodeGearCopiesEveryInventorySlotNameVerbatim(t *testing.T) {
	// Decode copies whatever slot names the export carries; it does not validate them
	// against a fixed slot list (that is the addon's and the planner's own decoders'
	// job before they ever write or trust a string) — this test only proves every real
	// slot name round-trips, not that Decode enforces the list.
	var entries []string
	want := map[string]int{}
	for i, slot := range inventorySlots {
		itemID := 10000 + i
		entries = append(entries, slot+"="+strconv.Itoa(itemID))
		want[slot] = itemID
	}
	export := "FS1:1.60.1.69893:warrior:orc:0/0/0:" + strings.Join(entries, ",")
	d, ok := Decode(export)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if !reflect.DeepEqual(d.Gear, want) {
		t.Fatalf("gear = %v, want %v", d.Gear, want)
	}
}
