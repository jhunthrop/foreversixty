package coach

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// estimatedCostMargin multiplies a cost the client estimated from token
// counts (llm.Response.CostEstimated) before it is added to the day's spend,
// so an estimate can only over-count against the cap, never under-count.
const estimatedCostMargin = 1.25

// Refusal says why a generation may not start; RefusalNone admits it.
type Refusal string

const (
	RefusalNone      Refusal = ""
	RefusalSwitchOff Refusal = "switch_off"
	RefusalSpendCap  Refusal = "spend_cap"
	// RefusalAccountCap is one account's cap: its generations or spend today, or its
	// spend this month. Verdict.Period says which period holds.
	RefusalAccountCap Refusal = "account_cap"
	// RefusalError is returned with a non-nil error: the caps could not be
	// checked, so nothing may be generated (fail closed to the rule-only review).
	RefusalError Refusal = "error"
)

// CapPeriod is the period an account cap holds for.
type CapPeriod string

const (
	CapPeriodNone  CapPeriod = ""
	CapPeriodDay   CapPeriod = "day"
	CapPeriodMonth CapPeriod = "month"
)

// Verdict is the caps' answer: the refusal, and for an account cap the period that holds
// and the UTC day it lifts (tomorrow, or the first of next month).
type Verdict struct {
	Refusal Refusal
	Period  CapPeriod
	Until   time.Time
}

// TierSource names the plan whose price bounds an account's spend (entitlements.Store's
// PlanTier): a guild tier, entitlements.PlanPremium, or "" for none.
type TierSource interface {
	PlanTier(ctx context.Context, userID int64) (string, error)
}

// Caps enforces the kill switch, the daily spend cap, and each account's caps: its daily
// generations, its daily spend and its monthly spend by plan tier. Days and months are UTC.
type Caps struct {
	DB          *pgxpool.Pool
	Enabled     bool
	SpendCapUSD float64
	PerAccount  int
	Account     config.AccountSpendCaps
	Tiers       TierSource
	Now         func() time.Time
}

// NewCaps builds Caps from the coach configuration so the switch and limits
// cannot drift from it. A nil pool, tier source or clock is an error, not a later panic.
func NewCaps(cfg config.CoachConfig, pool *pgxpool.Pool, tiers TierSource, now func() time.Time) (*Caps, error) {
	if pool == nil {
		return nil, errors.New("coach: caps need a database pool")
	}
	if tiers == nil {
		return nil, errors.New("coach: caps need a plan tier source")
	}
	if now == nil {
		return nil, errors.New("coach: caps need a clock")
	}
	return &Caps{
		DB: pool, Enabled: cfg.Enabled, SpendCapUSD: cfg.DailySpendCapUSD,
		PerAccount: cfg.DailyReviewsPerAccount, Account: cfg.AccountSpendCaps, Tiers: tiers, Now: now,
	}, nil
}

const spendSQL = `select coalesce((select usd::float8 from llm_spend where day = $1), 0)`

const generationsSQL = `select coalesce((select generations from coach_account_day where day = $1 and user_id = $2), 0)`

// accountSpendSQL is one account's spend today ($2) and over its month so far ($3..$2).
const accountSpendSQL = `
select coalesce(sum(usd) filter (where day = $2), 0)::float8, coalesce(sum(usd), 0)::float8
from coach_account_day where user_id = $1 and day between $3 and $2`

// admitSQL reserves one slot in a single statement: a first generation inserts
// the row, later ones increment it only while below the cap, and the row lock
// of the conflicting update serialises concurrent callers. No row returned
// means the account is at its cap.
const admitSQL = `
insert into coach_account_day (day, user_id, generations) values ($1, $2, 1)
on conflict (day, user_id) do update set generations = coach_account_day.generations + 1
  where coach_account_day.generations < $3
returning generations`

const recordSQL = `
insert into llm_spend (day, usd, calls, estimated_calls) values ($1, $2, 1, $3)
on conflict (day) do update set usd = llm_spend.usd + excluded.usd,
  calls = llm_spend.calls + 1, estimated_calls = llm_spend.estimated_calls + excluded.estimated_calls`

const recordAccountSQL = `
insert into coach_account_day (day, user_id, usd) values ($1, $2, $3)
on conflict (day, user_id) do update set usd = coach_account_day.usd + excluded.usd`

// Admit decides whether one generation may start for accountID now and, when it
// may, counts it against the account's UTC day in the same statement. The switch,
// today's spend and the account's spend are checked first; a spend refusal counts
// nothing. Cached reviews never reach Admit.
//
// Contract: the caller generates only when the verdict's refusal is RefusalNone and the
// error is nil. On a database error Admit returns RefusalError with the error, so a caller
// that checks only the refusal still fails closed.
func (c *Caps) Admit(ctx context.Context, accountID int64) (Verdict, error) {
	day := c.today()
	if v, err := c.precheck(ctx, day, accountID); v.Refusal != RefusalNone || err != nil {
		return v, err
	}
	var generations int
	err := c.DB.QueryRow(ctx, admitSQL, day, accountID, c.PerAccount).Scan(&generations)
	if errors.Is(err, pgx.ErrNoRows) {
		return c.accountCap(day, CapPeriodDay), nil
	}
	if err != nil {
		return Verdict{Refusal: RefusalError}, fmt.Errorf("coach: reserve a generation: %w", err)
	}
	return Verdict{}, nil
}

// Peek answers what Admit would for accountID now without counting anything: the switch,
// today's spend, the account's spend, then its generations so far. The fight route uses it
// to decide between queueing a generation and the busy answer (§15.4); the job's Admit
// still counts.
func (c *Caps) Peek(ctx context.Context, accountID int64) (Verdict, error) {
	day := c.today()
	if v, err := c.precheck(ctx, day, accountID); v.Refusal != RefusalNone || err != nil {
		return v, err
	}
	var generations int
	if err := c.DB.QueryRow(ctx, generationsSQL, day, accountID).Scan(&generations); err != nil {
		return Verdict{Refusal: RefusalError}, fmt.Errorf("coach: read the account's generations: %w", err)
	}
	if generations >= c.PerAccount {
		return c.accountCap(day, CapPeriodDay), nil
	}
	return Verdict{}, nil
}

// precheck is every refusal that counts nothing: the switch, the day's spend cap, a
// per-account generation cap below one, then the account's monthly and daily spend.
func (c *Caps) precheck(ctx context.Context, day time.Time, accountID int64) (Verdict, error) {
	if !c.Enabled {
		return Verdict{Refusal: RefusalSwitchOff}, nil
	}
	var spent float64
	if err := c.DB.QueryRow(ctx, spendSQL, day).Scan(&spent); err != nil {
		return Verdict{Refusal: RefusalError}, fmt.Errorf("coach: read today's spend: %w", err)
	}
	if spent >= c.SpendCapUSD {
		return Verdict{Refusal: RefusalSpendCap}, nil
	}
	if c.PerAccount < 1 {
		return c.accountCap(day, CapPeriodDay), nil
	}
	return c.accountSpendVerdict(ctx, day, accountID)
}

// accountSpendVerdict refuses an account whose month (by its plan tier) or day is spent;
// the month is named first, since it holds longer.
func (c *Caps) accountSpendVerdict(ctx context.Context, day time.Time, accountID int64) (Verdict, error) {
	tier, err := c.Tiers.PlanTier(ctx, accountID)
	if err != nil {
		return Verdict{Refusal: RefusalError}, fmt.Errorf("coach: read the account's plan tier: %w", err)
	}
	var today, month float64
	if err := c.DB.QueryRow(ctx, accountSpendSQL, accountID, day, monthStart(day)).Scan(&today, &month); err != nil {
		return Verdict{Refusal: RefusalError}, fmt.Errorf("coach: read the account's spend: %w", err)
	}
	switch {
	case month >= c.monthlyCap(tier):
		return c.accountCap(day, CapPeriodMonth), nil
	case today >= c.Account.DailyUSD:
		return c.accountCap(day, CapPeriodDay), nil
	}
	return Verdict{}, nil
}

// monthlyCap is tier's monthly cap; an account with no plan (a moderator) gets Premium's,
// the lowest.
func (c *Caps) monthlyCap(tier string) float64 {
	switch tier {
	case entitlements.TierRaidGuild:
		return c.Account.MonthlyRaidGuildUSD
	case entitlements.TierGuild:
		return c.Account.MonthlyGuildUSD
	}
	return c.Account.MonthlyPremiumUSD
}

// accountCap is the account-cap verdict for period, lifting the next UTC day or month.
func (c *Caps) accountCap(day time.Time, period CapPeriod) Verdict {
	until := day.AddDate(0, 0, 1)
	if period == CapPeriodMonth {
		until = monthStart(day).AddDate(0, 1, 0)
	}
	return Verdict{Refusal: RefusalAccountCap, Period: period, Until: until}
}

// RecordSpend adds one call's cost to today's llm_spend row and to accountID's day row, in
// one transaction. An estimated cost is multiplied by estimatedCostMargin first.
func (c *Caps) RecordSpend(ctx context.Context, accountID int64, res llm.Response) error {
	if res.CostUSD < 0 {
		return errors.New("coach: a negative call cost cannot be recorded")
	}
	if accountID <= 0 {
		return errors.New("coach: spend needs the account that opened the review")
	}
	cost, estimated := res.CostUSD, 0
	if res.CostEstimated {
		cost, estimated = cost*estimatedCostMargin, 1
	}
	day := c.today()
	err := pgx.BeginFunc(ctx, c.DB, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, recordSQL, day, cost, estimated); err != nil {
			return err
		}
		_, err := tx.Exec(ctx, recordAccountSQL, day, accountID, cost)
		return err
	})
	if err != nil {
		return fmt.Errorf("coach: record spend: %w", err)
	}
	return nil
}

// today is the current UTC calendar day, as a date at midnight UTC.
func (c *Caps) today() time.Time {
	now := c.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

// monthStart is the first of day's UTC month.
func monthStart(day time.Time) time.Time {
	return time.Date(day.Year(), day.Month(), 1, 0, 0, 0, 0, time.UTC)
}
