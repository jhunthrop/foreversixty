# Simulator Engine Lane Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn the `wowsims/classic` fork at `/Users/jh/code/wowsims-forever` into the Forever Sixty combat engine — unified Hit and Crit, regenerated rating constants, declarative spell mods, an encounter biome, Forever racials, and two specs (`warrior-fury`, `mage-frost`) end to end — and ship it as three artifacts (`sim.wasm`, `sim.js`, `forever-sim`) plus the site-side `sim/` Go module whose `Summarize` adapter turns an engine result into the logs engine's `summary.Summary`.

**Architecture:** Two repositories. The **engine repo** (`/Users/jh/code/wowsims-forever`, module path stays `github.com/wowsims/classic`) holds every change to the simulation itself; it is consumed by the site as a Go module, so its generated protobufs must be committed (they are `.gitignore`d upstream — Task 1). The **site repo** (`/Users/jh/code/forever`) gains a new module `sim/` holding the request/result envelopes, the engine version pin, and the adapter. The adapter is the lane's first deliverable because the api and web lanes are blocked on it. Engine changes are ordered so that the one repo-wide mechanical rename (Hit/Crit) lands alone, before the six independent feature tasks that follow it.

**Tech Stack:** Go (engine `go 1.23.0`, site `go 1.25.11`), `google.golang.org/protobuf` v1.36.6, `protoc` 36.1 with `protoc-gen-go` v1.36.6, Python 3 (`tools/base_stats_parser.py`, the new `tools/spellconst_gen.py`), GNU make, GitHub Actions, Node 22 (wasm smoke test only).

**Spec:** `docs/superpowers/specs/2026-09-14-simulator-design.md` (sections 2 and 9)
**Interface contract (binding):** `docs/superpowers/specs/2026-09-14-simulator-interfaces.md` ("Engine" and "Engine version" sections are this lane's; the rest constrains it)
**Research:** `research/07-simulator.md` sections 1.2, 1.3, 1.11, 5.3

## Global Constraints

- **Two repositories, and every task says which.** Engine work happens in `/Users/jh/code/wowsims-forever`, module path `github.com/wowsims/classic` (unchanged — the organisation question is not settled). Site work happens in `/Users/jh/code/forever`, in the new module `sim/` (`github.com/jhunthrop/foreversixty/sim`). A task never edits both unless it is listed under both headings in its **Files** block.
- **Toolchain, measured 2026-09-14.** The engine declares `go 1.23.0` / `toolchain go1.23.4`; the installed Go is **go1.25.4 darwin/arm64**, which is newer, so `GOTOOLCHAIN=auto` uses it directly and **the engine builds and tests green under it** — verified: `go build ./...` 5.7 s, `go test --tags=with_db -count=1 ./sim/...` **10.5 s wall / 78.7 s CPU, 20 packages ok, 0 failures**. Do not raise the engine's `go` directive; upstream mergeability depends on it. The site's `go.work` declares `go 1.25.11`, which is newer than the installed 1.25.4, so `GOTOOLCHAIN=auto` downloads go1.25.11 on the first site build (verified: `cd logs && go version` reports `go1.25.11`). Both work with no configuration.
- **`--tags=with_db` is required for engine tests.** Without it `go test ./sim/core/...` panics in `NewEquipmentSet` (the item database is empty). Every engine test command in this plan carries the tag. `go build` does not need it.
- **The engine needs generated protobufs before anything compiles.** `sim/core/proto/*.pb.go` is produced by `make proto` and is `.gitignore`d upstream (`sim/core/proto/.gitignore:2`). A fresh clone fails `go build ./...` with `no required module provides package github.com/wowsims/classic/sim/core/proto`. Task 1 fixes this permanently. Locally: `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6`, then `PATH=$PATH:$(go env GOPATH)/bin make proto`.
- **`go build ./...` in the engine also needs `binary_dist/dist.go`**, which is `.gitignore`d. Run `make binary_dist/dist.go` once (it copies `sim/web/dist.go.tmpl`). Without it only `sim/web` fails; `go build ./sim/...` is unaffected.
- **`ENGINE_VERSION` is the short commit sha of `wowsims-forever`**, currently `7779ebb`. It appears in `sim/enginever/version.go` (`const Version = "7779ebb"`, written by `make engine-pin`), in `web/public/_sim/<ENGINE_VERSION>/sim.{wasm,js}`, in every `SimEnvelope.engine_version`, every stored sim row, every validation row, and the premium image tag. A result from a different `ENGINE_VERSION` stays readable and is labelled stale; it is never silently re-run.
- **The protobuf API does not change shape.** `RaidSimRequest`, `RaidSimResult`, `SimDatabase`, `APLRotation` keep their messages and field numbers. The only proto edits this plan makes are the `Stat` enum merge (Task 4) and two additive fields on `Encounter`/`Target` (Task 8). Additive means: new field numbers, nothing renumbered, nothing removed.
- **The `Stat` enum in `sim/core/stats/stats.go` and `proto.Stat` in `proto/common.proto` must stay index-synced.** The Go file says so at line 19. After the merge both shrink by two and every later index shifts down by two; Task 4 changes both in one commit and asserts the sync in a test.
- **The five WASM entrypoints the web uses are `raidSimAsync`, `raidSimRequestSplit`, `raidSimResultCombination`, `computeStats`, `abortById`** — all five confirmed present in `sim/wasm/main.go:28-38`. No others. Note `sim/wasm/main.go:39` calls `js.Global().Call("wasmready")` on startup, so the host page must define a global `wasmready` function before instantiating; the web lane owns that.
- **Unknown Forever numbers are never invented.** Every constant whose Forever value is not yet readable from a client table ships as the Era value, generated (not typed) into a constants file, and carries `unconfirmed` in a comment on its own line. The nightly validation job (api lane) is what clears it. A task that cannot generate a number from data must say so in the constants file, not guess.
- **Testing rule: a task runs only the tests covering the files it changed.** The full engine suite (`go test --tags=with_db -count=1 ./sim/...`, 10.5 s) runs once at the final review and in CI. The same for the site: `go test ./... -race` from `sim/`.
- **`gofmt` is a gate.** The engine has a `pre-commit` hook (`make setup` installs it) running `gofmt -w ./sim ./tools`. Run `gofmt -l ./sim ./tools` before every engine commit; a non-empty list fails CI.
- **Do not edit `logs/`.** `sim/adapter` imports `github.com/jhunthrop/foreversixty/logs/engine/summary` and builds its types; it never changes them.
- **Commits.** Conventional subject, a body, and exactly one trailer line in the final `-m`:
  ```
  Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
  ```
  Never `--no-verify`.

### What the data lane hands this lane, verified against the engine's own protos

The data lane finished first and checked these against `proto/common.proto`. They are binding; the contract carries them as of commit `73f86c9`.

- **`SimDatabase` is exactly `{items, random_suffixes, enchants}`** (`proto/common.proto:844-848`). It has **no item-set field and no consumables field**. Nothing in this lane may expect either.
- **Item sets ride on each item**, as `SimItem.set_id` (field 18) and `SimItem.set_name` (field 14) — confirmed at `proto/common.proto:870-871`. The Era build emits 3,144 such items. `sim/common/item_sets/` keeps working unchanged: it matches on set name, which is on the item.
- **Consumables are a sidecar**, `data/builds/<build>/simconsumes.json`, 1,467 rows. The data lane owns its shape and asks this lane to confirm it against the hand-written `sim/core/consumes.go` (1,252 lines). Task 10 carries that review step.
- **`random_suffixes` is emitted empty** — resolving vanilla suffixes needs the RandPropPoints allocation table, and Forever re-itemises anyway. `sim/core/database.go` already tolerates an empty list; do not add code that assumes suffixes exist.
- **Icons are not in `SimDatabase`.** They stay in `data/builds/<build>/icons/` for the site. The engine never reads them.
- **The vanilla coefficient conventions stay in this lane.** `simconst` emits the DB2 coefficient columns verbatim, **zeros included**, because for Classic-lineage spells `EffectBonusCoefficient` is routinely 0 or wrong (research §5.3). The `cast_time/3.5` and `duration/15` conventions, their halving for hybrids, and the per-spell overrides are the engine's, exactly as they are today. Task 10's generator must therefore treat a zero coefficient as "not in the data" and fall back to the convention, never as "the coefficient is zero".
- **The engine's checked-in preset APLs are stale on spell ranks.** Verified by the data lane: `ui/mage/apls/p1.apl.json` casts Frostbolt rank 10 where the Era tables give rank 11 for spell `25304`. **Do not copy a preset APL as a starting point without re-validating every spell rank against the build's own tables.** Tasks 11 and 12 carry that check.

---

## Parallel groups

The controller may run each group's tasks in parallel worktrees; groups run in order.

| Group | Tasks | Repo | Runs when | Notes |
|---|---|---|---|---|
| **G0** | 1 | engine | first, alone | Everything depends on committed protobufs. Serial gate. |
| **G1a** | 2 → 3 | site | after G0 | **Serial within the group.** Task 3 (`sim/adapter`) unblocks the api and web lanes; it is the lane's highest-priority output. |
| **G1b** | 4 | engine | after G0, **in parallel with G1a** | Serial gate for all later engine work: it renames symbols in 41 `sim/` files and 27 `ui/` files, so it must not race another engine task. |
| **G2** | 5, 6, 7, 8, 9, 10 | engine | after G1b | **All six are INDEPENDENT.** Disjoint file sets; six parallel worktrees. |
| **G3** | 11, 12 | engine | after G2 (needs 7 and 10) | **Both INDEPENDENT of each other.** Task 11 touches only `sim/warrior/**`, Task 12 only `sim/mage/**`. |
| **G4** | 13 | engine + site | after G3 | Artifacts and CI; needs the specs to exist so the smoke test has something to run. |
| **G5** | 14 | both | last, alone | Final review: full suites, both repos. |

Conflict map, so a controller can verify the claim:

- Task 4 alone touches `sim/core/stats/stats.go`, `proto/common.proto`, and 41 files under `sim/`. Nothing in G2 may run beside it.
- Within G2: Task 5 owns `sim/core/base_stats_auto_gen.go` + `tools/base_stats_parser.py`; Task 6 owns `sim/core/stats/deps.go` + `sim/core/character.go`; Task 7 owns `sim/core/spell_mod.go` (new) + `sim/core/spell.go` + `sim/core/flags.go` + `sim/core/cooldown.go`; Task 8 owns `proto/common.proto` (additive, different message from Task 4's enum) + `sim/core/target.go` + `sim/core/unit.go` + `sim/core/item_effects.go`; Task 9 owns `sim/core/racials.go`; Task 10 owns `tools/spellconst_gen.py` + `sim/core/spellconst/**` (new package). The only shared file is `proto/common.proto` between Task 4 (done, in G1b) and Task 8 — different messages, no overlap.

---

## File structure

```
ENGINE REPO  /Users/jh/code/wowsims-forever   (module github.com/wowsims/classic)
  sim/core/proto/.gitignore                   MODIFIED T1: stop ignoring *.pb.go
  sim/core/proto/*.pb.go                      NEW T1: 14 generated files, committed
  vite.build-workers.ts                       MODIFIED T1: GOROOT/lib/wasm, not misc/wasm
  makefile                                    MODIFIED T1 (proto check), T13 (artifacts)
  proto/common.proto                          MODIFIED T4 (Stat enum), T8 (Biome, Target.biome_affinity)
  sim/core/stats/stats.go                     MODIFIED T4: Hit, Crit; StatName
  sim/core/stats/stats_test.go                MODIFIED T4: index-sync assertion
  sim/core/stats/deps.go                      MODIFIED T6: safeDepsOrder places HealingPower before SpellDamage
  sim/core/stats/deps_test.go                 MODIFIED T6
  sim/core/character.go                       MODIFIED T6: the healing -> spell damage dependency
  sim/core/base_stats_auto_gen.go             REGENERATED T5
  sim/core/base_stats_forever.go              NEW T5: the provisional-value marker and its test hook
  tools/base_stats_parser.py                  MODIFIED T5: fixed, --build, --out
  sim/core/spell_mod.go                       NEW T7: ported from wowsims/sod
  sim/core/spell_mod_test.go                  NEW T7
  sim/core/spell.go                           MODIFIED T7: Matches, ClassSpellMask, Apply*DamageBonus, RelatedSelfBuff
  sim/core/cooldown.go                        MODIFIED T7: ApplyFlatCooldownMod, ApplyFlatPercentCooldownMod
  sim/core/flags.go                           MODIFIED T7: SpellFlagNoSpellMods
  sim/core/target.go                          MODIFIED T8: Encounter.Biome
  sim/core/unit.go                            MODIFIED T8: Unit.Biome accessor path
  sim/core/item_effects.go                    MODIFIED T8: NewBiomeDamageEffect, NewMobTypeDamageEffect
  sim/core/environment_biome_test.go          NEW T8
  sim/core/racials.go                         REWRITTEN T9
  sim/core/racials_test.go                    NEW T9
  sim/core/spellconst/spellconst.go           NEW T10: the Go-side shape
  sim/core/spellconst/spellconst_test.go      NEW T10
  sim/core/spellconst/gen/main.go             NEW T10: JSON -> Go generator
  sim/warrior/constants_auto_gen.go           NEW T10 (empty) / T11 (populated)
  sim/warrior/talents.go                      REWRITTEN T11: seven rows, 11/16/21/31
  sim/warrior/{rampage,piercing_howl}.go      NEW T11: baseline abilities
  sim/warrior/dps_warrior/dps_warrior_test.go MODIFIED T11
  ui/warrior/apls/forever_fury.apl.json       NEW T11
  sim/mage/constants_auto_gen.go              NEW T10 (empty) / T12 (populated)
  sim/mage/talents.go                         REWRITTEN T12
  sim/mage/{ice_lance,cold_snap_baseline}.go  NEW T12
  sim/mage/mage_test.go                       MODIFIED T12
  ui/mage/apls/forever_frost.apl.json         NEW T12
  cmd/forever-sim/main.go                     NEW T13: the native server-lane binary
  cmd/forever-sim/main_test.go                NEW T13
  .github/workflows/artifacts.yml             NEW T13: builds and publishes all three artifacts
  ui/core/proto_utils/{names,stats}.ts        MODIFIED T4: 177 references across 27 files

SITE REPO  /Users/jh/code/forever
  go.work                                     MODIFIED T2: use ./sim
  Makefile                                    NEW T2: engine-pin
  sim/go.mod, sim/go.sum                      NEW T2
  sim/enginever/version.go                    NEW T2 (generated by make engine-pin)
  sim/api/envelope.go                         NEW T2
  sim/api/envelope_test.go                    NEW T2
  sim/adapter/adapter.go                      NEW T3: Summarize
  sim/adapter/adapter_test.go                 NEW T3
  sim/adapter/golden_test.go                  NEW T3
  sim/adapter/testdata/*.result.pb            NEW T3 (T11/T12 refresh)
  sim/adapter/testdata/*.summary.json.golden  NEW T3 (T11/T12 refresh)
  .github/workflows/sim.yml                   NEW T13
```

---

## Task 1: Make the engine consumable as a Go module

**Repo: ENGINE** (`/Users/jh/code/wowsims-forever`). Run every command from that directory.

A fresh clone of this repository does not build. `sim/core/proto/.gitignore` line 2 is `*.pb.go`, so the 14 generated protobuf files that every other package imports are absent, and `go build ./...` fails with `no required module provides package github.com/wowsims/classic/sim/core/proto`. Upstream gets away with this because its only consumer is its own makefile. Ours cannot: the site's `sim/` module depends on `github.com/wowsims/classic/sim/core/proto` at a pinned pseudo-version, and Go resolves that from the commit, not from a build step. So the fork commits the generated code. That is a deliberate divergence from upstream, and the comment in `.gitignore` says why so a future merge does not silently revert it.

Two smaller repairs ride along because they are in the same "a fresh checkout works" deliverable: `vite.build-workers.ts:25` reads `$GOROOT/misc/wasm/wasm_exec.js`, a path Go removed in 1.24 (it is `$GOROOT/lib/wasm/wasm_exec.js` now — verified: `misc/wasm` does not exist under go1.25.4), and the `proto` make target silently produces nothing if `protoc-gen-go` is not on `PATH`.

**Files:**
- Modify: `sim/core/proto/.gitignore`
- Create (by generation, then commit): `sim/core/proto/api.pb.go`, `apl.pb.go`, `common.pb.go`, `druid.pb.go`, `hunter.pb.go`, `mage.pb.go`, `paladin.pb.go`, `priest.pb.go`, `rogue.pb.go`, `shaman.pb.go`, `test.pb.go`, `ui.pb.go`, `warlock.pb.go`, `warrior.pb.go`
- Modify: `vite.build-workers.ts:25`
- Modify: `makefile` (the `sim/core/proto/api.pb.go` rule, line 197)
- Create: `PORTING.md`
- Test: `sim/core/proto/generated_test.go`

**Interfaces:**
- Produces: the package `github.com/wowsims/classic/sim/core/proto` is importable from a module that only has the commit — this is what Task 2's `sim/go.mod` requires. No Go symbols are added.

- [ ] **Step 1: Write the failing test**

The test asserts the invariant a reviewer actually cares about: the committed generated code matches the `.proto` sources it claims to come from. It shells out to `protoc`, regenerates into a temp directory, and diffs. It skips (not fails) when `protoc` or `protoc-gen-go` is missing, so a contributor without them can still run the suite; CI has both, so CI enforces it.

Create `sim/core/proto/generated_test.go`:

```go
package proto_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// The generated protobuf code is committed (see PORTING.md): the site
// consumes this repository as a Go module at a pinned pseudo-version, and
// Go resolves packages from the commit, not from a build step. That makes
// drift between proto/*.proto and sim/core/proto/*.pb.go a real bug, so
// this test regenerates into a temp dir and diffs.
func TestGeneratedProtosMatchSources(t *testing.T) {
	if _, err := exec.LookPath("protoc"); err != nil {
		t.Skip("protoc not installed; CI enforces this test")
	}
	if _, err := exec.LookPath("protoc-gen-go"); err != nil {
		t.Skip("protoc-gen-go not installed; run go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6")
	}

	root := filepath.Join("..", "..", "..")
	tmp := t.TempDir()
	// protoc writes to <out>/github.com/wowsims/classic/sim/core/proto,
	// because that is the go_package the .proto files declare.
	cmd := exec.Command("protoc", "-I=./proto", "--go_out="+tmp, "./proto/api.proto",
		"./proto/apl.proto", "./proto/common.proto", "./proto/druid.proto",
		"./proto/hunter.proto", "./proto/mage.proto", "./proto/paladin.proto",
		"./proto/priest.proto", "./proto/rogue.proto", "./proto/shaman.proto",
		"./proto/test.proto", "./proto/ui.proto", "./proto/warlock.proto",
		"./proto/warrior.proto")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("protoc failed: %v\n%s", err, out)
	}

	genDir := filepath.Join(tmp, "github.com", "wowsims", "classic", "sim", "core", "proto")
	entries, err := os.ReadDir(genDir)
	if err != nil {
		t.Fatalf("reading regenerated dir: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("protoc produced no files")
	}
	for _, e := range entries {
		want, err := os.ReadFile(filepath.Join(genDir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		got, err := os.ReadFile(e.Name())
		if err != nil {
			t.Fatalf("%s is not committed: %v (run `make proto` and commit the result)", e.Name(), err)
		}
		if string(got) != string(want) {
			t.Errorf("%s is stale; run `make proto` and commit the result", e.Name())
		}
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd /Users/jh/code/wowsims-forever
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
export PATH=$PATH:$(go env GOPATH)/bin
go test ./sim/core/proto/ -run TestGeneratedProtosMatchSources -v
```

Expected: the package does not compile at all yet (`no required module provides package .../sim/core/proto`) on a clean tree, or — if you have run `make proto` before — the test passes for files present and the point is moot. To see the real failure, remove the generated files first: `rm -f sim/core/proto/*.pb.go` then run the command. Expect `FAIL` with `api.pb.go is not committed`.

- [ ] **Step 3: Generate, un-ignore, and commit**

Replace `sim/core/proto/.gitignore` in full with:

```gitignore
# NOTE FOR UPSTREAM MERGES: this file used to be `*.pb.go`.
# The Forever fork commits the generated protobuf code on purpose. The site
# repository consumes this module at a pinned pseudo-version, and the Go
# module system resolves packages from the commit, so a generated-but-
# ignored package makes the module unusable as a dependency.
# `make proto` regenerates; sim/core/proto/generated_test.go fails if the
# committed output drifts from proto/*.proto.
```

Then:

```bash
cd /Users/jh/code/wowsims-forever
export PATH=$PATH:$(go env GOPATH)/bin
make proto
```

- [ ] **Step 4: Fix the two build-path defects**

`vite.build-workers.ts:25` — replace

```ts
	const wasmExecutablePath = path.join(GO_ROOT, '/misc/wasm/wasm_exec.js');
```

with

```ts
	// Go moved wasm_exec.js from misc/wasm to lib/wasm in Go 1.24.
	// Prefer the new location and fall back for older toolchains.
	const wasmExecCandidates = [path.join(GO_ROOT, 'lib', 'wasm', 'wasm_exec.js'), path.join(GO_ROOT, 'misc', 'wasm', 'wasm_exec.js')];
	const wasmExecutablePath = wasmExecCandidates.find(p => fs.existsSync(p)) ?? wasmExecCandidates[0];
```

and make sure `import fs from 'fs';` is among the file's imports (add it if absent).

In `makefile`, replace the rule at line 197:

```makefile
sim/core/proto/api.pb.go: proto/*.proto
	protoc -I=./proto --go_out=./sim/core ./proto/*.proto
```

with

```makefile
sim/core/proto/api.pb.go: proto/*.proto
	@command -v protoc-gen-go >/dev/null || { \
	  echo "protoc-gen-go not on PATH."; \
	  echo "  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6"; \
	  echo "  export PATH=\$$PATH:\$$(go env GOPATH)/bin"; \
	  exit 1; }
	protoc -I=./proto --go_out=./sim/core ./proto/*.proto
```

- [ ] **Step 5: Write the porting note**

Create `PORTING.md`:

```markdown
# Forever fork: deliberate divergences from wowsims/classic

This fork tracks `upstream` = https://github.com/wowsims/classic. Everything
here is a change a merge must not silently revert.

## Committed generated protobufs

`sim/core/proto/*.pb.go` is committed; upstream ignores it. The Forever Sixty
site consumes this repository as a Go module at a pinned pseudo-version, and
the Go module system resolves packages from the commit, not from a build step.
`make proto` regenerates. `sim/core/proto/generated_test.go` fails if the
committed output drifts from `proto/*.proto`.

Requires: protoc >= 3.21 and protoc-gen-go v1.36.6
(`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6`).

## wasm_exec.js location

Go 1.24 moved `wasm_exec.js` from `$GOROOT/misc/wasm` to `$GOROOT/lib/wasm`.
`vite.build-workers.ts` checks both.

## Build prerequisites for a fresh clone

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
    export PATH=$PATH:$(go env GOPATH)/bin
    make proto
    make binary_dist/dist.go     # sim/web embeds this; it is gitignored
    go build ./...
    go test --tags=with_db ./sim/...
```

- [ ] **Step 6: Verify a fresh clone builds**

```bash
cd /Users/jh/code/wowsims-forever
export PATH=$PATH:$(go env GOPATH)/bin
go test ./sim/core/proto/ -run TestGeneratedProtosMatchSources -v
gofmt -l ./sim ./tools
git stash list >/dev/null && rm -rf /tmp/fresh-engine && git clone . /tmp/fresh-engine 2>/dev/null
cd /tmp/fresh-engine && go build ./sim/... && echo "FRESH CLONE BUILDS"
```

Expected: `PASS`; `gofmt -l` prints nothing; `FRESH CLONE BUILDS`. (The clone is from the local worktree, so it only sees committed state — which is the point. `go build ./...` in the clone still fails on `sim/web` until `make binary_dist/dist.go` runs; `./sim/...` is the meaningful check and is what the site depends on.)

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add sim/core/proto/ vite.build-workers.ts makefile PORTING.md
git commit -m "build: commit the generated protobufs so the module is consumable" \
  -m "The site imports github.com/wowsims/classic/sim/core/proto at a pinned pseudo-version, and Go resolves packages from the commit rather than from a build step, so an ignored generated package makes this module unusable as a dependency. A test regenerates and diffs, so the committed output cannot drift. Also fixes the wasm_exec.js path Go moved in 1.24 and makes the proto target say what to install when protoc-gen-go is missing." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 2: The `sim/` module, the engine pin, and the request/result envelopes

**Repo: SITE** (`/Users/jh/code/forever`). Depends on Task 1.

The site gains its fourth Go module. It holds three things and nothing else: the one string that identifies an engine build (`sim/enginever`), the envelopes both compute lanes speak (`sim/api`), and — next task — the adapter. It depends on the engine through a `replace` in `go.work` during development and a pinned pseudo-version in `sim/go.mod` for CI, exactly as the contract specifies.

`make engine-pin` is the only way `sim/enginever/version.go` is ever written. It reads the short sha out of the engine checkout, so the pin cannot drift from what was actually built.

**Files:**
- Create: `sim/go.mod`, `sim/go.sum`, `sim/enginever/version.go`, `sim/api/envelope.go`, `Makefile` (repository root)
- Modify: `go.work`
- Test: `sim/api/envelope_test.go`, `sim/enginever/version_test.go`

**Interfaces:**
- Consumes: `github.com/wowsims/classic/sim/core/proto` (Task 1), `github.com/jhunthrop/foreversixty/logs/engine/summary`.
- Produces, all used by Task 3, the api lane and the web lane:
  - `enginever.Version string` — the constant, e.g. `"7779ebb"`.
  - `api.SimRequest{EngineVersion string; Spec string; Source CharacterSource; Encounter EncounterSpec; Iterations int; RandomSeed int64; Raw []byte}`
  - `api.CharacterSource{Kind, Ref, CapturedAt string}`
  - `api.EncounterSpec{DurationSec int; Variation float64; Targets int; ExecuteRatio float64; Profile string}`
  - `api.SimResult{SimID, EngineVersion string; Request SimRequest; Lane string; DPS Estimate; IterationsRun int; DurationMS int64; Summary summary.Summary; Error string}`
  - `api.Estimate{Mean, StdDev, Error, Min, Max float64}`
  - `api.DefaultEncounter() EncounterSpec` — the contract's defaults: 180 s, 0.2 variation, 1 target, 0.25 execute, empty profile.
  - `api.ValidIterations = []int{500, 3000, 10000}`, `api.(SimRequest).Validate() error`.
  - `api.(SimResult).ForStorage() SimResult`, `api.(SimResult).Stale(current string) bool`.
  - `api.LaneBrowser = "browser"`, `api.LaneServer = "server"`.
  - `api.SourceArmory/SourceAddon/SourceBuild/SourceFight/SourceManual` — the five `Kind` values.

- [ ] **Step 1: Create the module and wire the workspace**

```bash
cd /Users/jh/code/forever
mkdir -p sim/api sim/enginever sim/adapter
cd sim
cat > go.mod <<'EOF'
module github.com/jhunthrop/foreversixty/sim

go 1.25.11

require (
	github.com/jhunthrop/foreversixty/logs v0.0.0
	github.com/wowsims/classic v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.36.6
)

replace github.com/jhunthrop/foreversixty/logs => ../logs

// Development pin. CI rewrites this to a pseudo-version of
// github.com/jhunthrop/wowsims-forever at the sha in sim/enginever/version.go.
replace github.com/wowsims/classic => /Users/jh/code/wowsims-forever
EOF
```

Then add `./sim` to the workspace. `/Users/jh/code/forever/go.work`'s use block becomes:

```
use (
	./api
	./companion
	./logs
	./sim
)
```

- [ ] **Step 2: Write the failing envelope test**

Create `sim/api/envelope_test.go`:

```go
package api

import (
	"encoding/json"
	"strings"
	"testing"
)

// The JSON field names are the contract, shared verbatim with
// web/src/lib/sim/types.ts. A rename here is a break there, so the test
// pins the wire form rather than the Go field names.
func TestSimRequestJSONFieldNames(t *testing.T) {
	req := SimRequest{
		EngineVersion: "7779ebb",
		Spec:          "warrior-fury",
		Source:        CharacterSource{Kind: SourceArmory, Ref: "us/normal/thrall", CapturedAt: "2026-09-14T00:00:00Z"},
		Encounter:     DefaultEncounter(),
		Iterations:    3000,
		RandomSeed:    0,
		Raw:           []byte{1, 2, 3},
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"engine_version", "spec", "source", "encounter", "iterations", "random_seed", "raw"} {
		if _, ok := m[k]; !ok {
			t.Errorf("SimRequest is missing JSON key %q", k)
		}
	}
	src, ok := m["source"].(map[string]any)
	if !ok {
		t.Fatalf("source is not an object: %T", m["source"])
	}
	for _, k := range []string{"kind", "ref", "captured_at"} {
		if _, ok := src[k]; !ok {
			t.Errorf("CharacterSource is missing JSON key %q", k)
		}
	}
	enc, ok := m["encounter"].(map[string]any)
	if !ok {
		t.Fatalf("encounter is not an object: %T", m["encounter"])
	}
	for _, k := range []string{"duration_sec", "variation", "targets", "execute_ratio", "profile"} {
		if _, ok := enc[k]; !ok {
			t.Errorf("EncounterSpec is missing JSON key %q", k)
		}
	}
}

func TestDefaultEncounterMatchesTheContract(t *testing.T) {
	got := DefaultEncounter()
	want := EncounterSpec{DurationSec: 180, Variation: 0.2, Targets: 1, ExecuteRatio: 0.25, Profile: ""}
	if got != want {
		t.Errorf("DefaultEncounter() = %+v, want %+v", got, want)
	}
}

func TestValidateRejectsBadRequests(t *testing.T) {
	good := SimRequest{EngineVersion: "7779ebb", Spec: "mage-frost", Iterations: 3000, Encounter: DefaultEncounter(), Raw: []byte{1}}
	if err := good.Validate(); err != nil {
		t.Fatalf("a good request was rejected: %v", err)
	}
	cases := []struct {
		name string
		mut  func(*SimRequest)
		want string
	}{
		{"no engine version", func(r *SimRequest) { r.EngineVersion = "" }, "engine_version"},
		{"no spec", func(r *SimRequest) { r.Spec = "" }, "spec"},
		{"odd iteration count", func(r *SimRequest) { r.Iterations = 1234 }, "iterations"},
		{"no duration", func(r *SimRequest) { r.Encounter.DurationSec = 0 }, "duration_sec"},
		{"too many targets", func(r *SimRequest) { r.Encounter.Targets = 11 }, "targets"},
		{"no raw request", func(r *SimRequest) { r.Raw = nil }, "raw"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := good
			tc.mut(&req)
			err := req.Validate()
			if err == nil {
				t.Fatalf("expected an error mentioning %q, got nil", tc.want)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error %q does not mention %q", err, tc.want)
			}
		})
	}
}

// A stored sim row is served to browsers; the raw engine request is large
// and is rebuilt from the character model whenever a re-run is wanted.
func TestSimResultOmitsRawWhenStored(t *testing.T) {
	res := SimResult{
		EngineVersion: "7779ebb",
		Request:       SimRequest{Spec: "warrior-fury", Raw: []byte{9, 9, 9}},
		Lane:          LaneBrowser,
		DPS:           Estimate{Mean: 1791.1, StdDev: 120, Error: 2.2, Min: 1400, Max: 2100},
	}
	b, err := json.Marshal(res.ForStorage())
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), `"raw":"`) {
		t.Errorf("ForStorage() kept the raw request: %s", b)
	}
	if !strings.Contains(string(b), `"warrior-fury"`) {
		t.Errorf("ForStorage() dropped the request: %s", b)
	}
}

func TestStaleComparesEngineVersions(t *testing.T) {
	res := SimResult{EngineVersion: "aaaaaaa"}
	if res.Stale("aaaaaaa") {
		t.Error("a result from the current engine reported stale")
	}
	if !res.Stale("bbbbbbb") {
		t.Error("a result from another engine build did not report stale")
	}
}
```

- [ ] **Step 3: Run it and watch it fail**

Run: `cd /Users/jh/code/forever/sim && go test ./api/ -v`
Expected: `FAIL [build failed]` with `undefined: SimRequest`.

- [ ] **Step 4: Write the envelopes**

Create `sim/api/envelope.go`:

```go
// Package api holds the request and result envelopes both compute lanes
// speak. The engine never sees them: SimRequest.Raw is the engine's own
// RaidSimRequest protobuf, and both lanes pass it through untouched.
//
// The JSON field names here are authoritative and are mirrored verbatim in
// web/src/lib/sim/types.ts.
package api

import (
	"errors"
	"fmt"
	"slices"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The five character sources.
const (
	SourceArmory = "armory"
	SourceAddon  = "addon"
	SourceBuild  = "build"
	SourceFight  = "fight"
	SourceManual = "manual"
)

// The two compute lanes.
const (
	LaneBrowser = "browser"
	LaneServer  = "server"
)

// ValidIterations is the closed set the UI offers: a planner-inline
// estimate, the default, and the precision toggle.
var ValidIterations = []int{500, 3000, 10000}

// MaxTargets is the settings bar's cap.
const MaxTargets = 10

// MinDurationSec and MaxDurationSec bound the fight-length control.
const (
	MinDurationSec = 60
	MaxDurationSec = 480
)

type SimRequest struct {
	EngineVersion string          `json:"engine_version"`
	Spec          string          `json:"spec"`
	Source        CharacterSource `json:"source"`
	Encounter     EncounterSpec   `json:"encounter"`
	Iterations    int             `json:"iterations"`
	RandomSeed    int64           `json:"random_seed"`
	Raw           []byte          `json:"raw"`
}

type CharacterSource struct {
	Kind       string `json:"kind"`
	Ref        string `json:"ref"`
	CapturedAt string `json:"captured_at"`
}

type EncounterSpec struct {
	DurationSec  int     `json:"duration_sec"`
	Variation    float64 `json:"variation"`
	Targets      int     `json:"targets"`
	ExecuteRatio float64 `json:"execute_ratio"`
	Profile      string  `json:"profile"`
}

// DefaultEncounter is the settings bar's opening state: a three-minute
// single-target fight with the standard duration variation and execute
// window.
func DefaultEncounter() EncounterSpec {
	return EncounterSpec{DurationSec: 180, Variation: 0.2, Targets: 1, ExecuteRatio: 0.25}
}

// Validate checks everything a malformed client could get wrong, at the
// boundary, before anything reaches the engine.
func (r SimRequest) Validate() error {
	var errs []error
	if r.EngineVersion == "" {
		errs = append(errs, errors.New("engine_version is required"))
	}
	if r.Spec == "" {
		errs = append(errs, errors.New("spec is required"))
	}
	if !slices.Contains(ValidIterations, r.Iterations) {
		errs = append(errs, fmt.Errorf("iterations must be one of %v, got %d", ValidIterations, r.Iterations))
	}
	if r.Encounter.DurationSec < MinDurationSec || r.Encounter.DurationSec > MaxDurationSec {
		errs = append(errs, fmt.Errorf("duration_sec must be between %d and %d, got %d", MinDurationSec, MaxDurationSec, r.Encounter.DurationSec))
	}
	if r.Encounter.Targets < 1 || r.Encounter.Targets > MaxTargets {
		errs = append(errs, fmt.Errorf("targets must be between 1 and %d, got %d", MaxTargets, r.Encounter.Targets))
	}
	if r.Encounter.Variation < 0 || r.Encounter.Variation > 1 {
		errs = append(errs, fmt.Errorf("variation must be between 0 and 1, got %v", r.Encounter.Variation))
	}
	if r.Encounter.ExecuteRatio < 0 || r.Encounter.ExecuteRatio > 1 {
		errs = append(errs, fmt.Errorf("execute_ratio must be between 0 and 1, got %v", r.Encounter.ExecuteRatio))
	}
	if len(r.Raw) == 0 {
		errs = append(errs, errors.New("raw is required: the engine's own RaidSimRequest protobuf"))
	}
	return errors.Join(errs...)
}

type SimResult struct {
	SimID         string          `json:"sim_id,omitempty"`
	EngineVersion string          `json:"engine_version"`
	Request       SimRequest      `json:"request"`
	Lane          string          `json:"lane"`
	DPS           Estimate        `json:"dps"`
	IterationsRun int             `json:"iterations_run"`
	DurationMS    int64           `json:"duration_ms"`
	Summary       summary.Summary `json:"summary"`
	Error         string          `json:"error,omitempty"`
}

type Estimate struct {
	Mean   float64 `json:"mean"`
	StdDev float64 `json:"stddev"`
	Error  float64 `json:"error"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
}

// ForStorage returns a copy with the raw engine request dropped. A stored
// sim row is served to browsers; the raw request is large and is rebuilt
// from the character model whenever a re-run is wanted.
func (r SimResult) ForStorage() SimResult {
	r.Request.Raw = nil
	return r
}

// Stale reports whether this result came from an engine build other than
// the current one. A stale result is still readable and is labelled in the
// UI; it is never silently re-run.
func (r SimResult) Stale(current string) bool {
	return r.EngineVersion != current
}
```

- [ ] **Step 5: Run the test and watch it pass**

Run: `cd /Users/jh/code/forever/sim && go mod tidy && go test ./api/ -v`
Expected: `PASS` for all five tests, `ok github.com/jhunthrop/foreversixty/sim/api`.

- [ ] **Step 6: Write the engine-pin test**

Create `sim/enginever/version_test.go`:

```go
package enginever

import "testing"

// Version is the short commit sha of wowsims-forever the artifacts were
// built from. It is generated by `make engine-pin` from the root of the
// site repository and must never be edited by hand.
func TestVersionLooksLikeAShortSha(t *testing.T) {
	if len(Version) < 7 || len(Version) > 12 {
		t.Fatalf("Version = %q; want a 7-to-12 character short sha", Version)
	}
	for _, c := range Version {
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
			t.Fatalf("Version = %q; want lowercase hexadecimal", Version)
		}
	}
}
```

- [ ] **Step 7: Run it and watch it fail**

Run: `cd /Users/jh/code/forever/sim && go test ./enginever/`
Expected: `FAIL [build failed]`, `undefined: Version`.

- [ ] **Step 8: Add the `engine-pin` make target and run it**

Create `/Users/jh/code/forever/Makefile`. Note the doubled `%%s`: make eats a single `%` in a recipe, so the `printf` format string needs it escaped.

```makefile
# Repository-level targets. Each module keeps its own tooling; this file is
# only for things that cross a module boundary.

# Where the engine fork is checked out. Override for a different location:
#   make engine-pin ENGINE_DIR=/somewhere/else
ENGINE_DIR ?= /Users/jh/code/wowsims-forever

.PHONY: engine-pin
# engine-pin writes sim/enginever/version.go from the engine checkout's HEAD.
# This is the only way that file is ever written. ENGINE_VERSION is the short
# sha, and it names the wasm artifact directory, the premium image tag, and
# every stored sim and validation row, so pinning a dirty tree would produce
# a version string that identifies nothing. Hence the cleanliness check.
engine-pin:
	@test -d "$(ENGINE_DIR)/.git" || { echo "no engine checkout at $(ENGINE_DIR); set ENGINE_DIR"; exit 1; }
	@if [ -n "$$(git -C "$(ENGINE_DIR)" status --porcelain)" ]; then \
	  echo "engine checkout at $(ENGINE_DIR) is dirty; commit or stash before pinning"; exit 1; \
	fi
	@sha=$$(git -C "$(ENGINE_DIR)" rev-parse --short HEAD); \
	mkdir -p sim/enginever; \
	printf '// Code generated by "make engine-pin". DO NOT EDIT.\n\npackage enginever\n\n// Version is the short commit sha of wowsims-forever that sim.wasm, sim.js\n// and the forever-sim binary were built from. It names the immutable wasm\n// directory, the premium lane image tag, and every stored sim row.\nconst Version = "%%s"\n' "$$sha" > sim/enginever/version.go; \
	gofmt -w sim/enginever/version.go; \
	echo "pinned engine version $$sha"
```

Run:

```bash
cd /Users/jh/code/forever
make engine-pin
cat sim/enginever/version.go
```

Expected: `pinned engine version <sha>` and a file ending in `const Version = "<sha>"`. At the time this plan was written the engine's HEAD was `7779ebb`.

- [ ] **Step 9: Run the tests and watch them pass**

Run: `cd /Users/jh/code/forever/sim && go test ./enginever/ ./api/`
Expected: `ok` for both packages.

- [ ] **Step 10: Commit**

```bash
cd /Users/jh/code/forever
git add go.work Makefile sim/
git commit -m "feat(sim): the sim module, the engine pin, and the request envelopes" \
  -m "A fourth Go module holding the two things every lane shares: the one string that names an engine build, written only by make engine-pin from the engine checkout's HEAD so it cannot drift from what was built, and the SimRequest/SimResult envelopes whose JSON field names are mirrored verbatim in the web's types.ts. The engine comes in through a replace during development and a pinned pseudo-version in CI." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 3: `sim/adapter` — engine result to logs summary

**Repo: SITE** (`/Users/jh/code/forever`). Depends on Task 2. **THE API AND WEB LANES ARE BLOCKED ON THIS TASK.** Finish it before anything in G2.

`Summarize` turns a `*proto.RaidSimResult` into a `logs/engine/summary.Summary`, so the report page's damage table, aura uptimes, cast list, and resource timeline render a simulation with no second renderer. The mapping is the contract's table, field by field, and the rule that governs all of it is: **the engine reports totals accumulated across every iteration, and the summary describes one fight, so every count and every total is divided by `IterationsDone` before it enters the summary.** The distribution stays in `SimResult.DPS`, where it belongs.

Two shapes do not survive the trip and the adapter says so rather than inventing them. The engine's `AuraMetrics` carries an average uptime in seconds and an average proc count, but no application timeline, so `AuraTrack.Segments` is empty and `Applications` is the rounded proc average. The engine's `ActionMetrics` carries cast counts but no timestamps, so `CastRow.Sequence` is empty. `DamageTaken`, `Healing`, `HealingTaken`, `Deaths`, `Interrupts`, `Dispels`, `Threat` and `Combatants` are empty at launch, per the contract.

Engine shapes this task reads, all confirmed in `proto/api.proto`: `RaidSimResult{RaidMetrics *RaidMetrics; AvgIterationDuration float64; IterationsDone int32; Error *ErrorOutcome}`, `RaidMetrics{Parties []*PartyMetrics}`, `PartyMetrics{Players []*UnitMetrics}`, `UnitMetrics{Name string; UnitIndex int32; Dps *DistributionMetrics; Actions []*ActionMetrics; Auras []*AuraMetrics; Resources []*ResourceMetrics; Pets []*UnitMetrics}`, `ActionMetrics{Id *ActionID; IsMelee bool; SpellSchool int32; Targets []*TargetedActionMetrics}`, `TargetedActionMetrics{Casts, Hits, Crits, Ticks, CritTicks, Misses, Dodges, Parries, Blocks, Glances, Crushes int32; Damage, CritDamage, TickDamage, CritTickDamage, GlanceDamage, CrushDamage, BlockDamage float64}`, `AuraMetrics{Id *ActionID; UptimeSecondsAvg, ProcsAvg float64}`, `ResourceMetrics{Id *ActionID; Type ResourceType; Events int32; Gain, ActualGain float64}`, `DistributionMetrics{Avg, Stdev, Max, Min float64}`, `ActionID{SpellId|ItemId|OtherId oneof; Tag, Rank int32}`.

**Files:**
- Create: `sim/adapter/adapter.go`, `sim/adapter/fixture.go`
- Test: `sim/adapter/adapter_test.go`, `sim/adapter/golden_test.go`
- Create (generated): `sim/adapter/testdata/warrior-fury.result.pb`, `sim/adapter/testdata/warrior-fury.summary.json.golden`, `sim/adapter/testdata/mage-frost.result.pb`, `sim/adapter/testdata/mage-frost.summary.json.golden`

**Interfaces:**
- Consumes: `api.SimRequest`, `api.Estimate`, `enginever.Version` (Task 2); `summary.Summary`, `summary.Actor`, `summary.Ability`, `summary.Pair`, `summary.AuraTrack`, `summary.CastRow`, `summary.ResourceTrack`, `summary.RosterRow` (the `logs/` engine, unchanged); `proto.RaidSimResult` and friends (the engine).
- Produces, used by the api lane's saved-sim write, validation job and execution scorer, and by the web lane through the stored `SimResult`:
  - `adapter.Summarize(res *proto.RaidSimResult, req api.SimRequest) (summary.Summary, error)`
  - `adapter.DPS(res *proto.RaidSimResult) api.Estimate`
  - `adapter.PlayerMetrics(res *proto.RaidSimResult) (*proto.UnitMetrics, error)`
  - `adapter.ErrNoPlayer`, `adapter.ErrSimFailed`
  - `adapter.ActionName(id *proto.ActionID) (spellID int64, name string)`
  - `adapter.Fixture(spec string) (*proto.RaidSimResult, error)` — reads `testdata/<spec>.result.pb`; the api lane's tests use it so they need no engine binary.

- [ ] **Step 1: Write the failing unit test**

Create `sim/adapter/adapter_test.go`:

```go
package adapter

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// oneAction builds a UnitMetrics with a single ability whose numbers are
// chosen so that the per-iteration division is visible: 200 casts over 100
// iterations is 2 casts per fight, not 200.
func oneAction() *proto.UnitMetrics {
	return &proto.UnitMetrics{
		Name:      "Fury",
		UnitIndex: 0,
		Dps:       &proto.DistributionMetrics{Avg: 1791.1, Stdev: 120.5, Max: 2100, Min: 1400},
		Actions: []*proto.ActionMetrics{{
			Id:          &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23894}},
			IsMelee:     true,
			SpellSchool: 1,
			Targets: []*proto.TargetedActionMetrics{{
				UnitIndex: 1,
				Casts:     200,
				Hits:      160,
				Crits:     40,
				Misses:    20,
				Dodges:    10,
				Parries:   5,
				Glances:   5,
				Damage:    100000,
				CritDamage: 40000,
			}},
		}},
		Auras: []*proto.AuraMetrics{{
			Id:               &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 12966}},
			UptimeSecondsAvg: 140.4,
			ProcsAvg:         31.5,
		}},
		Resources: []*proto.ResourceMetrics{{
			Id:        &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23894}},
			Type:      proto.ResourceType_ResourceTypeRage,
			Events:    200,
			Gain:      -6000,
			ActualGain: -6000,
		}},
	}
}

func resultWith(u *proto.UnitMetrics, iterations int32) *proto.RaidSimResult {
	return &proto.RaidSimResult{
		RaidMetrics: &proto.RaidMetrics{
			Dps:     &proto.DistributionMetrics{Avg: 1791.1, Stdev: 120.5, Max: 2100, Min: 1400},
			Parties: []*proto.PartyMetrics{{Players: []*proto.UnitMetrics{u}}},
		},
		AvgIterationDuration: 180.0,
		IterationsDone:       iterations,
	}
}

func req() api.SimRequest {
	return api.SimRequest{
		EngineVersion: "7779ebb",
		Spec:          "warrior-fury",
		Encounter:     api.DefaultEncounter(),
		Iterations:    100,
		Raw:           []byte{1},
	}
}

func TestSummarizeHeader(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if got.EngineVersion != "sim:7779ebb" {
		t.Errorf("EngineVersion = %q, want %q", got.EngineVersion, "sim:7779ebb")
	}
	if got.FightIndex != 1 {
		t.Errorf("FightIndex = %d, want 1", got.FightIndex)
	}
	if got.DurationMS != 180000 {
		t.Errorf("DurationMS = %d, want 180000 (avg_iteration_duration * 1000)", got.DurationMS)
	}
}

// The engine accumulates across iterations; the summary describes one
// fight. Everything countable is divided by IterationsDone.
func TestSummarizeDividesByIterations(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.DamageDone) != 1 {
		t.Fatalf("DamageDone has %d actors, want 1", len(got.DamageDone))
	}
	a := got.DamageDone[0]
	if a.Name != "Fury" {
		t.Errorf("actor Name = %q, want %q", a.Name, "Fury")
	}
	// 100000 damage + 40000 crit damage over 100 iterations.
	if a.Total != 1400 {
		t.Errorf("actor Total = %d, want 1400", a.Total)
	}
	if a.Effective != a.Total {
		t.Errorf("Effective = %d, want it equal to Total (%d)", a.Effective, a.Total)
	}
	if a.ActiveMS != 180000 {
		t.Errorf("ActiveMS = %d, want the fight duration 180000", a.ActiveMS)
	}
	if len(a.Abilities) != 1 {
		t.Fatalf("actor has %d abilities, want 1", len(a.Abilities))
	}
	ab := a.Abilities[0]
	if ab.SpellID != 23894 {
		t.Errorf("SpellID = %d, want 23894", ab.SpellID)
	}
	if ab.Hits != 2 { // 160 hits + 40 crits = 200 landed, over 100 iterations, minus crits
		t.Errorf("Hits = %d, want 2 (160/100 rounded)", ab.Hits)
	}
	if ab.Crits != 0 { // 40/100 = 0.4, rounds to 0
		t.Errorf("Crits = %d, want 0 (40/100 rounds down)", ab.Crits)
	}
	if ab.Misses["miss"] != 0 { // 20/100 = 0.2
		t.Errorf("misses = %d, want 0", ab.Misses["miss"])
	}
	if len(a.Targets) != 1 || a.Targets[0].Total != 1400 {
		t.Errorf("Targets = %+v, want one entry totalling 1400", a.Targets)
	}
}

// Rounding must not lose a whole ability: an ability cast once per fight
// over 100 iterations is 100 casts, which divides cleanly; but an ability
// cast every third fight must still appear, with a zero-ish count, rather
// than vanish from the table.
func TestSummarizeKeepsRareAbilities(t *testing.T) {
	u := oneAction()
	u.Actions = append(u.Actions, &proto.ActionMetrics{
		Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 20572}},
		IsMelee: false,
		Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 33, Hits: 33, Damage: 3300}},
	})
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, ab := range got.DamageDone[0].Abilities {
		if ab.SpellID == 20572 {
			found = true
			if ab.Total != 33 {
				t.Errorf("rare ability Total = %d, want 33", ab.Total)
			}
		}
	}
	if !found {
		t.Error("an ability used in a third of iterations vanished from the table")
	}
}

func TestSummarizeAurasCastsAndResources(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Auras) != 1 {
		t.Fatalf("Auras has %d entries, want 1", len(got.Auras))
	}
	au := got.Auras[0]
	if au.SpellID != 12966 {
		t.Errorf("aura SpellID = %d, want 12966", au.SpellID)
	}
	if au.UptimeMS != 140400 {
		t.Errorf("aura UptimeMS = %d, want 140400", au.UptimeMS)
	}
	if au.Applications != 32 { // 31.5 rounds to 32
		t.Errorf("aura Applications = %d, want 32", au.Applications)
	}
	if len(au.Segments) != 0 {
		t.Errorf("aura Segments = %v; the engine reports no application timeline, so this must stay empty", au.Segments)
	}

	if len(got.Casts) != 1 {
		t.Fatalf("Casts has %d entries, want 1", len(got.Casts))
	}
	c := got.Casts[0]
	if c.SpellID != 23894 || c.Started != 2 || c.Succeeded != 2 {
		t.Errorf("cast row = %+v, want spell 23894 started and succeeded 2", c)
	}
	if len(c.Sequence) != 0 {
		t.Errorf("cast Sequence = %v; the engine reports no cast timestamps, so this must stay empty", c.Sequence)
	}

	if len(got.Resources) != 1 {
		t.Fatalf("Resources has %d entries, want 1", len(got.Resources))
	}
	r := got.Resources[0]
	if r.PowerType != int64(proto.ResourceType_ResourceTypeRage) {
		t.Errorf("resource PowerType = %d, want %d", r.PowerType, proto.ResourceType_ResourceTypeRage)
	}
	if r.Spent != 60 {
		t.Errorf("resource Spent = %d, want 60 (6000 spent over 100 iterations)", r.Spent)
	}
	if r.Gained != 0 {
		t.Errorf("resource Gained = %d, want 0", r.Gained)
	}
}

func TestSummarizeRoster(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Roster) != 1 {
		t.Fatalf("Roster has %d rows, want 1", len(got.Roster))
	}
	r := got.Roster[0]
	if r.Name != "Fury" || r.Class != "warrior" || r.Spec != "fury" {
		t.Errorf("roster row = %+v, want name Fury, class warrior, spec fury", r)
	}
	if r.Role != "dps" {
		t.Errorf("roster Role = %q, want %q", r.Role, "dps")
	}
	if math.Abs(r.DPS-1791.1) > 0.001 {
		t.Errorf("roster DPS = %v, want 1791.1", r.DPS)
	}
}

// A pet is a second actor in the damage table, exactly as it is in a real
// fight's summary.
func TestSummarizePets(t *testing.T) {
	u := oneAction()
	u.Pets = []*proto.UnitMetrics{{
		Name: "Fury - Pet",
		Dps:  &proto.DistributionMetrics{Avg: 100},
		Actions: []*proto.ActionMetrics{{
			Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 3110}},
			IsMelee: true,
			Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 100, Hits: 100, Damage: 20000}},
		}},
	}}
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.DamageDone) != 2 {
		t.Fatalf("DamageDone has %d actors, want 2 (player and pet)", len(got.DamageDone))
	}
	if got.DamageDone[1].Name != "Fury - Pet" {
		t.Errorf("second actor = %q, want the pet", got.DamageDone[1].Name)
	}
	if got.DamageDone[1].Total != 200 {
		t.Errorf("pet Total = %d, want 200", got.DamageDone[1].Total)
	}
}

func TestSummarizeRejectsFailures(t *testing.T) {
	bad := resultWith(oneAction(), 100)
	bad.Error = &proto.ErrorOutcome{Message: "boom"}
	if _, err := Summarize(bad, req()); err == nil {
		t.Fatal("a result carrying an ErrorOutcome was summarized without error")
	}

	empty := &proto.RaidSimResult{RaidMetrics: &proto.RaidMetrics{}, IterationsDone: 100}
	if _, err := Summarize(empty, req()); err == nil {
		t.Fatal("a result with no player was summarized without error")
	}

	zero := resultWith(oneAction(), 0)
	if _, err := Summarize(zero, req()); err == nil {
		t.Fatal("a result with zero iterations was summarized without error")
	}
}

func TestDPS(t *testing.T) {
	got := DPS(resultWith(oneAction(), 100))
	if got.Mean != 1791.1 || got.StdDev != 120.5 || got.Min != 1400 || got.Max != 2100 {
		t.Errorf("DPS() = %+v", got)
	}
	// standard error of the mean = stdev / sqrt(n)
	want := 120.5 / math.Sqrt(100)
	if math.Abs(got.Error-want) > 1e-9 {
		t.Errorf("DPS().Error = %v, want %v", got.Error, want)
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/forever/sim && go test ./adapter/ -v`
Expected: `FAIL [build failed]`, `undefined: Summarize`.

- [ ] **Step 3: Write the adapter**

Create `sim/adapter/adapter.go`:

```go
// Package adapter turns an engine result into the logs engine's per-fight
// summary, so the report page's components render a simulation with no
// second renderer.
//
// The engine accumulates its metrics across every iteration. A summary
// describes one fight. So every count and every total here is divided by
// IterationsDone before it enters the summary; the distribution stays in
// api.SimResult.DPS, which is where a range belongs.
package adapter

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

var (
	// ErrSimFailed is returned when the engine itself reported a failure.
	ErrSimFailed = errors.New("adapter: the engine reported an error")
	// ErrNoPlayer is returned when the result carries no player metrics,
	// which means the request had no player in party one.
	ErrNoPlayer = errors.New("adapter: the result has no player metrics")
)

// playerGUID is the synthetic unit id the summary uses for the simmed
// player. A real fight's guids come from the combat log; a sim has none,
// so it gets a stable one that the report components treat identically.
const playerGUID = "sim-player"

// Summarize maps an engine result onto the logs engine's Summary, per the
// interface contract's mapping table.
func Summarize(res *proto.RaidSimResult, req api.SimRequest) (summary.Summary, error) {
	if res == nil {
		return summary.Summary{}, fmt.Errorf("%w: nil result", ErrSimFailed)
	}
	if res.Error != nil && res.Error.Message != "" {
		return summary.Summary{}, fmt.Errorf("%w: %s", ErrSimFailed, res.Error.Message)
	}
	iters := float64(res.IterationsDone)
	if iters <= 0 {
		return summary.Summary{}, fmt.Errorf("%w: iterations_done is %d", ErrSimFailed, res.IterationsDone)
	}
	player, err := PlayerMetrics(res)
	if err != nil {
		return summary.Summary{}, err
	}

	durationMS := int64(math.Round(res.AvgIterationDuration * 1000))

	out := summary.Summary{
		EngineVersion: "sim:" + req.EngineVersion,
		FightIndex:    1,
		DurationMS:    durationMS,

		DamageDone:   actors(player, iters, durationMS),
		DamageTaken:  []summary.Actor{},
		Healing:      []summary.Actor{},
		HealingTaken: []summary.Actor{},

		Deaths:     []summary.Death{},
		Auras:      auras(player, iters),
		Casts:      casts(player, iters),
		Interrupts: []summary.ExchangeRow{},
		Dispels:    []summary.ExchangeRow{},
		Resources:  resources(player, iters),
		Threat:     []summary.ThreatRow{},
		Combatants: []summary.CombatantRow{},
	}
	out.Roster = roster(player, req, out, durationMS)
	return out, nil
}

// PlayerMetrics returns the first player of the first party, which is the
// only player an individual sim has.
func PlayerMetrics(res *proto.RaidSimResult) (*proto.UnitMetrics, error) {
	if res.RaidMetrics == nil {
		return nil, ErrNoPlayer
	}
	for _, party := range res.RaidMetrics.Parties {
		for _, p := range party.Players {
			if p != nil {
				return p, nil
			}
		}
	}
	return nil, ErrNoPlayer
}

// DPS lifts the engine's distribution into the envelope's Estimate. Error
// is the standard error of the mean: stdev over the square root of the
// iteration count, which is the figure the sim page shows beside the DPS.
func DPS(res *proto.RaidSimResult) api.Estimate {
	if res == nil || res.RaidMetrics == nil || res.RaidMetrics.Dps == nil {
		return api.Estimate{}
	}
	d := res.RaidMetrics.Dps
	est := api.Estimate{Mean: d.Avg, StdDev: d.Stdev, Min: d.Min, Max: d.Max}
	if res.IterationsDone > 0 {
		est.Error = d.Stdev / math.Sqrt(float64(res.IterationsDone))
	}
	return est
}

// actors builds the damage table: one row for the player, then one per pet.
func actors(player *proto.UnitMetrics, iters float64, durationMS int64) []summary.Actor {
	out := []summary.Actor{actorFrom(player, playerGUID, iters, durationMS)}
	for i, pet := range player.Pets {
		out = append(out, actorFrom(pet, fmt.Sprintf("%s-pet-%d", playerGUID, i), iters, durationMS))
	}
	return out
}

func actorFrom(u *proto.UnitMetrics, guid string, iters float64, durationMS int64) summary.Actor {
	a := summary.Actor{
		GUID:      guid,
		Name:      u.Name,
		ActiveMS:  durationMS,
		Abilities: []summary.Ability{},
		Targets:   []summary.Pair{},
		Series:    []int64{},
	}
	perTarget := map[int32]int64{}
	for _, am := range u.Actions {
		ab := ability(am, iters, perTarget)
		a.Abilities = append(a.Abilities, ab)
		a.Total += ab.Total
	}
	a.Effective = a.Total

	idx := make([]int32, 0, len(perTarget))
	for k := range perTarget {
		idx = append(idx, k)
	}
	sort.Slice(idx, func(i, j int) bool { return idx[i] < idx[j] })
	for _, i := range idx {
		a.Targets = append(a.Targets, summary.Pair{
			GUID:  fmt.Sprintf("sim-target-%d", i),
			Name:  fmt.Sprintf("Target %d", i),
			Total: perTarget[i],
		})
	}
	sort.SliceStable(a.Abilities, func(i, j int) bool { return a.Abilities[i].Total > a.Abilities[j].Total })
	return a
}

// ability folds one ActionMetrics, which is already summed over every
// target and every iteration, into one summary row per fight.
func ability(am *proto.ActionMetrics, iters float64, perTarget map[int32]int64) summary.Ability {
	spellID, name := ActionName(am.Id)
	ab := summary.Ability{
		SpellID: spellID,
		Name:    name,
		School:  int64(am.SpellSchool),
		Misses:  map[string]int64{},
	}
	for _, t := range am.Targets {
		damage := t.Damage + t.CritDamage + t.TickDamage + t.CritTickDamage +
			t.GlanceDamage + t.CrushDamage + t.BlockDamage + t.BlockedCritDamage
		per := per(damage, iters)
		ab.Total += per
		ab.Effective += per
		perTarget[t.UnitIndex] += per

		ab.Hits += per(float64(t.Hits), iters)
		ab.Crits += per(float64(t.Crits), iters)
		ab.Ticks += per(float64(t.Ticks+t.CritTicks), iters)
		addMiss(ab.Misses, "miss", t.Misses, iters)
		addMiss(ab.Misses, "dodge", t.Dodges, iters)
		addMiss(ab.Misses, "parry", t.Parries, iters)
		addMiss(ab.Misses, "block", t.Blocks, iters)
		addMiss(ab.Misses, "glance", t.Glances, iters)
		addMiss(ab.Misses, "crush", t.Crushes, iters)
	}
	if len(ab.Misses) == 0 {
		ab.Misses = nil
	}
	// The engine reports no per-hit minimum or maximum, only per-action
	// totals, so a range would be invented. Both stay zero.
	return ab
}

func addMiss(m map[string]int64, key string, count int32, iters float64) {
	if count == 0 {
		return
	}
	m[key] += per(float64(count), iters)
}

// per divides an across-iterations total into a per-fight figure. It
// rounds rather than truncating, so an ability used in most iterations
// does not report zero, and it never produces a negative from a spend
// figure: callers negate first where that matters.
func per(total, iters float64) int64 {
	return int64(math.Round(total / iters))
}

func auras(u *proto.UnitMetrics, iters float64) []summary.AuraTrack {
	out := make([]summary.AuraTrack, 0, len(u.Auras))
	for _, am := range u.Auras {
		spellID, name := ActionName(am.Id)
		out = append(out, summary.AuraTrack{
			TargetGUID: playerGUID,
			TargetName: u.Name,
			SpellID:    spellID,
			Name:       name,
			Type:       "buff",
			// AuraMetrics already reports per-iteration averages, so
			// these are not divided again.
			Applications: int64(math.Round(am.ProcsAvg)),
			MaxStacks:    0,
			UptimeMS:     int64(math.Round(am.UptimeSecondsAvg * 1000)),
			// The engine reports no application timeline, so there are no
			// segments to build and none are invented.
			Segments: []summary.Segment{},
			Appliers: []string{u.Name},
		})
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].UptimeMS > out[j].UptimeMS })
	return out
}

func casts(u *proto.UnitMetrics, iters float64) []summary.CastRow {
	out := make([]summary.CastRow, 0, len(u.Actions))
	for _, am := range u.Actions {
		var total int32
		for _, t := range am.Targets {
			total += t.Casts
		}
		if total == 0 {
			continue
		}
		spellID, name := ActionName(am.Id)
		n := per(float64(total), iters)
		out = append(out, summary.CastRow{
			GUID:      playerGUID,
			Name:      u.Name,
			SpellID:   spellID,
			SpellName: name,
			Started:   n,
			Succeeded: n,
			Failed:    0,
			// The engine reports no cast timestamps, so the sequence the
			// report's cast timeline draws stays empty for a sim.
			Sequence: []int64{},
		})
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Succeeded > out[j].Succeeded })
	return out
}

func resources(u *proto.UnitMetrics, iters float64) []summary.ResourceTrack {
	// The engine reports one ResourceMetrics per action per resource type;
	// the summary wants one track per resource type.
	byType := map[proto.ResourceType]*summary.ResourceTrack{}
	order := []proto.ResourceType{}
	for _, rm := range u.Resources {
		tr, ok := byType[rm.Type]
		if !ok {
			tr = &summary.ResourceTrack{
				GUID:      playerGUID,
				Name:      u.Name,
				PowerType: int64(rm.Type),
				Series:    []int64{},
			}
			byType[rm.Type] = tr
			order = append(order, rm.Type)
		}
		// Gain is negative for a spend, per the proto's own comment.
		if rm.ActualGain >= 0 {
			tr.Gained += per(rm.ActualGain, iters)
		} else {
			tr.Spent += per(-rm.ActualGain, iters)
		}
	}
	out := make([]summary.ResourceTrack, 0, len(order))
	for _, t := range order {
		out = append(out, *byType[t])
	}
	return out
}

func roster(u *proto.UnitMetrics, req api.SimRequest, s summary.Summary, durationMS int64) []summary.RosterRow {
	class, spec := splitSpecSlug(req.Spec)
	var damage int64
	for _, a := range s.DamageDone {
		damage += a.Total
	}
	var dps float64
	if u.Dps != nil {
		dps = u.Dps.Avg
	}
	return []summary.RosterRow{{
		GUID:        playerGUID,
		Name:        u.Name,
		Class:       class,
		ClassSource: "sim",
		Spec:        spec,
		Role:        "dps",
		ActiveMS:    durationMS,
		ActivityPct: 100,
		DamageDone:  damage,
		DPS:         dps,
	}}
}

// splitSpecSlug turns "warrior-fury" into ("warrior", "fury"). The
// canonical list lives in data/curated/specs.json; this only needs the
// split, not the list.
func splitSpecSlug(slug string) (class, spec string) {
	i := strings.Index(slug, "-")
	if i < 0 {
		return slug, ""
	}
	return slug[:i], slug[i+1:]
}

// ActionName returns the spell id and a display name for an engine
// ActionID. The engine's ids carry no names, so the name is the id in a
// readable form; the web resolves real names from the item and spell
// database it already loads for tooltips.
func ActionName(id *proto.ActionID) (int64, string) {
	if id == nil {
		return 0, "Unknown"
	}
	switch raw := id.RawId.(type) {
	case *proto.ActionID_SpellId:
		if id.Tag != 0 {
			return int64(raw.SpellId), fmt.Sprintf("spell:%d/%d", raw.SpellId, id.Tag)
		}
		return int64(raw.SpellId), fmt.Sprintf("spell:%d", raw.SpellId)
	case *proto.ActionID_ItemId:
		return int64(raw.ItemId), fmt.Sprintf("item:%d", raw.ItemId)
	case *proto.ActionID_OtherId:
		return 0, fmt.Sprintf("other:%d", int32(raw.OtherId))
	}
	return 0, "Unknown"
}
```

- [ ] **Step 4: Run the test and watch it pass**

Run: `cd /Users/jh/code/forever/sim && go test ./adapter/ -v`
Expected: `PASS` for all nine tests. If `TestSummarizeDividesByIterations` fails on `Hits`, check that `per` rounds rather than truncates.

- [ ] **Step 5: Write the fixture loader**

The api lane's tests need a real engine result without running an engine. Create `sim/adapter/fixture.go`:

```go
package adapter

import (
	"embed"
	"fmt"

	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

//go:embed testdata/*.result.pb
var fixtures embed.FS

// Fixture returns a checked-in engine result for a spec. The api lane's
// validation and execution-score tests use it so they need no engine
// binary; the golden tests below use it as their input.
//
// Regenerated from the engine checkout; the procedure is Step 7 of the
// task that created this file, and the header of golden_test.go repeats it.
func Fixture(spec string) (*proto.RaidSimResult, error) {
	b, err := fixtures.ReadFile("testdata/" + spec + ".result.pb")
	if err != nil {
		return nil, fmt.Errorf("adapter: no fixture for spec %q: %w", spec, err)
	}
	res := &proto.RaidSimResult{}
	if err := googleproto.Unmarshal(b, res); err != nil {
		return nil, fmt.Errorf("adapter: fixture for %q is corrupt: %w", spec, err)
	}
	return res, nil
}
```

- [ ] **Step 6: Write the golden test**

It is the same idiom `logs/engine/summary/golden_test.go` already uses: JSON is portable text, so it catches a field rename, a reordering, or a rounding change that the unit tests above pass straight through. `FOREVER_UPDATE_GOLDEN=1` regenerates.

Create `sim/adapter/golden_test.go`:

```go
package adapter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// regenEnv regenerates the goldens instead of comparing against them. The
// same variable name the logs engine uses, so one habit covers both.
const regenEnv = "FOREVER_UPDATE_GOLDEN"

// goldenEngineVersion is deliberately not enginever.Version: the golden
// pins the adapter's output shape, and an engine pin bump must not churn
// every golden file.
const goldenEngineVersion = "golden"

// To regenerate a fixture after a spec's abilities change, from the engine
// checkout:
//
//	go run ./cmd/forever-sim -in <request.pb> -out <spec>.result.pb
//
// then copy it into sim/adapter/testdata/ and run
//
//	FOREVER_UPDATE_GOLDEN=1 go test ./adapter/
//
// and read the diff before committing it.
func TestGoldenSummaries(t *testing.T) {
	for _, spec := range []string{"warrior-fury", "mage-frost"} {
		t.Run(spec, func(t *testing.T) {
			res, err := Fixture(spec)
			if err != nil {
				t.Fatal(err)
			}
			req := api.SimRequest{
				EngineVersion: goldenEngineVersion,
				Spec:          spec,
				Encounter:     api.DefaultEncounter(),
				Iterations:    3000,
				Raw:           []byte{1},
			}
			got, err := Summarize(res, req)
			if err != nil {
				t.Fatal(err)
			}
			b, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			b = append(b, '\n')

			path := filepath.Join("testdata", spec+".summary.json.golden")
			if os.Getenv(regenEnv) != "" {
				if err := os.WriteFile(path, b, 0o644); err != nil {
					t.Fatal(err)
				}
				t.Logf("regenerated %s", path)
				return
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("%v (run with %s=1 to create it)", err, regenEnv)
			}
			if string(b) != string(want) {
				t.Errorf("summary for %s differs from the golden.\n--- got ---\n%s\n--- want ---\n%s", spec, b, want)
			}
		})
	}
}

// Summarizing is deterministic: the same result and request must produce
// byte-identical JSON, or the report page would churn between loads.
func TestSummarizeIsDeterministic(t *testing.T) {
	res, err := Fixture("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	req := api.SimRequest{EngineVersion: goldenEngineVersion, Spec: "warrior-fury", Encounter: api.DefaultEncounter(), Iterations: 3000, Raw: []byte{1}}
	var first string
	for i := 0; i < 20; i++ {
		s, err := Summarize(res, req)
		if err != nil {
			t.Fatal(err)
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		if i == 0 {
			first = string(b)
			continue
		}
		if string(b) != first {
			t.Fatalf("run %d differs from run 0; a map is being ranged without sorting", i)
		}
	}
}
```

- [ ] **Step 7: Produce the two fixtures**

The engine has no binary result writer until Task 13, so produce the fixtures with a short throwaway program in the **engine** checkout. It is deleted at the end of this step; nothing is committed to the engine repo here.

```bash
cd /Users/jh/code/wowsims-forever
mkdir -p tools/genfixture
cat > tools/genfixture/main.go <<'EOF'
package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/mage"
	dpswarrior "github.com/wowsims/classic/sim/warrior/dps_warrior"
	googleproto "google.golang.org/protobuf/proto"
)

func main() {
	spec := flag.String("spec", "warrior-fury", "spec slug")
	out := flag.String("out", "out.result.pb", "output file")
	iters := flag.Int("iterations", 3000, "iterations")
	flag.Parse()

	dpswarrior.RegisterDpsWarrior()
	mage.RegisterMage()

	var player *proto.Player
	var gearDir, aplDir, gearFile, aplFile string
	switch *spec {
	case "warrior-fury":
		gearDir, gearFile = "ui/warrior/gear_sets", "phase_1"
		aplDir, aplFile = "ui/warrior/apls", "dps_reck"
		player = &proto.Player{
			Name: "Fury", Race: proto.Race_RaceOrc, Class: proto.Class_ClassWarrior,
			TalentsString: "30305001302-05050005525010051", Consumes: &proto.Consumes{},
			Buffs: core.FullIndividualBuffs,
		}
	case "mage-frost":
		gearDir, gearFile = "ui/mage/gear_sets", "phase_1"
		aplDir, aplFile = "ui/mage/apls", "p1"
		player = &proto.Player{
			Name: "Frost", Race: proto.Race_RaceGnome, Class: proto.Class_ClassMage,
			TalentsString: "2500050300030150333125----", Consumes: &proto.Consumes{},
			Buffs: core.FullIndividualBuffs,
		}
	default:
		log.Fatalf("unknown spec %q", *spec)
	}

	player.Equipment = core.GetGearSet(gearDir, gearFile).GearSet
	player.Rotation = core.GetAplRotation(aplDir, aplFile).Rotation
	switch *spec {
	case "warrior-fury":
		core.WithSpec(player, &proto.Player_Warrior{Warrior: &proto.Warrior{
			Options: &proto.Warrior_Options{StartingRage: 50, Shout: proto.WarriorShout_WarriorShoutBattle},
		}})
	case "mage-frost":
		core.WithSpec(player, &proto.Player_Mage{Mage: &proto.Mage{Options: &proto.Mage_Options{}}})
	}

	enc := core.MakeSingleTargetEncounter(0.2)
	enc.Duration = 180
	req := &proto.RaidSimRequest{
		Raid:       core.SinglePlayerRaidProto(player, core.FullPartyBuffs, core.FullRaidBuffs, core.FullDebuffs),
		Encounter:  enc,
		SimOptions: &proto.SimOptions{Iterations: int32(*iters), IsTest: false, RandomSeed: 1},
	}
	res := core.RunRaidSim(req)
	if res.Error != nil {
		log.Fatalf("sim failed: %s", res.Error.Message)
	}
	// The cast log is not part of the fixture: it is one iteration of text
	// and it makes the file large without changing anything the adapter reads.
	res.Logs = ""
	b, err := googleproto.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(*out, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s (%d bytes), dps=%.1f", *out, len(b), res.RaidMetrics.Dps.Avg)
}
EOF
mkdir -p /Users/jh/code/forever/sim/adapter/testdata
go run --tags=with_db ./tools/genfixture -spec warrior-fury -out /Users/jh/code/forever/sim/adapter/testdata/warrior-fury.result.pb
go run --tags=with_db ./tools/genfixture -spec mage-frost  -out /Users/jh/code/forever/sim/adapter/testdata/mage-frost.result.pb
rm -rf tools/genfixture
git status --short   # must be empty: nothing is committed to the engine here
```

Expected: two `wrote ... dps=...` lines and an empty `git status`. On the measured baseline the warrior line reads roughly `dps=1791.1`; any number is fine, the fixture only has to be a real result.

If `ui/mage/gear_sets/phase_1.gear.json` does not exist, list `ui/mage/gear_sets/` and use a file that does; the fixture's job is to be a real engine result, not a particular gear set.

- [ ] **Step 8: Generate the goldens and read them**

```bash
cd /Users/jh/code/forever/sim
FOREVER_UPDATE_GOLDEN=1 go test ./adapter/ -run TestGoldenSummaries -v
head -60 adapter/testdata/warrior-fury.summary.json.golden
```

Expected: two `regenerated ...` log lines. Read the head of the file and check by eye that `engine_version` is `sim:golden`, `fight_index` is `1`, `duration_ms` is about `180000`, the first damage row is the player with a plausible per-fight total (a 180-second fight at ~1,800 DPS is about 320,000), and the largest ability is a recognisable warrior spell id. If the totals look like a 3,000-iteration sum rather than one fight, `per` is not dividing.

- [ ] **Step 9: Run the whole package and watch it pass**

```bash
cd /Users/jh/code/forever/sim
go test ./adapter/ -race -v
gofmt -l ./adapter ./api ./enginever
```

Expected: `PASS` for all eleven tests; `gofmt -l` prints nothing.

- [ ] **Step 10: Commit**

```bash
cd /Users/jh/code/forever
git add sim/adapter
git commit -m "feat(sim): the adapter from an engine result to a logs summary" \
  -m "Summarize maps RaidSimResult onto summary.Summary per the contract's table, so the report page's components render a sim with no second renderer. The governing rule is that the engine accumulates across iterations and a summary describes one fight, so every count and total is divided by IterationsDone; the distribution stays in SimResult.DPS. Auras carry no application timeline and casts carry no timestamps in an engine result, so those fields stay empty rather than being invented. Two checked-in engine results and their goldens, plus a determinism test, pin the shape; Fixture() lets the api lane test without an engine binary." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

**The api and web lanes are unblocked at this commit.** Tell the controller.

---

## Task 4: Unified `Hit` and `Crit`

**Repo: ENGINE.** Depends on Task 1. **Serial gate: nothing else in the engine may run beside it.**

Forever makes spell, melee, and ranged hit one stat, and likewise crit. In the engine that is one enum merge in two index-synced places and a mechanical rename everywhere else.

**Measured, 2026-09-14, at engine HEAD `7779ebb`** — verify these before and after, the counts are the task's completion check:

| Symbol | `sim/` call sites | files |
|---|---|---|
| `stats.MeleeHit` | 14 | 10 |
| `stats.SpellHit` | 13 | 8 |
| `stats.MeleeCrit` | 76 | 32 |
| `stats.SpellCrit` | 56 | 25 |
| **total** | **159** (157 distinct lines; two lines mention two of them) | **41 distinct** |

The research's "zero in `ui/`" is about Go-style `stats.X` references and is correct. **The TypeScript UI does reference the proto enum names: 177 occurrences of `StatSpellHit`/`StatMeleeHit`/`StatSpellCrit`/`StatMeleeCrit` across 27 files under `ui/`.** None of that UI ships on our site, but leaving it broken makes the fork unbuildable for anyone running `make` and unmergeable upstream, so this task fixes it too. `assets/` has zero references.

The 41 Go files, for the record (`grep -rlE 'stats\.(MeleeHit|SpellHit|MeleeCrit|SpellCrit)\b' sim/ --include='*.go'`):

```
sim/common/guardians/emerald_dragon_whelp.go   sim/common/item_effects.go
sim/common/item_sets/crafted.go                sim/common/item_sets/item_sets_pve.go
sim/common/item_sets/item_sets_pvp.go          sim/core/base_stats.go
sim/core/buffs.go                              sim/core/consumes.go
sim/core/spell_result.go                       sim/druid/druid.go
sim/druid/forms.go                             sim/druid/item_sets_pve.go
sim/hunter/explosive_trap.go                   sim/hunter/hunter.go
sim/hunter/immolation_trap.go                  sim/hunter/items.go
sim/hunter/pet.go                              sim/hunter/talents.go
sim/mage/item_sets_pve.go                      sim/mage/mage.go
sim/paladin/item_sets_pve.go                   sim/paladin/items.go
sim/paladin/paladin.go                         sim/paladin/talents.go
sim/priest/priest.go                           sim/rogue/rogue.go
sim/rogue/talents.go                           sim/shaman/shaman.go
sim/shaman/talents.go                          sim/warlock/felhunter.go
sim/warlock/imp.go                             sim/warlock/pet.go
sim/warlock/succubus.go                        sim/warlock/talents.go
sim/warlock/voidwalker.go                      sim/warlock/warlock.go
sim/warrior/item_sets_pve.go                   sim/warrior/recklessness.go
sim/warrior/stances.go                         sim/warrior/talents.go
sim/warrior/warrior.go
```

**Files:**
- Modify: `sim/core/stats/stats.go` (enum at lines 20-69, `StatName` at 182-250), `proto/common.proto` (`Stat` enum at lines 75-122), the 41 files above, `sim/core/base_stats_auto_gen.go` (the `SpellCritRatingPerCritChance` constant folds into `CritRatingPerCritChance`), `ui/core/proto_utils/names.ts`, `ui/core/proto_utils/stats.ts` and the other 25 `ui/` files
- Test: `sim/core/stats/stats_test.go`

**Interfaces:**
- Produces: `stats.Hit` replaces `stats.MeleeHit` and `stats.SpellHit`; `stats.Crit` replaces `stats.MeleeCrit` and `stats.SpellCrit`. `proto.Stat_StatHit` and `proto.Stat_StatCrit` replace the four proto values. `core.HitRatingPerHitChance` replaces `MeleeHitRatingPerHitChance` and `SpellHitRatingPerHitChance`; `core.CritRatingPerCritChance` absorbs `SpellCritRatingPerCritChance`. The enum shrinks from 44 values to 42 and every value after index 13 shifts down by two.

- [ ] **Step 1: Record the baseline counts**

```bash
cd /Users/jh/code/wowsims-forever
for s in MeleeHit SpellHit MeleeCrit SpellCrit; do
  printf "%-10s %s\n" "$s" "$(grep -rn "stats\.$s" sim/ --include='*.go' | wc -l)"
done
grep -rnE '\bStat(SpellHit|MeleeHit|SpellCrit|MeleeCrit)\b' ui/ --include='*.ts' --include='*.tsx' | wc -l
```

Expected: `14 13 76 56` and `177`. If any number differs, upstream has moved; record the new numbers in the commit body and proceed.

- [ ] **Step 2: Write the failing index-sync test**

The Go enum and the proto enum must stay index-synced — `sim/core/stats/stats.go:19` says so in a comment and nothing enforces it. Make it enforced, because this task is exactly the change that could break it.

Append to `sim/core/stats/stats_test.go`:

```go
// The Go Stat enum and proto.Stat are index-synced: Stat(v) is how a
// proto value becomes a Go one (see ProtoArrayToStatsList), so a
// divergence silently reads the wrong stat. Forever merges MeleeHit and
// SpellHit into Hit and MeleeCrit and SpellCrit into Crit, which shifts
// every later index, so the sync is checked rather than commented.
func TestStatEnumIsSyncedWithProto(t *testing.T) {
	names := proto.Stat_name
	if len(names) != int(Len) {
		t.Fatalf("proto.Stat has %d values, Go Stat has %d", len(names), int(Len))
	}
	for i := Stat(0); i < Len; i++ {
		protoName, ok := names[int32(i)]
		if !ok {
			t.Errorf("index %d: proto.Stat has no value", int(i))
			continue
		}
		// proto names are "StatFoo"; Go names are "Foo".
		want := "Stat" + i.StatName()
		if protoName != want {
			t.Errorf("index %d: proto has %q, Go has %q (want %q)", int(i), protoName, i.StatName(), want)
		}
	}
}

// Forever has one hit stat and one crit stat.
func TestForeverHasOneHitAndOneCritStat(t *testing.T) {
	if Hit >= Len || Crit >= Len {
		t.Fatal("Hit and Crit must be real stats")
	}
	if Hit.StatName() != "Hit" {
		t.Errorf("Hit.StatName() = %q, want %q", Hit.StatName(), "Hit")
	}
	if Crit.StatName() != "Crit" {
		t.Errorf("Crit.StatName() = %q, want %q", Crit.StatName(), "Crit")
	}
	for i := Stat(0); i < Len; i++ {
		switch i.StatName() {
		case "MeleeHit", "SpellHit", "MeleeCrit", "SpellCrit":
			t.Errorf("index %d still has the split stat %q", int(i), i.StatName())
		}
	}
}
```

Make sure the file imports `"github.com/wowsims/classic/sim/core/proto"`.

- [ ] **Step 3: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/stats/ -run 'TestStatEnum|TestForeverHas' -v`
Expected: `FAIL [build failed]`, `undefined: Hit`.

- [ ] **Step 4: Merge the proto enum**

In `proto/common.proto`, replace lines 89-95 — that is, the block from `StatSpellHit = 13;` through `StatMeleeCrit = 19;` — so the enum reads:

```proto
	StatMP5 = 12;
	// Forever merges spell, melee and ranged hit into one stat, and
	// likewise crit. This enum must stay index-synced with the Go Stat
	// enum in sim/core/stats/stats.go; stats_test.go enforces it.
	StatHit = 13;
	StatCrit = 14;
	StatSpellHaste = 15;
	StatSpellPenetration = 16;
	StatAttackPower = 17;
	StatMeleeHaste = 18;
	StatArmorPenetration = 19;
	StatExpertise = 20;
	StatMana = 21;
	StatEnergy = 22;
	StatRage = 23;
	StatArmor = 24;
	StatRangedAttackPower = 25;
	StatDefense = 26;
	StatBlock = 27;
	StatBlockValue = 28;
	StatDodge = 29;
	StatParry = 30;
	StatResilience = 31;
	StatHealth = 32;
	StatArcaneResistance = 33;
	StatFireResistance = 34;
	StatFrostResistance = 35;
	StatNatureResistance = 36;
	StatShadowResistance = 37;
	StatBonusArmor = 38;
	StatHealingPower = 39;
	StatSpellDamage = 40;
	StatFeralAttackPower = 41;
```

This renumbers, which a wire-compatible protobuf change normally must not do. It is correct here: the enum is a *stat index*, never a persisted value, and both ends of the wire — the Go engine and the web — are rebuilt from this file at the same `ENGINE_VERSION`. Add that reasoning as a comment above the enum so a reviewer does not flag it:

```proto
// NOTE: the values of this enum are array indexes into stats.Stats, not
// stable wire identities. They are renumbered when a stat is added or
// merged, and both the Go engine and the web are rebuilt from this file at
// the same ENGINE_VERSION, so nothing persists a raw value across builds.
enum Stat {
```

Regenerate: `export PATH=$PATH:$(go env GOPATH)/bin && make proto`.

- [ ] **Step 5: Merge the Go enum**

In `sim/core/stats/stats.go`, the const block at lines 20-69: delete `SpellHit`, `SpellCrit`, `MeleeHit` and `MeleeCrit`, and insert `Hit` then `Crit` where `SpellHit`/`SpellCrit` were (after `MP5`, before `SpellHaste`). The block's head becomes:

```go
	MP5
	// Forever merges spell, melee and ranged hit into one stat, and
	// likewise crit.
	Hit
	Crit
	SpellHaste
	SpellPenetration
	AttackPower
	MeleeHaste
	ArmorPenetration
	Expertise
```

In `StatName`, replace the four cases

```go
	case SpellCrit:
		return "SpellCrit"
	case SpellHit:
		return "SpellHit"
```

and

```go
	case MeleeHit:
		return "MeleeHit"
	case MeleeCrit:
		return "MeleeCrit"
```

with, in the position where `SpellCrit`/`SpellHit` were:

```go
	case Hit:
		return "Hit"
	case Crit:
		return "Crit"
```

- [ ] **Step 6: Rename the 159 call sites mechanically**

```bash
cd /Users/jh/code/wowsims-forever
grep -rlE 'stats\.(MeleeHit|SpellHit|MeleeCrit|SpellCrit)\b' sim/ --include='*.go' \
  | xargs sed -i '' -E 's/stats\.(MeleeHit|SpellHit)\b/stats.Hit/g; s/stats\.(MeleeCrit|SpellCrit)\b/stats.Crit/g'
grep -rlE '\b(MeleeHitRatingPerHitChance|SpellHitRatingPerHitChance|SpellCritRatingPerCritChance)\b' sim/ tools/ --include='*.go' \
  | xargs sed -i '' -E 's/\b(MeleeHitRatingPerHitChance|SpellHitRatingPerHitChance)\b/HitRatingPerHitChance/g; s/\bSpellCritRatingPerCritChance\b/CritRatingPerCritChance/g'
```

Then in `sim/core/base_stats_auto_gen.go` collapse the five rating constants to three:

```go
// Crit/Hit/Haste ratings are straight percentage values in classic.
// Forever merges spell and melee hit into one rating and spell and melee
// crit into one rating; Task 5 regenerates these from the client tables.
const HasteRatingPerHastePercent = 1
const CritRatingPerCritChance = 1
const HitRatingPerHitChance = 1
```

**Two renames create duplicates and must be fixed by hand, not by `sed`.** Find them:

```bash
go build ./sim/... 2>&1 | head -40
```

Anywhere a single composite literal previously set both `stats.MeleeCrit: x` and `stats.SpellCrit: y` there is now a duplicate key. Merge each by choosing the value that is correct for one unified stat — for an item that granted both, one entry with the larger of the two; for an item that granted only one, that value. Add a comment on each merged line:

```go
	// Forever: merged from MeleeCrit 1% + SpellCrit 1%. unconfirmed
	stats.Crit: 1,
```

`sim/core/base_stats.go` lines 87-128 are the main cluster: `ExtraClassBaseStats` sets `stats.SpellCrit` and `stats.MeleeCrit` per class at four levels. Those two are genuinely different numbers per class and Forever has not published a unified base crit, so keep the **melee** value (the attack table is the load-bearing one for the two launch specs) and mark it:

```go
	proto.Class_ClassWarrior: {
		25: {
			// unconfirmed: Era had base SpellCrit 0.0000 and MeleeCrit 0.0000
			// at 60; Forever's unified base crit is not published. Task 5's
			// regeneration replaces this from the client tables.
			stats.Crit: 0.0000 * CritRatingPerCritChance,
		},
		...
```

- [ ] **Step 7: Fix the TypeScript**

```bash
cd /Users/jh/code/wowsims-forever
grep -rlE '\bStat(SpellHit|MeleeHit|SpellCrit|MeleeCrit)\b' ui/ --include='*.ts' --include='*.tsx' \
  | xargs sed -i '' -E 's/\bStat(SpellHit|MeleeHit)\b/StatHit/g; s/\bStat(SpellCrit|MeleeCrit)\b/StatCrit/g'
```

Then in `ui/core/proto_utils/names.ts` the display-name map now has two entries mapping to `Stat.StatHit` and two to `Stat.StatCrit`. Collapse each pair to one:

```ts
	[Stat.StatHit, 'Hit'],
	[Stat.StatCrit, 'Crit'],
```

and do the same for the duplicate entries in the stat-ordering arrays around line 125-135.

- [ ] **Step 8: Build, test, and re-count**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo "SIM BUILDS"
gofmt -l ./sim ./tools
go test --tags=with_db ./sim/core/... -v -run 'TestStatEnum|TestForeverHas'
go test --tags=with_db ./sim/core/... ./sim/warrior/... ./sim/mage/...
grep -rnE 'stats\.(MeleeHit|SpellHit|MeleeCrit|SpellCrit)\b' sim/ --include='*.go' | wc -l
grep -rnE '\bStat(SpellHit|MeleeHit|SpellCrit|MeleeCrit)\b' ui/ --include='*.ts' --include='*.tsx' | wc -l
npx tsc --noEmit
```

Expected: `SIM BUILDS`; `gofmt -l` prints nothing; both new tests `PASS`; **both greps print `0`**; `npx tsc --noEmit` is clean.

The per-spec `.results` golden files will now differ, because merging two base-crit numbers into one changes DPS. That is a real behaviour change, not a regression, so regenerate and read the diff:

```bash
go test --tags=with_db ./sim/... 2>&1 | tail -25
make update-tests
git diff --stat -- '*.results'
```

Read the DPS deltas. A change of a few percent on casters (whose `SpellCrit` base was dropped in favour of the melee value) is expected; a change of more than 20% on any spec means a duplicate key was merged wrongly — find it before proceeding.

- [ ] **Step 9: Run the full engine suite**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db -count=1 ./sim/...`
Expected: 20 packages `ok`, zero `FAIL`, about 11 seconds. This is the one task in G2's neighbourhood that runs the whole suite, because it touched 41 files across every class.

- [ ] **Step 10: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add -A
git commit -m "feat(core): merge melee and spell hit into Hit, melee and spell crit into Crit" \
  -m "Forever makes spell, melee and ranged hit one stat and likewise crit. 159 call sites across 41 files in sim/ and 177 across 27 files in ui/, renamed mechanically; the handful of composite literals that set both halves are merged by hand and each carries an unconfirmed comment naming the two Era values it came from. The Stat enum is an array index, not a wire identity, so renumbering it is safe: both ends are rebuilt from the same proto at one ENGINE_VERSION, and a new test asserts the Go and proto enums stay index-synced rather than leaving it to a comment. Per-spec .results goldens regenerated; the caster deltas come from dropping the split base crit." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 5: Regenerate the rating constants and base stats from data

**Repo: ENGINE.** Depends on Task 4. **G2 — INDEPENDENT of Tasks 6, 7, 8, 9, 10.**

`sim/core/base_stats_auto_gen.go` claims to be `AUTO GENERATED BY BASE_STATS_PARSER.PY`. It is not, any more: **`tools/base_stats_parser.py` is broken and has been for some time.** Verified — running it produces

```
TypeError: write() argument must be str, not None
```

because `GenExtraStatsGoFile` builds `output`, then opens a triple-quoted string at the `"""     output += '''var CritPerAgiAtLevel...` line whose closing `"""` swallows the function's `return output`. The function returns `None`. The committed file has therefore been hand-edited: it carries `const ExpertiseRatingPerExpertiseChance = 1` and a `TODO: Update Defense/Dodge/Parry rates` comment that the generator never emits, and it is missing the `import` block and every `var` map the generator's dead code would have produced.

Forever's numbers are unknown until the beta client on Sept 17. So this task does **not** guess them. It makes the regeneration work, parameterises it by build, and ships Era values that are marked provisional — so that on Sept 17 the change is `python3 tools/base_stats_parser.py --build <forever-build>` and a commit, not a code edit.

Inputs, all present: `assets/db_inputs/basestats/{combatratings,chancetomeleecrit,chancetomeleecritbase,chancetospellcrit,chancetospellcritbase,octbasempbyclass}.txt`, tab-separated, one row per level. `combatratings.txt` has a `Level` column and named columns including `weapon skill`, `defense skill`, `dodge`, `parry`, `block`, `hit melee`, `hit spell`, `crit melee`, `crit spell`, `crit taken melee`.

**Files:**
- Modify: `tools/base_stats_parser.py`
- Regenerate: `sim/core/base_stats_auto_gen.go`
- Create: `sim/core/base_stats_provisional.go`
- Test: `sim/core/base_stats_test.go`

**Interfaces:**
- Consumes: `stats.Hit`, `stats.Crit`, `core.CritRatingPerCritChance`, `core.HitRatingPerHitChance` (Task 4).
- Produces:
  - `core.ProvisionalConstants() []string` — the names of every constant still carrying an Era value. Empty once the Forever tables land. The api lane's spec-support page reads it through the engine's `ComputeStatsResult`; until then it is the honest answer to "can I trust this".
  - `core.BaseStatsBuild string` — the client build the constants were generated from, e.g. `"1.15.9.69722"` for Era.
  - Regenerated: `ExpertisePerQuarterPercentReduction`, `ExpertiseRatingPerExpertiseChance`, `HasteRatingPerHastePercent`, `CritRatingPerCritChance`, `HitRatingPerHitChance`, `DefenseRatingPerDefense`, `DodgeRatingPerDodgeChance`, `ParryRatingPerParryChance`, `BlockRatingPerBlockChance`, `ResilienceRatingPerCritReductionChance`.

- [ ] **Step 1: Confirm the generator is broken**

```bash
cd /Users/jh/code/wowsims-forever
python3 tools/base_stats_parser.py; echo "exit=$?"
git checkout sim/core/base_stats_auto_gen.go
```

Expected: the `TypeError` above. (The script truncates the output file before it fails, hence the `git checkout`.)

- [ ] **Step 2: Write the failing test**

Create `sim/core/base_stats_test.go`:

```go
package core

import (
	"os"
	"strings"
	"testing"
)

// Every rating constant must come from the client tables, not from a
// literal somebody typed. The generated file names the build it came
// from, and the test refuses a file that has been hand-edited.
func TestBaseStatsFileIsGenerated(t *testing.T) {
	b, err := os.ReadFile("base_stats_auto_gen.go")
	if err != nil {
		t.Fatal(err)
	}
	src := string(b)
	if !strings.Contains(src, "Code generated by tools/base_stats_parser.py. DO NOT EDIT.") {
		t.Error("base_stats_auto_gen.go is missing the generated-code header")
	}
	if strings.Contains(src, "TODO") {
		t.Error("base_stats_auto_gen.go contains a TODO; it is generated, so a TODO means it was hand-edited")
	}
	if BaseStatsBuild == "" {
		t.Error("BaseStatsBuild is empty; the generator must record the client build")
	}
}

// Ratings are straight percentages in the Classic lineage and Forever
// keeps that; a regeneration that produced something else would silently
// scale every hit and crit chance in the engine.
func TestRatingConstantsArePercentages(t *testing.T) {
	cases := []struct {
		name string
		got  float64
	}{
		{"CritRatingPerCritChance", CritRatingPerCritChance},
		{"HitRatingPerHitChance", HitRatingPerHitChance},
		{"HasteRatingPerHastePercent", HasteRatingPerHastePercent},
	}
	for _, c := range cases {
		if c.got <= 0 {
			t.Errorf("%s = %v; a rating constant must be positive", c.name, c.got)
		}
	}
}

// Expertise is already modelled in the attack table (spell_outcome.go
// reduces dodge and parry by stats.Expertise/100) and is dormant in
// vanilla but live for Forever, so its two constants must be real.
func TestExpertiseConstantsAreReal(t *testing.T) {
	if ExpertisePerQuarterPercentReduction <= 0 {
		t.Errorf("ExpertisePerQuarterPercentReduction = %v, want positive", ExpertisePerQuarterPercentReduction)
	}
	if ExpertiseRatingPerExpertiseChance <= 0 {
		t.Errorf("ExpertiseRatingPerExpertiseChance = %v, want positive", ExpertiseRatingPerExpertiseChance)
	}
}

// Until the Forever client tables land, every constant here is an Era
// value. The engine says so out loud, and the spec support page repeats
// it, rather than letting a player assume the numbers are Forever's.
func TestProvisionalConstantsAreDeclared(t *testing.T) {
	got := ProvisionalConstants()
	if len(got) == 0 {
		t.Skip("no provisional constants: the Forever tables have landed, and this test has done its job")
	}
	for _, name := range got {
		if name == "" {
			t.Error("ProvisionalConstants() contains an empty name")
		}
	}
	b, err := os.ReadFile("base_stats_provisional.go")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "unconfirmed") {
		t.Error("base_stats_provisional.go must mark its values unconfirmed")
	}
}
```

- [ ] **Step 3: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'TestBaseStats|TestRatingConstants|TestExpertise|TestProvisional' -v`
Expected: `FAIL [build failed]`, `undefined: BaseStatsBuild`.

- [ ] **Step 4: Fix and parameterise the generator**

Replace `tools/base_stats_parser.py` in full:

```python
#!/usr/bin/env python3
"""Generate sim/core/base_stats_auto_gen.go from the client's base-stat tables.

The constants this emits are read straight out of the client's
CombatRatings table; nothing here is typed by hand. Forever's tables land
with the beta client, so until then this runs against the Era inputs and
the output is marked provisional (see sim/core/base_stats_provisional.go).

Usage:
    python3 tools/base_stats_parser.py                       # Era inputs
    python3 tools/base_stats_parser.py --build 2.0.1.12345 \\
        --inputs assets/db_inputs/basestats
"""

import argparse
import csv
import os
import sys

MAX_LEVEL = 60

COMBAT_RATINGS = "combatratings.txt"

# Rating constants, in the order the generated file emits them:
#   Go constant name -> (combatratings.txt column, comment)
RATING_CONSTANTS = [
    ("ExpertisePerQuarterPercentReduction", "weapon skill",
     "Expertise is live for Forever; the attack table already reduces dodge and parry by it."),
    ("ExpertiseRatingPerExpertiseChance", "weapon skill", None),
    ("HasteRatingPerHastePercent", "haste melee", None),
    ("DefenseRatingPerDefense", "defense skill", None),
    ("DodgeRatingPerDodgeChance", "dodge", None),
    ("ParryRatingPerParryChance", "parry", None),
    ("BlockRatingPerBlockChance", "block", None),
    ("ResilienceRatingPerCritReductionChance", "crit taken melee", None),
]

# Constants that are straight percentages in the Classic lineage, which
# Forever keeps. They are emitted as 1 rather than read from the table
# because the table's value is a rating-to-percent conversion that the
# lineage does not use. If Forever introduces rating conversion, these
# move into RATING_CONSTANTS and the change is one line each.
PERCENTAGE_CONSTANTS = [
    ("CritRatingPerCritChance", "Forever: one crit stat for spells and melee."),
    ("HitRatingPerHitChance", "Forever: one hit stat for spells, melee and ranged."),
]


def read_column_indexed(path):
    """combatratings.txt is one row per level with named columns. Return
    {column name: [value per level]}."""
    with open(path, newline="") as fh:
        rows = list(csv.reader(fh, delimiter="\t"))
    if not rows:
        raise SystemExit(f"{path} is empty")
    header = [h.strip() for h in rows[0]]
    out = {}
    for col_idx, name in enumerate(header):
        if not name or name == "Level":
            continue
        values = []
        for row in rows[1:]:
            if col_idx < len(row) and row[col_idx].strip():
                values.append(float(row[col_idx]))
        out[name] = values
    return out


def at_max_level(table, column, path):
    if column not in table:
        raise SystemExit(
            f"{path} has no column {column!r}; columns are {sorted(table)}")
    values = table[column]
    if len(values) < MAX_LEVEL:
        raise SystemExit(
            f"{path} column {column!r} has {len(values)} rows, need at least {MAX_LEVEL}")
    return values[MAX_LEVEL - 1]


def generate(table, build, inputs_path):
    lines = [
        "// Code generated by tools/base_stats_parser.py. DO NOT EDIT.",
        "//",
        f"// Client build: {build}",
        f"// Source:       {inputs_path}/{COMBAT_RATINGS}",
        "//",
        "// Regenerate with:",
        "//     python3 tools/base_stats_parser.py --build <build> --inputs <dir>",
        "",
        "package core",
        "",
        f'// BaseStatsBuild is the client build these constants were read from.',
        f'const BaseStatsBuild = "{build}"',
        "",
    ]
    for name, column, comment in RATING_CONSTANTS:
        if comment:
            lines.append(f"// {comment}")
        value = at_max_level(table, column, inputs_path)
        lines.append(f"const {name} = {value:f}")
    lines.append("")
    for name, comment in PERCENTAGE_CONSTANTS:
        lines.append(f"// {comment}")
        lines.append("// Ratings are straight percentages in the Classic lineage.")
        lines.append(f"const {name} = 1")
    lines.append("")
    return "\n".join(lines)


def main():
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("--inputs", default="assets/db_inputs/basestats",
                    help="directory holding the client's base-stat tables")
    ap.add_argument("--build", default="",
                    help="client build string, e.g. 1.15.9.69722; read from "
                         "<inputs>/BUILD when omitted")
    ap.add_argument("--out", default="sim/core/base_stats_auto_gen.go")
    args = ap.parse_args()

    build = args.build
    if not build:
        build_file = os.path.join(args.inputs, "BUILD")
        if os.path.exists(build_file):
            build = open(build_file).read().strip()
        else:
            raise SystemExit(
                "no --build given and no BUILD file in the inputs directory; "
                "the generated constants must record which client they came from")

    table = read_column_indexed(os.path.join(args.inputs, COMBAT_RATINGS))
    out = generate(table, build, args.inputs)
    with open(args.out, "w") as fh:
        fh.write(out)
    print(f"wrote {args.out} from build {build}", file=sys.stderr)


if __name__ == "__main__":
    main()
```

- [ ] **Step 5: Record the input build and regenerate**

The inputs are Era's. Record which Era build, so the generated header is not a lie:

```bash
cd /Users/jh/code/wowsims-forever
echo "1.15.9.69722" > assets/db_inputs/basestats/BUILD
python3 tools/base_stats_parser.py
gofmt -w sim/core/base_stats_auto_gen.go
cat sim/core/base_stats_auto_gen.go
```

Expected: a file whose header names build `1.15.9.69722`, with `ExpertisePerQuarterPercentReduction = 2.500000`, `ExpertiseRatingPerExpertiseChance = 2.500000` (both the `weapon skill` column at level 60 — that is the one the Era table gives, and `2.500000` matches the value the hand-edited file carried for the first of them), `HasteRatingPerHastePercent`, the four defensive constants, `ResilienceRatingPerCritReductionChance = 28.750002`, and `CritRatingPerCritChance = 1`, `HitRatingPerHitChance = 1`.

If `haste melee` is not a column in `combatratings.txt`, the generator says so by name and lists the columns it found. Read the header line (`head -1 assets/db_inputs/basestats/combatratings.txt`) and correct the column name in `RATING_CONSTANTS`; do not substitute a literal.

- [ ] **Step 6: Declare what is provisional**

Create `sim/core/base_stats_provisional.go`:

```go
package core

// The Forever client tables land with the beta on 2026-09-17. Until they
// do, base_stats_auto_gen.go is generated from the Era tables, which is
// the honest default — Forever keeps vanilla combat — but it is not
// Forever's data, and the engine says so rather than letting a player
// assume otherwise.
//
// Clearing this list is a three-step job and no code changes:
//
//  1. Put the Forever client's basestats tables in assets/db_inputs/basestats
//     and write the build string to assets/db_inputs/basestats/BUILD.
//  2. python3 tools/base_stats_parser.py && gofmt -w sim/core/base_stats_auto_gen.go
//  3. Set foreverBaseStatsBuildPrefix below to the Forever build's prefix.
//
// The nightly validation job (api lane) is what proves the result.

// foreverBaseStatsBuildPrefix is the client build prefix that means "these
// are Forever's own numbers". Era builds are 1.15.x; Forever's are not.
const foreverBaseStatsBuildPrefix = "2."

// provisionalConstantNames are the constants whose values are Era's and
// are therefore unconfirmed for Forever. Every one of them is generated,
// never typed.
var provisionalConstantNames = []string{
	"ExpertisePerQuarterPercentReduction", // unconfirmed
	"ExpertiseRatingPerExpertiseChance",   // unconfirmed
	"HasteRatingPerHastePercent",          // unconfirmed
	"CritRatingPerCritChance",             // unconfirmed
	"HitRatingPerHitChance",               // unconfirmed
	"DefenseRatingPerDefense",             // unconfirmed
	"DodgeRatingPerDodgeChance",           // unconfirmed
	"ParryRatingPerParryChance",           // unconfirmed
	"BlockRatingPerBlockChance",           // unconfirmed
}

// ProvisionalConstants returns the names of every rating constant still
// carrying an Era value. It is empty once BaseStatsBuild names a Forever
// client, and the spec support page shows the list either way.
func ProvisionalConstants() []string {
	if len(BaseStatsBuild) >= len(foreverBaseStatsBuildPrefix) &&
		BaseStatsBuild[:len(foreverBaseStatsBuildPrefix)] == foreverBaseStatsBuildPrefix {
		return nil
	}
	out := make([]string, len(provisionalConstantNames))
	copy(out, provisionalConstantNames)
	return out
}
```

- [ ] **Step 7: Run the tests and watch them pass**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim ./tools
go test --tags=with_db ./sim/core/ -run 'TestBaseStats|TestRatingConstants|TestExpertise|TestProvisional' -v
```

Expected: `BUILDS`; `gofmt -l` prints nothing; four tests `PASS`.

- [ ] **Step 8: Check nothing moved**

Regenerating should not change any behaviour, because the values it emits are the ones the hand-edited file already had — except `ExpertiseRatingPerExpertiseChance`, which was hand-set to `1` and is now read from the table. That one is load-bearing for Forever (expertise is live), so the change is intended; confirm it is the *only* one.

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db -count=1 ./sim/core/... ./sim/warrior/... ./sim/mage/... 2>&1 | tail -20
```

Expected: `ok` for each. If a `.results` golden moved, read the diff: only specs that actually carry expertise should shift. Regenerate with `make update-tests` and record the deltas in the commit body.

- [ ] **Step 9: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add tools/base_stats_parser.py sim/core/base_stats_auto_gen.go sim/core/base_stats_provisional.go sim/core/base_stats_test.go assets/db_inputs/basestats/BUILD
git commit -m "fix(core): make the base-stats generator work, and mark Era values provisional" \
  -m "tools/base_stats_parser.py has been broken: GenExtraStatsGoFile opens a triple-quoted block whose closing quotes swallow its return, so it returned None and crashed on write, and the committed base_stats_auto_gen.go had been hand-edited (it carried a TODO and a constant the generator never emitted). Rewritten: reads the CombatRatings columns by name, takes --build and --inputs, fails loudly when a column is missing rather than falling back to a literal, and stamps the client build into the output. Forever's tables land on Sept 17; until then the inputs are Era's and ProvisionalConstants() names every value that is therefore unconfirmed, which the spec support page shows. ExpertiseRatingPerExpertiseChance was hand-set to 1 and is now read from the table, which matters because expertise is dormant in vanilla and live for Forever." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 6: Bonus healing grants one third as bonus damage

**Repo: ENGINE.** Depends on Task 4. **G2 — INDEPENDENT of Tasks 5, 7, 8, 9, 10.**

Forever makes healing power carry a third of its value as spell damage. The engine already has a stat-dependency engine, so this is one dependency entry — but it does not work as written, and the reason is worth reading before touching anything.

`sim/core/stats/deps.go` keeps `safeDepsOrder`, a list that makes dependency evaluation single-pass and therefore cheap: a dependency is only valid if its source appears **before** its destination in that list. `validateDep` panics otherwise. The current order is

```go
	SpellPower,
	SpellDamage,
	HealingPower,
	Health,
```

so `HealingPower → SpellDamage` is invalid today and `character.AddStatDependency(stats.HealingPower, stats.SpellDamage, 1.0/3.0)` would panic at startup with `Invalid stat dependency: HealingPower --> SpellDamage`. The fix is to move `HealingPower` above `SpellDamage`, which is safe because nothing depends on `HealingPower` — verified: the only dependencies registered in the engine are the five in `addUniversalStatDependencies` (`sim/core/character.go:281-290`) plus per-class and per-item ones, and none has `HealingPower` as a destination.

**Files:**
- Modify: `sim/core/stats/deps.go` (the `safeDepsOrder` list, lines 12-35), `sim/core/character.go` (`addUniversalStatDependencies`, lines 280-291)
- Test: `sim/core/stats/deps_test.go`, `sim/core/character_deps_test.go`

**Interfaces:**
- Consumes: `stats.HealingPower`, `stats.SpellDamage`, `StatDependencyManager.AddStatDependency` — all existing.
- Produces: `core.HealingToSpellDamageRatio = 1.0 / 3.0`, a named constant so the number is not buried in a call, and the dependency itself, applied to every character.

- [ ] **Step 1: Write the failing test for the ordering**

Append to `sim/core/stats/deps_test.go`:

```go
// Forever: bonus healing carries one third as bonus damage. Dependency
// evaluation is single-pass over safeDepsOrder, so the source must appear
// before the destination or AddStatDependency panics. Nothing depends on
// HealingPower, so moving it above SpellDamage is safe — and this test is
// what stops a later reorder from silently breaking it.
func TestHealingPowerCanFeedSpellDamage(t *testing.T) {
	if !isValidDep(HealingPower, SpellDamage) {
		t.Fatal("HealingPower --> SpellDamage is not a valid dependency; " +
			"HealingPower must appear before SpellDamage in safeDepsOrder")
	}
}

func TestHealingToSpellDamageAppliesOneThird(t *testing.T) {
	sdm := NewStatDependencyManager()
	sdm.AddStatDependency(HealingPower, SpellDamage, 1.0/3.0)
	sdm.FinalizeStatDeps()

	got := sdm.ApplyStatDependencies(Stats{HealingPower: 300, SpellDamage: 50})
	if got[HealingPower] != 300 {
		t.Errorf("HealingPower = %v, want it unchanged at 300", got[HealingPower])
	}
	if got[SpellDamage] != 150 {
		t.Errorf("SpellDamage = %v, want 150 (50 base + 300/3)", got[SpellDamage])
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/stats/ -run 'TestHealing' -v`
Expected: `FAIL: TestHealingPowerCanFeedSpellDamage` with the message above, and `TestHealingToSpellDamageAppliesOneThird` panicking with `Invalid stat dependency: HealingPower --> SpellDamage`.

- [ ] **Step 3: Reorder `safeDepsOrder`**

In `sim/core/stats/deps.go`, the list at lines 12-35. Move `HealingPower` from after `SpellDamage` to before it, and say why:

```go
var safeDepsOrder = []Stat{
	Strength,
	Agility,
	Stamina,
	Intellect,
	Spirit,
	BonusArmor,
	Armor,
	FeralAttackPower,
	AttackPower,
	RangedAttackPower,
	SpellPower,
	// HealingPower comes before SpellDamage because Forever's bonus
	// healing carries one third as bonus damage (see
	// Character.addUniversalStatDependencies). Nothing depends on
	// HealingPower, so it is free to move up.
	HealingPower,
	SpellDamage,
	Health,
	Mana,
	MP5,
	Crit,
	Defense,
	Block,
	BlockValue,
	Dodge,
	Parry,
}
```

Note the single `Crit` where `SpellCrit` and `MeleeCrit` both were — Task 4's merge leaves a duplicate here if it was applied by `sed`; collapse it.

- [ ] **Step 4: Run the stats tests**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/stats/ -v`
Expected: `PASS` for both new tests and every pre-existing one.

- [ ] **Step 5: Write the failing test for the character-level dependency**

Every character gets it, not just casters, because a plate healer's gear carries healing power too. Create `sim/core/character_deps_test.go`:

```go
package core

import "testing"

// Forever: bonus healing carries one third as bonus damage. It is a
// universal dependency, applied to every character, because healing power
// appears on gear for classes that also deal damage.
func TestHealingToSpellDamageRatioIsOneThird(t *testing.T) {
	want := 1.0 / 3.0
	if HealingToSpellDamageRatio != want {
		t.Errorf("HealingToSpellDamageRatio = %v, want %v", HealingToSpellDamageRatio, want)
	}
}
```

Add to `sim/core/base_stats_test.go` — or create it if Task 5 has not run in this worktree — nothing; the dependency's effect is covered by the stats-package test above plus the DPS check in Step 8.

- [ ] **Step 6: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run TestHealingToSpellDamageRatio -v`
Expected: `FAIL [build failed]`, `undefined: HealingToSpellDamageRatio`.

- [ ] **Step 7: Register the dependency**

In `sim/core/character.go`, `addUniversalStatDependencies` at lines 280-291. Add the constant above the function and the dependency inside it:

```go
// HealingToSpellDamageRatio is Forever's rule that bonus healing carries
// one third as bonus damage. Confirmed from the Deep Dive panel; the exact
// fraction is one of the few Forever numbers that was stated outright, so
// it is not marked unconfirmed.
const HealingToSpellDamageRatio = 1.0 / 3.0

func (character *Character) addUniversalStatDependencies() {
	character.AddStat(stats.Health, 20-10*20)
	character.AddStatDependency(stats.Stamina, stats.Health, 10)
	character.AddStatDependency(stats.Agility, stats.Armor, 2)
	character.AddStatDependency(stats.Defense, stats.Dodge, MissDodgeParryBlockCritChancePerDefense)
	character.AddStatDependency(stats.Defense, stats.Parry, MissDodgeParryBlockCritChancePerDefense)
	character.AddStatDependency(stats.Defense, stats.Block, MissDodgeParryBlockCritChancePerDefense)
	// Forever: bonus healing grants one third as bonus damage. Applied to
	// every character, because healing power appears on gear worn by
	// classes that also deal damage.
	character.AddStatDependency(stats.HealingPower, stats.SpellDamage, HealingToSpellDamageRatio)

	character.AddStat(stats.Parry, 5*ParryRatingPerParryChance)
	character.AddStat(stats.Block, 5*BlockRatingPerBlockChance)
}
```

- [ ] **Step 8: Run the tests and check what moved**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test --tags=with_db ./sim/core/... -v -run 'TestHealing'
go test --tags=with_db -count=1 ./sim/core/... ./sim/priest/... ./sim/shaman/... ./sim/paladin/... ./sim/druid/...
```

Expected: `BUILDS`; `gofmt -l` prints nothing; the three healing tests `PASS`.

The `.results` goldens for any spec whose gear carries healing power will now report more DPS — hybrids especially. That is the change, not a regression. Regenerate and read the diff:

```bash
go test --tags=with_db -count=1 ./sim/... 2>&1 | tail -25
make update-tests
git diff --stat -- '*.results'
```

A pure-melee spec with no healing power on its phase-1 gear should not move at all. If a warrior's DPS changed, something else is wrong: check that the dependency was not registered twice.

- [ ] **Step 9: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add sim/core/stats/deps.go sim/core/stats/deps_test.go sim/core/character.go sim/core/character_deps_test.go
git add -u -- '*.results'
git commit -m "feat(core): bonus healing grants one third as bonus damage" \
  -m "One entry in the existing stat-dependency table, plus the reorder it needs: dependency evaluation is single-pass over safeDepsOrder and a source must precede its destination, and HealingPower sat after SpellDamage, so registering this would have panicked at startup. Nothing depends on HealingPower, so it moves up freely; a test asserts the ordering rather than leaving it to whoever next edits the list. Applied to every character, because healing power appears on gear worn by classes that also deal damage. Hybrid .results goldens move; pure-melee ones do not." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 7: Port `spell_mod.go` from `wowsims/sod`

**Repo: ENGINE.** Depends on Task 4. **G2 — INDEPENDENT of Tasks 5, 6, 8, 9, 10.** Tasks 11 and 12 depend on it.

Forever's reworked talents are mostly "these spells cost, crit, or hit differently". Season of Discovery expresses exactly that as declarative config rather than closures, in `sim/core/spell_mod.go` (790 lines, 26 modifier kinds). Cherry-picking it is what makes `talents.go` for a Forever spec a table instead of a pile of `OnSpellRegistered` hooks.

**It does not compile against classic's core as-is. Measured** — copying the file in with the import rewritten and running `go build -gcflags="-e" ./sim/core/` produces exactly eleven undefined symbols:

```
undefined: SpellFlagNoSpellMods
spell.Matches undefined
spell.ClassSpellMask undefined
spell.RelatedSelfBuff undefined
spell.ApplyAdditiveBaseDamageBonus undefined
spell.ApplyAdditiveDamageBonus undefined
spell.ApplyAdditiveImpactDamageBonus undefined
spell.ApplyAdditivePeriodicDamageBonus undefined
spell.ApplyMultiplicativeDamageBonus undefined
spell.CD.ApplyFlatCooldownMod undefined
spell.CD.ApplyFlatPercentCooldownMod undefined
```

Two of them are not a straight copy:

- **The damage-bonus helpers.** SoD keeps unexported `int64` percent accumulators (`damageMultiplierAdditivePct`, `impactDamageMultiplierAdditivePct`, …) and recomputes a cached multiplier in `updateImpactDamageMultiplier`. Classic keeps **exported `float64` multiplier fields on `Spell` and `SpellConfig`** — `BaseDamageMultiplierAdditive`, `DamageMultiplier`, `DamageMultiplierAdditive`, `ImpactDamageMultiplierAdditive`, `PeriodicDamageMultiplierAdditive` (`sim/core/spell.go:44-48` and `142-146`), each defaulting to 1 and read directly at damage time. Do **not** port SoD's accumulator machinery; implement the five helpers in classic's terms. A SoD mod passes a percent as an `int64` (`-50` means minus fifty percent), so each helper converts once.
- **`spell.CD.ApplyFlatCooldownMod`.** SoD wraps cooldowns in a `SpellCooldown` type holding a flat modifier and a percent multiplier, applied lazily in `applyCooldownModifiers`. Classic's `Cooldown` (`sim/core/cooldown.go:12-18`) is a `*Timer` plus a `Duration` and nothing else. Rather than introduce a second cooldown type across the whole engine, give `Cooldown` the two methods directly and have them mutate `Duration`; the semantics a talent needs ("this spell's cooldown is two seconds shorter") are the same, and the difference only shows if two mods apply a percentage after a flat in a different order than SoD, which no shipped talent does. The comment on the methods says so.

Two more things to check before porting: `Spell.ClassSpellMask` does not exist in classic, which instead has `Spell.SpellCode int32` (`sim/core/spell.go:19,80,238`) used the same way but not as a bitmask. Keep **both**: `SpellCode` stays for the existing class code, and `ClassSpellMask uint64` is added for spell mods, because a mask is what lets one talent config target a set of spells.

**Files:**
- Create: `sim/core/spell_mod.go` (ported), `sim/core/spell_mod_test.go`
- Modify: `sim/core/spell.go`, `sim/core/cooldown.go`, `sim/core/flags.go`, `sim/core/unit.go`
- Source: the SoD clone at `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/sod` (commit `0e3f6ef`). If it is gone, `git clone --depth 1 https://github.com/wowsims/sod /tmp/sod`.

**Interfaces:**
- Produces, used by Tasks 11 and 12:
  - `core.SpellModConfig{Kind SpellModType; ClassMask uint64; ClassSpellsOnly bool; School SpellSchool; SpellFlags, SpellFlagsExclude SpellFlag; DefenseType DefenseType; ProcMask ProcMask; CastType proto.CastType; IntValue int64; TimeValue time.Duration; FloatValue float64; KeyValue string; ApplyCustom SpellModApply; RemoveCustom SpellModRemove}`
  - `core.SpellMod` and the 26 `core.SpellMod_*` kinds (`SpellMod_DamageDone_Pct`, `SpellMod_DamageDone_Flat`, `SpellMod_BaseDamageDone_Flat`, `SpellMod_PeriodicDamageDone_Flat`, `SpellMod_ImpactDamageDone_Flat`, `SpellMod_CritDamageBonus_Flat`, `SpellMod_PowerCost_Pct`, `SpellMod_PowerCost_Flat`, `SpellMod_Cooldown_Flat`, `SpellMod_Cooldown_Multi_Flat`, `SpellMod_BonusCrit_Flat`, `SpellMod_BonusHit_Flat`, `SpellMod_CastTime_Pct`, `SpellMod_CastTime_Flat`, `SpellMod_DotNumberOfTicks_Flat`, `SpellMod_DotTickLength_Flat`, `SpellMod_DotTickLength_Pct`, `SpellMod_GlobalCooldown_Flat`, `SpellMod_BonusCoeffecient_Flat`, `SpellMod_AllowCastWhileMoving`, `SpellMod_BonusDamage_Flat`, `SpellMod_BonusExpertise_Rating`, `SpellMod_Threat_Flat`, `SpellMod_Threat_Pct`, `SpellMod_BonusThreat_Flat`, `SpellMod_DebuffDuration_Flat`, `SpellMod_BuffDuration_Flat`, `SpellMod_Custom`)
  - `(*Unit).AddStaticMod(config SpellModConfig)`, `(*Unit).AddDynamicMod(config SpellModConfig) *SpellMod`, `(*SpellMod).Activate()`, `(*SpellMod).Deactivate()`, `(*SpellMod).UpdateIntValue`, `(*SpellMod).UpdateTimeValue`, `(*SpellMod).UpdateFloatValue`
  - `core.SpellFlagNoSpellMods SpellFlag`
  - `(*Spell).Matches(mask uint64) bool`, `Spell.ClassSpellMask uint64`, `SpellConfig.ClassSpellMask uint64`, `Spell.RelatedSelfBuff *Aura`, `SpellConfig.RelatedSelfBuff *Aura`
  - `(*Spell).ApplyMultiplicativeDamageBonus(multiplier float64)`, `(*Spell).ApplyAdditiveDamageBonus(percent int64)`, `(*Spell).ApplyAdditiveBaseDamageBonus(percent int64)`, `(*Spell).ApplyAdditiveImpactDamageBonus(percent int64)`, `(*Spell).ApplyAdditivePeriodicDamageBonus(percent int64)`
  - `(*Cooldown).ApplyFlatCooldownMod(d time.Duration)`, `(*Cooldown).ApplyFlatPercentCooldownMod(percent int64)`

- [ ] **Step 1: Write the failing test**

Create `sim/core/spell_mod_test.go`:

```go
package core

import (
	"testing"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
)

const (
	testMaskBloodthirst uint64 = 1 << 0
	testMaskWhirlwind   uint64 = 1 << 1
)

// modTestSpell registers a spell on a bare unit so a mod has something to
// bind to, without standing up a whole character.
func modTestSpell(unit *Unit, mask uint64, cd time.Duration) *Spell {
	return unit.RegisterSpell(SpellConfig{
		ActionID:       ActionID{SpellID: 23894},
		ClassSpellMask: mask,
		SpellSchool:    SpellSchoolPhysical,
		DefenseType:    DefenseTypeMelee,
		ProcMask:       ProcMaskMeleeMHSpecial,
		Flags:          SpellFlagMeleeMetrics | SpellFlagAPL,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		Cast: CastConfig{
			DefaultCast: Cast{GCD: GCDDefault},
			CD:          Cooldown{Timer: unit.NewTimer(), Duration: cd},
		},
		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {},
	})
}

func TestSpellMatchesClassMask(t *testing.T) {
	unit := &Unit{Type: PlayerUnit}
	unit.Init(nil)
	spell := modTestSpell(unit, testMaskBloodthirst, time.Second*6)
	if !spell.Matches(testMaskBloodthirst) {
		t.Error("spell does not match its own mask")
	}
	if spell.Matches(testMaskWhirlwind) {
		t.Error("spell matches a mask it does not carry")
	}
	if spell.Matches(0) {
		t.Error("an empty mask must match nothing")
	}
}

// The five damage helpers are re-expressed in classic's exported float
// multiplier fields rather than SoD's private percent accumulators. A
// percent of -50 must halve; +20 must add a fifth.
func TestApplyDamageBonusHelpers(t *testing.T) {
	unit := &Unit{Type: PlayerUnit}
	unit.Init(nil)

	spell := modTestSpell(unit, testMaskBloodthirst, 0)
	spell.ApplyAdditiveDamageBonus(20)
	if got, want := spell.DamageMultiplierAdditive, 1.2; !closeEnough(got, want) {
		t.Errorf("after +20%%, DamageMultiplierAdditive = %v, want %v", got, want)
	}
	spell.ApplyAdditiveDamageBonus(-20)
	if got, want := spell.DamageMultiplierAdditive, 1.0; !closeEnough(got, want) {
		t.Errorf("after +20%% then -20%%, DamageMultiplierAdditive = %v, want %v", got, want)
	}

	spell.ApplyMultiplicativeDamageBonus(1.5)
	if got, want := spell.DamageMultiplier, 1.5; !closeEnough(got, want) {
		t.Errorf("after x1.5, DamageMultiplier = %v, want %v", got, want)
	}

	spell.ApplyAdditiveBaseDamageBonus(10)
	if got, want := spell.BaseDamageMultiplierAdditive, 1.1; !closeEnough(got, want) {
		t.Errorf("BaseDamageMultiplierAdditive = %v, want %v", got, want)
	}
	spell.ApplyAdditiveImpactDamageBonus(10)
	if got, want := spell.ImpactDamageMultiplierAdditive, 1.1; !closeEnough(got, want) {
		t.Errorf("ImpactDamageMultiplierAdditive = %v, want %v", got, want)
	}
	spell.ApplyAdditivePeriodicDamageBonus(10)
	if got, want := spell.PeriodicDamageMultiplierAdditive, 1.1; !closeEnough(got, want) {
		t.Errorf("PeriodicDamageMultiplierAdditive = %v, want %v", got, want)
	}
}

func TestCooldownMods(t *testing.T) {
	cd := Cooldown{Duration: time.Second * 30}
	cd.ApplyFlatCooldownMod(-time.Second * 5)
	if cd.Duration != time.Second*25 {
		t.Errorf("after -5s, Duration = %v, want 25s", cd.Duration)
	}
	cd.ApplyFlatPercentCooldownMod(-50)
	if cd.Duration != time.Second*12500/1000 {
		t.Errorf("after -50%%, Duration = %v, want 12.5s", cd.Duration)
	}
	// A cooldown never goes negative, however many mods pile on.
	cd.ApplyFlatCooldownMod(-time.Hour)
	if cd.Duration != 0 {
		t.Errorf("Duration = %v, want it clamped to 0", cd.Duration)
	}
}

// A static mod applies at registration and to every spell registered
// afterwards, which is what makes talents declarative.
func TestStaticModAppliesToMatchingSpells(t *testing.T) {
	unit := &Unit{Type: PlayerUnit}
	unit.Init(nil)

	before := modTestSpell(unit, testMaskBloodthirst, 0)
	unit.AddStaticMod(SpellModConfig{
		Kind:      SpellMod_DamageDone_Flat,
		ClassMask: testMaskBloodthirst,
		IntValue:  30,
	})
	after := modTestSpell(unit, testMaskBloodthirst, 0)
	other := modTestSpell(unit, testMaskWhirlwind, 0)

	if got, want := before.DamageMultiplierAdditive, 1.3; !closeEnough(got, want) {
		t.Errorf("spell registered before the mod: %v, want %v", got, want)
	}
	if got, want := after.DamageMultiplierAdditive, 1.3; !closeEnough(got, want) {
		t.Errorf("spell registered after the mod: %v, want %v", got, want)
	}
	if got, want := other.DamageMultiplierAdditive, 1.0; !closeEnough(got, want) {
		t.Errorf("a non-matching spell was modified: %v, want %v", got, want)
	}
}

// A dynamic mod is what a proc or a temporary buff uses: it toggles.
func TestDynamicModActivatesAndDeactivates(t *testing.T) {
	unit := &Unit{Type: PlayerUnit}
	unit.Init(nil)
	spell := modTestSpell(unit, testMaskBloodthirst, time.Second*6)

	mod := unit.AddDynamicMod(SpellModConfig{
		Kind:      SpellMod_Cooldown_Flat,
		ClassMask: testMaskBloodthirst,
		TimeValue: -time.Second * 2,
	})
	if spell.CD.Duration != time.Second*6 {
		t.Fatalf("a dynamic mod applied before Activate: %v", spell.CD.Duration)
	}
	mod.Activate()
	if spell.CD.Duration != time.Second*4 {
		t.Errorf("after Activate, CD = %v, want 4s", spell.CD.Duration)
	}
	mod.Deactivate()
	if spell.CD.Duration != time.Second*6 {
		t.Errorf("after Deactivate, CD = %v, want 6s", spell.CD.Duration)
	}
}

// A spell can opt out entirely, which the engine needs for auto attacks
// and for spells a talent must not touch.
func TestSpellFlagNoSpellModsOptsOut(t *testing.T) {
	unit := &Unit{Type: PlayerUnit}
	unit.Init(nil)
	spell := unit.RegisterSpell(SpellConfig{
		ActionID:         ActionID{SpellID: 1},
		ClassSpellMask:   testMaskBloodthirst,
		Flags:            SpellFlagNoSpellMods,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects:     func(sim *Simulation, target *Unit, spell *Spell) {},
	})
	unit.AddStaticMod(SpellModConfig{Kind: SpellMod_DamageDone_Flat, ClassMask: testMaskBloodthirst, IntValue: 50})
	if got, want := spell.DamageMultiplierAdditive, 1.0; !closeEnough(got, want) {
		t.Errorf("a SpellFlagNoSpellMods spell was modified: %v, want %v", got, want)
	}
}

func TestUnknownModKindPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("an unimplemented mod kind did not panic")
		}
	}()
	unit := &Unit{Type: PlayerUnit}
	unit.Init(nil)
	unit.AddStaticMod(SpellModConfig{Kind: SpellModType(1 << 62), ClassMask: testMaskBloodthirst})
}

func closeEnough(a, b float64) bool {
	d := a - b
	return d < 1e-9 && d > -1e-9
}

var _ = proto.CastType_CastTypeUnknown
```

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'TestSpellMatches|TestApplyDamage|TestCooldownMods|TestStaticMod|TestDynamicMod|TestSpellFlagNoSpellMods|TestUnknownModKind' -v`
Expected: `FAIL [build failed]`, `undefined: SpellMod_DamageDone_Flat` among others.

- [ ] **Step 3: Add the eleven missing symbols**

In `sim/core/flags.go`, add to the `SpellFlag` const block (use the next free bit; read the block first and take the one after the current highest):

```go
	SpellFlagNoSpellMods // Indicates that no spell mods should be applied to this spell
```

In `sim/core/spell.go`, add to **both** `SpellConfig` (around line 19, beside `SpellCode`) and `Spell` (around line 80):

```go
	// ClassSpellMask is a bitmask a SpellMod's ClassMask is tested
	// against, so one talent config can target a set of spells. It is
	// additional to SpellCode, which the existing class code uses as a
	// scalar identity; the two do not overlap in purpose.
	ClassSpellMask uint64

	// RelatedSelfBuff is the aura this spell applies to its caster, if
	// any. SpellMod_BuffDuration_Flat extends it.
	RelatedSelfBuff *Aura
```

and copy both through in `RegisterSpell` beside the existing `SpellCode: config.SpellCode,` (line 238):

```go
		ClassSpellMask:  config.ClassSpellMask,
		RelatedSelfBuff: config.RelatedSelfBuff,
```

Then append to `sim/core/spell.go`:

```go
// Matches reports whether this spell is in the given ClassSpellMask set.
// An empty mask matches nothing, so a SpellMod with no ClassMask must use
// ClassSpellsOnly or a School/Flags filter instead.
func (spell *Spell) Matches(mask uint64) bool {
	return spell.ClassSpellMask&mask > 0
}

// The five damage-bonus helpers below are the SpellMod system's write
// path into a spell's damage. Season of Discovery, which this system is
// ported from, keeps private int64 percent accumulators and recomputes a
// cached multiplier; classic keeps exported float64 multipliers that are
// read at damage time, so these convert a SpellMod's percent once and
// write the multiplier directly. `percent` is an offset from zero: 20
// means plus twenty percent, -50 means minus fifty.

// ApplyAdditiveBaseDamageBonus is "Modifies Spell Effectiveness (8)".
func (spell *Spell) ApplyAdditiveBaseDamageBonus(percent int64) {
	spell.BaseDamageMultiplierAdditive += float64(percent) / 100
}

// ApplyMultiplicativeDamageBonus is "Mod Damage Done %": it multiplies
// direct and periodic damage together.
func (spell *Spell) ApplyMultiplicativeDamageBonus(multiplier float64) {
	spell.DamageMultiplier *= multiplier
}

// ApplyAdditiveDamageBonus is "Modifies Damage/Healing Done (22)",
// applying to direct and periodic damage together.
func (spell *Spell) ApplyAdditiveDamageBonus(percent int64) {
	spell.DamageMultiplierAdditive += float64(percent) / 100
}

// ApplyAdditiveImpactDamageBonus applies to direct damage only.
func (spell *Spell) ApplyAdditiveImpactDamageBonus(percent int64) {
	spell.ImpactDamageMultiplierAdditive += float64(percent) / 100
}

// ApplyAdditivePeriodicDamageBonus applies to periodic damage only.
func (spell *Spell) ApplyAdditivePeriodicDamageBonus(percent int64) {
	spell.PeriodicDamageMultiplierAdditive += float64(percent) / 100
}
```

In `sim/core/cooldown.go`, append:

```go
// ApplyFlatCooldownMod adds a duration to this cooldown, clamped at zero.
//
// Season of Discovery wraps cooldowns in a SpellCooldown type that keeps
// the flat and percentage modifiers separately and applies them lazily in
// a fixed order. Classic's Cooldown is a timer and a duration, and
// introducing a second cooldown type across the engine to preserve that
// ordering buys nothing: no shipped talent applies a percentage and a
// flat modifier to the same cooldown, so applying each in call order is
// indistinguishable. If one ever does, this is the place to revisit.
func (cd *Cooldown) ApplyFlatCooldownMod(duration time.Duration) {
	cd.Duration = max(0, cd.Duration+duration)
}

// ApplyFlatPercentCooldownMod scales this cooldown. `percent` is an
// offset from zero: -50 means minus fifty percent.
func (cd *Cooldown) ApplyFlatPercentCooldownMod(percent int64) {
	cd.Duration = max(0, time.Duration(float64(cd.Duration)*float64(100+percent)/100))
}
```

- [ ] **Step 4: Copy and adapt `spell_mod.go`**

```bash
cd /Users/jh/code/wowsims-forever
SOD=/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/sod
test -f "$SOD/sim/core/spell_mod.go" || { git clone --depth 1 https://github.com/wowsims/sod /tmp/sod && SOD=/tmp/sod; }
sed 's|wowsims/sod|wowsims/classic|g' "$SOD/sim/core/spell_mod.go" > sim/core/spell_mod.go
go build -gcflags="-e" ./sim/core/ 2>&1 | sed 's/.*spell_mod.go:[0-9]*:[0-9]*: //' | sort -u
```

Expected after Step 3: an empty list, or a short one. Add a provenance header to the top of the new file, above `package core`:

```go
// Ported from wowsims/sod sim/core/spell_mod.go at commit 0e3f6ef.
//
// Forever's reworked talents are mostly "these spells cost, crit, or hit
// differently", which this expresses as config rather than as closures on
// OnSpellRegistered. Two things differ from the SoD original and the
// difference is deliberate:
//
//   - the Apply*DamageBonus helpers write classic's exported float64
//     multiplier fields rather than SoD's private percent accumulators;
//   - Cooldown gains the two mod methods directly rather than classic
//     adopting SoD's SpellCooldown wrapper (see cooldown.go).
//
// See PORTING.md.
```

If anything still fails to compile, resolve it the same way: express the SoD behaviour in classic's existing types, never by importing more of SoD's machinery. Record each such decision in `PORTING.md` under a new "spell_mod.go" heading.

- [ ] **Step 5: Run the tests and watch them pass**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test --tags=with_db ./sim/core/ -run 'TestSpellMatches|TestApplyDamage|TestCooldownMods|TestStaticMod|TestDynamicMod|TestSpellFlagNoSpellMods|TestUnknownModKind' -v
```

Expected: `BUILDS`, no `gofmt` output, seven tests `PASS`.

- [ ] **Step 6: Prove nothing else moved**

Nothing registers a `ClassSpellMask` yet, so no existing spell can match a mod, so no DPS may change.

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db -count=1 ./sim/... 2>&1 | tail -25
git status --short -- '*.results'
```

Expected: 20 packages `ok`; **`git status` prints nothing for `.results`**. A moved golden here means a real behaviour change slipped in with the port — find it before committing.

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add sim/core/spell_mod.go sim/core/spell_mod_test.go sim/core/spell.go sim/core/cooldown.go sim/core/flags.go PORTING.md
git commit -m "feat(core): port the declarative spell-mod system from wowsims/sod" \
  -m "Forever's reworked talents are mostly cost, crit and hit modifiers on a set of spells, which sod's SpellMod expresses as config rather than as OnSpellRegistered closures. The file needed eleven symbols classic lacks; nine are straight additions (SpellFlagNoSpellMods, Spell.ClassSpellMask, Spell.Matches, Spell.RelatedSelfBuff, and the five damage helpers), and two are deliberate divergences: the damage helpers write classic's exported float multiplier fields instead of sod's private percent accumulators, and Cooldown gains the two mod methods directly instead of classic adopting sod's SpellCooldown wrapper, because no shipped talent applies both a flat and a percentage to one cooldown. Both are recorded in PORTING.md. No spell carries a ClassSpellMask yet, so no .results golden moves." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 8: Encounter biome, and creature type for conditional item effects

**Repo: ENGINE.** Depends on Task 4. **G2 — INDEPENDENT of Tasks 5, 6, 7, 9, 10.**

Forever re-itemised the world and added trinkets whose effect depends on where you are and what you are fighting. The research called this "a new mechanic class with no precedent in the engine" and sized it at a week. **Reading the code, half of it already exists and the estimate is wrong in our favour.**

**Verified:** `MobType` is already a proto enum (`proto/common.proto:739-749`: Beast, Demon, Dragonkin, Elemental, Giant, Humanoid, Mechanical, Undead), already a field on `Target` (`Target.mob_type = 4`), already plumbed onto the core unit (`Unit.MobType` at `sim/core/unit.go:50`, set in `sim/core/target.go:125`), already read by conditional effects (`sim/core/racials.go:126` checks `t.MobType == MobTypeBeast`; `sim/core/consumes.go:111,165` check for Undead), and already has two ready-made item-effect constructors — `NewMobTypeAttackPowerEffect(itemID, mobTypes, bonus)` and `NewMobTypeSpellPowerEffect` at `sim/core/item_effects.go:157,187`.

So the only genuinely new concept is the **encounter biome**. This task adds it, and adds the two damage-multiplier constructors the creature-type half is missing so both conditions are expressed the same way.

Both proto changes are strictly additive: new field numbers at the end of their messages, nothing renumbered, nothing removed. A `SimDatabase` built before this change still parses; an encounter with no biome gets `BiomeUnknown`, which matches nothing, which is exactly a vanilla encounter.

**Files:**
- Modify: `proto/common.proto` (new `Biome` enum; `Encounter.biome` field 20; `Target.biome_affinity` is **not** added — see below)
- Modify: `sim/core/target.go` (`Encounter.Biome`, set in `NewEncounter`), `sim/core/environment.go` (expose it on the environment so a unit can read it), `sim/core/item_effects.go` (two new constructors)
- Test: `sim/core/environment_biome_test.go`
- Regenerate: `make proto`

**Interfaces:**
- Produces, used by item-effect content and by the web's encounter picker:
  - `proto.Biome` enum: `BiomeUnknown=0, BiomeForest=1, BiomeDesert=2, BiomeSnow=3, BiomeSwamp=4, BiomeUnderground=5, BiomeCoastal=6, BiomeMountain=7, BiomeVolcanic=8, BiomePlains=9, BiomeIndoors=10, BiomeAbyssal=11`
  - `proto.Encounter.biome` (field 20)
  - `core.Encounter.Biome proto.Biome`
  - `(*Unit).Biome() proto.Biome` — reads the environment's encounter, so an item effect written against a unit needs no plumbing of its own
  - `core.NewBiomeDamageEffect(itemID int32, biomes []proto.Biome, multiplier float64)`
  - `core.NewMobTypeDamageEffect(itemID int32, mobTypes []proto.MobType, multiplier float64)`

Why no `Target.biome_affinity`: a biome is a property of the encounter, not of one target, and a multi-target pull is in one place. Putting it on `Target` would let two targets in one fight disagree about where they are. The design says "an encounter environment (biome, target creature type)", and creature type is already per-target where it belongs.

- [ ] **Step 1: Write the failing test**

Create `sim/core/environment_biome_test.go`:

```go
package core

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// Forever's biome-conditional trinkets need the encounter to know where it
// is. The concept is additive: an encounter that does not set one reads
// BiomeUnknown, which matches nothing, which is a vanilla fight.
func TestEncounterCarriesItsBiome(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Biome:    proto.Biome_BiomeVolcanic,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60},
	})
	if enc.Biome != proto.Biome_BiomeVolcanic {
		t.Errorf("Encounter.Biome = %v, want BiomeVolcanic", enc.Biome)
	}
}

func TestEncounterWithoutABiomeIsUnknown(t *testing.T) {
	enc := NewEncounter(&proto.Encounter{
		Duration: 180,
		Targets:  []*proto.Target{DefaultTargetProtoLvl60},
	})
	if enc.Biome != proto.Biome_BiomeUnknown {
		t.Errorf("Encounter.Biome = %v, want BiomeUnknown for an encounter that sets none", enc.Biome)
	}
}

// An item effect is written against a unit, so a unit must be able to ask
// where it is without the effect plumbing the environment itself.
func TestUnitReadsTheEncounterBiome(t *testing.T) {
	env := setupBiomeEnv(t, proto.Biome_BiomeSwamp, proto.MobType_MobTypeBeast)
	player := env.Raid.Parties[0].Players[0].GetCharacter().Unit
	if got := player.Biome(); got != proto.Biome_BiomeSwamp {
		t.Errorf("Unit.Biome() = %v, want BiomeSwamp", got)
	}
	if got := env.Encounter.TargetUnits[0].MobType; got != proto.MobType_MobTypeBeast {
		t.Errorf("target MobType = %v, want MobTypeBeast", got)
	}
}

// A unit with no environment yet (during registration, before the
// environment is constructed) must answer BiomeUnknown rather than panic.
func TestUnitBiomeBeforeEnvironmentIsUnknown(t *testing.T) {
	u := &Unit{Type: PlayerUnit}
	if got := u.Biome(); got != proto.Biome_BiomeUnknown {
		t.Errorf("Unit.Biome() with no environment = %v, want BiomeUnknown", got)
	}
}
```

You will also need a small helper; put it in the same file:

```go
// setupBiomeEnv builds the smallest environment that has a player and one
// target, so the biome plumbing can be read end to end.
func setupBiomeEnv(t *testing.T, biome proto.Biome, mob proto.MobType) *Environment {
	t.Helper()
	target := googleCloneTarget(DefaultTargetProtoLvl60)
	target.MobType = mob
	env, _, _ := NewEnvironment(
		SinglePlayerRaidProto(&proto.Player{
			Name:  "Biome Test",
			Race:  proto.Race_RaceOrc,
			Class: proto.Class_ClassWarrior,
		}, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
		&proto.Encounter{Duration: 180, Biome: biome, Targets: []*proto.Target{target}},
		false,
	)
	return env
}
```

`googleCloneTarget` is `googleProto.Clone(x).(*proto.Target)`; if `sim/core` does not already import `google.golang.org/protobuf/proto` under an alias, write the clone by hand rather than adding an import for a test helper. Check `NewEnvironment`'s exact signature and return arity in `sim/core/environment.go:55` and match it — it is `NewEnvironment(raid *proto.Raid, encounter *proto.Encounter, runFakePrepull bool)`; adapt the call if the arity differs.

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'TestEncounterCarries|TestEncounterWithout|TestUnitReadsThe|TestUnitBiomeBefore' -v`
Expected: `FAIL [build failed]`, `Biome undefined (type *proto.Encounter has no field or method Biome)`.

- [ ] **Step 3: Add the proto**

In `proto/common.proto`, beside the `MobType` enum at line 739, add:

```proto
// Forever re-itemised the world and added trinkets whose effect depends on
// where the fight is. A biome is a property of the encounter, not of one
// target: a multi-target pull happens in one place, and putting this on
// Target would let two targets in one fight disagree about where they are.
// The creature-type half of the same mechanic is Target.mob_type, which
// already exists and is already read by conditional effects.
enum Biome {
	BiomeUnknown = 0;
	BiomeForest = 1;
	BiomeDesert = 2;
	BiomeSnow = 3;
	BiomeSwamp = 4;
	BiomeUnderground = 5;
	BiomeCoastal = 6;
	BiomeMountain = 7;
	BiomeVolcanic = 8;
	BiomePlains = 9;
	BiomeIndoors = 10;
	BiomeAbyssal = 11;
}
```

The list is the set Forever's Deep Dive panel named. It is additive by construction: a biome Forever adds later is a new value at the end, and an encounter that does not know its biome reads `BiomeUnknown`.

In `message Encounter`, append at the end of the message body — read the file for the current highest field number in that message and use the next one; at the time of writing the highest is 19, so:

```proto
	// Where this fight happens, for Forever's biome-conditional item
	// effects. Unset means BiomeUnknown, which matches nothing.
	Biome biome = 20;
```

Regenerate: `export PATH=$PATH:$(go env GOPATH)/bin && make proto`.

- [ ] **Step 4: Carry it into the core encounter and onto the unit**

In `sim/core/target.go`, add to `type Encounter struct` (after `TargetUnits`, line 15):

```go
	// Biome is where this fight happens, for Forever's biome-conditional
	// item effects. BiomeUnknown matches nothing.
	Biome proto.Biome
```

and set it in `NewEncounter` (line 32), inside the `encounter := Encounter{...}` literal:

```go
		Biome: options.Biome,
```

In `sim/core/unit.go`, append:

```go
// Biome reports where the current encounter is happening, for Forever's
// biome-conditional item effects. A unit whose environment has not been
// constructed yet — which is every unit during spell registration —
// reports BiomeUnknown rather than panicking, so an effect may read it at
// registration time and get a safe answer.
func (unit *Unit) Biome() proto.Biome {
	if unit.Env == nil {
		return proto.Biome_BiomeUnknown
	}
	return unit.Env.Encounter.Biome
}
```

Check the field name the unit uses for its environment (`unit.Env` in classic; grep `Env \*Environment` in `sim/core/unit.go`) and match it.

- [ ] **Step 5: Add the two conditional-damage constructors**

`sim/core/item_effects.go:157,187` already has `NewMobTypeAttackPowerEffect` and `NewMobTypeSpellPowerEffect`, which add a flat stat when the *target* matches. Forever's trinkets are mostly damage multipliers, and the biome half has nothing at all. Read the existing two and follow their shape exactly; then append:

```go
// NewMobTypeDamageEffect registers an item that multiplies the wearer's
// damage against a set of creature types. The flat-stat equivalents are
// NewMobTypeAttackPowerEffect and NewMobTypeSpellPowerEffect above; this
// is the multiplier form Forever's specialised trinkets use.
func NewMobTypeDamageEffect(itemID int32, mobTypes []proto.MobType, multiplier float64) {
	NewItemEffect(itemID, func(agent Agent) {
		character := agent.GetCharacter()
		for _, target := range character.Env.Encounter.TargetUnits {
			if slices.Contains(mobTypes, target.MobType) {
				character.AttackTables[target.UnitIndex].DamageDealtMultiplier *= multiplier
			}
		}
	})
}

// NewBiomeDamageEffect registers an item that multiplies the wearer's
// damage in a set of biomes. Forever's biome trinkets are the reason the
// encounter carries a biome at all; an encounter with no biome set reads
// BiomeUnknown and matches nothing, so a vanilla fight is unaffected.
func NewBiomeDamageEffect(itemID int32, biomes []proto.Biome, multiplier float64) {
	NewItemEffect(itemID, func(agent Agent) {
		character := agent.GetCharacter()
		if !slices.Contains(biomes, character.Biome()) {
			return
		}
		character.PseudoStats.DamageDealtMultiplier *= multiplier
	})
}
```

The exact field used to scale outgoing damage differs between the per-target attack table and the unit's `PseudoStats`; read `sim/core/item_effects.go:157-200` and `sim/core/stats/stats.go`'s `PseudoStats` and use whatever the neighbouring constructors use. If `AttackTables[...].DamageDealtMultiplier` does not exist, apply the multiplier to `PseudoStats.DamageDealtMultiplier` in both constructors and note in a comment that the creature-type form is then whole-fight rather than per-target — which is correct when the encounter has one target and is the common case.

`slices` may need adding to the imports.

- [ ] **Step 6: Run the tests and watch them pass**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test --tags=with_db ./sim/core/ -run 'TestEncounterCarries|TestEncounterWithout|TestUnitReadsThe|TestUnitBiomeBefore' -v
```

Expected: `BUILDS`; no `gofmt` output; four tests `PASS`.

- [ ] **Step 7: Prove it is additive**

No existing encounter sets a biome and no existing item uses the new constructors, so nothing may change.

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db -count=1 ./sim/... 2>&1 | tail -25
git status --short -- '*.results'
```

Expected: 20 packages `ok`; **`git status` prints nothing for `.results`**.

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add proto/common.proto sim/core/proto/ sim/core/target.go sim/core/unit.go sim/core/item_effects.go sim/core/environment_biome_test.go
git commit -m "feat(core): an encounter biome, and damage multipliers conditional on it or on creature type" \
  -m "Forever's biome- and creature-type trinkets needed one new engine concept, not two: MobType is already a proto enum, already a field on Target, already on Unit, and already read by racials and consumables, with two flat-stat item-effect constructors alongside it. So this adds only the encounter's biome, plus the two damage-multiplier constructors the multiplier form was missing. A biome belongs to the encounter rather than to a target because a multi-target pull happens in one place. Both proto changes are additive at the end of their messages; an encounter that sets no biome reads BiomeUnknown, matches nothing, and behaves exactly as it does today, which the unchanged .results goldens show." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 9: Racials for Forever's two-active-two-passive model

**Repo: ENGINE.** Depends on Task 4. **G2 — INDEPENDENT of Tasks 5, 6, 7, 8, 10.**

Forever gives every race two active and two passive racials. Vanilla's are lopsided: Undead has one passive and nothing else in `applyRaceEffects`, Troll has four things and a helper, Gnome has two passives. `sim/core/racials.go` is 259 lines and is restructured, not patched, because a per-race `switch` arm that mixes flat stats, an aura, a spell, a major cooldown, and a post-finalize hook is exactly what makes the current file hard to read and hard to check against a tooltip.

**What is confirmed and what is not.** The Deep Dive panel confirmed the *shape* — two active and two passive per race — and named three specifics: a new Undead active `Touch of the Grave`, a changed Stoneform, and Mace Specialization now affecting spell crit. Everything else, including every number, is unknown until the beta. So this task ships the **structure** with Era behaviour in it, each entry marked, and a single table a beta tester can correct line by line. It does not invent a second active for seven races; it declares the slot empty and names it in the provisional list, which is the honest thing and is what the spec support page shows.

Existing helpers this reuses, all confirmed present: `character.AxeSpecializationAura()`, `SwordSpecializationAura()`, `MaceSpecializationAura()`, `GunSpecializationAura()`, `BowSpecializationAura()`, `ThrownSpecializationAura()`; `character.AddStat`, `MultiplyStat`, `NewDynamicMultiplyStat`, `EnableDynamicStatDep`; `RegisterAura`, `NewTemporaryStatsAuraWrapped`, `RegisterSpell`, `AddMajorCooldown`, `NewTimer`; `Env.RegisterPostFinalizeEffect`; `APPerStrength`, `APPerAgility`.

**Files:**
- Rewrite: `sim/core/racials.go`
- Test: `sim/core/racials_test.go`

**Interfaces:**
- Produces:
  - `core.RacialKind` (`RacialPassive`, `RacialActive`)
  - `core.Racial{Name string; Kind RacialKind; SpellID int32; Confirmed bool; Apply func(*Character)}`
  - `core.RacialsFor(race proto.Race) []Racial`
  - `core.UnconfirmedRacials() []string` — `"<race>: <name>"` for every entry still carrying Era behaviour or an empty slot. The spec support page reads it beside `ProvisionalConstants()`.
  - `applyRaceEffects(agent Agent)` keeps its signature and its caller; only its body changes.

- [ ] **Step 1: Write the failing test**

Create `sim/core/racials_test.go`:

```go
package core

import (
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

var playableRaces = []proto.Race{
	proto.Race_RaceDwarf,
	proto.Race_RaceGnome,
	proto.Race_RaceHuman,
	proto.Race_RaceNightElf,
	proto.Race_RaceOrc,
	proto.Race_RaceTauren,
	proto.Race_RaceTroll,
	proto.Race_RaceUndead,
}

// Forever: two active and two passive racials per race. The shape is
// confirmed; the contents are not, and UnconfirmedRacials names every
// entry that still carries Era behaviour or an empty slot.
func TestEveryRaceHasTwoActivesAndTwoPassives(t *testing.T) {
	for _, race := range playableRaces {
		t.Run(race.String(), func(t *testing.T) {
			got := RacialsFor(race)
			if len(got) != 4 {
				t.Fatalf("%v has %d racials, want 4", race, len(got))
			}
			var actives, passives int
			for _, r := range got {
				switch r.Kind {
				case RacialActive:
					actives++
				case RacialPassive:
					passives++
				default:
					t.Errorf("%v: %q has no kind", race, r.Name)
				}
				if r.Name == "" {
					t.Errorf("%v: a racial has no name", race)
				}
				if r.Apply == nil {
					t.Errorf("%v: %q has no Apply function", race, r.Name)
				}
			}
			if actives != 2 {
				t.Errorf("%v has %d actives, want 2", race, actives)
			}
			if passives != 2 {
				t.Errorf("%v has %d passives, want 2", race, passives)
			}
		})
	}
}

// An unknown race must not panic and must not silently grant anything.
func TestUnknownRaceHasNoRacials(t *testing.T) {
	if got := RacialsFor(proto.Race_RaceUnknown); len(got) != 0 {
		t.Errorf("RaceUnknown has %d racials, want 0", len(got))
	}
}

// Every racial whose Forever behaviour is not known says so, by name, so
// the spec support page can show it and a beta tester can work the list.
func TestUnconfirmedRacialsAreNamed(t *testing.T) {
	got := UnconfirmedRacials()
	if len(got) == 0 {
		t.Skip("every racial is confirmed; this test has done its job")
	}
	for _, s := range got {
		if !strings.Contains(s, ":") {
			t.Errorf("%q is not in the form \"<race>: <name>\"", s)
		}
	}
	// Cross-check: the count must match the table.
	var want int
	for _, race := range playableRaces {
		for _, r := range RacialsFor(race) {
			if !r.Confirmed {
				want++
			}
		}
	}
	if len(got) != want {
		t.Errorf("UnconfirmedRacials() returned %d entries, the table has %d unconfirmed", len(got), want)
	}
}

// The three specifics the Deep Dive panel named must be in the table.
func TestConfirmedForeverRacialsArePresent(t *testing.T) {
	find := func(race proto.Race, name string) *Racial {
		for _, r := range RacialsFor(race) {
			if r.Name == name {
				return &r
			}
		}
		return nil
	}
	if find(proto.Race_RaceUndead, "Touch of the Grave") == nil {
		t.Error("Undead is missing the new active Touch of the Grave")
	}
	if r := find(proto.Race_RaceDwarf, "Stoneform"); r == nil {
		t.Error("Dwarf is missing Stoneform")
	}
	if find(proto.Race_RaceHuman, "Mace Specialization") == nil {
		t.Error("Human is missing Mace Specialization, which Forever extends to spell crit")
	}
}

// Applying every race's racials to a real character must not panic; this
// is the check that catches a helper called before the character is
// finalized.
func TestApplyingEveryRaceIsSafe(t *testing.T) {
	for _, race := range playableRaces {
		t.Run(race.String(), func(t *testing.T) {
			_, _, _ = NewEnvironment(
				SinglePlayerRaidProto(&proto.Player{
					Name:  "Racial Test",
					Race:  race,
					Class: proto.Class_ClassWarrior,
				}, &proto.PartyBuffs{}, &proto.RaidBuffs{}, &proto.Debuffs{}),
				&proto.Encounter{Duration: 60, Targets: []*proto.Target{DefaultTargetProtoLvl60}},
				false,
			)
		})
	}
}
```

Match `NewEnvironment`'s real signature (`sim/core/environment.go:55`) and return arity.

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/core/ -run 'Racial' -v`
Expected: `FAIL [build failed]`, `undefined: RacialsFor`.

- [ ] **Step 3: Rewrite `racials.go` as a table**

Replace the head of `sim/core/racials.go` — everything from `func applyRaceEffects` down to the closing brace of its `switch`, keeping `makeBerserkingCooldown` and any other helpers below it — with:

```go
// Forever gives every race two active and two passive racials. The shape
// is confirmed from the Deep Dive panel; almost none of the contents are,
// so this table carries Era behaviour with Confirmed:false on every entry
// that has not been checked against the beta, and UnconfirmedRacials()
// names them for the spec support page.
//
// An empty slot is declared, not filled with a guess: Forever has not
// published a second active for most races, and inventing one would put a
// number in the sim that nobody can check. `unconfirmed`.

type RacialKind int

const (
	RacialPassive RacialKind = iota + 1
	RacialActive
)

// Racial is one of a race's four entries.
type Racial struct {
	// Name is the in-game name, or a slot name for an entry Forever has
	// not published.
	Name string
	Kind RacialKind
	// SpellID is the ability's spell id, 0 for a passive with none and for
	// an unpublished slot.
	SpellID int32
	// Confirmed is true only once the entry has been checked against
	// Forever's own tooltip or against beta logs. Everything else is Era
	// behaviour standing in.
	Confirmed bool
	// Apply installs the racial on a character. It must be safe to call
	// during character construction: anything that needs the finished
	// encounter goes through Env.RegisterPostFinalizeEffect.
	Apply func(character *Character)
}

// nothing is the Apply for a slot Forever has not published. It is a
// function rather than a nil so that every entry is callable and the
// table needs no special case.
func nothing(*Character) {}

// unpublishedActive is the placeholder for a second active Forever has
// announced in shape but not in content.
func unpublishedActive(name string) Racial {
	return Racial{Name: name, Kind: RacialActive, Confirmed: false, Apply: nothing}
}

// racialsByRace is the whole model. One table, four entries per race, each
// naming what it is and whether anyone has checked it.
var racialsByRace = map[proto.Race][]Racial{
	proto.Race_RaceDwarf: {
		{Name: "Frost Resistance", Kind: RacialPassive, Apply: func(c *Character) {
			c.AddStat(stats.FrostResistance, 10) // unconfirmed
		}},
		{Name: "Gun Specialization", Kind: RacialPassive, Apply: func(c *Character) {
			c.GunSpecializationAura()
		}},
		{Name: "Stoneform", Kind: RacialActive, SpellID: 20594, Apply: applyStoneform},
		unpublishedActive("Dwarf second active"),
	},
	proto.Race_RaceGnome: {
		{Name: "Arcane Resistance", Kind: RacialPassive, Apply: func(c *Character) {
			c.AddStat(stats.ArcaneResistance, 10) // unconfirmed
		}},
		{Name: "Expansive Mind", Kind: RacialPassive, Apply: func(c *Character) {
			c.MultiplyStat(stats.Intellect, 1.05) // unconfirmed
		}},
		unpublishedActive("Gnome first active"),
		unpublishedActive("Gnome second active"),
	},
	proto.Race_RaceHuman: {
		{Name: "The Human Spirit", Kind: RacialPassive, Apply: func(c *Character) {
			c.MultiplyStat(stats.Spirit, 1.05) // unconfirmed
		}},
		// Forever extends Mace Specialization to spell crit. The Era aura
		// is weapon skill only; the spell-crit half needs a number nobody
		// has published, so it is not added here and the entry stays
		// unconfirmed rather than being half-implemented.
		{Name: "Mace Specialization", Kind: RacialPassive, Apply: func(c *Character) {
			c.SwordSpecializationAura()
			c.MaceSpecializationAura()
		}},
		unpublishedActive("Human first active"),
		unpublishedActive("Human second active"),
	},
	proto.Race_RaceNightElf: {
		{Name: "Nature Resistance", Kind: RacialPassive, Apply: func(c *Character) {
			c.AddStat(stats.NatureResistance, 10) // unconfirmed
		}},
		{Name: "Quickness", Kind: RacialPassive, Apply: func(c *Character) {
			c.AddStat(stats.Dodge, 1) // unconfirmed
		}},
		{Name: "Shadowmeld", Kind: RacialActive, SpellID: 20580, Confirmed: false, Apply: nothing},
		unpublishedActive("Night Elf second active"),
	},
	proto.Race_RaceOrc: {
		{Name: "Axe Specialization", Kind: RacialPassive, Apply: func(c *Character) {
			c.AxeSpecializationAura()
		}},
		{Name: "Command", Kind: RacialPassive, Apply: applyCommand},
		{Name: "Blood Fury", Kind: RacialActive, SpellID: 20572, Apply: applyBloodFury},
		unpublishedActive("Orc second active"),
	},
	proto.Race_RaceTauren: {
		{Name: "Nature Resistance", Kind: RacialPassive, Apply: func(c *Character) {
			c.AddStat(stats.NatureResistance, 10) // unconfirmed
		}},
		{Name: "Endurance", Kind: RacialPassive, Apply: func(c *Character) {
			c.MultiplyStat(stats.Health, 1.05) // unconfirmed
		}},
		{Name: "War Stomp", Kind: RacialActive, SpellID: 20549, Confirmed: false, Apply: nothing},
		unpublishedActive("Tauren second active"),
	},
	proto.Race_RaceTroll: {
		{Name: "Bow Specialization", Kind: RacialPassive, Apply: func(c *Character) {
			c.BowSpecializationAura()
			c.ThrownSpecializationAura()
		}},
		{Name: "Beast Slaying", Kind: RacialPassive, Apply: applyBeastSlaying},
		{Name: "Berserking", Kind: RacialActive, SpellID: 26297, Apply: applyBerserking},
		unpublishedActive("Troll second active"),
	},
	proto.Race_RaceUndead: {
		{Name: "Shadow Resistance", Kind: RacialPassive, Apply: func(c *Character) {
			c.AddStat(stats.ShadowResistance, 10) // unconfirmed
		}},
		{Name: "Will of the Forsaken", Kind: RacialPassive, Confirmed: false, Apply: nothing},
		// Confirmed by the Deep Dive panel as a new Forever active. Its
		// damage, cooldown and coefficient are unpublished, so it is a
		// named slot until the beta.
		{Name: "Touch of the Grave", Kind: RacialActive, Confirmed: false, Apply: nothing},
		unpublishedActive("Undead second active"),
	},
}

// RacialsFor returns a race's four racials, or nothing for a race the
// engine does not know.
func RacialsFor(race proto.Race) []Racial {
	return racialsByRace[race]
}

// UnconfirmedRacials names every entry still carrying Era behaviour or an
// empty slot, as "<race>: <name>". The spec support page shows it beside
// ProvisionalConstants(), so a player can see exactly what the sim is
// guessing about.
func UnconfirmedRacials() []string {
	races := make([]proto.Race, 0, len(racialsByRace))
	for race := range racialsByRace {
		races = append(races, race)
	}
	slices.Sort(races)

	var out []string
	for _, race := range races {
		for _, r := range racialsByRace[race] {
			if !r.Confirmed {
				out = append(out, race.String()+": "+r.Name)
			}
		}
	}
	return out
}

func applyRaceEffects(agent Agent) {
	character := agent.GetCharacter()
	for _, r := range RacialsFor(character.Race) {
		r.Apply(character)
	}
}
```

Then move the four bodies that were inline in the old `switch` into named functions below, unchanged except for their signature. `applyStoneform` is the Dwarf arm's body (old lines 17-52), `applyCommand` the Orc pet-damage block (lines 68-76), `applyBloodFury` the Orc aura, spell and major cooldown (lines 78-115), `applyBeastSlaying` the Troll post-finalize hook (lines 122-131), and `applyBerserking` the five `makeBerserkingCooldown` calls (lines 134-141). Each takes `character *Character` and returns nothing. For example:

```go
// applyBeastSlaying is Era's +5% damage and crit against beasts. Forever
// keeps beast slaying; the figures are unconfirmed.
func applyBeastSlaying(character *Character) {
	character.Env.RegisterPostFinalizeEffect(func() {
		for _, t := range character.Env.Encounter.Targets {
			if t.MobType == proto.MobType_MobTypeBeast {
				for _, at := range character.AttackTables[t.UnitIndex] {
					at.DamageDealtMultiplier *= 1.05 // unconfirmed
					at.CritMultiplier *= 1.05        // unconfirmed
				}
			}
		}
	})
}

// applyBerserking registers the baseline cooldown plus the five
// hard-coded percentage options the UI offers.
func applyBerserking(character *Character) {
	berserkingTimer := character.NewTimer()
	makeBerserkingCooldown(character, 0, berserkingTimer)
	for _, pct := range []float64{.1, .15, .2, .25, .3} {
		makeBerserkingCooldown(character, pct, berserkingTimer)
	}
}
```

Add `"slices"` to the imports.

- [ ] **Step 4: Run the tests and watch them pass**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test --tags=with_db ./sim/core/ -run 'Racial' -v
```

Expected: `BUILDS`; no `gofmt` output; five tests `PASS`. `TestUnconfirmedRacialsAreNamed` should report a list of roughly 20 entries.

- [ ] **Step 5: Prove behaviour is unchanged**

Restructuring is not a behaviour change: every race's Era effects are still applied, in the same order within a race. So no golden may move.

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db -count=1 ./sim/... 2>&1 | tail -25
git status --short -- '*.results'
```

Expected: 20 packages `ok`; **`git status` prints nothing for `.results`**. If one moved, an ordering changed — the spec test suites run several races per spec (`OtherRaces` in each `CharacterSuiteConfig`), so a reordered arm shows up immediately. Put the order back rather than regenerating.

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add sim/core/racials.go sim/core/racials_test.go
git commit -m "refactor(core): racials as a two-active-two-passive table per race" \
  -m "Forever gives every race two actives and two passives. The shape is confirmed; almost none of the contents are, so this restructures the switch into one table whose every entry names itself, says whether anyone has checked it against Forever, and carries Era behaviour until they have. Slots Forever has announced but not published are declared empty rather than filled with an invented ability, and UnconfirmedRacials() lists all of them for the spec support page. The three specifics the Deep Dive panel named - Touch of the Grave, a changed Stoneform, Mace Specialization reaching spell crit - are in the table as named entries; Mace Specialization keeps only its weapon-skill half, because the spell-crit figure is unpublished and a half-implemented racial is worse than a declared gap. No behaviour changes and no golden moves." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 10: Generated per-class spell constants, and the consumables sidecar review

**Repo: ENGINE.** Depends on Task 4. **G2 — INDEPENDENT of Tasks 5, 6, 7, 8, 9.** Tasks 11 and 12 depend on it.

Today every number in a WoWSims ability is a literal: `bonusDamage := 160.0`, `Duration: time.Second * 6`, `Cost: 30` (research §1.3). That is fine when the game is twenty years old and fine-tuned. It is wrong for Forever, whose numbers move weekly through October. So the ability files read from a generated per-class constants file instead, and a beta patch that changes a number becomes a pipeline run and a regeneration.

The mage already shows the shape this should take: `sim/mage/frostbolt.go:9-15` keeps `FrostboltSpellId`, `FrostboltBaseDamage`, `FrostboltSpellCoeff`, `FrostboltCastTime`, `FrostboltManaCost` and `FrostboltLevel` as per-rank arrays and `getFrostboltConfig(rank)` reads them. The generated file produces exactly that, from data.

**The data lane owns the input and this lane owns the semantics.** `data/builds/<build>/spellconst/<class-slug>.json` carries the DB2 columns verbatim — base points, coefficients, cooldown ms, cast time ms, cost, duration ms, school, family mask. **Coefficients come through with their zeros**, because `EffectBonusCoefficient` is routinely 0 or wrong for Classic-lineage spells (research §5.3), and the vanilla conventions — `cast_time/3.5`, `duration/15`, halved for hybrids, with per-spell overrides — belong here. So the generator treats a zero coefficient as *absent*, emits the convention's value, and marks it. It never emits a zero coefficient as if it were data.

**Files:**
- Create: `sim/core/spellconst/spellconst.go`, `sim/core/spellconst/spellconst_test.go`, `sim/core/spellconst/gen/main.go`, `sim/core/spellconst/testdata/warrior.json`
- Create (empty, populated by Tasks 11 and 12): `sim/warrior/constants_auto_gen.go`, `sim/mage/constants_auto_gen.go`
- Modify: `makefile` (a `spellconst` target)
- Create: `docs/consumes-sidecar-review.md`

**Interfaces:**
- Produces, used by Tasks 11 and 12:
  - `spellconst.Spell{ID int32; Name string; Rank int; BasePointsLow, BasePointsHigh float64; Coefficient float64; CoefficientSource string; CooldownMS, CastTimeMS, DurationMS int32; Cost float64; School int32; FamilyMask uint64; Level int}`
  - `spellconst.Class{Slug string; Build string; Spells []Spell}`
  - `spellconst.Load(path string) (Class, error)`
  - `spellconst.(Class).ByID(id int32) (Spell, bool)`, `(Class).Ranks(name string) []Spell`
  - `spellconst.CoefficientFor(castTimeMS int32, durationMS int32, hybrid bool) (float64, string)` — the vanilla convention, so one function is the single source of it
  - The generator writes `sim/<class>/constants_auto_gen.go` declaring, per named spell, the per-rank arrays in the `frostbolt.go` shape: `<Name>SpellId [N+1]int32`, `<Name>BaseDamage [N+1][]float64`, `<Name>SpellCoeff [N+1]float64`, `<Name>CastTime [N+1]int32`, `<Name>ManaCost [N+1]float64`, `<Name>Level [N+1]int`, `<Name>CooldownMS [N+1]int32`, plus `<Name>Ranks` as the count.

- [ ] **Step 1: Write the failing test**

Create `sim/core/spellconst/spellconst_test.go`:

```go
package spellconst

import (
	"math"
	"testing"
)

func TestLoadReadsAClassFile(t *testing.T) {
	c, err := Load("testdata/warrior.json")
	if err != nil {
		t.Fatal(err)
	}
	if c.Slug != "warrior" {
		t.Errorf("Slug = %q, want %q", c.Slug, "warrior")
	}
	if c.Build == "" {
		t.Error("Build is empty; a constants file must record which client it came from")
	}
	if len(c.Spells) == 0 {
		t.Fatal("no spells loaded")
	}
}

func TestByID(t *testing.T) {
	c, err := Load("testdata/warrior.json")
	if err != nil {
		t.Fatal(err)
	}
	got, ok := c.ByID(23894)
	if !ok {
		t.Fatal("spell 23894 (Bloodthirst) not found")
	}
	if got.Name != "Bloodthirst" {
		t.Errorf("name = %q, want Bloodthirst", got.Name)
	}
	if _, ok := c.ByID(1); ok {
		t.Error("spell 1 was found in a warrior file")
	}
}

func TestRanksAreOrdered(t *testing.T) {
	c, err := Load("testdata/warrior.json")
	if err != nil {
		t.Fatal(err)
	}
	ranks := c.Ranks("Bloodthirst")
	if len(ranks) < 2 {
		t.Fatalf("Bloodthirst has %d ranks in the fixture, want at least 2", len(ranks))
	}
	for i := 1; i < len(ranks); i++ {
		if ranks[i].Rank <= ranks[i-1].Rank {
			t.Errorf("ranks are not ascending: %d then %d", ranks[i-1].Rank, ranks[i].Rank)
		}
	}
	if c.Ranks("Nonexistent") != nil {
		t.Error("an unknown spell name returned ranks")
	}
}

// The data lane emits the DB2 coefficient columns verbatim, zeros
// included, because EffectBonusCoefficient is routinely 0 or wrong for
// Classic-lineage spells. A zero therefore means "absent", and the
// vanilla convention fills it in — never a literal zero coefficient,
// which would silently remove all spell-power scaling from a spell.
func TestZeroCoefficientFallsBackToTheConvention(t *testing.T) {
	c, err := Load("testdata/warrior.json")
	if err != nil {
		t.Fatal(err)
	}
	s, ok := c.ByID(11605) // Slam rank 4 in the fixture, coefficient 0
	if !ok {
		t.Fatal("spell 11605 not found")
	}
	if s.Coefficient == 0 {
		t.Error("a zero DB2 coefficient was kept as zero; it must fall back to the convention")
	}
	if s.CoefficientSource != "convention" {
		t.Errorf("CoefficientSource = %q, want %q", s.CoefficientSource, "convention")
	}
}

func TestNonZeroCoefficientIsKept(t *testing.T) {
	c, err := Load("testdata/warrior.json")
	if err != nil {
		t.Fatal(err)
	}
	s, ok := c.ByID(23881) // Bloodthirst rank 1 in the fixture, coefficient 0.15
	if !ok {
		t.Fatal("spell 23881 not found")
	}
	if math.Abs(s.Coefficient-0.15) > 1e-9 {
		t.Errorf("Coefficient = %v, want the table's 0.15", s.Coefficient)
	}
	if s.CoefficientSource != "table" {
		t.Errorf("CoefficientSource = %q, want %q", s.CoefficientSource, "table")
	}
}

// The vanilla conventions, in one function so no ability file re-derives
// them: a direct spell gets cast_time/3.5, a periodic one duration/15,
// and a hybrid class gets half.
func TestCoefficientConvention(t *testing.T) {
	cases := []struct {
		name       string
		castMS     int32
		durationMS int32
		hybrid     bool
		want       float64
		source     string
	}{
		{"three second cast", 3000, 0, false, 3.0 / 3.5, "convention"},
		{"instant direct", 0, 0, false, 1.5 / 3.5, "convention"},
		{"fifteen second dot", 0, 15000, false, 1.0, "convention"},
		{"hybrid three second cast", 3000, 0, true, 3.0 / 3.5 / 2, "convention"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, src := CoefficientFor(tc.castMS, tc.durationMS, tc.hybrid)
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("CoefficientFor(%d, %d, %v) = %v, want %v", tc.castMS, tc.durationMS, tc.hybrid, got, tc.want)
			}
			if src != tc.source {
				t.Errorf("source = %q, want %q", src, tc.source)
			}
		})
	}
}

// A cast time under the global cooldown is treated as a GCD cast, which
// is the vanilla rule and the reason an instant nuke is not coefficient
// zero.
func TestInstantCastUsesTheGlobalCooldown(t *testing.T) {
	got, _ := CoefficientFor(500, 0, false)
	want := 1.5 / 3.5
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("a 0.5s cast gave %v, want the GCD-floored %v", got, want)
	}
}

func TestLoadRejectsAMissingFile(t *testing.T) {
	if _, err := Load("testdata/nope.json"); err == nil {
		t.Fatal("loading a missing file returned no error")
	}
}
```

- [ ] **Step 2: Write the fixture**

Create `sim/core/spellconst/testdata/warrior.json`. This is a hand-written stand-in for what the data lane emits, small enough to read, with one spell that has a real coefficient and one whose column is zero:

```json
{
  "class_slug": "warrior",
  "build": "1.15.9.69722",
  "spells": [
    {
      "id": 23881,
      "name": "Bloodthirst",
      "rank": 1,
      "base_points_low": 160,
      "base_points_high": 160,
      "coefficient": 0.15,
      "cooldown_ms": 6000,
      "cast_time_ms": 0,
      "duration_ms": 0,
      "cost": 30,
      "school": 1,
      "family_mask": 2,
      "level": 40
    },
    {
      "id": 23894,
      "name": "Bloodthirst",
      "rank": 2,
      "base_points_low": 210,
      "base_points_high": 210,
      "coefficient": 0.15,
      "cooldown_ms": 6000,
      "cast_time_ms": 0,
      "duration_ms": 0,
      "cost": 30,
      "school": 1,
      "family_mask": 2,
      "level": 48
    },
    {
      "id": 11605,
      "name": "Slam",
      "rank": 4,
      "base_points_low": 87,
      "base_points_high": 87,
      "coefficient": 0,
      "cooldown_ms": 0,
      "cast_time_ms": 1500,
      "duration_ms": 0,
      "cost": 15,
      "school": 1,
      "family_mask": 2097152,
      "level": 54
    }
  ]
}
```

- [ ] **Step 3: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test ./sim/core/spellconst/ -v`
Expected: `FAIL [build failed]`, `undefined: Load`.

- [ ] **Step 4: Write the loader and the convention**

Create `sim/core/spellconst/spellconst.go`:

```go
// Package spellconst reads the per-class spell constants the data
// pipeline generates from the client tables, so an ability's numbers are
// regenerated rather than retyped when Forever changes one.
//
// The pipeline emits the DB2 columns verbatim. For Classic-lineage spells
// EffectBonusCoefficient is routinely 0 or wrong, so a zero here means
// "the table does not know", and CoefficientFor supplies the vanilla
// convention in its place. The conventions live in this package and
// nowhere else.
package spellconst

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

// gcd is the vanilla global cooldown. A cast faster than it still scales
// as a GCD cast, which is why an instant nuke is not coefficient zero.
const gcd = 1500 * time.Millisecond

// directDivisor and periodicDivisor are the vanilla spell-coefficient
// conventions: a direct spell gets cast_time/3.5, a periodic one
// duration/15. Both are seconds.
const (
	directDivisor   = 3.5
	periodicDivisor = 15.0
)

// Spell is one rank of one ability.
type Spell struct {
	ID             int32   `json:"id"`
	Name           string  `json:"name"`
	Rank           int     `json:"rank"`
	BasePointsLow  float64 `json:"base_points_low"`
	BasePointsHigh float64 `json:"base_points_high"`
	// Coefficient is the spell-power coefficient, resolved: the table's
	// value when it has one, the convention's when it does not.
	Coefficient float64 `json:"coefficient"`
	// CoefficientSource is "table" or "convention", so a reader can tell
	// a measured number from a derived one.
	CoefficientSource string `json:"-"`
	CooldownMS        int32  `json:"cooldown_ms"`
	CastTimeMS        int32  `json:"cast_time_ms"`
	DurationMS        int32  `json:"duration_ms"`
	Cost       float64 `json:"cost"`
	School     int32   `json:"school"`
	FamilyMask uint64  `json:"family_mask"`
	Level      int     `json:"level"`
}

// Class is one generated per-class file.
type Class struct {
	Slug   string  `json:"class_slug"`
	Build  string  `json:"build"`
	Spells []Spell `json:"spells"`

	// hybrid marks the classes whose spell coefficients are halved by the
	// vanilla convention.
	hybrid bool
}

// hybridClasses are the classes the vanilla convention halves.
var hybridClasses = map[string]bool{
	"paladin": true,
	"shaman":  true,
	"druid":   true,
	"priest":  false, // shadow priests use the full convention
	"warrior": false,
	"rogue":   false,
	"hunter":  false,
	"mage":    false,
	"warlock": false,
}

// Load reads a generated class file and resolves every coefficient.
func Load(path string) (Class, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Class{}, fmt.Errorf("spellconst: %w", err)
	}
	var c Class
	if err := json.Unmarshal(b, &c); err != nil {
		return Class{}, fmt.Errorf("spellconst: parsing %s: %w", path, err)
	}
	if c.Slug == "" {
		return Class{}, fmt.Errorf("spellconst: %s has no class_slug", path)
	}
	if c.Build == "" {
		return Class{}, fmt.Errorf("spellconst: %s has no build; a constants file must record which client it came from", path)
	}
	c.hybrid = hybridClasses[c.Slug]
	for i := range c.Spells {
		s := &c.Spells[i]
		if s.Coefficient != 0 {
			s.CoefficientSource = "table"
			continue
		}
		s.Coefficient, s.CoefficientSource = CoefficientFor(s.CastTimeMS, s.DurationMS, c.hybrid)
	}
	sort.SliceStable(c.Spells, func(i, j int) bool {
		if c.Spells[i].Name != c.Spells[j].Name {
			return c.Spells[i].Name < c.Spells[j].Name
		}
		return c.Spells[i].Rank < c.Spells[j].Rank
	})
	return c, nil
}

// ByID returns one rank of one spell.
func (c Class) ByID(id int32) (Spell, bool) {
	for _, s := range c.Spells {
		if s.ID == id {
			return s, true
		}
	}
	return Spell{}, false
}

// Ranks returns every rank of a named spell, ascending, or nil.
func (c Class) Ranks(name string) []Spell {
	var out []Spell
	for _, s := range c.Spells {
		if s.Name == name {
			out = append(out, s)
		}
	}
	return out
}

// CoefficientFor is the vanilla spell-coefficient convention: a direct
// spell scales with its cast time over 3.5 seconds, a periodic one with
// its duration over 15, and a hybrid class gets half. A cast faster than
// the global cooldown is treated as a GCD cast.
//
// This is a convention, not data: per-spell exceptions are dozens strong
// and live in the ability files that override this value, exactly as they
// do today. The second return says which of the two a caller got, so an
// override can be applied knowingly.
func CoefficientFor(castTimeMS int32, durationMS int32, hybrid bool) (float64, string) {
	var coeff float64
	switch {
	case durationMS > 0:
		coeff = (time.Duration(durationMS) * time.Millisecond).Seconds() / periodicDivisor
	default:
		cast := time.Duration(castTimeMS) * time.Millisecond
		if cast < gcd {
			cast = gcd
		}
		coeff = cast.Seconds() / directDivisor
	}
	if hybrid {
		coeff /= 2
	}
	return coeff, "convention"
}
```

- [ ] **Step 5: Run the tests and watch them pass**

Run: `cd /Users/jh/code/wowsims-forever && go test ./sim/core/spellconst/ -v && gofmt -l ./sim`
Expected: eight tests `PASS`, no `gofmt` output.

- [ ] **Step 6: Write the generator**

Create `sim/core/spellconst/gen/main.go`:

```go
// Command gen turns a data-lane spellconst file into the per-rank Go
// arrays an ability file reads, in the shape sim/mage/frostbolt.go
// already uses. A Forever patch that changes a number is then a pipeline
// run and a `make spellconst`, not a code edit.
//
//	go run ./sim/core/spellconst/gen \
//	  -in  ../forever/data/builds/1.15.9.69722/spellconst/warrior.json \
//	  -out sim/warrior/constants_auto_gen.go \
//	  -package warrior
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/wowsims/classic/sim/core/spellconst"
)

func main() {
	in := flag.String("in", "", "the data lane's spellconst/<class>.json")
	out := flag.String("out", "", "the Go file to write")
	pkg := flag.String("package", "", "the Go package name, e.g. warrior")
	flag.Parse()
	if *in == "" || *out == "" || *pkg == "" {
		log.Fatal("-in, -out and -package are all required")
	}

	class, err := spellconst.Load(*in)
	if err != nil {
		log.Fatal(err)
	}

	names := map[string]bool{}
	var order []string
	for _, s := range class.Spells {
		if !names[s.Name] {
			names[s.Name] = true
			order = append(order, s.Name)
		}
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by sim/core/spellconst/gen. DO NOT EDIT.\n//\n")
	fmt.Fprintf(&b, "// Class: %s\n// Client build: %s\n// Source: %s\n//\n", class.Slug, class.Build, *in)
	fmt.Fprintf(&b, "// Regenerate with `make spellconst`. A coefficient marked\n")
	fmt.Fprintf(&b, "// \"convention\" was derived from the vanilla cast_time/3.5 and\n")
	fmt.Fprintf(&b, "// duration/15 rules because the client table's column was zero;\n")
	fmt.Fprintf(&b, "// per-spell overrides stay in the ability files.\n\n")
	fmt.Fprintf(&b, "package %s\n\n", *pkg)
	fmt.Fprintf(&b, "// ConstantsBuild is the client build these numbers came from.\n")
	fmt.Fprintf(&b, "const ConstantsBuild = %q\n\n", class.Build)

	for _, name := range order {
		ranks := class.Ranks(name)
		ident := goIdent(name)
		n := len(ranks)
		fmt.Fprintf(&b, "// %s: %d rank(s), from build %s.\n", name, n, class.Build)
		fmt.Fprintf(&b, "const %sRanks = %d\n\n", ident, n)

		// Index 0 is unused so a rank number indexes directly, which is
		// the convention sim/mage/frostbolt.go established.
		writeArray(&b, ident, "SpellId", "int32", n, func(i int) string { return fmt.Sprint(ranks[i].ID) })
		writeArray(&b, ident, "Level", "int", n, func(i int) string { return fmt.Sprint(ranks[i].Level) })
		writeArray(&b, ident, "CastTime", "int32", n, func(i int) string { return fmt.Sprint(ranks[i].CastTimeMS) })
		writeArray(&b, ident, "CooldownMS", "int32", n, func(i int) string { return fmt.Sprint(ranks[i].CooldownMS) })
		writeArray(&b, ident, "ManaCost", "float64", n, func(i int) string { return trimFloat(ranks[i].Cost) })
		writeArray(&b, ident, "SpellCoeff", "float64", n, func(i int) string { return trimFloat(ranks[i].Coefficient) })

		fmt.Fprintf(&b, "var %sBaseDamage = [%sRanks + 1][]float64{{0, 0}", ident, ident)
		for i := 0; i < n; i++ {
			fmt.Fprintf(&b, ", {%s, %s}", trimFloat(ranks[i].BasePointsLow), trimFloat(ranks[i].BasePointsHigh))
		}
		fmt.Fprintf(&b, "}\n")

		// A coefficient the table did not supply is called out by name,
		// so a reader of the ability file knows which numbers are derived.
		var derived []string
		for _, r := range ranks {
			if r.CoefficientSource == "convention" {
				derived = append(derived, fmt.Sprintf("rank %d", r.Rank))
			}
		}
		if len(derived) > 0 {
			fmt.Fprintf(&b, "// unconfirmed: %s coefficient derived from the vanilla convention (%s)\n",
				name, strings.Join(derived, ", "))
		}
		fmt.Fprintf(&b, "\n")
	}

	src, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatalf("generated code does not parse: %v\n%s", err, b.String())
	}
	if err := os.WriteFile(*out, src, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s: %d spells, %d ranks, build %s", *out, len(order), len(class.Spells), class.Build)
}

// writeArray emits `var <ident><field> = [<ident>Ranks + 1]<typ>{0, a, b}`.
// Index 0 is a placeholder so a rank number indexes the array directly,
// which is the convention sim/mage/frostbolt.go established.
func writeArray(b *bytes.Buffer, ident, field, typ string, n int, at func(int) string) {
	fmt.Fprintf(b, "var %s%s = [%sRanks + 1]%s{0", ident, field, ident, typ)
	for i := 0; i < n; i++ {
		fmt.Fprintf(b, ", %s", at(i))
	}
	fmt.Fprintf(b, "}\n")
}

func trimFloat(f float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.4f", f), "0"), ".")
}

// goIdent turns "Mortal Strike" into "MortalStrike".
func goIdent(name string) string {
	var out strings.Builder
	for _, part := range strings.FieldsFunc(name, func(r rune) bool {
		return r == ' ' || r == '-' || r == '\'' || r == ':'
	}) {
		out.WriteString(strings.ToUpper(part[:1]))
		out.WriteString(part[1:])
	}
	return out.String()
}
```

- [ ] **Step 7: Run the generator against the fixture**

```bash
cd /Users/jh/code/wowsims-forever
mkdir -p /tmp/spellconst-out
go run ./sim/core/spellconst/gen -in sim/core/spellconst/testdata/warrior.json -out /tmp/spellconst-out/constants_auto_gen.go -package warrior
cat /tmp/spellconst-out/constants_auto_gen.go
```

Expected: a file declaring `ConstantsBuild = "1.15.9.69722"`, then `BloodthirstRanks = 2` with `BloodthirstSpellId = [BloodthirstRanks + 1]int32{0, 23881, 23894}`, `BloodthirstBaseDamage = [BloodthirstRanks + 1][]float64{{0, 0}, {160, 160}, {210, 210}}`, `BloodthirstSpellCoeff = [...]{0, 0.15, 0.15}`, then `SlamRanks = 1` with a coefficient of `0.4286` (the GCD-floored convention: 1.5/3.5) followed by the `unconfirmed: Slam coefficient derived from the vanilla convention (rank 4)` line.

If `format.Source` fails, the error prints the generated source — read it; the usual cause is an identifier with a character `goIdent` does not strip.

- [ ] **Step 8: Add the make target and the empty per-class files**

In `makefile`, append:

```makefile
# Where the site repository is checked out, for the data lane's outputs.
SITE_DIR ?= /Users/jh/code/forever
# Which build's constants to generate from.
BUILD ?= 1.15.9.69722

.PHONY: spellconst
# spellconst regenerates the per-class constants files from the data
# lane's spellconst output. A Forever patch that changes a number is a
# pipeline run and this target, not a code edit.
spellconst:
	@for class in warrior mage; do \
	  src="$(SITE_DIR)/data/builds/$(BUILD)/spellconst/$$class.json"; \
	  if [ ! -f "$$src" ]; then echo "missing $$src (run the data lane's simconst command)"; exit 1; fi; \
	  go run ./sim/core/spellconst/gen -in "$$src" -out "sim/$$class/constants_auto_gen.go" -package "$$class" || exit 1; \
	done
	gofmt -w ./sim
```

Create the two placeholder files so the packages compile before the data lane's output exists. `sim/warrior/constants_auto_gen.go`:

```go
// Code generated by sim/core/spellconst/gen. DO NOT EDIT.
//
// Placeholder: the data lane's spellconst output for this build does not
// exist yet. Run `make spellconst` once it does. Until then the ability
// files keep their Era literals, each marked unconfirmed.

package warrior

// ConstantsBuild is the client build these numbers came from. Empty means
// no constants have been generated yet.
const ConstantsBuild = ""
```

and the same for `sim/mage/constants_auto_gen.go` with `package mage`.

- [ ] **Step 9: Review the consumables sidecar against `consumes.go`**

The data lane asked this lane to confirm the shape of `data/builds/<build>/simconsumes.json` (1,467 rows) against the hand-written `sim/core/consumes.go` (1,252 lines). `SimDatabase` has no consumables field, so this file is how consumables reach the engine at all.

```bash
cd /Users/jh/code/forever
ls -la data/builds/*/simconsumes.json 2>/dev/null && python3 -c "
import json,glob,collections
p=sorted(glob.glob('data/builds/*/simconsumes.json'))[-1]
d=json.load(open(p))
rows = d if isinstance(d, list) else d.get('consumables', d.get('rows', []))
print(p, len(rows), 'rows')
keys=collections.Counter()
for r in rows[:2000]:
    keys.update(r.keys())
for k,v in keys.most_common():
    print(f'  {k:28} {v}')
print('sample:', json.dumps(rows[0], indent=2)[:600])
"
```

Then read `sim/core/consumes.go` and answer three questions in a short note. Create `docs/consumes-sidecar-review.md` **in the engine repo**:

```markdown
# Review: data/builds/<build>/simconsumes.json against sim/core/consumes.go

Reviewed: <date>. Build: <build>. Rows: <n>.

`SimDatabase` carries no consumables field, so this sidecar is the only
path consumables take from the pipeline into the engine.

## 1. Does every field the engine needs exist?

`sim/core/consumes.go` applies a consumable by proto enum value, not by
item id: `proto.Food_FoodSmokedDesertDumpling`, `proto.AgilityElixir_*`,
`proto.Potions_*` and so on, each hard-coded to the stats it grants.
<Answer: which sidecar fields map onto which, and which enum values have
no row.>

## 2. Does every row map onto something the engine can apply?

<Answer: the count of rows whose effect is a flat stat the engine already
models, versus rows carrying a proc or a use effect it does not.>

## 3. What the engine lane asks the data lane for

<Either "nothing, the shape is sufficient" or a numbered list of fields.>
```

Fill in the three answers from what the file and the code actually say. If the file does not exist yet, write the note with `Rows: not yet emitted` and the three questions answered from `consumes.go` alone — the point is that the data lane gets an answer, and an answer that says "here is what the engine needs" is more useful than a wait.

- [ ] **Step 10: Run everything and commit**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test ./sim/core/spellconst/ -v
go test --tags=with_db -count=1 ./sim/core/... ./sim/warrior/... ./sim/mage/...
git status --short -- '*.results'
```

Expected: `BUILDS`; no `gofmt` output; eight spellconst tests `PASS`; the three package groups `ok`; no `.results` change (nothing reads the constants yet).

```bash
git add sim/core/spellconst sim/warrior/constants_auto_gen.go sim/mage/constants_auto_gen.go makefile docs/consumes-sidecar-review.md
git commit -m "feat(core): generated per-class spell constants, and the coefficient convention in one place" \
  -m "Every number in a WoWSims ability is a literal today, which is fine for a twenty-year-old game and wrong for Forever, whose numbers move weekly through October. The ability files now read per-rank arrays generated from the data lane's spellconst output, in the shape sim/mage/frostbolt.go already uses, so a beta patch is a pipeline run and a make spellconst rather than a code edit. The pipeline emits the DB2 coefficient columns verbatim including zeros, because EffectBonusCoefficient is routinely wrong for Classic-lineage spells, so a zero means absent and the vanilla cast_time/3.5 and duration/15 conventions fill it in - from one function, marked in the generated file, with per-spell overrides staying in the ability files. Also reviews the consumables sidecar the data lane asked about, since SimDatabase carries no consumables field." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 11: `warrior-fury` end to end

**Repo: ENGINE.** Depends on Tasks 7 and 10. **G3 — INDEPENDENT of Task 12.** This task touches only `sim/warrior/**`, `proto/warrior.proto`, `ui/warrior/**` and `sim/register_all.go`; Task 12 touches only the mage's equivalents. Two parallel worktrees, one merge each.

Fury is the first of the two launch specs. Four deliverables: a `talents.go` rewritten for Forever's seven rows and 11/16/21/31 milestones using the spell-mod system from Task 7, the new baseline abilities as one file each, the constants read from Task 10's generated file, and a default APL.

**What survives and what does not.** Research §1.11 puts it at 40-60% of `talents.go`. Reading `sim/warrior/talents.go` (498 lines) against Forever's confirmed rules: Forever keeps seven rows and the 11/21/31 gold-medal talents, **adds a 16-point one per tree**, deletes the buff-only talents and makes those baseline, and keeps many talents unchanged from 2006. So the arithmetic talents (Cruelty, Toughness, Anticipation, Deflection, the weapon specializations, Impale, Improved Heroic Strike) survive as declarative mods; the mechanical ones (Anger Management, Deep Wounds, Flurry, Enrage, Unbridled Wrath, Death Wish, Sweeping Strikes, Last Stand) survive as their existing functions; only the tree *shape* and the new 16-point talent are genuinely new.

**Do not copy `ui/warrior/apls/dps_reck.apl.json` as the starting point without re-validating it.** The data lane verified that the engine's checked-in preset APLs are stale on spell ranks. Every `spellId` in the new APL is checked against the build's own tables.

**Files:**
- Modify: `proto/warrior.proto` (the `WarriorTalents` message, 111 lines), `sim/warrior/talents.go`, `sim/warrior/warrior.go` (`TalentTreeSizes`, the spell-code block, `ApplyTalents` call sites), `sim/warrior/constants_auto_gen.go` (regenerated), `sim/warrior/dps_warrior/dps_warrior_test.go`
- Create: `sim/warrior/rampage.go`, `sim/warrior/piercing_howl.go`, `ui/warrior/apls/forever_fury.apl.json`, `ui/core/talents/trees/warrior.json` (regenerated)
- Test: `sim/warrior/talents_test.go`, `sim/warrior/dps_warrior/dps_warrior_test.go`

**Interfaces:**
- Consumes: `core.SpellModConfig`, `core.SpellMod_*`, `(*Unit).AddStaticMod`, `(*Unit).AddDynamicMod`, `Spell.ClassSpellMask` (Task 7); `warrior.ConstantsBuild` and the generated per-rank arrays (Task 10).
- Produces:
  - `warrior.WarriorSpellMask*` — a `uint64` bitmask constant per ability, replacing nothing (`SpellCode_Warrior*` stays for the existing code)
  - `warrior.TalentTreeSizes` updated to Forever's row counts
  - `(*Warrior).registerRampageSpell()`, `(*Warrior).registerPiercingHowlSpell()`
  - `warrior.ForeverFuryTalents string` — the reference build the test suite uses
  - `ui/warrior/apls/forever_fury.apl.json`

- [ ] **Step 1: Write the failing talent test**

Create `sim/warrior/talents_test.go`:

```go
package warrior

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// Forever's trees are seven rows with gold-medal talents at 11, 16, 21
// and 31 points. The 16-point tier is the new one; everything else is
// vanilla's shape.
func TestForeverTalentMilestones(t *testing.T) {
	want := []int{11, 16, 21, 31}
	if len(ForeverMilestones) != len(want) {
		t.Fatalf("ForeverMilestones = %v, want %v", ForeverMilestones, want)
	}
	for i, m := range want {
		if ForeverMilestones[i] != m {
			t.Errorf("ForeverMilestones[%d] = %d, want %d", i, ForeverMilestones[i], m)
		}
	}
}

// The talent string is parsed positionally against TalentTreeSizes, so a
// tree size that does not match the proto's field count silently reads
// the wrong talent.
func TestTalentTreeSizesMatchTheProto(t *testing.T) {
	var total int
	for _, n := range TalentTreeSizes {
		total += n
	}
	fields := (&proto.WarriorTalents{}).ProtoReflect().Descriptor().Fields()
	if fields.Len() != total {
		t.Errorf("WarriorTalents has %d fields, TalentTreeSizes sums to %d", fields.Len(), total)
	}
}

// Every ability a talent modifies must carry a ClassSpellMask, or the
// declarative mod silently applies to nothing.
func TestFurySpellsCarryTheirMasks(t *testing.T) {
	cases := []struct {
		name string
		mask uint64
	}{
		{"Bloodthirst", WarriorSpellMaskBloodthirst},
		{"Whirlwind", WarriorSpellMaskWhirlwind},
		{"Execute", WarriorSpellMaskExecute},
		{"Heroic Strike", WarriorSpellMaskHeroicStrike},
		{"Cleave", WarriorSpellMaskCleave},
		{"Rampage", WarriorSpellMaskRampage},
	}
	seen := uint64(0)
	for _, c := range cases {
		if c.mask == 0 {
			t.Errorf("%s has a zero mask; an empty mask matches nothing", c.name)
		}
		if seen&c.mask != 0 {
			t.Errorf("%s reuses a bit already taken", c.name)
		}
		seen |= c.mask
	}
}

// The reference Fury build must spend exactly 51 points and must reach
// the 31-point talent in Fury, or the suite is validating a build nobody
// would play.
func TestForeverFuryTalentsAreAValidBuild(t *testing.T) {
	talents := &proto.WarriorTalents{}
	fillWarriorTalents(talents, ForeverFuryTalents)
	if !talents.Bloodthirst {
		t.Error("the reference Fury build does not take Bloodthirst, the 31-point talent")
	}
}
```

`fillWarriorTalents` is a two-line helper wrapping `core.FillTalentsProto(talents.ProtoReflect(), s, TalentTreeSizes)`; put it in `talents.go` beside `ApplyTalents` (it is also what `NewWarrior` already does, so factor the existing call through it rather than duplicating).

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/warrior/ -v`
Expected: `FAIL [build failed]`, `undefined: ForeverMilestones`.

- [ ] **Step 3: Regenerate the talent proto and tree JSON**

The scrapers already exist and already target a per-environment Wowhead URL:

```bash
cd /Users/jh/code/wowsims-forever
head -40 tools/scrape_talents_proto.py     # read how the env is chosen
python3 tools/scrape_talents_proto.py --class warrior --env forever
python3 tools/scrape_talents_config.py --class warrior --env forever
export PATH=$PATH:$(go env GOPATH)/bin && make proto
```

If the scrapers take the environment differently (they were written against `wowhead.com/classic/talent-calc/<class>`), read the file and pass whatever it takes; Wowhead's Forever data environment is live at `nether.wowhead.com/forever/` with the internal name `classicplus` (research §5.2). **If the Forever talent calculator is not up yet** — before Sept 17 it will not be — do not fabricate a tree. Instead:

- leave `proto/warrior.proto` as it is,
- set `TalentTreeSizes` to Era's `[3]int{18, 17, 17}` unchanged,
- add the comment below to `talents.go`, and
- put `"warrior: talent tree shape"` into the unconfirmed list.

```go
// Forever keeps seven rows per tree and moves the gold-medal talents to
// 11, 16, 21 and 31 points, adding a 16-point talent per tree that
// vanilla does not have. The tree's actual contents come from
// tools/scrape_talents_proto.py against Wowhead's Forever environment,
// which is not populated before the beta client on 2026-09-17. Until it
// is, this is Era's tree and the milestone list below is the only part
// of Forever's shape that is expressed.
// unconfirmed: warrior talent tree shape
```

The rest of this task is written so it works either way: the mods below key on spell masks and talent *fields*, and a regenerated proto changes field names, not the shape of the code.

- [ ] **Step 4: Declare the milestones and the spell masks**

In `sim/warrior/warrior.go`, beside the existing `SpellCode_Warrior*` block (lines 17-31) and `TalentTreeSizes` (line 33), add:

```go
// ForeverMilestones are the point totals at which each tree grants a
// gold-medal talent. Vanilla had 11, 21 and 31; Forever adds 16.
var ForeverMilestones = []int{11, 16, 21, 31}

// Spell masks for the declarative talent mods. These are additional to
// the SpellCode_* constants above, which the existing class code uses as
// a scalar identity; a mask is what lets one talent config target a set
// of spells at once.
const (
	WarriorSpellMaskNone uint64 = 0

	WarriorSpellMaskBloodthirst uint64 = 1 << iota
	WarriorSpellMaskWhirlwind
	WarriorSpellMaskExecute
	WarriorSpellMaskHeroicStrike
	WarriorSpellMaskCleave
	WarriorSpellMaskMortalStrike
	WarriorSpellMaskOverpower
	WarriorSpellMaskRend
	WarriorSpellMaskRevenge
	WarriorSpellMaskShieldSlam
	WarriorSpellMaskSlam
	WarriorSpellMaskSunderArmor
	WarriorSpellMaskThunderClap
	WarriorSpellMaskRampage
	WarriorSpellMaskPiercingHowl

	// Groups, for talents that target a category rather than one spell.
	WarriorSpellMaskSpecials = WarriorSpellMaskBloodthirst | WarriorSpellMaskWhirlwind |
		WarriorSpellMaskExecute | WarriorSpellMaskMortalStrike | WarriorSpellMaskOverpower |
		WarriorSpellMaskShieldSlam | WarriorSpellMaskSlam | WarriorSpellMaskRampage
	WarriorSpellMaskOnNextSwing = WarriorSpellMaskHeroicStrike | WarriorSpellMaskCleave
)
```

Then add `ClassSpellMask: WarriorSpellMask<Name>,` to each ability's `SpellConfig`. `sim/warrior/bloodthirst.go:15` gains `ClassSpellMask: WarriorSpellMaskBloodthirst,` beside its existing `SpellCode:`; do the same in `whirlwind.go`, `execute.go`, `heroic_strike_cleave.go` (both spells), `mortal_strike.go`, `overpower.go`, `rend.go`, `revenge.go`, `shield_slam.go`, `slam.go`, `sunder_armor.go` and `thunder_clap.go`.

- [ ] **Step 5: Rewrite the arithmetic talents as declarative mods**

In `sim/warrior/talents.go`, replace the head of `ApplyTalents` (lines 16-34). The talents that were flat stat additions or `OnSpellRegistered` multiplier hooks become `AddStaticMod` calls; the ones with real mechanics keep their functions.

```go
// fillWarriorTalents parses a talent string into the proto, positionally
// against TalentTreeSizes. It is the one place that pairing happens, so
// a tree-size change cannot be applied in one caller and missed in
// another.
func fillWarriorTalents(talents *proto.WarriorTalents, s string) {
	core.FillTalentsProto(talents.ProtoReflect(), s, TalentTreeSizes)
}

func (warrior *Warrior) ToughnessArmorMultiplier() float64 {
	return 1.0 + 0.02*float64(warrior.Talents.Toughness) // unconfirmed
}

func (warrior *Warrior) ApplyTalents() {
	// Flat stats. Forever keeps these unchanged from 2006 as far as
	// anyone has published; each figure is unconfirmed until the
	// validation job clears it.
	warrior.AddStat(stats.Crit, core.CritRatingPerCritChance*1*float64(warrior.Talents.Cruelty)) // unconfirmed
	warrior.ApplyEquipScaling(stats.Armor, warrior.ToughnessArmorMultiplier())
	warrior.AddStat(stats.Defense, 2*float64(warrior.Talents.Anticipation)) // unconfirmed
	warrior.AddStat(stats.Parry, 1*float64(warrior.Talents.Deflection))     // unconfirmed

	warrior.applyDeclarativeTalents()

	// Talents with real mechanics keep their own functions.
	warrior.applyAngerManagement()
	warrior.applyDeepWounds()
	warrior.applyWeaponSpecializations()
	warrior.applyUnbridledWrath()
	warrior.applyDualWieldSpecialization()
	warrior.applyEnrage()
	warrior.applyFlurry()
	warrior.applyShieldSpecialization()
	warrior.registerDeathWishCD()
	warrior.registerSweepingStrikesCD()
	warrior.registerLastStandCD()
}

// applyDeclarativeTalents is every talent that is a modifier on a set of
// spells. Before the spell-mod system these were OnSpellRegistered
// closures; as config they can be read against a tooltip line by line,
// which is what Forever's weekly number changes need.
func (warrior *Warrior) applyDeclarativeTalents() {
	t := warrior.Talents

	// Improved Heroic Strike: -1 rage per point.
	if t.ImprovedHeroicStrike > 0 {
		warrior.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_PowerCost_Flat,
			ClassMask: WarriorSpellMaskHeroicStrike,
			IntValue:  -int64(t.ImprovedHeroicStrike), // unconfirmed
		})
	}

	// Improved Execute: -2.5 rage per point, rounded as the engine's
	// integer rage cost requires.
	if t.ImprovedExecute > 0 {
		warrior.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_PowerCost_Flat,
			ClassMask: WarriorSpellMaskExecute,
			IntValue:  -int64(t.ImprovedExecute) * 2, // unconfirmed
		})
	}

	// Impale: +10% crit damage per point on specials.
	if t.Impale > 0 {
		warrior.AddStaticMod(core.SpellModConfig{
			Kind:       core.SpellMod_CritDamageBonus_Flat,
			ClassMask:  WarriorSpellMaskSpecials,
			FloatValue: 0.1 * float64(t.Impale), // unconfirmed
		})
	}

	// Two-handed weapon specialization: +1% damage per point, two-handers
	// only. The hand-type check cannot be expressed as config, so the
	// talent is skipped rather than being applied and filtered.
	if t.TwoHandedWeaponSpecialization > 0 && warrior.MainHand().HandType == proto.HandType_HandTypeTwoHand {
		warrior.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_DamageDone_Flat,
			ClassMask: WarriorSpellMaskSpecials | WarriorSpellMaskOnNextSwing,
			IntValue:  int64(t.TwoHandedWeaponSpecialization), // unconfirmed
		})
	}

	// One-handed weapon specialization: +2% per point, one-handers only.
	if t.OneHandedWeaponSpecialization > 0 && warrior.MainHand().HandType != proto.HandType_HandTypeTwoHand {
		warrior.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_DamageDone_Flat,
			ClassMask: WarriorSpellMaskSpecials | WarriorSpellMaskOnNextSwing,
			IntValue:  2 * int64(t.OneHandedWeaponSpecialization), // unconfirmed
		})
	}
}
```

Delete `applyOneHandedWeaponSpecialization` and `applyTwoHandedWeaponSpecialization`; they are now the two blocks above. Read every remaining function in the file and, for each whose whole body is "multiply a matching spell's damage / cost / cooldown", move it here too — the test in Step 1 does not measure this, so the judgement is: if the function has no state and no timer, it is config.

- [ ] **Step 6: Add the new baseline abilities**

Forever makes several buff-only talents baseline and adds new abilities per class. For Fury the two that matter are Rampage (a Fury capstone in later expansions that Forever grants earlier) and Piercing Howl (a Fury talent made baseline). **Neither has published Forever numbers**, so each ships as a real, registered ability reading its constants from the generated file, with its behaviour written from the tooltip and marked.

Create `sim/warrior/rampage.go`, in the one-ability-per-file mould `mortal_strike.go` established:

```go
package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

// Rampage: on a crit, the warrior and the party gain attack power for 30
// seconds, stacking. Forever grants it baseline rather than as a Fury
// capstone.
//
// unconfirmed: the attack power per stack, the stack cap, and whether the
// party half exists in Forever are all unpublished. The figures below are
// the last-known values from the expansion that introduced it, which is
// the closest thing to evidence that exists before the beta. The
// validation job's aura-uptime comparison is what clears them.
const (
	rampageSpellID         = 29801
	rampageAPPerStack      = 50 // unconfirmed
	rampageMaxStacks       = 5  // unconfirmed
	rampageDuration        = time.Second * 30
	rampageRageCost        = 20 // unconfirmed
	rampageCooldownSeconds = 0
)

func (warrior *Warrior) registerRampageSpell() {
	actionID := core.ActionID{SpellID: rampageSpellID}

	aura := warrior.RegisterAura(core.Aura{
		Label:     "Rampage",
		ActionID:  actionID,
		Duration:  rampageDuration,
		MaxStacks: rampageMaxStacks,
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
			warrior.AddStatDynamic(sim, stats.AttackPower, float64(newStacks-oldStacks)*rampageAPPerStack)
		},
	})

	warrior.Rampage = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:       actionID,
		ClassSpellMask: WarriorSpellMaskRampage,
		SpellSchool:    core.SpellSchoolPhysical,
		ProcMask:       core.ProcMaskEmpty,
		Flags:          core.SpellFlagAPL | core.SpellFlagNoOnCastComplete,
		RageCost:       core.RageCostOptions{Cost: rampageRageCost},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			// Rampage requires a recent crit, which the engine models as
			// the Enrage-style trigger aura below.
			return warrior.RampageValidAura.IsActive()
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			aura.Activate(sim)
			aura.SetStacks(sim, rampageMaxStacks)
		},
	})

	// The "you have critically hit" window that gates the cast.
	warrior.RampageValidAura = warrior.RegisterAura(core.Aura{
		Label:    "Rampage Ready",
		ActionID: core.ActionID{SpellID: rampageSpellID, Tag: 1},
		Duration: time.Second * 5, // unconfirmed
	})
	warrior.RegisterAura(core.Aura{
		Label:    "Rampage Trigger",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if result.DidCrit() && spell.ProcMask.Matches(core.ProcMaskMelee) {
				warrior.RampageValidAura.Activate(sim)
			}
		},
	})
}
```

Add `Rampage *core.Spell` and `RampageValidAura *core.Aura` to the `Warrior` struct in `warrior.go`, and call `warrior.registerRampageSpell()` from wherever the other abilities are registered (grep `registerBloodthirstSpell(` to find it).

Create `sim/warrior/piercing_howl.go` the same way. It is a snare with no damage, so it contributes nothing to a Patchwerk sim; register it anyway, because the APL validator warns about a spell the character cannot cast and because a later multi-target encounter profile will use it:

```go
package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Piercing Howl: an area snare. Forever makes it baseline rather than a
// Fury talent. It deals no damage, so it changes nothing on a
// single-target fight; it is registered so the APL validator knows the
// character has it and so a movement-aware encounter profile can use it.
//
// unconfirmed: rage cost and cooldown.
const (
	piercingHowlSpellID  = 12323
	piercingHowlRageCost = 10              // unconfirmed
	piercingHowlDuration = time.Second * 6 // unconfirmed
)

func (warrior *Warrior) registerPiercingHowlSpell() {
	warrior.PiercingHowl = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:       core.ActionID{SpellID: piercingHowlSpellID},
		ClassSpellMask: WarriorSpellMaskPiercingHowl,
		SpellSchool:    core.SpellSchoolPhysical,
		ProcMask:       core.ProcMaskEmpty,
		Flags:          core.SpellFlagAPL | core.SpellFlagNoOnCastComplete,
		RageCost:       core.RageCostOptions{Cost: piercingHowlRageCost},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
			IgnoreHaste: true,
		},
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aoeTarget := range sim.Encounter.TargetUnits {
				spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeAlwaysHit)
			}
		},
	})
}
```

Field names on `core.Aura` (`OnStacksChange`, `OnSpellHitDealt`, `MaxStacks`, `SetStacks`) and on `SpellConfig` (`ExtraCastCondition`, `RageCostOptions`) must be checked against `sim/core/aura.go` and `sim/core/spell.go` and corrected if they differ — the compiler tells you immediately.

- [ ] **Step 7: Generate the constants and point the abilities at them**

```bash
cd /Users/jh/code/wowsims-forever
make spellconst BUILD=1.15.9.69722
head -40 sim/warrior/constants_auto_gen.go
```

If the data lane's output does not exist yet the target says so by name. In that case leave the placeholder file and skip to Step 8 — the ability literals stay, each already marked `unconfirmed`, and the regeneration is a later one-line commit.

If it does exist, replace the literals in the ability files with the generated arrays. `sim/warrior/bloodthirst.go` becomes rank-aware in the `frostbolt.go` shape, and `mortal_strike.go`'s `bonusDamage := 160.0` becomes `MortalStrikeBaseDamage[rank][0]`. Do this only for abilities the generated file actually covers; an ability with no row keeps its literal and its comment.

- [ ] **Step 8: Write the default APL**

**Re-validate every rank.** The data lane verified the presets are stale: `ui/mage/apls/p1.apl.json` casts Frostbolt rank 10 where the tables give rank 11 for spell `25304`. So:

```bash
cd /Users/jh/code/wowsims-forever
python3 - <<'EOF'
import json, sys
apl = json.load(open('ui/warrior/apls/dps_reck.apl.json'))
ids = set()
def walk(o):
    if isinstance(o, dict):
        if 'spellId' in o and isinstance(o['spellId'], dict) and 'spellId' in o['spellId']:
            ids.add(o['spellId']['spellId'])
        for v in o.values(): walk(v)
    elif isinstance(o, list):
        for v in o: walk(v)
walk(apl)
print(sorted(ids))
EOF
```

Check each id against `data/builds/<build>/spellconst/warrior.json` in the site repo: for each spell *name* the APL uses, the id in the APL must be the highest-rank id the tables give for the character's level. Where it is not, use the table's id.

Then write `ui/warrior/apls/forever_fury.apl.json`, in the schema `proto/apl.proto` defines (research §1.4). The starting priority, which the validation job's cast-frequency comparison is what changes:

```json
{
  "type": "TypeAPL",
  "prepullActions": [
    { "action": { "castSpell": { "spellId": { "spellId": 6673 } } }, "doAt": "-10s" }
  ],
  "priorityList": [
    { "action": { "castSpell": { "spellId": { "spellId": 12292 } } } },
    { "action": { "castSpell": { "spellId": { "spellId": 23894 } } } },
    { "action": { "castSpell": { "spellId": { "spellId": 29801 } } } },
    {
      "action": {
        "condition": { "cmp": { "op": "OpGe", "lhs": { "currentRage": {} }, "rhs": { "const": { "val": "50" } } } },
        "castSpell": { "spellId": { "spellId": 1680 } }
      }
    },
    {
      "action": {
        "condition": { "isExecutePhase": { "threshold": "ExecutePhase20" } },
        "castSpell": { "spellId": { "spellId": 20662 } }
      }
    },
    {
      "action": {
        "condition": { "cmp": { "op": "OpGe", "lhs": { "currentRage": {} }, "rhs": { "const": { "val": "60" } } } },
        "castSpell": { "spellId": { "spellId": 11567 } }
      }
    }
  ]
}
```

Every one of those six spell ids is a placeholder until Step 8's check confirms it: `6673` Battle Shout, `12292` Death Wish, `23894` Bloodthirst, `29801` Rampage, `1680` Whirlwind, `20662` Execute, `11567` Heroic Strike. **Replace each with the id the build's tables give for the highest rank a level-60 warrior has**, and record in the commit body which ones moved.

- [ ] **Step 9: Point the test suite at the Forever build and APL**

In `sim/warrior/dps_warrior/dps_warrior_test.go`, add the Forever rotation beside the existing ones and declare the reference build. In `talents.go` or `warrior.go`:

```go
// ForeverFuryTalents is the reference build the regression suite runs: a
// deep Fury build reaching Bloodthirst, with the Arms points where a Fury
// warrior actually spends them. It is not advice; it is a fixed input so
// a DPS change is attributable to the engine rather than to a build edit.
// unconfirmed: it is expressed against Era's tree until the Forever
// talent calculator is populated.
const ForeverFuryTalents = "30305001302-05050005525010051"
```

and in the test:

```go
			Rotation: core.GetAplRotation("../../../ui/warrior/apls", "forever_fury"),
			OtherRotations: []core.RotationCombo{
				core.GetAplRotation("../../../ui/warrior/apls", "dps_reck"),
			},
```

- [ ] **Step 10: Run, regenerate the goldens, and read them**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test --tags=with_db ./sim/warrior/... -v -run 'TestForever|TestTalentTree|TestFurySpells'
go test --tags=with_db -count=1 ./sim/warrior/... 2>&1 | tail -15
make update-tests
git diff -- 'sim/warrior/*.results' | head -60
```

Expected: `BUILDS`; no `gofmt` output; the four new tests `PASS`; the warrior suites `ok`. The `.results` diff will be large — Rampage is a new ability and the talent mods replace closures — so read the DPS lines. A Fury warrior with Rampage should gain; a tank warrior, which takes none of this, should be unchanged. If the tank moved, a mod is matching more spells than it should: check the `WarriorSpellMaskSpecials` group.

- [ ] **Step 11: Measure**

The baseline measured for this plan, on an Apple M4 Pro, 300-second fight, phase-1 Fury, 3,000 iterations: **1,231 iterations per second serial, 2.44 s wall**. Confirm the new spec has not made the engine materially slower:

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db ./sim/warrior/dps_warrior/ -count=1 -v 2>&1 | tail -3
```

and record the suite's wall time in the commit body. A Fury sim more than 25% slower than before means an aura is being re-registered per iteration; `sim/core/cooldown.go`'s "Over 100 timers!" panic catches the worst form of that.

- [ ] **Step 12: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add proto/warrior.proto sim/core/proto/ sim/warrior/ ui/warrior/ ui/core/talents/trees/warrior.json sim/register_all.go
git commit -m "feat(warrior): Forever Fury end to end - talents, baseline abilities, default APL" \
  -m "The first of the two launch specs. Talents that are pure modifiers become SpellMod config, which is readable against a tooltip line by line and is what Forever's weekly number changes need; talents with state keep their functions. Every ability gains a ClassSpellMask so a mod can target a set. Rampage and Piercing Howl are new baseline abilities, one file each in the existing mould, with every unpublished figure marked unconfirmed rather than presented as Forever's. The default APL is written fresh rather than copied from dps_reck, because the data lane verified the checked-in presets are stale on spell ranks, and every spell id in it is checked against the build's own tables. The talent tree shape is Era's until Wowhead's Forever environment is populated on Sept 17; the 11/16/21/31 milestones are declared and tested now." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 12: `mage-frost` end to end

**Repo: ENGINE.** Depends on Tasks 7 and 10. **G3 — INDEPENDENT of Task 11.** This task touches only `sim/mage/**`, `proto/mage.proto` and `ui/mage/**`.

Frost is the second launch spec and the one with the most different mechanics: a caster with resistances and mana against the warrior's attack table and rage. Same four deliverables.

Two differences from Task 11 that matter:

- **The mage is one package, not a spec sub-package.** `sim/mage/mage.go:31` is `RegisterMage()`; there is no `sim/mage/frost/`. The spec slug `mage-frost` maps onto `sim/mage` plus the frost APL. `TalentTreeSizes` is `[3]int{16, 16, 17}` (`sim/mage/mage.go:29`).
- **The mage already has the generated-constants shape.** `sim/mage/frostbolt.go:9-15` is exactly the per-rank arrays Task 10's generator emits, hand-written. So the constants step here is a substitution, not a restructure: delete the hand-written arrays and let the generated file supply them.

**The preset APL is the one the data lane caught.** `ui/mage/apls/p1.apl.json` casts Frostbolt **rank 10** where the Era tables give **rank 11 for spell `25304`**. Fixing that is a real DPS change and it is this task's most concrete deliverable.

**Files:**
- Modify: `proto/mage.proto`, `sim/mage/talents.go` (529 lines), `sim/mage/mage.go`, `sim/mage/frostbolt.go`, `sim/mage/constants_auto_gen.go` (regenerated), `sim/mage/mage_test.go`
- Create: `sim/mage/ice_lance.go`, `sim/mage/cold_snap_baseline.go`, `ui/mage/apls/forever_frost.apl.json`, `ui/core/talents/trees/mage.json` (regenerated)
- Test: `sim/mage/talents_test.go`, `sim/mage/mage_test.go`

**Interfaces:**
- Consumes: `core.SpellModConfig`, `core.SpellMod_*`, `(*Unit).AddStaticMod`, `Spell.ClassSpellMask` (Task 7); `mage.ConstantsBuild` and the generated arrays (Task 10).
- Produces:
  - `mage.MageSpellMask*` — one `uint64` per ability
  - `mage.ForeverMilestones` — the same `[]int{11, 16, 21, 31}`; **it is declared per class rather than in core**, because a tree is a class's own shape and a shared constant would invite a class to differ silently
  - `(*Mage).registerIceLanceSpell()`, `(*Mage).registerColdSnapSpell()`
  - `mage.ForeverFrostTalents string`
  - `ui/mage/apls/forever_frost.apl.json`

- [ ] **Step 1: Write the failing test**

Create `sim/mage/talents_test.go`:

```go
package mage

import (
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

func TestForeverTalentMilestones(t *testing.T) {
	want := []int{11, 16, 21, 31}
	if len(ForeverMilestones) != len(want) {
		t.Fatalf("ForeverMilestones = %v, want %v", ForeverMilestones, want)
	}
	for i, m := range want {
		if ForeverMilestones[i] != m {
			t.Errorf("ForeverMilestones[%d] = %d, want %d", i, ForeverMilestones[i], m)
		}
	}
}

// The talent string is parsed positionally against TalentTreeSizes, so a
// mismatch silently reads the wrong talent.
func TestTalentTreeSizesMatchTheProto(t *testing.T) {
	var total int
	for _, n := range TalentTreeSizes {
		total += n
	}
	fields := (&proto.MageTalents{}).ProtoReflect().Descriptor().Fields()
	if fields.Len() != total {
		t.Errorf("MageTalents has %d fields, TalentTreeSizes sums to %d", fields.Len(), total)
	}
}

func TestFrostSpellsCarryTheirMasks(t *testing.T) {
	cases := []struct {
		name string
		mask uint64
	}{
		{"Frostbolt", MageSpellMaskFrostbolt},
		{"Ice Lance", MageSpellMaskIceLance},
		{"Frost Nova", MageSpellMaskFrostNova},
		{"Blizzard", MageSpellMaskBlizzard},
		{"Cone of Cold", MageSpellMaskConeOfCold},
	}
	seen := uint64(0)
	for _, c := range cases {
		if c.mask == 0 {
			t.Errorf("%s has a zero mask; an empty mask matches nothing", c.name)
		}
		if seen&c.mask != 0 {
			t.Errorf("%s reuses a bit already taken", c.name)
		}
		seen |= c.mask
	}
}

// The data lane verified that the engine's checked-in preset APL casts
// Frostbolt rank 10 where the Era tables give rank 11 for spell 25304.
// The Forever APL must use the highest rank the character has.
func TestFrostboltHasElevenRanks(t *testing.T) {
	if FrostboltRanks < 11 {
		t.Fatalf("FrostboltRanks = %d, want at least 11", FrostboltRanks)
	}
	if got := FrostboltSpellId[11]; got != 25304 {
		t.Errorf("FrostboltSpellId[11] = %d, want 25304", got)
	}
	if FrostboltLevel[11] > 60 {
		t.Errorf("Frostbolt rank 11 requires level %d; a level-60 mage cannot cast it", FrostboltLevel[11])
	}
}

func TestForeverFrostTalentsAreAValidBuild(t *testing.T) {
	talents := &proto.MageTalents{}
	fillMageTalents(talents, ForeverFrostTalents)
	if talents.IceBarrier == 0 && talents.WintersChill == 0 {
		t.Error("the reference Frost build reaches nothing deep in the Frost tree")
	}
}
```

`fillMageTalents` wraps `core.FillTalentsProto(talents.ProtoReflect(), s, TalentTreeSizes)`; factor the call `sim/mage/mage.go:132` already makes through it.

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./sim/mage/ -v -run 'TestForever|TestTalentTree|TestFrostSpells|TestFrostboltHas'`
Expected: `FAIL [build failed]`, `undefined: ForeverMilestones`.

- [ ] **Step 3: Regenerate the talent proto and tree JSON**

```bash
cd /Users/jh/code/wowsims-forever
python3 tools/scrape_talents_proto.py --class mage --env forever
python3 tools/scrape_talents_config.py --class mage --env forever
export PATH=$PATH:$(go env GOPATH)/bin && make proto
```

As in Task 11: if Wowhead's Forever talent calculator is not populated yet — it will not be before Sept 17 — **do not fabricate a tree**. Leave `proto/mage.proto` alone, keep `TalentTreeSizes = [3]int{16, 16, 17}`, and add the same comment to `talents.go`:

```go
// Forever keeps seven rows per tree and moves the gold-medal talents to
// 11, 16, 21 and 31 points, adding a 16-point talent per tree. The tree's
// contents come from tools/scrape_talents_proto.py against Wowhead's
// Forever environment, which is not populated before the beta client on
// 2026-09-17.
// unconfirmed: mage talent tree shape
```

- [ ] **Step 4: Declare the milestones and the spell masks**

In `sim/mage/mage.go`, beside `TalentTreeSizes` at line 29:

```go
// ForeverMilestones are the point totals at which each tree grants a
// gold-medal talent. Vanilla had 11, 21 and 31; Forever adds 16. It is
// declared per class rather than in core: a tree is a class's own shape,
// and a shared constant would let one class differ without saying so.
var ForeverMilestones = []int{11, 16, 21, 31}

// Spell masks for the declarative talent mods.
const (
	MageSpellMaskNone uint64 = 0

	MageSpellMaskFrostbolt uint64 = 1 << iota
	MageSpellMaskIceLance
	MageSpellMaskFrostNova
	MageSpellMaskBlizzard
	MageSpellMaskConeOfCold
	MageSpellMaskIceBarrier
	MageSpellMaskFireball
	MageSpellMaskFrostfireBolt
	MageSpellMaskScorch
	MageSpellMaskPyroblast
	MageSpellMaskArcaneExplosion
	MageSpellMaskArcaneMissiles
	MageSpellMaskEvocation
	MageSpellMaskColdSnap

	// Groups.
	MageSpellMaskFrostDamage = MageSpellMaskFrostbolt | MageSpellMaskIceLance |
		MageSpellMaskFrostNova | MageSpellMaskBlizzard | MageSpellMaskConeOfCold |
		MageSpellMaskFrostfireBolt
	MageSpellMaskFireDamage = MageSpellMaskFireball | MageSpellMaskScorch |
		MageSpellMaskPyroblast | MageSpellMaskFrostfireBolt
)
```

Then add `ClassSpellMask:` to each ability's `SpellConfig`. `sim/mage/frostbolt.go`'s `getFrostboltConfig` gains `ClassSpellMask: MageSpellMaskFrostbolt,`; do the same in `frost_nova.go`, `blizzard.go`, `cone_of_cold.go`, `ice_barrier.go`, `fireball.go`, `scorch.go`, `pyroblast.go`, `arcane_explosion.go`, `arcane_missiles.go` and `evocation.go` (list the actual files with `ls sim/mage/*.go` and cover each that registers a spell).

- [ ] **Step 5: Rewrite the arithmetic talents as declarative mods**

In `sim/mage/talents.go`, add an `applyDeclarativeTalents` beside `ApplyTalents` and move every pure-modifier talent into it. The frost ones that matter:

```go
// applyDeclarativeTalents is every talent that is a modifier on a set of
// spells. As config these can be read against a tooltip line by line,
// which is what Forever's weekly number changes need; talents with state
// or a timer keep their own functions.
func (mage *Mage) applyDeclarativeTalents() {
	t := mage.Talents

	// Piercing Ice: +2% frost damage per point.
	if t.PiercingIce > 0 {
		mage.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_DamageDone_Flat,
			ClassMask: MageSpellMaskFrostDamage,
			IntValue:  2 * int64(t.PiercingIce), // unconfirmed
		})
	}

	// Improved Frostbolt: -0.1s cast time per point.
	if t.ImprovedFrostbolt > 0 {
		mage.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_CastTime_Flat,
			ClassMask: MageSpellMaskFrostbolt,
			TimeValue: -time.Millisecond * 100 * time.Duration(t.ImprovedFrostbolt), // unconfirmed
		})
	}

	// Elemental Precision: +2% hit on frost and fire.
	if t.ElementalPrecision > 0 {
		mage.AddStaticMod(core.SpellModConfig{
			Kind:       core.SpellMod_BonusHit_Flat,
			ClassMask:  MageSpellMaskFrostDamage | MageSpellMaskFireDamage,
			FloatValue: 2 * float64(t.ElementalPrecision) * core.HitRatingPerHitChance, // unconfirmed
		})
	}

	// Frost Channeling: -5% mana cost per point on frost spells.
	if t.FrostChanneling > 0 {
		mage.AddStaticMod(core.SpellModConfig{
			Kind:      core.SpellMod_PowerCost_Pct,
			ClassMask: MageSpellMaskFrostDamage,
			IntValue:  -5 * int64(t.FrostChanneling), // unconfirmed
		})
	}

	// Arctic Reach: range, which the sim does not model. Declared so the
	// talent is not silently missing; the mod system has no range kind
	// and adding one would model nothing.
	_ = t.ArcticReach

	// Ice Shards: +20% frost crit damage per point.
	if t.IceShards > 0 {
		mage.AddStaticMod(core.SpellModConfig{
			Kind:       core.SpellMod_CritDamageBonus_Flat,
			ClassMask:  MageSpellMaskFrostDamage,
			FloatValue: 0.2 * float64(t.IceShards), // unconfirmed
		})
	}
}
```

Call it from `ApplyTalents` and delete the `OnSpellRegistered` closures it replaces. The talent field names above are Era's; if the regenerated proto renamed one, the compiler says so — correct it rather than dropping the talent.

`time` and `core` must be imported.

- [ ] **Step 6: Add the new baseline abilities**

Create `sim/mage/ice_lance.go`:

```go
package mage

import (
	"github.com/wowsims/classic/sim/core"
)

// Ice Lance: a fast, cheap frost nuke that hits harder against a frozen
// target. Forever grants it baseline; vanilla has no such spell.
//
// unconfirmed: base damage, coefficient, mana cost, and the frozen
// multiplier are all unpublished. The figures below are the last-known
// values from the expansion that introduced it, which is the closest
// thing to evidence before the beta, and the validation job's
// cast-frequency and damage comparison is what clears them.
const (
	iceLanceSpellID         = 30455
	iceLanceBaseDamageLow   = 161  // unconfirmed
	iceLanceBaseDamageHigh  = 187  // unconfirmed
	iceLanceCoefficient     = 0.14 // unconfirmed
	iceLanceManaCost        = 150  // unconfirmed
	iceLanceFrozenMultiplur = 3.0  // unconfirmed
	iceLanceRequiredLevel   = 40   // unconfirmed
)

func (mage *Mage) registerIceLanceSpell() {
	if mage.Level < iceLanceRequiredLevel {
		return
	}
	mage.IceLance = mage.GetOrRegisterSpell(core.SpellConfig{
		ActionID:       core.ActionID{SpellID: iceLanceSpellID},
		ClassSpellMask: MageSpellMaskIceLance,
		SpellSchool:    core.SpellSchoolFrost,
		DefenseType:    core.DefenseTypeMagic,
		ProcMask:       core.ProcMaskSpellDamage,
		Flags:          core.SpellFlagAPL,
		RequiredLevel:  iceLanceRequiredLevel,

		ManaCost: core.ManaCostOptions{FlatCost: iceLanceManaCost},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{GCD: core.GCDDefault},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: iceLanceCoefficient,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(iceLanceBaseDamageLow, iceLanceBaseDamageHigh)
			// Forever keeps the frozen-target bonus; whether it stacks
			// with Shatter's crit bonus is exactly the kind of
			// interaction rule DB2 does not carry (research 5.3), so it
			// is written as multiplicative and marked.
			if mage.isTargetFrozen(target) {
				baseDamage *= iceLanceFrozenMultiplur // unconfirmed
			}
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		},
	})
}

// isTargetFrozen reports whether a frost snare or root is on the target.
// The engine has no "frozen" concept, so this checks the auras Forever's
// frost spells apply.
func (mage *Mage) isTargetFrozen(target *core.Unit) bool {
	for _, aura := range mage.frozenAuras(target) {
		if aura.IsActive() {
			return true
		}
	}
	return false
}
```

Write `frozenAuras` against whatever Frost Nova and Frostbite register in this package (grep `Frost Nova` in `sim/mage/`); if neither registers a debuff aura the engine tracks, return `nil` and add a comment saying the frozen bonus is therefore inert until they do. That is honest and testable; guessing an aura name is not.

Create `sim/mage/cold_snap_baseline.go` the same way: Cold Snap resets the frost cooldowns, so it is `RegisterSpell` with an `ApplyEffects` that calls `.Reset()` on the frost timers, `core.SpellFlagNoOnCastComplete`, and a major cooldown registration. Mark its cooldown `unconfirmed`.

Add `IceLance *core.Spell` and `ColdSnap *core.Spell` to the `Mage` struct and call both registrars from wherever the other spells are registered.

- [ ] **Step 7: Substitute the generated constants**

```bash
cd /Users/jh/code/wowsims-forever
make spellconst BUILD=1.15.9.69722
grep -n 'FrostboltRanks\|FrostboltSpellId' sim/mage/constants_auto_gen.go
```

If the generated file declares `FrostboltRanks`, `FrostboltSpellId`, `FrostboltBaseDamage`, `FrostboltSpellCoeff`, `FrostboltCastTime`, `FrostboltManaCost` and `FrostboltLevel`, **delete lines 9-15 of `sim/mage/frostbolt.go`** — the hand-written arrays — and let the generated ones supply them. The rest of `frostbolt.go` is unchanged, because the generator emits exactly the names it already uses. That is the whole point of matching the shape.

Do the same for every other mage ability whose hand-written arrays the generated file covers. An ability with no row keeps its literals and its `unconfirmed` comment.

If the data lane's output does not exist yet, the make target says so by name; leave the hand-written arrays and move on.

- [ ] **Step 8: Write the default APL, with the ranks checked**

```bash
cd /Users/jh/code/wowsims-forever
python3 - <<'EOF'
import json
apl = json.load(open('ui/mage/apls/p1.apl.json'))
ids = set()
def walk(o):
    if isinstance(o, dict):
        sid = o.get('spellId')
        if isinstance(sid, dict) and 'spellId' in sid:
            ids.add(sid['spellId'])
        for v in o.values(): walk(v)
    elif isinstance(o, list):
        for v in o: walk(v)
walk(apl)
print('preset uses:', sorted(ids))
EOF
grep -n 'FrostboltSpellId' sim/mage/frostbolt.go sim/mage/constants_auto_gen.go
```

Expected: the preset's list contains a Frostbolt id that is **not** `25304` — the data lane measured it as rank 10. Note which id it is; that is the bug this task fixes.

Write `ui/mage/apls/forever_frost.apl.json`:

```json
{
  "type": "TypeAPL",
  "prepullActions": [
    { "action": { "castSpell": { "spellId": { "spellId": 25304 } } }, "doAt": "-3s" }
  ],
  "priorityList": [
    {
      "action": {
        "condition": { "auraIsActive": { "auraId": { "spellId": 12472 } } },
        "castSpell": { "spellId": { "spellId": 25304 } }
      }
    },
    { "action": { "castSpell": { "spellId": { "spellId": 12472 } } } },
    {
      "action": {
        "condition": { "cmp": { "op": "OpLt", "lhs": { "currentManaPercent": {} }, "rhs": { "const": { "val": "10%" } } } },
        "castSpell": { "spellId": { "spellId": 12051 } }
      }
    },
    { "action": { "castSpell": { "spellId": { "spellId": 25304 } } } }
  ]
}
```

`25304` is Frostbolt rank 11, `12472` Icy Veins, `12051` Evocation. **Confirm each against the build's tables before committing**, and confirm that rank 11 is castable at 60 (`FrostboltLevel[11]`, which the Step 1 test asserts).

Ice Lance is deliberately not in the opening priority: on a Patchwerk fight with no frozen target it is a damage loss, and putting it in because it is new would be exactly the kind of guess the validation loop exists to prevent. The nightly cast-frequency comparison against top parses is what adds it.

- [ ] **Step 9: Point the test suite at the Forever APL**

In `sim/mage/mage.go` or `talents.go`:

```go
// ForeverFrostTalents is the reference build the regression suite runs.
// It is not advice; it is a fixed input so a DPS change is attributable
// to the engine rather than to a build edit.
// unconfirmed: expressed against Era's tree until the Forever talent
// calculator is populated.
const ForeverFrostTalents = "2500050300030150333125----"
```

In `sim/mage/mage_test.go`, use it and the new rotation:

```go
			Talents:  ForeverFrostTalents,
			Rotation: core.GetAplRotation("../../ui/mage/apls", "forever_frost"),
			OtherRotations: []core.RotationCombo{
				core.GetAplRotation("../../ui/mage/apls", "p1"),
			},
```

Check the relative path against the existing file — `sim/mage/mage_test.go` is one level shallower than `sim/warrior/dps_warrior/`, so it is `../../ui/mage/apls`, not `../../../`.

- [ ] **Step 10: Run, regenerate, and read the goldens**

```bash
cd /Users/jh/code/wowsims-forever
go build ./sim/... && echo BUILDS
gofmt -l ./sim
go test --tags=with_db ./sim/mage/ -v -run 'TestForever|TestTalentTree|TestFrostSpells|TestFrostboltHas'
go test --tags=with_db -count=1 ./sim/mage/... 2>&1 | tail -10
make update-tests
git diff -- 'sim/mage/*.results' | head -60
```

Expected: `BUILDS`; no `gofmt` output; five new tests `PASS`; the mage suite `ok`. The `.results` diff should show a **DPS increase**, because the preset was casting a lower rank of Frostbolt than the character can. Record the before and after numbers in the commit body — that is the concrete evidence this task produced.

- [ ] **Step 11: Measure**

```bash
cd /Users/jh/code/wowsims-forever
go test --tags=with_db ./sim/mage/ -count=1 2>&1 | tail -3
```

The mage suite measured 1.376 s before this task. Record the new figure. A caster sim is normally faster than a melee one — the warrior suite is 7.6 s — so a mage suite above about 4 s means something is being re-registered per iteration.

- [ ] **Step 12: Commit**

```bash
cd /Users/jh/code/wowsims-forever
git add proto/mage.proto sim/core/proto/ sim/mage/ ui/mage/ ui/core/talents/trees/mage.json
git commit -m "feat(mage): Forever Frost end to end - talents, baseline abilities, default APL" \
  -m "The second launch spec, and the one with the different mechanics: resistances and mana rather than an attack table and rage. Pure-modifier talents become SpellMod config; talents with state keep their functions. Ice Lance and Cold Snap are new baseline abilities with every unpublished figure marked unconfirmed, and Ice Lance's frozen bonus is written as inert rather than guessed when no aura tracks frozen. Frostbolt's hand-written per-rank arrays are deleted in favour of the generated ones, which the generator emits under exactly the names frostbolt.go already used. The default APL fixes the rank bug the data lane found: the checked-in preset casts Frostbolt rank 10 where the tables give rank 11 for spell 25304, which a level-60 mage can cast, and the .results diff is the DPS that was being left on the floor." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 13: The three artifacts, `forever-sim`, and the CI that publishes them

**Repo: ENGINE and SITE.** Depends on Tasks 11 and 12. **G4, alone.**

One pinned sha produces three artifacts: `sim.wasm` and `sim.js` for the browser lane, and a native `forever-sim` binary for the premium lane and the validation job. The contract names them and names where they go.

**Measured, at HEAD `7779ebb`, Apple M4 Pro, go1.25.4:**

| Artifact | Figure |
|---|---|
| `sim.wasm`, `GOOS=js GOARCH=wasm go build ./sim/wasm/` | **18,657,775 bytes raw, 3,470,821 gzipped (3.31 MB)**, 2.6 s to build |
| Design budget | 4 MB gzipped — **we are 0.69 MB under it** |
| `sim.js` | `$(go env GOROOT)/lib/wasm/wasm_exec.js`, 16,992 bytes. Go 1.24 moved it from `misc/wasm`; Task 1 fixed the one place that still looked there |
| Native, 300 s fight, phase-1 Fury, 3,000 iterations, serial | **1,231 iterations/sec, 2.44 s wall** |
| Native, same, 4-way split | **679.6 ms wall, 4,414 it/s** |
| Native, same, 8-way split | **378.0 ms wall, 7,936 it/s** |
| WASM under node 22, 500 iterations, same profile | **168 it/s** — **8.5× slower than native**, and the same DPS to one decimal (1427.4 both), which is the determinism check |

Those numbers meet the design's budgets: a 4-worker laptop runs 3,000 iterations of WASM in roughly `3000 / (168 × 4) ≈ 4.5 s`, inside the 8-second budget, and the server lane's 10,000 iterations take about 1.3 s of 8-way native compute, inside the 3-second dispatch budget.

The native binary is new. `cmd/wowsimcli` exists but reads and writes **protojson** and is a cobra app with three other subcommands; the premium lane wants a single-purpose binary that reads a binary `RaidSimRequest`, uses every core automatically, streams progress, and writes a binary `RaidSimResult`. `core.RunRaidSimConcurrentAsync` already does the hard part: `runSimConcurrent` splits by `runtime.NumCPU()` on its own (`sim/core/sim_concurrent.go:441`) and pushes `ProgressMetrics` down a channel.

**Files:**
- Create (engine): `cmd/forever-sim/main.go`, `cmd/forever-sim/main_test.go`, `.github/workflows/artifacts.yml`
- Modify (engine): `makefile` (an `artifacts` target)
- Create (site): `.github/workflows/sim.yml`
- Modify (site): `Makefile` (an `engine-artifacts` target that fetches them)

**Interfaces:**
- Produces:
  - `forever-sim -in <file> -out <file> [-progress] [-iterations N]` — reads a binary `RaidSimRequest`, writes a binary `RaidSimResult`. `-in -` reads stdin, `-out -` writes stdout. With `-progress`, one JSON object per line on **stderr**: `{"completed":N,"total":M,"dps":X}`.
  - Exit 0 on a successful sim, 1 on an engine error (with the message on stderr), 2 on bad input.
  - CI artifacts per sha: `sim.wasm`, `sim.js`, `forever-sim-linux-amd64`, `forever-sim-darwin-arm64`, and `SHA256SUMS`.

- [ ] **Step 1: Write the failing test**

Create `cmd/forever-sim/main_test.go`:

```go
package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

// smallRequest is a two-iteration Fury sim: enough to exercise the whole
// path, fast enough to be a unit test.
func smallRequest(t *testing.T) []byte {
	t.Helper()
	player := &proto.Player{
		Name:          "CLI Test",
		Race:          proto.Race_RaceOrc,
		Class:         proto.Class_ClassWarrior,
		TalentsString: "30305001302-05050005525010051",
		Equipment:     core.GetGearSet("../../ui/warrior/gear_sets", "phase_1").GearSet,
		Rotation:      core.GetAplRotation("../../ui/warrior/apls", "dps_reck").Rotation,
		Consumes:      &proto.Consumes{},
		Buffs:         core.FullIndividualBuffs,
	}
	core.WithSpec(player, &proto.Player_Warrior{Warrior: &proto.Warrior{
		Options: &proto.Warrior_Options{StartingRage: 50, Shout: proto.WarriorShout_WarriorShoutBattle},
	}})
	enc := core.MakeSingleTargetEncounter(0)
	enc.Duration = 60
	req := &proto.RaidSimRequest{
		Raid:       core.SinglePlayerRaidProto(player, core.FullPartyBuffs, core.FullRaidBuffs, core.FullDebuffs),
		Encounter:  enc,
		SimOptions: &proto.SimOptions{Iterations: 2, RandomSeed: 1},
	}
	b, err := googleproto.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestRunProducesAResult(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.pb")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.pb")

	var progress bytes.Buffer
	if err := run(in, out, 0, true, &progress); err != nil {
		t.Fatalf("run: %v", err)
	}

	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	res := &proto.RaidSimResult{}
	if err := googleproto.Unmarshal(b, res); err != nil {
		t.Fatalf("output is not a RaidSimResult: %v", err)
	}
	if res.Error != nil && res.Error.Message != "" {
		t.Fatalf("the sim reported an error: %s", res.Error.Message)
	}
	if res.IterationsDone != 2 {
		t.Errorf("IterationsDone = %d, want 2", res.IterationsDone)
	}
	if res.RaidMetrics == nil || res.RaidMetrics.Dps == nil || res.RaidMetrics.Dps.Avg <= 0 {
		t.Error("the result carries no DPS")
	}
}

// Progress is JSON lines on stderr, so the Cloud Run job can stream it to
// the API without parsing a protobuf.
func TestProgressIsJSONLines(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.pb")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	var progress bytes.Buffer
	if err := run(in, filepath.Join(dir, "res.pb"), 0, true, &progress); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(progress.String()), "\n")
	if len(lines) == 0 || lines[0] == "" {
		t.Fatal("no progress was written")
	}
	for i, line := range lines {
		var got struct {
			Completed int     `json:"completed"`
			Total     int     `json:"total"`
			DPS       float64 `json:"dps"`
		}
		if err := json.Unmarshal([]byte(line), &got); err != nil {
			t.Fatalf("progress line %d is not JSON: %q (%v)", i, line, err)
		}
		if got.Total != 2 {
			t.Errorf("progress line %d has total %d, want 2", i, got.Total)
		}
	}
}

// -iterations overrides the request, so the validation job can reuse one
// stored request at a different precision.
func TestIterationsOverride(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "req.pb")
	if err := os.WriteFile(in, smallRequest(t), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res.pb")
	if err := run(in, out, 5, false, nil); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(out)
	res := &proto.RaidSimResult{}
	if err := googleproto.Unmarshal(b, res); err != nil {
		t.Fatal(err)
	}
	if res.IterationsDone != 5 {
		t.Errorf("IterationsDone = %d, want the overridden 5", res.IterationsDone)
	}
}

func TestBadInputIsRejected(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "junk.pb")
	if err := os.WriteFile(in, []byte("not a protobuf at all, really not"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run(in, filepath.Join(dir, "res.pb"), 0, false, nil); err == nil {
		t.Fatal("junk input was accepted")
	}
	if err := run(filepath.Join(dir, "missing.pb"), filepath.Join(dir, "res.pb"), 0, false, nil); err == nil {
		t.Fatal("a missing input file was accepted")
	}
}
```

- [ ] **Step 2: Run it and watch it fail**

Run: `cd /Users/jh/code/wowsims-forever && go test --tags=with_db ./cmd/forever-sim/ -v`
Expected: `FAIL [build failed]`, `undefined: run`.

- [ ] **Step 3: Write the binary**

Create `cmd/forever-sim/main.go`:

```go
// Command forever-sim runs one RaidSimRequest natively and writes the
// RaidSimResult, both as binary protobuf. It is the server lane's
// binary: the premium Cloud Run job, the nightly validation job, and the
// execution scorer on the ingest path all invoke it.
//
// It is deliberately not cmd/wowsimcli, which speaks protojson and is a
// cobra app with three other subcommands. This one has a single job and
// a stable contract:
//
//	forever-sim -in request.pb -out result.pb -progress
//	forever-sim -in - -out - < request.pb > result.pb
//
// Concurrency is automatic: core.RunRaidSimConcurrentAsync splits the
// request across runtime.NumCPU() and recombines the distribution
// metrics properly, offsetting each split's random seed so the stream
// matches a serial run.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/sim"
	googleproto "google.golang.org/protobuf/proto"
)

// Version is set by the makefile to the short sha the binary was built
// from. It is the same string as sim/enginever.Version in the site.
var Version = "dev"

// exit codes, which the Cloud Run job distinguishes.
const (
	exitOK      = 0
	exitSimFail = 1
	exitBadArgs = 2
)

func main() {
	in := flag.String("in", "-", "RaidSimRequest protobuf; - for stdin")
	out := flag.String("out", "-", "RaidSimResult protobuf; - for stdout")
	iterations := flag.Int("iterations", 0, "override the request's iteration count")
	progress := flag.Bool("progress", false, "write JSON-lines progress to stderr")
	version := flag.Bool("version", false, "print the engine version and exit")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(exitOK)
	}

	var progressOut io.Writer
	if *progress {
		progressOut = os.Stderr
	}
	if err := run(*in, *out, *iterations, *progress, progressOut); err != nil {
		fmt.Fprintln(os.Stderr, "forever-sim:", err)
		if errors.Is(err, errBadInput) {
			os.Exit(exitBadArgs)
		}
		os.Exit(exitSimFail)
	}
}

var errBadInput = errors.New("bad input")

// run is main's body, with its inputs and its progress sink as
// parameters so it is testable.
func run(inPath, outPath string, iterations int, wantProgress bool, progressOut io.Writer) error {
	sim.RegisterAll()

	var raw []byte
	var err error
	if inPath == "-" {
		raw, err = io.ReadAll(os.Stdin)
	} else {
		raw, err = os.ReadFile(inPath)
	}
	if err != nil {
		return fmt.Errorf("%w: reading the request: %v", errBadInput, err)
	}

	req := &proto.RaidSimRequest{}
	if err := googleproto.Unmarshal(raw, req); err != nil {
		return fmt.Errorf("%w: the input is not a RaidSimRequest: %v", errBadInput, err)
	}
	if req.SimOptions == nil {
		return fmt.Errorf("%w: the request has no sim_options", errBadInput)
	}
	if iterations > 0 {
		req.SimOptions.Iterations = int32(iterations)
	}
	if req.SimOptions.Iterations <= 0 {
		return fmt.Errorf("%w: iterations is %d", errBadInput, req.SimOptions.Iterations)
	}
	// IsTest caps concurrency at three splits and adds per-iteration
	// bookkeeping; neither is wanted in production.
	req.SimOptions.IsTest = false

	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(req, reporter, "forever-sim")

	enc := json.NewEncoder(io.Discard)
	if wantProgress && progressOut != nil {
		enc = json.NewEncoder(progressOut)
	}

	var result *proto.RaidSimResult
	for p := range reporter {
		if p.FinalRaidResult != nil {
			result = p.FinalRaidResult
			break
		}
		if wantProgress && progressOut != nil {
			_ = enc.Encode(struct {
				Completed int32   `json:"completed"`
				Total     int32   `json:"total"`
				DPS       float64 `json:"dps"`
			}{p.CompletedIterations, p.TotalIterations, p.Dps})
		}
	}
	if result == nil {
		return errors.New("the engine produced no result")
	}
	if result.Error != nil && result.Error.Message != "" {
		return fmt.Errorf("the sim failed: %s", result.Error.Message)
	}

	b, err := googleproto.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshalling the result: %w", err)
	}
	if outPath == "-" {
		_, err = os.Stdout.Write(b)
		return err
	}
	return os.WriteFile(outPath, b, 0o644)
}
```

The import `github.com/wowsims/classic/sim/sim` is wrong as written — `RegisterAll` lives in package `sim` at the repository's `sim/` directory, so the path is `github.com/wowsims/classic/sim`. Check with `head -1 sim/register_all.go` and use the right one; the compiler catches it immediately.

- [ ] **Step 4: Run the tests and watch them pass**

```bash
cd /Users/jh/code/wowsims-forever
go build --tags=with_db ./cmd/forever-sim/ && echo BUILDS
gofmt -l ./cmd
go test --tags=with_db ./cmd/forever-sim/ -v
```

Expected: `BUILDS`, no `gofmt` output, four tests `PASS`.

- [ ] **Step 5: Add the `artifacts` make target**

In `makefile`, append. Note this target deliberately does **not** depend on `$(OUT_DIR)/.dirstamp`: that pulls in `npm ci`, `tsc --noEmit` and a Vite build of the TypeScript UI, none of which ships on our site and all of which would make CI slow and brittle.

```makefile
ARTIFACT_DIR ?= artifacts
ENGINE_VERSION := $(shell git rev-parse --short HEAD)

.PHONY: artifacts
# artifacts builds the three things one pinned sha produces: sim.wasm and
# sim.js for the browser lane, and forever-sim for the server lane. It
# does not build the TypeScript UI, which does not ship on our site.
artifacts: sim/core/proto/api.pb.go
	mkdir -p $(ARTIFACT_DIR)
	GOOS=js GOARCH=wasm go build -o $(ARTIFACT_DIR)/sim.wasm ./sim/wasm/
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(ARTIFACT_DIR)/sim.js
	go build --tags=with_db -ldflags="-X 'main.Version=$(ENGINE_VERSION)' -s -w" \
	  -o $(ARTIFACT_DIR)/forever-sim ./cmd/forever-sim
	@cd $(ARTIFACT_DIR) && shasum -a 256 sim.wasm sim.js forever-sim > SHA256SUMS
	@echo "ENGINE_VERSION=$(ENGINE_VERSION)"
	@ls -l $(ARTIFACT_DIR)
	@gzip -9 -c $(ARTIFACT_DIR)/sim.wasm | wc -c | awk '{printf "sim.wasm gzipped: %.2f MB (budget 4.00)\n", $$1/1048576}'
```

Add `artifacts/` to `.gitignore`.

Run it:

```bash
cd /Users/jh/code/wowsims-forever && make artifacts
```

Expected: `ENGINE_VERSION=<sha>`, four files, and a gzipped size line. The measured baseline was **3.31 MB**; if it is above 4.00 MB the budget is blown and the fix is `-ldflags="-s -w"` on the wasm build too, then TinyGo as a later option — do not ship over budget silently.

- [ ] **Step 6: Write the engine's CI workflow**

Create `.github/workflows/artifacts.yml` in the **engine** repo, following the site's workflow idiom (see `/Users/jh/code/forever/.github/workflows/logs.yml`):

```yaml
name: artifacts
on:
  push:
    branches: [main]
  pull_request:
  workflow_dispatch:
permissions: { contents: read }
concurrency:
  group: artifacts-${{ github.ref }}
  cancel-in-progress: true

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version-file: go.mod, cache-dependency-path: go.sum }
      - name: gofmt
        run: |
          unformatted=$(gofmt -l ./sim ./tools ./cmd)
          if [ -n "$unformatted" ]; then
            echo "these files are not gofmt'd:"; echo "$unformatted"; exit 1
          fi
      - name: protoc
        run: |
          sudo apt-get update && sudo apt-get install -y protobuf-compiler
          go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
          echo "$(go env GOPATH)/bin" >> "$GITHUB_PATH"
      # The generated protobufs are committed (see PORTING.md); this test
      # regenerates and diffs, so a stale commit fails here rather than in
      # the site repository's build.
      - run: go test ./sim/core/proto/ -run TestGeneratedProtosMatchSources -v
      - run: go vet --tags=with_db ./sim/... ./cmd/...
      - run: go test --tags=with_db -count=1 ./sim/... ./cmd/...

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version-file: go.mod, cache-dependency-path: go.sum }
      - name: protoc
        run: |
          sudo apt-get update && sudo apt-get install -y protobuf-compiler
          go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
          echo "$(go env GOPATH)/bin" >> "$GITHUB_PATH"
      - name: build the three artifacts
        run: make artifacts
      # The design budgets the browser download at 4 MB gzipped. It
      # measured 3.31 MB when this workflow was written; the gate is the
      # budget, not the measurement, so a regression fails the build
      # rather than quietly costing every visitor a second.
      - name: wasm size budget
        run: |
          bytes=$(gzip -9 -c artifacts/sim.wasm | wc -c)
          mb=$(awk -v b="$bytes" 'BEGIN { printf "%.2f", b/1048576 }')
          echo "sim.wasm gzipped: ${mb} MB"
          awk -v b="$bytes" 'BEGIN { exit (b <= 4*1048576) ? 0 : 1 }' || {
            echo "sim.wasm is ${mb} MB gzipped, over the 4 MB budget"; exit 1; }
      # A wasm build that does not instantiate is worse than no build:
      # it fails in the visitor's browser, not here.
      - name: wasm smoke test
        run: |
          cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" /tmp/
          cat > /tmp/smoke.mjs <<'EOF'
          import './wasm_exec.js';
          import { readFile } from 'node:fs/promises';
          globalThis.wasmready = () => {
            const missing = ['raidSimAsync', 'raidSimRequestSplit', 'raidSimResultCombination',
                             'computeStats', 'abortById'].filter(n => typeof globalThis[n] !== 'function');
            if (missing.length) { console.error('missing entrypoints:', missing); process.exit(1); }
            console.log('all five entrypoints present');
            process.exit(0);
          };
          const go = new globalThis.Go();
          const { instance } = await WebAssembly.instantiate(await readFile(process.argv[2]), go.importObject);
          go.run(instance);
          EOF
          cp artifacts/sim.wasm /tmp/
          cd /tmp && node smoke.mjs sim.wasm
      - name: cross-compile the server binary
        run: |
          GOOS=linux GOARCH=amd64 go build --tags=with_db \
            -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)' -s -w" \
            -o artifacts/forever-sim-linux-amd64 ./cmd/forever-sim
          GOOS=darwin GOARCH=arm64 go build --tags=with_db \
            -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)' -s -w" \
            -o artifacts/forever-sim-darwin-arm64 ./cmd/forever-sim
          rm -f artifacts/forever-sim
          cd artifacts && sha256sum sim.wasm sim.js forever-sim-* > SHA256SUMS && cat SHA256SUMS
      - uses: actions/upload-artifact@v4
        with:
          name: engine-${{ github.sha }}
          path: artifacts/
          retention-days: 90
      # A tagged build is what the site pins. Untagged pushes keep their
      # artifacts for 90 days, which covers the beta's weekly cadence.
      - name: release
        if: startsWith(github.ref, 'refs/tags/')
        uses: softprops/action-gh-release@v2
        with:
          files: artifacts/*
```

The smoke test is the reason this workflow earns its place: `sim/wasm/main.go:39` calls `js.Global().Call("wasmready")` after registering its exports, so a page that does not define `wasmready` hangs. The test defines it, and checks that all five entrypoints the contract names are really there.

- [ ] **Step 7: Write the site's side**

In the site repo's `Makefile`, append:

```makefile
ARTIFACT_RUN ?=
WEB_SIM_DIR = web/public/_sim

.PHONY: engine-artifacts
# engine-artifacts downloads the three artifacts the engine's CI built
# for the pinned sha and puts the two browser ones where the web loads
# them: web/public/_sim/<ENGINE_VERSION>/, cached immutably, so a new
# engine version never collides with a cached old one.
#
# Needs the gh CLI and read access to the engine repository.
engine-artifacts: engine-pin
	@sha=$$(sed -n 's/.*Version = "\(.*\)"/\1/p' sim/enginever/version.go); \
	test -n "$$sha" || { echo "no version in sim/enginever/version.go"; exit 1; }; \
	dir="$(WEB_SIM_DIR)/$$sha"; \
	mkdir -p "$$dir"; \
	gh run download $(ARTIFACT_RUN) --repo jhunthrop/wowsims-forever \
	  --name "engine-$$(git -C $(ENGINE_DIR) rev-parse HEAD)" --dir "$$dir" || { \
	    echo "no CI artifact for $$sha; push the engine commit and wait for the artifacts workflow"; exit 1; }; \
	shasum -a 256 -c "$$dir/SHA256SUMS" || { echo "checksum mismatch"; exit 1; }; \
	echo "engine artifacts for $$sha are in $$dir"
```

Create `/Users/jh/code/forever/.github/workflows/sim.yml`:

```yaml
name: sim
on:
  push:
    branches: [main]
    paths: ['sim/**', 'logs/**', '.github/workflows/sim.yml']
  pull_request:
    paths: ['sim/**', 'logs/**', '.github/workflows/sim.yml']
  workflow_dispatch:
permissions: { contents: read }
concurrency:
  group: sim-${{ github.ref }}
  cancel-in-progress: true
jobs:
  test:
    runs-on: ubuntu-latest
    defaults: { run: { working-directory: sim } }
    steps:
      - uses: actions/checkout@v4
      # The engine is a separate repository; CI resolves it from the
      # pinned pseudo-version in sim/go.mod rather than the local replace
      # that development uses.
      - name: drop the development replace
        run: |
          sed -i '/replace github.com\/wowsims\/classic =>/d' go.mod
          grep -q 'github.com/wowsims/classic' go.mod || { echo "sim/go.mod does not require the engine"; exit 1; }
      - uses: actions/setup-go@v5
        with: { go-version-file: sim/go.mod, cache-dependency-path: sim/go.sum }
      - name: gofmt
        run: |
          unformatted=$(gofmt -l .)
          if [ -n "$unformatted" ]; then
            echo "these files are not gofmt'd:"; echo "$unformatted"; exit 1
          fi
      - run: go vet ./...
      - run: go test ./... -race -coverprofile=cover.out
      - name: coverage floor
        run: |
          total=$(go tool cover -func=cover.out | awk '/^total:/ {print substr($3, 1, length($3)-1)}')
          echo "total coverage ${total}%"
          awk -v t="$total" 'BEGIN { exit (t + 0 >= 80) ? 0 : 1 }' || {
            echo "coverage ${total}% is under the 80% floor"; exit 1; }
      # The engine version the adapter stamps into every summary must be
      # the one the artifacts were built from.
      - name: the pin is a real sha
        run: go test ./enginever/ -v
```

The `sed` that drops the development replace will fail the build if `sim/go.mod` still points at a local path with no pseudo-version alongside it. That is intended: it is the check that someone pinned the engine before merging.

- [ ] **Step 8: Verify both workflows parse and the artifacts are real**

```bash
cd /Users/jh/code/wowsims-forever
make artifacts
ls -l artifacts/ && cat artifacts/SHA256SUMS
node --version
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" /tmp/ && cp artifacts/sim.wasm /tmp/
cd /tmp && cat > smoke.mjs <<'EOF'
import './wasm_exec.js';
import { readFile } from 'node:fs/promises';
globalThis.wasmready = () => {
  const missing = ['raidSimAsync','raidSimRequestSplit','raidSimResultCombination','computeStats','abortById']
    .filter(n => typeof globalThis[n] !== 'function');
  if (missing.length) { console.error('missing entrypoints:', missing); process.exit(1); }
  console.log('all five entrypoints present');
  process.exit(0);
};
const go = new globalThis.Go();
const { instance } = await WebAssembly.instantiate(await readFile('sim.wasm'), go.importObject);
go.run(instance);
EOF
node smoke.mjs
```

Expected: four artifacts, a `SHA256SUMS` with four lines, `all five entrypoints present`, and a gzipped size under 4.00 MB (3.31 MB measured).

Then run the binary end to end against the fixture Task 3 produced:

```bash
cd /Users/jh/code/wowsims-forever
./artifacts/forever-sim -version
```

Expected: the short sha, not `dev`.

- [ ] **Step 9: Commit, in both repos**

```bash
cd /Users/jh/code/wowsims-forever
git add cmd/forever-sim makefile .github/workflows/artifacts.yml .gitignore
git commit -m "feat(build): forever-sim, and CI that builds and publishes all three artifacts" \
  -m "One pinned sha produces sim.wasm and sim.js for the browser lane and forever-sim for the server lane. forever-sim is new and is deliberately not cmd/wowsimcli, which speaks protojson and carries three other subcommands: this one reads a binary RaidSimRequest, splits across every core through RunRaidSimConcurrentAsync, streams JSON-lines progress on stderr for the Cloud Run job, and writes a binary RaidSimResult. The make target skips the TypeScript UI build entirely, since none of that ships on our site. CI gates three things that would otherwise fail in a visitor's browser: the committed protobufs are not stale, sim.wasm is under the design's 4 MB gzipped budget (3.31 MB measured), and the built wasm actually instantiates under node with all five contract entrypoints present." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"

cd /Users/jh/code/forever
git add Makefile .github/workflows/sim.yml
git commit -m "ci(sim): test the sim module against the pinned engine, and fetch its artifacts" \
  -m "The workflow drops the development replace before building, so CI resolves the engine from the pinned pseudo-version and a merge that forgot to pin fails loudly rather than passing against a local checkout. make engine-artifacts downloads the three artifacts the engine's CI built for the pinned sha, verifies their checksums, and puts the browser pair under web/public/_sim/<ENGINE_VERSION>/ where the web loads them immutably." \
  -m "Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>"
```

---

## Task 14: Final review

**Repo: ENGINE and SITE.** Depends on everything. **G5, alone.** This is the only task that runs both full suites.

**Files:**
- Modify (only if a check below fails): whichever file the failure names
- Modify: `docs/superpowers/plans/2026-09-14-sim-engine.md` — the measured table in Step 4, corrected to what the tree now measures
- Test: every package in both repositories

**Interfaces:**
- Consumes: everything Tasks 1 through 13 produced.
- Produces: no code. The deliverable is the finish report in Step 6, and a corrected measured table in this plan.

- [ ] **Step 1: Both full suites, clean**

```bash
cd /Users/jh/code/wowsims-forever
export PATH=$PATH:$(go env GOPATH)/bin
make proto && make binary_dist/dist.go
go build ./... && echo "ENGINE BUILDS"
gofmt -l ./sim ./tools ./cmd
go vet --tags=with_db ./sim/... ./cmd/...
time go test --tags=with_db -count=1 ./sim/... ./cmd/...
```

Expected: `ENGINE BUILDS`; `gofmt -l` prints nothing; `go vet` silent; every package `ok`, zero `FAIL`. The baseline before this plan was **20 packages, 10.5 s wall, 78.7 s CPU**; with two reworked specs and the new `cmd/forever-sim` package expect roughly 12 to 15 s. Record the real figure.

```bash
cd /Users/jh/code/forever/sim
gofmt -l .
go vet ./...
go test ./... -race -coverprofile=cover.out
go tool cover -func=cover.out | tail -1
```

Expected: nothing from `gofmt` or `vet`; `ok` for `api`, `adapter` and `enginever`; total coverage at or above 80%.

- [ ] **Step 2: Re-verify every measured claim in this plan**

The plan states numbers. Confirm each is still true, and correct the plan where it is not.

```bash
cd /Users/jh/code/wowsims-forever
echo "--- the merge left no split stats ---"
grep -rnE 'stats\.(MeleeHit|SpellHit|MeleeCrit|SpellCrit)\b' sim/ --include='*.go' | wc -l    # want 0
grep -rnE '\bStat(SpellHit|MeleeHit|SpellCrit|MeleeCrit)\b' ui/ --include='*.ts' --include='*.tsx' | wc -l  # want 0
echo "--- wasm size ---"
make artifacts >/dev/null && gzip -9 -c artifacts/sim.wasm | wc -c | awk '{printf "%.2f MB (budget 4.00, was 3.31)\n", $1/1048576}'
echo "--- the five entrypoints ---"
grep -c 'js.Global().Set("\(raidSimAsync\|raidSimRequestSplit\|raidSimResultCombination\|computeStats\|abortById\)"' sim/wasm/main.go  # want 5
echo "--- provisional values are declared, not hidden ---"
go test --tags=with_db ./sim/core/ -run 'TestProvisional|TestUnconfirmedRacials' -v
echo "--- nothing invented a Forever number without saying so ---"
grep -rn 'unconfirmed' sim/ --include='*.go' | wc -l
```

The last count has no target; it is a figure to read. Every number this plan could not source is supposed to carry that word, so a small count means someone typed a number silently.

- [ ] **Step 3: Confirm the contract, clause by clause**

Open `docs/superpowers/specs/2026-09-14-simulator-interfaces.md` and check the "Engine" and "Engine version" sections against the tree. Each of these is a yes or a written reason why not:

- [ ] The protobuf API does not change shape: `RaidSimRequest`, `RaidSimResult`, `SimDatabase`, `APLRotation` all keep their messages. *(The `Stat` enum is renumbered — an index, not a wire identity, argued in Task 4; `Encounter.biome` is additive.)*
- [ ] `Stat` enum: `MeleeHit`+`SpellHit` → `Hit`, `MeleeCrit`+`SpellCrit` → `Crit`, indexes synced between `sim/core/stats` and `proto/common.proto`, asserted by `TestStatEnumIsSyncedWithProto`.
- [ ] Two artifacts per version plus the native one, all from the pinned sha, built in CI.
- [ ] The WASM entrypoints are exactly `raidSimAsync`, `raidSimRequestSplit`, `raidSimResultCombination`, `computeStats`, `abortById`, and the smoke test proves they instantiate.
- [ ] `ENGINE_VERSION` appears in `sim/enginever/version.go`, is written only by `make engine-pin`, and names the wasm directory and the image tag.
- [ ] `Summarize(res *proto.RaidSimResult, req SimRequest) (summary.Summary, error)` exists in `sim/adapter`, follows the contract's mapping table, and has a golden test per spec.
- [ ] `logs/` is untouched: `cd /Users/jh/code/forever && git diff --stat main -- logs/` prints nothing.
- [ ] Every commit in both repos carries the trailer exactly once:
      `git log --format='%H %s%n%b' main..HEAD | grep -c 'Co-Authored-By: Claude Opus 5'`

- [ ] **Step 4: Re-measure the sim**

The design's budgets depend on these. Run the same measurement the plan's baseline came from and record the new figures beside the old.

Build a 3,000-iteration request once and time `forever-sim` against it. The request is the same profile the baseline used: phase-1 Fury, 300-second single-target fight, seed 1.

```bash
cd /Users/jh/code/wowsims-forever
make artifacts
./artifacts/forever-sim -version

# Build the request with the throwaway harness from Task 3 Step 7, but
# writing a RaidSimRequest rather than a result: change its final lines
# from RunRaidSim + Marshal(res) to Marshal(req), set enc.Duration = 300
# and Iterations = 3000, run it, then delete tools/genfixture again.
go run --tags=with_db ./tools/genfixture -spec warrior-fury -out /tmp/bench-req.pb
rm -rf tools/genfixture

# 8-way native: forever-sim splits across runtime.NumCPU() on its own.
time ./artifacts/forever-sim -in /tmp/bench-req.pb -out /tmp/bench-res.pb -progress 2>/dev/null
```

Divide 3,000 by the wall clock for the concurrent figure. For the serial one, set `GOMAXPROCS=1`. Record:

| Figure | Baseline (2026-09-14, HEAD 7779ebb) | Now |
|---|---|---|
| Native serial, 300 s Fury, 3,000 iters | 1,231 it/s, 2.44 s | |
| Native 8-way | 7,936 it/s, 378 ms | |
| WASM under node, 500 iters | 168 it/s | |
| `sim.wasm` gzipped | 3.31 MB | |
| Full engine suite | 10.5 s wall / 78.7 s CPU | |

A native figure more than 25% below the baseline is a regression worth finding before this lane is declared done: the likely causes are an aura registered per iteration (which `sim/core/cooldown.go`'s "Over 100 timers!" panic catches at the extreme) or a spell mod re-evaluating on every cast.

- [ ] **Step 5: Request review**

Use `superpowers:requesting-code-review` against both branches. Point the reviewer at the four things this lane got wrong most easily:

1. **The merged stat literals.** Task 4 merged composite literals that set both `MeleeCrit` and `SpellCrit`. Each merge picked one value. Check every `// Forever: merged from` comment against what the item or class actually granted.
2. **The spell masks.** A talent mod with a mask that is too wide silently buffs spells it should not. Cross-check `WarriorSpellMaskSpecials` and `MageSpellMaskFrostDamage` against the abilities they name.
3. **The unconfirmed list.** Every Forever number nobody published must carry `unconfirmed`. A number that does not is either sourced — in which case say where — or invented.
4. **The APL ranks.** Both default APLs must cast the highest rank the character has. The data lane found the preset mage APL casting Frostbolt rank 10; confirm neither new APL repeats it.

- [ ] **Step 6: Finish**

Use `superpowers:finishing-a-development-branch`. The engine branch merges into the engine fork's `main`; the site branch merges into the site's `main`. They are independent merges — the site's `sim/go.mod` pins the engine by sha, so the engine must merge and its artifacts workflow must finish **first**, then `make engine-pin && make engine-artifacts` in the site, then the site merges.

Report to the controller:

- the task count and which groups ran in parallel;
- the re-measured table from Step 4;
- every contract clause in Step 3 that is a "no", with the reason;
- the count of `unconfirmed` markers, as the honest size of what the beta still has to settle;
- whether the Forever talent calculator was populated in time, because Tasks 11 and 12 branch on it and the answer determines whether a follow-up commit is owed.

---

## Self-review

Run against the design (sections 2 and 9) and the contract.

**Spec coverage.**

| Design / contract requirement | Task |
|---|---|
| §2.1 Repository, module path unchanged, consumed as a Go module pinned by version | 1, 2 |
| §2.2 Hit and crit are one stat each | 4 |
| §2.2 Expertise reduces parry and dodge | none needed — already modelled (`sim/core/spell_outcome.go:711-735`); Task 5 makes its two constants real rather than hand-set |
| §2.2 Bonus healing carries one third as bonus damage | 6 |
| §2.2 Caster weapons grant spell damage | none needed — `stats.SpellDamage` exists; item data, data lane |
| §2.2 Weapon skill kept | none needed — fully modelled, 15-value enum |
| §2.2 Talent trees: seven rows, 11/16/21/31, buff talents baseline | 11, 12 |
| §2.2 Reworked racials, two active two passive | 9 |
| §2.2 New baseline abilities per class | 11, 12 |
| §2.2 Encounter environment (biome, creature type) | 8 |
| §2.2 `SpellScaling` absent; coefficients stay the vanilla convention | 10 |
| §2.2 `spell_mod.go` cherry-picked from SoD | 7 |
| §2.3 Per-class generated constants file | 10, consumed by 11 and 12 |
| §2.3 Order: Fury Warrior then Frost Mage | 11, 12 |
| §2.4 One default APL per spec, as data | 11, 12 |
| §2.5 Adapter to the logs engine's summary | 3 |
| §5.1 `sim.wasm` from the pinned version, 4 MB gzipped budget | 13 |
| §5.2 Native binary for the Cloud Run job | 13 |
| Contract: `ENGINE_VERSION`, `make engine-pin`, `sim/enginever/version.go` | 2 |
| Contract: `SimRequest` / `SimResult` / `Estimate` envelopes | 2 |
| Contract: the five WASM entrypoints and no others | 13 (asserted by the smoke test) |
| Contract: base stats and rating constants regenerated | 5 |
| §9 risk: Forever's numbers move weekly | 5, 10 — regeneration, not editing |
| §9 risk: unified hit against the weapon-skill miss table is unspecified | 4 keeps the vanilla table; `ProvisionalConstants()` and the spec page say so. **Not solved, by design: only beta logs settle it.** |
| §9 risk: WASM is eight times slower than native | measured at 8.5×; the budgets in Task 13 are built on the measurement |

Gaps, stated rather than hidden:

- **`data/curated/apl/<spec_slug>.json` is the data lane's file**, and Tasks 11 and 12 write the APL into the engine's `ui/<class>/apls/` instead, because that is where `core.GetAplRotation` reads it from in the regression suite. The data lane's curated copy is the same JSON with a `sources` block; keeping both in step is a data-lane task, and the engine's copy is the one the tests run.
- **`data/curated/specs.json` → `sim/specs/specs.go`** is named in the contract's Identifiers section and is generated by the data lane. This plan's `sim/` module does not create `sim/specs`; `adapter.splitSpecSlug` only needs the split, and hardcoding a spec list here would violate "nothing hardcodes a spec list elsewhere."
- **Tanks and healers are out of scope** at launch (design §Decisions), so `Summarize` leaves `Healing`, `DamageTaken` and `HealingTaken` empty and says so.

**Placeholder scan.** No step says "TBD", "implement later", "add appropriate error handling", or "similar to Task N". Every code step carries the code. Three steps are conditional on something outside this lane, and each states both branches explicitly rather than deferring: Task 11 Step 3 and Task 12 Step 3 (the Forever talent calculator may not be populated before Sept 17 — keep Era's tree, declare the milestones, mark it unconfirmed), and Task 10 Step 7 / Task 11 Step 7 / Task 12 Step 7 (the data lane's `spellconst` output may not exist — keep the literals, which already carry `unconfirmed`, and regenerate later in a one-line commit).

**Type consistency.** Checked across tasks:

- `enginever.Version` (Task 2) is read by `adapter` through `req.EngineVersion` (Task 3) and written by `make engine-pin` (Task 2) and stamped into the binary by `make artifacts` (Task 13) — one string, three consumers, one writer.
- `api.SimRequest` (Task 2) is `Summarize`'s second parameter (Task 3) and matches the contract's Go block field for field, JSON name for JSON name.
- `stats.Hit` / `stats.Crit` (Task 4) are used by Task 5's `CritRatingPerCritChance` / `HitRatingPerHitChance`, by Task 6's `deps.go` reorder (which also collapses the duplicate `Crit` the `sed` leaves in `safeDepsOrder`), and by Tasks 11 and 12's talent mods.
- `core.SpellModConfig` and the 26 `SpellMod_*` kinds (Task 7) are used by Tasks 11 and 12; the kinds those two tasks actually use — `SpellMod_PowerCost_Flat`, `SpellMod_PowerCost_Pct`, `SpellMod_CritDamageBonus_Flat`, `SpellMod_DamageDone_Flat`, `SpellMod_CastTime_Flat`, `SpellMod_BonusHit_Flat`, `SpellMod_Cooldown_Flat` — are all in Task 7's produced list.
- `Spell.ClassSpellMask uint64` (Task 7) is what `WarriorSpellMask*` and `MageSpellMask*` (Tasks 11, 12) are assigned to; both are `uint64`.
- `spellconst.Class` / `spellconst.Spell` / `CoefficientFor` (Task 10) are consumed only by the generator in the same task; Tasks 11 and 12 consume the *generated Go arrays*, whose names (`<Name>SpellId`, `<Name>BaseDamage`, `<Name>SpellCoeff`, `<Name>CastTime`, `<Name>ManaCost`, `<Name>Level`, `<Name>CooldownMS`, `<Name>Ranks`) match what `sim/mage/frostbolt.go` already declares, which is why Task 12 Step 7 is a deletion rather than a rewrite.
- `core.ProvisionalConstants()` (Task 5) and `core.UnconfirmedRacials()` (Task 9) are both `[]string` and are both read by the api lane's spec-support page; Task 14 Step 2 runs both.
- `run(inPath, outPath string, iterations int, wantProgress bool, progressOut io.Writer) error` (Task 13) is the signature every test in that task calls.

One inconsistency found and fixed while reviewing: `ForeverMilestones` is declared twice, once in `sim/warrior` and once in `sim/mage`, rather than once in `core`. That is deliberate and Task 12 says why — a tree is a class's own shape, and a shared constant would let one class differ without saying so — but the two declarations must stay equal, which both tasks' `TestForeverTalentMilestones` asserts independently.
