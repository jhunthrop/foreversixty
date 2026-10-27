package discordbot

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	linkStartPath = "/v1/discord/link/start"
	linkPath      = "/v1/discord/link"
	// linkStartsPerMinute caps link starts per IP.
	linkStartsPerMinute = 20

	oauthAccessDenied = "access_denied"

	// accountReturn is the account page's link return: the outcome follows the =.
	accountReturn = "/account?discord_link="
	youAnchor     = "#you"
	// signedOutNext is where sign-in returns a member who started a link signed out.
	signedOutNext = "/account?discord_link=start" + youAnchor

	linkLinked   = "linked"
	linkDenied   = "denied"
	linkFailed   = "failed"
	linkConflict = "conflict"
)

// mountLink registers the link routes; Mount calls it.
func (s *Service) mountLink(mux *http.ServeMux, trustedHops int) {
	limit := httpx.RateLimitPer(linkStartsPerMinute, time.Minute, trustedHops)
	mux.Handle("GET "+linkStartPath, limit(http.HandlerFunc(s.linkStart)))
	mux.HandleFunc("DELETE "+linkPath, auth.RequireSession(s.unlink))
}

// linkStartURL is where the link button sends the member.
func (s *Service) linkStartURL() string { return s.d.APIBaseURL + linkStartPath }

// LinkAvailable implements auth.DiscordLinkView.
func (s *Service) LinkAvailable() bool { return s.d.Notify != nil && s.d.OAuth != nil }

// DiscordLink implements auth.DiscordLinkView: the account's link, or, while unlinked, a
// guild whose bot a link would unlock.
func (s *Service) DiscordLink(ctx context.Context, userID int64) (*auth.DiscordLinkBlock, *string, error) {
	link, linked, err := s.d.Store.LinkByUser(ctx, userID)
	if err != nil {
		return nil, nil, err
	}
	if linked {
		return &auth.DiscordLinkBlock{DisplayName: link.DisplayName}, nil, nil
	}
	suggested, err := s.d.Store.SuggestedGuild(ctx, userID)
	return nil, suggested, err
}

func (s *Service) accountRedirect(w http.ResponseWriter, r *http.Request, outcome string) {
	http.Redirect(w, r, s.d.PublicBaseURL+accountReturn+outcome+youAnchor, http.StatusFound)
}

// linkStart sends a signed-in member to Discord's identify page, and a signed-out one
// through sign-in first, back to the account page, which resumes the link.
func (s *Service) linkStart(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() {
		http.Redirect(w, r, s.d.PublicBaseURL+"/login?next="+url.QueryEscape(signedOutNext), http.StatusFound)
		return
	}
	if actor.IsDevice() {
		s.accountRedirect(w, r, linkFailed)
		return
	}
	state, err := s.d.Notify.StartState(r.Context(), notify.StateKindLink, actor.UserID, nil)
	if err != nil {
		s.log.Error("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "link state", "err", err)
		s.accountRedirect(w, r, linkFailed)
		return
	}
	http.Redirect(w, r, s.d.OAuth.LinkURL(state), http.StatusFound)
}

// linkCallback finishes a link at the shared callback. Only the Discord user id and display
// name are kept; the code and the access token never leave the exchange.
func (s *Service) linkCallback(w http.ResponseWriter, r *http.Request, cs notify.ConnectState) {
	httpx.CachePrivateNoStore(w)
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() || actor.IsDevice() || actor.UserID != cs.UserID || cs.Expired {
		s.accountRedirect(w, r, linkFailed)
		return
	}
	query := r.URL.Query()
	code := query.Get("code")
	if code == "" || query.Get("error") == oauthAccessDenied {
		s.accountRedirect(w, r, linkDenied)
		return
	}
	s.accountRedirect(w, r, s.finishLink(r, cs.UserID, code))
}

// finishLink exchanges the code and saves the link, returning the outcome word.
func (s *Service) finishLink(r *http.Request, userID int64, code string) string {
	who, err := s.d.OAuth.ExchangeIdentity(r.Context(), code)
	if err != nil {
		s.log.Warn("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "link exchange", "err", err)
		return linkFailed
	}
	err = s.d.Store.SaveLink(r.Context(), userID, who.ID, who.DisplayName)
	if errors.Is(err, ErrLinkConflict) {
		return linkConflict
	}
	if err != nil {
		s.log.Error("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "link save", "err", err)
		return linkFailed
	}
	s.log.Info("discordbot", "op", "link", "user_id", userID)
	return linkLinked
}

// unlink removes the caller's link; 204 whether or not one existed.
func (s *Service) unlink(w http.ResponseWriter, r *http.Request) {
	if err := s.d.Store.DeleteLink(r.Context(), auth.ActorFrom(r.Context()).UserID); err != nil {
		s.log.Error("discordbot", "id", httpx.RequestIDFrom(r.Context()), "op", "unlink", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not unlink just now", nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
