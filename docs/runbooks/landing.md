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
   - checks that every workflow the batch's paths reach (`tools/land_workflows.py`) has a
     run. A push whose head sha the remote already held, say from a lane branch pushed at the
     same commit, carries no commits, so GitHub starts no path-filtered workflow for it (land
     2026-11-07, bnet-refresh: only companion ran). land.sh dispatches each missing workflow
     on the land branch (`gh workflow run`) and waits for it; one that still never starts is
     a timeout, never a landing;
   - leaves out a workflow turned off on GitHub (`disabled_manually`, as the refresh hold kept data,
     bis and engine until 2026-11-12), saying "skipping <workflow>: disabled (disabled_manually); not gating this
     landing" for each; any other disabled state fails the landing, so a workflow switched off by
     mistake never lets a batch through quietly (2026-11-08: data.yml's `addon/**` path held every
     addon batch);
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

## Launch mode

Owner, 2026-11-03 (WoW Forever launches 4 Nov): smoke gates only until after launch. One
switch, the repository variable `LAUNCH_MODE`; every check workflow reads `vars.LAUNCH_MODE`.
While it is `true`:

- What still runs (target 5 to 8 minutes of wall time on a land branch):
  - web `build`: check, lint, format, the guide prose test
    (`src/content/guides/_prose.test.ts`) and the fixture build. deploy runs on main as before.
  - api `check`: the engine-free `go vet` and `go build ./...` with `GOWORK=off`, and the
    workspace `go vet`. `image` and `deploy` run on main only.
  - sim `test`: gofmt, `go vet`, `go build ./...`. `rotations`: apl-check and styles-check.
  - data: ruff and a byte-compile. tools: shellcheck and actionlint. addon: luacheck and the
    Lua 5.1 and Data.lua parse checks. logs: gofmt, vet, build. companion: gofmt, vet and the
    headless build (its release on main is unchanged).
- What is skipped (jobs show as skipped, which `tools/land.sh` counts as green):
  - the full vitest beyond the prose test;
  - the e2e shards, real-engine, lhci and lhci-alert;
  - every Go test run and `-race`: the api's four test shards, the sim suite and its
    artifacts job (the wasm smoke), logs' tests, fuzz and bench, the image tests;
  - the iOS and Android suites and mobile-codegen's drift checks;
  - the companion tests and the window build;
  - addon busted, data's pytest and `--check` gates, the tools script tests and coach;
  - the fork suite (engine.yml is not a push check; leave its schedule alone).
- `smoke.yml` runs on main after every green web or api run (each ends in its deploy):
  `tools/prod-smoke.sh` fetches the home page, /guides, one guide per class (each must carry
  its Priority list), /bis/warrior/fury, /tiers, /sim, the API's /health and /go/discord, and
  fails with a summary naming every URL that answered wrong. Run it by hand with
  `gh workflow run smoke` or `bash tools/prod-smoke.sh`.
- Lanes run no local tests: commit and push straight to `land/<n>`, then
  `tools/land.sh land/<n>`. The smoke gate is the land branch's run and smoke.yml after the
  deploy.

Turning it off (after launch):

1. `gh variable set LAUNCH_MODE --body false`.
2. Run the full suite once on main HEAD by dispatching every check workflow by file name:

   ```
   for w in web api sim tools addon logs coach companion ios android mobile-codegen; do
     gh workflow run "$w.yml" --ref main
   done
   ```

   Not data.yml: its dispatch always runs the fetch job, which downloads and commits a
   build and then starts a ranking (see "The refresh chain"). Run its `test` job's steps
   locally instead (in data/), or let the next push that touches data/ run it.
   An empty land branch is not enough: a land branch is diffed against main, so an empty
   one starts only companion.
3. Fix what it finds, in one landing, before the next feature batch.

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
`/<tree>/` line after the `!` lines re-includes it). data.yml's fetch job, and the deploy and
release jobs, which no landing runs, keep the full checkout.

`tools/land.sh` needs `gh` (signed in), `jq` and Python 3 with PyYAML. Tunables: `LAND_POLL_SECONDS` (30),
`LAND_SETTLE_SECONDS` (90), `LAND_TIMEOUT_SECONDS` (5400). Tests: `bash tools/tests/test_land.sh`.

## The refresh chain

Since 2026-11-12 the data, engine and ranking workflows run on events; their schedules
(client 03:41, engine 05:07, bis 06:23 UTC) stay as the safety net. Every link skips when
nothing it reads has moved (engine's `Follow-Inputs`, bis's `Inputs:` trailer), and bis's
`bis` concurrency group never runs two rankings at once. No link loops: every bot commit is
pushed with `GITHUB_TOKEN`, which starts no workflow, so bis's commit to `data/builds/*/bis`,
`addon-data.json` and `Data.lua` never re-enters data.yml's push run, and data.yml's own
commits never re-enter data.yml.

| Event | Starts | How |
|---|---|---|
| wago.tools lists a newer Forever build (client's 03:41 watch, or `gh workflow run client.yml -f build=<build>`) | data.yml's fetch for it, then `tools/activate-build.sh` | `workflow_call` inside client.yml |
| The activation lands on main | web, api, sim, addon-release (the `addon-v<version>` tag it made) and engine with force; bis with force only when the pinned engine already follows the build | `gh workflow run` from the script |
| engine.yml pins a fork that follows the active build | web, api, sim and bis with force | `gh workflow run` from the pin job (also after it auto-merges a green `engine: pin <sha>` pull request) |
| A person pushes `sim/enginever/version.go` to main (a sim lane's hand pin) | bis | `on.push` paths filter in bis.yml |
| A dispatched data.yml run finishes green on main (a hand refresh: `gh workflow run data.yml -f product=<key> -f build=<build>`) | bis | `on.workflow_run` in bis.yml; a push-triggered data run starts nothing |
| client's push-check finds a new hotfix push for the active build | data.yml's fetch again (refetch), then engine; engine's pin then ranks | `workflow_call`, then `gh workflow run engine.yml` from client's refollow job |
| bis commits new picks | the web and api deploys | `gh workflow run` from bis.yml |

By hand: `gh workflow run data.yml -f product=<key> -f build=<build>` fetches and then ranks;
`gh workflow run engine.yml -f force=true` follows the active build and ranks on a green pin;
`gh workflow run bis.yml -f force=true` ranks whatever main holds now. A workflow_run or push
start of bis checks out main's head at start time, not the event's sha.

bis's shape (since 2026-11-12, when a full ranking outgrew one job): a `prepare` job
checks out main's head, commits the wowhead crawl, reads the engine verdict and the inputs hash,
and prints the matrix (`tools/bis-matrix.sh`: one `rank <class>` leg per class, every spec with a
written rotation); each leg runs `make bis` with `-all -specs <its specs>` from the sha prepare
pushed and uploads its `<spec>.json` and `.md` as the artifact `bis-picks-<class>`; one `commit`
job downloads every leg's picks over its checkout, runs the sanity tests, the dps-per-point
re-measure and the addon data, and pushes the one `chore(bis)` commit, then dispatches web and
api. The wall clock is the priest leg (Holy and Discipline, about 80 minutes on a fast host, 120
on a slow one) plus a few minutes at each end. A leg that does not finish green fails the
commit job by name and nothing lands; the green legs' picks stay 14 days
(`gh run download <run> -p 'bis-picks-*'`). A spec that fails inside a green leg keeps its
previous file, as before. A ranking stalled on the engine for a second night is the red
`stalled` job.

To prove a bis.yml change without a nightly: push it to a lane branch and dispatch two quick
specs there, `gh workflow run bis.yml --ref <branch> -f force=true -f specs=rogue-assassination,hunter-marksmanship`.
Off main the run pushes and deploys nothing: the crawl and the picks commit are built and
printed. On main, `-f specs=...` ranks only those specs and the commit's `Inputs:` trailer is
marked partial, so the next full run still ranks everything.

The addon tag stays a hand step. bis commits the new picks into `addon/ForeverSixty/Data.lua`
on main, but players get them only with the next `addon-v<version>` release: a
`chore(addon): <version>` commit (TOC version and changelog heading) and the tag, as
`tools/activate-build.sh` does for a client move. Nothing tags on a re-rank.
