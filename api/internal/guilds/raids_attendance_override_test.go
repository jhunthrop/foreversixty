package guilds

import (
	"errors"
	"net/http"
	"testing"
)

const benchwarmerKey = "us/hardcore/benchwarmer"

func (w raidWorld) setOverride(t *testing.T, raidID int64, key, status string) {
	t.Helper()
	if err := w.h.store.SetAttendanceOverride(t.Context(), w.gid, raidID, key, status, w.officer); err != nil {
		t.Fatal(err)
	}
}

func (w raidWorld) clearOverride(t *testing.T, raidID int64, key string) {
	t.Helper()
	if err := w.h.store.ClearAttendanceOverride(t.Context(), w.gid, raidID, key); err != nil {
		t.Fatal(err)
	}
}

// A fix writes the effective outcome every count reads, survives a recompute of the log, and
// clearing it puts the log's value back.
func TestAttendanceOverrideSurvivesARecomputeAndClearRestoresTheLog(t *testing.T) {
	w, raid := matchedNight(t)
	w.setOverride(t, raid.ID, ghostKey, OverridePresent)
	if r := rowsByKey(w.attendance(t, raid.ID))[ghostKey]; r.Outcome != OutcomePresent {
		t.Fatalf("ghost fixed present = %+v", r)
	}
	nights, err := w.h.store.raidAttendance(t.Context(), w.gid, []string{ghostKey}, attendanceWindow)
	if err != nil {
		t.Fatal(err)
	}
	if a := nights[ghostKey]; a.Present != 1 || a.Nights != 1 {
		t.Fatalf("ghost nights = %+v, want 1 of 1", a)
	}
	if err := w.h.store.RecomputeAttendance(t.Context(), raid.ID); err != nil {
		t.Fatal(err)
	}
	if r := rowsByKey(w.attendance(t, raid.ID))[ghostKey]; r.Outcome != OutcomePresent {
		t.Fatalf("ghost after recompute = %+v", r)
	}
	w.clearOverride(t, raid.ID, ghostKey)
	if r := rowsByKey(w.attendance(t, raid.ID))[ghostKey]; r.Outcome != OutcomeNoShow {
		t.Fatalf("ghost cleared = %+v, want the log's no-show", r)
	}
	w.clearOverride(t, raid.ID, ghostKey) // a second clear is a no-op
}

func TestAttendanceOverrideStatusesMapToTheEffectiveRow(t *testing.T) {
	w, raid := matchedNight(t)
	w.setOverride(t, raid.ID, drifterKey, OverrideAbsent) // unsigned, in the log
	w.setOverride(t, raid.ID, ghostKey, OverrideAbsent)   // signed In, not in the log
	w.setOverride(t, raid.ID, member2Key, OverrideBench)  // signed Late, in the log
	w.setOverride(t, raid.ID, officerKey, OverrideExcused)
	by := rowsByKey(w.attendance(t, raid.ID))
	if by[drifterKey].Outcome != OutcomeAbsent || by[ghostKey].Outcome != OutcomeNoShow {
		t.Fatalf("absent: drifter %+v ghost %+v", by[drifterKey], by[ghostKey])
	}
	if by[member2Key].Outcome != OutcomePresent || by[member2Key].Excused {
		t.Fatalf("bench earns the night: %+v", by[member2Key])
	}
	if o := by[officerKey]; !o.Excused || o.Outcome != OutcomePresent {
		t.Fatalf("excused keeps the log's outcome: %+v", o)
	}
	w.clearOverride(t, raid.ID, officerKey)
	if o := rowsByKey(w.attendance(t, raid.ID))[officerKey]; o.Excused {
		t.Fatalf("cleared excused = %+v", o)
	}
}

// A verified raider with no row (not signed up, not in the log) gets one for the fix, and it
// goes with the fix. A character off the roster has nothing to fix.
func TestAttendanceOverrideOnARaiderWithNoRow(t *testing.T) {
	w, raid := matchedNight(t)
	w.seedRaider(t, "bench@example.com", "member", raidChar{benchwarmerKey, "Benchwarmer", "paladin", true})
	w.setOverride(t, raid.ID, benchwarmerKey, OverrideBench)
	r, ok := rowsByKey(w.attendance(t, raid.ID))[benchwarmerKey]
	if !ok || r.Outcome != OutcomePresent || r.Name != "Benchwarmer" || r.PullsTotal != 4 {
		t.Fatalf("benchwarmer = %+v (row %v)", r, ok)
	}
	if err := w.h.store.RecomputeAttendance(t.Context(), raid.ID); err != nil {
		t.Fatal(err)
	}
	if r := rowsByKey(w.attendance(t, raid.ID))[benchwarmerKey]; r.Outcome != OutcomePresent {
		t.Fatalf("benchwarmer after recompute = %+v", r)
	}
	w.clearOverride(t, raid.ID, benchwarmerKey)
	if _, ok := rowsByKey(w.attendance(t, raid.ID))[benchwarmerKey]; ok {
		t.Fatal("the bare row stayed after its fix was cleared")
	}
	for _, key := range []string{"us/hardcore/nobody"} {
		if err := w.h.store.SetAttendanceOverride(t.Context(), w.gid, raid.ID, key, OverridePresent, w.officer); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%s: err = %v, want ErrNotFound", key, err)
		}
	}
}

func TestExcusingAFixedRowClearsTheFixAndClearingTheMatchDropsThem(t *testing.T) {
	w, raid := matchedNight(t)
	w.setOverride(t, raid.ID, ghostKey, OverridePresent)
	if err := w.h.store.SetExcused(t.Context(), raid.ID, ghostKey, true); err != nil {
		t.Fatal(err)
	}
	if r := rowsByKey(w.attendance(t, raid.ID))[ghostKey]; r.Outcome != OutcomeNoShow || !r.Excused {
		t.Fatalf("ghost excused over a fix = %+v", r)
	}
	w.setOverride(t, raid.ID, drifterKey, OverrideAbsent)
	if _, err := w.h.store.ClearMatch(t.Context(), w.gid, raid.ID, w.officer); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := w.h.pool.QueryRow(t.Context(),
		`select count(*) from guild_raid_attendance_overrides where raid_id = $1`, raid.ID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("%d fixes outlived the match", n)
	}
	if err := w.h.store.SetAttendanceOverride(t.Context(), w.gid, raid.ID, ghostKey, OverridePresent, w.officer); !errors.Is(err, ErrNoLog) {
		t.Fatalf("fix on an unmatched raid: err = %v, want ErrNoLog", err)
	}
}

func TestOverrideRoutesShowAutoFixAndEffective(t *testing.T) {
	w, raid := matchedNight(t)
	path := w.path("/raids/%d/attendance/%s/override", raid.ID, ghostKey)
	w.as(w.officer)
	res := w.h.do(http.MethodPut, path, `{"status":"present"}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("put override = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	g := attendanceKeys(v)[ghostKey]
	if g.Outcome != OutcomePresent || g.AutoOutcome != OutcomeNoShow || g.Override == nil ||
		g.Override.Status != OverridePresent || g.Override.ByName == "" || g.Override.At.IsZero() {
		t.Fatalf("ghost = %+v override %+v", g, g.Override)
	}
	if v.Night == nil || v.Night.NotInLog != 0 {
		t.Fatalf("night = %+v, want the fix counted", v.Night)
	}
	if d := attendanceKeys(v)[drifterKey]; d.Override != nil || d.AutoOutcome != d.Outcome {
		t.Fatalf("an unfixed row = %+v", d)
	}
	var back RaidView
	w.h.data(w.h.do(http.MethodDelete, path, ""), &back)
	if g := attendanceKeys(back)[ghostKey]; g.Override != nil || g.Outcome != OutcomeNoShow {
		t.Fatalf("cleared ghost = %+v", g)
	}
	expectError(t, w.h.do(http.MethodPut, path, `{"status":"late"}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, w.h.do(http.MethodPut, path, `junk`), http.StatusBadRequest, "invalid")
	expectError(t, w.h.do(http.MethodPut, w.path("/raids/%d/attendance/us/hardcore/nobody/override", raid.ID),
		`{"status":"present"}`), http.StatusUnprocessableEntity, "invalid")
}

func TestOverrideRoutesAreForOfficersWithThePlan(t *testing.T) {
	w, raid := matchedNight(t)
	path := w.path("/raids/%d/attendance/%s/override", raid.ID, ghostKey)
	routes := []struct{ method, body string }{{http.MethodPut, `{"status":"present"}`}, {http.MethodDelete, ""}}
	w.as(w.member)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, path, rt.body), http.StatusForbidden, "forbidden")
	}
	w.as(0)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, path, rt.body), http.StatusNotFound, "not_found")
	}
	w.as(w.officer)
	w.setContested(t, true)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, path, rt.body), http.StatusConflict, "claim_contested")
	}
	w.setContested(t, false)
	w.dropPlan(t)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, path, rt.body), http.StatusForbidden, "guild_plan_required")
	}
}
