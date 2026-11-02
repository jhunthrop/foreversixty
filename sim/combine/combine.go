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
	"github.com/jhunthrop/foreversixty/sim/adapter"
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

// poolEstimates pools one estimate read from every part, weighting each by
// the iterations behind it.
//
// The pooled variance is the within-part variance plus the spread between
// the part means. Dropping the second term would report a tighter error
// than a serial run, and the sim page would lie about its own precision.
// Min and Max are extremes over the parts; a part whose Max is zero
// reported no distribution at all, and folding its Min in would pull the
// run's minimum to zero - a figure no iteration produced.
func poolEstimates(parts []api.SimResult, read func(api.SimResult) api.Estimate) api.Estimate {
	var total int
	var mean float64
	for _, p := range parts {
		total += p.IterationsRun
		mean += read(p).Mean * float64(p.IterationsRun)
	}
	mean /= float64(total)

	var pooled float64
	for _, p := range parts {
		e := read(p)
		d := e.Mean - mean
		pooled += float64(p.IterationsRun) * (e.StdDev*e.StdDev + d*d)
	}
	pooled /= float64(total)

	out := api.Estimate{
		Mean:   mean,
		StdDev: math.Sqrt(pooled),
		Error:  math.Sqrt(pooled) / math.Sqrt(float64(total)),
	}
	first := true
	for _, p := range parts {
		e := read(p)
		if e.Max == 0 {
			continue
		}
		if first {
			out.Min, out.Max = e.Min, e.Max
			first = false
			continue
		}
		out.Min = math.Min(out.Min, e.Min)
		out.Max = math.Max(out.Max, e.Max)
	}
	return out
}

// Results combines partial results into one.
//
// DPS (and a role's healing and tank blocks) is pooled, IterationsRun summed, DurationMS the slowest part's, and
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

	out.DPS = poolEstimates(parts, func(p api.SimResult) api.Estimate { return p.DPS })
	var err error
	if out.Healing, out.Tank, err = poolRoles(parts); err != nil {
		return api.SimResult{}, err
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
	// Summary.DurationMS is the fight clock every per-fight figure in the
	// summary is divided by - the table's per-second totals, an aura's
	// uptime share, an actor's activity percent - and it is a different
	// field from out.DurationMS above (the wall clock this whole combine
	// took). weightSummaries just merged every part's damage table into
	// one, and out.DPS above is the iteration-weighted pooled mean of
	// every part's own DPS - but out.Summary.DurationMS, copied in with
	// the rest of parts[0].Summary, is still part 0's OWN duration: the
	// one number that made part 0's own table agree with part 0's own
	// headline. Two shards of one run almost never share a mean DPS or a
	// fight length (different iterations, different encounter
	// variation), so leaving it at part 0's reopens exactly the
	// disagreement DeriveDurationMS exists to close, once combine is the
	// thing doing the dividing.
	partZeroDurationMS := out.Summary.DurationMS
	out.Summary.DurationMS = adapter.DeriveDurationMS(
		adapter.SumActorTotals(out.Summary.DamageDone), out.DPS.Mean, fallbackDurationMS(parts, total),
	)
	// The player's ActiveMS is the fight's own duration, and a measured
	// pet's is its share of the fight, the same rules sim/adapter.Summarize
	// keeps. Both are restated here because weightSummaries carries each
	// merged actor's ActiveMS in from whichever part first introduced that
	// GUID, and that part's own duration is no longer the summary's.
	shares := petActiveShares(parts, total)
	for i, row := range out.Summary.DamageDone {
		active := out.Summary.DurationMS
		if row.PetUptime {
			active = int64(math.Round(shares[row.GUID] * float64(out.Summary.DurationMS)))
		}
		out.Summary.DamageDone[i].ActiveMS = active
	}
	// Auras, casts and resources are still part 0's own - weightSummaries'
	// own doc comment calls that a "share the largest part already
	// represents within sampling error" - but the clock they were
	// measured against just moved. An aura's UptimeMS is a MILLISECOND
	// COUNT, measured over part 0's OWN mean iteration length; dividing
	// it by the newly-derived (and, ordinarily, shorter) combined
	// duration inflates its uptime share past what part 0 itself ever
	// measured - which is how every full-uptime buff read an impossible
	// 100.9% (2026-09-21 result-page review round 3, E8). The fix
	// rescales UptimeMS by the same ratio the clock itself just moved
	// by, which preserves the one thing that is actually meaningful
	// across a duration change - the UPTIME SHARE part 0 measured, not
	// the millisecond count - so an aura that was up 100% of part 0's
	// own fight still reads exactly 100%, and one at 61% still reads 61%.
	rescaleToDuration(out.Summary.Auras, partZeroDurationMS, out.Summary.DurationMS)
	out.Request.Iterations = total
	return out, nil
}

// fallbackDurationMS is DeriveDurationMS's degenerate-case answer for a
// COMBINED summary: the iteration-share-weighted average of every
// part's own duration, the same "weigh by iteration share" rule
// weightSummaries applies to every other per-fight figure it merges.
// DeriveDurationMS only reaches for it when the merged table has no
// damage or the pooled mean is non-positive - nothing happened, or the
// rate has no well-defined duration - and a blended measured duration is
// the sane answer either way, the same role sim/adapter's own
// AvgIterationDuration fallback plays for one part.
// petActiveShares is each measured pet row's share of the fight it was
// summoned for, weighted by each part's iterations: a pet out for all of
// one shard and half of an equal other was out for three quarters.
func petActiveShares(parts []api.SimResult, total int) map[string]float64 {
	shares := map[string]float64{}
	for _, p := range parts {
		if p.Summary.DurationMS <= 0 {
			continue
		}
		w := float64(p.IterationsRun) / float64(total)
		for _, row := range p.Summary.DamageDone {
			if row.PetUptime {
				share := min(float64(row.ActiveMS)/float64(p.Summary.DurationMS), 1)
				shares[row.GUID] += w * share
			}
		}
	}
	return shares
}

func fallbackDurationMS(parts []api.SimResult, total int) int64 {
	var weighted float64
	for _, p := range parts {
		weighted += float64(p.Summary.DurationMS) * float64(p.IterationsRun) / float64(total)
	}
	return int64(math.Round(weighted))
}

// rescaleToDuration rescales every aura's UptimeMS in place by newMS/oldMS -
// the ratio Summary.DurationMS just moved by - so the UPTIME SHARE
// (UptimeMS/DurationMS) part 0 measured survives the clock changing under
// it, rather than the millisecond count (2026-09-21 result-page review
// round 3, E8; see Results' own comment on the call site for the fuller
// account). Rescaling preserves the share exactly for an always-up aura
// (UptimeMS == oldMS becomes exactly newMS) and proportionally for every
// other one.
//
// oldMS <= 0 leaves every row untouched - part 0 itself had no well-defined
// duration to measure a share against, the same degenerate case
// DeriveDurationMS's own fallback exists for, and guessing a scale from
// nothing would invent a number rather than report one. The result is
// still clamped to newMS: math.Round on a scale computed from two
// independently-derived integers can round an always-up aura's UptimeMS
// one millisecond past the very duration it was just rescaled to agree
// with, and an uptime a hair over 100% is the exact defect this function
// exists to close.
func rescaleToDuration(auras []summary.AuraTrack, oldMS, newMS int64) {
	if oldMS <= 0 {
		return
	}
	scale := float64(newMS) / float64(oldMS)
	for i := range auras {
		rescaled := int64(math.Round(float64(auras[i].UptimeMS) * scale))
		if rescaled > newMS {
			rescaled = newMS
		}
		auras[i].UptimeMS = rescaled
	}
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
	sums := map[string]map[abilityKey]*rowSums{}
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
				sums[src.GUID] = map[abilityKey]*rowSums{}
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
					shell.Min, shell.Max = ab.Min, ab.Max
					shell.Misses = nil
					rows[src.GUID][k] = &shell
					sums[src.GUID][k] = &rowSums{}
					rowOrder[src.GUID] = append(rowOrder[src.GUID], k)
					cur = &shell
				}
				sums[src.GUID][k].add(ab, w)
				// Min and Max are extremes, not averages: weighting them
				// would report a spread no iteration ever saw.
				cur.Min = min(cur.Min, ab.Min)
				cur.Max = max(cur.Max, ab.Max)
			}
		}
	}

	merged := make([]summary.Actor, 0, len(order))
	for _, guid := range order {
		a := *actors[guid]
		a.Abilities = make([]summary.Ability, 0, len(rowOrder[guid]))
		a.Total, a.Effective = 0, 0
		for _, k := range rowOrder[guid] {
			ab := sums[guid][k].rounded(*rows[guid][k])
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

// rowSums is one merged row's weighed counters, kept unrounded until every
// part is in. Rounding each part's share first lost a count a part saw less
// than twice per fight: eight parts with 3 Ignite crits a fight each rounded
// 3/8 to 0 and the row read a 0.0% crit rate (round 7, page lane).
type rowSums struct {
	total, effective, overheal, overkill, absorbed float64
	resisted, blocked, hits, crits, ticks          float64
	misses                                         map[string]float64
}

// add weighs one part's row into the sums.
func (s *rowSums) add(ab summary.Ability, w float64) {
	s.total += float64(ab.Total) * w
	s.effective += float64(ab.Effective) * w
	s.overheal += float64(ab.Overheal) * w
	s.overkill += float64(ab.Overkill) * w
	s.absorbed += float64(ab.Absorbed) * w
	s.resisted += float64(ab.Resisted) * w
	s.blocked += float64(ab.Blocked) * w
	s.hits += float64(ab.Hits) * w
	s.crits += float64(ab.Crits) * w
	s.ticks += float64(ab.Ticks) * w
	for kind, n := range ab.Misses {
		if s.misses == nil {
			s.misses = map[string]float64{}
		}
		s.misses[kind] += float64(n) * w
	}
}

// rounded is row with the summed counters, each rounded once.
func (s *rowSums) rounded(row summary.Ability) summary.Ability {
	r := func(v float64) int64 { return int64(math.Round(v)) }
	row.Total, row.Effective = r(s.total), r(s.effective)
	row.Overheal, row.Overkill, row.Absorbed = r(s.overheal), r(s.overkill), r(s.absorbed)
	row.Resisted, row.Blocked = r(s.resisted), r(s.blocked)
	row.Hits, row.Crits, row.Ticks = r(s.hits), r(s.crits), r(s.ticks)
	row.Misses = nil
	for kind, n := range s.misses {
		if r(n) == 0 {
			continue // "0 dodge" is not something a row says (adapter addMiss)
		}
		if row.Misses == nil {
			row.Misses = make(map[string]int64, len(s.misses))
		}
		row.Misses[kind] = r(n)
	}
	return row
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
