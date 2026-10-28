package reports

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func rosterWithLevel(level int64) summary.Summary {
	return summary.Summary{Roster: []summary.RosterRow{{GUID: "Player-1", Name: "A", Level: level}}}
}

func TestIngestOverlaysLevelFromRebuild(t *testing.T) {
	posted := rosterWithLevel(0)
	overlayDerivedRoster(&posted, rosterWithLevel(60))
	if got := posted.Roster[0].Level; got != 60 {
		t.Fatalf("level = %d, want 60 from the rebuild", got)
	}
}

func TestIngestKeepsAPostedLevel(t *testing.T) {
	posted := rosterWithLevel(58)
	overlayDerivedRoster(&posted, rosterWithLevel(60))
	if got := posted.Roster[0].Level; got != 58 {
		t.Fatalf("level = %d, want the posted 58", got)
	}
}

func TestIngestLeavesLevelZeroWhenTheRebuildHasNone(t *testing.T) {
	posted := rosterWithLevel(0)
	overlayDerivedRoster(&posted, rosterWithLevel(0))
	if got := posted.Roster[0].Level; got != 0 {
		t.Fatalf("level = %d, want 0", got)
	}
}
