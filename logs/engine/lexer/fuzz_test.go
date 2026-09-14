// logs/engine/lexer/fuzz_test.go
package lexer

import (
	"strings"
	"testing"
)

// FuzzSplitParams checks the CSV splitter's invariants on arbitrary input.
// It is the first thing every log line touches, so it sees whatever an
// upload contains.
func FuzzSplitParams(f *testing.F) {
	seeds := []string{
		"",
		",",
		",,,,,,,,,,",
		`SPELL_DAMAGE,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0`,
		`COMBATANT_INFO,Player-4184-000000A1,0,[(1,2,(3),(4),(5))],[6,7]`,
		`a,"b,c",d`,
		`a,"b\"c",d`,
		`"`,
		`\`,
		`[`,
		`(((((`,
		`)))))`,
		`"unterminated,quote`,
		"nil,nil,nil",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		got := SplitParams(s)

		// Always at least one field: a splitter that returned none would
		// make Params[0] indexing unsafe everywhere downstream.
		if len(got) == 0 {
			t.Fatalf("SplitParams(%q) returned no fields", s)
		}
		// Every byte of output came from the input, and separators,
		// quotes and escapes are dropped, so the output can never grow.
		total := 0
		for _, p := range got {
			total += len(p)
		}
		if total > len(s) {
			t.Fatalf("SplitParams(%q) produced %d bytes from %d", s, total, len(s))
		}
		// With no quoting, nesting or escaping in play the splitter must
		// agree with the standard library exactly.
		if !strings.ContainsAny(s, `"\[]()`) {
			want := strings.Split(s, ",")
			if len(want) != len(got) {
				t.Fatalf("SplitParams(%q) = %q, strings.Split = %q", s, got, want)
			}
			for i := range want {
				if want[i] != got[i] {
					t.Fatalf("SplitParams(%q)[%d] = %q, strings.Split = %q", s, i, got[i], want[i])
				}
			}
		}
		// Splitting is stable: the same input gives the same fields.
		again := SplitParams(s)
		if len(again) != len(got) {
			t.Fatalf("SplitParams(%q) is not stable: %d fields then %d", s, len(got), len(again))
		}
		for i := range got {
			if again[i] != got[i] {
				t.Fatalf("SplitParams(%q) is not stable at %d: %q then %q", s, i, got[i], again[i])
			}
		}
	})
}
