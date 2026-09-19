# Simulator parity — `sim` module lane implementation plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make the site's `sim` Go module speak every request kind Raidbots parity needs — bulk combination runs (Top Gear, Droptimizer, talent compare), stat weights, Smart Sim precision, the full encounter and cooldown vocabulary, and a sample-iteration cast log — so the browser and the server lanes run the same planner, the same statistics and the same mapping, written once in Go.

**Architecture:** Three layers, all inside `sim/`. `sim/api` gains the envelope's new shapes and every rule a client could get wrong (kinds, bulk spec, precision ladders, caps, target error). `sim/bulk` is a new package that owns *which* sims run — expansion, staging, ranking — and is compiled into both artifacts. `sim/request`, `sim/adapter`, `sim/internal/simdb`, `sim/cmd/wasm` and `sim/cmd/forever-sim` carry the new fields across the protobuf boundary and expose the planner as three wasm exports and one native loop. Nothing statistical is ever written in TypeScript.

**Tech Stack:** Go (site modules on `go 1.25`, `GOTOOLCHAIN=auto`; the engine fork stays `go 1.23.0`), `google.golang.org/protobuf`, the pinned `github.com/wowsims/classic` fork at `/Users/jh/code/wowsims-forever`, `GOOS=js GOARCH=wasm` for the browser artifact, GNU make, GitHub Actions. **No new module dependency is added by this lane.**

**Spec:** `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` (sections 2, 4.1, 4.2, 4.3, 5.1, 7, 10.1, 11 are this lane). **Binding contract:** `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` — sections 1, 2, 3, 4 and 1.7 are this lane's interface; sections 5, 6, 8 and 9 constrain it. Supporting: `docs/superpowers/specs/2026-09-14-simulator-interfaces.md` (the contract this one amends), `sim/request/IDS.md`, `sim/README.md`.

---

## Global Constraints

Every task's requirements implicitly include this section.

**Toolchain**

- The repository is a Go workspace: `go.work` at the repository root lists `./api ./companion ./logs ./sim` and declares `go 1.25.11`. `sim/go.mod` declares `go 1.25.11`. Export `GOTOOLCHAIN=auto` in every shell so the declared toolchain is fetched rather than refused.
- **The engine fork stays on `go 1.23.0`.** Never edit `go.mod` in `/Users/jh/code/wowsims-forever`, and never raise its language version to compile something of ours. Nothing in this plan writes Go into the fork.
- **`--tags=with_db` is for the fork's own tests only.** Neither `sim.wasm` nor `forever-sim` is ever built with it: that tag loads vanilla's item table and turns on `core.WITH_DB`, which panics at init for any Forever item effect. Both artifacts embed the active build's rows through `sim/internal/simdb`. If you run a test inside the fork, `go test --tags=with_db ./...` from `/Users/jh/code/wowsims-forever`; from `sim/`, never.
- **Never run `make update-tests`.** It regenerates the fork's own expected-results files and would paper over a real engine change.
- `//go:embed` resolves at compile time, so **`make simdb` must have run** before any `go build`, `go vet` or `go test` in `sim/`. Run it once per checkout and again whenever `data/builds/<build>/simdb.bin` changes.
- **gofmt every file you touch**: `gofmt -w <paths>`. CI fails on `gofmt -l .` inside `sim/`.
- **Tests are scoped to the package the task touches**: `go test ./sim/bulk/...` from the worktree root, never `go test ./...`. The full suite and the 80% coverage floor are CI's job (`.github/workflows/sim.yml`).
- `go vet ./sim/<pkg>/...` after any task that adds a file.

**Branch and commits**

- Work in an isolated worktree created with the `superpowers:using-git-worktrees` skill, on branch `sim-parity-module`. Run every command from that worktree's root.
- **One commit per task**, conventional-commit type: `feat(sim): …` for behaviour, `test(sim): …` for a test-only task, `docs(sim): …` for the contract amendment, `ci:`/`chore:` for the pipeline.
- Every commit ends with exactly one trailer line, and it is this one:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  It is the plan's, and it stands. If the session running a task carries an attribution reminder naming a different model, that reminder is not this lane's; use the line above.
- **Never `git stash`** — the stash stack is shared with every other worktree. Set work aside with a temporary WIP commit instead.
- Never amend a commit that already exists; never bypass a git hook.

**The contract is the authority**

- `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` wins over anything you find in the code. Where the code today disagrees, change the code.
- Where the contract is silent or names something that does not exist, **add the name to the contract in its own commit first** (that is the contract's own rule, stated in its preamble) and then code against it. Task 1 is exactly that commit; do not scatter further amendments without one.

**Contract values, verbatim**

- Kinds: `run`, `gear`, `talents`, `drops`, `weights`. Derived, never sent.
- Precisions: `fast`, `normal`, `high`. Ladders: fast `100 → 1000 → 3000`, normal `1000 → 3000`, high `1000 → 10000`.
- `StepIterations = 1000`. `LaneIterationCeiling = {browser: 30000, server: 100000}`. `Caps = {browser: 400, server: 20000}`.
- Lanes are `api.LaneBrowser = "browser"` and `api.LaneServer = "server"`; they already exist.
- Origins: `equipped`, `bag`, `bank`, `search`, `drop:<source-id>`, `set:<name>`.
- Target levels 60–63, default 63 (`api.BossLevel`). Target types: `humanoid | undead | beast | demon | dragonkin | elemental | giant | mechanical | unknown`.
- Movement kinds: `away`, `casting`.
- The engine pin is `sim/enginever/version.go`'s `Version`, **`edc0c8e9a`** before Task 2. Never type a sha anywhere else; `make engine-pin` is the only writer of that file.
- The wasm exports are JSON strings in and out, failures `{"error": "..."}`. After this lane there are seven: `simRun`, `simSplit`, `simCombine`, `simAbort`, `simPlan`, `simRank`, `simWeights`, plus the `simEngineVersion` string global and the `wasmready()` handshake.
- `forever-sim` flags stay `-in -out -out-proto -iterations -progress -version`; exit codes stay `0` ok, `1` sim failure, `2` bad args/input, `130` aborted.

---

## File Structure

**New files**

| File | Responsibility |
| --- | --- |
| `sim/api/bulk.go` | `BulkSpec`, `Candidate`, `TalentLoadout`, `GearSet`, precision ladders, `Caps`, `ErrCapExceeded`, bulk validation. |
| `sim/api/bulk_test.go` | Table tests for bulk validation and the ladders. |
| `sim/api/weights.go` | `WeightsSpec`, `StatWeight`, weights validation. |
| `sim/api/weights_test.go` | Table tests for weights validation. |
| `sim/request/styles.go` | The fight-style table: one style id expands to encounter fields. |
| `sim/request/styles_test.go` | Every style expands; the committed JSON matches. |
| `sim/request/styles.json` | Generated by `go run ./internal/genstyles`; the web lane's fixture. |
| `sim/request/internal_cooldowns.go` → `sim/request/cooldowns.go` | `CooldownSpec` → `proto.Cooldowns`. |
| `sim/request/cooldowns_test.go` | Cooldown id grammar and timings. |
| `sim/request/weights.go` | `BuildWeights`: a weights `SimRequest` → `proto.StatWeightsRequest`. |
| `sim/request/weights_test.go` | Stat id mapping, reference stat, unknown stat. |
| `sim/internal/genstyles/main.go` | Writes `sim/request/styles.json`. |
| `sim/internal/simdb/items.go` | Exported, protobuf-free item rows for `sim/bulk`. |
| `sim/internal/simdb/items_test.go` | Row shape, slot vocabulary, the fields Expand needs. |
| `sim/bulk/doc.go` | Package comment: what decides which sims run. |
| `sim/bulk/enchants.go` | The build's `enchants.json` table. |
| `sim/bulk/enchants_test.go` | Loader and fit rules. |
| `sim/bulk/expand.go` | `Expand`/`ExpandWith`, placement, eligibility, shapes, validity, cap. |
| `sim/bulk/expand_test.go` | One table test per rule the contract lists. |
| `sim/bulk/plan.go` | `Plan`/`PlanWith`, `StageRequests`, `Combination`. |
| `sim/bulk/plan_test.go` | Stage 1 shape, seeds, equipped first. |
| `sim/bulk/rank.go` | `Rank`: cuts, next stage, final result, deltas, groups. |
| `sim/bulk/rank_test.go` | Staging arithmetic on fabricated results. |
| `sim/bulk/testdata/enchants.excerpt.json` | Ten-row excerpt of a build's enchant table. |

**Modified files**

| File | Change |
| --- | --- |
| `sim/api/envelope.go` | `Kind()`, `Bulk`, `Weights`, `TargetError`, encounter additions, `CooldownSpec`, result additions, `Progress` additions, `ValidateLane`, `NeedsMoreIterations`. |
| `sim/request/request.go` | `encounter()` gains movement, targets-over-time, dummy, level, armor, type; `Player.Cooldowns`. |
| `sim/request/buffs.go` | `<id>:improved` graded ids. |
| `sim/request/vocabulary.go` | World-buff and stat sections, graded ids, `IDsMarkdown`. |
| `sim/request/IDS.md` | Regenerated. |
| `sim/request/consumables.go` | `ItemID(key)` reverse lookup for cooldown ids. |
| `sim/adapter/adapter.go` | `Weights()` and `Sample()`. |
| `sim/runner/runner.go`, `sim/runner/native.go`, `sim/runner/fixture.go` | `Progress` widened to `api.Progress`. |
| `sim/cmd/wasm/main.go` | `simPlan`, `simRank`, `simWeights`. |
| `sim/cmd/forever-sim/main.go` | Kind dispatch: bulk loop, weights, Smart Sim loop; stage progress lines. |
| `Makefile` | `sim-styles`, `styles-check`; `simdb-check` and `apl-check` unchanged but newly depended on. |
| `.github/workflows/sim.yml` | Style check, enchant-table path, four-candidate Top Gear smoke. |
| `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` | Task 1's amendment only. |

---

### Task 1: Amend the contract with the names this lane needs and the contract does not carry

The contract's preamble says a lane that needs a name not written there adds it there first, in its own commit. Three such names exist. This task writes them and nothing else; no Go changes.

**Files:**
- Modify: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` (section 5, and a new 1.8)

**Interfaces:**
- Consumes: nothing.
- Produces: the engine-lane and data-lane obligations Task 2 verifies — `SimItem.unique`, `SimItem.required_level`, `SimItem.faction_restriction`, `SimItem.random_suffix_options`; `RaidSimResult.sample_iteration` as a `SampleIteration` message; and the rule that a bulk request's `Iterations` equals its precision's final stage count.

- [ ] **Step 1: Add the SimItem fields to section 5**

Append to section 5 ("Engine fork (behind the pin)"), after the `Encounter` block:

````markdown
`proto/common.proto` `SimItem` gains the four fields `sim/bulk`'s
expansion rules need and the sim database does not carry today. Every
one of them already exists on `UIItem` (or on the client's `ItemSparse`
row the data lane reads), so this is a widening of the reduced message,
not new data:

```
bool unique = 20;                           // UIItem.unique
int32 required_level = 21;                  // ItemSparse.RequiredLevel
UIItem.FactionRestriction faction_restriction = 22;
repeated int32 random_suffix_options = 23;  // UIItem.random_suffix_options
```

The data lane fills them in `pipeline/simdb/items.py` from the same
`ItemSparse`/`Item` pair it already reads, and `python -m pipeline
genproto` re-vendors `proto/common.proto` into `data/proto/` first.
Without them `Expand` cannot enforce unique-equipped, required level or
faction, which section 3 requires of it.
````

- [ ] **Step 2: Add the sample-iteration message to section 5**

Replace the last sentence of section 5 ("`RaidSimResult` gains `sample_iteration` …") with:

````markdown
`RaidSimResult` gains `sample_iteration`, the cast log and resource
readings of the median-DPS iteration:

```
message SampleCast {
  double at_seconds = 1;            // negative during the pre-pull
  ActionID action_id = 2;
  string target = 3;                // the target's name; empty for a self-cast
  map<string, int32> resources = 4; // rage, energy, mana, combo_points after the cast
}
message SampleIteration { repeated SampleCast casts = 1; }
```

`SampleIteration sample_iteration = 8;` on `RaidSimResult`. `sim/adapter`
maps it onto `[]api.SampleCast`, taking the row id and display name from
`adapter.ActionName`, so a sample row keys against the cast table.
````

- [ ] **Step 3: Add section 1.8, the two iteration rules the envelope needs**

Append after section 1.7:

````markdown
### 1.8 Iteration counts outside the closed set

`ValidIterations` (500, 3000, 10000) is the settings bar's closed set and
stays the rule for a plain fixed-count run. Two requests are outside it
and are validated by their own rule instead:

- **A bulk request.** `Iterations` must equal the final stage count of
  its `Precision` — 3,000 for `fast` and `normal`, 10,000 for `high` —
  so a stored bulk request says what it was run to without a second
  field. Individual stage requests are parts and are built with
  `ValidatePart`.
- **A target-error run.** When `TargetError > 0`, `Iterations` is a
  ceiling: it must be a positive multiple of `StepIterations` and at most
  the lane's `LaneIterationCeiling`. `MaxIterations` does not bound it.

`SimRequest.Validate()` applies the largest lane's numbers, because the
envelope does not carry a lane. `SimRequest.ValidateLane(lane)` applies
one lane's `Caps` and `LaneIterationCeiling` and is what the api handler
and the page call.

`WeightsSpec.Reference` is required; the envelope does not default it.
Its per-spec default is `data/curated/specs.json`'s `reference_stat`,
which the api lane serves on `GET /v1/specs` and the page sends.

Stat ids are the engine's `Stat` enum names in lower snake case with the
`Stat` prefix stripped, published in `sim/request/IDS.md` the way buff
ids are. The engine has no plain `haste`: it is `spell_haste` and
`melee_haste`, and section 1.4's illustrative list is read that way.
````

- [ ] **Step 4: Commit**

```bash
git add docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md
git commit -m "docs(sim): the SimItem fields, sample-iteration message and iteration rules the parity contract was missing

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 2: Bump the engine pin and prove the new engine surface is there

Every later task in this plan compiles against fields the fork and the data lane add. This task is the gate: it moves the pin, refreshes the embedded database, and adds tests that fail loudly when the surface is absent, so nothing downstream is written against a field that silently does not exist.

**Do not start this task until the engine lane has landed `Encounter.movement`, `Encounter.targets_over_time`, `Encounter.target_dummy`, `RaidSimResult.sample_iteration` and the four `SimItem` fields, and the data lane has re-run `python -m pipeline genproto simdb` for the active build.** Step 1 tells you whether they have.

**Files:**
- Modify: `sim/enginever/version.go` (generated — by `make engine-pin` only)
- Create: `sim/internal/simdb/surface_test.go`
- Modify: `sim/request/mapping_test.go`

**Interfaces:**
- Consumes: Task 1's contract amendment.
- Produces: a pinned engine whose `proto.Encounter` has `Movement`, `TargetsOverTime`, `TargetDummy`; whose `proto.RaidSimResult` has `SampleIteration`; whose `proto.SimItem` has `Unique`, `RequiredLevel`, `FactionRestriction`, `RandomSuffixOptions`; and an embedded `simdb.bin` whose rows populate them.

- [ ] **Step 1: Check the fork and the build carry the new surface**

Run:

```bash
grep -n 'movement\|targets_over_time\|target_dummy\|sample_iteration' \
  /Users/jh/code/wowsims-forever/proto/common.proto /Users/jh/code/wowsims-forever/proto/api.proto
grep -n 'unique\|required_level\|faction_restriction\|random_suffix_options' \
  /Users/jh/code/wowsims-forever/proto/common.proto
```

Expected: every name appears. If any is missing, **stop and report that the engine lane has not landed**; do not proceed and do not add the field yourself.

- [ ] **Step 2: Move the pin and refresh the embedded database**

```bash
make engine-pin
make simdb
```

Expected: `pinned engine version <new sha>` and `embedded data/builds/<build>/simdb.bin (… bytes)`. `make engine-pin` refuses a dirty fork checkout; commit in the fork first if it complains.

- [ ] **Step 3: Write the failing surface test**

Create `sim/internal/simdb/surface_test.go`:

```go
package simdb

import "testing"

// The parity contract's section 5 widens SimItem with the four fields
// sim/bulk's expansion rules read. A build whose rows leave them all at
// zero is a stale simdb.bin, and Expand would then quietly stop
// enforcing unique-equipped and required level - the two rules whose
// absence looks like a correct answer. So the embedded table is asked
// whether anything at all sets them.
func TestTheEmbeddedTableCarriesTheFieldsExpansionNeeds(t *testing.T) {
	db, err := load()
	if err != nil {
		t.Fatal(err)
	}
	var unique, level, faction, suffixes int
	for _, it := range db.Items {
		if it.Unique {
			unique++
		}
		if it.RequiredLevel > 0 {
			level++
		}
		if it.FactionRestriction != 0 {
			faction++
		}
		if len(it.RandomSuffixOptions) > 0 {
			suffixes++
		}
	}
	t.Logf("items=%d unique=%d required_level=%d faction=%d suffixes=%d",
		len(db.Items), unique, level, faction, suffixes)
	for _, c := range []struct {
		name  string
		count int
	}{
		{"unique", unique},
		{"required_level", level},
		{"faction_restriction", faction},
		{"random_suffix_options", suffixes},
	} {
		if c.count == 0 {
			t.Errorf("no item sets %s; re-run the data lane's `python -m pipeline genproto simdb` for the active build", c.name)
		}
	}
}
```

- [ ] **Step 4: Run it to verify it passes against the refreshed build**

Run: `go test ./sim/internal/simdb/... -run TestTheEmbeddedTableCarriesTheFieldsExpansionNeeds -v`
Expected: PASS, with a log line naming four non-zero counts. A zero count means Step 2's `make simdb` copied a build the data lane has not regenerated — go back, do not weaken the test.

- [ ] **Step 5: Add the engine-side guard to the request package**

Append to `sim/request/mapping_test.go`:

```go
// The parity contract's section 5 adds three encounter fields and one
// result field to the fork. sim/request and sim/adapter are written
// against them, so a pin moved backwards is a compile error here rather
// than a field silently left at its zero value in a shipped artifact.
func TestThePinnedEngineCarriesTheParityFields(t *testing.T) {
	e := &proto.Encounter{
		Movement:        &proto.MovementPattern{IntervalSeconds: 45, DurationSeconds: 5},
		TargetsOverTime: []*proto.TargetCountAt{{AtSeconds: 0, Count: 1}},
		TargetDummy:     true,
	}
	if e.GetMovement().GetIntervalSeconds() != 45 || len(e.GetTargetsOverTime()) != 1 || !e.GetTargetDummy() {
		t.Fatal("the pinned engine's Encounter does not carry the parity fields")
	}
	r := &proto.RaidSimResult{SampleIteration: &proto.SampleIteration{}}
	if r.GetSampleIteration() == nil {
		t.Fatal("the pinned engine's RaidSimResult has no sample_iteration")
	}
}
```

- [ ] **Step 6: Run the request tests**

Run: `go test ./sim/request/... -run TestThePinnedEngineCarriesTheParityFields -v`
Expected: PASS.

- [ ] **Step 7: Format and commit**

```bash
gofmt -w sim/enginever/version.go sim/internal/simdb/surface_test.go sim/request/mapping_test.go
git add sim/enginever/version.go sim/internal/simdb/surface_test.go sim/request/mapping_test.go
git commit -m "feat(sim): pin the engine that carries movement, target counts, the dummy and the sample iteration

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 3: The envelope's kinds, bulk block, precision ladders and caps

**Files:**
- Create: `sim/api/bulk.go`
- Create: `sim/api/bulk_test.go`
- Modify: `sim/api/envelope.go` (add the `Bulk` field and call `validateBulk` from `validate`)

**Interfaces:**
- Consumes: `api.LaneBrowser`, `api.LaneServer`, `api.SimRequest` (existing).
- Produces:
  - `api.KindRun/KindGear/KindTalents/KindDrops/KindWeights` string constants
  - `func (r SimRequest) Kind() string`
  - `api.BulkSpec{Mode, Candidates, Talents, Sets, Locked, Precision, Cap}`
  - `api.Candidate{Slot, ItemID, Enchant, Suffix, Origin}`
  - `api.TalentLoadout{Name, Talents}`, `api.GearSet{Name, Gear}`
  - `api.PrecisionFast/PrecisionNormal/PrecisionHigh`, `api.Precisions`
  - `api.Ladder{Iterations []int, Cuts []Cut}`, `api.Cut{Fraction, Top, SlackSE}`, `api.Ladders map[string]Ladder`
  - `func FinalIterations(precision string) (int, bool)`
  - `api.Caps map[string]int`, `api.ErrCapExceeded{Cap, Combinations}`
  - `api.OriginEquipped/OriginBag/OriginBank/OriginSearch/OriginDropPrefix/OriginSetPrefix`

- [ ] **Step 1: Write the failing test**

Create `sim/api/bulk_test.go`:

```go
package api

import (
	"errors"
	"strings"
	"testing"
)

// gear is a minimal valid Top Gear request. Every case below starts
// from it and breaks one thing, so a failure names the rule it broke.
func gear() SimRequest {
	req := runReq()
	req.Iterations = 3000
	req.Bulk = &BulkSpec{
		Mode:       KindGear,
		Precision:  PrecisionNormal,
		Cap:        Caps[LaneBrowser],
		Candidates: []Candidate{{Slot: "head", ItemID: 16963, Origin: OriginBag}},
	}
	return req
}

func TestKindIsDerivedFromTheRequest(t *testing.T) {
	cases := []struct {
		name string
		req  SimRequest
		want string
	}{
		{"a plain run", runReq(), KindRun},
		{"top gear", gear(), KindGear},
		{"droptimizer", func() SimRequest {
			r := gear()
			r.Bulk.Mode = KindDrops
			return r
		}(), KindDrops},
		{"talent compare", func() SimRequest {
			r := gear()
			r.Bulk.Mode = KindTalents
			return r
		}(), KindTalents},
		// The weights kind is Task 5's; its case lives in
		// TestWeightsKindAndShape, beside the type that makes it
		// possible.
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.req.Kind(); got != c.want {
				t.Errorf("Kind() = %q, want %q", got, c.want)
			}
		})
	}
}

func TestBulkValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string // a substring of the error; "" means the request is legal
	}{
		{"the happy path", func(*SimRequest) {}, ""},
		{"an unknown mode", func(r *SimRequest) { r.Bulk.Mode = "everything" }, "bulk.mode"},
		{"an unknown precision", func(r *SimRequest) { r.Bulk.Precision = "exact" }, "bulk.precision"},
		{"a cap over the largest lane's", func(r *SimRequest) { r.Bulk.Cap = Caps[LaneServer] + 1 }, "bulk.cap"},
		{"a cap of zero", func(r *SimRequest) { r.Bulk.Cap = 0 }, "bulk.cap"},
		{"a candidate on a locked slot", func(r *SimRequest) { r.Bulk.Locked = []string{"head"} }, "locked"},
		{"an unknown locked slot", func(r *SimRequest) { r.Bulk.Locked = []string{"tabard"} }, "bulk.locked"},
		{"an origin outside the vocabulary", func(r *SimRequest) { r.Bulk.Candidates[0].Origin = "wishful" }, "origin"},
		{"a drop origin with no source", func(r *SimRequest) { r.Bulk.Candidates[0].Origin = "drop:" }, "origin"},
		{"a set origin with a name", func(r *SimRequest) { r.Bulk.Candidates[0].Origin = "set:my AQ set" }, ""},
		{"a candidate with no item", func(r *SimRequest) { r.Bulk.Candidates[0].ItemID = 0 }, "item_id"},
		{"an unknown candidate slot", func(r *SimRequest) { r.Bulk.Candidates[0].Slot = "tabard" }, "bulk.candidates"},
		{"gear with nothing to try", func(r *SimRequest) { r.Bulk.Candidates = nil }, "at least one"},
		{"talents mode with a loadout", func(r *SimRequest) {
			r.Bulk.Mode = KindTalents
			r.Bulk.Candidates = nil
			r.Bulk.Talents = []TalentLoadout{{Name: "Deep Fury", Talents: "30305001302-05050005525010051"}}
		}, ""},
		{"talents mode with no loadout", func(r *SimRequest) {
			r.Bulk.Mode = KindTalents
			r.Bulk.Candidates = nil
		}, "at least one talent loadout"},
		{"talents mode carrying candidates", func(r *SimRequest) {
			r.Bulk.Mode = KindTalents
			r.Bulk.Talents = []TalentLoadout{{Name: "Deep Fury", Talents: "30305001302-05050005525010051"}}
		}, "no candidates"},
		{"a loadout with no name", func(r *SimRequest) {
			r.Bulk.Talents = []TalentLoadout{{Talents: "30305001302-05050005525010051"}}
		}, "talent loadout"},
		{"drops mode from a boss", func(r *SimRequest) {
			r.Bulk.Mode = KindDrops
			r.Bulk.Candidates[0].Origin = "drop:raid:mc:lucifron"
		}, ""},
		{"drops mode from a bag", func(r *SimRequest) { r.Bulk.Mode = KindDrops }, "every candidate"},
		{"a set with no gear", func(r *SimRequest) { r.Bulk.Sets = []GearSet{{Name: "PvP"}} }, "gear set"},
		{"iterations that are not the precision's final count", func(r *SimRequest) { r.Iterations = 500 }, "iterations must be 3000"},
		{"high precision at ten thousand", func(r *SimRequest) {
			r.Bulk.Precision = PrecisionHigh
			r.Iterations = 10000
		}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := gear()
			c.edit(&req)
			err := req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// The ladders are the fork's fast_mode made explicit, and both lanes
// read them, so their shape is pinned rather than trusted: one cut per
// gap between stages, and a final count that is a number the settings
// bar offers.
func TestTheLaddersAreWellFormed(t *testing.T) {
	want := map[string][]int{
		PrecisionFast:   {100, 1000, 3000},
		PrecisionNormal: {1000, 3000},
		PrecisionHigh:   {1000, 10000},
	}
	for _, p := range Precisions {
		l, ok := Ladders[p]
		if !ok {
			t.Fatalf("no ladder for precision %q", p)
		}
		if len(l.Cuts) != len(l.Iterations)-1 {
			t.Errorf("%s: %d stages and %d cuts; there is one cut per gap", p, len(l.Iterations), len(l.Cuts))
		}
		if got := l.Iterations; !equalInts(got, want[p]) {
			t.Errorf("%s ladder is %v, the contract says %v", p, got, want[p])
		}
		final, ok := FinalIterations(p)
		if !ok || final != l.Iterations[len(l.Iterations)-1] {
			t.Errorf("FinalIterations(%q) = %d, %v", p, final, ok)
		}
	}
	if _, ok := FinalIterations("exact"); ok {
		t.Error("FinalIterations accepted a precision that is not in the vocabulary")
	}
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCapExceededCarriesBothNumbers(t *testing.T) {
	err := error(ErrCapExceeded{Cap: 400, Combinations: 1280})
	var capped ErrCapExceeded
	if !errors.As(err, &capped) {
		t.Fatal("ErrCapExceeded does not match errors.As")
	}
	if capped.Cap != 400 || capped.Combinations != 1280 {
		t.Errorf("got %+v", capped)
	}
	for _, want := range []string{"400", "1280"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("%q does not name %s", err, want)
		}
	}
}
```

Add the shared `runReq()` helper to `sim/api/envelope_test.go` if one does not already exist under another name; if `envelope_test.go` already builds a valid request in a helper, use that helper's name here instead and delete this note. The helper is:

```go
// runReq is a valid plain run: the smallest request Validate accepts.
func runReq() SimRequest {
	return SimRequest{
		EngineVersion: enginever.Version,
		Spec:          "warrior-fury",
		Source:        CharacterSource{Kind: SourceManual},
		Character: CharacterSpec{
			Name:    "Thrall",
			Race:    "orc",
			Class:   "warrior",
			Level:   SimLevel,
			Talents: "30305001302-05050005525010051",
		},
		Encounter:  DefaultEncounter(),
		Iterations: 3000,
		RandomSeed: 7,
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/api/... -run 'TestKindIsDerived|TestBulkValidation|TestTheLadders|TestCapExceeded' -v`
Expected: FAIL to compile — `undefined: KindGear`, `undefined: BulkSpec`, and so on.

- [ ] **Step 3: Write `sim/api/bulk.go`**

```go
package api

// The bulk block: one request shape for every combination tool.
//
// Top Gear, Droptimizer and talent compare ask the same question - a
// base character, a set of substitutions, a ranked answer - so they are
// one kind of request with one planner behind it (sim/bulk) rather than
// three. Mode says which of the three asked; it is also the request's
// Kind, which is why the three mode strings ARE the three kind strings
// and not a parallel vocabulary that could drift from them.

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// The five kinds of request. A kind is derived from the request, never
// sent: a client that could name its own kind could name one the
// request's own shape contradicts, and every row keyed on it would be
// wrong in a way nothing could detect afterwards.
const (
	KindRun     = "run"
	KindGear    = "gear"
	KindTalents = "talents"
	KindDrops   = "drops"
	KindWeights = "weights"
)

// BulkModes is the closed set BulkSpec.Mode is checked against. They are
// the three bulk kinds, in the order the tools were built.
var BulkModes = []string{KindGear, KindTalents, KindDrops}

// Kind reports which tool produced this request.
//
// Task 5 adds the weights arm; until then a request with no bulk
// block is a run.
func (r SimRequest) Kind() string {
	if r.Bulk != nil {
		return r.Bulk.Mode
	}
	return KindRun
}

// BulkSpec is everything that turns one character into many sims.
type BulkSpec struct {
	Mode       string          `json:"mode"`
	Candidates []Candidate     `json:"candidates"`
	Talents    []TalentLoadout `json:"talents,omitempty"`
	Sets       []GearSet       `json:"sets,omitempty"`
	// Locked names slots that are never substituted, whatever the
	// candidate list says. It is the page's "keep what I have here".
	Locked    []string `json:"locked,omitempty"`
	Precision string   `json:"precision"`
	// Cap is the lane's combination cap, echoed into the request so a
	// saved row says what bounded it. The planner refuses an expansion
	// above it; it never silently trims.
	Cap int `json:"cap"`
}

// Candidate is one item the planner may substitute in.
type Candidate struct {
	// Slot is the planner slot name, or "" for "wherever it fits",
	// which is how a ring, a trinket or a one-hander is offered.
	Slot   string `json:"slot"`
	ItemID int    `json:"item_id"`
	// Enchant of 0 inherits the equipped item's enchant for the slot it
	// lands in, where that enchant fits; it does not mean "none".
	Enchant int    `json:"enchant,omitempty"`
	Suffix  int    `json:"suffix,omitempty"`
	Origin  string `json:"origin"`
}

// TalentLoadout is one named build to try.
type TalentLoadout struct {
	Name    string `json:"name"`
	Talents string `json:"talents"`
}

// GearSet is a whole-gear alternative: one candidate that replaces every
// slot at once. It is what Raidbots' legacy Gear Compare was for.
type GearSet struct {
	Name string     `json:"name"`
	Gear []GearSlot `json:"gear"`
}

// Where a candidate came from. The page shows it, history reads it, and
// Droptimizer groups its results by the source id a drop origin names.
const (
	OriginEquipped   = "equipped"
	OriginBag        = "bag"
	OriginBank       = "bank"
	OriginSearch     = "search"
	OriginDropPrefix = "drop:"
	OriginSetPrefix  = "set:"
)

// plainOrigins are the origins that are a whole word.
var plainOrigins = []string{OriginEquipped, OriginBag, OriginBank, OriginSearch}

// The three precisions. They differ in how many stages the planner runs
// and how many iterations the last one gets.
const (
	PrecisionFast   = "fast"
	PrecisionNormal = "normal"
	PrecisionHigh   = "high"
)

// Precisions is the closed set, weakest first.
var Precisions = []string{PrecisionFast, PrecisionNormal, PrecisionHigh}

// Cut is what survives one stage.
//
// Fraction and Top are alternatives: exactly one is non-zero. SlackSE
// widens either by keeping anything whose interval still overlaps the
// last survivor's, which is what stops a stage of 100 iterations from
// throwing away the eventual winner over noise.
type Cut struct {
	Fraction float64 `json:"fraction,omitempty"`
	Top      int     `json:"top,omitempty"`
	SlackSE  float64 `json:"slack_se"`
}

// Ladder is a precision's staging plan: the iteration count of each
// stage and the cut between each pair of them.
type Ladder struct {
	Iterations []int `json:"iterations"`
	Cuts       []Cut `json:"cuts"`
}

// Ladders is the fork's fast_mode ladder made explicit, so the browser
// and the server agree on it rather than each guessing. The equipped set
// runs in every stage, so every delta is paired.
var Ladders = map[string]Ladder{
	PrecisionFast: {
		Iterations: []int{100, 1000, 3000},
		Cuts:       []Cut{{Fraction: 0.25, SlackSE: 2}, {Top: 10, SlackSE: 2}},
	},
	PrecisionNormal: {
		Iterations: []int{1000, 3000},
		Cuts:       []Cut{{Top: 10, SlackSE: 2}},
	},
	PrecisionHigh: {
		Iterations: []int{1000, 10000},
		Cuts:       []Cut{{Top: 20, SlackSE: 2}},
	},
}

// FinalIterations is the iteration count a precision's last stage runs
// at, which is what a bulk request's Iterations field must say.
func FinalIterations(precision string) (int, bool) {
	l, ok := Ladders[precision]
	if !ok || len(l.Iterations) == 0 {
		return 0, false
	}
	return l.Iterations[len(l.Iterations)-1], true
}

// Caps is how many combinations each lane will plan. The browser's
// number is 400 because 400 combinations at 100 iterations is about a
// minute on a mid laptop with eight workers; the server's is the fork's
// own guard against an expansion nothing could finish.
var Caps = map[string]int{LaneBrowser: 400, LaneServer: 20000}

// ErrCapExceeded is returned when an expansion is larger than the lane
// allows. It carries both numbers because the page and the API both say
// them out loud - "31,200 combinations, and this lane plans 20,000" -
// rather than trimming the list and reporting a ranking of a subset
// nobody chose.
type ErrCapExceeded struct {
	Cap          int `json:"cap"`
	Combinations int `json:"combinations"`
}

func (e ErrCapExceeded) Error() string {
	return fmt.Sprintf("the plan is %d combinations, and this lane plans at most %d", e.Combinations, e.Cap)
}

// validateBulk checks everything about the bulk block that a malformed
// client could get wrong, against the largest lane's cap; ValidateLane
// narrows it to one lane.
func (b *BulkSpec) validate(iterations int) []error {
	var errs []error
	if !slices.Contains(BulkModes, b.Mode) {
		errs = append(errs, fmt.Errorf("bulk.mode must be one of %v, got %q", BulkModes, b.Mode))
	}
	if !slices.Contains(Precisions, b.Precision) {
		errs = append(errs, fmt.Errorf("bulk.precision must be one of %v, got %q", Precisions, b.Precision))
	}
	if final, ok := FinalIterations(b.Precision); ok && iterations != final {
		errs = append(errs, fmt.Errorf("a %s bulk request runs its last stage at %d, so iterations must be %d, got %d", b.Precision, final, final, iterations))
	}
	largest := 0
	for _, c := range Caps {
		largest = max(largest, c)
	}
	if b.Cap <= 0 || b.Cap > largest {
		errs = append(errs, fmt.Errorf("bulk.cap must be between 1 and %d, got %d", largest, b.Cap))
	}

	locked := map[string]bool{}
	for _, slot := range b.Locked {
		if !slices.Contains(GearSlots, slot) {
			errs = append(errs, fmt.Errorf("bulk.locked names %q, which is not a gear slot", slot))
			continue
		}
		locked[slot] = true
	}

	for i, c := range b.Candidates {
		if c.ItemID <= 0 {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d] has no item_id", i))
		}
		if c.Slot != "" && !slices.Contains(GearSlots, c.Slot) {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d] names slot %q, which is not a gear slot", i, c.Slot))
		}
		if c.Slot != "" && locked[c.Slot] {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d] is on %q, which is locked", i, c.Slot))
		}
		if err := validateOrigin(c.Origin); err != nil {
			errs = append(errs, fmt.Errorf("bulk.candidates[%d]: %w", i, err))
		}
	}

	for i, l := range b.Talents {
		if l.Name == "" || l.Talents == "" {
			errs = append(errs, fmt.Errorf("bulk.talents[%d] is a talent loadout with no name or no talents", i))
		}
	}
	for i, s := range b.Sets {
		if s.Name == "" || len(s.Gear) == 0 {
			errs = append(errs, fmt.Errorf("bulk.sets[%d] is a gear set with no name or no gear", i))
		}
	}

	switch b.Mode {
	case KindGear:
		if len(b.Candidates) == 0 && len(b.Talents) == 0 && len(b.Sets) == 0 {
			errs = append(errs, errors.New("a gear request needs at least one candidate, talent loadout or set"))
		}
	case KindTalents:
		if len(b.Talents) == 0 {
			errs = append(errs, errors.New("a talents request needs at least one talent loadout"))
		}
		if len(b.Candidates) != 0 || len(b.Sets) != 0 {
			errs = append(errs, errors.New("a talents request carries no candidates and no sets; gear is locked"))
		}
	case KindDrops:
		for i, c := range b.Candidates {
			if !strings.HasPrefix(c.Origin, OriginDropPrefix) {
				errs = append(errs, fmt.Errorf("a drops request needs every candidate to come from a source; bulk.candidates[%d] is %q", i, c.Origin))
			}
		}
		if len(b.Candidates) == 0 {
			errs = append(errs, errors.New("a drops request needs at least one candidate"))
		}
	}
	return errs
}

// validateOrigin holds an origin to the vocabulary. A prefixed origin
// must name something after the colon: "drop:" alone would group a
// Droptimizer result under a source that does not exist.
func validateOrigin(origin string) error {
	if slices.Contains(plainOrigins, origin) {
		return nil
	}
	for _, prefix := range []string{OriginDropPrefix, OriginSetPrefix} {
		if rest, ok := strings.CutPrefix(origin, prefix); ok {
			if rest == "" {
				return fmt.Errorf("origin %q names nothing after the prefix", origin)
			}
			return nil
		}
	}
	return fmt.Errorf("origin must be one of %v, %q<id> or %q<name>, got %q",
		plainOrigins, OriginDropPrefix, OriginSetPrefix, origin)
}
```

- [ ] **Step 4: Add the field and the slot vocabulary to the envelope**

In `sim/api/envelope.go`, add to `SimRequest` after `RandomSeed`:

```go
	// Bulk turns one character into many sims. A request without it is
	// exactly today's single run; see sim/api/bulk.go.
	Bulk *BulkSpec `json:"bulk,omitempty"`
```

and add, beside `MaxTargets`:

```go
// GearSlots is the slot vocabulary the envelope validates against. It is
// the planner's own list and the order sim/request builds the engine's
// positional equipment array in, which is why that package holds the
// table and this one names it; request.SlotNames() is its source and
// TestGearSlotsMatchTheRequestTable proves the two agree.
var GearSlots = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
	"legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
	"main_hand", "off_hand", "ranged",
}
```

and call the bulk validation from `validate`, immediately before the final `return`:

```go
	if r.Bulk != nil {
		errs = append(errs, r.Bulk.validate(r.Iterations)...)
	}
```

Finally, in `validate`, the closed-set iteration check must not also fire for a bulk request, which runs at its precision's final count. Replace the `switch` on iterations with:

```go
	switch {
	case r.Bulk != nil:
		// A bulk request's count is the precision's, checked by
		// BulkSpec.validate against the ladder rather than against the
		// settings bar's closed set.
	case closedSet && !slices.Contains(ValidIterations, r.Iterations):
		errs = append(errs, fmt.Errorf("iterations must be one of %v, got %d", ValidIterations, r.Iterations))
	case !closedSet && (r.Iterations <= 0 || r.Iterations > MaxIterations):
		errs = append(errs, fmt.Errorf("a split part's iterations must be between 1 and %d, got %d", MaxIterations, r.Iterations))
	}
```

- [ ] **Step 5: Prove the slot vocabulary against the request table**

Append to `sim/request/mapping_test.go`:

```go
// api.GearSlots and request.slotOrder are the same list in two packages
// - the envelope cannot import request, because request imports the
// engine and the api module must not. So the two are proved equal here
// rather than kept equal by hand.
func TestGearSlotsMatchTheRequestTable(t *testing.T) {
	if len(api.GearSlots) != len(slotOrder) {
		t.Fatalf("api.GearSlots has %d entries, slotOrder has %d", len(api.GearSlots), len(slotOrder))
	}
	for i, want := range slotOrder {
		if api.GearSlots[i] != want {
			t.Errorf("slot %d is %q in the envelope and %q here", i, api.GearSlots[i], want)
		}
	}
}
```

- [ ] **Step 6: Run the tests**

Run: `go test ./sim/api/... ./sim/request/... -run 'Kind|Bulk|Ladder|CapExceeded|GearSlots' -v`
Expected: PASS.

- [ ] **Step 7: Run the whole api and request packages, then vet and format**

Run: `go test ./sim/api/... ./sim/request/...` then `go vet ./sim/api/... ./sim/request/...` and `gofmt -l sim/api sim/request`
Expected: PASS, no vet findings, no unformatted files.

- [ ] **Step 8: Commit**

```bash
git add sim/api/bulk.go sim/api/bulk_test.go sim/api/envelope.go sim/api/envelope_test.go sim/request/mapping_test.go
git commit -m "feat(sim): the envelope's bulk block, kinds, precision ladders and lane caps

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 4: Smart Sim — target error, the lane ceilings, and the one place that decides "another step?"

Both lanes loop: the page runs stages of `StepIterations` through `simSplit`/`simRun`/`simCombine`, the binary loops inside itself. The decision to run another step is Go, in one function, so the two can never disagree about when a run is precise enough.

**Files:**
- Modify: `sim/api/envelope.go`
- Create: `sim/api/precision_test.go`

**Interfaces:**
- Consumes: `api.SimRequest`, `api.SimResult`, `api.Estimate`, `api.Caps` (Task 3).
- Produces:
  - `SimRequest.TargetError float64` (`json:"target_error,omitempty"`)
  - `const StepIterations = 1000`
  - `var LaneIterationCeiling = map[string]int{LaneBrowser: 30000, LaneServer: 100000}`
  - `func (r SimRequest) ValidateLane(lane string) error`
  - `func NeedsMoreIterations(res SimResult, req SimRequest) bool`
  - `func NextStepIterations(res SimResult, req SimRequest) int`

- [ ] **Step 1: Write the failing test**

Create `sim/api/precision_test.go`:

```go
package api

import (
	"strings"
	"testing"
)

func targetErrorReq(target float64, ceiling int) SimRequest {
	req := runReq()
	req.TargetError = target
	req.Iterations = ceiling
	return req
}

func TestTargetErrorValidation(t *testing.T) {
	cases := []struct {
		name string
		req  SimRequest
		want string
	}{
		{"half a percent up to thirty thousand", targetErrorReq(0.005, 30000), ""},
		{"a ceiling that is not a whole number of steps", targetErrorReq(0.005, 30500), "multiple of 1000"},
		{"a ceiling of zero", targetErrorReq(0.005, 0), "multiple of 1000"},
		{"a ceiling past the largest lane's", targetErrorReq(0.005, 200000), "at most 100000"},
		{"a negative target", targetErrorReq(-0.1, 30000), "target_error"},
		{"a target of one whole", targetErrorReq(1.5, 30000), "target_error"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// Validate uses the largest lane's numbers because the envelope carries
// no lane. ValidateLane is what the api handler and the page call, and
// it is where a browser request asking for a server-sized run is
// refused.
func TestValidateLaneNarrowsToOneLane(t *testing.T) {
	req := targetErrorReq(0.005, 100000)
	if err := req.Validate(); err != nil {
		t.Fatalf("the largest lane's ceiling was refused by Validate: %v", err)
	}
	if err := req.ValidateLane(LaneServer); err != nil {
		t.Fatalf("the server lane refused its own ceiling: %v", err)
	}
	err := req.ValidateLane(LaneBrowser)
	if err == nil {
		t.Fatal("the browser lane accepted 100000 iterations")
	}
	if !strings.Contains(err.Error(), "30000") {
		t.Errorf("error %q does not name the browser's ceiling", err)
	}

	big := gear()
	big.Bulk.Cap = Caps[LaneServer]
	if err := big.ValidateLane(LaneServer); err != nil {
		t.Fatalf("the server lane refused its own cap: %v", err)
	}
	if err := big.ValidateLane(LaneBrowser); err == nil || !strings.Contains(err.Error(), "400") {
		t.Errorf("the browser lane did not refuse a server-sized cap: %v", err)
	}
	if err := req.ValidateLane("moon"); err == nil || !strings.Contains(err.Error(), "lane") {
		t.Errorf("an unknown lane was accepted: %v", err)
	}
}

func TestNeedsMoreIterations(t *testing.T) {
	cases := []struct {
		name  string
		req   SimRequest
		res   SimResult
		want  bool
		steps int // NextStepIterations; -1 means do not check
	}{
		{
			name:  "a fixed-count run never steps",
			req:   runReq(),
			res:   SimResult{IterationsRun: 3000, DPS: Estimate{Mean: 1000, Error: 50}},
			want:  false,
			steps: 0,
		},
		{
			name:  "still outside the target",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 2000, DPS: Estimate{Mean: 1000, Error: 20}},
			want:  true,
			steps: StepIterations,
		},
		{
			name:  "inside the target",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 2000, DPS: Estimate{Mean: 1000, Error: 4}},
			want:  false,
			steps: 0,
		},
		{
			name:  "exactly on the target is inside it",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 2000, DPS: Estimate{Mean: 1000, Error: 5}},
			want:  false,
			steps: 0,
		},
		{
			name:  "the ceiling wins",
			req:   targetErrorReq(0.005, 3000),
			res:   SimResult{IterationsRun: 3000, DPS: Estimate{Mean: 1000, Error: 90}},
			want:  false,
			steps: 0,
		},
		{
			name:  "a short last step rather than an overshoot",
			req:   targetErrorReq(0.005, 3000),
			res:   SimResult{IterationsRun: 2500, DPS: Estimate{Mean: 1000, Error: 90}},
			want:  true,
			steps: 500,
		},
		{
			name:  "a failed run does not step",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 1000, Error: "the engine died", DPS: Estimate{Mean: 0}},
			want:  false,
			steps: 0,
		},
		{
			name:  "a stopped run does not step",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 1000, Aborted: true, DPS: Estimate{Mean: 1000, Error: 90}},
			want:  false,
			steps: 0,
		},
		{
			name:  "no dps yet is not precision",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 1000, DPS: Estimate{Mean: 0, Error: 0}},
			want:  true,
			steps: StepIterations,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := NeedsMoreIterations(c.res, c.req); got != c.want {
				t.Errorf("NeedsMoreIterations = %v, want %v", got, c.want)
			}
			if c.steps >= 0 {
				if got := NextStepIterations(c.res, c.req); got != c.steps {
					t.Errorf("NextStepIterations = %d, want %d", got, c.steps)
				}
			}
		})
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/api/... -run 'TargetError|ValidateLane|NeedsMore' -v`
Expected: FAIL to compile — `undefined: StepIterations`, `undefined: NeedsMoreIterations`.

- [ ] **Step 3: Add the field and the constants to `sim/api/envelope.go`**

Add to `SimRequest`, after `Bulk`. (The `Weights` field is Task 5's, added beside the type it needs.)

```go
	// TargetError, when above zero, turns Iterations into a CEILING:
	// the run continues in steps of StepIterations until DPS.Error over
	// DPS.Mean is at or under it, or Iterations is reached. Zero is
	// today's fixed-count run.
	//
	// This is Raidbots' Smart Sim, made visible. The results line says
	// which of the two ended the run, which is why the loop is a loop
	// over whole results rather than a number the engine is handed.
	TargetError float64 `json:"target_error,omitempty"`
```

Add beside `ValidIterations`:

```go
// StepIterations is how many iterations a target-error run adds at a
// time. It is one number for both lanes: the browser's pool runs a step
// as an ordinary sharded run and the native binary runs it in one go,
// and a step of a different size in each would make the two report
// different iteration counts for the same request.
const StepIterations = 1000

// LaneIterationCeiling bounds a target-error run that never reaches its
// target. Without it a flat distribution would run until the tab was
// closed.
var LaneIterationCeiling = map[string]int{LaneBrowser: 30000, LaneServer: 100000}
```

- [ ] **Step 4: Extend `validate` with the target-error rule**

Replace the iteration `switch` written in Task 3 with:

```go
	switch {
	case r.Bulk != nil:
		// A bulk request's count is the precision's, checked by
		// BulkSpec.validate against the ladder.
	case r.TargetError > 0:
		// A target-error run's Iterations is a ceiling, not a choice
		// from the settings bar, so the closed set does not apply and
		// neither does MaxIterations. It has to be a whole number of
		// steps, because a step is what the loop adds.
		errs = append(errs, validateCeiling(r.Iterations, largestCeiling())...)
	case closedSet && !slices.Contains(ValidIterations, r.Iterations):
		errs = append(errs, fmt.Errorf("iterations must be one of %v, got %d", ValidIterations, r.Iterations))
	case !closedSet && (r.Iterations <= 0 || r.Iterations > MaxIterations):
		errs = append(errs, fmt.Errorf("a split part's iterations must be between 1 and %d, got %d", MaxIterations, r.Iterations))
	}
	if r.TargetError < 0 || r.TargetError >= 1 {
		errs = append(errs, fmt.Errorf("target_error is a fraction of the mean, so it must be between 0 and 1, got %v", r.TargetError))
	}
```

and add the two helpers at the end of the file:

```go
// largestCeiling is the biggest target-error ceiling any lane allows.
// Validate uses it because the envelope carries no lane; ValidateLane
// narrows it.
func largestCeiling() int {
	out := 0
	for _, c := range LaneIterationCeiling {
		out = max(out, c)
	}
	return out
}

// validateCeiling holds a target-error run's ceiling to whole steps and
// to a lane's limit.
func validateCeiling(iterations, ceiling int) []error {
	var errs []error
	if iterations <= 0 || iterations%StepIterations != 0 {
		errs = append(errs, fmt.Errorf("a target-error run's iterations is a ceiling and must be a positive multiple of %d, got %d", StepIterations, iterations))
	}
	if iterations > ceiling {
		errs = append(errs, fmt.Errorf("a target-error run's iterations is at most %d on this lane, got %d", ceiling, iterations))
	}
	return errs
}
```

- [ ] **Step 5: Add `ValidateLane`**

Append to `sim/api/envelope.go`:

```go
// ValidateLane is Validate plus the two limits that belong to a lane
// rather than to the request: the combination cap and the target-error
// ceiling. The envelope carries no lane - a request is a question, and
// the same question can be asked of either - so Validate applies the
// largest lane's numbers and this applies one lane's. The api handler
// calls it with LaneServer and the page with LaneBrowser.
func (r SimRequest) ValidateLane(lane string) error {
	var errs []error
	if err := r.Validate(); err != nil {
		errs = append(errs, err)
	}
	capped, ok := Caps[lane]
	ceiling := LaneIterationCeiling[lane]
	if !ok {
		return errors.Join(append(errs, fmt.Errorf("lane must be %q or %q, got %q", LaneBrowser, LaneServer, lane))...)
	}
	if r.Bulk != nil && r.Bulk.Cap > capped {
		errs = append(errs, fmt.Errorf("the %s lane plans at most %d combinations, and bulk.cap is %d", lane, capped, r.Bulk.Cap))
	}
	if r.TargetError > 0 && r.Iterations > ceiling {
		errs = append(errs, fmt.Errorf("the %s lane runs at most %d iterations, and iterations is %d", lane, ceiling, r.Iterations))
	}
	return errors.Join(errs...)
}
```

- [ ] **Step 6: Add the two loop decisions**

Append to `sim/api/envelope.go`:

```go
// NeedsMoreIterations reports whether a target-error run should run
// another step.
//
// It lives here, in Go, because BOTH lanes loop and they must agree.
// The browser runs each step as an ordinary sharded run through
// simSplit/simRun/simCombine and asks this between steps; forever-sim
// runs the step itself and asks the same function. A copy of this
// arithmetic in TypeScript is exactly the drift the module exists to
// prevent - the two would stop at different precisions and the same
// request would report a different error bar depending on where it ran.
//
// A run that failed or was stopped never steps: there is nothing to
// refine, and stepping would turn one bad answer into several.
func NeedsMoreIterations(res SimResult, req SimRequest) bool {
	if req.TargetError <= 0 || res.Error != "" || res.Aborted {
		return false
	}
	if res.IterationsRun >= req.Iterations {
		return false
	}
	if res.DPS.Mean <= 0 {
		// No mean yet, so no relative error either. One more step is
		// the only way to learn anything, and the ceiling above is what
		// stops it being forever.
		return true
	}
	return res.DPS.Error/res.DPS.Mean > req.TargetError
}

// NextStepIterations is how many iterations the next step runs, and 0
// when there is no next step. The last step is shortened rather than
// overshooting, so a ceiling of 3,000 is 3,000 and never 3,500.
func NextStepIterations(res SimResult, req SimRequest) int {
	if !NeedsMoreIterations(res, req) {
		return 0
	}
	return min(StepIterations, req.Iterations-res.IterationsRun)
}
```

- [ ] **Step 7: Run the tests**

Run: `go test ./sim/api/... -run 'TargetError|ValidateLane|NeedsMore' -v`
Expected: PASS.

- [ ] **Step 8: Run the whole package, vet, format, commit**

```bash
go test ./sim/api/...
go vet ./sim/api/...
gofmt -w sim/api/envelope.go sim/api/precision_test.go
git add sim/api/envelope.go sim/api/precision_test.go
git commit -m "feat(sim): Smart Sim - target error, lane ceilings, and one place that decides another step

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 5: The weights block

**Files:**
- Create: `sim/api/weights.go`
- Create: `sim/api/weights_test.go`
- Modify: `sim/api/envelope.go` (the `Weights` field, the weights arm of `Kind()`, and the validation call)

**Interfaces:**
- Consumes: `api.SimRequest`, `api.Kind()` (Task 3).
- Produces: `SimRequest.Weights *WeightsSpec`, `api.WeightsSpec{Stats []string, Reference string}`, `api.StatWeight{Stat string, Weight float64, Error float64}`, and `Kind()` returning `KindWeights`.

- [ ] **Step 1: Write the failing test**

Create `sim/api/weights_test.go`:

```go
package api

import (
	"strings"
	"testing"
)

func weightsReq() SimRequest {
	req := runReq()
	req.Weights = &WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit", "hit"},
		Reference: "attack_power",
	}
	return req
}

func TestWeightsValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string
	}{
		{"the happy path", func(*SimRequest) {}, ""},
		{"no stats", func(r *SimRequest) { r.Weights.Stats = nil }, "at least one stat"},
		{"no reference", func(r *SimRequest) { r.Weights.Reference = "" }, "weights.reference"},
		{"a reference that is not weighed", func(r *SimRequest) { r.Weights.Reference = "strength" }, "one of the stats it weighs"},
		{"a repeated stat", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", "crit"}
			r.Weights.Reference = "crit"
		}, "listed twice"},
		{"an empty stat id", func(r *SimRequest) {
			r.Weights.Stats = []string{"crit", ""}
			r.Weights.Reference = "crit"
		}, "empty stat id"},
		{"weights and bulk together", func(r *SimRequest) {
			r.Bulk = &BulkSpec{Mode: KindGear, Precision: PrecisionNormal, Cap: 10,
				Candidates: []Candidate{{Slot: "head", ItemID: 1, Origin: OriginBag}}}
		}, "a request is one kind"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := weightsReq()
			c.edit(&req)
			err := req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// The stat ids themselves are the engine's, so the envelope does not
// carry a copy of the list; sim/request resolves them and IDS.md
// publishes them. What the envelope owns is the shape.
func TestWeightsKindAndShape(t *testing.T) {
	req := weightsReq()
	if req.Kind() != KindWeights {
		t.Errorf("Kind() = %q, want %q", req.Kind(), KindWeights)
	}
	w := StatWeight{Stat: "crit", Weight: 1.0, Error: 0.04}
	if w.Stat != "crit" || w.Weight != 1.0 || w.Error != 0.04 {
		t.Errorf("StatWeight round trip: %+v", w)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/api/... -run Weights -v`
Expected: FAIL to compile — `undefined: WeightsSpec`.

- [ ] **Step 3: Write `sim/api/weights.go`**

```go
package api

// Stat weights.
//
// The engine already computes them, so this is a request shape and a
// result shape, not arithmetic. The caution belongs to the page: a
// weight is a linear guess at a non-linear thing, and simming the
// actual items is the better answer. The envelope's job is to make the
// request honest - a reference stat that is not among the stats being
// weighed would normalise against a number nobody computed.

import (
	"errors"
	"fmt"
)

// WeightsSpec asks for stat weights instead of a DPS number.
type WeightsSpec struct {
	// Stats are IDS.md stat ids: the engine's Stat enum names in lower
	// snake case with the prefix stripped, "attack_power", "crit",
	// "spell_haste". sim/request resolves them; an id the engine has no
	// stat for is refused there rather than weighed as nothing.
	Stats []string `json:"stats"`
	// Reference is the stat normalised to exactly 1.0. It is required:
	// its per-spec default lives in data/curated/specs.json and the
	// caller sends it, because the envelope has no spec table and a
	// default invented here would be a second source of truth.
	Reference string `json:"reference"`
}

// StatWeight is one stat's answer. Reference's Weight is exactly 1.
type StatWeight struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
	Error  float64 `json:"error"`
}

func (w *WeightsSpec) validate() []error {
	var errs []error
	if len(w.Stats) == 0 {
		errs = append(errs, errors.New("weights needs at least one stat to weigh"))
	}
	seen := map[string]bool{}
	for _, s := range w.Stats {
		if s == "" {
			errs = append(errs, errors.New("weights.stats carries an empty stat id"))
			continue
		}
		if seen[s] {
			errs = append(errs, fmt.Errorf("weights.stats lists %q twice", s))
		}
		seen[s] = true
	}
	switch {
	case w.Reference == "":
		errs = append(errs, errors.New("weights.reference is required; the spec's default is in data/curated/specs.json"))
	case !seen[w.Reference]:
		errs = append(errs, fmt.Errorf("weights.reference is %q, which is not one of the stats it weighs (%v); the reference is normalised to 1 and there would be nothing to normalise", w.Reference, w.Stats))
	}
	return errs
}
```

- [ ] **Step 4: Add the field and the weights arm of `Kind()`**

Add to `SimRequest` in `sim/api/envelope.go`, after `Bulk`:

```go
	// Weights asks for stat weights instead of a DPS run.
	Weights *WeightsSpec `json:"weights,omitempty"`
```

and replace `Kind()` in `sim/api/bulk.go` with its finished form:

```go
// Kind reports which tool produced this request.
func (r SimRequest) Kind() string {
	switch {
	case r.Bulk != nil:
		return r.Bulk.Mode
	case r.Weights != nil:
		return KindWeights
	default:
		return KindRun
	}
}
```

- [ ] **Step 5: Call the validation from `validate`**

In `sim/api/envelope.go`, beside the bulk call added in Task 3:

```go
	if r.Bulk != nil {
		errs = append(errs, r.Bulk.validate(r.Iterations)...)
	}
	if r.Weights != nil {
		errs = append(errs, r.Weights.validate()...)
	}
	if r.Bulk != nil && r.Weights != nil {
		// Kind() reads Bulk first, so a request carrying both would run
		// as a bulk and silently answer a different question from the
		// one the weights block asked.
		errs = append(errs, errors.New("a request is one kind: it carries bulk or weights, never both"))
	}
```

- [ ] **Step 6: Run the tests**

Run: `go test ./sim/api/... -run Weights -v`
Expected: PASS.

- [ ] **Step 7: Run the package, vet, format, commit**

```bash
go test ./sim/api/...
go vet ./sim/api/...
gofmt -w sim/api/weights.go sim/api/weights_test.go sim/api/envelope.go
git add sim/api/weights.go sim/api/weights_test.go sim/api/envelope.go
git commit -m "feat(sim): the envelope's stat weights block

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 6: The encounter's new fields and the character's cooldowns

**Files:**
- Modify: `sim/api/envelope.go`
- Create: `sim/api/encounter_test.go`

**Interfaces:**
- Consumes: `api.EncounterSpec`, `api.CharacterSpec`, `api.BossLevel`.
- Produces:
  - `EncounterSpec.Style string`, `.Movement *Movement`, `.TargetsOverTime []TargetCount`, `.TargetLevel int`, `.TargetArmor int`, `.TargetType string`, `.Dummy bool`
  - `api.Movement{IntervalSec, DurationSec, Kind}`, `api.MovementAway`, `api.MovementCasting`, `api.MovementKinds`
  - `api.TargetCount{AtSec, Count}`
  - `api.TargetTypes []string`, `api.TargetTypeUnknown`
  - `api.MinTargetLevel = 60`, `api.MaxTargetLevel = 63`, `api.TargetArmorByLevel map[int]int`, `func TargetArmorFor(level, override int) int`
  - `CharacterSpec.Cooldowns []CooldownSpec`, `api.CooldownSpec{ID, AtSec}`

- [ ] **Step 1: Write the failing test**

Create `sim/api/encounter_test.go`:

```go
package api

import (
	"strings"
	"testing"
)

func TestEncounterAdditionsValidation(t *testing.T) {
	cases := []struct {
		name string
		edit func(*SimRequest)
		want string
	}{
		{"the default encounter is unchanged and legal", func(*SimRequest) {}, ""},
		{"a style label rides along", func(r *SimRequest) { r.Encounter.Style = "heavy-movement" }, ""},
		{"movement away", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{IntervalSec: 20, DurationSec: 5, Kind: MovementAway}
		}, ""},
		{"movement with an unknown kind", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{IntervalSec: 20, DurationSec: 5, Kind: "teleport"}
		}, "movement.kind"},
		{"movement longer than its interval", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{IntervalSec: 5, DurationSec: 20, Kind: MovementAway}
		}, "shorter than the interval"},
		{"movement with no interval", func(r *SimRequest) {
			r.Encounter.Movement = &Movement{DurationSec: 5, Kind: MovementAway}
		}, "movement.interval_sec"},
		{"a target-count timeline", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}}
		}, ""},
		{"a timeline that does not start at zero", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 40, Count: 3}}
		}, "starts at 0"},
		{"a timeline out of order", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}, {AtSec: 20, Count: 5}}
		}, "in time order"},
		{"a timeline over the target cap", func(r *SimRequest) {
			r.Encounter.TargetsOverTime = []TargetCount{{AtSec: 0, Count: MaxTargets + 1}}
		}, "targets_over_time"},
		{"target level 60", func(r *SimRequest) { r.Encounter.TargetLevel = 60 }, ""},
		{"target level 64", func(r *SimRequest) { r.Encounter.TargetLevel = 64 }, "target_level"},
		{"an armor override", func(r *SimRequest) { r.Encounter.TargetArmor = 2500 }, ""},
		{"negative armor", func(r *SimRequest) { r.Encounter.TargetArmor = -1 }, "target_armor"},
		{"a target type", func(r *SimRequest) { r.Encounter.TargetType = "undead" }, ""},
		{"an unknown target type", func(r *SimRequest) { r.Encounter.TargetType = "murloc" }, "target_type"},
		{"the dummy", func(r *SimRequest) { r.Encounter.Dummy = true }, ""},
		{"a cooldown on cooldown", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{ID: "spell:1719"}}
		}, ""},
		{"a cooldown at fixed times", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{ID: "spell:1719", AtSec: []float64{0, 90}}}
		}, ""},
		{"a cooldown with no id", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{AtSec: []float64{0}}}
		}, "cooldowns"},
		{"a cooldown used before the pull", func(r *SimRequest) {
			r.Character.Cooldowns = []CooldownSpec{{ID: "spell:1719", AtSec: []float64{-1}}}
		}, "at_sec"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := runReq()
			c.edit(&req)
			err := req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// Zero armor means the level's preset, not a naked target. Getting that
// backwards would report every sim against an unarmoured boss.
func TestTargetArmorFor(t *testing.T) {
	cases := []struct {
		level, override, want int
	}{
		{BossLevel, 0, TargetArmorByLevel[BossLevel]},
		{60, 0, TargetArmorByLevel[60]},
		{0, 0, TargetArmorByLevel[BossLevel]}, // an unset level is the default boss
		{BossLevel, 2500, 2500},
		{99, 0, TargetArmorByLevel[BossLevel]}, // a level with no preset falls back
	}
	for _, c := range cases {
		if got := TargetArmorFor(c.level, c.override); got != c.want {
			t.Errorf("TargetArmorFor(%d, %d) = %d, want %d", c.level, c.override, got, c.want)
		}
	}
	for level := MinTargetLevel; level <= MaxTargetLevel; level++ {
		if TargetArmorByLevel[level] <= 0 {
			t.Errorf("no armor preset for target level %d", level)
		}
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/api/... -run 'EncounterAdditions|TargetArmorFor' -v`
Expected: FAIL to compile — `undefined: Movement`.

- [ ] **Step 3: Extend `EncounterSpec` and `CharacterSpec` in `sim/api/envelope.go`**

Replace `EncounterSpec` with:

```go
type EncounterSpec struct {
	DurationSec  int     `json:"duration_sec"`
	Variation    float64 `json:"variation"`
	Targets      int     `json:"targets"`
	ExecuteRatio float64 `json:"execute_ratio"`
	Profile      string  `json:"profile"`

	// Style is the page preset's LABEL and nothing more: the fields
	// below carry what it expanded to. Storing the label as well as the
	// fields is what lets a saved sim say "Heavy movement" after the
	// preset's numbers have been retuned, and lets an edited request
	// stop claiming a preset it no longer matches.
	Style string `json:"style,omitempty"`

	// Movement schedules time out of melee or out of casting.
	Movement *Movement `json:"movement,omitempty"`

	// TargetsOverTime OVERRIDES Targets when set: a dungeon pull is a
	// boss and then packs, not a fixed count.
	TargetsOverTime []TargetCount `json:"targets_over_time,omitempty"`

	// TargetLevel is 60 to 63; 0 means BossLevel, which is what every
	// sim fought before this field existed.
	TargetLevel int `json:"target_level,omitempty"`

	// TargetArmor overrides the level's preset. 0 means the preset -
	// NOT an unarmoured target.
	TargetArmor int `json:"target_armor,omitempty"`

	// TargetType changes what Hunter and Warlock abilities do. "" is
	// TargetTypeUnknown, which is what a target dummy is.
	TargetType string `json:"target_type,omitempty"`

	// Dummy is the training dummy: no debuffs, no execute window, no
	// armor reduction.
	Dummy bool `json:"dummy,omitempty"`
}

// Movement is a repeating window the player spends away from the target
// or unable to cast.
type Movement struct {
	IntervalSec int    `json:"interval_sec"`
	DurationSec int    `json:"duration_sec"`
	Kind        string `json:"kind"`
}

// The two kinds of movement window. Away is out of melee range with no
// casting; Casting interrupts spells while melee continues.
const (
	MovementAway    = "away"
	MovementCasting = "casting"
)

// MovementKinds is the closed set.
var MovementKinds = []string{MovementAway, MovementCasting}

// TargetCount is one step of a target-count timeline: from AtSec, this
// many targets are alive.
type TargetCount struct {
	AtSec int `json:"at_sec"`
	Count int `json:"count"`
}

// The target levels the settings bar offers.
const (
	MinTargetLevel = SimLevel
	MaxTargetLevel = BossLevel
)

// TargetArmorByLevel is the armor a target of each level carries when
// the request does not override it. The boss row is the engine's own
// preset (sim/encounters/default_presets.go); the three below it are
// the Classic armor-by-level table the same presets were derived from.
// They are here rather than in sim/request because the page shows the
// preset beside the override control, and a second copy there would
// drift from the number the sim actually ran.
var TargetArmorByLevel = map[int]int{60: 3000, 61: 3210, 62: 3421, 63: 3731}

// TargetArmorFor resolves an encounter's armor: the override when it is
// set, otherwise the level's preset, otherwise the boss's.
func TargetArmorFor(level, override int) int {
	if override > 0 {
		return override
	}
	if armor, ok := TargetArmorByLevel[level]; ok {
		return armor
	}
	return TargetArmorByLevel[BossLevel]
}

// TargetTypeUnknown is a target with no creature type, which is what a
// training dummy is and what every sim fought before this field.
const TargetTypeUnknown = "unknown"

// TargetTypes is the closed set, matching the engine's MobType enum.
// sim/request proves the pairing against the enum.
var TargetTypes = []string{
	"beast", "demon", "dragonkin", "elemental", "giant",
	"humanoid", "mechanical", "undead", TargetTypeUnknown,
}
```

Add to `CharacterSpec`, after `Consumes`:

```go
	// Cooldowns is when to use the major cooldowns and potions the
	// rotation would otherwise fire on cooldown.
	Cooldowns []CooldownSpec `json:"cooldowns,omitempty"`
```

and beside it:

```go
// CooldownSpec pins one cooldown's timings.
type CooldownSpec struct {
	// ID is "spell:<id>", "item:<id>", or a consumable id from IDS.md,
	// which the build's consumable table resolves to an item.
	ID string `json:"id"`
	// AtSec are the times to use it, in fight seconds. Each value is
	// one usage; usages past the list happen as soon as the rotation
	// allows. An empty list is "on cooldown", which is the default the
	// engine already has.
	AtSec []float64 `json:"at_sec,omitempty"`
}
```

- [ ] **Step 4: Validate them**

Add to `validate`, after the existing `execute_ratio` check:

```go
	errs = append(errs, validateEncounterAdditions(r.Encounter)...)
	for i, cd := range r.Character.Cooldowns {
		if cd.ID == "" {
			errs = append(errs, fmt.Errorf("character.cooldowns[%d] has no id", i))
		}
		for j, at := range cd.AtSec {
			if at < 0 {
				errs = append(errs, fmt.Errorf("character.cooldowns[%d].at_sec[%d] is %v; a cooldown is used during the fight, and the pre-pull is the rotation's job", i, j, at))
			}
		}
	}
```

and add the helper at the end of the file:

```go
// validateEncounterAdditions checks the fields the parity contract added
// to EncounterSpec. They are all optional, so every check is on a value
// the client actually sent.
func validateEncounterAdditions(e EncounterSpec) []error {
	var errs []error
	if m := e.Movement; m != nil {
		if !slices.Contains(MovementKinds, m.Kind) {
			errs = append(errs, fmt.Errorf("encounter.movement.kind must be one of %v, got %q", MovementKinds, m.Kind))
		}
		if m.IntervalSec <= 0 {
			errs = append(errs, fmt.Errorf("encounter.movement.interval_sec must be positive, got %d", m.IntervalSec))
		}
		if m.DurationSec <= 0 {
			errs = append(errs, fmt.Errorf("encounter.movement.duration_sec must be positive, got %d", m.DurationSec))
		}
		if m.IntervalSec > 0 && m.DurationSec >= m.IntervalSec {
			errs = append(errs, fmt.Errorf("encounter.movement.duration_sec (%d) must be shorter than the interval (%d), or the player never stands still", m.DurationSec, m.IntervalSec))
		}
	}
	if len(e.TargetsOverTime) > 0 {
		if e.TargetsOverTime[0].AtSec != 0 {
			errs = append(errs, fmt.Errorf("encounter.targets_over_time starts at 0, not %d; the fight has targets from the pull", e.TargetsOverTime[0].AtSec))
		}
		prev := -1
		for i, tc := range e.TargetsOverTime {
			if tc.AtSec <= prev {
				errs = append(errs, fmt.Errorf("encounter.targets_over_time must be in time order; entry %d is at %d after %d", i, tc.AtSec, prev))
			}
			prev = tc.AtSec
			if tc.Count < 1 || tc.Count > MaxTargets {
				errs = append(errs, fmt.Errorf("encounter.targets_over_time[%d].count must be between 1 and %d, got %d", i, MaxTargets, tc.Count))
			}
		}
	}
	if e.TargetLevel != 0 && (e.TargetLevel < MinTargetLevel || e.TargetLevel > MaxTargetLevel) {
		errs = append(errs, fmt.Errorf("encounter.target_level must be between %d and %d, got %d", MinTargetLevel, MaxTargetLevel, e.TargetLevel))
	}
	if e.TargetArmor < 0 {
		errs = append(errs, fmt.Errorf("encounter.target_armor must not be negative, got %d; 0 means the level's preset", e.TargetArmor))
	}
	if e.TargetType != "" && !slices.Contains(TargetTypes, e.TargetType) {
		errs = append(errs, fmt.Errorf("encounter.target_type must be one of %v, got %q", TargetTypes, e.TargetType))
	}
	return errs
}
```

- [ ] **Step 5: Run the tests**

Run: `go test ./sim/api/... -run 'EncounterAdditions|TargetArmorFor' -v`
Expected: PASS.

- [ ] **Step 6: Run the package, vet, format, commit**

```bash
go test ./sim/api/...
go vet ./sim/api/...
gofmt -w sim/api/envelope.go sim/api/encounter_test.go
git add sim/api/envelope.go sim/api/encounter_test.go
git commit -m "feat(sim): movement, target-count timelines, target level, armor, type, dummy and cooldown timings

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 7: The result's new fields and the progress line's stage

**Files:**
- Modify: `sim/api/envelope.go`
- Create: `sim/api/result_test.go`
- Modify: `sim/combine/combine.go` (`shape` must ignore the fields a stage varies)
- Modify: `sim/combine/combine_test.go`

**Interfaces:**
- Consumes: `api.Estimate`, `api.GearSlot`, `api.StatWeight` (Task 5).
- Produces:
  - `SimResult.Combos []Combo`, `.Equipped *Estimate`, `.Stages []Stage`, `.Weights []StatWeight`, `.Sample []SampleCast`
  - `api.Combo{Substitutions, DPS, Delta, Group}`
  - `api.Substitution{Kind, Slot, ItemID, Enchant, Suffix, Name, Talents, Origin}` with `SubstitutionItem/Talents/Set`
  - `api.Stage{Iterations, Combos}`
  - `api.SampleCast{AtMS, SpellID, Name, Target, Resources}`
  - `Progress.Stage`, `.CombosDone`, `.CombosTotal`

- [ ] **Step 1: Write the failing test**

Create `sim/api/result_test.go`:

```go
package api

import (
	"encoding/json"
	"testing"
)

// The web mirrors these tags verbatim, so the JSON names are pinned
// here rather than left to whoever edits the struct next. A plain run's
// result must not grow a single key: every new field is omitempty, and
// a run that filled none of them marshals exactly as it did before.
func TestAPlainResultGainsNoKeys(t *testing.T) {
	b, err := json.Marshal(SimResult{})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"combos", "equipped", "stages", "weights", "sample"} {
		if _, ok := got[key]; ok {
			t.Errorf("a plain result carries %q; every bulk field is omitempty", key)
		}
	}
}

func TestBulkResultJSONNames(t *testing.T) {
	res := SimResult{
		Combos: []Combo{{
			Substitutions: []Substitution{
				{Kind: SubstitutionItem, Slot: "main_hand", ItemID: 19352, Enchant: 2568, Origin: OriginBag},
				{Kind: SubstitutionTalents, Name: "Deep Fury", Talents: "30305001302-05050005525010051"},
				{Kind: SubstitutionSet, Name: "my AQ set"},
			},
			DPS:   Estimate{Mean: 1100},
			Delta: Estimate{Mean: 41, Error: 9},
			Group: 0,
		}},
		Equipped: &Estimate{Mean: 1059},
		Stages:   []Stage{{Iterations: 100, Combos: 38}, {Iterations: 1000, Combos: 10}},
		Weights:  []StatWeight{{Stat: "crit", Weight: 1, Error: 0.04}},
		Sample:   []SampleCast{{AtMS: -1500, SpellID: 1719, Name: "spell:1719", Resources: map[string]int{"rage": 0}}},
	}
	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"combos", "equipped", "stages", "weights", "sample"} {
		if _, ok := got[key]; !ok {
			t.Errorf("a bulk result does not carry %q", key)
		}
	}

	var back SimResult
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Combos) != 1 || back.Combos[0].Delta.Mean != 41 || back.Combos[0].Substitutions[0].ItemID != 19352 {
		t.Errorf("round trip lost something: %+v", back.Combos)
	}
	if back.Sample[0].AtMS != -1500 || back.Sample[0].Resources["rage"] != 0 {
		t.Errorf("round trip lost the sample: %+v", back.Sample)
	}
}

// The progress payload is the page's partial result, so the three bulk
// fields are absent for a plain run rather than zero-valued keys the
// page has to ignore.
func TestProgressGainsTheStageFields(t *testing.T) {
	b, err := json.Marshal(Progress{IterationsRun: 100, DPS: Estimate{Mean: 900}})
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"iterations_run":100,"dps":{"mean":900,"stddev":0,"error":0,"min":0,"max":0}}` {
		t.Errorf("a plain run's progress is %s", b)
	}
	b, err = json.Marshal(Progress{IterationsRun: 100, Stage: 2, CombosDone: 31, CombosTotal: 96})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"stage", "combos_done", "combos_total"} {
		if _, ok := got[key]; !ok {
			t.Errorf("a bulk run's progress does not carry %q", key)
		}
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/api/... -run 'PlainResult|BulkResultJSON|ProgressGains' -v`
Expected: FAIL to compile — `undefined: Combo`.

- [ ] **Step 3: Extend `SimResult` and `Progress`**

Add to `SimResult`, after `Aborted`:

```go
	// Combos is the ranked answer to a bulk request, best first.
	Combos []Combo `json:"combos,omitempty"`
	// Equipped is the base character at the FINAL stage, which is what
	// every Delta is measured against. It is a pointer so that "no
	// baseline" and "a baseline of zero" are different things.
	Equipped *Estimate `json:"equipped,omitempty"`
	// Stages is what the planner actually ran, so the page can say
	// "3 stages, 38 combinations" without recomputing the ladder.
	Stages []Stage `json:"stages,omitempty"`
	// Weights is the answer to a weights request.
	Weights []StatWeight `json:"weights,omitempty"`
	// Sample is one iteration's casts in order - the median-DPS one.
	// It is a table view of the cast log, not a guide, and the page
	// says so.
	Sample []SampleCast `json:"sample,omitempty"`
```

Add the types beside `Estimate`:

```go
// Combo is one substitution set's result.
type Combo struct {
	Substitutions []Substitution `json:"substitutions"`
	DPS           Estimate       `json:"dps"`
	// Delta is against Equipped, paired at the same stage, WITH ITS OWN
	// ERROR - which is what lets the page say "within error" honestly
	// instead of ranking noise.
	Delta Estimate `json:"delta"`
	// Group is 0 for the leader's within-error group, then 1, 2, and so
	// on. Rows in one group are the same answer and the page ranks them
	// the same.
	Group int `json:"group"`
}

// The three things a combination can substitute.
const (
	SubstitutionItem    = "item"
	SubstitutionTalents = "talents"
	SubstitutionSet     = "set"
)

// Substitution is one change from the base character.
type Substitution struct {
	Kind string `json:"kind"`
	// Slot is where an item landed. A ring or a trinket says which of
	// the two slots it took, because that is the answer the player
	// acts on.
	Slot    string `json:"slot,omitempty"`
	ItemID  int    `json:"item_id,omitempty"`
	Enchant int    `json:"enchant,omitempty"`
	Suffix  int    `json:"suffix,omitempty"`
	// Name is the loadout's or the set's.
	Name    string `json:"name,omitempty"`
	Talents string `json:"talents,omitempty"`
	Origin  string `json:"origin,omitempty"`
}

// Stage is one rung of the ladder, as run.
type Stage struct {
	Iterations int `json:"iterations"`
	Combos     int `json:"combos"`
}

// SampleCast is one cast of the sample iteration.
type SampleCast struct {
	// AtMS is negative during the pre-pull, which the page renders as
	// its own section.
	AtMS int64 `json:"at_ms"`
	// SpellID is the summary's row id for the action, from
	// adapter.ActionName, so a sample row keys against the cast table.
	SpellID int64  `json:"spell_id"`
	Name    string `json:"name"`
	Target  string `json:"target,omitempty"`
	// Resources are what the player held AFTER the cast: rage, energy,
	// mana, combo_points.
	Resources map[string]int `json:"resources,omitempty"`
}
```

Replace `Progress` with:

```go
type Progress struct {
	IterationsRun int      `json:"iterations_run"`
	DPS           Estimate `json:"dps"`
	// The three below are a bulk run's, and absent for a plain one:
	// stage 1 of 3, 31 of 96 combinations. Stage numbers start at 1.
	Stage       int `json:"stage,omitempty"`
	CombosDone  int `json:"combos_done,omitempty"`
	CombosTotal int `json:"combos_total,omitempty"`
}
```

- [ ] **Step 4: Teach `combine.shape` about the fields a stage varies**

`combine.Results` refuses parts that "ask a different question". A stage's parts share everything, so nothing changes there — but a whole-run result now carries `Bulk`, and the shape comparison must still work when the page splits a *stage* request (whose `Bulk` is nil). Add the two fields a step or a stage legitimately varies to `sim/combine/combine.go`'s `shape`:

```go
// shape is a request with the fields a split or a Smart Sim step is
// allowed to vary cleared, so two parts of one run compare equal.
//
// TargetError joins the seed and the count because a step of a
// target-error run is a fixed-count part of it: the loop owns the
// target, the part does not.
func shape(r api.SimRequest) api.SimRequest {
	r.RandomSeed = 0
	r.Iterations = 0
	r.TargetError = 0
	return r
}
```

Add to `sim/combine/combine_test.go`:

```go
// A Smart Sim step is a part of a target-error run. The parts carry no
// target of their own, and the whole run's result does, so the shape
// comparison must not call them different questions.
func TestPartsOfATargetErrorRunCombine(t *testing.T) {
	base := part()
	base.Request.TargetError = 0.005
	base.Request.Iterations = 30000
	other := part()
	other.Request.TargetError = 0
	out, err := Results([]api.SimResult{base, other})
	if err != nil {
		t.Fatalf("two steps of one run were called different questions: %v", err)
	}
	if out.IterationsRun != base.IterationsRun+other.IterationsRun {
		t.Errorf("iterations_run = %d", out.IterationsRun)
	}
}
```

Use whatever helper `combine_test.go` already has for a valid partial result in place of `part()`; if it has none, add one that returns a `SimResult` with `EngineVersion: enginever.Version`, a valid `Request`, `IterationsRun: 750` and a positive `DPS.Mean`.

- [ ] **Step 5: Run the tests**

Run: `go test ./sim/api/... ./sim/combine/... -v`
Expected: PASS.

- [ ] **Step 6: Vet, format, commit**

```bash
go vet ./sim/api/... ./sim/combine/...
gofmt -w sim/api/envelope.go sim/api/result_test.go sim/combine/combine.go sim/combine/combine_test.go
git add sim/api/envelope.go sim/api/result_test.go sim/combine/combine.go sim/combine/combine_test.go
git commit -m "feat(sim): combos, equipped, stages, weights and the sample log on the result

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 8: The fight-style table, as a Go map and the JSON the page is tested against

A style is a page preset, but the expansion is Go: the page applies it in TypeScript and a second, hand-kept copy of the table would drift. So the Go map is the source, it is generated to JSON, the JSON is committed, and the web lane's test compares its own table against that file.

**Files:**
- Create: `sim/request/styles.go`
- Create: `sim/request/styles_test.go`
- Create: `sim/request/styles.json` (generated)
- Create: `sim/internal/genstyles/main.go`

**Interfaces:**
- Consumes: `api.EncounterSpec`, `api.Movement`, `api.TargetCount`, `api.DefaultEncounter()` (Task 6).
- Produces:
  - `request.StyleIDs []string` — the vocabulary, in the page's display order
  - `func request.ExpandStyle(id string, base api.EncounterSpec) (api.EncounterSpec, bool)`
  - `func request.StylesJSON() ([]byte, error)`
  - `sim/request/styles.json`

- [ ] **Step 1: Write the failing test**

Create `sim/request/styles_test.go`:

```go
package request

import (
	"bytes"
	"os"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// The contract's section 1.6 table, transcribed. It is written out
// here rather than derived from the map it checks, because a table
// that checked itself would agree with any typo.
func TestEveryStyleExpandsToTheContractsFields(t *testing.T) {
	cases := []struct {
		id      string
		targets int
		execute float64
		move    *api.Movement
		overT   []api.TargetCount
		dummy   bool
	}{
		{id: "patchwerk", targets: 1, execute: 0.25},
		{id: "execute", targets: 1, execute: 0.35},
		{id: "light-movement", targets: 1, execute: 0.25,
			move: &api.Movement{IntervalSec: 45, DurationSec: 5, Kind: api.MovementAway}},
		{id: "heavy-movement", targets: 1, execute: 0.25,
			move: &api.Movement{IntervalSec: 20, DurationSec: 5, Kind: api.MovementAway}},
		{id: "cleave-2", targets: 2, execute: 0.25},
		{id: "cleave-3", targets: 3, execute: 0.25},
		{id: "cleave-5", targets: 5, execute: 0.25},
		{id: "dungeon", targets: 1, execute: 0, overT: []api.TargetCount{
			{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}, {AtSec: 80, Count: 5},
			{AtSec: 130, Count: 3}, {AtSec: 160, Count: 1}}},
		{id: "dummy", targets: 1, execute: 0, dummy: true},
	}
	if len(cases) != len(StyleIDs) {
		t.Fatalf("the contract lists %d styles and StyleIDs has %d: %v", len(cases), len(StyleIDs), StyleIDs)
	}
	for _, c := range cases {
		t.Run(c.id, func(t *testing.T) {
			if !slices.Contains(StyleIDs, c.id) {
				t.Fatalf("StyleIDs does not carry %q", c.id)
			}
			got, ok := ExpandStyle(c.id, api.DefaultEncounter())
			if !ok {
				t.Fatalf("ExpandStyle(%q) is not a style", c.id)
			}
			if got.Style != c.id {
				t.Errorf("style label is %q", got.Style)
			}
			if got.Targets != c.targets {
				t.Errorf("targets = %d, want %d", got.Targets, c.targets)
			}
			if got.ExecuteRatio != c.execute {
				t.Errorf("execute_ratio = %v, want %v", got.ExecuteRatio, c.execute)
			}
			if got.Dummy != c.dummy {
				t.Errorf("dummy = %v, want %v", got.Dummy, c.dummy)
			}
			switch {
			case c.move == nil && got.Movement != nil:
				t.Errorf("movement = %+v, want none", got.Movement)
			case c.move != nil && got.Movement == nil:
				t.Errorf("movement is absent, want %+v", c.move)
			case c.move != nil && *got.Movement != *c.move:
				t.Errorf("movement = %+v, want %+v", *got.Movement, *c.move)
			}
			if len(got.TargetsOverTime) != len(c.overT) {
				t.Fatalf("targets_over_time = %v, want %v", got.TargetsOverTime, c.overT)
			}
			for i := range c.overT {
				if got.TargetsOverTime[i] != c.overT[i] {
					t.Errorf("targets_over_time[%d] = %+v, want %+v", i, got.TargetsOverTime[i], c.overT[i])
				}
			}
		})
	}
	if _, ok := ExpandStyle("naxx", api.DefaultEncounter()); ok {
		t.Error("ExpandStyle accepted a style that is not in the vocabulary")
	}
}

// Every expansion has to survive the envelope's own validation, or the
// page would offer a preset the run then refuses.
func TestEveryStyleExpandsToALegalEncounter(t *testing.T) {
	for _, id := range StyleIDs {
		t.Run(id, func(t *testing.T) {
			req := fury()
			got, ok := ExpandStyle(id, req.Encounter)
			if !ok {
				t.Fatalf("no such style")
			}
			req.Encounter = got
			if err := req.Validate(); err != nil {
				t.Errorf("the %q preset does not validate: %v", id, err)
			}
		})
	}
}

// The fields a style does NOT set are the player's, and an expansion
// that reset them would silently throw away a chosen fight length.
func TestExpandStyleKeepsWhatTheStyleDoesNotSet(t *testing.T) {
	base := api.DefaultEncounter()
	base.DurationSec = 300
	base.Variation = 0.1
	base.TargetLevel = 61
	base.TargetArmor = 2500
	base.TargetType = "undead"
	got, _ := ExpandStyle("cleave-3", base)
	if got.DurationSec != 300 || got.Variation != 0.1 || got.TargetLevel != 61 || got.TargetArmor != 2500 || got.TargetType != "undead" {
		t.Errorf("the style overwrote what it does not own: %+v", got)
	}
}

// styles.json is the web lane's fixture: its own style table is tested
// against this file, so the two cannot drift. A stale copy is a failing
// test here rather than a disagreement nobody notices.
func TestStylesJSONIsCommitted(t *testing.T) {
	want, err := StylesJSON()
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile("styles.json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.TrimSpace(got), bytes.TrimSpace(want)) {
		t.Error("sim/request/styles.json is stale; run `go run ./internal/genstyles` from sim/")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/request/... -run Style -v`
Expected: FAIL to compile — `undefined: StyleIDs`.

- [ ] **Step 3: Write `sim/request/styles.go`**

```go
package request

// The fight styles.
//
// A style is a page preset: a name a player picks, and the encounter
// fields it stands for. The expansion lives here, in Go, because the
// page applies it too and a second copy of the table in TypeScript
// would drift the first time a preset was retuned - the page would
// offer "Heavy movement" and the sim would run last month's numbers.
// So this map is the source, `go run ./internal/genstyles` renders it
// to styles.json, and the web lane's test compares its own table
// against that file.
//
// A style sets only the fields it owns. Fight length, variation, target
// level, armor and type are the player's and survive a style change,
// which is what makes the style control a preset rather than a reset.

import (
	"encoding/json"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// style is one preset's expansion, as a delta over the base encounter.
type style struct {
	Targets         int                `json:"targets"`
	ExecuteRatio    float64            `json:"execute_ratio"`
	Movement        *api.Movement      `json:"movement,omitempty"`
	TargetsOverTime []api.TargetCount  `json:"targets_over_time,omitempty"`
	Dummy           bool               `json:"dummy,omitempty"`
}

// StyleIDs is the vocabulary, in the order the page lists it.
var StyleIDs = []string{
	"patchwerk",
	"execute",
	"light-movement",
	"heavy-movement",
	"cleave-2",
	"cleave-3",
	"cleave-5",
	"dungeon",
	"dummy",
}

// The movement windows the two movement styles use. Five seconds is
// one global cooldown plus a step; the interval is what makes light
// and heavy different.
const (
	movementWindowSec   = 5
	lightMovementEvery  = 45
	heavyMovementEvery  = 20
	patchwerkExecute    = 0.25
	executeHeavyExecute = 0.35
)

// styles is the contract's section 1.6 table.
var styles = map[string]style{
	"patchwerk": {Targets: 1, ExecuteRatio: patchwerkExecute},
	"execute":   {Targets: 1, ExecuteRatio: executeHeavyExecute},
	"light-movement": {Targets: 1, ExecuteRatio: patchwerkExecute,
		Movement: &api.Movement{IntervalSec: lightMovementEvery, DurationSec: movementWindowSec, Kind: api.MovementAway}},
	"heavy-movement": {Targets: 1, ExecuteRatio: patchwerkExecute,
		Movement: &api.Movement{IntervalSec: heavyMovementEvery, DurationSec: movementWindowSec, Kind: api.MovementAway}},
	"cleave-2": {Targets: 2, ExecuteRatio: patchwerkExecute},
	"cleave-3": {Targets: 3, ExecuteRatio: patchwerkExecute},
	"cleave-5": {Targets: 5, ExecuteRatio: patchwerkExecute},
	// A boss, then packs of three and five, then back down: the shape
	// of a dungeon pull. No execute window, because a pack dies from
	// full health to nothing rather than sliding down a health bar.
	"dungeon": {Targets: 1, ExecuteRatio: 0, TargetsOverTime: []api.TargetCount{
		{AtSec: 0, Count: 1},
		{AtSec: 40, Count: 3},
		{AtSec: 80, Count: 5},
		{AtSec: 130, Count: 3},
		{AtSec: 160, Count: 1},
	}},
	// A training dummy has no debuffs, no execute and no armor
	// reduction; the engine's target_dummy flag is what turns all
	// three off at once.
	"dummy": {Targets: 1, ExecuteRatio: 0, Dummy: true},
}

// ExpandStyle applies a style to an encounter, keeping every field the
// style does not own. The label rides along in Style so a saved sim can
// say which preset produced it.
func ExpandStyle(id string, base api.EncounterSpec) (api.EncounterSpec, bool) {
	s, ok := styles[id]
	if !ok {
		return base, false
	}
	out := base
	out.Style = id
	out.Targets = s.Targets
	out.ExecuteRatio = s.ExecuteRatio
	out.Dummy = s.Dummy
	out.Movement = nil
	if s.Movement != nil {
		// A copy: the map's pointer is shared by every caller, and a
		// caller that edited the encounter it got back would retune the
		// preset for the whole process.
		m := *s.Movement
		out.Movement = &m
	}
	out.TargetsOverTime = slices.Clone(s.TargetsOverTime)
	return out, true
}

// StylesJSON renders the table as the web lane's fixture: an array in
// StyleIDs order, each entry the style id and its expansion.
func StylesJSON() ([]byte, error) {
	type row struct {
		ID string `json:"id"`
		style
	}
	out := make([]row, 0, len(StyleIDs))
	for _, id := range StyleIDs {
		out = append(out, row{ID: id, style: styles[id]})
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(b, '\n'), nil
}
```

- [ ] **Step 4: Write the generator**

Create `sim/internal/genstyles/main.go`:

```go
// Command genstyles writes sim/request/styles.json, the fight-style
// table the web lane's own style table is tested against.
//
//	go run ./internal/genstyles
//
// TestStylesJSONIsCommitted fails when the committed file and this
// output disagree, so retuning a preset in Go and forgetting the page
// is a failing test rather than two products disagreeing quietly.
package main

import (
	"log"
	"os"

	"github.com/jhunthrop/foreversixty/sim/request"
)

func main() {
	out := "request/styles.json"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	b, err := request.StylesJSON()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(out, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s", out)
}
```

- [ ] **Step 5: Generate the file and run the tests**

```bash
cd sim && go run ./internal/genstyles && cd ..
go test ./sim/request/... -run Style -v
```

Expected: `wrote request/styles.json`, then PASS on all four style tests.

- [ ] **Step 6: Vet, format, commit**

```bash
go vet ./sim/request/... ./sim/internal/genstyles/...
gofmt -w sim/request/styles.go sim/request/styles_test.go sim/internal/genstyles/main.go
git add sim/request/styles.go sim/request/styles_test.go sim/request/styles.json sim/internal/genstyles/main.go
git commit -m "feat(sim): the fight-style table, and the JSON the page is tested against

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 9: Map the new encounter fields onto the engine

**Files:**
- Modify: `sim/request/request.go` (`encounter`, `biomeFor` untouched)
- Modify: `sim/request/request_test.go`

**Interfaces:**
- Consumes: `api.EncounterSpec` additions (Task 6), the pinned engine's `proto.MovementPattern`, `proto.TargetCountAt`, `Encounter.TargetDummy` (Task 2).
- Produces: `request.Build`/`BuildWith` filling `Encounter.Movement`, `Encounter.TargetsOverTime`, `Encounter.TargetDummy`, and every `Target`'s `Level`, `MobType` and `Stats[StatArmor]`.

- [ ] **Step 1: Write the failing test**

Append to `sim/request/request_test.go`:

```go
// The five encounter fields the parity contract added all reach the
// engine, and the target the sim has always built is unchanged when
// none of them is set.
func TestEncounterCarriesTheParityFields(t *testing.T) {
	req := fury()
	req.Encounter.Movement = &api.Movement{IntervalSec: 20, DurationSec: 5, Kind: api.MovementCasting}
	req.Encounter.TargetsOverTime = []api.TargetCount{{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}}
	req.Encounter.TargetLevel = 61
	req.Encounter.TargetArmor = 2500
	req.Encounter.TargetType = "undead"
	req.Encounter.Dummy = true

	got, err := Build(req)
	if err != nil {
		t.Fatal(err)
	}
	e := got.Encounter
	if m := e.GetMovement(); m == nil || m.IntervalSeconds != 20 || m.DurationSeconds != 5 || !m.CastingOnly {
		t.Errorf("movement = %+v", e.GetMovement())
	}
	if len(e.GetTargetsOverTime()) != 2 || e.GetTargetsOverTime()[1].AtSeconds != 40 || e.GetTargetsOverTime()[1].Count != 3 {
		t.Errorf("targets_over_time = %+v", e.GetTargetsOverTime())
	}
	if !e.GetTargetDummy() {
		t.Error("target_dummy is not set")
	}
	// A timeline overrides the fixed count, and the pool is sized to
	// the largest count the timeline ever reaches: the engine activates
	// and deactivates targets from it rather than creating them.
	if len(e.Targets) != 3 {
		t.Fatalf("the target pool is %d, want the timeline's maximum of 3", len(e.Targets))
	}
	for i, target := range e.Targets {
		if target.Level != 61 {
			t.Errorf("target %d is level %d", i, target.Level)
		}
		if target.MobType != proto.MobType_MobTypeUndead {
			t.Errorf("target %d is %v", i, target.MobType)
		}
		if got := target.Stats[proto.Stat_StatArmor]; got != 2500 {
			t.Errorf("target %d armor = %v, want the override 2500", i, got)
		}
	}
}

// Nothing set is the fight the sim has always built, plus the armor
// preset the contract now says every target carries.
func TestEncounterDefaultsAreUnchanged(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	e := got.Encounter
	if e.GetMovement() != nil || len(e.GetTargetsOverTime()) != 0 || e.GetTargetDummy() {
		t.Errorf("a plain request set a parity field: %+v", e)
	}
	if len(e.Targets) != 1 {
		t.Fatalf("targets = %d", len(e.Targets))
	}
	target := e.Targets[0]
	if target.Level != api.BossLevel {
		t.Errorf("level = %d, want %d", target.Level, api.BossLevel)
	}
	if target.MobType != proto.MobType_MobTypeHumanoid {
		t.Errorf("mob_type = %v, want humanoid, which is what the sim has always fought", target.MobType)
	}
	if got := target.Stats[proto.Stat_StatArmor]; got != float64(api.TargetArmorByLevel[api.BossLevel]) {
		t.Errorf("armor = %v, want the boss preset %d", got, api.TargetArmorByLevel[api.BossLevel])
	}
}

// The target-type vocabulary is the engine's enum, and a name that
// resolved to nothing would fight a creature with no type and change
// what Hunter and Warlock abilities do without a word.
func TestTargetTypesMatchTheEngineEnum(t *testing.T) {
	for _, id := range api.TargetTypes {
		if _, ok := mobTypes[id]; !ok {
			t.Errorf("api.TargetTypes lists %q, which this package cannot map", id)
		}
	}
	for value, name := range proto.MobType_name {
		id := strcase.Snake(strings.TrimPrefix(name, "MobType"))
		if _, ok := mobTypes[id]; !ok {
			t.Errorf("the engine has MobType %s (%d) and no id maps to it", name, value)
		}
	}
	if len(mobTypes) != len(proto.MobType_name) {
		t.Errorf("mobTypes has %d entries, the enum has %d", len(mobTypes), len(proto.MobType_name))
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/request/... -run 'EncounterCarries|EncounterDefaults|TargetTypes' -v`
Expected: FAIL to compile — `undefined: mobTypes`; and the default test fails on armor.

- [ ] **Step 3: Replace `encounter` in `sim/request/request.go`**

```go
// mobTypes maps our target-type ids onto the engine's enum. The ids are
// the enum names in lower snake case with the prefix stripped, and
// TestTargetTypesMatchTheEngineEnum holds the map to the enum in both
// directions: a creature type the engine models and we cannot name is
// a Hunter's Slaying bonus nobody can ask for.
var mobTypes = map[string]proto.MobType{
	"beast":                   proto.MobType_MobTypeBeast,
	"demon":                   proto.MobType_MobTypeDemon,
	"dragonkin":               proto.MobType_MobTypeDragonkin,
	"elemental":               proto.MobType_MobTypeElemental,
	"giant":                   proto.MobType_MobTypeGiant,
	"humanoid":                proto.MobType_MobTypeHumanoid,
	"mechanical":              proto.MobType_MobTypeMechanical,
	"undead":                  proto.MobType_MobTypeUndead,
	api.TargetTypeUnknown:     proto.MobType_MobTypeUnknown,
}

// targetCount is how many targets the encounter needs BUILT. A
// target-count timeline overrides the fixed count, and the engine
// activates and deactivates targets from a pool, so the pool is the
// largest count the timeline ever reaches.
func targetCount(e api.EncounterSpec) int {
	if len(e.TargetsOverTime) == 0 {
		return e.Targets
	}
	most := 1
	for _, tc := range e.TargetsOverTime {
		most = max(most, tc.Count)
	}
	return most
}

// targetStats is one target's stat array, with the encounter's armor in
// it. The engine indexes the array by proto.Stat, so it is built to the
// enum's length rather than to the highest index we happen to set.
func targetStats(e api.EncounterSpec) []float64 {
	out := make([]float64, len(proto.Stat_name))
	out[proto.Stat_StatArmor] = float64(api.TargetArmorFor(e.TargetLevel, e.TargetArmor))
	return out
}

// movement turns our window into the engine's pattern. Our two kinds
// are the engine's one boolean: "casting" interrupts spells without
// moving, "away" leaves melee range too.
func movement(m *api.Movement) *proto.MovementPattern {
	if m == nil {
		return nil
	}
	return &proto.MovementPattern{
		IntervalSeconds: float64(m.IntervalSec),
		DurationSeconds: float64(m.DurationSec),
		CastingOnly:     m.Kind == api.MovementCasting,
	}
}

// targetsOverTime turns our timeline into the engine's.
func targetsOverTime(steps []api.TargetCount) []*proto.TargetCountAt {
	if len(steps) == 0 {
		return nil
	}
	out := make([]*proto.TargetCountAt, len(steps))
	for i, s := range steps {
		out[i] = &proto.TargetCountAt{AtSeconds: float64(s.AtSec), Count: int32(s.Count)}
	}
	return out
}

func encounter(e api.EncounterSpec) *proto.Encounter {
	below20, below25, below35 := executeProportions(e.ExecuteRatio)
	// A target dummy has no execute window: nothing kills it, so its
	// health never falls. The envelope's Dummy flag is the one place
	// that says so, rather than the page being asked to zero the ratio
	// as well as tick the box.
	if e.Dummy {
		below20, below25, below35 = 0, 0, 0
	}
	level := e.TargetLevel
	if level == 0 {
		level = api.BossLevel
	}
	mob, ok := mobTypes[e.TargetType]
	if !ok {
		// "" is not a type the page offers; it is the shape of every
		// request written before the field existed, and those fought a
		// humanoid. Keeping that is what stops the pin bump changing
		// every stored spec's number.
		mob = proto.MobType_MobTypeHumanoid
	}
	targets := make([]*proto.Target, targetCount(e))
	for i := range targets {
		targets[i] = &proto.Target{
			Id:        targetDummyID,
			Name:      targetDummyName,
			Level:     int32(level),
			MobType:   mob,
			Stats:     targetStats(e),
			TankIndex: targetNotTanked,
		}
	}
	return &proto.Encounter{
		Duration:             float64(e.DurationSec),
		Biome:                biomeFor(e.Profile),
		DurationVariation:    float64(e.DurationSec) * e.Variation,
		ExecuteProportion_20: below20,
		ExecuteProportion_25: below25,
		ExecuteProportion_35: below35,
		Targets:              targets,
		Movement:             movement(e.Movement),
		TargetsOverTime:      targetsOverTime(e.TargetsOverTime),
		TargetDummy:          e.Dummy,
	}
}
```

Add `"strings"` and the `strcase` import to `request_test.go` if they are not already there.

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/request/... -run 'Encounter|TargetTypes' -v`
Expected: PASS.

- [ ] **Step 5: Run the whole package**

Run: `go test ./sim/request/...`
Expected: PASS. If `rotations_smoke_test.go` now produces different DPS because every target carries armor where none did before, that is the intended change — the smoke test asserts a positive DPS, not a number, so it should still pass. If it asserts a number, update the number in the same commit and say so in the message.

- [ ] **Step 6: Vet, format, commit**

```bash
go vet ./sim/request/...
gofmt -w sim/request/request.go sim/request/request_test.go
git add sim/request/request.go sim/request/request_test.go
git commit -m "feat(sim): movement, target-count timelines, the dummy, target level, armor and type reach the engine

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 10: Cooldown timings onto `proto.Cooldowns`

**Files:**
- Create: `sim/request/cooldowns.go`
- Create: `sim/request/cooldowns_test.go`
- Modify: `sim/request/consumables.go` (the reverse lookup)
- Modify: `sim/request/request.go` (`BuildWith` sets `player.Cooldowns`)

**Interfaces:**
- Consumes: `api.CooldownSpec` (Task 6), `request.Consumables`.
- Produces:
  - `func (c *Consumables) ItemID(key string) (int64, bool)`
  - `request.ErrUnknownCooldown`
  - `player.Cooldowns` filled from `CharacterSpec.Cooldowns`

- [ ] **Step 1: Write the failing test**

Create `sim/request/cooldowns_test.go`:

```go
package request

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func testConsumables(t *testing.T) *Consumables {
	t.Helper()
	f, err := os.Open("testdata/simconsumes.excerpt.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	table, err := LoadConsumables(f)
	if err != nil {
		t.Fatal(err)
	}
	return table
}

func TestCooldownIDGrammar(t *testing.T) {
	table := testConsumables(t)
	cases := []struct {
		name     string
		id       string
		wantKind string // "spell" or "item"
		wantRaw  int32
		wantErr  bool
	}{
		{name: "a spell", id: "spell:1719", wantKind: "spell", wantRaw: 1719},
		{name: "an item", id: "item:13452", wantKind: "item", wantRaw: 13452},
		{name: "a consumable by name", id: "elixir_of_the_mongoose", wantKind: "item", wantRaw: 13452},
		{name: "a spell with no number", id: "spell:", wantErr: true},
		{name: "a spell that is not a number", id: "spell:mongoose", wantErr: true},
		{name: "a bare word the table does not know", id: "elixir_of_nothing", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := cooldownAction(c.id, table)
			if c.wantErr {
				if err == nil {
					t.Fatalf("%q was accepted as %+v", c.id, got)
				}
				if !errors.Is(err, ErrUnknownCooldown) {
					t.Errorf("error %v is not ErrUnknownCooldown", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			switch c.wantKind {
			case "spell":
				if got.GetSpellId() != c.wantRaw {
					t.Errorf("spell_id = %d, want %d", got.GetSpellId(), c.wantRaw)
				}
			case "item":
				if got.GetItemId() != c.wantRaw {
					t.Errorf("item_id = %d, want %d", got.GetItemId(), c.wantRaw)
				}
			}
		})
	}
}

// The excerpt's ids are whatever the ten rows carry; if 13452 is not
// "Elixir of the Mongoose" in testdata/simconsumes.excerpt.json, read
// the file and use a row that is there. The grammar is what is under
// test, not the particular item.

func TestBuildCarriesCooldownTimings(t *testing.T) {
	req := fury()
	req.Character.Cooldowns = []api.CooldownSpec{
		{ID: "spell:1719", AtSec: []float64{0, 90}},
		{ID: "spell:20572"},
	}
	got, err := BuildWith(req, Options{Consumables: testConsumables(t)})
	if err != nil {
		t.Fatal(err)
	}
	cds := got.Raid.Parties[0].Players[0].GetCooldowns().GetCooldowns()
	if len(cds) != 2 {
		t.Fatalf("cooldowns = %d, want 2", len(cds))
	}
	if cds[0].Id.GetSpellId() != 1719 || len(cds[0].Timings) != 2 || cds[0].Timings[1] != 90 {
		t.Errorf("first cooldown = %+v", cds[0])
	}
	// No timings is "on cooldown", which the engine spells as an empty
	// list rather than as an absent entry: the entry is what makes the
	// cooldown eligible at all.
	if cds[1].Id.GetSpellId() != 20572 || len(cds[1].Timings) != 0 {
		t.Errorf("second cooldown = %+v", cds[1])
	}
}

func TestBuildWithNoCooldownsLeavesTheMessageAbsent(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if cd := got.Raid.Parties[0].Players[0].GetCooldowns(); cd != nil && len(cd.Cooldowns) != 0 {
		t.Errorf("a request with no cooldowns produced %+v", cd)
	}
}

func TestACooldownNamedByItemNeedsTheTable(t *testing.T) {
	req := fury()
	req.Character.Cooldowns = []api.CooldownSpec{{ID: "elixir_of_the_mongoose"}}
	_, err := Build(req) // Build carries no consumable table
	if err == nil {
		t.Fatal("a consumable cooldown resolved without a table")
	}
	if !strings.Contains(err.Error(), "no consumable table") {
		t.Errorf("error %q does not say the table is missing", err)
	}
}

// Two items whose names normalise to one key cannot be told apart, so
// the reverse lookup refuses rather than picking one.
func TestConsumablesItemIDRefusesAnAmbiguousKey(t *testing.T) {
	table := &Consumables{names: map[int64]string{1: "twin", 2: "twin", 3: "alone"}}
	if _, ok := table.ItemID("twin"); ok {
		t.Error("an ambiguous key resolved")
	}
	if id, ok := table.ItemID("alone"); !ok || id != 3 {
		t.Errorf("ItemID(alone) = %d, %v", id, ok)
	}
	if _, ok := table.ItemID("absent"); ok {
		t.Error("an absent key resolved")
	}
	if _, ok := (*Consumables)(nil).ItemID("alone"); ok {
		t.Error("a nil table resolved a key")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/request/... -run 'Cooldown|ConsumablesItemID' -v`
Expected: FAIL to compile — `undefined: cooldownAction`.

- [ ] **Step 3: Add the reverse lookup to `sim/request/consumables.go`**

```go
// ItemID is key's item, and the reverse of key(): the cooldown panel
// names a potion by the same id the consumable panel does, and the
// engine wants an item id for its ActionID.
//
// It refuses an ambiguous key rather than picking one. Two client items
// whose names normalise to the same key are two different potions, and
// pinning the wrong one's timing would be a silently different run.
func (c *Consumables) ItemID(key string) (int64, bool) {
	if c == nil {
		return 0, false
	}
	var found int64
	var seen int
	for id, k := range c.names {
		if k == key {
			found = id
			seen++
		}
	}
	if seen != 1 {
		return 0, false
	}
	return found, true
}
```

- [ ] **Step 4: Write `sim/request/cooldowns.go`**

```go
package request

// Cooldown timings.
//
// The engine already carries them: proto.Cooldowns is a list of
// (ActionID, timings), where an empty timing list means "on cooldown,
// whenever the rotation allows" and a value is one usage at a fixed
// second. All this file does is name an action the way the settings bar
// names things - "spell:1719", "item:13452", or a consumable id from
// IDS.md - and refuse one it cannot resolve, because a major cooldown
// quietly dropped is a DPS number that is wrong and says nothing.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// ErrUnknownCooldown is returned for a cooldown id this module cannot
// turn into an engine action.
var ErrUnknownCooldown = errors.New("request: unknown cooldown")

const (
	spellPrefix = "spell"
	// itemPrefix is declared in consumables.go and is "item".
)

// cooldownAction turns one cooldown id into the engine's ActionID.
func cooldownAction(id string, table *Consumables) (*proto.ActionID, error) {
	head, tail, qualified := strings.Cut(id, ":")
	switch {
	case qualified && head == spellPrefix:
		n, err := strconv.ParseInt(tail, 10, 32)
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("%w: %q names no spell", ErrUnknownCooldown, id)
		}
		return &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: int32(n)}}, nil
	case qualified && head == itemPrefix:
		n, err := strconv.ParseInt(tail, 10, 32)
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("%w: %q names no item", ErrUnknownCooldown, id)
		}
		return &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: int32(n)}}, nil
	}
	// A bare id is a consumable's, which only the build's table can
	// turn into an item. Saying the table is missing beats resolving
	// the id to nothing.
	if table == nil {
		return nil, fmt.Errorf("%w: %q is a consumable and no consumable table is loaded; pass one in Options, from data/builds/<build>/simconsumes.json", ErrUnknownCooldown, id)
	}
	itemID, ok := table.ItemID(id)
	if !ok {
		return nil, fmt.Errorf("%w: %q is not a spell, an item, or a consumable this build's table names exactly once", ErrUnknownCooldown, id)
	}
	return &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: int32(itemID)}}, nil
}

// cooldownsFor builds the engine's Cooldowns message, or nil when the
// character pinned none - which is the engine's own default and must
// stay distinguishable from "a message with no entries".
func cooldownsFor(specs []api.CooldownSpec, table *Consumables) (*proto.Cooldowns, error) {
	if len(specs) == 0 {
		return nil, nil
	}
	out := &proto.Cooldowns{Cooldowns: make([]*proto.Cooldown, 0, len(specs))}
	for _, cd := range specs {
		action, err := cooldownAction(cd.ID, table)
		if err != nil {
			return nil, err
		}
		out.Cooldowns = append(out.Cooldowns, &proto.Cooldown{Id: action, Timings: cd.AtSec})
	}
	return out, nil
}
```

If `spellPrefix` collides with an existing declaration, drop the `const` block here and use the existing one.

- [ ] **Step 5: Call it from `BuildWith`**

In `sim/request/request.go`, after `buffs, err := buffsFor(ch.Buffs)`:

```go
	cooldowns, err := cooldownsFor(ch.Cooldowns, opt.Consumables)
	if err != nil {
		return nil, err
	}
```

and add to the `proto.Player` literal, after `Consumes: cons,`:

```go
		Cooldowns:   cooldowns,
```

- [ ] **Step 6: Run the tests**

Run: `go test ./sim/request/... -run 'Cooldown|ConsumablesItemID' -v`
Expected: PASS. If the excerpt fixture does not carry item 13452, read `sim/request/testdata/simconsumes.excerpt.json`, pick a row it does carry, and use that id and name in the test — do not add a row to the fixture.

- [ ] **Step 7: Run the package, vet, format, commit**

```bash
go test ./sim/request/...
go vet ./sim/request/...
gofmt -w sim/request/cooldowns.go sim/request/cooldowns_test.go sim/request/consumables.go sim/request/request.go
git add sim/request/cooldowns.go sim/request/cooldowns_test.go sim/request/consumables.go sim/request/request.go
git commit -m "feat(sim): cooldown timings reach the engine's Cooldowns message

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 11: Graded buff ids, the world-buff section, and the stat vocabulary

**Files:**
- Modify: `sim/request/buffs.go`
- Modify: `sim/request/vocabulary.go`
- Modify: `sim/request/vocabulary_test.go`
- Modify: `sim/request/IDS.md` (regenerated)

**Interfaces:**
- Consumes: the engine's `TristateEffect`, `IndividualBuffs` and `Stat` descriptors.
- Produces:
  - `<id>:improved` accepted by `buffsFor` for every tristate field
  - `request.KnownBuffs()` listing both forms
  - `request.WorldBuffs() []string`
  - `request.KnownStats() []string`, `func request.ParseStat(id string) (proto.Stat, bool)`
  - `sim/request/IDS.md` with a **World buffs** section and a **Stats** section

- [ ] **Step 1: Write the failing test**

Append to `sim/request/vocabulary_test.go`:

```go
// A graded buff is the engine's TristateEffect: value one is the plain
// version and value two the talented one. The settings bar needs both,
// so every tristate field answers to "<id>" and "<id>:improved".
func TestGradedBuffIds(t *testing.T) {
	got, err := buffsFor([]string{"battle_shout:improved", "power_word_fortitude"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.BattleShout != proto.TristateEffect_TristateEffectImproved {
		t.Errorf("battle_shout:improved set %v, want the improved version", got.Raid.BattleShout)
	}
	if got.Raid.PowerWordFortitude != proto.TristateEffect_TristateEffectRegular {
		t.Errorf("the plain id set %v", got.Raid.PowerWordFortitude)
	}

	// A field that is not graded has no improved form, and accepting
	// one would silently apply the plain buff under a name that
	// promised more.
	if _, err := buffsFor([]string{"blessing_of_kings:improved"}); err == nil {
		t.Error("an improved form of a boolean buff was accepted")
	}
	if _, err := buffsFor([]string{"no_such_buff:improved"}); err == nil {
		t.Error("an improved form of an unknown buff was accepted")
	}
}

func TestKnownBuffsListsBothForms(t *testing.T) {
	ids := KnownBuffs()
	for _, want := range []string{"battle_shout", "battle_shout:improved"} {
		if !slices.Contains(ids, want) {
			t.Errorf("KnownBuffs does not list %q", want)
		}
	}
	if slices.Contains(ids, "blessing_of_kings:improved") {
		t.Error("KnownBuffs lists an improved form for a buff that has none")
	}
	for _, id := range ids {
		if _, err := buffsFor([]string{id}); err != nil {
			t.Errorf("KnownBuffs lists %q, which does not resolve: %v", id, err)
		}
	}
}

// The world buffs are a section of IndividualBuffs, and the engine
// marks them only with a comment. So the list is written down here and
// held to the message: the engine's world-buff fields are numbered from
// worldBuffFirstField up, and a new one that nobody added to the list
// fails rather than going missing from the settings bar.
func TestWorldBuffsCoverTheEnginesWorldBuffSection(t *testing.T) {
	names := map[string]bool{}
	for _, id := range WorldBuffs() {
		names[id] = true
	}
	desc := (&proto.IndividualBuffs{}).ProtoReflect().Descriptor()
	fields := desc.Fields()
	var expected int
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if fd.Number() < worldBuffFirstField {
			continue
		}
		expected++
		if !names[string(fd.Name())] {
			t.Errorf("IndividualBuffs.%s is a world buff and WorldBuffs() does not list it", fd.Name())
		}
	}
	if len(names) != expected {
		t.Errorf("WorldBuffs() has %d entries and IndividualBuffs has %d world-buff fields", len(names), expected)
	}
	if !slices.IsSorted(WorldBuffs()) {
		t.Error("WorldBuffs() is not sorted; the list is an interface and must be stable")
	}
}

// Stat ids are the engine's enum, so the weights panel is built from
// the engine rather than from a table that would drift.
func TestKnownStatsAreTheEnginesEnum(t *testing.T) {
	ids := KnownStats()
	if !slices.IsSorted(ids) {
		t.Error("KnownStats is not sorted")
	}
	if len(ids) != len(proto.Stat_name) {
		t.Errorf("KnownStats has %d entries, the enum has %d", len(ids), len(proto.Stat_name))
	}
	for _, id := range ids {
		if _, ok := ParseStat(id); !ok {
			t.Errorf("KnownStats lists %q, which does not resolve", id)
		}
	}
	for _, want := range []string{"agility", "attack_power", "crit", "hit", "spell_power", "spell_haste", "melee_haste"} {
		if !slices.Contains(ids, want) {
			t.Errorf("KnownStats does not list %q", want)
		}
	}
	// The contract's section 1.4 says "haste"; the engine has two.
	if _, ok := ParseStat("haste"); ok {
		t.Error("ParseStat resolved \"haste\"; the engine has spell_haste and melee_haste and no plain haste")
	}
	if got, _ := ParseStat("crit"); got != proto.Stat_StatCrit {
		t.Errorf("ParseStat(crit) = %v", got)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/request/... -run 'GradedBuff|KnownBuffsListsBoth|WorldBuffs|KnownStats' -v`
Expected: FAIL to compile — `undefined: WorldBuffs`.

- [ ] **Step 3: Teach `buffsFor` the improved suffix**

In `sim/request/buffs.go`, add beside `ErrUnknownBuff`:

```go
// improvedSuffix marks the talented version of a graded buff:
// "battle_shout:improved". The plain id is the plain version, which is
// what it has always meant, so every stored request keeps its meaning.
const improvedSuffix = ":improved"
```

and replace `enableField` with:

```go
// enableField turns on the named field of the first message that has
// one, and reports whether any did. An id carrying improvedSuffix asks
// for the second non-zero value of a graded field, and matches nothing
// else: a boolean buff has no improved form, and answering one with the
// plain buff would apply less than the id promised.
func enableField(targets []protoreflect.ProtoMessage, id string) bool {
	name, improved := strings.CutSuffix(id, improvedSuffix)
	for _, target := range targets {
		msg := target.ProtoReflect()
		fd := msg.Descriptor().Fields().ByName(protoreflect.Name(name))
		if fd == nil {
			continue
		}
		if improved {
			value, ok := gradedValue(fd)
			if !ok {
				return false
			}
			msg.Set(fd, protoreflect.ValueOfEnum(value))
			return true
		}
		switch fd.Kind() {
		case protoreflect.BoolKind:
			msg.Set(fd, protoreflect.ValueOfBool(true))
		case protoreflect.Int32Kind:
			// A count, not a flag: innervates, power infusions, atiesh
			// stacks. One is the smallest thing "on" can mean.
			msg.Set(fd, protoreflect.ValueOfInt32(1))
		case protoreflect.EnumKind:
			msg.Set(fd, protoreflect.ValueOfEnum(onEnumValue(fd)))
		default:
			continue
		}
		return true
	}
	return false
}

// gradedValue is the improved value of a graded field: the SECOND
// non-zero value of its enum. A field with fewer than two is not
// graded and has no improved form.
func gradedValue(fd protoreflect.FieldDescriptor) (protoreflect.EnumNumber, bool) {
	if fd.Kind() != protoreflect.EnumKind {
		return 0, false
	}
	values := fd.Enum().Values()
	var seen int
	for i := 0; i < values.Len(); i++ {
		n := values.Get(i).Number()
		if n == 0 {
			continue
		}
		seen++
		if seen == 2 {
			return n, true
		}
	}
	return 0, false
}
```

Add `"strings"` to the import block if it is not there.

- [ ] **Step 4: Publish both forms, the world buffs and the stats**

In `sim/request/vocabulary.go`, extend `buffVocabulary` so that a graded field emits a second entry:

```go
			id := string(fd.Name())
			if claimed[id] {
				continue
			}
			claimed[id] = true
			out = append(out, vocabularyEntry{id: id, field: id, owner: string(desc.Name())})
			// A graded field has a talented version, and the settings
			// bar offers it as a third state rather than a second
			// checkbox, so it needs an id of its own.
			if _, graded := gradedValue(fd); graded {
				out = append(out, vocabularyEntry{id: id + improvedSuffix, field: id, owner: string(desc.Name())})
			}
```

and append to the same file:

```go
// worldBuffFirstField is where IndividualBuffs' world-buff section
// starts. The engine marks the section with a comment and nothing
// machine-readable, so the boundary is this number and
// TestWorldBuffsCoverTheEnginesWorldBuffSection holds the list below to
// it: a world buff the engine gains and this list misses is a failing
// test rather than a buff no player can tick.
const worldBuffFirstField = 7

// WorldBuffs lists the world-buff ids, sorted. They are IndividualBuffs
// fields like any other - buffsFor needs no special case - but the
// settings bar groups them, so the grouping is published rather than
// guessed at from the names.
func WorldBuffs() []string {
	desc := (&proto.IndividualBuffs{}).ProtoReflect().Descriptor()
	fields := desc.Fields()
	out := make([]string, 0, 8)
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if fd.Number() < worldBuffFirstField {
			continue
		}
		out = append(out, string(fd.Name()))
	}
	sort.Strings(out)
	return out
}

// statIDs is the engine's Stat enum, keyed by the id the weights panel
// sends: the value name in lower snake case with the prefix stripped,
// so StatAttackPower is "attack_power". It is built from the descriptor
// rather than written down, so a stat the engine gains is offerable the
// day the pin moves.
var statIDs = func() map[string]proto.Stat {
	out := make(map[string]proto.Stat, len(proto.Stat_name))
	for value, name := range proto.Stat_name {
		out[strcase.Snake(strings.TrimPrefix(name, "Stat"))] = proto.Stat(value)
	}
	return out
}()

// ParseStat maps a stat id onto the engine's enum.
func ParseStat(id string) (proto.Stat, bool) {
	s, ok := statIDs[id]
	return s, ok
}

// KnownStats lists every stat id a weights request may name, sorted.
func KnownStats() []string {
	out := make([]string, 0, len(statIDs))
	for id := range statIDs {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

// statVocabulary names every stat and the engine value it selects.
func statVocabulary() []vocabularyEntry {
	out := make([]vocabularyEntry, 0, len(statIDs))
	for id, s := range statIDs {
		out = append(out, vocabularyEntry{id: id, field: s.String(), owner: "Stat"})
	}
	return out
}
```

- [ ] **Step 5: Add the two sections to `IDsMarkdown`**

In `IDsMarkdown`, replace the sentence "A boolean field is turned on, a count is set to one, and a graded field (the engine's `TristateEffect`) is set to its plain version; the settings bar has no id for the improved form yet." with:

```
A boolean field is turned on, a count is set to one, and a graded field
(the engine's `TristateEffect`) has two ids: the plain name is the
plain version and `<id>:improved` is the talented one. A buff that is
not graded has no `:improved` form and naming one is an error.
```

Then, after the Buffs table, insert:

```go
	b.WriteString(`
### World buffs

The settings bar groups these separately: they are ` + "`IndividualBuffs`" + `
fields like any other and take the same ids, but a player ticks them
as a group — the Dire Maul tribute buffs, the Zandalar and Warchief's
world enchantments, Rallying Cry, Songflower, Sayge's fortune.

`)
	for _, id := range WorldBuffs() {
		fmt.Fprintf(&b, "- `%s`\n", id)
	}
```

and, after the Professions table, append:

```go
	b.WriteString(`
## Stats

` + "`api.SimRequest.Weights`" + `'s ` + "`stats`" + ` and ` + "`reference`" + ` are these ids: the
engine's ` + "`Stat`" + ` enum value names in lower snake case with the ` + "`Stat`" + `
prefix stripped. There is no plain ` + "`haste`" + `: Forever's engine carries
` + "`spell_haste`" + ` and ` + "`melee_haste`" + ` and they are different stats. The
reference stat is normalised to exactly 1.0 and must be one of the
stats being weighed.

| id | engine enum |
| --- | --- |
`)
	for _, e := range sortedByID(statVocabulary()) {
		fmt.Fprintf(&b, "| `%s` | Stat.%s |\n", e.id, e.field)
	}
```

- [ ] **Step 6: Regenerate IDS.md and run the tests**

```bash
cd sim && go run ./internal/genids && cd ..
go test ./sim/request/... -v
```

Expected: `wrote request/IDS.md`, then PASS — including `TestIDsMarkdownIsCommitted`.

- [ ] **Step 7: Vet, format, commit**

```bash
go vet ./sim/request/...
gofmt -w sim/request/buffs.go sim/request/vocabulary.go sim/request/vocabulary_test.go
git add sim/request/buffs.go sim/request/vocabulary.go sim/request/vocabulary_test.go sim/request/IDS.md
git commit -m "feat(sim): graded buff ids, the world-buff grouping and the stat vocabulary

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 12: Item rows `sim/bulk` can read without touching a protobuf

`sim/request` and `sim/adapter` are the only two packages that speak protobuf; `sim/bulk` must not become a third. So `sim/internal/simdb` grows a plain-Go view of the rows the planner reads.

**Files:**
- Create: `sim/internal/simdb/items.go`
- Create: `sim/internal/simdb/items_test.go`

**Interfaces:**
- Consumes: the embedded `proto.SimDatabase` and the four `SimItem` fields Task 2 pinned.
- Produces:
  - `simdb.Item{ID, Name, Slots, ArmorType, WeaponType, HandType, Classes, RequiredLevel, Unique, Faction, SuffixOptions}`
  - `func simdb.Lookup(id int) (Item, bool)`
  - `func simdb.Len() (int, error)`
  - `simdb.HandTwo`, `HandMain`, `HandOff`, `HandOne`, `HandNone`
  - `simdb.FactionAlliance`, `FactionHorde`

- [ ] **Step 1: Write the failing test**

Create `sim/internal/simdb/items_test.go`:

```go
package simdb

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// Every slot name a row can produce must be a slot the envelope and
// sim/request know, or an expansion would place an item somewhere
// nothing can equip it.
func TestEverySlotNameIsInTheVocabulary(t *testing.T) {
	n, err := Len()
	if err != nil {
		t.Fatal(err)
	}
	if n == 0 {
		t.Fatal("the embedded table is empty")
	}
	db, err := load()
	if err != nil {
		t.Fatal(err)
	}
	var placeable int
	for _, row := range db.Items {
		item, ok := Lookup(int(row.Id))
		if !ok {
			t.Fatalf("item %d is in the table and Lookup does not find it", row.Id)
		}
		if len(item.Slots) > 0 {
			placeable++
		}
		for _, slot := range item.Slots {
			if !slices.Contains(api.GearSlots, slot) {
				t.Errorf("item %d (%s) is placeable in %q, which is not a gear slot", item.ID, item.Name, slot)
			}
		}
	}
	if placeable == 0 {
		t.Error("no item in the build is placeable in any slot")
	}
	t.Logf("items=%d placeable=%d", n, placeable)
}

// The engine's own eligibility rules, re-expressed: a ring and a
// trinket go in either of two slots, a one-hander in either hand, a
// two-hander and a main-hand-only weapon in the main hand, an off-hand
// in the off hand.
func TestSlotsFollowTheEnginesEligibilityRules(t *testing.T) {
	cases := []struct {
		name  string
		row   *proto.SimItem
		slots []string
	}{
		{"a helm", &proto.SimItem{Type: proto.ItemType_ItemTypeHead}, []string{"head"}},
		{"a ring", &proto.SimItem{Type: proto.ItemType_ItemTypeFinger}, []string{"finger1", "finger2"}},
		{"a trinket", &proto.SimItem{Type: proto.ItemType_ItemTypeTrinket}, []string{"trinket1", "trinket2"}},
		{"a bow", &proto.SimItem{Type: proto.ItemType_ItemTypeRanged}, []string{"ranged"}},
		{"a two-hander", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeTwoHand}, []string{"main_hand"}},
		{"a main-hand", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeMainHand}, []string{"main_hand"}},
		{"an off-hand", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeOffHand}, []string{"off_hand"}},
		{"a one-hander", &proto.SimItem{Type: proto.ItemType_ItemTypeWeapon, HandType: proto.HandType_HandTypeOneHand}, []string{"main_hand", "off_hand"}},
		{"an unknown type", &proto.SimItem{Type: proto.ItemType_ItemTypeUnknown}, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := itemFrom(c.row)
			if !slices.Equal(got.Slots, c.slots) {
				t.Errorf("slots = %v, want %v", got.Slots, c.slots)
			}
		})
	}
}

func TestRowFieldsCross(t *testing.T) {
	got := itemFrom(&proto.SimItem{
		Id:                  19352,
		Name:                "Vis'kag the Bloodletter",
		Type:                proto.ItemType_ItemTypeWeapon,
		HandType:            proto.HandType_HandTypeOneHand,
		WeaponType:          proto.WeaponType_WeaponTypeSword,
		ArmorType:           proto.ArmorType_ArmorTypeUnknown,
		ClassAllowlist:      []proto.Class{proto.Class_ClassWarrior, proto.Class_ClassRogue},
		RequiredLevel:       60,
		Unique:              true,
		FactionRestriction:  proto.UIItem_FACTION_RESTRICTION_HORDE_ONLY,
		RandomSuffixOptions: []int32{1805, 1806},
	})
	if got.ID != 19352 || got.Name != "Vis'kag the Bloodletter" {
		t.Errorf("identity: %+v", got)
	}
	if got.HandType != HandOne || got.WeaponType != "sword" {
		t.Errorf("weapon: %+v", got)
	}
	if !slices.Equal(got.Classes, []string{"warrior", "rogue"}) {
		t.Errorf("classes = %v", got.Classes)
	}
	if got.RequiredLevel != 60 || !got.Unique || got.Faction != FactionHorde {
		t.Errorf("restrictions: %+v", got)
	}
	if !slices.Equal(got.SuffixOptions, []int{1805, 1806}) {
		t.Errorf("suffixes = %v", got.SuffixOptions)
	}
}

func TestLookupMissesCleanly(t *testing.T) {
	if _, ok := Lookup(0); ok {
		t.Error("item 0 resolved")
	}
	if _, ok := Lookup(-1); ok {
		t.Error("a negative id resolved")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/internal/simdb/... -run 'SlotName|SlotsFollow|RowFields|LookupMisses' -v`
Expected: FAIL to compile — `undefined: Lookup`.

- [ ] **Step 3: Write `sim/internal/simdb/items.go`**

```go
package simdb

// The item rows the planner reads.
//
// sim/bulk decides which sims run, and to do that it needs to know what
// an item is: where it goes, who can wear it, whether two of it may be
// worn at once. Those facts are in the same embedded database the
// engine loads, and this file is the view of them that is NOT a
// protobuf.
//
// That matters because sim/request and sim/adapter are the only two
// packages in the product that touch one, and a third would be a third
// place the engine's schema leaks into our own vocabulary. So every
// enum here is a string in OUR words - the planner's slot names, class
// slugs, "two_hand" rather than HandTypeTwoHand - and the mapping
// happens once, here.

import (
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
)

// Hand types, in our words.
const (
	HandNone = ""
	HandOne  = "one_hand"
	HandTwo  = "two_hand"
	HandMain = "main_hand"
	HandOff  = "off_hand"
)

// Faction restrictions, in our words. The empty string is "either".
const (
	FactionAny      = ""
	FactionAlliance = "alliance"
	FactionHorde    = "horde"
)

// Item is one row, in the planner's vocabulary.
type Item struct {
	ID   int
	Name string
	// Slots is every gear slot this item can be equipped in, in the
	// order the planner tries them: a ring is finger1 then finger2, a
	// one-hander is main_hand then off_hand. Empty means nothing can
	// equip it, which is most of a client's item table.
	Slots []string
	// ArmorType is "cloth", "leather", "mail", "plate" or "".
	ArmorType string
	// WeaponType is "sword", "axe", "shield" and the rest, or "".
	WeaponType string
	// HandType is one of the Hand constants.
	HandType string
	// Classes are the class slugs allowed to use it. Empty means every
	// class.
	Classes []string
	// RequiredLevel is 0 when the item has none.
	RequiredLevel int
	// Unique says only one may be equipped at a time.
	Unique bool
	// Faction is one of the Faction constants.
	Faction string
	// SuffixOptions are the random suffixes this item can roll.
	SuffixOptions []int
}

// slotsByType is the engine's own eligibility table (core.database.go's
// itemTypeToSlotsMap), in our slot names. Weapons are absent because a
// weapon's slot cannot be decided from its type alone.
var slotsByType = map[proto.ItemType][]string{
	proto.ItemType_ItemTypeHead:     {"head"},
	proto.ItemType_ItemTypeNeck:     {"neck"},
	proto.ItemType_ItemTypeShoulder: {"shoulder"},
	proto.ItemType_ItemTypeBack:     {"back"},
	proto.ItemType_ItemTypeChest:    {"chest"},
	proto.ItemType_ItemTypeWrist:    {"wrist"},
	proto.ItemType_ItemTypeHands:    {"hands"},
	proto.ItemType_ItemTypeWaist:    {"waist"},
	proto.ItemType_ItemTypeLegs:     {"legs"},
	proto.ItemType_ItemTypeFeet:     {"feet"},
	proto.ItemType_ItemTypeFinger:   {"finger1", "finger2"},
	proto.ItemType_ItemTypeTrinket:  {"trinket1", "trinket2"},
	proto.ItemType_ItemTypeRanged:   {"ranged"},
}

// slotsByHand is the weapon half of the same table.
var slotsByHand = map[proto.HandType][]string{
	proto.HandType_HandTypeTwoHand:  {"main_hand"},
	proto.HandType_HandTypeMainHand: {"main_hand"},
	proto.HandType_HandTypeOffHand:  {"off_hand"},
	proto.HandType_HandTypeOneHand:  {"main_hand", "off_hand"},
}

var handNames = map[proto.HandType]string{
	proto.HandType_HandTypeUnknown:  HandNone,
	proto.HandType_HandTypeOneHand:  HandOne,
	proto.HandType_HandTypeTwoHand:  HandTwo,
	proto.HandType_HandTypeMainHand: HandMain,
	proto.HandType_HandTypeOffHand:  HandOff,
}

var factionNames = map[proto.UIItem_FactionRestriction]string{
	proto.UIItem_FACTION_RESTRICTION_UNSPECIFIED:   FactionAny,
	proto.UIItem_FACTION_RESTRICTION_ALLIANCE_ONLY: FactionAlliance,
	proto.UIItem_FACTION_RESTRICTION_HORDE_ONLY:    FactionHorde,
}

// trimmed is an enum value name in our words: the value name without
// its enum's prefix, lower snake case. ArmorTypeMail is "mail",
// WeaponTypeTwoHandedSword is "two_handed_sword", ClassWarrior is
// "warrior". The unknown value of every one of these enums is 0 and
// reads as "", which is what "this item has no armor type" means.
func trimmed(name, prefix string) string {
	name = strings.TrimPrefix(name, prefix)
	if name == "Unknown" || name == "" {
		return ""
	}
	return strcase.Snake(name)
}

// itemFrom maps one engine row into our vocabulary.
func itemFrom(row *proto.SimItem) Item {
	slots := slotsByType[row.Type]
	if row.Type == proto.ItemType_ItemTypeWeapon {
		slots = slotsByHand[row.HandType]
	}
	classes := make([]string, 0, len(row.ClassAllowlist))
	for _, c := range row.ClassAllowlist {
		classes = append(classes, trimmed(c.String(), "Class"))
	}
	suffixes := make([]int, 0, len(row.RandomSuffixOptions))
	for _, s := range row.RandomSuffixOptions {
		suffixes = append(suffixes, int(s))
	}
	out := Item{
		ID:            int(row.Id),
		Name:          row.Name,
		Slots:         slots,
		ArmorType:     trimmed(row.ArmorType.String(), "ArmorType"),
		WeaponType:    trimmed(row.WeaponType.String(), "WeaponType"),
		HandType:      handNames[row.HandType],
		RequiredLevel: int(row.RequiredLevel),
		Unique:        row.Unique,
		Faction:       factionNames[row.FactionRestriction],
	}
	if len(classes) > 0 {
		out.Classes = classes
	}
	if len(suffixes) > 0 {
		out.SuffixOptions = suffixes
	}
	return out
}

// items is the whole table, indexed, built once. The map is never
// handed out: Lookup returns a copy, whose only slices are the shared
// read-only ones built above, so a caller cannot retune the build's
// item table for the rest of the process.
var items = sync.OnceValues(func() (map[int]Item, error) {
	db, err := load()
	if err != nil {
		return nil, err
	}
	out := make(map[int]Item, len(db.Items))
	for _, row := range db.Items {
		out[int(row.Id)] = itemFrom(row)
	}
	return out, nil
})

// Lookup is one item of the active build. A miss is an item the build
// does not carry, which sim/bulk refuses rather than sims.
func Lookup(id int) (Item, bool) {
	table, err := items()
	if err != nil {
		return Item{}, false
	}
	it, ok := table[id]
	return it, ok
}

// Len is how many items the build carries. It is what a caller uses to
// prove the table loaded before it starts asking about ids.
func Len() (int, error) {
	table, err := items()
	if err != nil {
		return 0, err
	}
	return len(table), nil
}
```

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/internal/simdb/... -v`
Expected: PASS, with a log line naming the item and placeable counts.

- [ ] **Step 5: Vet, format, commit**

```bash
go vet ./sim/internal/simdb/...
gofmt -w sim/internal/simdb/items.go sim/internal/simdb/items_test.go
git add sim/internal/simdb/items.go sim/internal/simdb/items_test.go
git commit -m "feat(sim): a protobuf-free view of the build's item rows for the planner

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 13: The build's enchant table

`SimEnchant` carries an effect id and a stat array and nothing about where an enchant may go, so enchant fit cannot come from `simdb`. The contract puts it in `data/builds/<build>/enchants.json` (section 6.2), which is a build's file and is loaded rather than embedded — exactly as `simconsumes.json` already is.

**Files:**
- Create: `sim/bulk/doc.go`
- Create: `sim/bulk/enchants.go`
- Create: `sim/bulk/enchants_test.go`
- Create: `sim/bulk/testdata/enchants.excerpt.json`

**Interfaces:**
- Consumes: `simdb.Item` (Task 12).
- Produces:
  - `bulk.Enchants`, `func bulk.LoadEnchants(io.Reader) (*Enchants, error)`
  - `func (e *Enchants) Fits(effectID int, slot string, item simdb.Item, class string) error`
  - `func (e *Enchants) Len() int`
  - `bulk.ErrUnknownEnchant`, `bulk.ErrEnchantDoesNotFit`, `bulk.ErrNoEnchantTable`
  - `bulk.Options{Enchants *Enchants}`

- [ ] **Step 1: Write the fixture**

Create `sim/bulk/testdata/enchants.excerpt.json` — a ten-row excerpt in the contract's section 6.2 shape. Copy ten real rows out of `data/builds/<build>/enchants.json` once the data lane has published it; until then, this hand-written excerpt is the fixture and the loader is tested against it:

```json
[
  { "id": 1503, "name": "Enchant Weapon - Crusader", "icon": "inv_misc_note_01",
    "slots": ["main_hand", "off_hand"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 2568, "name": "Enchant 2H Weapon - Superior Impact", "icon": "inv_misc_note_01",
    "slots": ["main_hand"], "item_types": ["two_hand"], "classes": [], "stats": {}, "phase": 1 },
  { "id": 1506, "name": "Libram of Rapidity", "icon": "inv_misc_book_11",
    "slots": ["head", "legs"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 849, "name": "Enchant Cloak - Greater Resistance", "icon": "inv_misc_note_01",
    "slots": ["back"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 1891, "name": "Enchant Chest - Greater Stats", "icon": "inv_misc_note_01",
    "slots": ["chest"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 1885, "name": "Enchant Bracer - Superior Strength", "icon": "inv_misc_note_01",
    "slots": ["wrist"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 927, "name": "Enchant Gloves - Agility", "icon": "inv_misc_note_01",
    "slots": ["hands"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 911, "name": "Enchant Boots - Greater Agility", "icon": "inv_misc_note_01",
    "slots": ["feet"], "item_types": [], "classes": [], "stats": {}, "phase": 1 },
  { "id": 2503, "name": "Enchant Shield - Greater Stamina", "icon": "inv_misc_note_01",
    "slots": ["off_hand"], "item_types": ["shield"], "classes": [], "stats": {}, "phase": 1 },
  { "id": 2646, "name": "Mana Oil", "icon": "inv_potion_99",
    "slots": ["main_hand"], "item_types": [], "classes": ["mage", "priest", "warlock"], "stats": {}, "phase": 2 }
]
```

- [ ] **Step 2: Write the failing test**

Create `sim/bulk/enchants_test.go`:

```go
package bulk

import (
	"errors"
	"os"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

func testEnchants(t *testing.T) *Enchants {
	t.Helper()
	f, err := os.Open("testdata/enchants.excerpt.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	table, err := LoadEnchants(f)
	if err != nil {
		t.Fatal(err)
	}
	if table.Len() != 10 {
		t.Fatalf("the excerpt carries %d enchants, want 10", table.Len())
	}
	return table
}

func TestEnchantFit(t *testing.T) {
	table := testEnchants(t)
	twoHander := simdb.Item{ID: 1, HandType: simdb.HandTwo, WeaponType: "two_handed_axe"}
	oneHander := simdb.Item{ID: 2, HandType: simdb.HandOne, WeaponType: "sword"}
	shield := simdb.Item{ID: 3, WeaponType: "shield"}
	helm := simdb.Item{ID: 4}

	cases := []struct {
		name    string
		effect  int
		slot    string
		item    simdb.Item
		class   string
		wantErr error
	}{
		{name: "a weapon enchant on a weapon", effect: 1503, slot: "main_hand", item: oneHander, class: "warrior"},
		{name: "a weapon enchant in the off hand", effect: 1503, slot: "off_hand", item: oneHander, class: "warrior"},
		{name: "a weapon enchant on a helm", effect: 1503, slot: "head", item: helm, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a two-hand enchant on a two-hander", effect: 2568, slot: "main_hand", item: twoHander, class: "warrior"},
		{name: "a two-hand enchant on a one-hander", effect: 2568, slot: "main_hand", item: oneHander, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a shield enchant on a shield", effect: 2503, slot: "off_hand", item: shield, class: "warrior"},
		{name: "a shield enchant on a one-hander", effect: 2503, slot: "off_hand", item: oneHander, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "a class-restricted oil for a mage", effect: 2646, slot: "main_hand", item: oneHander, class: "mage"},
		{name: "a class-restricted oil for a warrior", effect: 2646, slot: "main_hand", item: oneHander, class: "warrior", wantErr: ErrEnchantDoesNotFit},
		{name: "an enchant the build has never heard of", effect: 999999, slot: "head", item: helm, class: "warrior", wantErr: ErrUnknownEnchant},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := table.Fits(c.effect, c.slot, c.item, c.class)
			if !errors.Is(err, c.wantErr) {
				t.Errorf("Fits = %v, want %v", err, c.wantErr)
			}
		})
	}
}

// No table means no named enchant can be checked. Saying so beats
// accepting an enchant that may not exist in this build.
func TestNoTableRefusesANamedEnchant(t *testing.T) {
	var table *Enchants
	err := table.Fits(1503, "main_hand", simdb.Item{HandType: simdb.HandOne}, "warrior")
	if !errors.Is(err, ErrNoEnchantTable) {
		t.Errorf("Fits without a table = %v, want ErrNoEnchantTable", err)
	}
	if table.Len() != 0 {
		t.Error("a nil table has a length")
	}
}

func TestLoadEnchantsRefusesRubbish(t *testing.T) {
	for _, body := range []string{``, `[]`, `[{"name":"no id"}]`, `{"not":"an array"}`} {
		if _, err := LoadEnchants(stringsReader(body)); err == nil {
			t.Errorf("LoadEnchants accepted %q", body)
		}
	}
}
```

Add a tiny helper at the bottom of the file:

```go
func stringsReader(s string) *strings.Reader { return strings.NewReader(s) }
```

with `"strings"` imported.

- [ ] **Step 3: Run the test to verify it fails**

Run: `go test ./sim/bulk/... -v`
Expected: FAIL — no such package.

- [ ] **Step 4: Write `sim/bulk/doc.go`**

```go
// Package bulk decides WHICH sims run.
//
// Top Gear, Droptimizer and talent compare are one question - a base
// character, a set of substitutions, a ranked answer - so they are one
// planner. Expand turns a request's candidates into every valid
// combination; Plan turns those into the first stage's requests; Rank
// scores a finished stage and returns either the next one or the
// finished SimResult.
//
// Nothing here runs a sim. The browser's worker pool runs each stage's
// requests through the sharding it already has, and forever-sim runs
// them in a loop; both call these three functions and neither does any
// statistics of its own. That is the point: the ladder, the cuts, the
// deltas and the within-error grouping are written once, in Go, and
// compiled into both artifacts, so the two lanes cannot rank the same
// candidates differently.
//
// This package touches no protobuf. Item facts come from
// sim/internal/simdb's plain-Go view and enchant facts from the
// build's enchants.json, loaded by the caller.
package bulk
```

- [ ] **Step 5: Write `sim/bulk/enchants.go`**

```go
package bulk

// The build's enchant table.
//
// data/builds/<build>/enchants.json is the data lane's list of every
// enchant the build carries, with the slots and item types it may go
// on. It is not in the sim database: the engine's SimEnchant is an
// effect id and a stat array, because the engine is told where an
// enchant went rather than asked whether it could go there. The
// planner has to ask.
//
// The file belongs to a build, so it is loaded and passed in, the way
// sim/request takes the consumable table - embedding it would pin one
// build into both artifacts.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

var (
	// ErrUnknownEnchant is returned for an effect id the build's table
	// does not carry.
	ErrUnknownEnchant = errors.New("bulk: unknown enchant")
	// ErrEnchantDoesNotFit is returned when the enchant exists and
	// cannot go on that item in that slot.
	ErrEnchantDoesNotFit = errors.New("bulk: the enchant does not fit")
	// ErrNoEnchantTable is returned when a candidate names an enchant
	// and no table was loaded. Accepting it would let a request carry
	// an enchant this build has never had.
	ErrNoEnchantTable = errors.New("bulk: no enchant table is loaded")
)

// Enchant is one row of the build's table.
type Enchant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	// Slots are the gear slots this enchant may be applied in.
	Slots []string `json:"slots"`
	// ItemTypes narrows it further where the slot is not enough: a
	// two-hand enchant and a shield enchant both go in a hand slot.
	// Empty means every item that fits the slot. The vocabulary is
	// simdb's: a hand type ("two_hand") or a weapon type ("shield").
	ItemTypes []string `json:"item_types"`
	// Classes are the class slugs allowed to use it; empty is every
	// class.
	Classes []string           `json:"classes"`
	Stats   map[string]float64 `json:"stats"`
	Phase   int                `json:"phase"`
}

// Enchants is a build's enchant table. The zero value is unusable; use
// LoadEnchants. A nil table answers every Fits with ErrNoEnchantTable.
type Enchants struct {
	byID map[int]Enchant
}

// LoadEnchants reads a build's enchants.json.
func LoadEnchants(r io.Reader) (*Enchants, error) {
	var rows []Enchant
	if err := json.NewDecoder(r).Decode(&rows); err != nil {
		return nil, fmt.Errorf("bulk: enchants.json is not readable: %w", err)
	}
	if len(rows) == 0 {
		return nil, errors.New("bulk: enchants.json carries no enchants")
	}
	out := &Enchants{byID: make(map[int]Enchant, len(rows))}
	for _, row := range rows {
		if row.ID == 0 {
			return nil, fmt.Errorf("bulk: enchants.json has a row with no id: %+v", row)
		}
		out.byID[row.ID] = row
	}
	return out, nil
}

// Len is how many enchants the table carries.
func (e *Enchants) Len() int {
	if e == nil {
		return 0
	}
	return len(e.byID)
}

// Lookup is one enchant.
func (e *Enchants) Lookup(id int) (Enchant, bool) {
	if e == nil {
		return Enchant{}, false
	}
	row, ok := e.byID[id]
	return row, ok
}

// Fits reports whether an enchant may go on this item in this slot for
// this class, and says why not when it may not.
func (e *Enchants) Fits(effectID int, slot string, item simdb.Item, class string) error {
	if e == nil {
		return fmt.Errorf("%w, so enchant %d cannot be checked; pass one in Options, from data/builds/<build>/enchants.json", ErrNoEnchantTable, effectID)
	}
	row, ok := e.byID[effectID]
	if !ok {
		return fmt.Errorf("%w: %d", ErrUnknownEnchant, effectID)
	}
	if !slices.Contains(row.Slots, slot) {
		return fmt.Errorf("%w: %q goes in %v, not %q", ErrEnchantDoesNotFit, row.Name, row.Slots, slot)
	}
	if len(row.ItemTypes) > 0 && !slices.Contains(row.ItemTypes, item.HandType) && !slices.Contains(row.ItemTypes, item.WeaponType) {
		return fmt.Errorf("%w: %q goes on %v, and item %d is a %s %s", ErrEnchantDoesNotFit, row.Name, row.ItemTypes, item.ID, item.HandType, item.WeaponType)
	}
	if len(row.Classes) > 0 && !slices.Contains(row.Classes, class) {
		return fmt.Errorf("%w: %q is for %v, not a %s", ErrEnchantDoesNotFit, row.Name, row.Classes, class)
	}
	return nil
}

// Options are the inputs the planner needs that the module cannot
// embed, the way sim/request.Options carries the consumable table.
type Options struct {
	// Enchants resolves a candidate's named enchant. Nil means a
	// candidate may only INHERIT the equipped item's enchant; naming
	// one fails at the boundary rather than carrying an enchant this
	// build may not have.
	Enchants *Enchants
}
```

- [ ] **Step 6: Run the tests**

Run: `go test ./sim/bulk/... -v`
Expected: PASS.

- [ ] **Step 7: Vet, format, commit**

```bash
go vet ./sim/bulk/...
gofmt -w sim/bulk
git add sim/bulk
git commit -m "feat(sim): the planner package and the build's enchant table

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 14: `Expand`, part one — placements, eligibility and enchant inheritance

**Files:**
- Create: `sim/bulk/expand.go`
- Create: `sim/bulk/expand_test.go`

**Interfaces:**
- Consumes: `api.SimRequest`, `api.BulkSpec`, `api.Candidate`, `api.GearSlot`, `api.Substitution` (Tasks 3, 7); `simdb.Lookup`, `simdb.Item` (Task 12); `bulk.Options`, `bulk.Enchants` (Task 13).
- Produces:
  - `bulk.Combination{Request api.SimRequest, Substitutions []api.Substitution}`
  - `func bulk.Expand(req api.SimRequest) ([]Combination, error)`
  - `func bulk.ExpandWith(req api.SimRequest, opt Options) ([]Combination, error)`
  - `bulk.ErrUnknownItem`, `bulk.ErrNotBulk`
  - unexported `placement{Slot, Gear, Candidate}` and `placements(req, opt)` which Task 15 builds combinations from

- [ ] **Step 1: Write the failing test**

Create `sim/bulk/expand_test.go`:

```go
package bulk

import (
	"errors"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// The real item ids below are the warrior fixture's: they are in the
// active build, which is what makes the eligibility rules answerable
// at all. If one of them ever leaves the build,
// TestFixtureGearResolves in sim/internal/simdb fails first and names
// it; replace it here with another row of the same kind.
const (
	itemHelm       = 12640  // head
	itemOneHander  = 19352  // one-hand sword
	itemTwoHander  = 19019  // two-hand sword
	itemOffHand    = 17075  // off-hand
	itemRing       = 19382  // finger
	itemTrinket    = 19406  // trinket
	enchantWeapon  = 1503   // Crusader: main_hand, off_hand
	enchantTwoHand = 2568   // two-hand only
	enchantHead    = 1506   // head, legs
)

// base is a fury warrior with a full set of the fixture's own gear, so
// every slot the tests substitute into is already filled.
func base() api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          "warrior-fury",
		Source:        api.CharacterSource{Kind: api.SourceManual},
		Character: api.CharacterSpec{
			Name:    "Thrall",
			Race:    "orc",
			Class:   "warrior",
			Level:   api.SimLevel,
			Talents: "30305001302-05050005525010051",
			Gear: []api.GearSlot{
				{Slot: "head", ItemID: itemHelm, Enchant: enchantHead},
				{Slot: "main_hand", ItemID: itemOneHander, Enchant: enchantWeapon},
				{Slot: "off_hand", ItemID: itemOffHand},
				{Slot: "finger1", ItemID: itemRing},
				{Slot: "trinket1", ItemID: itemTrinket},
			},
		},
		Encounter:  api.DefaultEncounter(),
		Iterations: 3000,
		RandomSeed: 7,
	}
}

// withBulk is base with a bulk block that names the given candidates.
func withBulk(mode string, candidates ...api.Candidate) api.SimRequest {
	req := base()
	req.Bulk = &api.BulkSpec{
		Mode:       mode,
		Precision:  api.PrecisionNormal,
		Cap:        api.Caps[api.LaneServer],
		Candidates: candidates,
	}
	return req
}

func candidate(slot string, id int) api.Candidate {
	return api.Candidate{Slot: slot, ItemID: id, Origin: api.OriginBag}
}

// slotsOf lists the slots one combination substitutes into, sorted, so
// a case can assert a shape without depending on iteration order.
func slotsOf(c Combination) []string {
	out := make([]string, 0, len(c.Substitutions))
	for _, s := range c.Substitutions {
		if s.Kind == api.SubstitutionItem {
			out = append(out, s.Slot)
		}
	}
	slices.Sort(out)
	return out
}

// gearAt is the item the combination's request has in a slot.
func gearAt(c Combination, slot string) api.GearSlot {
	for _, g := range c.Request.Character.Gear {
		if g.Slot == slot {
			return g
		}
	}
	return api.GearSlot{}
}

func TestExpandRefusesARequestWithNoBulkBlock(t *testing.T) {
	if _, err := Expand(base()); !errors.Is(err, ErrNotBulk) {
		t.Errorf("Expand of a plain run = %v, want ErrNotBulk", err)
	}
}

func TestExpandRefusesAnItemTheBuildDoesNotCarry(t *testing.T) {
	_, err := Expand(withBulk(api.KindGear, candidate("head", 999999999)))
	if !errors.Is(err, ErrUnknownItem) {
		t.Errorf("Expand = %v, want ErrUnknownItem", err)
	}
}

// A candidate with no slot goes wherever it fits, which is what makes
// "try this ring in both slots" and "try this sword in either hand"
// one rule rather than two.
func TestACandidateWithNoSlotGoesWhereverItFits(t *testing.T) {
	cases := []struct {
		name  string
		item  int
		slots []string
	}{
		{"a ring", itemRing + 1, []string{"finger1", "finger2"}},
		{"a trinket", itemTrinket + 1, []string{"trinket1", "trinket2"}},
		{"a one-hander", itemOneHander, []string{"main_hand", "off_hand"}},
		{"a two-hander", itemTwoHander, []string{"main_hand"}},
		{"a helm", itemHelm, []string{"head"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Expand(withBulk(api.KindDrops, api.Candidate{ItemID: c.item, Origin: "drop:raid:mc:lucifron"}))
			if err != nil {
				t.Fatal(err)
			}
			var slots []string
			for _, combo := range got {
				slots = append(slots, slotsOf(combo)...)
			}
			slices.Sort(slots)
			if !slices.Equal(slots, c.slots) {
				t.Errorf("placed in %v, want %v", slots, c.slots)
			}
		})
	}
}

func TestACandidateNamingASlotGoesOnlyThere(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, candidate("finger2", itemRing+1)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || !slices.Equal(slotsOf(got[0]), []string{"finger2"}) {
		t.Fatalf("got %d combinations in %v", len(got), slotsOf(got[0]))
	}
}

func TestACandidateNamingASlotItDoesNotFitIsSkipped(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, candidate("head", itemOneHander)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("a sword was placed on a head: %v", slotsOf(got[0]))
	}
}

// A candidate the character cannot equip is SKIPPED, not refused: a
// boss's loot table is the whole table, and Droptimizer sends it.
func TestCandidatesTheCharacterCannotEquipAreSkipped(t *testing.T) {
	cases := []struct {
		name string
		edit func(*api.SimRequest)
		item int
	}{
		{"another class's item", func(r *api.SimRequest) { r.Character.Class = "mage"; r.Spec = "mage-frost" }, itemOneHander},
		{"a locked slot", func(r *api.SimRequest) { r.Bulk.Locked = []string{"finger1", "finger2"} }, itemRing + 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := withBulk(api.KindDrops, api.Candidate{ItemID: c.item, Origin: "drop:raid:mc:lucifron"})
			c.edit(&req)
			got, err := Expand(req)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != 0 {
				t.Errorf("%d combinations survived: %v", len(got), slotsOf(got[0]))
			}
		})
	}
}

// An enchant of zero inherits the equipped item's, where it fits. That
// is what stops Top Gear ranking an unenchanted upgrade below an
// enchanted item the player already wears.
func TestEnchantInheritance(t *testing.T) {
	table := testEnchants(t)
	cases := []struct {
		name        string
		candidate   api.Candidate
		wantEnchant int
	}{
		{"inherits the equipped weapon enchant", api.Candidate{Slot: "main_hand", ItemID: itemOneHander + 1, Origin: api.OriginBag}, enchantWeapon},
		{"inherits nothing into an unenchanted slot", api.Candidate{Slot: "off_hand", ItemID: itemOffHand + 1, Origin: api.OriginBag}, 0},
		{"keeps the enchant it was given", api.Candidate{Slot: "head", ItemID: itemHelm, Enchant: enchantHead, Origin: api.OriginBag}, enchantHead},
		{"does not inherit an enchant that does not fit", api.Candidate{Slot: "main_hand", ItemID: itemTwoHander, Origin: api.OriginBag}, enchantWeapon},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ExpandWith(withBulk(api.KindGear, c.candidate), Options{Enchants: table})
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != 1 {
				t.Fatalf("got %d combinations", len(got))
			}
			slot := got[0].Substitutions[0].Slot
			if g := gearAt(got[0], slot); g.Enchant != c.wantEnchant {
				t.Errorf("%s carries enchant %d, want %d", slot, g.Enchant, c.wantEnchant)
			}
		})
	}
}

// A two-hand enchant cannot be inherited onto a one-hander, so a
// candidate that would have inherited one carries none instead. The
// case above covers the reverse; this is the one that used to silently
// apply an enchant the item cannot hold.
func TestAnInheritedEnchantThatDoesNotFitIsDropped(t *testing.T) {
	req := base()
	req.Character.Gear = []api.GearSlot{{Slot: "main_hand", ItemID: itemTwoHander, Enchant: enchantTwoHand}}
	req.Bulk = &api.BulkSpec{Mode: api.KindGear, Precision: api.PrecisionNormal, Cap: 400,
		Candidates: []api.Candidate{candidate("main_hand", itemOneHander)}}
	got, err := ExpandWith(req, Options{Enchants: testEnchants(t)})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d combinations", len(got))
	}
	if g := gearAt(got[0], "main_hand"); g.Enchant != 0 {
		t.Errorf("a two-hand enchant was inherited onto a one-hander: %d", g.Enchant)
	}
}

// A named enchant needs a table. Without one the candidate is refused
// rather than run with an enchant the build may not carry.
func TestANamedEnchantWithoutATableIsRefused(t *testing.T) {
	c := candidate("head", itemHelm)
	c.Enchant = enchantHead
	_, err := Expand(withBulk(api.KindGear, c))
	if !errors.Is(err, ErrNoEnchantTable) {
		t.Errorf("Expand = %v, want ErrNoEnchantTable", err)
	}
}

// A named enchant that does not fit is the page sending something it
// should not, and it is refused rather than dropped: the player asked
// for that enchant and would otherwise read a ranking of something
// else.
func TestANamedEnchantThatDoesNotFitIsRefused(t *testing.T) {
	c := candidate("head", itemHelm)
	c.Enchant = enchantWeapon
	_, err := ExpandWith(withBulk(api.KindGear, c), Options{Enchants: testEnchants(t)})
	if !errors.Is(err, ErrEnchantDoesNotFit) {
		t.Errorf("Expand = %v, want ErrEnchantDoesNotFit", err)
	}
}

// The substitution the result carries says everything the page shows
// in a chip: the slot it took, the item, the enchant it ended up with,
// the suffix and where the candidate came from.
func TestTheSubstitutionCarriesTheWholeChip(t *testing.T) {
	c := api.Candidate{Slot: "finger2", ItemID: itemRing + 1, Suffix: 1805, Origin: "drop:raid:mc:lucifron"}
	got, err := Expand(withBulk(api.KindDrops, c))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d combinations", len(got))
	}
	s := got[0].Substitutions[0]
	if s.Kind != api.SubstitutionItem || s.Slot != "finger2" || s.ItemID != c.ItemID || s.Suffix != 1805 || s.Origin != c.Origin {
		t.Errorf("substitution = %+v", s)
	}
	if g := gearAt(got[0], "finger2"); g.ItemID != c.ItemID || g.Suffix != 1805 {
		t.Errorf("the request does not carry the substitution: %+v", g)
	}
}

// Every combination is a runnable request: the bulk block is gone (a
// stage request is a plain run) and nothing else about the character
// moved.
func TestACombinationsRequestIsAPlainRun(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, candidate("head", itemHelm)))
	if err != nil {
		t.Fatal(err)
	}
	if got[0].Request.Bulk != nil {
		t.Error("a combination's request still carries a bulk block")
	}
	if got[0].Request.Spec != base().Spec || got[0].Request.Character.Talents != base().Character.Talents {
		t.Error("a combination changed something other than the gear")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/bulk/... -run Expand -v`
Expected: FAIL to compile — `undefined: Expand`.

- [ ] **Step 3: Write `sim/bulk/expand.go`**

```go
package bulk

// Expansion: which substitutions are even possible.
//
// The rules are the ones the page states in words on the Top Gear
// page, and each has a table test: an item goes only where its
// inventory type allows; rings and trinkets are tried in both slots;
// a one-hander is tried in either hand; a class, level or faction the
// character does not have means the item is never simmed; a locked
// slot is never substituted; and an enchant is inherited from the
// equipped item where it fits.
//
// A candidate the character cannot equip is SKIPPED rather than
// refused. Droptimizer sends a boss's whole loot table, plate and
// cloth together, and refusing the request would make the tool
// unusable; the page shows the count it planned, so a skipped
// candidate is visible rather than silent.
//
// A candidate that is malformed - an item the build has never heard
// of, a named enchant that cannot go where it was sent - IS refused.
// That is the page sending something it should not, and running it
// would answer a different question from the one asked.

import (
	"errors"
	"fmt"
	"slices"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
)

var (
	// ErrNotBulk is returned when a request with no bulk block is
	// handed to the planner.
	ErrNotBulk = errors.New("bulk: the request carries no bulk block")
	// ErrUnknownItem is returned for a candidate the active build has
	// no row for. It is a refusal rather than a skip: an id nothing
	// resolves is a client bug, and the engine would die mid-run on it.
	ErrUnknownItem = errors.New("bulk: the build has no such item")
)

// Combination is one substitution set and the request that runs it.
type Combination struct {
	Request       api.SimRequest     `json:"request"`
	Substitutions []api.Substitution `json:"substitutions"`
}

// placement is one candidate in one slot: the gear row it produces and
// the substitution chip that describes it.
type placement struct {
	Slot         string
	Gear         api.GearSlot
	Substitution api.Substitution
}

// Expand lists every valid combination for req, with no build tables.
// See ExpandWith.
func Expand(req api.SimRequest) ([]Combination, error) {
	return ExpandWith(req, Options{})
}

// ExpandWith lists every valid combination for req.
func ExpandWith(req api.SimRequest, opt Options) ([]Combination, error) {
	if req.Bulk == nil {
		return nil, ErrNotBulk
	}
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("bulk: %w", err)
	}
	places, err := placements(req, opt)
	if err != nil {
		return nil, err
	}
	return combinations(req, places)
}

// baseGear indexes the character's equipped gear by slot.
func baseGear(req api.SimRequest) map[string]api.GearSlot {
	out := make(map[string]api.GearSlot, len(req.Character.Gear))
	for _, g := range req.Character.Gear {
		out[g.Slot] = g
	}
	return out
}

// placements turns every candidate into the (slot, gear) pairs it can
// produce. A candidate that fits nowhere contributes none.
func placements(req api.SimRequest, opt Options) ([]placement, error) {
	equipped := baseGear(req)
	locked := req.Bulk.Locked
	var out []placement
	for _, c := range req.Bulk.Candidates {
		item, ok := simdb.Lookup(c.ItemID)
		if !ok {
			return nil, fmt.Errorf("%w: %d", ErrUnknownItem, c.ItemID)
		}
		if !usable(item, req.Character) {
			continue
		}
		for _, slot := range item.Slots {
			if c.Slot != "" && c.Slot != slot {
				continue
			}
			if slices.Contains(locked, slot) {
				continue
			}
			enchant, err := enchantFor(c, slot, item, req.Character.Class, equipped, opt)
			if err != nil {
				return nil, err
			}
			out = append(out, placement{
				Slot: slot,
				Gear: api.GearSlot{Slot: slot, ItemID: c.ItemID, Enchant: enchant, Suffix: c.Suffix},
				Substitution: api.Substitution{
					Kind:    api.SubstitutionItem,
					Slot:    slot,
					ItemID:  c.ItemID,
					Enchant: enchant,
					Suffix:  c.Suffix,
					Origin:  c.Origin,
				},
			})
		}
	}
	return out, nil
}

// usable reports whether this character could wear the item at all.
// Class, level and faction are three ways of saying "not for you", and
// all three end the same way: the item is never simmed.
func usable(item simdb.Item, ch api.CharacterSpec) bool {
	if len(item.Classes) > 0 && !slices.Contains(item.Classes, ch.Class) {
		return false
	}
	if item.RequiredLevel > ch.Level {
		return false
	}
	if item.Faction != simdb.FactionAny && item.Faction != factionOf(ch.Race) {
		return false
	}
	return true
}

// hordeRaces are the races that cannot loot an Alliance-only item. The
// list is short and stable - Forever's Skyborne choose a faction at
// creation and the race slug says which - so it lives here rather than
// pulling the build's race table into the planner.
var hordeRaces = []string{"orc", "tauren", "troll", "undead", "windshaper-skyborne"}

// factionOf is the faction a race belongs to.
func factionOf(race string) string {
	if slices.Contains(hordeRaces, race) {
		return simdb.FactionHorde
	}
	return simdb.FactionAlliance
}

// enchantFor resolves a candidate's enchant.
//
// A named enchant is checked and kept, or refused. An unnamed one
// INHERITS the equipped item's for that slot, where that enchant fits
// the new item - which is what Raidbots does and what the fork's
// auto_enchant does - and is dropped where it does not, because a
// two-hand enchant cannot ride onto a one-hander.
//
// Inheritance with no table loaded is not an error: the enchant came
// from the character the player sent, and refusing it would make Top
// Gear unusable wherever the table has not been published. A NAMED
// enchant with no table is an error, because nothing can say whether
// this build has it.
func enchantFor(c api.Candidate, slot string, item simdb.Item, class string, equipped map[string]api.GearSlot, opt Options) (int, error) {
	if c.Enchant != 0 {
		if err := opt.Enchants.Fits(c.Enchant, slot, item, class); err != nil {
			return 0, err
		}
		return c.Enchant, nil
	}
	inherited := equipped[slot].Enchant
	if inherited == 0 {
		return 0, nil
	}
	if err := opt.Enchants.Fits(inherited, slot, item, class); err != nil {
		if errors.Is(err, ErrNoEnchantTable) {
			return inherited, nil
		}
		return 0, nil
	}
	return inherited, nil
}
```

`combinations` is Task 15; write it as a stub that returns one combination per placement so this task's tests can run, and replace it there:

```go
// combinations is Task 15. Until then, one combination per placement.
func combinations(req api.SimRequest, places []placement) ([]Combination, error) {
	out := make([]Combination, 0, len(places))
	for _, p := range places {
		out = append(out, apply(req, []placement{p}, nil, nil))
	}
	return out, nil
}
```

and the shared applier, which Task 15 keeps:

```go
// apply builds one combination: the base character with these
// placements swapped in, optionally a talent loadout and optionally a
// whole gear set, and the chips that describe the change.
//
// The request it produces is a PLAIN RUN - no bulk block - because
// that is what a stage request is, and leaving the block on would make
// every stage request expand again.
func apply(req api.SimRequest, places []placement, loadout *api.TalentLoadout, set *api.GearSet) Combination {
	out := req
	out.Bulk = nil
	out.Weights = nil
	out.TargetError = 0

	gear := baseGear(req)
	if set != nil {
		gear = make(map[string]api.GearSlot, len(set.Gear))
		for _, g := range set.Gear {
			gear[g.Slot] = g
		}
	}
	for _, p := range places {
		gear[p.Slot] = p.Gear
	}
	// In the envelope's own slot order, so two combinations that equip
	// the same set produce byte-identical requests and the seeds line
	// up across stages.
	out.Character.Gear = make([]api.GearSlot, 0, len(gear))
	for _, slot := range api.GearSlots {
		if g, ok := gear[slot]; ok {
			out.Character.Gear = append(out.Character.Gear, g)
		}
	}

	subs := make([]api.Substitution, 0, len(places)+2)
	if set != nil {
		subs = append(subs, api.Substitution{Kind: api.SubstitutionSet, Name: set.Name})
	}
	for _, p := range places {
		subs = append(subs, p.Substitution)
	}
	if loadout != nil {
		out.Character.Talents = loadout.Talents
		subs = append(subs, api.Substitution{
			Kind:    api.SubstitutionTalents,
			Name:    loadout.Name,
			Talents: loadout.Talents,
		})
	}
	return Combination{Request: out, Substitutions: subs}
}
```

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/bulk/... -run Expand -v`
Expected: PASS. The eligibility cases depend on the constants at the top of the file being real rows of the active build. If one is not, the case fails with "placed in [] want [...]". To find real ids, add this test to `sim/bulk/expand_test.go`, run it, read the log, put the ids into the constants, and **delete the test before committing**:

```go
func TestFindIDs(t *testing.T) {
	want := map[string]int{"two_hand": 0, "off_hand": 0, "one_hand": 0}
	for id := 1; id < 400000 && !allFound(want); id++ {
		item, ok := simdb.Lookup(id)
		if !ok {
			continue
		}
		if _, interesting := want[item.HandType]; interesting && want[item.HandType] == 0 {
			want[item.HandType] = id
		}
	}
	t.Logf("%+v", want)
	t.Fail()
}

func allFound(m map[string]int) bool {
	for _, v := range m {
		if v == 0 {
			return false
		}
	}
	return true
}
```

- [ ] **Step 5: Vet, format, commit**

```bash
go vet ./sim/bulk/...
gofmt -w sim/bulk
git add sim/bulk
git commit -m "feat(sim): the planner's placement and eligibility rules, with enchant inheritance

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 15: `Expand`, part two — combination shapes, validity and the cap

**Files:**
- Modify: `sim/bulk/expand.go` (replace the `combinations` stub)
- Modify: `sim/bulk/expand_test.go`

**Interfaces:**
- Consumes: `placement`, `apply`, `baseGear` (Task 14); `api.ErrCapExceeded`, `api.Caps` (Task 3).
- Produces: the real `combinations`, enforcing one item per slot, two-hand versus main-plus-off-hand, duplicate rings and trinkets, same-name rings and trinkets, unique-equipped, the talents and sets dimensions, and the cap.

- [ ] **Step 1: Write the failing test**

Append to `sim/bulk/expand_test.go`:

```go
// Gear mode is every valid combination; drops and talents mode are one
// substitution at a time. That is the whole difference between Top
// Gear and Droptimizer, and it is the mode that says which.
func TestGearModeCombinesAndDropsModeDoesNot(t *testing.T) {
	two := []api.Candidate{candidate("head", itemHelm), candidate("finger2", itemRing+1)}
	gearMode, err := Expand(withBulk(api.KindGear, two...))
	if err != nil {
		t.Fatal(err)
	}
	// head-or-not times finger2-or-not, minus the "neither" case,
	// which is the equipped set and is not a combination.
	if len(gearMode) != 3 {
		t.Errorf("gear mode produced %d combinations, want 3", len(gearMode))
	}
	var both int
	for _, c := range gearMode {
		if len(slotsOf(c)) == 2 {
			both++
		}
	}
	if both != 1 {
		t.Errorf("gear mode produced %d combinations that change both slots, want 1", both)
	}

	dropsCandidates := []api.Candidate{
		{Slot: "head", ItemID: itemHelm, Origin: "drop:raid:mc:lucifron"},
		{Slot: "finger2", ItemID: itemRing + 1, Origin: "drop:raid:mc:lucifron"},
	}
	dropsMode, err := Expand(withBulk(api.KindDrops, dropsCandidates...))
	if err != nil {
		t.Fatal(err)
	}
	if len(dropsMode) != 2 {
		t.Errorf("drops mode produced %d combinations, want 2", len(dropsMode))
	}
	for _, c := range dropsMode {
		if len(slotsOf(c)) != 1 {
			t.Errorf("drops mode changed %d slots at once", len(slotsOf(c)))
		}
	}
}

// Two candidates for one slot are alternatives, never both at once.
func TestTwoCandidatesForOneSlotAreAlternatives(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear,
		candidate("head", itemHelm), candidate("head", itemHelm+1)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d combinations, want 2", len(got))
	}
	for _, c := range got {
		if len(c.Substitutions) != 1 {
			t.Errorf("two items landed in one slot: %+v", c.Substitutions)
		}
	}
}

func TestWeaponShapesAndDuplicates(t *testing.T) {
	cases := []struct {
		name       string
		candidates []api.Candidate
		edit       func(*api.SimRequest)
		wantSlots  [][]string // the combinations, as sorted slot lists
	}{
		{
			name:       "a two-hander competes with main-plus-off-hand",
			candidates: []api.Candidate{{ItemID: itemTwoHander, Origin: api.OriginBag}},
			// The off hand is emptied rather than the combination
			// being thrown away: a two-hander IS a shape, and refusing
			// it would mean Top Gear never ranked one.
			wantSlots: [][]string{{"main_hand"}},
		},
		{
			name:       "a one-hander is tried in both hands",
			candidates: []api.Candidate{{ItemID: itemOneHander + 1, Origin: api.OriginBag}},
			wantSlots:  [][]string{{"main_hand"}, {"off_hand"}},
		},
		{
			name: "dual wield tries both orders",
			candidates: []api.Candidate{
				{ItemID: itemOneHander + 1, Origin: api.OriginBag},
				{ItemID: itemOneHander + 2, Origin: api.OriginBag},
			},
			wantSlots: [][]string{
				{"main_hand"}, {"off_hand"},
				{"main_hand"}, {"off_hand"},
				{"main_hand", "off_hand"}, {"main_hand", "off_hand"},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := withBulk(api.KindGear, c.candidates...)
			if c.edit != nil {
				c.edit(&req)
			}
			got, err := Expand(req)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(c.wantSlots) {
				var shapes [][]string
				for _, combo := range got {
					shapes = append(shapes, slotsOf(combo))
				}
				t.Fatalf("got %d combinations %v, want %d %v", len(got), shapes, len(c.wantSlots), c.wantSlots)
			}
		})
	}
}

// A two-hander in the main hand empties the off hand; nothing may be
// wielded beside it.
func TestATwoHanderEmptiesTheOffHand(t *testing.T) {
	got, err := Expand(withBulk(api.KindGear, api.Candidate{ItemID: itemTwoHander, Origin: api.OriginBag}))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d combinations", len(got))
	}
	if g := gearAt(got[0], "off_hand"); g.ItemID != 0 {
		t.Errorf("a two-hander left item %d in the off hand", g.ItemID)
	}
}

// The engine's own validity rules: one item cannot be in both ring
// slots, nor both trinket slots, and two rings or trinkets sharing a
// name are the same item at two qualities.
func TestDuplicateRingsAndTrinketsAreRefused(t *testing.T) {
	// finger1 already holds itemRing; offering it for finger2 would
	// produce a character wearing two of one ring.
	got, err := Expand(withBulk(api.KindGear, candidate("finger2", itemRing)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("the same ring was equipped twice: %+v", got[0].Request.Character.Gear)
	}

	got, err = Expand(withBulk(api.KindGear, candidate("trinket2", itemTrinket)))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("the same trinket was equipped twice")
	}
}

// Unique-equipped is one at a time anywhere, not one per slot.
func TestUniqueEquippedIsRespected(t *testing.T) {
	// itemUniqueRing must be a row whose Unique flag is set. Find one
	// with the helper below and replace the constant if this id is not
	// unique in the active build.
	req := base()
	req.Character.Gear = append(req.Character.Gear, api.GearSlot{Slot: "finger2", ItemID: itemUniqueRing})
	req.Bulk = &api.BulkSpec{Mode: api.KindGear, Precision: api.PrecisionNormal, Cap: 400,
		Candidates: []api.Candidate{candidate("finger1", itemUniqueRing)}}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("a unique-equipped item was equipped twice: %+v", got[0].Request.Character.Gear)
	}
}

// Talent loadouts are a dimension of their own: in gear mode they
// multiply the gear combinations, and in talents mode they are the
// only candidates.
func TestTalentLoadoutsAreADimension(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Bulk.Talents = []api.TalentLoadout{
		{Name: "Deep Fury", Talents: "30305001302-05050005525010052"},
		{Name: "Two-hand Arms", Talents: "30305001302-05050005525010053"},
	}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	// (the helm, or not) times (own talents, Deep Fury, Two-hand Arms),
	// minus the equipped set: 2 * 3 - 1 = 5.
	if len(got) != 5 {
		t.Errorf("got %d combinations, want 5", len(got))
	}

	only := withBulk(api.KindTalents)
	only.Bulk.Candidates = nil
	only.Bulk.Talents = req.Bulk.Talents
	got, err = Expand(only)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("talents mode produced %d combinations, want 2", len(got))
	}
	for _, c := range got {
		if len(c.Substitutions) != 1 || c.Substitutions[0].Kind != api.SubstitutionTalents {
			t.Errorf("a talents-mode combination substituted gear: %+v", c.Substitutions)
		}
		if c.Request.Character.Talents == base().Character.Talents {
			t.Error("a loadout did not reach the request")
		}
	}
}

// A set replaces every slot at once, so it is an alternative to the
// whole gear product rather than a member of it.
func TestASetReplacesEverySlot(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Bulk.Sets = []api.GearSet{{Name: "my AQ set", Gear: []api.GearSlot{
		{Slot: "head", ItemID: itemHelm + 1},
		{Slot: "main_hand", ItemID: itemTwoHander},
	}}}
	got, err := Expand(req)
	if err != nil {
		t.Fatal(err)
	}
	// the helm on its own, and the set: two combinations.
	if len(got) != 2 {
		t.Fatalf("got %d combinations, want 2", len(got))
	}
	var set *Combination
	for i, c := range got {
		if len(c.Substitutions) == 1 && c.Substitutions[0].Kind == api.SubstitutionSet {
			set = &got[i]
		}
	}
	if set == nil {
		t.Fatal("no combination substituted the set")
	}
	if len(set.Request.Character.Gear) != 2 {
		t.Errorf("the set did not replace the whole character's gear: %+v", set.Request.Character.Gear)
	}
	if set.Substitutions[0].Name != "my AQ set" {
		t.Errorf("the chip does not name the set: %+v", set.Substitutions[0])
	}
}

// The cap is a refusal with both numbers, never a silent trim: a
// ranking of a subset nobody chose looks exactly like a ranking.
func TestTheCapRefusesRatherThanTrims(t *testing.T) {
	req := withBulk(api.KindGear,
		candidate("head", itemHelm), candidate("head", itemHelm+1),
		candidate("finger2", itemRing+1), candidate("trinket2", itemTrinket+1))
	req.Bulk.Cap = 3
	_, err := Expand(req)
	var capped api.ErrCapExceeded
	if !errors.As(err, &capped) {
		t.Fatalf("Expand = %v, want ErrCapExceeded", err)
	}
	if capped.Cap != 3 || capped.Combinations <= 3 {
		t.Errorf("ErrCapExceeded = %+v", capped)
	}
}
```

Add the constant beside the others at the top of the file, and a helper that finds one so the number can be replaced without guessing:

```go
// itemUniqueRing is any unique-equipped ring in the active build.
// TestFindAUniqueRing prints candidates when this one stops being one.
const itemUniqueRing = 19432

func TestFindAUniqueRing(t *testing.T) {
	item, ok := simdb.Lookup(itemUniqueRing)
	if ok && item.Unique && slices.Contains(item.Slots, "finger1") {
		return
	}
	t.Errorf("item %d is no longer a unique ring in this build; pick another from the build's table", itemUniqueRing)
}
```

with `"github.com/jhunthrop/foreversixty/sim/internal/simdb"` imported by the test file.

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/bulk/... -run 'GearModeCombines|TwoCandidates|WeaponShapes|TwoHanderEmpties|Duplicate|Unique|TalentLoadouts|ASet|TheCap' -v`
Expected: FAIL — the stub `combinations` produces one combination per placement, so the gear-mode, talents, sets and cap cases all fail.

- [ ] **Step 3: Replace the stub in `sim/bulk/expand.go`**

```go
// combinations builds every combination the mode allows.
//
// Gear mode is a product: each slot offers "keep what is equipped" or
// one of its candidates, and the talent dimension offers "the
// character's own build" or one of the loadouts. A set is not part of
// that product - it replaces every slot at once - so each set is its
// own arm, crossed with the talent dimension only.
//
// Drops and talents mode are one substitution at a time: a Droptimizer
// answer is "this boss's sword is worth 41 DPS", and a product would
// answer a question nobody asked and blow the cap doing it.
//
// The equipped set is never a combination: Plan runs it separately, in
// every stage, so that every delta is paired.
func combinations(req api.SimRequest, places []placement) ([]Combination, error) {
	var out []Combination
	if req.Bulk.Mode == api.KindGear {
		out = gearCombinations(req, places)
	} else {
		out = singleCombinations(req, places)
	}
	out = slices.DeleteFunc(out, func(c Combination) bool {
		return !valid(c.Request.Character.Gear)
	})
	if len(out) > req.Bulk.Cap {
		return nil, api.ErrCapExceeded{Cap: req.Bulk.Cap, Combinations: len(out)}
	}
	return out, nil
}

// singleCombinations is one substitution at a time: every placement on
// its own, then every talent loadout on its own.
func singleCombinations(req api.SimRequest, places []placement) []Combination {
	out := make([]Combination, 0, len(places)+len(req.Bulk.Talents))
	for _, p := range places {
		out = append(out, apply(req, []placement{p}, nil, nil))
	}
	for i := range req.Bulk.Talents {
		out = append(out, apply(req, nil, &req.Bulk.Talents[i], nil))
	}
	return out
}

// gearCombinations is the product.
func gearCombinations(req api.SimRequest, places []placement) []Combination {
	// One bucket per slot, in the envelope's slot order so the product
	// is enumerated the same way every time and two runs of the same
	// request produce the same combination order.
	bySlot := map[string][]placement{}
	for _, p := range places {
		bySlot[p.Slot] = append(bySlot[p.Slot], p)
	}
	slots := make([]string, 0, len(bySlot))
	for _, slot := range api.GearSlots {
		if len(bySlot[slot]) > 0 {
			slots = append(slots, slot)
		}
	}

	// Every choice of at most one placement per slot, including none.
	sets := [][]placement{nil}
	for _, slot := range slots {
		next := make([][]placement, 0, len(sets)*(len(bySlot[slot])+1))
		for _, chosen := range sets {
			next = append(next, chosen)
			for _, p := range bySlot[slot] {
				next = append(next, append(append([]placement(nil), chosen...), p))
			}
		}
		sets = next
	}

	// The talent dimension: the character's own build, then each
	// loadout. A nil loadout is "their own".
	loadouts := make([]*api.TalentLoadout, 0, len(req.Bulk.Talents)+1)
	loadouts = append(loadouts, nil)
	for i := range req.Bulk.Talents {
		loadouts = append(loadouts, &req.Bulk.Talents[i])
	}

	out := make([]Combination, 0, len(sets)*len(loadouts))
	for _, chosen := range sets {
		for _, loadout := range loadouts {
			if len(chosen) == 0 && loadout == nil {
				// The equipped set with its own talents is the
				// baseline, and Plan runs it. Including it here would
				// rank the character against itself.
				continue
			}
			out = append(out, apply(req, chosen, loadout, nil))
		}
	}
	// A named set is a whole-gear alternative, so it is its own arm
	// rather than a member of the product above.
	for i := range req.Bulk.Sets {
		for _, loadout := range loadouts {
			out = append(out, apply(req, nil, loadout, &req.Bulk.Sets[i]))
		}
	}
	return out
}

// valid is the engine's own isValidEquipment, over our gear list: no
// two-hander beside an off-hand, no item in both ring or both trinket
// slots, no two rings or trinkets sharing a name (which is the same
// item at two qualities), and nothing unique-equipped worn twice.
//
// It is re-expressed here rather than called because the engine's copy
// works on a protobuf and this package holds no protobuf, and because
// a combination that reached the engine and was refused there would
// cost a whole sim to learn.
func valid(gear []api.GearSlot) bool {
	byID := map[int]int{}
	names := map[string]int{}
	var mainHand, offHand simdb.Item
	for _, g := range gear {
		item, ok := simdb.Lookup(g.ItemID)
		if !ok {
			continue
		}
		byID[g.ItemID]++
		if byID[g.ItemID] > 1 && item.Unique {
			return false
		}
		switch g.Slot {
		case "main_hand":
			mainHand = item
		case "off_hand":
			offHand = item
		case "finger1", "finger2", "trinket1", "trinket2":
			// One item cannot be in both of a pair, and two items of
			// one name are the same thing at two qualities.
			if byID[g.ItemID] > 1 {
				return false
			}
			names[item.Name]++
			if names[item.Name] > 1 {
				return false
			}
		}
	}
	return !(mainHand.HandType == simdb.HandTwo && offHand.ID != 0)
}
```

Update `apply` so a two-hander clears the off hand — replace the placement loop in `apply` with:

```go
	for _, p := range places {
		gear[p.Slot] = p.Gear
		// A two-hander leaves no room for an off-hand. Clearing it is
		// what makes "two-hand versus main-plus-off-hand" two competing
		// shapes rather than one combination the engine would refuse.
		if p.Slot == "main_hand" {
			if item, ok := simdb.Lookup(p.Gear.ItemID); ok && item.HandType == simdb.HandTwo {
				delete(gear, "off_hand")
			}
		}
	}
```

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/bulk/... -v`
Expected: PASS, including Task 14's cases.

- [ ] **Step 5: Vet, format, commit**

```bash
go vet ./sim/bulk/...
gofmt -w sim/bulk
git add sim/bulk
git commit -m "feat(sim): combination shapes, weapon validity, unique-equipped and the cap

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 16: `Plan` — the first stage

**Files:**
- Create: `sim/bulk/plan.go`
- Create: `sim/bulk/plan_test.go`

**Interfaces:**
- Consumes: `ExpandWith`, `Combination`, `Options` (Tasks 13–15); `api.Ladders`, `api.FinalIterations` (Task 3).
- Produces:
  - `bulk.StageRequests{Stage int, Iterations int, Requests []api.SimRequest, Combos []Combination}` with JSON tags `stage`, `iterations`, `requests`, `combos`
  - `func bulk.Plan(req api.SimRequest) (StageRequests, error)`
  - `func bulk.PlanWith(req api.SimRequest, opt Options) (StageRequests, error)`
  - `func bulk.stageRequests(req api.SimRequest, stage, iterations int, combos []Combination) StageRequests` (used by `Rank`)

- [ ] **Step 1: Write the failing test**

Create `sim/bulk/plan_test.go`:

```go
package bulk

import (
	"encoding/json"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

func TestPlanIsTheFirstRungOfTheLadder(t *testing.T) {
	cases := []struct {
		precision  string
		iterations int
	}{
		{api.PrecisionFast, 100},
		{api.PrecisionNormal, 1000},
		{api.PrecisionHigh, 1000},
	}
	for _, c := range cases {
		t.Run(c.precision, func(t *testing.T) {
			req := withBulk(api.KindGear, candidate("head", itemHelm))
			req.Bulk.Precision = c.precision
			final, _ := api.FinalIterations(c.precision)
			req.Iterations = final

			stage, err := Plan(req)
			if err != nil {
				t.Fatal(err)
			}
			if stage.Stage != 1 {
				t.Errorf("stage = %d, want 1", stage.Stage)
			}
			if stage.Iterations != c.iterations {
				t.Errorf("iterations = %d, want %d", stage.Iterations, c.iterations)
			}
			if len(stage.Requests) != len(stage.Combos)+1 {
				t.Fatalf("%d requests and %d combos; the equipped set is the extra one",
					len(stage.Requests), len(stage.Combos))
			}
			for i, r := range stage.Requests {
				if r.Iterations != c.iterations {
					t.Errorf("request %d asks for %d iterations", i, r.Iterations)
				}
				if r.Bulk != nil {
					t.Errorf("request %d still carries a bulk block", i)
				}
				if err := r.ValidatePart(); err != nil {
					t.Errorf("request %d is not runnable: %v", i, err)
				}
			}
		})
	}
}

// Requests[0] is always the equipped set: it is the baseline every
// delta is measured against, and the page and the binary both index it
// by position rather than searching for it.
func TestTheEquippedSetIsRequestZero(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm+1))
	stage, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	equipped := stage.Requests[0]
	if len(equipped.Character.Gear) != len(req.Character.Gear) {
		t.Fatalf("request 0 has %d items, the character has %d", len(equipped.Character.Gear), len(req.Character.Gear))
	}
	for i, g := range req.Character.Gear {
		if equipped.Character.Gear[i] != g {
			t.Errorf("request 0 slot %d is %+v, the character wears %+v", i, equipped.Character.Gear[i], g)
		}
	}
	if equipped.Character.Talents != req.Character.Talents {
		t.Error("request 0 does not carry the character's own talents")
	}
}

// Every request in a stage shares one seed, so the combinations and
// the equipped set see the same rolls and the delta between them is a
// paired comparison rather than two independent samples. Stages differ,
// so a combination that got a lucky stream at 100 iterations does not
// keep it at 1,000.
func TestSeedsArePairedWithinAStageAndDifferBetweenThem(t *testing.T) {
	req := withBulk(api.KindGear, candidate("head", itemHelm), candidate("finger2", itemRing+1))
	first, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	seed := first.Requests[0].RandomSeed
	for i, r := range first.Requests {
		if r.RandomSeed != seed {
			t.Errorf("request %d has seed %d, request 0 has %d; a stage is paired", i, r.RandomSeed, seed)
		}
	}
	second := stageRequests(req, 2, 1000, first.Combos)
	if second.Requests[0].RandomSeed == seed {
		t.Error("stage 2 reuses stage 1's seed")
	}
	// Deterministic: planning twice produces the same seeds, so a
	// re-run of a saved request reproduces the run.
	again, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	if again.Requests[0].RandomSeed != seed {
		t.Error("two plans of one request disagree about the seed")
	}
}

// The stage object crosses the wasm boundary as JSON, so its names are
// the contract's and are pinned here.
func TestStageRequestsJSONNames(t *testing.T) {
	stage, err := Plan(withBulk(api.KindGear, candidate("head", itemHelm)))
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(stage)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"stage", "iterations", "requests", "combos"} {
		if _, ok := got[key]; !ok {
			t.Errorf("a stage does not carry %q", key)
		}
	}
	var back StageRequests
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Stage != stage.Stage || len(back.Requests) != len(stage.Requests) || len(back.Combos) != len(stage.Combos) {
		t.Errorf("round trip lost something: %+v", back)
	}
}

func TestPlanRefusesAPlainRun(t *testing.T) {
	if _, err := Plan(base()); err == nil {
		t.Error("Plan accepted a request with no bulk block")
	}
}

func TestPlanRefusesAnExpansionWithNothingInIt(t *testing.T) {
	// Every candidate is for a class the character is not, so nothing
	// survives eligibility. A plan of zero combinations is a ranking
	// of nothing, and saying so beats running the equipped set alone
	// and calling it a Top Gear.
	req := withBulk(api.KindGear, candidate("head", itemHelm))
	req.Character.Class = "mage"
	req.Spec = "mage-frost"
	if _, err := Plan(req); err == nil {
		t.Error("Plan accepted an expansion with no combinations")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/bulk/... -run Plan -v`
Expected: FAIL to compile — `undefined: Plan`.

- [ ] **Step 3: Write `sim/bulk/plan.go`**

```go
package bulk

// Staging: how many sims at what precision.
//
// The ladder is api.Ladders, so the page can show it before the run
// starts. This file turns one rung into runnable requests.
//
// Two things about the requests matter and are easy to get wrong.
//
// The equipped set is Requests[0] of EVERY stage. It is not an
// optimisation to run it once at the end: a delta is only honest
// against a baseline measured the same way, and a baseline from a
// 100-iteration stage compared against a 3,000-iteration finalist
// would put the whole error of the cheap run into every delta.
//
// Every request in one stage shares a seed. The engine advances its
// seed once per iteration, so two requests with one seed see the same
// rolls, and the difference between them is a PAIRED comparison -
// which is what makes a 41-DPS gain detectable at 1,000 iterations
// instead of 10,000. Stages get different seeds so that a combination
// that got a lucky stream once does not keep it.

import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// StageRequests is one rung: what to run, and what each run means.
type StageRequests struct {
	Stage      int `json:"stage"`
	Iterations int `json:"iterations"`
	// Requests[0] is ALWAYS the equipped set; Requests[1:] are
	// parallel to Combos.
	Requests []api.SimRequest `json:"requests"`
	Combos   []Combination    `json:"combos"`
}

// Plan is the first stage, with no build tables. See PlanWith.
func Plan(req api.SimRequest) (StageRequests, error) {
	return PlanWith(req, Options{})
}

// PlanWith is the first stage: every combination plus the equipped set,
// each as a request at the first rung's iteration count.
func PlanWith(req api.SimRequest, opt Options) (StageRequests, error) {
	combos, err := ExpandWith(req, opt)
	if err != nil {
		return StageRequests{}, err
	}
	if len(combos) == 0 {
		return StageRequests{}, errors.New("bulk: nothing to rank; every candidate was for another class, another faction, a locked slot, or a slot it does not fit")
	}
	ladder, ok := api.Ladders[req.Bulk.Precision]
	if !ok || len(ladder.Iterations) == 0 {
		return StageRequests{}, fmt.Errorf("bulk: no ladder for precision %q", req.Bulk.Precision)
	}
	return stageRequests(req, 1, ladder.Iterations[0], combos), nil
}

// stageRequests builds one stage's runnable requests. Rank calls it for
// every rung after the first.
func stageRequests(req api.SimRequest, stage, iterations int, combos []Combination) StageRequests {
	// A stage's seed is the request's plus the stage number, so a
	// re-run of a saved request reproduces the whole ladder and no two
	// stages share a stream.
	seed := req.RandomSeed + int64(stage)

	equipped := req
	equipped.Bulk = nil
	equipped.Weights = nil
	equipped.TargetError = 0
	equipped.Iterations = iterations
	equipped.RandomSeed = seed

	out := StageRequests{
		Stage:      stage,
		Iterations: iterations,
		Requests:   make([]api.SimRequest, 0, len(combos)+1),
		Combos:     combos,
	}
	out.Requests = append(out.Requests, equipped)
	for _, c := range combos {
		r := c.Request
		r.Iterations = iterations
		r.RandomSeed = seed
		out.Requests = append(out.Requests, r)
	}
	return out
}
```

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/bulk/... -run Plan -v`
Expected: PASS.

- [ ] **Step 5: Run the package, vet, format, commit**

```bash
go test ./sim/bulk/...
go vet ./sim/bulk/...
gofmt -w sim/bulk
git add sim/bulk
git commit -m "feat(sim): the planner's first stage, with the equipped set paired at one seed

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 17: `Rank`, part one — scoring a stage and cutting to the next

**Files:**
- Create: `sim/bulk/rank.go`
- Create: `sim/bulk/rank_test.go`

**Interfaces:**
- Consumes: `StageRequests`, `stageRequests` (Task 16); `api.Ladder`, `api.Cut`, `api.Ladders` (Task 3).
- Produces:
  - `func bulk.Rank(req api.SimRequest, stage StageRequests, results []api.SimResult) (next *StageRequests, final *api.SimResult, err error)`
  - unexported `scored{Combo Combination, DPS, Delta api.Estimate}`, `score(...)`, `applyCut(...)`
  - `bulk.ErrStageMismatch`

- [ ] **Step 1: Write the failing test**

Create `sim/bulk/rank_test.go`:

```go
package bulk

import (
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// resultsFor fabricates one finished result per request, with the DPS
// given. Nothing here runs the engine: staging is arithmetic and is
// tested as arithmetic.
func resultsFor(stage StageRequests, means []float64, stderr float64) []api.SimResult {
	out := make([]api.SimResult, len(stage.Requests))
	for i := range stage.Requests {
		out[i] = api.SimResult{
			EngineVersion: enginever.Version,
			Request:       stage.Requests[i],
			Lane:          api.LaneBrowser,
			IterationsRun: stage.Iterations,
			DPS:           api.Estimate{Mean: means[i], StdDev: stderr * 10, Error: stderr},
		}
	}
	return out
}

// planOf is a stage with n distinguishable candidates: n helms, each a
// real row, so expansion keeps them all in their own slot.
func planOf(t *testing.T, precision string, n int) (api.SimRequest, StageRequests) {
	t.Helper()
	req := base()
	final, _ := api.FinalIterations(precision)
	req.Iterations = final
	req.Bulk = &api.BulkSpec{Mode: api.KindGear, Precision: precision, Cap: api.Caps[api.LaneServer]}
	for i := 0; i < n; i++ {
		req.Bulk.Candidates = append(req.Bulk.Candidates, candidate("head", helmIDs[i]))
	}
	stage, err := Plan(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(stage.Combos) != n {
		t.Fatalf("planned %d combinations, want %d; helmIDs may be stale", len(stage.Combos), n)
	}
	return req, stage
}

func TestRankRefusesResultsThatDoNotMatchTheStage(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	short := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 5)[:3]
	if _, _, err := Rank(req, stage, short); !errors.Is(err, ErrStageMismatch) {
		t.Errorf("Rank = %v, want ErrStageMismatch", err)
	}
}

func TestRankPropagatesAFailedRun(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 4)
	results := resultsFor(stage, []float64{1000, 1010, 1020, 1030, 1040}, 5)
	results[2].Error = "the engine died"
	_, _, err := Rank(req, stage, results)
	if err == nil {
		t.Fatal("Rank ignored a failed run")
	}
}

// Normal is two rungs: 1,000 then 3,000, keeping the top ten plus
// anything still overlapping the tenth.
func TestNormalKeepsTheTopTenPlusTies(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 14)
	// The equipped set, then fourteen combinations 10 DPS apart, so
	// nothing is a tie at an error of 1.
	means := []float64{1000}
	for i := 0; i < 14; i++ {
		means = append(means, float64(1100-10*i))
	}
	next, final, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 2 produced a final result")
	}
	if next.Stage != 2 || next.Iterations != 3000 {
		t.Errorf("next stage is %d at %d iterations", next.Stage, next.Iterations)
	}
	if len(next.Combos) != 10 {
		t.Errorf("kept %d combinations, want the top 10", len(next.Combos))
	}
	if len(next.Requests) != 11 {
		t.Errorf("next stage has %d requests, want 10 plus the equipped set", len(next.Requests))
	}

	// With a large error every candidate overlaps the tenth, so every
	// one survives: the cut keeps the top ten PLUS the ties, and
	// throwing a winner away over noise is what the slack exists to
	// prevent.
	next, _, err = Rank(req, stage, resultsFor(stage, means, 50))
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Combos) != 14 {
		t.Errorf("kept %d combinations at an error of 50, want all 14", len(next.Combos))
	}
}

// Fast is three rungs, and the first keeps a quarter.
func TestFastKeepsAQuarterThenTheTopTen(t *testing.T) {
	req, stage := planOf(t, api.PrecisionFast, 40)
	means := []float64{1000}
	for i := 0; i < 40; i++ {
		means = append(means, float64(1400-10*i))
	}
	if stage.Iterations != 100 {
		t.Fatalf("stage 1 runs %d iterations, want 100", stage.Iterations)
	}
	next, final, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 1 of 3 produced a final result")
	}
	if next.Stage != 2 || next.Iterations != 1000 {
		t.Errorf("next stage is %d at %d iterations", next.Stage, next.Iterations)
	}
	if len(next.Combos) != 10 {
		t.Errorf("kept %d of 40, want a quarter", len(next.Combos))
	}

	means2 := means[:len(next.Requests)]
	third, final, err := Rank(req, *next, resultsFor(*next, means2, 1))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil {
		t.Fatal("stage 2 of 3 produced a final result")
	}
	if third.Stage != 3 || third.Iterations != 3000 {
		t.Errorf("third stage is %d at %d iterations", third.Stage, third.Iterations)
	}
}

// A cut never keeps nothing, however small the fraction.
func TestACutAlwaysKeepsAtLeastOne(t *testing.T) {
	req, stage := planOf(t, api.PrecisionFast, 2)
	next, _, err := Rank(req, stage, resultsFor(stage, []float64{1000, 1100, 900}, 1))
	if err != nil {
		t.Fatal(err)
	}
	if len(next.Combos) < 1 {
		t.Fatal("a quarter of two kept nothing")
	}
}

// The survivors are the BEST ones, in order, and the next stage's
// requests are parallel to them with the equipped set still first.
func TestTheSurvivorsAreTheBestInOrder(t *testing.T) {
	req, stage := planOf(t, api.PrecisionNormal, 14)
	means := []float64{1000}
	for i := 0; i < 14; i++ {
		// Ascending, so the LAST candidate is the best and a planner
		// that kept the first ten would be caught.
		means = append(means, float64(1000+10*i))
	}
	next, _, err := Rank(req, stage, resultsFor(stage, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	best := next.Combos[0].Substitutions[0].ItemID
	if best != helmIDs[13] {
		t.Errorf("the leader is item %d, want %d", best, helmIDs[13])
	}
	if next.Requests[0].Character.Gear[0].ItemID != req.Character.Gear[0].ItemID {
		t.Error("the equipped set is not request 0 of the next stage")
	}
	for i, c := range next.Combos {
		if next.Requests[i+1].Character.Gear[0].ItemID != c.Substitutions[0].ItemID {
			t.Errorf("request %d does not match combo %d", i+1, i)
		}
	}
}
```

Add the helm id list beside the other constants in `expand_test.go`:

```go
// helmIDs are forty head items of the active build, used to plan a
// stage with a known number of distinguishable candidates. Any forty
// heads will do; TestHelmIDsAreHeads keeps them honest.
var helmIDs = []int{ /* forty head item ids from the active build */ }

func TestHelmIDsAreHeads(t *testing.T) {
	if len(helmIDs) < 40 {
		t.Fatalf("helmIDs has %d entries, the rank tests need 40", len(helmIDs))
	}
	seen := map[int]bool{}
	for _, id := range helmIDs {
		item, ok := simdb.Lookup(id)
		if !ok || !slices.Contains(item.Slots, "head") {
			t.Errorf("item %d is not a head of this build", id)
		}
		if seen[id] {
			t.Errorf("helmIDs lists %d twice", id)
		}
		seen[id] = true
	}
}
```

Fill the list by running this once, reading the log, pasting the ids in, and **deleting the finder**:

```go
func TestFindHelms(t *testing.T) {
	var out []int
	for id := 1; id < 400000 && len(out) < 40; id++ {
		if item, ok := simdb.Lookup(id); ok && slices.Contains(item.Slots, "head") && !item.Unique {
			out = append(out, id)
		}
	}
	t.Logf("%v", out)
	t.Fail()
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/bulk/... -run Rank -v`
Expected: FAIL to compile — `undefined: Rank`.

- [ ] **Step 3: Write the scoring and cutting half of `sim/bulk/rank.go`**

```go
package bulk

// Ranking: what a finished stage means, and what runs next.
//
// Everything statistical about a bulk run is here, and nowhere else.
// The page does not recompute a delta, a standard error or a tie; it
// renders what this returns. That is what stops the browser and the
// server from disagreeing about who won.

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// ErrStageMismatch is returned when the results do not line up with the
// stage they claim to answer. Scoring them anyway would attribute one
// combination's DPS to another's substitutions.
var ErrStageMismatch = errors.New("bulk: the results do not match the stage")

// scored is one combination with this stage's numbers.
type scored struct {
	Combo Combination
	DPS   api.Estimate
	Delta api.Estimate
}

// Rank scores a finished stage and returns either the next stage or,
// after the last rung, the finished result. Exactly one of next and
// final is non-nil.
func Rank(req api.SimRequest, stage StageRequests, results []api.SimResult) (*StageRequests, *api.SimResult, error) {
	if req.Bulk == nil {
		return nil, nil, ErrNotBulk
	}
	ladder, ok := api.Ladders[req.Bulk.Precision]
	if !ok {
		return nil, nil, fmt.Errorf("bulk: no ladder for precision %q", req.Bulk.Precision)
	}
	ranked, equipped, err := score(stage, results)
	if err != nil {
		return nil, nil, err
	}
	if stage.Stage < 1 || stage.Stage > len(ladder.Iterations) {
		return nil, nil, fmt.Errorf("%w: stage %d, and the %s ladder has %d", ErrStageMismatch, stage.Stage, req.Bulk.Precision, len(ladder.Iterations))
	}
	if stage.Stage == len(ladder.Iterations) {
		final := finalResult(req, stage, ranked, equipped, results[0])
		return nil, &final, nil
	}
	kept := applyCut(ranked, ladder.Cuts[stage.Stage-1])
	combos := make([]Combination, 0, len(kept))
	for _, s := range kept {
		combos = append(combos, s.Combo)
	}
	next := stageRequests(req, stage.Stage+1, ladder.Iterations[stage.Stage], combos)
	return &next, nil, nil
}

// score pairs each result with its combination and computes the delta
// against the equipped set.
//
// The delta's error is the two errors added in quadrature. The runs
// share a seed, so they are positively correlated and the true error of
// the difference is SMALLER than this; the envelope carries no
// covariance, so the conservative figure is the one reported. Saying a
// gain is "within error" when it is real costs a player nothing; the
// reverse costs them an upgrade.
func score(stage StageRequests, results []api.SimResult) ([]scored, api.Estimate, error) {
	if len(results) != len(stage.Requests) {
		return nil, api.Estimate{}, fmt.Errorf("%w: %d results for %d requests", ErrStageMismatch, len(results), len(stage.Requests))
	}
	if len(stage.Combos) != len(stage.Requests)-1 {
		return nil, api.Estimate{}, fmt.Errorf("%w: %d combinations and %d requests; the equipped set is the extra one", ErrStageMismatch, len(stage.Combos), len(stage.Requests))
	}
	for i, res := range results {
		if res.Error != "" {
			return nil, api.Estimate{}, fmt.Errorf("bulk: run %d of stage %d failed: %s", i, stage.Stage, res.Error)
		}
		if res.Aborted {
			return nil, api.Estimate{}, fmt.Errorf("bulk: run %d of stage %d was stopped before it finished", i, stage.Stage)
		}
	}
	equipped := results[0].DPS
	out := make([]scored, 0, len(stage.Combos))
	for i, combo := range stage.Combos {
		dps := results[i+1].DPS
		out = append(out, scored{
			Combo: combo,
			DPS:   dps,
			Delta: api.Estimate{
				Mean:   dps.Mean - equipped.Mean,
				StdDev: math.Hypot(dps.StdDev, equipped.StdDev),
				Error:  math.Hypot(dps.Error, equipped.Error),
			},
		})
	}
	// Best first, and ties broken by the substitution chip so two runs
	// of one request produce the same order.
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].DPS.Mean != out[j].DPS.Mean {
			return out[i].DPS.Mean > out[j].DPS.Mean
		}
		return chipKey(out[i].Combo) < chipKey(out[j].Combo)
	})
	return out, equipped, nil
}

// chipKey is a stable ordering for two combinations of equal DPS.
func chipKey(c Combination) string {
	var b []byte
	for _, s := range c.Substitutions {
		b = fmt.Appendf(b, "%s|%s|%d|%d|%d|%s;", s.Kind, s.Slot, s.ItemID, s.Enchant, s.Suffix, s.Name)
	}
	return string(b)
}

// applyCut keeps the cut's survivors plus anything still overlapping
// the last of them.
//
// The slack is what makes a 100-iteration stage safe. At that count a
// combination's error is large, so the order is mostly noise; cutting
// strictly to the top quarter would throw away the eventual winner
// roughly as often as not. Keeping anything whose interval still
// reaches the cut's costs a few more sims and cannot lose a winner to
// one unlucky stage.
func applyCut(ranked []scored, cut api.Cut) []scored {
	if len(ranked) == 0 {
		return ranked
	}
	keep := cut.Top
	if cut.Fraction > 0 {
		keep = int(math.Round(float64(len(ranked)) * cut.Fraction))
	}
	keep = max(keep, 1)
	if keep >= len(ranked) {
		return ranked
	}
	// The last survivor's lower bound is the bar; anything whose upper
	// bound still reaches it is a tie with the cut.
	bar := ranked[keep-1].DPS.Mean - cut.SlackSE*ranked[keep-1].DPS.Error
	for keep < len(ranked) {
		s := ranked[keep]
		if s.DPS.Mean+cut.SlackSE*s.DPS.Error < bar {
			break
		}
		keep++
	}
	return ranked[:keep]
}
```

`finalResult` is Task 18; write it as a stub that returns a result with only `Request` and `Equipped` set so this task's tests compile, and replace it there:

```go
// finalResult is Task 18.
func finalResult(req api.SimRequest, stage StageRequests, ranked []scored, equipped api.Estimate, base api.SimResult) api.SimResult {
	out := base
	out.Request = req
	out.Equipped = &equipped
	return out
}
```

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/bulk/... -run Rank -v`
Expected: PASS.

- [ ] **Step 5: Vet, format, commit**

```bash
go vet ./sim/bulk/...
gofmt -w sim/bulk
git add sim/bulk
git commit -m "feat(sim): scoring a stage, paired deltas, and the cut to the next rung

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 18: `Rank`, part two — the finished result, and the within-error groups

**Files:**
- Modify: `sim/bulk/rank.go` (replace the `finalResult` stub)
- Modify: `sim/bulk/rank_test.go`

**Interfaces:**
- Consumes: `scored`, `Rank` (Task 17); `api.Combo`, `api.Stage`, `api.SimResult` (Task 7).
- Produces: the real `finalResult`, filling `Combos`, `Equipped`, `Stages`, `DPS`, `IterationsRun` and `Summary`, with `Combo.Group` set by overlapping delta intervals.

- [ ] **Step 1: Write the failing test**

Append to `sim/bulk/rank_test.go`:

```go
// The last rung produces a SimResult, not another stage.
func TestTheLastRungProducesTheResult(t *testing.T) {
	req, first := planOf(t, api.PrecisionNormal, 4)
	means := []float64{1000, 1100, 1050, 1020, 990}
	next, final, err := Rank(req, first, resultsFor(first, means, 2))
	if err != nil {
		t.Fatal(err)
	}
	if final != nil || next == nil {
		t.Fatal("the first of two rungs finished the run")
	}

	last := *next
	lastMeans := means[:len(last.Requests)]
	next, final, err = Rank(req, last, resultsFor(last, lastMeans, 2))
	if err != nil {
		t.Fatal(err)
	}
	if next != nil || final == nil {
		t.Fatal("the last rung did not finish the run")
	}
	if final.Equipped == nil || final.Equipped.Mean != 1000 {
		t.Errorf("equipped = %+v, want the baseline 1000", final.Equipped)
	}
	if final.DPS.Mean != 1000 {
		t.Errorf("the result's own DPS is %v; a bulk result's headline number is the equipped set's", final.DPS.Mean)
	}
	if final.IterationsRun != last.Iterations {
		t.Errorf("iterations_run = %d, want the last stage's %d", final.IterationsRun, last.Iterations)
	}
	if final.Request.Bulk == nil {
		t.Error("the result does not carry the request that produced it")
	}
	if len(final.Combos) != len(last.Combos) {
		t.Fatalf("%d combos in the result and %d in the last stage", len(final.Combos), len(last.Combos))
	}
	if final.Combos[0].DPS.Mean != 1100 || final.Combos[0].Delta.Mean != 100 {
		t.Errorf("the leader is %+v", final.Combos[0])
	}
	if len(final.Combos[0].Substitutions) == 0 {
		t.Error("the leader carries no substitution chips")
	}
	// Every stage the ladder ran, in order, with what it ran.
	if len(final.Stages) != 2 {
		t.Fatalf("stages = %+v, want two", final.Stages)
	}
	if final.Stages[0].Iterations != 1000 || final.Stages[0].Combos != 4 {
		t.Errorf("stage 1 = %+v", final.Stages[0])
	}
	if final.Stages[1].Iterations != 3000 || final.Stages[1].Combos != len(last.Combos) {
		t.Errorf("stage 2 = %+v", final.Stages[1])
	}
	// The summary is the equipped set's: a bulk report renders the
	// baseline character's breakdown beside the ranking.
	if final.Summary.EngineVersion == "" {
		t.Error("the result carries no summary")
	}
}

// "Within error" is an overlap of delta intervals with the group's
// leader, and the page ranks a group the same. Getting this wrong
// makes noise look like a decision.
func TestWithinErrorGroups(t *testing.T) {
	cases := []struct {
		name   string
		means  []float64 // the equipped set first
		stderr float64
		groups []int // per combination, best first
	}{
		{
			name:   "three clearly separated",
			means:  []float64{1000, 1100, 1050, 1010},
			stderr: 1,
			groups: []int{0, 1, 2},
		},
		{
			name:   "all one answer",
			means:  []float64{1000, 1100, 1099, 1098},
			stderr: 40,
			groups: []int{0, 0, 0},
		},
		{
			name:   "two tied, then one apart",
			means:  []float64{1000, 1100, 1098, 1000},
			stderr: 2,
			groups: []int{0, 0, 1},
		},
		{
			name:   "a single combination is its own group",
			means:  []float64{1000, 1100},
			stderr: 2,
			groups: []int{0},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req, first := planOf(t, api.PrecisionNormal, len(c.means)-1)
			next, _, err := Rank(req, first, resultsFor(first, c.means, c.stderr))
			if err != nil {
				t.Fatal(err)
			}
			last := *next
			_, final, err := Rank(req, last, resultsFor(last, c.means[:len(last.Requests)], c.stderr))
			if err != nil {
				t.Fatal(err)
			}
			if len(final.Combos) != len(c.groups) {
				t.Fatalf("%d combos, want %d", len(final.Combos), len(c.groups))
			}
			for i, want := range c.groups {
				if final.Combos[i].Group != want {
					t.Errorf("combo %d (%v DPS) is group %d, want %d",
						i, final.Combos[i].DPS.Mean, final.Combos[i].Group, want)
				}
			}
		})
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/bulk/... -run 'TheLastRung|WithinError' -v`
Expected: FAIL — the stub fills no combos.

- [ ] **Step 3: Replace the stub in `sim/bulk/rank.go`**

```go
// finalResult turns the last stage into the answer.
//
// The headline DPS and the summary are the EQUIPPED set's: a Top Gear
// report shows the character the player has, with the ranking beside
// it, and a headline taken from the winner would tell them they
// already do 1,100 DPS.
func finalResult(req api.SimRequest, stage StageRequests, ranked []scored, equipped api.Estimate, base api.SimResult) api.SimResult {
	out := base
	out.Request = req
	out.DPS = equipped
	out.Equipped = &equipped
	out.IterationsRun = stage.Iterations
	out.Combos = make([]api.Combo, 0, len(ranked))
	for _, s := range ranked {
		out.Combos = append(out.Combos, api.Combo{
			Substitutions: s.Combo.Substitutions,
			DPS:           s.DPS,
			Delta:         s.Delta,
		})
	}
	group(out.Combos)
	out.Stages = stagesOf(req, stage, len(ranked))
	return out
}

// group numbers the within-error bands.
//
// A band starts at the first combination not yet in one, and every
// following combination whose delta interval still overlaps THAT
// leader's joins it. The intervals are one standard error either side,
// which is the interval the page draws.
//
// Comparing each row to its group's leader rather than to its
// predecessor is deliberate: chaining would walk a long tail of
// overlapping neighbours into one band whose ends do not overlap at
// all, and the page would rank a real 40-DPS gap as a tie.
func group(combos []api.Combo) {
	current := -1
	var leader api.Estimate
	for i := range combos {
		d := combos[i].Delta
		if current < 0 || d.Mean+d.Error < leader.Mean-leader.Error {
			current++
			leader = d
		}
		combos[i].Group = current
	}
}

// stagesOf is what the ladder actually ran: every rung's iteration
// count, and how many combinations it ran. The counts before the last
// are recomputed from the cuts rather than remembered, because Rank is
// called once per stage and holds no state between calls - the page
// and the binary each drive the loop and neither carries a history.
func stagesOf(req api.SimRequest, last StageRequests, finalCombos int) []api.Stage {
	ladder := api.Ladders[req.Bulk.Precision]
	out := make([]api.Stage, 0, len(ladder.Iterations))
	for i, iterations := range ladder.Iterations {
		stage := api.Stage{Iterations: iterations}
		switch {
		case i == len(ladder.Iterations)-1:
			stage.Combos = finalCombos
		case i == last.Stage-1:
			stage.Combos = len(last.Combos)
		}
		out = append(out, stage)
	}
	return out
}
```

`stagesOf` as written fills only the last rung and the one it was handed, which is wrong for a three-rung ladder. Make `Rank` carry the history instead: add a field to `StageRequests` and fill it in `stageRequests`.

In `sim/bulk/plan.go`, add to `StageRequests`:

```go
	// Ran is every stage before this one, so the finished result can
	// say what the ladder did without anybody holding a history. It
	// travels with the stage across the wasm boundary, which is the
	// only place it could live: Rank is called once per stage and the
	// page drives the loop.
	Ran []api.Stage `json:"ran,omitempty"`
```

and in `stageRequests`, after `out := StageRequests{...}`, nothing changes — instead have `Rank` pass the history through. In `Rank`, replace the `next := stageRequests(...)` line with:

```go
	next := stageRequests(req, stage.Stage+1, ladder.Iterations[stage.Stage], combos)
	next.Ran = append(append([]api.Stage(nil), stage.Ran...),
		api.Stage{Iterations: stage.Iterations, Combos: len(stage.Combos)})
```

and replace `stagesOf` with:

```go
// stagesOf is what the ladder actually ran: the stages the request
// carried through the loop, plus the one that just finished.
func stagesOf(stage StageRequests) []api.Stage {
	return append(append([]api.Stage(nil), stage.Ran...),
		api.Stage{Iterations: stage.Iterations, Combos: len(stage.Combos)})
}
```

and its call site in `finalResult`:

```go
	out.Stages = stagesOf(stage)
```

Drop the now-unused `finalCombos` parameter and the `req` parameter from `stagesOf`.

- [ ] **Step 4: Run the tests**

Run: `go test ./sim/bulk/... -v`
Expected: PASS.

- [ ] **Step 5: Add the fast-ladder history case**

Append to `sim/bulk/rank_test.go`:

```go
// Three rungs, so the result's stage list is three entries and each
// says what that rung actually ran.
func TestAThreeRungLadderReportsEveryStage(t *testing.T) {
	req, first := planOf(t, api.PrecisionFast, 40)
	means := []float64{1000}
	for i := 0; i < 40; i++ {
		means = append(means, float64(1400-10*i))
	}
	second, _, err := Rank(req, first, resultsFor(first, means, 1))
	if err != nil {
		t.Fatal(err)
	}
	third, _, err := Rank(req, *second, resultsFor(*second, means[:len(second.Requests)], 1))
	if err != nil {
		t.Fatal(err)
	}
	_, final, err := Rank(req, *third, resultsFor(*third, means[:len(third.Requests)], 1))
	if err != nil {
		t.Fatal(err)
	}
	if final == nil {
		t.Fatal("the third rung did not finish the run")
	}
	want := []api.Stage{
		{Iterations: 100, Combos: 40},
		{Iterations: 1000, Combos: len(second.Combos)},
		{Iterations: 3000, Combos: len(third.Combos)},
	}
	if len(final.Stages) != 3 {
		t.Fatalf("stages = %+v", final.Stages)
	}
	for i, w := range want {
		if final.Stages[i] != w {
			t.Errorf("stage %d = %+v, want %+v", i+1, final.Stages[i], w)
		}
	}
}
```

- [ ] **Step 6: Run, vet, format, commit**

```bash
go test ./sim/bulk/...
go vet ./sim/bulk/...
gofmt -w sim/bulk
git add sim/bulk
git commit -m "feat(sim): the finished bulk result, its stage history and the within-error groups

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 19: Stat weights — the engine request and the result mapping

**Files:**
- Create: `sim/request/weights.go`
- Create: `sim/request/weights_test.go`
- Modify: `sim/adapter/adapter.go`
- Modify: `sim/adapter/adapter_test.go`

**Interfaces:**
- Consumes: `api.WeightsSpec`, `api.StatWeight` (Task 5); `request.ParseStat` (Task 11); `request.BuildWith`.
- Produces:
  - `func request.BuildWeights(req api.SimRequest, opt Options) (*proto.StatWeightsRequest, error)`
  - `request.ErrNotWeights`, `request.ErrUnknownStat`
  - `func adapter.Weights(res *proto.StatWeightsResult, req api.SimRequest) ([]api.StatWeight, error)`
  - `adapter.ErrNoWeights`

- [ ] **Step 1: Write the failing test for the request half**

Create `sim/request/weights_test.go`:

```go
package request

import (
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

func weights() api.SimRequest {
	req := fury()
	req.Weights = &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit", "hit"},
		Reference: "attack_power",
	}
	return req
}

// The weights request is the SAME player, buffs, encounter and options
// as a DPS run - that is the whole point of reusing BuildWith - plus
// the stats to weigh.
func TestBuildWeightsIsTheSameRunPlusStats(t *testing.T) {
	got, err := BuildWeights(weights(), Options{})
	if err != nil {
		t.Fatal(err)
	}
	run, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if got.Player.GetName() != run.Raid.Parties[0].Players[0].GetName() {
		t.Error("the weights request carries a different player")
	}
	if got.Player.GetRotation() == nil {
		t.Error("the weights request carries no rotation")
	}
	if got.Player.GetDatabase() != nil {
		t.Error("BuildWeights attached a database; sim/internal/simdb.AttachWeights does that, once, at the caller")
	}
	if got.Encounter.GetDuration() != run.Encounter.GetDuration() {
		t.Error("the weights request fights a different encounter")
	}
	if got.SimOptions.GetIterations() != int32(fury().Iterations) {
		t.Errorf("iterations = %d", got.SimOptions.GetIterations())
	}
	if got.RaidBuffs == nil || got.PartyBuffs == nil || got.Debuffs == nil {
		t.Error("the weights request lost the buffs")
	}
	if got.Tanks == nil {
		t.Error("tanks is nil; the engine indexes it")
	}
	want := []proto.Stat{proto.Stat_StatAgility, proto.Stat_StatAttackPower, proto.Stat_StatCrit, proto.Stat_StatHit}
	if len(got.StatsToWeigh) != len(want) {
		t.Fatalf("stats_to_weigh = %v", got.StatsToWeigh)
	}
	for i, w := range want {
		if got.StatsToWeigh[i] != w {
			t.Errorf("stats_to_weigh[%d] = %v, want %v", i, got.StatsToWeigh[i], w)
		}
	}
	if got.EpReferenceStat != proto.Stat_StatAttackPower {
		t.Errorf("ep_reference_stat = %v", got.EpReferenceStat)
	}
}

func TestBuildWeightsRefusals(t *testing.T) {
	if _, err := BuildWeights(fury(), Options{}); !errors.Is(err, ErrNotWeights) {
		t.Error("BuildWeights accepted a plain run")
	}
	req := weights()
	req.Weights.Stats = []string{"haste", "attack_power"}
	req.Weights.Reference = "attack_power"
	_, err := BuildWeights(req, Options{})
	if !errors.Is(err, ErrUnknownStat) {
		t.Errorf("BuildWeights = %v, want ErrUnknownStat for a stat the engine does not carry", err)
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `go test ./sim/request/... -run BuildWeights -v`
Expected: FAIL to compile — `undefined: BuildWeights`.

- [ ] **Step 3: Write `sim/request/weights.go`**

```go
package request

// The stat weights request.
//
// The engine computes weights by running the same character twice per
// stat, once with a little more of it and once with a little less, so
// a weights request IS a raid sim request with a list of stats
// attached. Building it by taking BuildWith's output apart is what
// keeps the two from diverging: a buff, a consumable or a spec option
// that reaches a DPS run reaches the weights run by construction,
// rather than by somebody remembering to copy the line.

import (
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

var (
	// ErrNotWeights is returned when a request with no weights block is
	// handed to BuildWeights.
	ErrNotWeights = errors.New("request: the request carries no weights block")
	// ErrUnknownStat is returned for a stat id the engine has no Stat
	// value for. Weighing it as nothing would report a weight of zero
	// for a stat the player asked about.
	ErrUnknownStat = errors.New("request: unknown stat")
)

// BuildWeights turns a validated weights SimRequest into the engine's
// StatWeightsRequest.
func BuildWeights(req api.SimRequest, opt Options) (*proto.StatWeightsRequest, error) {
	if req.Weights == nil {
		return nil, ErrNotWeights
	}
	run, err := BuildWith(req, opt)
	if err != nil {
		return nil, err
	}
	stats := make([]proto.Stat, 0, len(req.Weights.Stats))
	for _, id := range req.Weights.Stats {
		s, ok := ParseStat(id)
		if !ok {
			return nil, fmt.Errorf("%w: %q; the ids are in sim/request/IDS.md", ErrUnknownStat, id)
		}
		stats = append(stats, s)
	}
	reference, ok := ParseStat(req.Weights.Reference)
	if !ok {
		return nil, fmt.Errorf("%w: reference %q", ErrUnknownStat, req.Weights.Reference)
	}
	party := run.Raid.Parties[0]
	return &proto.StatWeightsRequest{
		Player:     party.Players[0],
		RaidBuffs:  run.Raid.Buffs,
		PartyBuffs: party.Buffs,
		Debuffs:    run.Raid.Debuffs,
		Encounter:  run.Encounter,
		SimOptions: run.SimOptions,
		// The engine indexes this slice; a nil one is a panic rather
		// than a raid with no tank.
		Tanks:           []*proto.UnitReference{},
		StatsToWeigh:    stats,
		EpReferenceStat: reference,
	}, nil
}
```

- [ ] **Step 4: Attach the database for a weights request**

The item rows reach the engine on `Player.Database`, and `simdb.Attach` takes a `RaidSimRequest`. Add the sibling to `sim/internal/simdb/simdb.go`:

```go
// AttachWeights puts the database on a stat weights request's player,
// which is the same door as Attach's: proto.Player.Database, folded in
// by core.NewCharacter. A weights run equips the character the same
// way a DPS run does, so without it every item id resolves to nothing.
func AttachWeights(req *proto.StatWeightsRequest) error {
	db, err := load()
	if err != nil {
		return err
	}
	if req == nil || req.Player == nil {
		return nil
	}
	req.Player.Database = db
	return nil
}
```

- [ ] **Step 5: Write the failing test for the adapter half**

Append to `sim/adapter/adapter_test.go`:

```go
// The engine answers with a UnitStats array indexed by proto.Stat; the
// envelope answers with named rows in the order the request asked for
// them, normalised so the reference stat is exactly 1.
func TestWeightsMapping(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit"},
		Reference: "attack_power",
	}}
	stats := make([]float64, len(proto.Stat_name))
	stdev := make([]float64, len(proto.Stat_name))
	stats[proto.Stat_StatAgility] = 1.1
	stats[proto.Stat_StatAttackPower] = 1.0
	stats[proto.Stat_StatCrit] = 12.0
	stdev[proto.Stat_StatAgility] = 0.02
	stdev[proto.Stat_StatAttackPower] = 0.01
	stdev[proto.Stat_StatCrit] = 0.30

	res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
		Weights:      &proto.UnitStats{Stats: stats},
		WeightsStdev: &proto.UnitStats{Stats: stdev},
	}}
	got, err := Weights(res, req)
	if err != nil {
		t.Fatal(err)
	}
	want := []api.StatWeight{
		{Stat: "agility", Weight: 1.1, Error: 0.02},
		{Stat: "attack_power", Weight: 1.0, Error: 0.01},
		{Stat: "crit", Weight: 12.0, Error: 0.30},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d weights, want %d: %+v", len(got), len(want), got)
	}
	for i := range want {
		if got[i].Stat != want[i].Stat || got[i].Weight != want[i].Weight || got[i].Error != want[i].Error {
			t.Errorf("weight %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}

// A reference weight of zero cannot be normalised against, and
// reporting infinities would render as a table of blanks.
func TestWeightsRefusesAZeroReference(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{Stats: []string{"crit"}, Reference: "crit"}}
	stats := make([]float64, len(proto.Stat_name))
	res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
		Weights:      &proto.UnitStats{Stats: stats},
		WeightsStdev: &proto.UnitStats{Stats: stats},
	}}
	if _, err := Weights(res, req); err == nil {
		t.Error("a zero reference weight was normalised")
	}
}

func TestWeightsRefusals(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{Stats: []string{"crit"}, Reference: "crit"}}
	if _, err := Weights(nil, req); !errors.Is(err, ErrNoWeights) {
		t.Error("a nil result was accepted")
	}
	if _, err := Weights(&proto.StatWeightsResult{}, req); !errors.Is(err, ErrNoWeights) {
		t.Error("a result with no dps block was accepted")
	}
	failed := &proto.StatWeightsResult{Error: &proto.ErrorOutcome{Message: "the engine died"}}
	if _, err := Weights(failed, req); err == nil {
		t.Error("an engine failure was reported as weights")
	}
	if _, err := Weights(&proto.StatWeightsResult{}, api.SimRequest{}); err == nil {
		t.Error("a request with no weights block was accepted")
	}
}
```

- [ ] **Step 6: Write the adapter half**

Append to `sim/adapter/adapter.go`:

```go
// ErrNoWeights is returned when a stat weights result carries no DPS
// weight block. There is nothing to report and nothing to normalise.
var ErrNoWeights = errors.New("adapter: the result carries no stat weights")

// Weights lifts the engine's stat weights into the envelope's named
// rows.
//
// The engine answers with an array indexed by proto.Stat, which is a
// hundred-odd slots of mostly zero; the envelope answers with the
// stats the request asked for, in the order it asked, normalised so
// the reference stat is exactly 1. Normalising here rather than on the
// page is what makes the "copy for Pawn" string and the table the same
// numbers.
//
// The engine already normalises when asked, through EpReferenceStat,
// but it reports both the raw weights and the EP values and the two
// are easy to confuse; taking the raw weights and dividing is one
// arithmetic, in one place, that cannot pick the wrong block.
func Weights(res *proto.StatWeightsResult, req api.SimRequest) ([]api.StatWeight, error) {
	if req.Weights == nil {
		return nil, fmt.Errorf("%w: the request asked for none", ErrNoWeights)
	}
	if res == nil {
		return nil, fmt.Errorf("%w: nil result", ErrNoWeights)
	}
	if res.Error != nil && res.Error.Message != "" {
		return nil, fmt.Errorf("%w: %s", ErrSimFailed, res.Error.Message)
	}
	values := res.GetDps()
	if values.GetWeights() == nil {
		return nil, ErrNoWeights
	}
	raw := values.GetWeights().GetStats()
	stdev := values.GetWeightsStdev().GetStats()
	at := func(s proto.Stat, from []float64) float64 {
		if int(s) >= len(from) {
			return 0
		}
		return from[s]
	}

	reference, ok := request.ParseStat(req.Weights.Reference)
	if !ok {
		return nil, fmt.Errorf("%w: reference %q", ErrNoWeights, req.Weights.Reference)
	}
	scale := at(reference, raw)
	if scale == 0 {
		return nil, fmt.Errorf("%w: the reference stat %q weighs nothing, so nothing can be normalised against it", ErrNoWeights, req.Weights.Reference)
	}

	out := make([]api.StatWeight, 0, len(req.Weights.Stats))
	for _, id := range req.Weights.Stats {
		s, ok := request.ParseStat(id)
		if !ok {
			return nil, fmt.Errorf("%w: %q", ErrNoWeights, id)
		}
		out = append(out, api.StatWeight{
			Stat:   id,
			Weight: at(s, raw) / scale,
			Error:  at(s, stdev) / scale,
		})
	}
	return out, nil
}
```

Add `"github.com/jhunthrop/foreversixty/sim/request"` to the adapter's imports.

**If that import creates a cycle** — `sim/request` must not import `sim/adapter`, and today it does not — it is fine. Check with `go build ./sim/adapter/...` before going on; if a cycle appears, move `ParseStat`/`KnownStats` into a new leaf package `sim/internal/statid` imported by both, and say so in the commit message.

- [ ] **Step 7: Run the tests**

Run: `go test ./sim/request/... ./sim/adapter/... ./sim/internal/simdb/... -run 'Weights' -v`
Expected: PASS.

- [ ] **Step 8: Run the packages, vet, format, commit**

```bash
go test ./sim/request/... ./sim/adapter/... ./sim/internal/simdb/...
go vet ./sim/request/... ./sim/adapter/... ./sim/internal/simdb/...
gofmt -w sim/request sim/adapter sim/internal/simdb
git add sim/request sim/adapter sim/internal/simdb
git commit -m "feat(sim): the stat weights request and its normalised result rows

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 20: The sample iteration

**Files:**
- Modify: `sim/adapter/adapter.go`
- Modify: `sim/adapter/adapter_test.go`
- Modify: `sim/cmd/wasm/main.go` and `sim/cmd/forever-sim/main.go` (one line each: fill `Sample`)

**Interfaces:**
- Consumes: `proto.SampleIteration` (Task 1's amendment, Task 2's pin), `adapter.ActionName`.
- Produces: `func adapter.Sample(res *proto.RaidSimResult) []api.SampleCast`, and both lanes filling `SimResult.Sample`.

- [ ] **Step 1: Write the failing test**

Append to `sim/adapter/adapter_test.go`:

```go
// One iteration's casts, in order, with the pre-pull negative and each
// row keyed the same way the cast table's rows are - so the page can
// join a sample row to its cast summary without matching on a name.
func TestSampleMapsTheEnginesCastLog(t *testing.T) {
	res := &proto.RaidSimResult{SampleIteration: &proto.SampleIteration{Casts: []*proto.SampleCast{
		{AtSeconds: -1.5, ActionId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 1719}}},
		{AtSeconds: 0, ActionId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23881}},
			Target: "Target Dummy", Resources: map[string]int32{"rage": 26}},
		{AtSeconds: 1.502, ActionId: &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 13442}}},
		{AtSeconds: 2, ActionId: &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionAttack}}},
	}}}

	got := Sample(res)
	if len(got) != 4 {
		t.Fatalf("got %d casts, want 4", len(got))
	}
	if got[0].AtMS != -1500 {
		t.Errorf("the pre-pull cast is at %d ms, want -1500", got[0].AtMS)
	}
	if got[0].SpellID != 1719 || got[0].Name != "spell:1719" {
		t.Errorf("an untagged spell is %d %q, want its own id", got[0].SpellID, got[0].Name)
	}
	if got[1].Target != "Target Dummy" || got[1].Resources["rage"] != 26 {
		t.Errorf("cast 1 = %+v", got[1])
	}
	// Rounded, not truncated: 1.502 s is 1502 ms.
	if got[2].AtMS != 1502 {
		t.Errorf("cast 2 is at %d ms, want 1502", got[2].AtMS)
	}
	// An item and an "other" action get derived ids, exactly as the
	// cast table's rows do, so the two tables key alike.
	wantItemID, wantItemName := ActionName(&proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 13442}})
	if got[2].SpellID != wantItemID || got[2].Name != wantItemName {
		t.Errorf("the item cast is %d %q, want %d %q", got[2].SpellID, got[2].Name, wantItemID, wantItemName)
	}
	if got[3].Name != "other:attack" {
		t.Errorf("the white swing is %q", got[3].Name)
	}
}

// A result with no sample is not an error: an aborted run and an older
// engine both produce one, and the page renders the card only when
// there are rows.
func TestSampleOfNothingIsNothing(t *testing.T) {
	if got := Sample(nil); got != nil {
		t.Errorf("Sample(nil) = %+v", got)
	}
	if got := Sample(&proto.RaidSimResult{}); got != nil {
		t.Errorf("a result with no sample produced %+v", got)
	}
	if got := Sample(&proto.RaidSimResult{SampleIteration: &proto.SampleIteration{}}); got != nil {
		t.Errorf("an empty sample produced %+v", got)
	}
}

// The checked-in fixtures are real engine output, so the mapping is
// proved against one rather than only against hand-built messages.
func TestSampleOfTheWarriorFixture(t *testing.T) {
	res, err := Fixture("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	got := Sample(res)
	if len(got) == 0 {
		t.Skip("the checked-in fixture predates sample_iteration; refresh it with `forever-sim -out-proto`")
	}
	var last int64 = -1 << 62
	for i, c := range got {
		if c.AtMS < last {
			t.Errorf("cast %d is at %d ms, after %d; the sample is in cast order", i, c.AtMS, last)
		}
		last = c.AtMS
		if c.Name == "" {
			t.Errorf("cast %d has no name", i)
		}
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `go test ./sim/adapter/... -run Sample -v`
Expected: FAIL to compile — `undefined: Sample`.

- [ ] **Step 3: Write `Sample` in `sim/adapter/adapter.go`**

```go
// Sample is one iteration's casts, in order: the median-DPS iteration
// the engine records alongside the aggregate metrics.
//
// It is a table view of the cast log the timeline already renders, so
// every row takes its id and its name from ActionName - the same
// derivation the Casts table uses. Without that, an item's on-use and
// a spell of the same raw number would look like one action in one
// table and two in the other.
//
// No sample is not an error. An aborted run has none, and neither does
// a result from an engine older than the field; the page renders the
// card only when there are rows.
func Sample(res *proto.RaidSimResult) []api.SampleCast {
	casts := res.GetSampleIteration().GetCasts()
	if len(casts) == 0 {
		return nil
	}
	out := make([]api.SampleCast, 0, len(casts))
	for _, c := range casts {
		id, name := ActionName(c.GetActionId())
		row := api.SampleCast{
			// Rounded, not truncated: a cast at 1.5015 s is 1502 ms in
			// the timeline and must be 1502 here, or the two views of
			// one iteration disagree about when something happened.
			AtMS:    int64(math.Round(c.GetAtSeconds() * 1000)),
			SpellID: id,
			Name:    name,
			Target:  c.GetTarget(),
		}
		if len(c.GetResources()) > 0 {
			row.Resources = make(map[string]int, len(c.GetResources()))
			for k, v := range c.GetResources() {
				row.Resources[k] = int(v)
			}
		}
		out = append(out, row)
	}
	return out
}
```

- [ ] **Step 4: Fill it on both lanes**

In `sim/cmd/wasm/main.go`, in `simRun`'s final `return result(api.SimResult{...})`, add:

```go
		Sample:        adapter.Sample(engineRes),
```

In `sim/cmd/forever-sim/main.go`, in `Execute`, beside `base.Summary = sum`:

```go
	base.Sample = adapter.Sample(engineRes)
```

- [ ] **Step 5: Keep the sample out of a combined run**

`combine.Results` takes the first part's non-numeric fields, and a sample from one shard of a split run is one iteration of a quarter of the run — which is fine and is what the contract asks for. But a *stage* request's sample would be carried into a bulk result and shown beside a ranking it does not belong to. Add to `finalResult` in `sim/bulk/rank.go`, beside `out.IterationsRun`:

```go
	// The sample belongs to the equipped set, which is results[0] -
	// base here - so it is already the right one. Stated rather than
	// assumed, because the obvious alternative (the winner's) would be
	// a cast log of a character the player does not have.
	out.Sample = base.Sample
```

- [ ] **Step 6: Run the tests**

Run: `go test ./sim/adapter/... ./sim/bulk/... ./sim/cmd/forever-sim/... -v`
Expected: PASS. `TestSampleOfTheWarriorFixture` may skip until the fixtures are refreshed; that is expected and the skip message says how.

- [ ] **Step 7: Vet, format, commit**

```bash
go vet ./sim/adapter/... ./sim/bulk/... ./sim/cmd/...
gofmt -w sim/adapter sim/bulk sim/cmd
git add sim/adapter sim/bulk sim/cmd
git commit -m "feat(sim): the sample iteration's cast log on the result

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 21: `simPlan` and `simRank`

**Files:**
- Modify: `sim/cmd/wasm/main.go`
- Modify: `sim/cmd/wasm/main_test.go`

**Interfaces:**
- Consumes: `bulk.Plan`, `bulk.Rank`, `bulk.StageRequests` (Tasks 16–18).
- Produces:
  - `simPlan(requestJSON)` → a `StageRequests` JSON, or `{"error": "..."}`
  - `simRank(requestJSON, stageJSON, resultsJSON)` → `{"next": <stage>}` or `{"result": <SimResult>}`, or `{"error": "..."}`

- [ ] **Step 1: Write the failing test**

`sim/cmd/wasm/main.go` is `//go:build js && wasm`, so its package tests cannot run on the host. `main_test.go` already exists for the host-testable parts; keep the new logic out of the `js.Func` wrappers so it is testable the same way. Append to `sim/cmd/wasm/main_test.go`:

```go
// The two bulk exports are thin: the wrappers unwrap js.Value and the
// functions below do the work, so the work is testable on the host
// and the wrappers carry nothing that could be wrong.
func TestPlanJSONRoundTrip(t *testing.T) {
	req := bulkFixtureRequest(t)
	out := planJSON(req)
	var stage bulk.StageRequests
	if err := json.Unmarshal([]byte(out), &stage); err != nil {
		t.Fatalf("simPlan returned %s", out)
	}
	if stage.Stage != 1 || len(stage.Requests) != len(stage.Combos)+1 {
		t.Errorf("stage = %+v", stage)
	}
}

func TestPlanJSONReportsAnError(t *testing.T) {
	out := planJSON("not json")
	var e struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(out), &e); err != nil || e.Error == "" {
		t.Errorf("simPlan of rubbish returned %s", out)
	}
}

func TestRankJSONReturnsTheNextStageThenTheResult(t *testing.T) {
	reqJSON := bulkFixtureRequest(t)
	stageJSON := planJSON(reqJSON)

	var stage bulk.StageRequests
	if err := json.Unmarshal([]byte(stageJSON), &stage); err != nil {
		t.Fatal(err)
	}
	resultsJSON := fakeResultsJSON(t, stage)

	out := rankJSON(reqJSON, stageJSON, resultsJSON)
	var first struct {
		Next   *bulk.StageRequests `json:"next"`
		Result *api.SimResult      `json:"result"`
	}
	if err := json.Unmarshal([]byte(out), &first); err != nil {
		t.Fatalf("simRank returned %s", out)
	}
	if first.Next == nil || first.Result != nil {
		t.Fatalf("the first rung returned %s", out)
	}

	nextJSON, err := json.Marshal(first.Next)
	if err != nil {
		t.Fatal(err)
	}
	out = rankJSON(reqJSON, string(nextJSON), fakeResultsJSON(t, *first.Next))
	var last struct {
		Next   *bulk.StageRequests `json:"next"`
		Result *api.SimResult      `json:"result"`
	}
	if err := json.Unmarshal([]byte(out), &last); err != nil {
		t.Fatalf("simRank returned %s", out)
	}
	if last.Next != nil || last.Result == nil {
		t.Fatalf("the last rung returned %s", out)
	}
	if len(last.Result.Combos) == 0 || last.Result.Equipped == nil {
		t.Errorf("the finished result is %+v", last.Result)
	}
}

func TestRankJSONReportsAMismatch(t *testing.T) {
	reqJSON := bulkFixtureRequest(t)
	out := rankJSON(reqJSON, planJSON(reqJSON), `[]`)
	var e struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(out), &e); err != nil || e.Error == "" {
		t.Errorf("simRank of no results returned %s", out)
	}
}
```

with two helpers at the bottom of the file:

```go
// bulkFixtureRequest is a two-candidate Top Gear over the checked-in
// warrior fixture, as JSON.
func bulkFixtureRequest(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "adapter", "testdata", "warrior-fury.request.json"))
	if err != nil {
		t.Fatal(err)
	}
	var req api.SimRequest
	if err := json.Unmarshal(b, &req); err != nil {
		t.Fatal(err)
	}
	req.EngineVersion = enginever.Version
	req.Iterations = 3000
	head := req.Character.Gear[0]
	req.Bulk = &api.BulkSpec{
		Mode:      api.KindGear,
		Precision: api.PrecisionNormal,
		Cap:       api.Caps[api.LaneBrowser],
		Candidates: []api.Candidate{
			{Slot: "head", ItemID: head.ItemID, Origin: api.OriginBag},
			{Slot: "neck", ItemID: req.Character.Gear[1].ItemID, Origin: api.OriginBag},
		},
	}
	out, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

// fakeResultsJSON answers a stage with plausible numbers, so the wasm
// boundary can be tested without running the engine on the host.
func fakeResultsJSON(t *testing.T, stage bulk.StageRequests) string {
	t.Helper()
	out := make([]api.SimResult, len(stage.Requests))
	for i := range stage.Requests {
		out[i] = api.SimResult{
			EngineVersion: enginever.Version,
			Request:       stage.Requests[i],
			Lane:          api.LaneBrowser,
			IterationsRun: stage.Iterations,
			DPS:           api.Estimate{Mean: 1000 + float64(i)*10, StdDev: 200, Error: 4},
			Summary:       adapter.EmptySummary(),
		}
	}
	b, err := json.Marshal(out)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
```

**Note the build tags.** `main.go` is `//go:build js && wasm` and `main_test.go` is `//go:build !js`, so on the host the test file is the only file in the package and today it says there is nothing to test. Put the JSON functions (`planJSON`, `rankJSON`, and Task 22's `weightsJSON`) in a new `sim/cmd/wasm/exports.go` with **no build tag**, so it compiles under both, and put the tests above in `main_test.go` beside the existing note (or in a new `exports_test.go`, also `//go:build !js`). That is why these functions take and return strings rather than `js.Value`: the boundary becomes host-testable, and the `js.Func` wrappers in `main.go` carry nothing that could be wrong. Replace the existing "there is nothing here to unit test" comment in `main_test.go` with a line saying the JSON bodies are tested here and only the `js.Value` unwrapping is not.

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./sim/cmd/wasm/... -v`
Expected: FAIL to compile — `undefined: planJSON`.

- [ ] **Step 3: Write `sim/cmd/wasm/exports.go`**

```go
package main

// The JSON half of the exports.
//
// Everything the browser calls is a JSON string in and a JSON string
// out, and the js.Func wrappers in main.go do nothing but unwrap the
// arguments. The work lives here, with NO build tag, so it is compiled
// and tested on the host: a bug in the planner's boundary would
// otherwise only ever be found in a browser.

import (
	"encoding/json"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
)

// planJSON is simPlan's body: a bulk SimRequest in, the first stage
// out.
//
// The page then runs each of stage.requests through the sharding it
// already has - simSplit, simRun per worker, simCombine - and hands
// the results back to simRank in the same order. No statistics happen
// in TypeScript.
func planJSON(requestJSON string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	stage, err := bulk.Plan(req)
	if err != nil {
		return errorJSON(err.Error())
	}
	return encodeOrError(stage)
}

// rankJSON is simRank's body: the request, the stage it answers, and
// one result per request in that stage's order.
//
// It returns {"next": stage} or {"result": SimResult}, never both,
// which is how the page knows whether to loop.
func rankJSON(requestJSON, stageJSON, resultsJSON string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return errorJSON("the request is not valid JSON: " + err.Error())
	}
	var stage bulk.StageRequests
	if err := strictDecode(stageJSON, &stage); err != nil {
		return errorJSON("the stage is not valid JSON: " + err.Error())
	}
	var results []api.SimResult
	if err := strictDecode(resultsJSON, &results); err != nil {
		return errorJSON("the results are not valid JSON: " + err.Error())
	}
	next, final, err := bulk.Rank(req, stage, results)
	if err != nil {
		return errorJSON(err.Error())
	}
	out := struct {
		Next   *bulk.StageRequests `json:"next,omitempty"`
		Result *api.SimResult      `json:"result,omitempty"`
	}{Next: next, Result: final}
	return encodeOrError(out)
}

// strictDecode refuses a field the type does not carry, the way
// decodeRequest does: a page sending something this wasm cannot honour
// is a version mismatch, and running anyway would drop part of it.
func strictDecode(s string, into any) error {
	dec := json.NewDecoder(strings.NewReader(s))
	dec.DisallowUnknownFields()
	return dec.Decode(into)
}

// encodeOrError marshals, or returns the {"error": ...} shape.
func encodeOrError(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return errorJSON(err.Error())
	}
	return string(b)
}
```

Move `decodeRequest` and `errorJSON` out of `main.go` into this file unchanged, so both build tags see them.

- [ ] **Step 4: Add the wrappers to `sim/cmd/wasm/main.go`**

Register them beside the others in `main()`:

```go
	js.Global().Set("simPlan", js.FuncOf(simPlan))
	js.Global().Set("simRank", js.FuncOf(simRank))
```

and add:

```go
// simPlan(requestJSON) returns the first stage of a bulk run:
// {"stage":1,"iterations":100,"requests":[...],"combos":[...]}.
func simPlan(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorJSON("simPlan takes (requestJSON)")
	}
	return planJSON(args[0].String())
}

// simRank(requestJSON, stageJSON, resultsJSON) scores a finished stage
// and returns {"next": stage} or {"result": SimResult}.
func simRank(_ js.Value, args []js.Value) any {
	if len(args) < 3 {
		return errorJSON("simRank takes (requestJSON, stageJSON, resultsJSON)")
	}
	return rankJSON(args[0].String(), args[1].String(), args[2].String())
}
```

Update the package comment's "exactly four functions" to name all seven, and the sentence about the handshake to say the host page is told when all seven exist.

- [ ] **Step 5: Run the tests and build the wasm**

```bash
go test ./sim/cmd/wasm/... -v
(cd sim && GOOS=js GOARCH=wasm go build -o /dev/null ./cmd/wasm)
```

Expected: PASS, and the wasm build succeeds.

- [ ] **Step 6: Vet, format, commit**

```bash
go vet ./sim/cmd/wasm/...
gofmt -w sim/cmd/wasm
git add sim/cmd/wasm
git commit -m "feat(sim): simPlan and simRank, with the JSON boundary testable on the host

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 22: `simWeights`

**Files:**
- Modify: `sim/cmd/wasm/exports.go`
- Modify: `sim/cmd/wasm/main.go`
- Modify: `sim/cmd/wasm/main_test.go`

**Interfaces:**
- Consumes: `request.BuildWeights` (Task 19), `simdb.AttachWeights` (Task 19), `adapter.Weights` (Task 19), `core.StatWeightsAsync`.
- Produces: `simWeights(requestJSON, callbackId)` → a `SimResult` with `Weights`, progress through the existing `simProgress` global.

- [ ] **Step 1: Write the failing test**

Append to `sim/cmd/wasm/main_test.go`:

```go
// The weights export's host-testable half is the shape of its
// refusals: running the engine needs a browser, but a malformed
// request must come back as a result, not a throw, on every lane.
func TestWeightsJSONRefusals(t *testing.T) {
	for _, c := range []struct {
		name string
		body string
		want string
	}{
		{"not json", "{", "valid JSON"},
		{"a plain run", plainFixtureRequest(t), "weights"},
	} {
		t.Run(c.name, func(t *testing.T) {
			out := weightsJSON(c.body, "cb")
			var res api.SimResult
			if err := json.Unmarshal([]byte(out), &res); err != nil {
				t.Fatalf("weightsJSON returned %s", out)
			}
			if res.Error == "" {
				t.Fatalf("a bad request returned no error: %s", out)
			}
			if !strings.Contains(res.Error, c.want) {
				t.Errorf("error %q does not mention %q", res.Error, c.want)
			}
			// Every export answers in the same shape, so the page
			// never has to distinguish a throw from a result.
			if res.Summary.DamageDone == nil {
				t.Error("a failed weights run carries a null summary; it must carry the empty one")
			}
		})
	}
}
```

with the fixture helper:

```go
// plainFixtureRequest is the checked-in warrior request, unchanged, as
// JSON: a run with no weights block.
func plainFixtureRequest(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "adapter", "testdata", "warrior-fury.request.json"))
	if err != nil {
		t.Fatal(err)
	}
	var req api.SimRequest
	if err := json.Unmarshal(b, &req); err != nil {
		t.Fatal(err)
	}
	req.EngineVersion = enginever.Version
	out, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `go test ./sim/cmd/wasm/... -run WeightsJSON -v`
Expected: FAIL to compile — `undefined: weightsJSON`.

- [ ] **Step 3: Split the export in two**

The decoding and the refusals are host-testable; the engine call is not. In `sim/cmd/wasm/exports.go` (no build tag):

```go
// weightsRunner is the engine half of simWeights, supplied by the
// wasm build. The host build has none, so weightsJSON's refusals are
// testable without syscall/js and the engine call is not duplicated.
var weightsRunner func(req api.SimRequest, callbackID string) (api.SimResult, error)

// weightsJSON is simWeights' body: a weights SimRequest in, a
// SimResult with Weights out. Progress is reported through the same
// simProgress global a plain run uses.
func weightsJSON(requestJSON, callbackID string) string {
	req, err := decodeRequest(requestJSON)
	if err != nil {
		return failJSON(req, "the request is not valid JSON: "+err.Error())
	}
	if req.Weights == nil {
		return failJSON(req, "simWeights takes a request with a weights block; this one has none")
	}
	if err := req.ValidateLane(api.LaneBrowser); err != nil {
		return failJSON(req, err.Error())
	}
	if weightsRunner == nil {
		return failJSON(req, "this build has no engine; simWeights runs only in the browser")
	}
	res, err := weightsRunner(req, callbackID)
	if err != nil {
		return failJSON(req, err.Error())
	}
	return encodeOrError(stamp(res))
}
```

Move `fail`, `stopped`, `result` out of `main.go` into `exports.go` as well, renamed so the untagged file owns the shape:

```go
// stamp fills the two fields every export sets the same way.
// EngineVersion is enginever.Version, never the request's claim: a
// row's provenance is a fact about the binary that produced it.
func stamp(res api.SimResult) api.SimResult {
	res.EngineVersion = enginever.Version
	res.Lane = api.LaneBrowser
	return res
}

// failJSON wraps an error as a SimResult, so every export returns the
// same shape and the worker never has to distinguish a throw from a
// result. The empty summary rather than a zero one: a nil Go slice
// marshals as null, and the page would need a null check per key on
// the path least likely to be exercised.
func failJSON(req api.SimRequest, msg string) string {
	return encodeOrError(stamp(api.SimResult{Request: req, Error: msg, Summary: adapter.EmptySummary()}))
}
```

and have `main.go`'s existing `fail`, `stopped` and `result` call them, or delete them and call these directly — whichever leaves fewer names. Keep the comments.

- [ ] **Step 4: Write the engine half in `sim/cmd/wasm/main.go`**

```go
func init() {
	weightsRunner = runWeights
}

// runWeights is the engine half of simWeights.
//
// The engine runs the same character twice per stat, a little above
// and a little below, so this is a raid sim request with a stat list
// - which is exactly what request.BuildWeights builds. Progress is
// the engine's own per-sim ticks, reported through the simProgress
// global a plain run already uses, so the page's progress handling is
// unchanged.
func runWeights(req api.SimRequest, callbackID string) (api.SimResult, error) {
	start := time.Now()
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return api.SimResult{}, err
	}
	// Forever's own item rows: a weights run equips the character the
	// same way a DPS run does.
	if err := simdb.AttachWeights(engineReq); err != nil {
		return api.SimResult{}, err
	}

	reporter := make(chan *proto.ProgressMetrics, 32)
	core.StatWeightsAsync(engineReq, reporter, callbackID)

	var engineRes *proto.StatWeightsResult
	for p := range reporter {
		if p.FinalWeightResult != nil {
			engineRes = p.FinalWeightResult
			break
		}
		if cb := js.Global().Get("simProgress"); cb.Type() == js.TypeFunction {
			if b, err := json.Marshal(api.Progress{
				IterationsRun: int(p.CompletedIterations),
				DPS:           api.Estimate{Mean: p.Dps},
				// A weights run is many sims, so the page shows the
				// same "n of m" line a bulk stage does rather than a
				// bare iteration count that restarts per stat.
				CombosDone:  int(p.CompletedSims),
				CombosTotal: int(p.TotalSims),
			}); err == nil {
				cb.Invoke(callbackID, string(b))
			}
		}
	}
	if engineRes == nil {
		return api.SimResult{}, errors.New("the engine produced no weights")
	}
	weights, err := adapter.Weights(engineRes, req)
	if err != nil {
		return api.SimResult{}, err
	}
	return api.SimResult{
		Request:       req,
		IterationsRun: req.Iterations,
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       adapter.EmptySummary(),
		Weights:       weights,
	}, nil
}
```

and register the export in `main()`:

```go
	js.Global().Set("simWeights", js.FuncOf(simWeights))
```

with:

```go
// simWeights(requestJSON, callbackId) computes stat weights and
// returns a SimResult with Weights filled. Progress is reported the
// way simRun reports it.
func simWeights(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return errorJSON("simWeights takes (requestJSON, callbackId)")
	}
	return weightsJSON(args[0].String(), args[1].String())
}
```

- [ ] **Step 5: Run the tests and build the wasm**

```bash
go test ./sim/cmd/wasm/... -v
(cd sim && GOOS=js GOARCH=wasm go build -o /dev/null ./cmd/wasm)
```

Expected: PASS and a successful wasm build.

- [ ] **Step 6: Check the size budget still holds**

```bash
make artifacts
```

Expected: the final line reports the gzipped size under 4 MB. `sim/bulk` and the weights path add Go, not engine, so the increase should be small; if the budget is breached, **stop and report it** rather than raising the number.

- [ ] **Step 7: Vet, format, commit**

```bash
go vet ./sim/cmd/wasm/...
gofmt -w sim/cmd/wasm
git add sim/cmd/wasm
git commit -m "feat(sim): simWeights, over the engine's own StatWeightsAsync

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 23: Widen the runner's progress callback

The contract's section 8 says progress rows carry `stage`, `combos_done` and `combos_total`. `sim/runner` is how the api module reads a native run's progress, and its callback carries two bare numbers. Widening it to `api.Progress` is a **breaking change across the module boundary**: the api module's callers stop compiling until they follow. That is intended and the api lane's plan has the matching change; this task's verification is scoped to `sim/`.

**Files:**
- Modify: `sim/runner/runner.go`
- Modify: `sim/runner/native.go`
- Modify: `sim/runner/fixture.go`
- Modify: `sim/runner/runner_test.go`

**Interfaces:**
- Consumes: `api.Progress` (Task 7).
- Produces: `type runner.Progress func(api.Progress)`; `Native` parsing `stage`, `combos_done`, `combos_total` off the tick line.

- [ ] **Step 1: Write the failing test**

In `sim/runner/runner_test.go`, change every `func(done int, mean float64)` callback to `func(p api.Progress)` and add:

```go
// A bulk run's progress lines carry the stage and the combination
// counts, and the api lane shows them as "stage 2 of 3 - 31 of 96
// combinations". A runner that dropped them would leave the page
// watching an iteration count that restarts every combination.
func TestNativeReportsStageProgress(t *testing.T) {
	stderr := strings.Join([]string{
		`{"completed":100,"total":400,"dps":900,"stage":1,"combos_done":1,"combos_total":4}`,
		`{"completed":200,"total":400,"dps":950,"stage":1,"combos_done":2,"combos_total":4}`,
		`{"log":"Running 400 iterations"}`,
	}, "\n")
	var seen []api.Progress
	n := &Native{Binary: fakeBinary(t, resultJSON(t), stderr, 0)}
	if _, err := n.Run(context.Background(), validRequest(), func(p api.Progress) {
		seen = append(seen, p)
	}); err != nil {
		t.Fatal(err)
	}
	if len(seen) != 2 {
		t.Fatalf("saw %d progress ticks, want 2: %+v", len(seen), seen)
	}
	if seen[1].IterationsRun != 200 || seen[1].DPS.Mean != 950 {
		t.Errorf("tick 1 = %+v", seen[1])
	}
	if seen[1].Stage != 1 || seen[1].CombosDone != 2 || seen[1].CombosTotal != 4 {
		t.Errorf("tick 1 lost the stage fields: %+v", seen[1])
	}
}

// A plain run's ticks carry no stage, and the fields stay zero rather
// than being invented.
func TestNativeReportsPlainProgress(t *testing.T) {
	stderr := `{"completed":750,"total":3000,"dps":1010}`
	var seen []api.Progress
	n := &Native{Binary: fakeBinary(t, resultJSON(t), stderr, 0)}
	if _, err := n.Run(context.Background(), validRequest(), func(p api.Progress) {
		seen = append(seen, p)
	}); err != nil {
		t.Fatal(err)
	}
	if len(seen) != 1 || seen[0].Stage != 0 || seen[0].CombosTotal != 0 {
		t.Errorf("a plain run's tick = %+v", seen)
	}
}
```

Use whatever helpers `runner_test.go` already has in place of `fakeBinary`, `resultJSON` and `validRequest`; the file builds a stub binary today and those are its names or their equivalents.

- [ ] **Step 2: Run it to verify it fails**

Run: `go test ./sim/runner/... -v`
Expected: FAIL to compile — the callback signature no longer matches.

- [ ] **Step 3: Widen the type in `sim/runner/runner.go`**

```go
// Progress is called as a run advances, so a caller can refine what it
// is showing. It is never called after Run returns.
//
// It takes the whole api.Progress rather than a count and a mean
// because a bulk run's progress is "stage 2 of 3, 31 of 96
// combinations" and an iteration count alone would restart on every
// combination. The same struct is what the browser's simProgress
// callback carries, so the two lanes report progress in one shape.
type Progress func(p api.Progress)
```

- [ ] **Step 4: Parse the new fields in `sim/runner/native.go`**

Replace the `tick` struct and its use:

```go
// tick is one progress line on stderr, as forever-sim writes it with
// -progress. The three bulk fields are absent for a plain run.
type tick struct {
	Completed   int     `json:"completed"`
	Total       int     `json:"total"`
	DPS         float64 `json:"dps"`
	Stage       int     `json:"stage"`
	CombosDone  int     `json:"combos_done"`
	CombosTotal int     `json:"combos_total"`
}
```

and where the goroutine calls `onProgress`:

```go
			if onProgress != nil {
				onProgress(api.Progress{
					IterationsRun: t.Completed,
					DPS:           api.Estimate{Mean: t.DPS},
					Stage:         t.Stage,
					CombosDone:    t.CombosDone,
					CombosTotal:   t.CombosTotal,
				})
			}
```

- [ ] **Step 5: Follow in `sim/runner/fixture.go`**

Change its `onProgress` calls to build an `api.Progress`. If the fixture runner reports a single mid-run tick, keep it at the same completion point and leave the three bulk fields zero.

- [ ] **Step 6: Run the tests**

Run: `go test ./sim/runner/... -v`
Expected: PASS.

- [ ] **Step 7: Check what this breaks outside the module**

Run: `go build ./...` from the worktree root.
Expected: the `api` module fails to compile at its `runner.Progress` call sites. **That is the intended, contract-mandated break.** Record the exact failing file and line in the commit message so the api lane can land its side; do not edit the api module from this lane.

- [ ] **Step 8: Vet, format, commit**

```bash
go vet ./sim/runner/...
gofmt -w sim/runner
git add sim/runner
git commit -m "feat(sim)!: the runner's progress callback carries the whole api.Progress

A bulk run's progress is a stage and a combination count, not an
iteration count, so the callback carries api.Progress rather than two
numbers - the same struct the browser's simProgress callback carries.

The api module's call sites need the matching change; they are listed
in the parity contract's section 8.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 24: `forever-sim` runs the plan–rank loop

**Files:**
- Create: `sim/cmd/forever-sim/bulk.go`
- Create: `sim/cmd/forever-sim/bulk_test.go`
- Modify: `sim/cmd/forever-sim/main.go`

**Interfaces:**
- Consumes: `bulk.PlanWith`, `bulk.Rank`, `bulk.StageRequests`, `bulk.Options`, `bulk.LoadEnchants` (Tasks 13–18); `Execute` (existing).
- Produces:
  - `func executeBulk(req api.SimRequest, opt bulk.Options, progress io.Writer) (api.SimResult, error)`
  - a `-enchants <path>` flag
  - progress lines carrying `stage`, `combos_done`, `combos_total`

- [ ] **Step 1: Write the failing test**

Create `sim/cmd/forever-sim/bulk_test.go`:

```go
package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// bulkRequest is the checked-in warrior fixture with two candidates,
// at the fewest iterations that still answers: the binary's job here
// is the LOOP, and the engine's numbers are tested in sim/adapter.
func bulkRequest(t *testing.T) api.SimRequest {
	t.Helper()
	req := fixtureRequest(t, "warrior-fury")
	req.EngineVersion = enginever.Version
	req.Iterations = 3000
	req.Encounter.DurationSec = 60
	req.Bulk = &api.BulkSpec{
		Mode:      api.KindGear,
		Precision: api.PrecisionNormal,
		Cap:       api.Caps[api.LaneServer],
		Candidates: []api.Candidate{
			{Slot: "head", ItemID: req.Character.Gear[0].ItemID, Origin: api.OriginBag},
			{Slot: "neck", ItemID: req.Character.Gear[1].ItemID, Origin: api.OriginBag},
		},
	}
	return req
}

// The whole loop, end to end, at the smallest useful size. It is slow
// (two stages of three sims), so it is one test rather than a table.
func TestExecuteBulkRunsEveryStage(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	var progress bytes.Buffer
	req := bulkRequest(t)
	res, err := executeBulk(req, bulk.Options{}, &progress)
	if err != nil {
		t.Fatal(err)
	}
	if res.Equipped == nil || res.Equipped.Mean <= 0 {
		t.Fatalf("equipped = %+v", res.Equipped)
	}
	if len(res.Combos) != 3 {
		t.Errorf("got %d combos, want the head, the neck and both", len(res.Combos))
	}
	if len(res.Stages) != 2 || res.Stages[0].Iterations != 1000 || res.Stages[1].Iterations != 3000 {
		t.Errorf("stages = %+v", res.Stages)
	}
	if res.Request.Bulk == nil {
		t.Error("the result lost the request that produced it")
	}
	if res.Lane != api.LaneServer {
		t.Errorf("lane = %q", res.Lane)
	}
	if res.EngineVersion != enginever.Version {
		t.Errorf("engine_version = %q", res.EngineVersion)
	}
	// A within-error group is assigned for every row, and the leader
	// is group 0.
	if res.Combos[0].Group != 0 {
		t.Errorf("the leader is group %d", res.Combos[0].Group)
	}

	// Every stderr line is JSON, and the bulk ones carry the stage.
	var stages, combos int
	for _, line := range strings.Split(strings.TrimSpace(progress.String()), "\n") {
		if line == "" {
			continue
		}
		var tick struct {
			Stage       int `json:"stage"`
			CombosDone  int `json:"combos_done"`
			CombosTotal int `json:"combos_total"`
			Log         string
		}
		if err := json.Unmarshal([]byte(line), &tick); err != nil {
			t.Fatalf("a progress line is not JSON: %s", line)
		}
		if tick.Stage > stages {
			stages = tick.Stage
		}
		if tick.CombosTotal > combos {
			combos = tick.CombosTotal
		}
	}
	if stages != 2 {
		t.Errorf("progress reported %d stages, want 2", stages)
	}
	if combos == 0 {
		t.Error("progress never reported a combination count")
	}
}

func TestExecuteBulkRefusesAPlainRun(t *testing.T) {
	if _, err := executeBulk(fixtureRequest(t, "warrior-fury"), bulk.Options{}, nil); err == nil {
		t.Error("executeBulk accepted a request with no bulk block")
	}
}
```

Use the existing helper in `main_test.go` that loads a fixture request in place of `fixtureRequest`; if there is none, add one that reads `../../adapter/testdata/<spec>.request.json` and decodes it strictly.

- [ ] **Step 2: Run it to verify it fails**

Run: `go test ./sim/cmd/forever-sim/... -run ExecuteBulk -v`
Expected: FAIL to compile — `undefined: executeBulk`.

- [ ] **Step 3: Write `sim/cmd/forever-sim/bulk.go`**

```go
package main

// The native plan-rank loop.
//
// This is the same three functions the browser calls - bulk.Plan, a
// run per request, bulk.Rank - in a loop instead of across a worker
// pool. Nothing about WHICH sims run or HOW they are ranked lives
// here; that is sim/bulk's, compiled into both artifacts, which is
// what stops the premium lane and the free lane disagreeing about who
// won.
//
// Concurrency is inside each run: core.RunRaidSimConcurrentAsync
// already splits one sim across every CPU. Running the stage's sims
// in parallel on top of that would oversubscribe the 4-CPU job and
// make the progress line meaningless.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/bulk"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// executeBulk runs every stage of a bulk request and returns the
// ranked result.
func executeBulk(req api.SimRequest, opt bulk.Options, progress io.Writer) (api.SimResult, error) {
	start := time.Now()
	stage, err := bulk.PlanWith(req, opt)
	if err != nil {
		return api.SimResult{}, fmt.Errorf("%w: %v", errBadInput, err)
	}
	for {
		results := make([]api.SimResult, 0, len(stage.Requests))
		for i, r := range stage.Requests {
			// The stage's own progress, written before each sim rather
			// than after, so a long stage says what it is doing rather
			// than going quiet and then jumping.
			writeStageTick(progress, stage, i)
			// The per-sim ticks are suppressed: one line per iteration
			// batch times ninety-six combinations is a stderr stream
			// nobody can read, and the api lane's reader wants the
			// stage counts.
			res, err := Execute(r, nil)
			if err != nil {
				if errors.Is(err, adapter.ErrAborted) {
					// A stopped bulk run returns what it has: the
					// stage it reached, marked partial, the way a
					// stopped single run does.
					return abortedBulk(req, stage, start), err
				}
				return api.SimResult{}, err
			}
			results = append(results, res)
		}
		writeStageTick(progress, stage, len(stage.Requests))

		next, final, err := bulk.Rank(req, stage, results)
		if err != nil {
			return api.SimResult{}, err
		}
		if final != nil {
			final.EngineVersion = enginever.Version
			final.Lane = api.LaneServer
			final.DurationMS = time.Since(start).Milliseconds()
			return *final, nil
		}
		stage = *next
	}
}

// writeStageTick writes one bulk progress line. done counts the
// requests of this stage that have finished, and the equipped set is
// one of them - the page says "31 of 96" about the whole stage, not
// about the combinations alone, because that is what the bar fills to.
func writeStageTick(progress io.Writer, stage bulk.StageRequests, done int) {
	if progress == nil {
		return
	}
	_ = json.NewEncoder(progress).Encode(struct {
		Completed   int     `json:"completed"`
		Total       int     `json:"total"`
		DPS         float64 `json:"dps"`
		Stage       int     `json:"stage"`
		CombosDone  int     `json:"combos_done"`
		CombosTotal int     `json:"combos_total"`
	}{
		Completed:   done * stage.Iterations,
		Total:       len(stage.Requests) * stage.Iterations,
		Stage:       stage.Stage,
		CombosDone:  done,
		CombosTotal: len(stage.Requests),
	})
}

// abortedBulk is what a stopped bulk run writes: the request, the
// stage it reached, and nothing pretending to be a ranking.
func abortedBulk(req api.SimRequest, stage bulk.StageRequests, start time.Time) api.SimResult {
	return api.SimResult{
		EngineVersion: enginever.Version,
		Request:       req,
		Lane:          api.LaneServer,
		Aborted:       true,
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       adapter.EmptySummary(),
		Stages:        []api.Stage{{Iterations: stage.Iterations, Combos: len(stage.Combos)}},
	}
}
```

- [ ] **Step 4: Dispatch on the kind in `main.go`**

Add the flag beside the others in `main`:

```go
	enchants := flag.String("enchants", "", "data/builds/<build>/enchants.json; needed only when a bulk request names an enchant rather than inheriting one")
```

and pass it down. Change `run` to:

```go
// run is main's body, with its files and its progress sink as
// parameters so it is testable.
func run(inPath, outPath, enchantPath string, iterations int, progress io.Writer) error {
	req, err := load(inPath, iterations)
	if err != nil {
		return err
	}
	opt, err := bulkOptions(enchantPath)
	if err != nil {
		return err
	}
	res, err := dispatch(req, opt, progress)
	if err != nil && !errors.Is(err, adapter.ErrAborted) {
		return err
	}
	b, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		return fmt.Errorf("marshalling the result: %w", marshalErr)
	}
	if writeErr := write(outPath, b); writeErr != nil {
		return writeErr
	}
	return err
}

// dispatch picks the pipeline the request's kind asks for. The kind is
// derived from the request, never passed as a flag: an operator who
// could say "run this as a bulk" could say it of a request with no
// bulk block.
func dispatch(req api.SimRequest, opt bulk.Options, progress io.Writer) (api.SimResult, error) {
	switch req.Kind() {
	case api.KindGear, api.KindTalents, api.KindDrops:
		return executeBulk(req, opt, progress)
	case api.KindWeights:
		return executeWeights(req, progress)
	default:
		return ExecuteToTarget(req, progress)
	}
}

// bulkOptions loads the build tables a bulk request may need. An
// absent path is not an error: a request whose candidates all inherit
// their enchants needs no table, and requiring one would make every
// Droptimizer job depend on a file it never reads.
func bulkOptions(enchantPath string) (bulk.Options, error) {
	if enchantPath == "" {
		return bulk.Options{}, nil
	}
	f, err := os.Open(enchantPath)
	if err != nil {
		return bulk.Options{}, fmt.Errorf("%w: reading the enchant table: %v", errBadInput, err)
	}
	defer f.Close()
	table, err := bulk.LoadEnchants(f)
	if err != nil {
		return bulk.Options{}, fmt.Errorf("%w: %v", errBadInput, err)
	}
	return bulk.Options{Enchants: table}, nil
}
```

`ExecuteToTarget` and `executeWeights` are Task 25. Until then, make `dispatch`'s default arm `Execute(req, progress)` and its weights arm return an error saying so, and replace both there.

Update the two call sites of `run`/`runProto` in `main` to pass `*enchants`, and give `runProto` the same parameter (ignored — it has one caller and that caller runs a plain sim). Update `main_test.go`'s calls to `run` for the new parameter.

Update the command's doc comment: the usage lines gain `-enchants`, and a paragraph says the binary detects the request's kind and runs the plan–rank loop itself for a bulk one.

- [ ] **Step 5: Run the tests**

Run: `go test ./sim/cmd/forever-sim/... -v`
Expected: PASS. The bulk test runs the engine and takes tens of seconds; `-short` skips it.

- [ ] **Step 6: Vet, format, commit**

```bash
go vet ./sim/cmd/forever-sim/...
gofmt -w sim/cmd/forever-sim
git add sim/cmd/forever-sim
git commit -m "feat(sim): the native plan-rank loop and its stage progress lines

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 25: `forever-sim` runs weights and the Smart Sim loop

**Files:**
- Create: `sim/cmd/forever-sim/target.go`
- Create: `sim/cmd/forever-sim/target_test.go`
- Modify: `sim/cmd/forever-sim/main.go`

**Interfaces:**
- Consumes: `api.NeedsMoreIterations`, `api.NextStepIterations` (Task 4); `combine.Results`; `request.BuildWeights`, `simdb.AttachWeights`, `adapter.Weights` (Task 19).
- Produces:
  - `func ExecuteToTarget(req api.SimRequest, progress io.Writer) (api.SimResult, error)`
  - `func executeWeights(req api.SimRequest, progress io.Writer) (api.SimResult, error)`

- [ ] **Step 1: Write the failing test**

Create `sim/cmd/forever-sim/target_test.go`:

```go
package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// A fixed-count run is one sim: the loop must not turn today's
// requests into several.
func TestExecuteToTargetIsOneRunWithoutATarget(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.Iterations = 500
	req.Encounter.DurationSec = 60
	res, err := ExecuteToTarget(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.IterationsRun != 500 {
		t.Errorf("iterations_run = %d, want exactly the 500 asked for", res.IterationsRun)
	}
}

// A target-error run steps until the error is inside the target or the
// ceiling is reached, and reports the total.
func TestExecuteToTargetSteps(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.Encounter.DurationSec = 60
	// A target nothing reaches, so the ceiling ends it and the count
	// is exactly the ceiling rather than a step past it.
	req.TargetError = 0.00001
	req.Iterations = 3 * api.StepIterations
	res, err := ExecuteToTarget(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.IterationsRun != 3*api.StepIterations {
		t.Errorf("iterations_run = %d, want the ceiling %d", res.IterationsRun, 3*api.StepIterations)
	}
	if res.DPS.Mean <= 0 || res.DPS.Error <= 0 {
		t.Errorf("dps = %+v", res.DPS)
	}
	if res.Summary.EngineVersion == "" {
		t.Error("a stepped run lost its summary")
	}

	// A target anything reaches stops early, so the two runs differ.
	req.TargetError = 0.5
	early, err := ExecuteToTarget(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if early.IterationsRun != api.StepIterations {
		t.Errorf("a slack target ran %d iterations, want one step", early.IterationsRun)
	}
}

func TestExecuteWeights(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the engine")
	}
	req := fixtureRequest(t, "warrior-fury")
	req.Iterations = 500
	req.Encounter.DurationSec = 60
	req.Weights = &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit"},
		Reference: "attack_power",
	}
	res, err := executeWeights(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Weights) != 3 {
		t.Fatalf("got %d weights: %+v", len(res.Weights), res.Weights)
	}
	if res.Weights[1].Stat != "attack_power" || res.Weights[1].Weight != 1 {
		t.Errorf("the reference stat is %+v, want a weight of exactly 1", res.Weights[1])
	}
	if res.Lane != api.LaneServer || res.EngineVersion == "" {
		t.Errorf("provenance: lane=%q engine=%q", res.Lane, res.EngineVersion)
	}
	if res.Summary.DamageDone == nil {
		t.Error("a weights result carries a null summary; it must carry the empty one")
	}
}
```

- [ ] **Step 2: Run it to verify it fails**

Run: `go test ./sim/cmd/forever-sim/... -run 'ExecuteToTarget|ExecuteWeights' -v`
Expected: FAIL to compile — `undefined: ExecuteToTarget`.

- [ ] **Step 3: Write `sim/cmd/forever-sim/target.go`**

```go
package main

// Smart Sim and stat weights, natively.
//
// The stepping decision is api.NeedsMoreIterations, which the browser
// asks too - the loop is written twice because the two lanes drive it
// differently (a worker pool there, a for loop here), but the QUESTION
// is asked once, in Go, so the two stop at the same precision.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/combine"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// ExecuteToTarget is Execute, plus the target-error loop.
//
// A request with no target is one run, unchanged: this is the default
// path and it must stay exactly today's behaviour.
//
// A request with one runs steps of api.StepIterations and pools them
// with combine.Results - the same pooling the browser's worker pool
// uses for a sharded run, so a stepped run's error bar is computed the
// same way whichever lane produced it. Each step's seed is offset by
// the iterations before it, as combine.Split does, so the stream of
// a stepped run matches a single run of the same total.
func ExecuteToTarget(req api.SimRequest, progress io.Writer) (api.SimResult, error) {
	if req.TargetError <= 0 {
		return Execute(req, progress)
	}
	start := time.Now()
	var parts []api.SimResult
	pooled := api.SimResult{Request: req}
	seed := req.RandomSeed

	for {
		step := api.NextStepIterations(pooled, req)
		if step == 0 {
			break
		}
		part := req
		part.TargetError = 0
		part.Iterations = step
		part.RandomSeed = seed
		seed += int64(step)

		res, err := Execute(part, progress)
		if err != nil {
			if errors.Is(err, adapter.ErrAborted) && len(parts) > 0 {
				// What completed is still an answer, and the caller
				// asked for the stop.
				break
			}
			return api.SimResult{}, err
		}
		parts = append(parts, res)

		pooled, err = combine.Results(parts)
		if err != nil {
			return api.SimResult{}, fmt.Errorf("pooling the steps: %w", err)
		}
		// The pooled result's request is part zero's, which carries no
		// target; the run's own request is the one with the target,
		// and it is what the loop and the stored row must both see.
		pooled.Request = req
	}
	if len(parts) == 0 {
		return api.SimResult{}, errors.New("a target-error run made no steps; iterations must be at least one step")
	}
	pooled.EngineVersion = enginever.Version
	pooled.Lane = api.LaneServer
	pooled.DurationMS = time.Since(start).Milliseconds()
	return pooled, nil
}

// executeWeights computes stat weights natively, over the engine's own
// StatWeightsAsync - the same call the browser makes.
func executeWeights(req api.SimRequest, progress io.Writer) (api.SimResult, error) {
	registerOnce.Do(engineRegisterAll)
	start := time.Now()

	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return api.SimResult{}, fmt.Errorf("%w: %v", errBadInput, err)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return api.SimResult{}, err
	}

	reporter := make(chan *proto.ProgressMetrics, 32)
	id := runID()
	defer onInterrupt(id)()
	core.StatWeightsAsync(engineReq, reporter, id)

	var enc *json.Encoder
	if progress != nil {
		enc = json.NewEncoder(progress)
	}
	var engineRes *proto.StatWeightsResult
	for p := range reporter {
		if p.FinalWeightResult != nil {
			engineRes = p.FinalWeightResult
			break
		}
		if enc != nil {
			// A weights run is many sims, so the sim counts are the
			// honest progress and the iteration count alone would
			// restart per stat. They ride the same three fields a
			// bulk stage uses.
			_ = enc.Encode(struct {
				Completed   int32   `json:"completed"`
				Total       int32   `json:"total"`
				DPS         float64 `json:"dps"`
				CombosDone  int32   `json:"combos_done"`
				CombosTotal int32   `json:"combos_total"`
			}{p.CompletedIterations, p.TotalIterations, p.Dps, p.CompletedSims, p.TotalSims})
		}
	}
	if engineRes == nil {
		return api.SimResult{}, errors.New("the engine produced no weights")
	}
	weights, err := adapter.Weights(engineRes, req)
	if err != nil {
		return api.SimResult{}, err
	}
	return api.SimResult{
		EngineVersion: enginever.Version,
		Request:       req,
		Lane:          api.LaneServer,
		IterationsRun: req.Iterations,
		DurationMS:    time.Since(start).Milliseconds(),
		Summary:       adapter.EmptySummary(),
		Weights:       weights,
	}, nil
}
```

`registerOnce.Do(engineRegisterAll)` must match what `execute` in `main.go` already does — it calls `registerOnce.Do(engine.RegisterAll)`. Use the identical expression rather than inventing `engineRegisterAll`.

- [ ] **Step 4: Wire the two arms of `dispatch`**

In `main.go`, replace Task 24's placeholder arms with the real calls: `api.KindWeights` → `executeWeights(req, progress)`, default → `ExecuteToTarget(req, progress)`.

- [ ] **Step 5: Run the tests**

Run: `go test ./sim/cmd/forever-sim/... -v`
Expected: PASS. These run the engine and take a minute or two; `-short` skips them.

- [ ] **Step 6: Vet, format, commit**

```bash
go vet ./sim/cmd/forever-sim/...
gofmt -w sim/cmd/forever-sim
git add sim/cmd/forever-sim
git commit -m "feat(sim): the native Smart Sim loop and stat weights run

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 26: The Makefile targets and the pipeline

**Files:**
- Modify: `Makefile`
- Modify: `.github/workflows/sim.yml`
- Modify: `.gitignore` (the copied style file, if it is copied)

**Interfaces:**
- Consumes: `go run ./internal/genstyles` (Task 8); the seven wasm exports (Tasks 21, 22).
- Produces: `make sim-styles`, `make styles-check`, and a `sim.yml` that proves the style copy, the enchant-table path and a four-candidate Top Gear plan end to end.

- [ ] **Step 1: Add the style targets to the Makefile**

After the `simdb` block:

```make
WEB_SIM_DATA  = web/src/data/sim-styles.json
SIM_STYLES    = sim/request/styles.json

.PHONY: sim-styles
# sim-styles copies the fight-style table where the web reads it.
#
# The table is Go - a style expands to encounter fields, and the page
# applies the same expansion, so a second hand-kept copy in TypeScript
# would drift the first time a preset was retuned. `go run
# ./internal/genstyles` renders the Go map to sim/request/styles.json,
# this copies it to web/src/data/, and the web lane's own test
# compares its table against that file.
#
# Unlike simdb.bin, BOTH files are committed: the web build reads the
# copy, and a generated file the build depends on cannot be
# git-ignored on a runner that never runs make.
sim-styles: $(SIM_STYLES)
	@cp "$(SIM_STYLES)" "$(WEB_SIM_DATA)"
	@echo "copied $(SIM_STYLES) to $(WEB_SIM_DATA)"

$(SIM_STYLES): sim/request/styles.go
	@(cd sim && go run ./internal/genstyles)

.PHONY: styles-check
# styles-check proves both derived copies against the Go table, the way
# apl-check proves the rotations. A retuned preset that never reached
# the page is a failing pipeline rather than two products quietly
# disagreeing about what "Heavy movement" means.
styles-check:
	@(cd sim && go run ./internal/genstyles /tmp/styles.check.json)
	@diff -u "$(SIM_STYLES)" /tmp/styles.check.json || { \
	  echo "$(SIM_STYLES) is stale; run \`cd sim && go run ./internal/genstyles\`"; exit 1; }
	@diff -u "$(SIM_STYLES)" "$(WEB_SIM_DATA)" || { \
	  echo "$(WEB_SIM_DATA) is stale; run \`make sim-styles\`"; exit 1; }
	@echo "the fight-style table matches in all three places"
```

- [ ] **Step 2: Generate and copy, then check**

```bash
make sim-styles
make styles-check
```

Expected: `copied …`, then `the fight-style table matches in all three places`.

- [ ] **Step 3: Extend the sim workflow's paths and checks**

In `.github/workflows/sim.yml`, add to **both** `paths:` lists:

```
'data/builds/*/enchants.json', 'web/src/data/sim-styles.json'
```

and add a step to the `rotations` job, after the `apl-check` step (renaming the job's comment to say it proves both derived tables):

```yaml
      - run: make styles-check
```

`styles-check` needs Go, and the `rotations` job has none. Add the setup step at the top of that job, after the checkout:

```yaml
      - uses: actions/setup-go@v5
        with: { go-version-file: sim/go.mod, cache-dependency-path: sim/go.sum }
      - uses: ./.github/actions/pin-engine
      - name: the active build's item database
        run: make simdb
```

`sim/request/styles.go` imports `sim/api` only, but the package it lives in imports the engine, so the pin and the database are both needed to run the generator.

- [ ] **Step 4: Extend the wasm smoke test with a Top Gear plan**

In the `artifacts` job's `wasm smoke test` step, inside `globalThis.wasmready`, after the existing geared-run assertions and before the final `console.log`, add:

```js
            // Top Gear, end to end, at 100 iterations: simPlan, then
            // each stage request through the sharding a single run
            // uses, then simRank until it hands back a result. This
            // is the only gate that proves the planner compiled into
            // the wasm and that the three new exports agree about
            // their JSON.
            const gear = JSON.parse(gearJSON);
            const bulkReq = { ...geared, iterations: 3000, random_seed: 3,
              bulk: { mode: 'gear', precision: 'normal', cap: 400,
                candidates: [
                  { slot: 'head',   item_id: gear.find(g => g.slot === 'head').item_id,   origin: 'bag' },
                  { slot: 'neck',   item_id: gear.find(g => g.slot === 'neck').item_id,   origin: 'bag' },
                  { slot: 'wrist',  item_id: gear.find(g => g.slot === 'wrist').item_id,  origin: 'bag' },
                  { slot: 'hands',  item_id: gear.find(g => g.slot === 'hands').item_id,  origin: 'bag' },
                ] } };
            let stage = JSON.parse(globalThis.simPlan(JSON.stringify(bulkReq)));
            if (stage.error) { console.error('simPlan:', stage.error); process.exit(1); }
            // Four candidates in four distinct slots: every subset but
            // the empty one, which is the equipped set and is
            // requests[0].
            if (stage.combos.length !== 15) {
              console.error('simPlan planned', stage.combos.length, 'combinations, want 15'); process.exit(1);
            }
            if (stage.requests.length !== stage.combos.length + 1) {
              console.error('simPlan requests', stage.requests.length); process.exit(1);
            }
            let stagesRun = 0, topGear = null;
            while (topGear === null) {
              stagesRun++;
              if (stagesRun > 4) { console.error('the ladder did not finish'); process.exit(1); }
              // 100 iterations, not the stage's own count: this is a
              // smoke test of the plumbing, and sixteen sims at 1,000
              // would dominate the job.
              const cheap = stage.requests.map(r => ({ ...r, iterations: 100 }));
              const results = cheap.map((r, i) => {
                const parts = JSON.parse(globalThis.simSplit(JSON.stringify(r), 2));
                if (parts.error) { console.error('simSplit:', parts.error); process.exit(1); }
                const done = parts.map((p, j) => JSON.parse(globalThis.simRun(JSON.stringify(p), `tg${i}-${j}`)));
                for (const d of done) { if (d.error) { console.error('simRun:', d.error); process.exit(1); } }
                const one = JSON.parse(globalThis.simCombine(JSON.stringify(done)));
                if (one.error) { console.error('simCombine:', one.error); process.exit(1); }
                return one;
              });
              const ranked = JSON.parse(globalThis.simRank(
                JSON.stringify(bulkReq), JSON.stringify(stage), JSON.stringify(results)));
              if (ranked.error) { console.error('simRank:', ranked.error); process.exit(1); }
              if (ranked.result) { topGear = ranked.result; } else { stage = ranked.next; }
            }
            if (!topGear.equipped || !(topGear.equipped.mean > 0)) {
              console.error('no equipped baseline:', topGear.equipped); process.exit(1);
            }
            if (!topGear.combos || topGear.combos.length === 0) {
              console.error('no ranked combinations'); process.exit(1);
            }
            for (const combo of topGear.combos) {
              if (typeof combo.delta.mean !== 'number' || typeof combo.group !== 'number') {
                console.error('a combo is missing its delta or group:', combo); process.exit(1);
              }
            }
            if (topGear.stages.length !== 2) {
              console.error('stages:', topGear.stages); process.exit(1);
            }
```

and extend the final `console.log` and the export check:

```js
            const missing = ['simRun', 'simSplit', 'simCombine', 'simAbort', 'simPlan', 'simRank', 'simWeights']
              .filter(n => typeof globalThis[n] !== 'function');
```

```js
            console.log('seven exports, dps', one.dps.mean.toFixed(1), 'with',
                        one.summary.damage_done[0].abilities.length, 'abilities in the summary,',
                        progressSeen, 'progress calls, geared dps', g.dps.mean.toFixed(1),
                        '| top gear:', topGear.combos.length, 'ranked over', stagesRun, 'stages,',
                        'best delta', topGear.combos[0].delta.mean.toFixed(1));
```

- [ ] **Step 5: Run the smoke test locally before pushing**

```bash
make artifacts
cp artifacts/sim.wasm artifacts/sim.js /tmp/
cp sim/adapter/testdata/warrior-fury.request.json /tmp/
# paste the workflow's heredoc into /tmp/smoke.mjs, then:
cd /tmp && node smoke.mjs
```

Expected: the "seven exports … top gear: N ranked over 2 stages" line and exit 0. If the combination count is not 15, read the failure: it means one of the four slots' candidates was skipped, which is an eligibility bug and not a smoke-test bug.

- [ ] **Step 6: Run the whole module's suite once, as CI will**

```bash
(cd sim && go vet ./... && go test ./... -race)
(cd sim && gofmt -l .)
```

Expected: PASS, no vet findings, no unformatted files. Then check the coverage floor the pipeline enforces:

```bash
(cd sim && go test ./... -coverprofile=cover.out && go tool cover -func=cover.out | tail -1)
```

Expected: 80% or better. If `sim/bulk` drags it under, the gap is a rule with no table test — find it and add the case rather than lowering the floor.

- [ ] **Step 7: Commit**

```bash
git add Makefile .github/workflows/sim.yml web/src/data/sim-styles.json sim/request/styles.json
git commit -m "ci: prove the fight-style table and run a four-candidate Top Gear through the wasm

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

## Notes for the executor

**Order.** Tasks 1 and 2 are gates: nothing else compiles until the fork and the data lane have landed their side. Tasks 3–7 are the envelope and are independent of each other in spirit but touch one file, so run them in order. Tasks 8–11 (`sim/request`) and 12–18 (`simdb` and `sim/bulk`) are two lanes that do not touch the same files and can run in parallel once 3–7 are in. Tasks 19–22 need both. Tasks 23–25 need everything, and 26 is last.

**What this lane does not do.** It does not write the `data/builds/<build>/enchants.json` or `loot.json` files (data lane), does not touch `web/` except for the generated `sim-styles.json` copy, does not add the `kind` column or migration 0014 (api lane), and does not edit the engine fork.

**The one cross-lane break.** Task 23 changes `runner.Progress`, and the `api` module stops compiling until its call sites follow. That is the contract's section 8 requirement; the api lane's plan carries the matching change. Do not fix it from here.

**If a test in this plan asserts an item id that is not in the active build**, the build was regenerated. Find a replacement with the finder tests the plan describes, update the constant, and say so in the commit message. Never weaken the rule the test was proving.

---
