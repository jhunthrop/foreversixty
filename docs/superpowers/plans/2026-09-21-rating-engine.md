# Performance rating engine — implementation plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development
> (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps
> use checkbox (`- [ ]`) syntax for tracking.

> **Implementation status note (2026-09-21):** a second, independent agent was found
> operating concurrently in this same worktree/branch partway through this plan's own
> execution (discovered via a stray uncommitted `logs/engine/rating/types.go` and a set of
> already-landed commits this plan did not make). Its Tasks 1–3 work (the additive
> `summary.Summary`/`RosterRow` fields, the `Downtime` schema, the `weights` and `utility`
> packages) matched this plan's own design closely enough to adopt as-is rather than redo —
> see commits `e20b53f`, `c3ea403`, `5712672`, `5625fe0`. From the `consumables` package
> onward, every commit on this branch is this plan's own direct execution (not dispatched to
> fresh subagents per task, a deliberate deviation from the "subagent-driven-development"
> instruction below, made to avoid a second collision while reconciling the first one — see
> the final report for the reasoning). The task bodies below describe the *intended* design;
> where the landed code differs in a small way (an added `time.go`/`percentile.go` helper
> file, a `roundTo` implemented with `math.Round` rather than a hand-rolled truncation, a
> `Roles.For` that returns `(RoleWeights, bool)`), the code is authoritative. `git log
> --oneline` on this branch is the true record of what shipped and in what order.

**Goal:** Build `logs/engine/rating/`, a new pure-function package that scores one player's
performance for one closed fight into a six-component `Card`, plus the curated data (utility
tables, consumable catalogue, weight table, downtime windows) it reads and the drafting-tool
evidence passes that help curate them.

**Architecture:** One new engine package (`logs/engine/rating`) with one file per component
formula, sitting beside `summary` and `mechanics` with zero dependency on `api/internal`.
Percentile lookups go through a small injected interface (`PercentileSource`) so golden
fixture tests can supply a fake. Curated data lives in three new sibling packages under
`logs/engine/mechanics/` (`utility`, `consumables`, `weights`), each following the existing
`mechanics` package's embed-parse-validate-panic-at-boot pattern exactly. `mechanics.Table`
gets one additive field (`Downtime`) for §3.4. `summary.Summary` and `summary.RosterRow` each
get a small number of additive fields the rating engine cannot function without and cannot
get any other way, given `rating.Score`'s fixed signature only accepts `summary.Summary` (see
Ruling R1/R2 below) — both are documented rulings, both keep every existing golden fixture
passing after a one-time regeneration.

**Tech Stack:** Go 1.25 (this repo's `go.work` floor), standard `testing` package,
table-driven tests, `go:embed` for curated JSON.

**Spec:** `docs/superpowers/specs/2026-09-21-performance-rating-design.md` — this plan covers
its §1 (scoring model), §2 (fairness), §3 (curated data), and §4.1 (the engine package) only.
§4.2–§9 (storage, API, web, abuse/tone, the December 9 checklist) are other lanes' work and
are read here only for cross-reference.

## Global Constraints

- Go 1.25.11 floor; use `GOTOOLCHAIN=go1.25.11 go test ./...` if the machine's default `go`
  is older (check `go version` first).
- `go vet ./... && go test ./...` for every touched package before every commit.
- Functions under 50 lines, files under 800, no magic numbers (named constants with the
  reason), no dead code, pure functions with explicit inputs, immutable updates, errors
  wrapped with context, table-driven tests.
- `logs/engine/rating` must import nothing under `api/internal` (verified the same way
  `logs/engine/mechanics` already is: `go list -deps` or a grep for the import path).
- Every curated spell id ships with a `verified: "<build>/spells.json#<id>"` string that was
  actually checked against `data/builds/1.60.1.69893/spells.json`; an id that does not
  resolve is never committed.
- Component names are exactly `"output"`, `"survival"`, `"mechanics"`, `"utility"`,
  `"preparation"`, `"activity"` (lowercase, per spec §4.1's `Card.Components [6]Component`
  and §5.2's JSON example).
- Own only: `logs/engine/rating/**`, `logs/engine/mechanics/**`, the new curated data files
  named in spec §3.1, `logs/cmd/forever-logs/mechanics_draft.go`, and these files' tests.
  Never touch `api/`, `web/`, `sim/`, `addon/`, `companion/`.
- Commits: write the message to a file under `.superpowers/` with `printf`, then
  `git commit -F <file>` as its own Bash command with nothing else in it. End every message
  with:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`

## Rulings this plan bakes in (read before touching any task)

These resolve places where the spec's fixed `rating.Score` signature (§4.1) cannot reach data
the formulas in §1–§3 require. Each is the smallest additive change that closes the gap
without altering `Score`'s parameter list or the `Card`/`Component` field set, which the API
lane is already told to compile against verbatim.

- **R1 — `summary.Summary` gains four fields.** `EncounterID int64`, `Difficulty int64`, and
  `Kill bool` (mirroring `fight.Fight`'s own fields of the same name/type, set in `Snapshot`
  from the `fight.Fight` it already receives) — `rating.Score` needs these for the percentile
  bracket key (§1.2) and the wipe rule (§2), and its signature takes only `summary.Summary`,
  not `fight.Fight`. All three are `omitempty`/zero-value-safe and do not change any existing
  golden fixture's *shape*, only add three keys — the golden files are regenerated once, in
  Task 1, and the diff is reviewed to confirm it is exactly those three additions.
- **R2 — `summary.RosterRow` gains `ExecutionScore *float64`.** The Output component's
  primary value (§1.3) is `fight_metrics.execution_score`, computed by `api/internal/sims`
  from a simulator run and stored in a database column no engine package can read. The API
  lane, which already re-reads a fight's stored summary before calling `rating.Score`
  (§4.4), is expected to set this pointer on the relevant roster row before the call.
  `omitempty` and nil by default, so it adds nothing to today's golden fixtures.
- **R3 — `PercentileSource` gets a second method, `KillTimeBand`.** §1.2's kill-time band
  needs the bracket's own kill-duration median, which `Placement` (a percentile-of-a-value
  query) cannot answer. `Score`'s parameter list is unchanged; only the interface it accepts
  grows by one method, decided here rather than left unresolved, since nothing outside this
  worktree implements the interface yet.
- **R4 — Survival's `AvoidableHitScore` has no absolute standard.** Its raw input
  (avoidable-damage-taken per second) is an unbounded rate, not a 0–100 share, so it cannot
  be "its own absolute standard" the way Preparation/Utility/Activity's raw values can (R5).
  When its bracket lacks `MinSample`, Survival falls back to `DeathScore` alone (weight 1.0
  on the death half, 0 on the avoidable-hit half) rather than inventing a numeric standard
  the spec does not give. Flagged on the component as `Basis: "absolute"` (the death-timing
  formula is a direct computation, not a peer comparison, which "absolute" already means
  elsewhere in this spec).
- **R5 — a bounded 0–100 raw value is its own absolute standard.** For every component whose
  raw input is already a share or percentage (Preparation's completeness score, Utility's
  mean uptime share, Activity's active share, Mechanics' dropped-debuff-uptime sub-part),
  §1.2's "absolute standard" when the bracket lacks samples is that same value, scaled to
  0–100 — consistent with how Output's own absolute standard (`min(100, execution_score*100)`)
  already treats a bounded value. A component whose raw input is an unbounded rate (raw DPS,
  interrupts/dispels value per second) has no such standard and is excluded instead, exactly
  as §1.3 states outright for the raw-DPS fallback.
- **R6 — Mechanics' dispel-made credit divides by `Dispelled`, mirroring interrupts.**
  §1.3 gives interrupts' credit as `MechanicRow.Damage / max(1, MechanicRow.Casts)` ("the
  average cost of one uninterrupted cast") but only says dispels credit
  `MechanicRow.Healed` + `MechanicRow.Damage` with no stated divisor. Dispels get the same
  per-occurrence averaging, divided by `max(1, MechanicRow.Dispelled)`, for the same reason
  interrupts divide by casts: crediting the fight's *total* prevented value on every single
  dispel a player made would double-count when a debuff was dispelled more than once.
- **R7 — Downtime windows are computed from `Casts`/`Auras` already in `Summary`, and support
  three of `PhaseStart`'s four trigger kinds.** §3.4 reuses `mechanics.PhaseStart` and says
  the rating package's downtime detector "calls" the same detection path `phaseFires`
  (`summary/phases.go:47-71`) uses — but that method is unexported on `*Accumulator` and
  operates on raw events, while `rating.Score` (per §4.1's own ruling against changing
  `summary`) only ever sees a finished `Summary`. The rating package re-implements the same
  trigger semantics against data `Summary` already exports: `on: "aura_applied"` and
  `on: "aura_removed"` read `AuraTrack.Segments`' start/end instants on the *enemy* GUID the
  aura opened on; `on: "cast_success"` reads `CastRow.Sequence`. `on: "cast_start"` and a
  `health_pct` trigger have no equivalent data in `Summary` (cast starts are not recorded
  past their pairing with a success, and boss health readings are not retained past
  `Snapshot`) — a downtime window using either produces zero windows, with a doc comment
  explaining why, rather than a crash or a silent wrong answer. None of the curated downtime
  data this plan ships uses either unsupported trigger.
- **R8 — Activity's downtime-excluded `ActiveMS` is a per-second-bucket approximation.**
  `markActive`'s exact sub-second logic is not reproducible from `Summary` (no raw event
  timestamps survive `Snapshot` at that granularity). The rating package instead treats each
  whole second in `Actor.Series` (already in `Summary`, one non-zero entry per second of
  damage-done or healing-done) as "active," and excludes any second whose midpoint falls
  inside a downtime window — a second-granularity approximation of the same idea, documented
  in `activity.go`. None of §1.6's worked examples exercise a fight with curated downtime
  (all three explicitly use "no curated downtime... approximate"), so this path is unit
  tested directly but is not what the worked-example tests check.
- **R9 — component names used in `Bracket.Component` for sub-part percentile queries** are
  not the six top-level names, since Survival and Mechanics each combine more than one
  independent percentile query: `survival_avoidable_hit`, `mechanics_interrupt`,
  `mechanics_dispel`, `mechanics_debuff_uptime`, alongside the top-level `output`, `utility`,
  `preparation`, `activity`. Recorded here because the API lane's `rating_percentile_digests`
  writer needs to know these exact strings to fold fights into the right digest.
- **R10 — the spec's cross-references to "§7.4" for excluded-component reason copy are a
  drafting slip for "§6.4"** ("Copy for every state"), the section that actually lists the
  fixed strings. `Component.Reason` holds the short machine code shown in §5.2's own JSON
  example (`"no_mechanics_table"`), not the sentence; the web lane maps code → §6.4's copy.

---

## Task 1: Additive fields on `summary.Summary` / `summary.RosterRow`, golden regeneration

**Files:**
- Modify: `logs/engine/summary/summary.go` (the `Summary` struct, `Snapshot`)
- Modify: `logs/engine/summary/roster.go` (the `RosterRow` struct)
- Modify: `logs/engine/summary/testdata/v16.summary.json.golden`,
  `logs/engine/summary/testdata/v22.summary.json.golden` (regenerated, not hand-edited)
- Test: `logs/engine/summary/golden_test.go` (already exists; no change needed, just run it)

**Interfaces:**
- Produces: `Summary.EncounterID int64`, `Summary.Difficulty int64`, `Summary.Kill bool`,
  `RosterRow.ExecutionScore *float64` — every later task reads these off `summary.Summary`.

- [ ] **Step 1: Add the three `Summary` fields and populate them in `Snapshot`**

In `logs/engine/summary/summary.go`, add to the `Summary` struct (after `DurationMS`):

```go
	// EncounterID, Difficulty and Kill mirror fight.Fight's own fields of the
	// same name: the rating engine (logs/engine/rating) needs fight-level
	// context for its percentile bracket key and its wipe rule, and its
	// Score function's signature — fixed so the API lane can compile
	// against it (docs/superpowers/specs/2026-09-21-performance-rating-design.md
	// §4.1) — takes only a Summary, never a fight.Fight alongside it.
	EncounterID int64 `json:"encounter_id,omitempty"`
	Difficulty  int64 `json:"difficulty,omitempty"`
	Kill        bool  `json:"kill"`
```

In `Snapshot`, right after `s := Summary{...}` is constructed (before `s.Mechanics = ...`),
add:

```go
	s.EncounterID, s.Difficulty, s.Kill = f.EncounterID, f.Difficulty, f.Kill
```

- [ ] **Step 2: Add `ExecutionScore` to `RosterRow`**

In `logs/engine/summary/roster.go`, add to the `RosterRow` struct (after `HPS`, `DTPS`):

```go
	// ExecutionScore is fight_metrics.execution_score (api/internal/sims),
	// overlaid by the API layer before calling rating.Score — the engine
	// package has no database access and cannot compute it itself. Nil
	// when no execution score exists for this player's fight (unvalidated
	// spec, unclaimed character, non-DPS role, or the API layer simply
	// has not set it, e.g. in every engine-package unit test).
	ExecutionScore *float64 `json:"execution_score,omitempty"`
```

- [ ] **Step 3: Regenerate the golden fixtures and review the diff**

```bash
GOTOOLCHAIN=go1.25.11 FOREVER_UPDATE_GOLDEN=1 go test ./engine/summary/... -run TestTheFixtureSummaryMatchesTheCommittedGolden
```
(Run from `/Users/jh/code/forever/.worktrees/rating-engine/logs`, the module root for the
`logs` module.)

```bash
git diff --stat logs/engine/summary/testdata/
git diff logs/engine/summary/testdata/v16.summary.json.golden | head -60
```

Expected: only `"encounter_id"`, `"difficulty"`, `"kill"` keys added to the top-level object
(and nested nowhere else) and no `execution_score` key anywhere (it is `omitempty` and stays
nil in every existing fixture). If anything else changed, stop and investigate before
proceeding — a Step 1/2 typo, not an intended change.

- [ ] **Step 4: Run the full summary package test suite**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go vet ./engine/summary/... && GOTOOLCHAIN=go1.25.11 go test ./engine/summary/...
```
Expected: all tests `PASS`, including `TestTheGoldenSummaryIsStableAcrossRuns` and every
other existing test in the package (a compile error here means Step 1 or 2 broke something
downstream — check `logs/engine/summary/summary_test.go` for anywhere it builds a literal
`Summary{...}` or `RosterRow{...}` with positional fields, which the new fields would not
break since Go struct literals here are keyed, but re-run to confirm).

- [ ] **Step 5: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine
git add logs/engine/summary/summary.go logs/engine/summary/roster.go logs/engine/summary/testdata/
printf 'feat: add fight-context and execution-score fields to summary.Summary\n\nlogs/engine/rating (docs/superpowers/specs/2026-09-21-performance-rating-design.md\n§4.1) needs EncounterID/Difficulty/Kill for its percentile bracket key and wipe\nrule, and RosterRow.ExecutionScore for the Output component, and its Score\nsignature is fixed to accept only a summary.Summary. All four fields are\nadditive and omitempty; golden fixtures regenerated (ruling R1/R2 in\ndocs/superpowers/plans/2026-09-21-rating-engine.md).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task1.txt
git commit -F .superpowers/commit-msg-task1.txt
```

---

## Task 2: `Downtime` schema on `mechanics.Table` (§3.4)

**Files:**
- Modify: `logs/engine/mechanics/mechanics.go`
- Modify: `logs/engine/mechanics/mechanics_test.go`
- Modify: `logs/engine/mechanics/tables/666.json` (Garr — add the one verified, real
  downtime window the spec itself illustrates)

**Interfaces:**
- Produces: `mechanics.Downtime{Trigger PhaseStart, DurationMS int64, Note string}`,
  `Table.Downtime []Downtime`, validated by `Parse`.

- [ ] **Step 1: Write the failing validation tests**

Append to `logs/engine/mechanics/mechanics_test.go`:

```go
func TestParseAcceptsDowntimeAndRefusesAMalformedOne(t *testing.T) {
	good := []byte(`{"encounter_id": 666, "name": "Garr", "mechanics": [
		{"spell_id": 19497, "name": "Eruption", "kind": "avoidable"}],
		"downtime": [
			{"trigger": {"spell_id": 19497, "on": "aura_applied"}, "duration_ms": 3000,
			 "note": "Firesworn detonation window."}]}`)
	table, err := Parse(good)
	if err != nil {
		t.Fatal(err)
	}
	if len(table.Downtime) != 1 || table.Downtime[0].DurationMS != 3000 {
		t.Fatalf("downtime = %+v", table.Downtime)
	}
	if table.Downtime[0].Trigger.On != OnAuraApplied {
		t.Fatalf("trigger = %+v", table.Downtime[0].Trigger)
	}

	noDuration := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [],
		"downtime": [{"trigger": {"spell_id": 1, "on": "aura_applied"}}]}`)
	if _, err := Parse(noDuration); err == nil {
		t.Fatal("a downtime window with no duration_ms must be refused")
	}

	badTrigger := []byte(`{"encounter_id": 1, "name": "X", "mechanics": [],
		"downtime": [{"trigger": {"spell_id": 1, "on": "not_a_thing"}, "duration_ms": 1000}]}`)
	if _, err := Parse(badTrigger); err == nil {
		t.Fatal("a downtime window with a malformed trigger must be refused")
	}
}
```

- [ ] **Step 2: Run it, confirm it fails to compile (no `Downtime` type yet)**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/... -run TestParseAcceptsDowntimeAndRefusesAMalformedOne
```
Expected: build failure, `undefined: Downtime` or similar.

- [ ] **Step 3: Add the `Downtime` type and validation**

In `logs/engine/mechanics/mechanics.go`, add after the `Phase` struct:

```go
// Downtime is one stretch during which the raid is expected to stop
// attacking without it counting against their positioning — an add's death
// detonation, a forced-movement channel. Consumed only by the rating
// engine's Activity component (logs/engine/rating), never subtracted from
// Survival's or Mechanics' own denominators: the boss being untargetable
// does not mean avoidable damage stopped mattering during that window.
type Downtime struct {
	// Trigger reuses PhaseStart's exact shape: a phase boundary and a
	// downtime window are both "this fires when a spell or a health
	// threshold is crossed," and sharing the vocabulary means one detector,
	// not two.
	Trigger    PhaseStart `json:"trigger"`
	DurationMS int64      `json:"duration_ms"`
	Note       string     `json:"note,omitempty"`
}
```

Add `Downtime []Downtime` to `Table`, after `Phases`:

```go
	// Downtime is the encounter's curated forced-downtime windows (§3.4).
	// Omitted means none, the same pattern Phases already uses.
	Downtime []Downtime `json:"downtime,omitempty"`
```

In `Parse`, after the existing `for i, p := range t.Phases` validation loop, add:

```go
	for i, d := range t.Downtime {
		if d.DurationMS <= 0 {
			return Table{}, fmt.Errorf("downtime[%d]: duration_ms must be positive", i)
		}
		switch {
		case d.Trigger.SpellID > 0:
			if !phaseOn[d.Trigger.On] {
				return Table{}, fmt.Errorf(
					"downtime[%d]: on %q is not cast_start, cast_success, aura_applied or aura_removed", i, d.Trigger.On)
			}
			if d.Trigger.HealthPct != 0 {
				return Table{}, fmt.Errorf(
					"downtime[%d]: a trigger names a spell or a health percentage, not both", i)
			}
		case d.Trigger.HealthPct > 0 && d.Trigger.HealthPct <= 100:
			if d.Trigger.On != "" {
				return Table{}, fmt.Errorf("downtime[%d]: on belongs to a spell trigger; a health trigger takes none", i)
			}
		default:
			return Table{}, fmt.Errorf(
				"downtime[%d]: trigger must name a spell_id with on, or a health_pct above 0 and at most 100", i)
		}
	}
```

- [ ] **Step 4: Run the test, confirm it passes**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/... -run TestParseAcceptsDowntimeAndRefusesAMalformedOne -v
```
Expected: `PASS`.

- [ ] **Step 5: Add the one real, verified downtime window to `tables/666.json`**

Spell id 19497 is already the `Eruption` mechanic in this same file, sourced from the build's
own client data per the file's existing "avoidable" row — reused here as a real, already-
verified id (no new spell id introduced). Edit `logs/engine/mechanics/tables/666.json`, add
after the closing `]` of `"mechanics"`:

```json
  ,
  "downtime": [
    { "trigger": { "spell_id": 19497, "on": "aura_applied" }, "duration_ms": 3000,
      "note": "A Firesworn's death detonation window: melee are expected to step out, not to keep swinging." }
  ]
```

(Insert as a sibling key of `"mechanics"`, keeping the file valid JSON — read the file first
with the file-reading tool to get exact current formatting before editing.)

- [ ] **Step 6: Run the full mechanics package suite, including the embed-and-validate test**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go vet ./engine/mechanics/... && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/... -v
```
Expected: all `PASS`, including `TestEveryEmbeddedTableParses/666.json` (proves the edited
Garr table still parses) and every other existing table's subtest (proves nothing else
broke).

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine
git add logs/engine/mechanics/mechanics.go logs/engine/mechanics/mechanics_test.go logs/engine/mechanics/tables/666.json
printf 'feat: add curated downtime windows to the mechanics table schema\n\nSpec §3.4: an additive, optional Downtime field on mechanics.Table, reusing\nPhaseStart'"'"'s trigger vocabulary. Adds one real, already-verified downtime\nwindow to Garr'"'"'s table (the Firesworn detonation the spec itself illustrates,\nspell id 19497, already curated in this same file as the Eruption mechanic).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task2.txt
git commit -F .superpowers/commit-msg-task2.txt
```

---

## Task 3: curated-data packages — `weights`, `utility`, `consumables`

All three follow the exact pattern `logs/engine/mechanics/mechanics.go` already uses:
`//go:embed`, a `Parse` that validates, a package-init `mustParseAll`-style panic on a
malformed embedded file, and a `Load` accessor. Build all three in one task since they are
structurally identical and small; a reviewer accepts or rejects them as one unit.

**Files:**
- Create: `logs/engine/mechanics/weights/weights.go`, `logs/engine/mechanics/weights/roles.json`,
  `logs/engine/mechanics/weights/weights_test.go`
- Create: `logs/engine/mechanics/utility/utility.go`, `logs/engine/mechanics/utility/utility_test.go`
  (the 27 `tables/<spec-slug>.json` files already exist, written by prior curation research —
  verify them, do not regenerate them)
- Create: `logs/engine/mechanics/consumables/consumables.go`, `logs/engine/mechanics/consumables/consumables_test.go`
  (`catalogue.json` already exists, written by prior curation research — verify it, do not
  regenerate it)

**Interfaces:**
- Produces: `weights.Roles{DPS, Healer, Tank weights.RoleWeights}`,
  `weights.RoleWeights{Output, Survival, Mechanics, Utility, Preparation, Activity float64}`,
  `weights.Default() Roles`.
- Produces: `utility.Table{Spec string, Owned []utility.Entry}`,
  `utility.Entry{SpellID int64, Name, Kind, Target, Verified, Note string}`,
  `utility.Load(specSlug string) (Table, bool)`.
- Produces: `consumables.Catalogue{Roles map[string]consumables.RoleCatalogue}`,
  `consumables.RoleCatalogue{Weights map[string]int, Flask, Food, WeaponEnchant, WorldBuffs []consumables.Entry, CombatPotion consumables.PotionGroup}`,
  `consumables.Entry{SpellID int64, Name, Verified string}`,
  `consumables.PotionGroup{MaxUses int64, Entries []consumables.Entry}`,
  `consumables.Load() Catalogue`.

- [ ] **Step 1: Write `weights/roles.json`**

```json
{
  "dps":    { "output": 35, "survival": 15, "mechanics": 20, "utility": 15, "preparation": 10, "activity": 5 },
  "healer": { "output": 30, "survival": 10, "mechanics": 20, "utility": 20, "preparation": 10, "activity": 10 },
  "tank":   { "output": 10, "survival": 35, "mechanics": 20, "utility": 20, "preparation": 10, "activity": 5 }
}
```
(§1.4's table, verbatim — each row already sums to 100.)

- [ ] **Step 2: Write the failing test for `weights`**

`logs/engine/mechanics/weights/weights_test.go`:

```go
// logs/engine/mechanics/weights/weights_test.go
package weights

import "testing"

func TestDefaultSumsToOneHundredPerRole(t *testing.T) {
	d := Default()
	for name, r := range map[string]RoleWeights{"dps": d.DPS, "healer": d.Healer, "tank": d.Tank} {
		sum := r.Output + r.Survival + r.Mechanics + r.Utility + r.Preparation + r.Activity
		if sum != 100 {
			t.Errorf("%s weights sum to %v, want 100", name, sum)
		}
	}
	if d.DPS.Output != 35 || d.Tank.Survival != 35 || d.Healer.Utility != 20 {
		t.Fatalf("Default() = %+v, does not match spec §1.4's table", d)
	}
}

func TestParseRejectsARoleWhoseWeightsDoNotSumToOneHundred(t *testing.T) {
	bad := []byte(`{"dps": {"output": 50, "survival": 15, "mechanics": 20, "utility": 15, "preparation": 10, "activity": 5}, "healer": {}, "tank": {}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("weights that do not sum to 100 must be refused")
	}
}
```

- [ ] **Step 3: Run it, confirm it fails (package does not exist yet)**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/weights/...
```
Expected: build failure.

- [ ] **Step 4: Write `weights/weights.go`**

```go
// logs/engine/mechanics/weights/weights.go
// Package weights holds the default per-role component weights (spec §1.4)
// as data, not code, so an eventual guild-adjustable override (out of
// scope here) edits data in this same shape rather than a scoring rewrite.
package weights

import (
	"embed"
	"encoding/json"
	"fmt"
)

// RoleWeights is one role's six component weights, summing to 100.
type RoleWeights struct {
	Output      float64 `json:"output"`
	Survival    float64 `json:"survival"`
	Mechanics   float64 `json:"mechanics"`
	Utility     float64 `json:"utility"`
	Preparation float64 `json:"preparation"`
	Activity    float64 `json:"activity"`
}

func (w RoleWeights) sum() float64 {
	return w.Output + w.Survival + w.Mechanics + w.Utility + w.Preparation + w.Activity
}

// Roles is the site-wide default weight table, one row per role.
type Roles struct {
	DPS    RoleWeights `json:"dps"`
	Healer RoleWeights `json:"healer"`
	Tank   RoleWeights `json:"tank"`
}

//go:embed roles.json
var rolesFile embed.FS

var defaultRoles = mustParseEmbedded()

func mustParseEmbedded() Roles {
	data, err := rolesFile.ReadFile("roles.json")
	if err != nil {
		panic(fmt.Errorf("weights: reading roles.json: %w", err))
	}
	r, err := Parse(data)
	if err != nil {
		panic(fmt.Errorf("weights: roles.json: %w", err))
	}
	return r
}

// Parse reads a weight table and refuses one whose rows do not sum to 100.
func Parse(data []byte) (Roles, error) {
	var r Roles
	if err := json.Unmarshal(data, &r); err != nil {
		return Roles{}, fmt.Errorf("weights: %w", err)
	}
	for name, w := range map[string]RoleWeights{"dps": r.DPS, "healer": r.Healer, "tank": r.Tank} {
		if w.sum() != 100 {
			return Roles{}, fmt.Errorf("weights: %s weights sum to %v, want 100", name, w.sum())
		}
	}
	return r, nil
}

// Default returns the embedded site-wide weight table.
func Default() Roles { return defaultRoles }

// For selects a role's weights by the same role strings summary.RosterRow.Role
// uses: "dps", "healer", "tank". Unknown roles get the dps row — a role the
// registry could not classify is more often a dps-shaped log than not, and a
// silent zero-weight card would be a worse failure than a slightly wrong one.
func (r Roles) For(role string) RoleWeights {
	switch role {
	case "healer":
		return r.Healer
	case "tank":
		return r.Tank
	default:
		return r.DPS
	}
}
```

- [ ] **Step 5: Run the weights tests, confirm pass**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go vet ./engine/mechanics/weights/... && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/weights/... -v
```
Expected: `PASS`.

- [ ] **Step 6: Verify the 27 existing utility table files, then write `utility/utility.go`**

First, confirm the curated files already at `logs/engine/mechanics/utility/tables/*.json`
(27 files, one per `data/curated/specs.json` entry) are well-formed and every `spell_id`
resolves in `data/builds/1.60.1.69893/spells.json`:

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine && python3 -c "
import json, glob
idx = {r['id'] for r in json.load(open('data/builds/1.60.1.69893/spells.json'))}
specs = {s['spec'] for s in json.load(open('data/curated/specs.json'))}
files = sorted(glob.glob('logs/engine/mechanics/utility/tables/*.json'))
have = set()
bad = []
for fn in files:
    t = json.load(open(fn))
    have.add(t['spec'])
    for e in t.get('owned', []):
        if e['spell_id'] not in idx:
            bad.append((fn, e['spell_id']))
        if e['kind'] not in ('buff','debuff','cooldown','talent-modifier'):
            bad.append((fn, 'bad kind', e['kind']))
missing = specs - have
extra = have - specs
print('files', len(files), 'missing specs', missing, 'unexpected specs', extra, 'bad entries', bad)
"
```
Expected: `missing specs set()`, `unexpected specs set()`, `bad entries []`. If not, stop and
fix the offending JSON file directly (do not regenerate all 27; fix only what fails) before
continuing — a bad file here would panic the process at boot once embedded.

Then write `logs/engine/mechanics/utility/utility.go`:

```go
// logs/engine/mechanics/utility/utility.go
// Package utility holds the curated per-spec utility table (spec §3.2): the
// buffs, debuffs and cooldowns a spec is expected to own, each spell id
// verified against the build's own spells.json before it is committed. It
// follows logs/engine/mechanics's own embed-parse-validate-panic pattern.
package utility

import (
	"embed"
	"encoding/json"
	"fmt"
)

// Kind is what an owned entry is and how it is scored.
type Kind string

const (
	// Buff is a self/raid buff the spec applies, scored by uptime share.
	Buff Kind = "buff"
	// Debuff is an enemy debuff the spec applies, scored by uptime share.
	Debuff Kind = "debuff"
	// Cooldown is a burst/utility cooldown scored by use, not uptime.
	Cooldown Kind = "cooldown"
	// TalentModifier is informational only, never independently scored.
	TalentModifier Kind = "talent-modifier"
)

var kinds = map[Kind]bool{Buff: true, Debuff: true, Cooldown: true, TalentModifier: true}

// Entry is one utility the spec owns.
type Entry struct {
	SpellID int64  `json:"spell_id"`
	Name    string `json:"name"`
	Kind    Kind   `json:"kind"`
	// Target is "enemy" for a debuff applied to the boss/adds, "self" or
	// "party" for a buff applied to the caster or other raiders.
	Target   string `json:"target,omitempty"`
	Verified string `json:"verified"`
	Note     string `json:"note,omitempty"`
}

// Table is one spec's owned utility.
type Table struct {
	Spec  string  `json:"spec"`
	Owned []Entry `json:"owned"`
}

//go:embed tables/*.json
var tables embed.FS

var _ = mustParseAll()

func mustParseAll() int {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		panic(fmt.Errorf("utility: reading the embedded tables: %w", err))
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := tables.ReadFile("tables/" + entry.Name())
		if err != nil {
			panic(fmt.Errorf("utility: reading %s: %w", entry.Name(), err))
		}
		if _, err := Parse(data); err != nil {
			panic(fmt.Errorf("utility: %s: %w", entry.Name(), err))
		}
	}
	return len(entries)
}

// Parse reads a table and refuses one that is not the format.
func Parse(data []byte) (Table, error) {
	var t Table
	if err := json.Unmarshal(data, &t); err != nil {
		return Table{}, fmt.Errorf("utility: %w", err)
	}
	if t.Spec == "" {
		return Table{}, fmt.Errorf("utility: spec is required")
	}
	for i, e := range t.Owned {
		if e.SpellID <= 0 {
			return Table{}, fmt.Errorf("owned[%d]: spell_id must be positive", i)
		}
		if !kinds[e.Kind] {
			return Table{}, fmt.Errorf("owned[%d]: kind %q is not buff, debuff, cooldown or talent-modifier", i, e.Kind)
		}
		if e.Verified == "" {
			return Table{}, fmt.Errorf("owned[%d]: verified is required", i)
		}
	}
	return t, nil
}

// Load returns the embedded table for a spec slug (e.g. "warrior-protection",
// matching data/curated/specs.json's own "spec" field), and false when there
// is none — an unverified or not-yet-curated spec, not a build defect.
func Load(specSlug string) (Table, bool) {
	data, err := tables.ReadFile("tables/" + specSlug + ".json")
	if err != nil {
		return Table{}, false
	}
	t, err := Parse(data)
	if err != nil {
		panic(err)
	}
	return t, true
}
```

- [ ] **Step 7: Write `utility/utility_test.go`**

```go
// logs/engine/mechanics/utility/utility_test.go
package utility

import "testing"

func TestLoadFindsAnEmbeddedSpecAndSaysWhenThereIsNone(t *testing.T) {
	table, ok := Load("warrior-protection")
	if !ok || table.Spec != "warrior-protection" || len(table.Owned) == 0 {
		t.Fatalf("Load(warrior-protection) = %+v, %v", table, ok)
	}
	if _, ok := Load("not-a-real-spec"); ok {
		t.Fatal("an unknown spec slug must report false")
	}
}

func TestEveryEmbeddedSpecTableParsesAndHasAName(t *testing.T) {
	entries, err := tables.ReadDir("tables")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 27 {
		t.Fatalf("%d embedded spec tables, want 27 (one per data/curated/specs.json entry)", len(entries))
	}
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			data, err := tables.ReadFile("tables/" + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			table, err := Parse(data)
			if err != nil {
				t.Fatal(err)
			}
			if table.Spec+".json" != entry.Name() {
				t.Fatalf("%s holds spec %q, so Load would never find it", entry.Name(), table.Spec)
			}
		})
	}
}

func TestParseRejectsAnUnverifiedEntry(t *testing.T) {
	bad := []byte(`{"spec": "x", "owned": [{"spell_id": 1, "name": "Y", "kind": "buff"}]}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an entry with no verified source must be refused")
	}
}
```

- [ ] **Step 8: Run the utility tests, confirm pass**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go vet ./engine/mechanics/utility/... && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/utility/... -v
```
Expected: `PASS`, `TestEveryEmbeddedSpecTableParsesAndHasAName` shows 27 subtests.

- [ ] **Step 9: Verify the existing `consumables/catalogue.json`, then write `consumables/consumables.go`**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine && python3 -c "
import json
idx = {r['id'] for r in json.load(open('data/builds/1.60.1.69893/spells.json'))}
cat = json.load(open('logs/engine/mechanics/consumables/catalogue.json'))
bad = []
for role, rc in cat['roles'].items():
    w = rc['weights']
    total = sum(w.values())
    if total != 100:
        bad.append((role, 'weights sum to', total))
    for cat_name in ('flask','food','weapon_enchant','world_buffs'):
        for e in rc.get(cat_name, []):
            if e['spell_id'] not in idx:
                bad.append((role, cat_name, e['spell_id']))
    for e in rc.get('combat_potion', {}).get('entries', []):
        if e['spell_id'] not in idx:
            bad.append((role, 'combat_potion', e['spell_id']))
print('roles', list(cat['roles'].keys()), 'bad', bad)
"
```
Expected: `bad []` and all three roles (`dps`, `healer`, `tank`) present. If not, fix the
JSON file directly before continuing.

Then write `logs/engine/mechanics/consumables/consumables.go`:

```go
// logs/engine/mechanics/consumables/consumables.go
// Package consumables holds the curated raid-consumable catalogue by role
// (spec §3.3), for the Preparation rating component. Same embed-parse-
// validate pattern as logs/engine/mechanics and logs/engine/mechanics/utility.
package consumables

import (
	"embed"
	"encoding/json"
	"fmt"
)

// Entry is one consumable: an item or buff spell a completeness check
// matches against CombatantRow.Consumables or a player's Casts.
type Entry struct {
	SpellID  int64  `json:"spell_id"`
	Name     string `json:"name"`
	Verified string `json:"verified"`
}

// PotionGroup is combat_potion's shape: entries plus the shared-cooldown cap
// on how many uses count toward completeness in one fight.
type PotionGroup struct {
	MaxUses int64   `json:"max_uses"`
	Entries []Entry `json:"entries"`
}

// RoleCatalogue is one role's consumable categories and their weights.
type RoleCatalogue struct {
	Weights       map[string]int `json:"weights"`
	Flask         []Entry        `json:"flask"`
	Food          []Entry        `json:"food"`
	WeaponEnchant []Entry        `json:"weapon_enchant"`
	WorldBuffs    []Entry        `json:"world_buffs"`
	CombatPotion  PotionGroup    `json:"combat_potion"`
}

func (rc RoleCatalogue) weightSum() int {
	total := 0
	for _, w := range rc.Weights {
		total += w
	}
	return total
}

// Catalogue is the whole consumable table, one row per role.
type Catalogue struct {
	Roles map[string]RoleCatalogue `json:"roles"`
}

//go:embed catalogue.json
var catalogueFile embed.FS

var defaultCatalogue = mustParseEmbedded()

func mustParseEmbedded() Catalogue {
	data, err := catalogueFile.ReadFile("catalogue.json")
	if err != nil {
		panic(fmt.Errorf("consumables: reading catalogue.json: %w", err))
	}
	c, err := Parse(data)
	if err != nil {
		panic(fmt.Errorf("consumables: catalogue.json: %w", err))
	}
	return c
}

const weightsTotal = 100

// Parse reads a catalogue and refuses one whose role weights do not sum to
// 100 or whose entries carry no verified source.
func Parse(data []byte) (Catalogue, error) {
	var c Catalogue
	if err := json.Unmarshal(data, &c); err != nil {
		return Catalogue{}, fmt.Errorf("consumables: %w", err)
	}
	for role, rc := range c.Roles {
		if sum := rc.weightSum(); sum != weightsTotal {
			return Catalogue{}, fmt.Errorf("consumables: %s weights sum to %d, want %d", role, sum, weightsTotal)
		}
		all := append(append(append(append([]Entry{}, rc.Flask...), rc.Food...), rc.WeaponEnchant...), rc.WorldBuffs...)
		all = append(all, rc.CombatPotion.Entries...)
		for _, e := range all {
			if e.SpellID <= 0 {
				return Catalogue{}, fmt.Errorf("consumables: %s: an entry has no positive spell_id", role)
			}
			if e.Verified == "" {
				return Catalogue{}, fmt.Errorf("consumables: %s: %q has no verified source", role, e.Name)
			}
		}
	}
	return c, nil
}

// Load returns the embedded catalogue.
func Load() Catalogue { return defaultCatalogue }

// For selects a role's catalogue by the same role strings
// summary.RosterRow.Role uses: "dps", "healer", "tank". ok is false for an
// unknown role.
func (c Catalogue) For(role string) (RoleCatalogue, bool) {
	rc, ok := c.Roles[role]
	return rc, ok
}
```

- [ ] **Step 10: Write `consumables/consumables_test.go`**

```go
// logs/engine/mechanics/consumables/consumables_test.go
package consumables

import "testing"

func TestLoadHasAllThreeRolesWithWeightsSummingToOneHundred(t *testing.T) {
	c := Load()
	for _, role := range []string{"dps", "healer", "tank"} {
		rc, ok := c.For(role)
		if !ok {
			t.Fatalf("no catalogue for role %q", role)
		}
		if rc.weightSum() != 100 {
			t.Errorf("%s weights sum to %d, want 100", role, rc.weightSum())
		}
	}
	if _, ok := c.For("not-a-role"); ok {
		t.Fatal("an unknown role must report false")
	}
}

func TestParseRejectsWeightsThatDoNotSumToOneHundred(t *testing.T) {
	bad := []byte(`{"roles": {"dps": {"weights": {"flask": 50}, "flask": [], "food": [], "weapon_enchant": [], "world_buffs": [], "combat_potion": {"max_uses": 1, "entries": []}}}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("weights that do not sum to 100 must be refused")
	}
}

func TestParseRejectsAnEntryWithNoVerifiedSource(t *testing.T) {
	bad := []byte(`{"roles": {"dps": {"weights": {"flask": 100}, "flask": [{"spell_id": 1, "name": "X"}], "food": [], "weapon_enchant": [], "world_buffs": [], "combat_potion": {"max_uses": 1, "entries": []}}}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("an entry with no verified source must be refused")
	}
}
```

- [ ] **Step 11: Run the consumables tests, confirm pass, then run all three packages together**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go vet ./engine/mechanics/... && GOTOOLCHAIN=go1.25.11 go test ./engine/mechanics/... -v
```
Expected: `PASS` across `mechanics`, `mechanics/weights`, `mechanics/utility`,
`mechanics/consumables`.

- [ ] **Step 12: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine
git add logs/engine/mechanics/weights/ logs/engine/mechanics/utility/ logs/engine/mechanics/consumables/
printf 'feat: add curated weights, utility and consumables data packages\n\nSpec §3.1-3.3: three new curated-data packages under logs/engine/mechanics,\nfollowing the existing package'"'"'s embed-parse-validate-panic-at-boot pattern.\nweights/roles.json is spec §1.4'"'"'s default weight table verbatim.\nutility/tables holds one verified per-spec utility file for all 27 specs in\ndata/curated/specs.json (every spell_id checked against\ndata/builds/1.60.1.69893/spells.json before being committed).\nconsumables/catalogue.json holds the by-role raid consumable catalogue,\nsame verification standard.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task3.txt
git commit -F .superpowers/commit-msg-task3.txt
```

---

## Task 4: `logs/engine/rating` core types, constants and the `combine` helper

This task does **not** write `Score` itself — that needs all six components, which land in
Tasks 5–7 — but everything `Score` (Task 8) will be built from: the exported types the API
lane compiles against verbatim, and the renormalise/weighted-mean/cap arithmetic (§1.1, §1.5)
as a standalone, independently-testable function.

**Files:**
- Create: `logs/engine/rating/rating.go`
- Create: `logs/engine/rating/combine.go`
- Create: `logs/engine/rating/combine_test.go`
- Create: `logs/engine/rating/rating_test.go` (package-shape smoke test only; the real
  behavioural tests land per-component in Tasks 5–7 and end-to-end in Task 8)

**Interfaces:**
- Produces: `Card`, `Component`, `Moment`, `Assignment`, `ModelInfo`, `CuratedTables`,
  `Bracket`, `PercentileSource` — **exactly** as spec §4.1 gives `Card`/`Component`/
  `Assignment`, and per Rulings R3/R9 for `PercentileSource`/component-name constants.
- Produces: `combine(components *[6]Component, w weights.RoleWeights,
  survivalDeathScoreZero, capEnabled bool, capThreshold float64) (overallUncapped, overall
  float64, capped bool, basis string)` — Task 8's `Score` calls this after building all six
  components.
- Produces: `rosterRow(fight summary.Summary, player string) (summary.RosterRow, bool)`,
  `specSlug(class, spec string) string`, `round2(x float64) float64` — every component task
  (5–7) uses these.
- Produces the constants: `ComponentNameOutput` / `Survival` / `Mechanics` / `Utility` /
  `Preparation` / `Activity`; `BasisPercentile`, `BasisAbsolute`; `RoleDPS`, `RoleHealer`,
  `RoleTank`; `BandFast`, `BandTypical`, `BandSlow`; `MinSample = 20`;
  `DefaultCapThreshold = 40.0`; `ReasonNoMechanicsTable`, `ReasonNoUtilityTable`,
  `ReasonNoConsumableCatalogue`, `ReasonWipe`, `ReasonNotEnoughSamples`.

- [ ] **Step 1: Write `rating.go` — types and constants**

```go
// logs/engine/rating/rating.go
// Package rating scores one player's performance for one closed fight
// against the six components docs/superpowers/specs/2026-09-21-performance-
// rating-design.md §1 defines: pure functions over an already-finished
// summary.Summary plus curated tables, with no network or database access,
// so a unit test can hand it a hand-built fixture and a fake percentile
// source. Sibling to logs/engine/summary and logs/engine/mechanics, never
// imported by either (§4.1's own ruling: a rating is a post-fight
// computation, not something summary.Accumulator.Snapshot needs to finish
// mid-fight).
package rating

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/weights"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The six component names, exactly as they appear in Component.Name and in
// the site's public API (spec §4.1, §5.2).
const (
	ComponentNameOutput      = "output"
	ComponentNameSurvival    = "survival"
	ComponentNameMechanics   = "mechanics"
	ComponentNameUtility     = "utility"
	ComponentNamePreparation = "preparation"
	ComponentNameActivity    = "activity"
)

// componentOrder is the fixed [6]Component slot order Card.Components uses,
// matching the table in spec §1.4 top to bottom.
var componentOrder = [6]string{
	ComponentNameOutput, ComponentNameSurvival, ComponentNameMechanics,
	ComponentNameUtility, ComponentNamePreparation, ComponentNameActivity,
}

// A component's basis (spec §1.2): which of the two rules produced its score.
const (
	BasisPercentile = "percentile"
	BasisAbsolute   = "absolute"
)

// The three roles summary.RosterRow.Role already uses (roster.go's own
// role() function), reused here rather than a second vocabulary.
const (
	RoleDPS    = "dps"
	RoleHealer = "healer"
	RoleTank   = "tank"
)

// The three kill-time bands, spec §1.2.
const (
	BandFast    = "fast"
	BandTypical = "typical"
	BandSlow    = "slow"
)

// MinSample is spec §1.2's RULING: a percentile computed from fewer than
// this many kills is a percentile of noise dressed as a number, so the
// absolute standard is used instead until a bracket crosses it.
const MinSample = 20

// DefaultCapThreshold is spec §1.5's site-wide default: the score an
// avoidable-death-triggered cap holds Overall at. A guild-adjustable
// setting (out of scope here, officer tooling); ModelInfo carries whatever
// value a caller wants, this is only the default a caller not overriding it
// should pass.
const DefaultCapThreshold = 40.0

// The fixed machine-readable exclusion reasons Component.Reason carries
// (spec §5.2's own JSON example shows "no_mechanics_table" as exactly this
// shape; the web lane maps these to §6.4's reader-facing copy — RULING R10,
// the spec's own "§7.4" cross-references are a drafting slip for "§6.4").
const (
	ReasonNoMechanicsTable      = "no_mechanics_table"
	ReasonNoUtilityTable        = "no_utility_table"
	ReasonNoConsumableCatalogue = "no_consumable_catalogue"
	ReasonWipe                  = "wipe"
	ReasonNotEnoughSamples      = "not_enough_samples"
)

// Card is one player's rating for one fight (spec §4.1, verbatim).
type Card struct {
	Overall         float64
	OverallUncapped float64
	OverallCapped   bool
	Components      [6]Component
	Basis           string
	ModelVersion    string
	KillTimeBand    string
}

// Component is one of the six parts of a Card (spec §4.1, verbatim).
type Component struct {
	Name       string
	Score      float64
	Weight     float64
	Basis      string
	Percentile *float64
	BracketN   int64
	Excluded   bool
	Reason     string
	Moments    []Moment
}

// Moment is one specific instant a component's score is built from — a
// death, an avoidable hit, an interrupt made, a dispel made, a dropped
// debuff — the "opens into the moments in the log" data spec §6.1/§7.1
// describe. Anchor, when set, is a DOM anchor id an existing tab already
// renders (spec §6.1's "Jumping to a moment"): "death-<guid>-<at_ms>" for a
// death, matching DeathsTab.svelte's own convention.
type Moment struct {
	Kind      string `json:"kind"`
	AtMS      int64  `json:"at_ms"`
	SpellID   int64  `json:"spell_id,omitempty"`
	SpellName string `json:"spell_name,omitempty"`
	Avoidable bool   `json:"avoidable,omitempty"`
	Anchor    string `json:"anchor,omitempty"`
}

// Assignment is an officer-marked job for one player on one fight (spec §2,
// verbatim). Score's assignments parameter is always an empty slice today —
// no caller populates it — but the type and Score's signature accept it now
// so officer tooling can be built against a stable contract later without a
// model change.
type Assignment struct {
	PlayerKey string
	Job       string
	FromMS    int64
	ToMS      int64
}

// ModelInfo is every guild-adjustable setting Score needs, bundled so
// Score's signature does not grow every time officer tooling adds one
// (spec §4.1).
type ModelInfo struct {
	ModelVersion string
	Weights      weights.Roles
	// CapEnabled and CapThreshold are §1.5's cap: a guild may turn it off
	// entirely, or (in principle) recompute its threshold. A caller wanting
	// the site default passes CapEnabled: true, CapThreshold:
	// DefaultCapThreshold.
	CapEnabled   bool
	CapThreshold float64
}

// CuratedTables bundles the curated data a fight's rating reads, already
// selected for this player's encounter/spec/role by the caller: Mechanics
// is the fight's own encounter table (nil if none curated yet), Utility is
// this player's spec's owned-utility table (nil if none curated for this
// spec), Consumables is this player's role's consumable catalogue row (nil
// only if the embedded catalogue is somehow missing a role, which
// mechanics/consumables.Parse already refuses at boot).
type CuratedTables struct {
	Mechanics   *mechanics.Table
	Utility     *utility.Table
	Consumables *consumables.RoleCatalogue
}

// Bracket is the percentile-digest key spec §1.2 defines:
// (encounter_id, difficulty, spec, role, kill_time_band, component), plus
// Kill (spec §2's wipe rule: a wipe's digests are kept apart from kills').
// Component is one of the eight strings this package queries with —
// ComponentNameOutput/Utility/Preparation/Activity for the four
// single-query top-level components, plus the four sub-part names
// mechanics_interrupt/mechanics_dispel/mechanics_debuff_uptime and
// survival_avoidable_hit for Mechanics' and Survival's independently-
// percentiled parts (RULING R9 — these are not top-level Component.Name
// values, only digest keys).
type Bracket struct {
	EncounterID  int64
	Difficulty   int64
	Spec         string
	Role         string
	KillTimeBand string
	Kill         bool
	Component    string
}

// The Bracket.Component values for Mechanics' and Survival's sub-parts,
// which each need their own percentile digest distinct from any of the six
// top-level component names (RULING R9).
const (
	ComponentSurvivalAvoidableHit  = "survival_avoidable_hit"
	ComponentMechanicsInterrupt    = "mechanics_interrupt"
	ComponentMechanicsDispel       = "mechanics_dispel"
	ComponentMechanicsDebuffUptime = "mechanics_debuff_uptime"
)

// PercentileSource is the small interface Score reads distributions
// through (spec §4.1), so logs/engine never depends on api/internal: the
// API layer wires a concrete implementation backed by
// api/internal/digest.Unmarshal/Placement over rows read from
// rating_percentile_digests.
type PercentileSource interface {
	// Placement is spec §1.2's percentile-within-bracket lookup: the share
	// of the bracket's other values this value beats, in 0..1, and how many
	// kills the bracket has. ok is false when the source has nothing for
	// this bracket at all (a brand-new encounter/spec/role combination).
	Placement(bracket Bracket, value float64) (pct float64, n int64, ok bool)
	// KillTimeBand classifies durationMS against the (encounterID,
	// difficulty) bracket's own trailing kill-duration median (spec §1.2).
	// ok is false when the source has no kill-duration digest for this
	// bracket yet, in which case the caller uses BandTypical — spec §1.2's
	// own note that an unclassifiable band "degenerates to a no-op, not a
	// wrong answer" (RULING R3: added to the interface spec §4.1 names only
	// Placement on, since Placement alone cannot answer a median query).
	KillTimeBand(encounterID, difficulty int64, durationMS int64) (band string, ok bool)
}

// rosterRow finds player's row in the fight's roster.
func rosterRow(fight summary.Summary, player string) (summary.RosterRow, bool) {
	for _, r := range fight.Roster {
		if r.GUID == player {
			return r, true
		}
	}
	return summary.RosterRow{}, false
}

// round2 rounds to two decimal places: spec §1.5's storage precision ("the
// stored value keeps two decimal places... so a recompute is stable to the
// same input").
func round2(x float64) float64 {
	return float64(int64(x*100+sign(x)*0.5)) / 100
}

func sign(x float64) float64 {
	if x < 0 {
		return -1
	}
	return 1
}
```

- [ ] **Step 2: Write `rating_test.go` — the package-shape smoke test**

```go
// logs/engine/rating/rating_test.go
package rating

import "testing"

func TestRosterRowFindsAPlayerAndSaysWhenThereIsNone(t *testing.T) {
	fight := fixtureSummary(t) // defined in combine_test.go's shared helpers, see Task 5
	if _, ok := rosterRow(fight, "does-not-exist"); ok {
		t.Fatal("an unknown guid must report false")
	}
}

func TestRound2RoundsToTwoDecimalPlaces(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{70.004, 70.0}, {70.005, 70.01}, {-1.005, -1.01}, {0, 0}, {99.995, 100.0},
	}
	for _, c := range cases {
		if got := round2(c.in); got != c.want {
			t.Errorf("round2(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
```

This references `fixtureSummary`, which Step 3 below defines in this same task (shared by
every later task's tests) so it is available from the start.

- [ ] **Step 3: Write `combine.go`**

```go
// logs/engine/rating/combine.go
package rating

import "github.com/jhunthrop/foreversixty/logs/engine/mechanics/weights"

// weightOf returns a component's site-default weight for one role's weight
// row, by name.
func weightOf(w weights.RoleWeights, name string) float64 {
	switch name {
	case ComponentNameOutput:
		return w.Output
	case ComponentNameSurvival:
		return w.Survival
	case ComponentNameMechanics:
		return w.Mechanics
	case ComponentNameUtility:
		return w.Utility
	case ComponentNamePreparation:
		return w.Preparation
	default:
		return w.Activity
	}
}

// combine applies spec §1.1's renormalisation and §1.5's weighted mean and
// Survival-catastrophe cap to six already-scored components, in place:
// every non-excluded Component's Weight field is overwritten with its
// renormalised share (spec §4.1's "the *renormalised* weight actually
// applied"). survivalDeathScoreZero is true only when Survival was scored
// (not excluded) and its death half alone floored to zero (spec §1.5's cap
// condition) — Component carries no field for "the death half
// specifically," so Task 7's Survival scorer threads it through here
// separately rather than it being re-derived from Component.Score, which
// mixes the death and avoidable-hit halves together.
func combine(
	components *[6]Component, w weights.RoleWeights,
	survivalDeathScoreZero, capEnabled bool, capThreshold float64,
) (overallUncapped, overall float64, capped bool, basis string) {
	var totalWeight float64
	for _, c := range components {
		if !c.Excluded {
			totalWeight += weightOf(w, c.Name)
		}
	}

	basisKinds := map[string]bool{}
	var sum float64
	for i := range components {
		c := &components[i]
		if c.Excluded {
			c.Weight = 0
			continue
		}
		if totalWeight <= 0 {
			c.Weight = 0
			continue
		}
		share := weightOf(w, c.Name) / totalWeight
		c.Weight = round2(share * 100)
		sum += share * c.Score
		if c.Basis != "" {
			basisKinds[c.Basis] = true
		}
	}

	overallUncapped = round2(sum)
	overall = overallUncapped
	capped = survivalDeathScoreZero
	if capped && capEnabled && overallUncapped > capThreshold {
		overall = capThreshold
	}

	switch len(basisKinds) {
	case 1:
		for k := range basisKinds {
			basis = k
		}
	case 0:
		basis = ""
	default:
		basis = "mixed"
	}
	return overallUncapped, overall, capped, basis
}
```

- [ ] **Step 4: Write `combine_test.go`, including the shared `fixtureSummary` test helper**

```go
// logs/engine/rating/combine_test.go
package rating

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/weights"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// fixtureSummary is the shared minimal fixture every test file in this
// package builds on and mutates via its own helpers — a one-player, one-
// death, 180-second fight on Shazzrah (encounter_id 667), matching spec
// §1.6's worked examples' own setup (a 3-minute Molten Core kill).
func fixtureSummary(t *testing.T) summary.Summary {
	t.Helper()
	const guid = "Player-1-00000001"
	return summary.Summary{
		EngineVersion: "test", DurationMS: 180000, EncounterID: 667, Difficulty: 0, Kill: true,
		Roster: []summary.RosterRow{
			{GUID: guid, Name: "Fixture", Class: "Warrior", Spec: "Fury", Role: RoleDPS, DPS: 500, ActiveMS: 163800, ActivityPct: 91},
		},
	}
}

// fakePercentiles is a hand-fed PercentileSource: every component test
// supplies exactly the (bracket.Component, value) -> (pct, n) pairs its
// scenario needs, so the golden-fixture-style determinism of §1.2's rule is
// exercised without any database.
type fakePercentiles struct {
	placements map[string]fakePlacement
	band       string
	bandOK     bool
}

type fakePlacement struct {
	pct float64
	n   int64
	ok  bool
}

func (f fakePercentiles) Placement(bracket Bracket, value float64) (float64, int64, bool) {
	p, ok := f.placements[bracket.Component]
	if !ok {
		return 0, 0, false
	}
	return p.pct, p.n, p.ok
}

func (f fakePercentiles) KillTimeBand(encounterID, difficulty int64, durationMS int64) (string, bool) {
	return f.band, f.bandOK
}

func TestCombineRenormalisesWeightsAcrossExcludedComponents(t *testing.T) {
	w := weights.Default().Tank // Output 10, Survival 35, Mechanics 20, Utility 20, Preparation 10, Activity 5
	components := [6]Component{
		{Name: ComponentNameOutput, Score: 55},
		{Name: ComponentNameSurvival, Score: 76.06},
		{Name: ComponentNameMechanics, Excluded: true, Reason: ReasonNoMechanicsTable},
		{Name: ComponentNameUtility, Score: 91},
		{Name: ComponentNamePreparation, Score: 100},
		{Name: ComponentNameActivity, Score: 70},
	}
	overallUncapped, overall, capped, _ := combine(&components, w, false, true, DefaultCapThreshold)
	if capped {
		t.Fatal("no cap condition was signalled; capped must be false")
	}
	// Mechanics' 20 points redistribute across the other five in proportion
	// to their own weights (10, 35, 15... wait Utility/Preparation/Activity
	// keep their own weights too): matches spec §1.6's tank worked example.
	if want := 75.1; roundTo(overallUncapped, 1) != want {
		t.Fatalf("overallUncapped = %v, want %v (spec §1.6's tank worked example)", overallUncapped, want)
	}
	if overall != overallUncapped {
		t.Fatalf("overall = %v, want %v (no cap fired)", overall, overallUncapped)
	}
	if components[2].Weight != 0 {
		t.Errorf("an excluded component's Weight must be 0, got %v", components[2].Weight)
	}
	var sumWeights float64
	for _, c := range components {
		sumWeights += c.Weight
	}
	if roundTo(sumWeights, 1) != 100 {
		t.Errorf("renormalised weights sum to %v, want 100", sumWeights)
	}
}

func TestCombineAppliesTheCatastropheCapWhenEnabled(t *testing.T) {
	w := weights.Default().DPS
	components := [6]Component{
		{Name: ComponentNameOutput, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNameSurvival, Score: 40, Basis: BasisAbsolute},
		{Name: ComponentNameMechanics, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNameUtility, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNamePreparation, Score: 90, Basis: BasisPercentile},
		{Name: ComponentNameActivity, Score: 90, Basis: BasisPercentile},
	}
	overallUncapped, overall, capped, basis := combine(&components, w, true, true, DefaultCapThreshold)
	if !capped {
		t.Fatal("survivalDeathScoreZero was true; capped must be true")
	}
	if overallUncapped <= DefaultCapThreshold {
		t.Fatalf("overallUncapped = %v, fixture is built to exceed the cap", overallUncapped)
	}
	if overall != DefaultCapThreshold {
		t.Fatalf("overall = %v, want the cap threshold %v", overall, DefaultCapThreshold)
	}
	if basis != "mixed" {
		t.Fatalf("basis = %q, want mixed (percentile and absolute both present)", basis)
	}
}

func TestCombineLeavesOverallUncappedWhenCapDisabled(t *testing.T) {
	w := weights.Default().DPS
	components := [6]Component{
		{Name: ComponentNameOutput, Score: 90}, {Name: ComponentNameSurvival, Score: 40},
		{Name: ComponentNameMechanics, Score: 90}, {Name: ComponentNameUtility, Score: 90},
		{Name: ComponentNamePreparation, Score: 90}, {Name: ComponentNameActivity, Score: 90},
	}
	overallUncapped, overall, capped, _ := combine(&components, w, true, false, DefaultCapThreshold)
	if !capped {
		t.Fatal("the condition still fired even though the cap is disabled")
	}
	if overall != overallUncapped {
		t.Fatalf("overall = %v, want it to equal overallUncapped (%v) when the cap is disabled", overall, overallUncapped)
	}
}

// roundTo rounds to n decimal places, for comparing arithmetic against the
// spec's own worked-example precision without a flaky float equality check.
func roundTo(x float64, n int) float64 {
	m := 1.0
	for i := 0; i < n; i++ {
		m *= 10
	}
	return float64(int64(x*m+0.5)) / m
}
```

- [ ] **Step 5: Run the tests, confirm they fail to compile (package doesn't exist yet), then implement, then pass**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go test ./engine/rating/... -v
```
Write `rating.go` and `combine.go` from Steps 1 and 3 above (if not already written), then
re-run. Expected: all four tests `PASS`, in particular
`TestCombineRenormalisesWeightsAcrossExcludedComponents` reproducing 75.1 — spec §1.6's tank
worked example's own overall figure — from `combine` alone, before any component's real
formula exists.

- [ ] **Step 6: `go vet` and commit**

```bash
cd /Users/jh/code/forever/.worktrees/rating-engine/logs && GOTOOLCHAIN=go1.25.11 go vet ./engine/rating/...
cd /Users/jh/code/forever/.worktrees/rating-engine
git add logs/engine/rating/
printf 'feat: add logs/engine/rating package skeleton (types, combine)\n\nSpec §4.1: Card/Component/Assignment exactly as given, plus ModelInfo,\nCuratedTables, Bracket and PercentileSource (extended with KillTimeBand per\nruling R3 in docs/superpowers/plans/2026-09-21-rating-engine.md, since\nPlacement alone cannot answer a median-duration query). combine() implements\nspec §1.1'"'"'s renormalisation and §1.5'"'"'s weighted mean and catastrophe cap,\nverified standalone against the tank worked example'"'"'s own arithmetic (75.1)\nbefore any component formula exists. Score() itself lands once all six\ncomponents do (Tasks 5-8).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-task4.txt
git commit -F .superpowers/commit-msg-task4.txt
```

---
