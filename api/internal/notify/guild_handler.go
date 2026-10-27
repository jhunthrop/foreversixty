package notify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/refusal"
)

const (
	guildChannelPath = "/v1/guilds/{id}/notifications/discord"
	guildTestPath    = guildChannelPath + "/test"

	noGuildMessage   = "no such guild"
	noChannelMessage = "no channel is connected"
)

// GuildStanding is a viewer's standing in a guild as the channel routes need it. The
// guilds package answers it, so the rank, plan and freeze rules stay there.
type GuildStanding struct {
	// Officer is a verified officer or leader of the guild.
	Officer bool
	// Plan is the guild plan (the officer views feature, a grant included).
	Plan bool
	// Frozen is a contested claim: officer actions wait for a moderator.
	Frozen bool
	// Name is the guild's display name; Path its page on the site ("/guild/us/normal/slug").
	Name, Path string
}

// GuildAccess answers GuildStanding; found is false when the guild does not exist.
type GuildAccess interface {
	GuildStanding(ctx context.Context, guildID, userID int64) (standing GuildStanding, found bool, err error)
}

// GuildConnectEnabled reports whether the guild channel is on in this deployment: the
// Discord connect is configured and the guild lookups are wired.
func (s *Service) GuildConnectEnabled() bool {
	return s != nil && s.OAuth != nil && s.Guilds != nil && s.GuildStore != nil
}

// GuildChannel is the guild's saved channel for the settings page and the raid view.
func (s *Service) GuildChannel(ctx context.Context, guildID int64) (GuildWebhook, bool, error) {
	return s.GuildStore.Get(ctx, guildID)
}

func (s *Service) mountGuild(mux *http.ServeMux) {
	if !s.GuildConnectEnabled() {
		return
	}
	limit := httpx.RateLimitByKey(testsPerMinute, time.Minute, auth.AccountRateLimitKey)
	connectLimit := httpx.RateLimitByKey(connectsPerMinute, time.Minute, auth.AccountRateLimitKey)
	mux.Handle("GET "+guildConnectPath, connectLimit(auth.RequireSession(s.guildConnect)))
	mux.Handle("POST "+guildTestPath, limit(auth.RequireSession(s.guildTest)))
	// Remove is never behind the plan: stopping posts must always work.
	mux.HandleFunc("DELETE "+guildChannelPath, auth.RequireSession(s.guildRemove))
}

// guildOfficer is GuildOfficerGate for the webhook channel routes.
func (s *Service) guildOfficer(w http.ResponseWriter, r *http.Request, needPlan bool) (int64, GuildStanding, bool) {
	return GuildOfficerGate(w, r, s.Guilds, needPlan, func(op string, err error) {
		s.logger().Error("notify", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	})
}

// GuildOfficerGate runs the guild routes' checks in order: a verified officer or leader
// (404 for anyone else, so a non-member never learns the route), the guild plan when
// needPlan (403 guild_plan_required), and the claim freeze (409). A lookup error is logged
// through fail and answered 500. ok is false once it has answered.
func GuildOfficerGate(w http.ResponseWriter, r *http.Request, access GuildAccess, needPlan bool, fail func(op string, err error)) (int64, GuildStanding, bool) {
	guildID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || guildID <= 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", noGuildMessage, nil)
		return 0, GuildStanding{}, false
	}
	standing, found, err := access.GuildStanding(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	switch {
	case err != nil:
		fail("guild standing", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that just now", nil)
	case !found || !standing.Officer:
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", noGuildMessage, nil)
	case needPlan && !standing.Plan:
		httpx.WriteError(w, r, http.StatusForbidden, "guild_plan_required", refusal.RaidPlanRequired, nil)
	case standing.Frozen:
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested", refusal.ClaimContested, nil)
	default:
		return guildID, standing, true
	}
	return 0, GuildStanding{}, false
}

func (s *Service) guildTest(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, standing, ok := s.guildOfficer(w, r, true)
	if !ok {
		return
	}
	url, err := s.openGuild(r.Context(), guildID)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", noChannelMessage, nil)
		return
	}
	if err != nil {
		s.fail(w, r, "guild open", err)
		return
	}
	embed := GuildTestEmbed(s.PublicBaseURL, standing.Name, standing.Path, time.Now())
	if err := s.Poster.Post(r.Context(), url, embed); err != nil {
		// A manual test never counts toward the five failures (testFailed changes nothing).
		s.testFailed(w, r, err)
		return
	}
	// A test that lands switches a switched-off channel back on.
	if err := s.GuildStore.RecordSuccess(r.Context(), guildID); err != nil {
		s.fail(w, r, "guild record success", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) guildRemove(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	guildID, _, ok := s.guildOfficer(w, r, false)
	if !ok {
		return
	}
	saved := s.guildURL(r.Context(), guildID)
	if err := s.GuildStore.Delete(r.Context(), guildID); err != nil {
		s.fail(w, r, "guild delete", err)
		return
	}
	if saved != "" {
		s.deleteAtDiscord(r.Context(), saved, "guild remove")
	}
	w.WriteHeader(http.StatusNoContent)
}

// openGuild unseals the guild's webhook URL; ErrNotFound when none is saved.
func (s *Service) openGuild(ctx context.Context, guildID int64) (string, error) {
	box, err := s.GuildStore.TokenBox(ctx, guildID)
	if err != nil {
		return "", err
	}
	url, err := s.Sealer.Open(box)
	if err != nil {
		return "", fmt.Errorf("notify: open guild webhook %d: %w", guildID, err)
	}
	return url, nil
}

// guildURL is openGuild for Discord-side cleanup: no channel and a failed read both read
// as "" (the failure is logged), so the caller's own change never fails on cleanup.
func (s *Service) guildURL(ctx context.Context, guildID int64) string {
	url, err := s.openGuild(ctx, guildID)
	if errors.Is(err, ErrNotFound) {
		return ""
	}
	if err != nil {
		s.logOpenFailure(err)
		return ""
	}
	return url
}
