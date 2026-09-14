// logs/engine/event/malformed_test.go
package event

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// decodeNoPanic decodes one line and turns a panic into a test failure, so
// a regression reads as "this line panicked" rather than as a crashed test
// binary. Decode's contract is that a malformed line costs one event.
func decodeNoPanic(t *testing.T, d *Decoder, ln lexer.Line) (e Event) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Decode panicked on %q: %v", ln.Raw, r)
		}
	}()
	return d.Decode(ln)
}

// line lexes one log line, so a test case is written the way the file is.
func line(t *testing.T, text string) lexer.Line {
	t.Helper()
	var out []lexer.Line
	l := lexer.New()
	emit := func(ln lexer.Line) error { out = append(out, ln); return nil }
	if err := l.Feed([]byte(text+"\n"), 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 {
		t.Fatalf("lexed %d lines from %q, want 1", len(out), text)
	}
	return out[0]
}

// narrowLog is the sample layout.Infer counts to build a row whose own
// arithmetic runs off the end of the line: no header, no SPELL_CAST_SUCCESS
// and no SWING_DAMAGE, so Advanced stays 0, and every event is narrower
// than the common header plus its prefix.
const narrowLog = `9/26 20:10:00.000  EMOTE,Creature-0-1,foe
9/26 20:10:00.100  EMOTE,Creature-0-1,foe
9/26 20:10:00.200  SPELL_ENERGIZE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,116
9/26 20:10:00.300  SPELL_AURA_APPLIED,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511
9/26 20:10:00.400  ENCOUNTER_START,9001,"Warden Kelthas"
`

func inferredNarrowRow(t *testing.T) layout.Layout {
	t.Helper()
	var ls []lexer.Line
	l := lexer.New()
	emit := func(ln lexer.Line) error { ls = append(ls, ln); return nil }
	if err := l.Feed([]byte(narrowLog), 0, emit); err != nil {
		t.Fatal(err)
	}
	if err := l.Flush(emit); err != nil {
		t.Fatal(err)
	}
	return layout.Infer(ls)
}

// TestMalformedLinesNeverPanic is the table of every shape that used to
// take the decoder out with an index out of range. Each must come back as
// a parse_error or an unknown event with the raw text kept.
func TestMalformedLinesNeverPanic(t *testing.T) {
	classic := layout.ClassicWiki()
	inferred := inferredNarrowRow(t)

	cases := []struct {
		name string
		lay  layout.Layout
		text string
	}{
		{"classic ENCOUNTER_START short", classic, `9/26 20:10:00.000  ENCOUNTER_START,9001`},
		{"classic ENCOUNTER_END short", classic, `9/26 20:10:00.000  ENCOUNTER_END,9001`},
		{"classic ZONE_CHANGE bare", classic, `9/26 20:10:00.000  ZONE_CHANGE`},
		{"classic MAP_CHANGE short", classic, `9/26 20:10:00.000  MAP_CHANGE,1`},
		{"classic EMOTE short", classic, `9/26 20:10:00.000  EMOTE,Creature-0-1`},
		{"classic ENVIRONMENTAL_DAMAGE short", classic, `9/26 20:10:00.000  ENVIRONMENTAL_DAMAGE,x`},
		{"classic SPELL_HEAL_ABSORBED bare", classic, `9/26 20:10:00.000  SPELL_HEAL_ABSORBED,x`},
		{"classic SPELL_ABSORBED short", classic, `9/26 20:10:00.000  SPELL_ABSORBED,x,y`},
		{"classic UNIT_DIED short", classic, `9/26 20:10:00.000  UNIT_DIED,x`},
		{"classic ENCHANT_APPLIED short", classic, `9/26 20:10:00.000  ENCHANT_APPLIED,x`},
		{"classic SWING_DAMAGE short", classic, `9/26 20:10:00.000  SWING_DAMAGE,x,y`},
		{"inferred short EMOTE", inferred, `9/26 20:10:00.000  EMOTE,Creature-0-1,foe`},
		{"inferred narrow SPELL_ENERGIZE", inferred,
			`9/26 20:10:00.000  SPELL_ENERGIZE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,116`},
		{"inferred narrow AURA_APPLIED", inferred,
			`9/26 20:10:00.000  SPELL_AURA_APPLIED,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511`},
		{"inferred short ENCOUNTER_START", inferred, `9/26 20:10:00.000  ENCOUNTER_START,9001,"Warden Kelthas"`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ln := line(t, c.text)
			e := decodeNoPanic(t, NewDecoder(c.lay, fixtureBase), ln)
			if e.Kind != ParseError && e.Kind != Unknown {
				t.Fatalf("kind = %s, want parse_error or unknown", e.Kind)
			}
			if e.Raw != ln.Raw {
				t.Errorf("raw = %q, want the whole line %q", e.Raw, ln.Raw)
			}
			if e.Kind == ParseError && e.Error == "" {
				t.Error("a parse_error with no message says nothing to the conformance report")
			}
		})
	}
}

// TestSpellHealAbsorbedAtTheClassicWidthLeavesTotalUnset covers the
// off-by-one against the Classic row's own contract: it allows width 20,
// where the wiki's suffix stops before totalAmount.
func TestSpellHealAbsorbedAtTheClassicWidthLeavesTotalUnset(t *testing.T) {
	units := `Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0`
	short := `9/26 20:10:00.000  SPELL_HEAL_ABSORBED,` + units +
		`,116,"Frostbolt",0x10,Player-4184-000000A2,"Sunwick-Nightslayer",0x512,0x0,17,"Shield",0x2,400`
	ln := line(t, short)
	if got := len(ln.Params); got != 20 {
		t.Fatalf("the fixture line has %d fields, want the 20 the Classic row allows", got)
	}
	e := decodeNoPanic(t, NewDecoder(layout.ClassicWiki(), fixtureBase), ln)
	if e.Kind != HealAbsorbed {
		t.Fatalf("kind = %s, want heal_absorbed", e.Kind)
	}
	if e.Amount.V != 400 || !e.Amount.OK {
		t.Errorf("amount = %+v, want 400", e.Amount)
	}
	if e.Total.OK {
		t.Errorf("total = %+v, want unset: at width 20 the field is not on the line", e.Total)
	}

	// At width 21 the field is there and must be read.
	e = decodeNoPanic(t, NewDecoder(layout.ClassicWiki(), fixtureBase), line(t, short+",900"))
	if !e.Total.OK || e.Total.V != 900 {
		t.Errorf("total = %+v, want 900 at width 21", e.Total)
	}
}

// TestASpecialWithNoDeclaredWidthIsKeptRaw pins the inverted default: a row
// that does not say how wide an event is must not be taken as permission to
// index into it.
func TestASpecialWithNoDeclaredWidthIsKeptRaw(t *testing.T) {
	l := layout.RetailV16()
	l.Specials["ENCOUNTER_START"] = layout.Special{}
	ln := line(t, `9/26 20:10:00.000  ENCOUNTER_START,9001,"Warden Kelthas",16,20,2450`)
	e := decodeNoPanic(t, NewDecoder(l, fixtureBase), ln)
	if e.Kind != Unknown {
		t.Fatalf("kind = %s, want unknown", e.Kind)
	}
	if e.Raw != ln.Raw {
		t.Errorf("raw = %q, want the whole line", e.Raw)
	}
	if e.Encounter != nil {
		t.Error("the line was decoded against a row that declares no width for it")
	}
}

// TestEverySpecialInEveryRegisteredRowDeclaresItsWidths is the standing
// guard on the inverted default: a new row that forgets a width list would
// otherwise silently keep those lines raw.
func TestEverySpecialInEveryRegisteredRowDeclaresItsWidths(t *testing.T) {
	for _, row := range layout.Rows() {
		for name, spec := range row.Specials {
			if len(spec.Widths) == 0 {
				t.Errorf("row %q declares no widths for %s", row.Name, name)
			}
		}
	}
}

// TestEveryDeclaredSpecialWidthDecodesWithoutPanicking walks every width
// every registered row allows and feeds a line of exactly that width made
// of one repeated placeholder field, which is the cheapest way to prove no
// branch indexes past a width its own row permits.
func TestEveryDeclaredSpecialWidthDecodesWithoutPanicking(t *testing.T) {
	for _, row := range layout.Rows() {
		for name, spec := range row.Specials {
			for _, w := range spec.Widths {
				t.Run(fmt.Sprintf("%s/%s/%d", row.Name, name, w), func(t *testing.T) {
					params := append([]string{name}, strings.Split(strings.Repeat("0 ", w-1), " ")[:w-1]...)
					ln := lexer.Line{Stamp: "9/26 20:10:00.000", Raw: strings.Join(params, ","), Params: params}
					decodeNoPanic(t, NewDecoder(row, fixtureBase), ln)
				})
			}
		}
	}
}
