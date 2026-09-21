package addon

import "testing"

func TestValidateGuildNameNormalisesAndBounds(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		want   string
		wantOK bool
	}{
		{name: "ordinary name passes through", input: "Iron Vanguard", want: "Iron Vanguard", wantOK: true},
		{name: "leading and trailing space trimmed", input: "  Iron Vanguard  ", want: "Iron Vanguard", wantOK: true},
		{name: "exactly 24 runes is allowed", input: "123456789012345678901234", want: "123456789012345678901234", wantOK: true},
		{name: "25 runes is refused", input: "1234567890123456789012345", wantOK: false},
		{name: "a control character is refused", input: "Iron\tVanguard", wantOK: false},
		{name: "empty after trimming is refused", input: "   ", wantOK: false},
		{name: "invalid UTF-8 is refused", input: string([]byte{0xff, 0xfe}), wantOK: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := validateGuildName(c.input)
			if ok != c.wantOK {
				t.Fatalf("ok = %v, want %v", ok, c.wantOK)
			}
			if ok && got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}

// TestValidateGuildNameRejectsFormatCharacters is B (2026-09-21 second
// security review response): Cf format characters must be refused
// alongside the pre-existing Cc control-character check, since none of
// them can appear in a real WoW guild name and a bidi override in
// particular can make a displayed name misleading.
func TestValidateGuildNameRejectsFormatCharacters(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{name: "zero-width space (U+200B)", input: "Iron\u200bVanguard"},
		{name: "zero-width joiner (U+200D)", input: "Iron\u200dVanguard"},
		{name: "right-to-left override (U+202E)", input: "Iron\u202eVanguard"},
		{name: "byte-order mark (U+FEFF)", input: "\ufeffIron Vanguard"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, ok := validateGuildName(c.input); ok {
				t.Fatalf("a name containing %s should be refused", c.name)
			}
		})
	}
}

func TestValidateGuildNameNormalisesToNFC(t *testing.T) {
	// "e" + combining acute accent (NFD) must normalise to the single
	// precomposed "é" (NFC) so two exports that differ only in Unicode
	// normalisation form resolve to the same guild.
	decomposed := "Café" // "Cafe" + combining acute over the e
	precomposed := "Café" // "Café"
	got, ok := validateGuildName(decomposed)
	if !ok {
		t.Fatal("a decomposed-but-otherwise-valid name should be accepted")
	}
	if got != precomposed {
		t.Fatalf("got %q (%d runes), want the NFC form %q (%d runes)", got, len([]rune(got)), precomposed, len([]rune(precomposed)))
	}
}
