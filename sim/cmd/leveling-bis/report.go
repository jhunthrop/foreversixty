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
func dpsComparisonPhrase(delta, primaryDPS, otherDPS, setDPS float64) string {
	if math.Abs(primaryDPS-setDPS) <= finishedSetEpsilon {
		return fmt.Sprintf("+%.1f DPS (%.1f vs %.1f set DPS)", delta, primaryDPS, otherDPS)
	}
	return fmt.Sprintf("+%.1f DPS over it", delta)
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
	// EffectUnmodelled is true when the picked item carries an
	// effect_text the engine does NOT implement (effectids_generated.go)
	// -- this lane's brief, item 3: such a candidate is still scored on
	// its plain stats (score.go never saw the effect either way), but
	// the page shows "proc not simulated" on it rather than letting the
	// Score/Verified columns imply the whole item was accounted for.
	// False (omitted) for a candidate with no effect_text at all, and
	// for one whose effect IS implemented -- see rank.go's
	// hasImplementedEffect, the single predicate both this flag and the
	// effect-verification pass itself read.
	EffectUnmodelled bool `json:"effect_unmodelled,omitempty"`
	// Ties is every other candidate that scored identically to this
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

// trinketZeroGainThresholdDPS is the noise floor rankTrinketSlot's own
// baseline-relative gain (scored.MeasuredGainDPS) must clear before a
// trinket counts as having any real, measured value at all - this
// lane's brief, item 2: "a trinket or relic whose measured/estimated
// gain is < 0.05 DPS publishes empty with no_dps_value". Below this, a
// trinket's own real contribution (Rune of Perfection's spell
// penetration/stamina) is indistinguishable from noise, not a genuine
// DPS gain a player should read as a reason to chase this item.
//
// Known limitation (bis-ranker-integrity-3 lane report): a trinket
// whose ONLY stat is completely inert for this spec (Rune of
// Perfection/Rune of Duty/Onyxia Blood Talisman - nothing the engine's
// resource model reads at all) measures bit-identical to the baseline,
// exactly 0.0000, every time - this threshold catches those
// deterministically. A trinket carrying Spirit specifically (Ankh of
// Life) can still perturb a MANA-using spec's own RNG consumption
// order (a slightly different mp5 tick timing shifts which random
// draws a 100-iteration trinket-rank sim happens to make) enough to
// read anywhere from small-negative to several DPS on either side of
// this threshold, inconsistently across bands, for paladin-retribution
// and shaman-enhancement specifically (dogfooded directly: Ankh of
// Life measured exactly 0.0000 at every band for druid-feral, which
// spends no mana in Cat Form, but 0.03-1.3 DPS for the two mana-using
// hybrids) - this is the same scale of noise verify.go's own swapMargin
// (1%) already exists to filter for the swap pass, at 300 iterations;
// this pass runs at trinketRankIterations (100, a deliberately smaller
// nightly-budget spend) and could not fully suppress it within this
// lane's scope. The fix as shipped is still strictly better than
// before it (previously nothing here could ever fire for ANY trinket -
// see rankTrinketSlot's own doc), and reliably closes the exact,
// reported defect (a bare off-axis-stat trinket like Rune of
// Perfection). A mana-using hybrid's own Ankh of Life pick may still
// occasionally survive on sim noise; a future lane could raise this
// pass's own iteration count or switch to a relative margin (matching
// swapMargin's own convention) if that residual case needs closing too.
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

// tieAlternative is one equally-scored item slotRow.Ties names.
type tieAlternative struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
}

// alternativesLimit bounds how many candidates slotRow.Alternatives
// carries beyond the pick itself (this lane's brief: "the next best 3
// candidates by score after the pick") - enough for a player who
// cannot get the picked item's own source to see a genuine fallback,
// without ballooning the JSON with a slot's whole candidate pool.
const alternativesLimit = 3

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
	seen := map[int]bool{pk.Item.ID: true}
	excluded := func(c scored) bool {
		if seen[c.ID] {
			return true
		}
		return isPairMateItem(c)
	}
	add := func(out []alternativeRow, c scored) []alternativeRow {
		seen[c.ID] = true
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
		return append(out, alternativeRow{
			ItemID:     c.ID,
			ItemName:   c.Name,
			SourceKind: c.Source.Kind,
			Source:     c.Source.Label,
			DPSDelta:   dpsDelta,
		})
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
		if len(out) >= alternativesLimit {
			break
		}
		if excluded(tie) || !realAlternative(tie) {
			continue
		}
		out = add(out, tie)
	}
	for _, c := range list {
		if len(out) >= alternativesLimit {
			break
		}
		if excluded(c) || !realAlternative(c) {
			continue
		}
		out = add(out, c)
	}

	if sw != nil && pk.RunnerUp != nil && !isPairMateItem(*pk.RunnerUp) {
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

	sort.SliceStable(out, func(i, j int) bool {
		if out[i].DPSDelta != out[j].DPSDelta {
			return out[i].DPSDelta > out[j].DPSDelta
		}
		return out[i].ItemID < out[j].ItemID
	})
	return out
}

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
// BaselineDPS by more than swapMargin), but NOT always true when
// !sw.Beat: swapMargin (verify.go) deliberately keeps the scored pick
// on a runner-up that measured HIGHER but not by enough to clear the
// noise margin (verify.go's own Serpent's Shoulders/Mantle of Honor
// example, 78.7 vs 78.6), which left this branch computing a genuinely
// positive SwapDPS-BaselineDPS and publishing it as a "verified"
// alternative's dps_delta - the exact contract violation the
// controller's own review named (bis-ranker-integrity-2 lane, item 9:
// "every alternative's dps_delta <= 0 after the swap stage, verified
// or not"; 55 slots across ret/feral/enhancement, the same item pair
// flipping sign between factions on a coin-flip-sized margin). Capped
// at 0 here for exactly that case: a runner-up that did not clear the
// promotion bar is reported as "at best a tie", never as a positive,
// sim-measured win the pick was not actually given.
func swapMeasuredDelta(sw swapResult) float64 {
	if sw.Beat {
		return sw.BaselineDPS - sw.SwapDPS
	}
	delta := sw.SwapDPS - sw.BaselineDPS
	if delta > 0 {
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
	Spec              string      `json:"spec"`
	Band              int         `json:"band"`
	Faction           string      `json:"faction"`
	Race              string      `json:"race"`
	Talents           string      `json:"talents"`
	TalentPoints      int         `json:"talent_points"`
	Weights           []weightRow `json:"weights"`
	Slots             []slotRow   `json:"slots"`
	SetDPS            float64     `json:"set_dps"`
	NoSourceCount     int         `json:"no_source_count"`
	NoSourceSample    []string    `json:"no_source_sample,omitempty"`
	NewAtBand         []string    `json:"new_at_band"`
	WeightsRunSeconds float64     `json:"weights_run_seconds"`
	VerifyRunSeconds  float64     `json:"verify_run_seconds"`
	VerifyErrors      []string    `json:"verify_errors,omitempty"`
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
}

// scoreUnitReferenceStatPoints is bandReport.ScoreUnit's only value
// today - see that field's own doc.
const scoreUnitReferenceStatPoints = "reference_stat_points"

type weightRow struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
	// Error is this weight's standard error, in the same units as
	// Weight (sim/adapter.Weights' own StatWeight.Error, carried
	// through unchanged) -- the page shows it as "±". See
	// isWeightSignificant's own doc for why this command publishes it
	// at all when the general-purpose /sim/weights tool's own
	// significance bar (sim/adapter.go's Insignificant field) is more
	// lenient than the one this command applies below.
	Error float64 `json:"error"`
	// Insignificant is this command's OWN, stricter call (see
	// isWeightSignificant), not sim/adapter's Insignificant field: the
	// page greys this row out and the Pawn/planner consumers of this
	// JSON should not treat it as a real number.
	Insignificant bool `json:"insignificant"`
}

// significanceErrorFraction and isWeightSignificant now live in
// weights.go, alongside effectiveWeights -- the one function that
// turns a raw stat-weights result into the plain numbers score()
// ranks by, zeroing exactly the weights this file's own
// isWeightSignificant call marks Insignificant below.

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
		row := slotRow{Slot: slot}
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
			row.Score = pk.Item.Score
			// This lane's brief, item 7: a candidate a real sim actually
			// measured (MeasuredDPS > 0 - trinkets.go/rank.go/sets.go's
			// own tournaments, or a swap promotion below) publishes that
			// measurement instead of score()'s stat estimate - never
			// both, so a reader cannot compare this row's Score against
			// another row's Score across two different units.
			simDecided := pk.Item.MeasuredDPS > 0
			if simDecided {
				row.SimDPS = pk.Item.MeasuredDPS
				row.Score = 0
			}
			row.EffectUnmodelled = pk.Item.EffectText != "" && !hasImplementedEffect(pk.Item.candidate)
			for _, tie := range pk.Ties {
				row.Ties = append(row.Ties, tieAlternative{ItemID: tie.ID, ItemName: tie.Name})
			}
			var swForSlot *swapResult
			if s, ok := swapBySlot[slot]; ok {
				swForSlot = &s
			}
			row.Alternatives = buildAlternatives(pk, slot, bySlot[slot], picks, referenceDPSPerPoint, swForSlot, setDPS)
			if pk.Item.HasSource {
				row.Source = pk.Item.Source.Label
				row.SourceKind = pk.Item.Source.Kind
			}
			var realSimPromotion bool
			switch {
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
					row.SwapNote = fmt.Sprintf("beat the scored pick %s (id %d) in the sim: %s", pk.RunnerUp.Name, pk.RunnerUp.ID, dpsComparisonPhrase(delta, sw.SwapDPS, sw.BaselineDPS, setDPS))
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
					if pk.RunnerUp != nil && !alternativeCarriesRealEvidence(row.Alternatives) {
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
						row.SwapNote = fmt.Sprintf("confirmed by the sim against %s (id %d): kept the pick, %s", pk.RunnerUp.Name, pk.RunnerUp.ID, dpsComparisonPhrase(delta, sw.BaselineDPS, sw.SwapDPS, setDPS))
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
			trinketEffectExempt := isTrinketSlot && pk.Item.EffectText != ""
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
			trinketLowGain := isTrinketSlot && pk.Item.GainMeasured && pk.Item.MeasuredGainDPS < trinketZeroGainThresholdDPS
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
				row = slotRow{Slot: slot, EmptyReason: effectNotModelledReason, EffectUnmodelled: true}
			case trinketLowGain:
				row = slotRow{Slot: slot, EmptyReason: noDPSValueReason}
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
				row = slotRow{Slot: slot, EmptyReason: noDPSValueReason}
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
}

// writeSpecReport writes path per the specReport contract above.
// GeneratedAt is RFC3339, in UTC so two runs on different machines (a
// dev's laptop, the nightly workflow's runner) produce comparable
// timestamps rather than each in its own local zone.
func writeSpecReport(path, spec, build string, reports []bandReport) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	out := specReport{
		Spec:          spec,
		Build:         build,
		EngineVersion: enginever.Version,
		GeneratedAt:   time.Now().UTC().Format(time.RFC3339),
		Bands:         reports,
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
			fmt.Fprintf(&b, "### Band %d (%s, %s)\n\n", r.Band, r.Race, r.Talents)
			fmt.Fprintf(&b, "Set DPS (verified): %.1f. Weights run: %.1fs. Verify run: %.1fs. %d eligible items had no known source.\n\n",
				r.SetDPS, r.WeightsRunSeconds, r.VerifyRunSeconds, r.NoSourceCount)

			fmt.Fprintf(&b, "Stat weights (normalized to %s = 1.0, error under %.0f%% of the weight to publish - see report.go's isWeightSignificant): ", spec.ReferenceStat, significanceErrorFraction*100)
			parts := make([]string, len(r.Weights))
			for i, w := range r.Weights {
				if w.Insignificant {
					parts[i] = fmt.Sprintf("%s=not significant (%.3f ± %.3f)", w.Stat, w.Weight, w.Error)
				} else {
					parts[i] = fmt.Sprintf("%s=%.3f ± %.3f", w.Stat, w.Weight, w.Error)
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
						alts := make([]string, len(row.Ties))
						for i, t := range row.Ties {
							alts[i] = fmt.Sprintf("%s (%d)", t.ItemName, t.ItemID)
						}
						item += " (or " + strings.Join(alts, ", ") + ")"
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
