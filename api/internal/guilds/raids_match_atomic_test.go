package guilds

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// failAttendanceWrites makes every attendance write for raidID fail until the test ends, so the
// recompute inside a match fails.
func (w raidWorld) failAttendanceWrites(t *testing.T, raidID int64) {
	t.Helper()
	w.exec(t, fmt.Sprintf(`
		create or replace function test_fail_attendance() returns trigger language plpgsql as $$
		begin
		  if new.raid_id = %d then raise exception 'test: attendance write refused'; end if;
		  return new;
		end $$`, raidID))
	w.exec(t, `create trigger test_fail_attendance before insert or update on guild_raid_attendance
		for each row execute function test_fail_attendance()`)
	t.Cleanup(func() {
		_, _ = w.h.pool.Exec(context.Background(), `drop trigger if exists test_fail_attendance on guild_raid_attendance`)
		_, _ = w.h.pool.Exec(context.Background(), `drop function if exists test_fail_attendance()`)
	})
}

func TestAFailedRecomputeRollsTheHandMatchBack(t *testing.T) {
	w, raid := matchedNight(t)
	seedNight(t, w.h.pool, w.gid, "other", onyLogZone, raid.StartsAt.Add(2*time.Hour), []seedFight{{
		encounterID: 1084, kill: true, duration: 5 * time.Minute,
		players: []seedPlayer{{key: memberKeyA, name: "Kraggor", class: "warrior", role: "dps"}}}})
	before := w.attendance(t, raid.ID)
	w.failAttendanceWrites(t, raid.ID)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "other", MatchedHand); err == nil {
		t.Fatal("a re-match whose recompute failed succeeded")
	}
	if r := w.raid(t, raid.ID); r.ReportID == nil || *r.ReportID != "night" {
		t.Fatalf("raid holds %v, want the old log kept", r.ReportID)
	}
	if after := w.attendance(t, raid.ID); len(after) != len(before) {
		t.Fatalf("rows %d, want the old %d", len(after), len(before))
	}
}

func TestAFailedRecomputeRollsTheAutoMatchBack(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.signUp(t, raid.ID, officerKey, "Thornhide", "druid", w.officer, SignupIn)
	w.oneFightNight(t, "log1", onyLogZone, start.Add(30*time.Minute))
	w.failAttendanceWrites(t, raid.ID)
	if _, ok, err := w.h.store.AutoMatch(t.Context(), "log1"); err == nil || ok {
		t.Fatalf("auto match = %v %v, want the failed recompute's error", ok, err)
	}
	if r := w.raid(t, raid.ID); r.ReportID != nil || r.Status != RaidPlanned {
		t.Fatalf("raid = %+v, want it untouched", r)
	}
}
