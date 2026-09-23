package subscribe

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

func Mount(mux *http.ServeMux, s *Service, log *slog.Logger, trustedProxyHops int) {
	limited := httpx.RateLimit(10, trustedProxyHops)
	mux.Handle("POST /v1/subscribe", limited(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096)).Decode(&body); err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with an email field", map[string]string{"email": "required"})
			return
		}
		// Subscribe only touches the store on the request path; mail is sent
		// afterward (see Service.Subscribe), so an error here is always an
		// unexpected store error, never a mail failure.
		err := s.Subscribe(r.Context(), body.Email)
		switch {
		case errors.Is(err, ErrInvalidEmail):
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "enter a valid email address", map[string]string{"email": "invalid"})
		case err != nil:
			log.Error("subscribe", "id", httpx.RequestIDFrom(r.Context()), "op", "subscribe", "err", err)
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
		httpx.CachePrivate(w)
		ok, err := s.Confirm(r.Context(), r.URL.Query().Get("token"))
		if err != nil {
			log.Error("subscribe", "id", httpx.RequestIDFrom(r.Context()), "op", "confirm", "err", err)
		}
		redirect(w, r, err == nil && ok, "/subscribed")
	})
	mux.HandleFunc("GET /v1/subscribe/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		httpx.CachePrivate(w)
		ok, err := s.Unsubscribe(r.Context(), r.URL.Query().Get("token"))
		if err != nil {
			log.Error("subscribe", "id", httpx.RequestIDFrom(r.Context()), "op", "unsubscribe", "err", err)
		}
		redirect(w, r, err == nil && ok, "/unsubscribed")
	})
}
