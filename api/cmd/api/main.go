package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/addon"
	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/bnetbuild"
	"github.com/jhunthrop/foreversixty/api/internal/bnetimport"
	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/dataaddon"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
	"github.com/jhunthrop/foreversixty/api/internal/parse"
	"github.com/jhunthrop/foreversixty/api/internal/phase"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/site"
	"github.com/jhunthrop/foreversixty/api/internal/subscribe"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/runner"
)

var version = "dev" // set with -ldflags "-X main.version=<git sha>"

// resendCooldown and maxSendsPerHour guard against confirmation-resend
// abuse: an attacker (or a confused client) hammering POST /v1/subscribe for
// the same address can otherwise mail-bomb that address, and hammering
// distinct addresses can otherwise exhaust the mail provider's quota.
const (
	resendCooldown  = 15 * time.Minute
	maxSendsPerHour = 300

	// shutdownBudget is the one deadline for the whole stop: the server's shutdown and the drain
	// of detached report completion hooks share it, inside Cloud Run's 10 s grace period.
	shutdownBudget = 9 * time.Second
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// The image is the service and every job: Cloud Run runs it with
	// the job's name as its first container argument.
	if len(os.Args) > 1 {
		runJobTable(log, allCommands(), os.Args[1:], os.Exit)
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
		case sims.ValidateJobCommand:
			if err := runValidate(context.Background(), log); err != nil {
				log.Error(sims.ValidateJobCommand, "err", err)
				os.Exit(1)
			}
			return
		case rating.BackfillJobCommand:
			if err := runRatingBackfill(context.Background(), log); err != nil {
				log.Error(rating.BackfillJobCommand, "err", err)
				os.Exit(1)
			}
			return
		case rating.RerateFlaggedJobCommand:
			if err := runRerateFlagged(context.Background(), log); err != nil {
				log.Error(rating.RerateFlaggedJobCommand, "err", err)
				os.Exit(1)
			}
			return
		case dataaddon.JobCommand:
			if err := runDataAddon(context.Background(), log); err != nil {
				log.Error(dataaddon.JobCommand, "err", err)
				os.Exit(1)
			}
			return
		case bnetimport.RefreshJobCommand:
			if err := runBnetRefresh(context.Background(), log); err != nil {
				log.Error(bnetimport.RefreshJobCommand, "err", err)
				os.Exit(1)
			}
			return
		case "grant":
			if err := runGrant(context.Background(), log, os.Args[2:]); err != nil {
				log.Error("grant", "err", err)
				os.Exit(1)
			}
			return
		case "grandfather-officer-tools":
			if err := runGrandfatherOfficerTools(context.Background(), log, os.Args[2:]); err != nil {
				log.Error("grandfather-officer-tools", "err", err)
				os.Exit(1)
			}
			return
		case "revoke":
			if err := runRevoke(context.Background(), log, os.Args[2:]); err != nil {
				log.Error("revoke", "err", err)
				os.Exit(1)
			}
			return
		case "stripe-archive-prices":
			if err := runStripeArchivePrices(context.Background(), log, os.Args[2:]); err != nil {
				log.Error("stripe-archive-prices", "err", err)
				os.Exit(1)
			}
			return
		case "billing-daily":
			if err := runBillingDaily(context.Background(), log); err != nil {
				log.Error("billing-daily", "err", err)
				os.Exit(1)
			}
			return
		case "raid-posts":
			if err := runRaidPosts(context.Background(), log); err != nil {
				log.Error("raid-posts", "err", err)
				os.Exit(1)
			}
			return
		case "stripe-reconcile":
			if err := runStripeReconcile(context.Background(), log); err != nil {
				log.Error("stripe-reconcile", "err", err)
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
// service and the job need exactly this much. Each step logs how long it
// took: a cold start of the service was measured at 13 s end to end
// (2026-10-19) and the split between the database waking, the migration
// pass and the pool's first connection is what decides where to fix it.
func start(ctx context.Context) (config.Config, *pgxpool.Pool, error) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	began := time.Now()
	cfg, err := config.Load(os.Getenv)
	if err != nil {
		return config.Config{}, nil, err
	}
	if err := cfg.ValidateStripeKeyEnvironment(); err != nil {
		return config.Config{}, nil, err
	}
	log.Info("startup", "step", "config", "ms", time.Since(began).Milliseconds())

	began = time.Now()
	if err := db.Migrate(cfg.MigrateDatabaseURL); err != nil {
		return config.Config{}, nil, fmt.Errorf("migrate: %w", err)
	}
	log.Info("startup", "step", "migrate", "ms", time.Since(began).Milliseconds())

	began = time.Now()
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return config.Config{}, nil, err
	}
	log.Info("startup", "step", "connect", "ms", time.Since(began).Milliseconds())
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

// newReportsService builds the reports.Service exactly as serve wires
// it, pulled into its own function so main_test.go can assert the
// wiring itself: Guilds must always be populated, since a
// reports.Service with Guilds left nil silently disables the
// contested-and-frozen-claim report-edit freeze (D, third security
// review response) rather than failing loudly - reports.Service.Guilds
// being nil elsewhere (every test harness that does not care about
// claim disputes) is deliberately still supported, so the guard belongs
// here, at the one call site that matters for a real deployment, not as
// a universal nil-check inside reports.Service itself.
//
// hooks is the guilds service: told when a patch sets a report's guild
// (GuildMoved, releasing a match another guild's raid held) and when it
// sets the guild of a complete report (Completed).
func newReportsService(reportStore *reports.Store, authStore *auth.Store, guildStore *guilds.Store,
	rankStore *rankings.Store, hooks reports.GuildHooks, cfg config.Config, log *slog.Logger) *reports.Service {
	return &reports.Service{
		Store: reportStore, Accounts: authStore, Guilds: guildStore, Rank: rankStore, Completed: hooks, Moved: hooks,
		PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL, Log: log,
	}
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
	log.Info("parse-report", "report", args[0])
	return parse.Report(ctx, newParseDeps(cfg, pool, client, treeData, log), args[0])
}

// newParseDeps builds the parse-report job's dependencies on the job's
// own pool: the completion hook is a guilds service of its own, the
// way the job's Discord poster is a notify service of its own.
func newParseDeps(cfg config.Config, pool *pgxpool.Pool, client parse.Objects, treeData *trees.Data, log *slog.Logger) parse.Deps {
	deps := parse.Deps{
		Reports: &reports.Store{Pool: pool}, Objects: client, Log: log,
		Rank:      &rankings.Store{Pool: pool, Specs: inferrer(treeData)},
		Completed: newJobGuilds(cfg, pool, log),
	}
	if notifier := newJobNotifier(context.Background(), cfg, pool, log); notifier != nil {
		deps.Notify = notifier
	}
	return deps
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
	deps := sims.JobDeps{
		Store: &sims.Store{Pool: pool}, Put: client, Engine: simEngine(log), Log: log,
	}
	if notifier := newJobNotifier(ctx, cfg, pool, log); notifier != nil {
		deps.Notify = notifier
	}
	return sims.Run(ctx, deps, args[0])
}

// runValidate is the nightly Cloud Run job: measure every spec's
// fidelity against the top parses and publish the figures.
func runValidate(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("%s needs R2 credentials", sims.ValidateJobCommand)
	}
	at := phase.At(time.Now().UTC())
	specs := sims.DPSSpecs()
	log.Info(sims.ValidateJobCommand, "specs", len(specs), "phase", at,
		"engine", enginever.Version)
	return sims.Validate(ctx, sims.ValidateDeps{
		Store:  &sims.Store{Pool: pool},
		Top:    &sims.ParseReader{Pool: pool, Get: client, Log: log},
		Engine: simEngine(log), Build: sims.NoBuilder{},
		Scores: &rankings.Store{Pool: pool}, Log: log,
	}, specs, at, enginever.Version)
}

// runRatingBackfill is the nightly (and on-demand) Cloud Run job: recompute every
// rating_scores row not at logs/engine/rating's current model version.
func runRatingBackfill(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("%s needs R2 credentials", rating.BackfillJobCommand)
	}
	n, err := rating.Backfill(ctx, rating.BackfillDeps{
		Store: ratingStoreWithCoach(cfg, pool, log), Summaries: client, Log: log,
	}, rating.BackfillBatchSize)
	if err != nil {
		return err
	}
	log.Info(rating.BackfillJobCommand, "recomputed", n)
	return nil
}

// runRerateFlagged is the on-demand Cloud Run job that brings the stored Forever Score back
// into line once reports are flagged (rating.RerateFlagged): every character whose score a
// flagged report moved (in it, a guild-mate of its night, or an alt on the same account) has
// its newest clean fight re-rated and its score rolled up again. Idempotent, so it is safe to
// run again or on a schedule.
func runRerateFlagged(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	client := objects(cfg, log)
	if client == nil {
		return fmt.Errorf("%s needs R2 credentials", rating.RerateFlaggedJobCommand)
	}
	return rerateFlagged(ctx, rating.BackfillDeps{
		Store: ratingStoreWithCoach(cfg, pool, log), Summaries: client, Log: log,
	}, log)
}

// rerateFlagged is the job's run once its dependencies are wired: it logs the counts of each
// group, then fails when any fight could not be re-rated (every fight is still tried first),
// so the job exits non-zero and its failure alert fires.
func rerateFlagged(ctx context.Context, d rating.BackfillDeps, log *slog.Logger) error {
	c, err := rating.RerateFlagged(ctx, d)
	if err != nil {
		return err
	}
	for _, g := range []struct {
		name   string
		counts rating.RerateGroup
	}{{"flagged", c.Flagged}, {"guildmates", c.Guildmates}, {"linked", c.Linked}} {
		log.Info(rating.RerateFlaggedJobCommand, "group", g.name, "characters", g.counts.Characters,
			"fights", g.counts.Fights, "rerated", g.counts.Rerated, "failed", g.counts.Failed)
	}
	if failed := c.Failed(); failed > 0 {
		return fmt.Errorf("%d of %d fights could not be re-rated (each logged above)", failed, c.Fights())
	}
	return nil
}

// runDataAddon is the nightly Cloud Run job: aggregate every public rated
// character and guild into the Forever Sixty Data addon's Data.lua (or its
// region-split form) and publish it to DATA_ADDON_BUCKET for
// addon-data-release.yml to package. DATA_ADDON_BUCKET is read directly
// via os.Getenv, not added to config.Config: api/internal/config is
// outside this lane's file ownership (docs/superpowers/plans/
// 2026-09-21-data-addon.md), and this is the one job whose bucket address
// a future lane can fold into Config the normal way without this lane
// having touched that file first.
func runDataAddon(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()

	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return fmt.Errorf("trees: %s: %w", cfg.TreeDataDir, err)
	}
	build := "unknown"
	if active, ok := treeData.Active(); ok {
		build = active.Version
	} else {
		log.Warn(dataaddon.JobCommand, "err", "no client build available",
			"effect", "every gear component reads 0 and no no-logs record names a spec")
	}
	// The scores' gear component (rating.ScoreV2 through sims.GearScore) reads the item
	// tables and per-build weights, exactly as the API server wires them.
	sims.SetGearData(treeData, cfg.TreeDataDir)

	bucket := os.Getenv("DATA_ADDON_BUCKET")
	var uploader dataaddon.Uploader
	if bucket != "" {
		gcs, err := dataaddon.NewGCS(ctx)
		if err != nil {
			return fmt.Errorf("%s: %w", dataaddon.JobCommand, err)
		}
		uploader = gcs
	} else {
		log.Warn(dataaddon.JobCommand, "err", "DATA_ADDON_BUCKET is not set",
			"effect", "Data.lua is rendered but not published")
	}

	result, err := dataaddon.Run(ctx, dataaddon.Deps{
		Store: &dataaddon.Store{Pool: pool}, Trees: treeData, Upload: uploader, Bucket: bucket,
		Build: build, Now: time.Now().UTC(), Log: log,
	})
	if err != nil {
		return err
	}
	log.Info(dataaddon.JobCommand, "characters", result.Characters, "guilds", result.Guilds,
		"skipped", result.SkippedRows, "files", result.Files)
	return nil
}

// runBnetRefresh is the nightly Cloud Run job (spec §4.4): re-sync every
// stale 'bnet'-sourced character's guild membership from Blizzard's
// public data (no user token needed), and run the namespace probe
// (spec §5).
func runBnetRefresh(ctx context.Context, log *slog.Logger) error {
	cfg, pool, err := start(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()
	if !cfg.BattleNetConfigured() {
		return fmt.Errorf("%s needs Battle.net credentials", bnetimport.RefreshJobCommand)
	}
	treeData, err := trees.Load(cfg.TreeDataDir)
	if err != nil {
		return fmt.Errorf("trees: %s: %w", cfg.TreeDataDir, err)
	}
	tables, ok, err := bnetbuild.LoadTables(cfg.TreeDataDir, treeData)
	if err != nil {
		return fmt.Errorf("bnetbuild: load tables: %w", err)
	}
	if !ok {
		log.Warn(bnetimport.RefreshJobCommand, "err", "no client build available",
			"effect", "characters are re-synced but no Blizzard-sourced build is (re-)encoded")
	}
	svc := &bnetimport.Service{Pool: pool, Client: bnetClient(cfg, log), Regions: cfg.BnetRegions, Log: log, Tables: tables}
	if !cfg.BnetImportCharacters {
		// Probe only: the namespace_appeared warning is the launch-day signal, and no
		// character is re-synced from a namespace that is not Forever's (bnetimport.Stub).
		svc.Probe(ctx, cfg.BnetProbeGames)
		log.Info(bnetimport.RefreshJobCommand, "state", "character import is off, probed namespaces only")
		return nil
	}
	result, err := svc.RunRefresh(ctx, cfg.BnetProbeGames)
	if err != nil {
		return err
	}
	log.Info(bnetimport.RefreshJobCommand, "considered", result.Considered, "refreshed", result.Refreshed,
		"skipped", result.Skipped, "rate_limited", result.RateLimited)
	return nil
}

// bnetClient builds the Blizzard game-data/profile client every
// Battle.net-import call site (the login import, the nightly refresh)
// shares. It carries the service logger so the client's own WARN lines
// (a failed realm detail, a rate limit) are structured JSON like every
// other package's, rather than the unstructured text slog.Default()
// writes to stderr — invisible to a structured-log query, spec A1.
func bnetClient(cfg config.Config, log *slog.Logger) *bnetapi.Client {
	return bnetapi.New(bnetapi.Config{
		ClientID: cfg.BnetClientID, ClientSecret: cfg.BnetClientSecret,
		Game: cfg.BnetProfileGame, Regions: cfg.BnetRegions, Log: log,
	})
}

// simEngine is what every simulator job and the submit handler use: the
// real binary when the image carries one, and the checked-in fixture
// when it does not, so a deployment without the artifact still answers
// instead of failing. It runs sims and it counts bulk requests without
// running them (`forever-sim -plan`), which is why one value serves both
// the jobs and the service.
func simEngine(log *slog.Logger) sims.Engine {
	if _, err := os.Stat(runner.DefaultBinary); err == nil {
		return &runner.Native{}
	}
	log.Warn("sims", "state", "no engine binary at "+runner.DefaultBinary,
		"effect", "sims answer from the checked-in fixture result")
	return &runner.Fixture{}
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
	log.Info("trees", "dir", cfg.TreeDataDir, "versions", treeData.Versions(), "active", activeVersion(treeData), "active_from", treeData.ActiveFrom())

	// Ranking rows are written into monthly partitions, which have to
	// exist before the first fight of the month closes.
	partitions := &db.PartitionJob{Pool: pool, Log: log, Tables: []string{"fight_metrics", db.RatingsTable, db.CoachProfilesTable}}
	if err := partitions.Run(ctx); err != nil {
		return fmt.Errorf("partitions: %w", err)
	}
	// DataDir backs the guild control centre's own BiS band reads (api/internal/bis) -
	// the same TreeDataDir every other build-scoped read in this file uses. Trees backs
	// the loot tab's fallback candidate tier (class usability + item level from the
	// client's own per-class item tables, already loaded above).
	guildStore := newGuildStore(cfg, pool, treeData)
	membership := &guilds.MembershipJob{Store: guildStore, Log: log}
	if err := membership.Run(ctx); err != nil {
		return fmt.Errorf("guilds membership sweep: %w", err)
	}

	buildStore := &builds.Store{Pool: pool, Log: log}
	views := builds.NewViews(buildStore, log)
	// Built here, ahead of siteDeps, so the shared build page can find
	// a build's simmed DPS for its card description; simStore is
	// reused below for the sim service and scorer, which need the
	// same pool.
	simStore := &sims.Store{Pool: pool}
	siteDeps := &site.Deps{
		Store: buildStore, Data: treeData, PublicBaseURL: cfg.PublicBaseURL, Views: views,
		Sims: simStore, Log: log,
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

	authStore := &auth.Store{Pool: pool, Trees: treeData, Log: log}
	// The guild control centre's own claimed_by_name (home) and loot award
	// attribution (loot) resolve a user id to a battletag through this, the same
	// Accounts dependency guilds.Service already takes below.
	guildStore.Accounts = authStore
	entStore := &entitlements.Store{Pool: pool}
	authenticator := &auth.Authenticator{
		Store: authStore, CookieDomain: cfg.SessionCookieDomain,
		Secure: strings.HasPrefix(cfg.PublicBaseURL, "https://"), Log: log,
	}
	accounts := &auth.Service{
		Store: authStore, Auth: authenticator, Mailer: mailer, Entitlements: entStore,
		GuildCapabilities: guildStore, GuildAttendance: guildStore, Standing: guildStore, Sims: simsAccountView{store: simStore},
		PublicBaseURL: cfg.PublicBaseURL, APIBaseURL: cfg.APIBaseURL, Log: log,
	}
	if cfg.DevLogin {
		accounts.DevLogin = true
		log.Warn("auth", "state", "FS_DEV_LOGIN is on",
			"effect", "POST /v1/dev/login signs anyone in as a seeded account; loopback only, never for a deployment")
	}
	if cfg.BattleNetConfigured() {
		accounts.BNet = auth.NewBattleNet(cfg.BnetClientID, cfg.BnetClientSecret, cfg.BnetRedirectURL)
		bnetSvcClient := bnetClient(cfg, log)
		bnetTables, ok, err := bnetbuild.LoadTables(cfg.TreeDataDir, treeData)
		if err != nil {
			return fmt.Errorf("bnetbuild: load tables: %w", err)
		}
		if !ok {
			log.Warn("auth", "err", "no client build available",
				"effect", "Battle.net sign-in works but no character gets a Blizzard-sourced build")
		}
		if cfg.BnetImportCharacters {
			accounts.Importer = &bnetimport.Service{
				Pool: pool, Client: bnetSvcClient, Regions: cfg.BnetRegions, Log: log, Tables: bnetTables,
			}
		} else {
			// Blizzard serves no Forever namespace yet; see bnetimport.Stub.
			accounts.Importer = &bnetimport.Stub{Regions: cfg.BnetRegions}
			log.Warn("auth", "state", "battle.net character import is off",
				"effect", "sign-in works, no characters are imported; set BNET_IMPORT_CHARACTERS=true once Blizzard serves Forever")
		}
		// Warm the realm cache before traffic arrives (spec A1): a first
		// sign-in must never pay the cost of a cold per-region realm
		// fetch inside the import's 5s budget.
		go bnetSvcClient.WarmRealms(context.Background(), cfg.BnetRegions)
	} else {
		log.Warn("auth", "state", "battle.net is not configured", "effect", "email sign-in only")
	}

	reportStore := &reports.Store{Pool: pool}
	rankStore := &rankings.Store{Pool: pool, Specs: inferrer(treeData), APIBaseURL: cfg.APIBaseURL, Trees: treeData}
	ratingStore := &rating.Store{Pool: pool, Log: log}
	client := objects(cfg, log)

	guildService := &guilds.Service{
		Store: guildStore, Accounts: authStore, Supporters: entStore, Entitlements: entStore,
		PublicBaseURL: cfg.PublicBaseURL, Log: log,
	}
	deps := server.Deps{
		Version: version, Log: log, AllowedOrigin: cfg.PublicBaseURL,
		Subscribe: subscribeSvc, Builds: buildsSvc, Site: siteDeps,
		Auth: authenticator, Accounts: accounts,
		Reports:  newReportsService(reportStore, authStore, guildStore, rankStore, guildService, cfg, log),
		Rankings: &rankings.Service{Store: rankStore, Log: log, Supporters: entStore},
		Addon: &addon.Service{
			Store: &addon.Store{Pool: pool, Log: log}, Builds: buildStore, Data: treeData,
			Accounts: authStore, Guilds: guildStore, Sims: simStore, Log: log,
		},
		Guilds: guildService,
		Rating: &rating.Service{
			Store: ratingStore, Reports: reportStore, Accounts: authStore,
			Owners: authStore, Entitlements: entStore, Log: log,
		},
		TrustedProxyHops: cfg.TrustedProxyHops,
	}

	// The character page carries the Forever Score and the gear panel, which live in the rating
	// lane and share one read of the character.
	deps.Rankings.Scores = deps.Rating
	deps.Rankings.ScoreGear = deps.Rating
	deps.Coach = mountCoach(ctx, coachWiring{cfg: cfg, pool: pool, deps: &deps, rating: ratingStore, reports: reportStore, auth: authStore, entitle: entStore, sims: simStore, summaries: bucketOrNil(client), log: log})

	deps.Sims = &sims.Service{
		Store: simStore, Accounts: entStore, Planner: simEngine(log),
		EngineVersion: enginever.Version, Log: log,
	}
	// The addon outbox's sim requests start the same server runs; the
	// job runner is attached to deps.Sims below, when the job is reachable.
	deps.Addon.SimRequests = deps.Sims
	// sims.GearScore reads the item tables and the per-build weights.
	sims.SetGearData(treeData, cfg.TreeDataDir)
	wireNotify(cfg, pool, entStore, accounts, &deps, log)
	wireDiscord(ctx, cfg, pool, &deps, accounts, treeData, log)

	accounts.Seats = wireBilling(billingWiring{cfg: cfg, pool: pool, deps: &deps, entitle: entStore, authStore: authStore,
		guildStore: guildStore, mailer: mailer, log: log})
	// After every producer exists (guilds, addon, billing, seats): the events router reaches them all.
	wireEvents(ctx, cfg, pool, &deps, log)

	var sampler *parse.Worker
	var scorer *sims.Scorer
	var rater *rating.Rater
	if client != nil {
		deps.Reports.Signer = client
		deps.Guilds.R2 = client
		sampler = parse.NewWorker(parse.Deps{
			Reports: reportStore, Objects: client, Rank: rankStore, Log: log,
		})
		go sampler.Run(ctx)
		deps.Ingest = &reports.Ingest{
			Store: reportStore, Put: client, Rank: rankStore, Samp: sampler, Completed: deps.Guilds, Log: log,
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
		deps.Ingest.Levels = authStore
		rater = rating.NewRater(rating.RateDeps{Store: ratingStore, Log: log})
		go rater.Run(ctx)
		deps.Ingest.Rate = rater
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
		Addr:              net.JoinHostPort(cfg.ListenHost, cfg.Port),
		Handler:           server.NewRouter(deps),
		ReadHeaderTimeout: 5 * time.Second,
	}

	serveErr := make(chan error, 1)
	go func() {
		log.Info("listening", "port", cfg.Port, "version", version)
		serveErr <- srv.ListenAndServe()
	}()

	var stopCtx context.Context
	select {
	case err := <-serveErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("serve: %w", err)
		}
		var cancel context.CancelFunc
		stopCtx, cancel = context.WithTimeout(context.Background(), shutdownBudget)
		defer cancel()
	case <-ctx.Done():
		log.Info("shutting down")
		var cancel context.CancelFunc
		stopCtx, cancel = context.WithTimeout(context.Background(), shutdownBudget)
		defer cancel()
		if err := srv.Shutdown(stopCtx); err != nil {
			log.Error("shutdown", "err", err)
		}
		<-serveErr // wait for the listener goroutine to actually return
	}
	drainCompletions(stopCtx, deps, log)
	drainDiscordBot(stopCtx, deps, log)
	subscribeSvc.Wait()
	views.Close()
	if sampler != nil {
		sampler.Close()
	}
	if scorer != nil {
		scorer.Close()
	}
	if rater != nil {
		rater.Close()
	}
	log.Info("stopped")
	return nil
}
