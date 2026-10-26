package guilds

import (
	"reflect"
	"testing"
	"time"
)

// match.zone_name is the matched log's zone, not the raid's planned one.
func TestMatchZoneNameIsTheLogsZone(t *testing.T) {
	w := newRaidWorld(t, true)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	seedNight(t, w.h.pool, w.gid, "elsewhere", "Molten Core", start.Add(time.Hour), []seedFight{
		{encounterID: 1084, kill: true, duration: 5 * time.Minute},
	})
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "elsewhere", MatchedHand); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	v := w.getRaid(t, raid.ID)
	if v.Match == nil || v.Match.ZoneName != "Molten Core" {
		t.Fatalf("match = %+v, want the log's zone Molten Core (raid zone %q)", v.Match, v.ZoneName)
	}
}

// A log that named no zone falls back to the raid's zone name.
func TestMatchZoneNameFallsBackToTheRaidsZone(t *testing.T) {
	w := newRaidWorld(t, true)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	seedNight(t, w.h.pool, w.gid, "nozone", "", start.Add(time.Hour), []seedFight{
		{encounterID: 1084, kill: true, duration: 5 * time.Minute},
	})
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "nozone", MatchedHand); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	v := w.getRaid(t, raid.ID)
	if v.Match == nil || v.Match.ZoneName != v.ZoneName || v.ZoneName == "" {
		t.Fatalf("match = %+v, want the raid's zone %q", v.Match, v.ZoneName)
	}
}

// The batched read returns every raid's rows in AttendanceRows' order, and an empty slice for
// a raid with none.
func TestAttendanceRowsForBatchesPerRaidInTheSameOrder(t *testing.T) {
	w, raid := attendanceNight(t)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	other := w.raidAt(t, raidZoneOny, pastStart().Add(-14*24*time.Hour))
	w.signUp(t, other.ID, officerKey, "Thornhide", "druid", w.officer, SignupIn)
	seedNight(t, w.h.pool, w.gid, "night3", onyLogZone, other.StartsAt.Add(time.Hour), []seedFight{
		{encounterID: 1084, kill: true, duration: 5 * time.Minute,
			players: []seedPlayer{{key: officerKey, name: "Thornhide", class: "druid", role: "dps"}}},
	})
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, other.ID, "night3", MatchedHand); err != nil {
		t.Fatal(err)
	}
	empty := w.raidAt(t, raidZoneOny, pastStart().Add(-28*24*time.Hour))

	got, err := w.h.store.attendanceRowsFor(t.Context(), []int64{raid.ID, other.ID, empty.ID})
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range []Raid{raid, other} {
		want := w.attendance(t, r.ID)
		if len(want) == 0 || !reflect.DeepEqual(got[r.ID], want) {
			t.Fatalf("raid %d: batched %+v, single %+v", r.ID, got[r.ID], want)
		}
	}
	if rows, ok := got[empty.ID]; !ok || rows == nil || len(rows) != 0 {
		t.Fatalf("raid without rows = %v (present %v), want an empty slice", rows, ok)
	}
}
