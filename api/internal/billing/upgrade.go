// api/internal/billing/upgrade.go
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

// previewValidFor is how long an upgrade preview's proration date stays good
// for the upgrade itself; past it the amounts shown may no longer be what
// Stripe charges.
//
// The date is not signed or tied to a preview: any instant in the last
// previewValidFor is accepted. That is deliberate. Only a payer can upgrade,
// the charge lands on the guild's own invoice, and Stripe bounds the
// proration date to the current period, so a forged date can only move the
// payer's own proration by minutes.
const previewValidFor = 10 * time.Minute

// prorationClockSkew tolerates a proration date slightly ahead of our clock.
const prorationClockSkew = time.Minute

// UpgradePreview is GET /v1/guilds/{id}/billing/upgrade-preview.
// ProrationDate (unix seconds) goes back in the upgrade's body.
type UpgradePreview struct {
	AmountDueCents  int64  `json:"amount_due_cents"`
	NextAmountCents int64  `json:"next_amount_cents"`
	Interval        string `json:"interval"`
	RenewsAt        string `json:"renews_at"`
	ProrationDate   int64  `json:"proration_date"`
}

// upgradeTarget is a payer's Guild-tier Stripe plan and the Raid guild price
// it moves to.
type upgradeTarget struct {
	actor     guildActor
	plan      *entitlements.GuildBilling
	lookupKey string
	priceID   string
}

// priceChangeError is a ChangePrice failure, told apart from a refusal and
// from an internal error once the guild lock is released.
type priceChangeError struct{ err error }

func (e priceChangeError) Error() string { return e.err.Error() }
func (e priceChangeError) Unwrap() error { return e.err }

// upgradable reports whether gb is a live Guild-tier Stripe whole plan, and
// the Raid guild key it moves to.
func upgradable(gb *entitlements.GuildBilling) (string, bool) {
	key, ok := catalog.UpgradeLookupKey(gb.PriceLookupKey)
	return key, ok && gb.Active && !gb.Scoped() && gb.Tier == catalog.TierGuild && gb.StripeSubscriptionID != ""
}

func writeNotUpgradable(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusConflict, "not_upgradable", "only a Guild plan can move to Raid guild", nil)
}

// resolveUpgrade answers every refusal itself: 503 without Stripe, the
// actor refusals, 403 guild_plan_required without a live plan, and 409
// not_upgradable for anything but a Guild-tier Stripe row.
func (s *Service) resolveUpgrade(w http.ResponseWriter, r *http.Request, op string) (upgradeTarget, bool) {
	if s.Gateway == nil {
		writeUnavailable(w, r)
		return upgradeTarget{}, false
	}
	a, ok := s.actingActor(w, r, op, isPayer)
	if !ok {
		return upgradeTarget{}, false
	}
	gb, ok := s.livePlan(w, r, op, a.guildID)
	if !ok {
		return upgradeTarget{}, false
	}
	key, ok := upgradable(gb)
	if !ok {
		writeNotUpgradable(w, r)
		return upgradeTarget{}, false
	}
	priceID, err := s.Gateway.PriceIDForLookupKey(r.Context(), key)
	if err != nil {
		s.fail(w, r, op, err, "could not price the upgrade just now")
		return upgradeTarget{}, false
	}
	return upgradeTarget{actor: a, plan: gb, lookupKey: key, priceID: priceID}, true
}

func (s *Service) upgradePreview(w http.ResponseWriter, r *http.Request) {
	t, ok := s.resolveUpgrade(w, r, "upgrade preview")
	if !ok {
		return
	}
	ctx := r.Context()
	moved, err := s.alreadyOnTarget(ctx, t)
	if err != nil {
		s.upgradeFailed(w, r, "upgrade preview", t, err)
		return
	}
	if moved {
		writeNotUpgradable(w, r)
		return
	}
	p, err := s.Gateway.PreviewPriceChange(ctx, t.plan.StripeSubscriptionID, t.priceID)
	if err != nil {
		s.upgradeFailed(w, r, "upgrade preview", t, err)
		return
	}
	price, _ := catalog.Lookup(t.lookupKey)
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, UpgradePreview{
		AmountDueCents: p.AmountDueCents, NextAmountCents: p.NextAmountCents, Interval: string(price.Interval),
		RenewsAt: *rfc3339(&p.PeriodEnd), ProrationDate: p.ProrationDate.Unix(),
	})
}

// alreadyOnTarget catches an upgrade Stripe applied but our row missed (a
// half-done apply): the subscription is already on the Raid guild price, so
// the row is brought up to date instead of previewing a second move.
func (s *Service) alreadyOnTarget(ctx context.Context, t upgradeTarget) (bool, error) {
	done, err := s.liveOnTarget(ctx, t)
	if err != nil || !done {
		return false, err
	}
	return true, s.upsertSubscription(ctx, t.plan.StripeSubscriptionID, s.upgradeActor(t))
}

// liveOnTarget reports whether Stripe's item already bills the target price.
func (s *Service) liveOnTarget(ctx context.Context, t upgradeTarget) (bool, error) {
	sub, err := s.Gateway.GetSubscription(ctx, t.plan.StripeSubscriptionID)
	if err != nil {
		return false, err
	}
	price := itemPrice(sub)
	return price != nil && price.ID == t.priceID, nil
}

func (s *Service) upgradeActor(t upgradeTarget) string {
	return fmt.Sprintf("billing_upgrade:%d", t.actor.userID)
}

func (s *Service) upgrade(w http.ResponseWriter, r *http.Request) {
	t, ok := s.resolveUpgrade(w, r, "upgrade")
	if !ok {
		return
	}
	var in struct {
		ProrationDate int64 `json:"proration_date"`
	}
	if r.ContentLength != 0 {
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
			return
		}
	}
	prorationDate, fresh := s.freshProrationDate(in.ProrationDate)
	if !fresh {
		httpx.WriteError(w, r, http.StatusConflict, "preview_expired", "the upgrade price has changed; check it again", nil)
		return
	}
	ctx := r.Context()
	err := s.Store.WithGuildLock(ctx, t.actor.guildID, func(ctx context.Context) error {
		return s.changeToRaidGuild(ctx, t, prorationDate)
	})
	if !s.answerChangeError(w, r, t, err) {
		return
	}
	// upsertSubscription takes the guild lock itself, so it runs after release.
	if err := s.upsertSubscription(ctx, t.plan.StripeSubscriptionID, s.upgradeActor(t)); err != nil {
		s.fail(w, r, "upgrade", err, "the upgrade went through; reload to see it")
		return
	}
	read, err := s.buildGuildBillingRead(ctx, t.actor)
	if err != nil {
		s.fail(w, r, "upgrade", err, "the upgrade went through; reload to see it")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, read)
}

// changeToRaidGuild is the upgrade's locked half: the row is re-read, so a
// transfer accepted since resolveUpgrade is never charged on the retiring
// subscription.
func (s *Service) changeToRaidGuild(ctx context.Context, t upgradeTarget, prorationDate time.Time) error {
	gb, err := s.Entitlements.GuildBilling(ctx, t.actor.guildID)
	if err != nil {
		return err
	}
	if gb == nil || gb.StripeSubscriptionID != t.plan.StripeSubscriptionID || gb.PriceLookupKey != t.plan.PriceLookupKey {
		return &checkoutRefusal{status: http.StatusConflict, code: "not_upgradable",
			message: "the plan changed; reload to see it"}
	}
	if _, ok := upgradable(gb); !ok {
		return &checkoutRefusal{status: http.StatusConflict, code: "not_upgradable",
			message: "only a Guild plan can move to Raid guild"}
	}
	// An open transfer Checkout priced its subscription at the Guild key; an
	// upgrade now would be lost when that subscription takes the row.
	if err := s.refuseWhileCheckoutOpen(ctx, t.actor.guildID); err != nil {
		return err
	}
	// A double submit reaches here a second time before the row is upserted;
	// Stripe's own item says whether the move already happened.
	if done, err := s.liveOnTarget(ctx, t); err != nil || done {
		return err
	}
	if err := s.Gateway.ChangePrice(ctx, gb.StripeSubscriptionID, t.priceID, prorationDate); err != nil {
		return priceChangeError{err: err}
	}
	return nil
}

// answerChangeError answers a failed locked half and reports whether the
// upgrade may go on (err was nil).
func (s *Service) answerChangeError(w http.ResponseWriter, r *http.Request, t upgradeTarget, err error) bool {
	var refusal *checkoutRefusal
	var changeErr priceChangeError
	switch {
	case err == nil:
		return true
	case errors.As(err, &refusal):
		httpx.WriteError(w, r, refusal.status, refusal.code, refusal.message, refusal.fields)
	case errors.As(err, &changeErr):
		s.priceChangeFailed(w, r, t, changeErr.err)
	default:
		s.fail(w, r, "upgrade", err, "could not change the plan just now")
	}
	return false
}

// freshProrationDate reads the preview's proration date, false when it is
// missing, older than previewValidFor, or ahead of our clock.
func (s *Service) freshProrationDate(unix int64) (time.Time, bool) {
	if unix <= 0 {
		return time.Time{}, false
	}
	at, now := time.Unix(unix, 0), s.now()
	return at, now.Sub(at) <= previewValidFor && !at.After(now.Add(prorationClockSkew))
}

// upgradeFailed answers a failed preview: a preview never charges, so the
// caller may try again. The log names ids only.
func (s *Service) upgradeFailed(w http.ResponseWriter, r *http.Request, op string, t upgradeTarget, err error) {
	s.logUpgradeError(r, op, t, err)
	httpx.WriteError(w, r, http.StatusConflict, "upgrade_failed", "that did not go through; nothing was charged", nil)
}

// priceChangeFailed answers a failed ChangePrice. Only an error Stripe
// raised before applying anything (a card decline or an invalid request)
// is a clean "nothing was charged" 409; a timeout or a server error may
// follow an update Stripe applied, so it answers 502 and the page reloads.
func (s *Service) priceChangeFailed(w http.ResponseWriter, r *http.Request, t upgradeTarget, err error) {
	s.logUpgradeError(r, "upgrade", t, err)
	if declinedAtStripe(err) {
		httpx.WriteError(w, r, http.StatusConflict, "upgrade_failed", "that did not go through; nothing was charged", nil)
		return
	}
	httpx.WriteError(w, r, http.StatusBadGateway, "upgrade_unconfirmed",
		"we could not confirm the upgrade; check back shortly", nil)
}

// declinedAtStripe is a Stripe card or invalid-request error: refused
// before Stripe changed the subscription.
func declinedAtStripe(err error) bool {
	var stripeErr *stripe.Error
	return errors.As(err, &stripeErr) &&
		(stripeErr.Type == stripe.ErrorTypeCard || stripeErr.Type == stripe.ErrorTypeInvalidRequest)
}

func (s *Service) logUpgradeError(r *http.Request, op string, t upgradeTarget, err error) {
	s.logger().Warn("billing", "id", httpx.RequestIDFrom(r.Context()), "op", op,
		"subscription_id", t.plan.StripeSubscriptionID, "price_id", t.priceID, "err", err)
}
