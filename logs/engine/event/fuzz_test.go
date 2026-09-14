// logs/engine/event/fuzz_test.go
package event

import (
	"os"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/lexer"
)

// malformedSeeds are the shapes that used to take the decoder out with an
// index out of range, kept as fuzz seeds so the corpus starts at the
// interesting edge rather than having to find it again.
var malformedSeeds = []string{
	"",
	"\n",
	"9/26 20:10:00.000  ",
	"9/26 20:10:00.000  ZONE_CHANGE",
	"9/26 20:10:00.000  ENCOUNTER_START,9001",
	"9/26 20:10:00.000  ENCOUNTER_END,9001",
	"9/26 20:10:00.000  MAP_CHANGE,1",
	"9/26 20:10:00.000  EMOTE,Creature-0-1",
	"9/26 20:10:00.000  EMOTE,Creature-0-1,foe",
	"9/26 20:10:00.000  ENVIRONMENTAL_DAMAGE,x",
	"9/26 20:10:00.000  SPELL_HEAL_ABSORBED,x",
	"9/26 20:10:00.000  SPELL_ABSORBED,x,y",
	"9/26 20:10:00.000  UNIT_DIED,x",
	"9/26 20:10:00.000  SWING_DAMAGE,x,y",
	`9/26 20:10:00.000  SPELL_ENERGIZE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,116`,
	`9/26 20:10:00.000  SPELL_AURA_APPLIED,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511`,
	"9/26 20:10:00.000  COMBAT_LOG_VERSION",
	"9/26 20:10:00.000  COMBATANT_INFO,Player-4184-000000A1",
	"not a log line at all",
	",,,,,,,,,,",
	`"`,
}

// FuzzDecode feeds arbitrary bytes through the lexer and decodes every
// line that comes out against the retail row, the Classic row, and a row
// inferred from the input itself — the last of these being the one whose
// field counts an attacker controls. Decode's contract is that it never
// fails: a line it cannot read costs one parse_error, not the process.
func FuzzDecode(f *testing.F) {
	text, err := os.ReadFile("testdata/v16.log")
	if err != nil {
		f.Fatal(err)
	}
	for _, ln := range strings.Split(string(text), "\n") {
		if strings.TrimSpace(ln) != "" {
			f.Add(ln)
		}
	}
	for _, ln := range malformedSeeds {
		f.Add(ln)
	}

	f.Fuzz(func(t *testing.T, in string) {
		var lines []lexer.Line
		lx := lexer.New()
		emit := func(ln lexer.Line) error { lines = append(lines, ln); return nil }
		if err := lx.Feed([]byte(in), 0, emit); err != nil {
			t.Fatalf("lexer refused a chunk at offset 0: %v", err)
		}
		if err := lx.Flush(emit); err != nil {
			t.Fatalf("lexer flush: %v", err)
		}

		rows := []layout.Layout{layout.RetailV16(), layout.ClassicWiki(), layout.Infer(lines)}
		for _, row := range rows {
			d := NewDecoder(row, fixtureBase)
			for _, ln := range lines {
				e := d.Decode(ln)
				switch e.Kind {
				case ParseError:
					if e.Error == "" {
						t.Fatalf("row %q: parse_error with no message on %q", row.Name, ln.Raw)
					}
				case Unknown:
					// Unknown keeps the raw text so nothing is lost.
					if e.Raw != ln.Raw {
						t.Fatalf("row %q: unknown event dropped its raw text on %q", row.Name, ln.Raw)
					}
				}
			}
		}
	})
}
