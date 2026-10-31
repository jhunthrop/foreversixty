# Client build move hardening (2026-10-30)

Lane: `wt/client-hardening`, rebased on origin/main. Built on the two rehearsal commits (`6a22fa669` watcher warnings, `11f891be0` `pipeline carry-bis`) and the rehearsal report, `design/reviews/2026-10-30-launch-rehearsal.md`.

Nothing was pushed, tagged or dispatched. Production, the engine fork and the guide FS1 builds were not touched.

## Commits

| Sha | What |
|---|---|
| `80f909271` | **Test before pushing.** `tools/activation-gate.sh` runs, in parallel, the whole web vitest suite and `go test -p 1` over every API package whose Go files name `active-build`, `ActiveBuild` or `data/builds` (20 today, found by that grep). `activate-build.sh` runs it after the `--check` gates and before `git add`. When it is red, nothing is committed, pushed or tagged, and the failing tests go to the summary. `client.yml`'s activate job gains Postgres, `TEST_DATABASE_URL`, pin-engine and `npm ci`. The push uses `--autostash`. `--no-gate` is allowed only with `--no-commit`. |
| `be108631c` | **The hit-cap test follows the build.** `_spec.test.ts` decodes the frost guide's build against `active-build.json` and asks `schoolHitTalent`, the same way the page does. The e2e walk checks only the line's shape. |
| `180fb4eac` | **The ranking waits for the engine pin.** `tools/engine/pinned-build.sh` reads the fork's `FOREVER_BUILD` at the pinned sha. `activate-build.sh` dispatches `bis` only when that stamp matches the new build; otherwise it writes "ranking waits for engine follow". `bis.yml` refuses a mismatched engine even when forced. It re-measures `dps-per-point.json` once the file is still the carry, commits it, and drops it from the input hash. `engine.yml`'s green pin dispatches `bis` with force, and the auto-merge of a hand-fixed pin pull request now dispatches the consumers too. |
| `4356d1db3` | **No PTR activation.** `client.json` gains `test_product_pattern` (`ptr\|test\|^wow[a-z]?t$\|^wowz$`). It is checked before every other rule: such a key is never Forever and is not warned about. `load_config` refuses a listed product that matches it. |
| `77849c0ea` | **Readable failure summaries.** Each stage of `activate-build.sh` has a step name and a hand command. An EXIT trap writes the step, the reason (the step's output, or the failing tests) and the exact command. It goes to stderr, the run summary and an `::error` annotation, plus an undo line on local runs. Dispatch failures are collected and named. `client.yml` passes `PRODUCT`, and report steps cover setup, watch and fetch failures. |
| `6d652f162` | **The new scripts are executable.** `core.fileMode` is off, so the gate and `pinned-build.sh` had gone in as 100644. The fake move below caught it: the gate failed with "Permission denied". A test now asserts the index mode of every script the workflows call. |
| `c1f0300ce` | **Runbook.** `docs/runbooks/launch-client-move.md` covers the eve and launch night, finishing each kind of stop, hand-dispatching for a new key or a 2.0 renumber, and rollback. |

## Verification

- **data pytest (full):** 1608 passed, 9 skipped. ruff is clean.
- **Shell tests:**
  - `tools/tests/test_activation.sh`: 48 assertions. They cover the gate's parsing with stubbed npm and go, the real `activate-build.sh` in scratch repositories (re-encode stop, under Actions, unfetched build, long log, argument cases) and the script modes.
  - `tools/engine/tests/test_scripts.sh`: 51 assertions.
  - `actionlint` is clean on client, bis and engine; `shellcheck` is clean on the new scripts.
- **Web:** `_spec.test.ts` passes, 11/11. With "Elemental Precision" renamed in the build's `mage.json`, the new test passes and the old one fails. eslint and prettier are clean.
- **API:** no API package was changed. The gate ran the 20 active-build packages green against the private database `forever_gc_harden`, which has since been dropped. The full gate took 8m47s on the laptop.
- **Dry run, 1.60.1.70707 `--no-commit`,** on a scratch clone: the run stops at the re-encode. The summary names the step, lists all six unspent mage builds (Arcane Focus, Improved Fireball, Elemental Precision, Critical Mass) and gives the re-spend-then-rerun command. The printed undo line restores the clone to clean.
- **Full fake move, `1.60.1.70999`** (a restamped copy of 70658), on a scratch clone with a local bare origin and a failing `gh` stub:
  - every step passes, and the gate is green;
  - the move commits, pushes to the local origin and tags `addon-v2.0.21`;
  - it writes "ranking waits for engine follow" and does not dispatch `bis`;
  - then it fails, naming each dispatch that could not run.

  The clone and the origin were removed afterwards.

## Fix round (review: docs/superpowers/reports/2026-10-30-client-hardening-review.md, FIX FIRST)

| Sha | Finding | What |
|---|---|---|
| `7840e0cdc` | H1 | **The gate runs what web's deploy waits on.** The web leg runs `npm run check`, `lint`, `format:check`, the real-data `CF_PAGES=1 npm run build` (with `FOREVER_DATA` and `PUBLIC_SIM_ENGINE` unset) and `npm test`, in parallel with the API leg. They run in sequence because check's real sync must come before the build and the test's fixture sync after it. Every red script is listed. The new end-to-end shell test proves that a prerender failure stops the move with nothing committed, pushed or tagged, and that the same move with a green gate pushes and tags. The real gate on 70658 is green in 3m53s locally. |
| `69180e076` | H2, L3 | **The dps-per-point re-measure runs only on a move.** The activation's carry writes `dps-per-point.carried-from`, and bis.yml re-measures only when that marker exists. No build on main has one, so tonight's nightly does not re-measure. The step has `continue-on-error` and a 90-minute timeout. The figures and the marker's removal are staged only on success; on failure the partial file is restored before addon-data renders. `dps-per-point.*` leaves the input hash. |
| `5ff7dd48d` | M1 | **A stalled ranking turns the run red.** A new `tools/engine/ranking-wait.sh` returns rank, wait or stalled. Stalled means a second night (the move is 20 hours old or more) or an unreadable stamp. The run warns on both, and on stalled it fails at its last step, after the crawl commit. The repo opens no issues from any workflow, so the run fails instead. |
| `bb2b57fec` | M2 | **A rerun resumes a move stopped after its push.** The post-push steps (tag, engine check, dispatches) are now `finish_move` and safe to rerun. `resume_move` finishes a move only when its own commit set the addon version and `addon-v<version>` is not on origin or was never released. `client.yml` resumes on a hand dispatch of the active build, and on a nightly where the TOC version has no tag. The end-to-end test runs the move three times: the first stops at a refused tag, the rerun tags and dispatches, and the third has nothing to do. |
| `cd716e472` | M3 | **The hit line must name a talent.** Unless the build is in `HIT_LINE_FALLBACK_BUILDS` (empty), the test asserts that the frost guide names a school-hit talent. **The move to 70707 will stop at the gate on this test** until the mage lane adds Precision to `SCHOOL_HIT_TALENTS` or the build to the allowlist. |
| `5686388bf` | M4 | **The shell tests run in CI.** `tools.yml` runs both shell suites, shellcheck over tools/ and actionlint over client, bis, engine and tools. |
| `fb27289af` | M5 | **A ranking redeploys the API too.** bis.yml also dispatches `api.yml`. |
| `600e034ca` | L1, L2, L6 | **Every stop gets a report.** The trap reports every non-zero status, since usage errors exit before it is set. The link-rewrite grep can no longer fail the pipeline, and the status listing can no longer SIGPIPE. A per-step breadcrumb, a 50-minute timeout on the activate step and an `always()` report step cover timeouts and cancels. |
| `b8668c8f1` | L5 | **Runbook.** On a move night the engine follows at about 04:00. The runbook now covers the new gate, the marker, the red stall, the resume and the 70707 hit-line decision. |

**Skipped:** L4. The API leg stays the active-build subset without `-race` or `go vet`. An API red after the tag does not strand the addon the way a web red does, and the review rates it a note. The header says so.

**The input hash:** merging this branch forces one full re-ranking (about 70 runner-minutes) on the first nightly, with no re-measure. bis.yml is a hashed input, and any change to the hash formula also changes its value, so dropping the workflow file from the hash would not avoid it. The ranking is deterministic, so the picks should come out the same. If that night matters, dispatch bis.yml at a chosen time instead.

Fix-round verification:

- `test_activation.sh`: 91 assertions; `test_scripts.sh`: 59.
- shellcheck and actionlint are clean.
- data pytest: 1608 passed, 9 skipped.
- `_spec.test.ts`: 11/11 on 70658, and red with the allowlist message when the talent is renamed.
- The real gate is green.
- The 70707 `--no-commit` dry run still stops at the re-encode, with the readable summary and annotation.
- The private database has been dropped.

## Concerns

1. **The hit line on 70707 needs a decision before the move.** The talent is now called "Precision", which is not in `SCHOOL_HIT_TALENTS`, so the gate stops the move on the frost test (M3) until someone adds it there, or adds the build to `HIT_LINE_FALLBACK_BUILDS`. This is rehearsal recommendation 4, for the mage lane.
2. **`builds.test.ts`'s `SIGNATURE_TALENTS` still names talents.** I left it on purpose: it states curated rotation intent, and deriving it from the APLs diverges from the table for 12 specs. A rename of one of those talents fails that test by name, after a hand re-spend.
3. **Gate runtime.** It took 3m53s on the laptop with the real build; an earlier vitest-only run took 8m47s while the laptop was loaded. CI has not been measured. The activate step's timeout is 50 minutes.
4. **`tools.yml` is new.** Its first run on main is the first time the shell suites run on Linux (GNU tools rather than BSD).
5. **The dps-per-point re-measure has never run in CI.** Its runtime is unmeasured. It now runs only on a move night and is failure-isolated, but a measured run (`go run ./sim/cmd/dps-per-point -repo-root .`) by dispatch before 4 November would de-risk launch night.
6. **The re-encode hand path still needs the mage respec.** The other lane's 70707 respec has to land before tonight's nightly can move on its own. Until then, every nightly stops at the re-encode, now with the readable summary.
7. **A stray `/private/tmp/claude-501/go.work`** (rehearsal recommendation 7) is still there. Scratch clones with their own `go.work` are unaffected.
