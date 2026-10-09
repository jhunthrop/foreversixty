package score

import (
	"math"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// HealingResult is what a healer did against a heal profile's fake raid.
type HealingResult struct {
	// Effective is healing that landed on a health bar (or soaked damage,
	// for a shield) per second; Raw includes the overheal.
	Effective, Raw api.Estimate
	// ManaLastsSec is when the healer first ran out of mana, averaged over
	// the iterations; an iteration that never did contributes the time it
	// would have, projected from its mana left (the engine's own
	// convention), so a value past the fight's length means mana to spare.
	ManaLastsSec float64
	// ManaSpent and EffectiveHealed are totals over every iteration, so
	// their ratio is healing per mana without a per-iteration rounding.
	ManaSpent       float64
	EffectiveHealed float64
}

// HealingPerMana is effective healing per point of mana spent.
func (r HealingResult) HealingPerMana() float64 {
	if r.ManaSpent <= 0 {
		return 0
	}
	return r.EffectiveHealed / r.ManaSpent
}

// OverhealShare is the share of raw healing that overhealed, 0 to 1.
func (r HealingResult) OverhealShare() float64 {
	if r.Raw.Mean <= 0 {
		return 0
	}
	return math.Max(0, 1-r.Effective.Mean/r.Raw.Mean)
}

// HealingResultOf reads a healer's metrics from a run of the given length.
func HealingResultOf(player *proto.UnitMetrics, iterations int32, durationSec float64) HealingResult {
	out := HealingResult{
		Effective:    EstimateOf(player.EffectiveHps, iterations),
		Raw:          EstimateOf(player.Hps, iterations),
		ManaLastsSec: player.GetTto().GetAvg(),
		ManaSpent:    ManaSpent(player.Resources),
	}
	out.EffectiveHealed = out.Effective.Mean * durationSec * float64(iterations)
	return out
}

// ManaSpent is every point of mana the healer paid out over the run: the
// engine records a spend as a negative gain on the action that cost it.
func ManaSpent(resources []*proto.ResourceMetrics) float64 {
	spent := 0.0
	for _, r := range resources {
		if r.Type == proto.ResourceType_ResourceTypeMana && r.Gain < 0 {
			spent -= r.Gain
		}
	}
	return spent
}

// Block is the envelope's healing block for the run.
func (r HealingResult) Block() *api.HealingResult {
	return &api.HealingResult{
		EffectiveHPS: r.Effective,
		HPS:          r.Raw,
		ManaLastsSec: r.ManaLastsSec,
		HPM:          r.HealingPerMana(),
	}
}
