// api/internal/guilds/me_queries_test.go
package guilds

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// meStatements is how many statements GET /v1/me runs for a verified officer of n claimed
// guilds, each on a Stripe plan paid by another account, so every per-guild read runs.
func meStatements(t *testing.T, n int) int64 {
	t.Helper()
	h := newHTTPHarness(t)
	viewer := seedUser(t, h.pool, "me-counted-viewer@example.com")
	for i := 0; i < n; i++ {
		gid := seedGuild(t, h.pool, fmt.Sprintf("Me Counted %d", i))
		payer := seedRanked(t, h, gid, fmt.Sprintf("Mepayer%d", i), "leader", true)
		seedCharacter(t, h.pool, gid, viewer, fmt.Sprintf("us/hardcore/Meviewer%d", i), "officer", true)
		syncMembership(t, h.pool, gid, viewer)
		setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, gid, payer)
		seedStripePlan(t, h, gid, "guild", "active", daysFromNow(20), false, &payer)
	}
	store, counter := countedStore(t)
	svc := &auth.Service{
		Store: &auth.Store{Pool: store.Pool}, Entitlements: &entitlements.Store{Pool: store.Pool},
		GuildCapabilities: store, Standing: store,
		Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	mux := http.NewServeMux()
	auth.Mount(mux, svc, 0)
	actor := auth.Actor{UserID: viewer, Role: "user", Method: "session"}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), actor)))
	}))
	t.Cleanup(srv.Close)
	res, err := http.Get(srv.URL + "/v1/me")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("GET /v1/me = %d, want 200", res.StatusCode)
	}
	return counter.n.Load()
}

// TestMe_QueryCountFixed: the whole /v1/me read does not grow with the guild count - the
// capabilities, the officer plans and the payers' names are each one query.
func TestMe_QueryCountFixed(t *testing.T) {
	const want = 24
	one, three := meStatements(t, 1), meStatements(t, 3)
	if one != three || one != want {
		t.Fatalf("GET /v1/me ran %d statements for 1 guild and %d for 3, want %d for both", one, three, want)
	}
}
