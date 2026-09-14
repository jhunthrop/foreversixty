package reports

import (
	"context"
	"encoding/json"
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
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

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
		`truncate users, reports, fight_metrics, percentile_digests, moderation cascade`); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureMetricsPartitions(context.Background(), pool, engine.FixtureBase, 1); err != nil {
		t.Fatal(err)
	}
	return pool
}

// fakeSigner signs by naming the key, which is all a test needs to see.
type fakeSigner struct{}

func (fakeSigner) PresignGet(_ context.Context, key string, ttl time.Duration) (string, error) {
	return "https://r2.example/signed/" + key + "?ttl=" + ttl.String(), nil
}

// harness is the report surface mounted over the test database, with a
// local directory for R2 and a fixed actor on every request.
type harness struct {
	t      *testing.T
	store  *Store
	dir    string
	files  *store.Dir
	server *httptest.Server
	actor  auth.Actor
	owner  int64
	// reportID is the report a store test seeded fights into.
	reportID string
	service  *Service
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	pool := testPool(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	if os.Getenv("TEST_LOG") != "" {
		quiet = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
	accounts := &auth.Store{Pool: pool}
	owner, err := accounts.UpsertEmailUser(context.Background(), "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	h := &harness{
		t: t, store: &Store{Pool: pool}, dir: t.TempDir(),
		owner: owner.ID,
	}
	h.files = store.NewDir(h.dir)
	h.actor = auth.Actor{UserID: owner.ID, Role: "user", Method: "device", DeviceID: "device-1"}

	h.service = &Service{
		Store: h.store, Accounts: accounts, Signer: fakeSigner{},
		PublicBaseURL: "https://foreversixty.gg", APIBaseURL: "https://api.foreversixty.gg", Log: quiet,
	}
	mux := http.NewServeMux()
	Mount(mux, h.service)
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h
}

// do sends a request as the harness's current actor.
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

// data decodes an envelope's data half.
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

// errorFields reads the fields map of a failing envelope.
func (h *harness) errorFields(res *http.Response) map[string]string {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		Error struct {
			Fields map[string]string `json:"fields"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	return env.Error.Fields
}

// createReport makes a report with the given visibility and returns it.
func (h *harness) createReport(visibility string) string {
	h.t.Helper()
	res := h.json(http.MethodPost, "/v1/reports",
		`{"title":"Tuesday","visibility":"`+visibility+`","zone":"Blackrock Spire",
		  "logging_character":{"region":"us","ruleset":"hardcore","name":"Baelgrim"}}`)
	var out struct {
		ID string `json:"id"`
	}
	h.data(res, &out)
	return out.ID
}

// fileExists reports whether the local object store holds a key.
func (h *harness) fileExists(key string) bool {
	h.t.Helper()
	_, err := os.Stat(h.dir + "/" + key)
	return err == nil
}

// readFile reads an object out of the local store.
func (h *harness) readFile(key string) []byte {
	h.t.Helper()
	b, err := os.ReadFile(h.dir + "/" + key)
	if err != nil {
		h.t.Fatal(err)
	}
	return b
}

// asSession switches the harness to a browser session, which is what
// the upload routes require.
func (h *harness) asSession() {
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
}

// seedFight writes one fight row directly, for the store tests.
func (h *harness) seedFight(index int, encounterID int64, start, end int64) {
	h.t.Helper()
	rec := FightRecord{
		ReportID: h.reportID, Index: index, Name: "Warden Kelthas", Kill: true,
		DurationMS: 30000, StartMS: time.Now().UnixMilli(), Verified: true,
		Players: []string{"Player-1", "Player-2"}, Deaths: 1, NPCKills: 2,
		RawStart: &start, RawEnd: &end,
	}
	if encounterID != 0 {
		d, s := int64(8), int64(5)
		rec.EncounterID, rec.Difficulty, rec.Size = &encounterID, &d, &s
	}
	if _, err := h.store.UpsertFight(context.Background(), rec); err != nil {
		h.t.Fatal(err)
	}
}
