package linkedrole

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const (
	// StartPath is the Role Connections Verification URL's path and the account row's link.
	StartPath = "/v1/discord/linked-role"
	// startsPerMinute caps starts per IP, as the Discord link start does.
	startsPerMinute = 20

	oauthAccessDenied = "access_denied"

	// accountReturn is the account page's return: the outcome follows the =.
	accountReturn = "/account?discord_role="
	youAnchor     = "#you"
	// signedOutNext is where sign-in returns a visitor who started signed out; the account
	// page goes straight on to StartPath.
	signedOutNext = accountReturn + outcomeStart + youAnchor

	outcomeLinked = "linked"
	outcomeDenied = "denied"
	outcomeFailed = "failed"
	outcomeStart  = "start"
)

// States is the shared Discord callback's state store; *notify.Service satisfies it.
type States interface {
	StartState(ctx context.Context, kind string, userID int64, guildID *int64) (string, error)
	OnState(kind string, h notify.StateHandler)
}

// Handler serves the verification start; its callback rides the shared Discord redirect URI
// (notify's callback), which hands it the taken state.
type Handler struct {
	Syncer        *Syncer
	States        States
	PublicBaseURL string
	Log           *slog.Logger
}

// NewHandler builds the handler and registers its callback with the shared state store.
func NewHandler(syncer *Syncer, states States, publicBaseURL string, log *slog.Logger) *Handler {
	h := &Handler{Syncer: syncer, States: states, PublicBaseURL: publicBaseURL, Log: log}
	if h.Log == nil {
		h.Log = slog.Default()
	}
	states.OnState(notify.StateKindLinkedRole, h.callback)
	return h
}

// Mount registers the start route, rate-limited per IP.
func Mount(mux *http.ServeMux, h *Handler, trustedHops int) {
	limit := httpx.RateLimitPer(startsPerMinute, time.Minute, trustedHops)
	mux.Handle("GET "+StartPath, limit(http.HandlerFunc(h.start)))
}

func (h *Handler) accountRedirect(w http.ResponseWriter, r *http.Request, outcome string) {
	http.Redirect(w, r, h.PublicBaseURL+accountReturn+outcome+youAnchor, http.StatusFound)
}

// start sends a signed-in member to Discord's consent page under a fresh state bound to the
// account, and a signed-out visitor through sign-in first (the account page then resumes).
// Every redirect target is built here from configured bases; nothing from the request is
// echoed into a Location.
func (h *Handler) start(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() {
		http.Redirect(w, r, h.PublicBaseURL+"/login?next="+url.QueryEscape(signedOutNext), http.StatusFound)
		return
	}
	if actor.IsDevice() {
		h.accountRedirect(w, r, outcomeFailed)
		return
	}
	state, err := h.States.StartState(r.Context(), notify.StateKindLinkedRole, actor.UserID, nil)
	if err != nil {
		h.Log.Error("linkedrole", "id", httpx.RequestIDFrom(r.Context()), "op", "state", "err", err)
		h.accountRedirect(w, r, outcomeFailed)
		return
	}
	http.Redirect(w, r, h.Syncer.Discord.AuthorizeURL(state), http.StatusFound)
}

// callback finishes at the shared callback once the state row was taken (single use). It
// fails unless the session is the account that started it, inside the state's ten minutes.
func (h *Handler) callback(w http.ResponseWriter, r *http.Request, cs notify.ConnectState) {
	httpx.CachePrivateNoStore(w)
	actor := auth.ActorFrom(r.Context())
	if !actor.Signed() || actor.IsDevice() || actor.UserID != cs.UserID || cs.Expired {
		h.accountRedirect(w, r, outcomeFailed)
		return
	}
	query := r.URL.Query()
	if discordError := query.Get("error"); discordError != "" {
		outcome := outcomeFailed
		if discordError == oauthAccessDenied {
			outcome = outcomeDenied
		}
		h.accountRedirect(w, r, outcome)
		return
	}
	code := query.Get("code")
	if code == "" {
		h.accountRedirect(w, r, outcomeDenied)
		return
	}
	h.accountRedirect(w, r, h.finish(r, cs.UserID, code))
}

// finish exchanges the code, reads the Discord account and links it, returning the outcome.
func (h *Handler) finish(r *http.Request, userID int64, code string) string {
	ctx := r.Context()
	t, err := h.Syncer.Discord.Exchange(ctx, code)
	if err != nil {
		h.Log.Warn("linkedrole", "id", httpx.RequestIDFrom(ctx), "op", "exchange", "err", err)
		return outcomeFailed
	}
	discordUserID, err := h.Syncer.Discord.UserID(ctx, t.Access)
	if err != nil {
		h.Log.Warn("linkedrole", "id", httpx.RequestIDFrom(ctx), "op", "identify", "err", err)
		return outcomeFailed
	}
	if err := h.Syncer.Link(ctx, userID, discordUserID, t); err != nil {
		h.Log.Error("linkedrole", "id", httpx.RequestIDFrom(ctx), "op", "link", "err", err)
		return outcomeFailed
	}
	h.Log.Info("linkedrole", "op", "linked", "user_id", userID)
	return outcomeLinked
}
