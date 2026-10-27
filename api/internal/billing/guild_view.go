// api/internal/billing/guild_view.go
package billing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/seats"
)

// guildPlanScope values, as the web reads them.
const (
	scopeWholePlan    = "plan"
	scopeOfficerTools = "officer_tools"
	statusCanceled    = "canceled"
)

// GuildBillingRead is GET /v1/guilds/{id}/billing. TrialUsed is top level
// because Billing is null before the guild has a plan.
type GuildBillingRead struct {
	Billing            *GuildPlanView  `json:"billing"`
	TrialUsed          bool            `json:"trial_used"`
	VerifiedCount      int             `json:"verified_count"`
	Seats              []seats.Seat    `json:"seats"`
	Transfer           *TransferView   `json:"transfer"`
	TransferCandidates []seats.Officer `json:"transfer_candidates"`
	Can                []string        `json:"can"`
}

// GuildPlanView is the guild's plan row, ended or live.
type GuildPlanView struct {
	Scope             string  `json:"scope"`
	Tier              string  `json:"tier"`
	Status            string  `json:"status"`
	CurrentPeriodEnd  *string `json:"current_period_end"`
	CancelAtPeriodEnd bool    `json:"cancel_at_period_end"`
	TrialEnd          *string `json:"trial_end"`
	catalog.PriceFacts
	// BilledBy is empty for a grant: there is no billing contact to name.
	BilledBy             string `json:"billed_by"`
	YouAreBillingContact bool   `json:"you_are_billing_contact"`
}

// TransferView is the guild's plan transfer offer, as the reader sees it.
type TransferView struct {
	ToName   string  `json:"to_name"`
	ToYou    bool    `json:"to_you"`
	FromName string  `json:"from_name"`
	At       string  `json:"at"`
	State    string  `json:"state"`
	StartsAt *string `json:"starts_at"`
	catalog.PriceFacts
}

func (s *Service) guildBilling(w http.ResponseWriter, r *http.Request) {
	a, ok := s.guildActorFor(w, r, "guild billing")
	if !ok {
		return
	}
	read, err := s.buildGuildBillingRead(r.Context(), a)
	if err != nil {
		s.fail(w, r, "guild billing", err, "could not read billing just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, read)
}

// buildGuildBillingRead sweeps the guild's seats, then reads every fact the
// settings page draws.
func (s *Service) buildGuildBillingRead(ctx context.Context, a guildActor) (GuildBillingRead, error) {
	now := s.now()
	if err := s.Seats.SweepGuild(ctx, a.guildID, now); err != nil {
		return GuildBillingRead{}, err
	}
	gb, err := s.Entitlements.GuildBilling(ctx, a.guildID)
	if err != nil {
		return GuildBillingRead{}, err
	}
	transferable := transferablePlan(gb)
	read := GuildBillingRead{Can: canFor(a, transferable), TransferCandidates: []seats.Officer{}}
	if gb != nil {
		if read.Billing, err = s.planView(ctx, gb, a.userID); err != nil {
			return GuildBillingRead{}, err
		}
	}
	if read.TrialUsed, err = s.Store.GuildTrialUsed(ctx, a.guildID); err != nil {
		return GuildBillingRead{}, err
	}
	if read.VerifiedCount, err = s.Store.VerifiedCount(ctx, a.guildID); err != nil {
		return GuildBillingRead{}, err
	}
	if read.Seats, err = s.Seats.List(ctx, a.guildID, now); err != nil {
		return GuildBillingRead{}, err
	}
	if read.Transfer, err = s.transferView(ctx, a, gb); err != nil {
		return GuildBillingRead{}, err
	}
	if transferable && a.subscriber && !a.frozen {
		if read.TransferCandidates, err = s.transferCandidates(ctx, a); err != nil {
			return GuildBillingRead{}, err
		}
	}
	return read, nil
}

// canFor is what the reader may do: a payer manages billing and seats, the
// subscriber of a live Stripe plan also transfers it. A frozen claimant does
// nothing.
func canFor(a guildActor, transferable bool) []string {
	can := []string{}
	if a.frozen {
		return can
	}
	if a.payer {
		can = append(can, canBilling, canSeats)
	}
	if a.subscriber && transferable {
		can = append(can, canTransfer)
	}
	return can
}

func (s *Service) planView(ctx context.Context, gb *entitlements.GuildBilling, viewerID int64) (*GuildPlanView, error) {
	v := &GuildPlanView{
		Scope: scopeWholePlan, Tier: gb.Tier, Status: gb.Status,
		CurrentPeriodEnd: rfc3339(gb.CurrentPeriodEnd), CancelAtPeriodEnd: gb.CancelAtPeriodEnd,
		TrialEnd: rfc3339(gb.TrialEnd), PriceFacts: catalog.FactsOf(gb.PriceLookupKey, gb.UnitAmountCents),
	}
	if gb.Scoped() {
		v.Scope = scopeOfficerTools
	}
	// An expired grant keeps status 'active' in its row; it reads as ended.
	if !gb.Active && entitlements.StatusGrantsAccess(gb.Status) {
		v.Status = statusCanceled
	}
	if gb.BillingUserID != nil {
		u, err := s.Users.User(ctx, *gb.BillingUserID)
		if err != nil {
			return nil, fmt.Errorf("billing: billing user of guild row %d: %w", gb.ID, err)
		}
		v.BilledBy = u.PublicName()
		v.YouAreBillingContact = *gb.BillingUserID == viewerID
	}
	return v, nil
}

// transferCandidates is every verified officer the subscriber could hand the
// plan to: everyone but themselves.
func (s *Service) transferCandidates(ctx context.Context, a guildActor) ([]seats.Officer, error) {
	officers, err := s.Seats.Officers(ctx, a.guildID)
	if err != nil {
		return nil, err
	}
	out := []seats.Officer{}
	for _, o := range officers {
		if o.UserID == a.userID {
			continue
		}
		if o.Name, err = s.displayName(ctx, a.guildID, o.UserID, o.Name); err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return out, nil
}

// displayName is the account's character name in the guild, or its public
// name when it has no character to show.
func (s *Service) displayName(ctx context.Context, guildID, userID int64, known string) (string, error) {
	if known != "" {
		return known, nil
	}
	name, err := s.Seats.MemberName(ctx, guildID, userID)
	if err != nil || name != "" {
		return name, err
	}
	u, err := s.Users.User(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("billing: name of user %d: %w", userID, err)
	}
	return u.PublicName(), nil
}
