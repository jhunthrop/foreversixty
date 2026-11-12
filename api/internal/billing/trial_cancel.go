// api/internal/billing/trial_cancel.go
//
// Cancelling a free trial (spec 2 and 4.3): one click and one confirm, from the account page,
// guild settings, or the reminder's link. Every path sets cancel_at_period_end on the one
// trialing subscription, so the plan stays on to the trial end and Stripe never charges; keeping
// the plan clears it. Opening a link never cancels anything: only its POST does.
package billing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	statusTrialing = "trialing"
	// sourceStripe is the entitlement source of a row Stripe drives.
	sourceStripe = "stripe"
	// cancelLinkPath is the site page a reminder's link opens; the token rides in its query and
	// from there only in POST bodies to the API.
	cancelLinkPath = "/premium/cancel?t="
	// linkExpiredMessage is the page's own sentence for a dead link.
	linkExpiredMessage = "This link has expired or was already used. You can cancel from your account page."
)

// mountTrial registers the trial routes.
func mountTrial(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	read := httpx.RateLimitPer(termsRateLimitPerHour, time.Hour, trustedProxyHops)
	write := httpx.RateLimitPer(billingRateLimitPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/billing/checkout/terms", read(auth.RequireSession(s.terms)))
	mux.Handle("POST /v1/billing/trial/cancel", write(auth.RequireBrowserSession(s.trialCancel)))
	mux.Handle("POST /v1/billing/trial/keep", write(auth.RequireBrowserSession(s.trialKeep)))
	mux.Handle("POST /v1/billing/cancel-link/preview", read(http.HandlerFunc(s.cancelLinkPreview)))
	mux.Handle("POST /v1/billing/cancel-link", write(http.HandlerFunc(s.cancelLink)))
}

// TrialChangeView answers a trial cancel or keep.
type TrialChangeView struct {
	EndsAt            string `json:"ends_at"`
	CancelAtPeriodEnd bool   `json:"cancel_at_period_end"`
}

// CancelLinkView is what a cancel link names: the plan word (premium, guild or raid_guild), the
// guild for a guild plan, and the trial's end and price.
type CancelLinkView struct {
	Plan     string  `json:"plan"`
	Guild    *string `json:"guild"`
	EndsAt   string  `json:"ends_at"`
	Canceled bool    `json:"canceled"`
	catalog.PriceFacts
}

type trialRequest struct {
	GuildID *int64 `json:"guild_id"`
}

type tokenRequest struct {
	Token string `json:"token"`
}

// trialRow is the caller's trialing subscription.
type trialRow struct {
	subscriptionID string
	trialEnd       time.Time
}

var (
	errNotTrialing = &checkoutRefusal{status: http.StatusConflict, code: "not_trialing", message: "this plan is not in a free trial"}
	errNotPayerRow = &checkoutRefusal{status: http.StatusForbidden, code: "forbidden", message: "only the guild's payer can do that"}
)

func (s *Service) trialCancel(w http.ResponseWriter, r *http.Request) {
	s.changeTrial(w, r, "trial_cancel", true)
}

func (s *Service) trialKeep(w http.ResponseWriter, r *http.Request) {
	s.changeTrial(w, r, "trial_keep", false)
}

// changeTrial sets (cancel) or clears (keep) cancel_at_period_end on the caller's trialing
// Premium, or on the guild's trialing plan when the caller pays for it.
func (s *Service) changeTrial(w http.ResponseWriter, r *http.Request, op string, cancel bool) {
	if s.Gateway == nil {
		writeUnavailable(w, r)
		return
	}
	var in trialRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	userID := auth.ActorFrom(r.Context()).UserID
	row, err := s.trialRowOf(r.Context(), userID, in.GuildID)
	if err == nil {
		err = s.setTrialCancel(r.Context(), row.subscriptionID, cancel, op+":user:"+strconv.FormatInt(userID, 10))
	}
	var refusal *checkoutRefusal
	switch {
	case errors.As(err, &refusal):
		httpx.WriteError(w, r, refusal.status, refusal.code, refusal.message, nil)
	case err != nil:
		s.fail(w, r, op, err, "could not change the trial just now")
	default:
		httpx.WriteOK(w, r, http.StatusOK, TrialChangeView{EndsAt: row.trialEnd.UTC().Format(time.RFC3339), CancelAtPeriodEnd: cancel})
	}
}

// trialRowOf is the caller's own trialing Premium, or the guild's trialing plan the caller pays.
func (s *Service) trialRowOf(ctx context.Context, userID int64, guildID *int64) (trialRow, error) {
	if guildID == nil {
		b, err := s.Entitlements.PersonalBilling(ctx, userID)
		if err != nil {
			return trialRow{}, err
		}
		if b == nil || b.Source != sourceStripe || b.Status != statusTrialing || b.CurrentPeriodEnd == nil {
			return trialRow{}, errNotTrialing
		}
		return trialRow{subscriptionID: b.StripeSubscriptionID, trialEnd: *b.CurrentPeriodEnd}, nil
	}
	gb, err := s.Entitlements.GuildBilling(ctx, *guildID)
	if err != nil {
		return trialRow{}, err
	}
	if gb == nil || gb.BillingUserID == nil || *gb.BillingUserID != userID {
		return trialRow{}, errNotPayerRow
	}
	if gb.StripeSubscriptionID == "" || gb.Status != statusTrialing || gb.TrialEnd == nil {
		return trialRow{}, errNotTrialing
	}
	return trialRow{subscriptionID: gb.StripeSubscriptionID, trialEnd: *gb.TrialEnd}, nil
}

// setTrialCancel tells Stripe, then re-reads the subscription into the entitlement so the account
// page shows the change at once instead of waiting for the webhook. A failed re-read is logged:
// the webhook writes the same state.
func (s *Service) setTrialCancel(ctx context.Context, subscriptionID string, cancel bool, actor string) error {
	var err error
	if cancel {
		err = s.Gateway.CancelSubscriptionAtPeriodEnd(ctx, subscriptionID)
	} else {
		err = s.Gateway.ResumeSubscription(ctx, subscriptionID)
	}
	if err != nil {
		return err
	}
	if err := s.upsertSubscription(ctx, subscriptionID, actor); err != nil {
		s.logger().Error("billing", "op", "trial_resync", "subscription_id", subscriptionID, "err", err)
	}
	return nil
}

// readToken decodes a cancel link body; ok is false (and the 410 written) for a token of the
// wrong shape, which no stored hash could match.
func readToken(w http.ResponseWriter, r *http.Request) (string, bool) {
	var in tokenRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return "", false
	}
	if raw, err := decodeToken(in.Token); err != nil || len(raw) != cancelTokenBytes {
		writeLinkExpired(w, r)
		return "", false
	}
	return in.Token, true
}

func writeLinkExpired(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusGone, "link_expired", linkExpiredMessage, nil)
}

// cancelLinkPreview names what a link would cancel. It changes nothing.
func (s *Service) cancelLinkPreview(w http.ResponseWriter, r *http.Request) {
	token, ok := readToken(w, r)
	if !ok {
		return
	}
	t, ok, err := s.Store.CancelTokenTarget(r.Context(), token, s.now())
	switch {
	case err != nil:
		s.fail(w, r, "cancel_link_preview", err, "could not read the link just now")
	case !ok:
		writeLinkExpired(w, r)
	default:
		httpx.WriteOK(w, r, http.StatusOK, cancelLinkView(t))
	}
}

// cancelLinkView is a target as the page shows it.
func cancelLinkView(t CancelTarget) CancelLinkView {
	v := CancelLinkView{Plan: t.Plan, EndsAt: t.TrialEnd.UTC().Format(time.RFC3339), Canceled: t.CancelAtPeriodEnd,
		PriceFacts: catalog.FactsOf(t.PriceLookupKey, t.UnitAmountCents)}
	if t.Plan == entitlements.PlanGuild {
		v.Plan = t.Tier
		guild := t.Guild
		v.Guild = &guild
	}
	return v
}

// cancelLink cancels the one subscription the link names at its period end, once.
func (s *Service) cancelLink(w http.ResponseWriter, r *http.Request) {
	if s.Gateway == nil {
		writeUnavailable(w, r)
		return
	}
	token, ok := readToken(w, r)
	if !ok {
		return
	}
	ctx, now := r.Context(), s.now()
	t, ok, err := s.Store.CancelTokenTarget(ctx, token, now)
	if err == nil && ok {
		ok, err = s.Store.UseCancelToken(ctx, token, now)
	}
	switch {
	case err != nil:
		s.fail(w, r, "cancel_link", err, "could not cancel just now")
		return
	case !ok:
		writeLinkExpired(w, r)
		return
	}
	if err := s.setTrialCancel(ctx, t.StripeSubscriptionID, true, "cancel_link:entitlement:"+strconv.FormatInt(t.EntitlementID, 10)); err != nil {
		if relErr := s.Store.ReleaseCancelToken(ctx, token); relErr != nil {
			err = errors.Join(err, relErr)
		}
		s.fail(w, r, "cancel_link", err, "could not cancel just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, TrialChangeView{EndsAt: t.TrialEnd.UTC().Format(time.RFC3339), CancelAtPeriodEnd: true})
}

// cancelLinkURL is the reminder's link for a token.
func (s *Service) cancelLinkURL(token string) string { return s.siteURL(cancelLinkPath + token) }
