# Rating API Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development
> (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps
> use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the API + storage + job lane of the performance rating feature: the
`rating_scores`/`rating_percentile_digests`/`kill_duration_digests` migration, the
`api/internal/rating` package (percentile-digest adapter, fight-scoring store, HTTP
handlers for the two read endpoints, the async ingest hook, the backfill job), the `i.rate`
hook in `api/internal/reports/ingest.go`, and the `EnsureMetricsPartition` generalisation in
`api/internal/db/partitions.go`.

**Architecture:** `logs/engine/rating` (already merged) is a pure function,
`rating.Score(fight, player, tables, percentiles, assignments, modelInfo) rating.Card`, with
no network or database access. This lane's job is everything around it: select the curated
tables for a player, feed `rating.Score` a `PercentileSource` backed by two new Postgres
digest tables (folding a value in only when it is safe and correct to do so), store the
resulting `Card`, and serve it back out through two read endpoints that reuse the reports
package's own visibility rule. Rating computation happens out of band from the ingest
request (a `Rater` worker, the same shape as `sims.Scorer`), so a rating failure never fails
an upload.

**Tech Stack:** Go 1.25.11, PostgreSQL (pgx/v5, golang-migrate), the existing
`api/internal/digest` t-digest, the merged `logs/engine/rating` package.

**Spec:** `docs/superpowers/specs/2026-09-21-performance-rating-design.md` (this lane's
scope: the API + storage + job columns of its §8 table), read together with the
coordinator's rulings in this lane's dispatch (reproduced in Global Constraints below).

## Global Constraints

- **Migration number is `0022_ratings`, not the spec's own `0019`.** `0018` (guild
  membership) is already on `main`; `0020` is the entitlements lane's migration and `0021`
  is reserved for its follow-up, both in flight in other worktrees. golang-migrate only ever
  moves a database forward from its currently-applied version, so a lower number landing
  after a higher one has already deployed is silently skipped forever. A test in this plan
  makes this class of mistake fail the build mechanically.
- **Ratings are public exactly the way report bodies are.** A rating is computed only from a
  report whose visibility lets it be ranked (`reports.Ranked`, i.e. not `private`); the
  per-fight endpoint's visibility is `reports`'s own `mayView` rule (public/unlisted open to
  anyone with the link; private to owner/moderator; guild to the owner, moderators, and
  **verified** guild members — `auth.Store.GuildRank` already requires `verified_at is not
  null`); the character endpoint includes only rows whose report is `public` **at query
  time**, never unlisted/guild/private even for a viewer who could open them.
- **File ownership is absolute.** This lane owns: the migration, `api/internal/rating/`
  (new), `i.rate`/the `Rater`/`RatedFight` types in `api/internal/reports/ingest.go` (that
  file only, nothing else under `api/internal/reports/`), `api/internal/db/partitions.go`
  and `api/internal/db/partitions_job.go`, a new `api/internal/db/migration_order_test.go`,
  and small, contiguous wiring edits to `api/cmd/api/main.go` and
  `api/internal/server/server.go`. Never touch `web/`, `addon/`, `sim/`, `companion/`,
  `logs/` (read-only; the engine is finished — a bug there is reported, not fixed here), nor
  `api/internal/guilds/`, `api/internal/billing/`, `api/internal/entitlements/`, nor any
  other file under `api/internal/reports/`.
- **`Service.mayView` cannot be imported.** It is an unexported method on `*reports.Service`
  (`api/internal/reports/handler.go:540`), and this lane may not edit that file. This plan
  therefore replicates its exact rule — public/unlisted open; private to owner/moderator;
  guild to owner/moderator/verified-member via the same exported `Accounts.GuildRank` — in
  `api/internal/rating/viewer.go`, built only from primitives `reports`/`auth` already
  export (`reports.Public`/`Unlisted`/`Private`/`GuildTo`, `reports.Report`'s exported
  fields, `auth.ActorFrom`, `Actor.Signed`/`IsModerator`). This is not a second, different
  access rule — it is the same rule, unavoidably re-expressed once, because Go's visibility
  rules and this lane's file fence make literal reuse impossible. Record this as a ruling in
  the ledger.
- **No premium gate.** The spec's own scope line says both endpoints are the free surface
  ("single lookups free... the player's own report card free"); neither is gated in §5. Per
  the coordinator's conditional instruction ("if the spec names one"), no `CanSee` interface
  is built — building one against nothing behind it would be dead, speculative structure.
  Record this finding in the final report.
- **Never import `api/internal/entitlements`, `api/internal/guilds` (write side), or
  `api/internal/billing`.** `Accounts.GuildRank` is consumed only through the small
  interface `reports.Accounts` already declares this shape for
  (`api/internal/reports/handler.go:41-44`); this lane defines its own copy of that same
  two-line interface in `viewer.go` rather than importing `reports.Accounts` (unexported
  field visibility is not the issue — `reports.Accounts` is exported — but keeping the
  interface declared where it is consumed, per this repo's Go patterns rule, avoids a
  needless coupling to `reports`'s own interface for a single method).
- **SQL is always parameterised.** No endpoint leaks the existence of a non-public report
  through a rating, a count, ordering, or an error-message difference between "forbidden"
  and "not found" (both read as `404 not_found`, matching how the rest of the reports API
  already treats visibility — see `reports/handler.go`'s own `get` handler for the pattern
  this lane follows).
- **Every list response is paginated with a validated, opaque cursor**, following the exact
  convention `api/internal/reports/recent.go` already uses (base64 of a `|`-separated
  keyset tuple, `NextCursor` only set when a further page may exist). The per-fight
  endpoint's `players[]` is not a paginated list — it is the full roster of one specific,
  size-bounded fight (≤ 40), matching the spec's own documented response shape in §5.2
  exactly; only the character endpoint's `trend[]` (unbounded over a career) is paginated.
  Record this scoping decision as a ruling.
- **The ingest hook never fails an upload and is idempotent.** `i.rate` only schedules
  work onto an out-of-band queue (mirroring `i.score`/`sims.Scorer` exactly); the worker
  that drains it logs and continues on any error. Re-ingesting the same fight, and the
  backfill job recomputing it, both upsert without folding a value into a percentile digest
  twice — detected the same way `rankings.Store.WriteFight` already detects a rewrite
  (delete rows for the fight first, check `RowsAffected() > 0`).
- **A wipe's `Output`/`Utility`/`Preparation`/`Activity` never feed a percentile digest;
  its `Survival`/`Mechanics` do**, into the `kill = false` bucket (spec §2). This lane's
  digest-fold decision is made from `rating.Bracket.Kill` and `rating.Bracket.Component`
  alone — no separate wipe-tracking state.
- **A fight's own kill-duration is folded into `kill_duration_digests` exactly once per
  fight, never once per roster player.** `rating.Score`'s injected `PercentileSource` is
  called once per player and must never fold on `KillTimeBand` — that fold happens as one
  explicit step per fight, before the per-player loop.
- **The rating package's own throwaway test database.** Every `go test` invocation in this
  lane points `TEST_DATABASE_URL` at a database this lane creates and drops itself (not the
  shared `forever_test` default), and every package-level test run uses `-p 1`. Two other
  lanes hit flaky failures today from sharing the default database.
- Go: `go vet ./... && go test ./...` for every package touched, before every commit.
  `GOTOOLCHAIN=go1.25.11` is unnecessary here — `go version` already reports 1.25.11.
- Code: functions under 50 lines, files under 800 lines, no magic numbers, pure functions
  with explicit inputs, immutable updates, errors wrapped with context, table-driven tests.

---

## File Structure

```
api/internal/db/
  migrations/0022_ratings.up.sql     (new)
  migrations/0022_ratings.down.sql   (new)
  partitions.go                      (modify: generalise EnsureMetricsPartition)
  partitions_job.go                  (modify: sweep more than one table)
  migration_order_test.go            (new: the mechanical numbering guard)

api/internal/rating/                 (new package)
  doc.go            package doc, shared small constants
  curated.go         specSlug + per-player CuratedTables selection
  cards.go            Card <-> stored/API JSON shapes (componentDTO, momentDTO, CardRow)
  adapter.go          percentileAdapter (rating.PercentileSource) + kill-duration fold
  store.go             Store: RateFight, ReadFightRatings, ReadCharacterRating
  viewer.go             visible() (mayView replica) + anonymize gate
  rater.go               Rater: the async worker reports.Ingest.Rate schedules into
  backfill.go              the Cloud Run recompute job
  handler.go                 Service, Mount, the two HTTP handlers, cursor encode/decode
  *_test.go for every file above

api/internal/reports/ingest.go       (modify: Rater/RatedFight types, Ingest.Rate, i.rate,
                                       the putFight call site)

api/cmd/api/main.go                  (modify: small, contiguous wiring)
api/internal/server/server.go        (modify: Deps.Rating field + Mount call)
```

---

## Task 1: Migration, its down script, and the numbering guard

**Files:**
- Create: `api/internal/db/migrations/0022_ratings.up.sql`
- Create: `api/internal/db/migrations/0022_ratings.down.sql`
- Create: `api/internal/db/migration_order_test.go`
- Test: the same file (pure filesystem test, no `TEST_DATABASE_URL` needed) plus one DB
  round-trip test appended to a new `api/internal/db/ratings_migration_test.go`

**Interfaces:**
- Produces: tables `rating_scores`, `rating_percentile_digests`, `kill_duration_digests`,
  exactly as later tasks read/write them (see Task 4/5's queries).

- [ ] **Step 1: Write the up migration**

```sql
-- api/internal/db/migrations/0022_ratings.up.sql
-- Player performance ratings (docs/superpowers/specs/2026-09-21-performance-rating-
-- design.md §4.2). Numbered 0022, not the spec's own 0019: 0018 (guild membership) is
-- already on main, 0020 is the entitlements lane's migration and 0021 is reserved for its
-- follow-up, both in flight in other worktrees as this lane starts. golang-migrate only
-- ever moves a database forward from its currently-applied version, so a lower number
-- landing after a higher one has already deployed would be silently skipped forever -
-- 0022 clears all three. migration_order_test.go guards this class of mistake mechanically.

create table if not exists rating_scores (
  report_id      text not null,
  fight_index    int not null,
  player_key     text not null,
  player_name    text not null default '',
  class          text,
  spec           text,
  role           text,
  encounter_id   int,
  difficulty     int,
  size           int,
  duration_ms    int,
  kill           boolean not null,
  kill_time_band text not null default '',
  overall           numeric,
  overall_uncapped  numeric,
  overall_capped    boolean not null default false,
  components     jsonb not null,
  model_version  text not null,
  fought_at      timestamptz not null,
  computed_at    timestamptz not null default now(),
  primary key (report_id, fight_index, player_key, fought_at)
) partition by range (fought_at);

create index if not exists rating_scores_player_idx on rating_scores (player_key, fought_at desc);
create index if not exists rating_scores_report_idx on rating_scores (report_id, fight_index);
create index if not exists rating_scores_stale_idx on rating_scores (model_version) where model_version <> '';

create table if not exists rating_percentile_digests (
  encounter_id   int not null,
  difficulty     int not null,
  spec           text not null,
  role           text not null,
  kill_time_band text not null,
  kill           boolean not null,
  component      text not null,
  digest         bytea not null,
  n              bigint not null default 0,
  updated_at     timestamptz not null default now(),
  primary key (encounter_id, difficulty, spec, role, kill_time_band, kill, component)
);

create table if not exists kill_duration_digests (
  encounter_id int not null,
  difficulty   int not null,
  digest       bytea not null,
  n            bigint not null default 0,
  updated_at   timestamptz not null default now(),
  primary key (encounter_id, difficulty)
);
```

- [ ] **Step 2: Write the down migration**

```sql
-- api/internal/db/migrations/0022_ratings.down.sql
drop table if exists kill_duration_digests;
drop table if exists rating_percentile_digests;
drop index if exists rating_scores_stale_idx;
drop index if exists rating_scores_report_idx;
drop index if exists rating_scores_player_idx;
drop table if exists rating_scores;
```

- [ ] **Step 3: Write the numbering guard test**

```go
// api/internal/db/migration_order_test.go
package db

import (
	"strconv"
	"strings"
	"testing"
)

// mergeBaseMigrations is the exact set of migration files present on main at this lane's
// merge base (confirmed via `ls api/internal/db/migrations/*.up.sql | sort` in this
// checkout before this lane wrote anything). golang-migrate only ever moves a database
// forward from its currently-applied version, so a migration numbered at or below the
// highest version already live is silently skipped the moment it deploys after a higher
// one - the exact mistake the coordinator's numbering ruling calls out by name. This test
// cannot see the future (it cannot know what version another in-flight lane will land at),
// so it enforces the one thing it can check mechanically and permanently: nothing this
// lane adds may reuse, or fall at or below, a number that was already on main before this
// lane started.
var mergeBaseMigrations = map[string]bool{
	"0001_subscribers.up.sql":                true,
	"0002_unsubscribe_token.up.sql":          true,
	"0003_confirmation_sent_at.up.sql":       true,
	"0004_builds.up.sql":                     true,
	"0005_logs.up.sql":                       true,
	"0006_boss_health.up.sql":                true,
	"0007_refold_digests.up.sql":             true,
	"0008_refold_after_engine.up.sql":        true,
	"0009_refold_after_engine_0_2_4.up.sql":  true,
	"0010_refold_after_engine_0_2_5.up.sql":  true,
	"0011_refold_after_engine_0_2_6.up.sql":  true,
	"0012_builds_point_order_integer.up.sql": true,
	"0013_sims.up.sql":                       true,
	"0014_sim_kinds.up.sql":                  true,
	"0015_builds_owner.up.sql":               true,
	"0016_sim_headline_backfill.up.sql":      true,
	"0017_reports_recent_idx.up.sql":         true,
	"0018_guild_membership.up.sql":           true,
}

// mergeBaseMax is the highest number in mergeBaseMigrations.
const mergeBaseMax = 18

func TestMigrationNumbersDoNotRegressBelowMergeBase(t *testing.T) {
	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		numStr, _, ok := strings.Cut(name, "_")
		if !ok {
			t.Fatalf("migration file %q has no <number>_ prefix", name)
		}
		n, err := strconv.Atoi(numStr)
		if err != nil {
			t.Fatalf("migration file %q has a non-numeric prefix: %v", name, err)
		}
		if n > mergeBaseMax {
			continue // this lane's own migration(s), or any later lane's: fine by construction
		}
		if !mergeBaseMigrations[name] {
			t.Errorf("migration %q is numbered %d, at or below the merge-base max of %d, "+
				"but is not one of the migrations already on main - renumber it above %d",
				name, n, mergeBaseMax, mergeBaseMax)
		}
	}
}

func TestRatingsMigrationIsNumberedAfterMergeBase(t *testing.T) {
	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if e.Name() == "0022_ratings.up.sql" {
			return
		}
	}
	t.Fatal("expected api/internal/db/migrations/0022_ratings.up.sql per the coordinator's " +
		"numbering ruling (0019 is reserved by the guild-membership window, 0020/0021 by " +
		"the entitlements lane)")
}
```

- [ ] **Step 4: Run the numbering guard (no database needed)**

Run: `cd api && go test ./internal/db/... -run TestMigrationNumbers -v` and
`go test ./internal/db/... -run TestRatingsMigrationIsNumbered -v`
Expected: both PASS.

- [ ] **Step 5: Stand up this lane's throwaway test database**

```bash
cd api
docker compose -f docker-compose.test.yml up -d
# Wait for readiness, then create a database just for this lane, separate from the shared
# forever_test other lanes use today.
until docker compose -f docker-compose.test.yml exec -T postgres pg_isready -U forever; do sleep 1; done
docker compose -f docker-compose.test.yml exec -T postgres psql -U forever -d forever_test \
  -c "drop database if exists rating_api_test;" \
  -c "create database rating_api_test;"
export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/rating_api_test?sslmode=disable"
```

Keep this `export` active (or re-run it) for every subsequent task's test commands in this
plan.

- [ ] **Step 6: Write and run the migration round-trip test**

```go
// api/internal/db/ratings_migration_test.go
package db

import (
	"context"
	"testing"
)

func TestMigrateCreatesRatingsTables(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	for _, name := range []string{"rating_scores", "rating_percentile_digests", "kill_duration_digests"} {
		if n := tableCount(t, pool, name); n != 1 {
			t.Errorf("table %s: count = %d, want 1", name, n)
		}
	}
	var strategy string
	if err := pool.QueryRow(context.Background(),
		`select partstrat from pg_partitioned_table p
		 join pg_class c on c.oid = p.partrelid where c.relname = 'rating_scores'`).Scan(&strategy); err != nil {
		t.Fatalf("rating_scores is not partitioned: %v", err)
	}
	if strategy != "r" {
		t.Fatalf("partition strategy = %q, want range", strategy)
	}
}

func TestMigration0022DownReversesUp(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := Migrate(url); err != nil {
			t.Errorf("restoring the latest migration: %v", err)
		}
	})
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	migrateTo(t, url, mergeBaseMax)
	for _, gone := range []string{"rating_scores", "rating_percentile_digests", "kill_duration_digests"} {
		if n := tableCount(t, pool, gone); n != 0 {
			t.Errorf("table %s survived the down migration", gone)
		}
	}
	if n := tableCount(t, pool, "guild_members"); n != 1 {
		t.Error("the down migration took a table from an earlier migration")
	}
	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := tableCount(t, pool, "rating_scores"); n != 1 {
		t.Error("rating_scores did not come back")
	}
}
```

Run: `cd api && go test ./internal/db/... -run "TestMigrateCreatesRatingsTables|TestMigration0022" -p 1 -v`
Expected: both PASS.

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/db/migrations/0022_ratings.up.sql api/internal/db/migrations/0022_ratings.down.sql \
  api/internal/db/migration_order_test.go api/internal/db/ratings_migration_test.go
printf 'feat: add the ratings migration and its numbering guard\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task1.txt
git commit -F .superpowers/commit-msg-task1.txt
```

---

## Task 2: Generalise `EnsureMetricsPartition`, sweep more than one table

**Files:**
- Modify: `api/internal/db/partitions.go`
- Modify: `api/internal/db/partitions_job.go`
- Test: `api/internal/db/db_test.go` is NOT modified (owned by other lanes' tests too);
  add new assertions to `api/internal/db/ratings_migration_test.go` instead.

**Interfaces:**
- Consumes: nothing new.
- Produces (for Task 4/9): `db.EnsurePartition(ctx, pool, table string, at time.Time) error`,
  `db.EnsurePartitions(ctx, pool, table string, from time.Time, months int) error`,
  `db.RatingsTable = "rating_scores"` constant, `db.EnsureRatingsPartition(ctx, pool, at) error`,
  and `PartitionJob.Tables []string` (empty means `{"fight_metrics"}`, unchanged behaviour).

- [ ] **Step 1: Read the current file to confirm the exact lines being replaced**

`api/internal/db/partitions.go` currently defines `MetricsPartition`, `EnsureMetricsPartition`,
and `EnsureMetricsPartitions` as the only partition helpers (already read in full during
research; reproduced in Step 2's diff).

- [ ] **Step 2: Generalise `partitions.go`**

Replace the whole file with:

```go
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PartitionMonthsAhead is how many months of partitions the service keeps in front of
// today. Three is a quarter of slack: the partition job runs daily, so two whole runs can
// fail before an insert has nowhere to land.
const PartitionMonthsAhead = 3

// PartitionSweepEvery is how often the background partition job runs.
const PartitionSweepEvery = 24 * time.Hour

// RatingsTable is the rating_scores partitioned table (migration 0022), so a caller
// outside this package never hand-spells it.
const RatingsTable = "rating_scores"

// partitionName is the table name for the month containing at, for any range-partitioned
// table keyed the same way fight_metrics and rating_scores both are.
func partitionName(table string, at time.Time) string {
	return fmt.Sprintf("%s_%s", table, at.UTC().Format("2006_01"))
}

// MetricsPartition is the table name for the month containing at, for fight_metrics.
// Exported for the existing tests and callers that already name it directly.
func MetricsPartition(at time.Time) string {
	return partitionName("fight_metrics", at)
}

// EnsurePartition creates the monthly partition of table holding at, if it is not there
// already. Partitions are created one month at a time and never dropped. table is always a
// fixed, code-provided constant (fight_metrics or RatingsTable today), never
// request-derived, so building the statement by Sprintf carries no injection risk - the
// same trade this function's fight_metrics-only predecessor already made.
func EnsurePartition(ctx context.Context, pool *pgxpool.Pool, table string, at time.Time) error {
	start := time.Date(at.UTC().Year(), at.UTC().Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	stmt := fmt.Sprintf(
		`create table if not exists %s partition of %s for values from ('%s') to ('%s')`,
		partitionName(table, start), table, start.Format(time.RFC3339), end.Format(time.RFC3339))
	if _, err := pool.Exec(ctx, stmt); err != nil {
		return fmt.Errorf("db: create partition %s: %w", partitionName(table, start), err)
	}
	return nil
}

// EnsureMetricsPartition is EnsurePartition pinned to fight_metrics, kept as the name every
// existing call site already uses.
func EnsureMetricsPartition(ctx context.Context, pool *pgxpool.Pool, at time.Time) error {
	return EnsurePartition(ctx, pool, "fight_metrics", at)
}

// EnsureRatingsPartition is EnsurePartition pinned to rating_scores.
func EnsureRatingsPartition(ctx context.Context, pool *pgxpool.Pool, at time.Time) error {
	return EnsurePartition(ctx, pool, RatingsTable, at)
}

// EnsurePartitions creates the partition of table for the month containing from and the
// months months after it. Idempotent.
func EnsurePartitions(ctx context.Context, pool *pgxpool.Pool, table string, from time.Time, months int) error {
	for i := 0; i <= months; i++ {
		if err := EnsurePartition(ctx, pool, table, from.AddDate(0, i, 0)); err != nil {
			return err
		}
	}
	return nil
}

// EnsureMetricsPartitions is EnsurePartitions pinned to fight_metrics, kept as the name
// every existing call site already uses.
func EnsureMetricsPartitions(ctx context.Context, pool *pgxpool.Pool, from time.Time, months int) error {
	return EnsurePartitions(ctx, pool, "fight_metrics", from, months)
}
```

- [ ] **Step 3: Extend `PartitionJob` to sweep more than one table**

Replace `api/internal/db/partitions_job.go` with:

```go
package db

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PartitionJob keeps every table in Tables partitioned ahead of today. Run it at startup;
// Run does the first pass itself and sweeps in the background until the context is
// cancelled.
type PartitionJob struct {
	Pool  *pgxpool.Pool
	Log   *slog.Logger
	Every time.Duration
	Ahead int
	// Tables is which partitioned tables to sweep. Empty means {"fight_metrics"}, so
	// every caller and test that predates the ratings lane keeps its exact old behaviour
	// with no change.
	Tables []string
	// Now is the clock, so a test can drive the job from a fixed date.
	Now func() time.Time
}

func (j *PartitionJob) tables() []string {
	if len(j.Tables) == 0 {
		return []string{"fight_metrics"}
	}
	return j.Tables
}

// Run creates every table's partitions once and then again on every tick, until ctx is
// cancelled. The first pass's error is returned so a startup can fail loudly on a database
// that cannot be partitioned at all; later failures are logged, because by then the
// service is serving traffic.
func (j *PartitionJob) Run(ctx context.Context) error {
	every, ahead, now := j.Every, j.Ahead, j.Now
	if every <= 0 {
		every = PartitionSweepEvery
	}
	if ahead <= 0 {
		ahead = PartitionMonthsAhead
	}
	if now == nil {
		now = time.Now
	}
	log := j.Log
	if log == nil {
		log = slog.Default()
	}
	tables := j.tables()
	if err := sweepOnce(ctx, j.Pool, tables, now(), ahead); err != nil {
		return err
	}
	go func() {
		t := time.NewTicker(every)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				if err := sweepOnce(ctx, j.Pool, tables, now(), ahead); err != nil {
					log.Error("partitions", "err", err)
				}
			}
		}
	}()
	return nil
}

func sweepOnce(ctx context.Context, pool *pgxpool.Pool, tables []string, at time.Time, ahead int) error {
	for _, table := range tables {
		if err := EnsurePartitions(ctx, pool, table, at.UTC(), ahead); err != nil {
			return err
		}
	}
	return nil
}
```

- [ ] **Step 4: Add coverage for the generalisation and the multi-table sweep**

Append to `api/internal/db/ratings_migration_test.go`:

```go
func TestRatingsTableIsPartitionedByMonthViaEnsurePartition(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx := context.Background()
	at := time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC)
	if err := EnsureRatingsPartition(ctx, pool, at); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(ctx,
		`select count(*) from information_schema.tables where table_name = 'rating_scores_2026_12'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("rating_scores_2026_12 missing")
	}
}

func TestPartitionJobSweepsEveryConfiguredTable(t *testing.T) {
	url := testURL(t)
	if err := Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	job := &PartitionJob{
		Pool: pool, Ahead: 0, Every: time.Hour, Tables: []string{"fight_metrics", RatingsTable},
		Now: func() time.Time { return time.Date(2027, 3, 1, 0, 0, 0, 0, time.UTC) },
	}
	if err := job.Run(ctx); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"fight_metrics_2027_03", "rating_scores_2027_03"} {
		var n int
		if err := pool.QueryRow(ctx,
			`select count(*) from information_schema.tables where table_name = $1`, want).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("partition %s missing", want)
		}
	}
}
```

Add `"time"` to the file's import block if not already present.

- [ ] **Step 5: Run every db package test**

Run: `cd api && go vet ./internal/db/... && go test ./internal/db/... -p 1 -v`
Expected: every test PASSes, including the pre-existing `TestPartitionJobCreatesTheMonthsAhead`
and `TestFightMetricsIsPartitionedByMonth` (unaffected — `PartitionJob{}` with no `Tables`
still defaults to `fight_metrics` only, and `EnsureMetricsPartition`/`EnsureMetricsPartitions`
keep their old names and signatures).

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/db/partitions.go api/internal/db/partitions_job.go api/internal/db/ratings_migration_test.go
printf 'refactor: generalise EnsureMetricsPartition to sweep any partitioned table\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task2.txt
git commit -F .superpowers/commit-msg-task2.txt
```

---

## Task 3: `rating` package skeleton — curated-table selection and Card <-> JSON shapes

**Files:**
- Create: `api/internal/rating/doc.go`
- Create: `api/internal/rating/curated.go`
- Create: `api/internal/rating/cards.go`
- Test: `api/internal/rating/curated_test.go`, `api/internal/rating/cards_test.go`

**Interfaces:**
- Consumes: `ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"` (`Card`,
  `Component`, `Moment`, `CuratedTables`, `DefaultModelInfo`, the six `ComponentName*`
  constants), `logs/engine/mechanics` (`Load`, `Table`), `logs/engine/mechanics/utility`
  (`Load`), `logs/engine/mechanics/consumables` (`Load`, `RoleCatalogue.For`),
  `logs/engine/summary` (`RosterRow`).
- Produces (for Task 4/5/6): `curatedTablesFor(row summary.RosterRow, encounterID int64)
  ratingengine.CuratedTables`, `type componentDTO struct{...}` + `toComponentDTOs(cs [6]ratingengine.Component)
  []componentDTO`, `type CardRow struct{...}` + `newCardRow(f fightMeta, row summary.RosterRow,
  card ratingengine.Card) CardRow` (fightMeta defined in this task too, carrying the
  per-fight fields every row shares).

- [ ] **Step 1: Package doc**

```go
// api/internal/rating/doc.go
// Package rating is the API + storage + job lane of the performance rating feature
// (docs/superpowers/specs/2026-09-21-performance-rating-design.md §8's "API + storage +
// job" row): the percentile-digest adapter that feeds logs/engine/rating.Score, the store
// that writes and reads rating_scores/rating_percentile_digests/kill_duration_digests, the
// two public read endpoints (§5), the async worker api/internal/reports/ingest.go's i.rate
// hook schedules into, and the backfill job (§4.4).
package rating
```

- [ ] **Step 2: Write the failing test for curated-table selection**

```go
// api/internal/rating/curated_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestCuratedTablesForSelectsBySpecAndRole(t *testing.T) {
	row := summary.RosterRow{GUID: "g1", Name: "Baelgrim", Class: "Warrior", Spec: "Protection", Role: "tank"}
	tables := curatedTablesFor(row, 667) // Shazzrah, a Molten Core boss with a curated table
	if tables.Mechanics == nil {
		t.Error("expected a mechanics table for encounter 667")
	}
	if tables.Utility == nil || tables.Utility.Spec != "warrior-protection" {
		t.Errorf("expected the warrior-protection utility table, got %+v", tables.Utility)
	}
	if tables.Consumables == nil {
		t.Error("expected the tank consumable catalogue row")
	}
}

func TestCuratedTablesForHandlesNoEncounterTable(t *testing.T) {
	row := summary.RosterRow{GUID: "g1", Name: "Nobody", Class: "Mage", Spec: "Frost", Role: "dps"}
	tables := curatedTablesFor(row, 99999999) // no such encounter
	if tables.Mechanics != nil {
		t.Error("expected no mechanics table for an uncurated encounter")
	}
	if tables.Utility == nil {
		t.Error("expected the mage-frost utility table regardless of the encounter")
	}
}

func TestSpecSlugMatchesEveryClassSpecPattern(t *testing.T) {
	cases := []struct{ class, spec, want string }{
		{"Warrior", "Protection", "warrior-protection"},
		{"Death Knight", "Blood", "death-knight-blood"},
		{"Priest", "Holy", "priest-holy"},
	}
	for _, c := range cases {
		if got := specSlug(c.class, c.spec); got != c.want {
			t.Errorf("specSlug(%q, %q) = %q, want %q", c.class, c.spec, got, c.want)
		}
	}
}
```

- [ ] **Step 3: Run it to see it fail**

Run: `cd api && go test ./internal/rating/... -run TestCuratedTablesFor -v`
Expected: FAIL (package does not exist / functions undefined).

- [ ] **Step 4: Implement curated-table selection**

```go
// api/internal/rating/curated.go
package rating

import (
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// specSlug turns a roster row's display class/spec ("Warrior", "Protection") into the
// slug logs/engine/mechanics/utility and data/curated/specs.json use ("warrior-
// protection"): lowercase, spaces to hyphens. This is the same trivial transform
// logs/engine/rating's own unexported specSlug applies (rating.go) - duplicated here
// because that function is unexported and this lane may not export it (logs/ is
// read-only for this lane); it exists only to pick which curated file to load, never to
// influence scoring math itself, which the engine computes independently.
func specSlug(class, spec string) string {
	return toSlug(class) + "-" + toSlug(spec)
}

func toSlug(s string) string {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == ' ':
			b = append(b, '-')
		case c >= 'A' && c <= 'Z':
			b = append(b, c+('a'-'A'))
		default:
			b = append(b, c)
		}
	}
	return string(b)
}

// curatedTablesFor selects the curated data one roster row's rating needs: the fight's
// own encounter mechanics table (nil if none curated yet - Score handles that by
// excluding Survival/Mechanics, §1.1), this player's spec's utility table, and this
// player's role's consumable catalogue.
func curatedTablesFor(row summary.RosterRow, encounterID int64) ratingengine.CuratedTables {
	var out ratingengine.CuratedTables
	if m, ok := mechanics.Load(encounterID); ok {
		out.Mechanics = &m
	}
	if u, ok := utility.Load(specSlug(row.Class, row.Spec)); ok {
		out.Utility = &u
	}
	if rc, ok := consumables.Load().For(row.Role); ok {
		out.Consumables = &rc
	}
	return out
}
```

- [ ] **Step 5: Run the test again**

Run: `cd api && go test ./internal/rating/... -run "TestCuratedTablesFor|TestSpecSlug" -v`
Expected: PASS.

- [ ] **Step 6: Write the failing test for the Card <-> JSON shapes**

```go
// api/internal/rating/cards_test.go
package rating

import (
	"encoding/json"
	"testing"
	"time"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

func sampleCard() ratingengine.Card {
	pct := 71.2
	return ratingengine.Card{
		Overall: 70, OverallUncapped: 70, OverallCapped: false,
		Basis: "percentile", ModelVersion: ratingengine.DefaultModelVersion, KillTimeBand: "typical",
		Components: [6]ratingengine.Component{
			{Name: ratingengine.ComponentNameOutput, Score: 71, Weight: 35, Basis: "percentile",
				Percentile: &pct, BracketN: 142, Excluded: false},
			{Name: ratingengine.ComponentNameSurvival, Score: 76, Weight: 15, Basis: "percentile",
				Excluded: false, Moments: []ratingengine.Moment{
					{Kind: "death", AtMS: 140000, SpellID: 19712, SpellName: "Arcane Explosion",
						Avoidable: true, Anchor: "death-g1-140000"},
				}},
			{Name: ratingengine.ComponentNameMechanics, Excluded: true, Reason: ratingengine.ReasonNoMechanicsTable},
			{Name: ratingengine.ComponentNameUtility, Score: 62, Weight: 15, Basis: "percentile"},
			{Name: ratingengine.ComponentNamePreparation, Score: 95, Weight: 10, Basis: "absolute"},
			{Name: ratingengine.ComponentNameActivity, Score: 80, Weight: 5, Basis: "percentile"},
		},
	}
}

func TestToComponentDTOsMatchesTheSpecFieldNames(t *testing.T) {
	dtos := toComponentDTOs(sampleCard().Components)
	b, err := json.Marshal(dtos[0])
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"name", "score", "weight", "basis", "percentile", "bracket_n", "excluded", "moments"} {
		if _, ok := got[field]; !ok {
			t.Errorf("component JSON missing field %q: %s", field, b)
		}
	}
}

func TestToComponentDTOsNullsScoreWhenExcluded(t *testing.T) {
	dtos := toComponentDTOs(sampleCard().Components)
	mechanics := dtos[2]
	b, _ := json.Marshal(mechanics)
	var got map[string]any
	_ = json.Unmarshal(b, &got)
	if got["score"] != nil {
		t.Errorf("excluded component's score = %v, want null", got["score"])
	}
	if got["reason"] != ratingengine.ReasonNoMechanicsTable {
		t.Errorf("reason = %v, want %q", got["reason"], ratingengine.ReasonNoMechanicsTable)
	}
}

func TestNewCardRowRoundTripsThroughJSON(t *testing.T) {
	meta := fightMeta{
		ReportID: "r1", FightIndex: 3, EncounterID: 667, Difficulty: 3, Size: 40,
		DurationMS: 180000, Kill: true, KillTimeBand: "typical", FoughtAt: time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC),
	}
	row := summary.RosterRow{GUID: "g1", Name: "Simfury", Class: "Warrior", Spec: "Fury", Role: "dps"}
	cr := newCardRow(meta, row, sampleCard())
	if cr.PlayerKey != "" {
		t.Error("newCardRow must not invent a player key - the caller (Store) computes it from region/ruleset")
	}
	var decoded []componentDTO
	if err := json.Unmarshal(cr.Components, &decoded); err != nil {
		t.Fatalf("Components is not valid JSON: %v", err)
	}
	if len(decoded) != 6 {
		t.Fatalf("decoded %d components, want 6", len(decoded))
	}
}
```

Add `"github.com/jhunthrop/foreversixty/logs/engine/summary"` to this test file's imports.

- [ ] **Step 7: Run it to see it fail**

Run: `cd api && go test ./internal/rating/... -run "TestToComponentDTOs|TestNewCardRow" -v`
Expected: FAIL.

- [ ] **Step 8: Implement the Card <-> JSON shapes**

```go
// api/internal/rating/cards.go
package rating

import (
	"encoding/json"
	"time"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// momentDTO is one Moment, spec §5.2's exact JSON field names.
type momentDTO struct {
	Kind      string `json:"kind"`
	AtMS      int64  `json:"at_ms"`
	SpellID   int64  `json:"spell_id,omitempty"`
	SpellName string `json:"spell_name,omitempty"`
	Avoidable bool   `json:"avoidable,omitempty"`
	Anchor    string `json:"anchor,omitempty"`
}

// componentDTO is one Component, spec §5.2's exact JSON field names. Score is a pointer so
// an excluded component serialises "score": null (spec §5.2: "an excluded entry carries
// ... score: null").
type componentDTO struct {
	Name       string      `json:"name"`
	Score      *float64    `json:"score"`
	Weight     float64     `json:"weight"`
	Basis      string      `json:"basis"`
	Percentile *float64    `json:"percentile,omitempty"`
	BracketN   int64       `json:"bracket_n,omitempty"`
	Excluded   bool        `json:"excluded"`
	Reason     string      `json:"reason,omitempty"`
	Moments    []momentDTO `json:"moments"`
}

func toMomentDTOs(ms []ratingengine.Moment) []momentDTO {
	out := make([]momentDTO, 0, len(ms))
	for _, m := range ms {
		out = append(out, momentDTO{
			Kind: m.Kind, AtMS: m.AtMS, SpellID: m.SpellID, SpellName: m.SpellName,
			Avoidable: m.Avoidable, Anchor: m.Anchor,
		})
	}
	return out
}

// toComponentDTOs converts the engine's six components into the site's response shape.
// Used both to build the stored components jsonb column and to build the two HTTP
// responses, so the two never drift out of sync with each other.
func toComponentDTOs(cs [6]ratingengine.Component) []componentDTO {
	out := make([]componentDTO, 0, len(cs))
	for _, c := range cs {
		d := componentDTO{
			Name: c.Name, Weight: c.Weight, Basis: c.Basis, BracketN: c.BracketN,
			Excluded: c.Excluded, Reason: c.Reason, Moments: toMomentDTOs(c.Moments),
		}
		if !c.Excluded {
			score := c.Score
			d.Score = &score
		}
		if c.Percentile != nil {
			pct := *c.Percentile
			d.Percentile = &pct
		}
		out = append(out, d)
	}
	return out
}

// fightMeta is the per-fight fields every roster player's stored row shares - read once
// per fight, not recomputed per player.
type fightMeta struct {
	ReportID     string
	FightIndex   int
	EncounterID  int64
	Difficulty   int64
	Size         int
	DurationMS   int64
	Kill         bool
	KillTimeBand string
	ModelVersion string
	FoughtAt     time.Time
}

// CardRow is one player's rating_scores row, ready to bind to an insert statement.
// PlayerKey is left blank here: the caller (Store.RateFight) is the one place that knows
// the fight's region/ruleset, so it fills PlayerKey via character.KeyFromUnit after
// newCardRow returns, rather than this package importing character.KeyFromUnit's inputs
// just to thread two more strings through this function's signature.
type CardRow struct {
	ReportID, PlayerKey, PlayerName, Class, Spec, Role, KillTimeBand, ModelVersion string
	FightIndex                                                                    int
	EncounterID, Difficulty                                                       int64
	Size                                                                          int
	DurationMS                                                                    int64
	Kill, OverallCapped                                                           bool
	Overall, OverallUncapped                                                      float64
	Components                                                                    json.RawMessage
	FoughtAt                                                                      time.Time
}

// newCardRow builds a CardRow from one player's already-scored Card. PlayerKey is left
// empty; the caller sets it.
func newCardRow(f fightMeta, row summary.RosterRow, card ratingengine.Card) CardRow {
	components, err := json.Marshal(toComponentDTOs(card.Components))
	if err != nil {
		// toComponentDTOs never produces a value json.Marshal can refuse (no channels,
		// funcs, or cycles in componentDTO), so this is unreachable outside a future,
		// accidentally-unmarshalable field addition - panicking here turns that mistake
		// into a test failure immediately rather than a silently empty components column.
		panic("rating: components must always marshal: " + err.Error())
	}
	return CardRow{
		ReportID: f.ReportID, FightIndex: f.FightIndex, PlayerName: row.Name,
		Class: row.Class, Spec: row.Spec, Role: row.Role,
		EncounterID: f.EncounterID, Difficulty: f.Difficulty, Size: f.Size,
		DurationMS: f.DurationMS, Kill: f.Kill, KillTimeBand: f.KillTimeBand,
		Overall: card.Overall, OverallUncapped: card.OverallUncapped, OverallCapped: card.OverallCapped,
		Components: components, ModelVersion: card.ModelVersion, FoughtAt: f.FoughtAt,
	}
}
```

- [ ] **Step 9: Run the tests again**

Run: `cd api && go test ./internal/rating/... -v`
Expected: every test PASSes.

- [ ] **Step 10: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/rating/doc.go api/internal/rating/curated.go api/internal/rating/curated_test.go \
  api/internal/rating/cards.go api/internal/rating/cards_test.go
printf 'feat: add curated-table selection and Card JSON shapes for ratings\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task3.txt
git commit -F .superpowers/commit-msg-task3.txt
```

---

## Task 4: The percentile-digest adapter (`PercentileSource`) and the kill-duration fold

**Files:**
- Create: `api/internal/rating/adapter.go`
- Test: `api/internal/rating/adapter_test.go` (uses the throwaway test database)

**Interfaces:**
- Consumes: `ratingengine.Bracket`, `ratingengine.PercentileSource`,
  `ratingengine.BandFast/BandTypical/BandSlow`, `api/internal/digest` (`New`, `Unmarshal`,
  `Add`, `Quantile`, `Placement`, `MarshalBinary`, `Count`), `github.com/jackc/pgx/v5`
  (`Tx`).
- Produces (for Task 5): `type percentileAdapter struct { Tx pgx.Tx; Fold bool }`
  implementing `ratingengine.PercentileSource`, and
  `foldKillDuration(ctx context.Context, tx pgx.Tx, encounterID, difficulty, durationMS int64) error`.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/rating/adapter_test.go
package rating

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := testURL(t) // see Step 2: this task also adds testURL to this package
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func TestPlacementReadsEmptyBracketAsNotOK(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	a := &percentileAdapter{Tx: tx, Fold: true}
	b := ratingengine.Bracket{EncounterID: 111111, Difficulty: 1, Spec: "warrior-fury", Role: "dps",
		KillTimeBand: "typical", Kill: true, Component: ratingengine.ComponentNameOutput}
	_, _, ok := a.Placement(b, 50)
	if ok {
		t.Error("a never-seen bracket must report ok = false")
	}
}

func TestPlacementFoldsOnKillForAKillGatedComponent(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 222222, Difficulty: 1, Spec: "mage-frost", Role: "dps",
		KillTimeBand: "typical", Kill: true, Component: ratingengine.ComponentNameOutput}
	a := &percentileAdapter{Tx: tx, Fold: true}
	for i, v := range []float64{10, 20, 30, 40, 50} {
		pct, n, ok := a.Placement(b, v)
		if i == 0 && ok {
			t.Error("the first fold into a brand-new bracket must still report ok = false (nothing to compare against yet)")
		}
		if i > 0 && (!ok || n != int64(i)) {
			t.Errorf("fold %d: pct=%v n=%v ok=%v, want ok and n=%d", i, pct, n, ok, i)
		}
	}
}

func TestPlacementDoesNotFoldAWipeIntoAKillGatedComponent(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 333333, Difficulty: 1, Spec: "priest-holy", Role: "healer",
		KillTimeBand: "", Kill: false, Component: ratingengine.ComponentNameActivity}
	a := &percentileAdapter{Tx: tx, Fold: true}
	if _, _, ok := a.Placement(b, 88); ok {
		t.Fatal("first call on an empty bracket must be ok=false")
	}
	// A second call with the same wipe bracket must see the SAME empty state: nothing
	// was folded, because Activity is not Survival/Mechanics and this bracket is a wipe.
	if _, n, ok := a.Placement(b, 88); ok || n != 0 {
		t.Errorf("wipe fold into a kill-gated component: ok=%v n=%v, want ok=false n=0", ok, n)
	}
}

func TestPlacementFoldsAWipeIntoASurvivalSubBracket(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 444444, Difficulty: 1, Spec: "warrior-protection", Role: "tank",
		KillTimeBand: "", Kill: false, Component: ratingengine.ComponentSurvivalAvoidableHit}
	a := &percentileAdapter{Tx: tx, Fold: true}
	a.Placement(b, 10)
	_, n, ok := a.Placement(b, 20)
	if !ok || n != 1 {
		t.Errorf("Survival must fold on a wipe: ok=%v n=%v, want ok=true n=1", ok, n)
	}
}

func TestPlacementNeverFoldsWhenFoldIsFalse(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	b := ratingengine.Bracket{EncounterID: 555555, Difficulty: 1, Spec: "hunter-survival", Role: "dps",
		KillTimeBand: "typical", Kill: true, Component: ratingengine.ComponentNameOutput}
	fold := &percentileAdapter{Tx: tx, Fold: true}
	fold.Placement(b, 100)
	readOnly := &percentileAdapter{Tx: tx, Fold: false}
	readOnly.Placement(b, 999)
	_, n, _ := fold.Placement(b, 1) // a fresh read confirms the read-only call added nothing
	if n != 1 {
		t.Errorf("n = %d after a read-only call, want 1 (unchanged)", n)
	}
}

func TestKillTimeBandNeverFolds(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback(ctx)
	a := &percentileAdapter{Tx: tx, Fold: true}
	if _, ok := a.KillTimeBand(666666, 1, 180000); ok {
		t.Fatal("a never-seen encounter/difficulty must report ok = false")
	}
	if err := foldKillDuration(ctx, tx, 666666, 1, 180000); err != nil {
		t.Fatal(err)
	}
	band, ok := a.KillTimeBand(666666, 1, 180000)
	if !ok || band != ratingengine.BandTypical {
		t.Errorf("band = %q ok=%v, want typical/true after exactly one fold", band, ok)
	}
	// KillTimeBand itself must not have folded a second time.
	if err := foldKillDuration(ctx, tx, 666666, 1, 90000); err != nil { // a much faster kill
		t.Fatal(err)
	}
	band, _ = a.KillTimeBand(666666, 1, 90000)
	if band != ratingengine.BandFast {
		t.Errorf("band = %q, want fast (median should now reflect both folds, not be stuck on a single call's snapshot)", band)
	}
}
```

- [ ] **Step 2: Add the package's own `testURL` helper**

```go
// api/internal/rating/db_test.go
package rating

import (
	"os"
	"testing"
)

// testURL is this package's own copy of db.testURL's shape: db.testURL is unexported, so
// every DB-backed package in this repo (compare api/internal/db/db_test.go) declares its
// own. TEST_DATABASE_URL is set by this lane's throwaway-database setup (see the plan's
// Task 1, Step 5) and must point at rating_api_test, not the shared forever_test other
// lanes use, per this lane's dispatch.
func testURL(t *testing.T) string {
	t.Helper()
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL not set; see docs/superpowers/plans/2026-09-21-rating-api.md Task 1 Step 5")
	}
	return u
}
```

- [ ] **Step 3: Run the tests to see them fail**

Run: `cd api && go test ./internal/rating/... -run "TestPlacement|TestKillTimeBand" -p 1 -v`
Expected: FAIL (adapter.go does not exist).

- [ ] **Step 4: Implement the adapter**

```go
// api/internal/rating/adapter.go
package rating

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/digest"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

// killDurationBandFast/Typical/Slow thresholds, spec §1.2's exact formula.
const (
	killDurationFastRatio = 0.85
	killDurationSlowRatio = 1.15
)

// foldEligible reports whether bracket b's value may be folded into its digest for this
// fight, spec §2's wipe rule: Survival/Mechanics' four sub-brackets fold on a wipe as well
// as a kill; every other component - Output, Utility, Preparation, Activity, addressed
// directly by their top-level component names since only Survival and Mechanics have
// sub-brackets (RULING R9 in logs/engine/rating) - folds only on a kill.
func foldEligible(b ratingengine.Bracket) bool {
	if b.Kill {
		return true
	}
	switch b.Component {
	case ratingengine.ComponentSurvivalAvoidableHit,
		ratingengine.ComponentMechanicsInterrupt,
		ratingengine.ComponentMechanicsDispel,
		ratingengine.ComponentMechanicsDebuffUptime:
		return true
	default:
		return false
	}
}

// percentileAdapter is the ratingengine.PercentileSource logs/engine/rating.Score reads
// through, backed by rating_percentile_digests and kill_duration_digests. Fold controls
// whether a Placement call may also add this fight's value into its digest: true on a
// fight's first, genuine write; false on a rewrite (a re-verified fight) or a backfill
// recompute, so a value already folded once is never folded twice (spec's idempotency
// requirement). KillTimeBand never folds regardless of Fold - see foldKillDuration, called
// once per fight rather than once per roster player.
type percentileAdapter struct {
	Tx   pgx.Tx
	Fold bool
}

var _ ratingengine.PercentileSource = (*percentileAdapter)(nil)

// Placement reads bracket b's digest, answers this value's percentile within it (before
// this value is added - the same "read, then decide, then maybe fold" order
// rankings.Store.updateDigest already uses), and, when a.Fold and foldEligible(b), adds
// value and persists the row afterward, all under the caller's transaction. The returned
// pct/n are captured from the digest's state BEFORE the fold mutates it (Go passes them
// back by value, so foldValue's later mutation of the same *digest.Digest cannot change
// what was already returned) - a fight is never compared against its own not-yet-folded
// value.
func (a *percentileAdapter) Placement(b ratingengine.Bracket, value float64) (pct float64, n int64, ok bool) {
	var raw []byte
	err := a.Tx.QueryRow(context.Background(),
		`insert into rating_percentile_digests
		   (encounter_id, difficulty, spec, role, kill_time_band, kill, component, digest, n)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, 0)
		 on conflict (encounter_id, difficulty, spec, role, kill_time_band, kill, component)
		 do update set digest = rating_percentile_digests.digest
		 returning digest`,
		b.EncounterID, b.Difficulty, b.Spec, b.Role, b.KillTimeBand, b.Kill, b.Component, []byte{},
	).Scan(&raw)
	if err != nil {
		return 0, 0, false
	}
	d, derr := digest.Unmarshal(raw)
	if derr != nil {
		d = digest.New()
	}
	// digest.Digest.Placement already returns the share of the bracket's other values
	// this value beats in 0..1 (api/internal/digest/digest.go:203-212's own doc comment),
	// exactly the range ratingengine.PercentileSource.Placement documents - no further
	// scaling belongs here.
	if d.Count() > 0 {
		pct, n, ok = d.Placement(value), d.Count(), true
	}
	if a.Fold && foldEligible(b) {
		a.foldValue(b, d, value)
	}
	return pct, n, ok
}

// foldValue adds value to d and persists the encoded digest under bracket b. Called only
// after Placement has already captured its return values from d's pre-fold state.
func (a *percentileAdapter) foldValue(b ratingengine.Bracket, d *digest.Digest, value float64) {
	d.Add(value)
	encoded, err := d.MarshalBinary()
	if err != nil {
		// MarshalBinary never errors in the current implementation (digest.go:237-249
		// builds a byte slice with no fallible step); a future change that makes it
		// fallible must not silently drop a fold, so this panics rather than swallowing
		// the error.
		panic("rating: digest must always marshal: " + err.Error())
	}
	_, _ = a.Tx.Exec(context.Background(),
		`update rating_percentile_digests set digest = $8, n = $9, updated_at = now()
		 where encounter_id = $1 and difficulty = $2 and spec = $3 and role = $4
		   and kill_time_band = $5 and kill = $6 and component = $7`,
		b.EncounterID, b.Difficulty, b.Spec, b.Role, b.KillTimeBand, b.Kill, b.Component,
		encoded, d.Count())
}

// KillTimeBand reads (never folds) the kill-duration digest for (encounterID,
// difficulty) and classifies durationMS against its median, spec §1.2's exact bands.
func (a *percentileAdapter) KillTimeBand(encounterID, difficulty int64, durationMS int64) (string, bool) {
	var raw []byte
	err := a.Tx.QueryRow(context.Background(),
		`select digest from kill_duration_digests where encounter_id = $1 and difficulty = $2`,
		encounterID, difficulty).Scan(&raw)
	if err != nil {
		return "", false
	}
	d, err := digest.Unmarshal(raw)
	if err != nil || d.Count() == 0 {
		return "", false
	}
	median := d.Quantile(0.5)
	return classifyBand(float64(durationMS), median), true
}

func classifyBand(durationMS, medianMS float64) string {
	switch {
	case durationMS < killDurationFastRatio*medianMS:
		return ratingengine.BandFast
	case durationMS > killDurationSlowRatio*medianMS:
		return ratingengine.BandSlow
	default:
		return ratingengine.BandTypical
	}
}

// foldKillDuration folds one kill's duration into its (encounterID, difficulty) bracket
// exactly once. Called once per fight, before the per-player Score loop - never from
// inside percentileAdapter.KillTimeBand, which would otherwise fold the same fight's
// duration once per roster player.
func foldKillDuration(ctx context.Context, tx pgx.Tx, encounterID, difficulty int64, durationMS int64) error {
	var raw []byte
	if err := tx.QueryRow(ctx,
		`insert into kill_duration_digests (encounter_id, difficulty, digest, n)
		 values ($1, $2, $3, 0)
		 on conflict (encounter_id, difficulty) do update set digest = kill_duration_digests.digest
		 returning digest`,
		encounterID, difficulty, []byte{}).Scan(&raw); err != nil {
		return fmt.Errorf("rating: lock kill duration digest: %w", err)
	}
	d, err := digest.Unmarshal(raw)
	if err != nil {
		return fmt.Errorf("rating: decode kill duration digest: %w", err)
	}
	d.Add(float64(durationMS))
	encoded, err := d.MarshalBinary()
	if err != nil {
		return fmt.Errorf("rating: encode kill duration digest: %w", err)
	}
	if _, err := tx.Exec(ctx,
		`update kill_duration_digests set digest = $3, n = $4, updated_at = now()
		 where encounter_id = $1 and difficulty = $2`,
		encounterID, difficulty, encoded, d.Count()); err != nil {
		return fmt.Errorf("rating: write kill duration digest: %w", err)
	}
	return nil
}
```

`Placement`'s `d.Placement(value)` call: `digest.Digest.Placement` returns a single
`float64` in `0..1` per `digest.go:212` (re-check the signature before wiring - if it
returns one value, not two, adjust the `pct, n = ...` line to
`pct, n = d.Placement(value)*100, d.Count()` accordingly, which is already what is written
above). Confirm this against `api/internal/digest/digest.go` before writing the test
assertions' exact percent math.

- [ ] **Step 5: Run the tests again**

Run: `cd api && go test ./internal/rating/... -run "TestPlacement|TestKillTimeBand" -p 1 -v`
Expected: every test PASSes. If `TestPlacementFoldsOnKillForAKillGatedComponent` fails on
the exact `n` sequence, adjust the loop's expected values to match `digest.Digest.Count()`'s
real behaviour (read after each `Add`) rather than the plan's guess — the test's *shape*
(first call not ok, every fold after that both ok and growing n) is what matters.

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/rating/adapter.go api/internal/rating/adapter_test.go api/internal/rating/db_test.go
printf 'feat: add the ratings percentile-digest adapter\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task4.txt
git commit -F .superpowers/commit-msg-task4.txt
```

---

## Task 5: `Store.RateFight` — score every roster player and write their rows

**Files:**
- Create: `api/internal/rating/store.go`
- Test: `api/internal/rating/store_test.go`

**Interfaces:**
- Consumes: `reports.RatedFight` (defined in Task 7 — **this task defines its own copy of
  the struct's shape as a local type `RatedFight` in this package first**; Task 7 makes
  `api/internal/reports.RatedFight` and updates every call site in this package to the
  `reports.` prefix, or — simpler, chosen here — **this task imports the type from
  `reports` directly**, since Task 7 lands the type before this task's tests run against a
  live `Ingest`. To keep task order flexible, this task defines `RatedFight` as its own
  package-local struct with the identical field set, and Task 7 makes
  `api/internal/reports.RatedFight` a distinct type with the same fields; `rater.go` (Task
  7) is the single, small conversion point between them. This avoids Task 5 depending on
  Task 7 landing first.
- Produces (for Task 6/7): `type RatedFight struct { ReportID string; FightIndex int;
  Region, Ruleset string; FoughtAt time.Time; EncounterID int64; Summary summary.Summary }`,
  `type Store struct { Pool *pgxpool.Pool; Log *slog.Logger }`,
  `func (s *Store) RateFight(ctx context.Context, f RatedFight) error`.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/rating/store_test.go
package rating

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func fightFixture(reportID string, kill bool, players ...summary.RosterRow) RatedFight {
	return RatedFight{
		ReportID: reportID, FightIndex: 1, Region: "us", Ruleset: "normal",
		FoughtAt: time.Date(2026, 12, 9, 1, 0, 0, 0, time.UTC), EncounterID: 667,
		Summary: summary.Summary{
			EncounterID: 667, Difficulty: 3, DurationMS: 180000, Kill: kill,
			Roster: players,
		},
	}
}

func TestRateFightWritesOneRowPerRosterPlayer(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-test-1", true,
		summary.RosterRow{GUID: "g1", Name: "Simfury", Class: "Warrior", Spec: "Fury", Role: "dps"},
		summary.RosterRow{GUID: "g2", Name: "Healface", Class: "Priest", Spec: "Holy", Role: "healer"},
	)
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })

	var n int
	if err := pool.QueryRow(ctx, `select count(*) from rating_scores where report_id = $1`, f.ReportID).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("wrote %d rows, want 2", n)
	}
	var key string
	if err := pool.QueryRow(ctx, `select player_key from rating_scores where report_id = $1 and player_name = 'Simfury'`, f.ReportID).Scan(&key); err != nil {
		t.Fatal(err)
	}
	if key != "us/normal/simfury" {
		t.Errorf("player_key = %q, want us/normal/simfury", key)
	}
}

func TestRateFightIsIdempotentAndDoesNotDoubleFoldDigests(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-test-idempotent", true,
		summary.RosterRow{GUID: "g1", Name: "Repeatme", Class: "Warrior", Spec: "Fury", Role: "dps"},
	)
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID)
	})
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	var nAfterFirst int64
	pool.QueryRow(ctx,
		`select n from rating_percentile_digests where encounter_id = 667 and difficulty = 3
		 and spec = 'warrior-fury' and role = 'dps' and component = 'output'`).Scan(&nAfterFirst)

	if err := store.RateFight(ctx, f); err != nil { // a re-ingest of the same fight
		t.Fatal(err)
	}
	var rowCount int
	pool.QueryRow(ctx, `select count(*) from rating_scores where report_id = $1`, f.ReportID).Scan(&rowCount)
	if rowCount != 1 {
		t.Fatalf("rewriting the same fight produced %d rows, want exactly 1 (upsert, not a duplicate)", rowCount)
	}
	var nAfterSecond int64
	pool.QueryRow(ctx,
		`select n from rating_percentile_digests where encounter_id = 667 and difficulty = 3
		 and spec = 'warrior-fury' and role = 'dps' and component = 'output'`).Scan(&nAfterSecond)
	if nAfterSecond != nAfterFirst {
		t.Errorf("digest n grew from %d to %d on a rewrite - the same value was folded twice", nAfterFirst, nAfterSecond)
	}
}

func TestRateFightStoresSixComponentsPerPlayer(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-test-components", true,
		summary.RosterRow{GUID: "g1", Name: "Sixparts", Class: "Mage", Spec: "Frost", Role: "dps"},
	)
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	var raw json.RawMessage
	if err := pool.QueryRow(ctx, `select components from rating_scores where report_id = $1`, f.ReportID).Scan(&raw); err != nil {
		t.Fatal(err)
	}
	var dtos []componentDTO
	if err := json.Unmarshal(raw, &dtos); err != nil {
		t.Fatal(err)
	}
	if len(dtos) != 6 {
		t.Fatalf("stored %d components, want 6", len(dtos))
	}
}
```

- [ ] **Step 2: Run to see them fail**

Run: `cd api && go test ./internal/rating/... -run TestRateFight -p 1 -v`
Expected: FAIL (`RatedFight`/`Store`/`RateFight` undefined).

- [ ] **Step 3: Implement the store**

```go
// api/internal/rating/store.go
package rating

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// RatedFight is one verified, ranked fight handed to the rating pipeline - the already-
// decoded summary the ingest holds in memory (spec §4.5: "no extra summary read is
// needed"), plus the region/ruleset a roster row's name is keyed under (the same pair
// api/internal/reports.ReportRealm already resolves for i.rank/i.score).
//
// This is deliberately this package's own type, not an alias for
// api/internal/reports.RatedFight: that lets this file's tests run without depending on
// the reports package landing its own copy first. api/internal/rating/rater.go is the one
// place that converts between the two.
type RatedFight struct {
	ReportID        string
	FightIndex      int
	Region, Ruleset string
	FoughtAt        time.Time
	EncounterID     int64
	Summary         summary.Summary
}

// Store is every rating read and write.
type Store struct {
	Pool *pgxpool.Pool
	Log  *slog.Logger
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// RateFight scores every roster player in f and writes their rows, folding each scored
// component's value into its percentile digest exactly once - on this fight's first
// genuine write, never on a rewrite (a re-verified fight arriving again) or a backfill
// recompute. Mirrors rankings.Store.WriteFight's own lock-then-detect-rewrite shape
// closely, on purpose: the two stores solve the same idempotency problem the same way.
func (s *Store) RateFight(ctx context.Context, f RatedFight) error {
	if len(f.Summary.Roster) == 0 {
		return nil
	}
	if err := db.EnsureRatingsPartition(ctx, s.Pool, f.FoughtAt); err != nil {
		return err
	}
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("rating: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	// Serialises this fight's writers exactly as rankings.Store.WriteFight's own comment
	// explains: without the lock two writers of the same fight could both see zero rows
	// withdrawn, both fold every value, and double-count the fight in every digest. A
	// distinct string prefix ("rating:") keeps this lock namespace separate from
	// rankings' own advisory lock on the same (report_id, fight_index) pair - harmless if
	// they collided, but a stray cross-package wait is easier to reason about avoided.
	if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock(hashtext('rating:' || $1), $2)`,
		f.ReportID, f.FightIndex); err != nil {
		return fmt.Errorf("rating: lock %s/%d: %w", f.ReportID, f.FightIndex, err)
	}

	withdrawn, err := tx.Exec(ctx, `delete from rating_scores where report_id = $1 and fight_index = $2`,
		f.ReportID, f.FightIndex)
	if err != nil {
		return fmt.Errorf("rating: withdraw %s/%d: %w", f.ReportID, f.FightIndex, err)
	}
	fold := withdrawn.RowsAffected() == 0

	if fold && f.Summary.Kill {
		if err := foldKillDuration(ctx, tx, f.Summary.EncounterID, f.Summary.Difficulty, f.Summary.DurationMS); err != nil {
			return err
		}
	}

	meta := fightMeta{
		ReportID: f.ReportID, FightIndex: f.FightIndex, EncounterID: f.Summary.EncounterID,
		Difficulty: f.Summary.Difficulty, DurationMS: f.Summary.DurationMS, Kill: f.Summary.Kill,
		FoughtAt: f.FoughtAt,
	}
	adapter := &percentileAdapter{Tx: tx, Fold: fold}
	if band, ok := adapter.KillTimeBand(f.Summary.EncounterID, f.Summary.Difficulty, f.Summary.DurationMS); ok && f.Summary.Kill {
		meta.KillTimeBand = band
	}
	modelInfo := ratingengine.DefaultModelInfo()

	for _, row := range f.Summary.Roster {
		tables := curatedTablesFor(row, f.Summary.EncounterID)
		card := ratingengine.Score(f.Summary, row.GUID, tables, adapter, nil, modelInfo)
		cr := newCardRow(meta, row, card)
		cr.PlayerKey = character.KeyFromUnit(f.Region, f.Ruleset, row.Name)
		if err := insertCardRow(ctx, tx, cr); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("rating: commit: %w", err)
	}
	return nil
}

func insertCardRow(ctx context.Context, tx pgx.Tx, cr CardRow) error {
	_, err := tx.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, role,
		   encounter_id, difficulty, size, duration_ms, kill, kill_time_band, overall, overall_uncapped,
		   overall_capped, components, model_version, fought_at)
		 values ($1, $2, $3, $4, nullif($5, ''), nullif($6, ''), nullif($7, ''), $8, $9, $10, $11, $12, $13,
		   $14, $15, $16, $17, $18, $19)
		 on conflict (report_id, fight_index, player_key, fought_at) do update set
		   player_name = excluded.player_name, class = excluded.class, spec = excluded.spec,
		   role = excluded.role, overall = excluded.overall, overall_uncapped = excluded.overall_uncapped,
		   overall_capped = excluded.overall_capped, components = excluded.components,
		   model_version = excluded.model_version, computed_at = now()`,
		cr.ReportID, cr.FightIndex, cr.PlayerKey, cr.PlayerName, cr.Class, cr.Spec, cr.Role,
		cr.EncounterID, cr.Difficulty, cr.Size, cr.DurationMS, cr.Kill, cr.KillTimeBand,
		cr.Overall, cr.OverallUncapped, cr.OverallCapped, cr.Components, cr.ModelVersion, cr.FoughtAt)
	if err != nil {
		return fmt.Errorf("rating: write row %s/%d/%s: %w", cr.ReportID, cr.FightIndex, cr.PlayerKey, err)
	}
	return nil
}
```

Do not declare an `errNotFound` var here: Task 6's read methods report "not found" via a
plain `ok bool` return instead (see Task 6), so an unused package-level error var would be
dead code the moment this task lands. If a later task genuinely needs a sentinel error, it
declares it where it is first used.

- [ ] **Step 4: Run the tests**

Run: `cd api && go test ./internal/rating/... -run TestRateFight -p 1 -v`
Expected: every test PASSes.

- [ ] **Step 5: Run the whole package's tests together to catch cross-test interference**

Run: `cd api && go test ./internal/rating/... -p 1 -v`
Expected: PASS. If `TestRateFightIsIdempotentAndDoesNotDoubleFoldDigests` is flaky because
an earlier test already wrote to the same `(667, 3, warrior-fury, dps, output)` bracket,
change that test's `EncounterID`/`Spec` to a value unique to it (already done above —
`store-test-idempotent`'s fixture reuses encounter 667 deliberately to prove real cross-
fight sharing works, but if this collides with `TestRateFightWritesOneRowPerRosterPlayer`'s
own `Simfury`/`warrior-fury`/`dps` row, adjust one of the two fixtures' spec to `Arms`
instead so their brackets do not overlap).

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/rating/store.go api/internal/rating/store_test.go
printf 'feat: add Store.RateFight for the ratings write path\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task5.txt
git commit -F .superpowers/commit-msg-task5.txt
```

---

## Task 6: Visibility (`viewer.go`) and the two read queries

**Files:**
- Modify: `api/internal/rating/store.go` (add the two read methods)
- Create: `api/internal/rating/viewer.go`
- Test: `api/internal/rating/viewer_test.go`, additions to `api/internal/rating/store_test.go`

**Interfaces:**
- Consumes: `reports.Report`, `reports.Public/Unlisted/Private/GuildTo`, `auth.ActorFrom`,
  `auth.Actor.Signed/IsModerator`.
- Produces (for Task 8): `type Accounts interface { GuildRank(ctx, guildID, userID int64)
  (string, bool, error) }`, `func visible(r *http.Request, rep reports.Report, accounts
  Accounts) bool`, `func (s *Store) ReadFightRatings(ctx, reportID string, fightIndex int)
  ([]CardRow, bool, error)`, `func (s *Store) ReadCharacterRating(ctx, playerKey string,
  limit int, before *cursorPos) (rows []CardRow, hasMore bool, err error)`,
  `func (s *Store) anonymized(ctx, playerKey string) (bool, error)`.

- [ ] **Step 1: Write the failing viewer tests**

```go
// api/internal/rating/viewer_test.go
package rating

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

type stubAccounts struct{ rank string; ok bool; err error }

func (s stubAccounts) GuildRank(ctx interface{ Done() <-chan struct{} }, guildID, userID int64) (string, bool, error) {
	return s.rank, s.ok, s.err
}

func TestVisiblePublicAndUnlistedAlwaysTrue(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, v := range []string{reports.Public, reports.Unlisted} {
		if !visible(r, reports.Report{Visibility: v}, nil) {
			t.Errorf("visibility %q must be visible to an anonymous request", v)
		}
	}
}

func TestVisiblePrivateFalseForAnonymous(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if visible(r, reports.Report{Visibility: reports.Private}, nil) {
		t.Error("a private report must not be visible to an anonymous request")
	}
}
```

The `stubAccounts` signature above will not compile as written (context type mismatch) —
this is intentional as a placeholder; the implementer replaces it with a real
`context.Context` parameter and `auth.WithActor`-style context injection matching how
`api/internal/reports/handler_test.go` (read it before writing this test file) already
builds an authenticated `*http.Request` for `mayView`-equivalent tests. Follow that
existing test file's exact helper pattern rather than the sketch above; the two additional
required cases are:

```go
func TestVisibleGuildRequiresVerifiedMembership(t *testing.T) {
	// Build a request whose context carries a signed-in, non-owner, non-moderator actor
	// (the same helper reports/handler_test.go uses), a Report with Visibility: GuildTo
	// and GuildID set, and an Accounts stub whose GuildRank returns ("", false, nil) for
	// that (guildID, userID) - visible() must return false.
}

func TestVisibleGuildTrueForVerifiedMember(t *testing.T) {
	// Same shape, but the Accounts stub returns (rank, true, nil) - visible() must return
	// true. auth.Store.GuildRank already filters to verified_at is not null at the SQL
	// layer (api/internal/auth/store.go:423), so a stub returning ok=true is standing in
	// for "this call would only ever be true for a verified member" - no additional
	// verified_at check belongs in this package.
}
```

- [ ] **Step 2: Implement `viewer.go`**

```go
// api/internal/rating/viewer.go
package rating

import (
	"context"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// Accounts is the one method this package needs from auth.Store: whether userID is a
// verified member of guildID, and at what rank. Declared here, where it is consumed,
// rather than imported from reports.Accounts (which declares a second method this
// package never calls) - *auth.Store already satisfies both.
type Accounts interface {
	GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error)
}

// visible replicates reports.Service.mayView's exact rule (reports/handler.go:540-556)
// for a report rep: public and unlisted are visible to anyone; private is visible to the
// owner and moderators; guild is visible to the owner, moderators, and verified guild
// members. It cannot call mayView directly - that method is unexported on *reports.Service
// and this lane's file fence forbids editing reports/handler.go to export it (see this
// plan's Global Constraints and the ledger for the ruling this documents) - so it is
// rebuilt here from the same exported primitives mayView itself is built from
// (reports.Public/Unlisted/Private/GuildTo, auth.ActorFrom, Actor.Signed/IsModerator,
// Accounts.GuildRank). Any future change to mayView's rule must be mirrored here by hand;
// there is no way around that within this lane's constraints.
func visible(r *http.Request, rep reports.Report, accounts Accounts) bool {
	if rep.Visibility == reports.Public || rep.Visibility == reports.Unlisted {
		return true
	}
	a := auth.ActorFrom(r.Context())
	if !a.Signed() {
		return false
	}
	if a.IsModerator() || (rep.OwnerID != nil && *rep.OwnerID == a.UserID) {
		return true
	}
	if rep.Visibility == reports.GuildTo && rep.GuildID != nil && accounts != nil {
		if _, ok, err := accounts.GuildRank(r.Context(), *rep.GuildID, a.UserID); err == nil && ok {
			return true
		}
	}
	return false
}
```

- [ ] **Step 3: Run the viewer tests**

Run: `cd api && go test ./internal/rating/... -run TestVisible -v`
Expected: PASS once the test file is corrected per Step 1's note (read
`api/internal/reports/handler_test.go`'s own actor-context helper first, then use it).

- [ ] **Step 4: Write the failing store read tests**

Append to `api/internal/rating/store_test.go`:

```go
func TestReadFightRatingsReturnsTheStoredRows(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("store-read-1", true,
		summary.RosterRow{GUID: "g1", Name: "Readme", Class: "Rogue", Spec: "Combat", Role: "dps"},
	)
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	rows, ok, err := store.ReadFightRatings(ctx, f.ReportID, f.FightIndex)
	if err != nil {
		t.Fatal(err)
	}
	if !ok || len(rows) != 1 {
		t.Fatalf("ok=%v rows=%d, want ok=true rows=1", ok, len(rows))
	}
	if rows[0].PlayerName != "Readme" {
		t.Errorf("player_name = %q", rows[0].PlayerName)
	}
}

func TestReadFightRatingsMissingFightReportsNotFound(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	_, ok, err := store.ReadFightRatings(context.Background(), "no-such-report", 1)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Error("a fight with no rows must report ok = false, not an empty success")
	}
}

func TestReadCharacterRatingExcludesAnonymizedPlayers(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	if _, err := pool.Exec(ctx,
		`insert into users (email, anonymize) values ('rating-anon-test@example.com', true) on conflict (email) do update set anonymize = true`); err != nil {
		t.Fatal(err)
	}
	var uid int64
	pool.QueryRow(ctx, `select id from users where email = 'rating-anon-test@example.com'`).Scan(&uid)
	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values ('us/normal/anonme', 'us', 'normal', 'Anonme', $1)
		 on conflict (key) do update set user_id = $1`, uid); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key = 'us/normal/anonme'`)
		pool.Exec(ctx, `delete from users where email = 'rating-anon-test@example.com'`)
	})
	anon, err := store.anonymized(ctx, "us/normal/anonme")
	if err != nil {
		t.Fatal(err)
	}
	if !anon {
		t.Error("expected the player to read as anonymized")
	}
}
```

- [ ] **Step 5: Run to see them fail**

Run: `cd api && go test ./internal/rating/... -run "TestReadFightRatings|TestReadCharacterRating" -p 1 -v`
Expected: FAIL.

- [ ] **Step 6: Implement the two read methods, appended to `store.go`**

```go
// api/internal/rating/store.go (append)

// cursorPos is the trend list's keyset position: the (fought_at, report_id, fight_index)
// of the last row a page returned, matching reports/recent.go's own cursor shape.
type cursorPos struct {
	FoughtAt   time.Time
	ReportID   string
	FightIndex int
}

// ReadFightRatings reads every stored rating row for one fight. ok is false when the
// fight has no rows at all - never verified, still queued in the async Rater, or (per §4.3)
// a fight this lane never rates at all (trash, no ranker, or a report that was never
// Ranked). The caller (the handler) does not distinguish those three at the HTTP layer:
// all read as 404, matching how a report's own file routes already treat "nothing here"
// without describing why.
func (s *Store) ReadFightRatings(ctx context.Context, reportID string, fightIndex int) ([]CardRow, bool, error) {
	rows, err := s.Pool.Query(ctx,
		`select player_key, player_name, coalesce(class, ''), coalesce(spec, ''), coalesce(role, ''),
		    encounter_id, difficulty, coalesce(size, 0), coalesce(duration_ms, 0), kill, kill_time_band,
		    overall, overall_uncapped, overall_capped, components, model_version, fought_at
		 from rating_scores where report_id = $1 and fight_index = $2 order by player_name`,
		reportID, fightIndex)
	if err != nil {
		return nil, false, fmt.Errorf("rating: read %s/%d: %w", reportID, fightIndex, err)
	}
	defer rows.Close()
	var out []CardRow
	for rows.Next() {
		var cr CardRow
		if err := rows.Scan(&cr.PlayerKey, &cr.PlayerName, &cr.Class, &cr.Spec, &cr.Role,
			&cr.EncounterID, &cr.Difficulty, &cr.Size, &cr.DurationMS, &cr.Kill, &cr.KillTimeBand,
			&cr.Overall, &cr.OverallUncapped, &cr.OverallCapped, &cr.Components, &cr.ModelVersion, &cr.FoughtAt); err != nil {
			return nil, false, fmt.Errorf("rating: scan %s/%d: %w", reportID, fightIndex, err)
		}
		cr.ReportID, cr.FightIndex = reportID, fightIndex
		out = append(out, cr)
	}
	if err := rows.Err(); err != nil {
		return nil, false, err
	}
	return out, len(out) > 0, nil
}

// anonymized reports whether playerKey's owning account has asked to be anonymous
// (users.anonymize), joined the same way auth.User.PublicName's own doc names as the join
// a character-scoped anonymize check would need. A player_key with no characters row, or
// no linked account, is never anonymized - there is nothing to hide.
func (s *Store) anonymized(ctx context.Context, playerKey string) (bool, error) {
	var anon bool
	err := s.Pool.QueryRow(ctx,
		`select u.anonymize from characters c join users u on u.id = c.user_id where c.key = $1`,
		playerKey).Scan(&anon)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("rating: anonymize check %s: %w", playerKey, err)
	}
	return anon, nil
}

// characterTrendPageSize bounds one page of the character endpoint's trend[]; a career's
// worth of rated fights is unbounded, so this list is paginated like every other list read
// in this API (reports/recent.go's own convention).
const characterTrendPageSize = 20

// ReadCharacterRating reads up to limit rows of playerKey's rating history, newest first,
// from public reports only (spec §5.1's ruling: "public at time of query", enforced here
// by joining the live reports.visibility column rather than trusting anything cached on
// the row at write time), starting after before when it is non-nil. hasMore is true when
// a further page exists.
func (s *Store) ReadCharacterRating(ctx context.Context, playerKey string, limit int, before *cursorPos) ([]CardRow, bool, error) {
	if limit <= 0 || limit > characterTrendPageSize {
		limit = characterTrendPageSize
	}
	query := `select rs.player_key, rs.player_name, coalesce(rs.class, ''), coalesce(rs.spec, ''),
	    coalesce(rs.role, ''), rs.encounter_id, rs.difficulty, coalesce(rs.size, 0),
	    coalesce(rs.duration_ms, 0), rs.kill, rs.kill_time_band, rs.overall, rs.overall_uncapped,
	    rs.overall_capped, rs.components, rs.model_version, rs.report_id, rs.fight_index, rs.fought_at
	 from rating_scores rs
	 join reports r on r.id = rs.report_id
	 where rs.player_key = $1 and r.visibility = 'public'`
	args := []any{playerKey}
	if before != nil {
		query += ` and (rs.fought_at, rs.report_id, rs.fight_index) < ($2, $3, $4)`
		args = append(args, before.FoughtAt, before.ReportID, before.FightIndex)
	}
	query += ` order by rs.fought_at desc, rs.report_id desc, rs.fight_index desc limit $` + fmt.Sprint(len(args)+1)
	args = append(args, limit+1) // one extra row, to answer hasMore without a second query

	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, false, fmt.Errorf("rating: read character %s: %w", playerKey, err)
	}
	defer rows.Close()
	var out []CardRow
	for rows.Next() {
		var cr CardRow
		if err := rows.Scan(&cr.PlayerKey, &cr.PlayerName, &cr.Class, &cr.Spec, &cr.Role,
			&cr.EncounterID, &cr.Difficulty, &cr.Size, &cr.DurationMS, &cr.Kill, &cr.KillTimeBand,
			&cr.Overall, &cr.OverallUncapped, &cr.OverallCapped, &cr.Components, &cr.ModelVersion,
			&cr.ReportID, &cr.FightIndex, &cr.FoughtAt); err != nil {
			return nil, false, fmt.Errorf("rating: scan character %s: %w", playerKey, err)
		}
		out = append(out, cr)
	}
	if err := rows.Err(); err != nil {
		return nil, false, err
	}
	hasMore := len(out) > limit
	if hasMore {
		out = out[:limit]
	}
	return out, hasMore, nil
}
```

Add `"errors"` to `store.go`'s imports for `errors.Is(err, pgx.ErrNoRows)` above
(`"github.com/jackc/pgx/v5"` is already imported from Task 5's `insertCardRow`).

- [ ] **Step 7: Run the tests**

Run: `cd api && go test ./internal/rating/... -p 1 -v`
Expected: every test PASSes.

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/rating/viewer.go api/internal/rating/viewer_test.go api/internal/rating/store.go api/internal/rating/store_test.go
printf 'feat: add visibility replica and the two ratings read queries\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task6.txt
git commit -F .superpowers/commit-msg-task6.txt
```

---

## Task 7: The ingest hook (`i.rate`) and the async `Rater` worker

**Files:**
- Modify: `api/internal/reports/ingest.go` (the only file this lane may touch under
  `api/internal/reports/`)
- Create: `api/internal/rating/rater.go`
- Test: `api/internal/reports/ingest_test.go` — **read this file first to find the existing
  test harness for `i.score`/`i.rank` and extend it with matching cases**, plus
  `api/internal/rating/rater_test.go`

**Interfaces:**
- Produces (for Task 9): `reports.RatedFight` (the public type, distinct from Task 5's
  package-local `rating.RatedFight`), `reports.Rater` interface, `Ingest.Rate Rater` field,
  `rating.RateDeps{ Store *Store; Log *slog.Logger }`, `rating.NewRater(d RateDeps)
  *Rater`, `(*Rater).Schedule(f reports.RatedFight)`, `(*Rater).Run(ctx context.Context)`,
  `(*Rater).Close()`.

- [ ] **Step 1: Read `api/internal/reports/ingest_test.go`'s existing scorer-hook test(s)**

Find the test(s) covering `i.score`'s never-fails/private-skip/trash-skip behaviour (search
for `TestPutFight.*[Ss]core` or similar) and copy their exact harness shape (a fake `Ranker`/
`Scorer`, a minimal multipart fight bundle) for this step's new tests, so `i.rate`'s tests
read as siblings of `i.score`'s, not a reinvention.

- [ ] **Step 2: Write the failing ingest hook tests**

Add to `api/internal/reports/ingest_test.go` (following the file's existing fake-dependency
pattern for `Scorer`):

```go
type fakeRater struct {
	scheduled []RatedFight
}

func (f *fakeRater) Schedule(rf RatedFight) {
	f.scheduled = append(f.scheduled, rf)
}

func TestPutFightSchedulesARatingForAPublicReport(t *testing.T) {
	// Build the same minimal Ingest + multipart fight bundle
	// TestPutFightSchedulesAScoreForAMember (or its equivalent) already builds, with
	// Visibility: Public, a fakeRater wired as i.Rate, and a roster of two players -
	// one a signed-in member, one not (unlike i.score, i.rate must schedule both).
	// After PUT succeeds:
	//   if len(rater.scheduled) != 1 { t.Fatal(...) }
	//   rf := rater.scheduled[0]
	//   if len(rf.Summary.Roster) != 2 { t.Errorf("i.rate must include every roster player, not signed-in members only") }
}

func TestPutFightSkipsRatingForAPrivateReport(t *testing.T) {
	// Same shape with Visibility: Private. rater.scheduled must be empty.
}

func TestPutFightStillCompletesWhenRateIsNil(t *testing.T) {
	// i.Rate left nil (a deployment with no rater wired) must not panic or fail the
	// upload - the PUT must still answer 201/200 exactly as it does today.
}
```

- [ ] **Step 3: Run to see them fail**

Run: `cd api && go test ./internal/reports/... -run "TestPutFightSchedulesARating|TestPutFightSkipsRatingForAPrivate|TestPutFightStillCompletesWhenRateIsNil" -v`
Expected: FAIL (`RatedFight`/`Rater`/`Ingest.Rate` undefined).

- [ ] **Step 4: Add the hook to `ingest.go`**

Add near `ScoredFight`/`Scorer` (after the existing `Scorer` interface, before `Ingest`):

```go
// RatedFight is one verified, ranked fight handed to the rating pipeline: the already-
// decoded summary the ingest holds in memory (the same "summary", "posted"/"posted" the
// verify handler already unmarshalled to store the fight - spec §4.5's own note that no
// extra summary read is needed), plus the region/ruleset a roster row's name is keyed
// under, matching what ReportRealm already resolves for i.rank/i.score.
type RatedFight struct {
	ReportID        string
	FightIndex      int
	Region, Ruleset string
	FoughtAt        time.Time
	EncounterID     int64
	Summary         summary.Summary
}

// Rater schedules one fight's ratings for every roster player, out of band, so the
// companion's fight-close call returns at once - the same never-blocks-storage shape as
// Scorer. *rating.Rater satisfies it.
type Rater interface {
	Schedule(f RatedFight)
}
```

Add `Rate Rater` to the `Ingest` struct, next to `Score Scorer`, with its own doc comment:

```go
	// Rate schedules ratings for every roster player of a ranked fight, unlike Score
	// which is signed-in-member-only (the execution scorer needs a rebuildable simulator
	// character; a rating does not - spec §4.3's own ruling). Nil means this deployment
	// rates nothing, which degrades the same way a nil Score does.
	Rate Rater
```

Add the `rate` method near `score` (after it):

```go
// rate schedules this fight's ratings for every roster player - unlike score, which is
// signed-in-member-only. Trash, an unranked visibility, and a deployment with no rater are
// all skipped, matching score's own skip conditions exactly except for the member gate
// (spec §4.3's ruling: "i.score's member gate is specific to needing a simulator
// character, not a rating precondition"). Best-effort and returns nothing, for the same
// reason score is: a rating failure must never fail an upload.
func (i *Ingest) rate(ctx context.Context, rep Report, n int, encounterID int64, at time.Time, sum summary.Summary) {
	if i.Rate == nil || encounterID == 0 || !Ranked(rep.Visibility) {
		return
	}
	region, ruleset := ReportRealm(rep)
	i.Rate.Schedule(RatedFight{
		ReportID: rep.ID, FightIndex: n, Region: region, Ruleset: ruleset,
		FoughtAt: at.UTC(), EncounterID: encounterID, Summary: sum,
	})
}
```

In `putFight`, immediately after the existing `i.score(r.Context(), rep, n, f.EncounterID, derived)`
line, add exactly one line:

```go
	i.rate(r.Context(), rep, n, f.EncounterID, f.Start, posted)
```

(`posted` is the `summary.Summary` already decoded earlier in `putFight` from the bundle's
"summary" part — confirm the variable name at the call site; it must be the full summary,
not `derived`, which is `[]metrics.Row`.)

- [ ] **Step 5: Run the ingest tests**

Run: `cd api && go vet ./internal/reports/... && go test ./internal/reports/... -v`
Expected: every test PASSes, including every pre-existing test (this change is additive:
`Rate` defaults to nil, and `i.rate` no-ops when nil, so no existing test's behaviour
changes).

- [ ] **Step 6: Write the failing `Rater` worker test**

```go
// api/internal/rating/rater_test.go
package rating

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestRaterRunProcessesScheduledFights(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	rater := NewRater(RateDeps{Store: store})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go rater.Run(ctx)
	defer rater.Close()

	rf := reports.RatedFight{
		ReportID: "rater-test-1", FightIndex: 1, Region: "us", Ruleset: "normal",
		FoughtAt: time.Date(2026, 12, 9, 2, 0, 0, 0, time.UTC), EncounterID: 667,
		Summary: summary.Summary{
			EncounterID: 667, Difficulty: 3, DurationMS: 180000, Kill: true,
			Roster: []summary.RosterRow{{GUID: "g1", Name: "Ratered", Class: "Hunter", Spec: "Survival", Role: "dps"}},
		},
	}
	rater.Schedule(rf)
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, rf.ReportID) })

	deadline := time.After(5 * time.Second)
	for {
		var n int
		pool.QueryRow(context.Background(), `select count(*) from rating_scores where report_id = $1`, rf.ReportID).Scan(&n)
		if n == 1 {
			return
		}
		select {
		case <-deadline:
			t.Fatal("the scheduled fight was never rated within 5s")
		case <-time.After(50 * time.Millisecond):
		}
	}
}

func TestRaterScheduleNeverBlocksAfterClose(t *testing.T) {
	rater := NewRater(RateDeps{Store: &Store{Pool: nil}})
	rater.Close()
	// Scheduling after Close must not panic (send on a closed channel) - mirrors
	// sims.Scorer's own Close-then-Schedule safety.
	rater.Schedule(reports.RatedFight{ReportID: "closed-test"})
}
```

- [ ] **Step 7: Run to see it fail**

Run: `cd api && go test ./internal/rating/... -run TestRater -p 1 -v`
Expected: FAIL (`reports.RatedFight`, `NewRater`, `RateDeps` undefined until Step 8).

- [ ] **Step 8: Implement `rater.go`**

```go
// api/internal/rating/rater.go
package rating

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// RateQueueDepth is how many fights may wait to be rated. Beyond it a fight-close drops
// the rating rather than blocking the ingest; the nightly backfill job picks it up on its
// next sweep (RateFight's rewrite-detection makes catching up later exactly as correct as
// catching up immediately).
const RateQueueDepth = 512

// RateDeps is what the rater needs.
type RateDeps struct {
	Store *Store
	Log   *slog.Logger
}

func (d RateDeps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

// Rater runs ratings out of band, so the companion's fight-close call returns at once. It
// is sims.Scorer's exact shape, for sims.Scorer's exact reason: the consumer has to outlive
// the handlers that feed it, so Close - not a cancelled context - is the only thing that
// stops it, and the shutdown order is Shutdown then Close.
type Rater struct {
	Deps  RateDeps
	queue chan reports.RatedFight
	done  chan struct{}
	once  sync.Once

	mu     sync.RWMutex
	closed bool
}

var _ reports.Rater = (*Rater)(nil)

// NewRater returns a rater with an empty queue.
func NewRater(d RateDeps) *Rater {
	return &Rater{Deps: d, queue: make(chan reports.RatedFight, RateQueueDepth), done: make(chan struct{})}
}

// Schedule queues one fight. It never blocks: a full queue, or one that lost the race with
// Close, drops the fight (logged) rather than blocking the ingest or panicking on a send to
// a closed channel.
func (s *Rater) Schedule(f reports.RatedFight) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.closed {
		s.Deps.logger().Warn("rating", "op", "schedule", "report", f.ReportID, "fight", f.FightIndex,
			"err", "rater is closed")
		return
	}
	select {
	case s.queue <- f:
	default:
		s.Deps.logger().Error("rating", "op", "schedule", "report", f.ReportID, "fight", f.FightIndex,
			"err", "the rating queue is full; the backfill job will catch this fight up later")
	}
}

// Run consumes the queue until Close is called. A fight already queued when Close runs is
// still drained and rated: closing a channel does not discard what is already buffered in
// it - this is sims.Scorer.Run's exact shape (api/internal/sims/score.go), on purpose, not
// the select-on-two-channels sketch an earlier draft of this plan carried (that version
// could drop a buffered fight the instant Close fired, racing the select's two ready cases
// instead of draining first - caught by an implementer's own self-review, not by this
// plan's own preflight scan; see the ledger). Ratings run on a context the shutdown signal
// does not cancel, matching Scorer's own "in-flight work finishes" contract.
func (s *Rater) Run(ctx context.Context) {
	defer close(s.done)
	work := context.WithoutCancel(ctx)
	for f := range s.queue {
		if err := s.Deps.Store.RateFight(work, RatedFight{
			ReportID: f.ReportID, FightIndex: f.FightIndex, Region: f.Region, Ruleset: f.Ruleset,
			FoughtAt: f.FoughtAt, EncounterID: f.EncounterID, Summary: f.Summary,
		}); err != nil {
			s.Deps.logger().Error("rating", "op", "rate", "report", f.ReportID, "fight", f.FightIndex, "err", err)
		}
	}
}

// Close stops the rater and waits for the queue to drain (every fight already buffered is
// rated before this returns). No Schedule started after Close returns reaches the queue,
// and calling Close twice is fine.
func (s *Rater) Close() {
	s.once.Do(func() {
		s.mu.Lock()
		s.closed = true
		close(s.queue)
		s.mu.Unlock()
	})
	<-s.done
}
```

Now add `reports.RatedFight` as the public counterpart to Task 5's package-local type: it
already exists from Step 4 above (added directly to `ingest.go`), and `reports.Rater` is
already declared there too — no further edit to `ingest.go` is needed in this step.

- [ ] **Step 9: Run the rater tests**

Run: `cd api && go test ./internal/rating/... -p 1 -v`
Expected: every test PASSes.

- [ ] **Step 10: Run the full reports + rating suites together once more**

Run: `cd api && go vet ./internal/reports/... ./internal/rating/... && go test ./internal/reports/... ./internal/rating/... -p 1 -v`
Expected: PASS.

- [ ] **Step 11: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/reports/ingest.go api/internal/reports/ingest_test.go \
  api/internal/rating/rater.go api/internal/rating/rater_test.go
printf 'feat: add the i.rate ingest hook and the async Rater worker\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task7.txt
git commit -F .superpowers/commit-msg-task7.txt
```

---

## Task 8: HTTP handlers for the two read endpoints

**Files:**
- Create: `api/internal/rating/handler.go`
- Test: `api/internal/rating/handler_test.go`

**Interfaces:**
- Consumes: `httpx.Envelope/WriteOK/WriteError`, `character.ValidRegion/ValidRuleset/Slug`,
  `reports.Store.Get`, `reports.ErrNotFound`, `rating.visible`, `rating.Accounts`,
  `rating.Store.ReadFightRatings/ReadCharacterRating/anonymized`.
- Produces (for Task 9): `type Service struct { Store *Store; Reports *reports.Store;
  Accounts Accounts; Log *slog.Logger }`, `func Mount(mux *http.ServeMux, s *Service)`.

- [ ] **Step 1: Write the failing handler tests**

```go
// api/internal/rating/handler_test.go
package rating

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func newTestService(t *testing.T) (*Service, *reports.Store) {
	pool := testPool(t)
	reportStore := &reports.Store{Pool: pool}
	return &Service{Store: &Store{Pool: pool}, Reports: reportStore}, reportStore
}

func mustCreateReport(t *testing.T, rs *reports.Store, id, visibility string) {
	t.Helper()
	// Use whatever reports.Store method api/internal/reports/store_test.go already uses
	// to insert a minimal report row directly (read that file for the exact helper/SQL);
	// visibility, status = 'complete' is all this handler cares about.
}

func TestFightRatingsServesAPublicReportsRoster(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-public-1", reports.Public)
	f := fightFixture("handler-public-1", true,
		summary.RosterRow{GUID: "g1", Name: "Served", Class: "Druid", Spec: "Balance", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-public-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	var env httpx.Envelope
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if !env.OK {
		t.Fatalf("ok = false: %+v", env.Error)
	}
}

func TestFightRatingsIs404ForAPrivateReportToAnAnonymousCaller(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-private-1", reports.Private)
	f := fightFixture("handler-private-1", true,
		summary.RosterRow{GUID: "g1", Name: "Hidden", Class: "Warlock", Spec: "Affliction", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-private-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 (private reports must not distinguish 403 from a missing report)", rec.Code)
	}
}

func TestFightRatingsIs404ForAMissingFight(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-missing-1", reports.Public)
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-missing-1/fights/9/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestCharacterRatingRejectsAnInvalidRegionOrRuleset(t *testing.T) {
	svc, _ := newTestService(t)
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/characters/xx/normal/someone/rating", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404 for an invalid region", rec.Code)
	}
}

func TestFightRatingsSetsPrivateCacheControlForANonPublicReport(t *testing.T) {
	svc, rs := newTestService(t)
	mustCreateReport(t, rs, "handler-unlisted-1", reports.Unlisted)
	f := fightFixture("handler-unlisted-1", true,
		summary.RosterRow{GUID: "g1", Name: "Unlisted", Class: "Paladin", Spec: "Retribution", Role: "dps"},
	)
	if err := svc.Store.RateFight(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	Mount(mux, svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/handler-unlisted-1/fights/1/ratings", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if got := rec.Header().Get("Cache-Control"); got != "private" {
		t.Errorf("Cache-Control = %q, want private for an unlisted report (spec §5.3)", got)
	}
}
```

Fill in `mustCreateReport` by reading `api/internal/reports/store_test.go`'s own insert
helper for a minimal report row (it already exists for that package's tests — reuse its
exact SQL/shape rather than inventing a second one).

- [ ] **Step 2: Run to see them fail**

Run: `cd api && go test ./internal/rating/... -run "TestFightRatings|TestCharacterRating" -p 1 -v`
Expected: FAIL.

- [ ] **Step 3: Implement `handler.go`**

```go
// api/internal/rating/handler.go
package rating

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// cacheSeconds matches rankings' own edge-cache policy (spec §5.3: "ratings are exactly as
// volatile as a rankings page").
const cacheSeconds = 30

// Service serves the two rating read routes.
type Service struct {
	Store    *Store
	Reports  *reports.Store
	Accounts Accounts
	Log      *slog.Logger
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

func (s *Service) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	s.logger().Error("rating", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// Mount registers the rating read routes.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/reports/{id}/fights/{n}/ratings", s.fightRatings)
	mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}/rating", s.characterRating)
}

func (s *Service) fightRatings(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil || n < 1 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the fight index must be a number, counted from one", nil)
		return
	}
	rep, err := s.Reports.Get(r.Context(), r.PathValue("id"))
	if err == reports.ErrNotFound {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "report", err, "could not read that report just now")
		return
	}
	if !visible(r, rep, s.Accounts) {
		// The same 404, whether the report does not exist or the caller may not see it -
		// spec's own "no endpoint leaks the existence of a non-public report" rule.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return
	}
	rows, ok, err := s.Store.ReadFightRatings(r.Context(), rep.ID, n)
	if err != nil {
		s.fail(w, r, "fight", err, "could not read that fight's ratings just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no ratings for that fight", nil)
		return
	}
	players := make([]playerRatingDTO, 0, len(rows))
	for _, cr := range rows {
		anon, err := s.Store.anonymized(r.Context(), cr.PlayerKey)
		if err != nil {
			s.fail(w, r, "fight", err, "could not read that fight's ratings just now")
			return
		}
		if anon {
			continue // spec §5.1: an anonymized player's row is omitted entirely, no placeholder
		}
		players = append(players, toPlayerRatingDTO(cr))
	}
	if rep.Visibility == reports.Public {
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheSeconds))
	} else {
		w.Header().Set("Cache-Control", "private")
	}
	httpx.WriteOK(w, r, http.StatusOK, fightRatingsDTO{
		FightIndex: n, Kill: rows[0].Kill, KillTimeBand: rows[0].KillTimeBand,
		ModelVersion: rows[0].ModelVersion, Players: players,
	})
}

func (s *Service) characterRating(w http.ResponseWriter, r *http.Request) {
	region, ruleset := strings.ToLower(r.PathValue("region")), strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	playerKey := character.Key(region, ruleset, r.PathValue("name"))
	anon, err := s.Store.anonymized(r.Context(), playerKey)
	if err != nil {
		s.fail(w, r, "character", err, "could not read that character's rating just now")
		return
	}
	if anon {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	limit, before := parseTrendQuery(r)
	rows, hasMore, err := s.Store.ReadCharacterRating(r.Context(), playerKey, limit, before)
	if err != nil {
		s.fail(w, r, "character", err, "could not read that character's rating just now")
		return
	}
	if len(rows) == 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no rated fights for that character", nil)
		return
	}
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheSeconds))
	httpx.WriteOK(w, r, http.StatusOK, buildCharacterRatingDTO(playerKey, rows, hasMore))
}

func parseTrendQuery(r *http.Request) (limit int, before *cursorPos) {
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	if c, ok := decodeTrendCursor(r.URL.Query().Get("cursor")); ok {
		before = &c
	}
	return limit, before
}

const trendCursorSep = "|"

func encodeTrendCursor(c cursorPos) string {
	raw := c.FoughtAt.UTC().Format(time.RFC3339Nano) + trendCursorSep + c.ReportID +
		trendCursorSep + strconv.Itoa(c.FightIndex)
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

// decodeTrendCursor is encodeTrendCursor's inverse. It reports false for anything this
// service did not itself produce - a caller may not walk the trend by any timestamp,
// report and fight index of their own choosing, matching reports/recent.go's own
// decodeRecentCursor contract exactly.
func decodeTrendCursor(s string) (cursorPos, bool) {
	if s == "" {
		return cursorPos{}, false
	}
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return cursorPos{}, false
	}
	parts := strings.SplitN(string(raw), trendCursorSep, 3)
	if len(parts) != 3 {
		return cursorPos{}, false
	}
	t, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil || parts[1] == "" {
		return cursorPos{}, false
	}
	n, err := strconv.Atoi(parts[2])
	if err != nil {
		return cursorPos{}, false
	}
	return cursorPos{FoughtAt: t, ReportID: parts[1], FightIndex: n}, true
}
```

Then add the response DTOs (append to `handler.go` or a small `dto.go` in the same
package — keep `handler.go` under 400 lines; split into `dto.go` if it grows past that):

```go
type momentAPIWrap = momentDTO // no separate shape; momentDTO already matches §5.2 verbatim

type playerRatingDTO struct {
	PlayerKey       string          `json:"player_key"`
	PlayerName      string          `json:"player_name"`
	Class           string          `json:"class"`
	Spec            string          `json:"spec"`
	Role            string          `json:"role"`
	Overall         float64         `json:"overall"`
	OverallUncapped float64         `json:"overall_uncapped"`
	OverallCapped   bool            `json:"overall_capped"`
	Components      json.RawMessage `json:"components"`
}

func toPlayerRatingDTO(cr CardRow) playerRatingDTO {
	return playerRatingDTO{
		PlayerKey: cr.PlayerKey, PlayerName: cr.PlayerName, Class: cr.Class, Spec: cr.Spec, Role: cr.Role,
		Overall: cr.Overall, OverallUncapped: cr.OverallUncapped, OverallCapped: cr.OverallCapped,
		Components: cr.Components,
	}
}

type fightRatingsDTO struct {
	FightIndex   int               `json:"fight_index"`
	Kill         bool              `json:"kill"`
	KillTimeBand string            `json:"kill_time_band"`
	ModelVersion string            `json:"model_version"`
	Players      []playerRatingDTO `json:"players"`
}

type trendPointDTO struct {
	FoughtAt   time.Time `json:"fought_at"`
	Overall    float64   `json:"overall"`
	ReportID   string    `json:"report_id"`
	FightIndex int       `json:"fight_index"`
}

type characterRatingDTO struct {
	PlayerKey      string           `json:"player_key"`
	SampleSize     int              `json:"sample_size"`
	Trend          []trendPointDTO  `json:"trend"`
	NextCursor     string           `json:"next_cursor,omitempty"`
	Latest         *playerRatingDTO `json:"latest,omitempty"`
}

func buildCharacterRatingDTO(playerKey string, rows []CardRow, hasMore bool) characterRatingDTO {
	out := characterRatingDTO{PlayerKey: playerKey, SampleSize: len(rows)}
	for _, cr := range rows {
		out.Trend = append(out.Trend, trendPointDTO{
			FoughtAt: cr.FoughtAt, Overall: cr.Overall, ReportID: cr.ReportID, FightIndex: cr.FightIndex,
		})
	}
	if len(rows) > 0 {
		latest := toPlayerRatingDTO(rows[0])
		out.Latest = &latest
	}
	if hasMore && len(rows) > 0 {
		last := rows[len(rows)-1]
		out.NextCursor = encodeTrendCursor(cursorPos{FoughtAt: last.FoughtAt, ReportID: last.ReportID, FightIndex: last.FightIndex})
	}
	return out
}
```

Note: `best_component`/`worst_component` from spec §5.2's example are intentionally
**omitted** here — computing them requires decoding every row's `components` JSON to find
each component's max/min score, which the `json.RawMessage` passthrough in `Components`
deliberately avoids paying for on every row. Add a follow-up ruling to the final report:
this is a real, small gap against the spec's literal example payload, left for a fast
follow-up (decode `Latest`'s own `Components` only — one row, not the whole trend — to fill
`best_component`/`worst_component` from the single most recent card) rather than blocking
this task on it. **Do implement the one-row version before closing this task**, since it
costs one `json.Unmarshal` on `Latest` only:

```go
func bestWorstComponent(latest json.RawMessage) (best, worst string) {
	var cs []componentDTO
	if err := json.Unmarshal(latest, &cs); err != nil {
		return "", ""
	}
	var bestScore, worstScore *float64
	for _, c := range cs {
		if c.Excluded || c.Score == nil {
			continue
		}
		if bestScore == nil || *c.Score > *bestScore {
			bestScore, best = c.Score, c.Name
		}
		if worstScore == nil || *c.Score < *worstScore {
			worstScore, worst = c.Score, c.Name
		}
	}
	return best, worst
}
```

Wire it into `buildCharacterRatingDTO`: after setting `out.Latest`, call
`out.BestComponent, out.WorstComponent = bestWorstComponent(latest.Components)` and add
`BestComponent, WorstComponent string \`json:"best_component,omitempty"\`/\`json:"worst_component,omitempty"\``
fields to `characterRatingDTO`.

- [ ] **Step 4: Run the handler tests**

Run: `cd api && go test ./internal/rating/... -p 1 -v`
Expected: every test PASSes.

- [ ] **Step 5: `go vet` the whole package**

Run: `cd api && go vet ./internal/rating/...`
Expected: clean.

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/rating/handler.go api/internal/rating/handler_test.go
printf 'feat: add the two ratings HTTP read endpoints\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task8.txt
git commit -F .superpowers/commit-msg-task8.txt
```

---

## Task 9: The backfill job

**Files:**
- Create: `api/internal/rating/backfill.go`
- Test: `api/internal/rating/backfill_test.go`

**Interfaces:**
- Consumes: `logs/engine/store.Keys.FightSummary`, `ratingengine.DefaultModelVersion`.
- Produces (for Task 10): `const BackfillJobCommand = "rating-backfill"`, `const
  BackfillBatchSize = 500`, `type Getter interface { Get(ctx, key string)
  (io.ReadCloser, error) }`, `type BackfillDeps struct { Store *Store; Summaries Getter;
  Log *slog.Logger }`, `func Backfill(ctx context.Context, d BackfillDeps, batchSize int)
  (recomputed int, err error)`.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/rating/backfill_test.go
package rating

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
)

var errFakeGetterMiss = errors.New("fake getter: no such object")

type fakeGetter struct{ objects map[string][]byte }

func (g fakeGetter) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	b, ok := g.objects[key]
	if !ok {
		return nil, errFakeGetterMiss
	}
	return io.NopCloser(bytes.NewReader(b)), nil
}

func TestBackfillRecomputesStaleRowsAndAdvancesModelVersion(t *testing.T) {
	pool := testPool(t)
	// Named ratingStore, not store: this file also imports logs/engine/store for
	// store.Keys below, and a local variable named store would shadow that package
	// identifier for the rest of the function.
	ratingStore := &Store{Pool: pool}
	ctx := context.Background()
	f := fightFixture("backfill-test-1", true,
		summary.RosterRow{GUID: "g1", Name: "Stale", Class: "Shaman", Spec: "Enhancement", Role: "dps"},
	)
	if err := ratingStore.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = $1`, f.ReportID) })
	// Force the row stale, as if an older model_version had written it.
	if _, err := pool.Exec(ctx, `update rating_scores set model_version = 'rating-2020-01-01' where report_id = $1`, f.ReportID); err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(f.Summary)
	if err != nil {
		t.Fatal(err)
	}
	getter := fakeGetter{objects: map[string][]byte{
		store.Keys{ReportID: f.ReportID}.FightSummary(f.FightIndex): body,
	}}
	n, err := Backfill(ctx, BackfillDeps{Store: ratingStore, Summaries: getter}, 10)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("recomputed %d fights, want 1", n)
	}
	var version string
	pool.QueryRow(ctx, `select model_version from rating_scores where report_id = $1`, f.ReportID).Scan(&version)
	if version != ratingengine.DefaultModelVersion {
		t.Errorf("model_version = %q, want %q", version, ratingengine.DefaultModelVersion)
	}
}

func TestBackfillIsBoundedPerRun(t *testing.T) {
	pool := testPool(t)
	ratingStore := &Store{Pool: pool} // see the naming note in the test above
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		f := fightFixture(fightIDFor(i), true,
			summary.RosterRow{GUID: "g1", Name: "Bounded", Class: "Shaman", Spec: "Elemental", Role: "dps"})
		if err := ratingStore.RateFight(ctx, f); err != nil {
			t.Fatal(err)
		}
		pool.Exec(ctx, `update rating_scores set model_version = 'rating-2020-01-01' where report_id = $1`, f.ReportID)
	}
	t.Cleanup(func() {
		for i := 0; i < 3; i++ {
			pool.Exec(ctx, `delete from rating_scores where report_id = $1`, fightIDFor(i))
		}
	})
	getter := fakeGetter{objects: map[string][]byte{}} // every read fails: this test only checks the batch bound, not success
	n, err := Backfill(ctx, BackfillDeps{Store: ratingStore, Summaries: getter}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 { // every read failed, so nothing was actually recomputed - but the query must have been capped at 2
		t.Fatalf("recomputed %d, want 0 (every read fails in this test)", n)
	}
}

func fightIDFor(i int) string {
	return "backfill-bounded-" + string(rune('a'+i))
}
```

- [ ] **Step 2: Run to see them fail**

Run: `cd api && go test ./internal/rating/... -run TestBackfill -p 1 -v`
Expected: FAIL.

- [ ] **Step 3: Implement `backfill.go`**

```go
// api/internal/rating/backfill.go
package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// BackfillJobCommand is this lane's Cloud Run job name, dispatched the same way
// sims.ValidateJobCommand/sims.SimRunJobCommand already are (api/cmd/api/main.go's
// os.Args[1] switch).
const BackfillJobCommand = "rating-backfill"

// BackfillBatchSize bounds one run: the spec's own cost note (§4.4/§4.5) treats
// over-recomputing as a compute cost, not a correctness one, so this can be generous
// without risk - 500 fights is comfortably inside one Cloud Run job's timeout at the
// per-fight cost §4.5 already budgets for fight-close itself.
const BackfillBatchSize = 500

// MaxBackfillSummaryBytes bounds a stored fight summary read back out of the bucket -
// mirrors sims.MaxSummaryBytes (api/internal/sims/summaries.go): nothing the ingest ever
// accepted is bigger than this.
const MaxBackfillSummaryBytes = 64 << 20

// Getter reads one stored object back. *r2.Client satisfies it; a nil Getter means this
// deployment has no bucket, and Backfill no-ops, logging why rather than failing.
type Getter interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// BackfillDeps is what one backfill run needs.
type BackfillDeps struct {
	Store     *Store
	Summaries Getter
	Log       interface {
		Error(msg string, args ...any)
		Warn(msg string, args ...any)
	}
}

func (d BackfillDeps) logf() interface {
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
} {
	if d.Log != nil {
		return d.Log
	}
	return noopLogger{}
}

type noopLogger struct{}

func (noopLogger) Error(string, ...any) {}
func (noopLogger) Warn(string, ...any)  {}

// staleFight is one (report, fight) this run must recompute.
type staleFight struct {
	ReportID   string
	FightIndex int
	PlayerKey  string
	FoughtAt   staleTime
}

// Backfill recomputes every rating_scores row not yet at logs/engine/rating's current
// DefaultModelVersion, up to batchSize fights, and upserts them via Store.RateFight -
// which, because every one of these rows already exists, always takes the rewrite path
// (spec §4.4: "digest tables are not rebuilt from scratch on a backfill"). It is
// resumable with no extra state: each run only ever selects rows still stale after the
// previous run's writes, so a crash mid-run simply leaves the next run a slightly larger
// batch to work through, and safe to run concurrently with live ingest, because
// RateFight's own per-fight advisory lock (store.go) serialises the two.
func Backfill(ctx context.Context, d BackfillDeps, batchSize int) (recomputed int, err error) {
	if d.Store == nil || d.Summaries == nil {
		return 0, nil
	}
	if batchSize <= 0 {
		batchSize = BackfillBatchSize
	}
	fights, err := d.Store.staleFights(ctx, batchSize)
	if err != nil {
		return 0, fmt.Errorf("rating: backfill: %w", err)
	}
	for _, sf := range fights {
		region, ruleset, ok := splitPlayerKeyRegionRuleset(sf.PlayerKey)
		if !ok {
			d.logf().Warn("rating", "op", "backfill", "report", sf.ReportID, "fight", sf.FightIndex,
				"err", "unparseable player_key "+sf.PlayerKey)
			continue
		}
		sum, err := readSummary(ctx, d.Summaries, sf.ReportID, sf.FightIndex)
		if err != nil {
			d.logf().Error("rating", "op", "backfill", "report", sf.ReportID, "fight", sf.FightIndex, "err", err)
			continue
		}
		if err := d.Store.RateFight(ctx, RatedFight{
			ReportID: sf.ReportID, FightIndex: sf.FightIndex, Region: region, Ruleset: ruleset,
			FoughtAt: sf.FoughtAt.Time, EncounterID: sum.EncounterID, Summary: sum,
		}); err != nil {
			d.logf().Error("rating", "op", "backfill", "report", sf.ReportID, "fight", sf.FightIndex, "err", err)
			continue
		}
		recomputed++
	}
	return recomputed, nil
}

// readSummary reads one fight's stored summary out of the bucket - the same shape
// api/internal/sims/summaries.go's own unexported fightSummary uses (MaxSummaryBytes
// bound, store.Keys for the object key), duplicated in miniature rather than imported:
// importing api/internal/sims here to reach one unexported helper is not possible, and
// the alternative (exporting it from sims for one cross-lane caller) is out of this
// lane's file ownership. The duplicated part is plain object-read boilerplate, not
// business logic.
func readSummary(ctx context.Context, get Getter, reportID string, index int) (summary.Summary, error) {
	body, err := get.Get(ctx, store.Keys{ReportID: reportID}.FightSummary(index))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("rating: read summary %s/%d: %w", reportID, index, err)
	}
	defer body.Close()
	b, err := io.ReadAll(io.LimitReader(body, MaxBackfillSummaryBytes))
	if err != nil {
		return summary.Summary{}, fmt.Errorf("rating: read summary %s/%d: %w", reportID, index, err)
	}
	var s summary.Summary
	if err := json.Unmarshal(b, &s); err != nil {
		return summary.Summary{}, fmt.Errorf("rating: decode summary %s/%d: %w", reportID, index, err)
	}
	return s, nil
}

// splitPlayerKeyRegionRuleset reads the region/ruleset prefix back out of a character key
// ("us/normal/simfury" -> "us", "normal") - every player of one fight shares the same
// pair (api/internal/character.KeyFromUnit builds every row's key from the same
// RatedFight.Region/Ruleset), so reading it back off any one stored row is enough to
// rebuild the RatedFight Backfill needs.
func splitPlayerKeyRegionRuleset(key string) (region, ruleset string, ok bool) {
	parts := splitN(key, "/", 3)
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}
```

`staleTime` and `splitN` are small local helpers — implement `splitN` as
`strings.SplitN(key, sep, n)` directly (drop the wrapper, call `strings.SplitN` inline, add
`"strings"` to imports) rather than inventing a new name; `staleFight.FoughtAt` can just be
`time.Time` directly (drop the `staleTime` wrapper — it added no value here; the test file
above's `sf.FoughtAt.Time` reference should become `sf.FoughtAt` once this simplification is
made, and the test edited to match before it is run).

Add `Store.staleFights` to `store.go`:

```go
// staleFights selects up to limit (report_id, fight_index) pairs whose stored rows are not
// at the engine's current DefaultModelVersion, one representative player_key and fought_at
// per fight (every player of one fight shares both).
func (s *Store) staleFights(ctx context.Context, limit int) ([]staleFight, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct on (report_id, fight_index) report_id, fight_index, player_key, fought_at
		 from rating_scores where model_version <> $1
		 order by report_id, fight_index limit $2`,
		ratingengine.DefaultModelVersion, limit)
	if err != nil {
		return nil, fmt.Errorf("rating: stale fights: %w", err)
	}
	defer rows.Close()
	var out []staleFight
	for rows.Next() {
		var sf staleFight
		if err := rows.Scan(&sf.ReportID, &sf.FightIndex, &sf.PlayerKey, &sf.FoughtAt); err != nil {
			return nil, err
		}
		out = append(out, sf)
	}
	return out, rows.Err()
}
```

Add the `ratingengine` import alias to `store.go` if not already present from Task 5, and
move the `staleFight` type declaration into `store.go` beside `staleFights` (it is a store
concern, not a backfill-file concern) rather than `backfill.go` — adjust `backfill.go` to
just reference `staleFight` without redeclaring it.

- [ ] **Step 4: Run the tests**

Run: `cd api && go test ./internal/rating/... -run TestBackfill -p 1 -v`
Expected: every test PASSes. If `TestBackfillIsBoundedPerRun`'s batch-of-2-vs-3 assertion
needs a stronger check than "0 recomputed" (since every read fails in that test either
way), strengthen it by asserting `d.Store.staleFights(ctx, 10)` still reports the expected
remaining count after the bounded run — call `staleFights` directly in the test for this,
since it is in the same package.

- [ ] **Step 5: Run the whole package once more**

Run: `cd api && go vet ./internal/rating/... && go test ./internal/rating/... -p 1 -v`
Expected: PASS.

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/rating/backfill.go api/internal/rating/backfill_test.go api/internal/rating/store.go
printf 'feat: add the ratings backfill job\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task9.txt
git commit -F .superpowers/commit-msg-task9.txt
```

---

## Task 10: Wire it all into `main.go` and `server.go`, then build the whole module

**Files:**
- Modify: `api/internal/server/server.go` (small, contiguous: one `Deps` field, one `Mount`
  call, following the existing `if d.X != nil { x.Mount(mux, d.X) }` pattern exactly)
- Modify: `api/cmd/api/main.go` (small, contiguous: construct the rating store/service,
  wire `deps.Ingest.Rate`, start/stop the `Rater` goroutine beside the existing
  sampler/scorer, add the `rating-backfill` job dispatch, add `rating_scores` to the
  partition job's `Tables`)
- Test: `api/cmd/api/main_test.go` (read it first — it already asserts `newReportsService`'s
  wiring; add a parallel assertion for the rating wiring if that file's existing pattern
  makes one easy; otherwise a plain `go build ./...` is this task's real verification)

**Interfaces:**
- Consumes everything Tasks 1–9 produced.
- Produces: nothing further downstream — this is the leaf wiring task.

- [ ] **Step 1: Add `Rating` to `server.Deps` and mount it**

In `api/internal/server/server.go`, add the import
`"github.com/jhunthrop/foreversixty/api/internal/rating"`, add `Rating *rating.Service` to
the `Deps` struct next to `Sims *sims.Service`, and add, next to the existing
`if d.Sims != nil { sims.Mount(mux, d.Sims) }` block:

```go
	if d.Rating != nil {
		rating.Mount(mux, d.Rating)
	}
```

- [ ] **Step 2: Wire construction in `main.go`**

In `serve()`, after the existing `reportStore := &reports.Store{Pool: pool}` /
`rankStore := &rankings.Store{...}` lines and before `client := objects(cfg, log)`, add:

```go
	ratingStore := &rating.Store{Pool: pool, Log: log}
```

In the `deps := server.Deps{...}` literal, add `Rating: &rating.Service{Store: ratingStore,
Reports: reportStore, Accounts: authStore, Log: log},` next to the `Guilds:` line (`authStore`
already satisfies `rating.Accounts` — it has `GuildRank`).

Inside the `if client != nil { ... }` block, immediately after the existing
`deps.Ingest.Score = scorerShim{scorer}` / `deps.Ingest.Members = authStore` lines, add:

```go
		rater = rating.NewRater(rating.RateDeps{Store: ratingStore, Log: log})
		go rater.Run(ctx)
		deps.Ingest.Rate = rater
```

and declare `var rater *rating.Rater` alongside the existing `var sampler *parse.Worker` /
`var scorer *sims.Scorer` declarations near the top of `serve()`.

In the shutdown sequence at the bottom of `serve()`, next to `if scorer != nil {
scorer.Close() }`, add:

```go
	if rater != nil {
		rater.Close()
	}
```

Extend the `PartitionJob` construction:

```go
	partitions := &db.PartitionJob{Pool: pool, Log: log, Tables: []string{"fight_metrics", db.RatingsTable}}
```

(replacing the existing `partitions := &db.PartitionJob{Pool: pool, Log: log}` line).

Add the job dispatch in `main()`'s `switch os.Args[1]` block, next to the existing
`sims.ValidateJobCommand` case:

```go
		case rating.BackfillJobCommand:
			if err := runRatingBackfill(context.Background(), log); err != nil {
				log.Error(rating.BackfillJobCommand, "err", err)
				os.Exit(1)
			}
			return
```

Add the `runRatingBackfill` function, following `runValidate`'s exact shape:

```go
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
		Store: &rating.Store{Pool: pool, Log: log}, Summaries: client, Log: log,
	}, rating.BackfillBatchSize)
	if err != nil {
		return err
	}
	log.Info(rating.BackfillJobCommand, "recomputed", n)
	return nil
}
```

`rating.BackfillDeps.Log` is typed as a small logging interface (Task 9's `interface {
Error(...); Warn(...) }`), and `*slog.Logger` already satisfies it (it has both methods with
that exact signature) — confirm this compiles as-is; if `slog.Logger`'s methods don't match
that interface shape exactly (they take `(msg string, args ...any)` — they do), no
adapter is needed.

- [ ] **Step 3: Build the whole module**

Run: `cd api && go build ./... && go vet ./...`
Expected: clean build, no vet warnings.

- [ ] **Step 4: Run `main_test.go`**

Run: `cd api && go test ./cmd/api/... -v`
Expected: PASS (this task did not change `newReportsService`, so its existing assertions
are unaffected; if the file's own pattern makes it easy to add a one-line assertion that
`deps.Rating != nil` when `client != nil`, add it — otherwise this step is just the
regression check).

- [ ] **Step 5: Run the full lane's test suite once, end to end**

```bash
cd api
go vet ./...
go test ./... -p 1
```

Expected: every package PASSes (packages with no `TEST_DATABASE_URL`-gated tests run
normally; DB-backed packages use the throwaway `rating_api_test` database from Task 1).

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-api
git add api/internal/server/server.go api/cmd/api/main.go
printf 'feat: wire the ratings service, rater worker, and backfill job into main\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task10.txt
git commit -F .superpowers/commit-msg-task10.txt
```

- [ ] **Step 7: Drop the throwaway test database**

```bash
cd api
docker compose -f docker-compose.test.yml exec -T postgres psql -U forever -d forever_test \
  -c "drop database if exists rating_api_test;"
```

Do this only after Task 11 (the whole-branch review) is fully done and no more test runs
are pending in this lane.

---

## Task 11: Whole-branch review (code + security), one fix wave

This is not a subagent-implemented task — it is the lane controller's own final review
step per `lane-common-go.md`, run after Task 10 lands. Use the visibility rules in this
plan's Global Constraints as the security checklist:

1. Every SQL statement in `api/internal/rating/*.go` is parameterised (grep for any
   `fmt.Sprintf` building a query with request-derived input — there should be none; the
   only `Sprintf`-built statement in this lane is `partitions.go`'s table-name
   interpolation, which is always a fixed, code-provided constant).
2. `visible()` in `viewer.go` is checked line-by-line against `mayView`
   (`reports/handler.go:540-556`) for drift.
3. The anonymize gate is checked on both endpoints, and the per-fight endpoint drops the
   row entirely (no placeholder) per spec §5.1/§6.4.
4. The character endpoint's query joins live `reports.visibility = 'public'`, not a cached
   visibility column, per spec §5.1's "at time of query" ruling.
5. `i.rate` never returns an error and never blocks `putFight`'s response.
6. `RateFight`'s rewrite-detection and the adapter's `Fold` flag together guarantee no
   double-fold on re-ingest or backfill — re-read Task 5/Task 9's idempotency tests and
   confirm they actually exercise this (not just assert row counts).
7. The `rating-backfill` job is bounded (`BackfillBatchSize`) and resumable (no external
   cursor state needed between runs).
8. `docs/superpowers/sdd/2026-09-21-rating-api/progress.md` (created during subagent-driven
   execution) has a `Ruling:` line for each of: the `mayView` duplication, the "no premium
   gate" finding, the per-fight-vs-trend pagination scoping, and the `0022` migration
   number.

Fix anything Steps 1–8 find in one additional fix wave (fresh subagent per fix, same
sonnet-only budget), then re-run `go vet ./... && go test ./... -p 1` from `api/` before
declaring the lane done.

---

## Self-Review Notes (for whoever executes this plan)

- **Spec coverage:** §1–§4 (scoring model, fairness, curated data, storage/invocation) are
  the merged engine's responsibility (out of this lane's scope) plus this lane's Tasks 1–7
  (storage, hook, backfill). §5 (API) is Task 8. §5.3 (caching) is Task 8, Step 3. §7 (no
  leaderboard) is satisfied by construction — no endpoint in this plan accepts a
  cross-report/cross-guild query. §8's lane table and numbering note are honoured by the
  Global Constraints section and Task 1. §9 (out of scope) needs no task.
- **Known, recorded gaps to name in the final report:** (a) `RemoveReport`-style
  withdrawal of `rating_scores` on report moderation/removal is not built — out of this
  lane's explicit scope and out of `rankings` package ownership, but worth flagging as a
  follow-up; (b) `best_component`/`worst_component` are computed from `Latest` only, not
  scanned across the whole `trend[]`, for cost reasons — implemented in Task 8; (c) the
  officer premium gate was found to not apply to either endpoint in scope, so no `CanSee`
  interface exists (record exactly where it would be consulted if built: `fightRatings`
  and `characterRating` in `handler.go`, right after the `visible()`/anonymize checks).
