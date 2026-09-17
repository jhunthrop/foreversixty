// logs/engine/summary/phases.go
package summary

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// Phase is one named stretch of the fight, from its trigger to the next
// trigger or the fight's end.
type Phase struct {
	Name    string `json:"name"`
	StartMS int64  `json:"start_ms"`
	EndMS   int64  `json:"end_ms"`
}

// firstPhaseName is what the stretch before the first trigger is called. The
// table names the phases its triggers open; the pull's own opening has no
// trigger and no entry, and every encounter calls it the same thing.
const firstPhaseName = "Phase 1"

// notePhase records the first instant each of the encounter's phase triggers
// fired. A trigger fires once: a boss that casts its opener twice is still in
// the phase its first cast began, and a health threshold crossed again after a
// heal did not start the phase over.
func (a *Accumulator) notePhase(e event.Event) {
	if a.opt.Mechanics == nil || len(a.opt.Mechanics.Phases) == 0 {
		return
	}
	for i, p := range a.opt.Mechanics.Phases {
		if _, fired := a.phaseAt[i]; fired {
			continue
		}
		if a.phaseFires(p.Starts, e) {
			a.phaseAt[i] = e.Time
		}
	}
}

// phaseFires reports whether this event is the trigger. A spell trigger is the
// enemies' own line -- a player casting the same id is not the boss changing
// phase -- and a health trigger is the boss's own health, read off the advanced
// block of any line that describes the unit the table names, which is the same
// identity the fight list reads a wipe percentage from.
func (a *Accumulator) phaseFires(start mechanics.PhaseStart, e event.Event) bool {
	if start.SpellID > 0 {
		if e.Spell.ID != start.SpellID || a.isPlayer(e.Source.GUID) {
			return false
		}
		switch start.On {
		case mechanics.OnCastStart:
			return e.Kind == event.CastStart
		case mechanics.OnCastSuccess:
			return e.Kind == event.CastSuccess
		case mechanics.OnAuraApplied:
			return e.Kind == event.AuraApplied
		case mechanics.OnAuraRemoved:
			return e.Kind == event.AuraRemoved
		}
		return false
	}
	if !e.Adv.OK || e.Adv.MaxHP <= 0 || e.Adv.InfoGUID == "" {
		return false
	}
	if a.name(e.Adv.InfoGUID) != a.opt.Mechanics.Name {
		return false
	}
	return float64(e.Adv.CurrentHP)/float64(e.Adv.MaxHP)*100 <= start.HealthPct
}

// phaseRows renders the fight's phases: each from its trigger to the next
// trigger or the fight's end, with the pull's own opening in front of them.
// A trigger that never fired is not a phase the fight reached and is left out.
// A fight that reached none at all has no phases: a single band called "Phase 1"
// across the whole pull says nothing a reader can act on, and a preset for it
// would be the whole-fight preset under another name.
func (a *Accumulator) phaseRows(durationMS int64) []Phase {
	out := []Phase{}
	if a.opt.Mechanics == nil || len(a.phaseAt) == 0 {
		return out
	}
	type fired struct {
		name string
		at   int64
	}
	rows := []fired{{name: firstPhaseName, at: 0}}
	for i, p := range a.opt.Mechanics.Phases {
		if t, ok := a.phaseAt[i]; ok {
			rows = append(rows, fired{name: p.Name, at: a.ms(t)})
		}
	}
	// Stable, and by instant rather than by table order: a curated list can name a
	// health threshold that a fight crossed before a cast the table lists first.
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].at < rows[j].at })
	for i, r := range rows {
		end := durationMS
		if i+1 < len(rows) {
			end = rows[i+1].at
		}
		if end < r.at {
			end = r.at
		}
		out = append(out, Phase{Name: r.name, StartMS: r.at, EndMS: end})
	}
	return out
}
