package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/subscribe"
)

var version = "dev" // set with -ldflags "-X main.version=<git sha>"

// resendCooldown and maxSendsPerHour guard against confirmation-resend
// abuse: an attacker (or a confused client) hammering POST /v1/subscribe for
// the same address can otherwise mail-bomb that address, and hammering
// distinct addresses can otherwise exhaust the mail provider's quota.
const (
	resendCooldown  = 15 * time.Minute
	maxSendsPerHour = 300

	shutdownTimeout = 10 * time.Second
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		log.Error("startup", "err", err)
		os.Exit(1)
	}
	if err := db.Migrate(cfg.MigrateDatabaseURL); err != nil {
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
		Store:           &subscribe.Store{Pool: pool},
		Mailer:          mail.NewResend(cfg.ResendAPIKey, cfg.MailFrom, nil),
		PublicBaseURL:   cfg.PublicBaseURL,
		APIBaseURL:      cfg.APIBaseURL,
		Logger:          log,
		ResendCooldown:  resendCooldown,
		MaxSendsPerHour: maxSendsPerHour,
	}
	srv := &http.Server{
		Addr: ":" + cfg.Port,
		Handler: server.NewRouter(server.Deps{
			Version:          version,
			Log:              log,
			AllowedOrigin:    cfg.PublicBaseURL,
			Subscribe:        svc,
			TrustedProxyHops: cfg.TrustedProxyHops,
		}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serveErr := make(chan error, 1)
	go func() {
		log.Info("listening", "port", cfg.Port, "version", version)
		serveErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-serveErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("serve", "err", err)
			os.Exit(1)
		}
	case <-ctx.Done():
		log.Info("shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error("shutdown", "err", err)
		}
		<-serveErr // wait for the listener goroutine to actually return
	}
	svc.Wait()
	log.Info("stopped")
}
