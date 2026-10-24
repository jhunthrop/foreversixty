# Engine follow: the engine tracks client builds on its own

Owner delegated the design on 2026-10-23: "follow our standards and make the best
user experience and accuracy as possible."

## Intent

Every night the data lane mines the client, and the morning a new Forever build
appears the client watcher moves the site, the addon and the ranking onto it.
The engine does not move. Its talent trees, spell constants, stat tables, relic
mods and set rows are all generated from the same client tables the site already
has in `data/builds/<build>/`, but a person runs the generators, reads the
conformance report, pushes the fork, and pins the site. Today the fork's tables
are stamped 1.60.1.70425 while the site serves 1.60.1.70532, two builds behind,
and nothing says so. The leveling BiS, the stat weights and every simulated
number on the site came from an engine that never saw the current client.

This design makes the engine follow the client the way the site already does:
automatically for everything a generator can prove, with one clear, complete
hand-off for the part only a person can judge, and a pin on the site the same
night the fork is ready.

## Goals

- When client data that affects damage output changes, the engine regenerates
  from it without a person, and the site pins the result the same night.
- A change the generators cannot absorb (a talent whose mechanic is new, a rate
  the engine carries in hand-written Go, a new trainable the engine does not
  register, a racial text change) reaches the owner as one pull request in the
  fork with a report that says exactly what changed and what the engine still
  gets wrong, so the hand work starts from a finished list, not a hunt.
- The site never pins an engine that fails the site's own tests, and never
  runs a ranking against an engine whose tables are stamped for another build
  without saying so.
- A person can run the same follow locally with one command and get the same
  result, so the automation is a schedule around a tool, not a second
  procedure.

## Non-goals

- Implementing new talent mechanics or rate changes automatically. The
  generators carry numbers; mechanics stay hand-written Go behind the
  conformance report, as now.
- Absorbing server-side hotfix tables. The owner's 22 Oct call stands: hotfix
  tables stay as seeded from 70450 until an official client carries them.
- Changing what the data lane mines or how `data.yml` runs.
- Item data. The engine already reads the site's `simdb.bin` at build time, so
  item changes follow today without this design.

## Anti-goals

- A pin that lands on the site and turns `main` red, or a ranking that runs
  against half-regenerated tables. Automation that breaks the morning's picks
  is worse than the drift it replaces.
- A silent follow. If the engine's numbers moved, the owner can find in one
  place what moved and why.
- A second copy of the generator sequence living in a workflow YAML. The
  sequence is one script in the fork; the workflow calls it.

## Where the pieces are today

| Generated artifact in the fork | Generator | Reads from the site build dir |
|---|---|---|
| `proto/<class>.proto`, `sim/<class>/talents_auto_gen.go`, `ui/core/talents/trees/<class>.json` | `tools/talentgen` then `make proto` | `talents/<class>.json` |
| `sim/<class>/constants_auto_gen.go` | `sim/core/spellconst/gen` (`make spellconst`) | `spellconst/<class>.json` |
| `sim/core/base_stats_auto_gen.go`, `sim/core/base_stats_levels_auto_gen.go` | `tools/base_stats_parser.py` after `tools/levels_from_site.py` | `gametables/*`, `levels.json` |
| `sim/<class>/relic_mods_auto_gen.go` | `tools/relicmods/gen.py` (`make relicmods`) | `raw/ItemXItemEffect.csv`, `raw/ItemEffect.csv`, `raw/SpellEffect.csv`, `raw/Spell.csv`, `raw/SpellName.csv`, `items/<class>.json` |
| `sim/core/client_sets_gen.go` | site repo `tools/gen_set_rows.py <build dir> <fork dir>` | `sets.json`, raw set tables, loot index |
| `sim/core/testdata/conformance/client/<class>.json` and `client/trainables/<class>.json` | verbatim copies | `spellconst/<class>.json`, `trainables/<class>.json` |
| `tools/talentgen/testdata/<class>.json` | verbatim copies | `talents/<class>.json` |
| `assets/db_inputs/gametables/<build>/*`, `assets/db_inputs/levels/<build>.json` | verbatim copy and `levels_from_site.py` | `gametables/*`, `levels.json` |
| `sim/core/testdata/conformance/<class>.golden.md`, the generated block of `SUMMARY.md` | `FOREVER_UPDATE_GOLDEN=1 go test --tags=with_db ./sim/conformance/...` | the copied client tables |

Hand-written, checked by the report but never generated: every ability file's
mechanics, every talent's rate in `sim/<class>/talents.go`, `sim/core/racials.go`,
`sim/conformance/dispositions.go` (why each unregistered trainable is not
modelled), `TALENTS.md` (talent rates against Forever's text).

Pins on the site: `sim/enginever/version.go`, `web/src/lib/sim/version.ts`, the
two sim fixtures (`make engine-pin`), and `data/proto/ENGINE_SHA` with the
vendored proto bindings (`python -m pipeline simproto --engine <fork>`).
`data/tests/test_genproto.py` fails when the two pins disagree.

The fork's integration branch is `master`: the pinned sha is on it, the owner's
work lands on it, and the `forever` branch is a stale mirror two commits behind
it. Everything below reads and writes `master`.

The raw client tables (`data/builds/<build>/raw/*.csv`) are gitignored on the
site; `python -m pipeline fetch --product <key> --build <build>` downloads them
from wago.tools (no token; the `WAGO_API_TOKEN` secret is for publishing the
addon), and the build's `manifest.json` records the product key it was fetched
with. Relic mods and set
rows read those tables, so the job fetches before it follows.

Defects the design corrects on the way:

- Three fork tests carry the build as a literal (`sim/core/talents/talents_test.go`
  wants `"1.60.1.70425"` and warrior tree sizes `{17,17,18}`;
  `tools/talentgen/main_test.go` wants `TalentsBuild = "1.60.1.70425"` and
  compares fixtures against an absolute `/Users/jh/...` path).
- The fork makefile's `BUILD` and `TALENT_BUILD` defaults name two different
  stale builds.
- The fork's CI (`test.yml`) runs on pushes to `forever`, not `master`, so the
  branch that matters has no push checks.
- The fork's CI is red on every recent run: `TestGeneratedProtosMatchSources`
  reports every `.pb.go` stale because the runner's apt `protobuf-compiler` is
  not the `protoc v36.2` the committed files were generated with. The committed
  files are right; the runner's protoc is wrong. CI installs the protoc release
  the committed headers name instead of the distribution package.

## The design

### One command in the fork: `make follow`

A script `tools/follow.sh` in the fork, reached as `make follow SITE_DIR=<site>
BUILD=<build>`, runs the whole generator sequence in the table above, in
dependency order, against one site build directory, and ends with the
conformance goldens and summary regenerated. It refuses to run on a dirty
checkout, and it stops at the first generator that fails. It writes the build it
followed to one file, `assets/db_inputs/FOREVER_BUILD`, which the makefile's
`BUILD` and `TALENT_BUILD` read as their default and which the tests read in
place of literals. Every generated header already names the build; the script
asserts after each generator that the header's build is the one requested.

The script is the only definition of the sequence. The workflow calls it; a
person calls it; the plan's local verification calls it.

A regenerated talent tree can remove or rename a proto field that hand-written
Go still names, so the engine may not compile after the generators run. The
script does not stop there: it records the build errors, leaves the goldens as
they were (they cannot be regenerated without a compiling engine), and still
writes the report, whose verdict is then `review` with a leading "Build"
section carrying the compiler output. The hand work that fixes the build is
exactly what the review pull request is for.

The script owns an explicit list of the paths it may write (the generated
artifacts in the table above, `FOREVER_BUILD`, the goldens and the summary's
generated block) and prints that list on exit. The commit stages exactly those
paths, never `git add -A`, so an incidental file from `make proto` or the
untracked report can never ride along. Writing a path outside the list is a
script defect and fails the run.

On the first run there is no `FOREVER_BUILD`; the script seeds it from the
`build` field of `tools/talentgen/testdata/warrior.json` (1.60.1.70425 today)
before it starts, so the talent diff below has a previous build to compare
against. The previous build's talent tables are the fixtures themselves, which
are exact copies of that build's `talents/<class>.json`; the site checkout is
not consulted for them, since an old build directory may no longer exist there.

### Fork tests read the build from the fixtures

The three literal-carrying tests change so a follow needs no code edit:

- `TestTreeSizesMatchTheClient` asserts the reader's sizes equal the generated
  `warrior.TalentTreeSizes` and that the reader's `Build` equals the fixture's
  own `build` field.
- `TestGoFileCarriesSizesBuildAndIDs` takes the expected build from
  `testdata/<class>.json`.
- `TestFixturesMatchTheSiteData` reads the site checkout from `FOREVER_SITE_DIR`
  (skipping with a message when unset) and the build from the fixture.

### The follow report

After the goldens regenerate, the script writes `FOLLOW-REPORT.md` (untracked,
under the fork's `dist/`) from a before/after diff of what it touched:

- per class: `match`, `client-scripted`, `mismatch`, `declared, differs` and
  `unclassified` counts before and after, and each row that moved into
  `mismatch`, `declared, differs` or `unclassified`, with its diff column;
- talents: nodes added, removed, or whose description, rank count or spell id
  changed between the previously followed build and this one, with the old and
  new text, because the engine's rate for that talent lives in hand-written Go
  the generators never touch;
- trainables: abilities new to a class's table that have no disposition;
- racials: any `forever_changes` text that changed in `races.json`, because
  `racials.go` is hand-written from that text;
- stat tables, relic mods and set rows: a one-line "regenerated, N rows
  changed" each, from the generated-file diff.

The report's first line is its verdict, machine-readable as `verdict: clean` or
`verdict: review`. It is `clean` only when all of the following hold:

- Row level: no row that was `match`, `client-scripted` or `declared, matches`
  before is `mismatch`, `declared, differs` or `unclassified` after. Rows are
  keyed by class, level, spell id and rank. A row that is new and already
  `match` is fine; a row that disappeared is fine (the client dropped the
  spell); a row that is new and not `match` is a worsening.
- Class level: for every class, the `mismatch`, `declared, differs` and
  `unclassified` counts after are each less than or equal to before. A class
  appearing or disappearing from the goldens is `review`.
- The talents section is empty: no node added, removed, or changed in
  description, rank count or spell id.
- The trainables section is empty: no new ability without a disposition.
- The racials section is empty: no `forever_changes` text changed.

Anything else is `review`. Counts of `match` or `client-scripted` moving down
while the row-level rule holds (a spell removed by the client) is still
`clean`. The report is what the workflow pastes into the step summary and the
pull request.

### The workflow: `engine.yml` in the site repo

The site repo is already the control plane (the client watcher dispatches every
consumer of a new build), so the workflow lives there and the fork stays a
library.

Triggers, mirroring `bis.yml`:

- `schedule` at 05:07 UTC, after the data lane's nightly commits and before the
  ranking at 06:23, so a follow that lands is ranked the same night.
- `workflow_dispatch` with `build` (default: the active build) and `force`.
- The client watcher's activate step dispatches it with `force=true` right
  after it dispatches the ranking, so a new build is followed the same morning
  rather than carried forward for a day.

Also a `dry_run` input (default false): run everything, push nothing, open
nothing, and upload the report and the generated diff as a workflow artifact.
This is how the workflow is proven before the fork token exists and how a
person previews a follow.

Concurrency group `engine`, `cancel-in-progress: false`, so a scheduled run and
the watcher's dispatch queue rather than race; the second one sees the hash
unchanged and skips. The ranking's own `bis` group does the same for the
`--force` dispatch the pin step sends: it queues behind the 06:23 run if one is
in flight, and whichever runs second finds the inputs unchanged and skips.

Skip-when-unchanged, the same shape as the ranking's: the job hashes the build
id together with the committed damage-relevant inputs for the build
(`spellconst/`, `talents/`, `trainables/`, `gametables/`, `levels.json`,
`sets.json`, `races.json`) and the fork's generator sources, and compares it
with the `Follow-Inputs:` trailer on the fork master's last follow commit.
Equal and not forced: skip with a notice. The raw tables are not hashed (they
are not in the tree); they are refetched and the relic mods and set rows
regenerated on every follow, and a new build always follows because its id is
in the hash. A follow whose only generated change is the build stamp still
commits (the 70338 follow was exactly that), so the fork's `FOREVER_BUILD`
always names the build the engine was last proven against.

Otherwise the job:

1. Checks out the site at the triggering commit and the fork's `master` into
   `engine/`; installs Go, the protoc release the committed headers name,
   protoc-gen-go v1.36.6, Python and uv.
2. Fetches the build's raw tables: `uv run python -m pipeline fetch --product
   <manifest product> --build $BUILD` in `data/`.
3. Runs `make follow SITE_DIR=$GITHUB_WORKSPACE BUILD=$BUILD` in `engine/`.
4. Runs the fork's suite exactly as the fork's CI runs it (the proto-match test,
   `make binary_dist/dist.go`, `go test --tags=with_db ./sim/...` and
   `./tools/...`, the wasm cross-compile), with `FOREVER_SITE_DIR` set so the
   fixture comparison runs rather than skips.
5. Reads the report's verdict.

**Verdict `clean` and the suite green**: commit in the fork as
`forever-bis-bot` with the message `chore(client): the engine follows client
<build>` and a `Follow-Inputs: <hash>` trailer, push to `master`, then go on to
the pin. The push uses the fork token, and a push made with a personal token
does trigger the fork's own `test.yml` (only `GITHUB_TOKEN` pushes are silent),
so the fork's CI runs on the same commit as a second gate. `test.yml` gains
`master` in its push branches so this and later human pushes are covered.

**Verdict `review`, or the suite red**: push the same commit to a branch
`follow/<build>` in the fork and open (or update, never duplicate) one pull
request titled `engine: follow client <build>` whose body is the report, the
test failures if any, and the one-line next step: finish the hand work on this
branch, merge, and the pin lands the next night (or on dispatch). The pin step
does not run. No issue is opened on the site; the pull request is the single
notification, and the step summary carries the same report.

The first follow is expected to be `review`: two builds of drift, including a
reworked Frost tree, is exactly the case the report exists for.

### The pin step

The pin is its own job in the same workflow so it also runs when the fork
advanced by hand (a review PR merged, or any engine fix) with nothing to
regenerate. It runs whenever fork `master` differs from the site's pin **and**
the fork's `FOREVER_BUILD` equals the active build **and** the fork's `test`
workflow has concluded successfully for that sha. A half-followed fork is never
pinned; neither is a commit the fork's own CI has not passed. If that CI run is
still in progress the job waits for it, up to 30 minutes; if it is absent or
failed, the job does not pin and the step summary says which of the three
conditions failed. Any engine change that meets the three conditions pins and
re-ranks that night: that is the intent, since an engine fix the owner merged
is a fix the site should carry.

1. In the site checkout, rewrite the engine replace to fork master's sha (the
   same `go mod edit -replace` the `pin-engine` action does) and run the site's
   sim suite (`go test ./...` in `sim/` with the active build's `make simdb`)
   and the data tests that read both pins (`uv run pytest tests/test_genproto.py`
   with `FOREVER_ENGINE_PATH` set to the fork checkout).
2. Green: run `make engine-pin ENGINE_DIR=engine` and `uv run python -m
   pipeline simproto --engine ../engine` in `data/`, commit as
   `forever-bis-bot` with `chore(engine): pin <sha> (the engine follows client
   <build>)`, push to `main` with the activate script's rebase-and-retry loop,
   and dispatch `web.yml`, `api.yml`, `sim.yml` and `bis.yml --force`, as the
   activate script does, because a bot push triggers nothing.
3. Red: open (or update) one pull request on the site, `engine: pin <sha>`,
   with the pin commit and the failing test output, so the ladder goldens or
   whatever moved can be regenerated by hand from a branch that already carries
   the pin. `main` keeps the old pin.

### The drift guard

`sim.yml`'s test job gains one step after `pin-engine`: read the pinned
module's `FOREVER_BUILD` and compare it with `web/src/data/active-build.json`.
Equal: nothing. Different: a `::warning` annotation naming both builds and the
engine workflow, and the same line in the step summary. It warns rather than
fails because the follow window between activation and pin is a normal state
for a few hours; a red `main` for every sim push during it would teach people
to ignore red.

### Tokens

The fork is public, so reading needs nothing. Pushing branches and opening pull
requests in the fork needs a fine-grained personal access token scoped to
`jhunthrop/wowsims-forever` with contents and pull-requests read/write, stored
as the site repo secret `ENGINE_REPO_TOKEN`. The owner creates it; the design
and the runbook name it, never its value. Fine-grained tokens expire (one year
at most); the runbook records the expiry month beside the name. The fork's
`master` is unprotected, so the direct push in the clean path needs nothing
more. Until the token exists the workflow runs only with `dry_run`. Everything
on the site side uses the job's own `GITHUB_TOKEN`, as the ranking and the
client watcher do today, plus the existing `WAGO_API_TOKEN` for the raw fetch.

### Runbook

`docs/engine-follow.md` in the site repo: what runs when, how to read the report,
the two pull-request shapes and what to do with each, how to run `make follow`
and `make engine-pin` by hand, how to force a follow, and the token's name and
expiry. It replaces nothing (there was no written procedure) and the memory file
that described the hand procedure points to it.

## What the owner sees

Nothing, on a night when the client changed only what the generators carry: the
fork advances, the site pins, the ranking runs, and the morning's picks are
from an engine that saw the current client. The step summary has the report
for anyone who looks.

One pull request in the fork, on a morning when a talent's text changed or a
spell stopped matching: its body is the full list of what the engine now gets
wrong and the talents whose rates to check, nothing more to find. Merging it is
the whole hand-off; the pin follows that night.

One pull request on the site, rarely, when a correct engine change moved a
number the site's own goldens pin: the branch already carries the pin, and the
failing test names the golden to regenerate.

## Constraints

- No secret value in any file, log or chat; the token is named only.
- Commits as `forever-bis-bot <bot@foreversixty.gg>`, conventional type and the
  repo's sentence-style description; the bot's commits carry no
  `Co-Authored-By` trailer, like the ranking's.
- The ranking's inputs hash already covers `sim/`, so a pin re-ranks without
  any change to `bis.yml`.
- `go test` on the fork needs `--tags=with_db`; the site's sim tests need
  `make simdb` first (compile-time embeds).
- Python invoked as `python3 -I` from outside the fork for the relic generator,
  as its make target already does.
- The follow never edits a hand-written file. A generator whose output requires
  a hand edit to compile is a defect in the generator, fixed in the fork, not
  patched in the script.

## Order of work

The fork's CI must be green on `master` before any automation pushes to it, so
the protoc fix and the `master` push trigger land first and are proven on a
real run. The follow script and the test refactors come next, proven locally
against the site's 70532 tables (the first real follow, which is expected to
end `review`). The workflow lands last, run in `dry_run` until the token
exists, then once for real.

## Left to the builder

- The exact before/after diff mechanics for the report (a copy of the
  pre-follow goldens and fixtures to a temp dir is fine; the counts come from
  the golden tables, keyed as the verdict rule says).
- Step ordering and caching in the workflow, the retry count on the push, and
  the pull-request update mechanics (`gh pr edit` on an existing PR for the
  same branch, never a second PR).
- How `make follow` locates `gen_set_rows.py` (it is in the site repo; the
  script runs it with the site dir it was given).
- How the protoc release is installed on the runner (a pinned GitHub release
  download is fine) and how the version is kept in one place the test and the
  workflow both read.
