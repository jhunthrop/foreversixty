// logs/engine/summary/mechanics.go
package summary

import (
	"sort"

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
		row := MechanicRow{SpellID: m.SpellID, Name: m.Name, Kind: m.Kind, Note: m.Note}
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
			row.Damage, row.Healed = a.enemySpellTotals(m.EffectIDs())
		case mechanics.Dispel:
			row.Damage, row.Healed = a.enemySpellTotals(m.EffectIDs())
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

// enemySpellTotals is what listed spells did in the enemies' hands: the
// effective damage they dealt (the damage-done tables hold only damage to the
// other side, so a player's own copy of an id is not in it) and the healing
// they gave the enemies, so an interrupt or dispel that went through can be
// ranked by its cost. A channel's damage tick and heal carry their own ids,
// which the table lists as the cast's effects.
func (a *Accumulator) enemySpellTotals(spellIDs []int64) (damage, healed int64) {
	listed := map[int64]bool{}
	for _, id := range spellIDs {
		listed[id] = true
	}
	for guid, t := range a.damageDone {
		if a.isPlayer(guid) {
			continue
		}
		for key, ab := range t.abilities {
			if listed[key.spellID] {
				damage += ab.Effective
			}
		}
	}
	for guid, t := range a.healingDone {
		if a.isPlayer(guid) {
			continue
		}
		for key, ab := range t.abilities {
			if listed[key.spellID] {
				healed += ab.Effective
			}
		}
	}
	return damage, healed
}
