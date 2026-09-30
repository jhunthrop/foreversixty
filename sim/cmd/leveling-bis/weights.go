package main

import (
	"fmt"
	"math"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// significanceErrorFraction is this command's own publication bar
// (2026-09-28 weights-effects lane; owner, looking at a level-20
// hunter's weights: "these stat weights look like garbage"). The
// general-purpose /sim/weights tool's own bar
// (sim/adapter.go's `Insignificant: errAmt >= math.Abs(weight)`) only
// catches a weight the error swallows entirely -- appropriate for a
// tool a player can re-run at higher precision on demand. This
// command's sweep runs a fixed, small budget (100 iterations per
// direction by default -- see weightsIterations's own flag doc) and
// publishes every weight on the BiS page's aside as a fact a player
// acts on without re-running anything, so it holds every row to a
// tighter bar: an error under 25% of the weight's own value, or the
// row reports "not significant" instead of a number nobody should
// trust.
const significanceErrorFraction = 0.25

// weightsRetryIterationsFactor is how much main.go's own retry-then-
// fallback guard (this lane's brief, item 1's second half) multiplies
// -weights-iterations by for a band whose reference stat measured "not
// positive beyond its own error" (referenceMeasurementReason):
// warlock-destruction band 60's own repro (bis-ranker-integrity-11)
// re-ran at 4x (400 iterations/direction instead of 100) and its raw
// standard error tightened from ±0.1332 to ±0.0642 - almost exactly
// the sqrt(4)=2x this sweep's own merged-sample-count math
// (adapter.Weights' own sampleCount comment) predicts - but the
// measured reference DPS/point itself stayed negative (-0.1039), so
// 4x alone did not resolve that particular band; the guard's fallback
// half (the nearest lower band's own significant weights) exists for
// exactly that case.
const weightsRetryIterationsFactor = 4

// isWeightSignificant applies significanceErrorFraction to one
// engine-reported weight. A weight of exactly zero is never
// significant (there is nothing for a 25%-of-value bar to compare
// against, and a hard-capped or unmoved stat reads back as
// weight=0, error=0 the same way sim/adapter.go's own comment
// describes for its own field).
func isWeightSignificant(w api.StatWeight) bool {
	if w.Weight == 0 {
		return false
	}
	return w.Error < significanceErrorFraction*math.Abs(w.Weight)
}

// effectiveWeights is the ONE place a raw stat-weights result
// (sim/adapter.Weights' per-stat api.StatWeight, each with its own
// Weight and Error) turns into the plain stat->weight numbers that
// score() (score.go), and anything downstream that ranks by score,
// actually multiplies against an item's stats.
//
// The rule is narrower than "zero anything isWeightSignificant calls
// insignificant" -- this lane tried that first and the controller's
// own before/after measurement rejected it: zeroing every
// insignificant weight (not just negative ones) DROPPED mage-frost's
// own verified set_dps at band 20 for both factions (alliance
// 29.8->29.2, horde 28.1->27.0) versus the un-zeroed baseline,
// because this command's 100-iteration sweep calls a real, useful
// POSITIVE weight "insignificant" more often than its 25% error bar
// should be trusted to gate ranking on -- report.go's Insignificant
// flag is calibrated for "should a player trust this printed number",
// not "should this stat be allowed to influence which item wins a
// slot". Discarding it lost real signal.
//
// A NEGATIVE weight is different: in this engine, no stat lowers a
// damage spec's own DPS, so a negative weight -- whatever
// isWeightSignificant says about it -- is measurement noise around a
// true value at or near zero, never a real "this stat hurts" signal.
// That is night-bis-sanity's own finding (~40 low-level cloth items,
// Evergreen Gloves/Featherbead Bracers among them, publishing a
// NEGATIVE score because such a noise-negative weight landed on one
// of the item's stats and score() dotted it straight into the total)
// -- so effectiveWeights zeroes exactly the weights that can only ever
// be noise (Weight <= 0) and otherwise trusts the sweep's own number,
// significant or not.
//
// The raw weights -- Weight, Error and Insignificant exactly as the
// engine reported and this command judged them -- still reach the
// published JSON unchanged (report.go's buildReport reads the
// map[string]api.StatWeight the caller passes it directly, not this
// function's output): a player or Pawn-style consumer reading the
// aside still sees the real number, its error bar and whether this
// command trusts it enough to print as a fact -- only ranking's own,
// narrower rule is decided here.
// referenceMeasurementReason reports why a band's whole weights sweep
// cannot be trusted - "" when it can (the ordinary case, every band
// but one seen so far).
//
// Every other published Weights[i].Weight is raw[i]/scale, where scale
// is this same band's raw, un-normalised DPS delta for one point of
// the spec's own reference stat (simrun.go's referenceStatRawWeight;
// bandReport.ReferenceDPSPerPoint). Dividing by a scale that is at or
// below zero flips or garbles every ratio: warlock-demonology band 60
// published reference_dps_per_point -0.189 (spell power made the raid
// dummy hit LESS hard at this band's gear - noise, not a real effect)
// with intellect 4.55, hit -5.57, crit -1.77, spell_haste -2.24, NONE
// insignificant - a caster page reading that would tell a player
// Intellect helps roughly as much as Spell Power and Hit actively
// hurts, both false. A weight born of two negative-noise raw deltas
// can even land positive (two negatives divide to a positive) and
// still be worthless - "not insignificant" here is not enough cover;
// the whole band has to publish as untrustworthy, not just each stat
// on its own noise bar (isWeightSignificant).
//
// The check is "positive beyond its own error" (this lane's brief):
// wresult[referenceStat].Error is sim/adapter.Weights' own
// errAmt/scale, and scale is exactly referenceDPSPerPoint (both read
// the identical raw array entry - simrun.go's own doc), so
// wresult[referenceStat].Error*referenceDPSPerPoint recovers the raw
// standard error without asking the engine for it a second time.
// referenceDPSPerPoint must clear that error, not just be positive,
// so a reference sitting at its own noise floor (small and positive,
// but not distinguishably so) is caught the same way a negative one
// is.
func referenceMeasurementReason(referenceStat string, wresult map[string]api.StatWeight, referenceDPSPerPoint float64) string {
	trustworthy, rawStderr := referenceMeasurementTrustworthy(referenceStat, wresult, referenceDPSPerPoint)
	if trustworthy {
		return ""
	}
	return fmt.Sprintf(
		"reference stat %s measured %.4f ± %.4f DPS per point at this band's gear - not positive beyond its own error, so no weight this band measured can be trusted",
		referenceStat, referenceDPSPerPoint, rawStderr,
	)
}

// referenceMeasurementTrustworthy is the arithmetic
// referenceMeasurementReason and fallbackWeightsReason (below) both
// build their wording from - split out so main.go's retry-then-fallback
// guard (this lane's brief, item 1's second half) and the fallback
// band's own reason text can each report the SAME raw standard error
// referenceMeasurementReason already computed, rather than recomputing
// it with a second formula that could drift from this one.
func referenceMeasurementTrustworthy(referenceStat string, wresult map[string]api.StatWeight, referenceDPSPerPoint float64) (trustworthy bool, rawStderr float64) {
	ref, ok := wresult[referenceStat]
	if !ok {
		// Every real request's WeightStats carries its own
		// ReferenceStat (weightsRequest's own doc), so this is
		// unreachable in production; leave the band alone rather than
		// invent a reason for a case that cannot happen with real
		// data.
		return true, 0
	}
	rawStderr = math.Abs(ref.Error * referenceDPSPerPoint)
	return positiveBeyondError(referenceDPSPerPoint, rawStderr), rawStderr
}

// positiveBeyondError is this command's one shared significance test -
// used by referenceMeasurementTrustworthy above (a band's own
// reference-stat measurement, this lane's brief item 1) and by
// report.go's trinketLowGain gate (a trinket's own measured DPS gain,
// item 2): delta is trusted only when it clears its own standard
// error, not merely when it is positive - a delta sitting inside its
// own error bar (small and positive, or negative) is measurement
// noise, whichever of the two quantities it happens to be. Strict `>`,
// not `>=`: a delta exactly AT its own error (or a zero-noise 0-vs-0)
// is not distinguishably positive either, the same boundary
// sim/adapter.Weights' own Insignificant flag draws.
func positiveBeyondError(delta, stdErr float64) bool {
	return delta > stdErr
}

// trinketGainSignificanceMultiplier is report.go's trinketLowGain gate
// own bar on a trinket's MeasuredGainDPS/MeasuredGainStdErr
// (trinkets.go) - bis-ranker-integrity-11's brief, item 2. A bare 1x
// (positiveBeyondError's own bar, ~84% one-sided confidence) is not
// tight enough here: dogfooded directly (this lane's report) at
// trinketRankIterations (100), a trinket with NO real DPS relevance at
// all for the spec wearing it - Fire Ruby (a mage-only Fire Ward/Fire
// Blast interaction) on hunter-beast-mastery band 50 Alliance - still
// measured gain 0.773 against its own combined stdErr 0.7125 (a
// same-magnitude coin flip against Sanctified Orb's genuinely-noise
// 0.691/0.699 right next to it), clearing a bare 1x bar and winning
// the slot anyway. 2x (~95% one-sided confidence, the ordinary
// scientific convention for "not just noise") correctly fails that
// same case (0.773 < 1.425) while still passing every genuine small
// gain this lane re-measured (Frozen Heart of the Mountain's own +9
// Hit: 3.03 DPS against a 0.98 stdErr for rogue-assassination band 50
// Horde, 4.11 against 1.64 for band 60 Alliance - both comfortably
// beyond 2x).
const trinketGainSignificanceMultiplier = 2.0

// trinketGainSignificant is report.go's trinketLowGain gate: whether a
// trinket's own measured DPS gain (over the no-trinket baseline,
// trinkets.go) clears trinketGainSignificanceMultiplier times its own
// combined standard error - stricter than positiveBeyondError's bare
// 1x (see trinketGainSignificanceMultiplier's own doc for why this
// gate needs the wider margin and that one does not).
func trinketGainSignificant(gainDPS, gainStdErr float64) bool {
	return gainDPS > trinketGainSignificanceMultiplier*gainStdErr
}

// Known limitation (bis-ranker-integrity-11 lane report): this gate
// only ever catches a candidate's gain being pure SAMPLING noise
// around a true value at or near zero - it cannot catch a gain that
// is itself real and reproducible by the engine's own math for a
// reason unrelated to the trinket's own stated effect. Dogfooded
// directly: hunter-beast-mastery band 60 (both factions) still
// measures Fire Ruby/Burst of Knowledge/Second Wind/Sanctified Orb
// each gaining 1-7 DPS over the no-trinket baseline, and the gain does
// NOT shrink toward zero as trinketRankIterations rises (100 to 2000,
// a 20x rerun of this exact repro) - only gainStdErr shrinks, exactly
// as it would for a REAL, non-noise effect, so no iteration count or
// significance multiplier this lane could choose would ever gate it.
// Every simmable candidate in that slot's own pool measures a
// similarly-shaped positive gain regardless of whether its own effect
// text has anything to do with this spec (Sanctified Orb's mana
// restore and Frozen Heart of the Mountain's own +9 Hit rating measure
// the same order of magnitude there) - the likely cause is the spec's
// own APL's unconditional `autocastOtherCooldowns` action (this
// spec's own data/curated/apl/hunter-beast-mastery.json) crediting ANY
// equipped on-use item with some economy-of-action value the moment
// it is off cooldown, independent of what the item's own effect
// actually does - an engine-side (sim/core's item-effect/APL
// interaction) question, not a ranker one, and the engine fork is
// read-only for this lane. Flagged for the controller as a follow-up:
// either the engine's own autocastOtherCooldowns handling needs to
// stop crediting a no-op on-use effect, or trinketShortlist needs a
// genuinely spec-aware relevance check (does this item's own effect
// reference a resource/spell this spec's class actually has) rather
// than a DPS-measurement significance test, since the measurement
// here is not wrong, only the premise that a positive measurement
// implies real value is.

// fallbackWeightsReason is the weights_reason main.go publishes when a
// band's own sweep - even re-run once at weightsRetryIterationsFactor
// iterations (main.go's own guard) - still fails
// referenceMeasurementTrustworthy: this band's picks are ranked and
// verified against fromBand's own last-trusted weights instead of an
// empty map (an empty slot must only ever mean no candidate exists,
// never that a sweep was noisy - this lane's brief), and the report
// says exactly that rather than a bare "not significant".
func fallbackWeightsReason(referenceStat string, wresult map[string]api.StatWeight, referenceDPSPerPoint float64, band, fromBand int) string {
	_, rawStderr := referenceMeasurementTrustworthy(referenceStat, wresult, referenceDPSPerPoint)
	return fmt.Sprintf(
		"weights carried from band %d: band %d's own sweep measured %.4f ± %.4f DPS per point at this band's gear - not positive beyond its own error",
		fromBand, band, referenceDPSPerPoint, rawStderr,
	)
}

func effectiveWeights(weights map[string]api.StatWeight) map[string]float64 {
	out := make(map[string]float64, len(weights))
	for stat, w := range weights {
		if w.Weight > 0 {
			out[stat] = w.Weight
		}
	}
	return out
}
