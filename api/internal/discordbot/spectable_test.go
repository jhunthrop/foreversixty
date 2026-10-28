package discordbot

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

const specPath = "../../../design/specs/2026-10-27-discord-bot.md"

var backticked = regexp.MustCompile("`([^`]+)`")

// specCell is the backticked segments of a spec table row's string cell, found
// by the row's first-column label, as web/src/test-support/spec-table.ts does.
func specCell(t *testing.T, heading, label string) []string {
	t.Helper()
	return specCellIn(t, specPath, heading, label)
}

var whitespaceRun = regexp.MustCompile(`\s+`)

// collapseSpaces turns every whitespace run into one space: spec strings may wrap inside their backticks.
func collapseSpaces(s string) string { return whitespaceRun.ReplaceAllString(s, " ") }

// specCellIn is specCell for the spec at path.
func specCellIn(t *testing.T, path, heading, label string) []string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	spec := string(raw)
	start := strings.Index(spec, heading)
	if start == -1 {
		t.Fatalf("spec has no heading %q", heading)
	}
	rows := strings.Split(spec[start:], "\n")
	row, found := "", false
	for _, prefix := range []string{"| " + label + " |", "| " + label} {
		for _, line := range rows {
			if strings.HasPrefix(line, prefix) {
				row, found = line, true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		t.Fatalf("spec %q has no row starting %q", heading, label)
	}
	cells := strings.Split(row, "|")
	if len(cells) < 3 {
		t.Fatalf("spec row %q has no string cell", label)
	}
	var out []string
	for _, m := range backticked.FindAllStringSubmatch(cells[2], -1) {
		out = append(out, m[1])
	}
	return out
}

// specString is one segment of a row with {placeholder} markers filled in.
func specString(t *testing.T, heading, label string, index int, fill map[string]string) string {
	t.Helper()
	segments := specCell(t, heading, label)
	if index >= len(segments) {
		t.Fatalf("spec %q row %q has no segment %d", heading, label, index)
	}
	text := segments[index]
	for k, v := range fill {
		text = strings.Replace(text, "{"+k+"}", v, 1)
	}
	return text
}

func TestSpecTableReadsTheCopy(t *testing.T) {
	if got := specString(t, "### 6.2 In Discord", "Any failure", 0, nil); got != commonCopy.Failure {
		t.Fatalf("failure sentence = %q", got)
	}
}
