// logs/engine/summary/threat.go
package summary

import "github.com/jhunthrop/foreversixty/logs/engine/event"

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
