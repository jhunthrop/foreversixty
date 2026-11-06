package guilds

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// seedLoggedRaid inserts a done Onyxia raid at start with a public log and the attendance rows
// given as key -> {signed, outcome} (signed "" is no sign-up).
func seedLoggedRaid(t *testing.T, w raidWorld, n int, start time.Time, rows map[string][2]string) int64 {
	t.Helper()
	report := fmt.Sprintf("recent-%d-%d", w.gid, n)
	seedRankedReport(t, w.h.pool, w.gid, report)
	var id int64
	if err := w.h.pool.QueryRow(t.Context(), `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, report_id, created_by)
		values ($1, $2, 'Onyxia', $3, 180, 'done', $4, $5) returning id`, w.gid, raidZoneOny, start, report, w.officer).Scan(&id); err != nil {
		t.Fatal(err)
	}
	for key, r := range rows {
		var signed *string
		if r[0] != "" {
			signed = &r[0]
		}
		if _, err := w.h.pool.Exec(t.Context(), `insert into guild_raid_attendance (raid_id, character_key, name, class, signed, outcome)
			values ($1, $2, $2, 'warrior', $3, $4)`, id, key, signed, r[1]); err != nil {
			t.Fatal(err)
		}
	}
	return id
}

func TestRecentAttendanceGrid(t *testing.T) {
	w := newRaidWorld(t, true)
	base := time.Now().Add(-30 * 24 * time.Hour).UTC().Truncate(time.Hour)
	in, present := [2]string{"in", OutcomePresent}, [2]string{"", OutcomePresent}
	noShow, bench := [2]string{"in", OutcomeNoShow}, [2]string{"bench", OutcomeAbsent}
	nights := []map[string][2]string{
		{officerKey: in, memberKeyA: in},
		{officerKey: in, memberKeyA: in, member2Key: bench},
		{officerKey: present, memberKeyA: noShow, member2Key: in},
		{officerKey: in, memberKeyA: noShow, member2Key: in},
		{officerKey: in},
	}
	for i, rows := range nights {
		seedLoggedRaid(t, w, i, base.Add(time.Duration(i)*7*24*time.Hour), rows)
	}
	// The officer and the member are from before joined_at (null: in the guild for every night);
	// member2 joined after the first night.
	if _, err := w.h.pool.Exec(t.Context(), `update guild_members set joined_at = case when user_id = $2 then $3::timestamptz end
		where guild_id = $1`, w.gid, w.member2, base.Add(24*time.Hour)); err != nil {
		t.Fatal(err)
	}

	w.as(w.member)
	expectError(t, w.h.do(http.MethodGet, w.path("/attendance/recent"), ""), http.StatusForbidden, "forbidden")

	w.as(w.officer)
	var v RecentAttendanceView
	w.h.data(w.h.do(http.MethodGet, w.path("/attendance/recent"), ""), &v)
	if len(v.Raids) != 4 || !v.Raids[0].StartsAt.Before(v.Raids[3].StartsAt) || v.Raids[0].ZoneName != "Onyxia's Lair" {
		t.Fatalf("raids = %+v", v.Raids)
	}
	got := map[string]RecentAttendanceRow{}
	for _, r := range v.Rows {
		got[r.Name] = r
	}
	if len(got) != 2 {
		t.Fatalf("default rows = %+v", v.Rows)
	}
	if k := got["Kraggor"]; !k.Slipping || fmt.Sprint(k.Cells) != "[present absent absent absent]" {
		t.Fatalf("Kraggor = %+v", k)
	}
	if m := got["Mirelle"]; m.Slipping || fmt.Sprint(m.Cells) != "[bench present present absent]" {
		t.Fatalf("Mirelle = %+v", m)
	}

	w.h.data(w.h.do(http.MethodGet, w.path("/attendance/recent?n=5&all=1"), ""), &v)
	got = map[string]RecentAttendanceRow{}
	for _, r := range v.Rows {
		got[r.Name] = r
	}
	if len(v.Raids) != 5 || len(got) != 3 || fmt.Sprint(got["Mirelle"].Cells) != "[none bench present present absent]" ||
		fmt.Sprint(got["Thornhide"].Cells) != "[present present present present present]" {
		t.Fatalf("all rows = %+v", v.Rows)
	}
	expectError(t, w.h.do(http.MethodGet, w.path("/attendance/recent?n=0"), ""), http.StatusUnprocessableEntity, "invalid")
}

func TestOfficerActions(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	expectError(t, w.h.do(http.MethodGet, w.path("/officer/actions"), ""), http.StatusForbidden, "forbidden")

	w.as(w.officer)
	var a OfficerActions
	w.h.data(w.h.do(http.MethodGet, w.path("/officer/actions"), ""), &a)
	if a.Approvals.Count != 1 || a.Approvals.Names[0] != "Unverifiedalt" {
		t.Fatalf("approvals = %+v", a.Approvals)
	}
	if a.NotSigned.RaidID == nil || *a.NotSigned.RaidID != raid.ID || a.NotSigned.Count != 3 {
		t.Fatalf("not_signed = %+v", a.NotSigned)
	}
	if a.DropsToAward != nil {
		t.Fatalf("drops with no logged night = %+v", a.DropsToAward)
	}

	seedLoggedRaid(t, w, 0, time.Now().Add(-48*time.Hour), map[string][2]string{officerKey: {"in", OutcomePresent}})
	w.h.data(w.h.do(http.MethodGet, w.path("/officer/actions"), ""), &a)
	if d := a.DropsToAward; d == nil || d.Tracked || d.Count != nil || d.RaidID == nil || d.Awarded != 0 || d.Label == "" {
		t.Fatalf("drops = %+v", a.DropsToAward)
	}
}
