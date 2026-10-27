// api/internal/billing/checkout.go
package billing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

type checkoutRequest struct {
	Plan       string `json:"plan"`
	Interval   string `json:"interval"`
	GuildID    *int64 `json:"guild_id"`
	Intent     string `json:"intent"`
	ComebackID *int64 `json:"comeback_id"`
}

// intentTransfer is the checkout intent of a guild's billing transfer,
// which only an open plan_transfers offer to the caller unlocks.
const intentTransfer = "transfer"

// transferTrialFloor is the least time ahead a transfer's trial_end may
// be: Stripe's 48-hour minimum plus an hour of slack for clock skew.
const transferTrialFloor = 49 * time.Hour

// Where Stripe returns the buyer, appended to the public base URL.
const (
	checkoutSuccessPath = "/premium/checkout?status=success&session_id={CHECKOUT_SESSION_ID}"
	checkoutCancelPath  = "/premium?canceled=1"
)

// now is the handler's clock; tests set Service.Now.
func (s *Service) now() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}

func (s *Service) checkout(w http.ResponseWriter, r *http.Request) {
	if s.Gateway == nil {
		writeUnavailable(w, r)
		return
	}
	var in checkoutRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	offer, interval, ok := parseCheckoutRequest(in)
	if !ok {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"plan must be premium, guild or raid_guild, interval monthly or yearly", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	userIDStr := strconv.FormatInt(actor.UserID, 10)
	metadata := map[string]string{"plan": offer.EntitlementPlan(), "user_id": userIDStr}

	if offer == catalog.OfferPremium {
		if in.ComebackID != nil {
			metadata["comeback_id"] = strconv.FormatInt(*in.ComebackID, 10)
		}
		s.premiumCheckout(w, r, premiumCheckoutRequest{
			userID: actor.UserID, interval: interval, comebackID: in.ComebackID,
			clientRef: "user:" + userIDStr, metadata: metadata,
		})
		return
	}
	if in.GuildID == nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "guild_id is required for a guild plan", nil)
		return
	}
	metadata["guild_id"] = strconv.FormatInt(*in.GuildID, 10)
	transfer := in.Intent == intentTransfer
	if transfer {
		// Carried onto the Subscription itself (subscription_data.metadata)
		// so the webhook can tell a deliberate billing transfer apart from
		// the double-billing race: the new subscription becomes the one
		// entitlements points at.
		metadata["intent"] = intentTransfer
	}
	s.guildCheckout(w, r, guildCheckoutRequest{
		guildID: *in.GuildID, userID: actor.UserID, transfer: transfer, offer: offer, interval: interval,
		clientRef: fmt.Sprintf("guild:%d:user:%d", *in.GuildID, actor.UserID), metadata: metadata,
	})
}

// parseCheckoutRequest validates the plan and interval, and refuses a
// comeback outside Premium.
func parseCheckoutRequest(in checkoutRequest) (catalog.Offer, catalog.Interval, bool) {
	offer, ok := catalog.ParseOffer(in.Plan)
	if !ok {
		return "", "", false
	}
	if in.Intent != "" && (in.Intent != intentTransfer || offer == catalog.OfferPremium) {
		return "", "", false
	}
	if in.ComebackID != nil {
		// The comeback's own price decides the interval, so none is checked.
		return offer, "", offer == catalog.OfferPremium
	}
	interval := catalog.Interval(in.Interval)
	return offer, interval, interval == catalog.Monthly || interval == catalog.Yearly
}

// checkoutTrial is a trial on a checkout: days from the start, or an end
// instant.
type checkoutTrial struct {
	days int64
	end  *time.Time
}

// createSession resolves the Customer and price and creates the session.
func (s *Service) createSession(ctx context.Context, userID int64, lookupKey, clientRef string,
	metadata map[string]string, trial *checkoutTrial) (CheckoutSession, error) {
	customerID, err := s.customerFor(ctx, userID)
	if err != nil {
		return CheckoutSession{}, err
	}
	priceID, err := s.Gateway.PriceIDForLookupKey(ctx, lookupKey)
	if err != nil {
		return CheckoutSession{}, err
	}
	params := CheckoutParams{
		CustomerID: customerID, PriceID: priceID, ClientReferenceID: clientRef, Metadata: metadata,
		SuccessURL: s.PublicBaseURL + checkoutSuccessPath,
		CancelURL:  s.PublicBaseURL + checkoutCancelPath,
	}
	if trial != nil {
		params.TrialPeriodDays, params.TrialEnd = trial.days, trial.end
	}
	return s.Gateway.CreateCheckoutSession(ctx, params)
}

// writeCheckoutResult answers a checkout path's outcome: the session's
// URL, a refusal's own 4xx, or a logged 500.
func (s *Service) writeCheckoutResult(w http.ResponseWriter, r *http.Request, session CheckoutSession, err error) {
	if err != nil {
		var refusal *checkoutRefusal
		if errors.As(err, &refusal) {
			httpx.WriteError(w, r, refusal.status, refusal.code, refusal.message, refusal.fields)
			return
		}
		s.fail(w, r, "checkout", err, "could not start checkout just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"checkout_url": session.URL})
}

type premiumCheckoutRequest struct {
	userID     int64
	interval   catalog.Interval
	comebackID *int64
	clientRef  string
	metadata   map[string]string
}

// premiumCheckout is the personal path, serialized per user by the same
// lock the webhook's upsert takes, so the "already subscribed" read below
// and the Customer creation cannot race a double click or an in-flight
// webhook. A live premium row (active, trialing or past_due, including
// one set to cancel at the period's end) answers 409 with the portal
// hint: a second subscription would be flagged as a duplicate by
// entitlements.UpsertStripe and set to cancel, and the member would lose
// access when the first one ended despite having paid for the second.
func (s *Service) premiumCheckout(w http.ResponseWriter, r *http.Request, req premiumCheckoutRequest) {
	var result CheckoutSession
	err := s.Store.WithUserLock(r.Context(), req.userID, func(ctx context.Context) error {
		current, err := s.Entitlements.PersonalBilling(ctx, req.userID)
		if err != nil {
			return err
		}
		if current != nil && current.Active {
			return alreadyOnPremiumRefusal()
		}
		lookupKey, err := s.premiumLookupKey(ctx, req)
		if err != nil {
			return err
		}
		session, err := s.createSession(ctx, req.userID, lookupKey, req.clientRef, req.metadata, nil)
		if err != nil {
			return err
		}
		result = session
		return nil
	})
	s.writeCheckoutResult(w, r, result, err)
}

// premiumLookupKey is a comeback's remembered price, else the price on
// offer now. A comeback that is not the caller's, has expired or was used
// answers 404.
func (s *Service) premiumLookupKey(ctx context.Context, req premiumCheckoutRequest) (string, error) {
	if req.comebackID == nil {
		return catalog.CheckoutLookupKey(catalog.OfferPremium, req.interval, s.now()), nil
	}
	key, ok, err := s.Store.Comeback(ctx, *req.comebackID, req.userID, s.now())
	if err != nil {
		return "", err
	}
	if price, known := catalog.Lookup(key); !ok || !known || price.Offer != catalog.OfferPremium {
		return "", &checkoutRefusal{status: http.StatusNotFound, code: "not_found", message: "that offer is not available"}
	}
	return key, nil
}

type guildCheckoutRequest struct {
	guildID   int64
	userID    int64
	transfer  bool
	offer     catalog.Offer
	interval  catalog.Interval
	clientRef string
	metadata  map[string]string
}

// guildCheckout is the guild-plan checkout path, wholly serialized per
// guild by an advisory lock held from the entitlement read through
// Checkout Session creation and the pending_checkouts write (security
// review fix, 2026-09-21): without it, two officers checking out for the
// same guild at once could each read "no active plan," each create a
// session, and each pay.
func (s *Service) guildCheckout(w http.ResponseWriter, r *http.Request, req guildCheckoutRequest) {
	var result CheckoutSession
	err := s.Store.WithGuildLock(r.Context(), req.guildID, func(ctx context.Context) error {
		refusal, err := s.checkGuildCheckout(ctx, req.guildID, req.userID, req.transfer)
		if err != nil {
			return err
		}
		if refusal != nil {
			return refusal
		}
		lookupKey, trial, err := s.guildTerms(ctx, req)
		if err != nil {
			return err
		}
		session, err := s.createSession(ctx, req.userID, lookupKey, req.clientRef, req.metadata, trial)
		if err != nil {
			return err
		}
		if err := s.Store.RecordPendingCheckout(ctx, req.guildID, req.userID, session.SessionID, session.ExpiresAt); err != nil {
			return err
		}
		result = session
		return nil
	})
	s.writeCheckoutResult(w, r, result, err)
}

// guildTerms is the price and trial of a guild checkout. A first purchase
// gets the guild's one 7-day trial and the price on offer now.
func (s *Service) guildTerms(ctx context.Context, req guildCheckoutRequest) (string, *checkoutTrial, error) {
	if req.transfer {
		return s.transferTerms(ctx, req)
	}
	used, err := s.Store.GuildTrialUsed(ctx, req.guildID)
	if err != nil {
		return "", nil, err
	}
	var trial *checkoutTrial
	if !used {
		trial = &checkoutTrial{days: catalog.GuildTrialDays}
	}
	return catalog.CheckoutLookupKey(req.offer, req.interval, s.now()), trial, nil
}

// transferTerms needs the open offer to the caller, keeps the plan's
// current price (so a founder price survives the window's close) and
// trials until the old period ends, floored at transferTrialFloor. The old
// subscription id rides in the metadata for the webhook.
func (s *Service) transferTerms(ctx context.Context, req guildCheckoutRequest) (string, *checkoutTrial, error) {
	fromSubscriptionID, ok, err := s.Store.OpenTransferTo(ctx, req.guildID, req.userID)
	if err != nil {
		return "", nil, err
	}
	current, err := s.Entitlements.GuildBilling(ctx, req.guildID)
	if err != nil {
		return "", nil, err
	}
	if !ok || current == nil || !current.Active || current.Scoped() || current.StripeSubscriptionID != fromSubscriptionID {
		return "", nil, &checkoutRefusal{status: http.StatusConflict, code: "no_transfer_offer",
			message: "there is no billing transfer offered to you for this guild"}
	}
	req.metadata["from_subscription_id"] = fromSubscriptionID
	lookupKey := current.PriceLookupKey
	if lookupKey == "" {
		// A row with no key gets the price on offer for its own tier,
		// never the client's plan.
		tierOffer, _ := catalog.ParseOffer(current.Tier)
		lookupKey = catalog.CheckoutLookupKey(tierOffer, req.interval, s.now())
	}
	trialEnd := s.now().Add(transferTrialFloor)
	if end := current.CurrentPeriodEnd; end != nil && end.After(trialEnd) {
		trialEnd = *end
	}
	return lookupKey, &checkoutTrial{end: &trialEnd}, nil
}
