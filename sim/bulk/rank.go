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
	"sort"

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
		final := finalResult(req, stage, ranked, equipped, results[0])
		return nil, &final, nil
	}
	kept := applyCut(ranked, ladder.Cuts[stage.Stage-1])
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
	}
	equipped := results[0].DPS
	out := make([]scored, 0, len(stage.Combos))
	for i, combo := range stage.Combos {
		dps := results[i+1].DPS
		out = append(out, scored{
			Combo: combo,
			DPS:   dps,
			Delta: api.Estimate{
				Mean:   dps.Mean - equipped.Mean,
				StdDev: math.Hypot(dps.StdDev, equipped.StdDev),
				Error:  math.Hypot(dps.Error, equipped.Error),
			},
		})
	}
	// Best first, and ties broken by the substitution chip so two runs
	// of one request produce the same order.
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].DPS.Mean != out[j].DPS.Mean {
			return out[i].DPS.Mean > out[j].DPS.Mean
		}
		return chipKey(out[i].Combo) < chipKey(out[j].Combo)
	})
	return out, equipped, nil
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
	// The last survivor's lower bound is the bar; anything whose upper
	// bound still reaches it is a tie with the cut.
	bar := ranked[keep-1].DPS.Mean - cut.SlackSE*ranked[keep-1].DPS.Error
	for keep < len(ranked) {
		s := ranked[keep]
		if s.DPS.Mean+cut.SlackSE*s.DPS.Error < bar {
			break
		}
		keep++
	}
	return ranked[:keep]
}

// finalResult is Task 17.
func finalResult(req api.SimRequest, stage StageRequests, ranked []scored, equipped api.Estimate, base api.SimResult) api.SimResult {
	out := base
	out.Request = req
	out.Equipped = &equipped
	return out
}
