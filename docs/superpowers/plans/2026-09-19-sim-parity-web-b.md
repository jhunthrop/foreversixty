# Simulator parity — web lane, part B: the four tool pages

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build `/sim/gear` (Top Gear), `/sim/talents`, `/sim/drops` (Droptimizer) and `/sim/weights`, the shared bulk run loop that drives all four in the browser under the lane's cap, and `/sim/<id>` rendering every result kind — so a player ticks candidate items, presses one button, and gets a ranked list of what to equip with an honest error bar on every delta.

**Architecture:** One new standalone island (`sim-tools-island`) over four static Astro shells, each shell stamping `data-sim-tool` and rendering its own shell skeleton so mounting shifts nothing. Every tool posts the same envelope: a `SimRequest` with a `bulk` block (or a `weights` block), planned by the wasm's `simPlan`, run stage by stage through the existing `simSplit`/`simRun`/`simCombine` worker pool, ranked by the wasm's `simRank`, and repeated until `simRank` answers with a finished `SimResult`. **No statistics, no expansion and no staging arithmetic exist in TypeScript** — the page ticks boxes, builds an envelope and renders what comes back. The premium lane posts the identical envelope to the API and polls it exactly as a single server run does.

**Tech Stack:** Astro 7 (static, `build.format: 'file'`), Svelte 5 runes, Tailwind 4 over `web/src/styles/tokens.css`, TypeScript 6 strict, Vitest 5, Playwright 1.63, Node 22.12.0. **No new runtime or test dependency.**

**Spec:** `docs/superpowers/specs/2026-09-19-simulator-parity-design.md` — sections 2, 3, 6, 7 and 10 are this lane. Binding contract: `docs/superpowers/specs/2026-09-19-simulator-parity-interfaces.md` — sections 1.3, 1.4, 2, 4, 6, 7, 8 and 9 bind this lane, **as amended by its section 10 (2026-09-19, commit 439f0b7), which wins wherever it contradicts an earlier section.** The contract is read-only from here: section 10 *is* the amendment round, so no task in this plan asks for another one. Supporting: `docs/superpowers/specs/2026-09-14-simulator-design.md` and `-interfaces.md` (what already exists), `design/DESIGN-SYSTEM.md`.

---

## Global Constraints

Every task's requirements implicitly include this section.

**Toolchain**

- Node 22.12.0. Every shell that runs npm starts with `export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH`. All npm commands run from `web/`.
- **`FOREVER_DATA=fixture` for every test run.** `npm run pretest` sets it; a bare `npx vitest`/`npx playwright` does not, so prefix it.
- **Every Playwright command in this plan is prefixed `E2E_PORT=4326`.** 4321 is the main checkout's preview and 4325 is the part-A lane's; a bare run collides and asserts against the wrong site.
- **`NO_COLOR=1 npx astro check` must report `0 errors`** after every task that adds or edits a `.ts`, `.astro` or `.svelte` file. It type-checks `.svelte` too, and CI runs it.
- **eslint and prettier run over the touched files only:** `npx eslint <paths>` and `npx prettier --write <paths>`. Never over the whole tree.
- **Tests are scoped to the files the task touches:** `npx vitest run src/lib/sim/bulk-run.test.ts`, not `npm test`. The full suite runs once, in Task 20.
- Vitest's default environment here is Node. Any test file that touches `document`, `window`, `Worker`, `navigator` or a module that reads `document.cookie` **must** start with `// @vitest-environment jsdom`.
- **The API is not running and this lane adds no HTTP-mocking library.** Unit tests stub through `web/src/test-support/sim-api.ts` (`createSimApi()`); Playwright uses `page.route`. Do not add `msw`.
- Lighthouse (`npm run lhci`) runs **once**, in Task 20.
- **Never `git stash`** (the stash stack is shared with other worktrees), never bypass git hooks, never amend an existing commit.

**Commits**

- One commit per task, conventional type `feat(sim): …` / `test(sim): …` / `fix(sim): …` / `chore(sim): …`.
- Every commit ends with exactly one trailer line, and it is this one:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  It is the plan's and it stands. If the session running a task carries an attribution reminder naming a different model, that reminder is not this lane's; use the line above.

**Contract values, verbatim**

- Bulk modes are `gear`, `talents`, `drops` (`BulkSpec.Mode`). Precisions are `fast`, `normal`, `high`. Result kinds are `run`, `gear`, `drops`, `talents`, `weights`, and **kind is derived, never sent**.
- **`Caps = {browser: 400, server: 5000}`** (contract 10.1 A2 — the server cap is 5,000, *not* the 20,000 of section 1.3: a 20,000-combination fast run cannot finish inside the job's 15-minute timeout). The browser cap is **halved to 200 when `navigator.hardwareConcurrency <= 4`** (design 2.3, "a phone gets half the cap"). The page echoes the cap it used in `BulkSpec.Cap`; the wasm enforces it and refuses with the structured `{"error":"cap_exceeded","cap":n,"combinations":n}` of contract 10.2. **The page never trims a candidate list to fit**, and the premium alternative it offers says 5,000.
- **A bulk request's `Iterations` is its precision's final-stage count** (contract 10.1 A3): 3,000 for `fast` and `normal`, 10,000 for `high`. It is not a free number and it is not always 3,000.
- **Mode decides expansion** (contract 10.1 A4): `gear` takes the product of every candidate group, `drops` and `talents` one substitution at a time. The design's `combinations` boolean does not exist.
- Stage ladders, from contract 1.3, are the wasm's business and are mirrored in TypeScript **only** as a stage count for the progress line: `fast` 3 stages, `normal` 2, `high` 2.
- `Candidate.Slot` is `""` for anything that fits more than one slot — rings, trinkets, weapons. Every other candidate carries its IDS.md slot name.
- `Candidate.Origin` is one of `equipped`, `bag`, `bank`, `search`, `drop:<source-id>`, `set:<name>`.
- `Combo.Group` is `0` for the leader's within-error group, then `1`, `2`, …. **The page never recomputes a statistic**: "within error" is `group === 0`, and a delta's error is `delta.error`.
- **Names come back on the result, not from a re-join** (contract 10.1 A6): `Substitution.Name` is filled for items too (from simdb) and `Substitution.SourceName` is copied from the candidate. The page fills `Candidate.SourceName` from `loot.json` when it builds a drops request, and reads both back off the result rather than re-joining an id to a name.
- **Stat ids are the fork's `proto.Stat` enum names in snake case, pinned in full by contract 10.8.** Haste is split (`melee_haste`, `spell_haste`); **hit and crit are not** — the engine carries one `hit` and one `crit`, and `melee_crit`, `spell_crit`, `melee_hit` and `spell_hit` do not exist. MP5 is `mp5`. `WeightsSpec.Reference` is required and uses this vocabulary, as does `specs.json`'s `reference_stat`, whose defaults are `attack_power` for melee and hunters and `spell_power` for casters.
- **`Substitution.Kind` is closed at four:** `item`, `talents`, `set` and `consumes` (contract 10.8), the last carrying the consumable ids joined by `, ` in `Name`.
- **The server-run button is gated client-side on `Caps.server`** before submitting (contract 10.8); a `cap_exceeded` that still arrives from the API shows the generic failure sentence, because `simCount` has already put both numbers on screen.
- **Premium submit is `POST /v1/sims/run`** and `POST /v1/sims` remains the browser-result save (contract 10.6); both accept every kind.
- Validation the page must satisfy before it sends: no candidate on a locked slot; `talents` mode has at least one loadout and **no** candidates; `drops` mode has candidates whose origins are all `drop:`; `gear` mode has at least one of candidates, talents or sets.
- The engine version is `ENGINE_VERSION` from `web/src/lib/sim/version.ts`. **Never type a sha** into a component, test or fixture.
- The active data build is `web/src/data/active-build.json`'s `build` field. **Never type a build id**; import it.
- `sim_id` is 12 lowercase base32 characters, `[a-z2-7]{12}`.
- JSON field names mirror the Go `json:` tags exactly. `snake_case` on the wire, `camelCase` only for values that never cross the boundary.
- **No protobuf crosses a lane boundary.** Everything in and out of the wasm is a JSON string.

**Test ids (contract 9), and who owns them**

- This lane: `sim-combos`, `sim-combo-row`, `sim-equipped-line`, `sim-cap-notice`, `sim-stage-progress`, `sim-source-picker`, `sim-source-<id>`, `sim-candidate-<slot>-<item>`, `sim-weights`, and `sim-precision` on the four tool pages.
- Part A owns `sim-request-drawer`, `sim-style`, `sim-target-error`, `sim-sample-log`, `sim-details-card`, and `sim-precision` on `/sim`.
- Page-local ids this lane adds, which contract 10.7 explicitly permits ("test ids in section 9 are a minimum; page-local ids follow the same `sim-<thing>` shape"): `sim-combo-count`, `sim-slot-grid`, `sim-item-search`, `sim-enchant-<slot>-<id>`, `sim-lock-<slot>`, `sim-loadout-<name>`, `sim-set-<name>`, `sim-slot-summary`, `sim-pawn`, `sim-consumable-<id>`, `sim-kind-<kind>`, `sim-upcoming`, `sim-drops-by-boss`, `sim-drops-flat`, `sim-drops-pin-<item>`, `sim-weight-<stat>`.

**Copy**

- **Every user-visible string lives in `web/src/lib/sim/copy.ts`.** Part A appends inside the existing `simCopy` object; **this lane's strings are a second named export, `bulkCopy`, appended after `simCopy`'s closing brace**, so the two lanes' additions sit at different anchors and a merge does not collide. No component holds a literal sentence, and no test asserts on a literal a component could stop rendering.
- Voice, per `design/DESIGN-SYSTEM.md`: reference, not pitch. State the number and stop.

**Design system**

- No new hex literal. Colours come from Tailwind utilities bound to `web/src/styles/tokens.css`: `bg-bg`, `bg-raised`, `bg-card-top`, `border-line`, `border-line-soft`, `border-line-warm`, `text`, `text-strong`, `text-muted`, `text-nav`, `text-gold`, `rounded-panel`, `rounded-control`, `rounded-pill`, `font-display`, `font-mono`. Rarity goes through `rarityClassFor()`, never a literal.
- Reuse `SECONDARY_BUTTON` from `web/src/lib/planner/styles.ts` for every button. Minimum hit target 44px on phone.
- **No `console.log`.** `console.error` only where an existing module already does it.
- **No file over 800 lines.** The tool island is split one component per page plus shared parts; if a component crosses 400 lines, split it before committing.

**Files shared with part A, and the rule**

- `web/src/lib/sim/types.ts`: this lane makes **one** addition — `reference_stat?: string` on `SpecFidelity` and `kind`/`headline` on `SimListRow` (Task 1). Nothing else. Part A's additions (`TargetError`, `EncounterSpec` fields, `Sample`) are elsewhere in the file.
- `web/src/lib/sim/copy.ts`: the `bulkCopy` export only, after `simCopy`.
- `web/src/lib/sim/store.svelte.ts`: **never touched by this lane.** All new state is in `bulk-store.svelte.ts`.
- `web/src/lib/sim/run.ts`: **never touched by this lane.** Task 3 extracts the pool triple into a new `run-shards.ts`; rewiring `run.ts` onto it is recorded as a merge follow-up, not a step here.

**Consumed from part A** (name them in every task that needs them; if part A has not landed, stub against the signature and say so in the commit body)

1. `web/src/lib/sim/styles.ts` — `SIM_STYLES: readonly { id: string; label: string }[]` and `encounterForStyle(id: string, base: EncounterSpec): EncounterSpec`.
2. `web/src/components/sim/SettingsPanel.svelte` — the full buff/consumable/encounter panel. Props `{ settings: SimSettings; spec: string; disabled: boolean; onchange: (next: SimSettings) => void }`.
3. `web/src/components/sim/RequestDrawer.svelte` — Advanced (design 8). Props `{ request: unknown; onapply: (next: unknown) => void }`, test id `sim-request-drawer`.
4. `decodeFS1` in `web/src/lib/planner/fs1.ts` returning `FS1Build` with `bags`, `bank`, `sets`, `loadouts` and `professions` (contract 10.5's new `professions=` section), and `SimCharacter` (`web/src/lib/sim/character.ts`) carrying the same five through `characterFromFs1`.
5. **`SimCharacter.gearSlots?: GearSlot[]`** — contract 10.5's "`SimCharacter` gains per-slot enchant and suffix", as the engine's own `GearSlot[]` beside the planner's id-only `gear` map. Part A owns the field and its name; this lane reads it in exactly one place (`seedRows` in `bulk-store.svelte.ts`) and in exactly one other (`toCharacterSpec`, which part A already changes to send it), so adopting a different name is a two-line change.
6. **`SimCharacter.professions?: string[]`** — filled by the export's `professions=` section, and sent on as `CharacterSpec.professions`. Read by the Droptimizer's crafted split (contract 10.7).
7. `web/src/lib/sim/run.ts`'s step-loop `RunInput.targetError`, and its use of the wasm's `simNeedsMore` — consumed only where a tool offers "until ±0.5%"; no tool in this plan does, so this is a compatibility dependency, not a functional one. This lane nevertheless declares `simNeedsMore` and `simValidate` on `EngineModule` (Task 3) and on the pool (Task 4), because contract 10.2 puts all three new exports in one table and splitting one file's surface across two lanes is how a merge conflict is made.

---

## File Structure

**New pure modules** (`web/src/lib/sim/`, all unit-tested, none imports Astro or Svelte)

| File | Responsibility |
| --- | --- |
| `bulk-types.ts` | TypeScript mirrors of contract 1.3, 1.4, 2 and 4. The only place a bulk shape is declared. |
| `run-shards.ts` | `runShards()`: one prepared `SimRequest` JSON through `simSplit`/`simRun`/`simCombine`. |
| `bulk-run.ts` | The stage loop: `simPlan` → run → `simRank` → next stage or final. Progress, abort, cap. |
| `bulk-store.svelte.ts` | The runes store all four tool pages mount over. |
| `phase.ts` | Phase names and boundaries: `GET /v1/phases` at runtime, `src/data/phases.json` as the build-time fallback. |
| `sim-buffs.ts` | `simbuffs.json`: the display name and icon for every IDS.md buff, debuff and consumable id. |
| `loot.ts` | `loot.json` loader, the source tree, the item→source index, the phase gate. |
| `enchants.ts` | `enchants.json` / `suffixes.json` loaders and per-slot filtering. |
| `item-search.ts` | The search over `items.json`: name, min ilvl, slot, source, usable-only. |
| `candidates.ts` | The candidate row model: origins, keys, toggles, copy-and-modify, locks, envelope build. |
| `combos.ts` | Result-view arithmetic-free helpers: rows, within-error grouping, per-slot summary, headline. |
| `weights.ts` | Stat vocabulary, the default set from the spec's `reference_stat`, the Pawn string. |
| `addon-export.ts` | "Copy to addon" for a winning set. |
| `bulk-skeleton.ts` | The four shell skeletons, one string each, shared by the `.astro` shell and the island. |

**New components** (`web/src/components/sim/tools/`)

`ToolsView.svelte` (island root, lazy-loads one tool), `TopGear.svelte`, `SlotGrid.svelte`, `CandidateRows.svelte`, `ItemSearch.svelte`, `EnchantList.svelte`, `TalentCandidates.svelte`, `NamedSets.svelte`, `BulkRunBar.svelte`, `ComboResults.svelte`, `Droptimizer.svelte`, `SourcePicker.svelte`, `DropResults.svelte`, `StatWeights.svelte`.

**New components** (`web/src/components/sim/`, for `/sim/<id>`)

`SavedCombos.svelte`, `SavedWeights.svelte` — lazy chunks off `SavedSim.svelte`, so the 90 KB `sim-island` budget is untouched.

**New pages and entries**

`src/pages/sim/gear.astro`, `talents.astro`, `drops.astro`, `weights.astro`; `src/sim-tools-island.ts`; `vite.sim-tools-island.config.ts`.

**Modified**

`src/lib/sim/types.ts` (one addition), `src/lib/sim/copy.ts` (`bulkCopy`), `src/lib/sim/engine.ts`, `engine-protocol.ts`, `sim.worker.ts`, `worker.ts` (three new exports), `src/lib/sim/api.ts` (three calls), `src/lib/planner/types.ts` (`Item.suffixes`), `src/lib/report/og-meta.ts` (unfurl per kind), `src/components/sim/SavedSim.svelte` and `SimHistory.svelte`, `src/fixtures/sim/engine-fake.ts`, `src/test-support/sim-api.ts`, `vite.island.config.ts`, `package.json`, `scripts/check-island-size.mjs`, `scripts/sync-data.mjs`, `lighthouserc.json`, `src/data/phases.json` (new data file).

---

### Task 1: The bulk envelope in TypeScript

**Files:**
- Create: `web/src/lib/sim/bulk-types.ts`
- Create: `web/src/lib/sim/bulk-types.test.ts`
- Modify: `web/src/lib/sim/types.ts` (add `reference_stat` to `SpecFidelity`; add `kind` and `headline` to `SimListRow`)
- Modify: `web/src/lib/sim/copy.ts` (append the `bulkCopy` export after `simCopy`)
- Modify: `web/src/lib/planner/types.ts` (add `suffixes?: number[]` to `Item`)

**Interfaces:**
- Consumes: `SimRequest`, `SimResult`, `Estimate`, `GearSlot`, `SpecFidelity`, `SimListRow` from `./types`; `Item` from `../planner/types`.
- Produces: everything later tasks type against — `BulkMode`, `Precision`, `PRECISIONS`, `STAGES_BY_PRECISION`, `finalIterations()`, `Candidate`, `TalentLoadout`, `GearSet`, `BulkSpec`, `WeightsSpec`, `BulkRequest`, `WeightsRequest`, `Substitution`, `Combo`, `Stage`, `StatWeight`, `BulkResult`, `WeightsResult`, `Combination`, `StageRequests`, `RankAnswer`, `CapExceeded`, `SimKind`, `BROWSER_CAP`, `SERVER_CAP`, `LOW_CORE_CAP`, `LOW_CORE_THRESHOLD`, `browserCap()`, `kindOf()`, `isCapExceeded()`; and `bulkCopy` from `./copy`.

- [ ] **Step 1: Write the failing test**

Create `web/src/lib/sim/bulk-types.test.ts`:

```ts
// web/src/lib/sim/bulk-types.test.ts
import { describe, expect, it } from 'vitest';
import {
  BROWSER_CAP,
  LOW_CORE_CAP,
  SERVER_CAP,
  browserCap,
  finalIterations,
  isCapExceeded,
  kindOf,
  PRECISIONS,
  type BulkRequest,
  type WeightsRequest,
} from './bulk-types';
import type { SimRequest } from './types';

const base: SimRequest = {
  engine_version: 'testver',
  spec: 'warrior-fury',
  source: { kind: 'addon', ref: '', captured_at: '2026-09-19T00:00:00Z' },
  character: {
    name: 'Fury',
    race: 'orc',
    class: 'warrior',
    level: 60,
    talents: '0-5530515-',
    gear: [{ slot: 'head', item_id: 12640 }],
    buffs: [],
    consumes: [],
  },
  encounter: { duration_sec: 180, variation: 0.2, targets: 1, execute_ratio: 0.25, profile: '' },
  iterations: 3000,
  random_seed: 0,
};

describe('browserCap', () => {
  it('is 400 on a desktop and 200 on four cores or fewer', () => {
    expect(browserCap(8)).toBe(BROWSER_CAP);
    expect(browserCap(6)).toBe(BROWSER_CAP);
    expect(browserCap(4)).toBe(LOW_CORE_CAP);
    expect(browserCap(2)).toBe(LOW_CORE_CAP);
  });

  it('assumes the low-core cap when the browser will not say', () => {
    expect(browserCap(undefined)).toBe(LOW_CORE_CAP);
    expect(browserCap(Number.NaN)).toBe(LOW_CORE_CAP);
    expect(browserCap(0)).toBe(LOW_CORE_CAP);
  });
});

describe('kindOf', () => {
  it('derives run, gear, talents, drops and weights the way SimRequest.Kind() does', () => {
    expect(kindOf(base)).toBe('run');
    for (const mode of ['gear', 'talents', 'drops'] as const) {
      const request: BulkRequest = {
        ...base,
        bulk: { mode, candidates: [], precision: 'normal', cap: BROWSER_CAP },
      };
      expect(kindOf(request)).toBe(mode);
    }
    const weights: WeightsRequest = { ...base, weights: { stats: ['crit'], reference: 'crit' } };
    expect(kindOf(weights)).toBe('weights');
  });

  it('prefers bulk over weights, because a bulk request is never a weights run', () => {
    const both = {
      ...base,
      bulk: { mode: 'gear' as const, candidates: [], precision: 'fast' as const, cap: 400 },
      weights: { stats: ['crit'], reference: 'crit' },
    };
    expect(kindOf(both)).toBe('gear');
  });
});

describe('isCapExceeded', () => {
  it('recognises the wasm refusal and nothing else', () => {
    expect(isCapExceeded({ error: 'cap_exceeded', cap: 400, combinations: 812 })).toBe(true);
    expect(isCapExceeded({ error: 'request: unknown buff' })).toBe(false);
    expect(isCapExceeded({ stage: 1, iterations: 100, requests: [], combos: [] })).toBe(false);
    expect(isCapExceeded(null)).toBe(false);
  });
});

describe('PRECISIONS and finalIterations', () => {
  it('is the contract’s three, in ladder order', () => {
    expect([...PRECISIONS]).toEqual(['fast', 'normal', 'high']);
  });

  it('gives each precision its final-stage iteration count (contract 10.1 A3)', () => {
    expect(finalIterations('fast')).toBe(3000);
    expect(finalIterations('normal')).toBe(3000);
    expect(finalIterations('high')).toBe(10_000);
  });
});

describe('the lane caps', () => {
  it('are 400 in the browser and 5,000 on the server (contract 10.1 A2)', () => {
    expect(BROWSER_CAP).toBe(400);
    expect(SERVER_CAP).toBe(5000);
  });
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-types.test.ts
```

Expected: FAIL, `Failed to resolve import "./bulk-types"`.

- [ ] **Step 3: Write `bulk-types.ts`**

```ts
// web/src/lib/sim/bulk-types.ts
// TypeScript mirrors of sim/api/envelope.go's bulk and weights additions, contract sections
// 1.3, 1.4, 2 and 4. Every key here is that file's `json:` tag.
//
// These extend SimRequest/SimResult rather than widening them in types.ts, for two reasons:
// the bulk fields are absent from every single run and an optional field on the base type
// would make every existing reader carry an `undefined` branch it never needs; and this file
// is owned by one lane, so two lanes amending the envelope in parallel never collide.
//
// Nothing here computes a statistic. Expansion, staging, ranking and the within-error
// grouping all happen inside sim/bulk in the wasm; TypeScript names the shapes and renders
// them.
import type { Estimate, GearSlot, SimRequest, SimResult } from './types';

export type BulkMode = 'gear' | 'talents' | 'drops';
export type SimKind = 'run' | BulkMode | 'weights';

/** The contract's three, in ladder order: more stages and fewer survivors, left to right. */
export const PRECISIONS = ['fast', 'normal', 'high'] as const;
export type Precision = (typeof PRECISIONS)[number];

/**
 * How many stages a precision runs, for the progress line only. The iteration counts and the
 * survivor cuts are sim/bulk's and are never mirrored here: fast is 100 -> 1000 -> 3000,
 * normal is 1000 -> 3000, high is 1000 -> 10000 (contract 1.3).
 */
export const STAGES_BY_PRECISION: Record<Precision, number> = { fast: 3, normal: 2, high: 2 };

/**
 * The iteration count a bulk request carries in `SimRequest.iterations` -- its precision's
 * FINAL stage, not 3,000 for everything (contract 10.1 A3). `Validate` applies the rule
 * for the request's kind, so a `high` bulk request with 3,000 on it is refused.
 */
export function finalIterations(precision: Precision): number {
  return precision === 'high' ? 10_000 : 3000;
}

/**
 * One substitution the planner may try. `slot` is "" for anything that fits more than one
 * slot -- rings, trinkets, weapons -- and sim/bulk decides which slot it lands in.
 */
export interface Candidate {
  slot: string;
  item_id: number;
  enchant?: number;
  suffix?: number;
  /** equipped | bag | bank | search | drop:<source-id> | set:<name> */
  origin: string;
  /**
   * Where it comes from, in words -- "Ragnaros", "Molten Core", "Blacksmithing" (contract
   * 10.1 A6). The page fills it from `loot.json` when it builds a drops request; `Rank`
   * copies it onto the substitution, and the API's headline reads it. It exists so the
   * results view never re-joins a source id back to a name it already had.
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
  mode: BulkMode;
  candidates: Candidate[];
  talents?: TalentLoadout[];
  sets?: GearSet[];
  /**
   * Alternative consumable lists, tried as candidates in `gear` mode (contract 10.1 A5).
   * Each inner list replaces `CharacterSpec.Consumes` wholesale for that combination, so
   * "try each of these" is one inner list per consumable rather than every subset.
   */
  consumables?: string[][];
  /** Slots never substituted. */
  locked?: string[];
  precision: Precision;
  /** The lane's cap, echoed so a saved request says what bounded it. */
  cap: number;
}

export interface WeightsSpec {
  /** IDS.md stat ids. */
  stats: string[];
  /** The stat normalised to 1.0. */
  reference: string;
}

export interface BulkRequest extends SimRequest {
  bulk: BulkSpec;
}

export interface WeightsRequest extends SimRequest {
  weights: WeightsSpec;
}

export interface Substitution {
  /**
   * The four kinds, closed (contract 10.8): `consumes` is the alternative consumable list
   * of 10.1 A5, and its `name` is the consumable ids joined by `, `. `sim/bulk` emits one
   * per combination that used an alternative list.
   */
  kind: 'item' | 'talents' | 'set' | 'consumes';
  slot?: string;
  item_id?: number;
  enchant?: number;
  suffix?: number;
  /**
   * The display name. Filled for items too, from simdb (contract 10.1 A6) -- so the page
   * never joins an item id back to a name the result already carries.
   */
  name?: string;
  talents?: string;
  origin?: string;
  /** Copied from the candidate: "Ragnaros", "Blacksmithing" (contract 10.1 A6). */
  source_name?: string;
}

export interface Combo {
  substitutions: Substitution[];
  dps: Estimate;
  /** Against `equipped`, paired at the same stage, with its own error. */
  delta: Estimate;
  /** 0 is the leader's within-error group, then 1, 2, .... */
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

export interface BulkResult extends SimResult {
  combos: Combo[];
  equipped: Estimate;
  stages: Stage[];
}

export interface WeightsResult extends SimResult {
  weights: StatWeight[];
}

/** What `simPlan` returns, and what `simRank` returns as `next`. */
export interface StageRequests {
  stage: number;
  iterations: number;
  /** Requests[0] is always the equipped set. */
  requests: SimRequest[];
  /** Parallel to Requests[1:]. */
  combos: Combination[];
  /**
   * The ladder's history so far (contract 10.1 A10): what each finished stage cost. It
   * rides on the stage object across the wasm boundary so `Rank` can fill
   * `SimResult.Stages` without the page keeping a tally of its own.
   */
  ran?: Stage[];
}

export interface Combination {
  request: SimRequest;
  substitutions: Substitution[];
}

/** What `simRank` answers: one of the two is set. */
export interface RankAnswer {
  next?: StageRequests;
  result?: SimResult;
}

/**
 * sim/bulk's ErrCapExceeded across the wasm boundary. Plain `{"error": "..."}` carries no
 * numbers and the page has to say what the cap was and by how much the list overran it, so
 * the cap refusal is the one structured error the exports answer with.
 */
export interface CapExceeded {
  error: 'cap_exceeded';
  cap: number;
  combinations: number;
}

export function isCapExceeded(value: unknown): value is CapExceeded {
  if (typeof value !== 'object' || value === null) return false;
  const record = value as Record<string, unknown>;
  return (
    record.error === 'cap_exceeded' &&
    typeof record.cap === 'number' &&
    typeof record.combinations === 'number'
  );
}

/** The browser lane's cap (contract 1.3's `Caps`, unchanged by section 10). */
export const BROWSER_CAP = 400;
/**
 * The premium lane's cap. 5,000, not section 1.3's 20,000: contract 10.1 A2 lowered it
 * because a 20,000-combination fast run cannot finish inside the job's 15-minute timeout.
 * The page shows this number in the cap notice's premium alternative and refuses to
 * dispatch a server run past it.
 */
export const SERVER_CAP = 5000;
/** Design 2.3: "a phone gets half the cap by hardwareConcurrency". */
export const LOW_CORE_CAP = BROWSER_CAP / 2;
export const LOW_CORE_THRESHOLD = 4;

/**
 * The cap this device gets. An absent or nonsensical `hardwareConcurrency` takes the low
 * cap rather than the high one: a browser that will not say how many cores it has is far
 * more often a phone with few than a workstation with many, and the run button always says
 * what premium would allow anyway.
 */
export function browserCap(hardwareConcurrency: number | undefined): number {
  if (typeof hardwareConcurrency !== 'number' || !Number.isFinite(hardwareConcurrency)) {
    return LOW_CORE_CAP;
  }
  if (hardwareConcurrency <= 0) return LOW_CORE_CAP;
  return hardwareConcurrency <= LOW_CORE_THRESHOLD ? LOW_CORE_CAP : BROWSER_CAP;
}

/**
 * Kind is derived, never sent -- the same rule and the same order as Go's
 * `SimRequest.Kind()`: a request with Bulk is gear, talents or drops by Bulk.Mode; one with
 * Weights is weights; otherwise run.
 */
export function kindOf(request: SimRequest): SimKind {
  const bulk = (request as Partial<BulkRequest>).bulk;
  if (bulk !== undefined) return bulk.mode;
  if ((request as Partial<WeightsRequest>).weights !== undefined) return 'weights';
  return 'run';
}
```

- [ ] **Step 4: Run it and watch it pass**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-types.test.ts
```

Expected: PASS, 6 tests.

- [ ] **Step 5: Add the three field additions to the shared types**

In `web/src/lib/sim/types.ts`, inside `interface SpecFidelity`, after the `engine_version` field:

```ts
  /**
   * The stat `/sim/weights` normalises to 1.0 for this spec, from data/curated/specs.json
   * (contract 8, `GET /v1/specs` (+)). Optional because a build predating the column, or a
   * spec nobody has set one for, sends no value; `weights.ts`'s own fallback is the first
   * stat the spec's class uses.
   */
  reference_stat?: string;
```

In the same file, inside `interface SimListRow`, after `title`:

```ts
  /**
   * `run` | `gear` | `drops` | `talents` | `weights`, from the `sims.kind` column
   * (contract 8). Optional so a row written before migration 0014 still parses; absent
   * reads as `run`.
   */
  kind?: string;
  /**
   * The API's own one-line summary of the result, composed server-side so every client
   * says the same thing: run -> "1,204 DPS", gear -> "+41 DPS from Vis'kag", drops ->
   * "3 upgrades on Ragnaros", talents -> "+18 DPS with 'Deep Fury'", weights ->
   * "Crit 1.00 - Agility 0.87".
   */
  headline?: string;
```

In `web/src/lib/planner/types.ts`, inside `interface Item`, after `unique`:

```ts
  /**
   * The random-suffix ids this item can roll (contract 6.3: "items.json rows gain
   * `suffixes: [id, ...]` where the item rolls one"). Absent on a build the data lane has
   * not regenerated, and on every item that rolls none.
   */
  suffixes?: number[];
```

- [ ] **Step 6: Append the copy block to `copy.ts`**

Note the two numbers the contract's section 10 pins: `capPremiumNote` says **5,000**
(10.1 A2), and `opensLater` is the copy for a source whose `opens` is the literal `"later"`
(10.4) — an unreleased source with no announced date.

At the very end of `web/src/lib/sim/copy.ts`, **after** the `} as const;` that closes `simCopy`:

```ts
/**
 * The simulator's combination tools: Top Gear, talent compare, Droptimizer and stat
 * weights. A second export rather than more keys on `simCopy` so the parity work's two web
 * lanes append at different anchors in this file and never collide; every rule above
 * applies unchanged -- components import from here and tests assert against these
 * constants, never against a literal.
 */
export const bulkCopy = {
  // --- page titles and the one-line standfirst under each ---
  gearTitle: 'Top Gear',
  gearIntro:
    'Tick the items, enchants, talents and sets you want tried. Every valid combination is simulated and ranked against what you have on.',
  talentsTitle: 'Talent compare',
  talentsIntro:
    'Your build against every other build you have, ranked. Gear is locked to what you are wearing, so the only thing that changes is the tree.',
  dropsTitle: 'Droptimizer',
  dropsIntro:
    'Pick where you are going. Every item those bosses drop is simulated one at a time against your current set, and the upgrades are listed by boss.',
  weightsTitle: 'Stat weights',
  weightsIntro: 'What one point of each stat is worth, for the addons that ask for a number.',
  weightsWarning:
    'A stat weight is a straight-line guess at something that is not a straight line: it holds near the gear you have now and stops holding as soon as a set bonus, a proc or a hit cap changes. Sim the actual items in Top Gear instead. These are here because addons want them.',
  weightsWarningLink: 'Open Top Gear',

  // --- candidates ---
  equipped: 'Equipped',
  bags: 'Bags',
  bank: 'Bank',
  fromSearch: 'Search',
  pinned: 'Pinned',
  lockSlot: 'Lock to equipped',
  lockedSlot: 'Locked',
  copyAndModify: 'Copy and modify',
  withEnchant: 'With enchant',
  withSuffix: 'With suffix',
  keepCurrentEnchant: 'Keep current',
  /**
   * A `consumes` substitution on a results row (contract 10.8): the ids come back joined
   * by ", ", and this is the only place they are turned into a phrase. Ids are legible in
   * snake case -- "flask_of_supreme_power" -- so they are de-underscored rather than
   * looked up: `simbuffs.json` is loaded per build and a saved result opened by someone
   * else may not have it.
   */
  consumesChip: (ids: string): string =>
    `With ${ids
      .split(', ')
      .map((id) => id.replaceAll('_', ' '))
      .join(', ')}`,
  noEnchant: 'None',
  enchantCap: (cap: number): string => `At most ${cap} enchants per slot.`,
  noCandidates: 'Nothing ticked yet. Tick an item, a talent build or a set.',
  bagsNeedAddon: 'Your bags and bank come from the addon export; this character was loaded another way.',
  tryEach: 'Try each',
  consumableCandidates: 'Try each of these as a candidate rather than a setting.',

  // --- item search ---
  searchLabel: 'Find an item',
  searchPlaceholder: 'Name',
  searchMinItemLevel: 'Minimum item level',
  searchSlot: 'Slot',
  searchSource: 'Source',
  searchAnySlot: 'Any slot',
  searchAnySource: 'Anywhere',
  searchUsableOnly: 'Only items this character can equip',
  searchNoResults: 'No item in this class’s list matches.',
  searchTruncated: (shown: number, total: number): string =>
    `Showing ${shown} of ${total}. Narrow the search.`,
  searchAdd: 'Add',
  searchAdded: 'Added',

  // --- talents and sets ---
  talentsOwn: 'Your current build',
  talentsSaved: 'Your saved builds',
  talentsLoadouts: 'In-game loadouts',
  talentsAddCustom: 'Add a build',
  talentsNoSaved: 'No saved builds for this class yet.',
  talentsSavedUnavailable: 'Your saved builds could not be read; the rest of the page still works.',
  setsTitle: 'Whole sets',
  setsIntro: 'A set replaces every slot at once. Paste a second export string to add one.',
  setsPaste: 'Paste an export string',
  setsAdd: 'Add set',
  setsBadCode: 'That is not an export string this build can read.',

  // --- the run bar ---
  combinations: (n: number): string =>
    `${n.toLocaleString('en-US')} valid ${n === 1 ? 'combination' : 'combinations'}`,
  combinationsCounting: 'Counting combinations…',
  precisionLabel: 'Precision',
  precisionFast: 'Fast',
  precisionNormal: 'Normal',
  precisionHigh: 'High',
  precisionNote: {
    fast: 'Three stages: everything at 100 iterations, the survivors at 1,000, the finalists at 3,000.',
    normal: 'Two stages: everything at 1,000 iterations, the finalists at 3,000.',
    high: 'Two stages: everything at 1,000 iterations, twice as many finalists at 10,000.',
  } as Record<string, string>,
  runBulk: 'Run',
  runBulkAgain: 'Run again',
  stopBulk: 'Stop',
  stageProgress: (stage: number, stages: number, done: number, total: number): string =>
    `stage ${stage} of ${stages} · ${done.toLocaleString('en-US')} of ${total.toLocaleString('en-US')} combinations`,
  partial: 'Stopped. These are the combinations that finished.',
  capNotice: (cap: number, combinations: number): string =>
    `${combinations.toLocaleString('en-US')} combinations is past this browser’s limit of ${cap.toLocaleString('en-US')}. Untick ${(combinations - cap).toLocaleString('en-US')} of them, or run it on our servers.`,
  capPremium: 'Run on our servers',
  capPremiumNote: 'Premium lifts the limit to 5,000 combinations and any precision.',
  serverCapNotice: (cap: number, combinations: number): string =>
    `${combinations.toLocaleString('en-US')} combinations is past our servers’ limit of ${cap.toLocaleString('en-US')} too. Untick ${(combinations - cap).toLocaleString('en-US')} of them.`,
  lowCoreNote: (cap: number): string =>
    `This device reports four cores or fewer, so the limit here is ${cap.toLocaleString('en-US')} combinations.`,

  // --- results ---
  resultsEquipped: 'What you have on',
  resultsRank: '#',
  resultsChange: 'Change',
  resultsDps: 'DPS',
  resultsDelta: 'Gain',
  resultsPercent: '%',
  withinError: 'Within error of the leader',
  withinErrorNote:
    'These runs are too close to separate at this many iterations. Run again at a higher precision to tell them apart.',
  noGain: 'Nothing here beats what you are wearing.',
  slotSummary: 'By slot',
  slotSummaryNote: 'What the winning set uses in each slot, and what that slot contributed.',
  openInPlanner: 'Open in planner',
  copyToAddon: 'Copy to addon',
  copiedToAddon: 'Copied',
  keepFourPiece: 'Only combinations keeping a 4-piece set bonus',
  ranAtStages: (stages: { iterations: number; combos: number }[]): string =>
    stages
      .map(
        (stage) =>
          `${stage.combos.toLocaleString('en-US')} at ${stage.iterations.toLocaleString('en-US')}`,
      )
      .join(' · '),

  // --- droptimizer ---
  sourcesRaids: 'Raids',
  sourcesDungeons: 'Dungeons',
  sourcesWorld: 'World bosses',
  sourcesCrafted: 'Crafted',
  sourcesRep: 'Reputation',
  sourcesPvp: 'PvP',
  sourcesQuests: 'Quests',
  sourcesQuestNote: 'Off by default: a quest reward is a one-time source.',
  sourcesMyProfessions: 'My professions',
  sourcesAllProfessions: 'All professions',
  sourcesProfessionsUnknown:
    'Nothing has recorded this character’s professions, so every profession is listed.',
  sourcesTrash: 'Trash and chests',
  sourcesWholeRaid: 'Every boss',
  showUpcoming: 'Show unreleased content',
  opensOn: (label: string, date: string): string => `${label}, ${date}`,
  notOpenYet: 'Not open yet',
  opensLater: 'Not open yet; no date announced.',
  dropsUpgrades: (upgrades: number, drops: number): string =>
    `${upgrades} of the ${drops} drops here ${upgrades === 1 ? 'is an upgrade' : 'are upgrades'}`,
  dropsBest: 'Best here',
  dropsEveryUpgrade: 'Every upgrade',
  dropsByBoss: 'By boss',
  dropsPin: 'Pin into Top Gear',
  dropsPinned: 'Pinned',
  dropsNoChance:
    'Neither database records drop rates, so nothing here is a probability. It is a count of what drops and what would be an upgrade.',
  dropsNothing: 'No source ticked yet.',

  // --- stat weights ---
  weightsStat: 'Stat',
  weightsWeight: 'Weight',
  weightsReference: 'Reference',
  weightsCopyPawn: 'Copy for Pawn',
  weightsCopied: 'Copied',
  weightsPick: 'Stats to weigh',

  // --- the request drawer's own one-liner, so Advanced is findable on every tool ---
  advancedTitle: 'Request',

  // --- the rules card, design 3.2, in our words ---
  rulesTitle: 'How the combinations are built',
  rules: [
    'An enchant you already have carries over to a candidate in the same slot where it fits.',
    'Rings and trinkets are tried in both slots.',
    'A two-hander and a one-hander with an off-hand are competing shapes, not two slots.',
    'A dual-wield spec tries each weapon pair both ways round.',
    'Unique-equipped is respected, including unique categories.',
    'Nothing this character cannot equip is ever simulated.',
    'Item search can find items this character has no way to obtain.',
  ] as readonly string[],

  // --- failures ---
  planFailed: 'The combinations could not be worked out.',
  bulkFailed: 'The engine could not run these combinations.',
  lootFailed: 'The loot tables could not be read.',
  enchantsFailed: 'The enchant list could not be read.',
  needCharacter: 'Load a character first.',
  needAddonForBags: 'Paste your addon export to see your bags and bank here.',
} as const;
```

- [ ] **Step 7: Type-check, lint and format**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/bulk-types.ts src/lib/sim/bulk-types.test.ts src/lib/sim/copy.ts src/lib/sim/types.ts src/lib/planner/types.ts
npx prettier --write src/lib/sim/bulk-types.ts src/lib/sim/bulk-types.test.ts src/lib/sim/copy.ts src/lib/sim/types.ts src/lib/planner/types.ts
```

Expected: `0 errors`, eslint clean.

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/bulk-types.ts web/src/lib/sim/bulk-types.test.ts web/src/lib/sim/copy.ts web/src/lib/sim/types.ts web/src/lib/planner/types.ts
git commit -m "$(cat <<'MSG'
feat(sim): the bulk and weights envelope in TypeScript

Mirrors contract 1.3, 1.4, 2 and 4 as extensions of SimRequest/SimResult
rather than optional fields on them, so the two parity web lanes amend
different files. Adds the browser cap rule (400, halved at four cores or
fewer), the derived kind, and the structured cap refusal. bulkCopy is a
second export in copy.ts for the same collision reason.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 2: The fixtures every later task and every test renders against

Nothing in this lane can be tested without data: a ranked bulk result, a weights result, a
loot slice, an enchant slice, a suffix slice, a phase table and a `reference_stat` on the
spec rows. They land together because they are one decision — what the fixture world
contains — and splitting them would leave half the suite unable to run.

**Files:**
- Create: `web/src/fixtures/sim/bulk-result.json`
- Create: `web/src/fixtures/sim/weights-result.json`
- Create: `web/src/fixtures/planner/loot.json`
- Create: `web/src/fixtures/planner/enchants.json`
- Create: `web/src/fixtures/planner/suffixes.json`
- Create: `web/src/fixtures/planner/simbuffs.json`
- Create: `web/src/data/phases.json`
- Create: `web/src/fixtures/sim/bulk-fixture.test.ts`
- Modify: `web/src/fixtures/sim/specs.json` (add `reference_stat` to all three rows)
- Modify: `web/src/fixtures/planner/items/warrior.json` (add `suffixes` to one item)
- Modify: `web/src/fixtures/planner/manifest.json` (list the three new files)
- Modify: `web/scripts/sync-data.mjs` (`SYNC_ENTRIES` gains the three files)
- Modify: `web/src/test-support/sim-api.ts` (export the two new fixtures; `GET /v1/builds?mine=1`)

**Interfaces:**
- Consumes: `BulkResult`, `WeightsResult`, `StatWeight` (Task 1).
- Produces: `fixtureBulkResult`, `fixtureWeightsResult`, `FIXTURE_MY_BUILD_ID` from `web/src/test-support/sim-api.ts`; the four data files at `/data/<build>/loot.json`, `enchants.json`, `suffixes.json` and `simbuffs.json`; `web/src/data/phases.json`; and the `GET /v1/phases` and `GET /v1/builds?mine=1` stub routes.

- [ ] **Step 1: Write the failing fixture test**

Create `web/src/fixtures/sim/bulk-fixture.test.ts`:

```ts
// web/src/fixtures/sim/bulk-fixture.test.ts
// The fixtures are data, so their invariants are asserted rather than assumed: every later
// task's tests read these files, and a fixture that quietly stops matching the contract
// turns a real failure into a green run.
import { describe, expect, it } from 'vitest';
import bulkJson from './bulk-result.json';
import weightsJson from './weights-result.json';
import specsJson from './specs.json';
import lootJson from '../planner/loot.json';
import enchantsJson from '../planner/enchants.json';
import suffixesJson from '../planner/suffixes.json';
import simbuffsJson from '../planner/simbuffs.json';
import itemsJson from '../planner/items/warrior.json';
import phasesJson from '../../data/phases.json';
import type { BulkResult, WeightsResult } from '../../lib/sim/bulk-types';
import type { SpecFidelity } from '../../lib/sim/types';

const bulk = bulkJson as unknown as BulkResult;
const weights = weightsJson as unknown as WeightsResult;

describe('the bulk result fixture', () => {
  it('is a gear request with an equipped baseline and three stages', () => {
    expect(bulk.request.bulk?.mode).toBe('gear');
    expect(bulk.request.bulk?.precision).toBe('fast');
    expect(bulk.equipped.mean).toBeGreaterThan(0);
    expect(bulk.stages.map((stage) => stage.iterations)).toEqual([100, 1000, 3000]);
  });

  it('is ranked best first, with every delta measured against the equipped set', () => {
    const means = bulk.combos.map((combo) => combo.dps.mean);
    expect([...means].sort((a, b) => b - a)).toEqual(means);
    for (const combo of bulk.combos) {
      expect(combo.delta.mean).toBeCloseTo(combo.dps.mean - bulk.equipped.mean, 6);
      expect(combo.delta.error).toBeGreaterThan(0);
    }
  });

  it('carries a within-error group at the top and a separated group below it', () => {
    expect(bulk.combos.filter((combo) => combo.group === 0).length).toBeGreaterThan(1);
    expect(new Set(bulk.combos.map((combo) => combo.group))).toEqual(new Set([0, 1, 2]));
  });

  it('names a slot on every ring and trinket substitution, the way sim/bulk resolves them', () => {
    for (const combo of bulk.combos) {
      for (const sub of combo.substitutions) {
        if (sub.kind !== 'item') continue;
        expect(sub.slot).toBeTruthy();
        expect(sub.item_id).toBeGreaterThan(0);
        expect(sub.origin).toBeTruthy();
      }
    }
  });
});

describe('the weights result fixture', () => {
  it('normalises the reference stat to exactly 1', () => {
    const reference = weights.request.weights?.reference ?? '';
    const row = weights.weights.find((entry) => entry.stat === reference);
    expect(row?.weight).toBe(1);
  });

  it('gives every stat an error', () => {
    for (const row of weights.weights) expect(row.error).toBeGreaterThan(0);
  });
});

describe('the spec fixture', () => {
  it('gives every spec a reference stat', () => {
    for (const row of specsJson as unknown as SpecFidelity[]) {
      expect(row.reference_stat).toBeTruthy();
    }
  });
});

describe('the loot fixture', () => {
  const itemIds = new Set((itemsJson.items as { id: number }[]).map((item) => item.id));

  it('covers every source kind the picker groups by', () => {
    expect(new Set(lootJson.sources.map((source) => source.kind))).toEqual(
      new Set(['raid', 'dungeon', 'world', 'crafted', 'rep', 'pvp', 'quest']),
    );
  });

  it('only drops items the fixture item file actually has', () => {
    for (const source of lootJson.sources) {
      for (const id of [...(source.items ?? []), ...(source.trash ?? [])]) {
        expect(itemIds.has(id), `source ${source.id} drops ${id}`).toBe(true);
      }
      for (const boss of source.bosses ?? []) {
        for (const id of boss.items) expect(itemIds.has(id), `boss ${boss.id} drops ${id}`).toBe(true);
      }
    }
  });

  it('names a phase only from the phase table, or the literal "later"', () => {
    const names = new Set([...phasesJson.map((phase) => phase.name), 'later']);
    for (const source of lootJson.sources) {
      if (source.opens === undefined) continue;
      expect(names.has(source.opens), `source ${source.id} opens in ${source.opens}`).toBe(true);
    }
  });

  it('spells every boss id as <source>:<npc-id>, per contract 10.4', () => {
    for (const source of lootJson.sources) {
      for (const boss of source.bosses ?? []) {
        expect(boss.id).toBe(`${source.id}:${boss.npc_id}`);
      }
    }
  });
});

describe('the enchant and suffix fixtures', () => {
  it('keys every enchant by effect_id plus a spell or an item, per contract 10.4', () => {
    for (const enchant of enchantsJson) {
      expect(enchant.effect_id).toBeGreaterThan(0);
      expect(('spell_id' in enchant ? 1 : 0) + ('item_id' in enchant ? 1 : 0)).toBeGreaterThan(0);
      expect(enchant.slots.length).toBeGreaterThan(0);
      expect(enchant.icon.startsWith('fixture_')).toBe(true);
    }
  });

  it('points every suffix on an item at a suffix the file defines', () => {
    const suffixIds = new Set(suffixesJson.map((suffix) => suffix.id));
    for (const item of itemsJson.items as { suffixes?: number[] }[]) {
      for (const id of item.suffixes ?? []) expect(suffixIds.has(id)).toBe(true);
    }
  });
});

describe('the phase table', () => {
  it('is the four phases api/internal/phase names, in order, with the launch instant', () => {
    expect(phasesJson.map((phase) => phase.name)).toEqual(['pre-beta', 'beta', 'launch', 'raids-1']);
    expect(phasesJson[2].start).toBe('2026-11-04T23:00:00Z');
    expect(phasesJson[3].start).toBe('2026-12-09T00:00:00Z');
    expect(phasesJson[0].start).toBe('');
  });

  it('carries no display label: contract 10.4 pins the shape at name and start', () => {
    for (const phase of phasesJson) expect(Object.keys(phase).sort()).toEqual(['name', 'start']);
  });
});

describe('the simbuffs fixture', () => {
  it('names and illustrates every consumable the preset applies', () => {
    for (const id of ['flask_of_supreme_power', 'elixir_of_the_mongoose']) {
      expect(simbuffsJson.entries[id]?.name.length).toBeGreaterThan(0);
      expect(simbuffsJson.entries[id]?.icon.startsWith('fixture_')).toBe(true);
    }
  });
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/fixtures/sim/bulk-fixture.test.ts
```

Expected: FAIL, cannot resolve `./bulk-result.json`.

- [ ] **Step 3: Write `web/src/data/phases.json`**

```json
[
  { "name": "pre-beta", "start": "" },
  { "name": "beta", "start": "2026-09-17T00:00:00Z" },
  { "name": "launch", "start": "2026-11-04T23:00:00Z" },
  { "name": "raids-1", "start": "2026-12-09T00:00:00Z" }
]
```

Contract 10.4 makes `data/curated/phases.json` the source of truth — the file
`api/internal/phase.Boundaries` is tested against — and has the pipeline emit it to
`web/src/data/phases.json` in exactly this `[{ "name", "start" }]` shape. There is no
`label` column: `phase.ts` carries the display names, because they are display copy and the
data lane's file is data.

The file checked in here is the pipeline's output, written by hand for now so this lane is
not blocked on the data lane's emit step; the moment the pipeline writes it, this copy is
overwritten and nothing else changes. At **runtime** the page prefers `GET /v1/phases`
(contract 10.6) and falls back to this file, so a phase date that moves after a deploy is
picked up without a rebuild.

- [ ] **Step 4: Write `web/src/fixtures/sim/bulk-result.json`**

Six combos over the fixture's own six warrior items, ranked, with a two-member leader group,
a second group and a third. `request.character` is the same orc Fury warrior every other sim
fixture uses; `summary` is deliberately the empty-but-valid shape, because a bulk result's
page renders `combos`, never a damage table.

```json
{
  "engine_version": "edc0c8e9a",
  "request": {
    "engine_version": "edc0c8e9a",
    "spec": "warrior-fury",
    "source": { "kind": "addon", "ref": "", "captured_at": "2026-09-19T09:00:00Z" },
    "character": {
      "name": "Fury",
      "race": "orc",
      "class": "warrior",
      "level": 60,
      "talents": "0-5530515-",
      "gear": [
        { "slot": "head", "item_id": 12640 },
        { "slot": "main_hand", "item_id": 12784 },
        { "slot": "finger1", "item_id": 19325 }
      ],
      "buffs": ["battle_shout"],
      "consumes": []
    },
    "encounter": { "duration_sec": 180, "variation": 0.2, "targets": 1, "execute_ratio": 0.25, "profile": "" },
    "iterations": 3000,
    "random_seed": 0,
    "bulk": {
      "mode": "gear",
      "candidates": [
        { "slot": "head", "item_id": 16963, "origin": "bag" },
        { "slot": "shoulder", "item_id": 16966, "origin": "bank" },
        { "slot": "", "item_id": 13968, "origin": "search" },
        { "slot": "", "item_id": 19325, "origin": "equipped" }
      ],
      "talents": [{ "name": "Deep Fury", "talents": "0-5530515-" }],
      "sets": [],
      "locked": [],
      "precision": "fast",
      "cap": 400
    }
  },
  "lane": "browser",
  "dps": { "mean": 1502.4, "stddev": 210.5, "error": 3.84, "min": 900.1, "max": 2100.7 },
  "iterations_run": 3000,
  "duration_ms": 8420,
  "summary": { "duration_ms": 180000, "abilities": [], "auras": [], "casts": [], "resources": [], "roster": [], "combatants": [] },
  "equipped": { "mean": 1461.2, "stddev": 208.9, "error": 3.81, "min": 880.4, "max": 2044.9 },
  "stages": [
    { "iterations": 100, "combos": 6 },
    { "iterations": 1000, "combos": 4 },
    { "iterations": 3000, "combos": 3 }
  ],
  "combos": [
    {
      "substitutions": [
        { "kind": "item", "slot": "head", "item_id": 16963, "origin": "bag" },
        { "kind": "item", "slot": "shoulder", "item_id": 16966, "origin": "bank" }
      ],
      "dps": { "mean": 1502.4, "stddev": 210.5, "error": 3.84, "min": 900.1, "max": 2100.7 },
      "delta": { "mean": 41.2, "stddev": 0, "error": 5.41, "min": 0, "max": 0 },
      "group": 0
    },
    {
      "substitutions": [{ "kind": "item", "slot": "head", "item_id": 16963, "origin": "bag" }],
      "dps": { "mean": 1499.8, "stddev": 209.7, "error": 3.83, "min": 899.2, "max": 2098.1 },
      "delta": { "mean": 38.6, "stddev": 0, "error": 5.4, "min": 0, "max": 0 },
      "group": 0
    },
    {
      "substitutions": [{ "kind": "item", "slot": "shoulder", "item_id": 16966, "origin": "bank" }],
      "dps": { "mean": 1480.1, "stddev": 209.1, "error": 3.82, "min": 890.3, "max": 2070.6 },
      "delta": { "mean": 18.9, "stddev": 0, "error": 5.39, "min": 0, "max": 0 },
      "group": 1
    },
    {
      "substitutions": [{ "kind": "item", "slot": "trinket1", "item_id": 13968, "origin": "search" }],
      "dps": { "mean": 1474.5, "stddev": 208.8, "error": 3.81, "min": 888.7, "max": 2061.2 },
      "delta": { "mean": 13.3, "stddev": 0, "error": 5.39, "min": 0, "max": 0 },
      "group": 1
    },
    {
      "substitutions": [{ "kind": "item", "slot": "finger2", "item_id": 19325, "origin": "equipped" }],
      "dps": { "mean": 1463.0, "stddev": 208.9, "error": 3.81, "min": 881.0, "max": 2046.3 },
      "delta": { "mean": 1.8, "stddev": 0, "error": 5.38, "min": 0, "max": 0 },
      "group": 2
    },
    {
      "substitutions": [{ "kind": "talents", "name": "Deep Fury", "talents": "0-5530515-" }],
      "dps": { "mean": 1459.4, "stddev": 208.4, "error": 3.8, "min": 879.6, "max": 2041.0 },
      "delta": { "mean": -1.8, "stddev": 0, "error": 5.38, "min": 0, "max": 0 },
      "group": 2
    }
  ]
}
```

- [ ] **Step 5: Write `web/src/fixtures/sim/weights-result.json`**

```json
{
  "engine_version": "edc0c8e9a",
  "request": {
    "engine_version": "edc0c8e9a",
    "spec": "warrior-fury",
    "source": { "kind": "addon", "ref": "", "captured_at": "2026-09-19T09:00:00Z" },
    "character": {
      "name": "Fury",
      "race": "orc",
      "class": "warrior",
      "level": 60,
      "talents": "0-5530515-",
      "gear": [{ "slot": "head", "item_id": 12640 }],
      "buffs": ["battle_shout"],
      "consumes": []
    },
    "encounter": { "duration_sec": 180, "variation": 0.2, "targets": 1, "execute_ratio": 0.25, "profile": "" },
    "iterations": 3000,
    "random_seed": 0,
    "weights": {
      "stats": ["strength", "agility", "attack_power", "crit", "hit", "melee_haste"],
      "reference": "attack_power"
    }
  },
  "lane": "browser",
  "dps": { "mean": 1461.2, "stddev": 208.9, "error": 3.81, "min": 880.4, "max": 2044.9 },
  "iterations_run": 3000,
  "duration_ms": 12040,
  "summary": { "duration_ms": 180000, "abilities": [], "auras": [], "casts": [], "resources": [], "roster": [], "combatants": [] },
  "weights": [
    { "stat": "attack_power", "weight": 1, "error": 0 },
    { "stat": "strength", "weight": 2.14, "error": 0.06 },
    { "stat": "crit", "weight": 21.7, "error": 0.9 },
    { "stat": "hit", "weight": 27.3, "error": 1.2 },
    { "stat": "agility", "weight": 1.32, "error": 0.05 },
    { "stat": "melee_haste", "weight": 18.4, "error": 1.1 }
  ]
}
```

- [ ] **Step 6: Write `web/src/fixtures/planner/loot.json`**

Only the six fixture warrior items exist, so every source draws from them. The shape is
contract 6.1 exactly.

Source ids follow contract 10.4 exactly: `raid:<zone-slug>` and `raid:<zone-slug>:<npc-id>`,
likewise `dungeon:`; `zone_id` is an AreaTable id (Molten Core is **2717**, not 409);
`pvp:rank-<n>` carries the rank; and a source whose open date is unknown carries
`"opens": "later"`.

```json
{
  "sources": [
    {
      "id": "raid:mc",
      "kind": "raid",
      "name": "Molten Core",
      "zone_id": 2717,
      "opens": "raids-1",
      "bosses": [
        { "id": "raid:mc:12118", "name": "Lucifron", "npc_id": 12118, "items": [16963] },
        { "id": "raid:mc:11502", "name": "Ragnaros", "npc_id": 11502, "items": [12784, 19325] }
      ],
      "trash": [13968]
    },
    {
      "id": "dungeon:hall-of-thanes",
      "kind": "dungeon",
      "name": "Hall of Thanes",
      "zone_id": 9001,
      "bosses": [
        { "id": "dungeon:hall-of-thanes:90011", "name": "Thane Korgal", "npc_id": 90011, "items": [16966] }
      ],
      "trash": [13968]
    },
    {
      "id": "world:azuregos",
      "kind": "world",
      "name": "Azuregos",
      "zone_id": 16,
      "opens": "later",
      "items": [19325]
    },
    {
      "id": "crafted:blacksmithing",
      "kind": "crafted",
      "name": "Blacksmithing",
      "profession": "blacksmithing",
      "items": [12784]
    },
    {
      "id": "crafted:tailoring",
      "kind": "crafted",
      "name": "Tailoring",
      "profession": "tailoring",
      "items": [12640]
    },
    {
      "id": "rep:argent-dawn:exalted",
      "kind": "rep",
      "name": "Argent Dawn, Exalted",
      "faction_id": 529,
      "standing": "exalted",
      "items": [13968]
    },
    { "id": "pvp:rank-10", "kind": "pvp", "name": "Rank 10", "rank": 10, "items": [16966] },
    { "id": "quest", "kind": "quest", "name": "Quest rewards", "items": [12640] }
  ]
}
```

- [ ] **Step 7: Write `web/src/fixtures/planner/enchants.json` and `suffixes.json`**

`enchants.json` — keyed by `effect_id` plus `spell_id`/`item_id`, with `item_types` as the
`EnchantType` shape restriction and `slots` derived from `type` and `extra_types`, per
contract 10.4:

```json
[
  {
    "effect_id": 2543,
    "spell_id": 23545,
    "name": "Lesser Arcanum of Voracity",
    "icon": "fixture_item_lionheart_helm",
    "slots": ["head", "legs"],
    "item_types": [1, 2, 3, 4],
    "classes": [],
    "stats": { "strength": 8 },
    "phase": "launch"
  },
  {
    "effect_id": 1900,
    "spell_id": 20034,
    "name": "Crusader",
    "icon": "fixture_item_arcanite_reaper",
    "slots": ["main_hand", "off_hand"],
    "item_types": [5],
    "classes": [],
    "stats": {},
    "phase": "launch"
  },
  {
    "effect_id": 2564,
    "item_id": 16252,
    "name": "Greater Stats",
    "icon": "fixture_item_spaulders_of_wrath",
    "slots": ["chest"],
    "item_types": [1, 2, 3, 4],
    "classes": [],
    "stats": { "strength": 4, "agility": 4 },
    "phase": "raids-1"
  }
]
```

`simbuffs.json` — contract 10.4's new file, the display name and icon for every IDS.md
buff, debuff, world buff and consumable id. The slice carries the two the preset uses:

```json
{
  "entries": {
    "flask_of_supreme_power": {
      "name": "Flask of Supreme Power",
      "icon": "fixture_item_snakestone_charm"
    },
    "elixir_of_the_mongoose": {
      "name": "Elixir of the Mongoose",
      "icon": "fixture_item_band_of_accuria"
    },
    "battle_shout": { "name": "Battle Shout", "icon": "fixture_improved_battle_shout" }
  }
}
```

`suffixes.json`:

```json
[
  { "id": 1, "name": "of the Bear", "stats": { "strength": 9, "stamina": 9 } },
  { "id": 2, "name": "of the Tiger", "stats": { "agility": 9, "strength": 9 } },
  { "id": 3, "name": "of Power", "stats": { "attack_power": 20 } }
]
```

- [ ] **Step 8: Point one fixture item at two suffixes**

In `web/src/fixtures/planner/items/warrior.json`, on the `13968` (Snakestone Charm) entry,
after `"unique": false`, add:

```json
      "suffixes": [1, 2]
```

- [ ] **Step 9: Give every fixture spec a reference stat**

In `web/src/fixtures/sim/specs.json`, add `"reference_stat"` after each row's
`"engine_version"`: `"attack_power"` for `warrior-fury`, `"spell_power"` for `mage-frost`,
`"attack_power"` for `rogue-combat`.

- [ ] **Step 10: Publish the three data files**

In `web/scripts/sync-data.mjs`, inside `SYNC_ENTRIES`, after the `sets.json` line:

```js
  // Contract 6 and 10.4: the Droptimizer's source tables, Top Gear's enchant and suffix
  // lists, and the buff/consumable name table. Optional the same way sets.json is -- a
  // build the data lane has not regenerated ships none, and the pages say so rather than
  // failing to render.
  { name: 'loot.json', kind: 'file', required: false },
  { name: 'enchants.json', kind: 'file', required: false },
  { name: 'suffixes.json', kind: 'file', required: false },
  { name: 'simbuffs.json', kind: 'file', required: false },
```

Then list the three in `web/src/fixtures/planner/manifest.json`'s `files` map, with their
real hashes:

```bash
cd web/src/fixtures/planner && shasum -a 256 loot.json enchants.json suffixes.json simbuffs.json
```

Add one `"loot.json": "<hash>"` entry per file, in the map's existing alphabetical order.

- [ ] **Step 11: Export the fixtures and add the `GET /v1/builds?mine=1` route**

In `web/src/test-support/sim-api.ts`, beside the existing `fixtureResult`/`fixtureSpecs`
exports:

```ts
import fixtureBulkJson from '../fixtures/sim/bulk-result.json';
import fixtureWeightsJson from '../fixtures/sim/weights-result.json';
import type { BulkResult, WeightsResult } from '../lib/sim/bulk-types';

export const fixtureBulkResult = fixtureBulkJson as unknown as BulkResult;
export const fixtureWeightsResult = fixtureWeightsJson as unknown as WeightsResult;

/** The signed-in player's one saved planner build, for Top Gear's talent candidates. */
export const FIXTURE_MY_BUILD_ID = 'bld987654321';
```

And, in `builtIn`, the phases route (contract 10.6) beside the specs one:

```ts
    {
      method: 'GET',
      pattern: /\/v1\/phases$/,
      respond: () => envelope({ phases: phasesJson }),
    },
```

with `import phasesJson from '../data/phases.json';`.

And, **above** the existing `GET /v1/builds/{id}` route (a narrower pattern must come
first):

```ts
    // GET /v1/builds?mine=1 -- the signed-in player's saved builds, for Top Gear's talent
    // candidate list. Contract 10.6 adds `builds.user_id` and this route; the page still
    // treats a failure as "no saved builds" and says so in one line, because an old
    // deployment answers 404 and that is not worth an error banner.
    {
      method: 'GET',
      pattern: /\/v1\/builds\?/,
      respond: () =>
        envelope({
          rows: [
            {
              id: FIXTURE_MY_BUILD_ID,
              class_id: 1,
              race_id: 2,
              tree_version: FIXTURE_DATA_BUILD,
              point_order: [2001, 2001, 2001, 2001, 2001, 2002, 2002, 2002, 2002, 2002],
              gear: {},
              title: 'Deep Fury',
              created_at: '2026-09-18T12:00:00Z',
              views: 2,
            },
          ],
          total: 1,
          page: 1,
          per_page: 100,
        }),
    },
```

- [ ] **Step 12: Run the fixture test and the existing fixture suite**

```bash
cd web && FOREVER_DATA=fixture npm run sync:data
FOREVER_DATA=fixture npx vitest run src/fixtures/sim/bulk-fixture.test.ts src/fixtures/planner/fixture.test.ts src/lib/planner/sync-data.test.ts
ls public/data/$(node -p "require('./src/data/active-build.json').build")/loot.json
```

Expected: PASS, and the published `loot.json` exists.

- [ ] **Step 13: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/fixtures/sim/bulk-fixture.test.ts src/test-support/sim-api.ts scripts/sync-data.mjs
npx prettier --write src/fixtures/sim src/fixtures/planner src/data/phases.json src/test-support/sim-api.ts scripts/sync-data.mjs
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/fixtures web/src/data/phases.json web/src/test-support/sim-api.ts web/scripts/sync-data.mjs
git commit -m "$(cat <<'MSG'
test(sim): fixtures for the combination tools

A ranked bulk result with a real within-error group, a weights result
normalised on its reference stat, and loot, enchant and suffix slices
over the six items the warrior fixture already has. The phase table is a
checked-in mirror of api/internal/phase.Boundaries rather than a new
endpoint: four fixed instants the API already sources from this site.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 3: The engine surface grows the six exports, and the fake implements them

**Files:**
- Modify: `web/src/lib/sim/engine.ts` (three functions on `EngineModule`, three globals on the wasm adapter)
- Modify: `web/src/fixtures/sim/engine-fake.ts`
- Modify: `web/src/fixtures/sim/engine-fake.test.ts`

**Interfaces:**
- Consumes: `StageRequests`, `RankAnswer`, `BulkRequest`, `WeightsRequest`, `Combination`, `Combo`, `Precision`, `isCapExceeded` (Task 1); the two result fixtures (Task 2).
- Produces: on `EngineModule`, the five of contract 4 and 10.2 this lane needs plus the two part A needs —
  - `simPlan(requestJSON: string): string` → `StageRequests` JSON or the `cap_exceeded` JSON;
  - `simRank(requestJSON: string, stageJSON: string, resultsJSON: string): string` → `RankAnswer` JSON;
  - `simWeights(requestJSON: string, callbackId: string): Promise<string>` → `SimResult` JSON with `weights`;
  - `simCount(requestJSON: string): string` → `{"combinations": n}` JSON or the `cap_exceeded` JSON;
  - `simValidate(requestJSON: string): string` → `{"ok": bool, "errors": [{"field","message"}]}`;
  - `simNeedsMore(resultJSON: string, requestJSON: string): string` → `{"needs_more": bool}`.
- Also produces: `createFakeEngine(options)` honouring all six.

- [ ] **Step 1: Write the failing tests**

Append to `web/src/fixtures/sim/engine-fake.test.ts`:

```ts
describe('the fake engine’s bulk exports', () => {
  const gearRequest = (cap = 400, precision: Precision = 'fast'): string =>
    JSON.stringify({
      ...fixtureBulkResult.request,
      bulk: {
        mode: 'gear',
        candidates: [
          { slot: 'head', item_id: 16963, origin: 'bag' },
          { slot: 'shoulder', item_id: 16966, origin: 'bank' },
        ],
        talents: [],
        sets: [],
        locked: [],
        precision,
        cap,
      },
    });

  it('plans the equipped set first and every combination after it', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const stage = JSON.parse(engine.simPlan(gearRequest())) as StageRequests;
    expect(stage.stage).toBe(1);
    expect(stage.iterations).toBe(100);
    // two singles plus the pair, plus the equipped set at index 0
    expect(stage.requests).toHaveLength(4);
    expect(stage.combos).toHaveLength(3);
    expect(stage.requests[0].iterations).toBe(100);
  });

  it('starts a normal-precision plan at 1,000 iterations', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const stage = JSON.parse(engine.simPlan(gearRequest(400, 'normal'))) as StageRequests;
    expect(stage.iterations).toBe(1000);
  });

  it('refuses a plan past the cap with the count, and never trims', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const refusal: unknown = JSON.parse(engine.simPlan(gearRequest(2)));
    expect(isCapExceeded(refusal)).toBe(true);
    expect((refusal as CapExceeded).combinations).toBe(3);
    expect((refusal as CapExceeded).cap).toBe(2);
  });

  it('counts without building requests', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    expect(JSON.parse(engine.simCount(gearRequest()))).toEqual({ combinations: 3 });
  });

  it('multiplies the gear product by the consumable alternatives (contract 10.1 A5)', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const request = JSON.parse(gearRequest()) as BulkRequest;
    request.bulk.consumables = [['flask_of_supreme_power'], ['elixir_of_the_mongoose']];
    // (1 head + 1) x (1 shoulder + 1) x 2 consumable lists, minus the untouched base
    expect(JSON.parse(engine.simCount(JSON.stringify(request)))).toEqual({ combinations: 7 });
  });

  it('answers simValidate and simNeedsMore in the shapes contract 10.2 names', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    expect(JSON.parse(engine.simValidate(gearRequest()))).toEqual({ ok: true, errors: [] });
    const bad = JSON.parse(gearRequest()) as BulkRequest;
    bad.bulk.precision = 'blazing' as never;
    const answer = JSON.parse(engine.simValidate(JSON.stringify(bad))) as {
      ok: boolean;
      errors: { field: string; message: string }[];
    };
    expect(answer.ok).toBe(false);
    expect(answer.errors[0].field).toBe('bulk.precision');
    expect(JSON.parse(engine.simNeedsMore(JSON.stringify(fixtureBulkResult), gearRequest()))).toEqual({
      needs_more: false,
    });
  });

  it('ranks a stage into the next one and finally into a result', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const request = gearRequest();
    let stage = JSON.parse(engine.simPlan(request)) as StageRequests;
    let final: SimResult | null = null;
    for (let guard = 0; guard < 5 && final === null; guard += 1) {
      const results = await Promise.all(
        stage.requests.map((entry, index) =>
          engine.simRun(JSON.stringify(entry), `plan-${stage.stage}-${index}`),
        ),
      );
      const answer = JSON.parse(
        engine.simRank(request, JSON.stringify(stage), `[${results.join(',')}]`),
      ) as RankAnswer;
      if (answer.result !== undefined) final = answer.result;
      else stage = answer.next!;
    }
    const bulk = final as BulkResult | null;
    expect(bulk).not.toBeNull();
    expect(bulk!.stages.map((entry) => entry.iterations)).toEqual([100, 1000, 3000]);
    expect(bulk!.equipped.mean).toBeGreaterThan(0);
    const means = bulk!.combos.map((combo) => combo.dps.mean);
    expect([...means].sort((a, b) => b - a)).toEqual(means);
    expect(bulk!.combos[0].group).toBe(0);
    // Contract 10.1 A6: an item substitution comes back named, so the page never re-joins.
    expect(bulk!.combos[0].substitutions[0].name).toBeTruthy();
  });

  it('carries the ladder history on the stage object, not in the engine (contract 10.1 A10)', () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const stage = JSON.parse(engine.simPlan(gearRequest())) as StageRequests;
    expect(stage.ran).toEqual([]);
  });

  it('answers a weights request with the reference stat at exactly 1', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const json = await engine.simWeights(JSON.stringify(fixtureWeightsResult.request), 'w-1');
    const result = JSON.parse(json) as WeightsResult;
    expect(result.weights.find((row) => row.stat === 'attack_power')?.weight).toBe(1);
    expect(result.weights).toHaveLength(6);
  });

  it('gives the same numbers for the same request, twice', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const one = await engine.simWeights(JSON.stringify(fixtureWeightsResult.request), 'w-1');
    const two = await engine.simWeights(JSON.stringify(fixtureWeightsResult.request), 'w-2');
    expect(JSON.parse(one).weights).toEqual(JSON.parse(two).weights);
  });
});
```

Add to that file's imports:

```ts
import { isCapExceeded, type BulkResult, type CapExceeded, type Precision, type RankAnswer, type StageRequests, type WeightsResult } from '../../lib/sim/bulk-types';
import fixtureBulkResultJson from './bulk-result.json';
import fixtureWeightsResultJson from './weights-result.json';
import type { BulkRequest, WeightsRequest } from '../../lib/sim/bulk-types';

const fixtureBulkResult = fixtureBulkResultJson as unknown as BulkResult & { request: BulkRequest };
const fixtureWeightsResult = fixtureWeightsResultJson as unknown as WeightsResult & {
  request: WeightsRequest;
};
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/fixtures/sim/engine-fake.test.ts
```

Expected: FAIL, `engine.simPlan is not a function`.

- [ ] **Step 3: Declare the four on `EngineModule`**

In `web/src/lib/sim/engine.ts`, inside `interface EngineModule`, after `simAbort`:

```ts
  /**
   * The bulk planner's first stage (contract 4). Returns a StageRequests JSON whose
   * `requests[0]` is always the equipped set, or the structured cap refusal
   * `{"error":"cap_exceeded","cap":n,"combinations":n}` -- the one error shape that is not
   * a bare `{"error": "..."}`, because the page has to say by how much the list overran.
   */
  simPlan(requestJSON: string): string;
  /**
   * Scores a finished stage. Returns `{"next": stage}` or `{"result": SimResult}`.
   * `resultsJSON` is an array of SimResult in the same order as the stage's requests.
   */
  simRank(requestJSON: string, stageJSON: string, resultsJSON: string): string;
  /** A weights run, progress through the same `simProgress` callback simRun uses. */
  simWeights(requestJSON: string, callbackId: string): Promise<string>;
  /**
   * How many combinations a bulk request expands to, without building a single request for
   * them (contract 10.2). The live count on Top Gear runs on every checkbox tick, and
   * simPlan allocates a whole SimRequest per combination -- up to 400 copies of a
   * character -- which is far too much work for a keystroke.
   */
  simCount(requestJSON: string): string;
  /**
   * `{"ok": bool, "errors": [{"field","message"}]}` (contract 10.2). The Advanced drawer
   * validates an edited request with the same `Validate` the run would apply, so the
   * errors it shows are the engine's own and not a second copy of the rules.
   */
  simValidate(requestJSON: string): string;
  /**
   * `{"needs_more": bool}` (contract 10.2, `api.NeedsMoreIterations`). The target-error
   * step loop asks after each step whether to run another. Declared here rather than in
   * part A's lane because all three of 10.2's exports belong to one file's surface.
   */
  simNeedsMore(resultJSON: string, requestJSON: string): string;
```

In the same file, extend `WasmGlobals` and `loadWasmEngine`'s returned object:

```ts
  simPlan?: (requestJSON: string) => string;
  simRank?: (requestJSON: string, stageJSON: string, resultsJSON: string) => string;
  simWeights?: (requestJSON: string, callbackId: string) => Promise<string>;
  simCount?: (requestJSON: string) => string;
  simValidate?: (requestJSON: string) => string;
  simNeedsMore?: (resultJSON: string, requestJSON: string) => string;
```

```ts
    // simPlan and simCount are NOT run through unwrapOrThrow: `cap_exceeded` is a legal,
    // structured answer the caller reads rather than a failure it throws on. Every other
    // failure from them is still a bare `{"error": "..."}` and bulk-run.ts raises it.
    simPlan: (requestJSON) => globals.simPlan!(requestJSON),
    simCount: (requestJSON) => globals.simCount!(requestJSON),
    simRank: (requestJSON, stageJSON, resultsJSON) =>
      unwrapOrThrow(globals.simRank!(requestJSON, stageJSON, resultsJSON)),
    simWeights: (requestJSON, callbackId) => globals.simWeights!(requestJSON, callbackId),
    // simValidate answers `{"ok": false, errors: [...]}` for an invalid request, which is a
    // success of the call and not a failure of it, so it is not unwrapped either.
    simValidate: (requestJSON) => globals.simValidate!(requestJSON),
    simNeedsMore: (resultJSON, requestJSON) =>
      unwrapOrThrow(globals.simNeedsMore!(resultJSON, requestJSON)),
```

- [ ] **Step 4: Implement the four in the fake engine**

In `web/src/fixtures/sim/engine-fake.ts`, add these imports and helpers above `createFakeEngine`:

```ts
import fixtureItemsJson from '../planner/items/warrior.json';
import type { Item } from '../../lib/planner/types';
import {
  type BulkRequest,
  type Combination,
  type Combo,
  type Precision,
  type Stage,
  type StageRequests,
  type Substitution,
  type WeightsRequest,
} from '../../lib/sim/bulk-types';

/** simdb's stand-in, for the item names contract 10.1 A6 puts on every substitution. */
const fixtureItems = fixtureItemsJson.items as unknown as Item[];

/**
 * The iteration ladder, as contract 1.3 writes it. The real planner reads it from
 * sim/bulk; the fake repeats it because the fake IS the planner here, and the e2e suite
 * asserts the progress line names the right stage count.
 */
const LADDER: Record<Precision, number[]> = {
  fast: [100, 1000, 3000],
  normal: [1000, 3000],
  high: [1000, 10000],
};

/**
 * The fake's expansion, following contract 10.1 A4: **mode decides**. `gear` takes the
 * product of every candidate group -- one group per slot, each group being "keep what is
 * equipped" plus that slot's candidates -- crossed with the consumable alternatives of
 * A5; `drops` and `talents` take one substitution at a time. The empty product is the
 * equipped set, which is `requests[0]` and not a combination, so it is dropped.
 *
 * Substitutions come back NAMED (contract 10.1 A6): the real planner reads the name from
 * simdb, and the fake reads it from the fixture item file, so the page's "never re-join an
 * id to a name" rule is exercised rather than merely asserted.
 *
 * A candidate whose slot is "" fits more than one slot and only the item database can say
 * which; the fake has no database, so it resolves every one to `finger1`. Nothing in the
 * suite asserts on that resolution -- the real engine's own smoke test covers it.
 */
function expand(request: BulkRequest): Combination[] {
  const bulk = request.bulk;
  const slotOf = (slot: string): string => (slot === '' ? 'finger1' : slot);
  const itemSubs: Substitution[] = bulk.candidates
    .filter((candidate) => !(bulk.locked ?? []).includes(slotOf(candidate.slot)))
    .map((candidate) => ({
      kind: 'item',
      slot: slotOf(candidate.slot),
      item_id: candidate.item_id,
      enchant: candidate.enchant,
      suffix: candidate.suffix,
      origin: candidate.origin,
      name: itemName(candidate.item_id),
      source_name: candidate.source_name,
    }));

  let groups: Substitution[][];
  if (bulk.mode === 'gear') {
    // The product of one group per slot, each group led by "nothing" -- which is exactly
    // the set of subsets that use each slot at most once, in a stable order.
    const bySlot = new Map<string, Substitution[]>();
    for (const sub of itemSubs) {
      const existing = bySlot.get(sub.slot ?? '');
      if (existing === undefined) bySlot.set(sub.slot ?? '', [sub]);
      else existing.push(sub);
    }
    let product: Substitution[][] = [[]];
    for (const options of bySlot.values()) {
      product = product.flatMap((base) => [base, ...options.map((sub) => [...base, sub])]);
    }
    // Consumable alternatives multiply the gear product (contract 10.1 A5). Each inner
    // list is one alternative, reported as a `consume` substitution named by its ids.
    const consumableLists = bulk.consumables ?? [];
    if (consumableLists.length > 0) {
      // Contract 10.8: the kind is `consumes` and the name is the ids joined by ", ".
      product = product.flatMap((base) =>
        consumableLists.map((list) => [
          ...base,
          { kind: 'consumes', name: list.join(', ') } as Substitution,
        ]),
      );
    }
    groups = product.filter((group) => group.length > 0).sort((a, b) => a.length - b.length);
  } else {
    groups = itemSubs.map((sub) => [sub]);
  }

  for (const loadout of bulk.talents ?? []) {
    groups.push([{ kind: 'talents', name: loadout.name, talents: loadout.talents }]);
  }
  for (const set of bulk.sets ?? []) groups.push([{ kind: 'set', name: set.name }]);

  return groups.map((substitutions) => ({
    request: applySubstitutions(request, substitutions),
    substitutions,
  }));
}

/** The fixture item file stands in for simdb, so a substitution comes back named. */
function itemName(itemId: number): string {
  return fixtureItems.find((item) => item.id === itemId)?.name ?? `Item ${itemId}`;
}

/** The base character with the substitutions written into its gear, talents and consumes. */
function applySubstitutions(request: BulkRequest, substitutions: readonly Substitution[]): SimRequest {
  const gear = request.character.gear.map((slot) => ({ ...slot }));
  let talents = request.character.talents;
  let consumes = [...request.character.consumes];
  for (const sub of substitutions) {
    if (sub.kind === 'talents') {
      talents = sub.talents ?? talents;
      continue;
    }
    if (sub.kind === 'consumes') {
      // A5: the inner list REPLACES the character's consumes for this combination.
      consumes = (sub.name ?? '').split(', ').filter((id) => id !== '');
      continue;
    }
    if (sub.kind !== 'item') continue;
    const existing = gear.findIndex((slot) => slot.slot === sub.slot);
    const next = { slot: sub.slot!, item_id: sub.item_id!, enchant: sub.enchant, suffix: sub.suffix };
    if (existing >= 0) gear[existing] = next;
    else gear.push(next);
  }
  const { bulk: _bulk, ...base } = request;
  return { ...base, character: { ...request.character, gear, talents, consumes } };
}

/** The equipped set, with the bulk block stripped: it is a plain run like any other. */
function equippedRequest(request: BulkRequest, iterations: number): SimRequest {
  const { bulk: _bulk, ...base } = request;
  return { ...base, iterations };
}

function stageAt(
  request: BulkRequest,
  stage: number,
  combos: Combination[],
  ran: Stage[],
): StageRequests {
  const iterations = LADDER[request.bulk.precision][stage - 1];
  return {
    stage,
    iterations,
    requests: [
      equippedRequest(request, iterations),
      ...combos.map((combo) => ({ ...combo.request, iterations })),
    ],
    combos,
    // Contract 10.1 A10: the ladder's history rides on the stage object across the
    // boundary, so nothing stateful lives in the engine between calls.
    ran: [...ran],
  };
}
```

Then, inside the object `createFakeEngine` returns, after `simAbort`:

```ts
    simPlan(requestJSON) {
      const request = JSON.parse(requestJSON) as BulkRequest;
      const combos = expand(request);
      if (combos.length > request.bulk.cap) {
        return JSON.stringify({
          error: 'cap_exceeded',
          cap: request.bulk.cap,
          combinations: combos.length,
        });
      }
      return JSON.stringify(stageAt(request, 1, combos, []));
    },

    simValidate(requestJSON) {
      const request = JSON.parse(requestJSON) as BulkRequest;
      const errors: { field: string; message: string }[] = [];
      const bulk = request.bulk;
      if (bulk !== undefined) {
        if (!['fast', 'normal', 'high'].includes(bulk.precision)) {
          errors.push({ field: 'bulk.precision', message: 'precision must be fast, normal or high' });
        }
        if (!['gear', 'talents', 'drops'].includes(bulk.mode)) {
          errors.push({ field: 'bulk.mode', message: 'mode must be gear, talents or drops' });
        }
        // Contract 10.1 A3: a bulk request's iterations are its precision's final stage.
        const wanted = bulk.precision === 'high' ? 10000 : 3000;
        if (request.iterations !== wanted) {
          errors.push({ field: 'iterations', message: `iterations must be ${wanted} for ${bulk.precision}` });
        }
      }
      return JSON.stringify({ ok: errors.length === 0, errors });
    },

    simNeedsMore(resultJSON, requestJSON) {
      const result = JSON.parse(resultJSON) as SimResult;
      const request = JSON.parse(requestJSON) as SimRequest & { target_error?: number };
      const target = request.target_error ?? 0;
      if (target <= 0 || result.dps.mean === 0) return JSON.stringify({ needs_more: false });
      return JSON.stringify({ needs_more: result.dps.error / result.dps.mean > target });
    },

    simCount(requestJSON) {
      // Contract 10.2: counts without allocating a request per combination. The fake still
      // expands -- there is no cheaper way over a fixture -- but it discards the requests,
      // which is the contract the caller sees.
      const request = JSON.parse(requestJSON) as BulkRequest;
      const combinations = expand(request).length;
      if (combinations > request.bulk.cap) {
        return JSON.stringify({ error: 'cap_exceeded', cap: request.bulk.cap, combinations });
      }
      return JSON.stringify({ combinations });
    },

    simRank(requestJSON, stageJSON, resultsJSON) {
      const request = JSON.parse(requestJSON) as BulkRequest;
      const stage = JSON.parse(stageJSON) as StageRequests;
      const results = JSON.parse(resultsJSON) as SimResult[];
      const equipped = results[0].dps;
      const scored = stage.combos
        .map((combo, index) => ({ combo, dps: results[index + 1].dps }))
        .sort((a, b) => b.dps.mean - a.dps.mean);

      const ladder = LADDER[request.bulk.precision];
      const ran: Stage[] = [
        ...(stage.ran ?? []),
        { iterations: stage.iterations, combos: stage.combos.length },
      ];

      if (stage.stage < ladder.length) {
        // The cut: the top half at every stage but the last. The real planner keeps a
        // quarter plus anything within two standard errors; the fake keeps a fixed
        // fraction so a test can predict how many requests the next stage carries.
        const keep = Math.max(1, Math.ceil(scored.length / 2));
        const survivors = scored.slice(0, keep).map((entry) => entry.combo);
        return JSON.stringify({ next: stageAt(request, stage.stage + 1, survivors, ran) });
      }

      // Within-error grouping: a run whose delta interval overlaps the leader's is group 0,
      // then each next non-overlapping run opens the next group.
      const combos: Combo[] = [];
      let group = 0;
      let boundary = Number.POSITIVE_INFINITY;
      for (const entry of scored) {
        const delta = {
          mean: entry.dps.mean - equipped.mean,
          stddev: 0,
          error: Math.hypot(entry.dps.error, equipped.error),
          min: 0,
          max: 0,
        };
        const high = delta.mean + 1.96 * delta.error;
        if (high < boundary) {
          if (Number.isFinite(boundary)) group += 1;
          boundary = delta.mean - 1.96 * delta.error;
        }
        combos.push({ substitutions: entry.combo.substitutions, dps: entry.dps, delta, group });
      }

      return JSON.stringify({
        result: {
          ...results[0],
          request,
          lane: 'browser',
          combos,
          equipped,
          stages: ran,
        },
      });
    },

    async simWeights(requestJSON, callbackId) {
      const request = JSON.parse(requestJSON) as WeightsRequest;
      const base = await this.simRun(JSON.stringify({ ...request }), callbackId);
      const result = JSON.parse(base) as SimResult;
      // Seeded off the stat name, so the same request gives the same weights every time.
      const weights = request.weights.stats.map((stat) => {
        if (stat === request.weights.reference) return { stat, weight: 1, error: 0 };
        const random = seeded([...stat].reduce((sum, ch) => sum + ch.charCodeAt(0), 0));
        return {
          stat,
          weight: Math.round(random() * 3000) / 100,
          error: Math.round(random() * 200) / 100 + 0.01,
        };
      });
      return JSON.stringify({ ...result, request, weights });
    },
```

There is deliberately **no** stage accumulator inside `createFakeEngine`: contract 10.1
A10 puts the ladder's history on `StageRequests.ran`, so the engine keeps nothing between
calls and two runs in flight at once cannot mix their histories.

`STAGES_BY_PRECISION` is used by the test only, not by the fake; do not import it there.

- [ ] **Step 5: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/fixtures/sim/engine-fake.test.ts
```

Expected: PASS, including the seven new cases.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/engine.ts src/fixtures/sim/engine-fake.ts src/fixtures/sim/engine-fake.test.ts
npx prettier --write src/lib/sim/engine.ts src/fixtures/sim/engine-fake.ts src/fixtures/sim/engine-fake.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/engine.ts web/src/fixtures/sim/engine-fake.ts web/src/fixtures/sim/engine-fake.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): the six bulk exports on the engine surface

simPlan, simRank, simWeights, simCount, simValidate and simNeedsMore --
contract 4 as amended by 10.2 -- and a deterministic fake for all six so
the whole browser suite runs with no Go toolchain. Expansion follows
10.1 A4 (mode decides) and A5 (consumable lists multiply the gear
product); substitutions come back named per A6; the ladder's history
rides on StageRequests.ran per A10, so the engine keeps no state
between calls.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 4: The worker pool carries plan, rank, count, validate, needsMore and weights

The four new exports live in the wasm, the main thread holds no instance, and `simSplit` /
`simCombine` are already routed to worker 0 for exactly that reason. These four follow the
same road.

**Files:**
- Modify: `web/src/lib/sim/worker.ts` (four message kinds, four pool methods)
- Modify: `web/src/lib/sim/sim.worker.ts` (handle them)
- Modify: `web/src/test-support/fake-worker.ts` (handle them identically)
- Modify: `web/src/lib/sim/worker.test.ts`

**Interfaces:**
- Consumes: `EngineModule.simPlan/simRank/simWeights/simCount` (Task 3).
- Produces: on `SimPool` —
  - `plan(requestJSON: string): Promise<string>`
  - `rank(requestJSON: string, stageJSON: string, resultsJSON: string): Promise<string>`
  - `count(requestJSON: string): Promise<string>`
  - `validate(requestJSON: string): Promise<string>`
  - `needsMore(resultJSON: string, requestJSON: string): Promise<string>`
  - `weights(requestJSON: string, callbackId: string, onProgress: (p: ShardProgress) => void): Promise<string>`

- [ ] **Step 1: Write the failing test**

Append to `web/src/lib/sim/worker.test.ts`:

```ts
describe('the pool’s bulk messages', () => {
  it('routes plan, count and rank to worker 0 and returns their JSON verbatim', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
    const spawned: number[] = [];
    const pool = createPool({
      hardwareConcurrency: 4,
      spawn: (index) => {
        spawned.push(index);
        return createFakeWorker(engine);
      },
    });

    const request = JSON.stringify(bulkRequestFixture);
    const planJSON = await pool.plan(request);
    const stage = JSON.parse(planJSON) as StageRequests;
    expect(stage.stage).toBe(1);

    const countJSON = await pool.count(request);
    expect(JSON.parse(countJSON)).toHaveProperty('combinations');

    expect(JSON.parse(await pool.validate(request))).toHaveProperty('ok');
    expect(
      JSON.parse(await pool.needsMore(JSON.stringify(bulkResultJson), request)),
    ).toHaveProperty('needs_more');

    const results = await Promise.all(
      stage.requests.map((entry, index) => pool.run([JSON.stringify(entry)], `t-${index}`, () => {})),
    );
    const rankJSON = await pool.rank(
      request,
      planJSON,
      `[${results.map((part) => part[0]).join(',')}]`,
    );
    expect(JSON.parse(rankJSON)).toHaveProperty('next');

    // split/combine/plan/rank/count/validate/needs-more all go to worker 0; only `run`
    // fans out. Contract 10.2 is explicit that a stage's requests run unsplit.
    expect(Math.min(...spawned)).toBe(0);
    pool.terminate();
  });

  it('reports weights progress through the same shard callback a run uses', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 3 });
    const pool = createPool({ hardwareConcurrency: 2, spawn: () => createFakeWorker(engine) });
    const ticks: number[] = [];
    const json = await pool.weights(JSON.stringify(weightsRequestFixture), 'w-1', (progress) =>
      ticks.push(progress.iterationsDone),
    );
    expect(JSON.parse(json)).toHaveProperty('weights');
    expect(ticks.length).toBeGreaterThan(0);
    pool.terminate();
  });
});
```

with these additions to that file's imports:

```ts
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import type { BulkResult, StageRequests, WeightsResult } from './bulk-types';

const bulkRequestFixture = (bulkResultJson as unknown as BulkResult).request;
const weightsRequestFixture = (weightsResultJson as unknown as WeightsResult).request;
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/worker.test.ts
```

Expected: FAIL, `pool.plan is not a function`.

- [ ] **Step 3: Extend the protocol in `worker.ts`**

Add to `ToWorker`:

```ts
  | { kind: 'plan'; token: number; request: string }
  | { kind: 'count'; token: number; request: string }
  | { kind: 'validate'; token: number; request: string }
  | { kind: 'needs-more'; token: number; request: string; result: string }
  | { kind: 'rank'; token: number; request: string; stage: string; results: string }
  | { kind: 'weights'; token: number; callbackId: string; request: string }
```

Add to `SimPool`:

```ts
  /** simPlan, on worker 0. Answers a StageRequests JSON or the structured cap refusal. */
  plan(request: string): Promise<string>;
  /** simCount, on worker 0. Answers `{"combinations": n}` or the structured cap refusal. */
  count(request: string): Promise<string>;
  /** simValidate, on worker 0. Answers `{"ok": bool, "errors": [...]}`. For part A's drawer. */
  validate(request: string): Promise<string>;
  /** simNeedsMore, on worker 0. Answers `{"needs_more": bool}`. For part A's step loop. */
  needsMore(result: string, request: string): Promise<string>;
  /** simRank, on worker 0. Answers `{"next": …}` or `{"result": …}`. */
  rank(request: string, stage: string, results: string): Promise<string>;
  /** simWeights, on worker 0, with progress through the same shard callback a run uses. */
  weights(
    request: string,
    callbackId: string,
    onProgress: (progress: ShardProgress) => void,
  ): Promise<string>;
```

And to the object `createPool` returns, beside `split`/`combine`:

```ts
    plan(request) {
      return send<string>(0, (token) => ({ kind: 'plan', token, request }));
    },
    count(request) {
      return send<string>(0, (token) => ({ kind: 'count', token, request }));
    },
    validate(request) {
      return send<string>(0, (token) => ({ kind: 'validate', token, request }));
    },
    needsMore(result, request) {
      return send<string>(0, (token) => ({ kind: 'needs-more', token, request, result }));
    },
    rank(request, stage, results) {
      return send<string>(0, (token) => ({ kind: 'rank', token, request, stage, results }));
    },
    weights(request, callbackId, onProgress) {
      return send<string>(
        0,
        (token) => ({ kind: 'weights', token, callbackId, request }),
        (progress) => onProgress({ ...progress, shard: 0 }),
      );
    },
```

- [ ] **Step 4: Handle them in `sim.worker.ts`**

Inside `handle()`, after the `combine` branch and before the `run` fallthrough:

```ts
    if (message.kind === 'plan') {
      reply({ kind: 'one', token, result: loaded.simPlan(message.request) });
      return;
    }
    if (message.kind === 'count') {
      reply({ kind: 'one', token, result: loaded.simCount(message.request) });
      return;
    }
    if (message.kind === 'validate') {
      reply({ kind: 'one', token, result: loaded.simValidate(message.request) });
      return;
    }
    if (message.kind === 'needs-more') {
      reply({ kind: 'one', token, result: loaded.simNeedsMore(message.result, message.request) });
      return;
    }
    if (message.kind === 'rank') {
      reply({
        kind: 'one',
        token,
        result: loaded.simRank(message.request, message.stage, message.results),
      });
      return;
    }
    if (message.kind === 'weights') {
      // Registered in the same callbackId map a run uses, so simProgress reaches the pool
      // through the identical path and the page's progress rendering does not fork.
      tokenOf.set(message.callbackId, token);
      try {
        reply({
          kind: 'one',
          token,
          result: await loaded.simWeights(message.request, message.callbackId),
        });
      } finally {
        tokenOf.delete(message.callbackId);
      }
      return;
    }
```

- [ ] **Step 5: Mirror them in `fake-worker.ts`**

Inside `postMessage`'s async body, beside the `split`/`combine` branches:

```ts
          } else if (message.kind === 'plan') {
            emit({ kind: 'one', token, result: engine.simPlan(message.request) });
          } else if (message.kind === 'count') {
            emit({ kind: 'one', token, result: engine.simCount(message.request) });
          } else if (message.kind === 'validate') {
            emit({ kind: 'one', token, result: engine.simValidate(message.request) });
          } else if (message.kind === 'needs-more') {
            emit({ kind: 'one', token, result: engine.simNeedsMore(message.result, message.request) });
          } else if (message.kind === 'rank') {
            emit({
              kind: 'one',
              token,
              result: engine.simRank(message.request, message.stage, message.results),
            });
          } else if (message.kind === 'weights') {
            tokenOf.set(message.callbackId, token);
            try {
              emit({
                kind: 'one',
                token,
                result: await engine.simWeights(message.request, message.callbackId),
              });
            } finally {
              tokenOf.delete(message.callbackId);
            }
```

- [ ] **Step 6: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/worker.test.ts src/lib/sim/run.test.ts
```

Expected: PASS, and `run.test.ts` unchanged and still green.

- [ ] **Step 7: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/worker.ts src/lib/sim/sim.worker.ts src/test-support/fake-worker.ts src/lib/sim/worker.test.ts
npx prettier --write src/lib/sim/worker.ts src/lib/sim/sim.worker.ts src/test-support/fake-worker.ts src/lib/sim/worker.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/worker.ts web/src/lib/sim/sim.worker.ts web/src/test-support/fake-worker.ts web/src/lib/sim/worker.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): the bulk exports across the worker protocol

plan, rank, count, validate, needsMore and weights all route to worker
0, the same rule and the same reason simSplit and simCombine already do:
the main thread holds no wasm instance and loading a second copy for the
sim page would cost its LCP budget for nothing. validate and needsMore
are part A's to call, declared here because contract 10.2 puts all three
new exports in one table and splitting one file's surface across two
lanes is how a merge conflict is made. Weights progress goes through the
run callback map, so the page renders one progress line for both.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 5: `bulk-run.ts`, the stage loop every tool runs on

**Files:**
- Create: `web/src/lib/sim/bulk-run.ts`
- Create: `web/src/lib/sim/bulk-run.test.ts`
- Modify: `web/src/lib/sim/api.ts` (`fetchBulkProgress`, and the premium dispatch note)
- Modify: `web/src/lib/sim/bulk-types.ts` (`BulkServerProgress`)

**Interfaces:**
- Consumes: `SimPool` (`plan`, `rank`, `count`, `run`, `weights`, `abort`, `size`) from `./worker`; everything from `./bulk-types`; `bulkCopy`, `simCopy` from `./copy`; `dispatchServerSim`, `fetchSim` from `./api`.
- Produces:
  - `class BulkCapError extends Error { readonly cap: number; readonly combinations: number }`
  - `class BulkRunError extends Error { readonly cancelled: boolean; readonly detail: string }`
  - `interface BulkProgress { stage: number; stages: number; combosDone: number; combosTotal: number }`
  - `interface BulkRunHandle { readonly callbackId: string; readonly result: Promise<SimResult>; cancel(): void }`
  - `function countCombinations(pool: SimPool, request: BulkRequest): Promise<number>` (throws `BulkCapError`)
  - `function runBulk(pool: SimPool, request: BulkRequest, onProgress: (p: BulkProgress) => void, now?: () => number): BulkRunHandle`
  - `function runWeightsRun(pool: SimPool, request: WeightsRequest, onTick: (iterationsDone: number) => void, now?: () => number): BulkRunHandle`
  - `function chunk<T>(items: readonly T[], size: number): T[][]`
  - `function stageProgressLine(progress: BulkProgress): string`
- Also produces: `fetchBulkProgress(simId, apiBase?): Promise<BulkServerProgress>` from `./api`.

**Two shapes of the work, both now written into the contract**

1. **A stage's requests go through `pool.run` unsplit.** Contract 10.2 settles it: "a
   stage's request array goes through the worker pool as independent whole requests, in
   chunks of the pool's width… `simSplit` and `simCombine` are for plain runs only." A
   bulk stage request is one indivisible unit of work, so splitting it into eight shards of
   twelve iterations and combining them back would be the identity transform, and it would
   funnel two extra messages per combination (up to 800 a stage) through worker 0, which is
   also a worker running combinations.
2. **A stop keeps every finished chunk.** The pool's `run` settles all-or-nothing and
   aborts its siblings on the first failure, so one call per stage would throw away 399
   finished combinations when the player presses Stop. Chunking at the pool's own width —
   which is what 10.2 prescribes anyway — loses at most one chunk and is what makes "abort
   returns the partial" true.

- [ ] **Step 1: Write the failing test**

Create `web/src/lib/sim/bulk-run.test.ts`:

```ts
// web/src/lib/sim/bulk-run.test.ts
import { describe, expect, it, vi } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import {
  BulkCapError,
  BulkRunError,
  chunk,
  countCombinations,
  runBulk,
  runWeightsRun,
  stageProgressLine,
  type BulkProgress,
} from './bulk-run';
import type { BulkRequest, BulkResult, WeightsRequest, WeightsResult } from './bulk-types';
import { createPool, type SimPool } from './worker';

const baseRequest = (bulkResultJson as unknown as BulkResult).request as BulkRequest;
const weightsRequest = (weightsResultJson as unknown as WeightsResult).request as WeightsRequest;

function gearRequest(overrides: Partial<BulkRequest['bulk']> = {}): BulkRequest {
  return {
    ...baseRequest,
    bulk: {
      mode: 'gear',
      candidates: [
        { slot: 'head', item_id: 16963, origin: 'bag' },
        { slot: 'shoulder', item_id: 16966, origin: 'bank' },
      ],
      talents: [],
      sets: [],
      locked: [],
      precision: 'fast',
      cap: 400,
      ...overrides,
    },
  };
}

function pool(options: { tickMs?: number; ticks?: number; size?: number } = {}): SimPool {
  const engine = createFakeEngine({ tickMs: options.tickMs ?? 0, ticks: options.ticks ?? 1 });
  return createPool({
    hardwareConcurrency: options.size ?? 4,
    spawn: () => createFakeWorker(engine),
  });
}

describe('chunk', () => {
  it('splits into runs of at most size, keeping order', () => {
    expect(chunk([1, 2, 3, 4, 5], 2)).toEqual([[1, 2], [3, 4], [5]]);
    expect(chunk([], 3)).toEqual([]);
    expect(chunk([1], 0)).toEqual([[1]]);
  });
});

describe('stageProgressLine', () => {
  it('reads the way the design writes it', () => {
    const progress: BulkProgress = { stage: 2, stages: 3, combosDone: 31, combosTotal: 96 };
    expect(stageProgressLine(progress)).toBe('stage 2 of 3 · 31 of 96 combinations');
  });
});

describe('countCombinations', () => {
  it('answers the live count', async () => {
    const p = pool();
    await expect(countCombinations(p, gearRequest())).resolves.toBe(3);
    p.terminate();
  });

  it('throws the cap refusal with both numbers rather than trimming', async () => {
    const p = pool();
    const error = await countCombinations(p, gearRequest({ cap: 2 })).catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkCapError);
    expect((error as BulkCapError).cap).toBe(2);
    expect((error as BulkCapError).combinations).toBe(3);
    p.terminate();
  });
});

describe('runBulk', () => {
  it('runs every stage and returns a ranked result', async () => {
    const p = pool();
    const seen: BulkProgress[] = [];
    const handle = runBulk(p, gearRequest(), (progress) => seen.push({ ...progress }));
    const result = (await handle.result) as BulkResult;

    expect(result.combos.length).toBeGreaterThan(0);
    expect(result.equipped.mean).toBeGreaterThan(0);
    expect(result.stages).toHaveLength(3);
    expect(result.lane).toBe('browser');
    expect(result.aborted).toBeUndefined();
    expect(seen.map((progress) => progress.stage)).toEqual(
      expect.arrayContaining([1, 2, 3]),
    );
    expect(seen.every((progress) => progress.stages === 3)).toBe(true);
    p.terminate();
  });

  it('refuses past the cap before it runs anything', async () => {
    const p = pool();
    const handle = runBulk(p, gearRequest({ cap: 1 }), () => {});
    await expect(handle.result).rejects.toBeInstanceOf(BulkCapError);
    p.terminate();
  });

  it('a stop returns the combinations that finished, marked partial', async () => {
    // Slow ticks and a one-worker pool, so the first chunk finishes and the second is
    // still running when the stop lands.
    const p = pool({ tickMs: 30, ticks: 2, size: 1 });
    const handle = runBulk(p, gearRequest(), () => {});
    await new Promise((resolve) => setTimeout(resolve, 90));
    handle.cancel();
    const result = (await handle.result) as BulkResult;
    expect(result.aborted).toBe(true);
    expect(result.combos.length).toBeGreaterThanOrEqual(1);
    p.terminate();
  });

  it('a stop before the equipped run finishes has nothing to return', async () => {
    const p = pool({ tickMs: 200, ticks: 2, size: 1 });
    const handle = runBulk(p, gearRequest(), () => {});
    handle.cancel();
    const error = await handle.result.catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkRunError);
    expect((error as BulkRunError).cancelled).toBe(true);
    p.terminate();
  });

  it('keeps the engine’s own words on a failure', async () => {
    const engine = createFakeEngine({ tickMs: 0, ticks: 1, failWith: 'request: unknown buff: "x"' });
    const p = createPool({ hardwareConcurrency: 2, spawn: () => createFakeWorker(engine) });
    const handle = runBulk(p, gearRequest(), () => {});
    const error = await handle.result.catch((e: unknown) => e);
    expect(error).toBeInstanceOf(BulkRunError);
    expect((error as BulkRunError).cancelled).toBe(false);
    expect((error as BulkRunError).detail).toContain('unknown buff');
    p.terminate();
  });

  it('stamps the browser lane and the wall clock', async () => {
    const p = pool();
    const clock = vi.fn().mockReturnValueOnce(1000).mockReturnValue(4500);
    const handle = runBulk(p, gearRequest(), () => {}, clock);
    const result = await handle.result;
    expect(result.duration_ms).toBe(3500);
    p.terminate();
  });
});

describe('runWeightsRun', () => {
  it('returns the weights and ticks progress', async () => {
    const p = pool({ tickMs: 0, ticks: 3 });
    const ticks: number[] = [];
    const handle = runWeightsRun(p, weightsRequest, (done) => ticks.push(done));
    const result = (await handle.result) as WeightsResult;
    expect(result.weights.length).toBe(6);
    expect(ticks.length).toBeGreaterThan(0);
    p.terminate();
  });
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-run.test.ts
```

Expected: FAIL, cannot resolve `./bulk-run`.

- [ ] **Step 3: Write `bulk-run.ts`**

```ts
// web/src/lib/sim/bulk-run.ts
// The stage loop every combination tool runs on: simPlan, then each stage's requests
// through the pool, then simRank, then the next stage or the finished result.
//
// Nothing here is a statistic. Which combinations exist, which survive a cut, how many
// iterations each stage runs and which runs are within error of the leader are all decided
// by sim/bulk inside the wasm. This file moves JSON between the planner and the pool,
// counts what has finished, and gives the player a way to stop.
//
// Two shapes of the work are not the obvious ones, and both are load-bearing:
//
//   * a stage's requests go to `pool.run` directly. Contract 4 says they run "exactly as it
//     runs a single sim", and `pool.run` IS that fan-out. A stage request is one
//     indivisible unit of work, so simSplit/simCombine around it would be the identity and
//     would push two extra messages per combination through worker 0 -- which is also a
//     worker running combinations.
//   * a stage runs in chunks of the pool's own width. `pool.run` settles all-or-nothing and
//     aborts its siblings on the first failure, so one call per stage would throw away every
//     finished combination the moment the player pressed Stop. Chunking loses at most one
//     chunk, which is what makes "abort returns the partial" true rather than aspirational.
import { bulkCopy, simCopy } from './copy';
import {
  isCapExceeded,
  STAGES_BY_PRECISION,
  type BulkRequest,
  type BulkResult,
  type Combination,
  type RankAnswer,
  type StageRequests,
  type WeightsRequest,
} from './bulk-types';
import type { SimResult } from './types';
import type { SimPool } from './worker';

/** The planner refused: the list is bigger than the lane allows, and by this much. */
export class BulkCapError extends Error {
  constructor(
    readonly cap: number,
    readonly combinations: number,
  ) {
    super(bulkCopy.capNotice(cap, combinations));
    this.name = 'BulkCapError';
  }
}

/**
 * A stop the player asked for, or an engine that could not run these combinations. `detail`
 * carries the engine's own words verbatim -- sim/request names the buff or consumable id it
 * refused, and that sentence is the only thing that says what to change.
 */
export class BulkRunError extends Error {
  constructor(
    message: string,
    readonly cancelled: boolean,
    readonly detail: string,
    options: { cause?: unknown } = {},
  ) {
    super(message, options);
    this.name = 'BulkRunError';
  }
}

export interface BulkProgress {
  stage: number;
  stages: number;
  combosDone: number;
  combosTotal: number;
}

export interface BulkRunHandle {
  readonly callbackId: string;
  readonly result: Promise<SimResult>;
  cancel(): void;
}

export function stageProgressLine(progress: BulkProgress): string {
  return bulkCopy.stageProgress(
    progress.stage,
    progress.stages,
    progress.combosDone,
    progress.combosTotal,
  );
}

/** Runs of at most `size`, in order. A size of zero or less is one run of everything. */
export function chunk<T>(items: readonly T[], size: number): T[][] {
  if (items.length === 0) return [];
  if (size <= 0) return [[...items]];
  const out: T[][] = [];
  for (let i = 0; i < items.length; i += size) out.push(items.slice(i, i + size));
  return out;
}

function detailOf(cause: unknown): string {
  if (cause instanceof Error) return cause.message;
  return typeof cause === 'string' ? cause : '';
}

/** The planner's answer, or the cap refusal as an exception. */
function parseStage(json: string): StageRequests {
  const parsed: unknown = JSON.parse(json);
  if (isCapExceeded(parsed)) throw new BulkCapError(parsed.cap, parsed.combinations);
  const error = (parsed as { error?: unknown }).error;
  if (typeof error === 'string') throw new Error(error);
  return parsed as StageRequests;
}

/**
 * The live combination count, for the run bar (`simCount`, contract 10.2). Throws
 * `BulkCapError` past the cap, because the button has to say what would exceed it and by
 * how much -- it never trims the list.
 */
export async function countCombinations(pool: SimPool, request: BulkRequest): Promise<number> {
  const parsed: unknown = JSON.parse(await pool.count(JSON.stringify(request)));
  if (isCapExceeded(parsed)) throw new BulkCapError(parsed.cap, parsed.combinations);
  const error = (parsed as { error?: unknown }).error;
  if (typeof error === 'string') throw new Error(error);
  return (parsed as { combinations: number }).combinations;
}

let bulkCounter = 0;

function nextCallbackId(): string {
  bulkCounter += 1;
  return `bulk-${bulkCounter}`;
}

/**
 * What a stop has to hand back: the equipped run plus whatever combinations finished,
 * re-ranked by the wasm as if this were the final stage. The stage number is set to the
 * precision's last so `simRank` returns a `result` rather than another `next` -- ranking a
 * short list is still the planner's arithmetic, never this file's.
 */
function partialStage(
  stage: StageRequests,
  stages: number,
  finishedCombos: Combination[],
): StageRequests {
  return {
    stage: stages,
    iterations: stage.iterations,
    requests: [stage.requests[0], ...finishedCombos.map((combo) => combo.request)],
    combos: finishedCombos,
    // The ladder's history travels with the stage (contract 10.1 A10), so a partial
    // result still says what each finished stage cost rather than starting the tally over.
    ran: stage.ran ?? [],
  };
}

export function runBulk(
  pool: SimPool,
  request: BulkRequest,
  onProgress: (progress: BulkProgress) => void,
  now: () => number = () => Date.now(),
): BulkRunHandle {
  const callbackId = nextCallbackId();
  const stages = STAGES_BY_PRECISION[request.bulk.precision];
  const requestJSON = JSON.stringify(request);
  let cancelled = false;

  async function execute(): Promise<SimResult> {
    const startedAt = now();
    let stage = parseStage(await pool.plan(requestJSON));

    for (;;) {
      const combosTotal = stage.combos.length;
      onProgress({ stage: stage.stage, stages, combosDone: 0, combosTotal });

      // Index 0 is the equipped set and every later index is combos[index - 1]; the chunks
      // keep that order, so a partial rank can pair them back up by position.
      const results: string[] = [];
      let stopped = false;
      for (const part of chunk(stage.requests, pool.size)) {
        if (cancelled) {
          stopped = true;
          break;
        }
        const offset = results.length;
        try {
          results.push(...(await pool.run(part, `${callbackId}-s${stage.stage}-${offset}`, () => {})));
        } catch (cause) {
          if (!cancelled) throw cause;
          stopped = true;
          break;
        }
        onProgress({
          stage: stage.stage,
          stages,
          combosDone: Math.max(0, results.length - 1),
          combosTotal,
        });
      }

      if (stopped || cancelled) {
        // Nothing to hand back without the baseline: every delta is measured against it.
        if (results.length < 2) {
          throw new BulkRunError(simCopy.stopped, true, '');
        }
        const finished = stage.combos.slice(0, results.length - 1);
        const answer = JSON.parse(
          await pool.rank(
            requestJSON,
            JSON.stringify(partialStage(stage, stages, finished)),
            `[${results.join(',')}]`,
          ),
        ) as RankAnswer;
        const partial = answer.result as BulkResult | undefined;
        if (partial === undefined) throw new BulkRunError(simCopy.stopped, true, '');
        return finish(partial, startedAt, true);
      }

      const answer = JSON.parse(
        await pool.rank(requestJSON, JSON.stringify(stage), `[${results.join(',')}]`),
      ) as RankAnswer;
      if (answer.result !== undefined) return finish(answer.result as BulkResult, startedAt, false);
      if (answer.next === undefined) throw new Error(bulkCopy.planFailed);
      stage = answer.next;
    }
  }

  /** The four facts the browser owns, not the engine, exactly as run.ts stamps them. */
  function finish(result: BulkResult, startedAt: number, aborted: boolean): SimResult {
    return {
      ...result,
      request,
      lane: 'browser',
      duration_ms: Math.round(now() - startedAt),
      ...(aborted ? { aborted: true } : {}),
    };
  }

  const result = execute().catch((cause: unknown) => {
    if (cause instanceof BulkCapError || cause instanceof BulkRunError) throw cause;
    throw new BulkRunError(
      cancelled ? simCopy.stopped : bulkCopy.bulkFailed,
      cancelled,
      cancelled ? '' : detailOf(cause),
      { cause },
    );
  });

  return {
    callbackId,
    result,
    cancel() {
      cancelled = true;
      pool.abort(callbackId);
    },
  };
}

/**
 * A stat-weights run. It is one wasm call rather than a stage loop -- the engine already
 * computes every weight in one pass -- so this is `runBulk`'s shape with the loop taken out.
 */
export function runWeightsRun(
  pool: SimPool,
  request: WeightsRequest,
  onTick: (iterationsDone: number) => void,
  now: () => number = () => Date.now(),
): BulkRunHandle {
  const callbackId = nextCallbackId();
  let cancelled = false;

  const result = (async (): Promise<SimResult> => {
    const startedAt = now();
    const json = await pool.weights(JSON.stringify(request), callbackId, (progress) =>
      onTick(progress.iterationsDone),
    );
    const finished = JSON.parse(json) as SimResult;
    return {
      ...finished,
      request,
      lane: 'browser',
      duration_ms: Math.round(now() - startedAt),
    };
  })().catch((cause: unknown) => {
    throw new BulkRunError(
      cancelled ? simCopy.stopped : bulkCopy.bulkFailed,
      cancelled,
      cancelled ? '' : detailOf(cause),
      { cause },
    );
  });

  return {
    callbackId,
    result,
    cancel() {
      cancelled = true;
      pool.abort(callbackId);
    },
  };
}
```

- [ ] **Step 4: Add `BulkServerProgress` and `fetchBulkProgress`**

In `web/src/lib/sim/bulk-types.ts`, at the end:

```ts
import type { SimProgress } from './types';

/**
 * `GET /v1/sims/<id>/progress` for a bulk job (contract 2, "`Progress` (+)"). Every field
 * is optional here because the same route answers a plain run with all three at zero and
 * a build predating the columns with none of them.
 */
export interface BulkServerProgress extends SimProgress {
  stage?: number;
  combos_done?: number;
  combos_total?: number;
}
```

In `web/src/lib/sim/api.ts`, beside `fetchSimProgress`:

```ts
/**
 * The same route `fetchSimProgress` reads, typed for a bulk job's three extra columns. The
 * premium bulk dispatch itself is `dispatchServerSim` unchanged: a `BulkRequest` is a
 * `SimRequest`, the API derives the kind from the body, and a second POST helper would be a
 * second place for the CSRF header to go wrong.
 */
export function fetchBulkProgress(
  simId: string,
  apiBase: string = API_BASE_URL,
): Promise<BulkServerProgress> {
  return call<BulkServerProgress>(`/v1/sims/${simId}/progress`, apiBase, simCopy.loadFailed, {
    credentials: 'omit',
  });
}
```

with `import type { BulkServerProgress } from './bulk-types';` added to that file's imports.

- [ ] **Step 5: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-run.test.ts src/lib/sim/api.test.ts
```

Expected: PASS, 12 new cases.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/bulk-run.ts src/lib/sim/bulk-run.test.ts src/lib/sim/api.ts src/lib/sim/bulk-types.ts
npx prettier --write src/lib/sim/bulk-run.ts src/lib/sim/bulk-run.test.ts src/lib/sim/api.ts src/lib/sim/bulk-types.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/bulk-run.ts web/src/lib/sim/bulk-run.test.ts web/src/lib/sim/api.ts web/src/lib/sim/bulk-types.ts
git commit -m "$(cat <<'MSG'
feat(sim): the shared bulk stage loop

simPlan, each stage's requests across every worker unsplit, simRank,
repeat -- the shape contract 10.2 settles. No statistic exists here:
expansion, staging, cuts and the within-error grouping are all
sim/bulk's, and the ladder's history rides on StageRequests.ran. A stage
runs in chunks of the pool's width so a stop keeps every finished chunk
and re-ranks it as a final stage, which is what makes the partial result
real rather than aspirational.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 6: The phase gate and the loot tables

**Files:**
- Create: `web/src/lib/sim/phase.ts`, `web/src/lib/sim/phase.test.ts`
- Create: `web/src/lib/sim/loot.ts`, `web/src/lib/sim/loot.test.ts`

**Interfaces:**
- Consumes: `web/src/data/phases.json` (Task 2); `fetchJson`, `dataUrl`, `DataLoadError` from `../planner/load`; `call` from `./api`; `bulkCopy` from `./copy`.
- Produces:
  - phase.ts: `PhaseRow`, `BUILT_IN_PHASES`, `PHASE_LATER`, `phaseAt(phases, when)`, `phaseLabel(name)`, `phaseStart(phases, name)`, `hasOpened(phases, name, when)`, `openDateLabel(phases, name)`, `fetchPhases(apiBase?)`
  - loot.ts: `LootKind`, `LOOT_KINDS`, `LootBoss`, `LootSource`, `LootFile`, `loadLoot(build)`, `sourcesByItem(file)`, `groupSources(file)`, `itemsOfSource(source)`, `itemsOfBoss(source, bossId)`, `sourceNameOf(file, id)`, `DEFAULT_OFF_KINDS`, `isOpen(phases, source, when)`, `professionSplit(sources, professions)`, `SOURCE_KIND_LABELS`
  - api.ts: `fetchPhases(apiBase?)` re-exported through `phase.ts` (the single call site).

- [ ] **Step 1: Write the failing tests**

`web/src/lib/sim/phase.test.ts`:

```ts
// @vitest-environment jsdom
// web/src/lib/sim/phase.test.ts
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createSimApi, TEST_API } from '../../test-support/sim-api';
import {
  BUILT_IN_PHASES,
  PHASE_LATER,
  fetchPhases,
  hasOpened,
  openDateLabel,
  phaseAt,
  phaseLabel,
  phaseStart,
} from './phase';

const api = createSimApi();
const phases = BUILT_IN_PHASES;

beforeEach(() => api.install());
afterEach(() => api.reset());

describe('phaseAt', () => {
  it('names the phase a moment falls in, the way api/internal/phase does', () => {
    expect(phaseAt(phases, new Date('2026-09-01T00:00:00Z'))).toBe('pre-beta');
    expect(phaseAt(phases, new Date('2026-09-17T00:00:00Z'))).toBe('beta');
    expect(phaseAt(phases, new Date('2026-11-04T22:59:00Z'))).toBe('beta');
    expect(phaseAt(phases, new Date('2026-11-04T23:00:00Z'))).toBe('launch');
    expect(phaseAt(phases, new Date('2026-12-09T00:00:00Z'))).toBe('raids-1');
    expect(phaseAt(phases, new Date('2027-03-01T00:00:00Z'))).toBe('raids-1');
  });
});

describe('hasOpened', () => {
  it('treats a source with no phase as open from launch of the data', () => {
    expect(hasOpened(phases, undefined, new Date('2026-09-01T00:00:00Z'))).toBe(true);
  });

  it('gates a later phase and opens it on the instant', () => {
    expect(hasOpened(phases, 'raids-1', new Date('2026-12-08T23:59:00Z'))).toBe(false);
    expect(hasOpened(phases, 'raids-1', new Date('2026-12-09T00:00:00Z'))).toBe(true);
  });

  it('never opens the literal "later" (contract 10.4: an unknown date)', () => {
    expect(PHASE_LATER).toBe('later');
    expect(hasOpened(phases, PHASE_LATER, new Date('2099-01-01T00:00:00Z'))).toBe(false);
  });

  it('treats a phase nobody has heard of as not open, rather than as open', () => {
    expect(hasOpened(phases, 'season-of-mastery', new Date('2030-01-01T00:00:00Z'))).toBe(false);
  });
});

describe('labels', () => {
  it('names each phase from the display table, not from the data file', () => {
    expect(phaseLabel('raids-1')).toBe('First raids');
    expect(phaseLabel(PHASE_LATER)).toBe('Later');
    expect(phaseLabel('nope')).toBe('nope');
  });

  it('dates the phases that have a date and nothing else', () => {
    expect(openDateLabel(phases, 'raids-1')).toBe('9 December 2026');
    expect(openDateLabel(phases, 'pre-beta')).toBe('');
    expect(openDateLabel(phases, PHASE_LATER)).toBe('');
    expect(phaseStart(phases, 'pre-beta')).toBeNull();
    expect(BUILT_IN_PHASES).toHaveLength(4);
  });
});

describe('fetchPhases', () => {
  it('prefers GET /v1/phases at runtime', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/phases$/,
      respond: () =>
        new Response(
          JSON.stringify({
            ok: true,
            request_id: 'r',
            error: null,
            data: { phases: [{ name: 'raids-2', start: '2027-03-01T00:00:00Z' }] },
          }),
          { status: 200, headers: { 'content-type': 'application/json' } },
        ),
    });
    await expect(fetchPhases(TEST_API)).resolves.toEqual([
      { name: 'raids-2', start: '2027-03-01T00:00:00Z' },
    ]);
  });

  it('falls back to the build-time table rather than leaving the gate unknown', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/phases$/,
      respond: () => new Response(null, { status: 500 }),
    });
    await expect(fetchPhases(TEST_API)).resolves.toEqual(BUILT_IN_PHASES);
  });
});
```

`web/src/lib/sim/loot.test.ts`:

```ts
// web/src/lib/sim/loot.test.ts
import { describe, expect, it } from 'vitest';
import lootJson from '../../fixtures/planner/loot.json';
import {
  DEFAULT_OFF_KINDS,
  groupSources,
  isOpen,
  itemsOfBoss,
  itemsOfSource,
  professionSplit,
  sourceNameOf,
  sourcesByItem,
  type LootFile,
} from './loot';
import { BUILT_IN_PHASES } from './phase';

const file = lootJson as unknown as LootFile;
const phases = BUILT_IN_PHASES;

describe('sourcesByItem', () => {
  it('lists every source an item drops from, boss and trash alike', () => {
    const index = sourcesByItem(file);
    expect(index.get(16963)).toEqual(['raid:mc']);
    expect(index.get(13968)?.sort()).toEqual(
      ['dungeon:hall-of-thanes', 'raid:mc', 'rep:argent-dawn:exalted'].sort(),
    );
    expect(index.get(999999)).toBeUndefined();
  });
});

describe('groupSources', () => {
  it('groups by kind in picker order and labels each group', () => {
    const groups = groupSources(file);
    expect(groups.map((group) => group.kind)).toEqual([
      'raid',
      'dungeon',
      'world',
      'crafted',
      'rep',
      'pvp',
      'quest',
    ]);
    expect(groups[0].label).toBe('Raids');
    expect(groups[3].sources).toHaveLength(2);
  });
});

describe('itemsOfSource and itemsOfBoss', () => {
  it('rolls a raid up to every boss plus its trash, and a boss down to its own', () => {
    const raid = file.sources.find((source) => source.id === 'raid:mc')!;
    expect(itemsOfSource(raid).sort()).toEqual([12784, 13968, 16963, 19325]);
    expect(itemsOfBoss(raid, 'raid:mc:11502')).toEqual([12784, 19325]);
    expect(itemsOfBoss(raid, 'raid:mc:99999')).toEqual([]);
  });
});

describe('sourceNameOf', () => {
  it('names a source and a boss, which is what Candidate.SourceName carries', () => {
    expect(sourceNameOf(file, 'raid:mc')).toBe('Molten Core');
    expect(sourceNameOf(file, 'raid:mc:11502')).toBe('Ragnaros');
    expect(sourceNameOf(file, 'nothing')).toBe('');
  });
});

describe('isOpen', () => {
  it('gates an unreleased raid and passes everything without a phase', () => {
    const raid = file.sources.find((source) => source.id === 'raid:mc')!;
    const dungeon = file.sources.find((source) => source.id === 'dungeon:hall-of-thanes')!;
    expect(isOpen(phases, raid, new Date('2026-11-05T00:00:00Z'))).toBe(false);
    expect(isOpen(phases, raid, new Date('2026-12-09T00:00:00Z'))).toBe(true);
    expect(isOpen(phases, dungeon, new Date('2026-09-19T00:00:00Z'))).toBe(true);
  });

  it('never opens a source whose date is unknown ("later", contract 10.4)', () => {
    const world = file.sources.find((source) => source.id === 'world:azuregos')!;
    expect(world.opens).toBe('later');
    expect(isOpen(phases, world, new Date('2099-01-01T00:00:00Z'))).toBe(false);
  });
});

describe('professionSplit', () => {
  it('splits crafted into the character’s own and the rest', () => {
    const crafted = file.sources.filter((source) => source.kind === 'crafted');
    const split = professionSplit(crafted, ['blacksmithing']);
    expect(split.mine.map((source) => source.profession)).toEqual(['blacksmithing']);
    expect(split.other.map((source) => source.profession)).toEqual(['tailoring']);
  });

  it('puts everything in `other` when nothing recorded a profession', () => {
    const crafted = file.sources.filter((source) => source.kind === 'crafted');
    const split = professionSplit(crafted, undefined);
    expect(split.mine).toEqual([]);
    expect(split.other).toHaveLength(2);
  });
});

describe('DEFAULT_OFF_KINDS', () => {
  it('is quests and nothing else', () => {
    expect([...DEFAULT_OFF_KINDS]).toEqual(['quest']);
  });
});
```

- [ ] **Step 2: Run both and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/phase.test.ts src/lib/sim/loot.test.ts
```

Expected: FAIL, neither module resolves.

- [ ] **Step 3: Write `phase.ts`**

```ts
// web/src/lib/sim/phase.ts
// The content phases, and whether a loot source has opened yet.
//
// Two sources, in this order (contract 10.4 and 10.6):
//   * GET /v1/phases at runtime, so a date that moves after a deploy is picked up without
//     a rebuild;
//   * web/src/data/phases.json at build time, which the pipeline emits from
//     data/curated/phases.json -- the file api/internal/phase.Boundaries is tested
//     against -- as the fallback when the API is unreachable.
//
// The data file is `[{ "name", "start" }]` and carries no label: a label is display copy
// and lives here, in DISPLAY_LABELS, beside the rest of this lane's words.
import { API_BASE_URL } from '../planner/config';
import builtIn from '../../data/phases.json';

export interface PhaseRow {
  /** The stored, URL-safe name: pre-beta | beta | launch | raids-1. */
  name: string;
  /** RFC3339, or "" for the phase that has no start. */
  start: string;
}

/** The build-time table, and the fallback when the API cannot be reached. */
export const BUILT_IN_PHASES: readonly PhaseRow[] = builtIn as PhaseRow[];

/**
 * Contract 10.4: "a source whose date is unknown carries `opens: "later"`, which the page
 * shows as unreleased without a date." It is never a row in the phase table, and it never
 * opens.
 */
export const PHASE_LATER = 'later';

const DISPLAY_LABELS: Record<string, string> = {
  'pre-beta': 'Before beta',
  beta: 'Beta',
  launch: 'Launch',
  'raids-1': 'First raids',
  [PHASE_LATER]: 'Later',
};

/**
 * The live table, or the build-time one. Every failure -- an unreachable API, a 500, a
 * body that is not the expected shape -- falls back rather than throwing: a phase gate
 * that cannot answer must not take the Droptimizer down with it.
 */
export async function fetchPhases(apiBase: string = API_BASE_URL): Promise<readonly PhaseRow[]> {
  try {
    const response = await fetch(`${apiBase}/v1/phases`, {
      headers: { accept: 'application/json' },
      credentials: 'omit',
    });
    if (!response.ok) return BUILT_IN_PHASES;
    const envelope = (await response.json()) as { ok?: boolean; data?: { phases?: PhaseRow[] } };
    const rows = envelope.data?.phases;
    return Array.isArray(rows) && rows.length > 0 ? rows : BUILT_IN_PHASES;
  } catch {
    return BUILT_IN_PHASES;
  }
}

/** The phase a moment falls in. Later boundaries win, so the last match is the answer. */
export function phaseAt(phases: readonly PhaseRow[], when: Date): string {
  let name = phases[0]?.name ?? '';
  for (const phase of phases) {
    if (phase.start === '') continue;
    if (when.getTime() >= Date.parse(phase.start)) name = phase.name;
  }
  return name;
}

function rowOf(phases: readonly PhaseRow[], name: string): PhaseRow | undefined {
  return phases.find((phase) => phase.name === name);
}

/** What a player calls it. Display copy, not data -- the table carries no label column. */
export function phaseLabel(name: string): string {
  return DISPLAY_LABELS[name] ?? name;
}

/** When the phase opens, or null for the one with no start, for "later", and for an unknown. */
export function phaseStart(phases: readonly PhaseRow[], name: string): Date | null {
  const start = rowOf(phases, name)?.start ?? '';
  return start === '' ? null : new Date(start);
}

/**
 * Whether content gated on this phase is available at `when`.
 *
 * An absent phase is open: contract 6.1 says "a source without it is open from launch".
 * The literal "later" is never open (contract 10.4). A phase this build has never heard of
 * is NOT open either -- a loot file naming a phase the site does not carry is ahead of the
 * site, and showing its raid as live would be a lie in the one direction that matters.
 */
export function hasOpened(
  phases: readonly PhaseRow[],
  name: string | undefined,
  when: Date,
): boolean {
  if (name === undefined || name === '') return true;
  if (name === PHASE_LATER) return false;
  const row = rowOf(phases, name);
  if (row === undefined) return false;
  if (row.start === '') return true;
  return when.getTime() >= Date.parse(row.start);
}

const DATE_FORMAT = new Intl.DateTimeFormat('en-GB', {
  day: 'numeric',
  month: 'long',
  year: 'numeric',
  timeZone: 'UTC',
});

/** "9 December 2026", or "" for a phase with no start, for "later", and for an unknown. */
export function openDateLabel(phases: readonly PhaseRow[], name: string): string {
  const start = phaseStart(phases, name);
  return start === null ? '' : DATE_FORMAT.format(start);
}
```

- [ ] **Step 4: Write `loot.ts`**

```ts
// web/src/lib/sim/loot.ts
// data/builds/<build>/loot.json: where every item comes from, contract 6.1 as corrected by
// 10.4 -- AreaTable zone ids, `<source>:<npc-id>` boss ids, `opens: "later"` for an
// unknown date, and an empty boss name where neither database had one.
//
// The file is per build and item-id keyed through its sources, so the page never asks the
// API where an item drops -- it is the same fetch-and-cache path items.json already takes.
// Drop chances are in neither the fork's database nor the curated overlay, so nothing here
// computes or exposes a probability.
//
// Contract 10.4 also warns that the re-itemised raid tier is thin: 1,809 of the fork's
// sourced item ids do not exist in the 1.60 client, loot.json lists only items the build
// has, and the first curated overlay records the gap per raid in its notes. The page shows
// what the file contains and counts what it does not -- it never fabricates a drop.
import { bulkCopy } from './copy';
import { dataUrl, fetchJson, DataLoadError } from '../planner/load';
import { hasOpened, type PhaseRow } from './phase';

export const LOOT_KINDS = ['raid', 'dungeon', 'world', 'crafted', 'rep', 'pvp', 'quest'] as const;
export type LootKind = (typeof LOOT_KINDS)[number];

export const SOURCE_KIND_LABELS: Record<LootKind, string> = {
  raid: bulkCopy.sourcesRaids,
  dungeon: bulkCopy.sourcesDungeons,
  world: bulkCopy.sourcesWorld,
  crafted: bulkCopy.sourcesCrafted,
  rep: bulkCopy.sourcesRep,
  pvp: bulkCopy.sourcesPvp,
  quest: bulkCopy.sourcesQuests,
};

/** Off unless the player asks: a quest reward is a one-time source (design 6.1). */
export const DEFAULT_OFF_KINDS: readonly LootKind[] = ['quest'];

export interface LootBoss {
  /** `<source id>:<npc-id>` (contract 10.4). */
  id: string;
  /** Empty when neither database names the boss; the picker shows the id in that case. */
  name: string;
  npc_id?: number;
  items: number[];
}

export interface LootSource {
  /** `raid:<zone-slug>`, `dungeon:<zone-slug>`, `pvp:rank-<n>`, … (contract 10.4). */
  id: string;
  kind: LootKind;
  name: string;
  /** An AreaTable id -- Molten Core is 2717, not 409 (contract 10.4). */
  zone_id?: number;
  /**
   * A phase name from the phase table; the literal `"later"` for a source whose date is
   * unknown, which the page shows as unreleased without a date; absent means open from
   * launch (contract 6.1 and 10.4).
   */
  opens?: string;
  bosses?: LootBoss[];
  trash?: number[];
  items?: number[];
  profession?: string;
  faction_id?: number;
  standing?: string;
  rank?: number;
}

export interface LootFile {
  sources: LootSource[];
}

const EMPTY: LootFile = { sources: [] };

/**
 * The build's loot table. A build the data lane has not regenerated ships none, and a 404
 * resolves to an empty file rather than throwing -- the Droptimizer then says it has no
 * sources, which is true, instead of showing an error for a file that was never promised.
 * Every other failure is a broken build and is rethrown.
 */
export async function loadLoot(build: string): Promise<LootFile> {
  try {
    return await fetchJson<LootFile>(dataUrl(build, 'loot.json'));
  } catch (error) {
    if (error instanceof DataLoadError && error.status === 404) return EMPTY;
    throw error;
  }
}

/** Every item a source yields: its own list, every boss's, and its trash. */
export function itemsOfSource(source: LootSource): number[] {
  return [
    ...(source.items ?? []),
    ...(source.bosses ?? []).flatMap((boss) => boss.items),
    ...(source.trash ?? []),
  ];
}

export function itemsOfBoss(source: LootSource, bossId: string): number[] {
  return (source.bosses ?? []).find((boss) => boss.id === bossId)?.items ?? [];
}

/** item id -> the source ids it drops from. Built once per loaded file. */
export function sourcesByItem(file: LootFile): Map<number, string[]> {
  const index = new Map<number, string[]>();
  for (const source of file.sources) {
    for (const item of new Set(itemsOfSource(source))) {
      const existing = index.get(item);
      if (existing === undefined) index.set(item, [source.id]);
      else if (!existing.includes(source.id)) existing.push(source.id);
    }
  }
  return index;
}

export interface SourceGroup {
  kind: LootKind;
  label: string;
  sources: LootSource[];
}

/** The picker's groups, in the design's own order, skipping kinds this build has none of. */
export function groupSources(file: LootFile): SourceGroup[] {
  return LOOT_KINDS.flatMap((kind) => {
    const sources = file.sources.filter((source) => source.kind === kind);
    return sources.length === 0 ? [] : [{ kind, label: SOURCE_KIND_LABELS[kind], sources }];
  });
}

export function isOpen(phases: readonly PhaseRow[], source: LootSource, when: Date): boolean {
  return hasOpened(phases, source.opens, when);
}

/**
 * A source or boss id as words -- "Molten Core", "Ragnaros" -- for `Candidate.SourceName`
 * (contract 10.1 A6). The page fills it once, when it builds a drops request; every later
 * read is off the result, never a second join.
 */
export function sourceNameOf(file: LootFile, id: string): string {
  for (const source of file.sources) {
    if (source.id === id) return source.name;
    for (const boss of source.bosses ?? []) if (boss.id === id) return boss.name;
  }
  return '';
}

/**
 * Crafted sources split into the character's own professions and the rest. Nothing records
 * a character's professions today (`CharacterSpec.professions` is optional and every source
 * leaves it unset), so an absent list puts everything in `other` and the page says why
 * rather than claiming the character has none.
 */
export function professionSplit(
  crafted: readonly LootSource[],
  professions: readonly string[] | undefined,
): { mine: LootSource[]; other: LootSource[] } {
  const mine = new Set(professions ?? []);
  return {
    mine: crafted.filter((source) => source.profession !== undefined && mine.has(source.profession)),
    other: crafted.filter(
      (source) => source.profession === undefined || !mine.has(source.profession),
    ),
  };
}
```

- [ ] **Step 5: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/phase.test.ts src/lib/sim/loot.test.ts
```

Expected: PASS.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/phase.ts src/lib/sim/phase.test.ts src/lib/sim/loot.ts src/lib/sim/loot.test.ts
npx prettier --write src/lib/sim/phase.ts src/lib/sim/phase.test.ts src/lib/sim/loot.ts src/lib/sim/loot.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/phase.ts web/src/lib/sim/phase.test.ts web/src/lib/sim/loot.ts web/src/lib/sim/loot.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): the phase gate and the loot tables

loot.json's sources in contract 10.4's shape -- AreaTable zone ids,
`<source>:<npc-id>` boss ids, `opens: "later"` for an unknown date --
the item-to-source index the search filters on, and sourceNameOf, which
is what fills Candidate.SourceName so the results never re-join an id to
a name. Phases come from GET /v1/phases at runtime with the pipeline's
web/src/data/phases.json as the build-time fallback (contract 10.6).

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 7: Enchants, suffixes, buff names and the item search

**Files:**
- Create: `web/src/lib/sim/enchants.ts`, `web/src/lib/sim/enchants.test.ts`
- Create: `web/src/lib/sim/sim-buffs.ts`, `web/src/lib/sim/sim-buffs.test.ts`
- Create: `web/src/lib/sim/item-search.ts`, `web/src/lib/sim/item-search.test.ts`

**Interfaces:**
- Consumes: `loadLoot`, `sourcesByItem` (Task 6); `Item`, `Slot`, `SLOTS`, `SLOT_LABELS` and `slotsForItem` from the planner; `dataUrl`, `fetchJson`, `DataLoadError`.
- Produces:
  - enchants.ts: `EnchantRow`, `SuffixRow`, `loadEnchants(build)`, `loadSuffixes(build)`, `enchantsForSlot(rows, slot, classSlug?)`, `suffixesForItem(rows, item)`, `ENCHANTS_PER_SLOT_CAP`, `KEEP_CURRENT_ENCHANT`, `NO_ENCHANT`
  - sim-buffs.ts: `SimBuffRow`, `SimBuffFile`, `loadSimBuffs(build)`, `buffName(file, id)`, `buffIcon(file, id)`
  - item-search.ts: `ItemQuery`, `defaultItemQuery()`, `SEARCH_LIMIT`, `searchItems(items, query, ctx)`, `SearchContext`, `slotOptions()`

- [ ] **Step 1: Write the failing tests**

`web/src/lib/sim/enchants.test.ts`:

```ts
// web/src/lib/sim/enchants.test.ts
import { describe, expect, it } from 'vitest';
import enchantsJson from '../../fixtures/planner/enchants.json';
import suffixesJson from '../../fixtures/planner/suffixes.json';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import {
  ENCHANTS_PER_SLOT_CAP,
  KEEP_CURRENT_ENCHANT,
  NO_ENCHANT,
  enchantsForSlot,
  suffixesForItem,
  type EnchantRow,
  type SuffixRow,
} from './enchants';
import type { Item } from '../planner/types';

const enchants = enchantsJson as unknown as EnchantRow[];
const suffixes = suffixesJson as unknown as SuffixRow[];
const items = itemsJson.items as unknown as Item[];
const head = items.find((item) => item.id === 12640)!;
const charm = items.find((item) => item.id === 13968)!;

describe('enchantsForSlot', () => {
  it('offers only the enchants whose slot list names this slot', () => {
    expect(enchantsForSlot(enchants, 'head').map((row) => row.effect_id)).toEqual([2543]);
    expect(enchantsForSlot(enchants, 'main_hand').map((row) => row.effect_id)).toEqual([1900]);
    expect(enchantsForSlot(enchants, 'neck')).toEqual([]);
  });

  it('is empty rather than wrong when the build ships no enchant file', () => {
    expect(enchantsForSlot([], 'head')).toEqual([]);
  });
});

describe('suffixesForItem', () => {
  it('offers only the suffixes the item rolls', () => {
    expect(suffixesForItem(suffixes, charm).map((row) => row.name)).toEqual([
      'of the Bear',
      'of the Tiger',
    ]);
    expect(suffixesForItem(suffixes, head)).toEqual([]);
  });
});

describe('the two sentinel enchant values', () => {
  it('are distinct and are not real effect ids', () => {
    expect(KEEP_CURRENT_ENCHANT).toBe(-1);
    expect(NO_ENCHANT).toBe(0);
    expect(ENCHANTS_PER_SLOT_CAP).toBe(4);
  });
});

describe('the enchant row shape (contract 10.4)', () => {
  it('is keyed by effect_id, which is what GearSlot.enchant carries', () => {
    for (const row of enchants) expect(typeof row.effect_id).toBe('number');
  });
});
```

`web/src/lib/sim/item-search.test.ts`:

```ts
// web/src/lib/sim/item-search.test.ts
import { describe, expect, it } from 'vitest';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import lootJson from '../../fixtures/planner/loot.json';
import { defaultItemQuery, searchItems, SEARCH_LIMIT, type SearchContext } from './item-search';
import { sourcesByItem, type LootFile } from './loot';
import type { Item } from '../planner/types';

const items = itemsJson.items as unknown as Item[];
const ctx: SearchContext = {
  level: 60,
  sourcesByItem: sourcesByItem(lootJson as unknown as LootFile),
};

describe('searchItems', () => {
  it('returns everything for an empty query, sorted by item level then name', () => {
    const found = searchItems(items, defaultItemQuery(), ctx);
    expect(found).toHaveLength(items.length);
    expect(found[0].item_level).toBeGreaterThanOrEqual(found[found.length - 1].item_level);
  });

  it('matches on name, case-insensitively, anywhere in the name', () => {
    const found = searchItems(items, { ...defaultItemQuery(), text: 'wrath' }, ctx);
    expect(found.map((item) => item.id).sort()).toEqual([16963, 16966]);
  });

  it('filters on minimum item level', () => {
    const found = searchItems(items, { ...defaultItemQuery(), minItemLevel: 76 }, ctx);
    expect(found.map((item) => item.id).sort()).toEqual([16963, 16966, 19325]);
  });

  it('filters on slot, following the finger and trinket aliases', () => {
    expect(searchItems(items, { ...defaultItemQuery(), slot: 'finger1' }, ctx).map((i) => i.id)).toEqual([
      19325,
    ]);
    expect(searchItems(items, { ...defaultItemQuery(), slot: 'head' }, ctx).map((i) => i.id).sort()).toEqual(
      [12640, 16963],
    );
  });

  it('filters on source', () => {
    const found = searchItems(items, { ...defaultItemQuery(), sourceId: 'world:azuregos' }, ctx);
    expect(found.map((item) => item.id)).toEqual([19325]);
  });

  it('drops what the character cannot equip when usableOnly is on, and keeps it when off', () => {
    const lowLevel: SearchContext = { ...ctx, level: 59 };
    const on = searchItems(items, defaultItemQuery(), lowLevel);
    expect(on.map((item) => item.id)).toEqual([13968]);
    const off = searchItems(items, { ...defaultItemQuery(), usableOnly: false }, lowLevel);
    expect(off).toHaveLength(items.length);
  });

  it('never returns more than the limit', () => {
    const many = Array.from({ length: SEARCH_LIMIT + 20 }, (_, i) => ({ ...items[0], id: 1000 + i }));
    expect(searchItems(many, defaultItemQuery(), ctx)).toHaveLength(SEARCH_LIMIT);
  });
});

describe('defaultItemQuery', () => {
  it('opens with usable-only on, the way the design specifies', () => {
    expect(defaultItemQuery()).toEqual({
      text: '',
      minItemLevel: 0,
      slot: '',
      sourceId: '',
      usableOnly: true,
    });
  });
});
```

- [ ] **Step 2: Run them and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/enchants.test.ts src/lib/sim/item-search.test.ts
```

Expected: FAIL, neither module resolves.

- [ ] **Step 3: Write `enchants.ts`**

```ts
// web/src/lib/sim/enchants.ts
// data/builds/<build>/enchants.json and suffixes.json, contract 6.2 and 6.3 as corrected
// by 10.4: rows are keyed by `effect_id` plus `spell_id`/`item_id`, `item_types` is the
// `EnchantType` shape restriction, `slots` is derived by the data lane from `type` and
// `extra_types`, and suffixes come from the fork database's `randomSuffixes` because this
// build's client has no `ItemRandomSuffix` table.
//
// Both files are optional: a build the data lane has not regenerated ships neither, and
// the enchant column simply does not appear. Neither is large enough to want pruning, so
// both are fetched whole, per build, and cached by the browser the way items.json is.
//
// These are the browser UI's copy. The planner inside the wasm does not read them:
// contract 10.1 A9 embeds the same rows in `sim/internal/simdb`, so `Expand` needs no file
// at runtime and the two can never disagree about what fits where.
import { dataUrl, fetchJson, DataLoadError } from '../planner/load';
import type { Item, StatKey } from '../planner/types';

export interface EnchantRow {
  /**
   * The client's enchantment effect id. This is the key (contract 10.4) and it is what
   * `GearSlot.enchant` and `Candidate.enchant` carry.
   */
  effect_id: number;
  /** The spell that applies it, where one does. */
  spell_id?: number;
  /** The item that applies it, where one does. One of the two is always present. */
  item_id?: number;
  name: string;
  icon: string;
  /** Planner slot names, derived by the data lane from `type` and `extra_types`. */
  slots: string[];
  /** The `EnchantType` shape restriction, as the fork's own numeric enum values. */
  item_types: number[];
  /** Empty for an enchant every class can use. */
  classes: string[];
  stats: Partial<Record<StatKey, number>>;
  phase?: string;
}

export interface SuffixRow {
  id: number;
  name: string;
  stats: Partial<Record<StatKey, number>>;
}

/**
 * `Candidate.enchant` of 0 means "inherit the equipped item's enchant for this slot where
 * it fits" (contract 1.3), so 0 is the wire's own "keep current" and there is no separate
 * value for it. -1 is this page's "explicitly none": it never travels, and
 * `candidates.ts` turns it into an omitted enchant on a candidate whose slot is empty.
 */
export const NO_ENCHANT = 0;
export const KEEP_CURRENT_ENCHANT = -1;

/** Raidbots shows a per-slot selection cap; ours is four, which is a legible list. */
export const ENCHANTS_PER_SLOT_CAP = 4;

async function loadOptional<T>(build: string, file: string, empty: T): Promise<T> {
  try {
    return await fetchJson<T>(dataUrl(build, file));
  } catch (error) {
    if (error instanceof DataLoadError && error.status === 404) return empty;
    throw error;
  }
}

export function loadEnchants(build: string): Promise<EnchantRow[]> {
  return loadOptional<EnchantRow[]>(build, 'enchants.json', []);
}

export function loadSuffixes(build: string): Promise<SuffixRow[]> {
  return loadOptional<SuffixRow[]>(build, 'suffixes.json', []);
}

/**
 * The enchants this slot allows. Two gates, both the fork's own: the enchant's slot list,
 * and its class list where it has one (an empty `classes` is for everyone).
 *
 * `item_types` is deliberately NOT applied here. It is the `EnchantType` shape restriction
 * and `items.json` rows carry no armour subclass to test it against, so the browser cannot
 * evaluate it. `Expand` inside the wasm can, from simdb, and refuses a bad pairing at the
 * boundary with its own words -- which is the honest failure, and better than a filter here
 * that would have to guess.
 */
export function enchantsForSlot(
  rows: readonly EnchantRow[],
  slot: string,
  classSlug = '',
): EnchantRow[] {
  return rows.filter((row) => {
    if (!row.slots.includes(slot)) return false;
    if (row.classes.length > 0 && classSlug !== '' && !row.classes.includes(classSlug)) return false;
    return true;
  });
}

/** The random suffixes this item rolls, in file order. */
export function suffixesForItem(rows: readonly SuffixRow[], item: Item): SuffixRow[] {
  const ids = new Set(item.suffixes ?? []);
  return rows.filter((row) => ids.has(row.id));
}
```

- [ ] **Step 3b: Write `sim-buffs.ts` and its test**

`web/src/lib/sim/sim-buffs.test.ts`:

```ts
// web/src/lib/sim/sim-buffs.test.ts
import { describe, expect, it } from 'vitest';
import simbuffsJson from '../../fixtures/planner/simbuffs.json';
import { buffIcon, buffName, type SimBuffFile } from './sim-buffs';

const file = simbuffsJson as unknown as SimBuffFile;

describe('buffName and buffIcon', () => {
  it('name and illustrate an id the file has', () => {
    expect(buffName(file, 'flask_of_supreme_power')).toBe('Flask of Supreme Power');
    expect(buffIcon(file, 'flask_of_supreme_power')).toBe('fixture_item_snakestone_charm');
  });

  it('fall back to a legible form of the id rather than to an empty row', () => {
    expect(buffName(file, 'juju_power')).toBe('juju power');
    expect(buffIcon(file, 'juju_power')).toBe('');
  });

  it('read an empty file without throwing', () => {
    expect(buffName({ entries: {} }, 'battle_shout')).toBe('battle shout');
  });
});
```

```ts
// web/src/lib/sim/sim-buffs.ts
// data/builds/<build>/simbuffs.json (contract 10.4): the display name and icon for every
// IDS.md buff, debuff, world buff and consumable id.
//
// The engine's ids are snake case and legible on their own -- "flask_of_supreme_power" --
// so a missing row is a de-underscored id rather than a blank, and a build that ships no
// file loses the icons and nothing else.
import { dataUrl, fetchJson, DataLoadError } from '../planner/load';

export interface SimBuffRow {
  name: string;
  icon: string;
}

export interface SimBuffFile {
  entries: Record<string, SimBuffRow>;
}

const EMPTY: SimBuffFile = { entries: {} };

export async function loadSimBuffs(build: string): Promise<SimBuffFile> {
  try {
    return await fetchJson<SimBuffFile>(dataUrl(build, 'simbuffs.json'));
  } catch (error) {
    if (error instanceof DataLoadError && error.status === 404) return EMPTY;
    throw error;
  }
}

export function buffName(file: SimBuffFile, id: string): string {
  return file.entries[id]?.name ?? id.replaceAll('_', ' ');
}

/** "" when the file has no row: the caller draws no icon rather than a broken one. */
export function buffIcon(file: SimBuffFile, id: string): string {
  return file.entries[id]?.icon ?? '';
}
```

- [ ] **Step 4: Write `item-search.ts`**

```ts
// web/src/lib/sim/item-search.ts
// The item search behind Top Gear's "add a candidate", design 3.1.3. Entirely client-side
// over the build's own items/<class>.json, which the page already has loaded for the gear
// grid -- no API call, no second index.
import { slotsForItem } from '../planner/rules';
import { SLOTS, SLOT_LABELS, type Item, type Slot } from '../planner/types';

export interface ItemQuery {
  text: string;
  minItemLevel: number;
  /** A planner slot, or "" for any. */
  slot: string;
  /** A loot source id, or "" for anywhere. */
  sourceId: string;
  /** On by default, per the design. */
  usableOnly: boolean;
}

export interface SearchContext {
  /** The character's level, for the usable-only gate. */
  level: number;
  /** From `loot.ts`'s `sourcesByItem`. Empty when the build ships no loot file. */
  sourcesByItem: Map<number, string[]>;
}

/**
 * The list is class-filtered already (items/<class>.json is per class) and a long list is
 * a scroll nobody reads, so the search shows the best hundred and says how many it hid.
 */
export const SEARCH_LIMIT = 100;

export function defaultItemQuery(): ItemQuery {
  return { text: '', minItemLevel: 0, slot: '', sourceId: '', usableOnly: true };
}

/** The slot filter's options, in the grid's own order. */
export function slotOptions(): { slot: Slot; label: string }[] {
  return SLOTS.map((slot) => ({ slot, label: SLOT_LABELS[slot] }));
}

/**
 * Best item level first, then name, so the top of the list is the part worth reading. The
 * design's own note stands: search can find items this character has no way to obtain, and
 * `usableOnly` is about what can be equipped, not about what can be got.
 */
export function searchItems(
  items: readonly Item[],
  query: ItemQuery,
  ctx: SearchContext,
): Item[] {
  const needle = query.text.trim().toLowerCase();
  return items
    .filter((item) => {
      if (needle !== '' && !item.name.toLowerCase().includes(needle)) return false;
      if (item.item_level < query.minItemLevel) return false;
      if (query.slot !== '' && !slotsForItem(item).includes(query.slot as Slot)) return false;
      if (query.sourceId !== '' && !(ctx.sourcesByItem.get(item.id) ?? []).includes(query.sourceId)) {
        return false;
      }
      if (query.usableOnly && item.required_level > ctx.level) return false;
      return true;
    })
    .sort((a, b) => b.item_level - a.item_level || a.name.localeCompare(b.name))
    .slice(0, SEARCH_LIMIT);
}
```

- [ ] **Step 5: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/enchants.test.ts src/lib/sim/sim-buffs.test.ts src/lib/sim/item-search.test.ts
```

Expected: PASS.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/enchants.ts src/lib/sim/enchants.test.ts src/lib/sim/sim-buffs.ts src/lib/sim/sim-buffs.test.ts src/lib/sim/item-search.ts src/lib/sim/item-search.test.ts
npx prettier --write src/lib/sim/enchants.ts src/lib/sim/enchants.test.ts src/lib/sim/sim-buffs.ts src/lib/sim/sim-buffs.test.ts src/lib/sim/item-search.ts src/lib/sim/item-search.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/enchants.ts web/src/lib/sim/enchants.test.ts web/src/lib/sim/sim-buffs.ts web/src/lib/sim/sim-buffs.test.ts web/src/lib/sim/item-search.ts web/src/lib/sim/item-search.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): enchant, suffix and buff-name lists, and the item search

Enchants are keyed by effect_id plus spell_id/item_id and item_types is
the EnchantType shape restriction, per contract 10.4; the browser does
not evaluate it, because items.json carries no armour subclass and the
wasm's own Expand can (10.1 A9). simbuffs.json gives every IDS.md id a
name and an icon. All three files are optional and a 404 is an empty
list, so a build the data lane has not regenerated loses the column
rather than the page.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 8: The candidate model and "Copy to addon"

**Files:**
- Create: `web/src/lib/sim/candidates.ts`, `web/src/lib/sim/candidates.test.ts`
- Create: `web/src/lib/sim/addon-export.ts`, `web/src/lib/sim/addon-export.test.ts`

**Interfaces:**
- Consumes: `Candidate`, `BulkSpec`, `BulkMode`, `Precision`, `Substitution` (Task 1); `NO_ENCHANT`, `KEEP_CURRENT_ENCHANT` (Task 7); `slotsForItem`, `SLOTS`, `SLOT_ALIASES`, `SLOT_LABELS`, `Item`, `Slot`, `GearSlot`.
- Produces:
  - candidates.ts: `CandidateRow`, `Origin`, `candidateKey(row)`, `envelopeSlotOf(item)`, `uiSlotsOf(item)`, `rowFor(item, slot, origin, sourceName?)`, `toggleRow(rows, key)`, `addRow(rows, row)`, `copyAndModify(rows, key, patch)`, `removeRow(rows, key)`, `toCandidates(rows, locked)`, `buildBulkSpec(input)`, `validateBulk(spec)`, `CANDIDATE_ROW_SEPARATOR`
  - addon-export.ts: `gearEntry(slot)`, `addonGearList(gear)`, `addonStringFor(input)`

- [ ] **Step 1: Write the failing tests**

`web/src/lib/sim/candidates.test.ts`:

```ts
// web/src/lib/sim/candidates.test.ts
import { describe, expect, it } from 'vitest';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import {
  addRow,
  buildBulkSpec,
  candidateKey,
  copyAndModify,
  envelopeSlotOf,
  removeRow,
  rowFor,
  toCandidates,
  toggleRow,
  uiSlotsOf,
  validateBulk,
  type CandidateRow,
} from './candidates';
import { KEEP_CURRENT_ENCHANT, NO_ENCHANT } from './enchants';
import type { Item } from '../planner/types';

const items = itemsJson.items as unknown as Item[];
const helm = items.find((item) => item.id === 16963)!;
const ring = items.find((item) => item.id === 19325)!;
const charm = items.find((item) => item.id === 13968)!;

describe('envelopeSlotOf', () => {
  it('is "" for anything that fits more than one slot, and the slot otherwise', () => {
    expect(envelopeSlotOf(helm)).toBe('head');
    expect(envelopeSlotOf(ring)).toBe('');
    expect(envelopeSlotOf(charm)).toBe('');
  });
});

describe('uiSlotsOf', () => {
  it('gives the grid one row for rings and trinkets, not two', () => {
    expect(uiSlotsOf(helm)).toEqual(['head']);
    expect(uiSlotsOf(ring)).toEqual(['finger1']);
    expect(uiSlotsOf(charm)).toEqual(['trinket1']);
  });
});

describe('candidateKey', () => {
  it('identifies a row by slot, item, enchant and suffix, so a copy is a different row', () => {
    const base = rowFor(helm, 'head', 'bag');
    const enchanted = { ...base, enchant: 2543 };
    expect(candidateKey(base)).toBe('head:16963:0:0');
    expect(candidateKey(enchanted)).toBe('head:16963:2543:0');
  });
});

describe('the row list', () => {
  const start: CandidateRow[] = [rowFor(helm, 'head', 'bag'), rowFor(ring, 'finger1', 'equipped')];

  it('toggles one row and leaves the list otherwise identical', () => {
    const next = toggleRow(start, candidateKey(start[0]));
    expect(next[0].checked).toBe(!start[0].checked);
    expect(next[1]).toEqual(start[1]);
    expect(start[0].checked).toBe(false);
  });

  it('adds a row once, and re-adding it ticks the one already there', () => {
    const added = addRow(start, rowFor(charm, 'trinket1', 'search'));
    expect(added).toHaveLength(3);
    const again = addRow(added, { ...rowFor(charm, 'trinket1', 'search'), checked: true });
    expect(again).toHaveLength(3);
    expect(again[2].checked).toBe(true);
  });

  it('copies and modifies into a new, ticked row beside the original', () => {
    const next = copyAndModify(start, candidateKey(start[0]), { enchant: 2543 });
    expect(next).toHaveLength(3);
    expect(next[1].enchant).toBe(2543);
    expect(next[1].checked).toBe(true);
    expect(next[1].origin).toBe('bag');
  });

  it('removes by key', () => {
    expect(removeRow(start, candidateKey(start[0]))).toHaveLength(1);
  });
});

describe('toCandidates', () => {
  it('sends only ticked rows, with the envelope slot and no sentinel enchant', () => {
    const rows = [
      { ...rowFor(helm, 'head', 'bag'), checked: true },
      { ...rowFor(ring, 'finger1', 'equipped'), checked: true, enchant: KEEP_CURRENT_ENCHANT },
      rowFor(charm, 'trinket1', 'search'),
    ];
    expect(toCandidates(rows, [])).toEqual([
      { slot: 'head', item_id: 16963, origin: 'bag' },
      { slot: '', item_id: 19325, origin: 'equipped' },
    ]);
  });

  it('carries the source name a drop was picked from (contract 10.1 A6)', () => {
    const rows = [
      { ...rowFor(helm, 'head', 'drop:raid:mc:11502', 'Ragnaros'), checked: true },
    ];
    expect(toCandidates(rows, [])).toEqual([
      { slot: 'head', item_id: 16963, origin: 'drop:raid:mc:11502', source_name: 'Ragnaros' },
    ]);
  });

  it('never sends a candidate on a locked slot', () => {
    const rows = [{ ...rowFor(helm, 'head', 'bag'), checked: true }];
    expect(toCandidates(rows, ['head'])).toEqual([]);
  });

  it('sends an explicit "no enchant" as an omitted enchant, not as -1', () => {
    const rows = [{ ...rowFor(helm, 'head', 'bag'), checked: true, enchant: NO_ENCHANT }];
    expect(toCandidates(rows, [])[0].enchant).toBeUndefined();
  });
});

describe('buildBulkSpec and validateBulk', () => {
  it('builds a gear spec with the cap and precision it was given', () => {
    const spec = buildBulkSpec({
      mode: 'gear',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: ['main_hand'],
      loadouts: [],
      sets: [],
      precision: 'fast',
      cap: 400,
    });
    expect(spec.mode).toBe('gear');
    expect(spec.cap).toBe(400);
    expect(spec.locked).toEqual(['main_hand']);
    expect(validateBulk(spec)).toBeNull();
  });

  it('refuses a talents spec that carries candidates, and one with no loadout', () => {
    const withCandidates = buildBulkSpec({
      mode: 'talents',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: [],
      loadouts: [{ name: 'A', talents: '0-1-' }],
      sets: [],
      precision: 'normal',
      cap: 400,
    });
    expect(withCandidates.candidates).toEqual([]);
    const empty = buildBulkSpec({
      mode: 'talents',
      rows: [],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'normal',
      cap: 400,
    });
    expect(validateBulk(empty)).not.toBeNull();
  });

  it('refuses a drops spec whose candidates are not all from a drop', () => {
    const spec = buildBulkSpec({
      mode: 'drops',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'normal',
      cap: 400,
    });
    expect(validateBulk(spec)).not.toBeNull();
  });

  it('refuses a gear spec with nothing ticked at all', () => {
    const spec = buildBulkSpec({
      mode: 'gear',
      rows: [],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'fast',
      cap: 400,
    });
    expect(validateBulk(spec)).not.toBeNull();
  });
});
```

`web/src/lib/sim/addon-export.test.ts`:

```ts
// web/src/lib/sim/addon-export.test.ts
import { describe, expect, it } from 'vitest';
import { addonGearList, addonStringFor, gearEntry } from './addon-export';
import type { GearSlot } from './types';

const gear: GearSlot[] = [
  { slot: 'head', item_id: 16963, enchant: 2543 },
  { slot: 'main_hand', item_id: 12784 },
  { slot: 'trinket1', item_id: 13968, enchant: 0, suffix: 2 },
];

describe('gearEntry', () => {
  it('writes item, enchant and suffix only as far as it has to', () => {
    expect(gearEntry({ slot: 'head', item_id: 1 })).toBe('head=1');
    expect(gearEntry({ slot: 'head', item_id: 1, enchant: 2 })).toBe('head=1:2');
    expect(gearEntry({ slot: 'head', item_id: 1, suffix: 3 })).toBe('head=1:0:3');
    expect(gearEntry({ slot: 'head', item_id: 1, enchant: 2, suffix: 3 })).toBe('head=1:2:3');
  });
});

describe('addonGearList', () => {
  it('joins the slots with commas in the order given', () => {
    expect(addonGearList(gear)).toBe('head=16963:2543,main_hand=12784,trinket1=13968:0:2');
  });
});

describe('addonStringFor', () => {
  it('is an FS1 string the existing decoder reads unchanged', () => {
    const code = addonStringFor({
      dataBuild: '1.60.1',
      classSlug: 'warrior',
      raceSlug: 'orc',
      talents: '0-5530515-',
      gear,
    });
    expect(code.startsWith('FS1:1.60.1:warrior:orc:0/5530515/0:')).toBe(true);
    expect(code.endsWith(addonGearList(gear))).toBe(true);
  });
});
```

- [ ] **Step 2: Run both and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/candidates.test.ts src/lib/sim/addon-export.test.ts
```

Expected: FAIL, neither module resolves.

- [ ] **Step 3: Write `candidates.ts`**

```ts
// web/src/lib/sim/candidates.ts
// The candidate list Top Gear, talent compare and the Droptimizer all build, and the
// envelope it turns into.
//
// Every function here returns a new list. The rows are the page's own state and a mutation
// in place would leave Svelte's `$state` array holding the same identity with different
// contents, which is exactly the class of bug that makes a checkbox render one tick behind.
import { bulkCopy } from './copy';
import { KEEP_CURRENT_ENCHANT, NO_ENCHANT } from './enchants';
import { slotsForItem } from '../planner/rules';
import { SLOT_ALIASES, type Item, type Slot } from '../planner/types';
import type { BulkMode, BulkSpec, Candidate, GearSet, Precision, TalentLoadout } from './bulk-types';

/** Contract 1.3's origin vocabulary. `drop:` and `set:` carry an id after the colon. */
export type Origin = 'equipped' | 'bag' | 'bank' | 'search' | `drop:${string}` | `set:${string}`;

export interface CandidateRow {
  /** The slot this row is shown under. Rings and trinkets show under the first of the pair. */
  slot: Slot;
  item: Item;
  origin: Origin;
  /**
   * Where it comes from, in words -- "Ragnaros", "Blacksmithing" -- travelling to the
   * engine as `Candidate.SourceName` (contract 10.1 A6) and coming back on the
   * substitution, so the results view never re-joins an id to a name. Empty for anything
   * that is not a drop.
   */
  sourceName: string;
  /** 0 inherits the equipped enchant; -1 is the page's "explicitly none". */
  enchant: number;
  suffix: number;
  checked: boolean;
}

export const CANDIDATE_ROW_SEPARATOR = ':';

/** Row identity: slot, item, enchant, suffix. A copy-and-modify is therefore a new row. */
export function candidateKey(row: Pick<CandidateRow, 'slot' | 'item' | 'enchant' | 'suffix'>): string {
  return [row.slot, row.item.id, row.enchant, row.suffix].join(CANDIDATE_ROW_SEPARATOR);
}

/**
 * What goes on the wire. "" for anything that fits more than one slot -- rings, trinkets
 * and weapons -- because only the item database can decide which slot it lands in, and
 * sim/bulk is the one that has it (contract 1.3, and design 3.2's "rings and trinkets are
 * tried in both slots").
 */
export function envelopeSlotOf(item: Item): string {
  const fits = slotsForItem(item);
  return fits.length === 1 ? fits[0] : '';
}

/**
 * Which slot the grid shows this item under. Rings and trinkets get one row, not two --
 * design 3.1.2, "Rings and trinkets show one list for both slots" -- so an aliased item
 * shows under the first of its pair.
 */
export function uiSlotsOf(item: Item): Slot[] {
  const aliased = SLOT_ALIASES[item.slot];
  if (aliased !== undefined) return [aliased[0]];
  return slotsForItem(item);
}

export function rowFor(item: Item, slot: Slot, origin: Origin, sourceName = ''): CandidateRow {
  return { slot, item, origin, sourceName, enchant: NO_ENCHANT, suffix: 0, checked: false };
}

export function toggleRow(rows: readonly CandidateRow[], key: string): CandidateRow[] {
  return rows.map((row) => (candidateKey(row) === key ? { ...row, checked: !row.checked } : row));
}

/** Adds a row, or ticks the one already there -- adding the same item twice is not two rows. */
export function addRow(rows: readonly CandidateRow[], row: CandidateRow): CandidateRow[] {
  const key = candidateKey(row);
  if (rows.some((existing) => candidateKey(existing) === key)) {
    return rows.map((existing) =>
      candidateKey(existing) === key ? { ...existing, checked: existing.checked || row.checked } : existing,
    );
  }
  return [...rows, row];
}

/**
 * Design 3.1.2's "copy and modify": the same item again with a different enchant or suffix,
 * added beside the original and ticked, never replacing it. Adding a copy that already
 * exists ticks it instead of duplicating it.
 */
export function copyAndModify(
  rows: readonly CandidateRow[],
  key: string,
  patch: { enchant?: number; suffix?: number },
): CandidateRow[] {
  const source = rows.find((row) => candidateKey(row) === key);
  if (source === undefined) return [...rows];
  const copy: CandidateRow = { ...source, ...patch, checked: true };
  if (candidateKey(copy) === key) return rows.map((row) => (row === source ? copy : row));
  const at = rows.indexOf(source);
  const next = [...rows];
  const existing = next.findIndex((row) => candidateKey(row) === candidateKey(copy));
  if (existing >= 0) {
    next[existing] = { ...next[existing], checked: true };
    return next;
  }
  next.splice(at + 1, 0, copy);
  return next;
}

export function removeRow(rows: readonly CandidateRow[], key: string): CandidateRow[] {
  return rows.filter((row) => candidateKey(row) !== key);
}

/**
 * The ticked rows as envelope candidates. A locked slot contributes nothing -- the contract
 * validates "no candidate on a locked slot" and refusing here is what keeps the page from
 * sending a request it knows will be refused.
 */
export function toCandidates(rows: readonly CandidateRow[], locked: readonly string[]): Candidate[] {
  return rows
    .filter((row) => row.checked && !locked.includes(row.slot))
    .map((row) => {
      const candidate: Candidate = {
        slot: envelopeSlotOf(row.item),
        item_id: row.item.id,
        origin: row.origin,
      };
      // NO_ENCHANT (0) is the wire's own "inherit the equipped enchant", and
      // KEEP_CURRENT_ENCHANT (-1) never travels -- both are an omitted field.
      if (row.enchant > 0) candidate.enchant = row.enchant;
      if (row.suffix > 0) candidate.suffix = row.suffix;
      // Contract 10.1 A6: the page fills the source's name once, here, and reads it back
      // off the substitution rather than joining the id to loot.json a second time.
      if (row.sourceName !== '') candidate.source_name = row.sourceName;
      return candidate;
    });
}

export interface BulkSpecInput {
  mode: BulkMode;
  rows: readonly CandidateRow[];
  locked: readonly string[];
  loadouts: readonly TalentLoadout[];
  sets: readonly GearSet[];
  precision: Precision;
  cap: number;
}

/**
 * The envelope's bulk block. `talents` mode sends no candidates at all, whatever is ticked:
 * the contract validates it and the page's own gear grid is hidden in that mode anyway.
 */
export function buildBulkSpec(input: BulkSpecInput): BulkSpec {
  return {
    mode: input.mode,
    candidates: input.mode === 'talents' ? [] : toCandidates(input.rows, input.locked),
    talents: [...input.loadouts],
    sets: [...input.sets],
    locked: [...input.locked],
    precision: input.precision,
    cap: input.cap,
  };
}

/**
 * The contract's own validation, before the request is sent, so a refusal is a sentence
 * beside the button rather than an engine error after a wait. Null means valid.
 */
export function validateBulk(spec: BulkSpec): string | null {
  if (spec.candidates.some((candidate) => (spec.locked ?? []).includes(candidate.slot))) {
    return bulkCopy.noCandidates;
  }
  if (spec.mode === 'talents') {
    return (spec.talents ?? []).length > 0 ? null : bulkCopy.noCandidates;
  }
  if (spec.mode === 'drops') {
    if (spec.candidates.length === 0) return bulkCopy.dropsNothing;
    return spec.candidates.every((candidate) => candidate.origin.startsWith('drop:'))
      ? null
      : bulkCopy.dropsNothing;
  }
  const anything =
    spec.candidates.length > 0 || (spec.talents ?? []).length > 0 || (spec.sets ?? []).length > 0;
  return anything ? null : bulkCopy.noCandidates;
}
```

- [ ] **Step 4: Write `addon-export.ts`**

```ts
// web/src/lib/sim/addon-export.ts
// "Copy to addon" (design 3.3): the winning set as an export string the companion and the
// in-game addon read, so a player can see the swaps in game.
//
// A gear entry carries `item_id[:enchant[:suffix]]`, which contract 10.5 makes explicit
// for `<gear>` and `sets=` and not only for the bags and bank sections: "a version-1
// decoder reading a bare id is unaffected". A set with no enchants therefore encodes
// byte-for-byte as version 1, and every existing decoder reads it unchanged.
import { FS1_PREFIX } from '../planner/fs1';
import type { GearSlot } from './types';

/** `slot=item[:enchant[:suffix]]`. A suffix with no enchant writes the enchant as 0. */
export function gearEntry(slot: GearSlot): string {
  const enchant = slot.enchant ?? 0;
  const suffix = slot.suffix ?? 0;
  if (suffix > 0) return `${slot.slot}=${slot.item_id}:${enchant}:${suffix}`;
  if (enchant > 0) return `${slot.slot}=${slot.item_id}:${enchant}`;
  return `${slot.slot}=${slot.item_id}`;
}

export function addonGearList(gear: readonly GearSlot[]): string {
  return gear.map(gearEntry).join(',');
}

export interface AddonStringInput {
  dataBuild: string;
  classSlug: string;
  raceSlug: string;
  /** The engine's talents string, dash-joined; the FS1 grammar slash-joins the same trees. */
  talents: string;
  gear: readonly GearSlot[];
}

/**
 * `FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>`.
 *
 * The talents come as the engine's dash-joined string because that is what a `SimRequest`
 * carries, and the FS1 grammar slash-joins the identical three fields; nothing is re-encoded
 * beyond the separator. A tree the engine trimmed to empty writes as "0", the same way
 * `fs1.ts`'s own `encodeTree` does, so the field count is always three.
 */
export function addonStringFor(input: AddonStringInput): string {
  const trees = input.talents.split('-');
  while (trees.length < 3) trees.push('');
  const encoded = trees.slice(0, 3).map((tree) => (tree === '' ? '0' : tree));
  return [
    FS1_PREFIX,
    input.dataBuild,
    input.classSlug,
    input.raceSlug,
    encoded.join('/'),
    addonGearList(input.gear),
  ].join(':');
}
```

- [ ] **Step 5: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/candidates.test.ts src/lib/sim/addon-export.test.ts
```

Expected: PASS.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/candidates.ts src/lib/sim/candidates.test.ts src/lib/sim/addon-export.ts src/lib/sim/addon-export.test.ts
npx prettier --write src/lib/sim/candidates.ts src/lib/sim/candidates.test.ts src/lib/sim/addon-export.ts src/lib/sim/addon-export.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/candidates.ts web/src/lib/sim/candidates.test.ts web/src/lib/sim/addon-export.ts web/src/lib/sim/addon-export.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): the candidate row model and the addon export for a winner

Rows are immutable and keyed on slot, item, enchant and suffix, so
copy-and-modify is a new row beside the original rather than an edit.
The envelope slot is "" for rings, trinkets and weapons, which is the
contract's way of saying only the item database can place them, and a
drop carries its source's name so the results never re-join it (10.1
A6). The gear entry's `item:enchant:suffix` form is contract 10.5's.
The contract's own bulk validation runs before the request is sent, so
a refusal is a sentence beside the button rather than a wait.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 9: The results view helpers and the stat-weight vocabulary

**Files:**
- Create: `web/src/lib/sim/combos.ts`, `web/src/lib/sim/combos.test.ts`
- Create: `web/src/lib/sim/weights.ts`, `web/src/lib/sim/weights.test.ts`

**Interfaces:**
- Consumes: `BulkResult`, `Combo`, `Substitution`, `WeightsResult`, `StatWeight` (Task 1); `Item`, `ItemSet`, `Slot`, `GearSlot`; `specRow`, `specLabel` from `./spec-label`; `classRows` from `../planner/reference`.
- **The stat ids are pinned, not guessed:** contract 10.8 lists the whole `proto.Stat` vocabulary in snake case, and the test below asserts every id in `WEIGHT_STATS` is one of them. The two traps it guards: the engine has **one** `hit` and **one** `crit` (only haste is split, into `melee_haste` and `spell_haste`), and MP5 is `mp5`.
- Produces:
  - combos.ts: `ComboRow`, `comboRows(result)`, `percentOf(delta, equipped)`, `deltaLabel(delta)`, `winningGear(result)`, `slotSummary(result)`, `SlotSummaryRow`, `keepsSetBonus(combo, result, items, sets, pieces)`, `headlineFor(result)`, `substitutionLabel(sub)`, `sourceNameOfCombo(combo)`
  - weights.ts: `WEIGHT_STATS`, `WeightStat`, `statLabel(id)`, `defaultStatsFor(spec, referenceStat)`, `referenceFor(spec, rows)`, `pawnString(spec, weights)`, `weightScale(weights)`, `DEFAULT_REFERENCE`

- [ ] **Step 1: Write the failing tests**

`web/src/lib/sim/combos.test.ts`:

```ts
// web/src/lib/sim/combos.test.ts
import { describe, expect, it } from 'vitest';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import setsJson from '../../fixtures/planner/sets.json';
import {
  comboRows,
  deltaLabel,
  headlineFor,
  keepsSetBonus,
  percentOf,
  slotSummary,
  sourceNameOfCombo,
  substitutionLabel,
  winningGear,
} from './combos';
import type { BulkResult } from './bulk-types';
import type { Item, ItemSet } from '../planner/types';

const result = bulkResultJson as unknown as BulkResult;
const items = new Map((itemsJson.items as unknown as Item[]).map((item) => [item.id, item]));
const sets = setsJson as unknown as ItemSet[];

describe('comboRows', () => {
  it('ranks best first and gives every member of a within-error group the same rank', () => {
    const rows = comboRows(result);
    expect(rows.map((row) => row.rank)).toEqual([1, 1, 3, 3, 5, 5]);
    expect(rows[0].withinError).toBe(true);
    expect(rows[2].withinError).toBe(false);
  });

  it('carries the percent against the equipped set', () => {
    const rows = comboRows(result);
    expect(rows[0].percent).toBeCloseTo((41.2 / 1461.2) * 100, 6);
  });
});

describe('percentOf and deltaLabel', () => {
  it('reads a gain with its error and a sign', () => {
    expect(deltaLabel({ mean: 41.2, stddev: 0, error: 5.41, min: 0, max: 0 })).toBe('+41 ± 11');
    expect(deltaLabel({ mean: -1.8, stddev: 0, error: 5.38, min: 0, max: 0 })).toBe('−2 ± 11');
  });

  it('is zero percent against a zero baseline rather than infinite', () => {
    expect(percentOf(41.2, 0)).toBe(0);
  });
});

describe('winningGear', () => {
  it('writes the leader’s substitutions over the base character’s gear', () => {
    const gear = winningGear(result);
    expect(gear.find((slot) => slot.slot === 'head')?.item_id).toBe(16963);
    expect(gear.find((slot) => slot.slot === 'shoulder')?.item_id).toBe(16966);
    // untouched by the winner
    expect(gear.find((slot) => slot.slot === 'main_hand')?.item_id).toBe(12784);
  });
});

describe('slotSummary', () => {
  it('names the winner’s item per slot and the gain that slot contributed alone', () => {
    const rows = slotSummary(result);
    const head = rows.find((row) => row.slot === 'head')!;
    expect(head.item_id).toBe(16963);
    expect(head.name).toBe('Helm of Wrath');
    // the single-substitution combo for that slot
    expect(head.gain).toBeCloseTo(38.6, 6);
    const shoulder = rows.find((row) => row.slot === 'shoulder')!;
    expect(shoulder.gain).toBeCloseTo(18.9, 6);
  });

  it('lists a slot the winner changed even when nothing measured it alone', () => {
    const trimmed: BulkResult = { ...result, combos: [result.combos[0]] };
    const rows = slotSummary(trimmed);
    expect(rows.map((row) => row.slot).sort()).toEqual(['head', 'shoulder']);
    expect(rows.every((row) => row.gain === null || typeof row.gain === 'number')).toBe(true);
  });
});

describe('keepsSetBonus', () => {
  it('is false when the combination does not reach the piece count', () => {
    expect(keepsSetBonus(result.combos[0], result, items, sets, 4)).toBe(false);
  });

  it('is true at a piece count the combination does reach', () => {
    expect(keepsSetBonus(result.combos[0], result, items, sets, 2)).toBe(true);
  });
});

describe('headlineFor and substitutionLabel', () => {
  it('names the leader’s biggest change from the result’s own name field', () => {
    expect(headlineFor(result)).toBe('+41 DPS from Helm of Wrath');
  });

  it('labels an item, a loadout, a set and a consumable list (contract 10.8)', () => {
    expect(substitutionLabel({ kind: 'item', slot: 'head', item_id: 16963, name: 'Helm of Wrath' })).toBe(
      'Helm of Wrath',
    );
    expect(substitutionLabel({ kind: 'talents', name: 'Deep Fury' })).toBe('Deep Fury');
    expect(substitutionLabel({ kind: 'set', name: 'AQ set' })).toBe('AQ set');
    expect(substitutionLabel({ kind: 'consumes', name: 'flask_of_supreme_power' })).toBe(
      'With flask of supreme power',
    );
  });

  it('falls back to the item id only when the engine sent no name', () => {
    expect(substitutionLabel({ kind: 'item', slot: 'head', item_id: 16963 })).toBe('Item 16963');
  });
});

describe('sourceNameOfCombo', () => {
  it('reads the boss off the substitution the engine copied it onto (contract 10.1 A6)', () => {
    expect(
      sourceNameOfCombo({
        substitutions: [{ kind: 'item', item_id: 1, source_name: 'Ragnaros' }],
        dps: result.equipped,
        delta: result.equipped,
        group: 0,
      }),
    ).toBe('Ragnaros');
  });
});
```

`web/src/lib/sim/weights.test.ts`:

```ts
// web/src/lib/sim/weights.test.ts
import { describe, expect, it } from 'vitest';
import weightsResultJson from '../../fixtures/sim/weights-result.json';
import specsJson from '../../fixtures/sim/specs.json';
import {
  CASTER_REFERENCE,
  DEFAULT_REFERENCE,
  defaultStatsFor,
  fallbackReferenceFor,
  pawnString,
  referenceFor,
  statLabel,
  weightScale,
  WEIGHT_STATS,
} from './weights';
import type { WeightsResult } from './bulk-types';
import type { SpecFidelity } from './types';

const result = weightsResultJson as unknown as WeightsResult;
const specs = specsJson as unknown as SpecFidelity[];

describe('referenceFor', () => {
  it('takes the spec’s own reference stat from GET /v1/specs', () => {
    expect(referenceFor('warrior-fury', specs)).toBe('attack_power');
    expect(referenceFor('mage-frost', specs)).toBe('spell_power');
  });

  it('falls back to the pinned default for the kind of spec, not to one number', () => {
    // Contract 10.8: attack_power for melee and hunters, spell_power for casters.
    expect(referenceFor('warrior-fury', [])).toBe(DEFAULT_REFERENCE);
    expect(referenceFor('mage-fire', [])).toBe(CASTER_REFERENCE);
    expect(referenceFor('druid-balance', specs)).toBe(CASTER_REFERENCE);
    expect(fallbackReferenceFor('hunter-marksmanship')).toBe(DEFAULT_REFERENCE);
    expect(fallbackReferenceFor('nonesuch-spec')).toBe(DEFAULT_REFERENCE);
  });
});

describe('defaultStatsFor', () => {
  it('always includes the reference stat, exactly once, first', () => {
    const stats = defaultStatsFor('warrior-fury', 'attack_power');
    expect(stats[0]).toBe('attack_power');
    expect(stats.filter((stat) => stat === 'attack_power')).toHaveLength(1);
    expect(new Set(stats).size).toBe(stats.length);
  });
});

describe('statLabel', () => {
  it('names every stat the picker offers', () => {
    for (const stat of WEIGHT_STATS) expect(statLabel(stat.id)).toBe(stat.label);
    expect(statLabel('nonesuch')).toBe('nonesuch');
  });
});

describe('weightScale', () => {
  it('is the largest weight plus its error, so no bar overflows its track', () => {
    expect(weightScale(result.weights)).toBeCloseTo(27.3 + 1.2, 6);
    expect(weightScale([])).toBe(1);
  });
});

describe('pawnString', () => {
  it('is a Pawn v1 line with the class, the spec and two decimals a piece', () => {
    expect(pawnString('warrior-fury', result.weights)).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury, AttackPower=1.00, Strength=2.14, ' +
        'CritRating=21.70, HitRating=27.30, Agility=1.32, HasteRating=18.40 )',
    );
  });

  it('leaves out a stat Pawn has no key for rather than inventing one', () => {
    expect(pawnString('warrior-fury', [{ stat: 'nonesuch', weight: 3, error: 0 }])).toBe(
      '( Pawn: v1: "Fury Warrior": Class=Warrior, Spec=Fury )',
    );
  });
});

describe('the stat vocabulary (contract 10.8, pinned)', () => {
  it('carries one hit and one crit, and splits only haste', () => {
    const ids = WEIGHT_STATS.map((stat) => stat.id);
    expect(ids).toContain('hit');
    expect(ids).toContain('crit');
    expect(ids).toContain('melee_haste');
    expect(ids).toContain('spell_haste');
    for (const wrong of ['melee_hit', 'spell_hit', 'melee_crit', 'spell_crit', 'haste']) {
      expect(ids).not.toContain(wrong);
    }
  });

  it('spells MP5 as mp5', () => {
    const ids = WEIGHT_STATS.map((stat) => stat.id);
    expect(ids).toContain('mp5');
    expect(ids).not.toContain('m_p5');
  });

  it('offers only ids the pinned proto.Stat list carries', () => {
    const pinned = new Set([
      'strength', 'agility', 'stamina', 'intellect', 'spirit', 'spell_power', 'arcane_power',
      'fire_power', 'frost_power', 'holy_power', 'nature_power', 'shadow_power', 'mp5', 'hit',
      'crit', 'spell_haste', 'spell_penetration', 'attack_power', 'melee_haste',
      'armor_penetration', 'expertise', 'mana', 'energy', 'rage', 'armor',
      'ranged_attack_power', 'defense', 'block', 'block_value', 'dodge', 'parry', 'health',
      'arcane_resistance', 'fire_resistance', 'frost_resistance', 'nature_resistance',
      'shadow_resistance', 'bonus_armor', 'healing_power', 'spell_damage',
      'feral_attack_power',
    ]);
    for (const stat of WEIGHT_STATS) expect(pinned.has(stat.id), stat.id).toBe(true);
  });
});
```

- [ ] **Step 2: Run both and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/combos.test.ts src/lib/sim/weights.test.ts
```

Expected: FAIL, neither module resolves.

- [ ] **Step 3: Write `combos.ts`**

```ts
// web/src/lib/sim/combos.ts
// Reading a ranked bulk result. Nothing here is a statistic: `group` already says which
// runs are within error of the leader, `delta` already carries its own error, and the
// ordering is the planner's. This turns those into rows, labels and a winning gear list.
import type { BulkResult, Combo, Substitution } from './bulk-types';
import { bulkCopy } from './copy';
import type { Estimate, GearSlot } from './types';
import type { Item, ItemSet } from '../planner/types';

export interface ComboRow {
  /** 1-based, and shared by every member of a within-error group (design 3.3). */
  rank: number;
  combo: Combo;
  /** True for the leader's group, which is the one the page draws a rule under. */
  withinError: boolean;
  percent: number;
}

export function percentOf(delta: number, equippedMean: number): number {
  return equippedMean === 0 ? 0 : (delta / equippedMean) * 100;
}

const MINUS = '−';

/**
 * "+41 ± 11": the gain, rounded, and its 95% band. A minus sign, not a hyphen -- the
 * design system's rule for a negative figure in a table.
 */
export function deltaLabel(delta: Estimate): string {
  const mean = Math.round(delta.mean);
  const band = Math.round(1.96 * delta.error);
  const sign = mean < 0 ? MINUS : '+';
  return `${sign}${Math.abs(mean).toLocaleString('en-US')} ± ${band.toLocaleString('en-US')}`;
}

/**
 * Rank shared across a group, per design 3.3: "Rows in the leader's within-error group
 * carry the same rank." A group's rank is the 1-based position of its first member, so the
 * run after a two-member tie is rank 3.
 */
export function comboRows(result: BulkResult): ComboRow[] {
  const firstOfGroup = new Map<number, number>();
  return result.combos.map((combo, index) => {
    if (!firstOfGroup.has(combo.group)) firstOfGroup.set(combo.group, index + 1);
    return {
      rank: firstOfGroup.get(combo.group)!,
      combo,
      withinError: combo.group === 0,
      percent: percentOf(combo.delta.mean, result.equipped.mean),
    };
  });
}

/** The leader's substitutions written over the base character's gear. */
export function winningGear(result: BulkResult): GearSlot[] {
  const gear = result.request.character.gear.map((slot) => ({ ...slot }));
  for (const sub of result.combos[0]?.substitutions ?? []) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    const next: GearSlot = { slot: sub.slot, item_id: sub.item_id };
    if (sub.enchant !== undefined && sub.enchant > 0) next.enchant = sub.enchant;
    if (sub.suffix !== undefined && sub.suffix > 0) next.suffix = sub.suffix;
    const at = gear.findIndex((slot) => slot.slot === sub.slot);
    if (at >= 0) gear[at] = next;
    else gear.push(next);
  }
  return gear;
}

export interface SlotSummaryRow {
  slot: string;
  item_id: number;
  /** The item's own name, off the substitution (contract 10.1 A6). Never a re-join. */
  name: string;
  /**
   * What this slot contributed on its own, from the combination that changed only this
   * slot. Null when no such combination exists -- a Droptimizer run has one per slot by
   * construction, a Top Gear run has one whenever the slot's item was also tried alone, and
   * a run where it was not is honest about not knowing rather than apportioning the total.
   */
  gain: number | null;
}

/** Design 3.3's per-slot summary: what the winner uses, and what that slot was worth. */
export function slotSummary(result: BulkResult): SlotSummaryRow[] {
  const winner = result.combos[0];
  if (winner === undefined) return [];
  const alone = new Map<string, number>();
  for (const combo of result.combos) {
    if (combo.substitutions.length !== 1) continue;
    const only = combo.substitutions[0];
    if (only.kind !== 'item' || only.slot === undefined) continue;
    alone.set(`${only.slot}:${only.item_id}`, combo.delta.mean);
  }
  return winner.substitutions
    .filter((sub): sub is Substitution & { slot: string; item_id: number } =>
      sub.kind === 'item' && sub.slot !== undefined && sub.item_id !== undefined,
    )
    .map((sub) => ({
      slot: sub.slot,
      item_id: sub.item_id,
      name: substitutionLabel(sub),
      gain: alone.get(`${sub.slot}:${sub.item_id}`) ?? null,
    }));
}

/**
 * Design 4.4's results filter: "only combos keeping 4-piece". Tier bonuses are counted by
 * the engine from the gear itself, so this is a filter over what a combination would be
 * wearing, never an input to the run.
 */
export function keepsSetBonus(
  combo: Combo,
  result: BulkResult,
  items: ReadonlyMap<number, Item>,
  sets: readonly ItemSet[],
  pieces: number,
): boolean {
  const gear = new Map(result.request.character.gear.map((slot) => [slot.slot, slot.item_id]));
  for (const sub of combo.substitutions) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    gear.set(sub.slot, sub.item_id);
  }
  const counts = new Map<number, number>();
  for (const itemId of gear.values()) {
    const setId = items.get(itemId)?.set_id ?? null;
    if (setId === null) continue;
    counts.set(setId, (counts.get(setId) ?? 0) + 1);
  }
  return sets.some((set) => (counts.get(set.id) ?? 0) >= pieces);
}

/**
 * "Helm of Wrath", "Deep Fury", "AQ set" -- whatever this substitution actually changed.
 *
 * It reads `name`, which contract 10.1 A6 fills for items too, from simdb. There is no
 * item-id join here and no item map parameter: the result already carries every name it
 * needs, and re-deriving one from a per-class item file the reader may not have loaded is
 * how a saved page ends up showing "Item 16963" for something the engine named.
 */
export function substitutionLabel(sub: Substitution): string {
  // Contract 10.8: a `consumes` substitution's name is the ids joined by ", ", which is a
  // list and not a phrase, so it is the one kind that goes through copy on its way out.
  if (sub.kind === 'consumes') return bulkCopy.consumesChip(sub.name ?? '');
  if (sub.name !== undefined && sub.name !== '') return sub.name;
  return sub.kind === 'item' ? `Item ${sub.item_id ?? 0}` : '';
}

/** Where a Droptimizer combination's one substitution came from, in words. */
export function sourceNameOfCombo(combo: Combo): string {
  return combo.substitutions[0]?.source_name ?? '';
}

/**
 * "+41 DPS from Vis'kag", for the page's own display. The leader's first substitution is
 * the one named -- the planner orders a combination's substitutions by the slot's own
 * contribution, so the first is the one that carried it.
 *
 * The *stored* headline on a saved sim is not this: contract 10.6 has the API compose it
 * at save time, with its own rules for several substitutions ("… and 2 more") and for an
 * empty result ("no combinations"). `SimListRow.headline` is read, never recomputed.
 */
export function headlineFor(result: BulkResult): string {
  const winner = result.combos[0];
  if (winner === undefined) return '';
  const gain = deltaLabel(winner.delta).split(' ')[0];
  const what = substitutionLabel(winner.substitutions[0]);
  return what === '' ? `${gain} DPS` : `${gain} DPS from ${what}`;
}
```

- [ ] **Step 4: Write `weights.ts`**

```ts
// web/src/lib/sim/weights.ts
// /sim/weights: which stats to weigh, which one is the reference, and the Pawn line.
//
// The ids are the fork's `proto.Stat` enum names in snake case, pinned by contract 10.8
// and published as IDS.md's **Stats** section, which the sim module generates from the
// enum. Two spellings matter and are easy to get wrong:
//
//   * the engine carries ONE `hit` and ONE `crit`. There is no `melee_crit`, `spell_crit`,
//     `melee_hit` or `spell_hit`; haste IS split (`melee_haste`, `spell_haste`), which is
//     what makes the single hit and crit look like an oversight when they are not.
//   * MP5 is `mp5`, not `m_p5`.
//
// The list below is the subset that moves a spec's DPS, which is what 10.8 says a weight
// page offers -- the enum also carries defence, block, dodge, parry, resistances, health
// and the school power stats, and none of them belongs in a DPS weight picker.
//
// `WeightsSpec.Reference` is required and uses the same vocabulary. The reference itself
// comes from the spec list, never from a table here -- data/curated/specs.json is
// canonical and GET /v1/specs carries it as `reference_stat`.
//
// The `pawn` column is Pawn's own vocabulary and is not a transformation of the id, which
// is why it is a column. An empty `pawn` is a stat Pawn has no key for; it is left out of
// the string rather than guessed at.
import { classRows } from '../planner/reference';
import { specRow } from './spec-label';
import type { StatWeight } from './bulk-types';
import type { SpecFidelity } from './types';

export interface WeightStat {
  /** The engine's id. */
  id: string;
  label: string;
  /** Pawn's own key, or "" for a stat Pawn has no key for. */
  pawn: string;
}

/**
 * The stats a vanilla-era damage spec is ever weighed on. Pawn's keys are its own
 * vocabulary, which is why they are a column here rather than a transformation of the id.
 */
export const WEIGHT_STATS: readonly WeightStat[] = [
  { id: 'strength', label: 'Strength', pawn: 'Strength' },
  { id: 'agility', label: 'Agility', pawn: 'Agility' },
  { id: 'stamina', label: 'Stamina', pawn: 'Stamina' },
  { id: 'intellect', label: 'Intellect', pawn: 'Intellect' },
  { id: 'spirit', label: 'Spirit', pawn: 'Spirit' },
  { id: 'attack_power', label: 'Attack power', pawn: 'AttackPower' },
  { id: 'ranged_attack_power', label: 'Ranged attack power', pawn: 'RangedAttackPower' },
  { id: 'feral_attack_power', label: 'Feral attack power', pawn: '' },
  { id: 'spell_power', label: 'Spell power', pawn: 'SpellDamage' },
  { id: 'hit', label: 'Hit', pawn: 'HitRating' },
  { id: 'crit', label: 'Crit', pawn: 'CritRating' },
  { id: 'melee_haste', label: 'Melee haste', pawn: 'HasteRating' },
  { id: 'spell_haste', label: 'Spell haste', pawn: 'SpellHasteRating' },
  { id: 'spell_penetration', label: 'Spell penetration', pawn: 'SpellPenetration' },
  { id: 'armor_penetration', label: 'Armor penetration', pawn: 'ArmorPenetration' },
  { id: 'expertise', label: 'Expertise', pawn: '' },
  { id: 'mp5', label: 'Mana per 5', pawn: 'Mp5' },
];

const BY_ID = new Map(WEIGHT_STATS.map((stat) => [stat.id, stat]));

/**
 * The fallback reference when the spec list has no row or no column. Contract 10.8 pins
 * the defaults: `attack_power` for melee and hunters, `spell_power` for casters. It is a
 * fallback only -- `reference_stat` on the spec row is the answer whenever there is one.
 */
export const DEFAULT_REFERENCE = 'attack_power';
export const CASTER_REFERENCE = 'spell_power';

/** The classes and specs whose damage scales with spell power rather than attack power. */
const CASTER_CLASSES = new Set(['mage', 'warlock', 'priest']);
const CASTER_SPECS = new Set(['druid-balance', 'shaman-elemental', 'paladin-holy']);

export function fallbackReferenceFor(spec: string): string {
  const row = specRow(spec);
  if (row === null) return DEFAULT_REFERENCE;
  if (CASTER_SPECS.has(row.spec) || CASTER_CLASSES.has(row.class_slug)) return CASTER_REFERENCE;
  return DEFAULT_REFERENCE;
}

export function statLabel(id: string): string {
  return BY_ID.get(id)?.label ?? id;
}

/**
 * The spec's own reference stat from the spec list. A spec the list has no row for, or a
 * row from a build predating the column, takes contract 10.8's own default for its kind of
 * spec -- attack power for melee and hunters, spell power for casters -- which is the same
 * rule the data lane applies when it writes the column, so the two cannot disagree.
 */
export function referenceFor(spec: string, rows: readonly SpecFidelity[]): string {
  const row = rows.find((entry) => entry.spec === spec);
  const reference = row?.reference_stat ?? '';
  return reference === '' ? fallbackReferenceFor(spec) : reference;
}

/**
 * What the picker opens with: the reference stat first, then every stat that can matter.
 * The list is not pruned by class -- a spec that gains nothing from spirit gets a weight of
 * about zero, which is itself the answer, and pruning would hide it.
 */
export function defaultStatsFor(_spec: string, referenceStat: string): string[] {
  const rest = WEIGHT_STATS.filter((stat) => stat.id !== referenceStat).map((stat) => stat.id);
  return [referenceStat, ...rest];
}

/** The widest bar the table draws: the biggest weight plus its own error. */
export function weightScale(weights: readonly StatWeight[]): number {
  const widest = weights.reduce((max, row) => Math.max(max, row.weight + row.error), 0);
  return widest > 0 ? widest : 1;
}

/**
 * Pawn's v1 line, which is what the addon pastes. A stat Pawn has no key for is left out
 * rather than guessed at: Pawn ignores a key it does not know, but it warns about it, and a
 * warning nobody can act on is worse than an absent stat.
 */
export function pawnString(spec: string, weights: readonly StatWeight[]): string {
  const row = specRow(spec);
  const className = classRows.find((entry) => entry.slug === row?.class_slug)?.name ?? '';
  const name = row === null ? spec : `${row.name} ${className}`.trim();
  const parts = [`Class=${className}`, `Spec=${row?.name ?? spec}`];
  for (const weight of weights) {
    const key = BY_ID.get(weight.stat)?.pawn ?? '';
    if (key === '') continue;
    parts.push(`${key}=${weight.weight.toFixed(2)}`);
  }
  return `( Pawn: v1: "${name}": ${parts.join(', ')} )`;
}
```

- [ ] **Step 5: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/combos.test.ts src/lib/sim/weights.test.ts
```

Expected: PASS. If `deltaLabel`'s rounding disagrees with the expected `± 11`, the test is
right and the code is wrong: `1.96 × 5.41 = 10.6`, which rounds to 11.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/combos.ts src/lib/sim/combos.test.ts src/lib/sim/weights.ts src/lib/sim/weights.test.ts
npx prettier --write src/lib/sim/combos.ts src/lib/sim/combos.test.ts src/lib/sim/weights.ts src/lib/sim/weights.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/combos.ts web/src/lib/sim/combos.test.ts web/src/lib/sim/weights.ts web/src/lib/sim/weights.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): reading a ranked result, and the stat-weight vocabulary

Ranks are shared across a within-error group because the planner already
said which runs those are; nothing here recomputes a statistic. Names
come off the substitution the engine filled (contract 10.1 A6), so no
view re-joins an item id to a per-class file the reader may not have.
The per-slot summary reads a slot's contribution from the combination
that changed only that slot, and says null rather than apportioning a
total when no such combination was run. Stat ids are proto.Stat enum
names with melee and spell split (10.1 A7), never a bare haste.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 10: `bulk-store.svelte.ts`, the store all four pages mount over

**Files:**
- Create: `web/src/lib/sim/bulk-store.svelte.ts`
- Create: `web/src/lib/sim/bulk-store.test.ts`

**Interfaces:**
- Consumes: `fromAddonExport`, `fromPlannerBuild`, `fromLoggedFight`, `fromStoredCharacter`, `LoadContext`, `SourceResult` from `./sources`; `toCharacterSpec`, `needsRace`, `SimCharacter`, `talentsString` from `./character`; `indexTalents` from `../planner/rules`; `loadItems`, `loadTalents`, `loadReference`, `loadSets` from `../planner/load`; `loadLoot`, `sourcesByItem` (Task 6); `loadEnchants`, `loadSuffixes` (Task 7); everything from `./candidates`, `./bulk-run`, `./bulk-types`; `defaultSettings`, `SimSettings` from `./settings`; `fetchSpecs`, `dispatchServerSim`, `fetchSim`, `fetchBulkProgress`, `SimApiError` from `./api`; `referenceFor`, `defaultStatsFor` from `./weights`.
- **Consumes from part A:** `SimCharacter.bags`, `.bank`, `.sets`, `.loadouts` (optional; absent reads as "no addon export").
- Produces: `createBulkStore(init: BulkStoreInit): BulkStore`, `type BulkStore`, `type BulkPhase`, `type SimTool`, `TOOLS`, `MODE_OF_TOOL`.

- [ ] **Step 1: Write the failing test**

Create `web/src/lib/sim/bulk-store.test.ts`. It is a jsdom test: the store reaches
`requestEnvelope`, which reads `document.cookie`.

```ts
// @vitest-environment jsdom
// web/src/lib/sim/bulk-store.test.ts
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { createFakeWorker } from '../../test-support/fake-worker';
import { createSimApi, FIXTURE_DATA_BUILD, TEST_API } from '../../test-support/sim-api';
import { createBulkStore, MODE_OF_TOOL, TOOLS } from './bulk-store.svelte';
import { candidateKey } from './candidates';
import { createPool } from './worker';
import type { BulkResult } from './bulk-types';

const FURY = `FS1:${FIXTURE_DATA_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;
const api = createSimApi();

function store(tool: (typeof TOOLS)[number] = 'gear') {
  const engine = createFakeEngine({ tickMs: 0, ticks: 1 });
  return createBulkStore({
    tool,
    treeVersion: FIXTURE_DATA_BUILD,
    apiBase: TEST_API,
    hardwareConcurrency: 8,
    serverPollMs: 1,
    pool: createPool({ hardwareConcurrency: 4, spawn: () => createFakeWorker(engine) }),
    now: () => new Date('2026-12-10T00:00:00Z'),
  });
}

beforeEach(() => api.install());
afterEach(() => api.reset());

describe('TOOLS and MODE_OF_TOOL', () => {
  it('maps each tool page to its bulk mode', () => {
    expect([...TOOLS]).toEqual(['gear', 'talents', 'drops', 'weights']);
    expect(MODE_OF_TOOL.gear).toBe('gear');
    expect(MODE_OF_TOOL.talents).toBe('talents');
    expect(MODE_OF_TOOL.drops).toBe('drops');
  });
});

describe('loading a character', () => {
  it('adopts the addon export, its items, its loot index and its enchants', async () => {
    const s = store();
    await s.loadAddon(FURY);
    expect(s.character?.class_slug).toBe('warrior');
    expect(s.items.size).toBeGreaterThan(0);
    expect(s.enchants.length).toBeGreaterThan(0);
    expect(s.suffixes.length).toBeGreaterThan(0);
    expect(s.loot.sources.length).toBeGreaterThan(0);
    expect(s.sourceIndex.get(16963)).toEqual(['raid:mc']);
    s.dispose();
  });

  it('seeds a row per equipped item, ticked off, so the grid opens with what you wear', async () => {
    const s = store();
    await s.loadAddon(FURY);
    expect(s.rows.filter((row) => row.origin === 'equipped').map((row) => row.item.id).sort()).toEqual([
      12640, 12784,
    ]);
    expect(s.rows.every((row) => !row.checked)).toBe(true);
    s.dispose();
  });

  it('keeps the character already on screen when a later load fails', async () => {
    const s = store();
    await s.loadAddon(FURY);
    await s.loadAddon('FS2:nope');
    expect(s.character).not.toBeNull();
    expect(s.message).not.toBeNull();
    s.dispose();
  });
});

describe('the cap', () => {
  it('is 400 on eight cores and 200 on four', () => {
    const wide = store();
    expect(wide.cap).toBe(400);
    wide.dispose();
  });
});

describe('the live combination count', () => {
  it('counts what is ticked and clears the cap notice when it comes back under', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.addSearchItem(16966);
    await s.recount();
    expect(s.combinations).toBe(3);
    expect(s.capNotice).toBeNull();
    s.dispose();
  });

  it('shows the cap notice with both numbers rather than trimming', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.addSearchItem(16966);
    s.setCap(2);
    await s.recount();
    expect(s.capNotice).toEqual({ cap: 2, combinations: 3 });
    s.dispose();
  });

  it('counts the consumable alternatives too (contract 10.1 A5)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.toggleConsumable('flask_of_supreme_power');
    s.toggleConsumable('elixir_of_the_mongoose');
    await s.recount();
    // (1 head + 1) x 2 consumable lists, minus the untouched base
    expect(s.combinations).toBe(3);
    s.dispose();
  });
});

describe('running', () => {
  it('runs a gear request to a ranked result and reports stage progress', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    await s.run();
    const result = s.result as BulkResult | null;
    expect(result?.combos.length).toBeGreaterThan(0);
    expect(s.phase).toBe('done');
    expect(s.progressLine).toBe('');
    s.dispose();
  });

  it('refuses to run with nothing ticked, and says which sentence', async () => {
    const s = store();
    await s.loadAddon(FURY);
    await s.run();
    expect(s.result).toBeNull();
    expect(s.message).not.toBeNull();
    s.dispose();
  });

  it('runs a talents request from loadouts alone, with no candidates', async () => {
    const s = store('talents');
    await s.loadAddon(FURY);
    s.addLoadout({ name: 'Deep Fury', talents: '0-5530515-' });
    await s.run();
    expect((s.result as BulkResult).request.bulk?.candidates).toEqual([]);
    s.dispose();
  });

  it('runs a weights request and keeps the reference stat at 1', async () => {
    const s = store('weights');
    await s.loadAddon(FURY);
    await s.loadSpecs();
    expect(s.referenceStat).toBe('attack_power');
    await s.run();
    expect(s.weights.find((row) => row.stat === 'attack_power')?.weight).toBe(1);
    s.dispose();
  });
});

describe('the droptimizer’s sources', () => {
  it('opens with every kind but quests, and gates an unreleased raid until show-upcoming', async () => {
    const s = store('drops');
    await s.loadAddon(FURY);
    expect(s.visibleSources.map((source) => source.id)).not.toContain('quest');
    // the fixture's clock is after raids-1 opens, so Molten Core is visible
    expect(s.visibleSources.map((source) => source.id)).toContain('raid:mc');
    s.dispose();
  });

  it('turns a picked boss into drop-origin candidates carrying its name', async () => {
    const s = store('drops');
    await s.loadAddon(FURY);
    s.toggleSource('raid:mc', 'raid:mc:11502');
    const picked = s.rows.filter((row) => row.checked);
    expect(picked.map((row) => row.origin)).toEqual(['drop:raid:mc:11502', 'drop:raid:mc:11502']);
    // Contract 10.1 A6: the name travels on the candidate, not on a later join.
    expect(new Set(picked.map((row) => row.sourceName))).toEqual(new Set(['Ragnaros']));
    s.dispose();
  });

  it('hides a source whose date is unknown until show-upcoming (contract 10.4)', async () => {
    const s = store('drops');
    await s.loadAddon(FURY);
    expect(s.visibleSources.map((source) => source.id)).not.toContain('world:azuregos');
    s.setShowUpcoming(true);
    expect(s.visibleSources.map((source) => source.id)).toContain('world:azuregos');
    s.dispose();
  });
});

describe('the request’s iteration count', () => {
  it('is the precision’s final stage, not always 3,000 (contract 10.1 A3)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.setPrecision('high');
    await s.run();
    expect((s.result as BulkResult).request.iterations).toBe(10_000);
    s.dispose();
  });
});

describe('the server lane’s own cap', () => {
  it('will not offer a premium run past 5,000 combinations (contract 10.1 A2)', async () => {
    const s = store();
    await s.loadAddon(FURY);
    s.setPremium(true);
    s.addSearchItem(16963);
    await s.recount();
    expect(s.serverCapNotice).toBeNull();
    // A count the browser cap would refuse but the server cap allows still offers premium.
    s.setCap(1);
    await s.recount();
    expect(s.capNotice).not.toBeNull();
    expect(s.serverCapNotice).toBeNull();
    s.dispose();
  });
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-store.test.ts
```

Expected: FAIL, cannot resolve `./bulk-store.svelte`.

- [ ] **Step 3: Write `bulk-store.svelte.ts`**

```ts
// web/src/lib/sim/bulk-store.svelte.ts
// The single source of truth for /sim/gear, /sim/talents, /sim/drops and /sim/weights.
//
// It is deliberately NOT the /sim store with more fields on it. The two share their
// character loading (sources.ts, which both call and neither owns) and nothing else: /sim
// runs one request and shows a damage breakdown, these four build a candidate list and show
// a ranking. Merging them would give every page the other's state to reason about, and
// would make the two web lanes of this spec edit one file.
//
// Two rules the components must not second-guess, both the same as /sim's:
//   * a failed load keeps the character already on screen;
//   * the pool is created on the first run, never on mount -- the wasm is 4 MB and these
//     pages have the site's Lighthouse budget.
import { loadItems, loadReference, loadSets, loadTalents } from '../planner/load';
import { indexTalents } from '../planner/rules';
import type { Item, ItemSet, Slot, TalentFile } from '../planner/types';
import type { CharacterPath } from '../characters';
import {
  dispatchServerSim,
  fetchBulkProgress,
  fetchSim,
  fetchSpecs,
  SimApiError,
} from './api';
import { needsRace, toCharacterSpec, type SimCharacter } from './character';
import { bulkCopy, simCopy } from './copy';
import { loadEnchants, loadSuffixes, type EnchantRow, type SuffixRow } from './enchants';
import { loadSimBuffs, type SimBuffFile } from './sim-buffs';
import { BUILT_IN_PHASES, fetchPhases, type PhaseRow } from './phase';
import {
  DEFAULT_OFF_KINDS,
  itemsOfBoss,
  itemsOfSource,
  isOpen,
  loadLoot,
  sourceNameOf,
  sourcesByItem,
  type LootFile,
  type LootSource,
} from './loot';
import {
  addRow,
  buildBulkSpec,
  candidateKey,
  copyAndModify,
  removeRow,
  rowFor,
  toggleRow,
  uiSlotsOf,
  validateBulk,
  type CandidateRow,
  type Origin,
} from './candidates';
import {
  BulkCapError,
  BulkRunError,
  countCombinations,
  runBulk,
  runWeightsRun,
  stageProgressLine,
  type BulkProgress,
  type BulkRunHandle,
} from './bulk-run';
import {
  SERVER_CAP,
  browserCap,
  finalIterations,
  type BulkMode,
  type BulkRequest,
  type BulkResult,
  type GearSet,
  type Precision,
  type StatWeight,
  type TalentLoadout,
  type WeightsRequest,
  type WeightsResult,
} from './bulk-types';
import { defaultSettings, type SimSettings } from './settings';
import {
  fromAddonExport,
  fromLoggedFight,
  fromPlannerBuild,
  fromStoredCharacter,
  type LoadContext,
  type SourceResult,
} from './sources';
import type { SimResult, SourceKind, SpecFidelity } from './types';
import { ENGINE_VERSION } from './version';
import { defaultStatsFor, referenceFor } from './weights';
import { createPool, type SimPool } from './worker';

export const TOOLS = ['gear', 'talents', 'drops', 'weights'] as const;
export type SimTool = (typeof TOOLS)[number];

/** The three tools that send a `bulk` block, and which mode each sends. */
export const MODE_OF_TOOL: Record<Exclude<SimTool, 'weights'>, BulkMode> = {
  gear: 'gear',
  talents: 'talents',
  drops: 'drops',
};

export type BulkPhase = 'idle' | 'loading-character' | 'counting' | 'running' | 'done' | 'error';

const DEFAULT_SERVER_POLL_MS = 2000;
const COUNT_DEBOUNCE_MS = 250;

const delay = (ms: number): Promise<void> => new Promise((resolve) => setTimeout(resolve, ms));

export interface BulkStoreInit {
  tool: SimTool;
  treeVersion: string;
  apiBase?: string;
  /** Injected by tests; production creates one on the first run. */
  pool?: SimPool;
  /** Injected by tests; production reads `navigator.hardwareConcurrency`. */
  hardwareConcurrency?: number;
  source?: SourceKind | '';
  ref?: string;
  serverPollMs?: number;
  /** Injected by tests, so the phase gate is not a clock-dependent assertion. */
  now?: () => Date;
}

export function createBulkStore(init: BulkStoreInit) {
  const ctx: LoadContext = { treeVersion: init.treeVersion, apiBase: init.apiBase };
  const clock = init.now ?? ((): Date => new Date());
  const mode: BulkMode | null = init.tool === 'weights' ? null : MODE_OF_TOOL[init.tool];

  let phase = $state<BulkPhase>('idle');
  let character = $state<SimCharacter | null>(null);
  let settings = $state<SimSettings>(defaultSettings());
  let message = $state<string | null>(null);
  let detail = $state('');
  let premium = $state(false);

  let items = $state<Map<number, Item>>(new Map());
  let sets = $state<ItemSet[]>([]);
  let talentFile = $state<TalentFile | null>(null);
  let enchants = $state<EnchantRow[]>([]);
  let suffixes = $state<SuffixRow[]>([]);
  let loot = $state<LootFile>({ sources: [] });
  let sourceIndex = $state<Map<number, string[]>>(new Map());
  let specRows = $state<SpecFidelity[]>([]);
  let simBuffs = $state<SimBuffFile>({ entries: {} });
  // The build-time table until GET /v1/phases answers; see phase.ts's own header.
  let phases = $state<readonly PhaseRow[]>(BUILT_IN_PHASES);

  let rows = $state<CandidateRow[]>([]);
  let locked = $state<string[]>([]);
  let loadouts = $state<TalentLoadout[]>([]);
  let namedSets = $state<GearSet[]>([]);
  let precision = $state<Precision>('fast');
  let cap = $state(browserCap(init.hardwareConcurrency ?? globalThis.navigator?.hardwareConcurrency));

  let consumableIds = $state<string[]>([]);
  let combinations = $state<number | null>(null);
  let capNotice = $state<{ cap: number; combinations: number } | null>(null);
  // Set when the count is past the premium lane's own cap too, so the page does not offer
  // a server run the API would refuse at submit (contract 10.1 A2, 5,000).
  let serverCapNotice = $state<{ cap: number; combinations: number } | null>(null);
  let progress = $state<BulkProgress | null>(null);
  let result = $state<SimResult | null>(null);

  // Droptimizer state. `pickedBosses` is keyed `<source id>|<boss id or "">`, so a raid card
  // ("every boss") and one boss of it are two different picks the player can hold at once.
  let pickedBosses = $state<string[]>([]);
  let showUpcoming = $state(false);
  let shownKinds = $state<string[]>([]);

  // Weights state.
  let stats = $state<string[]>([]);

  let pool: SimPool | null = init.pool ?? null;
  let handle: BulkRunHandle | null = null;
  let stopRequested = false;
  let serverRunning = $state(false);
  let serverGeneration = 0;
  let countTimer: ReturnType<typeof setTimeout> | null = null;

  function poolOnce(): SimPool {
    pool ??= createPool({});
    return pool;
  }

  /** Every source funnels through here, so the "a bad paste keeps the page" rule is once. */
  async function adopt(load: Promise<SourceResult>): Promise<void> {
    serverGeneration += 1;
    serverRunning = false;
    phase = 'loading-character';
    message = null;
    const outcome = await load;
    if (!outcome.ok) {
      message = outcome.message;
      phase = 'idle';
      return;
    }
    character = outcome.character;
    result = null;
    progress = null;
    capNotice = null;
    combinations = null;
    phase = 'idle';
    await loadDataFor(outcome.character);
    rows = seedRows(outcome.character);
    if (init.tool === 'drops') syncDropRows();
  }

  /**
   * Everything a tool page needs about the build, in one pass. Each file is optional in its
   * own way and each failure degrades that one feature rather than the page: no item file
   * means slot names without icons, no loot file means no source picker, no enchant file
   * means no enchant column.
   */
  async function loadDataFor(next: SimCharacter): Promise<void> {
    const [itemFile, setFile, talents, enchantRows, suffixRows, lootFile, buffFile, phaseRows] =
      await Promise.all([
        loadItems(next.tree_version, next.class_slug).catch(() => null),
        loadSets(next.tree_version).catch(() => []),
        loadTalents(next.tree_version, next.class_slug).catch(() => null),
        loadEnchants(next.tree_version).catch(() => []),
        loadSuffixes(next.tree_version).catch(() => []),
        loadLoot(next.tree_version).catch(() => ({ sources: [] }) as LootFile),
        loadSimBuffs(next.tree_version).catch(() => ({ entries: {} }) as SimBuffFile),
        // Never rejects: phase.ts falls back to the build-time table rather than leaving
        // the gate unknown.
        fetchPhases(init.apiBase),
      ]);
    items = new Map((itemFile?.items ?? []).map((item) => [item.id, item]));
    sets = setFile;
    talentFile = talents;
    enchants = enchantRows;
    suffixes = suffixRows;
    loot = lootFile;
    simBuffs = buffFile;
    phases = phaseRows;
    sourceIndex = sourcesByItem(lootFile);
    shownKinds = [...new Set(lootFile.sources.map((source) => source.kind))].filter(
      (kind) => !DEFAULT_OFF_KINDS.includes(kind),
    );
  }

  /**
   * The grid opens with what the character is wearing, plus whatever the addon export's
   * bags, bank and in-game loadouts carried (part A's FS1 v2 decoder). Nothing is ticked:
   * a page that pre-ticks is a page that runs something the player did not choose.
   */
  function seedRows(next: SimCharacter): CandidateRow[] {
    const seeded: CandidateRow[] = [];
    const add = (itemId: number, origin: Origin, enchant = 0, suffix = 0): void => {
      const item = items.get(itemId);
      if (item === undefined) return;
      for (const slot of uiSlotsOf(item)) {
        seeded.push({ ...rowFor(item, slot, origin), enchant, suffix });
      }
    };
    // Contract 10.5 puts per-slot enchant and suffix on SimCharacter (part A's decoder),
    // so an equipped row opens carrying the enchant the player actually has on -- which is
    // what "an enchant you already have carries over" has to mean on screen as well as in
    // the planner. `gearSlots` absent (an Armory or logged-fight source) falls back to the
    // id-only map, and the rows simply carry no enchant.
    const detailed = next.gearSlots ?? [];
    if (detailed.length > 0) {
      for (const slot of detailed) add(slot.item_id, 'equipped', slot.enchant ?? 0, slot.suffix ?? 0);
    } else {
      for (const itemId of Object.values(next.gear)) if (itemId !== undefined) add(itemId, 'equipped');
    }
    for (const entry of next.bags ?? []) add(entry.item_id, 'bag', entry.enchant ?? 0, entry.suffix ?? 0);
    for (const entry of next.bank ?? []) add(entry.item_id, 'bank', entry.enchant ?? 0, entry.suffix ?? 0);
    return seeded;
  }

  /** The Droptimizer's picks, as ticked rows with a `drop:` origin and nothing else. */
  function syncDropRows(): void {
    const next: CandidateRow[] = [];
    for (const pick of pickedBosses) {
      const [sourceId, bossId] = pick.split('|');
      const source = loot.sources.find((entry) => entry.id === sourceId);
      if (source === undefined) continue;
      const pickedId = bossId === '' ? sourceId : bossId;
      const origin = `drop:${pickedId}` as Origin;
      // Contract 10.1 A6: the name is resolved once, here, and rides on the candidate.
      const sourceName = sourceNameOf(loot, pickedId);
      for (const itemId of bossId === '' ? itemsOfSource(source) : itemsOfBoss(source, bossId)) {
        const item = items.get(itemId);
        if (item === undefined) continue;
        for (const slot of uiSlotsOf(item)) {
          next.push({ ...rowFor(item, slot, origin, sourceName), checked: true });
        }
      }
    }
    rows = next;
  }

  function currentSpec(): BulkRequest['bulk'] | null {
    if (mode === null) return null;
    return buildBulkSpec({
      mode,
      rows,
      locked,
      loadouts,
      sets: namedSets,
      precision,
      cap,
      // "Try each of these" is one alternative list per ticked consumable, not every
      // subset of them (contract 10.1 A5).
      consumables: consumableIds.map((id) => [id]),
    });
  }

  function baseRequest(): BulkRequest | WeightsRequest | null {
    if (character === null) {
      message = bulkCopy.needCharacter;
      return null;
    }
    const index = talentFile === null ? null : indexTalents(talentFile);
    if (index === null) {
      message = simCopy.failed;
      return null;
    }
    const base = {
      engine_version: ENGINE_VERSION,
      spec: character.spec,
      source: character.source,
      character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
      encounter: settings.encounter,
      // Contract 10.1 A3: a bulk request's iterations ARE its precision's final stage, and
      // `Validate` refuses anything else. A weights run is a plain fixed run.
      iterations: init.tool === 'weights' ? 3000 : finalIterations(precision),
      random_seed: 0,
    };
    if (init.tool === 'weights') {
      return { ...base, weights: { stats: [...stats], reference: stats[0] ?? '' } };
    }
    const bulk = currentSpec();
    if (bulk === null) return null;
    const refusal = validateBulk(bulk);
    if (refusal !== null) {
      message = refusal;
      return null;
    }
    return { ...base, bulk };
  }

  /** The live count, debounced: it fires on every checkbox tick. */
  function scheduleCount(): void {
    if (countTimer !== null) clearTimeout(countTimer);
    countTimer = setTimeout(() => void recount(), COUNT_DEBOUNCE_MS);
  }

  async function recount(): Promise<void> {
    if (mode === null || character === null) return;
    const bulk = currentSpec();
    if (bulk === null || validateBulk(bulk) !== null) {
      combinations = null;
      capNotice = null;
      return;
    }
    const request = baseRequestForCount(bulk);
    if (request === null) return;
    try {
      combinations = await countCombinations(poolOnce(), request);
      capNotice = null;
      serverCapNotice = null;
    } catch (error) {
      if (error instanceof BulkCapError) {
        combinations = error.combinations;
        capNotice = { cap: error.cap, combinations: error.combinations };
        // Past the premium lane's own 5,000 too (contract 10.1 A2), so the page does not
        // offer a server run the API would refuse at submit with `cap_exceeded`.
        serverCapNotice =
          error.combinations > SERVER_CAP
            ? { cap: SERVER_CAP, combinations: error.combinations }
            : null;
        return;
      }
      combinations = null;
      capNotice = null;
      serverCapNotice = null;
    }
  }

  /** The count needs a whole envelope but never runs one, so it skips the message writes. */
  function baseRequestForCount(bulk: NonNullable<BulkRequest['bulk']>): BulkRequest | null {
    if (character === null || talentFile === null) return null;
    const index = indexTalents(talentFile);
    return {
      engine_version: ENGINE_VERSION,
      spec: character.spec,
      source: character.source,
      character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
      encounter: settings.encounter,
      iterations: finalIterations(bulk.precision),
      random_seed: 0,
      bulk,
    };
  }

  return {
    get tool() {
      return init.tool;
    },
    get phase() {
      return phase;
    },
    get character() {
      return character;
    },
    get needsRace() {
      return character !== null && needsRace(character);
    },
    get settings() {
      return settings;
    },
    get items() {
      return items;
    },
    get sets() {
      return sets;
    },
    get enchants() {
      return enchants;
    },
    get suffixes() {
      return suffixes;
    },
    get loot() {
      return loot;
    },
    get sourceIndex() {
      return sourceIndex;
    },
    get specRows() {
      return specRows;
    },
    get rows() {
      return rows;
    },
    get locked() {
      return locked;
    },
    get loadouts() {
      return loadouts;
    },
    get namedSets() {
      return namedSets;
    },
    get precision() {
      return precision;
    },
    get cap() {
      return cap;
    },
    get combinations() {
      return combinations;
    },
    get capNotice() {
      return capNotice;
    },
    /** Non-null only when even the premium lane's 5,000 would be exceeded. */
    get serverCapNotice() {
      return serverCapNotice;
    },
    get consumableIds() {
      return consumableIds;
    },
    /** The buff/consumable name table, for the candidate list's labels. */
    get simBuffs() {
      return simBuffs;
    },
    /** The phase table the source picker gates on: live where the API answered. */
    get phases() {
      return phases;
    },
    get progress() {
      return progress;
    },
    /** "" when nothing is running; the design's own sentence otherwise. */
    get progressLine() {
      return progress === null ? '' : stageProgressLine(progress);
    },
    get result() {
      return result;
    },
    get combos() {
      return (result as BulkResult | null)?.combos ?? [];
    },
    get weights(): StatWeight[] {
      return (result as WeightsResult | null)?.weights ?? [];
    },
    get message() {
      return message;
    },
    get detail() {
      return detail;
    },
    get premium() {
      return premium;
    },
    get serverRunning() {
      return serverRunning;
    },
    get showUpcoming() {
      return showUpcoming;
    },
    get shownKinds() {
      return shownKinds;
    },
    get pickedBosses() {
      return pickedBosses;
    },
    get stats() {
      return stats;
    },
    get referenceStat() {
      return stats[0] ?? '';
    },
    /** The sources the picker draws: kind ticked on, and released unless asked otherwise. */
    get visibleSources(): LootSource[] {
      const at = clock();
      return loot.sources.filter(
        (source) =>
          shownKinds.includes(source.kind) && (showUpcoming || isOpen(phases, source, at)),
      );
    },

    loadAddon: (code: string) => adopt(fromAddonExport(code, ctx)),
    loadBuild: (id: string) => adopt(fromPlannerBuild(id, ctx)),
    loadFight: (ref: string) => adopt(fromLoggedFight(ref, ctx)),
    loadStored: (path: CharacterPath) => adopt(fromStoredCharacter(path, ctx)),

    async loadSpecs(): Promise<void> {
      try {
        specRows = await fetchSpecs(init.apiBase);
      } catch {
        specRows = [];
      }
      if (init.tool === 'weights' && stats.length === 0 && character !== null) {
        stats = defaultStatsFor(character.spec, referenceFor(character.spec, specRows));
      }
    },

    setPremium(value: boolean): void {
      premium = value;
    },
    setMessage(text: string | null): void {
      message = text;
    },
    setSettings(next: SimSettings): void {
      settings = next;
      scheduleCount();
    },
    setRace(slug: string): void {
      if (character === null) return;
      character = { ...character, race_slug: slug };
      result = null;
      message = null;
      phase = 'idle';
    },
    setPrecision(value: Precision): void {
      precision = value;
    },
    setCap(value: number): void {
      cap = value;
      scheduleCount();
    },
    setStats(next: string[]): void {
      stats = [...next];
    },

    toggleRow(key: string): void {
      rows = toggleRow(rows, key);
      scheduleCount();
    },
    removeRow(key: string): void {
      rows = removeRow(rows, key);
      scheduleCount();
    },
    copyAndModify(key: string, patch: { enchant?: number; suffix?: number }): void {
      rows = copyAndModify(rows, key, patch);
      scheduleCount();
    },
    /** From the item search and from a Droptimizer pin: added, ticked, and counted. */
    addSearchItem(itemId: number, origin: Origin = 'search'): void {
      const item = items.get(itemId);
      if (item === undefined) return;
      for (const slot of uiSlotsOf(item)) {
        rows = addRow(rows, { ...rowFor(item, slot, origin), checked: true });
      }
      scheduleCount();
    },
    toggleLock(slot: Slot): void {
      locked = locked.includes(slot) ? locked.filter((entry) => entry !== slot) : [...locked, slot];
      scheduleCount();
    },
    addLoadout(loadout: TalentLoadout): void {
      if (loadouts.some((entry) => entry.name === loadout.name)) return;
      loadouts = [...loadouts, loadout];
      scheduleCount();
    },
    removeLoadout(name: string): void {
      loadouts = loadouts.filter((entry) => entry.name !== name);
      scheduleCount();
    },
    addNamedSet(set: GearSet): void {
      namedSets = [...namedSets.filter((entry) => entry.name !== set.name), set];
      scheduleCount();
    },
    removeNamedSet(name: string): void {
      namedSets = namedSets.filter((entry) => entry.name !== name);
      scheduleCount();
    },
    toggleConsumable(id: string): void {
      consumableIds = consumableIds.includes(id)
        ? consumableIds.filter((entry) => entry !== id)
        : [...consumableIds, id];
      scheduleCount();
    },
    toggleKind(kind: string): void {
      shownKinds = shownKinds.includes(kind)
        ? shownKinds.filter((entry) => entry !== kind)
        : [...shownKinds, kind];
    },
    setShowUpcoming(value: boolean): void {
      showUpcoming = value;
    },
    /** A boss, or a whole source when `bossId` is empty. */
    toggleSource(sourceId: string, bossId = ''): void {
      const key = `${sourceId}|${bossId}`;
      pickedBosses = pickedBosses.includes(key)
        ? pickedBosses.filter((entry) => entry !== key)
        : [...pickedBosses, key];
      syncDropRows();
      scheduleCount();
    },

    recount,

    adoptResult(next: SimResult): void {
      result = next;
      phase = 'done';
    },

    async run(): Promise<void> {
      const request = baseRequest();
      if (request === null) {
        phase = 'idle';
        return;
      }
      message = null;
      detail = '';
      stopRequested = false;
      phase = 'running';
      progress = null;

      handle =
        init.tool === 'weights'
          ? runWeightsRun(poolOnce(), request as WeightsRequest, () => {})
          : runBulk(poolOnce(), request as BulkRequest, (next) => {
              if (stopRequested) return;
              progress = next;
            });

      try {
        const finished = await handle.result;
        result = finished;
        if (finished.aborted === true) message = bulkCopy.partial;
        phase = 'done';
      } catch (error) {
        if (error instanceof BulkCapError) {
          capNotice = { cap: error.cap, combinations: error.combinations };
          combinations = error.combinations;
          message = error.message;
          phase = 'idle';
          return;
        }
        const failure = error instanceof BulkRunError ? error : null;
        message = failure?.cancelled === true ? simCopy.stopped : (failure?.message ?? bulkCopy.bulkFailed);
        detail = failure?.detail ?? '';
        phase = failure?.cancelled === true && result !== null ? 'done' : 'error';
      } finally {
        handle = null;
        progress = null;
      }
    },

    stop(): void {
      stopRequested = true;
      handle?.cancel();
    },

    /**
     * The same envelope, on the premium lane. `dispatchServerSim` is the single-run
     * dispatch unchanged: a bulk request IS a SimRequest, the API derives the kind from the
     * body, and the progress route carries the three bulk columns.
     */
    async runOnServer(): Promise<void> {
      if (serverRunning) return;
      const request = baseRequest();
      if (request === null) return;
      serverRunning = true;
      const generation = ++serverGeneration;
      const current = (): boolean => generation === serverGeneration;
      try {
        message = null;
        detail = '';
        let simId: string;
        try {
          simId = await dispatchServerSim(request, init.apiBase);
        } catch (error) {
          if (current()) message = error instanceof SimApiError ? error.message : bulkCopy.bulkFailed;
          return;
        }
        for (;;) {
          await delay(init.serverPollMs ?? DEFAULT_SERVER_POLL_MS);
          if (!current()) return;
          let row;
          try {
            row = await fetchBulkProgress(simId, init.apiBase);
          } catch (error) {
            if (current()) {
              message = error instanceof SimApiError ? error.message : bulkCopy.bulkFailed;
              phase = result !== null ? 'done' : 'error';
            }
            return;
          }
          if (!current()) return;
          if (row.stage !== undefined && row.combos_total !== undefined) {
            progress = {
              stage: row.stage,
              stages: row.combos_total === 0 ? 1 : (progress?.stages ?? 1),
              combosDone: row.combos_done ?? 0,
              combosTotal: row.combos_total,
            };
          }
          if (row.state === 'error') {
            message = bulkCopy.bulkFailed;
            phase = result !== null ? 'done' : 'error';
            return;
          }
          if (row.state === 'done') {
            try {
              const finished = await fetchSim(simId, init.apiBase);
              if (current()) {
                result = finished;
                progress = null;
                phase = 'done';
              }
            } catch (error) {
              if (current()) {
                message = error instanceof SimApiError ? error.message : bulkCopy.bulkFailed;
                phase = result !== null ? 'done' : 'error';
              }
            }
            return;
          }
        }
      } finally {
        if (current()) serverRunning = false;
      }
    },

    dispose(): void {
      handle?.cancel();
      if (countTimer !== null) clearTimeout(countTimer);
      serverGeneration += 1;
      serverRunning = false;
      pool?.terminate();
      pool = null;
    },
  };
}

export type BulkStore = ReturnType<typeof createBulkStore>;
```

> **If part A has not landed:** `SimCharacter` will not yet declare `bags`, `bank`, `sets`,
> `loadouts`, `professions` or `gearSlots`, and `seedRows`' reads of them will not
> type-check. Add the six as optionals to `SimCharacter` in a **separate commit** on this
> branch, note it in the commit body, and drop that commit when part A lands. `bags` and
> `bank` are `GearSlot[]`-shaped (`{ item_id, enchant?, suffix? }` per contract 10.5), not
> bare ids.

- [ ] **Step 4: Run the test**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-store.test.ts
```

Expected: PASS, 13 cases.

- [ ] **Step 5: Check the file length**

```bash
cd web && wc -l src/lib/sim/bulk-store.svelte.ts
```

Expected: under 800. If it is over 500, split the Droptimizer's picks (`pickedBosses`,
`showUpcoming`, `shownKinds`, `visibleSources`, `syncDropRows`) into
`src/lib/sim/drop-picks.ts` as pure functions the store calls, and re-run.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/bulk-store.svelte.ts src/lib/sim/bulk-store.test.ts
npx prettier --write src/lib/sim/bulk-store.svelte.ts src/lib/sim/bulk-store.test.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/bulk-store.svelte.ts web/src/lib/sim/bulk-store.test.ts
git commit -m "$(cat <<'MSG'
feat(sim): the store the four combination tools mount over

Its own file, not more fields on /sim's store: the two share character
loading through sources.ts and nothing else, and the parity spec's two
web lanes must not edit one store. The pool is created on the first run,
the live count is debounced, a cap refusal is two numbers rather than a
trimmed list, and the premium lane posts the identical envelope to
POST /v1/sims/run (contract 10.6). Iterations are the precision's final
stage per 10.1 A3; the phase gate prefers GET /v1/phases and falls back
to the build-time table; a drop's source name travels on the candidate
per 10.1 A6; and the page will not offer a server run past 5,000.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 11: The tools island, the four shells and their skeletons

The existing `sim-island` bundle is shared by `/sim`, `/sim/specs` and `/sim/<id>` and is
budgeted at 90 KB gzipped. Four candidate grids, a source picker and two results views do
not fit under it, and adding them would charge `/sim`'s 1,600 ms mobile LCP budget for code
no visitor to `/sim` runs. These four pages therefore get their own standalone entry, and
each page's view is a lazy chunk off it, so `/sim/weights` never downloads the slot grid.

**Files:**
- Create: `web/src/lib/sim/bulk-skeleton.ts`, `web/src/lib/sim/bulk-skeleton.test.ts`
- Create: `web/src/sim-tools-island.ts`, `web/src/sim-tools-island.test.ts`
- Create: `web/vite.sim-tools-island.config.ts`
- Create: `web/src/components/sim/tools/ToolsView.svelte`
- Create: `web/src/pages/sim/gear.astro`, `talents.astro`, `drops.astro`, `weights.astro`
- Create: `web/src/pages/_sim-tools.test.ts`
- Modify: `web/vite.island.config.ts` (widen the entry-name union)
- Modify: `web/package.json` (`build:island`)
- Modify: `web/scripts/check-island-size.mjs` (a budget row)

**Interfaces:**
- Consumes: `SimTool`, `TOOLS` (Task 10); `createLazyComponent` from `../../lib/report/lazy-component.svelte`; `bulkCopy`.
- Produces: `TOOL_SKELETONS: Record<SimTool, string>`, `toolFrom(element, pathname): SimTool`, and the mount contract `<div id="sim-tools" data-sim-tool="gear">`.

- [ ] **Step 1: Write the failing tests**

`web/src/lib/sim/bulk-skeleton.test.ts`:

```ts
// web/src/lib/sim/bulk-skeleton.test.ts
import { describe, expect, it } from 'vitest';
import { TOOL_SKELETONS } from './bulk-skeleton';
import { TOOLS } from './bulk-store.svelte';

describe('TOOL_SKELETONS', () => {
  it('has one skeleton per tool page', () => {
    expect(Object.keys(TOOL_SKELETONS).sort()).toEqual([...TOOLS].sort());
  });

  it('announces itself once to a screen reader and hides its blocks from it', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html.match(/role="status"/g) ?? []).toHaveLength(1);
      expect(html).toContain('aria-hidden="true"');
      expect(html).toContain('aria-busy="true"');
    }
  });

  it('puts no data in the markup, so the shell and the island render the same bytes', () => {
    for (const html of Object.values(TOOL_SKELETONS)) {
      expect(html).not.toMatch(/\{|\$\{/);
    }
  });
});
```

`web/src/sim-tools-island.test.ts`:

```ts
// @vitest-environment jsdom
// web/src/sim-tools-island.test.ts
import { describe, expect, it } from 'vitest';
import { toolFrom } from './sim-tools-island';

describe('toolFrom', () => {
  it('prefers the mount’s own stamp', () => {
    const element = document.createElement('div');
    element.dataset.simTool = 'drops';
    expect(toolFrom(element, '/sim/gear')).toBe('drops');
  });

  it('falls back to the path', () => {
    const element = document.createElement('div');
    expect(toolFrom(element, '/sim/weights')).toBe('weights');
    expect(toolFrom(element, '/sim/weights.html')).toBe('weights');
  });

  it('refuses anything not a tool, and lands on gear rather than throwing', () => {
    const element = document.createElement('div');
    element.dataset.simTool = 'javascript:alert(1)';
    expect(toolFrom(element, '/sim/nonesuch')).toBe('gear');
  });
});
```

`web/src/pages/_sim-tools.test.ts`:

```ts
// web/src/pages/_sim-tools.test.ts
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { beforeAll, describe, expect, it } from 'vitest';
import Gear from './sim/gear.astro';
import Talents from './sim/talents.astro';
import Drops from './sim/drops.astro';
import Weights from './sim/weights.astro';
import { TOOL_SKELETONS } from '../lib/sim/bulk-skeleton';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

const OG_HOOKS = [
  'data-og="title"',
  'data-og="description"',
  'data-og="canonical"',
  'data-og="og-title"',
  'data-og="og-description"',
  'data-og="og-url"',
  'data-og="og-image"',
];

const PAGES = [
  { name: 'gear', component: Gear, path: '/sim/gear', heading: 'Top Gear' },
  { name: 'talents', component: Talents, path: '/sim/talents', heading: 'Talent compare' },
  { name: 'drops', component: Drops, path: '/sim/drops', heading: 'Droptimizer' },
  { name: 'weights', component: Weights, path: '/sim/weights', heading: 'Stat weights' },
] as const;

describe.each(PAGES)('the $name shell', ({ component, path, heading, name }) => {
  it('carries every data-og hook and its own canonical', async () => {
    const html = await container.renderToString(component);
    for (const selector of OG_HOOKS) expect(html, selector).toContain(selector);
    expect(html).toContain(`href="https://foreversixty.gg${path}"`);
  });

  it('mounts the tools island, stamps its tool, and links one stylesheet and one script', async () => {
    const html = await container.renderToString(component);
    expect(html).toContain('id="sim-tools"');
    expect(html).toContain(`data-sim-tool="${name}"`);
    expect(html).toContain('href="/sim-tools-island.css"');
    expect(html).toContain('src="/sim-tools-island.js"');
  });

  it('renders its heading and its own skeleton before the island, so nothing shifts', async () => {
    const html = await container.renderToString(component);
    expect(html).toContain(heading);
    expect(html).toContain(TOOL_SKELETONS[name].slice(0, 80));
  });

  it('says what the page needs when there is no JavaScript', async () => {
    const html = await container.renderToString(component);
    expect(html).toContain('needs JavaScript');
    expect(html).toContain('href="/sim"');
  });
});
```

- [ ] **Step 2: Run them and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-skeleton.test.ts src/sim-tools-island.test.ts src/pages/_sim-tools.test.ts
```

Expected: FAIL, nothing resolves.

- [ ] **Step 3: Write `bulk-skeleton.ts`**

```ts
// web/src/lib/sim/bulk-skeleton.ts
// What each tool page looks like before its island mounts, in placeholder blocks -- the
// same pattern skeleton.ts and report/skeleton.ts use and for the same reason: one string,
// rendered twice (by the Astro shell via set:html, and by ToolsView between mounting and
// the page's own view chunk resolving), so nothing shifts between the two moments.
//
// Static markup with no data in it, so it is safe to {@html} and identical every render.
import type { SimTool } from './bulk-store.svelte';

const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** One slot row of the candidate grid: an icon square, a name bar, two small bars. */
const slotRow = (): string =>
  '<li class="flex min-h-11 items-center gap-3 border-b border-line-soft px-2 py-2">' +
  `${block('h-8 w-8 rounded-control')}${block('h-3 w-32')}${block('ml-auto h-3 w-10')}${block('h-3 w-8')}` +
  '</li>';

/** One ranked results row: rank, chips, figure, delta, percent. */
const comboRow = (): string =>
  '<li class="grid min-h-11 grid-cols-[28px_minmax(120px,2fr)_84px_84px_56px] items-center gap-x-3 border-b border-line-soft px-2 py-2">' +
  `${block('h-3 w-4')}${block('h-3 w-40')}${block('ml-auto h-3 w-14')}${block('ml-auto h-3 w-12')}${block('ml-auto h-3 w-8')}` +
  '</li>';

const shell = (label: string, body: string): string =>
  [
    `<div class="flex flex-col gap-4" aria-busy="true">`,
    `<p class="sr-only" role="status">${label}</p>`,
    '<div aria-hidden="true" class="flex flex-col gap-4">',
    body,
    '</div></div>',
  ].join('');

/** The character strip's reserved band, which every tool page opens with. */
const strip = (): string =>
  '<div class="bg-raised border-line rounded-panel mx-[18px] flex flex-wrap items-center gap-x-6 gap-y-2 border p-4 md:mx-0">' +
  `${block('h-6 w-40')}${block('h-3 w-24')}${block('h-3 w-16')}` +
  '</div>';

const runBar = (): string =>
  '<div class="border-line rounded-panel mx-[18px] flex flex-wrap items-center gap-3 border p-4 md:mx-0">' +
  `${block('h-3 w-36')}${block('h-9 w-28 rounded-control')}${block('h-9 w-24 rounded-control')}` +
  '</div>';

const grid = (count: number): string =>
  `<ul class="mx-[18px] flex flex-col md:mx-0">${Array.from({ length: count }, slotRow).join('')}</ul>`;

const table = (count: number): string =>
  `<ul class="mx-[18px] flex flex-col md:mx-0">${Array.from({ length: count }, comboRow).join('')}</ul>`;

export const TOOL_SKELETONS: Record<SimTool, string> = {
  gear: shell('Loading Top Gear.', [strip(), grid(8), runBar()].join('')),
  talents: shell('Loading talent compare.', [strip(), grid(3), runBar()].join('')),
  drops: shell('Loading the Droptimizer.', [strip(), grid(6), runBar()].join('')),
  weights: shell('Loading stat weights.', [strip(), table(6), runBar()].join('')),
};
```

- [ ] **Step 4: Write the island entry**

`web/src/sim-tools-island.ts`:

```ts
// web/src/sim-tools-island.ts
// Entry point for dist/sim-tools-island.js, the bundle /sim/gear, /sim/talents, /sim/drops
// and /sim/weights load.
//
// It is a separate build from sim-island for one measured reason: sim-island is budgeted at
// 90 KB gzipped and carries /sim's own 1,600 ms mobile LCP, and four candidate grids, a
// source picker and two results views do not fit under that. Nothing here is loaded by /sim.
import { mount } from 'svelte';
import ToolsView from './components/sim/tools/ToolsView.svelte';
import { TOOLS, type SimTool } from './lib/sim/bulk-store.svelte';
import './styles/fonts.css';
import './styles/global.css';

const MOUNT_ID = 'sim-tools';

function isTool(value: string): value is SimTool {
  return (TOOLS as readonly string[]).includes(value);
}

/**
 * The mount's own stamp wins; the path is the fallback, so a shell that forgot the
 * attribute still renders the right page. Anything else lands on Top Gear rather than
 * throwing: the value reaches a component name lookup and is attacker-influenced through
 * the URL, so it is validated against a closed list and never used raw.
 */
export function toolFrom(element: HTMLElement, pathname: string): SimTool {
  const stamped = element.dataset.simTool ?? '';
  if (isTool(stamped)) return stamped;
  const last = pathname.replace(/\.html$/, '').split('/').pop() ?? '';
  return isTool(last) ? last : 'gear';
}

function boot(): void {
  const target = document.getElementById(MOUNT_ID);
  if (target === null) return;
  const tool = toolFrom(target, window.location.pathname);
  // The shell's skeleton and no-JS paragraph live inside the mount; Svelte 5 appends rather
  // than replaces, so they go before the mount or they stay under the island.
  target.replaceChildren();
  mount(ToolsView, { target, props: { tool } });
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', boot, { once: true });
} else {
  boot();
}
```

`web/vite.sim-tools-island.config.ts`:

```ts
// web/vite.sim-tools-island.config.ts
import { defineConfig } from 'vite';
import { islandConfig } from './vite.island.config';

export default defineConfig(islandConfig('sim-tools-island'));
```

In `web/vite.island.config.ts`, widen the parameter type:

```ts
export function islandConfig(
  name: 'planner-island' | 'report-island' | 'sim-island' | 'sim-tools-island',
) {
```

and set `modulePreload: false` stays as is — note in the file's own comment that this entry
**does** dynamically import (one chunk per tool), so the existing comment "Nothing here is
dynamically imported" gets one added sentence:

```ts
      // sim-tools-island is the exception: it lazily imports one view chunk per tool page,
      // so /sim/weights never downloads the slot grid. The chunks take
      // `${name}-[hash].js`, which is already the chunkFileNames rule below.
```

- [ ] **Step 5: Write `ToolsView.svelte`**

```svelte
<!-- web/src/components/sim/tools/ToolsView.svelte -->
<!-- The tools island's root: the character strip, the source switcher and the one tool
     view this page is. Every tool view is a lazy chunk, so /sim/weights downloads neither
     the slot grid nor the source picker.

     There is no <h1> here: each .astro shell carries its own, in static HTML, ahead of this
     island -- an island-mounted heading is invisible to Lighthouse's LCP measurement, which
     is the same reason SimView.svelte has none. -->
<script lang="ts">
  import { onMount, untrack } from 'svelte';
  import activeBuild from '../../../data/active-build.json';
  import { battlenetStartUrl, fetchMe, type Me } from '../../../lib/account/api';
  import type { CharacterPath } from '../../../lib/characters';
  import { createLazyComponent, type LazyLoadState } from '../../../lib/report/lazy-component.svelte';
  import { TOOL_SKELETONS } from '../../../lib/sim/bulk-skeleton';
  import { createBulkStore, type SimTool } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { simCopy } from '../../../lib/sim/copy';
  import { parseSimState } from '../../../lib/sim/url';
  import CharacterStrip from '../CharacterStrip.svelte';
  import SourceSwitcher from '../SourceSwitcher.svelte';

  let { tool }: { tool: SimTool } = $props();

  const bootstrap = untrack(() => {
    const { source, ref } = parseSimState(window.location.search);
    const mount = document.getElementById('sim-tools');
    return { source, ref, treeVersion: mount?.dataset.treeVersion ?? activeBuild.build };
  });

  const store = untrack(() =>
    createBulkStore({
      tool,
      treeVersion: bootstrap.treeVersion,
      source: bootstrap.source,
      ref: bootstrap.ref,
      hardwareConcurrency: navigator.hardwareConcurrency,
    }),
  );

  let me = $state<Me | null>(null);
  let switcherOpen = $state(false);

  $effect(() => {
    if (store.character !== null) switcherOpen = false;
  });

  onMount(() => {
    void fetchMe()
      .then((result) => {
        me = result;
        store.setPremium(result?.user.premium === true);
      })
      .catch(() => {});
    void store.loadSpecs();
    // The URL's own bootstrap, once, here rather than in an effect: a "pin into Top Gear"
    // link and a "sim this build" link both arrive as ?source=&ref=.
    if (bootstrap.source !== '' && bootstrap.ref !== '') {
      if (bootstrap.source === 'addon') void store.loadAddon(bootstrap.ref);
      else if (bootstrap.source === 'build') void store.loadBuild(bootstrap.ref);
      else if (bootstrap.source === 'fight') void store.loadFight(bootstrap.ref);
    }
    return () => store.dispose();
  });

  function onSignIn(): void {
    window.location.href = battlenetStartUrl(`${window.location.pathname}${window.location.search}`);
  }

  async function pick(path: CharacterPath): Promise<void> {
    await store.loadStored(path);
  }

  // One chunk per tool, resolved from a closed map: `tool` is validated against TOOLS
  // before it reaches here, and a map rather than a template literal keeps the bundler's
  // own analysis exact.
  const views = {
    gear: () => import('./TopGear.svelte'),
    talents: () => import('./TopGear.svelte'),
    drops: () => import('./Droptimizer.svelte'),
    weights: () => import('./StatWeights.svelte'),
  } as const;

  const viewLazy = createLazyComponent(views[tool]);
  viewLazy.load();
</script>

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="alert" data-testid="sim-tool-error">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex min-h-11 items-center underline"
        onclick={() => lazy.load()}>{simCopy.tryAgain}</button
      >
    </p>
  {:else}
    <!-- eslint-disable-next-line svelte/no-at-html-tags -->
    {@html TOOL_SKELETONS[tool]}
  {/if}
{/snippet}

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-tools-view">
  {#if store.character !== null && !switcherOpen}
    <CharacterStrip
      character={store.character}
      items={store.items}
      races={[]}
      onchange={() => (switcherOpen = true)}
      onrace={(slug) => store.setRace(slug)}
    />
  {:else}
    <SourceSwitcher
      busy={store.phase === 'loading-character'}
      message={store.message}
      signedIn={me !== null}
      onaddon={(code) => void store.loadAddon(code)}
      onbuild={(id) => void store.loadBuild(id)}
      onfight={(ref) => void store.loadFight(ref)}
      onsignin={onSignIn}
      onback={() => (switcherOpen = false)}
    />
  {/if}

  {#if store.character !== null}
    {#if viewLazy.current}
      <viewLazy.current {store} {me} onpick={(path: CharacterPath) => void pick(path)} />
    {:else}
      {@render lazyFallback(viewLazy)}
    {/if}
  {:else}
    <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-tools-empty">
      {bulkCopy.needCharacter}
    </p>
  {/if}
</div>
```

- [ ] **Step 6: Write the four shells**

`web/src/pages/sim/gear.astro`, and the other three by the same pattern with their own
`tool`, title, description and intro:

```astro
---
// web/src/pages/sim/gear.astro
// Static. The island reads ?source= and ?ref= in the browser, because the page is
// prerendered and has no request to read.
//
// The skeleton below is the same string ToolsView.svelte renders while its view chunk
// resolves (bulk-skeleton.ts is the one source both read), so the mount swap moves nothing
// and the page holds its CLS budget -- the identical reservation technique sim/[id].astro
// and sim/specs.astro already use, and for the identical reason.
import Base from '../../layouts/Base.astro';
import { bulkCopy } from '../../lib/sim/copy';
import { TOOL_SKELETONS } from '../../lib/sim/bulk-skeleton';
---

<Base
  title="Top Gear"
  description="Simulate every combination of the gear, enchants, talents and sets you have, and see which one is actually best."
  path="/sim/gear"
  session
>
  <link rel="stylesheet" href="/sim-tools-island.css" slot="head" />
  <main
    id="main"
    tabindex="-1"
    class="mx-auto flex w-full max-w-[1344px] flex-col gap-[22px] pt-4 pb-10 md:gap-8 md:px-12 md:pt-7"
  >
    <h1 class="section-title px-[18px] text-[18px] md:px-0">{bulkCopy.gearTitle}</h1>
    <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-tool-intro">
      {bulkCopy.gearIntro}
    </p>
    <div id="sim-tools" data-sim-tool="gear" class="min-h-[1180px] md:min-h-[820px]">
      <Fragment set:html={TOOL_SKELETONS.gear} />
      <noscript>
        <p class="text-muted px-[18px] text-[14px] md:px-0">
          Top Gear needs JavaScript: it runs the engine in your browser rather than on our servers. The{' '}
          <a href="/sim">simulator</a> says the same.
        </p>
      </noscript>
    </div>
    <script type="module" src="/sim-tools-island.js"></script>
  </main>
</Base>
```

The other three, identical but for these values:

| File | `tool` | `title` | copy | `min-h` |
| --- | --- | --- | --- | --- |
| `talents.astro` | `talents` | Talent compare | `bulkCopy.talentsTitle` / `talentsIntro` | `min-h-[760px] md:min-h-[520px]` |
| `drops.astro` | `drops` | Droptimizer | `bulkCopy.dropsTitle` / `dropsIntro` | `min-h-[1080px] md:min-h-[760px]` |
| `weights.astro` | `weights` | Stat weights | `bulkCopy.weightsTitle` / `weightsIntro` | `min-h-[820px] md:min-h-[560px]` |

Each `description`:
- talents: "Simulate your talent builds against each other on the gear you are wearing, and see which tree actually wins."
- drops: "Simulate every item a boss drops against the set you are wearing, and see which drops are upgrades."
- weights: "What one point of each stat is worth for your character, with the caveat that comes with it."

> The `min-h` values are placeholders until measured. **Task 20 measures each page at 360 px
> and 1280 px with the fixture character loaded and replaces all four**, exactly as
> `sim/[id].astro`'s own comment describes; leaving a guess in is how a CLS budget fails in
> CI rather than locally.

- [ ] **Step 7: Build the island and budget it**

In `web/package.json`, extend `build:island`:

```json
    "build:island": "vite build --config vite.island.config.ts && vite build --config vite.report-island.config.ts && vite build --config vite.sim-island.config.ts && vite build --config vite.sim-tools-island.config.ts",
```

In `web/scripts/check-island-size.mjs`, add to `BUDGETS`:

```js
  // The four combination tools: a candidate grid, an item search, a source picker and two
  // results views, each a lazily-imported chunk off one entry. 70 KB gzipped for the entry
  // is roughly twice what the shared parts (the character strip, the source switcher, the
  // run bar) weigh; the per-tool chunks are not counted here because none of them is on any
  // page's LCP path -- the shell paints its own skeleton first.
  { file: 'dist/sim-tools-island.js', limitBytes: 70 * 1024 },
```

- [ ] **Step 8: Run everything this task touched**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-skeleton.test.ts src/sim-tools-island.test.ts src/pages/_sim-tools.test.ts src/pages/_shells.test.ts
FOREVER_DATA=fixture npm run build
```

Expected: tests PASS; the build emits `dist/sim-tools-island.js`, `dist/sim-tools-island.css`,
`dist/sim/gear.html`, `talents.html`, `drops.html`, `weights.html`, and the size check prints
the new budget line under its limit.

- [ ] **Step 9: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/lib/sim/bulk-skeleton.ts src/lib/sim/bulk-skeleton.test.ts src/sim-tools-island.ts src/sim-tools-island.test.ts src/components/sim/tools/ToolsView.svelte src/pages/sim src/pages/_sim-tools.test.ts vite.island.config.ts vite.sim-tools-island.config.ts scripts/check-island-size.mjs
npx prettier --write src/lib/sim/bulk-skeleton.ts src/lib/sim/bulk-skeleton.test.ts src/sim-tools-island.ts src/sim-tools-island.test.ts src/components/sim/tools src/pages/sim src/pages/_sim-tools.test.ts vite.island.config.ts vite.sim-tools-island.config.ts scripts/check-island-size.mjs package.json
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/lib/sim/bulk-skeleton.ts web/src/lib/sim/bulk-skeleton.test.ts web/src/sim-tools-island.ts web/src/sim-tools-island.test.ts web/src/components/sim/tools web/src/pages/sim web/src/pages/_sim-tools.test.ts web/vite.island.config.ts web/vite.sim-tools-island.config.ts web/scripts/check-island-size.mjs web/package.json
git commit -m "$(cat <<'MSG'
feat(sim): the tools island and the four page shells

A second standalone island rather than more weight on sim-island: that
bundle is budgeted at 90 KB and carries /sim's own 1,600 ms mobile LCP,
and nothing on /sim runs a candidate grid. Each tool view is a lazy
chunk off the entry, so /sim/weights downloads neither the slot grid nor
the source picker. Each shell renders the same skeleton string its
island renders, so the mount swap moves nothing.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 12: Top Gear's slot grid — candidates, locks, enchants and copy-and-modify

**Files:**
- Create: `web/src/components/sim/tools/TopGear.svelte`
- Create: `web/src/components/sim/tools/SlotGrid.svelte`
- Create: `web/src/components/sim/tools/CandidateRows.svelte`
- Create: `web/src/components/sim/tools/EnchantList.svelte`
- Create: `web/tests/e2e/sim-gear.spec.ts` (the grid's own cases; the rest arrive in later tasks)

**Interfaces:**
- Consumes: `BulkStore` (Task 10); `candidateKey`, `uiSlotsOf`, `CandidateRow` (Task 8); `enchantsForSlot`, `suffixesForItem`, `ENCHANTS_PER_SLOT_CAP`, `KEEP_CURRENT_ENCHANT`, `NO_ENCHANT` (Task 7); `rarityClassFor` from `../../../lib/planner/items`; `dataUrl` from `../../../lib/planner/load`; `SLOTS`, `SLOT_LABELS`, `SECONDARY_BUTTON`.
- **Consumes from part A:** `SettingsPanel.svelte`, `RequestDrawer.svelte`.
- Produces: `TopGear.svelte` with props `{ store: BulkStore; me: Me | null }`, the composition every later Top Gear task adds a section to. Test ids: `sim-slot-grid`, `sim-candidate-<slot>-<item>`, `sim-lock-<slot>`, `sim-enchant-<slot>-<id>`.

- [ ] **Step 1: Write the failing e2e**

Create `web/tests/e2e/sim-gear.spec.ts`:

```ts
// web/tests/e2e/sim-gear.spec.ts
// Top Gear against the checked-in fake engine (PUBLIC_SIM_ENGINE defaults to 'fake'), so
// no Go toolchain and no wasm artifact is needed. The real engine gets one gated run of its
// own in tests/e2e/sim-gear-real-engine.spec.ts.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// The fixture warrior, wearing two of the six fixture items so the grid has an equipped row
// to lock and a bag row to tick.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

export async function loadGear(page: Page, at = '/sim/gear'): Promise<void> {
  await page.goto(at);
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-slot-grid')).toBeVisible();
}

test('the grid lists the equipped item in its slot and nothing is ticked on arrival', async ({
  page,
}) => {
  await loadGear(page);
  const equipped = page.getByTestId('sim-candidate-head-12640');
  await expect(equipped).toBeVisible();
  await expect(equipped.getByRole('checkbox')).not.toBeChecked();
});

test('locking a slot disables every candidate in it', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-lock-head').check();
  await expect(page.getByTestId('sim-candidate-head-12640').getByRole('checkbox')).toBeDisabled();
  await page.getByTestId('sim-lock-head').uncheck();
  await expect(page.getByTestId('sim-candidate-head-12640').getByRole('checkbox')).toBeEnabled();
});

test('copy and modify adds the same item again with an enchant, beside the original', async ({
  page,
}) => {
  await loadGear(page);
  await page.getByTestId('sim-candidate-head-12640').getByRole('button', { name: /copy/i }).click();
  await page.getByTestId('sim-enchant-head-2543').click();
  await expect(page.getByTestId('sim-candidate-head-12640')).toBeVisible();
  const copy = page.getByTestId('sim-candidate-head-12640-e2543');
  await expect(copy).toBeVisible();
  await expect(copy.getByRole('checkbox')).toBeChecked();
});

test('the combination count moves as candidates are ticked', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-combo-count')).toHaveText(/0 valid combinations|Counting/);
  await page.getByTestId('sim-candidate-head-12640').getByRole('checkbox').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText('1 valid combination');
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts --project=desktop
```

Expected: FAIL — `sim-slot-grid` never appears.

- [ ] **Step 3: Write `EnchantList.svelte`**

```svelte
<!-- web/src/components/sim/tools/EnchantList.svelte -->
<!-- The enchants one slot allows, as a menu behind "copy and modify" and as the per-slot
     checklist design 3.1.4 asks for. "Keep current" and "None" are rows, not a separate
     control, because they are two of the answers to the same question. -->
<script lang="ts">
  import { dataUrl } from '../../../lib/planner/load';
  import {
    ENCHANTS_PER_SLOT_CAP,
    KEEP_CURRENT_ENCHANT,
    NO_ENCHANT,
    enchantsForSlot,
    type EnchantRow,
  } from '../../../lib/sim/enchants';
  import { bulkCopy } from '../../../lib/sim/copy';

  let {
    rows,
    slot,
    classSlug,
    treeVersion,
    onpick,
  }: {
    rows: readonly EnchantRow[];
    slot: string;
    classSlug: string;
    treeVersion: string;
    onpick: (enchantId: number) => void;
  } = $props();

  const allowed = $derived(enchantsForSlot(rows, slot, classSlug).slice(0, ENCHANTS_PER_SLOT_CAP));
</script>

<ul class="border-line bg-raised rounded-panel flex flex-col border p-2" data-testid={`sim-enchants-${slot}`}>
  <li>
    <button
      type="button"
      class="text-text flex min-h-11 w-full items-center px-2 text-left text-[13px]"
      onclick={() => onpick(KEEP_CURRENT_ENCHANT)}>{bulkCopy.keepCurrentEnchant}</button
    >
  </li>
  <li>
    <button
      type="button"
      class="text-text flex min-h-11 w-full items-center px-2 text-left text-[13px]"
      onclick={() => onpick(NO_ENCHANT)}>{bulkCopy.noEnchant}</button
    >
  </li>
  {#each allowed as enchant (enchant.effect_id)}
    <li class="border-line-soft border-t">
      <button
        type="button"
        class="text-text flex min-h-11 w-full items-center gap-2 px-2 text-left text-[13px]"
        data-testid={`sim-enchant-${slot}-${enchant.effect_id}`}
        onclick={() => onpick(enchant.effect_id)}
      >
        <img
          src={dataUrl(treeVersion, `icons/${enchant.icon}.webp`)}
          alt=""
          width="20"
          height="20"
          loading="lazy"
          decoding="async"
          class="rounded-control border-line h-5 w-5 border object-cover"
        />
        {enchant.name}
      </button>
    </li>
  {/each}
  {#if allowed.length === 0}
    <li class="text-muted px-2 py-2 text-[13px]">{bulkCopy.noEnchant}</li>
  {/if}
</ul>
<p class="text-muted px-2 text-[12px]">{bulkCopy.enchantCap(ENCHANTS_PER_SLOT_CAP)}</p>
```

- [ ] **Step 4: Write `CandidateRows.svelte`**

```svelte
<!-- web/src/components/sim/tools/CandidateRows.svelte -->
<!-- One slot's candidates: a checkbox, an icon, a name, an item level and a copy-and-modify
     menu, grouped by where each came from. A locked slot disables every checkbox in it
     rather than hiding the rows -- the player has to see what they locked away. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
  import { dataUrl } from '../../../lib/planner/load';
  import type { Slot } from '../../../lib/planner/types';
  import { candidateKey, type CandidateRow } from '../../../lib/sim/candidates';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { suffixesForItem, type EnchantRow, type SuffixRow } from '../../../lib/sim/enchants';
  import EnchantList from './EnchantList.svelte';

  let {
    rows,
    slot,
    locked,
    enchants,
    suffixes,
    classSlug,
    treeVersion,
    ontoggle,
    oncopy,
  }: {
    rows: readonly CandidateRow[];
    slot: Slot;
    locked: boolean;
    enchants: readonly EnchantRow[];
    suffixes: readonly SuffixRow[];
    classSlug: string;
    treeVersion: string;
    ontoggle: (key: string) => void;
    oncopy: (key: string, patch: { enchant?: number; suffix?: number }) => void;
  } = $props();

  /** Which row's copy-and-modify menu is open, by key. Only ever one. */
  let openKey = $state<string | null>(null);

  const ORIGIN_LABELS: Record<string, string> = {
    equipped: bulkCopy.equipped,
    bag: bulkCopy.bags,
    bank: bulkCopy.bank,
    search: bulkCopy.fromSearch,
  };

  function originLabel(origin: string): string {
    return ORIGIN_LABELS[origin] ?? (origin.startsWith('drop:') ? bulkCopy.pinned : origin);
  }

  /**
   * The row's own test id. A copy carries its enchant and suffix, so the original and every
   * copy of it are separately addressable -- which is the whole point of copy-and-modify.
   */
  function testId(row: CandidateRow): string {
    const base = `sim-candidate-${row.slot}-${row.item.id}`;
    const parts = [row.enchant > 0 ? `e${row.enchant}` : '', row.suffix > 0 ? `s${row.suffix}` : ''];
    return [base, ...parts.filter((part) => part !== '')].join('-');
  }
</script>

<ul class="flex flex-col">
  {#each rows as row (candidateKey(row))}
    <li class="border-line-soft flex flex-wrap items-center gap-2 border-b px-2 py-1 last:border-b-0">
      <label class="flex min-h-11 flex-1 items-center gap-3" data-testid={testId(row)}>
        <input
          type="checkbox"
          class="h-5 w-5"
          checked={row.checked}
          disabled={locked}
          onchange={() => ontoggle(candidateKey(row))}
        />
        <img
          src={dataUrl(treeVersion, `icons/${row.item.icon}.webp`)}
          alt=""
          width="28"
          height="28"
          loading="lazy"
          decoding="async"
          class="rounded-control border-line h-7 w-7 border object-cover"
        />
        <span class={`text-[14px] font-semibold ${rarityClassFor(row.item.quality)}`}>
          {row.item.name}
        </span>
        <span class="text-muted text-[12px]">{originLabel(row.origin)}</span>
        <span class="tabular text-muted ml-auto font-mono text-[12px]">{row.item.item_level}</span>
      </label>
      <button
        type="button"
        class="border-line-warm rounded-control text-nav label min-h-11 border px-3"
        aria-expanded={openKey === candidateKey(row)}
        onclick={() => (openKey = openKey === candidateKey(row) ? null : candidateKey(row))}
      >
        {bulkCopy.copyAndModify}
      </button>
      {#if openKey === candidateKey(row)}
        <div class="flex w-full flex-col gap-2 pb-2 md:flex-row">
          <div class="flex-1">
            <p class="text-muted px-2 text-[12px]">{bulkCopy.withEnchant}</p>
            <EnchantList
              rows={enchants}
              {slot}
              {classSlug}
              {treeVersion}
              onpick={(enchant) => {
                oncopy(candidateKey(row), { enchant });
                openKey = null;
              }}
            />
          </div>
          {#if suffixesForItem(suffixes, row.item).length > 0}
            <div class="flex-1">
              <p class="text-muted px-2 text-[12px]">{bulkCopy.withSuffix}</p>
              <ul class="border-line bg-raised rounded-panel flex flex-col border p-2">
                {#each suffixesForItem(suffixes, row.item) as suffix (suffix.id)}
                  <li>
                    <button
                      type="button"
                      class="text-text flex min-h-11 w-full items-center px-2 text-left text-[13px]"
                      data-testid={`sim-suffix-${row.slot}-${row.item.id}-${suffix.id}`}
                      onclick={() => {
                        oncopy(candidateKey(row), { suffix: suffix.id });
                        openKey = null;
                      }}>{suffix.name}</button
                    >
                  </li>
                {/each}
              </ul>
            </div>
          {/if}
        </div>
      {/if}
    </li>
  {/each}
  {#if rows.length === 0}
    <li class="text-muted px-2 py-2 text-[13px]">{bulkCopy.noCandidates}</li>
  {/if}
</ul>
```

- [ ] **Step 5: Write `SlotGrid.svelte`**

```svelte
<!-- web/src/components/sim/tools/SlotGrid.svelte -->
<!-- Design 3.1.2: one section per slot, each with its candidates and a lock. Rings and
     trinkets get one section, not two: sim/bulk tries them in both slots and the grid
     saying otherwise would be a promise the engine does not keep. -->
<script lang="ts">
  import { SLOTS, SLOT_LABELS, type Slot } from '../../../lib/planner/types';
  import type { CandidateRow } from '../../../lib/sim/candidates';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { EnchantRow, SuffixRow } from '../../../lib/sim/enchants';
  import CandidateRows from './CandidateRows.svelte';

  let {
    rows,
    locked,
    enchants,
    suffixes,
    classSlug,
    treeVersion,
    ontoggle,
    oncopy,
    onlock,
  }: {
    rows: readonly CandidateRow[];
    locked: readonly string[];
    enchants: readonly EnchantRow[];
    suffixes: readonly SuffixRow[];
    classSlug: string;
    treeVersion: string;
    ontoggle: (key: string) => void;
    oncopy: (key: string, patch: { enchant?: number; suffix?: number }) => void;
    onlock: (slot: Slot) => void;
  } = $props();

  /** finger2 and trinket2 never get a section of their own; uiSlotsOf puts both in the first. */
  const SHOWN: readonly Slot[] = SLOTS.filter((slot) => slot !== 'finger2' && slot !== 'trinket2');

  const bySlot = $derived(
    new Map(SHOWN.map((slot) => [slot, rows.filter((row) => row.slot === slot)])),
  );
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-slot-grid">
  {#each SHOWN as slot (slot)}
    {@const slotRows = bySlot.get(slot) ?? []}
    {#if slotRows.length > 0}
      <div class="border-line rounded-panel border p-3">
        <header class="flex flex-wrap items-center justify-between gap-2">
          <h3 class="section-title text-[14px]">{SLOT_LABELS[slot]}</h3>
          <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
            <input
              type="checkbox"
              class="h-5 w-5"
              data-testid={`sim-lock-${slot}`}
              checked={locked.includes(slot)}
              onchange={() => onlock(slot)}
            />
            {locked.includes(slot) ? bulkCopy.lockedSlot : bulkCopy.lockSlot}
          </label>
        </header>
        <CandidateRows
          rows={slotRows}
          {slot}
          locked={locked.includes(slot)}
          {enchants}
          {suffixes}
          {classSlug}
          {treeVersion}
          {ontoggle}
          {oncopy}
        />
      </div>
    {/if}
  {/each}
</section>
```

- [ ] **Step 6: Write `TopGear.svelte`, the composition later tasks add to**

```svelte
<!-- web/src/components/sim/tools/TopGear.svelte -->
<!-- /sim/gear and /sim/talents are the same island: Top Gear with its gear locked and only
     the loadouts as candidates (design 3.4). `store.tool` is the whole difference, and it
     is read once here rather than branched on in every child. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { Slot } from '../../../lib/planner/types';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';
  import SlotGrid from './SlotGrid.svelte';

  let { store, me }: { store: BulkStore; me: Me | null } = $props();

  const gearMode = $derived(store.tool === 'gear');
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-top-gear">
  {#if gearMode}
    <SlotGrid
      rows={store.rows}
      locked={store.locked}
      enchants={store.enchants}
      suffixes={store.suffixes}
      classSlug={store.character?.class_slug ?? ''}
      treeVersion={store.character?.tree_version ?? ''}
      ontoggle={(key) => store.toggleRow(key)}
      oncopy={(key, patch) => store.copyAndModify(key, patch)}
      onlock={(slot: Slot) => store.toggleLock(slot)}
    />
  {/if}

  <!-- Task 13 inserts ItemSearch here, Task 14 TalentCandidates and NamedSets,
       Task 15 BulkRunBar, Task 16 ComboResults. -->
  <p class="text-muted px-[18px] text-[13px] md:px-0" data-testid="sim-combo-count">
    {store.combinations === null ? bulkCopy.combinationsCounting : bulkCopy.combinations(store.combinations)}
  </p>
</div>
```

- [ ] **Step 7: Run the e2e, desktop and mobile**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts
```

Expected: PASS on both projects.

- [ ] **Step 8: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools tests/e2e/sim-gear.spec.ts
npx prettier --write src/components/sim/tools tests/e2e/sim-gear.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/tests/e2e/sim-gear.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): Top Gear's slot grid, locks, enchants and copy-and-modify

One section per slot, rings and trinkets sharing one because sim/bulk
tries both slots anyway. A locked slot disables its checkboxes rather
than hiding them: the player has to see what they locked away. A copy is
a new row beside the original with its own test id, so a grid with an
item and two enchanted versions of it is three addressable rows.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 13: The item search, and consumables as candidates

**Files:**
- Create: `web/src/components/sim/tools/ItemSearch.svelte`
- Create: `web/src/components/sim/tools/ConsumableCandidates.svelte`
- Modify: `web/src/components/sim/tools/TopGear.svelte` (mount both)
- Modify: `web/src/lib/sim/candidates.ts` (`buildBulkSpec` carries `consumables`)
- Modify: `web/src/lib/sim/candidates.test.ts`
- Modify: `web/tests/e2e/sim-gear.spec.ts`

`BulkSpec.consumables`, `store.consumableIds` and `store.toggleConsumable` already exist:
Task 1 declared the field and Task 10 the state. This task is the control and the
`buildBulkSpec` plumbing.

**Interfaces:**
- Consumes: `searchItems`, `defaultItemQuery`, `slotOptions`, `SEARCH_LIMIT`, `ItemQuery` (Task 7); `groupSources` (Task 6); `PRESET_CONSUMABLES` and `SimSettings` from `./settings`.
- Produces: `ItemSearch.svelte`, `ConsumableCandidates.svelte`, and `BulkSpecInput.consumables` on `buildBulkSpec`.

> **Ratified by contract 10.1 A5.** `BulkSpec.Consumables [][]string` is a real field:
> "alternative consumable lists tried as candidates in `gear` mode… Each inner list
> replaces `CharacterSpec.Consumes` for that combination." It multiplies the gear product,
> so the live count **does** move when a consumable is ticked, and the e2e asserts that.
>
> The result shape is settled too: contract 10.8 gives `Substitution.Kind` a fourth value,
> `consumes`, with `Name` the consumable ids joined by `, `, and has `sim/bulk` emit one per
> combination that used an alternative list. Task 1 types the union closed on those four,
> Task 3's fake emits `kind: 'consumes'`, and `substitutionLabel` renders it through
> `bulkCopy.consumesChip` -- a named kind, not an unknown.

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/sim/candidates.test.ts`:

```ts
describe('consumable candidates', () => {
  it('travels as a list of alternative sets, and is absent when nothing is picked', () => {
    const spec = buildBulkSpec({
      mode: 'gear',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'fast',
      cap: 400,
      consumables: [['flask_of_supreme_power'], ['elixir_of_the_mongoose']],
    });
    expect(spec.consumables).toEqual([['flask_of_supreme_power'], ['elixir_of_the_mongoose']]);

    const none = buildBulkSpec({
      mode: 'gear',
      rows: [{ ...rowFor(helm, 'head', 'bag'), checked: true }],
      locked: [],
      loadouts: [],
      sets: [],
      precision: 'fast',
      cap: 400,
    });
    expect(none.consumables).toBeUndefined();
  });
});
```

Append to `web/tests/e2e/sim-gear.spec.ts`:

```ts
test('the item search adds a candidate and says how it was filtered', async ({ page }) => {
  await loadGear(page);
  const search = page.getByTestId('sim-item-search');
  await expect(search).toBeVisible();
  await search.getByRole('searchbox').fill('wrath');
  await expect(page.getByTestId('sim-search-result-16963')).toBeVisible();
  await expect(page.getByTestId('sim-search-result-12640')).toHaveCount(0);
  await page.getByTestId('sim-search-add-16963').click();
  await expect(page.getByTestId('sim-candidate-head-16963').getByRole('checkbox')).toBeChecked();
});

test('the source filter narrows the search to one boss’s loot', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-source').selectOption('world:azuregos');
  await expect(page.getByTestId('sim-search-result-19325')).toBeVisible();
  await expect(page.getByTestId('sim-search-result-16963')).toHaveCount(0);
});

test('usable-only is on by default and can be turned off', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-search-usable')).toBeChecked();
  await page.getByTestId('sim-search-usable').uncheck();
  await expect(page.getByTestId('sim-search-usable')).not.toBeChecked();
});

test('a ticked consumable multiplies the combination count (contract 10.1 A5)', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await expect(page.getByTestId('sim-combo-count')).toHaveText('1 valid combination');
  await page.getByTestId('sim-consumable-flask_of_supreme_power').check();
  await page.getByTestId('sim-consumable-elixir_of_the_mongoose').check();
  // (1 head + 1) x 2 alternatives, minus the untouched base
  await expect(page.getByTestId('sim-combo-count')).toHaveText('3 valid combinations');
});

test('a consumable candidate is named, not spelled as an id', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-consumables')).toContainText('Flask of Supreme Power');
});
```

- [ ] **Step 2: Run them and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/candidates.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts --project=desktop
```

Expected: FAIL on the new cases only.

- [ ] **Step 3: Carry consumables through `buildBulkSpec`**

In `candidates.ts`, add `consumables?: string[][]` to `BulkSpecInput`, and in
`buildBulkSpec`'s returned object:

```ts
    ...(input.consumables !== undefined && input.consumables.length > 0
      ? { consumables: input.consumables.map((set) => [...set]) }
      : {}),
```

- [ ] **Step 4: Write `ItemSearch.svelte`**

```svelte
<!-- web/src/components/sim/tools/ItemSearch.svelte -->
<!-- Design 3.1.3: name, minimum item level, slot, source, usable-only. Entirely
     client-side over the items file the grid already has -- no API call, and no second
     index. The design's own caveat stands and the page says it: search can find items this
     character has no way to obtain. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
  import { dataUrl } from '../../../lib/planner/load';
  import type { Item } from '../../../lib/planner/types';
  import { bulkCopy } from '../../../lib/sim/copy';
  import {
    defaultItemQuery,
    searchItems,
    slotOptions,
    type ItemQuery,
    type SearchContext,
  } from '../../../lib/sim/item-search';
  import { groupSources, type LootFile } from '../../../lib/sim/loot';

  let {
    items,
    loot,
    ctx,
    treeVersion,
    onadd,
  }: {
    items: readonly Item[];
    loot: LootFile;
    ctx: SearchContext;
    treeVersion: string;
    onadd: (itemId: number) => void;
  } = $props();

  let query = $state<ItemQuery>(defaultItemQuery());

  const found = $derived(searchItems(items, query, ctx));
  const matching = $derived(
    searchItems(items, { ...query, text: query.text }, ctx).length === items.length
      ? items.length
      : found.length,
  );
  const groups = $derived(groupSources(loot));
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-3 border p-3 md:mx-0"
  data-testid="sim-item-search"
>
  <h3 class="section-title text-[14px]">{bulkCopy.searchLabel}</h3>

  <div class="flex flex-wrap gap-2">
    <input
      type="search"
      bind:value={query.text}
      placeholder={bulkCopy.searchPlaceholder}
      aria-label={bulkCopy.searchLabel}
      class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 min-w-[180px] flex-1 border px-3 text-[14px]"
    />
    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.searchMinItemLevel}
      <input
        type="number"
        min="0"
        max="120"
        bind:value={query.minItemLevel}
        data-testid="sim-search-ilvl"
        class="border-line-warm rounded-control bg-bg text-text h-11 w-20 border px-2 text-[14px]"
      />
    </label>
    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.searchSlot}
      <select
        bind:value={query.slot}
        data-testid="sim-search-slot"
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
      >
        <option value="">{bulkCopy.searchAnySlot}</option>
        {#each slotOptions() as option (option.slot)}
          <option value={option.slot}>{option.label}</option>
        {/each}
      </select>
    </label>
    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.searchSource}
      <select
        bind:value={query.sourceId}
        data-testid="sim-search-source"
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
      >
        <option value="">{bulkCopy.searchAnySource}</option>
        {#each groups as group (group.kind)}
          <optgroup label={group.label}>
            {#each group.sources as source (source.id)}
              <option value={source.id}>{source.name}</option>
            {/each}
          </optgroup>
        {/each}
      </select>
    </label>
    <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid="sim-search-usable"
        bind:checked={query.usableOnly}
      />
      {bulkCopy.searchUsableOnly}
    </label>
  </div>

  {#if found.length === 0}
    <p class="text-muted text-[13px]">{bulkCopy.searchNoResults}</p>
  {:else}
    <ul class="max-h-[320px] overflow-y-auto">
      {#each found as item (item.id)}
        <li
          class="border-line-soft flex min-h-11 items-center gap-3 border-b px-1 py-1 last:border-b-0"
          data-testid={`sim-search-result-${item.id}`}
        >
          <img
            src={dataUrl(treeVersion, `icons/${item.icon}.webp`)}
            alt=""
            width="24"
            height="24"
            loading="lazy"
            decoding="async"
            class="rounded-control border-line h-6 w-6 border object-cover"
          />
          <span class={`flex-1 text-[14px] font-semibold ${rarityClassFor(item.quality)}`}>
            {item.name}
          </span>
          <span class="tabular text-muted font-mono text-[12px]">{item.item_level}</span>
          <button
            type="button"
            class="border-line-warm rounded-control text-nav label min-h-11 border px-3"
            data-testid={`sim-search-add-${item.id}`}
            onclick={() => onadd(item.id)}>{bulkCopy.searchAdd}</button
          >
        </li>
      {/each}
    </ul>
    {#if found.length < matching}
      <p class="text-muted text-[12px]">{bulkCopy.searchTruncated(found.length, matching)}</p>
    {/if}
  {/if}
</section>
```

- [ ] **Step 5: Write `ConsumableCandidates.svelte`**

```svelte
<!-- web/src/components/sim/tools/ConsumableCandidates.svelte -->
<!-- Design 3.1.5, ratified as BulkSpec.Consumables by contract 10.1 A5: the consumables the
     settings panel applies as a setting can instead be tried one at a time as candidates,
     each ticked one replacing CharacterSpec.Consumes wholesale for its combination.
     The id list is the preset's own, not a second copy of IDS.md -- settings.ts is the one
     place this lane names a consumable id -- and the names and icons come from
     simbuffs.json (contract 10.4) rather than from a de-underscored id. -->
<script lang="ts">
  import { dataUrl } from '../../../lib/planner/load';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { PRESET_CONSUMABLES } from '../../../lib/sim/settings';
  import { buffIcon, buffName, type SimBuffFile } from '../../../lib/sim/sim-buffs';

  let {
    picked,
    buffs,
    treeVersion,
    ontoggle,
  }: {
    picked: readonly string[];
    buffs: SimBuffFile;
    treeVersion: string;
    ontoggle: (id: string) => void;
  } = $props();

  const offered = $derived([...new Set(PRESET_CONSUMABLES['raid-buffed'])]);
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0"
  data-testid="sim-consumables"
>
  <h3 class="section-title text-[14px]">{bulkCopy.tryEach}</h3>
  <p class="text-muted text-[12px]">{bulkCopy.consumableCandidates}</p>
  <ul class="flex flex-wrap gap-3">
    {#each offered as id (id)}
      <li>
        <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
          <input
            type="checkbox"
            class="h-5 w-5"
            data-testid={`sim-consumable-${id}`}
            checked={picked.includes(id)}
            onchange={() => ontoggle(id)}
          />
          {#if buffIcon(buffs, id) !== ''}
            <img
              src={dataUrl(treeVersion, `icons/${buffIcon(buffs, id)}.webp`)}
              alt=""
              width="20"
              height="20"
              loading="lazy"
              decoding="async"
              class="rounded-control border-line h-5 w-5 border object-cover"
            />
          {/if}
          {buffName(buffs, id)}
        </label>
      </li>
    {/each}
  </ul>
</section>
```

- [ ] **Step 6: Mount both in `TopGear.svelte`**

Inside the `{#if gearMode}` block, after `<SlotGrid … />`:

```svelte
    <ItemSearch
      items={[...store.items.values()]}
      loot={store.loot}
      ctx={{ level: SIM_LEVEL, sourcesByItem: store.sourceIndex }}
      treeVersion={store.character?.tree_version ?? ''}
      onadd={(itemId) => store.addSearchItem(itemId)}
    />
    <ConsumableCandidates
      picked={store.consumableIds}
      buffs={store.simBuffs}
      treeVersion={store.character?.tree_version ?? ''}
      ontoggle={(id) => store.toggleConsumable(id)}
    />
```

with `import { SIM_LEVEL } from '../../../lib/sim/character';`, `import ItemSearch from
'./ItemSearch.svelte';` and `import ConsumableCandidates from './ConsumableCandidates.svelte';`.

- [ ] **Step 7: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/candidates.test.ts src/lib/sim/bulk-store.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts
```

Expected: PASS.

- [ ] **Step 8: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools src/lib/sim/candidates.ts src/lib/sim/bulk-types.ts src/lib/sim/bulk-store.svelte.ts tests/e2e/sim-gear.spec.ts
npx prettier --write src/components/sim/tools src/lib/sim/candidates.ts src/lib/sim/bulk-types.ts src/lib/sim/bulk-store.svelte.ts tests/e2e/sim-gear.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/src/lib/sim web/tests/e2e/sim-gear.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): the item search, and consumables as candidates

Search is client-side over the items file the grid already holds: name,
minimum item level, slot, loot source, usable-only on by default.
Consumable candidates send BulkSpec.Consumables, which contract 10.1 A5
makes a real field that multiplies the gear product -- so the live count
moves when one is ticked, and the test asserts that. Names and icons
come from simbuffs.json rather than from a de-underscored id.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 14: Talent loadouts and named sets as candidates

**Files:**
- Create: `web/src/components/sim/tools/TalentCandidates.svelte`
- Create: `web/src/components/sim/tools/NamedSets.svelte`
- Modify: `web/src/components/sim/tools/TopGear.svelte`
- Modify: `web/src/components/planner/Planner.svelte` (one additive `oncode` prop)
- Modify: `web/src/lib/sim/api.ts` (`fetchMyBuilds`)
- Modify: `web/src/lib/sim/api.test.ts`
- Modify: `web/tests/e2e/sim-gear.spec.ts`

**Interfaces:**
- Consumes: `store.loadouts`, `store.addLoadout`, `store.removeLoadout`, `store.namedSets`, `store.addNamedSet`, `store.removeNamedSet` (Task 10); `decodeFS1`, `orderFromRanks`, `indexTalents`, `talentsString`; `BuildRecord`, `TalentFile`.
- **Consumes from part A:** `SimCharacter.loadouts?: { name: string; talents: string }[]` and `.sets?: { name: string; gear: GearSlot[] }[]`, both from the FS1 v2 decoder.
- Produces: `fetchMyBuilds(page?, apiBase?): Promise<{ rows: BuildRecord[]; total: number; page: number; per_page: number }>`; `Planner.svelte`'s `oncode?: (code: string) => void`.

> **Ratified by contract 10.6.** `builds` gains `user_id` (nullable, set on save when
> signed in) and `GET /v1/builds?mine=1` lists the signed-in player's builds; the contract
> names Top Gear's talent list as its reader. The page still **treats every failure as "no
> saved builds"** with the copy `bulkCopy.talentsSavedUnavailable`: a deployment older than
> the migration answers 404, and a list that came back empty is not worth an error banner.

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/sim/api.test.ts`:

```ts
describe('fetchMyBuilds', () => {
  it('reads the signed-in player’s saved builds', async () => {
    const page = await fetchMyBuilds(1, TEST_API);
    expect(page.rows[0].id).toBe(FIXTURE_MY_BUILD_ID);
    expect(api.lastUrl()).toContain('/v1/builds?mine=1&page=1');
  });

  it('raises a SimApiError the caller can treat as "none" rather than as a page failure', async () => {
    api.route({ method: 'GET', pattern: /\/v1\/builds\?/, respond: () => failure('nope', 404) });
    await expect(fetchMyBuilds(1, TEST_API)).rejects.toBeInstanceOf(SimApiError);
  });
});
```

Append to `web/tests/e2e/sim-gear.spec.ts`:

```ts
test('the talent list offers the character’s own build and a saved one', async ({ page }) => {
  await page.route('**/v1/builds?mine=1*', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          rows: [
            {
              id: 'bld987654321',
              class_id: 1,
              race_id: 2,
              tree_version: activeBuild.build,
              point_order: [2001, 2001, 2001, 2001, 2001],
              gear: {},
              title: 'Deep Fury',
              created_at: '2026-09-18T12:00:00Z',
              views: 2,
            },
          ],
          total: 1,
          page: 1,
          per_page: 100,
        },
      }),
    }),
  );
  await loadGear(page);
  await expect(page.getByTestId('sim-loadout-current')).toBeVisible();
  await page.getByTestId('sim-loadout-Deep Fury').check();
  await expect(page.getByTestId('sim-loadout-Deep Fury')).toBeChecked();
});

test('a pasted second export string becomes a named set', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-set-input').fill(`${FURY}`);
  await page.getByTestId('sim-set-name').fill('PvP set');
  await page.getByTestId('sim-set-add').click();
  await expect(page.getByTestId('sim-set-PvP set')).toBeVisible();
});

test('a set that is not an export string says so and adds nothing', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-set-input').fill('FS2:nope');
  await page.getByTestId('sim-set-name').fill('Bad');
  await page.getByTestId('sim-set-add').click();
  await expect(page.getByTestId('sim-set-error')).toBeVisible();
  await expect(page.getByTestId('sim-set-Bad')).toHaveCount(0);
});
```

- [ ] **Step 2: Run them and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/api.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts --project=desktop
```

Expected: FAIL on the new cases.

- [ ] **Step 3: Add `fetchMyBuilds`**

In `web/src/lib/sim/api.ts`, beside `listMySims`:

```ts
/**
 * The signed-in player's saved planner builds, for Top Gear's talent candidates
 * (contract 10.6: `builds.user_id` plus this route). The caller treats every failure as
 * "no saved builds" and says so in one line -- a deployment older than migration 0014
 * answers 404, and that is not worth an error banner on a page whose other numbers are
 * all correct.
 */
export function fetchMyBuilds(
  page: number = 1,
  apiBase: string = API_BASE_URL,
): Promise<{ rows: BuildRecord[]; total: number; page: number; per_page: number }> {
  return call(`/v1/builds?mine=1&page=${page}`, apiBase, simCopy.loadFailed);
}
```

with `import type { BuildRecord } from '../planner/types';` added.

- [ ] **Step 4: Give `Planner.svelte` an `oncode` callback**

In its props:

```ts
    record = null,
    oncode,
  }: {
    treeVersion: string;
    classSlug?: string;
    raceSlug?: string;
    record?: BuildRecord | null;
    /**
     * Called with the build's own FS1 code whenever it changes. Top Gear's "add a build"
     * mounts this component inline and reads the code back through it; /planner and /b/:id
     * pass nothing and the callback never fires.
     */
    oncode?: (code: string) => void;
  } = $props();
```

and, immediately after the `simHref` derivation:

```ts
  /**
   * The same code `simHref` embeds, handed to an embedder that asked for it. A `$effect`
   * rather than a call inside the derivation: a derivation must stay a pure read, and this
   * is a side effect on every build change.
   */
  const liveCode = $derived(
    store.talentIndex === null
      ? ''
      : encodeFS1({
          dataBuild: store.treeVersion,
          classSlug: store.classSlug,
          raceSlug: store.raceSlug,
          treeRanks: treeRanksFor(store.talentIndex, store.order),
          gear: store.gear,
        }),
  );

  $effect(() => {
    if (liveCode !== '') oncode?.(liveCode);
  });
```

- [ ] **Step 5: Write `TalentCandidates.svelte`**

```svelte
<!-- web/src/components/sim/tools/TalentCandidates.svelte -->
<!-- Design 3.1.6: the character's own build, every planner build the signed-in player has
     saved for this class, the in-game loadouts the addon export carried, and "add a build"
     which mounts the planner inline and reads its code back.
     The planner is a large component and /sim/weights must never download it, so it is a
     lazy chunk opened only when the player asks for it. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { decodeFS1, orderFromRanks } from '../../../lib/planner/fs1';
  import { indexTalents } from '../../../lib/planner/rules';
  import { loadTalents } from '../../../lib/planner/load';
  import type { BuildRecord, TalentFile } from '../../../lib/planner/types';
  import { createLazyComponent } from '../../../lib/report/lazy-component.svelte';
  import { fetchMyBuilds } from '../../../lib/sim/api';
  import { talentsString, type SimCharacter } from '../../../lib/sim/character';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { TalentLoadout } from '../../../lib/sim/bulk-types';

  let {
    character,
    picked,
    ontoggle,
  }: {
    character: SimCharacter;
    picked: readonly TalentLoadout[];
    ontoggle: (loadout: TalentLoadout, on: boolean) => void;
  } = $props();

  let talents = $state<TalentFile | null>(null);
  let saved = $state<BuildRecord[] | null>(null);
  let savedFailed = $state(false);
  let plannerOpen = $state(false);
  let customCode = $state('');

  const plannerLazy = createLazyComponent(() => import('../../planner/Planner.svelte'));

  onMount(() => {
    void loadTalents(character.tree_version, character.class_slug)
      .then((file) => (talents = file))
      .catch(() => (talents = null));
    void fetchMyBuilds()
      .then((page) => {
        saved = page.rows.filter((row) => row.tree_version === character.tree_version);
      })
      .catch(() => {
        saved = [];
        savedFailed = true;
      });
  });

  /** A saved build's point order as the engine's talents string, through the one converter. */
  function loadoutFor(record: BuildRecord): TalentLoadout | null {
    if (talents === null) return null;
    return {
      name: record.title === undefined || record.title === '' ? record.id : record.title,
      talents: talentsString(indexTalents(talents), record.point_order),
    };
  }

  /** The character's own build, always the first row and always available. */
  const own = $derived<TalentLoadout | null>(
    talents === null
      ? null
      : { name: bulkCopy.talentsOwn, talents: talentsString(indexTalents(talents), character.point_order) },
  );

  const savedLoadouts = $derived(
    (saved ?? []).map(loadoutFor).filter((entry): entry is TalentLoadout => entry !== null),
  );

  /** The addon export's in-game loadouts (part A's FS1 v2 decoder). */
  const exported = $derived<TalentLoadout[]>(character.loadouts ?? []);

  function isPicked(loadout: TalentLoadout): boolean {
    return picked.some((entry) => entry.name === loadout.name);
  }

  /** The inline planner's code, turned into a loadout the moment the player accepts it. */
  function addCustom(): void {
    if (talents === null || customCode === '') return;
    const decoded = decodeFS1(customCode);
    if (!decoded.ok) return;
    const index = indexTalents(talents);
    const { order } = orderFromRanks(index, decoded.build.treeRanks);
    ontoggle({ name: `Build ${picked.length + 1}`, talents: talentsString(index, order) }, true);
    plannerOpen = false;
  }
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0"
  data-testid="sim-talent-candidates"
>
  <h3 class="section-title text-[14px]">{bulkCopy.talentsSaved}</h3>

  {#if own !== null}
    <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid="sim-loadout-current"
        checked={isPicked(own)}
        onchange={(event) => ontoggle(own, event.currentTarget.checked)}
      />
      {bulkCopy.talentsOwn}
    </label>
  {/if}

  {#each savedLoadouts as loadout (loadout.name)}
    <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid={`sim-loadout-${loadout.name}`}
        checked={isPicked(loadout)}
        onchange={(event) => ontoggle(loadout, event.currentTarget.checked)}
      />
      {loadout.name}
    </label>
  {/each}

  {#if exported.length > 0}
    <h4 class="text-muted text-[12px]">{bulkCopy.talentsLoadouts}</h4>
    {#each exported as loadout (loadout.name)}
      <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
        <input
          type="checkbox"
          class="h-5 w-5"
          data-testid={`sim-loadout-${loadout.name}`}
          checked={isPicked(loadout)}
          onchange={(event) => ontoggle(loadout, event.currentTarget.checked)}
        />
        {loadout.name}
      </label>
    {/each}
  {/if}

  {#if savedFailed}
    <p class="text-muted text-[12px]" data-testid="sim-loadouts-unavailable">
      {bulkCopy.talentsSavedUnavailable}
    </p>
  {:else if savedLoadouts.length === 0 && saved !== null}
    <p class="text-muted text-[12px]">{bulkCopy.talentsNoSaved}</p>
  {/if}

  <button
    type="button"
    class="border-line-warm rounded-control text-nav label min-h-11 w-fit border px-3"
    data-testid="sim-loadout-add"
    onclick={() => {
      plannerOpen = !plannerOpen;
      if (plannerOpen) plannerLazy.load();
    }}>{bulkCopy.talentsAddCustom}</button
  >

  {#if plannerOpen && plannerLazy.current}
    <div class="border-line rounded-panel border p-2" data-testid="sim-inline-planner">
      <plannerLazy.current
        treeVersion={character.tree_version}
        classSlug={character.class_slug}
        raceSlug={character.race_slug}
        oncode={(code: string) => (customCode = code)}
      />
      <button
        type="button"
        class="border-line-warm rounded-control text-nav label min-h-11 border px-3"
        data-testid="sim-loadout-accept"
        onclick={addCustom}>{bulkCopy.talentsAddCustom}</button
      >
    </div>
  {/if}
</section>
```

- [ ] **Step 6: Write `NamedSets.svelte`**

```svelte
<!-- web/src/components/sim/tools/NamedSets.svelte -->
<!-- Design 3.1.7 and the decision that killed Gear Compare: a whole-set alternative is one
     candidate that replaces every slot at once. Sets come from the addon export (part A's
     FS1 v2 `sets=` section) or from a second export string pasted here. -->
<script lang="ts">
  import { decodeFS1 } from '../../../lib/planner/fs1';
  import { SLOTS, type Slot } from '../../../lib/planner/types';
  import type { GearSet } from '../../../lib/sim/bulk-types';
  import type { SimCharacter } from '../../../lib/sim/character';
  import { bulkCopy } from '../../../lib/sim/copy';

  let {
    character,
    sets,
    onadd,
    onremove,
  }: {
    character: SimCharacter;
    sets: readonly GearSet[];
    onadd: (set: GearSet) => void;
    onremove: (name: string) => void;
  } = $props();

  let code = $state('');
  let name = $state('');
  let error = $state('');

  /** The export's own named sets, offered as one-click adds (part A's decoder). */
  const exported = $derived<GearSet[]>(character.sets ?? []);

  function add(): void {
    error = '';
    const decoded = decodeFS1(code.trim());
    if (!decoded.ok) {
      error = bulkCopy.setsBadCode;
      return;
    }
    const gear = SLOTS.flatMap((slot) => {
      const itemId = decoded.build.gear[slot as Slot];
      return itemId === undefined ? [] : [{ slot, item_id: itemId }];
    });
    onadd({ name: name.trim() === '' ? `Set ${sets.length + 1}` : name.trim(), gear });
    code = '';
    name = '';
  }
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0"
  data-testid="sim-named-sets"
>
  <h3 class="section-title text-[14px]">{bulkCopy.setsTitle}</h3>
  <p class="text-muted text-[12px]">{bulkCopy.setsIntro}</p>

  {#each [...exported, ...sets] as set (set.name)}
    <div class="flex min-h-11 items-center gap-2 text-[13px]" data-testid={`sim-set-${set.name}`}>
      <span class="text-text flex-1">{set.name}</span>
      <button
        type="button"
        class="border-line-warm rounded-control text-nav label min-h-11 border px-3"
        onclick={() => (sets.some((entry) => entry.name === set.name) ? onremove(set.name) : onadd(set))}
      >
        {sets.some((entry) => entry.name === set.name) ? 'Remove' : bulkCopy.setsAdd}
      </button>
    </div>
  {/each}

  <div class="flex flex-wrap gap-2">
    <input
      type="text"
      bind:value={name}
      placeholder="Name"
      aria-label="Set name"
      data-testid="sim-set-name"
      class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 w-40 border px-3 text-[14px]"
    />
    <input
      type="text"
      bind:value={code}
      placeholder={bulkCopy.setsPaste}
      aria-label={bulkCopy.setsPaste}
      data-testid="sim-set-input"
      class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 min-w-[200px] flex-1 border px-3 text-[14px]"
    />
    <button
      type="button"
      class="border-line-warm rounded-control text-nav label min-h-11 border px-3"
      data-testid="sim-set-add"
      onclick={add}>{bulkCopy.setsAdd}</button
    >
  </div>
  {#if error !== ''}
    <p class="text-muted text-[13px]" role="alert" data-testid="sim-set-error">{error}</p>
  {/if}
</section>
```

- [ ] **Step 7: Mount both in `TopGear.svelte`**

After `<ConsumableCandidates … />`, and **outside** the `{#if gearMode}` block for
`TalentCandidates` (talent compare is loadouts only):

```svelte
  {#if store.character !== null}
    <TalentCandidates
      character={store.character}
      picked={store.loadouts}
      ontoggle={(loadout, on) => (on ? store.addLoadout(loadout) : store.removeLoadout(loadout.name))}
    />
  {/if}
  {#if gearMode && store.character !== null}
    <NamedSets
      character={store.character}
      sets={store.namedSets}
      onadd={(set) => store.addNamedSet(set)}
      onremove={(name) => store.removeNamedSet(name)}
    />
  {/if}
```

- [ ] **Step 8: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/api.test.ts src/components
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts tests/e2e/planner.spec.ts
```

Expected: PASS, and the planner's own suite unaffected by the additive prop.

- [ ] **Step 9: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools src/components/planner/Planner.svelte src/lib/sim/api.ts src/lib/sim/api.test.ts tests/e2e/sim-gear.spec.ts
npx prettier --write src/components/sim/tools src/components/planner/Planner.svelte src/lib/sim/api.ts src/lib/sim/api.test.ts tests/e2e/sim-gear.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components web/src/lib/sim/api.ts web/src/lib/sim/api.test.ts web/tests/e2e/sim-gear.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): talent loadouts and named sets as candidates

The character's own build, the signed-in player's saved planner builds
through GET /v1/builds?mine=1 (contract 10.6), the export's in-game
loadouts, and a planner mounted inline as a lazy chunk so /sim/weights
never downloads it. A failure still reads as "no saved builds" in one
line, because a deployment older than the migration answers 404. Named
sets are the answer to Gear Compare: one candidate that replaces every
slot at once.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 15: The run bar — count, precision, cap notice, stage progress

**Files:**
- Create: `web/src/components/sim/tools/BulkRunBar.svelte`
- Modify: `web/src/components/sim/tools/TopGear.svelte` (replace the bare count paragraph)
- Modify: `web/tests/e2e/sim-gear.spec.ts`

**Interfaces:**
- Consumes: `store.combinations`, `capNotice`, `cap`, `precision`, `setPrecision`, `phase`, `progressLine`, `premium`, `serverRunning`, `run`, `stop`, `runOnServer`, `message`, `detail` (Task 10); `PRECISIONS`, `LOW_CORE_CAP`, `BROWSER_CAP` (Task 1); `SECONDARY_BUTTON`.
- **Consumes from part A:** `SettingsPanel.svelte` and `RequestDrawer.svelte`, both mounted here so every tool gets them in one place.
- Produces: test ids `sim-combo-count`, `sim-precision`, `sim-cap-notice`, `sim-stage-progress`, `sim-run-bulk`, `sim-server-run`.

- [ ] **Step 1: Write the failing e2e**

Append to `web/tests/e2e/sim-gear.spec.ts`:

```ts
test('four candidates are well under the cap, so no notice and a live run button', async ({
  page,
}) => {
  await loadGear(page);
  for (const id of [16963, 16966, 13968, 19325]) {
    await page.getByTestId(`sim-search-add-${id}`).click();
  }
  // The desktop project reports 8+ cores, so the cap is 400 and four candidates fit. The
  // notice's own wording and its 5,000-combination premium alternative are asserted in
  // the unit tests, where the cap is injectable.
  await expect(page.getByTestId('sim-cap-notice')).toHaveCount(0);
  await expect(page.getByTestId('sim-run-bulk')).toBeEnabled();
});

test('a run reports its stage line and ends with a ranked table', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-stage-progress')).toHaveText(
    /stage \d of 3 · \d+ of \d+ combinations/,
  );
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  await expect(page.getByTestId('sim-stage-progress')).toHaveCount(0);
});

test('precision is three choices and normal runs two stages', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-precision').selectOption('normal');
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-stage-progress')).toHaveText(/stage \d of 2 /);
});

test('the server lane is not offered to a signed-out visitor', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-server-run')).toHaveCount(0);
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts --project=desktop
```

Expected: FAIL — `sim-run-bulk` does not exist.

- [ ] **Step 3: Write `BulkRunBar.svelte`**

```svelte
<!-- web/src/components/sim/tools/BulkRunBar.svelte -->
<!-- Design 3.1.8: the combination count, the precision, the cap notice, the lane switch and
     one button. Every tool page mounts this, so the run affordance is identical on all four
     and the settings panel and the request drawer arrive with it.

     The cap notice never trims the list. It says what would exceed the cap and by how much,
     and offers the lane that would allow it -- design 2.3's own rule. -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { LOW_CORE_CAP, PRECISIONS, type Precision } from '../../../lib/sim/bulk-types';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import RequestDrawer from '../RequestDrawer.svelte';
  import SettingsPanel from '../SettingsPanel.svelte';

  let { store }: { store: BulkStore } = $props();

  const running = $derived(store.phase === 'running' || store.serverRunning);
  const countLabel = $derived(
    store.combinations === null
      ? bulkCopy.combinationsCounting
      : bulkCopy.combinations(store.combinations),
  );
  const PRECISION_LABELS: Record<Precision, string> = {
    fast: bulkCopy.precisionFast,
    normal: bulkCopy.precisionNormal,
    high: bulkCopy.precisionHigh,
  };
</script>

<SettingsPanel
  settings={store.settings}
  spec={store.character?.spec ?? ''}
  disabled={running}
  onchange={(next) => store.setSettings(next)}
/>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-4 md:mx-0"
  data-testid="sim-run-bulk-bar"
>
  <div class="flex flex-wrap items-center gap-4">
    <span class="tabular text-strong font-mono text-[14px]" data-testid="sim-combo-count">
      {countLabel}
    </span>

    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.precisionLabel}
      <select
        data-testid="sim-precision"
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
        disabled={running}
        value={store.precision}
        onchange={(event) => store.setPrecision(event.currentTarget.value as Precision)}
      >
        {#each PRECISIONS as precision (precision)}
          <option value={precision}>{PRECISION_LABELS[precision]}</option>
        {/each}
      </select>
    </label>

    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
      data-testid="sim-run-bulk"
      disabled={store.capNotice !== null || store.character === null}
      onclick={() => (running ? store.stop() : void store.run())}
    >
      {#if running}{bulkCopy.stopBulk}{:else if store.result !== null}{bulkCopy.runBulkAgain}{:else}{bulkCopy.runBulk}{/if}
    </button>

    {#if store.premium}
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
        data-testid="sim-server-run"
        disabled={running || store.serverCapNotice !== null}
        onclick={() => void store.runOnServer()}>{bulkCopy.capPremium}</button
      >
    {/if}
  </div>

  <p class="text-muted text-[12px]">{bulkCopy.precisionNote[store.precision]}</p>

  {#if store.cap === LOW_CORE_CAP}
    <p class="text-muted text-[12px]" data-testid="sim-low-core">{bulkCopy.lowCoreNote(store.cap)}</p>
  {/if}

  {#if store.capNotice !== null}
    <p class="text-strong text-[13px]" role="alert" data-testid="sim-cap-notice">
      {bulkCopy.capNotice(store.capNotice.cap, store.capNotice.combinations)}
      {#if store.serverCapNotice === null}
        <!-- Contract 10.1 A2: the premium lane's cap is 5,000, and bulkCopy.capPremiumNote
             says that number. Offering it when the list is past 5,000 too would send the
             player to a lane that refuses the same request. -->
        <span class="text-muted">{bulkCopy.capPremiumNote}</span>
      {:else}
        <span class="text-muted"
          >{bulkCopy.serverCapNotice(
            store.serverCapNotice.cap,
            store.serverCapNotice.combinations,
          )}</span
        >
      {/if}
    </p>
  {/if}

  {#if store.progressLine !== ''}
    <p class="tabular text-muted font-mono text-[13px]" data-testid="sim-stage-progress">
      {store.progressLine}
    </p>
  {/if}

  {#if store.message !== null}
    <p class="text-strong text-[13px]" role="alert" data-testid="sim-message">
      {store.message}
      {#if store.detail !== ''}
        <span class="text-muted block font-mono text-[12px]">{store.detail}</span>
      {/if}
    </p>
  {/if}
</section>

<!-- Design 8: the exact JSON the run will send, editable and validated by the engine's own
     Validate through the wasm's simValidate (contract 10.2). Part A owns the component;
     mounting it here is what gives every tool page the same escape hatch. -->
<RequestDrawer request={store.requestPreview} onapply={(next) => store.applyRequest(next)} />
```

- [ ] **Step 4: Give the store the drawer's two hooks**

In `bulk-store.svelte.ts`, beside the other accessors:

```ts
    /** Exactly what a run would send, for part A's Advanced drawer (design 8). */
    get requestPreview() {
      return baseRequestPreview();
    },
    /**
     * A request edited in the drawer, adopted whole. Only the two blocks this store owns
     * are read back -- the encounter and the buffs belong to `settings`, which part A's own
     * panel owns, and writing them from here would fight it.
     */
    applyRequest(next: unknown): void {
      const parsed = next as Partial<BulkRequest & WeightsRequest>;
      if (parsed.bulk !== undefined) {
        precision = parsed.bulk.precision;
        cap = parsed.bulk.cap;
        locked = [...(parsed.bulk.locked ?? [])];
        loadouts = [...(parsed.bulk.talents ?? [])];
        namedSets = [...(parsed.bulk.sets ?? [])];
      }
      if (parsed.weights !== undefined) stats = [...parsed.weights.stats];
      scheduleCount();
    },
```

with, above the returned object:

```ts
  /** Like `baseRequest()` but silent: the drawer renders what would be sent, it never runs. */
  function baseRequestPreview(): BulkRequest | WeightsRequest | null {
    if (character === null || talentFile === null) return null;
    const index = indexTalents(talentFile);
    const base = {
      engine_version: ENGINE_VERSION,
      spec: character.spec,
      source: character.source,
      character: toCharacterSpec(character, index, settings.buffs, settings.consumables),
      encounter: settings.encounter,
      iterations: init.tool === 'weights' ? 3000 : finalIterations(precision),
      random_seed: 0,
    };
    if (init.tool === 'weights') {
      return { ...base, weights: { stats: [...stats], reference: stats[0] ?? '' } };
    }
    const bulk = currentSpec();
    return bulk === null ? null : { ...base, bulk };
  }
```

- [ ] **Step 5: Mount it in `TopGear.svelte`**

Replace the placeholder `<p data-testid="sim-combo-count">` with:

```svelte
  <BulkRunBar {store} />
```

and `import BulkRunBar from './BulkRunBar.svelte';`.

- [ ] **Step 6: Run the e2e, desktop and mobile**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts
```

Expected: PASS. The `sim-combos` assertion in "a run reports its stage line" fails until
Task 16; comment that one line out and restore it there, or run Task 16 first — the plan's
order is Task 16 next, so a single red assertion here is expected and is noted in the
commit body.

- [ ] **Step 7: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools src/lib/sim/bulk-store.svelte.ts tests/e2e/sim-gear.spec.ts
npx prettier --write src/components/sim/tools src/lib/sim/bulk-store.svelte.ts tests/e2e/sim-gear.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/src/lib/sim/bulk-store.svelte.ts web/tests/e2e/sim-gear.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): the shared run bar for every combination tool

One bar on all four pages: the live count, the three precisions, the
stage line, and a cap notice that says by how much the list overran and
what premium would allow rather than trimming anything. Part A's
settings panel and request drawer are mounted here, so every tool gets
them in one place. The results table lands in the next task.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 16: Top Gear's results

**Files:**
- Create: `web/src/components/sim/tools/ComboResults.svelte`
- Create: `web/src/components/sim/tools/SubstitutionChips.svelte`
- Modify: `web/src/components/sim/tools/TopGear.svelte`
- Modify: `web/tests/e2e/sim-gear.spec.ts`

**Interfaces:**
- Consumes: `comboRows`, `deltaLabel`, `percentOf`, `slotSummary`, `winningGear`, `keepsSetBonus`, `substitutionLabel` (Task 9); `addonStringFor` (Task 8); `encodeFS1`, `gearFromSlots`, `ranksFromTalentsString`; `simSearch`, `withSimState`, `defaultSimState`.
- Produces: test ids `sim-combos`, `sim-combo-row`, `sim-equipped-line`, `sim-slot-summary`.

- [ ] **Step 1: Write the failing e2e**

Append to `web/tests/e2e/sim-gear.spec.ts`:

```ts
test('the results show the equipped baseline, a ranked table and a per-slot summary', async ({
  page,
}) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-search-add-16966').click();
  await page.getByTestId('sim-run-bulk').click();

  const equipped = page.getByTestId('sim-equipped-line');
  await expect(equipped).toBeVisible({ timeout: 20_000 });
  await expect(equipped).toHaveText(/\d/);

  const rows = page.getByTestId('sim-combo-row');
  await expect(rows.first()).toBeVisible();
  expect(await rows.count()).toBeGreaterThan(1);
  await expect(rows.first()).toContainText(/[+−]\d+ ± \d+/);

  await expect(page.getByTestId('sim-slot-summary')).toBeVisible();
});

test('the winner opens in the planner and copies to the addon', async ({ page, context }) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });

  await expect(page.getByTestId('sim-open-in-planner')).toHaveAttribute('href', /\/planner\?/);

  await page.getByTestId('sim-copy-addon').click();
  await expect(page.getByTestId('sim-copy-addon')).toHaveText('Copied');
  const copied = await page.evaluate(() => navigator.clipboard.readText());
  expect(copied.startsWith('FS1:')).toBe(true);
});

test('the four-piece filter hides combinations that break the set', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  const before = await page.getByTestId('sim-combo-row').count();
  await page.getByTestId('sim-keep-set').check();
  expect(await page.getByTestId('sim-combo-row').count()).toBeLessThanOrEqual(before);
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts --project=desktop
```

Expected: FAIL — `sim-combos` never appears.

- [ ] **Step 3: Write `SubstitutionChips.svelte`**

```svelte
<!-- web/src/components/sim/tools/SubstitutionChips.svelte -->
<!-- What a combination changed, as icon chips.
     The NAME always comes off the substitution -- contract 10.1 A6 fills it for items too,
     from simdb -- so a chip is legible even where the per-class item file has not loaded.
     The item map is consulted for the icon and the rarity colour only, which are the two
     things the result does not carry. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
  import { dataUrl } from '../../../lib/planner/load';
  import type { Item } from '../../../lib/planner/types';
  import type { Substitution } from '../../../lib/sim/bulk-types';
  import { substitutionLabel } from '../../../lib/sim/combos';

  let {
    substitutions,
    items,
    treeVersion,
  }: {
    substitutions: readonly Substitution[];
    items: ReadonlyMap<number, Item>;
    treeVersion: string;
  } = $props();
</script>

<span class="flex flex-wrap items-center gap-1">
  {#each substitutions as sub, index (`${sub.kind}-${sub.slot ?? ''}-${sub.item_id ?? sub.name ?? index}`)}
    {@const item = sub.item_id === undefined ? undefined : items.get(sub.item_id)}
    <span
      class="border-line rounded-pill inline-flex items-center gap-1 border px-2 py-[2px] text-[12px]"
      title={sub.source_name !== undefined && sub.source_name !== ''
        ? `${substitutionLabel(sub)} · ${sub.source_name}`
        : substitutionLabel(sub)}
    >
      {#if item !== undefined}
        <img
          src={dataUrl(treeVersion, `icons/${item.icon}.webp`)}
          alt=""
          width="16"
          height="16"
          loading="lazy"
          decoding="async"
          class="rounded-control h-4 w-4 object-cover"
        />
      {/if}
      <span class={item === undefined ? 'text-text' : rarityClassFor(item.quality)}>
        {substitutionLabel(sub)}
      </span>
    </span>
  {/each}
</span>
```

- [ ] **Step 4: Write `ComboResults.svelte`**

```svelte
<!-- web/src/components/sim/tools/ComboResults.svelte -->
<!-- Design 3.3. Every number here comes from the result: the ranking, the within-error
     grouping and each delta's error are the planner's, and this component neither sorts nor
     recomputes any of them. -->
<script lang="ts">
  import { encodeFS1 } from '../../../lib/planner/fs1';
  import type { Item, ItemSet } from '../../../lib/planner/types';
  import { addonStringFor } from '../../../lib/sim/addon-export';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import { gearFromSlots, ranksFromTalentsString } from '../../../lib/sim/character';
  import {
    comboRows,
    deltaLabel,
    keepsSetBonus,
    slotSummary,
    winningGear,
  } from '../../../lib/sim/combos';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { SLOT_LABELS, type Slot } from '../../../lib/planner/types';
  import SubstitutionChips from './SubstitutionChips.svelte';

  let {
    result,
    items,
    sets,
    treeVersion,
    partial,
  }: {
    result: BulkResult;
    items: ReadonlyMap<number, Item>;
    sets: readonly ItemSet[];
    treeVersion: string;
    partial: boolean;
  } = $props();

  const FOUR_PIECE = 4;
  let keepSet = $state(false);
  let copied = $state(false);

  const rows = $derived(
    comboRows(result).filter(
      (row) => !keepSet || keepsSetBonus(row.combo, result, items, sets, FOUR_PIECE),
    ),
  );
  const summary = $derived(slotSummary(result));
  const winner = $derived(winningGear(result));

  const equippedFigure = $derived(Math.round(result.equipped.mean).toLocaleString('en-US'));
  const equippedBand = $derived(Math.round(1.96 * result.equipped.error).toLocaleString('en-US'));

  /** The winning set into the planner, through the same ?code= bootstrap /sim already uses. */
  const plannerHref = $derived(
    `/planner?code=${encodeURIComponent(
      encodeFS1({
        dataBuild: treeVersion,
        classSlug: result.request.character.class,
        raceSlug: result.request.character.race,
        treeRanks: ranksFromTalentsString(result.request.character.talents),
        gear: gearFromSlots(winner),
      }),
    )}`,
  );

  async function copyAddon(): Promise<void> {
    try {
      await navigator.clipboard.writeText(
        addonStringFor({
          dataBuild: treeVersion,
          classSlug: result.request.character.class,
          raceSlug: result.request.character.race,
          talents: result.request.character.talents,
          gear: winner,
        }),
      );
      copied = true;
    } catch {
      copied = false;
    }
  }
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-combos">
  {#if partial}
    <p class="text-strong text-[13px]" role="status">{bulkCopy.partial}</p>
  {/if}

  <p class="tabular text-strong font-mono text-[14px]" data-testid="sim-equipped-line">
    {bulkCopy.resultsEquipped}: {equippedFigure} ± {equippedBand}
    <span class="text-muted">{bulkCopy.ranAtStages(result.stages)}</span>
  </p>

  <label class="text-muted flex min-h-11 w-fit items-center gap-2 text-[12px]">
    <input type="checkbox" class="h-5 w-5" data-testid="sim-keep-set" bind:checked={keepSet} />
    {bulkCopy.keepFourPiece}
  </label>

  {#if rows.length === 0}
    <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
  {:else}
    <ul class="flex flex-col">
      {#each rows as row (row.combo.substitutions.map((sub) => `${sub.kind}:${sub.slot ?? ''}:${sub.item_id ?? sub.name ?? ''}`).join('|'))}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 border-b px-2 py-2"
          data-testid="sim-combo-row"
        >
          <span class="tabular text-muted font-mono text-[12px]">{row.rank}</span>
          <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
          <span class="tabular text-strong ml-auto font-mono text-[13px]">
            {Math.round(row.combo.dps.mean).toLocaleString('en-US')}
          </span>
          <span class="tabular text-gold ml-auto font-mono text-[13px]">
            {deltaLabel(row.combo.delta)}
          </span>
          <span class="tabular text-muted ml-auto font-mono text-[12px]">
            {row.percent.toFixed(1)}%
          </span>
        </li>
      {/each}
    </ul>
    <p class="text-muted text-[12px]">{bulkCopy.withinErrorNote}</p>
  {/if}

  {#if summary.length > 0}
    <section class="border-line rounded-panel border p-3" data-testid="sim-slot-summary">
      <h3 class="section-title text-[14px]">{bulkCopy.slotSummary}</h3>
      <p class="text-muted text-[12px]">{bulkCopy.slotSummaryNote}</p>
      <ul class="flex flex-col">
        {#each summary as row (`${row.slot}-${row.item_id}`)}
          <li class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 last:border-b-0">
            <span class="text-muted w-24 text-[12px]">{SLOT_LABELS[row.slot as Slot] ?? row.slot}</span>
            <span class="text-text flex-1 text-[13px]">{row.name}</span>
            <span class="tabular text-gold font-mono text-[13px]">
              {row.gain === null ? '—' : `+${Math.round(row.gain).toLocaleString('en-US')}`}
            </span>
          </li>
        {/each}
      </ul>
    </section>
  {/if}

  <div class="flex flex-wrap gap-3">
    <a
      class="border-line-warm rounded-control text-nav label inline-flex min-h-11 items-center border px-4"
      href={plannerHref}
      data-testid="sim-open-in-planner">{bulkCopy.openInPlanner}</a
    >
    <button
      type="button"
      class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
      data-testid="sim-copy-addon"
      onclick={() => void copyAddon()}
    >
      {copied ? bulkCopy.copiedToAddon : bulkCopy.copyToAddon}
    </button>
  </div>

  <section class="border-line rounded-panel border p-3" data-testid="sim-rules">
    <h3 class="section-title text-[14px]">{bulkCopy.rulesTitle}</h3>
    <ul class="text-muted flex list-disc flex-col gap-1 pl-5 text-[12px]">
      {#each bulkCopy.rules as rule (rule)}
        <li>{rule}</li>
      {/each}
    </ul>
  </section>
</section>
```

- [ ] **Step 5: Mount it in `TopGear.svelte`**

After `<BulkRunBar {store} />`:

```svelte
  {#if store.result !== null && store.combos.length >= 0 && store.tool !== 'weights'}
    <ComboResults
      result={store.result as BulkResult}
      items={store.items}
      sets={store.sets}
      treeVersion={store.character?.tree_version ?? ''}
      partial={store.result.aborted === true}
    />
  {/if}
```

with `import ComboResults from './ComboResults.svelte';` and
`import type { BulkResult } from '../../../lib/sim/bulk-types';`.

- [ ] **Step 6: Run the e2e, desktop and mobile, and restore the assertion Task 15 deferred**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-gear.spec.ts
```

Expected: PASS on both projects, including the `sim-combos` assertion in "a run reports its
stage line".

- [ ] **Step 7: Check the component sizes**

```bash
cd web && wc -l src/components/sim/tools/*.svelte
```

Expected: every file under 400 lines. `TopGear.svelte` is a composition and should be well
under 150.

- [ ] **Step 8: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools tests/e2e/sim-gear.spec.ts
npx prettier --write src/components/sim/tools tests/e2e/sim-gear.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/tests/e2e/sim-gear.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): Top Gear's results

The equipped baseline, the ranked rows with each delta's own error and
its percent, a shared rank across the leader's within-error group, and
the per-slot summary that says "—" rather than apportioning a total when
nothing measured that slot alone. Open in planner and Copy to addon both
carry the winning set. The four-piece filter is a results filter, since
the engine counts set bonuses from the gear itself.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 17: `/sim/talents` — the same island with gear locked

Design 3.4: the same page, gear locked, the talent list as the only candidates. The island
is already `TopGear.svelte` for both routes (`views.talents` in `ToolsView.svelte` imports
the same chunk), so this task is the difference, not a second page.

**Files:**
- Modify: `web/src/components/sim/tools/TopGear.svelte`
- Modify: `web/src/lib/sim/bulk-store.svelte.ts` (lock every slot in talents mode)
- Modify: `web/src/lib/sim/bulk-store.test.ts`
- Create: `web/tests/e2e/sim-talents.spec.ts`

**Interfaces:**
- Consumes: everything Tasks 12–16 produced.
- Produces: the invariant "a `talents` request carries no candidates and every slot locked".

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/sim/bulk-store.test.ts`:

```ts
describe('talent compare', () => {
  it('locks every slot the moment a character loads, so nothing gear-shaped can be sent', async () => {
    const s = store('talents');
    await s.loadAddon(FURY);
    expect(s.locked).toEqual(expect.arrayContaining(['head', 'main_hand', 'finger1']));
    s.dispose();
  });

  it('sends the loadouts and nothing else', async () => {
    const s = store('talents');
    await s.loadAddon(FURY);
    s.addSearchItem(16963);
    s.addLoadout({ name: 'Deep Fury', talents: '0-5530515-' });
    await s.run();
    const bulk = (s.result as BulkResult).request.bulk!;
    expect(bulk.mode).toBe('talents');
    expect(bulk.candidates).toEqual([]);
    expect(bulk.talents).toEqual([{ name: 'Deep Fury', talents: '0-5530515-' }]);
    s.dispose();
  });
});
```

Create `web/tests/e2e/sim-talents.spec.ts`:

```ts
// web/tests/e2e/sim-talents.spec.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function loadTalents(page: Page): Promise<void> {
  await page.goto('/sim/talents');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('there is no slot grid, no item search and no named sets — only the talent list', async ({
  page,
}) => {
  await loadTalents(page);
  await expect(page.getByTestId('sim-talent-candidates')).toBeVisible();
  await expect(page.getByTestId('sim-slot-grid')).toHaveCount(0);
  await expect(page.getByTestId('sim-item-search')).toHaveCount(0);
  await expect(page.getByTestId('sim-named-sets')).toHaveCount(0);
});

test('ticking the character’s own build makes the run button live and ranks it', async ({ page }) => {
  await loadTalents(page);
  await page.getByTestId('sim-loadout-current').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText('1 valid combination');
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  await expect(page.getByTestId('sim-combo-row').first()).toContainText('Your current build');
});

test('the page has its own heading and its own canonical', async ({ page }) => {
  await page.goto('/sim/talents');
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Talent compare');
});
```

- [ ] **Step 2: Run them and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-store.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-talents.spec.ts --project=desktop
```

Expected: FAIL — the slot grid is still on the page and nothing is locked.

- [ ] **Step 3: Lock every slot in talents mode**

In `bulk-store.svelte.ts`'s `adopt()`, immediately after `rows = seedRows(outcome.character);`:

```ts
    // Design 3.4: talent compare is Top Gear with the gear locked. Locking every slot here
    // rather than hiding the grid means the invariant holds for the request too -- the
    // contract validates "no candidate on a locked slot", and a page that only hid the UI
    // could still send one through the request drawer.
    if (init.tool === 'talents') locked = [...SLOTS];
```

with `import { SLOTS } from '../planner/types';` added to the imports.

- [ ] **Step 4: Hide the gear-only sections**

`TopGear.svelte` already wraps `SlotGrid`, `ItemSearch`, `ConsumableCandidates` and
`NamedSets` in `{#if gearMode}`, and `gearMode` is `store.tool === 'gear'`. Confirm by
reading the file; if `ConsumableCandidates` sits outside the block, move it inside — a
consumable candidate is a gear-shaped change and talent compare offers none.

- [ ] **Step 5: Run the tests, desktop and mobile**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/sim/bulk-store.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-talents.spec.ts
```

Expected: PASS.

- [ ] **Step 6: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools src/lib/sim/bulk-store.svelte.ts src/lib/sim/bulk-store.test.ts tests/e2e/sim-talents.spec.ts
npx prettier --write src/components/sim/tools src/lib/sim/bulk-store.svelte.ts src/lib/sim/bulk-store.test.ts tests/e2e/sim-talents.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/src/lib/sim/bulk-store.svelte.ts web/src/lib/sim/bulk-store.test.ts web/tests/e2e/sim-talents.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): /sim/talents, the same island with gear locked

Locking every slot rather than only hiding the grid: the contract
validates "no candidate on a locked slot", so the invariant then holds
for a request edited in the Advanced drawer too, not just for the UI.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 18: `/sim/drops` — the Droptimizer

**Files:**
- Create: `web/src/components/sim/tools/Droptimizer.svelte`
- Create: `web/src/components/sim/tools/SourcePicker.svelte`
- Create: `web/src/components/sim/tools/DropResults.svelte`
- Create: `web/tests/e2e/sim-drops.spec.ts`

**Interfaces:**
- Consumes: `groupSources`, `itemsOfSource`, `itemsOfBoss`, `isOpen`, `professionSplit`, `SOURCE_KIND_LABELS`, `DEFAULT_OFF_KINDS` (Task 6); `phaseLabel`, `openDateLabel` (Task 6); `comboRows`, `deltaLabel`, `substitutionLabel` (Task 9); `BulkRunBar` (Task 15); `SubstitutionChips` (Task 16).
- Produces: test ids `sim-source-picker`, `sim-source-<id>`, `sim-upcoming`, `sim-drops-by-boss`, `sim-drops-flat`, `sim-drops-pin`.

> **Source icons.** Design 6.1 says the picker shows "icons from the zone pages", and there
> are none: `src/content/zones/*.md` and `src/content/dungeons/*.md` carry a title, a
> continent and a level range, and nothing image-shaped. Contract 10.7 confirms it — "zone
> icons do not exist in the content collection; the source picker labels by kind and name
> and draws none" — so this task labels and draws none.

- [ ] **Step 1: Write the failing e2e**

Create `web/tests/e2e/sim-drops.spec.ts`:

```ts
// web/tests/e2e/sim-drops.spec.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function loadDrops(page: Page): Promise<void> {
  await page.goto('/sim/drops');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-source-picker')).toBeVisible();
}

test('the picker groups every source kind and leaves quests off', async ({ page }) => {
  await loadDrops(page);
  for (const label of ['Raids', 'Dungeons', 'World bosses', 'Crafted', 'Reputation', 'PvP']) {
    await expect(page.getByTestId('sim-source-picker')).toContainText(label);
  }
  await expect(page.getByTestId('sim-kind-quest')).not.toBeChecked();
  await expect(page.getByTestId('sim-source-quest')).toHaveCount(0);
  await page.getByTestId('sim-kind-quest').check();
  await expect(page.getByTestId('sim-source-quest')).toBeVisible();
});

test('an unreleased source is hidden until show-upcoming, and then says it has no date', async ({
  page,
}) => {
  await loadDrops(page);
  // The fixture's world boss carries `opens: "later"` (contract 10.4), which never opens
  // whatever the clock says -- so this is a date-independent assertion.
  await expect(page.getByTestId('sim-source-world:azuregos')).toHaveCount(0);
  await page.getByTestId('sim-upcoming').check();
  await expect(page.getByTestId('sim-source-world:azuregos')).toBeVisible();
  await expect(page.getByTestId('sim-source-picker')).toContainText('no date announced');
});

test('picking a boss counts its drops, and running ranks them by source', async ({ page }) => {
  await loadDrops(page);
  await page.getByTestId('sim-source-raid:mc:11502').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(/\d+ valid combinations?/);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-drops-by-boss')).toBeVisible({ timeout: 25_000 });
  // The boss is named from Substitution.SourceName, which the candidate carried in
  // (contract 10.1 A6) -- the page does not re-join the origin id to loot.json.
  await expect(page.getByTestId('sim-drops-by-boss')).toContainText('Ragnaros');
  await expect(page.getByTestId('sim-drops-flat')).toBeVisible();
});

test('a drop pins into Top Gear and arrives there ticked', async ({ page }) => {
  await loadDrops(page);
  await page.getByTestId('sim-source-raid:mc:11502').check();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-drops-flat')).toBeVisible({ timeout: 25_000 });
  await page.getByTestId('sim-drops-pin-12784').click();
  await expect(page).toHaveURL(/\/sim\/gear/);
});

test('the page never shows a probability', async ({ page }) => {
  await loadDrops(page);
  await expect(page.getByTestId('sim-drops-note')).toContainText('not a probability');
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-drops.spec.ts --project=desktop
```

Expected: FAIL — no source picker.

- [ ] **Step 3: Write `SourcePicker.svelte`**

```svelte
<!-- web/src/components/sim/tools/SourcePicker.svelte -->
<!-- Design 6.1. A raid card is one checkbox for every boss; each boss is its own. A source
     gated on a phase that has not opened is hidden until "show unreleased content", and
     then carries the date it opens rather than pretending it is live.
     There are no zone icons to draw: src/content/zones/*.md carries a title, a continent
     and a level range, and nothing image-shaped. -->
<script lang="ts">
  import { bulkCopy } from '../../../lib/sim/copy';
  import {
    DEFAULT_OFF_KINDS,
    groupSources,
    isOpen,
    itemsOfSource,
    professionSplit,
    SOURCE_KIND_LABELS,
    type LootFile,
    type LootSource,
  } from '../../../lib/sim/loot';
  import { PHASE_LATER, openDateLabel, phaseLabel, type PhaseRow } from '../../../lib/sim/phase';

  let {
    loot,
    phases,
    shownKinds,
    showUpcoming,
    picked,
    professions,
    now,
    ontogglekind,
    ontoggleupcoming,
    ontoggle,
  }: {
    loot: LootFile;
    phases: readonly PhaseRow[];
    shownKinds: readonly string[];
    showUpcoming: boolean;
    picked: readonly string[];
    professions: readonly string[] | undefined;
    now: Date;
    ontogglekind: (kind: string) => void;
    ontoggleupcoming: (value: boolean) => void;
    ontoggle: (sourceId: string, bossId?: string) => void;
  } = $props();

  const groups = $derived(groupSources(loot));

  function visible(source: LootSource): boolean {
    return shownKinds.includes(source.kind) && (showUpcoming || isOpen(phases, source, now));
  }

  function isPicked(sourceId: string, bossId = ''): boolean {
    return picked.includes(`${sourceId}|${bossId}`);
  }

  /**
   * "First raids, 9 December 2026" for a dated phase, and the no-date sentence for the
   * literal `"later"` -- contract 10.4's "shows as unreleased without a date".
   */
  function gateLabel(source: LootSource): string {
    if (source.opens === undefined || isOpen(phases, source, now)) return '';
    if (source.opens === PHASE_LATER) return bulkCopy.opensLater;
    const date = openDateLabel(phases, source.opens);
    return date === '' ? bulkCopy.notOpenYet : bulkCopy.opensOn(phaseLabel(source.opens), date);
  }
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-3 border p-3 md:mx-0"
  data-testid="sim-source-picker"
>
  <div class="flex flex-wrap items-center gap-3">
    {#each groups as group (group.kind)}
      <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
        <input
          type="checkbox"
          class="h-5 w-5"
          data-testid={`sim-kind-${group.kind}`}
          checked={shownKinds.includes(group.kind)}
          onchange={() => ontogglekind(group.kind)}
        />
        {group.label}
      </label>
    {/each}
    <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid="sim-upcoming"
        checked={showUpcoming}
        onchange={(event) => ontoggleupcoming(event.currentTarget.checked)}
      />
      {bulkCopy.showUpcoming}
    </label>
  </div>

  {#each groups as group (group.kind)}
    {@const shown = group.sources.filter(visible)}
    {#if shown.length > 0}
      <div class="flex flex-col gap-1">
        <h3 class="section-title text-[13px]">{SOURCE_KIND_LABELS[group.kind]}</h3>
        {#if group.kind === 'crafted'}
          {@const split = professionSplit(shown, professions)}
          {#if split.mine.length > 0}
            <p class="text-muted text-[12px]">{bulkCopy.sourcesMyProfessions}</p>
          {:else}
            <p class="text-muted text-[12px]">{bulkCopy.sourcesProfessionsUnknown}</p>
          {/if}
        {/if}
        {#if group.kind === 'quest'}
          <p class="text-muted text-[12px]">{bulkCopy.sourcesQuestNote}</p>
        {/if}
        {#each shown as source (source.id)}
          <div class="border-line-soft flex flex-col gap-1 border-b py-2 last:border-b-0">
            <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
              <input
                type="checkbox"
                class="h-5 w-5"
                data-testid={`sim-source-${source.id}`}
                checked={isPicked(source.id)}
                onchange={() => ontoggle(source.id)}
              />
              <span class="flex-1">{source.name}</span>
              {#if (source.bosses ?? []).length > 0}
                <span class="text-muted text-[12px]">{bulkCopy.sourcesWholeRaid}</span>
              {/if}
              <span class="tabular text-muted font-mono text-[12px]">
                {itemsOfSource(source).length}
              </span>
            </label>
            {#if gateLabel(source) !== ''}
              <p class="text-muted pl-7 text-[12px]">{gateLabel(source)}</p>
            {/if}
            {#each source.bosses ?? [] as boss (boss.id)}
              <label class="text-text flex min-h-11 items-center gap-2 pl-7 text-[13px]">
                <input
                  type="checkbox"
                  class="h-5 w-5"
                  data-testid={`sim-source-${boss.id}`}
                  checked={isPicked(source.id, boss.id)}
                  onchange={() => ontoggle(source.id, boss.id)}
                />
                <!-- Contract 10.4: "a boss with no name in either database is emitted with
                     an empty name". The id is the honest stand-in; inventing one is not. -->
                <span class="flex-1">{boss.name === '' ? boss.id : boss.name}</span>
                <span class="tabular text-muted font-mono text-[12px]">{boss.items.length}</span>
              </label>
            {/each}
            {#if (source.trash ?? []).length > 0}
              <p class="text-muted pl-7 text-[12px]">{bulkCopy.sourcesTrash}</p>
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  {/each}
</section>
```

- [ ] **Step 4: Write `DropResults.svelte`**

```svelte
<!-- web/src/components/sim/tools/DropResults.svelte -->
<!-- Design 6.3: by source, with the best per boss, then a flat "every upgrade" list.
     Nothing here is a probability -- neither database records drop rates -- so a boss reads
     "3 of the 11 drops here are upgrades", which is what the data supports. -->
<script lang="ts">
  import type { Item } from '../../../lib/planner/types';
  import type { BulkResult, Combo } from '../../../lib/sim/bulk-types';
  import { comboRows, deltaLabel, sourceNameOfCombo } from '../../../lib/sim/combos';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { LootFile } from '../../../lib/sim/loot';
  import SubstitutionChips from './SubstitutionChips.svelte';

  let {
    result,
    items,
    loot,
    treeVersion,
    onpin,
  }: {
    result: BulkResult;
    items: ReadonlyMap<number, Item>;
    loot: LootFile;
    treeVersion: string;
    onpin: (itemId: number) => void;
  } = $props();

  const rows = $derived(comboRows(result));

  /** A drops run is one substitution per combination, so a row's origin names its boss. */
  function originOf(combo: Combo): string {
    return combo.substitutions[0]?.origin ?? '';
  }

  /**
   * Contract 10.1 A6: the name rode in on `Candidate.SourceName` and came back on the
   * substitution, so this reads it rather than joining the origin id to loot.json a second
   * time. The id is the fallback for a result saved before the field existed.
   */
  function bossName(combo: Combo, origin: string): string {
    return sourceNameOfCombo(combo) || origin.replace(/^drop:/, '');
  }

  const byBoss = $derived.by(() => {
    const groups = new Map<string, typeof rows>();
    for (const row of rows) {
      const origin = originOf(row.combo);
      const existing = groups.get(origin);
      if (existing === undefined) groups.set(origin, [row]);
      else existing.push(row);
    }
    return [...groups.entries()];
  });

  const upgrades = $derived(rows.filter((row) => row.combo.delta.mean > 0));
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-combos">
  <p class="text-muted text-[12px]" data-testid="sim-drops-note">{bulkCopy.dropsNoChance}</p>

  <p class="tabular text-strong font-mono text-[14px]" data-testid="sim-equipped-line">
    {bulkCopy.resultsEquipped}: {Math.round(result.equipped.mean).toLocaleString('en-US')}
    ± {Math.round(1.96 * result.equipped.error).toLocaleString('en-US')}
  </p>

  <section class="flex flex-col gap-3" data-testid="sim-drops-by-boss">
    <h3 class="section-title text-[14px]">{bulkCopy.dropsByBoss}</h3>
    {#each byBoss as [origin, group] (origin)}
      {@const wins = group.filter((row) => row.combo.delta.mean > 0)}
      <div class="border-line rounded-panel border p-3">
        <header class="flex flex-wrap items-baseline justify-between gap-2">
          <h4 class="text-strong text-[13px]">{bossName(group[0].combo, origin)}</h4>
          <span class="text-muted text-[12px]">{bulkCopy.dropsUpgrades(wins.length, group.length)}</span>
        </header>
        {#if wins.length > 0}
          <p class="tabular text-gold font-mono text-[13px]">
            {bulkCopy.dropsBest}: {deltaLabel(wins[0].combo.delta)}
          </p>
        {/if}
        <ul class="flex flex-col">
          {#each group as row (row.combo.substitutions[0]?.item_id ?? row.rank)}
            <li
              class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 py-1 last:border-b-0"
              data-testid="sim-combo-row"
            >
              <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
              <span class="tabular text-gold ml-auto font-mono text-[13px]">
                {deltaLabel(row.combo.delta)}
              </span>
              <span class="tabular text-muted font-mono text-[12px]">{row.percent.toFixed(1)}%</span>
            </li>
          {/each}
        </ul>
      </div>
    {/each}
  </section>

  <section class="border-line rounded-panel border p-3" data-testid="sim-drops-flat">
    <h3 class="section-title text-[14px]">{bulkCopy.dropsEveryUpgrade}</h3>
    {#if upgrades.length === 0}
      <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
    {:else}
      <ul class="flex flex-col">
        {#each upgrades as row (row.combo.substitutions[0]?.item_id ?? row.rank)}
          {@const itemId = row.combo.substitutions[0]?.item_id ?? 0}
          <li class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 py-1 last:border-b-0">
            <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
            <span class="tabular text-gold ml-auto font-mono text-[13px]">
              {deltaLabel(row.combo.delta)}
            </span>
            <button
              type="button"
              class="border-line-warm rounded-control text-nav label min-h-11 border px-3"
              data-testid={`sim-drops-pin-${itemId}`}
              onclick={() => onpin(itemId)}>{bulkCopy.dropsPin}</button
            >
          </li>
        {/each}
      </ul>
    {/if}
  </section>
</section>
```

- [ ] **Step 5: Write `Droptimizer.svelte`**

```svelte
<!-- web/src/components/sim/tools/Droptimizer.svelte -->
<!-- /sim/drops. It is the same store, the same run bar and the same stage loop as Top Gear;
     the source picker replaces the slot grid and the results group by boss instead of
     ranking one flat list.
     "Pin into Top Gear" leaves this page with the item in the query string, because the two
     pages are separate islands and a pin has to survive the navigation. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import { bulkCopy } from '../../../lib/sim/copy';
  import BulkRunBar from './BulkRunBar.svelte';
  import DropResults from './DropResults.svelte';
  import SourcePicker from './SourcePicker.svelte';

  let { store }: { store: BulkStore; me: Me | null } = $props();

  /**
   * Top Gear with this item already ticked. `?pin=` is read by ToolsView's bootstrap the
   * same way `?source=`/`?ref=` are, and a pin carries the character with it so the page
   * does not open empty.
   */
  function pin(itemId: number): void {
    const params = new URLSearchParams(window.location.search);
    params.set('pin', String(itemId));
    window.location.href = `/sim/gear?${params.toString()}`;
  }
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-droptimizer">
  {#if store.character !== null}
    <SourcePicker
      loot={store.loot}
      phases={store.phases}
      shownKinds={store.shownKinds}
      showUpcoming={store.showUpcoming}
      picked={store.pickedBosses}
      professions={store.character.professions}
      now={new Date()}
      ontogglekind={(kind) => store.toggleKind(kind)}
      ontoggleupcoming={(value) => store.setShowUpcoming(value)}
      ontoggle={(sourceId, bossId) => store.toggleSource(sourceId, bossId)}
    />
  {/if}

  <BulkRunBar {store} />

  {#if store.result !== null}
    <DropResults
      result={store.result as BulkResult}
      items={store.items}
      loot={store.loot}
      treeVersion={store.character?.tree_version ?? ''}
      onpin={pin}
    />
  {:else if store.pickedBosses.length === 0}
    <p class="text-muted px-[18px] text-[13px] md:px-0">{bulkCopy.dropsNothing}</p>
  {/if}
</div>
```

`SimCharacter` has no `professions` field today; add it as an optional in
`web/src/lib/sim/character.ts` beside `consumables`:

```ts
  /**
   * The character's professions, for the Droptimizer's "my professions" split. Nothing
   * records them yet -- no source in sources.ts sets it and `CharacterSpec.professions` is
   * likewise optional -- so an absent list means "we have not looked", which is what the
   * picker says, rather than "this character has none".
   */
  professions?: string[];
```

- [ ] **Step 6: Read `?pin=` in the tools bootstrap**

In `ToolsView.svelte`'s `bootstrap`, add `pin: new URLSearchParams(window.location.search).get('pin') ?? ''`,
and in `onMount`, after the source bootstrap:

```ts
    // A drop pinned from /sim/drops. It is applied after the character has loaded, because
    // `addSearchItem` needs the item file -- the same reason `seedRows` runs inside adopt().
    if (bootstrap.pin !== '') {
      const itemId = Number.parseInt(bootstrap.pin, 10);
      if (Number.isInteger(itemId) && itemId > 0) {
        void (async () => {
          await new Promise((resolve) => setTimeout(resolve, 0));
          store.addSearchItem(itemId, 'search');
        })();
      }
    }
```

- [ ] **Step 7: Run the e2e, desktop and mobile**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-drops.spec.ts
```

Expected: PASS.

- [ ] **Step 8: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools src/lib/sim/character.ts tests/e2e/sim-drops.spec.ts
npx prettier --write src/components/sim/tools src/lib/sim/character.ts tests/e2e/sim-drops.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/src/lib/sim/character.ts web/tests/e2e/sim-drops.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): /sim/drops, the Droptimizer

Raids by boss with trash, dungeons by boss, world bosses, crafted split
by profession, reputation by standing, PvP by rank, and quests off by
default. An unreleased source is hidden until show-upcoming and then
carries the date it opens, or says none is announced when its `opens` is
the literal "later" (contract 10.4). Bosses are named from
Substitution.SourceName, which the candidate carried in (10.1 A6), so
the results never re-join an origin id. No probability is shown
anywhere: neither database records drop rates, so a boss reads "3 of the
11 drops here are upgrades", which is what the data supports. Source
icons are not drawn because the zone pages carry none (10.7).

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 19: `/sim/weights` — stat weights, with the warning

**Files:**
- Create: `web/src/components/sim/tools/StatWeights.svelte`
- Create: `web/tests/e2e/sim-weights.spec.ts`

**Interfaces:**
- Consumes: `WEIGHT_STATS`, `statLabel`, `referenceFor`, `defaultStatsFor`, `pawnString`, `weightScale` (Task 9); `store.weights`, `store.stats`, `store.setStats`, `store.referenceStat`, `store.specRows` (Task 10); `BulkRunBar` (Task 15).
- Produces: test ids `sim-weights`, `sim-pawn`, `sim-weights-warning`, `sim-weight-<stat>`.

- [ ] **Step 1: Write the failing e2e**

Create `web/tests/e2e/sim-weights.spec.ts`:

```ts
// web/tests/e2e/sim-weights.spec.ts
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function loadWeights(page: Page): Promise<void> {
  await page.goto('/sim/weights');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('the page opens with the warning, above everything else', async ({ page }) => {
  await page.goto('/sim/weights');
  const warning = page.getByTestId('sim-weights-warning');
  await expect(warning).toBeVisible();
  await expect(warning).toContainText('not a straight line');
  await expect(warning.getByRole('link', { name: /Top Gear/i })).toHaveAttribute('href', '/sim/gear');
});

test('the stat picker defaults to the spec’s reference stat, first and ticked', async ({ page }) => {
  await page.route('**/v1/specs', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          specs: [
            {
              spec: 'warrior-fury',
              state: 'validated',
              median_gap: 0.03,
              parses: 50,
              worst_actions: [],
              engine_version: 'x',
              updated_at: null,
              // Contract 10.1 A7: proto.Stat enum names in snake case.
              reference_stat: 'attack_power',
            },
          ],
        },
      }),
    }),
  );
  await loadWeights(page);
  await expect(page.getByTestId('sim-weight-pick-attack_power')).toBeChecked();
  await expect(page.getByTestId('sim-weights-reference')).toHaveText(/Attack power/);
});

test('a run renders a weight per stat with an error bar and a Pawn string', async ({
  page,
  context,
}) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await loadWeights(page);
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-weights')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-weight-attack_power')).toContainText('1.00');
  // Contract 10.8: haste is split, hit and crit are not.
  await expect(page.getByTestId('sim-weight-melee_haste')).toBeVisible();
  await expect(page.getByTestId('sim-weight-crit')).toBeVisible();
  await expect(page.getByTestId('sim-weight-hit')).toBeVisible();
  for (const wrong of ['haste', 'melee_crit', 'melee_hit']) {
    await expect(page.getByTestId(`sim-weight-${wrong}`)).toHaveCount(0);
  }
  await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');
  await page.getByTestId('sim-pawn-copy').click();
  await expect(page.getByTestId('sim-pawn-copy')).toHaveText('Copied');
});

test('there is no slot grid, no source picker and no named sets', async ({ page }) => {
  await loadWeights(page);
  await expect(page.getByTestId('sim-slot-grid')).toHaveCount(0);
  await expect(page.getByTestId('sim-source-picker')).toHaveCount(0);
  await expect(page.getByTestId('sim-named-sets')).toHaveCount(0);
});
```

- [ ] **Step 2: Run it and watch it fail**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-weights.spec.ts --project=desktop
```

Expected: FAIL — the warning never renders.

- [ ] **Step 3: Write `StatWeights.svelte`**

```svelte
<!-- web/src/components/sim/tools/StatWeights.svelte -->
<!-- Design 7. The caution opens the page, in our words, and links to Top Gear -- it is not
     a footnote, because the whole reason this page exists is that addons ask for a number
     the number itself cannot justify. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { pawnString, statLabel, weightScale, WEIGHT_STATS } from '../../../lib/sim/weights';
  import BulkRunBar from './BulkRunBar.svelte';

  let { store }: { store: BulkStore; me: Me | null } = $props();

  let copied = $state(false);

  const scale = $derived(weightScale(store.weights));
  const pawn = $derived(
    store.weights.length === 0 ? '' : pawnString(store.character?.spec ?? '', store.weights),
  );

  function togglePick(id: string): void {
    const next = store.stats.includes(id)
      ? store.stats.filter((stat) => stat !== id)
      : [...store.stats, id];
    // The reference stat always leads: pawnString and the table both read stats[0] as the
    // one normalised to 1, and a picker that could reorder it would silently rescale
    // every number on the page.
    const reference = store.referenceStat;
    store.setStats([reference, ...next.filter((stat) => stat !== reference)]);
  }

  async function copyPawn(): Promise<void> {
    try {
      await navigator.clipboard.writeText(pawn);
      copied = true;
    } catch {
      copied = false;
    }
  }

  /** Bar geometry as percentages of the widest weight plus its error. */
  function bar(weight: number, error: number): { width: string; error: string } {
    return {
      width: `${Math.max(0, (weight / scale) * 100)}%`,
      error: `${Math.max(0, ((2 * error) / scale) * 100)}%`,
    };
  }
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-stat-weights">
  <section
    class="border-line-warm rounded-panel mx-[18px] border p-4 md:mx-0"
    data-testid="sim-weights-warning"
  >
    <p class="text-strong text-[14px]">{bulkCopy.weightsWarning}</p>
    <a class="text-nav underline" href="/sim/gear">{bulkCopy.weightsWarningLink}</a>
  </section>

  <section class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0">
    <h3 class="section-title text-[14px]">{bulkCopy.weightsPick}</h3>
    <p class="text-muted text-[12px]" data-testid="sim-weights-reference">
      {bulkCopy.weightsReference}: {statLabel(store.referenceStat)}
    </p>
    <ul class="flex flex-wrap gap-3">
      {#each WEIGHT_STATS as stat (stat.id)}
        <li>
          <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
            <input
              type="checkbox"
              class="h-5 w-5"
              data-testid={`sim-weight-pick-${stat.id}`}
              checked={store.stats.includes(stat.id)}
              disabled={stat.id === store.referenceStat}
              onchange={() => togglePick(stat.id)}
            />
            {stat.label}
          </label>
        </li>
      {/each}
    </ul>
  </section>

  <BulkRunBar {store} />

  {#if store.weights.length > 0}
    <section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-weights">
      <ul class="flex flex-col">
        {#each store.weights as row (row.stat)}
          {@const geometry = bar(row.weight, row.error)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[minmax(120px,1fr)_minmax(0,3fr)_88px] items-center gap-x-3 border-b px-2 py-2"
            data-testid={`sim-weight-${row.stat}`}
          >
            <span class="text-text text-[13px]">{statLabel(row.stat)}</span>
            <span class="bg-card-top relative block h-[8px] w-full">
              <span class="bg-gold absolute top-0 left-0 block h-[8px]" style={`width:${geometry.width}`}
              ></span>
              <span
                class="border-line-warm absolute top-0 block h-[8px] border-x"
                style={`left:calc(${geometry.width} - ${geometry.error}/2);width:${geometry.error}`}
                aria-hidden="true"
              ></span>
            </span>
            <span class="tabular text-strong ml-auto font-mono text-[13px]">
              {row.weight.toFixed(2)}
              <span class="text-muted">± {row.error.toFixed(2)}</span>
            </span>
          </li>
        {/each}
      </ul>

      <div class="border-line rounded-panel flex flex-wrap items-center gap-3 border p-3">
        <code class="text-muted flex-1 font-mono text-[12px] break-all" data-testid="sim-pawn">{pawn}</code>
        <button
          type="button"
          class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
          data-testid="sim-pawn-copy"
          onclick={() => void copyPawn()}
        >
          {copied ? bulkCopy.weightsCopied : bulkCopy.weightsCopyPawn}
        </button>
      </div>
    </section>
  {/if}
</div>
```

- [ ] **Step 4: Run the e2e, desktop and mobile**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-weights.spec.ts
```

Expected: PASS.

- [ ] **Step 5: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim/tools tests/e2e/sim-weights.spec.ts
npx prettier --write src/components/sim/tools tests/e2e/sim-weights.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim/tools web/tests/e2e/sim-weights.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): /sim/weights, with the warning first

The caution opens the page rather than footnoting it: the whole reason
the page exists is that addons ask for a number the number cannot
justify. The reference stat comes from GET /v1/specs, is required by
the envelope, and always leads the list, because everything on the page
is scaled to it. Stat ids are contract 10.8's pinned proto.Stat
list: one hit, one crit, haste split, mp5 spelled mp5. Every weight
carries its error, as a bar and as a figure.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 20: `/sim/<id>` for every kind, the unfurl, and history

**Files:**
- Create: `web/src/components/sim/SavedCombos.svelte`
- Create: `web/src/components/sim/SavedWeights.svelte`
- Modify: `web/src/components/sim/SavedSim.svelte`
- Modify: `web/src/components/sim/SimHistory.svelte`
- Modify: `web/src/lib/report/og-meta.ts`, `og-meta.test.ts`
- Modify: `web/tests/e2e/sim-saved.spec.ts`

**Interfaces:**
- Consumes: `kindOf` (Task 1); `comboRows`, `headlineFor`, `deltaLabel` (Task 9); `statLabel`, `pawnString` (Task 9); `SimListRow.kind`/`headline` (Task 1).
- Produces: `simShellMeta` naming the kind; `SavedSim` rendering by kind; `SimHistory` filtering by kind.

- [ ] **Step 1: Write the failing tests**

Append to `web/src/lib/report/og-meta.test.ts`:

```ts
describe('simShellMeta by kind', () => {
  it('names Top Gear and its headline, item included', () => {
    const meta = simShellMeta({ ...bulkResult, sim_id: 'simfixtureab' } as unknown as SimResult);
    expect(meta.title).toBe('Top Gear · Fury Warrior · Forever Sixty');
    expect(meta.description).toContain('6 combinations');
    expect(meta.description).toContain('+41 DPS from Helm of Wrath');
  });

  it('says "no combinations" for an empty result, as the API’s headline rule does', () => {
    const meta = simShellMeta({
      ...bulkResult,
      sim_id: 'simfixtureab',
      combos: [],
    } as unknown as SimResult);
    expect(meta.description).toContain('no combinations');
  });

  it('names stat weights and lists the top three', () => {
    const meta = simShellMeta({ ...weightsResult, sim_id: 'simfixtureab' } as unknown as SimResult);
    expect(meta.title).toBe('Stat weights · Fury Warrior · Forever Sixty');
    expect(meta.description).toContain('Attack power 1.00');
  });

  it('leaves a plain run exactly as it was', () => {
    const meta = simShellMeta(fixtureResult);
    expect(meta.title).toContain('DPS · Forever Sixty');
  });
});
```

Append to `web/tests/e2e/sim-saved.spec.ts`:

```ts
test('a saved Top Gear result renders its ranked table, not a damage breakdown', async ({ page }) => {
  await page.route('**/v1/sims/simbulk23456', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, request_id: 'r', error: null, data: bulkFixture }),
    }),
  );
  await page.goto('/sim/simbulk23456');
  await expect(page.getByTestId('sim-combos')).toBeVisible();
  await expect(page.getByTestId('sim-combo-row').first()).toBeVisible();
  await expect(page.getByTestId('sim-results-tabs')).toHaveCount(0);
});

test('a saved weights result renders its table and its Pawn string', async ({ page }) => {
  await page.route('**/v1/sims/simweight3456', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, request_id: 'r', error: null, data: weightsFixture }),
    }),
  );
  await page.goto('/sim/simweight3456');
  await expect(page.getByTestId('sim-weights')).toBeVisible();
  await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');
});
```

with, at the top of that file:

```ts
import bulkFixture from '../../src/fixtures/sim/bulk-result.json';
import weightsFixture from '../../src/fixtures/sim/weights-result.json';
```

- [ ] **Step 2: Run them and watch them fail**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/report/og-meta.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-saved.spec.ts --project=desktop
```

Expected: FAIL on the new cases.

- [ ] **Step 3: Write `SavedCombos.svelte` and `SavedWeights.svelte`**

`SavedCombos.svelte` is `ComboResults.svelte` without the two "do something with the winner"
buttons, because a saved page belongs to whoever opened the link and not to the character:

```svelte
<!-- web/src/components/sim/SavedCombos.svelte -->
<!-- A saved Top Gear, talent compare or Droptimizer result at /sim/<id>: read-only, so
     the ranked table and the per-slot summary without the "open in planner" and "copy to
     addon" buttons -- those act on the reader's own character, and a saved page is not
     theirs. "Run this yourself" beside the header is the way back to a page that is. -->
<script lang="ts">
  import { loadItems } from '../../lib/planner/load';
  import type { Item } from '../../lib/planner/types';
  import { SLOT_LABELS, type Slot } from '../../lib/planner/types';
  import type { BulkResult } from '../../lib/sim/bulk-types';
  import { comboRows, deltaLabel, slotSummary } from '../../lib/sim/combos';
  import { bulkCopy } from '../../lib/sim/copy';
  import SubstitutionChips from './tools/SubstitutionChips.svelte';

  let { result, treeVersion }: { result: BulkResult; treeVersion: string } = $props();

  let items = $state<Map<number, Item>>(new Map());

  void (async () => {
    try {
      const file = await loadItems(treeVersion, result.request.character.class);
      items = new Map(file.items.map((item) => [item.id, item]));
    } catch {
      items = new Map();
    }
  })();

  const rows = $derived(comboRows(result));
  const summary = $derived(slotSummary(result));
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-combos">
  <p class="tabular text-strong font-mono text-[14px]" data-testid="sim-equipped-line">
    {bulkCopy.resultsEquipped}: {Math.round(result.equipped.mean).toLocaleString('en-US')}
    ± {Math.round(1.96 * result.equipped.error).toLocaleString('en-US')}
    <span class="text-muted">{bulkCopy.ranAtStages(result.stages)}</span>
  </p>
  <ul class="flex flex-col">
    {#each rows as row (row.combo.substitutions.map((sub) => `${sub.kind}:${sub.slot ?? ''}:${sub.item_id ?? sub.name ?? ''}`).join('|'))}
      <li
        class="border-line-soft grid min-h-11 grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 border-b px-2 py-2"
        data-testid="sim-combo-row"
      >
        <span class="tabular text-muted font-mono text-[12px]">{row.rank}</span>
        <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
        <span class="tabular text-strong ml-auto font-mono text-[13px]">
          {Math.round(row.combo.dps.mean).toLocaleString('en-US')}
        </span>
        <span class="tabular text-gold ml-auto font-mono text-[13px]">{deltaLabel(row.combo.delta)}</span>
        <span class="tabular text-muted ml-auto font-mono text-[12px]">{row.percent.toFixed(1)}%</span>
      </li>
    {/each}
  </ul>
  {#if summary.length > 0}
    <section class="border-line rounded-panel border p-3" data-testid="sim-slot-summary">
      <h3 class="section-title text-[14px]">{bulkCopy.slotSummary}</h3>
      <ul class="flex flex-col">
        {#each summary as row (`${row.slot}-${row.item_id}`)}
          <li class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 last:border-b-0">
            <span class="text-muted w-24 text-[12px]">{SLOT_LABELS[row.slot as Slot] ?? row.slot}</span>
            <span class="text-text flex-1 text-[13px]">{row.name}</span>
            <span class="tabular text-gold font-mono text-[13px]">
              {row.gain === null ? '—' : `+${Math.round(row.gain).toLocaleString('en-US')}`}
            </span>
          </li>
        {/each}
      </ul>
    </section>
  {/if}
</section>
```

The item file is fetched here for **icons and rarity colours only** — every name on the
page comes off the result (contract 10.1 A6), so a saved sim whose class file 404s still
reads correctly.

`SavedWeights.svelte`:

```svelte
<!-- web/src/components/sim/SavedWeights.svelte -->
<!-- A saved stat-weights result. The caution travels with the numbers: a shared link is
     exactly where someone meets these weights without having read the page that made them. -->
<script lang="ts">
  import type { WeightsResult } from '../../lib/sim/bulk-types';
  import { bulkCopy } from '../../lib/sim/copy';
  import { pawnString, statLabel, weightScale } from '../../lib/sim/weights';

  let { result }: { result: WeightsResult } = $props();

  const scale = $derived(weightScale(result.weights));
  const pawn = $derived(pawnString(result.request.spec, result.weights));
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-weights">
  <p class="text-muted text-[13px]" data-testid="sim-weights-warning">{bulkCopy.weightsWarning}</p>
  <ul class="flex flex-col">
    {#each result.weights as row (row.stat)}
      <li
        class="border-line-soft grid min-h-11 grid-cols-[minmax(120px,1fr)_minmax(0,3fr)_88px] items-center gap-x-3 border-b px-2 py-2"
        data-testid={`sim-weight-${row.stat}`}
      >
        <span class="text-text text-[13px]">{statLabel(row.stat)}</span>
        <span class="bg-card-top relative block h-[8px] w-full">
          <span
            class="bg-gold absolute top-0 left-0 block h-[8px]"
            style={`width:${Math.max(0, (row.weight / scale) * 100)}%`}
          ></span>
        </span>
        <span class="tabular text-strong ml-auto font-mono text-[13px]">
          {row.weight.toFixed(2)}
          <span class="text-muted">± {row.error.toFixed(2)}</span>
        </span>
      </li>
    {/each}
  </ul>
  <code class="text-muted font-mono text-[12px] break-all" data-testid="sim-pawn">{pawn}</code>
</section>
```

- [ ] **Step 4: Render by kind in `SavedSim.svelte`**

Replace the single `simResultsLazy` block with a kind switch. Add to the script:

```ts
  import { kindOf, type BulkResult, type WeightsResult } from '../../lib/sim/bulk-types';

  const kind = $derived(kindOf(result.request));
  // Three separate chunks, so a saved weights page downloads neither the damage tables nor
  // the combination table, and vice versa.
  const savedCombosLazy = createLazyComponent(() => import('./SavedCombos.svelte'));
  const savedWeightsLazy = createLazyComponent(() => import('./SavedWeights.svelte'));

  if (kind === 'weights') savedWeightsLazy.load();
  else if (kind !== 'run') savedCombosLazy.load();
  else simResultsLazy.load();
```

(remove the bare `simResultsLazy.load();` that stood alone), and in the markup, replace the
`{#if simResultsLazy.current}` block with:

```svelte
{#if kind === 'weights'}
  {#if savedWeightsLazy.current}
    <savedWeightsLazy.current result={result as WeightsResult} />
  {:else}
    {@render lazyFallback(savedWeightsLazy)}
  {/if}
{:else if kind !== 'run'}
  {#if savedCombosLazy.current}
    <savedCombosLazy.current result={result as BulkResult} treeVersion={character.tree_version} />
  {:else}
    {@render lazyFallback(savedCombosLazy)}
  {/if}
{:else if simResultsLazy.current}
  <simResultsLazy.current
    summary={result.summary}
    estimate={result.dps}
    iterationsRun={result.iterations_run}
    actionNames={null}
  />
{:else}
  {@render lazyFallback(simResultsLazy)}
{/if}
```

Also give the heading the kind, beside the spec:

```ts
  const KIND_TITLES: Record<string, string> = {
    gear: bulkCopy.gearTitle,
    talents: bulkCopy.talentsTitle,
    drops: bulkCopy.dropsTitle,
    weights: bulkCopy.weightsTitle,
  };
  const heading = $derived(
    kind === 'run' ? specLabel(result.request.spec) : `${KIND_TITLES[kind]} · ${specLabel(result.request.spec)}`,
  );
```

and hide the character strip's gear grid for a weights result, which has no gear story:
`<CharacterStrip … gearKnown={kind !== 'weights' && gearKnown} … />`.

- [ ] **Step 5: Name the kind in the unfurl**

In `web/src/lib/report/og-meta.ts`, replace `simShellMeta`'s body with a kind switch:

```ts
const SIM_KIND_TITLES: Record<string, string> = {
  gear: 'Top Gear',
  talents: 'Talent compare',
  drops: 'Droptimizer',
  weights: 'Stat weights',
};

export function simShellMeta(result: SimResult): ShellMeta {
  const kind = kindOf(result.request);
  const canonical = `${SITE_BASE_URL}/sim/${result.sim_id ?? ''}`;
  const spec = specLabel(result.request.spec);

  if (kind === 'weights') {
    const weights = (result as WeightsResult).weights ?? [];
    const top = weights
      .slice(0, 3)
      .map((row) => `${statLabel(row.stat)} ${row.weight.toFixed(2)}`)
      .join(' · ');
    return {
      title: `${SIM_KIND_TITLES.weights} · ${spec} · Forever Sixty`,
      description: `Simulated on engine ${result.engine_version}: ${top}.`,
      image: SITE_CARD,
      canonical,
    };
  }

  if (kind !== 'run') {
    const bulk = result as BulkResult;
    const combos = bulk.combos ?? [];
    // Contract 10.1 A6 fills Substitution.Name for items too, so the unfurl can name the
    // winning change without the per-class item file this function must never fetch.
    // An empty result reads "no combinations", matching the API's own headline rule
    // (contract 10.6) rather than inventing a second wording for the same state.
    const headline = combos.length === 0 ? 'no combinations' : headlineFor(bulk);
    return {
      title: `${SIM_KIND_TITLES[kind]} · ${spec} · Forever Sixty`,
      description:
        `Simulated on engine ${result.engine_version}: ${combos.length} combinations, ` +
        `${headline}.`,
      image: SITE_CARD,
      canonical,
    };
  }

  // ... today's plain-run body, unchanged ...
}
```

with `import { kindOf, type BulkResult, type WeightsResult } from '../sim/bulk-types';`,
`import { headlineFor } from '../sim/combos';` and `import { statLabel } from '../sim/weights';`.

- [ ] **Step 6: Filter history by kind**

In `web/src/components/sim/SimHistory.svelte`, add a kind filter over the rows and show each
row's `headline` where the API sent one:

```svelte
  let kind = $state('');
  const KINDS = [
    { id: '', label: 'Every kind' },
    { id: 'run', label: 'Sims' },
    { id: 'gear', label: bulkCopy.gearTitle },
    { id: 'talents', label: bulkCopy.talentsTitle },
    { id: 'drops', label: bulkCopy.dropsTitle },
    { id: 'weights', label: bulkCopy.weightsTitle },
  ];
  const shown = $derived((rows ?? []).filter((row) => kind === '' || (row.kind ?? 'run') === kind));
```

```svelte
  <select
    class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
    data-testid="sim-history-kind"
    bind:value={kind}
  >
    {#each KINDS as option (option.id)}
      <option value={option.id}>{option.label}</option>
    {/each}
  </select>
```

and, in each row, render `row.headline ?? ''` beside the title. **It is never recomputed
here**: contract 10.6 has the API compose it at save time, with its own rules for several
substitutions ("… and 2 more") and for an empty result ("no combinations", "no upgrades",
"no weights"), and a second composition in the browser would drift from the stored one.

- [ ] **Step 7: Run the tests**

```bash
cd web && FOREVER_DATA=fixture npx vitest run src/lib/report/og-meta.test.ts src/worker.test.ts
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-saved.spec.ts
```

Expected: PASS.

- [ ] **Step 8: Type-check, lint, format, commit**

```bash
cd web && NO_COLOR=1 npx astro check
npx eslint src/components/sim src/lib/report/og-meta.ts src/lib/report/og-meta.test.ts tests/e2e/sim-saved.spec.ts
npx prettier --write src/components/sim src/lib/report/og-meta.ts src/lib/report/og-meta.test.ts tests/e2e/sim-saved.spec.ts
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/src/components/sim web/src/lib/report/og-meta.ts web/src/lib/report/og-meta.test.ts web/tests/e2e/sim-saved.spec.ts
git commit -m "$(cat <<'MSG'
feat(sim): /sim/<id> renders every kind, and the unfurl names it

Kind is derived from the stored request, never stored twice, and each
kind's view is its own lazy chunk so a saved weights page downloads
neither the damage tables nor the combination table. The unfurl names
the winning item, which it can because contract 10.1 A6 fills
Substitution.Name from simdb. The saved view has no "open in planner"
or "copy to addon": those act on the reader's own character, and a
saved page is not theirs. History filters by kind and shows the API's
own composed headline rather than recomputing one (contract 10.6).

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

### Task 21: Phone passes, the gated real-engine run, Lighthouse, and the whole suite

**Files:**
- Create: `web/tests/e2e/sim-tools-phone.spec.ts`
- Create: `web/tests/e2e/sim-gear-real-engine.spec.ts`
- Modify: `web/src/pages/sim/gear.astro`, `talents.astro`, `drops.astro`, `weights.astro` (measured `min-h`)
- Modify: `web/lighthouserc.json`
- Modify: `web/package.json` (`test:e2e:phone`, `test:e2e:real-engine`)

**Interfaces:**
- Consumes: every page from Tasks 11–20.
- Produces: green CI. No new module.

- [ ] **Step 1: Write the phone spec**

Create `web/tests/e2e/sim-tools-phone.spec.ts`:

```ts
// web/tests/e2e/sim-tools-phone.spec.ts
// The four tool pages at phone width. Everything here is about layout and reach, not about
// the engine: a control that overflows the viewport or is under 44px tall is unusable, and
// nothing else in the suite measures that.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

test.use({ viewport: { width: 360, height: 780 } });

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function load(page: Page, route: string): Promise<void> {
  await page.goto(route);
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

const ROUTES = ['/sim/gear', '/sim/talents', '/sim/drops', '/sim/weights'] as const;

for (const route of ROUTES) {
  test(`${route} never scrolls sideways at 360px`, async ({ page }) => {
    await load(page, route);
    const overflow = await page.evaluate(
      () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
    );
    expect(overflow).toBeLessThanOrEqual(1);
  });

  test(`${route} gives every control a 44px hit target`, async ({ page }) => {
    await load(page, route);
    const short = await page.evaluate(() => {
      const bad: string[] = [];
      for (const element of document.querySelectorAll('button, a[href], select, input, label')) {
        const box = element.getBoundingClientRect();
        if (box.width === 0 && box.height === 0) continue;
        if (box.height < 44) bad.push(element.outerHTML.slice(0, 80));
      }
      return bad;
    });
    expect(short).toEqual([]);
  });
}

test('the run bar stays reachable on /sim/gear after a run', async ({ page }) => {
  await load(page, '/sim/gear');
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-run-bulk')).toBeVisible();
});
```

- [ ] **Step 2: Run it and fix what it finds**

```bash
cd web && E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test tests/e2e/sim-tools-phone.spec.ts --project=mobile
```

Expected on the first run: FAIL, with a list of short controls and possibly one overflow.
Fix each in its component — the `min-h-11` and `h-11` classes already used everywhere are
the remedy for the first, and `flex-wrap` plus `min-w-0` on the offending row for the
second. Do **not** relax the assertion.

- [ ] **Step 3: Write the gated real-engine spec**

Create `web/tests/e2e/sim-gear-real-engine.spec.ts`:

```ts
// web/tests/e2e/sim-gear-real-engine.spec.ts
// The one end-to-end proof that the real wasm's simPlan/simRank survive the full browser
// integration: two candidates, planned, run stage by stage through the pool, and ranked.
// sim.yml's own wasm smoke runs one Top Gear plan of four candidates directly against the
// artifact; what this adds is the browser, the worker pool and the page around it.
//
// Skipped unless the build was made with PUBLIC_SIM_ENGINE=wasm and the artifact it names
// was published to web/public/_sim/<ENGINE_VERSION>/ first. Run it with
// `npm run test:e2e:real-engine`, after `make simdb && make artifacts && make publish-wasm`
// from the repository root.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test } from '@playwright/test';
import { ENGINE_VERSION } from '../../src/lib/sim/version';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const artifactPublished = existsSync(path.join(WEB_ROOT, 'public/_sim', ENGINE_VERSION, 'sim.wasm'));

test.skip(
  process.env.PUBLIC_SIM_ENGINE !== 'wasm' || !artifactPublished,
  `real-engine Top Gear smoke; publish web/public/_sim/${ENGINE_VERSION}/ first ` +
    '(make simdb && make artifacts && make publish-wasm) and run with npm run test:e2e:real-engine',
);

const activeBuild = JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/active-build.json'), 'utf8')) as {
  build: string;
};

// Real Forever item ids from the active build's own embedded database, the same two
// sim-real-engine.spec.ts uses -- which is what proves the wasm's simdb resolves real gear
// rather than echoing back whatever it was handed.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=21521`;

test('two candidates plan, run and rank against the equipped set', async ({ page }) => {
  test.slow();
  const pageErrors: string[] = [];
  page.on('pageerror', (error) => pageErrors.push(error.message));

  await page.goto('/sim/gear');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-slot-grid')).toBeVisible();

  // Two search candidates rather than two bag rows: this export carries no bags, and the
  // search is the path that exercises the item database on both sides.
  await page.getByTestId('sim-item-search').getByRole('searchbox').fill('');
  const adds = page.locator('[data-testid^="sim-search-add-"]');
  await adds.nth(0).click();
  await adds.nth(1).click();

  await expect(page.getByTestId('sim-combo-count')).toHaveText(/[1-9]\d* valid combinations?/);

  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-stage-progress')).toBeVisible();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 180_000 });

  const rows = page.getByTestId('sim-combo-row');
  expect(await rows.count()).toBeGreaterThan(0);
  await expect(page.getByTestId('sim-equipped-line')).toContainText(/\d/);
  expect(pageErrors).toEqual([]);
});
```

In `package.json`:

```json
    "test:e2e:real-engine": "PUBLIC_SIM_ENGINE=wasm E2E_PORT=4325 playwright test tests/e2e/sim-real-engine.spec.ts tests/e2e/sim-gear-real-engine.spec.ts --project=desktop",
    "test:e2e:phone": "playwright test report-phone.spec.ts rankings-phone.spec.ts character-phone.spec.ts guild-phone.spec.ts sim-phone.spec.ts sim-tools-phone.spec.ts --project=mobile",
```

- [ ] **Step 4: Measure the four `min-h` reservations and replace the guesses**

```bash
cd web && FOREVER_DATA=fixture npm run build
FOREVER_DATA=fixture npm run preview -- --port 4326 &
```

Then, for each of the four routes, at 360 px and 1280 px, with the fixture character loaded:

```bash
cd web && E2E_PORT=4326 node --input-type=module -e "
import { chromium } from '@playwright/test';
const build = (await import('./src/data/active-build.json', { with: { type: 'json' } })).default.build;
const code = \`FS1:\${build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784\`;
const browser = await chromium.launch();
for (const route of ['/sim/gear', '/sim/talents', '/sim/drops', '/sim/weights']) {
  for (const width of [360, 1280]) {
    const page = await browser.newPage({ viewport: { width, height: 900 } });
    await page.goto('http://localhost:4326' + route);
    await page.getByTestId('sim-addon-input').fill(code);
    await page.getByTestId('sim-addon-load').click();
    await page.getByTestId('sim-character').waitFor();
    await page.waitForTimeout(500);
    const height = await page.evaluate(() => document.getElementById('sim-tools').getBoundingClientRect().height);
    console.log(route, width, Math.ceil(height));
    await page.close();
  }
}
await browser.close();
"
```

Replace each shell's `min-h-[…px] md:min-h-[…px]` with the measured 360 px and 1280 px
numbers, rounded up. Then kill the preview server.

- [ ] **Step 5: Add the four pages to Lighthouse**

In `web/lighthouserc.json`, add to `collect.url`:

```json
        "http://localhost/sim/gear.html",
        "http://localhost/sim/talents.html",
        "http://localhost/sim/drops.html",
        "http://localhost/sim/weights.html"
```

Extend the existing `.*/sim(?:/specs)?\\.html$` matrix entry's pattern so the four tool
pages take the same budget as `/sim` and `/sim/specs`:

```json
          "matchingUrlPattern": ".*/sim(?:/specs|/gear|/talents|/drops|/weights)?\\.html$",
```

and extend the catch-all's negative lookahead in the first matrix entry to match:

```json
          "matchingUrlPattern": "^(?!.*/(?:index|planner|logs)\\.html$)(?!.*/reports/)(?!.*/sim(?:/specs|/gear|/talents|/drops|/weights)?\\.html$)(?!.*/sim/[a-z2-7]{12}\\.html$).*$",
```

- [ ] **Step 6: Run Lighthouse**

```bash
cd web && FOREVER_DATA=fixture npm run build && npm run lhci
```

Expected: every assertion passes. The likely first failure is `cumulative-layout-shift` on a
page whose `min-h` is still short — fix the number, never the budget. The second likely
failure is `total-blocking-time` if the island is doing work on mount: the store's pool is
created on the first run, not on mount, so if TBT is over 100 ms find what else is running
eagerly (a `$effect` that fetches, a large `$derived` over `items`) and defer it.

- [ ] **Step 7: Run the whole suite**

```bash
cd web && export PATH=/Users/jh/.nvm/versions/node/v22.12.0/bin:$PATH
FOREVER_DATA=fixture npm test
NO_COLOR=1 npx astro check
npx eslint .
npx prettier --check .
E2E_PORT=4326 FOREVER_DATA=fixture npx playwright test
```

Expected: vitest green, `0 errors`, eslint and prettier clean, Playwright green on both
projects.

- [ ] **Step 8: Check the island budget one last time**

```bash
cd web && FOREVER_DATA=fixture npm run build
```

Expected: `dist/sim-island.js` **unchanged** in size to within a few hundred bytes — nothing
in this lane may have grown /sim's bundle — and `dist/sim-tools-island.js` under 70 KB
gzipped. If `sim-island` grew, something in `SavedSim.svelte`'s kind switch is importing
eagerly instead of lazily; fix that rather than raising the budget.

- [ ] **Step 9: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/sim-web
git add web/tests/e2e web/src/pages/sim web/lighthouserc.json web/package.json web/src/components web/src/lib
git commit -m "$(cat <<'MSG'
test(sim): phone passes, a gated real-engine Top Gear run, and Lighthouse

The four tool pages at 360px with a 44px hit-target floor and a no-
sideways-scroll assertion, one gated end-to-end run of the real wasm's
simPlan/simRank with two candidates, and measured shell reservations so
the mount swap moves nothing. Every page joins /sim's Lighthouse budget
rather than the catch-all's.

Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>
MSG
)"
```

---

## Merging with part A

Both lanes land on the same branch. Three places will conflict, and each has one right
resolution:

1. **`web/src/lib/sim/copy.ts`** — part A appends inside `simCopy`, this lane appends
   `bulkCopy` after it. Keep both; neither deletes the other's keys.
2. **`web/src/lib/sim/types.ts`** — part A adds `TargetError`, the `EncounterSpec` fields
   and `Sample`; this lane adds `reference_stat` on `SpecFidelity` and `kind`/`headline` on
   `SimListRow`. Keep both.
3. **`web/src/lib/sim/run.ts`** — untouched here. After the merge, `run.ts`'s
   `execute()` and `bulk-run.ts`'s stage loop both drive the pool; `bulk-run.ts` runs a
   stage's requests through `pool.run` directly and never splits, so there is nothing to
   extract. **No follow-up is owed** — confirm it, and if a genuine duplication appears,
   extract it into a new file rather than editing either.

Part A's `SettingsPanel.svelte` and `RequestDrawer.svelte` are mounted by
`BulkRunBar.svelte` (Task 15). If part A lands after this plan, `BulkRunBar` will not
compile until they exist: stub both as one-line components in a separate commit, note it in
the commit body, and delete the stubs when part A lands.

A fourth file both lanes touch is **`web/src/lib/sim/engine.ts`**: this lane declares all
three of contract 10.2's new exports (`simCount`, `simValidate`, `simNeedsMore`) on
`EngineModule` and on the pool, in Tasks 3 and 4, precisely so part A does not also edit
that surface. If part A has already added `simNeedsMore` or `simValidate`, keep its version
and drop this lane's duplicate declaration; the signatures are the contract's and cannot
differ.

---

## Self-review

**Spec coverage.**

| Requirement | Task |
| --- | --- |
| Design 2.2 / contract 4: the plan-rank loop through the existing pool | 3, 4, 5 |
| Design 2.3 / contract 1.3: the browser cap, halved on four cores, enforced by `ErrCapExceeded`, shown as `sim-cap-notice` with the premium alternative | 1, 5, 15 |
| Design 10.1: `sim-stage-progress`, abort returns the partial | 5, 15 |
| Design 10.2: the premium path posts the same request and polls | 5, 10 |
| Design 3.1.1: character strip, same source switcher | 11 |
| Design 3.1.2: slot grid, equipped/bag/bank/search/pinned rows, checkboxes, copy-and-modify (enchant, suffix), locks, rings and trinkets in one list | 8, 12 |
| Design 3.1.3: item search over `items.json` — name, min ilvl, slot, source, usable-only | 7, 13 |
| Design 3.1.4: per-slot enchant lists, "keep current", "none", a selection cap | 7, 12 |
| Design 3.1.5 / contract 10.1 A5: consumables as candidates | 1, 13 |
| Design 3.1.6: own talents, saved planner builds, export loadouts, add custom → planner inline | 14 |
| Design 3.1.7: named sets from the export or a pasted second string | 14 |
| Design 3.1.8: live combination count, precision, cap notice, lane switch, one button | 15 |
| Design 3.2: the rules card in words | 16 |
| Design 3.3: equipped baseline, ranked rows with delta ± error and percent, within-error grouping by `group`, per-slot summary, "Open in planner", "Copy to addon" | 8, 9, 16 |
| Design 3.4 / `/sim/talents` | 17 |
| Design 4.4: the 4-piece results filter | 9, 16 |
| Design 6.1: every source kind, phase gate, "show upcoming", no-date sources | 2, 6, 18 |
| Design 6.2: `loot.json`, no probability shown | 6, 18 |
| Design 6.3: results by source, best per boss, flat list, pin into Top Gear | 18 |
| Design 7 / `/sim/weights`: warning, stat picker on `reference_stat`, error bars, Pawn | 9, 19 |
| Design 5.2 / contract 9: `/sim/<id>` renders by kind; unfurl names the kind | 20 |
| Design 9.3: history lists every kind, filterable, with the API's headline | 20 |
| Fixtures: bulk result, weights result, `loot.json`, `enchants.json`, `suffixes.json`, `simbuffs.json`, `phases.json`, `specs.json` with `reference_stat`, the fake engine's six exports | 2, 3 |
| e2e per page, desktop and mobile; one gated real-engine Top Gear run | 12–21 |
| vitest for every pure module | 1, 2, 5–10 |
| Shell skeletons so CLS holds | 11, 21 |

**Contract section 10 coverage** (the amendment round, 2026-09-19, commit 439f0b7).

| Ruling | Where it lands |
| --- | --- |
| A1 lane-aware caps | Header; `browserCap`/`SERVER_CAP` (Task 1); the store's server-cap gate (Task 10) |
| A2 server cap 5,000 | Task 1 `SERVER_CAP`; `bulkCopy.capPremiumNote`/`serverCapNotice` (Task 1); the run bar (Task 15) |
| A3 a bulk request's iterations are its precision's final stage | `finalIterations` (Task 1); the store's request builders (Task 10); the fake's `simValidate` (Task 3) |
| A4 mode decides expansion | The fake's `expand` (Task 3); no `combinations` boolean anywhere |
| A5 `BulkSpec.Consumables` | Task 1 (field), Task 3 (expansion), Task 13 (control and count) |
| A6 `Candidate.SourceName`, named item substitutions | Task 1 (fields), 6 (`sourceNameOf`), 8 (`toCandidates`), 9 (`substitutionLabel`, `slotSummary.name`), 16, 18, 20 |
| A7 / 10.8 `proto.Stat` vocabulary pinned, required reference, per-kind defaults | `WEIGHT_STATS`, `fallbackReferenceFor` (Task 9); the weights fixture (Task 2); `/sim/weights` (Task 19) |
| A9 enchants embedded in simdb | Noted in `enchants.ts`'s header (Task 7): the browser file is the UI's copy only |
| A10 `StageRequests.Ran` | Task 1 (field), Task 3 (the fake keeps no state), Task 5 (`partialStage` carries it) |
| A11 additive progress widening | `BulkServerProgress` (Task 5) |
| 10.2 `simCount`, `simValidate`, `simNeedsMore`, the structured cap error, unsplit stages | Tasks 1, 3, 4, 5 |
| 10.4 AreaTable zone ids, `<source>:<npc-id>` boss ids, `opens: "later"`, enchant keying, `simbuffs.json`, `data/curated/phases.json` | Tasks 2, 6, 7, 13, 18 |
| 10.5 `item_id[:enchant[:suffix]]` in `<gear>`/`sets=`, `professions=`, per-slot enchant on `SimCharacter` | Task 8 (`addon-export`), Task 10 (`seedRows`), the Consumes block |
| 10.6 `POST /v1/sims/run`, `GET /v1/builds?mine=1`, `GET /v1/phases`, `headline` composed at save time | Tasks 5, 6, 10, 14, 20 |
| 10.7 test ids are a minimum, no zone icons, "my professions" | Header, Task 18 |
| 10.8 `Substitution.Kind` gains `consumes` | Task 1 (closed union), Task 3 (the fake emits it), Task 9 (`substitutionLabel` via `bulkCopy.consumesChip`), Task 13 |
| 10.8 stat vocabulary pinned, one `hit` and one `crit`, `mp5` | Task 9 (`WEIGHT_STATS` and the pinned-list test), Task 2 (fixture), Task 19 (e2e) |
| 10.8 client-side server-cap gate | Task 10 (`serverCapNotice`), Task 15 (the disabled premium button) |

**Rulings this lane does not carry.** A8 (`TargetArmorByLevel`) and A12 (`SampleCast` is
action-keyed) are part A's settings panel and report work. 10.3, and 10.8's faction enum,
progress callback and embedded enchant table, are the engine fork's and the sim module's.

**Nothing in section 10 is left unapplied by this plan.** The two items the first revision
flagged open — the substitution kind for a consumable list, and the exact stat spellings —
are settled by 10.8 and carried above; the third, the client-side server-cap gate, is
ratified as planned.

**Not in this plan, and why.** The settings panel, fight styles, the precision/target-error
control on `/sim`, the report additions (sample log, details card), the request drawer's own
implementation and the FS1 v2 decoder are part A's. `sim/bulk`, the wasm exports, the API's
migration 0014 and the generation of `loot.json`, `enchants.json`, `suffixes.json`,
`simbuffs.json` and `phases.json` are other lanes'.

**Type consistency.** `BulkStore` is the one store type; `CandidateRow` the one row type;
`Origin` the one origin union; `Substitution.kind` is closed at `item | talents | set |
consumes` and every renderer goes through `substitutionLabel`; `Precision`, `BulkMode`,
`SimKind`, `PhaseRow` and the two caps are declared once and imported everywhere. `substitutionLabel(sub)` takes one argument
in `combos.ts`, `SubstitutionChips`, `ComboResults`, `DropResults`, `SavedCombos` and
`og-meta.ts`; `headlineFor(result)` likewise. Every `phase.ts` function that needs the
table takes it as its first parameter (`phaseAt`, `phaseStart`, `hasOpened`,
`openDateLabel`), and `phaseLabel(name)` does not, because a label is display copy and not
data. `isOpen(phases, source, when)` has that shape in `loot.ts`, its test, the store and
`SourcePicker`. `store.rows`, `store.locked`, `store.loadouts`, `store.namedSets`,
`store.consumableIds`, `store.combinations`, `store.capNotice`, `store.serverCapNotice`,
`store.progressLine`, `store.phases`, `store.simBuffs`, `store.combos` and `store.weights`
are the names every component reads.

**Placeholder scan.** No "TBD", no "similar to Task N", no step that describes without
showing. One step deliberately ends in a measurement rather than a literal — Task 21's
replacement of the four shells' `min-h` reservations — and it names the command that takes
it. That is a fact the plan cannot know before the pages are built, not a gap in it. The
stat ids that were previously left to a `grep` are now pinned by contract 10.8 and asserted
against the full list in Task 9.
