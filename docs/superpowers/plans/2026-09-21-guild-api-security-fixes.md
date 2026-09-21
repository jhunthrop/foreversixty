# Guild API Security Fix Round Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development
> to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix seven findings (2 Critical, 2 High, 2 Medium, 1 Low) from an independent
security review of the `guild-api` branch, per the coordinator's explicit rulings (R1-R7),
before the branch may merge.

**Architecture:** Every fix lands in the same packages the original branch already built
(`api/internal/guilds`, `api/internal/addon`), plus one migration amendment (0018 is
unreleased and may be edited in place, not renumbered). No new packages except one new file,
`api/internal/guilds/contest.go`, for the claim contest/resolve flow.

**Tech Stack:** Go 1.25, pgx/v5, Postgres 16, `golang.org/x/text/unicode/norm` (already an
indirect dependency in `api/go.mod`, promoted to direct by this work).

**Spec:** `docs/superpowers/specs/2026-09-21-guild-membership-design.md` — see the two
"Amendment, 2026-09-21 (security review response)" blocks in §2.4 and §3.3, already written
and committed ahead of this plan; this plan implements exactly what they describe.

**Prior plan this extends:** `docs/superpowers/plans/2026-09-21-guild-api.md` (already fully
implemented, reviewed, and merged into this branch — see its own ledger, now deleted; the
branch head this plan starts from is `bbc60e3`).

## Global Constraints

- Same rules as the prior plan: `go vet ./... && go test -p 1 ./...` (from `api/`) must
  pass before every commit; `-p 1` avoids spurious deadlocks from parallel packages sharing
  the test Postgres container; `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable`
  exported before any DB-touching `go test`; Go 1.25.11 already active, no `GOTOOLCHAIN=`
  needed.
- File ownership (unchanged from the prior plan): this lane owns `api/internal/db/migrations/0018_*`,
  `api/internal/addon/*.go`, `api/internal/guilds/*.go`, `api/internal/server/server.go`,
  `api/cmd/api/main.go`. **Never edit** `api/internal/reports/handler.go`,
  `api/internal/reports/store.go`, anything under `api/internal/sims/` or
  `api/internal/rankings/`, or anything outside `api/`. `api/internal/auth/store.go` may be
  read but this plan needs no changes to it — `auth.User.BnetSub` is already exported.
- **Migration 0018 is amended in place, not renumbered** — it has never shipped (confirmed:
  production has zero rows in `guilds`/`guild_members`/`guild_characters`), so editing
  `0018_guild_membership.up.sql`/`.down.sql` directly is safe and is what the coordinator
  explicitly authorised. Do not create `0019`.
- Commits: conventional subjects (`feat:`, `fix:`, `test:`, `docs:`), written to a file
  under this worktree's `.superpowers/` with `printf`, committed with `git commit -F <file>`
  as its own Bash command, ending with exactly:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`
  (This is a lane-wide, coordinator-set convention overriding any generic per-session
  attribution default naming a different model — see this branch's own commit history for
  precedent; every commit on this branch uses this exact text.)
- Parameterised SQL only. Functions under 50 lines, files under 800, table-driven tests,
  errors wrapped with `fmt.Errorf("<package>: <what>: %w", err)`.
- Failing test first, per finding, for every behavioral change in this plan — this is a
  security fix round; there is no "trust the description" shortcut.
- **New response/error shapes this plan introduces** (the web lane builds against these —
  list them verbatim in your final report): `ClaimStateView{state, since?}` exposed as
  `claim` on both `SettingsView` and `HomeView`; `POST /v1/guilds/{id}/claim/contest` →
  `{status: "contested"}`; `POST /v1/guilds/{id}/claim/resolve` body
  `{outcome: "uphold"|"release"|"transfer"}` → `{status: "resolved", outcome}`; new error
  code `409 claim_contested` on `approve`, `remove` (unless self/moderator),
  `PATCH .../settings`, and `POST .../invite/rotate`; `SettingsView.ClaimedBy`/`ClaimPending`
  are now actually populated (previously always `null`).

---

## Task 1: Migration 0018 amendment — verified_by, contest columns, claim-attempt table, case-insensitive guild identity

**Files:**
- Modify: `api/internal/db/migrations/0018_guild_membership.up.sql`
- Modify: `api/internal/db/migrations/0018_guild_membership.down.sql`
- Modify: `api/internal/db/db_test.go`

**Interfaces:**
- Produces: `guild_characters.verified_by` (text, check `claim|officer|invite|logs`,
  nullable); `guilds.claim_contested_at` (timestamptz), `guilds.claim_contested_by` (bigint
  FK users); new table `guild_claim_attempts (id, user_id, attempted_at)`; new unique index
  `guilds_region_ruleset_lower_name_idx` on `(region, ruleset, lower(name))`.

- [ ] **Step 1: Read the current migration files**

Read `api/internal/db/migrations/0018_guild_membership.up.sql` and `.down.sql` in full
before editing — this task appends to the up file and prepends to the down file; your
old_string/new_string edits must match the current byte-for-byte content, not a guess.

- [ ] **Step 2: Append to the up migration**

At the end of `0018_guild_membership.up.sql`, append:

```sql

-- Security hardening (2026-09-21 review response — see the spec's dated
-- amendment blocks in §2.4 and §3.3).

-- Which of the four corroboration paths actually set verified_at, so a
-- claim release/transfer can un-verify precisely the rows the claim
-- itself vouched for and nothing a character separately earned through
-- officer approval, an invite, or log corroboration.
alter table guild_characters add column if not exists verified_by text
  check (verified_by in ('claim', 'officer', 'invite', 'logs'));

-- A claim may be contested while pending or already claimed; a
-- moderator resolves it (uphold, release, or transfer).
alter table guilds add column if not exists claim_contested_at timestamptz;
alter table guilds add column if not exists claim_contested_by bigint references users (id) on delete set null;

-- Rate-limits an account to one claim (successful or pending) per
-- rolling 30 days, across every guild - a small, guilds-owned table
-- rather than a column on auth's own users table.
create table if not exists guild_claim_attempts (
  id           bigserial primary key,
  user_id      bigint not null references users (id) on delete cascade,
  attempted_at timestamptz not null default now()
);
create index if not exists guild_claim_attempts_user_idx on guild_claim_attempts (user_id, attempted_at desc);

-- Guild identity is case-insensitive: one guild per (region, ruleset,
-- lower(name)). The pre-existing exact-text unique(region, ruleset,
-- name) constraint (migration 0005) stays - the new index is strictly
-- stricter and subsumes it, so both coexist harmlessly.
create unique index if not exists guilds_region_ruleset_lower_name_idx
  on guilds (region, ruleset, lower(name));
```

- [ ] **Step 3: Prepend to the down migration**

At the very start of `0018_guild_membership.down.sql` (before its existing first line),
insert:

```sql
-- Security hardening additions (2026-09-21), reversed first, in exact
-- reverse order of the up migration's appended block.
drop index if exists guilds_region_ruleset_lower_name_idx;
drop table if exists guild_claim_attempts;
alter table guilds drop column if exists claim_contested_by;
alter table guilds drop column if exists claim_contested_at;
alter table guild_characters drop column if exists verified_by;

```

(A trailing blank line before the file's existing first `drop` statement is fine and matches
this file's existing style of one statement per line.)

- [ ] **Step 4: Write the failing tests**

Append to `api/internal/db/db_test.go`:

```go
func TestMigration0018SecurityHardeningColumnsRoundTrip(t *testing.T) {
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

	for _, check := range []struct{ table, column string }{
		{"guild_characters", "verified_by"},
		{"guilds", "claim_contested_at"},
		{"guilds", "claim_contested_by"},
	} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = $1 and column_name = $2`,
			check.table, check.column).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 1 {
			t.Errorf("%s.%s should exist at the latest migration", check.table, check.column)
		}
	}
	if n := tableCount(t, pool, "guild_claim_attempts"); n != 1 {
		t.Fatal("guild_claim_attempts should exist at the latest migration")
	}
	if n := indexCount(t, pool, "guilds_region_ruleset_lower_name_idx"); n != 1 {
		t.Fatal("guilds_region_ruleset_lower_name_idx should exist at the latest migration")
	}

	migrateTo(t, url, 17)
	for _, check := range []struct{ table, column string }{
		{"guild_characters", "verified_by"},
		{"guilds", "claim_contested_at"},
		{"guilds", "claim_contested_by"},
	} {
		var n int
		if err := pool.QueryRow(context.Background(),
			`select count(*) from information_schema.columns where table_name = $1 and column_name = $2`,
			check.table, check.column).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("%s.%s survived the down migration", check.table, check.column)
		}
	}
	if n := tableCount(t, pool, "guild_claim_attempts"); n != 0 {
		t.Error("guild_claim_attempts survived the down migration")
	}
	if n := indexCount(t, pool, "guilds_region_ruleset_lower_name_idx"); n != 0 {
		t.Error("guilds_region_ruleset_lower_name_idx survived the down migration")
	}
	// 0018's own base table must survive its own down migration's new
	// prefix - only the newly-appended block should be reversed here.
	if n := tableCount(t, pool, "guild_characters"); n != 0 {
		t.Error("guild_characters itself should still be gone at version 17 (this assertion documents the base table's own down migration still runs)")
	}

	if err := Migrate(url); err != nil {
		t.Fatalf("migrating up again: %v", err)
	}
	if n := tableCount(t, pool, "guild_claim_attempts"); n != 1 {
		t.Error("guild_claim_attempts did not come back")
	}
}

func TestGuildsAreCaseInsensitiveByRegionAndRuleset(t *testing.T) {
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
	if _, err := pool.Exec(ctx, `truncate guilds cascade`); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Iron Vanguard')`); err != nil {
		t.Fatal(err)
	}
	_, err = pool.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'IRON VANGUARD')`)
	if err == nil {
		t.Fatal("a case-variant name in the same region/ruleset should violate the case-insensitive unique index")
	}
	// A different ruleset, or a genuinely different name, is unaffected.
	if _, err := pool.Exec(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Iron Vanguard')`); err != nil {
		t.Fatalf("a different ruleset should not collide: %v", err)
	}
}
```

- [ ] **Step 5: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/db/... -run 'TestMigration0018SecurityHardeningColumnsRoundTrip|TestGuildsAreCaseInsensitiveByRegionAndRuleset' -v
```

Expected: FAIL — the columns/table/index do not exist yet.

- [ ] **Step 6: Apply the migration edits (Steps 2-3 above) and re-run**

```bash
cd api && go test -p 1 ./internal/db/... -run 'TestMigration0018SecurityHardeningColumnsRoundTrip|TestGuildsAreCaseInsensitiveByRegionAndRuleset' -v
```

Expected: PASS, both tests.

- [ ] **Step 7: Run the full db package to confirm no regression**

```bash
cd api && go test -p 1 ./internal/db/... -v 2>&1 | tail -60
```

Expected: PASS, every test including `TestMigration0018DownReversesUp` (the original
round-trip test from the prior plan) and `TestGuildCharactersCharacterKeyIsUniqueAcrossGuilds`.

- [ ] **Step 8: Commit**

```bash
cd api && git add internal/db/migrations/0018_guild_membership.up.sql \
  internal/db/migrations/0018_guild_membership.down.sql internal/db/db_test.go
printf 'feat(db): amend migration 0018 for the security review response\n\nAdds guild_characters.verified_by, guilds.claim_contested_at/by, the\nguild_claim_attempts rate-limit table, and a case-insensitive unique\nindex on guild identity. 0018 is unreleased, so amended in place\nrather than a new migration.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-1.txt
git commit -F ../.superpowers/commit-msg-sec-1.txt
```

---

## Task 2: `guilds/store.go` — advisory lock, contest fields, account-wide verification helper

**Files:**
- Modify: `api/internal/guilds/store.go`
- Modify: `api/internal/guilds/store_test.go`

**Interfaces:**
- Consumes: `guild_characters.verified_by`, `guilds.claim_contested_at`/`claim_contested_by`
  (Task 1).
- Produces: `Guild` struct gains `ClaimContestedAt *time.Time`, `ClaimContestedBy *int64`;
  `func (s *Store) contested(ctx, guildID int64) (bool, error)`;
  `func setVerifiedForAccount(ctx, tx pgx.Tx, guildID, userID int64, source string) error`
  (package-level, same shape as `RecomputeMembership`/`ReleaseClaimIfLost`, for callers that
  hold their own transaction); `RecomputeMembership` now takes an advisory lock first — no
  signature change, existing callers are unaffected.

This task is foundational: Tasks 3, 5, 6, 7, 8 all depend on the `Guild` struct's two new
fields, `contested`, and `setVerifiedForAccount`.

- [ ] **Step 1: Write the failing tests**

Append to `api/internal/guilds/store_test.go`:

```go
func TestGetGuildReadsContestFields(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	contester := seedUser(t, pool, "contester@example.com")
	if _, err := pool.Exec(ctx,
		`update guilds set claim_contested_at = now(), claim_contested_by = $2 where id = $1`, gid, contester); err != nil {
		t.Fatal(err)
	}
	g, err := s.getGuild(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	if g.ClaimContestedAt == nil || g.ClaimContestedBy == nil || *g.ClaimContestedBy != contester {
		t.Fatalf("g.ClaimContestedAt/By = %v, %v, want set and %d", g.ClaimContestedAt, g.ClaimContestedBy, contester)
	}
}

func TestContestedReportsWhetherAGuildsClaimIsDisputed(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")

	yes, err := s.contested(ctx, gid)
	if err != nil || yes {
		t.Fatalf("contested = %v, %v, want false on a fresh guild", yes, err)
	}
	if _, err := pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	yes, err = s.contested(ctx, gid)
	if err != nil || !yes {
		t.Fatalf("contested = %v, %v, want true once claim_contested_at is set", yes, err)
	}
}

func TestSetVerifiedForAccountVerifiesEveryRowAndRecordsSource(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	uid := seedUser(t, pool, "multichar@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/main", "leader", false)
	seedCharacter(t, pool, gid, uid, "us/hardcore/alt", "member", false)

	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := setVerifiedForAccount(ctx, tx, gid, uid, "claim"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}

	rows, err := pool.Query(ctx,
		`select verified_at is not null, verified_by from guild_characters where guild_id = $1 and user_id = $2`, gid, uid)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var verified bool
		var by *string
		if err := rows.Scan(&verified, &by); err != nil {
			t.Fatal(err)
		}
		if !verified || by == nil || *by != "claim" {
			t.Fatalf("row: verified = %v, verified_by = %v, want true/claim", verified, by)
		}
		n++
	}
	if n != 2 {
		t.Fatalf("verified %d rows, want 2 (both of the account's characters)", n)
	}

	// A second call with a different source must not reclassify an
	// already-verified row - the FIRST path that verified it is what a
	// later release/transfer un-verifies.
	tx, err = pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := setVerifiedForAccount(ctx, tx, gid, uid, "officer"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	var by string
	if err := pool.QueryRow(ctx,
		`select verified_by from guild_characters where character_key = 'us/hardcore/main'`).Scan(&by); err != nil {
		t.Fatal(err)
	}
	if by != "claim" {
		t.Fatalf("verified_by = %q after a second call, want it to stay claim (first writer wins)", by)
	}
}

func TestRecomputeMembershipSerialisesConcurrentCallersForTheSameGuild(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	var wg sync.WaitGroup
	errs := make([]error, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			uid := seedUser(t, pool, fmt.Sprintf("lockrace-%d@example.com", i))
			seedCharacter(t, pool, gid, uid, fmt.Sprintf("us/hardcore/lockrace%d", i), "member", true)
			tx, err := pool.Begin(ctx)
			if err != nil {
				errs[i] = err
				return
			}
			if err := RecomputeMembership(ctx, tx, gid, &uid); err != nil {
				errs[i] = err
				return
			}
			errs[i] = tx.Commit(ctx)
		}(i)
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Fatalf("goroutine %d: %v", i, err)
		}
	}
	var n int
	if err := pool.QueryRow(ctx, `select count(*) from guild_members where guild_id = $1`, gid).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 20 {
		t.Fatalf("guild_members rows = %d, want 20 (one per concurrently-added account, none lost to the race)", n)
	}
}
```

Add `"fmt"` and `"sync"` to `store_test.go`'s import block if not already present (the file
already imports `context`, `testing`).

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/guilds/... -run 'TestGetGuildReadsContestFields|TestContestedReports|TestSetVerifiedForAccount|TestRecomputeMembershipSerialises' -v
```

Expected: FAIL to compile — `ClaimContestedAt`, `contested`, `setVerifiedForAccount` undefined.

- [ ] **Step 3: Update `store.go`**

Change the `Guild` struct:

```go
// Guild is one row of the guilds table, as this package's handlers need it.
type Guild struct {
	ID                    int64
	Region, Ruleset, Name string
	DefaultVisibility     string
	ClaimedBy             *int64
	ClaimPendingBy        *int64
	ClaimRequestedAt      *time.Time
	ClaimContestedAt      *time.Time
	ClaimContestedBy      *int64
	OfficerMaxRankIndex   int
	InviteTokenRotatedAt  *time.Time
}
```

Change `getGuild`:

```go
func (s *Store) getGuild(ctx context.Context, id int64) (Guild, error) {
	var g Guild
	err := s.Pool.QueryRow(ctx,
		`select id, region, ruleset, name, default_visibility, claimed_by, claim_pending_by,
		        claim_requested_at, claim_contested_at, claim_contested_by,
		        officer_max_rank_index, invite_token_rotated_at
		 from guilds where id = $1`, id).
		Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name, &g.DefaultVisibility, &g.ClaimedBy,
			&g.ClaimPendingBy, &g.ClaimRequestedAt, &g.ClaimContestedAt, &g.ClaimContestedBy,
			&g.OfficerMaxRankIndex, &g.InviteTokenRotatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Guild{}, ErrNotFound
	}
	if err != nil {
		return Guild{}, fmt.Errorf("guilds: read %d: %w", id, err)
	}
	return g, nil
}
```

Add, near `IsMember`:

```go
// contested reports whether guildID's claim is currently disputed - the
// freeze gate every officer-power route in this package checks before
// acting, so a disputed claim cannot be used to entrench itself while a
// moderator investigates.
func (s *Store) contested(ctx context.Context, guildID int64) (bool, error) {
	var yes bool
	err := s.Pool.QueryRow(ctx,
		`select claim_contested_at is not null from guilds where id = $1`, guildID).Scan(&yes)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, ErrNotFound
	}
	if err != nil {
		return false, fmt.Errorf("guilds: contested: %w", err)
	}
	return yes, nil
}
```

Add the advisory lock as the first statement in `RecomputeMembership`:

```go
func RecomputeMembership(ctx context.Context, tx pgx.Tx, guildID int64, userID *int64) error {
	// Serialises every recompute for this guild against every other -
	// PutExports, ApproveCharacter, RemoveCharacter, Claim,
	// UpdateSettings and the sweep jobs can all call this concurrently
	// from independent transactions; without a lock, two concurrent
	// SELECT-then-UPSERT passes can each compute from a stale
	// pre-lock snapshot, and the later-committing transaction's stale
	// values win, transiently overwriting a just-set verified_at/rank
	// until the next recompute. Transaction-scoped (released on commit
	// or rollback) and keyed on the whole guild, not one account, so a
	// whole-guild recompute (userID nil) and a single-account one still
	// serialise against each other rather than leaving a gap.
	if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock($1)`, guildID); err != nil {
		return fmt.Errorf("guilds: recompute membership: lock guild %d: %w", guildID, err)
	}
	if _, err := tx.Exec(ctx, `
		insert into guild_members (guild_id, user_id, rank, verified_at, refreshed_at)
		select $1, gc.user_id,
		       case max(case when gc.verified_at is not null then
		                   case gc.rank when 'leader' then 3 when 'officer' then 2 else 1 end
		                 end)
		         when 3 then 'leader' when 2 then 'officer' else 'member' end,
		       min(gc.verified_at) filter (where gc.verified_at is not null),
		       now()
		from guild_characters gc
		where gc.guild_id = $1 and ($2::bigint is null or gc.user_id = $2)
		group by gc.user_id
		on conflict (guild_id, user_id) do update set
		  rank = excluded.rank, verified_at = excluded.verified_at, refreshed_at = now()
	`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: recompute membership for guild %d: %w", guildID, err)
	}
	if _, err := tx.Exec(ctx, `
		delete from guild_members m
		where m.guild_id = $1 and ($2::bigint is null or m.user_id = $2)
		  and not exists (
		    select 1 from guild_characters gc where gc.guild_id = m.guild_id and gc.user_id = m.user_id
		  )
	`, guildID, userID); err != nil {
		return fmt.Errorf("guilds: prune membership for guild %d: %w", guildID, err)
	}
	return nil
}
```

Add, near `ReleaseClaimIfLost`:

```go
// setVerifiedForAccount marks every one of userID's guild_characters
// rows in guildID verified via source, without overwriting an
// already-verified row's original source - the claim flow trusts every
// alt the claiming/confirming account holds in the guild (§2.4), not
// just the one claiming character, so this is account-wide rather than
// the single-character scope ApproveCharacter/AcceptInvite/VerifyByLogs
// each use for their own UPDATE statements.
func setVerifiedForAccount(ctx context.Context, tx pgx.Tx, guildID, userID int64, source string) error {
	if _, err := tx.Exec(ctx,
		`update guild_characters set verified_at = coalesce(verified_at, now()),
		   verified_by = coalesce(verified_by, $3)
		 where guild_id = $1 and user_id = $2`, guildID, userID, source); err != nil {
		return fmt.Errorf("guilds: set verified (%s) for guild %d: %w", source, guildID, err)
	}
	return nil
}
```

- [ ] **Step 4: Run to verify it passes**

```bash
cd api && go vet ./internal/guilds/... && go test -p 1 -race ./internal/guilds/... -run 'TestGetGuildReadsContestFields|TestContestedReports|TestSetVerifiedForAccount|TestRecomputeMembershipSerialises' -v
```

Expected: PASS, all four new tests, including under `-race` (the advisory lock is a
Postgres-side serialisation, not a Go-side mutex, so `-race` here is only checking the Go
test harness's own concurrency, e.g. the `errs` slice writes — each goroutine writes its own
index, so this is race-clean by construction).

- [ ] **Step 5: Run the full package to confirm zero regressions**

```bash
cd api && go test -p 1 ./internal/guilds/... -v 2>&1 | tail -100
```

Expected: PASS, every prior test (29 from the original plan) plus the four new ones.

- [ ] **Step 6: Commit**

```bash
cd api && git add internal/guilds/store.go internal/guilds/store_test.go
printf 'feat(guilds): add the advisory lock, contest fields, and account-wide verify helper\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-2.txt
git commit -F ../.superpowers/commit-msg-sec-2.txt
```

---

## Task 3: `guilds/claim.go` — R1a/b (Battle.net + rate limit) and R2 (same-account auto-confirm)

**Files:**
- Modify: `api/internal/guilds/claim.go`
- Modify: `api/internal/guilds/claim_test.go`
- Modify: `api/internal/guilds/handler.go` (the `claim` HTTP handler needs to fetch the
  actor's `auth.User` to compute `hasBattleNetIdentity` before calling the now-changed
  `Store.Claim` signature)

**Interfaces:**
- Consumes: `setVerifiedForAccount` (Task 2).
- Produces (signature changes — every existing caller in this codebase is fixed by this
  task; no other task calls these): `func (s *Store) Claim(ctx, guildID, userID int64,
  hasBattleNetIdentity bool) (ClaimResult, error)`; `func AutoConfirmClaimIfPending(ctx, tx
  pgx.Tx, guildID, triggeringUserID int64) error` (Task 4's addon-package call site must be
  updated to pass the triggering account's userID — that update is Task 4's job, not this
  one, since it lives in a different package; this task only changes the function itself).

- [ ] **Step 1: Write the failing tests**

Append to `api/internal/guilds/claim_test.go`:

```go
func TestClaimRefusesAnAccountWithNoBattleNetIdentity(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "email-only-gm@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/emailgm", "leader", false)

	if _, err := s.Claim(ctx, gid, uid, false); !errors.Is(err, ErrNoBattleNetIdentity) {
		t.Fatalf("Claim with no Battle.net identity = %v, want ErrNoBattleNetIdentity", err)
	}
	result, err := s.Claim(ctx, gid, uid, true)
	if err != nil || result.Status != "confirmed" {
		t.Fatalf("Claim with a Battle.net identity = %+v, %v, want confirmed", result, err)
	}
}

func TestClaimRateLimitsToOnePerAccountPerThirtyDays(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "repeat-claimant@example.com")
	g1 := seedGuild(t, pool, "First")
	g2 := seedGuild(t, pool, "Second")
	seedCharacter(t, pool, g1, uid, "us/hardcore/first", "leader", false)
	if _, err := s.Claim(ctx, g1, uid, true); err != nil {
		t.Fatal(err)
	}
	if err := s.ReleaseClaim(ctx, g1, uid, false); err != nil {
		t.Fatal(err)
	}
	seedCharacter(t, pool, g2, uid, "us/hardcore/second", "leader", false)
	if _, err := s.Claim(ctx, g2, uid, true); !errors.Is(err, ErrClaimRateLimited) {
		t.Fatalf("a second claim within 30 days = %v, want ErrClaimRateLimited (even after releasing the first)", err)
	}
}

func TestClaimRefusesASecondGuildWhileAnotherIsStillHeld(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "two-guild-claimant@example.com")
	g1 := seedGuild(t, pool, "First")
	g2 := seedGuild(t, pool, "Second")
	seedCharacter(t, pool, g1, uid, "us/hardcore/heldfirst", "leader", false)
	if _, err := s.Claim(ctx, g1, uid, true); err != nil {
		t.Fatal(err)
	}
	seedCharacter(t, pool, g2, uid, "us/hardcore/heldsecond", "leader", false)
	if _, err := s.Claim(ctx, g2, uid, true); !errors.Is(err, ErrAlreadyClaimsAnotherGuild) {
		t.Fatalf("claiming a second guild while still holding the first = %v, want ErrAlreadyClaimsAnotherGuild", err)
	}
}

func TestAutoConfirmClaimIfPendingRefusesTheSameAccountsSecondForgedCharacter(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	attacker := seedUser(t, pool, "self-confirm@example.com")
	seedCharacter(t, pool, gid, attacker, "us/hardcore/alt1", "officer", false)
	if _, err := s.Claim(ctx, gid, attacker, true); err != nil {
		t.Fatal(err)
	}
	var pendingBy *int64
	pool.QueryRow(ctx, `select claim_pending_by from guilds where id = $1`, gid).Scan(&pendingBy)
	if pendingBy == nil || *pendingBy != attacker {
		t.Fatal("the attacker's claim should be pending")
	}

	// Same account, a second forged character at rank 0 - must NOT
	// auto-confirm its own pending claim.
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := AutoConfirmClaimIfPending(ctx, tx, gid, attacker); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatal("a same-account rank-0 signal must not auto-confirm the account's own pending claim")
	}

	// A genuinely distinct account's rank-0 export still auto-confirms.
	gm := seedUser(t, pool, "real-gm@example.com")
	seedCharacter(t, pool, gid, gm, "us/hardcore/realgm", "leader", false)
	tx, err = pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := AutoConfirmClaimIfPending(ctx, tx, gid, gm); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != attacker {
		t.Fatalf("claimed_by = %v, want the originally pending %d, confirmed by a genuinely distinct account", claimedBy, attacker)
	}
}
```

You must also update every EXISTING test in `claim_test.go` that calls `s.Claim(ctx, gid,
userID)` or `AutoConfirmClaimIfPending(ctx, tx, gid)` — both now take one more argument.
Read the file, find every call site, and add the new argument:
- `s.Claim(ctx, guildID, userID)` → `s.Claim(ctx, guildID, userID, true)` (every existing
  test's claimant is meant to succeed on the merits being tested, so `true` — a linked
  Battle.net identity — is the correct value unless the test is specifically about that
  check, which only the new test above is).
- `AutoConfirmClaimIfPending(ctx, tx, guildID)` → `AutoConfirmClaimIfPending(ctx, tx,
  guildID, <the id of the account whose export triggered it in that test's scenario>)` — in
  `TestAutoConfirmClaimIfPendingFiresOnTheGuildMastersOwnExport` (the existing test from the
  prior plan), that is the `gm` account's id, which the test already has in scope.

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/guilds/... -run 'Claim|AutoConfirm' -v
```

Expected: FAIL to compile (signature mismatches at every existing call site) until Step 1's
full update is applied, then FAIL on the four new tests specifically (undefined
`ErrNoBattleNetIdentity`/`ErrClaimRateLimited`/`ErrAlreadyClaimsAnotherGuild`, wrong
`Claim`/`AutoConfirmClaimIfPending` arity).

- [ ] **Step 3: Rewrite `claim.go`**

Replace the whole file with:

```go
// api/internal/guilds/claim.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	ErrAlreadyClaimed           = errors.New("guilds: already claimed")
	ErrClaimPending             = errors.New("guilds: a claim is already pending")
	ErrNotEligible              = errors.New("guilds: no officer or leader character in this guild")
	ErrSameAccount              = errors.New("guilds: cannot confirm your own pending claim")
	ErrNoPendingClaim           = errors.New("guilds: no pending claim")
	ErrNotClaimant              = errors.New("guilds: not the claimant")
	ErrNoBattleNetIdentity      = errors.New("guilds: claiming as guild master requires a linked Battle.net account")
	ErrClaimRateLimited         = errors.New("guilds: only one claim attempt per account every 30 days")
	ErrAlreadyClaimsAnotherGuild = errors.New("guilds: this account already holds another guild's claim")
)

// claimRateLimitWindow and claimRateLimitReason: an account may attempt
// at most one claim (successful or pending) per rolling window, and may
// hold at most one claimed guild at a time - both per the 2026-09-21
// security review response (spec §2.4's amendment).
const claimRateLimitWindow = 30 * 24 * time.Hour

// ClaimResult is what Claim answers with.
type ClaimResult struct {
	Status    string     `json:"status"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// eligibleClaimRank reads the caller's highest raw rank in guildID
// straight from guild_characters, deliberately bypassing the
// verified-only guild_members/GuildRank: claiming is itself the
// corroboration mechanism, so it does not require verified_at
// beforehand.
func (s *Store) eligibleClaimRank(ctx context.Context, guildID, userID int64) (string, bool, error) {
	var rank string
	err := s.Pool.QueryRow(ctx,
		`select rank from guild_characters where guild_id = $1 and user_id = $2 and rank in ('officer', 'leader')
		 order by (rank = 'leader') desc limit 1`, guildID, userID).Scan(&rank)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("guilds: eligible claim rank: %w", err)
	}
	return rank, true, nil
}

// checkClaimRateLimit enforces the 2026-09-21 hardening: at most one
// currently-claimed guild per account, and at most one claim attempt
// (successful or pending) per account per rolling 30 days. Checked
// before either branch of Claim acts. The "already claims another
// guild" check runs first: holding a claim always also means a recent
// attempt exists, so checking attempts first would make the
// already-claims case unreachable in practice - a caller who still
// holds a guild always hears about that specifically, not a generic
// rate limit, even though both are true.
func (s *Store) checkClaimRateLimit(ctx context.Context, userID int64) error {
	var claimedElsewhere int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from guilds where claimed_by = $1`, userID).Scan(&claimedElsewhere); err != nil {
		return fmt.Errorf("guilds: claim rate limit: %w", err)
	}
	if claimedElsewhere > 0 {
		return ErrAlreadyClaimsAnotherGuild
	}
	var recent int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from guild_claim_attempts where user_id = $1 and attempted_at >= now() - interval '30 days'`,
		userID).Scan(&recent); err != nil {
		return fmt.Errorf("guilds: claim rate limit: %w", err)
	}
	if recent > 0 {
		return ErrClaimRateLimited
	}
	return nil
}

// Claim starts or completes a claim on guildID for userID. A
// guild-master-rank character claims immediately - the bootstrap for an
// unclaimed guild, since it has no verified member who could vouch for
// a second signal (spec §2.4's amendment records why this stays
// single-signal rather than being closed outright) - hardened by
// requiring a linked Battle.net identity and the account-wide rate
// limit above. An officer's claim goes pending until a second officer
// confirms it, or the guild master's own export does
// (AutoConfirmClaimIfPending).
func (s *Store) Claim(ctx context.Context, guildID, userID int64, hasBattleNetIdentity bool) (ClaimResult, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return ClaimResult{}, err
	}
	if g.ClaimedBy != nil {
		return ClaimResult{}, ErrAlreadyClaimed
	}
	if g.pendingActive(time.Now()) {
		return ClaimResult{}, ErrClaimPending
	}
	rank, ok, err := s.eligibleClaimRank(ctx, guildID, userID)
	if err != nil {
		return ClaimResult{}, err
	}
	if !ok {
		return ClaimResult{}, ErrNotEligible
	}
	if rank == "leader" && !hasBattleNetIdentity {
		return ClaimResult{}, ErrNoBattleNetIdentity
	}
	if err := s.checkClaimRateLimit(ctx, userID); err != nil {
		return ClaimResult{}, err
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if _, err := tx.Exec(ctx, `insert into guild_claim_attempts (user_id) values ($1)`, userID); err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: record attempt: %w", err)
	}

	if rank == "leader" {
		if _, err := tx.Exec(ctx, `update guilds set claimed_by = $2 where id = $1`, guildID, userID); err != nil {
			return ClaimResult{}, fmt.Errorf("guilds: claim: %w", err)
		}
		if err := setVerifiedForAccount(ctx, tx, guildID, userID, "claim"); err != nil {
			return ClaimResult{}, err
		}
		if err := RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
			return ClaimResult{}, err
		}
		if err := tx.Commit(ctx); err != nil {
			return ClaimResult{}, fmt.Errorf("guilds: claim: commit: %w", err)
		}
		return ClaimResult{Status: "confirmed"}, nil
	}

	if _, err := tx.Exec(ctx,
		`update guilds set claim_pending_by = $2, claim_requested_at = now() where id = $1`,
		guildID, userID); err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return ClaimResult{}, fmt.Errorf("guilds: claim: commit: %w", err)
	}
	expires := time.Now().Add(ClaimPendingTTL)
	return ClaimResult{Status: "pending", ExpiresAt: &expires}, nil
}

// ConfirmClaim completes a pending officer claim: a second, distinct
// officer/leader account vouches for the pending claimant. Returns the
// claimant's user id so the handler can look up their battletag.
func (s *Store) ConfirmClaim(ctx context.Context, guildID, confirmerID int64) (int64, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return 0, err
	}
	if !g.pendingActive(time.Now()) {
		return 0, ErrNoPendingClaim
	}
	claimant := *g.ClaimPendingBy
	if claimant == confirmerID {
		return 0, ErrSameAccount
	}
	if _, ok, err := s.eligibleClaimRank(ctx, guildID, confirmerID); err != nil {
		return 0, err
	} else if !ok {
		return 0, ErrNotEligible
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("guilds: confirm claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`update guilds set claimed_by = $2, claim_pending_by = null, claim_requested_at = null where id = $1`,
		guildID, claimant); err != nil {
		return 0, fmt.Errorf("guilds: confirm claim: %w", err)
	}
	if err := setVerifiedForAccount(ctx, tx, guildID, claimant, "claim"); err != nil {
		return 0, err
	}
	if err := RecomputeMembership(ctx, tx, guildID, &claimant); err != nil {
		return 0, err
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("guilds: confirm claim: commit: %w", err)
	}
	return claimant, nil
}

// AutoConfirmClaimIfPending confirms a pending officer claim the moment
// any account's export shows guild-master rank for guildID, per the
// design's "failing that, by the guild master's export" - the guild
// master need never call the confirm endpoint themselves. Requires
// triggeringUserID (the account whose export produced the rank-0
// signal) to differ from claim_pending_by, mirroring ConfirmClaim's own
// ErrSameAccount check - a single account must not be able to
// self-confirm its own pending claim with a second forged character
// (2026-09-21 security review finding). A pending claim older than
// ClaimPendingTTL has already expired and is left untouched; a fresh
// Claim call is what clears it. Called from addon.Store.syncGuild in
// the same transaction as the character write that made this account's
// rank "leader".
func AutoConfirmClaimIfPending(ctx context.Context, tx pgx.Tx, guildID, triggeringUserID int64) error {
	var claimant int64
	err := tx.QueryRow(ctx,
		`update guilds set claimed_by = claim_pending_by, claim_pending_by = null, claim_requested_at = null
		 where id = $1 and claim_pending_by is not null and claim_pending_by != $2
		   and claim_requested_at >= now() - interval '14 days'
		 returning claimed_by`, guildID, triggeringUserID).Scan(&claimant)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("guilds: auto-confirm claim for guild %d: %w", guildID, err)
	}
	if err := setVerifiedForAccount(ctx, tx, guildID, claimant, "claim"); err != nil {
		return err
	}
	return RecomputeMembership(ctx, tx, guildID, &claimant)
}

// ReleaseClaim clears guildID's claim. Only the current claimant or a
// moderator may do it.
func (s *Store) ReleaseClaim(ctx context.Context, guildID, userID int64, moderator bool) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	if !moderator && (g.ClaimedBy == nil || *g.ClaimedBy != userID) {
		return ErrNotClaimant
	}
	if _, err := s.Pool.Exec(ctx, `update guilds set claimed_by = null where id = $1`, guildID); err != nil {
		return fmt.Errorf("guilds: release claim: %w", err)
	}
	return nil
}
```

- [ ] **Step 4: Update `handler.go`'s `claim` HTTP handler**

The handler must fetch the actor's `auth.User` to compute `hasBattleNetIdentity` before
calling the now-three-argument `Store.Claim`, and map the two new error cases. Replace the
existing `claim` function in `handler.go` with:

```go
func (s *Service) claim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	u, err := s.Accounts.User(r.Context(), actor.UserID)
	if err != nil {
		s.fail(w, r, "claim", err, "could not claim that guild just now")
		return
	}
	result, err := s.Store.Claim(r.Context(), guildID, actor.UserID, u.BnetSub != "")
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrAlreadyClaimed), errors.Is(err, ErrClaimPending), errors.Is(err, ErrAlreadyClaimsAnotherGuild):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "that guild already has a claim, or you already hold another guild's claim", nil)
	case errors.Is(err, ErrClaimRateLimited):
		httpx.WriteError(w, r, http.StatusTooManyRequests, "rate_limited",
			"you may only attempt one guild claim every 30 days", nil)
	case errors.Is(err, ErrNoBattleNetIdentity):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"claiming as guild master requires a linked Battle.net account", nil)
	case errors.Is(err, ErrNotEligible):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you need an officer or leader character in this guild to claim it", nil)
	case err != nil:
		s.fail(w, r, "claim", err, "could not claim that guild just now")
	default:
		s.logger().Info("guilds", "op", "claim", "guild_id", guildID, "user_id", actor.UserID, "status", result.Status)
		httpx.WriteOK(w, r, http.StatusOK, result)
	}
}
```

- [ ] **Step 5: Run to verify it passes**

```bash
cd api && go vet ./internal/guilds/... && go test -p 1 ./internal/guilds/... -v 2>&1 | tail -120
```

Expected: PASS, every test in the package — the four new tests plus every pre-existing one
with its call sites updated.

- [ ] **Step 6: Commit**

```bash
cd api && git add internal/guilds/claim.go internal/guilds/claim_test.go internal/guilds/handler.go
printf 'feat(guilds): require Battle.net identity and rate-limit the claim flow; fix the same-account auto-confirm bypass\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-3.txt
git commit -F ../.superpowers/commit-msg-sec-3.txt
```

---

## Task 4: `addon` package — case-insensitive guild identity, boundary validation, the R2 call-site fix, and the concurrency test

**Files:**
- Create: `api/internal/addon/guildname.go`
- Create: `api/internal/addon/guildname_test.go`
- Modify: `api/internal/addon/addon.go`
- Modify: `api/internal/addon/addon_test.go`
- Modify: `api/cmd/api/main.go` (one field added to the existing `addon.Store{...}` literal)

**Interfaces:**
- Consumes: `guilds.AutoConfirmClaimIfPending`'s new four-argument signature (Task 3).
- Produces: `func validateGuildName(name string) (string, bool)`; `addon.Store` gains `Log
  *slog.Logger` and an unexported `logger()` method; `resolveGuild` now matches/inserts
  case-insensitively.

- [ ] **Step 1: Write the failing `validateGuildName` tests**

```go
// api/internal/addon/guildname_test.go
package addon

import "testing"

func TestValidateGuildNameNormalisesAndBounds(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    string
		wantOK  bool
	}{
		{name: "ordinary name passes through", input: "Iron Vanguard", want: "Iron Vanguard", wantOK: true},
		{name: "leading and trailing space trimmed", input: "  Iron Vanguard  ", want: "Iron Vanguard", wantOK: true},
		{name: "exactly 24 runes is allowed", input: "123456789012345678901234", want: "123456789012345678901234", wantOK: true},
		{name: "25 runes is refused", input: "1234567890123456789012345", wantOK: false},
		{name: "a control character is refused", input: "Iron\tVanguard", wantOK: false},
		{name: "empty after trimming is refused", input: "   ", wantOK: false},
		{name: "invalid UTF-8 is refused", input: string([]byte{0xff, 0xfe}), wantOK: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := validateGuildName(c.input)
			if ok != c.wantOK {
				t.Fatalf("ok = %v, want %v", ok, c.wantOK)
			}
			if ok && got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestValidateGuildNameNormalisesToNFC(t *testing.T) {
	// "e" + combining acute accent (NFD) must normalise to the single
	// precomposed "é" (NFC) so two exports that differ only in Unicode
	// normalisation form resolve to the same guild.
	decomposed := "Café" // "Cafe" + combining acute over the e
	precomposed := "Café"      // "Café"
	got, ok := validateGuildName(decomposed)
	if !ok {
		t.Fatal("a decomposed-but-otherwise-valid name should be accepted")
	}
	if got != precomposed {
		t.Fatalf("got %q (%d runes), want the NFC form %q (%d runes)", got, len([]rune(got)), precomposed, len([]rune(precomposed)))
	}
}
```

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && go test ./internal/addon/... -run TestValidateGuildName -v
```

Expected: FAIL to compile — `validateGuildName` undefined.

- [ ] **Step 3: Write `guildname.go`**

```go
// api/internal/addon/guildname.go
package addon

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// maxGuildNameRunes is the game's own guild name limit.
const maxGuildNameRunes = 24

// validateGuildName normalises a decoded guild name (Unicode NFC,
// trimmed) and rejects one that cannot be a real WoW guild name: not
// valid UTF-8 (a hand-crafted percent-encoded sequence can decode to
// bytes that are not), empty after trimming, longer than the game's own
// 24-character limit, or containing a control character. Part of the
// 2026-09-21 security review response (spec §3.3's amendment) - a
// malformed or oversized name must never reach a database write
// unvalidated.
func validateGuildName(name string) (string, bool) {
	if !utf8.ValidString(name) {
		return "", false
	}
	normalized := strings.TrimSpace(norm.NFC.String(name))
	if normalized == "" || utf8.RuneCountInString(normalized) > maxGuildNameRunes {
		return "", false
	}
	for _, r := range normalized {
		if unicode.IsControl(r) {
			return "", false
		}
	}
	return normalized, true
}
```

- [ ] **Step 4: Run to verify it passes**

```bash
cd api && go test ./internal/addon/... -run TestValidateGuildName -v
```

Expected: PASS, all cases including the NFC normalisation test.

- [ ] **Step 5: Write the failing `resolveGuild` case-insensitivity test and the rank-bound/logging test**

Append to `api/internal/addon/addon_test.go`:

```go
func TestPutExportsResolvesGuildsCaseInsensitively(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:0"},
	}); err != nil {
		t.Fatal(err)
	}
	alt := h.owner + 1
	if _, err := h.pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'alt-case@example.com') on conflict (id) do nothing`, alt); err != nil {
		t.Fatal(err)
	}
	if err := h.store.PutExports(ctx, alt, []Export{
		{Name: "Caseshifter", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|guild=IRON%20VANGUARD:5"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guilds where region = 'us' and ruleset = 'hardcore'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("guilds rows = %d, want 1 (case-insensitive match, not a second decoy guild)", n)
	}
	var name string
	if err := h.pool.QueryRow(ctx, `select name from guilds where region = 'us' and ruleset = 'hardcore'`).Scan(&name); err != nil {
		t.Fatal(err)
	}
	if name != "Iron Vanguard" {
		t.Fatalf("name = %q, want the first writer's casing, Iron Vanguard", name)
	}
}

func TestPutExportsRejectsAnOutOfRangeRankIndexAsANoOpNotAnAbort(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "Baelgrim", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Overflow:99999"},
	})
	if err != nil {
		t.Fatalf("an out-of-range rank index must be a silent no-op, not a batch-aborting error: %v", err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/baelgrim'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("no guild_characters row should have been created for the out-of-range rank")
	}
	if err := h.pool.QueryRow(ctx, `select count(*) from guilds where name = 'Overflow'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("the guild itself should never have been created from invalid data")
	}
}

func TestPutExportsContinuesTheBatchPastOneBadCharacter(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "BadOne", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Overflow:99999"},
		{Name: "GoodTwo", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:mage:gnome:0/0/0:|guild=RealGuild:0"},
	}); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/goodtwo'`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatal("the second, valid character in the same batch must still sync, even though the first was invalid")
	}
}
```

- [ ] **Step 6: Run to verify these fail**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/addon/... -run 'CaseInsensitively|OutOfRangeRankIndex|ContinuesTheBatch' -v
```

Expected: FAIL — today, `99999` overflows `rank_index smallint` and aborts the whole
`PutExports` call with an error (the pre-fix behavior the review flagged); case-variant
names currently create two `guilds` rows.

- [ ] **Step 7: Add `Log` to `addon.Store` and rewrite `resolveGuild`/`syncGuild`**

In `addon.go`, change the `Store` struct and add a logger helper (mirroring `Service`'s
existing pattern in the same file):

```go
// Store is the addon's two tables.
type Store struct {
	Pool *pgxpool.Pool
	// Log is used only for warning about a malformed guild= section
	// during a sync (validateGuildName/rank-bound failures) - never for
	// routine operation. Defaults to slog.Default() when nil.
	Log *slog.Logger
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}
```

Replace `resolveGuild` with a case-insensitive, race-safe version:

```go
// resolveGuild finds or creates the guild an export names, matching
// case-insensitively on (region, ruleset, lower(name)) - two exports
// differing only in casing must resolve to the same guilds row, the
// same way the pre-existing public guild page already matches (2026-
// 09-21 security review response, spec §3.3's amendment). The first
// writer's casing is kept as the display name; a concurrent insert
// racing on the same case-insensitive name is tolerated by falling back
// to the row the winner created.
func resolveGuild(ctx context.Context, tx pgx.Tx, region, ruleset, name string) (id int64, officerMax int, err error) {
	err = tx.QueryRow(ctx,
		`select id, officer_max_rank_index from guilds where region = $1 and ruleset = $2 and lower(name) = lower($3)`,
		region, ruleset, name).Scan(&id, &officerMax)
	if err == nil {
		return id, officerMax, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return 0, 0, fmt.Errorf("addon: read guild %s: %w", name, err)
	}
	err = tx.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ($1, $2, $3)
		 on conflict (region, ruleset, (lower(name))) do nothing
		 returning id, officer_max_rank_index`,
		region, ruleset, name).Scan(&id, &officerMax)
	if errors.Is(err, pgx.ErrNoRows) {
		// A concurrent insert of a case-variant name won the race; read
		// the row it created.
		err = tx.QueryRow(ctx,
			`select id, officer_max_rank_index from guilds where region = $1 and ruleset = $2 and lower(name) = lower($3)`,
			region, ruleset, name).Scan(&id, &officerMax)
	}
	if err != nil {
		return 0, 0, fmt.Errorf("addon: create/read guild %s: %w", name, err)
	}
	return id, officerMax, nil
}
```

Add the rank-index bound constant, and update `syncGuild` to validate the name and rank
before ever calling `resolveGuild`, logging (not erroring) on a malformed section, and
update the `AutoConfirmClaimIfPending` call site for its new four-argument signature:

```go
// maxRankIndex is a WoW guild's highest real rank index (10 ranks, 0-9).
const maxRankIndex = 9
```

Replace the body of `syncGuild` from the `name, rankIndex, ok := ParseFS1Guild(export)` line
through the `!ok` branch with:

```go
	name, rankIndex, ok := ParseFS1Guild(export)
	if ok {
		var validName bool
		name, validName = validateGuildName(name)
		if !validName || rankIndex < 0 || rankIndex > maxRankIndex {
			s.logger().Warn("addon", "op", "guild_sync", "character_key", key,
				"reason", "invalid guild= section: bad name or rank index outside 0-9")
			ok = false
		}
	}
	if !ok {
		if prevGuildID == nil {
			return nil
		}
		if _, err := tx.Exec(ctx, `delete from guild_characters where character_key = $1`, key); err != nil {
			return fmt.Errorf("addon: clear guild for %s: %w", key, err)
		}
		return afterGuildChange(ctx, tx, *prevGuildID, prevUserID)
	}
```

And change the `AutoConfirmClaimIfPending` call site (unchanged surrounding code, just the
new argument):

```go
	if rank == "leader" {
		if err := guilds.AutoConfirmClaimIfPending(ctx, tx, guildID, userID); err != nil {
			return err
		}
	}
```

- [ ] **Step 8: Write the concurrency test (spec §5's "concurrent PutExports plus approve")**

Append to `api/internal/addon/addon_test.go`:

```go
func TestConcurrentPutExportsAndApproveDoNotLoseAWrite(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	if err := h.store.PutExports(ctx, h.owner, []Export{
		{Name: "RaceMain", Region: "us", Ruleset: "hardcore",
			Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=RaceGuild:5"},
	}); err != nil {
		t.Fatal(err)
	}
	var gid int64
	if err := h.pool.QueryRow(ctx, `select id from guilds where name = 'RaceGuild'`).Scan(&gid); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank) values ($1, 'us/hardcore/racealt', $2, 'member')`,
		gid, h.owner); err != nil {
		t.Fatal(err)
	}
	guildStore := &guilds.Store{Pool: h.pool}

	var wg sync.WaitGroup
	var err1, err2 error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err1 = h.store.PutExports(ctx, h.owner, []Export{
			{Name: "RaceMain", Region: "us", Ruleset: "hardcore",
				Export: "FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=RaceGuild:5"},
		})
	}()
	go func() {
		defer wg.Done()
		err2 = guildStore.ApproveCharacter(ctx, gid, "us/hardcore/racealt")
	}()
	wg.Wait()
	if err1 != nil {
		t.Fatalf("concurrent PutExports: %v", err1)
	}
	if err2 != nil {
		t.Fatalf("concurrent ApproveCharacter: %v", err2)
	}

	// The advisory lock in RecomputeMembership serialises the two racing
	// recomputes; the final guild_members row must reflect BOTH
	// guild_characters rows correctly, not a lost update from one racing
	// past the other's stale snapshot.
	var rank string
	var verified bool
	if err := h.pool.QueryRow(ctx,
		`select rank, verified_at is not null from guild_members where guild_id = $1 and user_id = $2`,
		gid, h.owner).Scan(&rank, &verified); err != nil {
		t.Fatal(err)
	}
	if !verified {
		t.Fatal("the approved character's verification must not have been lost to the race")
	}
}
```

Add `"sync"` and `"github.com/jhunthrop/foreversixty/api/internal/guilds"` to
`addon_test.go`'s import block if not already present (the `guilds` import is already there
from the prior plan's Task 8 auto-confirm integration test).

- [ ] **Step 9: Update `main.go`**

In `api/cmd/api/main.go`, the existing `Addon: &addon.Service{Store: &addon.Store{Pool:
pool}, ...}` literal's inner `addon.Store{Pool: pool}` gains `Log: log`:

```go
		Addon: &addon.Service{
			Store: &addon.Store{Pool: pool, Log: log}, Builds: buildStore, Data: treeData, Log: log,
		},
```

- [ ] **Step 10: Run to verify everything passes**

```bash
cd api && go vet ./internal/addon/... ./cmd/api/... && \
  go test -p 1 ./internal/addon/... -v 2>&1 | tail -150
```

Expected: PASS, every test in the package — the new `validateGuildName` tests, the three new
`PutExports` boundary-validation/case-insensitivity tests, the concurrency test, and every
pre-existing test with zero regressions.

- [ ] **Step 11: Commit**

```bash
cd api && git add internal/addon/guildname.go internal/addon/guildname_test.go \
  internal/addon/addon.go internal/addon/addon_test.go cmd/api/main.go
printf 'feat(addon): case-insensitive guild identity, boundary validation, and the R2 call-site fix\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-4.txt
git commit -F ../.superpowers/commit-msg-sec-4.txt
```

---

## Task 5: `guilds/contest.go` (new) — the claim contest/resolve flow

**Files:**
- Create: `api/internal/guilds/contest.go`
- Create: `api/internal/guilds/contest_test.go`
- Modify: `api/internal/guilds/handler.go` (two new routes)

**Interfaces:**
- Consumes: `Guild.ClaimContestedAt`/`ClaimContestedBy`, `setVerifiedForAccount` (Task 2);
  `eligibleClaimRank`, `ErrSameAccount`, `ErrNotEligible` (Task 3, unchanged by this task).
- Produces: `type ClaimStateView struct{ State string; Since *time.Time }`;
  `func claimState(g Guild, now time.Time) ClaimStateView` (used by Task 7's settings fix
  and Task 8's home fix — both land after this task); `Store.ContestClaim`,
  `Store.ResolveClaim`; `contestClaim`/`resolveClaim` HTTP handlers;
  `ErrNoActiveClaim`, `ErrAlreadyContested`, `ErrInvalidOutcome`.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/guilds/contest_test.go
package guilds

import (
	"context"
	"errors"
	"testing"
)

func TestContestClaimRequiresAnActiveClaim(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	contester := seedUser(t, pool, "contester1@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/contester1", "officer", false)

	if err := s.ContestClaim(ctx, gid, contester); !errors.Is(err, ErrNoActiveClaim) {
		t.Fatalf("contesting an unclaimed guild = %v, want ErrNoActiveClaim", err)
	}
}

func TestContestClaimNeedsARawOfficerOrLeaderCharacterOnADifferentAccount(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/claimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}

	if err := s.ContestClaim(ctx, gid, claimant); !errors.Is(err, ErrSameAccount) {
		t.Fatalf("the claimant contesting their own claim = %v, want ErrSameAccount", err)
	}

	noStanding := seedUser(t, pool, "nostanding@example.com")
	seedCharacter(t, pool, gid, noStanding, "us/hardcore/nostanding", "member", false)
	if err := s.ContestClaim(ctx, gid, noStanding); !errors.Is(err, ErrNotEligible) {
		t.Fatalf("a plain member contesting = %v, want ErrNotEligible", err)
	}

	contester := seedUser(t, pool, "realcontester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/realcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}
	var contestedBy *int64
	pool.QueryRow(ctx, `select claim_contested_by from guilds where id = $1`, gid).Scan(&contestedBy)
	if contestedBy == nil || *contestedBy != contester {
		t.Fatalf("claim_contested_by = %v, want %d", contestedBy, contester)
	}

	second := seedUser(t, pool, "secondcontester@example.com")
	seedCharacter(t, pool, gid, second, "us/hardcore/secondcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, second); !errors.Is(err, ErrAlreadyContested) {
		t.Fatalf("contesting an already-contested claim = %v, want ErrAlreadyContested", err)
	}
}

func TestResolveClaimUphold(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "upheld-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/upheldclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "upheld-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/upheldcontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}

	if err := s.ResolveClaim(ctx, gid, "uphold"); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	var contestedAt any
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	pool.QueryRow(ctx, `select claim_contested_at from guilds where id = $1`, gid).Scan(&contestedAt)
	if claimedBy == nil || *claimedBy != claimant {
		t.Fatalf("claimed_by = %v, want unchanged (%d)", claimedBy, claimant)
	}
	if contestedAt != nil {
		t.Fatal("claim_contested_at should be cleared after an uphold")
	}
}

func TestResolveClaimReleaseUnverifiesOnlyClaimSourcedRows(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "released-claimant@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/releasedclaimant", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	// A second, officer-verified row on a DIFFERENT account, verified via
	// approval rather than the claim - must survive the release.
	other := seedUser(t, pool, "officer-verified@example.com")
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, verified_at, verified_by)
		 values ($1, 'us/hardcore/officerverified', $2, 'member', now(), 'officer')`, gid, other); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "release-contester@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/releasecontester", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}

	if err := s.ResolveClaim(ctx, gid, "release"); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy != nil {
		t.Fatal("claimed_by should be cleared after a release")
	}
	var claimantVerified, otherVerified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/releasedclaimant'`).Scan(&claimantVerified)
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/officerverified'`).Scan(&otherVerified)
	if claimantVerified {
		t.Fatal("the claim-sourced row should be un-verified after a release")
	}
	if !otherVerified {
		t.Fatal("the officer-sourced row must survive the release untouched")
	}
}

func TestResolveClaimTransferMovesTheClaimAndVerification(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	claimant := seedUser(t, pool, "transferred-from@example.com")
	seedCharacter(t, pool, gid, claimant, "us/hardcore/transferredfrom", "leader", false)
	if _, err := s.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}
	contester := seedUser(t, pool, "transferred-to@example.com")
	seedCharacter(t, pool, gid, contester, "us/hardcore/transferredto", "officer", false)
	if err := s.ContestClaim(ctx, gid, contester); err != nil {
		t.Fatal(err)
	}

	if err := s.ResolveClaim(ctx, gid, "transfer"); err != nil {
		t.Fatal(err)
	}
	var claimedBy *int64
	pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != contester {
		t.Fatalf("claimed_by = %v, want the contesting account %d", claimedBy, contester)
	}
	var contesterVerified, originalVerified bool
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/transferredto'`).Scan(&contesterVerified)
	pool.QueryRow(ctx, `select verified_at is not null from guild_characters where character_key = 'us/hardcore/transferredfrom'`).Scan(&originalVerified)
	if !contesterVerified {
		t.Fatal("the new claimant's character should be verified")
	}
	if originalVerified {
		t.Fatal("the original claimant's claim-sourced verification should not survive the transfer")
	}
}

func TestClaimStateReflectsEachPhase(t *testing.T) {
	now := time.Now()
	unclaimed := Guild{}
	if got := claimState(unclaimed, now); got.State != "unclaimed" {
		t.Fatalf("state = %q, want unclaimed", got.State)
	}
	pendingBy := int64(1)
	requested := now
	pending := Guild{ClaimPendingBy: &pendingBy, ClaimRequestedAt: &requested}
	if got := claimState(pending, now); got.State != "pending" {
		t.Fatalf("state = %q, want pending", got.State)
	}
	claimedByID := int64(2)
	claimed := Guild{ClaimedBy: &claimedByID}
	if got := claimState(claimed, now); got.State != "claimed" {
		t.Fatalf("state = %q, want claimed", got.State)
	}
	contested := Guild{ClaimedBy: &claimedByID, ClaimContestedAt: &requested}
	if got := claimState(contested, now); got.State != "contested" {
		t.Fatalf("state = %q, want contested", got.State)
	}
}
```

Add `"time"` to `contest_test.go`'s import block (alongside `context`, `errors`, `testing`).

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/guilds/... -run 'ContestClaim|ResolveClaim|ClaimStateReflects' -v
```

Expected: FAIL to compile.

- [ ] **Step 3: Write `contest.go`**

```go
// api/internal/guilds/contest.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

var (
	ErrNoActiveClaim    = errors.New("guilds: no claim to contest")
	ErrAlreadyContested = errors.New("guilds: claim already contested")
	ErrInvalidOutcome   = errors.New("guilds: invalid resolution outcome")
)

// ClaimStateView is the claim state object GET /v1/guilds/{id}/settings
// and GET /v1/guilds/{id}/home both expose (2026-09-21 security review
// response, spec §2.4's amendment).
type ClaimStateView struct {
	State string     `json:"state"`
	Since *time.Time `json:"since,omitempty"`
}

// claimState derives the four-phase claim state a guild is in as of now.
func claimState(g Guild, now time.Time) ClaimStateView {
	switch {
	case g.ClaimContestedAt != nil:
		return ClaimStateView{State: "contested", Since: g.ClaimContestedAt}
	case g.ClaimedBy != nil:
		return ClaimStateView{State: "claimed"}
	case g.pendingActive(now):
		return ClaimStateView{State: "pending", Since: g.ClaimRequestedAt}
	default:
		return ClaimStateView{State: "unclaimed"}
	}
}

// ContestClaim flags guildID's current claim (pending or claimed) as
// disputed by contesterID, freezing the claimant's officer powers until
// a moderator resolves it. The contester must hold a raw officer/leader
// rank character in this guild (eligibleClaimRank - the same raw,
// unverified-tolerant check the claim flow itself uses, since disputing
// a possibly-forged claim cannot itself require verification the dispute
// exists to question) on a different account from the current
// claimant/pending claimant.
func (s *Store) ContestClaim(ctx context.Context, guildID, contesterID int64) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	current := g.ClaimedBy
	if current == nil && g.pendingActive(time.Now()) {
		current = g.ClaimPendingBy
	}
	if current == nil {
		return ErrNoActiveClaim
	}
	if g.ClaimContestedAt != nil {
		return ErrAlreadyContested
	}
	if *current == contesterID {
		return ErrSameAccount
	}
	if _, ok, err := s.eligibleClaimRank(ctx, guildID, contesterID); err != nil {
		return err
	} else if !ok {
		return ErrNotEligible
	}
	if _, err := s.Pool.Exec(ctx,
		`update guilds set claim_contested_at = now(), claim_contested_by = $2 where id = $1`,
		guildID, contesterID); err != nil {
		return fmt.Errorf("guilds: contest claim: %w", err)
	}
	return nil
}

// ResolveClaim is a moderator's decision on a contested claim: uphold
// (dismiss the contest, claim stands), release (clear the claim and
// un-verify every character whose only verification source was the
// claim), or transfer (move the claim, and the same claim-sourced
// verification, to the contesting account).
func (s *Store) ResolveClaim(ctx context.Context, guildID int64, outcome string) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	if g.ClaimContestedAt == nil {
		return ErrNoActiveClaim
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("guilds: resolve claim: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	switch outcome {
	case "uphold":
		if _, err := tx.Exec(ctx,
			`update guilds set claim_contested_at = null, claim_contested_by = null where id = $1`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (uphold): %w", err)
		}
	case "release":
		if _, err := tx.Exec(ctx,
			`update guilds set claimed_by = null, claim_pending_by = null, claim_requested_at = null,
			   claim_contested_at = null, claim_contested_by = null where id = $1`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (release): %w", err)
		}
		if _, err := tx.Exec(ctx,
			`update guild_characters set verified_at = null, verified_by = null
			 where guild_id = $1 and verified_by = 'claim'`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (release): un-verify: %w", err)
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return err
		}
	case "transfer":
		if g.ClaimContestedBy == nil {
			return fmt.Errorf("guilds: resolve claim (transfer): no contesting account recorded")
		}
		newClaimant := *g.ClaimContestedBy
		if _, err := tx.Exec(ctx,
			`update guilds set claimed_by = $2, claim_pending_by = null, claim_requested_at = null,
			   claim_contested_at = null, claim_contested_by = null where id = $1`, guildID, newClaimant); err != nil {
			return fmt.Errorf("guilds: resolve claim (transfer): %w", err)
		}
		if _, err := tx.Exec(ctx,
			`update guild_characters set verified_at = null, verified_by = null
			 where guild_id = $1 and verified_by = 'claim'`, guildID); err != nil {
			return fmt.Errorf("guilds: resolve claim (transfer): un-verify: %w", err)
		}
		if err := setVerifiedForAccount(ctx, tx, guildID, newClaimant, "claim"); err != nil {
			return err
		}
		if err := RecomputeMembership(ctx, tx, guildID, nil); err != nil {
			return err
		}
	default:
		return ErrInvalidOutcome
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("guilds: resolve claim: commit: %w", err)
	}
	return nil
}

func (s *Service) contestClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	err := s.Store.ContestClaim(r.Context(), guildID, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrNoActiveClaim):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "there is no claim on this guild to contest", nil)
	case errors.Is(err, ErrAlreadyContested):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "this claim is already contested", nil)
	case errors.Is(err, ErrSameAccount):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you cannot contest your own claim", nil)
	case errors.Is(err, ErrNotEligible):
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you need an officer or leader character in this guild to contest its claim", nil)
	case err != nil:
		s.fail(w, r, "contest_claim", err, "could not contest that claim just now")
	default:
		s.logger().Info("guilds", "op", "claim_contest", "guild_id", guildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "contested"})
	}
}

type resolveClaimInput struct {
	Outcome string `json:"outcome"`
}

func (s *Service) resolveClaim(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	if !actor.IsModerator() {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "only a moderator may resolve a contested claim", nil)
		return
	}
	var in resolveClaimInput
	if err := decodeJSON(w, r, &in); err != nil {
		return
	}
	if in.Outcome != "uphold" && in.Outcome != "release" && in.Outcome != "transfer" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a resolution outcome",
			map[string]string{"outcome": "one of uphold, release, transfer"})
		return
	}
	err := s.Store.ResolveClaim(r.Context(), guildID, in.Outcome)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
	case errors.Is(err, ErrNoActiveClaim):
		httpx.WriteError(w, r, http.StatusConflict, "conflict", "there is no contested claim on this guild", nil)
	case err != nil:
		s.fail(w, r, "resolve_claim", err, "could not resolve that claim just now")
	default:
		s.logger().Info("guilds", "op", "claim_resolve", "guild_id", guildID,
			"moderator_id", actor.UserID, "outcome", in.Outcome)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "resolved", "outcome": in.Outcome})
	}
}
```

- [ ] **Step 4: Wire the two routes into `handler.go`'s existing `Mount`**

```go
	mux.HandleFunc("POST /v1/guilds/{id}/claim/contest", auth.RequireSession(s.contestClaim))
	mux.HandleFunc("POST /v1/guilds/{id}/claim/resolve", auth.RequireSession(s.resolveClaim))
```

Add these two lines to the existing `Mount` function, next to the other claim routes. The
prior plan's `Mount` had no per-route rate limiting on `claim`/`confirm`/`release`; match
that for `contest` too (the router-wide per-IP limiter in `server.NewRouter` already covers
every route, `contest`'s DB-level `ErrAlreadyContested`/`ErrNoActiveClaim` checks already
bound repeat calls to a no-op, and a moderator-only `resolve` needs no IP limiter at all).

- [ ] **Step 5: Run to verify it passes**

```bash
cd api && go vet ./internal/guilds/... && go test -p 1 ./internal/guilds/... -v 2>&1 | tail -150
```

Expected: PASS, every test including the seven new ones.

- [ ] **Step 6: Commit**

```bash
cd api && git add internal/guilds/contest.go internal/guilds/contest_test.go internal/guilds/handler.go
printf 'feat(guilds): add the claim contest and moderator-resolve flow\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-5.txt
git commit -F ../.superpowers/commit-msg-sec-5.txt
```

---

## Task 6: `guilds/roster.go` — R3 (rank protects rank) and the contested-claim freeze

**Files:**
- Modify: `api/internal/guilds/roster.go`
- Modify: `api/internal/guilds/roster_test.go`

**Interfaces:**
- Consumes: `Store.contested` (Task 2); `Guild.ClaimedBy` (existing).
- Produces (signature change): `CharacterOwner(ctx, guildID int64, characterKey string)
  (userID int64, rank string, err error)` — was `(int64, error)`; new
  `Store.mayRemoveCharacter(ctx, guildID, actorID, ownerID int64, targetRank string,
  moderator, verifiedOfficer bool) (bool, error)`.

- [ ] **Step 1: Write the failing tests**

Append to `api/internal/guilds/roster_test.go`:

```go
func TestAnOfficerCannotRemoveTheGuildMastersCharacter(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	gm := seedUser(t, h.pool, "protected-gm@example.com")
	seedCharacter(t, h.pool, gid, gm, "us/hardcore/protectedgm", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, gm, gid); err != nil {
		t.Fatal(err)
	}
	officer := seedUser(t, h.pool, "unprivileged-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/unprivilegedofficer", "officer", true)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/protectedgm", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("an officer removing the guild master = %d, want 403", res.StatusCode)
	}
	var claimedBy *int64
	h.pool.QueryRow(ctx, `select claimed_by from guilds where id = $1`, gid).Scan(&claimedBy)
	if claimedBy == nil || *claimedBy != gm {
		t.Fatalf("claimed_by = %v, want unchanged (%d) - the claim must stay intact", claimedBy, gm)
	}
	var n int
	h.pool.QueryRow(ctx, `select count(*) from guild_characters where character_key = 'us/hardcore/protectedgm'`).Scan(&n)
	if n != 1 {
		t.Fatal("the guild master's row must still exist")
	}
}

func TestTheAccountHoldingTheClaimCanRemoveAnOfficerButNotAnotherLeaderRow(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	gm := seedUser(t, h.pool, "gm-removes@example.com")
	seedCharacter(t, h.pool, gid, gm, "us/hardcore/gmremoves", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, gm, gid); err != nil {
		t.Fatal(err)
	}
	officer := seedUser(t, h.pool, "removable-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/removableofficer", "officer", true)
	otherLeader := seedUser(t, h.pool, "other-leader@example.com")
	seedCharacter(t, h.pool, gid, otherLeader, "us/hardcore/otherleader", "leader", true)

	h.actor = auth.Actor{UserID: gm, Role: "user", Method: "session"}
	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/removableofficer", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("the claim-holding account removing an officer = %d, want 200", res.StatusCode)
	}

	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/otherleader", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("the claim-holding account removing a DIFFERENT leader-rank row = %d, want 403", res.StatusCode)
	}
}

func TestAContestedClaimFreezesApproveAndRemove(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	claimant := seedUser(t, h.pool, "frozen-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/frozenclaimant", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1, claim_contested_at = now() where id = $2`, claimant, gid); err != nil {
		t.Fatal(err)
	}
	unverified := seedUser(t, h.pool, "frozen-unverified@example.com")
	seedCharacter(t, h.pool, gid, unverified, "us/hardcore/frozenunverified", "member", false)

	h.actor = auth.Actor{UserID: claimant, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/frozenunverified/approve", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("approve while contested = %d, want 409", res.StatusCode)
	}

	otherOfficer := seedUser(t, h.pool, "frozen-other-officer@example.com")
	seedCharacter(t, h.pool, gid, otherOfficer, "us/hardcore/frozenotherofficer", "officer", true)
	h.actor = auth.Actor{UserID: otherOfficer, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/frozenunverified", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("remove by an officer while contested = %d, want 409", res.StatusCode)
	}

	// Self-removal still works during a contest - leaving is never frozen.
	h.actor = auth.Actor{UserID: unverified, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/frozenunverified", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("self-removal while contested = %d, want 200 (never frozen)", res.StatusCode)
	}
}

func TestApproveCharacterRecordsVerifiedByOfficer(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "verify-source@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/verifysource", "member", false)
	officer := seedUser(t, h.pool, "verify-source-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/verifysourceofficer", "officer", true)

	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/characters/us/hardcore/verifysource/approve", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("approve = %d, want 200", res.StatusCode)
	}
	var by string
	h.pool.QueryRow(ctx, `select verified_by from guild_characters where character_key = 'us/hardcore/verifysource'`).Scan(&by)
	if by != "officer" {
		t.Fatalf("verified_by = %q, want officer", by)
	}
}
```

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/guilds/... -run 'CannotRemoveTheGuildMaster|CanRemoveAnOfficerButNot|FreezesApproveAndRemove|RecordsVerifiedByOfficer' -v
```

Expected: FAIL — today an officer CAN remove the guild master (200, not 403); a contested
claim does not freeze anything yet; `verified_by` is never set by approve.

- [ ] **Step 3: Update `roster.go`**

Change `ApproveCharacter`'s UPDATE to record its source:

```go
	err = tx.QueryRow(ctx,
		`update guild_characters set verified_at = coalesce(verified_at, now()),
		   verified_by = coalesce(verified_by, 'officer')
		 where guild_id = $1 and character_key = $2 returning user_id`,
		guildID, characterKey).Scan(&userID)
```

Change `CharacterOwner` to also return the row's rank:

```go
// CharacterOwner reads which account a guild_characters row belongs to
// and its current rank, so a handler can decide "the character's own
// account" and apply the rank-protects-rank rule (mayRemoveCharacter)
// before acting.
func (s *Store) CharacterOwner(ctx context.Context, guildID int64, characterKey string) (userID int64, rank string, err error) {
	err = s.Pool.QueryRow(ctx,
		`select user_id, rank from guild_characters where guild_id = $1 and character_key = $2`,
		guildID, characterKey).Scan(&userID, &rank)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, "", ErrNotFound
	}
	if err != nil {
		return 0, "", fmt.Errorf("guilds: character owner: %w", err)
	}
	return userID, rank, nil
}
```

Add, after `CharacterOwner`:

```go
// mayRemoveCharacter applies the rank-protects-rank rule (2026-09-21
// security review response, spec §3.3's amendment): a member-rank row
// may be removed by its own account, a verified officer/leader, or a
// moderator; an officer-rank row adds only the account currently
// holding the guild's claim to that list (not any officer); a
// leader-rank row is removable only by its own account or a moderator -
// never by the account holding the claim either, since a second
// leader-rank row belongs to a different real character than the
// claimant's own. This is what stops an officer from stripping the real
// guild master's membership and then claiming the now-unclaimed guild.
func (s *Store) mayRemoveCharacter(ctx context.Context, guildID, actorID, ownerID int64, targetRank string, moderator, verifiedOfficer bool) (bool, error) {
	if moderator || actorID == ownerID {
		return true, nil
	}
	switch targetRank {
	case "leader":
		return false, nil
	case "officer":
		g, err := s.getGuild(ctx, guildID)
		if err != nil {
			return false, err
		}
		return g.ClaimedBy != nil && *g.ClaimedBy == actorID, nil
	default: // member
		return verifiedOfficer, nil
	}
}
```

Replace `approveCharacter`'s body (after the existing `verifiedOfficerOrLeader` authz check
succeeds, before the `s.Store.ApproveCharacter` call) to add the contested-freeze check:

```go
func (s *Service) approveCharacter(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "approve", err, "could not approve that character just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to approve a character", nil)
		return
	}
	if contested, err := s.Store.contested(r.Context(), guildID); err != nil {
		s.fail(w, r, "approve", err, "could not approve that character just now")
		return
	} else if contested {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	err = s.Store.ApproveCharacter(r.Context(), guildID, key)
	switch {
	case errors.Is(err, ErrNotFound):
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on this guild's roster", nil)
	case err != nil:
		s.fail(w, r, "approve", err, "could not approve that character just now")
	default:
		s.logger().Info("guilds", "op", "character_approve", "guild_id", guildID, "character_key", key,
			"user_id", auth.ActorFrom(r.Context()).UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]string{"character_key": key, "status": "approved"})
	}
}
```

Replace `removeCharacter` entirely:

```go
func (s *Service) removeCharacter(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	key, ok := characterKeyFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	owner, targetRank, err := s.Store.CharacterOwner(r.Context(), guildID, key)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on this guild's roster", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "remove", err, "could not remove that character just now")
		return
	}

	selfOrModerator := owner == actor.UserID || actor.IsModerator()
	var verifiedOfficer bool
	if !selfOrModerator {
		verifiedOfficer, err = s.verifiedOfficerOrLeader(r, guildID)
		if err != nil {
			s.fail(w, r, "remove", err, "could not remove that character just now")
			return
		}
	}
	allowed, err := s.Store.mayRemoveCharacter(r.Context(), guildID, actor.UserID, owner, targetRank, actor.IsModerator(), verifiedOfficer)
	if err != nil {
		s.fail(w, r, "remove", err, "could not remove that character just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you do not have standing to remove this character", nil)
		return
	}
	// Leaving is never frozen by a contested claim - only another
	// account's officer action against someone else's row is.
	if !selfOrModerator {
		if contested, err := s.Store.contested(r.Context(), guildID); err != nil {
			s.fail(w, r, "remove", err, "could not remove that character just now")
			return
		} else if contested {
			httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
				"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
			return
		}
	}
	if err := s.Store.RemoveCharacter(r.Context(), guildID, key); err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on this guild's roster", nil)
			return
		}
		s.fail(w, r, "remove", err, "could not remove that character just now")
		return
	}
	s.logger().Info("guilds", "op", "character_remove", "guild_id", guildID, "character_key", key, "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]string{"status": "removed"})
}
```

- [ ] **Step 4: Run to verify it passes**

```bash
cd api && go vet ./internal/guilds/... && go test -p 1 ./internal/guilds/... -v 2>&1 | tail -150
```

Expected: PASS, every test including the four new ones and the pre-existing
`TestRemoveCharacterIsTheOwnerOrAVerifiedOfficer` (a plain member-rank row, still removable
by a verified officer — that leg is unaffected by the rank-protects-rank rule, which only
adds restrictions for `officer`/`leader` rows).

- [ ] **Step 5: Commit**

```bash
cd api && git add internal/guilds/roster.go internal/guilds/roster_test.go
printf 'feat(guilds): rank protects rank on character removal; freeze roster actions during a contested claim\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-6.txt
git commit -F ../.superpowers/commit-msg-sec-6.txt
```

---

## Task 7: `guilds/invite.go`, `jobs.go`, `settings.go` — verified_by propagation, the settings claim-state fix, and remaining contested-freeze checks

**Files:**
- Modify: `api/internal/guilds/invite.go`
- Modify: `api/internal/guilds/invite_test.go`
- Modify: `api/internal/guilds/jobs.go`
- Modify: `api/internal/guilds/jobs_test.go`
- Modify: `api/internal/guilds/settings.go`
- Modify: `api/internal/guilds/settings_test.go`

**Interfaces:**
- Consumes: `Store.contested` (Task 2); `ClaimStateView`, `claimState` (Task 5).
- Produces: no new exported names — this task only changes existing behavior.

These three files each get one small, same-shape change (`verified_by` on the UPDATE that
sets `verified_at`, or a contested-freeze check before an officer-power write), plus
`settings.go`'s specific fix for the previously-unpopulated `ClaimedBy`/`ClaimPending`
fields and the new `Claim` field. Batched into one task because each individual change is a
few lines with its own small test, not because they share a review surface beyond that.

- [ ] **Step 1: Write the failing tests**

Append to `api/internal/guilds/invite_test.go`:

```go
func TestAcceptInviteRecordsVerifiedByInvite(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	gid := seedGuild(t, pool, "Forever")
	token, _, err := s.RotateInvite(ctx, gid)
	if err != nil {
		t.Fatal(err)
	}
	uid := seedUser(t, pool, "invite-source@example.com")
	if _, err := s.AcceptInvite(ctx, token, uid); err != nil {
		t.Fatal(err)
	}
	var by string
	if err := pool.QueryRow(ctx,
		`select verified_by from guild_characters where character_key = $1`, syntheticKey(uid)).Scan(&by); err != nil {
		t.Fatal(err)
	}
	if by != "invite" {
		t.Fatalf("verified_by = %q, want invite", by)
	}
}

func TestRotateInviteIsFrozenDuringAContestedClaim(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "frozen-rotate@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/frozenrotate", "officer", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/invite/rotate", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("invite rotate while contested = %d, want 409", res.StatusCode)
	}
}
```

Append to `api/internal/guilds/jobs_test.go`:

```go
func TestVerifyByLogsRecordsVerifiedByLogs(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "logs-source@example.com")
	gid := seedGuild(t, pool, "Forever")
	seedCharacter(t, pool, gid, uid, "us/hardcore/logssource", "member", false)

	first := time.Now().Add(-20 * 24 * time.Hour)
	for i, days := range []float64{20, 5} {
		id := fmt.Sprintf("logsource%d", i)
		if _, err := pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, 'guild', 'complete', $4)`,
			id, uid, gid, first.Add(time.Duration(20-days)*24*time.Hour)); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, players) values ($1, 0, $2)`,
			id, []string{"us/hardcore/logssource"}); err != nil {
			t.Fatal(err)
		}
	}
	if err := s.VerifyByLogs(ctx); err != nil {
		t.Fatal(err)
	}
	var by string
	if err := pool.QueryRow(ctx,
		`select verified_by from guild_characters where character_key = 'us/hardcore/logssource'`).Scan(&by); err != nil {
		t.Fatal(err)
	}
	if by != "logs" {
		t.Fatalf("verified_by = %q, want logs", by)
	}
}
```

Append to `api/internal/guilds/settings_test.go`:

```go
func TestSettingsExposesClaimStateAndPopulatesClaimant(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	claimant := seedUser(t, h.pool, "settings-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/settingsclaimant", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, claimant, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: claimant, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/settings", gid), "")
	var view SettingsView
	h.data(res, &view)
	if view.Claim.State != "claimed" {
		t.Fatalf("claim.state = %q, want claimed", view.Claim.State)
	}
	if view.ClaimedBy == nil || view.ClaimedBy.Battletag == "" {
		t.Fatalf("claimed_by = %v, want populated (this was the bug the review found: always null)", view.ClaimedBy)
	}
}

func TestPatchSettingsIsFrozenDuringAContestedClaim(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	officer := seedUser(t, h.pool, "frozen-settings@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/frozensettings", "officer", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claim_contested_at = now() where id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}
	res := h.do(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/settings", gid), `{"default_visibility":"public"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("patch settings while contested = %d, want 409", res.StatusCode)
	}
}
```

`settings_test.go` needs the HTTP harness (`newHTTPHarness`, from Task 11 of the prior
plan) and `auth`/`net/http`/`fmt` imports — check whether it already imports the `auth`
package and `net/http`; add whichever it's missing (`fmt` is almost certainly already
imported by an existing `Sprintf` call in the file; verify rather than assume).

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/guilds/... -run 'RecordsVerifiedByInvite|FrozenDuringAContestedClaim|RecordsVerifiedByLogs|ExposesClaimState' -v
```

Expected: FAIL — `verified_by` is never set by invite accept or the log-corroboration
sweep; nothing is frozen yet; `SettingsView.Claim`/`ClaimedBy` are unpopulated (the exact
review finding).

- [ ] **Step 3: Update `invite.go`**

Change `AcceptInvite`'s insert to record its source:

```go
	if _, err := tx.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, source, verified_at, verified_by, refreshed_at)
		 values ($1, $2, $3, 'member', 'invite', now(), 'invite', now())
		 on conflict (guild_id, character_key) do update set
		   user_id = excluded.user_id, refreshed_at = now(),
		   verified_at = coalesce(guild_characters.verified_at, now()),
		   verified_by = coalesce(guild_characters.verified_by, 'invite')`,
		g.ID, key, userID); err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: %w", err)
	}
```

Add the contested-freeze check to `rotateInvite`, right after the existing
`verifiedOfficerOrLeader` check succeeds:

```go
	if contested, err := s.Store.contested(r.Context(), guildID); err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	} else if contested {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
```

- [ ] **Step 4: Update `jobs.go`**

Change `VerifyByLogs`'s UPDATE to record its source (unconditional, not `coalesce` — this
statement's own `WHERE gc.verified_at is null` guarantees every row it touches had no prior
`verified_by` either):

```go
	rows, err := tx.Query(ctx, `
		update guild_characters gc
		set verified_at = now(), verified_by = 'logs'
		where gc.verified_at is null
		  and (
		    select count(distinct r2.created_at::date)
		    from fights f2 join reports r2 on r2.id = f2.report_id
		    where r2.guild_id = gc.guild_id
		      and gc.character_key = any(f2.players)
		      and r2.created_at >= now() - interval '30 days'
		  ) >= 2
		returning guild_id, user_id
	`)
```

- [ ] **Step 5: Update `settings.go`**

Add `Claim ClaimStateView` to the `SettingsView` struct:

```go
type SettingsView struct {
	DefaultVisibility   string            `json:"default_visibility"`
	OfficerMaxRankIndex int               `json:"officer_max_rank_index"`
	ClaimedBy           *MemberRef        `json:"claimed_by"`
	ClaimPending        *ClaimPendingView `json:"claim_pending"`
	Claim               ClaimStateView    `json:"claim"`
	Invite              InviteView        `json:"invite"`
}
```

Set it in `Store.Settings`:

```go
func (s *Store) Settings(ctx context.Context, guildID int64) (SettingsView, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return SettingsView{}, err
	}
	return SettingsView{
		DefaultVisibility: g.DefaultVisibility, OfficerMaxRankIndex: g.OfficerMaxRankIndex,
		Invite: InviteView{RotatedAt: g.InviteTokenRotatedAt},
		Claim:  claimState(g, time.Now()),
	}, nil
}
```

Fix `getSettings` to actually populate `ClaimedBy`/`ClaimPending` (the review's exact
"declared but never populated" finding) — the enrichment needs `s.Accounts.User`, which only
the handler (not `Store`) has, so it happens here, reading the guild's raw claim fields via
the package-private `s.Store.getGuild` (same package, callable):

```go
func (s *Service) getSettings(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	allowed, err := s.verifiedOfficerOrModerator(r, guildID)
	if err != nil {
		s.fail(w, r, "settings", err, "could not read those settings just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to see its settings", nil)
		return
	}
	view, err := s.Store.Settings(r.Context(), guildID)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "settings", err, "could not read those settings just now")
		return
	}
	g, err := s.Store.getGuild(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "settings", err, "could not read those settings just now")
		return
	}
	if g.ClaimedBy != nil {
		if u, err := s.Accounts.User(r.Context(), *g.ClaimedBy); err == nil {
			view.ClaimedBy = &MemberRef{Battletag: u.PublicName()}
		}
	}
	if g.pendingActive(time.Now()) {
		if u, err := s.Accounts.User(r.Context(), *g.ClaimPendingBy); err == nil {
			view.ClaimPending = &ClaimPendingView{
				By: MemberRef{Battletag: u.PublicName()}, ExpiresAt: g.ClaimRequestedAt.Add(ClaimPendingTTL),
			}
		}
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}
```

Add the contested-freeze check to `patchSettings`, right after its own
`verifiedOfficerOrModerator` check succeeds, before decoding the body:

```go
	if contested, err := s.Store.contested(r.Context(), guildID); err != nil {
		s.fail(w, r, "patch_settings", err, "could not change those settings just now")
		return
	} else if contested {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
```

- [ ] **Step 6: Run to verify it passes**

```bash
cd api && go vet ./internal/guilds/... && go test -p 1 ./internal/guilds/... -v 2>&1 | tail -180
```

Expected: PASS, every test in the package.

- [ ] **Step 7: Commit**

```bash
cd api && git add internal/guilds/invite.go internal/guilds/invite_test.go \
  internal/guilds/jobs.go internal/guilds/jobs_test.go \
  internal/guilds/settings.go internal/guilds/settings_test.go
printf 'feat(guilds): record verified_by on invite/logs, freeze invite rotate and settings when contested, fix the unpopulated claim fields on settings\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-7.txt
git commit -F ../.superpowers/commit-msg-sec-7.txt
```

---

## Task 8: `guilds/home.go` — R4 (report visibility) and the home claim-state field

**Files:**
- Modify: `api/internal/guilds/home.go`
- Modify: `api/internal/guilds/home_test.go`

**Interfaces:**
- Consumes: `ClaimStateView`, `claimState` (Task 5).
- Produces (signature changes): `HomeReports(ctx, guildID, userID int64, verified bool,
  before *HomeCursor) ([]HomeReport, error)`; `Home(ctx, guildID, userID int64, verified
  bool, before *HomeCursor) (HomeView, error)` — both gain `userID`/`verified`.

- [ ] **Step 1: Write the failing tests**

Append to `api/internal/guilds/home_test.go`:

```go
func TestHomeReportsHidesPrivateAndUnlistedFromOthers(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	owner := seedUser(t, h.pool, "home-visibility-owner@example.com")
	seedCharacter(t, h.pool, gid, owner, "us/hardcore/homevisibilityowner", "member", true)
	viewer := seedUser(t, h.pool, "home-visibility-viewer@example.com")
	seedCharacter(t, h.pool, gid, viewer, "us/hardcore/homevisibilityviewer", "member", true)

	for _, r := range []struct{ id, visibility string }{
		{"homevispublic", "public"}, {"homevisguild", "guild"},
		{"homevisprivate", "private"}, {"homevisunlisted", "unlisted"},
	} {
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, $4, 'complete', now())`, r.id, owner, gid, r.visibility); err != nil {
			t.Fatal(err)
		}
	}

	h.actor = auth.Actor{UserID: viewer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	seen := map[string]bool{}
	for _, r := range view.Reports {
		seen[r.ID] = true
	}
	if !seen["homevispublic"] || !seen["homevisguild"] {
		t.Fatalf("a verified viewer should see public and guild reports: %+v", view.Reports)
	}
	if seen["homevisprivate"] || seen["homevisunlisted"] {
		t.Fatalf("a verified viewer who is not the owner must never see private or unlisted reports: %+v", view.Reports)
	}

	h.actor = auth.Actor{UserID: owner, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	h.data(res, &view)
	seen = map[string]bool{}
	for _, r := range view.Reports {
		seen[r.ID] = true
	}
	for _, id := range []string{"homevispublic", "homevisguild", "homevisprivate", "homevisunlisted"} {
		if !seen[id] {
			t.Fatalf("the owner should see every one of their own reports regardless of visibility: %+v", view.Reports)
		}
	}
}

func TestHomeReportsHidesGuildVisibilityFromAnUnverifiedMember(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	owner := seedUser(t, h.pool, "home-unverified-owner@example.com")
	seedCharacter(t, h.pool, gid, owner, "us/hardcore/homeunverifiedowner", "member", true)
	unverified := seedUser(t, h.pool, "home-unverified-viewer@example.com")
	seedCharacter(t, h.pool, gid, unverified, "us/hardcore/homeunverifiedviewer", "member", false)

	for _, r := range []struct{ id, visibility string }{
		{"unverifiedpublic", "public"}, {"unverifiedguild", "guild"},
	} {
		if _, err := h.pool.Exec(ctx,
			`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
			 values ($1, $2, $3, $4, 'complete', now())`, r.id, owner, gid, r.visibility); err != nil {
			t.Fatal(err)
		}
	}

	h.actor = auth.Actor{UserID: unverified, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	seen := map[string]bool{}
	for _, r := range view.Reports {
		seen[r.ID] = true
	}
	if !seen["unverifiedpublic"] {
		t.Fatal("an unverified member should still see public reports")
	}
	if seen["unverifiedguild"] {
		t.Fatal("an unverified member must not see guild-visibility reports")
	}
}

func TestHomeExposesClaimState(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")
	member := seedUser(t, h.pool, "home-claim-state@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/homeclaimstate", "member", true)
	claimant := seedUser(t, h.pool, "home-claim-state-claimant@example.com")
	seedCharacter(t, h.pool, gid, claimant, "us/hardcore/homeclaimstateclaimant", "leader", true)
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, claimant, gid); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var view HomeView
	h.data(res, &view)
	if view.Claim.State != "claimed" {
		t.Fatalf("claim.state = %q, want claimed", view.Claim.State)
	}
}
```

- [ ] **Step 2: Run to verify it fails**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go test -p 1 ./internal/guilds/... -run 'HomeReportsHides|HomeExposesClaimState' -v
```

Expected: FAIL — today every report with the guild's `guild_id` returns regardless of
visibility, and `HomeView` carries no `claim` field at all.

- [ ] **Step 3: Update `home.go`**

Add `Claim ClaimStateView` to `HomeView`:

```go
type HomeView struct {
	Guild      GuildIdentity  `json:"guild"`
	Claim      ClaimStateView `json:"claim"`
	Reports    []HomeReport   `json:"reports"`
	NextCursor string         `json:"next_cursor,omitempty"`
	Roster     []RosterRow    `json:"roster"`
}
```

Change `HomeReports` to filter by visibility:

```go
// HomeReports lists this guild's reports from the trailing week, newest
// first, keyset-paginated. A member sees their own report regardless of
// visibility, every public report regardless of their own verification,
// and a guild-visible report only once they are verified - never a
// private or unlisted row that is not their own, even once verified: an
// unlisted report is discoverable by every member on this page, which
// defeats its whole point (2026-09-21 security review response, spec
// §3.3's amendment).
func (s *Store) HomeReports(ctx context.Context, guildID, userID int64, verified bool, before *HomeCursor) ([]HomeReport, error) {
	since := time.Now().Add(-homeReportsWindow)
	const columns = `r.id, r.title, r.created_at,
	       (select count(*) from fights f where f.report_id = r.id),
	       (select count(*) from fights f where f.report_id = r.id and f.kill)`
	visibility := `(r.owner_id = $3 or r.visibility = 'public')`
	if verified {
		visibility = `(r.owner_id = $3 or r.visibility = 'public' or r.visibility = 'guild')`
	}
	query := `select ` + columns + ` from reports r
		where r.guild_id = $1 and r.created_at >= $2 and ` + visibility + `
		order by r.created_at desc, r.id desc limit $4`
	args := []any{guildID, since, userID, HomeReportsPerPage}
	if before != nil {
		query = `select ` + columns + ` from reports r
			where r.guild_id = $1 and r.created_at >= $2 and ` + visibility + `
			  and (r.created_at, r.id) < ($4, $5)
			order by r.created_at desc, r.id desc limit $6`
		args = []any{guildID, since, userID, before.CreatedAt, before.ID, HomeReportsPerPage}
	}
	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("guilds: home reports: %w", err)
	}
	defer rows.Close()
	out := []HomeReport{}
	for rows.Next() {
		var h HomeReport
		if err := rows.Scan(&h.ID, &h.Title, &h.CreatedAt, &h.FightCount, &h.KillCount); err != nil {
			return nil, fmt.Errorf("guilds: home reports: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}
```

Change `Home` to thread `userID`/`verified` through and set `Claim`:

```go
func (s *Store) Home(ctx context.Context, guildID, userID int64, verified bool, before *HomeCursor) (HomeView, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return HomeView{}, err
	}
	reports, err := s.HomeReports(ctx, guildID, userID, verified, before)
	if err != nil {
		return HomeView{}, err
	}
	roster, err := s.HomeRoster(ctx, guildID)
	if err != nil {
		return HomeView{}, err
	}
	view := HomeView{
		Guild: GuildIdentity{ID: g.ID, Region: g.Region, Ruleset: g.Ruleset, Name: g.Name},
		Claim: claimState(g, time.Now()),
		Reports: reports, Roster: roster,
	}
	if len(reports) == HomeReportsPerPage {
		last := reports[len(reports)-1]
		view.NextCursor = encodeHomeCursor(last.CreatedAt, last.ID)
	}
	return view, nil
}
```

Change the `home` handler to determine `verified` (via the same `Accounts.GuildRank` every
other officer-power check in this package already uses) and pass both new arguments through:

```go
func (s *Service) home(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	member, err := s.Store.IsMember(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "home", err, "could not load that guild's home just now")
		return
	}
	if !member {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you are not a member of that guild", nil)
		return
	}
	_, verified, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "home", err, "could not load that guild's home just now")
		return
	}
	var before *HomeCursor
	if v := r.URL.Query().Get("cursor"); v != "" {
		c, ok := decodeHomeCursor(v)
		if !ok {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a cursor",
				map[string]string{"cursor": "the next_cursor a previous page returned"})
			return
		}
		before = &c
	}
	view, err := s.Store.Home(r.Context(), guildID, actor.UserID, verified, before)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "home", err, "could not load that guild's home just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}
```

- [ ] **Step 4: Run to verify it passes**

```bash
cd api && go vet ./internal/guilds/... && go test -p 1 ./internal/guilds/... -v 2>&1 | tail -180
```

Expected: PASS, every test in the package, including the three new ones and every
pre-existing home test (`TestHomeRequiresMembership`, `TestHomeRosterHonoursConsent`,
`TestHomeReportsListsOnlyTheTrailingWeek` — that last one's own reports are all owned by the
test's own actor, so the new visibility filter does not change its result: the `owner_id =
$3` clause always matches for that test's seeded reports).

- [ ] **Step 5: Commit**

```bash
cd api && git add internal/guilds/home.go internal/guilds/home_test.go
printf 'fix(guilds): the guild home report list now respects visibility\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-8.txt
git commit -F ../.superpowers/commit-msg-sec-8.txt
```

---

## Task 9: Whole-round verification and final review

**Files:** none modified directly; this task is verification plus whatever a fix wave from
its review touches.

- [ ] **Step 1: Full test run, whole module, sequential**

```bash
cd api && export TEST_DATABASE_URL="postgres://forever:forever@localhost:5434/forever_test?sslmode=disable"
go vet ./... && go test -p 1 -cover ./... 2>&1 | tee /tmp/guild-api-sec-final-test-run.txt | tail -100
```

Expected: every package `ok`, including `internal/guilds`, `internal/addon`,
`internal/auth`, `internal/reports`, `internal/db`, `internal/server`. Any `FAIL` is a
regression this fix round introduced and must be fixed before this task closes.

- [ ] **Step 2: Confirm the migration boundary is still exactly `0018`**

```bash
ls api/internal/db/migrations/ | grep -E '^00(18|19|20)'
```

Expected: only `0018_guild_membership.up.sql` and `.down.sql` — amended in place, per the
coordinator's explicit authorisation, never renumbered.

- [ ] **Step 3: Confirm the file-ownership boundary held**

```bash
cd api && git diff main --stat -- internal/reports/handler.go internal/reports/store.go \
  internal/sims internal/rankings internal/auth/store.go
```

Expected: empty output for every path except (if this fix round's own work required it —
it should not have) nothing at all; this fix round makes zero changes to any of these paths,
including `auth/store.go` (R1a's Battle.net check reads the already-exported
`auth.User.BnetSub` from the handler layer, never modifying `auth` itself).

- [ ] **Step 4: Re-verify each of the coordinator's seven findings directly, with the exact test that proves it fixed**

Run each focused test named below and confirm PASS; this is the evidence for the final
report's "each finding as FIXED with the test that proves it":

```bash
cd api && go test -p 1 ./internal/guilds/... -run 'TestClaimRefusesAnAccountWithNoBattleNetIdentity|TestClaimRateLimitsToOnePerAccountPerThirtyDays|TestClaimRefusesASecondGuildWhileAnotherIsStillHeld|TestContestClaim|TestResolveClaim' -v
go test -p 1 ./internal/guilds/... -run 'TestAutoConfirmClaimIfPendingRefusesTheSameAccountsSecondForgedCharacter' -v
go test -p 1 ./internal/guilds/... -run 'TestAnOfficerCannotRemoveTheGuildMastersCharacter|TestTheAccountHoldingTheClaimCanRemoveAnOfficerButNotAnotherLeaderRow' -v
go test -p 1 ./internal/guilds/... -run 'TestHomeReportsHidesPrivateAndUnlistedFromOthers|TestHomeReportsHidesGuildVisibilityFromAnUnverifiedMember' -v
go test -p 1 ./internal/db/... -run 'TestGuildsAreCaseInsensitiveByRegionAndRuleset' -v
go test -p 1 ./internal/addon/... -run 'TestPutExportsResolvesGuildsCaseInsensitively|TestPutExportsRejectsAnOutOfRangeRankIndexAsANoOpNotAnAbort|TestPutExportsContinuesTheBatchPastOneBadCharacter' -v
go test -p 1 ./internal/guilds/... -run 'TestRecomputeMembershipSerialisesConcurrentCallersForTheSameGuild' -v
go test -p 1 ./internal/addon/... -run 'TestConcurrentPutExportsAndApproveDoNotLoseAWrite' -v
```

Expected: PASS on every command.

- [ ] **Step 5: Dispatch the whole-round review**

Use `superpowers:requesting-code-review`. Give the reviewer, verbatim:
- This plan's Global Constraints section.
- The coordinator's original seven-finding list (R1-R7) from their dispatch message, and
  the security reviewer's original finding text (both are reproduced in this plan's Task
  descriptions above, task by task).
- The two "Amendment, 2026-09-21" blocks in the spec (`docs/superpowers/specs/2026-09-21-guild-membership-design.md`
  §2.4 and §3.3).
- Instruction: for each of the seven findings, confirm (a) the fix is actually present in
  the diff, (b) the test that claims to prove it actually exercises the vulnerable path
  (not a vacuous assertion), and (c) no other path in the same file re-introduces the same
  hole (e.g., check every place `verified_at` is ever set now also sets `verified_by`; check
  every officer-power route now checks `contested`, not just the four named).
- Instruction: specifically re-attempt the CRITICAL exploit sketch from the original
  review — forge a `rankIndex == 0` export naming a real unclaimed guild, `POST
  /v1/guilds/{id}/claim` — and confirm it now requires a Battle.net identity, is
  rate-limited, and is contestable, rather than merely being "harder."

- [ ] **Step 6: One fix wave**

Address every CRITICAL and HIGH finding from Step 5 directly (dispatch a fresh
implementer + reviewer pair per this lane's house rules for anything non-trivial). Re-run
Step 1 after any fix.

- [ ] **Step 7: Write the final report**

Per the coordinator's request: branch head sha, test results (Step 1's full run), each of
the seven findings marked FIXED with the exact test name that proves it (Step 4's list),
and the new/changed request/response shapes for the web lane (already listed in this plan's
Global Constraints — copy verbatim: `ClaimStateView`, the `claim` field on
`SettingsView`/`HomeView`, the two new endpoints and their bodies, the new `409
claim_contested` code, and the fixed `SettingsView.ClaimedBy`/`ClaimPending`).

- [ ] **Step 8: Final commit (only if Step 6 changed anything)**

```bash
cd api && git add -A
printf 'fix: address whole-round review findings for the guild API security fix round\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-sec-9.txt
git commit -F ../.superpowers/commit-msg-sec-9.txt
```
