package guilds

import (
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
)

// furyHeadPick is the Fury band's head pick at band, read from the real data.
func furyHeadPick(t *testing.T, s *Store, band int) int {
	t.Helper()
	b, ok := s.loadBandAt("warrior", "Fury", "alliance", band)
	if !ok {
		t.Fatalf("no Fury band %d in the data", band)
	}
	pick := b.BySlot()["head"].ItemID
	if pick <= 0 {
		t.Fatalf("Fury band %d has no head pick", band)
	}
	return pick
}

func furyRow(key, name string, level int, gear map[string]int) RosterRow {
	return RosterRow{CharacterKey: key, Name: name, Verified: true, className: "warrior", specStr: "Fury",
		faction: "alliance", Level: &level, gear: gear}
}

// A member needs an item when their spec's band at their level names it and they do not wear
// it; one who wears it has it; one whose gear is unknown is counted, never listed.
func TestGuildItemNeeds(t *testing.T) {
	s := realTreesStore(t)
	helm := furyHeadPick(t, s, 60)
	roster := []RosterRow{
		furyRow("us/pvp/needer", "Needer", 60, shirtOnly()),
		furyRow("us/pvp/wearer", "Wearer", 60, map[string]int{"head": helm}),
		furyRow("us/pvp/unknown", "Unknown", 60, nil),
		furyRow("us/pvp/lowbie", "Lowbie", 24, shirtOnly()),
		{CharacterKey: "us/pvp/unverified", Name: "Unverified", className: "warrior", specStr: "Fury", faction: "alliance", gear: shirtOnly()},
	}
	got := s.guildItemNeeds(roster, helm)
	if len(got.Needs) != 1 || got.Needs[0].Name != "Needer" || got.Needs[0].Band != 60 || got.Needs[0].Slot != "head" {
		t.Fatalf("needs = %+v", got.Needs)
	}
	if len(got.Has) != 1 || got.Has[0].Name != "Wearer" {
		t.Fatalf("has = %+v", got.Has)
	}
	if got.UnknownGear != 1 || got.ItemID != helm {
		t.Fatalf("got %+v", got)
	}
}

// The band follows the member's level: a leveler's list is their band's, not the raid's.
func TestGuildItemNeedsReadsEachMembersBand(t *testing.T) {
	s := realTreesStore(t)
	low := furyHeadPick(t, s, 40)
	if low == furyHeadPick(t, s, 60) {
		t.Skip("the data's band 40 and 60 head picks are the same item")
	}
	got := s.guildItemNeeds([]RosterRow{
		furyRow("us/pvp/forty", "Forty", 44, shirtOnly()),
		furyRow("us/pvp/sixty", "Sixty", 60, shirtOnly()),
	}, low)
	if len(got.Needs) != 1 || got.Needs[0].Name != "Forty" || got.Needs[0].Band != 40 {
		t.Fatalf("needs = %+v", got.Needs)
	}
}

func TestAppContractItemNeeds(t *testing.T) {
	holder := ItemHolder{CharacterKey: "us/pvp/kraggor", Name: "Kraggor", Class: "Warrior", Spec: "Fury"}
	body := map[string]any{"ok": true, "error": nil, "request_id": "x", "data": ItemNeeds{
		ItemID: 18205, Needs: []ItemNeed{{ItemHolder: holder, Level: 60, Band: 60, Slot: "neck", Pick: true}},
		Has: []ItemHolder{holder}, UnknownGear: 3,
	}}
	contract.AssertValue(t, "guild.item-needs", body, contract.Options{})
}

func TestItemNeedsIsForMembers(t *testing.T) {
	w := newRaidWorld(t, true)
	stranger := seedUser(t, w.h.pool, "stranger@example.com")
	w.as(stranger)
	if res := w.h.do(http.MethodGet, w.path("/items/16963/needs"), ""); res.StatusCode != http.StatusNotFound {
		t.Fatalf("stranger = %d", res.StatusCode)
	}
	w.as(w.member)
	if res := w.h.do(http.MethodGet, w.path("/items/0/needs"), ""); res.StatusCode != http.StatusNotFound {
		t.Fatalf("bad item = %d", res.StatusCode)
	}
	res := w.h.do(http.MethodGet, w.path("/items/16963/needs"), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("member = %d: %s", res.StatusCode, readBody(res))
	}
	var got ItemNeeds
	w.h.data(res, &got)
	if got.ItemID != 16963 || got.Needs == nil || got.Has == nil {
		t.Fatalf("got %+v", got)
	}
}
