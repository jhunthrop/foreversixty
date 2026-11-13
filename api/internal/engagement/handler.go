package engagement

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// AdminPath is the owner's figures (/admin/activity).
const AdminPath = "/v1/admin/engagement"

type reporter interface {
	Report(ctx context.Context, now time.Time) (Report, error)
}

// Service serves the owner's figures.
type Service struct {
	Reports reporter
	// Now is the clock; nil is time.Now.
	Now func() time.Time
	Log *slog.Logger
}

// Mount registers the owner's read.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET "+AdminPath, auth.RequireAdmin(s.report))
}

func (s *Service) report(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if s.Now != nil {
		now = s.Now()
	}
	rep, err := s.Reports.Report(r.Context(), now)
	if err != nil {
		log := s.Log
		if log == nil {
			log = slog.Default()
		}
		log.Error("engagement", "id", httpx.RequestIDFrom(r.Context()), "op", "report", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read the activity figures just now", nil)
		return
	}
	httpx.CachePrivateNoStore(w)
	httpx.WriteOK(w, r, http.StatusOK, rep)
}
