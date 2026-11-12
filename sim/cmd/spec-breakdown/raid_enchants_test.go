package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// The raid pass ranks and measures every published set wearing its
// permanent enchants (leveling-bis enchants.go), and the entry's slots
// carry them as enchant_id. A breakdown that read the items alone simmed a
// raid caster 8 to 10% under its own published set_dps (2026-11-11), so the
// request wears what the entry wears.
func TestRaidEntryWearsItsPublishedEnchants(t *testing.T) {
	o := options{repoRoot: repoRoot, spec: "mage-fire", preset: "raid", band: 60, faction: "alliance", iterations: 1, seed: 1}
	entry, err := loadEntry(o)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int{}
	for _, s := range entry.Slots {
		if s.EnchantID != 0 {
			want[s.Slot] = s.EnchantID
		}
	}
	if len(want) == 0 {
		t.Fatal("the published mage-fire raid entry carries no enchant_id; the test needs one")
	}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	assertEnchants(t, req.Character.Gear, want)
}

// A bare entry carries no enchants, and the breakdown adds none.
func TestBareEntryWearsNoEnchant(t *testing.T) {
	o := options{repoRoot: repoRoot, spec: "mage-fire", preset: "bare", band: 60, faction: "alliance", iterations: 1, seed: 1}
	req, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	assertEnchants(t, req.Character.Gear, map[string]int{})
}

func assertEnchants(t *testing.T, gear []api.GearSlot, want map[string]int) {
	t.Helper()
	got := map[string]int{}
	for _, g := range gear {
		if g.Enchant != 0 {
			got[g.Slot] = g.Enchant
		}
	}
	for slot, id := range want {
		if got[slot] != id {
			t.Errorf("%s: request enchant %d, published entry %d", slot, got[slot], id)
		}
	}
	for slot, id := range got {
		if want[slot] == 0 {
			t.Errorf("%s: request wears enchant %d the entry does not carry", slot, id)
		}
	}
}
