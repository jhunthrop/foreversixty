package coach

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

var sampleToday = time.Date(2026, 12, 10, 0, 0, 0, 0, time.UTC)

const sampleSecret = "sk-secret-key"

func sampleRand() *rand.Rand { return rand.New(rand.NewPCG(1, 2)) }

func sampleLogger(buf *bytes.Buffer) *slog.Logger { return slog.New(slog.NewJSONHandler(buf, nil)) }

// seedReviews inserts n reviews per spec generated from the given instant, each with a pack and
// a rating row naming the character, and removes them at cleanup.
func seedReviews(t *testing.T, pool *pgxpool.Pool, at time.Time, perSpec map[string]int) {
	t.Helper()
	if err := db.EnsurePartitions(context.Background(), pool, "rating_scores", at, 1); err != nil {
		t.Fatal(err)
	}
	for spec, n := range perSpec {
		for i := 0; i < n; i++ {
			seedReview(t, pool, at.Add(time.Duration(i)*time.Minute), spec, i)
		}
	}
}

func seedReview(t *testing.T, pool *pgxpool.Pool, at time.Time, spec string, i int) {
	t.Helper()
	ctx := context.Background()
	rid := fmt.Sprintf("t37-%s-%s-%d", at.Format("0102"), spec, i)
	key := "us/normal/" + rid
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from coach_review_samples where report_id = $1`, rid)
		pool.Exec(ctx, `delete from rating_scores where report_id = $1`, rid)
		pool.Exec(ctx, `delete from reports where id = $1`, rid)
	})
	exec := func(q string, args ...any) {
		t.Helper()
		if _, err := pool.Exec(ctx, q, args...); err != nil {
			t.Fatalf("%s: %v", q, err)
		}
	}
	exec(`insert into reports (id, visibility, status) values ($1, 'public', 'complete')`, rid)
	exec(`insert into coach_packs (report_id, fight_index, player_key, pack_version, pack)
	      values ($1, 1, $2, 'p1', jsonb_build_object('context', jsonb_build_object('spec', $3::text)))`, rid, key, spec)
	exec(`insert into coach_reviews (report_id, fight_index, player_key, pack_version, model_version, persona_version,
	        review, source, validation, generated_at, gateway_request_id)
	      values ($1, 1, $2, 'p1', 'm1', 'v1', '{"opening":"x"}', 'model', '["a problem"]', $3, $4)`, rid, key, at, sampleSecret)
	exec(`insert into rating_scores (report_id, fight_index, player_key, player_name, kill, components, model_version, fought_at)
	      values ($1, 1, $2, 'Name'||$3::text, true, '{}', 'x', $4)`, rid, key, spec, at)
}

func countSamples(t *testing.T, pool *pgxpool.Pool, day time.Time) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select count(*) from coach_review_samples where day = $1 and report_id like 't37-%'`, day).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestSampleDay_TwentySpreadAcrossSpecs(t *testing.T) {
	pool := writerPool(t)
	seedReviews(t, pool, sampleToday.Add(time.Hour), map[string]int{"a": 15, "b": 15, "c": 15, "d": 15})
	seedReviews(t, pool, sampleToday.Add(-12*time.Hour), map[string]int{"e": 5})
	n, err := SampleDay(context.Background(), pool, sampleToday, sampleRand(), slog.Default())
	if err != nil || n != SampleSize || countSamples(t, pool, sampleToday) != SampleSize {
		t.Fatalf("n=%d err=%v", n, err)
	}
	rows, err := pool.Query(context.Background(),
		`select spec, count(*) from coach_review_samples where day = $1 and report_id like 't37-%' group by spec`, sampleToday)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var spec string
		var c int
		if err := rows.Scan(&spec, &c); err != nil || spec == "e" || c != 5 {
			t.Fatalf("spec %s count %d err %v", spec, c, err)
		}
	}
}

func TestSampleDay_FewerThanTwentyTakesAll(t *testing.T) {
	pool := writerPool(t)
	seedReviews(t, pool, sampleToday.Add(time.Hour), map[string]int{"warrior-fury": 4, "priest-holy": 3})
	n, err := SampleDay(context.Background(), pool, sampleToday, sampleRand(), slog.Default())
	if err != nil || n != 7 || countSamples(t, pool, sampleToday) != 7 {
		t.Fatalf("n=%d err=%v", n, err)
	}
}

func TestSampleDay_Idempotent(t *testing.T) {
	pool := writerPool(t)
	seedReviews(t, pool, sampleToday.Add(time.Hour), map[string]int{"a": 15, "b": 15})
	var buf bytes.Buffer
	first, err := SampleDay(context.Background(), pool, sampleToday, DaySeed(sampleToday), sampleLogger(&buf))
	if err != nil || first != SampleSize {
		t.Fatalf("first n=%d err=%v", first, err)
	}
	buf.Reset()
	second, err := SampleDay(context.Background(), pool, sampleToday, DaySeed(sampleToday), sampleLogger(&buf))
	if err != nil || second != 0 || countSamples(t, pool, sampleToday) != SampleSize || buf.Len() != 0 {
		t.Fatalf("second n=%d err=%v rows=%d log=%q", second, err, countSamples(t, pool, sampleToday), buf.String())
	}
}

func TestSampleDay_SameSeedSamePick(t *testing.T) {
	pool := writerPool(t)
	seedReviews(t, pool, sampleToday.Add(time.Hour), map[string]int{"a": 15, "b": 15})
	ctx := context.Background()
	pick := func() string {
		pool.Exec(ctx, `delete from coach_review_samples where report_id like 't37-%'`)
		if _, err := SampleDay(ctx, pool, sampleToday, DaySeed(sampleToday), slog.Default()); err != nil {
			t.Fatal(err)
		}
		var s string
		pool.QueryRow(ctx, `select string_agg(report_id, ',' order by report_id) from coach_review_samples where report_id like 't37-%'`).Scan(&s)
		return s
	}
	if a, b := pick(), pick(); a != b {
		t.Fatalf("picks differ:\n%s\n%s", a, b)
	}
}

func TestSampleDay_LogsOneEntryEachWithoutPersonalData(t *testing.T) {
	pool := writerPool(t)
	seedReviews(t, pool, sampleToday.Add(time.Hour), map[string]int{"a": 15, "b": 15})
	var buf bytes.Buffer
	if _, err := SampleDay(context.Background(), pool, sampleToday, sampleRand(), sampleLogger(&buf)); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if got := strings.Count(out, "\n"); got != SampleSize {
		t.Fatalf("%d entries, want %d", got, SampleSize)
	}
	for _, banned := range []string{sampleSecret, "Name", "t37-", "us/normal", "opening", "a problem"} {
		if strings.Contains(out, banned) {
			t.Fatalf("log carries %q: %s", banned, out)
		}
	}
}

func TestSamplesRoute_ModeratorsOnly(t *testing.T) {
	pool := writerPool(t)
	seedReviews(t, pool, sampleToday.Add(time.Hour), map[string]int{"a": 15, "b": 15})
	if _, err := SampleDay(context.Background(), pool, sampleToday, sampleRand(), slog.Default()); err != nil {
		t.Fatal(err)
	}
	h := SamplesHandler(pool)
	get := func(a auth.Actor, query string) *httptest.ResponseRecorder {
		r := httptest.NewRequest(http.MethodGet, "/v1/admin/coach/samples"+query, nil)
		w := httptest.NewRecorder()
		h(w, r.WithContext(auth.WithActor(r.Context(), a)))
		return w
	}
	w := get(auth.Actor{UserID: 1, Role: "moderator"}, "?day=2026-12-10")
	if w.Code != 200 || w.Header().Get("Cache-Control") != "private, no-store" ||
		strings.Count(w.Body.String(), `"character_name"`) != SampleSize {
		t.Fatalf("moderator: %d %v", w.Code, w.Header())
	}
	if c := get(auth.Actor{UserID: 2, Role: "user"}, "?day=2026-12-10").Code; c != 403 {
		t.Fatalf("member %d", c)
	}
	if c := get(auth.Actor{}, "?day=2026-12-10").Code; c != 401 {
		t.Fatalf("signed out %d", c)
	}
	if c := get(auth.Actor{UserID: 1, Role: "admin"}, "?day=nope").Code; c != 400 {
		t.Fatalf("bad day %d", c)
	}
}
