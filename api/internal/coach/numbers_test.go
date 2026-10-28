package coach

import (
	"slices"
	"testing"
)

func norms(tokens []NumberToken) []string {
	out := make([]string, len(tokens))
	for i, t := range tokens {
		out[i] = t.Norm
	}
	return out
}

func TestExtractNumbers(t *testing.T) {
	got := norms(ExtractNumbers("31 of a possible 40, 9.1 a minute, 44%, 2,140 at 3:05, gaps 1:12–1:25 and 2:40-2:51"))
	want := []string{"31", "40", "9.1", "44%", "2140", "3:05", "1:12–1:25", "2:40–2:51"}
	if !slices.Equal(got, want) {
		t.Fatalf("%v", got)
	}
}

func TestNormalise_EveryWayToWriteANumber(t *testing.T) {
	cases := map[string][]string{
		"44 percent":              {"44%"},
		"44 per cent":             {"44%"},
		"44 %":                    {"44%"},
		"2.1k":                    {"2100"},
		"2140":                    {"2140"},
		"9.10 casts":              {"9.1"},
		"03:05":                   {"3:05"},
		"1:12 to 1:25":            {"1:12–1:25"},
		"1:12 — 1:25":             {"1:12–1:25"},
		"at 0:05 and 12:30":       {"0:05", "12:30"},
		"a 1,234,567 total":       {"1234567"},
		"(31/40).":                {"31", "40"},
		"no digits here at all":   nil,
		"percent without a digit": nil,
	}
	for in, want := range cases {
		got := norms(ExtractNumbers(in))
		if !slices.Equal(got, want) && len(got)+len(want) != 0 {
			t.Errorf("%q: got %v, want %v", in, got, want)
		}
	}
}

func TestNormalise_Kinds(t *testing.T) {
	kinds := map[string]string{"44%": KindPercent, "3:05": KindTime, "1:12–1:25": KindRange, "31": KindNumber}
	for in, want := range kinds {
		tokens := ExtractNumbers(in)
		if len(tokens) != 1 || tokens[0].Kind != want {
			t.Errorf("%q: %+v, want kind %s", in, tokens, want)
		}
	}
}

func TestNumberWords(t *testing.T) {
	got := NumberWords("Four casts, a dozen gaps, twelve deaths, Twenty-one, one, two and three, an often tenth")
	want := []string{"four", "dozen", "twelve", "twenty", "tenth"}
	if !slices.Equal(got, want) {
		t.Fatalf("%v", got)
	}
}
