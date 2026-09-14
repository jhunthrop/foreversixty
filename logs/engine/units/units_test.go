// logs/engine/units/units_test.go
package units

import (
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

var at = time.Date(2026, 9, 26, 20, 10, 0, 0, time.UTC)

func TestParseGUID(t *testing.T) {
	for _, tc := range []struct {
		in       string
		kind     Kind
		serverID int64
		npcID    int64
	}{
		{"Player-4184-000000A1", KindPlayer, 4184, 0},
		{"Creature-0-2085-2284-7855-169753-0000AA0001", KindCreature, 2085, 169753},
		{"Pet-0-2085-2284-7855-165189-01000000B1", KindPet, 2085, 165189},
		{"Vehicle-0-2085-2284-7855-170000-0000AA0002", KindVehicle, 2085, 170000},
		{"GameObject-0-2085-2284-7855-335621-0000AA0003", KindGameObject, 2085, 335621},
		{"0000000000000000", KindNone, 0, 0},
		{"", KindNone, 0, 0},
		{"Something-Else", KindUnknown, 0, 0},
	} {
		t.Run(tc.in, func(t *testing.T) {
			g := Parse(tc.in)
			if g.Kind != tc.kind || g.ServerID != tc.serverID || g.NPCID != tc.npcID {
				t.Fatalf("got kind=%s server=%d npc=%d, want %s %d %d",
					g.Kind, g.ServerID, g.NPCID, tc.kind, tc.serverID, tc.npcID)
			}
		})
	}
}

func TestFlagHelpers(t *testing.T) {
	if !Hostile(0xa48) || Friendly(0xa48) {
		t.Error("0xa48 is a hostile NPC")
	}
	if !Friendly(0x511) || Hostile(0x511) {
		t.Error("0x511 is a friendly player")
	}
	if 0x1114&FlagTypePet == 0 {
		t.Error("0x1114 is a pet")
	}
}

func TestRegistryNamesUnitsAndTracksTimes(t *testing.T) {
	r := NewRegistry(Options{})
	r.Observe(event.Event{
		Time:   at,
		Kind:   event.Damage,
		Source: event.Unit{GUID: "Player-4184-000000A1", Name: "Baelgrim-Nightslayer", Flags: 0x511},
		Dest:   event.Unit{GUID: "Creature-0-2085-2284-7855-169753-0000AA0001", Name: "Hollow Sentinel", Flags: 0xa48},
	})
	r.Observe(event.Event{
		Time:   at.Add(5 * time.Second),
		Kind:   event.Damage,
		Source: event.Unit{GUID: "Player-4184-000000A1", Name: "Baelgrim-Nightslayer", Flags: 0x511},
		Dest:   event.Unit{GUID: "Creature-0-2085-2284-7855-169753-0000AA0001", Name: "Hollow Sentinel", Flags: 0xa48},
	})
	u, ok := r.Get("Player-4184-000000A1")
	if !ok || !u.IsPlayer() || u.Name != "Baelgrim-Nightslayer" {
		t.Fatalf("player = %+v ok=%v", u, ok)
	}
	if !u.FirstSeen.Equal(at) || !u.LastSeen.Equal(at.Add(5*time.Second)) {
		t.Errorf("times = %s to %s", u.FirstSeen, u.LastSeen)
	}
	if got := r.Name("Creature-0-2085-2284-7855-169753-0000AA0001"); got != "Hollow Sentinel" {
		t.Errorf("name = %q", got)
	}
	if got := r.Name("nobody"); got != "nobody" {
		t.Errorf("an unknown GUID must return itself, got %q", got)
	}
	if len(r.Players()) != 1 || len(r.All()) != 2 {
		t.Errorf("players=%d all=%d", len(r.Players()), len(r.All()))
	}
}

func TestPetOwnerFromASummonAndFromTheAdvancedBlock(t *testing.T) {
	r := NewRegistry(Options{})
	owner := "Player-4184-000000A4"
	pet := "Pet-0-2085-2284-7855-165189-01000000B1"
	r.Observe(event.Event{
		Time: at, Kind: event.Summon,
		Source: event.Unit{GUID: owner, Name: "Thalgrit-Nightslayer", Flags: 0x511},
		Dest:   event.Unit{GUID: pet, Name: "Ashfang", Flags: 0x1114},
	})
	if got := r.Owner(pet); got != owner {
		t.Fatalf("owner after summon = %q, want %q", got, owner)
	}

	r2 := NewRegistry(Options{})
	r2.Observe(event.Event{
		Time: at, Kind: event.Damage,
		Source: event.Unit{GUID: pet, Name: "Ashfang", Flags: 0x1114},
		Dest:   event.Unit{GUID: "Creature-0-2085-2284-7855-169753-0000AA0001", Name: "Hollow Sentinel", Flags: 0xa48},
		Adv:    event.Advanced{OK: true, InfoGUID: pet, OwnerGUID: owner},
	})
	if got := r2.Owner(pet); got != owner {
		t.Fatalf("owner from the advanced block = %q, want %q", got, owner)
	}
	if got := r2.Owner(owner); got != owner {
		t.Errorf("a unit with no owner returns itself, got %q", got)
	}
}

func TestOwnerStopsOnACycle(t *testing.T) {
	r := NewRegistry(Options{})
	a, b := "Pet-0-1-1-1-1-1-A", "Pet-0-1-1-1-1-1-B"
	r.Observe(event.Event{Time: at, Kind: event.Summon,
		Source: event.Unit{GUID: a}, Dest: event.Unit{GUID: b}})
	r.Observe(event.Event{Time: at, Kind: event.Summon,
		Source: event.Unit{GUID: b}, Dest: event.Unit{GUID: a}})
	if got := r.Owner(a); got != a && got != b {
		t.Fatalf("a cycle must terminate, got %q", got)
	}
}

func TestClassFromCombatantInfoBeatsInference(t *testing.T) {
	guid := "Player-4184-000000A1"
	r := NewRegistry(Options{
		ClassBySpec:  RetailSpecClass,
		ClassBySpell: map[int64]string{116: "Mage"},
	})
	// An inference first.
	r.Observe(event.Event{Time: at, Kind: event.CastSuccess,
		Source: event.Unit{GUID: guid, Name: "Baelgrim-Nightslayer", Flags: 0x511},
		Spell:  event.Spell{ID: 116, Name: "Frostbolt"}})
	u, _ := r.Get(guid)
	if u.Class != "Mage" || u.ClassSource != "inferred" {
		t.Fatalf("inferred class = %q from %q", u.Class, u.ClassSource)
	}
	// Then the authoritative answer.
	r.Observe(event.Event{Time: at, Kind: event.CombatantInfo,
		Combatant: &event.Combatant{GUID: guid, SpecID: 73, ItemLevel: 183}})
	u, _ = r.Get(guid)
	if u.Class != "Warrior" || u.ClassSource != "combatant_info" || u.ItemLevel != 183 {
		t.Fatalf("after COMBATANT_INFO = %+v", u)
	}
	// A later inference must not undo it.
	r.Observe(event.Event{Time: at, Kind: event.CastSuccess,
		Source: event.Unit{GUID: guid, Flags: 0x511},
		Spell:  event.Spell{ID: 116}})
	u, _ = r.Get(guid)
	if u.Class != "Warrior" {
		t.Fatalf("inference overwrote COMBATANT_INFO: %+v", u)
	}
}

func TestNoClassTableMeansNoClaimedClass(t *testing.T) {
	r := NewRegistry(Options{})
	r.Observe(event.Event{Time: at, Kind: event.CastSuccess,
		Source: event.Unit{GUID: "Player-4184-000000A1", Flags: 0x511},
		Spell:  event.Spell{ID: 116}})
	u, _ := r.Get("Player-4184-000000A1")
	if u.Class != "" || u.ClassSource != "" {
		t.Fatalf("class was guessed without a table: %+v", u)
	}
}

func TestStateRoundTrip(t *testing.T) {
	r := NewRegistry(Options{ClassBySpec: RetailSpecClass})
	r.Observe(event.Event{Time: at, Kind: event.Summon,
		Source: event.Unit{GUID: "Player-4184-000000A4", Name: "Thalgrit-Nightslayer", Flags: 0x511},
		Dest:   event.Unit{GUID: "Pet-0-2085-2284-7855-165189-01000000B1", Name: "Ashfang", Flags: 0x1114}})
	revived := RestoreRegistry(Options{ClassBySpec: RetailSpecClass}, r.State())
	if got := revived.Owner("Pet-0-2085-2284-7855-165189-01000000B1"); got != "Player-4184-000000A4" {
		t.Fatalf("owner lost across restore: %q", got)
	}
	if len(revived.All()) != len(r.All()) {
		t.Fatalf("unit count %d != %d", len(revived.All()), len(r.All()))
	}
}

func TestAllIsSortedForDeterminism(t *testing.T) {
	r := NewRegistry(Options{})
	for _, g := range []string{"Player-1-C", "Player-1-A", "Player-1-B"} {
		r.Observe(event.Event{Time: at, Source: event.Unit{GUID: g, Flags: 0x511}})
	}
	all := r.All()
	for i := 1; i < len(all); i++ {
		if all[i-1].GUID >= all[i].GUID {
			t.Fatalf("All is not sorted: %q then %q", all[i-1].GUID, all[i].GUID)
		}
	}
}
