package coach

// F3 fixtures: a player whose damage series and roster active time agree, so
// ratingengine.IdleWindows (per-second spans) and ActivityIdleMS (time alive less the
// roster's ActiveMS when no window is in force) measure the same dead time.

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	idleBossGUID      = "Creature-0-1-1084-1-10184-0000000001"
	idleDowntimeSpell = 99_901
	secondMS          = 1_000
)

// idleOf is in's player fighting for fightMS: one damage point every second, none inside
// the idle spans, a cast at the pull, no deaths, and the roster's ActiveMS equal to the
// seconds with damage.
func idleOf(t *testing.T, in Input, fightMS int64, idle ...summary.Segment) Input {
	t.Helper()
	guid := in.Row.GUID
	f := withDuration(in.Fight, fightMS)
	f = withCasts(f, guid, 23881, "Bloodthirst", 0)
	f.Deaths = slices.DeleteFunc(slices.Clone(f.Deaths), func(d summary.Death) bool { return d.GUID == guid })
	series, active := damageSeries(fightMS, idle)
	f.DamageDone = append(slices.DeleteFunc(slices.Clone(f.DamageDone), func(a summary.Actor) bool { return a.GUID == guid }),
		summary.Actor{GUID: guid, ActiveMS: active, Series: series})
	f.Healing = slices.DeleteFunc(slices.Clone(f.Healing), func(a summary.Actor) bool { return a.GUID == guid })
	f.Roster = slices.Clone(f.Roster)
	for i := range f.Roster {
		if f.Roster[i].GUID == guid {
			f.Roster[i].ActiveMS = active
			f.Roster[i].Deaths = 0
		}
	}
	in.Fight = f
	in.Row = rosterRowOf(t, f, guid)
	return in
}

// damageSeries is one 1-damage reading per second of fightMS, 0 inside the idle spans,
// and the active milliseconds (the seconds with damage).
func damageSeries(fightMS int64, idle []summary.Segment) ([]int64, int64) {
	series := repeat(1, int(fightMS/secondMS))
	for _, s := range idle {
		for sec := s.StartMS / secondMS; sec < s.EndMS/secondMS; sec++ {
			series[sec] = 0
		}
	}
	var active int64
	for _, v := range series {
		if v != 0 {
			active += secondMS
		}
	}
	return series, active
}

func furyIdle(t *testing.T, fightMS int64, idle ...summary.Segment) Input {
	t.Helper()
	return idleOf(t, furyInput(t), fightMS, idle...)
}

func protIdle(t *testing.T, fightMS int64, idle ...summary.Segment) Input {
	t.Helper()
	return idleOf(t, protInput(t), fightMS, idle...)
}

// holyIdle is a healer with the same gap; F3 must still ignore her.
func holyIdle(t *testing.T, fightMS int64, idle ...summary.Segment) Input {
	t.Helper()
	return idleOf(t, holyInput(t), fightMS, idle...)
}

// furyIdleInsideDowntime idles 40 s to 52 s, exactly the 12 s downtime a boss cast at
// 40 s opens.
func furyIdleInsideDowntime(t *testing.T) Input {
	t.Helper()
	in := furyIdle(t, 200_000, span(40_000, 52_000))
	in.Fight.Casts = append(slices.Clone(in.Fight.Casts), summary.CastRow{GUID: idleBossGUID, SpellID: idleDowntimeSpell,
		SpellName: "Fixture Downtime", Started: 1, Succeeded: 1, Sequence: []int64{40_000}})
	in.Tables.Mechanics = &mechanics.Table{EncounterID: onyxiaEncounterID, Name: onyxiaName, Downtime: []mechanics.Downtime{{
		Trigger:    mechanics.PhaseStart{SpellID: idleDowntimeSpell, On: mechanics.OnCastSuccess},
		DurationMS: 12_000, Note: "fixture",
	}}}
	return in
}

// withAssignment is in with an officer's mark on the player over [fromMS, toMS).
func withAssignment(in Input, fromMS, toMS int64) Input {
	in.Tables.Assignments = append(slices.Clone(in.Tables.Assignments),
		ratingengine.Assignment{PlayerKey: in.Row.GUID, FromMS: fromMS, ToMS: toMS})
	return in
}

// withFirstCast is in with the player's only cast at atMS.
func withFirstCast(in Input, atMS int64) Input {
	in.Fight = withCasts(withoutCasts(in.Fight, in.Row.GUID), in.Row.GUID, 23881, "Bloodthirst", atMS)
	return in
}

// idlePool is a pool of n players whose median is active percent of counted time.
func idlePool(active float64, n int) Refs {
	return poolOf(n, Profile{ActivePct: &active})
}
