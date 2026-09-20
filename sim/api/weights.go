package api

// Stat weights.
//
// The engine already computes them, so this is a request shape and a
// result shape, not arithmetic. The caution belongs to the page: a
// weight is a linear guess at a non-linear thing, and simming the
// actual items is the better answer. The envelope's job is to make the
// request honest - a reference stat that is not among the stats being
// weighed would normalise against a number nobody computed.

import (
	"errors"
	"fmt"
	"slices"
)

// WeightsSpec asks for stat weights instead of a DPS number.
type WeightsSpec struct {
	// Stats are IDS.md stat ids: the engine's Stat enum names in lower
	// snake case with the prefix stripped, "attack_power", "crit",
	// "spell_haste". sim/request resolves them; an id the engine has no
	// stat for is refused there rather than weighed as nothing.
	Stats []string `json:"stats"`
	// Reference is the stat normalised to exactly 1.0. It is
	// REQUIRED (contract A7): its per-spec default lives in
	// data/curated/specs.json, reaches Go as the generated
	// sim/specs struct's ReferenceStat, and is served to the page on
	// GET /v1/specs. The envelope does not read it - defaulting here
	// would make the envelope a second source of truth for a fact
	// the data lane owns, and a request that did not say which stat
	// it normalised against would be unreproducible.
	Reference string `json:"reference"`
}

// StatWeight is one stat's answer. Reference's Weight is exactly 1.
type StatWeight struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
	// Error is the standard error of the mean over the weighted run's
	// own iterations (sim/adapter.Weights divides the engine's raw
	// standard deviation by sqrt(N), the same conversion adapter.DPS
	// makes for the headline number) - the uncertainty from THIS run
	// not having sampled forever. It is not the whole uncertainty: an
	// 8-widely-seeded empirical check of the warrior-fury sample
	// request (simfury-weights.json, "normal", WeightsIterationsFactor
	// 8) found the run-to-run spread of the weight itself exceeds this
	// figure by roughly 1.2x for a stat that scales damage smoothly
	// (strength) but by up to about 10x for a stat that gates the
	// rotation's own decisions (crit, expertise, melee_haste) - see
	// task-5-report.md for the table. Error is therefore a lower bound
	// on the true uncertainty for those stats; a correct estimate
	// needs a replication estimator (several independent sub-sweeps
	// averaged together), deliberately deferred rather than built into
	// this round.
	Error float64 `json:"error"`
	// Insignificant is true when this weight cannot be told apart from
	// zero: either the run's own error swamps the weight (Error >=
	// abs(Weight)), or the engine never moved the stat at all - a
	// hard-capped stat is skipped by the engine's sweep and its Weight
	// and Error both read back as the result array's zero value, which
	// satisfies the same >= test. The comparison is on the absolute
	// value, so a stat that genuinely hurts DPS (a real negative
	// weight bigger than its error) is not flagged.
	//
	// The page greys an insignificant row rather than presenting it
	// as a confident number; the Pawn string omits it rather than
	// exporting a stat weight of 0.00 that looks measured but is not.
	//
	// No omitempty: the web needs a literal `false` on the wire to
	// tell "significant" from "field absent" (an older client, a
	// fixture predating this field).
	Insignificant bool `json:"insignificant"`
}

// KnownStats is the closed vocabulary a weights request may name: the
// ids sim/request/IDS.md's Stats section publishes, which
// sim/internal/statid generates from the engine's own Stat enum
// (proto.Stat_name). sim/api may not import the engine's proto - no
// protobuf crosses a lane boundary - so this is a copy rather than a
// call across that boundary; sim/request's
// TestAPIKnownStatsMatchTheGeneratedVocabulary proves the copy has
// not drifted from the generator's own list. It is sorted so an error
// message that lists it reads the same way twice.
var KnownStats = []string{
	"agility", "arcane_power", "arcane_resistance", "armor",
	"armor_penetration", "attack_power", "block", "block_value",
	"bonus_armor", "crit", "defense", "dodge", "energy", "expertise",
	"feral_attack_power", "fire_power", "fire_resistance", "frost_power",
	"frost_resistance", "healing_power", "health", "hit", "holy_power",
	"intellect", "mana", "melee_haste", "mp5", "nature_power",
	"nature_resistance", "parry", "rage", "ranged_attack_power",
	"shadow_power", "shadow_resistance", "spell_damage", "spell_haste",
	"spell_penetration", "spell_power", "spirit", "stamina", "strength",
}

func (w *WeightsSpec) validate() []error {
	var errs []error
	if len(w.Stats) == 0 {
		errs = append(errs, errors.New("weights needs at least one stat to weigh"))
	}
	// Bounded before the per-id walk below: a request with more stats
	// than the vocabulary has cannot be legal no matter what the ids
	// are (there is nowhere for a 42nd distinct, known id to come
	// from), and refusing it here in one error is both the honest
	// answer and what keeps an oversized Stats array from being
	// walked id by id for nothing.
	if len(w.Stats) > len(KnownStats) {
		errs = append(errs, fmt.Errorf("weights.stats has %d entries; the vocabulary only has %d",
			len(w.Stats), len(KnownStats)))
		return errs
	}
	seen := map[string]bool{}
	for _, s := range w.Stats {
		if s == "" {
			errs = append(errs, errors.New("weights.stats carries an empty stat id"))
			continue
		}
		if seen[s] {
			errs = append(errs, fmt.Errorf("weights.stats has %q listed twice", s))
			continue
		}
		seen[s] = true
		if !slices.Contains(KnownStats, s) {
			errs = append(errs, fmt.Errorf("weights.stats has %q, which is not a known stat id; see sim/request/IDS.md", s))
		}
	}
	switch {
	case w.Reference == "":
		errs = append(errs, errors.New("weights.reference is required; the spec's default is in data/curated/specs.json"))
	case !seen[w.Reference]:
		errs = append(errs, fmt.Errorf("weights.reference is %q, which is not one of the stats it weighs (%v); the reference is normalised to 1 and there would be nothing to normalise", w.Reference, w.Stats))
	}
	return errs
}

// WeightsIterationsFactor multiplies a weights request's own
// Iterations before the sweep runs, on top of the halving the
// engine's own buildStatWeightRequests applies for low/high RNG
// parity.
//
// It exists for a reason iteration count does not usually need
// spelling out: sim/core/statweight.go's WeightsStdev
// (buildStatWeightRequests / computeStatWeights in the wowsims fork)
// is the *population* standard deviation of the per-iteration
// low/baseline and high/baseline deltas - sim/core/utils.go's
// aggregator computes sqrt(sumSq/n - mean*mean), with no /sqrt(n).
// That quantity converges to a constant as the sample count grows; it
// is not sampling noise iteration count can shrink on its own.
// sim/adapter.Weights converts it to the standard error the page and
// the Pawn string need - the same conversion adapter.DPS already
// makes for the headline number - by dividing by sqrt(N), and N is
// this same multiplied count. Exporting the factor once, rather than
// letting BuildWeights, WeightsIterations and adapter.Weights each
// carry their own copy, is what keeps the cost estimate, the actual
// run and the error conversion from disagreeing about how many
// samples a weight is really built from.
//
// A target of "error under 10% of weight" turned out to assume the
// printed error was the whole uncertainty; it is not (see
// StatWeight.Insignificant's comment and StatWeight.Error's), so this
// factor is no longer chosen against that number. It buys a *stable*
// weight - one that does not move much if the same request runs
// again - not a truthful error bar; the printed error stays a lower
// bound regardless of the factor. What actually bounds the factor is
// the job budget: api/internal/sims/run.go's checkWeightsSize costs a
// weights request through this same WeightsIterations, and the
// largest legal request (MaxIterations, every id in KnownStats) has
// to fit api/internal/sims.BulkBudget. 8 is the largest power of two
// that does; 16 does not (see
// TestTheMaxWeightsRequestFitsTheBudget/task-5-report.md for the
// numbers). At 8, an 8-widely-seeded run of the warrior-fury sample
// request (simfury-weights.json, "normal"=3000) puts most weighed
// stats' cross-seed spread within a small multiple of their weight -
// crit's spread was about 7% of its own weight - which is the
// standard this factor is actually held to.
const WeightsIterationsFactor = 8

// WeightsIterations is the engine's own stat-weights arithmetic
// (sim/core/statweight.go in the wowsims fork,
// buildStatWeightRequests and runStatWeights), expressed exactly
// rather than approximated:
//
// The request's own Iterations is multiplied by
// WeightsIterationsFactor (see its doc for why) and then halved once,
// for RNG parity between the "more of this stat" and "less of this
// stat" runs, and that halved count becomes the size of EVERY sim
// pass the sweep runs: one baseline pass, then one low pass and one
// high pass per distinct stat being weighed. The reference stat is
// always one of them (WeightsSpec.validate requires it), and a
// repeated id is refused there too, so the distinct count is exactly
// len(req.Weights.Stats). That makes the total iterations run:
//
//	(Iterations * WeightsIterationsFactor / 2) * (1 + 2*len(Stats))
//
// the baseline's one pass plus two passes per stat, every pass at
// half the multiplied count. It returns 0 for a request with no
// Weights block, the way LadderIterations has nothing to cost for a
// request with no Bulk block.
func WeightsIterations(req SimRequest) int {
	if req.Weights == nil {
		return 0
	}
	return (req.Iterations * WeightsIterationsFactor / 2) * (1 + 2*len(req.Weights.Stats))
}
