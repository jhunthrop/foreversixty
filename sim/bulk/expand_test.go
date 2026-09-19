package bulk

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

// The ids below are real rows of the active build, checked directly
// against simdb.Lookup and simdb.Enchants rather than assumed: the
// brief this test was written from predates the current data, and
// most of its ids either had no row at all or named a different kind
// of item than the test needed (its "two-hander" was a one-hander;
// its "weapon enchant" was a head/legs enchant). itemHelm and
// enchantHead are also part of the warrior-fury fixture that
// TestFixtureGearResolves (sim/internal/simdb) checks on every run; the
// rest were found with a short throwaway scan of the embedded tables
// for a row of the needed kind (hand type, restricted-to-a-class, an
// item_types-restricted enchant) usable by an orc warrior. If one of
// these ever leaves the build, the eligibility cases below fail with
// "placed in [] want [...]" or an unexpected error, which is the
// signal to look it up again the same way.
const (
	itemHelm      = 12640 // head: Lionheart Helm (also in the warrior-fury fixture)
	itemOneHander = 7116  // one-hand, warrior-only: Heirloom Dagger (main_hand, off_hand)
	itemTwoHander = 3488  // two-hand: Copper Battle Axe (main_hand only)
	// itemOffHand+1 (11863) is White Bone Shredder, an off_hand-only item.
	itemOffHand = 11862
	// itemRing+1 (19382) is Pure Elementium Band, finger1/finger2.
	itemRing = 19381
	// itemTrinket+1 (19406) is Drake Fang Talisman, trinket1/trinket2.
	itemTrinket    = 19405
	enchantWeapon  = 1900 // Enchant Weapon - Crusader: main_hand, off_hand
	enchantTwoHand = 1903 // Enchant 2H Weapon - Major Spirit: two-hand only
	enchantHead    = 1506 // Lesser Arcanum of Voracity: head, legs (also in the fixture)
)

// base is a fury warrior with a full set of the fixture's own gear, so
// every slot the tests substitute into is already filled.
func base() api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          "warrior-fury",
		Source:        api.CharacterSource{Kind: api.SourceManual},
		Character: api.CharacterSpec{
			Name:    "Thrall",
			Race:    "orc",
			Class:   "warrior",
			Level:   api.SimLevel,
			Talents: "30305001302-05050005525010051",
			Gear: []api.GearSlot{
				{Slot: "head", ItemID: itemHelm, Enchant: enchantHead},
				{Slot: "main_hand", ItemID: itemOneHander, Enchant: enchantWeapon},
				{Slot: "off_hand", ItemID: itemOffHand},
				{Slot: "finger1", ItemID: itemRing},
				{Slot: "trinket1", ItemID: itemTrinket},
			},
		},
		Encounter:  api.DefaultEncounter(),
		Iterations: 3000,
		RandomSeed: 7,
	}
}

// withBulk is base with a bulk block that names the given candidates.
func withBulk(mode string, candidates ...api.Candidate) api.SimRequest {
	req := base()
	req.Bulk = &api.BulkSpec{
		Mode:       mode,
		Precision:  api.PrecisionNormal,
		Cap:        api.Caps[api.LaneServer],
		Candidates: candidates,
	}
	return req
}

func candidate(slot string, id int) api.Candidate {
	return api.Candidate{Slot: slot, ItemID: id, Origin: api.OriginBag}
}

// slotsOf lists the slots one combination substitutes into, sorted, so
// a case can assert a shape without depending on iteration order.
func slotsOf(c Combination) []string {
	out := make([]string, 0, len(c.Substitutions))
	for _, s := range c.Substitutions {
		if s.Kind == api.SubstitutionItem {
			out = append(out, s.Slot)
		}
	}
	slices.Sort(out)
	return out
}

// gearAt is the item the combination's request has in a slot.
func gearAt(c Combination, slot string) api.GearSlot {
	for _, g := range c.Request.Character.Gear {
		if g.Slot == slot {
			return g
		}
	}
	return api.GearSlot{}
}

func TestExpandRefusesARequestWithNoBulkBlock(t *testing.T) {
	if _, err := Expand(base()); !errors.Is(err, ErrNotBulk) {
		t.Errorf("Expand of a plain run = %v, want ErrNotBulk", err)
	}
}

func TestExpandRefusesAnItemTheBuildDoesNotCarry(t *testing.T) {
	_, err := Expand(withBulk(api.KindGear, candidate("head", 999999999)))
	if !errors.Is(err, ErrUnknownItem) {
		t.Errorf("Expand = %v, want ErrUnknownItem", err)
	}
}

// A candidate with no slot goes wherever it fits, which is what makes
// "try this ring in both slots" and "try this sword in either hand"
// one rule rather than two.
func TestACandidateWithNoSlotGoesWhereverItFits(t *testing.T) {
	cases := []struct {
		name  string
		item  int
		slots []string
	}{
		{"a ring", itemRing + 1, []string{"finger1", "finger2"}},
		{"a trinket", itemTrinket + 1, []string{"trinket1", "trinket2"}},
		{"a one-hander", itemOneHander, []string{"main_hand", "off_hand"}},
		{"a two-hander", itemTwoHander, []string{"main_hand"}},
		{"a helm", itemHelm, []string{"head"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Expand(withBulk(api.KindDrops, api.Candidate{ItemID: c.item, Origin: "drop:raid:mc:lucifron"}))
			if err != nil {
				t.Fatal(err)
			}
			var slots []string
			for _, combo := range got {
				slots = append(slots, slotsOf(combo)...)
			}
			slices.Sort(slots)
			if !slices.Equal(slots, c.slots) {
				t.Errorf("placed in %v, want %v", slots, c.slots)
			}
		})
	}
}

func TestACandidateNamingASlotGoesOnlyThere(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, candidate("finger2", itemRing+1)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || !slices.Equal(slotsOf(got[0]), []string{"finger2"}) {
		t.Fatalf("got %d combinations in %v", len(got), slotsOf(got[0]))
	}
}

func TestACandidateNamingASlotItDoesNotFitIsSkipped(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, candidate("head", itemOneHander)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("a sword was placed on a head: %v", slotsOf(got[0]))
	}
}

// A candidate the character cannot equip is SKIPPED, not refused: a
// boss's loot table is the whole table, and Droptimizer sends it.
func TestCandidatesTheCharacterCannotEquipAreSkipped(t *testing.T) {
	cases := []struct {
		name string
		edit func(*api.SimRequest)
		item int
	}{
		{"another class's item", func(r *api.SimRequest) { r.Character.Class = "mage"; r.Spec = "mage-frost" }, itemOneHander},
		{"a locked slot", func(r *api.SimRequest) { r.Bulk.Locked = []string{"finger1", "finger2"} }, itemRing + 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := withBulk(api.KindDrops, api.Candidate{ItemID: c.item, Origin: "drop:raid:mc:lucifron"})
			c.edit(&req)
			got, err := Expand(req)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != 0 {
				t.Errorf("%d combinations survived: %v", len(got), slotsOf(got[0]))
			}
		})
	}
}

// An enchant of zero inherits the equipped item's, where it fits. That
// is what stops Top Gear ranking an unenchanted upgrade below an
// enchanted item the player already wears.
// enchantFixture answers the expansion tests' enchant questions from
// a tiny allowlist. The fit RULES are tested in sim/internal/simdb
// against the build's own table; what these tests need is a table
// whose answers they can predict, so that a build whose enchant rows
// move does not silently turn an inheritance assertion green.
type enchantFixture struct{}

func (enchantFixture) Fits(effectID int, slot string, item simdb.Item, _ string) error {
	switch effectID {
	case enchantWeapon:
		if slot == "main_hand" || slot == "off_hand" {
			return nil
		}
	case enchantTwoHand:
		if slot == "main_hand" && item.HandType == simdb.HandTwo {
			return nil
		}
	case enchantHead:
		if slot == "head" || slot == "legs" {
			return nil
		}
	default:
		return fmt.Errorf("%w: %d", simdb.ErrUnknownEnchant, effectID)
	}
	return fmt.Errorf("%w: enchant %d in %q", simdb.ErrEnchantDoesNotFit, effectID, slot)
}

func TestEnchantInheritance(t *testing.T) {
	table := enchantFixture{}
	cases := []struct {
		name        string
		candidate   api.Candidate
		wantEnchant int
	}{
		{"inherits the equipped weapon enchant", api.Candidate{Slot: "main_hand", ItemID: itemOneHander + 1, Origin: api.OriginBag}, enchantWeapon},
		{"inherits nothing into an unenchanted slot", api.Candidate{Slot: "off_hand", ItemID: itemOffHand + 1, Origin: api.OriginBag}, 0},
		{"keeps the enchant it was given", api.Candidate{Slot: "head", ItemID: itemHelm, Enchant: enchantHead, Origin: api.OriginBag}, enchantHead},
		{"does not inherit an enchant that does not fit", api.Candidate{Slot: "main_hand", ItemID: itemTwoHander, Origin: api.OriginBag}, enchantWeapon},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ExpandWith(withBulk(api.KindGear, c.candidate), Options{Enchants: table})
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != 1 {
				t.Fatalf("got %d combinations", len(got))
			}
			slot := got[0].Substitutions[0].Slot
			if g := gearAt(got[0], slot); g.Enchant != c.wantEnchant {
				t.Errorf("%s carries enchant %d, want %d", slot, g.Enchant, c.wantEnchant)
			}
		})
	}
}

// A two-hand enchant cannot be inherited onto a one-hander, so a
// candidate that would have inherited one carries none instead. The
// case above covers the reverse; this is the one that used to silently
// apply an enchant the item cannot hold.
func TestAnInheritedEnchantThatDoesNotFitIsDropped(t *testing.T) {
	req := base()
	req.Character.Gear = []api.GearSlot{{Slot: "main_hand", ItemID: itemTwoHander, Enchant: enchantTwoHand}}
	req.Bulk = &api.BulkSpec{Mode: api.KindGear, Precision: api.PrecisionNormal, Cap: 400,
		Candidates: []api.Candidate{candidate("main_hand", itemOneHander)}}
	got, err := ExpandWith(req, Options{Enchants: enchantFixture{}})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d combinations", len(got))
	}
	if g := gearAt(got[0], "main_hand"); g.Enchant != 0 {
		t.Errorf("a two-hand enchant was inherited onto a one-hander: %d", g.Enchant)
	}
}

// With no Options, Expand uses the table embedded beside the items,
// so nothing has to be fetched or passed in (contract A9). An effect
// id the build has never carried is refused there, not here.
func TestExpandUsesTheEmbeddedEnchantTable(t *testing.T) {
	c := candidate("head", itemHelm)
	c.Enchant = 999999
	_, err := Expand(withBulk(api.KindGear, c))
	if !errors.Is(err, simdb.ErrUnknownEnchant) {
		t.Errorf("Expand = %v, want simdb.ErrUnknownEnchant from the embedded table", err)
	}
}

// A named enchant that does not fit is the page sending something it
// should not, and it is refused rather than dropped: the player asked
// for that enchant and would otherwise read a ranking of something
// else.
func TestANamedEnchantThatDoesNotFitIsRefused(t *testing.T) {
	c := candidate("head", itemHelm)
	c.Enchant = enchantWeapon
	_, err := ExpandWith(withBulk(api.KindGear, c), Options{Enchants: enchantFixture{}})
	if !errors.Is(err, simdb.ErrEnchantDoesNotFit) {
		t.Errorf("Expand = %v, want simdb.ErrEnchantDoesNotFit", err)
	}
}

// The substitution the result carries says everything the page shows
// in a chip: the slot it took, the item, the enchant it ended up with,
// the suffix and where the candidate came from.
func TestTheSubstitutionCarriesTheWholeChip(t *testing.T) {
	c := api.Candidate{Slot: "finger2", ItemID: itemRing + 1, Suffix: 1805,
		Origin: "drop:raid:mc:lucifron", SourceName: "Lucifron"}
	got, err := Expand(withBulk(api.KindDrops, c))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d combinations", len(got))
	}
	s := got[0].Substitutions[0]
	if s.Kind != api.SubstitutionItem || s.Slot != "finger2" || s.ItemID != c.ItemID || s.Suffix != 1805 || s.Origin != c.Origin {
		t.Errorf("substitution = %+v", s)
	}
	if s.SourceName != "Lucifron" {
		t.Errorf("the chip lost the source name: %+v", s)
	}
	// The item's own name, so the API can compose a headline without
	// an item table of its own.
	want, _ := simdb.Lookup(c.ItemID)
	if s.Name != want.Name || s.Name == "" {
		t.Errorf("the chip names the item %q, the build calls it %q", s.Name, want.Name)
	}
	if g := gearAt(got[0], "finger2"); g.ItemID != c.ItemID || g.Suffix != 1805 {
		t.Errorf("the request does not carry the substitution: %+v", g)
	}
}

// Every combination is a runnable request: the bulk block is gone (a
// stage request is a plain run) and nothing else about the character
// moved.
func TestACombinationsRequestIsAPlainRun(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, candidate("head", itemHelm)))
	if err != nil {
		t.Fatal(err)
	}
	if got[0].Request.Bulk != nil {
		t.Error("a combination's request still carries a bulk block")
	}
	if got[0].Request.Spec != base().Spec || got[0].Request.Character.Talents != base().Character.Talents {
		t.Error("a combination changed something other than the gear")
	}
}
