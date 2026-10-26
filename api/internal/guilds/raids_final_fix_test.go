package guilds

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

// foreignRaid is a raid of another guild, in the given zone and status.
func (w raidWorld) foreignRaid(t *testing.T, zone, status string) int64 {
	t.Helper()
	other := seedGuild(t, w.h.pool, "Foreign "+status+zone)
	var id int64
	if err := w.h.pool.QueryRow(t.Context(), `
		insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
		values ($1, $2, 'Elsewhere', $3, 180, $4, $5) returning id`,
		other, zone, pastStart(), status, w.officer).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func TestAForeignRaidIsRefusedOnTheLootAndReadinessRoutes(t *testing.T) {
	w, _ := matchedNight(t)
	foreign := w.foreignRaid(t, raidZoneOny, RaidDone)
	w.as(w.officer)
	expectError(t, w.postAward(&foreign, officerKey), http.StatusBadRequest, "invalid")
	expectError(t, w.h.do(http.MethodGet, w.path("/readiness?raid=%d", foreign), ""), http.StatusNotFound, "not_found")
}

func TestAnAwardNamesTheRaidItRefuses(t *testing.T) {
	w, _ := matchedNight(t)
	elsewhere := w.raidAt(t, "raid:barrow-deeps", pastStart().Add(-24*time.Hour))
	w.exec(t, `update guild_raids set status = 'done' where id = `+strconv.FormatInt(elsewhere.ID, 10))
	w.as(w.officer)
	res := w.postAward(&elsewhere.ID, officerKey)
	body := readBody(res)
	if res.StatusCode != http.StatusBadRequest || !strings.Contains(body, strconv.FormatInt(elsewhere.ID, 10)) || !strings.Contains(body, `"invalid"`) {
		t.Fatalf("award for a done raid of another zone = %d: %s", res.StatusCode, body)
	}
}

func TestReadinessForADoneOrCancelledRaidIs404(t *testing.T) {
	w, night := matchedNight(t)
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodGet, w.path("/readiness?raid=%d", night.ID), ""), http.StatusNotFound, "not_found")
	cancelled := w.createRaid(t)
	w.exec(t, `update guild_raids set status = 'cancelled' where id = `+strconv.FormatInt(cancelled.ID, 10))
	expectError(t, w.h.do(http.MethodGet, w.path("/readiness?raid=%d", cancelled.ID), ""), http.StatusNotFound, "not_found")
}

func TestAStaleOnlyRowsNudgeNamesTheStaleSync(t *testing.T) {
	w, raid := readinessRaidWorld(t)
	w.as(w.officer)
	var v ReadinessView
	w.h.data(w.h.do(http.MethodGet, w.path("/readiness?raid=%d", raid.ID), ""), &v)
	row := readinessRowsByKey(v)[memberKeyA]
	if row.NudgeText == nil || strings.Contains(*row.NudgeText, "every readiness check passes") ||
		!strings.Contains(*row.NudgeText, "Kraggor:") || !strings.Contains(*row.NudgeText, "last sync was 10 days before the raid") {
		t.Fatalf("nudge = %v", row.NudgeText)
	}
}

// R-A1: the plan gates writes; an officer without it still reads everything.
func TestAnOfficerWithoutThePlanReadsEveryRowAndTheFullReport(t *testing.T) {
	w, night, _ := reportNight(t)
	w.dropPlan(t)
	w.as(w.officer)
	v := w.getRaid(t, night.ID)
	if got := attendanceKeys(v); len(got) != 7 {
		t.Fatalf("officer without the plan sees %d rows, want 7", len(got))
	}
	if v.Report == nil || len(v.Report.Deaths) != 2 || len(v.Report.Moves.Rows) != 2 {
		t.Fatalf("officer's report = %+v", v.Report)
	}
	w.as(w.member)
	if got := attendanceKeys(w.getRaid(t, night.ID)); len(got) != 3 {
		t.Fatalf("a member sees %d rows, want their own 3", len(got))
	}
}

func TestPostedAtIsUTC(t *testing.T) {
	w, night, _ := postWorld(t)
	w.as(w.officer)
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("post = %d", res.StatusCode)
	}
	rep := w.getRaid(t, night.ID).Report
	if rep.PostedAt == nil || rep.PostedAt.Location() != time.UTC {
		t.Fatalf("posted_at = %v", rep.PostedAt)
	}
}

func TestAHardPostErrorReleasesTheClaim(t *testing.T) {
	w, night, events := postWorld(t)
	events.reportErr = errors.New("boom")
	w.as(w.officer)
	expectError(t, w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""), http.StatusInternalServerError, "internal")
	if rep := w.getRaid(t, night.ID).Report; rep.PostedAt != nil {
		t.Fatalf("posted_at kept after a hard error: %v", rep.PostedAt)
	}
	events.reportErr = nil
	if res := w.h.do(http.MethodPost, w.path("/raids/%d/report/post", night.ID), ""); res.StatusCode != http.StatusOK {
		t.Fatalf("retry = %d: %s", res.StatusCode, readBody(res))
	}
}

// A release only clears the claim it was handed: not another guild's raid, not a newer claim.
func TestReleasingAClaimIsScopedToTheGuildAndTheClaim(t *testing.T) {
	w, night, _ := postWorld(t)
	store := w.h.store
	at, ok, err := store.ClaimReportPost(t.Context(), w.gid, night.ID)
	if err != nil || !ok {
		t.Fatalf("claim = %v, %v", ok, err)
	}
	if err := store.ReleaseReportPost(t.Context(), w.gid+1000, night.ID, at); err != nil {
		t.Fatal(err)
	}
	if err := store.ReleaseReportPost(t.Context(), w.gid, night.ID, at.Add(-time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := store.ClaimReportPost(t.Context(), w.gid, night.ID); ok {
		t.Fatal("a release for another guild or claim cleared the stamp")
	}
	if err := store.ReleaseReportPost(t.Context(), w.gid, night.ID, at); err != nil {
		t.Fatal(err)
	}
	if _, ok, _ := store.ClaimReportPost(t.Context(), w.gid, night.ID); !ok {
		t.Fatal("the matching release did not clear the stamp")
	}
}
