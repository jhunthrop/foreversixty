package server

import (
	"net/http"

	"github.com/PLACEHOLDER/forever/api/internal/httpx"
)

type Deps struct {
	Version string
}

func NewRouter(d Deps) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"version": d.Version})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
	})
	return mux
}
