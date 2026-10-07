package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// Ranking trinkets by engine-verified DPS instead of score().
//
// score() is the weighted sum of an item's STAT block; a trinket's real
// value is almost always its proc or on-use effect, which carries no
// stats at all (see score.go's own doc). Every trinket therefore scores
// at or near 0 under the normal pipeline, which the prototype's report
// flagged as dishonest: pick()'s tie-break (item id, ascending) would
// otherwise decide trinket1/trinket2 by which id happens to be lowest,
// not which trinket is actually strongest.
//
// This lane's brief: "rank trinket candidates by the verify step
// (engine, top-N by item level) instead of the stat score." trinketTopN
// candidates per trinket slot, ordered by item level (a real, if rough,
// proxy for "how good a trinket this tier of content dropped" when
// stats cannot rank them), are each equipped in turn alongside the rest
// of this band's picks and run through a real DPS sim; the one that
// measures highest becomes the pick, the next becomes the runner-up -
// so trinket1/trinket2's Verified flag in the final report means the
// same thing it means for every other slot (a real sim, not a stat
// formula, decided it).
//
// Item-level-only selection has its own blind spot (bis-ranker-integrity
// lane, 2026-09-29): a trinket that IS mostly plain stats (Neltharion's
// Tear, spell_power 44 + hit 20 - no proc at all) can sit at a lower
// item level than several proc/on-use trinkets from the same or a later
// raid tier, so it never got a single sim and its own, much higher,
// score() total showed up only as an unverified "alternatives" entry
// the published pick had never actually been measured against (exactly
// the dishonest tie-break this file's own opening paragraph already
// named, just moved from "which id is lowest" to "which axis got
// asked"). trinketShortlist (below) fixes this by adding the top
// trinketTopN candidates BY SCORE to the real-sim pool alongside the
// item-level ones, so the best plain stat-stick and the best proc
// candidates are always tested against EACH OTHER, not just against
// their own axis.
const (
	// trinketTopN bounds how many of a slot's item-level-ranked
	// candidates get their own sim: bounded because this runs once per
	// trinket slot per band per faction per spec, and the nightly
	// budget (this lane's brief) has to fit every written spec.
	trinketTopN = 5
	// trinketRankIterations is fewer than verifyIterations (300): this
	// pass only has to ORDER a handful of candidates relative to each
	// other, not settle a single close call the way verifyBand's final
	// swap check does, so it spends less of the nightly budget per
	// candidate. Same fixed verifySeed as every other sim in this
	// command, so a report is reproducible.
	trinketRankIterations = 100
	// trinketShortlistBound is the generous cap trinketShortlist only
	// enforces once its QUALIFYING set (hasPositivelyWeightedStat OR
	// hasImplementedEffect OR hasUseEffect - see trinketShortlist's own
	// doc) is itself larger than this: an ordinary slot's qualifying
	// pool (a handful of stat-carrying or effect-carrying trinkets)
	// never approaches it, so it only ever trims the rare slot whose
	// real candidate count is unusually large, never the ordinary case
	// this lane's brief fixes (bis-ranker-integrity-8, 2026-09-30).
	trinketShortlistBound = 24
)

// excludePairMate drops a slot's own pair-mate (by id, or by name - a
// lower/higher-quality version of "the same trinket" - matching
// excludePaired's own rule in pick.go) from list, so neither
// topByItemLevel nor topByScore below can ever ask to equip the same
// physical trinket twice.
func excludePairMate(list []scored, excludeID int, excludeName string) []scored {
	candidates := make([]scored, 0, len(list))
	for _, c := range list {
		if c.ID == excludeID || (excludeName != "" && c.Name == excludeName) {
			continue
		}
		candidates = append(candidates, c)
	}
	return candidates
}

// topByItemLevel returns up to trinketTopN candidates from list, highest
// item level first (ties broken by id, ascending, for a stable order
// across runs). list is assumed already pair-mate-filtered
// (excludePairMate).
func topByItemLevel(list []scored, excludeID int, excludeName string) []scored {
	candidates := excludePairMate(list, excludeID, excludeName)
	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].ItemLevel != candidates[j].ItemLevel {
			return candidates[i].ItemLevel > candidates[j].ItemLevel
		}
		return candidates[i].ID < candidates[j].ID
	})
	if len(candidates) > trinketTopN {
		candidates = candidates[:trinketTopN]
	}
	return candidates
}

// hasPositivelyWeightedStat reports whether c carries a positive
// amount of any stat this spec's own weights (score.go's own units -
// post rating-to-percent conversion, bis-ranker-integrity-8) actually
// value (statWeight, score.go's own helper, so an "attack_power"
// entry is checked the same combined way score() itself weighs it).
//
// This is the regression fix (bis-ranker-integrity-8, 2026-09-30):
// hunter-marksmanship Alliance band 60's trinket1/trinket2 both
// published no_dps_value with zero alternatives on nightly 5036d76c
// (the first nightly after the rating-to-percent conversion, score.go)
// because that conversion shrank every rating-family stat's score()
// contribution 10-14x, which silently pushed rating trinkets like
// Frozen Heart of the Mountain (+9 hit) out of BOTH the old
// topByItemLevel and topByScore top-5 buckets before either ever ran
// a sim - the exact "which axis got asked" blind spot this file's own
// history (above) already named and fixed once for score() itself,
// just reintroduced at a different scale by the unit change. Ranking
// by SCORE or ITEM LEVEL was always only ever a proxy for "does this
// candidate carry real, spec-relevant value" - this function asks
// that question directly and exactly, so no top-N cut on either proxy
// axis can ever again silently drop a candidate the spec's own
// weights say is worth something.
func hasPositivelyWeightedStat(c candidate, weights map[string]float64) bool {
	for stat, amount := range c.Stats {
		if amount > 0 && statWeight(stat, weights) > 0 {
			return true
		}
	}
	return false
}

// hasUseEffect reports whether c's own effect_text is a real on-USE
// effect - the client's own "Use: ..." tooltip prefix, exactly as
// data/builds/<build>/items/<class>.json states it (classItem.EffectText,
// data.go) - regardless of whether the engine actually implements that
// effect (hasImplementedEffect, rank.go, is a stricter, separate
// question). This lane's brief's third qualifying condition: an
// on-use trinket the engine cannot yet fire is still a real item a
// player can click, and its own passive stats (if any, valued by
// hasPositivelyWeightedStat already) are never the whole story for an
// on-use trinket - it deserves a real sim and an honest
// effect_unmodelled flag (report.go, trinketEffectUnmodelled below),
// not silent exclusion before either.
func hasUseEffect(c candidate) bool {
	return strings.HasPrefix(c.EffectText, "Use:")
}

// trinketShortlist is rankTrinketSlot's real candidate pool: every
// pair-mate-filtered candidate that QUALIFIES - carries a positively-
// weighted stat (hasPositivelyWeightedStat), OR an engine-implemented
// effect (hasImplementedEffect - a trinket's real value is almost
// always its proc/on-use effect, invisible to score() entirely, this
// file's own package doc), OR a real on-use effect_text the engine
// does not yet implement (hasUseEffect) - is simmed for the
// tournament. No cap applies to that qualifying set unless it is
// itself larger than trinketShortlistBound, a generous bound chosen
// so the ordinary slot (a handful of qualifying candidates) is never
// trimmed at all: the exact regression this lane fixes
// (bis-ranker-integrity-8, 2026-09-30) was a top-5-by-score/top-5-by-
// item-level cap applied BEFORE any relevance check, which silently
// dropped a real, weighted-stat trinket whose score had simply moved
// (the rating-to-percent conversion, score.go) without touching its
// actual DPS value at all.
//
// Only when the qualifying set exceeds trinketShortlistBound does
// this function trim: by SCORE (already post-rating-conversion,
// list's own best-score-first contract - candidatesBySlot's own doc)
// down toward the bound, while still keeping the qualifying set's own
// top-by-item-level candidates (topByItemLevel, trinketTopN) even if
// their score fell outside that cut - a higher-item-level candidate
// remains a legitimate, real-world upgrade path this file has always
// protected (see the union this function used to keep unconditionally,
// in its history above), it just no longer stands in for "was this
// candidate worth a sim at all" the way it used to.
func trinketShortlist(list []scored, excludeID int, excludeName string, weights map[string]float64) []scored {
	filtered := excludePairMate(list, excludeID, excludeName)

	qualifying := make([]scored, 0, len(filtered))
	for _, c := range filtered {
		if hasPositivelyWeightedStat(c.candidate, weights) || hasImplementedEffect(c.candidate) || hasUseEffect(c.candidate) {
			qualifying = append(qualifying, c)
		}
	}
	if len(qualifying) <= trinketShortlistBound {
		return qualifying
	}

	byScore := qualifying[:trinketShortlistBound]
	byItemLevel := topByItemLevel(qualifying, 0, "")

	seen := make(map[int]bool, len(byScore)+len(byItemLevel))
	out := make([]scored, 0, len(byScore)+len(byItemLevel))
	for _, group := range [][]scored{byScore, byItemLevel} {
		for _, c := range group {
			if seen[c.ID] {
				continue
			}
			seen[c.ID] = true
			out = append(out, c)
		}
	}
	return out
}

// trinketEffectUnmodelled reports whether c's own effect is exactly the
// case report.go's EffectUnmodelled flag publishes: a real, named
// on-hit/on-use/proc effect (rank.go's carriesEffect) this tournament's
// own sim could not actually verify (rank.go's effectVerifiedInSim) -
// either because the engine does not implement it at all (Serenity
// Field's own Spirit self-buff, for one) or because it does, but this
// build's own simdb silently stripped the item before the sim ever ran
// (rank.go's effectVerifiedInSim doc: Hand of Justice 11815). A
// candidate with no effect at all (carriesEffect false) is NOT unmodelled
// in this sense: its whole value is stats, which this tournament's own
// real sim already measures exactly as faithfully as any other
// stat-only trinket.
func trinketEffectUnmodelled(c candidate) bool {
	return carriesEffect(c) && !effectVerifiedInSim(c)
}

// rankTrinketSlot replaces picks[slot] with the engine-verified best of
// its qualifying candidates (trinketShortlist; and the runner-up with
// the second best), leaving every other slot's pick untouched. It
// returns a new
// map rather than mutating picks (this command's own immutability
// rule - see the other pick.go/band.go functions, which all return new
// values instead of editing their input).
//
// A candidate whose own sim fails is skipped and logged, the same
// resilience verify.go's own per-slot swap loop uses (one bad
// candidate should not lose the whole slot's ranking); if every
// candidate fails, picks is returned unchanged (the slot keeps
// whatever score()-based pick it already had - almost always the
// lowest-id eligible trinket, which the report's swap_note on that row
// still leaves honestly unverified) and the caller is told so via the
// returned notes slice, one line per skipped candidate.
func rankTrinketSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored, slot string, weights map[string]float64) (map[string]slotPick, []string) {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}

	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok && picks[mate].Item != nil {
		mateID, mateName = picks[mate].Item.ID, picks[mate].Item.Name
	}
	candidates := trinketShortlist(bySlot[slot], mateID, mateName, weights)
	if len(candidates) == 0 {
		return out, nil
	}
	return rankByMeasuredGain(runner, spec, race, classSlug, level, talents, picks, slot, candidates)
}

// rankByMeasuredGain is rankTrinketSlot's tournament, for any slot whose
// candidates carry no scorable stats worth ranking by: every candidate is
// equipped in turn alongside the rest of picks and run through a real DPS
// sim, the highest measured becomes picks[slot] (the next its runner-up),
// and the winner's gain over the empty slot is measured at escalating
// precision (trinketAdaptiveGain) for report.go's low-gain gate. The
// trinket slots and the relic slot (relics.go) share it; it returns a new
// map rather than mutating picks.
func rankByMeasuredGain(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, slot string, candidates []scored) (map[string]slotPick, []string) {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}

	type measured struct {
		item   scored
		dps    float64
		stdErr float64
	}
	var results []measured
	var notes []string
	for _, c := range candidates {
		gear := swapSlot(picks, slot, c.ID, false)
		req := plainRequest(spec, bandCharacter("trinket-rank", race, classSlug, spec.Spec, level, talents, gear), trinketRankIterations, verifySeed)
		dps, stdErr, err := runner.RunPlainDPSWithError(req)
		if err != nil {
			notes = append(notes, formatTrinketRankError(slot, c, err))
			continue
		}
		results = append(results, measured{item: c, dps: dps, stdErr: stdErr})
	}
	if len(results) == 0 {
		return out, notes
	}
	sort.SliceStable(results, func(i, j int) bool { return results[i].dps > results[j].dps })

	// Hybrid sweep, bis-ranker-integrity-4 lane, item 4: Serenity Field
	// (a Spirit self-buff the engine does not model, published with its
	// own effect_unmodelled: true) beat a real combat trinket at band 60
	// ret/enhancement by a margin smaller than this 100-iteration
	// tournament's own noise - a coin flip decided the slot, not a real
	// DPS difference. An unmodelled candidate's own measured "DPS" is
	// exactly the stat-only DPS a modelled candidate ALSO gets credited
	// with (score.go's own doc: this tournament runs a real sim either
	// way, so a trinket with no proc/on-use the engine can fire measures
	// no differently from one whose stats alone happen to be worth the
	// same amount) - the two are not being compared unfairly, only
	// noisily, at this iteration count. The fix: a modelled trinket
	// keeps the slot over a currently-leading unmodelled one unless the
	// unmodelled one clears beatsByMargin's own 1% bar (verify.go, the
	// same bar this command already trusts for every other close-call
	// swap decision), rather than raising trinketRankIterations to shrink
	// the noise floor by brute force - the lane report names the reason
	// this one was picked without a side-by-side timing run: doubling or
	// more the per-candidate iteration count multiplies the cost of
	// EVERY trinket tournament this command runs, every band, every
	// faction, every written spec, every night, to fix a failure mode
	// this shared, already-trusted margin closes for free.
	//
	// This lane's brief (bis-ranker-integrity-6), item 3: the ORIGINAL
	// version of this margin check searched arbitrarily far down
	// results for the first MODELLED candidate, however weak, and
	// demoted results[0] to it whenever results[0] did not clear the
	// margin over THAT candidate - hunter-beast-mastery band 60
	// Alliance's own trinket2 repro: Thunderbrew's Boot Flask (+8
	// Spirit, a real, always-applied STAT - not its own separate,
	// unrelated, unimplemented "drunken fire breath" effect_text) and
	// Frozen Heart of the Mountain (+9 Hit, same story) both measured
	// real, positive gains (0.77 and 0.47 DPS) from ordinary stats the
	// engine unconditionally simulates - but BOTH also happen to carry
	// an unrelated effect_text the engine does not implement, so the
	// old search skipped past both of them looking for a "modelled"
	// candidate and landed on Darkmoon Card: Heroism, whose own
	// IMPLEMENTED effect is a pure self-heal (zero DPS relevance) -
	// measuring EXACTLY the no-trinket baseline. The slot then
	// published empty (report.go's trinketLowGain: 0.00 < the 0.05
	// floor), discarding two real, better, stat-driven trinkets in
	// favour of a demonstrably worse one, purely because "modelled"
	// was being used as a proxy for "trustworthy" when the candidate's
	// own measured number was never in question - only an UNMODELLED
	// EFFECT's contribution is unverifiable; a candidate's plain STATS
	// are always faithfully simulated regardless of what else its
	// effect_text says (score.go's own doc makes exactly this point
	// about score(), and it is equally true of a real engine sim).
	//
	// The margin check that follows now only ever compares results[0]
	// against results[1] - the ONE case the Serenity Field finding
	// above actually describes ("beat a real combat trinket... by a
	// margin smaller than this tournament's own noise"): an unmodelled
	// leader's own measured lead over the VERY NEXT candidate is close
	// enough that the leader's own extra, unmodelled effect_text could
	// be the entire (unverifiable) reason for it, so a real, modelled
	// runner-up right behind it is trusted instead. When results[1] is
	// ALSO unmodelled (this lane's own repro), there is no modelled
	// candidate immediately behind results[0] to defer to at all - both
	// numbers are equally real measurements of equally-real stats, so
	// the plain highest one wins, exactly as it would if neither
	// candidate carried an effect_text in the first place.
	winner := 0
	if len(results) > 1 && trinketEffectUnmodelled(results[0].item.candidate) && !trinketEffectUnmodelled(results[1].item.candidate) {
		if !beatsByMargin(results[0].dps, results[1].dps) {
			winner = 1
		}
	}
	// runnerUp is whichever OTHER candidate has the next-highest measured
	// dps - results is sorted descending, so the first index that is not
	// winner is already that candidate, whether winner stayed at 0 (the
	// ordinary case) or moved to demote an unmodelled results[0] (in
	// which case that demoted candidate - real top measured dps, just
	// not enough to clear the margin - is exactly the runner-up a
	// reader most wants to see).
	runnerUpIdx := -1
	for i := range results {
		if i != winner {
			runnerUpIdx = i
			break
		}
	}

	// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 2: every
	// candidate's own dps above is the full SET's absolute DPS wearing
	// it, always positive regardless of whether the trinket itself does
	// anything at all - report.go's old zero-value check read
	// MeasuredDPS > 0 as "this trinket was really measured, so it must
	// carry real value", which is true of every trinket a real trinket-
	// less character's own gear already produces hundreds of DPS
	// without. baselineGear drops this slot entirely (swapSlot's own
	// itemID-0 shape - see its doc) so baselineDPS is what THIS set
	// produces with NOTHING in the trinket slot at all; every
	// candidate's own real GAIN is its own dps minus this one shared
	// number. A failed baseline sim is logged and every candidate below
	// is left with GainMeasured false (report.go's zero-value gate only
	// trusts a gain it actually has - see scored.GainMeasured's own
	// doc): this pass already has a real MeasuredDPS ranking from
	// results above, so a baseline failure costs the gain check, not
	// the ranking itself.
	baselineGear := swapSlot(picks, slot, 0, false)
	baselineReq := plainRequest(spec, bandCharacter("trinket-rank-baseline", race, classSlug, spec.Spec, level, talents, baselineGear), trinketRankIterations, verifySeed)
	baselineDPS, baselineStdErr, baselineErr := runner.RunPlainDPSWithError(baselineReq)
	if baselineErr != nil {
		notes = append(notes, slot+": measuring the no-trinket baseline failed: "+baselineErr.Error())
	}

	// gainStdErr is MeasuredGainDPS's own standard error: the candidate
	// run and the baseline run are two independent sims, so their
	// difference's error is the two combined in quadrature (the
	// ordinary rule for the error of a difference of independent
	// means) - this lane's brief, item 2.
	gainStdErr := func(candidateStdErr float64) float64 {
		return math.Sqrt(candidateStdErr*candidateStdErr + baselineStdErr*baselineStdErr)
	}

	best := results[winner].item
	// This lane's brief, item 7: best/runnerUp's own MeasuredDPS
	// (scored's own doc) records the real, per-candidate full-set DPS
	// this very tournament measured - a trinket has no scorable stats
	// at all (this file's own package doc), so score()'s Score field
	// stays whatever near-zero value it always was; buildReport reads
	// MeasuredDPS, not Score, to decide what this row publishes.
	best.MeasuredDPS = results[winner].dps
	if baselineErr == nil {
		gainDPS := results[winner].dps - baselineDPS
		stdErr := gainStdErr(results[winner].stdErr)
		// This lane's brief (bis-ranker-integrity-12), item 1: the
		// ordering pass above only ever runs at trinketRankIterations
		// (100) - enough to ORDER the shortlist, but the WINNER's own
		// gain vs baseline is what trinketGainSignificant (weights.go)
		// gates the whole slot on, and a gain sitting close to that bar
		// can swing across it purely from noise: two otherwise-identical
		// reruns of this exact command against the live engine (this
		// lane's own repro, paladin-retribution band 60) measured
		// Second Wind's own set DPS at 186.96 and 188.28 with nothing
		// else different, and the committed nightly's own Alliance/Horde
		// pair for the same band+slot picked two DIFFERENT trinkets
		// (Second Wind / Burst of Knowledge) for exactly this reason. A
		// specifically high-variance spec can fail EVERY trinket in a
		// band this way at the flat iteration count (controller-flagged:
		// mage-fire band 60, both factions, both trinket slots -
		// Ignite's own variance keeps the gain inside 2x its own error
		// for every shortlisted candidate). trinketAdaptiveGain re-measures
		// only this ONE winning candidate, at escalating iteration
		// counts, until its own error is comfortably below its own gain
		// or a ceiling is reached - see its own doc for why only the
		// winner, not the whole shortlist.
		if refinedGain, refinedStdErr, err := trinketAdaptiveGain(runner, spec, race, classSlug, level, talents, picks, slot, best.ID, gainDPS, stdErr); err == nil {
			gainDPS, stdErr = refinedGain, refinedStdErr
		} else {
			notes = append(notes, slot+": adaptive gain re-measurement for "+best.Name+" failed, keeping the "+strconv.Itoa(trinketRankIterations)+"-iteration measurement: "+err.Error())
		}
		best.MeasuredGainDPS = gainDPS
		best.MeasuredGainStdErr = stdErr
		best.GainMeasured = true
	}
	sp := slotPick{Item: &best}
	if runnerUpIdx != -1 {
		runnerUp := results[runnerUpIdx].item
		runnerUp.MeasuredDPS = results[runnerUpIdx].dps
		if baselineErr == nil {
			runnerUp.MeasuredGainDPS = results[runnerUpIdx].dps - baselineDPS
			runnerUp.MeasuredGainStdErr = gainStdErr(results[runnerUpIdx].stdErr)
			runnerUp.GainMeasured = true
		}
		sp.RunnerUp = &runnerUp
	}
	out[slot] = sp
	return out, notes
}

func formatTrinketRankError(slot string, c scored, err error) string {
	return slot + ": ranking candidate " + c.Name + " failed: " + err.Error()
}

// measureTrinketGain is rankTrinketSlot's own per-candidate measurement
// (a candidate run, a no-trinket baseline run, then
// trinketAdaptiveGain's own escalation) pulled out into its own
// function so a caller outside the ordinary shortlist tournament -
// reconcileFactionTrinkets (faction_trinkets.go, this lane's brief
// bis-ranker-integrity-15) - can measure exactly one candidate's own
// gain on a character (picks/race/talents) it was never shortlisted
// for, at the same precision rankTrinketSlot's own winner gets,
// without re-running that whole tournament: this lane's brief's own
// words, "a contained post-pass, not a cache threaded through the
// tournaments" - rankTrinketSlot and trinketShortlist themselves are
// unchanged by this lane.
//
// itemID, not a scored candidate, is the only input this needs beyond
// the character context: the candidate run only ever equips an item
// id (swapSlot's own contract), so the caller's own copy of the
// candidate (whichever faction's bySlot it came from) never has to
// travel through this function at all.
func measureTrinketGain(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, slot string, itemID int) (absoluteDPS, gainDPS, gainStdErr float64, err error) {
	gear := swapSlot(picks, slot, itemID, false)
	req := plainRequest(spec, bandCharacter("trinket-reconcile", race, classSlug, spec.Spec, level, talents, gear), trinketRankIterations, verifySeed)
	dps, stdErr, err := runner.RunPlainDPSWithError(req)
	if err != nil {
		return 0, 0, 0, err
	}
	baselineGear := swapSlot(picks, slot, 0, false)
	baselineReq := plainRequest(spec, bandCharacter("trinket-reconcile-baseline", race, classSlug, spec.Spec, level, talents, baselineGear), trinketRankIterations, verifySeed)
	baselineDPS, baselineStdErr, err := runner.RunPlainDPSWithError(baselineReq)
	if err != nil {
		return 0, 0, 0, err
	}
	gainDPS = dps - baselineDPS
	gainStdErr = math.Sqrt(stdErr*stdErr + baselineStdErr*baselineStdErr)
	if refinedGain, refinedStdErr, adaptErr := trinketAdaptiveGain(runner, spec, race, classSlug, level, talents, picks, slot, itemID, gainDPS, gainStdErr); adaptErr == nil {
		gainDPS, gainStdErr = refinedGain, refinedStdErr
	}
	return dps, gainDPS, gainStdErr, nil
}

// trinketGainAdaptivePrecisionDivisor is trinketAdaptiveGain's own
// target: keep escalating the winning candidate's own gain measurement
// until its combined standard error sits under gain/4 - this lane's
// brief (bis-ranker-integrity-12), item 1, verbatim ("stderr <
// gain/4"). Comfortably inside trinketGainSignificanceMultiplier's own
// 2x significance bar (weights.go), so the escalation loop is
// confirming the gain, not merely chasing the same bar it started
// short of.
const trinketGainAdaptivePrecisionDivisor = 4.0

// trinketGainAdaptiveIterationCeiling bounds how far trinketAdaptiveGain
// will escalate a single winning candidate's own gain measurement -
// weights.go's own trinketGainSignificanceMultiplier doc already found
// that even a 20x rerun (100 to 2000 iterations) of a genuinely
// zero-relevance trinket does not shrink its gain toward zero (an
// engine-side question, not this loop's), so this loop only ever
// existing to shrink the ERROR, not chase a gain that is not really
// there, a 2000-iteration ceiling matches that same finding rather
// than inventing a new number.
const trinketGainAdaptiveIterationCeiling = 2000

// trinketAdaptiveGain re-measures winnerID's own gain over the
// no-trinket baseline (rankTrinketSlot's own call site: the ordering
// pass's WINNER only, never every shortlist candidate, so the added
// cost is one slot's one item per band per faction, not the whole
// tournament) at increasing iteration counts - doubling each round,
// capped at trinketGainAdaptiveIterationCeiling - until the gain's own
// combined standard error clears trinketGainAdaptivePrecisionDivisor,
// or a re-run errors (the caller already has a valid measurement from
// the ordering pass to fall back on in that case, so the error is
// returned rather than panicking a whole band).
//
// This is this lane's brief, item 1's actual mechanism finding: a gain
// sitting close to trinketGainSignificant's own bar can swing across
// it purely from sampling noise (Fire Ruby: 0.77 DPS gain against a
// 0.71 combined stdErr at the flat 100-iteration pass, hunter-beast-
// mastery band 50 - ranker-11's own repro) or from which faction
// happened to be simmed at all (this lane's own repro: two otherwise-
// identical reruns of paladin-retribution band 60 measured Second
// Wind's own set DPS at 186.96 and 188.28 with nothing else different,
// and the committed nightly's Alliance/Horde pair picked two different
// trinkets for exactly this reason) - and a specifically high-variance
// spec can fail EVERY trinket in a band this way at the flat iteration
// count (controller-flagged: mage-fire band 60, both factions, both
// trinket slots - Ignite's own variance). Escalating only the winner's
// own gain, rather than raising trinketRankIterations for the whole
// shortlist, keeps this fix's added cost to the one candidate that
// actually decides the slot.
func trinketAdaptiveGain(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, slot string, winnerID int, gainDPS, stdErr float64) (float64, float64, error) {
	iterations := trinketRankIterations
	for gainDPS > 0 && stdErr >= gainDPS/trinketGainAdaptivePrecisionDivisor && iterations < trinketGainAdaptiveIterationCeiling {
		iterations *= 2
		if iterations > trinketGainAdaptiveIterationCeiling {
			iterations = trinketGainAdaptiveIterationCeiling
		}
		gear := swapSlot(picks, slot, winnerID, false)
		req := plainRequest(spec, bandCharacter("trinket-rank-adaptive", race, classSlug, spec.Spec, level, talents, gear), iterations, verifySeed)
		candDPS, candStdErr, err := runner.RunPlainDPSWithError(req)
		if err != nil {
			return gainDPS, stdErr, err
		}
		baselineGear := swapSlot(picks, slot, 0, false)
		baseReq := plainRequest(spec, bandCharacter("trinket-rank-adaptive-baseline", race, classSlug, spec.Spec, level, talents, baselineGear), iterations, verifySeed)
		baseDPS, baseStdErr, err := runner.RunPlainDPSWithError(baseReq)
		if err != nil {
			return gainDPS, stdErr, err
		}
		gainDPS = candDPS - baseDPS
		stdErr = math.Sqrt(candStdErr*candStdErr + baseStdErr*baseStdErr)
	}
	return gainDPS, stdErr, nil
}
