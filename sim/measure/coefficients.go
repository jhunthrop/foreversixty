package measure

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// Coefficient is a spell's observed damage against a known character
// sheet. EffectBonusCoefficient is routinely 0 or wrong for
// Classic-lineage spells, so the observed figure is the only evidence
// there is until the sim's own prediction can be diffed against it.
type Coefficient struct {
	SpellID   int64  `json:"spell_id"`
	SpellName string `json:"spell_name"`

	Hits       int     `json:"hits"`
	MeanDamage float64 `json:"mean_damage"`
	StdDev     float64 `json:"stddev"`
	// Observed is mean damage per point of spell power, and is zero when
	// Input.SpellPower is zero. It is not the coefficient: the base
	// damage has not been subtracted, because the base is what the
	// constants file is for. It is the figure to diff a candidate
	// coefficient against.
	Observed float64 `json:"observed"`

	Enough bool `json:"enough"`
}

// MeasureCoefficients reports the mean and spread of each direct spell's
// damage. Critical, glancing and blocked hits are excluded, because each
// carries its own multiplier and would widen the spread without telling
// anyone anything.
func MeasureCoefficients(in Input) []Coefficient {
	amounts := map[int64][]float64{}
	names := map[int64]string{}
	var order []int64

	for _, e := range in.Events {
		if e.Kind != event.Damage || isPeriodic(e) || !byActor(e, in.Actor) {
			continue
		}
		if e.Name != "SPELL_DAMAGE" || !e.Amount.OK {
			continue
		}
		if e.Critical.OK && e.Critical.V {
			continue
		}
		if e.Glancing.OK && e.Glancing.V {
			continue
		}
		// A partial block took damage away from the landing, so the
		// amount is not the coefficient's: leaving it in dragged the
		// mean down and widened the spread, which is the one figure
		// this function exists to produce. The doc above always said
		// blocked hits were excluded; now they are.
		if e.Blocked.OK && e.Blocked.V > 0 {
			continue
		}
		if _, ok := amounts[e.Spell.ID]; !ok {
			names[e.Spell.ID] = e.Spell.Name
			order = append(order, e.Spell.ID)
		}
		amounts[e.Spell.ID] = append(amounts[e.Spell.ID], float64(e.Amount.V))
	}

	out := make([]Coefficient, 0, len(order))
	for _, id := range order {
		xs := amounts[id]
		row := Coefficient{
			SpellID:    id,
			SpellName:  names[id],
			Hits:       len(xs),
			MeanDamage: mean(xs),
			StdDev:     stddev(xs),
			Enough:     len(xs) >= in.MinSamples,
		}
		if in.SpellPower > 0 {
			row.Observed = row.MeanDamage / in.SpellPower
		}
		out = append(out, row)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Hits > out[j].Hits })
	return out
}
