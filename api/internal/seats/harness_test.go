package seats

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

const testBaseURL = "https://foreversixty.gg"

var testNow = time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)

type fakeCanceller struct {
	ids     []string
	resumed []string
	err     error
	// during runs inside the cancel call, before it answers.
	during func(ctx context.Context)
}

func (f *fakeCanceller) ResumeSubscription(_ context.Context, id string) error {
	f.resumed = append(f.resumed, id)
	return nil
}

func (f *fakeCanceller) CancelSubscriptionAtPeriodEnd(ctx context.Context, id string) error {
	if f.during != nil {
		f.during(ctx)
	}
	if f.err != nil {
		return f.err
	}
	f.ids = append(f.ids, id)
	return nil
}

type fakeFrozen struct{ frozen bool }

func (f fakeFrozen) FrozenClaimant(context.Context, int64, int64) (bool, error) { return f.frozen, nil }

type harness struct {
	pool   *pgxpool.Pool
	svc    *Service
	mailer *mail.Fake
	cancel *fakeCanceller
	mux    *http.ServeMux
	seq    atomic.Int64
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
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
		`truncate users, guilds, entitlements, characters cascade`); err != nil {
		t.Fatal(err)
	}
	h := &harness{pool: pool, mailer: &mail.Fake{}, cancel: &fakeCanceller{}}
	h.svc = &Service{
		Store:  &Store{Pool: pool, Entitlements: &entitlements.Store{Pool: pool}},
		Mailer: h.mailer, Canceller: h.cancel, PublicBaseURL: testBaseURL,
		Now: func() time.Time { return testNow },
	}
	h.mux = http.NewServeMux()
	Mount(h.mux, h.svc, 1)
	return h
}

func (h *harness) exec(t *testing.T, sql string, args ...any) {
	t.Helper()
	if _, err := h.pool.Exec(context.Background(), sql, args...); err != nil {
		t.Fatal(err)
	}
}

func (h *harness) user(t *testing.T, name string) int64 {
	t.Helper()
	var id int64
	if err := h.pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`, name+"@example.com").Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// member makes a verified guild member with one character named name.
func (h *harness) member(t *testing.T, guildID int64, name, rank string) int64 {
	t.Helper()
	uid := h.user(t, name)
	key := fmt.Sprintf("us/hardcore/%s-%d", name, h.seq.Add(1))
	h.exec(t, `insert into characters (key, region, ruleset, name, class, user_id) values ($1,'us','hardcore',$2,'mage',$3)`,
		key, name, uid)
	h.exec(t, `insert into guild_characters (guild_id, character_key, user_id, rank, verified_at) values ($1,$2,$3,$4,now())`,
		guildID, key, uid, rank)
	h.exec(t, `insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at) values ($1,$2,$3,now(),now())`,
		guildID, uid, rank)
	return uid
}

// guild makes a claimed guild on tier with an active Stripe plan billed to its leader.
func (h *harness) guild(t *testing.T, tier string) (guildID, leader int64) {
	t.Helper()
	if err := h.pool.QueryRow(context.Background(),
		`insert into guilds (region, ruleset, name) values ('us','hardcore','Olympus') returning id`).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	leader = h.member(t, guildID, "Leader", "leader")
	h.exec(t, `update guilds set claimed_by = $2, claimed_at = now() where id = $1`, guildID, leader)
	h.exec(t, `insert into entitlements (guild_id, plan, source, status, tier, stripe_subscription_id, billing_user_id, current_period_end)
		values ($1,'guild','stripe','active',$2,'sub_guild',$3, now() + interval '30 days')`, guildID, tier, leader)
	return guildID, leader
}

// payingPremium gives userID an active personal Stripe Premium ending at end.
func (h *harness) payingPremium(t *testing.T, userID int64, end time.Time) {
	t.Helper()
	h.exec(t, `insert into entitlements (user_id, plan, source, status, stripe_subscription_id, billing_user_id, current_period_end, price_lookup_key, unit_amount_cents)
		values ($1,'premium','stripe','active',$2,$1,$3,$4,500)`,
		userID, fmt.Sprintf("sub_user_%d", userID), end, catalog.LookupKeyPremiumFounderMonthly)
}

// payingPremiumCanceling is payingPremium for a subscription already set to end.
func (h *harness) payingPremiumCanceling(t *testing.T, userID int64, end time.Time) {
	t.Helper()
	h.payingPremium(t, userID, end)
	h.exec(t, `update entitlements set cancel_at_period_end = true where user_id = $1`, userID)
}

func (h *harness) request(method, path string, userID int64, body any) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&buf).Encode(body)
	}
	r := httptest.NewRequest(method, path, &buf)
	r.Header.Set("Content-Type", "application/json")
	r = r.WithContext(auth.WithActor(r.Context(), auth.Actor{UserID: userID, Role: "member", Method: "session"}))
	w := httptest.NewRecorder()
	h.mux.ServeHTTP(w, r)
	return w
}

func envelope(t *testing.T, w *httptest.ResponseRecorder) httpx.Envelope {
	t.Helper()
	var e httpx.Envelope
	if err := json.NewDecoder(w.Body).Decode(&e); err != nil {
		t.Fatalf("decode: %v (%s)", err, w.Body.String())
	}
	return e
}
