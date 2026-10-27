// api/internal/billing/transfer.go
package billing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// transferablePlan reports whether gb is a live whole plan billed through
// Stripe: the only kind a subscriber can hand to another officer.
func transferablePlan(gb *entitlements.GuildBilling) bool {
	return gb != nil && gb.Active && !gb.Scoped() && gb.StripeSubscriptionID != ""
}

// outgoingStillBilling reports whether an accepted transfer's outgoing
// subscription is still in its last period. The incoming subscription's
// trial runs to that period's end (or later), so while it lasts the
// accepted row is what keeps stripe-reconcile from flagging the outgoing
// subscription as an orphan.
func outgoingStillBilling(t *planTransfer, gb *entitlements.GuildBilling, now time.Time) bool {
	return t != nil && t.state == transferAccepted && t.fromSubscriptionID != gb.StripeSubscriptionID &&
		gb.TrialEnd != nil && gb.TrialEnd.After(now)
}

// transferView is the guild's transfer as the reader sees it: an offer or
// decline away from the current subscription, or an accepted transfer whose
// outgoing subscription still runs. Anything older is history and not shown.
func (s *Service) transferView(ctx context.Context, a guildActor, gb *entitlements.GuildBilling) (*TransferView, error) {
	if gb == nil {
		return nil, nil
	}
	t, err := s.Store.GuildTransfer(ctx, a.guildID)
	if err != nil || t == nil {
		return nil, err
	}
	startsAt := gb.CurrentPeriodEnd
	switch {
	case outgoingStillBilling(t, gb, s.now()):
		startsAt = gb.TrialEnd
	case t.state == transferAccepted || t.fromSubscriptionID != gb.StripeSubscriptionID:
		return nil, nil
	}
	v := &TransferView{
		ToYou: t.toUserID == a.userID, State: t.state, At: *rfc3339(&t.offeredAt),
		StartsAt: rfc3339(startsAt), PriceFacts: priceFactsOf(gb.PriceLookupKey, gb.UnitAmountCents),
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
	// The subscriber check waits for the plan read, so a guild with no plan
	// answers guild_plan_required to every officer.
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
// webhook's acceptance of an earlier offer.
func (s *Service) offerTransfer(ctx context.Context, a guildActor, toUserID int64) error {
	gb, err := s.Entitlements.GuildBilling(ctx, a.guildID)
	if err != nil {
		return err
	}
	if !transferablePlan(gb) {
		return &checkoutRefusal{status: http.StatusForbidden, code: "guild_plan_required",
			message: "this guild has no guild plan to transfer"}
	}
	if !a.subscriber {
		return &checkoutRefusal{status: http.StatusForbidden, code: "forbidden",
			message: "only the plan's subscriber can transfer it"}
	}
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
	current, err := s.Store.GuildTransfer(ctx, a.guildID)
	if err != nil {
		return err
	}
	if outgoingStillBilling(current, gb, s.now()) {
		return &checkoutRefusal{status: http.StatusConflict, code: "transfer_in_progress",
			message: "the last transfer finishes when the previous subscription's period ends"}
	}
	return s.Store.OfferTransfer(ctx, a.guildID, a.userID, toUserID, gb.StripeSubscriptionID)
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
		return s.Store.WithdrawTransfer(ctx, a.guildID)
	})
}

func (s *Service) declineTransfer(w http.ResponseWriter, r *http.Request) {
	a, ok := s.actingActor(w, r, "decline transfer", anyOfficer)
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
