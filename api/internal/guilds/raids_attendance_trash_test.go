package guilds

import (
	"testing"
	"time"
)

// A log of trash pulls only has fight_metrics rows but no encounter fight, so there is no
// first pull to measure from: the recompute must still succeed and leave every offset nil.
func TestATrashOnlyLogRecomputesWithNoFirstPullOffsets(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.signUp(t, raid.ID, officerKey, "Thornhide", "druid", w.officer, SignupIn)
	players := []seedPlayer{{key: officerKey, name: "Thornhide", class: "druid", role: "dps"}}
	seedNight(t, w.h.pool, w.gid, "trashnight", onyLogZone, start.Add(time.Hour), []seedFight{
		{offset: 0, duration: 2 * time.Minute, players: players},
		{offset: 10 * time.Minute, duration: 2 * time.Minute, players: players},
	})
	// seedNight writes metrics for boss pulls only; a trash pull's rows are added by hand.
	for i := range 2 {
		if _, err := w.h.pool.Exec(t.Context(),
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role,
				metric_dps, deaths, duration_ms, kill, fought_at)
			 values ('trashnight', $1, $2, 'Thornhide', 'druid', '', 'dps', 100, 0, 120000, false, $3)`,
			i, officerKey, start.Add(time.Hour)); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "trashnight", MatchedHand); err != nil {
		t.Fatalf("matching a trash-only log: %v", err)
	}
	if err := w.h.store.RecomputeAttendance(t.Context(), raid.ID); err != nil {
		t.Fatalf("recomputing a trash-only log: %v", err)
	}
	rows := w.attendance(t, raid.ID)
	if len(rows) == 0 {
		t.Fatal("no attendance rows were written")
	}
	for _, r := range rows {
		if r.FirstSeenOffsetS != nil || r.Outcome == OutcomeLate {
			t.Fatalf("%s: offset %v outcome %s, want no offset and not late", r.CharacterKey, r.FirstSeenOffsetS, r.Outcome)
		}
	}
}

// Trash packs count on neither side: a raider in two trash pulls and both boss pulls of a
// two-boss-pull night is "2 of 2", never "4 of 2" (round 3 e2e, item 1).
func TestPullsPresentCountsBossPullsOnly(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.signUp(t, raid.ID, officerKey, "Thornhide", "druid", w.officer, SignupIn)
	players := []seedPlayer{{key: officerKey, name: "Thornhide", class: "druid", role: "dps"}}
	seedNight(t, w.h.pool, w.gid, "mixednight", onyLogZone, start.Add(time.Hour), []seedFight{
		{offset: 0, duration: 2 * time.Minute, players: players},
		{offset: 5 * time.Minute, duration: 2 * time.Minute, players: players},
		{offset: 10 * time.Minute, duration: 5 * time.Minute, encounterID: 1084, players: players},
		{offset: 20 * time.Minute, duration: 5 * time.Minute, encounterID: 1084, kill: true, players: players},
	})
	for i := range 2 {
		if _, err := w.h.pool.Exec(t.Context(),
			`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role,
				metric_dps, deaths, duration_ms, kill, fought_at)
			 values ('mixednight', $1, $2, 'Thornhide', 'druid', '', 'dps', 100, 0, 120000, false, $3)`,
			i, officerKey, start.Add(time.Hour)); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "mixednight", MatchedHand); err != nil {
		t.Fatal(err)
	}
	if err := w.h.store.RecomputeAttendance(t.Context(), raid.ID); err != nil {
		t.Fatal(err)
	}
	for _, r := range w.attendance(t, raid.ID) {
		if r.CharacterKey != officerKey {
			continue
		}
		if r.PullsPresent != 2 || r.PullsTotal != 2 {
			t.Fatalf("pulls %d of %d, want 2 of 2", r.PullsPresent, r.PullsTotal)
		}
		return
	}
	t.Fatal("no attendance row for the raider")
}
