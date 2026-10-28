// api/cmd/api/rebuild_preparation.go
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
)

// runRebuildPreparation is the on-demand Cloud Run job that re-rates every stored fight under
// the current model version and refolds the preparation digests. It never migrates (start
// refuses to in a job), so it runs from the deployed sha only. Idempotent; -dry-run counts
// without writing; a crashed run resumes with the -after cursor its last progress line logged.
func runRebuildPreparation(ctx context.Context, log *slog.Logger, args []string) error {
	opts, err := parseRebuildFlags(args, io.Discard)
	if err != nil {
		return err
	}
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("%s needs R2 credentials", rating.RebuildPreparationJobCommand)
	}
	opts.Progress = logRebuildProgress(log)
	counts, err := rating.RebuildPreparationWith(ctx, rating.BackfillDeps{
		Store: &rating.Store{Pool: pool, Log: log}, Summaries: client, Log: log,
	}, opts)
	logRebuildCounts(log, "done", counts, opts.DryRun)
	if err != nil {
		return err
	}
	if counts.Failed > 0 {
		return fmt.Errorf("%d fights could not be rebuilt (each logged above)", counts.Failed)
	}
	return nil
}

// parseRebuildFlags reads -dry-run, -batch and -after (report:fightIndex).
func parseRebuildFlags(args []string, out io.Writer) (rating.RebuildOptions, error) {
	fs := flag.NewFlagSet(rating.RebuildPreparationJobCommand, flag.ContinueOnError)
	fs.SetOutput(out)
	dry := fs.Bool("dry-run", false, "count what would be folded and re-rated; write nothing")
	batch := fs.Int("batch", rating.RebuildBatchSize, "fights per batch")
	after := fs.String("after", "", "resume the fold past this cursor (report:fightIndex)")
	if err := fs.Parse(args); err != nil {
		return rating.RebuildOptions{}, err
	}
	if *batch <= 0 {
		return rating.RebuildOptions{}, fmt.Errorf("-batch must be positive, got %d", *batch)
	}
	opts := rating.RebuildOptions{DryRun: *dry, BatchSize: *batch}
	if *after != "" {
		cursor, err := parseFightCursor(*after)
		if err != nil {
			return rating.RebuildOptions{}, err
		}
		opts.After = &cursor
	}
	return opts, nil
}

func parseFightCursor(s string) (rating.FightCursor, error) {
	i := strings.LastIndex(s, ":")
	if i <= 0 {
		return rating.FightCursor{}, fmt.Errorf("-after %q: want report:fightIndex", s)
	}
	index, err := strconv.Atoi(s[i+1:])
	if err != nil {
		return rating.FightCursor{}, fmt.Errorf("-after %q: fight index: %w", s, err)
	}
	return rating.FightCursor{ReportID: s[:i], FightIndex: index}, nil
}

func logRebuildProgress(log *slog.Logger) func(string, rating.RebuildCounts, rating.FightCursor) {
	return func(phase string, c rating.RebuildCounts, cur rating.FightCursor) {
		logRebuildCounts(log, phase, c, false, "after", fmt.Sprintf("%s:%d", cur.ReportID, cur.FightIndex))
	}
}

func logRebuildCounts(log *slog.Logger, phase string, c rating.RebuildCounts, dry bool, extra ...any) {
	log.Info(rating.RebuildPreparationJobCommand, append([]any{"phase", phase, "dry_run", dry,
		"folded", c.Folded, "rerated", c.Rerated, "rolled_up", c.RolledUp, "failed", c.Failed}, extra...)...)
}
