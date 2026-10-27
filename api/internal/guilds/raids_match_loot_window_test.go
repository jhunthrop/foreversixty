package guilds

import (
	"testing"
	"time"
)

// A log hand-matched days after its night never sweeps a later raid's awards: the window ends
// six hours after the raid's own end.
func TestALateHandMatchDoesNotSweepALaterRaidsAwards(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart().Add(-3 * 24 * time.Hour)
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "late", onyLogZone, start.Add(time.Hour))
	// The log was uploaded three days after the night.
	w.exec(t, `update reports set created_at = now() - interval '1 hour' where id = 'late'`)
	during := w.seedAward(t, 0, officerKey, start.Add(30*time.Minute), nil)
	later := w.seedAward(t, 1, memberKeyA, start.Add(2*24*time.Hour), nil)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "late", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if got, _ := w.awardRaid(t, during); !onRaid(got, raid.ID) {
		t.Fatalf("the night's award is on %v, want raid %d", got, raid.ID)
	}
	if got, _ := w.awardRaid(t, later); got != nil {
		t.Fatalf("a later night's award was swept onto raid %d", *got)
	}
}
