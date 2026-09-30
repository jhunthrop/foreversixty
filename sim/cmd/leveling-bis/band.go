package main

import "sort"

// sourceFor answers loot.json's own question for one item at one
// band: does it have a source usable at this level, and if so which.
//
// Raid sources are excluded below level 60 (this lane's brief: "a
// leveling list is about what a leveling character can get" - the
// leveling-bis design doc's own words). An item with more than one
// source (rare in today's data; loot.json's kinds barely overlap)
// takes the highest-priority one in a fixed kind order, so the choice
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
//
// wowhead-world-drops lane, 2026-09-29: the priority order is now
// quest, boss (dungeon/raid, by highest chance - bestBoss below), rep,
// vendor, crafted, world_drop, pvp, world - a real boss kill now
// outranks a reputation/vendor/crafted purchase, reversing the
// 2026-09-28 night-bis-sources lane's own order below. That lane put
// "rep"/"vendor" ahead of "dungeon" because a fixed kind order picked
// whichever dungeon SOURCE (not boss) an item's array entry happened to
// list first, with no way to tell a real boss from a trash mob dumped
// in the same zone bucket - tenet 7's "never show an arbitrary trash
// mob when a boss ... exists" now has the data (ItemChances) to make
// that distinction directly, so a boss ranking below rep/vendor no
// longer buys anything a chance-aware pick could not do better; ranking
// it above them instead means the report says "kill this named boss"
// rather than "buy this from a vendor" whenever a real boss drop
// exists, which is the more specific, more direct instruction (tenet
// 5's "accuracy first" reading of what a leveling player actually
// wants: the fastest, most reliable route, not merely a purchasable
// one).
//
// "vendor" still sits right after "rep": a vendor purchase needs only
// gold and the right faction/level (both already gated elsewhere -
// sourceObtainable's Side check and eligible.go's factionRestriction
// check), so it is at least as reliable while leveling as a reputation
// grind and more reliable than an ungated open-world/pvp source. It was
// missing entirely until the 2026-09-28 night-bis-sources lane, even
// though loot.json has carried vendor sources (and the web's LOOT_KINDS
// has known the kind) since the vendor/zone kinds landed - a
// vendor-only item fell through every case here and reported "no known
// source".
//
// "rep" outranks "vendor": a reputation quartermaster's own vendor row
// duplicates its rep tier's item list verbatim (data.go's
// vendorInheritsRepStandingGate, 2026-09-28 night-bis-sources lane's own
// defect 2), so an item with both sources is the identical purchase
// either way once the vendor row correctly inherits the rep gate - the
// report should say what a player actually has to do ("Silverwing
// Sentinels (honored)"), not the generic "vendor" label that a plain
// kind order would prefer.
//
// "world_drop" sits right after "crafted": loot.json's own generic,
// bind-on-equip world-drop bucket (pipeline.loot.classicdb's and, since
// this lane, pipeline.loot.wowhead's own synthetic source) IS obtainable
// at any level in its own range - buy it off the auction house - so it
// ranks above "pvp"/"world" (each needs a specific grind), but below
// every source that names an exact, single place to go, and below a
// real boss (bossKinds), matching tenet 7's own ranking: "World drop" is
// the least specific true thing this pipeline can say about an item.
var sourceKindPriority = []string{
	"quest", "rep", "vendor", "crafted", "world_drop", "pvp", "world",
}

// bossKinds are the two kinds loot.json gives a per-boss chance
// (lootBoss.ItemChances): a real dungeon or raid kill, checked as one
// combined priority tier (bestBoss below) right after "quest" - see
// sourceFor's own doc for why a boss now outranks rep/vendor/crafted.
var bossKinds = [2]string{"dungeon", "raid"}

// repStandingObtainable is the highest reputation standing a leveling
// character is assumed to reach: friendly and honored come from playing
// the zone or a few battlegrounds; revered and exalted are endgame grinds,
// so a rep reward behind them counts as obtainable only at 60.
var repStandingObtainable = map[string]bool{"friendly": true, "honored": true}

// sourceObtainable is whether one source can actually be used by a
// character of this faction and level: a reputation only the other side
// can earn never is; a revered/exalted reward is only at 60.
//
// itemFactionRestriction is the CANDIDATE ITEM's own faction_restriction
// (candidate.FactionRestriction, "" for an unrestricted item) - the
// client's own hard gate, already checked once by eligible.go before this
// item's candidate ever reaches buildBandPool. A mined rep source's Side
// is wrong for a real pair of items in today's loot.json: Scout's
// Medallion (item 20442, horde_only) and Sentinel's Medallion (item
// 20444, alliance_only) are both battleground-reputation honored
// rewards, sold by an in-faction vendor (Kelm Hargunth in the Barrens;
// Illiyana Moonblaze in Ashenvale - wowhead's own pages confirm Scout's
// is Horde and Sentinel's is Alliance), but loot.json's mined rep source
// has each under the OTHER side's WSG faction (Scout's under Silverwing
// Sentinels/889, Sentinel's under Warsong Outriders/890) - the exact
// inverse of the item's own client-stated restriction. Since a
// faction-restricted item can only ever reach this function already
// carrying the one faction that can equip it, that hard restriction is
// trusted over a mined rep Side that contradicts it, rather than
// rejecting the item's only source and leaving the slot empty (this is
// what emptied hunter's level-20 neck slot for both sides at once: each
// Medallion was excluded for its own faction by the Side check below,
// and for the other faction by the item's own restriction).
func sourceObtainable(s itemSource, level int, faction, itemFactionRestriction string) bool {
	// This lane's brief, item 3: a source data/curated/loot/forever-raid-
	// phases.json curated a content phase for (Opens non-empty - every
	// Era raid this build's generator emits, "later" for six of them and
	// "raids-1" for Onyxia's Lair) is refused at EVERY band, not only
	// below 60: that file's own notes say plainly "Nothing raids at
	// launch on 4 November; the first tier opens on 9 December" - a
	// leveling list is about what a launch-day character can get
	// (band.go's own sourceFor doc), and no raid is open on launch day
	// at all, band 60 included. Fixes the worst finding in this lane's
	// brief: all nine caster specs' band-60 main_hand was Atiesh
	// (Naxxramas' own legendary quest chain), and every trinket slot's
	// "verified" pick across every melee/caster spec reviewed was a
	// Naxxramas/BWL/AQ trinket.
	if s.Opens != "" {
		return false
	}
	if s.Side != "" && s.Side != faction && itemFactionRestriction != faction {
		return false
	}
	// Standing, not Kind == "rep": a "vendor" source that
	// vendorInheritsRepStandingGate (data.go) gated from a matching rep
	// source carries the same non-empty Standing a "rep" source does,
	// and must be checked the same way - every other kind (quest,
	// dungeon, crafted, pvp, world, raid, and an ordinary vendor with
	// no matching rep source) never sets Standing at all, so this is
	// exactly as narrow as the old Kind check for them.
	if s.Standing == "" {
		return true
	}
	return level >= 60 || repStandingObtainable[s.Standing]
}

func sourceFor(id, level int, faction, itemFactionRestriction string, idx lootIndex) (itemSource, bool) {
	srcs := idx[id]
	if len(srcs) == 0 {
		return itemSource{}, false
	}
	obtainable := make([]itemSource, 0, len(srcs))
	for _, s := range srcs {
		if !sourceObtainable(s, level, faction, itemFactionRestriction) {
			continue
		}
		if s.Kind == "raid" && level < 60 {
			continue
		}
		obtainable = append(obtainable, s)
	}
	if quest, ok := bestOfKind(obtainable, "quest"); ok {
		return quest, true
	}
	if boss, ok := bestBoss(obtainable); ok {
		return boss, true
	}
	for _, kind := range sourceKindPriority[1:] { // ["quest", ...] - quest is already checked
		if s, ok := bestOfKind(obtainable, kind); ok {
			return s, true
		}
	}
	return itemSource{}, false
}

// bestOfKind is the highest-Chance obtainable source of exactly this
// kind - the tie-break `sourceFor` needs whenever more than one source
// of the same kind names the same item (rare, but a `world`/`vendor`
// item can have two). Every non-boss kind carries no chance data in
// practice (Chance stays 0 for all of them), so this degrades to "the
// first one in loot.json's own order" exactly like the old byKind map
// did, for every kind sourceFor still walks in a fixed order below.
func bestOfKind(obtainable []itemSource, kind string) (itemSource, bool) {
	var best itemSource
	found := false
	for _, s := range obtainable {
		if s.Kind != kind || (found && s.Chance <= best.Chance) {
			continue
		}
		best, found = s, true
	}
	return best, found
}

// bestBoss is the highest-Chance obtainable dungeon OR raid source -
// sourceFor's own doc: a real boss kill (bossKinds) is checked as one
// combined tier, by drop chance, rather than as two separate kind slots
// in a fixed order the way rep/vendor/crafted still are. An item
// dropping from more than one boss (two different dungeons, or a
// dungeon and a raid) picks whichever one a leveling character is
// actually more likely to see it from; an unknown chance (0, the
// overwhelming majority of today's dungeon/raid rows - only classic-db
// states one) never loses a tie against another unknown chance, so
// "first boss loot.json happens to list" is still the fallback exactly
// like it always was, not a regression this lane introduces.
func bestBoss(obtainable []itemSource) (itemSource, bool) {
	var best itemSource
	found := false
	for _, s := range obtainable {
		if (s.Kind != bossKinds[0] && s.Kind != bossKinds[1]) || (found && s.Chance <= best.Chance) {
			continue
		}
		best, found = s, true
	}
	return best, found
}

// legendaryQuality is the client's own quality tier for an orange/
// legendary item (items.json's own "quality" field, confirmed against
// Atiesh's four ids, both Thunderfury ids and both Sulfuras ids - this
// lane's brief, item 2).
const legendaryQuality = 5

// legendaryGatedLater reports whether c is a legendary - this lane's
// brief, item 2: replaces data.go's raidLockedQuestOpens hand list (which
// needed its own manual quest-id entry for Atiesh AND, separately, for
// Thunderfury, and would need a fresh one for every future legendary this
// build ever adds) with one rule that catches every legendary regardless
// of which source kind loot.json's own pipeline classified it under. A
// legendary is exactly as launch-day-unobtainable as a direct raid drop
// whether a character reaches it through a raid boss, a class-quest chain
// gated behind one (Atiesh's own Naxxramas chain, Thunderfury's Molten
// Core Bindings), or - should this build ever add one - a vendor or
// crafted recipe; no source kind makes a legendary launch-day-obtainable.
// raidLockedQuestOpens keeps the Ahn'Qiraj war-effort quest ids (8756/
// 8789/8790): those rewards are not legendaries (their own items.json
// quality is well below 5) and are gated by the war-effort RAID EVENT,
// a fact this rule has no way to see.
func legendaryGatedLater(c candidate) bool {
	return c.Quality == legendaryQuality
}

// pvpRankCap is the highest PvP rank reward this lane's own ruling
// (owner, 2026-09-29) treats as something "a fresh level 60 can
// realistically get": rank 10 is a few weeks of casual battlegrounds,
// rank 11+ needs sustained top-bracket play with rank decay, which a
// leveling list - "what a launch-day character can get", band.go's
// own sourceFor doc - should not default to any more than it defaults
// to a raid drop. Unlike a raid/legendary gate (sourceObtainable's
// Opens check, legendaryGatedLater above), a rank-11+ reward is NOT
// removed from the candidate pool: it is still a real, obtainable
// (just unlikely) item, so it stays sourced and scored, able to
// appear as a labelled alternative ("Rank 18") - only pick()'s own
// picking pool (main.go's runSpec, pickBySlot) excludes it from being
// the DEFAULT pick. See pick.go's excludeAbovePvpRankCap for exactly
// where that line is drawn.
const pvpRankCap = 10

// pvpRankExceedsCap reports whether s carries a PvP rank above
// pvpRankCap. Checked on Rank alone, not Kind == "pvp": a PvP rank
// quartermaster's own vendor row (Captain O'Neal, Rank 18) duplicates
// the identical item list its matching "pvp" source names, and
// data.go's vendorInheritsPvpRankGate copies Rank onto that vendor row
// for exactly this reason - "vendor" outranks "pvp" in
// sourceKindPriority, so sourceFor picks the vendor row for a
// dual-listed item every time, and a Kind == "pvp" check here would
// never see it. Rank is 0 (the Go zero value) for every source
// vendorInheritsPvpRankGate did not touch and every non-pvp kind
// loot.json itself never populates a nonzero rank for (confirmed
// against the whole build: no non-"pvp" source carries one), so this
// is still false by construction for an ordinary vendor/quest/rep/etc.
// source.
func pvpRankExceedsCap(s itemSource) bool {
	return s.Rank > pvpRankCap
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
	// Coverage is planner slot -> how many items eligible() passed for
	// this band+faction fan out to that slot, and how many of those
	// sourceFor() could actually find a source for (lane
	// rank-guardrails' guardrail A: "a review can no longer confirm
	// what is there instead of testing what should be there" - a slot
	// whose coverage is 2 eligible/0 sourced is a slot the ranker can
	// never fill no matter how good its scoring gets, and today's page
	// has no way to say that other than the per-band NoSourceCount,
	// which is not broken out per slot). Counted for every eligible
	// item regardless of whether buildBandPool's own crossClassSetItem
	// exclusion later drops it from Scored/NoSource -- a cross-class
	// set item is still real, equippable, (in)sourced gear a player
	// could look up, and hiding it from this count would silently
	// undercount exactly the honesty gap this field exists to surface.
	// An item with zero Slots (see the loop below) cannot be attributed
	// to any slot and is not counted here either, matching NoSource's
	// own silent drop for the same case.
	Coverage map[string]coverageRow
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
func buildBandPool(items []candidate, idx lootIndex, classSlug string, level int, faction string, weights map[string]float64, referenceDPSPerPoint float64) bandPool {
	var out bandPool
	out.Coverage = make(map[string]coverageRow)
	for _, c := range items {
		if !eligible(c, classSlug, level, faction) {
			continue
		}
		if legendaryGatedLater(c) {
			// This lane's brief, item 2: gated exactly like a raid drop
			// (below), before sourceFor ever runs - a legendary has no
			// "obtainable via a different, non-raid kind" case sourceFor
			// could otherwise find (loadLootIndex's quest-kind rows are
			// the actual gap this closes: Atiesh/Thunderfury's own quest
			// source carries no raid Opens tag at all). Counted as
			// not-sourced, same accounting a raid-gated item already
			// gets, so Coverage/NoSourceCount stay an honest count of
			// "eligible items this band could not actually put in a
			// leveling character's hands" rather than silently excluding
			// legendaries from that count.
			addCoverage(out.Coverage, c.Slots, false)
			if len(c.Slots) > 0 {
				out.NoSource = append(out.NoSource, c)
			}
			continue
		}
		src, ok := sourceFor(c.ID, level, faction, c.FactionRestriction, idx)
		addCoverage(out.Coverage, c.Slots, ok)
		if crossClassSetItem(c, classSlug) {
			out.CrossClassSet = append(out.CrossClassSet, c)
			continue
		}
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
			Score:     score(c, c.Slots[0], weights, referenceDPSPerPoint),
			Source:    src,
			HasSource: true,
		})
	}
	sort.Slice(out.NoSource, func(i, j int) bool { return out.NoSource[i].ID < out.NoSource[j].ID })
	return out
}

// coverageRow is one slot's answer to "how many eligible items exist,
// and how many of those had a source" - report.go's bandReport.Coverage
// publishes exactly this shape (json tags live there, next to the
// other published fields; this file only computes the counts).
type coverageRow struct {
	Eligible int `json:"eligible"`
	Sourced  int `json:"sourced"`
}

// addCoverage fans one already-eligible candidate's coverage out to
// every planner slot it occupies (the same slots.go/data.go
// plannerSlots expansion candidatesBySlot uses), incrementing Eligible
// always and Sourced when sourceFor found it a usable source. A
// candidate with no Slots (see buildBandPool's own doc: an item the
// per-class file could not map to a planner slot at all) contributes
// to no slot's count, same as it contributes to none of Scored or
// NoSource either.
func addCoverage(cov map[string]coverageRow, slots []string, sourced bool) {
	for _, slot := range slots {
		row := cov[slot]
		row.Eligible++
		if sourced {
			row.Sourced++
		}
		cov[slot] = row
	}
}
