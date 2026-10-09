package combine

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/score"
)

// poolRoles pools the healing and tank blocks of a role_metrics run. Parts
// of one run either all carry a block or none does; the request they share
// decides which, so a mix is a corrupt set of parts.
func poolRoles(parts []api.SimResult) (*api.HealingResult, *api.TankResult, error) {
	healing, tank := 0, 0
	for _, p := range parts {
		if p.Healing != nil {
			healing++
		}
		if p.Tank != nil {
			tank++
		}
	}
	if (healing != 0 && healing != len(parts)) || (tank != 0 && tank != len(parts)) {
		return nil, nil, fmt.Errorf("%w: only some parts carry the role's figures", ErrMixedParts)
	}
	var pooledHealing *api.HealingResult
	if healing > 0 {
		pooledHealing = poolHealing(parts)
	}
	var pooledTank *api.TankResult
	if tank > 0 {
		var err error
		if pooledTank, err = poolTank(parts); err != nil {
			return nil, nil, err
		}
	}
	return pooledHealing, pooledTank, nil
}

// iterationWeighted is the mean of a per-part scalar weighted by the
// iterations behind each part.
func iterationWeighted(parts []api.SimResult, read func(api.SimResult) float64) float64 {
	var sum, total float64
	for _, p := range parts {
		sum += read(p) * float64(p.IterationsRun)
		total += float64(p.IterationsRun)
	}
	return sum / total
}

// poolHealing pools the healer's estimates. Healing per mana is a ratio of
// totals, so it is pooled the same way: each part's mana spent is its
// effective healing over its own ratio, and the pooled ratio is the summed
// healing over the summed mana (the fight length is the same for every
// part, so it cancels).
func poolHealing(parts []api.SimResult) *api.HealingResult {
	var healed, mana float64
	for _, p := range parts {
		h := p.Healing.EffectiveHPS.Mean * float64(p.IterationsRun)
		healed += h
		if p.Healing.HPM > 0 {
			mana += h / p.Healing.HPM
		}
	}
	out := &api.HealingResult{
		EffectiveHPS: poolEstimates(parts, func(p api.SimResult) api.Estimate { return p.Healing.EffectiveHPS }),
		HPS:          poolEstimates(parts, func(p api.SimResult) api.Estimate { return p.Healing.HPS }),
		ManaLastsSec: iterationWeighted(parts, func(p api.SimResult) float64 { return p.Healing.ManaLastsSec }),
	}
	if mana > 0 {
		out.HPM = healed / mana
	}
	return out
}

// poolTank pools the three measured terms and then scores them. The score
// is nonlinear in them, so pooling per-part scores would not be the score
// of the pooled run; sim/score computes it from the pooled terms, exactly as
// it does for a serial run.
func poolTank(parts []api.SimResult) (*api.TankResult, error) {
	run := score.TankRunResult{
		DTPS: poolEstimates(parts, func(p api.SimResult) api.Estimate { return p.Tank.DTPS }),
		TPS:  poolEstimates(parts, func(p api.SimResult) api.Estimate { return p.Tank.TPS }),
		TMI:  poolEstimates(parts, func(p api.SimResult) api.Estimate { return p.Tank.TMI }),
		// Maximum health is the same deterministic number in every part.
		Health:        parts[0].Tank.Health,
		ChanceOfDeath: iterationWeighted(parts, func(p api.SimResult) float64 { return p.Tank.ChanceOfDeath }),
	}
	block, err := score.TankBlockOf(run, parts[0].Request.Character.Level)
	if err != nil {
		return nil, fmt.Errorf("combine: scoring the pooled tank run: %w", err)
	}
	return block, nil
}
