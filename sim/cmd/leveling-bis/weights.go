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
	ref, ok := wresult[referenceStat]
	if !ok {
		// Every real request's WeightStats carries its own
		// ReferenceStat (weightsRequest's own doc), so this is
		// unreachable in production; leave the band alone rather than
		// invent a reason for a case that cannot happen with real
		// data.
		return ""
	}
	rawStderr := math.Abs(ref.Error * referenceDPSPerPoint)
	if referenceDPSPerPoint > rawStderr {
		return ""
	}
	return fmt.Sprintf(
		"reference stat %s measured %.4f ± %.4f DPS per point at this band's gear - not positive beyond its own error, so no weight this band measured can be trusted",
		referenceStat, referenceDPSPerPoint, rawStderr,
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
