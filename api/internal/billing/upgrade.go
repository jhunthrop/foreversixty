// api/internal/billing/upgrade.go
package billing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// previewValidFor is how long an upgrade preview's proration date stays good
// for the upgrade itself; past it the amounts shown may no longer be what
// Stripe charges.
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
	key, upgradable := catalog.UpgradeLookupKey(gb.PriceLookupKey)
	if gb.Tier != catalog.TierGuild || gb.StripeSubscriptionID == "" || !upgradable {
		httpx.WriteError(w, r, http.StatusConflict, "not_upgradable", "only a Guild plan can move to Raid guild", nil)
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
	p, err := s.Gateway.PreviewPriceChange(r.Context(), t.plan.StripeSubscriptionID, t.priceID)
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

func (s *Service) upgrade(w http.ResponseWriter, r *http.Request) {
	var in struct {
		ProrationDate int64 `json:"proration_date"`
	}
	if r.ContentLength != 0 {
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
			return
		}
	}
	t, ok := s.resolveUpgrade(w, r, "upgrade")
	if !ok {
		return
	}
	prorationDate, fresh := s.freshProrationDate(in.ProrationDate)
	if !fresh {
		httpx.WriteError(w, r, http.StatusConflict, "preview_expired", "the upgrade price has changed; check it again", nil)
		return
	}
	ctx := r.Context()
	if err := s.Gateway.ChangePrice(ctx, t.plan.StripeSubscriptionID, t.priceID, prorationDate); err != nil {
		s.upgradeFailed(w, r, "upgrade", t, err)
		return
	}
	actor := fmt.Sprintf("billing_upgrade:%d", t.actor.userID)
	if err := s.upsertSubscription(ctx, t.plan.StripeSubscriptionID, actor); err != nil {
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

// freshProrationDate reads the preview's proration date, false when it is
// missing, older than previewValidFor, or ahead of our clock.
func (s *Service) freshProrationDate(unix int64) (time.Time, bool) {
	if unix <= 0 {
		return time.Time{}, false
	}
	at, now := time.Unix(unix, 0), s.now()
	return at, now.Sub(at) <= previewValidFor && !at.After(now.Add(prorationClockSkew))
}

// upgradeFailed answers a refused price change: nothing changed, so the
// caller may try again. The log names ids only.
func (s *Service) upgradeFailed(w http.ResponseWriter, r *http.Request, op string, t upgradeTarget, err error) {
	s.logger().Warn("billing", "id", httpx.RequestIDFrom(r.Context()), "op", op,
		"subscription_id", t.plan.StripeSubscriptionID, "price_id", t.priceID, "err", err)
	httpx.WriteError(w, r, http.StatusConflict, "upgrade_failed", "that did not go through; nothing was charged", nil)
}
