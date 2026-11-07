package main

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/activity"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/bnetimport"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// activityWiring is what wireActivity reads.
type activityWiring struct {
	cfg      config.Config
	pool     *pgxpool.Pool
	trees    *trees.Data
	auth     *auth.Store
	accounts *auth.Service
	log      *slog.Logger
}

// curatedOrEmpty reads the feed's curated files; without them the feed still serves, with
// no boards and no first-kill rows.
func curatedOrEmpty(cfg config.Config, log *slog.Logger) activity.Curated {
	c, err := activity.LoadCurated(cfg.CuratedDir)
	if err != nil {
		log.Warn("activity", "err", err, "effect", "no spec boards, no first raid kills in the feed")
	}
	return c
}

// wireActivity builds the site feed (design/specs/2026-11-07-home-b.md) into deps: its
// routes, its writers on the report, guild raid, addon and Battle.net paths, and the hourly
// loop. Run it after the Ingest, Sims and Guilds services exist.
func wireActivity(ctx context.Context, w activityWiring, deps *server.Deps) {
	store := &activity.Store{Pool: w.pool}
	curated := curatedOrEmpty(w.cfg, w.log)
	writer := &activity.ReportWriter{Store: store, RaidZones: curated.RaidZones, Log: w.log}
	if deps.Reports != nil {
		deps.Reports.Completed = reports.CompletedAll{deps.Reports.Completed, writer}
	}
	if deps.Ingest != nil {
		deps.Ingest.Completed = reports.CompletedAll{deps.Ingest.Completed, writer}
	}
	if deps.Guilds != nil {
		deps.Guilds.Activity = &activity.GuildRaidWriter{Store: store, Log: w.log}
	}
	home := &activity.Home{Store: store, Sections: w.cfg.App.HomeSections}
	if deps.Sims != nil && deps.Sims.Jobs != nil {
		home.Imports = deps.Sims
		if deps.Addon != nil {
			deps.Addon.ImportSims = deps.Sims
		}
		if importer, ok := w.accounts.Importer.(*bnetimport.Service); ok {
			importer.ImportSims = deps.Sims
		}
	}
	deps.Activity = &activity.Service{Store: store, Guilds: w.auth, Home: home, Log: w.log}
	loop := &activity.Loop{
		Store: store, DataDir: w.cfg.TreeDataDir, Curated: curated, Log: w.log,
		Build: bis.SourceBuild(w.cfg.TreeDataDir, activeVersion(w.trees)),
	}
	go loop.Run(ctx)
}

// refreshImportSims is the nightly refresh's auto-sim: a sims service of the job's own that
// can start the sim job, or nil when it cannot be reached.
func refreshImportSims(ctx context.Context, cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) bnetimport.ImportSimmer {
	runner, err := jobs.NewCloudRun(ctx, cfg.SimJobProject, cfg.SimJobRegion, cfg.SimJobName)
	if err != nil {
		log.Warn(bnetimport.RefreshJobCommand, "err", err, "effect", "no auto-sim of refreshed gear")
		return nil
	}
	return &sims.Service{Store: &sims.Store{Pool: pool}, Jobs: runner, EngineVersion: enginever.Version, Log: log}
}
