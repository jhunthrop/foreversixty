package guilds

import (
	"net/http"
	"testing"
)

// The hand match refuses a log that is not evidence (private or flagged), as the picker never
// lists one.
func TestHandMatchRefusesAPrivateOrFlaggedLog(t *testing.T) {
	for name, fix := range map[string]string{
		"private": `update reports set visibility = 'private' where id = 'night'`,
		"flagged": `update reports set flagged = 'tampered' where id = 'night'`,
	} {
		t.Run(name, func(t *testing.T) {
			w, raid := attendanceNight(t)
			w.grantPlan(t)
			w.exec(t, fix)
			w.as(w.officer)
			expectError(t, w.h.do(http.MethodPut, w.path("/raids/%d/match", raid.ID), `{"report_id":"night"}`),
				http.StatusUnprocessableEntity, "invalid")
			if r := w.raid(t, raid.ID); r.ReportID != nil {
				t.Fatalf("raid holds %v, want no log", *r.ReportID)
			}
		})
	}
}

// A log flagged after its match: the raid keeps the match so an officer can see and clear it,
// but night, report and attendance read null and its rows count in nobody's record.
func TestALogFlaggedAfterTheMatchShowsNothingAndCountsForNobody(t *testing.T) {
	w, raid := matchedNight(t)
	keys := []string{officerKey}
	before, err := w.h.store.raidAttendance(t.Context(), w.gid, keys, attendanceWindow)
	if err != nil || before[officerKey].Nights != 1 {
		t.Fatalf("record before = %+v %v, want one night", before, err)
	}
	w.exec(t, `update reports set flagged = 'tampered' where id = 'night'`)

	w.as(w.officer)
	v := w.getRaid(t, raid.ID)
	if v.Match == nil || v.Match.ReportID != "night" {
		t.Fatalf("match = %+v, want the flagged log still named", v.Match)
	}
	if v.Match.Pulls != 0 || v.Match.Kills != 0 || v.Match.DurationMS != 0 {
		t.Fatalf("match figures = %+v, want none read from a flagged log", v.Match)
	}
	if v.Night != nil || v.Report != nil || v.Attendance != nil {
		t.Fatalf("night %+v report %+v attendance %+v, want all null", v.Night, v.Report, v.Attendance)
	}
	after, err := w.h.store.raidAttendance(t.Context(), w.gid, keys, attendanceWindow)
	if err != nil || len(after) != 0 {
		t.Fatalf("record after = %+v %v, want no night", after, err)
	}
	if src, err := w.h.store.attendanceSource(t.Context(), w.gid); err != nil || src != AttendanceSourceReports {
		t.Fatalf("attendance source = %q %v, want reports", src, err)
	}
	// A recompute while the log is flagged leaves the stored rows alone for when it returns.
	if err := w.h.store.RecomputeAttendance(t.Context(), raid.ID); err != nil {
		t.Fatal(err)
	}
	w.exec(t, `update reports set flagged = null where id = 'night'`)
	back, err := w.h.store.raidAttendance(t.Context(), w.gid, keys, attendanceWindow)
	if err != nil || back[officerKey] != before[officerKey] {
		t.Fatalf("record once public again = %+v %v, want %+v", back, err, before)
	}
}

// The moves block never compares against a previous raid whose log is not evidence.
func TestMovesSkipAPreviousRaidWithAPrivateLog(t *testing.T) {
	w, night, _ := reportNight(t)
	w.exec(t, `update reports set visibility = 'private' where id = 'prev'`)
	w.as(w.officer)
	v := w.getRaid(t, night.ID)
	if v.Report == nil || v.Report.Moves.Previous != nil || len(v.Report.Moves.Rows) != 0 {
		t.Fatalf("moves = %+v, want none against a private log", v.Report)
	}
}
