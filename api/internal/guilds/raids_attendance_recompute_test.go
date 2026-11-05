package guilds

import (
	"testing"
	"time"
)

// ghostOutcome is Ghost's attendance outcome on raidID ("" without a row).
func (w raidWorld) ghostOutcome(t *testing.T, raidID int64) string {
	t.Helper()
	return rowsByKey(w.attendance(t, raidID))[ghostKey].Outcome
}

// An officer's correction to a logged raid's sign-ups rebuilds its attendance.
func TestACorrectionToADoneRaidRecomputesAttendance(t *testing.T) {
	w, raid := matchedNight(t)
	if got := w.ghostOutcome(t, raid.ID); got != OutcomeNoShow {
		t.Fatalf("ghost before = %q", got)
	}
	cmd := w.cmd(raid.ID, w.officer, ghostKey, SignupOut)
	if _, err := w.svcSignUp(cmd); err != nil {
		t.Fatal(err)
	}
	if got := w.ghostOutcome(t, raid.ID); got != OutcomeAbsent {
		t.Fatalf("ghost after the correction = %q, want absent", got)
	}
}

// A correction made while the log was private counts once the log is visible again.
func TestAVisibilityChangeRecomputesTheRaidHoldingTheLog(t *testing.T) {
	w, raid := matchedNight(t)
	w.exec(t, `update reports set visibility = 'private' where id = 'night'`)
	if _, err := w.svcSignUp(w.cmd(raid.ID, w.officer, ghostKey, SignupOut)); err != nil {
		t.Fatal(err)
	}
	if got := w.ghostOutcome(t, raid.ID); got != OutcomeNoShow {
		t.Fatalf("a private log's rows moved: %q", got)
	}
	w.exec(t, `update reports set visibility = 'guild' where id = 'night'`)
	w.h.svc.ReportVisibilityChanged(t.Context(), "night")
	if got := w.ghostOutcome(t, raid.ID); got != OutcomeAbsent {
		t.Fatalf("ghost once the log is visible = %q, want absent", got)
	}
}

// A log made visible after it completed finds its raid.
func TestAVisibilityChangeMatchesALogNoRaidHolds(t *testing.T) {
	w := newRaidWorld(t, false)
	start := pastStart()
	raid := w.raidAt(t, raidZoneOny, start)
	w.oneFightNight(t, "hidden", onyLogZone, start.Add(30*time.Minute))
	w.exec(t, `update reports set visibility = 'private' where id = 'hidden'`)
	w.h.svc.ReportCompleted(t.Context(), "hidden")
	if r := w.raid(t, raid.ID); r.ReportID != nil {
		t.Fatalf("a private log matched: %+v", r)
	}
	w.exec(t, `update reports set visibility = 'guild' where id = 'hidden'`)
	w.h.svc.ReportVisibilityChanged(t.Context(), "hidden")
	if r := w.raid(t, raid.ID); r.ReportID == nil || *r.ReportID != "hidden" || r.Status != RaidDone {
		t.Fatalf("raid after the log was made visible = %+v", r)
	}
}
