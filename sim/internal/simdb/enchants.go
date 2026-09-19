package simdb

// The build's enchant table.
//
// It rides with the item database and for the same reason: sim/bulk has
// to know whether an enchant may go where a candidate is being sent, and
// the engine's own SimEnchant is an effect id and a stat array - the
// engine is TOLD where an enchant went rather than asked whether it
// could go there. The planner has to ask.
//
// It is embedded rather than fetched (contract A9): the browser would
// otherwise need a round trip before it could plan, and the native
// binary a file path. `make simdb` copies
// data/builds/<build>/enchants.json here beside simdb.bin; the copy is
// git-ignored and the source is committed. The data lane publishes the
// same rows for the browser UI's own enchant picker.
//
// The fork's own UIEnchant table keys ROWS by effect id, but effect id
// is not unique: 19 of this build's 150 ids are shared by two or three
// otherwise-unrelated rows (data/pipeline/models.py's EnchantRecord.id
// documents this - it repeats an effect across every slot it can go
// in). Effect 41, for instance, is both "Enchant Bracer - Minor Health"
// (wrist) and "Enchant Chest - Minor Health" (chest): two spells that
// happen to share one SpellItemEnchantment row. The engine's own
// SimEnchant carries only the effect id and a stat array, so the engine
// itself cannot tell those rows apart either - it applies whichever
// stats effect 41 carries, wherever it is cast. A merged view, one
// Enchant per id whose Slots and ItemTypes are the UNION of every
// sharer's, is therefore the faithful model of what the engine will
// actually accept: EnchantFits is a guard rail against Expand proposing
// a placement the engine would refuse, and it must never refuse one the
// engine would allow. So effect 41 legitimately fits both wrist and
// chest, and effect 241 (a "two_hand" row sharing an id with a "normal"
// one) legitimately fits both a two-hander and a one-hander.

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"sync"
)

var (
	// ErrUnknownEnchant is returned for an effect id the build's table
	// does not carry.
	ErrUnknownEnchant = errors.New("simdb: unknown enchant")
	// ErrEnchantDoesNotFit is returned when the enchant exists and
	// cannot go on that item in that slot.
	ErrEnchantDoesNotFit = errors.New("simdb: the enchant does not fit")
)

//go:embed enchants.json
var rawEnchants []byte

// The two item_types values enchants.json's ENCHANT_TYPES vocabulary
// carries that are real restrictions on the item rather than the slot
// alone. "two_hand" happens to spell the same as HandTwo, and "shield"
// happens to spell the same as the WeaponType this package's Item.
// carries for a shield; that coincidence is the fork's vocabulary, not
// this package's choice.
const (
	itemTypeTwoHand = "two_hand"
	itemTypeShield  = "shield"
)

// Enchant is one row of the build's table, merged across every
// SpellItemEnchantment row that shares its effect id (see the package
// doc comment).
type Enchant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	// Slots are the gear slots this enchant may be applied in: the
	// union of every sharer's slots.
	Slots []string `json:"slots"`
	// ItemTypes narrows it further where the slot is not enough: the
	// union of every sharer's item_types. "two_hand" and "shield" are
	// real restrictions (see EnchantFits); every other value - "normal"
	// and "kit" in this build, and any value a later build introduces
	// that this package does not yet know - is slot-only and imposes
	// no further restriction, since the slots check above already says
	// where it goes.
	ItemTypes []string `json:"item_types"`
	// Classes are the class slugs allowed to use it: the union of
	// every sharer's classes. Empty means every class.
	Classes []string           `json:"classes"`
	Stats   map[string]float64 `json:"stats"`
	// Phase is the lowest non-zero phase among its sharers, or 0 if
	// every sharer is phase 0.
	Phase int `json:"phase"`
}

// enchantRow is one line of enchants.json, before rows sharing an
// effect id are merged into one Enchant. spell_id and item_id
// distinguish sharers in the source data but are not carried into
// Enchant: sim/bulk asks whether an effect id fits somewhere, never
// which spell cast it.
type enchantRow struct {
	ID        int                `json:"id"`
	Name      string             `json:"name"`
	Icon      string             `json:"icon"`
	Slots     []string           `json:"slots"`
	ItemTypes []string           `json:"item_types"`
	Classes   []string           `json:"classes"`
	Stats     map[string]float64 `json:"stats"`
	Phase     int                `json:"phase"`
}

// parseEnchants decodes and merges the table. It is separate from the
// embedded load so the tests can read a small fixture through the same
// code.
func parseEnchants(b []byte) (map[int]Enchant, error) {
	var rows []enchantRow
	if err := json.Unmarshal(b, &rows); err != nil {
		return nil, fmt.Errorf("simdb: the enchant table is not readable: %w", err)
	}
	if len(rows) == 0 {
		return nil, errors.New("simdb: the enchant table carries no enchants")
	}
	out := make(map[int]Enchant, len(rows))
	for _, row := range rows {
		if row.ID == 0 {
			return nil, fmt.Errorf("simdb: the enchant table has a row with no id: %+v", row)
		}
		existing, ok := out[row.ID]
		if !ok {
			// The first sharer in file order supplies Name, Icon and
			// Stats; a later sharer only widens Slots, ItemTypes,
			// Classes and lowers Phase.
			out[row.ID] = Enchant{
				ID:        row.ID,
				Name:      row.Name,
				Icon:      row.Icon,
				Slots:     unique(row.Slots),
				ItemTypes: unique(row.ItemTypes),
				Classes:   unique(row.Classes),
				Stats:     row.Stats,
				Phase:     row.Phase,
			}
			continue
		}
		existing.Slots = union(existing.Slots, row.Slots)
		existing.ItemTypes = union(existing.ItemTypes, row.ItemTypes)
		existing.Classes = union(existing.Classes, row.Classes)
		existing.Phase = lowestNonZero(existing.Phase, row.Phase)
		out[row.ID] = existing
	}
	return out, nil
}

// unique is a sorted copy of ss with duplicates removed, so a single
// row's own repeats (none observed, but not assumed) sort the same way
// a merge's union does.
func unique(ss []string) []string {
	return union(nil, ss)
}

// union is the sorted, deduplicated union of a and b.
func union(a, b []string) []string {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}
	out := make([]string, 0, len(a)+len(b))
	out = append(out, a...)
	out = append(out, b...)
	slices.Sort(out)
	return slices.Compact(out)
}

// lowestNonZero is the lower of a and b, treating 0 as "no opinion"
// rather than as a value: two sharers at phase 0 and 4 merge to 4, not
// 0.
func lowestNonZero(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 || b > a {
		return a
	}
	return b
}

var enchantTable = sync.OnceValues(func() (map[int]Enchant, error) {
	return parseEnchants(rawEnchants)
})

// Enchants is the whole table of the active build.
func Enchants() (map[int]Enchant, error) { return enchantTable() }

// LookupEnchant is one enchant.
func LookupEnchant(id int) (Enchant, bool) {
	table, err := enchantTable()
	if err != nil {
		return Enchant{}, false
	}
	e, ok := table[id]
	return e, ok
}

// EnchantFits reports whether an enchant may go on this item in this
// slot for this class, and says why not when it may not.
func EnchantFits(effectID int, slot string, item Item, class string) error {
	table, err := enchantTable()
	if err != nil {
		return err
	}
	return fits(table, effectID, slot, item, class)
}

// fits is EnchantFits over a given table, so the rules are testable
// against the fixture as well as against the build.
func fits(table map[int]Enchant, effectID int, slot string, item Item, class string) error {
	row, ok := table[effectID]
	if !ok {
		return fmt.Errorf("%w: %d", ErrUnknownEnchant, effectID)
	}
	if !slices.Contains(row.Slots, slot) {
		return fmt.Errorf("%w: %q goes in %v, not %q", ErrEnchantDoesNotFit, row.Name, row.Slots, slot)
	}
	if len(row.Classes) > 0 && !slices.Contains(row.Classes, class) {
		return fmt.Errorf("%w: %q is for %v, not a %s", ErrEnchantDoesNotFit, row.Name, row.Classes, class)
	}
	if !itemTypesAllow(row.ItemTypes, item) {
		return fmt.Errorf("%w: %q goes on %v, and item %d is a %s %s", ErrEnchantDoesNotFit, row.Name, row.ItemTypes, item.ID, item.HandType, item.WeaponType)
	}
	return nil
}

// itemTypesAllow reports whether item satisfies at least one of
// itemTypes. Empty is unrestricted. "two_hand" and "shield" are real
// restrictions on the item; every other value - including one this
// build has never produced - is slot-only and always satisfied, so a
// later build's new EnchantType cannot silently make a legal enchant
// unplaceable.
func itemTypesAllow(itemTypes []string, item Item) bool {
	if len(itemTypes) == 0 {
		return true
	}
	for _, t := range itemTypes {
		switch t {
		case itemTypeTwoHand:
			if item.HandType == HandTwo {
				return true
			}
		case itemTypeShield:
			if item.WeaponType == itemTypeShield {
				return true
			}
		default:
			return true
		}
	}
	return false
}
