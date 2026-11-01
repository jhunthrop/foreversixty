package livesurface

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// testNow is the harness clock: the board's Tuesday 8 December, 17:39 UTC.
var testNow = time.Date(2026, 12, 8, 17, 39, 27, 0, time.UTC)

type fakeAccounts struct {
	user  auth.User
	chars []auth.Character
}

func (f *fakeAccounts) User(context.Context, int64) (auth.User, error) { return f.user, nil }
func (f *fakeAccounts) Characters(context.Context, int64) ([]auth.Character, error) {
	return f.chars, nil
}

type fakeRaids struct {
	byKey   map[string]guilds.WidgetRaid
	asked   string
	raidEnd map[int64]time.Time
}

func (f *fakeRaids) WidgetRaidFor(_ context.Context, _ int64, key string, _ time.Time) (guilds.WidgetRaid, error) {
	f.asked = key
	return f.byKey[key], nil
}

func (f *fakeRaids) LiveRaidEnd(_ context.Context, _ int64, raidID int64) (time.Time, bool, error) {
	end, ok := f.raidEnd[raidID]
	return end, ok, nil
}

type fakeUpgrades map[string]*WidgetUpgrades

func (f fakeUpgrades) WidgetUpgrades(_ context.Context, key string) (*WidgetUpgrades, error) {
	return f[key], nil
}

type fakeSims struct {
	running int
	owned   map[string]bool
}

func (f *fakeSims) RunningServerRuns(context.Context, int64) (int, error) { return f.running, nil }
func (f *fakeSims) OwnsServerRun(_ context.Context, _ int64, id string) (bool, error) {
	return f.owned[id], nil
}

type fakeEntitlements struct{ premium bool }

func (f *fakeEntitlements) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	if f.premium {
		return true, "", nil
	}
	return false, entitlements.ReasonNoPlan, nil
}

// harness serves the routes with fakes for everything but the token store, which is real
// when TEST_DATABASE_URL is set.
type harness struct {
	t        *testing.T
	svc      *Service
	accounts *fakeAccounts
	raids    *fakeRaids
	sims     *fakeSims
	ent      *fakeEntitlements
	actor    auth.Actor
	server   *httptest.Server
	pool     *pgxpool.Pool
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	h := &harness{
		t: t, accounts: &fakeAccounts{}, raids: &fakeRaids{byKey: map[string]guilds.WidgetRaid{}, raidEnd: map[int64]time.Time{}},
		sims: &fakeSims{owned: map[string]bool{}}, ent: &fakeEntitlements{},
		actor: auth.Actor{UserID: 1, Role: "user", Method: "device", DeviceID: "ph", DeviceKind: auth.DeviceKindApp},
	}
	h.svc = &Service{
		Accounts: h.accounts, Raids: h.raids, Sims: h.sims, Entitlements: h.ent,
		Now: func() time.Time { return testNow },
	}
	mux := http.NewServeMux()
	Mount(mux, h.svc)
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h
}

// withStore gives the harness a real token store and a user and device row for its actor.
func (h *harness) withStore() *harness {
	h.t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		h.t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		h.t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		h.t.Fatal(err)
	}
	h.t.Cleanup(pool.Close)
	ctx := context.Background()
	if _, err := pool.Exec(ctx, `truncate users cascade`); err != nil {
		h.t.Fatal(err)
	}
	var uid int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('w@example.com') returning id`).Scan(&uid); err != nil {
		h.t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `insert into devices (id, user_id, name, platform, token_hash, kind)
		values ('ph', $1, 'iPhone', 'ios', 'h', 'app')`, uid); err != nil {
		h.t.Fatal(err)
	}
	h.pool, h.actor.UserID = pool, uid
	h.svc.Tokens = &Store{Pool: pool}
	return h
}

func (h *harness) do(method, path, body string) (*http.Response, []byte) {
	h.t.Helper()
	var rd io.Reader
	if body != "" {
		rd = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, h.server.URL+path, rd)
	if err != nil {
		h.t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.t.Fatal(err)
	}
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		h.t.Fatal(err)
	}
	return res, raw
}

func (h *harness) want(res *http.Response, raw []byte, status int) []byte {
	h.t.Helper()
	if res.StatusCode != status {
		h.t.Fatalf("%s %s = %d, want %d: %s", res.Request.Method, res.Request.URL.Path, res.StatusCode, status, raw)
	}
	return raw
}
