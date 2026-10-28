package coach

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

var capsNow = time.Date(2026, 10, 27, 15, 0, 0, 0, time.UTC)

func capsPool(t *testing.T) *pgxpool.Pool {
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
	if _, err := pool.Exec(context.Background(), `truncate users, llm_spend cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

func capsUser(t *testing.T, pool *pgxpool.Pool, email string) int64 {
	t.Helper()
	var id int64
	if err := pool.QueryRow(context.Background(),
		`insert into users (email) values ($1) returning id`, email).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

func newCaps(pool *pgxpool.Pool) *Caps {
	return &Caps{DB: pool, Enabled: true, SpendCapUSD: 25, PerAccount: 30, Now: func() time.Time { return capsNow }}
}

func mustAdmit(t *testing.T, c *Caps, id int64) Refusal {
	t.Helper()
	r, err := c.Admit(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func TestAdmit_SwitchOff(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "off@example.com")
	c := newCaps(pool)
	c.Enabled = false
	if r := mustAdmit(t, c, uid); r != RefusalSwitchOff {
		t.Fatalf("got %q", r)
	}
	var n int
	if err := pool.QueryRow(context.Background(), `select count(*) from coach_account_day`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("counted %d, %v", n, err)
	}
}

func TestAdmit_AccountCapAtThirty(t *testing.T) {
	pool := capsPool(t)
	a, b := capsUser(t, pool, "a@example.com"), capsUser(t, pool, "b@example.com")
	c := newCaps(pool)
	for i := 0; i < 30; i++ {
		if r := mustAdmit(t, c, a); r != RefusalNone {
			t.Fatalf("admit %d refused: %q", i+1, r)
		}
	}
	if r := mustAdmit(t, c, a); r != RefusalAccountCap {
		t.Fatalf("31st: %q", r)
	}
	if r := mustAdmit(t, c, b); r != RefusalNone {
		t.Fatalf("other account: %q", r)
	}
}

func TestAdmit_SpendCapFromRecordedCost(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "spend@example.com")
	c := newCaps(pool)
	for i := 0; i < 2; i++ {
		if err := c.RecordSpend(context.Background(), llm.Response{CostUSD: 12.5}); err != nil {
			t.Fatal(err)
		}
	}
	if r := mustAdmit(t, c, uid); r != RefusalSpendCap {
		t.Fatalf("got %q", r)
	}
	var n int
	if err := pool.QueryRow(context.Background(), `select count(*) from coach_account_day`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("a refused spend counted %d, %v", n, err)
	}
}

func TestAdmit_JustUnderSpendCapIsAdmitted(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "under@example.com")
	c := newCaps(pool)
	if err := c.RecordSpend(context.Background(), llm.Response{CostUSD: 24.99}); err != nil {
		t.Fatal(err)
	}
	if r := mustAdmit(t, c, uid); r != RefusalNone {
		t.Fatalf("got %q", r)
	}
}

func TestAdmit_NewUTCDayResets(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "day@example.com")
	now := capsNow
	c := newCaps(pool)
	c.Now = func() time.Time { return now }
	c.PerAccount = 1
	if err := c.RecordSpend(context.Background(), llm.Response{CostUSD: 30}); err != nil {
		t.Fatal(err)
	}
	if r := mustAdmit(t, c, uid); r != RefusalSpendCap {
		t.Fatalf("same day: %q", r)
	}
	now = time.Date(2026, 10, 27, 23, 59, 59, 0, time.UTC)
	if r := mustAdmit(t, c, uid); r != RefusalSpendCap {
		t.Fatalf("last second of the day: %q", r)
	}
	now = time.Date(2026, 10, 28, 0, 0, 0, 0, time.UTC)
	if r := mustAdmit(t, c, uid); r != RefusalNone {
		t.Fatalf("next day: %q", r)
	}
	if r := mustAdmit(t, c, uid); r != RefusalAccountCap {
		t.Fatalf("next day second: %q", r)
	}
}

func TestAdmit_ConcurrentNeverExceedsTheCap(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "race@example.com")
	c := newCaps(pool)
	var admitted, refused atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := c.Admit(context.Background(), uid)
			switch {
			case err != nil:
				t.Error(err)
			case r == RefusalNone:
				admitted.Add(1)
			case r == RefusalAccountCap:
				refused.Add(1)
			default:
				t.Errorf("unexpected %q", r)
			}
		}()
	}
	wg.Wait()
	if admitted.Load() != 30 || refused.Load() != 10 {
		t.Fatalf("admitted %d refused %d", admitted.Load(), refused.Load())
	}
}

func TestAdmit_ZeroPerAccountRefuses(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "zero@example.com")
	c := newCaps(pool)
	c.PerAccount = 0
	if r := mustAdmit(t, c, uid); r != RefusalAccountCap {
		t.Fatalf("got %q", r)
	}
}

func TestRecordSpend_CountsEstimatedCallsWithMargin(t *testing.T) {
	pool := capsPool(t)
	c := newCaps(pool)
	ctx := context.Background()
	if err := c.RecordSpend(ctx, llm.Response{CostUSD: 1, CostEstimated: true}); err != nil {
		t.Fatal(err)
	}
	if err := c.RecordSpend(ctx, llm.Response{CostUSD: 1}); err != nil {
		t.Fatal(err)
	}
	var usd float64
	var calls, estimated int
	if err := pool.QueryRow(ctx, `select usd::float8, calls, estimated_calls from llm_spend where day = '2026-10-27'`).
		Scan(&usd, &calls, &estimated); err != nil {
		t.Fatal(err)
	}
	want := 1 + estimatedCostMargin
	if calls != 2 || estimated != 1 || usd < want-1e-6 || usd > want+1e-6 {
		t.Fatalf("usd %v calls %d estimated %d, want usd %v", usd, calls, estimated, want)
	}
}

func TestRecordSpend_ConcurrentUpsertsSum(t *testing.T) {
	pool := capsPool(t)
	c := newCaps(pool)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.RecordSpend(context.Background(), llm.Response{CostUSD: 0.5}); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	var usd float64
	var calls int
	if err := pool.QueryRow(context.Background(), `select usd::float8, calls from llm_spend`).Scan(&usd, &calls); err != nil || calls != 20 || usd != 10 {
		t.Fatalf("usd %v calls %d %v", usd, calls, err)
	}
}

func TestRecordSpend_RejectsNegativeCost(t *testing.T) {
	c := newCaps(nil)
	if err := c.RecordSpend(context.Background(), llm.Response{CostUSD: -1}); err == nil {
		t.Fatal("negative cost accepted")
	}
}
