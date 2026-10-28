package coach

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// estimatedCostMargin multiplies a cost the client estimated from token
// counts (llm.Response.CostEstimated) before it is added to the day's spend,
// so an estimate can only over-count against the cap, never under-count.
const estimatedCostMargin = 1.25

// Refusal says why a generation may not start; RefusalNone admits it.
type Refusal string

const (
	RefusalNone       Refusal = ""
	RefusalSwitchOff  Refusal = "switch_off"
	RefusalSpendCap   Refusal = "spend_cap"
	RefusalAccountCap Refusal = "account_cap"
	// RefusalError is returned with a non-nil error: the caps could not be
	// checked, so nothing may be generated (fail closed to the rule-only review).
	RefusalError Refusal = "error"
)

// Caps enforces the kill switch, the daily spend cap and the per-account
// daily generation cap. Days are UTC days.
type Caps struct {
	DB          *pgxpool.Pool
	Enabled     bool
	SpendCapUSD float64
	PerAccount  int
	Now         func() time.Time
}

// NewCaps builds Caps from the coach configuration so the switch and limits
// cannot drift from it. A nil pool or clock is an error, not a later panic.
func NewCaps(cfg config.CoachConfig, pool *pgxpool.Pool, now func() time.Time) (*Caps, error) {
	if pool == nil {
		return nil, errors.New("coach: caps need a database pool")
	}
	if now == nil {
		return nil, errors.New("coach: caps need a clock")
	}
	return &Caps{
		DB: pool, Enabled: cfg.Enabled, SpendCapUSD: cfg.DailySpendCapUSD,
		PerAccount: cfg.DailyReviewsPerAccount, Now: now,
	}, nil
}

const spendSQL = `select coalesce((select usd::float8 from llm_spend where day = $1), 0)`

const generationsSQL = `select coalesce((select generations from coach_account_day where day = $1 and user_id = $2), 0)`

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

// Admit decides whether one generation may start for accountID now and, when it
// may, counts it against the account's UTC day in the same statement. Spend is
// checked from the recorded costs first; a spend refusal counts nothing. Cached
// reviews never reach Admit.
//
// Contract: the caller generates only when the result is (RefusalNone, nil). On
// a database error Admit returns (RefusalError, err), so a caller that checks
// only the refusal still fails closed.
func (c *Caps) Admit(ctx context.Context, accountID int64) (Refusal, error) {
	if !c.Enabled {
		return RefusalSwitchOff, nil
	}
	day := c.today()
	if refusal, err := c.dayRefusal(ctx, day); refusal != RefusalNone || err != nil {
		return refusal, err
	}
	var generations int
	err := c.DB.QueryRow(ctx, admitSQL, day, accountID, c.PerAccount).Scan(&generations)
	if errors.Is(err, pgx.ErrNoRows) {
		return RefusalAccountCap, nil
	}
	if err != nil {
		return RefusalError, fmt.Errorf("coach: reserve a generation: %w", err)
	}
	return RefusalNone, nil
}

// Peek answers what Admit would for accountID now without counting anything: the switch,
// today's spend, then the account's generations so far. The fight route uses it to decide
// between queueing a generation and the busy answer (§15.4); the job's Admit still counts.
func (c *Caps) Peek(ctx context.Context, accountID int64) (Refusal, error) {
	if !c.Enabled {
		return RefusalSwitchOff, nil
	}
	day := c.today()
	if refusal, err := c.dayRefusal(ctx, day); refusal != RefusalNone || err != nil {
		return refusal, err
	}
	var generations int
	if err := c.DB.QueryRow(ctx, generationsSQL, day, accountID).Scan(&generations); err != nil {
		return RefusalError, fmt.Errorf("coach: read the account's generations: %w", err)
	}
	if generations >= c.PerAccount {
		return RefusalAccountCap, nil
	}
	return RefusalNone, nil
}

// dayRefusal is the refusal that holds for every account today: the spend cap, or a
// per-account cap below one.
func (c *Caps) dayRefusal(ctx context.Context, day time.Time) (Refusal, error) {
	var spent float64
	if err := c.DB.QueryRow(ctx, spendSQL, day).Scan(&spent); err != nil {
		return RefusalError, fmt.Errorf("coach: read today's spend: %w", err)
	}
	if spent >= c.SpendCapUSD {
		return RefusalSpendCap, nil
	}
	if c.PerAccount < 1 {
		return RefusalAccountCap, nil
	}
	return RefusalNone, nil
}

// RecordSpend adds one call's cost to today's llm_spend row in one upsert.
// An estimated cost is multiplied by estimatedCostMargin first.
func (c *Caps) RecordSpend(ctx context.Context, res llm.Response) error {
	if res.CostUSD < 0 {
		return errors.New("coach: a negative call cost cannot be recorded")
	}
	cost, estimated := res.CostUSD, 0
	if res.CostEstimated {
		cost, estimated = cost*estimatedCostMargin, 1
	}
	if _, err := c.DB.Exec(ctx, recordSQL, c.today(), cost, estimated); err != nil {
		return fmt.Errorf("coach: record spend: %w", err)
	}
	return nil
}

// today is the current UTC calendar day, as a date at midnight UTC.
func (c *Caps) today() time.Time {
	now := c.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
