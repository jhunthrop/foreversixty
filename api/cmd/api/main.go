package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/PLACEHOLDER/forever/api/internal/config"
	"github.com/PLACEHOLDER/forever/api/internal/server"
)

var version = "dev" // set with -ldflags "-X main.version=<git sha>"

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		log.Error("startup", "err", err)
		os.Exit(1)
	}
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           server.NewRouter(server.Deps{Version: version, Log: log, AllowedOrigin: cfg.PublicBaseURL}),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Info("listening", "port", cfg.Port, "version", version)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("serve", "err", err)
		os.Exit(1)
	}
}
