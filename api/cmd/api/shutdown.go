package main

import (
	"context"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/server"
)

// drainCompletions waits, until ctx's deadline (the stop's shared budget), for the report
// completion hooks the last requests started (a raid match each), so a deploy does not drop them.
func drainCompletions(ctx context.Context, deps server.Deps, log *slog.Logger) {
	if deps.Reports != nil {
		if err := deps.Reports.DrainCompletions(ctx); err != nil {
			log.Error("shutdown", "step", "report_hooks", "err", err)
		}
	}
	if deps.Ingest != nil {
		if err := deps.Ingest.DrainCompletions(ctx); err != nil {
			log.Error("shutdown", "step", "ingest_hooks", "err", err)
		}
	}
}

// drainDiscordBot waits, within the stop's shared budget, for the deferred
// answers the last interactions started.
func drainDiscordBot(ctx context.Context, deps server.Deps, log *slog.Logger) {
	if deps.DiscordBot == nil {
		return
	}
	if err := deps.DiscordBot.Drain(ctx); err != nil {
		log.Error("shutdown", "step", "discord_answers", "err", err)
	}
}
