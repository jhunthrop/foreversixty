// api/internal/guilds/recruit_handler.go
//
// The recruitment routes: the public card and its officer edit, a player's own applications,
// and the officer list, votes and decisions. Officer routes gate on verifiedOfficerOrLeader (403)
// and, for writes, the claim freeze (officerNotFrozen); an applicant's routes on their account.
package guilds

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	// applyPerHour is the per-IP cap on filing applications.
	applyPerHour = 10
	// recruitCacheSeconds is the public card's edge cache.
	recruitCacheSeconds = 60
	recruitFailMessage  = "could not load that guild's recruiting just now"
	applyFailMessage    = "could not file that application just now"
	officerAppsFail     = "could not load applications just now"
)

// mountRecruitRoutes registers the recruitment, application, profile and note routes.
func mountRecruitRoutes(mux *http.ServeMux, s *Service, trustedProxyHops int) {
	mux.HandleFunc("GET /v1/guilds/{id}/recruit", s.getRecruit)
	mux.HandleFunc("PATCH /v1/guilds/{id}/recruit", auth.RequireSession(s.patchRecruit))
	apply := httpx.RateLimitPer(applyPerHour, time.Hour, trustedProxyHops)
	mux.Handle("POST /v1/guilds/{id}/applications", apply(auth.RequireSession(s.apply)))
	mux.HandleFunc("GET /v1/guilds/{id}/applications/mine", auth.RequireSession(s.myApplications))
	mux.HandleFunc("POST /v1/guilds/{id}/applications/{app}/withdraw", auth.RequireSession(s.withdraw))
	mux.HandleFunc("GET /v1/guilds/{id}/applications", auth.RequireSession(s.officerApplications))
	mux.HandleFunc("PUT /v1/guilds/{id}/applications/{app}/vote", auth.RequireSession(s.putVote))
	mux.HandleFunc("DELETE /v1/guilds/{id}/applications/{app}/vote", auth.RequireSession(s.deleteVote))
	mux.HandleFunc("POST /v1/guilds/{id}/applications/{app}/accept", auth.RequireSession(s.decideRoute(true)))
	mux.HandleFunc("POST /v1/guilds/{id}/applications/{app}/decline", auth.RequireSession(s.decideRoute(false)))
	mountMemberRoutes(mux, s)
}

func writeGuildNotFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
}

// officerGuild answers the guild id for a verified officer or guild master, writing the 404 or
// 403 itself otherwise.
func (s *Service) officerGuild(w http.ResponseWriter, r *http.Request, op, failMessage string) (int64, bool) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return 0, false
	}
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, op, err, failMessage)
		return 0, false
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you must be a verified officer of this guild", nil)
		return 0, false
	}
	return guildID, true
}

// officerWriteGuild is officerGuild plus the claim freeze.
func (s *Service) officerWriteGuild(w http.ResponseWriter, r *http.Request, op, failMessage string) (int64, bool) {
	guildID, ok := s.officerGuild(w, r, op, failMessage)
	if !ok || !s.officerNotFrozen(w, r, guildID, op, failMessage) {
		return 0, false
	}
	return guildID, true
}

// applicationIDFrom parses {app}, writing the 404 itself when it is not an id.
func applicationIDFrom(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("app"), 10, 64)
	if err != nil || id <= 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such application", nil)
		return 0, false
	}
	return id, true
}

// getRecruit is GET /v1/guilds/{id}/recruit: public.
func (s *Service) getRecruit(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return
	}
	card, err := s.Store.RecruitCard(r.Context(), guildID)
	if errors.Is(err, ErrNotFound) {
		writeGuildNotFound(w, r)
		return
	}
	if err != nil {
		s.fail(w, r, "recruit", err, recruitFailMessage)
		return
	}
	httpx.CachePublic(w, recruitCacheSeconds*time.Second, 0)
	httpx.WriteOK(w, r, http.StatusOK, card)
}

// patchRecruit is PATCH /v1/guilds/{id}/recruit: officers edit the card and the recruit role.
func (s *Service) patchRecruit(w http.ResponseWriter, r *http.Request) {
	guildID, ok := s.officerWriteGuild(w, r, "patch_recruit", recruitFailMessage)
	if !ok {
		return
	}
	var in RecruitPatch
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	var invalid ErrInvalidRecruit
	err := s.Store.UpdateRecruit(r.Context(), guildID, in)
	switch {
	case errors.As(err, &invalid):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a valid recruiting setting",
			map[string]string{invalid.Field: invalid.Rule})
		return
	case errors.Is(err, ErrNotFound):
		writeGuildNotFound(w, r)
		return
	case err != nil:
		s.fail(w, r, "patch_recruit", err, recruitFailMessage)
		return
	}
	s.logger().Info("guilds", "op", "recruit_update", "guild_id", guildID, "user_id", auth.ActorFrom(r.Context()).UserID)
	s.writeRecruitSettings(w, r, guildID)
}

// writeRecruitSettings answers the card with the officer-only recruit role beside it.
func (s *Service) writeRecruitSettings(w http.ResponseWriter, r *http.Request, guildID int64) {
	card, err := s.Store.RecruitCard(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "recruit_settings", err, recruitFailMessage)
		return
	}
	roleID, err := s.Store.RecruitRoleID(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "recruit_settings", err, recruitFailMessage)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"card": card, "discord_role_id": roleID})
}

// apply is POST /v1/guilds/{id}/applications.
func (s *Service) apply(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return
	}
	var in ApplicationInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	actor := auth.ActorFrom(r.Context())
	a, err := s.Store.Apply(r.Context(), guildID, actor.UserID, in)
	if err != nil {
		s.writeApplyError(w, r, err)
		return
	}
	s.logger().Info("guilds", "op", "apply", "guild_id", guildID, "user_id", actor.UserID, "application_id", a.ID)
	s.publishApplicant(r.Context(), guildID, a)
	httpx.WriteOK(w, r, http.StatusCreated, a)
}

func (s *Service) writeApplyError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		writeGuildNotFound(w, r)
	case errors.Is(err, ErrInvalidApplication):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a valid application", map[string]string{
			"character_key": "one of your characters", "message": "at most 1000 characters",
			"availability": "at most 200 characters"})
	case errors.Is(err, ErrNotApplicantsChar):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "apply with one of your own characters", nil)
	case errors.Is(err, ErrWrongRealm):
		httpx.WriteError(w, r, http.StatusBadRequest, "wrong_realm", "that character is not on this guild's region and ruleset", nil)
	case errors.Is(err, ErrNotRecruiting):
		httpx.WriteError(w, r, http.StatusConflict, "not_recruiting", "this guild is not taking applications right now", nil)
	case errors.Is(err, ErrAlreadyApplied):
		httpx.WriteError(w, r, http.StatusConflict, "already_applied", "you already have an application open with this guild", nil)
	case errors.Is(err, ErrAlreadyMember):
		httpx.WriteError(w, r, http.StatusConflict, "already_member", "you are already in this guild", nil)
	default:
		s.fail(w, r, "apply", err, applyFailMessage)
	}
}

// myApplications is GET /v1/guilds/{id}/applications/mine: the applicant's own, newest first.
func (s *Service) myApplications(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return
	}
	apps, err := s.Store.MyApplications(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "my_applications", err, applyFailMessage)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": apps})
}

// withdraw is POST /v1/guilds/{id}/applications/{app}/withdraw: the applicant's own only.
func (s *Service) withdraw(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		writeGuildNotFound(w, r)
		return
	}
	id, ok := applicationIDFrom(w, r)
	if !ok {
		return
	}
	a, err := s.Store.Withdraw(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID, id)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such application", nil)
	case errors.Is(err, ErrApplicationClosed):
		httpx.WriteError(w, r, http.StatusConflict, "application_closed", "that application is already decided or withdrawn", nil)
	case err != nil:
		s.fail(w, r, "withdraw", err, applyFailMessage)
	default:
		httpx.WriteOK(w, r, http.StatusOK, a)
	}
}
