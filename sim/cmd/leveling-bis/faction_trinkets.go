package main

import (
	"fmt"
	"math"
)

// trinketSlots is the fixed pair of planner slots
// reconcileFactionTrinkets ever touches - every loop below names the
// same two strings, from the one place that decides what they are.
var trinketSlots = []string{"trinket1", "trinket2"}

// factionTrinketInputs is one faction's own inputs
// reconcileFactionTrinkets needs beyond its own current picks - the
// two are kept separate (rather than folding BySlot into a picks-only
// signature) because a candidate crossing over from the OTHER faction
// must adopt THIS faction's own scored entry (its own Source, resolved
// against this faction's side of the loot index) when one exists,
// never the source faction's.
type factionTrinketInputs struct {
	// Faction is "alliance" or "horde" - band.go's sourceObtainable/
	// sourceFor's own faction string, used to decide whether a
	// candidate is faction-neutral and, on a crossover, to resolve
	// this faction's own source for it.
	Faction string
	// Race is the character race this faction's own bandCharacter
	// calls build with - measureTrinketGain's own character context.
	Race string
	// BySlot is this faction's own scored candidate pool
	// (candidatesBySlot's per-faction output, main.go's own call
	// site) - candidateForFaction (below) looks a crossing item up
	// here first, so it publishes with this faction's own Source
	// label rather than a source resolved defensively from scratch.
	BySlot map[string][]scored
}

// isFactionNeutralCandidate reports whether c carries no
// faction_restriction of its own AND has a real source both factions
// can actually use at level (sourceFor, band.go's own obtainability
// rules - the same ones buildBandPool already trusts to decide
// whether this exact item reaches either faction's own candidate
// pool) - this lane's brief's own definition of "faction-neutral":
// "no faction_restriction, every source obtainable by both factions."
// A restricted item never reaches here with a chance to cross at all
// (this lane's brief, test (d)): reconcileTrinketDirection's own call
// site only ever offers a slot's pick as a candidate to cross when
// this reports true for it.
func isFactionNeutralCandidate(c candidate, level int, idx lootIndex) bool {
	if c.FactionRestriction != "" {
		return false
	}
	if _, ok := sourceFor(c.ID, level, "alliance", c.FactionRestriction, idx); !ok {
		return false
	}
	_, ok := sourceFor(c.ID, level, "horde", c.FactionRestriction, idx)
	return ok
}

// trinketVerdictTrustworthy reports whether c's own measured gain is
// one a reader can already trust - real and significant
// (trinketGainSignificant, weights.go, the same bar report.go's own
// trinketLowGain gate uses) - as opposed to a gain that was measured
// but sat inside its own noise floor. reconcileTrinketDirection's own
// "is target empty" check (below) reads this as its "or no_dps_value"
// half: a slot whose own pick's gain never cleared this bar is, from a
// reader's point of view, exactly as unconvincing as a literally empty
// slot, whichever physical item happens to be sitting in it.
func trinketVerdictTrustworthy(c *scored) bool {
	return c != nil && c.GainMeasured && trinketGainSignificant(c.MeasuredGainDPS, c.MeasuredGainStdErr)
}

// gainAlreadyMeasuredOnTarget reports whether target already carries a
// TRUSTWORTHY measured gain for itemID - true only when itemID is
// target's own current pick or its own runner-up (the two candidates a
// tournament retains a real MeasuredGainDPS for at all, trinkets.go's
// own rankTrinketSlot doc) AND that measurement is itself significant
// (trinketVerdictTrustworthy, above) - so
// reconcileTrinketDirection can reuse that real number instead of
// running a fresh sim, but only when the existing number is actually
// worth trusting. A target whose own pick happens to be the identical
// item but whose own gain sat inside its own noise floor reports false
// here on purpose: reusing that same noisy number would refresh
// nothing (this lane's own druid-feral/druid-balance band 50 repro,
// this file's own doc on reconcileFactionTrinkets) - the one honest
// fix is a fresh, adaptively-escalated measurement, exactly the one
// reconcileTrinketDirection runs when this reports false.
func gainAlreadyMeasuredOnTarget(target slotPick, itemID int) (absoluteDPS, gainDPS, gainStdErr float64, ok bool) {
	if target.Item != nil && target.Item.ID == itemID && trinketVerdictTrustworthy(target.Item) {
		return target.Item.MeasuredDPS, target.Item.MeasuredGainDPS, target.Item.MeasuredGainStdErr, true
	}
	if target.RunnerUp != nil && target.RunnerUp.ID == itemID && trinketVerdictTrustworthy(target.RunnerUp) {
		return target.RunnerUp.MeasuredDPS, target.RunnerUp.MeasuredGainDPS, target.RunnerUp.MeasuredGainStdErr, true
	}
	return 0, 0, 0, false
}

// gainsIndistinguishable reports whether two independently measured
// trinket gains differ by no more than their own combined standard
// error (the same combination-in-quadrature rule trinkets.go's own
// gainStdErr uses for a candidate-vs-baseline difference, here applied
// to two DIFFERENT items' own gains) - reconcileTrinketDirection's own
// "keep the numerically higher one" step needs this to tell a genuine,
// reproducible edge apart from the same near-tie this file's own doc
// already names as a live repro (Second Wind/Burst of Knowledge):
// a small, same-magnitude gap here is not evidence either faction's
// own pick is actually better, only that its own tournament's noise
// happened to land it a little ahead.
func gainsIndistinguishable(gainA, stdErrA, gainB, stdErrB float64) bool {
	combined := math.Sqrt(stdErrA*stdErrA + stdErrB*stdErrB)
	return math.Abs(gainA-gainB) <= combined
}

// candidateForFaction is the scored record reconcileTrinketDirection
// publishes for a crossing item on the TARGET faction: target's own
// bySlot[slot] entry when one exists (the ordinary case for a real
// faction-neutral item - it is eligible and sourced for target too,
// exactly as isFactionNeutralCandidate just confirmed, so
// buildBandPool already scored and sourced it for this faction on its
// own). Falls back to source's own scored copy with target's own
// Source resolved directly (sourceFor) only when target's pool
// somehow never carried it (a slot-narrowing pass upstream of
// bySlot[slot] excluded it for a reason unrelated to faction - this
// lane's report names any case this fallback actually fires for) -
// never source's OWN Source label, which would publish what the OTHER
// faction's player does to get it, not this one's.
func candidateForFaction(pool []scored, itemID int, fallback scored, level int, faction string, idx lootIndex) scored {
	for _, c := range pool {
		if c.ID == itemID {
			return c
		}
	}
	resolved := fallback
	if src, ok := sourceFor(itemID, level, faction, fallback.FactionRestriction, idx); ok {
		resolved.Source, resolved.HasSource = src, true
	}
	return resolved
}

// clonePicksMap is this file's own copy of the immutable-update pattern
// every other picks-mutating pass in this command already uses
// (rankTrinketSlot, rankSlotWithEffects, trySetCompletion, applySwaps,
// clearTrinketPlaceholders - pick.go's own doc): a new map, never an
// edit to the caller's own.
func clonePicksMap(picks map[string]slotPick) map[string]slotPick {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}
	return out
}

// reconcileFactionTrinkets is this lane's brief (bis-ranker-
// integrity-15, twelfth sweep, ranker-14's own deferred diagnosis): a
// faction-neutral trinket - no faction_restriction, obtainable by both
// sides (isFactionNeutralCandidate, above) - is worth the identical
// amount to both factions, modulo a real racial stat difference,
// because both factions' characters are otherwise identical for this
// spec's own rotation and gear. rankTrinketSlot's own 2x-stderr
// significance gate (ranker-11) and adaptive re-measure (ranker-12)
// each run once PER FACTION at a bounded iteration count, so a
// measurement sitting near either bar can land on opposite sides of it
// for the two factions purely from independent sim noise, not from any
// real difference between them.
//
// This lane's own repro, dogfooded directly against the live engine
// (druid-feral/druid-balance band 50, both factions - see the lane
// report for the full before/after): the SAME item, Frozen Heart of
// the Mountain (id 249469, crafted, faction-neutral), was each
// faction's own trinket2 winner in both specs - Alliance's own
// tournament measured its gain at 0.00 +/- 0.65 DPS (inside its own
// noise floor, so report.go's trinketLowGain gate correctly published
// the slot empty), Horde's own tournament measured the IDENTICAL item
// at 0.86 +/- 0.21 DPS (comfortably significant) in the same run. Nothing
// about the item or the character differs between the two publishes -
// only which faction's own 100-iteration tournament happened to land
// closer to the true value. A literal "is this the same item id"
// check would wrongly call that a no-op (this file's own first draft
// did exactly that, caught by this dogfood run, not by a synthetic
// test); trinketVerdictTrustworthy/gainAlreadyMeasuredOnTarget (above)
// instead ask "is TARGET's own verdict for this item already
// trustworthy", which is false here even though the item id matches,
// so reconcileTrinketDirection (below) re-measures it on Alliance's
// own character at the same adaptive escalation rankTrinketSlot's own
// winner already gets, rather than leaving a published "no_dps_value"
// standing next to the other faction's own real, significant number
// for the identical item.
//
// This is a contained POST-pass (this lane's brief's own words): run
// once per band, after both factions' own rankTrinketSlot tournaments
// finish (main.go's own call site, between the loop's trinket-ranking
// half and its rankSlotWithEffects/trySetCompletion/verifyBand half) -
// never a cache threaded through trinketShortlist/rankTrinketSlot
// themselves, which this lane leaves exactly as every prior lane did.
// It reconciles trinket1, then trinket2, each slot in both directions
// (Alliance's own pick offered to Horde, then Horde's own - possibly
// now-updated - pick offered back to Alliance) exactly once -
// reconcileTrinketDirection's own doc has the one-directional rule; a
// second full call on this function's own result changes nothing
// further (TestReconcileFactionTrinketsSecondRoundIsANoOp), because
// every direction's own first question is "is target's own verdict for
// this item already trustworthy", which a converged slot always
// answers yes to.
func reconcileFactionTrinkets(runner engineRunner, spec specInfo, classSlug string, level int, talents string, idx lootIndex, alliance factionTrinketInputs, alliancePicks map[string]slotPick, horde factionTrinketInputs, hordePicks map[string]slotPick) (map[string]slotPick, map[string]slotPick, []string) {
	var notes []string
	for _, slot := range trinketSlots {
		hordePicks = reconcileTrinketDirection(runner, spec, classSlug, level, talents, idx, slot, alliance, alliancePicks, horde, hordePicks, &notes)
		alliancePicks = reconcileTrinketDirection(runner, spec, classSlug, level, talents, idx, slot, horde, hordePicks, alliance, alliancePicks, &notes)
	}
	return alliancePicks, hordePicks, notes
}

// reconcileTrinketDirection applies this lane's brief's reconcile rule
// one direction, for one slot: source's own pick, when it is a real,
// faction-neutral candidate, is offered to target. target adopts it -
// with target's OWN measured numbers, never source's - when target's
// own verdict for it is not already trustworthy
// (trinketVerdictTrustworthy): target's slot is empty, target's own
// pick is this SAME item but its own gain never cleared significance
// (this file's own doc above has the live repro), or target's own pick
// is a DIFFERENT faction-neutral candidate that measures a lower gain,
// on target, than source's pick does. A former, DIFFERENT pick target
// is demoted from becomes the new pick's runner-up, so it flows into
// target's own published alternatives and a real verifyBand re-check
// exactly the way any other slot's runner-up already does (verify.go's
// own doc - this reconciliation intentionally runs BEFORE verifyBand,
// so nothing downstream needs to know a crossover happened at all).
//
// Racials are the one legitimate reason the two factions can differ
// (this lane's brief): when target's own measured gain for source's
// pick is negative beyond its own error (negativeBeyondError,
// weights.go) - a real, reproducible loss, not sampling noise around
// zero - target's own verdict is kept and FactionNote (slotPick,
// pick.go) names why, rather than silently leaving the row's own
// history of the comparison invisible.
//
// A target whose own current pick is real, different, and already
// trustworthy, but NOT faction-neutral (almost always
// faction-restricted), is never touched here at all: that pick is
// target's own legitimate, faction-specific verdict, and source's
// neutral candidate - obtainable on target's own side too, by
// definition - is not evidence target's own restricted item is wrong,
// only that it has one more competitor this pass does not referee
// between two DIFFERENT faction-restricted items (this lane's brief,
// test (d): "a faction-restricted trinket never crosses").
func reconcileTrinketDirection(runner engineRunner, spec specInfo, classSlug string, level int, talents string, idx lootIndex, slot string, source factionTrinketInputs, sourcePicks map[string]slotPick, target factionTrinketInputs, targetPicks map[string]slotPick, notes *[]string) map[string]slotPick {
	srcItem := sourcePicks[slot].Item
	if srcItem == nil || !isFactionNeutralCandidate(srcItem.candidate, level, idx) {
		return targetPicks
	}

	targetPick := targetPicks[slot]
	targetTrustworthy := trinketVerdictTrustworthy(targetPick.Item)
	if targetTrustworthy && targetPick.Item.ID == srcItem.ID {
		// Already the same physical item on both sides, and target's
		// own verdict for it is already trustworthy - nothing to
		// reconcile. This is also what makes a second full round a
		// no-op: once a slot converges on a trustworthy verdict, this
		// check short-circuits every later call for it, in either
		// direction.
		return targetPicks
	}
	if targetTrustworthy && !isFactionNeutralCandidate(targetPick.Item.candidate, level, idx) {
		return targetPicks
	}

	absOnTarget, gainOnTarget, stdErrOnTarget, alreadyMeasured := gainAlreadyMeasuredOnTarget(targetPick, srcItem.ID)
	if !alreadyMeasured {
		dps, gainDPS, gainStdErr, err := measureTrinketGain(runner, spec, target.Race, classSlug, level, talents, targetPicks, slot, srcItem.ID)
		if err != nil {
			*notes = append(*notes, fmt.Sprintf("%s: measuring %s (id %d) on %s for faction reconciliation failed, keeping %s's own pick: %v", slot, srcItem.Name, srcItem.ID, target.Race, target.Faction, err))
			return targetPicks
		}
		absOnTarget, gainOnTarget, stdErrOnTarget = dps, gainDPS, gainStdErr
	}

	if negativeBeyondError(gainOnTarget, stdErrOnTarget) {
		out := clonePicksMap(targetPicks)
		row := out[slot]
		row.FactionNote = fmt.Sprintf("%s measures lower on %s (racial)", srcItem.Name, target.Race)
		out[slot] = row
		*notes = append(*notes, fmt.Sprintf("%s: kept %s's own verdict over %s (id %d) from %s - measures %.2f ± %.2f DPS gain on %s, a racial difference", slot, target.Faction, srcItem.Name, srcItem.ID, source.Faction, gainOnTarget, stdErrOnTarget, target.Race))
		return out
	}

	// This check is NOT gated on targetTrustworthy: dogfooded directly
	// against the live engine (rogue-assassination/rogue-subtlety/
	// shaman-enhancement band 50, this lane's own report) - Alliance's
	// own Molten Heart of the Mountain (id 249470, gain 0.30, itself
	// inside its own noise floor and about to publish empty) was being
	// REPLACED by Horde's Frozen Heart of the Mountain re-measured on
	// Alliance at 0.00 DPS, a strictly WORSE number, purely because
	// "target not yet trustworthy" was being read as "anything beats
	// it". A candidate that already measures a higher gain than
	// source's own pick does on target must never be demoted for a
	// weaker one, whether or not either number clears significance -
	// significance decides what PUBLISHES, not which of two real
	// measurements is the better one to keep.
	//
	// data-followups-10 lane, 2026-09-30: that rule alone let two
	// DIFFERENT faction-neutral items stand as each faction's own pick
	// with NO record of the comparison at all whenever each one's own
	// gain happened to sit fractionally above the other's, purely from
	// two independent noisy measurements - the exact shape this file's
	// own doc above already names as a live repro (Second Wind/Burst of
	// Knowledge, 186.96 vs 188.28 DPS with nothing else different,
	// paladin-retribution band 60): each side's own bare `<=` compare
	// went the "keep, say nothing" way in BOTH directions, so the two
	// factions published two different trinkets for an identical slot
	// with no FactionNote explaining why, unlike negativeBeyondError's
	// own branch above (a REAL, reproducible loss) or the swap branch
	// below (a REAL, reproducible win) - both of which always leave a
	// note. gainsIndistinguishable asks the same "is this difference
	// bigger than the two measurements' own combined noise" question
	// negativeBeyondError already asks for the zero-vs-gain case, here
	// applied to two nonzero gains: when the gap is not real, this is
	// a genuine cross-faction tie (tenet 8: unverifiable is labelled,
	// never left as an unexplained fact), so target keeps its own pick
	// WITH a note recording the comparison, same shape as the racial
	// branch above, rather than silently diverging.
	if targetPick.Item != nil && targetPick.Item.ID != srcItem.ID && targetPick.Item.GainMeasured {
		if gainsIndistinguishable(gainOnTarget, stdErrOnTarget, targetPick.Item.MeasuredGainDPS, targetPick.Item.MeasuredGainStdErr) {
			out := clonePicksMap(targetPicks)
			row := out[slot]
			row.FactionNote = fmt.Sprintf(
				"%s (from %s) and %s's own pick %s measure statistically indistinguishable on %s (%.2f ± %.2f vs %.2f ± %.2f DPS gain) -- each faction's own tournament winner kept",
				srcItem.Name, source.Faction, target.Faction, targetPick.Item.Name, target.Race,
				gainOnTarget, stdErrOnTarget, targetPick.Item.MeasuredGainDPS, targetPick.Item.MeasuredGainStdErr,
			)
			// This note names Item itself ("target's own pick") - only
			// true while Item goes on to publish as this slot's pick.
			// report.go's own zero-value gate reads this to drop the
			// note instead of publishing it (bis-ranker-integrity-16,
			// item 3) if Item's own measured gain later fails
			// trinketGainSignificant and the row empties.
			row.FactionNoteNeedsPick = true
			out[slot] = row
			*notes = append(*notes, fmt.Sprintf(
				"%s: kept %s's own verdict over %s (id %d) from %s - gains are statistically indistinguishable on %s (%.2f ± %.2f vs %.2f ± %.2f DPS), a cross-faction tie rather than a real difference",
				slot, target.Faction, srcItem.Name, srcItem.ID, source.Faction, target.Race,
				gainOnTarget, stdErrOnTarget, targetPick.Item.MeasuredGainDPS, targetPick.Item.MeasuredGainStdErr,
			))
			return out
		}
		if gainOnTarget <= targetPick.Item.MeasuredGainDPS {
			return targetPicks
		}
	}

	newItem := candidateForFaction(target.BySlot[slot], srcItem.ID, *srcItem, level, target.Faction, idx)
	newItem.MeasuredDPS = absOnTarget
	newItem.MeasuredGainDPS = gainOnTarget
	newItem.MeasuredGainStdErr = stdErrOnTarget
	newItem.GainMeasured = true

	out := clonePicksMap(targetPicks)
	newPick := slotPick{Item: &newItem}
	switch {
	case targetPick.Item != nil && targetPick.Item.ID == srcItem.ID:
		// The identical item was already target's own pick, just with
		// an untrustworthy gain (this file's own doc, above) - refreshed
		// in place, so target's own original runner-up (if any) carries
		// over unchanged rather than being replaced by the item it is
		// already sitting beside.
		newPick.RunnerUp = targetPick.RunnerUp
		*notes = append(*notes, fmt.Sprintf("%s: %s's own pick %s (id %d) re-measured at %.2f DPS gain on %s (was inside its own noise floor) after %s measured the identical item as significant", slot, target.Faction, srcItem.Name, srcItem.ID, gainOnTarget, target.Race, source.Faction))
	case targetPick.Item != nil:
		demoted := *targetPick.Item
		newPick.RunnerUp = &demoted
		*notes = append(*notes, fmt.Sprintf("%s: %s adopted %s (id %d) from %s - faction-neutral, measures %.2f DPS gain on %s vs %.2f for the previous pick %s (id %d), which moves to the alternatives", slot, target.Faction, srcItem.Name, srcItem.ID, source.Faction, gainOnTarget, target.Race, targetPick.Item.MeasuredGainDPS, targetPick.Item.Name, targetPick.Item.ID))
	default:
		*notes = append(*notes, fmt.Sprintf("%s: %s adopted %s (id %d) from %s - faction-neutral, measures %.2f DPS gain on %s, the slot was empty", slot, target.Faction, srcItem.Name, srcItem.ID, source.Faction, gainOnTarget, target.Race))
	}
	out[slot] = newPick
	return out
}
