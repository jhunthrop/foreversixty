# Performance rating engine — implementation plan

> **For agentic workers:** this plan was executed directly by the lane
> controller (not dispatched task-by-task to fresh subagents) because the
> six components share constants, helper functions and a single
> `componentCtx`, and re-deriving the spec's several genuinely ambiguous
> formulas independently per subagent risked inconsistency. Subagents were
> used for the pieces that *are* independent: curated-data authoring
> (utility tables, consumable catalogue) and the `mechanics-draft` CLI
> extension. See the progress ledger for every ruling made along the way.

**Goal:** Implement the ENGINE lane of
`docs/superpowers/specs/2026-09-21-performance-rating-design.md`: the new
`logs/engine/rating` package (pure functions scoring one player's fight,
§4.1), the curated-data schemas and files it reads (§3), and the
`downtime` extension to the existing mechanics table (§3.4), with golden
and worked-example tests (§8, §1.6).

**Architecture:** `logs/engine/rating` is a new, sibling package to
`summary` and `mechanics` (never imported by either). It takes an
already-closed `summary.Summary`, a bundle of curated tables
(`CuratedTables`), a `PercentileSource` interface the caller wires against
`api/internal/digest` (never imported directly — keeps `logs/engine` free
of `api/internal`), an `[]Assignment` (always empty today), and a
`ModelInfo` (weights + cap + version), and returns a `Card`. Every
component (`output`, `survival`, `mechanics`, `utility`, `preparation`,
`activity`) is its own file with its own pure scoring function, called
from `score.go`'s orchestrator, which renormalises weights over
non-excluded components and applies §1.5's cap.

**Tech Stack:** Go 1.25 (`GOTOOLCHAIN=go1.25.11` on this machine),
`encoding/json` + `embed` for curated data (matching the existing
`mechanics` package pattern), the existing `logs/engine/summary` golden
test harness reused for regression tests.

**Spec:** `docs/superpowers/specs/2026-09-21-performance-rating-design.md`
(sections 1, 2, 3, 4.1 are this lane's scope; §8's testing lane split and
worked examples in §1.6 gate what "done" means).

## Global Constraints

- Own only: `logs/engine/rating/**` (new), `logs/engine/mechanics/**`
  (downtime schema extension + new `utility/`, `consumables/`, `weights/`
  subpackages), the curated data files §3.1 names, `logs/cmd/forever-logs`
  only for the §3.5 drafting-tool extension, and tests for all of the
  above. Never touch `api/`, `web/`, `sim/`, `addon/`, `companion/`, or any
  other existing engine package beyond additive fields this spec requires.
- `logs/engine` must never import `api/internal` — percentile data arrives
  only through the injected `PercentileSource` interface.
- A curated spell id must be verified against
  `data/builds/1.60.1.69893/spells.json`; an unverifiable id is left out
  with a comment, never guessed.
- Go: `go vet ./... && go test ./...` for touched packages before every
  commit; table-driven tests; functions <50 lines, files <800 lines, no
  magic numbers, immutable updates, wrapped errors.
- The worked examples in spec §1.6 (DPS, healer, tank) must exist as tests
  reproducing the spec's own numbers; an arithmetic error in the spec is
  fixed in the spec in the same commit, noted in the report.
- Reuse the existing golden-fixture harness (`logs/engine/summary/golden_test.go`
  pattern: table of cases, `FOREVER_UPDATE_GOLDEN` env var, JSON
  comparison) rather than inventing a second one.

---

## File map

```
logs/engine/summary/summary.go     modify: += Summary.EncounterID, Difficulty, Kill (additive, omitempty where zero-valued makes sense; Kill has no omitempty since false is meaningful)
logs/engine/summary/roster.go      modify: += RosterRow.ExecutionScore *float64 `json:"execution_score,omitempty"`
logs/engine/summary/testdata/*.golden   regenerated (new fields only)

logs/engine/mechanics/mechanics.go modify: += Table.Downtime []Downtime, Downtime type, PhaseStart reused as Trigger, validation
logs/engine/mechanics/tables/666.json   modify: += a real, verified downtime entry for Garr

logs/engine/mechanics/weights/roles.go       new: RoleWeights, Roles types, embed roles.json, Load()
logs/engine/mechanics/weights/roles.json     new: the default §1.4 weight table as data
logs/engine/mechanics/weights/roles_test.go  new

logs/engine/mechanics/utility/utility.go        new: Entry, Table types, embed tables/*.json, Load(spec)
logs/engine/mechanics/utility/tables/*.json     new: 27 files (curated by subagent, verified)
logs/engine/mechanics/utility/utility_test.go   new

logs/engine/mechanics/consumables/consumables.go      new: Entry, PotionGroup, RoleCatalogue, Catalogue types, embed catalogue.json, Load()
logs/engine/mechanics/consumables/catalogue.json       new: curated by subagent, verified
logs/engine/mechanics/consumables/consumables_test.go  new

logs/engine/rating/types.go           new: Card, Component, Moment, Assignment, ModelInfo, CuratedTables, component name + basis + reason constants
logs/engine/rating/bracket.go         new: Bracket, PercentileSource, MinSample, killTimeBand, evaluate()
logs/engine/rating/specs.go           new: specSlug(class, spec string) — copy of the class+spec→slug table logs/engine cannot reach data/curated/specs.json for
logs/engine/rating/output.go          new: scoreOutput
logs/engine/rating/survival.go        new: scoreSurvival, deathScoreFor, avoidableDamagePerSecond
logs/engine/rating/mechanics_component.go  new: scoreMechanics + 3 parts
logs/engine/rating/utility_component.go    new: scoreUtility
logs/engine/rating/preparation.go     new: scorePreparation
logs/engine/rating/activity.go        new: scoreActivity, downtime window math
logs/engine/rating/score.go           new: Score() orchestrator — renormalisation, cap, rounding
logs/engine/rating/*_test.go          new: unit tests per file
logs/engine/rating/golden_test.go     new: reuses summary's v16/v22 golden Summary fixtures
logs/engine/rating/worked_examples_test.go  new: §1.6's three tables, exact numbers
logs/engine/rating/testdata/*.golden  new

logs/cmd/forever-logs/mechanics_draft.go  modify: += utility evidence pass, downtime evidence pass (§3.5)
```

---

## Task list (executed directly; TDD per task, `go vet && go test` before each commit)

### Task 1 — Additive Summary/RosterRow fields the fixed `rating.Score` signature needs

`rating.Score(fight summary.Summary, player string, ...)` takes only a
`summary.Summary`, never `fight.Fight` — but §1.2's bracket key needs
`EncounterID`/`Difficulty`, §2's wipe rule needs `Kill`, and §1.3's Output
needs `execution_score` (computed by the API layer, stored outside
`summary.Summary` entirely, in `fight_metrics`). **Ruling:** add
`Summary.EncounterID int64`, `Summary.Difficulty int64`, `Summary.Kill
bool` (populated from `fight.Fight` in `Snapshot`, mirroring how
`FightIndex`/`DurationMS` already are), and `RosterRow.ExecutionScore
*float64 \`json:"execution_score,omitempty"\`` (left nil by the log
engine; the API layer overlays each player's already-computed
`fight_metrics.execution_score` onto the roster row it rebuilds before
calling `rating.Score`, per §4.4's "re-reads each fight's stored summary"
flow). Both are additive/omitempty-safe; regenerate the two existing
golden fixtures and confirm the diff is exactly these new fields.

- Write test asserting `Snapshot` copies `f.EncounterID`/`f.Difficulty`/`f.Kill` onto the returned `Summary`.
- Add the fields, wire them in `Snapshot`.
- `FOREVER_UPDATE_GOLDEN=1 go test ./logs/engine/summary/... -run TestTheFixtureSummaryMatchesTheCommittedGolden`, review the diff (must be additive only), commit.

### Task 2 — `Downtime` schema extension on `mechanics.Table`

Add, in `logs/engine/mechanics/mechanics.go`:

```go
// Downtime is a forced stretch a listed trigger opens, during which the
// boss is untargetable or the raid is expected to step out (§3.4). It
// extends Activity's denominator only — Survival and Mechanics still
// count a hit landed during one, since "untargetable" does not mean
// "avoidable damage stopped mattering."
type Downtime struct {
	Trigger    PhaseStart `json:"trigger"`
	DurationMS int64      `json:"duration_ms"`
	Note       string     `json:"note,omitempty"`
}
```

`Table.Downtime []Downtime \`json:"downtime,omitempty"\`` — omitted means
none, same convention `Phases` already uses. Validate each entry's
`Trigger` with the exact same switch `Parse` already runs for a `Phase`'s
`Starts` (extract that validation into a shared `validatePhaseStart`
helper both call, DRY). Add a verified downtime entry to `tables/666.json`
(Garr): `spell_id: 19497` ("Eruption," already a verified id in that same
file) with `on: "aura_applied"`, since the Firesworn's death detonation
is exactly this pattern and 19497 is already curated, verified data in
this repo. Add `TestParseAcceptsDowntimeAndRefusesAMalformedTrigger` and
`TestEveryEmbeddedTableParses` continues to cover 666.json unchanged
(no new required fields on existing tables).

### Task 3 — `logs/engine/mechanics/weights` package

`RoleWeights{Output, Survival, Mechanics, Utility, Preparation, Activity float64}`,
`Roles{DPS, Healer, Tank RoleWeights}`. `roles.json` holds §1.4's table
verbatim (35/15/20/15/10/5 dps, 30/10/20/20/10/10 healer,
10/35/20/20/10/5 tank). `Parse` validates each role's six weights sum to
100. Embed + `Load()` (panic-at-init like `mechanics.mustParseAll`).
`(Roles) For(role string) RoleWeights` selects by `"dps"|"healer"|"tank"`,
defaulting to DPS weights with a `bool` "found" return for an unknown
role string (defensive; never expected in practice since `RosterRow.Role`
only ever produces those three values, per `roster.go`'s `role()`).

### Task 4 — `logs/engine/mechanics/utility` package

`Entry{SpellID int64, Name, Kind, Target, Verified, Note string}`,
`Table{Spec string, Owned []Entry}`. `Kind` one of `buff|debuff|cooldown|
talent-modifier` (constants, validated at parse — same `kinds` map pattern
`mechanics.go` uses). Embed `tables/*.json`, `Load(specSlug string) (Table, bool)`.
The 27 files are curated separately (parallel data-authoring subagent);
this task's test (`TestEveryEmbeddedUtilityTableParses`) walks the embed
FS the same way `mechanics.TestEveryEmbeddedTableParses` does, and a
`TestEveryCuratedSpecHasAUtilityFile` cross-checks file names against
`data/curated/specs.json`'s 27 `spec` values (plain `os.ReadFile` at test
time, not `//go:embed` — a test may reach outside its module the way an
embed cannot).

### Task 5 — `logs/engine/mechanics/consumables` package

`Entry{SpellID int64, Name, Verified string}`, `PotionGroup{MaxUses int,
Entries []Entry}`, `RoleWeights{Flask, Food, WeaponEnchant, WorldBuffs,
CombatPotion float64}`, `RoleCatalogue{Weights RoleWeights, Flask, Food,
WeaponEnchant, WorldBuffs []Entry, CombatPotion PotionGroup}`,
`Catalogue{Roles map[string]RoleCatalogue}`. Validate each role's five
weights sum to 100 (mirrors `mechanics.Parse`'s field checks). Embed
`catalogue.json` (curated by the parallel subagent), `Load() Catalogue`.

### Task 6 — `logs/engine/rating` core types (`types.go`, `bracket.go`, `specs.go`)

`Card`, `Component`, `Moment`, `Assignment`, `ModelInfo`, `CuratedTables`
exactly as spec §4.1 gives their fields (see spec text — transcribed
verbatim, JSON tags snake_case matching §5.2's example payload).
`Bracket{EncounterID, Difficulty int64; Spec, Role, KillTimeBand string;
Kill bool; Component string}`. `PercentileSource` interface:

```go
type PercentileSource interface {
	Placement(bracket Bracket, value float64) (pct float64, n int64, ok bool)
	// RULING: added beyond the spec's own Placement-only sketch (§4.1) —
	// §1.2's kill_time_band needs a bracket's median kill duration, which
	// Placement (a percentile of a *given* value) cannot answer, and
	// Score's signature is fixed (§4.1) with no room for a second
	// parameter. Smallest necessary extension to the one interface the
	// spec already names.
	KillTimeBandMedian(encounterID, difficulty int64) (medianMS float64, ok bool)
}
```

`MinSample = 20`. `evaluate(ps, bracket, value *float64, absolute func() (float64, bool)) (score float64, basis string, pct *float64, n int64, excluded bool)` implements §1.2's exact rule. `killTimeBand(durationMS int64, medianMS float64) string` implements §1.2's threshold formula as a pure function (unit-tested directly against the three band boundaries). `specSlug(class, spec string) (string, bool)` is a literal Go map transcribed from `data/curated/specs.json` (27 entries) — **ruling:** duplicated here, not imported, because `logs/engine` cannot reach `api/internal`'s `specs` package or `//go:embed` a file outside its own module tree; a test (`TestSpecSlugsMatchCuratedSpecsJSON`) reads `../../../data/curated/specs.json` via plain `os.ReadFile` (not embed) at test time and fails if the two drift.

### Task 7 — `output.go`

`scoreOutput(c componentCtx) Component` per §1.3's Output formula: wipe →
excluded (`wipe`); `RosterRow.ExecutionScore != nil` → percentile-or-
absolute (`min(100, execScore*100)`) via `evaluate`; else raw `DPS`/`HPS`
fallback (healer role → `HPS`, else `DPS`) with no absolute standard.

### Task 8 — `survival.go`

`deathScoreFor` implements the exact penalty formula (`70 * causeMultiplier *
timeRemaining`, multipliers 1.00/0.50/0.15, floored at 0) and builds
`death` Moments with `anchor = "death-<guid>-<at_ms>"` matching
`DeathsTab.svelte`'s existing anchor convention (§6.1). `avoidableDamagePerSecond`
sums `MechanicsBlock.Rows[kind=avoidable].Players[guid=player].Damage`,
excluding rows whose `Role` equals the player's own role, divided by
fight-seconds. `scoreSurvival` returns `(Component, deathScore float64,
scored bool)` — the extra two values feed §1.5's cap, computed even when
the *displayed* component ends up excluded for insufficient percentile
sample (the cap only cares whether the player's own deaths already
floored DeathScore, not whether the sub-percentile could also be ranked).
No mechanics table → whole component excluded (`no_mechanics_table`),
`scored=false`.

### Task 9 — `mechanics_component.go`

Three independently-excludable parts, each its own percentile bracket
(`Component` key `"mechanics_interrupt"`, `"mechanics_dispel"`,
`"mechanics_debuff_uptime"` — **ruling:** the spec's schema names one
`component` column but Mechanics is the only top-level component built
from three differently-scaled raw values, so the three parts need three
digests; recorded here for the API lane, which owns the
`rating_percentile_digests` writer). Interrupts: `Σ (Damage/max(1,Casts)) *
x.Count` over this player's `ExchangeRow{Kind:"interrupt"}` rows, divided
by time-alive-seconds. Dispels: same shape, `(Healed+Damage)/max(1,Dispelled)
* x.Count`. Own dispellable-debuff uptime: `Σ UptimeMS` over this player's
`DEBUFF` `AuraTrack`s whose spell classifies `dispel`, percentile inverted
(`100 - pct`). Average with equal weight over the parts present; `Basis`
is `"mixed"` when more than one part is present, else that one part's own
basis. No mechanics table → whole component excluded; a table with none
of the three parts computable → excluded (`no_mechanics_rows`).

### Task 10 — `utility_component.go`

For every `owned` entry not `talent-modifier`: `buff`/`debuff` → mean
uptime share (`UptimeMS/DurationMS`) across every `AuraTrack` of that
spell id the player is an applier of (a raid-wide buff can have many
target tracks); `cooldown` → used-at-least-once-in-the-fight as a binary
1.0/0.0 share (**ruling:** the schema has no `expected_uses` field to
divide a cast count against "fight length" the way §3.2's prose gestures
at; binary use/no-use is the simplest honest signal available today,
flagged for a follow-up schema field if this proves too coarse). Mean the
per-entry shares, percentile-within-bracket. Append a `threat_not_modeled`
Moment (never excluding the component) when `Summary.Threat`'s
`Complete` flag is false, per §1.3's Threat note. No utility table or
empty `Owned` → excluded (`no_utility_table`).

### Task 11 — `preparation.go`

Per-category presence check (flask/food/weapon_enchant/world_buffs)
against `RosterRow.Consumables`, crediting the category's full weight if
any one catalogue entry in that category is present; combat potion
credited proportionally, `uses/max_uses` capped at 1.0, `uses` = sum of
`CastRow.Succeeded` over the catalogue's potion spell ids for this player.
Sum of credited weights (0-100) is the raw value, percentile-within-
bracket, no absolute standard. No catalogue or no entry for this role →
excluded (`no_consumable_catalogue`).

### Task 12 — `activity.go`

`timeAliveMS` = this player's own death `AtMS`, or `DurationMS` if no
death, clamped to `DurationMS`. `downtimeOverlapMS` walks
`Table.Downtime`, finds every trigger instant via
`downtimeTriggerInstants` (reading `Summary.Casts[].Sequence` for
`cast_success`, `Summary.Auras[].Segments` start/end for
`aura_applied`/`aura_removed` — **ruling:** `cast_start` and `health_pct`
triggers are not derivable from a finished `Summary` (`Casts.Sequence`
only records successful-cast instants; no boss-health timeline survives
past `Snapshot`) — a downtime entry using either contributes zero windows,
a documented gap rather than a silent wrong answer, matching the spirit
of every other honestly-flagged gap in this spec (Threat, tank execution
score)), merges overlapping windows, clips to `[0, timeAliveMS)`. No
curated downtime → `approximate=true`, falls back to whole-fight
`RosterRow.ActivityPct/100`, **never excluded** for this specific reason
(§1.1). With curated downtime, `activeMSExcludingDowntime` approximates a
downtime-adjusted active time from the player's own per-second
`DamageDone`/`Healing` `Series` (one-second buckets, the same events
`markActive` folds, at the granularity `Summary` actually exposes): a
bucket with nonzero output inside a downtime window does not count as
active *or* denominator time. `activeShare = activeMS /
(timeAliveMS-downtimeMS)`, `value = 100*activeShare`,
percentile-within-bracket.

### Task 13 — `score.go` orchestrator

```go
func Score(fight summary.Summary, player string, tables CuratedTables,
	percentiles PercentileSource, assignments []Assignment, now ModelInfo) Card
```

Builds `componentCtx` (roster row lookup, role, spec slug via `specSlug`,
kill-time band via `classifyKillTimeBand`), calls all six `score*`
functions, computes renormalised weights (§1.1: excluded components' raw
weight redistributed proportionally across scored ones — `w'_c = w_c *
(totalWeight / Σ w_scored)`), computes `OverallUncapped = Σ (w'_c/100) *
score_c`, applies §1.5's cap (`OverallCapped = deathScoreZero &&
now.CapEnabled`, `Overall = min(OverallUncapped, now.CapThreshold)` when
capped else `OverallUncapped`), rounds `Overall`/`OverallUncapped`/every
`Component.Score` to the nearest integer for the `Card`'s numbers **only
where the spec's own worked examples round** — the stored/tested values
keep 2 decimals per §1.5's "the stored value keeps two decimal places," so
`Card` itself stores the *unrounded* float64s (the API lane's storage
layer / web layer does display rounding) — this task's tests assert
against the spec's rounded display numbers by rounding the returned
`Card` values before comparing, not by asserting the `Card` fields are
pre-rounded integers.

### Task 14 — Golden fixture tests

`logs/engine/rating/golden_test.go` loads
`../summary/testdata/v16.summary.json.golden` and
`v22.summary.json.golden` (`os.ReadFile` + `json.Unmarshal` into
`summary.Summary` — no log parsing needed, per §8's own note that these
files are "exactly the shape `rating.Score` consumes"), builds a minimal
hand-written `CuratedTables`/fake `PercentileSource` per fixture, calls
`Score` for every roster player, and compares the resulting `[]Card`
(sorted by player GUID for determinism) against a new
`testdata/v16.rating.json.golden` / `v22.rating.json.golden`, following
the exact same `FOREVER_UPDATE_GOLDEN` env var and stability-across-runs
pattern as `logs/engine/summary/golden_test.go`.

### Task 15 — Worked-example tests (§1.6)

`worked_examples_test.go`: three tests (`TestWorkedExampleDPSWarriorFury`,
`TestWorkedExampleHealerPriestHoly`, `TestWorkedExampleTankWarriorProtection`),
each hand-building a minimal `summary.Summary` (one death at the stated
`AtMS` with the stated `KillingBlow.SpellID`, `DurationMS: 180000`,
`Roster` row with the stated role, `ExecutionScore`/`metric_dps`/`metric_hps`
as the example states) plus a fake `PercentileSource` whose `Placement`
returns exactly the percentile the table states for each raw value, and
asserting `Score(...)`'s `Overall` (rounded) equals the worked example's
final number. **If arithmetic in the spec doesn't reproduce from its own
stated inputs, fix the spec's number in the same commit** (see the tank
example's Mechanics-excluded renormalisation, which the plan verified by
hand: weights `10,35,15,10,5` (Output/Survival/Utility/Preparation/
Activity) sum to 75; `+20*(own/75)` gives Output `+2.67`, Survival
`+9.33`, Utility `+4`, Preparation `+2.67`, Activity `+1.33` — these check
out against the spec's own arithmetic).

### Task 16 — `mechanics-draft` evidence-pass extension (§3.5)

Two additive evidence passes in `logs/cmd/forever-logs/mechanics_draft.go`,
run over the same drafted pulls: a utility pass (every aura a player was
the sole/majority source of, not already in the fight's mechanics
vocabulary, proposed as a candidate `owned` entry with a `"draft:"`-
prefixed note naming fight count and mean uptime) and a downtime pass
(an enemy cast/aura the phase-candidate scan already collects that
repeats *within* one phase with a consistent duration, proposed as a
downtime candidate instead of a phase candidate). Output alongside the
existing `mechanics.Table` JSON, to stderr evidence lines (same
convention the phase evidence already uses).

### Task 17 — Final verification and whole-branch review

`GOTOOLCHAIN=go1.25.11 go vet ./... && GOTOOLCHAIN=go1.25.11 go test ./...`
from `logs/`, confirm every existing golden file still matches (or was
deliberately, minimally regenerated per Task 1), dispatch a code-reviewer
pass over the diff, fix-wave any findings, update the progress ledger with
every ruling, commit.
