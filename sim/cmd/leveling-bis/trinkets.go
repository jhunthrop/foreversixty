package main

import (
	"sort"
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

// trinketShortlist is rankTrinketSlot's real candidate pool: the union
// of topByItemLevel's top trinketTopN (highest item level - a rough
// proxy for "how good a trinket this tier of content dropped", the
// only ranking this file had before this lane), the top trinketTopN
// BY SCORE (list is already best-score-first - candidatesBySlot's own
// contract - so this is simply its own head, pair-mate-filtered), and
// (bis-ranker-integrity-6, item 2) every candidate whose own effect the
// engine actually implements (implementedEffectTrinkets, below) -
// deduplicated by item id.
//
// This lane's brief (bis-ranker-integrity, 2026-09-29), item 1's own
// mechanism finding: a plain-stat trinket that out-scores everything
// else in the pool (Neltharion's Tear, spell_power 44 + hit 20, against
// mage-fire's Naxxramas picks) can sit at a LOWER item level than five
// higher-ilvl on-use/proc trinkets from the same or a later raid tier,
// so topByItemLevel alone never gave it a single sim - the real
// tournament below crowned whichever proc trinket won among candidates
// that never had to face the best plain stat-stick in the pool at all.
// bySlot[slot]'s own score() ranking cannot crown the WINNER by itself
// either (score() has no notion of a proc - trinkets.go's own package
// doc), but it is exactly the list this file needs to make sure the
// best stat-based candidate gets a fair, real-sim shot alongside the
// best item-level-based ones, so whichever axis actually wins the real
// sim becomes the pick, and the loser becomes an honestly-Verified
// alternative instead of an unverified score() estimate the pick never
// actually faced (report.go's buildAlternatives, this lane's other
// fix).
// implementedEffectTrinkets returns every candidate in list whose own
// on-hit/on-use/proc effect the engine actually implements
// (hasImplementedEffect, rank.go) - this lane's brief
// (bis-ranker-integrity-6), item 2: Hand of Justice (11815, a real
// Blackrock Depths drop, effectids_generated.go's own table) never
// reached a single sim for any melee spec at any band, because its
// stat block is empty (score() has nothing to rank it by - trinkets.go's
// own package doc) AND five higher-item-level trinkets (Darkmoon Faire
// cards among them, ilvl 66-75 at band 60) always filled topByItemLevel
// first. An item the engine CAN measure a real, implemented effect for
// must reach the tournament regardless of its structured stats or its
// item level relative to the rest of the pool - the whole reason this
// file runs a real sim at all is to value exactly what score() cannot
// see, and item level is only ever a proxy for that, never a
// substitute for actually asking the engine. A trinket with no
// implemented effect (or no effect at all) is untouched by this
// bucket - Blackhand's Breadth (13965) is a real BRD drop too, but its
// own crit-chance proc is not in effectids_generated.go, so the engine
// could not measure it any better than score() already fails to; force-
// including it here would spend a real sim on a candidate this command
// still could not value, not fix anything this lane's brief asks for.
func implementedEffectTrinkets(list []scored) []scored {
	out := make([]scored, 0, len(list))
	for _, c := range list {
		if hasImplementedEffect(c.candidate) {
			out = append(out, c)
		}
	}
	return out
}

func trinketShortlist(list []scored, excludeID int, excludeName string) []scored {
	filtered := excludePairMate(list, excludeID, excludeName)
	byItemLevel := topByItemLevel(filtered, 0, "")
	byScore := filtered
	if len(byScore) > trinketTopN {
		byScore = byScore[:trinketTopN]
	}
	byImplementedEffect := implementedEffectTrinkets(filtered)
	seen := make(map[int]bool, len(byItemLevel)+len(byScore)+len(byImplementedEffect))
	out := make([]scored, 0, len(byItemLevel)+len(byScore)+len(byImplementedEffect))
	for _, group := range [][]scored{byItemLevel, byScore, byImplementedEffect} {
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
// on-hit/on-use/proc effect (EffectText non-empty) the engine does NOT
// actually implement (rank.go's effectImplemented) - Serenity Field's
// own Spirit self-buff, for one. A candidate with no effect at all
// (EffectText == "") is NOT unmodelled in this sense: its whole value
// is stats, which this tournament's own real sim already measures
// exactly as faithfully as any other stat-only trinket.
func trinketEffectUnmodelled(c candidate) bool {
	return c.EffectText != "" && !hasImplementedEffect(c)
}

// rankTrinketSlot replaces picks[slot] with the engine-verified best of
// its top-item-level candidates (and the runner-up with the second
// best), leaving every other slot's pick untouched. It returns a new
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
func rankTrinketSlot(runner engineRunner, spec specInfo, race, classSlug string, level int, talents string, picks map[string]slotPick, bySlot map[string][]scored, slot string) (map[string]slotPick, []string) {
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}

	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok && picks[mate].Item != nil {
		mateID, mateName = picks[mate].Item.ID, picks[mate].Item.Name
	}
	candidates := trinketShortlist(bySlot[slot], mateID, mateName)
	if len(candidates) == 0 {
		return out, nil
	}

	type measured struct {
		item scored
		dps  float64
	}
	var results []measured
	var notes []string
	for _, c := range candidates {
		gear := swapSlot(picks, slot, c.ID, false)
		req := plainRequest(spec, bandCharacter("trinket-rank", race, classSlug, spec.Spec, level, talents, gear), trinketRankIterations, verifySeed)
		dps, err := runner.RunPlainDPS(req)
		if err != nil {
			notes = append(notes, formatTrinketRankError(slot, c, err))
			continue
		}
		results = append(results, measured{item: c, dps: dps})
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
	baselineDPS, baselineErr := runner.RunPlainDPS(baselineReq)
	if baselineErr != nil {
		notes = append(notes, slot+": measuring the no-trinket baseline failed: "+baselineErr.Error())
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
		best.MeasuredGainDPS = results[winner].dps - baselineDPS
		best.GainMeasured = true
	}
	sp := slotPick{Item: &best}
	if runnerUpIdx != -1 {
		runnerUp := results[runnerUpIdx].item
		runnerUp.MeasuredDPS = results[runnerUpIdx].dps
		if baselineErr == nil {
			runnerUp.MeasuredGainDPS = results[runnerUpIdx].dps - baselineDPS
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
