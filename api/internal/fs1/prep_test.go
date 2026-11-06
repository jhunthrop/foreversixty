package fs1

import (
	"reflect"
	"testing"
)

const prepHead = "FS1:1.60.1.70790:warrior:human:0/0/0:head=16866|level=60"

func TestDecodeReadsTheRaidPrepSections(t *testing.T) {
	d, ok := Decode(prepHead + "|bags=16309|bank=12344|school=52:10:0:15:5|supplies=13510:1,20452:6,20452:4|locks=13704,12382|quests=6502:1,6602:0|owned=16309:1|who=Tank:Realm")
	if !ok {
		t.Fatal("ok = false")
	}
	if d.Resistances == nil || *d.Resistances != (Resistances{Fire: 52, Nature: 10, Frost: 0, Shadow: 15, Arcane: 5}) {
		t.Fatalf("resistances = %+v", d.Resistances)
	}
	if !d.HasSupplies || !reflect.DeepEqual(d.Supplies, map[int]int{13510: 1, 20452: 10}) {
		t.Fatalf("supplies = %v has=%v", d.Supplies, d.HasSupplies)
	}
	if !d.HasKeyring || !reflect.DeepEqual(d.Keyring, []int{13704, 12382}) {
		t.Fatalf("keyring = %v has=%v", d.Keyring, d.HasKeyring)
	}
	if !d.HasQuests || !reflect.DeepEqual(d.Quests, map[int]bool{6502: true, 6602: false}) {
		t.Fatalf("quests = %v has=%v", d.Quests, d.HasQuests)
	}
	if !d.HasOwned || !reflect.DeepEqual(d.Owned, map[int]bool{16309: true}) {
		t.Fatalf("owned = %v has=%v", d.Owned, d.HasOwned)
	}
	if !d.HasBags || !reflect.DeepEqual(d.Bank, []int{12344}) {
		t.Fatalf("bags=%v bank=%v", d.HasBags, d.Bank)
	}
	if !d.HasWho || d.CharacterName != "Tank" {
		t.Fatalf("who after the prep sections = %q", d.CharacterName)
	}
}

func TestDecodeOldExportHasNoRaidPrep(t *testing.T) {
	d, ok := Decode(prepHead + "|who=Old:Realm")
	if !ok {
		t.Fatal("ok = false")
	}
	if d.Resistances != nil || d.HasSupplies || d.HasKeyring || d.HasQuests || d.HasOwned || d.HasBags {
		t.Fatalf("old export reads prep: %+v", d)
	}
	if d.Supplies == nil || d.Keyring == nil || d.Quests == nil || d.Bank == nil {
		t.Fatal("absent prep collections must be empty, not nil")
	}
}

func TestDecodeEmptyRaidPrepSectionsAreRead(t *testing.T) {
	d, ok := Decode(prepHead + "|supplies=|locks=|quests=")
	if !ok || !d.HasSupplies || !d.HasKeyring || !d.HasQuests {
		t.Fatalf("empty sections must still read as present: %+v", d)
	}
	if len(d.Supplies) != 0 || len(d.Keyring) != 0 || len(d.Quests) != 0 {
		t.Fatalf("empty sections carry entries: %+v", d)
	}
}

func TestDecodeDropsMalformedRaidPrep(t *testing.T) {
	d, ok := Decode(prepHead + "|school=1:2:3|supplies=x:1,13510:,13510:2,0:4|locks=a,13704,-1|quests=6502:2,6602:1,x:1")
	if !ok {
		t.Fatal("a malformed prep section must not fail the decode")
	}
	if d.Resistances != nil {
		t.Fatalf("a four-school section read as %+v", d.Resistances)
	}
	if !reflect.DeepEqual(d.Supplies, map[int]int{13510: 2}) {
		t.Fatalf("supplies = %v", d.Supplies)
	}
	if !reflect.DeepEqual(d.Keyring, []int{13704}) {
		t.Fatalf("keyring = %v", d.Keyring)
	}
	if !reflect.DeepEqual(d.Quests, map[int]bool{6602: true}) {
		t.Fatalf("quests = %v", d.Quests)
	}
}
