package server

import (
	"log/slog"
	"net/http"

	"github.com/PLACEHOLDER/forever/api/internal/httpx"
	"github.com/PLACEHOLDER/forever/api/internal/subscribe"
)

type Deps struct {
	Version       string
	Log           *slog.Logger
	AllowedOrigin string
	Subscribe     *subscribe.Service
}

func NewRouter(d Deps) http.Handler {
	if d.Log == nil {
		d.Log = slog.Default()
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"version": d.Version})
	})
	if d.Subscribe != nil {
		subscribe.Mount(mux, d.Subscribe)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
	})
	return httpx.Chain(mux, httpx.RequestID(), httpx.Recover(d.Log), httpx.Logger(d.Log), httpx.CORS(d.AllowedOrigin), httpx.RateLimit(120))
}
