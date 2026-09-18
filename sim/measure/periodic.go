package measure

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// PeriodicCrit is what one periodic spell's ticks say about two rules the
// client tables do not carry: whether periodic damage can crit at all,
// and what multiplier a critical tick uses.
type PeriodicCrit struct {
	SpellID   int64  `json:"spell_id"`
	SpellName string `json:"spell_name"`
	School    int64  `json:"school"`
	// Physical separates the two answers, because Forever may let
	// physical dots crit and magic ones not, or use different
	// multipliers. Reported separately rather than averaged.
	Physical bool `json:"physical"`

	Ticks     int `json:"ticks"`
	CritTicks int `json:"crit_ticks"`
	// CanCrit is true if any tick carried the critical flag. One is
	// enough to settle the question; zero out of many is evidence the
	// other way, which is why Ticks is printed beside it.
	CanCrit bool `json:"can_crit"`

	MeanNormal float64 `json:"mean_normal"`
	MeanCrit   float64 `json:"mean_crit"`
	// Multiplier is MeanCrit over MeanNormal, zero when either side has
	// no samples.
	Multiplier float64 `json:"multiplier"`

	Enough bool `json:"enough"`
}

// MeasurePeriodic answers, per periodic spell, whether its ticks crit and
// by how much.
func MeasurePeriodic(in Input) []PeriodicCrit {
	type acc struct {
		row     PeriodicCrit
		normals []float64
		crits   []float64
	}
	bySpell := map[int64]*acc{}
	var order []int64

	for _, e := range in.Events {
		if e.Kind != event.Damage || !isPeriodic(e) || !byActor(e, in.Actor) {
			continue
		}
		if !e.Amount.OK {
			continue
		}
		a, ok := bySpell[e.Spell.ID]
		if !ok {
			a = &acc{row: PeriodicCrit{
				SpellID:   e.Spell.ID,
				SpellName: e.Spell.Name,
				School:    e.Spell.School,
				Physical:  e.Spell.School == physicalSchool,
			}}
			bySpell[e.Spell.ID] = a
			order = append(order, e.Spell.ID)
		}
		a.row.Ticks++
		amount := float64(e.Amount.V)
		if e.Critical.OK && e.Critical.V {
			a.row.CritTicks++
			a.row.CanCrit = true
			a.crits = append(a.crits, amount)
		} else {
			a.normals = append(a.normals, amount)
		}
	}

	out := make([]PeriodicCrit, 0, len(order))
	for _, id := range order {
		a := bySpell[id]
		a.row.MeanNormal = mean(a.normals)
		a.row.MeanCrit = mean(a.crits)
		if a.row.MeanNormal > 0 && a.row.MeanCrit > 0 {
			a.row.Multiplier = a.row.MeanCrit / a.row.MeanNormal
		}
		a.row.Enough = a.row.Ticks >= in.MinSamples
		out = append(out, a.row)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Ticks > out[j].Ticks })
	return out
}
