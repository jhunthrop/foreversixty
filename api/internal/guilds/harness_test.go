// api/internal/guilds/harness_test.go
package guilds

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

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
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
	if _, err := pool.Exec(context.Background(), `truncate users, guilds, reports, fight_metrics cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// seedUser inserts a bare account and returns its id.
func seedUser(t *testing.T, pool *pgxpool.Pool, email string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`, email).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedGuild inserts a guild and returns its id.
func seedGuild(t *testing.T, pool *pgxpool.Pool, name string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', $1) returning id`, name).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// seedCharacter inserts a guild_characters row directly.
func seedCharacter(t *testing.T, pool *pgxpool.Pool, guildID, userID int64, key, rank string, verified bool) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		insert into guild_characters (guild_id, character_key, user_id, rank, verified_at)
		values ($1, $2, $3, $4, case when $5 then now() else null end)`,
		guildID, key, userID, rank, verified); err != nil {
		t.Fatal(err)
	}
}

// httpHarness is the HTTP-level guilds test harness: a real auth.Store
// for Accounts (the same one production wires in), a Service mounted on
// a test server, and a fixed actor every request carries.
type httpHarness struct {
	t      *testing.T
	pool   *pgxpool.Pool
	store  *Store
	actor  auth.Actor
	server *httptest.Server
}

func newHTTPHarness(t *testing.T) *httpHarness {
	t.Helper()
	pool := testPool(t)
	accounts := &auth.Store{Pool: pool}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	store := &Store{Pool: pool}
	svc := &Service{Store: store, Accounts: accounts, Log: quiet}
	mux := http.NewServeMux()
	Mount(mux, svc, 0)
	h := &httpHarness{t: t, pool: pool, store: store}
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h
}

func (h *httpHarness) do(method, path, body string) *http.Response {
	h.t.Helper()
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, h.server.URL+path, reader)
	if err != nil {
		h.t.Fatal(err)
	}
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

// data decodes an envelope's data half, failing the test on a failure envelope.
func (h *httpHarness) data(res *http.Response, into any) {
	h.t.Helper()
	defer res.Body.Close()
	var env struct {
		OK    bool            `json:"ok"`
		Data  json.RawMessage `json:"data"`
		Error *struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		h.t.Fatal(err)
	}
	if !env.OK {
		h.t.Fatalf("envelope failure: %+v", env.Error)
	}
	if into != nil {
		if err := json.Unmarshal(env.Data, into); err != nil {
			h.t.Fatal(err)
		}
	}
}
