package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// runnerFactory binds a job runner to one Cloud Run job (a CloudRun runs one job, F23).
type runnerFactory func(ctx context.Context, job string) (jobs.Runner, error)

// coachWiring is what mountCoach takes from serve: the concrete stores serve already builds.
// newRunner is nil in production (Cloud Run in cfg's sim-job project and region).
type coachWiring struct {
	cfg       config.Config
	pool      *pgxpool.Pool
	deps      *server.Deps
	rating    *rating.Store
	reports   *reports.Store
	auth      *auth.Store
	entitle   *entitlements.Store
	sims      *sims.Store
	summaries coach.Getter
	log       *slog.Logger
	newRunner runnerFactory
}

// mountCoach builds the coach service (the model client only while the switch is on), sets
// the serving rating store's hook to the profile writer and the guilds' coach holders. It
// answers nil, with one warn line naming the effect, when the coach cannot run (no bucket,
// an unreadable COACH_DATA_DIR); the coach routes then stay unmounted. With the switch off
// the routes answer rule-only.
func mountCoach(ctx context.Context, w coachWiring) *coach.Service {
	if w.summaries == nil {
		w.log.Warn("coach", "state", "no bucket", "effect", "the coach routes are not served and no coach profiles are written")
		return nil
	}
	svc, err := newCoachService(ctx, w)
	if err != nil {
		w.log.Warn("coach", "state", "not built", "err", err, "effect", "the coach routes are not served and no coach profiles are written")
		return nil
	}
	w.rating.Hook = profileWriter(svc.Tables, w.log)
	if w.deps != nil && w.deps.Guilds != nil {
		w.deps.Guilds.CoachHolders = svc
	}
	return svc
}

// bucketOrNil keeps a nil *r2.Client from becoming a non-nil coach.Getter.
func bucketOrNil(c *r2.Client) coach.Getter {
	if c == nil {
		return nil
	}
	return c
}

// coachTables is the production TableSource: the image's build data and COACH_DATA_DIR's rule files.
func coachTables(cfg config.Config, simStore *sims.Store) *coach.FileTables {
	return &coach.FileTables{BuildsDir: cfg.TreeDataDir, CoachDir: cfg.Coach.DataDir, Specs: simStore}
}

// profileWriter is the rating hook that writes coach profiles: one construction for the
// service and the jobs.
func profileWriter(tables coach.TableSource, log *slog.Logger) *coach.ProfileWriter {
	return &coach.ProfileWriter{Tables: tables, Log: log}
}

// ratingStoreWithCoach is a rating store whose every rated fight writes coach profiles: the
// jobs that re-rate (backfill, rerate-flagged, rebuild-preparation) build theirs with it.
// They hook whether or not the service could mount the coach: profiles need only the build
// data and the rule files (a spec without them is skipped), never the bucket.
func ratingStoreWithCoach(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) *rating.Store {
	return &rating.Store{Pool: pool, Log: log, Hook: profileWriter(coachTables(cfg, &sims.Store{Pool: pool}), log)}
}

// newCoachService builds the service from the wiring. It fails on an unreadable
// COACH_DATA_DIR, a model client the config cannot make, or caps it cannot build.
func newCoachService(ctx context.Context, w coachWiring) (*coach.Service, error) {
	if err := readableDir(w.cfg.Coach.DataDir); err != nil {
		return nil, fmt.Errorf("COACH_DATA_DIR: %w", err)
	}
	caps, err := coach.NewCaps(w.cfg.Coach, w.pool, time.Now)
	if err != nil {
		return nil, err
	}
	model, err := coachModel(w.cfg.Coach)
	if err != nil {
		return nil, err
	}
	tables := coachTables(w.cfg, w.sims)
	store := &coach.Store{DB: w.pool}
	return &coach.Service{
		Store: store, Reports: w.reports, Accounts: w.auth, Owners: w.auth, Entitlements: w.entitle,
		Summaries: w.summaries, Tables: tables, Refs: coachRefs(ctx, w, tables),
		Personas: &coach.Personas{Dir: w.cfg.Coach.DataDir},
		Generator: &coach.Generator{LLM: model, Caps: caps, Store: store, Model: w.cfg.Coach.LLMModel,
			Temperature: w.cfg.Coach.Temperature, Log: w.log, Now: time.Now},
		Caps: caps, Jobs: w.runner(ctx, coach.ReviewJobCommand, "nothing is queued; reviews are rule-only"),
		Enabled: w.cfg.Coach.Enabled, PublicBaseURL: w.cfg.PublicBaseURL, Log: w.log, Now: time.Now,
	}, nil
}

func readableDir(dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", dir)
	}
	return nil
}

// coachRefs resolves the top logs and, for a validated DPS spec, the sim. Without a
// coach-simref runner no spec counts as validated, so nothing waits on a sim that cannot run.
func coachRefs(ctx context.Context, w coachWiring, tables *coach.FileTables) *coach.Resolver {
	refs := &coach.Resolver{Pool: &coach.Pool{DB: w.pool}, Sims: &coach.SimRefs{DB: w.pool, Log: w.log},
		Validated: w.sims, EngineVersion: enginever.Version, BuildAt: buildAt(tables, w.log)}
	if runner := w.runner(ctx, coach.SimRefJobCommand, "no sim references; the top logs are the only basis"); runner != nil {
		refs.Sims.Jobs = runner
	} else {
		refs.Validated = noSimValidated{}
	}
	return refs
}

// noSimValidated validates no spec: the coach then never asks for a sim reference.
type noSimValidated struct{}

func (noSimValidated) Validated(context.Context, string) (bool, error) { return false, nil }

// buildAt adapts FileTables.BuildAt to the Resolver's func: a fight before every known build
// is "" (no sim key matches it), logged.
func buildAt(tables *coach.FileTables, log *slog.Logger) func(time.Time) string {
	return func(t time.Time) string {
		b, err := tables.BuildAt(t)
		if err != nil {
			log.Warn("coach", "op", "build_at", "at", t, "err", err)
			return ""
		}
		return b
	}
}

// runner binds a runner to job, or answers nil after one warn line naming effect.
func (w coachWiring) runner(ctx context.Context, job, effect string) jobs.Runner {
	factory := w.newRunner
	if factory == nil {
		factory = func(ctx context.Context, job string) (jobs.Runner, error) {
			return jobs.NewCloudRun(ctx, w.cfg.SimJobProject, w.cfg.SimJobRegion, job)
		}
	}
	r, err := factory(ctx, job)
	if err == nil && r == nil {
		return nil // the factory binds nothing for this job, on purpose
	}
	if err != nil {
		w.log.Warn("jobs", "state", "the "+job+" job cannot be reached", "err", err, "effect", effect)
		return nil
	}
	return r
}

// coachModel is the gateway client while the switch is on, and a client that refuses every
// call while it is off (the caps refuse first, so it is never asked).
func coachModel(c config.CoachConfig) (llm.Client, error) {
	if !c.Enabled {
		return switchedOff{}, nil
	}
	return llm.New(llm.Config{
		BaseURL: c.LLMBaseURL, Model: c.LLMModel, APIKey: llm.NewSecret(c.LLMAPIKey), Providers: c.Providers(),
		PriceInPerMTok: c.LLMPriceInPerMTok, PriceOutPerMTok: c.LLMPriceOutPerMTok,
	}, nil)
}

// switchedOff is the model client of a coach whose switch is off.
type switchedOff struct{}

var errSwitchedOff = errors.New("coach: the model is switched off (COACH_LLM_ENABLED)")

func (switchedOff) Complete(context.Context, llm.Request) (llm.Response, error) {
	return llm.Response{}, errSwitchedOff
}
