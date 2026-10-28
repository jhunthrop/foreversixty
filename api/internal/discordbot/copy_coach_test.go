package discordbot

import (
	"testing"
)

const (
	coachSpecPath    = "../../../design/specs/2026-10-28-log-coach.md"
	coachSpecHeading = "### 8.3 Discord"
)

// coachSpecWords is the §8.3 row's words for key, whitespace collapsed as the coach package's
// specQuoted does.
func coachSpecWords(t *testing.T, key string) string {
	t.Helper()
	segments := specCellIn(t, coachSpecPath, coachSpecHeading, key)
	if len(segments) != 1 {
		t.Fatalf("§8.3 %s has %d quoted strings", key, len(segments))
	}
	return collapseSpaces(segments[0])
}

func TestCoachCopyIsTheSpec(t *testing.T) {
	for key, got := range map[string]string{
		"discordFieldName": coachCopy.FieldName,
		"discordLine":      coachCopy.Line,
		"discordLineOne":   coachCopy.LineOne,
		"discordLineAll":   coachCopy.LineAll,
		"discordLineNamed": coachCopy.Named,
		"discordPending":   coachCopy.Pending,
		"raidNight":        coachCopy.RaidNight,
	} {
		if want := coachSpecWords(t, key); got != want {
			t.Errorf("%s = %q, spec %q", key, got, want)
		}
	}
}
