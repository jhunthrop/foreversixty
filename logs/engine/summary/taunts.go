// logs/engine/summary/taunts.go
package summary

import (
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

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

// tauntSpells are the taunts the deaths and threat views mark: single-target
// taunts only, each one cast at the enemy it pulls. Data awaiting Forever's own
// spell ids; the retail sample's Hand of Reckoning and Provoke are here so the
// feature can be reviewed against it.
//
// Deliberately absent until Forever's ids land, because a Taunt row names a
// target and these name the wrong one: Challenging Shout (1161) and Challenging
// Roar (5209) are untargeted area taunts, and Righteous Defense (31789) is cast
// on the ally whose attackers it pulls, so its Dest is a raider, not an enemy.
var tauntSpells = map[int64]bool{
	355: true, 694: true, 6795: true, 17735: true,
	56222: true, 62124: true, 115546: true, 116189: true, 185245: true,
}

// noteTaunt records a taunt cast at an enemy. Called from Add for every event.
// A taunt with no destination is no taunt this view can show: the nil GUID
// renders as "Environment", so it is rejected alongside the empty one.
func (a *Accumulator) noteTaunt(e event.Event) {
	if e.Kind != event.CastSuccess || !tauntSpells[e.Spell.ID] ||
		e.Dest.GUID == "" || e.Dest.GUID == units.NoGUID {
		return
	}
	a.taunts = append(a.taunts, Taunt{
		AtMS: a.ms(e.Time), SourceGUID: e.Source.GUID, SourceName: a.name(e.Source.GUID),
		TargetGUID: e.Dest.GUID, TargetName: a.name(e.Dest.GUID), SpellID: e.Spell.ID, SpellName: e.Spell.Name,
	})
}

// tauntRows renders the taunt list, in the order the casts arrived. Like
// threatPairs it is empty rather than nil when there was nothing: `taunts: []`
// in the json says the parse kept taunts and found none, where a missing key
// says the report was parsed before taunts existed, and the report reads the
// two differently.
func (a *Accumulator) tauntRows() []Taunt {
	out := make([]Taunt, 0, len(a.taunts))
	return append(out, a.taunts...)
}
