# The current character (Lane B) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development
> (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps
> use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Give the site a single "current character" the whole product agrees on: a
localStorage-remembered pointer, a visible chip, and the hand-off links (Quick Sim →
planner, Droptimizer/Top Gear rows → planner, paste boxes → `/addon`, Share → confirm) that
close the dead ends the architect's review found.

**Architecture:** A shared pointer type and three pure functions
(`web/src/lib/current-character.ts`) that every consumer reads/writes through, built on the
`{ source, ref }` shape the sim tab strip already uses. `sources.ts`'s loaders write the
pointer on every successful load (one seam, reused by both the sim store and the tools
island). A new chip component renders it. Existing "Open in planner" plumbing
(`codeForCharacterSpec`, `plannerHrefFor` in `sim/character.ts`) is reused wherever a full
character is already in memory; the new module's own `plannerHrefFor`/`simHrefFor` are the
fallback for a page that only has the bare pointer.

**Tech Stack:** Astro (static pages) + Svelte 5 islands (runes: `$state`/`$derived`/`$effect`),
TypeScript, Vitest (`svelte/server` `render()` for component tests), Playwright (e2e).

**Spec:** `docs/superpowers/specs/2026-09-21-one-product-design.md`, section 1 (this lane's
scope in full) and section 0 (binding on every lane). Evidence:
`.superpowers/journeys/architect/review.md` sections 2-4.

## Global Constraints

(Spec section 0, binding on every task below.)

- Design system: `design/DESIGN-SYSTEM.md`. No emoji, no marketing buttons, secondary
  buttons only, restrained motion, 44px hit targets on phone, dark only.
- Honest copy: never promise what is not built; say what a page is for and who it is for.
  No exclamation marks.
- No layout shift: anything that appears after hydration reserves its space or sits below
  existing content. Lighthouse budgets in `web/lighthouserc.json` must hold.
- Signed-out first: every feature works without an account; an account only adds.
- localStorage is a convenience, never the source of truth: wrap every read and write in
  try/catch, render correctly without it, never read it during SSR/prerender.
- Tests: unit tests for every pure function, component tests where the repo has them
  (`svelte/server`'s `render()`, see `StatWeights.test.ts`), e2e for each new hand-off
  (desktop and mobile Playwright projects). `FOREVER_DATA=fixture npm run sync` before web
  tests. Node 22.12 via nvm.
- File ownership (this lane): `web/src/components/sim/**`, `web/src/components/planner/**`,
  `web/src/lib/sim/**`, `web/src/lib/planner/**`, `web/src/lib/current-character.ts`,
  `web/src/components/CurrentCharacterChip.svelte`, the `/sim*` and `/planner` page files,
  and their tests. Never edit `Header.astro`, `Footer.astro`, `Base.astro`, `index.astro`,
  `addon.astro`, `account.astro`, `logs.astro`, anything under `web/src/components/report/`,
  rankings, or `api/`.
- Code: functions under 50 lines, files under 800 (several copy/store files are already at
  the ceiling: `src/lib/sim/store.svelte.ts` 798, `src/lib/sim/bulk-store.svelte.ts` 839,
  `src/lib/sim/copy.ts` 1247 — put new copy in a new small module, add to these files only
  the minimum lines a task strictly needs). No magic numbers, no dead code, immutable
  updates, every visible string in a copy module.
- Never save a public build or sim, never upload, never sign in on production. Tests use the
  fake engine and route stubs, as the existing specs do.
- Commits: message via `printf` to a file under `.superpowers/`, then
  `git commit -F <file>` alone. Conventional subjects. End every message with exactly:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`

## Facts every task can rely on (verified in code, not re-derived per task)

- `SourceKind` (sim, `web/src/lib/sim/types.ts:12`) is
  `'armory' | 'addon' | 'build' | 'fight' | 'manual'`. The NEW `CurrentCharacter`'s own
  source union is `'addon' | 'build' | 'fight' | 'armory' | 'code'` (spec's exact code
  block) — `'code'` is the pointer's name for what the sim lib calls `'manual'`. These are
  two different types in two different modules; never conflate them.
- The tab strip's existing `TabSource`/`tabStateFor`/`syncTabHrefs`
  (`web/src/lib/sim/tabs.ts`) already carries `{ kind, ref }` across all six `/sim*` tabs and
  is NOT being replaced — `current-character.ts` is a new, separate, localStorage-backed
  layer above it, for pages that have no URL state of their own (a bare `/sim`/`/planner`
  load, or Lane A's non-sim/planner pages).
- `web/src/lib/sim/sources.ts` exports four loaders — `fromStoredCharacter`,
  `fromAddonExport`, `fromPlannerBuild`, `fromLoggedFight` — every one returning
  `Promise<SourceResult>` (`{ ok: true; character: SimCharacter } | { ok: false; message }`).
  `web/src/lib/sim/store.svelte.ts:63-84` has a fifth, currently-private loader,
  `fromPlannerCode(code, ctx)`, for an unsaved planner build's `?code=` link — it stamps
  `kind: 'manual'`. Task 3 moves it into `sources.ts` as a fifth exported loader,
  `fromManualCode`, so both the sim store and the (currently code-blind) tools island can
  use it.
- `SimCharacter` (`web/src/lib/sim/character.ts`) carries `name`, `spec`, `class_slug`,
  `source: CharacterSource`, `point_order`. `specLabel(spec: string): string`
  (`web/src/lib/sim/spec-label.ts`) renders `"Fury Warrior"` from a spec slug — the chip's
  label is `${character.name} · ${specLabel(character.spec)}`.
- `codeForCharacterSpec(spec: CharacterSpec, dataBuild: string): string` and
  `plannerHrefFor(character: SimCharacter, index: TalentIndex | null): string` already exist
  in `web/src/lib/sim/character.ts` and are reused everywhere a full `SimCharacter` is
  already in memory (`CharacterStrip.svelte`, `ComboResults.svelte`). The NEW
  `current-character.ts`'s own `plannerHrefFor`/`simHrefFor` take the bare
  `CurrentCharacter` pointer (no character object) and are for pages/moments that only have
  that.
- `BASE_LEVEL` (`web/src/lib/planner/types.ts`, `FIRST_POINT_LEVEL - 1` = 9) means a level-60
  character has spent 51 points (`talentLevel()` in `character.ts` = `min(60, 9 + order.length)`).
  "Fewer than 51 points" is `character.point_order.length < 51`.
- `/b/<id>` is the saved-build permalink, built by the API as
  `PUBLIC_BASE_URL + "/b/" + id` (`web/src/lib/planner/share.ts:41-47`'s own comment); it is
  NOT a page under `web/src/pages` (served by the API, out of this lane's file ownership).
- Component tests in this repo render with `import { render } from 'svelte/server'` and
  assert on the returned HTML string (see `web/src/components/sim/tools/StatWeights.test.ts`)
  — there is no `@testing-library/svelte`. Use the same pattern for
  `CurrentCharacterChip.test.ts`.
- No file in `web/src` touches `localStorage` today (grepped, zero hits outside this plan's
  own new code) — `current-character.ts` sets the house style: every read/write wrapped in
  try/catch, a `Storage` parameter defaulting to `globalThis.localStorage` so tests inject a
  fake and the module never reads `window` at import time (SSR-safe).
- `web/playwright.config.ts`: `testDir: './tests/e2e'`, `E2E_PORT` env var, `desktop`/`mobile`
  Playwright projects, `FOREVER_DATA=fixture` by default.

---

### Task 1: `current-character.ts` — the pointer module

**Files:**
- Create: `web/src/lib/current-character.ts`
- Create: `web/src/lib/current-character-copy.ts`
- Test: `web/src/lib/current-character.test.ts`

**Interfaces:**
- Produces (used by every later task):
  ```ts
  export type CurrentCharacterSource = 'addon' | 'build' | 'fight' | 'armory' | 'code';
  export interface CurrentCharacter {
    source: CurrentCharacterSource;
    ref: string;
    label: string;
    classSlug: string;
    savedAt: string; // ISO
  }
  export function readCurrent(storage?: Storage): CurrentCharacter | null;
  export function writeCurrent(value: CurrentCharacter, storage?: Storage): void;
  export function clearCurrent(storage?: Storage): void;
  export function plannerHrefFor(current: CurrentCharacter): string;
  export function simHrefFor(current: CurrentCharacter, tab?: SimTabId): string;
  ```
  `SimTabId` is `'quick-sim' | 'gear' | 'drops' | 'talents' | 'weights' | 'specs'`
  (`web/src/lib/sim/tabs.ts`'s `SimTabEntry['id']`) — import the type, do not redeclare it.
- Consumes: `SIM_TABS`, `tabHref`, `defaultSimState`, `withSimState`, `MAX_CODE` from
  `web/src/lib/sim/tabs.ts` / `web/src/lib/sim/url.ts` (already exist, read-only reuse).

**Ruling (write this into the ledger verbatim when this task lands):** `plannerHrefFor`
builds `/planner?code=<ref>` for `source === 'addon' | 'code'` (the spec's own comment says
`ref` IS the FS1 string for these two kinds), `/b/<ref>` for `source === 'build'` (the
existing saved-build permalink — `ref` is the build id), and the honest partial fallback
`/planner?class=<classSlug>` for `'fight'`/`'armory'` (the pointer alone carries no gear or
talents for these kinds — the same class-only fallback `sim/character.ts`'s own
`plannerHrefFor` already uses when it has no talent index). `simHrefFor` builds
`?code=<ref>` for `'addon'`/`'code'` and `?source=<source>&ref=<ref>` for the other three,
through `tabHref`/`withSimState`/`defaultSimState` — never a hand-built query string.

- [ ] **Step 1: Write the failing tests**

```ts
// web/src/lib/current-character.test.ts
// @vitest-environment jsdom
import { describe, expect, it } from 'vitest';
import {
  clearCurrent,
  plannerHrefFor,
  readCurrent,
  simHrefFor,
  writeCurrent,
  type CurrentCharacter,
} from './current-character';

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: (index) => [...map.keys()][index] ?? null,
    get length() {
      return map.size;
    },
  };
}

const sample: CurrentCharacter = {
  source: 'addon',
  ref: 'FS1:1:warrior:orc:0/0/0:',
  label: 'Simfury · Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-21T00:00:00.000Z',
};

describe('readCurrent / writeCurrent / clearCurrent', () => {
  it('round-trips a written value', () => {
    const storage = fakeStorage();
    writeCurrent(sample, storage);
    expect(readCurrent(storage)).toEqual(sample);
  });

  it('is null when nothing has been written', () => {
    expect(readCurrent(fakeStorage())).toBeNull();
  });

  it('is null for malformed JSON rather than throwing', () => {
    const storage = fakeStorage();
    storage.setItem('fs.currentCharacter', '{not json');
    expect(readCurrent(storage)).toBeNull();
  });

  it('refuses a write past the 16 KB cap', () => {
    const storage = fakeStorage();
    const huge: CurrentCharacter = { ...sample, ref: 'x'.repeat(20_000) };
    writeCurrent(huge, storage);
    expect(readCurrent(storage)).toBeNull();
  });

  it('clear removes the stored value', () => {
    const storage = fakeStorage();
    writeCurrent(sample, storage);
    clearCurrent(storage);
    expect(readCurrent(storage)).toBeNull();
  });

  it('never throws when storage.getItem throws (private browsing)', () => {
    const storage = fakeStorage();
    storage.getItem = () => {
      throw new Error('blocked');
    };
    expect(readCurrent(storage)).toBeNull();
  });

  it('never throws when storage.setItem throws (private browsing / quota)', () => {
    const storage = fakeStorage();
    storage.setItem = () => {
      throw new Error('blocked');
    };
    expect(() => writeCurrent(sample, storage)).not.toThrow();
  });
});

describe('plannerHrefFor', () => {
  it('is a ?code= link for an addon-sourced pointer', () => {
    expect(plannerHrefFor(sample)).toBe(`/planner?code=${encodeURIComponent(sample.ref)}`);
  });

  it('is a ?code= link for a code-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'code' })).toBe(
      `/planner?code=${encodeURIComponent(sample.ref)}`,
    );
  });

  it('is the saved-build permalink for a build-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'build', ref: 'b1' })).toBe('/b/b1');
  });

  it('is the honest class-only fallback for a fight-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'fight', ref: 'abc:1' })).toBe(
      '/planner?class=warrior',
    );
  });

  it('is the honest class-only fallback for an armory-sourced pointer', () => {
    expect(plannerHrefFor({ ...sample, source: 'armory', ref: 'us/normal/simfury' })).toBe(
      '/planner?class=warrior',
    );
  });
});

describe('simHrefFor', () => {
  it('is a ?code= link on the bare /sim path for an addon-sourced pointer', () => {
    expect(simHrefFor(sample)).toBe(`/sim?code=${encodeURIComponent(sample.ref)}`);
  });

  it('carries ?source=&ref= for a build-sourced pointer', () => {
    expect(simHrefFor({ ...sample, source: 'build', ref: 'b1' })).toBe('/sim?source=build&ref=b1');
  });

  it('targets the named tab’s own path', () => {
    expect(simHrefFor({ ...sample, source: 'build', ref: 'b1' }, 'drops')).toBe(
      '/sim/drops?source=build&ref=b1',
    );
  });
});
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd web && npx vitest run src/lib/current-character.test.ts`
Expected: FAIL — `current-character.ts` does not exist yet.

- [ ] **Step 3: Write the copy module**

```ts
// web/src/lib/current-character-copy.ts
// Strings for the current-character pointer, the chip, and the bare-load restore banner.
// Shared across every page that mounts CurrentCharacterChip.svelte -- sim, planner, and
// (Lane A) /account, /character/<key>, /addon -- so the wording can never drift between
// pages that describe the same pointer.
export const currentCharacterCopy = {
  restoredNote: 'Restored your last character.',
  forget: 'Forget',
  openInPlanner: 'Open in planner',
  openInSimulator: 'Open in simulator',
  copyAddonCode: 'Copy addon code',
  copiedAddonCode: 'Copied',
  noCharacterLine: 'No character loaded. Paste an addon export in the planner or the simulator.',
  getTheAddon: "Don't have an export? Get the addon.",
} as const;
```

- [ ] **Step 4: Write the pointer module**

```ts
// web/src/lib/current-character.ts
// The site's one "current character" pointer: written on every successful character load
// (sources.ts), read by CurrentCharacterChip.svelte and by a bare /sim or /planner load to
// restore it. A convenience, never the source of truth -- every read and write is wrapped in
// try/catch and this module never touches `window`/`localStorage` at import time, only
// inside the functions below, so importing it is safe during SSR/prerender (Global
// Constraint: "never read it during SSR/prerender").
import { SIM_TABS, tabHref, type SimTabEntry } from './sim/tabs';
import { defaultSimState, withSimState } from './sim/url';

export type CurrentCharacterSource = 'addon' | 'build' | 'fight' | 'armory' | 'code';

export interface CurrentCharacter {
  source: CurrentCharacterSource;
  /** For 'addon' and 'code' the FS1 string itself; otherwise the existing ref format. */
  ref: string;
  /** What the chip shows: "Simfury · Fury Warrior". Derived once at load, never parsed back. */
  label: string;
  classSlug: string;
  savedAt: string; // ISO
}

export type SimTabId = SimTabEntry['id'];

const STORAGE_KEY = 'fs.currentCharacter';
/** An FS1 export with a full bank is several KB; this stays well clear of it while refusing
 *  a pathological write (the spec's own cap). */
const MAX_STORED_BYTES = 16_384;

function storageOf(storage: Storage | undefined): Storage | null {
  if (storage !== undefined) return storage;
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

export function readCurrent(storage?: Storage): CurrentCharacter | null {
  const target = storageOf(storage);
  if (target === null) return null;
  try {
    const raw = target.getItem(STORAGE_KEY);
    if (raw === null) return null;
    const parsed: unknown = JSON.parse(raw);
    if (!isCurrentCharacter(parsed)) return null;
    return parsed;
  } catch {
    return null;
  }
}

function isCurrentCharacter(value: unknown): value is CurrentCharacter {
  if (typeof value !== 'object' || value === null) return false;
  const candidate = value as Record<string, unknown>;
  return (
    typeof candidate.source === 'string' &&
    typeof candidate.ref === 'string' &&
    typeof candidate.label === 'string' &&
    typeof candidate.classSlug === 'string' &&
    typeof candidate.savedAt === 'string'
  );
}

export function writeCurrent(value: CurrentCharacter, storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    const serialised = JSON.stringify(value);
    if (serialised.length > MAX_STORED_BYTES) return;
    target.setItem(STORAGE_KEY, serialised);
  } catch {
    // Private browsing, quota exceeded, or a disabled storage API: the pointer is a
    // convenience, so a failed write is silently skipped rather than surfaced.
  }
}

export function clearCurrent(storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    target.removeItem(STORAGE_KEY);
  } catch {
    // Same as writeCurrent: nothing to surface.
  }
}

/** The saved-build permalink; not a page under web/src/pages (API-served, share.ts's own
 *  `cardUrlFor` comment: `PUBLIC_BASE_URL + "/b/" + id`). */
function buildPermalink(id: string): string {
  return `/b/${id}`;
}

export function plannerHrefFor(current: CurrentCharacter): string {
  if (current.source === 'addon' || current.source === 'code') {
    return `/planner?code=${encodeURIComponent(current.ref)}`;
  }
  if (current.source === 'build') return buildPermalink(current.ref);
  // 'fight' | 'armory': the pointer alone carries no gear or talents for these kinds -- the
  // same honest class-only fallback sim/character.ts's own plannerHrefFor uses when it has
  // no talent index to encode a full FS1 code from.
  return `/planner?class=${encodeURIComponent(current.classSlug)}`;
}

export function simHrefFor(current: CurrentCharacter, tab: SimTabId = 'quick-sim'): string {
  const entry = SIM_TABS.find((row) => row.id === tab) ?? SIM_TABS[0];
  if (current.source === 'addon' || current.source === 'code') {
    return tabHref(entry.href, withSimState(defaultSimState(), { code: current.ref }));
  }
  return tabHref(
    entry.href,
    withSimState(defaultSimState(), { source: current.source, ref: current.ref }),
  );
}
```

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/current-character.test.ts`
Expected: PASS, all cases.

- [ ] **Step 6: Lint and type-check**

Run: `cd web && npx astro check && npm run lint`
Expected: no new errors.

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): add the current-character pointer module\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-1
git add web/src/lib/current-character.ts web/src/lib/current-character-copy.ts web/src/lib/current-character.test.ts
git commit -F .superpowers/commit-msg-1
```

---

### Task 2: `CurrentCharacterChip.svelte`

**Files:**
- Create: `web/src/components/CurrentCharacterChip.svelte`
- Test: `web/src/components/CurrentCharacterChip.test.ts`

**Interfaces:**
- Consumes: `CurrentCharacter`, `plannerHrefFor`, `simHrefFor`, `readCurrent`, `clearCurrent`
  from `web/src/lib/current-character.ts` (Task 1); `currentCharacterCopy` from
  `web/src/lib/current-character-copy.ts` (Task 1).
- Produces: a Svelte component, props below, mounted by later tasks on `/sim*` and
  `/planner` (Lane A mounts it elsewhere later; not this lane's job to wire those pages).

```ts
let {
  current,
  onforget,
  hasOwnPasteBox = false,
}: {
  /** null renders the "no character" line, unless hasOwnPasteBox. */
  current: CurrentCharacter | null;
  onforget: () => void;
  /** True on a page that already has its own paste box (sim, planner): renders nothing
   *  when current is null, per the spec ("renders nothing on pages that have their own
   *  paste box... and one line elsewhere"). */
  hasOwnPasteBox?: boolean;
} = $props();
```

- [ ] **Step 1: Write the failing test**

```ts
// web/src/components/CurrentCharacterChip.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CurrentCharacterChip from './CurrentCharacterChip.svelte';
import { currentCharacterCopy } from '../lib/current-character-copy';
import type { CurrentCharacter } from '../lib/current-character';

const current: CurrentCharacter = {
  source: 'addon',
  ref: 'FS1:1:warrior:orc:0/0/0:',
  label: 'Simfury · Fury Warrior',
  classSlug: 'warrior',
  savedAt: '2026-09-21T00:00:00.000Z',
};

describe('CurrentCharacterChip', () => {
  it('shows the label and the three links when a character is loaded', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).toContain('Simfury · Fury Warrior');
    expect(body).toContain(currentCharacterCopy.openInPlanner);
    expect(body).toContain(currentCharacterCopy.openInSimulator);
    expect(body).toContain(currentCharacterCopy.forget);
    expect(body).toContain('/planner?code=');
    expect(body).toContain('/sim?code=');
  });

  it('renders the no-character line on a page without its own paste box', () => {
    const { body } = render(CurrentCharacterChip, { props: { current: null, onforget: () => {} } });
    expect(body).toContain(currentCharacterCopy.noCharacterLine);
  });

  it('renders nothing on a page that already has its own paste box', () => {
    const { body } = render(CurrentCharacterChip, {
      props: { current: null, onforget: () => {}, hasOwnPasteBox: true },
    });
    expect(body.trim()).toBe('');
  });

  it('fixes its own height so resolving does not move content (data-testid anchor present)', () => {
    const { body } = render(CurrentCharacterChip, { props: { current, onforget: () => {} } });
    expect(body).toContain('data-testid="current-character-chip"');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/components/CurrentCharacterChip.test.ts`
Expected: FAIL — component does not exist.

- [ ] **Step 3: Write the component**

```svelte
<!-- web/src/components/CurrentCharacterChip.svelte -->
<!-- Design 4 (architect review): the visible half of the current-character pointer. Reads
     the pointer it is given (never localStorage itself -- the mounting page owns when to
     read/restore, so this stays a pure render of one prop) and shows up to three links plus
     Forget, in class colour. Fixed min-height so resolving from "no character" to "loaded"
     -- or the reverse, after Forget -- never moves content below it (Global Constraint: no
     layout shift). -->
<script lang="ts">
  import { plannerHrefFor, simHrefFor, type CurrentCharacter } from '../lib/current-character';
  import { currentCharacterCopy } from '../lib/current-character-copy';

  let {
    current,
    onforget,
    hasOwnPasteBox = false,
  }: {
    current: CurrentCharacter | null;
    onforget: () => void;
    hasOwnPasteBox?: boolean;
  } = $props();
</script>

{#if current !== null}
  <section
    class="border-line bg-raised rounded-panel flex min-h-11 flex-wrap items-center gap-3 border px-3 py-2 text-[13px]"
    data-testid="current-character-chip"
  >
    <span class="text-strong font-semibold" data-testid="current-character-label">{current.label}</span>
    <a class="text-nav" href={plannerHrefFor(current)} data-testid="current-character-planner">
      {currentCharacterCopy.openInPlanner}
    </a>
    <a class="text-nav" href={simHrefFor(current)} data-testid="current-character-sim">
      {currentCharacterCopy.openInSimulator}
    </a>
    <button
      type="button"
      class="text-muted ml-auto min-h-11 md:min-h-0"
      onclick={onforget}
      data-testid="current-character-forget"
    >
      {currentCharacterCopy.forget}
    </button>
  </section>
{:else if !hasOwnPasteBox}
  <p class="text-muted min-h-11 text-[13px] md:min-h-0" data-testid="current-character-chip">
    {currentCharacterCopy.noCharacterLine}
  </p>
{/if}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/components/CurrentCharacterChip.test.ts`
Expected: PASS.

- [ ] **Step 5: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/CurrentCharacterChip.svelte src/components/CurrentCharacterChip.test.ts`

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): add CurrentCharacterChip\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-2
git add web/src/components/CurrentCharacterChip.svelte web/src/components/CurrentCharacterChip.test.ts
git commit -F .superpowers/commit-msg-2
```

---

### Task 3: `sources.ts` — a fifth exported loader (`fromManualCode`) + write the pointer on every load

**Files:**
- Modify: `web/src/lib/sim/sources.ts`
- Modify: `web/src/lib/sim/store.svelte.ts:63-84` (remove the now-duplicated private
  `fromPlannerCode`, call the new `fromManualCode` instead)
- Create: `web/src/lib/sim/current-character-bridge.ts` (the `SimCharacter → CurrentCharacter`
  adapter; kept out of `sources.ts` and out of the source-kind-agnostic
  `lib/current-character.ts` on purpose — it is the one place sim-lib types meet the pointer)
- Modify: `web/src/lib/sim/sources.test.ts`
- Modify: `web/src/lib/sim/store.test.ts` (remove/adjust any test asserting on the now-moved
  private `fromPlannerCode`; re-point it at `fromManualCode` from `sources.ts` if such a test
  exists — check with `grep -n fromPlannerCode web/src/lib/sim/store.test.ts` first)

**Interfaces:**
- Produces:
  ```ts
  // sources.ts, alongside the existing four
  export async function fromManualCode(code: string, ctx: LoadContext): Promise<SourceResult>;
  ```
  ```ts
  // current-character-bridge.ts
  import type { SimCharacter } from './character';
  import type { CurrentCharacterSource } from '../current-character';
  export function recordCurrentCharacter(
    character: SimCharacter,
    source: CurrentCharacterSource,
    ref: string,
    storage?: Storage,
  ): void;
  ```
- Consumes: `writeCurrent`, `CurrentCharacter` from `web/src/lib/current-character.ts`
  (Task 1); `specLabel` from `web/src/lib/sim/spec-label.ts` (existing).

**Ruling:** the pointer's `ref` differs by source kind and is not always derivable from the
returned `SimCharacter` alone (an addon-sourced character carries no field holding the
original pasted FS1 string). Every call site below passes `ref` explicitly, using the value
each loader already has as its own parameter or local variable — no re-encoding.

- [ ] **Step 1: Write the failing tests**

```ts
// Add to web/src/lib/sim/sources.test.ts (existing file; import additions alongside the
// existing ones at the top).
// @vitest-environment jsdom  <- add this pragma if the file does not already have it
import { readCurrent } from '../current-character';
// ... (keep existing imports)

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: (index) => [...map.keys()][index] ?? null,
    get length() {
      return map.size;
    },
  };
}

describe('fromManualCode', () => {
  it('decodes the same as fromAddonExport, stamped "manual"', async () => {
    const result = await fromManualCode('FS1:1:warrior:orc:0/0/0:', {
      treeVersion: '1',
    });
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.source.kind).toBe('manual');
  });
});

describe('current-character pointer writes', () => {
  it('fromAddonExport writes an addon-sourced pointer with the pasted code as ref', async () => {
    const storage = fakeStorage();
    const code = 'FS1:1:warrior:orc:0/0/0:';
    await fromAddonExport(code, { treeVersion: '1' }, storage);
    const pointer = readCurrent(storage);
    expect(pointer?.source).toBe('addon');
    expect(pointer?.ref).toBe(code);
    expect(pointer?.classSlug).toBe('warrior');
  });

  it('fromManualCode writes a code-sourced pointer with the pasted code as ref', async () => {
    const storage = fakeStorage();
    const code = 'FS1:1:warrior:orc:0/0/0:';
    await fromManualCode(code, { treeVersion: '1' }, storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'code', ref: code });
  });

  it('writes nothing when the load fails', async () => {
    const storage = fakeStorage();
    await fromAddonExport('garbage', { treeVersion: '1' }, storage);
    expect(readCurrent(storage)).toBeNull();
  });
});
```

Note for the implementer: read the existing `sources.test.ts` first for its fixture setup
(fake `fetch`, `treeVersion`, reference/talent fixtures) and reuse it rather than
reintroducing a second fixture rig — grep for `describe('fromAddonExport'` to find the
existing tests this task's new `describe` blocks sit beside.

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts`
Expected: FAIL — `fromManualCode` not exported, loaders take no `storage` param yet.

- [ ] **Step 3: Write `current-character-bridge.ts`**

```ts
// web/src/lib/sim/current-character-bridge.ts
// The one place a SimCharacter becomes a CurrentCharacter pointer. Kept separate from
// current-character.ts (source-kind agnostic, no sim-lib import) and from sources.ts (keeps
// that file's own header claim -- "the four/five ways a character reaches the simulator" --
// free of a second concern).
import { writeCurrent, type CurrentCharacter, type CurrentCharacterSource } from '../current-character';
import type { SimCharacter } from './character';
import { specLabel } from './spec-label';

export function recordCurrentCharacter(
  character: SimCharacter,
  source: CurrentCharacterSource,
  ref: string,
  storage?: Storage,
): void {
  const pointer: CurrentCharacter = {
    source,
    ref,
    label: `${character.name} · ${specLabel(character.spec)}`,
    classSlug: character.class_slug,
    savedAt: new Date().toISOString(),
  };
  writeCurrent(pointer, storage);
}
```

- [ ] **Step 4: Add `fromManualCode` to `sources.ts` and thread an optional `storage` param through all five loaders**

Read `web/src/lib/sim/sources.ts` in full before editing (already done in planning; the
implementer should re-read it, since exact line numbers will have shifted once earlier tasks
land). Changes:

1. Add the import: `import { recordCurrentCharacter } from './current-character-bridge';`
2. Add `storage?: Storage` as a final parameter on `fromStoredCharacter`, `fromAddonExport`,
   `fromPlannerBuild`, `fromLoggedFight`, and the new `fromManualCode` (below). Every caller
   in the codebase (`store.svelte.ts`, `bulk-store.svelte.ts`) calls these with two
   arguments today; adding an optional third is source-compatible, no other call site needs
   to change.
3. At each `ok: true` return site, call `recordCurrentCharacter` before returning, with the
   `ref` each function already has to hand:
   - `fromStoredCharacter`'s addon branch: `recordCurrentCharacter(result.character, 'addon', code, storage)`
     (the local `code` variable already decoded from `input.gear`).
   - `fromStoredCharacter`'s non-addon branch: `recordCurrentCharacter(result.character, input.source === 'armory' ? 'armory' : 'fight', characterKey, storage)`
     (input.source here is never `'addon'`, since that branch is the `if` above; it is
     `'armory'` or `'fight'` per `SimInput.source`'s own doc comment).
   - `fromAddonExport`: `recordCurrentCharacter(character.character, 'addon', code, storage)`
     where `code` is the function's own first parameter.
   - `fromPlannerBuild`: `recordCurrentCharacter(character.character, 'build', record.id, storage)`.
   - `fromLoggedFight`: `recordCurrentCharacter(character.character, 'fight', ref, storage)`
     where `ref` is the function's own first parameter (already the exact `<report>:<fight>`
     or, after Task 6, `<report>:<fight>:<guid>` string).
   - `fromManualCode` (new, placed near `fromPlannerBuild` in the file): identical body to
     `store.svelte.ts`'s current private `fromPlannerCode` (copy it verbatim, changing
     nothing but the name and adding the `storage` param and the
     `recordCurrentCharacter(result.character, 'code', code, storage)` call before returning
     `ok: true`).

   Concretely, `fromManualCode`:

   ```ts
   /**
    * An unsaved planner build's own FS1 code, into a `'manual'`-sourced character. Moved
    * here from store.svelte.ts (Task 3, current-character spec) so the tools island --
    * previously blind to `?code=` entirely -- can bootstrap from one too, through the same
    * loader the sim page already used.
    */
   export async function fromManualCode(
     code: string,
     ctx: LoadContext,
     storage?: Storage,
   ): Promise<SourceResult> {
     const classSlug = code.trim().split(':')[2] ?? '';
     let talents;
     let classes;
     let races;
     try {
       [talents, { classes, races }] = await Promise.all([
         loadTalents(ctx.treeVersion, classSlug),
         reference(ctx),
       ]);
     } catch {
       return { ok: false, message: simCopy.characterFailed };
     }
     const result = characterFromFs1(code, talents, classes, races, {
       kind: 'manual',
       ref: '',
       captured_at: new Date().toISOString(),
     });
     if (result.ok) recordCurrentCharacter(result.character, 'code', code, storage);
     return result;
   }
   ```

   Every other `ok: true` site gets the equivalent one-line insertion — capture the
   character in a local (`const character = characterFromFs1(...)` etc., where the function
   does not already do so) so `recordCurrentCharacter` can read it before the `return`.

- [ ] **Step 5: Update `store.svelte.ts` to use the moved loader**

Remove the private `fromPlannerCode` function (lines 55-84 per the pre-task read; re-locate
by searching for the function name) and its now-unused imports (`characterFromFs1`, if
nothing else in the file needs it — check with `grep -n characterFromFs1
src/lib/sim/store.svelte.ts` before removing the import). Replace every call site
(`fromPlannerCode(code, ctx)` / `fromPlannerCode(init.code, ctx)`) with
`fromManualCode(code, ctx)` / `fromManualCode(init.code, ctx)`, importing `fromManualCode`
from `./sources` alongside the other four loaders already imported there.

- [ ] **Step 6: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts src/lib/sim/store.test.ts`
Expected: PASS.

- [ ] **Step 7: Run the full sim lib suite (nothing else regresses)**

Run: `cd web && npx vitest run src/lib/sim`
Expected: PASS.

- [ ] **Step 8: Lint, type-check, prettier, file-size check**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/sim/sources.ts src/lib/sim/sources.test.ts src/lib/sim/store.svelte.ts src/lib/sim/current-character-bridge.ts && wc -l src/lib/sim/store.svelte.ts src/lib/sim/sources.ts`
Expected: `store.svelte.ts` shrinks (a function moved out); `sources.ts` grows by roughly
the size of the moved function plus five one-line pointer-write calls; neither exceeds 800.

- [ ] **Step 9: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): write the current-character pointer on every successful load\n\nMoves the manual-code loader into sources.ts (fromManualCode) so the\ntools island can share it, and writes the pointer from all five loaders.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-3
git add web/src/lib/sim/sources.ts web/src/lib/sim/sources.test.ts web/src/lib/sim/store.svelte.ts web/src/lib/sim/store.test.ts web/src/lib/sim/current-character-bridge.ts
git commit -F .superpowers/commit-msg-3
```

---

### Task 4: the tools island bootstraps from `?code=` and from the stored pointer; `?instance=` preselect

**Files:**
- Modify: `web/src/lib/sim/bulk-store.svelte.ts`
- Modify: `web/src/lib/sim/tabs.ts` (flip `supportsCode` to `true` for all six tabs)
- Modify: `web/src/lib/sim/tabs.test.ts` (update the assertion the flip changes)
- Modify: `web/src/components/sim/tools/ToolsView.svelte`
- Modify: `web/src/lib/sim/bulk-store.test.ts`

**Interfaces:**
- Consumes: `fromManualCode` (Task 3, `sources.ts`); `readCurrent` (Task 1,
  `current-character.ts`); `LootSource.id` shape `'<kind>:<zone-slug>'`
  (`web/src/lib/sim/loot.ts:44-46`, existing).
- Produces: `BulkStoreInit.code?: string`, `createBulkStore(...).loadCode(code: string): Promise<void>`
  (mirrors the existing `loadAddon`/`loadBuild`/`loadFight` methods at
  `bulk-store.svelte.ts`'s return object).

- [ ] **Step 1: Write the failing tests**

```ts
// Add to web/src/lib/sim/bulk-store.test.ts, beside the existing loadAddon/loadBuild tests
// (read the file first to match its existing fixture/mock-fetch setup).
it('loadCode adopts a character from an FS1 code, the same as loadAddon does from a paste', async () => {
  const store = createBulkStore({ tool: 'gear', treeVersion: FIXTURE_TREE_VERSION /* match existing fixture const name */ });
  await store.loadCode('FS1:1:warrior:orc:0/0/0:');
  expect(store.character?.class_slug).toBe('warrior');
});

it('BulkStoreInit.code bootstraps a character on creation, same as source/ref does', async () => {
  const store = createBulkStore({
    tool: 'gear',
    treeVersion: FIXTURE_TREE_VERSION,
    code: 'FS1:1:warrior:orc:0/0/0:',
  });
  // Bootstrapping from init.code happens on first access the same way source/ref does in
  // ToolsView.svelte's onMount -- this store-level test calls loadCode directly if the
  // store does not self-bootstrap from init.code; read the existing source/ref handling in
  // createBulkStore (init.source / init.ref are read by ToolsView.svelte's onMount, NOT by
  // the store itself) and follow that exact pattern: init.code is a value ToolsView reads
  // and calls store.loadCode(...) with, not a self-triggering store field. Adjust this test
  // to assert `store.loadCode` exists and behaves as above; drop the "bootstraps on
  // creation" framing if createBulkStore does not self-bootstrap for source/ref either
  // (verify against the real file before writing this test, per the note below).
});
```

Note for the implementer: before writing the second test, re-read
`web/src/lib/sim/bulk-store.svelte.ts`'s `BulkStoreInit` and confirm whether `source`/`ref`
are consumed inside `createBulkStore` itself or only read by `ToolsView.svelte`'s `onMount`
(the recon for this plan found the latter — `ToolsView.svelte`'s `onMount` calls
`store.loadAddon`/`loadBuild`/`loadFight` explicitly; `init.source`/`init.ref` do not appear
to be read inside `createBulkStore` itself, only carried on the init object). If so, `code`
should behave identically: `BulkStoreInit.code` is not read by `createBulkStore`, only the
new `loadCode` method is exposed, and `ToolsView.svelte` (Step 4 below) is what actually
calls it during its own `onMount`. Delete the second test above and keep only the first if
that is what the real file shows — do not assert behaviour the file does not have.

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd web && npx vitest run src/lib/sim/bulk-store.test.ts`
Expected: FAIL — `loadCode` does not exist.

- [ ] **Step 3: Add `loadCode` to `bulk-store.svelte.ts`**

Import `fromManualCode` alongside the existing four loader imports from `./sources`. Add one
line beside the existing `loadAddon`/`loadBuild`/`loadFight` in the returned object:

```ts
loadAddon: (code: string) => adopt(fromAddonExport(code, ctx)),
loadBuild: (id: string) => adopt(fromPlannerBuild(id, ctx)),
loadFight: (ref: string) => adopt(fromLoggedFight(ref, ctx)),
loadCode: (code: string) => adopt(fromManualCode(code, ctx)),
loadStored: (path: CharacterPath) => adopt(fromStoredCharacter(path, ctx)),
```

- [ ] **Step 4: Flip `supportsCode` in `tabs.ts`**

```ts
// web/src/lib/sim/tabs.ts -- SIM_TABS, replace all six `supportsCode` values with `true`,
// and update the doc comment on SimTabEntry.supportsCode: it now reads, in full:
//
// /**
//  * Whether this tab's own destination can bootstrap a character from `?code=`. Every tab
//  * can, as of the current-character spec (2026-09-21): the tools island
//  * (ToolsView.svelte / bulk-store.svelte.ts) gained its own `loadCode`, matching
//  * SimView.svelte's long-standing `fromPlannerCode` (now `sources.ts`'s
//  * `fromManualCode`). Kept as a field, not simplified away, because `tabStateFor` still
//  * needs to know per tab whether a fallback code is safe to offer -- the day any tab
//  * loses that ability again, this is the one place to flip back.
//  */
```

- [ ] **Step 5: Update the flipped assertion in `tabs.test.ts`**

```ts
// Replace this existing test's body:
it('marks only the two SimView-served tabs as able to bootstrap from ?code=', () => {
  const supportsCode = SIM_TABS.filter((tab) => tab.supportsCode).map((tab) => tab.id);
  expect(supportsCode).toEqual(['quick-sim', 'specs']);
});
// with:
it('marks every tab able to bootstrap from ?code=, now the tools island has loadCode too', () => {
  expect(SIM_TABS.every((tab) => tab.supportsCode)).toBe(true);
});
```

- [ ] **Step 6: Wire `ToolsView.svelte`'s bootstrap**

In the `bootstrap` object (`untrack(() => { ... })` near the top of the `<script>`), add:

```ts
code: params.get('code') ?? '',
instance: params.get('instance') ?? '',
```

In the `onMount` body, before the existing `if (bootstrap.source !== '' && bootstrap.ref !== '')`
block, add the `?code=` branch and, when the URL carries neither `?code=` nor
`?source=&ref=`, fall back to the stored pointer:

```ts
if (bootstrap.code !== '') {
  void store.loadCode(bootstrap.code);
} else if (bootstrap.source !== '' && bootstrap.ref !== '') {
  if (bootstrap.source === 'addon') void store.loadAddon(bootstrap.ref);
  else if (bootstrap.source === 'build') void store.loadBuild(bootstrap.ref);
  else if (bootstrap.source === 'fight') void store.loadFight(bootstrap.ref);
} else {
  const stored = readCurrent();
  if (stored !== null) {
    if (stored.source === 'addon' || stored.source === 'code') void store.loadCode(stored.ref);
    else if (stored.source === 'build') void store.loadBuild(stored.ref);
    else if (stored.source === 'fight') void store.loadFight(stored.ref);
    // 'armory': no loader here reads a bare CharacterPath from a ref string alone; skip
    // rather than guess (the existing loadStored(path) takes a structured CharacterPath,
    // not this pointer's plain ref string -- restoring an armory-sourced pointer on the
    // tools island is out of this task's scope, honestly skipped rather than half-built).
  }
}
```

Import `readCurrent` from `'../../../lib/current-character'`.

For the `?instance=` preselect: add a second `$effect`, modelled on the existing pin
`$effect` immediately below it in the same file (same `phase === 'idle'` gate, same
apply-once guard):

```ts
let instanceApplied = false;
$effect(() => {
  if (instanceApplied || store.phase !== 'idle' || store.character === null) return;
  if (tool !== 'drops' || bootstrap.instance === '') return;
  instanceApplied = true;
  const match = [...store.loot.sources].find((source) => source.id.endsWith(`:${bootstrap.instance}`));
  if (match !== undefined) store.toggleSource(match.id);
});
```

(`store.loot.sources` and `store.toggleSource` already exist on `BulkStore` — confirm the
exact getter/method names by reading the current `bulk-store.svelte.ts` return object before
writing this; the plan's earlier recon read them as `get loot()` returning `LootFile` with a
`sources` array, and `toggleSource(sourceId: string, bossId = '')`.)

- [ ] **Step 7: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/sim/bulk-store.test.ts src/lib/sim/tabs.test.ts`
Expected: PASS.

- [ ] **Step 8: Full sim lib + component suite**

Run: `cd web && npx vitest run src/lib/sim src/components/sim`
Expected: PASS (no regression from the `supportsCode` flip — search for any other test
asserting the old `['quick-sim', 'specs']` list, e.g. in a `ToolsView` or `SimTabs` test, and
update it the same way).

- [ ] **Step 9: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/sim/bulk-store.svelte.ts src/lib/sim/tabs.ts src/components/sim/tools/ToolsView.svelte`

- [ ] **Step 10: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): the tools island bootstraps from ?code= and the stored character\n\nAdds loadCode to the bulk store, flips every sim tab to supportsCode,\nand wires ?instance= to preselect a Droptimizer source.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-4
git add web/src/lib/sim/bulk-store.svelte.ts web/src/lib/sim/bulk-store.test.ts web/src/lib/sim/tabs.ts web/src/lib/sim/tabs.test.ts web/src/components/sim/tools/ToolsView.svelte
git commit -F .superpowers/commit-msg-4
```

---

### Task 5: `/sim` restores the stored pointer on a bare load, and mounts the chip

**Files:**
- Modify: `web/src/components/sim/SimView.svelte`
- Modify: `web/src/lib/sim/store.svelte.ts` (only if the bare-load check needs a store-level
  hook; prefer doing this entirely in `SimView.svelte`'s own mount logic, since the store's
  file is at the size ceiling)

**Interfaces:**
- Consumes: `readCurrent`, `clearCurrent`, `CurrentCharacter` (Task 1); `CurrentCharacterChip`
  (Task 2); `currentCharacterCopy.restoredNote` (Task 1).

- [ ] **Step 1: Read `SimView.svelte`'s current bootstrap logic in full**

Locate the block that decides what to bootstrap from (`request` > `code` > `source`/`ref`,
per the existing doc comment on `SimStoreInit`). This is where the "bare" check belongs: bare
means none of `?req=`, `?code=`, `?source=`+`?ref=` are present in
`window.location.search`.

- [ ] **Step 2: Write the failing component test**

```ts
// Add to an existing SimView-adjacent test file if one exists (check for
// web/src/components/sim/SimView.test.ts first with `find`); if none exists, this step's
// assertion is covered by Task 7's e2e spec instead — component-level `svelte/server`
// rendering cannot exercise `onMount`/`$effect` (server rendering skips both), so this
// behaviour is unit-tested at the level that IS reachable without a browser: the "what to
// bootstrap from" decision, extracted as a small pure function.
```

Extract the bare-load decision into a pure, testable function rather than leaving it
inline in `<script>` (`$effect`/`onMount` bodies are not unit-testable without a browser;
a pure function next to them is):

```ts
// Add to web/src/lib/sim/store.svelte.ts, near bootstrapSource (it is the same kind of
// pure URL-shape decision bootstrapSource already is):
/** True when none of the URL's own bootstrap params are present -- the load-nothing case a
 *  stored pointer restore is for. */
export function isBareSimUrl(search: string): boolean {
  const params = new URLSearchParams(search);
  return params.get('req') === null && params.get('code') === null && (params.get('source') === null || params.get('ref') === null);
}
```

```ts
// Test, added to web/src/lib/sim/store.test.ts:
describe('isBareSimUrl', () => {
  it('is true for an empty query', () => {
    expect(isBareSimUrl('')).toBe(true);
  });
  it('is false when ?code= is present', () => {
    expect(isBareSimUrl('?code=FS1:1:warrior:orc:0/0/0:')).toBe(false);
  });
  it('is false when ?source=&ref= are both present', () => {
    expect(isBareSimUrl('?source=build&ref=b1')).toBe(false);
  });
  it('is true when only ?source= is present without ?ref=', () => {
    expect(isBareSimUrl('?source=build')).toBe(true);
  });
});
```

- [ ] **Step 3: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/sim/store.test.ts -t isBareSimUrl`
Expected: FAIL — not exported yet.

- [ ] **Step 4: Implement `isBareSimUrl` in `store.svelte.ts`** (code above)

- [ ] **Step 5: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/sim/store.test.ts -t isBareSimUrl`
Expected: PASS.

- [ ] **Step 6: Wire `SimView.svelte`**

In the mount/bootstrap block, when `isBareSimUrl(window.location.search)` is true, read the
stored pointer and adopt it (same source-to-loader mapping as Task 4's `ToolsView.svelte`
addition — `loadAddon`/`loadFight`/`loadBuild` for `store.fromPlannerCode` calls, use
`store.loadCode`-equivalent method that already exists on `createSimStore`'s return object
for the `'addon' | 'code'` case; check the exact exposed method name, likely
`applyCode`/`loadCode` — read `store.svelte.ts`'s return object to confirm before wiring).
Add a `let restored = $state(false);` flag, set it true only when this fallback path is the
one that loaded the character (not when the URL itself carried a source). Render, directly
above `<CharacterStrip ...>` at line ~531:

```svelte
{#if restored}
  <p class="text-muted px-[18px] text-[13px] md:px-0" data-testid="sim-restored-note">
    {currentCharacterCopy.restoredNote}
    <button
      type="button"
      class="text-nav ml-1"
      onclick={() => {
        clearCurrent();
        restored = false;
      }}
      data-testid="sim-restored-forget"
    >
      {currentCharacterCopy.forget}
    </button>
  </p>
{/if}
```

This sits inside the same `min-h-[1044px] md:min-h-[531px]` reserved region `sim.astro`
already budgets for the island's hydrated content (Global Constraint: no layout shift is
already satisfied by that existing reservation — do not add a second one).

- [ ] **Step 7: Run the sim component suite**

Run: `cd web && npx vitest run src/components/sim src/lib/sim`
Expected: PASS.

- [ ] **Step 8: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/sim/SimView.svelte src/lib/sim/store.svelte.ts src/lib/sim/store.test.ts`

- [ ] **Step 9: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): /sim restores the stored character on a bare load\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-5
git add web/src/components/sim/SimView.svelte web/src/lib/sim/store.svelte.ts web/src/lib/sim/store.test.ts
git commit -F .superpowers/commit-msg-5
```

---

### Task 6: `fromLoggedFight` accepts the optional `<guid>` ref part

**Files:**
- Modify: `web/src/lib/sim/sources.ts`
- Modify: `web/src/lib/sim/sources.test.ts`

**Interfaces:**
- Produces: `parseFightRef(ref: string)` now also returns an optional `guid`; `fromLoggedFight`
  selects the named combatant when present.

- [ ] **Step 1: Write the failing tests**

```ts
// Add to sources.test.ts, beside the existing parseFightRef/fromLoggedFight tests.
describe('parseFightRef with a combatant guid', () => {
  it('parses the third part as guid', () => {
    expect(parseFightRef('abcdefabcdef:2:Player-1234')).toEqual({
      reportId: 'abcdefabcdef',
      fightIndex: 2,
      guid: 'Player-1234',
    });
  });

  it('leaves guid undefined when the ref has only two parts, same as before', () => {
    expect(parseFightRef('abcdefabcdef:2')).toEqual({
      reportId: 'abcdefabcdef',
      fightIndex: 2,
      guid: undefined,
    });
  });
});

describe('fromLoggedFight with a combatant guid', () => {
  it('selects the named combatant instead of the first-dps default', async () => {
    // Reuse this file's existing fetch/report-summary fixture setup (grep the existing
    // fromLoggedFight describe block for how `fetchReportMeta`/`fetchSummary` are stubbed)
    // and add a second dps-role combatant to the roster fixture with a distinct name and
    // guid, so the assertion can tell which one loaded.
    const result = await fromLoggedFight('report1:1:guid-of-second-dps', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('SecondDps');
  });

  it('falls back to the first-dps rule when no guid is given, unchanged from before', async () => {
    const result = await fromLoggedFight('report1:1', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('FirstDps');
  });

  it('falls back to the first-dps rule when the named guid is not in the roster', async () => {
    const result = await fromLoggedFight('report1:1:not-a-real-guid', ctx);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.character.name).toBe('FirstDps');
  });
});
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts -t "combatant guid"`
Expected: FAIL.

- [ ] **Step 3: Implement**

```ts
// sources.ts -- widen the regex and parseFightRef's return type:
const FIGHT_REF = /^([a-z2-7]{12}):(\d{1,6})(?::(.+))?$/;

export function parseFightRef(
  ref: string,
): { reportId: string; fightIndex: number; guid?: string } | null {
  const match = FIGHT_REF.exec(ref.trim());
  if (match === null) return null;
  return {
    reportId: match[1],
    fightIndex: Number.parseInt(match[2], 10),
    guid: match[3],
  };
}
```

In `fromLoggedFight`, change the roster/combatant selection:

```ts
const roster =
  (parsed.guid !== undefined
    ? summary.roster.find((row) => row.class !== undefined && row.guid === parsed.guid)
    : undefined) ?? summary.roster.find((row) => row.class !== undefined && row.role === 'dps');
```

(Keep everything else in the function unchanged — `combatant`, `classSlug`, etc. all read
off whichever `roster` row this resolves to, same as before.)

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts`
Expected: PASS, including every pre-existing test in the file (the regex/selection change
must not alter behaviour for a two-part ref).

- [ ] **Step 5: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/sim/sources.ts src/lib/sim/sources.test.ts`

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): fromLoggedFight reads an optional combatant guid\n\nA report link can now name the exact combatant\n(<report>:<fight>:<guid>); falls back to the first-dps rule when\nabsent, unchanged from before. Lane A writes the per-combatant link.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-6
git add web/src/lib/sim/sources.ts web/src/lib/sim/sources.test.ts
git commit -F .superpowers/commit-msg-6
```

---

### Task 7: "Plan it" on every Droptimizer / Top Gear upgrade row

**Files:**
- Modify: `web/src/lib/sim/combos.ts`
- Modify: `web/src/lib/sim/combos.test.ts`
- Modify: `web/src/components/sim/tools/ComboResults.svelte`
- Modify: `web/src/components/sim/tools/DropResults.svelte`
- Modify: `web/src/components/sim/tools/DropResults.test.ts`
- Modify: `web/src/lib/sim/copy.ts` — add exactly one new key to the existing `bulkCopy`
  object (`planIt: 'Plan it'`); this is the one place adding to `sim/copy.ts` (1247 lines,
  already over the file's own soft ceiling) is still correct rather than a new module: it is
  a single short key beside its siblings `openInPlanner`/`dropsPin`, not new structure, and
  splitting one key into its own module would scatter `bulkCopy`'s own vocabulary rather
  than concentrate it.

**Interfaces:**
- Produces:
  ```ts
  // combos.ts
  export function gearForCombo(baseGear: readonly GearSlot[], combo: Combo): GearSlot[];
  ```
  `winningGear(result)` becomes `gearForCombo(result.request.character.gear, result.combos[0])`
  internally (no external signature change — same DRY the file's own header already asks
  for).

- [ ] **Step 1: Write the failing test**

```ts
// Add to web/src/lib/sim/combos.test.ts, beside the existing winningGear tests (read them
// first for the exact fixture Combo/GearSlot shapes already in use in this file).
describe('gearForCombo', () => {
  it('writes one substitution over the base gear, same as winningGear does for the leader', () => {
    const base: GearSlot[] = [{ slot: 'head', item_id: 1 }];
    const combo: Combo = {
      substitutions: [{ kind: 'item', slot: 'head', item_id: 2 }],
      dps: ZERO_ESTIMATE, // reuse this file's existing zero-estimate fixture constant
      delta: ZERO_ESTIMATE,
      group: 0,
    };
    expect(gearForCombo(base, combo)).toEqual([{ slot: 'head', item_id: 2 }]);
  });

  it('agrees with winningGear on the leader combo', () => {
    // Reuse this file's existing BulkResult fixture (the one winningGear's own tests use).
    expect(gearForCombo(fixtureResult.request.character.gear, fixtureResult.combos[0])).toEqual(
      winningGear(fixtureResult),
    );
  });

  it('removes an emptied off-hand rather than writing item_id 0, same rule as winningGear', () => {
    const base: GearSlot[] = [
      { slot: 'main_hand', item_id: 1 },
      { slot: 'off_hand', item_id: 2 },
    ];
    const combo: Combo = {
      substitutions: [
        { kind: 'item', slot: 'main_hand', item_id: 3 },
        { kind: 'item', slot: 'off_hand', item_id: 0 },
      ],
      dps: ZERO_ESTIMATE,
      delta: ZERO_ESTIMATE,
      group: 0,
    };
    expect(gearForCombo(base, combo)).toEqual([{ slot: 'main_hand', item_id: 3 }]);
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/sim/combos.test.ts -t gearForCombo`
Expected: FAIL.

- [ ] **Step 3: Refactor `winningGear` into `gearForCombo`**

```ts
// combos.ts -- replace the body of winningGear with a call to the new, generalised function.
export function gearForCombo(baseGear: readonly GearSlot[], combo: Combo): GearSlot[] {
  let gear: GearSlot[] = baseGear.map((slot) => ({ ...slot }));
  for (const sub of combo.substitutions) {
    if (sub.kind !== 'item' || sub.slot === undefined || sub.item_id === undefined) continue;
    if (isEmptiedOffHand(sub)) {
      gear = gear.filter((slot) => slot.slot !== sub.slot);
      continue;
    }
    const next: GearSlot = { slot: sub.slot, item_id: sub.item_id };
    if (sub.enchant !== undefined && sub.enchant > 0) next.enchant = sub.enchant;
    if (sub.suffix !== undefined && sub.suffix > 0) next.suffix = sub.suffix;
    const at = gear.findIndex((slot) => slot.slot === sub.slot);
    gear = at >= 0 ? gear.map((slot, index) => (index === at ? next : slot)) : [...gear, next];
  }
  return gear;
}

export function winningGear(result: BulkResult): GearSlot[] {
  return gearForCombo(result.request.character.gear, result.combos[0] ?? { substitutions: [], dps: result.equipped, delta: result.equipped, group: 0 });
}
```

Keep every doc comment already on `winningGear` and `isEmptiedOffHand` — move the detailed
"why gear_slots, why the sentinel" explanation onto `gearForCombo` (it is now the function
that actually does the writing) and leave a one-line pointer on `winningGear` ("the leader's
own combo, through `gearForCombo` below").

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/sim/combos.test.ts`
Expected: PASS, including every pre-existing `winningGear` test unchanged.

- [ ] **Step 5: Add the `planIt` copy key**

```ts
// web/src/lib/sim/copy.ts, inside the existing bulkCopy object, beside `openInPlanner`:
planIt: 'Plan it',
```

- [ ] **Step 6: Add a per-row "Plan it" link to `ComboResults.svelte`**

In the `{#each rows as row, index (...)}` block, inside the existing row `<div role="row">`,
beside the percent cell, add a sixth cell (widen the grid template columns from
`28px_minmax(0,2fr)_84px_96px_56px` to `28px_minmax(0,2fr)_84px_96px_56px_auto` in both the
header row and the data row):

```svelte
<span role="cell">
  <a
    class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
    href={`/planner?code=${encodeURIComponent(codeForCharacterSpec({ ...result.request.character, gear: gearForCombo(result.request.character.gear, row.combo) }, treeVersion))}`}
    data-testid={`sim-combo-plan-it-${comboKey(row)}`}
  >
    {bulkCopy.planIt}
  </a>
</span>
```

Add a header cell too (`<span role="columnheader"></span>`, empty label — the action column
needs no header text, same pattern the existing table already avoids elsewhere for an
action-only column; check the design system for whether an empty `columnheader` needs an
`aria-label` — if so add `aria-label={bulkCopy.planIt}` to it). Define `comboKey` in this
file the same way `DropResults.svelte` already does (copy that function verbatim — both
files independently need a stable per-row key; a shared helper is not worth adding for four
lines duplicated between two already-separate components rendering two different layouts of
the same data). Import `gearForCombo` from `../../../lib/sim/combos` alongside the other
`combos.ts` imports already at the top of the file.

- [ ] **Step 7: Add a per-row "Plan it" link to `DropResults.svelte`**

In the "every upgrade" flat list (`{#each upgrades as row (comboKey(row))}`), beside the
existing "Pin into Top Gear" button, add:

```svelte
<a
  class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
  href={`/planner?code=${encodeURIComponent(codeForCharacterSpec({ ...result.request.character, gear: gearForCombo(result.request.character.gear, row.combo) }, treeVersion))}`}
  data-testid={`sim-drops-plan-it-${pinId}`}
>
  {bulkCopy.planIt}
</a>
```

Import `codeForCharacterSpec` from `'../../../lib/sim/character'` and `gearForCombo` from
`'../../../lib/sim/combos'` (add to the existing `combos.ts` import line, which already
imports `comboRows`, `deltaLabel`, `sourceNameOfCombo`).

- [ ] **Step 8: Extend `DropResults.test.ts`**

```ts
// Add to the existing describe block, using the file's existing fixture BulkResult.
it('renders a Plan it link on every upgrade row, targeting the planner', () => {
  const { body } = render(DropResults, { props: { /* ...existing props from the file's own setup */ } });
  expect(body).toContain('/planner?code=');
  expect(body).toContain(bulkCopy.planIt);
});
```

- [ ] **Step 9: Run the component tests**

Run: `cd web && npx vitest run src/components/sim/tools/ComboResults.test.ts src/components/sim/tools/DropResults.test.ts src/lib/sim/combos.test.ts`

Note: `ComboResults.svelte` has no existing `.test.ts` (confirmed by the earlier file
listing) — if this task's review finds the component untestable without one, add
`web/src/components/sim/tools/ComboResults.test.ts` using the exact `svelte/server`
`render()` pattern `StatWeights.test.ts` and `DropResults.test.ts` already use, asserting the
same two facts (a `/planner?code=` href and the `bulkCopy.planIt` label appear once per row).

- [ ] **Step 10: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/sim/combos.ts src/lib/sim/combos.test.ts src/components/sim/tools/ComboResults.svelte src/components/sim/tools/DropResults.svelte src/components/sim/tools/DropResults.test.ts src/lib/sim/copy.ts`

- [ ] **Step 11: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): a Plan it link on every Droptimizer and Top Gear upgrade row\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-7
git add web/src/lib/sim/combos.ts web/src/lib/sim/combos.test.ts web/src/components/sim/tools/ComboResults.svelte web/src/components/sim/tools/ComboResults.test.ts web/src/components/sim/tools/DropResults.svelte web/src/components/sim/tools/DropResults.test.ts web/src/lib/sim/copy.ts
git commit -F .superpowers/commit-msg-7
```

---

### Task 8: "Get the addon" in every paste box

**Files:**
- Modify: `web/src/components/sim/SourceSwitcher.svelte`
- Modify: `web/src/components/planner/ImportBox.svelte`

**Interfaces:**
- Consumes: `currentCharacterCopy.getTheAddon` (Task 1).

**Before writing code:** run
`grep -rln "textarea" web/src/components --include="*.svelte"` again against the
then-current tree (Tasks 1-7 may not have added a new paste box, but re-verify — the
dispatch's own warning is specifically about missing a second instance of a control).
Confirm the result is still exactly `ImportBox.svelte`, `Planner.svelte` (composes
`ImportBox`, not a second box), `QueriesView.svelte` (report SQL editor, not a character
paste box — out of scope), `SourceSwitcher.svelte`, `RequestDrawer.svelte` (a JSON request
editor, not a character paste box — out of scope). `SourceSwitcher.svelte` is mounted by
both `SimView.svelte` and `ToolsView.svelte` (one shared component — confirmed by Task 3/4's
own reading), so editing it once covers all six `/sim*` tabs; `ImportBox.svelte` is mounted
once, by `Planner.svelte`.

- [ ] **Step 1: `SourceSwitcher.svelte`**

Add, directly below the existing scope note (`data-testid="sim-sources-scope-note"`):

```svelte
<p class="text-muted text-[12px]">
  <a href="/addon" class="text-nav">{currentCharacterCopy.getTheAddon}</a>
</p>
```

Import `currentCharacterCopy` from `'../../lib/current-character-copy'`.

- [ ] **Step 2: `ImportBox.svelte`**

Add, directly below the existing `{#each notes as note}` block (or, when there are no notes,
still rendered — place it outside the `{#if failure}`/`{#each notes}` conditionals so it is
always present):

```svelte
<p class="text-muted text-[13px]">
  <a href="/addon" class="text-nav">{currentCharacterCopy.getTheAddon}</a>
</p>
```

Import `currentCharacterCopy` from `'../../lib/current-character-copy'`.

- [ ] **Step 3: Write/extend component tests**

`SourceSwitcher.svelte` has no `.test.ts` today (confirmed by the file listing) and neither
does `ImportBox.svelte`. Add both, using the `svelte/server` `render()` pattern:

```ts
// web/src/components/sim/SourceSwitcher.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SourceSwitcher from './SourceSwitcher.svelte';
import { currentCharacterCopy } from '../../lib/current-character-copy';

describe('SourceSwitcher', () => {
  it('links to /addon', () => {
    const { body } = render(SourceSwitcher, {
      props: { busy: false, message: null, signedIn: false, onaddon: () => {}, onbuild: () => {}, onfight: () => {}, onsignin: () => {} },
    });
    expect(body).toContain('href="/addon"');
    expect(body).toContain(currentCharacterCopy.getTheAddon);
  });
});
```

```ts
// web/src/components/planner/ImportBox.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ImportBox from './ImportBox.svelte';
import { currentCharacterCopy } from '../../lib/current-character-copy';

describe('ImportBox', () => {
  it('links to /addon', () => {
    // talents/activeBuild props: reuse this component's own existing fixture pattern --
    // check web/src/lib/planner/load.test.ts or a Planner.svelte test for a minimal
    // TalentIndex fixture already in the repo before inventing one.
    const { body } = render(ImportBox, { props: { talents: FIXTURE_TALENT_INDEX, activeBuild: '1', onimport: () => {} } });
    expect(body).toContain('href="/addon"');
    expect(body).toContain(currentCharacterCopy.getTheAddon);
  });
});
```

- [ ] **Step 4: Run tests**

Run: `cd web && npx vitest run src/components/sim/SourceSwitcher.test.ts src/components/planner/ImportBox.test.ts`
Expected: PASS.

- [ ] **Step 5: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/sim/SourceSwitcher.svelte src/components/sim/SourceSwitcher.test.ts src/components/planner/ImportBox.svelte src/components/planner/ImportBox.test.ts`

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): a Get the addon link in every paste box\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-8
git add web/src/components/sim/SourceSwitcher.svelte web/src/components/sim/SourceSwitcher.test.ts web/src/components/planner/ImportBox.svelte web/src/components/planner/ImportBox.test.ts
git commit -F .superpowers/commit-msg-8
```

---

### Task 9: the below-60 framing line, and the honest "simming at 60" note

**Files:**
- Modify: `web/src/components/sim/ScopeNote.astro`
- Modify: `web/src/lib/sim/copy.ts` (two new `simCopy` keys)
- Modify: `web/src/components/sim/CharacterStrip.svelte`

**Interfaces:** none new (both are copy + render changes on existing components).

- [ ] **Step 1: Add the two copy keys**

```ts
// web/src/lib/sim/copy.ts, inside simCopy, beside the existing `scopeNote`:
belowSixtyFraming:
  'The simulator models level 60 characters. Below 60, plan your build in the planner and come back.',
// beside the existing levelSuffix-adjacent strings (openInPlanner, changeSource):
simmedAtSixty: 'Simmed as a level 60 with these talents.',
```

- [ ] **Step 2: `ScopeNote.astro`**

```astro
<p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-scope-note">
  {simCopy.scopeNote}
</p>
<p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-below-sixty-note">
  {simCopy.belowSixtyFraming}
</p>
```

This is static, server-rendered content (the file has no `<script>` block at all — it is a
pure Astro partial), so it costs nothing against the Lighthouse/CLS budget: it exists before
hydration and never moves.

- [ ] **Step 3: `CharacterStrip.svelte`'s honest note**

Change the existing `levelSuffix` derivation (currently only appends "· N talent points")
to also say the character is simmed at 60:

```ts
const levelSuffix = $derived(
  character.talent_level < SIM_LEVEL
    ? ` · ${character.point_order.length} talent points (${simCopy.simmedAtSixty})`
    : '',
);
```

- [ ] **Step 4: Write/extend tests**

`ScopeNote.astro` has no `.test.ts` (Astro partials in this repo are not unit-tested the same
way — confirm by checking whether any other `.astro` file under `src/components/sim` has a
sibling `.test.ts`; if none do, this step is e2e-only, covered by Task 12).

```ts
// Add to web/src/lib/sim/character.test.ts (existing file) or create a small assertion
// directly in a CharacterStrip test if one already exists; otherwise add
// web/src/components/sim/CharacterStrip.test.ts:
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterStrip from './CharacterStrip.svelte';
import { simCopy } from '../../lib/sim/copy';
// build a minimal SimCharacter fixture with point_order.length === 30 (well under 51) and
// talent_level accordingly < 60 -- reuse an existing fixture from character.test.ts if one
// fits, rather than hand-building a new one.

it('says the character is simmed at 60 when it has fewer than 51 points', () => {
  const { body } = render(CharacterStrip, { props: { character: underLeveled, items: new Map() } });
  expect(body).toContain(simCopy.simmedAtSixty);
});

it('says nothing extra for a full 51-point build', () => {
  const { body } = render(CharacterStrip, { props: { character: fullBuild, items: new Map() } });
  expect(body).not.toContain(simCopy.simmedAtSixty);
});
```

(`CharacterStrip.svelte` requires `onchange`/`onrace` callback props too — check the exact
required prop list from the file read in this plan's recon and supply no-op functions for
both.)

- [ ] **Step 5: Run tests**

Run: `cd web && npx vitest run src/components/sim/CharacterStrip.test.ts`
Expected: PASS.

- [ ] **Step 6: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/sim/ScopeNote.astro src/lib/sim/copy.ts src/components/sim/CharacterStrip.svelte src/components/sim/CharacterStrip.test.ts`

- [ ] **Step 7: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): honest below-60 and simmed-at-60 framing on the simulator\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-9
git add web/src/components/sim/ScopeNote.astro web/src/lib/sim/copy.ts web/src/components/sim/CharacterStrip.svelte web/src/components/sim/CharacterStrip.test.ts
git commit -F .superpowers/commit-msg-9
```

---

### Task 10: the planner writes the pointer on load, restores it on a bare load, and mounts the chip

**Files:**
- Modify: `web/src/components/planner/Planner.svelte`
- Modify: `web/src/lib/planner/current-character-planner.ts` (new small module: the
  Planner-side bridge, mirroring Task 3's `current-character-bridge.ts` for the sim side —
  kept separate because it depends on `BuildDraft`/`ClassRow` shapes the sim-side module has
  no reason to import)
- Test: `web/src/lib/planner/current-character-planner.test.ts`

**Interfaces:**
- Produces:
  ```ts
  // current-character-planner.ts
  export function recordPlannerCharacter(
    classSlug: string,
    code: string, // the fresh FS1 v2 code for what is now loaded
    storage?: Storage,
  ): void;
  export function isBarePlannerUrl(search: string): boolean;
  ```
- Consumes: `writeCurrent`, `readCurrent`, `CurrentCharacter` (Task 1); `CurrentCharacterChip`
  (Task 2); `decodeFS1` (existing, `lib/planner/fs1.ts`); `encodeFS1V2` (existing).

**Ruling:** the planner can only deterministically restore a bare-loaded pointer for
`'code'` and `'build'` sources — a `'code'` pointer decodes directly with the planner's
existing `decodeFS1`, a `'build'` pointer already has a permalink route
(`plannerHrefFor`/`/b/<id>`, not this page's own bare-load restore at all — a `'build'`
pointer's honest handling on `/planner` itself, not `/b/<id>`, is to do nothing, since
`/planner` has no `?build=` bootstrap and inventing one is out of this task's scope). So the
planner's own bare-load restore only ever fires for a `'code'`-sourced pointer. Document this
ruling in the ledger when this task lands.

- [ ] **Step 1: Write the failing tests**

```ts
// web/src/lib/planner/current-character-planner.test.ts
import { describe, expect, it } from 'vitest';
import { isBarePlannerUrl, recordPlannerCharacter } from './current-character-planner';
import { readCurrent } from '../current-character';

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: (index) => [...map.keys()][index] ?? null,
    get length() {
      return map.size;
    },
  };
}

describe('isBarePlannerUrl', () => {
  it('is true for an empty query', () => {
    expect(isBarePlannerUrl('')).toBe(true);
  });
  it('is false when ?code= is present', () => {
    expect(isBarePlannerUrl('?code=FS1:1:warrior:orc:0/0/0:')).toBe(false);
  });
  it('is false when ?class= is present', () => {
    expect(isBarePlannerUrl('?class=warrior')).toBe(false);
  });
});

describe('recordPlannerCharacter', () => {
  it('writes a code-sourced pointer', () => {
    const storage = fakeStorage();
    recordPlannerCharacter('warrior', 'FS1:1:warrior:orc:0/0/0:', storage);
    expect(readCurrent(storage)).toMatchObject({ source: 'code', classSlug: 'warrior' });
  });
});
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd web && npx vitest run src/lib/planner/current-character-planner.test.ts`
Expected: FAIL — module does not exist.

- [ ] **Step 3: Implement**

```ts
// web/src/lib/planner/current-character-planner.ts
// The planner side of the current-character bridge (Task 3's current-character-bridge.ts
// is the sim side; kept separate, see that file's own header for why).
import { writeCurrent, type CurrentCharacter } from '../current-character';

export function isBarePlannerUrl(search: string): boolean {
  const params = new URLSearchParams(search);
  return params.get('code') === null && params.get('class') === null && params.get('race') === null;
}

/** No spec label here: the planner has no SimCharacter, only a class slug and a fresh FS1
 *  code. The chip's label falls back to the class slug itself, title-cased, rather than
 *  guessing a spec off talent points the player has not necessarily finished spending. */
function titleCase(slug: string): string {
  return slug.replace(/(^|-)([a-z])/g, (_, sep: string, letter: string) => `${sep === '-' ? ' ' : ''}${letter.toUpperCase()}`);
}

export function recordPlannerCharacter(classSlug: string, code: string, storage?: Storage): void {
  const pointer: CurrentCharacter = {
    source: 'code',
    ref: code,
    label: titleCase(classSlug),
    classSlug,
    savedAt: new Date().toISOString(),
  };
  writeCurrent(pointer, storage);
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/planner/current-character-planner.test.ts`
Expected: PASS.

- [ ] **Step 5: Wire `Planner.svelte`**

Read the current `<script>` block in full (it changed once already in Task 1-9's own edits
to neighbouring files only if any touched `Planner.svelte` directly — none did, so the file
read during this plan's recon is still accurate). Two additions:

1. **Write on load.** After the `store` is created (`untrack(() => createPlannerStore(...))`),
   add an effect that writes the pointer whenever the store's own class/order/gear settle
   into something worth remembering — but only on an actual *load* event (import, code
   decode, or `record` prop), never on every keystroke-level talent edit (that would thrash
   localStorage on every click and is not what "on every successful load" means). The
   simplest correct trigger: call `recordPlannerCharacter` once, synchronously, right after
   each of the three load paths already in this file:
   - Inside `ImportBox`'s `onimport` handler (`Planner.svelte` passes one to
     `<ImportBox onimport={...}>` — find it and extend it): after the store applies the
     imported build, encode a fresh FS1 v2 code the same way `SharePanel.svelte`'s
     `addonCode` derivation does (`encodeFS1V2`/`codeForCharacterSpec` pattern) and call
     `recordPlannerCharacter(classSlug, code)`.
   - After a successful `?code=` decode (`decoded?.ok` branch already in this file): call
     `recordPlannerCharacter(decoded.build.classSlug, codeParam)` (the original query
     param is already the FS1 string — no re-encoding needed here).
   - When `record !== null` (the `/b/:id` mount): call
     `recordCurrentCharacter`-equivalent for a build source — actually simplest and most
     honest: `writeCurrent({ source: 'build', ref: record.id, label: record.title ?? classRow.name, classSlug: classRow.slug, savedAt: new Date().toISOString() })`,
     imported directly from `'../../lib/current-character'` (this one path needs the
     'build' shape `current-character-planner.ts`'s `recordPlannerCharacter` does not
     produce — write it inline rather than adding a second helper function for one call
     site).

2. **Restore on a bare load.** Near the top of the mount logic, before `decoded`/`codeParam`
   are computed, if `isBarePlannerUrl(window.location.search)` is true, read the stored
   pointer; if it exists and `source === 'code'`, treat its `ref` exactly as `codeParam`
   would have been (feed it through the same `decodeFS1` path already in the file) and set a
   `restored = $state(true)` flag.

3. **Mount the chip and the restored note.** In the template, above the existing
   `<ImportBox ...>` (or wherever the page's top-level layout begins — check the file's
   render block), add:

```svelte
{#if restored}
  <p class="text-muted text-[13px]" data-testid="planner-restored-note">
    {currentCharacterCopy.restoredNote}
    <button type="button" class="text-nav ml-1" onclick={() => { clearCurrent(); restored = false; }} data-testid="planner-restored-forget">
      {currentCharacterCopy.forget}
    </button>
  </p>
{/if}
```

`CurrentCharacterChip` itself is NOT mounted inline in the planner's own always-visible
paste-box area (the spec: "renders nothing on pages that have their own paste box") — do not
add it to `Planner.svelte`'s template at all. This task's chip-mounting scope is therefore
limited to the restored-note banner above; `CurrentCharacterChip.svelte` stays reserved for
Lane A's non-paste-box pages and any future planner surface that is not the paste box itself.

- [ ] **Step 6: Run the planner component/lib suite**

Run: `cd web && npx vitest run src/components/planner src/lib/planner`
Expected: PASS.

- [ ] **Step 7: Lint, type-check, prettier, file size**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/planner/Planner.svelte src/lib/planner/current-character-planner.ts src/lib/planner/current-character-planner.test.ts && wc -l src/components/planner/Planner.svelte`

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): the planner writes and restores the current-character pointer\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-10
git add web/src/components/planner/Planner.svelte web/src/lib/planner/current-character-planner.ts web/src/lib/planner/current-character-planner.test.ts
git commit -F .superpowers/commit-msg-10
```

---

### Task 11: Share asks before it writes

**Files:**
- Modify: `web/src/components/planner/SharePanel.svelte`
- Modify: `web/src/lib/planner/copy.ts` (new, small — first planner-specific copy module;
  every other planner component today reaches into `addonCopy`/`simCopy` for its strings,
  which is why this task creates the first dedicated one rather than adding to either)
- Test: `web/src/lib/planner/copy.test.ts` (trivial: the module exports the exact strings
  used below — matches the convention other `copy.ts` files in this repo do not bother
  testing directly, so keep this to one assertion per string, or skip the test file
  entirely and let `SharePanel.test.ts`, added below, be the only place these strings are
  asserted — prefer that: one fewer file, same coverage)

**Interfaces:** none new outside this component; `share()`'s existing signature and the
`saveBuild`/`cardUrlFor` calls it makes are unchanged, only gated behind a new confirm step.

- [ ] **Step 1: Add the copy module**

```ts
// web/src/lib/planner/copy.ts
// Planner-specific copy. New module (the planner's other components read addonCopy/simCopy
// for their strings today; this is the first string that is genuinely the planner's own).
export const plannerCopy = {
  shareConfirmTitle: 'Share this build',
  shareConfirmBody:
    'This saves your talents and gear to a link anyone can open. Nothing else about you is shared.',
  shareConfirmProceed: 'Share anyway',
  shareConfirmCancel: 'Cancel',
  shareConfirmCopyCode: 'Copy addon code instead',
  shareConfirmCopyLink: 'Copy an unsaved link instead',
  copiedUnsavedLink: 'Copied',
} as const;
```

- [ ] **Step 2: Write the failing component test**

```ts
// web/src/components/planner/SharePanel.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SharePanel from './SharePanel.svelte';
import { plannerCopy } from '../../lib/planner/copy';
// Build the minimal PlannerStore/LiveDps fixture this component needs -- check
// web/src/lib/planner/store.test.ts for an existing createPlannerStore fixture to reuse
// rather than hand-rolling a second one, and a LiveDps stub with state: 'idle'.

describe('SharePanel', () => {
  it('shows the confirm step, not a saved link, on first render of a spent build', () => {
    const { body } = render(SharePanel, { props: { store: fixtureStore, live: fixtureLive } });
    expect(body).not.toContain('data-testid="share-link"');
  });

  it('the Share button reads Share, not Saving, before anything is confirmed', () => {
    const { body } = render(SharePanel, { props: { store: fixtureStore, live: fixtureLive } });
    expect(body).toContain('>Share<');
  });
});
```

`svelte/server`'s `render()` cannot click a button and observe the resulting state (no
interactivity in server rendering), so this task's *behavioural* assertion — that clicking
"Share anyway" is what actually calls `saveBuild`, and that "Cancel" does not — is an e2e
concern (Task 12), not this component test. This component test only asserts the *initial*
render shows the confirm affordance rather than jumping straight to a save.

- [ ] **Step 3: Run test to verify it fails**

Run: `cd web && npx vitest run src/components/planner/SharePanel.test.ts`
Expected: whichever assertion the current unconfirmed code fails (today's `share()` runs
immediately on click, with no confirm step rendered at all before that — the test as written
may already pass on the "not Saving" assertion; the important failing state is Step 4's new
markup not existing yet, so add a third assertion that pins it):

```ts
it('renders the confirm panel body copy once the player asks to share', () => {
  // Server-side render cannot simulate a click; assert the confirm markup exists in the
  // component's output somewhere (Svelte renders both branches' possible content only when
  // reachable via a prop, not a click, in a server render -- if the confirm panel is gated
  // by local $state defaulting to false, this assertion instead checks the CANCELLED/initial
  // render does NOT already show the saved-link markup, which is the meaningful regression
  // guard at this test level). Replace with the two `not.toContain` assertions above if a
  // stateful confirm-then-click test does not fit svelte/server rendering -- confirmed by
  // this plan's recon that no existing test in this repo exercises a click-through Svelte
  // island this way; e2e is where "Share anyway" vs "Cancel" is actually exercised.
});
```

- [ ] **Step 4: Implement the confirm gate**

Add local state and a guarded `share()`:

```ts
let confirmOpen = $state(false);

function requestShare(): void {
  confirmOpen = true;
}

function cancelShare(): void {
  confirmOpen = false;
}

async function confirmedShare(): Promise<void> {
  confirmOpen = false;
  await share(); // the existing function, body unchanged
}
```

Change the existing Share `<button onclick={share}>` to `<button onclick={requestShare}>`
(keep its existing `disabled={saving || store.spent === 0}` guard). Add, directly below the
button row, gated on `confirmOpen`:

```svelte
{#if confirmOpen}
  <div class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-3" role="alertdialog" data-testid="share-confirm">
    <p class="text-strong text-[14px]">{plannerCopy.shareConfirmTitle}</p>
    <p class="text-muted text-[13px]">{plannerCopy.shareConfirmBody}</p>
    <div class="flex flex-wrap gap-3">
      <button type="button" class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4" onclick={() => void confirmedShare()} data-testid="share-confirm-proceed">
        {plannerCopy.shareConfirmProceed}
      </button>
      <button type="button" class={NEUTRAL_BUTTON} onclick={cancelShare} data-testid="share-confirm-cancel">
        {plannerCopy.shareConfirmCancel}
      </button>
      <button type="button" class={NEUTRAL_BUTTON} disabled={addonCode === ''} onclick={() => { copyToClipboard(addonCode, 'addon'); cancelShare(); }} data-testid="share-confirm-copy-code">
        {plannerCopy.shareConfirmCopyCode}
      </button>
      <button
        type="button"
        class={NEUTRAL_BUTTON}
        onclick={() => {
          const unsavedLink = `${window.location.origin}/planner?code=${encodeURIComponent(addonCode === '' ? '' : addonCode)}`;
          void copyToClipboard(unsavedLink, 'link');
          cancelShare();
        }}
        data-testid="share-confirm-copy-unsaved"
      >
        {plannerCopy.shareConfirmCopyLink}
      </button>
    </div>
  </div>
{/if}
```

Note: the "unsaved link" is a `?code=` link built from `addonCode` — but `addonCode` (this
file's existing `$derived`) is the in-game ADDON paste string (`addonCodeFor`), not
necessarily the URL-safe FS1 code `?code=` expects. Re-derive the unsaved link correctly:
reuse `codeForCharacterSpec`/`toCharacterSpec`/`characterFromPlanner` (already imported in
this file for the card-sim feature) to build a real FS1 code, the same way
`ComboResults.svelte`'s own `plannerHref` does:

```ts
const unsavedPlannerLink = $derived.by(() => {
  const character = characterFromPlanner(store);
  if (character === null || store.talentIndex === null) return '';
  const spec = toCharacterSpec(character, store.talentIndex, [], []);
  return `${window.location.origin}/planner?code=${encodeURIComponent(codeForCharacterSpec(spec, store.treeVersion))}`;
});
```

(`window.location.origin` is only safe inside the click handler or an `onMount`-guarded
derivation, never evaluated during SSR — since `SharePanel.svelte` is always inside the
`client:load` `Planner` island, this is fine, but guard with `typeof window !== 'undefined'`
defensively, following this file's own existing `fromQuery`-style guard pattern elsewhere in
the planner components.) Use `unsavedPlannerLink` in the button's `onclick` instead of the
ad-hoc string built above.

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd web && npx vitest run src/components/planner/SharePanel.test.ts`
Expected: PASS.

- [ ] **Step 6: Full planner component/lib suite**

Run: `cd web && npx vitest run src/components/planner src/lib/planner`

- [ ] **Step 7: Lint, type-check, prettier**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/planner/SharePanel.svelte src/components/planner/SharePanel.test.ts src/lib/planner/copy.ts`

- [ ] **Step 8: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'feat(web): Share asks before it writes a build to a public link\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-11
git add web/src/components/planner/SharePanel.svelte web/src/components/planner/SharePanel.test.ts web/src/lib/planner/copy.ts
git commit -F .superpowers/commit-msg-11
```

---

### Task 12: e2e coverage for the new hand-offs

**Files:**
- Create: `web/tests/e2e/current-character.spec.ts`

**Interfaces:** none new; exercises everything Tasks 1-11 built, end to end.

- [ ] **Step 1: Read one existing sim/planner e2e spec first**

Run: `ls web/tests/e2e` and open one spec that already visits `/sim` and `/planner` (grep for
`page.goto('/sim'` across `web/tests/e2e`) to match this repo's existing fixture-data
assumptions, `data-testid` conventions, and `test.describe`/desktop-mobile project structure
before writing anything new.

- [ ] **Step 2: Write the spec**

```ts
// web/tests/e2e/current-character.spec.ts
import { expect, test } from '@playwright/test';

// A known-good FS1 v2 code for the fixture data build's Warrior tree, at a partial spend
// (well under 51 points) so the "simmed at 60" note is exercised too. Confirm the exact
// string against an existing e2e spec that already pastes one (grep tests/e2e for `FS1:`)
// and reuse that spec's own fixture string rather than inventing a new one that may not
// decode against the fixture data build.
const FIXTURE_CODE = 'FS1:<fixture-build>:warrior:orc:0/0/0:'; // replace with the real fixture string found in Step 1

test.describe('current character', () => {
  test('loading a character on /sim and then opening /planner bare restores it', async ({ page }) => {
    await page.goto(`/sim?code=${encodeURIComponent(FIXTURE_CODE)}`);
    await expect(page.getByTestId('sim-character')).toBeVisible();
    await page.goto('/planner');
    await expect(page.getByTestId('planner-restored-note')).toBeVisible();
  });

  test('Forget on the restored note clears the pointer for the next bare load', async ({ page }) => {
    await page.goto(`/sim?code=${encodeURIComponent(FIXTURE_CODE)}`);
    await page.goto('/sim');
    await expect(page.getByTestId('sim-restored-note')).toBeVisible();
    await page.getByTestId('sim-restored-forget').click();
    await page.goto('/sim');
    await expect(page.getByTestId('sim-restored-note')).toHaveCount(0);
  });

  test('every paste box links to /addon', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.locator('a[href="/addon"]')).toBeVisible();
    await page.goto('/planner');
    await expect(page.locator('a[href="/addon"]')).toBeVisible();
  });

  test('Droptimizer\'s every-upgrade rows offer Plan it', async ({ page }) => {
    await page.goto(`/sim/drops?code=${encodeURIComponent(FIXTURE_CODE)}`);
    // Drive the fixture source picker/run bar the same way an existing sim/drops e2e spec
    // already does (grep tests/e2e for an existing droptimizer spec and copy its run steps)
    // ...
    await expect(page.locator('a[data-testid^="sim-drops-plan-it-"]').first()).toBeVisible();
  });

  test('the below-60 framing line is on /sim', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('sim-below-sixty-note')).toBeVisible();
  });

  test('Share shows a confirm step before writing', async ({ page }) => {
    await page.goto('/planner');
    // spend at least one point so Share is enabled -- copy an existing planner e2e spec's
    // own "spend a point" steps here.
    await page.getByRole('button', { name: 'Share' }).click();
    await expect(page.getByTestId('share-confirm')).toBeVisible();
    await expect(page.getByTestId('share-link')).toHaveCount(0);
  });
});
```

Fill in every `...` placeholder above with real steps copied from an existing spec before
this task is considered done — the plan's own "no placeholders" rule applies to this file
exactly as much as to any `.ts` module; the ellipses here exist only because the exact
existing-spec helper calls were not re-read line-for-line while writing this plan, and the
implementer must read them and inline the real calls.

- [ ] **Step 3: Run the desktop project**

Run: `cd web && E2E_PORT=4411 npx playwright test tests/e2e/current-character.spec.ts --project=desktop`
Expected: PASS. Stop a stale preview first if needed: `npx astro preview stop`.

- [ ] **Step 4: Run the mobile project**

Run: `cd web && E2E_PORT=4411 npx playwright test tests/e2e/current-character.spec.ts --project=mobile`
Expected: PASS.

- [ ] **Step 5: Lint, prettier**

Run: `cd web && npm run lint && npx prettier --check tests/e2e/current-character.spec.ts`

- [ ] **Step 6: Commit**

```bash
cd /Users/jh/code/forever/.worktrees/op-character
printf 'test(web): e2e coverage for the current-character hand-offs\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-12
git add web/tests/e2e/current-character.spec.ts
git commit -F .superpowers/commit-msg-12
```

---

### Task 13: whole-branch review and fix wave

Not a build task — the final gate `subagent-driven-development` runs after Task 12: a fresh
reviewer reads the full diff since this plan's first commit against spec section 1 and
section 0's Global Constraints, checks every bullet in the spec's "Lane B also owns" list has
a landed task, and checks for the specific trap named in this lane's dispatch (a second,
missed instance of a control a fix only reached once — re-run the `textarea`/paste-box grep
from Task 8 one more time against the FINAL tree, and re-check every "Open in planner"/"Plan
it" surface named in spec section 1 has the link, not just the one this plan happened to
touch first). One fix wave, ledger entries for every ruling, then the lane's final report per
`lane-common.md`.

- [ ] **Step 1: Run the full scoped check suite**

```bash
cd /Users/jh/code/forever/.worktrees/op-character/web
export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12
FOREVER_DATA=fixture npm run sync
npx vitest run src/lib/current-character.test.ts src/lib/sim src/lib/planner src/components/sim src/components/planner src/components/CurrentCharacterChip.test.ts
npx astro check
npm run lint
npx prettier --check src/lib/current-character.ts src/lib/current-character-copy.ts src/components/CurrentCharacterChip.svelte src/lib/sim/sources.ts src/lib/sim/bulk-store.svelte.ts src/lib/sim/tabs.ts src/lib/sim/combos.ts src/lib/sim/copy.ts src/components/sim/SimView.svelte src/components/sim/tools/ComboResults.svelte src/components/sim/tools/DropResults.svelte src/components/sim/SourceSwitcher.svelte src/components/sim/CharacterStrip.svelte src/components/sim/ScopeNote.astro src/components/planner/Planner.svelte src/components/planner/SharePanel.svelte src/components/planner/ImportBox.svelte src/lib/planner/copy.ts src/lib/planner/current-character-planner.ts
npx astro preview stop || true
E2E_PORT=4411 npx playwright test tests/e2e/current-character.spec.ts
```

- [ ] **Step 2: Spec-coverage pass**

Check off, against the final tree, every clause of spec section 1: pointer shape and module
signatures (Task 1); chip (Task 2); pointer written on every successful load (Task 3);
restored on bare `/sim`/`/planner` (Tasks 5, 10); tools island bootstraps from `?code=` and
the pointer (Task 4); tab strip carries the character (already true via `syncTabHrefs`,
unaffected by this plan — confirm no task broke it); "Get the addon" in every paste box
(Task 8); "Open in planner" on Quick Sim (already true via `CharacterStrip.svelte`'s existing
`plannerHrefFor` link — confirm still present and not accidentally duplicated by Task 5's
restored-note wiring); "Plan it" on Droptimizer/Top Gear rows (Task 7); below-60 framing
(Task 9); "simmed at 60" honesty (Task 9); Share confirm (Task 11); `fromLoggedFight`'s third
ref part (Task 6); `/sim/drops`'s `?instance=` (Task 4).

- [ ] **Step 3: Fix wave**

File-scoped fixes only, same TDD cycle as any task above; no new scope.

- [ ] **Step 4: Ledger and final report**

Write `.superpowers/sdd/2026-09-21-op-character/progress.md` with every ruling made across
Tasks 1-13 as `Ruling: <decision> — <why> — <cost if wrong>` (the plan's own "Ruling"
callouts in Tasks 1 and 10 are the seed list; add every other judgment call an implementer or
reviewer made along the way). Then produce the lane's final report per `lane-common.md`:
branch head sha; what shipped per spec bullet; test counts from the final run above; every
ruling; anything left undone and why (name it explicitly if, e.g., the planner's
`?source=&ref=` bootstrap for `'fight'`/`'armory'` pointers was ruled out of scope in Task
10 — that is a real, intentional gap, not an oversight, and the report should say so in one
sentence); files another lane must know this lane touched (nothing outside this lane's
ownership list should appear in `git diff --stat` — verify that as the report's last line).
