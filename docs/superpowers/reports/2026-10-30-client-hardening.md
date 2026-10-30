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

## Concerns

1. **The hit-line fallback is quiet.** In 70685/70707 the talent is "Precision", which is not in `SCHOOL_HIT_TALENTS`. The guide therefore says "from talents", and the test now accepts that. The copy decision (Precision now covers all spells) is still open: rehearsal recommendation 4, for the mage lane.
2. **`builds.test.ts`'s `SIGNATURE_TALENTS` still names talents.** I left it on purpose: it states curated rotation intent, and deriving it from the APLs diverges from the table for 12 specs. A rename of one of those talents fails that test by name, after a hand re-spend.
3. **Gate runtime.** It took 8m47s on the laptop. CI has not been measured. The activate job's timeout is 60 minutes.
4. **The new shell tests are not run in CI.** Neither is the existing `tools/engine/tests/test_scripts.sh`. They should be added to a workflow.
5. **The dps-per-point re-measure is new work in `bis.yml`,** with an unmeasured runtime, and it runs only on move nights. It detects the carry by byte-equality with the nearest older build's file.
6. **The re-encode hand path still needs the mage respec.** The other lane's 70707 respec has to land before tonight's nightly can move on its own. Until then, every nightly stops at the re-encode, now with the readable summary.
7. **A stray `/private/tmp/claude-501/go.work`** (rehearsal recommendation 7) is still there. Scratch clones with their own `go.work` are unaffected.
