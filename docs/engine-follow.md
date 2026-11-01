# The engine follows the client

The wowsims fork (`jhunthrop/wowsims-forever`, branch `master`) is the sim engine behind every simulated number on the site. Its talent trees, spell constants, stat tables, relic mods and set rows are generated from the same client tables the site mines into `data/builds/<build>/`. `engine.yml` regenerates them when the client moves and pins the site to the result. The design is `docs/superpowers/specs/2026-10-23-engine-follow-design.md`; this page is how to operate it.

## What runs when

`.github/workflows/engine.yml` runs at 05:07 UTC every day, after the client watcher and the data lane's commits and before the ranking at 06:23. It also runs the morning `tools/activate-build.sh` moves the site to a new build, dispatched with `force`.

It has two jobs.

- **follow** hashes the build id, the site's committed inputs for it and the fork's generator sources, and compares the hash with the `Follow-Inputs:` trailer on the fork's last follow commit. Equal and not forced: it skips. Otherwise it fetches the build's raw tables, runs the fork's `tools/follow.sh`, runs the fork's own suite and reads the report's verdict.
- **pin** runs whenever the fork's `master` is not the site's pin, the fork's `assets/db_inputs/FOREVER_BUILD` names the active build, and the fork's `test` workflow passed for that sha (it waits up to 30 minutes for a run in progress). If any condition fails it does nothing and the step summary says which. Otherwise it pins, runs the site's sim suite and the data lane's pin test, and either pushes the pin to `main` and dispatches `web.yml`, `api.yml`, `sim.yml` and `bis.yml --force`, or opens a site pull request.

A half-followed fork is never pinned, and neither is a commit the fork's own CI has not passed.

## The two verdicts

The follow report starts with `verdict: clean` or `verdict: review`.

**Clean and the fork's suite green.** The bot commits `chore(client): the engine follows client <build>` to the fork's `master`. The fork's CI runs on that commit, and the pin follows. Nothing for you to do.

**Review, or the suite red.** The same commit goes to the branch `follow/<build>` in the fork, with one pull request titled `engine: follow client <build>`. A later run updates it; it never opens a second. The pin does not run.

A review verdict means the generators moved something only a person can judge: a talent whose text, rank count or spell id changed, an ability with no disposition, a racial text change, or a conformance row that got worse. The first real run is expected to end here, for 1.60.1.70532: the Frost tree rework removed Piercing Ice, `sim/mage/talents.go` still names it, and the engine does not compile until a person fixes that.

## The fork pull request: `engine: follow client <build>`

The body is the follow report: what changed per class, the talents to check, the trainables to register or dispose of, the racials to re-read, and the fork suite's result with the failing output if it is red. The branch holds the regenerated files and nothing else.

Do the hand work on that branch: fix what no longer compiles, correct the talent rates the report lists, add dispositions, fix the racials. Run the fork's suite. Merge to `master`. The pin lands the next night, or at once if you dispatch `engine.yml`.

A merged `follow/<build>` left on the fork is stale and the next follow overwrites it, but enabling "Automatically delete head branches" on the fork (Settings, General, Pull Requests) avoids the stale-branch case entirely.

## The site pull request: `engine: pin <sha>`

The fork passed its own CI and follows the active build, but the site is red against it: `go vet` or the sim suite in `sim/`, `make apl-check`, `make styles-check`, or `data/tests/test_genproto.py`. `main` keeps the old pin. The pull request comes from `engine-pin/<sha>`, which already carries the pin commit, and its body is the failing output. It is opened with the job token, so no workflow runs on it; the person merging runs the checks by pushing to the branch or dispatching `sim.yml` on it.

Almost always a correct engine change moved a number a site golden pins. Check out the branch, regenerate what the failing test names, read the diff to confirm the new answer is right, and merge.

## Pull requests need a repository setting

The site repository has "Allow GitHub Actions to create and approve pull requests" on (Settings, Actions, General, Workflow permissions; turned on 2026-10-24). If it is ever turned off again, a red pin only pushes `engine-pin/<sha>` and warns that it could not open the pull request; open it by hand from the branch. The fork's pull requests use `ENGINE_REPO_TOKEN` and are not affected.

## `ENGINE_REPO_TOKEN`

The site repository secret `ENGINE_REPO_TOKEN` is a fine-grained personal access token on `jhunthrop/wowsims-forever` with Contents and Pull requests read/write. It pushes to the fork's `master`, pushes `follow/<build>` and opens its pull request. Everything on the site side uses the job's own `GITHUB_TOKEN`; the raw fetch needs no token.

Expires: no expiration (created 2026-10-24). If it is ever recreated with an expiry, write the month here and renew before it, or live runs stop.

Until the secret exists, every run that is not already a dry run turns into one and warns. Nothing is pushed or opened.

## Forcing a follow and previewing one

Follow now, even if nothing the follow reads has changed:

```bash
gh workflow run engine.yml --ref main -f force=true
```

Follow a specific build instead of the active one:

```bash
gh workflow run engine.yml --ref main -f force=true -f build=<build>
```

Preview: run everything, push nothing, open nothing.

```bash
gh workflow run engine.yml --ref main -f dry_run=true
```

The report, the follow paths, the suite log and the generated diff are uploaded as the artifact `follow-<build>` on the run.

## By hand

The workflow is a schedule around the same commands. In the fork, with the site's raw tables for the build fetched first. The product is the `"product"` in `data/builds/<build>/manifest.json` (`wowf` for Forever):

```bash
cd /path/to/forever/data
uv run python -m pipeline fetch --product <the "product" in data/builds/<build>/manifest.json> --build <build>

cd /path/to/wowsims-forever
make follow SITE_DIR=/path/to/forever BUILD=<build>
```

`fetch` also layers the server's hotfix push for the build into the raw tables (`data/pipeline/hotfix_push.py`; the file as fetched stays in `raw/shipped/`), so the fork's raw-table generators (relicmods, periodiccrit, eureka, the set rows) read the server's rows. `uv run python -m pipeline layer-hotfix-push --build <build> --offline` re-layers from the cached push without the network.

`make follow` runs `tools/follow.sh`, which refuses a dirty checkout, stops at the first failing generator, writes `assets/db_inputs/FOREVER_BUILD` and the report at `dist/FOLLOW-REPORT.md`, and lists the paths it wrote in `dist/FOLLOW-PATHS`. The script exits 0 for clean and 3 for review. `make` cannot pass 3 through, so `make follow` exits 2 on review; run `tools/follow.sh` directly when the code matters.

Run the fork's suite (`go test --tags=with_db ./sim/... ./tools/...`) and commit the paths in `dist/FOLLOW-PATHS`. To pin that fork commit on the site, from the site root:

```bash
make engine-pin ENGINE_DIR=/path/to/wowsims-forever
cd data && uv run python -m pipeline simproto --engine /path/to/wowsims-forever
```

Both pins move together. `data/tests/test_genproto.py` fails when they disagree.

## Why the drift guard warns

`sim.yml`'s test job compares the pinned engine's `FOREVER_BUILD` with `web/src/data/active-build.json`. When they differ it prints a warning annotation and a summary line; it does not fail.

The hours between a build's activation and the engine pin are a normal state. Failing every sim push on `main` during them would teach people to ignore red. The warning is how a drift that lasts longer than a night gets noticed.
