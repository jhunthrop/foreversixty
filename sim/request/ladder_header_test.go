package request

import (
	"strings"
	"testing"
)

// The golden header's gear prose must agree with ladderGearProfiles: shamans
// never dual wield in Forever (owner rule), so no spec whose profile carries a
// shield may be named a dual-wielder, or one-hand "for both hands".
func TestLadderRulesHeaderGearProseMatchesProfiles(t *testing.T) {
	header := strings.Join(strings.Fields(ladderRulesHeader), " ")
	clauses := []string{
		sentenceAround(t, header, "dual-wield in this build"),
		sentenceAround(t, header, "for both hands"),
	}
	for spec, profile := range ladderGearProfiles {
		if !profile.Shield {
			continue
		}
		for _, clause := range clauses {
			if strings.Contains(clause, spec) {
				t.Errorf("ladderRulesHeader names %s, a shield profile, in %q", spec, clause)
			}
		}
	}
}

// sentenceAround returns the clause of text holding marker, bounded by the
// nearest semicolon or bullet before it and the nearest closing parenthesis or
// semicolon after it.
func sentenceAround(t *testing.T, text, marker string) string {
	t.Helper()
	at := strings.Index(text, marker)
	if at < 0 {
		t.Fatalf("ladderRulesHeader has no %q", marker)
	}
	start := max(strings.LastIndex(text[:at], ";"), strings.LastIndex(text[:at], "- "), 0)
	end := len(text)
	if rest := strings.IndexAny(text[at:], ");"); rest >= 0 {
		end = at + rest
	}
	return text[start:end]
}
