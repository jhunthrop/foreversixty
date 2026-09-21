package bulk

// Ranking: what a finished stage means, and what runs next.
//
// Everything statistical about a bulk run is here, and nowhere else.
// The page does not recompute a delta, a standard error or a tie; it
// renders what this returns. That is what stops the browser and the
// server from disagreeing about who won.

import (
	"errors"
	"fmt"
	"math"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// ErrStageMismatch is returned when the results do not line up with the
// stage they claim to answer. Scoring them anyway would attribute one
// combination's DPS to another's substitutions.
var ErrStageMismatch = errors.New("bulk: the results do not match the stage")

// scored is one combination with this stage's numbers.
type scored struct {
	Combo Combination
	DPS   api.Estimate
	Delta api.Estimate
}

// Rank scores a finished stage and returns either the next stage or,
// after the last rung, the finished result. Exactly one of next and
// final is non-nil.
func Rank(req api.SimRequest, stage StageRequests, results []api.SimResult) (*StageRequests, *api.SimResult, error) {
	if req.Bulk == nil {
		return nil, nil, ErrNotBulk
	}
	ladder, ok := api.Ladders[req.Bulk.Precision]
	if !ok {
		return nil, nil, fmt.Errorf("bulk: no ladder for precision %q", req.Bulk.Precision)
	}
	ranked, equipped, err := score(stage, results)
	if err != nil {
		return nil, nil, err
	}
	if stage.Stage < 1 || stage.Stage > len(ladder.Iterations) {
		return nil, nil, fmt.Errorf("%w: stage %d, and the %s ladder has %d", ErrStageMismatch, stage.Stage, req.Bulk.Precision, len(ladder.Iterations))
	}
	if stage.Stage == len(ladder.Iterations) {
		// stage.Ran is every stage BEFORE this one; this last stage
		// itself was never appended (Rank only appends when building a
		// NEXT stage, below, and there is no next one here). Task 17's
		// finalResult needs stage.Ran PLUS one entry for this stage
		// (stage.Iterations, len(stage.Combos)) to fill SimResult.Stages
		// completely.
		final := finalResult(req, stage, ranked, equipped, results[0])
		return nil, &final, nil
	}
	kept := ranked
	if req.Bulk.Mode != api.KindTalents {
		// dps-minmaxer round 3, finding 5: talents mode submits exactly
		// the loadouts the player ticked -- a handful of builds chosen
		// on purpose, never the hundreds of auto-generated combinations
		// Top Gear's own cut exists to narrow. Fast precision's first
		// cut keeps a quarter of the field; at two loadouts that rounds
		// to one, so a comparison build that is simply worse (not
		// noise) never survives to the final result, and the page
		// cannot rank a build it was never told the answer for. Every
		// other mode keeps the cut unchanged.
		kept = applyCut(ranked, ladder.Cuts[stage.Stage-1])
	}
	combos := make([]Combination, 0, len(kept))
	for _, s := range kept {
		combos = append(combos, s.Combo)
	}
	// The documented protocol (plan.go's package doc, contract A10):
	// Rank appends the stage it just ran onto the Ran it was handed,
	// then threads that forward as the next stageRequests call's
	// history, so the LAST StageRequests either side ever receives
	// carries the whole ladder run so far.
	ran := append(slices.Clone(stage.Ran), api.Stage{Iterations: stage.Iterations, Combos: len(stage.Combos)})
	next := stageRequests(req, stage.Stage+1, ladder.Iterations[stage.Stage], combos, ran)
	return &next, nil, nil
}

// score pairs each result with its combination and computes the delta
// against the equipped set.
//
// The delta's error is the two errors added in quadrature. The runs
// share a seed, so they are positively correlated and the true error of
// the difference is SMALLER than this; the envelope carries no
// covariance, so the conservative figure is the one reported. Saying a
// gain is "within error" when it is real costs a player nothing; the
// reverse costs them an upgrade.
func score(stage StageRequests, results []api.SimResult) ([]scored, api.Estimate, error) {
	if len(results) != len(stage.Requests) {
		return nil, api.Estimate{}, fmt.Errorf("%w: %d results for %d requests", ErrStageMismatch, len(results), len(stage.Requests))
	}
	if len(stage.Combos) != len(stage.Requests)-1 {
		return nil, api.Estimate{}, fmt.Errorf("%w: %d combinations and %d requests; the equipped set is the extra one", ErrStageMismatch, len(stage.Combos), len(stage.Requests))
	}
	for i, res := range results {
		if res.Error != "" {
			return nil, api.Estimate{}, fmt.Errorf("bulk: run %d of stage %d failed: %s", i, stage.Stage, res.Error)
		}
		if res.Aborted {
			return nil, api.Estimate{}, fmt.Errorf("bulk: run %d of stage %d was stopped before it finished", i, stage.Stage)
		}
		// A result the right SIZE but the wrong ORDER would otherwise
		// pass every check above and silently attribute one
		// combination's DPS to another's substitutions - exactly what
		// this error's own doc comment says never happens. Stage
		// requests run through a worker pool with no ordering
		// guarantee of their own (spec 10.2); this is the only place
		// that checks results stayed lined up with the requests that
		// produced them.
		if res.IterationsRun != stage.Iterations {
			return nil, api.Estimate{}, fmt.Errorf("%w: result %d ran %d iterations, stage %d runs %d", ErrStageMismatch, i, res.IterationsRun, stage.Stage, stage.Iterations)
		}
		if err := sameSubstitutedFields(res.Request, stage.Requests[i]); err != nil {
			return nil, api.Estimate{}, fmt.Errorf("%w: result %d does not carry stage request %d's %v; results must stay in request order", ErrStageMismatch, i, i, err)
		}
		if math.IsNaN(res.DPS.Mean) || math.IsInf(res.DPS.Mean, 0) || math.IsNaN(res.DPS.Error) || math.IsInf(res.DPS.Error, 0) {
			return nil, api.Estimate{}, fmt.Errorf("bulk: run %d of stage %d reported a non-finite DPS (mean %v, error %v); a NaN mean would sort as tied with everything and a NaN bar would make the cut keep everyone silently", i, stage.Stage, res.DPS.Mean, res.DPS.Error)
		}
	}
	equipped := results[0].DPS
	// ranked pairs a scored combination with its precomputed sort key,
	// so the key - which walks every substitution's chip - is built
	// once per combination rather than twice per comparison inside the
	// sort (O(n) chipKey calls instead of O(n log n) of them, which
	// matters at the server's 5,000-combination cap).
	type ranked struct {
		scored
		key string
	}
	tmp := make([]ranked, 0, len(stage.Combos))
	for i, combo := range stage.Combos {
		dps := results[i+1].DPS
		tmp = append(tmp, ranked{
			scored: scored{
				Combo: combo,
				DPS:   dps,
				Delta: api.Estimate{
					Mean:   dps.Mean - equipped.Mean,
					StdDev: math.Hypot(dps.StdDev, equipped.StdDev),
					Error:  math.Hypot(dps.Error, equipped.Error),
				},
			},
			key: chipKey(combo),
		})
	}
	// Best first, and ties broken by the substitution chip so two runs
	// of one request produce the same order.
	slices.SortStableFunc(tmp, func(a, b ranked) int {
		switch {
		case a.DPS.Mean > b.DPS.Mean:
			return -1
		case a.DPS.Mean < b.DPS.Mean:
			return 1
		case a.key < b.key:
			return -1
		case a.key > b.key:
			return 1
		default:
			return 0
		}
	})
	out := make([]scored, len(tmp))
	for i, r := range tmp {
		out[i] = r.scored
	}
	return out, equipped, nil
}

// sameSubstitutedFields reports which field of a result's own request
// disagrees with the stage request it claims to answer, or nil if none
// does. It is score's order guard, and it must cover EVERY field
// bulk's apply can change, not just gear.
//
// Gear alone leaves the guard inert for a whole mode: in talents mode
// apply never touches gear, so every combination AND the equipped
// baseline carry byte-identical gear lists and any permutation of the
// results passes. A three-loadout talents plan fed results permuted
// [0,3,2,1] was accepted and ranked exactly backwards. The same hole
// swallows any two gear-mode combinations that differ only in their
// talent loadout or their consumable list - the other two dimensions
// of the product - which is every request with a Consumables or
// Talents dimension and no gear candidates in play.
//
// The fields are exactly apply's writes: Character.Gear,
// Character.Talents (a talent loadout) and Character.Consumes (an
// alternative consumable list). A substitution that changed anything
// else would need a line here too, or this guard goes quiet for it in
// the same way.
func sameSubstitutedFields(got, want api.SimRequest) error {
	switch {
	case !slices.Equal(got.Character.Gear, want.Character.Gear):
		return errors.New("gear")
	case got.Character.Talents != want.Character.Talents:
		return errors.New("talents")
	case !slices.Equal(got.Character.Consumes, want.Character.Consumes):
		return errors.New("consumables")
	}
	return nil
}

// chipKey is a stable ordering for two combinations of equal DPS.
func chipKey(c Combination) string {
	var b []byte
	for _, s := range c.Substitutions {
		b = fmt.Appendf(b, "%s|%s|%d|%d|%d|%s;", s.Kind, s.Slot, s.ItemID, s.Enchant, s.Suffix, s.Name)
	}
	return string(b)
}

// applyCut keeps the cut's survivors plus anything still overlapping
// the last of them.
//
// The slack is what makes a 100-iteration stage safe. At that count a
// combination's error is large, so the order is mostly noise; cutting
// strictly to the top quarter would throw away the eventual winner
// roughly as often as not. Keeping anything whose interval still
// reaches the cut's costs a few more sims and cannot lose a winner to
// one unlucky stage.
func applyCut(ranked []scored, cut api.Cut) []scored {
	if len(ranked) == 0 {
		return ranked
	}
	keep := cut.Top
	if cut.Fraction > 0 {
		keep = int(math.Round(float64(len(ranked)) * cut.Fraction))
	}
	keep = max(keep, 1)
	if keep >= len(ranked) {
		return ranked
	}
	// The last of the cut's own survivors sets the bar; every
	// candidate past it is checked against that FIXED bar
	// independently, per api.Cut's own doc: "keeping ANYTHING whose
	// interval still overlaps the last survivor's" - not "a
	// contiguous run starting there". Error varies candidate to
	// candidate, so the overlap predicate is not monotone in
	// mean-sorted order: a tight, non-overlapping candidate can sit
	// ranked ABOVE a wide, overlapping one. Stopping at (or including
	// up to) the first failure would either drop or wrongly keep a
	// candidate the slack exists to protect, so this is a filter over
	// the whole tail, not a shortened or widened prefix.
	bar := ranked[keep-1].DPS.Mean - cut.SlackSE*ranked[keep-1].DPS.Error
	out := append([]scored(nil), ranked[:keep]...)
	for _, s := range ranked[keep:] {
		if s.DPS.Mean+cut.SlackSE*s.DPS.Error >= bar {
			out = append(out, s)
		}
	}
	return out
}

// finalResult turns the last stage into the answer.
//
// The headline DPS and the summary are the EQUIPPED set's: a Top Gear
// report shows the character the player has, with the ranking beside
// it, and a headline taken from the winner would tell them they
// already do 1,100 DPS.
//
// Two fields of that equipped run are dropped rather than inherited,
// because they describe the ONE SIM it was and not the ladder this
// result is about. See the assignments below.
func finalResult(req api.SimRequest, stage StageRequests, ranked []scored, equipped api.Estimate, base api.SimResult) api.SimResult {
	out := base
	out.Request = req
	out.DPS = equipped
	out.Equipped = &equipped
	out.IterationsRun = stage.Iterations
	// A bulk result carries no cast log. Every stage request sets
	// NoSample (plan.go's stamp) precisely so the engine never builds
	// one, but that is the planner asking nicely: a caller that
	// assembled its own stage requests, or an engine that filled the
	// sample anyway, would have a 3,000-iteration median cast log
	// copied out of the equipped run into every Top Gear result - a
	// sample of one combination presented as the run's. Contract 10.3
	// says a stage sim's cast log is never read; this is the half of
	// that which does not depend on anyone else's cooperation.
	out.Sample = nil
	// The equipped run's own wall time is not the ladder's. base is
	// the LAST stage's equipped sim - one of dozens - so inheriting
	// its DurationMS reports a fraction of the run as the whole of it.
	//
	// Rank cannot measure the right number either: it is called once
	// per stage, holds no state between calls, and on the browser side
	// each call is a separate trip across the wasm boundary with the
	// page driving the loop. So it reports nothing, and the lane that
	// drove the ladder stamps the wall time it measured - the same way
	// each lane stamps Lane and EngineVersion, which Rank equally
	// cannot know (sim/cmd/forever-sim's executeBulk does exactly
	// this). Zero here means "not measured", and both lanes now get
	// zero from the shared code instead of one of them getting a
	// misleading number.
	//
	// The alternative considered and rejected: carry a duration on
	// api.Stage and sum the ladder's stages here. That would give both
	// lanes one identical, lane-independent number, but api.Stage's
	// shape is contract 10 JSON crossing the wasm boundary, and
	// widening it is an envelope change rather than a fix.
	out.DurationMS = 0
	out.Combos = make([]api.Combo, 0, len(ranked))
	for _, s := range ranked {
		out.Combos = append(out.Combos, api.Combo{
			Substitutions: s.Combo.Substitutions,
			DPS:           s.DPS,
			Delta:         s.Delta,
		})
	}
	group(out.Combos)
	out.Stages = stagesOf(stage)
	return out
}

// group numbers the within-error bands.
//
// A band starts at the first combination not yet in one, and every
// following combination whose delta interval still overlaps THAT
// leader's joins it. The intervals are one standard error either side,
// which is the interval the page draws.
//
// Comparing each row to its group's leader rather than to its
// predecessor is deliberate: chaining would walk a long tail of
// overlapping neighbours into one band whose ends do not overlap at
// all, and the page would rank a real 40-DPS gap as a tie.
//
// The comparison is between DELTA intervals, not DPS intervals, so
// the equipped set's own error enters both sides of every combo-vs-
// combo comparison and makes a tie slightly likelier than a direct
// pairwise test of the two candidates' DPS would. That is deliberate
// and consistent with score's stated conservatism (see its doc
// comment): a later reader should not "fix" it by comparing DPS
// directly.
func group(combos []api.Combo) {
	current := -1
	var leader api.Estimate
	for i := range combos {
		d := combos[i].Delta
		if current < 0 || d.Mean+d.Error < leader.Mean-leader.Error {
			current++
			leader = d
		}
		combos[i].Group = current
	}
}

// stagesOf is what the ladder actually ran: the stages the request
// carried through the loop (contract A10), plus the one that just
// finished. Rank is called once per stage and holds no state between
// calls, so stage.Ran - threaded across the wasm boundary inside
// StageRequests - is the only place that history can live.
func stagesOf(stage StageRequests) []api.Stage {
	return append(append([]api.Stage(nil), stage.Ran...),
		api.Stage{Iterations: stage.Iterations, Combos: len(stage.Combos)})
}
