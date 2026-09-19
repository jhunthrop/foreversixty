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

// combinations is Task 14. Until then, one combination per placement.
func combinations(req api.SimRequest, places []placement) ([]Combination, error) {
	out := make([]Combination, 0, len(places))
	for _, p := range places {
		out = append(out, apply(req, []placement{p}, nil, nil, nil))
	}
	return out, nil
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
