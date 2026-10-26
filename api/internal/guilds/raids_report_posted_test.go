package guilds

import (
	"testing"
	"time"
)

func (w raidWorld) postedAt(t *testing.T, raidID int64) *time.Time {
	t.Helper()
	var at *time.Time
	if err := w.h.pool.QueryRow(t.Context(), `select report_posted_at from guild_raids where id = $1`, raidID).Scan(&at); err != nil {
		t.Fatal(err)
	}
	return at
}

// A corrected log can be posted: a clear or a re-match to another log resets report_posted_at.
func TestAClearOrARematchLetsTheCorrectedReportBePosted(t *testing.T) {
	w, raid := matchedNight(t)
	seedNight(t, w.h.pool, w.gid, "other", onyLogZone, raid.StartsAt.Add(2*time.Hour), []seedFight{{
		encounterID: 1084, kill: true, duration: 5 * time.Minute,
		players: []seedPlayer{{key: officerKey, name: "Thornhide", class: "druid", role: "dps"}}}})
	claim := func() {
		t.Helper()
		if _, ok, err := w.h.store.ClaimReportPost(t.Context(), w.gid, raid.ID); err != nil || !ok {
			t.Fatalf("claim = %v %v", ok, err)
		}
	}
	claim()
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if w.postedAt(t, raid.ID) == nil {
		t.Fatal("matching the same log again reset the post")
	}
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "other", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if at := w.postedAt(t, raid.ID); at != nil {
		t.Fatalf("after a re-match posted_at = %v, want null", at)
	}
	claim()
	if _, err := w.h.store.ClearMatch(t.Context(), w.gid, raid.ID, w.officer); err != nil {
		t.Fatal(err)
	}
	if at := w.postedAt(t, raid.ID); at != nil {
		t.Fatalf("after a clear posted_at = %v, want null", at)
	}
}
