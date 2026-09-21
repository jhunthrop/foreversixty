// api/internal/guilds/handler.go
package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// maxJSONBody is the ceiling on the small JSON bodies these routes take.
const maxJSONBody = 8 << 10

// inviteAcceptPerHour is the contract's redemption rate limit.
const inviteAcceptPerHour = 20

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
	Log      *slog.Logger
}

// Mount registers every /v1/guilds/{id}/... mutation route and the
// invite-accept route. Reads of a public guild page stay with
// rankings.Mount; this package only ever mutates or serves a signed-in
// member's own view.
func Mount(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	mux.HandleFunc("POST /v1/guilds/{id}/claim", auth.RequireSession(s.claim))
	mux.HandleFunc("POST /v1/guilds/{id}/claim/confirm", auth.RequireSession(s.confirmClaim))
	mux.HandleFunc("POST /v1/guilds/{id}/claim/release", auth.RequireSession(s.releaseClaim))
	mux.HandleFunc("GET /v1/guilds/{id}/settings", auth.RequireSession(s.getSettings))
	mux.HandleFunc("PATCH /v1/guilds/{id}/settings", auth.RequireSession(s.patchSettings))
	mux.HandleFunc("POST /v1/guilds/{id}/invite/rotate", auth.RequireSession(s.rotateInvite))
	accept := httpx.RateLimitPer(inviteAcceptPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/guilds/invite/{token}/accept", accept(auth.RequireSession(s.acceptInvite)))
	mux.HandleFunc("POST /v1/guilds/{id}/characters/{region}/{ruleset}/{name}/approve", auth.RequireSession(s.approveCharacter))
	mux.HandleFunc("DELETE /v1/guilds/{id}/characters/{region}/{ruleset}/{name}", auth.RequireSession(s.removeCharacter))
	mux.HandleFunc("PATCH /v1/guilds/{id}/members/me", auth.RequireSession(s.patchConsent))
	mux.HandleFunc("DELETE /v1/guilds/{id}/members/me", auth.RequireSession(s.leaveGuild))
	mux.HandleFunc("GET /v1/guilds/{id}/home", auth.RequireSession(s.home))
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
	actor := auth.ActorFrom(r.Context())
	rank, ok, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
	if err != nil || !ok {
		return false, err
	}
	return rank == "officer" || rank == "leader", nil
}

// decodeJSON decodes r's body into v, capped at maxJSONBody, answering 400
// in the envelope on failure. Callers return immediately when it errors.
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
	result, err := s.Store.Claim(r.Context(), guildID, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrAlreadyClaimed), errors.Is(err, ErrClaimPending):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that guild already has a claim", nil)
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
