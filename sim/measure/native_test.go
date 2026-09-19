package measure

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// The cap and the budget are two readings of one measurement, so the
// cap has to fit inside the budget. A fast run is 100 + 1,000
// iterations over the surviving quarter, plus the finalists at 3,000;
// the first stage dominates and is what this bounds.
func TestTheServerCapFitsTheJobBudget(t *testing.T) {
	first := api.Ladders[api.PrecisionFast].Iterations[0]
	need := (api.Caps[api.LaneServer] + 1) * first
	if budget := NativeIterationBudget(); need > budget {
		t.Errorf("a full %s expansion's first stage is %d iterations and the job budget is %d; the cap and %d iterations per CPU-second disagree",
			api.PrecisionFast, need, budget, NativeIterationsPerCPUSecond)
	}
}

func TestTheBudgetIsPositive(t *testing.T) {
	if NativeIterationBudget() <= 0 {
		t.Fatal("the job budget is not positive")
	}
}
