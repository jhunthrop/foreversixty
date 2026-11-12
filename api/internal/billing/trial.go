// api/internal/billing/trial.go
//
// Free trials told plainly (design/specs/2026-11-08-premium-trial.md): the server decides every
// trial. decideTrial is the one rule the checkout and the chooser's terms both read, so the
// dates the chooser shows are the dates the checkout uses.
package billing

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// The terms' code words.
const (
	codeApplied   = "applied"
	codeUnknown   = "unknown"
	codeTrialUsed = "trial_used"
)

// trialDecision is a checkout's trial: days (0 for none) and what a creator's code did ("" for
// no code).
type trialDecision struct {
	days int64
	code string
}

// decideTrial answers the trial a plan with defaultDays gets for a buyer whose trial is spent or
// not, with an optional normalized creator code. A code that is not live leaves the default; a
// spent trial has none, code or not.
func (s *Service) decideTrial(ctx context.Context, code string, spent bool, defaultDays int64) (trialDecision, error) {
	base := trialDecision{days: defaultDays}
	if spent {
		base.days = 0
	}
	if code == "" {
		return base, nil
	}
	live := false
	if s.Creators != nil {
		var err error
		if live, err = s.Creators.CodeIsLive(ctx, code); err != nil {
			return trialDecision{}, err
		}
	}
	switch {
	case !live:
		base.code = codeUnknown
		return base, nil
	case spent:
		return trialDecision{code: codeTrialUsed}, nil
	}
	return trialDecision{days: s.Creators.CodeTrialDays(), code: codeApplied}, nil
}

type termsRequest struct {
	Plan        string `json:"plan"`
	Interval    string `json:"interval"`
	GuildID     *int64 `json:"guild_id"`
	CreatorCode string `json:"creator_code"`
}

// TermsView is POST /v1/billing/checkout/terms: what a checkout started now would sell.
type TermsView struct {
	TrialDays  int64   `json:"trial_days"`
	EndsAt     *string `json:"ends_at"`
	PriceCents int64   `json:"price_cents"`
	Interval   string  `json:"interval"`
	Founder    bool    `json:"founder"`
	Code       *string `json:"code"`
}

// terms answers the chooser (spec 3.3). It reads only; the checkout re-checks everything.
func (s *Service) terms(w http.ResponseWriter, r *http.Request) {
	var in termsRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBillingBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return
	}
	offer, ok := catalog.ParseOffer(in.Plan)
	interval := catalog.Interval(in.Interval)
	if !ok || (interval != catalog.Monthly && interval != catalog.Yearly) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"plan must be premium, guild or raid_guild, interval monthly or yearly", nil)
		return
	}
	if offer != catalog.OfferPremium && (in.GuildID == nil || *in.GuildID <= 0) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "guild_id is required for a guild plan", nil)
		return
	}
	view, err := s.termsFor(r.Context(), offer, interval, auth.ActorFrom(r.Context()).UserID, in.GuildID, creatorsNormalize(in.CreatorCode))
	if err != nil {
		s.fail(w, r, "terms", err, "could not read the terms just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// termsFor is the terms of one plan for the caller (a guild plan: for that guild).
func (s *Service) termsFor(ctx context.Context, offer catalog.Offer, interval catalog.Interval, userID int64,
	guildID *int64, code string) (TermsView, error) {
	now := s.now()
	price, _ := catalog.Lookup(catalog.CheckoutLookupKey(offer, interval, now))
	spent, defaultDays, err := s.trialSpent(ctx, offer, userID, guildID)
	if err != nil {
		return TermsView{}, err
	}
	d, err := s.decideTrial(ctx, code, spent, defaultDays)
	if err != nil {
		return TermsView{}, err
	}
	view := TermsView{TrialDays: d.days, PriceCents: price.AmountCents, Interval: string(interval), Founder: price.Founder}
	if d.days > 0 {
		end := now.Add(time.Duration(d.days) * 24 * time.Hour)
		view.EndsAt = rfc3339(&end)
	}
	if d.code != "" {
		view.Code = &d.code
	}
	return view, nil
}

// trialSpent is whether the buyer's (Premium) or the guild's trial is gone, and the plan's
// default trial length.
func (s *Service) trialSpent(ctx context.Context, offer catalog.Offer, userID int64, guildID *int64) (bool, int64, error) {
	if offer == catalog.OfferPremium {
		spent, err := s.Store.PremiumTrialSpent(ctx, userID)
		return spent, catalog.PremiumTrialDays, err
	}
	used, err := s.Store.GuildTrialUsed(ctx, *guildID)
	return used, catalog.GuildTrialDays, err
}
