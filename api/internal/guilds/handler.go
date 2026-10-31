// api/internal/guilds/handler.go
package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// maxJSONBody is the ceiling on the small JSON bodies these routes take.
const maxJSONBody = 8 << 10

// inviteAcceptPerHour is the contract's redemption rate limit.
const inviteAcceptPerHour = 20

// contestPerHour is the per-IP cap on claim/contest attempts (A2,
// 2026-09-21 second security review response): "rate-limited like
// claim" per the spec, applied here as the same kind of per-IP
// httpx.RateLimitPer wrap the invite-accept route uses, at a much
// lower ceiling than invite-accept's 20/hour since a successful
// contest can freeze a guild's officer tools - a bulk-account attempt
// against many guilds must be slowed at the IP layer even before the
// per-account, 30-day DB-level limit (checkContestRateLimit) ever
// gets involved.
const contestPerHour = 5

// Accounts is the part of auth.Store the guilds handlers need — the same
// two methods reports.Accounts already reads.
type Accounts interface {
	User(ctx context.Context, id int64) (auth.User, error)
	GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error)
}

// Service serves every guild-mutation route.
type Service struct {
	Store    *Store
	Accounts Accounts
	// R2 is the object store the crest routes write to, read from, and delete from
	// (crest.go). nil in a deployment or test harness with no object storage
	// configured - every crest route then answers 503 rather than panicking on a nil
	// client, the same shape reports.Service.Signer being nil already takes.
	R2 Objects
	// Supporters adds the supporter mark to the home roster; nil means nobody shows it.
	Supporters Supporters
	// Entitlements answers the home viewer's officer_tools flag; nil means false.
	Entitlements Entitlements
	// Discord is the guild's raid channel (notify); nil leaves settings.discord out and
	// every raid view's discord block false.
	Discord DiscordChannels
	// Bot is the Forever Sixty bot; nil leaves bot_available false.
	Bot DiscordBot
	// RaidEvents posts the raid embeds to the guild's channel; nil posts nothing.
	RaidEvents RaidEvents
	// CoachHolders tells whether any raider in a log holds log_coach; nil: never.
	CoachHolders CoachHolders
	// PublicBaseURL is the site origin the raid posts link to.
	PublicBaseURL string
	Log           *slog.Logger
	// RaidPosts is told when a raid changes so its channel post can be refreshed; nil
	// tells nothing.
	RaidPosts RaidPostMarker
	// SignupLimiter is the per-account sign-up budget the PUT route and SignUp share; nil
	// builds the default on first use.
	SignupLimiter *httpx.KeyLimiter
	limiterOnce   sync.Once
	// Events publishes the player notifications (events.Router); nil publishes nothing.
	Events events.Publisher
	// AppSignups (APP_NOTIFICATIONS_ENABLED) lets a signed-in app's token answer the sign-up
	// PUT, the route the apps' In, Late and Out buttons call. Off, it takes a browser session.
	AppSignups bool
	// raidPostBudget bounds one request's raid posts together; zero means the default.
	raidPostBudget time.Duration
}

// Mount registers every /v1/guilds/{id}/... mutation route and the
// invite-accept route. Reads of a public guild page stay with
// rankings.Mount; this package only ever mutates or serves a signed-in
// member's own view.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	mux.HandleFunc("POST /v1/guilds/{id}/claim", auth.RequireSession(s.claim))
	mux.HandleFunc("POST /v1/guilds/{id}/claim/confirm", auth.RequireSession(s.confirmClaim))
	mux.HandleFunc("POST /v1/guilds/{id}/claim/release", auth.RequireSession(s.releaseClaim))
	contest := httpx.RateLimitPer(contestPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/guilds/{id}/claim/contest", contest(auth.RequireSession(s.contestClaim)))
	mux.HandleFunc("POST /v1/guilds/{id}/claim/resolve", auth.RequireSession(s.resolveClaim))
	// No auth.RequireSession wrap (fifth security review response, same
	// hidden-standing pattern as moderationClaims below): the handler
	// itself answers 404 for a non-moderator.
	mux.HandleFunc("POST /v1/guilds/{id}/claim/reopen", s.reopenClaimCooldown)
	mux.HandleFunc("GET /v1/guilds/{id}/settings", auth.RequireSession(s.getSettings))
	mux.HandleFunc("PATCH /v1/guilds/{id}/settings", auth.RequireSession(s.patchSettings))
	mux.HandleFunc("POST /v1/guilds/{id}/invite/rotate", auth.RequireSession(s.rotateInvite))
	accept := httpx.RateLimitPer(inviteAcceptPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/guilds/invite/{token}/accept", accept(auth.RequireSession(s.acceptInvite)))
	// The invite page's lookup: public (a signed-out visitor sees the same card), under
	// its own prefix because /v1/guilds/invite/{token}/... would collide with the
	// rankings route GET /v1/guilds/{region}/{ruleset}/{name}.
	mux.Handle("GET /v1/guild-invites/{token}", accept(http.HandlerFunc(s.peekInvite)))
	mux.HandleFunc("POST /v1/guilds/{id}/characters/{region}/{ruleset}/{name}/approve", auth.RequireSession(s.approveCharacter))
	mux.HandleFunc("POST /v1/guilds/{id}/roster/approve-all", auth.RequireSession(s.approveAll))
	mux.HandleFunc("DELETE /v1/guilds/{id}/characters/{region}/{ruleset}/{name}", auth.RequireSession(s.removeCharacter))
	mux.HandleFunc("PATCH /v1/guilds/{id}/members/me", auth.RequireSession(s.patchConsent))
	mux.HandleFunc("DELETE /v1/guilds/{id}/members/me", auth.RequireSession(s.leaveGuild))
	mux.HandleFunc("PUT /v1/guilds/{id}/crest", auth.RequireSession(s.putCrest))
	mux.HandleFunc("DELETE /v1/guilds/{id}/crest", auth.RequireSession(s.deleteCrest))
	// No RequireSession wrap: a guild's crest is a public mark, served wherever the
	// faction logo it falls back to would be, to a signed-out browser included.
	mux.HandleFunc("GET /v1/guilds/{id}/crest.webp", s.getCrest)
	mux.HandleFunc("GET /v1/guilds/{id}/home", auth.RequireSession(s.home))
	// Raids is reachable signed-out (design spec §4.0: public reports only) - no
	// RequireSession wrap; the handler reads whatever actor the auth middleware already
	// put in context (the zero Actor for an anonymous request) and the visibility clause
	// does the rest, a flagged report included (shown only to a verified member who owns it).
	mux.HandleFunc("GET /v1/guilds/{id}/raids", s.raids)
	// The raid calendar (design/specs/2026-10-26-guild-raids.md §9). The raid read is
	// session-optional like /raids and answers 404 to anyone but a verified member (R1).
	mux.HandleFunc("GET /v1/guilds/{id}/raids/{raid}", s.getRaid)
	officerLimit := officerRaidRateLimit()
	mux.Handle("POST /v1/guilds/{id}/raids", officerLimit(auth.RequireSession(s.createRaid)))
	mux.Handle("PATCH /v1/guilds/{id}/raids/{raid}", officerLimit(auth.RequireSession(s.patchRaid)))
	for action := range raidStatusMoves {
		mux.Handle("POST /v1/guilds/{id}/raids/{raid}/"+action, officerLimit(auth.RequireSession(s.raidStatusAction(action))))
	}
	mountRaidMatchRoutes(mux, s, officerLimit)
	mountRaidReportRoutes(mux, s, officerLimit)
	mux.Handle("PUT /v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}",
		s.signupRateLimit()(s.signupAuth()(s.putSignup)))
	// Progression is public (design spec §4.0) - no RequireSession wrap - and its counts
	// depend on the viewer: the handler reads whatever actor the auth middleware put in
	// context, and a verified member gets the guild's own numbers (private nights counted,
	// a private response) where anyone else gets the public guild page's (shareable).
	mux.HandleFunc("GET /v1/guilds/{id}/progression", s.progression)
	// Readiness is member/officer only (design spec §4.0: "Not shown" to the public).
	// With ?raid= it is a raid route and keeps R1 (readinessRoute).
	mux.HandleFunc("GET /v1/guilds/{id}/readiness", s.readinessRoute)
	// CSV exports are officer reads by rank in any plan state (history is never locked).
	mux.HandleFunc("GET /v1/guilds/{id}/export/attendance.csv", auth.RequireSession(s.exportAttendance))
	mux.HandleFunc("GET /v1/guilds/{id}/export/loot.csv", auth.RequireSession(s.exportLoot))
	// Loot is member-and-officer (member read-only, officer gets the Award control).
	mux.HandleFunc("GET /v1/guilds/{id}/loot", auth.RequireSession(s.loot))
	mux.HandleFunc("POST /v1/guilds/{id}/loot/awards", auth.RequireSession(s.createLootAward))
	mux.HandleFunc("DELETE /v1/guilds/{id}/loot/awards/{award_id}", auth.RequireSession(s.deleteLootAward))
	// No auth.RequireSession wrap (item 4, fourth security review
	// response): the handler itself answers 404 for a non-moderator,
	// including an unauthenticated caller, so the route's existence is
	// never advertised by a 401/403 that a wrap would otherwise leak.
	mux.HandleFunc("GET /v1/moderation/claims", s.moderationClaims)
}

// signupLimiter is the shared sign-up budget, built once when none was injected.
func (s *Service) signupLimiter() *httpx.KeyLimiter {
	s.limiterOnce.Do(func() {
		if s.SignupLimiter == nil {
			s.SignupLimiter = httpx.NewKeyLimiter(signupsPerMinute, time.Minute)
		}
	})
	return s.SignupLimiter
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("guilds", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// guildIDFrom parses the {id} path value every route in this package takes.
func guildIDFrom(r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return id, err == nil && id > 0
}

// verifiedOfficerOrLeader reports whether the actor is a verified
// officer or leader of guildID — no moderator bypass, unlike
// verifiedOfficerOrModerator (settings), matching the spec's own auth
// column for invite rotation and roster approve.
func (s *Service) verifiedOfficerOrLeader(r *http.Request, guildID int64) (bool, error) {
	return s.officerOrLeaderOf(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
}

// officerOrLeaderOf is verifiedOfficerOrLeader for an account named rather than read from a
// request.
func (s *Service) officerOrLeaderOf(ctx context.Context, guildID, userID int64) (bool, error) {
	rank, ok, err := s.Accounts.GuildRank(ctx, guildID, userID)
	if err != nil || !ok {
		return false, err
	}
	return rank == "officer" || rank == "leader", nil
}

// freezeCheck applies the officer-power freeze gate every officer route
// in this package shares (A4, 2026-09-21 security review response,
// simplified by the fourth response): a route is blocked with 409
// claim_contested whenever the guild's claim is contested - always,
// with no exception for an established or otherwise-corroborated
// claim, since both narrower rules this package tried were found
// gameable by a squatter and the fourth response deleted the escape
// rather than patch it again. moderatorExempt lets a route that
// already gave the caller moderator standing (settings) also give them
// the bypass every moderator implicitly needs to be the one who
// resolves a freeze in the first place; a route with no moderator
// standing at all passes false and the freeze blocks everyone the same
// way.
func (s *Service) freezeCheck(ctx context.Context, guildID int64, moderatorExempt bool) (bool, error) {
	if moderatorExempt {
		return false, nil
	}
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return false, err
	}
	return claimState(g, time.Now()).Frozen, nil
}

// decodeJSON decodes r's body into v, capped at maxJSONBody, answering 400
// in the envelope on failure. Callers return immediately when it errors.
// signupAuth is who may call the sign-up PUT: a browser session, and a signed-in app's
// token too when the apps' sign-up is on. A companion's token never signs up.
func (s *Service) signupAuth() func(http.HandlerFunc) http.HandlerFunc {
	if s.AppSignups {
		return auth.RequireSession
	}
	return auth.RequireBrowserSession
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v any) error {
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(v); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON", nil)
		return err
	}
	return nil
}

func (s *Service) claim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	u, err := s.Accounts.User(r.Context(), actor.UserID)
	if err != nil {
		s.fail(w, r, "claim", err, "could not claim that guild just now")
		return
	}
	result, err := s.Store.Claim(r.Context(), guildID, actor.UserID, u.BnetSub != "")
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrAlreadyClaimed), errors.Is(err, ErrClaimPending), errors.Is(err, ErrAlreadyClaimsAnotherGuild):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that guild already has a claim, or you already hold another guild's claim", nil)
	case errors.Is(err, ErrClaimRateLimited):
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited",
			"you may only attempt one guild claim every 30 days", nil)
	case errors.Is(err, ErrNoBattleNetIdentity):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"claiming as guild master requires a linked Battle.net account", nil)
	case errors.Is(err, ErrNotEligible):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you need an officer or leader character in this guild to claim it", nil)
	case err != nil:
		s.fail(w, r, "claim", err, "could not claim that guild just now")
	default:
		s.logger().Info("guilds", "op", "claim", "guild_id", guildID, "user_id", actor.UserID, "status", result.Status)
		httpx.WriteOK(w, r, http.StatusOK, result)
	}
}

func (s *Service) confirmClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	claimant, err := s.Store.ConfirmClaim(r.Context(), guildID, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound), errors.Is(err, ErrNoPendingClaim):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no pending claim on that guild", nil)
	case errors.Is(err, ErrSameAccount), errors.Is(err, ErrNotEligible):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you need a second, distinct officer or leader character to confirm this claim", nil)
	case err != nil:
		s.fail(w, r, "confirm_claim", err, "could not confirm that claim just now")
	default:
		u, uerr := s.Accounts.User(r.Context(), claimant)
		if uerr != nil {
			s.fail(w, r, "confirm_claim", uerr, "could not confirm that claim just now")
			return
		}
		s.logger().Info("guilds", "op", "claim_confirm", "guild_id", guildID, "confirmer_id", actor.UserID, "claimant_id", claimant)
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{
			"status": "confirmed", "claimed_by": map[string]string{"battletag": u.PublicName()},
		})
	}
}

func (s *Service) releaseClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := s.Store.ReleaseClaim(r.Context(), guildID, actor.UserID, actor.IsModerator())
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrNotClaimant):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you do not hold this guild's claim", nil)
	case err != nil:
		s.fail(w, r, "release_claim", err, "could not release that claim just now")
	default:
		s.logger().Info("guilds", "op", "claim_release", "guild_id", guildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "released"})
	}
}

// CoachHolders answers the report_ready post's coach sentence, which names nobody.
type CoachHolders interface {
	// AnyCoachHolder reports whether any raider in the log holds log_coach on an owned character.
	AnyCoachHolder(ctx context.Context, reportID string) (bool, error)
}
