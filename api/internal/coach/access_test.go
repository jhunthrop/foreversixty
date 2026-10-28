package coach

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	t30Report = "t30-coach-report"
	t30Fight  = 1
	t30GuildA = "t30 Guild A"
	t30GuildB = "t30 Guild B"
)

// fakeOwners answers ownership from a map of user -> owned keys.
type fakeOwners map[int64]map[string]bool

func (f fakeOwners) OwnsCharacter(_ context.Context, userID int64, key string) (bool, error) {
	return f[userID][key], nil
}

// fakeEnt grants log_coach to coach users and officer_views per guild.
type fakeEnt struct {
	coach    map[int64]bool
	officers map[int64]map[int64]bool // guild -> user
}

func (f fakeEnt) Can(_ context.Context, userID int64, feat entitlements.Feature) (bool, entitlements.Reason, error) {
	return feat == entitlements.FeatureLogCoach && f.coach[userID], "", nil
}

func (f fakeEnt) CanForGuild(_ context.Context, guildID, userID int64, feat entitlements.Feature) (bool, entitlements.Reason, error) {
	return feat == entitlements.FeatureOfficerViews && f.officers[guildID][userID], "", nil
}

// fakeRanker makes every listed user a member of every guild.
type fakeRanker map[int64]bool

func (f fakeRanker) GuildRank(_ context.Context, _, userID int64) (string, bool, error) {
	return "member", f[userID], nil
}

// fakeBucket serves summaries by key and counts the reads.
type fakeBucket struct {
	mu    sync.Mutex
	files map[string][]byte
	gets  int
}

func (b *fakeBucket) Get(_ context.Context, key string) (io.ReadCloser, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.gets++
	data, ok := b.files[key]
	if !ok {
		return nil, os.ErrNotExist
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}

func (b *fakeBucket) reads() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.gets
}

// t30 is the fight route's world: one public report with the Fury warrior's pull, and
// the accounts the tests view it as.
type t30 struct {
	pool                               *pgxpool.Pool
	svc                                *Service
	jobs                               *jobs.Fake
	llm                                *llm.Fake
	bucket                             *fakeBucket
	owners                             fakeOwners
	ent                                fakeEnt
	ranker                             fakeRanker
	furyKey                            string
	freeOwner, premium, stranger, mod  int64
	officerA, officerB, guildA, guildB int64
	mux                                *http.ServeMux
}

func newT30(t *testing.T) *t30 {
	t.Helper()
	pool := capsPool(t)
	f := &t30{pool: pool, jobs: &jobs.Fake{}, llm: &llm.Fake{}, owners: fakeOwners{}, ranker: fakeRanker{},
		ent: fakeEnt{coach: map[int64]bool{}, officers: map[int64]map[int64]bool{}}}
	f.seedDB(t)
	sum := furyThreeFindings(t).Fight
	raw, err := json.Marshal(sum)
	if err != nil {
		t.Fatal(err)
	}
	f.bucket = &fakeBucket{files: map[string][]byte{store.Keys{ReportID: t30Report}.FightSummary(t30Fight): raw}}
	f.furyKey = character.KeyFromUnit("us", character.RulesetNormal, furyRowOf(t, sum.Roster).Name)
	f.seedUsers(t)
	f.svc = f.service(t, coachDirWith(t), false)
	f.mux = http.NewServeMux()
	Mount(f.mux, f.svc)
	return f
}

func (f *t30) seedDB(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	for _, q := range []string{
		`delete from coach_profiles`,
		`delete from reports where id = '` + t30Report + `'`,
		`delete from guilds where name like 't30 %'`,
		`delete from characters where name = 't30 fury'`,
	} {
		if _, err := f.pool.Exec(ctx, q); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() { f.pool.Exec(context.Background(), `delete from reports where id = $1`, t30Report) })
	f.guildA, f.guildB = t30Guild(t, f.pool, t30GuildA), t30Guild(t, f.pool, t30GuildB)
	if _, err := f.pool.Exec(ctx, `insert into reports (id, visibility, status, guild_id) values ($1, 'public', 'complete', $2)`,
		t30Report, f.guildA); err != nil {
		t.Fatal(err)
	}
	if _, err := f.pool.Exec(ctx, `insert into fights (report_id, fight_index, encounter_id, name, difficulty, size, kill, duration_ms, start_ms)
		values ($1, $2, $3, $4, $5, 40, true, $6, $7)`, t30Report, t30Fight, onyxiaEncounterID, onyxiaName, raidDifficulty40,
		reviewLongMS, capsNow.UnixMilli()); err != nil { // the pull is on the newest build
		t.Fatal(err)
	}
}

func t30Guild(t *testing.T, pool *pgxpool.Pool, name string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'normal', $1) returning id`, name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedUsers: the free owner owns the Fury warrior's character, a verified member of guild
// A; the premium user owns it too (two accounts are never both owners in production, but
// each test views as one of them).
func (f *t30) seedUsers(t *testing.T) {
	t.Helper()
	f.freeOwner = capsUser(t, f.pool, "t30-free@example.com")
	f.premium = capsUser(t, f.pool, "t30-premium@example.com")
	f.stranger = capsUser(t, f.pool, "t30-stranger@example.com")
	f.mod = capsUser(t, f.pool, "t30-mod@example.com")
	f.officerA = capsUser(t, f.pool, "t30-officer-a@example.com")
	f.officerB = capsUser(t, f.pool, "t30-officer-b@example.com")
	f.owners[f.freeOwner] = map[string]bool{f.furyKey: true}
	f.owners[f.premium] = map[string]bool{f.furyKey: true}
	f.ent.coach[f.premium] = true
	f.ent.officers[f.guildA] = map[int64]bool{f.officerA: true}
	f.ent.officers[f.guildB] = map[int64]bool{f.officerB: true}
	f.linkCharacter(t, f.freeOwner, true)
}

// linkCharacter makes userID the character's account, a guild-A member verified or not.
func (f *t30) linkCharacter(t *testing.T, userID int64, verified bool) {
	t.Helper()
	ctx := context.Background()
	if _, err := f.pool.Exec(ctx, `insert into characters (key, region, ruleset, name, user_id) values ($1, 'us', 'normal', 't30 fury', $2)
		on conflict (key) do update set user_id = excluded.user_id`, f.furyKey, userID); err != nil {
		t.Fatal(err)
	}
	verifiedAt := any(nil)
	if verified {
		verifiedAt = capsNow
	}
	if _, err := f.pool.Exec(ctx, `insert into guild_members (guild_id, user_id, verified_at) values ($1, $2, $3)
		on conflict (guild_id, user_id) do update set verified_at = excluded.verified_at`, f.guildA, userID, verifiedAt); err != nil {
		t.Fatal(err)
	}
}

func (f *t30) service(t *testing.T, coachDir string, validated bool) *Service {
	t.Helper()
	store := &Store{DB: f.pool}
	caps := newCaps(f.pool)
	now := func() time.Time { return capsNow }
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	return &Service{
		Store: store, Reports: &reports.Store{Pool: f.pool}, Accounts: f.ranker, Owners: f.owners, Entitlements: f.ent,
		Summaries: f.bucket, Tables: &FileTables{BuildsDir: repoBuilds, CoachDir: coachDir},
		Refs: &Resolver{Pool: &Pool{DB: f.pool}, Sims: &SimRefs{DB: f.pool, Jobs: &jobs.Fake{}},
			Validated: fakeValidated{"warrior-fury": validated}, EngineVersion: "t30-engine", BuildAt: func(time.Time) string { return "t30-build" }},
		Personas:  &Personas{Dir: fixtureDir},
		Generator: &Generator{LLM: f.llm, Caps: caps, Store: store, Model: "m", Log: log, Now: now},
		Names:     buildNames(t), Caps: caps, Jobs: f.jobs, Enabled: true, Log: log, Now: now,
	}
}

func (f *t30) rep(t *testing.T) reports.Report {
	t.Helper()
	rep, err := f.svc.Reports.Get(context.Background(), t30Report)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

// get calls the route as actor and decodes the envelope's data (or error).
func (f *t30) get(t *testing.T, actor auth.Actor, query string) (int, http.Header, map[string]any) {
	t.Helper()
	url := "/v1/reports/" + t30Report + "/fights/1/coach?player=" + f.furyKey + query
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(auth.WithActor(req.Context(), actor))
	rec := httptest.NewRecorder()
	f.mux.ServeHTTP(rec, req)
	var env map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatalf("%d %s", rec.Code, rec.Body.String())
	}
	if data, ok := env["data"].(map[string]any); ok {
		return rec.Code, rec.Header(), data
	}
	errBody, _ := env["error"].(map[string]any)
	return rec.Code, rec.Header(), errBody
}

func user(id int64) auth.Actor { return auth.Actor{UserID: id, Role: "user", Method: "session"} }

func furyRowOf(t *testing.T, roster []summary.RosterRow) summary.RosterRow {
	t.Helper()
	for _, r := range roster {
		if r.GUID == furyGUID {
			return r
		}
	}
	t.Fatal("no Fury warrior in the roster")
	return summary.RosterRow{}
}

func TestAccess_OwnerFreeIsTeaserPremiumIsFull(t *testing.T) {
	f := newT30(t)
	ctx := context.Background()
	for _, c := range []struct {
		actor auth.Actor
		want  Access
	}{
		{user(f.freeOwner), AccessTeaser}, {user(f.premium), AccessFull}, {user(f.stranger), AccessNone},
		{auth.Actor{}, AccessNone}, {user(f.officerA), AccessOfficer}, {user(f.officerB), AccessNone},
		{auth.Actor{UserID: f.mod, Role: "moderator"}, AccessOfficer},
	} {
		got, account, err := f.svc.AccessFor(ctx, c.actor, f.rep(t), f.furyKey)
		if err != nil || got != c.want || (got != AccessNone && account != c.actor.UserID) {
			t.Errorf("actor %d: %q account %d err %v, want %q", c.actor.UserID, got, account, err, c.want)
		}
	}
}

func TestAccess_PrivateReportHidesEvenTheOfficerView(t *testing.T) {
	f := newT30(t)
	if _, err := f.pool.Exec(context.Background(), `update reports set visibility = 'private' where id = $1`, t30Report); err != nil {
		t.Fatal(err)
	}
	got, _, err := f.svc.AccessFor(context.Background(), user(f.officerA), f.rep(t), f.furyKey)
	if err != nil || got != AccessNone {
		t.Fatalf("%q %v", got, err)
	}
}
