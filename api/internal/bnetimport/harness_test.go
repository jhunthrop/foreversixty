// api/internal/bnetimport/harness_test.go
package bnetimport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
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
	if _, err := pool.Exec(context.Background(), `truncate users, guilds, characters cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func seedUser(t *testing.T, pool *pgxpool.Pool) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`, fmt.Sprintf("u%d@example.com", time.Now().UnixNano())).
		Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// blizzardFixture is a tiny httptest-backed stand-in for the whole
// Blizzard API surface bnetimport reads, keyed by exact request path and
// query (the same shape as bnetapi's own fixtureServer, duplicated here
// rather than exported cross-package so each package's tests stay
// self-contained).
type blizzardFixture struct {
	server   *httptest.Server
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
}

func newBlizzardFixture(t *testing.T) *blizzardFixture {
	t.Helper()
	f := &blizzardFixture{handlers: map[string]func(http.ResponseWriter, *http.Request){}}
	f.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Fixtures are keyed without the locale bnetapi always appends.
		key := r.Method + " " + strings.Replace(r.URL.RequestURI(), "&locale=en_US", "", 1)
		h, ok := f.handlers[key]
		if !ok {
			t.Fatalf("blizzardFixture: unexpected request %s", key)
			return
		}
		h(w, r)
	}))
	t.Cleanup(f.server.Close)
	f.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "app-tok", "expires_in": 3600})
	return f
}

func (f *blizzardFixture) json(method, path string, status int, body any) {
	f.handlers[method+" "+path] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}
}

func (f *blizzardFixture) client() *bnetapi.Client {
	c := bnetapi.New(bnetapi.Config{
		TokenURL: f.server.URL + "/token",
		APIHost:  func(string) string { return f.server.URL },
		Game:     "classic1x", Regions: []string{"us"},
	})
	c.RetryBackoff = time.Millisecond
	return c
}

func newTestService(t *testing.T, pool *pgxpool.Pool, f *blizzardFixture) *Service {
	return &Service{Pool: pool, Client: f.client(), Regions: []string{"us"}}
}

// realms registers a region's realm index and one detail per realm, the
// two calls Realms makes, so a refresh test can resolve its rows' rulesets.
// realmTypes maps slug to Blizzard's realm type ("PVP", "NORMAL", ...).
func (f *blizzardFixture) realms(region string, realmTypes map[string]string) {
	index := []map[string]any{}
	id := 1
	for slug, realmType := range realmTypes {
		index = append(index, map[string]any{"id": id, "name": slug, "slug": slug})
		f.json(http.MethodGet, "/data/wow/realm/"+slug+"?namespace=dynamic-classic1x-"+region, http.StatusOK,
			map[string]any{"type": map[string]string{"type": realmType}, "category": "Classic Era"})
		id++
	}
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-classic1x-"+region, http.StatusOK,
		map[string]any{"realms": index})
}
