package fs1

import (
	"reflect"
	"testing"
)

const dropsHead = "FS1:1.60.1.70790:warrior:human:0/0/0:head=16900"

func TestDecodeDropsReadsCorpsesItemsAndWinners(t *testing.T) {
	d, ok := Decode(dropsHead + "|drops=1762480000:249:1084:master:Creature-0-1-249-7-10184-0A:16900x1@Thr%C3%A2ll,18205x2@A+B;1762480100:249:0::x:15410x1|who=Tester:Ashbringer")
	if !ok || !d.HasDrops {
		t.Fatalf("decode = %v, has drops %v", ok, d.HasDrops)
	}
	want := []DropCorpse{
		{At: 1762480000, InstanceID: 249, EncounterID: 1084, Method: "master", Source: "Creature-0-1-249-7-10184-0A",
			Items: []DropItem{{ItemID: 16900, Count: 1, Winners: []string{"Thrâll"}}, {ItemID: 18205, Count: 2, Winners: []string{"A", "B"}}}},
		{At: 1762480100, InstanceID: 249, Source: "x", Items: []DropItem{{ItemID: 15410, Count: 1, Winners: []string{}}}},
	}
	if !reflect.DeepEqual(d.Drops, want) {
		t.Fatalf("drops = %+v", d.Drops)
	}
	if d.CharacterName != "Tester" {
		t.Fatalf("the section after drops= is still read: %q", d.CharacterName)
	}
}

func TestDecodeDropsDropsMalformedEntries(t *testing.T) {
	d, _ := Decode(dropsHead + "|drops=bad;0:1:2::x:1x1;1:2:3:Master::1x1;1:2:3::g.u:1x1;5:1:1:::zz,0x1,1x0,1x41,7x1")
	if len(d.Drops) != 1 || len(d.Drops[0].Items) != 1 || d.Drops[0].Items[0].ItemID != 7 {
		t.Fatalf("drops = %+v", d.Drops)
	}
}

func TestDecodeWithoutDropsIsNotTracked(t *testing.T) {
	d, _ := Decode(dropsHead)
	if d.HasDrops || d.Drops == nil || len(d.Drops) != 0 {
		t.Fatalf("drops = %+v, has %v", d.Drops, d.HasDrops)
	}
}
