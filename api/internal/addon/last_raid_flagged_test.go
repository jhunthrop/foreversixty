package addon

import (
	"context"
	"testing"
	"time"
)

// Task L6.12: a report flagged tampered or unverified counts nowhere a player looks. The
// inbox's last_raid is the character's newest raid that is evidence (presence.EvidenceReport):
// a newer flagged raid is not its last raid.
func TestLastRaidIsNeverAFlaggedReport(t *testing.T) {
	h := newHarness(t)
	cleanL63(t, h)
	h.exportL63(l63Export)
	now := time.Now().UTC()
	h.seedRatedFight(ratedFight{report: "l63-clean", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-3 * 24 * time.Hour), overall: 60, kill: true, missed: []string{}})
	h.seedRatedFight(ratedFight{report: "l63-flagged", visibility: "public", zone: "Molten Core", boss: "Lucifron",
		index: 1, encounter: 663, foughtAt: now.Add(-1 * 24 * time.Hour), overall: 99, kill: true, missed: []string{}})
	if _, err := h.pool.Exec(context.Background(), `update reports set flagged = 'tampered' where id = 'l63-flagged'`); err != nil {
		t.Fatal(err)
	}

	s := &Service{Store: h.store, Log: quietLog()}
	got := messagesOf(s.messages(context.Background(), h.owner, now), messageLastRaid)
	if len(got) != 1 || got[0].ReportID != "l63-clean" {
		t.Fatalf("last_raid = %+v, want l63-clean (the newer flagged raid is no raid)", got)
	}
}
