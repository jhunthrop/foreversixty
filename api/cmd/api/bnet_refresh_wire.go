package main

import (
	"context"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bnetimport"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// wireBnetRefresh turns on the automatic Battle.net refresh when the character import is
// on (the importer is a *bnetimport.Service): the sealed sign-in token, the refresh on a
// signed-in visit, and the hourly in-process loop. Run it after wireActivity, which gives
// the service its auto-sim. Without NOTIFY_ENCRYPTION_KEY no token is kept and every visit
// refresh reads public profiles.
func wireBnetRefresh(ctx context.Context, cfg config.Config, accounts *auth.Service, authenticator *auth.Authenticator,
	log *slog.Logger) {
	svc, ok := accounts.Importer.(*bnetimport.Service)
	if !ok {
		return
	}
	if len(cfg.NotifyEncryptionKey) == 0 {
		log.Warn("bnetimport", "state", "no token sealer", "effect", "visit refreshes read public profiles only")
	} else if sealer, err := notify.NewSealer(cfg.NotifyEncryptionKey); err != nil {
		log.Error("bnetimport", "state", "no token sealer", "err", err)
	} else {
		svc.Sealer = sealer
	}
	authenticator.Visits = svc
	go (&bnetimport.RefreshLoop{Service: svc}).Run(ctx)
	wireDirectory(ctx, cfg, svc, log)
}
