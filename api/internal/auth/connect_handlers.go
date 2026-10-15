package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/textx"
)

type connectStartRequest struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// connectStart mints a connect for a companion that has no account yet.
func (s *Service) connectStart(w http.ResponseWriter, r *http.Request) {
	var in connectStartRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a name and platform", nil)
		return
	}
	c, err := s.Store.CreateConnect(r.Context(),
		textx.Trim(in.Name, maxDeviceName), textx.Trim(in.Platform, maxDevicePlatform))
	if err != nil {
		s.fail(w, r, "connect", err, "could not start a connection just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusCreated, map[string]any{
		"connect_id": c.ConnectID,
		"user_code":  c.UserCode,
		"url":        s.PublicBaseURL + "/connect?code=" + c.UserCode,
		"expires_in": int(ConnectTTL.Seconds()),
		"interval":   ConnectPollInterval,
	})
}

// connectPoll is the companion's poll. The id is the only secret, and every
// miss looks the same: a 404 not_found envelope.
func (s *Service) connectPoll(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	id := r.PathValue("connect_id")
	token, hash := NewDeviceToken()
	st, err := s.Store.PollConnect(r.Context(), id, hash)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such connection", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "connect poll", err, "could not check that connection just now")
		return
	}
	switch st.Status {
	case "approved":
		// The only time the token is ever readable.
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{
			"status": "approved", "device_id": st.Device.ID, "token": token, "account_label": st.AccountLabel,
		})
	case "used":
		httpx.WriteOK(w, r, http.StatusGone, map[string]string{"status": "used"})
	default:
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": st.Status})
	}
}

// connectCode reads the code a person typed or followed, tolerant of case,
// hyphens and spaces. A malformed one is as good as unknown.
func connectCode(w http.ResponseWriter, r *http.Request, raw string) (string, bool) {
	code, ok := NormalizeUserCode(raw)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "that code is not valid or has expired", nil)
	}
	return code, ok
}

func (s *Service) connectLookup(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivate(w)
	code, ok := connectCode(w, r, r.URL.Query().Get("code"))
	if !ok {
		return
	}
	in, err := s.Store.LookupConnect(r.Context(), code)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "that code is not valid or has expired", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "connect lookup", err, "could not look that code up just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{
		"name": in.Name, "platform": in.Platform, "expires_in": in.ExpiresIn,
	})
}

type connectDecision struct {
	Code string `json:"code"`
}

func (s *Service) connectDecide(w http.ResponseWriter, r *http.Request, approve bool) {
	var in connectDecision
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with a code", nil)
		return
	}
	code, ok := connectCode(w, r, strings.TrimSpace(in.Code))
	if !ok {
		return
	}
	info, err := s.Store.DecideConnect(r.Context(), code, ActorFrom(r.Context()).UserID, approve)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "that code is not valid or has expired", nil)
	case errors.Is(err, ErrConnectDecided):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that connection was already answered", nil)
	case err != nil:
		s.fail(w, r, "connect decide", err, "could not answer that connection just now")
	case approve:
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"name": info.Name, "platform": info.Platform})
	default:
		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *Service) connectApprove(w http.ResponseWriter, r *http.Request) { s.connectDecide(w, r, true) }
func (s *Service) connectDeny(w http.ResponseWriter, r *http.Request)    { s.connectDecide(w, r, false) }
