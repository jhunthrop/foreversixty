package measure

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// ProcRate is a proc aura's observed rate and its observed internal
// cooldown. Neither is in DB2 at all: ItemEffect.TriggerType says
// on-equip, on-use or on-proc and nothing more, and an ICD is script-side.
type ProcRate struct {
	SpellID   int64  `json:"spell_id"`
	SpellName string `json:"spell_name"`

	Procs  int `json:"procs"`
	Swings int `json:"swings"`
	Casts  int `json:"casts"`

	PerSwing float64 `json:"per_swing"`
	PerCast  float64 `json:"per_cast"`
	// MinGap is the shortest interval between two procs. It is a lower
	// bound on the internal cooldown, and a tight one once the proc has
	// fired often enough: a real ICD shows as a hard floor that many
	// samples never cross.
	MinGap time.Duration `json:"min_gap"`

	Enough bool `json:"enough"`
}

// MeasureProcs counts each aura the actor gains on itself, against the
// swings and casts it made, and records the tightest gap between two
// applications.
func MeasureProcs(in Input) []ProcRate {
	var swings, casts int
	for _, e := range in.Events {
		if !byActor(e, in.Actor) {
			continue
		}
		switch e.Name {
		case "SWING_DAMAGE", "SWING_MISSED":
			swings++
		case "SPELL_CAST_SUCCESS":
			casts++
		}
	}

	type acc struct {
		row  ProcRate
		last time.Time
	}
	bySpell := map[int64]*acc{}
	var order []int64

	for _, e := range in.Events {
		// A proc is an aura the actor applies to itself. A buff from a
		// raid member is not a proc, so both ends must be the actor.
		if e.Kind != event.AuraApplied && e.Kind != event.AuraRefresh {
			continue
		}
		if e.Source.Name != in.Actor || e.Dest.Name != in.Actor {
			continue
		}
		a, ok := bySpell[e.Spell.ID]
		if !ok {
			a = &acc{row: ProcRate{SpellID: e.Spell.ID, SpellName: e.Spell.Name}}
			bySpell[e.Spell.ID] = a
			order = append(order, e.Spell.ID)
		}
		if !a.last.IsZero() {
			gap := e.Time.Sub(a.last)
			if a.row.MinGap == 0 || gap < a.row.MinGap {
				a.row.MinGap = gap
			}
		}
		a.last = e.Time
		a.row.Procs++
	}

	out := make([]ProcRate, 0, len(order))
	for _, id := range order {
		a := bySpell[id]
		a.row.Swings = swings
		a.row.Casts = casts
		if swings > 0 {
			a.row.PerSwing = float64(a.row.Procs) / float64(swings)
		}
		if casts > 0 {
			a.row.PerCast = float64(a.row.Procs) / float64(casts)
		}
		// The sample that matters is the number of chances, not the
		// number of procs: ten procs in twelve swings says nothing.
		a.row.Enough = swings+casts >= in.MinSamples
		out = append(out, a.row)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Procs > out[j].Procs })
	return out
}
