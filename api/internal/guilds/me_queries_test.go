// api/internal/guilds/me_queries_test.go
package guilds

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

// simsView adapts *sims.Store to auth.SimsView as cmd/api's simsAccountView does.
type simsView struct{ store *sims.Store }

func (v simsView) FinishedUnseen(ctx context.Context, userID int64, limit int) ([]auth.FinishedSim, error) {
	runs, err := v.store.FinishedUnseen(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	out := make([]auth.FinishedSim, 0, len(runs))
	for _, r := range runs {
		out = append(out, auth.FinishedSim{SimID: r.SimID, Spec: r.Spec, Title: r.Title, DPS: r.DPS,
			FinishedAt: r.FinishedAt.UTC().Format(time.RFC3339)})
	}
	return out, nil
}

func (v simsView) RunsThisMonth(ctx context.Context, userID int64) (int, error) {
	return v.store.RunsThisMonth(ctx, userID)
}

// seedOwnedCharacter inserts the characters row key names, owned by userID.
func seedOwnedCharacter(t *testing.T, h *httpHarness, key string, userID int64) {
	t.Helper()
	if _, err := h.pool.Exec(context.Background(),
		`insert into characters (key, region, ruleset, name, class, user_id, realm_name, level, faction, source, imported_at, refreshed_at)
		 values ($1, 'us', 'hardcore', $1, 'warrior', $2, 'Whitemane', 60, 'alliance', 'bnet', now(), now())`,
		key, userID); err != nil {
		t.Fatal(err)
	}
}

// meStatements is how many statements GET /v1/me runs for a verified officer of n claimed
// guilds, each on a Stripe plan paid by another account, so every per-guild read runs.
func meStatements(t *testing.T, n int) int64 {
	t.Helper()
	h := newHTTPHarness(t)
	viewer := seedUser(t, h.pool, "me-counted-viewer@example.com")
	for i := 0; i < n; i++ {
		gid := seedGuild(t, h.pool, fmt.Sprintf("Me Counted %d", i))
		payer := seedRanked(t, h, gid, fmt.Sprintf("Mepayer%d", i), "leader", true)
		key := fmt.Sprintf("us/hardcore/Meviewer%d", i)
		seedCharacter(t, h.pool, gid, viewer, key, "officer", true)
		seedOwnedCharacter(t, h, key, viewer)
		seedDoneRaid(t, h, gid, viewer, time.Now().Add(-time.Hour), seedAtt{key: key, signed: "in", outcome: "present"})
		syncMembership(t, h.pool, gid, viewer)
		setGuild(t, h, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, gid, payer)
		seedStripePlan(t, h, gid, "guild", "active", daysFromNow(20), false, &payer)
	}
	if err := (&entitlements.Store{Pool: h.pool}).Grant(context.Background(), entitlements.Subject{UserID: &viewer},
		entitlements.PlanPremium, nil, nil, "owner", "t", nil); err != nil {
		t.Fatal(err) // server_sims on, so the finished-runs read runs too
	}
	store, counter := countedStore(t)
	svc := &auth.Service{
		Store: &auth.Store{Pool: store.Pool}, Entitlements: &entitlements.Store{Pool: store.Pool},
		GuildCapabilities: store, GuildAttendance: store, Standing: store,
		Sims: simsView{store: &sims.Store{Pool: store.Pool}},
		Log:  slog.New(slog.NewTextHandler(io.Discard, nil)),
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
	const want = 28
	one, three := meStatements(t, 1), meStatements(t, 3)
	if one != three || one != want {
		t.Fatalf("GET /v1/me ran %d statements for 1 guild and %d for 3, want %d for both", one, three, want)
	}
}
