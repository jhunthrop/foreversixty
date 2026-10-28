// api/internal/billing/transfer.go
package billing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	stripe "github.com/stripe/stripe-go/v82"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// transferablePlan reports whether gb is a live whole plan billed through
// Stripe: the only kind a subscriber can hand to another officer.
func transferablePlan(gb *entitlements.GuildBilling) bool {
	return entitlements.LiveWholePlan(gb) && gb.StripeSubscriptionID != ""
}

// transferStartsAt is when a transfer's new subscriber first pays: the
// current period's end, floored at transferTrialFloor because Stripe
// refuses a Checkout trial_end closer than 48 hours. Checkout trials the
// new subscription to exactly this instant.
func transferStartsAt(periodEnd *time.Time, now time.Time) time.Time {
	start := now.Add(transferTrialFloor)
	if periodEnd != nil && periodEnd.After(start) {
		return *periodEnd
	}
	return start
}

// acceptedFromEarlier reports whether t is an accepted transfer away from a
// subscription other than gb's current one: its outgoing subscription may
// still be running out its last period.
func acceptedFromEarlier(t *planTransfer, gb *entitlements.GuildBilling) bool {
	return t != nil && t.state == transferAccepted && t.fromSubscriptionID != gb.StripeSubscriptionID
}

// transferView is the guild's transfer as the reader sees it: an offer or
// decline away from the current subscription, or an accepted transfer while
// the incoming subscription's trial covers the outgoing one's last period.
// Anything older is history and not shown. Display only: no Stripe call.
func (s *Service) transferView(ctx context.Context, a guildActor, gb *entitlements.GuildBilling) (*TransferView, error) {
	if gb == nil {
		return nil, nil
	}
	t, err := s.Store.GuildTransfer(ctx, a.guildID)
	if err != nil || t == nil {
		return nil, err
	}
	now := s.now()
	var startsAt time.Time
	switch {
	case acceptedFromEarlier(t, gb) && gb.TrialEnd != nil && gb.TrialEnd.After(now):
		startsAt = *gb.TrialEnd
	case t.state == transferAccepted || t.fromSubscriptionID != gb.StripeSubscriptionID:
		return nil, nil
	default:
		startsAt = transferStartsAt(gb.CurrentPeriodEnd, now)
	}
	v := &TransferView{
		ToYou: t.toUserID == a.userID, State: t.state, At: *rfc3339(&t.offeredAt),
		StartsAt: rfc3339(&startsAt), PriceFacts: catalog.FactsOf(gb.PriceLookupKey, gb.UnitAmountCents),
	}
	if t.decidedAt != nil {
		v.At = *rfc3339(t.decidedAt)
	}
	if v.FromName, err = s.displayName(ctx, a.guildID, t.fromUserID, ""); err != nil {
		return nil, err
	}
	if v.ToName, err = s.displayName(ctx, a.guildID, t.toUserID, ""); err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Service) sendTransfer(w http.ResponseWriter, r *http.Request) {
	if s.Gateway == nil {
		writeUnavailable(w, r)
		return
	}
	// The subscriber check waits for the plan read inside the lock, so a
	// guild with no plan answers guild_plan_required to every officer.
	a, ok := s.actingActor(w, r, "transfer", anyOfficer)
	if !ok {
		return
	}
	var in struct {
		ToUserID int64 `json:"to_user_id"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil || in.ToUserID <= 0 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a to_user_id", nil)
		return
	}
	err := s.Store.WithGuildLock(r.Context(), a.guildID, func(ctx context.Context) error {
		return s.offerTransfer(ctx, a, in.ToUserID)
	})
	s.answerTransfer(w, r, "transfer", err)
}

// offerTransfer is sendTransfer's locked half, serialized against the
// webhook's acceptance of an earlier offer. Every fact is re-read here.
func (s *Service) offerTransfer(ctx context.Context, a guildActor, toUserID int64) error {
	gb, err := s.Entitlements.GuildBilling(ctx, a.guildID)
	if err != nil {
		return err
	}
	if !transferablePlan(gb) {
		return &checkoutRefusal{status: http.StatusForbidden, code: "guild_plan_required",
			message: "this guild has no guild plan to transfer"}
	}
	if gb.BillingUserID == nil || *gb.BillingUserID != a.userID {
		return &checkoutRefusal{status: http.StatusForbidden, code: "forbidden",
			message: "only the plan's subscriber can transfer it"}
	}
	if err := s.checkTransferTarget(ctx, a, toUserID); err != nil {
		return err
	}
	if err := s.refuseWhileCheckoutOpen(ctx, a.guildID); err != nil {
		return err
	}
	current, err := s.Store.GuildTransfer(ctx, a.guildID)
	if err != nil {
		return err
	}
	if acceptedFromEarlier(current, gb) {
		if err := s.checkOutgoingEnded(ctx, current.fromSubscriptionID); err != nil {
			return err
		}
	}
	return s.Store.OfferTransfer(ctx, a.guildID, a.userID, toUserID, gb.StripeSubscriptionID)
}

// checkTransferTarget refuses anyone but another verified officer or leader.
func (s *Service) checkTransferTarget(ctx context.Context, a guildActor, toUserID int64) error {
	if toUserID == a.userID {
		return notMemberRefusal()
	}
	rank, verified, err := s.Accounts.GuildRank(ctx, a.guildID, toUserID)
	if err != nil {
		return err
	}
	if !verified || (rank != "officer" && rank != "leader") {
		return notMemberRefusal()
	}
	return nil
}

// checkOutgoingEnded refuses a new offer while an accepted transfer's
// outgoing subscription has not ended at Stripe: the accepted row is what
// keeps retiredByTransfer and stripe-reconcile from treating it as a stray,
// so it must not be overwritten until Stripe has ended it. A subscription
// Stripe no longer knows has ended too; refusing on it would block every
// offer for good.
func (s *Service) checkOutgoingEnded(ctx context.Context, subscriptionID string) error {
	ended, err := s.outgoingEnded(ctx, subscriptionID)
	if err != nil || ended {
		return err
	}
	return transferInProgress("the last transfer finishes when the previous subscription's period ends")
}

// outgoingEnded reports whether a transfer's outgoing subscription has
// ended at Stripe: canceled, expired unpaid, or no longer known there.
func (s *Service) outgoingEnded(ctx context.Context, subscriptionID string) (bool, error) {
	sub, err := s.Gateway.GetSubscription(ctx, subscriptionID)
	if errors.Is(err, ErrSubscriptionNotFound) {
		s.logger().Warn("billing", "op", "outgoing subscription missing at stripe",
			"subscription_id", subscriptionID, "effect", "treated as ended")
		return true, nil
	}
	if err != nil {
		return false, fmt.Errorf("billing: outgoing subscription %s: %w", subscriptionID, err)
	}
	return sub.Status == stripe.SubscriptionStatusCanceled || sub.Status == stripe.SubscriptionStatusIncompleteExpired, nil
}

// refuseWhileCheckoutOpen refuses a transfer or plan change while a
// Checkout for the guild is open: the subscription it makes was priced
// and pointed at the plan as it stood when the session began.
func (s *Service) refuseWhileCheckoutOpen(ctx context.Context, guildID int64) error {
	pending, err := s.Store.PendingGuildCheckout(ctx, guildID)
	if err != nil {
		return err
	}
	if pending {
		return transferInProgress("a checkout for this guild is open; try again once it finishes")
	}
	return nil
}

func transferInProgress(message string) *checkoutRefusal {
	return &checkoutRefusal{status: http.StatusConflict, code: "transfer_in_progress", message: message}
}

func notMemberRefusal() *checkoutRefusal {
	return &checkoutRefusal{status: http.StatusConflict, code: "not_member",
		message: "that account is not another verified officer of this guild"}
}

func (s *Service) withdrawTransfer(w http.ResponseWriter, r *http.Request) {
	a, ok := s.actingActor(w, r, "withdraw transfer", isSubscriber)
	if !ok {
		return
	}
	s.closeTransfer(w, r, "withdraw transfer", a, func(ctx context.Context) (bool, error) {
		return s.Store.WithdrawTransfer(ctx, a.guildID, a.userID)
	})
}

// declineTransfer skips the frozen-claimant check: declining an offer
// changes nobody's billing, so a contested claim is no reason to keep it.
func (s *Service) declineTransfer(w http.ResponseWriter, r *http.Request) {
	a, ok := s.guildActorFor(w, r, "decline transfer")
	if !ok {
		return
	}
	s.closeTransfer(w, r, "decline transfer", a, func(ctx context.Context) (bool, error) {
		return s.Store.DeclineTransfer(ctx, a.guildID, a.userID)
	})
}

// closeTransfer runs a withdraw or decline under the guild lock; no open
// offer to close answers 404.
func (s *Service) closeTransfer(w http.ResponseWriter, r *http.Request, op string, a guildActor,
	closeOffer func(context.Context) (bool, error)) {
	err := s.Store.WithGuildLock(r.Context(), a.guildID, func(ctx context.Context) error {
		if err := s.refuseWhileCheckoutOpen(ctx, a.guildID); err != nil {
			return err
		}
		closed, err := closeOffer(ctx)
		if err == nil && !closed {
			return &checkoutRefusal{status: http.StatusNotFound, code: "no_transfer_offer",
				message: "there is no open transfer offer"}
		}
		return err
	})
	s.answerTransfer(w, r, op, err)
}

// answerTransfer is every transfer write's answer: 204, a refusal, or 500.
func (s *Service) answerTransfer(w http.ResponseWriter, r *http.Request, op string, err error) {
	var refusal *checkoutRefusal
	switch {
	case err == nil:
		w.WriteHeader(http.StatusNoContent)
	case errors.As(err, &refusal):
		httpx.WriteError(w, r, refusal.status, refusal.code, refusal.message, refusal.fields)
	default:
		s.fail(w, r, op, err, "could not change the transfer just now")
	}
}
