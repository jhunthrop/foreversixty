package creators

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// fakeGateway is Stripe in memory: accounts, invoices, refunds and transfers.
type fakeGateway struct {
	mu         sync.Mutex
	accounts   map[string]AccountFacts
	invoices   map[string]InvoiceMoney
	refunds    map[string][]Refund
	transfers  map[string]TransferParams
	creates    int
	failNext   error
	connectErr error
	// available is the platform balance in cents; -1 is unlimited.
	available int64
	nextID    int
}

func newFakeGateway() *fakeGateway {
	return &fakeGateway{accounts: map[string]AccountFacts{}, invoices: map[string]InvoiceMoney{},
		refunds: map[string][]Refund{}, transfers: map[string]TransferParams{}, available: -1}
}

func (g *fakeGateway) id(prefix string) string {
	g.nextID++
	return prefix + strconv.Itoa(g.nextID)
}

func (g *fakeGateway) CreateExpressAccount(context.Context, int64) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	id := g.id("acct_")
	g.accounts[id] = AccountFacts{ID: id}
	return id, nil
}

func (g *fakeGateway) AccountLink(_ context.Context, account, ret, refresh string) (string, error) {
	return "https://connect.stripe.test/setup/" + account + "?return=" + ret + "&refresh=" + refresh, nil
}

func (g *fakeGateway) LoginLink(_ context.Context, account string) (string, error) {
	return "https://connect.stripe.test/express/" + account, nil
}

func (g *fakeGateway) CheckConnect(context.Context) error { return g.connectErr }

func (g *fakeGateway) Account(_ context.Context, id string) (AccountFacts, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.accounts[id], nil
}

func (g *fakeGateway) InvoiceMoney(_ context.Context, id string) (InvoiceMoney, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.invoices[id], nil
}

func (g *fakeGateway) Refunds(_ context.Context, charge string) ([]Refund, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.refunds[charge], nil
}

func (g *fakeGateway) AvailableUSD(context.Context) (int64, error) {
	if g.available < 0 {
		return 1 << 40, nil
	}
	return g.available, nil
}

func (g *fakeGateway) FindTransfer(_ context.Context, group string) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.transfers[group]; ok {
		return "tr_" + group, nil
	}
	return "", nil
}

func (g *fakeGateway) CreateTransfer(_ context.Context, p TransferParams) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.available >= 0 && p.AmountCents > g.available {
		return "", ErrInsufficientBalance
	}
	if g.failNext != nil {
		err := g.failNext
		g.failNext = nil
		return "", err
	}
	g.creates++
	g.transfers[p.Group] = p
	return "tr_" + p.Group, nil
}

// harness is a real database behind a fake Stripe, with the clock pinned.
type harness struct {
	t    *testing.T
	pool *pgxpool.Pool
	gw   *fakeGateway
	svc  *Service
	mux  *http.ServeMux
	now  time.Time
	logs *bytes.Buffer
}

func newHarness(t *testing.T) *harness {
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
	ctx := context.Background()
	for _, sql := range []string{
		`truncate users, entitlements, creator_refunds, creator_revenue, creator_players, creator_payouts, creator_invites,
		 creator_applications, go_link_clicks cascade`,
		`delete from creators`,
		`delete from go_links where label like 'Creator: %'`,
	} {
		if _, err := pool.Exec(ctx, sql); err != nil {
			t.Fatal(err)
		}
	}
	h := &harness{t: t, pool: pool, gw: newFakeGateway(), now: at("2026-12-09T12:00:00Z"), logs: &bytes.Buffer{}}
	h.svc = &Service{Store: &Store{Pool: pool}, Entitlements: &entitlements.Store{Pool: pool}, Gateway: h.gw,
		PublicBaseURL: "https://foreversixty.gg", PayoutsEnabled: true,
		Log: slog.New(slog.NewTextHandler(h.logs, nil)), Now: func() time.Time { return h.now }}
	h.mux = http.NewServeMux()
	Mount(h.mux, h.svc, 1)
	return h
}

func (h *harness) exec(sql string, args ...any) {
	h.t.Helper()
	if _, err := h.pool.Exec(context.Background(), sql, args...); err != nil {
		h.t.Fatalf("%s: %v", sql, err)
	}
}

// user makes an account signed up at created, with a Battle.net login.
func (h *harness) user(created time.Time) int64 {
	h.t.Helper()
	var id int64
	if err := h.pool.QueryRow(context.Background(), `insert into users (bnet_sub, battletag, created_at)
		values (gen_random_uuid()::text, 'player#' || floor(random()*100000)::text, $1) returning id`, created).Scan(&id); err != nil {
		h.t.Fatal(err)
	}
	return id
}

// character gives the account a Forever character imported from Battle.net.
func (h *harness) character(userID int64, key string) {
	h.exec(`insert into characters (key, region, ruleset, name, user_id, source) values ($1, 'us', 'normal', $1, $2, 'bnet')`, key, userID)
}

// creator adds a creator through the service.
func (h *harness) creator(name, start string) Creator {
	h.t.Helper()
	cap := int64(DefaultCapCents)
	c, err := h.svc.Add(context.Background(), AddInput{Input: Input{Name: name, ChannelURL: "https://twitch.tv/" + name,
		Slug: name, Code: NormalizeCode(name) + "1", StartsOn: start, CapCents: &cap}})
	if err != nil {
		h.t.Fatal(err)
	}
	return c
}

// ready gives the creator a payable Express account.
func (h *harness) ready(c Creator) Creator {
	h.t.Helper()
	account := "acct_ready_" + strconv.FormatInt(c.ID, 10)
	h.gw.accounts[account] = AccountFacts{ID: account, DetailsSubmitted: true, PayoutsEnabled: true, TransfersActive: true}
	if err := h.svc.Store.SetStripeAccount(context.Background(), c.ID, account); err != nil {
		h.t.Fatal(err)
	}
	if err := h.svc.AccountUpdated(context.Background(), account); err != nil {
		h.t.Fatal(err)
	}
	got, err := h.svc.Store.Creator(context.Background(), c.ID)
	if err != nil || got.StripeStatus != StatusReady {
		h.t.Fatalf("ready: %+v %v", got, err)
	}
	return got
}

// returning makes a link player of c who comes back: a character and a second day.
func (h *harness) returning(c Creator, signup time.Time, key string) int64 {
	h.t.Helper()
	uid := h.user(signup)
	if ok, err := h.svc.Store.AttributeLink(context.Background(), uid, c.Slug, signup.Add(-time.Hour)); err != nil || !ok {
		h.t.Fatalf("attribute: %v %v", ok, err)
	}
	h.character(uid, key)
	if ok, err := h.svc.Store.MarkVisit(context.Background(), uid, signup.Add(26*time.Hour)); err != nil || !ok {
		h.t.Fatalf("visit: %v %v", ok, err)
	}
	return uid
}

// pays records a $20.00 net payment by the player in the month of at: a $10.00 share, so a month
// with it clears the $10 minimum payout.
func (h *harness) pays(uid int64, invoice string, paid time.Time) {
	h.t.Helper()
	h.gw.invoices[invoice] = InvoiceMoney{ChargeID: "ch_" + invoice, PaidAt: paid, AmountCents: 2100, FeeCents: 100, NetCents: 2000}
	if err := h.svc.InvoicePaid(context.Background(), uid, invoice); err != nil {
		h.t.Fatal(err)
	}
}

func (h *harness) statement(c Creator, month string) Statement {
	h.t.Helper()
	m, _ := ParseMonth(month)
	c, _ = h.svc.Store.Creator(context.Background(), c.ID)
	st, err := h.svc.Store.Statement(context.Background(), c, m, h.now)
	if err != nil {
		h.t.Fatal(err)
	}
	return st
}

func (h *harness) request(method, path string, actor auth.Actor, body any) *httptest.ResponseRecorder {
	h.t.Helper()
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			h.t.Fatal(err)
		}
	}
	r := httptest.NewRequest(method, path, &buf)
	r = r.WithContext(auth.WithActor(r.Context(), actor))
	w := httptest.NewRecorder()
	h.mux.ServeHTTP(w, r)
	return w
}

var (
	anonymous = auth.Actor{}
	owner     = auth.Actor{UserID: 1, Role: "admin", Method: "session"}
)

func member(id int64) auth.Actor { return auth.Actor{UserID: id, Role: "user", Method: "session"} }
