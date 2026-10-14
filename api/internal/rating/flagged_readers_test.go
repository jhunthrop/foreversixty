package rating

// Task L6.12: a report flagged tampered or unverified counts nowhere a player looks, and the
// backfill stops rating one. The rating endpoints' character history (ReadCharacterRating:
// the fight cards and trend) and the backfill's pick of unrated fights (staleFights) select
// their reports through presence.EvidenceReport.

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The character's rating history never lists a flagged report's card.
func TestReadCharacterRatingLeavesOutAFlaggedReport(t *testing.T) {
	fx := seedLoggedRaider(t)
	ctx := context.Background()
	flagged := sv2Report + "c"
	inReport := func() int {
		t.Helper()
		rows, _, err := fx.store.ReadCharacterRating(ctx, fx.main, characterTrendPageSize, nil)
		if err != nil {
			t.Fatal(err)
		}
		n := 0
		for _, r := range rows {
			if r.ReportID == flagged {
				n++
			}
		}
		return n
	}
	if n := inReport(); n != 4 {
		t.Fatalf("before the flag the history lists %d of the report's cards, want its 4", n)
	}
	flagReport(t, fx, flagged)
	if n := inReport(); n != 0 {
		t.Errorf("the history lists %d cards of a flagged report, want none", n)
	}
}

// The backfill never rates a flagged report's unrated fights; the same fight in an
// unflagged report is picked.
func TestStaleFightsSkipsAFlaggedReportsUnratedFights(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	t.Cleanup(func() {
		pool.Exec(context.Background(), `delete from reports where id like '000l612-%'`)
	})
	at := time.Now().Add(-24 * time.Hour)
	for _, r := range []struct{ id, flagged string }{{"000l612-clean", ""}, {"000l612-flagged", "tampered"}} {
		mustExec(t, pool, `insert into reports (id, title, visibility, zone, status, flagged, logging_character)
			values ($1, 'L6.12', 'public', 'Molten Core', 'complete', nullif($2, ''), 'us/normal/l612logger')`,
			r.id, r.flagged)
		mustExec(t, pool, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, kill, start_ms)
			values ($1, 1, 990699, 'L6.12 boss', 3, true, $2)`, r.id, at.UnixMilli())
	}
	picked, err := (&Store{Pool: pool}).staleFights(ctx, BackfillBatchSize)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, sf := range picked {
		got[sf.ReportID] = true
	}
	if !got["000l612-clean"] {
		t.Fatalf("the clean report's unrated fight was not picked: %+v", picked)
	}
	if got["000l612-flagged"] {
		t.Errorf("the flagged report's unrated fight was picked for rating")
	}
}

// The per-fight ratings endpoint answers a flagged report's fight 404 "flagged": its cards are
// not evidence. The same fight in the same report, unflagged, is served.
func TestFightRatingsIs404FlaggedForAFlaggedReport(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-flagged-1", reports.Public)
	f := fightFixture("handler-flagged-1", true,
		summary.RosterRow{GUID: "g1", Name: "Forged", Class: "Rogue", Spec: "Combat", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	get := func() *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-flagged-1/fights/1/ratings", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		return rec
	}
	if rec := get(); rec.Code != http.StatusOK {
		t.Fatalf("before the flag: status = %d, body = %s", rec.Code, rec.Body.String())
	}
	mustExec(t, rs.Pool, `update reports set flagged = 'tampered' where id = 'handler-flagged-1'`)
	rec := get()
	var env httpx.Envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusNotFound || env.Error == nil || env.Error.Code != "flagged" {
		t.Fatalf("status = %d, body = %s; want 404 with error code \"flagged\"", rec.Code, rec.Body.String())
	}
}

// A flagged report the caller may not open answers the plain not_found, never "flagged":
// the flag is asked only once the report is visible, so it says nothing about a report the
// caller could not see.
func TestFightRatingsIsPlainNotFoundForAFlaggedReportTheCallerCannotOpen(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-flagged-private-1", reports.Private)
	f := fightFixture("handler-flagged-private-1", true,
		summary.RosterRow{GUID: "g1", Name: "Hidden", Class: "Rogue", Spec: "Combat", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mustExec(t, rs.Pool, `update reports set flagged = 'tampered' where id = 'handler-flagged-private-1'`)
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-flagged-private-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var env httpx.Envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusNotFound || env.Error == nil || env.Error.Code != "not_found" {
		t.Fatalf("status = %d, body = %s; want 404 with error code \"not_found\"", rec.Code, rec.Body.String())
	}
}
