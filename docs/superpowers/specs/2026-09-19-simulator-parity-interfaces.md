# Simulator parity: the interfaces

**Date:** 2026-09-19
**Design:** `2026-09-19-simulator-parity-design.md`
**Amends:** `2026-09-14-simulator-interfaces.md`. Everything there stands;
this document adds to it. Where a type below already exists, the new
fields are marked `+`.

The lanes (engine fork, `sim` module, data, API, web, addon) code against
this document and nothing else. A lane that needs a name not written here
adds it here first, in its own commit, and the other lanes pick it up.

## 1. Request envelope (`sim/api/envelope.go`)

### 1.1 Kinds

```go
const (
    KindRun     = "run"      // one character, one result (today's request)
    KindGear    = "gear"     // Top Gear: combinations of substitutions
    KindTalents = "talents"  // talent loadouts only
    KindDrops   = "drops"    // one substitution at a time, grouped by source
    KindWeights = "weights"  // stat weights
)

// Kind is derived, never sent: a request with Bulk is gear, talents or
// drops by Bulk.Mode; one with Weights is weights; otherwise run.
func (r SimRequest) Kind() string
```

### 1.2 SimRequest

```go
type SimRequest struct {
    // ... existing fields unchanged ...
    Bulk    *BulkSpec    `json:"bulk,omitempty"`     // +
    Weights *WeightsSpec `json:"weights,omitempty"`  // +
    // TargetError, when > 0, turns Iterations into a ceiling: the run
    // continues in steps of StepIterations until DPS.Error/DPS.Mean is
    // at or under TargetError or Iterations is reached. 0 is today's
    // fixed-count run.
    TargetError float64 `json:"target_error,omitempty"` // +
}

const StepIterations = 1000
var LaneIterationCeiling = map[string]int{LaneBrowser: 30000, LaneServer: 100000}
```

### 1.3 BulkSpec

```go
type BulkSpec struct {
    Mode       string          `json:"mode"`        // gear | talents | drops
    Candidates []Candidate     `json:"candidates"`
    Talents    []TalentLoadout `json:"talents,omitempty"`
    Sets       []GearSet       `json:"sets,omitempty"`
    Locked     []string        `json:"locked,omitempty"`   // slots never substituted
    Precision  string          `json:"precision"`          // fast | normal | high
    Cap        int             `json:"cap"`                // the lane's cap, echoed so a saved request says what bounded it
}

type Candidate struct {
    Slot    string `json:"slot"`              // IDS.md slot vocabulary; "" means "wherever it fits" (rings, trinkets, weapons)
    ItemID  int    `json:"item_id"`
    Enchant int    `json:"enchant,omitempty"` // 0 inherits the equipped enchant for the slot where it fits
    Suffix  int    `json:"suffix,omitempty"`
    Origin  string `json:"origin"`            // equipped | bag | bank | search | drop:<source-id> | set:<name>
}

type TalentLoadout struct {
    Name    string `json:"name"`
    Talents string `json:"talents"` // positional string, same as CharacterSpec.Talents
}

type GearSet struct {
    Name string     `json:"name"`
    Gear []GearSlot `json:"gear"`
}

const (
    PrecisionFast   = "fast"
    PrecisionNormal = "normal"
    PrecisionHigh   = "high"
)

// Stages by precision: iterations per stage, and what survives each cut.
// The equipped set runs in every stage.
//   fast:   100 → top 25% (+ within 2 SE of the cut) → 1000 → top 10 (+ ties) → 3000
//   normal: 1000 → top 10 (+ ties) → 3000
//   high:   1000 → top 20 (+ ties) → 10000
var Caps = map[string]int{LaneBrowser: 400, LaneServer: 20000}
```

Validation: `Mode` in the set; `Precision` in the set; `Cap` at most the
lane's; no candidate on a locked slot; every `Origin` matches the pattern;
`talents` mode has at least one loadout and no candidates; `drops` mode
has candidates whose origins are all `drop:`; `gear` has at least one of
candidates, talents or sets. Expansion above `Cap` is refused with
`ErrCapExceeded{Cap, Combinations}`.

### 1.4 WeightsSpec

```go
type WeightsSpec struct {
    Stats     []string `json:"stats"`     // IDS.md stat ids: strength, agility, attack_power, crit, hit, haste, spell_power, ... 
    Reference string   `json:"reference"` // the stat normalised to 1.0; the spec's default from data/curated/specs.json
}
```

### 1.5 EncounterSpec

```go
type EncounterSpec struct {
    // ... existing: DurationSec, Variation, Targets, ExecuteRatio, Profile ...
    Style           string        `json:"style,omitempty"`             // + label only; see 1.6
    Movement        *Movement     `json:"movement,omitempty"`          // +
    TargetsOverTime []TargetCount `json:"targets_over_time,omitempty"` // + overrides Targets when set
    TargetLevel     int           `json:"target_level,omitempty"`      // + 60..63, default 63
    TargetArmor     int           `json:"target_armor,omitempty"`      // + 0 means the level's preset
    TargetType      string        `json:"target_type,omitempty"`       // + humanoid | undead | beast | demon | dragonkin | elemental | giant | mechanical | unknown
    Dummy           bool          `json:"dummy,omitempty"`             // + no debuffs, no execute, no armor reduction
}

type Movement struct {
    IntervalSec int    `json:"interval_sec"`
    DurationSec int    `json:"duration_sec"`
    Kind        string `json:"kind"` // away (out of melee, no casting) | casting (spells interrupted, melee continues)
}

type TargetCount struct {
    AtSec int `json:"at_sec"`
    Count int `json:"count"`
}
```

### 1.6 Styles

A style is a page preset that expands to encounter fields; the envelope
stores the fields and keeps `Style` as the label. Vocabulary and expansion:

| Style id | Targets | ExecuteRatio | Movement | TargetsOverTime | Dummy |
| --- | --- | --- | --- | --- | --- |
| `patchwerk` | 1 | 0.25 | none | none | false |
| `execute` | 1 | 0.35 | none | none | false |
| `light-movement` | 1 | 0.25 | 45 s / 5 s / away | none | false |
| `heavy-movement` | 1 | 0.25 | 20 s / 5 s / away | none | false |
| `cleave-2`, `cleave-3`, `cleave-5` | 2, 3, 5 | 0.25 | none | none | false |
| `dungeon` | 1 | 0 | none | 0 s: 1, 40 s: 3, 80 s: 5, 130 s: 3, 160 s: 1 | false |
| `dummy` | 1 | 0 | none | none | true |

### 1.7 Buffs, consumables, cooldowns

- IDS.md gains graded ids: `<id>:improved` for every `TristateEffect`
  field (e.g. `battle_shout:improved`). The plain id stays the plain form.
- IDS.md gains a **World buffs** section listing every `IndividualBuffs`
  world-buff field by its snake-case name.
- `CharacterSpec` gains `Cooldowns []CooldownSpec` (+):

```go
type CooldownSpec struct {
    ID    string    `json:"id"`     // spell id as "spell:<id>" or a consumable id from IDS.md
    AtSec []float64 `json:"at_sec"` // times to use; empty means "on cooldown"
}
```

## 2. Result envelope

```go
type SimResult struct {
    // ... existing fields unchanged ...
    Combos   []Combo      `json:"combos,omitempty"`   // + ranked, best first
    Equipped *Estimate    `json:"equipped,omitempty"` // + the base character at the final stage
    Stages   []Stage      `json:"stages,omitempty"`   // +
    Weights  []StatWeight `json:"weights,omitempty"`  // +
    Sample   []SampleCast `json:"sample,omitempty"`   // + one iteration's casts (the median-DPS iteration)
}

type Combo struct {
    Substitutions []Substitution `json:"substitutions"`
    DPS           Estimate       `json:"dps"`
    Delta         Estimate       `json:"delta"` // against Equipped, paired at the same stage
    Group         int            `json:"group"` // 0 for the leader's within-error group, then 1, 2, ...
}

type Substitution struct {
    Kind    string `json:"kind"`              // item | talents | set
    Slot    string `json:"slot,omitempty"`    // item: the slot it went into (rings and trinkets say which)
    ItemID  int    `json:"item_id,omitempty"`
    Enchant int    `json:"enchant,omitempty"`
    Suffix  int    `json:"suffix,omitempty"`
    Name    string `json:"name,omitempty"`    // talents/set: the loadout or set name
    Talents string `json:"talents,omitempty"`
    Origin  string `json:"origin,omitempty"`  // copied from the candidate
}

type Stage struct {
    Iterations int `json:"iterations"`
    Combos     int `json:"combos"`
}

type StatWeight struct {
    Stat   string  `json:"stat"`
    Weight float64 `json:"weight"` // Reference stat is exactly 1
    Error  float64 `json:"error"`
}

type SampleCast struct {
    AtMS      int64            `json:"at_ms"` // negative during pre-pull
    SpellID   int64            `json:"spell_id"`
    Name      string           `json:"name"`
    Target    string           `json:"target,omitempty"`
    Resources map[string]int   `json:"resources,omitempty"` // rage, energy, mana, combo_points after the cast
}
```

`Progress` (+): `Stage int`, `CombosDone int`, `CombosTotal int`, all zero
for a plain run.

## 3. `sim/bulk` (new package)

```go
// Expand lists every valid combination for req. It reads item rows from
// the same simdb the engine loads. ErrCapExceeded carries the count.
func Expand(req api.SimRequest) ([]Combination, error)

type Combination struct {
    Request       api.SimRequest    // the base character with the substitutions applied
    Substitutions []api.Substitution
}

// Plan is the first stage: the combinations plus the equipped set, each
// as a request at the stage's iteration count. Stage numbers start at 1.
func Plan(req api.SimRequest) (StageRequests, error)

type StageRequests struct {
    Stage      int
    Iterations int
    Requests   []api.SimRequest // Requests[0] is always the equipped set
    Combos     []Combination    // parallel to Requests[1:]
}

// Rank scores a finished stage. It returns the next stage or, after the
// final stage, the SimResult with Combos, Equipped and Stages filled.
func Rank(req api.SimRequest, stage StageRequests, results []api.SimResult) (next *StageRequests, final *api.SimResult, err error)
```

Rules Expand enforces, each with a table test: slot fit by inventory type
(`sim/internal/simdb`), class allowlist, required level, faction, unique
and unique-category limits, rings and trinkets in both slots, two-hand
versus main-plus-off-hand, dual-wield weapon order, enchant slot and item
type fit, enchant inheritance from the equipped item, `Locked`.

## 4. wasm exports (`sim/cmd/wasm`)

All JSON strings in and out, failures `{"error": "..."}`, as today.

| Export | In | Out |
| --- | --- | --- |
| `simPlan(requestJSON)` | a bulk SimRequest | `{"stage":1,"iterations":100,"requests":[SimRequest,...],"combos":[Combination,...]}` |
| `simRank(requestJSON, stageJSON, resultsJSON)` | the request, the stage object simPlan/simRank returned, an array of SimResult in the same order | `{"next": stage}` or `{"result": SimResult}` |
| `simWeights(requestJSON, callbackId)` | a weights SimRequest | SimResult with Weights; progress via `simProgress` |

`simRun` is unchanged; the page runs each stage's requests through
`simSplit`/`simRun`/`simCombine` exactly as it runs a single sim. The
native `forever-sim` binary detects the kind and runs the plan-rank loop
itself; `-progress` lines gain `stage`, `combos_done`, `combos_total`.

## 5. Engine fork (behind the pin)

`proto/common.proto` `Encounter` gains:

```
message MovementPattern { double interval_seconds = 1; double duration_seconds = 2; bool casting_only = 3; }
MovementPattern movement = 10;
message TargetCountAt { double at_seconds = 1; int32 count = 2; }
repeated TargetCountAt targets_over_time = 11;
bool target_dummy = 12;
```

Behaviour: `movement` schedules `MovementHandler` moves out of range for
`duration_seconds` every `interval_seconds` (or, with `casting_only`,
interrupts casting without moving); `targets_over_time` activates and
deactivates targets from a pool sized to the maximum count;
`target_dummy` disables debuff application, execute windows and the
target's armor reduction. `RaidSimResult` gains `sample_iteration` with
the cast log and resource readings of the median-DPS iteration.

## 6. Data files

### 6.1 `data/builds/<build>/loot.json`

```json
{
  "sources": [
    { "id": "raid:mc", "kind": "raid", "name": "Molten Core", "zone_id": 409, "opens": "raids-1",
      "bosses": [ { "id": "raid:mc:lucifron", "name": "Lucifron", "npc_id": 12118, "items": [16800, 16803] } ],
      "trash": [ 17011 ] },
    { "id": "dungeon:brd", "kind": "dungeon", ... },
    { "id": "world:azuregos", "kind": "world", "items": [...] },
    { "id": "crafted:blacksmithing", "kind": "crafted", "profession": "blacksmithing", "items": [...] },
    { "id": "rep:argent-dawn:exalted", "kind": "rep", "faction_id": 529, "standing": "exalted", "items": [...] },
    { "id": "pvp:rank-10", "kind": "pvp", "rank": 10, "items": [...] },
    { "id": "quest", "kind": "quest", "items": [...] }
  ]
}
```

Generated by the data lane from the fork database's `sources` (AtlasLoot
and Wowhead), joined to the build's zones, then overlaid by
`data/curated/loot/*.json`, which use the same shape plus `sources` and
`notes` like every curated file and may add, replace or remove a source or
an item. `opens` is a phase name from `api/internal/phase`; a source
without it is open from launch.

### 6.2 `data/builds/<build>/enchants.json`

`[ { "id": <effect_id>, "name", "icon", "slots": ["head", ...], "item_types": [...], "classes": [...], "stats": {...}, "phase" } ]`
from the fork database's `UIEnchant`, with Forever's additions overlaid
from `data/curated/enchants.json`.

### 6.3 `data/builds/<build>/suffixes.json`

`[ { "id", "name", "stats": {...} } ]` from `ItemRandomSuffix`; `items.json`
rows gain `suffixes: [id, ...]` where the item rolls one.

### 6.4 Build validation

The pipeline fails when any `ItemSparse` row has a non-zero
`SocketType_*` (the gem signal from the design's section 4.4).

## 7. The addon export, version 2 (`FS1`)

```
FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>|bags=<items>|bank=<items>|sets=<name>=<gear>;...|loadouts=<name>=<t1>/<t2>/<t3>;...
```

- Everything before the first `|` is version 1 unchanged, so every
  existing decoder keeps working.
- `<items>` is `item_id[:enchant[:suffix]]` joined by `,`; only
  equippable items; empty sections are omitted.
- `<gear>` inside `sets=` is the version-1 gear list; `<name>` is URL-encoded.
- Order of sections is fixed as written; unknown sections are ignored by
  the decoder and reported in its result as `ignored: [name]`.
- The decoder (`web/src/lib/planner/fs1.ts`) returns `FS1Build` (+)
  `bags`, `bank`, `sets`, `loadouts`, all optional. `MAX_CODE_LENGTH`
  rises to 16,384.
- The encoder lands in the in-game addon's `Export.lua` (the phase-2
  addon design); the companion passes strings through untouched.

## 8. API

- `POST /v1/sims`: unchanged path and premium gate. A bulk or weights
  request is validated with the server lane's cap; a cap breach answers
  `400 cap_exceeded` with `{ "cap": 20000, "combinations": 31200 }`. A bulk
  request the planner estimates past the job timeout answers `400
  too_large` with the estimate in seconds.
- `sims` (+): `kind text not null default 'run'`; migration `0014_sim_kinds`.
  `Store.Queue` sets it from `req.Kind()`.
- `GET /v1/sims/<id>`: unchanged; the result blob carries the new fields.
- `GET /v1/sims?mine=1` (+) accepts `kind=`; rows carry `kind` and
  `headline` (the API composes: run → "1,204 DPS"; gear → "+41 DPS from
  Vis'kag"; drops → "3 upgrades on Ragnaros"; talents → "+18 DPS with
  'Deep Fury'"; weights → "Crit 1.00 · Agility 0.87").
- Progress rows carry `stage`, `combos_done`, `combos_total`.
- `GET /v1/specs` (+) rows carry `reference_stat`.

## 9. Web

Routes: `/sim/gear`, `/sim/talents`, `/sim/drops`, `/sim/weights`, each
an island over the same store as `/sim`. `/sim/<id>` renders by
`result.request.Kind()`. Test ids: `sim-combos`, `sim-combo-row`,
`sim-equipped-line`, `sim-cap-notice`, `sim-stage-progress`,
`sim-source-picker`, `sim-source-<id>`, `sim-candidate-<slot>-<item>`,
`sim-weights`, `sim-request-drawer`, `sim-style`, `sim-precision`,
`sim-target-error`, `sim-sample-log`, `sim-details-card`.

Share URLs carry the whole request as today; a bulk request above the URL
budget is shared by its saved id only.

## 10. Amendments after planning (2026-09-19)

Six lane plans were written against sections 1–9 and reported where the
contract was ambiguous, wrong, or silent. Every ruling below is binding;
where it contradicts an earlier section, this section wins. Plans that
carried their own "amend the contract" task drop it: this is that task.

### 10.1 Envelope and validation

- **A1. Lane-aware caps.** `SimRequest` carries no lane. `Validate` checks
  `Bulk.Cap` against the largest lane's cap; `ValidateLane(lane string)`
  checks against that lane's and is what the API handler and the page call.
- **A2. Server cap is 5,000 combinations**, not 20,000: at the measured
  native rate a 20,000-combination fast run cannot finish in the job's
  15-minute timeout. `Caps = {browser: 400, server: 5000}`. The job timeout
  stays 15 minutes. The `too_large` estimate uses
  `measure.NativeIterationsPerCPUSecond` (the constant `sim/measure`
  publishes from its benchmark, 1,218 today) × 4 CPUs × 840 seconds.
- **A3. Iterations.** A plain fixed run keeps the closed set
  `ValidIterations`. A bulk request's `Iterations` is its precision's
  final-stage count (3,000; 10,000 for `high`). A `TargetError` run's
  `Iterations` is a positive multiple of `StepIterations` at or under
  `LaneIterationCeiling[lane]`. `Validate` applies the rule for the request's
  kind. `MinDurationSec = 20`, `MaxDurationSec = 600`.
- **A4. Mode decides expansion**; the design's `combinations` boolean is
  gone. `gear` takes the product of every candidate group, `drops` and
  `talents` one substitution at a time.
- **A5. `BulkSpec.Consumables [][]string`** (+): alternative consumable
  lists tried as candidates in `gear` mode (design 3.1 item 5). Each inner
  list replaces `CharacterSpec.Consumes` for that combination.
- **A6. `Candidate.SourceName string`** (+) and **`Substitution.Name` is
  filled for items too** (from simdb) and gains **`SourceName`** (copied
  from the candidate). The page fills `SourceName` from `loot.json` when it
  builds a drops request; the API's headline reads it.
- **A7. `WeightsSpec.Reference` is required.** The stat vocabulary is the
  fork's `proto.Stat` enum names in snake case (`spell_haste` and
  `melee_haste`, never a bare `haste`), published as a new **Stats** section
  in IDS.md that the sim module generates from the enum. `specs.json`'s
  `reference_stat` (data lane) and the generated `sim/specs` struct's
  `ReferenceStat` (+) use that vocabulary; the API serves it on `/v1/specs`.
- **A8. `TargetArmorByLevel`** = `{60: 3300, 61: 3444, 62: 3588, 63: 3731}`:
  the engine's own 3,731 at 63 and a linear fall to the level-60 figure.
  Ratified here; a better source replaces the three interior numbers.
- **A9. `ExpandWith(req, Options)`** exists beside `Expand`. Options carry
  the enchant table. The wasm and the native binary get enchants the way
  they get items: `make simdb` embeds the fork database's enchants beside
  its items in `sim/internal/simdb`, so `Expand` needs no file at runtime
  and `enchants.json` (data lane) is the same rows for the browser UI.
- **A10. `StageRequests.Ran []api.Stage`** (+) carries the ladder's history
  across the wasm boundary so `Rank` can fill `SimResult.Stages`.
- **A11. Progress widening is additive.** `runner.Progress` gains `Stage`,
  `CombosDone`, `CombosTotal`; the type is not replaced, so the `api`
  module keeps compiling when the `sim` module merges first.
- **A12. Sample rows carry action keys, not names.** `SampleCast` is
  `{ at_ms, action, target, resources }` where `action` is the summary's
  action-key form (`spell:23881`, `item:13503`, `other:melee`) and the
  page resolves the display name with `resolveActionName` exactly as it
  does for cast rows. `SpellID` and `Name` are removed from section 2.

### 10.2 wasm and native exports (section 4 additions)

| Export | In | Out |
| --- | --- | --- |
| `simCount(requestJSON)` | a bulk request | `{"combinations": n}` — counts without allocating requests |
| `simNeedsMore(resultJSON, requestJSON)` | a result and its request | `{"needs_more": bool}` (the Go original is `api.NeedsMoreIterations`) |
| `simValidate(requestJSON)` | any request | `{"ok": bool, "errors": [{"field","message"}]}` |

A cap breach from `simPlan` or `simCount` is
`{"error":"cap_exceeded","cap":n,"combinations":n}`. The native binary
gains `-plan` (print `simCount`'s answer and the first stage, run nothing),
which is how the API counts combinations without importing `sim/internal`.

Stage requests run unsplit: a stage's request array goes through the
worker pool as independent whole requests, in chunks of the pool's width,
which is what makes "abort returns the partial" true. `simSplit` and
`simCombine` are for plain runs only.

### 10.3 Engine fork (section 5 additions)

- `SimOptions.sample_iteration = 10` (bool) opts in; the sim module sets it
  for plain runs only, never for bulk stages.
- `RaidSimResult.sample_iteration = 8` of
  `SampleIteration { double dps = 1; double duration_seconds = 2; repeated SampleCast casts = 3; }`,
  `SampleCast { int64 at_ms = 1; ActionID action_id = 2; string target = 3; map<string,int32> resources = 4; }`.
- `SimItem` (the reduced item message simdb stores) gains
  `bool unique = 20; int32 required_level = 21; UIItem.FactionRestriction faction_restriction = 22; repeated int32 random_suffix_options = 23;`
  and the data lane fills them in `pipeline/simdb/items.py`.
- `targets_over_time`: the engine pads the target list by repeating the
  last target up to the timeline's maximum; the site sends one target.
- `target_dummy`: the raid debuff panel is not applied and nothing lowers
  the target's armor from any source; the player's own debuffs still land.
- Two pins move together: `sim/enginever/version.go` (sim module lane) and
  `data/proto/ENGINE_SHA` with the vendored protos (data lane, `python -m
  pipeline genproto`), in that order, in the same round.

### 10.4 Data files (section 6 corrections)

- Zone ids are AreaTable ids (Molten Core is 2717). Source ids are
  `raid:<zone-slug>` and `raid:<zone-slug>:<npc-id>`, likewise `dungeon:`;
  a boss with no name in either database is emitted with an empty name.
  `pvp:rank-<n>` comes from `ItemSparse.RequiredPVPRank`. Plain vendors
  and unnamed open-world drops have no kind and are dropped, counted in the
  pipeline's log.
- `opens` is a phase name; a source whose date is unknown carries
  `"opens": "later"`, which the page shows as unreleased without a date.
- The curated overlay shape is `{ "sources": [...provenance...], "notes":
  "...", "add": [loot sources], "replace": [loot sources], "remove": [ids] }`.
- `enchants.json` rows are keyed by `effect_id` plus `spell_id`/`item_id`;
  `item_types` is the `EnchantType` shape restriction; `slots` derives from
  `type` and `extra_types`.
- Suffixes come from the fork database's `randomSuffixes`; the client has
  no `ItemRandomSuffix` table on this build.
- New file **`data/builds/<build>/simbuffs.json`**:
  `{ "entries": { "<id>": { "name", "icon" } } }` for every IDS.md buff,
  debuff, world buff and consumable id.
- New file **`data/curated/phases.json`**: `[ { "name", "start" } ]`, the
  source of truth `api/internal/phase.Boundaries` is tested against, emitted
  to `web/src/data/phases.json` by the pipeline.
- The re-itemised raid tier: 1,809 of the fork's sourced item ids do not
  exist in the 1.60 client. `loot.json` lists only items the build has;
  the first overlay records the gap per raid in its notes. Raid Droptimizer
  is honest and thin until sourced replacements are curated.

### 10.5 Addon export (section 7 corrections)

`<gear>` entries and `sets=` gear lists use `item_id[:enchant[:suffix]]`
too; a version-1 decoder reading a bare id is unaffected. A
`professions=<slug>,<slug>` section follows `loadouts=`. `SimCharacter`
gains per-slot enchant and suffix.

### 10.6 API (section 8 corrections)

- The premium submit is **`POST /v1/sims/run`** (the existing route);
  `POST /v1/sims` remains the browser-result save. Both accept every kind.
- `cap_exceeded` and `too_large` fields travel as decimal strings
  (`httpx.ErrorBody.Fields` is `map[string]string`).
- Migration `0014_sim_kinds` adds `kind`, `headline text`, `stage int`,
  `combos_done int`, `combos_total int`. `headline` is composed at save
  time by the rules in 8 plus: several substitutions read "… and 2 more";
  empty results read "no combinations", "no upgrades", "no weights".
- `builds` gains `user_id` (nullable, set on save when signed in) and
  **`GET /v1/builds?mine=1`** lists the signed-in player's builds; Top Gear's
  talent list reads it.
- **`GET /v1/phases`** returns `phases.json`'s boundaries.

### 10.7 Web

- Test ids in section 9 are a minimum; page-local ids follow the same
  `sim-<thing>` shape.
- Zone icons do not exist in the content collection; the source picker
  labels by kind and name and draws none.
- "My professions" reads `CharacterSpec.Profession`, filled by the export's
  new section; absent that, the crafted picker shows all professions and
  says why.

### 10.8 Rulings from the plan revisions

- **Faction enum.** `UIItem.FactionRestriction` lives in `ui.proto`, which
  imports `common.proto`, so `SimItem` cannot reference it. `SimItem`
  declares its own nested `FactionRestriction` enum with identical value
  names and numbers, pinned to `UIItem`'s by a test.
- **Progress callback.** `runner.Progress` is a function type and stays
  byte-identical. The `sim` module adds `StageProgress func(api.Progress)`
  and a `StageRunner` interface that `Native` and `Fixture` implement; the
  API adopts the wider callback in its own task.
- **Enchant table in the engine.** `SimEnchant` is not widened. `make simdb`
  copies `data/builds/<build>/enchants.json` to
  `sim/internal/simdb/enchants.json` for `go:embed`; `bulk.Options.Enchants`
  is a test override over that default.
- **Consumable substitutions.** `Substitution.Kind` gains `consumes`, with
  `Name` the consumable ids joined by `, `; `sim/bulk` emits one per
  combination that used an alternative list.
- **Stat vocabulary, pinned.** The `proto.Stat` enum in snake case, with
  `MP5` spelled `mp5`: `strength, agility, stamina, intellect, spirit,
  spell_power, arcane_power, fire_power, frost_power, holy_power,
  nature_power, shadow_power, mp5, hit, crit, spell_haste,
  spell_penetration, attack_power, melee_haste, armor_penetration,
  expertise, mana, energy, rage, armor, ranged_attack_power, defense,
  block, block_value, dodge, parry, health, arcane_resistance,
  fire_resistance, frost_resistance, nature_resistance, shadow_resistance,
  bonus_armor, healing_power, spell_damage, feral_attack_power`. There is
  no `melee_crit`, `spell_crit`, `melee_hit` or `spell_hit`: the engine
  carries one `hit` and one `crit`. Weight pages offer the subset that
  moves a spec's DPS; `reference_stat` defaults are `attack_power` for
  melee and hunters, `spell_power` for casters.
- **Client-side server cap.** The page gates the server-run button on
  `Caps.server` before submitting; a server `cap_exceeded` that still
  arrives shows the generic failure sentence. Acceptable: the numbers are
  already on screen from `simCount`.
- **Overlay `replace` is partial.** A `replace` entry names an `id` and only
  the keys it changes; unnamed keys keep the generated value; a differing
  `kind` is refused.
- **Faction restriction comes from the fork database**, not the client
  (`ItemSparse.AllowableRace` is unset on this build). `items.json` gains a
  `faction_restriction` column from one fork-derived pass and
  `pipeline/simdb` reads it; the loot step runs before the simdb step.
