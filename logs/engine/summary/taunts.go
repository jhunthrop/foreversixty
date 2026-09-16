// logs/engine/summary/taunts.go
package summary

import "github.com/jhunthrop/foreversixty/logs/engine/event"

// Taunt is one taunt cast: who taunted what, when.
type Taunt struct {
	AtMS       int64  `json:"at_ms"`
	SourceGUID string `json:"source_guid"`
	SourceName string `json:"source_name"`
	TargetGUID string `json:"target_guid"`
	TargetName string `json:"target_name"`
	SpellID    int64  `json:"spell_id"`
	SpellName  string `json:"spell_name"`
}

// tauntSpells are the taunts the deaths and threat views mark. Data awaiting
// Forever's own spell ids; the retail sample's Hand of Reckoning and Provoke are
// here so the feature can be reviewed against it.
var tauntSpells = map[int64]bool{
	355: true, 694: true, 1161: true, 5209: true, 6795: true, 17735: true,
	31789: true, 56222: true, 62124: true, 115546: true, 116189: true, 185245: true,
}

// noteTaunt records a taunt cast at an enemy. Called from Add for every event.
func (a *Accumulator) noteTaunt(e event.Event) {
	if e.Kind != event.CastSuccess || !tauntSpells[e.Spell.ID] || e.Dest.GUID == "" {
		return
	}
	a.taunts = append(a.taunts, Taunt{
		AtMS: a.ms(e.Time), SourceGUID: e.Source.GUID, SourceName: a.name(e.Source.GUID),
		TargetGUID: e.Dest.GUID, TargetName: a.name(e.Dest.GUID), SpellID: e.Spell.ID, SpellName: e.Spell.Name,
	})
}
