// logs/engine/rating/combine.go
package rating

import "github.com/jhunthrop/foreversixty/logs/engine/mechanics/weights"

// weightOf returns a component's site-default weight for one role's weight
// row, by name.
func weightOf(w weights.RoleWeights, name string) float64 {
	switch name {
	case ComponentNameOutput:
		return w.Output
	case ComponentNameSurvival:
		return w.Survival
	case ComponentNameMechanics:
		return w.Mechanics
	case ComponentNameUtility:
		return w.Utility
	case ComponentNamePreparation:
		return w.Preparation
	default:
		return w.Activity
	}
}

// combine applies spec §1.1's renormalisation and §1.5's weighted mean and
// Survival-catastrophe cap to six already-scored components, in place:
// every non-excluded Component's Weight field is overwritten with its
// renormalised share (spec §4.1's "the *renormalised* weight actually
// applied"). survivalDeathScoreZero is true only when Survival was scored
// (not excluded) and its death half alone floored to zero (spec §1.5's cap
// condition) — Component carries no field for "the death half
// specifically," so the Survival scorer threads it through here separately
// rather than it being re-derived from Component.Score, which mixes the
// death and avoidable-hit halves together.
func combine(
	components *[6]Component, w weights.RoleWeights,
	survivalDeathScoreZero, capEnabled bool, capThreshold float64,
) (overallUncapped, overall float64, capped bool, basis string) {
	var totalWeight float64
	for _, c := range components {
		if !c.Excluded {
			totalWeight += weightOf(w, c.Name)
		}
	}

	basisKinds := map[string]bool{}
	var sum float64
	for i := range components {
		c := &components[i]
		if c.Excluded {
			c.Weight = 0
			continue
		}
		if totalWeight <= 0 {
			c.Weight = 0
			continue
		}
		share := weightOf(w, c.Name) / totalWeight
		c.Weight = round2(share * 100)
		sum += share * c.Score
		if c.Basis != "" {
			basisKinds[c.Basis] = true
		}
	}

	overallUncapped = round2(sum)
	overall = overallUncapped
	capped = survivalDeathScoreZero
	if capped && capEnabled && overallUncapped > capThreshold {
		overall = capThreshold
	}

	switch len(basisKinds) {
	case 1:
		for k := range basisKinds {
			basis = k
		}
	case 0:
		basis = ""
	default:
		basis = "mixed"
	}
	return overallUncapped, overall, capped, basis
}
