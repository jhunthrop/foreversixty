package seats

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/billingmail"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// Refusals the routes map to HTTP statuses.
var (
	ErrNotPayer           = errors.New("seats: only the guild's payer can do that")
	ErrNoPlan             = errors.New("seats: the guild has no plan with seats")
	ErrNoSuchSlot         = errors.New("seats: no such seat")
	ErrNotMember          = errors.New("seats: that account is not a verified member")
	ErrTaken              = errors.New("seats: the seat or the account already holds a seat")
	ErrNoOffer            = errors.New("seats: no offer to act on")
	ErrNotHeld            = errors.New("seats: the seat has no holder to move")
	ErrClaimContested     = errors.New("seats: the guild's claim is contested")
	ErrBillingUnavailable = errors.New("seats: billing is not configured")
)

// ErrLocked refuses a move while the seat's lock runs.
type ErrLocked struct{ MovableAt time.Time }

func (e ErrLocked) Error() string {
	return "seats: the seat moves again on " + e.MovableAt.Format(time.RFC3339)
}

// Canceller stops a Stripe subscription at the end of its period.
// billing.Gateway satisfies it.
type Canceller interface {
	CancelSubscriptionAtPeriodEnd(ctx context.Context, id string) error
}

// FrozenClaimant is guilds.Store's freeze hook; nil makes the check a no-op.
type FrozenClaimant interface {
	FrozenClaimant(ctx context.Context, guildID, userID int64) (bool, error)
}

// Service gives, moves, offers and releases seats.
type Service struct {
	Store *Store
	// Mailer sends the seat emails; nil sends nothing.
	Mailer mail.Mailer
	// Canceller is nil when Stripe is not configured.
	Canceller     Canceller
	Guilds        FrozenClaimant
	PublicBaseURL string
	Log           *slog.Logger
	// Now is the clock; nil means time.Now.
	Now func() time.Time
}

func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func wrap(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("seats: %s: %w", op, err)
}

func (s *Service) send(ctx context.Context, msgs []mail.Message) {
	if s.Mailer == nil {
		return
	}
	for _, m := range msgs {
		if err := s.Mailer.Send(ctx, m); err != nil {
			s.logger().Error("seats", "op", "send email", "subject", m.Subject, "err", err)
		}
	}
}

func (s *Service) accountURL() string { return s.PublicBaseURL + "/account" }

// frozen refuses a disputed claimant.
func (s *Service) frozen(ctx context.Context, guildID, userID int64) error {
	if s.Guilds == nil {
		return nil
	}
	frozen, err := s.Guilds.FrozenClaimant(ctx, guildID, userID)
	if err != nil {
		return wrap("freeze check", err)
	}
	if frozen {
		return ErrClaimContested
	}
	return nil
}

// authorize is the check every payer action starts with: not frozen, a
// payer, and a plan with seats. It returns the plan's tier.
func (s *Service) authorize(ctx context.Context, guildID, actorID int64) (string, error) {
	if err := s.frozen(ctx, guildID, actorID); err != nil {
		return "", err
	}
	ok, err := s.Store.IsPayer(ctx, guildID, actorID)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrNotPayer
	}
	tier, err := s.Store.activeTier(ctx, guildID)
	if err != nil {
		return "", err
	}
	if catalog.SeatsForTier(tier) == 0 {
		return "", ErrNoPlan
	}
	return tier, nil
}

func (s *Service) checkSlot(tier string, slot int) error {
	if slot < 1 || slot > catalog.SeatsForTier(tier) {
		return ErrNoSuchSlot
	}
	return nil
}

// payingRow is the recipient's own live Stripe Premium, nil when they pay for none.
type payingRow struct {
	endsAt time.Time
}

func (s *Service) recipientPays(ctx context.Context, userID int64, now time.Time) (*payingRow, error) {
	b, err := s.Store.Entitlements.PersonalBilling(ctx, userID)
	if err != nil {
		return nil, wrap("recipient billing", err)
	}
	if b == nil || b.Source != "stripe" || !b.Active {
		return nil, nil
	}
	ends := now
	if b.CurrentPeriodEnd != nil {
		ends = *b.CurrentPeriodEnd
	}
	return &payingRow{endsAt: ends}, nil
}

func loadSeat(ctx context.Context, tx pgx.Tx, guildID int64, slot int) (seatRow, error) {
	r, err := scanSeatRow(tx.QueryRow(ctx,
		`select `+seatColumns+` from premium_seats where guild_id = $1 and slot = $2 for update`, guildID, slot))
	if errors.Is(err, pgx.ErrNoRows) {
		return seatRow{slot: slot}, nil
	}
	return r, wrap("load seat", err)
}

// Give hands a slot to a member: a payer is offered it, anyone else holds
// it at once.
func (s *Service) Give(ctx context.Context, guildID int64, slot int, actorID, toUserID int64) (Seat, error) {
	tier, err := s.authorize(ctx, guildID, actorID)
	if err != nil {
		return Seat{}, err
	}
	if err := s.checkSlot(tier, slot); err != nil {
		return Seat{}, err
	}
	now := s.now()
	pays, err := s.recipientPays(ctx, toUserID, now)
	if err != nil {
		return Seat{}, err
	}
	return s.inLock(ctx, guildID, slot, now, func(tx pgx.Tx, row seatRow) ([]mail.Message, error) {
		if row.state != "" && row.state != StateEmpty {
			return nil, ErrTaken
		}
		return s.assign(ctx, tx, guildID, slot, actorID, toUserID, pays, now)
	})
}

// Move releases a held seat's holder and gives it to someone else, once per
// lock period.
func (s *Service) Move(ctx context.Context, guildID int64, slot int, actorID, toUserID int64) (Seat, error) {
	tier, err := s.authorize(ctx, guildID, actorID)
	if err != nil {
		return Seat{}, err
	}
	if err := s.checkSlot(tier, slot); err != nil {
		return Seat{}, err
	}
	now := s.now()
	pays, err := s.recipientPays(ctx, toUserID, now)
	if err != nil {
		return Seat{}, err
	}
	return s.inLock(ctx, guildID, slot, now, func(tx pgx.Tx, row seatRow) ([]mail.Message, error) {
		if row.state != StateHeld && row.state != StateAccepted {
			return nil, ErrNotHeld
		}
		if row.movableAt != nil && row.movableAt.After(now) {
			return nil, ErrLocked{MovableAt: *row.movableAt}
		}
		if row.holderID != nil && *row.holderID == toUserID {
			return nil, ErrTaken
		}
		removed, err := s.release(ctx, tx, guildID, row, now)
		if err != nil {
			return nil, err
		}
		msgs, err := s.assign(ctx, tx, guildID, slot, actorID, toUserID, pays, now)
		if removed != nil {
			msgs = append(msgs, *removed)
		}
		return msgs, err
	})
}

// Withdraw takes back an offer that was not yet answered.
func (s *Service) Withdraw(ctx context.Context, guildID int64, slot int, actorID int64) (Seat, error) {
	tier, err := s.authorize(ctx, guildID, actorID)
	if err != nil {
		return Seat{}, err
	}
	if err := s.checkSlot(tier, slot); err != nil {
		return Seat{}, err
	}
	return s.inLock(ctx, guildID, slot, s.now(), func(tx pgx.Tx, row seatRow) ([]mail.Message, error) {
		if row.state != StateOffered {
			return nil, ErrNoOffer
		}
		_, err := tx.Exec(ctx, clearSeatSQL, guildID, slot, nil, nil)
		return nil, wrap("withdraw", err)
	})
}

// inLock runs fn on the slot's row under the guild lock, sends the emails fn
// returns after the commit, and answers the slot as it then reads.
func (s *Service) inLock(ctx context.Context, guildID int64, slot int, now time.Time,
	fn func(tx pgx.Tx, row seatRow) ([]mail.Message, error)) (Seat, error) {
	var seat Seat
	var outbox []mail.Message
	err := s.Store.withGuildLock(ctx, guildID, func(tx pgx.Tx) error {
		row, err := loadSeat(ctx, tx, guildID, slot)
		if err != nil {
			return err
		}
		if outbox, err = fn(tx, row); err != nil {
			return err
		}
		after, err := loadSeat(ctx, tx, guildID, slot)
		if err != nil {
			return err
		}
		seat, err = s.Store.view(ctx, tx, guildID, slot, after, now)
		return err
	})
	if err != nil {
		return Seat{}, err
	}
	s.send(ctx, outbox)
	return seat, nil
}

// assign writes the slot for toUserID: offered to a payer, held by anyone
// else. It returns the email to send after commit.
func (s *Service) assign(ctx context.Context, tx pgx.Tx, guildID int64, slot int, actorID, toUserID int64,
	pays *payingRow, now time.Time) ([]mail.Message, error) {
	var member bool
	if err := tx.QueryRow(ctx, `select exists (select 1 from guild_members
		where guild_id = $1 and user_id = $2 and verified_at is not null)`, guildID, toUserID).Scan(&member); err != nil {
		return nil, wrap("recipient membership", err)
	}
	if !member {
		return nil, ErrNotMember
	}
	state, offeredAt, givenAt, movableAt := StateHeld, (*time.Time)(nil), &now, now.Add(catalog.SeatMoveLock)
	var lock *time.Time = &movableAt
	if pays != nil {
		state, offeredAt, givenAt, lock = StateOffered, &now, nil, nil
	}
	if _, err := tx.Exec(ctx, `
		insert into premium_seats (guild_id, slot, state, holder_user_id, given_by, given_at, offered_at, movable_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (guild_id, slot) do update set state = excluded.state,
		  holder_user_id = excluded.holder_user_id, given_by = excluded.given_by,
		  given_at = excluded.given_at, offered_at = excluded.offered_at, movable_at = excluded.movable_at,
		  starts_at = null, holder_gave_up_lookup_key = null, returned_name = null, returned_at = null`,
		guildID, slot, state, toUserID, actorID, givenAt, offeredAt, lock); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrTaken
		}
		return nil, wrap("assign seat", err)
	}
	msg, err := s.assignedEmail(ctx, tx, guildID, actorID, toUserID, pays)
	if err != nil {
		return nil, err
	}
	return []mail.Message{msg}, nil
}

func (s *Service) assignedEmail(ctx context.Context, tx pgx.Tx, guildID, actorID, toUserID int64, pays *payingRow) (mail.Message, error) {
	var guildName, to string
	if err := tx.QueryRow(ctx, `select g.name, u.email from guilds g, users u where g.id = $1 and u.id = $2`,
		guildID, toUserID).Scan(&guildName, &to); err != nil {
		return mail.Message{}, wrap("email parties", err)
	}
	from, err := characterOf(ctx, tx, guildID, actorID)
	if err != nil {
		return mail.Message{}, err
	}
	fromName := from.name
	if fromName == "" {
		fromName = "A guild officer"
	}
	if pays != nil {
		return billingmail.SeatOfferedEmail(billingmail.SeatOffered{
			To: to, Guild: guildName, From: fromName, OwnEndsOn: pays.endsAt, AccountURL: s.accountURL()}), nil
	}
	return billingmail.SeatGivenEmail(billingmail.SeatGiven{
		To: to, Guild: guildName, From: fromName, AccountURL: s.accountURL()}), nil
}

// Accept takes the seat offered to userID. When they pay for their own
// Premium it is set to end with its period and the seat starts then.
func (s *Service) Accept(ctx context.Context, guildID, userID int64) error {
	if err := s.frozen(ctx, guildID, userID); err != nil {
		return err
	}
	b, err := s.Store.Entitlements.PersonalBilling(ctx, userID)
	if err != nil {
		return wrap("own billing", err)
	}
	paying := b != nil && b.Source == "stripe" && b.Active
	if paying && s.Canceller == nil {
		return ErrBillingUnavailable
	}
	now := s.now()
	return s.Store.withGuildLock(ctx, guildID, func(tx pgx.Tx) error {
		row, err := scanSeatRow(tx.QueryRow(ctx, `select `+seatColumns+` from premium_seats
			where guild_id = $1 and holder_user_id = $2 for update`, guildID, userID))
		if errors.Is(err, pgx.ErrNoRows) || (err == nil && row.state != StateOffered) {
			return ErrNoOffer
		}
		if err != nil {
			return wrap("load offer", err)
		}
		startsAt, gaveUp := now, (*string)(nil)
		if paying {
			if startsAt, gaveUp, err = s.stepDown(ctx, b.StripeSubscriptionID, b.CancelAtPeriodEnd, b.CurrentPeriodEnd, b.PriceLookupKey, now); err != nil {
				return err
			}
		}
		_, err = tx.Exec(ctx, `update premium_seats set state = 'accepted', starts_at = $3, given_at = $4,
			movable_at = $5, holder_gave_up_lookup_key = $6 where guild_id = $1 and slot = $2`,
			guildID, row.slot, startsAt, now, now.Add(catalog.SeatMoveLock), gaveUp)
		return wrap("accept seat", err)
	})
}

// stepDown ends the holder's own subscription at its period end and
// returns when the seat starts and which price they gave up.
func (s *Service) stepDown(ctx context.Context, subscriptionID string, alreadyCanceling bool,
	periodEnd *time.Time, lookupKey string, now time.Time) (time.Time, *string, error) {
	if !alreadyCanceling {
		if err := s.Canceller.CancelSubscriptionAtPeriodEnd(ctx, subscriptionID); err != nil {
			return time.Time{}, nil, wrap("cancel own subscription", err)
		}
	}
	startsAt := now
	if periodEnd != nil {
		startsAt = *periodEnd
	}
	var key *string
	if lookupKey != "" {
		key = &lookupKey
	}
	return startsAt, key, nil
}

// Decline turns down the seat offered to userID.
func (s *Service) Decline(ctx context.Context, guildID, userID int64) error {
	if err := s.frozen(ctx, guildID, userID); err != nil {
		return err
	}
	return s.Store.withGuildLock(ctx, guildID, func(tx pgx.Tx) error {
		tag, err := tx.Exec(ctx, `update premium_seats set state = 'empty', holder_user_id = null, given_by = null,
			given_at = null, offered_at = null, starts_at = null, movable_at = null, holder_gave_up_lookup_key = null
			where guild_id = $1 and holder_user_id = $2 and state = 'offered'`, guildID, userID)
		if err != nil {
			return wrap("decline", err)
		}
		if tag.RowsAffected() == 0 {
			return ErrNoOffer
		}
		return nil
	})
}
