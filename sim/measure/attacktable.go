package measure

import (
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// AttackTableRow is the observed one-roll attack table against one
// target. With a known character sheet it is what settles how Forever's
// unified Hit stat interacts with the vanilla weapon-skill miss table,
// which research 5.3 records as completely unspecified by anything
// public and load-bearing for every melee spec.
type AttackTableRow struct {
	// AttackType is "melee", "ranged" or "special": the three tables
	// vanilla rolls separately.
	AttackType  string `json:"attack_type"`
	TargetName  string `json:"target_name"`
	TargetLevel int64  `json:"target_level"`

	Swings int `json:"swings"`

	Miss   float64 `json:"miss"`
	Dodge  float64 `json:"dodge"`
	Parry  float64 `json:"parry"`
	Glance float64 `json:"glance"`
	Crit   float64 `json:"crit"`
	Block  float64 `json:"block"`

	Enough bool `json:"enough"`
}

type tableKey struct {
	attackType string
	target     string
	level      int64
}

// MeasureAttackTable counts every outcome of every swing against each
// target and divides by the total. A rate is a fraction of all swings,
// including misses, because that is what the one-roll table produces.
func MeasureAttackTable(in Input) []AttackTableRow {
	type acc struct {
		row                                     AttackTableRow
		miss, dodge, parry, glance, crit, block int
	}
	rows := map[tableKey]*acc{}
	var order []tableKey

	// The target's level is NOT on the swing: a line the player deals
	// carries the PLAYER's advanced block, and a SWING_MISSED carries no
	// block at all. It is on the line whose block describes the target -
	// SWING_DAMAGE_LANDED - so the levels are gathered once, by creature
	// GUID, and looked up by the swing's DEST guid. Keying off the
	// swing's own block instead would file every row under the player's
	// level (or, for the misses, under 0) and split one target in two.
	levels := TargetLevels(in.Events)
	// Which spells are physical, so that a spell-prefixed line can be
	// told from a melee special. See attackTypeOf.
	physical := physicalSpells(in.Events)

	for _, e := range in.Events {
		if !byActor(e, in.Actor) {
			continue
		}
		at := attackTypeOf(e, physical)
		if at == "" {
			continue
		}
		key := tableKey{attackType: at, target: e.Dest.Name, level: levels[e.Dest.GUID]}
		a, ok := rows[key]
		if !ok {
			a = &acc{row: AttackTableRow{
				AttackType: at, TargetName: e.Dest.Name, TargetLevel: key.level,
			}}
			rows[key] = a
			order = append(order, key)
		}
		a.row.Swings++
		switch {
		case e.Kind == event.Missed:
			switch strings.ToUpper(e.MissType) {
			case "MISS":
				a.miss++
			case "DODGE":
				a.dodge++
			case "PARRY":
				a.parry++
			case "BLOCK":
				a.block++
			}
		case e.Critical.OK && e.Critical.V:
			a.crit++
		case e.Glancing.OK && e.Glancing.V:
			a.glance++
		}
		// A blocked hit still lands, so Blocked is counted from the
		// damage event's own field rather than from a miss type.
		if e.Blocked.OK && e.Blocked.V > 0 {
			a.block++
		}
	}

	out := make([]AttackTableRow, 0, len(order))
	for _, key := range order {
		a := rows[key]
		n := float64(a.row.Swings)
		if n > 0 {
			a.row.Miss = float64(a.miss) / n
			a.row.Dodge = float64(a.dodge) / n
			a.row.Parry = float64(a.parry) / n
			a.row.Glance = float64(a.glance) / n
			a.row.Crit = float64(a.crit) / n
			a.row.Block = float64(a.block) / n
		}
		a.row.Enough = a.row.Swings >= in.MinSamples
		out = append(out, a.row)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Swings > out[j].Swings })
	return out
}

// physicalSpells is the set of spell ids the log shows dealing physical
// damage. A spell miss carries no school - the _MISSED suffix is
// missType and isOffHand and nothing else - so the only way to know
// which table a missed ability rolled on is what its own damage lines
// said elsewhere in the same log.
func physicalSpells(events []event.Event) map[int64]bool {
	out := map[int64]bool{}
	for _, e := range events {
		if e.Kind != event.Damage || e.Spell.ID == 0 || !e.School.OK {
			continue
		}
		if e.School.V == physicalSchool {
			out[e.Spell.ID] = true
		}
	}
	return out
}

// attackTypeOf classifies an event into one of vanilla's three attack
// tables, or "" for anything that does not roll on one.
//
// A spell-prefixed line is a melee special only when it is PHYSICAL: the
// one-roll table with its dodge, parry and glance outcomes is the
// weapon table, and a Frostbolt rolls on the spell hit table instead.
// Mixing the two would report a dodge rate diluted by every spell cast
// in the log. A spell MISS states no school, so it is placed by what the
// same spell's damage lines said (physicalSpells) and dropped when the
// log never says.
//
// SWING_DAMAGE_LANDED is deliberately absent. It restates the
// SWING_DAMAGE that preceded it, and counting both would double every
// landed hit while leaving the misses alone, which would halve every
// observed miss rate. logs/engine/summary hit the same trap.
func attackTypeOf(e event.Event, physical map[int64]bool) string {
	switch e.Name {
	case "SWING_DAMAGE", "SWING_MISSED":
		return "melee"
	case "RANGE_DAMAGE", "RANGE_MISSED":
		return "ranged"
	case "SPELL_DAMAGE":
		if e.School.OK && e.School.V != physicalSchool {
			return ""
		}
		return "special"
	case "SPELL_MISSED":
		if !physical[e.Spell.ID] {
			return ""
		}
		return "special"
	}
	return ""
}
