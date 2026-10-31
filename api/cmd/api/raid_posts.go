// api/cmd/api/raid_posts.go
//
// raid-posts is the Cloud Scheduler job that runs every ten minutes: it sends the raid posts
// and day-before reminders that have come due, then brings the bot's posts up to date, then
// (with APP_NOTIFICATIONS_ENABLED) publishes the time-based player events and prunes the old
// delivery log. It is the clock that works while the service is cold. It logs counts only.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// raidPostsDrainPasses is how many bot drain passes follow the schedule run.
const raidPostsDrainPasses = 5

// raidPostsBudget ends the run cleanly, with its log line, before the job's 5 minute task timeout kills it.
const raidPostsBudget = 4 * time.Minute

// newGuildStore is the guilds store serve and the job share, so their reads cannot drift.
func newGuildStore(cfg config.Config, pool *pgxpool.Pool, treeData *trees.Data) *guilds.Store {
	return &guilds.Store{Pool: pool, DataDir: cfg.TreeDataDir, Trees: treeData, APIBaseURL: cfg.APIBaseURL}
}

func runRaidPosts(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	svc, err := newRaidPostsServices(cfg, pool, log)
	if err != nil {
		return err
	}
	return raidPosts(ctx, log, svc, time.Now())
}

// raidScheduler sends the posts and reminders that are due (guilds.Service).
type raidScheduler interface {
	RunRaidSchedule(ctx context.Context, now time.Time) (guilds.ScheduleRun, error)
}

// raidDrainer runs one bot drain pass (discordbot.Drainer).
type raidDrainer interface {
	DrainOnce(ctx context.Context) error
}

// eventScheduler publishes the time-based player events that are due (guilds.Service).
type eventScheduler interface {
	RunEventSchedule(ctx context.Context, now time.Time) (guilds.EventScheduleRun, error)
}

// deliveryPruner forgets old deliveries (events.Store).
type deliveryPruner interface {
	PruneDeliveries(ctx context.Context, now time.Time, limit int) (int64, error)
}

// pruneBatch bounds the deliveries one run forgets.
const pruneBatch = 5000

// raidPostsServices is what the job runs. Drainer is nil when the bot is not configured, and
// Events and Prune are nil until the app notifications are on; tests build it with fakes.
type raidPostsServices struct {
	Guilds  raidScheduler
	Drainer raidDrainer
	Events  eventScheduler
	Prune   deliveryPruner
}

// newRaidPostsServices builds the guilds, notify and bot wiring serve uses for the raid router,
// but starts no drainer goroutine and registers no commands.
func newRaidPostsServices(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) (raidPostsServices, error) {
	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return raidPostsServices{}, fmt.Errorf("trees: %w", err)
	}
	entStore := &entitlements.Store{Pool: pool}
	guildService := &guilds.Service{
		Store:    newGuildStore(cfg, pool, treeData),
		Accounts: &auth.Store{Pool: pool, Log: log}, Supporters: entStore, Entitlements: entStore,
		PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
	deps := server.Deps{Guilds: guildService}
	if n := newNotifyService(cfg, pool, entStore, log); n != nil {
		deps.Notify = n
		wireGuildRaidChannel(n, guildService)
	}
	svc := raidPostsServices{Guilds: guildService}
	if router := newEventsRouter(context.Background(), cfg, pool, deps.Notify, log); cfg.Push.Enabled && router != nil {
		guildService.Events = router
		svc.Events, svc.Prune = guildService, &events.Store{Pool: pool}
	}
	if !botEnabled(cfg, deps, log) {
		return svc, nil
	}
	d := botDeps(cfg, pool, deps, log)
	bot := discordbot.New(d)
	guildService.Bot = bot
	guildService.RaidEvents = raidRouterFor(bot, deps)
	guildService.RaidPosts = d.Store
	svc.Drainer = bot.NewDrainer(guildService)
	return svc, nil
}

// raidPosts runs the schedule, then the drain passes. A schedule failure does not skip the
// drain (rows already queued still go out) but fails the job.
func raidPosts(ctx context.Context, log *slog.Logger, svc raidPostsServices, now time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, raidPostsBudget)
	defer cancel()
	var errs []error
	run, err := svc.Guilds.RunRaidSchedule(ctx, now)
	if err != nil {
		errs = append(errs, fmt.Errorf("schedule: %w", err))
	}
	if svc.Drainer != nil {
		for i := 0; i < raidPostsDrainPasses && ctx.Err() == nil; i++ {
			if err := svc.Drainer.DrainOnce(ctx); err != nil {
				errs = append(errs, fmt.Errorf("drain: %w", err))
				break
			}
		}
	}
	log.Info("raid-posts", "posted", run.Posted, "reminded", run.Reminded, "failed", run.Failed)
	if err := runEvents(ctx, log, svc, now); err != nil {
		errs = append(errs, err)
	}
	if err := errors.Join(errs...); err != nil {
		return fmt.Errorf("raid-posts: %w", err)
	}
	return nil
}

// runEvents publishes the due player events and prunes the delivery log; a no-op until the
// app notifications are on.
func runEvents(ctx context.Context, log *slog.Logger, svc raidPostsServices, now time.Time) error {
	if svc.Events == nil {
		return nil
	}
	var errs []error
	run, err := svc.Events.RunEventSchedule(ctx, now)
	if err != nil {
		errs = append(errs, fmt.Errorf("events: %w", err))
	}
	var pruned int64
	if svc.Prune != nil {
		if pruned, err = svc.Prune.PruneDeliveries(ctx, now, pruneBatch); err != nil {
			errs = append(errs, fmt.Errorf("prune: %w", err))
		}
	}
	log.Info("raid-posts", "events_posted", run.Posted, "nudged", run.Nudged, "short", run.Short,
		"starting", run.Starting, "pending", run.Pending, "pruned", pruned)
	return errors.Join(errs...)
}
