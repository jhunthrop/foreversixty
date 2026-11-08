package apptesters

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// The two routes.
const (
	SignupPath = "/v1/app-testers"
	AdminPath  = "/v1/admin/app-testers"
)

// SignupsPerIPPerHour is the public form's budget per address, on top of the router's.
const SignupsPerIPPerHour = 5

const maxBody = 1 << 10

type testers interface {
	Add(ctx context.Context, email, platform string) (bool, error)
	List(ctx context.Context) ([]Tester, error)
}

// Service serves the signup and the owner's list.
type Service struct {
	Testers          testers
	Log              *slog.Logger
	TrustedProxyHops int
}

// Mount registers the public signup (no auth; the router's Authenticator still
// holds a signed-in browser to the double-submit CSRF check) and the owner's list.
func Mount(mux *http.ServeMux, s *Service) {
	limited := httpx.RateLimitPer(SignupsPerIPPerHour, time.Hour, s.TrustedProxyHops)
	mux.Handle("POST "+SignupPath, limited(http.HandlerFunc(s.signup)))
	mux.HandleFunc("GET "+AdminPath, auth.RequireAdmin(s.list))
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error) {
	s.logger().Error("apptesters", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not save that just now; try again", nil)
}

type signupInput struct {
	Email    string `json:"email"`
	Platform string `json:"platform"`
}

// fields is every reason a signup is refused, by field, and the normalized address.
func (in signupInput) check() (string, map[string]string) {
	fields := map[string]string{}
	email, err := NormalizeEmail(in.Email)
	if err != nil {
		fields["email"] = "enter one email address, like you@example.com"
	}
	if !ValidPlatform(in.Platform) {
		fields["platform"] = "pick iPhone or Android"
	}
	return email, fields
}

func (s *Service) signup(w http.ResponseWriter, r *http.Request) {
	httpx.CachePrivateNoStore(w)
	var in signupInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with an email and a platform", nil)
		return
	}
	email, fields := in.check()
	if len(fields) > 0 {
		httpx.WriteError(w, r, http.StatusUnprocessableEntity, "invalid", "that signup cannot be saved", fields)
		return
	}
	created, err := s.Testers.Add(r.Context(), email, in.Platform)
	if err != nil {
		s.fail(w, r, "signup", err)
		return
	}
	if created {
		httpx.WriteOK(w, r, http.StatusCreated, map[string]string{"status": "created"})
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "exists"})
}

func (s *Service) list(w http.ResponseWriter, r *http.Request) {
	rows, err := s.Testers.List(r.Context())
	if err != nil {
		s.fail(w, r, "list", err)
		return
	}
	if rows == nil {
		rows = []Tester{}
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"testers": rows})
}
