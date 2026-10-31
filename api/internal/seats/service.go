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
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// Refusals the routes map to HTTP statuses.
var (
	ErrNotPayer           = errors.New("seats: only the guild's payer can do that")
	ErrNoPlan             = errors.New("seats: the guild has no plan with seats")
	ErrNoSuchSlot         = errors.New("seats: that seat is outside the plan's seat count")
	ErrNotMember          = errors.New("seats: that account is not a verified member")
	ErrTaken              = errors.New("seats: the seat or the account already holds a seat")
	ErrNoOffer            = errors.New("seats: no offer to act on")
	ErrNotHeld            = errors.New("seats: the seat has no holder to move")
	ErrOfferPending       = errors.New("seats: an offer is waiting on that seat; withdraw it first")
	ErrClaimContested     = errors.New("seats: the guild's claim is contested")
	ErrBillingUnavailable = errors.New("seats: billing is not configured")
)

// ErrLocked refuses a move while the seat's lock runs.
type ErrLocked struct{ MovableAt time.Time }

func (e ErrLocked) Error() string {
	return "seats: the seat moves again on " + e.MovableAt.Format(time.RFC3339)
}

// Canceller stops a Stripe subscription at the end of its period, and
// resumes it. billing.Gateway satisfies it.
type Canceller interface {
	CancelSubscriptionAtPeriodEnd(ctx context.Context, id string) error
	// ResumeSubscription undoes the cancel when accepting a seat fails after it.
	ResumeSubscription(ctx context.Context, id string) error
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
	// Events publishes premium.seat to the member handed a seat; nil publishes nothing.
	Events events.Publisher
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

// send mails after commit; a client that has gone away must not cancel it.
func (s *Service) send(ctx context.Context, msgs []mail.Message) {
	if s.Mailer == nil {
		return
	}
	ctx = context.WithoutCancel(ctx)
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

// ownPremium is what a recipient pays for their own Premium.
type ownPremium struct {
	endsAt time.Time
}

// recipientPays is nil when the recipient pays for no Premium of their own.
func (s *Service) recipientPays(ctx context.Context, userID int64, now time.Time) (*ownPremium, error) {
	b, err := s.Store.Entitlements.PersonalBilling(ctx, userID)
	if err != nil {
		return nil, wrap("recipient billing", err)
	}
	if b == nil || !paysPremium(b.Source, b.Status) || !b.Active {
		return nil, nil
	}
	ends := now
	if b.CurrentPeriodEnd != nil {
		ends = *b.CurrentPeriodEnd
	}
	return &ownPremium{endsAt: ends}, nil
}

func loadSeat(ctx context.Context, tx pgx.Tx, guildID int64, slot int) (seatRow, error) {
	r, err := scanSeatRow(tx.QueryRow(ctx,
		`select `+seatColumns+` from premium_seats where guild_id = $1 and slot = $2 for update`, guildID, slot))
	if errors.Is(err, pgx.ErrNoRows) {
		return seatRow{slot: slot}, nil
	}
	return r, wrap("load seat", err)
}

// Give hands an empty slot to a member: a payer is offered it, anyone else
// holds it at once.
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
	seat, err := s.inLock(ctx, guildID, slot, now, func(tx pgx.Tx, row seatRow) ([]mail.Message, error) {
		if row.pendingID != nil {
			return nil, ErrOfferPending
		}
		if row.state != "" && row.state != StateEmpty {
			return nil, ErrTaken
		}
		return s.assign(ctx, tx, guildID, slot, actorID, toUserID, pays, now)
	})
	if err == nil {
		s.publishSeat(ctx, guildID, toUserID, seat, pays != nil)
	}
	return seat, err
}

// Move gives a held seat to someone else, once per lock period. A
// non-paying recipient takes it at once; a paying one is offered it and the
// holder keeps it until they accept.
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
	seat, err := s.inLock(ctx, guildID, slot, now, func(tx pgx.Tx, row seatRow) ([]mail.Message, error) {
		if row.state != StateHeld && row.state != StateAccepted {
			return nil, ErrNotHeld
		}
		if row.pendingID != nil {
			return nil, ErrOfferPending
		}
		if row.movableAt != nil && row.movableAt.After(now) {
			return nil, ErrLocked{MovableAt: *row.movableAt}
		}
		if row.holderID != nil && *row.holderID == toUserID {
			return nil, ErrTaken
		}
		if pays != nil {
			return s.offerMove(ctx, tx, guildID, row, actorID, toUserID, pays, now)
		}
		removed, err := s.release(ctx, tx, guildID, row, now)
		if err != nil {
			return nil, err
		}
		msgs, err := s.assign(ctx, tx, guildID, slot, actorID, toUserID, nil, now)
		if removed != nil {
			msgs = append(msgs, *removed)
		}
		return msgs, err
	})
	if err == nil {
		s.publishSeat(ctx, guildID, toUserID, seat, pays != nil)
	}
	return seat, err
}

// offerMove records toUserID as the seat's offeree, leaving the holder in place.
func (s *Service) offerMove(ctx context.Context, tx pgx.Tx, guildID int64, row seatRow, actorID, toUserID int64,
	pays *ownPremium, now time.Time) ([]mail.Message, error) {
	member, err := isVerifiedMember(ctx, tx, guildID, toUserID)
	if err != nil {
		return nil, err
	}
	if !member {
		return nil, ErrNotMember
	}
	if err := refuseSeated(ctx, tx, guildID, toUserID); err != nil {
		return nil, err
	}
	if _, err := tx.Exec(ctx, `update premium_seats set pending_user_id = $3, pending_by = $4, pending_offered_at = $5
		where guild_id = $1 and slot = $2`, guildID, row.slot, toUserID, actorID, now); err != nil {
		return nil, mapUnique("offer move", err)
	}
	msg, err := s.assignedEmail(ctx, tx, guildID, actorID, toUserID, pays)
	if err != nil {
		return nil, err
	}
	return []mail.Message{msg}, nil
}

// refuseSeated refuses a recipient who already holds, or waits on an offer
// of, any seat in the guild: accepting would collide with that seat.
func refuseSeated(ctx context.Context, tx pgx.Tx, guildID, userID int64) error {
	var seated bool
	if err := tx.QueryRow(ctx, `select exists (select 1 from premium_seats
		where guild_id = $1 and $2 in (holder_user_id, pending_user_id))`, guildID, userID).Scan(&seated); err != nil {
		return wrap("seated check", err)
	}
	if seated {
		return ErrTaken
	}
	return nil
}

const pgUniqueViolation = "23505"

func mapUnique(op string, err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
		return ErrTaken
	}
	return wrap(op, err)
}

// Withdraw takes back an offer that was not yet answered: a seat offered to
// its first holder empties, a move offer clears and the holder stays.
func (s *Service) Withdraw(ctx context.Context, guildID int64, slot int, actorID int64) (Seat, error) {
	tier, err := s.authorize(ctx, guildID, actorID)
	if err != nil {
		return Seat{}, err
	}
	if err := s.checkSlot(tier, slot); err != nil {
		return Seat{}, err
	}
	return s.inLock(ctx, guildID, slot, s.now(), func(tx pgx.Tx, row seatRow) ([]mail.Message, error) {
		switch {
		case row.pendingID != nil:
			return nil, wrap("withdraw move", clearPending(ctx, tx, guildID, slot))
		case row.state == StateOffered:
			_, err := tx.Exec(ctx, clearSeatSQL, guildID, slot, nil, nil)
			return nil, wrap("withdraw", err)
		}
		return nil, ErrNoOffer
	})
}

func clearPending(ctx context.Context, tx pgx.Tx, guildID int64, slot int) error {
	_, err := tx.Exec(ctx, `update premium_seats set pending_user_id = null, pending_by = null,
		pending_offered_at = null where guild_id = $1 and slot = $2`, guildID, slot)
	return err
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
	pays *ownPremium, now time.Time) ([]mail.Message, error) {
	member, err := isVerifiedMember(ctx, tx, guildID, toUserID)
	if err != nil {
		return nil, err
	}
	if !member {
		return nil, ErrNotMember
	}
	if err := refuseSeated(ctx, tx, guildID, toUserID); err != nil {
		return nil, err
	}
	state, offeredAt, givenAt := StateHeld, (*time.Time)(nil), &now
	movableAt := now.Add(catalog.SeatMoveLock)
	lock := &movableAt
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
		return nil, mapUnique("assign seat", err)
	}
	msg, err := s.assignedEmail(ctx, tx, guildID, actorID, toUserID, pays)
	if err != nil {
		return nil, err
	}
	return []mail.Message{msg}, nil
}

func (s *Service) assignedEmail(ctx context.Context, tx pgx.Tx, guildID, actorID, toUserID int64, pays *ownPremium) (mail.Message, error) {
	guildName, to, err := emailParties(ctx, tx, guildID, toUserID)
	if err != nil {
		return mail.Message{}, err
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

// offerRowSQL finds the seat userID has been offered: as its first holder, or
// as the offeree of a move.
const offerRowSQL = `select ` + seatColumns + ` from premium_seats
	where guild_id = $1 and ((holder_user_id = $2 and state = 'offered') or pending_user_id = $2)
	order by slot limit 1 for update`

// Accept takes the seat offered to userID. A member who pays for their own
// Premium has it set to end with its period and the seat starts then; the
// Stripe call is the last step before commit, undone if the commit fails.
func (s *Service) Accept(ctx context.Context, guildID, userID int64) error {
	if err := s.frozen(ctx, guildID, userID); err != nil {
		return err
	}
	b, err := s.Store.Entitlements.PersonalBilling(ctx, userID)
	if err != nil {
		return wrap("own billing", err)
	}
	paying := b != nil && paysPremium(b.Source, b.Status) && b.Active
	now := s.now()
	var outbox []mail.Message
	err = s.Store.withGuildLockStep(ctx, guildID, func(tx pgx.Tx) (*externalStep, error) {
		row, err := scanSeatRow(tx.QueryRow(ctx, offerRowSQL, guildID, userID))
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoOffer
		}
		if err != nil {
			return nil, wrap("load offer", err)
		}
		if err := s.offerStillStands(ctx, tx, guildID, userID, row); err != nil {
			return nil, err
		}
		if paying && s.Canceller == nil {
			return nil, ErrBillingUnavailable
		}
		return s.takeSeat(ctx, tx, guildID, userID, row, b, paying, now, &outbox)
	})
	if err != nil {
		return err
	}
	s.send(ctx, outbox)
	return nil
}

// offerStillStands re-checks, under the lock and before anything outside the
// database is touched, that the plan still has the slot and the offeree is
// still a verified member.
func (s *Service) offerStillStands(ctx context.Context, tx pgx.Tx, guildID, userID int64, row seatRow) error {
	tier, err := s.Store.activeTier(ctx, guildID)
	if err != nil {
		return err
	}
	if catalog.SeatsForTier(tier) == 0 {
		return ErrNoPlan
	}
	if err := s.checkSlot(tier, row.slot); err != nil {
		return err
	}
	member, err := isVerifiedMember(ctx, tx, guildID, userID)
	if err != nil {
		return err
	}
	if !member {
		return ErrNotMember
	}
	return nil
}

// takeSeat writes the acceptance. For a move offer the old holder is
// released first. It returns the Stripe cancel as the commit's last step.
func (s *Service) takeSeat(ctx context.Context, tx pgx.Tx, guildID, userID int64, row seatRow, b *entitlements.Billing,
	paying bool, now time.Time, outbox *[]mail.Message) (*externalStep, error) {
	giver := row.givenBy
	if row.pendingID != nil {
		giver = row.pendingBy
		removed, err := s.release(ctx, tx, guildID, row, now)
		if err != nil {
			return nil, err
		}
		if removed != nil {
			*outbox = append(*outbox, *removed)
		}
	}
	startsAt, cancelNow, gaveUp := now, false, (*string)(nil)
	if paying {
		cancelNow = !b.CancelAtPeriodEnd
		if b.CurrentPeriodEnd != nil {
			startsAt = *b.CurrentPeriodEnd
		}
		if cancelNow && b.PriceLookupKey != "" {
			gaveUp = &b.PriceLookupKey
		}
	}
	lockFrom := startsAt
	if now.After(lockFrom) {
		lockFrom = now
	}
	if _, err := tx.Exec(ctx, `update premium_seats set state = 'accepted', holder_user_id = $3, given_by = $4,
		starts_at = $5, given_at = $6, offered_at = null, movable_at = $7, holder_gave_up_lookup_key = $8,
		returned_name = null, returned_at = null, pending_user_id = null, pending_by = null, pending_offered_at = null
		where guild_id = $1 and slot = $2`,
		guildID, row.slot, userID, giver, startsAt, now, lockFrom.Add(catalog.SeatMoveLock), gaveUp); err != nil {
		return nil, mapUnique("accept seat", err)
	}
	if !cancelNow {
		return nil, nil
	}
	return s.cancelStep(b.StripeSubscriptionID), nil
}

func (s *Service) cancelStep(subscriptionID string) *externalStep {
	return &externalStep{
		do: func(ctx context.Context) error {
			return wrap("cancel own subscription", s.Canceller.CancelSubscriptionAtPeriodEnd(ctx, subscriptionID))
		},
		undo: func(ctx context.Context) {
			if err := s.Canceller.ResumeSubscription(ctx, subscriptionID); err != nil {
				s.logger().Error("seats", "op", "resume after failed accept", "subscription", subscriptionID, "err", err)
				return
			}
			s.logger().Error("seats", "op", "accept failed", "subscription", subscriptionID,
				"effect", "own subscription resumed")
		},
	}
}

// Decline turns down the seat offered to userID: a first offer empties the
// slot, a move offer clears and the holder stays.
func (s *Service) Decline(ctx context.Context, guildID, userID int64) error {
	if err := s.frozen(ctx, guildID, userID); err != nil {
		return err
	}
	return s.Store.withGuildLock(ctx, guildID, func(tx pgx.Tx) error {
		row, err := scanSeatRow(tx.QueryRow(ctx, offerRowSQL, guildID, userID))
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNoOffer
		}
		if err != nil {
			return wrap("load offer", err)
		}
		if row.pendingID != nil {
			return wrap("decline move", clearPending(ctx, tx, guildID, row.slot))
		}
		_, err = tx.Exec(ctx, clearSeatSQL, guildID, row.slot, nil, nil)
		return wrap("decline", err)
	})
}
