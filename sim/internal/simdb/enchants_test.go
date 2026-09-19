package simdb

import (
	"errors"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// excerpt is ten real rows cut from data/builds/1.60.1.69893/enchants.json
// (see testdata/enchants.excerpt.json's comment-free header for the exact
// ids). Two of the ten ids are shared: 41 (a wrist row and a chest row,
// both "normal") and 241 (a "two_hand" row and a "normal" row), which is
// why the merged table carries eight enchants, not ten - parseEnchants
// merges same-id rows rather than dropping the later one (see enchants.go's
// package doc comment).
func excerpt(t *testing.T) map[int]Enchant {
	t.Helper()
	b, err := os.ReadFile("testdata/enchants.excerpt.json")
	if err != nil {
		t.Fatal(err)
	}
	table, err := parseEnchants(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(table) != 8 {
		t.Fatalf("the excerpt carries %d enchants, want 8 (ten rows, two shared ids)", len(table))
	}
	return table
}

func TestEnchantFit(t *testing.T) {
	table := excerpt(t)
	twoHander := Item{ID: 1, HandType: HandTwo, WeaponType: "two_handed_axe"}
	oneHander := Item{ID: 2, HandType: HandOne, WeaponType: "sword"}
	shield := Item{ID: 3, WeaponType: "shield"}
	helm := Item{ID: 4}
	wristItem := Item{ID: 5}
	chestArmor := Item{ID: 6}
	feetItem := Item{ID: 7}

	cases := []struct {
		name    string
		effect  int
		slot    string
		item    Item
		class   string
		wantErr error
	}{
		{name: "a normal weapon enchant on a one-hander in the main hand", effect: 1900, slot: "main_hand", item: oneHander, class: "warrior"},
		{name: "a normal weapon enchant on a two-hander too", effect: 1900, slot: "main_hand", item: twoHander, class: "warrior"},
		{name: "a normal weapon enchant in the off hand", effect: 1900, slot: "off_hand", item: oneHander, class: "warrior"},
		{name: "a normal weapon enchant on a helm", effect: 1900, slot: "head", item: helm, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a two-hand enchant on a two-hander", effect: 2646, slot: "main_hand", item: twoHander, class: "warrior"},
		{name: "a two-hand enchant on a one-hander", effect: 2646, slot: "main_hand", item: oneHander, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a shield enchant on a shield", effect: 863, slot: "off_hand", item: shield, class: "warrior"},
		{name: "a shield enchant on a one-hander", effect: 863, slot: "off_hand", item: oneHander, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "an unrestricted kit on the chest", effect: 15, slot: "chest", item: chestArmor, class: "warrior"},
		{name: "a class-restricted kit for the right class", effect: 2588, slot: "head", item: helm, class: "mage"},
		{name: "a class-restricted kit for the wrong class", effect: 2588, slot: "head", item: helm, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a shared id fits the wrist, from its wrist sharer", effect: 41, slot: "wrist", item: wristItem, class: "warrior"},
		{name: "a shared id also fits the chest, from its chest sharer", effect: 41, slot: "chest", item: chestArmor, class: "warrior"},
		{name: "a shared id does not fit a slot neither sharer named", effect: 41, slot: "feet", item: feetItem, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a shared id whose sharers disagree on item_types fits a two-hander", effect: 241, slot: "main_hand", item: twoHander, class: "warrior"},
		{name: "the same shared id also fits a one-hander, from its normal sharer", effect: 241, slot: "main_hand", item: oneHander, class: "warrior"},
		{name: "an enchant the build has never heard of", effect: 999999, slot: "head", item: helm, class: "warrior", wantErr: ErrUnknownEnchant},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := fits(table, c.effect, c.slot, c.item, c.class)
			if !errors.Is(err, c.wantErr) {
				t.Errorf("fits = %v, want %v", err, c.wantErr)
			}
		})
	}
}

func TestParseEnchantsRefusesRubbish(t *testing.T) {
	for _, body := range []string{``, `[]`, `[{"name":"no id"}]`, `{"not":"an array"}`} {
		if _, err := parseEnchants([]byte(body)); err == nil {
			t.Errorf("parseEnchants accepted %q", body)
		}
	}
}

// A shared id's merged row keeps the first sharer's name (file order),
// and its lowest non-zero phase where sharers disagree.
func TestParseEnchantsMergesSharedIds(t *testing.T) {
	table := excerpt(t)

	got, ok := table[41]
	if !ok {
		t.Fatal("41 missing")
	}
	if got.Name != "Enchant Bracer - Minor Health" {
		t.Errorf("41's merged name = %q, want the first sharer's", got.Name)
	}
	wantSlots := []string{"chest", "wrist"}
	if !slices.Equal(got.Slots, wantSlots) {
		t.Errorf("41's merged slots = %v, want %v", got.Slots, wantSlots)
	}

	got, ok = table[241]
	if !ok {
		t.Fatal("241 missing")
	}
	wantTypes := []string{"normal", "two_hand"}
	if !slices.Equal(got.ItemTypes, wantTypes) {
		t.Errorf("241's merged item_types = %v, want %v", got.ItemTypes, wantTypes)
	}

	got, ok = table[2588]
	if !ok {
		t.Fatal("2588 missing")
	}
	if got.Phase != 4 {
		t.Errorf("2588's phase = %d, want 4 (its own, unshared)", got.Phase)
	}
}

// No real shared id in this build mixes a class-restricted sharer with
// an unrestricted one (all 19 shared ids in the real file agree on
// Classes), so this scenario is exercised with a synthetic fixture
// rather than a cut from enchants.json. An empty Classes means "every
// class"; merging it with a restriction must stay "every class", not
// keep the restriction, because EnchantFits must never refuse a class
// the engine would actually accept through the unrestricted sharer.
func TestParseEnchantsClassRestrictionIsAbsorbedByAnUnrestrictedSharer(t *testing.T) {
	synthetic := []byte(`[
		{"id": 90000, "name": "Restricted Sharer", "icon": "x", "slots": ["head"], "item_types": ["kit"], "classes": ["mage"], "stats": {}, "phase": 0, "spell_id": 1, "item_id": 0},
		{"id": 90000, "name": "Unrestricted Sharer", "icon": "x", "slots": ["legs"], "item_types": ["kit"], "classes": [], "stats": {}, "phase": 0, "spell_id": 2, "item_id": 0}
	]`)
	table, err := parseEnchants(synthetic)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := table[90000]
	if !ok {
		t.Fatal("90000 missing")
	}
	if len(got.Classes) != 0 {
		t.Errorf("merged Classes = %v, want empty (the unrestricted sharer should absorb the restriction)", got.Classes)
	}
	// A mage keeps fitting through the unrestricted sharer's slot.
	if err := fits(table, 90000, "legs", Item{ID: 1}, "mage"); err != nil {
		t.Errorf("fits = %v, want nil", err)
	}
	// So does a warrior: the restriction did not survive the merge.
	if err := fits(table, 90000, "legs", Item{ID: 1}, "warrior"); err != nil {
		t.Errorf("fits = %v, want nil", err)
	}
}

func TestMergeClasses(t *testing.T) {
	cases := []struct {
		name string
		a, b []string
		want []string
	}{
		{name: "both unrestricted", a: nil, b: nil, want: nil},
		{name: "one restricted, one not, absorbs to unrestricted", a: []string{"mage"}, b: nil, want: nil},
		{name: "not restricted, one is, absorbs to unrestricted", a: nil, b: []string{"warrior"}, want: nil},
		{name: "both restricted unions", a: []string{"mage"}, b: []string{"warrior"}, want: []string{"mage", "warrior"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := mergeClasses(c.a, c.b)
			if !slices.Equal(got, c.want) {
				t.Errorf("mergeClasses(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
			}
		})
	}
}

// The embedded copy is what both artifacts ship, so it has to load and
// it has to carry the slot vocabulary the planner speaks. A stale copy
// is a failing test here rather than an enchant silently refused in a
// visitor's browser.
func TestTheEmbeddedEnchantTableLoads(t *testing.T) {
	table, err := Enchants()
	if err != nil {
		t.Fatal(err)
	}
	if len(table) == 0 {
		t.Fatal("the embedded enchant table is empty; run `make simdb`")
	}
	var placeable int
	for id, e := range table {
		if e.ID != id {
			t.Errorf("enchant %d is keyed as %d", e.ID, id)
		}
		for _, slot := range e.Slots {
			if !slices.Contains(api.GearSlots, slot) {
				t.Errorf("enchant %d (%s) goes in %q, which is not a gear slot", e.ID, e.Name, slot)
			}
		}
		if len(e.Slots) > 0 {
			placeable++
		}
	}
	if placeable == 0 {
		t.Error("no enchant in the build goes anywhere")
	}
	t.Logf("enchants=%d placeable=%d", len(table), placeable)
	if _, ok := LookupEnchant(0); ok {
		t.Error("enchant 0 resolved")
	}
}

// item_types pinned to the four values this build's enchants.json
// actually carries ("staff" is in the fork's EnchantType enum but no
// row in this build uses it). itemTypesAllow treats any value it does
// not recognize as slot-only rather than a hard reject (see enchants.go),
// so a later build introducing "staff" cannot silently make a legal
// enchant unplaceable - but this test pins the known set so a new value
// showing up is noticed here rather than absorbed without comment.
func TestEnchantItemTypesAreTheFourThisBuildCarries(t *testing.T) {
	table, err := Enchants()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]bool{"normal": true, "two_hand": true, "shield": true, "kit": true}
	got := map[string]bool{}
	for _, e := range table {
		for _, itemType := range e.ItemTypes {
			got[itemType] = true
		}
	}
	for itemType := range got {
		if !want[itemType] {
			t.Errorf("enchants.json now carries item_type %q, which itemTypesAllow has never seen and "+
				"treats as slot-only; confirm that is still right, then add it to this test's want set", itemType)
		}
	}
	for itemType := range want {
		if !got[itemType] {
			t.Logf("the active build carries no enchant with item_type %q anymore", itemType)
		}
	}
}

// The error messages name the enchant and the item, because the page
// shows them and a bare "does not fit" is unactionable.
func TestEnchantErrorsSayWhy(t *testing.T) {
	err := fits(excerpt(t), 2646, "main_hand", Item{ID: 2, HandType: HandOne, WeaponType: "sword"}, "warrior")
	for _, want := range []string{"Agility", "two_hand", "one_hand"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("%q does not mention %q", err, want)
		}
	}
}
