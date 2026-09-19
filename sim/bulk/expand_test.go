package bulk

import (
	"errors"
	"fmt"
	"runtime"
	"slices"
	"strings"
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
//
// Every constant below - not just the ones this file happens to
// Lookup itself - names the kind of row its name claims: a bare
// itemRing IS a ring, itemOneHander+1 and +2 ARE one-handers too, not
// just itemOneHander. Task 14 uses several of these bare and at +1/+2,
// so a neighbour that merely compiles today but resolves to nothing -
// or to the wrong kind - would pass this task's tests and fail the
// next one's.
const (
	itemHelm = 12640 // head: Lionheart Helm (also in the warrior-fury fixture)
	// itemHelm2 is a second, unrelated real head item (Whitesoul Helm),
	// for the "two candidates for one slot" and cap-counting cases.
	// itemHelm+1 (12641) is Invulnerable Mail, a CHEST item - Task 14's
	// own brief assumed a +1 neighbour the way itemRing and
	// itemOneHander have one, but head has no such neighbour in this
	// build, so this constant is a distinct id found the same way
	// itemClassLocked and itemAllianceOnly were: a short scan of the
	// embedded table for a row of the needed kind.
	itemHelm2 = 12633 // head: Whitesoul Helm
	// itemOneHander, +1 and +2 are three consecutive real one-handers
	// (Silverbane Slicer, Guardian's Maulers, Swampspine Crusher; all
	// main_hand/off_hand, unrestricted, level 40), for the dual-wield
	// placement cases and their +1/+2 neighbours.
	itemOneHander = 277243
	itemTwoHander = 3488 // two-hand: Copper Battle Axe (main_hand only)
	// itemOffHand and +1 are two consecutive real off_hand-only items:
	// Grand Marshal's Left Hand Blade, High Warlord's Left Claw.
	itemOffHand = 18847
	// itemRing and +1 are two consecutive real rings, finger1/finger2:
	// Stalwart Watcher's Signet, Ferocious Watcher's Signet.
	itemRing = 275980
	// itemTrinket and +1 are two consecutive real trinkets,
	// trinket1/trinket2: Weakness Analyzer, Serenity Field.
	itemTrinket = 272438
	// itemClassLocked is a one-hander restricted to warriors only
	// (Heirloom Dagger) - the one constant above that must NOT be
	// usable by every class, for the "another class's item" case
	// below. It is kept separate from itemOneHander because no run of
	// three consecutive one-handers in the build has a class-locked
	// base: unrestricted and class-locked are two different needs,
	// so they get two different constants.
	itemClassLocked = 7116
	// itemAllianceOnly is Lorekeeper's Staff: class-unrestricted and
	// level 48, so faction is the only reason an orc cannot equip it -
	// isolating the faction branch of usable() the way itemClassLocked
	// isolates the class branch.
	itemAllianceOnly = 19571
	enchantWeapon    = 1900 // Enchant Weapon - Crusader: main_hand, off_hand
	enchantTwoHand   = 1903 // Enchant 2H Weapon - Major Spirit: two-hand only
	enchantHead      = 1506 // Lesser Arcanum of Voracity: head, legs (also in the fixture)
	// itemUniqueRing is any unique-equipped ring in the active build.
	// TestFindAUniqueRing prints candidates when this one stops being one.
	itemUniqueRing = 19432
)

func TestFindAUniqueRing(t *testing.T) {
	item, ok := simdb.Lookup(itemUniqueRing)
	if ok && item.Unique && slices.Contains(item.Slots, "finger1") {
		return
	}
	t.Errorf("item %d is no longer a unique ring in this build; pick another from the build's table", itemUniqueRing)
}

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
		{"another class's item", func(r *api.SimRequest) { r.Character.Class = "mage"; r.Spec = "mage-frost" }, itemClassLocked},
		{"a locked slot", func(r *api.SimRequest) { r.Bulk.Locked = []string{"finger1", "finger2"} }, itemRing + 1},
		{"an alliance-only item", func(*api.SimRequest) {}, itemAllianceOnly},
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

// usable's three branches - class, level, faction - above are also
// exercised end-to-end through real build rows, but the build's own
// level cap means no item in it carries a RequiredLevel above
// api.SimLevel (60): the sim never runs any other level, so nothing
// upstream of usable ever itemises one. That branch can only be
// proven directly, against a synthetic row, which this table does for
// all three so they are checked the same way rather than two real
// rows and one that cannot exist.
func TestUsableEligibilityRules(t *testing.T) {
	warrior := api.CharacterSpec{Class: "warrior", Race: "orc", Level: api.SimLevel}
	cases := []struct {
		name string
		item simdb.Item
		ch   api.CharacterSpec
		want bool
	}{
		{"open to every class", simdb.Item{}, warrior, true},
		{"restricted to a class the character has", simdb.Item{Classes: []string{"warrior"}}, warrior, true},
		{"restricted to a class the character does not have", simdb.Item{Classes: []string{"mage"}}, warrior, false},
		{"at the character's level", simdb.Item{RequiredLevel: api.SimLevel}, warrior, true},
		{"above the character's level", simdb.Item{RequiredLevel: api.SimLevel + 1}, warrior, false},
		{"faction-open", simdb.Item{Faction: simdb.FactionAny}, warrior, true},
		{"the character's own faction", simdb.Item{Faction: simdb.FactionHorde}, warrior, true},
		{"the other faction", simdb.Item{Faction: simdb.FactionAlliance}, warrior, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := usable(c.item, c.ch); got != c.want {
				t.Errorf("usable(%+v, %s/%s@%d) = %v, want %v", c.item, c.ch.Race, c.ch.Class, c.ch.Level, got, c.want)
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
		{"inherits an enchant that has no item-type restriction onto a two-hander", api.Candidate{Slot: "main_hand", ItemID: itemTwoHander, Origin: api.OriginBag}, enchantWeapon},
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

// Gear mode is every valid combination; drops and talents mode are one
// substitution at a time. That is the whole difference between Top
// Gear and Droptimizer, and it is the mode that says which.
func TestGearModeCombinesAndDropsModeDoesNot(t *testing.T) {
	two := []api.Candidate{candidate("head", itemHelm), candidate("finger2", itemRing+1)}
	gearMode, err := Expand(withBulk(api.KindGear, two...))
	if err != nil {
		t.Fatal(err)
	}
	// head-or-not times finger2-or-not, minus the "neither" case,
	// which is the equipped set and is not a combination.
	if len(gearMode) != 3 {
		t.Errorf("gear mode produced %d combinations, want 3", len(gearMode))
	}
	var both int
	for _, c := range gearMode {
		if len(slotsOf(c)) == 2 {
			both++
		}
	}
	if both != 1 {
		t.Errorf("gear mode produced %d combinations that change both slots, want 1", both)
	}

	dropsCandidates := []api.Candidate{
		{Slot: "head", ItemID: itemHelm, Origin: "drop:raid:mc:lucifron"},
		{Slot: "finger2", ItemID: itemRing + 1, Origin: "drop:raid:mc:lucifron"},
	}
	dropsMode, err := Expand(withBulk(api.KindDrops, dropsCandidates...))
	if err != nil {
		t.Fatal(err)
	}
	if len(dropsMode) != 2 {
		t.Errorf("drops mode produced %d combinations, want 2", len(dropsMode))
	}
	for _, c := range dropsMode {
		if len(slotsOf(c)) != 1 {
			t.Errorf("drops mode changed %d slots at once", len(slotsOf(c)))
		}
	}
}

// Two candidates for one slot are alternatives, never both at once.
func TestTwoCandidatesForOneSlotAreAlternatives(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear,
		candidate("head", itemHelm), candidate("head", itemHelm2)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d combinations, want 2", len(got))
	}
	for _, c := range got {
		if len(c.Substitutions) != 1 {
			t.Errorf("two items landed in one slot: %+v", c.Substitutions)
		}
	}
}

func TestWeaponShapesAndDuplicates(t *testing.T) {
	cases := []struct {
		name       string
		candidates []api.Candidate
		edit       func(*api.SimRequest)
		wantSlots  [][]string // the combinations, as sorted slot lists
	}{
		{
			name:       "a two-hander competes with main-plus-off-hand",
			candidates: []api.Candidate{{ItemID: itemTwoHander, Origin: api.OriginBag}},
			// The off hand is emptied rather than the combination
			// being thrown away: a two-hander IS a shape, and refusing
			// it would mean Top Gear never ranked one.
			wantSlots: [][]string{{"main_hand"}},
		},
		{
			name:       "a one-hander is tried in both hands",
			candidates: []api.Candidate{{ItemID: itemOneHander + 1, Origin: api.OriginBag}},
			wantSlots:  [][]string{{"main_hand"}, {"off_hand"}},
		},
		{
			name: "dual wield tries both orders",
			candidates: []api.Candidate{
				{ItemID: itemOneHander + 1, Origin: api.OriginBag},
				{ItemID: itemOneHander + 2, Origin: api.OriginBag},
			},
			wantSlots: [][]string{
				{"main_hand"}, {"off_hand"},
				{"main_hand"}, {"off_hand"},
				{"main_hand", "off_hand"}, {"main_hand", "off_hand"},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := withBulk(api.KindGear, c.candidates...)
			if c.edit != nil {
				c.edit(&req)
			}
			got, err := Expand(req)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(c.wantSlots) {
				var shapes [][]string
				for _, combo := range got {
					shapes = append(shapes, slotsOf(combo))
				}
				t.Fatalf("got %d combinations %v, want %d %v", len(got), shapes, len(c.wantSlots), c.wantSlots)
			}
		})
	}
}

// A two-hander in the main hand empties the off hand; nothing may be
// wielded beside it.
func TestATwoHanderEmptiesTheOffHand(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, api.Candidate{ItemID: itemTwoHander, Origin: api.OriginBag}))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d combinations", len(got))
	}
	if g := gearAt(got[0], "off_hand"); g.ItemID != 0 {
		t.Errorf("a two-hander left item %d in the off hand", g.ItemID)
	}
}

// The engine's own validity rules: one item cannot be in both ring
// slots, nor both trinket slots, and two rings or trinkets sharing a
// name are the same item at two qualities.
func TestDuplicateRingsAndTrinketsAreRefused(t *testing.T) {
	// finger1 already holds itemRing; offering it for finger2 would
	// produce a character wearing two of one ring.
	got, err := Expand(withBulk(api.KindGear, candidate("finger2", itemRing)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("the same ring was equipped twice: %+v", got[0].Request.Character.Gear)
	}

	got, err = Expand(withBulk(api.KindGear, candidate("trinket2", itemTrinket)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("the same trinket was equipped twice")
	}
}

// Unique-equipped is one at a time anywhere, not one per slot.
func TestUniqueEquippedIsRespected(t *testing.T) {
	// itemUniqueRing must be a row whose Unique flag is set. Find one
	// with the helper below and replace the constant if this id is not
	// unique in the active build.
	req := base()
	req.Character.Gear = append(req.Character.Gear, api.GearSlot{Slot: "finger2", ItemID: itemUniqueRing})
	req.Bulk = &api.BulkSpec{Mode: api.KindGear, Precision: api.PrecisionNormal, Cap: 400,
		Candidates: []api.Candidate{candidate("finger1", itemUniqueRing)}}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("a unique-equipped item was equipped twice: %+v", got[0].Request.Character.Gear)
	}
}

// Talent loadouts are a dimension of their own: in gear mode they
// multiply the gear combinations, and in talents mode they are the
// only candidates.
func TestTalentLoadoutsAreADimension(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Bulk.Talents = []api.TalentLoadout{
		{Name: "Deep Fury", Talents: "30305001302-05050005525010052"},
		{Name: "Two-hand Arms", Talents: "30305001302-05050005525010053"},
	}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	// (the helm, or not) times (own talents, Deep Fury, Two-hand Arms),
	// minus the equipped set: 2 * 3 - 1 = 5.
	if len(got) != 5 {
		t.Errorf("got %d combinations, want 5", len(got))
	}

	only := withBulk(api.KindTalents)
	only.Bulk.Candidates = nil
	only.Bulk.Talents = req.Bulk.Talents
	got, err = Expand(only)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("talents mode produced %d combinations, want 2", len(got))
	}
	for _, c := range got {
		if len(c.Substitutions) != 1 || c.Substitutions[0].Kind != api.SubstitutionTalents {
			t.Errorf("a talents-mode combination substituted gear: %+v", c.Substitutions)
		}
		if c.Request.Character.Talents == base().Character.Talents {
			t.Error("a loadout did not reach the request")
		}
	}
}

// Alternative consumable lists are a dimension too, and each REPLACES
// the character's own list rather than adding to it.
func TestConsumableListsAreADimension(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Character.Consumes = []string{"elixir_of_the_mongoose"}
	req.Bulk.Consumables = [][]string{
		{"flask_of_the_titans"},
		{"juju_power", "juju_might"},
	}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	// (the helm, or not) times (own, flask, jujus), minus the
	// equipped set with its own consumables: 2 * 3 - 1 = 5.
	if len(got) != 5 {
		t.Fatalf("got %d combinations, want 5", len(got))
	}
	var replaced int
	for _, c := range got {
		for _, sub := range c.Substitutions {
			if sub.Kind != api.SubstitutionConsumes {
				continue
			}
			replaced++
			// The chip's label is the ids joined, so the API can
			// compose a headline from Name whatever the kind is.
			if sub.Name != strings.Join(sub.Consumes, ", ") || sub.Name == "" {
				t.Errorf("the consumables chip is named %q for %v", sub.Name, sub.Consumes)
			}
			if len(c.Request.Character.Consumes) != len(sub.Consumes) {
				t.Errorf("the request carries %v and the chip says %v",
					c.Request.Character.Consumes, sub.Consumes)
			}
			for i := range sub.Consumes {
				if c.Request.Character.Consumes[i] != sub.Consumes[i] {
					t.Errorf("the request carries %v and the chip says %v",
						c.Request.Character.Consumes, sub.Consumes)
				}
			}
			if sub.Consumes[0] == "elixir_of_the_mongoose" {
				t.Error("a consumables chip named the character's own list")
			}
		}
	}
	// Two alternatives times two gear choices.
	if replaced != 4 {
		t.Errorf("%d combinations changed the consumables, want 4", replaced)
	}
	// The combinations that did NOT change them still carry the
	// character's own.
	for _, c := range got {
		if len(c.Substitutions) == 1 && c.Substitutions[0].Kind == api.SubstitutionItem {
			if len(c.Request.Character.Consumes) != 1 || c.Request.Character.Consumes[0] != "elixir_of_the_mongoose" {
				t.Errorf("a gear-only combination lost the character's consumables: %v", c.Request.Character.Consumes)
			}
		}
	}
}

// An empty alternative consumables list is a real thing to compare
// against - api validation blesses it as "no consumables" - so its
// chip needs a label too, not strings.Join(nil, ", ") = "".
func TestAnEmptyConsumablesListIsNamed(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Bulk.Consumables = [][]string{{}}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, c := range got {
		for _, sub := range c.Substitutions {
			if sub.Kind != api.SubstitutionConsumes {
				continue
			}
			found = true
			if sub.Name != api.NoConsumablesLabel {
				t.Errorf("the empty consumables chip is named %q, want %q", sub.Name, api.NoConsumablesLabel)
			}
			if len(sub.Consumes) != 0 {
				t.Errorf("the empty consumables chip carries %v", sub.Consumes)
			}
			if len(c.Request.Character.Consumes) != 0 {
				t.Errorf("the request carries %v for the empty alternative", c.Request.Character.Consumes)
			}
		}
	}
	if !found {
		t.Fatal("no combination substituted the empty consumables list")
	}
}

// A set replaces every slot at once, so it is an alternative to the
// whole gear product rather than a member of it.
func TestASetReplacesEverySlot(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Bulk.Sets = []api.GearSet{{Name: "my AQ set", Gear: []api.GearSlot{
		{Slot: "head", ItemID: itemHelm2},
		{Slot: "main_hand", ItemID: itemTwoHander},
	}}}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	// the helm on its own, and the set: two combinations.
	if len(got) != 2 {
		t.Fatalf("got %d combinations, want 2", len(got))
	}
	var set *Combination
	for i, c := range got {
		if len(c.Substitutions) == 1 && c.Substitutions[0].Kind == api.SubstitutionSet {
			set = &got[i]
		}
	}
	if set == nil {
		t.Fatal("no combination substituted the set")
	}
	if len(set.Request.Character.Gear) != 2 {
		t.Errorf("the set did not replace the whole character's gear: %+v", set.Request.Character.Gear)
	}
	if set.Substitutions[0].Name != "my AQ set" {
		t.Errorf("the chip does not name the set: %+v", set.Substitutions[0])
	}
}

// A set that is not valid equipment on its own - here, a two-hander
// beside an off-hand - is refused outright rather than silently
// dropped: apply's two-hand clearing only runs on candidate
// placements, never on a set's own gear list, so an invalid set would
// otherwise have every combination it produces deleted by valid() and
// vanish from the result with nothing to say why. That is exactly the
// malformed input the package doc says gets refused, not skipped.
func TestAnInvalidSetIsRefused(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Bulk.Sets = []api.GearSet{{Name: "bad set", Gear: []api.GearSlot{
		{Slot: "main_hand", ItemID: itemTwoHander},
		{Slot: "off_hand", ItemID: itemOffHand},
	}}}
	_, err := Expand(req)
	if !errors.Is(err, ErrInvalidSet) {
		t.Fatalf("Expand = %v, want ErrInvalidSet", err)
	}
	if !strings.Contains(err.Error(), "bad set") {
		t.Errorf("error %q does not name the set", err)
	}
}

// The cap is a refusal with both numbers, never a silent trim: a
// ranking of a subset nobody chose looks exactly like a ranking.
func TestTheCapRefusesRatherThanTrims(t *testing.T) {
	req := withBulk(api.KindGear,
		candidate("head", itemHelm), candidate("head", itemHelm2),
		candidate("finger2", itemRing+1), candidate("trinket2", itemTrinket+1))
	req.Bulk.Cap = 3
	_, err := Expand(req)
	var capped api.ErrCapExceeded
	if !errors.As(err, &capped) {
		t.Fatalf("Expand = %v, want ErrCapExceeded", err)
	}
	if capped.Cap != 3 || capped.Combinations <= 3 {
		t.Errorf("ErrCapExceeded = %+v", capped)
	}
}

// The count and the plan are the same number, always. A count the
// page shows and a plan the run executes that disagreed would be a
// cap notice nobody could act on.
func TestCountAgreesWithExpand(t *testing.T) {
	req := withBulk(api.KindGear,
		candidate("head", itemHelm), candidate("head", itemHelm2),
		candidate("finger2", itemRing+1), candidate("trinket2", itemTrinket+1))
	combos, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	n, err := Count(req)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(combos) {
		t.Errorf("Count = %d, Expand = %d", n, len(combos))
	}
}

// A count over the cap is the same refusal, with the same numbers -
// not just the same Cap, but the exact same Combinations Expand would
// have reported: combinations counts every combination it enumerates
// before deciding what to retain, so the two can never disagree even
// though Count never builds the requests Expand returns.
func TestCountRefusesOverTheCap(t *testing.T) {
	req := withBulk(api.KindGear,
		candidate("head", itemHelm), candidate("head", itemHelm2),
		candidate("finger2", itemRing+1), candidate("trinket2", itemTrinket+1))
	req.Bulk.Cap = 3
	_, expandErr := Expand(req)
	_, countErr := Count(req)
	var expandCapped, countCapped api.ErrCapExceeded
	if !errors.As(expandErr, &expandCapped) {
		t.Fatalf("Expand = %v, want ErrCapExceeded", expandErr)
	}
	if !errors.As(countErr, &countCapped) {
		t.Fatalf("Count = %v, want ErrCapExceeded", countErr)
	}
	if countCapped.Cap != 3 {
		t.Errorf("ErrCapExceeded = %+v", countCapped)
	}
	if countCapped != expandCapped {
		t.Errorf("Count's cap error = %+v, Expand's = %+v", countCapped, expandCapped)
	}
}

// A regression test for retention itself, not just for the count: a
// future change that went back to `out := make([]Combination, 0,
// len(sets)*len(loadouts)*len(drinks))` would still return the right
// answer and pass every other test in this package - the count would
// still be exact - but would allocate proportionally to the whole
// product instead of to the cap. This exercises combinations' own
// retention logic directly (talents mode, so gearCombinations and
// walkGearChoices are not involved), pinning it at the small side of a
// large gap: a 2,000-loadout request capped at 10 must not leave
// anywhere near 2,000 combinations' worth of api.SimRequest behind.
func TestCombinationsRetentionStaysAtTheCap(t *testing.T) {
	req := withBulk(api.KindTalents)
	req.Bulk.Candidates = nil
	for i := 0; i < 2000; i++ {
		req.Bulk.Talents = append(req.Bulk.Talents, api.TalentLoadout{
			Name:    fmt.Sprintf("loadout-%d", i),
			Talents: "30305001302-05050005525010051",
		})
	}
	req.Bulk.Cap = 10

	places, err := placements(req, Options{})
	if err != nil {
		t.Fatal(err)
	}

	// On a cap breach, combinations discards out and returns
	// (nil, ErrCapExceeded{...}) either way - whether it retained at
	// most Cap+1 throughout or retained all 2,000 and threw them away
	// at the end. The return value alone cannot tell those apart, so
	// retentionProbe watches len(out) DURING enumeration, which is the
	// only place the guard's effect is observable at all.
	var peak int
	retentionProbe = func(retained int) {
		if retained > peak {
			peak = retained
		}
	}
	defer func() { retentionProbe = nil }()

	_, err = combinations(req, places)

	var capped api.ErrCapExceeded
	if !errors.As(err, &capped) || capped.Combinations != 2000 {
		t.Fatalf("combinations(...) = %v, want ErrCapExceeded{Combinations: 2000}", err)
	}
	// The guard is `if len(out) <= req.Bulk.Cap`, so out can reach
	// Cap+1 elements (the one append that proves the guard is about to
	// stop, not Cap itself) before it stops growing. Anything larger
	// means retention is no longer bounded by the cap.
	if want := req.Bulk.Cap + 1; peak > want {
		t.Errorf("combinations retained %d combinations at once for a cap of %d, want at most %d", peak, req.Bulk.Cap, want)
	}
}

// A regression test for walkGearChoices' laziness specifically: nine
// slots with three candidates apiece is the exact shape that used to
// materialise the whole [][]placement product - hundreds of megabytes
// - before yielding a single combination. Synthetic slot names (not
// real gear slots) keep this independent of simdb and of
// sameWeaponTwice, which only ever looks at main_hand/off_hand.
func TestWalkGearChoicesDoesNotMaterialiseTheProduct(t *testing.T) {
	const slotCount, candidatesPerSlot = 9, 3
	bySlot := map[string][]placement{}
	slots := make([]string, 0, slotCount)
	for i := 0; i < slotCount; i++ {
		slot := fmt.Sprintf("synthetic-slot-%d", i)
		slots = append(slots, slot)
		for j := 0; j < candidatesPerSlot; j++ {
			bySlot[slot] = append(bySlot[slot], placement{
				Slot: slot,
				Gear: api.GearSlot{Slot: slot, ItemID: i*100 + j},
			})
		}
	}
	// "keep" plus each slot's candidates, to the power of the slot
	// count: (1+3)^9 = 262,144 - the shape the re-review measured at
	// 739.4 MB before this fix.
	want := 1
	for range slots {
		want *= candidatesPerSlot + 1
	}

	runtime.GC()
	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	var got int
	walkGearChoices(slots, bySlot, func(chosen []placement) { got++ })

	var after runtime.MemStats
	runtime.ReadMemStats(&after)

	if got != want {
		t.Fatalf("walkGearChoices yielded %d combinations, want %d", got, want)
	}
	// Two orders of magnitude below the 739.4 MB the eager product used
	// to cost for this exact shape, and well above ordinary GC noise.
	const ceiling = 20 << 20 // 20 MB
	if delta := after.HeapAlloc - before.HeapAlloc; delta > ceiling {
		t.Errorf("walkGearChoices grew the heap by %d bytes (%.1f MB), want under %d (%.0f MB)",
			delta, float64(delta)/1e6, ceiling, float64(ceiling)/1e6)
	}
}
