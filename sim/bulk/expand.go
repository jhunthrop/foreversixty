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
	places, err := placements(req, opt)
	if err != nil {
		return nil, err
	}
	return combinations(req, places)
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
func placements(req api.SimRequest, opt Options) ([]placement, error) {
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

// combinations builds every combination the mode allows.
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
func combinations(req api.SimRequest, places []placement) ([]Combination, error) {
	var out []Combination
	if req.Bulk.Mode == api.KindGear {
		out = gearCombinations(req, places)
	} else {
		out = singleCombinations(req, places)
	}
	out = slices.DeleteFunc(out, func(c Combination) bool {
		return !valid(c.Request.Character.Gear)
	})
	if len(out) > req.Bulk.Cap {
		return nil, api.ErrCapExceeded{Cap: req.Bulk.Cap, Combinations: len(out)}
	}
	return out, nil
}

// singleCombinations is one substitution at a time: every placement on
// its own, then every talent loadout on its own.
func singleCombinations(req api.SimRequest, places []placement) []Combination {
	out := make([]Combination, 0, len(places)+len(req.Bulk.Talents))
	for _, p := range places {
		out = append(out, apply(req, []placement{p}, nil, nil, nil))
	}
	for i := range req.Bulk.Talents {
		out = append(out, apply(req, nil, &req.Bulk.Talents[i], nil, nil))
	}
	return out
}

// gearCombinations is the product.
func gearCombinations(req api.SimRequest, places []placement) []Combination {
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

	// Every choice of at most one placement per slot, including none.
	sets := [][]placement{nil}
	for _, slot := range slots {
		next := make([][]placement, 0, len(sets)*(len(bySlot[slot])+1))
		for _, chosen := range sets {
			next = append(next, chosen)
			for _, p := range bySlot[slot] {
				next = append(next, append(append([]placement(nil), chosen...), p))
			}
		}
		sets = next
	}
	// A one-hander with no slot named expands into two placements, one
	// per hand, for the product above to choose between; choosing BOTH
	// in one combination would wield one physical item in two slots at
	// once. Two DIFFERENT one-handers dual-wielding each other is
	// unaffected, since their items differ.
	sets = slices.DeleteFunc(sets, sameWeaponTwice)

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

	out := make([]Combination, 0, len(sets)*len(loadouts)*len(drinks))
	for _, chosen := range sets {
		for _, loadout := range loadouts {
			for _, drink := range drinks {
				if len(chosen) == 0 && loadout == nil && drink == nil {
					// The equipped set with its own talents and its
					// own consumables is the baseline, and Plan runs
					// it. Including it here would rank the character
					// against itself.
					continue
				}
				out = append(out, apply(req, chosen, loadout, nil, drink))
			}
		}
	}
	// A named set is a whole-gear alternative, so it is its own arm
	// rather than a member of the product above.
	for i := range req.Bulk.Sets {
		for _, loadout := range loadouts {
			for _, drink := range drinks {
				out = append(out, apply(req, nil, loadout, &req.Bulk.Sets[i], drink))
			}
		}
	}
	return out
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
// item at two qualities), and nothing unique-equipped worn twice.
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
		subs = append(subs, api.Substitution{
			Kind: api.SubstitutionConsumes,
			// The ids joined by ", " (contract 10.8). Every other
			// kind of substitution puts its label in Name, and the
			// API composes a headline from that one field without
			// knowing what kind it is reading.
			Name:     strings.Join(consumes.List, ", "),
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
// be. What it skips is allocating the stage's REQUESTS, which is the
// expensive half and the reason the page can ask on every tick.
//
// It answers the cap the same way Expand does, so "1,280 against a
// cap of 400" is one message from one place.
func Count(req api.SimRequest) (int, error) { return CountWith(req, Options{}) }

// CountWith is Count with the planner's options.
func CountWith(req api.SimRequest, opt Options) (int, error) {
	combos, err := ExpandWith(req, opt)
	if err != nil {
		return 0, err
	}
	return len(combos), nil
}
