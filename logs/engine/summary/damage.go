// logs/engine/summary/damage.go
package summary

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Ability is one spell's contribution to an actor's row. SpellID 0 is melee.
type Ability struct {
	SpellID   int64            `json:"spell_id"`
	Name      string           `json:"name"`
	School    int64            `json:"school,omitempty"`
	Total     int64            `json:"total"`
	Effective int64            `json:"effective"`
	Overheal  int64            `json:"overheal,omitempty"`
	Overkill  int64            `json:"overkill,omitempty"`
	Absorbed  int64            `json:"absorbed,omitempty"`
	Resisted  int64            `json:"resisted,omitempty"`
	Blocked   int64            `json:"blocked,omitempty"`
	Hits      int64            `json:"hits"`
	Crits     int64            `json:"crits"`
	Ticks     int64            `json:"ticks"`
	Misses    map[string]int64 `json:"misses,omitempty"`
	Min       int64            `json:"min"`
	Max       int64            `json:"max"`
}

// Pair is one target's share of an actor's total.
type Pair struct {
	GUID  string `json:"guid"`
	Name  string `json:"name"`
	Total int64  `json:"total"`
}

// Actor is one row of a damage or healing table.
type Actor struct {
	GUID      string    `json:"guid"`
	Name      string    `json:"name"`
	Class     string    `json:"class,omitempty"`
	Total     int64     `json:"total"`
	Effective int64     `json:"effective"`
	Overheal  int64     `json:"overheal,omitempty"`
	Absorbed  int64     `json:"absorbed,omitempty"`
	ActiveMS  int64     `json:"active_ms"`
	Abilities []Ability `json:"abilities"`
	Targets   []Pair    `json:"targets"`
	Series    []int64   `json:"series"`
}

type actor struct {
	guid      string
	total     int64
	effective int64
	overheal  int64
	absorbed  int64
	abilities map[int64]*Ability
	targets   map[string]int64
	series    []int64
}

// activity tracks how long an actor spent doing something, so a row can
// show damage per active second as well as per fight second.
type activity struct {
	ms   int64
	last time.Time
}

func (a *Accumulator) markActive(guid string, at time.Time) {
	if guid == "" || guid == units.NoGUID {
		return
	}
	act, ok := a.active[guid]
	if !ok {
		a.active[guid] = &activity{ms: a.opt.ActiveGap.Milliseconds(), last: at}
		return
	}
	gap := at.Sub(act.last)
	if gap < 0 {
		gap = 0
	}
	if gap >= a.opt.ActiveGap {
		act.ms += a.opt.ActiveGap.Milliseconds()
	} else {
		act.ms += gap.Milliseconds()
	}
	act.last = at
}

func (a *Accumulator) table(m map[string]*actor, guid string) *actor {
	t, ok := m[guid]
	if !ok {
		t = &actor{guid: guid, abilities: map[int64]*Ability{}, targets: map[string]int64{}}
		m[guid] = t
	}
	return t
}

func (t *actor) ability(e event.Event) *Ability {
	ab, ok := t.abilities[e.Spell.ID]
	if !ok {
		name := e.Spell.Name
		if e.Spell.ID == 0 && name == "" {
			name = "Melee"
		}
		ab = &Ability{SpellID: e.Spell.ID, Name: name, School: e.Spell.School}
		t.abilities[e.Spell.ID] = ab
	}
	return ab
}

func (t *actor) addSeries(bucket int, amount int64) {
	for len(t.series) <= bucket {
		t.series = append(t.series, 0)
	}
	t.series[bucket] += amount
}

// addDamageAndHealing folds damage, healing, misses, and absorbs into the
// four tables. Pet output is credited to the pet's owner.
func (a *Accumulator) addDamageAndHealing(e event.Event) {
	switch e.Kind {
	case event.Damage:
		src := a.owner(e.Source.GUID)
		amount, effective := e.Amount.V, e.Effective()
		done := a.table(a.damageDone, src)
		a.fold(done, e, amount, effective, e.Dest.GUID)
		taken := a.table(a.damageTaken, e.Dest.GUID)
		a.fold(taken, e, amount, effective, src)
		a.markActive(src, e.Time)
		a.threat[src] += a.opt.Threat.Damage(e)

	case event.Heal:
		src := a.owner(e.Source.GUID)
		amount, effective := e.Amount.V, e.Effective()
		done := a.table(a.healingDone, src)
		a.fold(done, e, amount, effective, e.Dest.GUID)
		done.overheal += e.Overheal.V
		done.absorbed += e.Absorbed.V
		if ab := done.ability(e); ab != nil {
			ab.Overheal += e.Overheal.V
		}
		taken := a.table(a.healingTaken, e.Dest.GUID)
		a.fold(taken, e, amount, effective, src)
		taken.overheal += e.Overheal.V
		a.markActive(src, e.Time)
		a.threat[src] += a.opt.Threat.Healing(e)

	case event.Missed:
		src := a.owner(e.Source.GUID)
		done := a.table(a.damageDone, src)
		ab := done.ability(e)
		if ab.Misses == nil {
			ab.Misses = map[string]int64{}
		}
		ab.Misses[e.MissType]++
		taken := a.table(a.damageTaken, e.Dest.GUID)
		tab := taken.ability(e)
		if tab.Misses == nil {
			tab.Misses = map[string]int64{}
		}
		tab.Misses[e.MissType]++

	case event.Absorbed:
		// The shield's owner gets credit for the absorb as healing done:
		// this is the only place the absorbed amount is counted, since the
		// absorbed field on damage events is informational.
		caster := a.owner(e.ExtraUnit.GUID)
		if caster == "" {
			return
		}
		shield := event.Event{
			Time: e.Time, Kind: event.Heal,
			Source: e.ExtraUnit, Dest: e.Dest, Spell: e.ExtraSpell,
			Amount: e.Amount,
		}
		done := a.table(a.healingDone, caster)
		a.fold(done, shield, e.Amount.V, e.Amount.V, e.Dest.GUID)
		done.absorbed += e.Amount.V
		if ab := done.ability(shield); ab != nil {
			ab.Absorbed += e.Amount.V
		}
		a.markActive(caster, e.Time)
	}
}

func (a *Accumulator) fold(t *actor, e event.Event, amount, effective int64, target string) {
	t.total += amount
	t.effective += effective
	t.targets[target] += effective
	t.addSeries(a.bucket(e.Time), effective)

	ab := t.ability(e)
	ab.Total += amount
	ab.Effective += effective
	ab.Overkill += max(e.Overkill.V, 0)
	ab.Absorbed += e.Absorbed.V
	ab.Resisted += e.Resisted.V
	ab.Blocked += e.Blocked.V
	if e.Name == "SPELL_PERIODIC_DAMAGE" || e.Name == "SPELL_PERIODIC_HEAL" {
		ab.Ticks++
	} else {
		ab.Hits++
	}
	if e.Critical.V {
		ab.Crits++
	}
	if amount > ab.Max {
		ab.Max = amount
	}
	if ab.Min == 0 || (amount > 0 && amount < ab.Min) {
		ab.Min = amount
	}
}

// actors renders one table, sorted for determinism.
func (a *Accumulator) actors(m map[string]*actor) []Actor {
	out := make([]Actor, 0, len(m))
	for guid, t := range m {
		row := Actor{
			GUID: guid, Name: a.name(guid),
			Total: t.total, Effective: t.effective,
			Overheal: t.overheal, Absorbed: t.absorbed,
			Series: t.series,
		}
		if row.Series == nil {
			row.Series = []int64{}
		}
		if a.opt.Registry != nil {
			if u, ok := a.opt.Registry.Get(guid); ok {
				row.Class = u.Class
			}
		}
		if act, ok := a.active[guid]; ok {
			row.ActiveMS = act.ms
		}
		row.Abilities = make([]Ability, 0, len(t.abilities))
		for _, ab := range t.abilities {
			row.Abilities = append(row.Abilities, *ab)
		}
		sort.Slice(row.Abilities, func(i, j int) bool {
			if row.Abilities[i].Effective != row.Abilities[j].Effective {
				return row.Abilities[i].Effective > row.Abilities[j].Effective
			}
			return row.Abilities[i].SpellID < row.Abilities[j].SpellID
		})
		row.Targets = make([]Pair, 0, len(t.targets))
		for g, v := range t.targets {
			row.Targets = append(row.Targets, Pair{GUID: g, Name: a.name(g), Total: v})
		}
		sort.Slice(row.Targets, func(i, j int) bool {
			if row.Targets[i].Total != row.Targets[j].Total {
				return row.Targets[i].Total > row.Targets[j].Total
			}
			return row.Targets[i].GUID < row.Targets[j].GUID
		})
		out = append(out, row)
	}
	sortActors(out)
	return out
}
