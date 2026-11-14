// api/internal/directory/optout_handler.go
package directory

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// optOutPerMinute bounds each opt-out route per IP: a code start can cost a Blizzard call.
const optOutPerMinute = 10

// optOutBody is every opt-out route's request: the character's region and full name.
type optOutBody struct {
	Region string `json:"region"`
	Name   string `json:"name"`
}

// MountOptOuts serves the "Remove my character" page's three calls.
func MountOptOuts(mux *http.ServeMux, o *OptOuts, log *slog.Logger, trustedProxyHops int) {
	limited := httpx.RateLimit(optOutPerMinute, trustedProxyHops)
	h := optOutHandler{o: o, log: log}
	mux.Handle("POST /v1/directory/opt-out/code", limited(http.HandlerFunc(h.code)))
	mux.Handle("POST /v1/directory/opt-out/status", limited(http.HandlerFunc(h.status)))
	mux.Handle("POST /v1/directory/opt-out/account", limited(auth.RequireSession(h.account)))
}

type optOutHandler struct {
	o   *OptOuts
	log *slog.Logger
}

func (h optOutHandler) read(w http.ResponseWriter, r *http.Request) (optOutBody, bool) {
	httpx.CachePrivateNoStore(w)
	var body optOutBody
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1024)).Decode(&body); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with region and name", nil)
		return optOutBody{}, false
	}
	return body, true
}

func (h optOutHandler) code(w http.ResponseWriter, r *http.Request) {
	body, ok := h.read(w, r)
	if !ok {
		return
	}
	code, err := h.o.StartCode(r.Context(), body.Region, body.Name)
	if err != nil {
		h.fail(w, r, "code", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, code)
}

func (h optOutHandler) status(w http.ResponseWriter, r *http.Request) {
	body, ok := h.read(w, r)
	if !ok {
		return
	}
	state, err := h.o.Status(r.Context(), body.Region, body.Name)
	if err != nil {
		h.fail(w, r, "status", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"state": state})
}

func (h optOutHandler) account(w http.ResponseWriter, r *http.Request) {
	body, ok := h.read(w, r)
	if !ok {
		return
	}
	if err := h.o.RemoveOwned(r.Context(), auth.ActorFrom(r.Context()).UserID, body.Region, body.Name); err != nil {
		h.fail(w, r, "account", err)
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"state": OptOutRemoved})
}

// fail answers an opt-out error: the refusals by name, anything else as internal (logged by
// request id, never the character's name).
func (h optOutHandler) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	switch {
	case errors.Is(err, ErrBadCharacter):
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "enter a region and the character's full name",
			map[string]string{"name": "invalid"})
	case errors.Is(err, ErrNoSuchCharacter):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no public character has that name", nil)
	case errors.Is(err, ErrNotYours):
		httpx.WriteError(w, r, http.StatusForbidden, "not_yours", "that character is not on your Battle.net account", nil)
	default:
		h.log.Error("directory", "id", httpx.RequestIDFrom(r.Context()), "op", "opt_out_"+op, "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not do that just now", nil)
	}
}
