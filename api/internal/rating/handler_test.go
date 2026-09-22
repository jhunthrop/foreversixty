// api/internal/rating/handler_test.go
package rating

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func newTestService(t *testing.T) (*Service, *reports.Store) {
	pool := testPool(t)
	reportStore := &reports.Store{Pool: pool}
	return &Service{Store: &Store{Pool: pool}, Reports: reportStore}, reportStore
}

// testOwnerID upserts the one account this file's tests attach every seeded report to and
// returns its id. reports.owner_id references users(id); which account holds a seeded
// report does not matter to any test here, so every call shares the same email.
func testOwnerID(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ('rating-handler-test@example.com')
		 on conflict (email) do update set email = excluded.email
		 returning id`).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// mustCreateReport inserts a minimal report row directly, the same columns
// api/internal/reports/harness_test.go's own seedReport helper writes (that helper is
// unexported in package reports and this file is package rating, so it cannot be called
// directly). status is always 'complete': that is all this handler cares about.
func mustCreateReport(t *testing.T, rs *reports.Store, id, visibility string) {
	t.Helper()
	ownerID := testOwnerID(t, rs.Pool)
	if _, err := rs.Pool.Exec(context.Background(),
		`insert into reports (id, owner_id, title, visibility, zone, status, created_at)
		 values ($1, $2, 'Handler test', $3, 'Blackrock Spire', 'complete', now())
		 on conflict (id) do update set visibility = excluded.visibility, status = excluded.status`,
		id, ownerID, visibility); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		rs.Pool.Exec(context.Background(), `delete from reports where id = $1`, id)
		rs.Pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, id)
	})
}

func TestFightRatingsServesAPublicReportsRoster(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-public-1", reports.Public)
	f := fightFixture("handler-public-1", true,
		summary.RosterRow{GUID: "g1", Name: "Served", Class: "Druid", Spec: "Balance", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-public-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	var env httpx.Envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if !env.OK {
		t.Fatalf("ok = false: %+v", env.Error)
	}
}

func TestFightRatingsIs404ForAPrivateReportToAnAnonymousCaller(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-private-1", reports.Private)
	f := fightFixture("handler-private-1", true,
		summary.RosterRow{GUID: "g1", Name: "Hidden", Class: "Warlock", Spec: "Affliction", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-private-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 (private reports must not distinguish 403 from a missing report)", rec.Code)
	}
}

func TestFightRatingsIs404ForAMissingFight(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-missing-1", reports.Public)
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-missing-1/fights/9/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestCharacterRatingRejectsAnInvalidRegionOrRuleset(t *testing.T) {
	svc, _ := newTestService(t)
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/xx/normal/someone/rating", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 for an invalid region", rec.Code)
	}
}

func TestBestWorstComponentPicksTheHighestAndLowestNonExcludedScores(t *testing.T) {
	raw := json.RawMessage(`[
		{"name":"output","score":80,"weight":1,"basis":"b","excluded":false,"moments":[]},
		{"name":"survival","score":20,"weight":1,"basis":"b","excluded":false,"moments":[]},
		{"name":"utility","score":null,"weight":1,"basis":"b","excluded":true,"reason":"n/a","moments":[]}
	]`)
	best, worst := bestWorstComponent(raw)
	if best != "output" {
		t.Errorf("best = %q, want output", best)
	}
	if worst != "survival" {
		t.Errorf("worst = %q, want survival", worst)
	}
}

func TestBestWorstComponentReturnsEmptyOnUnmarshalFailure(t *testing.T) {
	best, worst := bestWorstComponent(json.RawMessage(`not json`))
	if best != "" || worst != "" {
		t.Errorf("best=%q worst=%q, want both empty on malformed input", best, worst)
	}
}

func TestFightRatingsSetsPrivateCacheControlForANonPublicReport(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-unlisted-1", reports.Unlisted)
	f := fightFixture("handler-unlisted-1", true,
		summary.RosterRow{GUID: "g1", Name: "Unlisted", Class: "Paladin", Spec: "Retribution", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-unlisted-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if got := rec.Header().Get("Cache-Control"); got != "private" {
		t.Errorf("Cache-Control = %q, want private for an unlisted report (spec §5.3)", got)
	}
}
