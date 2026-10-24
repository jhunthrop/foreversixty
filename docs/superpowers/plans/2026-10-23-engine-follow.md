# Engine Follow Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** The engine fork regenerates from each new client build and the site pins the result the same night, with one review pull request whenever a change needs a person.

**Architecture:** One script in the fork (`tools/follow.sh`, reached as `make follow`) runs every generator against a site build directory and writes a verdict-bearing report via a Go command (`tools/followreport`). One workflow on the site (`engine.yml`) fetches the raw tables, runs the follow, pushes a clean follow to fork `master` or opens a review PR, and in a second job pins fork `master` into the site when three conditions hold. The fork's CI is made green first so the automation has a gate to stand on.

**Tech Stack:** Go (fork tools and tests, `--tags=with_db`), bash (fork script, site workflow steps), Python 3 via `python3 -I` (existing generators, site data pipeline via uv), GitHub Actions, `gh`.

**Spec:** `docs/superpowers/specs/2026-10-23-engine-follow-design.md`

## Global Constraints

- Two checkouts: the fork at `/Users/jh/code/wowsims-forever` on branch `master` (the integration branch; `forever` is a stale mirror and is not touched), the site at `/Users/jh/code/forever` on `main`.
- Fork module path is `github.com/wowsims/classic`; the site consumes it via `replace github.com/wowsims/classic => github.com/jhunthrop/wowsims-forever@<sha>`.
- Fork tests run with `go test --tags=with_db`; the site's sim tests need `make simdb` (compile-time embeds) first.
- Committed `.pb.go` files were generated with `protoc v36.2` and `protoc-gen-go v1.36.6`; release asset `protoc-36.2-linux-x86_64.zip` at `https://github.com/protocolbuffers/protobuf/releases/download/v36.2/`.
- Bot identity for every automated commit: `forever-bis-bot <bot@foreversixty.gg>`; conventional type, sentence-style description, no `Co-Authored-By` trailer on bot commits. Human commits in this plan carry `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`.
- Commit messages: fork follow `chore(client): the engine follows client <build>` with trailer `Follow-Inputs: <hash>`; site pin `chore(engine): pin <sha> (the engine follows client <build>)`.
- The follow never edits a hand-written file; `git add` stages only the paths the script lists, never `-A`.
- The report's first line is exactly `verdict: clean` or `verdict: review`.
- Secrets are named, never valued: `ENGINE_REPO_TOKEN` (site secret, fine-grained PAT on `jhunthrop/wowsims-forever`, contents + pull-requests RW). No token is needed to read the fork or to fetch from wago.tools.
- Workflow schedule `7 5 * * *` UTC, concurrency group `engine`, `cancel-in-progress: false`.
- Python generators invoked as `python3 -I` from outside the fork directory, as the fork's `relicmods` target does.
- Nothing in this plan commits a follow's regenerated output to fork `master`; the first real follow lands through the workflow (expected verdict `review`).

## Review Focus

1. A talent tree change removes a proto field that hand-written Go names, so `go build ./sim/...` fails after `make talents`. Expected: the script still writes the report, verdict `review`, with a `## Build` section first. Test in Task 4 (`follow.sh` build-failure path) and Task 3 (`Render` with a non-empty build log).
2. A class whose golden has no `Disposition` column (no dispositions map) gains a new trainable row. Expected: counted as "no disposition", verdict `review`. Test in Task 3.
3. The site build directory for the previously followed build no longer exists (70425 is gone today). Expected: the talent diff uses the fork's `tools/talentgen/testdata` copies and never touches the site for the previous build. Test in Task 4 (dry run against 70532 with no 70425 directory present).
4. Two `engine.yml` runs queue (schedule plus watcher dispatch). Expected: the second sees the `Follow-Inputs` hash unchanged and skips with a notice, pushing nothing. Covered by the hash script's test in Task 5 (same inputs produce the same hash; the trailer lookup finds the commit).
5. Fork `master` has a newer commit than the pin but its `test` workflow has not run or failed. Expected: the pin job does not pin and the step summary names the failed condition. Test in Task 5 (`pin-conditions.sh` with a stubbed `gh`).

---

### Task 1: Fork CI runs on `master` with the protoc the committed protos name

**Repo:** fork (`/Users/jh/code/wowsims-forever`, branch `master`)

**Files:**
- Create: `.protoc-version` (content: `36.2`)
- Modify: `sim/core/proto/generated_test.go`
- Modify: `makefile` (the `sim/core/proto/api.pb.go` rule's guard)
- Modify: `.github/workflows/test.yml`

**Interfaces:**
- Produces: `.protoc-version`, read by the makefile guard, the test, and later by the site workflow (Task 5) to install the same release.

- [ ] **Step 1: Write the failing test** in `sim/core/proto/generated_test.go`

```go
func TestProtocVersionIsPinned(t *testing.T) {
	pinned, err := pinnedProtocVersion() // reads ../../../.protoc-version, trimmed
	if err != nil { t.Fatal(err) }
	if pinned != "36.2" { t.Errorf("pinned = %q, want the version the committed .pb.go headers name, 36.2", pinned) }
	if got := protocVersionFromOutput("libprotoc 36.2\n"); got != "36.2" { t.Errorf("parse = %q", got) }
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/proto/ -run TestProtocVersionIsPinned`
Expected: FAIL (undefined `pinnedProtocVersion`).

- [ ] **Step 3: Implement** `pinnedProtocVersion() (string, error)` and `protocVersionFromOutput(s string) string` in `generated_test.go`, and make `TestGeneratedProtosMatchSources` call `protoc --version` first: if the installed version differs from the pinned one, `t.Skipf` locally and `t.Fatalf` when `os.Getenv("CI") != ""`, with the message `protoc %s is installed; the committed .pb.go were generated with protoc %s (.protoc-version). Install that release; the files are not stale.`

- [ ] **Step 4: Guard the makefile.** In the `sim/core/proto/api.pb.go: proto/*.proto` rule, before `protoc -I=...`, add a check that `protoc --version` ends with the content of `.protoc-version`, else print `protoc $(installed) installed; this repository pins $(pinned) (.protoc-version)` and exit 1.

- [ ] **Step 5: Fix the workflow.** In `.github/workflows/test.yml`: `on.push.branches: [master, forever]`; replace the `apt-get install protobuf-compiler` line with a step that reads `VER=$(cat .protoc-version)`, downloads `protoc-$VER-linux-x86_64.zip` from the release URL in Global Constraints, unzips into `$HOME/.local`, and appends `$HOME/.local/bin` to `$GITHUB_PATH`. Keep `protoc-gen-go v1.36.6`.

- [ ] **Step 6: Run the proto test locally**

Run: `go test --tags=with_db ./sim/core/proto/ -v -run 'TestGeneratedProtosMatchSources|TestProtocVersionIsPinned'`
Expected: both PASS (local protoc is 36.2).

- [ ] **Step 7: Commit and push to master, then watch CI**

```bash
git add .protoc-version sim/core/proto/generated_test.go makefile .github/workflows/test.yml
git commit -m "ci: the test workflow runs on master with the protoc the committed protos name" -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
git push origin master
gh run watch -R jhunthrop/wowsims-forever "$(gh run list -R jhunthrop/wowsims-forever -w test --limit 1 --json databaseId --jq '.[0].databaseId')" --exit-status
```
Expected: the `test` workflow runs on the master push and concludes success. If any other test is red on master, fix it in this task before moving on: the spec's order of work needs a green master first.

---

### Task 2: The build comes from the fixtures, never a literal; `FOREVER_BUILD` names the followed build

**Repo:** fork

**Files:**
- Create: `assets/db_inputs/FOREVER_BUILD` (content: `1.60.1.70425`)
- Modify: `makefile` (`BUILD`, `TALENT_BUILD` defaults)
- Modify: `sim/core/talents/talents_test.go:35-45`
- Modify: `tools/talentgen/main_test.go:115-125, 325-345`

**Interfaces:**
- Produces: `assets/db_inputs/FOREVER_BUILD`, one line, the build every generated file is stamped for. Read by `make follow` (Task 4) and the site's pin conditions and drift guard (Tasks 5, 6).
- Produces: env var `FOREVER_SITE_DIR` honoured by `TestFixturesMatchTheSiteData`.

- [ ] **Step 1: Write the failing tests**

In `tools/talentgen/main_test.go`, replace the literal in `TestGoFileCarriesSizesBuildAndIDs` with a value read from the fixture:
```go
build := fixtureClass(t, "warrior").Build
want := fmt.Sprintf(`const TalentsBuild = %q`, build)
```
and rewrite `TestFixturesMatchTheSiteData` so `siteDir := os.Getenv("FOREVER_SITE_DIR")` (skip with `FOREVER_SITE_DIR is not set; the fixtures cannot be compared with their source here` when empty), `siteTalents := filepath.Join(siteDir, "data", "builds", build, "talents")` with `build` from the warrior fixture, and skip with `the site checkout has no %s; nothing to compare` when that directory is absent.

In `sim/core/talents/talents_test.go`, `TestTreeSizesMatchTheClient` asserts `c.Build` equals the `build` field read from the fixture file the `warrior(t)` helper loads (a two-field `json.Unmarshal` of the same bytes), and that each `c.TreeSizes()[i]` equals `len(c.Trees[i].Talents)` with the three summing to the class's node count; the `{17,17,18}` literal and the `"1.60.1.70425"` literal are gone.

- [ ] **Step 2: Run to verify they fail** (they fail to compile until the helpers exist, or fail on the literal mismatch after you change the fixture build in a scratch copy)

Run: `go test --tags=with_db ./sim/core/talents/ ./tools/talentgen/`

- [ ] **Step 3: Implement the fixture reads** (a `fixtureBuild(t, path) string` helper in each test package).

- [ ] **Step 4: Makefile.** Add `FOREVER_BUILD := $(shell cat assets/db_inputs/FOREVER_BUILD)` near the other variables and set `BUILD ?= $(FOREVER_BUILD)` and `TALENT_BUILD ?= $(FOREVER_BUILD)`, deleting the `1.60.1.70338` and `1.60.1.69893` literals and their comments' claims about defaults.

- [ ] **Step 5: Run the tests both ways**

Run: `go test --tags=with_db ./sim/core/talents/ ./tools/talentgen/` → PASS, with `TestFixturesMatchTheSiteData` SKIP (no `FOREVER_SITE_DIR`).
Run: `FOREVER_SITE_DIR=/Users/jh/code/forever go test --tags=with_db ./tools/talentgen/ -run TestFixturesMatchTheSiteData -v` → SKIP with the "has no 1.60.1.70425" message (that build directory is gone from the site).

- [ ] **Step 6: Commit**

```bash
git add assets/db_inputs/FOREVER_BUILD makefile sim/core/talents/talents_test.go tools/talentgen/main_test.go
git commit -m "refactor(test): the build comes from the fixtures, never a literal; FOREVER_BUILD names the followed build" -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 3: `tools/followreport`: the follow report and its verdict

**Repo:** fork

**Files:**
- Create: `tools/followreport/main.go` (flags, wiring, exit codes)
- Create: `tools/followreport/golden.go` (parse goldens)
- Create: `tools/followreport/compare.go` (row/class deltas and verdict)
- Create: `tools/followreport/talents.go` (talent and racial diffs)
- Create: `tools/followreport/render.go` (markdown)
- Create: `tools/followreport/*_test.go` with `testdata/` fixtures (small goldens, talents, races)

**Interfaces:**
- Produces the command `go run ./tools/followreport -before <dir> -after <dir> -prev-talents <dir> -talents <dir> -prev-races <file> -races <file> -build-log <file|''> -generated-diffstat <file|''> -out <file>`; exit 0 on `clean`, exit 3 on `review`, exit 1 on a parse error.
- Produces package functions used by the tests and by nothing else:
  - `ParseGolden(path string) (Golden, error)` where `Golden{Class string; Rows map[RowKey]Row; Trainables map[string]Trainable}`, `RowKey{Section, Spec string; Level int; SpellID int; Rank int}` (Section is `spells` or `talent-gated`), `Row{Verdict, Damage string}`, `Trainable{Name, Disposition string; HasDispositionColumn bool}`.
  - `CompareClass(before, after Golden) ClassDelta` with `Worsened []RowChange`, `Counts map[string][2]int` for `mismatch`, `declared, differs`, `unclassified`, and `NewUndisposed []string`.
  - `LoadTalents(dir string) (map[string]TalentsFile, error)` (one per class slug; the site's `talents/<class>.json` shape: `trees[].talents[]{id,name,max_rank,spell_id,ranks[]{spell_id,description}}`), `DiffTalents(prev, cur map[string]TalentsFile) TalentDelta` with `Added, Removed, Changed []TalentChange` keyed by node id; changed when `max_rank`, `spell_id`, or any rank's `description` differs.
  - `DiffRacials(prev, cur string) ([]RacialChange, error)` reading `races.json` and comparing each race's `forever_changes[].text` as a set.
  - `Verdict(classes map[string]ClassDelta, t TalentDelta, r []RacialChange, buildLog string, classesBefore, classesAfter []string) string` returning `"clean"` or `"review"` per the spec's rule.
  - `Render(r Report) string` whose first line is `verdict: <v>` and whose first section is `## Build` when the build log is non-empty.

- [ ] **Step 1: Write the failing tests** (`compare_test.go`, table-driven), one case per spec rule, using fixtures under `testdata/`:

```go
cases := []struct{ name, before, after string; wantVerdict string }{
	{"identical goldens", "mage-a", "mage-a", "clean"},
	{"a match row becomes mismatch", "mage-a", "mage-one-mismatch", "review"},
	{"a new row already match", "mage-a", "mage-plus-match-row", "clean"},
	{"a row the client dropped", "mage-a", "mage-minus-row", "clean"},
	{"a new row that is mismatch", "mage-a", "mage-plus-mismatch-row", "review"},
	{"declared, matches becomes declared, differs", "mage-a", "mage-damage-differs", "review"},
	{"a new trainable in a class with dispositions, unclassified", "mage-a", "mage-new-unclassified", "review"},
	{"a new trainable in a class with no Disposition column", "rogue-a", "rogue-new-trainable", "review"},
	{"a class disappears from the goldens", "two-classes", "one-class", "review"},
}
```
plus `TestDiffTalents` (added, removed, description changed, max_rank changed, spell_id changed, unchanged → empty), `TestDiffRacials` (text changed → one change; identical → none), `TestVerdictBuildLogIsReview`, and `TestRenderFirstLine` asserting `strings.HasPrefix(out, "verdict: clean\n")` and that a non-empty build log puts `## Build` before any class section.

- [ ] **Step 2: Run to verify they fail**

Run: `go test ./tools/followreport/`
Expected: FAIL (package does not build yet).

- [ ] **Step 3: Implement the five files** to the interfaces above. Golden parsing: find `## Spells`, `## Talent-gated spells`, `## Trainable abilities the engine does not register`; read each markdown table by its header row, locating columns by header text (`Spec`, `Level`, `SpellID`, `Rank`, `Verdict`, `Damage`, `Ability`, `Disposition`), never by index. Class name from the H1 (`# Mage: ...`).

- [ ] **Step 4: Run to verify they pass**

Run: `go test ./tools/followreport/ -cover`
Expected: PASS, coverage ≥ 80%.

- [ ] **Step 5: Commit**

```bash
git add tools/followreport
git commit -m "feat(tools): followreport diffs a follow's goldens, talents and racials and names a verdict" -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 4: `tools/follow.sh` and `make follow`

**Repo:** fork

**Files:**
- Create: `tools/follow.sh`
- Modify: `makefile` (add `follow` target)
- Modify: `PORTING.md` (one paragraph: the follow sequence lives in `tools/follow.sh`; run `make follow SITE_DIR=... BUILD=...`)
- Create: `sim/core/testdata/conformance/client/races.json` (a tracked copy of the site's `races.json`, so the racial diff has a previous build to compare against)

**Interfaces:**
- Consumes: `tools/followreport` (Task 3), `assets/db_inputs/FOREVER_BUILD` (Task 2), existing targets `talents`, `proto`, `spellconst`, `relicmods`, scripts `tools/levels_from_site.py`, `tools/base_stats_parser.py`, and the site's `tools/gen_set_rows.py`.
- Produces: `make follow SITE_DIR=<site> BUILD=<build>`; writes `dist/FOLLOW-REPORT.md` and `dist/FOLLOW-PATHS` (one path per line, the allowlist actually written); exit 0 on `clean`, 3 on `review`, 1 on a generator failure. Prints `verdict: <v>` as its last line.

- [ ] **Step 1: Write the script.** Sequence, each step aborting on failure except where noted:
  1. Refuse a dirty checkout (`git status --porcelain` non-empty → exit 1).
  2. `PREV=$(cat assets/db_inputs/FOREVER_BUILD 2>/dev/null)`; when missing, seed from the `build` field of `tools/talentgen/testdata/warrior.json`.
  3. Snapshot to `$TMP`: `sim/core/testdata/conformance/*.golden.md`, `tools/talentgen/testdata/*.json`, `sim/core/testdata/conformance/client/races.json`.
  4. Copies: `$SITE/data/builds/$BUILD/talents/*.json → tools/talentgen/testdata/`; `spellconst/*.json → sim/core/testdata/conformance/client/`; `trainables/*.json → .../client/trainables/`; `races.json → .../client/races.json`; `gametables/* → assets/db_inputs/gametables/$BUILD/`.
  5. `make talents TALENT_BUILD=$BUILD SITE_DIR=$SITE` (includes `make proto`).
  6. `make spellconst BUILD=$BUILD SITE_DIR=$SITE`.
  7. `python3 -I tools/levels_from_site.py $SITE/data/builds/$BUILD/levels.json assets/db_inputs/levels/$BUILD.json`; `python3 -I tools/base_stats_parser.py --inputs assets/db_inputs/gametables/$BUILD --build $BUILD --levels-json assets/db_inputs/levels/$BUILD.json`.
  8. `make relicmods BUILD=$BUILD SITE_DIR=$SITE`.
  9. `(cd /tmp && python3 -I $SITE/tools/gen_set_rows.py $SITE/data/builds/$BUILD $FORK)`.
  10. `go build --tags=with_db ./sim/... 2> $TMP/build.log`: on failure, keep going with `BUILD_LOG=$TMP/build.log` and skip step 11.
  11. `FOREVER_UPDATE_GOLDEN=1 go test --tags=with_db ./sim/conformance/...`.
  12. `echo $BUILD > assets/db_inputs/FOREVER_BUILD`.
  13. Stamp asserts: every regenerated file names `$BUILD` in its header (`talents_auto_gen.go`: `build $BUILD`; `constants_auto_gen.go`, `base_stats_auto_gen.go`, `relic_mods_auto_gen.go`: `Client build: $BUILD`; `client_sets_gen.go`: `client build $BUILD`; goldens when regenerated: `Client build: $BUILD`).
  14. `go run ./tools/followreport -before $TMP/goldens -after sim/core/testdata/conformance -prev-talents $TMP/talents -talents tools/talentgen/testdata -prev-races $TMP/races.json -races sim/core/testdata/conformance/client/races.json -build-log "${BUILD_LOG:-}" -generated-diffstat <(git diff --numstat -- <allowlist>) -out dist/FOLLOW-REPORT.md`.
  15. Write `dist/FOLLOW-PATHS` from the allowlist (the exact paths in the spec's table plus `FOREVER_BUILD` and `client/races.json`); fail if `git status --porcelain` names a path outside it.
  16. Print the verdict line; exit 0 or 3.

- [ ] **Step 2: Add the make target**

```make
.PHONY: follow
follow:
	tools/follow.sh "$(SITE_DIR)" "$(BUILD)"
```

- [ ] **Step 3: Prepare the site's raw tables for 70532** (gitignored output, no commit)

Run: `cd /Users/jh/code/forever/data && uv run python -m pipeline fetch --product wowf --build 1.60.1.70532`
Expected: `data/builds/1.60.1.70532/raw/*.csv` present (Spell.csv, SpellEffect.csv, ItemEffect.csv, ItemXItemEffect.csv, SpellName.csv, ItemSet.csv, ItemSetSpell.csv among them).

- [ ] **Step 4: Run the first follow in a throwaway worktree** (the spec forbids committing a follow's output here)

```bash
cd /Users/jh/code/wowsims-forever && git worktree add /tmp/follow-70532 master
cd /tmp/follow-70532 && make follow SITE_DIR=/Users/jh/code/forever BUILD=1.60.1.70532; echo "exit $?"
head -3 dist/FOLLOW-REPORT.md; wc -l dist/FOLLOW-PATHS
git status --porcelain | awk '{print $2}' | sort > /tmp/changed; sort dist/FOLLOW-PATHS | comm -23 /tmp/changed - 
```
Expected: exit 3 (`verdict: review`, two builds of drift including the Frost tree rework), the report's talents section non-empty, `comm` prints nothing (no path changed outside the allowlist). If `go build` failed in step 10, the report opens with `## Build`. Record in the ledger whether the build passed, since it tells the owner how much hand work the first review PR carries.

- [ ] **Step 5: Run the fork suite in the worktree** to learn the gate's outcome

Run: `cd /tmp/follow-70532 && FOREVER_SITE_DIR=/Users/jh/code/forever go test --tags=with_db ./sim/core/proto/ -run TestGeneratedProtosMatchSources && go test --tags=with_db -count=1 ./sim/... ./tools/... 2>&1 | tail -20`
Expected: whatever it is; the result goes in the ledger. Then `cd /Users/jh/code/wowsims-forever && git worktree remove --force /tmp/follow-70532`.

- [ ] **Step 6: Commit the script (not the follow's output)**

```bash
cd /Users/jh/code/wowsims-forever
git add tools/follow.sh makefile PORTING.md sim/core/testdata/conformance/client/races.json
git commit -m "feat(tools): make follow regenerates the engine from one site build and writes the follow report" -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
git push origin master
```
Expected: fork `test` workflow green on the push (the races.json copy and the script change no Go).

---

### Task 5: `engine.yml` on the site: follow job, pin job, dry run

**Repo:** site (`/Users/jh/code/forever`)

**Files:**
- Create: `.github/workflows/engine.yml`
- Create: `tools/engine/follow-inputs-hash.sh`
- Create: `tools/engine/pin-conditions.sh`
- Create: `tools/engine/tests/test_scripts.sh` (bats-free bash tests runnable with `bash tools/engine/tests/test_scripts.sh`)

**Interfaces:**
- Consumes: fork `make follow` and its `dist/FOLLOW-REPORT.md`, `dist/FOLLOW-PATHS` (Task 4); `.protoc-version` (Task 1); `assets/db_inputs/FOREVER_BUILD` (Task 2); site `make engine-pin ENGINE_DIR=<fork>`; `uv run python -m pipeline simproto --engine <fork>` in `data/`; `uv run python -m pipeline fetch --product <p> --build <b>`.
- Produces: `tools/engine/follow-inputs-hash.sh <site-dir> <fork-dir> <build>` prints a 16-hex hash of the build id, `data/builds/<build>/{spellconst,talents,trainables,gametables}` trees, `levels.json`, `sets.json`, `races.json` (via `git ls-tree -r HEAD` on the site) and the fork's generator sources (`git rev-parse HEAD:tools HEAD:sim/core/spellconst/gen` in the fork). Also `--last <fork-dir>` prints the `Follow-Inputs:` trailer of the newest fork master commit that has one, or `none`.
- Produces: `tools/engine/pin-conditions.sh <site-dir> <fork-dir> <active-build>` exits 0 when fork master sha ≠ `sim/enginever/version.go` sha prefix, fork `FOREVER_BUILD` == active build, and `gh run list -R jhunthrop/wowsims-forever -w test --commit <sha> --json conclusion` reports `success` (polling every 60 s up to 30 min while `in_progress`); otherwise exits 1 and prints one line per failed condition. `GH_BIN` env overrides the `gh` command for tests.

- [ ] **Step 1: Write the failing script tests** in `tools/engine/tests/test_scripts.sh`: (a) hash is stable across two calls and changes when a byte of a `spellconst` file changes in a scratch clone; (b) `--last` prints `none` on a repo with no trailer and the hash on one with a `Follow-Inputs:` trailer commit; (c) `pin-conditions.sh` with a stub `gh` printing `[{"conclusion":"failure"}]` exits 1 and names the CI condition; with `success` and matching build exits 0; with a mismatched `FOREVER_BUILD` exits 1 naming the build condition.

- [ ] **Step 2: Run to verify they fail**: `bash tools/engine/tests/test_scripts.sh` → fails (scripts absent).

- [ ] **Step 3: Implement the two scripts** to the interfaces above.

- [ ] **Step 4: Run the tests**: `bash tools/engine/tests/test_scripts.sh` → all pass.

- [ ] **Step 5: Write `engine.yml`.** Header comment in the house style (what it does, when, why the schedule). `on`: `schedule: '7 5 * * *'`, `workflow_dispatch` inputs `build` (string, default ''), `force` (boolean, default false), `dry_run` (boolean, default false). `permissions: { contents: read }`, `concurrency: { group: engine, cancel-in-progress: false }`.

  Job `follow` (`permissions: contents: write, actions: write`; `timeout-minutes: 90`):
  1. checkout site; checkout `jhunthrop/wowsims-forever@master` into `engine/` with `fetch-depth: 0`.
  2. setup-go (fork `go.mod`), protoc from `engine/.protoc-version` (same download step as Task 1), `protoc-gen-go@v1.36.6`, setup-uv + `uv sync` in `data/`.
  3. `BUILD` = input or `web/src/data/active-build.json`; `PRODUCT` = `manifest.json` product.
  4. `HASH=$(tools/engine/follow-inputs-hash.sh . engine "$BUILD")`, `LAST=$(... --last engine)`; skip with `::notice` when equal and not forced.
  5. `uv run python -m pipeline fetch --product "$PRODUCT" --build "$BUILD"` in `data/`.
  6. `make follow SITE_DIR=$GITHUB_WORKSPACE BUILD=$BUILD` in `engine/`; capture exit code (0 or 3) without failing the step; `cat engine/dist/FOLLOW-REPORT.md >> $GITHUB_STEP_SUMMARY`.
  7. Fork suite as the fork's `test.yml` runs it, in `engine/`, with `FOREVER_SITE_DIR=$GITHUB_WORKSPACE`; capture pass/fail.
  8. Upload `engine/dist/FOLLOW-REPORT.md` and `git -C engine diff` as artifact `follow-<build>`.
  9. If `dry_run`: stop here, `::notice::dry run; nothing pushed`.
  10. Commit in `engine/` as the bot, staging exactly `xargs git add < dist/FOLLOW-PATHS`, message per Global Constraints with `Follow-Inputs: $HASH`.
  11. Clean and suite green: `git push https://x-access-token:${{ secrets.ENGINE_REPO_TOKEN }}@github.com/jhunthrop/wowsims-forever HEAD:master`.
  12. Otherwise push `HEAD:refs/heads/follow/$BUILD` (force) and `gh pr create` or `gh pr edit` (when `gh pr list --head follow/$BUILD` finds one) with `GH_TOKEN=${{ secrets.ENGINE_REPO_TOKEN }}`, title `engine: follow client $BUILD`, body = report + test tail + the spec's one-line next step.

  Job `pin` (`needs: follow`, runs `if: always() && !inputs.dry_run`, `permissions: contents: write, actions: write`):
  1. checkout site and fork master (fresh, after the follow job's push).
  2. `tools/engine/pin-conditions.sh . engine "$BUILD"` → stop with the summary line when it exits 1.
  3. In `sim/`: `go mod edit -replace github.com/wowsims/classic=github.com/jhunthrop/wowsims-forever@$SHA && go mod tidy`; `make simdb` at root; `go test ./... -timeout 40m`. In `data/`: `FOREVER_ENGINE_PATH=$GITHUB_WORKSPACE/engine uv run pytest -q --no-cov tests/test_genproto.py`.
  4. Green: `make engine-pin ENGINE_DIR=engine`; `cd data && uv run python -m pipeline simproto --engine ../engine`; bot commit `chore(engine): pin $SHA (the engine follows client $BUILD)`; push to `main` with the rebase-and-retry loop copied from `bis.yml`; `gh workflow run web.yml api.yml sim.yml --ref main` and `gh workflow run bis.yml --ref main -f force=true`.
  5. Red: push `HEAD:refs/heads/engine-pin/$SHA` and `gh pr create`/`edit` on the site, title `engine: pin $SHA`, body = failing test tail.

- [ ] **Step 6: Commit and push the workflow, then dry run it**

```bash
git add .github/workflows/engine.yml tools/engine
git commit -m "ci(engine): the engine follows the client build nightly and pins when the fork is proven" -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
git push origin main
gh workflow run engine.yml --ref main -f dry_run=true -f force=true
gh run watch "$(gh run list -w engine.yml --limit 1 --json databaseId --jq '.[0].databaseId')" --exit-status; gh run download -n follow-1.60.1.70532 -D /tmp/follow-artifact && head -40 /tmp/follow-artifact/FOLLOW-REPORT.md
```
Expected: the follow job completes, the step summary carries the report with `verdict: review`, nothing is pushed, the pin job reports "dry run" and exits cleanly. Fix and rerun until this holds.

---

### Task 6: Activation hook, drift guard, runbook

**Repo:** site

**Files:**
- Modify: `tools/activate-build.sh:170-176` (the dispatch block)
- Modify: `.github/workflows/sim.yml` (test job, after `pin-engine`)
- Create: `docs/engine-follow.md`
- Modify: `.github/workflows/client.yml` header comment (names the engine dispatch among the consumers)

**Interfaces:**
- Consumes: `engine.yml` (Task 5), fork `FOREVER_BUILD` (Task 2) read from the pinned module dir via `go list -m -f '{{.Dir}}' github.com/wowsims/classic`.

- [ ] **Step 1: Activation hook.** Add to the `gh workflow run` block in `tools/activate-build.sh`: `gh workflow run engine.yml --ref main -f force=true && echo "dispatched the engine follow for $NEW"`, and the matching line in `client.yml`'s header comment.

- [ ] **Step 2: Drift guard.** In `sim.yml`'s `test` job, after `pin-engine`, a step `the engine follows the active build`:
```bash
dir=$(go list -m -f '{{.Dir}}' github.com/wowsims/classic)
engine=$(cat "$dir/assets/db_inputs/FOREVER_BUILD" 2>/dev/null || echo unknown)
active=$(sed -n 's/.*"build"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' ../web/src/data/active-build.json)
if [ "$engine" != "$active" ]; then
  echo "::warning title=engine behind the client::the pinned engine is stamped for $engine; the site serves $active. engine.yml follows it nightly or on dispatch."
  echo "engine $engine, active $active: the engine is behind the client" >> "$GITHUB_STEP_SUMMARY"
fi
```

- [ ] **Step 3: Runbook** `docs/engine-follow.md`: what runs when (05:07 UTC and on activation); the two verdicts; what a `follow/<build>` PR and an `engine-pin/<sha>` PR each mean and what to do; running `make follow` and `make engine-pin` by hand; forcing a follow (`gh workflow run engine.yml -f force=true`); the dry run; `ENGINE_REPO_TOKEN` (name, scopes, expiry month to be filled in by the owner when created); why the drift guard warns rather than fails.

- [ ] **Step 4: Verify** the activate script still parses (`bash -n tools/activate-build.sh`) and `sim.yml` is valid YAML (`python3 -I -c 'import yaml,sys; yaml.safe_load(open(".github/workflows/sim.yml"))'`), then push and confirm the sim workflow's run on main shows the warning (the engine is stamped 70425, the site 70532).

- [ ] **Step 5: Commit**

```bash
git add tools/activate-build.sh .github/workflows/sim.yml .github/workflows/client.yml docs/engine-follow.md
git commit -m "feat(ci): activation dispatches the engine follow, the sim suite warns when the engine trails the client, and the runbook says how it all moves" -m "Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
git push origin main
```

---

### Task 7: First real run (owner-gated)

**Repo:** site

The owner creates the fine-grained token and adds it as `ENGINE_REPO_TOKEN`. This task cannot start until then; the plan's executor stops here and says so.

- [ ] **Step 1:** `gh workflow run engine.yml --ref main -f force=true` and watch it.
- [ ] **Step 2:** Expected: a pull request `engine: follow client 1.60.1.70532` opens in the fork with the report; no pin lands (the pin job names the `FOREVER_BUILD` condition). Report the PR link and the report's headline counts to the owner.
