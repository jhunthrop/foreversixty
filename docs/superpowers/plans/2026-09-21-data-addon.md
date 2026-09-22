# Data Addon Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the nightly job (`api/cmd/api ... data-addon`) that reads public rating and guild data out of Postgres and writes the Forever Sixty Data addon's `Data.lua`, plus the GitHub Actions workflow that packages and publishes it.

**Architecture:** A new, self-contained Go package `api/internal/dataaddon` does everything: read the database (`store.go`), fold rows into per-character and per-guild aggregates (`aggregate.go`), serialize them to Lua deterministically (`write.go`, with a region-split fallback for a file over 4 MiB), and publish the result to a Cloud Storage bucket (`upload.go`). `job.go` wires the four together behind one `Run` entry point that `api/cmd/api/main.go` dispatches on the argument `data-addon`, exactly the way `rating.BackfillJobCommand` is already dispatched. A new GitHub Actions workflow (`addon-data-release.yml`) downloads the bucket's nightly output and republishes it through the same BigWigsMods packager `addon-release.yml` already uses.

**Tech Stack:** Go 1.25 (api module), `pgx/v5` against Postgres, `google.golang.org/api/storage/v1` (already an indirect-turned-direct dependency of `api`, via `google.golang.org/api` — no new module needed) for the Cloud Storage write, Lua 5.1-compatible output text, `busted` for the addon-side spec.

**Spec:** No separate spec document exists; the dispatch that started this lane (recorded in full below, under "Binding design (the dispatch)") is the binding design. This plan is this lane's spec coverage of it.

## Binding design (the dispatch)

> Your lane: "data-addon" — the nightly job that writes and publishes the Forever Sixty Data addon.
>
> **What exists.** `addon/ForeverSixtyData/` is the data addon: a TOC with the CurseForge id 1706707 and Wago id ANz7R564, an empty `Data.lua` showing the shape, and a README. `addon/ForeverSixty/Ratings.lua` is the READER in the main addon, and is the contract you must satisfy exactly: read it first. It expects a global `ForeverSixtyData` with `format = 1`, `generated` as an ISO 8601 UTC timestamp `YYYY-MM-DDTHH:MM:SSZ`, `build` (the data build), `characters` keyed `"<region>:<realm>:<name>"` lowercased with spaces to dashes (the reader builds the key from the client's region name such as "US", `GetRealmName()` and the character name, so the job must produce keys the same way from the site's own character data: work out from `api/internal/character/character.go` and the `rating_scores.player_key` format what the realm and name are, and write down in the plan exactly how the two key formats are reconciled; if the site's key carries a ruleset rather than a realm, the job must map back to the realm the addon will see, and you must say how), each row `{ rating, output, survival, mechanics, utility, preparation, activity, fights }` as numbers; and `guilds` keyed the same way with `{ name, progress, nights, roster, members = { "<name>", ... } }` where members are lowercased character names on the guild's realm. The reader treats anything missing as "no data", so the job may omit what it cannot compute, but must never write a wrong number.
>
> **The job.** A new `api/cmd/api` subcommand `data-addon` (dispatch it beside `rating-backfill` in `main.go`, a small contiguous change) in a new package `api/internal/dataaddon/`:
> 1. Reads, from the database, every character with at least one rating row in `rating_scores` joined to `reports` with `visibility = 'public'` (the same rule the rating read endpoints use: read `api/internal/rating/store.go`), excluding anonymized owners the way the character rating endpoint does. A character's `rating` is the mean of `overall` over their public rated fights in the last 90 days, rounded to an integer; each component is the mean of that component's score over the same fights where the component was not excluded (read the `components` jsonb shape from `api/internal/rating/dto.go` and the engine's `Card`); `fights` is the count. State this aggregation in the plan and in the `/ratings` page's wording if it differs from what the page says (you may not edit the web; report the discrepancy instead).
> 2. Reads guilds from the guild tables (`guilds`, `guild_members`, `guild_characters`, migration 0018): only VERIFIED members count toward the roster and the members list; `progress` is the site's existing progression summary if one is queryable (look at `api/internal/rankings/guilds.go` for what the public guild page shows) else omit; `nights` is distinct raid nights in the last 90 days from reports attached to the guild with `visibility = 'public'`.
> 3. Writes `Data.lua` deterministically: sorted keys, fixed number formatting, one row per line, a header comment with the generation time and the build, valid Lua (escape strings with `%q` semantics; names can contain apostrophes and non-ASCII). Size: write region-split files only if the single file would exceed 4 MB (then `Data.lua` becomes a loader of `Data-us.lua`, `Data-eu.lua`... listed in the TOC; the reader on the main addon needs no change because the global is the same); otherwise one file. Write a size line to the log either way.
> 4. Never fails a run on one bad row: a row it cannot format is skipped with a logged reason and counted.
> 5. Output goes to a path given by `--out` (default the repository's `addon/ForeverSixtyData/Data.lua` when run locally; in production the job writes to a path the workflow collects). Design the production path: the job runs on Cloud Run with no repository checkout, so the simplest contract is: the job writes the file to a Cloud Storage bucket object (`gs://<bucket>/data-addon/Data.lua`, bucket name from an env var `DATA_ADDON_BUCKET`, skipped with a log line when unset), and a GitHub Actions workflow `addon-data-release.yml`, scheduled nightly, downloads it, drops it into `addon/ForeverSixtyData/`, and runs the existing packager the way `addon-release.yml` does for the main addon, tagging the release with the date (`data-v2026.09.22`). Read `.github/workflows/addon-release.yml` carefully and reuse its packaging steps (its "give the packager its own checkout" trick is needed here too); the two secrets `CF_API_KEY` and `WAGO_API_TOKEN` already exist; the workflow needs a Google Cloud credential to read the bucket: use Workload Identity Federation if `.github/workflows/api.yml` already authenticates that way (read it), and name exactly what the owner must create (the bucket, the job, the scheduler, any IAM binding) in `api/README.md` next to the other job recipes. Do not create any cloud resource yourself.
> 6. Tests: a golden test that a fixed database fixture produces a byte-exact `Data.lua`; a test that the written file is valid Lua that defines the global (run it through the `lua` interpreter if present in the test environment, else parse it with a minimal check and say so); a test that the reader in `addon/ForeverSixty/Ratings.lua` accepts the job's output (drive busted from Go with `os/exec` if `busted` is on the path, else add a busted spec under `addon/tests/` that loads a checked-in sample the Go golden test also asserts against, so the two stay in lockstep); key-format tests for names with spaces, apostrophes and non-ASCII; the 90-day window; the anonymized exclusion; the public-only rule.
>
> Rules: own throwaway test database with `-p 1`; failing test first; ownership is `api/internal/dataaddon`, the `main.go` dispatch, `.github/workflows/addon-data-release.yml`, `api/README.md` additions, `addon/ForeverSixtyData/` (the README and the TOC's file list if region splitting needs it; the checked-in `Data.lua` stays the empty shape), and a busted spec under `addon/tests/`; do not edit `addon/ForeverSixty/` (if the reader needs a change, report it), `web/`, `logs/`, `sim/`, or other `api/internal` packages. Never end a turn with background work outstanding. Final review: code plus a security pass (the job reads private data and must publish only public-derived numbers; check the anonymized and visibility rules again there). Report the branch head, tests, the exact production setup steps for the owner, and the aggregation rule as implemented.

## Key-format reconciliation (dispatch item 1, required write-up)

`addon/ForeverSixty/Ratings.lua` builds its lookup key client-side as `slug(region) .. ":" .. slug(realm) .. ":" .. slug(name)`, where `region` comes from `GetCurrentRegion()` (`Export.REGION_NAMES`, e.g. `"US"`) and `realm` from `GetRealmName()`. `slug` is `tostring(x):lower():gsub("%s+", "-")`.

The site's own character identity has no realm at all. `api/internal/character/character.go`'s package doc says so directly: *"Forever has no realms: the Deep Dive panel replaced them with four rulesets per region. A character key is `<region>/<ruleset>/<name-slug>`."* `rating_scores.player_key` is written in exactly that format (`character.Key`/`character.KeyFromUnit`, used by `rating.Store.RateFight`).

The reconciliation: **this job writes the addon's `<realm>` position with the site's `ruleset`.** Evidence this is the intended mapping, not a guess:

- The companion addon itself already does this. `addon/ForeverSixty/Export.lua:263-267` writes `ForeverSixtyDB.characters[region .. "/" .. realm .. "/" .. name] = { ..., realm = realm, ruleset = realm, ... }` — it reads `GetRealmName()` once and uses that one string as both `realm` and `ruleset`, with the comment: *"`ruleset` is the realm name until the beta shows an API that states Forever's own ruleset (spike check 12)."*
- Since Forever has no realm concept server-side, `GetRealmName()` is the only per-account string a Classic client exposes that could carry it, so the whole codebase's working assumption is that a live client's realm name *is* (or reads as) one of the four ruleset names (`normal`, `pvp`, `rp`, `hardcore`), until the Sept 17 beta proves otherwise.

**Risk, stated plainly:** if the beta shows `GetRealmName()` returning something else (a real fantasy realm name distinct from the ruleset), every key this job writes will miss every lookup the addon performs, silently. The blast radius is small and matches the reader's own design: `Ratings.status()`/`forCharacter`/`forGuild` all treat a miss as "no data," never an error, so a wrong mapping degrades the addon to "shows nothing" rather than showing something false. This is the same hedge `Export.lua`'s own comment carries, not a new risk this job introduces. If the beta settles a different mapping, only `characterKey`/`guildKey` in `api/internal/dataaddon/keys.go` (Task 2) need to change.

The existing addon test fixture (`addon/tests/ratings_spec.lua`) uses `realm = "Ashbringer"` as a mock value; that is `Ratings.lua`'s own spec testing that the reader builds *some* string into a key correctly, not a claim about what Forever's real realm strings will be — `Ratings.lua` is agnostic to what the string means. This job's own new golden fixture (Task 10) uses ruleset-shaped values (`normal`, `pvp`) instead, matching what this job actually writes.

## Aggregation rule (dispatch item 1, and the `/ratings` page discrepancy)

**As implemented:**
- **Character `rating`:** `round(mean(overall))` over every `rating_scores` row for that `player_key` where (a) the joined `reports.visibility = 'public'`, (b) `rating_scores.fought_at >= now() - 90 days`, and (c) the player's owning account (via `characters.user_id` → `users.anonymize`) is not anonymized. A `player_key` with no `characters` row at all is not anonymized (nothing to hide), matching `rating.Store.anonymized`'s own rule exactly.
- **Each component** (`output`, `survival`, `mechanics`, `utility`, `preparation`, `activity`): `round(mean(score))` over the same fight set, counting only the fights where that component's stored `excluded` flag is `false`. A component with zero non-excluded fights in the window is **omitted from the row entirely** (dispatch: "may omit what it cannot compute"), never written as zero.
- **`fights`:** the count of fights the rating itself rested on (before per-component exclusion) — i.e. every public, in-window, non-anonymized fight for that character.
- **Guild `progress`:** `"<encounters killed at least once>/<encounters attempted>"`, counted from `fights` joined through `reports` where `reports.visibility <> 'private'` — the exact rule `rankings.Store.Guild`'s own progression query uses (`api/internal/rankings/guilds.go`), duplicated here rather than imported (see Task 7's comment for why). Omitted when the guild has zero attempted encounters.
- **Guild `nights`:** distinct UTC calendar dates (`date(reports.created_at)`) among that guild's `visibility = 'public'` reports with `created_at >= now() - 90 days`. The dispatch names "distinct raid nights" without defining one; this is this job's own definition, stated here as a ruling — a report that starts before and ends after UTC midnight would double-count, an accepted low-stakes approximation.
- **Guild `roster` / `members`:** the count and lowercased name-slugs of `guild_characters` rows for that guild with `verified_at is not null`, sorted alphabetically for determinism. A guild with zero verified members gets no row at all.
- Every number in the file is rounded to the nearest integer (Go's `math.Round`, ties away from zero) — the dispatch only says this explicitly for `rating`, but `Ratings.lua`'s own `L.ratingsTooltip = "Forever Sixty rating %d, over %d fights"` requires `rating` and `fights` to be whole numbers for Lua's `string.format("%d", ...)` to accept them at all, and rounding every field the same way keeps `Data.lua`'s formatting uniform and the golden test unambiguous (no fractional ties to adjudicate).

**Discrepancy with `web/src/pages/ratings.astro` (reported, not fixed — web is out of this lane's ownership):** the `/ratings` page and `web/src/components/CharacterRatingPanel.svelte` describe and display a character's rating as **the single most recent fight's card** (`data.latest`, rendered `Math.round(data.latest.overall)`), plus a trend chart of individual fights — there is no rolling/mean rating anywhere in the web surface today. This job's `rating` is a **90-day mean across every public rated fight**, which is a different number from what `CharacterRatingPanel.svelte` shows for the same character on the same day. Both are faithful to their own inputs (the API's `characterRatingDTO` has no aggregate field to reuse), but a player could see two different "rating" numbers for themselves — one in the addon, one on the site — until a future lane either adds a real aggregate to the API or changes the addon's contract to want `latest` instead of a mean. Flagging this for the web/rating-api lanes rather than changing either surface myself.

## Global Constraints

- Go 1.25.11 (this machine's default `go` already matches; no `GOTOOLCHAIN` override needed — verified `go version` in the worktree).
- `go vet ./... && go test ./...` (add `-p 1` for this package's own DB-backed tests) from `api/` before every commit.
- Own throwaway test database (`dataaddon_test`, a separate database name on the existing `docker-compose.test.yml` Postgres instance, port 5434 — not `forever_test`, which other lanes' tests truncate concurrently).
- Failing test first, every task.
- File ownership: `api/internal/dataaddon/` (new package, full ownership), `api/cmd/api/main.go` (only the dispatch `case` and its one handler function — a small, contiguous change), `.github/workflows/addon-data-release.yml` (new file), `api/README.md` (additions only, appended near the other job recipes), `addon/ForeverSixtyData/` (README and, only if region-splitting is exercised in production, the TOC's file list — the checked-in `Data.lua` keeps its current empty-shape content), and a new busted spec under `addon/tests/` plus its fixture under `addon/tests/fixtures/`.
- Never edit `addon/ForeverSixty/`, `web/`, `logs/`, `sim/`, or any other `api/internal/*` package (including `api/internal/config` — see the ruling in Task 11 for how `DATA_ADDON_BUCKET` is read without touching it).
- Never connect to the production database, never print a secret, never deploy.
- Commits: message via `printf` to a file under this worktree's `.superpowers/`, then `git commit -F <file>` as its own Bash command (no heredoc combined with commit, no `-n`, no "no-verify"). End every message with:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`
- Functions under 50 lines, files under 800, no magic numbers, no dead code, pure functions with explicit inputs, immutable updates, errors wrapped with context, table-driven tests where the shape fits.
- Model budget: every implementer/reviewer subagent runs on `sonnet`; `haiku` only for a purely mechanical single-file fix; never `opus`. One implementer at a time.

## File Structure

```
api/internal/dataaddon/
  doc.go             package doc, JobCommand, ratingWindow, maxSingleFileBytes constants
  keys.go            slug(), characterKey(), guildKey()
  aggregate.go        componentNames, fightScore/componentScore, decodeComponents, aggregateCharacter, roundHalfUp
  guild_aggregate.go  buildGuildRow, memberNameOf
  write.go            Data, characterRow/guildRow, Render, render, renderSplit, sortedKeys
  upload.go           Uploader interface, GCS client, Fake
  store.go            Store and every DB read
  job.go              Deps, Result, Run, Logger, manifestOf, splitCharacterKey
  *_test.go           one per file above, plus db_test.go (testURL/testPool) and a golden_test.go

api/cmd/api/main.go    + case dataaddon.JobCommand, + runDataAddon
api/README.md          + "The data-addon job" section

.github/workflows/addon-data-release.yml   new

addon/ForeverSixtyData/README.md   + one paragraph naming the nightly job
addon/tests/data_addon_spec.lua            new busted spec
addon/tests/fixtures/data_addon_sample.lua new fixture, byte-identical to the Go golden test's expected output
```

---

### Task 1: Package skeleton, throwaway test database, and key-format module

**Files:**
- Create: `api/internal/dataaddon/doc.go`
- Create: `api/internal/dataaddon/db_test.go`
- Create: `api/internal/dataaddon/keys.go`
- Create: `api/internal/dataaddon/keys_test.go`

**Interfaces:**
- Produces: `dataaddon.JobCommand` (string constant `"data-addon"`), `dataaddon.ratingWindow` (`time.Duration`, 90 days), `dataaddon.maxSingleFileBytes` (`int`, `4 << 20`), unexported `slug(s string) string`, `characterKey(region, ruleset, name string) string`, `guildKey(region, ruleset, name string) string`, and the test helpers `testURL(t *testing.T) string` / `testPool(t *testing.T) *pgxpool.Pool`.

- [ ] **Step 1: Start the throwaway Postgres and create this lane's own database**

Run (from the repository root):
```bash
cd api
docker compose -f docker-compose.test.yml up -d
PGPASSWORD=forever createdb -h localhost -p 5434 -U forever dataaddon_test
```
Expected: the `createdb` command exits 0 (or, if it already exists from a prior run of this task, prints `database "dataaddon_test" already exists` — either is fine to continue).

- [ ] **Step 2: Write `db_test.go` (this package's own copy of the shared testURL/testPool pattern)**

```go
// api/internal/dataaddon/db_test.go
package dataaddon

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// testURL is this package's own copy of the shape every DB-backed package in
// this repo declares (compare api/internal/rating/db_test.go):
// TEST_DATABASE_URL must point at dataaddon_test, not the shared
// forever_test other lanes' tests truncate concurrently, per this lane's
// dispatch ("own throwaway test database").
func testURL(t *testing.T) string {
	t.Helper()
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL not set; see docs/superpowers/plans/2026-09-21-data-addon.md Task 1 Step 1")
	}
	return u
}

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := testURL(t)
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
```

- [ ] **Step 3: Write the failing test for `slug`/`characterKey`/`guildKey`**

```go
// api/internal/dataaddon/keys_test.go
package dataaddon

import "testing"

func TestSlugLowercasesAndCollapsesWhitespaceToOneHyphen(t *testing.T) {
	cases := map[string]string{
		"Thoradin":        "thoradin",
		"Iron  Vanguard":  "iron-vanguard", // a doubled space still collapses to one hyphen
		"O'Malley":        "o'malley",      // an apostrophe is not whitespace: untouched
		"Mörk":            "mörk",          // non-ASCII passes through unescaped
	}
	for in, want := range cases {
		if got := slug(in); got != want {
			t.Errorf("slug(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCharacterKeyMatchesTheAddonReadersOwnFormat(t *testing.T) {
	// addon/ForeverSixty/Ratings.lua: characterKey(region, realm, name) =
	// slug(region) .. ":" .. slug(realm) .. ":" .. slug(name); this job's
	// realm position is the site's ruleset (see the plan's "Key-format
	// reconciliation" section).
	got := characterKey("US", "normal", "Thoradin")
	if want := "us:normal:thoradin"; got != want {
		t.Errorf("characterKey = %q, want %q", got, want)
	}
	got = characterKey("EU", "pvp", "Mörk")
	if want := "eu:pvp:mörk"; got != want {
		t.Errorf("characterKey = %q, want %q", got, want)
	}
}

func TestGuildKeyCollapsesTheGuildNamesSpaces(t *testing.T) {
	got := guildKey("US", "normal", "Iron Vanguard")
	if want := "us:normal:iron-vanguard"; got != want {
		t.Errorf("guildKey = %q, want %q", got, want)
	}
}
```

- [ ] **Step 4: Run the tests to verify they fail**

Run: `cd api && go test ./internal/dataaddon/... -run TestSlug -v`
Expected: FAIL — `slug` undefined (the package does not compile yet: no non-test file exists).

- [ ] **Step 5: Write `doc.go` and `keys.go`**

```go
// api/internal/dataaddon/doc.go
// Package dataaddon is the nightly job that writes the Forever Sixty Data
// addon's ForeverSixtyData global: every character with at least one public
// rated fight in the last 90 days, and every guild with at least one
// verified member, serialized to Data.lua (or Data-<region>.lua, when the
// combined file would exceed 4 MiB) for the addon-data-release.yml workflow
// to package and publish. addon/ForeverSixty/Ratings.lua is the reader this
// package's output must satisfy exactly; see
// docs/superpowers/plans/2026-09-21-data-addon.md for the key-format
// reconciliation and aggregation rule this package implements.
package dataaddon

import "time"

// JobCommand is this lane's Cloud Run job name, dispatched the same way
// rating.BackfillJobCommand already is (api/cmd/api/main.go's os.Args[1]
// switch).
const JobCommand = "data-addon"

// ratingWindow is "the last 90 days" the dispatch's aggregation rule names,
// for both a character's rating and a guild's night count.
const ratingWindow = 90 * 24 * time.Hour

// maxSingleFileBytes is the dispatch's own split threshold: a rendered
// Data.lua at or under this size ships as one file; larger, it ships as
// one file per region instead (see write.go's Render).
const maxSingleFileBytes = 4 << 20 // 4 MiB
```

```go
// api/internal/dataaddon/keys.go
package dataaddon

import (
	"regexp"
	"strings"
)

// whitespaceRun matches the same class the addon's own slug function does
// (Ratings.lua: tostring(text):lower():gsub("%s+", "-")). Lua's %s class is
// ASCII whitespace (space, tab, newline, CR, FF, VT); Go's \s matches it
// identically for the ASCII range every real character or guild name is
// drawn from.
var whitespaceRun = regexp.MustCompile(`\s+`)

// slug lowercases s and collapses every run of whitespace into a single
// hyphen -- byte-for-byte what the addon's own Ratings.lua computes
// client-side from GetRealmName(), a character name, and a guild name.
//
// This cannot reuse api/internal/character.Slug: that helper trims the
// string first and replaces only the literal space character one-for-one,
// which is a different function for a name with a tab or a doubled space
// (neither occurs in a real character name, but a guild name can have a
// doubled space, and the addon's own gsub is the contract this job has to
// match exactly, not approximate).
func slug(s string) string {
	return whitespaceRun.ReplaceAllString(strings.ToLower(s), "-")
}

// characterKey builds the addon's lookup key for a character:
// "<region>:<ruleset>:<name-slug>". The dispatch's "realm" position is
// filled with the site's ruleset -- see the plan's "Key-format
// reconciliation" section for why.
func characterKey(region, ruleset, name string) string {
	return slug(region) + ":" + slug(ruleset) + ":" + slug(name)
}

// guildKey is characterKey's twin for a guild: the same three parts, the
// guild's own name in the third.
func guildKey(region, ruleset, name string) string {
	return characterKey(region, ruleset, name)
}
```

- [ ] **Step 6: Run the tests to verify they pass**

Run: `cd api && go test ./internal/dataaddon/... -v`
Expected: PASS (every `Test...` in `keys_test.go`).

- [ ] **Step 7: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/doc.go internal/dataaddon/db_test.go internal/dataaddon/keys.go internal/dataaddon/keys_test.go
printf 'feat(api): data-addon package skeleton and the addon key format\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-1.txt
git commit -F ../.superpowers/commit-msg-1.txt
```

---

### Task 2: Character aggregation

**Files:**
- Create: `api/internal/dataaddon/aggregate.go`
- Create: `api/internal/dataaddon/aggregate_test.go`

**Interfaces:**
- Consumes: nothing from earlier tasks beyond the package existing.
- Produces: `componentNames []string` (`{"output","survival","mechanics","utility","preparation","activity"}`, `Ratings.lua`'s own order), `componentScore{Name string; Score *float64}`, `fightScore{Overall float64; Components []componentScore}`, `characterRow{Rating int; Components map[string]int; Fights int}`, `decodeComponents(raw json.RawMessage) []componentScore`, `aggregateCharacter(fights []fightScore) (characterRow, bool)`, `roundHalfUp(v float64) int`. Task 6 (job.go) and Task 10 (golden test) call `decodeComponents` and `aggregateCharacter` directly.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/dataaddon/aggregate_test.go
package dataaddon

import "testing"

func score(v float64) *float64 { return &v }

func TestAggregateCharacterAveragesOverallAndEachNonExcludedComponent(t *testing.T) {
	fights := []fightScore{
		{Overall: 80, Components: []componentScore{
			{Name: "output", Score: score(90)}, {Name: "survival", Score: score(70)},
			{Name: "mechanics", Score: score(85)}, {Name: "utility", Score: score(60)},
			{Name: "preparation", Score: score(95)}, {Name: "activity", Score: nil}, // excluded this fight
		}},
		{Overall: 82, Components: []componentScore{
			{Name: "output", Score: score(92)}, {Name: "survival", Score: score(74)},
			{Name: "mechanics", Score: score(83)}, {Name: "utility", Score: score(64)},
			{Name: "preparation", Score: score(93)}, {Name: "activity", Score: score(91)},
		}},
	}
	row, ok := aggregateCharacter(fights)
	if !ok {
		t.Fatal("aggregateCharacter reported no data for two fights")
	}
	if row.Rating != 81 {
		t.Errorf("rating = %d, want 81", row.Rating)
	}
	if row.Fights != 2 {
		t.Errorf("fights = %d, want 2", row.Fights)
	}
	want := map[string]int{"output": 91, "survival": 72, "mechanics": 84, "utility": 62, "preparation": 94, "activity": 91}
	for name, v := range want {
		if row.Components[name] != v {
			t.Errorf("component %s = %d, want %d", name, row.Components[name], v)
		}
	}
}

func TestAggregateCharacterOmitsAComponentWithNoNonExcludedFights(t *testing.T) {
	fights := []fightScore{
		{Overall: 88, Components: []componentScore{
			{Name: "output", Score: score(70)}, {Name: "survival", Score: nil},
			{Name: "mechanics", Score: score(70)}, {Name: "utility", Score: nil},
			{Name: "preparation", Score: nil}, {Name: "activity", Score: nil},
		}},
	}
	row, ok := aggregateCharacter(fights)
	if !ok {
		t.Fatal("aggregateCharacter reported no data for one fight")
	}
	if len(row.Components) != 2 {
		t.Fatalf("components = %v, want exactly output and mechanics", row.Components)
	}
	if row.Components["output"] != 70 || row.Components["mechanics"] != 70 {
		t.Errorf("components = %v", row.Components)
	}
	if _, present := row.Components["survival"]; present {
		t.Error("survival should be omitted, not written as any number")
	}
}

func TestAggregateCharacterReportsNoDataForZeroFights(t *testing.T) {
	_, ok := aggregateCharacter(nil)
	if ok {
		t.Error("aggregateCharacter should report ok=false for zero fights")
	}
}

func TestDecodeComponentsReadsExcludedAsNilScore(t *testing.T) {
	raw := []byte(`[{"name":"output","score":70,"excluded":false},{"name":"survival","score":null,"excluded":true,"reason":"wipe"}]`)
	got := decodeComponents(raw)
	if len(got) != 2 {
		t.Fatalf("decoded %d components, want 2", len(got))
	}
	if got[0].Name != "output" || got[0].Score == nil || *got[0].Score != 70 {
		t.Errorf("output = %+v", got[0])
	}
	if got[1].Name != "survival" || got[1].Score != nil {
		t.Errorf("survival = %+v, want a nil score", got[1])
	}
}

func TestDecodeComponentsReturnsNilForUnreadableJSON(t *testing.T) {
	if got := decodeComponents([]byte(`not json`)); got != nil {
		t.Errorf("decodeComponents(garbage) = %v, want nil", got)
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run 'TestAggregate|TestDecode' -v`
Expected: FAIL — undefined: `fightScore`, `componentScore`, `aggregateCharacter`, `decodeComponents`.

- [ ] **Step 3: Write `aggregate.go`**

```go
// api/internal/dataaddon/aggregate.go
package dataaddon

import (
	"encoding/json"
	"math"
)

// componentNames is the addon's six component keys, in Ratings.lua's own
// COMPONENTS order -- the order this package writes them in Data.lua, so a
// diff between two nights' files stays minimal.
var componentNames = []string{"output", "survival", "mechanics", "utility", "preparation", "activity"}

// componentScore is one fight's contribution to one component: nil when
// the engine excluded it for that fight (a stored componentDTO's score is
// JSON null exactly when excluded is true -- see api/internal/rating/cards.go).
type componentScore struct {
	Name  string
	Score *float64
}

// fightScore is one public, in-window, non-anonymized rated fight read for
// one character.
type fightScore struct {
	Overall    float64
	Components []componentScore
}

// characterRow is one character's line in Data.lua: the rounded rating,
// whichever components had at least one non-excluded fight to average
// (absent from the map for the rest -- "may omit what it cannot compute"),
// and the fight count the rating rests on.
type characterRow struct {
	Rating     int
	Components map[string]int
	Fights     int
}

// decodeComponents reads one fight's stored components jsonb
// (api/internal/rating/cards.go's componentDTO array) into the scores this
// package aggregates over. A row whose components column does not decode
// as the expected shape is not a reason to drop the whole fight -- its
// Overall still counts toward the rating -- so this returns nil rather
// than an error, and every component is simply excluded for that one
// fight: "never write a wrong number" outranks "use every fight for every
// field."
func decodeComponents(raw json.RawMessage) []componentScore {
	var decoded []struct {
		Name     string   `json:"name"`
		Score    *float64 `json:"score"`
		Excluded bool     `json:"excluded"`
	}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil
	}
	out := make([]componentScore, 0, len(decoded))
	for _, d := range decoded {
		if d.Excluded {
			out = append(out, componentScore{Name: d.Name, Score: nil})
			continue
		}
		out = append(out, componentScore{Name: d.Name, Score: d.Score})
	}
	return out
}

// aggregateCharacter folds one character's fights into its Data.lua row.
// ok is false for zero fights: a character with no public rated fights in
// the window has nothing to say, and gets no row at all rather than a
// zeroed one.
func aggregateCharacter(fights []fightScore) (row characterRow, ok bool) {
	if len(fights) == 0 {
		return characterRow{}, false
	}
	var overallSum float64
	sums := make(map[string]float64, len(componentNames))
	counts := make(map[string]int, len(componentNames))
	for _, f := range fights {
		overallSum += f.Overall
		for _, c := range f.Components {
			if c.Score == nil {
				continue
			}
			sums[c.Name] += *c.Score
			counts[c.Name]++
		}
	}
	components := make(map[string]int, len(componentNames))
	for _, name := range componentNames {
		if n := counts[name]; n > 0 {
			components[name] = roundHalfUp(sums[name] / float64(n))
		}
	}
	return characterRow{
		Rating:     roundHalfUp(overallSum / float64(len(fights))),
		Components: components,
		Fights:     len(fights),
	}, true
}

// roundHalfUp rounds to the nearest integer, ties away from zero -- Go's
// math.Round's own rule. Every number this package writes (rating and
// every component) takes this same rounding, so Data.lua's formatting
// stays uniform and its golden test has no fractional tie to adjudicate.
func roundHalfUp(v float64) int {
	return int(math.Round(v))
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -v`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/aggregate.go internal/dataaddon/aggregate_test.go
printf 'feat(api): data-addon character aggregation (90-day mean, per-component exclusion)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-2.txt
git commit -F ../.superpowers/commit-msg-2.txt
```

---

### Task 3: Guild row assembly

**Files:**
- Create: `api/internal/dataaddon/guild_aggregate.go`
- Create: `api/internal/dataaddon/guild_aggregate_test.go`

**Interfaces:**
- Consumes: nothing new.
- Produces: `guildRow{Name, Progress string; Nights, Roster int; Members []string}`, `guildIdentity{ID int64; Region, Ruleset, Name string}` (also used by Task 7's store), `buildGuildRow(g guildIdentity, memberKeys []string, nights, killed, total int) guildRow`, `memberNameOf(characterKey string) string`.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/dataaddon/guild_aggregate_test.go
package dataaddon

import "testing"

func TestMemberNameOfReadsTheNameSegmentOfACharacterKey(t *testing.T) {
	if got := memberNameOf("us/normal/thoradin"); got != "thoradin" {
		t.Errorf("memberNameOf = %q, want thoradin", got)
	}
	if got := memberNameOf("not-a-character-key"); got != "not-a-character-key" {
		t.Errorf("memberNameOf on an unparseable key should return it unchanged, got %q", got)
	}
}

func TestBuildGuildRowSortsMembersAndComputesProgress(t *testing.T) {
	g := guildIdentity{ID: 1, Region: "us", Ruleset: "normal", Name: "Iron Vanguard"}
	row := buildGuildRow(g, []string{"us/normal/thoradin", "us/normal/o'malley"}, 2, 2, 3)
	if row.Name != "Iron Vanguard" {
		t.Errorf("name = %q", row.Name)
	}
	if row.Progress != "2/3" {
		t.Errorf("progress = %q, want 2/3", row.Progress)
	}
	if row.Nights != 2 {
		t.Errorf("nights = %d, want 2", row.Nights)
	}
	if row.Roster != 2 {
		t.Errorf("roster = %d, want 2", row.Roster)
	}
	want := []string{"o'malley", "thoradin"}
	for i, name := range want {
		if row.Members[i] != name {
			t.Errorf("members[%d] = %q, want %q", i, row.Members[i], name)
		}
	}
}

func TestBuildGuildRowOmitsProgressWithNoAttemptedEncounters(t *testing.T) {
	g := guildIdentity{ID: 1, Region: "us", Ruleset: "normal", Name: "Iron Vanguard"}
	row := buildGuildRow(g, []string{"us/normal/thoradin"}, 0, 0, 0)
	if row.Progress != "" {
		t.Errorf("progress = %q, want empty (omitted)", row.Progress)
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run 'TestMemberNameOf|TestBuildGuildRow' -v`
Expected: FAIL — undefined: `guildIdentity`, `buildGuildRow`, `memberNameOf`.

- [ ] **Step 3: Write `guild_aggregate.go`**

```go
// api/internal/dataaddon/guild_aggregate.go
package dataaddon

import (
	"fmt"
	"sort"
	"strings"
)

// guildRow is one guild's line in Data.lua.
type guildRow struct {
	Name     string
	Progress string // "" when there is nothing to report (omitted from Data.lua)
	Nights   int
	Roster   int
	Members  []string // lowercased character-name slugs, sorted
}

// guildIdentity names one guild with at least one verified member -- the
// only kind this job publishes a row for.
type guildIdentity struct {
	ID              int64
	Region, Ruleset, Name string
}

// memberNameOf reads the name segment back out of a character key
// ("us/normal/thoradin" -> "thoradin") -- guild_characters.character_key
// is written in api/internal/character.Key's own format
// (<region>/<ruleset>/<name-slug>), and this is that format's inverse for
// the one segment this package needs. An unparseable key (should not
// happen: character_key is not-null and every write path uses
// character.Key) is returned unchanged rather than panicking, so a bad row
// degrades to an odd-looking member name instead of stopping the run.
func memberNameOf(key string) string {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 {
		return key
	}
	return parts[2]
}

// buildGuildRow assembles one guild's Data.lua row. memberKeys are that
// guild's verified guild_characters.character_key values; killed/total
// are attempted-encounter counts (see store.go's progressionByGuild).
func buildGuildRow(g guildIdentity, memberKeys []string, nights, killed, total int) guildRow {
	members := make([]string, 0, len(memberKeys))
	for _, key := range memberKeys {
		members = append(members, slug(memberNameOf(key)))
	}
	sort.Strings(members)
	progress := ""
	if total > 0 {
		progress = fmt.Sprintf("%d/%d", killed, total)
	}
	return guildRow{
		Name: g.Name, Progress: progress, Nights: nights,
		Roster: len(members), Members: members,
	}
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -v`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/guild_aggregate.go internal/dataaddon/guild_aggregate_test.go
printf 'feat(api): data-addon guild row assembly (progress, nights, verified roster)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-3.txt
git commit -F ../.superpowers/commit-msg-3.txt
```

---

### Task 4: Deterministic Lua serialization (single file)

**Files:**
- Create: `api/internal/dataaddon/write.go`
- Create: `api/internal/dataaddon/write_test.go`

**Interfaces:**
- Consumes: `characterRow`, `guildRow`, `componentNames` (Tasks 2-3).
- Produces: `Data{Generated time.Time; Build string; Characters map[string]characterRow; Guilds map[string]guildRow}`, `Render(d Data) (map[string][]byte, error)` (Task 5 adds the split path this function falls back to), `sortedKeys[V any](m map[string]V) []string`. Task 6 (job.go) calls `Render`; Task 10 (golden test) asserts its exact output.

- [ ] **Step 1: Write the failing test**

```go
// api/internal/dataaddon/write_test.go
package dataaddon

import (
	"strings"
	"testing"
	"time"
)

func TestRenderProducesOneSortedDeterministicFile(t *testing.T) {
	d := Data{
		Generated: time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC),
		Build:     "1.60.1.69893",
		Characters: map[string]characterRow{
			"us:normal:thoradin": {Rating: 81, Components: map[string]int{
				"output": 91, "survival": 72, "mechanics": 84, "utility": 62, "preparation": 94, "activity": 91,
			}, Fights: 2},
			"eu:pvp:mörk": {Rating: 88, Components: map[string]int{"output": 70, "mechanics": 70}, Fights: 1},
		},
		Guilds: map[string]guildRow{
			"us:normal:iron-vanguard": {
				Name: "Iron Vanguard", Progress: "2/3", Nights: 2, Roster: 2,
				Members: []string{"o'malley", "thoradin"},
			},
		},
	}
	files, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("files = %v, want exactly Data.lua", mapKeysOf(files))
	}
	got := string(files["Data.lua"])
	if !strings.Contains(got, `generated = "2026-09-22T04:00:00Z"`) {
		t.Error("missing generated field")
	}
	if !strings.Contains(got, `build = "1.60.1.69893"`) {
		t.Error("missing build field")
	}
	// eu:pvp:mörk sorts before us:normal:thoradin (byte order: 'e' < 'u').
	if strings.Index(got, "eu:pvp:mörk") > strings.Index(got, "us:normal:thoradin") {
		t.Error("characters are not sorted by key")
	}
	if !strings.Contains(got, `["us:normal:thoradin"] = { rating = 81, output = 91, survival = 72, mechanics = 84, utility = 62, preparation = 94, activity = 91, fights = 2 },`) {
		t.Errorf("thoradin row wrong, got:\n%s", got)
	}
	if !strings.Contains(got, `["eu:pvp:mörk"] = { rating = 88, output = 70, mechanics = 70, fights = 1 },`) {
		t.Errorf("mörk row wrong (should omit the four unscored components), got:\n%s", got)
	}
	if !strings.Contains(got, `["us:normal:iron-vanguard"] = { name = "Iron Vanguard", progress = "2/3", nights = 2, roster = 2, members = { "o'malley", "thoradin" } },`) {
		t.Errorf("guild row wrong, got:\n%s", got)
	}
}

func TestRenderOmitsProgressWhenEmpty(t *testing.T) {
	d := Data{
		Generated: time.Now(), Build: "x",
		Guilds: map[string]guildRow{"us:normal:new-guild": {Name: "New Guild", Nights: 0, Roster: 1, Members: []string{"a"}}},
	}
	files, err := Render(d)
	if err != nil {
		t.Fatal(err)
	}
	got := string(files["Data.lua"])
	if strings.Contains(got, "progress") {
		t.Errorf("progress should be omitted entirely, got:\n%s", got)
	}
}

func mapKeysOf[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run TestRender -v`
Expected: FAIL — undefined: `Data`, `Render`.

- [ ] **Step 3: Write `write.go`**

```go
// api/internal/dataaddon/write.go
package dataaddon

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"time"
)

// isoFormat is Ratings.lua's own pattern:
// "^(%d%d%d%d)%-(%d%d)%-(%d%d)T(%d%d):(%d%d):(%d%d)Z$".
const isoFormat = "2006-01-02T15:04:05Z"

// dataLuaPath is the checked-in path this content replaces at release
// time -- echoed in the file's own header comment.
const dataLuaPath = "addon/ForeverSixtyData/Data.lua"

// Data is everything one Render call writes.
type Data struct {
	Generated  time.Time
	Build      string
	Characters map[string]characterRow
	Guilds     map[string]guildRow
}

// Render returns the file(s) to publish: one ("Data.lua") when the whole
// file fits in maxSingleFileBytes, or a header plus one per region when it
// does not (renderSplit, write_split.go). Every file returned, together,
// defines the same ForeverSixtyData global the single-file form would --
// Ratings.lua reads one global regardless of how many files built it.
func Render(d Data) (map[string][]byte, error) {
	return renderWithLimit(d, maxSingleFileBytes)
}

// renderWithLimit is Render with the split threshold spelled out, so a
// test can exercise the split path without generating a real 4 MiB fixture.
func renderWithLimit(d Data, limit int) (map[string][]byte, error) {
	var single bytes.Buffer
	if err := render(&single, d); err != nil {
		return nil, err
	}
	if single.Len() <= limit {
		return map[string][]byte{"Data.lua": single.Bytes()}, nil
	}
	return renderSplit(d)
}

func render(w io.Writer, d Data) error {
	bw := bufio.NewWriter(w)
	writeHeader(bw, d)
	fmt.Fprintln(bw, "ForeverSixtyData = {")
	fmt.Fprintln(bw, "\tformat = 1,")
	fmt.Fprintf(bw, "\tgenerated = %q,\n", d.Generated.UTC().Format(isoFormat))
	fmt.Fprintf(bw, "\tbuild = %q,\n", d.Build)
	fmt.Fprintln(bw, "\tcharacters = {")
	for _, key := range sortedKeys(d.Characters) {
		writeCharacterRow(bw, key, d.Characters[key])
	}
	fmt.Fprintln(bw, "\t},")
	fmt.Fprintln(bw, "\tguilds = {")
	for _, key := range sortedKeys(d.Guilds) {
		writeGuildRow(bw, key, d.Guilds[key])
	}
	fmt.Fprintln(bw, "\t},")
	fmt.Fprintln(bw, "}")
	return bw.Flush()
}

func writeHeader(bw *bufio.Writer, d Data) {
	fmt.Fprintf(bw, "-- %s\n", dataLuaPath)
	fmt.Fprintf(bw, "-- Generated %s for data build %s by the nightly data-addon job (api/internal/dataaddon).\n",
		d.Generated.UTC().Format(isoFormat), d.Build)
	fmt.Fprintln(bw, "-- Do not edit by hand: the addon-data-release workflow overwrites this file every night.")
}

func writeCharacterRow(bw *bufio.Writer, key string, row characterRow) {
	fmt.Fprintf(bw, "\t\t[%q] = { rating = %d", key, row.Rating)
	for _, name := range componentNames {
		if v, ok := row.Components[name]; ok {
			fmt.Fprintf(bw, ", %s = %d", name, v)
		}
	}
	fmt.Fprintf(bw, ", fights = %d },\n", row.Fights)
}

func writeGuildRow(bw *bufio.Writer, key string, row guildRow) {
	fmt.Fprintf(bw, "\t\t[%q] = { name = %q", key, row.Name)
	if row.Progress != "" {
		fmt.Fprintf(bw, ", progress = %q", row.Progress)
	}
	fmt.Fprintf(bw, ", nights = %d, roster = %d, members = {", row.Nights, row.Roster)
	for i, m := range row.Members {
		if i > 0 {
			fmt.Fprint(bw, ", ")
		}
		fmt.Fprintf(bw, "%q", m)
	}
	fmt.Fprintln(bw, "} },")
}

// sortedKeys returns m's keys in ascending byte order, for deterministic
// output regardless of Go's randomized map iteration.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -run TestRender -v`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/write.go internal/dataaddon/write_test.go
printf 'feat(api): data-addon Data.lua serialization (sorted keys, %%q escaping)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-4.txt
git commit -F ../.superpowers/commit-msg-4.txt
```

---

### Task 5: Region-split serialization (the 4 MiB fallback)

**Files:**
- Create: `api/internal/dataaddon/write_split.go`
- Create: `api/internal/dataaddon/write_split_test.go`

**Interfaces:**
- Consumes: `Data`, `renderWithLimit` (Task 4).
- Produces: `renderSplit(d Data) (map[string][]byte, error)`, `regionsOf(d Data) []string`. `renderWithLimit` (Task 4) already calls `renderSplit`; this task only adds the function itself, so Task 4's code starts compiling.

- [ ] **Step 1: Write the failing test**

```go
// api/internal/dataaddon/write_split_test.go
package dataaddon

import (
	"strings"
	"testing"
	"time"
)

func TestRenderWithLimitSplitsByRegionWhenOverTheLimit(t *testing.T) {
	d := Data{
		Generated: time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC),
		Build:     "1.60.1.69893",
		Characters: map[string]characterRow{
			"us:normal:thoradin": {Rating: 81, Components: map[string]int{"output": 91}, Fights: 2},
			"eu:pvp:mörk":        {Rating: 88, Components: map[string]int{"output": 70}, Fights: 1},
		},
		Guilds: map[string]guildRow{
			"us:normal:iron-vanguard": {Name: "Iron Vanguard", Nights: 1, Roster: 1, Members: []string{"a"}},
		},
	}
	// A limit of 1 byte guarantees the single-file render (never under 1
	// byte) always exceeds it, forcing the split path deterministically.
	files, err := renderWithLimit(d, 1)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["Data.lua"]; !ok {
		t.Fatal("Data.lua (the header/init file) is missing from a split render")
	}
	if strings.Contains(string(files["Data.lua"]), "thoradin") {
		t.Error("the header file should carry no rows, only format/generated/build and empty tables")
	}
	us, ok := files["Data-us.lua"]
	if !ok {
		t.Fatal("Data-us.lua is missing")
	}
	if !strings.Contains(string(us), "thoradin") || strings.Contains(string(us), "mörk") {
		t.Errorf("Data-us.lua should carry only us: rows, got:\n%s", us)
	}
	if !strings.Contains(string(us), "ForeverSixtyData.characters[key] = row") {
		t.Errorf("Data-us.lua should merge into the shared global, got:\n%s", us)
	}
	eu, ok := files["Data-eu.lua"]
	if !ok {
		t.Fatal("Data-eu.lua is missing")
	}
	if !strings.Contains(string(eu), "mörk") || strings.Contains(string(eu), "thoradin") {
		t.Errorf("Data-eu.lua should carry only eu: rows, got:\n%s", eu)
	}
}

func TestRegionsOfListsEveryRegionAKeyMentions(t *testing.T) {
	d := Data{
		Characters: map[string]characterRow{"us:normal:a": {}, "eu:pvp:b": {}},
		Guilds:     map[string]guildRow{"kr:hardcore:c": {}},
	}
	got := regionsOf(d)
	want := []string{"eu", "kr", "us"}
	if len(got) != len(want) {
		t.Fatalf("regions = %v, want %v", got, want)
	}
	for i, r := range want {
		if got[i] != r {
			t.Errorf("regions[%d] = %q, want %q", i, got[i], r)
		}
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run 'TestRenderWithLimit|TestRegionsOf' -v`
Expected: FAIL — undefined: `renderSplit` (referenced by `renderWithLimit` in Task 4, so the whole package currently fails to compile once Task 4 landed — this task's Step 2 should already show a compile error, which is an acceptable "fails" form for this step).

- [ ] **Step 3: Write `write_split.go`**

```go
// api/internal/dataaddon/write_split.go
package dataaddon

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// renderSplit is Render's fallback for a combined file over
// maxSingleFileBytes: "Data.lua" becomes a small header/init file (format,
// generated, build, and two empty tables), and one "Data-<region>.lua" per
// region that has any character or guild row merges its own rows into the
// shared global. The TOC must list Data.lua before every Data-<region>.lua
// (addon/ForeverSixtyData/README.md documents the ordering requirement);
// addon-data-release.yml is the one place that ever needs to change the
// TOC's file list to match, since the checked-in TOC only ever lists
// Data.lua (see the plan's Task 12).
func renderSplit(d Data) (map[string][]byte, error) {
	out := map[string][]byte{}

	var header bytes.Buffer
	hw := bufio.NewWriter(&header)
	writeHeader(hw, d)
	fmt.Fprintln(hw, "-- The combined file exceeded 4 MiB, so this build is split by region; see Data-<region>.lua.")
	fmt.Fprintln(hw, "ForeverSixtyData = {")
	fmt.Fprintln(hw, "\tformat = 1,")
	fmt.Fprintf(hw, "\tgenerated = %q,\n", d.Generated.UTC().Format(isoFormat))
	fmt.Fprintf(hw, "\tbuild = %q,\n", d.Build)
	fmt.Fprintln(hw, "\tcharacters = {},")
	fmt.Fprintln(hw, "\tguilds = {},")
	fmt.Fprintln(hw, "}")
	if err := hw.Flush(); err != nil {
		return nil, err
	}
	out["Data.lua"] = header.Bytes()

	for _, region := range regionsOf(d) {
		var buf bytes.Buffer
		bw := bufio.NewWriter(&buf)
		fmt.Fprintf(bw, "-- addon/ForeverSixtyData/Data-%s.lua\n", region)
		fmt.Fprintln(bw, "-- One region's share of the nightly data-addon job's output; see Data.lua.")
		fmt.Fprintln(bw, "for key, row in pairs({")
		for _, key := range sortedKeys(d.Characters) {
			if !strings.HasPrefix(key, region+":") {
				continue
			}
			writeCharacterRow(bw, key, d.Characters[key])
		}
		fmt.Fprintln(bw, "}) do ForeverSixtyData.characters[key] = row end")
		fmt.Fprintln(bw, "for key, row in pairs({")
		for _, key := range sortedKeys(d.Guilds) {
			if !strings.HasPrefix(key, region+":") {
				continue
			}
			writeGuildRow(bw, key, d.Guilds[key])
		}
		fmt.Fprintln(bw, "}) do ForeverSixtyData.guilds[key] = row end")
		if err := bw.Flush(); err != nil {
			return nil, err
		}
		out["Data-"+region+".lua"] = buf.Bytes()
	}
	return out, nil
}

// regionsOf lists every region any character or guild key names, sorted --
// the leading segment before the first ":" of an addon key built by
// characterKey/guildKey.
func regionsOf(d Data) []string {
	set := map[string]bool{}
	for key := range d.Characters {
		set[strings.SplitN(key, ":", 2)[0]] = true
	}
	for key := range d.Guilds {
		set[strings.SplitN(key, ":", 2)[0]] = true
	}
	out := make([]string, 0, len(set))
	for r := range set {
		out = append(out, r)
	}
	sort.Strings(out)
	return out
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -v`
Expected: PASS (every test in the package so far).

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/write_split.go internal/dataaddon/write_split_test.go
printf 'feat(api): data-addon region-split Data.lua fallback for files over 4 MiB\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-5.txt
git commit -F ../.superpowers/commit-msg-5.txt
```

---

### Task 6: Database reads

**Files:**
- Create: `api/internal/dataaddon/store.go`
- Create: `api/internal/dataaddon/store_test.go`

**Interfaces:**
- Consumes: `guildIdentity` (Task 3).
- Produces: `Store{Pool *pgxpool.Pool}`, `characterFightRow{PlayerKey string; Overall float64; Components json.RawMessage}`, `(*Store).characterFights(ctx, since time.Time) ([]characterFightRow, error)`, `(*Store).guildsWithVerifiedMembers(ctx) ([]guildIdentity, error)`, `(*Store).verifiedMembers(ctx, guildIDs []int64) (map[int64][]string, error)`, `(*Store).nightsByGuild(ctx, guildIDs []int64, since time.Time) (map[int64]int, error)`, `(*Store).progressionByGuild(ctx, guildIDs []int64) (killed, total map[int64]int, err error)`. Task 8 (job.go) calls every one of these.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/dataaddon/store_test.go
package dataaddon

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// seedUser inserts a minimal users row and returns its id.
func seedUser(t *testing.T, pool interface {
	QueryRow(ctx context.Context, sql string, args ...any) interface {
		Scan(dest ...any) error
	}
}, anonymize bool) int64 {
	t.Helper()
	return 0 // replaced below; kept here only to name the shape this step's inline helpers use
}

func TestCharacterFightsAppliesTheWindowVisibilityAndAnonymizeRules(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var normalUser, anonUser int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Normal') returning id`).Scan(&normalUser); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag, anonymize) values ('Hidden', true) returning id`).Scan(&anonUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from users where id in ($1, $2)`, normalUser, anonUser)
	})

	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values
		 ('us/normal/thoradin', 'us', 'normal', 'Thoradin', $1),
		 ('us/normal/hiddenhero', 'us', 'normal', 'Hiddenhero', $2)`, normalUser, anonUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key in ('us/normal/thoradin', 'us/normal/hiddenhero')`)
	})

	insertReport := func(id, visibility string, createdAt time.Time) {
		if _, err := pool.Exec(ctx,
			`insert into reports (id, visibility, status, created_at) values ($1, $2, 'complete', $3)`,
			id, visibility, createdAt); err != nil {
			t.Fatal(err)
		}
	}
	insertReport("dataaddon-store-public", "public", now.Add(-2*24*time.Hour))
	insertReport("dataaddon-store-private", "private", now.Add(-2*24*time.Hour))
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from reports where id in ('dataaddon-store-public', 'dataaddon-store-private')`)
	})

	insertScore := func(reportID string, index int, playerKey string, overall float64, foughtAt time.Time) {
		// rating_scores is partitioned by month on fought_at (migration
		// 0022): a plain insert with no matching partition fails outright,
		// so every write through this test helper ensures one first, the
		// same way rating.Store.RateFight does before its own insert.
		if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
			 values ($1, $2, $3, $3, 1, false, $4, $4, false, '[]', 'test', $5)`,
			reportID, index, playerKey, overall, foughtAt); err != nil {
			t.Fatal(err)
		}
	}
	insertScore("dataaddon-store-public", 1, "us/normal/thoradin", 80, now.Add(-2*24*time.Hour))       // in window, public: counts
	insertScore("dataaddon-store-public", 2, "us/normal/thoradin", 999, now.Add(-100*24*time.Hour))    // out of window: excluded
	insertScore("dataaddon-store-private", 1, "us/normal/thoradin", 999, now.Add(-1*24*time.Hour))     // private report: excluded
	insertScore("dataaddon-store-public", 3, "us/normal/hiddenhero", 999, now.Add(-2*24*time.Hour))    // anonymized owner: excluded
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id in ('dataaddon-store-public', 'dataaddon-store-private')`)
	})

	store := &Store{Pool: pool}
	rows, err := store.characterFights(ctx, now.Add(-ratingWindow))
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatalf("rows = %d, want exactly 1 (thoradin's in-window public fight); got %+v", len(rows), rows)
	}
	if rows[0].PlayerKey != "us/normal/thoradin" || rows[0].Overall != 80 {
		t.Errorf("row = %+v", rows[0])
	}
}

func TestGuildQueriesReadVerifiedMembersNightsAndProgression(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var u1, u2 int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('One') returning id`).Scan(&u1); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Two') returning id`).Scan(&u2); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from users where id in ($1, $2)`, u1, u2) })

	var guildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Iron Vanguard') returning id`,
	).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guilds where id = $1`, guildID) })

	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, verified_at) values
		 ($1, 'us/normal/thoradin', $2, now()),
		 ($1, 'us/normal/notyet', $2, null)`, guildID, u1); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guild_characters where guild_id = $1`, guildID) })

	insertReport := func(id string, createdAt time.Time) {
		if _, err := pool.Exec(ctx,
			`insert into reports (id, guild_id, visibility, status, created_at) values ($1, $2, 'public', 'complete', $3)`,
			id, guildID, createdAt); err != nil {
			t.Fatal(err)
		}
	}
	insertReport("dataaddon-guild-1", now.Add(-2*24*time.Hour))
	insertReport("dataaddon-guild-2", now.Add(-10*24*time.Hour))
	t.Cleanup(func() { pool.Exec(ctx, `delete from reports where guild_id = $1`, guildID) })

	insertFight := func(reportID string, index int, encounterID int64, kill bool) {
		if _, err := pool.Exec(ctx,
			`insert into fights (report_id, fight_index, encounter_id, kill) values ($1, $2, $3, $4)`,
			reportID, index, encounterID, kill); err != nil {
			t.Fatal(err)
		}
	}
	insertFight("dataaddon-guild-1", 1, 100, true)
	insertFight("dataaddon-guild-1", 2, 101, false)
	insertFight("dataaddon-guild-2", 1, 100, true)
	insertFight("dataaddon-guild-2", 2, 102, true)
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from fights where report_id in ('dataaddon-guild-1', 'dataaddon-guild-2')`)
	})

	store := &Store{Pool: pool}
	identities, err := store.guildsWithVerifiedMembers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, g := range identities {
		if g.ID == guildID {
			found = true
		}
	}
	if !found {
		t.Fatal("the guild with a verified member should be listed")
	}

	members, err := store.verifiedMembers(ctx, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	if len(members[guildID]) != 1 || members[guildID][0] != "us/normal/thoradin" {
		t.Errorf("verified members = %v, want exactly [us/normal/thoradin]", members[guildID])
	}

	nights, err := store.nightsByGuild(ctx, []int64{guildID}, now.Add(-ratingWindow))
	if err != nil {
		t.Fatal(err)
	}
	if nights[guildID] != 2 {
		t.Errorf("nights = %d, want 2", nights[guildID])
	}

	killed, total, err := store.progressionByGuild(ctx, []int64{guildID})
	if err != nil {
		t.Fatal(err)
	}
	if killed[guildID] != 2 || total[guildID] != 3 {
		t.Errorf("progression = %d/%d, want 2/3", killed[guildID], total[guildID])
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run 'TestCharacterFights|TestGuildQueries' -v`
Expected: FAIL — undefined: `Store`.

- [ ] **Step 3: Write `store.go`**

```go
// api/internal/dataaddon/store.go
package dataaddon

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store is every database read this job needs.
type Store struct {
	Pool *pgxpool.Pool
}

// characterFightRow is one public, in-window, non-anonymized rated fight.
type characterFightRow struct {
	PlayerKey  string
	Overall    float64
	Components json.RawMessage
}

// characterFights reads every rating_scores row fought_at or after since,
// whose report is public, for a player_key whose owning account has not
// asked to be anonymized -- the same rule api/internal/rating.Store's
// anonymized and ReadCharacterRating apply per character, applied here
// across every character at once. A player_key with no characters row at
// all (never linked to an account) is not anonymized, matching
// rating.Store.anonymized's own "nothing to hide" rule exactly.
func (s *Store) characterFights(ctx context.Context, since time.Time) ([]characterFightRow, error) {
	rows, err := s.Pool.Query(ctx,
		`select rs.player_key, rs.overall, rs.components
		 from rating_scores rs
		 join reports r on r.id = rs.report_id
		 where r.visibility = 'public' and rs.fought_at >= $1
		   and not exists (
		     select 1 from characters c join users u on u.id = c.user_id
		     where c.key = rs.player_key and u.anonymize
		   )
		 order by rs.player_key`, since)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read character fights: %w", err)
	}
	defer rows.Close()
	var out []characterFightRow
	for rows.Next() {
		var r characterFightRow
		if err := rows.Scan(&r.PlayerKey, &r.Overall, &r.Components); err != nil {
			return nil, fmt.Errorf("dataaddon: scan character fight: %w", err)
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// guildsWithVerifiedMembers reads every guild that has at least one
// verified guild_characters row -- a guild with none has nothing
// corroborated to publish (dispatch: "only VERIFIED members count"), so it
// is left out of this list entirely and never gets a Data.lua row.
func (s *Store) guildsWithVerifiedMembers(ctx context.Context) ([]guildIdentity, error) {
	rows, err := s.Pool.Query(ctx,
		`select g.id, g.region, g.ruleset, g.name
		 from guilds g
		 where exists (
		   select 1 from guild_characters gc
		   where gc.guild_id = g.id and gc.verified_at is not null
		 )
		 order by g.id`)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read guilds: %w", err)
	}
	defer rows.Close()
	var out []guildIdentity
	for rows.Next() {
		var g guildIdentity
		if err := rows.Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name); err != nil {
			return nil, fmt.Errorf("dataaddon: scan guild: %w", err)
		}
		out = append(out, g)
	}
	return out, rows.Err()
}

// verifiedMembers reads every verified guild_characters row for the given
// guild ids, keyed by guild_id.
func (s *Store) verifiedMembers(ctx context.Context, guildIDs []int64) (map[int64][]string, error) {
	rows, err := s.Pool.Query(ctx,
		`select guild_id, character_key from guild_characters
		 where verified_at is not null and guild_id = any($1)
		 order by guild_id, character_key`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read verified members: %w", err)
	}
	defer rows.Close()
	out := map[int64][]string{}
	for rows.Next() {
		var id int64
		var key string
		if err := rows.Scan(&id, &key); err != nil {
			return nil, fmt.Errorf("dataaddon: scan verified member: %w", err)
		}
		out[id] = append(out[id], key)
	}
	return out, rows.Err()
}

// nightsByGuild counts, per guild id, the distinct UTC calendar dates of
// that guild's public reports created at or after since -- this job's own
// definition of "a raid night" (see the plan's aggregation-rule section).
func (s *Store) nightsByGuild(ctx context.Context, guildIDs []int64, since time.Time) (map[int64]int, error) {
	rows, err := s.Pool.Query(ctx,
		`select guild_id, count(distinct date(created_at))
		 from reports
		 where guild_id = any($1) and visibility = 'public' and created_at >= $2
		 group by guild_id`, guildIDs, since)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: read guild nights: %w", err)
	}
	defer rows.Close()
	out := map[int64]int{}
	for rows.Next() {
		var id int64
		var n int
		if err := rows.Scan(&id, &n); err != nil {
			return nil, fmt.Errorf("dataaddon: scan guild nights: %w", err)
		}
		out[id] = n
	}
	return out, rows.Err()
}

// progressionByGuild counts, per guild id, encounters killed at least once
// against every encounter attempted, from fights attached to a report the
// guild's own public page would show
// (api/internal/rankings/guilds.go's own Guild() progression query:
// r.visibility <> 'private'). Duplicated here rather than imported:
// api/internal/rankings is outside this lane's file ownership, and
// api/internal/rating/backfill.go's own readSummary already sets the
// precedent in this codebase for duplicating one small, owned query across
// a lane boundary rather than reaching across it.
func (s *Store) progressionByGuild(ctx context.Context, guildIDs []int64) (killed, total map[int64]int, err error) {
	rows, err := s.Pool.Query(ctx,
		`select r.guild_id, f.encounter_id, count(*) filter (where f.kill)
		 from fights f join reports r on r.id = f.report_id
		 where r.guild_id = any($1) and f.encounter_id is not null and r.visibility <> 'private'
		 group by r.guild_id, f.encounter_id`, guildIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("dataaddon: read guild progression: %w", err)
	}
	defer rows.Close()
	killed, total = map[int64]int{}, map[int64]int{}
	for rows.Next() {
		var id, encounterID int64
		var kills int
		if err := rows.Scan(&id, &encounterID, &kills); err != nil {
			return nil, nil, fmt.Errorf("dataaddon: scan guild progression: %w", err)
		}
		total[id]++
		if kills > 0 {
			killed[id]++
		}
	}
	return killed, total, rows.Err()
}
```

- [ ] **Step 4: Remove the placeholder `seedUser` helper**

`seedUser` in Step 1's test file was scaffolding to think through the shape and is unused by the final tests (each test inlines its own inserts) — delete it now so `go vet` does not flag dead code:

```bash
cd api
sed -i '' '/^func seedUser/,/^}/d' internal/dataaddon/store_test.go
```

- [ ] **Step 5: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -v -p 1`
Expected: PASS. (If `TEST_DATABASE_URL` is unset, these two tests SKIP rather than fail — export it per Task 1 Step 1 before running: `export TEST_DATABASE_URL='postgres://forever:forever@localhost:5434/dataaddon_test?sslmode=disable'`.)

- [ ] **Step 6: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/store.go internal/dataaddon/store_test.go
printf 'feat(api): data-addon database reads (character fights, guild membership, progression)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-6.txt
git commit -F ../.superpowers/commit-msg-6.txt
```

---

### Task 7: Cloud Storage upload

**Files:**
- Create: `api/internal/dataaddon/upload.go`
- Create: `api/internal/dataaddon/upload_test.go`

**Interfaces:**
- Consumes: nothing new (uses only `context`, `google.golang.org/api/storage/v1`, `google.golang.org/api/option` — `google.golang.org/api` is already a direct dependency of the `api` module per `api/go.mod`; no `go.mod`/`go.sum` change is needed).
- Produces: `Uploader` interface (`Upload(ctx context.Context, bucket, object string, data []byte) error`), `GCS` (real implementation), `NewGCS(ctx context.Context) (*GCS, error)`, `NewGCSWith(ctx context.Context, opts ...option.ClientOption) (*GCS, error)`, `FakeUploader` (test double recording calls). Task 8 (job.go) takes an `Uploader`; Task 9 (main.go) constructs `NewGCS`.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/dataaddon/upload_test.go
package dataaddon

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/api/option"
)

func TestFakeUploaderRecordsWhatWasUploaded(t *testing.T) {
	f := &FakeUploader{}
	if err := f.Upload(context.Background(), "my-bucket", "data-addon/Data.lua", []byte("x")); err != nil {
		t.Fatal(err)
	}
	if len(f.Calls) != 1 || f.Calls[0].Bucket != "my-bucket" || f.Calls[0].Object != "data-addon/Data.lua" {
		t.Fatalf("calls = %+v", f.Calls)
	}
	if string(f.Calls[0].Data) != "x" {
		t.Errorf("data = %q", f.Calls[0].Data)
	}
}

func TestFakeUploaderReportsItsConfiguredFailure(t *testing.T) {
	f := &FakeUploader{Err: context.DeadlineExceeded}
	if err := f.Upload(context.Background(), "b", "o", nil); err == nil {
		t.Fatal("the configured error should surface")
	}
	if len(f.Calls) != 0 {
		t.Fatal("a failed upload records nothing")
	}
}

func TestGCSUploadPUTsTheObjectToTheStorageJSONAPI(t *testing.T) {
	var gotPath, gotBucket string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotBucket = r.URL.Query().Get("name")
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"name": gotBucket})
	}))
	defer srv.Close()

	g, err := NewGCSWith(context.Background(), option.WithEndpoint(srv.URL), option.WithoutAuthentication())
	if err != nil {
		t.Fatal(err)
	}
	if err := g.Upload(context.Background(), "my-bucket", "data-addon/Data.lua", []byte("ForeverSixtyData = {}")); err != nil {
		t.Fatal(err)
	}
	if gotPath == "" {
		t.Fatal("no request reached the fake server")
	}
	if string(gotBody) != "" && !contains(string(gotBody), "ForeverSixtyData") {
		// The storage/v1 client sends a multipart body (metadata + media);
		// this only asserts the payload made it across, not the exact
		// multipart framing.
		t.Errorf("body = %q, want it to contain the uploaded content", gotBody)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})()
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run 'TestFakeUploader|TestGCSUpload' -v`
Expected: FAIL — undefined: `FakeUploader`, `NewGCSWith`.

- [ ] **Step 3: Write `upload.go`**

```go
// api/internal/dataaddon/upload.go
package dataaddon

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"google.golang.org/api/option"
	storage "google.golang.org/api/storage/v1"
)

// Uploader publishes one named object's bytes. *GCS satisfies it; a nil
// Uploader (or an empty bucket name) means this deployment has no bucket
// configured, and job.go's Run logs why and skips the upload rather than
// failing the run.
type Uploader interface {
	Upload(ctx context.Context, bucket, object string, data []byte) error
}

// GCS uploads through the Cloud Storage JSON API, using the runtime
// service account's Application Default Credentials -- the same
// credential model api/internal/jobs.CloudRun already uses for the Cloud
// Run Admin API, which is why this follows that file's own constructor
// shape (NewX / NewXWith) closely.
type GCS struct {
	svc *storage.Service
}

// NewGCS builds an uploader using Application Default Credentials. Nothing
// is dialled here.
func NewGCS(ctx context.Context) (*GCS, error) {
	return NewGCSWith(ctx)
}

// NewGCSWith is NewGCS with the client options spelled out, so a test can
// point the uploader at an in-process fake server.
func NewGCSWith(ctx context.Context, opts ...option.ClientOption) (*GCS, error) {
	svc, err := storage.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: storage client: %w", err)
	}
	return &GCS{svc: svc}, nil
}

// Upload writes data to gs://bucket/object, overwriting any existing
// object at that name -- every publish is a full overwrite, matching
// "Do not edit by hand: the addon-data-release workflow overwrites this
// file every night."
func (g *GCS) Upload(ctx context.Context, bucket, object string, data []byte) error {
	_, err := g.svc.Objects.Insert(bucket, &storage.Object{Name: object}).
		Media(bytes.NewReader(data)).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("dataaddon: upload gs://%s/%s: %w", bucket, object, err)
	}
	return nil
}

// FakeUploader records what it was asked to upload, for job.go's tests.
type FakeUploader struct {
	mu    sync.Mutex
	Calls []UploadCall
	Err   error
}

// UploadCall is one recorded FakeUploader.Upload invocation.
type UploadCall struct {
	Bucket, Object string
	Data           []byte
}

func (f *FakeUploader) Upload(_ context.Context, bucket, object string, data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.Calls = append(f.Calls, UploadCall{Bucket: bucket, Object: object, Data: append([]byte{}, data...)})
	return nil
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -v`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/upload.go internal/dataaddon/upload_test.go
printf 'feat(api): data-addon Cloud Storage upload (google.golang.org/api/storage/v1)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-7.txt
git commit -F ../.superpowers/commit-msg-7.txt
```

---

### Task 8: Job orchestration (`Run`)

**Files:**
- Create: `api/internal/dataaddon/job.go`
- Create: `api/internal/dataaddon/job_test.go`

**Interfaces:**
- Consumes: `Store` and every method (Task 6), `Uploader`/`FakeUploader` (Task 7), `Render` (Tasks 4-5), `aggregateCharacter`/`decodeComponents` (Task 2), `buildGuildRow` (Task 3), `characterKey`/`guildKey` (Task 1).
- Produces: `Logger` interface, `Deps{Store *Store; Upload Uploader; Bucket, Build string; Now time.Time; Log Logger}`, `Result{Characters, Guilds, SkippedRows int; Files map[string]int}`, `Run(ctx context.Context, d Deps) (Result, error)`. Task 9 (main.go) calls `Run` directly.

- [ ] **Step 1: Write the failing tests**

```go
// api/internal/dataaddon/job_test.go
package dataaddon

import (
	"context"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

type recordingLogger struct {
	warnings []string
}

func (l *recordingLogger) Info(string, ...any) {}
func (l *recordingLogger) Warn(msg string, args ...any) {
	l.warnings = append(l.warnings, msg)
}
func (l *recordingLogger) Error(string, ...any) {}

func TestRunSkipsAnUnparseablePlayerKeyAndCountsIt(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)
	foughtAt := now.Add(-2 * 24 * time.Hour)

	if _, err := pool.Exec(ctx,
		`insert into reports (id, visibility, status, created_at) values ('dataaddon-job-bad', 'public', 'complete', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from reports where id = 'dataaddon-job-bad'`) })

	// rating_scores is partitioned by month on fought_at (migration 0022);
	// ensure the partition exists before inserting into it, the same way
	// rating.Store.RateFight does before its own insert.
	if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
		   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
		 values ('dataaddon-job-bad', 1, 'not-a-valid-key', 'x', 1, false, 50, 50, false, '[]', 'test', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = 'dataaddon-job-bad'`) })

	log := &recordingLogger{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: now, Log: log,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.SkippedRows != 1 {
		t.Errorf("skipped = %d, want 1", result.SkippedRows)
	}
	if len(log.warnings) == 0 {
		t.Error("the skip should be logged")
	}
}

func TestRunSkipsTheUploadAndLogsWhyWhenNoBucketIsConfigured(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	log := &recordingLogger{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: time.Now().UTC(), Log: log,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Files["Data.lua"] == 0 {
		t.Error("Data.lua should still be rendered (and its size logged) even with no bucket")
	}
	found := false
	for _, w := range log.warnings {
		if w == "dataaddon" {
			found = true
		}
	}
	if !found {
		t.Error("the skipped upload should be logged")
	}
}

func TestRunUploadsEveryRenderedFileAndAManifest(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	up := &FakeUploader{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Upload: up, Bucket: "my-bucket",
		Build: "1.60.1.69893", Now: time.Now().UTC(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Files["Data.lua"] == 0 {
		t.Fatal("Data.lua should be rendered")
	}
	var sawData, sawManifest bool
	for _, c := range up.Calls {
		if c.Bucket != "my-bucket" {
			t.Errorf("bucket = %q, want my-bucket", c.Bucket)
		}
		if c.Object == "data-addon/Data.lua" {
			sawData = true
		}
		if c.Object == "data-addon/manifest.txt" {
			sawManifest = true
			if string(c.Data) != "Data.lua\n" {
				t.Errorf("manifest = %q, want \"Data.lua\\n\"", c.Data)
			}
		}
	}
	if !sawData || !sawManifest {
		t.Errorf("calls = %+v, want Data.lua and manifest.txt uploaded", up.Calls)
	}
}
```

- [ ] **Step 2: Run to verify failure**

Run: `cd api && go test ./internal/dataaddon/... -run TestRun -v`
Expected: FAIL — undefined: `Run`, `Deps`, `Result`.

- [ ] **Step 3: Write `job.go`**

```go
// api/internal/dataaddon/job.go
package dataaddon

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Logger is the subset of *slog.Logger this package needs -- *slog.Logger
// satisfies it structurally, with no adapter, so api/cmd/api/main.go can
// pass its own logger straight through.
type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// Deps is everything one run of the job needs.
type Deps struct {
	Store  *Store
	Upload Uploader // nil (or an empty Bucket) means this deployment has no bucket; the run still succeeds and logs why the upload was skipped
	Bucket string
	Build  string // the data build the ratings were computed against
	Now    time.Time
	Log    Logger
}

func (d Deps) logger() Logger {
	if d.Log != nil {
		return d.Log
	}
	return noopLogger{}
}

type noopLogger struct{}

func (noopLogger) Info(string, ...any)  {}
func (noopLogger) Warn(string, ...any)  {}
func (noopLogger) Error(string, ...any) {}

// Result is what one run produced.
type Result struct {
	Characters  int            // rows written
	Guilds      int            // rows written
	SkippedRows int            // rows a bad player_key could not be formatted for
	Files       map[string]int // filename -> byte size (the dispatch's "size line")
}

// Run reads the database, aggregates, renders Data.lua (or its region-split
// form), and uploads it, in that order. It never fails the whole run over
// one bad row (dispatch: "skipped with a logged reason and counted"); it
// only returns an error for something that stops the run outright -- a
// database read failing, a render failing, or an upload failing.
func Run(ctx context.Context, d Deps) (Result, error) {
	since := d.Now.Add(-ratingWindow)

	fights, err := d.Store.characterFights(ctx, since)
	if err != nil {
		return Result{}, err
	}
	byPlayer := map[string][]fightScore{}
	for _, f := range fights {
		byPlayer[f.PlayerKey] = append(byPlayer[f.PlayerKey], fightScore{
			Overall: f.Overall, Components: decodeComponents(f.Components),
		})
	}

	var skipped int
	characters := map[string]characterRow{}
	for playerKey, fs := range byPlayer {
		region, ruleset, name, ok := splitCharacterKey(playerKey)
		if !ok {
			d.logger().Warn("dataaddon", "op", "character", "err", "unparseable player_key "+playerKey)
			skipped++
			continue
		}
		row, ok := aggregateCharacter(fs)
		if !ok {
			continue
		}
		characters[characterKey(region, ruleset, name)] = row
	}

	guildIdentities, err := d.Store.guildsWithVerifiedMembers(ctx)
	if err != nil {
		return Result{}, err
	}
	guilds := map[string]guildRow{}
	if len(guildIdentities) > 0 {
		ids := make([]int64, len(guildIdentities))
		for i, g := range guildIdentities {
			ids[i] = g.ID
		}
		members, err := d.Store.verifiedMembers(ctx, ids)
		if err != nil {
			return Result{}, err
		}
		nights, err := d.Store.nightsByGuild(ctx, ids, since)
		if err != nil {
			return Result{}, err
		}
		killed, total, err := d.Store.progressionByGuild(ctx, ids)
		if err != nil {
			return Result{}, err
		}
		for _, g := range guildIdentities {
			guilds[guildKey(g.Region, g.Ruleset, g.Name)] = buildGuildRow(
				g, members[g.ID], nights[g.ID], killed[g.ID], total[g.ID])
		}
	}

	files, err := Render(Data{Generated: d.Now, Build: d.Build, Characters: characters, Guilds: guilds})
	if err != nil {
		return Result{}, fmt.Errorf("dataaddon: render: %w", err)
	}

	sizes := map[string]int{}
	for name, body := range files {
		sizes[name] = len(body)
		d.logger().Info("dataaddon", "op", "render", "file", name, "bytes", len(body))
	}
	if d.Upload == nil || d.Bucket == "" {
		d.logger().Warn("dataaddon", "op", "upload",
			"err", "no-op: DATA_ADDON_BUCKET is not set, this deployment publishes no addon data")
	} else {
		for name, body := range files {
			if err := d.Upload.Upload(ctx, d.Bucket, "data-addon/"+name, body); err != nil {
				return Result{}, fmt.Errorf("dataaddon: upload %s: %w", name, err)
			}
		}
		if err := d.Upload.Upload(ctx, d.Bucket, "data-addon/manifest.txt", manifestOf(files)); err != nil {
			return Result{}, fmt.Errorf("dataaddon: upload manifest: %w", err)
		}
	}

	return Result{
		Characters: len(characters), Guilds: len(guilds), SkippedRows: skipped, Files: sizes,
	}, nil
}

// manifestOf lists every published filename, one per line, sorted -- what
// addon-data-release.yml reads to know which files to download and which
// TOC lines to write, without guessing from a fixed list of regions that
// may not all be present.
func manifestOf(files map[string][]byte) []byte {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	var buf bytes.Buffer
	for _, name := range names {
		buf.WriteString(name)
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

// splitCharacterKey reads region/ruleset/name back out of a player_key
// ("us/normal/thoradin" -> "us", "normal", "thoradin", true) -- mirrors
// api/internal/rating/backfill.go's own splitPlayerKeyRegionRuleset,
// extended to also return the name segment this package needs.
func splitCharacterKey(key string) (region, ruleset, name string, ok bool) {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", false
	}
	return parts[0], parts[1], parts[2], true
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && go test ./internal/dataaddon/... -v -p 1`
Expected: PASS — every test in the package.

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/job.go internal/dataaddon/job_test.go
printf 'feat(api): data-addon job orchestration (Run: read, aggregate, render, publish)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-8.txt
git commit -F ../.superpowers/commit-msg-8.txt
```

---

### Task 9: `main.go` dispatch

**Files:**
- Modify: `api/cmd/api/main.go:82` (add a `case` beside `rating.BackfillJobCommand`) and after `runRatingBackfill` (currently ending at line 271, just before `simEngine` at line 278) (add `runDataAddon`)

**Interfaces:**
- Consumes: `dataaddon.JobCommand`, `dataaddon.Run`, `dataaddon.Deps`, `dataaddon.Store`, `dataaddon.NewGCS` (Tasks 1, 7, 8).
- Produces: nothing new for later tasks — this is the production wiring.

- [ ] **Step 1: Add the import**

In `api/cmd/api/main.go`, add one line to the import block (alphabetically, right after `"github.com/jhunthrop/foreversixty/api/internal/config"`):

```go
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/dataaddon"
	"github.com/jhunthrop/foreversixty/api/internal/db"
```

- [ ] **Step 2: Add the dispatch `case`**

Immediately after the existing block:

```go
		case rating.BackfillJobCommand:
			if err := runRatingBackfill(context.Background(), log); err != nil {
				log.Error(rating.BackfillJobCommand, "err", err)
				os.Exit(1)
			}
			return
```

add:

```go
		case dataaddon.JobCommand:
			if err := runDataAddon(context.Background(), log); err != nil {
				log.Error(dataaddon.JobCommand, "err", err)
				os.Exit(1)
			}
			return
```

- [ ] **Step 3: Add `runDataAddon` right after `runRatingBackfill`**

```go
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
	if latest, ok := treeData.Latest(); ok {
		build = latest.Version
	} else {
		log.Warn(dataaddon.JobCommand, "err", "no client build available",
			"effect", "Data.lua's build field reads \"unknown\"")
	}

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
		Store: &dataaddon.Store{Pool: pool}, Upload: uploader, Bucket: bucket,
		Build: build, Now: time.Now().UTC(), Log: log,
	})
	if err != nil {
		return err
	}
	log.Info(dataaddon.JobCommand, "characters", result.Characters, "guilds", result.Guilds,
		"skipped", result.SkippedRows, "files", result.Files)
	return nil
}
```

- [ ] **Step 4: Build and vet the whole module**

Run: `cd api && go build ./... && go vet ./...`
Expected: no output, exit 0. (There is no existing test for the `case` dispatch of `rating.BackfillJobCommand` either — `main_test.go` only covers `newReportsService`'s wiring — so a clean build is this step's pass signal, matching that established precedent.)

- [ ] **Step 5: Commit**

```bash
cd api && go vet ./...
git add cmd/api/main.go
printf 'feat(api): dispatch the data-addon nightly job from main.go\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-9.txt
git commit -F ../.superpowers/commit-msg-9.txt
```

---

### Task 10: Golden test, Lua validity check, and the busted lockstep spec

**Files:**
- Create: `api/internal/dataaddon/golden_test.go`
- Create: `addon/tests/fixtures/data_addon_sample.lua`
- Create: `addon/tests/data_addon_spec.lua`

**Interfaces:**
- Consumes: everything from Tasks 1-8 (`Store`, `Run`, `Deps`, `Render`).
- Produces: nothing new — this is the dispatch's three required tests (golden byte-exact output, valid-Lua check, reader-accepts-it check) plus the key-format edge cases (spaces, apostrophes, non-ASCII already partly covered by Task 1/6; this task adds the full end-to-end case).

- [ ] **Step 1: Write the checked-in fixture file** (this is the byte-exact expected output both the Go test and the busted spec below assert against — write it first, by hand, so Step 3's Go test has something fixed to compare to)

```lua
-- addon/ForeverSixtyData/Data.lua
-- Generated 2026-09-22T04:00:00Z for data build 1.60.1.69893 by the nightly data-addon job (api/internal/dataaddon).
-- Do not edit by hand: the addon-data-release workflow overwrites this file every night.
ForeverSixtyData = {
	format = 1,
	generated = "2026-09-22T04:00:00Z",
	build = "1.60.1.69893",
	characters = {
		["eu:pvp:mörk"] = { rating = 88, output = 70, mechanics = 70, fights = 1 },
		["us:normal:thoradin"] = { rating = 81, output = 91, survival = 72, mechanics = 84, utility = 62, preparation = 94, activity = 91, fights = 2 },
	},
	guilds = {
		["us:normal:iron-vanguard"] = { name = "Iron Vanguard", progress = "2/3", nights = 2, roster = 2, members = { "o'malley", "thoradin" } },
	},
}
```

Save it at `addon/tests/fixtures/data_addon_sample.lua`.

- [ ] **Step 2: Add `Rendered` to `Result` first, so the golden test can assert on exact bytes**

`job.go` (Task 8) does not currently hand back the rendered bytes, only their sizes. Modify `Result` in `api/internal/dataaddon/job.go`:

```go
// Result is what one run produced.
type Result struct {
	Characters  int
	Guilds      int
	SkippedRows int
	Files       map[string]int   // filename -> byte size (the dispatch's "size line")
	Rendered    map[string][]byte // filename -> content, for a caller (the golden test) that needs the bytes themselves
}
```

and in `Run`, change:

```go
	sizes := map[string]int{}
	for name, body := range files {
		sizes[name] = len(body)
		d.logger().Info("dataaddon", "op", "render", "file", name, "bytes", len(body))
	}
```

to also keep `files` on the returned `Result`, and change the final `return` statement from:

```go
	return Result{
		Characters: len(characters), Guilds: len(guilds), SkippedRows: skipped, Files: sizes,
	}, nil
```

to:

```go
	return Result{
		Characters: len(characters), Guilds: len(guilds), SkippedRows: skipped,
		Files: sizes, Rendered: files,
	}, nil
```

Run: `cd api && go build ./internal/dataaddon/...` — expected: builds clean (this is an additive field; Task 8's existing tests do not assert on `Files` exhaustively, so they keep passing unmodified).

- [ ] **Step 3: Write the failing golden test**

```go
// api/internal/dataaddon/golden_test.go
package dataaddon

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// goldenFixturePath is the one file both this test and
// addon/tests/data_addon_spec.lua assert against, so the Go job and the
// Lua reader can never silently drift out of lockstep.
const goldenFixturePath = "../../../addon/tests/fixtures/data_addon_sample.lua"

// TestGoldenDataLuaIsByteExact builds the exact database fixture the
// plan's "Key-format reconciliation" and "Aggregation rule" sections work
// through by hand: Thoradin (every component present, two in-window
// public fights, one out-of-window fight and one private-report fight
// that must both be excluded), Mörk (a non-ASCII name on another
// region/ruleset, only two of six components computable), Hiddenhero (an
// anonymized owner, must not appear at all), and the guild Iron Vanguard
// (two verified members -- one with an apostrophe in its name -- one
// unverified member excluded from the roster, two public reports on
// different UTC dates, and progression 2 killed of 3 attempted
// encounters). One call to Run must reproduce the checked-in fixture
// byte-for-byte.
func TestGoldenDataLuaIsByteExact(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)

	var normalUser, hiddenUser, memberUser int64
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Normal') returning id`).Scan(&normalUser); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag, anonymize) values ('Hidden', true) returning id`).Scan(&hiddenUser); err != nil {
		t.Fatal(err)
	}
	if err := pool.QueryRow(ctx, `insert into users (battletag) values ('Member') returning id`).Scan(&memberUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from users where id in ($1, $2, $3)`, normalUser, hiddenUser, memberUser) })

	if _, err := pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values
		 ('us/normal/thoradin', 'us', 'normal', 'Thoradin', $1),
		 ('us/normal/hiddenhero', 'us', 'normal', 'Hiddenhero', $2)`, normalUser, hiddenUser); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from characters where key in ('us/normal/thoradin', 'us/normal/hiddenhero')`)
	})

	var guildID int64
	if err := pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Iron Vanguard') returning id`,
	).Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guilds where id = $1`, guildID) })

	verifiedAt := now.Add(-5 * 24 * time.Hour)
	if _, err := pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, verified_at) values
		 ($1, 'us/normal/thoradin', $2, $4),
		 ($1, 'us/normal/o''malley', $3, $4),
		 ($1, 'us/normal/notyet', $2, null)`,
		guildID, normalUser, memberUser, verifiedAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from guild_characters where guild_id = $1`, guildID) })

	reports := []struct {
		id, visibility string
		guildID        *int64
		createdAt      time.Time
	}{
		{"golden-pub-1", "public", &guildID, now.Add(-2 * 24 * time.Hour)},
		{"golden-pub-2", "public", &guildID, now.Add(-10 * 24 * time.Hour)},
		{"golden-eu-1", "public", nil, now.Add(-3 * 24 * time.Hour)},
		{"golden-private-1", "private", nil, now.Add(-1 * 24 * time.Hour)},
	}
	for _, r := range reports {
		if _, err := pool.Exec(ctx,
			`insert into reports (id, guild_id, visibility, status, created_at) values ($1, $2, $3, 'complete', $4)`,
			r.id, r.guildID, r.visibility, r.createdAt); err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from reports where id in ('golden-pub-1', 'golden-pub-2', 'golden-eu-1', 'golden-private-1')`)
	})

	if _, err := pool.Exec(ctx,
		`insert into fights (report_id, fight_index, encounter_id, kill) values
		 ('golden-pub-1', 1, 100, true), ('golden-pub-1', 2, 101, false),
		 ('golden-pub-2', 1, 100, true), ('golden-pub-2', 2, 102, true)`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from fights where report_id in ('golden-pub-1', 'golden-pub-2')`) })

	thoradinA := `[{"name":"output","score":90,"excluded":false},{"name":"survival","score":70,"excluded":false},` +
		`{"name":"mechanics","score":85,"excluded":false},{"name":"utility","score":60,"excluded":false},` +
		`{"name":"preparation","score":95,"excluded":false},{"name":"activity","score":null,"excluded":true}]`
	thoradinB := `[{"name":"output","score":92,"excluded":false},{"name":"survival","score":74,"excluded":false},` +
		`{"name":"mechanics","score":83,"excluded":false},{"name":"utility","score":64,"excluded":false},` +
		`{"name":"preparation","score":93,"excluded":false},{"name":"activity","score":91,"excluded":false}]`
	mork := `[{"name":"output","score":70,"excluded":false},{"name":"survival","score":null,"excluded":true},` +
		`{"name":"mechanics","score":70,"excluded":false},{"name":"utility","score":null,"excluded":true},` +
		`{"name":"preparation","score":null,"excluded":true},{"name":"activity","score":null,"excluded":true}]`

	insertScore := func(reportID string, index int, playerKey string, overall float64, components string, foughtAt time.Time) {
		// rating_scores is partitioned by month on fought_at (migration
		// 0022); ensure the partition exists before inserting into it, the
		// same way rating.Store.RateFight does before its own insert. This
		// fixture spans two calendar months (the in-window fights in
		// September 2026, the deliberately-out-of-window fight about 100
		// days back in June 2026), so this must run per insert, not once.
		if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
			t.Fatal(err)
		}
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
			 values ($1, $2, $3, $3, 100, false, $4, $4, false, $5, 'test', $6)`,
			reportID, index, playerKey, overall, components, foughtAt); err != nil {
			t.Fatal(err)
		}
	}
	insertScore("golden-pub-1", 1, "us/normal/thoradin", 80, thoradinA, now.Add(-2*24*time.Hour))
	insertScore("golden-pub-2", 1, "us/normal/thoradin", 82, thoradinB, now.Add(-10*24*time.Hour))
	insertScore("golden-pub-1", 3, "us/normal/thoradin", 999, thoradinB, now.Add(-100*24*time.Hour)) // out of window
	insertScore("golden-private-1", 1, "us/normal/thoradin", 999, thoradinB, now.Add(-1*24*time.Hour)) // private report
	insertScore("golden-pub-1", 4, "us/normal/hiddenhero", 999, thoradinB, now.Add(-2*24*time.Hour))   // anonymized
	insertScore("golden-eu-1", 1, "eu/pvp/mörk", 88, mork, now.Add(-3*24*time.Hour))
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id in
			('golden-pub-1', 'golden-pub-2', 'golden-eu-1', 'golden-private-1')`)
	})

	result, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: now})
	if err != nil {
		t.Fatal(err)
	}
	if result.Characters != 2 || result.Guilds != 1 {
		t.Fatalf("result = %+v, want 2 characters and 1 guild", result)
	}

	want, err := os.ReadFile(goldenFixturePath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(result.Rendered["Data.lua"], want) {
		t.Errorf("Data.lua does not match the checked-in golden fixture.\ngot:\n%s\nwant:\n%s", result.Rendered["Data.lua"], want)
	}
	assertValidLua(t, result.Rendered["Data.lua"])
}

// assertValidLua runs the rendered file through the `lua` interpreter when
// one is on PATH (dispatch: "run it through the lua interpreter if present
// in the test environment"), asserting it parses and defines the global
// with no runtime error. Without one, it falls back to a minimal
// structural check and says so, per the dispatch's own fallback.
func assertValidLua(t *testing.T, data []byte) {
	t.Helper()
	path, err := exec.LookPath("lua")
	if err != nil {
		t.Log("no `lua` interpreter on PATH; falling back to a minimal structural check")
		if !bytes.Contains(data, []byte("ForeverSixtyData = {")) || !bytes.Contains(data, []byte("\n}")) {
			t.Errorf("does not look like a well-formed ForeverSixtyData table: %s", data)
		}
		return
	}
	tmp := filepath.Join(t.TempDir(), "check.lua")
	script := string(data) + "\nassert(type(ForeverSixtyData) == \"table\" and ForeverSixtyData.format == 1)\n"
	if err := os.WriteFile(tmp, []byte(script), 0o600); err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command(path, tmp).CombinedOutput(); err != nil {
		t.Errorf("lua rejected the rendered file: %v\n%s", err, out)
	}
}
```

- [ ] **Step 4: Run to verify pass**

Run: `cd api && TEST_DATABASE_URL='postgres://forever:forever@localhost:5434/dataaddon_test?sslmode=disable' go test ./internal/dataaddon/... -run TestGoldenDataLuaIsByteExact -v -p 1`
Expected: PASS. If it fails on the `bytes.Equal` comparison, the failure message prints both sides — compare byte-for-byte against Step 1's fixture file and fix whichever of the two has the typo (most likely culprit: a trailing-comma or spacing mismatch in the hand-written fixture from Step 1).

Then run the whole package once more: `cd api && go test ./internal/dataaddon/... -v -p 1` — expected: PASS, every test in the package.

- [ ] **Step 5: Write the busted spec that reads the same fixture through the real reader**

```lua
-- addon/tests/data_addon_spec.lua
-- The nightly data-addon job's own golden fixture
-- (tests/fixtures/data_addon_sample.lua, checked in lockstep with
-- api/internal/dataaddon's TestGoldenDataLuaIsByteExact) loaded exactly as
-- a packaged Data.lua would be, then read back through the real reader.
local helper = require("spec_helper")
local mock = require("wow_mock")

describe("the data-addon job's output, read by Ratings.lua", function()
	local Ratings

	local function start(realm, region)
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, realm = realm, region = region })
		_G.ForeverSixtyData = nil
		dofile("tests/fixtures/data_addon_sample.lua")
		Ratings = helper.load("Ratings")
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	it("is available", function()
		start("Normal", 1) -- Export.REGION_NAMES[1] = "US"
		assert.is_true(Ratings.status().available)
	end)

	it("reads Thoradin's card with every component present", function()
		start("Normal", 1)
		local card = Ratings.forCharacter("Thoradin")
		assert.are.equal(81, card.rating)
		assert.are.equal(2, card.fights)
		local byKey = {}
		for _, c in ipairs(card.components) do
			byKey[c.key] = c.score
		end
		assert.are.equal(91, byKey.output)
		assert.are.equal(72, byKey.survival)
		assert.are.equal(84, byKey.mechanics)
		assert.are.equal(62, byKey.utility)
		assert.are.equal(94, byKey.preparation)
		assert.are.equal(91, byKey.activity)
	end)

	it("reads a non-ASCII name on another region/ruleset with only its computable components", function()
		start("Pvp", 3) -- Export.REGION_NAMES[3] = "EU"
		local card = Ratings.forCharacter("Mörk")
		assert.are.equal(88, card.rating)
		assert.are.equal(1, card.fights)
		local byKey = {}
		for _, c in ipairs(card.components) do
			byKey[c.key] = c.score
		end
		assert.are.equal(70, byKey.output)
		assert.are.equal(70, byKey.mechanics)
		assert.is_nil(byKey.survival)
	end)

	it("reads the guild's progress, nights, roster and members", function()
		start("Normal", 1)
		local guild = Ratings.forGuild("Iron Vanguard")
		assert.are.equal("2/3", guild.progress)
		assert.are.equal(2, guild.nights)
		assert.are.equal(2, guild.roster)
		assert.are.equal("o'malley", guild.members[1])
		assert.are.equal("thoradin", guild.members[2])
	end)

	it("never sees the anonymized character or the unverified guild member", function()
		start("Normal", 1)
		assert.is_nil(Ratings.forCharacter("Hiddenhero"))
		local guild = Ratings.forGuild("Iron Vanguard")
		for _, name in ipairs(guild.members) do
			assert.are_not.equal("notyet", name)
		end
	end)
end)
```

- [ ] **Step 6: Run the busted spec if `busted` is on PATH; otherwise validate with `luac`**

Run: `cd addon && (command -v busted >/dev/null && busted --output=TAP tests/data_addon_spec.lua) || luac -p tests/data_addon_spec.lua tests/fixtures/data_addon_sample.lua`
Expected: with `busted` present, every `it(...)` passes (TAP `ok` lines, no `not ok`); without it (this machine has `lua`/`luac` but not `busted`/`luarocks` fully set up for it, per Task 1's discovery — confirm with `command -v busted`), `luac -p` exits 0 on both files, proving they are at least syntactically valid Lua 5.1 (the CI job `verify` in `.github/workflows/addon-release.yml` already runs the real `busted` suite on every push, so this spec still gets exercised for real there even if this machine cannot run it now).

- [ ] **Step 7: Commit**

```bash
cd api && go vet ./internal/dataaddon/...
git add internal/dataaddon/golden_test.go internal/dataaddon/job.go ../addon/tests/fixtures/data_addon_sample.lua ../addon/tests/data_addon_spec.lua
printf 'test(api,addon): data-addon golden fixture, Lua validity check, and the lockstep busted spec\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > ../.superpowers/commit-msg-10.txt
git commit -F ../.superpowers/commit-msg-10.txt
```

---

### Task 11: `addon/ForeverSixtyData/README.md` and a TOC comment correction

**Files:**
- Modify: `addon/ForeverSixtyData/README.md`
- Modify: `addon/ForeverSixtyData/Data.lua` (comment only — the shape itself is unchanged, per this lane's ownership rule)

**Interfaces:** none — documentation only.

- [ ] **Step 1: Read the current README and Data.lua**

Already read in full during research; `Data.lua`'s guild comment currently reads:
```lua
	-- ["<region>:<ruleset>:<guild-slug>"] = { name, progress, nights, roster }
```
This is missing `members` (the dispatch's own shape adds it: `roster` is the verified-member count, `members` is the sorted name list — see the plan's aggregation-rule section for how the two reconcile).

- [ ] **Step 2: Fix the comment (structure of the checked-in empty tables is untouched)**

```lua
	-- ["<region>:<ruleset>:<guild-slug>"] = { name, progress, nights, roster, members = { "<name-slug>", ... } }
```

- [ ] **Step 3: Add one paragraph to the README naming the nightly job**

Append to `addon/ForeverSixtyData/README.md`:

```markdown

## Where Data.lua comes from

`api/internal/dataaddon` (the `api` module's `data-addon` Cloud Run job) reads the database
nightly, aggregates every public rated character's last 90 days and every guild's verified
roster, and writes the real `Data.lua` to a Cloud Storage bucket. `addon-data-release.yml`
downloads it and runs this package through the same release pipeline `addon-release.yml`
uses for the main addon. See `api/README.md`'s "The data-addon job" section for the
production setup, and `docs/superpowers/plans/2026-09-21-data-addon.md` for the exact
aggregation rule and the region/ruleset-as-realm key mapping.
```

- [ ] **Step 4: Commit**

```bash
git add addon/ForeverSixtyData/README.md addon/ForeverSixtyData/Data.lua
printf 'docs(addon): name the nightly data-addon job and fix the guild-row comment\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-11.txt
git commit -F .superpowers/commit-msg-11.txt
```

---

### Task 12: `.github/workflows/addon-data-release.yml`

**Files:**
- Create: `.github/workflows/addon-data-release.yml`

**Interfaces:** none — CI only. Reuses the same `google-github-actions/auth@v2` pattern `.github/workflows/api.yml`'s `deploy` job already uses (`vars.GCP_WIF_PROVIDER`, `vars.GCP_DEPLOYER_SA`), and the same "give the packager its own checkout" trick `.github/workflows/addon-release.yml`'s `release` job already uses.

- [ ] **Step 1: Write the workflow**

```yaml
# .github/workflows/addon-data-release.yml
# Nightly: download the data-addon Cloud Run job's own output (published
# to gs://<DATA_ADDON_BUCKET>/data-addon/...) and publish it through the
# same BigWigsMods packager addon-release.yml uses for the main addon.
# Scheduled an hour after the data-addon job's own 04:00 UTC Cloud
# Scheduler run (api/README.md's "The data-addon job") so the bucket
# object is fresh before this runs.
name: addon-data-release
on:
  schedule:
    - cron: '0 5 * * *'
  workflow_dispatch:
    inputs:
      dry_run:
        description: 'Build the zip and publish nothing'
        type: boolean
        default: true
permissions:
  contents: read
concurrency:
  group: addon-data-release
  cancel-in-progress: false
jobs:
  release:
    # Skipped until the bucket variable exists (api/README.md "The
    # data-addon job" names exactly what the owner creates), the same gate
    # api.yml's own deploy job uses for GCP_WIF_PROVIDER.
    if: vars.DATA_ADDON_BUCKET != ''
    runs-on: ubuntu-latest
    permissions:
      contents: write
      id-token: write
    steps:
      - uses: actions/checkout@v4

      - uses: google-github-actions/auth@v2
        with:
          workload_identity_provider: ${{ vars.GCP_WIF_PROVIDER }}
          service_account: ${{ vars.GCP_DEPLOYER_SA }}
      - uses: google-github-actions/setup-gcloud@v2

      - name: the nightly job's manifest names what it published
        id: manifest
        run: |
          set -eu -o pipefail
          gcloud storage cat "gs://${{ vars.DATA_ADDON_BUCKET }}/data-addon/manifest.txt" > /tmp/manifest.txt
          test -s /tmp/manifest.txt || { echo "manifest.txt named no files"; exit 1; }
          echo "files=$(paste -sd, /tmp/manifest.txt)" >> "$GITHUB_OUTPUT"

      - name: download every file the manifest names
        run: |
          set -eu -o pipefail
          while IFS= read -r name; do
            [ -n "$name" ] || continue
            gcloud storage cp "gs://${{ vars.DATA_ADDON_BUCKET }}/data-addon/$name" "addon/ForeverSixtyData/$name"
          done < /tmp/manifest.txt

      - uses: leafo/gh-actions-lua@v10
        with: { luaVersion: '5.1.5' }
      - name: every published file parses on the client's own interpreter
        run: |
          set -eu -o pipefail
          for f in addon/ForeverSixtyData/*.lua; do luac -p "$f"; done

      - name: today's tag
        id: tag
        run: echo "tag=data-v$(date -u +%Y.%m.%d)" >> "$GITHUB_OUTPUT"

      # Same trick as addon-release.yml's own "give the packager its own
      # checkout" step, for the same reason: release.sh requires its
      # topdir to be a VCS checkout root, and this is a monorepo
      # subdirectory. This step also folds the manifest's file list into
      # the TOC when it names more than just Data.lua (a split night),
      # entirely inside this throwaway checkout -- the real repository's
      # checked-in TOC is never touched.
      - name: give the packager its own checkout of addon/ForeverSixtyData
        id: pkgroot
        run: |
          set -eu -o pipefail
          pkgroot="$RUNNER_TEMP/ForeverSixtyData-pkgroot"
          rm -rf "$pkgroot"
          mkdir -p "$pkgroot"
          git archive HEAD -- addon/ForeverSixtyData | tar -x -C "$pkgroot" --strip-components=2
          cp addon/ForeverSixtyData/*.lua "$pkgroot/"
          if [ "${{ steps.manifest.outputs.files }}" != "Data.lua" ]; then
            awk 'NR==FNR{n++; line[n]=$0; next} /^Data\.lua[[:space:]]*$/{for (i=1;i<=n;i++) print line[i]; next} {print}' \
              /tmp/manifest.txt "$pkgroot/ForeverSixtyData.toc" > /tmp/toc.new
            mv /tmp/toc.new "$pkgroot/ForeverSixtyData.toc"
          fi
          git -C "$pkgroot" init -q
          origin=$(git remote get-url origin)
          git -C "$pkgroot" remote add origin "$origin"
          git -C "$pkgroot" -c user.name=github-actions[bot] \
            -c user.email=41898282+github-actions[bot]@users.noreply.github.com \
            add -A
          git -C "$pkgroot" -c user.name=github-actions[bot] \
            -c user.email=41898282+github-actions[bot]@users.noreply.github.com \
            commit -q -m "addon-data-release packaging snapshot ($GITHUB_SHA)"
          git -C "$pkgroot" tag "${{ steps.tag.outputs.tag }}"
          echo "dir=$pkgroot" >> "$GITHUB_OUTPUT"

      - name: which publish targets are configured
        id: targets
        env:
          CF_API_KEY: ${{ secrets.CF_API_KEY }}
          WAGO_API_TOKEN: ${{ secrets.WAGO_API_TOKEN }}
          DRY_RUN: ${{ inputs.dry_run }}
        run: |
          if [ "$DRY_RUN" = "true" ]; then
            echo "publish=no" >> "$GITHUB_OUTPUT"
            echo "::notice::Dry run requested; the zip is built and nothing is published."
          elif [ "$CF_API_KEY" != "" ] || [ "$WAGO_API_TOKEN" != "" ]; then
            echo "publish=yes" >> "$GITHUB_OUTPUT"
          else
            echo "publish=no" >> "$GITHUB_OUTPUT"
            echo "::notice::Neither CF_API_KEY nor WAGO_API_TOKEN is set; the zip is built and nothing is published."
          fi

      - name: build the zip
        uses: BigWigsMods/packager@v2
        with:
          args: -d -t ${{ steps.pkgroot.outputs.dir }}

      - name: what the zip contains
        run: |
          set -eu -o pipefail
          mapfile -t zips < <(find "${{ steps.pkgroot.outputs.dir }}/.release" -name '*.zip')
          if [ "${#zips[@]}" -ne 1 ]; then
            echo "expected exactly one zip in .release, found ${#zips[@]}:"
            printf '%s\n' "${zips[@]}"
            exit 1
          fi
          zip="${zips[0]}"
          top=$(unzip -Z1 "$zip" | cut -d/ -f1 | sort -u)
          if [ "$top" != "ForeverSixtyData" ]; then echo "top-level folder is '$top'"; exit 1; fi
          unzip -Z1 "$zip" | grep -q 'ForeverSixtyData/Data.lua' || { echo "Data.lua is missing"; exit 1; }
          echo "zip=$zip" >> "$GITHUB_ENV"

      - name: attach it to the release
        if: github.event_name == 'schedule' || (github.event_name == 'workflow_dispatch' && inputs.dry_run != true)
        uses: softprops/action-gh-release@v2
        with:
          tag_name: ${{ steps.tag.outputs.tag }}
          files: ${{ env.zip }}
          generate_release_notes: true

      - name: publish to CurseForge and Wago Addons
        if: steps.targets.outputs.publish == 'yes' && (github.event_name == 'schedule' || (github.event_name == 'workflow_dispatch' && inputs.dry_run != true))
        uses: BigWigsMods/packager@v2
        env:
          CF_API_KEY: ${{ secrets.CF_API_KEY }}
          WAGO_API_TOKEN: ${{ secrets.WAGO_API_TOKEN }}
        with:
          args: -t ${{ steps.pkgroot.outputs.dir }}
```

- [ ] **Step 2: Validate the YAML parses**

Run: `python3 -c "import yaml, sys; yaml.safe_load(open('.github/workflows/addon-data-release.yml'))" 2>&1 || python3 -c "import json,sys; import yaml" 2>&1`

If `pyyaml` is not installed, fall back to: `ruby -ryaml -e "YAML.load_file('.github/workflows/addon-data-release.yml')"` or, minimally, `cat .github/workflows/addon-data-release.yml | head -1` plus a careful re-read of the indentation by eye. Expected: no parse error.

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/addon-data-release.yml
printf 'ci(addon): nightly workflow packages and publishes the data-addon\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-12.txt
git commit -F .superpowers/commit-msg-12.txt
```

---

### Task 13: `api/README.md` — the production setup recipe

**Files:**
- Modify: `api/README.md` (append a new subsection near "The performance ratings lane adds `rating-backfill`" in "First-time setup," plus one row-worthy mention where jobs are listed for `.github/workflows/api.yml`'s image-repoint loop — that loop itself is in `api/internal` ownership territory only insofar as it's a workflow file already outside this lane; note it as a follow-up instead of editing `api.yml`, which is not in this lane's ownership either)

**Interfaces:** none — documentation only.

- [ ] **Step 1: Append the job recipe**

Insert after the existing `rating-backfill` recipe block (ending `gcloud run jobs execute rating-backfill --region us-east1 --wait`) in `api/README.md`:

```markdown

### The data-addon job

`data-addon` (`api/internal/dataaddon`) aggregates every public rated character's last 90
days and every verified guild's roster into the Forever Sixty Data addon's `Data.lua`, and
publishes it to a Cloud Storage bucket for `.github/workflows/addon-data-release.yml` to
package nightly. Unlike every other job in this file, its bucket address is read directly
via `os.Getenv("DATA_ADDON_BUCKET")` in `main.go` rather than through `config.Config` — see
`docs/superpowers/plans/2026-09-21-data-addon.md`'s Task 9 for why.

One-time setup (in addition to "1. Google Cloud project setup" above):

```bash
gcloud services enable storage.googleapis.com
gcloud storage buckets create gs://foreversixty-addon-data --location us-east1 --uniform-bucket-level-access

# api-runtime (the service's own service account) needs to write the nightly file:
gcloud storage buckets add-iam-policy-binding gs://foreversixty-addon-data \
  --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com \
  --role roles/storage.objectAdmin

# the deployer service account (see "Then set up Workload Identity Federation..." above)
# needs read access, since addon-data-release.yml downloads through the same WIF identity
# api.yml's deploy job already uses:
gcloud storage buckets add-iam-policy-binding gs://foreversixty-addon-data \
  --member serviceAccount:<deployer service account email> \
  --role roles/storage.objectViewer

gcloud run jobs create data-addon \
  --image us-east1-docker.pkg.dev/foreversixty/api/api:latest \
  --region us-east1 --args data-addon \
  --cpu 1 --memory 512Mi --task-timeout 10m \
  --set-env-vars "$(tr '\n' ',' < .env.job),DATA_ADDON_BUCKET=foreversixty-addon-data" \
  --service-account api-runtime@foreversixty.iam.gserviceaccount.com

gcloud scheduler jobs create http data-addon-nightly \
  --schedule "0 4 * * *" \
  --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/data-addon:run" \
  --http-method POST --oauth-service-account-email api-runtime@foreversixty.iam.gserviceaccount.com
```

Then, as a GitHub repository variable (Settings → Secrets and variables → Actions →
Variables — the same place `GCP_WIF_PROVIDER` and `GCP_DEPLOYER_SA` already live):

```
DATA_ADDON_BUCKET = foreversixty-addon-data
```

`addon-data-release.yml` is gated on this variable existing (`if: vars.DATA_ADDON_BUCKET != ''`),
the same way `api.yml`'s own `deploy` job is gated on `GCP_WIF_PROVIDER`.

Like `parse-report`, `sim-run` and `sim-validate`, `.github/workflows/api.yml`'s `deploy`
job's "Point the jobs at the new image" loop should add `data-addon` to its job list so a
new deploy repoints it too — that file is outside this lane's ownership, so this is a
follow-up for whichever lane next touches `api.yml`, not done here.
```

- [ ] **Step 2: Commit**

```bash
git add api/README.md
printf 'docs(api): production setup recipe for the data-addon job (bucket, IAM, scheduler)\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-13.txt
git commit -F .superpowers/commit-msg-13.txt
```

---

## Final review (per the lane's own instructions)

After Task 13, run `superpowers:subagent-driven-development`'s whole-branch review with one fix wave, covering at minimum:

1. **Code review:** every task's diff together — DRY (no duplicate SQL beyond the one documented, deliberate `progressionByGuild` duplication of `rankings.Store.Guild`'s query), file/function size limits, immutability, error wrapping.
2. **Security pass**, explicitly re-checking (dispatch: "the job reads private data and must publish only public-derived numbers; check the anonymized and visibility rules again there"):
   - `characterFights` joins `reports.visibility = 'public'` (never `unlisted`, `guild`, or `private`) and excludes any `player_key` whose account has `users.anonymize = true`.
   - `progressionByGuild` uses `reports.visibility <> 'private'` (matching the public guild page's own rule, which does show `unlisted` progression counts as an aggregate, never the report itself) — confirm this still cannot leak an unlisted report's *existence* or *title*, since `dataaddon` never reads or writes `title`, `zone`, or `id` for any report.
   - `guildsWithVerifiedMembers`/`verifiedMembers` only ever read `guild_characters` rows with `verified_at is not null` — an unverified claim (`source = 'invite'`/`'claim'` with no corroboration) never reaches `Data.lua`.
   - No secret, credential, or raw database connection string is ever logged (`job.go`'s `Logger` calls only ever carry counts, filenames, byte sizes, and player_key/guild-id strings that are already public once published).
   - `GCS.Upload` never logs its own object bytes or bucket credentials.
3. Full suite once more from `api/`: `TEST_DATABASE_URL='postgres://forever:forever@localhost:5434/dataaddon_test?sslmode=disable' go test -race -cover -p 1 ./internal/dataaddon/... && go build ./... && go vet ./...`.
4. `cd addon && luac -p ForeverSixty/*.lua ForeverSixtyData/Data.lua tests/data_addon_spec.lua tests/fixtures/data_addon_sample.lua` — every file this lane touched still parses on Lua 5.1.

Write the ledger to `.superpowers/sdd/2026-09-21-data-addon/progress.md` per the lane's own instructions, with a `Ruling:` line for every judgment call this plan already made explicit (the ruleset-as-realm key mapping, the `roundHalfUp`-everywhere formatting choice, the "raid night = distinct UTC date" definition, the `DATA_ADDON_BUCKET` read-via-`os.Getenv` deviation from `config.Config`, and the `/ratings` page aggregation discrepancy).
