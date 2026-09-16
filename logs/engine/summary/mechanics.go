// logs/engine/summary/mechanics.go
package summary

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// MechanicsBlock is the fight's mechanics table, rendered against what
// actually happened. TableFound is false when the fight's encounter has no
// table, and the block is then empty.
type MechanicsBlock struct {
	TableFound bool          `json:"table_found"`
	Rows       []MechanicRow `json:"rows"`
}

// MechanicRow is one listed ability, with the fields for its kind filled
// in: avoidable and unavoidable abilities carry Players, interrupts carry
// Casts and Stopped, dispels carry Applied and Dispelled.
type MechanicRow struct {
	SpellID int64          `json:"spell_id"`
	Name    string         `json:"name"`
	Kind    mechanics.Kind `json:"kind"`
	Note    string         `json:"note,omitempty"`
	// Role is the table's: the one role meant to take this ability, if any.
	Role string `json:"role,omitempty"`
	// Avoidable and unavoidable: who it hit.
	Players []MechanicHit `json:"players,omitempty"`
	// Interrupt: casts the enemies started and how many were stopped.
	Casts   int64 `json:"casts,omitempty"`
	Stopped int64 `json:"stopped,omitempty"`
	// Dispel: applications on players and how many were dispelled.
	Applied   int64 `json:"applied,omitempty"`
	Dispelled int64 `json:"dispelled,omitempty"`
	// Interrupt and dispel: what the spell did when it went through -- the
	// damage it dealt the players and the healing it gave the enemies -- so
	// a drain that healed the boss eight times ranks by what it cost, not
	// by the count of eight.
	Damage int64 `json:"damage,omitempty"`
	Healed int64 `json:"healed,omitempty"`
}

// MechanicHit is one player's history with an avoidable or unavoidable
// ability.
type MechanicHit struct {
	GUID    string `json:"guid"`
	Name    string `json:"name"`
	Hits    int64  `json:"hits"`
	Damage  int64  `json:"damage"`
	FirstMS int64  `json:"first_ms"`
	LastMS  int64  `json:"last_ms"`
	// Killed is true when this player's killing blow was this mechanic.
	Killed bool `json:"killed"`
}

// mechanicCast is one enemy cast of a listed interrupt spell: when it began
// and, if a player stopped it, when.
type mechanicCast struct {
	start   time.Time
	stopped time.Time
}

// mechanicEffect is one tick of a listed effect: the damage it dealt a player
// or the healing it gave an enemy, and when.
type mechanicEffect struct {
	at     time.Time
	damage int64
	healed int64
}

// noteMechanicCast records, for the listed interrupt spells, each enemy cast
// beginning, each one being stopped, and each tick of the effects the table
// names for it. Called from Add for every event.
func (a *Accumulator) noteMechanicCast(e event.Event) {
	if a.opt.Mechanics == nil {
		return
	}
	switch e.Kind {
	case event.CastStart, event.CastSuccess:
		if m, ok := a.opt.Mechanics.Lookup(e.Spell.ID); ok && m.Kind == mechanics.Interrupt && !a.isPlayer(e.Source.GUID) {
			// A channel logs a start; an instant logs only a success. Either opens a
			// cast, unless a start already did a moment ago.
			casts := a.mechanicCasts[e.Spell.ID]
			if e.Kind == event.CastSuccess && len(casts) > 0 && casts[len(casts)-1].stopped.IsZero() &&
				e.Time.Sub(casts[len(casts)-1].start) < 15*time.Second {
				return
			}
			a.mechanicCasts[e.Spell.ID] = append(casts, mechanicCast{start: e.Time})
		}
	case event.Interrupt:
		casts := a.mechanicCasts[e.ExtraSpell.ID]
		if n := len(casts); n > 0 && casts[n-1].stopped.IsZero() {
			casts[n-1].stopped = e.Time
		}
	case event.Damage:
		if a.isPlayer(e.Dest.GUID) && !a.isPlayer(e.Source.GUID) {
			a.noteMechanicEffect(e.Spell.ID, mechanicEffect{at: e.Time, damage: e.Effective()})
		}
	case event.Heal:
		if !a.isPlayer(e.Source.GUID) && !a.isPlayer(e.Dest.GUID) {
			a.noteMechanicEffect(e.Spell.ID, mechanicEffect{at: e.Time, healed: e.Effective()})
		}
	}
}

func (a *Accumulator) noteMechanicEffect(spellID int64, effect mechanicEffect) {
	if _, listed := a.mechanicEffects[spellID]; listed {
		a.mechanicEffects[spellID] = append(a.mechanicEffects[spellID], effect)
		return
	}
	for _, m := range a.opt.Mechanics.Mechanics {
		if m.Kind != mechanics.Interrupt && m.Kind != mechanics.Dispel {
			continue
		}
		for _, id := range m.EffectIDs() {
			if id == spellID {
				a.mechanicEffects[spellID] = append(a.mechanicEffects[spellID], effect)
				return
			}
		}
	}
}

// noteMechanicHit records a listed ability landing on a player. Called from
// the Damage case of addDamageAndHealing, after the tables are folded.
func (a *Accumulator) noteMechanicHit(e event.Event) {
	if a.opt.Mechanics == nil {
		return
	}
	m, ok := a.opt.Mechanics.Lookup(e.Spell.ID)
	if !ok || (m.Kind != mechanics.Avoidable && m.Kind != mechanics.Unavoidable) {
		return
	}
	if !a.isPlayer(e.Dest.GUID) {
		return
	}
	// A mechanic is something the fight did to the raid. A spell id the table
	// lists that arrives from the player's own side -- a mirrored ability, a
	// friendly-fire cast, a pet -- is not the boss failing them, and counting
	// it would put a name on the problems list for someone else's mistake.
	if units.SameSide(e.Source.Flags, e.Dest.Flags) {
		return
	}
	bySpell := a.mechanicHits[e.Spell.ID]
	if bySpell == nil {
		bySpell = map[string]*MechanicHit{}
		a.mechanicHits[e.Spell.ID] = bySpell
	}
	hit := bySpell[e.Dest.GUID]
	if hit == nil {
		hit = &MechanicHit{GUID: e.Dest.GUID, Name: a.name(e.Dest.GUID), FirstMS: a.ms(e.Time)}
		bySpell[e.Dest.GUID] = hit
	}
	hit.Hits++
	hit.Damage += e.Effective()
	hit.LastMS = a.ms(e.Time)
}

// mechanicsBlock renders the table's rows for the fight. Interrupt and
// dispel rows read the casts, exchanges and aura tracks the accumulator
// already keeps.
func (a *Accumulator) mechanicsBlock(deaths []Death) MechanicsBlock {
	if a.opt.Mechanics == nil {
		return MechanicsBlock{Rows: []MechanicRow{}}
	}
	killedBy := map[int64]map[string]bool{}
	for _, d := range deaths {
		if d.KillingBlow == nil {
			continue
		}
		if killedBy[d.KillingBlow.SpellID] == nil {
			killedBy[d.KillingBlow.SpellID] = map[string]bool{}
		}
		killedBy[d.KillingBlow.SpellID][d.GUID] = true
	}
	stopped := map[int64]int64{}
	for _, x := range a.exchangeRows("interrupt") {
		stopped[x.ExtraSpellID] += x.Count
	}
	dispelled := map[int64]int64{}
	for _, x := range a.exchangeRows("dispel") {
		dispelled[x.ExtraSpellID] += x.Count
	}
	rows := make([]MechanicRow, 0, len(a.opt.Mechanics.Mechanics))
	for _, m := range a.opt.Mechanics.Mechanics {
		row := MechanicRow{SpellID: m.SpellID, Name: m.Name, Kind: m.Kind, Note: m.Note, Role: m.Role}
		switch m.Kind {
		case mechanics.Avoidable, mechanics.Unavoidable:
			for _, hit := range a.mechanicHits[m.SpellID] {
				h := *hit
				h.Killed = killedBy[m.SpellID][h.GUID]
				row.Players = append(row.Players, h)
			}
			sort.Slice(row.Players, func(i, j int) bool {
				if row.Players[i].Damage != row.Players[j].Damage {
					return row.Players[i].Damage > row.Players[j].Damage
				}
				return row.Players[i].GUID < row.Players[j].GUID
			})
		case mechanics.Interrupt:
			for _, c := range a.castRows() {
				if c.SpellID == m.SpellID && !a.isPlayer(c.GUID) {
					row.Casts += max(c.Started, c.Succeeded)
				}
			}
			row.Stopped = stopped[m.SpellID]
			if row.Casts < row.Stopped {
				row.Casts = row.Stopped
			}
			row.Damage, row.Healed = a.effectTotals(m, a.mechanicCasts[m.SpellID])
		case mechanics.Dispel:
			row.Damage, row.Healed = a.effectTotals(m, nil)
			for _, tr := range a.auraRows() {
				if tr.SpellID == m.SpellID && tr.Type == "DEBUFF" && a.isPlayer(tr.TargetGUID) {
					row.Applied += tr.Applications
				}
			}
			row.Dispelled = dispelled[m.SpellID]
		}
		rows = append(rows, row)
	}
	return MechanicsBlock{TableFound: true, Rows: rows}
}

// isPlayer reports whether the registry knows the guid as a player.
func (a *Accumulator) isPlayer(guid string) bool {
	u, ok := a.opt.Registry.Get(guid)
	return ok && u.IsPlayer()
}

// effectTotals is what a listed spell's effects did when it went through:
// every tick of the effect ids the table names for it, less the ticks of the
// casts a player stopped. A tick belongs to the latest cast begun before it,
// kicked or not: a kicked drain's ticks still in flight land after the kick
// and are the kick's cost, not the cost of letting it run. With no casts to
// lay ticks against (a dispel), every tick counts.
func (a *Accumulator) effectTotals(m mechanics.Mechanic, casts []mechanicCast) (damage, healed int64) {
	stopped := func(at time.Time) bool {
		var owner *mechanicCast
		for i := range casts {
			if !casts[i].start.After(at) {
				owner = &casts[i]
			}
		}
		return owner != nil && !owner.stopped.IsZero()
	}
	for _, id := range m.EffectIDs() {
		for _, tick := range a.mechanicEffects[id] {
			if stopped(tick.at) {
				continue
			}
			damage += tick.damage
			healed += tick.healed
		}
	}
	return damage, healed
}
