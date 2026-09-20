package bulk

// Expansion: which substitutions are even possible.
//
// The rules are the ones the page states in words on the Top Gear
// page, and each has a table test: an item goes only where its
// inventory type allows; rings and trinkets are tried in both slots;
// a one-hander is tried in either hand; a class, level or faction the
// character does not have means the item is never simmed; a locked
// slot is never substituted; and an enchant is inherited from the
// equipped item where it fits.
//
// A candidate the character cannot equip is SKIPPED rather than
// refused. Droptimizer sends a boss's whole loot table, plate and
// cloth together, and refusing the request would make the tool
// unusable; the page shows the count it planned, so a skipped
// candidate is visible rather than silent.
//
// A candidate that is malformed - an item the build has never heard
// of, a named enchant that cannot go where it was sent - IS refused.
// That is the page sending something it should not, and running it
// would answer a different question from the one asked.

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

var (
	// ErrNotBulk is returned when a request with no bulk block is
	// handed to the planner.
	ErrNotBulk = errors.New("bulk: the request carries no bulk block")
	// ErrUnknownItem is returned for a candidate the active build has
	// no row for. It is a refusal rather than a skip: an id nothing
	// resolves is a client bug, and the engine would die mid-run on it.
	ErrUnknownItem = errors.New("bulk: the build has no such item")
	// ErrInvalidSet is returned for a named gear set that is not valid
	// equipment on its own - a two-hander beside an off-hand, an item
	// worn twice. It is a refusal rather than a skip: a set the player
	// named and will look for in the results is exactly the malformed
	// input the package doc says is refused, and valid() quietly
	// dropping every combination that arm produces would leave the set
	// missing with no explanation.
	ErrInvalidSet = errors.New("bulk: a gear set is not valid equipment")
)

// Combination is one substitution set and the request that runs it.
type Combination struct {
	Request       api.SimRequest     `json:"request"`
	Substitutions []api.Substitution `json:"substitutions"`
}

// placement is one candidate in one slot: the gear row it produces and
// the substitution chip that describes it.
type placement struct {
	Slot         string
	Gear         api.GearSlot
	Substitution api.Substitution
}

// Expand lists every valid combination for req, with no build tables.
// See ExpandWith.
func Expand(req api.SimRequest) ([]Combination, error) {
	return ExpandWith(req, Options{})
}

// ExpandWith lists every valid combination for req.
func ExpandWith(req api.SimRequest, opt Options) ([]Combination, error) {
	if req.Bulk == nil {
		return nil, ErrNotBulk
	}
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("bulk: %w", err)
	}
	if err := validSets(req.Bulk.Sets); err != nil {
		return nil, err
	}
	places, err := placements(req, opt)
	if err != nil {
		return nil, err
	}
	return combinations(req, places)
}

// validSets refuses a named gear set that valid() would reject on its
// own. apply's two-hand-clears-the-off-hand rule (below) only runs on
// candidate placements, never on a set's own gear list - a set is a
// whole loadout the page sent as-is - so an internally invalid set
// would otherwise reach combinations, have every arm it produces
// silently deleted by the validity filter, and vanish from the result
// with nothing to say why.
func validSets(sets []api.GearSet) error {
	for _, set := range sets {
		if !valid(set.Gear) {
			return fmt.Errorf("%w: %q", ErrInvalidSet, set.Name)
		}
	}
	return nil
}

// baseGear indexes the character's equipped gear by slot.
func baseGear(req api.SimRequest) map[string]api.GearSlot {
	out := make(map[string]api.GearSlot, len(req.Character.Gear))
	for _, g := range req.Character.Gear {
		out[g.Slot] = g
	}
	return out
}

// placements turns every candidate into the (slot, gear) pairs it can
// produce. A candidate that fits nowhere contributes none.
//
// It asks simdb.Ready first, once, because every lookup below - and
// every lookup valid() does later - reads a table that failed to load
// as a table of misses. Without this, a corrupt embed would refuse
// the first candidate as ErrUnknownItem ("the build has no such
// item"), which is a true statement about the wrong thing and sends
// whoever reads it looking for a bad item id.
func placements(req api.SimRequest, opt Options) ([]placement, error) {
	if err := simdb.Ready(); err != nil {
		return nil, fmt.Errorf("bulk: the build's item table is unusable, so no candidate can be checked: %w", err)
	}
	equipped := baseGear(req)
	locked := req.Bulk.Locked
	var out []placement
	for _, c := range req.Bulk.Candidates {
		item, ok := simdb.Lookup(c.ItemID)
		if !ok {
			return nil, fmt.Errorf("%w: %d", ErrUnknownItem, c.ItemID)
		}
		if !usable(item, req.Character) {
			continue
		}
		for _, slot := range item.Slots {
			if c.Slot != "" && c.Slot != slot {
				continue
			}
			if slices.Contains(locked, slot) {
				continue
			}
			enchant, err := enchantFor(c, slot, item, req.Character.Class, equipped, opt)
			if err != nil {
				return nil, err
			}
			out = append(out, placement{
				Slot: slot,
				Gear: api.GearSlot{Slot: slot, ItemID: c.ItemID, Enchant: enchant, Suffix: c.Suffix},
				Substitution: api.Substitution{
					Kind:    api.SubstitutionItem,
					Slot:    slot,
					ItemID:  c.ItemID,
					Enchant: enchant,
					Suffix:  c.Suffix,
					Origin:  c.Origin,
					// The item's own name, from the build. The API
					// composes a headline - "+41 DPS from Vis'kag" -
					// at save time and has no item table; filling it
					// here is what lets a stored row still say what
					// it was about years later.
					Name:       item.Name,
					SourceName: c.SourceName,
				},
			})
		}
	}
	return out, nil
}

// usable reports whether this character could wear the item at all.
// Class, level and faction are three ways of saying "not for you", and
// all three end the same way: the item is never simmed.
func usable(item simdb.Item, ch api.CharacterSpec) bool {
	if len(item.Classes) > 0 && !slices.Contains(item.Classes, ch.Class) {
		return false
	}
	if item.RequiredLevel > ch.Level {
		return false
	}
	if item.Faction != simdb.FactionAny && item.Faction != factionOf(ch.Race) {
		return false
	}
	return true
}

// hordeRaces are the races that cannot loot an Alliance-only item. The
// list is short and stable - Forever's Skyborne choose a faction at
// creation and the race slug says which - so it lives here rather than
// pulling the build's race table into the planner.
//
// Keep in sync with data/builds/<build>/races.json, the pipeline's own
// race-to-faction table: a race added there without a matching entry
// here defaults to Alliance below, which drops every Horde-only
// candidate for that race from the plan with no error.
var hordeRaces = []string{"orc", "tauren", "troll", "undead", "windshaper-skyborne"}

// factionOf is the faction a race belongs to.
func factionOf(race string) string {
	if slices.Contains(hordeRaces, race) {
		return simdb.FactionHorde
	}
	return simdb.FactionAlliance
}

// enchantFor resolves a candidate's enchant.
//
// A named enchant is checked and kept, or REFUSED: the player asked
// for that enchant and would otherwise read a ranking of something
// else.
//
// An unnamed one INHERITS the equipped item's for that slot, where
// that enchant fits the new item - which is what Raidbots does and
// what the fork's auto_enchant does - and is DROPPED where it does
// not, because a two-hand enchant cannot ride onto a one-hander and
// the enchant came from the character rather than from the request.
//
// There is no "no table" case: the build's enchants are embedded
// beside its items (contract A9), so Options.enchants() always
// answers.
func enchantFor(c api.Candidate, slot string, item simdb.Item, class string, equipped map[string]api.GearSlot, opt Options) (int, error) {
	table := opt.enchants()
	if c.Enchant != 0 {
		if err := table.Fits(c.Enchant, slot, item, class); err != nil {
			return 0, err
		}
		return c.Enchant, nil
	}
	inherited := equipped[slot].Enchant
	if inherited == 0 {
		return 0, nil
	}
	if err := table.Fits(inherited, slot, item, class); err != nil {
		return 0, nil
	}
	return inherited, nil
}

// ExpandWorkBudget is how many combinations Expand and Count will
// ENUMERATE before they stop and refuse the request outright.
//
// Retention is bounded by the cap (see combinations, below), but
// retention was never the whole cost: apply builds a complete
// api.SimRequest for every member of the cross product before the
// retention decision is taken, so the WORK is proportional to the
// product no matter how little of it is kept. Nine armour slots with
// five real bag candidates apiece is 6^9 = 10,077,696 gear shapes and
// took 23 seconds on the machine this was measured on; ten apiece is
// 11^9 = 2.4 billion, which is hours. Those are ordinary Top Gear
// inputs - a player's bags and bank - and contract 10.2 has the page
// calling simCount on every candidate tick, single-threaded, on the
// browser's wasm heap. An unbounded count there is a frozen tab.
//
// 250,000 is chosen to sit far above any lane cap (the largest is
// api.Caps[api.LaneServer] = 5,000, fifty times smaller) and far below
// anything that takes a perceptible time to enumerate: it is roughly
// half a second of the measurement above. Reaching it therefore
// already proves the cap is breached, so the REFUSAL is exact even
// though the number quoted with it becomes an arithmetic upper bound
// (upperBound, below) rather than the enumerated count. A request that
// finishes enumerating under the budget still reports the precise
// count it always did.
//
// The one shape that trade gets wrong is a request whose validity
// filter rejects nearly everything - twenty-five identically named
// rings offered to both finger slots, twenty-five identically named
// trinkets to both trinket slots, which enumerates 457,000 shapes of
// which only 2,601 are valid. That is refused here and would have
// fitted. It is not a shape any of the three tools builds, and the
// alternative - enumerating without a bound so that case can be
// answered - is the freeze this constant exists to prevent.
const ExpandWorkBudget = 250_000

// combinations builds every combination the mode allows, retaining at
// most Cap+1 of them - enough to prove a cap breach without holding
// the whole product in memory at once - and enumerating at most
// ExpandWorkBudget of them, which is what bounds the TIME as well.
//
// Gear mode is a product: each slot offers "keep what is equipped" or
// one of its candidates, and the talent dimension offers "the
// character's own build" or one of the loadouts. A set is not part of
// that product - it replaces every slot at once - so each set is its
// own arm, crossed with the talent dimension only.
//
// Drops and talents mode are one substitution at a time: a Droptimizer
// answer is "this boss's sword is worth 41 DPS", and a product would
// answer a question nobody asked and blow the cap doing it.
//
// The equipped set is never a combination: Plan runs it separately, in
// every stage, so that every delta is paired.
//
// gearCombinations and singleCombinations deliver one combination at a
// time to keep, rather than building a slice: the talent and
// consumables dimensions alone multiply (a 400-loadout, 2,000-list
// gear request is 802,400 combinations before validity even runs), and
// so does the gear dimension on its own (nine single-item slots with
// three candidates apiece is 262,144 gear shapes before the other two
// dimensions touch it - see walkGearChoices, which walks that product
// lazily rather than building it). Contract 10.2 has the page calling
// Count on every candidate tick, in the browser, on a 32-bit wasm heap.
// keep runs valid() and counts every one of them - the reported
// Combinations is exact for any request that enumerates within the
// budget, and Expand/Count can never disagree with each other - but
// stops retaining once it already has enough to answer the cap, and
// stops enumerating altogether once the budget is spent. Retained
// memory is O(Cap); nothing proportional to any dimension's full
// product - gear, talents or consumables - is ever resident at once,
// and no enumeration outlives the budget.
//
// keep returns whether enumeration should continue, and every walker
// below threads that answer back out, so the budget is a real early
// exit rather than a count checked after the fact.
func combinations(req api.SimRequest, places []placement) ([]Combination, error) {
	// Non-nil from the start: an expansion with no valid combinations is
	// a real, empty result, not the absence of one, and should marshal
	// as [] rather than null once a caller across the wasm boundary is
	// reading this JSON.
	out := []Combination{}
	// total counts the VALID combinations, which is the number the cap
	// is about; seen counts every one enumerated, valid or not, which
	// is what the budget is about - an invalid combination costs a
	// whole apply() to discover.
	var total, seen int
	keep := func(c Combination) bool {
		seen++
		if valid(c.Request.Character.Gear) {
			total++
			if len(out) <= req.Bulk.Cap {
				out = append(out, c)
			}
			// retentionProbe exists only so a test can observe len(out)
			// DURING enumeration, not after: on a cap breach out is
			// discarded below in favor of the error, so a test that only
			// inspects the return value can never tell retained-to-the-cap
			// apart from retained-everything-then-thrown-away - both return
			// (nil, ErrCapExceeded{...}) either way.
			if retentionProbe != nil {
				retentionProbe(len(out))
			}
		}
		return seen < ExpandWorkBudget
	}
	var finished bool
	if req.Bulk.Mode == api.KindGear {
		finished = gearCombinations(req, places, keep)
	} else {
		finished = singleCombinations(req, places, keep)
	}
	if !finished {
		// The budget stopped the walk, so total is a partial count and
		// quoting it would understate the plan. upperBound is the
		// arithmetic size of the product, which is what the page shows
		// beside the cap.
		return nil, api.ErrCapExceeded{Cap: req.Bulk.Cap, Combinations: upperBound(req, places)}
	}
	if total > req.Bulk.Cap {
		return nil, api.ErrCapExceeded{Cap: req.Bulk.Cap, Combinations: total}
	}
	return out, nil
}

// upperBound is how large the expansion could be, by arithmetic
// instead of by enumeration: the product of each slot's candidates
// plus one ("keep what is equipped"), times the talent and consumable
// dimensions, plus the sets arm. It is O(number of slots) to compute
// and is only ever read when the work budget stopped the walk.
//
// It is an UPPER bound, never the answer: the validity rules - a
// two-hander beside an off-hand, one physical one-hander in both
// hands, a unique-equipped item worn twice - can only be answered by
// looking at a finished gear list, and every one of them can only
// remove combinations from this figure. Erring high is the right way
// round for a refusal that is already certain.
//
// The multiplications saturate rather than wrap: seventeen slots with
// twenty candidates apiece is 21^17, which is far past what an int
// holds, and a wrapped product could come back small enough to look
// like it fits.
func upperBound(req api.SimRequest, places []placement) int {
	if req.Bulk.Mode != api.KindGear {
		// One substitution at a time, so the arithmetic is exact: every
		// placement on its own, then every loadout on its own. This is
		// singleCombinations' own shape, written out.
		return len(places) + len(req.Bulk.Talents)
	}
	// The talent and consumables dimensions each offer "the character's
	// own" plus their alternatives, the way gearCombinations builds
	// them.
	dims := satMul(len(req.Bulk.Talents)+1, len(req.Bulk.Consumables)+1)
	perSlot := map[string]int{}
	for _, p := range places {
		perSlot[p.Slot]++
	}
	product := 1
	for _, n := range perSlot {
		product = satMul(product, n+1)
	}
	// A named set is its own arm, crossed with the other two dimensions
	// only - not a member of the gear product.
	return satAdd(satMul(product, dims), satMul(len(req.Bulk.Sets), dims))
}

// satMul multiplies without wrapping: an overflow saturates at
// math.MaxInt.
func satMul(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a > math.MaxInt/b {
		return math.MaxInt
	}
	return a * b
}

// satAdd adds without wrapping, for the same reason.
func satAdd(a, b int) int {
	if a > math.MaxInt-b {
		return math.MaxInt
	}
	return a + b
}

// retentionProbe is nil in production. A test sets it (and restores it
// to nil when done) to watch combinations' retained count as it grows,
// which is the only way to prove the len(out) <= req.Bulk.Cap guard is
// doing anything: on a cap breach, the return value alone cannot tell
// "retained at most Cap+1 throughout" apart from "retained everything,
// then discarded it for the error" - both produce the identical
// (nil, ErrCapExceeded{...}).
var retentionProbe func(retained int)

// singleCombinations is one substitution at a time: every placement on
// its own, then every talent loadout on its own. It reports whether it
// got all the way through, or stopped because keep said the work
// budget was spent.
func singleCombinations(req api.SimRequest, places []placement, keep func(Combination) bool) bool {
	for _, p := range places {
		if !keep(apply(req, []placement{p}, nil, nil, nil)) {
			return false
		}
	}
	for i := range req.Bulk.Talents {
		if !keep(apply(req, nil, &req.Bulk.Talents[i], nil, nil)) {
			return false
		}
	}
	return true
}

// gearCombinations is the product. Like singleCombinations it reports
// whether it finished rather than stopping on the work budget.
func gearCombinations(req api.SimRequest, places []placement, keep func(Combination) bool) bool {
	// One bucket per slot, in the envelope's slot order so the product
	// is enumerated the same way every time and two runs of the same
	// request produce the same combination order.
	bySlot := map[string][]placement{}
	for _, p := range places {
		bySlot[p.Slot] = append(bySlot[p.Slot], p)
	}
	slots := make([]string, 0, len(bySlot))
	for _, slot := range api.GearSlots {
		if len(bySlot[slot]) > 0 {
			slots = append(slots, slot)
		}
	}

	// The talent dimension: the character's own build, then each
	// loadout. A nil loadout is "their own".
	loadouts := make([]*api.TalentLoadout, 0, len(req.Bulk.Talents)+1)
	loadouts = append(loadouts, nil)
	for i := range req.Bulk.Talents {
		loadouts = append(loadouts, &req.Bulk.Talents[i])
	}
	// The consumable dimension, the same shape: the character's own
	// list, then each alternative. A nil choice is "their own".
	drinks := make([]*consumableChoice, 0, len(req.Bulk.Consumables)+1)
	drinks = append(drinks, nil)
	for i, list := range req.Bulk.Consumables {
		drinks = append(drinks, &consumableChoice{Index: i, List: list})
	}

	// Every choice of at most one placement per slot, including none -
	// walked lazily (see walkGearChoices) rather than built as a slice,
	// crossed with the talent and consumables dimensions as each one is
	// produced.
	if !walkGearChoices(slots, bySlot, func(chosen []placement) bool {
		for _, loadout := range loadouts {
			for _, drink := range drinks {
				if len(chosen) == 0 && loadout == nil && drink == nil {
					// The equipped set with its own talents and its
					// own consumables is the baseline, and Plan runs
					// it. Including it here would rank the character
					// against itself.
					continue
				}
				if !keep(apply(req, chosen, loadout, nil, drink)) {
					return false
				}
			}
		}
		return true
	}) {
		return false
	}
	// A named set is a whole-gear alternative, so it is its own arm
	// rather than a member of the product above.
	for i := range req.Bulk.Sets {
		for _, loadout := range loadouts {
			for _, drink := range drinks {
				if !keep(apply(req, nil, loadout, &req.Bulk.Sets[i], drink)) {
					return false
				}
			}
		}
	}
	return true
}

// walkGearChoices calls yield once for every choice of at most one
// placement per slot - "keep what is equipped" or one of a slot's
// candidates - crossed across every slot in slots, in the same order
// the eager version of this product used to build as a
// [][]placement: slots earlier in the list vary slower than slots
// later in it, and within one slot "keep" is tried before its
// candidates.
//
// It walks that product depth-first, one slot at a time, rather than
// building it as a slice: nine single-item slots with three candidates
// apiece is already 4^9 = 262,144 gear shapes before the talent and
// consumables dimensions even multiply it further, and a request with
// that many real, individually unremarkable bag candidates passes
// validation. Holding the whole cross product, even just as
// []placement slices with no requests built yet, is exactly the
// O(product) memory the cap exists to keep off the browser's wasm
// heap; walking it recursively keeps at most one partial choice per
// slot on the call stack, which is O(len(slots)).
//
// A one-hander with no slot named expands into two placements, one per
// hand, for this walk to choose between; yielding a chosen that picked
// BOTH would wield one physical item in two slots at once, so
// sameWeaponTwice filters at the leaf - the same place the eager
// version filtered the fully-built product, and the only place a
// chosen combination is complete enough to check. A leaf filtered
// there costs no apply() and so spends no work budget; it cannot run
// away with the walk either, because a filtered leaf needs the same
// item chosen in BOTH hands and those are at most a quarter of any
// product that has them at all.
//
// yield reports whether to carry on. Returning false unwinds the whole
// recursion immediately, which is what lets ExpandWorkBudget bound the
// TIME of an enumeration and not merely its memory; walkGearChoices
// passes that answer back to its own caller.
func walkGearChoices(slots []string, bySlot map[string][]placement, yield func([]placement) bool) bool {
	var walk func(i int, chosen []placement) bool
	walk = func(i int, chosen []placement) bool {
		if i == len(slots) {
			if sameWeaponTwice(chosen) {
				return true
			}
			return yield(chosen)
		}
		slot := slots[i]
		if !walk(i+1, chosen) {
			return false
		}
		for _, p := range bySlot[slot] {
			if !walk(i+1, append(append([]placement(nil), chosen...), p)) {
				return false
			}
		}
		return true
	}
	return walk(0, nil)
}

// sameWeaponTwice reports whether chosen substitutes the identical
// item into both main_hand and off_hand within the SAME combination.
// It only looks at what this combination actively substitutes, not at
// gear that is left equipped and unchanged: a candidate that happens
// to share an id with the character's own equipped weapon is a
// placement-eligibility question, answered in placements, not a
// weapon-shape question answered here.
//
// It compares item and suffix, not the resolved enchant: a one-hander
// with no slot named produces one placement per hand from the SAME
// candidate, and each independently inherits whatever enchant was
// already in that slot (contract: enchantFor), so the two placements'
// enchants can legitimately differ even though it is one physical
// item being offered for both hands.
func sameWeaponTwice(chosen []placement) bool {
	var main, off *placement
	for i := range chosen {
		switch chosen[i].Slot {
		case "main_hand":
			main = &chosen[i]
		case "off_hand":
			off = &chosen[i]
		}
	}
	if main == nil || off == nil {
		return false
	}
	return main.Gear.ItemID == off.Gear.ItemID && main.Gear.Suffix == off.Gear.Suffix
}

// valid is the engine's own isValidEquipment, over our gear list: no
// two-hander beside an off-hand, no item in both ring or both trinket
// slots, no two rings or trinkets sharing a name (which is the same
// item at two qualities), and nothing unique-equipped worn twice. It
// is not the whole of the product's shape rules on its own -
// sameWeaponTwice, just above, catches one physical one-hander being
// offered for both hands within a single combination, which valid
// cannot see because it only ever looks at one combination's finished
// gear list, never at which placements were chosen together to build
// it.
//
// It is re-expressed here rather than called because the engine's copy
// works on a protobuf and this package holds no protobuf, and because
// a combination that reached the engine and was refused there would
// cost a whole sim to learn.
func valid(gear []api.GearSlot) bool {
	byID := map[int]int{}
	names := map[string]int{}
	var mainHand, offHand simdb.Item
	for _, g := range gear {
		item, ok := simdb.Lookup(g.ItemID)
		if !ok {
			continue
		}
		byID[g.ItemID]++
		if byID[g.ItemID] > 1 && item.Unique {
			return false
		}
		switch g.Slot {
		case "main_hand":
			mainHand = item
		case "off_hand":
			offHand = item
		case "finger1", "finger2", "trinket1", "trinket2":
			// One item cannot be in both of a pair, and two items of
			// one name are the same thing at two qualities.
			if byID[g.ItemID] > 1 {
				return false
			}
			names[item.Name]++
			if names[item.Name] > 1 {
				return false
			}
		}
	}
	return !(mainHand.HandType == simdb.HandTwo && offHand.ID != 0)
}

// consumableChoice is one alternative consumable list from
// BulkSpec.Consumables, with the position the page put it in so the
// chip is stable across runs.
type consumableChoice struct {
	Index int
	List  []string
}

// apply builds one combination: the base character with these
// placements swapped in, optionally a talent loadout, optionally a
// whole gear set, optionally an alternative consumable list, and the
// chips that describe the change.
//
// The request it produces is a PLAIN RUN - no bulk block - because
// that is what a stage request is, and leaving the block on would make
// every stage request expand again.
func apply(req api.SimRequest, places []placement, loadout *api.TalentLoadout, set *api.GearSet, consumes *consumableChoice) Combination {
	out := req
	out.Bulk = nil
	out.Weights = nil
	out.TargetError = 0

	gear := baseGear(req)
	if set != nil {
		gear = make(map[string]api.GearSlot, len(set.Gear))
		for _, g := range set.Gear {
			gear[g.Slot] = g
		}
	}
	for _, p := range places {
		gear[p.Slot] = p.Gear
		// A two-hander leaves no room for an off-hand. Clearing it is
		// what makes "two-hand versus main-plus-off-hand" two competing
		// shapes rather than one combination the engine would refuse.
		//
		// This walks places in the order sameWeaponTwice's caller built
		// them, which follows api.GearSlots - main_hand ahead of
		// off_hand - so an off-hand placement chosen alongside a
		// two-hander is always set into gear AFTER this delete runs,
		// leaving it in gear so it gets caught by valid()'s own two-hand
		// check (above, in this file) instead of silently vanishing
		// here. The subs slice built further down still appends every
		// placement's chip unconditionally, though, so if api.GearSlots
		// ever put off_hand ahead of main_hand, this delete would start
		// firing AFTER an off-hand placement was set, removing it from
		// gear while its Substitution chip survived into subs - a
		// validated combination whose own chip names gear it does not
		// carry.
		if p.Slot == "main_hand" {
			if item, ok := simdb.Lookup(p.Gear.ItemID); ok && item.HandType == simdb.HandTwo {
				delete(gear, "off_hand")
			}
		}
	}
	// In the envelope's own slot order, so two combinations that equip
	// the same set produce byte-identical requests and the seeds line
	// up across stages.
	out.Character.Gear = make([]api.GearSlot, 0, len(gear))
	for _, slot := range api.GearSlots {
		if g, ok := gear[slot]; ok {
			out.Character.Gear = append(out.Character.Gear, g)
		}
	}

	subs := make([]api.Substitution, 0, len(places)+2)
	if set != nil {
		subs = append(subs, api.Substitution{Kind: api.SubstitutionSet, Name: set.Name})
	}
	for _, p := range places {
		subs = append(subs, p.Substitution)
	}
	if loadout != nil {
		out.Character.Talents = loadout.Talents
		subs = append(subs, api.Substitution{
			Kind:    api.SubstitutionTalents,
			Name:    loadout.Name,
			Talents: loadout.Talents,
		})
	}
	if consumes != nil {
		// REPLACES the character's own list, rather than adding to
		// it: "flask or two elixirs" is a choice, and merging them
		// would sim a character drinking both.
		out.Character.Consumes = slices.Clone(consumes.List)
		// The ids joined by ", " (contract 10.8). Every other kind of
		// substitution puts its label in Name, and the API composes a
		// headline from that one field without knowing what kind it
		// is reading - so an empty list, which validate() blesses as
		// "no consumables, a real thing to compare against," still
		// needs a label rather than joining to "".
		name := strings.Join(consumes.List, ", ")
		if name == "" {
			name = api.NoConsumablesLabel
		}
		subs = append(subs, api.Substitution{
			Kind:     api.SubstitutionConsumes,
			Name:     name,
			Consumes: slices.Clone(consumes.List),
		})
	}
	return Combination{Request: out, Substitutions: subs}
}

// Count is how many combinations req would expand to.
//
// It expands and counts rather than doing the arithmetic a second
// way: the validity rules - a two-hander beside an off-hand, a
// unique-equipped item worn twice - can only be answered by looking
// at the gear, so a closed-form count would be a different number
// from the one Plan runs, which is the one thing a count must never
// be. combinations enumerates the product one combination at a time
// and keeps at most Cap+1 of them, and it stops enumerating at
// ExpandWorkBudget (see both doc comments), so Count costs at most
// the budget's worth of work and the cap's worth of memory whatever
// the candidates describe. That - not "it allocates nothing", which
// was never true, since every combination is a built api.SimRequest
// by the time validity can be asked about it - is what makes it safe
// to call on every candidate tick (contract 10.2), including from
// the browser's wasm heap.
//
// It answers the cap the same way Expand does, so "1,280 against a
// cap of 400" is one message from one place. Past the budget both
// still refuse, and both quote the same arithmetic upper bound.
func Count(req api.SimRequest) (int, error) { return CountWith(req, Options{}) }

// CountWith is Count with the planner's options.
func CountWith(req api.SimRequest, opt Options) (int, error) {
	combos, err := ExpandWith(req, opt)
	if err != nil {
		return 0, err
	}
	return len(combos), nil
}
