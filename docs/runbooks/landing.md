# Landing a batch

With more than one lane running, lanes never push to main: one integrator lands them
together. The full CI suite runs once, on a `land/<n>` branch, and main is fast-forwarded to
that exact commit when it is green. Target: 30 to 45 minutes from merge to main.

1. Merge the lanes' branches onto `origin/main`, make the one engine pin, regenerate once and
   run the fast local checks (lint, typecheck, unit tests, the touched e2e specs, apl-check).
   No full Playwright, lhci, fork suite or `-race` sweeps locally: CI runs them.
2. Push the batch to a new land branch: `git push origin HEAD:land/<n>`.
3. Run `tools/land.sh land/<n>`. It:
   - waits for every check workflow GitHub started for the branch's head sha (web with every
     e2e shard, real-engine and lhci; sim; api in four test shards; data; tools; addon; logs;
     coach; companion; ios; android; mobile-codegen; whichever the batch's paths reach);
   - stops at the first failed job and lists every failed job with its link (exit 1);
   - when all are green, runs `git push origin <sha>:main`, which only succeeds as a
     fast-forward. If main moved meanwhile it says so (exit 2): merge `origin/main` into the
     batch, push it as a new `land/<n+1>` and land that;
   - deletes the land branch.
4. The push to main starts main's own runs and the deploys (web, api, the companion
   release). Watch them finish before the next landing.

What runs where:

- Every check workflow triggers on `land/**` pushes with the same path filters as main. A
  new branch is diffed against main, so the batch's whole diff decides what runs. GitHub
  reads at most 300 changed files for a path filter; a batch past that can miss a workflow,
  so keep regenerated trees out of a batch that does not need them.
- No deploy or release runs on `land/**`: web's and api's `deploy` jobs and the companion's
  `version`/`build`/`release` run only on main (or a tag).
- A `land/**` run is grouped by its sha and never cancelled, so nothing a landing waits on is
  lost to a later push. Main's runs are never cancelled either. Branch and pull request runs
  still cancel their superseded twin.
- The api's go test is split across four parallel jobs by `tools/go-test-shard.sh`, each with
  its own Postgres and `-race -p 1`; weights in `api/test-weights.txt`. Refresh the weights
  from a recent log when one shard runs much longer than the others.

- The slowest step under load was `actions/checkout`: the tree is ~3.5 GB and a cold checkout
  took 5 to 16 minutes when a landing started ~35 jobs at once (2026-11-03). No check job has
  a timeout under 20 minutes for that reason. See "Checkout" below.

## Checkout

Every check job that runs on a `land/**` push checks out only what it reads:

- A non-cone sparse checkout (`/*` and then `!/<tree>/` lines) leaves out the heavy trees no
  check reads: `data/diffs`, `design/reviews` (~0.9 GB of review captures; tests only ever
  write there), `design/mocks` (the api keeps `design/mocks/2026-10-29-discord-gear-panel/`,
  which a test reads), `docs/superpowers/reviews`, and `apps/ios` and `apps/android` in every
  job but the apps' own, `mobile-codegen` and web `build` (its unit suite holds the codegen
  drift checks). Everything else, small trees included, stays. A sparse checkout is also a
  `blob:none` partial clone, so the left-out trees cost nothing to fetch.
- `data/builds` (~2 GB on disk, every build, since several suites walk all of them) is left
  out of the checkout too and put back by `.github/actions/data-builds`: a restore from the
  Actions cache keyed by the tree's git id (`git rev-parse HEAD:data/builds`, ~150 MB
  compressed), then `git sparse-checkout add` and a `git status` that fails the job unless
  the tree is exactly the commit's. A miss checks it out from git. Only main saves the
  cache: a `land/**` branch reads main's caches and no other branch's, so a batch whose
  data/builds matches main's last pushed tree restores it, and one that moved it pays the git
  checkout once.
- Jobs that never read data/builds (addon, companion test, tools, logs bench) skip it.
- ios keeps the full checkout: its tests read android's test resources, and on the macOS
  runner the sparse partial clone measured slower than the full one (386 s against 78 s).
- companion's `changes` needs the history for its merge base and no file: `fetch-depth: 0`
  with `filter: blob:none`.

A job that starts reading one of the left-out trees adds it back in its own checkout (a
`/<tree>/` line after the `!` lines re-includes it). data.yml is held, and deploy and release
jobs, which no landing runs, keep the full checkout.

`tools/land.sh` needs `gh` (signed in) and `jq`. Tunables: `LAND_POLL_SECONDS` (30),
`LAND_SETTLE_SECONDS` (90), `LAND_TIMEOUT_SECONDS` (5400). Tests: `bash tools/tests/test_land.sh`.
