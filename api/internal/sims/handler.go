package sims

import (
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Service serves the simulator routes.
type Service struct {
	Store *Store
	// EngineVersion is the pinned engine build this deployment runs.
	EngineVersion string
	Log           *slog.Logger
}

// Mount registers every simulator route. Task 4 fills it in.
func Mount(mux *http.ServeMux, s *Service) {}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("sims", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}
