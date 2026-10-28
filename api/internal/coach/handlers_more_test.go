package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// t31 is the fight route's world (t30) with the night, recent and history routes mounted.
type t31 struct{ *t30 }

func newT31(t *testing.T) *t31 {
	t.Helper()
	f := &t31{newT30(t)}
	if _, err := f.pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	MountLines(f.mux, f.svc)
	return f
}

// rate makes fight a rated pull of key at foughtAt: the fight row and summary (copies of
// fight 1's) and the rating_scores row.
func (f *t31) rate(t *testing.T, key string, fight int, foughtAt time.Time) {
	t.Helper()
	ctx := context.Background()
	if fight != t30Fight {
		if _, err := f.pool.Exec(ctx, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, size, kill, duration_ms, start_ms)
			select report_id, $2, encounter_id, name, difficulty, size, kill, duration_ms, start_ms + $2 * 600000
			  from fights where report_id = $1 and fight_index = 1 on conflict do nothing`, t30Report, fight); err != nil {
			t.Fatal(err)
		}
		k := store.Keys{ReportID: t30Report}
		f.bucket.mu.Lock()
		f.bucket.files[k.FightSummary(fight)] = f.bucket.files[k.FightSummary(t30Fight)]
		f.bucket.mu.Unlock()
	}
	if err := db.EnsureRatingsPartition(ctx, f.pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	if _, err := f.pool.Exec(ctx, `insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, role,
		   encounter_id, difficulty, kill, overall, overall_uncapped, components, model_version, fought_at, deaths)
		values ($1, $2, $3, 'Tyrim', 'Warrior', 'Fury', 'dps', $4, 3, true, 50, 50, '[]', 'test', $5, 0)`,
		t30Report, fight, key, onyxiaEncounterID, foughtAt); err != nil {
		t.Fatal(err)
	}
}

// call GETs path as actor and decodes the envelope: its data, or its error.
func (f *t31) call(t *testing.T, actor auth.Actor, path string) (int, http.Header, any) {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	req = req.WithContext(auth.WithActor(req.Context(), actor))
	rec := httptest.NewRecorder()
	f.mux.ServeHTTP(rec, req)
	var env map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatalf("%d %s", rec.Code, rec.Body.String())
	}
	if e, ok := env["error"]; ok && e != nil {
		return rec.Code, rec.Header(), e
	}
	return rec.Code, rec.Header(), env["data"]
}

func asMap(t *testing.T, v any) map[string]any {
	t.Helper()
	m, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("not an object: %#v", v)
	}
	return m
}

func (f *t31) nightPath() string {
	return "/v1/reports/" + t30Report + "/coach/night?player=" + f.furyKey
}

func (f *t31) historyPath(cursor string) string {
	return "/v1/characters/" + f.furyKey + "/coach/history?cursor=" + cursor
}

func TestNightRoute_PremiumOwnerOnly(t *testing.T) {
	f := newT31(t)
	at := capsNow.Add(-24 * time.Hour)
	f.rate(t, f.furyKey, 1, at)
	f.rate(t, f.furyKey, 2, at.Add(time.Hour))
	code, h, body := f.call(t, user(f.premium), f.nightPath())
	m := asMap(t, body)
	if code != http.StatusOK || m["stem"] == "" || m["of"] != 2.0 || m["on"] != 2.0 || m["fight_index"] != 1.0 {
		t.Fatalf("%d %v", code, m)
	}
	if m["finding_count"].(float64) < 2 || h.Get("Cache-Control") != "private, no-store" {
		t.Fatalf("%v %q", m, h.Get("Cache-Control"))
	}
	if code, _, body := f.call(t, user(f.freeOwner), f.nightPath()); code != http.StatusForbidden || asMap(t, body)["code"] != codeNoPlan {
		t.Fatalf("free owner: %d %v", code, body)
	}
	if code, _, body := f.call(t, user(f.stranger), f.nightPath()); code != http.StatusNotFound || asMap(t, body)["code"] != codeNotFound {
		t.Fatalf("stranger: %d %v", code, body)
	}
}

func (f *t31) recent(t *testing.T, userID int64) map[string]any {
	t.Helper()
	code, h, body := f.call(t, user(userID), "/v1/me/coach/recent")
	if code != http.StatusOK || h.Get("Cache-Control") != "private, no-store" {
		t.Fatalf("%d %v", code, body)
	}
	return asMap(t, body)
}

func TestRecent_LastRaidWithin14Days(t *testing.T) {
	f := newT31(t)
	f.linkCharacter(t, f.premium, true)
	f.rate(t, f.furyKey, 1, capsNow.Add(-10*24*time.Hour))
	raid := asMap(t, f.recent(t, f.premium)["last_raid"])
	if raid["report_id"] != t30Report || raid["url"] == "" {
		t.Fatalf("%v", raid)
	}
	if _, err := f.pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	f.rate(t, f.furyKey, 1, capsNow.Add(-15*24*time.Hour))
	if got := f.recent(t, f.premium)["last_raid"]; got != nil {
		t.Fatalf("a raid 15 days ago is no last raid: %v", got)
	}
}

func TestRecent_FreeGetsCountAndFirstNoReviews(t *testing.T) {
	f := newT31(t)
	at := capsNow.Add(-2 * 24 * time.Hour)
	f.rate(t, f.furyKey, 1, at)
	f.rate(t, f.furyKey, 2, at.Add(time.Hour))
	f.seedReviews(t, 3)
	body := f.recent(t, f.freeOwner)
	raid := asMap(t, body["last_raid"])
	first := asMap(t, raid["first"])
	if first["stem"] == "" || first["of"] != 2.0 || body["pending"] != false {
		t.Fatalf("%v", body)
	}
	one, _ := Night([]NightFight{{FightIndex: 1, Pack: f.pack(t)}})
	if raid["finding_count"] != float64(2*one.FindingCount) {
		t.Fatalf("finding_count %v, want twice %d", raid["finding_count"], one.FindingCount)
	}
	if reviews, ok := body["reviews"].([]any); !ok || len(reviews) != 0 {
		t.Fatalf("a free account gets []: %#v", body["reviews"])
	}
}

func (f *t31) pack(t *testing.T) Pack {
	t.Helper()
	p, err := f.svc.PackFor(context.Background(), f.rep(t), t30Fight, f.furyKey)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

// seedReviews stores n reviews of the Fury character, fight i generated i minutes after
// capsNow, so the newest is fight n.
func (f *t31) seedReviews(t *testing.T, n int) {
	t.Helper()
	p := f.pack(t)
	for i := 1; i <= n; i++ {
		if err := f.svc.Store.SavePack(context.Background(), t30Report, i, f.furyKey, p); err != nil {
			t.Fatal(err)
		}
		k := NewReviewKey(t30Report, i, f.furyKey, p, "m", fixturePersona(t))
		doc := exampleReview()
		if err := f.svc.Store.SaveReview(context.Background(), StoredReview{Key: k, Doc: doc, Source: SourceRules,
			GeneratedAt: capsNow.Add(time.Duration(i) * time.Minute)}); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRecent_PremiumGetsTheNewest20Reviews(t *testing.T) {
	f := newT31(t)
	f.linkCharacter(t, f.premium, true)
	f.seedReviews(t, 25)
	reviews := f.recent(t, f.premium)["reviews"].([]any)
	if len(reviews) != historyPageSize {
		t.Fatalf("%d reviews", len(reviews))
	}
	first := asMap(t, reviews[0])
	if first["fight_index"] != 25.0 || asMap(t, reviews[19])["fight_index"] != 6.0 {
		t.Fatalf("newest first: %v … %v", first, reviews[19])
	}
	if first["boss"] == "" || first["spec"] == "" || first["kill"] != true {
		t.Fatalf("the pack's boss, spec and kill: %v", first)
	}
	if fix, ok := first["fix_first"].(map[string]any); !ok || fix["title"] == "" || len(fix) != 1 {
		t.Fatalf("fix_first is {title}: %v", first)
	}
}

// blockingBucket answers no read until release is closed.
type blockingBucket struct{ release chan struct{} }

func (b blockingBucket) Get(ctx context.Context, _ string) (io.ReadCloser, error) {
	select {
	case <-b.release:
	case <-ctx.Done():
	}
	return nil, fmt.Errorf("blocked")
}

func (f *t31) blockBucket(t *testing.T) {
	t.Helper()
	b := blockingBucket{release: make(chan struct{})}
	t.Cleanup(func() { close(b.release); time.Sleep(50 * time.Millisecond) })
	f.svc.Summaries = b
}

func TestRecent_ColdCacheIsPendingNotSlow(t *testing.T) {
	f := newT31(t)
	f.linkCharacter(t, f.premium, true)
	f.rate(t, f.furyKey, 1, capsNow.Add(-24*time.Hour))
	f.blockBucket(t)
	began := time.Now()
	body := f.recent(t, f.premium)
	if took := time.Since(began); took > recentBudget+time.Second {
		t.Fatalf("took %s", took)
	}
	if body["pending"] != true || asMap(t, body["last_raid"])["first"] != nil {
		t.Fatalf("%v", body)
	}
}

func TestHistory_OwnerPremiumPaged(t *testing.T) {
	f := newT31(t)
	f.owners[f.premium] = map[string]bool{f.furyKey: true}
	f.seedReviews(t, 25)
	code, _, body := f.call(t, user(f.premium), f.historyPath(""))
	page := asMap(t, body)
	if code != http.StatusOK || len(page["reviews"].([]any)) != 20 || page["next_cursor"] == "" {
		t.Fatalf("%d %v", code, page)
	}
	item := asMap(t, page["reviews"].([]any)[0])
	for _, field := range []string{"report_id", "fight_index", "player_key", "spec", "boss", "kill", "source", "generated_at", "fix_first"} {
		if _, ok := item[field]; !ok {
			t.Fatalf("history item lacks %s: %v", field, item)
		}
	}
	_, _, body = f.call(t, user(f.premium), f.historyPath(page["next_cursor"].(string)))
	page = asMap(t, body)
	rest := page["reviews"].([]any)
	if len(rest) != 5 || page["next_cursor"] != "" || asMap(t, rest[0])["fight_index"] != 5.0 {
		t.Fatalf("%v", page)
	}
}

func TestHistory_BadCursorIs400(t *testing.T) {
	f := newT31(t)
	if code, _, _ := f.call(t, user(f.premium), f.historyPath("not-a-cursor")); code != http.StatusBadRequest {
		t.Fatalf("%d", code)
	}
}

func TestHistory_NoPlanIs403(t *testing.T) {
	f := newT31(t)
	f.seedReviews(t, 2)
	code, _, body := f.call(t, user(f.freeOwner), f.historyPath(""))
	m := asMap(t, body)
	if code != http.StatusForbidden || m["code"] != codeNoPlan || m["reviews"] != nil {
		t.Fatalf("%d %v", code, m)
	}
}

func TestHistory_NotOwnerIs404(t *testing.T) {
	f := newT31(t)
	f.ent.coach[f.stranger] = true
	code, _, body := f.call(t, user(f.stranger), f.historyPath(""))
	_, _, unknown := f.call(t, user(f.stranger), "/v1/characters/us/normal/nobody/coach/history")
	if code != http.StatusNotFound || asMap(t, body)["code"] != asMap(t, unknown)["code"] {
		t.Fatalf("%d %v vs %v", code, body, unknown)
	}
}

// ownCharacters makes the premium account own n characters of the report's roster (the
// Fury warrior first), each with a rated pull of fight 1.
func (f *t31) ownCharacters(t *testing.T, n int) {
	t.Helper()
	f.linkCharacter(t, f.premium, true)
	f.rate(t, f.furyKey, 1, capsNow.Add(-time.Hour))
	for i := 1; i < n; i++ {
		key := fmt.Sprintf("us/normal/t31-alt-%d", i)
		if _, err := f.pool.Exec(context.Background(), `insert into characters (key, region, ruleset, name, user_id)
			values ($1, 'us', 'normal', $2, $3) on conflict (key) do update set user_id = excluded.user_id`,
			key, fmt.Sprintf("t31 alt %d", i), f.premium); err != nil {
			t.Fatal(err)
		}
		f.rate(t, key, 1, capsNow.Add(-time.Hour))
	}
	t.Cleanup(func() {
		if _, err := f.pool.Exec(context.Background(), `delete from characters where key like 'us/normal/t31-alt-%'`); err != nil {
			t.Error(err)
		}
	})
}

func TestNightLinesFor_ThreeCharactersAtMostAndPendingPastBudget(t *testing.T) {
	f := newT31(t)
	f.ownCharacters(t, 4)
	lines, err := f.svc.NightLinesFor(context.Background(), f.premium, t30Report, 30*time.Second)
	if err != nil || len(lines) != 1 {
		t.Fatalf("the alts have no line; the cap comes after the filter: %+v %v", lines, err)
	}
	if l := lines[0]; l.PlayerKey != f.furyKey || l.Pending || l.Line.Stem == "" || l.URL == "" {
		t.Fatalf("%+v", l)
	}
	f.blockBucket(t)
	if _, err := f.pool.Exec(context.Background(), `delete from coach_packs where report_id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	const budget = 300 * time.Millisecond
	began := time.Now()
	lines, err = f.svc.NightLinesFor(context.Background(), f.premium, t30Report, budget)
	if took := time.Since(began); took > budget+time.Second {
		t.Fatalf("one deadline for every character, took %s", took)
	}
	if err != nil || len(lines) != maxNightLines || !lines[0].Pending {
		t.Fatalf("past the budget every line is pending: %+v %v", lines, err)
	}
	if none, _ := f.svc.NightLinesFor(context.Background(), f.freeOwner, t30Report, time.Second); len(none) != 0 {
		t.Fatalf("no log_coach, no lines: %+v", none)
	}
}

func TestNightLinesFor_DeadRequestIsPendingNotAnError(t *testing.T) {
	f := newT31(t)
	f.ownCharacters(t, 1)
	if _, err := f.pool.Exec(context.Background(), `delete from coach_packs where report_id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	release := make(chan struct{})
	t.Cleanup(func() { close(release); time.Sleep(50 * time.Millisecond) })
	f.svc.Summaries = cancellingBucket{cancel: cancel, release: release}
	lines, err := f.svc.NightLinesFor(ctx, f.premium, t30Report, 30*time.Second)
	if err != nil || len(lines) != 1 || !lines[0].Pending {
		t.Fatalf("a request that ends mid-read keeps its line, pending: %+v %v", lines, err)
	}
}

// cancellingBucket ends the request on the first read, then blocks until released.
type cancellingBucket struct {
	cancel  context.CancelFunc
	release chan struct{}
}

func (b cancellingBucket) Get(context.Context, string) (io.ReadCloser, error) {
	b.cancel()
	<-b.release
	return nil, errors.New("released")
}

func TestNightLinesFor_FlaggedOrHiddenReportHasNoLines(t *testing.T) {
	f := newT31(t)
	f.ownCharacters(t, 1)
	ctx := context.Background()
	if _, err := f.pool.Exec(ctx, `update reports set flagged = 'spam' where id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	if lines, err := f.svc.NightLinesFor(ctx, f.premium, t30Report, time.Second); err != nil || len(lines) != 0 {
		t.Fatalf("flagged: %+v %v", lines, err)
	}
	if any, err := f.svc.AnyCoachHolder(ctx, t30Report); err != nil || any {
		t.Fatalf("a flagged report has no coach sentence: %v %v", any, err)
	}
	if _, err := f.pool.Exec(ctx, `update reports set flagged = null, visibility = 'private' where id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	if lines, err := f.svc.NightLinesFor(ctx, f.premium, t30Report, time.Second); err != nil || len(lines) != 0 {
		t.Fatalf("a private report the viewer may not view: %+v %v", lines, err)
	}
	if any, err := f.svc.AnyCoachHolder(ctx, t30Report); err != nil || any {
		t.Fatalf("a private report has no coach sentence: %v %v", any, err)
	}
}

func TestAnyCoachHolder(t *testing.T) {
	f := newT31(t)
	f.rate(t, f.furyKey, 1, capsNow.Add(-time.Hour))
	if any, err := f.svc.AnyCoachHolder(context.Background(), t30Report); err != nil || any {
		t.Fatalf("the free owner holds no log_coach: %v %v", any, err)
	}
	f.linkCharacter(t, f.premium, true)
	if any, err := f.svc.AnyCoachHolder(context.Background(), t30Report); err != nil || !any {
		t.Fatalf("the premium owner holds it: %v %v", any, err)
	}
}

func TestPackFor_SimRefThatCannotQueueStillYieldsAPack(t *testing.T) {
	for name, runner := range map[string]jobs.Runner{"launch error": &jobs.Fake{Err: errors.New("no cloud")}, "no runner": nil} {
		f := newT30(t)
		f.svc = f.service(t, coachDirWith(t), true)
		f.svc.Refs.Sims.Jobs = runner
		p, err := f.svc.PackFor(context.Background(), f.rep(t), t30Fight, f.furyKey)
		if err != nil || p.Version == "" {
			t.Fatalf("%s: %+v %v", name, p.Status, err)
		}
	}
}
