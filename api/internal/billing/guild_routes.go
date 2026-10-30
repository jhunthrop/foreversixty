// api/internal/billing/guild_routes.go
package billing

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

// guildReadRateLimitPerHour is the guild billing read's own per-IP budget.
// The settings page refetches it after every seat and transfer action, so it
// gets more room than the Stripe-minting writes' billingRateLimitPerHour.
const guildReadRateLimitPerHour = 120

// The billing read's `can` words.
const (
	canBilling  = "billing"
	canSeats    = "seats"
	canTransfer = "transfer"
)

// mountGuild registers the guild billing routes: the read, the Guild to Raid
// guild upgrade and the plan transfer offer.
func mountGuild(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	read := httpx.RateLimitPer(guildReadRateLimitPerHour, time.Hour, trustedProxyHops)
	write := httpx.RateLimitPer(billingRateLimitPerHour, time.Hour, trustedProxyHops)
	const base = "/v1/guilds/{id}/billing"
	mux.Handle("GET "+base, read(auth.RequireSession(s.guildBilling)))
	mux.Handle("GET "+base+"/upgrade-preview", write(auth.RequireSession(s.upgradePreview)))
	mux.Handle("POST "+base+"/upgrade", write(auth.RequireBrowserSession(s.upgrade)))
	mux.Handle("POST "+base+"/transfer", write(auth.RequireBrowserSession(s.sendTransfer)))
	mux.Handle("POST "+base+"/transfer/withdraw", write(auth.RequireBrowserSession(s.withdrawTransfer)))
	mux.Handle("POST "+base+"/transfer/decline", write(auth.RequireBrowserSession(s.declineTransfer)))
}

// guildActor is the caller of a guild billing route: a verified officer or
// leader of the guild, and what they may do there.
type guildActor struct {
	guildID    int64
	userID     int64
	payer      bool
	subscriber bool
	frozen     bool
}

// guildActorFor resolves the caller, answering the refusal itself when
// ok is false: 503 without seats wiring, 404 for a bad id or anyone who is
// neither a verified officer or leader of the guild nor its plan's verified
// billing user (a demoted billing user still pays, as on the seat routes).
func (s *Service) guildActorFor(w http.ResponseWriter, r *http.Request, op string) (guildActor, bool) {
	if s.Seats == nil {
		writeUnavailable(w, r)
		return guildActor{}, false
	}
	guildID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || guildID <= 0 {
		writeNoSuchGuild(w, r)
		return guildActor{}, false
	}
	a, admitted, err := s.loadGuildActor(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, op, err, "could not read billing just now")
		return guildActor{}, false
	}
	if !admitted {
		writeNoSuchGuild(w, r)
		return guildActor{}, false
	}
	return a, true
}

// loadGuildActor reads the caller's standing in the guild; admitted is
// false for anyone the billing routes do not serve.
func (s *Service) loadGuildActor(ctx context.Context, guildID, userID int64) (guildActor, bool, error) {
	a := guildActor{guildID: guildID, userID: userID}
	rank, verified, err := s.Accounts.GuildRank(ctx, guildID, userID)
	if err != nil || !verified {
		return guildActor{}, false, err
	}
	if a.subscriber, err = s.Seats.IsSubscriber(ctx, guildID, userID); err != nil {
		return guildActor{}, false, err
	}
	if rank != "officer" && rank != "leader" && !a.subscriber {
		return guildActor{}, false, nil
	}
	if s.Guilds != nil {
		if a.frozen, err = s.Guilds.FrozenClaimant(ctx, guildID, userID); err != nil {
			return guildActor{}, false, err
		}
	}
	if a.payer, err = s.Seats.IsPayer(ctx, guildID, userID); err != nil {
		return guildActor{}, false, err
	}
	return a, true, nil
}

// actingActor is guildActorFor for a write: a frozen claimant is refused,
// and allowed says whether the caller may make this write at all.
func (s *Service) actingActor(w http.ResponseWriter, r *http.Request, op string, allowed func(guildActor) bool) (guildActor, bool) {
	a, ok := s.guildActorFor(w, r, op)
	if !ok {
		return guildActor{}, false
	}
	if a.frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested", refusal.ClaimContested, nil)
		return guildActor{}, false
	}
	if !allowed(a) {
		writeNotPayer(w, r)
		return guildActor{}, false
	}
	return a, true
}

func writeNotPayer(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only the guild's payer can do that", nil)
}

func isPayer(a guildActor) bool      { return a.payer }
func isSubscriber(a guildActor) bool { return a.subscriber }
func anyOfficer(guildActor) bool     { return true }

// livePlan is the guild's live whole-plan row, answering 403
// guild_plan_required itself when there is none.
func (s *Service) livePlan(w http.ResponseWriter, r *http.Request, op string, guildID int64) (*entitlements.GuildBilling, bool) {
	gb, err := s.Entitlements.GuildBilling(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, op, err, "could not read billing just now")
		return nil, false
	}
	if gb == nil || !gb.Active || gb.Scoped() {
		httpx.WriteError(w, r, http.StatusForbidden, "guild_plan_required", "this guild has no guild plan", nil)
		return nil, false
	}
	return gb, true
}

func writeNoSuchGuild(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
}

// rfc3339 is an instant on the wire: UTC, RFC 3339, null when unknown.
func rfc3339(t *time.Time) *string {
	if t == nil {
		return nil
	}
	v := t.UTC().Format(time.RFC3339)
	return &v
}
