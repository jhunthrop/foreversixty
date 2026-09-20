package measure

import (
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

// fits is what the cap-versus-budget question actually answers today,
// per precision, with the overrun spelled out. It is a table of the
// STATE, not of the intent: normal and high do not fit, and a test
// that quietly asserted they did - or that only ever looked at fast,
// which is what this test used to do - would be hiding it.
//
// This is an open contract question, not a bug to patch here.
// api.Caps[api.LaneServer] = 5,000 is contract value A2 and
// NativeIterationBudget() is read off a measurement, so "make it
// pass" would mean silently moving one of them. What A2 never says is
// WHICH precision the cap is scoped to: the arithmetic behind it
// (sim/api's Caps doc, and NativeIterationsPerCPUSecond's) is written
// entirely about a FAST run, and normal and high are both a single
// 1,000-iteration first stage over every combination, which is ten
// times the fast ladder's first rung. Either the cap is per-precision
// (5,000 fast, ~4,000 normal, ~3,900 high), or the premium job's
// shape has to grow, or normal and high are not offered at the server
// cap at all. Referred up.
var fits = map[string]struct {
	fits    bool
	overrun int // iterations over the budget, 0 when it fits
}{
	// 5,001 x 100, then a quarter (1,250) + equipped x 1,000, then the
	// top 10 + equipped x 3,000 = 1,784,100, inside the 4,092,480
	// budget.
	api.PrecisionFast: {fits: true},
	// 5,001 x 1,000, then the top 10 + equipped x 3,000 = 5,034,000:
	// 941,520 over, 23%.
	api.PrecisionNormal: {overrun: 941520},
	// 5,001 x 1,000, then the top 20 + equipped x 10,000 = 5,211,000:
	// 1,118,520 over, 27%.
	api.PrecisionHigh: {overrun: 1118520},
}

// The cap and the budget are two readings of one measurement, so the
// cap has to fit inside the budget and the cap A2 rejected has to not
// fit. The whole ladder is what decides it: a fast run's second stage,
// a quarter of the candidates at 1,000 iterations each, is far larger
// than the first stage's 100 apiece and is the reason 20,000 does not
// fit.
//
// Every precision is checked, because the cap is one number and a
// server run may be requested at any of the three. See fits, above,
// for why two of them are asserted NOT to fit.
func TestTheServerCapFitsTheJobBudget(t *testing.T) {
	budget := NativeIterationBudget()
	for _, precision := range api.Precisions {
		want, ok := fits[precision]
		if !ok {
			t.Fatalf("api.Precisions has %q and this test has no row for it", precision)
		}
		ladder := api.Ladders[precision]
		need := api.LadderIterations(ladder, api.Caps[api.LaneServer])
		switch {
		case want.fits && need > budget:
			t.Errorf("a full %s expansion of the %s cap is %d iterations and the job budget is %d; the cap and %d iterations per CPU-second disagree",
				precision, api.LaneServer, need, budget, NativeIterationsPerCPUSecond)
		case !want.fits && need <= budget:
			t.Errorf("a full %s expansion of the %s cap is %d iterations and now fits the %d-iteration budget; it used to be %d over. Either the measurement moved or the cap did - the open question in this file's `fits` table may be settled, so update it rather than deleting this branch",
				precision, api.LaneServer, need, budget, want.overrun)
		case !want.fits && need-budget != want.overrun:
			t.Errorf("a full %s expansion of the %s cap is %d over the budget, and this test was written when it was %d over; the cap, the ladder or the measurement moved",
				precision, api.LaneServer, need-budget, want.overrun)
		}
	}

	// The cap A2 rejected has to not fit at the precision A2 argued
	// about, which is fast.
	if need := api.LadderIterations(api.Ladders[api.PrecisionFast], rejectedCap); need <= budget {
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
	got := api.LadderIterations(ladder, combinations)
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
