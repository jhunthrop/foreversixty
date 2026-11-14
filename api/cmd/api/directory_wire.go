// api/cmd/api/directory_wire.go
package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/bnetimport"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/directory"
)

// directoryJobs are the character directory's Cloud Run job commands, rows of extraJobs.
var directoryJobs = []jobCommand{{directory.JobCommand, runDirectoryRefresh}}

// directoryGame is the only game the directory reads: Forever's realmless profiles.
const directoryGame = "forever"

// directoryEnabled reports whether the directory may call Blizzard: Battle.net configured,
// the Forever game, and the CHARACTER_DIRECTORY switch not off.
func directoryEnabled(cfg config.Config) bool {
	return cfg.BattleNetConfigured() && cfg.CharacterDirectory && cfg.BnetProfileGame == directoryGame
}

// wireDirectory starts the directory's hourly refresh in this process when it is enabled,
// sharing the Battle.net import's pool and client (and so its app token).
func wireDirectory(ctx context.Context, cfg config.Config, importer *bnetimport.Service, log *slog.Logger) {
	if !directoryEnabled(cfg) {
		log.Warn("directory", "state", "off", "effect", "the character directory is not refreshed")
		return
	}
	svc := &directory.Service{Pool: importer.Pool, Client: importer.Client, Regions: cfg.BnetRegions, Log: log}
	go (&directory.Loop{Service: svc}).Run(ctx)
}

// runDirectoryRefresh is the directory-refresh job: one pass, its counts logged.
func runDirectoryRefresh(ctx context.Context, log *slog.Logger, _ []string) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	if !directoryEnabled(cfg) {
		return fmt.Errorf("%s: the directory is off (Battle.net, BNET_PROFILE_GAME or CHARACTER_DIRECTORY)", directory.JobCommand)
	}
	svc := &directory.Service{Pool: pool, Client: bnetClient(cfg, log), Regions: cfg.BnetRegions, Log: log}
	res, err := svc.Pass(ctx)
	directory.LogPass(svc, res)
	return err
}
