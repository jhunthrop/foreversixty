package creators

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// visitBudget bounds the background second-day stamp a signed-in request starts.
const visitBudget = 5 * time.Second

// LinkArrival is the tracked links' hook (golinks): an account's first link, sent with its first
// signed-in page, with when the click was. It attributes the account when the rules allow.
func (s *Service) LinkArrival(ctx context.Context, userID int64, slug string, clickedAt time.Time) error {
	ok, err := s.Store.AttributeLink(ctx, userID, slug, clickedAt)
	if err != nil {
		return err
	}
	if ok {
		s.logger().Info("creators", "op", "attribute", "source", "link", "user_id", userID)
	}
	return nil
}

// Visited is the authenticator's hook (auth.VisitHook): at most once per account per UTC day per
// instance it stamps a player's second day, in the background, so the request never waits.
func (s *Service) Visited(u auth.User) {
	now := s.now()
	day := now.Format(time.DateOnly)
	if last, seen := s.visits.Load(u.ID); seen && last == day {
		return
	}
	s.visits.Store(u.ID, day)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), visitBudget)
		defer cancel()
		if _, err := s.Store.MarkVisit(ctx, u.ID, now); err != nil {
			s.visits.Delete(u.ID)
			s.logger().Warn("creators", "op", "visit", "user_id", u.ID, "err", err)
		}
	}()
}

// CodeIsLive reports whether code is a creator's code checkout may honour (a creator in the program).
func (s *Service) CodeIsLive(ctx context.Context, code string) (bool, error) {
	return s.Store.CodeIsLive(ctx, NormalizeCode(code))
}

// CodeTrialDays is the free month a creator's code gives (billing's checkout reads it).
func (s *Service) CodeTrialDays() int64 { return CodeTrialDays }

// CheckoutCompleted is billing's hook on a completed Checkout: the creator code checkout carried in
// the subscription's metadata attributes the buyer, and the session's first invoice is recorded at
// once, because Stripe may deliver invoice.paid before this event and that invoice would be missed.
func (s *Service) CheckoutCompleted(ctx context.Context, userID int64, code string, at time.Time, invoiceID string) error {
	if code != "" {
		ok, err := s.Store.AttributeCode(ctx, userID, NormalizeCode(code), at)
		if err != nil {
			return err
		}
		if ok {
			s.logger().Info("creators", "op", "attribute", "source", "code", "user_id", userID)
		}
	}
	if invoiceID == "" {
		return nil
	}
	return s.InvoicePaid(ctx, userID, invoiceID)
}

// InvoicePaid is billing's hook on invoice.paid for the paying account. Only a creator's player
// costs a Stripe read; anyone else is one database lookup.
func (s *Service) InvoicePaid(ctx context.Context, userID int64, invoiceID string) error {
	if s.Gateway == nil {
		return nil
	}
	_, ok, err := s.Store.PlayerCreator(ctx, userID)
	if err != nil || !ok {
		return err
	}
	done, err := s.Store.HasRevenue(ctx, invoiceID)
	if err != nil || done {
		return err
	}
	money, err := s.Gateway.InvoiceMoney(ctx, invoiceID)
	if err != nil {
		return err
	}
	if money.AmountCents <= 0 {
		return nil
	}
	recorded, err := s.Store.RecordRevenue(ctx, userID, invoiceID, money)
	if err != nil {
		return err
	}
	if !recorded {
		return nil
	}
	s.logger().Info("creators", "op", "revenue", "user_id", userID)
	// A refund that landed before this payment was recorded found nothing to attach to.
	return s.ChargeRefunded(ctx, money.ChargeID)
}

// ChargeRefunded is billing's hook on charge.refunded: a refund of a recorded payment becomes a line.
func (s *Service) ChargeRefunded(ctx context.Context, chargeID string) error {
	if s.Gateway == nil {
		return nil
	}
	r, ok, err := s.Store.RevenueOfCharge(ctx, chargeID)
	if err != nil || !ok {
		return err
	}
	refunds, err := s.Gateway.Refunds(ctx, chargeID)
	if err != nil {
		return err
	}
	return s.Store.RecordRefunds(ctx, r, refunds)
}

// AccountUpdated is billing's hook on a connected account's account.updated: the status is
// re-read from Stripe rather than trusted from the event.
func (s *Service) AccountUpdated(ctx context.Context, accountID string) error {
	if s.Gateway == nil || accountID == "" {
		return nil
	}
	_, err := s.syncAccount(ctx, accountID)
	return err
}
