package subscribe

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PLACEHOLDER/forever/api/internal/httpx"
)

func Mount(mux *http.ServeMux, s *Service) {
	limited := httpx.RateLimit(10)
	mux.Handle("POST /v1/subscribe", limited(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096)).Decode(&body); err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with an email field", map[string]string{"email": "required"})
			return
		}
		err := s.Subscribe(r.Context(), body.Email)
		switch {
		case errors.Is(err, ErrInvalidEmail):
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "enter a valid email address", map[string]string{"email": "invalid"})
		case err != nil:
			httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not subscribe right now", nil)
		default:
			httpx.WriteOK(w, r, http.StatusAccepted, map[string]string{"status": "check your email"})
		}
	})))
	redirect := func(w http.ResponseWriter, r *http.Request, ok bool, good string) {
		target := s.PublicBaseURL + "/subscribe-invalid"
		if ok {
			target = s.PublicBaseURL + good
		}
		http.Redirect(w, r, target, http.StatusFound)
	}
	mux.HandleFunc("GET /v1/subscribe/confirm", func(w http.ResponseWriter, r *http.Request) {
		ok, err := s.Confirm(r.Context(), r.URL.Query().Get("token"))
		redirect(w, r, err == nil && ok, "/subscribed")
	})
	mux.HandleFunc("GET /v1/subscribe/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		ok, err := s.Unsubscribe(r.Context(), r.URL.Query().Get("token"))
		redirect(w, r, err == nil && ok, "/unsubscribed")
	})
}
