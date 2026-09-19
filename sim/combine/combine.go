// Package combine splits a run across workers and puts the pieces back
// together. The browser's worker pool and the server lane both use it,
// so the arithmetic is written once.
package combine

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
)

// ErrMixedParts is returned when the parts are not shares of one run.
// Pooling them would print one authoritative DPS for two different
// characters, two different fights or two different engines, and
// nothing downstream could tell.
var ErrMixedParts = errors.New("combine: the parts are not shares of one run")

// Split divides a request into n parts by iteration count.
//
// Each part's RandomSeed is offset by the iterations of every part
// before it. That is what the engine's own SplitSimRequestForConcurrency
// does, and the reason is that the engine increments its seed once per
// iteration: without the offset, four workers would run the same four
// thousand rolls and a split run would not match a serial one. A paired
// comparison depends on it.
//
// A part's iteration count is not one of api.ValidIterations - 3,000
// over four workers is 750 - so a part is built with
// request.Options{SplitPart: true}, which validates everything but the
// closed set the settings bar offers.
//
// Only part zero is allowed to produce the sample iteration; every
// later part is opted out. A sample costs a whole replayed iteration
// on a fresh Environment, so leaving it on would make an N-way split
// pay N replays and throw N-1 of them away - and combine.Results keeps
// part zero's, so the discarded ones were never going to be read. A
// request that had already opted out stays opted out: an N-way split
// of a stage request must not put the sample back.
func Split(req api.SimRequest, n int) ([]api.SimRequest, error) {
	if n <= 0 {
		return nil, fmt.Errorf("combine: split count must be positive, got %d", n)
	}
	if req.Iterations <= 0 {
		return nil, fmt.Errorf("combine: iterations must be positive, got %d", req.Iterations)
	}
	if n > req.Iterations {
		n = req.Iterations
	}

	per := req.Iterations / n
	out := make([]api.SimRequest, n)
	seed := req.RandomSeed
	for i := 0; i < n; i++ {
		part := req
		part.Iterations = per
		if i == 0 {
			// The remainder goes to the first part, as the engine does
			// it; spreading it would lose iterations to truncation.
			part.Iterations += req.Iterations % n
		}
		part.NoSample = req.NoSample || i != 0
		part.RandomSeed = seed
		seed += int64(part.Iterations)
		out[i] = part
	}
	return out, nil
}

// Results combines partial results into one.
//
// DPS is pooled, IterationsRun summed, DurationMS the slowest part's, and
// the damage table merged row by row by weightSummaries. Everything else
// is the FIRST part's, presented as the whole run's: the request (whose
// RandomSeed is the run's, because combine.Split gives part zero the
// original), the summary's aura, cast, resource and threat tables,
// which are shares and averages the largest part already represents
// within sampling error, and Sample, which is part zero's one recorded
// fight - combine.Split asks no other part for one, so there is exactly
// one to take. Weighting the damage table is the part worth doing
// exactly, because it is the one a viewer reads as a number.
func Results(parts []api.SimResult) (api.SimResult, error) {
	if len(parts) == 0 {
		return api.SimResult{}, errors.New("combine: no results")
	}
	var total int
	for i, p := range parts {
		if p.Error != "" {
			return api.SimResult{}, fmt.Errorf("combine: part %d failed: %s", i, p.Error)
		}
		if p.Aborted {
			return api.SimResult{}, fmt.Errorf("combine: part %d was stopped before it finished", i)
		}
		if p.IterationsRun <= 0 {
			return api.SimResult{}, fmt.Errorf("combine: part %d ran no iterations", i)
		}
		if err := sameRun(parts[0], p); err != nil {
			return api.SimResult{}, fmt.Errorf("%w: part %d %v", ErrMixedParts, i, err)
		}
		total += p.IterationsRun
	}

	out := parts[0]
	out.IterationsRun = total

	// Pooled mean: weight each part by the iterations behind it.
	var mean float64
	for _, p := range parts {
		mean += p.DPS.Mean * float64(p.IterationsRun)
	}
	mean /= float64(total)

	// Pooled variance is the within-part variance plus the spread
	// between the part means. Dropping the second term would report a
	// tighter error than a serial run, and the sim page would lie about
	// its own precision.
	var pooled float64
	for _, p := range parts {
		w := float64(p.IterationsRun)
		d := p.DPS.Mean - mean
		pooled += w * (p.DPS.StdDev*p.DPS.StdDev + d*d)
	}
	pooled /= float64(total)

	out.DPS = api.Estimate{
		Mean:   mean,
		StdDev: math.Sqrt(pooled),
		Error:  math.Sqrt(pooled) / math.Sqrt(float64(total)),
	}
	// Min and Max are extremes over the parts. A part whose Max is zero
	// reported no distribution at all, and folding its Min in would
	// pull the run's minimum to zero - a figure no iteration produced.
	first := true
	for _, p := range parts {
		if p.DPS.Max == 0 {
			continue
		}
		if first {
			out.DPS.Min, out.DPS.Max = p.DPS.Min, p.DPS.Max
			first = false
			continue
		}
		out.DPS.Min = math.Min(out.DPS.Min, p.DPS.Min)
		out.DPS.Max = math.Max(out.DPS.Max, p.DPS.Max)
	}

	out.DurationMS = 0
	for _, p := range parts {
		if p.DurationMS > out.DurationMS {
			// Wall clock of a parallel run is the slowest part, not the
			// sum: the parts ran at the same time.
			out.DurationMS = p.DurationMS
		}
	}

	out.Summary = weightSummaries(parts, total)
	out.Request.Iterations = total
	return out, nil
}

// weightSummaries averages the per-fight damage table by iteration
// share. Each part's summary is already a per-fight average (sim/adapter
// divides by IterationsDone), so combining them is a weighted mean of
// like quantities rather than a re-sum.
//
// Rows are merged by identity, never by position: an actor by GUID and
// an ability by the (SpellID, Via) pair the report keys its lists on.
// The adapter sorts a part's abilities by Total descending, so two parts
// of a split run order near-ties differently and a proc that fires in
// one part and not another changes the row count - merging by slice
// index would then add one spell's damage to another's and drop the odd
// row, silently. A row only a later part has is added rather than folded
// into whatever sat at its index.
//
// Nothing here touches a part. Every slice and map that is scaled is
// copied first, so Results can be called twice on the same slice and a
// caller that keeps its parts still has them.
func weightSummaries(parts []api.SimResult, total int) summary.Summary {
	out := parts[0].Summary
	if len(parts) == 1 {
		return out
	}

	// Actors keep the adapter's order - the player, then its pets - with
	// any actor only a later part saw appended. Abilities are re-sorted
	// on the way out, because their order is by damage and the damage is
	// what just changed.
	var order []string
	actors := map[string]*summary.Actor{}
	rows := map[string]map[abilityKey]*summary.Ability{}
	var rowOrder = map[string][]abilityKey{}

	for _, p := range parts {
		w := float64(p.IterationsRun) / float64(total)
		for _, src := range p.Summary.DamageDone {
			dst, seen := actors[src.GUID]
			if !seen {
				shell := src
				shell.Total, shell.Effective, shell.Overheal, shell.Absorbed = 0, 0, 0, 0
				shell.Abilities, shell.Targets, shell.Series = nil, nil, nil
				actors[src.GUID] = &shell
				rows[src.GUID] = map[abilityKey]*summary.Ability{}
				order = append(order, src.GUID)
				dst = &shell
			}
			// Total and Effective are NOT weighed here: they are
			// recomputed from the merged ability rows below, because
			// the adapter guarantees an actor's total is the sum of
			// its rows and weighing the two independently breaks that
			// by a rounding error per row.
			dst.Overheal += weigh(src.Overheal, w)
			dst.Absorbed += weigh(src.Absorbed, w)
			dst.Targets = mergeTargets(dst.Targets, src.Targets, w)
			dst.Series = mergeSeries(dst.Series, src.Series, w)

			for _, ab := range src.Abilities {
				k := abilityKey{SpellID: ab.SpellID, Via: ab.Via}
				cur, seen := rows[src.GUID][k]
				if !seen {
					shell := ab
					shell.Total, shell.Effective = 0, 0
					shell.Overheal, shell.Overkill, shell.Absorbed = 0, 0, 0
					shell.Resisted, shell.Blocked = 0, 0
					shell.Hits, shell.Crits, shell.Ticks = 0, 0, 0
					shell.Min, shell.Max = ab.Min, ab.Max
					shell.Misses = nil
					rows[src.GUID][k] = &shell
					rowOrder[src.GUID] = append(rowOrder[src.GUID], k)
					cur = &shell
				}
				cur.Total += weigh(ab.Total, w)
				cur.Effective += weigh(ab.Effective, w)
				cur.Overheal += weigh(ab.Overheal, w)
				cur.Overkill += weigh(ab.Overkill, w)
				cur.Absorbed += weigh(ab.Absorbed, w)
				cur.Resisted += weigh(ab.Resisted, w)
				cur.Blocked += weigh(ab.Blocked, w)
				cur.Hits += weigh(ab.Hits, w)
				cur.Crits += weigh(ab.Crits, w)
				cur.Ticks += weigh(ab.Ticks, w)
				// Min and Max are extremes, not averages: weighting them
				// would report a spread no iteration ever saw.
				cur.Min = min(cur.Min, ab.Min)
				cur.Max = max(cur.Max, ab.Max)
				cur.Misses = mergeMisses(cur.Misses, ab.Misses, w)
			}
		}
	}

	merged := make([]summary.Actor, 0, len(order))
	for _, guid := range order {
		a := *actors[guid]
		a.Abilities = make([]summary.Ability, 0, len(rowOrder[guid]))
		a.Total, a.Effective = 0, 0
		for _, k := range rowOrder[guid] {
			ab := *rows[guid][k]
			a.Total += ab.Total
			a.Effective += ab.Effective
			a.Abilities = append(a.Abilities, ab)
		}
		// The adapter's own order: damage descending. Ties break on the
		// row's identity rather than on where it happened to arrive, so
		// the combined table is the same table whatever order the
		// workers finished in.
		sort.Slice(a.Abilities, func(i, j int) bool {
			if a.Abilities[i].Total != a.Abilities[j].Total {
				return a.Abilities[i].Total > a.Abilities[j].Total
			}
			if a.Abilities[i].SpellID != a.Abilities[j].SpellID {
				return a.Abilities[i].SpellID < a.Abilities[j].SpellID
			}
			return a.Abilities[i].Via < a.Abilities[j].Via
		})
		merged = append(merged, a)
	}
	out.DamageDone = merged
	return out
}

// sameRun reports why two parts do not belong to one run, or nil.
//
// The three things that must match are the engine that ran them, the
// spec they ran, and the request itself apart from the two fields
// combine.Split is allowed to change: RandomSeed, which is offset per
// part so the streams do not repeat, and Iterations, which is the
// part's share. Everything else - gear, talents, buffs, encounter - is
// what the DPS is a number about.
func sameRun(first, p api.SimResult) error {
	if p.EngineVersion != first.EngineVersion {
		return fmt.Errorf("ran on engine %q, part 0 on %q", p.EngineVersion, first.EngineVersion)
	}
	if p.Request.Spec != first.Request.Spec {
		return fmt.Errorf("is spec %q, part 0 is %q", p.Request.Spec, first.Request.Spec)
	}
	if !reflect.DeepEqual(shape(p.Request), shape(first.Request)) {
		return errors.New("asks a different question from part 0")
	}
	return nil
}

// shape is a request with the fields a split or a Smart Sim step is
// allowed to vary cleared, so two parts of one run compare equal.
//
// TargetError joins the seed and the count because a step of a
// target-error run is a fixed-count part of it: the loop owns the
// target, the part does not.
//
// NoSample joins them because combine.Split sets it on every part but
// the first, and it changes nothing the DPS is a number about - it
// only says whether that part recorded one fight's casts on the way
// past. Without clearing it here, sameRun would reject the very parts
// Split produces with "asks a different question from part 0", and
// every split run - which is every browser run - would fail to
// combine.
func shape(r api.SimRequest) api.SimRequest {
	r.RandomSeed = 0
	r.Iterations = 0
	r.TargetError = 0
	r.NoSample = false
	return r
}

// abilityKey is the identity a summary row is rendered by: the spell and,
// for a pet's ability counted on its owner's row, which pet cast it.
type abilityKey struct {
	SpellID int64
	Via     string
}

// weigh takes an iteration-share of a per-fight figure. It ROUNDS: a
// truncation biases every row of a four-way split low by up to one per
// part, which on a table of forty rows is a visible shortfall against
// the same run done serially.
func weigh(v int64, w float64) int64 { return int64(math.Round(float64(v) * w)) }

// mergeTargets folds one part's per-target damage into the running table,
// keyed by the target's guid. The slice it returns is always the
// caller's own.
func mergeTargets(dst, src []summary.Pair, w float64) []summary.Pair {
	for _, p := range src {
		found := false
		for i := range dst {
			if dst[i].GUID == p.GUID {
				dst[i].Total += weigh(p.Total, w)
				found = true
				break
			}
		}
		if !found {
			p.Total = weigh(p.Total, w)
			dst = append(dst, p)
		}
	}
	return dst
}

// mergeSeries folds one part's per-second timeline into the running one.
// The index is a second, so the same second of two parts is the same
// quantity; a part that ran longer extends the series.
func mergeSeries(dst, src []int64, w float64) []int64 {
	for len(dst) < len(src) {
		dst = append(dst, 0)
	}
	for i, v := range src {
		dst[i] += weigh(v, w)
	}
	return dst
}

// mergeMisses folds one part's outcome counters into the running map,
// allocating rather than writing into the part's own.
func mergeMisses(dst, src map[string]int64, w float64) map[string]int64 {
	if len(src) == 0 {
		return dst
	}
	if dst == nil {
		dst = make(map[string]int64, len(src))
	}
	for k, v := range src {
		dst[k] += weigh(v, w)
	}
	return dst
}
