// api/internal/account/handler.go
package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Path is the deletion's route; GET and PATCH on it belong to the auth package.
const Path = "/v1/me"

// deletesPerHour is each account's budget: a typo or two, never a loop.
const deletesPerHour = 5

// maxBody bounds the body: one short typed string.
const maxBody = 1 << 10

// The answers' words. The guild sentences follow the guild-owner decision recorded in
// docs/superpowers/specs/2026-11-07-account-delete-design.md ("Guild owner decision").
const (
	msgInvalid      = "send what you typed to confirm"
	msgMismatch     = "Type your BattleTag or delete to confirm."
	fieldConfirm    = "your BattleTag or delete"
	msgStaff        = "Staff accounts are closed by support. Email help@foreversixty.gg."
	msgBilling      = "Could not cancel your subscription just now, so nothing was deleted. Try again in a few minutes."
	msgFailed       = "could not delete your account just now"
	msgGone         = "this account no longer exists"
	claimSentence   = "You hold the claim on %s. Release the claim on your account page first, so the in-game guild master can claim it again."
	planSentence    = "You pay for the Guild plan of %s, so deleting your account would end it for the whole guild right away, with no refund. Have another officer take over billing on the guild page, or cancel the plan from Manage billing, before you delete."
	guildNamesField = "guilds"
)

// SessionClearer expires the browser's session cookies; *auth.Authenticator satisfies it.
type SessionClearer interface {
	ClearSession(w http.ResponseWriter)
}

// Handler serves DELETE /v1/me.
type Handler struct {
	Service  *Service
	Sessions SessionClearer
}

// Mount registers the route behind a session (a browser or a signed-in app; never a
// companion) and the per-account budget.
func Mount(mux *http.ServeMux, h *Handler) {
	limit := httpx.RateLimitByKey(deletesPerHour, time.Hour, auth.AccountRateLimitKey)
	mux.Handle("DELETE "+Path, limit(auth.RequireSession(h.remove)))
}

type deleteBody struct {
	Confirm string `json:"confirm"`
}

func (h *Handler) remove(w http.ResponseWriter, r *http.Request) {
	var body deleteBody
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&body); err != nil ||
		strings.TrimSpace(body.Confirm) == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msgInvalid, map[string]string{"confirm": fieldConfirm})
		return
	}
	at, err := h.Service.Delete(r.Context(), auth.ActorFrom(r.Context()).UserID, body.Confirm)
	if err != nil {
		h.refuse(w, r, err)
		return
	}
	h.Sessions.ClearSession(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]time.Time{"deleted_at": at})
}

// refuse maps a deletion error to its answer.
func (h *Handler) refuse(w http.ResponseWriter, r *http.Request, err error) {
	var obligation *ObligationError
	switch {
	case errors.Is(err, ErrConfirmMismatch):
		httpx.WriteError(w, r, http.StatusBadRequest, "confirm_mismatch", msgMismatch, map[string]string{"confirm": fieldConfirm})
	case errors.As(err, &obligation):
		httpx.WriteError(w, r, http.StatusConflict, "guild_obligation", ObligationMessage(obligation.Guilds),
			map[string]string{guildNamesField: guildNames(obligation.Guilds, func(Obligation) bool { return true }, ", ")})
	case errors.Is(err, ErrStaffAccount):
		httpx.WriteError(w, r, http.StatusConflict, "staff_account", msgStaff, nil)
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized", msgGone, nil)
	case errors.Is(err, ErrBilling):
		h.Service.logger().Error("account", "op", "delete", "stage", "billing", "user_id", auth.ActorFrom(r.Context()).UserID, "err", err)
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "billing_unavailable", msgBilling, nil)
	default:
		h.Service.logger().Error("account", "op", "delete", "user_id", auth.ActorFrom(r.Context()).UserID, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", msgFailed, nil)
	}
}

// ObligationMessage is the refusal's sentence: one for the claims held, one for the plans
// paid, each naming its guilds.
func ObligationMessage(guilds []Obligation) string {
	var parts []string
	if names := guildNames(guilds, func(o Obligation) bool { return o.Claim }, ""); names != "" {
		parts = append(parts, fmt.Sprintf(claimSentence, names))
	}
	if names := guildNames(guilds, func(o Obligation) bool { return o.GuildPlan }, ""); names != "" {
		parts = append(parts, fmt.Sprintf(planSentence, names))
	}
	return strings.Join(parts, " ")
}

// guildNames lists the names of the guilds keep picks: joined by sep, or as prose ("A, B
// and C") when sep is empty.
func guildNames(guilds []Obligation, keep func(Obligation) bool, sep string) string {
	var names []string
	for _, g := range guilds {
		if keep(g) {
			names = append(names, g.GuildName)
		}
	}
	if sep != "" || len(names) < 2 {
		return strings.Join(names, sep)
	}
	return strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
}
