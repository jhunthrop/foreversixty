package creators

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// ErrNothingOwed: the month owes nothing, so there is nothing to approve.
var ErrNothingOwed = errors.New("creators: nothing owed this month")

// ErrBelowMinimum: the month owes less than MinPayoutCents; it carries over to the next month.
var ErrBelowMinimum = errors.New("creators: under the minimum payout; it carries over")

type querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// payoutRow is a recorded payout.
type payoutRow struct {
	ID          int64
	Status      string
	BountyCount int64
	Bounty      int64
	Share       int64
	Refund      int64
	Capped      int64
	Total       int64
	TransferID  *string
	PaidAt      *time.Time
}

func readPayout(ctx context.Context, q querier, creatorID int64, m Month, lock bool) (payoutRow, bool, error) {
	var p payoutRow
	sql := `select id, status, bounty_count, bounty_cents, share_cents, refund_cents, capped_cents, total_cents, transfer_id, paid_at
		from creator_payouts where creator_id = $1 and month = $2`
	if lock {
		sql += ` for update`
	}
	err := q.QueryRow(ctx, sql, creatorID, m.Start).Scan(&p.ID, &p.Status, &p.BountyCount, &p.Bounty, &p.Share, &p.Refund,
		&p.Capped, &p.Total, &p.TransferID, &p.PaidAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, false, nil
	}
	if err != nil {
		return p, false, fmt.Errorf("creators: read payout: %w", err)
	}
	return p, true, nil
}

// lineSet is a statement's lines with the keys that settle them.
type lineSet struct {
	lines    []Line
	players  []int64
	invoices []string
	refunds  []string
}

// statementLines are the creator's unsettled lines dated before end, or, with payoutID, the
// lines that payout settled.
func statementLines(ctx context.Context, q querier, creatorID int64, end time.Time, payoutID *int64) (lineSet, error) {
	var set lineSet
	bountyWhere := `p.bounty_payout_id is null and ` + returningSQL + ` and p.second_day_at < $2`
	revenueWhere := `payout_id is null and paid_at < $2`
	refundWhere := `payout_id is null and refunded_at < $2`
	var arg any = end
	if payoutID != nil {
		bountyWhere, revenueWhere, refundWhere = `p.bounty_payout_id = $2`, `payout_id = $2`, `payout_id = $2`
		arg = *payoutID
	}
	reads := []struct {
		kind, sql string
		keep      func(key string, player int64)
	}{
		{LineBounty, `select p.user_id::text, p.user_id, coalesce(p.second_day_at, p.created_at), ` + fmt.Sprint(BountyCents) + `::bigint
			from creator_players p where p.creator_id = $1 and ` + bountyWhere,
			func(_ string, player int64) { set.players = append(set.players, player) }},
		{LineShare, `select invoice_id, coalesce(user_id, 0), paid_at, share_cents from creator_revenue
			where creator_id = $1 and ` + revenueWhere,
			func(key string, _ int64) { set.invoices = append(set.invoices, key) }},
		{LineRefund, `select f.refund_id, coalesce(r.user_id, 0), f.refunded_at, f.share_cents from creator_refunds f
			join creator_revenue r on r.invoice_id = f.invoice_id where f.creator_id = $1 and f.` + refundWhere,
			func(key string, _ int64) { set.refunds = append(set.refunds, key) }},
	}
	for _, r := range reads {
		rows, err := q.Query(ctx, r.sql+` order by 3`, creatorID, arg)
		if err != nil {
			return set, fmt.Errorf("creators: %s lines: %w", r.kind, err)
		}
		for rows.Next() {
			var key string
			var player, amount int64
			var on time.Time
			if err := rows.Scan(&key, &player, &on, &amount); err != nil {
				rows.Close()
				return set, fmt.Errorf("creators: %s lines: %w", r.kind, err)
			}
			r.keep(key, player)
			set.lines = append(set.lines, Line{Kind: r.kind, Player: player, On: on.UTC().Format(time.DateOnly), AmountCents: amount})
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return set, fmt.Errorf("creators: %s lines: %w", r.kind, err)
		}
	}
	return set, nil
}

// sums totals a line set under the cap.
func sums(set lineSet, capCents int64) Statement {
	var st Statement
	for _, l := range set.lines {
		switch l.Kind {
		case LineBounty:
			st.BountyCount++
			st.BountyCents += l.AmountCents
		case LineShare:
			st.ShareCents += l.AmountCents
		case LineRefund:
			st.RefundCents += l.AmountCents
		}
	}
	st.TotalCents, st.CappedCents = Settle(st.BountyCents, st.ShareCents, st.RefundCents, capCents)
	st.Lines = set.lines
	if st.Lines == nil {
		st.Lines = []Line{}
	}
	return st
}

// Statement is the creator's month as the owner sees it at now.
func (s *Store) Statement(ctx context.Context, c Creator, m Month, now time.Time) (Statement, error) {
	p, found, err := readPayout(ctx, s.Pool, c.ID, m, false)
	if err != nil {
		return Statement{}, err
	}
	var st Statement
	if found {
		set, err := statementLines(ctx, s.Pool, c.ID, m.End(), &p.ID)
		if err != nil {
			return Statement{}, err
		}
		st = Statement{BountyCount: p.BountyCount, BountyCents: p.Bounty, ShareCents: p.Share, RefundCents: p.Refund,
			CappedCents: p.Capped, TotalCents: p.Total, PaidAt: p.PaidAt, TransferID: p.TransferID, Lines: set.lines}
		if st.Lines == nil {
			st.Lines = []Line{}
		}
		st.Status = StatementReady // a pending payout: Approve again finishes the transfer
		if p.Status != "pending" {
			st.Status = StatementPaid
		}
	} else {
		set, err := statementLines(ctx, s.Pool, c.ID, m.End(), nil)
		if err != nil {
			return Statement{}, err
		}
		st = sums(set, c.CapCents)
		st.Status = openStatus(c, m, now, st.TotalCents)
	}
	st.CreatorID, st.Name, st.Month = c.ID, c.Name, m.String()
	return st, nil
}

func openStatus(c Creator, m Month, now time.Time, total int64) string {
	switch {
	case !m.Closed(now):
		return StatementOpen
	case total == 0:
		return StatementNothing
	case total < MinPayoutCents:
		return StatementCarry
	case c.StripeStatus != StatusReady:
		return StatementBlocked
	default:
		return StatementReady
	}
}

// Claim settles a closed month's lines under a pending payout, or answers the payout already
// recorded for it (pending: the transfer is still to make; paid: done). It holds the creator's
// row for the transaction, so two approvals of the same month never both claim.
func (s *Store) Claim(ctx context.Context, creatorID int64, m Month, now time.Time) (payoutRow, Creator, error) {
	var out payoutRow
	var c Creator
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		var err error
		c, err = scanCreator(tx.QueryRow(ctx, `select `+creatorColumns+` from creators c where c.id = $1 for update`, creatorID))
		if err != nil {
			return err
		}
		p, found, err := readPayout(ctx, tx, creatorID, m, true)
		if err != nil || found {
			out = p
			return err
		}
		if !m.Closed(now) {
			return ErrMonthOpen
		}
		set, err := statementLines(ctx, tx, creatorID, m.End(), nil)
		if err != nil {
			return err
		}
		st := sums(set, c.CapCents)
		if st.TotalCents == 0 {
			return ErrNothingOwed
		}
		if st.TotalCents < MinPayoutCents {
			return ErrBelowMinimum
		}
		if c.StripeStatus != StatusReady || c.StripeAccountID == "" {
			return ErrNotPayable
		}
		out = payoutRow{Status: "pending", BountyCount: st.BountyCount, Bounty: st.BountyCents, Share: st.ShareCents,
			Refund: st.RefundCents, Capped: st.CappedCents, Total: st.TotalCents}
		if err := tx.QueryRow(ctx, `insert into creator_payouts (creator_id, month, status, bounty_count, bounty_cents,
			share_cents, refund_cents, capped_cents, total_cents) values ($1, $2, 'pending', $3, $4, $5, $6, $7, $8) returning id`,
			creatorID, m.Start, out.BountyCount, out.Bounty, out.Share, out.Refund, out.Capped, out.Total).Scan(&out.ID); err != nil {
			return err
		}
		return settleLines(ctx, tx, out.ID, set)
	})
	if err != nil && !errors.Is(err, ErrMonthOpen) && !errors.Is(err, ErrNothingOwed) && !errors.Is(err, ErrBelowMinimum) &&
		!errors.Is(err, ErrNotPayable) &&
		!errors.Is(err, ErrNotFound) {
		return out, c, fmt.Errorf("creators: claim %d %s: %w", creatorID, m, err)
	}
	return out, c, err
}

// settleLines marks every line of the set as the payout's, and stamps each paid player's
// qualifying character so it can never qualify a second account.
func settleLines(ctx context.Context, tx pgx.Tx, payoutID int64, set lineSet) error {
	if _, err := tx.Exec(ctx, `update creator_players p set bounty_payout_id = $1,
		qualified_character = (select min(c.key) from characters c where c.user_id = p.user_id and c.source = 'bnet'
		  and not exists (select 1 from creator_players q where q.qualified_character = c.key))
		where p.user_id = any($2)`, payoutID, set.players); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `update creator_revenue set payout_id = $1 where invoice_id = any($2)`, payoutID, set.invoices); err != nil {
		return err
	}
	_, err := tx.Exec(ctx, `update creator_refunds set payout_id = $1 where refund_id = any($2)`, payoutID, set.refunds)
	return err
}

// Release undoes a pending payout whose transfer Stripe refused: its lines are unsettled again and
// the payout row goes, so the statement is exactly as before the approve.
func (s *Store) Release(ctx context.Context, payoutID int64) error {
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		for _, sql := range []string{
			`update creator_players set bounty_payout_id = null, qualified_character = null where bounty_payout_id = $1`,
			`update creator_revenue set payout_id = null where payout_id = $1`,
			`update creator_refunds set payout_id = null where payout_id = $1`,
			`delete from creator_payouts where id = $1 and status = 'pending'`,
		} {
			if _, err := tx.Exec(ctx, sql, payoutID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("creators: release payout %d: %w", payoutID, err)
	}
	return nil
}

// MarkPaid records the transfer that paid a pending payout.
func (s *Store) MarkPaid(ctx context.Context, payoutID int64, transferID string, now time.Time) error {
	_, err := s.Pool.Exec(ctx, `update creator_payouts set status = 'paid', transfer_id = $2, paid_at = $3
		where id = $1 and status = 'pending'`, payoutID, transferID, now)
	if err != nil {
		return fmt.Errorf("creators: mark paid %d: %w", payoutID, err)
	}
	return nil
}

// Creators is every creator in the program (withEnded: and every removed one) with its figures at now.
func (s *Store) Creators(ctx context.Context, now time.Time, withEnded bool) ([]Creator, error) {
	rows, err := s.Pool.Query(ctx, `select `+creatorColumns+` from creators c where c.ended_at is null or $1 order by c.created_at, c.id`, withEnded)
	if err != nil {
		return nil, fmt.Errorf("creators: list: %w", err)
	}
	var list []Creator
	for rows.Next() {
		c, err := scanCreator(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		list = append(list, c)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("creators: list: %w", err)
	}
	out := make([]Creator, 0, len(list))
	for _, c := range list {
		withFigures, err := s.withFigures(ctx, c, now)
		if err != nil {
			return nil, err
		}
		out = append(out, withFigures)
	}
	return out, nil
}

func (s *Store) withFigures(ctx context.Context, c Creator, now time.Time) (Creator, error) {
	f := &c.Figures
	err := s.Pool.QueryRow(ctx, `select
		coalesce((select sum(clicks) from go_link_clicks where slug = $2), 0)::bigint,
		(select count(*) from creator_players p where p.creator_id = $1),
		(select count(*) from creator_players p where p.creator_id = $1 and `+returningSQL+`),
		coalesce((select sum(base_cents) from creator_revenue where creator_id = $1), 0)::bigint
		  - coalesce((select sum(amount_cents) from creator_refunds where creator_id = $1), 0)::bigint`,
		c.ID, c.Slug).Scan(&f.Clicks, &f.SignUps, &f.Returning, &f.RevenueCents)
	if err != nil {
		return c, fmt.Errorf("creators: figures %d: %w", c.ID, err)
	}
	m := MonthOf(now)
	st, err := s.Statement(ctx, c, m, now)
	if err != nil {
		return c, err
	}
	f.Month, f.OwedCents, f.Capped = m.String(), st.TotalCents, st.CappedCents > 0
	return c, nil
}

// Me is the creator's own page at now.
func (s *Store) Me(ctx context.Context, c Creator, now time.Time) (Me, error) {
	me := Me{Name: c.Name, Slug: c.Slug, Code: c.Code, CapCents: c.CapCents, StripeStatus: c.StripeStatus,
		Weeks: []Week{}, Payouts: []Payout{}}
	m := MonthOf(now)
	st, err := s.Statement(ctx, c, m, now)
	if err != nil {
		return me, err
	}
	me.Month, me.MonthCents = m.String(), st.TotalCents
	if err := s.Pool.QueryRow(ctx, `select count(*), count(*) filter (where `+returningSQL+`), count(*) filter (where `+payingSQL+`)
		from creator_players p where p.creator_id = $1`, c.ID).Scan(&me.Players, &me.Returning, &me.Paying); err != nil {
		return me, fmt.Errorf("creators: me counts: %w", err)
	}
	if me.Weeks, err = s.weeks(ctx, c.ID, now); err != nil {
		return me, err
	}
	me.Payouts, err = s.payouts(ctx, c.ID)
	return me, err
}

func (s *Store) weeks(ctx context.Context, creatorID int64, now time.Time) ([]Week, error) {
	rows, err := s.Pool.Query(ctx, `select date_trunc('week', p.signed_up_at at time zone 'utc')::date as wk, count(*),
		count(*) filter (where `+returningSQL+`), count(*) filter (where `+payingSQL+`)
		from creator_players p where p.creator_id = $1 group by wk order by wk desc limit $2`, creatorID, RecentWeeks)
	if err != nil {
		return nil, fmt.Errorf("creators: weeks: %w", err)
	}
	defer rows.Close()
	running := WeekOf(now)
	out := []Week{}
	for rows.Next() {
		var wk time.Time
		var signed, returning, paying int64
		if err := rows.Scan(&wk, &signed, &returning, &paying); err != nil {
			return nil, fmt.Errorf("creators: weeks: %w", err)
		}
		w := Week{WeekOf: wk.Format(time.DateOnly), SignedUp: signed}
		if !wk.Equal(running) {
			w.Returning, w.Paying = &returning, &paying
		}
		out = append(out, w)
	}
	return out, rows.Err()
}

func (s *Store) payouts(ctx context.Context, creatorID int64) ([]Payout, error) {
	rows, err := s.Pool.Query(ctx, `select month, total_cents, paid_at from creator_payouts
		where creator_id = $1 and status = 'paid' order by month desc`, creatorID)
	if err != nil {
		return nil, fmt.Errorf("creators: payouts: %w", err)
	}
	defer rows.Close()
	out := []Payout{}
	for rows.Next() {
		var month, paid time.Time
		var p Payout
		if err := rows.Scan(&month, &p.AmountCents, &paid); err != nil {
			return nil, fmt.Errorf("creators: payouts: %w", err)
		}
		p.Month, p.PaidAt = month.Format("2006-01"), paid.UTC()
		out = append(out, p)
	}
	return out, rows.Err()
}
