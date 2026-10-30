# Review: wt/client-hardening (2026-10-30)

Reviewed `origin/main...HEAD` (10 commits, 18 files). Read client.yml, bis.yml, engine.yml, data.yml (unchanged), web.yml, api.yml and the pin-engine action where they bear on the move.

Ran:
- `bash tools/tests/test_activation.sh`: ok 48, exit 0.
- `bash tools/engine/tests/test_scripts.sh`: ok 51, exit 0.
- `uv run pytest tests/test_bis_carry.py tests/test_clientwatch.py`: 34 passed.
- `tools/engine/pinned-build.sh .`: 1.60.1.70658. This is the active build, so tonight's bis.yml has `follows=true`.

Checked and found sound:
- Gate exit codes: `wait || x=$?`, an empty package list fails closed, a crashed parser fails closed under `set -e`, and `(suite)` fallbacks are in place. I found no path where the gate exits 0 on a failure.
- `run()` keeps pipefail status through `tee`.
- The no-gate guard.
- The re-encoder's loud failure and the three `--check` gates. Both are still run, both now through `run`, so a failure still stops the move.
- The PTR regex against today's `products` (`wowf`, `wow_cn_beta`).
- There is no dispatch loop. The chains are client -> engine(force) [-> bis(force) only when RANK_NOW], engine pin -> web/api/sim/bis, and bis -> web.
- The Postgres service plus `TEST_DATABASE_URL`. The job runs on the host, so `localhost:5433` is right; this matches api.yml.
- `npm ci`: `web/.node-version` exists and nothing in `pretest`/`sync` needs Go.
- pin-engine before `go test` (the go.work resolution).
- The order inside report_failure's EXIT trap.
- The report job's `if:` covers watch and fetch failures and does not double-report activate.

## CRITICAL

None.

## HIGH

### H1. The gate is not what web's deploy waits on, so the addon can still be tagged for a client the site cannot build

`tools/activation-gate.sh:8-12` and `tools/activate-build.sh:219-226`.

The gate's header says it runs "exactly what web.yml's `build` job runs before it deploys". It runs only `npm test`.

What web.yml actually runs before the site goes live:
- The `build` job runs `npm run check`, `lint`, `format:check`, `npm test` and a fixture `npm run build`.
- The `deploy` job runs `make simdb`, `make artifacts` and the **real-data** `CF_PAGES=1 npm run build`. That build is the one that prerenders every guide, BiS and planner page from `data/builds/<new>/`.

Why vitest misses this:
- `pretest` sets `FOREVER_DATA=fixture`, so almost every vitest case reads the fixture, not the new build.
- A new client that breaks a prerender breaks only the real build. Examples: a talent a guide names but the trees lost, a schema-rejected frontmatter field after the re-encode, or a bis file the page code rejects.

The result is the exact half-ship this lane exists to stop:
1. The gate goes green.
2. The move is pushed and `addon-v<x>` is tagged and published.
3. `web.yml`'s deploy goes red, and the site keeps serving the old build against a new addon.

**Fix:** add these to the gate, in parallel with the API leg:
- `npm run check` (about 1 min on CI);
- `npm run format:check` (about 35 s; the re-encode and link sed edit markdown);
- a real-data `npm run build` with `FOREVER_DATA` unset and no `PUBLIC_SIM_ENGINE` (about 2.5 min on CI, the same as build-real; it needs no wasm).

Then correct the header comment. That is still well inside the 60 minutes.

### H2. Merging tonight runs bis.yml's dps-per-point re-measure on the live build for the first time, before the picks are saved, with no runtime data and no failure isolation

`.github/workflows/bis.yml:261-280` and `.github/workflows/bis.yml:286-292`.

Why it runs tonight:
- Every `dps-per-point.json` from 70291 through 70658 has the same bytes (md5 `778c0d12...`). The file has only ever been carried; it was measured once by hand on 2026-10-12.
- The engine is already stamped for 70658, so `follows=true`.
- The bis.yml change alters `HEAD:.github/workflows/bis.yml` in the input hash.

So the first nightly after merge does all of this:
- forces a full `-all` ranking;
- then runs `go run ./sim/cmd/dps-per-point` for the first time in CI. That is every spec times 5 bands at 500 iterations, 8x for healers, and its runtime is unmeasured (the report's own concern 5);
- changes the addon's "+N DPS" figures and the API's loot values in production four days before launch.

The step sits between the ranking and the picks' save step with no `continue-on-error`. If it fails or times out, the night's ranking (about 70 runner-minutes) is discarded and nothing lands. On launch night that would also strand the carried picks with their `not_in_client` empty slots.

**Fix:**
1. Give the step `continue-on-error: true`. Stage dps-per-point.json only when the step succeeded, so the picks land without it if it fails.
2. Before Nov 4, run it once by `workflow_dispatch` on the current build to measure its runtime and review the diff.
3. Or deliberately keep it off the live build until launch, and say so in the runbook.

## MEDIUM

### M1. A stalled engine follow is green every night

`.github/workflows/bis.yml:86-101` and `bis.yml:210-212`.

`follows=false` makes `skip=true` and emits only a `::notice`. The job is green, and nothing ranks or redeploys.

The ranking also waits when the stamp cannot be read: a transient blobless `git show` fetch failure becomes "an unreadable build" and the run is still green.

After a launch move whose follow went to review (the likely case, since talent trees change), the carried picks and the dropped `not_in_client` slots stay live with no red signal anywhere.

**Fix:**
- Emit `::warning` instead of `::notice`.
- Fail the job, or open an issue, once the active build has waited more than one night. For example, compare the activation's date with today.
- Treat "unreadable" as a failure, not a wait.

### M2. A half-finished move is not idempotent; rerunning client.yml silently completes nothing

`tools/activate-build.sh:124` and `tools/activate-build.sh:270-310`.

After the push, a tag-push or dispatch failure is reported once. Every later run, nightly or by hand, says "already the active build; nothing to do" and exits 0. The addon is then never tagged or published, and web, api and engine are never dispatched, unless a person runs the printed commands from that one night's summary.

**Fix:** when `NEW == ACTIVE`, check that `addon-v<TOC version>` exists on origin at a revision carrying the move. If it does not, resume from the tag step and redo the dispatches. Then "redo the move" really is the hand command for every stop.

### M3. The hit-line tests now pass vacuously on the fallback

`web/src/pages/guides/_spec.test.ts:46-53` and `web/tests/e2e/guides-sections.spec.ts:114`.

`guideHitTalent` reproduces the page's own logic (`schoolHitTalent(...) ?? 'talents'`). If the build's talent is not in `SCHOOL_HIT_TALENTS`, or `decodeBuildTrees` returns undefined, the page and the test both say "talents" and the test passes. This is the 70685/70707 case, where "Precision" is the talent.

The e2e accepts `[^,]+`. A regression where the page stops naming any talent is now invisible. This is the report's concern 1; it is quieter than the report says, because no test can fail on it.

**Fix:** keep the derivation, but assert that `guideHitTalent('mage/frost')` is not `'talents'`. Allow an explicit, commented allowlist of builds where the fallback is the accepted copy, so moving to such a build is a decision rather than silence.

### M4. The new shell tests run in no workflow

`tools/tests/test_activation.sh` and `tools/engine/tests/test_scripts.sh`.

These 99 assertions are the only coverage of the gate's parsing, the failure report, the script modes and the executable-bit check. The fileMode regression fixed in 6d652f162 shows how easily the scripts break unnoticed.

**Fix:** add a job to an existing workflow. The cheapest fit is data.yml's test job or a small `tools.yml` on `tools/**`, `.github/workflows/{client,bis,engine}.yml` and `.github/actions/**`. It should run both scripts, plus `shellcheck` and `actionlint`, which the report says are clean.

### M5. bis.yml now changes data the API serves but redeploys only web

`.github/workflows/bis.yml:323-326`.

`api/internal/sims/lootvalues.go` and `gearscore.go` read `dps-per-point.json` from the image's `/data` copy of `data/builds`, and api/internal/bis reads the picks. The bis.yml run dispatches `web.yml` only, so the API serves the carried figures until something else redeploys it. Picks already lagged this way before this branch; the re-measure makes the gap matter for loot values.

**Fix:** also `gh workflow run api.yml --ref main` after a run that changed `bis/` or `dps-per-point.json`.

## LOW

### L1. A real step that exits 2 gets no summary

`tools/activate-build.sh:93`.

The trap skips the report for exit 2, to keep usage errors quiet. A wrapped command that exits 2 also gets no summary or annotation: argparse in any `uv run python -m pipeline ...`, or grep errors.

**Fix:** set a `USAGE_EXIT` flag in `usage()` and the argument guards, and skip only on that.

### L2. Unchanged lines that can stop the move, now under the readable report

- `tools/activate-build.sh:168`: the `grep -rl ... | while` loop under `pipefail` exits 1 when no guide links `foreversixty.gg/data/<active>/`. Only one guide does today.
- `tools/activate-build.sh:230`: `git status --short | head -40` could SIGPIPE on a very large status.

Neither has fired, but either would be reported against the wrong step. **Fix:** wrap the grep as `{ grep ... || true; }`, and replace the head with `awk 'NR<=40'`.

### L3. The staging line fails if the build has no dps-per-point.json

`.github/workflows/bis.yml:292`.

The staging line names `dps-per-point.json` explicitly. If the build has none, staging fails with "pathspec did not match", for example when an old build had none to carry. 70685 and 70707 have none on main today; they get one only through the activation's carry.

**Fix:** stage it only when `[ -f ... ]`.

### L4. The gate is a strict subset of api.yml's test job

`tools/activation-gate.sh:57-63`.

api.yml runs `go vet ./...` and `-race` across the whole module. The API image job and deploy are also not gated. An API red after the tag does not strand the addon the way the web does, so this is a note. Runtime is fine:
- api.yml's full `-race -p 1 ./...` takes about 16 min on CI (run 40979952412);
- vitest takes about 2 min (run 41079910252);
- the gated subset without `-race` will be lower;
- the 60-minute budget has room even with H1's additions.

### L5. Runbook timing is slightly off

`docs/runbooks/launch-client-move.md:15`.

The table puts the engine follow at 05:07. On a move night the activation dispatches `engine.yml -f force=true` right after the push, so the follow starts at about 04:00. The 05:07 schedule then queues behind it in the `engine` concurrency group. Say so, so the person watching does not wait for 05:07.

### L6. A timeout can produce a wrong summary, or none

`.github/workflows/client.yml:172`.

If the activate job hits `timeout-minutes: 60` mid-move, `steps.move.outcome` is not `failure`. The setup-failure step would then claim "nothing was moved", even when the timeout came after the push. Or `failure()` is false on a cancel, and no summary is written at all.

**Fix:** have the script write a breadcrumb file per step, and have the report step read it, with `if: always() && steps.move.outcome != 'success'`.

## Verdict

**FIX FIRST.** H1 leaves the main half-ship this lane exists to prevent: the addon tagged while the site's real build is red. H2 makes tonight's nightly the first-ever CI run of an unmeasured step that can discard the picks. Both fixes are small, and both should land before the eve check on Nov 3.
