package request

import (
	"errors"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestCooldownIDGrammar(t *testing.T) {
	table := testConsumables(t)
	cases := []struct {
		name     string
		id       string
		wantKind string // "spell" or "item"
		wantRaw  int32
		wantErr  bool
	}{
		{name: "a spell", id: "spell:1719", wantKind: "spell", wantRaw: 1719},
		{name: "an item", id: "item:13452", wantKind: "item", wantRaw: 13452},
		{name: "a consumable by name", id: "elixir_of_the_mongoose", wantKind: "item", wantRaw: 13452},
		{name: "a spell with no number", id: "spell:", wantErr: true},
		{name: "a spell that is not a number", id: "spell:mongoose", wantErr: true},
		{name: "a bare word the table does not know", id: "elixir_of_nothing", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := cooldownAction(c.id, table)
			if c.wantErr {
				if err == nil {
					t.Fatalf("%q was accepted as %+v", c.id, got)
				}
				if !errors.Is(err, ErrUnknownCooldown) {
					t.Errorf("error %v is not ErrUnknownCooldown", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			switch c.wantKind {
			case "spell":
				if got.GetSpellId() != c.wantRaw {
					t.Errorf("spell_id = %d, want %d", got.GetSpellId(), c.wantRaw)
				}
			case "item":
				if got.GetItemId() != c.wantRaw {
					t.Errorf("item_id = %d, want %d", got.GetItemId(), c.wantRaw)
				}
			}
		})
	}
}

// The excerpt's ids are whatever the ten rows carry; if 13452 is not
// "Elixir of the Mongoose" in testdata/simconsumes.excerpt.json, read
// the file and use a row that is there. The grammar is what is under
// test, not the particular item.

func TestBuildCarriesCooldownTimings(t *testing.T) {
	req := fury()
	req.Character.Cooldowns = []api.CooldownSpec{
		{ID: "spell:1719", AtSec: []float64{0, 90}},
		{ID: "spell:20572"},
	}
	got, err := BuildWith(req, Options{Consumables: testConsumables(t)})
	if err != nil {
		t.Fatal(err)
	}
	cds := got.Raid.Parties[0].Players[0].GetCooldowns().GetCooldowns()
	if len(cds) != 2 {
		t.Fatalf("cooldowns = %d, want 2", len(cds))
	}
	if cds[0].Id.GetSpellId() != 1719 || len(cds[0].Timings) != 2 || cds[0].Timings[1] != 90 {
		t.Errorf("first cooldown = %+v", cds[0])
	}
	// No timings is "on cooldown", which the engine spells as an empty
	// list rather than as an absent entry: the entry is what makes the
	// cooldown eligible at all.
	if cds[1].Id.GetSpellId() != 20572 || len(cds[1].Timings) != 0 {
		t.Errorf("second cooldown = %+v", cds[1])
	}
}

func TestBuildWithNoCooldownsLeavesTheMessageAbsent(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if cd := got.Raid.Parties[0].Players[0].GetCooldowns(); cd != nil && len(cd.Cooldowns) != 0 {
		t.Errorf("a request with no cooldowns produced %+v", cd)
	}
}

func TestACooldownNamedByItemNeedsTheTable(t *testing.T) {
	req := fury()
	req.Character.Cooldowns = []api.CooldownSpec{{ID: "elixir_of_the_mongoose"}}
	_, err := Build(req) // Build carries no consumable table
	if err == nil {
		t.Fatal("a consumable cooldown resolved without a table")
	}
	if !strings.Contains(err.Error(), "no consumable table") {
		t.Errorf("error %q does not say the table is missing", err)
	}
}

// Two items whose names normalise to one key cannot be told apart, so
// the reverse lookup refuses rather than picking one.
func TestConsumablesItemIDRefusesAnAmbiguousKey(t *testing.T) {
	table := &Consumables{names: map[int64]string{1: "twin", 2: "twin", 3: "alone"}}
	if _, ok := table.ItemID("twin"); ok {
		t.Error("an ambiguous key resolved")
	}
	if id, ok := table.ItemID("alone"); !ok || id != 3 {
		t.Errorf("ItemID(alone) = %d, %v", id, ok)
	}
	if _, ok := table.ItemID("absent"); ok {
		t.Error("an absent key resolved")
	}
	if _, ok := (*Consumables)(nil).ItemID("alone"); ok {
		t.Error("a nil table resolved a key")
	}
}
