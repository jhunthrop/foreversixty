package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/PLACEHOLDER/forever/api/internal/config"
	"github.com/PLACEHOLDER/forever/api/internal/db"
	"github.com/PLACEHOLDER/forever/api/internal/mail"
	"github.com/PLACEHOLDER/forever/api/internal/server"
	"github.com/PLACEHOLDER/forever/api/internal/subscribe"
)

var version = "dev" // set with -ldflags "-X main.version=<git sha>"

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		log.Error("startup", "err", err)
		os.Exit(1)
	}
	if err := db.Migrate(cfg.DatabaseURL); err != nil {
		log.Error("migrate", "err", err)
		os.Exit(1)
	}
	pool, err := db.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Error("db", "err", err)
		os.Exit(1)
	}
	defer pool.Close()
	svc := &subscribe.Service{
		Store:         &subscribe.Store{Pool: pool},
		Mailer:        mail.NewResend(cfg.ResendAPIKey, "Forever Sixty <hello@foreversixty.gg>", nil),
		PublicBaseURL: cfg.PublicBaseURL,
		APIBaseURL:    cfg.APIBaseURL,
	}
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           server.NewRouter(server.Deps{Version: version, Log: log, AllowedOrigin: cfg.PublicBaseURL, Subscribe: svc}),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Info("listening", "port", cfg.Port, "version", version)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("serve", "err", err)
		os.Exit(1)
	}
}
