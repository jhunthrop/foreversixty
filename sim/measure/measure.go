// Package measure recovers combat constants from a real combat log.
//
// The Classic-lineage client tables give item stats, spell attributes,
// cooldowns, costs, durations, talent grids and base stats. They do not
// give whether periodic damage crits (a server rule, not a per-spell
// field), the multiplier a critical tick uses, how a unified Hit stat
// interacts with the vanilla weapon-skill miss table, proc chances, or
// internal cooldowns. The rating conversion tables are game-table files
// inside the client package rather than DB2, and wago does not serve
// them. Every one of those is visible in a combat log.
//
// DIALECT. Forever writes combat-log version 22 with advanced logging on
// and PROJECT_ID 18, which logs/engine/layout.RetailV22() reads; the
// version alone selects the row. It is NOT the Classic dialect. The v22
// advanced block is nineteen fields and rides on every damage, heal,
// energize and cast-success line, carrying the acting unit's attack
// power, spell power, armour, power and level - so the character sheet a
// measurement is interpreted against comes out of the log itself rather
// than off a screenshot.
//
// The measurement functions return values and the command prints them,
// because this same code is the first input to the nightly validation
// job, which needs the numbers rather than a table on stdout.
//
// This package imports logs/engine read-only and never modifies it.
package measure

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
)

// DefaultMinSamples is the floor below which a figure prints as
// "insufficient data" rather than as a number. A proc rate from four
// swings is worse than no proc rate: it lands in a constants file and
// nobody re-checks it.
const DefaultMinSamples = 30

// Input is one measurement run.
type Input struct {
	// Events is the decoded log, from Load.
	Events []event.Event
	// Actor is the player whose actions are measured, by name. Every
	// figure is about this one unit.
	Actor string
	// SpellPower and AttackPower OVERRIDE what the log's advanced blocks
	// report for the actor. Leave them zero and the sheet is read from
	// the log, which is what a v22 log makes possible and what the
	// command does by default. Set them only when the log's own figures
	// are wrong or absent.
	SpellPower  float64
	AttackPower float64
	// MinSamples overrides DefaultMinSamples.
	MinSamples int
}

// Sheet is the actor's character sheet as the log reports it. Every
// figure is the median across the actor's own advanced blocks, not the
// mean, because a single buffed or debuffed line should not move it.
type Sheet struct {
	SpellPower  float64 `json:"spell_power"`
	AttackPower float64 `json:"attack_power"`
	Armor       float64 `json:"armor"`
	MaxHP       float64 `json:"max_hp"`
	Level       int64   `json:"level"`
	// Samples is how many advanced blocks the sheet was read from. Zero
	// means the log had advanced logging off and every figure that needs
	// a sheet is unavailable.
	Samples int `json:"samples"`
}

// Report is everything one log can say.
type Report struct {
	Actor  string `json:"actor"`
	Events int    `json:"events"`
	// FightCount counts ENCOUNTER_START events. An open-world log has
	// none, and that is not an error: it is why Incomplete lists what is
	// missing. It is deliberately NOT called Fights: Task 15B adds
	// Fights []Fight to this same struct, and one name for a count and a
	// list is how a refactor silently drops one of them.
	FightCount int   `json:"fight_count"`
	Sheet      Sheet `json:"sheet"`
	// Incomplete lists what this log could NOT support, in words, so the
	// printed table says "no level-63 target in this log" rather than
	// showing an empty attack table and letting the reader guess. The
	// field is named for what it holds: a field called Complete carrying
	// the incompleteness reads backwards.
	Incomplete []string `json:"incomplete,omitempty"`

	Periodic     []PeriodicCrit   `json:"periodic"`
	AttackTable  []AttackTableRow `json:"attack_table"`
	Procs        []ProcRate       `json:"procs"`
	Coefficients []Coefficient    `json:"coefficients"`
}

// Load reads a combat log and returns its decoded events. base seeds the
// clock for dialects whose timestamps carry no year: pass the file's
// modification time. Forever's own stamps carry the year, so base only
// matters for a log from another client.
func Load(path string, base time.Time) ([]event.Event, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("measure: %w", err)
	}
	s := session.New(session.Options{
		ReportID: "measure",
		// Forever writes combat-log version 22 with PROJECT_ID 18, which
		// this row reads; retail writes PROJECT_ID 1 and the same
		// dialect, so the version alone selects it. Infer is on so a log
		// with a damaged or missing header still parses by field count
		// rather than silently falling back to the v16 row.
		Layout: layout.RetailV22(),
		Infer:  true,
		Base:   base,
	})
	res, err := s.Feed(b, 0)
	if err != nil {
		return nil, fmt.Errorf("measure: feeding %s: %w", path, err)
	}
	last, err := s.Close()
	if err != nil {
		return nil, fmt.Errorf("measure: closing %s: %w", path, err)
	}
	out := make([]event.Event, 0, len(res.Events)+len(last.Events))
	out = append(out, res.Events...)
	out = append(out, last.Events...)
	return out, nil
}

// SheetFromEvents reads the actor's character sheet out of its own
// advanced blocks. A line the actor acted on carries the ACTOR's block;
// a line aimed at the actor carries the actor's block only on the
// _LANDED half, so the InfoGUID is matched rather than the source name.
func SheetFromEvents(events []event.Event, actor string) Sheet {
	var guid string
	for _, e := range events {
		if e.Source.Name == actor && e.Source.GUID != "" {
			guid = e.Source.GUID
			break
		}
	}
	if guid == "" {
		return Sheet{}
	}
	var sp, ap, armor, hp []float64
	var level int64
	for _, e := range events {
		if !e.Adv.OK || e.Adv.InfoGUID != guid {
			continue
		}
		sp = append(sp, float64(e.Adv.SpellPower))
		ap = append(ap, float64(e.Adv.AttackPower))
		armor = append(armor, float64(e.Adv.Armor))
		hp = append(hp, float64(e.Adv.MaxHP))
		level = e.Adv.Level
	}
	return Sheet{
		SpellPower:  median(sp),
		AttackPower: median(ap),
		Armor:       median(armor),
		MaxHP:       median(hp),
		// For a player the advanced block's Level field is item level,
		// not character level: the log writes one field with two
		// meanings. It is reported as-is and the table labels it.
		Level:   level,
		Samples: len(sp),
	}
}

// TargetLevels maps each creature GUID to the level its advanced blocks
// report. Only creatures: for a player that field is item level.
//
// A line the player deals carries the PLAYER's advanced block, so a
// creature's level comes from the lines where the creature is the acting
// or landed-on unit - in practice SWING_DAMAGE_LANDED and the creature's
// own swings. That is why a log with advanced logging off cannot key an
// attack table by level at all.
func TargetLevels(events []event.Event) map[string]int64 {
	out := map[string]int64{}
	for _, e := range events {
		if !e.Adv.OK || e.Adv.Level == 0 {
			continue
		}
		if !strings.HasPrefix(e.Adv.InfoGUID, "Creature-") && !strings.HasPrefix(e.Adv.InfoGUID, "Vehicle-") {
			continue
		}
		out[e.Adv.InfoGUID] = e.Adv.Level
	}
	return out
}

// Run measures everything this package knows how to measure.
func Run(in Input) (Report, error) {
	if len(in.Events) == 0 {
		return Report{}, errors.New("measure: no events; load a log first")
	}
	if in.Actor == "" {
		return Report{}, errors.New("measure: Actor is required; every figure is about one unit")
	}
	if in.MinSamples <= 0 {
		in.MinSamples = DefaultMinSamples
	}

	rep := Report{Actor: in.Actor, Events: len(in.Events)}
	rep.Sheet = SheetFromEvents(in.Events, in.Actor)
	// A flag overrides the log; the log fills in what no flag gave.
	if in.SpellPower == 0 {
		in.SpellPower = rep.Sheet.SpellPower
	}
	if in.AttackPower == 0 {
		in.AttackPower = rep.Sheet.AttackPower
	}
	for _, e := range in.Events {
		if e.Name == "ENCOUNTER_START" {
			rep.FightCount++
		}
	}

	rep.Periodic = MeasurePeriodic(in)
	rep.AttackTable = MeasureAttackTable(in)
	rep.Procs = MeasureProcs(in)
	rep.Coefficients = MeasureCoefficients(in)

	// Say what the log could not support, rather than printing an empty
	// table and letting the reader guess whether the tool is broken.
	if rep.Sheet.Samples == 0 {
		rep.Incomplete = append(rep.Incomplete,
			"no advanced blocks for this actor: advanced logging was off, or the name is wrong. Coefficients and the character sheet are unavailable.")
	}
	if rep.FightCount == 0 {
		rep.Incomplete = append(rep.Incomplete,
			"no ENCOUNTER_START in this log: it is open-world, so there is no COMBATANT_INFO, no per-fight segmentation, and nothing that needs the character's gear or talents at a pull.")
	}
	var sawBossLevel bool
	for _, row := range rep.AttackTable {
		if row.TargetLevel >= 63 {
			sawBossLevel = true
		}
	}
	if !sawBossLevel {
		rep.Incomplete = append(rep.Incomplete,
			"no level-63 target: the boss-level attack table, and therefore the unified-Hit fit, need an instance log.")
	}
	return rep, nil
}

// byActor reports whether an event's source is the measured actor.
func byActor(e event.Event, actor string) bool {
	return e.Source.Name == actor
}

// isPeriodic reports whether an event is a damage tick rather than a
// direct hit. The engine normalises both to event.Damage, so the raw
// name is what distinguishes them.
func isPeriodic(e event.Event) bool {
	return e.Name == "SPELL_PERIODIC_DAMAGE"
}

// physicalSchool is the combat-log school mask for physical damage.
const physicalSchool = 1

func mean(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func median(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	s := append([]float64(nil), xs...)
	sort.Float64s(s)
	n := len(s)
	if n%2 == 1 {
		return s[n/2]
	}
	return (s[n/2-1] + s[n/2]) / 2
}

func stddev(xs []float64) float64 {
	if len(xs) < 2 {
		return 0
	}
	m := mean(xs)
	var sum float64
	for _, x := range xs {
		d := x - m
		sum += d * d
	}
	return math.Sqrt(sum / float64(len(xs)))
}
