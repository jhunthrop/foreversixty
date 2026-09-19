package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/parse"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/site"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/subscribe"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/runner"
	"github.com/jhunthrop/foreversixty/sim/talents"
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
	// The image is the service and every job: Cloud Run runs it with
	// the job's name as its first container argument.
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case reports.ParseJobCommand:
			if err := runParse(context.Background(), log, os.Args[2:]); err != nil {
				log.Error(reports.ParseJobCommand, "err", err)
				os.Exit(1)
			}
			return
		case sims.SimRunJobCommand:
			if err := runSim(context.Background(), log, os.Args[2:]); err != nil {
				log.Error(sims.SimRunJobCommand, "err", err)
				os.Exit(1)
			}
			return
		}
	}
	if err := serve(log); err != nil {
		log.Error("startup", "err", err)
		os.Exit(1)
	}
}

// start loads the configuration, migrates, and connects. Both the
// service and the job need exactly this much.
func start(ctx context.Context) (config.Config, *pgxpool.Pool, error) {
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		return config.Config{}, nil, err
	}
	if err := db.Migrate(cfg.MigrateDatabaseURL); err != nil {
		return config.Config{}, nil, fmt.Errorf("migrate: %w", err)
	}
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return config.Config{}, nil, err
	}
	return cfg, pool, nil
}

// objects builds the R2 client, or nil when the deployment has no
// credentials for it: the service then serves everything that does not
// touch object storage.
func objects(cfg config.Config, log *slog.Logger) *r2.Client {
	if !cfg.R2Configured() {
		log.Warn("r2", "state", "not configured", "effect",
			"ingest, uploads, and report files are not served")
		return nil
	}
	client, err := r2.New(r2.Config{
		AccountID: cfg.R2AccountID, AccessKeyID: cfg.R2AccessKeyID,
		SecretAccessKey: cfg.R2SecretAccessKey, Bucket: cfg.R2Bucket, Endpoint: cfg.R2Endpoint,
	})
	if err != nil {
		log.Error("r2", "err", err)
		return nil
	}
	return client
}

// runParse is the Cloud Run job: parse one uploaded log into its
// report, then exit.
func runParse(ctx context.Context, log *slog.Logger, args []string) error {
	if len(args) != 1 || args[0] == "" {
		return fmt.Errorf("usage: api %s <report_id>", reports.ParseJobCommand)
	}
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("parse-report needs R2 credentials")
	}
	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return err
	}
	reportStore := &reports.Store{Pool: pool}
	log.Info("parse-report", "report", args[0])
	return parse.Report(ctx, parse.Deps{
		Reports: reportStore, Objects: client, Log: log,
		Rank: &rankings.Store{Pool: pool, Specs: inferrer(treeData)},
	}, args[0])
}

// runSim is the premium lane's Cloud Run job: run one sim natively
// and exit.
func runSim(ctx context.Context, log *slog.Logger, args []string) error {
	if len(args) != 1 || args[0] == "" {
		return fmt.Errorf("usage: api %s <sim_id>", sims.SimRunJobCommand)
	}
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("%s needs R2 credentials", sims.SimRunJobCommand)
	}
	log.Info(sims.SimRunJobCommand, "sim", args[0], "engine", enginever.Version)
	return sims.Run(ctx, sims.JobDeps{
		Store: &sims.Store{Pool: pool}, Put: client, Engine: simEngine(log), Log: log,
	}, args[0])
}

// simEngine is the runner every simulator job uses: the real binary
// when the image carries one, and the checked-in fixture when it does
// not, so a deployment without the artifact still answers instead of
// failing. Task 15 is what puts the binary there.
func simEngine(log *slog.Logger) runner.Runner {
	if _, err := os.Stat(runner.DefaultBinary); err == nil {
		return &runner.Native{}
	}
	log.Warn("sims", "state", "no engine binary at "+runner.DefaultBinary,
		"effect", "sims answer from the checked-in fixture result")
	return &runner.Fixture{}
}

// inferrer names specs from the newest client build's talent data.
func inferrer(data *trees.Data) *spec.Inferrer {
	b, ok := data.Latest()
	if !ok {
		return spec.New(nil)
	}
	return spec.New(b)
}

// talentLayouts loads the newest build's talent layout, which is what
// turns a fight's recorded talent ids into the engine's positional
// string. A build with no layout is not fatal: the scorer then refuses
// every character, which is what it already does for the missing race.
func talentLayouts(dir string, data *trees.Data, log *slog.Logger) *talents.Layouts {
	build, ok := data.Latest()
	if !ok {
		log.Warn("sims", "state", "no client build", "effect", "no execution scores")
		return nil
	}
	layouts, err := talents.Load(filepath.Join(dir, build.Version, "talents"))
	if err != nil {
		log.Warn("sims", "state", "no talent layout", "err", err, "effect", "no execution scores")
		return nil
	}
	return layouts
}

// scorerShim adapts the sims scorer to what the ingest asks for, so
// neither package has to import the other. The conversion compiles
// only while reports.ScoredFight and sims.FightAt have identical
// fields in identical order, which is the drift check.
type scorerShim struct{ s *sims.Scorer }

func (a scorerShim) Schedule(f reports.ScoredFight) {
	a.s.ScheduleFight(sims.FightAt(f))
}

func serve(log *slog.Logger) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()

	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return fmt.Errorf("trees: %s: %w", cfg.TreeDataDir, err)
	}
	for _, skipped := range treeData.Skipped() {
		log.Warn("trees", "skipped", skipped)
	}
	log.Info("trees", "dir", cfg.TreeDataDir, "versions", treeData.Versions())

	// Ranking rows are written into monthly partitions, which have to
	// exist before the first fight of the month closes.
	partitions := &db.PartitionJob{Pool: pool, Log: log}
	if err := partitions.Run(ctx); err != nil {
		return fmt.Errorf("partitions: %w", err)
	}

	buildStore := &builds.Store{Pool: pool, Log: log}
	views := builds.NewViews(buildStore, log)
	siteDeps := &site.Deps{
		Store: buildStore, Data: treeData, PublicBaseURL: cfg.PublicBaseURL, Views: views, Log: log,
	}
	buildsSvc := &builds.Service{
		Store: buildStore, Data: treeData, PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
	mailer := mail.NewResend(cfg.ResendAPIKey, cfg.MailFrom, nil)
	subscribeSvc := &subscribe.Service{
		Store: &subscribe.Store{Pool: pool}, Mailer: mailer,
		PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL, Logger: log,
		ResendCooldown: resendCooldown, MaxSendsPerHour: maxSendsPerHour,
	}

	authStore := &auth.Store{Pool: pool}
	authenticator := &auth.Authenticator{
		Store: authStore, CookieDomain: cfg.SessionCookieDomain,
		Secure: strings.HasPrefix(cfg.PublicBaseURL, "https://"), Log: log,
	}
	accounts := &auth.Service{
		Store: authStore, Auth: authenticator, Mailer: mailer,
		PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL, Log: log,
	}
	if cfg.BattleNetConfigured() {
		accounts.BNet = auth.NewBattleNet(cfg.BnetClientID, cfg.BnetClientSecret, cfg.BnetRedirectURL)
	} else {
		log.Warn("auth", "state", "battle.net is not configured", "effect", "email sign-in only")
	}

	reportStore := &reports.Store{Pool: pool}
	rankStore := &rankings.Store{Pool: pool, Specs: inferrer(treeData)}
	client := objects(cfg, log)

	deps := server.Deps{
		Version: version, Log: log, AllowedOrigin: cfg.PublicBaseURL,
		Subscribe: subscribeSvc, Builds: buildsSvc, Site: siteDeps,
		Auth: authenticator, Accounts: accounts,
		Reports: &reports.Service{
			Store: reportStore, Accounts: authStore, Rank: rankStore,
			PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL, Log: log,
		},
		Rankings: &rankings.Service{Store: rankStore, Log: log},
		Addon: &addon.Service{
			Store: &addon.Store{Pool: pool}, Builds: buildStore, Data: treeData, Log: log,
		},
		TrustedProxyHops: cfg.TrustedProxyHops,
	}

	simStore := &sims.Store{Pool: pool}
	deps.Sims = &sims.Service{
		Store: simStore, Accounts: authStore, EngineVersion: enginever.Version, Log: log,
	}

	var sampler *parse.Worker
	var scorer *sims.Scorer
	if client != nil {
		deps.Reports.Signer = client
		sampler = parse.NewWorker(parse.Deps{
			Reports: reportStore, Objects: client, Rank: rankStore, Log: log,
		})
		go sampler.Run(ctx)
		deps.Ingest = &reports.Ingest{
			Store: reportStore, Put: client, Rank: rankStore, Samp: sampler, Log: log,
		}
		// After deps.Ingest exists, never beside the sampler above it:
		// the next line needs the ingest to be there.
		scorer = sims.NewScorer(sims.ScoreDeps{
			Store: simStore, Scores: rankStore, Engine: simEngine(log),
			Build:         sims.CombatantBuilder{Talents: talentLayouts(cfg.TreeDataDir, treeData, log)},
			Summaries:     client,
			EngineVersion: enginever.Version, Log: log,
		})
		go scorer.Run(ctx)
		deps.Ingest.Score = scorerShim{scorer}
		deps.Ingest.Members = authStore
		if runner, err := jobs.NewCloudRun(ctx, cfg.ParseJobProject, cfg.ParseJobRegion, cfg.ParseJobName); err != nil {
			log.Warn("jobs", "state", "the parse job cannot be reached", "err", err,
				"effect", "whole-file uploads are not offered")
		} else {
			deps.Uploads = &reports.Uploads{
				Store: reportStore, R2: client, Jobs: runner, APIBaseURL: cfg.APIBaseURL, Log: log,
			}
		}
		// The bucket, for the buffs sim-input reads out of a stored
		// fight summary. Without it that read answers without buffs.
		deps.Sims.Summaries = client
		if runner, err := jobs.NewCloudRun(ctx, cfg.SimJobProject, cfg.SimJobRegion, cfg.SimJobName); err != nil {
			log.Warn("jobs", "state", "the sim job cannot be reached", "err", err,
				"effect", "running sims on our servers is not offered")
		} else {
			deps.Sims.Jobs = runner
		}
	}

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           server.NewRouter(deps),
		ReadHeaderTimeout: 5 * time.Second,
	}

	serveErr := make(chan error, 1)
	go func() {
		log.Info("listening", "port", cfg.Port, "version", version)
		serveErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-serveErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("serve: %w", err)
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
	subscribeSvc.Wait()
	views.Close()
	if sampler != nil {
		sampler.Close()
	}
	if scorer != nil {
		scorer.Close()
	}
	log.Info("stopped")
	return nil
}
