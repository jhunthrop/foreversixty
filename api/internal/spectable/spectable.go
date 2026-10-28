// Package spectable reads the backticked strings out of the experience specs' tables, so a
// package's copy tests pin its words to the spec without each carrying its own reader.
package spectable

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

var backticked = regexp.MustCompile("`([^`]+)`")

var whitespaceRun = regexp.MustCompile(`\s+`)

// Collapse is s with every whitespace run made one space.
func Collapse(s string) string { return whitespaceRun.ReplaceAllString(s, " ") }

// Cell is the backticked segments of the string cell of the spec table row whose first column
// starts with label, under the first occurrence of heading in the spec file at path.
func Cell(t *testing.T, path, heading, label string) []string {
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
