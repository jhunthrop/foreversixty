package main

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/request"
)

const enchantsFixture = `{
  "profiles": {
    "caster": {
      "wrist": {"id": 7665, "why": "+16 spell power"},
      "two_hand": {"id": 7662, "why": "Spellblasting"}
    },
    "melee": {
      "hands": {"id": 8206, "why": "+10 Agility"},
      "main_hand": {"id": 1900, "why": "Crusader"},
      "off_hand": {"id": 1900, "why": "Crusader"},
      "two_hand": {"id": 1900, "why": "Crusader"},
      "shield": {"id": 1886, "why": "+9 Stamina"}
    }
  },
  "specs": {
    "mage-fire": {"profile": "caster"},
    "druid-feral": {"profile": "melee", "slots": {"two_hand": null}},
    "bad-profile": {"profile": "nope"},
    "bad-key": {"profile": "caster", "slots": {"elbow": {"id": 7665}}},
    "unknown-id": {"profile": "caster", "slots": {"wrist": {"id": 99999999}}},
    "two-hand-on-one-hand": {"profile": "caster", "slots": {"main_hand": {"id": 7662}}}
  }
}`

func TestRaidEnchantsLoadAgainstTheBuildsTable(t *testing.T) {
	plan, err := parseRaidEnchants([]byte(enchantsFixture), "mage-fire", "mage", simdb.EnchantFits)
	if err != nil {
		t.Fatal(err)
	}
	want := enchantPlan{"wrist": 7665, "two_hand": 7662}
	if len(plan) != len(want) || plan["wrist"] != 7665 || plan["two_hand"] != 7662 {
		t.Errorf("plan = %v, want %v", plan, want)
	}
}

func TestASpecOverrideOfNullDropsTheProfilesSlot(t *testing.T) {
	plan, err := parseRaidEnchants([]byte(enchantsFixture), "druid-feral", "druid", simdb.EnchantFits)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := plan["two_hand"]; ok || plan["main_hand"] != 1900 {
		t.Errorf("plan = %v, want Crusader on one-handers and nothing on a two-hander", plan)
	}
}

func TestABadRaidEnchantPlanFailsLoudly(t *testing.T) {
	for spec, want := range map[string]string{
		"missing-spec":         "names no plan",
		"bad-profile":          "unknown profile",
		"bad-key":              "not a slot key",
		"unknown-id":           simdb.ErrUnknownEnchant.Error(),
		"two-hand-on-one-hand": simdb.ErrEnchantDoesNotFit.Error(),
	} {
		_, err := parseRaidEnchants([]byte(enchantsFixture), spec, "mage", simdb.EnchantFits)
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%s: err = %v, want one containing %q", spec, err, want)
		}
	}
}

// Every spec the ranker ranks has a plan, and every enchant in it is in the
// active build's enchants.json and fits its slot.
func TestTheCuratedRaidEnchantsLoadForEverySpec(t *testing.T) {
	specs, err := loadAllSpecs(integrationRepoRoot)
	if err != nil {
		t.Fatal(err)
	}
	for _, spec := range specs {
		plan, err := loadRaidEnchants(integrationRepoRoot, spec)
		if err != nil {
			t.Errorf("%s: %v", spec.Spec, err)
			continue
		}
		if len(plan) == 0 {
			t.Errorf("%s: an empty raid enchant plan", spec.Spec)
		}
	}
}

var enchantTestItems = map[int]simdb.Item{
	1: {ID: 1, HandType: simdb.HandTwo, WeaponType: "two_handed_sword"},
	2: {ID: 2, HandType: simdb.HandOne, WeaponType: "sword"},
	3: {ID: 3, WeaponType: "shield"},
	4: {ID: 4},
	5: {ID: 5, HandType: simdb.HandOne, WeaponType: "dagger"},
}

func lookupEnchantTestItem(id int) (simdb.Item, bool) {
	item, ok := enchantTestItems[id]
	return item, ok
}

func acceptAnyEnchant(int, string, simdb.Item, string) error { return nil }

var meleePlan = enchantPlan{"hands": 8206, "main_hand": 1900, "off_hand": 1901, "two_hand": 1902, "shield": 1886}

func enchantsOf(gear []api.GearSlot) map[string]int {
	out := map[string]int{}
	for _, g := range gear {
		out[g.Slot] = g.Enchant
	}
	return out
}

func TestWearPicksTheWeaponEnchantByTheWeaponsShape(t *testing.T) {
	cases := []struct {
		name string
		gear []api.GearSlot
		want map[string]int
	}{
		{"two-hander", []api.GearSlot{{Slot: "main_hand", ItemID: 1}}, map[string]int{"main_hand": 1902}},
		{"dual wield", []api.GearSlot{{Slot: "main_hand", ItemID: 2}, {Slot: "off_hand", ItemID: 5}}, map[string]int{"main_hand": 1900, "off_hand": 1901}},
		{"one-hander and shield", []api.GearSlot{{Slot: "main_hand", ItemID: 2}, {Slot: "off_hand", ItemID: 3}}, map[string]int{"main_hand": 1900, "off_hand": 1886}},
		{"held in off hand", []api.GearSlot{{Slot: "off_hand", ItemID: 4}}, map[string]int{"off_hand": 0}},
		{"armor and a ring", []api.GearSlot{{Slot: "hands", ItemID: 4}, {Slot: "finger1", ItemID: 4}}, map[string]int{"hands": 8206, "finger1": 0}},
		{"an item the build lacks", []api.GearSlot{{Slot: "hands", ItemID: 77}}, map[string]int{"hands": 0}},
		{"an enchant already worn", []api.GearSlot{{Slot: "hands", ItemID: 4, Enchant: 904}}, map[string]int{"hands": 904}},
	}
	for _, c := range cases {
		got := enchantsOf(meleePlan.wearWith("rogue", c.gear, lookupEnchantTestItem, acceptAnyEnchant))
		for slot, want := range c.want {
			if got[slot] != want {
				t.Errorf("%s: %s enchant = %d, want %d", c.name, slot, got[slot], want)
			}
		}
	}
}

func TestWearLeavesItsInputAlone(t *testing.T) {
	gear := []api.GearSlot{{Slot: "hands", ItemID: 4}}
	before := slices.Clone(gear)
	meleePlan.wearWith("rogue", gear, lookupEnchantTestItem, acceptAnyEnchant)
	if !slices.Equal(gear, before) {
		t.Errorf("wear changed its input: %v", gear)
	}
}

func TestWearPanicsOnAnEnchantThatDoesNotFit(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("wear put on an enchant the fit check refused")
		}
	}()
	refuse := func(int, string, simdb.Item, string) error { return errors.New("no") }
	meleePlan.wearWith("rogue", []api.GearSlot{{Slot: "hands", ItemID: 4}}, lookupEnchantTestItem, refuse)
}

// Items the active build's database carries, for the tests that go through
// the real lookup.
const (
	realGloves = 14146
	realRing   = 13001
)

func TestOnlyTheRaidPassWearsEnchants(t *testing.T) {
	gear := []api.GearSlot{{Slot: "hands", ItemID: realGloves}}
	plan := enchantPlan{"hands": 8206}
	bare := specInfo{Spec: "rogue-combat", ClassSlug: "rogue", RaidEnchants: plan}
	raid := bare
	raid.Applied = &request.ResolvedPreset{}

	ch := api.CharacterSpec{Level: 60, Gear: gear}
	if got := withRaidEnchants(bare, ch).Gear[0].Enchant; got != 0 {
		t.Errorf("bare pass wears enchant %d", got)
	}
	if got := withRaidEnchants(raid, ch).Gear[0].Enchant; got == 0 {
		t.Error("raid pass wears no enchant")
	}
	if got := plainRequest(raid, ch, 1, 1).Character.Gear[0].Enchant; got == 0 {
		t.Error("the raid pass's sim request carries no enchant")
	}
	if got := plainRequest(bare, ch, 1, 1).Character.Gear[0].Enchant; got != 0 {
		t.Errorf("the bare pass's sim request carries enchant %d", got)
	}
}

func TestTheRaidReportPublishesEachSlotsEnchant(t *testing.T) {
	spec := specInfo{Spec: "rogue-combat", ClassSlug: "rogue", RaidEnchants: enchantPlan{"hands": 8206}, Applied: &request.ResolvedPreset{}}
	rows := []slotRow{{Slot: "hands", ItemID: realGloves}, {Slot: "finger1", ItemID: realRing}, {Slot: "waist"}}
	got := withEnchantIDs(spec, rows)
	if got[0].EnchantID != 8206 || got[1].EnchantID != 0 || got[2].EnchantID != 0 {
		t.Errorf("enchant ids = %d %d %d, want 8206 0 0", got[0].EnchantID, got[1].EnchantID, got[2].EnchantID)
	}
	if rows[0].EnchantID != 0 {
		t.Error("withEnchantIDs changed its input")
	}
	spec.Applied = nil
	if withEnchantIDs(spec, rows)[0].EnchantID != 0 {
		t.Error("a bare report publishes an enchant")
	}
}

func TestOnlyARaidPassCreditsAWeaponItsEnchant(t *testing.T) {
	staff := candidate{ID: 1, ClassID: itemClassWeapon, TwoHand: true, Stats: map[string]float64{"spell_power": 10}}
	dagger := candidate{ID: 2, ClassID: itemClassWeapon, Stats: map[string]float64{"spell_power": 30}}
	ring := candidate{ID: 3, Stats: map[string]float64{"spell_power": 5}}
	items := []candidate{staff, dagger, ring}
	spec := specInfo{RaidEnchants: enchantPlan{"two_hand": 7662}}

	if got := withWeaponEnchantCredit(items, spec)[0].Stats["spell_power"]; got != 10 {
		t.Errorf("bare pass: the staff has %v spell power, want its own 10", got)
	}
	spec.Applied = &request.ResolvedPreset{}
	got := withWeaponEnchantCredit(items, spec)
	if got[0].Stats["spell_power"] != 75 {
		t.Errorf("the staff has %v spell power, want 10 + 65", got[0].Stats["spell_power"])
	}
	if got[1].Stats["spell_power"] != 30 || got[2].Stats["spell_power"] != 5 {
		t.Errorf("a one-hander or a ring took the two-hand enchant: %v %v", got[1].Stats, got[2].Stats)
	}
	if staff.Stats["spell_power"] != 10 {
		t.Error("withWeaponEnchantCredit changed its input")
	}
}
