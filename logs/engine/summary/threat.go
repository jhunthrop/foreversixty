// logs/engine/summary/threat.go
package summary

import (
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// ThreatModel turns damage and healing into threat. It is an interface
// because the vanilla model is per class, per stance, and per ability, and
// those coefficients are data, not code: they arrive with Forever's own
// numbers rather than being guessed here.
type ThreatModel interface {
	// Version identifies the model in the summary, so a report says which
	// numbers produced its threat table.
	Version() string
	// Complete reports whether the model has the per-class modifiers it
	// needs. A false here makes the report mark the threat table as
	// provisional rather than presenting a wrong number as a right one.
	Complete() bool
	// Damage is the threat one damage event generates.
	Damage(e event.Event) float64
	// Healing is the threat one heal generates.
	Healing(e event.Event) float64
}

// BaseThreat is the part of the vanilla model that is mechanically certain
// and version independent: a point of damage to a hostile unit is a point
// of threat, and a point of effective healing is half a point, spread over
// whatever the healer is in combat with.
//
// Modifiers holds the per-spell and per-stance multipliers. It is empty
// until Forever's beta log and the community tables settle the numbers, and
// Complete reports false while it is, so nothing downstream mistakes this
// for a finished threat table. Fill it from a data file; do not hard-code
// coefficients here.
type BaseThreat struct {
	// Modifiers maps a spell id to a multiplier applied to that spell's
	// threat. A missing entry means 1.0.
	Modifiers map[int64]float64
	// HealingCoefficient is the share of effective healing that becomes
	// threat. Vanilla's value is 0.5.
	HealingCoefficient float64
}

// Version names the model.
func (b BaseThreat) Version() string { return "base-1" }

// Complete reports whether the per-spell modifiers have been supplied.
func (b BaseThreat) Complete() bool { return len(b.Modifiers) > 0 }

// Damage returns the threat a damage event generates.
func (b BaseThreat) Damage(e event.Event) float64 {
	return float64(e.Effective()) * b.modifier(e.Spell.ID)
}

// Healing returns the threat a heal generates.
func (b BaseThreat) Healing(e event.Event) float64 {
	c := b.HealingCoefficient
	if c == 0 {
		c = 0.5
	}
	return float64(e.Effective()) * c * b.modifier(e.Spell.ID)
}

func (b BaseThreat) modifier(spellID int64) float64 {
	if m, ok := b.Modifiers[spellID]; ok {
		return m
	}
	return 1
}

// ThreatRow is one actor's threat for the fight.
type ThreatRow struct {
	GUID         string  `json:"guid"`
	Name         string  `json:"name"`
	Threat       float64 `json:"threat"`
	ModelVersion string  `json:"model_version"`
	Complete     bool    `json:"complete"`
}

// ThreatPair is one player's threat on one enemy for the fight.
type ThreatPair struct {
	GUID       string  `json:"guid"`
	Name       string  `json:"name"`
	TargetGUID string  `json:"target_guid"`
	TargetName string  `json:"target_name"`
	Threat     float64 `json:"threat"`
}

// creditThreat books threat from one player against one enemy. The caller
// is responsible for knowing enemy is hostile; creditThreat itself trusts
// it, the same way the damage and heal cases already trust the event's own
// flags rather than a registry lookup (see engage below).
func (a *Accumulator) creditThreat(player, enemy string, threat float64) {
	if threat == 0 || enemy == "" {
		return
	}
	by := a.threatBy[player]
	if by == nil {
		by = map[string]float64{}
		a.threatBy[player] = by
	}
	by[enemy] += threat
}

// engage marks a hostile unit as in the fight now. flags is the reaction
// the current event carries for guid, not a registry lookup: the registry's
// flags for a GUID are whatever the most recent event touching it said,
// which mid-stream can still be catching up to what a later event will
// reveal (Snapshot's doc comment promises a live-tail caller correct
// output at any point, so this must not depend on how much of the fight
// the registry has seen yet). The event in hand is already authoritative
// for its own units.
func (a *Accumulator) engage(guid string, flags uint32, at time.Time) {
	if guid == "" || !units.Hostile(flags) {
		return
	}
	a.engaged[guid] = at
}

// spreadThreat books healing threat over every hostile unit engaged within the window.
func (a *Accumulator) spreadThreat(player string, threat float64, at time.Time) {
	if threat == 0 {
		return
	}
	var live []string
	for guid, last := range a.engaged {
		if at.Sub(last) <= a.opt.EngagedWindow {
			live = append(live, guid)
		}
	}
	if len(live) == 0 {
		return
	}
	each := threat / float64(len(live))
	for _, guid := range live {
		a.creditThreat(player, guid, each)
	}
}

// threatPairs renders the per-target table, largest first.
func (a *Accumulator) threatPairs() []ThreatPair {
	out := []ThreatPair{}
	for player, by := range a.threatBy {
		for enemy, threat := range by {
			out = append(out, ThreatPair{GUID: player, Name: a.name(player), TargetGUID: enemy, TargetName: a.name(enemy), Threat: threat})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Threat != out[j].Threat {
			return out[i].Threat > out[j].Threat
		}
		return out[i].GUID+out[i].TargetGUID < out[j].GUID+out[j].TargetGUID
	})
	return out
}
