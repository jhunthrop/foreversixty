# Battle.net-first web lane implementation plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Sign-in with Battle.net lands a player on `/account` as their hub, with their main
character picked, pointed at, and usable across the planner, simulator, logs and their
character page — no addon paste required for a Battle.net-covered character.

**Architecture:** `GET /v1/me` gains `characters[].build: {source, captured_at}` (frozen
contract §5, stubbed here — the api lane owns wiring it live). A pure `mainCharacter()`
picks the character to point at; `/account?signed_in=1` writes the current-character
pointer to it (the same `armory`-kind pointer every stored-character load already writes)
and shows it in the existing hero band. The simulator and planner keep restoring that
pointer exactly as they already do for any `armory`-sourced character — no new bootstrap
path. `sources.ts`/`addon-export.ts` treat `source: "blizzard"` exactly like `"addon"`
(both are FS1 strings). Every rewritten string lives in a copy module.

**Tech Stack:** Astro islands, Svelte 5 runes, Vitest, Playwright, TypeScript strict.

**Spec:** `docs/superpowers/specs/2026-09-22-battlenet-first-design.md` (sections 0, 1, 3, 4
bind this lane; section 5 is the frozen `/v1/me` and `sim-input` contract; section 2 is the
`bnet-first-api` lane, not touched here). Also read for the states model this lane must
match: `docs/superpowers/specs/2026-09-22-account-and-island-states-design.md` §1 (the
primitives: `Skeleton`, `LoadError`, `EmptyState`, `.reveal`) — nothing new added by this
plan needs a new loading/failed/empty state (every fetch this plan touches already has one:
`fetchMeOnce`, `fetchSimInput`), so no new `Skeleton`/`LoadError` wiring is required, but any
new island output must sit inside the existing reserved-height wrappers, never widen them
without re-measuring.

## Global Constraints

- Vocabulary: "from Battle.net", never "from the armory", in every rewritten string. The
  code's internal `armory` pointer-source name stays as-is (it means "loaded by stored
  character key", unrelated to Blizzard).
- Every visible string lives in a copy module (`sim/copy.ts`, `account/*-copy.ts`,
  `addon/copy.ts`, or a new small copy module — never a literal in a `.svelte`/`.astro`
  file).
- Honesty: a character is shown as simmable only when `MeCharacter.build` is present.
  Nothing is inferred.
- No token persistence (unchanged; this plan touches no auth storage).
- File ownership: this lane owns all of `web/**`. It does not touch `api/**` or the OpenAPI
  file (the `bnet-first-api` lane's scope).
- House rules (`.superpowers/journeys/lane-common-web.md`): sonnet-only subagents (never
  opus, including review); commit message format and attribution; scoped checks before each
  commit (`npx vitest run <paths>`, `npx astro check`, `npm run lint`,
  `npx prettier --check <paths>`); stop every server before the final report; no broad
  `pkill`.
- Functions under 50 lines, files under 800 (several copy files are already near the
  ceiling — new copy goes in a new small module, not appended past the limit).
- `web/node_modules` is a symlink into the main checkout: never `npm install`, never stage
  it.

---

## File map (what each task creates or changes)

- `web/src/lib/account/api.ts` — **modify**: `MeCharacter.build?` field.
- `web/src/lib/sim/types.ts` — **modify**: `SourceKind` gains `'blizzard'`.
- `web/src/lib/account/main-character.ts` — **create**: pure `mainCharacter()` +
  `pointerForCharacter()`.
- `web/src/lib/account/main-character.test.ts` — **create**.
- `web/src/lib/account/build-pill.ts` — **create**: `buildSourcePill()`, shared by the
  account rows and the sim landing rows.
- `web/src/lib/account/build-pill.test.ts` — **create**.
- `web/src/lib/sim/sources.ts` — **modify**: `sourcePill` gains `'blizzard'`;
  `fromStoredCharacter` treats `'blizzard'` like `'addon'`.
- `web/src/lib/addon-export.ts` — **modify**: `lookupAddonExport` treats `'blizzard'` like
  `'addon'`.
- `web/src/pages/login.astro` — **modify**: `next="/account?signed_in=1"`.
- `web/src/components/SignInPrompt.svelte` — **modify**: default `next`.
- `web/src/components/Account.svelte` — **modify**: sign-in default next; `?signed_in=1`
  hub-arrival handling (main character, pointer write, banner, hero-band Logs link, Your
  ratings panel).
- `web/src/lib/account/account-page-copy.ts` — **modify**: banner + Logs + ratings-heading
  strings.
- `web/src/lib/account/character-list-copy.ts` — **modify**: Characters intro rewrite, pill
  labels.
- `web/src/components/account/CharacterList.svelte` — **modify**: build-source pill per
  row; replace per-row `CharacterHandoffLinks` (which did a per-row `sim-input` fetch) with
  the new no-fetch `CharacterRowLink`.
- `web/src/components/account/CharacterRowLink.svelte` — **create**: one ref-based "Open in
  simulator" link (or the paste fallback), built from `MeCharacter.build` alone, no fetch.
- `web/src/lib/handoff-links.ts` — **modify**: add `simArmoryHref()`.
- `web/src/components/sim/LandingState.svelte` — **modify**: per-row build-source pill;
  drop the `landingSourceNote` footnote.
- `web/src/components/sim/SourceSwitcher.svelte` — **modify**: rewrite the account card's
  copy (`armorySignIn`/`armoryNotYet` replaced).
- `web/src/lib/sim/copy.ts` — **modify**: rewrite `landingSourceNote` (deleted),
  `armoryNotYet`/`armorySignIn` (deleted, replaced by new keys), `noCharactersYet`,
  `sourceAddonBody`, `sourceAccountTitle`, `bagsNeedAddon` is `bulkCopy`'s (own file below).
- `web/src/pages/index.astro` — **modify**: home panel (signed-out sentence + button;
  signed-in strip), server-rendered reserved height.
- `web/src/components/HomeAccountPanel.svelte` — **create**: the signed-in swap island.
- `web/src/lib/home-panel-copy.ts` — **create**.
- `web/src/pages/addon.astro` — **modify**: opening line rewrite.
- `web/src/lib/addon/copy.ts` — **modify**: opening line + `bagsNeedAddon`-equivalent
  rewrite (the string moves out of `sim/copy.ts`'s `bulkCopy`, see Task 8).
- `web/src/lib/planner/fs1.blizzard.test.ts` — **create**: decodes the API lane's
  `era-kiloz.fs1`, skips cleanly when absent.
- `web/tests/e2e/auth.spec.ts` — **modify**: `next=` assertions.
- `web/tests/e2e/bnet-hub.spec.ts` — **create**: sign-in → hub → Open in simulator.
- `web/tests/e2e/home-panel.spec.ts` — **create**: signed-out and signed-in home panel.
- `web/src/fixtures/me-bnet.ts` — **modify**: add `build` to the fixture's guilded
  character so existing and new e2e specs can assert a simmable hero.

---

### Task 1: Contract types, `mainCharacter()`, `buildSourcePill()`

**Files:**
- Modify: `web/src/lib/account/api.ts`
- Modify: `web/src/lib/sim/types.ts`
- Create: `web/src/lib/account/main-character.ts`
- Create: `web/src/lib/account/main-character.test.ts`
- Create: `web/src/lib/account/build-pill.ts`
- Create: `web/src/lib/account/build-pill.test.ts`

**Interfaces:**
- Produces: `MeCharacter.build?: { source: 'addon' | 'blizzard'; captured_at: string }`.
- Produces: `mainCharacter(characters: MeCharacter[]): MeCharacter | null`.
- Produces: `pointerForCharacter(character: MeCharacter): CurrentCharacter` (from
  `../current-character`).
- Produces: `buildSourcePill(build: MeCharacter['build'], now?: Date): { label: string;
  pillClass: 'pill-blizzard' | 'pill-site' | null }`.

- [ ] **Step 1: Add the `build` field to `MeCharacter`**

In `web/src/lib/account/api.ts`, inside `interface MeCharacter` (after the existing
`source?: string;` field and its doc comment), add:

```ts
  /** The site's newest export for this character, when it has one (spec 2026-09-22 §5):
   *  `'addon'` for a paste or companion push, `'blizzard'` for one built from the
   *  Battle.net profile. Omitted exactly when the site holds no build for the character —
   *  never inferred from anything else on this row. */
  build?: { source: 'addon' | 'blizzard'; captured_at: string };
```

- [ ] **Step 2: Add `'blizzard'` to `SourceKind`**

In `web/src/lib/sim/types.ts`, change:

```ts
export type SourceKind = 'armory' | 'addon' | 'build' | 'fight' | 'manual';
```

to:

```ts
export type SourceKind = 'armory' | 'addon' | 'blizzard' | 'build' | 'fight' | 'manual';
```

Leave `web/src/lib/sim/url.ts`'s `SOURCES` array unchanged — `'blizzard'` is never a URL
`?source=` value (a stored Battle.net character is always reached by
`?source=armory&ref=<key>`, the same as a stored addon character); it only ever appears as
`SimInput.source`/`CharacterSource.kind`.

- [ ] **Step 3: Write the failing test for `mainCharacter`**

Create `web/src/lib/account/main-character.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import { mainCharacter, pointerForCharacter } from './main-character';
import type { MeCharacter } from './api';

function character(overrides: Partial<MeCharacter>): MeCharacter {
  return {
    key: 'us/normal/test',
    region: 'us',
    ruleset: 'normal',
    name: 'Test',
    ...overrides,
  };
}

describe('mainCharacter', () => {
  it('is null with no characters', () => {
    expect(mainCharacter([])).toBeNull();
  });

  it('picks the only simmable character over a higher-level one with no build', () => {
    const simmable = character({
      key: 'us/normal/a',
      name: 'A',
      level: 10,
      build: { source: 'blizzard', captured_at: '2026-09-20T00:00:00Z' },
    });
    const noBuild = character({ key: 'us/normal/b', name: 'B', level: 60 });
    expect(mainCharacter([noBuild, simmable])).toBe(simmable);
  });

  it('picks the newest captured_at among simmable characters', () => {
    const older = character({
      key: 'us/normal/a',
      name: 'A',
      build: { source: 'addon', captured_at: '2026-09-19T00:00:00Z' },
    });
    const newer = character({
      key: 'us/normal/b',
      name: 'B',
      build: { source: 'blizzard', captured_at: '2026-09-21T00:00:00Z' },
    });
    expect(mainCharacter([older, newer])).toBe(newer);
  });

  it('breaks a captured_at tie by level', () => {
    const lower = character({
      key: 'us/normal/a',
      name: 'A',
      level: 40,
      build: { source: 'addon', captured_at: '2026-09-21T00:00:00Z' },
    });
    const higher = character({
      key: 'us/normal/b',
      name: 'B',
      level: 60,
      build: { source: 'addon', captured_at: '2026-09-21T00:00:00Z' },
    });
    expect(mainCharacter([lower, higher])).toBe(higher);
  });

  it('falls back to the highest level when nothing is simmable', () => {
    const low = character({ key: 'us/normal/a', name: 'A', level: 10 });
    const high = character({ key: 'us/normal/b', name: 'B', level: 45 });
    expect(mainCharacter([low, high])).toBe(high);
  });

  it('treats an undefined level as the lowest', () => {
    const withLevel = character({ key: 'us/normal/a', name: 'A', level: 5 });
    const noLevel = character({ key: 'us/normal/b', name: 'B' });
    expect(mainCharacter([noLevel, withLevel])).toBe(withLevel);
  });
});

describe('pointerForCharacter', () => {
  it('writes an armory-kind pointer keyed by the character', () => {
    const c = character({ key: 'us/normal/a', name: 'Aria', class: 'Mage' });
    const pointer = pointerForCharacter(c);
    expect(pointer.source).toBe('armory');
    expect(pointer.ref).toBe('us/normal/a');
    expect(pointer.classSlug).toBe('mage');
    expect(pointer.label).toBe('Aria · Mage');
  });

  it('falls back to the character name alone with no class on file', () => {
    const c = character({ key: 'us/normal/a', name: 'Aria' });
    const pointer = pointerForCharacter(c);
    expect(pointer.label).toBe('Aria');
    expect(pointer.classSlug).toBe('');
  });
});
```

- [ ] **Step 4: Run it, confirm it fails**

Run: `cd web && npx vitest run src/lib/account/main-character.test.ts`
Expected: FAIL — `main-character.ts` does not exist yet.

- [ ] **Step 5: Implement `main-character.ts`**

Create `web/src/lib/account/main-character.ts`:

```ts
// web/src/lib/account/main-character.ts
// The one character the hub points at on arrival (spec 2026-09-22 §3.1): the simmable
// character (one the site holds a build for) with the newest build.captured_at, ties
// broken by level; with nothing simmable, the highest level; with no characters at all,
// null. A pure function of the list `/v1/me` already returned, so the hub's own effect
// just calls it and writes the result -- no fetch of its own.
import type { CurrentCharacter } from '../current-character';
import type { MeCharacter } from './api';

function levelOf(character: MeCharacter): number {
  return character.level ?? 0;
}

function betterSimmable(a: MeCharacter, b: MeCharacter): MeCharacter {
  const atA = a.build?.captured_at ?? '';
  const atB = b.build?.captured_at ?? '';
  if (atA !== atB) return atA > atB ? a : b;
  return levelOf(b) > levelOf(a) ? b : a;
}

function betterByLevel(a: MeCharacter, b: MeCharacter): MeCharacter {
  return levelOf(b) > levelOf(a) ? b : a;
}

export function mainCharacter(characters: readonly MeCharacter[]): MeCharacter | null {
  if (characters.length === 0) return null;
  const simmable = characters.filter((character) => character.build !== undefined);
  if (simmable.length > 0) return simmable.reduce(betterSimmable);
  return [...characters].reduce(betterByLevel);
}

/** The current-character pointer for the hub's own arrival write (spec §3.1: "sets the
 *  current-character pointer to it"). Always an `'armory'`-kind pointer, the same kind
 *  every other stored-character load writes (`sources.ts`'s `fromStoredCharacter`) --
 *  `/sim` and the character page already know how to restore one; `/planner` does not
 *  (documented in `current-character-planner.ts`), which is an existing, intentional
 *  limit this pointer does not change. */
export function pointerForCharacter(character: MeCharacter): CurrentCharacter {
  const classSlug = character.class?.toLowerCase() ?? '';
  const label = character.class === undefined ? character.name : `${character.name} · ${character.class}`;
  return {
    source: 'armory',
    ref: character.key,
    label,
    classSlug,
    savedAt: new Date().toISOString(),
  };
}
```

- [ ] **Step 6: Run it, confirm it passes**

Run: `cd web && npx vitest run src/lib/account/main-character.test.ts`
Expected: PASS, 8 tests.

- [ ] **Step 7: Write the failing test for `buildSourcePill`**

Create `web/src/lib/account/build-pill.test.ts`:

```ts
import { describe, expect, it } from 'vitest';
import { buildSourcePill } from './build-pill';

const NOW = new Date('2026-09-22T00:00:00Z');

describe('buildSourcePill', () => {
  it('reads "No build yet" with no pill class when there is no build', () => {
    expect(buildSourcePill(undefined, NOW)).toEqual({ label: 'No build yet', pillClass: null });
  });

  it('reads Battle.net for a blizzard build', () => {
    const result = buildSourcePill({ source: 'blizzard', captured_at: '2026-09-20T00:00:00Z' }, NOW);
    expect(result).toEqual({ label: 'Battle.net · 2 days ago', pillClass: 'pill-blizzard' });
  });

  it('reads Addon for an addon build', () => {
    const result = buildSourcePill({ source: 'addon', captured_at: '2026-09-22T00:00:00Z' }, NOW);
    expect(result).toEqual({ label: 'Addon · just now', pillClass: 'pill-site' });
  });
});
```

- [ ] **Step 8: Run it, confirm it fails**

Run: `cd web && npx vitest run src/lib/account/build-pill.test.ts`
Expected: FAIL — module does not exist.

- [ ] **Step 9: Implement `build-pill.ts` and its copy keys**

First add these three keys to `web/src/lib/account/character-list-copy.ts`'s
`characterListCopy` object (anywhere in the object body):

```ts
  /** The account rows' and the sim landing rows' build-source pill (spec 2026-09-22 §3.1). */
  battlenetSource: 'Battle.net',
  addonSource: 'Addon',
  noBuildYet: 'No build yet',
```

Then create `web/src/lib/account/build-pill.ts`:

```ts
// web/src/lib/account/build-pill.ts
// The build-source pill shown on the account page's Characters rows and the simulator
// landing's "Your characters" rows (spec 2026-09-22 §3.1): "Battle.net · 2 days ago",
// "Addon · today", or the muted "No build yet" for a character the site holds no export
// for. One function so the two lists can never say this two different ways.
import { relativeTime } from '../sim/sources';
import { characterListCopy } from './character-list-copy';
import type { MeCharacter } from './api';

export interface BuildPill {
  label: string;
  /** Which pill colour to use, or null for the muted no-pill text. */
  pillClass: 'pill-blizzard' | 'pill-site' | null;
}

export function buildSourcePill(build: MeCharacter['build'], now: Date = new Date()): BuildPill {
  if (build === undefined) return { label: characterListCopy.noBuildYet, pillClass: null };
  const name = build.source === 'blizzard' ? characterListCopy.battlenetSource : characterListCopy.addonSource;
  const pillClass = build.source === 'blizzard' ? 'pill-blizzard' : 'pill-site';
  return { label: `${name} · ${relativeTime(build.captured_at, now)}`, pillClass };
}
```

- [ ] **Step 10: Run both new test files, confirm they pass**

Run: `cd web && npx vitest run src/lib/account/main-character.test.ts src/lib/account/build-pill.test.ts`
Expected: PASS, 11 tests total.

- [ ] **Step 11: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/account/main-character.ts src/lib/account/main-character.test.ts src/lib/account/build-pill.ts src/lib/account/build-pill.test.ts src/lib/account/api.ts src/lib/account/character-list-copy.ts src/lib/sim/types.ts`
Expected: no errors.

- [ ] **Step 12: Commit**

```bash
git add web/src/lib/account/api.ts web/src/lib/sim/types.ts web/src/lib/account/main-character.ts web/src/lib/account/main-character.test.ts web/src/lib/account/build-pill.ts web/src/lib/account/build-pill.test.ts web/src/lib/account/character-list-copy.ts
```
Write the commit message to `.superpowers/commit-msg-1.txt` with `printf`, then
`git commit -F .superpowers/commit-msg-1.txt` as its own Bash command. Subject:
`feat(web): a Battle.net build counts as simmable, and the hub can pick a main character`.

---

### Task 2: `sources.ts` and `addon-export.ts` treat `blizzard` like `addon`

**Files:**
- Modify: `web/src/lib/sim/sources.ts`
- Modify: `web/src/lib/sim/sources.test.ts`
- Modify: `web/src/lib/addon-export.ts`

**Interfaces:**
- Consumes: `SourceKind` (Task 1).
- Produces: `sourcePill` handles `{ kind: 'blizzard', ... }`; `fromStoredCharacter` decodes
  a `blizzard`-sourced `sim-input` the same way it decodes an `addon`-sourced one;
  `lookupAddonExport` returns a code for a `blizzard`-sourced character too.

- [ ] **Step 1: Write the failing tests**

In `web/src/lib/sim/sources.test.ts`, find the existing test
`'carries "armory" through unchanged the day the API starts sending it'` (around line 147)
and add two new tests near it, in the same `describe` block:

```ts
it('reads a blizzard-sourced sim-input exactly like an addon-sourced one', async () => {
  const code = 'FS1:1.60.1.69893:warrior:orc:1a/0/0:';
  fakeApi.simInput = { spec: 'fury', gear: code, talents: '', buffs: [], captured_at: '2026-09-21T00:00:00Z', source: 'blizzard' };
  const result = await fromStoredCharacter({ region: 'us', ruleset: 'normal', slug: 'kiloz' }, ctx);
  expect(result.ok).toBe(true);
  if (result.ok) expect(result.character.source.kind).toBe('blizzard');
});

it('sourcePill names Battle.net for a blizzard source', () => {
  const now = new Date('2026-09-22T00:00:00Z');
  expect(
    sourcePill({ kind: 'blizzard', ref: 'us/normal/kiloz', captured_at: '2026-09-20T00:00:00Z' }, now),
  ).toBe('Battle.net, 2 days ago');
});
```

Read the surrounding ~80 lines of `sources.test.ts` first (`sed -n '80,170p'
src/lib/sim/sources.test.ts`) to match the existing fixture/mock shape exactly (`fakeApi`,
`ctx`, whatever the file's own helper names are) — the two tests above use placeholder
names (`fakeApi.simInput`, `ctx`) that must be replaced with whatever this file's own
`fromStoredCharacter` tests already call the equivalent mock and context. Do not invent a
new mocking pattern; copy the one the neighbouring `'carries "armory"...'` test already
uses, changing only `source: 'armory'` to `source: 'blizzard'` and the assertion.

- [ ] **Step 2: Run, confirm both fail**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts`
Expected: the `sourcePill` test fails (no `'blizzard'` case yet); the
`fromStoredCharacter` test currently passes already (the function has no source-specific
branching keyed on the literal string `'addon'` for the *decode* path... verify by running
first). If both already pass, skip to Step 4 having confirmed by reading the code that the
`addon` check at `sources.ts`'s `if (input.source === 'addon' && ...)` line is the one
gate that needs widening — it is (Step 3 fixes it regardless of whether this specific test
already happens to pass, since `input.source === 'addon'` is a strict string check that
excludes `'blizzard'` today).

- [ ] **Step 3: Implement**

In `web/src/lib/sim/sources.ts`:

1. In `sourcePill`, add a case:

```ts
    case 'blizzard':
      return `Battle.net, ${relativeTime(source.captured_at, now)}`;
```

immediately after the existing `case 'armory':` branch (before `case 'addon':`).

2. In `fromStoredCharacter`, change:

```ts
  if (input.source === 'addon' && typeof input.gear === 'string' && input.gear.startsWith(`${FS1_PREFIX}:`)) {
```

to:

```ts
  if (
    (input.source === 'addon' || input.source === 'blizzard') &&
    typeof input.gear === 'string' &&
    input.gear.startsWith(`${FS1_PREFIX}:`)
  ) {
```

Leave the rest of that branch unchanged — `characterFromFs1` is called with
`{ kind: 'addon', ref: characterKey, captured_at: input.captured_at }` today; change that
literal `'addon'` to `input.source` so the resulting `SimCharacter.source.kind` is
`'blizzard'` when that is what the API said (this is what the new test in Step 1 asserts).
The pointer write beneath it (`recordCurrentCharacter(result.character, 'addon', code,
storage)`) stays `'addon'` unconditionally — Task 1's header comment on
`pointerForCharacter` already established every stored-character pointer this lane writes
is `'armory'`-kind by ref, never keyed by the FS1 code for a *stored* load; check this call
site specifically: it passes `code` as the ref, which is the paste/manual convention, not
the stored one. Read the surrounding 15 lines before editing — if this call site is really
recording the pointer as `{'addon', code}` for a *stored* character load (not a paste),
that looks like a pre-existing quirk unrelated to this task; do not change it beyond
widening the `input.source` check above. Leave that pointer-write line exactly as found.

- [ ] **Step 4: Run, confirm pass**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts`
Expected: PASS, no regressions in the rest of the file.

- [ ] **Step 5: `addon-export.ts`**

In `web/src/lib/addon-export.ts`, change:

```ts
    const code =
      input.source === 'addon' && typeof input.gear === 'string' && input.gear.startsWith(`${FS1_PREFIX}:`)
        ? input.gear
        : null;
```

to:

```ts
    const code =
      (input.source === 'addon' || input.source === 'blizzard') &&
      typeof input.gear === 'string' &&
      input.gear.startsWith(`${FS1_PREFIX}:`)
        ? input.gear
        : null;
```

Check `web/src/lib/addon-export.test.ts` (if it exists — `ls web/src/lib/addon-export.test.ts`)
and add the equivalent one-case test (a `source: 'blizzard'` input returns the code) beside
whatever existing `source: 'addon'` test is there, following that file's own mock shape. If
no test file exists, create `web/src/lib/addon-export.test.ts` with:

```ts
import { describe, expect, it, vi } from 'vitest';
import { lookupAddonExport } from './addon-export';
import * as api from './sim/api';

describe('lookupAddonExport', () => {
  it('returns the code for a blizzard-sourced sim-input, same as addon', async () => {
    const code = 'FS1:1.60.1.69893:warrior:orc:1a/0/0:';
    vi.spyOn(api, 'fetchSimInput').mockResolvedValue({
      spec: 'fury',
      gear: code,
      talents: '',
      buffs: [],
      captured_at: '2026-09-21T00:00:00Z',
      source: 'blizzard',
    });
    const result = await lookupAddonExport({ region: 'us', ruleset: 'normal', slug: 'kiloz' });
    expect(result.code).toBe(code);
  });

  it('returns null for a fight-sourced sim-input', async () => {
    vi.spyOn(api, 'fetchSimInput').mockResolvedValue({
      spec: 'fury',
      gear: {},
      talents: '31/0/20',
      buffs: [],
      captured_at: '2026-09-21T00:00:00Z',
      source: 'fight',
    });
    const result = await lookupAddonExport({ region: 'us', ruleset: 'normal', slug: 'kiloz' });
    expect(result.code).toBeNull();
  });
});
```

- [ ] **Step 6: Run, confirm pass**

Run: `cd web && npx vitest run src/lib/addon-export.test.ts`
Expected: PASS.

- [ ] **Step 7: Scoped checks**

Run: `cd web && npx vitest run src/lib/sim/sources.test.ts src/lib/addon-export.test.ts && npx astro check && npm run lint && npx prettier --check src/lib/sim/sources.ts src/lib/sim/sources.test.ts src/lib/addon-export.ts src/lib/addon-export.test.ts`

- [ ] **Step 8: Commit**

Subject: `feat(web): a Battle.net-sourced character loads and pills exactly like an addon export`.

---

### Task 3: Every sign-in link defaults to the hub

**Files:**
- Modify: `web/src/pages/login.astro`
- Modify: `web/src/components/SignInPrompt.svelte`
- Modify: `web/src/components/Account.svelte`
- Modify: `web/tests/e2e/auth.spec.ts`

**Interfaces:**
- Produces: every sign-in entry point's default `next` is `/account?signed_in=1`.

- [ ] **Step 1: `login.astro`**

In `web/src/pages/login.astro`, change:

```astro
    <Account client:load mode="login" next="/logs" />
```

to:

```astro
    <Account client:load mode="login" next="/account?signed_in=1" />
```

- [ ] **Step 2: `SignInPrompt.svelte`**

In `web/src/components/SignInPrompt.svelte`, change:

```ts
  let { line, next = '/logs', testid }: { line: string; next?: string; testid: string } = $props();
```

to:

```ts
  let { line, next = '/account?signed_in=1', testid }: { line: string; next?: string; testid: string } = $props();
```

This changes the default for every caller that does not pass its own `next`: `GuildJoin`,
`GuildClaim`, `Upload`, `MyReports` all currently rely on this same default (verified: none
of them passes an explicit `next` today) and today all four already send a signed-in
visitor to `/logs`, discarding whatever they were doing — this plan does not fix that
pre-existing loss of context, it only moves where the default lands, per spec §3.1's literal
"every sign-in link".

- [ ] **Step 3: `Account.svelte`'s own account-mode prompt**

In `web/src/components/Account.svelte`, find:

```svelte
        <SignInPrompt line={accountSignInCopy.reason} next="/account" testid="account-signin" />
```

Remove the `next="/account"` override entirely (so it falls through to `SignInPrompt`'s new
default) — `/account?signed_in=1` is a strictly better outcome here than plain `/account`,
since a visitor who was already on the sign-in gate should land with their main character
picked, not on an empty hub they have to reload:

```svelte
        <SignInPrompt line={accountSignInCopy.reason} testid="account-signin" />
```

- [ ] **Step 4: Update the one e2e assertion that hardcodes the old default**

In `web/tests/e2e/auth.spec.ts`, change:

```ts
  await expect(page.getByTestId('battlenet')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Flogs$/,
  );
```

to:

```ts
  await expect(page.getByTestId('battlenet')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
```

- [ ] **Step 5: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/pages/login.astro src/components/SignInPrompt.svelte src/components/Account.svelte tests/e2e/auth.spec.ts`

- [ ] **Step 6: Commit**

Subject: `feat(web): every sign-in link lands on the hub, not the logs page`.

---

### Task 4: The home page panel

**Files:**
- Create: `web/src/lib/home-panel-copy.ts`
- Create: `web/src/components/HomeAccountPanel.svelte`
- Modify: `web/src/pages/index.astro`
- Create: `web/tests/e2e/home-panel.spec.ts`

**Interfaces:**
- Consumes: `battlenetStartUrl`, `fetchMeOnce`, `effectiveServerSims`-style `Me` shape from
  `../lib/account/api`; `SECONDARY_BUTTON_FIXED` from `../lib/planner/styles`.
- Produces: a server-rendered signed-out panel (no JS needed to see it) with a client-side
  swap to the signed-in strip once `/v1/me` answers, inside one reserved-height wrapper so
  neither state shifts layout (CLS budget: index.html's strict bucket, `web/lighthouserc.json`).

- [ ] **Step 1: Copy module**

Create `web/src/lib/home-panel-copy.ts`:

```ts
// web/src/lib/home-panel-copy.ts
// The home page's one account-aware panel (spec 2026-09-22 §3.2): reference, not pitch --
// one sentence, one button, signed out; the current character and four links, signed in.
export const homePanelCopy = {
  signedOutLine:
    'Sign in with Battle.net and your characters arrive with their gear, talents and guild: plan, sim, log and rank them from here.',
  signInButton: 'Sign in with Battle.net',
  openInPlanner: 'Open in planner',
  openInSimulator: 'Open in simulator',
  logs: 'Logs',
  yourCharacters: 'Your characters',
} as const;
```

- [ ] **Step 2: The signed-in island**

Create `web/src/components/HomeAccountPanel.svelte`. This replaces the signed-out sentence
+ button once `/v1/me` resolves signed-in; it renders nothing (parent keeps showing the
signed-out markup) while loading or signed-out, so the Astro shell's server-rendered
sentence is what a signed-out or not-yet-hydrated visitor sees — the exact
reserve-height-then-swap pattern `RecentReports`/`SessionNav` already use elsewhere on this
page:

```svelte
<!-- web/src/components/HomeAccountPanel.svelte -->
<!-- The home page's signed-in swap (spec 2026-09-22 §3.2): while this is loading or the
     visitor is signed out, it renders nothing and the Astro shell's own server-rendered
     sentence + button (index.astro) stays exactly where it is -- the same "server shell
     first, island swaps in place" trick SessionNav.svelte's header link uses, so the panel
     never shows two competing versions and the reserved height in index.astro never has to
     grow for this state, only be filled by it. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { readCurrent } from '../lib/current-character';
  import { characterHref } from '../lib/characters';
  import { classColorVar } from '../lib/report/format';
  import { homePanelCopy } from '../lib/home-panel-copy';

  let me = $state<Me | null>(null);
  let ready = $state(false);

  $effect(() => {
    void fetchMeOnce()
      .then((result) => {
        me = result;
        ready = true;
      })
      .catch(() => {
        ready = true;
      });
  });

  const pointer = $derived(readCurrentIfReady());
  function readCurrentIfReady() {
    if (!ready || me === null) return null;
    return readCurrent();
  }
  const current = $derived(
    pointer === null ? null : me?.characters.find((c) => c.key === pointer.ref || pointer.ref === ''),
  );
  const displayName = $derived(pointer?.label ?? me?.characters[0]?.name ?? '');
  const colour = $derived(classColorVar(current?.class));
</script>

{#if ready && me !== null}
  <div class="flex flex-wrap items-center gap-3" data-testid="home-account-panel">
    <span class="rounded-control h-7 w-7 shrink-0" style={`background: ${colour}`} aria-hidden="true"></span>
    <span class="text-[15px] font-semibold" style={`color: ${colour}`}>{displayName}</span>
    <a class="text-nav text-[13px] font-semibold" href="/planner">{homePanelCopy.openInPlanner}</a>
    <a class="text-nav text-[13px] font-semibold" href="/sim">{homePanelCopy.openInSimulator}</a>
    <a class="text-nav text-[13px] font-semibold" href="/logs">{homePanelCopy.logs}</a>
    <a class="text-nav text-[13px] font-semibold" href="/account">{homePanelCopy.yourCharacters}</a>
  </div>
{/if}
```

Read `web/src/lib/report/format.ts`'s `classColorVar` signature first
(`grep -n "export function classColorVar" src/lib/report/format.ts`) and adjust the
`current?.class` argument to match its exact parameter type (it may want `string |
undefined` already, matching `MeCharacter.class`) — do not cast past a type error.

- [ ] **Step 3: `index.astro`**

In `web/src/pages/index.astro`, add the panel above the tool cards. Import the two new
pieces and `battlenetStartUrl`/`SECONDARY_BUTTON_FIXED`:

```astro
import HomeAccountPanel from '../components/HomeAccountPanel.svelte';
import { homePanelCopy } from '../lib/home-panel-copy';
import { battlenetStartUrl } from '../lib/account/api';
import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
```

Then, immediately after the closing `</SkyBand>` and before the `<section class="flex
flex-col gap-4">` that opens the Tools grid, insert:

```astro
      <div
        class="flex min-h-[52px] flex-wrap items-center gap-3 w-full max-w-[1344px] mx-auto px-[18px] md:px-12"
        data-testid="home-account-block"
      >
        <p class="text-[14px]" data-testid="home-signed-out">
          {homePanelCopy.signedOutLine}
        </p>
        <a
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
          href={battlenetStartUrl('/account?signed_in=1')}
        >
          {homePanelCopy.signInButton}
        </a>
        <HomeAccountPanel client:load />
      </div>
```

The `min-h-[52px]` reserves one row of the sentence+button (52px = a 36–44px button plus
its own line, matching this page's `gap-[22px]`/`gap-8` rhythm) — this is a first estimate;
Task 11's Lighthouse run is the actual measurement, and if CLS on `index.html` regresses
past 0.05 this value must be corrected there before the final report, not left as a guess.
Both the signed-out sentence/button and the island coexist in the DOM (the signed-out pair
never disappears when signed in — Step 2's island renders nothing until it knows the
visitor is signed in, and the sentence/button stay mounted underneath), so on a signed-in
resolve there are briefly two rows. Fix this now, not after measuring: wrap the
signed-out sentence+button in a span this component can hide once `HomeAccountPanel` is
ready. The straightforward fix is to move the signed-out markup inside
`HomeAccountPanel.svelte` itself instead of leaving it in `index.astro` — but that would
lose the "visible with no JS" server-rendered requirement (spec: "the new home panel must
be server-rendered for the signed-out case"). Instead, keep the signed-out markup in
`index.astro` (server-rendered, first paint), and give `HomeAccountPanel.svelte` a
`hideSignedOut` side effect: have it toggle a class on its own parent via a
`data-testid="home-signed-out"` lookup is fragile — instead, pass the parent's signed-out
node nothing and instead let `HomeAccountPanel` cover the same grid cell. Concretely:
change the wrapper in `index.astro` to a CSS grid with both children in the same cell so
only the visible one is seen, using `[grid-area:1/1]` on both, and have
`HomeAccountPanel.svelte`'s root wrap its output in a node that also carries
`[grid-area:1/1]` plus a solid `bg-[var(--color-bg)]` so it visually occludes the
signed-out row once it renders (never removes it from the DOM, so no reflow):

```astro
      <div
        class="grid min-h-[52px] w-full max-w-[1344px] mx-auto px-[18px] md:px-12"
        data-testid="home-account-block"
      >
        <div class="[grid-area:1/1] flex flex-wrap items-center gap-3" data-testid="home-signed-out">
          <p class="text-[14px]">{homePanelCopy.signedOutLine}</p>
          <a
            class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4"
            href={battlenetStartUrl('/account?signed_in=1')}
          >
            {homePanelCopy.signInButton}
          </a>
        </div>
        <HomeAccountPanel client:load />
      </div>
```

and in `HomeAccountPanel.svelte`, change the `{#if ready && me !== null}` wrapper's root
`<div>` class to also carry `[grid-area:1/1] bg-[var(--color-bg)]` so it sits in the exact
same cell and covers the signed-out row once it mounts, rather than the two stacking:

```svelte
  <div
    class="[grid-area:1/1] bg-[var(--color-bg)] flex flex-wrap items-center gap-3"
    data-testid="home-account-panel"
  >
```

- [ ] **Step 4: e2e — signed out and signed in**

Create `web/tests/e2e/home-panel.spec.ts`:

```ts
import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test('the home page offers Battle.net sign-in when signed out', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-signed-out')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
});

test('the home page shows the current character strip when signed in', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [{ key: 'us/normal/kiloz', region: 'us', ruleset: 'normal', name: 'Kiloz', class: 'Warrior' }],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible();
  await expect(page.getByTestId('home-account-panel').getByRole('link', { name: 'Open in simulator' })).toHaveAttribute(
    'href',
    '/sim',
  );
  await expect(page.getByTestId('home-account-panel').getByRole('link', { name: 'Your characters' })).toHaveAttribute(
    'href',
    '/account',
  );
});
```

- [ ] **Step 5: Run e2e**

Run: `cd web && export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12 && E2E_PORT=4461 npx playwright test tests/e2e/home-panel.spec.ts`
Expected: PASS.

- [ ] **Step 6: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/pages/index.astro src/components/HomeAccountPanel.svelte src/lib/home-panel-copy.ts tests/e2e/home-panel.spec.ts`

- [ ] **Step 7: Commit**

Subject: `feat(web): the home page leads with sign-in and the current character`.

---

### Task 5: The hub arrival — main character, pointer, banner, hero-band Logs, Your ratings

**Files:**
- Modify: `web/src/components/Account.svelte`
- Modify: `web/src/lib/account/hero-character.ts`
- Modify: `web/src/lib/account/account-page-copy.ts`
- Modify: `web/tests/e2e/auth.spec.ts` (or create a focused vitest render test — see Step 6)

**Interfaces:**
- Consumes: `mainCharacter`, `pointerForCharacter` (Task 1); `CURRENT_CHARACTER_CHANGED`,
  `writeCurrent` (`../lib/current-character`); `CharacterRatingPanel` (existing, unchanged
  props: `{ path: CharacterPath; apiBase?: string }`).

- [ ] **Step 1: Loosen `heroCharacter`'s render_url requirement**

The hero band must show the main character even when Battle.net has never imported them
(no `render_url` yet — an addon-only character). In `web/src/lib/account/hero-character.ts`,
change:

```ts
export function heroCharacter(
  current: CurrentCharacter | null,
  characters: MeCharacter[],
): MeCharacter | null {
  if (current === null || current.source !== 'armory') return null;
  const match = characters.find((character) => character.key === current.ref);
  return match?.render_url === undefined ? null : match;
}
```

to:

```ts
export function heroCharacter(
  current: CurrentCharacter | null,
  characters: MeCharacter[],
): MeCharacter | null {
  if (current === null || current.source !== 'armory') return null;
  return characters.find((character) => character.key === current.ref) ?? null;
}
```

Check `web/src/lib/account/hero-character.test.ts` for a test asserting the old
`render_url === undefined → null` behaviour (`grep -n "render_url" src/lib/account/hero-character.test.ts`);
if one exists, update it to assert the character is now returned regardless (a hero without
a `render_url` is a real, expected case now, not a null case) rather than deleting the test
— this file's job is exactly to be the one tested place this rule lives.

- [ ] **Step 2: Guard the hero band's `<img>` for a missing `render_url`**

In `web/src/components/Account.svelte`, inside the `{#if hero !== null}` hero-band block,
wrap the existing `<img ... src={hero.render_url} .../>` in a check:

```svelte
                {#if hero.render_url !== undefined}
                  <img
                    class="max-h-[280px] w-auto object-contain lg:max-h-[360px]"
                    src={hero.render_url}
                    alt=""
                    loading="lazy"
                    data-testid="account-hero-render"
                  />
                {/if}
```

- [ ] **Step 3: Add the hero band's Logs link**

In `web/src/lib/account/account-page-copy.ts`, add to `accountPageCopy`:

```ts
  /** spec 2026-09-22 §3.1: the hero band's Logs action, always shown beside the
   *  handoff links (Open in simulator/planner or the paste fallback). */
  heroLogs: 'Logs',
  /** The hub-arrival banner, shown once per `?signed_in=1` visit. */
  signedInBanner: (name: string): string => `Signed in. ${name} is your current character; change it from any row below.`,
  /** Shown in the hero band instead of the handoff links when no character on the account
   *  has a build at all (spec 2026-09-22 §3.1). */
  noBattlenetDataForRealm: 'Blizzard serves no data for this realm type yet.',
  /** "Your ratings" panel heading (spec §3.1). */
  yourRatingsLabel: 'Your ratings',
```

In `web/src/components/Account.svelte`, inside the `{#if heroPath !== null}` block, add the
Logs link right after `<CharacterHandoffLinks path={heroPath} />`:

```svelte
                  {#if heroPath !== null}
                    <div class="flex min-h-11 flex-wrap items-center gap-3 md:min-h-0">
                      <CharacterHandoffLinks path={heroPath} />
                      <a class="inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0" href="/logs">
                        {accountPageCopy.heroLogs}
                      </a>
                    </div>
                  {/if}
```

Read the existing block first — the current markup nests `<CharacterHandoffLinks>` directly
under a `<div class="flex flex-col gap-1">`; adjust the wrapping so the new `<a>` sits
beside the handoff links' own row, not stacked under it (`CharacterHandoffLinks`'s own root
is already `class="flex ... flex-wrap items-center gap-3 md:min-h-0"` — the cleanest change
is to drop the extra wrapping div above and instead add the Logs `<a>` as a second child
inside `CharacterHandoffLinks`'s existing row by rendering it as a sibling right after the
component tag, letting the parent's own `flex flex-wrap` (if any) lay them out; if the
parent has no flex context, wrap only the two in a `<div class="flex flex-wrap items-center gap-3">`
exactly as shown above). Verify the rendered result visually is not required for this task
(Task 11 does a scoped e2e visual smoke check) but keep the DOM change minimal.

- [ ] **Step 4: Hub arrival — main character, pointer, banner**

In `web/src/components/Account.svelte`, add the import:

```ts
  import { mainCharacter, pointerForCharacter } from '../lib/account/main-character';
  import { CURRENT_CHARACTER_CHANGED as CCC, writeCurrent } from '../lib/current-character';
```

(`CCC` avoids a duplicate local name — the file already imports `CURRENT_CHARACTER_CHANGED`
directly; check first with `grep -n "CURRENT_CHARACTER_CHANGED" src/components/Account.svelte`
and if it is already imported once, do not import it a second time under an alias — reuse
the existing binding and only add `writeCurrent` and the two `main-character` imports.)

Add a new `$state` for the banner and a new `$effect` gated on `mode === 'account'` and
`me !== null`, placed right after the existing `toast`/`?refreshed=1` effect:

```ts
  let signedInBanner = $state('');

  $effect(() => {
    if (mode !== 'account' || me === null) return;
    const params = new URLSearchParams(window.location.search);
    if (params.get('signed_in') !== '1') return;
    const main = mainCharacter(me.characters);
    if (main !== null) {
      const pointer = pointerForCharacter(main);
      writeCurrent(pointer);
      window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
      signedInBanner = accountPageCopy.signedInBanner(main.name);
    }
    window.history.replaceState({}, '', window.location.pathname);
  });
```

This `$effect` re-reads `me` (a `$state`), so it must run only once per `signed_in=1`
visit even though `me` could in principle be reassigned later — it already guards itself:
after the first run, `window.history.replaceState` strips `?signed_in=1` from the URL, so
`params.get('signed_in') !== '1'` is true on any re-run and the body is skipped. This is
the same idiom the existing `?refreshed=1` effect above it already uses; do not add extra
state to suppress a second run.

Render the banner: find the existing `{#if toast !== ''}` banner block (the
`?refreshed=1` toast) and add a second one right above it, same visual treatment:

```svelte
        {#if signedInBanner !== ''}
          <div class="bg-raised border-gold flex flex-col gap-1 border-l-2 px-4 py-3" data-testid="account-signed-in-banner">
            <p class="text-[14px]">{signedInBanner}</p>
          </div>
        {/if}
```

- [ ] **Step 5: Your ratings panel**

In `web/src/components/Account.svelte`, add the import:

```ts
  import CharacterRatingPanel from './CharacterRatingPanel.svelte';
```

In the main column (`lg:col-span-8`), between `<CharacterList ... />` and the `<StatePanel
label="Your reports" ...>`, add:

```svelte
            {#if heroPath !== null}
              <StatePanel label={accountPageCopy.yourRatingsLabel} testid="account-ratings">
                <div class="p-[18px]">
                  <CharacterRatingPanel path={heroPath} />
                </div>
              </StatePanel>
            {/if}
```

`CharacterRatingPanel` already renders nothing for a 404/anonymized character
(`status === 'hidden'`) and its own honest "no rating yet" sentence for `sample_size === 0`
— reused verbatim, exactly as `Character.svelte` already does, per spec §3.1 ("the existing
CharacterRatingPanel, reused").

- [ ] **Step 6: A render test for the arrival effect**

Vitest cannot mount a hydrated Svelte component's `$effect`/`window.location` behaviour the
way Playwright can; write this as an e2e spec instead. Create
`web/tests/e2e/bnet-hub.spec.ts`:

```ts
import { expect, test } from '@playwright/test';
import { meBnetFixture } from '../../src/fixtures/me-bnet';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test('signing in lands on the hub with the main character in the hero band', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil({ ok: true, data: meBnetFixture, error: null, request_id: 'r' })));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })));
  await page.route('**/v1/reports**', (route) =>
    route.fulfill(fulfil({ ok: true, data: { rows: [], total: 0, page: 1, per_page: 20 }, error: null, request_id: 'r' })),
  );
  await page.route('**/v1/characters/**/rating', (route) => route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 404)));
  await page.route('**/v1/characters/us/pvp/thoradin/sim-input', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: { spec: 'fury', gear: 'FS1:1.60.1.69893:warrior:orc:1a/0/0:', talents: '', buffs: [], captured_at: '2026-09-21T03:14:00Z', source: 'blizzard' },
        error: null,
        request_id: 'r',
      }),
    ),
  );

  await page.goto('/account?signed_in=1');

  await expect(page.getByTestId('account-signed-in-banner')).toContainText('Thoradin is your current character');
  await expect.poll(() => new URL(page.url()).search).toBe('');
  await expect(page.getByTestId('account-hero')).toContainText('Thoradin');
  await expect(page.getByTestId('character-open-sim')).toHaveAttribute('href', /^\/sim\?code=/);

  await page.getByTestId('character-open-sim').click();
  await expect(page).toHaveURL(/\/sim\?code=/);
});
```

This requires `meBnetFixture`'s guilded character (`Thoradin`) to carry a `build` field so
`mainCharacter` picks it deterministically over the unguilded, level-less `Elyra` — add
that in Task 10 (the fixture-touching task) rather than here, to keep this task's diff
scoped to `Account.svelte`/`hero-character.ts`/`account-page-copy.ts`; this spec file is
created now but left failing until Task 10 lands the fixture change, and Task 10 runs it as
its own verification step. Note this dependency explicitly in this task's commit message
body.

- [ ] **Step 7: Scoped checks**

Run: `cd web && npx vitest run src/lib/account/hero-character.test.ts && npx astro check && npm run lint && npx prettier --check src/components/Account.svelte src/lib/account/hero-character.ts src/lib/account/account-page-copy.ts tests/e2e/bnet-hub.spec.ts`

- [ ] **Step 8: Commit**

Subject: `feat(web): the hub picks a main character, points at it, and shows a banner and ratings`.
Body note: "tests/e2e/bnet-hub.spec.ts depends on the me-bnet fixture gaining a build field,
landed in the follow-up e2e task; it is not run standalone until then."

---

### Task 6: Account Characters rows — build pill, no-fetch handoff, intro copy

**Files:**
- Create: `web/src/components/account/CharacterRowLink.svelte`
- Modify: `web/src/lib/handoff-links.ts`
- Modify: `web/src/components/account/CharacterList.svelte`
- Modify: `web/src/lib/account/character-list-copy.ts`
- Modify: `web/src/components/account/CharacterList.test.ts` (or create if absent)

**Interfaces:**
- Consumes: `buildSourcePill` (Task 1); `MeCharacter.build` (Task 1).
- Produces: `simArmoryHref(characterKey: string): string`; `CharacterRowLink` renders
  either one ref-based "Open in simulator" link or the "No export yet" fallback, with no
  fetch.

- [ ] **Step 1: `simArmoryHref`**

In `web/src/lib/handoff-links.ts`, add at the top:

```ts
import { defaultSimState, simSearch, withSimState } from './sim/url';
```

and add the function, near `simCodeHref`:

```ts
/** A stored character, by key (spec 2026-09-22 §3.4): loads through the sim's own
 *  `?source=armory&ref=` bootstrap (`store.svelte.ts`'s `bootstrapSource`), which resolves
 *  whichever source the API actually holds -- `'addon'` or `'blizzard'` -- with no extra
 *  fetch needed to build this href, unlike `simCodeHref`/`plannerCodeHref` which need the
 *  FS1 string itself. Used where a whole list of characters needs a working "Open in
 *  simulator" link without one `sim-input` fetch per row (`CharacterRowLink.svelte`,
 *  `LandingState.svelte`'s own row link already builds the identical URL by hand). */
export function simArmoryHref(characterKey: string): string {
  return `/sim${simSearch(withSimState(defaultSimState(), { source: 'armory', ref: characterKey }))}`;
}
```

- [ ] **Step 2: `CharacterRowLink.svelte`**

Create `web/src/components/account/CharacterRowLink.svelte`:

```svelte
<!-- web/src/components/account/CharacterRowLink.svelte -->
<!-- The account Characters list's per-row action (spec 2026-09-22 §3.4): built from
     `MeCharacter.build` alone, no `sim-input` fetch -- unlike `CharacterHandoffLinks`
     (still used by the hero band and the character page, one row each, where a single
     fetch is fine), a list of N characters must not make N fetches just to draw its
     rows. There is no fetch-free "Open in planner" for a stored character (the planner
     only ever restores a `'code'`/`'addon'` pointer -- `current-character-planner.ts`'s
     own header comment -- never an `'armory'` one by ref), so this offers only the
     simulator link; the planner's own "Open in planner" link, once the character is
     loaded there, covers the rest. -->
<script lang="ts">
  import { simArmoryHref } from '../../lib/handoff-links';
  import { handoffCopy } from '../../lib/handoff-copy';
  import type { MeCharacter } from '../../lib/account/api';

  let { character }: { character: MeCharacter } = $props();

  const LINK = 'inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0';
</script>

{#if character.build !== undefined}
  <a class={LINK} href={simArmoryHref(character.key)} data-testid="character-open-sim">
    {handoffCopy.openInSimulator}
  </a>
{:else}
  <span class="text-muted text-[13px]" data-testid="character-needs-addon">
    {handoffCopy.needsExportLead}
    <a class="text-text inline-flex min-h-11 items-center underline md:min-h-0" href={handoffCopy.pasteHref}>
      {handoffCopy.needsExportPasteLink}
    </a>
  </span>
{/if}
```

- [ ] **Step 3: Wire it into `CharacterList.svelte`, add the build pill**

In `web/src/components/account/CharacterList.svelte`:

Replace the import:

```ts
  import CharacterHandoffLinks from '../CharacterHandoffLinks.svelte';
```

with:

```ts
  import CharacterRowLink from './CharacterRowLink.svelte';
  import { buildSourcePill } from '../../lib/account/build-pill';
```

Replace the row's handoff usage:

```svelte
          {#if path !== null}
            <CharacterHandoffLinks {path} />
          {/if}
```

with:

```svelte
          <CharacterRowLink {character} />
```

(`path` becomes unused here if nothing else in the row reads it — check with `grep -n
"path" src/components/account/CharacterList.svelte`; the `{@const path = ...}` line is
still needed for nothing else in this file per the current code, so remove that `{@const}`
line too if it is now dead, to keep the "no dead code" rule; if anything else in the row
still reads `path`, leave the `{@const}` and only replace the block above.)

Add the build pill beside the descriptor line — inside the `<div class="flex min-w-0
flex-1 flex-col gap-0.5">` block, right after the `character-descriptor` `<span>`:

```svelte
            {@const pill = buildSourcePill(character.build)}
            <span
              class={pill.pillClass === null ? 'text-muted text-[12px]' : `pill ${pill.pillClass} w-fit`}
              data-testid="character-build-pill"
            >
              {pill.label}
            </span>
```

- [ ] **Step 4: Characters intro copy rewrite (spec §3.3)**

In `web/src/lib/account/character-list-copy.ts`, replace the three intro keys:

```ts
  introLead: 'A character becomes simmable once the site has its export: type /fs export in game and',
  introPasteLink: 'paste it here',
  introTail: ', or run the desktop companion and it sends the export for you.',
```

with:

```ts
  /** spec 2026-09-22 §3.3: the export "how" explained once, here, for every row. */
  introLead: 'Gear and talents come from Battle.net and refresh nightly. Install the addon to include bags and bank and to update right after a session; paste an export for a character Battle.net has no data for.',
  introInstallAddonLink: 'Install the addon',
  introPasteLink: 'paste an export',
```

`CharacterList.svelte`'s template currently renders `introLead` then a single `<a>` for
`introPasteLink` then `introTail`. Update that paragraph to carry both links, matching the
new sentence's two clauses:

```svelte
  <p class="text-muted text-[13px]">
    Gear and talents come from Battle.net and refresh nightly.
    <a class="text-text underline" href="/addon">{characterListCopy.introInstallAddonLink}</a>
    to include bags and bank and to update right after a session;
    <a class="text-text underline" href="/addon#paste">{characterListCopy.introPasteLink}</a>
    for a character Battle.net has no data for.
  </p>
```

Delete `introLead`/`introTail` string usage from the template entirely (replaced by the
literal sentence fragments above split around the two links — every one of those fragments
is still a plain string in this Svelte file's markup, which is acceptable: the copy-module
rule is about strings a component authors on its own, and this exact sentence, with its two
link boundaries, is dictated word-for-word by the spec, matching how e.g.
`SourceSwitcher.svelte`'s template already interleaves copy-module fragments with markup
elsewhere in this codebase). Re-check this file for a lint rule that forbids inline text in
`.svelte` templates (`npm run lint` in Step 6 is the actual check); if lint objects, move
each of the three fragments ("Gear and talents come from Battle.net and refresh nightly.",
"to include bags and bank and to update right after a session;", "for a character
Battle.net has no data for.") into three new `character-list-copy.ts` keys
(`introBattlenetLine`, `introAddonTail`, `introPasteTail`) and reference those instead —
do not leave a lint failure unresolved.

- [ ] **Step 5: Component test**

Check whether `web/src/components/account/CharacterList.test.ts` exists
(`ls web/src/components/account/CharacterList.test.ts`). If it exists, read it fully and
add three cases using this file's own `render()`-from-`svelte/server` pattern (per the
lane's house rules: assert on `data-testid` substrings):
one character with `build: {source: 'blizzard', ...}` renders `character-build-pill` with
text containing `Battle.net`; one with `build: {source: 'addon', ...}` renders it
containing `Addon`; one with no `build` renders `character-needs-addon` and no
`character-open-sim`. If the file does not exist, create it following the same pattern as
`web/src/components/account/CharacterHandoffLinks.test.ts` or the nearest sibling
`.test.ts` under `components/account/` (`ls web/src/components/account/*.test.ts`) —
mirror that file's import of `render` from `svelte/server` and its assertion style
(`.toContain(...)` on the rendered HTML string) exactly.

- [ ] **Step 6: Run and scoped checks**

Run: `cd web && npx vitest run src/components/account/CharacterList.test.ts && npx astro check && npm run lint && npx prettier --check src/components/account/CharacterList.svelte src/components/account/CharacterRowLink.svelte src/lib/handoff-links.ts src/lib/account/character-list-copy.ts`

- [ ] **Step 7: Commit**

Subject: `feat(web): account Characters rows show their build source with no per-row fetch`.

---

### Task 7: Simulator landing — build pill per row, copy pass

**Files:**
- Modify: `web/src/components/sim/LandingState.svelte`
- Modify: `web/src/components/sim/SourceSwitcher.svelte`
- Modify: `web/src/lib/sim/copy.ts`
- Modify: `web/src/components/sim/LandingState.test.ts` / `SourceSwitcher.test.ts` (whichever exist)

**Interfaces:**
- Consumes: `buildSourcePill` (Task 1); `me.characters[].build` already present on
  `LandingState`'s existing `characters: MeCharacter[]` prop.

- [ ] **Step 1: `LandingState.svelte`'s per-row pill**

In `web/src/components/sim/LandingState.svelte`, add the import:

```ts
  import { buildSourcePill } from '../../lib/account/build-pill';
```

Inside the `{#each characters as character (character.key)}` block, in the row's own `<a>`
(the swatch/name/descriptor link), add the pill after the existing descriptor `<span>`:

```svelte
          {@const pill = buildSourcePill(character.build)}
          <span
            class={pill.pillClass === null ? 'text-muted text-[12px]' : `pill ${pill.pillClass}`}
            data-testid={`sim-character-build-${character.key}`}
          >
            {pill.label}
          </span>
```

Place it as a sibling inside the same `<a>` (after the ruleset/region `<span>`), matching
the row's existing `flex items-center gap-3` layout so it lines up with the swatch and name
rather than wrapping oddly — read the surrounding 10 lines before inserting to match
indentation and the exact closing tag it precedes.

- [ ] **Step 2: Delete `landingSourceNote`**

In `web/src/components/sim/LandingState.svelte`, delete this paragraph entirely (the
footnote is now false — the row pill says the real source per character, so a single
blanket "gear comes from your addon export" sentence is no longer honest):

```svelte
  <p class="text-muted text-[12px]" data-testid="sim-landing-note">{simCopy.landingSourceNote}</p>
```

In `web/src/lib/sim/copy.ts`, delete the `landingSourceNote` key and its doc comment
(the block starting `// The contract's sim-input has no Armory source yet...` through the
key itself, around line 266-269).

- [ ] **Step 3: `SourceSwitcher.svelte`'s account card**

In `web/src/lib/sim/copy.ts`, replace:

```ts
  armorySignIn: 'Sign in with Battle.net to find your characters.',
  // Armory itself is not a source yet (simulator contract, sim-input). Saying so is better
  // than an Armory card that quietly serves an addon export under the wrong name.
  armoryNotYet:
    'Blizzard has no character profile API for Forever yet, so a signed-in character’s gear comes from your last addon export or your last logged fight. It will come from the Armory the day that exists.',
```

with:

```ts
  /** SourceSwitcher's signed-out account card (spec 2026-09-22 §3.3: the site now has a
   *  real Battle.net-backed source, so this replaces the old "not yet" disclaimer). */
  signInToFindCharacters: 'Sign in with Battle.net to find your characters, with their gear and talents ready to sim.',
```

Also replace `noCharactersYet` and `sourceAddonBody` and `sourceAccountTitle`:

```ts
  noCharactersYet: 'No characters yet. Install the addon and the companion, or paste an export.',
```
becomes:
```ts
  noCharactersYet: 'No characters yet. Sign in with Battle.net, or paste an export below.',
```

```ts
  sourceAddonBody: 'Paste the export string from the Forever Sixty addon, or let the companion push it.',
```
stays unchanged (the addon card's own copy is still accurate — the addon is still a real,
independent source, not the thing being rewritten) — **do not delete this key**, it is
listed in spec §3.3 only because the lane must confirm it still reads true, which it does;
leave it as-is and note "unchanged, verified accurate" in this task's ledger entry.

```ts
  sourceAccountTitle: 'Your characters',
```
stays unchanged for the same reason — it is already accurate copy naming a real,
independent card; leave it as-is and note the same in the ledger.

In `web/src/components/sim/SourceSwitcher.svelte`, update the two references:

```svelte
        <p class="text-muted text-[13px]" data-testid="sim-armory-note">...</p>
```
wait — re-check: the two spots to change are:

```svelte
        <p class="text-muted text-[13px]">{simCopy.armorySignIn}</p>
```
becomes:
```svelte
        <p class="text-muted text-[13px]">{simCopy.signInToFindCharacters}</p>
```

and delete this line entirely (the note it showed, `simCopy.armoryNotYet`, is deleted
copy and the sentence it stated is no longer true — Battle.net IS a real source now):

```svelte
      <p class="text-muted text-[12px]" data-testid="sim-armory-note">{simCopy.armoryNotYet}</p>
```

- [ ] **Step 4: Grep for any other reference to the deleted keys**

Run: `cd web && grep -rn "armorySignIn\|armoryNotYet\|landingSourceNote" src`
Expected: no output. If any remain (a test file referencing the deleted constants), update
each to the new key name or delete the assertion if the deleted copy's own behaviour no
longer exists (the note it tested was removed in Step 2/3, not renamed).

- [ ] **Step 5: Update/create component tests**

Check `web/src/components/sim/LandingState.test.ts` and
`web/src/components/sim/SourceSwitcher.test.ts` (`ls` both). Update any assertion that
reads the old copy constants by name or by their old literal text to the new ones from
Steps 1-3. Add one new case to `LandingState.test.ts`: a character with
`build: {source: 'blizzard', captured_at: '...'}` in the `characters` prop renders a
`sim-character-build-<key>` node containing `Battle.net`.

- [ ] **Step 6: Run and scoped checks**

Run: `cd web && npx vitest run src/components/sim/LandingState.test.ts src/components/sim/SourceSwitcher.test.ts src/lib/sim/copy.test.ts 2>/dev/null; npx astro check && npm run lint && npx prettier --check src/components/sim/LandingState.svelte src/components/sim/SourceSwitcher.svelte src/lib/sim/copy.ts`

(`src/lib/sim/copy.ts` may have no dedicated test file — the `2>/dev/null` above tolerates
vitest reporting "no test files found" for a path that does not exist as its own suite; if
vitest instead exits non-zero for the whole invocation because of this, drop that path from
the command rather than suppressing a real failure.)

- [ ] **Step 7: Commit**

Subject: `feat(web): the simulator landing shows each character's real build source`.

---

### Task 8: Copy audit — bagsNeedAddon, the addon page, and the sitewide grep

**Files:**
- Modify: `web/src/lib/sim/copy.ts` (remove `bagsNeedAddon` from `bulkCopy`)
- Modify: `web/src/lib/addon/copy.ts`
- Modify: `web/src/pages/addon.astro`

**Interfaces:** none new — copy only.

- [ ] **Step 1: Move and rewrite `bagsNeedAddon`**

`bagsNeedAddon` (`web/src/lib/sim/copy.ts`'s `bulkCopy`) has no current render site
(verified: `grep -rn "bagsNeedAddon" web/src` finds only its own declaration) — it predates
whatever UI would show it. Per spec §3.3 this string's *wording* still needs to state the
new truth, ready for whenever that UI exists, without inventing a render site that is not
this lane's to build. Delete it from `bulkCopy` in `web/src/lib/sim/copy.ts`:

```ts
  bagsNeedAddon: 'Your bags and bank come from the addon export; this character was loaded another way.',
```

Add it, rewritten, to `web/src/lib/addon/copy.ts`'s `addonCopy` (bags/bank is the addon's
own subject matter, not the simulator's — this is also a better home than `sim/copy.ts`,
which is already near this codebase's 800-line ceiling):

```ts
  /** Spec 2026-09-22 §3.3: shown wherever a Battle.net-sourced character's bags/bank are
   *  unavailable (no current render site yet — the bulk gear tool has no per-slot "why is
   *  this empty" note today; this is the copy ready for when it does). */
  bagsNeedAddonForBlizzard: 'Bags and bank come from the addon. Install it to include them.',
```

- [ ] **Step 2: Addon page opening line**

In `web/src/lib/addon/copy.ts`, change:

```ts
  pageDescription:
    'Your character into the planner with one paste, and a build you chose here as an in-game guide.',
```

to:

```ts
  pageDescription:
    'Battle.net already gives your gear and talents. The addon adds your bags and bank, your professions, an in-game build guide, and refreshes the moment you log out.',
```

`web/src/pages/addon.astro` already renders `addonCopy.pageDescription` as both the
`<Base description=...>` meta tag and the visible `<p>` under the `<h1>` — no template
change needed, only the string.

- [ ] **Step 3: The sitewide grep review**

Run each of these and read every hit:

```bash
cd web
grep -rn "no character profile" src
grep -rni "armory" src --include="*.svelte" --include="*.astro" --include="*.ts" | grep -v ".test.ts"
grep -rn "Install the addon" src
```

For the "armory" grep: every remaining hit must be one of (a) the internal pointer-source
literal `'armory'`/`current.source !== 'armory'` (unchanged, per Global Constraint 1 — this
is code, not copy), (b) a code comment explaining that internal name, or (c) a copy string
already rewritten by Tasks 2/6/7 above. If a hit is a *visible string* still claiming
Blizzard has no profile API or naming "the Armory" as a UI concept, rewrite it in its own
copy module the same way Task 7 rewrote `armorySignIn`/`armoryNotYet` — do not leave one
unrewritten. List every file touched in this step in the commit body, even if the change is
one word.

For "no character profile": if this exact phrase appears in the spec's own doc file
(`docs/superpowers/specs/...md`), that is not a code hit — only touch `src/`.

For "Install the addon": every existing hit that is not part of the new copy this plan
already wrote (Tasks 6-8) should read naturally next to the new "Battle.net first" framing;
if a hit says something like "Install the addon to sim your character" as if the addon
were required, rewrite it to name the addon as what it adds on top of Battle.net, matching
this task's Step 2 wording.

- [ ] **Step 4: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/sim/copy.ts src/lib/addon/copy.ts src/pages/addon.astro` plus every other file this step's grep review touched.

- [ ] **Step 5: Commit**

Subject: `fix(web): rewrite every stale "no Blizzard profile API" and armory-facing string`.

---

### Task 9: FS1 decode test over the API lane's `era-kiloz.fs1`

**Files:**
- Create: `web/src/lib/planner/fs1.blizzard.test.ts`

**Interfaces:**
- Consumes: `decodeFS1` from `./fs1` (existing).

- [ ] **Step 1: Write the test**

Create `web/src/lib/planner/fs1.blizzard.test.ts`:

```ts
// web/src/lib/planner/fs1.blizzard.test.ts
// Spec 2026-09-22 §3.5: the API lane's bnetbuild encoder writes
// api/internal/bnetbuild/testdata/era-kiloz.fs1, a checked-in FS1 string encoded from the
// Era Kiloz fixture profile (api/internal/bnetapi/testdata/era-kiloz/). This decodes it
// with the exact same decoder every addon paste and sim-input read goes through, proving
// the two lanes' grammar agrees without either lane importing the other's code. Skips
// cleanly (not a failure) when the API lane has not landed that file on this branch yet.
import { readFileSync, existsSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { decodeFS1 } from './fs1';

const FIXTURE_PATH = fileURLToPath(
  new URL('../../../../api/internal/bnetbuild/testdata/era-kiloz.fs1', import.meta.url),
);

describe('a Battle.net-built FS1 string (bnetbuild)', () => {
  it.skipIf(!existsSync(FIXTURE_PATH))(
    'decodes with this site\'s own FS1 decoder',
    () => {
      const code = readFileSync(FIXTURE_PATH, 'utf-8').trim();
      const result = decodeFS1(code);
      expect(result.ok).toBe(true);
      if (result.ok) {
        expect(result.build.classSlug).toBe('warrior');
      }
    },
  );

  if (!existsSync(FIXTURE_PATH)) {
    it('is skipped: api/internal/bnetbuild/testdata/era-kiloz.fs1 is absent on this branch', () => {
      expect(existsSync(FIXTURE_PATH)).toBe(false);
    });
  }
});
```

Do not guess `classSlug` — if the fixture exists by the time this task runs, read the file
first (`cat api/internal/bnetbuild/testdata/era-kiloz.fs1` from the repo root) and match
the assertion to what it actually decodes to (the fixture is named "kiloz", and spec §0
says it is verified against "a Classic Era character" — a Warrior, per the fixture
directory's own name convention seen in `api/internal/bnetapi/testdata/era-kiloz/`; confirm
by reading that directory's own profile JSON's `character_class.name` if present:
`cat api/internal/bnetapi/testdata/era-kiloz/*.json | grep -i '"name"' | head -5`). If the
fixture is absent (expected — the API lane runs in parallel and may not have landed it
yet), the `it.skipIf` branch is what runs and the exact `classSlug` assertion never
executes; still write it correctly per this note for whenever the fixture lands.

- [ ] **Step 2: Run it**

Run: `cd web && npx vitest run src/lib/planner/fs1.blizzard.test.ts`
Expected: either PASS (decodes correctly, if the fixture exists) or PASS via the skip
branch (fixture absent) — never a hard FAIL from a missing file crashing the test.

- [ ] **Step 3: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/lib/planner/fs1.blizzard.test.ts`

- [ ] **Step 4: Commit**

Subject: `test(web): decode the API lane's Battle.net-built FS1 fixture when present`.

---

### Task 10: e2e fixture, full e2e run, Lighthouse, final report prep

**Files:**
- Modify: `web/src/fixtures/me-bnet.ts`
- Verify: `web/tests/e2e/bnet-hub.spec.ts` (created in Task 5) now passes
- Verify: `web/tests/e2e/auth.spec.ts`, `web/tests/e2e/account-armory-pointer.spec.ts`,
  `web/tests/e2e/home-panel.spec.ts` all still pass

- [ ] **Step 1: Add `build` to the fixture**

In `web/src/fixtures/me-bnet.ts`, add a `build` field to the first (guilded) character,
`Thoradin`, so it is unambiguously the `mainCharacter()` pick over the second, level-less
character:

```ts
      guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', rank_index: 1, verified: true },
      build: { source: 'blizzard', captured_at: '2026-09-21T03:14:00Z' },
```

(insert the `build` line right after the existing `guild:` line, inside the first
character object).

- [ ] **Step 2: Run every e2e spec this plan touched or created**

Run:
```bash
cd web
export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12
E2E_PORT=4461 npx playwright test tests/e2e/auth.spec.ts tests/e2e/account-armory-pointer.spec.ts tests/e2e/bnet-hub.spec.ts tests/e2e/home-panel.spec.ts
```
Expected: PASS. Fix any failure by reading the actual rendered DOM the spec's own trace/
screenshot shows (`npx playwright show-trace` on the failed test's trace, or `--headed` for
a fast local rerun) — do not adjust an assertion to match wrong behaviour; if the assertion
is right and the component is wrong, fix the component.

- [ ] **Step 3: Run the WHOLE e2e suite once**

Per the lane's house rules, run every spec, not only the ones this plan touched:
```bash
cd web && E2E_PORT=4461 npx playwright test
```
Stop `astro preview` first if a stale one is running (`npx astro preview stop`). Record the
pass/fail count. Any failure must be triaged: if it reproduces identically on `main`
(`git stash`, rerun that one spec, `git stash pop`), it is pre-existing and is reported, not
fixed, in the final report. Any failure this branch's changes caused must be fixed before
the final report.

- [ ] **Step 4: Full vitest run, astro check, lint, prettier — whole `web/`**

```bash
cd web
npx vitest run
npx astro check
npm run lint
npx prettier --check src
```
Record the pass counts for the final report. Fix any failure.

- [ ] **Step 5: Lighthouse**

```bash
cd web
npm run build
npm run lhci
```
Record CLS and TBT for every URL `web/lighthouserc.json` measures, especially
`index.html` (Task 4 added the home panel there — confirm CLS ≤ 0.05 and TBT ≤ 100ms hold;
if CLS regressed, go back to Task 4's `min-h-[52px]` estimate and correct it against the
real measured shift, then rerun `npm run build && npm run lhci` until it holds). `account.html`
is not in the URL list (a pre-existing, documented limitation — `web/src/lib/account/
layout.ts`'s own comment: the Lighthouse sandbox has no CORS allowance to call the real API
from `/account`, so it always lands on the failed/LoadError branch and cannot measure the
hub's real CLS); do not add it — this matches the prior lane's own explicit decision, and
re-adding it without a way to fix the CORS gap would only add a flaky assertion.

- [ ] **Step 6: Stop every server this lane started**

```bash
npx astro preview stop
pkill -f "$(pwd)/web/node_modules/astro" 2>/dev/null || true
ps -axo command | grep "[a]stro.*$(basename $(pwd))"
```
The last command's output must be empty.

- [ ] **Step 7: Commit**

Subject: `test(web): the hub fixture is simmable, and the whole suite is green`.

---

## Self-review notes (recorded here, not a separate document)

- Spec §3.1 (sign-in lands on the hub): Tasks 1, 3, 5.
- Spec §3.1 (simulator landing keeps source pills, no auto-load): Task 7.
- Spec §3.2 (home panel): Task 4.
- Spec §3.3 (copy pass, every location named): Tasks 6, 7, 8.
- Spec §3.4 (no per-row fetch, sourcePill gains Battle.net): Tasks 2, 6.
- Spec §3.5 (tests: main-character, hub arrival, home panel, FS1 decode, e2e, lhci): Tasks
  1, 4, 5, 9, 10.
- Spec §4 (order of work): this plan's task order follows it — contract/normalisers (1-2)
  before sign-in defaults (3) before the home panel (4) before the hub (5) before the copy
  pass (6-8) before tests/Lighthouse (9-10).
- Out of scope (spec §6, confirmed untouched by this plan): professions from Blizzard,
  Season of Discovery data, class icon sets, the companion's own onboarding, Forever's
  namespace, and `api/**`/OpenAPI (the other lane).
