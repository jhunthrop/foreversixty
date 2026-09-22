// api/internal/billing/handler.go
package billing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// maxBillingBody bounds a checkout/portal request body: both are a
// handful of fields, a kilobyte is generous.
const maxBillingBody = 1 << 10

// billingRateLimitPerHour is POST /v1/billing/checkout and
// POST /v1/billing/portal's own per-IP budget (spec §3): generous for a
// signed-in member retrying a declined card, bounding a script that
// would otherwise mint many Stripe sessions for free.
const billingRateLimitPerHour = 20

// retentionGraceDays is the retention grace's length past a
// subscription's last known period end (spec §2.7).
const retentionGraceDays = 30 * 24 * time.Hour

// GuildRanker is the part of auth.Store this package needs: whether
// userID is a verified officer/leader of guildID (spec §2.3's
// VerifiedGuildRank, §0). *auth.Store satisfies it.
type GuildRanker interface {
	GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error)
}

// UserReader is the part of auth.Store this package needs to create a
// Stripe Customer with the account's email, when it has one (spec §2.4).
type UserReader interface {
	User(ctx context.Context, id int64) (auth.User, error)
}

// FrozenClaimant is guilds.Store's freeze hook (Ruling A/B): whether
// userID is guildID's disputed claimant right now. *guilds.Store
// satisfies it; nil (as every test that does not exercise a contested
// claim leaves it) makes the check a no-op, exactly like reports.Service.Guilds.
type FrozenClaimant interface {
	FrozenClaimant(ctx context.Context, guildID, userID int64) (bool, error)
}

// Entitlements is the part of entitlements.Store this package needs.
// *entitlements.Store satisfies it.
type Entitlements interface {
	GuildBilling(ctx context.Context, guildID int64) (*entitlements.GuildBilling, error)
	UpsertStripe(ctx context.Context, p entitlements.StripeUpsert) (entitlements.UpsertResult, error)
	SetCanceled(ctx context.Context, stripeSubscriptionID string, graceUntil *time.Time, actor string) error
	// StripeSubscriptionIDs is stripe-reconcile's database→Stripe read:
	// every stripe_subscription_id whose row still needs to be believed
	// live (spec §2.8).
	StripeSubscriptionIDs(ctx context.Context) ([]string, error)
	// RecordAnomaly is stripe-reconcile's Stripe→database read's writer:
	// a live Stripe subscription for one of our products with no
	// matching entitlements row at all (spec §2.8's last paragraph,
	// security review fix 2026-09-21).
	RecordAnomaly(ctx context.Context, userID, guildID *int64, plan string,
		kind entitlements.AnomalyKind, stripeSubscriptionID, actor string) error
}

// Service serves every billing route.
type Service struct {
	Store        *Store
	Entitlements Entitlements
	Accounts     GuildRanker
	Users        UserReader
	Guilds       FrozenClaimant
	// Gateway is nil when Stripe is not configured on this deployment
	// (no STRIPE_SECRET_KEY) — every route answers 503 rather than the
	// API failing to start (coordinator's lane constraint).
	Gateway       Gateway
	PublicBaseURL string
	WebhookSecret string
	Log           *slog.Logger
}

// Mount registers every billing route unconditionally: with Gateway nil
// they all answer 503 billing_unavailable rather than 404, so a
// deployment with no Stripe keys yet still describes the surface it will
// have once they are set.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	limit := httpx.RateLimitPer(billingRateLimitPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/billing/checkout", limit(auth.RequireSession(s.checkout)))
	mux.Handle("POST /v1/billing/portal", limit(auth.RequireSession(s.portal)))
	mux.HandleFunc("POST /v1/billing/webhook", s.webhook)
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("billing", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

func writeUnavailable(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusServiceUnavailable, "billing_unavailable",
		"billing is not configured on this deployment", nil)
}

type checkoutRequest struct {
	Plan     string `json:"plan"`
	Interval string `json:"interval"`
	GuildID  *int64 `json:"guild_id"`
	Intent   string `json:"intent"`
}

// lookupKeyFor maps a (plan, interval) pair to spec §2.1's fixed lookup
// key, the only place that mapping exists.
func lookupKeyFor(plan, interval string) (string, bool) {
	switch {
	case plan == entitlements.PlanPremium && interval == "monthly":
		return LookupKeyPremiumMonthly, true
	case plan == entitlements.PlanPremium && interval == "yearly":
		return LookupKeyPremiumYearly, true
	case plan == entitlements.PlanGuild && interval == "monthly":
		return LookupKeyGuildMonthly, true
	case plan == entitlements.PlanGuild && interval == "yearly":
		return LookupKeyGuildYearly, true
	default:
		return "", false
	}
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
	lookupKey, ok := lookupKeyFor(in.Plan, in.Interval)
	if !ok {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"plan must be premium or guild, interval monthly or yearly", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	userIDStr := strconv.FormatInt(actor.UserID, 10)
	metadata := map[string]string{"plan": in.Plan, "user_id": userIDStr}
	clientRef := "user:" + userIDStr

	if in.Plan == entitlements.PlanGuild {
		if in.GuildID == nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "guild_id is required for the guild plan", nil)
			return
		}
		metadata["guild_id"] = strconv.FormatInt(*in.GuildID, 10)
		clientRef = fmt.Sprintf("guild:%d:user:%d", *in.GuildID, actor.UserID)
		transfer := in.Intent == "transfer"
		if transfer {
			// Carried onto the Subscription itself (subscription_data.metadata,
			// same as every other field here) so upsertFromSubscription can
			// tell a deliberate "Take over billing" handoff (spec §2.9
			// RULING 10) apart from the double-billing security finding's
			// race: RULING 10 intends the new subscription to become the
			// one entitlements points at, with the old one left for the
			// guild to cancel by hand — the opposite of the duplicate guard
			// UpsertStripe otherwise applies here.
			metadata["intent"] = "transfer"
		}
		s.guildCheckout(w, r, *in.GuildID, actor.UserID, transfer, lookupKey, clientRef, metadata)
		return
	}

	customerID, err := s.customerFor(r.Context(), actor.UserID)
	if err != nil {
		s.fail(w, r, "checkout", err, "could not start checkout just now")
		return
	}
	priceID, err := s.Gateway.PriceIDForLookupKey(r.Context(), lookupKey)
	if err != nil {
		s.fail(w, r, "checkout", err, "could not start checkout just now")
		return
	}
	session, err := s.Gateway.CreateCheckoutSession(r.Context(), CheckoutParams{
		CustomerID: customerID, PriceID: priceID, ClientReferenceID: clientRef, Metadata: metadata,
		SuccessURL: s.PublicBaseURL + "/premium/checkout?status=success&session_id={CHECKOUT_SESSION_ID}",
		CancelURL:  s.PublicBaseURL + "/premium?canceled=1",
	})
	if err != nil {
		s.fail(w, r, "checkout", err, "could not start checkout just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"checkout_url": session.URL})
}

// checkoutRefusal carries a refusal's HTTP shape out of guildCheckout's
// WithGuildLock closure without conflating "the caller may not do this"
// with an actual failure (s.fail's 500 path) — errors.As in guildCheckout
// tells the two apart.
type checkoutRefusal struct {
	status  int
	code    string
	message string
	fields  map[string]string
}

func (e *checkoutRefusal) Error() string { return e.message }

// guildCheckout is the guild-plan checkout path, wholly serialized per
// guild by an advisory lock held from the entitlement read through
// Checkout Session creation and the pending_checkouts write (security
// review fix, 2026-09-21): without it, two officers checking out for the
// same guild at once could each read "no active plan," each create a
// session, and each pay — the second webhook's upsert would then silently
// overwrite the first subscription id (entitlements.UpsertStripe's
// duplicate guard is the second, independent layer against that half of
// the race, for a delivery this lock's window does not cover).
func (s *Service) guildCheckout(w http.ResponseWriter, r *http.Request, guildID, userID int64,
	transfer bool, lookupKey, clientRef string, metadata map[string]string) {
	var result CheckoutSession
	err := s.Store.WithGuildLock(r.Context(), guildID, func(ctx context.Context) error {
		refusal, err := s.checkGuildCheckout(ctx, guildID, userID, transfer)
		if err != nil {
			return err
		}
		if refusal != nil {
			return refusal
		}
		customerID, err := s.customerFor(ctx, userID)
		if err != nil {
			return err
		}
		priceID, err := s.Gateway.PriceIDForLookupKey(ctx, lookupKey)
		if err != nil {
			return err
		}
		session, err := s.Gateway.CreateCheckoutSession(ctx, CheckoutParams{
			CustomerID: customerID, PriceID: priceID, ClientReferenceID: clientRef, Metadata: metadata,
			SuccessURL: s.PublicBaseURL + "/premium/checkout?status=success&session_id={CHECKOUT_SESSION_ID}",
			CancelURL:  s.PublicBaseURL + "/premium?canceled=1",
		})
		if err != nil {
			return err
		}
		if err := s.Store.RecordPendingCheckout(ctx, guildID, userID, session.SessionID, session.ExpiresAt); err != nil {
			return err
		}
		result = session
		return nil
	})
	if err != nil {
		var refusal *checkoutRefusal
		if errors.As(err, &refusal) {
			httpx.WriteError(w, r, refusal.status, refusal.code, refusal.message, refusal.fields)
			return
		}
		s.fail(w, r, "checkout", err, "could not start checkout just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"checkout_url": result.URL})
}

// alreadyOnThePlanRefusal is the 409/portal_hint shape both an already-
// active entitlement and a pending, in-flight checkout answer with —
// from the caller's own point of view the two are indistinguishable
// ("someone has already started buying this guild the plan").
func alreadyOnThePlanRefusal() *checkoutRefusal {
	return &checkoutRefusal{status: http.StatusConflict, code: "conflict",
		message: "this guild already has the plan", fields: map[string]string{"portal_hint": "1"}}
}

// checkGuildCheckout is spec §2.3's three preconditions, Ruling B's
// frozen-claimant check, and (security review fix, 2026-09-21) the
// pending-checkout guard. A non-nil refusal is a caller-facing 4xx
// guildCheckout should answer with; a non-nil error is a genuine internal
// failure (a DB read that errored) which the caller instead routes to
// s.fail — logged with the real underlying error, never silently turned
// into an empty-message 500. Always called from inside WithGuildLock, so
// the entitlement and pending-checkout reads below are race-free against
// a second, concurrent caller for the same guild.
func (s *Service) checkGuildCheckout(ctx context.Context, guildID, userID int64, transfer bool) (*checkoutRefusal, error) {
	claimed, err := s.Store.GuildClaimed(ctx, guildID)
	if err != nil {
		return nil, err
	}
	if !claimed {
		return &checkoutRefusal{status: http.StatusForbidden, code: "forbidden", message: "this guild has not been claimed yet"}, nil
	}
	rank, verified, err := s.Accounts.GuildRank(ctx, guildID, userID)
	if err != nil {
		return nil, err
	}
	if !verified || (rank != "officer" && rank != "leader") {
		return &checkoutRefusal{status: http.StatusForbidden, code: "forbidden",
			message: "you must be a verified officer or leader of this guild"}, nil
	}
	if s.Guilds != nil {
		frozen, err := s.Guilds.FrozenClaimant(ctx, guildID, userID)
		if err != nil {
			return nil, err
		}
		if frozen {
			return &checkoutRefusal{status: http.StatusForbidden, code: "forbidden",
				message: "this guild's claim is contested; billing actions are frozen for the disputed claimant until a moderator resolves it"}, nil
		}
	}
	existing, err := s.Entitlements.GuildBilling(ctx, guildID)
	if err != nil {
		return nil, err
	}
	active := existing != nil && (existing.Status == "active" || existing.Status == "trialing" || existing.Status == "past_due")
	if active && !transfer {
		return alreadyOnThePlanRefusal(), nil
	}
	pending, err := s.Store.PendingGuildCheckout(ctx, guildID)
	if err != nil {
		return nil, err
	}
	if pending {
		return alreadyOnThePlanRefusal(), nil
	}
	return nil, nil
}

// customerFor returns userID's Stripe Customer id, creating and saving
// one on first use (spec §2.4). The mapping is written before any
// caller uses the id for anything else, so a crash between the two never
// leaves an orphaned, unmapped Customer visible to this codebase.
func (s *Service) customerFor(ctx context.Context, userID int64) (string, error) {
	if id, ok, err := s.Store.CustomerID(ctx, userID); err != nil {
		return "", err
	} else if ok {
		return id, nil
	}
	u, err := s.Users.User(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("billing: customer for %d: %w", userID, err)
	}
	email := ""
	if u.Email != nil {
		email = *u.Email
	}
	id, err := s.Gateway.CreateCustomer(ctx, email)
	if err != nil {
		return "", err
	}
	if err := s.Store.SaveCustomerID(ctx, userID, id); err != nil {
		return "", err
	}
	return id, nil
}

type portalRequest struct {
	GuildID *int64 `json:"guild_id"`
}

func (s *Service) portal(w http.ResponseWriter, r *http.Request) {
	if s.Gateway == nil {
		writeUnavailable(w, r)
		return
	}
	var in portalRequest
	if r.ContentLength != 0 {
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
			return
		}
	}
	actor := auth.ActorFrom(r.Context())
	billingUserID := actor.UserID
	if in.GuildID != nil {
		gb, err := s.Entitlements.GuildBilling(r.Context(), *in.GuildID)
		if err != nil {
			s.fail(w, r, "portal", err, "could not open billing just now")
			return
		}
		if gb == nil || gb.BillingUserID == nil || *gb.BillingUserID != actor.UserID {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found",
				"you are not this guild's billing contact", nil)
			return
		}
	}
	customerID, ok, err := s.Store.CustomerID(r.Context(), billingUserID)
	if err != nil {
		s.fail(w, r, "portal", err, "could not open billing just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no billing account on file yet", nil)
		return
	}
	url, err := s.Gateway.CreatePortalSession(r.Context(), customerID, s.PublicBaseURL+"/account")
	if err != nil {
		s.fail(w, r, "portal", err, "could not open billing just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"portal_url": url})
}
