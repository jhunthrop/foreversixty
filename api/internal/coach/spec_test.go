package coach

import (
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"
)

const specPath = "../../../docs/superpowers/specs/2026-10-27-log-coach-design.md"

var (
	specOnce  sync.Once
	specSpans []string
	specErr   error

	fencedBlock = regexp.MustCompile("(?s)```.*?```")
	inlineSpan  = regexp.MustCompile("`([^`]+)`")
	whitespace  = regexp.MustCompile(`\s+`)
)

// collapse turns every whitespace run, newlines included, into one space: most spec
// strings wrap across a line break inside their backticks.
func collapse(s string) string { return whitespace.ReplaceAllString(s, " ") }

// specQuoted reports whether s appears inside one backtick span of the log coach design
// spec, after collapsing whitespace runs in both. Fenced code blocks are not quotes.
func specQuoted(t *testing.T, s string) bool {
	t.Helper()
	specOnce.Do(func() {
		b, err := os.ReadFile(specPath)
		if err != nil {
			specErr = err
			return
		}
		text := fencedBlock.ReplaceAllString(string(b), "")
		for _, m := range inlineSpan.FindAllStringSubmatch(text, -1) {
			specSpans = append(specSpans, collapse(m[1]))
		}
	})
	if specErr != nil {
		t.Fatal(specErr)
	}
	want := collapse(s)
	for _, span := range specSpans {
		if strings.Contains(span, want) {
			return true
		}
	}
	return false
}

func TestSpecQuoted_FindsWrappedSpansAndRejectsOthers(t *testing.T) {
	// §6.1 wraps this line across a newline inside its backticks.
	if !specQuoted(t, "We couldn't tell this player's spec from the log.") {
		t.Error("a span wrapped across lines must match once collapsed")
	}
	if specQuoted(t, "Heal more next time.") {
		t.Error("a string the spec never quotes must not match")
	}
	// Prose outside backticks is not a quote: §6.2 says this only in running text.
	if specQuoted(t, "These are unchanged from the rule design.") {
		t.Error("unquoted prose must not match")
	}
}
