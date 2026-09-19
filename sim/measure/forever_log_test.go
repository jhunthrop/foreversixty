package measure

import (
	"strings"
	"testing"
	"time"
)

// foreverExcerpt is the committed excerpt of the first Forever beta log:
// open world, no instance, therefore no COMBATANT_INFO and no
// ENCOUNTER_START. It is read here through the same Load the tool uses,
// so a dialect regression in logs/engine fails this package too.
const foreverExcerpt = "../../logs/engine/event/testdata/forever-1.60.log"

var foreverBase = time.Date(2026, 9, 17, 0, 0, 0, 0, time.UTC)

func TestTheRealForeverLogParsesUnderRetailV22(t *testing.T) {
	events, err := Load(foreverExcerpt, foreverBase)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) < 50 {
		t.Fatalf("the excerpt produced %d events, want at least 50", len(events))
	}
	var parseErrors, unknown, advanced int
	for _, e := range events {
		switch e.Kind.String() {
		case "parse_error":
			parseErrors++
			t.Logf("parse error on line %d: %s", e.Line, e.Raw)
		case "unknown":
			unknown++
			t.Logf("unknown event %q on line %d", e.Name, e.Line)
		}
		if e.Adv.OK {
			advanced++
		}
	}
	if parseErrors != 0 {
		t.Errorf("%d parse errors in the real Forever log; the dialect is wrong", parseErrors)
	}
	if unknown != 0 {
		t.Errorf("%d unknown events in the real Forever log", unknown)
	}
	if advanced == 0 {
		t.Error("no advanced blocks decoded; the 19-field block is being read as suffix params")
	}
}

// The advanced block is what makes the character sheet readable without a
// screenshot. On a real line it carries attack power, spell power, armour
// and the unit's level.
func TestTheRealLogCarriesACharacterSheet(t *testing.T) {
	events, err := Load(foreverExcerpt, foreverBase)
	if err != nil {
		t.Fatal(err)
	}
	sheet := SheetFromEvents(events, "Tester02-Beta")
	if sheet.Samples == 0 {
		t.Fatal("no advanced blocks found for Tester02-Beta; the actor match or the block offset is wrong")
	}
	if sheet.MaxHP <= 0 {
		t.Errorf("Sheet.MaxHP = %v, want a positive figure from the advanced block", sheet.MaxHP)
	}
}

// Creature levels come out of the advanced blocks, and only creatures':
// for a player that field is item level.
func TestTheRealLogCarriesCreatureLevels(t *testing.T) {
	events, err := Load(foreverExcerpt, foreverBase)
	if err != nil {
		t.Fatal(err)
	}
	levels := TargetLevels(events)
	if len(levels) == 0 {
		t.Fatal("no creature levels found")
	}
	for guid, lvl := range levels {
		if !strings.HasPrefix(guid, "Creature-") && !strings.HasPrefix(guid, "Vehicle-") {
			t.Errorf("%q is not a creature guid but has a level", guid)
		}
		if lvl <= 0 || lvl > 100 {
			t.Errorf("%s has level %d", guid, lvl)
		}
	}
}

// The header restart is real and it is in this excerpt: the client wrote
// one header when it loaded with ADVANCED_LOG_ENABLED,0 and a second
// eleven seconds later when /combatlog began with it on. A tool that
// read the flag off the first line would call this log unadvanced, so
// nothing here reads the header: the test is Adv.OK, per event.
func TestTheRealLogSurvivesItsHeaderRestart(t *testing.T) {
	events, err := Load(foreverExcerpt, foreverBase)
	if err != nil {
		t.Fatal(err)
	}
	var headers, advanced int
	for _, e := range events {
		if e.Kind.String() == "header" {
			headers++
		}
		if e.Adv.OK {
			advanced++
		}
	}
	if headers < 2 {
		t.Fatalf("%d header lines; the excerpt carries two, the load header and the /combatlog header", headers)
	}
	if advanced == 0 {
		t.Error("no advanced blocks after the header restart; the layout was not re-selected")
	}
}

// This is the honest half of the task: the open-world log cannot support
// the boss attack table or anything keyed to COMBATANT_INFO, and the
// report has to say so rather than printing an empty table.
func TestTheRealLogReportsWhatItCannotMeasure(t *testing.T) {
	events, err := Load(foreverExcerpt, foreverBase)
	if err != nil {
		t.Fatal(err)
	}
	rep, err := Run(Input{Events: events, Actor: "Tester02-Beta", MinSamples: 5})
	if err != nil {
		t.Fatal(err)
	}
	if rep.FightCount != 0 {
		t.Errorf("FightCount = %d; the excerpt is open-world and has no ENCOUNTER_START", rep.FightCount)
	}
	joined := strings.Join(rep.Incomplete, "\n")
	for _, want := range []string{"ENCOUNTER_START", "level-63"} {
		if !strings.Contains(joined, want) {
			t.Errorf("Report.Incomplete does not mention %q:\n%s", want, joined)
		}
	}
	// Whatever it does measure must still be honestly marked: an excerpt
	// is far too small for any figure to clear the floor.
	for _, p := range rep.Periodic {
		if p.Enough {
			t.Errorf("%s reports Enough from an 85-line excerpt", p.SpellName)
		}
	}
}
