// api/internal/rating/handler_test.go
package rating

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
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

func TestFightRatingsRejectsAnInvalidFightIndex(t *testing.T) {
	svc, _ := newTestService(t)
	mux := http.NewServeMux()
	Mount(mux, svc)
	for _, n := range []string{"0", "-1", "abc"} {
		req := httptest.NewRequest(http.MethodGet, "/v1/reports/whatever-report/fights/"+n+"/ratings", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("fight index %q: status = %d, want 400", n, rec.Code)
		}
	}
}

// mustCreateGuildReport inserts a minimal guilds row and a report that belongs to it,
// visibility 'guild' - mustCreateReport's own insert never sets guild_id, and reports.
// guild_id references guilds(id), so a guild-visible report needs both rows to exist for
// the foreign key. Returns the report's owner id, so a caller can pick a non-member actor
// id that is guaranteed not to collide with it (visible()'s owner check would otherwise
// make a same-id "stranger" actor visible for the wrong reason - not because they are a
// verified guild member, but because they happen to share the owner's id).
func mustCreateGuildReport(t *testing.T, rs *reports.Store, reportID string) int64 {
	t.Helper()
	ctx := context.Background()
	ownerID := testOwnerID(t, rs.Pool)
	var guildID int64
	if err := rs.Pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', $1)
		 on conflict (region, ruleset, name) do update set name = excluded.name
		 returning id`, "Handler Test Guild "+reportID).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := rs.Pool.Exec(ctx,
		`insert into reports (id, owner_id, guild_id, title, visibility, zone, status, created_at)
		 values ($1, $2, $3, 'Handler guild test', 'guild', 'Blackrock Spire', 'complete', now())
		 on conflict (id) do update set visibility = excluded.visibility, guild_id = excluded.guild_id`,
		reportID, ownerID, guildID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		rs.Pool.Exec(ctx, `delete from reports where id = $1`, reportID)
		rs.Pool.Exec(ctx, `delete from rating_scores where report_id = $1`, reportID)
		rs.Pool.Exec(ctx, `delete from guilds where id = $1`, guildID)
	})
	return ownerID
}

func TestFightRatingsIsVisibleToAVerifiedGuildMember(t *testing.T) {
	svc, rs := newTestService(t)
	ownerID := mustCreateGuildReport(t, rs, "handler-guild-1")
	f := fightFixture("handler-guild-1", true,
		summary.RosterRow{GUID: "g1", Name: "Officer", Class: "Priest", Spec: "Shadow", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	svc.Accounts = stubAccounts{rank: "member", ok: true}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-guild-1/fights/1/ratings", nil)
	// ownerID+1, not a hardcoded id: visible() already returns true for the owner
	// regardless of stubAccounts, so this actor must be guaranteed distinct from the
	// owner for the test to actually exercise the guild-membership branch rather than the
	// owner branch.
	req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: ownerID + 1, Role: "user", Method: "session"}))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s (a verified guild member must see a guild report's ratings)",
			rec.Code, rec.Body.String())
	}
}

func TestFightRatingsIs404ForAGuildReportToANonMember(t *testing.T) {
	svc, rs := newTestService(t)
	ownerID := mustCreateGuildReport(t, rs, "handler-guild-2")
	f := fightFixture("handler-guild-2", true,
		summary.RosterRow{GUID: "g1", Name: "Stranger", Class: "Priest", Spec: "Shadow", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	svc.Accounts = stubAccounts{rank: "", ok: false}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-guild-2/fights/1/ratings", nil)
	req = req.WithContext(auth.WithActor(req.Context(), auth.Actor{UserID: ownerID + 1, Role: "user", Method: "session"}))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 for a non-member of a guild report", rec.Code)
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

// anonymize marks playerKey's owning account as anonymized, creating the account and the
// character link if they do not already exist, mirroring store_test.go's own
// TestReadCharacterRatingExcludesAnonymizedPlayers setup so both files stay consistent.
func anonymize(t *testing.T, pool *pgxpool.Pool, email, playerKey, region, ruleset, name string) {
	t.Helper()
	ctx := context.Background()
	var uid int64
	if err := pool.QueryRow(ctx,
		`insert into users (email, anonymize) values ($1, true)
		 on conflict (email) do update set anonymize = true
		 returning id`, email).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values ($1, $2, $3, $4, $5)
		 on conflict (key) do update set user_id = $5`,
		playerKey, region, ruleset, name, uid); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key = $1`, playerKey)
		pool.Exec(ctx, `delete from users where email = $1`, email)
	})
}

func TestVisiblePlayersOmitsAnonymizedPlayersRows(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-anon-1", reports.Public)
	f := fightFixture("handler-anon-1", true,
		summary.RosterRow{GUID: "g1", Name: "Seen", Class: "Hunter", Spec: "Marksmanship", Role: "dps"},
		summary.RosterRow{GUID: "g2", Name: "Hidden", Class: "Rogue", Spec: "Subtlety", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	anonymize(t, rs.Pool, "rating-handler-anon@example.com", "us/normal/hidden", "us", "normal", "Hidden")

	rows, _, err := svc.Store.ReadFightRatings(context.Background(), "handler-anon-1", 1)
	if err != nil {
		t.Fatal(err)
	}
	players, err := svc.visiblePlayers(context.Background(), rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(players) != 1 || players[0].PlayerName != "Seen" {
		t.Fatalf("visiblePlayers = %+v, want exactly the one non-anonymized row (spec §5.1: "+
			"an anonymized player's row is omitted entirely)", players)
	}
}

func TestFightRatingsResponseOmitsAnAnonymizedPlayersRow(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-anon-2", reports.Public)
	f := fightFixture("handler-anon-2", true,
		summary.RosterRow{GUID: "g1", Name: "Visible", Class: "Druid", Spec: "Feral", Role: "dps"},
		summary.RosterRow{GUID: "g2", Name: "Ghost", Class: "Mage", Spec: "Fire", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	anonymize(t, rs.Pool, "rating-handler-anon2@example.com", "us/normal/ghost", "us", "normal", "Ghost")

	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-anon-2/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	var env struct {
		Data fightRatingsDTO `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if len(env.Data.Players) != 1 || env.Data.Players[0].PlayerName != "Visible" {
		t.Fatalf("players = %+v, want exactly the one non-anonymized row", env.Data.Players)
	}
}

func TestCharacterRatingServesATrendForAPublicReport(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-char-1", reports.Public)
	f := fightFixture("handler-char-1", true,
		summary.RosterRow{GUID: "g1", Name: "Trendy", Class: "Shaman", Spec: "Elemental", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/Trendy/rating", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	var env struct {
		Data characterRatingDTO `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Data.PlayerKey != "us/normal/trendy" {
		t.Errorf("player_key = %q, want us/normal/trendy", env.Data.PlayerKey)
	}
	if env.Data.SampleSize != 1 || len(env.Data.Trend) != 1 {
		t.Fatalf("sample_size=%d len(trend)=%d, want 1 and 1", env.Data.SampleSize, len(env.Data.Trend))
	}
	if env.Data.Latest == nil || env.Data.Latest.PlayerName != "Trendy" {
		t.Fatalf("latest = %+v, want the one rated fight", env.Data.Latest)
	}
}

func TestCharacterRatingIs404WhenNoRatedFightsExist(t *testing.T) {
	svc, _ := newTestService(t)
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/nobodyhome/rating", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 for a character with no rated fights", rec.Code)
	}
}

func TestCharacterRatingIs404ForAnAnonymizedCharacter(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-char-anon-1", reports.Public)
	f := fightFixture("handler-char-anon-1", true,
		summary.RosterRow{GUID: "g1", Name: "Ducking", Class: "Warlock", Spec: "Demonology", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	anonymize(t, rs.Pool, "rating-handler-char-anon@example.com", "us/normal/ducking", "us", "normal", "Ducking")

	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/Ducking/rating", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 for an anonymized character (must read the same as no such character)", rec.Code)
	}
}

func TestTrendCursorRoundTripsThroughEncodeAndDecode(t *testing.T) {
	want := cursorPos{
		FoughtAt:   time.Date(2026, 12, 9, 1, 2, 3, 0, time.UTC),
		ReportID:   "handler-cursor-1",
		FightIndex: 4,
	}
	got, ok := decodeTrendCursor(encodeTrendCursor(want))
	if !ok {
		t.Fatal("decode of a cursor this package just encoded reported ok = false")
	}
	if !got.FoughtAt.Equal(want.FoughtAt) || got.ReportID != want.ReportID || got.FightIndex != want.FightIndex {
		t.Errorf("round trip = %+v, want %+v", got, want)
	}
}

func TestDecodeTrendCursorRejectsWhatThisPackageDidNotEncode(t *testing.T) {
	cases := []string{"", "not-base64!!", base64.RawURLEncoding.EncodeToString([]byte("too|few")), base64.RawURLEncoding.EncodeToString([]byte("bad-time|r1|3"))}
	for _, c := range cases {
		if _, ok := decodeTrendCursor(c); ok {
			t.Errorf("decodeTrendCursor(%q) = ok, want rejected", c)
		}
	}
}

func TestCharacterRatingSetsPublicCacheControl(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-char-cache-1", reports.Public)
	f := fightFixture("handler-char-cache-1", true,
		summary.RosterRow{GUID: "g1", Name: "Cached", Class: "Hunter", Spec: "Survival", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/us/normal/Cached/rating", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	want := "public, max-age=" + strconv.Itoa(cacheSeconds)
	if got := rec.Header().Get("Cache-Control"); got != want {
		t.Errorf("Cache-Control = %q, want %q (spec §5.3)", got, want)
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
