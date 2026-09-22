package sims

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
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
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// testEngine is the engine pin every test uses. It is read from the
// module's own constant rather than spelled as a sha, so a re-pin
// never edits a test.
var testEngine = enginever.Version

// errAnyway is a failure with nothing to say about itself, for the
// tests that only care that something went wrong.
var errAnyway = errors.New("it did not work")

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(),
		`truncate sims, sim_specs, addon_exports, fight_metrics, users cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// harness is the simulator surface mounted over the test database,
// with a local directory for R2 and a fixed actor on every request.
type harness struct {
	t            *testing.T
	store        *Store
	dir          string
	files        *store.Dir
	service      *Service
	server       *httptest.Server
	actor        auth.Actor
	owner        int64
	jobs         *jobs.Fake
	entitlements *fakeEntitlements
	planner      *fakePlanner
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	pool := testPool(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	if os.Getenv("TEST_LOG") != "" {
		quiet = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
	accounts := &auth.Store{Pool: pool}
	owner, err := accounts.UpsertEmailUser(context.Background(), "simmer@example.com")
	if err != nil {
		t.Fatal(err)
	}
	h := &harness{t: t, store: &Store{Pool: pool}, dir: t.TempDir(), owner: owner.ID}
	h.files = store.NewDir(h.dir)
	h.actor = auth.Actor{UserID: owner.ID, Role: "user", Method: "session"}
	h.jobs, h.entitlements = &jobs.Fake{}, &fakeEntitlements{}
	// A plan that fits: one combination, well inside both bounds. A test
	// that cares sets its own.
	h.planner = &fakePlanner{summary: simapi.PlanSummary{
		Kind: simapi.KindGear, Combinations: 1,
		Cap: simapi.Caps[simapi.LaneServer], IterationsTotal: 4000,
	}}
	h.service = &Service{
		Store: h.store, Accounts: h.entitlements, Jobs: h.jobs, Planner: h.planner,
		EngineVersion: testEngine, Log: quiet,
	}

	mux := http.NewServeMux()
	Mount(mux, h.service)
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h
}

// anonymous drops the harness's identity, for the routes that do not
// need one.
func (h *harness) anonymous() { h.actor = auth.Actor{} }

func (h *harness) do(method, path, contentType string, body io.Reader) *http.Response {
	h.t.Helper()
	r, err := http.NewRequest(method, h.server.URL+path, body)
	if err != nil {
		h.t.Fatal(err)
	}
	if contentType != "" {
		r.Header.Set("Content-Type", contentType)
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

func (h *harness) json(method, path, body string) *http.Response {
	return h.do(method, path, "application/json", strings.NewReader(body))
}

// data decodes an envelope's data half and fails the test when the
// envelope reports failure.
func (h *harness) data(res *http.Response, into any) {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		OK    bool            `json:"ok"`
		Data  json.RawMessage `json:"data"`
		Error *struct {
			Code   string            `json:"code"`
			Fields map[string]string `json:"fields"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	if !env.OK {
		h.t.Fatalf("envelope reports failure: %+v", env.Error)
	}
	if into != nil {
		if err := json.Unmarshal(env.Data, into); err != nil {
			h.t.Fatal(err)
		}
	}
}

// errorCode reads the code of a failing envelope.
func (h *harness) errorCode(res *http.Response) string {
	h.t.Helper()
	return h.errorBody(res).Code
}

// errorBody decodes a failing envelope's error object, code and
// message together. errorCode is the common case; a test that also
// needs the message reads it straight from here rather than through
// a second single-field accessor.
func (h *harness) errorBody(res *http.Response) struct {
	Code    string `json:"code"`
	Message string `json:"message"`
} {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	return env.Error
}

// ensureMetricsPartition makes the monthly fight_metrics partitions
// the seeded rows need.
func ensureMetricsPartition(h *harness) {
	h.t.Helper()
	if err := db.EnsureMetricsPartitions(h.t.Context(), h.store.Pool,
		time.Now().UTC().AddDate(0, -2, 0), 4); err != nil {
		h.t.Fatal(err)
	}
}

// aCharacter is a plausible simmed character: the envelope needs a
// race, a class and a level, and SimRequest.Validate rejects a request
// without them.
func aCharacter(class, race string) simapi.CharacterSpec {
	return simapi.CharacterSpec{
		Name: "Baelgrim", Race: race, Class: class, Level: simapi.SimLevel,
		Talents: "-0550000505021051-05",
		Gear: []simapi.GearSlot{
			{Slot: "main_hand", ItemID: 17182, Enchant: 2564},
			{Slot: "trinket1", ItemID: 19406},
		},
		Buffs:    []string{"blessing_of_kings"},
		Consumes: []string{"flask_of_the_titans"},
	}
}

// browserResult is a plausible finished browser run, for the tests
// that need one.
func browserResult(spec string, mean float64) simapi.SimResult {
	class, race := "warrior", "orc"
	if strings.HasPrefix(spec, "mage") {
		class, race = "mage", "gnome"
	}
	return simapi.SimResult{
		EngineVersion: testEngine, Lane: simapi.LaneBrowser,
		Request: simapi.SimRequest{
			EngineVersion: testEngine, Spec: spec, Iterations: defaultIterations,
			Source:    simapi.CharacterSource{Kind: simapi.SourceAddon, Ref: "us/normal/baelgrim"},
			Character: aCharacter(class, race),
			Encounter: withEncounterDefaults(simapi.EncounterSpec{}),
		},
		DPS:           simapi.Estimate{Mean: mean, StdDev: 80, Error: 1.5, Min: mean - 200, Max: mean + 200},
		IterationsRun: defaultIterations, DurationMS: 2400,
	}
}

// fakeEntitlements answers the server-sims Can() question without an
// entitlements table.
type fakeEntitlements struct {
	allowed bool
	err     error
}

func (f fakeEntitlements) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	if !f.allowed {
		return false, entitlements.ReasonNoPlan, f.err
	}
	return true, entitlements.ReasonEntitled, f.err
}

// fakePlanner answers the plan-only count without a binary. The real one
// invokes `forever-sim -plan` (contract 10.2); nothing about the
// handler's decision needs a subprocess to be tested.
//
// It can model either shape a real Planner answers a cap breach with:
// err set to a simapi.ErrCapExceeded reproduces Native's and a
// CapBreach-driven Fixture's (the request is refused outright, the
// call itself fails), while summary set to an over-cap PlanSummary
// with err left nil reproduces a PlanCombinations-driven Fixture's
// (the call succeeds with a summary the caller must still compare
// against its own Cap).
type fakePlanner struct {
	summary simapi.PlanSummary
	err     error
	asked   []simapi.SimRequest
}

func (f *fakePlanner) Plan(_ context.Context, req simapi.SimRequest) (simapi.PlanSummary, error) {
	f.asked = append(f.asked, req)
	if f.err != nil {
		return simapi.PlanSummary{}, f.err
	}
	return f.summary, nil
}
