package main

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/activity"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/parse"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// newParseDeps builds the parse-report job's dependencies on the job's
// own pool: the completion hook is a guilds service of its own, the
// way the job's Discord poster is a notify service of its own.
func newParseDeps(cfg config.Config, pool *pgxpool.Pool, client parse.Objects, treeData *trees.Data, log *slog.Logger) parse.Deps {
	deps := parseDeps(cfg, pool, client, treeData, newJobGuilds(cfg, pool, log), log)
	if notifier := newJobNotifier(context.Background(), cfg, pool, log); notifier != nil {
		deps.Notify = notifier
	}
	return deps
}

// newBackfillDeps is the parse job's wiring with every channel out: no notifier, and a
// guilds hook with no event router, so a backfilled night (docs/runbooks/report-backfill.md)
// still gets its zone, guild, raid match and feed line and tells nobody about a night that
// is already over.
func newBackfillDeps(cfg config.Config, pool *pgxpool.Pool, client parse.Objects, treeData *trees.Data, log *slog.Logger) parse.Deps {
	silent := &guilds.Service{Store: &guilds.Store{Pool: pool}, PublicBaseURL: cfg.PublicBaseURL, Log: log}
	return parseDeps(cfg, pool, client, treeData, silent, log)
}

// parseDeps is the wiring the parse job and the backfill share, with the guilds hook the
// caller chose.
func parseDeps(cfg config.Config, pool *pgxpool.Pool, client parse.Objects, treeData *trees.Data,
	guildHook reports.Completed, log *slog.Logger) parse.Deps {
	return parse.Deps{
		Reports: &reports.Store{Pool: pool}, Objects: client, Log: log,
		Rank: &rankings.Store{Pool: pool, Specs: inferrer(treeData)},
		Completed: reports.CompletedAll{guildHook, &activity.ReportWriter{
			Store: &activity.Store{Pool: pool}, RaidZones: curatedOrEmpty(cfg, log).RaidZones, Log: log}},
	}
}
