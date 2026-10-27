package auth

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

const (
	guildPlanScopeWhole        = "plan"
	guildPlanScopeOfficerTools = "officer_tools"
	// billingStatusCanceled is the ended-row status word BillingView shows.
	billingStatusCanceled = "canceled"
	// resubscribePath is where an ended seat's comeback link goes.
	resubscribePath = "/premium/checkout"
	// plansPath is where an ended seat whose price is unknown links to.
	plansPath = "/premium"
)

// BillingView is the caller's own personal premium subscription state,
// nil when they have none (spec §1.4).
type BillingView struct {
	Plan              string  `json:"plan"`
	Status            string  `json:"status"`
	CurrentPeriodEnd  *string `json:"current_period_end"`
	CancelAtPeriodEnd bool    `json:"cancel_at_period_end"`
	catalog.PriceFacts
}

// GuildBillingView is a guild's billing state, shown only to a verified
// officer/leader of that guild (spec §1.4; Ruling C).
type GuildBillingView struct {
	// Scope is "officer_tools" for a grant covering only the officer views,
	// "plan" for a whole guild plan.
	Scope                string  `json:"scope"`
	Tier                 string  `json:"tier"`
	YouHoldSeat          bool    `json:"you_hold_seat"`
	Status               string  `json:"status"`
	CurrentPeriodEnd     *string `json:"current_period_end"`
	CancelAtPeriodEnd    bool    `json:"cancel_at_period_end"`
	BilledBy             string  `json:"billed_by"`
	YouAreBillingContact bool    `json:"you_are_billing_contact"`
}

// SeatFactsSource is the part of the seats store the account read needs. It
// is an interface because the seats package imports this one.
type SeatFactsSource interface {
	SeatFactsForUser(ctx context.Context, userID int64, now time.Time) (SeatFacts, error)
}

// MeSeat is /v1/me's seat: the Premium seat the account holds.
type MeSeat struct {
	GuildName string  `json:"guild_name"`
	GivenBy   string  `json:"given_by"`
	Since     string  `json:"since"`
	StartsAt  *string `json:"starts_at"`
}

// MeSeatOffer is /v1/me's seat_offer; the own_* fields are the caller's own Premium row.
type MeSeatOffer struct {
	GuildID        int64   `json:"guild_id"`
	GuildName      string  `json:"guild_name"`
	FromName       string  `json:"from_name"`
	OfferedAt      string  `json:"offered_at"`
	OwnEndsAt      *string `json:"own_ends_at"`
	OwnAmountCents *int64  `json:"own_amount_cents"`
	OwnInterval    *string `json:"own_interval"`
}

// MeEndedSeat is /v1/me's ended_seat, shown for the comeback's lifetime.
type MeEndedSeat struct {
	GuildName      string  `json:"guild_name"`
	EndedAt        string  `json:"ended_at"`
	ResubscribeURL string  `json:"resubscribe_url"`
	AmountCents    *int64  `json:"amount_cents"`
	Interval       *string `json:"interval"`
}

// SeatViews are the three seat keys of /v1/me; each is omitted while absent.
type SeatViews struct {
	Seat      *MeSeat      `json:"seat,omitempty"`
	SeatOffer *MeSeatOffer `json:"seat_offer,omitempty"`
	EndedSeat *MeEndedSeat `json:"ended_seat,omitempty"`
}

// personalBillingView is the account page's view of the caller's own Premium row.
func personalBillingView(b *entitlements.Billing) *BillingView {
	return &BillingView{
		Plan: b.Plan, Status: personalBillingStatus(b), CancelAtPeriodEnd: b.CancelAtPeriodEnd,
		CurrentPeriodEnd: rfc3339Ptr(b.CurrentPeriodEnd),
		PriceFacts:       catalog.FactsOf(b.PriceLookupKey, b.UnitAmountCents),
	}
}

// personalBillingStatus is the status the account page shows: an expired
// grant keeps status 'active' in the row but is no longer live, so it shows
// as an ended row does, "canceled". Every other row shows its own status.
func personalBillingStatus(b *entitlements.Billing) string {
	if !b.Active && entitlements.StatusGrantsAccess(b.Status) {
		return billingStatusCanceled
	}
	return b.Status
}

// seatFacts reads the account's seat facts; none when no source is wired.
// A failed read is logged and leaves the seat keys out rather than failing /v1/me.
func (s *Service) seatFacts(ctx context.Context, userID int64) SeatFacts {
	if s.Seats == nil {
		return SeatFacts{}
	}
	facts, err := s.Seats.SeatFactsForUser(ctx, userID, time.Now())
	if err != nil {
		s.logger().ErrorContext(ctx, "auth: me: seat facts", "user", userID, "err", err)
		return SeatFacts{}
	}
	return facts
}

// seatViewsOf renders the seat facts; own is the caller's own Premium row.
func seatViewsOf(facts SeatFacts, own *BillingView) SeatViews {
	var v SeatViews
	if f := facts.Seat; f != nil {
		v.Seat = &MeSeat{
			GuildName: f.GuildName, GivenBy: f.GivenBy,
			Since: f.Since.UTC().Format(time.RFC3339), StartsAt: rfc3339Ptr(f.StartsAt),
		}
	}
	if f := facts.Offer; f != nil {
		offer := &MeSeatOffer{
			GuildID: f.GuildID, GuildName: f.GuildName, FromName: f.FromName,
			OfferedAt: f.OfferedAt.UTC().Format(time.RFC3339),
		}
		if own != nil {
			offer.OwnEndsAt, offer.OwnAmountCents, offer.OwnInterval = own.CurrentPeriodEnd, own.AmountCents, own.Interval
		}
		v.SeatOffer = offer
	}
	if f := facts.Ended; f != nil {
		v.EndedSeat = endedSeatView(f)
	}
	return v
}

func endedSeatView(f *EndedSeat) *MeEndedSeat {
	price := catalog.FactsOf(f.PriceLookupKey, 0)
	link := plansPath
	if price.Interval != nil {
		link = resubscribePath + "?plan=premium&interval=" + *price.Interval +
			"&comeback=" + strconv.FormatInt(f.ComebackID, 10)
	}
	return &MeEndedSeat{
		GuildName: f.GuildName, EndedAt: f.EndedAt.UTC().Format(time.RFC3339),
		ResubscribeURL: link,
		AmountCents:    price.AmountCents, Interval: price.Interval,
	}
}

// attachGuildPlans fills Guild.Plan for every guild in guilds where
// userID is a verified officer/leader and that guild currently has an
// active plan (spec §1.4; Ruling C). heldGuildIDs are the guilds the
// account holds a Premium seat in.
func (s *Service) attachGuildPlans(ctx context.Context, guilds []Guild, userID int64, heldGuildIDs []int64) ([]Guild, error) {
	if s.Entitlements == nil {
		return guilds, nil
	}
	out := make([]Guild, len(guilds))
	copy(out, guilds)
	for i, g := range out {
		if !g.Verified || (g.Rank != "officer" && g.Rank != "leader") {
			continue
		}
		gb, err := s.Entitlements.GuildBilling(ctx, g.ID)
		if err != nil {
			return nil, fmt.Errorf("auth: me: guild billing %d: %w", g.ID, err)
		}
		if gb == nil || !gb.Active {
			continue
		}
		billedBy, youAre := s.billedBy(ctx, gb, userID)
		out[i].Plan = &GuildBillingView{
			Scope: guildPlanScope(gb), Tier: gb.Tier, YouHoldSeat: slices.Contains(heldGuildIDs, g.ID),
			Status: gb.Status, CancelAtPeriodEnd: gb.CancelAtPeriodEnd,
			CurrentPeriodEnd: rfc3339Ptr(gb.CurrentPeriodEnd), BilledBy: billedBy, YouAreBillingContact: youAre,
		}
	}
	return out, nil
}

// billedBy names who pays for a guild plan: "" for a CLI grant, "you" for the caller.
func (s *Service) billedBy(ctx context.Context, gb *entitlements.GuildBilling, userID int64) (string, bool) {
	switch {
	case gb.BillingUserID == nil:
		return "", false
	case *gb.BillingUserID == userID:
		return "you", true
	}
	bu, err := s.Store.User(ctx, *gb.BillingUserID)
	if err != nil {
		return "a former member", false
	}
	return bu.PublicName(), false
}

// guildPlanScope names what a guild's live entitlement row covers.
func guildPlanScope(gb *entitlements.GuildBilling) string {
	if gb.Scoped() {
		return guildPlanScopeOfficerTools
	}
	return guildPlanScopeWhole
}
