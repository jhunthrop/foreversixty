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
	// PrePull is true for a taunt the fight holds only as its debuff landing:
	// the cast came before the pull's first event, so the Casts tab, which
	// counts casts, is one short of this list and both are right.
	PrePull bool `json:"pre_pull,omitempty"`
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

// tauntEchoMS is how long after a taunt's cast its debuff can land and still be
// the same taunt: the two lines arrive a few milliseconds apart in the log. The
// debuff can carry its own spell id (Provoke's cast is 115546, its debuff
// 116189), so the echo is matched by who taunted what, not by which id.
const tauntEchoMS = 100

// noteTaunt records a taunt cast at an enemy. Called from Add for every event.
// A taunt with no destination is no taunt this view can show: the nil GUID
// renders as "Environment", so it is rejected alongside the empty one.
//
// The cast is the record; the debuff it applies is the fallback. A tank who
// taunts the boss on the pull casts before the pull's first event, so the
// fight holds the debuff landing and not the cast, and the taunt that opened
// the fight would otherwise be the one taunt the list never shows.
func (a *Accumulator) noteTaunt(e event.Event) {
	if !tauntSpells[e.Spell.ID] || e.Dest.GUID == "" || e.Dest.GUID == units.NoGUID {
		return
	}
	key := tauntKey{e.Source.GUID, e.Dest.GUID}
	at := a.ms(e.Time)
	prePull := false
	switch e.Kind {
	case event.CastSuccess:
		if a.tauntCasts == nil {
			a.tauntCasts = map[tauntKey]int64{}
		}
		a.tauntCasts[key] = at
	case event.AuraApplied:
		if last, ok := a.tauntCasts[key]; ok && at-last <= tauntEchoMS {
			return
		}
		prePull = true
	default:
		return
	}
	a.taunts = append(a.taunts, Taunt{
		AtMS: at, SourceGUID: e.Source.GUID, SourceName: a.name(e.Source.GUID),
		TargetGUID: e.Dest.GUID, TargetName: a.name(e.Dest.GUID), SpellID: e.Spell.ID, SpellName: e.Spell.Name,
		PrePull: prePull,
	})
}

// tauntKey names one unit taunting one enemy, so a debuff can be matched to
// the cast that applied it whichever id the debuff carries.
type tauntKey struct {
	source, target string
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
