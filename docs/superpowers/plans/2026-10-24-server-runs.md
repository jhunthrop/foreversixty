# Server runs: "Your server runs" and "done while away" Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A payer who closes the tab still gets their server sims: a `Your server runs` list on the sim page (running with an estimate, done, failed with the reason and `Run again`) and one `Your <spec> sim finished: <dps> DPS` line in the character band on the next page load, dismissed by a click or by opening the result.

**Architecture:** The `sims` table gains `finished_at`, `failure_reason` and `seen_at` (migration 0045); the job and the run handler write them. `GET /v1/sims?mine=1&lane=server` lists server runs with the per-spec estimate; `POST /v1/sims/{id}/seen` marks one seen. `GET /v1/me` carries `finished_sims` (unseen done runs, newest three) and `server_sims.runs_this_month`. The web renders the list above the landing state on `/sim`, the done line in `CharacterHeader`, and the count in the account Premium block.

**Tech Stack:** Go (pgx v5, net/http mux), Svelte 5 islands, Astro, vitest, Playwright, the `test-support/sim-api.ts` fake for e2e.

**Spec:** Product: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` §4.3, §5, §6. Experience: `design/specs/2026-10-24-server-runs.md` (ux-designer, 2026-10-24); mock `design/mocks/gen_server_runs.py`. The experience spec's copy and states are the authority for every label below; where this plan quotes copy, the spec wins.

## Global Constraints

- Only the deployed API applies migrations; the lane adds `0045_*.up.sql`/`.down.sql`.
- No queue position, no cancel control, no progress percentage anywhere (product spec §4.3 rulings). Estimates read `about N minutes` (minimum `about a minute`), rounded to the minute, from the median of the last ten finished server runs for that spec; absent with fewer than three.
- Every write of `sims` state stays in `sims.Store` (`Queue`, `Advance`, `Finish`, `Fail`); the web never infers state from timing.
- `GET /v1/me` stays `httpx.CachePrivate`; the web reads it through `fetchMeOnce` and invalidates with `invalidateMe` after a mutation.
- Copy lives in the copy modules (`web/src/lib/sim/copy.ts`, `web/src/lib/billing/copy.ts`), verbatim from the experience spec; numbers render from data.
- Tests: API packages with `TEST_DATABASE_URL` and `-p 1`; web `npm run check`, vitest for touched files, the e2e specs this lane adds plus `sim-run.spec.ts`, `sim-history.spec.ts`, `account-billing.spec.ts`, `header-alignment.spec.ts`; the shared header changes, so the lane close runs the full web vitest and full Playwright suites (memory: shared layout needs the full suite).
- Commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.

## Review Focus

1. A payer with a run that finished while no tab was open: the next load of home shows the done line once; opening `/sim/<id>` marks it seen and the line is gone on the following load (Task 3 API test, Task 6 e2e).
2. A failed run: the row shows the plain-words reason from the job, and `Run again` re-submits the same request (Task 2 `failure_reason`, Task 5).
3. A running row for a spec with fewer than three finished runs shows no estimate and never a percentage (Task 2 estimate test, Task 5 unit test).
4. A free member (no `server_sims`) sees neither the list nor the done line, and `GET /v1/sims?mine=1&lane=server` still answers their own server rows if any exist (an expired payer keeps history) (Task 2, Task 5).
5. Three or more unseen finished runs: `/v1/me` carries only the newest three and the band shows one line for the newest with `and N more` per the spec (Task 3, Task 4).

---

### Task 1: Migration 0045 and the store's new columns

**Files:**
- Create: `api/internal/db/migrations/0045_sim_run_lifecycle.up.sql`, `.down.sql`
- Modify: `api/internal/sims/store.go` (`Row`, `Mine`, `Finish`, `Fail`, new `MarkSeen`, `Estimate`, `FinishedUnseen`, `RunsThisMonth`)
- Modify: `api/internal/sims/job.go:124-137` (`failed` passes the reason through; already does — verify `Fail` now stores it)
- Test: `api/internal/sims/store_test.go`

**Interfaces:**
- Produces:
  ```go
  // Row gains:
  Lane           string     `json:"lane"`
  FinishedAt     *time.Time `json:"finished_at,omitempty"`
  FailureReason  string     `json:"failure_reason,omitempty"`
  Seen           bool       `json:"seen"`                       // seen_at is not null (experience spec §10, decision 10)
  EstimateSeconds *int      `json:"estimate_seconds,omitempty"` // running/queued rows only
  func (s *Store) Mine(ctx, userID int64, page int, kind, lane string) (Page, error) // lane "" | "server" | "browser"
  func (s *Store) MarkSeen(ctx, id string, userID int64) (bool, error)             // false when not the caller's row or not done
  func (s *Store) FailedRequest(ctx, id string, userID int64) (simapi.SimRequest, error) // the stored request of the caller's own error row; ErrNotFound otherwise
  type FinishedRun struct { SimID, Spec, Title string; DPS float64; FinishedAt time.Time }
  func (s *Store) FinishedUnseen(ctx, userID int64, limit int) ([]FinishedRun, error)
  func (s *Store) RunsThisMonth(ctx, userID int64) (int, error)                      // lane=server, created_at >= date_trunc('month', now())
  func (s *Store) EstimateSeconds(ctx, spec string) (*int, error)                      // median duration_ms/1000 of the last 10 done server runs; nil under 3
  ```

- [ ] **Step 1: Migration**
```sql
alter table sims add column if not exists finished_at    timestamptz;
alter table sims add column if not exists failure_reason text not null default '';
alter table sims add column if not exists seen_at        timestamptz;
-- "Done while away": the member's unseen finished server runs, newest first.
create index if not exists sims_user_unseen_idx on sims (user_id, finished_at desc)
  where lane = 'server' and state = 'done' and seen_at is null;
```
Down drops the index and the three columns.

- [ ] **Step 2: Failing store tests** — `TestFinishSetsFinishedAt`, `TestFailStoresThePlainReasonAndFinishedAt`, `TestMineFiltersByLaneAndCarriesLifecycleFields` (a browser save and a server run; `lane=server` returns one row with `lane`, `finished_at`, `failure_reason`, `seen` as written), `TestMarkSeenOnlyForTheOwnersDoneRun` (owner+done → true and `seen_at` set; another user → false; a running row → false), `TestFailedRequestReturnsTheOwnersErrorRowOnly` (an error row's stored request round-trips; a done row, another user's row and an unknown id → `ErrNotFound`), `TestFinishedUnseenNewestFirstAndCapped` (four done unseen → three newest; seen one excluded), `TestRunsThisMonthCountsServerRunsOnly` (one browser save, two server runs this month, one server run last month via a `created_at` update → 2), `TestEstimateSecondsIsTheMedianOfTheLastTenDoneRuns` (two done → nil; three with `duration_ms` 60000/120000/600000 → 120; eleven rows → the oldest ignored).
- [ ] **Step 3: Run, expect FAIL.** `cd api && TEST_DATABASE_URL=... go test -p 1 ./internal/sims/ -run 'Finish|Fail|Mine|MarkSeen|FinishedUnseen|RunsThisMonth|Estimate'`
- [ ] **Step 4: Implement.** `Finish` and `Fail` set `finished_at = now()`; `Fail` writes `failure_reason = $why`. `Mine` adds the lane predicate (keep the one-predicate-two-queries pattern) and selects `lane, finished_at, failure_reason`; estimates are attached by the handler (Task 2), not the store query. `EstimateSeconds` reads `(result->>'duration_ms')::bigint` from the last ten `state='done' and lane='server' and spec=$1` rows ordered `finished_at desc nulls last, created_at desc` and takes the median in Go.
- [ ] **Step 5: Run, expect PASS; `go vet ./internal/sims/`.**
- [ ] **Step 6: Commit** `feat(sims): a run remembers when it finished, why it failed and whether its owner has seen it`

---

### Task 2: The list, seen and rerun routes

**Files:**
- Modify: `api/internal/sims/handler.go` (`Mount`, `mine`, new `seen`), `api/internal/sims/run.go` (extract `startRun(ctx, w, r, actor, req) ` from `run` so `rerun` shares the entitlement check, caps, size checks, `Queue` and `Jobs.Run`; new `rerun`), `api/internal/sims/store.go` only if a query is missing
- Modify: `api/openapi.yaml` (the `lane` query parameter on `GET /v1/sims`, the new `POST /v1/sims/{id}/seen`, the `Row` fields; `api/internal/server/openapi_test.go` checks mounted routes against it) and `api/README.md` where `/v1/sims` is listed
- Test: `api/internal/sims/handler_test.go`

**Interfaces:**
- Consumes: Task 1.
- Produces: `GET /v1/sims?mine=1&lane=server[&kind=…]` → `Page` of `Row` with `estimate_seconds` on queued/running rows (one `EstimateSeconds` call per distinct spec on the page); `lane` other than `server`/`browser`/empty → 400 `invalid`. `POST /v1/sims/{id}/seen` (session required) → 204 when marked, 404 `not_found` when not the caller's done row. `POST /v1/sims/{id}/rerun` (session required, mounted with the same `simsRunPerMinute` limiter and the same `Jobs/Accounts/Planner` guard as `run`) → 202 `{sim_id}` for the caller's own `error` row, re-running its stored request through the exact checks `run` applies (402 `premium_required`, `cap_exceeded`, `too_large`, 502 `upstream`); 404 `not_found` for a row that is not the caller's or not failed.

- [ ] **Step 1: Failing handler tests** — `TestMineLaneServerListsOnlyServerRunsWithEstimates` (seed three done server runs for the spec via `Queue`+`Finish` with `duration_ms`, one running; the running row carries `estimate_seconds`; done rows do not and carry `seen`; a browser save is absent), `TestMineRejectsAnUnknownLane`, `TestSeenMarksTheOwnersDoneRunAndRefusesOthers` (204 then `seen_at` set; a second user's row → 404; anonymous → 401), `TestRerunStartsTheFailedRunsStoredRequest` (a failed row → 202 with a new id; `jobs.Fake` ran `sim-run` for the new id; the new row's stored request equals the old one; the old row is unchanged), `TestRerunRefusesADoneRowAndAnotherUsersRow` (404 both), `TestRerunIsGatedByTheEntitlement` (fake entitlements false → 402).
- [ ] **Step 2: Run, expect FAIL.**
- [ ] **Step 3: Implement**: `mine` parses `lane` with the same refuse-unknown pattern as `kind`; after `Store.Mine`, for rows in `queued`/`running` fill `EstimateSeconds` memoised per spec. `seen` reads the actor, calls `MarkSeen`, writes 204 / 404. `run` becomes decode + `startRun`; `rerun` loads `FailedRequest` and calls `startRun` with it (the stored request already carries the engine pin and encounter defaults; `startRun` re-applies them anyway). Mount both under `auth.RequireSession`; `rerun` inside the `if s.Jobs != nil …` block with the limiter.
- [ ] **Step 4: Run, expect PASS; `go vet`.** Update `api/openapi.yaml` (`lane`, `seen`, `rerun`, the new `Row` fields) and the `/v1/sims` lines in `api/README.md`.
- [ ] **Step 5: Commit** `feat(sims): list server runs with estimates; mark one seen; run a failed one again`

---

### Task 3: `/v1/me` carries finished runs and this month's count

**Files:**
- Modify: `api/internal/auth/handler.go:340-420` (`Me`, `me`), `api/internal/auth/handler.go:47-71` (`Service` gains `Sims SimsView`)
- Modify: `api/cmd/api/main.go` (wire `sims.Store` into `auth.Service.Sims`)
- Test: `api/internal/auth/handler_test.go`

**Interfaces:**
- Consumes: Task 1's `FinishedUnseen`, `RunsThisMonth`.
- Produces:
  ```go
  // auth package, declared where used. sims already imports auth (RequireSession, ActorFrom),
  // so auth must NOT import sims: the interface and its row type are auth's own, and
  // api/cmd/api/main.go wires a two-method adapter over *sims.Store that converts
  // sims.FinishedRun into auth.FinishedSim.
  type FinishedSim struct { SimID, Spec, Title string; DPS float64; FinishedAt time.Time }
  type SimsView interface {
      FinishedUnseen(ctx context.Context, userID int64, limit int) ([]FinishedSim, error)
      RunsThisMonth(ctx context.Context, userID int64) (int, error)
  }
  // Me gains:
  FinishedSims []FinishedSim `json:"finished_sims"`           // always present, [] when none; newest first, max 3
  ServerSims   *ServerSimsView `json:"server_sims,omitempty"`  // nil when Sims is nil
  // FinishedSim marshals as sim_id, spec, title, dps, finished_at (RFC3339 via rfc3339Ptr's rule).
  type ServerSimsView struct { RunsThisMonth int `json:"runs_this_month"` }
  ```
  `finished_sims` is filled only when `Entitlements.ServerSims` is true (a lapsed payer's unseen rows stay unseen but unannounced; ruling below).

- [ ] **Step 1: Failing tests** — `TestMeListsUnseenFinishedServerRunsForAPayer` (grant premium; seed two done server runs via `sims.Store`, one seen → `finished_sims` has one with `sim_id`, `spec`, `dps`, `finished_at`; `server_sims.runs_this_month` == 2), `TestMeOmitsFinishedSimsForANonPayer` (same rows, no grant → `finished_sims` is `[]`, `server_sims.runs_this_month` still 2), `TestMeWorksWithNoSimsStoreConfigured` (`finished_sims: []`, no `server_sims` key).
- [ ] **Step 2: Run, expect FAIL.** **Step 3: Implement** (`Service.Sims SimsView`, nil-safe; the auth tests satisfy `SimsView` with a fake; `api/cmd/api/main.go:577` wires the adapter — put it in a small `api/cmd/api/simsview.go`, and note that the api binary builds). **Step 4: PASS, vet.** **Step 5: Commit** `feat(auth): /v1/me says which server runs finished while you were away`

---

### Task 4: Web types, copy and the seen/rerun client

**Files:**
- Modify: `web/src/lib/account/api.ts` (`Me` gains `finished_sims?: FinishedSim[]`, `server_sims?: { runs_this_month: number }`; new `FinishedSim` interface: `sim_id, spec, title, dps, finished_at`)
- Modify: `web/src/lib/sim/types.ts` (`SimListRow` gains `lane?`, `state?`, `finished_at?`, `failure_reason?`, `seen?`, `estimate_seconds?`)
- Modify: `web/src/lib/sim/api.ts` (`listMySims(page, apiBase, kind, lane?)` appends `&lane=`; new `markSimSeen(simId)` → `POST /v1/sims/{id}/seen` then `invalidateMe()`; new `rerunSim(simId)` → `POST /v1/sims/{id}/rerun` → `{sim_id}`; both through the existing `call`/CSRF helper the other POSTs use)
- Create: `web/src/lib/sim/server-runs-copy.ts` (every string of experience spec §6 for the list, the row states, `Run again`, the band line and its dismiss; functions take data: `runningLine(estimateSeconds, elapsedSeconds)`, `estimateWords(seconds)`, `failedLine(reason)`, `bandLine(specLabel, dps)`, `dismissLabel(specLabel)`, `moreLabel(n)`), `web/src/lib/sim/server-runs.ts` (view model: `rowState(row, now)` → `queued | running | running-overdue | running-no-estimate | done | done-unseen | error`; `estimateWords`: ≤ 1 min `about a minute`, 2–89 `about N minutes`, ≥ 90 `about N hours` rounded; overdue when elapsed > 1.5 × estimate; `ageWords(date, now)` relative inside 14 days else `Nov 25` via the header spec's rule — reuse `lib/dates` if it has one; `visibleRows(rows)` → newest five and `hasMore`)
- Modify: `web/src/lib/billing/copy.ts` (`premiumBlockCopy.rows.serverSims.on` becomes a function of `runs_this_month | undefined`: `On · 4 runs this month`, `On · 1 run this month`, `On · no runs yet this month`, `On` when undefined; `toLocaleString('en-US')`)
- Modify: `web/src/test-support/sim-api.ts` (built-in routes: `GET /v1/sims?mine=1&lane=server` → an empty page by default, `POST /v1/sims/{id}/seen` → 204, `POST /v1/sims/{id}/rerun` → 202 `{sim_id}`; recorded in `lastUrl`/`lastBody`)
- Test: `server-runs.test.ts`, `server-runs-copy.test.ts`, `api.test.ts` additions, `billing/copy.test.ts` additions

- [ ] Steps: failing vitest for every function above with the spec's examples (`estimateWords(60) === 'about a minute'`, `estimateWords(170) === 'about 3 minutes'`, `estimateWords(5400) === 'about 2 hours'`; overdue at elapsed 271 s with estimate 180; `visibleRows` of 7 → 5 + `hasMore`; the four account strings; `markSimSeen` posts and invalidates); implement; `npm run check` and the touched vitest files; commit `feat(web): server-run copy, view model and the seen/rerun client`.

### Task 5: The `Your server runs` list on `/sim`

**Files:**
- Create: `web/src/components/sim/ServerRuns.svelte` (experience spec §3.2–3.3, §4: the section, heading, panel, up to five rows via `visibleRows`, overflow footer, `EmptyState`, `LoadError` + `Try again`, skeleton reserve; one row = `ServerRunRow.svelte`), `web/src/components/sim/ServerRunRow.svelte` (grid `32px minmax(0,1.4fr) minmax(0,1.2fr) auto`, `ClassCrestRing` 32, unread dot with sr-only `Not opened yet`, state block, done rows are `<a href=/sim/{id}>` with `›`, `Run again` on error rows → `rerunSim`; button states pending `Starting…` / success `Started` / failure line / 402 line with `See plans`)
- Modify: `web/src/components/sim/SimView.svelte` (mount `ServerRuns` lazily above `SimRunBlock` when `effectiveServerSims(me)`; `loadServerRuns()` on session ready; re-read every 10 s while any row is queued/running and `document.visibilityState === 'visible'`; a row that flips to `done` while visible is marked seen at once via `markSimSeen`; when `store.serverRunning` ends, re-read; set `localStorage fs.serverRuns.live` (expiry 2 h) when a server run starts or a queued/running row is read, `fs.serverRuns.count` after every read)
- Modify: `web/src/pages/sim.astro` or the island shell (the reserve box: `html[data-session='1'][data-server-sims='1']` per §3.7, heights 64 desktop / 92 phone per row, `max(1, min(count, 5))`)
- Modify: the pre-paint inline script in `web/src/layouts/Base.astro:85-93` that stamps `data-session` on `<html>` so it also sets `data-server-sims='1'` from the cached `/v1/me` snapshot's `entitlements.server_sims`
- Modify: `web/src/components/account/AccountPremiumBlock.svelte` (the Server sims value from `me.server_sims?.runs_this_month` through the copy function; `Account.svelte` passes it)
- Test: `ServerRuns.test.ts`, `ServerRunRow.test.ts` (one test per row state's rendered strings and link/button presence; reduced-motion class; the 402 branch), `AccountPremiumBlock.test.ts` additions, `_sim.test.ts` for the reserve rule

- [ ] Steps: failing vitest per component state; implement; `npm run check`; vitest for touched files; e2e `sim-server-runs.spec.ts` written here (see Task 6 for its cases) and run; commit `feat(web): Your server runs on the sim page; the account row counts this month's runs`.

### Task 6: The done line in the character band and the saved-sim page's seen call

**Files:**
- Create: `web/src/components/character/FinishedSimLine.svelte` (§3.4: footer row inside the panel; `role="status"`; link to `/sim/<newest>` with the sentence, DPS in mono `--strong`, `Open result →` desktop / `→` phone; `and N more` → `/sim#server-runs`; 44×44 dismiss with `aria-label`; dismiss posts `markSimSeen` for every named run, `aria-busy` while in flight, failure text for 4 s; focus to the band's action on dismiss)
- Modify: `web/src/components/character/CharacterHeader.svelte` (the panel becomes a column: existing grid then the footer; render the line when `session.me.entitlements.server_sims && finished_sims.length > 0` and not dismissed locally; the pre-paint hint `html[data-finished='1']` reserves 44/56 px in the skeleton and ready states per §3.7; the 136 px minimum of the grid itself is unchanged), `web/src/styles/global.css` or the header's own styles for the hint reserve
- Modify: the `Base.astro` pre-paint script (same one as Task 5) to set `data-finished='1'` from the snapshot's `finished_sims.length > 0` or the `fs.serverRuns.live` flag
- Modify: `web/src/components/sim/SavedSim.svelte` (on mount, once, `markSimSeen(simId)` when `me` owns it: failures ignored)
- Test: `FinishedSimLine.test.ts` (one unseen; three → `and 2 more`; dismiss posts three seen calls and hides; failure text; phone variant drops `Open result`), `CharacterHeader.test.ts` additions (line present/absent; the hint reserve class), `SavedSim.test.ts` addition
- E2E (write and run here): `web/tests/e2e/sim-server-runs.spec.ts` — a payer with a running row (estimate → `takes about 3 minutes`), a running row without estimate, a done unseen row (dot, tinted), a done seen row, a failed row (`Failed: stopped before completion`, `Run again` → posts `/v1/sims/{id}/rerun`, button reads `Started`); a free member sees no list; the empty state; a Playwright CLS check of `/sim` at 390 and 1440 with the hints on (≤ 0.02). `web/tests/e2e/header-finished-sim.spec.ts` — home with one `finished_sims` → the line; dismiss posts seen and hides, band returns to its height; three → `and 2 more`; `[]` → no line; CLS check at 390 and 1440 hinted. `account-billing.spec.ts` gains the `On · 4 runs this month` assertion.
- Commit `feat(web): the band tells you a server sim finished while you were away`.

### Task 7: Changelog and lane close

- Changelog entry `web/src/content/changelog/2026-10-<day>-server-runs.md` (frontmatter as `2026-10-24-premium-page-truth.md`; past tense; dated the ship day).
- Lane close (controller): full web vitest + full Playwright + the lhci shard (the shared header changed); captures per experience spec §9 items 1–8 at 390/1440/2000 into `design/reviews/2026-10-24-server-runs/` with a README of verdicts; ux-designer and wow-player reviews on the captures; one fix round; merge by fast-forward; deploy; verify `/v1/me` on production carries `finished_sims` for the owner's account.

## Rulings made by the plan (owner to check)

- A lapsed payer keeps their rows in the list route (history is theirs) but gets no list and no done line on the web (`server_sims` false renders the free page byte for byte; `finished_sims` only for payers): the surfaces are a Premium benefit, the data is not.
- `seen` is the member's own action (open, dismiss, or watching a run finish in a visible tab); the API never auto-marks on the list read, so a glance at the list does not silence the band.
- Estimates come from the whole site's last ten finished runs for that spec, not the member's own.
- The experience spec's eleven open decisions (§11) are all accepted as written: the list above `SimRunBlock`; no separate line on `/sim`; dismiss marks every named run; `takes about N minutes` and `taking longer than usual`; the pre-paint hints with one accepted 44 px shift for an unhinted browser; the count words; five rows; the `rerun` route with the old row left reading `Started`; done runs in both lists; `seen` on rows; the e2e fails a run instead of cancelling.
- Running and failed rows are not links: `SavedSim.svelte` does not render a running run (checked: it has no progress handling).
