package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/request"
)

// finishedSetEpsilon is how close a row's own measured DPS number must
// be to this band's final, published SetDPS to trust it as "measured
// on the finished set" rather than a stale greedy-fill snapshot - this
// lane's brief, item 5. Both numbers come from the exact same fixed-
// seed sim request (verifySeed) whenever they really do describe the
// same gear, so they land bit-identical in practice; this only guards
// against float summation order, not real DPS differences.
const finishedSetEpsilon = 0.01

// dpsComparisonPhrase is the shared "how much, and whether the two raw
// numbers are trustworthy" clause both SwapNote branches in buildReport
// use - this lane's brief, item 5: paladin-retribution 60 Alliance
// published a header Set DPS of 182.9 while five different rows each
// named a different "X vs Y set DPS" pair, none of them 182.9, because
// each pair was measured before the OTHER four slots' own promotions
// were known - a player reads that as five contradictions with the
// header, not five honest partial snapshots. delta (always real - see
// slotRow.DPSDelta's own doc) is published unconditionally; the two
// raw numbers are added only when primaryDPS itself matches this
// band's own final setDPS, i.e. this really was the last word on the
// set's total, not a snapshot a later promotion superseded.
// bis-ranker-integrity-17 lane, player-review sweep 15/16: delta can be
// negative here (the "kept the pick" branch's own delta is
// sw.BaselineDPS - sw.SwapDPS, which is negative the moment the
// runner-up scored higher but not significantly enough to promote --
// warlock-destruction band 60 Alliance legs' own repro, -2.65). A
// literal "+" prefix in front of "%.1f" does not know that: "%.1f"
// already prints its own "-" for a negative value, so the two
// combined printed "+-2.7 DPS". "%+.1f" is Go's own sign flag -- it
// prints "+" for a non-negative value and "-" for a negative one,
// never both -- so it is the one formatting verb that is correct for
// both signs without a caller having to branch on delta's own sign
// first.
func dpsComparisonPhrase(delta, primaryDPS, otherDPS, setDPS float64) string {
	if math.Abs(primaryDPS-setDPS) <= finishedSetEpsilon {
		return fmt.Sprintf("%+.1f DPS (%.1f vs %.1f set DPS)", delta, primaryDPS, otherDPS)
	}
	return fmt.Sprintf("%+.1f DPS over it", delta)
}

// slotRow is one slot's line in a band's report: the JSON and the
// markdown table share this shape.
type slotRow struct {
	Slot       string `json:"slot"`
	ItemID     int    `json:"item_id,omitempty"`
	ItemName   string `json:"item_name,omitempty"`
	Source     string `json:"source,omitempty"`
	SourceKind string `json:"source_kind,omitempty"`
	// Score is score()'s own stat-weight estimate (bandReport.ScoreUnit,
	// reference-stat points) - published ONLY for a pick score() itself
	// decided. Omitted (0) whenever SimDPS below is set: this lane's
	// brief, item 7 ("one number per row") - the two are different
	// units a slot can be decided by, and publishing both under one
	// name is what let a reviewer compare an unrelated unit across two
	// rows and call it "the alternative outscores the verified pick"
	// (163 slots, the melee sweep's own systemic finding).
	Score float64 `json:"score,omitempty"`
	// SimDPS is the real, engine-measured full-set DPS that decided
	// THIS pick, published in place of Score whenever a real sim - not
	// score() - made the call: rankTrinketSlot (every trinket, always),
	// rankSlotWithEffects (a slot with an implemented-effect
	// candidate), trySetCompletion (a slot a winning set trial
	// replaced), or a swap promotion (verify.go's applySwaps). Sourced
	// from pk.Item.MeasuredDPS (scored's own doc) - this lane's brief,
	// item 7.
	SimDPS   float64 `json:"sim_dps,omitempty"`
	Verified bool    `json:"verified"`
	SwapNote string  `json:"swap_note,omitempty"`
	// DPSDelta is this row's own pick's measured DPS advantage over the
	// one comparator a real swap sim actually measured it against, in
	// THAT SAME run - this lane's brief, item 5: SimDPS/SwapNote's own
	// absolute numbers are each the full set's total AT THE MOMENT that
	// slot was decided (greedy fill measures one slot at a time against
	// whatever the OTHER slots already were, then keeps deciding more
	// slots afterward) - paladin-retribution 60 Alliance published a
	// header Set DPS of 182.9 while five different rows each named a
	// different, lower "X vs Y set DPS" pair (180.8, 176.4, 178.4,
	// 173.5, 174.2), none of them the header's own number, because five
	// different slots each promoted against a baseline measured before
	// the OTHER four promotions were known. The two raw numbers inside
	// one swap comparison are always internally consistent (SwapNote's
	// own doc), but comparing either one against the band's own final
	// SetDPS is not - DPSDelta is the one number that stays true
	// regardless: the gap SwapNote's own comparison measured, in that
	// one run, independent of anything decided before or after. Omitted
	// when this row was never compared to anything a sim actually
	// measured (a plain score()-decided pick with no swap result).
	DPSDelta *float64 `json:"dps_delta,omitempty"`
	// EffectUnmodelled is true when the picked item carries an effect
	// (effect_text, or a client proc/use/equip-behaviour spell: rank.go's
	// carriesEffect) the engine does NOT implement (effectids_generated.go)
	// -- this lane's brief, item 3: such a candidate is still scored on
	// its plain stats (score.go never saw the effect either way), but
	// the page shows "proc not simulated" on it rather than letting the
	// Score/Verified columns imply the whole item was accounted for.
	// False (omitted) for a candidate with no effect at all, and
	// for one whose effect IS implemented -- see rank.go's
	// hasImplementedEffect, the single predicate both this flag and the
	// effect-verification pass itself read.
	EffectUnmodelled bool `json:"effect_unmodelled,omitempty"`
	// SimStatus is notInSimReason ("not_in_sim") when this row's own
	// pick carries an item id this build's embedded item database does
	// not have (pk.Item.NotInSimDB, data.go's own doc) - this lane's
	// brief, item 1, generalising
	// rank.go's own effectVerifiedInSim gate (which already caught this
	// for an implemented-effect candidate's EffectUnmodelled flag,
	// above) to EVERY slot and every pass that can set MeasuredDPS
	// (trinkets.go's tournament, rank.go's effect tournament, sets.go's
	// set-completion trial, verify.go's swap promotion): simdb.Attach/
	// AttachWeights' own UnequipUnknown silently strips such an item
	// from the character before ANY sim this command runs ever builds
	// it, so nothing downstream ever actually measured it, whatever
	// pk.Item.MeasuredDPS below claims to have found. Forces simDecided
	// off (below), so the row keeps its plain score() estimate and
	// Score is never zeroed - "the pick stays score-decided" (this
	// lane's brief) - and forces Verified false and every sim-only field
	// (SimDPS, DPSDelta, SwapNote) empty, in the switch below, so the
	// row can never read as tested when it never was. Empty (omitted)
	// for every other row, including one whose pick was genuinely never
	// simmed at all (score()-only, no tournament/swap ever touched it) -
	// that case already publishes Verified: true with no SimDPS/
	// DPSDelta, which is not a false claim (see erroredSlots/default
	// case below): this field exists only for the specific, provably
	// false claim a stripped item's own MeasuredDPS would otherwise
	// make.
	SimStatus string `json:"sim_status,omitempty"`
	// row's pick (slotPick.Ties, pick.go's own doc; this lane's brief,
	// defect 4) - "or Blackwater Cutlass" as populated alternatives the
	// page can print beside a pick that was really an arbitrary
	// lowest-id tie-break rather than a unique best. Empty for a slot
	// the trinket/effect/set-completion passes decided by a real sim
	// instead of score() (see slotPick.Ties's own doc). Kept exactly as
	// it was before Alternatives existed (below) - a consumer already
	// reading Ties keeps working unchanged.
	Ties []tieAlternative `json:"ties,omitempty"`
	// Alternatives is up to alternativesLimit candidates a player who
	// cannot get this row's pick can fall back on: this lane's brief
	// (owner defect A, 2026-09-29) - warrior-arms horde band 20's
	// main_hand published Forsaken Greataxe with no record that Smite's
	// Mighty Hammer (a Deadmines drop, item 7230) was ever considered,
	// leaving a player who cannot or will not run that quest nothing to
	// fall back on. Candidates are Ties (score identical to the pick,
	// this lane's brief addendum) plus the next-best sourced candidates
	// by score() (buildAlternatives, this file), ordered by dps_delta
	// (alternativeRow.DPSDelta) descending - see that function's own
	// doc for exactly what "next-best" excludes (the pick itself, its
	// pair-mate, anything already listed as a tie) and why the final
	// order is not simply "ties, then score order" (owner review, tenet
	// 8, 2026-09-29: a runner-up verify.go actually simmed against the
	// pick can score higher yet measure worse, and must not publish a
	// dps_delta that outranks a tie or the pick it lost to).
	Alternatives []alternativeRow `json:"alternatives,omitempty"`
	// EmptyReason explains a slot with no ItemID at all beyond "nothing
	// eligible/sourced existed" (band.go's own NoSourceCount already
	// covers that case): "no_dps_value" means at least one candidate WAS
	// eligible and sourced, but every one of them contributed exactly
	// zero this spec's own score() could measure (score.go's weighted-
	// stat-plus-weapon-dps total), or - since bis-ranker-integrity-3,
	// 2026-09-29, this lane's brief item 2 - a trinket/relic whose own
	// real, engine-measured GAIN over an empty slot (GainMeasured/
	// MeasuredGainDPS, trinkets.go) never cleared
	// trinketZeroGainThresholdDPS: Sentinel's Medallion (Agility/Stamina)
	// for a mage, Rune of Perfection (spell penetration/stamina) for a
	// warrior, Ankh of Life (pure Spirit) for any physical-damage spec -
	// each published as "verified" BiS before its own fix, contributing
	// nothing any weight or real sim measured. "effect_not_modelled"
	// (effectNotModelledReason) is the OTHER empty reason a relic can
	// carry - not "measured zero", but "this relic's one real selling
	// point is its engraved effect and the engine cannot simulate it at
	// all", which score()'s plain stat fallback was never designed to
	// answer either way (isRelicCandidate, eligible.go). Never set for a
	// slot a real swap sim actually promoted (SwapNote non-empty - a
	// real, measured DPS gain, whatever score() says) or for a candidate
	// carrying an unmodelled effect outside a relic slot (EffectUnmodelled,
	// above - its own real value might not be zero at all, score() simply
	// cannot see it, and no relic-specific "cannot know at all" applies).
	EmptyReason string `json:"empty_reason,omitempty"`
	// LowValue is true for a weapon-slot pick (band.go's weaponSlots)
	// score() still measured at exactly 0 even after every real fix
	// this lane's brief, item 1 asks for - bis-ranker-integrity-2 lane:
	// a weapon slot must never publish empty (EmptyReason above is
	// never set for one), so a zero-scoring pick here is the best
	// defensible fallback pick.go's own promoteLowValueWeapon could
	// find (a candidate carrying one of the spec's own weight-stat
	// keys, else the highest item level) rather than an honest
	// "verified BiS", and the page should say so instead of implying
	// this weapon was actually the strongest option among real
	// contenders. Never set for a non-weapon slot, which still empties
	// exactly as before (EmptyReason == noDPSValueReason).
	LowValue bool `json:"low_value,omitempty"`
	// FactionNote explains the one legitimate reason a faction-neutral
	// trinket's own published pick differs between Alliance and Horde
	// (reconcileFactionTrinkets, faction_trinkets.go, this lane's brief
	// bis-ranker-integrity-15): a real racial stat difference this
	// faction's own re-measurement of the OTHER faction's pick found,
	// not the ranker's own per-faction sim noise landing on opposite
	// sides of a shared bar (twelfth sweep's defect, ranker-14). Empty
	// for every row that pass never touches, including a trinket slot
	// that simply was never faction-neutral to begin with.
	FactionNote string `json:"faction_note,omitempty"`
	// SetBonus is present only on a slot trySetCompletion (sets.go) adopted
	// for the set bonus it completes: the piece is worn for the bonus, not
	// for its own stats. Absent from every other row.
	SetBonus *setBonusNote `json:"set_bonus,omitempty"`
	// LabelSuffix disambiguates this row's own ItemName from a
	// DIFFERENT item sharing the exact same display name elsewhere in
	// this same row (an Alternatives entry, almost always - see
	// applyLabelSuffixForNameCollisions) - this lane's brief
	// (bis-ranker-integrity-16), item 6: warlock-destruction band 60
	// Alliance's own legs pick, "Sentinel's Silk Leggings" (id 237815,
	// ilvl 78, Forever's own reissue), sits beside an alternative ALSO
	// named "Sentinel's Silk Leggings" (id 22752, ilvl 65, the real
	// vanilla item, both sold by Illiyana Moonblaze) with nothing on
	// the row itself telling the two apart beyond a bare id a player
	// cannot look up in-game. "(ilvl 78)" here, the item's own real
	// item level, is the one distinguishing fact every other row on
	// the page already carries a tooltip for. Empty for every row
	// whose name is unique within its own Alternatives list, which is
	// nearly all of them.
	LabelSuffix string `json:"label_suffix,omitempty"`
}

// noDPSValueReason is EmptyReason's own published value for this lane's
// brief item 2 - a named constant so buildReport's own check and any
// consumer testing for it read the identical string.
const noDPSValueReason = "no_dps_value"

// effectNotModelledReason is EmptyReason's own value for a relic
// (libram/idol/totem) whose one real selling point - its engraved
// effect - the engine does not implement at all (bis-ranker-
// integrity-3, 2026-09-29, this lane's brief item 2): distinct from
// noDPSValueReason (which means "measured/estimated, and the answer
// really is zero") because this case is the opposite - the item's
// real value is specifically UNKNOWN, not zero, and publishing it as
// BiS on a plain-stats fallback score() was never designed to value
// (druid-balance's own Idol of the Huntress, an "Improved Swipe" -
// a Feral rune - engrave with nothing else worth scoring) would show
// an unverified guess as fact. The lane report lists every relic id
// this reason fires for, so the effect can be added to
// effectids_generated.go.
const effectNotModelledReason = "effect_not_modelled"

// notInSimReason is slotRow.SimStatus's own value for a pick this
// build's embedded item database does not carry - see that field's
// own doc (this lane's brief, item 1).
const notInSimReason = "not_in_sim"

// trinketZeroGainThresholdDPS is no longer the production gate (see
// history below) - it survives only as the test suite's own
// descriptive floor for "obviously zero" versus "obviously real" gain
// fixtures. bis-ranker-integrity-11's brief, item 2 replaced the
// production check with trinketGainSignificant (weights.go):
// trinketLowGain (below) now asks whether a trinket's own
// MeasuredGainDPS clears trinketGainSignificanceMultiplier times its
// own MeasuredGainStdErr (trinkets.go's engineRunner.RunPlainDPSWithError,
// the candidate and no-trinket-baseline runs' errors combined in
// quadrature), not a flat absolute DPS floor.
//
// History (bis-ranker-integrity-3 lane report, 2026-09-29): a flat
// 0.05 DPS floor reliably caught a trinket whose ONLY stat is
// completely inert for this spec (Rune of Perfection/Rune of Duty/
// Onyxia Blood Talisman), which measures bit-identical to the
// baseline every time, but could not resolve small, genuinely real
// gains (Frozen Heart of the Mountain's own +9 Hit, 0.47 DPS at band
// 60) from equally-small sim noise at trinketRankIterations (100) -
// nor, going the other way, could it catch a trinket whose own
// measured "gain" was PURE noise around an effect with zero true DPS
// relevance for this spec (Sanctified Orb's "Restores 340 Mana" for a
// spec with no mana resource, Fire Ruby's mage-only Fire Ward/Fire
// Blast interaction for any other class) when that noise happened to
// read above 0.05 by chance - bis-ranker-integrity-11's own repro:
// both won zero-delta ties against real, weighted-stat trinkets this
// way. trinketGainSignificant resolves both: a real small gain still
// clears its own (wider) error bar, and a genuinely-inert effect's
// noise only clears that bar as often as chance allows at the 2x
// multiplier's own confidence level, not every time it happens to
// land positive (see trinketGainSignificanceMultiplier's own doc for
// the dogfooded numbers that picked 2x over a bare 1x).
const trinketZeroGainThresholdDPS = 0.05

// twoHandEquippedReason is EmptyReason's own value for an off_hand
// slot with no pick at all because main_hand equipped a two-hander -
// this lane's brief, item 8: every empty slot carries an
// empty_reason, and this is the one case pick.go's own
// enforceTwoHandOffHandInvariant/off_hand-under-two-hander rule
// leaves behind (hunter-marksmanship band 40 alliance's own off_hand
// published neither an item nor a reason before this fix).
const twoHandEquippedReason = "two_hand_equipped"

// noSourcedItemReason is EmptyReason's own value for a slot with no
// pick at all because pick()'s own candidate list for it (bySlot,
// after any spec-specific narrowing: dagger-only restriction,
// pairing exclusion, ...) was genuinely empty - nothing sourced this
// band could put in the slot, as distinct from noDPSValueReason
// (something WAS sourced, it simply scored nothing) - this lane's
// brief, item 8.
const noSourcedItemReason = "no_sourced_item"

// swapNoteRunnerUpName is pk.RunnerUp.Name, with its own item level
// appended when it exactly matches pk.Item's own name under a
// different id - this lane's brief (bis-ranker-integrity-16), item 6:
// warlock-destruction band 60 Alliance's own legs swap_note read "beat
// the scored pick Sentinel's Silk Leggings (id 22752) in the sim" next
// to a pick ALSO named "Sentinel's Silk Leggings" (id 237815) - a
// sentence that reads as the pick beating itself, the id in
// parentheses being the only (easy to miss) fact telling the two
// apart. "(ilvl 65)" here is the same disambiguator
// applyLabelSuffixForNameCollisions publishes on the row itself.
func swapNoteRunnerUpName(pk slotPick) string {
	if pk.Item != nil && pk.RunnerUp != nil && pk.Item.Name == pk.RunnerUp.Name &&
		pk.Item.ID != pk.RunnerUp.ID && pk.RunnerUp.ItemLevel > 0 {
		return fmt.Sprintf("%s (ilvl %d)", pk.RunnerUp.Name, pk.RunnerUp.ItemLevel)
	}
	return pk.RunnerUp.Name
}

// emptyReasonForNilPick names why slot has no pick at all (pk.Item ==
// nil): pick()'s own loop (pick.go) only ever leaves a slot without
// an Item because (a) main_hand equipped a two-hander, so off_hand is
// deliberately left empty (pick()'s own "off_hand" case, and
// enforceTwoHandOffHandInvariant's later re-assertion of the same
// rule), or (b) this slot's own candidate list was empty once every
// spec-specific narrowing pick() applies ran (a dagger-only
// restriction leaving nothing, or a pairing exclusion removing the
// one candidate that existed) - band.go's Coverage row for the slot
// tells the two apart from a genuine "nothing was ever sourced here"
// only imprecisely (case (b) can still show a nonzero Eligible/
// Sourced count, since pick()'s own narrowing runs AFTER buildBandPool
// computed Coverage), so both land on noSourcedItemReason: either way,
// nothing usable existed for THIS spec's own picking pool, the same
// practical fact a reader needs regardless of which of the two caused
// it.
func emptyReasonForNilPick(slot string, picks map[string]slotPick) string {
	if slot == "off_hand" {
		if mh := picks["main_hand"].Item; mh != nil && mh.TwoHand {
			return twoHandEquippedReason
		}
	}
	return noSourcedItemReason
}

// publishableFactionNote is pk.FactionNote, or "" when that note's own
// wording names pk.Item as "this faction's own pick"
// (FactionNoteNeedsPick, pick.go's own doc) but the row about to
// publish no longer shows Item as its pick at all - this lane's brief
// (bis-ranker-integrity-16), item 3: shaman-elemental's Alliance band
// 50 trinket2 published empty (report.go's own trinketLowGain gate)
// while still carrying a note calling its own hidden item "alliance's
// own pick Molten Heart of the Mountain", an item that then named
// nowhere in the row at all. Every site that rebuilds row as an empty
// slotRow (relicEffectUnmodelled, trinketLowGain, the default
// zero-value case) calls this instead of reading pk.FactionNote
// directly, so a note only ever publishes while its own claim still
// holds; tenet 8: describe what actually publishes, or say nothing.
func publishableFactionNote(pk slotPick) string {
	if pk.FactionNoteNeedsPick {
		return ""
	}
	return pk.FactionNote
}

// tieAlternative is one equally-scored item slotRow.Ties names.
type tieAlternative struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

// reconcileTies is this lane's brief (bis-ranker-integrity-14), item 3:
// hunter-beast-mastery/hunter-marksmanship band 20 hands published
// Serpent Gloves as TYING Gloves of the Fang (pick.go's own
// slotPick.Ties - an exact SCORE match, computed before any real sim
// ever ran) while the very same row's own Alternatives entry for
// Gloves of the Fang showed dps_delta -0.08 - a real, sim-measured
// loss, from buildAlternatives' own swap-override (Gloves of the Fang
// is also pk.RunnerUp here, the one candidate verify.go's real swap
// pass actually tested against the pick). A score tie and a measured,
// above-noise loss cannot both be true; the tie label and the
// alternative's own delta must come from the same measurement.
//
// Alternatives is the stronger evidence whenever it exists for a given
// item id: buildAlternatives' own swap-override (its doc) only ever
// replaces a tie's default 0 delta with swapMeasuredDelta's real
// result when verify.go actually simmed that exact candidate -
// everything else in Alternatives for a genuine tie still publishes
// scoreDelta*referenceDPSPerPoint, which is exactly 0 for an identical
// score (toRow's own arithmetic). So a tie whose own alternatives row
// carries a nonzero DPSDelta was contradicted by a real measurement -
// dropped here instead of published as a tie the sim itself disproved.
// A tie with no matching alternatives row at all (buildAlternatives'
// own realAlternative/excluded gates can leave one out) or one whose
// alternatives row is still exactly 0 remains a genuine tie.
//
// What decides which of several genuinely-tied candidates the PICK
// itself is remains candidatesBySlot's own dead-stat-count-then-
// item-level rule (pick.go, ranker-12) - this function only reconciles
// the REPORTED tie list against the real evidence Alternatives already
// carries; it never changes which item is the pick.
func reconcileTies(ties []tieAlternative, alternatives []alternativeRow) []tieAlternative {
	if len(ties) == 0 {
		return ties
	}
	deltaByID := make(map[int]float64, len(alternatives))
	for _, a := range alternatives {
		deltaByID[a.ItemID] = a.DPSDelta
	}
	out := make([]tieAlternative, 0, len(ties))
	for _, tie := range ties {
		if delta, ok := deltaByID[tie.ItemID]; ok && delta != 0 {
			continue
		}
		out = append(out, tie)
	}
	return out
}

// alternativesLimit bounds how many candidates slotRow.Alternatives
// carries beyond the pick itself (this lane's brief: "the next best 3
// candidates by score after the pick") - enough for a player who
// cannot get the picked item's own source to see a genuine fallback,
// without ballooning the JSON with a slot's whole candidate pool.
const alternativesLimit = 3

// mdTieDisplayLimit bounds how many of a pick's tied items (slotRow.
// Ties) writeMarkdown names in its own Item cell before summarizing
// the rest as "and N more" - this lane's brief, item 4: the JSON's own
// `ties` field is uncapped (report.go publishes every real score tie,
// which a reader of the raw data may want in full), but the .md
// prototype page renders it inline in a single table cell, and a slot
// with 100+ ties (priest-shadow band 20 main_hand: 107/117 across
// factions) turned that cell into a wall of names longer than the
// rest of the table combined. 5, not alternativesLimit's 3: a true
// score tie is "equally good", a stronger claim than an ordinary
// runner-up alternative, so this cell keeps a little more room for it
// while still capping the pathological case.
const mdTieDisplayLimit = 5

// alternativeRow is one candidate slotRow.Alternatives names beyond
// the slot's own pick.
type alternativeRow struct {
	ItemID     int    `json:"item_id"`
	ItemName   string `json:"item_name"`
	SourceKind string `json:"source_kind"`
	Source     string `json:"source"`
	// DPSDelta is this row's own DPS gap against the pick - owner
	// review, tenet 8 (2026-09-29): the first cut of this field
	// published the raw score-unit delta under a "dps_delta" name with
	// nothing saying it was not actually DPS (warrior-arms horde band 20
	// read "Smite's Mighty Hammer -5.09" when the real gap is 0.23 DPS -
	// 5.09 SCORE points at this band's reference_dps_per_point of
	// 0.0444). bis-ranker-integrity-3 (2026-09-29), this lane's brief
	// item 1: a THIRD wow-player sweep still read this row's own score()
	// estimate (formerly published alongside DPSDelta as Score/
	// ScoreDelta) as "a positive gap against the pick", the exact
	// confusion those two fields existed to prevent - the field itself
	// carried no player meaning score()'s own units already didn't, and
	// has no test that ever asserted a consumer needed it. Removed
	// entirely: DPSDelta (already always <=0 for an unverified row, see
	// its own doc below) is the only per-alternative delta this JSON
	// publishes now, in one unit, always DPS. Score/score_unit stay on
	// the PICK (slotRow.Score) exactly as before - only the alternative
	// side of the comparison is gone.
	//
	// This lane's brief (bis-ranker-integrity, 2026-09-29): score() is
	// not the unit a trinket-rank/effect-rank/set-completion pick was
	// actually decided by (rankTrinketSlot, rankSlotWithEffects,
	// trySetCompletion each replace a slot's pick with a REAL sim's
	// winner, which score() cannot see the reason for at all - the exact
	// gap rank.go's own doc calls "no notion of a proc at all"), so an
	// alternative that never entered that real comparison - Neltharion's
	// Tear against mage-fire's Naxxramas trinket1, Kindling Stave against
	// warrior-arms' Blight, both true dogfood finds this lane's report
	// names - is not a fair, tested claim of "beats the pick" merely
	// because its raw stat score() is higher; score() structurally
	// undervalues exactly what made the real pick win. Tenet 8:
	// unverified data is never shown as fact, so buildAlternatives caps
	// every UNVERIFIED row's own DPSDelta at 0 (a candidate can publish
	// "at most a tie," never a positive, untested "beats the pick" claim)
	// - only Verified's own row, below, is allowed a positive number,
	// because it is the one candidate this command actually simmed
	// against the pick. A negative, unverified DPSDelta is left alone:
	// "this scores worse" is still useful fallback-ranking information,
	// and score() undervaluing the PICK's own real strength never makes
	// a genuinely-worse alternative look better than it is.
	DPSDelta float64 `json:"dps_delta"`
	// Verified is true for the one alternative (at most) verify.go's
	// own swap pass actually simmed against the pick (pk.RunnerUp at
	// buildAlternatives' own call site) - its DPSDelta is that sim's
	// REAL measured delta (swapMeasuredDelta), not the score-based
	// conversion above, and takes precedence: a runner-up that scored
	// higher than the pick but LOST the real sim (owner review, tenet
	// 8: warrior-arms horde band 20's Hammerbone scored 3.14 points
	// above the published Forsaken Greataxe pick, but measured 29.9 vs
	// 32.2 DPS and lost) must never publish a DPSDelta that makes it
	// look like the better fallback. Omitted (false) for every other
	// row - a score estimate this command never simmed at all.
	Verified bool `json:"verified,omitempty"`
	// SimDPS is this row's own real, engine-measured full-set DPS -
	// set only on the one Verified row above (this lane's brief, item
	// 7: "an alternative that was actually simmed" carries sim_dps,
	// same as a sim-decided pick does). swapAlternativeMeasuredDPS
	// (below) is the one place this is computed: sw.SwapDPS when this
	// row is still just the tested runner-up (!sw.Beat), or
	// sw.BaselineDPS when this row is the demoted FORMER pick
	// (sw.Beat promoted the other item into pk.Item, so THIS row's own
	// measured value is the number that used to be the baseline).
	// Omitted (0) for every unverified row.
	SimDPS float64 `json:"sim_dps,omitempty"`
	// LabelSuffix is slotRow.LabelSuffix's own doc, applied to this
	// alternative instead of the pick: set only when THIS alternative's
	// ItemName collides with a different item id elsewhere in the same
	// row (the pick itself, or another alternative) -
	// applyLabelSuffixForNameCollisions sets it on every colliding
	// entry, never only one side of the collision.
	LabelSuffix string `json:"label_suffix,omitempty"`
}

// buildAlternatives is slotRow.Alternatives' own builder: pk.Ties
// first (dps_delta 0, this lane's brief addendum), then list (bySlot's
// own score-sorted pool for this slot - candidatesBySlot's contract,
// unfiltered by pick()'s own per-slot narrowing, so a two-handed
// main_hand pick's alternatives still offer one-handers mixed with
// two-handers the way the brief asks for), until alternativesLimit
// total, excluding:
//   - the pick's own item (already the row's own ItemID/ItemName)
//   - the pick's pair-mate's item, by id or name (pairSlot, verify.go's
//     own doc: the same physical ring/trinket/weapon cannot also be
//     offered as a fallback for THIS slot when its pair-mate already
//     wears it)
//   - anything already listed as a tie, so a candidate never appears
//     twice
//
// list is bySlot[slot] as passed to buildReport - NOT the narrower,
// per-slot list pick() computes internally for a dual-wielder's
// off_hand (main_hand's one-handers merged in) or a dual-wielder's
// main_hand (two-handers excluded): reusing the base scoring pool
// (candidatesBySlot's own output) rather than re-deriving pick()'s
// internal, stateful narrowing is this function's own scope call - see
// this lane's report for why. off_hand of a two-handed main_hand pick
// is never reached here at all: buildReport only calls this for a
// slot whose pk.Item is non-nil, and enforceTwoHandOffHandInvariant
// (pick.go) guarantees off_hand's own Item is nil whenever main_hand
// is two-handed.
//
// referenceDPSPerPoint converts every score-based delta into a
// real DPSDelta (owner review, tenet 8). sw is verify.go's own
// swapResult for this slot, or nil when this slot had no runner-up to
// sim at all - when present, the ONE row matching pk.RunnerUp's own
// item id gets its DPSDelta and Verified overridden from that real sim
// rather than the score conversion (same owner review): a runner-up
// that outscored the pick but lost the real sim must never publish a
// DPSDelta that makes it look better than the item that beat it.
//
// The whole list is re-sorted by DPSDelta, descending, once every
// adjustment above is applied - not "ties, then score order": in the
// overwhelmingly common case those agree (bySlot's own list is
// best-score-first, the pick IS its own top entry, and a tie sits
// right behind it at the identical score, so nothing else can outscore
// either), but a tie is not guaranteed to be the effective maximum any
// more than list's own top-scoring non-pick entry is - either can be
// ahead of the other depending on the numbers, and a swap-corrected
// row in particular is the entire point of sorting at all: it was very
// likely appended near the top of the list (a demoted item usually
// scored ABOVE the pick, which is exactly why applySwaps had a runner-
// up worth testing in the first place), and its corrected DPSDelta is
// very likely now negative, so it must fall behind everything that
// still outranks it - including a tie at 0.

func buildAlternatives(pk slotPick, slot string, list []scored, picks map[string]slotPick, referenceDPSPerPoint float64, sw *swapResult, setDPS float64) []alternativeRow {
	if pk.Item == nil {
		return nil
	}
	var mateID int
	var mateName string
	if mate, ok := pairSlot[slot]; ok {
		if mp := picks[mate].Item; mp != nil {
			mateID, mateName = mp.ID, mp.Name
		}
	}
	// isPairMateItem is the mate-exclusion half of excluded (below) on
	// its own: buildAlternatives' own force-include step for a real
	// swap-tested runner-up (further down) needs to ask this WITHOUT
	// also asking "is this id already in out" the way excluded does -
	// that second question is true for the very row this step exists to
	// update, which would wrongly skip it.
	isPairMateItem := func(c scored) bool {
		return mateID != 0 && (c.ID == mateID || (mateName != "" && c.Name == mateName))
	}
	// isPickNameItem is this lane's brief, item 2 (bis-ranker-integrity-12):
	// a slot's own pick can have several real item ids sharing one
	// NAME (Sergeant Major's Cape: ids 16315/16336/16337, PvP rank 9
	// Alliance at req level 25/40/55 - three physically different
	// rewards a player earns at different points, but the same cape
	// from the reader's own point of view) - excluded's own id-based
	// check alone let a DIFFERENT id of the identical name reach
	// out as an "alternative" to itself (ten such pick-vs-alternative
	// pairs across the hybrid specs, this lane's own dogfood: the page
	// showed a pick and an alternative both reading "Sergeant Major's
	// Cape - PvP rank 9 - Sergeant Major - Alliance" with nothing
	// distinguishing them). A name match is excluded the same way an
	// id match already is - ranker-11's own dedupeAlternatives already
	// keeps alternatives unique AMONG THEMSELVES by name; this closes
	// the one case that missed, the pick's own name never being
	// checked against at all.
	isPickNameItem := func(c scored) bool {
		return c.Name == pk.Item.Name
	}
	// excluded no longer tracks "already added" - controller direction,
	// 2026-09-30 (Grand Marshal's Stave repro, buildAlternatives' own
	// doc below): the identical item id can reach this function twice,
	// through two different itemSource rows, and a plain "skip whatever
	// arrives second" guard makes the survivor depend on iteration
	// order (Ties before list, and whichever order a slot's own two
	// rows happen to sit in either one) rather than on
	// betterAlternative's own, order-independent rule. Only the pick's
	// own id and its pair-mate's are excluded outright here; a genuine
	// same-id duplicate is reconciled by upsert (below), not by being
	// silently dropped before it ever gets compared.
	excluded := func(c scored) bool {
		if c.ID == pk.Item.ID {
			return true
		}
		if isPickNameItem(c) {
			return true
		}
		return isPairMateItem(c)
	}
	// indexByID is upsert's own record of where each item id already
	// landed in out, so a same-id duplicate replaces that row (via
	// betterAlternative) instead of appending a second one.
	indexByID := make(map[int]int, alternativesLimit)
	toRow := func(c scored) alternativeRow {
		scoreDelta := c.Score - pk.Item.Score
		// DPSDelta's own doc, above: this row is not (yet) the one
		// candidate verify.go actually simmed against the pick (the
		// swap-override pass below is what promotes exactly one row to
		// that), so a positive score-unit delta here is never published
		// as a real, measured "beats the pick" claim - capped at 0. A
		// negative delta is left alone; only the positive/untested
		// direction is the tenet-8 problem this cap exists for.
		dpsDelta := scoreDelta * referenceDPSPerPoint
		if dpsDelta > 0 {
			dpsDelta = 0
		}
		return alternativeRow{
			ItemID:     c.ID,
			ItemName:   c.Name,
			SourceKind: c.Source.Kind,
			Source:     c.Source.Label,
			DPSDelta:   dpsDelta,
		}
	}
	// upsert appends c's own row, or - when this exact item id already
	// has one (indexByID) - keeps whichever of the two betterAlternative
	// prefers (report.go's own doc on that function). This is what
	// makes the survivor's SourceKind deterministic regardless of which
	// of an item's two source rows this function happens to see first.
	upsert := func(out []alternativeRow, c scored) []alternativeRow {
		row := toRow(c)
		if i, ok := indexByID[c.ID]; ok {
			if betterAlternative(row, out[i]) {
				out[i] = row
			}
			return out
		}
		indexByID[c.ID] = len(out)
		return append(out, row)
	}

	// This lane's brief, item 2: hunter-beast-mastery/marksmanship band
	// 40/50 main_hand published caster-stat weapons (spell power/
	// intellect staves) as alternatives next to a ranged spec's real
	// pick - those staves score() exactly 0 for a hunter (none of
	// score()'s weighted stats appear on them) and were never run
	// through any tournament here, so nothing about them is a real
	// alternative a player could act on. An alternative must carry at
	// least one of the spec's positively weighted stats (Score != 0)
	// or carry a real sim-measured delta from a tournament (the
	// swap-override force-include below, which runs regardless of
	// Score - that candidate WAS simmed, so it earns its place on real
	// evidence instead).
	realAlternative := func(c scored) bool { return c.Score != 0 }

	out := make([]alternativeRow, 0, alternativesLimit)
	for _, tie := range pk.Ties {
		if excluded(tie) || !realAlternative(tie) {
			continue
		}
		// The cap only ever blocks a genuinely NEW item id from
		// joining out - a duplicate of one already there still has to
		// reach upsert so betterAlternative gets to compare the two
		// (this lane's own doc above), whatever the cap already holds.
		if _, already := indexByID[tie.ID]; !already && len(out) >= alternativesLimit {
			continue
		}
		out = upsert(out, tie)
	}
	for _, c := range list {
		if excluded(c) || !realAlternative(c) {
			continue
		}
		if _, already := indexByID[c.ID]; !already && len(out) >= alternativesLimit {
			continue
		}
		out = upsert(out, c)
	}

	if sw != nil && pk.RunnerUp != nil && !isPairMateItem(*pk.RunnerUp) && !isPickNameItem(*pk.RunnerUp) {
		// bis-ranker-integrity-3, 2026-09-29, this lane's brief item 3:
		// pk.RunnerUp is the ONE candidate verify.go's own swap pass
		// actually simmed against the pick (swapBySlot, buildReport's own
		// call site) - but the tie/list loop above fills out purely by
		// score(), so a runner-up that lost badly enough on raw score()
		// to sit outside the top alternativesLimit candidates (shaman-
		// enhancement's own main_hand at bands 20/30/40: Diamond Hammer,
		// the real one-hand runner-up a dual-wield spec's own
		// twoHandBeatsPair pre-filter left as the only legal candidate to
		// test, ranked behind three untested two-handers on raw weapon
		// score alone) was silently dropped from the row entirely, even
		// though it is the ONE alternative this slot has real, measured
		// evidence about. The result read as "Verified: true" with
		// nothing anywhere on the row - no SimDPS on the pick (never
		// promoted, so never sim-decided by buildReport's own rule) and
		// no verified alternative either - which is indistinguishable
		// from "nothing was ever run" even though a real sim was. Force
		// room for it here if the tie/list loop did not already include
		// it, so "Verified: true" always has real evidence attached
		// somewhere on the row.
		found := false
		for _, a := range out {
			if a.ItemID == pk.RunnerUp.ID {
				found = true
				break
			}
		}
		if !found {
			if len(out) >= alternativesLimit {
				out = out[:alternativesLimit-1]
			}
			out = append(out, alternativeRow{
				ItemID:     pk.RunnerUp.ID,
				ItemName:   pk.RunnerUp.Name,
				SourceKind: pk.RunnerUp.Source.Kind,
				Source:     pk.RunnerUp.Source.Label,
			})
		}
		// data-followups-10 lane, 2026-09-30, item 7: only publish the
		// real sim measurement as confirmed evidence when it clears the
		// two runs' own combined standard error (sw.Significant,
		// verify.go's own doc) - a live repro (hunter-marksmanship band
		// 20, Serpent Gloves/Gloves of the Fang: both +6 Agility, the
		// other +4 Strength vs +7 Spell Power, NEITHER a stat this
		// spec's own weight_stats measures at all) kept publishing a
		// small, same-fixed-verifySeed "real" negative DPSDelta for the
		// runner-up for three regen sweeps running, purely from this
		// one trial's own sampling noise, with nothing checking whether
		// that gap was bigger than the trial's own error bar. An
		// insignificant swap leaves this row exactly as the tie/list
		// loop above already built it (score()'s own, already-capped-
		// at-0 estimate, Verified false) rather than overwriting a
		// noise-sized number with false confidence.
		if sw.Significant() {
			measured := swapMeasuredDelta(*sw)
			for i := range out {
				if out[i].ItemID == pk.RunnerUp.ID {
					out[i].DPSDelta = measured
					out[i].Verified = true
					// This lane's brief, item 7: the one row a real sim
					// actually measured publishes that measurement (SimDPS),
					// not score()'s stat estimate - the two units are not
					// comparable (score() is gone from this row entirely as
					// of bis-ranker-integrity-3, see alternativeRow's own
					// doc).
					//
					// bis-ranker-integrity-7 lane, item 2: this absolute
					// number is exactly as vulnerable to the shared-baseline
					// staleness this file's own finishedSetEpsilon guard
					// exists for (slotRow.SimDPS's own doc) - sw.BaselineDPS/
					// sw.SwapDPS describe THIS one single-slot trial, which
					// stops describing the finished set the moment some
					// OTHER slot also promoted this band. Withheld under the
					// same bar; DPSDelta above (a relative measurement,
					// always true of that one trial regardless of anything
					// else) still publishes.
					if simDPS := swapAlternativeMeasuredDPS(*sw); math.Abs(simDPS-setDPS) <= finishedSetEpsilon {
						out[i].SimDPS = simDPS
					}
					break
				}
			}
		}
	}

	// Controller direction, 2026-09-30 (priest-shadow band 60 Alliance
	// main_hand's own repro): Grand Marshal's Stave published twice,
	// once via loot.json's own "pvp:rank-18:alliance" source and once
	// via the quartermaster's own "vendor" row that
	// vendorInheritsPvpRankGate (data.go) copies that same rank onto -
	// two itemSource rows for the one physical item id, which can each
	// reach this function's own inputs (Ties/list/the swap force-
	// include above all key on Score, not on "have I already named
	// this id under a DIFFERENT source"). Alternatives are unique by
	// item id; dedupeAlternatives keeps the one row worth showing a
	// player for each id and drops the rest before the final sort
	// below ever sees them.
	out = dedupeAlternatives(out)

	sort.SliceStable(out, func(i, j int) bool {
		if out[i].DPSDelta != out[j].DPSDelta {
			return out[i].DPSDelta > out[j].DPSDelta
		}
		return out[i].ItemID < out[j].ItemID
	})
	return out
}

// slotItemLevelByID is item id -> ItemLevel for every candidate this
// slot's own pool (list, the same []scored buildAlternatives just
// read) or pick/runner-up carries - applyLabelSuffixForNameCollisions'
// own resolver. pk's Item/RunnerUp are added on top of list because a
// swap-promoted pick is not always still present in list by the time
// this runs (pick.go's own doc on the picks map replacing entries
// outright).
func slotItemLevelByID(list []scored, pk slotPick) func(id int) int {
	byID := make(map[int]int, len(list)+2)
	for _, c := range list {
		byID[c.ID] = c.ItemLevel
	}
	if pk.Item != nil {
		byID[pk.Item.ID] = pk.Item.ItemLevel
	}
	if pk.RunnerUp != nil {
		byID[pk.RunnerUp.ID] = pk.RunnerUp.ItemLevel
	}
	return func(id int) int { return byID[id] }
}

// slotStatsByID is item id -> Stats for the same pool slotItemLevelByID
// reads, and for the same reason: applyLabelSuffixForNameCollisions'
// own resolver for roleStatHint, once an item-level suffix can no
// longer tell two same-named items apart (bis-ranker-integrity-17
// lane's brief, item 4: three "Signet Ring of the Bronze Dragonflight"
// rings, all item_level 80, with a tank/melee/caster stat line each).
func slotStatsByID(list []scored, pk slotPick) func(id int) map[string]float64 {
	byID := make(map[int]map[string]float64, len(list)+2)
	for _, c := range list {
		byID[c.ID] = c.Stats
	}
	if pk.Item != nil {
		byID[pk.Item.ID] = pk.Item.Stats
	}
	if pk.RunnerUp != nil {
		byID[pk.RunnerUp.ID] = pk.RunnerUp.Stats
	}
	return func(id int) map[string]float64 { return byID[id] }
}

// roleStatLabels mirrors web/src/lib/sim/copy.ts's own `statLabel` map
// (Go and TypeScript do not import each other, the same reason
// LEVEL_BANDS is a copy in pipeline.addonrotation) for exactly the
// stats roleStatHint below ever names -- so a label this emits reads
// identically to the same stat's name everywhere else on the site.
var roleStatLabels = map[string]string{
	"defense":      "Defense",
	"block":        "Block",
	"parry":        "Parry",
	"dodge":        "Dodge",
	"spell_power":  "Spell power",
	"healing":      "Healing",
	"attack_power": "Attack power",
	"agility":      "Agility",
	"strength":     "Strength",
	"intellect":    "Intellect",
	"spirit":       "Spirit",
}

// roleStatOrder is the priority roleStatHint checks stats in, most
// role-defining first. A tank-signaling stat (defense/block/parry/
// dodge) outranks a caster one (spell_power/healing), which outranks a
// melee one (attack_power/agility), which outranks a plain primary stat
// (strength/intellect/spirit) -- every one of those beats the generic
// stats every piece of gear tends to carry regardless of role
// (stamina, hit, crit, mp5, ...), which is exactly why those generic
// stats are left out of roleStatLabels entirely: a ring with only
// stamina and hit never produces a hint, rather than a misleading one.
var roleStatOrder = []string{
	"defense", "block", "parry", "dodge",
	"spell_power", "healing",
	"attack_power", "agility",
	"strength", "intellect", "spirit",
}

// roleStatHint is "(Defense)"/"(Agility)"/"(Spell power)" for the
// first stat in roleStatOrder this item's own Stats names a nonzero
// value for, or "" when none of them do (an item whose only stats are
// generic ones a hint could not usefully distinguish by). Sweep 16's
// own repro (the three Signet Ring of the Bronze Dragonflight variants,
// applyLabelSuffixForNameCollisions' own doc) is the case this exists
// for: {stamina:24, strength:13, defense:7} -> "(Defense)" (a tank
// ring), {agility:24, stamina:13, hit:10} -> "(Agility)" (a melee
// ring), {intellect:9, stamina:8, spell_power:28, mp5:5} -> "(Spell
// power)" (a caster ring) -- three item-level-80 rings an ilvl suffix
// alone cannot tell apart, each correctly labelled by its own role.
func roleStatHint(stats map[string]float64) string {
	for _, key := range roleStatOrder {
		if stats[key] > 0 {
			return fmt.Sprintf("(%s)", roleStatLabels[key])
		}
	}
	return ""
}

// applyLabelSuffixForNameCollisions sets LabelSuffix on row itself and
// on any row.Alternatives entry whose ItemName equals row.ItemName but
// whose ItemID differs - two different items sharing one display name
// (warlock-destruction band 60 Alliance's own legs pick, this lane's
// brief item 6: "Sentinel's Silk Leggings" id 237815, ilvl 78,
// Forever's own reissue, sold beside a real vanilla item of the same
// name, id 22752, ilvl 65, by the same vendor), distinguished only by
// a bare id neither this page nor the in-game tooltip surfaces on its
// own. itemLevel resolves an id to its own item level; an id it has no
// answer for (0) publishes no suffix rather than a misleading
// "(ilvl 0)". Never set for the overwhelming majority of rows, whose
// name is unique within their own Alternatives list -
// dedupeAlternatives (buildAlternatives, above) already collapses a
// same-name collision BETWEEN two alternatives down to one row before
// this ever runs, so the only collision left to catch is the pick's
// own name against a surviving alternative.
//
// bis-ranker-integrity-17 lane's brief, item 4: an item-level suffix
// cannot tell apart a collision where the item levels ALSO tie -
// player-review sweep 15/16's own repro, "Signet Ring of the Bronze
// Dragonflight" (items 21200/21205/21210), three different rings, all
// item_level 80, with a tank/melee/caster stat line each. For any
// colliding alternative whose own item level equals row's (rowTies
// below), that ONE alternative gets a roleStatHint suffix instead of
// the useless "(ilvl 80)" both sides would otherwise share; row itself
// switches to its own roleStatHint the moment ANY alternative ties its
// level, since row.LabelSuffix is one field covering every collision
// it is in. A stat map with no role-defining stat at all (roleStatHint
// returns "") falls back to the ilvl suffix anyway: still true, just
// no longer the whole story, same as before this existed.
func applyLabelSuffixForNameCollisions(row *slotRow, itemLevel func(id int) int, statsOf func(id int) map[string]float64) {
	if row.ItemID == 0 || row.ItemName == "" {
		return
	}
	collides := false
	rowTiesOnLevel := false
	rowLevel := itemLevel(row.ItemID)
	for i := range row.Alternatives {
		if row.Alternatives[i].ItemName != row.ItemName || row.Alternatives[i].ItemID == row.ItemID {
			continue
		}
		collides = true
		altID := row.Alternatives[i].ItemID
		altLevel := itemLevel(altID)
		if altLevel > 0 && altLevel == rowLevel {
			rowTiesOnLevel = true
			if hint := roleStatHint(statsOf(altID)); hint != "" {
				row.Alternatives[i].LabelSuffix = hint
				continue
			}
		}
		if altLevel > 0 {
			row.Alternatives[i].LabelSuffix = fmt.Sprintf("(ilvl %d)", altLevel)
		}
	}
	if !collides {
		return
	}
	if rowTiesOnLevel {
		if hint := roleStatHint(statsOf(row.ItemID)); hint != "" {
			row.LabelSuffix = hint
			return
		}
	}
	if rowLevel > 0 {
		row.LabelSuffix = fmt.Sprintf("(ilvl %d)", rowLevel)
	}
}

// dedupeAlternatives keeps exactly one alternativeRow per ItemID, then
// exactly one per ItemName (bis-ranker-integrity-11's brief, item 3: "a
// player reads names") - buildAlternatives' own doc above has the
// concrete repro the ID pass closes (Grand Marshal's Stave, reached
// through both a "pvp" and a "vendor" itemSource row for the identical
// id); the NAME pass closes a different repro this lane dogfooded
// directly: Signet Ring of the Bronze Dragonflight publishing TWICE in
// the same band-60 finger1 Alternatives list under two different item
// ids (a player reading the list sees the same ring named back to back
// with no way to tell they are even different rows), the same shape
// Sentinel's/Scout's Medallion (band 30/40 neck), Highlander's/
// Defiler's Leather Girdle (band 50 waist) and Sentinel's Chain
// Leggings (band 60 legs) all share. Both passes keep the row
// betterAlternative prefers; ties resolve to whichever row this
// function saw first, so the result is deterministic regardless of
// map/slice iteration order upstream.
func dedupeAlternatives(rows []alternativeRow) []alternativeRow {
	byID := dedupeAlternativesBy(rows, func(r alternativeRow) any { return r.ItemID })
	return dedupeAlternativesBy(byID, func(r alternativeRow) any { return r.ItemName })
}

// dedupeAlternativesBy is dedupeAlternatives' own shared pass: keep
// exactly one row per key(row), the one betterAlternative prefers.
func dedupeAlternativesBy(rows []alternativeRow, key func(alternativeRow) any) []alternativeRow {
	indexByKey := make(map[any]int, len(rows))
	out := make([]alternativeRow, 0, len(rows))
	for _, r := range rows {
		k := key(r)
		if i, ok := indexByKey[k]; ok {
			if betterAlternative(r, out[i]) {
				out[i] = r
			}
			continue
		}
		indexByKey[k] = len(out)
		out = append(out, r)
	}
	return out
}

// betterAlternative reports whether candidate should replace incumbent
// as the one published row for the item id they share - controller
// direction, 2026-09-30: the row with the higher DPSDelta wins; on a
// tie, the "pvp" SourceKind wins over any other (a player recognises
// "Alliance PvP Rank 18" as the real requirement more readily than a
// vendor row that only carries a rank at all because
// vendorInheritsPvpRankGate, data.go, copied it there - the opposite
// question from band.go's own sourceKindPriority, which decides a
// SLOT'S OWN pick and prefers "vendor" for exactly the reason a plain
// gold purchase is simpler than a PvP grind; an alternative is not the
// pick, and here the more specific, more informative source wins
// instead).
func betterAlternative(candidate, incumbent alternativeRow) bool {
	if candidate.DPSDelta != incumbent.DPSDelta {
		return candidate.DPSDelta > incumbent.DPSDelta
	}
	return candidate.SourceKind == "pvp" && incumbent.SourceKind != "pvp"
}

// swapDeltaNoiseFloorDPS is how small a real, measured DPS delta must
// be (in absolute DPS, not swapMargin's relative 1%) before
// swapMeasuredDelta calls it an honest, exact tie (0) rather than
// publishing the signed number.
//
// Controller direction, bis-ranker-integrity-8, 2026-09-30 (the eighth
// wow-player sweep, day3/player-review-21/casters.md finding 1):
// several caster main_hand slots published an exact "dps_delta: 0" for
// a zero-caster-stat weapon (a rogue dagger, a stat-less epic sword,
// Teebu's Blazing Longsword) that verify.go's own swap sim (sw) had
// actually measured a real, small, POSITIVE delta over the picked
// spell-power staff - it just had not cleared swapMargin's own 1%
// PROMOTION bar, so no swap happened (sw.Beat == false) and the OLD
// version of this function (see its history below) clamped EVERY
// positive delta in that branch to exactly 0 regardless of size,
// which the web renders as "same DPS" even when the real measured gap
// was several DPS on a spec whose baseline DPS makes 1% a large
// absolute number. swapMargin still decides whether a runner-up is
// good enough to PROMOTE (verify.go's own job, unchanged); this floor
// decides only whether the ONE candidate this command actually simmed
// against the pick measured close enough to call an honest, exact tie
// in the row it publishes - a different question the old clamp
// wrongly answered by reusing swapMargin's relative bar for it.
const swapDeltaNoiseFloorDPS = 0.05

// swapMeasuredDelta is the real, sim-measured DPS delta between
// whichever item verify.go's own swap sim (sw) tried against the pick
// and the pick's own measured DPS - owner review, tenet 8. sw.SwapDPS
// is always the runner-up-at-verify-time's own measured DPS and
// sw.BaselineDPS is always the then-current pick's; applySwaps
// (verify.go) swaps which physical item plays "pick" versus
// "runner-up" once sw.Beat, but never re-labels these two numbers, so
// the item that is NOT the currently-published pick owns BaselineDPS
// when sw.Beat promoted the other one into the pick, and owns SwapDPS
// otherwise. Either way this starts as "the other item's measured DPS
// minus the pick's own measured DPS", negative whenever the pick's own
// real DPS is the higher of the two - always true when sw.Beat
// (Beat means the promoted item's SwapDPS beat the demoted item's
// BaselineDPS by more than swapMargin, so |delta| here is always at
// least swapMargin's own 1% of a real spec's DPS, far past
// swapDeltaNoiseFloorDPS), but NOT always true when !sw.Beat:
// swapMargin (verify.go) deliberately keeps the scored pick on a
// runner-up that measured HIGHER but not by enough to clear the noise
// margin (verify.go's own Serpent's Shoulders/Mantle of Honor example,
// 78.7 vs 78.6) - a genuinely small, real measurement, not nothing.
//
// bis-ranker-integrity-2 (2026-09-29) originally capped every positive
// delta in that branch at exactly 0, reasoning that "a runner-up that
// did not clear the promotion bar is reported as at best a tie" - but
// swapMargin is a RELATIVE bar (1% of baseline) built to decide
// promotion, not an absolute noise floor on the row's own published
// number; reusing it to zero out a real measurement made every one of
// these rows claim an exact tie regardless of how large the actual,
// measured gap was (bis-ranker-integrity-8's caster finding above).
// This function calls an exact tie (0) when the real measured delta
// itself is smaller than swapDeltaNoiseFloorDPS - an absolute floor on
// THIS row's own number, answering "is this measurement itself too
// small to trust" rather than "did it clear an unrelated promotion
// bar". A negative delta (a genuine, measured loss, however small) is
// never clamped by either rule below.
//
// Controller direction, 2026-09-30 (ninth wow-player sweep,
// day3/player-review-24/casters.md finding 1, lane bis-ranker-
// integrity-9): the noise floor alone still let a real, well-above-
// noise POSITIVE delta through the !sw.Beat branch - mage-arcane band
// 60's own Weakness Analyzer published dps_delta +2.34 next to the
// KEPT pick Talisman of Ascendance, warlock-destruction's own Orb of
// the Darkmoon +1.6 over the neck pick, both verified: true. !sw.Beat
// means beatsByMargin (verify.go) already found this candidate did NOT
// clear swapMargin, so the pick was kept on purpose - a positive
// number here reads as "the site chose the worse item", which is never
// true once the pick was kept. Every positive delta in the !sw.Beat
// branch is now an honest tie (0), whatever its size; only a genuine
// loss (delta <= 0) can ever publish a nonzero number here, and the
// noise floor still rounds a tiny loss to an exact tie the same way it
// always did.
func swapMeasuredDelta(sw swapResult) float64 {
	if sw.Beat {
		return sw.BaselineDPS - sw.SwapDPS
	}
	delta := sw.SwapDPS - sw.BaselineDPS
	if delta > 0 {
		return 0
	}
	if math.Abs(delta) < swapDeltaNoiseFloorDPS {
		return 0
	}
	return delta
}

// swapAlternativeMeasuredDPS is the ALTERNATIVE row's own measured
// full-set DPS from the same swap sim swapMeasuredDelta reads - this
// lane's brief, item 7. The alternative row is always "the OTHER
// item", never the currently-published pick (swapMeasuredDelta's own
// doc explains why): when sw.Beat, that other item is the DEMOTED
// former pick, whose own measured value is sw.BaselineDPS (it was the
// baseline the swap was measured against); when !sw.Beat, that other
// item is still just the tested runner-up, whose own measured value
// is sw.SwapDPS.
func swapAlternativeMeasuredDPS(sw swapResult) float64 {
	if sw.Beat {
		return sw.BaselineDPS
	}
	return sw.SwapDPS
}

// alternativeCarriesRealEvidence reports whether alts already names the
// one candidate a real sim actually measured against the pick (a
// Verified row - buildAlternatives' own contract, at most one) - this
// lane's brief, item 3: buildReport's own fallback SwapNote (above)
// only needs to state the real numbers itself when Alternatives does
// not already show them.
func alternativeCarriesRealEvidence(alts []alternativeRow) bool {
	for _, a := range alts {
		if a.Verified {
			return true
		}
	}
	return false
}

// bandReport is one band's whole answer for one faction: the pick per
// slot, the verification DPS, the diff against the previous band, and
// the counts an honest reader needs (how many eligible items had no
// known source).
type bandReport struct {
	Spec string `json:"spec"`
	Band int    `json:"band"`
	// Preset is the sim context the entry was measured under: "bare"
	// (the character and its class kit only) or "raid" (the curated Phase
	// 1 raid preset on top, level 60 only).
	Preset       string      `json:"preset"`
	Faction      string      `json:"faction"`
	Race         string      `json:"race"`
	Talents      string      `json:"talents"`
	TalentPoints int         `json:"talent_points"`
	Weights      []weightRow `json:"weights"`
	Slots        []slotRow   `json:"slots"`
	SetDPS       float64     `json:"set_dps"`
	// SetDPSPartial is true when at least one published slot's own pick
	// carries an item id this build's embedded item database does not
	// have (a Slots row with SimStatus == "not_in_sim", above) - this
	// lane's brief, item 1. simdb.Attach's own UnequipUnknown silently
	// empties that one item's equipment slot before EVERY sim this
	// band's own SetDPS was ever measured from (the baseline run,
	// every swap trial, applySwaps' own final re-measure), so SetDPS
	// itself was always a real, honestly-computed number for "this
	// set, minus that one item" - never for the full set this band
	// actually publishes. The two honest choices this lane's brief
	// offers are publishing only the simmable part (with this flag) or
	// omitting set_dps outright; omitting it would throw away a true
	// measurement of every OTHER slot's own real contribution just
	// because one relic/idol/item this build's client export never
	// carried a row for happened to also be the best pick somewhere in
	// the set, so this flag is the option chosen - set_dps stays
	// published, exactly as the engine already computed it, with this
	// flag naming what it is not a measurement of.
	SetDPSPartial     bool     `json:"set_dps_partial,omitempty"`
	NoSourceCount     int      `json:"no_source_count"`
	NoSourceSample    []string `json:"no_source_sample,omitempty"`
	NewAtBand         []string `json:"new_at_band"`
	WeightsRunSeconds float64  `json:"weights_run_seconds"`
	VerifyRunSeconds  float64  `json:"verify_run_seconds"`
	VerifyErrors      []string `json:"verify_errors,omitempty"`
	// Coverage is band.go's buildBandPool own per-slot count (lane
	// rank-guardrails' guardrail A): planner slot -> how many items
	// eligible() passed for this band+faction, and how many of those
	// sourceFor() could actually find a source for. A slot missing from
	// this map had zero eligible candidates at all (not even an
	// unsourced one) - see coverageRow's own doc for why a cross-class
	// set item still counts here even though it never reaches Slots or
	// NoSource. web/src/lib/bis/types.ts's BisBand.coverage is this
	// field's read contract.
	Coverage map[string]coverageRow `json:"coverage"`
	// ReferenceDPSPerPoint is the measured, un-normalised DPS this
	// band's weights run found for one point of the spec's own
	// reference stat (data/curated/specs.json's reference_stat) - this
	// lane's brief, item 2: the raw number every OTHER published
	// Weights[i].Weight is a ratio against (Weight_i/scale;
	// simrun.go's runWeights/referenceStatRawWeight read scale itself
	// straight off the engine's own result, since the normalised
	// weights map alone can never recover it - the reference stat's
	// own entry is exactly 1.0 by construction). Lets the page turn
	// "Strength 1.99" into "1 Attack Power = 0.07 DPS, Strength 1.99 (=
	// 0.14 DPS per point)" instead of publishing a bare, unitless
	// ratio with nothing saying what "1" means.
	//
	// A pointer, not a bare float64: nil (omitted from the JSON
	// entirely) when WeightsReason below is set, because a band whose
	// reference measurement is not trustworthy was never actually
	// divided into any of this band's Weights - publishing 0 here
	// would read as "this stat is worth zero DPS", a specific false
	// claim, not "unknown" (this lane's brief, item 2).
	ReferenceDPSPerPoint *float64 `json:"reference_dps_per_point,omitempty"`
	// ScoreUnit documents slotRow.Score's own unit for every consumer
	// of this JSON (alternativeRow carries no score at all as of
	// bis-ranker-integrity-3 - see its own doc) - this lane's brief,
	// item 7: "publish score ONLY where it is the stat-weight estimate
	// in reference-stat points, and document that in the JSON". Always
	// scoreUnitReferenceStatPoints today (score.go has exactly one
	// scoring function); a constant field rather than a bare doc
	// comment because a future second scoring unit must not silently
	// leave old JSON ambiguous about which one an already-written file
	// used.
	ScoreUnit string `json:"score_unit"`
	// WeightsReason is set only when this band's whole weights sweep
	// could not be trusted (referenceMeasurementReason, weights.go) -
	// this lane's brief, item 2: a caster reading "not significant" on
	// every row with no explanation would not know this band differs
	// from ordinary sweep noise. Empty on every ordinary band (the
	// overwhelming majority).
	WeightsReason string `json:"weights_reason,omitempty"`
	// ScaleReferenceStat is normalizeScaleFactors' own return
	// (weights.go): the weight_stats id whose own Weight every row's
	// ScaleFactor in Weights is divided by (the SimulationCraft-
	// familiar "top stat = 1.00" convention this lane's brief asks
	// for) - "" when no row qualified (see normalizeScaleFactors' own
	// doc for exactly when that is). Never a haste stat
	// (isHasteStat) - see that function's own doc for why.
	ScaleReferenceStat string `json:"scale_reference_stat,omitempty"`
	// HasteScaleFactor is hasteScaleFactorFromRows' own return
	// (weights.go) - owner correction, 2026-09-30, after player review:
	// haste is not a per-point stat (isHasteStat) and this lane's
	// weight rail no longer shows it as a table row at all, only as a
	// one-line caption ("Haste: 1.58 per 1%..."), so the site needs
	// this one number without having to find and re-read the haste
	// entry out of Weights itself. nil when this spec carries no haste
	// weight_stat, or when ScaleReferenceStat is "" (no per-point
	// anchor this band's own sweep trusted - see that field's own
	// doc).
	HasteScaleFactor *float64 `json:"haste_scale_factor,omitempty"`
	// HitToCap is how far the weights character is from the miss-table
	// caps (hitToCap's own doc), or for a caster from the spell hit cap
	// (spellHitToCap's own doc): the site's "hit to cap first" figure.
	// Omitted for a spec with no hit table to cap against.
	HitToCap *publishedHitToCap `json:"hit_to_cap,omitempty"`
	// HasteOnItems is bandHasHasteCandidate's own return (weights.go) -
	// owner correction, 2026-09-30, after the caption's own doubled-
	// suffix bug was found on screenshot review ("Haste: 1.58 per 1%,
	// per 1%"): whether any candidate this band's own eligible() pass
	// considered carries a nonzero haste stat, published so the site's
	// own caption ("Haste: <n> per 1%", plus ", not in the table
	// because no item at this band has it" only when this is false)
	// never has to infer that from a haste weightRow's own
	// Insignificant flag - which answers a different, statistical
	// question (was the sweep's own sample clean enough to trust the
	// number), not this plain inventory one. No omitempty: false is a
	// real, meaningful answer here, never "absent" - a file published
	// before this lane simply carries no `haste_on_items` key at all,
	// which the site's own load path (`normaliseBisFile`) defaults to
	// true, not false, on the theory that an unknown band should not
	// silently start SHOWING a claim ("no item has it") an older file
	// never made.
	HasteOnItems bool `json:"haste_on_items"`
	// WeightsLowConfidence is this lane's brief (ranker-weights-anchor),
	// item 3: true when this band's own primary-stat anchor row
	// (ScaleReferenceStat, normally) measured insignificant
	// (isWeightSignificant, weights.go) even after
	// primaryStatSignificanceCheck's own one-time retry at
	// primaryStatRetryIterationsFactor iterations. The anchor row is
	// still published as if significant either way (the brief's own
	// rule: "no primary stat ever published as 'not significant'" -
	// forceAnchorRowSignificant, weights.go) - this field is the
	// honest signal that the underlying measurement is still noisy,
	// without putting "not significant" back on the one row a reader
	// is told to trust as 1.00. omitempty: false (the overwhelming
	// majority of bands) is the ordinary, confident case; the site
	// does not render anything from this field yet (this lane's
	// brief: "web shows nothing new yet").
	WeightsLowConfidence bool `json:"weights_low_confidence,omitempty"`
	// Role, Profile and Metrics are set for a tank or a healer band only.
	// A tank band (annotateTankBand, score_tank_band.go) has Role "tank"
	// and Metrics a *tankMetricsReport, the figures the site headlines in
	// place of DPS; on it SetDPS is the tank's own damage and every
	// sim-decided figure (sim_dps, dps_delta) is in tank score, as
	// ScoreUnit says. A healer band (attachHealerFields, score_heal.go)
	// has Role "healer", the heal Profile it was measured under and
	// Metrics a *healingMetrics. A damage spec's entry carries none.
	Role    string `json:"role,omitempty"`
	Profile string `json:"profile,omitempty"`
	Metrics any    `json:"metrics,omitempty"`
}

// scoreUnitReferenceStatPoints is bandReport.ScoreUnit's only value
// today - see that field's own doc.
const scoreUnitReferenceStatPoints = "reference_stat_points"

type weightRow struct {
	// Stat is the weight_stats id (data/curated/specs.json) - for a
	// rating-family stat (hit/crit/dodge/parry/block/defense) this is
	// also the exact key candidate.Stats/data.go's ratingFactors uses
	// (Forever's engine unifies each rating family into one Stat; see
	// data.go's ratingStatColumns doc).
	Stat string `json:"stat"`
	// Weight is per RATING POINT for a rating-family stat (Unit ==
	// "rating"), because that is what the item's own tooltip and this
	// site show a player (this lane's brief, item 3): "Crit 1.44 RAP ·
	// 0.08 DPS per point" is true for a "+14 Crit" item, not for "+1%
	// crit chance". For every other stat, Weight is exactly what the
	// weights sweep measured, unchanged. publishWeightRatingUnits
	// (below) is the ONE place that divides a rating-family row's raw,
	// sim-unit weight down to this per-rating-point number -
	// buildReport itself still receives and computes against the raw
	// sim-unit weight throughout (isWeightSignificant's own 25%-of-
	// value bar runs on that raw number, before this conversion, so
	// dividing every row here by the same positive factor never
	// changes which rows it flags).
	Weight float64 `json:"weight"`
	// Error is this weight's standard error, in the SAME units as
	// Weight above - divided by RatingFactor alongside Weight for a
	// rating-family row, so "±" still means what Weight's own unit
	// says it means (sim/adapter.Weights' own StatWeight.Error is the
	// raw, sim-unit source; see isWeightSignificant's own doc for why
	// this command publishes an error bar at all).
	Error float64 `json:"error"`
	// Insignificant is this command's OWN, stricter call (see
	// isWeightSignificant), not sim/adapter's Insignificant field: the
	// page greys this row out and the Pawn/planner consumers of this
	// JSON should not treat it as a real number.
	Insignificant bool `json:"insignificant"`
	// Unit is "rating" for a rating-family stat - Weight/Error above
	// are per RATING POINT, not per sim unit (percent). Omitted for
	// every other stat: a plain sim-unit weight (Strength, Spell
	// Power, ...) needs no unit qualifier here (bandReport.ScoreUnit
	// is the separate, existing "what does Score mean" field; this is
	// "what does THIS ROW'S Weight mean").
	Unit string `json:"unit,omitempty"`
	// RatingFactor is this build's own gametables/combatratings.txt
	// level-60 rating points per 1% for Stat (data.go's
	// loadRatingFactors) - set only alongside Unit == "rating", so a
	// consumer can render "14 rating = 1%" in a tooltip without
	// hardcoding the client's own conversion table a second time, and
	// so the contract Weight == WeightPerPercent/RatingFactor is
	// checkable directly off this one row.
	RatingFactor float64 `json:"rating_factor,omitempty"`
	// WeightPerPercent is the weight exactly as the weights sweep
	// measured it, per SIM UNIT (one point of hit/crit/dodge/parry/
	// block/defense percentage, i.e. what score() (score.go) actually
	// multiplies a converted candidate's stat by) - published
	// alongside Weight above so nothing the sweep measured is lost
	// once Weight itself switches to per-rating-point for a
	// rating-family row. Equal to Weight for every non-rating-family
	// stat (RatingFactor unset).
	WeightPerPercent float64 `json:"weight_per_percent"`
	// ScaleFactor, DPSPerPoint and ScaleError are this lane's brief
	// (bis-weights-simc, owner: "we need to make the stat weights
	// align with simcraft stat weights output - that's what people
	// are familiar with"): normalizeScaleFactors (weights.go) fills
	// these on every row, alongside the existing Weight/Error/
	// WeightPerPercent fields above, which stay published unchanged.
	//
	// ScaleFactor is Weight (already per-point: per rating point for
	// a rating-family row, per 1% for a haste row, per plain point for
	// everything else) divided by this band's ScaleReferenceStat's own
	// Weight - the convention SimulationCraft's own scale-factor table
	// uses: the single highest-weighted PER-POINT stat prints as
	// exactly 1.00, every other stat as a fraction of it. A haste row
	// (isHasteStat) is never the anchor (vanilla haste is a flat
	// 1%-per-point stat, not comparable point-for-point against a
	// primary/rating stat - owner correction, 2026-09-30) but still
	// publishes its own ScaleFactor on the same divisor, so a reader
	// can see "haste is worth about 1.6x what top-stat is worth per
	// point" even though haste itself never sets that scale. The site's
	// own weight rail (web/src/lib/bis/panel-view.ts) does not render a
	// haste row in its table at all (second owner correction, same
	// date, after player review) - it reads this same number back out
	// of bandReport.HasteScaleFactor instead, for a one-line caption.
	ScaleFactor float64 `json:"scale_factor"`
	// DPSPerPoint is absolute DPS per point of Stat: Weight (per-point,
	// same units as ScaleFactor's own numerator) times this band's own
	// ReferenceDPSPerPoint (bandReport's own field - the measured DPS
	// for one point of the OLD, engine reference stat). Independent of
	// ScaleFactor/ScaleReferenceStat - it needs no anchor, only the raw
	// per-point weight and the band's own DPS-per-reference-point
	// measurement - so it still publishes even on a band with no
	// significant per-point stat to normalize against (ScaleFactor
	// would be 0 there; DPSPerPoint is not).
	DPSPerPoint float64 `json:"dps_per_point"`
	// ScaleError is Error on the same ScaleReferenceStat divisor as
	// ScaleFactor, so a reader comparing two rows' "±" is comparing the
	// same normalized units the rows' own ScaleFactor values are in,
	// not raw sim-unit error against a normalized value.
	ScaleError float64 `json:"scale_error"`
}

// isHasteStat is normalizeScaleFactors' own carve-out (owner
// correction, 2026-09-30): this command's two haste ids (data.go's
// convertRatingStats never touches either - haste has no
// gametables/combatratings.txt rating column in this ruleset, see
// ratingStatColumns's own doc) are always published per 1% of haste,
// never per rating point like crit/hit, so a haste row's own Weight is
// not directly comparable point-for-point against a primary stat's
// Weight - including it in the search for ScaleReferenceStat would let
// a haste row become the top-stat anchor and silently misrepresent
// every OTHER row's own scale factor (exactly what the mock's first,
// wrong draft did).
func isHasteStat(stat string) bool {
	return stat == "melee_haste" || stat == "spell_haste"
}

// significanceErrorFraction and isWeightSignificant now live in
// weights.go, alongside effectiveWeights -- the one function that
// turns a raw stat-weights result into the plain numbers score()
// ranks by, zeroing exactly the weights this file's own
// isWeightSignificant call marks Insignificant below.

// publishWeightRatingUnits returns a NEW []weightRow (this package's
// immutability rule) with every rating-family row (weightRow.Stat one
// of data.go's ratingStatColumns keys) republished per RATING point
// instead of per sim unit: Weight and Error divide by factors[Stat],
// Unit is set to "rating" and RatingFactor to the factor itself.
// WeightPerPercent is set to the ORIGINAL, un-divided Weight for
// every row (rating-family or not), so "weight ==
// weight_per_percent/rating_factor" holds exactly wherever
// rating_factor is published (this lane's brief, item 3's contract
// test) and nothing the weights sweep measured is lost.
//
// Called once per band+faction (main.go's runSpec, right after
// buildReport returns) rather than folded into buildReport itself:
// buildReport has dozens of existing callers in report_build_test.go
// that construct a bandReport directly from a raw sim-unit
// map[string]api.StatWeight and never touch a build directory at all
// - threading ratingFactors through buildReport's own signature would
// force every one of those tests to grow an unrelated parameter for a
// concern (rating units) none of them are about. Applying the
// conversion as a small, separate, well-tested pass over the rows
// buildReport already produced keeps that surface untouched while
// still reaching both outputs this lane's brief asks for (the JSON,
// via bandReport.Weights, and the markdown, via writeMarkdown reading
// the very same reports slice).
func publishWeightRatingUnits(rows []weightRow, factors ratingFactors) []weightRow {
	out := make([]weightRow, len(rows))
	for i, row := range rows {
		row.WeightPerPercent = row.Weight
		if factor, ok := factors[row.Stat]; ok {
			row.Unit = "rating"
			row.RatingFactor = factor
			row.Weight = row.Weight / factor
			row.Error = row.Error / factor
		}
		out[i] = row
	}
	return out
}

// noSourceSampleSize bounds how many unsourced item names the JSON
// and markdown carry - the count is exact, the sample is just enough
// to spot-check without shipping a multi-thousand-row list every run.
const noSourceSampleSize = 15

// buildReport assembles one band+faction's report from pick() output,
// the weights this band used, verification results, and the previous
// band's picks (nil for the first band run). bySlot is candidatesBySlot's
// own output for this band+faction (main.go's runSpec builds a fresh one
// per band iteration) - buildAlternatives reads it to fill each row's
// Alternatives; a nil/empty bySlot (every test but the ones this lane's
// brief adds) simply publishes no alternatives, same as an empty map
// would. referenceDPSPerPoint is runWeights' own second return
// (simrun.go) - this lane's brief, item 2.
// weightsReason is referenceMeasurementReason's own return
// (weights.go): "" for the ordinary band, or the reason this whole
// band's weights sweep could not be trusted - this lane's brief, item
// 2. A non-empty reason forces every weightRow to publish
// Insignificant regardless of what isWeightSignificant says about its
// own stat in isolation (a poisoned reference can make an individual
// ratio look deceptively clean - see referenceMeasurementReason's own
// doc), and ReferenceDPSPerPoint is omitted from the JSON entirely
// (nil, not a misleading 0) rather than publish the ratio every other
// weight was never actually normalised against.
func buildReport(spec specInfo, band int, faction, race, talents string, talentPoints int, weights map[string]api.StatWeight, weightOrder []string, picks map[string]slotPick, setDPS float64, swaps []swapResult, noSource []candidate, previous map[string]slotPick, weightsSeconds, verifySeconds float64, verifyErrors []string, coverage map[string]coverageRow, bySlot map[string][]scored, referenceDPSPerPoint float64, weightsReason string) bandReport {
	swapBySlot := make(map[string]swapResult, len(swaps))
	for _, s := range swaps {
		swapBySlot[s.Slot] = s
	}
	erroredSlots := make(map[string]bool, len(verifyErrors))
	for _, e := range verifyErrors {
		if slot, _, ok := strings.Cut(e, ":"); ok {
			erroredSlots[slot] = true
		}
	}

	rows := make([]slotRow, 0, len(slotOrder))
	for _, slot := range slotOrder {
		pk := picks[slot]
		row := slotRow{Slot: slot, FactionNote: pk.FactionNote}
		if pk.Item == nil {
			// This lane's brief, item 8: every empty slot carries an
			// empty_reason - before this, pick()'s own "off_hand under a
			// two-hander" case (and any slot whose own candidate list
			// was genuinely empty) published a bare {"slot": "off_hand"}
			// with no reason at all, the second "nothing here" shape the
			// page had alongside noDPSValueReason's own (hunter-
			// marksmanship band 40 alliance's own off_hand, named
			// directly in the brief).
			row.EmptyReason = emptyReasonForNilPick(slot, picks)
		}
		if pk.Item != nil {
			row.ItemID = pk.Item.ID
			row.ItemName = pk.Item.Name
			row.SetBonus = pk.SetBonus
			row.Score = pk.Item.Score
			// This lane's brief, item 7: a candidate a real sim actually
			// measured (MeasuredDPS > 0 - trinkets.go/rank.go/sets.go's
			// own tournaments, or a swap promotion below) publishes that
			// measurement instead of score()'s stat estimate - never
			// both, so a reader cannot compare this row's Score against
			// another row's Score across two different units.
			// notInSimDB (this lane's brief, item 1): pk.Item.NotInSimDB
			// (data.go's own doc - set once per spec run by
			// markNotInSimDB, off this build's real embedded item
			// database) is true when simdb.Attach/AttachWeights' own
			// UnequipUnknown would have silently emptied this exact
			// slot before every sim this command ran ever built the
			// character, so no run anywhere actually measured this
			// item whatever MeasuredDPS claims. Forcing simDecided off
			// here is what makes "the pick stays score-decided" true
			// below (Score is never zeroed) and what the SimStatus/
			// Verified branch further down reads to keep this row
			// honest.
			notInSimDB := pk.Item.NotInSimDB
			simDecided := pk.Item.MeasuredDPS > 0 && !notInSimDB
			if simDecided {
				row.SimDPS = pk.Item.MeasuredDPS
				row.Score = 0
			}
			if notInSimDB {
				row.SimStatus = notInSimReason
			}
			// effectVerifiedInSim (rank.go), not the bare "does the
			// engine's source claim this id" hasImplementedEffect: a
			// picked item whose effect the engine truly implements but
			// this build's own simdb silently stripped before every sim
			// ran (Hand of Justice 11815 - rank.go's own doc) never had
			// its effect exercised either, so it earns the same honest
			// label an unimplemented effect already gets here.
			row.EffectUnmodelled = carriesEffect(pk.Item.candidate) && !effectVerifiedInSim(pk.Item.candidate)
			for _, tie := range pk.Ties {
				row.Ties = append(row.Ties, tieAlternative{ItemID: tie.ID, ItemName: tie.Name})
			}
			var swForSlot *swapResult
			if s, ok := swapBySlot[slot]; ok {
				swForSlot = &s
			}
			row.Alternatives = buildAlternatives(pk, slot, bySlot[slot], picks, referenceDPSPerPoint, swForSlot, setDPS)
			// This lane's brief, item 3: reconcile row.Ties against
			// row.Alternatives' own, possibly sim-measured delta for the
			// same item id - reconcileTies' own doc, above.
			row.Ties = reconcileTies(row.Ties, row.Alternatives)
			// This lane's brief (bis-ranker-integrity-16), item 6: label
			// the pick and any alternative sharing its exact display
			// name under a different id (applyLabelSuffixForNameCollisions'
			// own doc) - dedupeAlternatives (buildAlternatives, above)
			// already collapses two ALTERNATIVES sharing a name down to
			// one row, so the only collision left to catch here is the
			// pick's own name against one of its surviving alternatives.
			applyLabelSuffixForNameCollisions(&row, slotItemLevelByID(bySlot[slot], pk), slotStatsByID(bySlot[slot], pk))
			if pk.Item.HasSource {
				row.Source = pk.Item.Source.Label
				row.SourceKind = pk.Item.Source.Kind
			}
			var realSimPromotion bool
			switch {
			case notInSimDB:
				// This lane's brief, item 1: never "verified" and never
				// carrying sim_dps/dps_delta - SimStatus (set above)
				// already tells the page why, so no SwapNote text
				// duplicates it here; DPSDelta/SimDPS simply never get
				// set in this branch (both start at their zero values).
				row.Verified = false
			case erroredSlots[slot]:
				row.Verified = false
				row.SwapNote = "the runner-up's verification sim failed (an engine-side error, not a scoring one - see verify_errors); the pick is unconfirmed against it"
			default:
				row.Verified = true
				switch sw, ok := swapBySlot[slot]; {
				case ok && sw.Beat && pk.RunnerUp != nil:
					// applySwaps already promoted the runner-up into pk.Item and
					// demoted the scored pick to pk.RunnerUp: this row IS the
					// measured winner, verified by that very run.
					delta := sw.SwapDPS - sw.BaselineDPS
					row.DPSDelta = &delta
					row.SwapNote = fmt.Sprintf("beat the scored pick %s (id %d) in the sim: %s", swapNoteRunnerUpName(pk), pk.RunnerUp.ID, dpsComparisonPhrase(delta, sw.SwapDPS, sw.BaselineDPS, setDPS))
					realSimPromotion = true
				case ok:
					// This lane's brief (bis-ranker-integrity-6), item 6:
					// sw.BaselineDPS is verifyBand's ONE shared baseline
					// sim, run ONCE for the whole band before any slot's
					// swap is tried - every unpromoted slot with a
					// runner-up carries this identical number. It equals
					// the finished set's own setDPS exactly when nothing
					// else in the band promoted (the common case: a prior
					// lane's own druid-feral band 50 repro, sw.BaselineDPS
					// 155.8 == setDPS 155.8), but the moment some OTHER
					// slot's swap DOES promote, applySwaps re-measures the
					// whole set and returns a NEW setDPS the shared
					// baseline never reflects - shaman-elemental band 60
					// Horde's own repro, eight unrelated slots all
					// publishing that one stale pre-promotion baseline as
					// though it were their own row's real finished-set
					// number, none of them equal to the band's own
					// published set_dps. Publish it only when it still
					// matches (finishedSetEpsilon, the same bar
					// dpsComparisonPhrase already holds the SwapNote text
					// to); otherwise omit sim_dps entirely rather than
					// print a number that no longer describes the
					// finished set - dps_delta (below) already carries
					// the one number that stays true regardless.
					if simDecided && math.Abs(sw.BaselineDPS-setDPS) <= finishedSetEpsilon {
						row.SimDPS = sw.BaselineDPS
					} else if simDecided {
						row.SimDPS = 0
					}
					// data-followups-10 lane, 2026-09-30, item 7: also
					// require sw.Significant (verify.go's own doc) -
					// this "kept the pick" branch is exactly the shape
					// that kept publishing a false-precision "confirmed"
					// DPSDelta for a pure sampling-noise gap (Serpent
					// Gloves/Gloves of the Fang, hunter-marksmanship band
					// 20, three sweeps running) with nothing checking
					// whether the measured gap cleared the trial's own
					// error bar. An insignificant swap leaves DPSDelta/
					// SwapNote unset here - the row stays Verified (a
					// real sim genuinely ran) without a false-confidence
					// number attached to it.
					if pk.RunnerUp != nil && !alternativeCarriesRealEvidence(row.Alternatives) && sw.Significant() {
						// bis-ranker-integrity-3, 2026-09-29, this lane's brief
						// item 3: a real sim DID run for this slot (verifyBand's
						// own swap pass) even though nothing was promoted - the
						// row otherwise has NO trace anywhere that this pick was
						// ever tested (never sim-decided, since nothing beat it
						// - SimDPS above stays 0; and the one candidate that WAS
						// tested is often not a visible Alternative at all,
						// either crowded out of the top alternativesLimit rows
						// by raw score() or - shaman-enhancement's own main_hand
						// at bands 20/30/40, this lane's own dogfood - correctly
						// excluded as the slot's own pair-mate's item: Diamond
						// Hammer is main_hand's real, tested runner-up here, but
						// it is ALSO off_hand's own pick, so buildAlternatives
						// rightly never offers the same physical weapon back as
						// a main_hand fallback). "Verified: true" published with
						// nothing anywhere backing it up read exactly like
						// "verified without a sim" to three sweeps of review in
						// a row, even though a sim genuinely ran; naming the
						// real numbers here, the same way a promoted swap's own
						// SwapNote already does, is this fix.
						delta := sw.BaselineDPS - sw.SwapDPS
						row.DPSDelta = &delta
						row.SwapNote = fmt.Sprintf("confirmed by the sim against %s (id %d): kept the pick, %s", swapNoteRunnerUpName(pk), pk.RunnerUp.ID, dpsComparisonPhrase(delta, sw.BaselineDPS, sw.SwapDPS, setDPS))
					}
				}
			}
			// bis-ranker-integrity-7 lane, item 2: the "case ok" branch
			// above (bis-ranker-integrity-6) only ever re-checked
			// sw.BaselineDPS against setDPS for a slot verifyBand tried
			// and did NOT promote. Two other paths publish an absolute
			// SimDPS that never goes through that check at all, and the
			// seventh wow-player sweep's hybrids report still finds both:
			//   - "case ok && sw.Beat" (just above): a promoted swap's own
			//     SwapDPS is measured with ONLY this slot swapped against
			//     the pre-swap baseline (verifyBand's own doc); applySwaps
			//     promotes every winning swap independently off that same
			//     shared baseline (its own doc), so the moment a SECOND
			//     slot promotes in the same band, this slot's own SwapDPS
			//     stops describing the finished set the same way a stale
			//     shared baseline did for bis-ranker-integrity-6.
			//   - a row simDecided (line ~788, above) but never entered
			//     swapBySlot at all: trySetCompletion (sets.go) sets
			//     MeasuredDPS directly and never sets RunnerUp, so
			//     verifyBand never tries it (its own "only a slot with a
			//     RunnerUp" doc) and neither switch case above ever sees
			//     it - a set-completion pick measured before some OTHER
			//     slot's own later swap promotion changed the total is
			//     exactly the same staleness, just never caught.
			// One shared rule closes both: whatever set row.SimDPS to a
			// nonzero number, above, it is only trustworthy when it still
			// equals THIS band's own finished set_dps (finishedSetEpsilon,
			// the same bar every other SimDPS check in this function
			// already holds itself to) - otherwise the row keeps its
			// SwapNote/DPSDelta (both independently true of the one sim
			// run that produced them) and omits sim_dps rather than print
			// a number that no longer describes the finished set.
			if row.SimDPS != 0 && math.Abs(row.SimDPS-setDPS) > finishedSetEpsilon {
				row.SimDPS = 0
			}
			// This lane's brief, item 2 (original), narrowed by
			// bis-ranker-integrity-2's own items 1 and 4: a slot score()
			// alone decided (never touched by rankTrinketSlot's own
			// real-sim tournament, and never redeemed by an unmodelled
			// effect or a real swap promotion) whose winning candidate
			// still scored exactly 0 contributes nothing this spec's own
			// weights can measure at all.
			//
			// A trinket is exempted only when it carries an EffectText -
			// score() cannot value ANY trinket's stats meaningfully
			// (trinkets.go's own doc), but a trinket with a real effect
			// (modelled or not) still has unquantified value score()
			// never claimed to see, unlike a genuinely bare stat-stick
			// trinket (Rune of Perfection: spell penetration + stamina,
			// neither weighted by warrior-arms; Rune of Duty: pure
			// resistance) which this lane's brief, item 4 calls exactly
			// as zero-value as Sentinel's/Scout's Medallion was.
			//
			// A weapon slot (band.go's weaponSlots) is NEVER emptied this
			// way at all - this lane's brief, item 1: "a weapon slot
			// should essentially never be zero value... hiding the slot
			// entirely is a worse failure mode". pick.go's own
			// promoteLowValueWeapon already chose the best defensible
			// fallback among this slot's own zero-scoring candidates
			// (one carrying a weight stat, else the highest item level);
			// this just flags that fallback honestly via LowValue instead
			// of blanking the row.
			isTrinketSlot := slot == "trinket1" || slot == "trinket2"
			trinketEffectExempt := isTrinketSlot && carriesEffect(pk.Item.candidate)
			// bis-ranker-integrity-3, 2026-09-29, this lane's brief item
			// 2: simDecided alone can never gate a TRINKET - rankTrinketSlot
			// always sets MeasuredDPS to the whole set's own absolute DPS
			// (scored.MeasuredGainDPS's own doc), which is positive for
			// every trinket regardless of whether it contributes anything,
			// so "simDecided" was true for every trinket unconditionally
			// and this whole zero-value gate could never actually reach
			// one (warrior-arms band 20's Rune of Perfection, +6 spell
			// penetration/+4 stamina, no effect_text, still published as
			// "verified" for three sweeps running). A trinket's own real
			// GAIN over an empty slot (GainMeasured/MeasuredGainDPS,
			// rankTrinketSlot's own baseline sim) is what actually
			// answers the question; a non-trinket slot's simDecided
			// (rankSlotWithEffects, trySetCompletion, a swap promotion)
			// is untouched - each of those already measures a genuine
			// comparative delta, not an absolute set total.
			//
			// This lane's brief (bis-ranker-integrity-16), item 4:
			// !trinketEffectExempt is required here, not just ORed in
			// further down the switch - a trinket rankTrinketSlot
			// tournament-ranked always has GainMeasured true (this
			// doc's own point above), so trinketLowGain fired for EVERY
			// low-gain trinket regardless of trinketEffectExempt, and
			// this case's own position ahead of `case ... ||
			// trinketEffectExempt || ...` in the switch below meant
			// that later case never got a chance to run for one: a
			// trinket with a real, unmodelled defensive effect_text
			// (Smoking Heart of the Mountain's armor buff, mage-arcane/
			// fire/frost band 50 Alliance trinket2's own tournament
			// winner) was emptied by this case exactly like a bare
			// stat-only trinket with nothing behind it at all, hiding
			// it next to its own tied alternative (Uther's Strength,
			// the same "real defensive effect, no measurable DPS gain"
			// shape) with no way for either to publish. druid-balance
			// band 50 Alliance trinket2 (also Uther's Strength) never
			// exposed this: it is EVERY band's own outright winner
			// there, but ranked via the SAME tournament and gated by
			// the SAME trinketLowGain - the two specs differ only in
			// which effect-bearing trinket happened to win, not in
			// whether this gate would have hidden it. Excluding an
			// effect-bearing item here restores trinketEffectExempt's
			// own stated purpose for every trinket the tournament ever
			// crowns, not only the ones lucky enough to never trigger
			// this case in the first place.
			// A relic is ranked by measured gain over an empty slot exactly as a
			// trinket is (relics.go), so it is gated on that gain the same way:
			// a modelled relic whose simulated gain is noise is no pick.
			isGainRankedSlot := isTrinketSlot || isRelicCandidate(pk.Item.candidate)
			trinketLowGain := isGainRankedSlot && !trinketEffectExempt && pk.Item.GainMeasured && !trinketGainSignificant(pk.Item.MeasuredGainDPS, pk.Item.MeasuredGainStdErr)
			// A relic (libram/idol/totem) whose one real selling point -
			// its engraved effect - the engine cannot simulate at all
			// never got a fair shot at EITHER exemption below: score()
			// cannot value it (same as a trinket, but a relic is not a
			// trinket slot, so trinketEffectExempt never applies to it),
			// and it was never sim-ranked (rankSlotWithEffects only runs
			// a slot with at least one IMPLEMENTED-effect candidate in
			// the pool - rank.go's own doc). The old EffectUnmodelled
			// exemption let it fall back to score()'s plain stat total
			// as though that were the whole story (druid-balance's own
			// Idol of the Huntress, an "Improved Swipe" - a FERAL rune -
			// engrave with no Balance-relevant stats at all, picked at 4
			// of 5 bands because it was simply never emptied). This lane's
			// brief, item 2: empty it with its own reason instead, so the
			// report can name exactly which relic ids need their effect
			// added to effectids_generated.go rather than silently
			// publishing an unmeasured guess as BiS.
			relicEffectUnmodelled := row.EffectUnmodelled && !simDecided && !realSimPromotion && isRelicCandidate(pk.Item.candidate)
			switch {
			case relicEffectUnmodelled:
				row = slotRow{Slot: slot, EmptyReason: effectNotModelledReason, EffectUnmodelled: true, FactionNote: publishableFactionNote(pk)}
			case trinketLowGain:
				// This lane's brief (bis-ranker-integrity-12), item 1:
				// the OLD version of this case replaced row wholesale
				// with a bare slotRow, discarding row.Alternatives
				// (built above, before this gate ever runs) along with
				// the now-hidden pick - shaman-elemental band 50
				// Horde's own repro, both trinket slots publishing
				// empty with literally zero alternatives even though
				// rankTrinketSlot's own tournament simmed a real
				// shortlist and paladin-retribution band 50's trinket2
				// (Alliance) losing Fire Ruby entirely as a fallback
				// while Horde's own trinket2 row - the SAME comparison,
				// just not gated - still showed it at -0.14 DPS. A
				// trinket whose own measured gain does not clear
				// trinketGainSignificant's bar is not evidence the
				// SLOT has no other real candidate worth a player's
				// attention; the pick's own alternatives (the runner-up
				// verify.go actually simmed, chief among them) are
				// still true, checked facts about this band's other
				// shortlisted trinkets and must survive the pick being
				// hidden.
				row = slotRow{Slot: slot, EmptyReason: noDPSValueReason, Alternatives: row.Alternatives, FactionNote: publishableFactionNote(pk)}
			case row.Score != 0 || row.EffectUnmodelled || realSimPromotion || trinketEffectExempt || simDecided:
				// Not zero-value at all, or redeemed by one of the
				// existing exemptions above - the row stands as computed.
				// A trinket that IS simDecided reaches this case only when
				// trinketLowGain above did NOT match: either its measured
				// gain cleared the noise floor, or the baseline sim failed
				// and GainMeasured is false - an unmeasured gain is not
				// evidence of zero value either (tenet 8: never publish a
				// claim this command could not check).
			case weaponSlots[slot]:
				row.LowValue = true
			default:
				row = slotRow{Slot: slot, EmptyReason: noDPSValueReason, FactionNote: publishableFactionNote(pk)}
			}
		}
		rows = append(rows, row)
	}

	// This lane's brief (bis-ranker-integrity-6), item 7: this used to
	// read picks[slot].Item directly - the RAW, pre-gate pick - so a
	// slot the zero-value gate above just emptied (relicEffectUnmodelled/
	// trinketLowGain/the default noDPSValueReason case, all of which
	// replace row with a bare slotRow carrying no ItemID at all) still
	// named its old, now-published-empty item here (mage-fire band 20
	// Alliance's own repro: "neck: Sentinel's Medallion", "trinket1:
	// Rune of Perfection", "trinket2: Rune of Duty" in new_at_band while
	// every one of those slots' own row published empty). new_at_band
	// must describe what this band ACTUALLY published, so it is
	// computed from rows - the final, gated slotRow list - instead.
	var newAt []string
	for _, row := range rows {
		if row.ItemID == 0 {
			continue
		}
		var prevID int
		if previous != nil && previous[row.Slot].Item != nil {
			prevID = previous[row.Slot].Item.ID
		}
		if prevID != row.ItemID {
			newAt = append(newAt, fmt.Sprintf("%s: %s", row.Slot, row.ItemName))
		}
	}

	// setDPSPartial: this lane's brief, item 1 - see SetDPSPartial's
	// own doc on bandReport for what it means and why publishing the
	// real (partial) set_dps rather than omitting it is the honest
	// choice here.
	setDPSPartial := false
	for _, row := range rows {
		if row.SimStatus == notInSimReason {
			setDPSPartial = true
			break
		}
	}

	sampleNames := make([]string, 0, noSourceSampleSize)
	for i, c := range noSource {
		if i >= noSourceSampleSize {
			break
		}
		sampleNames = append(sampleNames, fmt.Sprintf("%d %s", c.ID, c.Name))
	}

	wrows := make([]weightRow, 0, len(weightOrder))
	for _, id := range weightOrder {
		w := weights[id]
		// The reference stat's own row is always Weight==1.0 by
		// construction (simrun.go's runWeights doc: "the reference
		// stat's own entry in its output is always exactly 1.0") --
		// isWeightSignificant's 25%-of-value error bar can still flag
		// that row insignificant on a wide error, but
		// referenceMeasurementReason (weights.go) already ran the
		// real check for this exact stat: when it returns ""
		// (weightsReason == ""), the band's reference measurement
		// itself cleared its own noise floor, and the reference row
		// IS that measurement -- it is significant by definition, not
		// by isWeightSignificant's generic bar. Every other row still
		// goes through isWeightSignificant unchanged.
		isReferenceRow := id == spec.ReferenceStat
		wrows = append(wrows, weightRow{
			Stat:   id,
			Weight: w.Weight,
			Error:  w.Error,
			// weightsReason != "" overrides isWeightSignificant for
			// every row: the band's own reference measurement makes
			// none of them trustworthy, whatever their own noise bar
			// says (weightsReason's own doc).
			Insignificant: weightsReason != "" || (!isReferenceRow && !isWeightSignificant(w)),
		})
	}

	var referenceDPSPerPointOut *float64
	if weightsReason == "" {
		v := referenceDPSPerPoint
		referenceDPSPerPointOut = &v
	}

	return bandReport{
		Spec:                 spec.Spec,
		Band:                 band,
		Faction:              faction,
		Race:                 race,
		Talents:              talents,
		TalentPoints:         talentPoints,
		Weights:              wrows,
		Slots:                rows,
		SetDPS:               setDPS,
		SetDPSPartial:        setDPSPartial,
		NoSourceCount:        len(noSource),
		NoSourceSample:       sampleNames,
		NewAtBand:            nonNil(newAt),
		WeightsRunSeconds:    weightsSeconds,
		VerifyRunSeconds:     verifySeconds,
		VerifyErrors:         verifyErrors,
		Coverage:             coverage,
		ReferenceDPSPerPoint: referenceDPSPerPointOut,
		WeightsReason:        weightsReason,
		ScoreUnit:            scoreUnitReferenceStatPoints,
	}
}

// coverageTotals sums every slot's coverageRow into one band-wide
// figure: how many eligible items this band+faction saw across every
// slot, and how many of those had a source. Used by main.go's own
// one-line nightly log summary (this lane's brief, guardrail A) -
// giving a reader watching the run the same "N of M sourced" honesty
// the page's per-band coverage table will show, without having to
// open the JSON.
func coverageTotals(coverage map[string]coverageRow) (eligible, sourced int) {
	for _, row := range coverage {
		eligible += row.Eligible
		sourced += row.Sourced
	}
	return eligible, sourced
}

// worstCoveredSlot names the slot with the lowest sourced/eligible
// ratio (ties broken by the most eligible items, then by slot name,
// so the result is deterministic across runs) - the one line's own
// "and worst is X" clause points a reader straight at the slot most
// worth a data lane's attention, rather than making them scan the
// full per-slot map for it. A slot with zero eligible items is
// excluded: 0/0 is not a coverage gap, it is "nothing exists here to
// gate at all".
func worstCoveredSlot(coverage map[string]coverageRow) (slot string, row coverageRow, ok bool) {
	for _, s := range slotOrder {
		r, present := coverage[s]
		if !present || r.Eligible == 0 {
			continue
		}
		if !ok || ratio(r) < ratio(row) {
			slot, row, ok = s, r, true
		}
	}
	return slot, row, ok
}

func ratio(r coverageRow) float64 {
	if r.Eligible == 0 {
		return 1
	}
	return float64(r.Sourced) / float64(r.Eligible)
}

// coverageSummary is the one-line-per-band nightly log string this
// lane's brief asks for: the band-wide total, plus the single
// worst-covered slot so a reader does not have to open the JSON to
// see which slot most needs a data lane's attention.
func coverageSummary(coverage map[string]coverageRow) string {
	eligible, sourced := coverageTotals(coverage)
	slot, row, ok := worstCoveredSlot(coverage)
	if !ok {
		return fmt.Sprintf("%d/%d eligible items sourced", sourced, eligible)
	}
	return fmt.Sprintf("%d/%d eligible items sourced; worst slot %s (%d/%d)", sourced, eligible, slot, row.Sourced, row.Eligible)
}

// titleCase upper-cases a single lower-kebab word's first letter -
// "horde" -> "Horde" - for the markdown's faction headings. strings.
// Title is deprecated (it does not handle multi-word input correctly,
// which is not a problem here, but this project's lint config flags
// its use regardless), so this is the one-word case written out.
func titleCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// specReport is data/builds/<build>/bis/<spec>.json's whole shape -
// lane bis-web's read contract (this lane's brief, step 3): every
// bandReport field name is kept exactly as the prototype defined it;
// engine_version is the one field this wrapper adds beyond what wraps
// the array (spec, build, generated_at, bands).
type specReport struct {
	Spec          string       `json:"spec"`
	Build         string       `json:"build"`
	EngineVersion string       `json:"engine_version"`
	GeneratedAt   string       `json:"generated_at"`
	Bands         []bandReport `json:"bands"`
	// Presets states what each non-bare preset applied, by name: the
	// request-vocabulary ids with a label per id.
	Presets map[string]request.ResolvedPreset `json:"presets"`
	// HealProfile is the incoming-damage profile every healer entry in the
	// file was measured under, reasons included. Set only for a healer.
	HealProfile *request.HealProfile `json:"heal_profile,omitempty"`
}

// specReportOption adds an optional block to a spec report.
type specReportOption func(*specReport)

// withHealProfile publishes the profile a healer's file was measured under.
func withHealProfile(profile *request.HealProfile) specReportOption {
	return func(r *specReport) { r.HealProfile = profile }
}

// writeSpecReport writes path per the specReport contract above.
// GeneratedAt is RFC3339, in UTC so two runs on different machines (a
// dev's laptop, the nightly workflow's runner) produce comparable
// timestamps rather than each in its own local zone.
func writeSpecReport(path, spec, build string, reports []bandReport, presets map[string]request.ResolvedPreset, options ...specReportOption) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	out := specReport{
		Spec:          spec,
		Build:         build,
		EngineVersion: enginever.Version,
		GeneratedAt:   time.Now().UTC().Format(time.RFC3339),
		Bands:         reports,
		Presets:       presets,
	}
	for _, option := range options {
		option(&out)
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

// writeMarkdown renders every band+faction report for one spec, most
// recent band last, grouped by faction, in the shape this lane's
// brief asks for: a table per band (slot, item, source with faction
// badge, score, verified) and a "New at L" list.
func writeMarkdown(path string, spec specInfo, reports []bandReport) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Leveling BiS: %s\n\n", spec.Name)
	fmt.Fprintf(&b, "Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.\n\n")
	// Caster sweep, bis-ranker-integrity-4 lane, item 3: this sentence
	// used to name Earthweave Cloak as one of the two unified-crit
	// examples, but the build's own item file no longer states that
	// stat block for it (data/builds/1.60.1.70009/items/mage.json's own
	// Earthweave Cloak carries agility/hit, no crit at all today) -
	// worded generically instead of pinning it to whichever items
	// happen to carry this shape in any one build, which a data refresh
	// can freely change without making this sentence false.
	b.WriteString("Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.\n\n")

	byFaction := map[string][]bandReport{}
	var factions []string
	for _, r := range reports {
		if _, ok := byFaction[r.Faction]; !ok {
			factions = append(factions, r.Faction)
		}
		byFaction[r.Faction] = append(byFaction[r.Faction], r)
	}
	sort.Strings(factions)

	for _, faction := range factions {
		fmt.Fprintf(&b, "## %s\n\n", titleCase(faction))
		for _, r := range byFaction[faction] {
			fmt.Fprintf(&b, "### Band %d%s (%s, %s)\n\n", r.Band, presetHeadingSuffix(r.Preset), r.Race, r.Talents)
			fmt.Fprintf(&b, "Set DPS (verified): %.1f. Weights run: %.1fs. Verify run: %.1fs. %d eligible items had no known source.\n\n",
				r.SetDPS, r.WeightsRunSeconds, r.VerifyRunSeconds, r.NoSourceCount)

			fmt.Fprintf(&b, "Stat weights (normalized to %s = 1.0, error under %.0f%% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1%% hit/crit/dodge/parry/block/defense): ", spec.ReferenceStat, significanceErrorFraction*100)
			parts := make([]string, len(r.Weights))
			for i, w := range r.Weights {
				unit := ""
				if w.Unit == "rating" {
					unit = fmt.Sprintf(" per rating point (%.0f rating = 1%%, %.3f per %%)", w.RatingFactor, w.WeightPerPercent)
				}
				if w.Insignificant {
					parts[i] = fmt.Sprintf("%s=not significant (%.3f ± %.3f)%s", w.Stat, w.Weight, w.Error, unit)
				} else {
					parts[i] = fmt.Sprintf("%s=%.3f ± %.3f%s", w.Stat, w.Weight, w.Error, unit)
				}
			}
			fmt.Fprintln(&b, strings.Join(parts, ", "))
			b.WriteString("\n")

			// This lane's brief, item 3: a bare "Score" header, with
			// numbers like "7.0" sitting next to "sim-verified (244.9
			// DPS)" in the same column, reads as two unrelated units
			// with neither one named (fifth wow-player sweep, casters
			// finding 3/8). The column header now names its own unit -
			// score() is a weighted sum in the spec's own reference
			// stat (spec.ReferenceStat, data/curated/specs.json), not a
			// DPS number at all.
			fmt.Fprintf(&b, "| Slot | Item | Source | Score (%s points) | Verified | Alternatives |\n", spec.ReferenceStat)
			b.WriteString("|---|---|---|---|---|---|\n")
			for _, row := range r.Slots {
				item := "-"
				source := "-"
				score := ""
				verified := ""
				alternatives := ""
				if row.ItemID != 0 {
					item = fmt.Sprintf("%s (%d)", row.ItemName, row.ItemID)
					if len(row.Ties) > 0 {
						// This lane's brief, item 4: priest-shadow band 20
						// main_hand alone ties 107 (alliance) / 117 (horde)
						// other items - the web page never reads `ties` at
						// all and caps Alternatives at 3 (alternativesLimit),
						// so this .md-only cell printing every tied name made
						// one table cell longer than the rest of the table
						// combined (wow-player sweep, day3/player-review-33/
						// casters.md finding 8). Capped at mdTieDisplayLimit
						// names, "and N more" for the rest.
						shown := row.Ties
						var more int
						if len(shown) > mdTieDisplayLimit {
							more = len(shown) - mdTieDisplayLimit
							shown = shown[:mdTieDisplayLimit]
						}
						alts := make([]string, len(shown))
						for i, t := range shown {
							alts[i] = fmt.Sprintf("%s (%d)", t.ItemName, t.ItemID)
						}
						tieText := strings.Join(alts, ", ")
						if more > 0 {
							tieText += fmt.Sprintf(", and %d more", more)
						}
						item += " (or " + tieText + ")"
					}
					// This lane's brief, item 1 (bis-ranker-integrity-3,
					// 2026-09-29): a sim-decided row (SimDPS set, Score
					// deliberately zeroed - buildReport's own doc) used
					// to fall straight into FormatFloat(row.Score, ...)
					// here and print "0.0" right next to "Verified:
					// yes" (hybrids sweep, item 5: Dawn's Edge/Ebon
					// Hand/Annihilator all read this way) - the exact
					// "this item does nothing" misreading the sweep
					// flagged. A sim-decided row now prints the real,
					// measured DPS instead; only a score()-decided row
					// (or a genuinely zero-scoring LowValue weapon
					// fallback) prints a bare score number at all.
					switch {
					case row.SimDPS != 0:
						score = fmt.Sprintf("sim-verified (%.1f DPS)", row.SimDPS)
					case row.Score == 0 && row.DPSDelta != nil:
						// This lane's brief (bis-ranker-integrity-6), item
						// 6: a simDecided row whose own sim_dps buildReport
						// withheld (its measurement no longer matches the
						// band's own finished set_dps - buildReport's own
						// doc on this exact case) still has Score
						// deliberately zeroed, the same convention as an
						// ordinary sim-decided row (the comment above this
						// switch). Falling through to the reference-points
						// case below would print "0.0 <ref> points (0.00
						// DPS)" - the identical misreading this switch
						// already exists to prevent - so a row in exactly
						// this state publishes its one still-trustworthy
						// number, dps_delta, instead.
						score = fmt.Sprintf("sim-verified (%+.1f DPS vs the runner-up, not corroborated against the finished set)", *row.DPSDelta)
					case row.Score == 0 && !row.LowValue:
						// This lane's brief, item 4: 46 occurrences across
						// 8 caster specs (wow-player sweep, day3/player-
						// review-33/casters.md finding 9) - a pick with no
						// score AT ALL (Score deliberately zeroed because
						// the pick came from a real sim tournament -
						// buildReport's own doc - not because it truly
						// scores zero: that case sets LowValue instead,
						// excluded above) whose own SimDPS/DPSDelta both
						// happened to be withheld too (the finishedSetEpsilon
						// staleness guard, or an alternative already showing
						// the same evidence - buildReport's own doc on both)
						// fell through the two cases above and printed the
						// literal string "0.0 spell_power points (0.00 DPS)"
						// next to "Verified: yes" - read by a player as "this
						// item contributes zero DPS", the exact misreading
						// the SimDPS/DPSDelta cases above already exist to
						// prevent for every OTHER state that clears row.Score
						// == 0. Naming what actually happened - a real sim
						// tournament decided this slot - instead of a
						// fabricated zero score.
						score = "sim-decided (no score - a real sim tournament chose this pick)"
					case r.ReferenceDPSPerPoint != nil:
						// This lane's brief, item 3: the DPS conversion
						// beside the raw reference-stat-points number,
						// using this band's own measured
						// reference_dps_per_point (bandReport's own
						// doc) - the same rate the page can use to turn
						// "Strength 1.99" into "0.14 DPS per point"
						// elsewhere. Never a bare, unitless number.
						score = fmt.Sprintf("%.1f %s points (%.2f DPS)", row.Score, spec.ReferenceStat, row.Score*(*r.ReferenceDPSPerPoint))
					default:
						// No trustworthy reference_dps_per_point this
						// band (weights.go's referenceMeasurementReason
						// fired - r.WeightsReason is set) - still label
						// the unit even without a DPS conversion to
						// show beside it.
						score = fmt.Sprintf("%s %s points", strconv.FormatFloat(row.Score, 'f', 1, 64), spec.ReferenceStat)
					}
					verified = "yes"
					if !row.Verified {
						verified = "no - " + row.SwapNote
					}
					if row.Source != "" {
						source = fmt.Sprintf("%s [%s]", row.Source, row.SourceKind)
					} else {
						source = "no known source"
					}
					// Alternatives (this lane's brief, item 1): the
					// next-best sourced candidates after the pick, so a
					// player reading the human-facing markdown table -
					// not only a script reading the JSON - sees a real
					// fallback when the pick's own source is out of
					// reach.
					alternatives = "-"
					if len(row.Alternatives) > 0 {
						alts := make([]string, len(row.Alternatives))
						for i, a := range row.Alternatives {
							verifiedNote := ""
							if a.Verified {
								verifiedNote = ", sim-verified"
							}
							alts[i] = fmt.Sprintf("%s (%d, %+.2f DPS%s) [%s]", a.ItemName, a.ItemID, a.DPSDelta, verifiedNote, a.SourceKind)
						}
						alternatives = strings.Join(alts, "; ")
					}
				}
				fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s |\n", row.Slot, item, source, score, verified, alternatives)
			}
			b.WriteString("\n")

			if len(r.NewAtBand) > 0 {
				fmt.Fprintf(&b, "**New at %d:** %s\n\n", r.Band, strings.Join(r.NewAtBand, "; "))
			} else {
				b.WriteString("**New at this band:** nothing changed from the previous band.\n\n")
			}

			if len(r.NoSourceSample) > 0 {
				fmt.Fprintf(&b, "No-known-source sample (%d of %d, see the JSON for more): %s\n\n",
					len(r.NoSourceSample), r.NoSourceCount, strings.Join(r.NoSourceSample, "; "))
			}
		}
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

// nonNil is xs, or an empty slice for nil: the JSON contract promises an
// array, and encoding/json writes a nil slice as null, which the site's
// loader and the addon pipeline would then have to special-case.
func nonNil(xs []string) []string {
	if xs == nil {
		return []string{}
	}
	return xs
}

// presetHeadingSuffix labels a non-bare entry's heading in the markdown.
func presetHeadingSuffix(preset string) string {
	if preset == presetBare {
		return ""
	}
	return ", " + preset + " preset"
}
