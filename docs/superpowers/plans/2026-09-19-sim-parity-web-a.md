# Simulator parity — web, part A (settings, precision, report, input) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Bring the existing `/sim` page to Raidbots parity on everything that is not a new route — fight styles, the full encounter panel, Smart-Sim precision, the whole buff/debuff/consumable panel, the report additions, the request drawer, FS1 version 2 input and history rows by kind.

**Architecture:** Every new decision lives in a small, pure, unit-tested module under `web/src/lib/sim/` (or `web/src/lib/planner/fs1.ts` for the export string). Svelte components are pure renders of those modules over the island store, exactly as the lane already works. No statistics, no validation and no expansion arithmetic are written in TypeScript: the step loop asks the wasm (`simNeedsMore`), the request drawer asks the wasm (`simValidate`), and part B's combination count asks the wasm (`simCount`) — all three ratified in contract 10.2. Part B (`/sim/gear`, `/sim/talents`, `/sim/drops`, `/sim/weights`) mounts the same store and imports the modules listed in each task's **Interfaces → Produces** block.

**Tech Stack:** Astro 7, Svelte 5 runes islands, TypeScript strict, vitest (`src/**/*.test.ts`), Playwright (`tests/e2e/`, fake engine under `src/fixtures/sim/`), Tailwind 4, Lighthouse CI.

**Spec:**
- Design: `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` (sections 4, 5, 8, 9 are this lane's)
- Contract: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` (sections 1.2, 1.5, 1.6, 1.7, 2, 4, 7, 9 bind this lane)
- **Contract section 10 (commit `439f0b7`) is binding over all of the above** and settles every question the six lane plans raised. This lane's rulings are A3 (iterations and the 20–600 s duration range), A8 (`TargetArmorByLevel`), A11 (progress widening), A12 (`SampleCast` carries an action key), 10.2 (`simCount`, `simNeedsMore`, `simValidate`), 10.4 (`simbuffs.json`), 10.5 (FS1 gear enchants, suffixes and professions; `SimCharacter` per-slot enchant and suffix) and 10.7 (test ids are a minimum). A7 has the sim module generate IDS.md's `:improved` rows, World buffs and a new Stats section, all three of which this lane's generator reads.
- Base contract it amends: `docs/superpowers/specs/2026-09-14-simulator-interfaces.md`

## Global Constraints

Every task's requirements implicitly include all of these.

- **Node 22.12 via nvm.** Every shell in this plan starts with `export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH`. Run everything from `web/`.
- **`npx astro check` must report 0 errors, 0 warnings, 0 hints** before every commit.
- **eslint and prettier clean on touched files**: `npx eslint <files>` and `npx prettier --check <files>`.
- **`E2E_PORT`** selects the preview port for Playwright (`playwright.config.ts` reads it, default 4321). Use `E2E_PORT=4399` for this lane so a parallel checkout's preview is never reused.
- **Every user-visible string lives in `web/src/lib/sim/copy.ts`** (`simCopy`). No literal copy in a `.svelte` file, in a test assertion, or in a `.ts` module other than `copy.ts`.
- **`data-testid` vocabulary is the contract's section 9**: `sim-combos`, `sim-combo-row`, `sim-equipped-line`, `sim-cap-notice`, `sim-stage-progress`, `sim-source-picker`, `sim-source-<id>`, `sim-candidate-<slot>-<item>`, `sim-weights`, `sim-request-drawer`, `sim-style`, `sim-precision`, `sim-target-error`, `sim-sample-log`, `sim-details-card`. Page-local additions this plan introduces are named in the task that adds them and follow the same `sim-<thing>` shape.
- **No `console.log`, no `console.warn`, no `debugger`** anywhere.
- **Immutability**: every settings/state helper returns a new object. Never mutate a `$state` object in place.
- **Files stay under 800 lines**; prefer a new focused module over growing an existing one.
- **One commit per task**, message `feat(sim): …` or `test(sim): …`, wrapped at 72 columns, ending with the trailer:
  ```
  Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
  ```
- **Never `git stash`** (the stash stack is shared with other worktrees). Set work aside with a temporary WIP commit instead.
- **TDD**: the failing test is written and *run* before the implementation, every time.

## File Structure

New files, and what each one is responsible for.

| File | Responsibility |
| --- | --- |
| `src/lib/sim/kind.ts` | `SimKind`, `requestKind(request)` — the contract's derived kind, in TypeScript |
| `src/lib/sim/styles.ts` | The fight-style table (contract 1.6) and its expansion into `EncounterSpec` |
| `src/fixtures/sim/styles.json` | The same table as data, checked by this lane's test and by the sim module lane |
| `src/lib/sim/precision.ts` | fast/normal/high/`until ±0.5%`, step size, lane ceilings |
| `src/lib/sim/buffs.ts` | Grouping the engine's id vocabulary into the panel's sections; the three-way grade |
| `src/lib/sim/buff-names.ts` | Display name and icon for a buff/consumable id, from the build |
| `src/lib/sim/stats.ts` | Contract A7's stat vocabulary, for part B's weights page |
| `src/lib/sim/cooldowns.ts` | `CooldownSpec` rows: on cooldown / on pull / at a time / at execute |
| `src/lib/sim/details.ts` | The details card's figures: margin of error, iterations, processing time, lane |
| `src/lib/sim/sample-log.ts` | `result.sample` into table rows, pre-pull separated, resource columns |
| `src/lib/sim/request-json.ts` | Pretty-print / parse / apply an edited request |
| `src/lib/sim/history.ts` | `SimListRow` kind filtering and headline fallback |
| `src/components/sim/SettingsSheet.svelte` | The secondary encounter controls, in the site's `<details>` disclosure |
| `src/components/sim/BuffPanel.svelte` | The full grouped buff/debuff/consumable panel behind "Custom" |
| `src/components/sim/CooldownRows.svelte` | Cooldown timing rows inside the panel |
| `src/components/sim/DetailsCard.svelte` | Margin of error, iterations, processing time, engine version, lane |
| `src/components/sim/SampleLog.svelte` | The sample-iteration log table |
| `src/components/sim/RotationCard.svelte` | The rotation the run used, with the fidelity note |
| `src/components/sim/RequestDrawer.svelte` | The editable request JSON with inline validation errors |
| `scripts/sync-sim-ids.mjs` | Parses `sim/request/IDS.md` into `src/data/generated/sim-ids.json` |
| `src/fixtures/sim/sample.json` | A sample-iteration cast log the fake engine attaches to its result |

Modified: `src/lib/sim/types.ts`, `settings.ts`, `encounter.ts`, `run.ts`, `engine.ts`, `worker.ts`, `sim.worker.ts`, `store.svelte.ts`, `url.ts`, `copy.ts`, `character.ts`, `src/lib/planner/fs1.ts`, `src/fixtures/sim/engine-fake.ts`, `src/components/sim/SettingsBar.svelte`, `RunControl.svelte`, `SimResults.svelte`, `SimHistory.svelte`, `SimView.svelte`, `src/components/report/AuraTable.svelte`, `scripts/sync-data.mjs`, `package.json`.

---

### Task 1: The amended envelope in TypeScript, and the derived kind

The contract makes the Go `json:` tags authoritative and `web/src/lib/sim/types.ts` their verbatim mirror. Every later task in this plan and every task in part B reads these names, so they land first. `Kind` is derived and never sent (contract 1.1), so it is a function here too.

**Files:**
- Modify: `src/lib/sim/types.ts`
- Modify: `src/lib/sim/copy.ts`
- Create: `src/lib/sim/kind.ts`
- Create: `src/fixtures/sim/envelope-v2.json`
- Test: `src/lib/sim/kind.test.ts`

**Interfaces:**
- Consumes: nothing.
- Produces:
  - `types.ts`: `EncounterSpec` (+ `style?`, `movement?`, `targets_over_time?`, `target_level?`, `target_armor?`, `target_type?`, `dummy?`), `Movement`, `TargetCount`, `CooldownSpec`, `CharacterSpec` (+ `cooldowns?`), `SimRequest` (+ `bulk?`, `weights?`, `target_error?`), `BulkSpec` (with A5's `consumables?`), `Candidate` (with A6's `source_name?`), `TalentLoadout`, `GearSet`, `WeightsSpec`, `SimResult` (+ `combos?`, `equipped?`, `stages?`, `weights?`, `sample?`), `Combo`, `Substitution` (10.8: `kind` includes `consumes`), `Stage`, `StatWeight`, `SampleCast` (contract A12: `{ at_ms, action, target?, resources? }`), `SimProgress` (+ `stage?`, `combos_done?`, `combos_total?`), `SimListRow` (+ `kind`, `headline`), `STEP_ITERATIONS_DEFAULT`
  - `kind.ts`: `type SimKind = 'run' | 'gear' | 'talents' | 'drops' | 'weights'`, `SIM_KINDS: readonly SimKind[]`, `requestKind(request: Pick<SimRequest, 'bulk' | 'weights'>): SimKind`
  - `copy.ts`: `simCopy.substitutionKindLabel: Record<string, string>` — the four `Substitution.kind` words, `consumes` included (10.8)

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/kind.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import envelope from '../../fixtures/sim/envelope-v2.json';
import { simCopy } from './copy';
import { SIM_KINDS, requestKind } from './kind';
import type { SimRequest, SimResult } from './types';

// The fixture is the contract's own shapes as data. Assigning it to the mirrored types is
// the compile-time half of this test (`npx astro check`); the assertions below are the
// runtime half -- that the key names the fixture uses are the ones the code reads.
const fixture = envelope as unknown as { request: SimRequest; result: SimResult };

describe('requestKind', () => {
  it('is run for a request with neither bulk nor weights', () => {
    expect(requestKind({ bulk: undefined, weights: undefined })).toBe('run');
  });

  it('takes gear, talents and drops from bulk.mode', () => {
    for (const mode of ['gear', 'talents', 'drops'] as const) {
      expect(requestKind({ bulk: { ...fixture.request.bulk!, mode }, weights: undefined })).toBe(mode);
    }
  });

  it('is weights when a weights block is present', () => {
    expect(requestKind({ bulk: undefined, weights: { stats: ['crit'], reference: 'crit' } })).toBe(
      'weights',
    );
  });

  it('prefers bulk over weights so a malformed request never reports two kinds', () => {
    expect(
      requestKind({ bulk: { ...fixture.request.bulk!, mode: 'gear' }, weights: { stats: [], reference: '' } }),
    ).toBe('gear');
  });

  it('falls back to run for a bulk block with a mode nothing recognises', () => {
    expect(requestKind({ bulk: { ...fixture.request.bulk!, mode: 'nonsense' }, weights: undefined })).toBe(
      'run',
    );
  });

  it('lists the contract’s five kinds', () => {
    expect(SIM_KINDS).toEqual(['run', 'gear', 'talents', 'drops', 'weights']);
  });
});

describe('the amended envelope', () => {
  it('carries the new encounter fields under the contract’s names', () => {
    expect(fixture.request.encounter.style).toBe('light-movement');
    expect(fixture.request.encounter.movement).toEqual({
      interval_sec: 45,
      duration_sec: 5,
      kind: 'away',
    });
    expect(fixture.request.encounter.targets_over_time).toEqual([{ at_sec: 0, count: 1 }]);
    expect(fixture.request.encounter.target_level).toBe(63);
    expect(fixture.request.encounter.target_armor).toBe(0);
    expect(fixture.request.encounter.target_type).toBe('humanoid');
    expect(fixture.request.encounter.dummy).toBe(false);
    expect(fixture.request.target_error).toBe(0.005);
    expect(fixture.request.character.cooldowns).toEqual([{ id: 'spell:11305', at_sec: [0, 90] }]);
  });

  it('carries the new result fields under the contract’s names', () => {
    // Contract A12: an action key, never a name and never a spell id.
    expect(fixture.result.sample?.[0]).toEqual({
      at_ms: -1500,
      action: 'spell:11305',
      resources: { rage: 0 },
    });
    expect(fixture.result.stages).toEqual([{ iterations: 1000, combos: 4 }]);
    expect(fixture.result.combos?.[0].group).toBe(0);
    expect(fixture.result.weights?.[0]).toEqual({ stat: 'crit', weight: 1, error: 0.02 });
  });
});

describe('substitution kinds', () => {
  it('names all four, including contract 10.8’s consumes', () => {
    expect(Object.keys(simCopy.substitutionKindLabel).sort()).toEqual([
      'consumes',
      'item',
      'set',
      'talents',
    ]);
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/kind.test.ts
```

Expected: FAIL — `Failed to resolve import "./kind"` and `"../../fixtures/sim/envelope-v2.json"`.

- [ ] **Step 3: Add the fixture**

Create `src/fixtures/sim/envelope-v2.json`:

```json
{
  "request": {
    "engine_version": "edc0c8e9a",
    "spec": "warrior-fury",
    "source": { "kind": "addon", "ref": "", "captured_at": "2026-09-19T10:00:00Z" },
    "character": {
      "name": "Thrallgar",
      "race": "orc",
      "class": "warrior",
      "level": 60,
      "talents": "-5530515-",
      "gear": [{ "slot": "head", "item_id": 12640 }],
      "buffs": ["battle_shout:improved"],
      "consumes": ["flask_of_supreme_power"],
      "cooldowns": [{ "id": "spell:11305", "at_sec": [0, 90] }]
    },
    "encounter": {
      "duration_sec": 180,
      "variation": 0.2,
      "targets": 1,
      "execute_ratio": 0.25,
      "profile": "",
      "style": "light-movement",
      "movement": { "interval_sec": 45, "duration_sec": 5, "kind": "away" },
      "targets_over_time": [{ "at_sec": 0, "count": 1 }],
      "target_level": 63,
      "target_armor": 0,
      "target_type": "humanoid",
      "dummy": false
    },
    "iterations": 30000,
    "random_seed": 0,
    "target_error": 0.005,
    "bulk": {
      "mode": "gear",
      "candidates": [{ "slot": "head", "item_id": 16963, "origin": "bag" }],
      "talents": [{ "name": "Deep Fury", "talents": "-5530515-" }],
      "sets": [{ "name": "AQ set", "gear": [{ "slot": "head", "item_id": 21329 }] }],
      "locked": ["main_hand"],
      "precision": "normal",
      "cap": 400
    },
    "weights": { "stats": ["crit", "agility"], "reference": "crit" }
  },
  "result": {
    "engine_version": "edc0c8e9a",
    "request": {
      "engine_version": "edc0c8e9a",
      "spec": "warrior-fury",
      "source": { "kind": "addon", "ref": "", "captured_at": "2026-09-19T10:00:00Z" },
      "character": {
        "name": "Thrallgar",
        "race": "orc",
        "class": "warrior",
        "level": 60,
        "talents": "-5530515-",
        "gear": [],
        "buffs": [],
        "consumes": []
      },
      "encounter": {
        "duration_sec": 180,
        "variation": 0.2,
        "targets": 1,
        "execute_ratio": 0.25,
        "profile": ""
      },
      "iterations": 3000,
      "random_seed": 0
    },
    "lane": "browser",
    "dps": { "mean": 1204.5, "stddev": 180.2, "error": 3.3, "min": 700, "max": 1800 },
    "iterations_run": 3000,
    "duration_ms": 4200,
    "summary": {
      "engine_version": "edc0c8e9a",
      "fight_index": 0,
      "duration_ms": 180000,
      "damage_done": [],
      "damage_taken": [],
      "healing": [],
      "healing_taken": [],
      "deaths": [],
      "auras": [],
      "casts": [],
      "interrupts": [],
      "dispels": [],
      "resources": [],
      "threat": [],
      "threat_by_target": [],
      "taunts": [],
      "combatants": [],
      "roster": [],
      "mechanics": [],
      "phases": []
    },
    "sample": [
      { "at_ms": -1500, "action": "spell:11305", "resources": { "rage": 0 } },
      { "at_ms": 320, "action": "spell:25286", "target": "Target", "resources": { "rage": 42 } }
    ],
    "stages": [{ "iterations": 1000, "combos": 4 }],
    "equipped": { "mean": 1180, "stddev": 170, "error": 3.1, "min": 690, "max": 1760 },
    "combos": [
      {
        "substitutions": [{ "kind": "item", "slot": "head", "item_id": 16963, "origin": "bag" }],
        "dps": { "mean": 1230, "stddev": 175, "error": 3.2, "min": 700, "max": 1810 },
        "delta": { "mean": 50, "stddev": 20, "error": 1.1, "min": -10, "max": 110 },
        "group": 0
      }
    ],
    "weights": [{ "stat": "crit", "weight": 1, "error": 0.02 }]
  }
}
```

- [ ] **Step 4: Extend `types.ts`**

In `src/lib/sim/types.ts`, replace the `EncounterSpec` interface and add the new shapes. The existing `DEFAULT_ENCOUNTER`, `CharacterSpec`, `SimRequest`, `SimResult`, `SimProgress` and `SimListRow` declarations are edited in place; everything else in the file is untouched.

```ts
/** A movement window the APL's movement conditions honour (contract 1.5). */
export interface Movement {
  interval_sec: number;
  duration_sec: number;
  /** away: out of melee, no casting. casting: spells interrupted, melee continues. */
  kind: 'away' | 'casting';
}

/** A step in the target-count timeline. Overrides `targets` when the list is non-empty. */
export interface TargetCount {
  at_sec: number;
  count: number;
}

export const TARGET_TYPE_IDS = [
  'humanoid',
  'undead',
  'beast',
  'demon',
  'dragonkin',
  'elemental',
  'giant',
  'mechanical',
  'unknown',
] as const;
export type TargetType = (typeof TARGET_TYPE_IDS)[number];

export interface EncounterSpec {
  duration_sec: number;
  variation: number;
  targets: number;
  execute_ratio: number;
  /** "" | "patchwerk" | "encounter:<encounter_id>". */
  profile: string;
  /** The fight style's id: a label only. The fields above and below are what the engine reads. */
  style?: string;
  movement?: Movement | null;
  targets_over_time?: TargetCount[] | null;
  /** 60..63; 63 is the default the engine assumes when this is absent. */
  target_level?: number;
  /** 0 means the level's preset, resolved inside the engine. */
  target_armor?: number;
  target_type?: TargetType | '';
  /** No debuffs, no execute, no armor reduction. */
  dummy?: boolean;
}

/** When to use a cooldown. Empty `at_sec` means "on cooldown" (contract 1.7). */
export interface CooldownSpec {
  /** "spell:<id>" or a consumable id from IDS.md. */
  id: string;
  at_sec: number[];
}
```

`CharacterSpec` gains one field:

```ts
  professions?: string[];
  /** When to use each major cooldown and potion. Absent means "everything on cooldown". */
  cooldowns?: CooldownSpec[];
```

`SimRequest` gains three:

```ts
export interface SimRequest {
  engine_version: string;
  spec: string;
  source: CharacterSource;
  character: CharacterSpec;
  encounter: EncounterSpec;
  /** With `target_error` set this is the ceiling, not the count. */
  iterations: number;
  random_seed: number;
  /**
   * When > 0, the run continues in `STEP_ITERATIONS_DEFAULT` steps until
   * `dps.error / dps.mean` is at or under this, or `iterations` is reached. 0 is a
   * fixed-count run. Contract 1.2.
   */
  target_error?: number;
  bulk?: BulkSpec;
  weights?: WeightsSpec;
}

/** Contract 1.2. The size of one step of a target-error run. */
export const STEP_ITERATIONS_DEFAULT = 1000;

export interface Candidate {
  /** IDS.md slot vocabulary; "" means "wherever it fits" (rings, trinkets, weapons). */
  slot: string;
  item_id: number;
  /** 0 inherits the equipped enchant for the slot where it fits. */
  enchant?: number;
  suffix?: number;
  /** equipped | bag | bank | search | drop:<source-id> | set:<name>. */
  origin: string;
  /**
   * Contract A6: the human name of where it came from ("Ragnaros"), which the page fills
   * from `loot.json` and the API's headline reads back off the substitution.
   */
  source_name?: string;
}

export interface TalentLoadout {
  name: string;
  talents: string;
}

export interface GearSet {
  name: string;
  gear: GearSlot[];
}

export interface BulkSpec {
  /**
   * gear | talents | drops. Contract A4: the mode decides the expansion and the design's
   * `combinations` boolean is gone -- `gear` takes the product of every candidate group,
   * `drops` and `talents` one substitution at a time.
   */
  mode: string;
  candidates: Candidate[];
  talents?: TalentLoadout[];
  sets?: GearSet[];
  /**
   * Contract A5: alternative consumable lists tried as candidates in `gear` mode. Each
   * inner list replaces `CharacterSpec.Consumes` for that combination.
   */
  consumables?: string[][];
  /** Slots never substituted. */
  locked?: string[];
  /** fast | normal | high. */
  precision: string;
  /** The lane's cap, echoed so a saved request says what bounded it. */
  cap: number;
}

export interface WeightsSpec {
  stats: string[];
  /** The stat normalised to 1.0. */
  reference: string;
}
```

`SimResult` gains five, plus the four new row shapes:

```ts
export interface Substitution {
  /**
   * item | talents | set | consumes. Contract 10.8 adds `consumes`: `sim/bulk` emits one
   * per combination that used an alternative consumable list (`BulkSpec.consumables`),
   * and `name` is that list's ids joined by ", ".
   */
  kind: string;
  slot?: string;
  item_id?: number;
  enchant?: number;
  suffix?: number;
  /**
   * The loadout or set name — and, per contract A6, an item's name too, filled from
   * simdb, so a combo row reads without a second lookup. For a `consumes` substitution
   * (10.8) it is the consumable ids joined by ", ".
   */
  name?: string;
  talents?: string;
  origin?: string;
  /** Contract A6: copied from the candidate. */
  source_name?: string;
}

export interface Combo {
  substitutions: Substitution[];
  dps: Estimate;
  /** Against `equipped`, paired at the same stage. */
  delta: Estimate;
  /** 0 for the leader's within-error group, then 1, 2, … */
  group: number;
}

export interface Stage {
  iterations: number;
  combos: number;
}

export interface StatWeight {
  stat: string;
  /** The reference stat is exactly 1. */
  weight: number;
  error: number;
}

/**
 * One cast of the median-DPS iteration. `at_ms` is negative during the pre-pull.
 *
 * Contract A12: the row carries the summary's own action-key form (`spell:23881`,
 * `item:13503`, `other:melee`) and nothing else -- no display name and no spell id. The
 * page resolves the name with `resolveActionName`, exactly as it already does for every
 * cast row, so the sample table can never disagree with the cast table about what an
 * action is called.
 */
export interface SampleCast {
  at_ms: number;
  action: string;
  target?: string;
  /** rage, energy, mana, combo_points … after the cast. */
  resources?: Record<string, number>;
}
```

and on `SimResult` itself, after `aborted?: boolean;`:

```ts
  /** Ranked, best first. Bulk kinds only. */
  combos?: Combo[];
  /** The base character at the final stage. Bulk kinds only. */
  equipped?: Estimate;
  stages?: Stage[];
  weights?: StatWeight[];
  /** One iteration's casts, the median-DPS one. */
  sample?: SampleCast[];
```

`SimProgress` gains three, all zero for a plain run:

```ts
export interface SimProgress {
  state: 'queued' | 'running' | 'done' | 'error';
  iterations_done: number;
  dps?: number;
  stage?: number;
  combos_done?: number;
  combos_total?: number;
}
```

`SimListRow` gains two (contract 8: the API composes the headline, the web renders it):

```ts
export interface SimListRow {
  sim_id: string;
  spec: string;
  dps: number;
  engine_version: string;
  created_at: string;
  title: string;
  /** run | gear | drops | talents | weights. Absent on a row saved before migration 0014. */
  kind?: string;
  /** The API's own one-line summary, e.g. "+41 DPS from Vis'kag". */
  headline?: string;
}
```

- [ ] **Step 5: Name the four substitution kinds**

A `Substitution` is rendered by part B's combo rows, but the words are this lane's, because every user-visible string on the simulator lives in `copy.ts`. Append inside `simCopy`:

```ts
  /**
   * The four things a combination can substitute (contract 2, `kind`, as amended by
   * 10.8). `consumes` is an alternative consumable list tried as a candidate, and its
   * `name` is the ids joined by ", " -- so the row reads "Consumables: flask_of_supreme_power,
   * elixir_of_the_mongoose" until buff-names.ts (Task 9) is given the list to prettify.
   */
  substitutionKindLabel: {
    item: 'Item',
    talents: 'Talents',
    set: 'Set',
    consumes: 'Consumables',
  } as Record<string, string>,
```

- [ ] **Step 6: Write `kind.ts`**

```ts
// web/src/lib/sim/kind.ts
// The contract's derived kind (2026-09-19-simulator-parity-interfaces.md, 1.1): "Kind is
// derived, never sent". `api.SimRequest.Kind()` is the Go original; this is its mirror, so
// the page can render a saved sim by what its stored request actually is rather than by a
// field a client could set to anything.
import type { SimRequest } from './types';

export type SimKind = 'run' | 'gear' | 'talents' | 'drops' | 'weights';

export const SIM_KINDS: readonly SimKind[] = ['run', 'gear', 'talents', 'drops', 'weights'];

const BULK_MODES: readonly SimKind[] = ['gear', 'talents', 'drops'];

/**
 * Bulk wins over weights: a request carrying both is malformed, and reporting one kind
 * rather than two is what lets `/sim/<id>` pick a single renderer without a tie-break of
 * its own. A `bulk.mode` outside the vocabulary reads as a plain run -- the engine will
 * refuse the request anyway, and the page must not render a bulk table for it meanwhile.
 */
export function requestKind(request: Pick<SimRequest, 'bulk' | 'weights'>): SimKind {
  const mode = request.bulk?.mode;
  if (mode !== undefined) {
    return BULK_MODES.find((known) => known === mode) ?? 'run';
  }
  return request.weights === undefined ? 'run' : 'weights';
}
```

- [ ] **Step 7: Run the test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/kind.test.ts
```

Expected: PASS, 8 tests.

- [ ] **Step 8: Check types, lint and format**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/lib/sim/kind.ts src/lib/sim/kind.test.ts src/lib/sim/types.ts src/lib/sim/copy.ts && \
  npx prettier --check src/lib/sim/kind.ts src/lib/sim/kind.test.ts src/lib/sim/types.ts src/lib/sim/copy.ts src/fixtures/sim/envelope-v2.json
```

Expected: 0 errors from each.

- [ ] **Step 9: Commit**

```bash
cd /Users/jh/code/forever/web && git add src/lib/sim/types.ts src/lib/sim/kind.ts src/lib/sim/kind.test.ts src/lib/sim/copy.ts src/fixtures/sim/envelope-v2.json && \
git commit -m "feat(sim): the amended request and result envelope in TypeScript

Mirrors contract 1.2, 1.5, 1.7 and 2 as amended by section 10: the
encounter's style, movement, target-count timeline, target level,
armor, type and dummy flag; the request's target_error, bulk and
weights blocks; the result's combos, equipped, stages, weights and
sample, with A12's action-key sample row and 10.8's consumes
substitution kind. requestKind is the TypeScript mirror of
api.SimRequest.Kind().

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 2: The fight-style table

Contract 1.6 is a table of nine styles, each expanding to encounter fields. It lands twice: as a TypeScript table the page uses, and as `src/fixtures/sim/styles.json`, which the sim module lane checks its own Go table against. The fixture is the shared artefact; the test asserts the two agree, so neither side can drift.

**Files:**
- Create: `src/lib/sim/styles.ts`
- Create: `src/fixtures/sim/styles.json`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/styles.test.ts`

**Interfaces:**
- Consumes: `types.ts`'s `EncounterSpec`, `Movement`, `TargetCount` (Task 1).
- Produces:
  - `styles.ts`: `type FightStyleId`, `interface FightStyle`, `FIGHT_STYLES: readonly FightStyle[]`, `DEFAULT_STYLE_ID: FightStyleId`, `fightStyle(id: string): FightStyle | null`, `applyFightStyle(encounter: EncounterSpec, id: FightStyleId): EncounterSpec`
  - `copy.ts`: `simCopy.styleLabel: Record<string, string>`, `simCopy.fightStyle`, `simCopy.styleNote: Record<string, string>`

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/styles.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import stylesJson from '../../fixtures/sim/styles.json';
import { simCopy } from './copy';
import { DEFAULT_STYLE_ID, FIGHT_STYLES, applyFightStyle, fightStyle } from './styles';
import { DEFAULT_ENCOUNTER } from './types';

interface StyleRow {
  id: string;
  targets: number;
  execute_ratio: number;
  movement: { interval_sec: number; duration_sec: number; kind: string } | null;
  targets_over_time: { at_sec: number; count: number }[] | null;
  dummy: boolean;
}

// The fixture is the shared artefact: the sim module lane's own table test reads this same
// file, so a style added on one side and not the other fails on both.
const fixture = (stylesJson as { styles: StyleRow[] }).styles;

describe('FIGHT_STYLES', () => {
  it('is exactly the fixture, in the fixture’s order', () => {
    expect(FIGHT_STYLES.map((style) => style.id)).toEqual(fixture.map((row) => row.id));
  });

  it('expands each style to the fixture’s encounter fields', () => {
    for (const row of fixture) {
      const style = fightStyle(row.id);
      expect(style, row.id).not.toBeNull();
      expect({
        id: style!.id,
        targets: style!.targets,
        execute_ratio: style!.execute_ratio,
        movement: style!.movement,
        targets_over_time: style!.targets_over_time,
        dummy: style!.dummy,
      }).toEqual(row);
    }
  });

  it('is the contract’s nine ids', () => {
    expect(FIGHT_STYLES.map((style) => style.id)).toEqual([
      'patchwerk',
      'execute',
      'light-movement',
      'heavy-movement',
      'cleave-2',
      'cleave-3',
      'cleave-5',
      'dungeon',
      'dummy',
    ]);
  });

  it('opens on Patchwerk', () => {
    expect(DEFAULT_STYLE_ID).toBe('patchwerk');
    expect(fightStyle(DEFAULT_STYLE_ID)?.execute_ratio).toBe(DEFAULT_ENCOUNTER.execute_ratio);
  });

  it('names every style, and notes only the two that need one', () => {
    for (const style of FIGHT_STYLES) {
      expect(simCopy.styleLabel[style.id], style.id).toBeTruthy();
    }
    expect(Object.keys(simCopy.styleNote).sort()).toEqual(['heavy-movement', 'light-movement']);
  });

  it('answers null for an id nothing defines rather than guessing', () => {
    expect(fightStyle('raidbots-patchwerk')).toBeNull();
    expect(fightStyle('')).toBeNull();
  });
});

describe('applyFightStyle', () => {
  it('writes the style’s fields and the label, and keeps duration and variation', () => {
    const next = applyFightStyle({ ...DEFAULT_ENCOUNTER, duration_sec: 300, variation: 0.1 }, 'cleave-3');
    expect(next.style).toBe('cleave-3');
    expect(next.targets).toBe(3);
    expect(next.execute_ratio).toBe(0.25);
    expect(next.duration_sec).toBe(300);
    expect(next.variation).toBe(0.1);
    expect(next.dummy).toBe(false);
  });

  it('clears a previous style’s movement and timeline rather than leaving them behind', () => {
    const moving = applyFightStyle(DEFAULT_ENCOUNTER, 'heavy-movement');
    expect(moving.movement).toEqual({ interval_sec: 20, duration_sec: 5, kind: 'away' });
    const back = applyFightStyle(moving, 'patchwerk');
    expect(back.movement).toBeNull();
    expect(back.targets_over_time).toBeNull();
  });

  it('gives the dungeon pull its target-count timeline and no execute window', () => {
    const dungeon = applyFightStyle(DEFAULT_ENCOUNTER, 'dungeon');
    expect(dungeon.execute_ratio).toBe(0);
    expect(dungeon.targets_over_time).toEqual([
      { at_sec: 0, count: 1 },
      { at_sec: 40, count: 3 },
      { at_sec: 80, count: 5 },
      { at_sec: 130, count: 3 },
      { at_sec: 160, count: 1 },
    ]);
  });

  it('turns the dummy flag on for the dummy and off for everything else', () => {
    expect(applyFightStyle(DEFAULT_ENCOUNTER, 'dummy').dummy).toBe(true);
    expect(applyFightStyle(applyFightStyle(DEFAULT_ENCOUNTER, 'dummy'), 'execute').dummy).toBe(false);
  });

  it('never mutates the encounter it was given', () => {
    const base = { ...DEFAULT_ENCOUNTER };
    applyFightStyle(base, 'cleave-5');
    expect(base.targets).toBe(1);
    expect(base.style).toBeUndefined();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/styles.test.ts
```

Expected: FAIL — `Failed to resolve import "./styles"`.

- [ ] **Step 3: Write the fixture**

Create `src/fixtures/sim/styles.json`. This is the contract's section 1.6 table, verbatim, and the sim module lane reads the same file.

```json
{
  "note": "Contract 2026-09-19-simulator-parity-interfaces.md section 1.6. Shared by web/src/lib/sim/styles.ts and the sim module's own style table; both test against this file so neither can drift.",
  "styles": [
    {
      "id": "patchwerk",
      "targets": 1,
      "execute_ratio": 0.25,
      "movement": null,
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "execute",
      "targets": 1,
      "execute_ratio": 0.35,
      "movement": null,
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "light-movement",
      "targets": 1,
      "execute_ratio": 0.25,
      "movement": { "interval_sec": 45, "duration_sec": 5, "kind": "away" },
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "heavy-movement",
      "targets": 1,
      "execute_ratio": 0.25,
      "movement": { "interval_sec": 20, "duration_sec": 5, "kind": "away" },
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "cleave-2",
      "targets": 2,
      "execute_ratio": 0.25,
      "movement": null,
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "cleave-3",
      "targets": 3,
      "execute_ratio": 0.25,
      "movement": null,
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "cleave-5",
      "targets": 5,
      "execute_ratio": 0.25,
      "movement": null,
      "targets_over_time": null,
      "dummy": false
    },
    {
      "id": "dungeon",
      "targets": 1,
      "execute_ratio": 0,
      "movement": null,
      "targets_over_time": [
        { "at_sec": 0, "count": 1 },
        { "at_sec": 40, "count": 3 },
        { "at_sec": 80, "count": 5 },
        { "at_sec": 130, "count": 3 },
        { "at_sec": 160, "count": 1 }
      ],
      "dummy": false
    },
    {
      "id": "dummy",
      "targets": 1,
      "execute_ratio": 0,
      "movement": null,
      "targets_over_time": null,
      "dummy": true
    }
  ]
}
```

- [ ] **Step 4: Add the copy**

In `src/lib/sim/copy.ts`, append inside the `simCopy` object, before the closing `} as const;`:

```ts
  // --- Parity, contract 1.6: the fight styles. The ids are styles.ts's; the words are
  // ours. Raidbots' own names are in the design's table and are deliberately not used:
  // "Hectic Add Cleave" says nothing about how many adds there are.
  fightStyle: 'Fight style',
  styleLabel: {
    patchwerk: 'Patchwerk',
    execute: 'Execute heavy',
    'light-movement': 'Light movement',
    'heavy-movement': 'Heavy movement',
    'cleave-2': 'Cleave, 2 targets',
    'cleave-3': 'Cleave, 3 targets',
    'cleave-5': 'Cleave, 5 targets',
    dungeon: 'Dungeon pull',
    dummy: 'Target dummy',
  } as Record<string, string>,
  /**
   * Design risk 3: "a movement window is only as honest as each rotation's handling of
   * it", so the two movement styles carry the caution until the validation job has parses
   * for them. Every other style needs no note and has none.
   */
  styleNote: {
    'light-movement':
      'How much a movement window costs depends on the rotation’s own handling of it; no parse has measured this yet.',
    'heavy-movement':
      'How much a movement window costs depends on the rotation’s own handling of it; no parse has measured this yet.',
  } as Record<string, string>,
```

- [ ] **Step 5: Write `styles.ts`**

```ts
// web/src/lib/sim/styles.ts
// The fight styles, contract 1.6. A style is a page preset: it expands to encounter fields
// and leaves its own id behind as `encounter.style`, which is a label the engine ignores.
// The table is duplicated as src/fixtures/sim/styles.json because the sim module lane owns
// a Go copy of it; styles.test.ts asserts this file and that one agree, so the fixture --
// not either implementation -- is the artefact the two lanes share.
//
// Nothing here holds a word: labels and notes are simCopy.styleLabel and simCopy.styleNote,
// keyed by these ids, so a copy change is one diff in copy.ts.
import type { EncounterSpec, Movement, TargetCount } from './types';

export type FightStyleId =
  | 'patchwerk'
  | 'execute'
  | 'light-movement'
  | 'heavy-movement'
  | 'cleave-2'
  | 'cleave-3'
  | 'cleave-5'
  | 'dungeon'
  | 'dummy';

export interface FightStyle {
  id: FightStyleId;
  targets: number;
  execute_ratio: number;
  movement: Movement | null;
  targets_over_time: TargetCount[] | null;
  dummy: boolean;
}

/** The style a fresh settings state opens on. */
export const DEFAULT_STYLE_ID: FightStyleId = 'patchwerk';

const PLAIN_EXECUTE = 0.25;

/** A 5 s window out of melee, at the interval the style names. */
function away(intervalSec: number): Movement {
  return { interval_sec: intervalSec, duration_sec: 5, kind: 'away' };
}

function cleave(id: FightStyleId, targets: number): FightStyle {
  return {
    id,
    targets,
    execute_ratio: PLAIN_EXECUTE,
    movement: null,
    targets_over_time: null,
    dummy: false,
  };
}

export const FIGHT_STYLES: readonly FightStyle[] = [
  {
    id: 'patchwerk',
    targets: 1,
    execute_ratio: PLAIN_EXECUTE,
    movement: null,
    targets_over_time: null,
    dummy: false,
  },
  {
    id: 'execute',
    targets: 1,
    execute_ratio: 0.35,
    movement: null,
    targets_over_time: null,
    dummy: false,
  },
  {
    id: 'light-movement',
    targets: 1,
    execute_ratio: PLAIN_EXECUTE,
    movement: away(45),
    targets_over_time: null,
    dummy: false,
  },
  {
    id: 'heavy-movement',
    targets: 1,
    execute_ratio: PLAIN_EXECUTE,
    movement: away(20),
    targets_over_time: null,
    dummy: false,
  },
  cleave('cleave-2', 2),
  cleave('cleave-3', 3),
  cleave('cleave-5', 5),
  {
    id: 'dungeon',
    targets: 1,
    execute_ratio: 0,
    movement: null,
    targets_over_time: [
      { at_sec: 0, count: 1 },
      { at_sec: 40, count: 3 },
      { at_sec: 80, count: 5 },
      { at_sec: 130, count: 3 },
      { at_sec: 160, count: 1 },
    ],
    dummy: false,
  },
  {
    id: 'dummy',
    targets: 1,
    execute_ratio: 0,
    movement: null,
    targets_over_time: null,
    dummy: true,
  },
];

/** The style with this id, or null. An unknown id is never silently treated as Patchwerk. */
export function fightStyle(id: string): FightStyle | null {
  return FIGHT_STYLES.find((style) => style.id === id) ?? null;
}

/**
 * The style's fields written over an encounter. Every field a style owns is written on
 * every call, including the nulls and the false: switching from Heavy movement to
 * Patchwerk has to clear the movement block, and a partial write would leave the previous
 * style's fight running under the new style's name.
 *
 * Fight length and duration variation are the player's, not the style's, so they survive.
 */
export function applyFightStyle(encounter: EncounterSpec, id: FightStyleId): EncounterSpec {
  const style = fightStyle(id);
  if (style === null) return encounter;
  return {
    ...encounter,
    style: style.id,
    targets: style.targets,
    execute_ratio: style.execute_ratio,
    movement: style.movement === null ? null : { ...style.movement },
    targets_over_time:
      style.targets_over_time === null ? null : style.targets_over_time.map((step) => ({ ...step })),
    dummy: style.dummy,
  };
}
```

- [ ] **Step 6: Run the test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/styles.test.ts
```

Expected: PASS, 11 tests.

- [ ] **Step 7: Check, lint, format**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/lib/sim/styles.ts src/lib/sim/styles.test.ts src/lib/sim/copy.ts && \
  npx prettier --check src/lib/sim/styles.ts src/lib/sim/styles.test.ts src/lib/sim/copy.ts src/fixtures/sim/styles.json
```

Expected: 0 errors.

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/web && git add src/lib/sim/styles.ts src/lib/sim/styles.test.ts src/lib/sim/copy.ts src/fixtures/sim/styles.json && \
git commit -m "feat(sim): the fight-style table, shared with the sim module as a fixture

Contract 1.6's nine styles, each expanding to targets, execute ratio,
movement, target-count timeline and the dummy flag. styles.json is the
artefact the sim module lane tests its Go table against, so neither
side can drift; applyFightStyle always writes every field a style owns
so switching styles cannot leave the previous one's movement behind.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 3: The settings model — style, fight length 20 s to 10 min, variation, target level, armor, type, dummy

Design 4.1 widens the fight-length range and exposes six controls the bar does not have. All of it is arithmetic over `EncounterSpec`, so all of it is pure and unit-tested before a component reads it.

**Files:**
- Modify: `src/lib/sim/types.ts` (`DEFAULT_ENCOUNTER` only)
- Modify: `src/lib/sim/settings.ts`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/settings.test.ts`

**Interfaces:**
- Consumes: `styles.ts`'s `applyFightStyle`, `DEFAULT_STYLE_ID`, `FightStyleId` (Task 2); `types.ts`'s `TARGET_TYPE_IDS`, `CooldownSpec` (Task 1).
- Produces:
  - `settings.ts`: `MIN_DURATION_SEC = 20`, `MAX_DURATION_SEC = 600`, `DURATIONS`, `MAX_TARGETS = 10`, `VARIATIONS`, `MAX_VARIATION = 0.3`, `TARGET_LEVELS`, `DEFAULT_TARGET_LEVEL = 63`, `TARGET_ARMOR_BY_LEVEL`, `MAX_TARGET_ARMOR = 20000`, `TARGET_TYPES`, `interface SimSettings` (+ `cooldowns: CooldownSpec[]`), `defaultSettings()`, `withStyle`, `withDuration`, `withTargets`, `withVariation`, `withTargetLevel`, `withTargetArmor`, `withTargetType`, `withDummy`, `withExecutePhase`, `withPreset`, `executePhaseOn`, `styleIdOf(settings): string`, `settingsLabel`
  - `copy.ts`: `simCopy.styleCustom`, `variation`, `variationNote`, `targetLevel`, `targetArmor`, `targetArmorPreset(armor)`, `targetType`, `targetTypeAny`, `targetTypeLabel`, `dummyTarget`, `dummyNote`, `moreSettings`

- [ ] **Step 1: Write the failing test**

Replace the whole of `src/lib/sim/settings.test.ts` with this. The first two describes are today's tests with the two assertions the new defaults change; everything after is new.

```ts
import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import {
  BUFF_PRESETS,
  DURATIONS,
  MAX_DURATION_SEC,
  MAX_TARGETS,
  MAX_TARGET_ARMOR,
  MAX_VARIATION,
  MIN_DURATION_SEC,
  TARGET_ARMOR_BY_LEVEL,
  TARGET_LEVELS,
  TARGET_TYPES,
  VARIATIONS,
  defaultSettings,
  durationLabel,
  executePhaseOn,
  settingsLabel,
  styleIdOf,
  withDummy,
  withDuration,
  withExecutePhase,
  withPreset,
  withStyle,
  withTargetArmor,
  withTargetLevel,
  withTargetType,
  withTargets,
  withVariation,
} from './settings';

describe('defaultSettings', () => {
  it("is the contract's EncounterSpec defaults, raid-buffed, on Patchwerk", () => {
    const settings = defaultSettings();
    expect(settings.encounter).toEqual({
      duration_sec: 180,
      variation: 0.2,
      targets: 1,
      execute_ratio: 0.25,
      profile: '',
      style: 'patchwerk',
      movement: null,
      targets_over_time: null,
      target_level: 63,
      target_armor: 0,
      target_type: '',
      dummy: false,
    });
    expect(settings.preset).toBe('raid-buffed');
    expect(settings.buffs.length).toBeGreaterThan(0);
    expect(settings.cooldowns).toEqual([]);
  });
});

describe('DURATIONS and durationLabel', () => {
  it('runs from twenty seconds to ten minutes', () => {
    expect(DURATIONS[0]).toBe(MIN_DURATION_SEC);
    expect(MIN_DURATION_SEC).toBe(20);
    expect(DURATIONS.at(-1)).toBe(MAX_DURATION_SEC);
    expect(MAX_DURATION_SEC).toBe(600);
    expect(DURATIONS).toContain(180);
    expect(DURATIONS.slice(0, 4)).toEqual([20, 30, 45, 60]);
    // Past a minute the step is thirty seconds all the way to ten minutes.
    const past = DURATIONS.slice(3);
    expect(past.every((seconds, i) => i === 0 || seconds - past[i - 1] === 30)).toBe(true);
  });

  it('reads as a clock, not as seconds', () => {
    expect(durationLabel(20)).toBe('0:20');
    expect(durationLabel(180)).toBe('3:00');
    expect(durationLabel(600)).toBe('10:00');
  });
});

describe('the setters never mutate and always clamp', () => {
  it('keeps duration inside twenty seconds to ten minutes', () => {
    const base = defaultSettings();
    expect(withDuration(base, 5).encounter.duration_sec).toBe(MIN_DURATION_SEC);
    expect(withDuration(base, 9999).encounter.duration_sec).toBe(MAX_DURATION_SEC);
    expect(base.encounter.duration_sec).toBe(180);
  });

  it('keeps targets between one and ten', () => {
    const base = defaultSettings();
    expect(withTargets(base, 0).encounter.targets).toBe(1);
    expect(withTargets(base, 99).encounter.targets).toBe(MAX_TARGETS);
  });

  it('keeps variation between none and thirty per cent, in five-point steps', () => {
    const base = defaultSettings();
    expect(VARIATIONS).toEqual([0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3]);
    expect(MAX_VARIATION).toBe(0.3);
    expect(withVariation(base, -1).encounter.variation).toBe(0);
    expect(withVariation(base, 5).encounter.variation).toBe(MAX_VARIATION);
    expect(withVariation(base, 0.1).encounter.variation).toBe(0.1);
  });

  it('keeps the target level between sixty and sixty-three', () => {
    const base = defaultSettings();
    expect(TARGET_LEVELS).toEqual([60, 61, 62, 63]);
    expect(withTargetLevel(base, 42).encounter.target_level).toBe(60);
    expect(withTargetLevel(base, 99).encounter.target_level).toBe(63);
    expect(withTargetLevel(base, 61).encounter.target_level).toBe(61);
  });

  it('keeps target armor non-negative and bounded, and zero means the engine’s preset', () => {
    const base = defaultSettings();
    expect(withTargetArmor(base, -10).encounter.target_armor).toBe(0);
    expect(withTargetArmor(base, 999999).encounter.target_armor).toBe(MAX_TARGET_ARMOR);
    expect(withTargetArmor(base, 3731).encounter.target_armor).toBe(3731);
  });

  it('publishes contract A8’s armor preset for each level, so the control can name the figure', () => {
    expect(TARGET_ARMOR_BY_LEVEL).toEqual({ 60: 3300, 61: 3444, 62: 3588, 63: 3731 });
    expect(TARGET_LEVELS.every((level) => TARGET_ARMOR_BY_LEVEL[level] > 0)).toBe(true);
  });

  it('accepts only the contract’s target types, and the empty string for “any”', () => {
    const base = defaultSettings();
    expect(TARGET_TYPES).toContain('undead');
    expect(withTargetType(base, 'undead').encounter.target_type).toBe('undead');
    expect(withTargetType(base, 'gnome').encounter.target_type).toBe('');
    expect(withTargetType(base, '').encounter.target_type).toBe('');
  });

  it('turns the dummy on and off', () => {
    const base = defaultSettings();
    expect(withDummy(base, true).encounter.dummy).toBe(true);
    expect(withDummy(withDummy(base, true), false).encounter.dummy).toBe(false);
  });

  it('turns the execute phase off by zeroing the ratio, and back on to the default', () => {
    const base = defaultSettings();
    const off = withExecutePhase(base, false);
    expect(off.encounter.execute_ratio).toBe(0);
    expect(executePhaseOn(off)).toBe(false);
    expect(withExecutePhase(off, true).encounter.execute_ratio).toBe(0.25);
  });

  it('swaps the whole buff and consumable list with the preset, and leaves it alone for custom', () => {
    const base = defaultSettings();
    const solo = withPreset(base, 'solo');
    expect(solo.buffs).toEqual([]);
    expect(solo.consumables).toEqual([]);
    const custom = withPreset(base, 'custom');
    expect(custom.buffs).toEqual(base.buffs);
    expect(custom.preset).toBe('custom');
  });

  it('offers exactly the three presets the design names', () => {
    expect(BUFF_PRESETS.map((row) => row.id)).toEqual(['raid-buffed', 'solo', 'custom']);
  });
});

describe('the style and the controls beside it', () => {
  it('writes the style’s fields through applyFightStyle', () => {
    const cleave = withStyle(defaultSettings(), 'cleave-5');
    expect(cleave.encounter.targets).toBe(5);
    expect(styleIdOf(cleave)).toBe('cleave-5');
  });

  it('detaches from the style the moment a style-owned field is set by hand', () => {
    const cleave = withStyle(defaultSettings(), 'cleave-3');
    expect(styleIdOf(withTargets(cleave, 4))).toBe('');
    expect(styleIdOf(withExecutePhase(cleave, false))).toBe('');
    expect(styleIdOf(withDummy(cleave, true))).toBe('');
    // Fight length, variation, target level, armor and type are the player's, not the
    // style's: changing one keeps the style's name on the run.
    expect(styleIdOf(withDuration(cleave, 300))).toBe('cleave-3');
    expect(styleIdOf(withVariation(cleave, 0))).toBe('cleave-3');
    expect(styleIdOf(withTargetLevel(cleave, 60))).toBe('cleave-3');
    expect(styleIdOf(withTargetArmor(cleave, 2000))).toBe('cleave-3');
    expect(styleIdOf(withTargetType(cleave, 'undead'))).toBe('cleave-3');
  });

  it('names the detached state', () => {
    expect(simCopy.styleCustom).toBeTruthy();
  });
});

describe('settingsLabel', () => {
  it('is the one line a saved sim is titled with, unchanged by the new controls', () => {
    expect(settingsLabel(defaultSettings())).toBe('Raid-buffed, 3:00, single target');
    expect(settingsLabel(withTargets(withPreset(defaultSettings(), 'solo'), 4))).toBe(
      'Solo, 3:00, 4 targets',
    );
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/settings.test.ts
```

Expected: FAIL — `withStyle is not a function` and the `defaultSettings` deep-equal mismatch.

- [ ] **Step 3: Widen `DEFAULT_ENCOUNTER`**

In `src/lib/sim/types.ts`, replace the existing constant:

```ts
export const DEFAULT_ENCOUNTER: EncounterSpec = {
  duration_sec: 180,
  variation: 0.2,
  targets: 1,
  execute_ratio: 0.25,
  profile: '',
  // Every field a style owns is present from the start rather than appearing the first
  // time a style is chosen: `applyFightStyle` writes all of them on every call, and an
  // encounter that sometimes carries a key and sometimes does not makes the request
  // drawer's diff (Task 15) noisy for no reason.
  style: 'patchwerk',
  movement: null,
  targets_over_time: null,
  target_level: 63,
  target_armor: 0,
  target_type: '',
  dummy: false,
};
```

- [ ] **Step 4: Add the copy**

In `src/lib/sim/copy.ts`, append inside `simCopy`:

```ts
  /** The style select's option for an encounter no style describes any more. */
  styleCustom: 'Custom',
  moreSettings: 'More settings',
  variation: 'Length varies by',
  variationNote:
    'Every iteration draws its own fight length inside this band, the way real pulls do.',
  targetLevel: 'Target level',
  targetArmor: 'Target armor',
  /**
   * Zero means the preset for the chosen level, and contract A8 publishes the figure, so
   * the empty field names it rather than leaving the player guessing what they are about
   * to override.
   */
  targetArmorPreset: (armor: string): string => `${armor}, the preset for this level`,
  targetType: 'Target type',
  targetTypeAny: 'Any',
  targetTypeLabel: {
    humanoid: 'Humanoid',
    undead: 'Undead',
    beast: 'Beast',
    demon: 'Demon',
    dragonkin: 'Dragonkin',
    elemental: 'Elemental',
    giant: 'Giant',
    mechanical: 'Mechanical',
    unknown: 'Unknown',
  } as Record<string, string>,
  dummyTarget: 'Target dummy',
  dummyNote: 'No debuffs, no execute window and no armor reduction, the way a dummy fights back.',
```

`settingsFootnote` stays in `simCopy` for now: `SettingsBar.svelte` still renders it and `tests/e2e/sim-settings.spec.ts` still asserts it, and deleting a constant a component reads would fail `astro check` in the middle of this task. Task 4 removes all three together, once `variationNote` and the variation control have replaced the number the footnote used to hard-code.

- [ ] **Step 5: Rewrite `settings.ts`**

Replace the constants block and the setters. The `BUFF_PRESETS`/`PRESET_BUFFS`/`PRESET_CONSUMABLES` block and its long comment stay exactly as they are.

```ts
import { durationLabel, encounterLabel } from './encounter';
import { applyFightStyle, DEFAULT_STYLE_ID, type FightStyleId } from './styles';
import { DEFAULT_ENCOUNTER, TARGET_TYPE_IDS, type CooldownSpec, type EncounterSpec } from './types';

export { durationLabel, encounterLabel };

// Design 4.1 and contract A3, which ratifies both numbers on the Go side too
// (`api.MinDurationSec = 20`, `api.MaxDurationSec = 600`). The old floor was a minute and
// the old ceiling eight.
export const MIN_DURATION_SEC = 20;
export const MAX_DURATION_SEC = 600;
export const DURATION_STEP_SEC = 30;
export const MAX_TARGETS = 10;

/**
 * Three short lengths for the openers and the burst windows people actually ask about,
 * then thirty-second steps to ten minutes. A uniform step from twenty seconds would put
 * twenty options under a minute, which is a select nobody can use on a phone.
 */
export const DURATIONS: readonly number[] = [
  20,
  30,
  45,
  ...Array.from(
    { length: (MAX_DURATION_SEC - 60) / DURATION_STEP_SEC + 1 },
    (_, i) => 60 + i * DURATION_STEP_SEC,
  ),
];

export const MAX_VARIATION = 0.3;
export const VARIATIONS: readonly number[] = [0, 0.05, 0.1, 0.15, 0.2, 0.25, 0.3];

export const TARGET_LEVELS: readonly number[] = [60, 61, 62, 63];
export const DEFAULT_TARGET_LEVEL = 63;

/**
 * Contract A8: the engine's own 3,731 at level 63 and a linear fall to the level-60
 * figure. `target_armor: 0` still means "the level's preset" and is what the request
 * carries by default -- these numbers exist so the override control can say what it is
 * overriding rather than showing an empty field. A better source replaces the three
 * interior numbers on the Go side and here together.
 */
export const TARGET_ARMOR_BY_LEVEL: Record<number, number> = {
  60: 3300,
  61: 3444,
  62: 3588,
  63: 3731,
};

/** A generous bound on the override field; the presets above are far below it. */
export const MAX_TARGET_ARMOR = 20_000;
export const TARGET_TYPES: readonly string[] = TARGET_TYPE_IDS;
```

The settings shape gains one field, and the defaults come from `DEFAULT_ENCOUNTER`:

```ts
export interface SimSettings {
  encounter: EncounterSpec;
  preset: BuffPresetId;
  buffs: string[];
  consumables: string[];
  /** Cooldown timing rows (contract 1.7). Empty means "everything on cooldown". */
  cooldowns: CooldownSpec[];
}

export function defaultSettings(): SimSettings {
  return {
    encounter: applyFightStyle({ ...DEFAULT_ENCOUNTER }, DEFAULT_STYLE_ID),
    preset: 'raid-buffed',
    buffs: [...PRESET_BUFFS['raid-buffed']],
    consumables: [...PRESET_CONSUMABLES['raid-buffed']],
    cooldowns: [],
  };
}
```

Then the setters. `withPreset`, `executePhaseOn` and `settingsLabel` keep their current bodies; `withExecutePhase` and `withTargets` gain the detach, and the rest are new.

```ts
function clamp(value: number, low: number, high: number): number {
  return Math.min(high, Math.max(low, value));
}

/**
 * A style owns targets, the execute ratio, movement, the target-count timeline and the
 * dummy flag. Setting any of those by hand means the encounter is no longer the style's,
 * so the label goes -- a run that says "Cleave, 3 targets" while simming four targets
 * would be a lie in the saved sim's own title.
 */
function detached(encounter: EncounterSpec): EncounterSpec {
  return { ...encounter, style: '' };
}

/** The style this encounter still is, or "" once a style-owned field was changed by hand. */
export function styleIdOf(settings: SimSettings): string {
  return settings.encounter.style ?? '';
}

export function withStyle(settings: SimSettings, id: FightStyleId): SimSettings {
  return { ...settings, encounter: applyFightStyle(settings.encounter, id) };
}

export function withDuration(settings: SimSettings, seconds: number): SimSettings {
  const clamped = clamp(Math.round(seconds), MIN_DURATION_SEC, MAX_DURATION_SEC);
  return { ...settings, encounter: { ...settings.encounter, duration_sec: clamped } };
}

export function withTargets(settings: SimSettings, targets: number): SimSettings {
  const clamped = clamp(Math.round(targets), 1, MAX_TARGETS);
  return { ...settings, encounter: { ...detached(settings.encounter), targets: clamped } };
}

export function withVariation(settings: SimSettings, variation: number): SimSettings {
  const clamped = clamp(variation, 0, MAX_VARIATION);
  return { ...settings, encounter: { ...settings.encounter, variation: clamped } };
}

export function withTargetLevel(settings: SimSettings, level: number): SimSettings {
  const clamped = clamp(Math.round(level), TARGET_LEVELS[0], TARGET_LEVELS[TARGET_LEVELS.length - 1]);
  return { ...settings, encounter: { ...settings.encounter, target_level: clamped } };
}

export function withTargetArmor(settings: SimSettings, armor: number): SimSettings {
  const clamped = clamp(Math.round(armor), 0, MAX_TARGET_ARMOR);
  return { ...settings, encounter: { ...settings.encounter, target_armor: clamped } };
}

/** An id outside the contract's vocabulary reads as "any", never as itself. */
export function withTargetType(settings: SimSettings, type: string): SimSettings {
  const known = TARGET_TYPES.includes(type) ? (type as EncounterSpec['target_type']) : '';
  return { ...settings, encounter: { ...settings.encounter, target_type: known } };
}

export function withDummy(settings: SimSettings, on: boolean): SimSettings {
  return { ...settings, encounter: { ...detached(settings.encounter), dummy: on } };
}

export function withExecutePhase(settings: SimSettings, on: boolean): SimSettings {
  return {
    ...settings,
    encounter: {
      ...detached(settings.encounter),
      execute_ratio: on ? DEFAULT_ENCOUNTER.execute_ratio : 0,
    },
  };
}
```

- [ ] **Step 6: Run the test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/settings.test.ts src/lib/sim/styles.test.ts
```

Expected: PASS.

- [ ] **Step 7: Run the whole unit suite — `defaultSettings()` is used widely**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run
```

Expected: PASS. If `store.test.ts` or `url.test.ts` assert on the old encounter shape, widen those assertions to the new one rather than narrowing `defaultSettings()`.

- [ ] **Step 8: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/lib/sim/settings.ts src/lib/sim/settings.test.ts src/lib/sim/types.ts src/lib/sim/copy.ts && \
  npx prettier --check src/lib/sim/settings.ts src/lib/sim/settings.test.ts src/lib/sim/types.ts src/lib/sim/copy.ts && \
  git add -A src/lib/sim && \
  git commit -m "feat(sim): fight length to ten minutes, variation and the target controls

Design 4.1: twenty seconds to ten minutes, duration variation 0 to 30%,
target level 60-63, a target-armor override over the engine's preset,
target type, and the dummy. A style owns targets, execute, movement,
the target-count timeline and the dummy flag, so setting any of those
by hand detaches the encounter from the style's name rather than
leaving a saved sim titled for a fight it did not run.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 4: The settings bar — the style picker and the rest, with the phone disclosure

The bar keeps its five-control primary row (style replaces nothing; execute moves down) and gains a `<details>` disclosure for the six secondary controls. `<details>`/`<summary>` is the site's own collapsible — `report/Glossary.svelte` uses exactly this — so nothing new is invented, and a disclosure that is closed on first render adds no height to reflow.

**Files:**
- Modify: `src/components/sim/SettingsBar.svelte`
- Create: `src/components/sim/SettingsSheet.svelte`
- Modify: `tests/e2e/sim-settings.spec.ts`
- Modify: `tests/e2e/sim-phone.spec.ts`

**Interfaces:**
- Consumes: `settings.ts`'s setters and constants (Task 3), `styles.ts`'s `FIGHT_STYLES` (Task 2), `simCopy`.
- Produces: `SettingsSheet.svelte` — props `{ settings: SimSettings; disabled: boolean; onchange: (next: SimSettings) => void }`. Part B mounts both components unchanged.
- Test ids added: `sim-style` (contract 9), `sim-style-note`, `sim-settings-more`, `sim-variation`, `sim-target-level`, `sim-target-armor`, `sim-target-type`, `sim-dummy`. `sim-duration`, `sim-targets`, `sim-execute`, `sim-preset`, `sim-rotation` keep their names.

- [ ] **Step 1: Write the failing e2e test**

Replace `tests/e2e/sim-settings.spec.ts` with:

```ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('the settings bar reads the defaults and every control changes the settings', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  const settings = page.getByTestId('sim-settings');
  await expect(settings).toBeVisible();

  // Defaults: defaultSettings() (settings.ts) is Patchwerk, 3:00, single target,
  // raid-buffed -- and the rotation names the spec the loaded character carries, Fury.
  await expect(page.getByTestId('sim-style')).toHaveValue('patchwerk');
  await expect(page.getByTestId('sim-duration')).toHaveValue('180');
  await expect(page.getByTestId('sim-targets')).toHaveValue('1');
  await expect(page.getByTestId('sim-preset')).toHaveValue('raid-buffed');
  await expect(page.getByTestId('sim-rotation')).toHaveText('Default for Fury');
  await expect(page.getByTestId('sim-rotation-link')).toHaveAttribute('href', '/sim/specs#warrior-fury');

  // Fight style: the contract's nine, plus nothing. Choosing one writes its fields.
  const style = page.getByTestId('sim-style');
  await expect(style.locator('option')).toHaveCount(9);
  await style.selectOption('cleave-3');
  await expect(page.getByTestId('sim-targets')).toHaveValue('3');
  await expect(page.getByTestId('sim-style-note')).toBeHidden();

  // The two movement styles carry the honesty note; nothing else does.
  await style.selectOption('light-movement');
  await expect(page.getByTestId('sim-targets')).toHaveValue('1');
  await expect(page.getByTestId('sim-style-note')).toBeVisible();

  // Setting targets by hand detaches from the style: the select falls to "Custom".
  await page.getByTestId('sim-targets').selectOption('4');
  await expect(style).toHaveValue('');

  // Fight length is labelled as a clock and reaches ten minutes.
  const duration = page.getByTestId('sim-duration');
  await expect(duration.locator('option', { hasText: '0:20' })).toHaveCount(1);
  await expect(duration.locator('option', { hasText: '10:00' })).toHaveCount(1);
  await duration.selectOption('600');
  await expect(duration).toHaveValue('600');

  // The secondary controls live in the disclosure and are closed on arrival.
  const more = page.getByTestId('sim-settings-more');
  await expect(more).toBeVisible();
  await expect(page.getByTestId('sim-variation')).toBeHidden();
  await more.getByRole('group').or(more).locator('summary').click();

  await expect(page.getByTestId('sim-variation')).toBeVisible();
  await page.getByTestId('sim-variation').selectOption('0');
  await expect(page.getByTestId('sim-variation')).toHaveValue('0');

  await page.getByTestId('sim-target-level').selectOption('60');
  await expect(page.getByTestId('sim-target-level')).toHaveValue('60');

  await page.getByTestId('sim-target-armor').fill('3731');
  await expect(page.getByTestId('sim-target-armor')).toHaveValue('3731');

  await page.getByTestId('sim-target-type').selectOption('undead');
  await expect(page.getByTestId('sim-target-type')).toHaveValue('undead');

  const execute = page.getByTestId('sim-execute');
  await execute.uncheck();
  await expect(execute).not.toBeChecked();
  await execute.check();

  const dummy = page.getByTestId('sim-dummy');
  await dummy.check();
  await expect(dummy).toBeChecked();

  // The changes hold after the source switcher is reopened and the strip is brought back.
  await page.getByTestId('sim-change-source').click();
  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(duration).toHaveValue('600');
  await expect(page.getByTestId('sim-targets')).toHaveValue('4');
});

test('the settings bar is hidden until a character is loaded', async ({ page }) => {
  await page.goto('/sim');
  await expect(page.getByTestId('sim-settings')).toBeHidden();
  await expect(page.getByTestId('sim-empty')).toBeVisible();
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && E2E_PORT=4399 npx playwright test tests/e2e/sim-settings.spec.ts --project=desktop
```

Expected: FAIL — `getByTestId('sim-style')` resolves to nothing.

- [ ] **Step 3: Write `SettingsSheet.svelte`**

```svelte
<!-- web/src/components/sim/SettingsSheet.svelte -->
<!-- The six controls that change a fight without changing what most players are asking.
     A <details> disclosure, the same collapsible report/Glossary.svelte uses: closed on
     first render, so it adds no height to the settings bar's own reservation, and opening
     it is a user gesture, which is the one kind of layout shift the CLS budget excludes.

     Variation, target level, armor and type are the player's and never touch the fight
     style's name. Execute phase and the dummy are the style's, so setting either detaches
     the encounter from its style -- settings.ts's `detached`, not this component's. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import {
    MAX_TARGET_ARMOR,
    TARGET_ARMOR_BY_LEVEL,
    TARGET_LEVELS,
    TARGET_TYPES,
    VARIATIONS,
    executePhaseOn,
    withDummy,
    withExecutePhase,
    withTargetArmor,
    withTargetLevel,
    withTargetType,
    withVariation,
    type SimSettings,
  } from '../../lib/sim/settings';

  let {
    settings,
    disabled,
    onchange,
  }: { settings: SimSettings; disabled: boolean; onchange: (next: SimSettings) => void } = $props();

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 border px-3 text-[14px] font-semibold md:min-h-9';
  const percent = (value: number): string => `${Math.round(value * 100)}%`;
  // Contract A8's figure for whichever level is chosen, so an empty armor field says what
  // the engine will use instead of nothing at all.
  const armorPreset = $derived(
    simCopy.targetArmorPreset(
      (TARGET_ARMOR_BY_LEVEL[settings.encounter.target_level ?? 63] ?? 0).toLocaleString('en-US'),
    ),
  );
</script>

<details class="border-line-soft rounded-panel border" data-testid="sim-settings-more">
  <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-3 md:min-h-9">
    {simCopy.moreSettings}
  </summary>

  <div class="grid grid-cols-2 gap-3 p-3 pt-0 md:grid-cols-4">
    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.variation}</span>
      <select
        class={control}
        {disabled}
        title={simCopy.variationNote}
        value={String(settings.encounter.variation)}
        onchange={(event) => onchange(withVariation(settings, Number(event.currentTarget.value)))}
        data-testid="sim-variation"
      >
        {#each VARIATIONS as value (value)}
          <option value={String(value)}>{percent(value)}</option>
        {/each}
      </select>
    </label>

    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.targetLevel}</span>
      <select
        class={control}
        {disabled}
        value={String(settings.encounter.target_level ?? 63)}
        onchange={(event) => onchange(withTargetLevel(settings, Number(event.currentTarget.value)))}
        data-testid="sim-target-level"
      >
        {#each TARGET_LEVELS as level (level)}
          <option value={String(level)}>{level}</option>
        {/each}
      </select>
    </label>

    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.targetArmor}</span>
      <input
        type="number"
        min="0"
        max={MAX_TARGET_ARMOR}
        step="1"
        class={control}
        {disabled}
        placeholder={armorPreset}
        title={armorPreset}
        value={settings.encounter.target_armor === 0 ? '' : String(settings.encounter.target_armor)}
        onchange={(event) => onchange(withTargetArmor(settings, Number(event.currentTarget.value)))}
        data-testid="sim-target-armor"
      />
    </label>

    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.targetType}</span>
      <select
        class={control}
        {disabled}
        value={settings.encounter.target_type ?? ''}
        onchange={(event) => onchange(withTargetType(settings, event.currentTarget.value))}
        data-testid="sim-target-type"
      >
        <option value="">{simCopy.targetTypeAny}</option>
        {#each TARGET_TYPES as id (id)}
          <option value={id}>{simCopy.targetTypeLabel[id] ?? id}</option>
        {/each}
      </select>
    </label>

    <label class="flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="accent-gold h-5 w-5"
        {disabled}
        checked={executePhaseOn(settings)}
        onchange={(event) => onchange(withExecutePhase(settings, event.currentTarget.checked))}
        data-testid="sim-execute"
      />
      <span class="text-muted">{simCopy.executePhase}</span>
    </label>

    <label class="flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="accent-gold h-5 w-5"
        {disabled}
        title={simCopy.dummyNote}
        checked={settings.encounter.dummy === true}
        onchange={(event) => onchange(withDummy(settings, event.currentTarget.checked))}
        data-testid="sim-dummy"
      />
      <span class="text-muted">{simCopy.dummyTarget}</span>
    </label>
  </div>
</details>
```

- [ ] **Step 4: Rewrite the top of `SettingsBar.svelte`**

Replace the script's imports and the derived block:

```ts
  import { simCopy } from '../../lib/sim/copy';
  import {
    BUFF_PRESETS,
    DURATIONS,
    MAX_TARGETS,
    durationLabel,
    styleIdOf,
    withDuration,
    withPreset,
    withStyle,
    withTargets,
    type BuffPresetId,
    type SimSettings,
  } from '../../lib/sim/settings';
  import { FIGHT_STYLES, type FightStyleId } from '../../lib/sim/styles';
  import { specRow } from '../../lib/sim/spec-label';
  import SettingsSheet from './SettingsSheet.svelte';

  let {
    settings,
    spec,
    disabled,
    onchange,
  }: { settings: SimSettings; spec: string; disabled: boolean; onchange: (next: SimSettings) => void } =
    $props();

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 md:min-w-[7rem] border px-3 text-[14px] font-semibold md:min-h-9';
  const targets = Array.from({ length: MAX_TARGETS }, (_, i) => i + 1);
  const rotationName = $derived(specRow(spec)?.name ?? spec);
  const rotationLabel = $derived(`${simCopy.rotationPrefix} ${rotationName}`);
  const styleId = $derived(styleIdOf(settings));
  // Only the two movement styles carry a note (copy.ts's styleNote). Everything else
  // renders nothing at all rather than an empty paragraph that would reserve a line.
  const styleNote = $derived(simCopy.styleNote[styleId] ?? '');
</script>
```

Then in the template, put the style select first inside the existing `<div class="flex flex-wrap items-end gap-4 md:gap-5">`, replace the execute-phase `<label>` with nothing (it moved to the sheet), and replace the footnote paragraph:

```svelte
    <label class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.fightStyle}</span>
      <select
        class={control}
        {disabled}
        value={styleId}
        onchange={(event) => onchange(withStyle(settings, event.currentTarget.value as FightStyleId))}
        data-testid="sim-style"
      >
        <!-- The empty option exists only while the encounter has been detached from a
             style by hand (settings.ts's `detached`); it is never a thing to choose, so it
             is hidden the rest of the time rather than offered as a tenth style. -->
        {#if styleId === ''}
          <option value="">{simCopy.styleCustom}</option>
        {/if}
        {#each FIGHT_STYLES as style (style.id)}
          <option value={style.id}>{simCopy.styleLabel[style.id] ?? style.id}</option>
        {/each}
      </select>
    </label>
```

and, replacing `<p … data-testid="sim-settings-footnote">`:

```svelte
  {#if styleNote !== ''}
    <p class="text-muted text-[12px]" data-testid="sim-style-note">{styleNote}</p>
  {/if}

  <SettingsSheet {settings} {disabled} {onchange} />
```

Now delete `settingsFootnote` from `simCopy` (Task 3 left it in place because this component still read it). Nothing renders or asserts it any more: the rewritten spec in Step 1 dropped its assertion, and `variationNote` on the variation control states the number it used to hard-code.

- [ ] **Step 5: Run the e2e test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && E2E_PORT=4399 npx playwright test tests/e2e/sim-settings.spec.ts --project=desktop
```

Expected: PASS, 2 tests.

- [ ] **Step 6: Fix the phone audit**

`tests/e2e/sim-phone.spec.ts` sweeps every control for a 44px hit target and for horizontal scroll. Open the disclosure before the sweep of the loaded-character state so the sheet's own controls are measured, by adding this immediately after the spec's settings bar becomes visible:

```ts
  // The secondary encounter controls only exist while the disclosure is open; the audit
  // has to measure them, so it opens it the way a player does.
  await page.getByTestId('sim-settings-more').locator('summary').click();
  await expect(page.getByTestId('sim-variation')).toBeVisible();
```

Then run it:

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && E2E_PORT=4399 npx playwright test tests/e2e/sim-phone.spec.ts --project=mobile
```

Expected: PASS. If the four-column grid overflows at 412px, it is already `grid-cols-2` there — a failure means a `<select>`'s longest option is wider than half the gutter-inset row, and the fix is `min-w-0` on that label (already present) plus `truncate` on its `<span class="label">`, never a wider grid.

- [ ] **Step 7: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/components/sim/SettingsBar.svelte src/components/sim/SettingsSheet.svelte tests/e2e/sim-settings.spec.ts tests/e2e/sim-phone.spec.ts && \
  npx prettier --check src/components/sim/SettingsBar.svelte src/components/sim/SettingsSheet.svelte tests/e2e/sim-settings.spec.ts tests/e2e/sim-phone.spec.ts && \
  git add -A src/components/sim tests/e2e src/lib/sim/copy.ts && \
  git commit -m "feat(sim): the fight-style picker and the settings disclosure

The style select writes the contract's encounter fields and falls to
Custom the moment targets, execute or the dummy is set by hand. The six
secondary controls -- variation, target level, armor, type, execute and
the dummy -- live in a <details> disclosure, the same collapsible the
report's glossary uses, closed on arrival so it reserves no height.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 5: The three wasm exports of contract 10.2 — `simNeedsMore`, `simValidate`, `simCount`

None of the three decisions belongs in TypeScript. Whether a target-error run needs another step is the engine's arithmetic (`api.NeedsMoreIterations`); whether an edited request is legal is `api.SimRequest.Validate`; how many combinations a bulk request expands to is `sim/bulk`'s count. All three reach the page as wasm exports, routed to worker 0 the way `simSplit` and `simCombine` already are, because the main thread holds no engine instance.

`simCount` is part B's — `/sim/gear` and `/sim/drops` show the live combination count and the cap notice with it — but it lands here, in the one task that widens `EngineModule`, the pool's message protocol, the real worker, the fake worker and the fake engine. Adding it later would mean editing the same five files a second time for one more message.

**Files:**
- Modify: `src/lib/sim/engine.ts`
- Modify: `src/lib/sim/worker.ts`
- Modify: `src/lib/sim/sim.worker.ts`
- Modify: `src/test-support/fake-worker.ts`
- Modify: `src/fixtures/sim/engine-fake.ts`
- Test: `src/fixtures/sim/engine-fake.test.ts`, `src/lib/sim/worker.test.ts`

**Interfaces:**
- Consumes: `types.ts` (Task 1).
- Produces:
  - `engine.ts`: `EngineModule` (+ `simNeedsMore(resultJSON: string, requestJSON: string): string`, `simValidate(requestJSON: string): string`, `simCount(requestJSON: string): string`), `interface RequestValidation { ok: boolean; errors: RequestValidationError[] }`, `interface RequestValidationError { field: string; message: string }`, `type CountAnswer = { ok: true; combinations: number } | { ok: false; cap: number; combinations: number }`
  - `worker.ts`: `SimPool` (+ `needsMore(resultJSON: string, requestJSON: string): Promise<boolean>`, `validate(requestJSON: string): Promise<RequestValidation>`, `count(requestJSON: string): Promise<CountAnswer>`), `ToWorker` (+ `{ kind: 'needsMore' | 'validate' | 'count' }` variants)

**Export shapes, contract 10.2 verbatim:**

```
simNeedsMore(resultJSON, requestJSON) -> {"needs_more": true}          | {"error": "..."}
simValidate(requestJSON)              -> {"ok": false, "errors": [{"field":"iterations","message":"..."}]}
                                      | {"error": "..."}
simCount(requestJSON)                 -> {"combinations": 38}
                                      | {"error":"cap_exceeded","cap":400,"combinations":912}
```

All three are synchronous. `simNeedsMore` and `simValidate` fail the `errorJSON` way `simSplit` does, so `unwrapOrThrow` covers them. `simCount` does not go through `unwrapOrThrow`: a cap breach is not an exception, it is an answer with two numbers in it that the cap notice renders, and throwing away `cap` and `combinations` to raise `Error("cap_exceeded")` would lose exactly the part part B needs.

- [ ] **Step 1: Write the failing fake-engine test**

Append to `src/fixtures/sim/engine-fake.test.ts`:

```ts
describe('simNeedsMore', () => {
  const request = (over: Partial<SimRequest> = {}): string =>
    JSON.stringify({ ...fixtureRequest, iterations: 30_000, target_error: 0.005, ...over });
  const result = (mean: number, error: number, iterationsRun: number): string =>
    JSON.stringify({
      ...fixtureResult,
      dps: { mean, stddev: 100, error, min: 0, max: 0 },
      iterations_run: iterationsRun,
    });

  it('asks for more while the relative error is over the target', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 20, 2000), request()))).toEqual({
      needs_more: true,
    });
  });

  it('stops once the relative error is inside the target', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 4, 2000), request()))).toEqual({
      needs_more: false,
    });
  });

  it('stops at the ceiling however wide the band still is', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(1000, 99, 30_000), request()))).toEqual({
      needs_more: false,
    });
  });

  it('never asks for more on a fixed-count run', () => {
    const engine = createFakeEngine();
    expect(
      JSON.parse(engine.simNeedsMore(result(1000, 99, 500), request({ target_error: 0 }))),
    ).toEqual({ needs_more: false });
  });

  it('stops rather than dividing by a mean of zero', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simNeedsMore(result(0, 0, 1000), request()))).toEqual({
      needs_more: false,
    });
  });
});

describe('simValidate', () => {
  it('passes the fixture request', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simValidate(JSON.stringify(fixtureRequest)))).toEqual({
      ok: true,
      errors: [],
    });
  });

  it('names every field it refuses, so the drawer can put each one beside its line', () => {
    const engine = createFakeEngine();
    const broken = { ...fixtureRequest, spec: '', iterations: 0 };
    const answer = JSON.parse(engine.simValidate(JSON.stringify(broken))) as {
      ok: boolean;
      errors: { field: string; message: string }[];
    };
    expect(answer.ok).toBe(false);
    expect(answer.errors.map((row) => row.field).sort()).toEqual(['iterations', 'spec']);
    expect(answer.errors.every((row) => row.message.length > 0)).toBe(true);
  });

  it('answers the error envelope for a string that is not JSON at all', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simValidate('{nope'))).toHaveProperty('error');
  });
});

describe('simCount', () => {
  const bulk = (candidates: number, cap = 400): string =>
    JSON.stringify({
      ...fixtureRequest,
      bulk: {
        mode: 'gear',
        candidates: Array.from({ length: candidates }, (_, i) => ({
          slot: 'head',
          item_id: 16963 + i,
          origin: 'bag',
        })),
        precision: 'normal',
        cap,
      },
    });

  it('counts the combinations a bulk request expands to', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simCount(bulk(3)))).toEqual({ combinations: 3 });
  });

  it('answers cap_exceeded with both numbers rather than throwing them away', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simCount(bulk(5, 4)))).toEqual({
      error: 'cap_exceeded',
      cap: 4,
      combinations: 5,
    });
  });

  it('counts a request with no bulk block as no combinations at all', () => {
    const engine = createFakeEngine();
    expect(JSON.parse(engine.simCount(JSON.stringify(fixtureRequest)))).toEqual({ combinations: 0 });
  });
});
```

Add the two fixture consts at the top of the file if they are not already there:

```ts
const fixtureResult = fixtureResultJson as unknown as SimResult;
const fixtureRequest = fixtureResult.request;
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/fixtures/sim/engine-fake.test.ts
```

Expected: FAIL — `engine.simNeedsMore is not a function`.

- [ ] **Step 3: Widen the `EngineModule` interface**

In `src/lib/sim/engine.ts`, add to `EngineModule`:

```ts
  /**
   * Whether a target-error run has another step to do. The decision is the engine's, not
   * the page's: the page never computes a stopping rule over an estimate it did not pool.
   * Returns `{"needs_more": boolean}`, or the `{"error": …}` envelope.
   *
   * NOT in contract section 4 as written; see this task's note. The Go original is
   * `api.NeedsMoreIterations`.
   */
  simNeedsMore(resultJSON: string, requestJSON: string): string;
  /**
   * `api.SimRequest.Validate`, for the request drawer. Returns
   * `{"ok": boolean, "errors": [{"field", "message"}]}`, or the `{"error": …}` envelope for
   * a string that is not a request at all.
   */
  simValidate(requestJSON: string): string;
  /**
   * How many combinations a bulk request expands to, without allocating the requests
   * (contract 10.2). Part B's live combination count and cap notice. A breach answers
   * `{"error":"cap_exceeded","cap":n,"combinations":n}`, which is an answer and not a
   * failure, so this one is NOT passed through `unwrapOrThrow`.
   */
  simCount(requestJSON: string): string;
```

and, beside it:

```ts
export interface RequestValidationError {
  /** The request's own JSON path, e.g. "iterations" or "encounter.targets". */
  field: string;
  /** The engine's own sentence, shown verbatim. */
  message: string;
}

export interface RequestValidation {
  ok: boolean;
  errors: RequestValidationError[];
}

/** `simCount`'s two answers, both of them ordinary (contract 10.2). */
export type CountAnswer =
  | { ok: true; combinations: number }
  | { ok: false; cap: number; combinations: number };
```

Add the two to `WasmGlobals` and to the object `loadWasmEngine` returns:

```ts
  simNeedsMore?: (resultJSON: string, requestJSON: string) => string;
  simValidate?: (requestJSON: string) => string;
  simCount?: (requestJSON: string) => string;
```

```ts
    simNeedsMore: (resultJSON, requestJSON) =>
      unwrapOrThrow(globals.simNeedsMore!(resultJSON, requestJSON)),
    simValidate: (requestJSON) => unwrapOrThrow(globals.simValidate!(requestJSON)),
    // No unwrapOrThrow: `cap_exceeded` carries two numbers the page renders.
    simCount: (requestJSON) => globals.simCount!(requestJSON),
```

- [ ] **Step 4: Implement both in the fake engine**

In `src/fixtures/sim/engine-fake.ts`, add to the returned object:

```ts
    simNeedsMore(resultJSON, requestJSON) {
      try {
        const result = JSON.parse(resultJSON) as SimResult;
        const request = JSON.parse(requestJSON) as SimRequest;
        const target = request.target_error ?? 0;
        // Four ways to be done, and the real engine agrees on all four: this was not a
        // target-error run; nothing has been measured, so a relative error is undefined;
        // the band is inside the target; or the ceiling is reached.
        const needsMore =
          target > 0 &&
          result.dps.mean > 0 &&
          result.dps.error / result.dps.mean > target &&
          result.iterations_run < request.iterations;
        return JSON.stringify({ needs_more: needsMore });
      } catch (error) {
        return JSON.stringify({ error: error instanceof Error ? error.message : String(error) });
      }
    },

    simValidate(requestJSON) {
      let request: SimRequest;
      try {
        request = JSON.parse(requestJSON) as SimRequest;
      } catch (error) {
        return JSON.stringify({ error: error instanceof Error ? error.message : String(error) });
      }
      // A short stand-in for api.SimRequest.Validate: the checks the drawer's own tests
      // exercise. The real wasm runs the whole thing, and the drawer renders whatever
      // fields come back, so a fake that refuses fewer things cannot make the page wrong.
      const errors: { field: string; message: string }[] = [];
      if (request.engine_version === undefined || request.engine_version === '') {
        errors.push({ field: 'engine_version', message: 'engine_version is required' });
      }
      if (request.spec === undefined || request.spec === '') {
        errors.push({ field: 'spec', message: 'spec is required' });
      }
      if (!(request.iterations > 0)) {
        errors.push({ field: 'iterations', message: 'iterations must be a positive number' });
      }
      const targets = request.encounter?.targets;
      if (targets !== undefined && (targets < 1 || targets > 10)) {
        errors.push({ field: 'encounter.targets', message: 'targets must be between 1 and 10' });
      }
      return JSON.stringify({ ok: errors.length === 0, errors });
    },

    simCount(requestJSON) {
      let request: SimRequest;
      try {
        request = JSON.parse(requestJSON) as SimRequest;
      } catch (error) {
        return JSON.stringify({ error: error instanceof Error ? error.message : String(error) });
      }
      const bulk = request.bulk;
      if (bulk === undefined) return JSON.stringify({ combinations: 0 });
      // A stand-in for `sim/bulk`'s own expansion, which knows slot fit, unique-equipped
      // and weapon shapes; the real wasm counts properly. The fake counts what it can see
      // -- one combination per candidate, per talent loadout, per set -- which is enough
      // for the page's cap notice and its e2e to be exercised honestly.
      const combinations =
        (bulk.candidates?.length ?? 0) + (bulk.talents?.length ?? 0) + (bulk.sets?.length ?? 0);
      if (bulk.cap > 0 && combinations > bulk.cap) {
        return JSON.stringify({ error: 'cap_exceeded', cap: bulk.cap, combinations });
      }
      return JSON.stringify({ combinations });
    },
```

- [ ] **Step 5: Run the fake-engine test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/fixtures/sim/engine-fake.test.ts
```

Expected: PASS.

- [ ] **Step 6: Write the failing pool test**

Append to `src/lib/sim/worker.test.ts`:

```ts
describe('the pool routes the two synchronous exports to worker 0', () => {
  it('answers needsMore as a boolean', async () => {
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => createFakeWorker(createFakeEngine()) });
    const request = JSON.stringify({ ...fixtureResult.request, iterations: 30_000, target_error: 0.005 });
    const wide = JSON.stringify({
      ...fixtureResult,
      dps: { mean: 1000, stddev: 100, error: 40, min: 0, max: 0 },
      iterations_run: 1000,
    });
    const narrow = JSON.stringify({
      ...fixtureResult,
      dps: { mean: 1000, stddev: 100, error: 1, min: 0, max: 0 },
      iterations_run: 1000,
    });
    expect(await pool.needsMore(wide, request)).toBe(true);
    expect(await pool.needsMore(narrow, request)).toBe(false);
    pool.terminate();
  });

  it('answers validate as the parsed envelope', async () => {
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => createFakeWorker(createFakeEngine()) });
    const good = await pool.validate(JSON.stringify(fixtureResult.request));
    expect(good).toEqual({ ok: true, errors: [] });
    const bad = await pool.validate(JSON.stringify({ ...fixtureResult.request, spec: '' }));
    expect(bad.ok).toBe(false);
    expect(bad.errors[0].field).toBe('spec');
    pool.terminate();
  });

  it('answers count as a discriminated result, cap breach included', async () => {
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => createFakeWorker(createFakeEngine()) });
    const bulk = (cap: number): string =>
      JSON.stringify({
        ...fixtureResult.request,
        bulk: {
          mode: 'gear',
          candidates: [
            { slot: 'head', item_id: 16963, origin: 'bag' },
            { slot: 'head', item_id: 16964, origin: 'bag' },
          ],
          precision: 'normal',
          cap,
        },
      });
    expect(await pool.count(bulk(400))).toEqual({ ok: true, combinations: 2 });
    expect(await pool.count(bulk(1))).toEqual({ ok: false, cap: 1, combinations: 2 });
    pool.terminate();
  });
});
```

(`fixtureResult` is the file's existing import of `src/fixtures/sim/result.json`; add it if the file does not already have one.)

- [ ] **Step 7: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/worker.test.ts
```

Expected: FAIL — `pool.needsMore is not a function`.

- [ ] **Step 8: Extend the protocol and the pool**

In `src/lib/sim/worker.ts`, add two `ToWorker` variants (nothing is added to `FromWorker`: both answers are one JSON string, which `{kind: 'one'}` already carries):

```ts
export type ToWorker =
  | { kind: 'split'; token: number; request: string; shards: number }
  | { kind: 'run'; token: number; callbackId: string; request: string }
  | { kind: 'combine'; token: number; results: string[] }
  | { kind: 'needsMore'; token: number; result: string; request: string }
  | { kind: 'validate'; token: number; request: string }
  | { kind: 'count'; token: number; request: string }
  | { kind: 'abort'; callbackId: string };
```

Add the two methods to `SimPool`:

```ts
export interface SimPool {
  readonly size: number;
  split(request: string, shards: number): Promise<string[]>;
  run(
    shards: readonly string[],
    callbackId: string,
    onProgress: (progress: ShardProgress) => void,
  ): Promise<string[]>;
  combine(results: readonly string[]): Promise<string>;
  /** Whether a target-error run has another step to do. The engine decides, not the page. */
  needsMore(result: string, request: string): Promise<boolean>;
  /** `api.SimRequest.Validate` over an edited request. */
  validate(request: string): Promise<RequestValidation>;
  /** How many combinations a bulk request expands to; a cap breach is an answer, not a throw. */
  count(request: string): Promise<CountAnswer>;
  abort(callbackId: string): void;
  terminate(): void;
}
```

and implement them beside `combine` in the returned object:

```ts
    async needsMore(result, request) {
      const answer = await send<string>(0, (token) => ({ kind: 'needsMore', token, result, request }));
      return (JSON.parse(answer) as { needs_more?: boolean }).needs_more === true;
    },
    async validate(request) {
      const answer = await send<string>(0, (token) => ({ kind: 'validate', token, request }));
      return JSON.parse(answer) as RequestValidation;
    },
    async count(request) {
      const answer = await send<string>(0, (token) => ({ kind: 'count', token, request }));
      const parsed = JSON.parse(answer) as {
        error?: string;
        cap?: number;
        combinations?: number;
      };
      // `cap_exceeded` is the one error envelope on this lane that is a real answer; any
      // other error from the engine is a genuine failure and is thrown as one.
      if (parsed.error === 'cap_exceeded') {
        return { ok: false, cap: parsed.cap ?? 0, combinations: parsed.combinations ?? 0 };
      }
      if (parsed.error !== undefined) throw new Error(parsed.error);
      return { ok: true, combinations: parsed.combinations ?? 0 };
    },
```

with `import type { CountAnswer, RequestValidation } from './engine';` at the top.

- [ ] **Step 9: Handle both in the real worker and the fake one**

In `src/lib/sim/sim.worker.ts`, add two branches to `handle()`, after the `combine` branch and before the `run` fall-through:

```ts
    if (message.kind === 'needsMore') {
      reply({ kind: 'one', token, result: loaded.simNeedsMore(message.result, message.request) });
      return;
    }
    if (message.kind === 'validate') {
      reply({ kind: 'one', token, result: loaded.simValidate(message.request) });
      return;
    }
    if (message.kind === 'count') {
      reply({ kind: 'one', token, result: loaded.simCount(message.request) });
      return;
    }
```

In `src/test-support/fake-worker.ts`, the same two, inside the `postMessage` chain between the `combine` branch and the `else`:

```ts
          } else if (message.kind === 'needsMore') {
            emit({ kind: 'one', token, result: engine.simNeedsMore(message.result, message.request) });
          } else if (message.kind === 'validate') {
            emit({ kind: 'one', token, result: engine.simValidate(message.request) });
          } else if (message.kind === 'count') {
            emit({ kind: 'one', token, result: engine.simCount(message.request) });
```

`createBrokenWorker` needs no change: it already answers every non-abort message with `{kind: 'failed'}`.

- [ ] **Step 10: Run the pool test and the whole unit suite**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run
```

Expected: PASS.

- [ ] **Step 11: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/lib/sim/engine.ts src/lib/sim/worker.ts src/lib/sim/sim.worker.ts src/test-support/fake-worker.ts src/fixtures/sim/engine-fake.ts src/fixtures/sim/engine-fake.test.ts src/lib/sim/worker.test.ts && \
  npx prettier --check src/lib/sim/engine.ts src/lib/sim/worker.ts src/lib/sim/sim.worker.ts src/test-support/fake-worker.ts src/fixtures/sim/engine-fake.ts src/fixtures/sim/engine-fake.test.ts src/lib/sim/worker.test.ts && \
  git add -A src/lib/sim src/test-support src/fixtures/sim && \
  git commit -m "feat(sim): the three exports of contract 10.2, through the pool

simNeedsMore, simValidate and simCount on worker 0, beside simSplit and
simCombine. Whether a target-error run has another step, whether an
edited request is legal and how many combinations a bulk request
expands to are all the engine's decisions; none is reimplemented in
TypeScript. simCount's cap_exceeded keeps its two numbers rather than
being thrown as an Error, because the cap notice renders both.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 6: Precision — fast, normal, high, and "until ±0.5%"

Design 4.2. The three fixed counts plus the step loop, with the lane's ceiling. `run.ts` grows the loop; the store grows one field.

Two rulings bound it. **Contract A3**: a plain fixed run keeps the closed `ValidIterations` set, and a `TargetError` run's `Iterations` is instead "a positive multiple of `StepIterations` at or under `LaneIterationCeiling[lane]`", with `Validate` applying the rule for the request's kind — so nothing here needs a workaround for the closed set, and both ceilings are multiples of the step by construction. **Contract 10.2**: stage requests run unsplit, and `simSplit`/`simCombine` are for plain runs only — a target-error run *is* a plain run, so each of its steps still splits across the pool exactly as today's single run does. Only part B's bulk stages bypass the splitter.

**Files:**
- Create: `src/lib/sim/precision.ts`
- Modify: `src/lib/sim/run.ts`
- Modify: `src/lib/sim/store.svelte.ts`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/precision.test.ts`, `src/lib/sim/run.test.ts`, `src/lib/sim/store.test.ts`

**Interfaces:**
- Consumes: `worker.ts`'s `SimPool.needsMore` (Task 5), `types.ts`'s `STEP_ITERATIONS_DEFAULT` (Task 1).
- Produces:
  - `precision.ts`: `type PrecisionId = 'fast' | 'normal' | 'high' | 'target-error'`, `PRECISIONS: readonly PrecisionId[]`, `BULK_PRECISIONS: readonly ('fast'|'normal'|'high')[]`, `PRECISION_ITERATIONS: Record<'fast'|'normal'|'high', number>`, `TARGET_ERROR = 0.005`, `STEP_ITERATIONS`, `LANE_ITERATION_CEILING: Record<'browser'|'server', number>`, `CAPS: Record<'browser'|'server', number>` (contract A2), `interface PrecisionPlan { iterations: number; targetError: number; step: number }`, `precisionPlan(id, lane)`, `precisionOf(request)`, `relativeError(estimate)`
  - `run.ts`: `RunInput` (+ `targetError?`, `stepIterations?`, `pool` unchanged), `RunUpdate` (+ `relativeError: number`)
  - `store.svelte.ts`: `precisionId` getter, `setPrecisionId(id: PrecisionId)`, `lane` getter

- [ ] **Step 1: Write the failing precision test**

Create `src/lib/sim/precision.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import {
  BULK_PRECISIONS,
  CAPS,
  LANE_ITERATION_CEILING,
  PRECISIONS,
  PRECISION_ITERATIONS,
  STEP_ITERATIONS,
  TARGET_ERROR,
  precisionOf,
  precisionPlan,
  relativeError,
} from './precision';
import type { SimRequest } from './types';

describe('the precision vocabulary', () => {
  it('is the design’s three counts plus the target-error run', () => {
    expect(PRECISIONS).toEqual(['fast', 'normal', 'high', 'target-error']);
    expect(PRECISION_ITERATIONS).toEqual({ fast: 500, normal: 3000, high: 10_000 });
  });

  it('offers only the three fixed counts to a bulk request, which the contract’s BulkSpec takes', () => {
    expect(BULK_PRECISIONS).toEqual(['fast', 'normal', 'high']);
  });

  it('carries contract A2’s combination caps, the server one at five thousand', () => {
    expect(CAPS).toEqual({ browser: 400, server: 5000 });
  });

  it('carries the contract’s half a per cent, thousand-iteration step and lane ceilings', () => {
    expect(TARGET_ERROR).toBe(0.005);
    expect(STEP_ITERATIONS).toBe(1000);
    expect(LANE_ITERATION_CEILING).toEqual({ browser: 30_000, server: 100_000 });
  });

  it('keeps every ceiling a positive multiple of the step, which is contract A3’s rule', () => {
    for (const ceiling of Object.values(LANE_ITERATION_CEILING)) {
      expect(ceiling).toBeGreaterThan(0);
      expect(ceiling % STEP_ITERATIONS).toBe(0);
    }
  });
});

describe('precisionPlan', () => {
  it('turns a fixed precision into a count with no target and no step', () => {
    expect(precisionPlan('normal', 'browser')).toEqual({ iterations: 3000, targetError: 0, step: 0 });
    expect(precisionPlan('high', 'server')).toEqual({ iterations: 10_000, targetError: 0, step: 0 });
  });

  it('turns the target-error run into the lane’s ceiling, the target and the step', () => {
    expect(precisionPlan('target-error', 'browser')).toEqual({
      iterations: 30_000,
      targetError: 0.005,
      step: 1000,
    });
    expect(precisionPlan('target-error', 'server')).toEqual({
      iterations: 100_000,
      targetError: 0.005,
      step: 1000,
    });
  });
});

describe('precisionOf', () => {
  const request = (over: Partial<SimRequest>): SimRequest => ({ ...over } as SimRequest);

  it('reads a stored request back as the precision that produced it', () => {
    expect(precisionOf(request({ iterations: 500 }))).toBe('fast');
    expect(precisionOf(request({ iterations: 3000 }))).toBe('normal');
    expect(precisionOf(request({ iterations: 10_000 }))).toBe('high');
    expect(precisionOf(request({ iterations: 30_000, target_error: 0.005 }))).toBe('target-error');
  });

  it('falls back to normal for a count no precision names', () => {
    expect(precisionOf(request({ iterations: 1234 }))).toBe('normal');
  });
});

describe('relativeError', () => {
  it('is the error over the mean, and zero rather than infinity at a mean of zero', () => {
    expect(relativeError({ mean: 1000, stddev: 0, error: 5, min: 0, max: 0 })).toBeCloseTo(0.005);
    expect(relativeError({ mean: 0, stddev: 0, error: 5, min: 0, max: 0 })).toBe(0);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/precision.test.ts
```

Expected: FAIL — `Failed to resolve import "./precision"`.

- [ ] **Step 3: Write `precision.ts`**

```ts
// web/src/lib/sim/precision.ts
// Design 4.2 and contract 1.2. Three fixed counts and one stopping rule.
//
// "Until ±0.5%" is Raidbots' Smart Sim, made visible: the run continues in thousand-
// iteration steps until the DPS error is inside half a per cent, or until the lane's
// ceiling is reached, and the results line says which of the two stopped it. The decision
// itself is never taken here -- `SimPool.needsMore` asks the engine after every step (see
// run.ts). This module only carries the numbers.
//
// Contract A3 requires a target-error run's `Iterations` to be a positive multiple of
// `StepIterations` at or under the lane's ceiling; both ceilings below satisfy that, and
// precision.test.ts asserts it so a future ceiling cannot quietly break `Validate`.
import type { Estimate, SimRequest } from './types';
import { STEP_ITERATIONS_DEFAULT } from './types';

export type Lane = 'browser' | 'server';

export type PrecisionId = 'fast' | 'normal' | 'high' | 'target-error';

export const PRECISIONS: readonly PrecisionId[] = ['fast', 'normal', 'high', 'target-error'];

/** `BulkSpec.Precision` takes only the three fixed counts (contract 1.3). */
export const BULK_PRECISIONS: readonly Exclude<PrecisionId, 'target-error'>[] = [
  'fast',
  'normal',
  'high',
];

export const PRECISION_ITERATIONS: Record<Exclude<PrecisionId, 'target-error'>, number> = {
  fast: 500,
  normal: 3000,
  high: 10_000,
};

/** Half a per cent, which is what the control is named after. */
export const TARGET_ERROR = 0.005;

export const STEP_ITERATIONS = STEP_ITERATIONS_DEFAULT;

export const LANE_ITERATION_CEILING: Record<Lane, number> = { browser: 30_000, server: 100_000 };

/**
 * Contract A2: how many combinations a bulk request may expand to on each lane. The
 * server figure is 5,000, not the design's 20,000 -- at the measured native rate a
 * 20,000-combination fast run cannot finish inside the job's fifteen minutes.
 *
 * Nothing on `/sim` uses this; it lives beside the iteration ceilings because they are
 * the same kind of fact, and part B's cap notice would otherwise keep a second copy.
 */
export const CAPS: Record<Lane, number> = { browser: 400, server: 5000 };

export interface PrecisionPlan {
  /** The count for a fixed run, the ceiling for a target-error run. */
  iterations: number;
  /** 0 for a fixed run. */
  targetError: number;
  /** 0 for a fixed run. */
  step: number;
}

export function precisionPlan(id: PrecisionId, lane: Lane): PrecisionPlan {
  if (id === 'target-error') {
    return { iterations: LANE_ITERATION_CEILING[lane], targetError: TARGET_ERROR, step: STEP_ITERATIONS };
  }
  return { iterations: PRECISION_ITERATIONS[id], targetError: 0, step: 0 };
}

/**
 * A stored request read back as the control that produced it, for a saved sim and for a
 * request pasted into the drawer. A count no precision names -- a hand-edited request --
 * reads as `normal`: the control has to show something, and the request itself, not the
 * control, is what will be re-run.
 */
export function precisionOf(request: Pick<SimRequest, 'iterations' | 'target_error'>): PrecisionId {
  if ((request.target_error ?? 0) > 0) return 'target-error';
  const found = (Object.keys(PRECISION_ITERATIONS) as Exclude<PrecisionId, 'target-error'>[]).find(
    (id) => PRECISION_ITERATIONS[id] === request.iterations,
  );
  return found ?? 'normal';
}

/** The band as a fraction of the figure. Zero at a mean of zero rather than infinity. */
export function relativeError(estimate: Estimate): number {
  return estimate.mean === 0 ? 0 : estimate.error / estimate.mean;
}
```

- [ ] **Step 4: Run the precision test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/precision.test.ts
```

Expected: PASS, 8 tests.

- [ ] **Step 5: Write the failing run-loop test**

Append to `src/lib/sim/run.test.ts`:

```ts
describe('a target-error run', () => {
  /**
   * A fake engine whose relative error shrinks with every step, so the loop terminates on
   * the engine's own answer rather than on a count this test chose. `tickMs: 0` keeps it
   * instant; `ticks: 1` keeps the progress traffic down.
   */
  function stepPool(): SimPool {
    return createPool({
      hardwareConcurrency: 2,
      spawn: () => createFakeWorker(createFakeEngine({ tickMs: 0, ticks: 1 })),
    });
  }

  it('runs in thousand-iteration steps and stops when the engine says the band is inside the target', async () => {
    const pool = stepPool();
    const updates: RunUpdate[] = [];
    const handle = runSim(
      pool,
      { ...baseInput, iterations: 30_000, targetError: 0.005, stepIterations: 1000 },
      (update) => updates.push(update),
    );
    const result = await handle.result;

    // Every step is a multiple of the step size, nothing overshoots the ceiling, and the
    // pooled figure is the whole run's, not the last step's.
    expect(result.iterations_run % 1000).toBe(0);
    expect(result.iterations_run).toBeGreaterThanOrEqual(1000);
    expect(result.iterations_run).toBeLessThanOrEqual(30_000);
    expect(result.dps.error / result.dps.mean).toBeLessThanOrEqual(0.005);
    expect(result.request.target_error).toBe(0.005);
    // The ceiling is what the request carries, so a saved sim says what bounded it.
    expect(result.request.iterations).toBe(30_000);
    // The progress line has an error to show from the first step onwards.
    expect(updates.at(-1)?.relativeError).toBeLessThanOrEqual(0.005);
    expect(updates.at(-1)?.iterationsDone).toBe(result.iterations_run);
    pool.terminate();
  });

  it('stops at the ceiling when the band never closes', async () => {
    const pool = stepPool();
    const handle = runSim(
      pool,
      { ...baseInput, iterations: 3000, targetError: 0.0000001, stepIterations: 1000 },
      () => {},
    );
    const result = await handle.result;
    expect(result.iterations_run).toBe(3000);
    pool.terminate();
  });

  it('is today’s single pass when no target error is asked for', async () => {
    const pool = stepPool();
    const handle = runSim(pool, { ...baseInput, iterations: 500 }, () => {});
    const result = await handle.result;
    expect(result.iterations_run).toBe(500);
    expect(result.request.target_error).toBeUndefined();
    pool.terminate();
  });

  it('a stop between steps ends the run rather than starting another one', async () => {
    const pool = stepPool();
    const handle = runSim(
      pool,
      { ...baseInput, iterations: 30_000, targetError: 0.0000001, stepIterations: 1000 },
      () => {},
    );
    handle.cancel();
    await expect(handle.result).rejects.toThrow(SimRunError);
    pool.terminate();
  });
});
```

`baseInput` is the file's existing `RunInput` literal; add `import type { RunUpdate } from './run';` and `import type { SimPool } from './worker';` if the file lacks them.

- [ ] **Step 6: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/run.test.ts
```

Expected: FAIL — `result.request.target_error` is `undefined` and `iterations_run` is 30,000 in one pass.

- [ ] **Step 7: Grow the loop in `run.ts`**

Replace `RunInput`, `RunUpdate`, `buildSimRequest` and the body of `execute()`.

```ts
export interface RunInput {
  spec: string;
  source: CharacterSource;
  character: CharacterSpec;
  encounter: EncounterSpec;
  /** The count for a fixed run; the ceiling for a target-error run. */
  iterations: number;
  randomSeed?: number;
  /** > 0 turns this into a target-error run (contract 1.2). */
  targetError?: number;
  /** The size of one step of a target-error run. Ignored without `targetError`. */
  stepIterations?: number;
}

export interface RunUpdate {
  estimate: Estimate;
  iterationsDone: number;
  /** The fixed count, or the ceiling of a target-error run. */
  iterationsTotal: number;
  /** `error / mean`, for the "±41 DPS, 0.4%" half of the progress line. Zero before the first tick. */
  relativeError: number;
}

export function buildSimRequest(input: RunInput): SimRequest {
  const request: SimRequest = {
    engine_version: ENGINE_VERSION,
    spec: input.spec,
    source: input.source,
    character: input.character,
    encounter: input.encounter,
    iterations: input.iterations,
    random_seed: input.randomSeed ?? 0,
  };
  // Omitted rather than sent as 0: `target_error` is `omitempty` on the Go side, and a
  // request that carries the key with a zero in it reads, in the drawer and in a share
  // URL, as a deliberate choice rather than as today's fixed-count run.
  return (input.targetError ?? 0) > 0 ? { ...request, target_error: input.targetError } : request;
}
```

and inside `runSim`:

```ts
  async function execute(): Promise<SimResult> {
    const startedAt = now();
    const request = buildSimRequest(input);
    const requestJSON = JSON.stringify(request);
    const targetError = input.targetError ?? 0;
    const step = targetError > 0 ? Math.max(1, input.stepIterations ?? STEP_ITERATIONS_DEFAULT) : total;

    // Every shard result of every step, in order. simCombine pools the lot at the end, so
    // the figure a target-error run reports is the whole run's and not the last step's.
    const everyPart: string[] = [];
    let done = 0;

    onUpdate({ estimate: EMPTY_ESTIMATE, iterationsDone: 0, iterationsTotal: total, relativeError: 0 });

    for (;;) {
      const thisStep = Math.min(step, total - done);
      if (thisStep <= 0) break;

      // The step's own request: the same envelope with this step's count. The seed is left
      // alone -- 0 means "random", which is what makes each step independent, and a paired
      // run that pinned one gets the same pinned one every step, which is what pairing means.
      const stepJSON = JSON.stringify({ ...request, iterations: thisStep });
      const shards = await pool.split(stepJSON, Math.min(pool.size, thisStep));

      const latest = new Map<number, ShardProgress>();
      const before = done;
      const parts = await pool.run(shards, `${callbackId}-s${everyPart.length}`, (progress) => {
        latest.set(progress.shard, progress);
        const pooled = combineEstimate([...latest.values()]);
        onUpdate({
          estimate: pooled.estimate,
          iterationsDone: before + pooled.iterationsDone,
          iterationsTotal: total,
          relativeError: relativeError(pooled.estimate),
        });
      });

      everyPart.push(...parts);
      done += thisStep;

      const combinedJSON = await pool.combine(everyPart);
      const combined = JSON.parse(combinedJSON) as SimResult;
      onUpdate({
        estimate: combined.dps,
        iterationsDone: combined.iterations_run,
        iterationsTotal: total,
        relativeError: relativeError(combined.dps),
      });

      // The engine decides. Never a comparison written here: the page does not own a
      // stopping rule over an estimate it did not pool itself.
      if (targetError === 0 || done >= total) {
        return {
          ...combined,
          engine_version: ENGINE_VERSION,
          request,
          lane: 'browser',
          duration_ms: Math.round(now() - startedAt),
        };
      }
      if (!(await pool.needsMore(combinedJSON, requestJSON))) {
        return {
          ...combined,
          engine_version: ENGINE_VERSION,
          request,
          lane: 'browser',
          duration_ms: Math.round(now() - startedAt),
        };
      }
    }

    // Unreachable for any positive `total`; a zero-iteration request is refused at the
    // boundary long before it reaches here, and throwing says so rather than returning a
    // result nothing produced.
    throw new Error('sim run: no iterations to run');
  }
```

`cancel()` keeps its body; add the run's own abort prefix so a stop between steps kills the shards of the step in flight:

```ts
    cancel() {
      cancelled = true;
      pool.abort(callbackId);
    },
```

— unchanged, because the pool's abort rule is a prefix match on `<callbackId>-`, and every step's ids are `<callbackId>-s<n>-<shard>`. Add a guard at the top of the loop so a cancel that lands between two steps does not start a third:

```ts
      if (cancelled) throw new Error('sim run: stopped between steps');
```

New imports at the top of `run.ts`:

```ts
import { relativeError } from './precision';
import { STEP_ITERATIONS_DEFAULT } from './types';
```

- [ ] **Step 8: Run the run test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/run.test.ts
```

Expected: PASS. Existing `run.test.ts` cases that assert an `onUpdate` payload need `relativeError` added to their expected object.

- [ ] **Step 9: Write the failing store test**

Append to `src/lib/sim/store.test.ts`:

```ts
describe('precision', () => {
  it('opens on normal and carries the chosen precision into the request', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    expect(sim.precisionId).toBe('normal');

    sim.setPrecisionId('fast');
    await sim.run();
    expect(sim.result?.request.iterations).toBe(500);
    expect(sim.result?.request.target_error).toBeUndefined();
  });

  it('a target-error run sends the lane’s ceiling and the target', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    sim.setPrecisionId('target-error');
    expect(sim.iterationsTotal).toBe(0);
    await sim.run();
    expect(sim.result?.request.target_error).toBe(0.005);
    expect(sim.result?.request.iterations).toBe(30_000);
    expect(sim.result?.iterations_run).toBeLessThanOrEqual(30_000);
    expect(sim.relativeError).toBeGreaterThan(0);
  });
});
```

- [ ] **Step 10: Wire the store**

In `src/lib/sim/store.svelte.ts`:

- replace `let precision = $state<IterationCount>(3000);` with

```ts
  let precisionId = $state<PrecisionId>('normal');
  let relative = $state(0);
```

- the lane the ceiling comes from is the browser for `run()` and the server for `runOnServer()`; there is no third.
- replace the `precision` getter with:

```ts
    get precisionId() {
      return precisionId;
    },
    /** `error / mean` of the figure on screen, for the progress line and the details card. */
    get relativeError() {
      return relative;
    },
```

- replace `setPrecision` with:

```ts
    setPrecisionId(value: PrecisionId): void {
      precisionId = value;
    },
```

- in `run()`, replace `iterationsTotal = precision;` and the `RunInput` literal with:

```ts
      const plan = precisionPlan(precisionId, 'browser');
      iterationsTotal = plan.iterations;
      iterationsDone = 0;
      relative = 0;
      …
      const input: RunInput = {
        spec: character.spec,
        source: character.source,
        character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
        encounter: settings.encounter,
        iterations: plan.iterations,
        targetError: plan.targetError,
        stepIterations: plan.step,
      };
```

- in the `runSim` progress callback, add `relative = update.relativeError;`
- in `runOnServer()`, use `precisionPlan(precisionId, 'server')` the same way, and set `relative` from the finished result with `relativeError(finished.dps)`.
- in `restorePreviousResult()` and `adoptResult()`, set `relative = relativeError(result.dps)` / `relativeError(next.dps)` so the card is right after a saved sim is adopted.

New imports:

```ts
import { precisionPlan, relativeError, type PrecisionId } from './precision';
```

and drop the now-unused `IterationCount` import.

- [ ] **Step 11: Follow the rename through the components**

`RunControl.svelte` takes `precision: IterationCount` and `onprecision`; Task 7 rewrites that control. For now, change its two props to `precisionId: PrecisionId` / `onprecision: (value: PrecisionId) => void` and its checkbox to `checked={precisionId === 'high'}` / `onprecision(checked ? 'high' : 'normal')`, and update `SimView.svelte`'s two bindings. Task 7 replaces the control properly.

- [ ] **Step 12: Run the whole unit suite**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run
```

Expected: PASS.

- [ ] **Step 13: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/lib/sim src/components/sim && npx prettier --check src/lib/sim src/components/sim && \
  git add -A src/lib/sim src/components/sim && \
  git commit -m "feat(sim): fast, normal, high, and until the band is inside half a per cent

Design 4.2: the step loop runs a thousand iterations at a time and asks
the engine after each one, pooling every step's shards so the figure is
the whole run's. The ceiling is the lane's -- 30,000 in the browser,
100,000 on the server -- and the request carries it, so a saved sim
says what bounded it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 7: The precision control, the progress line and the details card

Design 4.2 and 5.1. The high-precision checkbox becomes a four-option select; the progress line gains the error; and a details card beside the results states margin of error, iterations, processing time, engine version (linking to `/sim/specs`) and lane.

Two different figures are on screen and they must never be confused, so they are labelled apart:

- **`± <n> DPS`** is the 95% confidence band, `1.96 × error` — what `estimate.ts`'s `confidenceBand` already computes and what the big figure has always shown.
- **`<n>%`** is the *relative standard error*, `error / mean` — the quantity `simNeedsMore` compares against `target_error`. The "until ±0.5%" control is named after this one, so the percent beside the progress line has to be this one or the control would stop at a number that does not match its own label.

**Files:**
- Create: `src/lib/sim/details.ts`
- Create: `src/components/sim/DetailsCard.svelte`
- Modify: `src/components/sim/RunControl.svelte`
- Modify: `src/components/sim/SimView.svelte`
- Modify: `src/components/sim/SavedSim.svelte`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/details.test.ts`, `tests/e2e/sim-run.spec.ts`

**Interfaces:**
- Consumes: `precision.ts` (Task 6), `estimate.ts`'s `confidenceBand`, `version.ts`'s `engineLabel`.
- Produces:
  - `details.ts`: `interface RunDetails`, `runDetails(result: SimResult): RunDetails`, `percentLabel(fraction: number): string`
  - `DetailsCard.svelte`: props `{ result: SimResult }`, testid `sim-details-card`
  - `RunControl.svelte` props: `precisionId: PrecisionId`, `relativeError: number`, `onprecision: (value: PrecisionId) => void`; testids `sim-precision` (the select) and `sim-target-error` (the line under it)

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/details.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/result.json';
import { percentLabel, runDetails } from './details';
import type { SimResult } from './types';

const base = fixture as unknown as SimResult;

function result(over: Partial<SimResult>): SimResult {
  return { ...base, ...over };
}

describe('runDetails', () => {
  it('reports the 95% band in DPS and the relative standard error as a percent', () => {
    const details = runDetails(
      result({ dps: { mean: 1000, stddev: 100, error: 4, min: 0, max: 0 } }),
    );
    // 1.96 * 4 = 7.84, rounded.
    expect(details.bandDps).toBe(8);
    expect(details.errorPercent).toBeCloseTo(0.004);
  });

  it('carries the iterations, the wall clock, the engine and the lane straight through', () => {
    const details = runDetails(result({ iterations_run: 7000, duration_ms: 12_345, lane: 'server' }));
    expect(details.iterations).toBe(7000);
    expect(details.processingMs).toBe(12_345);
    expect(details.lane).toBe('server');
    expect(details.engineVersion).toBe(base.engine_version);
  });

  it('says a target-error run stopped at the ceiling, and says nothing of the sort otherwise', () => {
    const capped = result({
      iterations_run: 30_000,
      request: { ...base.request, iterations: 30_000, target_error: 0.005 },
    });
    expect(runDetails(capped).hitCeiling).toBe(true);

    const converged = result({
      iterations_run: 4000,
      request: { ...base.request, iterations: 30_000, target_error: 0.005 },
    });
    expect(runDetails(converged).hitCeiling).toBe(false);

    const fixed = result({ iterations_run: 3000, request: { ...base.request, iterations: 3000 } });
    expect(runDetails(fixed).hitCeiling).toBe(false);
  });

  it('never divides by a mean of zero', () => {
    expect(runDetails(result({ dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 } })).errorPercent).toBe(
      0,
    );
  });
});

describe('percentLabel', () => {
  it('is two decimals and a sign-free per cent', () => {
    expect(percentLabel(0.004)).toBe('0.40%');
    expect(percentLabel(0.005)).toBe('0.50%');
    expect(percentLabel(0)).toBe('0.00%');
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/details.test.ts
```

Expected: FAIL — `Failed to resolve import "./details"`.

- [ ] **Step 3: Write `details.ts`**

```ts
// web/src/lib/sim/details.ts
// Design 5.1's details card: "margin of error, iterations, processing time, engine version
// and lane as a details card beside the results, the way Raidbots' sidebar reads".
//
// Two figures, never the same one twice: `bandDps` is the 95% confidence band the big
// figure already shows (1.96 x the standard error), and `errorPercent` is the relative
// standard error, which is the quantity `simNeedsMore` compares against `target_error`.
// The "until +/-0.5%" control is named after the second, so the second is what any percent
// on this lane means.
import { confidenceBand } from './estimate';
import { relativeError } from './precision';
import type { SimResult } from './types';

export interface RunDetails {
  /** The 95% confidence band, in DPS, rounded. */
  bandDps: number;
  /** `error / mean`, as a fraction. */
  errorPercent: number;
  iterations: number;
  processingMs: number;
  engineVersion: string;
  lane: 'browser' | 'server';
  /**
   * True when a target-error run reached the lane's ceiling with the band still open --
   * the design's "the results line says which" of the two things stopped it.
   */
  hitCeiling: boolean;
}

export function runDetails(result: SimResult): RunDetails {
  const target = result.request.target_error ?? 0;
  return {
    bandDps: Math.round(confidenceBand(result.dps)),
    errorPercent: relativeError(result.dps),
    iterations: result.iterations_run,
    processingMs: result.duration_ms,
    engineVersion: result.engine_version,
    lane: result.lane,
    hitCeiling: target > 0 && result.iterations_run >= result.request.iterations,
  };
}

/** A fraction as a per cent with two decimals: 0.004 is "0.40%". */
export function percentLabel(fraction: number): string {
  return `${(fraction * 100).toFixed(2)}%`;
}
```

- [ ] **Step 4: Add the copy**

In `src/lib/sim/copy.ts`, append inside `simCopy`, and delete `highPrecision` and `precisionNote` (the checkbox they named is gone):

```ts
  // --- Design 4.2: the precision control. ---
  precision: 'Precision',
  precisionLabel: {
    fast: 'Fast, 500 iterations',
    normal: 'Normal, 3,000 iterations',
    high: 'High, 10,000 iterations',
    'target-error': 'Until ±0.5%',
  } as Record<string, string>,
  /** Under the select while the target-error run is chosen; `ceiling` is the lane's. */
  targetErrorNote: (ceiling: string): string =>
    `Runs a thousand iterations at a time until the error is inside half a per cent, or until ${ceiling} iterations, whichever comes first.`,
  /** The results line when the ceiling, not the target, is what stopped the run. */
  targetErrorCeiling: (ceiling: string): string =>
    `Stopped at ${ceiling} iterations with the error still outside half a per cent.`,

  // --- Design 5.1: the details card. ---
  details: 'This run',
  detailsMargin: 'Margin of error',
  /** The 95% band and the relative standard error, side by side and never conflated. */
  detailsMarginValue: (band: string, percent: string): string => `± ${band} DPS · ${percent}`,
  detailsIterations: 'Iterations',
  detailsProcessing: 'Processing time',
  detailsEngine: 'Engine',
  detailsLane: 'Ran on',
  detailsLaneBrowser: 'your browser',
  detailsLaneServer: 'our servers',
```

- [ ] **Step 5: Write `DetailsCard.svelte`**

```svelte
<!-- web/src/components/sim/DetailsCard.svelte -->
<!-- Design 5.1. Five facts about the run itself, beside the results rather than mixed into
     them: none of them is about the character, and all five are what someone asks when
     they doubt the number. The engine version links to /sim/specs, which is where the
     honesty about each spec lives. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { percentLabel, runDetails } from '../../lib/sim/details';
  import type { SimResult } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';

  let { result }: { result: SimResult } = $props();

  const details = $derived(runDetails(result));
  const rows = $derived([
    {
      id: 'margin',
      label: simCopy.detailsMargin,
      value: simCopy.detailsMarginValue(
        details.bandDps.toLocaleString('en-US'),
        percentLabel(details.errorPercent),
      ),
    },
    {
      id: 'iterations',
      label: simCopy.detailsIterations,
      value: details.iterations.toLocaleString('en-US'),
    },
    {
      id: 'processing',
      label: simCopy.detailsProcessing,
      value: `${(details.processingMs / 1000).toFixed(1)} s`,
    },
    {
      id: 'lane',
      label: simCopy.detailsLane,
      value: details.lane === 'server' ? simCopy.detailsLaneServer : simCopy.detailsLaneBrowser,
    },
  ]);
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-details-card"
>
  <h2 class="section-title text-[15px]">{simCopy.details}</h2>
  <dl class="grid grid-cols-[minmax(0,1fr)_auto] gap-x-4 gap-y-1 text-[13px]">
    {#each rows as row (row.id)}
      <dt class="text-muted">{row.label}</dt>
      <dd class="tabular text-strong text-right font-mono" data-testid={`sim-details-${row.id}`}>
        {row.value}
      </dd>
    {/each}
    <dt class="text-muted">{simCopy.detailsEngine}</dt>
    <dd class="text-right">
      <a class="tabular font-mono text-[13px]" href="/sim/specs" data-testid="sim-details-engine">
        {engineLabel(details.engineVersion)}
      </a>
    </dd>
  </dl>
  {#if details.hitCeiling}
    <p class="text-muted text-[12px]" data-testid="sim-details-ceiling">
      {simCopy.targetErrorCeiling(result.request.iterations.toLocaleString('en-US'))}
    </p>
  {/if}
</section>
```

- [ ] **Step 6: Replace the precision checkbox in `RunControl.svelte`**

Props: swap `precision: IterationCount` for `precisionId: PrecisionId`, add `relativeError: number`, and change `onprecision` to `(value: PrecisionId) => void`. Imports lose `ITERATIONS`/`IterationCount` and gain:

```ts
  import { LANE_ITERATION_CEILING, PRECISIONS, type PrecisionId } from '../../lib/sim/precision';
  import { percentLabel } from '../../lib/sim/details';
```

Replace the whole `<label>`-with-checkbox block:

```svelte
    <label class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.precision}</span>
      <select
        class="border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 border px-3 text-[14px] font-semibold md:min-h-9"
        disabled={running || serverRunning}
        value={precisionId}
        onchange={(event) => onprecision(event.currentTarget.value as PrecisionId)}
        data-testid="sim-precision"
      >
        {#each PRECISIONS as id (id)}
          <option value={id}>{simCopy.precisionLabel[id] ?? id}</option>
        {/each}
      </select>
    </label>
```

and, immediately after the closing `</div>` of that control group:

```svelte
  {#if precisionId === 'target-error'}
    <p class="text-muted order-last w-full text-[12px]" data-testid="sim-target-error">
      {simCopy.targetErrorNote(LANE_ITERATION_CEILING.browser.toLocaleString('en-US'))}
    </p>
  {/if}
```

Widen the progress line so it carries the error once there is one (design 4.2: "the margin-of-error line is always shown, as is the iteration count and processing time"):

```ts
  const errorText = $derived(relativeError > 0 ? ` · ${percentLabel(relativeError)}` : '');
  const progressLine = $derived(
    phase === 'done'
      ? `${iterationsDone.toLocaleString('en-US')} ${simCopy.iterations}${errorText}`
      : `${iterationsDone.toLocaleString('en-US')} of ${iterationsTotal.toLocaleString('en-US')} ${simCopy.iterations}${errorText}`,
  );
```

- [ ] **Step 7: Mount the card**

In `SimView.svelte`, change the two `RunControl` props (`precisionId={store.precisionId}`, `relativeError={store.relativeError}`, `onprecision={(value) => store.setPrecisionId(value)}`) and render the card immediately after the results:

```svelte
        {#if store.result !== null && !comparing}
          <DetailsCard result={store.result} />
        {/if}
```

with `import DetailsCard from './DetailsCard.svelte';`. In `SavedSim.svelte`, render `<DetailsCard {result} />` immediately after the results block, so a saved sim carries the same card.

- [ ] **Step 8: Extend the e2e**

Append to `tests/e2e/sim-run.spec.ts`:

```ts
test('the precision select offers four choices and the details card states the run', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  const precision = page.getByTestId('sim-precision');
  await expect(precision).toHaveValue('normal');
  await expect(precision.locator('option')).toHaveCount(4);
  await expect(page.getByTestId('sim-target-error')).toBeHidden();

  await precision.selectOption('target-error');
  await expect(page.getByTestId('sim-target-error')).toContainText('30,000');

  // Fast keeps the browser suite quick; the card is the same card at every precision.
  await precision.selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });
  await expect(page.getByTestId('sim-details-iterations')).toHaveText('500');
  await expect(page.getByTestId('sim-details-margin')).toContainText('DPS');
  await expect(page.getByTestId('sim-details-margin')).toContainText('%');
  await expect(page.getByTestId('sim-details-lane')).toHaveText('your browser');
  await expect(page.getByTestId('sim-details-engine')).toHaveAttribute('href', '/sim/specs');
  await expect(page.getByTestId('sim-progress')).toContainText('%');
});
```

`FURY` is the constant the file already declares.

- [ ] **Step 9: Run both suites**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-run.spec.ts --project=desktop
```

Expected: PASS.

- [ ] **Step 10: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && npx eslint src/lib/sim src/components/sim tests/e2e && \
  npx prettier --check src/lib/sim src/components/sim tests/e2e && git add -A src/lib/sim src/components/sim tests/e2e && \
  git commit -m "feat(sim): the precision select, the error in the progress line, the details card

Four precisions in one select, the relative standard error beside the
iteration count while a run is in flight, and a card stating margin of
error, iterations, processing time, engine version and lane. The band
in DPS is the 95% interval and the per cent is the relative standard
error -- the quantity the target-error run is named after -- and the
two are labelled apart so neither can be read as the other.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 8: The id vocabulary, generated from IDS.md, and the panel's grouping

Design 4.3 wants the whole list, "grouped the way the engine groups them". The list itself is `sim/request/IDS.md`, generated by the engine from its own protobuf descriptors, and copying several hundred ids into `web/` would be a second copy to keep in step — which is exactly the reason `settings.ts`'s current comment gives for the preset being the only list it holds. So the vocabulary is generated into `src/data/generated/sim-ids.json` by a script, and `buffs.ts` holds only the grouping: a table from IDS.md's own "lands in" / "sets" column to a panel section.

Contract **A7** has the sim module lane regenerate IDS.md with three things it does not carry today: the graded `<id>:improved` rows (contract 1.7), a **World buffs** section, and a **Stats** section holding the fork's `proto.Stat` enum names in snake case. The parser reads all three now, so each appears with no change here the day it is generated — and the Stats vocabulary is what part B's `/sim/weights` builds its `WeightsSpec` from.

Contract **10.8** pins that vocabulary exactly, so `stats.ts` carries the list verbatim as its fallback and a test asserts the generated section agrees with it whenever one exists. The engine has **one `hit` and one `crit`**: there is no `melee_hit`, `spell_hit`, `melee_crit` or `spell_crit` anywhere on this lane. Haste *is* split — `spell_haste` and `melee_haste` — and `MP5` is spelled `mp5`.

**Files:**
- Create: `scripts/ids-md.mjs`
- Create: `scripts/sync-sim-ids.mjs`
- Create: `src/lib/sim/buffs.ts`
- Create: `src/lib/sim/stats.ts`
- Modify: `package.json`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/ids-md.test.ts`, `src/lib/sim/buffs.test.ts`, `src/lib/sim/stats.test.ts`

**Interfaces:**
- Consumes: nothing.
- Produces:
  - `scripts/ids-md.mjs`: `parseIdsMarkdown(markdown) -> { buffs: [{id, message}], consumables: [{id, sets}], professions: [string], worldBuffs: [string], stats: [string] }`
  - `src/data/generated/sim-ids.json`: that object
  - `buffs.ts`: `type BuffGroupId`, `BUFF_GROUPS: readonly BuffGroupId[]`, `interface BuffRow { id: string; group: BuffGroupId; graded: boolean; kind: 'buff' | 'consumable' }`, `interface SimIdsFile`, `buildCatalogue(ids): BuffRow[]`, `CATALOGUE: readonly BuffRow[]`, `rowsIn(group): BuffRow[]`, `type BuffGrade = 'off' | 'on' | 'improved'`, `IMPROVED_SUFFIX`, `gradeOf(selected, id)`, `setGrade(selected, id, grade)`
  - `stats.ts`: `PINNED_STATS: readonly string[]` (contract 10.8, verbatim), `SIM_STATS: readonly string[]` (the generated section when there is one, `PINNED_STATS` until then), `statLabel(id: string): string` — **part B's `/sim/weights` builds `WeightsSpec.stats` and its reference picker from these**
  - `copy.ts`: `simCopy.buffGroupLabel: Record<string, string>`, `simCopy.gradeLabel: Record<string, string>`, `simCopy.statLabel: Record<string, string>`

- [ ] **Step 1: Write the failing parser test**

Create `src/lib/sim/ids-md.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
// A plain .mjs so scripts/sync-sim-ids.mjs can import the same parser this test drives:
// the script runs under node before vitest ever starts, so the parser cannot be a .ts.
import { parseIdsMarkdown } from '../../../scripts/ids-md.mjs';

const SAMPLE = `# The sim request's id vocabulary

## Buffs

Prose the parser must skip, including a | pipe | in a sentence.

| id | lands in |
| --- | --- |
| \`battle_shout\` | RaidBuffs |
| \`battle_shout:improved\` | RaidBuffs |
| \`blessing_of_kings\` | IndividualBuffs |
| \`sunder_armor\` | Debuffs |
| \`atiesh_mage\` | PartyBuffs |
| \`songflower_serenade\` | IndividualBuffs |

## World buffs

| id |
| --- |
| \`songflower_serenade\` |

## Consumables

| id | sets |
| --- | --- |
| \`flask_of_supreme_power\` | Consumes.flask |
| \`main_hand_imbue:shadow_oil\` | Consumes.main_hand_imbue |
| \`food_grilled_squid\` | Consumes.food |

## Professions

| slug | engine enum |
| --- | --- |
| \`alchemy\` | Profession.Alchemy |

## Stats

| id | proto.Stat |
| --- | --- |
| \`attack_power\` | AttackPower |
| \`crit\` | Crit |
| \`melee_haste\` | MeleeHaste |
| \`mp5\` | MP5 |
| \`spell_haste\` | SpellHaste |
`;

describe('parseIdsMarkdown', () => {
  const parsed = parseIdsMarkdown(SAMPLE);

  it('reads every buff row with the message it lands in', () => {
    expect(parsed.buffs).toContainEqual({ id: 'battle_shout', message: 'RaidBuffs' });
    expect(parsed.buffs).toContainEqual({ id: 'sunder_armor', message: 'Debuffs' });
    expect(parsed.buffs).toContainEqual({ id: 'atiesh_mage', message: 'PartyBuffs' });
  });

  it('keeps a graded id as its own row, so the catalogue can pair it with the plain one', () => {
    expect(parsed.buffs).toContainEqual({ id: 'battle_shout:improved', message: 'RaidBuffs' });
  });

  it('reads the world-buff section when there is one', () => {
    expect(parsed.worldBuffs).toEqual(['songflower_serenade']);
  });

  it('reads consumables with the Consumes field they set', () => {
    expect(parsed.consumables).toContainEqual({
      id: 'main_hand_imbue:shadow_oil',
      sets: 'Consumes.main_hand_imbue',
    });
  });

  it('reads professions as plain slugs', () => {
    expect(parsed.professions).toEqual(['alchemy']);
  });

  it('reads the Stats section contract A7 adds, in the enum’s own snake case', () => {
    expect(parsed.stats).toEqual(['attack_power', 'crit', 'melee_haste', 'mp5', 'spell_haste']);
    // A7 and 10.8: haste is split, MP5 is `mp5`, and hit and crit are single stats.
    expect(parsed.stats).not.toContain('haste');
    expect(parsed.stats).not.toContain('melee_crit');
  });

  it('skips prose, headings and the separator row rather than reading them as ids', () => {
    expect(parsed.buffs.map((row) => row.id)).not.toContain('---');
    expect(parsed.buffs.map((row) => row.id)).not.toContain('id');
    expect(parsed.buffs).toHaveLength(6);
  });

  it('answers empty lists for a document with no tables at all', () => {
    expect(parseIdsMarkdown('# nothing here')).toEqual({
      buffs: [],
      consumables: [],
      professions: [],
      worldBuffs: [],
      stats: [],
    });
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/ids-md.test.ts
```

Expected: FAIL — cannot resolve `../../../scripts/ids-md.mjs`.

- [ ] **Step 3: Write the parser**

Create `web/scripts/ids-md.mjs`:

```js
// web/scripts/ids-md.mjs
// sim/request/IDS.md is generated by the engine from its own protobuf descriptors and is
// the only place the buff, consumable and profession vocabulary exists. This parses it, so
// web/ holds the grouping and never a second copy of the list. A plain .mjs because
// scripts/sync-sim-ids.mjs runs under node long before any TypeScript is compiled;
// src/lib/sim/ids-md.test.ts drives it through vitest all the same.

/** @typedef {{ id: string, message: string }} BuffIdRow */
/** @typedef {{ id: string, sets: string }} ConsumableIdRow */
/** @typedef {{ buffs: BuffIdRow[], consumables: ConsumableIdRow[], professions: string[], worldBuffs: string[], stats: string[] }} SimIds */

/** A markdown table cell's contents, with the backticks IDS.md wraps every id in removed. */
function cell(text) {
  return text.trim().replace(/^`|`$/g, '');
}

/** True for `| --- | --- |` and for the header row, neither of which is data. */
function isStructural(cells) {
  return cells.every((value) => /^-{3,}$/.test(value)) || cells[0] === 'id' || cells[0] === 'slug';
}

/**
 * Every table row under each `## ` heading, keyed by the heading's lowercased text. Prose
 * is skipped by requiring a line to start and end with a pipe.
 * @param {string} markdown
 * @returns {Map<string, string[][]>}
 */
function sectionsOf(markdown) {
  /** @type {Map<string, string[][]>} */
  const sections = new Map();
  let heading = '';
  for (const line of markdown.split('\n')) {
    const title = /^##\s+(.+?)\s*$/.exec(line);
    if (title !== null) {
      heading = title[1].toLowerCase();
      if (!sections.has(heading)) sections.set(heading, []);
      continue;
    }
    const trimmed = line.trim();
    if (heading === '' || !trimmed.startsWith('|') || !trimmed.endsWith('|')) continue;
    const cells = trimmed.slice(1, -1).split('|').map(cell);
    if (isStructural(cells)) continue;
    sections.get(heading)?.push(cells);
  }
  return sections;
}

/**
 * @param {string} markdown
 * @returns {SimIds}
 */
export function parseIdsMarkdown(markdown) {
  const sections = sectionsOf(markdown);
  const buffRows = sections.get('buffs') ?? [];
  const consumableRows = sections.get('consumables') ?? [];
  const professionRows = sections.get('professions') ?? [];
  const worldRows = sections.get('world buffs') ?? [];
  // Contract A7 adds both of these sections; absent ones read as empty lists, so this
  // parser works against today's IDS.md and against the regenerated one unchanged.
  const statRows = sections.get('stats') ?? [];
  return {
    buffs: buffRows.map((cells) => ({ id: cells[0], message: cells[1] ?? '' })),
    consumables: consumableRows.map((cells) => ({ id: cells[0], sets: cells[1] ?? '' })),
    professions: professionRows.map((cells) => cells[0]),
    worldBuffs: worldRows.map((cells) => cells[0]),
    stats: statRows.map((cells) => cells[0]),
  };
}
```

- [ ] **Step 4: Write the sync script and wire it in**

Create `web/scripts/sync-sim-ids.mjs`:

```js
// web/scripts/sync-sim-ids.mjs
// sim/request/IDS.md -> web/src/data/generated/sim-ids.json, before dev, check, test and
// build. The file is gitignored like everything else under src/data/generated: it is
// derived, and a committed copy is a copy that goes stale.
//
// A missing IDS.md is fatal rather than an empty file: the settings panel would render
// with no rows at all and say nothing about why, which is the failure mode the whole
// "an id the engine cannot map is an error at the boundary" rule exists to avoid.
import { mkdir, readFile, writeFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { parseIdsMarkdown } from './ids-md.mjs';

const webRoot = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = path.join(webRoot, '..');
const source = path.join(repoRoot, 'sim', 'request', 'IDS.md');
const target = path.join(webRoot, 'src', 'data', 'generated', 'sim-ids.json');

let markdown;
try {
  markdown = await readFile(source, 'utf8');
} catch (cause) {
  throw new Error(`sync-sim-ids: ${source} could not be read; the engine generates it with \`go run ./internal/genids\``, {
    cause,
  });
}

const ids = parseIdsMarkdown(markdown);
if (ids.buffs.length === 0 || ids.consumables.length === 0) {
  throw new Error(`sync-sim-ids: ${source} carried no buff or consumable table`);
}

await mkdir(path.dirname(target), { recursive: true });
await writeFile(target, `${JSON.stringify(ids, null, 2)}\n`, 'utf8');
```

In `package.json`, add the script and put it in the `sync` chain:

```json
    "sync:sim-ids": "node scripts/sync-sim-ids.mjs",
    "sync": "npm run sync:data && npm run sync:sim-ids && npm run sync:report-fixture && npm run sync:duckdb",
```

Run it once:

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npm run sync:sim-ids && head -20 src/data/generated/sim-ids.json
```

Expected: the file exists and its first buff row is `arcane_brilliance` / `RaidBuffs`.

- [ ] **Step 5: Write the failing catalogue test**

Create `src/lib/sim/buffs.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import generated from '../../data/generated/sim-ids.json';
import { simCopy } from './copy';
import {
  BUFF_GROUPS,
  CATALOGUE,
  IMPROVED_SUFFIX,
  buildCatalogue,
  gradeOf,
  rowsIn,
  setGrade,
} from './buffs';

describe('the catalogue', () => {
  it('groups every buff by the message IDS.md says it lands in', () => {
    const groups = new Map(CATALOGUE.map((row) => [row.id, row.group]));
    expect(groups.get('battle_shout')).toBe('raid-buffs');
    expect(groups.get('sunder_armor')).toBe('debuffs');
    expect(groups.get('atiesh_mage')).toBe('party-buffs');
    expect(groups.get('blessing_of_kings')).toBe('player-buffs');
  });

  it('puts the world buffs in their own section rather than among the blessings', () => {
    const groups = new Map(CATALOGUE.map((row) => [row.id, row.group]));
    for (const id of [
      'rallying_cry_of_the_dragonslayer',
      'songflower_serenade',
      'spirit_of_zandalar',
      'warchiefs_blessing',
      'fengus_ferocity',
      'moldars_moxie',
      'slipkiks_savvy',
      'sayges_fortune',
    ]) {
      expect(groups.get(id), id).toBe('world-buffs');
    }
  });

  it('groups consumables by the Consumes field they set', () => {
    const groups = new Map(CATALOGUE.map((row) => [row.id, row.group]));
    expect(groups.get('flask_of_supreme_power')).toBe('flask');
    expect(groups.get('elixir_of_the_mongoose')).toBe('battle-elixir');
    expect(groups.get('elixir_of_superior_defense')).toBe('guardian-elixir');
    expect(groups.get('food_grilled_squid')).toBe('food');
    expect(groups.get('main_hand_imbue:shadow_oil')).toBe('weapon-imbue');
    expect(groups.get('major_mana_potion')).toBe('potion');
    expect(groups.get('explosive_thorium_grenade')).toBe('explosive');
  });

  it('holds every id the engine publishes and invents none', () => {
    const published = new Set([
      ...generated.buffs.map((row) => row.id),
      ...generated.consumables.map((row) => row.id),
    ]);
    for (const row of CATALOGUE) {
      // A graded row is folded into its plain id and never appears on its own.
      expect(published.has(row.id), row.id).toBe(true);
    }
    for (const id of published) {
      if (id.endsWith(IMPROVED_SUFFIX)) continue;
      expect(
        CATALOGUE.some((row) => row.id === id),
        id,
      ).toBe(true);
    }
  });

  it('folds `<id>:improved` into the plain id as a grade, never as a second row', () => {
    const catalogue = buildCatalogue({
      buffs: [
        { id: 'battle_shout', message: 'RaidBuffs' },
        { id: 'battle_shout:improved', message: 'RaidBuffs' },
        { id: 'thorns', message: 'RaidBuffs' },
      ],
      consumables: [],
      professions: [],
      worldBuffs: [],
      stats: [],
    });
    expect(catalogue.map((row) => row.id)).toEqual(['battle_shout', 'thorns']);
    expect(catalogue.find((row) => row.id === 'battle_shout')?.graded).toBe(true);
    expect(catalogue.find((row) => row.id === 'thorns')?.graded).toBe(false);
  });

  it('names every group', () => {
    for (const group of BUFF_GROUPS) {
      expect(simCopy.buffGroupLabel[group], group).toBeTruthy();
    }
  });

  it('rowsIn is the group’s rows, alphabetical by id', () => {
    const flasks = rowsIn('flask').map((row) => row.id);
    expect(flasks.length).toBeGreaterThan(1);
    expect([...flasks].sort()).toEqual(flasks);
  });
});

describe('the three-way grade', () => {
  it('reads off, on and improved out of a selection', () => {
    expect(gradeOf([], 'battle_shout')).toBe('off');
    expect(gradeOf(['battle_shout'], 'battle_shout')).toBe('on');
    expect(gradeOf(['battle_shout:improved'], 'battle_shout')).toBe('improved');
  });

  it('never leaves both forms in the list', () => {
    expect(setGrade(['battle_shout'], 'battle_shout', 'improved')).toEqual(['battle_shout:improved']);
    expect(setGrade(['battle_shout:improved'], 'battle_shout', 'on')).toEqual(['battle_shout']);
    expect(setGrade(['battle_shout:improved'], 'battle_shout', 'off')).toEqual([]);
  });

  it('leaves every other id alone and its order intact', () => {
    expect(setGrade(['thorns', 'battle_shout', 'blood_pact'], 'battle_shout', 'off')).toEqual([
      'thorns',
      'blood_pact',
    ]);
  });

  it('does not mutate the list it was given', () => {
    const selected = ['battle_shout'];
    setGrade(selected, 'battle_shout', 'off');
    expect(selected).toEqual(['battle_shout']);
  });

  it('names the three grades', () => {
    expect(Object.keys(simCopy.gradeLabel).sort()).toEqual(['improved', 'off', 'on']);
  });
});
```

- [ ] **Step 6: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/buffs.test.ts
```

Expected: FAIL — `Failed to resolve import "./buffs"`.

- [ ] **Step 7: Add the copy**

In `src/lib/sim/copy.ts`, append inside `simCopy`. Leave `customPresetNote` where it is for now — `SettingsBar.svelte` still reads it as the Custom option's `title`, and Task 10 removes the two together once the panel it promised actually exists:

```ts
  // --- Design 4.3: the full buff, debuff and consumable panel, behind "Custom". ---
  buffPanel: 'Everything applied',
  buffPanelNote:
    'The engine’s own list. An id it cannot map fails the run and names itself, rather than being quietly dropped.',
  buffGroupLabel: {
    'raid-buffs': 'Raid buffs',
    'party-buffs': 'Party buffs',
    'player-buffs': 'On this player',
    'world-buffs': 'World buffs',
    debuffs: 'On the target',
    flask: 'Flasks',
    'battle-elixir': 'Battle elixirs',
    'guardian-elixir': 'Guardian elixirs',
    food: 'Food',
    'weapon-imbue': 'Weapon oils and stones',
    potion: 'Potions and runes',
    explosive: 'Explosives',
  } as Record<string, string>,
  gradeLabel: { off: 'Off', on: 'On', improved: 'Improved' } as Record<string, string>,
  /** The three-way control's own accessible name; the buff's name is beside it. */
  gradeFor: (name: string): string => `${name}, how good a version`,
```

- [ ] **Step 8: Write `buffs.ts`**

```ts
// web/src/lib/sim/buffs.ts
// The panel's sections, and nothing else. The id vocabulary itself is IDS.md's, generated
// by the engine from its own protobuf descriptors and copied into
// src/data/generated/sim-ids.json by scripts/sync-sim-ids.mjs -- so this file never holds
// a list of ids, only the table that says which section each engine message and each
// `Consumes` field belongs in. A buff the engine adds appears in the panel the next time
// the data is synced, with no diff here.
//
// Graded ids are contract 1.7: `<id>:improved` for every TristateEffect field. They are
// folded into the plain id as a three-way choice rather than shown as a second checkbox,
// because "Battle Shout" and "Battle Shout (improved)" both ticked is not a state the
// engine has.
import generated from '../../data/generated/sim-ids.json';

export const IMPROVED_SUFFIX = ':improved';

export type BuffGroupId =
  | 'raid-buffs'
  | 'party-buffs'
  | 'player-buffs'
  | 'world-buffs'
  | 'debuffs'
  | 'flask'
  | 'battle-elixir'
  | 'guardian-elixir'
  | 'food'
  | 'weapon-imbue'
  | 'potion'
  | 'explosive';

/** Panel order, top to bottom: what a raid gives you, then what you bring yourself. */
export const BUFF_GROUPS: readonly BuffGroupId[] = [
  'raid-buffs',
  'party-buffs',
  'player-buffs',
  'world-buffs',
  'debuffs',
  'flask',
  'battle-elixir',
  'guardian-elixir',
  'food',
  'weapon-imbue',
  'potion',
  'explosive',
];

export interface SimIdsFile {
  buffs: { id: string; message: string }[];
  consumables: { id: string; sets: string }[];
  professions: string[];
  worldBuffs: string[];
  /** Contract A7's Stats section. Read by stats.ts, not by the catalogue. */
  stats: string[];
}

export interface BuffRow {
  id: string;
  group: BuffGroupId;
  /** True when IDS.md also publishes `<id>:improved`, so the row is a three-way choice. */
  graded: boolean;
  kind: 'buff' | 'consumable';
}

const BY_MESSAGE: Record<string, BuffGroupId> = {
  RaidBuffs: 'raid-buffs',
  PartyBuffs: 'party-buffs',
  IndividualBuffs: 'player-buffs',
  Debuffs: 'debuffs',
};

/**
 * The `Consumes` field each consumable sets, to the kind the design groups by. The engine
 * has one field per effect and the panel has one section per shelf in the bank, so several
 * fields land in one section -- every stat-boosting elixir is a battle elixir, every
 * defensive or regenerative one a guardian elixir, which is the vanilla client's own split.
 */
const BY_CONSUMES_FIELD: Record<string, BuffGroupId> = {
  flask: 'flask',
  agility_elixir: 'battle-elixir',
  strength_buff: 'battle-elixir',
  attack_power_buff: 'battle-elixir',
  spell_power_buff: 'battle-elixir',
  fire_power_buff: 'battle-elixir',
  frost_power_buff: 'battle-elixir',
  shadow_power_buff: 'battle-elixir',
  hit_consumable: 'battle-elixir',
  bogling_root: 'battle-elixir',
  dragon_breath_chili: 'battle-elixir',
  armor_elixir: 'guardian-elixir',
  health_elixir: 'guardian-elixir',
  mana_regen_elixir: 'guardian-elixir',
  zanza_buff: 'guardian-elixir',
  food: 'food',
  alcohol: 'food',
  main_hand_imbue: 'weapon-imbue',
  off_hand_imbue: 'weapon-imbue',
  default_potion: 'potion',
  default_conjured: 'potion',
  filler_explosive: 'explosive',
  sapper_explosive: 'explosive',
};

/** "Consumes.main_hand_imbue" is the field `main_hand_imbue`. */
function consumesField(sets: string): string {
  const dot = sets.lastIndexOf('.');
  return dot === -1 ? sets : sets.slice(dot + 1);
}

export function buildCatalogue(ids: SimIdsFile): BuffRow[] {
  const world = new Set(ids.worldBuffs);
  const graded = new Set(
    ids.buffs
      .filter((row) => row.id.endsWith(IMPROVED_SUFFIX))
      .map((row) => row.id.slice(0, -IMPROVED_SUFFIX.length)),
  );

  const buffs: BuffRow[] = ids.buffs
    .filter((row) => !row.id.endsWith(IMPROVED_SUFFIX))
    .map((row) => ({
      id: row.id,
      group:
        row.message === 'IndividualBuffs' && world.has(row.id)
          ? ('world-buffs' as BuffGroupId)
          : (BY_MESSAGE[row.message] ?? 'player-buffs'),
      graded: graded.has(row.id),
      kind: 'buff' as const,
    }));

  const consumables: BuffRow[] = ids.consumables.map((row) => ({
    id: row.id,
    // An unmapped field lands with the potions rather than vanishing: a consumable the
    // panel cannot file is still a consumable the engine accepts, and hiding it would make
    // it unreachable except through the request drawer.
    group: BY_CONSUMES_FIELD[consumesField(row.sets)] ?? 'potion',
    graded: false,
    kind: 'consumable' as const,
  }));

  return [...buffs, ...consumables];
}

/**
 * IDS.md has no world-buff section yet (contract 1.7 adds one). Until it does, these are
 * the `IndividualBuffs` fields that are world buffs, read off IDS.md's own table by hand:
 * the four Dire Maul tribute buffs, the two capital-city buffs, Songflower and Sayge's.
 * `buildCatalogue` prefers the generated list whenever it is non-empty, so this disappears
 * the day the engine publishes the section, without a change to any caller.
 */
export const FALLBACK_WORLD_BUFFS: readonly string[] = [
  'fengus_ferocity',
  'moldars_moxie',
  'rallying_cry_of_the_dragonslayer',
  'sayges_fortune',
  'slipkiks_savvy',
  'songflower_serenade',
  'spirit_of_zandalar',
  'warchiefs_blessing',
];

const file = generated as SimIdsFile;

export const CATALOGUE: readonly BuffRow[] = buildCatalogue({
  ...file,
  worldBuffs: file.worldBuffs.length > 0 ? file.worldBuffs : [...FALLBACK_WORLD_BUFFS],
});

/** One group's rows, alphabetical by id so the panel's order never depends on IDS.md's. */
export function rowsIn(group: BuffGroupId): BuffRow[] {
  return CATALOGUE.filter((row) => row.group === group).sort((a, b) => a.id.localeCompare(b.id));
}

export type BuffGrade = 'off' | 'on' | 'improved';

export function gradeOf(selected: readonly string[], id: string): BuffGrade {
  if (selected.includes(`${id}${IMPROVED_SUFFIX}`)) return 'improved';
  return selected.includes(id) ? 'on' : 'off';
}

/**
 * Both forms of an id are removed and at most one is added back, so the list can never
 * carry `battle_shout` and `battle_shout:improved` at once -- a state `sim/request` would
 * resolve to whichever it read last, silently.
 */
export function setGrade(selected: readonly string[], id: string, grade: BuffGrade): string[] {
  const improved = `${id}${IMPROVED_SUFFIX}`;
  const without = selected.filter((entry) => entry !== id && entry !== improved);
  if (grade === 'off') return without;
  return [...without, grade === 'improved' ? improved : id];
}
```

- [ ] **Step 9: Write the failing stats test**

Create `src/lib/sim/stats.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import generated from '../../data/generated/sim-ids.json';
import { simCopy } from './copy';
import { PINNED_STATS, SIM_STATS, statLabel } from './stats';

describe('PINNED_STATS', () => {
  it('is contract 10.8’s list, verbatim and in its order', () => {
    expect(PINNED_STATS).toEqual([
      'strength',
      'agility',
      'stamina',
      'intellect',
      'spirit',
      'spell_power',
      'arcane_power',
      'fire_power',
      'frost_power',
      'holy_power',
      'nature_power',
      'shadow_power',
      'mp5',
      'hit',
      'crit',
      'spell_haste',
      'spell_penetration',
      'attack_power',
      'melee_haste',
      'armor_penetration',
      'expertise',
      'mana',
      'energy',
      'rage',
      'armor',
      'ranged_attack_power',
      'defense',
      'block',
      'block_value',
      'dodge',
      'parry',
      'health',
      'arcane_resistance',
      'fire_resistance',
      'frost_resistance',
      'nature_resistance',
      'shadow_resistance',
      'bonus_armor',
      'healing_power',
      'spell_damage',
      'feral_attack_power',
    ]);
  });

  it('carries one hit and one crit, and splits haste, which is 10.8’s whole point', () => {
    expect(PINNED_STATS).toContain('hit');
    expect(PINNED_STATS).toContain('crit');
    expect(PINNED_STATS).toContain('spell_haste');
    expect(PINNED_STATS).toContain('melee_haste');
    expect(PINNED_STATS).toContain('mp5');
    for (const forbidden of ['melee_hit', 'spell_hit', 'melee_crit', 'spell_crit', 'haste']) {
      expect(PINNED_STATS, forbidden).not.toContain(forbidden);
    }
  });

  it('names every one of them, so no weights row ever renders a raw id', () => {
    for (const stat of PINNED_STATS) {
      expect(simCopy.statLabel[stat], stat).toBeTruthy();
    }
  });

  it('names nothing 10.8 does not list, so a stale label cannot outlive its stat', () => {
    for (const named of Object.keys(simCopy.statLabel)) {
      expect(PINNED_STATS, named).toContain(named);
    }
  });
});

describe('SIM_STATS', () => {
  it('is the pinned list until IDS.md publishes a Stats section', () => {
    expect(SIM_STATS).toEqual(generated.stats.length > 0 ? generated.stats : PINNED_STATS);
  });

  it('agrees with the pinned list whenever IDS.md does publish one', () => {
    // The generated section and 10.8's pinning are the same vocabulary from two
    // directions; if they ever disagree, one of them is wrong and this says so loudly
    // rather than letting the page offer a stat the engine cannot weigh.
    if (generated.stats.length > 0) expect([...generated.stats].sort()).toEqual([...PINNED_STATS].sort());
  });
});

describe('statLabel', () => {
  it('prefers the copy table’s name', () => {
    expect(statLabel('attack_power')).toBe('Attack power');
    expect(statLabel('spell_haste')).toBe('Spell haste');
    expect(statLabel('mp5')).toBe('MP5');
  });

  it('humanises an id the copy table does not name, rather than showing the id raw', () => {
    expect(statLabel('some_new_stat')).toBe('Some new stat');
  });
});
```

- [ ] **Step 10: Write `stats.ts` and its copy**

Append to `simCopy`:

```ts
  /**
   * Stat names, for the weights page (part B). The vocabulary is contract 10.8's pinning
   * of the fork's `proto.Stat` enum in snake case: the engine carries ONE `hit` and ONE
   * `crit` -- there is no `melee_hit`, `spell_hit`, `melee_crit` or `spell_crit` -- while
   * haste IS split into `spell_haste` and `melee_haste`, and `MP5` is spelled `mp5`.
   * One key per id in `PINNED_STATS`, no more and no fewer; stats.test.ts asserts both
   * directions, so a renamed stat cannot leave a stale word behind.
   */
  statLabel: {
    strength: 'Strength',
    agility: 'Agility',
    stamina: 'Stamina',
    intellect: 'Intellect',
    spirit: 'Spirit',
    spell_power: 'Spell power',
    arcane_power: 'Arcane power',
    fire_power: 'Fire power',
    frost_power: 'Frost power',
    holy_power: 'Holy power',
    nature_power: 'Nature power',
    shadow_power: 'Shadow power',
    mp5: 'MP5',
    hit: 'Hit',
    crit: 'Crit',
    spell_haste: 'Spell haste',
    spell_penetration: 'Spell penetration',
    attack_power: 'Attack power',
    melee_haste: 'Melee haste',
    armor_penetration: 'Armor penetration',
    expertise: 'Expertise',
    mana: 'Mana',
    energy: 'Energy',
    rage: 'Rage',
    armor: 'Armor',
    ranged_attack_power: 'Ranged attack power',
    defense: 'Defense',
    block: 'Block',
    block_value: 'Block value',
    dodge: 'Dodge',
    parry: 'Parry',
    health: 'Health',
    arcane_resistance: 'Arcane resistance',
    fire_resistance: 'Fire resistance',
    frost_resistance: 'Frost resistance',
    nature_resistance: 'Nature resistance',
    shadow_resistance: 'Shadow resistance',
    bonus_armor: 'Bonus armor',
    healing_power: 'Healing power',
    spell_damage: 'Spell damage',
    feral_attack_power: 'Feral attack power',
  } as Record<string, string>,
```

Create `src/lib/sim/stats.ts`:

```ts
// web/src/lib/sim/stats.ts
// The stat vocabulary: the fork's `proto.Stat` enum names in snake case, pinned verbatim
// by contract 10.8 and generated into IDS.md's Stats section by the sim module (A7).
//
// Two things about this list are easy to get wrong and are wrong everywhere else:
//
//   * the engine carries ONE `hit` and ONE `crit`. There is no `melee_hit`, `spell_hit`,
//     `melee_crit` or `spell_crit`, and a page that offered either pair would be asking
//     for a weight the engine cannot compute;
//   * haste IS split -- `spell_haste` and `melee_haste` -- so there is no bare `haste`.
//
// `MP5` is spelled `mp5`. 10.8 also fixes the reference defaults the weights page starts
// on: `attack_power` for melee and hunters, `spell_power` for casters, served per spec as
// `reference_stat` on GET /v1/specs, so no page hard-codes one.
//
// This is here rather than in buffs.ts because a stat is not a buff: buffs.ts owns the
// panel's grouping and nothing else. Part B's /sim/weights is the only reader.
import generated from '../../data/generated/sim-ids.json';
import { simCopy } from './copy';

/** Contract 10.8's list, verbatim and in its order. */
export const PINNED_STATS: readonly string[] = [
  'strength',
  'agility',
  'stamina',
  'intellect',
  'spirit',
  'spell_power',
  'arcane_power',
  'fire_power',
  'frost_power',
  'holy_power',
  'nature_power',
  'shadow_power',
  'mp5',
  'hit',
  'crit',
  'spell_haste',
  'spell_penetration',
  'attack_power',
  'melee_haste',
  'armor_penetration',
  'expertise',
  'mana',
  'energy',
  'rage',
  'armor',
  'ranged_attack_power',
  'defense',
  'block',
  'block_value',
  'dodge',
  'parry',
  'health',
  'arcane_resistance',
  'fire_resistance',
  'frost_resistance',
  'nature_resistance',
  'shadow_resistance',
  'bonus_armor',
  'healing_power',
  'spell_damage',
  'feral_attack_power',
];

/**
 * The generated section when IDS.md has one, the pinned list until it does -- the same
 * rule `buffs.ts` uses for the world buffs, and for the same reason: the generator is the
 * long-term source of truth and the pinning is what makes the page correct today.
 * stats.test.ts asserts the two agree whenever both exist.
 */
const published = (generated as { stats: string[] }).stats;

export const SIM_STATS: readonly string[] = published.length > 0 ? published : PINNED_STATS;

/** "attack_power" reads as "Attack power". The copy table first, then plain casing. */
export function statLabel(id: string): string {
  const named = simCopy.statLabel[id];
  if (named !== undefined) return named;
  const words = id.split('_').join(' ');
  return words.charAt(0).toUpperCase() + words.slice(1);
}
```

- [ ] **Step 11: Run all three tests**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/ids-md.test.ts src/lib/sim/buffs.test.ts src/lib/sim/stats.test.ts
```

Expected: PASS. If "holds every id the engine publishes" fails, the missing id's `Consumes` field is not in `BY_CONSUMES_FIELD`; add the row rather than loosening the test.

- [ ] **Step 12: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint scripts/ids-md.mjs scripts/sync-sim-ids.mjs src/lib/sim/buffs.ts src/lib/sim/stats.ts src/lib/sim/buffs.test.ts src/lib/sim/stats.test.ts src/lib/sim/ids-md.test.ts && \
  npx prettier --check scripts src/lib/sim package.json && \
  git add -A scripts src/lib/sim package.json && \
  git commit -m "feat(sim): the engine's id vocabulary, generated, and the panel's grouping

sim/request/IDS.md is the only place the buff, consumable, profession
and stat vocabulary exists, so it is synced into
src/data/generated/sim-ids.json rather than copied by hand. buffs.ts
holds the table from the engine's message and Consumes field to a panel
section, folds the contract's <id>:improved rows into a three-way grade
on the plain id, and stats.ts carries contract A7's Stats section for
part B's weights page.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 9: Display names and icons for the ids

Design 4.3: "Each row is the id's display name and icon from the build". No such table exists in the build today, and this lane cannot invent one, so the loader is written the way `action-names.ts` already is: one optional file, and an honest fallback when it is absent — the humanised id, which is legible and true.

> **Ratified by contract 10.4.** `data/builds/<build>/simbuffs.json` is now a named data file with exactly the shape below — `{ "entries": { "<id>": { "name", "icon" } } }` for every IDS.md buff, debuff, world buff and consumable id — and the data lane owns it. The humanised fallback stays: the file is `required: false` in the sync, so the panel works before the data lane's first emission and degrades honestly if a build ever ships without it.

**Files:**
- Create: `src/lib/sim/buff-names.ts`
- Modify: `scripts/sync-data.mjs`
- Test: `src/lib/sim/buff-names.test.ts`

**Interfaces:**
- Consumes: `planner/load.ts`'s `dataUrl`, `fetchJson`.
- Produces: `buff-names.ts`: `interface BuffNames { entries: Record<string, { name: string; icon: string }> }`, `loadBuffNames(build: string): Promise<BuffNames>`, `EMPTY_BUFF_NAMES`, `buffLabel(id: string, names: BuffNames | null): string`, `buffIcon(build: string, id: string, names: BuffNames | null): string | null`

**The file shape (contract 10.4):** `data/builds/<build>/simbuffs.json`

```json
{ "entries": { "battle_shout": { "name": "Battle Shout", "icon": "ability_warrior_battleshout" } } }
```

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/buff-names.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import { EMPTY_BUFF_NAMES, buffIcon, buffLabel } from './buff-names';

const names = {
  entries: {
    battle_shout: { name: 'Battle Shout', icon: 'ability_warrior_battleshout' },
    'main_hand_imbue:shadow_oil': { name: 'Shadow Oil', icon: 'inv_potion_21' },
  },
};

describe('buffLabel', () => {
  it('is the build’s own name when it has one', () => {
    expect(buffLabel('battle_shout', names)).toBe('Battle Shout');
    expect(buffLabel('main_hand_imbue:shadow_oil', names)).toBe('Shadow Oil');
  });

  it('humanises the id when the build has no table, rather than showing “Unknown”', () => {
    expect(buffLabel('flask_of_supreme_power', EMPTY_BUFF_NAMES)).toBe('Flask of supreme power');
    expect(buffLabel('flask_of_supreme_power', null)).toBe('Flask of supreme power');
  });

  it('humanises a qualified id by its own half, keeping the qualifier', () => {
    expect(buffLabel('off_hand_imbue:frost_oil', EMPTY_BUFF_NAMES)).toBe('Off hand imbue: Frost oil');
  });

  it('leaves a client item id alone: there is nothing to humanise', () => {
    expect(buffLabel('item:13452', EMPTY_BUFF_NAMES)).toBe('item:13452');
  });
});

describe('buffIcon', () => {
  it('is the build’s icon path when the table names one', () => {
    expect(buffIcon('1.15.9', 'battle_shout', names)).toBe(
      '/data/1.15.9/icons/ability_warrior_battleshout.webp',
    );
  });

  it('is null when there is no icon, so the row renders without one', () => {
    expect(buffIcon('1.15.9', 'thorns', names)).toBeNull();
    expect(buffIcon('1.15.9', 'thorns', null)).toBeNull();
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/buff-names.test.ts
```

Expected: FAIL — `Failed to resolve import "./buff-names"`.

- [ ] **Step 3: Write `buff-names.ts`**

```ts
// web/src/lib/sim/buff-names.ts
// A display name and an icon for each id in the engine's vocabulary, from the build.
//
// The same shape and the same rule as action-names.ts: the build publishes a table, an id
// the table does not carry keeps a legible form of itself, and a build with no table at
// all renders humanised ids -- which is honest and is not worth an error banner on a page
// whose numbers are all correct. The table is `data/builds/<build>/simbuffs.json`
// (contract 10.4), owned by the data lane; until a build publishes one, every row falls
// back.
import { dataUrl, fetchJson } from '../planner/load';

export interface BuffNames {
  /** Engine id to its display name and the build's icon key (no extension). */
  entries: Record<string, { name: string; icon: string }>;
}

export const EMPTY_BUFF_NAMES: BuffNames = { entries: {} };

export function loadBuffNames(build: string): Promise<BuffNames> {
  return fetchJson<BuffNames>(dataUrl(build, 'simbuffs.json'));
}

/** "flask_of_supreme_power" reads as "Flask of supreme power". Casing, not translation. */
function humanise(snake: string): string {
  const words = snake.split('_').join(' ');
  return words.charAt(0).toUpperCase() + words.slice(1);
}

/**
 * An id's name. A qualified id ("main_hand_imbue:shadow_oil") humanises both halves and
 * keeps the qualifier, because the qualifier is the part that says which hand. A client
 * item id ("item:13452") is left exactly as it is: there is no word in it to raise, and
 * inventing "Item: 13452" would read as a name rather than as the raw id it is.
 */
export function buffLabel(id: string, names: BuffNames | null): string {
  const known = names?.entries[id]?.name;
  if (known !== undefined && known !== '') return known;
  const colon = id.indexOf(':');
  if (colon === -1) return humanise(id);
  const qualifier = id.slice(0, colon);
  const rest = id.slice(colon + 1);
  if (qualifier === 'item') return id;
  return `${humanise(qualifier)}: ${humanise(rest)}`;
}

/** The build's icon for an id, or null when there is none to draw. */
export function buffIcon(build: string, id: string, names: BuffNames | null): string | null {
  const icon = names?.entries[id]?.icon;
  return icon === undefined || icon === '' ? null : dataUrl(build, `icons/${icon}.webp`);
}
```

- [ ] **Step 4: Publish the file when the data lane has one**

In `scripts/sync-data.mjs`, add one row to `SYNC_ENTRIES`, beside `simconsumes.json`:

```js
  { name: 'simbuffs.json', kind: 'file', required: false },
```

`required: false` is the point: contract 10.4 names the file but the data lane emits it on its own schedule, so the panel works without it today and gains names and icons the day one appears.

- [ ] **Step 5: Run the test and the sync**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/buff-names.test.ts && FOREVER_DATA=fixture npm run sync:data
```

Expected: PASS, and the sync completes without complaining about the absent optional file.

- [ ] **Step 6: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && \
  npx eslint src/lib/sim/buff-names.ts src/lib/sim/buff-names.test.ts scripts/sync-data.mjs && \
  npx prettier --check src/lib/sim/buff-names.ts src/lib/sim/buff-names.test.ts scripts/sync-data.mjs && \
  git add -A src/lib/sim scripts && \
  git commit -m "feat(sim): display names and icons for the engine's buff ids

Reads contract 10.4's data/builds/<build>/simbuffs.json the way
action-names.ts reads its own table, and humanises the id when the
build has none. The file is optional in the sync, so the panel is not
blocked on the data lane's first emission.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 10: The full buff, debuff and consumable panel behind "Custom"

Design 4.3. Choosing "Custom" in the settings bar opens the whole list, grouped, each row the id's name and icon, graded ids as a three-way choice. Twelve groups is far too much to lay flat, so each group is the site's own `<details>` disclosure with the count of what is on in its summary — closed on arrival, so the panel costs one row of height until someone opens a section.

**Files:**
- Modify: `src/lib/sim/buffs.ts` (the selection helper)
- Create: `src/components/sim/BuffPanel.svelte`
- Modify: `src/lib/sim/store.svelte.ts` (load the names)
- Modify: `src/components/sim/SettingsBar.svelte`
- Modify: `src/components/sim/SimView.svelte`
- Test: `src/lib/sim/buffs.test.ts`, `tests/e2e/sim-buffs.spec.ts` (new)

**Interfaces:**
- Consumes: `buffs.ts` (Task 8), `buff-names.ts` (Task 9), `settings.ts` (Task 3).
- Produces:
  - `buffs.ts`: `interface Selection { buffs: string[]; consumables: string[] }`, `withGrade(selection: Selection, row: Pick<BuffRow, 'id' | 'kind'>, grade: BuffGrade): Selection`, `selectedIn(selection, group): string[]`
  - `store.svelte.ts`: `buffNames` getter
  - `BuffPanel.svelte`: props `{ settings: SimSettings; build: string; names: BuffNames | null; disabled: boolean; onchange: (next: SimSettings) => void }`
- Test ids added: `sim-buff-panel`, `sim-buff-group-<group>`, `sim-buff-<id>`.

- [ ] **Step 1: Write the failing selection test**

Append to `src/lib/sim/buffs.test.ts`:

```ts
describe('withGrade', () => {
  const empty = { buffs: [], consumables: [] };

  it('writes a buff into the buff list and a consumable into the consumable list', () => {
    expect(withGrade(empty, { id: 'battle_shout', kind: 'buff' }, 'on')).toEqual({
      buffs: ['battle_shout'],
      consumables: [],
    });
    expect(withGrade(empty, { id: 'flask_of_supreme_power', kind: 'consumable' }, 'on')).toEqual({
      buffs: [],
      consumables: ['flask_of_supreme_power'],
    });
  });

  it('carries the grade through to the id that is stored', () => {
    expect(withGrade(empty, { id: 'battle_shout', kind: 'buff' }, 'improved').buffs).toEqual([
      'battle_shout:improved',
    ]);
  });

  it('leaves the other list untouched', () => {
    const both = { buffs: ['thorns'], consumables: ['flask_of_supreme_power'] };
    expect(withGrade(both, { id: 'thorns', kind: 'buff' }, 'off')).toEqual({
      buffs: [],
      consumables: ['flask_of_supreme_power'],
    });
  });

  it('does not mutate the selection it was given', () => {
    const both = { buffs: ['thorns'], consumables: [] };
    withGrade(both, { id: 'thorns', kind: 'buff' }, 'off');
    expect(both.buffs).toEqual(['thorns']);
  });
});

describe('selectedIn', () => {
  it('counts what is on in one group, graded or plain', () => {
    const selection = { buffs: ['battle_shout:improved', 'thorns'], consumables: [] };
    expect(selectedIn(selection, 'raid-buffs').sort()).toEqual(['battle_shout', 'thorns']);
    expect(selectedIn(selection, 'debuffs')).toEqual([]);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/buffs.test.ts
```

Expected: FAIL — `withGrade is not a function`.

- [ ] **Step 3: Add the two helpers to `buffs.ts`**

```ts
/** The two id lists a `CharacterSpec` carries, and the shape the panel edits. */
export interface Selection {
  buffs: string[];
  consumables: string[];
}

/**
 * One row's grade, written into whichever of the two lists that row belongs to. The panel
 * spreads the answer over the settings object, so a component never decides which list an
 * id goes in -- the catalogue's `kind` does, and it came from IDS.md.
 */
export function withGrade(
  selection: Selection,
  row: Pick<BuffRow, 'id' | 'kind'>,
  grade: BuffGrade,
): Selection {
  if (row.kind === 'consumable') {
    return { ...selection, consumables: setGrade(selection.consumables, row.id, grade) };
  }
  return { ...selection, buffs: setGrade(selection.buffs, row.id, grade) };
}

/** The plain ids of one group that are on, for the group heading's count. */
export function selectedIn(selection: Selection, group: BuffGroupId): string[] {
  return rowsIn(group)
    .filter((row) => gradeOf(row.kind === 'consumable' ? selection.consumables : selection.buffs, row.id) !== 'off')
    .map((row) => row.id);
}
```

- [ ] **Step 4: Run the test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/buffs.test.ts
```

Expected: PASS.

- [ ] **Step 5: Write `BuffPanel.svelte`**

```svelte
<!-- web/src/components/sim/BuffPanel.svelte -->
<!-- Design 4.3, behind "Custom". Twelve sections, each the site's own <details>, each
     closed on arrival with the count of what is on in its summary: the whole vocabulary is
     several hundred rows and a flat list of them is not a control, it is a wall.

     Nothing here knows an id. The rows are buffs.ts's catalogue, which is IDS.md's own
     list read out of src/data/generated/sim-ids.json; the names and icons are the build's,
     with buff-names.ts's humanised fallback when the build has no table. A row this page
     could not file would still be reachable in the request drawer (Task 15), which is the
     escape hatch the design names. -->
<script lang="ts">
  import {
    BUFF_GROUPS,
    gradeOf,
    rowsIn,
    selectedIn,
    withGrade,
    type BuffGrade,
    type BuffGroupId,
  } from '../../lib/sim/buffs';
  import { buffIcon, buffLabel, type BuffNames } from '../../lib/sim/buff-names';
  import { simCopy } from '../../lib/sim/copy';
  import type { SimSettings } from '../../lib/sim/settings';

  let {
    settings,
    build,
    names,
    disabled,
    onchange,
  }: {
    settings: SimSettings;
    /** The character's data build, for the icon paths. */
    build: string;
    names: BuffNames | null;
    disabled: boolean;
    onchange: (next: SimSettings) => void;
  } = $props();

  const GRADES: BuffGrade[] = ['off', 'on', 'improved'];

  const selection = $derived({ buffs: settings.buffs, consumables: settings.consumables });

  function set(row: { id: string; kind: 'buff' | 'consumable' }, grade: BuffGrade): void {
    onchange({ ...settings, ...withGrade(selection, row, grade) });
  }

  function listOf(kind: 'buff' | 'consumable'): string[] {
    return kind === 'consumable' ? settings.consumables : settings.buffs;
  }

  function countIn(group: BuffGroupId): number {
    return selectedIn(selection, group).length;
  }
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-buff-panel"
>
  <h2 class="section-title text-[15px]">{simCopy.buffPanel}</h2>
  <p class="text-muted text-[12px]">{simCopy.buffPanelNote}</p>

  {#each BUFF_GROUPS as group (group)}
    {@const rows = rowsIn(group)}
    {#if rows.length > 0}
      <details class="border-line-soft rounded-panel border" data-testid={`sim-buff-group-${group}`}>
        <summary class="label text-nav flex min-h-11 cursor-pointer items-center gap-2 px-3 md:min-h-9">
          <span>{simCopy.buffGroupLabel[group] ?? group}</span>
          <span class="tabular text-muted font-mono text-[12px]">{countIn(group)}/{rows.length}</span>
        </summary>
        <ul class="grid grid-cols-1 gap-1 p-3 pt-0 md:grid-cols-2">
          {#each rows as row (row.id)}
            {@const label = buffLabel(row.id, names)}
            {@const icon = buffIcon(build, row.id, names)}
            {@const grade = gradeOf(listOf(row.kind), row.id)}
            <li class="flex min-h-11 items-center gap-2 text-[13px] md:min-h-9">
              {#if icon !== null}
                <!-- The name beside it carries the row, so the icon is decorative. -->
                <img src={icon} alt="" width="20" height="20" class="rounded-[2px]" loading="lazy" />
              {/if}
              {#if row.graded}
                <label class="flex min-w-0 flex-1 items-center gap-2">
                  <span class="truncate">{label}</span>
                  <select
                    class="border-line-warm rounded-control bg-raised text-text ml-auto min-h-11 border px-2 text-[12px] md:min-h-9"
                    {disabled}
                    aria-label={simCopy.gradeFor(label)}
                    value={grade}
                    onchange={(event) => set(row, event.currentTarget.value as BuffGrade)}
                    data-testid={`sim-buff-${row.id}`}
                  >
                    {#each GRADES as option (option)}
                      <option value={option}>{simCopy.gradeLabel[option]}</option>
                    {/each}
                  </select>
                </label>
              {:else}
                <label class="flex min-w-0 flex-1 items-center gap-2">
                  <input
                    type="checkbox"
                    class="accent-gold h-5 w-5 shrink-0"
                    {disabled}
                    checked={grade !== 'off'}
                    onchange={(event) => set(row, event.currentTarget.checked ? 'on' : 'off')}
                    data-testid={`sim-buff-${row.id}`}
                  />
                  <span class="truncate">{label}</span>
                </label>
              {/if}
            </li>
          {/each}
        </ul>
      </details>
    {/if}
  {/each}
</section>
```

- [ ] **Step 6: Load the names in the store**

In `src/lib/sim/store.svelte.ts`:

```ts
import { EMPTY_BUFF_NAMES, loadBuffNames, type BuffNames } from './buff-names';
```

```ts
  let buffNames = $state<BuffNames | null>(null);
  let loadedBuffNamesFor = '';
```

In `adopt()`, after `await ensureActionNames(...)`:

```ts
    await ensureBuffNames(outcome.character.tree_version);
```

and beside `ensureActionNames`:

```ts
  /**
   * The build's buff and consumable name table, once per build. A build with no table
   * renders humanised ids, which is legible and honest -- the same rule, and the same
   * reason, as `ensureActionNames` above.
   */
  async function ensureBuffNames(build: string): Promise<void> {
    if (build === '' || loadedBuffNamesFor === build) return;
    loadedBuffNamesFor = build;
    try {
      buffNames = await loadBuffNames(build);
    } catch {
      buffNames = EMPTY_BUFF_NAMES;
    }
  }
```

and the getter:

```ts
    get buffNames() {
      return buffNames;
    },
```

- [ ] **Step 7: Mount it**

In `SimView.svelte`, immediately after `<SettingsBar … />`:

```svelte
        {#if store.settings.preset === 'custom'}
          <BuffPanel
            settings={store.settings}
            build={store.character.tree_version}
            names={store.buffNames}
            disabled={store.phase === 'running' || store.serverRunning}
            onchange={(next) => store.setSettings(next)}
          />
        {/if}
```

with `import BuffPanel from './BuffPanel.svelte';`. In `SettingsBar.svelte`, drop `simCopy.customPresetNote` from the Custom option's `title` — the option keeps its label and loses the title attribute entirely — and then delete `customPresetNote` from `simCopy`. Its sentence promised this panel "arrives with Top Gear", and it has now arrived.

- [ ] **Step 8: Write the e2e**

Create `tests/e2e/sim-buffs.spec.ts`:

```ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('Custom opens the whole vocabulary, grouped, and every tick reaches the request', async ({
  page,
}) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  // The panel does not exist until Custom is chosen: a preset is a preset.
  await expect(page.getByTestId('sim-buff-panel')).toBeHidden();
  await page.getByTestId('sim-preset').selectOption('custom');
  await expect(page.getByTestId('sim-buff-panel')).toBeVisible();

  // Every group the design names has a section, and each one is closed on arrival.
  for (const group of [
    'raid-buffs',
    'party-buffs',
    'player-buffs',
    'world-buffs',
    'debuffs',
    'flask',
    'battle-elixir',
    'guardian-elixir',
    'food',
    'weapon-imbue',
    'potion',
    'explosive',
  ]) {
    await expect(page.getByTestId(`sim-buff-group-${group}`)).toBeVisible();
  }
  await expect(page.getByTestId('sim-buff-thorns')).toBeHidden();

  const raid = page.getByTestId('sim-buff-group-raid-buffs');
  await raid.locator('summary').click();
  await expect(page.getByTestId('sim-buff-thorns')).toBeVisible();
  await page.getByTestId('sim-buff-thorns').check();
  await expect(page.getByTestId('sim-buff-thorns')).toBeChecked();

  // A world buff is in its own section, not among the blessings.
  const world = page.getByTestId('sim-buff-group-world-buffs');
  await world.locator('summary').click();
  await expect(page.getByTestId('sim-buff-songflower_serenade')).toBeVisible();

  // The choices reach the engine: run at the fastest precision and read the request back
  // off the drawer-free path -- the saved request the result carries.
  await page.getByTestId('sim-precision').selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });
  const buffs = await page.evaluate(() => {
    const el = document.querySelector('[data-testid="sim-request-buffs"]');
    return el?.textContent ?? '';
  });
  expect(buffs === '' || buffs.includes('thorns')).toBe(true);
});
```

> The final assertion is deliberately loose until Task 15 puts the request on the page; it is tightened there. Leave the comment in.

- [ ] **Step 9: Run both suites**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-buffs.spec.ts --project=desktop
```

Expected: PASS.

- [ ] **Step 10: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && npx eslint src/lib/sim src/components/sim tests/e2e && \
  npx prettier --check src/lib/sim src/components/sim tests/e2e && git add -A src/lib/sim src/components/sim tests/e2e && \
  git commit -m "feat(sim): the whole buff, debuff and consumable panel behind Custom

Twelve grouped sections over IDS.md's own vocabulary, each a closed
disclosure with its count, graded ids as a three-way choice, and the
build's names and icons where it publishes them. Nothing here holds an
id: the catalogue is generated and the component only renders it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 11: Cooldown timing rows

Design 4.3's last sentence: "Cooldown timing for major cooldowns and potions (use on pull, at a time, at execute) rides along in the same panel; the engine's `Cooldowns` message already carries it." Contract 1.7 gives the shape: `CharacterSpec.Cooldowns []CooldownSpec`, where an empty `at_sec` means "on cooldown".

> **Ambiguity.** The web has no per-class table of major-cooldown spell ids — `spells.json` is a name table, not a cooldown list — so this task offers a timing row for every *selected* consumable in the Potions and Explosives groups, whose ids it does have, plus any `spell:<id>` row an edited or pasted request already carries, so those round-trip. Rows for class cooldowns arrive with a published list; the data lane is the owner.

**Files:**
- Create: `src/lib/sim/cooldowns.ts`
- Create: `src/components/sim/CooldownRows.svelte`
- Modify: `src/components/sim/BuffPanel.svelte`
- Modify: `src/lib/sim/character.ts` (carry the specs into `CharacterSpec`)
- Modify: `src/lib/sim/store.svelte.ts`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/cooldowns.test.ts`, `src/lib/sim/character.test.ts`

**Interfaces:**
- Consumes: `types.ts`'s `CooldownSpec`, `EncounterSpec` (Task 1); `buffs.ts`'s `rowsIn` (Task 8).
- Produces:
  - `cooldowns.ts`: `type CooldownMode = 'on-cooldown' | 'on-pull' | 'at-time' | 'at-execute'`, `COOLDOWN_MODES`, `interface CooldownRow { id: string; mode: CooldownMode; atSec: number }`, `executeStartSec(encounter)`, `specFor(id, mode, atSec, encounter): CooldownSpec`, `modeOf(spec, encounter): CooldownMode`, `rowsFor(ids, specs, encounter): CooldownRow[]`, `withCooldown(specs, id, mode, atSec, encounter): CooldownSpec[]`
  - `character.ts`: `toCharacterSpec(character, index, buffs, consumes, cooldowns?)`

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/cooldowns.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import {
  COOLDOWN_MODES,
  executeStartSec,
  modeOf,
  rowsFor,
  specFor,
  withCooldown,
} from './cooldowns';
import { DEFAULT_ENCOUNTER } from './types';

const encounter = { ...DEFAULT_ENCOUNTER, duration_sec: 200, execute_ratio: 0.25 };

describe('executeStartSec', () => {
  it('is where the execute window opens: the fight’s length less its execute share', () => {
    expect(executeStartSec(encounter)).toBe(150);
  });

  it('is the end of the fight when there is no execute window at all', () => {
    expect(executeStartSec({ ...encounter, execute_ratio: 0 })).toBe(200);
  });
});

describe('specFor and modeOf', () => {
  it('turns each mode into the times the engine takes', () => {
    expect(specFor('major_mana_potion', 'on-cooldown', 0, encounter).at_sec).toEqual([]);
    expect(specFor('major_mana_potion', 'on-pull', 0, encounter).at_sec).toEqual([0]);
    expect(specFor('major_mana_potion', 'at-time', 42, encounter).at_sec).toEqual([42]);
    expect(specFor('major_mana_potion', 'at-execute', 0, encounter).at_sec).toEqual([150]);
  });

  it('clamps a hand-typed time into the fight', () => {
    expect(specFor('x', 'at-time', -5, encounter).at_sec).toEqual([0]);
    expect(specFor('x', 'at-time', 9999, encounter).at_sec).toEqual([200]);
  });

  it('reads a stored spec back as the mode that wrote it', () => {
    expect(modeOf({ id: 'x', at_sec: [] }, encounter)).toBe('on-cooldown');
    expect(modeOf({ id: 'x', at_sec: [0] }, encounter)).toBe('on-pull');
    expect(modeOf({ id: 'x', at_sec: [150] }, encounter)).toBe('at-execute');
    expect(modeOf({ id: 'x', at_sec: [42] }, encounter)).toBe('at-time');
  });

  it('reads a multi-time spec as a plain time rather than claiming a mode it is not', () => {
    expect(modeOf({ id: 'x', at_sec: [0, 90] }, encounter)).toBe('at-time');
  });

  it('offers the four modes the design names', () => {
    expect(COOLDOWN_MODES).toEqual(['on-cooldown', 'on-pull', 'at-time', 'at-execute']);
  });
});

describe('rowsFor', () => {
  it('is one row per id, defaulting to on cooldown', () => {
    expect(rowsFor(['a', 'b'], [], encounter)).toEqual([
      { id: 'a', mode: 'on-cooldown', atSec: 0 },
      { id: 'b', mode: 'on-cooldown', atSec: 0 },
    ]);
  });

  it('takes a row’s mode and time from the stored spec', () => {
    expect(rowsFor(['a'], [{ id: 'a', at_sec: [42] }], encounter)).toEqual([
      { id: 'a', mode: 'at-time', atSec: 42 },
    ]);
  });

  it('keeps a stored spec whose id is not offered, so a pasted request round-trips', () => {
    expect(rowsFor(['a'], [{ id: 'spell:11305', at_sec: [0] }], encounter)).toEqual([
      { id: 'a', mode: 'on-cooldown', atSec: 0 },
      { id: 'spell:11305', mode: 'on-pull', atSec: 0 },
    ]);
  });
});

describe('withCooldown', () => {
  it('drops the spec entirely for “on cooldown”, which is the engine’s own default', () => {
    expect(withCooldown([{ id: 'a', at_sec: [0] }], 'a', 'on-cooldown', 0, encounter)).toEqual([]);
  });

  it('replaces a spec rather than appending a second one for the same id', () => {
    const once = withCooldown([], 'a', 'on-pull', 0, encounter);
    const twice = withCooldown(once, 'a', 'at-time', 60, encounter);
    expect(twice).toEqual([{ id: 'a', at_sec: [60] }]);
  });

  it('does not mutate the list it was given', () => {
    const specs = [{ id: 'a', at_sec: [0] }];
    withCooldown(specs, 'a', 'at-time', 10, encounter);
    expect(specs).toEqual([{ id: 'a', at_sec: [0] }]);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/cooldowns.test.ts
```

Expected: FAIL — `Failed to resolve import "./cooldowns"`.

- [ ] **Step 3: Write `cooldowns.ts`**

```ts
// web/src/lib/sim/cooldowns.ts
// Contract 1.7's CooldownSpec, as four choices a player recognises.
//
// The engine takes a list of seconds and an empty list means "on cooldown". "At execute"
// is not a fifth thing the engine knows -- it is the second the execute window opens,
// computed here from the encounter, so changing the fight length moves it with the fight
// rather than leaving a potion at a timestamp that is now past the pull.
import type { CooldownSpec, EncounterSpec } from './types';

export type CooldownMode = 'on-cooldown' | 'on-pull' | 'at-time' | 'at-execute';

export const COOLDOWN_MODES: readonly CooldownMode[] = [
  'on-cooldown',
  'on-pull',
  'at-time',
  'at-execute',
];

export interface CooldownRow {
  id: string;
  mode: CooldownMode;
  /** Meaningful for `at-time`; 0 otherwise. */
  atSec: number;
}

/** Where the execute window opens. With no execute window, the end of the fight. */
export function executeStartSec(encounter: EncounterSpec): number {
  return Math.round(encounter.duration_sec * (1 - encounter.execute_ratio));
}

export function specFor(
  id: string,
  mode: CooldownMode,
  atSec: number,
  encounter: EncounterSpec,
): CooldownSpec {
  switch (mode) {
    case 'on-cooldown':
      return { id, at_sec: [] };
    case 'on-pull':
      return { id, at_sec: [0] };
    case 'at-execute':
      return { id, at_sec: [executeStartSec(encounter)] };
    default:
      return { id, at_sec: [Math.min(encounter.duration_sec, Math.max(0, Math.round(atSec)))] };
  }
}

/**
 * The mode a stored spec was written by. A list of more than one time is a request the
 * panel cannot express, so it reads as `at-time` on its first entry rather than being
 * silently rewritten -- the drawer (Task 15) is where a multi-use schedule is edited.
 */
export function modeOf(spec: CooldownSpec, encounter: EncounterSpec): CooldownMode {
  if (spec.at_sec.length === 0) return 'on-cooldown';
  if (spec.at_sec.length === 1) {
    if (spec.at_sec[0] === 0) return 'on-pull';
    if (spec.at_sec[0] === executeStartSec(encounter)) return 'at-execute';
  }
  return 'at-time';
}

/**
 * A row for every id offered, plus a row for every stored spec whose id is not among them
 * -- so a request pasted into the drawer with a class cooldown in it keeps that cooldown
 * when the panel re-renders, instead of the panel quietly dropping what it cannot offer.
 */
export function rowsFor(
  ids: readonly string[],
  specs: readonly CooldownSpec[],
  encounter: EncounterSpec,
): CooldownRow[] {
  const byId = new Map(specs.map((spec) => [spec.id, spec]));
  const offered: CooldownRow[] = ids.map((id) => {
    const spec = byId.get(id);
    if (spec === undefined) return { id, mode: 'on-cooldown', atSec: 0 };
    return { id, mode: modeOf(spec, encounter), atSec: spec.at_sec[0] ?? 0 };
  });
  const extra: CooldownRow[] = specs
    .filter((spec) => !ids.includes(spec.id))
    .map((spec) => ({ id: spec.id, mode: modeOf(spec, encounter), atSec: spec.at_sec[0] ?? 0 }));
  return [...offered, ...extra];
}

/**
 * One row's answer, written into the spec list. "On cooldown" removes the spec rather than
 * storing an empty one: it is the engine's own default, and a request that spells out
 * every default is a request nobody can read in the drawer.
 */
export function withCooldown(
  specs: readonly CooldownSpec[],
  id: string,
  mode: CooldownMode,
  atSec: number,
  encounter: EncounterSpec,
): CooldownSpec[] {
  const without = specs.filter((spec) => spec.id !== id);
  if (mode === 'on-cooldown') return without;
  return [...without, specFor(id, mode, atSec, encounter)];
}
```

- [ ] **Step 4: Run it to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/cooldowns.test.ts
```

Expected: PASS, 12 tests.

- [ ] **Step 5: Carry the specs into the request**

Append to `src/lib/sim/character.test.ts`:

```ts
describe('toCharacterSpec cooldowns', () => {
  it('omits the key entirely when there are none, rather than sending an empty list', () => {
    const spec = toCharacterSpec(fixtureCharacter, fixtureIndex, [], [], []);
    expect(spec.cooldowns).toBeUndefined();
  });

  it('carries the rows through, copied rather than shared', () => {
    const rows = [{ id: 'major_mana_potion', at_sec: [0] }];
    const spec = toCharacterSpec(fixtureCharacter, fixtureIndex, [], [], rows);
    expect(spec.cooldowns).toEqual(rows);
    expect(spec.cooldowns).not.toBe(rows);
  });
});
```

(`fixtureCharacter` and `fixtureIndex` are the file's existing helpers; reuse whatever it already builds for its `toCharacterSpec` cases.)

Then in `src/lib/sim/character.ts`:

```ts
export function toCharacterSpec(
  character: SimCharacter,
  index: TalentIndex,
  buffs: string[],
  consumes: string[],
  cooldowns: readonly CooldownSpec[] = [],
): CharacterSpec {
  const spec: CharacterSpec = {
    name: character.name,
    race: character.race_slug,
    class: character.class_slug,
    level: SIM_LEVEL,
    talents: talentsString(index, character.point_order),
    gear: gearSlots(character.gear),
    buffs: [...buffs],
    consumes: [...consumes],
  };
  // Omitted rather than sent empty, for the same reason `professions` is: an empty list
  // would claim we had scheduled something and found nothing, when the truth is "every
  // cooldown on cooldown", which is what an absent field means to the engine.
  return cooldowns.length === 0 ? spec : { ...spec, cooldowns: cooldowns.map((row) => ({ ...row, at_sec: [...row.at_sec] })) };
}
```

with `import type { CharacterSource, CharacterSpec, CooldownSpec, GearSlot } from './types';`.

In `store.svelte.ts`, both `run()` and `runOnServer()` pass the settings' rows:

```ts
        character: toCharacterSpec(character, index, settings.buffs, settings.consumables, settings.cooldowns),
```

- [ ] **Step 6: Add the copy**

```ts
  // --- Design 4.3: the cooldown timing rows. ---
  cooldownTiming: 'When to use them',
  cooldownNote:
    'Only what is ticked above, plus anything a pasted request already schedules. Class cooldowns arrive when the build publishes their spell ids.',
  cooldownModeLabel: {
    'on-cooldown': 'On cooldown',
    'on-pull': 'On the pull',
    'at-time': 'At a time',
    'at-execute': 'At execute',
  } as Record<string, string>,
  cooldownAt: 'Second',
```

- [ ] **Step 7: Write `CooldownRows.svelte`**

```svelte
<!-- web/src/components/sim/CooldownRows.svelte -->
<!-- Design 4.3's last sentence, inside the same panel. One row per scheduled thing: the
     mode, and a seconds field only where the mode needs one -- an always-visible number
     input next to "On cooldown" would invite a time that is never read. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { COOLDOWN_MODES, rowsFor, withCooldown, type CooldownMode } from '../../lib/sim/cooldowns';
  import { buffLabel, type BuffNames } from '../../lib/sim/buff-names';
  import type { SimSettings } from '../../lib/sim/settings';

  let {
    settings,
    ids,
    names,
    disabled,
    onchange,
  }: {
    settings: SimSettings;
    /** The ids a row is offered for: whatever is ticked in Potions and Explosives. */
    ids: string[];
    names: BuffNames | null;
    disabled: boolean;
    onchange: (next: SimSettings) => void;
  } = $props();

  const rows = $derived(rowsFor(ids, settings.cooldowns, settings.encounter));

  function set(id: string, mode: CooldownMode, atSec: number): void {
    onchange({
      ...settings,
      cooldowns: withCooldown(settings.cooldowns, id, mode, atSec, settings.encounter),
    });
  }

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 border px-2 text-[12px] md:min-h-9';
</script>

{#if rows.length > 0}
  <details class="border-line-soft rounded-panel border" data-testid="sim-cooldowns">
    <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-3 md:min-h-9">
      {simCopy.cooldownTiming}
    </summary>
    <div class="flex flex-col gap-2 p-3 pt-0">
      <p class="text-muted text-[12px]">{simCopy.cooldownNote}</p>
      <ul class="flex flex-col gap-1">
        {#each rows as row (row.id)}
          <li class="flex min-h-11 flex-wrap items-center gap-2 text-[13px] md:min-h-9">
            <span class="min-w-0 flex-1 truncate">{buffLabel(row.id, names)}</span>
            <select
              class={control}
              {disabled}
              value={row.mode}
              onchange={(event) => set(row.id, event.currentTarget.value as CooldownMode, row.atSec)}
              data-testid={`sim-cooldown-${row.id}`}
            >
              {#each COOLDOWN_MODES as mode (mode)}
                <option value={mode}>{simCopy.cooldownModeLabel[mode]}</option>
              {/each}
            </select>
            {#if row.mode === 'at-time'}
              <input
                type="number"
                min="0"
                max={settings.encounter.duration_sec}
                step="1"
                class={`${control} w-20`}
                {disabled}
                aria-label={simCopy.cooldownAt}
                value={String(row.atSec)}
                onchange={(event) => set(row.id, 'at-time', Number(event.currentTarget.value))}
                data-testid={`sim-cooldown-at-${row.id}`}
              />
            {/if}
          </li>
        {/each}
      </ul>
    </div>
  </details>
{/if}
```

- [ ] **Step 8: Mount it in `BuffPanel.svelte`**

Add the import and, immediately before the closing `</section>`:

```svelte
  <CooldownRows
    {settings}
    {names}
    {disabled}
    ids={[...selectedIn(selection, 'potion'), ...selectedIn(selection, 'explosive')]}
    {onchange}
  />
```

with `import CooldownRows from './CooldownRows.svelte';`.

- [ ] **Step 9: Run the unit suite**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run
```

Expected: PASS.

- [ ] **Step 10: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && npx eslint src/lib/sim src/components/sim && \
  npx prettier --check src/lib/sim src/components/sim && git add -A src/lib/sim src/components/sim && \
  git commit -m "feat(sim): cooldown timing rows in the buff panel

Contract 1.7's CooldownSpec as four choices: on cooldown, on the pull,
at a time, at execute. \"At execute\" is computed from the encounter, so
changing the fight length moves it with the fight. A row is offered for
every ticked potion and explosive, and a spec the panel cannot offer is
kept rather than dropped, so a pasted request round-trips.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 12: Buff uptime counts, and the sample-iteration log

Design 5.1's second and third bullets.

**Buff uptime with counts.** `report/AuraTable.svelte` already renders an "Applied" column from `track.applications`, and `SimResults.svelte` already mounts it for both the Buffs and Debuffs tabs — so the column the design asks for exists and the work is to *prove* it, by giving the cell the testid the sim's own e2e can pin (the uptime cell has one; the count cell does not).

**Sample iteration log.** `result.sample` (contract 2, as amended by **A12**) is one iteration's casts in order, pre-pull separated, with the resources after each cast. Each row carries an *action key* — `spell:23881`, `item:13503`, `other:melee` — and neither a display name nor a spell id, so the table resolves names with `resolveActionName` exactly as the cast table does and the two can never disagree about what an action is called. A new results tab renders it.

**Files:**
- Create: `src/lib/sim/sample-log.ts`
- Create: `src/components/sim/SampleLog.svelte`
- Create: `src/fixtures/sim/sample.json`
- Modify: `src/components/report/AuraTable.svelte`
- Modify: `src/components/sim/SimResults.svelte`
- Modify: `src/fixtures/sim/engine-fake.ts`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/sample-log.test.ts`, `tests/e2e/sim-results.spec.ts`

**Interfaces:**
- Consumes: `types.ts`'s `SampleCast` (Task 1), `action-names.ts`'s `resolveActionName`.
- Produces:
  - `sample-log.ts`: `interface SampleRow`, `interface SampleLog { rows: SampleRow[]; columns: string[]; prePullCount: number }`, `RESOURCE_ORDER`, `sampleTime(atMs: number): string`, `sampleLog(sample: readonly SampleCast[] | undefined, names: ActionNames | null): SampleLog`
  - `SampleLog.svelte`: props `{ sample: SampleCast[] | undefined; actionNames: ActionNames | null }`, testid `sim-sample-log`
- Test id added: `aura-applied` on `AuraTable`'s count cell.

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/sample-log.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import { RESOURCE_ORDER, sampleLog, sampleTime } from './sample-log';
import type { SampleCast } from './types';

const names = { spell: { '25286': 'Heroic Strike', '11305': 'Bloodrage' }, item: {} };

const sample: SampleCast[] = [
  { at_ms: -1500, action: 'spell:11305', resources: { rage: 0 } },
  { at_ms: 320, action: 'spell:25286', target: 'Target', resources: { rage: 42 } },
  { at_ms: 1900, action: 'spell:25286', target: 'Target', resources: { rage: 12, mana: 300 } },
];

describe('sampleTime', () => {
  it('is a signed count of seconds, one decimal', () => {
    expect(sampleTime(-1500)).toBe('-1.5 s');
    expect(sampleTime(0)).toBe('0.0 s');
    expect(sampleTime(1900)).toBe('1.9 s');
  });
});

describe('sampleLog', () => {
  const log = sampleLog(sample, names);

  it('resolves every action key through the build’s table, never rendering the key', () => {
    expect(log.rows.map((row) => row.name)).toEqual(['Bloodrage', 'Heroic Strike', 'Heroic Strike']);
  });

  it('keeps an “other” key readable without a table, the way every cast row does', () => {
    expect(sampleLog([{ at_ms: 0, action: 'other:melee' }], null).rows[0].name).toBe('Melee');
  });

  it('marks the pre-pull casts and counts them, so the table can rule a line under them', () => {
    expect(log.rows.map((row) => row.prePull)).toEqual([true, false, false]);
    expect(log.prePullCount).toBe(1);
  });

  it('keeps the engine’s order rather than sorting, because the order is the point', () => {
    expect(log.rows.map((row) => row.atMs)).toEqual([-1500, 320, 1900]);
  });

  it('gives every row a key of its own, so two casts of one spell at one instant still render', () => {
    const doubled = sampleLog([sample[1], sample[1]], names);
    expect(new Set(doubled.rows.map((row) => row.key)).size).toBe(2);
  });

  it('columns are the resources that appear, in the engine’s own order', () => {
    expect(log.columns).toEqual(['mana', 'rage']);
    expect(RESOURCE_ORDER.indexOf('mana')).toBeLessThan(RESOURCE_ORDER.indexOf('rage'));
  });

  it('puts a resource nothing anticipated after the known ones rather than dropping it', () => {
    const odd = sampleLog(
      [{ at_ms: 0, action: 'spell:1', resources: { rage: 1, zeal: 2 } }],
      names,
    );
    expect(odd.columns).toEqual(['rage', 'zeal']);
  });

  it('is empty, not a throw, for a result with no sample at all', () => {
    expect(sampleLog(undefined, names)).toEqual({ rows: [], columns: [], prePullCount: 0 });
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/sample-log.test.ts
```

Expected: FAIL — `Failed to resolve import "./sample-log"`.

- [ ] **Step 3: Write `sample-log.ts`**

```ts
// web/src/lib/sim/sample-log.ts
// Design 5.1: "one iteration's casts in order with the pre-pull section separated,
// resources at each cast, and the note that it is one iteration and not a guide".
//
// The order is the engine's and is never sorted here: a cast log read out of order is not
// a cast log.
//
// Contract A12: a row carries an action key (`spell:23881`, `item:13503`, `other:melee`)
// and nothing else. Names go through resolveActionName like every other row on this lane,
// so a key never reaches the table and the sample and cast tables can never name one
// action two different ways.
import { resolveActionName, type ActionNames } from './action-names';
import type { SampleCast } from './types';

/**
 * The engine's resources, in the order a player reads them: the caster's pool first, then
 * the melee ones, then combo points. A resource this list does not anticipate is appended
 * rather than dropped -- the engine may well gain one, and a missing column would silently
 * hide it.
 */
export const RESOURCE_ORDER: readonly string[] = ['mana', 'energy', 'rage', 'focus', 'combo_points'];

export interface SampleRow {
  /** A key of its own: two casts can share an instant and a spell id. */
  key: string;
  atMs: number;
  /** True while the fight has not started. */
  prePull: boolean;
  time: string;
  name: string;
  target: string;
  resources: Record<string, number>;
}

export interface SampleLog {
  rows: SampleRow[];
  /** The resource columns this sample actually has, in RESOURCE_ORDER then alphabetical. */
  columns: string[];
  prePullCount: number;
}

/** "-1.5 s" before the pull, "1.9 s" after it. One decimal, because the engine's own ticks are 10ms. */
export function sampleTime(atMs: number): string {
  return `${(atMs / 1000).toFixed(1)} s`;
}

export function sampleLog(
  sample: readonly SampleCast[] | undefined,
  names: ActionNames | null,
): SampleLog {
  if (sample === undefined || sample.length === 0) return { rows: [], columns: [], prePullCount: 0 };

  const seen = new Set<string>();
  for (const cast of sample) {
    for (const key of Object.keys(cast.resources ?? {})) seen.add(key);
  }
  const known = RESOURCE_ORDER.filter((key) => seen.has(key));
  const extra = [...seen].filter((key) => !RESOURCE_ORDER.includes(key)).sort();

  const rows = sample.map((cast, index) => ({
    // The index is in the key because two casts can genuinely share an instant and an
    // action, and Svelte 5 throws on a repeated {#each} key.
    key: `${index}-${cast.at_ms}-${cast.action}`,
    atMs: cast.at_ms,
    prePull: cast.at_ms < 0,
    time: sampleTime(cast.at_ms),
    name: resolveActionName(cast.action, names),
    target: cast.target ?? '',
    resources: { ...(cast.resources ?? {}) },
  }));

  return { rows, columns: [...known, ...extra], prePullCount: rows.filter((row) => row.prePull).length };
}
```

- [ ] **Step 4: Run the test to verify it passes**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/sample-log.test.ts
```

Expected: PASS, 8 tests.

- [ ] **Step 5: Add the copy**

```ts
  // --- Design 5.1: the sample iteration log. ---
  tabSample: 'One iteration',
  sampleNote:
    'One iteration’s casts, in order. It is a sample of what the rotation did once, not a rotation guide, and the next iteration is a different fight.',
  sampleEmpty: 'This result carries no sample iteration.',
  samplePrePull: 'Before the pull',
  sampleTimeHeading: 'At',
  sampleCastHeading: 'Cast',
  sampleTargetHeading: 'On',
  resourceLabel: {
    mana: 'Mana',
    energy: 'Energy',
    rage: 'Rage',
    focus: 'Focus',
    combo_points: 'Combo',
  } as Record<string, string>,
```

- [ ] **Step 6: Write `SampleLog.svelte`**

```svelte
<!-- web/src/components/sim/SampleLog.svelte -->
<!-- Design 5.1. One iteration, in order, with the pre-pull ruled off from the fight and
     the resources after each cast. The note above it is not decoration: a table of one
     iteration's casts looks exactly like a rotation guide, and it is not one. -->
<script lang="ts">
  import type { ActionNames } from '../../lib/sim/action-names';
  import { simCopy } from '../../lib/sim/copy';
  import { sampleLog } from '../../lib/sim/sample-log';
  import type { SampleCast } from '../../lib/sim/types';

  let {
    sample,
    actionNames,
  }: { sample: SampleCast[] | undefined; actionNames: ActionNames | null } = $props();

  const log = $derived(sampleLog(sample, actionNames));
</script>

{#if log.rows.length === 0}
  <p class="text-muted p-4 text-[14px]" data-testid="sim-sample-empty">{simCopy.sampleEmpty}</p>
{:else}
  <div class="flex flex-col gap-2 p-2" data-testid="sim-sample-log">
    <p class="text-muted px-2 text-[12px]">{simCopy.sampleNote}</p>
    <!-- The resource columns make this the one table on the lane that can genuinely be
         wider than a phone, so it gets its own scroller rather than the page getting one. -->
    <div class="overflow-x-auto">
      <table class="w-full min-w-[420px] text-[13px]">
        <thead>
          <tr class="text-muted label">
            <th class="px-2 py-1 text-left">{simCopy.sampleTimeHeading}</th>
            <th class="px-2 py-1 text-left">{simCopy.sampleCastHeading}</th>
            <th class="px-2 py-1 text-left">{simCopy.sampleTargetHeading}</th>
            {#each log.columns as column (column)}
              <th class="px-2 py-1 text-right">{simCopy.resourceLabel[column] ?? column}</th>
            {/each}
          </tr>
        </thead>
        <tbody>
          {#each log.rows as row, index (row.key)}
            <!-- One rule, under the last pre-pull cast: the pull is the only boundary in
                 this table and a heading row would break the column alignment to say it. -->
            <tr
              class={`border-line-soft border-b ${
                index + 1 === log.prePullCount ? 'border-b-line-warm-strong' : ''
              }`}
              data-testid={`sim-sample-row-${row.key}`}
            >
              <td class="tabular px-2 py-1 font-mono">
                {row.time}{#if row.prePull}<span class="text-muted ml-1 text-[11px]"
                    >{simCopy.samplePrePull}</span
                  >{/if}
              </td>
              <td class="px-2 py-1 font-semibold">{row.name}</td>
              <td class="text-muted px-2 py-1">{row.target}</td>
              {#each log.columns as column (column)}
                <td class="tabular px-2 py-1 text-right font-mono">{row.resources[column] ?? ''}</td>
              {/each}
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}
```

- [ ] **Step 7: Add the tab, and pass the sample down**

In `SimResults.svelte`: add `{ id: 'sample', label: simCopy.tabSample }` to `TABS` after `timeline`; add a `sample` prop (`sample: SampleCast[] | undefined`); import `SampleLog`; and add the branch before the final `{:else}`:

```svelte
    {:else if tab === 'sample'}
      <SampleLog {sample} {actionNames} />
```

In `SimView.svelte` and `SavedSim.svelte`, pass `sample={result.sample}` (respectively `store.result.sample`) to the lazy `SimResults`.

- [ ] **Step 8: Give the count cell a testid**

In `src/components/report/AuraTable.svelte`, the applications cell becomes:

```svelte
          <span class="text-muted tabular text-right font-mono text-[13px]"
            ><span data-testid="aura-applied">{track.applications}</span><span
              class="label font-body ml-1.5 md:hidden">applied</span
            ></span
          >
```

The count column itself already existed; this only lets a test pin the number rather than the cell's whole text, matching what `aura-uptime` beside it already does.

- [ ] **Step 9: Give the fake engine a sample**

Create `src/fixtures/sim/sample.json`:

```json
[
  { "at_ms": -3000, "action": "spell:11305", "resources": { "rage": 0 } },
  { "at_ms": -1000, "action": "spell:6673", "resources": { "rage": 10 } },
  { "at_ms": 150, "action": "spell:25286", "target": "Target", "resources": { "rage": 31 } },
  { "at_ms": 1600, "action": "other:melee", "target": "Target", "resources": { "rage": 9 } },
  { "at_ms": 3050, "action": "spell:25286", "target": "Target", "resources": { "rage": 28 } },
  { "at_ms": 4400, "action": "spell:11605", "target": "Target", "resources": { "rage": 7 } }
]
```

In `engine-fake.ts`, import it and attach it to both the per-shard result and the combined one, so every path the page can reach carries a sample:

```ts
import sampleJson from './sample.json';

const fixtureSample = sampleJson as unknown as SimResult['sample'];
```

then add `sample: fixtureSample,` beside `summary: fixture.summary,` in `simRun`'s returned object and in `simCombine`'s.

- [ ] **Step 10: Extend the e2e**

Append to `tests/e2e/sim-results.spec.ts`:

```ts
test('the buffs tab counts applications and the sample tab shows one iteration', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await page.getByTestId('sim-precision').selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-results')).toBeVisible({ timeout: 30_000 });

  await page.getByTestId('sim-tab-buffs').click();
  await expect(page.getByTestId('aura-applied').first()).toHaveText(/^\d+$/);

  await page.getByTestId('sim-tab-sample').click();
  const log = page.getByTestId('sim-sample-log');
  await expect(log).toBeVisible();
  // The fixture's first two casts are before the pull and are labelled as such.
  await expect(log.getByText('Before the pull').first()).toBeVisible();
  // Names, never engine keys.
  await expect(log).not.toContainText('spell:25286');
  // Resources after each cast: the fixture's only pool is rage.
  await expect(log.locator('th', { hasText: 'Rage' })).toHaveCount(1);
});
```

- [ ] **Step 11: Run both suites**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-results.spec.ts --project=desktop
```

Expected: PASS. If the sample tab overflows the phone audit, its own `overflow-x-auto` wrapper is the fix and the page must never gain one.

- [ ] **Step 12: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && npx eslint src components tests 2>/dev/null; npx eslint src/lib/sim src/components tests/e2e && \
  npx prettier --check src/lib/sim src/components src/fixtures/sim tests/e2e && \
  git add -A src/lib/sim src/components src/fixtures/sim tests/e2e && \
  git commit -m "feat(sim): the sample-iteration log, and a testid on the aura count

Design 5.1's sample log as a results tab: one iteration's casts in the
engine's own order, the pre-pull ruled off, and the resources after
each cast in their own columns. Each row carries contract A12's action
key and the name is resolved with resolveActionName, so the sample and
cast tables cannot name one action two ways. The aura table already
carried an Applied count for both the buff and debuff tabs; it gains
the testid its uptime cell already had.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 13: The rotation card, with the fidelity note

Design 5.1's fourth bullet: "the rotation the run used with a link to its page, and for a non-validated spec the fidelity note". `/sim` already renders a one-line fidelity note under the settings bar (`spec-fidelity-note`); a *result* needs the same fact beside it, because a saved sim and a shared link have no settings bar at all.

**Files:**
- Create: `src/components/sim/RotationCard.svelte`
- Modify: `src/components/sim/SimView.svelte`
- Modify: `src/components/sim/SavedSim.svelte`
- Modify: `src/lib/sim/api.ts`, `src/components/sim/SavedSim.svelte` (the saved page needs the spec rows)
- Modify: `src/lib/sim/copy.ts`
- Test: `tests/e2e/sim-saved.spec.ts`

**Interfaces:**
- Consumes: `spec-state.ts`'s `mergeSpecRows`, `specPillClass`, `specStateLabel`, `specStateNote`; `spec-label.ts`'s `specRow`; `api.ts`'s `fetchSpecs`.
- Produces: `RotationCard.svelte` — props `{ spec: string; fidelity: SpecFidelity | null }`, testid `sim-rotation-card`.

- [ ] **Step 1: Write the failing e2e**

Append to `tests/e2e/sim-saved.spec.ts`:

```ts
test('a saved sim names the rotation it used and carries its fidelity', async ({ page }) => {
  await page.goto(`/sim/${SAVED_ID}`);
  const card = page.getByTestId('sim-rotation-card');
  await expect(card).toBeVisible();
  // The rotation is named by the spec's own display name, and links to its card.
  await expect(card.getByTestId('sim-rotation-card-link')).toHaveAttribute(
    'href',
    '/sim/specs#warrior-fury',
  );
  // The fixture's warrior-fury row is not validated, so the note is there; a validated
  // spec renders the card without one.
  await expect(card.getByTestId('sim-rotation-card-note')).toBeVisible();
});
```

`SAVED_ID` is the constant the file already declares for its prerendered fixture sim.

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && E2E_PORT=4399 npx playwright test tests/e2e/sim-saved.spec.ts --project=desktop
```

Expected: FAIL — no `sim-rotation-card`.

- [ ] **Step 3: Add the copy**

```ts
  // --- Design 5.1: the rotation card beside a result. ---
  rotationCard: 'Rotation',
  rotationCardBody: (name: string): string => `This run used the default rotation for ${name}.`,
```

- [ ] **Step 4: Write `RotationCard.svelte`**

```svelte
<!-- web/src/components/sim/RotationCard.svelte -->
<!-- Design 5.1. The settings bar's fidelity note says the same thing before a run; this
     says it beside the result, because a saved sim and a shared link have no settings bar
     and the number is exactly as good as the rotation that produced it.

     A validated spec renders the card with no note at all rather than a green "all is
     well" line: the absence of a caution is the message. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { specRow } from '../../lib/sim/spec-label';
  import { specPillClass, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';

  let { spec, fidelity }: { spec: string; fidelity: SpecFidelity | null } = $props();

  const name = $derived(specRow(spec)?.name ?? spec);
  const showNote = $derived(fidelity !== null && fidelity.state !== 'validated');
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-rotation-card"
>
  <h2 class="section-title text-[15px]">{simCopy.rotationCard}</h2>
  <p class="text-[13px]">
    {simCopy.rotationCardBody(name)}
    <a href={`/sim/specs#${spec}`} class="ml-1" data-testid="sim-rotation-card-link">
      {simCopy.rotationLink}
    </a>
  </p>
  {#if showNote && fidelity !== null}
    <p class="text-muted text-[13px]" data-testid="sim-rotation-card-note">
      <a href="/sim/specs" class={specPillClass(fidelity.state)}>{specStateLabel(fidelity.state)}</a>
      {specStateNote(fidelity.state)}
    </p>
  {/if}
</section>
```

- [ ] **Step 5: Mount it on `/sim`**

In `SimView.svelte`, immediately after the `DetailsCard`:

```svelte
        {#if store.result !== null && !comparing}
          <RotationCard spec={store.result.request.spec} fidelity={characterSpecRow} />
        {/if}
```

with `import RotationCard from './RotationCard.svelte';`.

- [ ] **Step 6: Mount it on `/sim/<id>`**

`SavedSim.svelte` has no spec list, so it fetches the one small, cacheable, credential-free GET the rest of the lane already shares. Add to its script:

```ts
  import { fetchSpecs } from '../../lib/sim/api';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';
  import RotationCard from './RotationCard.svelte';

  // One load, for this component's one unchanging result -- not an `$effect`, the same
  // reason the item-file load above is not one. A failed fetch leaves the card without a
  // fidelity note, which is the honest rendering of "we do not know yet".
  let fidelity = $state<SpecFidelity | null>(null);
  void (async () => {
    try {
      const rows = mergeSpecRows(await fetchSpecs());
      fidelity = rows.find((row) => row.spec === result.request.spec) ?? null;
    } catch {
      fidelity = null;
    }
  })();
```

and, after the `DetailsCard`:

```svelte
<RotationCard spec={result.request.spec} {fidelity} />
```

- [ ] **Step 7: Run the e2e**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && E2E_PORT=4399 npx playwright test tests/e2e/sim-saved.spec.ts tests/e2e/sim-results.spec.ts --project=desktop
```

Expected: PASS.

- [ ] **Step 8: Check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && npx eslint src/components/sim src/lib/sim tests/e2e && \
  npx prettier --check src/components/sim src/lib/sim tests/e2e && git add -A src/components/sim src/lib/sim tests/e2e && \
  git commit -m "feat(sim): the rotation card beside every result

Names the rotation the run used and links to its card, and carries the
fidelity note for a spec that is not validated. A saved sim has no
settings bar, so without this a shared link carried a number with no
statement of how good its rotation is.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 14: Report options — title, browser notification, open in a new tab

Design 5.4. The save form already pre-fills a title from `settingsLabel`; this promotes it to a field on the result itself, so it names the report before it is saved and is what the notification and the new tab carry.

**Files:**
- Create: `src/lib/sim/notify.ts`
- Modify: `src/lib/sim/store.svelte.ts`
- Modify: `src/components/sim/SimView.svelte`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/notify.test.ts`, `src/lib/sim/store.test.ts`, `tests/e2e/sim-saved.spec.ts`

**Interfaces:**
- Consumes: `settings.ts`'s `settingsLabel`.
- Produces:
  - `notify.ts`: `interface Notifier { permission: string; request(): Promise<string>; show(title: string, body: string): void }`, `browserNotifier(): Notifier | null`, `enableNotifications(notifier: Notifier | null): Promise<boolean>`, `notifyFinished(notifier: Notifier | null, enabled: boolean, title: string, body: string): boolean`
  - `store.svelte.ts`: `reportTitle` getter (falls back to `settingsLabel(settings)`), `setReportTitle(value: string)`
- Test ids added: `sim-report-title`, `sim-notify`, `sim-open-new-tab`.

- [ ] **Step 1: Write the failing notifier test**

Create `src/lib/sim/notify.test.ts`:

```ts
import { describe, expect, it, vi } from 'vitest';
import { enableNotifications, notifyFinished, type Notifier } from './notify';

function fake(permission: string, answer = 'granted'): Notifier & { shown: [string, string][] } {
  const shown: [string, string][] = [];
  return {
    permission,
    shown,
    request: vi.fn().mockResolvedValue(answer),
    show: (title, body) => shown.push([title, body]),
  };
}

describe('enableNotifications', () => {
  it('is true straight away when permission is already granted, and asks nothing', async () => {
    const notifier = fake('granted');
    expect(await enableNotifications(notifier)).toBe(true);
    expect(notifier.request).not.toHaveBeenCalled();
  });

  it('asks once when permission has not been decided', async () => {
    const notifier = fake('default', 'granted');
    expect(await enableNotifications(notifier)).toBe(true);
    expect(notifier.request).toHaveBeenCalledTimes(1);
  });

  it('is false when the answer is no, and never asks a denied browser again', async () => {
    const refused = fake('default', 'denied');
    expect(await enableNotifications(refused)).toBe(false);

    const denied = fake('denied');
    expect(await enableNotifications(denied)).toBe(false);
    expect(denied.request).not.toHaveBeenCalled();
  });

  it('is false where the browser has no notifications at all', async () => {
    expect(await enableNotifications(null)).toBe(false);
  });
});

describe('notifyFinished', () => {
  it('shows the title and the body once, and says it did', () => {
    const notifier = fake('granted');
    expect(notifyFinished(notifier, true, 'Fury, 3:00', '1,204 DPS')).toBe(true);
    expect(notifier.shown).toEqual([['Fury, 3:00', '1,204 DPS']]);
  });

  it('shows nothing when the player did not ask for it', () => {
    const notifier = fake('granted');
    expect(notifyFinished(notifier, false, 'a', 'b')).toBe(false);
    expect(notifier.shown).toEqual([]);
  });

  it('shows nothing without permission, however enabled the control is', () => {
    const notifier = fake('default');
    expect(notifyFinished(notifier, true, 'a', 'b')).toBe(false);
    expect(notifier.shown).toEqual([]);
  });

  it('never throws where there is no notifier', () => {
    expect(notifyFinished(null, true, 'a', 'b')).toBe(false);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/notify.test.ts
```

Expected: FAIL — `Failed to resolve import "./notify"`.

- [ ] **Step 3: Write `notify.ts`**

```ts
// web/src/lib/sim/notify.ts
// Design 5.4: "browser notification when a server run completes".
//
// The Notification API behind a two-method interface, so the decision logic is testable
// without a browser and so the store never touches `window` -- which is the rule
// store.svelte.ts's own header sets and which keeps it unit-testable.
//
// Permission is only ever asked for from the player's own click on the control. A page
// that asks on load is a page people permanently deny.

export interface Notifier {
  /** "default" | "granted" | "denied". */
  permission: string;
  request(): Promise<string>;
  show(title: string, body: string): void;
}

/** The real one, or null where the browser has no Notification API (Safari in a frame, and every SSR pass). */
export function browserNotifier(): Notifier | null {
  const ctor = (globalThis as { Notification?: typeof Notification }).Notification;
  if (ctor === undefined) return null;
  return {
    get permission() {
      return ctor.permission;
    },
    request: () => ctor.requestPermission(),
    show: (title, body) => {
      // A notification that throws must never take the results page down with it: a
      // browser can refuse to construct one even with permission (a private window, a
      // page that has lost its user gesture).
      try {
        // eslint-disable-next-line no-new
        new ctor(title, { body });
      } catch {
        /* no notification; the result is on screen regardless */
      }
    },
  };
}

/** True once notifications are usable. Asks at most once, and never a browser that said no. */
export async function enableNotifications(notifier: Notifier | null): Promise<boolean> {
  if (notifier === null) return false;
  if (notifier.permission === 'granted') return true;
  if (notifier.permission === 'denied') return false;
  return (await notifier.request()) === 'granted';
}

/** Shows one, if the player asked for them and the browser allows them. Returns whether it did. */
export function notifyFinished(
  notifier: Notifier | null,
  enabled: boolean,
  title: string,
  body: string,
): boolean {
  if (notifier === null || !enabled || notifier.permission !== 'granted') return false;
  notifier.show(title, body);
  return true;
}
```

- [ ] **Step 4: Write the failing store test**

Append to `src/lib/sim/store.test.ts`:

```ts
describe('the report title', () => {
  it('defaults to the settings clause and changes with the settings', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    expect(sim.reportTitle).toBe('Raid-buffed, 3:00, single target');
    sim.setSettings(withTargets(sim.settings, 4));
    expect(sim.reportTitle).toBe('Raid-buffed, 3:00, 4 targets');
  });

  it('keeps what the player typed, and an emptied field falls back rather than saving nothing', async () => {
    const sim = store();
    await sim.loadAddon(FURY);
    sim.setReportTitle('Pre-raid, no world buffs');
    expect(sim.reportTitle).toBe('Pre-raid, no world buffs');
    sim.setSettings(withTargets(sim.settings, 4));
    expect(sim.reportTitle).toBe('Pre-raid, no world buffs');
    sim.setReportTitle('   ');
    expect(sim.reportTitle).toBe('Raid-buffed, 3:00, 4 targets');
  });
});
```

with `import { withTargets } from './settings';` at the top of the file.

- [ ] **Step 5: Wire the store**

```ts
  // Empty means "the settings clause", which moves with the settings; anything the player
  // types wins until they clear it again. Blank-but-not-empty counts as empty: a title of
  // three spaces is not a title.
  let typedTitle = $state('');
```

```ts
    get reportTitle() {
      return typedTitle.trim() === '' ? settingsLabel(settings) : typedTitle;
    },
    setReportTitle(value: string): void {
      typedTitle = value;
    },
```

with `import { defaultSettings, settingsLabel, type SimSettings } from './settings';`. In `save()`, the default becomes the report title:

```ts
    async save(title?: string): Promise<string | null> {
      if (result === null) return null;
      const chosen = title ?? (typedTitle.trim() === '' ? settingsLabel(settings) : typedTitle);
      try {
        return await saveSim(result, init.apiBase, chosen);
      } catch {
        return null;
      }
    },
```

- [ ] **Step 6: Put the three controls on the page**

In `SimView.svelte`'s script:

```ts
  import { browserNotifier, enableNotifications, notifyFinished } from '../../lib/sim/notify';

  const notifier = untrack(() => browserNotifier());
  let notifyWanted = $state(false);
  // The id of the last result a notification was raised for, so a re-render never raises a
  // second one for the same run.
  let notifiedFor = $state('');

  async function toggleNotify(wanted: boolean): Promise<void> {
    notifyWanted = wanted && (await enableNotifications(notifier));
  }

  // Server runs only, per design 5.4: a browser run finishes on the tab you are looking at.
  $effect(() => {
    const finished = store.result;
    if (finished === null || finished.lane !== 'browser') {
      // `lane` is 'server' here; the key is the wall clock plus the figure, which no two
      // runs of one session share.
      const key = `${finished?.duration_ms ?? 0}-${finished?.iterations_run ?? 0}`;
      if (finished !== null && key !== notifiedFor) {
        notifiedFor = key;
        notifyFinished(
          notifier,
          notifyWanted,
          store.reportTitle,
          simCopy.notifyBody(Math.round(finished.dps.mean).toLocaleString('en-US')),
        );
      }
    }
  });
```

and, in the template, immediately above the save form's `<div data-testid="sim-save">`:

```svelte
        {#if store.result !== null}
          <div class="mx-[18px] flex flex-wrap items-end gap-3 md:mx-0">
            <label class="flex min-w-0 flex-1 flex-col gap-1 md:max-w-[420px]">
              <span class="label text-muted">{simCopy.reportTitleLabel}</span>
              <input
                type="text"
                class="border-line-warm rounded-control bg-raised text-text h-11 w-full border px-3 text-[14px]"
                value={store.reportTitle}
                onchange={(event) => store.setReportTitle(event.currentTarget.value)}
                data-testid="sim-report-title"
              />
            </label>
            {#if notifier !== null}
              <label class="flex min-h-11 items-center gap-2 text-[13px]">
                <input
                  type="checkbox"
                  class="accent-gold h-5 w-5"
                  checked={notifyWanted}
                  onchange={(event) => void toggleNotify(event.currentTarget.checked)}
                  data-testid="sim-notify"
                />
                <span class="text-muted">{simCopy.notifyLabel}</span>
              </label>
            {/if}
          </div>
        {/if}
```

and, inside the `{#if savedUrl !== null}` branch of the save form, after the copy button:

```svelte
            <a
              class="border-line-warm rounded-control text-nav label inline-flex min-h-11 items-center border px-4"
              href={savedUrl}
              target="_blank"
              rel="noopener"
              data-testid="sim-open-new-tab"
            >
              {simCopy.openInNewTab}
            </a>
```

Change `openSaveForm()` to seed from the report title rather than recomputing the label:

```ts
  function openSaveForm(): void {
    saveTitle = store.reportTitle;
    saveFailed = false;
    savedUrl = null;
    saveOpen = true;
  }
```

- [ ] **Step 7: Add the copy**

```ts
  // --- Design 5.4: report options. ---
  reportTitleLabel: 'Name this report',
  notifyLabel: 'Tell me when a server run finishes',
  notifyBody: (dps: string): string => `${dps} DPS. Your sim has finished.`,
  openInNewTab: 'Open in a new tab',
```

- [ ] **Step 8: Extend the e2e**

Append to `tests/e2e/sim-run.spec.ts`:

```ts
test('a finished run can be named, and the saved link opens in a new tab', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await page.getByTestId('sim-precision').selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });

  const title = page.getByTestId('sim-report-title');
  await expect(title).toHaveValue('Raid-buffed, 3:00, single target');
  await title.fill('Pre-raid, no world buffs');
  await page.getByTestId('sim-save-open').click();
  await expect(page.getByTestId('sim-save-title')).toHaveValue('Pre-raid, no world buffs');
});
```

- [ ] **Step 9: Run both suites, check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-run.spec.ts --project=desktop && \
  npx astro check && npx eslint src/lib/sim src/components/sim tests/e2e && \
  npx prettier --check src/lib/sim src/components/sim tests/e2e && \
  git add -A src/lib/sim src/components/sim tests/e2e && \
  git commit -m "feat(sim): report title, the finish notification and open in a new tab

Design 5.4. The title names the report before it is saved and is what
the save form, the notification and the new tab all carry. Permission
is asked for only from the player's own click on the control, and only
a server run raises a notification -- a browser run finishes on the tab
you are already looking at.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 15: The request drawer

Design 8, contract 9's `sim-request-drawer`. The exact JSON the run will send, editable, validated by the engine's own `Validate` with the errors inline, and two ways out: apply it to the page, or run it verbatim.

Two buttons rather than one, because they are genuinely different things. **Apply to the page** rebuilds the page's own state from the request — settings, precision and the character — and is what a shared request from another player is for. **Run this request** sends the edited JSON exactly as typed, which is the escape hatch the design names: "any field the panel does not expose is reachable here". Apply cannot carry a field the page has no control for; Run can.

> **Temporary, and lifted by Task 17.** Apply rebuilds the character through the FS1 route (`encodeFS1` over the request's class, race, talents and gear), which is how `SimView`'s existing "Run this yourself" already works. `SimCharacter.gear` is a map of item ids today, so a per-slot enchant or suffix in a pasted request is dropped by Apply and kept by Run; the drawer's copy says so. Contract 10.5 puts per-slot enchant and suffix on `SimCharacter` and lets an FS1 gear entry carry them, so **Task 17 upgrades this call to `encodeFS1V2`, makes Apply lossless, and replaces the note**. This task ships the honest version of the sentence rather than a promise.

**Files:**
- Create: `src/lib/sim/request-json.ts`
- Create: `src/components/sim/RequestDrawer.svelte`
- Modify: `src/lib/sim/store.svelte.ts`
- Modify: `src/components/sim/SimView.svelte`
- Modify: `src/lib/sim/copy.ts`
- Test: `src/lib/sim/request-json.test.ts`, `tests/e2e/sim-request.spec.ts` (new)

**Interfaces:**
- Consumes: `engine.ts`'s `RequestValidation` and `worker.ts`'s `SimPool.validate` (Task 5), `precision.ts`'s `precisionOf` (Task 6), `settings.ts` (Task 3), `buffs.ts` (Task 8).
- Produces:
  - `request-json.ts`: `formatRequest(request: SimRequest): string`, `type ParsedRequest = { ok: true; request: SimRequest } | { ok: false; message: string }`, `parseRequest(text: string): ParsedRequest`, `settingsFromRequest(request: SimRequest): SimSettings`, `MAX_REQUEST_CHARS`
  - `store.svelte.ts`: `buildRequest(): SimRequest | null`, `validateRequest(json: string): Promise<RequestValidation>`, `applyRequest(request: SimRequest): Promise<void>`, `runRequest(request: SimRequest): Promise<void>`
- Test ids added: `sim-request-drawer`, `sim-request-json`, `sim-request-errors`, `sim-request-error-<field>`, `sim-request-apply`, `sim-request-run`, `sim-request-share`, `sim-request-buffs`.

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/request-json.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import envelope from '../../fixtures/sim/envelope-v2.json';
import { MAX_REQUEST_CHARS, formatRequest, parseRequest, settingsFromRequest } from './request-json';
import { defaultSettings } from './settings';
import type { SimRequest } from './types';

const request = (envelope as unknown as { request: SimRequest }).request;

describe('formatRequest', () => {
  it('is indented JSON a person can edit, ending in a newline', () => {
    const text = formatRequest(request);
    expect(text.startsWith('{\n  "engine_version"')).toBe(true);
    expect(text.endsWith('\n')).toBe(true);
  });

  it('round-trips through parseRequest unchanged', () => {
    const parsed = parseRequest(formatRequest(request));
    expect(parsed.ok).toBe(true);
    if (parsed.ok) expect(parsed.request).toEqual(request);
  });
});

describe('parseRequest', () => {
  it('names the syntax error rather than saying “invalid”', () => {
    const parsed = parseRequest('{ "spec": ');
    expect(parsed.ok).toBe(false);
    if (!parsed.ok) expect(parsed.message.length).toBeGreaterThan(0);
  });

  it('refuses anything that is not a JSON object', () => {
    for (const text of ['[]', '"a string"', '42', 'null']) {
      expect(parseRequest(text).ok, text).toBe(false);
    }
  });

  it('refuses a request past the size bound before parsing it', () => {
    const parsed = parseRequest('{'.padEnd(MAX_REQUEST_CHARS + 1, ' '));
    expect(parsed.ok).toBe(false);
  });
});

describe('settingsFromRequest', () => {
  it('takes the encounter verbatim', () => {
    expect(settingsFromRequest(request).encounter).toEqual(request.encounter);
  });

  it('takes the buffs, consumables and cooldowns, and calls the preset custom', () => {
    const settings = settingsFromRequest(request);
    expect(settings.buffs).toEqual(request.character.buffs);
    expect(settings.consumables).toEqual(request.character.consumes);
    expect(settings.cooldowns).toEqual(request.character.cooldowns);
    // A pasted list is nobody's preset, and calling it "raid-buffed" would let the next
    // preset change silently discard it.
    expect(settings.preset).toBe('custom');
  });

  it('fills an absent encounter field from the defaults rather than leaving it undefined', () => {
    const thin = { ...request, encounter: { ...request.encounter, target_level: undefined } };
    expect(settingsFromRequest(thin as SimRequest).encounter.target_level).toBe(
      defaultSettings().encounter.target_level,
    );
  });

  it('copies rather than sharing, so editing the page never edits the pasted request', () => {
    const settings = settingsFromRequest(request);
    expect(settings.buffs).not.toBe(request.character.buffs);
    expect(settings.encounter).not.toBe(request.encounter);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/request-json.test.ts
```

Expected: FAIL — `Failed to resolve import "./request-json"`.

- [ ] **Step 3: Write `request-json.ts`**

```ts
// web/src/lib/sim/request-json.ts
// Design 8: "Advanced is a JSON editor, not a script language". SimulationCraft input is
// Raidbots' escape hatch; ours is the request envelope itself.
//
// Nothing here validates: `api.SimRequest.Validate` runs in the wasm and the drawer shows
// whatever it says. This module only turns a request into text, text back into a request,
// and a request back into the page's own settings.
import { defaultSettings, type SimSettings } from './settings';
import { simCopy } from './copy';
import type { SimRequest } from './types';

/**
 * A generous bound on a hand-edited request. A real one is a couple of kilobytes; a bulk
 * request with four hundred candidates is under sixty. This exists so a pasted megabyte
 * is refused before `JSON.parse` is asked to do any work on it -- the same reason
 * `fs1.ts` and `url.ts` each carry one.
 */
export const MAX_REQUEST_CHARS = 262_144;

/** Indented, and newline-terminated so a textarea's last line is editable. */
export function formatRequest(request: SimRequest): string {
  return `${JSON.stringify(request, null, 2)}\n`;
}

export type ParsedRequest = { ok: true; request: SimRequest } | { ok: false; message: string };

export function parseRequest(text: string): ParsedRequest {
  if (text.length > MAX_REQUEST_CHARS) return { ok: false, message: simCopy.requestTooLong };
  let parsed: unknown;
  try {
    parsed = JSON.parse(text);
  } catch (error) {
    // The parser's own message names the line and column, which is the only thing that
    // says where the typo is. Ours says what kind of thing went wrong.
    return { ok: false, message: `${simCopy.requestNotJson} ${error instanceof Error ? error.message : ''}`.trim() };
  }
  if (typeof parsed !== 'object' || parsed === null || Array.isArray(parsed)) {
    return { ok: false, message: simCopy.requestNotObject };
  }
  return { ok: true, request: parsed as SimRequest };
}

/**
 * The page's settings, rebuilt from a request. Every encounter field the request does not
 * carry falls back to the default rather than staying undefined: the controls are
 * `<select>`s over closed vocabularies and an undefined value renders as a blank option.
 *
 * The preset is always "custom". A pasted buff list is nobody's preset, and filing it
 * under one would mean the next preset change silently discarded it.
 */
export function settingsFromRequest(request: SimRequest): SimSettings {
  const base = defaultSettings();
  return {
    encounter: { ...base.encounter, ...request.encounter },
    preset: 'custom',
    buffs: [...(request.character.buffs ?? [])],
    consumables: [...(request.character.consumes ?? [])],
    cooldowns: (request.character.cooldowns ?? []).map((row) => ({ ...row, at_sec: [...row.at_sec] })),
  };
}
```

- [ ] **Step 4: Add the copy**

```ts
  // --- Design 8: the request drawer. ---
  requestDrawer: 'Request',
  requestNote:
    'The exact JSON this run sends. Edit it and run it as written: anything the panels above do not offer is reachable here.',
  requestApply: 'Apply to the page',
  // Task 17 replaces this sentence with the lossless one once contract 10.5's per-slot
  // enchant and suffix are on SimCharacter. Until then it states what actually happens.
  requestApplyNote:
    'Rebuilds the settings, the precision and the character from this request. Per-slot enchants and suffixes are not part of the page’s character yet, so Apply drops them; Run keeps them.',
  requestRun: 'Run this request',
  requestShare: 'Copy a link to this request',
  requestValid: 'The engine accepts this request.',
  requestTooLong: 'That request is too long to read.',
  requestNotJson: 'That is not JSON.',
  requestNotObject: 'A request is a JSON object.',
  requestShareTooLong:
    'This request is too long for a link. Save it and share the saved link instead.',
```

- [ ] **Step 5: Add the four store methods**

In `store.svelte.ts`:

```ts
    /**
     * The request the page would send right now, or null while there is no character or
     * no talent index. One function, so the drawer, the share link and the run can never
     * disagree about what "this request" is -- the same reason `buildSimRequest` exists.
     */
    buildRequest(): SimRequest | null {
      if (character === null) return null;
      const index = talents === null ? null : indexTalents(talents);
      if (index === null) return null;
      const plan = precisionPlan(precisionId, 'browser');
      return buildSimRequest({
        spec: character.spec,
        source: character.source,
        character: toCharacterSpec(character, index, settings.buffs, settings.consumables, settings.cooldowns),
        encounter: settings.encounter,
        iterations: plan.iterations,
        targetError: plan.targetError,
        stepIterations: plan.step,
      });
    },

    /** `api.SimRequest.Validate`, inside the wasm. Never a rule written here. */
    validateRequest(json: string): Promise<RequestValidation> {
      return poolOnce().validate(json);
    },

    /**
     * A pasted request as page state: settings and precision exactly, and the character
     * through the same FS1 route "Run this yourself" already uses. The gear a
     * `SimCharacter` carries is a map of item ids, so an enchant or suffix on a request's
     * gear slot is dropped here -- `runRequest` below is what keeps it.
     */
    async applyRequest(request: SimRequest): Promise<void> {
      settings = settingsFromRequest(request);
      precisionId = precisionOf(request);
      await adopt(
        fromPlannerCode(
          encodeFS1({
            dataBuild: init.treeVersion,
            classSlug: request.character.class,
            raceSlug: request.character.race,
            treeRanks: ranksFromTalentsString(request.character.talents),
            gear: gearFromSlots(request.character.gear),
          }),
          ctx,
        ),
      );
    },

    /** The edited request, run exactly as written. The escape hatch of design 8. */
    async runRequest(request: SimRequest): Promise<void> {
      message = null;
      detail = '';
      stopRequested = false;
      phase = 'loading-engine';
      iterationsTotal = request.iterations;
      iterationsDone = 0;
      relative = 0;
      phase = 'running';
      handle = runSim(
        poolOnce(),
        {
          spec: request.spec,
          source: request.source,
          character: request.character,
          encounter: request.encounter,
          iterations: request.iterations,
          randomSeed: request.random_seed,
          targetError: request.target_error,
          stepIterations: STEP_ITERATIONS,
        },
        (update) => {
          if (stopRequested) return;
          estimate = update.estimate;
          iterationsDone = update.iterationsDone;
          iterationsTotal = update.iterationsTotal;
          relative = update.relativeError;
        },
      );
      try {
        const finished = await handle.result;
        if (stopRequested) {
          message = simCopy.stopped;
          restorePreviousResult();
          phase = result !== null ? 'done' : 'idle';
          return;
        }
        result = finished;
        phase = 'done';
      } catch (error) {
        const failure = error instanceof SimRunError ? error : null;
        message = failure?.cancelled === true ? simCopy.stopped : (failure?.message ?? simCopy.failed);
        detail = failure?.detail ?? '';
        phase = 'error';
      } finally {
        handle = null;
      }
    },
```

New imports: `encodeFS1` from `../planner/fs1`, `gearFromSlots`, `ranksFromTalentsString` from `./character`, `precisionOf`, `STEP_ITERATIONS` from `./precision`, `settingsFromRequest` from `./request-json`, `type RequestValidation` from `./engine`, `type SimRequest` from `./types`.

- [ ] **Step 6: Write `RequestDrawer.svelte`**

```svelte
<!-- web/src/components/sim/RequestDrawer.svelte -->
<!-- Design 8. The exact JSON the run will send, editable, with the engine's own Validate
     behind it and its errors beside the fields they name.

     Two exits, because they are two different things: Apply rebuilds the page from the
     request and can only carry what the page has controls for; Run sends the text exactly
     as typed, which is what makes this the escape hatch rather than a second settings bar. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import type { RequestValidation } from '../../lib/sim/engine';
  import { formatRequest, parseRequest } from '../../lib/sim/request-json';
  import type { SimRequest } from '../../lib/sim/types';

  let {
    request,
    disabled,
    onvalidate,
    onapply,
    onrun,
    onshare,
  }: {
    /** The request the page would send now. Null while there is no character. */
    request: SimRequest | null;
    disabled: boolean;
    onvalidate: (json: string) => Promise<RequestValidation>;
    onapply: (request: SimRequest) => void;
    onrun: (request: SimRequest) => void;
    /** Returns the share URL, or null when the request is past the URL budget. */
    onshare: (request: SimRequest) => string | null;
  } = $props();

  // Seeded from the page's own request and then owned by the player: re-seeding it on
  // every settings change would throw away what they were typing.
  let text = $state('');
  let seeded = $state(false);
  let parseError = $state('');
  let validation = $state<RequestValidation | null>(null);
  let shared = $state('');
  let shareError = $state('');

  $effect(() => {
    if (!seeded && request !== null) {
      text = formatRequest(request);
      seeded = true;
    }
  });

  function reseed(): void {
    if (request !== null) text = formatRequest(request);
    parseError = '';
    validation = null;
  }

  async function check(): Promise<SimRequest | null> {
    shared = '';
    shareError = '';
    const parsed = parseRequest(text);
    if (!parsed.ok) {
      parseError = parsed.message;
      validation = null;
      return null;
    }
    parseError = '';
    validation = await onvalidate(text);
    return validation.ok ? parsed.request : null;
  }

  async function apply(): Promise<void> {
    const ok = await check();
    if (ok !== null) onapply(ok);
  }

  async function run(): Promise<void> {
    const ok = await check();
    if (ok !== null) onrun(ok);
  }

  async function share(): Promise<void> {
    const ok = await check();
    if (ok === null) return;
    const url = onshare(ok);
    if (url === null) {
      shareError = simCopy.requestShareTooLong;
      return;
    }
    shared = url;
    try {
      await navigator.clipboard.writeText(url);
    } catch {
      /* the field below holds it; a browser that refuses the clipboard is not an error */
    }
  }

  const button =
    'border-line-warm rounded-control text-nav label min-h-11 border px-4 disabled:opacity-50 md:min-h-9';
</script>

<details class="border-line bg-raised rounded-panel mx-[18px] border md:mx-0" data-testid="sim-request-drawer">
  <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-4 md:min-h-9">
    {simCopy.requestDrawer}
  </summary>
  <div class="flex flex-col gap-3 p-4 pt-0">
    <p class="text-muted text-[12px]">{simCopy.requestNote}</p>
    <textarea
      class="border-line-warm rounded-control bg-card-top text-text h-72 w-full border p-3 font-mono text-[12px]"
      spellcheck="false"
      bind:value={text}
      data-testid="sim-request-json"
    ></textarea>

    {#if parseError !== ''}
      <p role="alert" class="text-strong text-[13px]" data-testid="sim-request-errors">{parseError}</p>
    {:else if validation !== null && !validation.ok}
      <ul class="flex flex-col gap-1 text-[13px]" role="alert" data-testid="sim-request-errors">
        {#each validation.errors as row, index (`${row.field}-${index}`)}
          <li data-testid={`sim-request-error-${row.field}`}>
            <span class="text-strong font-mono">{row.field}</span>
            <span class="text-muted ml-2">{row.message}</span>
          </li>
        {/each}
      </ul>
    {:else if validation !== null}
      <p class="text-muted text-[13px]" data-testid="sim-request-valid">{simCopy.requestValid}</p>
    {/if}

    <div class="flex flex-wrap items-center gap-3">
      <button type="button" class={button} {disabled} onclick={() => void apply()} data-testid="sim-request-apply">
        {simCopy.requestApply}
      </button>
      <button type="button" class={button} {disabled} onclick={() => void run()} data-testid="sim-request-run">
        {simCopy.requestRun}
      </button>
      <button type="button" class={button} onclick={() => void share()} data-testid="sim-request-share">
        {simCopy.requestShare}
      </button>
      <button type="button" class={button} onclick={reseed} data-testid="sim-request-reset">
        {simCopy.cancel}
      </button>
    </div>
    <p class="text-muted text-[12px]">{simCopy.requestApplyNote}</p>

    {#if shareError !== ''}
      <p role="alert" class="text-strong text-[13px]" data-testid="sim-request-share-error">{shareError}</p>
    {:else if shared !== ''}
      <input
        type="text"
        readonly
        value={shared}
        class="border-line-warm rounded-control bg-raised text-text h-11 w-full border px-3 text-[13px]"
        data-testid="sim-request-share-link"
        onclick={(event) => event.currentTarget.select()}
      />
    {/if}

    <!-- The buff list as plain text, so a test (and a person) can read what is actually
         going to the engine without parsing the textarea. -->
    <p class="text-muted font-mono text-[11px]" data-testid="sim-request-buffs">
      {(request?.character.buffs ?? []).join(' ')}
    </p>
  </div>
</details>
```

- [ ] **Step 7: Mount it**

In `SimView.svelte`, after the settings bar and the buff panel:

```svelte
        <RequestDrawer
          request={store.buildRequest()}
          disabled={store.phase === 'running' || store.serverRunning}
          onvalidate={(json) => store.validateRequest(json)}
          onapply={(request) => void store.applyRequest(request)}
          onrun={(request) => void store.runRequest(request)}
          onshare={(request) => shareUrlFor(request)}
        />
```

`shareUrlFor` lands in Task 16; until then define it beside the other handlers as:

```ts
  function shareUrlFor(_request: SimRequest): string | null {
    return null;
  }
```

with a `// Task 16 fills this in.` comment, so this task's drawer is complete and its share button honestly reports "too long for a link" until the encoder exists.

- [ ] **Step 8: Write the e2e**

Create `tests/e2e/sim-request.spec.ts`:

```ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('the drawer shows the exact request, validates an edit and runs it', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  const drawer = page.getByTestId('sim-request-drawer');
  await drawer.locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const json = await editor.inputValue();
  expect(JSON.parse(json)).toMatchObject({ spec: 'warrior-fury', iterations: 3000 });

  // The engine's own Validate, with the field it refuses named beside the message.
  await editor.fill(JSON.stringify({ ...JSON.parse(json), spec: '' }, null, 2));
  await page.getByTestId('sim-request-run').click();
  await expect(page.getByTestId('sim-request-error-spec')).toBeVisible();

  // Syntax errors are caught before the engine is asked anything.
  await editor.fill('{ "spec": ');
  await page.getByTestId('sim-request-run').click();
  await expect(page.getByTestId('sim-request-errors')).toContainText('not JSON');

  // A legal edit runs exactly as written: 500 iterations, whatever the precision select says.
  await editor.fill(JSON.stringify({ ...JSON.parse(json), iterations: 500 }, null, 2));
  await page.getByTestId('sim-request-run').click();
  await expect(page.getByTestId('sim-request-valid')).toBeVisible();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });
  await expect(page.getByTestId('sim-details-iterations')).toHaveText('500');
});

test('a pasted request loads the page state', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const request = JSON.parse(await editor.inputValue());
  request.encounter.duration_sec = 600;
  request.encounter.targets = 5;
  request.character.buffs = ['thorns'];
  request.iterations = 500;
  await editor.fill(JSON.stringify(request, null, 2));
  await page.getByTestId('sim-request-apply').click();

  await expect(page.getByTestId('sim-duration')).toHaveValue('600');
  await expect(page.getByTestId('sim-targets')).toHaveValue('5');
  await expect(page.getByTestId('sim-preset')).toHaveValue('custom');
  await expect(page.getByTestId('sim-precision')).toHaveValue('fast');
  await expect(page.getByTestId('sim-request-buffs')).toHaveText('thorns');
});
```

- [ ] **Step 9: Run both suites, check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-request.spec.ts --project=desktop && \
  npx astro check && npx eslint src/lib/sim src/components/sim tests/e2e && \
  npx prettier --check src/lib/sim src/components/sim tests/e2e && \
  git add -A src/lib/sim src/components/sim tests/e2e && \
  git commit -m "feat(sim): the request drawer, validated by the engine's own Validate

Design 8. The exact JSON the run sends, editable, with simValidate's
errors beside the fields they name. Apply rebuilds the page from a
pasted request; Run sends the text exactly as typed, which is what
makes any field the panels do not expose reachable.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 16: A share URL of an edited request

Design 8's last sentence: "a share URL of an edited request is a full reproduction". Contract 9: "Share URLs carry the whole request as today; a bulk request above the URL budget is shared by its saved id only."

**Files:**
- Modify: `src/lib/sim/url.ts`
- Modify: `src/lib/sim/store.svelte.ts`
- Modify: `src/components/sim/SimView.svelte`
- Test: `src/lib/sim/url.test.ts`, `tests/e2e/sim-request.spec.ts`

**Interfaces:**
- Consumes: `request-json.ts` (Task 15).
- Produces:
  - `url.ts`: `SimState` (+ `req: string`), `MAX_REQUEST_PARAM`, `encodeRequestParam(request: SimRequest): string | null`, `decodeRequestParam(value: string): SimRequest | null`
  - `store.svelte.ts`: `SimStoreInit` (+ `request?: SimRequest`)

- [ ] **Step 1: Write the failing test**

Append to `src/lib/sim/url.test.ts`:

```ts
describe('a request in the URL', () => {
  const request = (envelope as unknown as { request: SimRequest }).request;

  it('round-trips a request through the query string', () => {
    const encoded = encodeRequestParam(request);
    expect(encoded).not.toBeNull();
    expect(decodeRequestParam(encoded!)).toEqual(request);
  });

  it('is URL-safe: no +, / or = to be mangled by a chat client', () => {
    expect(encodeRequestParam(request)!).toMatch(/^[A-Za-z0-9_-]+$/);
  });

  it('refuses a request past the budget rather than writing a link that will be cut', () => {
    const huge = {
      ...request,
      character: { ...request.character, buffs: Array.from({ length: 20_000 }, (_, i) => `b${i}`) },
    };
    expect(encodeRequestParam(huge)).toBeNull();
  });

  it('answers null for a value that is not a request, however it is malformed', () => {
    expect(decodeRequestParam('not-base64!!')).toBeNull();
    expect(decodeRequestParam(btoa('[1,2,3]').replaceAll('=', ''))).toBeNull();
    expect(decodeRequestParam('')).toBeNull();
  });

  it('parses and writes the req parameter beside the others', () => {
    const encoded = encodeRequestParam(request)!;
    const state = parseSimState(`?req=${encoded}`);
    expect(state.req).toBe(encoded);
    expect(simSearch({ ...defaultSimState(), req: encoded })).toBe(`?req=${encoded}`);
  });

  it('drops a req parameter past the budget instead of handing a truncated one to the decoder', () => {
    expect(parseSimState(`?req=${'a'.repeat(MAX_REQUEST_PARAM + 1)}`).req).toBe('');
  });
});
```

with, at the top of the file, `import envelope from '../../fixtures/sim/envelope-v2.json';`, `import type { SimRequest } from './types';` and the four new names added to the existing `./url` import.

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/url.test.ts
```

Expected: FAIL — `encodeRequestParam is not a function`.

- [ ] **Step 3: Extend `url.ts`**

```ts
/**
 * A base64url request is about a third larger than its JSON, and a real single-run request
 * is two to three kilobytes, so this leaves a comfortable margin under the ~8 KB every
 * browser and proxy handles. A request past it is shared by its saved id instead
 * (contract 9), which the drawer says in so many words.
 */
export const MAX_REQUEST_PARAM = 8192;

/** UTF-8 bytes to base64url, no padding: a chat client must not mangle a share link. */
function toBase64Url(text: string): string {
  const bytes = new TextEncoder().encode(text);
  let binary = '';
  for (const byte of bytes) binary += String.fromCharCode(byte);
  return btoa(binary).replaceAll('+', '-').replaceAll('/', '_').replaceAll('=', '');
}

function fromBase64Url(value: string): string {
  const padded = value.replaceAll('-', '+').replaceAll('_', '/');
  const binary = atob(padded.padEnd(Math.ceil(padded.length / 4) * 4, '='));
  const bytes = Uint8Array.from(binary, (character) => character.charCodeAt(0));
  return new TextDecoder().decode(bytes);
}

/** The request as a query value, or null when it is past the budget. */
export function encodeRequestParam(request: SimRequest): string | null {
  const encoded = toBase64Url(JSON.stringify(request));
  return encoded.length > MAX_REQUEST_PARAM ? null : encoded;
}

/**
 * A query value back into a request, or null. The query is attacker-controlled and this
 * value reaches the engine, so anything that is not a JSON object is refused here and the
 * engine's own Validate refuses the rest.
 */
export function decodeRequestParam(value: string): SimRequest | null {
  if (value === '' || value.length > MAX_REQUEST_PARAM) return null;
  try {
    const parsed: unknown = JSON.parse(fromBase64Url(value));
    if (typeof parsed !== 'object' || parsed === null || Array.isArray(parsed)) return null;
    return parsed as SimRequest;
  } catch {
    return null;
  }
}
```

`SimState`, `defaultSimState`, `parseSimState` and `simSearch` each gain the field:

```ts
export interface SimState {
  source: SourceKind | '';
  ref: string;
  code: string;
  /** A whole request, base64url-encoded: the drawer's share link (design 8). */
  req: string;
  mode: SimMode;
  fight: string;
}
```

```ts
  return { source: '', ref: '', code: '', req: '', mode: 'sim', fight: '' };
```

```ts
    req: bounded(params.get('req'), MAX_REQUEST_PARAM),
```

```ts
  if (state.req !== '') params.set('req', state.req);
```

with `import type { SimRequest, SourceKind } from './types';`.

- [ ] **Step 4: Bootstrap from it**

`SimStoreInit` gains one field, documented beside `code`:

```ts
  /**
   * A whole request from a share link (`/sim?req=…`, design 8). It wins over `code` and
   * over `source`/`ref`: it is the most specific thing a link can carry, and it carries
   * the settings and the precision as well as the character.
   */
  request?: SimRequest;
```

and the `ready` bootstrap gains a first branch:

```ts
  const ready: Promise<void> =
    init.request !== undefined
      ? applyRequestOnce(init.request)
      : init.code !== undefined && init.code !== ''
        ? adopt(fromPlannerCode(init.code, ctx))
        : (() => {
            const load =
              init.source !== undefined && init.ref !== undefined && init.ref !== ''
                ? bootstrapSource(init.source, init.ref, ctx)
                : null;
            return load === null ? Promise.resolve() : adopt(load);
          })();
```

`applyRequestOnce` is the body of the returned `applyRequest`, hoisted to a function declaration above `ready` so both can call it — the returned method becomes `applyRequest: (request: SimRequest) => applyRequestOnce(request),`.

In `SimView.svelte`'s `bootstrap`, read and decode it, and hand it across:

```ts
    const { source, ref, code, req, mode } = parseSimState(search);
    …
    return { treeVersion, source, ref, code, request: decodeRequestParam(req), mode, view };
```

```ts
    createSimStore({
      treeVersion: bootstrap.treeVersion,
      request: bootstrap.request ?? undefined,
      source: bootstrap.mode === 'compare' ? undefined : bootstrap.source,
      ref: bootstrap.mode === 'compare' ? undefined : bootstrap.ref,
      code: bootstrap.code,
    }),
```

and `shareUrlFor` becomes real:

```ts
  /** Design 8: a share URL of an edited request is a full reproduction. Null past the budget. */
  function shareUrlFor(request: SimRequest): string | null {
    const encoded = encodeRequestParam(request);
    if (encoded === null) return null;
    return `${window.location.origin}/sim${simSearch(withSimState(defaultSimState(), { req: encoded }))}`;
  }
```

with `encodeRequestParam` and `decodeRequestParam` added to the `./url` import.

- [ ] **Step 5: Extend the e2e**

Append to `tests/e2e/sim-request.spec.ts`:

```ts
test('a shared request link reproduces the whole page state', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-style').selectOption('cleave-5');
  await page.getByTestId('sim-duration').selectOption('600');
  await page.getByTestId('sim-request-drawer').locator('summary').click();
  await page.getByTestId('sim-request-share').click();

  const url = await page.getByTestId('sim-request-share-link').inputValue();
  expect(url).toContain('/sim?req=');

  await page.goto(url);
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-targets')).toHaveValue('5');
  await expect(page.getByTestId('sim-duration')).toHaveValue('600');
});
```

- [ ] **Step 6: Run both suites, check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-request.spec.ts --project=desktop && \
  npx astro check && npx eslint src/lib/sim src/components/sim tests/e2e && \
  npx prettier --check src/lib/sim src/components/sim tests/e2e && \
  git add -A src/lib/sim src/components/sim tests/e2e && \
  git commit -m "feat(sim): share an edited request as a link

Design 8: the request is base64url in ?req=, and opening the link
rebuilds the settings, the precision and the character from it. A
request past the 8 KB budget answers null and the drawer says to save
it and share the saved link instead, which is contract 9's rule.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 17: FS1 version 2, with per-slot enchants, suffixes and professions

Contract 7 as corrected by **10.5**. Everything before the first `|` is version 1 unchanged, so every existing decoder keeps working; the sections after it are fixed in order, and unknown ones are ignored and reported. 10.5 adds three things to what section 7 said:

- a `<gear>` entry and a `sets=` gear entry may carry `item_id[:enchant[:suffix]]`, exactly as a bag entry does — a version-1 decoder reading a bare id is unaffected, because a bare id is still a legal entry;
- a `professions=<slug>,<slug>` section follows `loadouts=`;
- `SimCharacter` gains per-slot enchant and suffix, which is what makes the request drawer's "Apply" lossless (Task 15's note comes off in this task).

**Files:**
- Modify: `src/lib/planner/fs1.ts`
- Modify: `src/lib/sim/character.ts`
- Modify: `src/lib/sim/url.ts` (`MAX_CODE`)
- Modify: `src/lib/sim/store.svelte.ts` (`applyRequest` becomes lossless)
- Modify: `src/lib/sim/copy.ts` (`requestApplyNote`)
- Test: `src/lib/planner/fs1.test.ts`, `src/lib/sim/character.test.ts`, `tests/e2e/sim-request.spec.ts`

**Interfaces:**
- Consumes: `planner/types.ts`'s `SLOTS`, `Gear`, `Slot`; `sim/types.ts`'s `GearSlot`.
- Produces:
  - `fs1.ts`: `MAX_CODE_LENGTH = 16_384` (exported now), `interface FS1Item { itemId: number; enchant?: number; suffix?: number }`, `interface FS1GearSlot { slot: Slot; itemId: number; enchant?: number; suffix?: number }`, `interface FS1Set { name: string; gear: FS1GearSlot[] }`, `interface FS1Loadout { name: string; treeRanks: number[][] }`, `FS1Build` (+ `gearSlots?: FS1GearSlot[]`, `bags: FS1Item[]`, `bank: FS1Item[]`, `sets: FS1Set[]`, `loadouts: FS1Loadout[]`, `professions: string[]`, `ignored: string[]`), `gearSlotsFrom(gear: Gear): FS1GearSlot[]`, `encodeFS1` (unchanged output), `encodeFS1V2(build: FS1Build): string`
  - `character.ts`: `SimCharacter` (+ `gear_slots: GearSlot[]`, `professions: string[]`, `bags: FS1Item[]`, `bank: FS1Item[]`, `sets: FS1Set[]`, `loadouts: FS1Loadout[]`) — **this is the block part B reads for its candidate lists**; `toCharacterSpec` sends `gear_slots` and `professions` when it has them.

- [ ] **Step 1: Write the failing test**

Append to `src/lib/planner/fs1.test.ts`:

```ts
describe('version 2 sections', () => {
  const V1 = 'FS1:1.15.9:warrior:orc:0/5530515/0:head=12640,main_hand=11726';

  it('reads a version 1 code unchanged, with the new fields empty', () => {
    const decoded = decodeFS1(V1);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.gear).toEqual({ head: 12640, main_hand: 11726 });
    expect(decoded.build.gearSlots).toEqual([
      { slot: 'head', itemId: 12640 },
      { slot: 'main_hand', itemId: 11726 },
    ]);
    expect(decoded.build.bags).toEqual([]);
    expect(decoded.build.bank).toEqual([]);
    expect(decoded.build.sets).toEqual([]);
    expect(decoded.build.loadouts).toEqual([]);
    expect(decoded.build.professions).toEqual([]);
    expect(decoded.build.ignored).toEqual([]);
  });

  it('reads an enchant and a suffix on a gear entry (contract 10.5), keeping the id map lossy', () => {
    const decoded = decodeFS1('FS1:1.15.9:warrior:orc:0/5530515/0:head=12640:2504,main_hand=11726:2505:1820');
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    // `gear` stays the planner's map of ids -- the strip, the planner link and the build
    // draft all read it and none of them models an enchant.
    expect(decoded.build.gear).toEqual({ head: 12640, main_hand: 11726 });
    // `gearSlots` is the whole truth, and is what SimCharacter and the request carry.
    expect(decoded.build.gearSlots).toEqual([
      { slot: 'head', itemId: 12640, enchant: 2504 },
      { slot: 'main_hand', itemId: 11726, enchant: 2505, suffix: 1820 },
    ]);
  });

  it('reads the professions section', () => {
    const decoded = decodeFS1(`${V1}|professions=engineering,blacksmithing`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.professions).toEqual(['engineering', 'blacksmithing']);
  });

  it('reads bags and bank, with the optional enchant and suffix', () => {
    const decoded = decodeFS1(`${V1}|bags=16963,17076:2504,19360:2505:1820|bank=12640`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.bags).toEqual([
      { itemId: 16963 },
      { itemId: 17076, enchant: 2504 },
      { itemId: 19360, enchant: 2505, suffix: 1820 },
    ]);
    expect(decoded.build.bank).toEqual([{ itemId: 12640 }]);
  });

  it('reads named sets whose gear is a gear list, enchants and suffixes included', () => {
    const decoded = decodeFS1(`${V1}|sets=AQ%20set=head=21329:2504,chest=21330;PvP=head=16963`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.sets).toEqual([
      {
        name: 'AQ set',
        gear: [
          { slot: 'head', itemId: 21329, enchant: 2504 },
          { slot: 'chest', itemId: 21330 },
        ],
      },
      { name: 'PvP', gear: [{ slot: 'head', itemId: 16963 }] },
    ]);
  });

  it('reads named loadouts as three trees each', () => {
    const decoded = decodeFS1(`${V1}|loadouts=Deep%20Fury=0/5530515/0;Arms=5530515/0/0`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.loadouts.map((row) => row.name)).toEqual(['Deep Fury', 'Arms']);
    expect(decoded.build.loadouts[0].treeRanks[1]).toEqual([5, 5, 3, 0, 5, 1, 5]);
  });

  it('ignores a section it does not know and reports its name', () => {
    const decoded = decodeFS1(`${V1}|bags=16963|quiver=1234|bank=12640`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.ignored).toEqual(['quiver']);
    expect(decoded.build.bags).toEqual([{ itemId: 16963 }]);
    expect(decoded.build.bank).toEqual([{ itemId: 12640 }]);
  });

  it('refuses an unreadable item entry rather than silently dropping it', () => {
    const decoded = decodeFS1(`${V1}|bags=16963,notanid`);
    expect(decoded.ok).toBe(false);
    if (decoded.ok) return;
    expect(decoded.message).toContain('notanid');
  });

  it('accepts a code up to the new bound and refuses one past it', () => {
    expect(MAX_CODE_LENGTH).toBe(16_384);
    const long = `${V1}|bags=${Array.from({ length: 900 }, () => '16963').join(',')}`;
    expect(long.length).toBeLessThan(MAX_CODE_LENGTH);
    expect(decodeFS1(long).ok).toBe(true);
    expect(decodeFS1('x'.repeat(MAX_CODE_LENGTH + 1)).ok).toBe(false);
  });
});

describe('encodeFS1V2', () => {
  it('is encodeFS1 exactly when there is nothing after the gear and nothing to enchant', () => {
    const build = {
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [5, 5, 3, 0, 5, 1, 5], []],
      gear: { head: 12640 },
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    };
    expect(encodeFS1V2(build)).toBe(encodeFS1(build));
  });

  it('writes an enchant and a suffix onto a gear entry, and round-trips them', () => {
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: { head: 12640 },
      gearSlots: [{ slot: 'head', itemId: 12640, enchant: 2504, suffix: 1820 }],
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    expect(code).toContain('head=12640:2504:1820');
    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) {
      expect(decoded.build.gearSlots).toEqual([
        { slot: 'head', itemId: 12640, enchant: 2504, suffix: 1820 },
      ]);
    }
  });

  it('writes the sections in the contract’s order and round-trips them', () => {
    const build = {
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [5, 5, 3, 0, 5, 1, 5], []],
      gear: { head: 12640 },
      bags: [{ itemId: 16963, enchant: 2504 }],
      bank: [{ itemId: 19360, enchant: 2505, suffix: 1820 }],
      sets: [{ name: 'AQ set', gear: [{ slot: 'head' as const, itemId: 21329 }] }],
      loadouts: [{ name: 'Deep Fury', treeRanks: [[], [5, 5, 3, 0, 5, 1, 5], []] }],
      professions: ['engineering', 'blacksmithing'],
      ignored: [],
    };
    const code = encodeFS1V2(build);
    expect(code.indexOf('|bags=')).toBeLessThan(code.indexOf('|bank='));
    expect(code.indexOf('|bank=')).toBeLessThan(code.indexOf('|sets='));
    expect(code.indexOf('|sets=')).toBeLessThan(code.indexOf('|loadouts='));
    expect(code.indexOf('|loadouts=')).toBeLessThan(code.indexOf('|professions='));

    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.bags).toEqual(build.bags);
    expect(decoded.build.bank).toEqual(build.bank);
    expect(decoded.build.sets).toEqual(build.sets);
    expect(decoded.build.loadouts[0].name).toBe('Deep Fury');
    expect(decoded.build.professions).toEqual(['engineering', 'blacksmithing']);
  });

  it('escapes a name carrying the separators it would otherwise break on', () => {
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      bags: [],
      bank: [],
      sets: [{ name: 'a;b=c|d', gear: [{ slot: 'head' as const, itemId: 1 }] }],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) expect(decoded.build.sets[0].name).toBe('a;b=c|d');
  });
});
```

Add `encodeFS1V2` and `MAX_CODE_LENGTH` to the file's existing `./fs1` import.

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/planner/fs1.test.ts
```

Expected: FAIL — `MAX_CODE_LENGTH` is not exported and `build.bags` is undefined.

- [ ] **Step 3: Extend `fs1.ts`**

Replace the constant and the `FS1Build` interface, then add the section parsing.

```ts
/**
 * Contract 7 raises this from 2,048: a version 2 code carries bags, bank, named sets and
 * named loadouts, and a full bank is several thousand characters on its own. It is still a
 * finite bound checked before any splitting, which is what keeps a hostile query value
 * from doing real work.
 */
export const MAX_CODE_LENGTH = 16_384;

/** One bag or bank item: `item_id[:enchant[:suffix]]`, with no slot in the string. */
export interface FS1Item {
  itemId: number;
  enchant?: number;
  suffix?: number;
}

/** One gear entry: `<slot>=item_id[:enchant[:suffix]]` (contract 10.5). */
export interface FS1GearSlot {
  slot: Slot;
  itemId: number;
  enchant?: number;
  suffix?: number;
}

export interface FS1Set {
  name: string;
  gear: FS1GearSlot[];
}

export interface FS1Loadout {
  name: string;
  /** One array per tree, one entry per talent in tab order -- FS1Build.treeRanks's shape. */
  treeRanks: number[][];
}

export interface FS1Build {
  dataBuild: string;
  classSlug: string;
  raceSlug: string;
  treeRanks: number[][];
  /**
   * The planner's map of slot to item id. Lossy by design: the strip, the planner link and
   * `BuildDraft` all read it and none of them models an enchant or a suffix.
   */
  gear: Gear;
  /**
   * The whole truth about the gear, enchants and suffixes included (contract 10.5). The
   * decoder always fills it; it is optional only so a caller holding nothing but a `Gear`
   * map can still build an `FS1Build` without writing `gearSlotsFrom(gear)` by hand.
   */
  gearSlots?: FS1GearSlot[];
  /** Version 2 (contract 7, corrected by 10.5). Empty for a version 1 code. */
  bags: FS1Item[];
  bank: FS1Item[];
  sets: FS1Set[];
  loadouts: FS1Loadout[];
  professions: string[];
  /** Section names the decoder did not recognise, reported rather than silently dropped. */
  ignored: string[];
}
```

The tree parsing and the gear parsing each move into a helper, because version 2 needs both again for loadouts and sets:

```ts
type Parsed<T> = { ok: true; value: T } | FS1Error;

function parseTrees(field: string): Parsed<number[][]> {
  const treeStrings = field.split('/');
  if (treeStrings.length !== TREES) {
    return { ok: false, message: `That code has ${treeStrings.length} talent trees; a build has ${TREES}.` };
  }
  const treeRanks: number[][] = [];
  for (const tree of treeStrings) {
    const ranks: number[] = [];
    for (const digit of tree) {
      const rank = Number.parseInt(digit, 36);
      if (Number.isNaN(rank)) {
        return { ok: false, message: `That code has an unreadable talent rank: ${digit}.` };
      }
      ranks.push(rank);
    }
    treeRanks.push(ranks);
  }
  return { ok: true, value: treeRanks };
}

/**
 * `<slot>=item_id[:enchant[:suffix]]`, joined by commas (contract 10.5). A bare id is
 * still legal, which is what keeps a version-1 string readable by this and a version-2
 * string's plain entries readable by a version-1 decoder.
 */
function parseGearList(field: string): Parsed<FS1GearSlot[]> {
  const slots: FS1GearSlot[] = [];
  if (field === '') return { ok: true, value: slots };
  for (const entry of field.split(',')) {
    const [slot, value] = entry.split('=');
    if (!(SLOTS as readonly string[]).includes(slot)) {
      return { ok: false, message: `That code names a slot this planner does not have: ${slot}.` };
    }
    // Digits only, never Number.parseInt on the whole field: parseInt stops at the first
    // non-digit and would silently turn "12640abc" into the item id 12640.
    const parts = (value ?? '').split(':');
    if (value === undefined || parts.length > 3 || parts.some((part) => !/^\d+$/.test(part))) {
      return { ok: false, message: `That code has an unreadable gear entry: ${entry}.` };
    }
    const [itemId, enchant, suffix] = parts.map((part) => Number.parseInt(part, 10));
    slots.push({
      slot: slot as Slot,
      itemId,
      ...(enchant === undefined ? {} : { enchant }),
      ...(suffix === undefined ? {} : { suffix }),
    });
  }
  return { ok: true, value: slots };
}

/** The planner's lossy view of a gear list: slot to item id, enchants dropped. */
function gearMapOf(slots: readonly FS1GearSlot[]): Gear {
  const gear: Gear = {};
  for (const entry of slots) gear[entry.slot] = entry.itemId;
  return gear;
}

/** The inverse, for a caller holding only the planner's map. */
export function gearSlotsFrom(gear: Gear): FS1GearSlot[] {
  return SLOTS.filter((slot) => gear[slot] !== undefined).map((slot) => ({
    slot,
    itemId: gear[slot] as number,
  }));
}

/** `item_id[:enchant[:suffix]]`, joined by commas. Each number is digits only, as gear is. */
function parseItemList(field: string): Parsed<FS1Item[]> {
  if (field === '') return { ok: true, value: [] };
  const items: FS1Item[] = [];
  for (const entry of field.split(',')) {
    const parts = entry.split(':');
    if (parts.length > 3 || parts.some((part) => !/^\d+$/.test(part))) {
      return { ok: false, message: `That code has an unreadable item entry: ${entry}.` };
    }
    const [itemId, enchant, suffix] = parts.map((part) => Number.parseInt(part, 10));
    items.push({
      itemId,
      ...(enchant === undefined ? {} : { enchant }),
      ...(suffix === undefined ? {} : { suffix }),
    });
  }
  return { ok: true, value: items };
}

/** `<name>=<payload>;…`, the name URL-encoded so it may carry `;`, `=` and `|`. */
function parseNamed(field: string): { name: string; payload: string }[] {
  if (field === '') return [];
  return field.split(';').map((entry) => {
    const split = entry.indexOf('=');
    return split === -1
      ? { name: decodeURIComponent(entry), payload: '' }
      : { name: decodeURIComponent(entry.slice(0, split)), payload: entry.slice(split + 1) };
  });
}
```

`decodeFS1` splits on `|` first, runs the existing version 1 parse on the head, and then walks the sections:

```ts
export function decodeFS1(code: string): FS1Result {
  if (code.length > MAX_CODE_LENGTH) {
    return { ok: false, message: 'That code is too long to read.' };
  }
  // Version 2 first, because everything before the first pipe is version 1 unchanged
  // (contract 7) -- so the version 1 parse below never has to know sections exist.
  const [head, ...sections] = code.trim().split('|');

  const parts = head.split(':');
  if (parts[0] !== FS1_PREFIX) {
    const named = parts[0] === undefined || parts[0] === '' ? 'unlabelled' : parts[0];
    return { ok: false, message: `That code is ${named}; this site reads ${FS1_PREFIX}.` };
  }
  if (parts.length < 6) return { ok: false, message: 'That code is missing its talent and gear fields.' };

  const [, dataBuild, classSlug, raceSlug, treeField, ...gearParts] = parts;
  const trees = parseTrees(treeField);
  if (!trees.ok) return trees;
  const gear = parseGearList(gearParts.join(':'));
  if (!gear.ok) return gear;

  const build: FS1Build = {
    dataBuild,
    classSlug,
    raceSlug,
    treeRanks: trees.value,
    gear: gearMapOf(gear.value),
    gearSlots: gear.value,
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
    professions: [],
    ignored: [],
  };

  for (const section of sections) {
    const split = section.indexOf('=');
    const name = split === -1 ? section : section.slice(0, split);
    const field = split === -1 ? '' : section.slice(split + 1);
    if (name === 'bags' || name === 'bank') {
      const items = parseItemList(field);
      if (!items.ok) return items;
      build[name] = items.value;
    } else if (name === 'sets') {
      for (const { name: setName, payload } of parseNamed(field)) {
        const setGear = parseGearList(payload);
        if (!setGear.ok) return setGear;
        build.sets.push({ name: setName, gear: setGear.value });
      }
    } else if (name === 'loadouts') {
      for (const { name: loadoutName, payload } of parseNamed(field)) {
        const loadoutTrees = parseTrees(payload);
        if (!loadoutTrees.ok) return loadoutTrees;
        build.loadouts.push({ name: loadoutName, treeRanks: loadoutTrees.value });
      }
    } else if (name === 'professions') {
      // Slugs, unvalidated here: IDS.md's profession list is the vocabulary and
      // sim/request refuses one it cannot map, naming it. Silently dropping a slug this
      // decoder did not recognise would hide exactly that error.
      build.professions = field === '' ? [] : field.split(',');
    } else if (name !== '') {
      // Contract 7: unknown sections are ignored by the decoder and reported in its
      // result. An addon a version ahead of the site is a thing that will happen, and
      // refusing its whole string would make the site useless the day it ships.
      build.ignored.push(name);
    }
  }

  return { ok: true, build };
}
```

And the encoder. `encodeFS1` keeps its exact current body and output; `encodeFS1V2` appends:

```ts
function encodeItems(items: readonly FS1Item[]): string {
  return items
    .map((item) =>
      [item.itemId, item.enchant, item.suffix]
        .filter((part): part is number => part !== undefined)
        .join(':'),
    )
    .join(',');
}

/** Version 1: bare ids, in SLOTS order. `encodeFS1`'s output is unchanged by 10.5. */
function encodeGearList(gear: Gear): string {
  return SLOTS.filter((slot) => gear[slot] !== undefined)
    .map((slot) => `${slot}=${gear[slot]}`)
    .join(',');
}

/** Version 2: `<slot>=item_id[:enchant[:suffix]]`, in SLOTS order (contract 10.5). */
function encodeGearSlots(slots: readonly FS1GearSlot[]): string {
  const bySlot = new Map(slots.map((entry) => [entry.slot, entry]));
  return SLOTS.filter((slot) => bySlot.has(slot))
    .map((slot) => {
      const entry = bySlot.get(slot) as FS1GearSlot;
      return `${slot}=${[entry.itemId, entry.enchant, entry.suffix]
        .filter((part): part is number => part !== undefined)
        .join(':')}`;
    })
    .join(',');
}

/**
 * Version 2: the version 1 string, then the sections in the contract's fixed order, each
 * omitted when empty. Names are URL-encoded, so a set called "a;b=c|d" survives -- the
 * three characters the grammar itself uses are the three a player is most likely to type.
 *
 * `encodeFS1` is untouched and still produces a version 1 string: the planner's own share
 * links and the "Sim this build" URL are version 1 and there is nothing in them to carry.
 */
export function encodeFS1V2(build: FS1Build): string {
  const sections: string[] = [];
  if (build.bags.length > 0) sections.push(`bags=${encodeItems(build.bags)}`);
  if (build.bank.length > 0) sections.push(`bank=${encodeItems(build.bank)}`);
  if (build.sets.length > 0) {
    sections.push(
      `sets=${build.sets
        .map((set) => `${encodeURIComponent(set.name)}=${encodeGearSlots(set.gear)}`)
        .join(';')}`,
    );
  }
  if (build.loadouts.length > 0) {
    sections.push(
      `loadouts=${build.loadouts
        .map(
          (loadout) =>
            `${encodeURIComponent(loadout.name)}=${encodeTrees(loadout.treeRanks)}`,
        )
        .join(';')}`,
    );
  }
  if (build.professions.length > 0) sections.push(`professions=${build.professions.join(',')}`);

  // The head is version 1 unless a gear entry actually has something to say beyond its id,
  // so a build with no enchants encodes byte-identically to `encodeFS1` -- which is what
  // the first test in this task asserts, and what keeps a version-2 string readable by a
  // version-1 decoder whenever it can be.
  const slots = build.gearSlots ?? gearSlotsFrom(build.gear);
  const rich = slots.some((entry) => entry.enchant !== undefined || entry.suffix !== undefined);
  const head = rich
    ? [FS1_PREFIX, build.dataBuild, build.classSlug, build.raceSlug, encodeTrees(build.treeRanks), encodeGearSlots(slots)].join(':')
    : encodeFS1(build);

  return [head, ...sections].join('|');
}
```

Two shared helpers fall out of this and both encoders use them, so neither can drift:

```ts
/** The three tree fields, slash-joined. `encodeFS1` and `encodeFS1V2` share it. */
function encodeTrees(treeRanks: readonly number[][]): string {
  return Array.from({ length: TREES }, (_, index) => encodeTree(treeRanks[index] ?? [])).join('/');
}
```

`encodeFS1`'s body becomes `[FS1_PREFIX, build.dataBuild, build.classSlug, build.raceSlug, encodeTrees(build.treeRanks), encodeGearList(build.gear)].join(':')` — the same output it has today, now expressed through the two helpers — and `encodeFS1V2`'s loadout section uses `encodeTrees(loadout.treeRanks)`.

- [ ] **Step 4: Raise the URL bound to match**

In `src/lib/sim/url.ts`, replace the `MAX_CODE` constant and its comment:

```ts
/**
 * An FS1 code is now version 2 (contract 7): three tree strings, seventeen gear entries,
 * and optionally a bag list, a bank list, named sets and named loadouts. `fs1.ts`'s own
 * decoder refuses anything over MAX_CODE_LENGTH for the same reason; this mirrors that
 * bound rather than importing it, since url.ts only ever needs to cap an
 * attacker-controlled query string before the code reaches a decoder at all.
 */
const MAX_CODE = 16_384;
```

- [ ] **Step 5: Carry the sections onto `SimCharacter`**

Append to `src/lib/sim/character.test.ts`:

```ts
describe('characterFromFs1 with version 2 sections', () => {
  it('carries bags, bank, sets, loadouts and professions onto the character', () => {
    const code = `${FURY_CODE}|bags=16963|bank=19360:2505:1820|sets=AQ=head=21329|loadouts=Arms=5530515/0/0|professions=engineering,alchemy`;
    const result = characterFromFs1(code, fixtureTalents, fixtureClasses, fixtureRaces, SOURCE);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.bags).toEqual([{ itemId: 16963 }]);
    expect(result.character.bank).toEqual([{ itemId: 19360, enchant: 2505, suffix: 1820 }]);
    expect(result.character.sets).toEqual([{ name: 'AQ', gear: [{ slot: 'head', itemId: 21329 }] }]);
    expect(result.character.loadouts[0].name).toBe('Arms');
    expect(result.character.professions).toEqual(['engineering', 'alchemy']);
  });

  it('carries a gear entry’s enchant and suffix (contract 10.5) onto gear_slots', () => {
    const result = characterFromFs1(
      'FS1:1.15.9.69722:warrior:orc:0/5530515/0:head=12640:2504:1820',
      fixtureTalents,
      fixtureClasses,
      fixtureRaces,
      SOURCE,
    );
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.gear_slots).toEqual([
      { slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 },
    ]);
    // The planner-facing map keeps the id alone, as every reader of it expects.
    expect(result.character.gear).toEqual({ head: 12640 });
  });

  it('gives a version 1 export the empty lists rather than leaving them undefined', () => {
    const result = characterFromFs1(FURY_CODE, fixtureTalents, fixtureClasses, fixtureRaces, SOURCE);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.character.bags).toEqual([]);
    expect(result.character.sets).toEqual([]);
    expect(result.character.professions).toEqual([]);
    expect(result.character.gear_slots.length).toBeGreaterThan(0);
  });
});

describe('toCharacterSpec with per-slot enchants and professions', () => {
  it('sends gear_slots verbatim rather than rebuilding the list from the id map', () => {
    const character = {
      ...fixtureCharacter,
      gear: { head: 12640 },
      gear_slots: [{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }],
      professions: ['engineering'],
    };
    const spec = toCharacterSpec(character, fixtureIndex, [], []);
    expect(spec.gear).toEqual([{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }]);
    expect(spec.professions).toEqual(['engineering']);
  });

  it('falls back to the id map for a character that has no slot list', () => {
    const character = { ...fixtureCharacter, gear: { head: 12640 }, gear_slots: [], professions: [] };
    const spec = toCharacterSpec(character, fixtureIndex, [], []);
    expect(spec.gear).toEqual([{ slot: 'head', item_id: 12640 }]);
    expect(spec.professions).toBeUndefined();
  });
});
```

(`FURY_CODE`, `fixtureTalents`, `fixtureClasses`, `fixtureRaces` and `SOURCE` are the file's existing helpers.)

In `src/lib/sim/character.ts`:

```ts
export interface SimCharacter {
  name: string;
  spec: string;
  class_slug: string;
  race_slug: string;
  talent_level: number;
  tree_version: string;
  point_order: number[];
  gear: Gear;
  buffs: string[];
  consumables: string[];
  source: CharacterSource;
  /**
   * The gear as the engine takes it, enchants and suffixes included (contract 10.5).
   * `gear` above stays the planner's map of ids -- the strip, the planner link and
   * `BuildDraft` all read it and none of them models an enchant -- and this is the
   * authoritative list `toCharacterSpec` sends. The two always agree on item ids.
   */
  gear_slots: GearSlot[];
  /** From the export's `professions=` section (contract 10.5). Empty for every other source. */
  professions: string[];
  /**
   * Version 2 export sections (contract 7, corrected by 10.5). Empty for every other
   * source and for a version 1 export. These are the candidate lists `/sim/gear`,
   * `/sim/talents` and `/sim/drops` read; nothing on `/sim` itself renders them. They keep
   * the decoder's own shapes, so part B converts an `FS1Set` into the envelope's `GearSet`
   * once, where it builds the bulk request.
   */
  bags: FS1Item[];
  bank: FS1Item[];
  sets: FS1Set[];
  loadouts: FS1Loadout[];
}
```

`characterFromFs1` fills them from the decoded build:

```ts
      gear: { ...decoded.build.gear },
      gear_slots: (decoded.build.gearSlots ?? []).map((entry) => ({
        slot: entry.slot,
        item_id: entry.itemId,
        ...(entry.enchant === undefined ? {} : { enchant: entry.enchant }),
        ...(entry.suffix === undefined ? {} : { suffix: entry.suffix }),
      })),
      professions: [...decoded.build.professions],
      buffs: [],
      consumables: [],
      source,
      bags: [...decoded.build.bags],
      bank: [...decoded.build.bank],
      sets: decoded.build.sets.map((set) => ({ name: set.name, gear: set.gear.map((e) => ({ ...e })) })),
      loadouts: decoded.build.loadouts.map((row) => ({
        name: row.name,
        treeRanks: row.treeRanks.map((tree) => [...tree]),
      })),
```

`fromBuildDraft` — the other constructor — gets the empty lists. A planner build has no
enchants, so its `gear_slots` is the id map converted, which keeps the "the two always
agree on item ids" promise true for every source:

```ts
    consumables: [...(extras.consumables ?? [])],
    source: extras.source,
    gear_slots: gearSlots(draft.gear ?? {}),
    professions: [],
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
```

and so does `SavedSim.svelte`'s `character` derived — where `gear_slots` is
`result.request.character.gear` itself, the richest form available on that page — plus any
other literal `astro check` names. Import the types in `character.ts`:

```ts
import { decodeFS1, orderFromRanks, type FS1Item, type FS1Loadout, type FS1Set } from '../planner/fs1';
```

Finally, `toCharacterSpec` sends the slot list and the professions:

```ts
  const spec: CharacterSpec = {
    name: character.name,
    race: character.race_slug,
    class: character.class_slug,
    level: SIM_LEVEL,
    talents: talentsString(index, character.point_order),
    // The slot list when the source gave one, the id map otherwise. Never both, and never
    // a merge: one of the two is the truth about this character's gear and it is this one.
    gear:
      character.gear_slots.length > 0
        ? character.gear_slots.map((slot) => ({ ...slot }))
        : gearSlots(character.gear),
    buffs: [...buffs],
    consumes: [...consumes],
    // Still omitted when empty, for the reason the original comment gives: an empty list
    // would claim we had looked and found none.
    ...(character.professions.length === 0 ? {} : { professions: [...character.professions] }),
  };
```

- [ ] **Step 6: Make the drawer's Apply lossless**

Contract 10.5 is what lifts Task 15's note. `applyRequest` round-trips the character through the FS1 route, so switching it from `encodeFS1` to `encodeFS1V2` with the request's own gear list carries the enchants and suffixes all the way back onto `SimCharacter.gear_slots`, and `toCharacterSpec` sends them again unchanged.

In `src/lib/sim/store.svelte.ts`:

```ts
    async applyRequest(request: SimRequest): Promise<void> {
      settings = settingsFromRequest(request);
      precisionId = precisionOf(request);
      // encodeFS1V2, not encodeFS1: contract 10.5 lets a gear entry carry
      // `item_id[:enchant[:suffix]]`, and `characterFromFs1` reads it straight back onto
      // `gear_slots`. Apply is therefore lossless for everything `CharacterSpec` models,
      // which is what makes a shared request a full reproduction rather than an
      // approximation of one.
      await adopt(
        fromPlannerCode(
          encodeFS1V2({
            dataBuild: init.treeVersion,
            classSlug: request.character.class,
            raceSlug: request.character.race,
            treeRanks: ranksFromTalentsString(request.character.talents),
            gear: gearFromSlots(request.character.gear),
            gearSlots: request.character.gear.map((slot) => ({
              slot: slot.slot as Slot,
              itemId: slot.item_id,
              ...(slot.enchant === undefined ? {} : { enchant: slot.enchant }),
              ...(slot.suffix === undefined ? {} : { suffix: slot.suffix }),
            })),
            bags: [],
            bank: [],
            sets: [],
            loadouts: [],
            professions: [...(request.character.professions ?? [])],
            ignored: [],
          }),
          ctx,
        ),
      );
    },
```

with `encodeFS1V2` replacing `encodeFS1` in the import and `import type { Slot } from '../planner/types';` added. Do the same in `SimView.svelte`'s `onRerunSaved`, which builds a code from a saved result's gear for exactly the same reason and loses exactly the same fields today.

Then replace the drawer's note in `src/lib/sim/copy.ts`:

```ts
  requestApplyNote:
    'Rebuilds the settings, the precision and the character from this request, enchants and suffixes included. Run sends the text exactly as typed instead, which is how a field no panel offers reaches the engine.',
```

and delete the `// Task 17 replaces this sentence …` comment above it.

- [ ] **Step 7: Prove Apply keeps an enchant**

Append to `tests/e2e/sim-request.spec.ts`:

```ts
test('Apply keeps a per-slot enchant and suffix', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const request = JSON.parse(await editor.inputValue());
  request.character.gear = [{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }];
  await editor.fill(JSON.stringify(request, null, 2));
  await page.getByTestId('sim-request-apply').click();

  // The drawer re-seeds from the page's own request only on its first render, so reopen
  // it on a fresh load to read what the page would now send.
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await page.getByTestId('sim-request-reset').click();
  const applied = JSON.parse(await editor.inputValue());
  expect(applied.character.gear).toEqual([
    { slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 },
  ]);
});
```

- [ ] **Step 8: Run the whole unit suite and the drawer e2e**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && npx astro check && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-request.spec.ts --project=desktop
```

Expected: PASS and 0 errors. `astro check` names every remaining `SimCharacter` literal that needs the new fields.

- [ ] **Step 9: Lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx eslint src/lib src/components tests/e2e && \
  npx prettier --check src/lib src/components tests/e2e && \
  git add -A src/lib src/components tests/e2e && \
  git commit -m "feat(sim): FS1 version 2 -- bags, bank, sets, loadouts, professions

Contract 7 as corrected by 10.5. Everything before the first pipe is
version 1 unchanged, so every existing decoder keeps working; the
sections after it are fixed in order and an unknown one is ignored and
named in the result rather than failing the whole string. A gear entry
may now carry item_id:enchant:suffix, SimCharacter gains gear_slots and
professions beside the planner's lossy id map, and the request drawer's
Apply is lossless because of it. MAX_CODE_LENGTH rises to 16,384 and
url.ts's own bound with it.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 18: History rows by kind, with the API's headline

Design 9.3 and contract 8: `GET /v1/sims?mine=1` accepts `kind=` and its rows carry `kind` and `headline`. The API composes the headline; the web renders it and filters.

**Files:**
- Create: `src/lib/sim/history.ts`
- Modify: `src/lib/sim/api.ts`
- Modify: `src/components/sim/SimHistory.svelte`
- Modify: `src/components/sim/SimView.svelte`
- Modify: `src/lib/sim/copy.ts`
- Modify: `src/test-support/sim-api.ts`
- Test: `src/lib/sim/history.test.ts`, `src/lib/sim/api.test.ts`, `tests/e2e/sim-history.spec.ts` (new)

**Interfaces:**
- Consumes: `kind.ts`'s `SIM_KINDS`, `SimKind` (Task 1).
- Produces:
  - `history.ts`: `type KindFilter = SimKind | 'all'`, `KIND_FILTERS: readonly KindFilter[]`, `kindOf(row: SimListRow): SimKind`, `headlineOf(row: SimListRow): string`, `titleOf(row: SimListRow): string`
  - `api.ts`: `listMySims(page?: number, apiBase?: string, kind?: KindFilter)`
- Test ids added: `sim-history-filter`, `sim-history-kind-<sim_id>`, `sim-history-headline-<sim_id>`.

- [ ] **Step 1: Write the failing test**

Create `src/lib/sim/history.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import { KIND_FILTERS, headlineOf, kindOf, titleOf } from './history';
import type { SimListRow } from './types';

const row = (over: Partial<SimListRow> = {}): SimListRow => ({
  sim_id: 'aaaaaaaaaaaa',
  spec: 'warrior-fury',
  dps: 1204.5,
  engine_version: 'edc0c8e9a',
  created_at: '2026-09-19T10:00:00Z',
  title: '',
  ...over,
});

describe('kindOf', () => {
  it('is the row’s kind', () => {
    expect(kindOf(row({ kind: 'gear' }))).toBe('gear');
  });

  it('is run for a row saved before the kind column existed, and for one the API mislabels', () => {
    expect(kindOf(row())).toBe('run');
    expect(kindOf(row({ kind: 'nonsense' }))).toBe('run');
  });
});

describe('headlineOf', () => {
  it('is the API’s own sentence when it sends one', () => {
    expect(headlineOf(row({ headline: '+41 DPS from Vis’kag' }))).toBe('+41 DPS from Vis’kag');
  });

  it('falls back to the figure, so a row before the API composed headlines still says something', () => {
    expect(headlineOf(row())).toBe('1,205 DPS');
    expect(headlineOf(row({ headline: '' }))).toBe('1,205 DPS');
  });
});

describe('titleOf', () => {
  it('is the player’s title, and the spec when they never gave one', () => {
    expect(titleOf(row({ title: 'Pre-raid' }))).toBe('Pre-raid');
    expect(titleOf(row())).toBe('Fury');
  });
});

describe('KIND_FILTERS', () => {
  it('is “all” and the contract’s five kinds, each named', () => {
    expect(KIND_FILTERS).toEqual(['all', 'run', 'gear', 'talents', 'drops', 'weights']);
    for (const filter of KIND_FILTERS) {
      expect(simCopy.kindLabel[filter], filter).toBeTruthy();
    }
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run src/lib/sim/history.test.ts
```

Expected: FAIL — `Failed to resolve import "./history"`.

- [ ] **Step 3: Add the copy**

```ts
  // --- Design 9.3: every kind in the history, with its headline. ---
  kindLabel: {
    all: 'Everything',
    run: 'Sim',
    gear: 'Top Gear',
    talents: 'Talents',
    drops: 'Droptimizer',
    weights: 'Stat weights',
  } as Record<string, string>,
  historyFilter: 'Show',
```

- [ ] **Step 4: Write `history.ts`**

```ts
// web/src/lib/sim/history.ts
// The saved-sim list, by kind. The API composes each row's headline (contract 8: run ->
// "1,204 DPS"; gear -> "+41 DPS from Vis'kag"; drops -> "3 upgrades on Ragnaros"), because
// only the API has the result blob in front of it -- the list route does not send one. The
// web renders whatever sentence arrives and falls back to the figure for a row saved
// before the column existed.
import { formatAmount } from '../report/format';
import { SIM_KINDS, type SimKind } from './kind';
import { specLabel } from './spec-label';
import type { SimListRow } from './types';

export type KindFilter = SimKind | 'all';

export const KIND_FILTERS: readonly KindFilter[] = ['all', ...SIM_KINDS];

/** A row's kind. Anything outside the vocabulary -- including absent -- reads as a run. */
export function kindOf(row: SimListRow): SimKind {
  return SIM_KINDS.find((kind) => kind === row.kind) ?? 'run';
}

export function headlineOf(row: SimListRow): string {
  const headline = row.headline ?? '';
  return headline === '' ? `${formatAmount(Math.round(row.dps))} DPS` : headline;
}

export function titleOf(row: SimListRow): string {
  return row.title === '' ? specLabel(row.spec) : row.title;
}
```

- [ ] **Step 5: Let the API call filter**

Append to `src/lib/sim/api.test.ts`:

```ts
describe('listMySims', () => {
  it('asks for every kind by default and for one when filtered', async () => {
    const api = createSimApi();
    api.install();
    await listMySims(1, 'https://api.test');
    expect(api.lastPath).toContain('/v1/sims?mine=1&page=1');
    expect(api.lastPath).not.toContain('kind=');
    await listMySims(1, 'https://api.test', 'gear');
    expect(api.lastPath).toContain('kind=gear');
    api.reset();
  });
});
```

If `createSimApi` does not already record the last path, add a `lastPath` field to it in `src/test-support/sim-api.ts` — one assignment in the handler it already has.

Then in `src/lib/sim/api.ts`:

```ts
export function listMySims(
  page: number = 1,
  apiBase: string = API_BASE_URL,
  kind: KindFilter = 'all',
): Promise<SimListPage> {
  // "all" is the absence of the parameter, not a value: contract 8 gives `kind=` a closed
  // vocabulary of five and adding a sixth for "no filter" would be a word the API has to
  // know about for no reason.
  const filter = kind === 'all' ? '' : `&kind=${kind}`;
  return call<SimListPage>(`/v1/sims?mine=1&page=${page}${filter}`, apiBase, simCopy.loadFailed);
}
```

with `import type { KindFilter } from './history';`.

- [ ] **Step 6: Render it**

`SimHistory.svelte` gains the filter and two cells:

```svelte
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { KIND_FILTERS, headlineOf, kindOf, titleOf, type KindFilter } from '../../lib/sim/history';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { SimListRow } from '../../lib/sim/types';

  let {
    rows,
    error,
    kind,
    onkind,
  }: {
    rows: SimListRow[] | null;
    error: string | null;
    kind: KindFilter;
    onkind: (next: KindFilter) => void;
  } = $props();
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-history">
  <div class="flex flex-wrap items-baseline justify-between gap-3">
    <h2 class="section-title text-[15px]">{simCopy.yourSims}</h2>
    <label class="flex items-center gap-2 text-[13px]">
      <span class="label text-muted">{simCopy.historyFilter}</span>
      <select
        class="border-line-warm rounded-control bg-raised text-text min-h-11 border px-2 text-[13px] md:min-h-9"
        value={kind}
        onchange={(event) => onkind(event.currentTarget.value as KindFilter)}
        data-testid="sim-history-filter"
      >
        {#each KIND_FILTERS as filter (filter)}
          <option value={filter}>{simCopy.kindLabel[filter]}</option>
        {/each}
      </select>
    </label>
  </div>
  {#if error !== null}
    <p role="alert" class="text-strong text-[13px]">{error}</p>
  {:else if rows === null}
    <p class="text-muted text-[13px]">{simCopy.historyLoading}</p>
  {:else if rows.length === 0}
    <p class="text-muted text-[13px]">{simCopy.historyEmpty}</p>
  {:else}
    <ul class="border-line bg-raised rounded-panel flex flex-col border">
      {#each rows as row (row.sim_id)}
        <li class="border-line-soft border-b last:border-b-0">
          <a
            class="grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-3 py-2 text-[14px] md:grid-cols-[88px_minmax(0,1.4fr)_minmax(0,1.2fr)_96px]"
            href={`/sim/${row.sim_id}`}
            data-testid={`sim-history-${row.sim_id}`}
          >
            <span class="pill pill-sample shrink-0" data-testid={`sim-history-kind-${row.sim_id}`}>
              {simCopy.kindLabel[kindOf(row)]}
            </span>
            <span class="text-strong truncate font-semibold">{titleOf(row)}</span>
            <!-- The API's own sentence replaces the bare DPS column: "3 upgrades on
                 Ragnaros" says what a Droptimizer row is and a number does not. -->
            <span class="text-muted truncate text-[13px]" data-testid={`sim-history-headline-${row.sim_id}`}>
              {headlineOf(row)}
            </span>
            <span class="tabular text-muted hidden text-right font-mono text-[13px] md:inline">
              {row.created_at.slice(0, 10)}
            </span>
          </a>
        </li>
      {/each}
    </ul>
  {/if}
</section>
```

`specLabel` is no longer used directly here (it moved into `titleOf`); drop the import if `astro check` flags it.

In `SimView.svelte`:

```ts
  let historyKind = $state<KindFilter>('all');

  async function loadHistory(): Promise<void> {
    historyError = null;
    try {
      const page = await listMySims(1, undefined, historyKind);
      historyRows = page.rows;
    } catch (error) {
      historyRows = null;
      historyError = error instanceof Error ? error.message : simCopy.loadFailed;
    }
  }

  function setHistoryKind(next: KindFilter): void {
    historyKind = next;
    historyRows = null;
    void loadHistory();
  }
```

```svelte
        <simHistoryLazy.current
          rows={historyRows}
          error={historyError}
          kind={historyKind}
          onkind={setHistoryKind}
        />
```

- [ ] **Step 7: Give the fixture a row per kind**

In `src/test-support/sim-api.ts`, replace the single list row with one per kind, and honour `kind=`:

```ts
    {
      method: 'GET',
      pattern: /\/v1\/sims(\?|$)/,
      respond: (_match, url) => {
        const wanted = new URL(url, 'https://api.test').searchParams.get('kind');
        const rows = [
          {
            sim_id: FIXTURE_SIM_ID,
            spec: 'warrior-fury',
            dps: fixtureResult.dps.mean,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-14T10:02:00Z',
            title: 'Raid-buffed, 3:00, single target',
            kind: 'run',
            headline: '1,204 DPS',
          },
          {
            sim_id: 'gearaaaaaaaa',
            spec: 'warrior-fury',
            dps: 1245,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-15T10:02:00Z',
            title: '',
            kind: 'gear',
            headline: '+41 DPS from Vis’kag',
          },
          {
            sim_id: 'dropsaaaaaaa',
            spec: 'warrior-fury',
            dps: 1210,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-16T10:02:00Z',
            title: '',
            kind: 'drops',
            headline: '3 upgrades on Ragnaros',
          },
          {
            sim_id: 'talentsaaaaa',
            spec: 'warrior-fury',
            dps: 1222,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-17T10:02:00Z',
            title: '',
            kind: 'talents',
            headline: '+18 DPS with ‘Deep Fury’',
          },
          {
            sim_id: 'weightsaaaaa',
            spec: 'warrior-fury',
            dps: 0,
            engine_version: fixtureResult.engine_version,
            created_at: '2026-09-18T10:02:00Z',
            title: '',
            kind: 'weights',
            headline: 'Crit 1.00 · Agility 0.87',
          },
        ];
        const filtered = wanted === null ? rows : rows.filter((row) => row.kind === wanted);
        return envelope({ rows: filtered, total: filtered.length, page: 1, per_page: 100 });
      },
    },
```

If the harness's `respond` takes only the regexp match, widen its signature to `(match, url)` — one parameter added where it is called, and every other handler ignores it.

- [ ] **Step 8: Write the e2e**

Create `tests/e2e/sim-history.spec.ts`:

```ts
import { expect, test } from '@playwright/test';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ data, error: null }),
});

const ROWS = [
  {
    sim_id: 'aaaaaaaaaaaa',
    spec: 'warrior-fury',
    dps: 1204,
    engine_version: 'edc0c8e9a',
    created_at: '2026-09-14T10:02:00Z',
    title: 'Raid-buffed, 3:00, single target',
    kind: 'run',
    headline: '1,204 DPS',
  },
  {
    sim_id: 'bbbbbbbbbbbb',
    spec: 'warrior-fury',
    dps: 1245,
    engine_version: 'edc0c8e9a',
    created_at: '2026-09-15T10:02:00Z',
    title: '',
    kind: 'gear',
    headline: '+41 DPS from Viskag',
  },
];

test('the history lists every kind and filters to one', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ user: { premium: false }, characters: [] })),
  );
  await page.route('**/v1/sims?mine=1*', (route) => {
    const kind = new URL(route.request().url()).searchParams.get('kind');
    const rows = kind === null ? ROWS : ROWS.filter((row) => row.kind === kind);
    return route.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 }));
  });

  await page.goto('/sim');
  await expect(page.getByTestId('sim-history')).toBeVisible();
  await expect(page.getByTestId('sim-history-kind-aaaaaaaaaaaa')).toHaveText('Sim');
  await expect(page.getByTestId('sim-history-kind-bbbbbbbbbbbb')).toHaveText('Top Gear');
  await expect(page.getByTestId('sim-history-headline-bbbbbbbbbbbb')).toHaveText('+41 DPS from Viskag');

  await page.getByTestId('sim-history-filter').selectOption('gear');
  await expect(page.getByTestId('sim-history-bbbbbbbbbbbb')).toBeVisible();
  await expect(page.getByTestId('sim-history-aaaaaaaaaaaa')).toHaveCount(0);
});
```

- [ ] **Step 9: Run both suites, check, lint, format, commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx vitest run && \
  E2E_PORT=4399 npx playwright test tests/e2e/sim-history.spec.ts --project=desktop && \
  npx astro check && npx eslint src tests/e2e && npx prettier --check src tests/e2e && \
  git add -A src tests/e2e && \
  git commit -m "feat(sim): history rows by kind, with the API's own headline

Design 9.3. Each row carries its kind as a pill and the API's sentence
in place of the bare DPS column -- \"3 upgrades on Ragnaros\" says what a
Droptimizer row is and a number does not. The filter sends kind= and
\"Everything\" sends nothing, since the contract's vocabulary has five
values and no sixth for \"no filter\".

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

### Task 19: Lighthouse — the settings sheet must not regress CLS

The `/sim.html` budget is in `lighthouserc.json`'s own matrix row and the shell's reservation is `sim.astro`'s `min-h-[1044px] md:min-h-[531px]` on `#sim`, measured against the real signed-out switcher. This task proves the whole plan's additions did not move it, and puts a check in place so the next one cannot either.

Why the additions should be safe, and what to verify rather than assume:

- Everything this plan adds below the character strip — the settings sheet, the buff panel, the request drawer, the details card, the rotation card — renders only once a character is loaded, which is never the first paint. The LCP element on a cold `/sim` is `SourceSwitcher`'s addon-card paragraph, which nothing here touches.
- Every disclosure is closed on first render, so the collapsed height is one summary row, and opening one is a user gesture — the one class of layout shift CLS excludes.
- The history panel now has a filter row above its list. It renders for a signed-in visitor only, and Lighthouse measures signed out, but the reservation still has to be right for a real member, so its own height is checked.

**Files:**
- Modify: `src/pages/sim.astro` (only if the measurement says so)
- Modify: `tests/e2e/sim-phone.spec.ts`
- Test: a measurement run of `lhci`

- [ ] **Step 1: Measure the signed-out shell against the hydrated island**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npm run build && E2E_PORT=4399 npm run preview -- --port 4399 &
```

Then, in a second shell:

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx playwright test --project=mobile -g 'never scrolls sideways' tests/e2e/sim-phone.spec.ts
```

and measure `#sim` before and after hydration with a one-off script:

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && node --input-type=module -e "
import { chromium } from 'playwright';
const browser = await chromium.launch();
for (const width of [360, 1280]) {
  const page = await browser.newPage({ viewport: { width, height: 800 } });
  await page.route('**/sim-island.js', (route) => route.abort());
  await page.goto('http://localhost:4399/sim');
  const before = await page.locator('#sim').boundingBox();
  await page.unroute('**/sim-island.js');
  await page.reload();
  await page.waitForSelector('[data-testid=\"sim-view\"]');
  const after = await page.locator('#sim').boundingBox();
  process.stdout.write(width + ': shell ' + before.height + ' hydrated ' + after.height + '\n');
  await page.close();
}
await browser.close();
"
```

Expected: the two heights agree at each width, within a pixel. If they do not, change the `min-h-*` on `#sim` in `sim.astro` to the hydrated number — never the other way around, and never by removing the reservation.

- [ ] **Step 2: Run the Lighthouse budget**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npm run lhci
```

Expected: `/sim.html`, `/sim/specs.html` and `/sim/<id>.html` all pass their matrix row, with `cumulative-layout-shift` at or under `0.05`.

- [ ] **Step 3: Pin the sheet's closed state in the phone audit**

Add to `tests/e2e/sim-phone.spec.ts`, in the loaded-character case, before the disclosure is opened:

```ts
  // Every disclosure this lane adds is closed on arrival. An open one would add height
  // after hydration, which is the shift the shell's own min-h reservation exists to stop.
  for (const testid of ['sim-settings-more', 'sim-request-drawer']) {
    await expect(page.getByTestId(testid)).toBeVisible();
    await expect(page.getByTestId(testid)).not.toHaveAttribute('open', '');
  }
```

- [ ] **Step 4: Run the phone audit and the whole browser suite**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && E2E_PORT=4399 npx playwright test --project=mobile && \
  E2E_PORT=4399 npx playwright test --project=desktop
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd /Users/jh/code/forever/web && npx astro check && npx eslint src tests && npx prettier --check src tests && \
  git add -A src tests && \
  git commit -m "test(sim): the settings sheet and the request drawer keep CLS at zero

Every disclosure this lane adds is closed on first render, and the
phone audit now asserts it: an open one would add height after
hydration, which is exactly the shift sim.astro's min-h reservation
exists to prevent. The shell reservation is re-measured against the
hydrated island at 360 and 1280.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>"
```

---

## Interfaces → Produces, for part B

Everything part B (`/sim/gear`, `/sim/talents`, `/sim/drops`, `/sim/weights`) may import from this lane, with the task that lands it.

| Module | Exports |
| --- | --- |
| `src/lib/sim/types.ts` (1) | `EncounterSpec`, `Movement`, `TargetCount`, `TARGET_TYPE_IDS`, `TargetType`, `CooldownSpec`, `CharacterSpec`, `GearSlot`, `SimRequest`, `BulkSpec`, `Candidate`, `TalentLoadout`, `GearSet`, `WeightsSpec`, `SimResult`, `Combo`, `Substitution`, `Stage`, `StatWeight`, `SampleCast` (A12: `{ at_ms, action, target?, resources? }`), `SimProgress`, `SimListRow`, `DEFAULT_ENCOUNTER`, `STEP_ITERATIONS_DEFAULT` |
| `src/lib/sim/kind.ts` (1) | `SimKind`, `SIM_KINDS`, `requestKind(request)` |
| `src/lib/sim/styles.ts` (2) | `FightStyleId`, `FightStyle`, `FIGHT_STYLES`, `DEFAULT_STYLE_ID`, `fightStyle(id)`, `applyFightStyle(encounter, id)` |
| `src/lib/sim/settings.ts` (3) | `SimSettings`, `defaultSettings()`, `withStyle`, `withDuration`, `withTargets`, `withVariation`, `withTargetLevel`, `withTargetArmor`, `withTargetType`, `withDummy`, `withExecutePhase`, `withPreset`, `executePhaseOn`, `styleIdOf`, `settingsLabel`, `durationLabel`, `encounterLabel`, `DURATIONS`, `MIN_DURATION_SEC`, `MAX_DURATION_SEC`, `MAX_TARGETS`, `VARIATIONS`, `MAX_VARIATION`, `TARGET_LEVELS`, `DEFAULT_TARGET_LEVEL`, `TARGET_ARMOR_BY_LEVEL` (A8), `MAX_TARGET_ARMOR`, `TARGET_TYPES`, `BUFF_PRESETS`, `PRESET_BUFFS`, `PRESET_CONSUMABLES` |
| `src/lib/sim/engine.ts` (5) | `EngineModule` (with `simNeedsMore`, `simValidate`, `simCount`), `RequestValidation`, `RequestValidationError`, `CountAnswer` |
| `src/lib/sim/worker.ts` (5) | `SimPool` (with `needsMore`, `validate`, `count`), `ToWorker`, `FromWorker`, `createPool`, `poolSize`, `MAX_WORKERS` |
| `src/lib/sim/precision.ts` (6) | `Lane`, `PrecisionId`, `PRECISIONS`, `BULK_PRECISIONS`, `PRECISION_ITERATIONS`, `TARGET_ERROR`, `STEP_ITERATIONS`, `LANE_ITERATION_CEILING`, `CAPS` (A2), `PrecisionPlan`, `precisionPlan(id, lane)`, `precisionOf(request)`, `relativeError(estimate)` |
| `src/lib/sim/run.ts` (6) | `RunInput` (with `targetError`, `stepIterations`), `RunUpdate` (with `relativeError`), `RunHandle`, `buildSimRequest`, `runSim`, `SimRunError` |
| `src/lib/sim/details.ts` (7) | `RunDetails`, `runDetails(result)`, `percentLabel(fraction)` |
| `src/lib/sim/buffs.ts` (8, 10) | `BuffGroupId`, `BUFF_GROUPS`, `BuffRow`, `SimIdsFile`, `buildCatalogue(ids)`, `CATALOGUE`, `rowsIn(group)`, `BuffGrade`, `IMPROVED_SUFFIX`, `gradeOf`, `setGrade`, `Selection`, `withGrade`, `selectedIn`, `FALLBACK_WORLD_BUFFS` |
| `src/lib/sim/stats.ts` (8) | `PINNED_STATS` (contract 10.8, verbatim), `SIM_STATS`, `statLabel(id)` — the `proto.Stat` vocabulary for `/sim/weights`: one `hit`, one `crit`, split haste, `mp5` |
| `src/lib/sim/buff-names.ts` (9) | `BuffNames`, `EMPTY_BUFF_NAMES`, `loadBuffNames(build)`, `buffLabel(id, names)`, `buffIcon(build, id, names)` |
| `src/lib/sim/cooldowns.ts` (11) | `CooldownMode`, `COOLDOWN_MODES`, `CooldownRow`, `executeStartSec`, `specFor`, `modeOf`, `rowsFor`, `withCooldown` |
| `src/lib/sim/sample-log.ts` (12) | `SampleRow`, `SampleLog`, `RESOURCE_ORDER`, `sampleTime`, `sampleLog(sample, names)` — resolves A12's action keys |
| `src/lib/sim/notify.ts` (14) | `Notifier`, `browserNotifier()`, `enableNotifications(notifier)`, `notifyFinished(notifier, enabled, title, body)` |
| `src/lib/sim/request-json.ts` (15) | `MAX_REQUEST_CHARS`, `formatRequest(request)`, `ParsedRequest`, `parseRequest(text)`, `settingsFromRequest(request)` |
| `src/lib/sim/url.ts` (16) | `SimState` (with `req`), `defaultSimState`, `parseSimState`, `simSearch`, `withSimState`, `MAX_REQUEST_PARAM`, `encodeRequestParam(request)`, `decodeRequestParam(value)` |
| `src/lib/planner/fs1.ts` (17) | `MAX_CODE_LENGTH`, `FS1Item`, `FS1GearSlot`, `FS1Set` (`gear: FS1GearSlot[]`), `FS1Loadout`, `FS1Build` (with `gearSlots?`, `bags`, `bank`, `sets`, `loadouts`, `professions`, `ignored`), `gearSlotsFrom(gear)`, `decodeFS1`, `encodeFS1`, `encodeFS1V2`, `orderFromRanks` |
| `src/lib/sim/character.ts` (17) | `SimCharacter` (with `gear_slots`, `professions`, `bags`, `bank`, `sets`, `loadouts`), `toCharacterSpec(character, index, buffs, consumes, cooldowns?)`, `characterFromFs1`, `fromBuildDraft`, `gearSlots`, `gearFromSlots`, `ranksFromTalentsString`, `talentsString`, `specOf`, `SIM_LEVEL` |
| `src/lib/sim/history.ts` (18) | `KindFilter`, `KIND_FILTERS`, `kindOf(row)`, `headlineOf(row)`, `titleOf(row)` |
| `src/lib/sim/api.ts` (18) | `listMySims(page?, apiBase?, kind?)`, plus the existing `saveSim`, `fetchSim`, `dispatchServerSim`, `fetchSimProgress`, `fetchSpecs`, `fetchSimInput`, `SimApiError` |
| `src/lib/sim/store.svelte.ts` (6, 10, 14, 15, 16) | `createSimStore(init)`, `SimStoreInit` (with `request?`), `SimPhase`; on the store: `precisionId`, `setPrecisionId`, `relativeError`, `buffNames`, `reportTitle`, `setReportTitle`, `buildRequest()`, `validateRequest(json)`, `applyRequest(request)`, `runRequest(request)`, plus everything it already exposed |
| `src/components/sim/` | `SettingsBar.svelte`, `SettingsSheet.svelte` (4), `BuffPanel.svelte` (10), `CooldownRows.svelte` (11), `DetailsCard.svelte` (7), `SampleLog.svelte` (12), `RotationCard.svelte` (13), `RequestDrawer.svelte` (15), `SimHistory.svelte` (18) |
| `src/fixtures/sim/` | `styles.json` (2), `envelope-v2.json` (1), `sample.json` (12), and `engine-fake.ts` with `simNeedsMore`/`simValidate` (5) |

Part B should **not** re-derive any of the following: a stopping rule for a target-error run (ask `pool.needsMore`), request validation (ask `pool.validate`), a combination count or a cap check (ask `pool.count`, whose `cap_exceeded` answer carries both numbers), the fight-style expansion (`applyFightStyle`), the id grouping (`CATALOGUE`), the stat vocabulary (`SIM_STATS`), or the ranking and grouping of bulk combos (the sim module's `simRank`).

Two shapes part B converts rather than this lane: an `FS1Set` (`{ name, gear: FS1GearSlot[] }`, the decoder's own shape) into the envelope's `GearSet` (`{ name, gear: GearSlot[] }`), and an `FS1Item` (no slot, because the export string carries none) into a `Candidate` once it has resolved the slot from the item table.

## Rulings applied, and what is still open

Contract section 10 (commit `439f0b7`) settled every question the first draft of this plan
raised. This section records how each ruling landed, so an executor reading one task does
not have to diff two specs to know why it says what it says.

**Applied.**

1. **10.2 ratifies `simNeedsMore` and `simValidate`** with exactly the shapes Task 5 had
   pinned; the task's "contract addition" caveat is gone and it now cites 10.2.
2. **A2, A4, A5 and A6 are part B's rulings, but their shapes live in this lane's files**,
   so Task 1 mirrors `BulkSpec.consumables`, `Candidate.source_name` and
   `Substitution.source_name`/`name`, records A4's "the mode decides the expansion" on
   `BulkSpec.mode`, and Task 6 publishes `CAPS = {browser: 400, server: 5000}` beside the
   iteration ceilings. Nothing on `/sim` reads any of them; they are here so part B does
   not keep a second copy.
3. **10.2 adds `simCount`.** Task 5 lands it beside the other two — same interface, same
   pool message, same fake engine — so part B does not have to edit those five files
   again. Its `cap_exceeded` answer keeps `cap` and `combinations` instead of being thrown
   as an `Error`, because the cap notice renders both numbers.
4. **A3 relaxes the iteration rule**: a fixed run keeps `ValidIterations`, a `TargetError`
   run's `Iterations` is a positive multiple of `StepIterations` at or under the lane's
   ceiling, and `MinDurationSec`/`MaxDurationSec` are 20 and 600. Task 6's workaround is
   gone and a test now pins both ceilings as multiples of the step; Task 3's duration
   comment cites A3 rather than asking for an amendment.
5. **A8 gives `TargetArmorByLevel` real numbers** (`{60: 3300, 61: 3444, 62: 3588, 63:
   3731}`). Task 3 publishes them and Task 4's armor field names the figure it is about to
   override instead of saying "Preset for the level".
6. **A12 changes `SampleCast`** to `{ at_ms, action, target, resources }`. Tasks 1 and 12,
   `envelope-v2.json` and `sample.json` all carry an action key now, resolved through
   `resolveActionName` exactly as a cast row is, and a test covers `other:melee`.
7. **10.4 names `simbuffs.json`** with the shape Task 9 proposed. The task cites 10.4, and
   the humanised fallback stays because the file is `required: false` in the sync.
8. **10.5 puts per-slot enchant and suffix on `SimCharacter`** and lets `<gear>` and
   `sets=` entries carry `item_id[:enchant[:suffix]]`, with a `professions=` section after
   `loadouts=`. Task 17 now adds `FS1GearSlot`, `FS1Build.gearSlots`, `FS1Build.professions`,
   `SimCharacter.gear_slots`, `SimCharacter.professions`, and upgrades the request drawer's
   "Apply" to `encodeFS1V2` so it is lossless. Task 15 states the limitation as temporary
   and names the task that lifts it.
9. **A7 has the sim module generate IDS.md's `:improved` rows, World buffs and a new Stats
   section.** Task 8's parser reads all three; the new `stats.ts` carries the stat
   vocabulary for part B's weights page and asserts there is no bare `haste`.
10. **10.7 makes section 9's test ids a minimum.** The page-local ids this plan adds are
    listed below and follow the same `sim-<thing>` shape; no amendment is needed.
11. **10.8 adds `consumes` to `Substitution.kind`** (with `name` the consumable ids joined
    by `, `) and **pins the stat vocabulary**. Task 1 mirrors the kind and names all four
    in `simCopy.substitutionKindLabel`; Task 8's `stats.ts` carries 10.8's list verbatim as
    `PINNED_STATS`, and its test pins one `hit`, one `crit`, split haste and `mp5`, and
    refuses `melee_hit`, `spell_hit`, `melee_crit`, `spell_crit` and a bare `haste`
    anywhere on the lane.
12. **A11 keeps the progress widening additive.** `SimProgress` mirrors `stage`,
    `combos_done` and `combos_total`, and nothing on `/sim` reads them — `sim-stage-progress`
    is part B's, as before.

13. **10.8's client-side server cap is a no-op on `/sim` and is deliberately left to part
    B.** A plain run expands to no combinations, so `Caps.server` gates nothing on the
    page this lane owns; `CAPS` is exported from `precision.ts` (Task 6) so the bulk pages
    gate their own server-run button on it, and a `cap_exceeded` that still arrives falls
    through to `simCopy.failed`, which is what `runOnServer` already does with every
    server error.

**Still open — none blocking, each with the behaviour this plan ships meanwhile.**

- **Major-cooldown spell ids are still unpublished.** No section of the contract gives the
  web a per-class list, so Task 11 offers a timing row for every ticked potion and
  explosive plus any `spell:<id>` an edited request already carries, and says so in the
  panel's own copy. A published list makes the rows appear with no change to
  `cooldowns.ts`.
- **Design 5.1's "the aura table gains a count column where it lacks one" describes a
  column that already exists.** `report/AuraTable.svelte` renders `track.applications` as
  "Applied" for both the buff and debuff tabs, so Task 12 adds only the testid its uptime
  cell already had. Nothing in section 10 contradicts this; it is recorded here so the
  gap between the design's wording and the code is deliberate rather than missed.

**Page-local test ids this plan adds** (10.7 permits them; listed so part B reuses the
shapes rather than inventing parallel ones): `sim-style-note`, `sim-settings-more`,
`sim-variation`, `sim-target-level`, `sim-target-armor`, `sim-target-type`, `sim-dummy`,
`sim-buff-panel`, `sim-buff-group-<group>`, `sim-buff-<id>`, `sim-cooldowns`,
`sim-cooldown-<id>`, `sim-cooldown-at-<id>`, `sim-details-<row>`, `sim-details-engine`,
`sim-details-ceiling`, `sim-sample-empty`, `sim-sample-row-<key>`, `sim-rotation-card`,
`sim-rotation-card-link`, `sim-rotation-card-note`, `sim-report-title`, `sim-notify`,
`sim-open-new-tab`, `sim-request-json`, `sim-request-errors`, `sim-request-error-<field>`,
`sim-request-valid`, `sim-request-apply`, `sim-request-run`, `sim-request-share`,
`sim-request-share-link`, `sim-request-share-error`, `sim-request-reset`,
`sim-request-buffs`, `sim-history-filter`, `sim-history-kind-<id>`,
`sim-history-headline-<id>`, and `aura-applied` on the shared report table.
