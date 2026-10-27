package guilds

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"
)

func jsonOf(t *testing.T, v any) string {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

func TestExtractedReadsMatchTheRoutes(t *testing.T) {
	t.Run("report", func(t *testing.T) {
		w, night, _ := reportNight(t)
		w.award(t, night.ID, officerKey)
		for _, who := range []int64{w.officer, w.member} {
			w.as(who)
			route := w.getRaid(t, night.ID)
			got, found, err := w.h.svc.ReportFor(t.Context(), w.gid, who, nil)
			if err != nil || !found {
				t.Fatalf("report for %d = %v %v", who, found, err)
			}
			if jsonOf(t, got.Raid) != jsonOf(t, route) || jsonOf(t, got.Report) != jsonOf(t, route.Report) {
				t.Fatalf("report for %d differs from the route:\n%s\n%s", who, jsonOf(t, got.Raid), jsonOf(t, route))
			}
			if !got.Member || got.Officer != (who == w.officer) || got.Public {
				t.Fatalf("report for %d: member %v officer %v public %v", who, got.Member, got.Officer, got.Public)
			}
		}
		outsider := seedUser(t, w.h.pool, "outsider@example.com") // signed in, no character in the guild
		for _, who := range []int64{0, outsider} {
			if _, found, err := w.h.svc.ReportFor(t.Context(), w.gid, who, nil); err != nil || found {
				t.Fatalf("%d read a guild-only log: %v %v", who, found, err)
			}
		}
		w.exec(t, `update reports set visibility = 'public' where id = 'night'`)
		for _, who := range []int64{0, outsider} {
			got, found, err := w.h.svc.ReportFor(t.Context(), w.gid, who, &night.ID)
			if err != nil || !found || !got.Public || got.Member || got.Officer || len(got.Report.Deaths) != 0 ||
				len(got.Report.Consumables.Rows) != 0 {
				t.Fatalf("report for %d (no member) = %+v %v %v", who, got, found, err)
			}
		}
		other := w.raidAt(t, raidZoneOny, time.Now().Add(72*time.Hour))
		if _, found, _ := w.h.svc.ReportFor(t.Context(), w.gid, w.officer, &other.ID); found {
			t.Fatal("a planned raid has a report")
		}
	})
	t.Run("attendance", func(t *testing.T) {
		w, _ := matchedNight(t)
		seedExport(t, w.h.pool, w.member, memberKeyA, "us", "hardcore", "Kraggor")
		seedExport(t, w.h.pool, w.member, memberKeyB, "us", "hardcore", "Sunderfel")
		seedExport(t, w.h.pool, w.member2, member2Key, "us", "hardcore", "Mirelle")
		w.as(w.member)
		res := w.h.do(http.MethodGet, w.path("/home"), "")
		if res.StatusCode != http.StatusOK {
			t.Fatalf("home = %d", res.StatusCode)
		}
		var home HomeView
		w.h.data(res, &home)
		chips := map[string]string{}
		for _, r := range home.Roster {
			chips[r.Name] = attendanceChip(r, AttendanceSourceRaids)
		}
		recs, err := w.h.svc.AttendanceFor(t.Context(), w.gid, w.member, "")
		if err != nil || len(recs) != 2 || recs[0].Name != "Kraggor" || recs[1].Name != "Sunderfel" {
			t.Fatalf("own records = %+v %v", recs, err)
		}
		for _, r := range recs {
			if r.Chip != chips[r.Name] {
				t.Fatalf("record %+v, roster chip %q", r, chips[r.Name])
			}
		}
		// Kraggor came as Sunderfel: the night counts once, on Kraggor.
		if recs[0].Chip != "Came 1 of 1 raids" || len(recs[0].Last) != 1 || recs[1].Chip != "" || len(recs[1].Last) != 0 {
			t.Fatalf("records = %+v", recs)
		}
		if l := recs[0].Last[0]; l.Outcome != OutcomePresent || l.CameAs != "Sunderfel" {
			t.Fatalf("kraggor line = %+v", l)
		}
		if _, err := w.h.svc.AttendanceFor(t.Context(), w.gid, w.member, "Mirelle"); !errors.Is(err, ErrOfficerOnly) {
			t.Fatalf("member naming another = %v", err)
		}
		if recs, err := w.h.svc.AttendanceFor(t.Context(), w.gid, w.member, "sunderfel"); err != nil || len(recs) != 1 {
			t.Fatalf("member naming their own = %+v %v", recs, err)
		}
		if recs, err := w.h.svc.AttendanceFor(t.Context(), w.gid, w.officer, "mirelle"); err != nil || len(recs) != 1 ||
			recs[0].Name != "Mirelle" || recs[0].Chip != chips["Mirelle"] {
			t.Fatalf("officer naming = %+v %v", recs, err)
		}
		if _, err := w.h.svc.AttendanceFor(t.Context(), w.gid, w.officer, "Nobody"); !errors.Is(err, ErrNoSuchMember) {
			t.Fatalf("unknown name = %v", err)
		}
		stranger := seedUser(t, w.h.pool, "stranger@example.com")
		if _, err := w.h.svc.AttendanceFor(t.Context(), w.gid, stranger, ""); !errors.Is(err, ErrNotMember) {
			t.Fatalf("non-member = %v", err)
		}
		if names, err := w.h.svc.AttendanceNames(t.Context(), w.gid, w.member, "", 25); err != nil || len(names) != 0 {
			t.Fatalf("member autocomplete = %v %v", names, err)
		}
	})
	t.Run("readiness", func(t *testing.T) {
		w, raid := readinessRaidWorld(t)
		w.as(w.officer)
		res := w.h.do(http.MethodGet, w.path("/readiness?raid=%d", raid.ID), "")
		if res.StatusCode != http.StatusOK {
			t.Fatalf("readiness = %d", res.StatusCode)
		}
		var board ReadinessView
		w.h.data(res, &board)
		failing := 0
		for _, r := range board.Rows {
			if r.Failing > 0 || r.NoExport {
				failing++
			}
		}
		got, err := w.h.svc.NextRaidReadiness(t.Context(), w.gid, w.officer, time.Now())
		if err != nil || got.Raid == nil || got.Raid.ID != raid.ID || got.Signed != len(board.Rows) || got.Failing != failing ||
			len(got.Rows) != failing || got.More != 0 {
			t.Fatalf("readiness = %+v %v (board %d rows, %d failing)", got, err, len(board.Rows), failing)
		}
		want := map[string]bool{"Kraggor · Synced 10 days before the raid": true, "Mirelle · No export yet": true}
		for _, line := range got.Rows {
			if !want[line] {
				t.Fatalf("row %q, want one of %v", line, want)
			}
		}
	})
}

func TestReadinessGateRunsOnEveryCall(t *testing.T) {
	w, _ := readinessRaidWorld(t)
	ctx := t.Context()
	if _, err := w.h.svc.NextRaidReadiness(ctx, w.gid, w.member, time.Now()); !errors.Is(err, ErrNotOfficer) {
		t.Fatalf("member = %v", err)
	}
	if _, err := w.h.pool.Exec(ctx, `update entitlements set tier = 'guild' where guild_id = $1`, w.gid); err != nil {
		t.Fatal(err)
	}
	if _, err := w.h.svc.NextRaidReadiness(ctx, w.gid, w.officer, time.Now()); !errors.Is(err, ErrRaidGuildRequired) {
		t.Fatalf("guild tier = %v", err)
	}
	w.dropPlan(t)
	if _, err := w.h.svc.NextRaidReadiness(ctx, w.gid, w.officer, time.Now()); !errors.Is(err, ErrPlanRequired) {
		t.Fatalf("lapsed = %v", err)
	}
	w.grantPlan(t)
	w.setContested(t, true)
	if _, err := w.h.svc.NextRaidReadiness(ctx, w.gid, w.officer, time.Now()); !errors.Is(err, ErrClaimContested) {
		t.Fatalf("frozen = %v", err)
	}
	w.setContested(t, false)
	got, err := w.h.svc.NextRaidReadiness(ctx, w.gid, w.officer, time.Now().Add(30*24*time.Hour))
	if err != nil || got.Raid != nil {
		t.Fatalf("none planned = %+v %v", got, err)
	}
}

func TestReadinessListsTenRowsThenTheRest(t *testing.T) {
	rows := []ReadinessRow{}
	for i := 0; i < 12; i++ {
		rows = append(rows, ReadinessRow{Name: "Raider", Failing: 1, Consumables: ConsumablesView{State: "short"}})
	}
	rows = append(rows, ReadinessRow{Name: "Ready"})
	got := readinessAnswerOf(Raid{ID: 1}, ReadinessView{Rows: rows})
	if got.Signed != 13 || got.Failing != 12 || len(got.Rows) != 10 || got.More != 2 || got.Rows[0] != "Raider · Short" {
		t.Fatalf("answer = %+v", got)
	}
}

// TestAttendanceChipIsTheRosterChipVerbatim: the web's own strings (roster-facts.ts
// attendanceChip; attendance-copy.ts rosterChip; nights.ts nightsLabel), in both modes.
func TestAttendanceChipIsTheRosterChipVerbatim(t *testing.T) {
	row := func(p, n int, long *Attendance) RosterRow {
		return RosterRow{Attendance: Attendance{Present: p, Nights: n}, Attendance30: long}
	}
	for _, c := range []struct {
		row    RosterRow
		source string
		want   string
	}{
		{row(5, 8, nil), AttendanceSourceReports, "5 of 8 nights"},
		{row(1, 1, nil), AttendanceSourceReports, "1 of 1 night"},
		{row(0, 0, nil), AttendanceSourceReports, ""},
		{row(7, 8, &Attendance{Present: 26, Nights: 30}), AttendanceSourceRaids, "Came 7 of 8 · 26 of 30 raids"},
		{row(3, 3, &Attendance{Present: 3, Nights: 3}), AttendanceSourceRaids, "Came 3 of 3 raids"},
		{row(0, 0, &Attendance{}), AttendanceSourceRaids, ""},
	} {
		if got := attendanceChip(c.row, c.source); got != c.want {
			t.Errorf("%+v %s = %q, want %q", c.row.Attendance, c.source, got, c.want)
		}
	}
}

// TestAnOfficersNameLookupIsStable: two verified characters of one name (two realms) resolve to
// the same one, the lower key, every time.
func TestAnOfficersNameLookupIsStable(t *testing.T) {
	w, _ := matchedNight(t)
	w.seedRaider(t, "twin@example.com", "member", raidChar{"us/aardvark/mirelle", "Mirelle", "mage", true})
	for i := 0; i < 5; i++ {
		recs, err := w.h.svc.AttendanceFor(t.Context(), w.gid, w.officer, "Mirelle")
		if err != nil || len(recs) != 1 || recs[0].Class != "mage" {
			t.Fatalf("lookup %d = %+v %v", i, recs, err)
		}
	}
}
