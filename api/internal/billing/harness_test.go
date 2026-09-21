// api/internal/billing/harness_test.go
package billing

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const testWebhookSecret = "whsec_test_fixture_secret_do_not_use_in_prod"
const testPublicBaseURL = "https://foreversixty.gg"

// harness wires a real database (entitlements.Store, guild_members reads)
// behind a fake Stripe gateway, and a real auth.Store for GuildRank/User —
// the same "real DB, faked external API" split every other package's own
// harness_test.go uses.
type harness struct {
	pool    *pgxpool.Pool
	gateway *FakeGateway
	svc     *Service
	mux     *http.ServeMux
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	pool := testPool(t)
	gw := &FakeGateway{Prices: map[string]string{
		LookupKeyPremiumMonthly: "price_premium_monthly",
		LookupKeyPremiumYearly:  "price_premium_yearly",
		LookupKeyGuildMonthly:   "price_guild_monthly",
		LookupKeyGuildYearly:    "price_guild_yearly",
	}}
	authStore := &auth.Store{Pool: pool}
	svc := &Service{
		Store: &Store{Pool: pool}, Entitlements: &entitlements.Store{Pool: pool},
		Accounts: authStore, Users: authStore, Guilds: nil, Gateway: gw,
		PublicBaseURL: testPublicBaseURL, WebhookSecret: testWebhookSecret, Log: slog.Default(),
	}
	mux := http.NewServeMux()
	Mount(mux, svc, 1)
	return &harness{pool: pool, gateway: gw, svc: svc, mux: mux}
}

func (h *harness) seedUser(t *testing.T, email string) int64 {
	t.Helper()
	var id int64
	if err := h.pool.QueryRow(t.Context(),
		`insert into users (email) values ($1) returning id`, email).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func (h *harness) seedGuild(t *testing.T, name string, claimed bool) int64 {
	t.Helper()
	var claimedBy any
	if claimed {
		claimedBy = h.seedUser(t, name+"-claimant@example.com")
	}
	var id int64
	if err := h.pool.QueryRow(t.Context(),
		`insert into guilds (region, ruleset, name, claimed_by, claimed_at)
		 values ('us', 'hardcore', $1, $2, case when $2::bigint is not null then now() end) returning id`,
		name, claimedBy).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func (h *harness) seedGuildMember(t *testing.T, guildID, userID int64, rank string) {
	t.Helper()
	if _, err := h.pool.Exec(t.Context(),
		`insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at)
		 values ($1, $2, $3, now(), now())`, guildID, userID, rank); err != nil {
		t.Fatal(err)
	}
}

// sessionRequest builds an authenticated, CSRF-valid request the way
// auth.Authenticator.Middleware would resolve one, without a real
// session cookie: it puts the actor on the context directly, exactly as
// every other package's own harness_test.go already does for a service
// mounted below the middleware layer.
func (h *harness) sessionRequest(t *testing.T, method, path string, userID int64, body any) *http.Request {
	t.Helper()
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatal(err)
		}
	}
	r := httptest.NewRequest(method, path, &buf)
	r.Header.Set("Content-Type", "application/json")
	return r.WithContext(auth.WithActor(r.Context(), auth.Actor{UserID: userID, Role: "member", Method: "session"}))
}

func (h *harness) do(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	h.mux.ServeHTTP(w, r)
	return w
}

func decodeEnvelope(t *testing.T, w *httptest.ResponseRecorder) httpx.Envelope {
	t.Helper()
	var e httpx.Envelope
	if err := json.NewDecoder(w.Body).Decode(&e); err != nil {
		t.Fatalf("decode envelope: %v (body: %s)", err, w.Body.String())
	}
	return e
}
