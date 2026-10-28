package discordbot

import (
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/spectable"
)

const specPath = "../../../design/specs/2026-10-27-discord-bot.md"

// specCell is the backticked segments of a spec table row's string cell, found
// by the row's first-column label, as web/src/test-support/spec-table.ts does.
func specCell(t *testing.T, heading, label string) []string {
	t.Helper()
	return spectable.Cell(t, specPath, heading, label)
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
