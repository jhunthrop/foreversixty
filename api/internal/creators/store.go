package creators

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const uniqueViolation = "23505"

// Store is the program's tables (the creators migration).
type Store struct{ Pool *pgxpool.Pool }

// returningSQL is "player p is a returning player": signed in with Battle.net, a second UTC day
// inside 14 days (second_day_at is only ever stamped inside it), and a Forever character from
// Battle.net that no other account already qualified with. A player already paid stays returning.
const returningSQL = `(p.bounty_payout_id is not null or (
	p.second_day_at is not null
	and exists (select 1 from users u where u.id = p.user_id and u.bnet_sub is not null)
	and exists (select 1 from characters c where c.user_id = p.user_id and c.source = 'bnet'
	  and not exists (select 1 from creator_players q where q.qualified_character = c.key and q.user_id <> p.user_id))))`

// payingSQL is "player p paid in the last year".
const payingSQL = `exists (select 1 from creator_revenue r where r.user_id = p.user_id and r.paid_at > now() - interval '13 months')`

const creatorColumns = `c.id, c.name, c.channel_url, c.slug, c.code, c.starts_on, c.cap_cents, c.stripe_status,
	c.user_id, coalesce(c.stripe_account_id, ''), c.ended_at is not null`

func scanCreator(row pgx.Row) (Creator, error) {
	var c Creator
	var starts time.Time
	err := row.Scan(&c.ID, &c.Name, &c.ChannelURL, &c.Slug, &c.Code, &starts, &c.CapCents, &c.StripeStatus,
		&c.UserID, &c.StripeAccountID, &c.Ended)
	if errors.Is(err, pgx.ErrNoRows) {
		return Creator{}, ErrNotFound
	}
	if err != nil {
		return Creator{}, fmt.Errorf("creators: read creator: %w", err)
	}
	c.StartsOn = starts.Format(time.DateOnly)
	return c, nil
}

// Creator reads one creator.
func (s *Store) Creator(ctx context.Context, id int64) (Creator, error) {
	return scanCreator(s.Pool.QueryRow(ctx, `select `+creatorColumns+` from creators c where c.id = $1`, id))
}

// CreatorOfUser reads the creator a site account is bound to.
func (s *Store) CreatorOfUser(ctx context.Context, userID int64) (Creator, error) {
	c, err := scanCreator(s.Pool.QueryRow(ctx, `select `+creatorColumns+` from creators c where c.user_id = $1 and c.ended_at is null`, userID))
	if errors.Is(err, ErrNotFound) {
		return Creator{}, ErrNotCreator
	}
	return c, err
}

// IsCreator reports whether the account is bound to a creator (GET /v1/me's menu item).
func (s *Store) IsCreator(ctx context.Context, userID int64) (bool, error) {
	var ok bool
	if err := s.Pool.QueryRow(ctx, `select exists (select 1 from creators where user_id = $1 and ended_at is null)`, userID).Scan(&ok); err != nil {
		return false, fmt.Errorf("creators: is creator %d: %w", userID, err)
	}
	return ok, nil
}

// CodeIsLive reports whether code belongs to a creator in the program.
func (s *Store) CodeIsLive(ctx context.Context, code string) (bool, error) {
	var ok bool
	if err := s.Pool.QueryRow(ctx, `select exists (select 1 from creators where code = $1 and ended_at is null)`, code).Scan(&ok); err != nil {
		return false, fmt.Errorf("creators: code: %w", err)
	}
	return ok, nil
}

// SlugOrCodeTaken reports whether a new creator's link or code is already used.
func (s *Store) SlugOrCodeTaken(ctx context.Context, slug, code string) (bool, error) {
	var taken bool
	err := s.Pool.QueryRow(ctx, `select exists (select 1 from go_links where slug = $1)
		or exists (select 1 from creators where code = $2)`, slug, code).Scan(&taken)
	if err != nil {
		return false, fmt.Errorf("creators: taken: %w", err)
	}
	return taken, nil
}

// Create adds the creator and its go link in one transaction.
func (s *Store) Create(ctx context.Context, in Input, startsOn time.Time, capCents int64) (Creator, error) {
	var id int64
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `insert into go_links (slug, target, label) values ($1, '/', $2)`,
			in.Slug, "Creator: "+in.Name); err != nil {
			return err
		}
		return tx.QueryRow(ctx, `insert into creators (name, channel_url, slug, code, starts_on, cap_cents)
			values ($1, $2, $3, $4, $5, $6) returning id`,
			in.Name, in.ChannelURL, in.Slug, in.Code, startsOn, capCents).Scan(&id)
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return Creator{}, ErrTaken
	}
	if err != nil {
		return Creator{}, fmt.Errorf("creators: create: %w", err)
	}
	return s.Creator(ctx, id)
}

// Update changes what the owner may change after adding: name, channel, start date, cap.
func (s *Store) Update(ctx context.Context, id int64, name, channel string, startsOn time.Time, capCents int64) error {
	tag, err := s.Pool.Exec(ctx, `update creators set name = $2, channel_url = $3, starts_on = $4, cap_cents = $5, updated_at = now()
		where id = $1`, id, name, channel, startsOn, capCents)
	if err != nil {
		return fmt.Errorf("creators: update %d: %w", id, err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// SetStripeAccount records the creator's Express account.
func (s *Store) SetStripeAccount(ctx context.Context, id int64, accountID string) error {
	_, err := s.Pool.Exec(ctx, `update creators set stripe_account_id = $2, stripe_status = case when stripe_status = 'none'
		then 'onboarding' else stripe_status end, updated_at = now() where id = $1`, id, accountID)
	if err != nil {
		return fmt.Errorf("creators: set account %d: %w", id, err)
	}
	return nil
}

// SetStripeStatus records an account's status, by creator id or by account id.
func (s *Store) SetStripeStatus(ctx context.Context, accountID, status string) error {
	_, err := s.Pool.Exec(ctx, `update creators set stripe_status = $2, stripe_checked_at = now(), updated_at = now()
		where stripe_account_id = $1`, accountID, status)
	if err != nil {
		return fmt.Errorf("creators: set status: %w", err)
	}
	return nil
}

// NewInvite replaces the creator's unused invites with one whose hash is given.
func (s *Store) NewInvite(ctx context.Context, creatorID int64, hash []byte, expires time.Time) error {
	return pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `delete from creator_invites where creator_id = $1 and used_at is null`, creatorID); err != nil {
			return fmt.Errorf("creators: clear invites: %w", err)
		}
		if _, err := tx.Exec(ctx, `insert into creator_invites (token_hash, creator_id, expires_at) values ($1, $2, $3)`,
			hash, creatorID, expires); err != nil {
			return fmt.Errorf("creators: invite: %w", err)
		}
		return nil
	})
}

// InviteCreator is the creator an open invite names, for the page to greet them by name.
func (s *Store) InviteCreator(ctx context.Context, hash []byte, now time.Time) (Creator, error) {
	c, err := scanCreator(s.Pool.QueryRow(ctx, `select `+creatorColumns+` from creators c
		join creator_invites i on i.creator_id = c.id
		where i.token_hash = $1 and i.used_at is null and i.expires_at > $2`, hash, now))
	if errors.Is(err, ErrNotFound) {
		return Creator{}, ErrBadInvite
	}
	return c, err
}

// forgetSelf drops a creator's own account from their unsettled lines: as a player, and its
// payments and refunds (spec section 1, self-referral). A line already paid stays paid.
func forgetSelf(ctx context.Context, tx pgx.Tx, creatorID, userID int64) error {
	for _, sql := range []string{
		`delete from creator_refunds f using creator_revenue r where f.invoice_id = r.invoice_id
		  and r.user_id = $1 and r.creator_id = $2 and f.payout_id is null`,
		`delete from creator_revenue r where r.user_id = $1 and r.creator_id = $2 and r.payout_id is null
		  and not exists (select 1 from creator_refunds f where f.invoice_id = r.invoice_id)`,
		`delete from creator_players where user_id = $1 and creator_id = $2 and bounty_payout_id is null`,
	} {
		if _, err := tx.Exec(ctx, sql, userID, creatorID); err != nil {
			return err
		}
	}
	return nil
}

// AcceptInvite binds the account to the invite's creator, once. The account stops counting as
// anyone's player of that creator (spec section 1, self-referral).
func (s *Store) AcceptInvite(ctx context.Context, hash []byte, userID int64, now time.Time) (Creator, error) {
	var creatorID int64
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		err := tx.QueryRow(ctx, `update creator_invites set used_at = $2
			where token_hash = $1 and used_at is null and expires_at > $2 returning creator_id`, hash, now).Scan(&creatorID)
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrBadInvite
		}
		if err != nil {
			return err
		}
		tag, err := tx.Exec(ctx, `update creators set user_id = $2, updated_at = now() where id = $1 and ended_at is null`, creatorID, userID)
		if err != nil {
			return err
		}
		if tag.RowsAffected() == 0 {
			return ErrBadInvite
		}
		return forgetSelf(ctx, tx, creatorID, userID)
	})
	var pgErr *pgconn.PgError
	switch {
	case errors.As(err, &pgErr) && pgErr.Code == uniqueViolation:
		return Creator{}, ErrTaken
	case errors.Is(err, ErrBadInvite):
		return Creator{}, err
	case err != nil:
		return Creator{}, fmt.Errorf("creators: accept invite: %w", err)
	}
	return s.Creator(ctx, creatorID)
}

// attributionTarget is what an attribution decision reads about the account and the creator.
type attributionTarget struct {
	creatorID     int64
	startsOn      time.Time
	creatorUserID *int64
	signedUpAt    time.Time
	attributed    bool
}

func (s *Store) target(ctx context.Context, where string, key any, userID int64) (attributionTarget, bool, error) {
	var t attributionTarget
	err := s.Pool.QueryRow(ctx, `select c.id, c.starts_on, c.user_id, u.created_at,
		exists (select 1 from creator_players p where p.user_id = u.id)
		from creators c, users u where `+where+` and c.ended_at is null and u.id = $2`, key, userID).
		Scan(&t.creatorID, &t.startsOn, &t.creatorUserID, &t.signedUpAt, &t.attributed)
	if errors.Is(err, pgx.ErrNoRows) {
		return t, false, nil
	}
	if err != nil {
		return t, false, fmt.Errorf("creators: attribution target: %w", err)
	}
	return t, true, nil
}

func (t attributionTarget) self(userID int64) bool {
	return t.creatorUserID != nil && *t.creatorUserID == userID
}

func (s *Store) attribute(ctx context.Context, t attributionTarget, userID int64, source string, touched time.Time) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `insert into creator_players (user_id, creator_id, source, touched_at, signed_up_at)
		values ($1, $2, $3, $4, $5) on conflict (user_id) do nothing`, userID, t.creatorID, source, touched, t.signedUpAt)
	if err != nil {
		return false, fmt.Errorf("creators: attribute %d: %w", userID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// AttributeLink makes the account a player of the creator whose link it arrived from, when the
// rules allow (LinkCounts, not the creator, not already anyone's). It reports whether it did.
func (s *Store) AttributeLink(ctx context.Context, userID int64, slug string, clickedAt time.Time) (bool, error) {
	t, ok, err := s.target(ctx, `c.slug = $1`, slug, userID)
	if ok && clickedAt.After(t.signedUpAt) {
		clickedAt = t.signedUpAt // no click time was sent: the arrival came with the sign-up
	}
	if err != nil || !ok || t.attributed || t.self(userID) || !LinkCounts(t.startsOn, clickedAt, t.signedUpAt) {
		return false, err
	}
	return s.attribute(ctx, t, userID, "link", clickedAt)
}

// AttributeCode makes the account a player of the creator whose code it redeemed at checkout.
func (s *Store) AttributeCode(ctx context.Context, userID int64, code string, at time.Time) (bool, error) {
	t, ok, err := s.target(ctx, `c.code = $1`, code, userID)
	if err != nil || !ok || t.attributed || t.self(userID) || !CodeCounts(t.startsOn, t.signedUpAt) {
		return false, err
	}
	return s.attribute(ctx, t, userID, "code", at)
}

// MarkVisit stamps a player's second day (SecondDay's rule, in SQL) and reports whether it did.
func (s *Store) MarkVisit(ctx context.Context, userID int64, seen time.Time) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update creator_players set second_day_at = $2
		where user_id = $1 and second_day_at is null
		  and $2 < signed_up_at + make_interval(secs => $3)
		  and ($2 at time zone 'utc')::date > (signed_up_at at time zone 'utc')::date`,
		userID, seen, ReturnWindow.Seconds())
	if err != nil {
		return false, fmt.Errorf("creators: visit %d: %w", userID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// PlayerCreator is the creator an account belongs to, false when none.
func (s *Store) PlayerCreator(ctx context.Context, userID int64) (int64, bool, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `select creator_id from creator_players where user_id = $1`, userID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("creators: player %d: %w", userID, err)
	}
	return id, true, nil
}

// HasRevenue reports whether an invoice is already recorded (a redelivered webhook).
func (s *Store) HasRevenue(ctx context.Context, invoiceID string) (bool, error) {
	var ok bool
	if err := s.Pool.QueryRow(ctx, `select exists (select 1 from creator_revenue where invoice_id = $1)`, invoiceID).Scan(&ok); err != nil {
		return false, fmt.Errorf("creators: has revenue: %w", err)
	}
	return ok, nil
}

// RecordRevenue records a player's paid invoice inside the 12-month window from their first
// payment, and reports whether it did.
func (s *Store) RecordRevenue(ctx context.Context, userID int64, invoiceID string, m InvoiceMoney) (bool, error) {
	base := PaymentBase(m.NetCents, m.TaxCents, m.TaxInFees)
	var recorded bool
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		var creatorID int64
		var first *time.Time
		err := tx.QueryRow(ctx, `select creator_id, first_paid_at from creator_players where user_id = $1 for update`, userID).
			Scan(&creatorID, &first)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		if err != nil {
			return err
		}
		if first != nil && !m.PaidAt.Before(first.AddDate(0, ShareMonths, 0)) {
			return nil
		}
		tag, err := tx.Exec(ctx, `insert into creator_revenue (invoice_id, creator_id, user_id, charge_id, paid_at,
			amount_cents, tax_cents, fee_cents, base_cents, share_cents)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) on conflict do nothing`,
			invoiceID, creatorID, userID, m.ChargeID, m.PaidAt, m.AmountCents, m.TaxCents, m.FeeCents, base, ShareOf(base))
		if err != nil {
			return err
		}
		recorded = tag.RowsAffected() == 1
		if !recorded {
			return nil
		}
		_, err = tx.Exec(ctx, `update creator_players set first_paid_at = least(coalesce(first_paid_at, $2), $2) where user_id = $1`,
			userID, m.PaidAt)
		return err
	})
	if err != nil {
		return false, fmt.Errorf("creators: record revenue %s: %w", invoiceID, err)
	}
	return recorded, nil
}

// revenueOfCharge is the recorded payment a refund belongs to.
type revenueOfCharge struct {
	invoiceID   string
	creatorID   int64
	amountCents int64
	taxCents    int64
}

// RevenueOfCharge reads the recorded payment a charge paid, false when it is none of a creator's.
func (s *Store) RevenueOfCharge(ctx context.Context, chargeID string) (revenueOfCharge, bool, error) {
	var r revenueOfCharge
	err := s.Pool.QueryRow(ctx, `select invoice_id, creator_id, amount_cents, tax_cents from creator_revenue where charge_id = $1`,
		chargeID).Scan(&r.invoiceID, &r.creatorID, &r.amountCents, &r.taxCents)
	if errors.Is(err, pgx.ErrNoRows) {
		return r, false, nil
	}
	if err != nil {
		return r, false, fmt.Errorf("creators: revenue of charge: %w", err)
	}
	return r, true, nil
}

// RecordRefunds records each refund once, as a line at half its taxless amount.
func (s *Store) RecordRefunds(ctx context.Context, r revenueOfCharge, refunds []Refund) error {
	for _, f := range refunds {
		amount := RefundBase(f.AmountCents, r.amountCents, r.taxCents)
		if _, err := s.Pool.Exec(ctx, `insert into creator_refunds (refund_id, invoice_id, creator_id, refunded_at, amount_cents, share_cents)
			values ($1, $2, $3, $4, $5, $6) on conflict do nothing`,
			f.ID, r.invoiceID, r.creatorID, f.Created, amount, ShareOf(amount)); err != nil {
			return fmt.Errorf("creators: record refund %s: %w", f.ID, err)
		}
	}
	return nil
}
