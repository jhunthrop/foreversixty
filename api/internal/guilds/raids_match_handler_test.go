package guilds

import (
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	nightDuration = 68 * time.Minute
	dayBefore     = -24 * time.Hour
)

// matchedNight is the attendance night with the guild plan, matched by hand and ended.
func matchedNight(t *testing.T) (raidWorld, Raid) {
	t.Helper()
	w, raid := attendanceNight(t)
	w.grantPlan(t)
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, raid.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	return w, raid
}

func (w raidWorld) grantPlan(t *testing.T) {
	t.Helper()
	if _, err := w.h.pool.Exec(t.Context(), scopedGrantSQL, w.gid); err != nil {
		t.Fatal(err)
	}
}

func (w raidWorld) dropPlan(t *testing.T) {
	t.Helper()
	if _, err := w.h.pool.Exec(t.Context(), `delete from entitlements where guild_id = $1`, w.gid); err != nil {
		t.Fatal(err)
	}
}

func (w raidWorld) setContested(t *testing.T, contested bool) {
	t.Helper()
	sql := `update guilds set claim_contested_at = null where id = $1`
	if contested {
		sql = `update guilds set claim_contested_at = now() where id = $1`
	}
	if _, err := w.h.pool.Exec(t.Context(), sql, w.gid); err != nil {
		t.Fatal(err)
	}
}

func attendanceKeys(v RaidView) map[string]AttendanceView {
	out := map[string]AttendanceView{}
	for _, a := range v.Attendance {
		out[a.CharacterKey] = a
	}
	return out
}

func TestAnUnmatchedRaidViewCarriesNullMatchNightAttendance(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	data := rawData(t, w.h.do(http.MethodGet, w.path("/raids/%d", raid.ID), ""))
	for _, key := range []string{"pre_done_status", "match", "night", "attendance"} {
		if got, ok := data[key]; !ok || string(got) != "null" {
			t.Fatalf("%s = %s (present %v), want null", key, got, ok)
		}
	}
}

func TestAMatchedRaidViewShowsTheMatchAndTheNightToEveryone(t *testing.T) {
	w, raid := matchedNight(t)
	for _, uid := range []int64{w.officer, w.member} {
		w.as(uid)
		v := w.getRaid(t, raid.ID)
		if v.Status != RaidDone || v.PreDoneStatus == nil || *v.PreDoneStatus != RaidPlanned {
			t.Fatalf("status %q pre_done %v", v.Status, v.PreDoneStatus)
		}
		m := v.Match
		if m == nil || m.ReportID != "night" || m.ZoneName != "Onyxia's Lair" || m.How != MatchedHand ||
			m.Pulls != 4 || m.Kills != 1 || m.DurationMS != nightDuration.Milliseconds() || m.UploadedAt.IsZero() {
			t.Fatalf("match = %+v", m)
		}
		want := NightView{Signed: 4, Came: 3, Late: 1, NotInLog: 1, Joined: 1, Excused: 0, Kills: 1, Wipes: 3,
			Deaths: 12, DurationMS: nightDuration.Milliseconds()}
		if v.Night == nil || *v.Night != want {
			t.Fatalf("night = %+v, want %+v", v.Night, want)
		}
	}
}

func TestAttendanceRowsGoToOfficersInFullAndToMembersOwnOnly(t *testing.T) {
	w, raid := matchedNight(t)
	w.as(w.officer)
	all := attendanceKeys(w.getRaid(t, raid.ID))
	if len(all) != 7 || !all[officerKey].Own || all[memberKeyA].Own || all[ghostKey].Outcome != OutcomeNoShow {
		t.Fatalf("officer rows = %+v", all)
	}
	w.as(w.member)
	mine := attendanceKeys(w.getRaid(t, raid.ID))
	if len(mine) != 3 {
		t.Fatalf("member rows = %+v", mine)
	}
	for _, key := range []string{memberKeyA, memberKeyB, memberKeyC} {
		if a, ok := mine[key]; !ok || !a.Own {
			t.Fatalf("member lacks own row %s: %+v", key, mine)
		}
	}
	if k := mine[memberKeyA]; k.CameAs == nil || *k.CameAs != "Sunderfel" {
		t.Fatalf("kraggor = %+v", k)
	}
	w.as(w.member2)
	if rows := attendanceKeys(w.getRaid(t, raid.ID)); len(rows) != 1 || rows[member2Key].Outcome != OutcomeLate {
		t.Fatalf("member2 rows = %+v", rows)
	}
}

func TestAnOfficerWithoutThePlanSeesOnlyTheirOwnAttendance(t *testing.T) {
	w, raid := matchedNight(t)
	w.dropPlan(t)
	w.as(w.officer)
	if rows := attendanceKeys(w.getRaid(t, raid.ID)); len(rows) != 1 || !rows[officerKey].Own {
		t.Fatalf("rows = %+v", rows)
	}
}

func TestMatchCandidatesIsOfficerOnlyAndMarksTheClosest(t *testing.T) {
	w, raid := attendanceNight(t)
	w.grantPlan(t)
	other := w.raidAt(t, raidZoneOny, pastStart().Add(dayBefore))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, other.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	path := w.path("/raids/%d/match-candidates", raid.ID)
	w.as(w.member)
	expectError(t, w.h.do(http.MethodGet, path, ""), http.StatusForbidden, "forbidden")
	w.as(w.officer)
	res := w.h.do(http.MethodGet, path, "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("candidates = %d: %s", res.StatusCode, readBody(res))
	}
	var out MatchCandidatesView
	w.h.data(res, &out)
	if !out.ZoneKnown || len(out.Rows) != 1 {
		t.Fatalf("candidates = %+v", out)
	}
	row := out.Rows[0]
	if row.ReportID != "night" || !row.Closest || row.Pulls != 4 || row.Kills != 1 ||
		row.MatchedRaid == nil || row.MatchedRaid.ID != other.ID {
		t.Fatalf("row = %+v", row)
	}
	expectError(t, w.h.do(http.MethodGet, w.path("/raids/999999/match-candidates"), ""), http.StatusNotFound, "not_found")
}

func TestPutMatchMatchesAndAnswersTheRaidView(t *testing.T) {
	w, raid := attendanceNight(t)
	w.grantPlan(t)
	path := w.path("/raids/%d/match", raid.ID)
	w.as(w.officer)
	res := w.h.do(http.MethodPut, path, `{"report_id":"night"}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("put match = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	if v.Status != RaidDone || v.Match == nil || v.Match.How != MatchedHand || v.Night == nil || len(v.Attendance) != 7 {
		t.Fatalf("view = %+v", v)
	}
}

func TestPutMatchRefusals(t *testing.T) {
	w, raid := attendanceNight(t)
	w.grantPlan(t)
	other := w.raidAt(t, raidZoneOny, pastStart().Add(dayBefore))
	if _, err := w.h.store.SetMatch(t.Context(), w.gid, other.ID, "night", MatchedHand); err != nil {
		t.Fatal(err)
	}
	future := w.createRaid(t)
	cancelled := w.raidAt(t, raidZoneOny, pastStart().Add(3*dayBefore))
	if _, err := w.h.store.SetRaidStatus(t.Context(), w.gid, cancelled.ID, []string{RaidPlanned}, RaidCancelled); err != nil {
		t.Fatal(err)
	}
	w.as(w.officer)
	put := func(id int64, body string) *http.Response {
		return w.h.do(http.MethodPut, w.path("/raids/%d/match", id), body)
	}
	expectError(t, put(raid.ID, `{"report_id":"night"}`), http.StatusConflict, "report_matched")
	expectError(t, put(future.ID, `{"report_id":"night"}`), http.StatusConflict, "conflict")
	expectError(t, put(cancelled.ID, `{"report_id":"night"}`), http.StatusConflict, "conflict")
	expectError(t, put(raid.ID, `{"report_id":"nope"}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put(raid.ID, `{"report_id":""}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put(raid.ID, `{}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put(raid.ID, `nonsense`), http.StatusBadRequest, "invalid")
	expectError(t, put(999999, `{"report_id":"night"}`), http.StatusNotFound, "not_found")
}

func TestMatchRoutesRefuseMembersAndNeedThePlanAndHonourTheFreeze(t *testing.T) {
	w, raid := matchedNight(t)
	routes := []struct{ method, path, body string }{
		{http.MethodGet, w.path("/raids/%d/match-candidates", raid.ID), ""},
		{http.MethodPut, w.path("/raids/%d/match", raid.ID), `{"report_id":"night"}`},
		{http.MethodDelete, w.path("/raids/%d/match", raid.ID), ""},
		{http.MethodPut, w.path("/raids/%d/attendance/%s", raid.ID, officerKey), `{"excused":true}`},
	}
	w.as(w.member)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, rt.path, rt.body), http.StatusForbidden, "forbidden")
	}
	w.setContested(t, true)
	w.as(w.officer)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, rt.path, rt.body), http.StatusConflict, "claim_contested")
	}
	w.setContested(t, false)
	w.dropPlan(t)
	for _, rt := range routes {
		expectError(t, w.h.do(rt.method, rt.path, rt.body), http.StatusForbidden, "guild_plan_required")
	}
}

func TestDeleteMatchClearsTheLogAndRestoresTheStatus(t *testing.T) {
	w, raid := matchedNight(t)
	path := w.path("/raids/%d/match", raid.ID)
	w.as(w.officer)
	res := w.h.do(http.MethodDelete, path, "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("delete match = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	if v.Status != RaidPlanned || v.Match != nil || v.Night != nil || v.Attendance != nil || v.PreDoneStatus != nil {
		t.Fatalf("view = %+v", v)
	}
	expectError(t, w.h.do(http.MethodDelete, path, ""), http.StatusConflict, "conflict")
	expectError(t, w.h.do(http.MethodDelete, w.path("/raids/999999/match"), ""), http.StatusNotFound, "not_found")
}

func TestPutAttendanceExcusesAndTheNightCountsIt(t *testing.T) {
	w, raid := matchedNight(t)
	path := w.path("/raids/%d/attendance/%s", raid.ID, ghostKey)
	w.as(w.officer)
	res := w.h.do(http.MethodPut, path, `{"excused":true}`)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("excuse = %d: %s", res.StatusCode, readBody(res))
	}
	var v RaidView
	w.h.data(res, &v)
	if !attendanceKeys(v)[ghostKey].Excused || v.Night == nil || v.Night.Excused != 1 {
		t.Fatalf("view = %+v night %+v", attendanceKeys(v)[ghostKey], v.Night)
	}
	var back RaidView
	w.h.data(w.h.do(http.MethodPut, path, `{"excused":false}`), &back)
	if attendanceKeys(back)[ghostKey].Excused || back.Night.Excused != 0 {
		t.Fatalf("unexcused view = %+v", back.Night)
	}
}

func TestPutAttendanceRefusals(t *testing.T) {
	w, raid := matchedNight(t)
	w.as(w.officer)
	put := func(key, body string) *http.Response {
		return w.h.do(http.MethodPut, w.path("/raids/%d/attendance/%s", raid.ID, key), body)
	}
	expectError(t, put("us/hardcore/nobody", `{"excused":true}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put(ghostKey, `{}`), http.StatusUnprocessableEntity, "invalid")
	expectError(t, put(ghostKey, `junk`), http.StatusBadRequest, "invalid")
	expectError(t, put(strings.Replace(ghostKey, "us/hardcore", "xx/yy", 1), `{"excused":true}`), http.StatusNotFound, "not_found")
}

func TestARaidsTabNightLinksToItsRaid(t *testing.T) {
	w, raid := matchedNight(t)
	w.oneFightNight(t, "unmatched", onyLogZone, pastStart().Add(-6*time.Hour))
	w.as(w.member)
	var page struct {
		Rows []struct {
			ID     string `json:"id"`
			RaidID *int64 `json:"raid_id"`
		} `json:"rows"`
	}
	w.h.data(w.h.do(http.MethodGet, w.path("/raids"), ""), &page)
	got := map[string]*int64{}
	for _, r := range page.Rows {
		got[r.ID] = r.RaidID
	}
	if got["night"] == nil || *got["night"] != raid.ID || got["unmatched"] != nil || len(got) != 2 {
		t.Fatalf("rows = %v", got)
	}
}
