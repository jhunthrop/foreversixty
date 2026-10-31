// Package engine is the API's one way of calling the logs engine. Every
// parse the API runs - the whole-file job, the ingest's verification of a
// companion bundle, the raw-sample check - goes through the options and
// the rebuild here, so all three agree on what the same bytes mean.
//
// The companion must build its sessions the same way, or verification
// would reject honest bundles: same fight.DefaultOptions, same
// summary.DefaultOptions, the engine's own retail spec tables, and no
// consumable or raid-buff tables (those change the summary, never a
// metrics row).
package engine

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Version is the engine version the API records on every report and fight.
const Version = session.Version

// classSpells is the table tools/gen-class-spells.mjs cuts from a client build: the class that
// owns each spell id, and the spec a talent's spell id defines.
//
//go:embed classspells.json
var classSpellsJSON []byte

var classSpells = func() (t struct {
	Classes map[string][]int64            `json:"classes"`
	Spec    map[string]map[string][]int64 `json:"spec"`
}) {
	if err := json.Unmarshal(classSpellsJSON, &t); err != nil {
		panic("engine: classspells.json: " + err.Error())
	}
	return t
}()

// ClassBySpell maps a spell id to the one class that owns it, for a log whose players carry no
// COMBATANT_INFO: a player who cast Rend is a Warrior. Spells shared by classes are absent.
func ClassBySpell() map[int64]string { return classBySpell }

var classBySpell = func() map[int64]string {
	out := map[int64]string{}
	for class, ids := range classSpells.Classes {
		for _, id := range ids {
			out[id] = class
		}
	}
	return out
}()

// SpecBySpell maps a spec-defining spell id (a talent's) to its spec's name.
func SpecBySpell() map[int64]string { return specBySpell }

var specBySpell = func() map[int64]string {
	out := map[int64]string{}
	for _, specs := range classSpells.Spec {
		for spec, ids := range specs {
			for _, id := range ids {
				out[id] = spec
			}
		}
	}
	return out
}()

// UnitOptions is the registry configuration shared by every parse. COMBATANT_INFO names a class
// and spec when the log has it; otherwise the class and spec are read off the spells cast.
func UnitOptions() units.Options {
	return units.Options{
		ClassBySpec:  units.RetailSpecClass,
		ClassBySpell: ClassBySpell(),
		SpecBySpell:  SpecBySpell(),
	}
}

// SummaryOptions is the accumulator configuration shared by every parse.
// reg is required: the summary resolves names and pet owners through it.
func SummaryOptions(reg *units.Registry) summary.Options {
	o := summary.DefaultOptions()
	o.Registry = reg
	o.SpecNames = units.RetailSpecName
	return o
}

// SessionOptions is the session configuration for a whole-file parse.
// base seeds the clock for dialects whose timestamps carry no year; pass
// the upload's creation time. keepEvents must be true for anything that
// writes Parquet.
func SessionOptions(reportID string, base time.Time, keepEvents bool) session.Options {
	o := session.Options{
		ReportID:   reportID,
		Base:       base,
		Infer:      true,
		KeepEvents: keepEvents,
		Units:      UnitOptions(),
		Fight:      fight.DefaultOptions(),
	}
	o.Summary = SummaryOptions(nil)
	return o
}

// LayoutByName finds a registered layout row. A report records the
// layout its parse settled on, so a later parse of part of the same log
// - the raw-sample check, which starts mid-file and so sees no header -
// can be told which dialect to read without guessing.
func LayoutByName(name string) (layout.Layout, bool) {
	for _, r := range layout.Rows() {
		if r.Name == name {
			return r, true
		}
	}
	return layout.Layout{}, false
}

// SessionOptionsFor is SessionOptions with the layout pinned by name.
// An unknown or empty name leaves the session to infer, as before.
func SessionOptionsFor(reportID string, base time.Time, keepEvents bool, layoutName string) session.Options {
	o := SessionOptions(reportID, base, keepEvents)
	if l, ok := LayoutByName(layoutName); ok {
		o.Layout = l
	}
	return o
}

// ErrNoEvents is returned by Rebuild when the Parquet holds no events at
// all, which no honest bundle ever does.
var ErrNoEvents = errors.New("engine: the events are empty")

// Header is the fight's identity. Rebuild cannot read it from the events:
// the Parquet schema keeps every event's kind and timing but not the
// ENCOUNTER_START payload, so an encounter's id, difficulty, size, name,
// and kill flag do not survive a write-and-read round trip.
//
// The caller therefore passes the header the uploader claimed, and
// verification checks everything the events themselves decide - damage,
// healing, damage taken, active time, deaths, duration, who was present,
// and each player's role - rather than the header.
//
// FOLLOW-UP (engine): give parquet.Row columns for the Encounter, Zone,
// and Combatant payloads. Rebuild can then take the header, the item
// levels, and the specs from the file, and verification can check them
// too. Until then the ingest trusts the header and the raw-sample job,
// which re-parses original log text, is what catches a forged one.
type Header struct {
	EncounterID int64
	Name        string
	Difficulty  int64
	Size        int64
	Kill        bool
	Zone        string
}

// Rebuild replays one fight's events - the Parquet a companion uploaded -
// through a fresh accumulator and returns the fight and summary the
// engine derives from them.
//
// The fight is assembled here rather than through fight.Segmenter on
// purpose: the segmenter decides encounter boundaries from
// ENCOUNTER_START payloads and trash boundaries from combat gaps, and
// neither survives the Parquet round trip, so a segmenter fed these
// events would split one honest fight into several. Everything it would
// derive from the events - start, end, the players present, the group,
// player and mob deaths, the group's kills - is derived the same way here,
// through the segmenter's own fight.Counter.
func Rebuild(index int, h Header, events []event.Event) (fight.Fight, summary.Summary, error) {
	if len(events) == 0 {
		return fight.Fight{}, summary.Summary{}, ErrNoEvents
	}
	f := fight.Fight{
		Index: index, Kind: fight.Encounter,
		EncounterID: h.EncounterID, Name: h.Name,
		Difficulty: h.Difficulty, Size: h.Size, Kill: h.Kill, Zone: h.Zone,
		Start: events[0].Time, End: events[0].Time,
	}
	if h.EncounterID == 0 {
		f.Kind, f.Name = fight.Trash, "Trash"
	}
	reg := units.NewRegistry(UnitOptions())
	acc := summary.New(SummaryOptions(reg))
	acc.Start(f.Start)
	counts := fight.NewCounter(fight.DefaultOptions())
	for _, e := range events {
		if e.Time.After(f.End) {
			f.End = e.Time
		}
		reg.Observe(e)
		acc.Add(e)
		counts.Observe(e)
	}
	counts.Apply(&f)
	f.StartOffset, f.EndOffset = events[0].Offset, events[len(events)-1].Offset
	f.StartLine, f.EndLine = events[0].Line, events[len(events)-1].Line
	return f, acc.Snapshot(f, Version), nil
}

// NameFromEvents picks a display name for a rebuilt fight. The
// ENCOUNTER_START payload with the boss's name does not survive the
// Parquet round trip (see Header), so the name is taken from the fight
// itself: the hostile unit that took the most damage, which on a boss
// fight is the boss.
//
// FOLLOW-UP (engine): once parquet.Row carries the Encounter payload,
// the real name travels with the events and this guess goes away. The
// whole-file parse job already writes the real name, because it reads
// the log text.
func NameFromEvents(h Header, s summary.Summary) string {
	if h.Name != "" {
		return h.Name
	}
	if h.EncounterID == 0 {
		return "Trash"
	}
	best := ""
	var most int64
	for _, a := range s.DamageTaken {
		if units.Parse(a.GUID).Kind == units.KindPlayer || a.Name == "" {
			continue
		}
		if a.Effective > most {
			best, most = a.Name, a.Effective
		}
	}
	if best != "" {
		return best
	}
	return "Encounter " + strconv.FormatInt(h.EncounterID, 10)
}
