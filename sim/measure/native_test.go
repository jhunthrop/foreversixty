package measure

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// rejectedCap is the cap the contract's A2 turned down. The server cap
// is 5,000 and not this, because at the measured native rate a
// 20,000-combination fast run does not finish inside the job's
// 15-minute timeout. The number is here so the arithmetic that made
// the decision is the thing under test, not a restatement of the
// answer.
const rejectedCap = 20000

// ladderIterations is what a whole bulk expansion costs: every stage
// of the precision's ladder, summed.
//
// The equipped set runs in every stage - that is what pairs a delta
// with something - so a stage of n surviving candidates is n+1 runs.
// Between stages a cut keeps a fraction of the survivors or a fixed
// top few.
//
// It is a floor, not the exact figure: Cut.SlackSE keeps anything
// whose interval still overlaps the last survivor's, so a real stage
// runs at least this many and usually a few more. That makes "over
// budget" a certainty and "under budget" the optimistic reading, which
// is the right way round for a cap.
func ladderIterations(l api.Ladder, combinations int) int {
	total := 0
	survivors := combinations
	for i, iterations := range l.Iterations {
		total += (survivors + 1) * iterations
		if i < len(l.Cuts) {
			survivors = survive(l.Cuts[i], survivors)
		}
	}
	return total
}

// survive applies one cut to a stage's surviving candidates.
func survive(c api.Cut, survivors int) int {
	switch {
	case c.Fraction > 0:
		return int(math.Ceil(c.Fraction * float64(survivors)))
	case c.Top > 0:
		return min(c.Top, survivors)
	default:
		return survivors
	}
}

// The cap and the budget are two readings of one measurement, so the
// cap has to fit inside the budget and the cap A2 rejected has to not
// fit. The whole ladder is what decides it: a fast run's second stage,
// a quarter of the candidates at 1,000 iterations each, is far larger
// than the first stage's 100 apiece and is the reason 20,000 does not
// fit.
func TestTheServerCapFitsTheJobBudget(t *testing.T) {
	ladder := api.Ladders[api.PrecisionFast]
	budget := NativeIterationBudget()

	if need := ladderIterations(ladder, api.Caps[api.LaneServer]); need > budget {
		t.Errorf("a full %s expansion of the %s cap is %d iterations and the job budget is %d; the cap and %d iterations per CPU-second disagree",
			api.PrecisionFast, api.LaneServer, need, budget, NativeIterationsPerCPUSecond)
	}
	if need := ladderIterations(ladder, rejectedCap); need <= budget {
		t.Errorf("a full %s expansion of %d combinations is %d iterations and the job budget is %d; A2 rejected that cap because it does not fit, so either the measurement moved or this arithmetic is wrong",
			api.PrecisionFast, rejectedCap, need, budget)
	}
}

// Every stage is counted, not just the first: a sum that quietly
// dropped the later stages would still pass the two bounds above by
// luck, which is exactly how the first version of this test managed to
// bound nothing. Worked out by hand on a fixed size, so it pins the
// arithmetic rather than whatever the cap happens to be.
func TestTheLadderSumCountsEveryStage(t *testing.T) {
	ladder := api.Ladders[api.PrecisionFast]
	const combinations = 5000
	// 5,001 x 100, then a quarter (1,250) + equipped x 1,000, then the
	// top 10 + equipped x 3,000.
	const want = 5001*100 + 1251*1000 + 11*3000
	got := ladderIterations(ladder, combinations)
	if got != want {
		t.Errorf("a %s expansion of %d combinations sums to %d, want %d", api.PrecisionFast, combinations, got, want)
	}
	if first := (combinations + 1) * ladder.Iterations[0]; got <= first {
		t.Errorf("the sum %d is not larger than the first stage alone (%d); the later stages are being dropped", got, first)
	}
}

func TestTheBudgetIsPositive(t *testing.T) {
	if NativeIterationBudget() <= 0 {
		t.Fatal("the job budget is not positive")
	}
}
