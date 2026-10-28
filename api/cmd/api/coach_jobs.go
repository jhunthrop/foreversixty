package main

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

// coachJobs are the coach's Cloud Run job commands, rows of extraJobs.
var coachJobs = []jobCommand{
	{coach.ReviewJobCommand, runCoachReview},
	{coach.SimRefJobCommand, runCoachSimRef},
	{coach.SamplesJobCommand, runCoachSamples},
}

// runCoachReview is the coach-review job: coach-review <report> <fight> <player_key>
// <account_id> <holder>. It loads the name dictionary once, then writes one review.
func runCoachReview(ctx context.Context, log *slog.Logger, args []string) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("%s needs R2 credentials", coach.ReviewJobCommand)
	}
	svc, err := newCoachService(ctx, jobWiring(cfg, pool, bucketOrNil(client), log))
	if err != nil {
		return err
	}
	if svc.Names, err = coachNames(cfg, coachTables(cfg, &sims.Store{Pool: pool})); err != nil {
		return err
	}
	return coach.RunReviewJob(ctx, svc, args)
}

// jobWiring is a job's coach wiring on its own pool: no server deps, its own stores.
func jobWiring(cfg config.Config, pool *pgxpool.Pool, bucket coach.Getter, log *slog.Logger) coachWiring {
	return coachWiring{
		cfg: cfg, pool: pool, reports: &reports.Store{Pool: pool},
		auth: &auth.Store{Pool: pool, Log: log}, entitle: &entitlements.Store{Pool: pool},
		sims: &sims.Store{Pool: pool}, summaries: bucket, log: log, newRunner: jobRunners(cfg),
	}
}

// jobRunners binds a job's runners: a job never queues a review (it is the review), so it
// binds no coach-review runner; a sim-reference miss still queues coach-simref.
func jobRunners(cfg config.Config) runnerFactory {
	return func(ctx context.Context, job string) (jobs.Runner, error) {
		if job == coach.ReviewJobCommand {
			return nil, nil
		}
		return jobs.NewCloudRun(ctx, cfg.SimJobProject, cfg.SimJobRegion, job)
	}
}

// coachNames is the name dictionary of the build in force now, with every encounter's boss.
func coachNames(cfg config.Config, tables *coach.FileTables) (*coach.NameDict, error) {
	build, err := tables.BuildAt(time.Now())
	if err != nil {
		return nil, fmt.Errorf("coach names: %w", err)
	}
	encounters, err := mechanics.All()
	if err != nil {
		return nil, err
	}
	return coach.LoadNames(filepath.Join(cfg.TreeDataDir, build), encounters)
}

// runCoachSimRef is the coach-simref job: coach-simref <engine> <build> <spec> <band>
// <duration bucket>. It sims the spec's ladder character and stores the reference.
func runCoachSimRef(ctx context.Context, log *slog.Logger, args []string) error {
	key, err := coach.ParseSimRefKey(args)
	if err != nil {
		return err
	}
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return fmt.Errorf("trees: %s: %w", cfg.TreeDataDir, err)
	}
	sims.SetGearData(treeData, cfg.TreeDataDir) // sims.BandPicks reads the item tables
	log.Info(coach.SimRefJobCommand, "spec", key.Spec, "band", key.Band, "build", key.Build)
	return coach.RunSimRef(ctx, coach.SimRefDeps{
		DB: pool, Engine: simEngine(log), Characters: ladderCharacters{dataDir: cfg.TreeDataDir},
		Tables: coachTables(cfg, &sims.Store{Pool: pool}),
	}, key)
}

// runCoachSamples is the nightly coach-samples job: coach-samples [YYYY-MM-DD], yesterday
// (UTC) when no day is named.
func runCoachSamples(ctx context.Context, log *slog.Logger, args []string) error {
	day, err := samplesDay(args, time.Now())
	if err != nil {
		return err
	}
	_, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	n, err := coach.SampleDay(ctx, pool, day, coach.DaySeed(day), log)
	if err != nil {
		return err
	}
	log.Info(coach.SamplesJobCommand, "day", day.Format(time.DateOnly), "sampled", n)
	return nil
}

// samplesDay is the day args name, else the UTC day before now.
func samplesDay(args []string, now time.Time) (time.Time, error) {
	switch len(args) {
	case 0:
		y, m, d := now.UTC().AddDate(0, 0, -1).Date()
		return time.Date(y, m, d, 0, 0, 0, 0, time.UTC), nil
	case 1:
		return coach.ParseSampleDay(args[0])
	}
	return time.Time{}, fmt.Errorf("usage: api %s [YYYY-MM-DD]", coach.SamplesJobCommand)
}
