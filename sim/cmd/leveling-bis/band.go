package main

import "sort"

// sourceFor answers loot.json's own question for one item at one
// band: does it have a source usable at this level, and if so which.
//
// Raid sources are excluded below level 60 (this lane's brief: "a
// leveling list is about what a leveling character can get" - the
// leveling-bis design doc's own words). An item with more than one
// source (rare in today's data; loot.json's kinds barely overlap)
// takes the first non-raid one in a fixed kind order, so the choice
// is deterministic across runs rather than dependent on loot.json's
// row order.
//
// TODO(bis-data): "no usable source" is reported as "no known
// source" full stop. The design doc's "zone drop, flagged lucky"
// case needs a zone-drop kind loot.json does not carry yet (see
// data.go's loadLootIndex doc) - once it does, this function should
// return a source with a Lucky marker instead of ok=false for one,
// and the caller (bandPool, below) should keep it out of the BiS
// pick but list it as a "lucky" candidate the way the design doc
// describes, rather than dropping it silently as this prototype does.
var sourceKindPriority = []string{"quest", "dungeon", "crafted", "rep", "pvp", "world", "raid"}

// repStandingObtainable is the highest reputation standing a leveling
// character is assumed to reach: friendly and honored come from playing
// the zone or a few battlegrounds; revered and exalted are endgame grinds,
// so a rep reward behind them counts as obtainable only at 60.
var repStandingObtainable = map[string]bool{"friendly": true, "honored": true}

// sourceObtainable is whether one source can actually be used by a
// character of this faction and level: a reputation only the other side
// can earn never is; a revered/exalted reward is only at 60.
func sourceObtainable(s itemSource, level int, faction string) bool {
	if s.Side != "" && s.Side != faction {
		return false
	}
	if s.Kind != "rep" {
		return true
	}
	return level >= 60 || repStandingObtainable[s.Standing]
}

func sourceFor(id, level int, faction string, idx lootIndex) (itemSource, bool) {
	srcs := idx[id]
	if len(srcs) == 0 {
		return itemSource{}, false
	}
	byKind := make(map[string]itemSource, len(srcs))
	for _, s := range srcs {
		if !sourceObtainable(s, level, faction) {
			continue
		}
		if _, seen := byKind[s.Kind]; !seen {
			byKind[s.Kind] = s
		}
	}
	for _, kind := range sourceKindPriority {
		s, ok := byKind[kind]
		if !ok {
			continue
		}
		if kind == "raid" && level < 60 {
			continue
		}
		return s, true
	}
	return itemSource{}, false
}

// bandPool is everything candidatesBySlot/pick need for one band and
// faction: every eligible, sourced item, scored - plus the ones that
// were eligible but had no usable source, kept only so the report can
// say how many and which.
type bandPool struct {
	Scored        []scored
	NoSource      []candidate
	CrossClassSet []candidate
	// NoDPSWeapon is every scored weapon candidate (main_hand, off_hand
	// or ranged) whose DPS field is 0 - see this file's own
	// weaponWithNoDPS doc for what that means and whose gap it is.
	NoDPSWeapon []candidate
}

// weaponSlots is which planner slots score.go's weaponAPStat also
// keys off of: the ones where an item's DPS, not just its flat stats,
// is part of its value.
var weaponSlots = map[string]bool{"main_hand": true, "off_hand": true, "ranged": true}

// weaponWithNoDPS reports whether c occupies a weapon slot but carries
// no dps figure (data/builds/<build>/items/<class>.json's own "dps"
// field, classItem.DPS in data.go) - this lane's brief: "weapons with
// no dps are lane data-weapons' job; rank by whatever dps the item
// file has and note the gap." score() already does the "rank by
// whatever dps the item file has" half (c.DPS defaults to 0 and simply
// contributes nothing to the weighted sum, the same as any other
// missing stat - no special-casing needed there). This function is the
// "note the gap" half: buildBandPool below collects every such weapon
// so main.go can log the count per band+faction, honest about which of
// this band's weapon picks the pipeline could not actually rank
// on damage.
func weaponWithNoDPS(c candidate) bool {
	if c.DPS > 0 {
		return false
	}
	for _, s := range c.Slots {
		if weaponSlots[s] {
			return true
		}
	}
	return false
}

// crossClassSetItem answers this lane's run-time finding: at band 60 a
// hunter candidate pool carried Bonescythe Armor pieces (e.g. id 22476,
// set_id 524) - a ROGUE-only Naxxramas set (the engine's own db.json
// carries classAllowlist:[6], proto.Class_ClassRogue, for every piece)
// that data/builds/<build>/items/hunter.json still lists, because that
// per-class file's own eligibility rule is armor type (leather, which
// hunters can wear below 40 and still equip above it), not the game's
// actual per-item class restriction - a data-pipeline gap outside this
// lane's scope. wowsims-forever's item_sets_pve.go registers that
// set's bonuses with `agent.(rogue.RogueAgent)`
// (sim/rogue/items_sets_pve.go's own unconditional type assertion),
// which panics building the engine's Environment for any OTHER class
// wearing 2+ pieces:
//
//	interface conversion: *hunter.Hunter is not rogue.RogueAgent: missing method GetRogue
//	  github.com/wowsims/classic/sim/rogue.init.func17 (items_sets_pve.go:269)
//	  .../sim/core.(*Character).applyItemSetBonusEffects (item_sets.go:159)
//	  .../sim/core.(*Character).applyAllEffects (character.go:348)
//	  .../sim/core.(*Raid).applyCharacterEffects (raid.go:291)
//	  .../sim/core.(*Environment).initialize (environment.go:132)
//	  .../sim/core.NewEnvironment (environment.go:56)
//	  .../sim/core.NewSim (sim.go:199)
//	  created by .../sim/core.runSimConcurrent (sim_concurrent.go:507)
//
// core.runSim's own recover (sim.go:134) DOES catch this - it reaches
// this command as an ordinary Go error (adapter.ResultError, "the
// engine reported an error: interface conversion: ..."), not a process
// crash. What actually breaks is verifyBand's own contract: a BASELINE
// failure (as opposed to a runner-up swap failure) has no DPS to
// report at all, so it propagates out of main's per-band loop and
// stops the ENTIRE nightly run over one bad candidate at one band for
// one spec - an availability bug, not a crash-safety one.
//
// The first fix this lane tried was excluding a set item whenever the
// embedded item database's own per-item Classes list (the same one
// eligible() and every other candidate check in this command already
// consults) did not name this class. It does not work: Classes there
// is "which classes could physically equip this item" (armor-type
// proficiency, data/pipeline/simdb/items.py's own _class_allowlist,
// built off the item's AllowableClass mask), not "which class this
// SET's bonus was authored for" - Bonescythe Armor's own Classes list
// names every class (any of them can wear the leather), and
// wowsims-forever's engine never checks class before invoking a set
// bonus at all (core.Character.applyItemSetBonusEffects matches
// equipped pieces by SetID/SetName only - see item_sets.go's
// GetActiveSetBonuses/applyItemSetBonusEffects); the per-class cast
// inside the Bonuses function is the only place "this class" is ever
// checked, and by then it is a type assertion, not a guard. That is
// the real engine bug this lane is filing a follow-up for (see the
// lane report): nothing under sim/core stops a non-rogue from
// registering 2+ Bonescythe pieces and reaching rogue.init.func17.
//
// The fix that does work has to name the set's TRUE owning class, which
// is not data any candidate here carries (loadCandidates.go's
// classItem/flatItem never see a set's name, only its numeric id) -
// setNativeClass (setclass_generated.go) supplies it, generated once
// from the engine source itself (every core.NewItemSet(...) call's
// Name, per class package, matched against assets/database/db.json's
// setName/setId - see that file's own doc for the exact procedure). A
// set id this map does not know is left alone (unmapped, not
// cross-class): the map is complete for wowsims-forever @
// sim/enginever.Version as of this lane's run, and a gap only appears
// if a later engine pin adds a set this table has not been
// regenerated against - see the lane report's follow-ups.
func crossClassSetItem(c candidate, classSlug string) bool {
	if c.SetID == nil {
		return false
	}
	owner, known := setNativeClass[*c.SetID]
	return known && owner != classSlug
}

// buildBandPool applies eligible(), the source rule, and score(), in
// that order, to the full candidate list for one band and faction. It
// does not pick: candidatesBySlot/pick (pick.go) do that from
// Scored.
func buildBandPool(items []candidate, idx lootIndex, classSlug string, level int, faction string, weights map[string]float64) bandPool {
	var out bandPool
	for _, c := range items {
		if !eligible(c, classSlug, level, faction) {
			continue
		}
		if crossClassSetItem(c, classSlug) {
			out.CrossClassSet = append(out.CrossClassSet, c)
			continue
		}
		src, ok := sourceFor(c.ID, level, faction, idx)
		if !ok {
			out.NoSource = append(out.NoSource, c)
			continue
		}
		if len(c.Slots) == 0 {
			continue
		}
		if weaponWithNoDPS(c) {
			out.NoDPSWeapon = append(out.NoDPSWeapon, c)
		}
		out.Scored = append(out.Scored, scored{
			candidate: c,
			Score:     score(c, c.Slots[0], weights),
			Source:    src,
			HasSource: true,
		})
	}
	sort.Slice(out.NoSource, func(i, j int) bool { return out.NoSource[i].ID < out.NoSource[j].ID })
	return out
}
