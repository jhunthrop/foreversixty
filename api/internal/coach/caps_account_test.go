package coach

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// spendOn writes usd to the account's row for day, as RecordSpend would have.
func spendOn(t *testing.T, pool *pgxpool.Pool, day string, userID int64, usd float64) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `insert into coach_account_day (day, user_id, usd) values ($1, $2, $3)
		on conflict (day, user_id) do update set usd = coach_account_day.usd + excluded.usd`, day, userID, usd); err != nil {
		t.Fatal(err)
	}
}

func generationsOf(t *testing.T, pool *pgxpool.Pool, userID int64) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`select coalesce(sum(generations), 0) from coach_account_day where user_id = $1`, userID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestAdmit_AccountAtItsDailySpendCapIsRefusedAnotherIsAdmitted(t *testing.T) {
	pool := capsPool(t)
	a, b := capsUser(t, pool, "a@example.com"), capsUser(t, pool, "b@example.com")
	c := newCaps(pool)
	if err := c.RecordSpend(context.Background(), a, llm.Response{CostUSD: 0.15}); err != nil {
		t.Fatal(err)
	}
	v := mustAdmitVerdict(t, c, a)
	want := Verdict{Refusal: RefusalAccountCap, Period: CapPeriodDay, Until: time.Date(2026, 10, 28, 0, 0, 0, 0, time.UTC)}
	if v != want {
		t.Fatalf("got %+v, want %+v", v, want)
	}
	if n := generationsOf(t, pool, a); n != 0 {
		t.Fatalf("a refused spend counted %d generations", n)
	}
	if r := mustAdmit(t, c, b); r != RefusalNone {
		t.Fatalf("other account: %q", r)
	}
}

func TestAdmit_MonthlyCapSumsTheMonthsDays(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "month@example.com")
	c := newCaps(pool)
	spendOn(t, pool, "2026-09-30", uid, 5)    // last month: not counted
	spendOn(t, pool, "2026-10-28", uid, 5)    // after today: not counted
	spendOn(t, pool, "2026-10-01", uid, 0.5)  // each day under the daily cap...
	spendOn(t, pool, "2026-10-13", uid, 0.49) // ...and the month just under $1
	if r := mustAdmit(t, c, uid); r != RefusalNone {
		t.Fatalf("under the month: %q", r)
	}
	spendOn(t, pool, "2026-10-20", uid, 0.01)
	v := mustAdmitVerdict(t, c, uid)
	want := Verdict{Refusal: RefusalAccountCap, Period: CapPeriodMonth, Until: time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC)}
	if v != want {
		t.Fatalf("got %+v, want %+v", v, want)
	}
	c.Now = func() time.Time { return time.Date(2026, 11, 2, 9, 0, 0, 0, time.UTC) }
	if r := mustAdmit(t, c, uid); r != RefusalNone {
		t.Fatalf("next month: %q", r)
	}
}

func TestAdmit_MonthlyCapIsTheAccountsTier(t *testing.T) {
	pool := capsPool(t)
	c := newCaps(pool)
	cases := []struct {
		tier string
		cap  float64
	}{
		{"", testAccountCaps.MonthlyPremiumUSD},
		{entitlements.PlanPremium, testAccountCaps.MonthlyPremiumUSD},
		{entitlements.TierGuild, testAccountCaps.MonthlyGuildUSD},
		{entitlements.TierRaidGuild, testAccountCaps.MonthlyRaidGuildUSD},
	}
	tiers := fakeTiers{}
	c.Tiers = tiers
	for i, tc := range cases {
		uid := capsUser(t, pool, "tier"+string(rune('a'+i))+"@example.com")
		tiers[uid] = tc.tier
		spendOn(t, pool, "2026-10-01", uid, tc.cap-0.01)
		if r := mustAdmit(t, c, uid); r != RefusalNone {
			t.Fatalf("%q just under %v: %q", tc.tier, tc.cap, r)
		}
		spendOn(t, pool, "2026-10-02", uid, 0.01)
		if v := mustAdmitVerdict(t, c, uid); v.Refusal != RefusalAccountCap || v.Period != CapPeriodMonth {
			t.Fatalf("%q at %v: %+v", tc.tier, tc.cap, v)
		}
	}
}

func TestPeek_CountsNothing(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "peek@example.com")
	c := newCaps(pool)
	for i := 0; i < 3; i++ {
		if v, err := c.Peek(context.Background(), uid); err != nil || v.Refusal != RefusalNone {
			t.Fatalf("%+v %v", v, err)
		}
	}
	spendOn(t, pool, "2026-10-27", uid, 0.15)
	if v, err := c.Peek(context.Background(), uid); err != nil || v.Refusal != RefusalAccountCap || v.Period != CapPeriodDay {
		t.Fatalf("%+v %v", v, err)
	}
	var rows int
	if err := pool.QueryRow(context.Background(), `select count(*) from coach_account_day`).Scan(&rows); err != nil || rows != 1 {
		t.Fatalf("rows %d %v", rows, err)
	}
	if n := generationsOf(t, pool, uid); n != 0 {
		t.Fatalf("peek counted %d generations", n)
	}
}

func TestRecordSpend_WritesTheDayAndTheAccountTogether(t *testing.T) {
	pool := capsPool(t)
	uid := capsUser(t, pool, "both@example.com")
	c := newCaps(pool)
	ctx := context.Background()
	if err := c.RecordSpend(ctx, uid, llm.Response{CostUSD: 0.002, CostEstimated: true}); err != nil {
		t.Fatal(err)
	}
	var day, account float64
	if err := pool.QueryRow(ctx, `select (select usd::float8 from llm_spend where day = '2026-10-27'),
		(select usd::float8 from coach_account_day where day = '2026-10-27' and user_id = $1)`, uid).Scan(&day, &account); err != nil {
		t.Fatal(err)
	}
	want := 0.002 * estimatedCostMargin
	if math.Abs(day-want) > 1e-9 || math.Abs(account-want) > 1e-9 {
		t.Fatalf("day %v account %v, want %v", day, account, want)
	}
}

func TestRecordSpend_AFailedAccountWriteLeavesTheDayUntouched(t *testing.T) {
	pool := capsPool(t)
	c := newCaps(pool)
	ctx := context.Background()
	if err := c.RecordSpend(ctx, 987654321, llm.Response{CostUSD: 1}); err == nil {
		t.Fatal("spend for an unknown account recorded")
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from llm_spend`).Scan(&n); err != nil || n != 0 {
		t.Fatalf("llm_spend rows %d %v: the two writes are not one transaction", n, err)
	}
}
