package textx

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestTrimNeverCutsMidRune(t *testing.T) {
	// "ä" is two bytes, so a ten-rune name is twenty bytes and every
	// odd bound falls inside a rune.
	name := strings.Repeat("ä", 10)
	for max := 0; max <= len(name)+2; max++ {
		got := Trim(name, max)
		if !utf8.ValidString(got) {
			t.Fatalf("Trim(%q, %d) = %q, which is not valid UTF-8", name, max, got)
		}
		if len(got) > max {
			t.Fatalf("Trim(%q, %d) = %q, %d bytes", name, max, got, len(got))
		}
	}
}

func TestTrim(t *testing.T) {
	for _, c := range []struct {
		name string
		in   string
		max  int
		want string
	}{
		{"short strings are returned whole", "raid.txt", 120, "raid.txt"},
		{"surrounding space goes first", "  raid.txt  ", 120, "raid.txt"},
		{"a long ASCII string is cut at the bound", "abcdef", 3, "abc"},
		{"a cut inside a rune backs off", "aä", 2, "a"},
		{"a bound of zero is empty", "abc", 0, ""},
		{"the largest whole-rune prefix under the bound wins",
			strings.Repeat("\u65e5", 20), 10, "\u65e5\u65e5\u65e5"},
		{"space is trimmed before the bound is applied", "   abcdef", 3, "abc"},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := Trim(c.in, c.max); got != c.want {
				t.Fatalf("Trim(%q, %d) = %q, want %q", c.in, c.max, got, c.want)
			}
		})
	}
}
