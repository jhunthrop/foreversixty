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
	// ManaLeft is the mana in the bar when the fight ends, averaged over the
	// iterations: the pool it started with plus every point gained (none over
	// the cap) less every point spent. Zero until WithManaPool sets the pool.
	ManaLeft float64
	// ManaPool is the mana the healer started the fight with.
	ManaPool float64
	// manaNet is the mana gained less spent per iteration, kept for WithManaPool.
	manaNet float64
}

// WithManaPool is the result with the starting pool known, so its mana left
// at the end can be stated. The receiver is not changed.
func (r HealingResult) WithManaPool(pool float64) HealingResult {
	r.ManaPool = pool
	r.ManaLeft = math.Max(0, pool+r.manaNet)
	return r
}

// ManaLeftShare is the mana left at the end as a share of the pool, 0 to 1.
func (r HealingResult) ManaLeftShare() float64 {
	if r.ManaPool <= 0 {
		return 0
	}
	return math.Min(1, r.ManaLeft/r.ManaPool)
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
	if iterations > 0 {
		out.manaNet = manaNet(player.Resources) / float64(iterations)
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

// manaNet is every point of mana the healer gained (none over the cap) less
// every point it spent, over the run.
func manaNet(resources []*proto.ResourceMetrics) float64 {
	net := 0.0
	for _, r := range resources {
		if r.Type == proto.ResourceType_ResourceTypeMana {
			net += r.ActualGain
		}
	}
	return net
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
