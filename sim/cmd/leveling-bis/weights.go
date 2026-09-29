package main

import (
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
// A weight isWeightSignificant calls insignificant is zeroed here
// rather than carried through at its raw (possibly negative) value.
// Night-bis-sanity's own finding is why: ~40 low-level cloth items
// (Evergreen Gloves, Featherbead Bracers, ...) across six caster specs
// published a NEGATIVE score because a weight the sweep itself could
// not distinguish from noise happened to land small and negative on
// one of the item's stats, and score() dotted it straight into the
// total. This command already tells the player which weights it
// trusts (report.go's Insignificant flag, isWeightSignificant above)
// -- this function is that same judgment applied BEFORE ranking, not
// just after publication: a weight nobody should trust should not
// move a ranking either.
//
// The raw weights -- Weight, Error and Insignificant exactly as the
// engine reported and this command judged them -- still reach the
// published JSON unchanged (report.go's buildReport reads the
// map[string]api.StatWeight the caller passes it directly, not this
// function's output): a player or Pawn-style consumer reading the
// aside still sees the real number and its error bar, only the
// RANKING stops trusting it.
func effectiveWeights(weights map[string]api.StatWeight) map[string]float64 {
	out := make(map[string]float64, len(weights))
	for stat, w := range weights {
		if isWeightSignificant(w) {
			out[stat] = w.Weight
		}
	}
	return out
}
