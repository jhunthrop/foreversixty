# Op hand-offs implementation plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Wire every hand-off the spec names outside the sim/planner tools themselves: the
homepage tools grid, `/addon`'s paste box, `/account` and `/character/<key>`'s per-character
links, the report's per-combatant Sim link, and the reference pages' links into the tools.

**Architecture:** One small, dependency-free URL-builder module
(`web/src/lib/handoff-links.ts`) is the single place every `/planner?code=`, `/sim?code=`,
`/sim?source=fight&ref=`, `/sim/drops?instance=` string gets built, so the coordinator can
retarget it at Lane B's `current-character.ts` API in one place at merge time. Everything
else is thin Astro/Svelte wiring around existing, read-only-imported planner/sim library
functions (`decodeFS1`, `plannerHref`, `fetchSimInput`) plus small new pure modules for the
two conditionals the reference pages need (zone match, loot match).

**Tech Stack:** Astro 5, Svelte 5 (runes), TypeScript, Vitest, Playwright. Node 22.12 via
nvm. `FOREVER_DATA=fixture` for all test runs.

**Spec:** `docs/superpowers/specs/2026-09-21-one-product-design.md` (section 0 binds every
lane; section 2 is this lane's scope; sections 1, 3, 4 name what the other three lanes own).

## Global Constraints

(Copied from spec section 0, plus this lane's own file-ownership rule.)

- Design system: `design/DESIGN-SYSTEM.md`. No emoji, no marketing buttons, secondary
  buttons only, restrained motion, 44px hit targets on phone, dark only.
- Honest copy: never promise what is not built; say what a page is for and who it is for.
  No exclamation marks.
- No layout shift: anything that appears after hydration reserves its space or sits below
  existing content. Lighthouse budgets in `web/lighthouserc.json` must hold.
- Signed-out first: every feature works without an account; an account only adds.
- localStorage is a convenience, never the source of truth: wrap every read and write in
  try/catch, render correctly without it, never read it during SSR/prerender. (This lane
  touches no localStorage directly — Lane B owns `current-character.ts` — but any code this
  lane writes that runs client-side must still degrade to the honest "nothing loaded"
  copy, never throw.)
- Tests: unit tests for every pure function, e2e for each new hand-off (desktop and mobile
  Playwright projects). `FOREVER_DATA=fixture npm run sync` before web tests. Node 22.12.
- **File ownership (this lane, "op-handoffs"):** never edit
  `web/src/components/sim/`, `web/src/components/planner/`, `web/src/lib/sim/`,
  `web/src/lib/planner/`, `Header.astro`, `Footer.astro`, `Base.astro`, `logs.astro`,
  rankings pages or components, or `api/`. Reading and read-only importing from
  `lib/planner/` and `lib/sim/` is fine (the codebase's own `Character.svelte` already
  imports `lib/sim/execution.ts` read-only); never write to those directories. Lane B is
  building `web/src/lib/current-character.ts` and
  `web/src/components/CurrentCharacterChip.svelte` in parallel — do not import either; if
  Lane B has not landed, every link this lane builds goes through `handoff-links.ts`
  directly.

---

## File structure

New files this plan creates:
- `web/src/lib/handoff-links.ts` — the four URL-builder functions, plus tests.
- `web/src/lib/addon-export.ts` — `lookupAddonExport`, the one API call this lane makes to
  learn whether a character has a usable addon export, plus tests.
- `web/src/components/CharacterHandoffLinks.svelte` — the shared "Open in simulator / Open
  in planner / Log in with the addon" row, used by both `Account.svelte` and
  `Character.svelte`.
- `web/src/components/AddonPasteBox.svelte` — the `/addon` page's own paste box.
- `web/src/lib/report/sim-link.ts` — the per-combatant Sim link builder, plus tests.
- `web/src/lib/reference/dungeon-links.ts` — the two pure conditionals `/dungeons/<slug>`
  needs (zone match, loot match), plus tests.
- `web/tests/e2e/handoffs-home.spec.ts`, `handoffs-addon.spec.ts`,
  `handoffs-account-character.spec.ts`, `handoffs-report.spec.ts`,
  `handoffs-reference.spec.ts` — new e2e coverage.

Modified files:
- `web/src/data/tools.json`, `web/src/pages/index.astro` — homepage grid.
- `web/src/pages/addon.astro`, `web/src/lib/addon/copy.ts` — paste box + honest beta copy.
- `web/src/components/Account.svelte` — Characters list gains hand-off links.
- `web/src/components/Character.svelte` — header gains hand-off links.
- `web/src/components/report/SummaryTab.svelte`, `DeathsTab.svelte`, `ReportView.svelte` —
  per-combatant Sim link.
- `web/src/pages/guides/[slug].astro` — planner link.
- `web/src/pages/dungeons/[slug].astro` — zone link, drops link.
- `web/src/pages/zones/index.astro` — coverage copy.

---

### Task 1: `handoff-links.ts`, the one URL-builder module

**Files:**
- Create: `web/src/lib/handoff-links.ts`
- Test: `web/src/lib/handoff-links.test.ts`

**Interfaces:**
- Produces: `plannerCodeHref(code: string): string`, `simCodeHref(code: string): string`,
  `simFightHref(reportId: string, fightIndex: number, guid?: string): string`,
  `simDropsHref(instanceSlug: string): string` — every later task in this plan imports one
  or more of these instead of building a query string by hand.

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/handoff-links.test.ts
import { describe, expect, it } from 'vitest';
import { plannerCodeHref, simCodeHref, simDropsHref, simFightHref } from './handoff-links';

describe('handoff-links', () => {
  it('builds a planner link carrying the FS1 code, URL-encoded', () => {
    expect(plannerCodeHref('FS1:1.60:warrior:human:0/0/0:')).toBe(
      '/planner?code=FS1%3A1.60%3Awarrior%3Ahuman%3A0%2F0%2F0%3A',
    );
  });

  it('builds a simulator link carrying the FS1 code, URL-encoded', () => {
    expect(simCodeHref('FS1:1.60:warrior:human:0/0/0:')).toBe(
      '/sim?code=FS1%3A1.60%3Awarrior%3Ahuman%3A0%2F0%2F0%3A',
    );
  });

  it('builds a fight-sourced sim link for the whole fight when no guid is given', () => {
    expect(simFightHref('abc2defg2hij', 3)).toBe('/sim?source=fight&ref=abc2defg2hij%3A3');
  });

  it('builds a fight-sourced sim link scoped to one combatant when a guid is given', () => {
    expect(simFightHref('abc2defg2hij', 3, 'Player-4184-000000A1')).toBe(
      '/sim?source=fight&ref=abc2defg2hij%3A3%3APlayer-4184-000000A1',
    );
  });

  it('builds a drops link preselecting an instance', () => {
    expect(simDropsHref('hall-of-thanes')).toBe('/sim/drops?instance=hall-of-thanes');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/handoff-links.test.ts`
Expected: FAIL — `Cannot find module './handoff-links'`.

- [ ] **Step 3: Write minimal implementation**

```typescript
// web/src/lib/handoff-links.ts
// Every hand-off link this lane builds, in one place, in exactly the four URL forms the
// spec names (docs/superpowers/specs/2026-09-21-one-product-design.md section 2): a pasted
// or carried FS1 code into the planner or the simulator, a fight-sourced character (whole
// fight or one combatant) into the simulator, and an instance preselected on the drops
// tool. No other module in this lane's scope should compose one of these paths by hand —
// that is what lets the coordinator retarget every link at Lane B's
// `current-character.ts` (`plannerHrefFor`/`simHrefFor`) in this one file at merge time.

export function plannerCodeHref(code: string): string {
  return `/planner?code=${encodeURIComponent(code)}`;
}

export function simCodeHref(code: string): string {
  return `/sim?code=${encodeURIComponent(code)}`;
}

/**
 * `<report>:<fight>` when `guid` is omitted (the existing whole-fight "Sim this fight"
 * link's own ref shape), extended to `<report>:<fight>:<guid>` when it is given — the
 * combatant GUID exactly as the report's roster carries it. Lane B's `fromLoggedFight`
 * reads the optional third part and falls back to today's first-dps rule when absent.
 */
export function simFightHref(reportId: string, fightIndex: number, guid?: string): string {
  const ref = guid === undefined ? `${reportId}:${fightIndex}` : `${reportId}:${fightIndex}:${guid}`;
  return `/sim?source=fight&ref=${encodeURIComponent(ref)}`;
}

export function simDropsHref(instanceSlug: string): string {
  return `/sim/drops?instance=${encodeURIComponent(instanceSlug)}`;
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/handoff-links.test.ts`
Expected: PASS, 5 tests.

- [ ] **Step 5: Scoped checks and commit**

```bash
cd web && npx vitest run src/lib/handoff-links.test.ts && npx astro check && npm run lint && npx prettier --check src/lib/handoff-links.ts src/lib/handoff-links.test.ts
```

```bash
printf 'feat(handoffs): add the one URL-builder module for every hand-off link\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-1.txt
git add web/src/lib/handoff-links.ts web/src/lib/handoff-links.test.ts
git commit -F .superpowers/commit-msg-1.txt
```

---

### Task 2: Homepage tools grid gains Simulator and The addon

**Files:**
- Modify: `web/src/data/tools.json`
- Modify: `web/src/pages/index.astro`
- Test: `web/tests/e2e/handoffs-home.spec.ts` (written in Task 10, not here — this task's
  own proof is `astro check` plus a manual read of the rendered order; Task 10 adds the
  Playwright assertion once every hand-off e2e spec is written together)

**Interfaces:**
- Consumes: nothing from earlier tasks.
- Produces: two new tool cards other tasks do not depend on.

- [ ] **Step 1: Reorder and extend `tools.json`**

Replace the file's contents so the order is exactly Planner, Simulator, Logs, Rankings, The
addon, then the three reference cards (their existing relative order kept):

```json
[
  {
    "title": "Build planner",
    "description": "Revamped talents and every race and class combo, including Skyborne and Undead Paladin. Share a build by link.",
    "href": "/planner",
    "status": {
      "label": "Live",
      "kind": "site"
    },
    "icon": "talents"
  },
  {
    "title": "Simulator",
    "description": "Estimate your DPS, find upgrades with the Droptimizer and Top Gear, and check stat weights.",
    "href": "/sim",
    "status": {
      "label": "Live",
      "kind": "site"
    },
    "icon": "simulator"
  },
  {
    "title": "Combat logs",
    "description": "Upload a log or log live with the companion. Every fight, every view, and your parse against everyone else's.",
    "href": "/logs",
    "status": {
      "label": "Live",
      "kind": "site"
    },
    "icon": "logs"
  },
  {
    "title": "Rankings",
    "description": "Every encounter's leaderboard by class and spec, updated the moment a fight ends.",
    "href": "/rankings",
    "status": {
      "label": "Live",
      "kind": "site"
    },
    "icon": "rankings"
  },
  {
    "title": "The addon",
    "description": "Paste an export to load your character into the planner or the simulator, then copy a build back into the game.",
    "href": "/addon",
    "status": {
      "label": "Live",
      "kind": "site"
    },
    "icon": "addon"
  },
  {
    "title": "Dungeons",
    "description": "All nine new dungeons with what is known today. Bosses and loot fill in from beta.",
    "href": "/dungeons",
    "status": {
      "label": "Live",
      "kind": "community"
    },
    "icon": "dungeon"
  },
  {
    "title": "Zone atlas",
    "description": "Mount Hyjal, Zephras Isle, Riverglades, Shen'dralar and the reworked old zones.",
    "href": "/zones",
    "status": {
      "label": "Live",
      "kind": "community"
    },
    "icon": "map"
  },
  {
    "title": "Class guides",
    "description": "What changed for each class in Forever, with the demo findings and every source.",
    "href": "/guides",
    "status": {
      "label": "Live",
      "kind": "community"
    },
    "icon": "guides"
  }
]
```

- [ ] **Step 2: Add the two new icons and a stable test hook in `index.astro`**

In `web/src/pages/index.astro`, find the `icons` record (currently `talents`, `logs`,
`rankings`, `guides`, `dungeon`, `map`) and add two entries:

```typescript
const icons: Record<string, string> = {
  talents:
    '<circle cx="12" cy="5" r="2"/><circle cx="6" cy="12" r="2"/><circle cx="18" cy="12" r="2"/><circle cx="12" cy="19" r="2"/><path d="M12 7v10M8 11l3-4M16 11l-3-4M8 13l3 4M16 13l-3 4"/>',
  simulator: '<path d="M4 20V10"/><path d="M10 20V4"/><path d="M16 20v-6"/><path d="M2 20h20"/>',
  logs: '<path d="M4 19V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v14"/><path d="M4 19h16"/><path d="M8 15v-4M12 15V8M16 15v-6"/>',
  rankings:
    '<path d="M8 21h8M12 17v4"/><path d="M6 4h12v4a6 6 0 0 1-12 0z"/><path d="M6 6H3v2a3 3 0 0 0 3 3M18 6h3v2a3 3 0 0 1-3 3"/>',
  addon: '<path d="M9 2v4"/><path d="M15 2v4"/><path d="M6 8h12l-1 5a5 5 0 0 1-10 0z"/><path d="M12 17v5"/>',
  guides:
    '<path d="M4 5a2 2 0 0 1 2-2h13v16H6a2 2 0 0 0-2 2z"/><path d="M4 19a2 2 0 0 1 2-2h13"/><path d="M9 7h6M9 11h4"/>',
  dungeon: '<path d="M4 20V8l8-5 8 5v12"/><path d="M9 20v-6h6v6"/>',
  map: '<path d="M3 6l6-2 6 2 6-2v14l-6 2-6-2-6 2z"/><path d="M9 4v14M15 6v14"/>',
};
```

Then find `<div class="grid grid-cols-2 lg:grid-cols-3 gap-4">` (the tools grid wrapper) and
add a stable test hook:

```astro
<div class="grid grid-cols-2 lg:grid-cols-3 gap-4" data-testid="tools-grid">
```

- [ ] **Step 3: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/data/tools.json src/pages/index.astro`
Expected: no errors. (No vitest for this task — `tools.json` has no pure-function logic of
its own; Task 10's e2e spec is the real proof.)

- [ ] **Step 4: Commit**

```bash
printf 'feat(homepage): add Simulator and The addon to the tools grid\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-2.txt
git add web/src/data/tools.json web/src/pages/index.astro
git commit -F .superpowers/commit-msg-2.txt
```

---

### Task 3: `/addon` gains a paste box and honest in-game-UI copy

**Files:**
- Create: `web/src/components/AddonPasteBox.svelte`
- Modify: `web/src/lib/addon/copy.ts`
- Modify: `web/src/pages/addon.astro`

**Interfaces:**
- Consumes: `decodeFS1` from `web/src/lib/planner/fs1.ts` (read-only import — this task
  never edits that file), `SECONDARY_BUTTON_FIXED` from `web/src/lib/planner/styles.ts`
  (read-only import), `plannerCodeHref`/`simCodeHref` from Task 1's `handoff-links.ts`.
- Produces: `AddonPasteBox.svelte`, a self-contained island with `data-testid`s
  `addon-paste-box`, `addon-paste-code`, `addon-paste-submit`, `addon-paste-error`,
  `addon-paste-planner`, `addon-paste-sim` — Task 10's e2e spec drives these.

- [ ] **Step 1: Add the new copy to `addon/copy.ts`**

In `web/src/lib/addon/copy.ts`, after the `// --- planner import box ---` block (before
`// --- /addon page ---`), add:

```typescript
  // --- /addon page's own paste box ---
  pasteTitle: 'Try an export',
  pastePlaceholder: 'FS1:…',
  pasteAction: 'Load',
  pasteOpenPlanner: 'Open in planner',
  pasteOpenSim: 'Open in simulator',
```

Inside the existing `// --- /addon page ---` block, after `flowInBody`, add the honest
beta-testing note (no screenshot, because none exists):

```typescript
  inGameTitle: 'The in-game window',
  inGameBody:
    'The addon window, tracker, talent glow, gear tab and minimap button are in beta testing in game. There is no screenshot here because we have not captured one yet.',
```

- [ ] **Step 2: Write the component**

```svelte
<!-- web/src/components/AddonPasteBox.svelte -->
<!-- /addon's own paste box: prove a pasted export decodes with the same decoder the
     planner uses (planner/fs1.ts, read-only import — this component never writes to that
     module) and hand back the two places it goes next. Unlike the planner's own
     ImportBox.svelte, this page never builds a talent draft: it only proves the string is
     readable and links on with the raw code, so it needs no TalentIndex and no active
     build to reconcile against. -->
<script lang="ts">
  import { addonCopy } from '../lib/addon/copy';
  import { decodeFS1 } from '../lib/planner/fs1';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { plannerCodeHref, simCodeHref } from '../lib/handoff-links';

  let code = $state('');
  let error = $state<string | null>(null);
  let loaded = $state<string | null>(null);

  function submit(): void {
    const trimmed = code.trim();
    const result = decodeFS1(trimmed);
    if (!result.ok) {
      error = result.message;
      loaded = null;
      return;
    }
    error = null;
    loaded = trimmed;
  }
</script>

<section
  class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4"
  data-testid="addon-paste-box"
>
  <h2 class="section-title text-[15px]">{addonCopy.pasteTitle}</h2>
  <textarea
    class="border-line rounded-control bg-card-top min-h-11 w-full border px-2 py-1 font-mono text-[13px]"
    rows="2"
    placeholder={addonCopy.pastePlaceholder}
    aria-label={addonCopy.pasteTitle}
    bind:value={code}
    data-testid="addon-paste-code"
  ></textarea>
  <button
    type="button"
    class={SECONDARY_BUTTON_FIXED}
    disabled={code.trim() === ''}
    onclick={submit}
    data-testid="addon-paste-submit"
  >
    {addonCopy.pasteAction}
  </button>
  {#if error !== null}
    <p class="text-strong text-[13px]" data-testid="addon-paste-error" role="alert">{error}</p>
  {/if}
  {#if loaded !== null}
    <div class="flex flex-wrap gap-3">
      <a
        class="text-gold text-[13px] font-semibold"
        href={plannerCodeHref(loaded)}
        data-testid="addon-paste-planner"
      >
        {addonCopy.pasteOpenPlanner}
      </a>
      <a
        class="text-gold text-[13px] font-semibold"
        href={simCodeHref(loaded)}
        data-testid="addon-paste-sim"
      >
        {addonCopy.pasteOpenSim}
      </a>
    </div>
  {/if}
</section>
```

- [ ] **Step 3: Wire it into `addon.astro`**

In `web/src/pages/addon.astro`, add the import:

```typescript
import AddonPasteBox from '../components/AddonPasteBox.svelte';
```

Add the third flow entry (after the existing `flowInTitle`/`flowInBody` one) in the `flows`
array:

```typescript
const flows = [
  { title: addonCopy.flowOutTitle, body: addonCopy.flowOutBody },
  { title: addonCopy.flowInTitle, body: addonCopy.flowInBody },
  { title: addonCopy.inGameTitle, body: addonCopy.inGameBody },
];
```

Place the paste box between the install-links `<Panel>` and the `flows` grid:

```astro
    <Panel>
      <ul class="flex flex-col gap-2 p-[18px] md:flex-row md:flex-wrap md:gap-6">
        {links.map((link) => (
          <li>
            <a
              class="inline-flex min-h-11 items-center text-[14px] font-semibold underline underline-offset-2"
              href={link.href}
              rel="noopener"
            >
              {link.label}
            </a>
          </li>
        ))}
      </ul>
    </Panel>

    <AddonPasteBox client:load />

    <div class="grid gap-4 md:grid-cols-2">
      {flows.map((flow) => (
        <Panel title={flow.title}>
          <p class="text-muted p-[18px] text-[14px] leading-relaxed">{flow.body}</p>
        </Panel>
      ))}
    </div>
```

- [ ] **Step 4: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/AddonPasteBox.svelte src/lib/addon/copy.ts src/pages/addon.astro`
Expected: no errors.

- [ ] **Step 5: Commit**

```bash
printf 'feat(addon): add a paste box and honest in-game beta copy to /addon\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-3.txt
git add web/src/components/AddonPasteBox.svelte web/src/lib/addon/copy.ts web/src/pages/addon.astro
git commit -F .superpowers/commit-msg-3.txt
```

---

### Task 4: Addon-export lookup and the shared `CharacterHandoffLinks` component

**Ruling (record in the ledger verbatim):** `/account` and `/character/<key>` cannot know
whether a character has a usable addon export without an API call of their own —
`MeCharacter` (`lib/account/api.ts`) and `CharacterPage.character`
(`lib/rankings/api.ts`) carry no source flag. The cheapest honest option is the one the
sim's own signed-in landing state already uses when a player picks a stored character:
`GET /v1/characters/{region}/{ruleset}/{name}/sim-input` (`fetchSimInput`,
`web/src/lib/sim/api.ts`, read-only import — never edited by this lane), read per
character, in parallel, best-effort. A character has a usable export only when the
response's `source` is `'addon'` and `gear` is an FS1 string — exactly
`fromStoredCharacter`'s own check (`web/src/lib/sim/sources.ts:112`) — because that is the
only shape a URL can carry into the sim or planner today (`bootstrapSource`,
`web/src/lib/sim/store.svelte.ts`, only resolves `addon`/`build`/`fight` refs; there is no
working `?source=armory&ref=` or `?source=stored&ref=` link to build even if there were).
Cost: one GET per character shown, bounded by how many characters an account or a
character page ever lists; a failed or non-addon lookup reads as "no export" rather than an
error banner.

**Files:**
- Create: `web/src/lib/addon-export.ts`
- Test: `web/src/lib/addon-export.test.ts`
- Create: `web/src/components/CharacterHandoffLinks.svelte`

**Interfaces:**
- Consumes: `fetchSimInput` from `web/src/lib/sim/api.ts` (read-only import), `FS1_PREFIX`
  from `web/src/lib/planner/fs1.ts` (read-only import), `CharacterPath` from
  `web/src/lib/characters.ts`, `plannerCodeHref`/`simCodeHref` from Task 1.
- Produces: `lookupAddonExport(path: CharacterPath, apiBase?: string): Promise<{ path: CharacterPath; code: string | null }>`
  and the `CharacterHandoffLinks.svelte` component (props `{ path: CharacterPath; apiBase?: string }`,
  `data-testid`s `character-handoff`, `character-open-sim`, `character-open-planner`,
  `character-needs-addon`) — Task 5 mounts it in Account.svelte and Character.svelte.

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/addon-export.test.ts
import { describe, expect, it, vi } from 'vitest';

const fetchSimInput = vi.fn();
vi.mock('./sim/api', () => ({ fetchSimInput: (...args: unknown[]) => fetchSimInput(...args) }));

const { lookupAddonExport } = await import('./addon-export');

const PATH = { region: 'us' as const, ruleset: 'normal' as const, slug: 'thrallgar' };

describe('lookupAddonExport', () => {
  it('carries the FS1 code when the newest source is an addon export', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'warrior-fury',
      gear: 'FS1:1.60.1.69893:warrior:human:0/0/0:',
      talents: '',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'addon',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.path).toBe(PATH);
    expect(result.code).toBe('FS1:1.60.1.69893:warrior:human:0/0/0:');
  });

  it('reads no export when the newest source is a logged fight', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'mage-fire',
      gear: { trinkets: [] },
      talents: '31/0/20',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'fight',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBeNull();
  });

  it('reads no export, not an error, when the API call fails', async () => {
    fetchSimInput.mockRejectedValueOnce(new Error('offline'));

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBeNull();
  });
});
```

`vi.mock` is hoisted above the import, so `fetchSimInput` (the outer `vi.fn()`) must be
declared before it — the factory closes over it rather than referencing
`vi.importActual`, since this module needs no other export from `./sim/api`. The dynamic
`await import('./addon-export')` after the mock is registered is required for the mock to
apply (a static top-of-file import would resolve before `vi.mock` hoists).

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/addon-export.test.ts`
Expected: FAIL — `Cannot find module './addon-export'`.

- [ ] **Step 3: Write minimal implementation**

```typescript
// web/src/lib/addon-export.ts
// Whether the API's newest source for a character is an addon export -- the one signal
// this site can turn into a working `/sim?source=addon&ref=` or `/planner?code=` link
// today. See the ruling in docs/superpowers/plans/2026-09-21-op-handoffs.md, Task 4: this
// makes the same read-only call the sim's own signed-in landing state makes
// (fetchSimInput, lib/sim/api.ts) and applies the same check fromStoredCharacter does
// (lib/sim/sources.ts) -- never a guess, never a second copy of that logic.
import { FS1_PREFIX } from './planner/fs1';
import type { CharacterPath } from './characters';
import { fetchSimInput } from './sim/api';

export interface AddonExportLookup {
  path: CharacterPath;
  /** The FS1 string, present only when the API's newest source for this character is the addon. */
  code: string | null;
}

export async function lookupAddonExport(
  path: CharacterPath,
  apiBase?: string,
): Promise<AddonExportLookup> {
  try {
    const input = await fetchSimInput(path, apiBase);
    const code =
      input.source === 'addon' && typeof input.gear === 'string' && input.gear.startsWith(`${FS1_PREFIX}:`)
        ? input.gear
        : null;
    return { path, code };
  } catch {
    return { path, code: null };
  }
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/addon-export.test.ts`
Expected: PASS, 3 tests.

- [ ] **Step 5: Write the shared component**

```svelte
<!-- web/src/components/CharacterHandoffLinks.svelte -->
<!-- Open in simulator / Open in planner for one character, or what is needed instead.
     Shared by Account.svelte's Characters list and Character.svelte's header: both need
     the exact same addon-export check (lib/addon-export.ts), so it lives once. -->
<script lang="ts">
  import { lookupAddonExport } from '../lib/addon-export';
  import { plannerCodeHref, simCodeHref } from '../lib/handoff-links';
  import type { CharacterPath } from '../lib/characters';

  let { path, apiBase = undefined }: { path: CharacterPath; apiBase?: string } = $props();

  let code = $state<string | null>(null);
  let status = $state<'loading' | 'ready'>('loading');

  // One load per `path`; the reference-equality guard is the same "requested vs resolved"
  // pattern Character.svelte's own effect uses, so a prop change mid-flight cannot land a
  // stale result over a newer one.
  $effect(() => {
    const requested = path;
    status = 'loading';
    code = null;
    void lookupAddonExport(requested, apiBase).then((result) => {
      if (result.path !== requested) return;
      code = result.code;
      status = 'ready';
    });
  });
</script>

<div
  class="flex min-h-11 flex-wrap items-center gap-3 md:min-h-0"
  data-testid="character-handoff"
>
  {#if status === 'loading'}
    <span class="invisible text-[13px]" aria-hidden="true">Open in simulator</span>
  {:else if code !== null}
    <a class="text-[13px] font-semibold" href={simCodeHref(code)} data-testid="character-open-sim">
      Open in simulator
    </a>
    <a class="text-[13px] font-semibold" href={plannerCodeHref(code)} data-testid="character-open-planner">
      Open in planner
    </a>
  {:else}
    <span class="text-muted text-[13px]" data-testid="character-needs-addon">
      Log in with the addon once to make this character simmable.
    </span>
  {/if}
</div>
```

- [ ] **Step 6: Scoped checks**

Run: `cd web && npx vitest run src/lib/addon-export.test.ts && npx astro check && npm run lint && npx prettier --check src/lib/addon-export.ts src/lib/addon-export.test.ts src/components/CharacterHandoffLinks.svelte`
Expected: no errors.

- [ ] **Step 7: Commit**

```bash
printf 'feat(handoffs): add the addon-export lookup and shared character hand-off links\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-4.txt
git add web/src/lib/addon-export.ts web/src/lib/addon-export.test.ts web/src/components/CharacterHandoffLinks.svelte
git commit -F .superpowers/commit-msg-4.txt
```

---

### Task 5: Wire hand-off links into `/account` and `/character/<key>`

**Files:**
- Modify: `web/src/components/Account.svelte`
- Modify: `web/src/components/Character.svelte`

**Interfaces:**
- Consumes: `CharacterHandoffLinks.svelte` (Task 4), `parseCharacterPath` from
  `web/src/lib/characters.ts` (already imported by `LandingState.svelte` the same way, read
  the comment there for why `MeCharacter.key` is already `<region>/<ruleset>/<slug>`).

- [ ] **Step 1: Read the current state of `Account.svelte`**

Read `web/src/components/Account.svelte` in full before editing (it was edited by the
coordinator for pairing/reports modes and `SignInPrompt.svelte` — keep every existing
behaviour). Confirm the Characters section still looks like:

```svelte
      <section class="flex flex-col gap-3">
        <h2 class="section-title text-[18px]">Characters</h2>
        {#if me!.characters.length === 0}
          <p class="text-muted text-[14px]">
            No characters linked yet. Sign in with Battle.net to link them.
          </p>
        {:else}
          <ul class="flex flex-col">
            {#each me!.characters as character (character.key)}
              <li class="border-line-soft flex min-h-11 items-center gap-3 border-b py-2 text-[14px]">
                <a href={characterHref(character.region, character.ruleset, character.name)}
                  >{character.name}</a
                >
                <span class="text-muted">
                  {rulesetLabel(character.ruleset)}
                  {character.region.toUpperCase()}
                </span>
              </li>
            {/each}
          </ul>
        {/if}
      </section>
```

If it has drifted from this, adapt the edit below to the real markup rather than
overwriting unrelated changes.

- [ ] **Step 2: Add the import and wire the row**

Add to the import block:

```typescript
  import CharacterHandoffLinks from './CharacterHandoffLinks.svelte';
  import { parseCharacterPath } from '../lib/characters';
```

Replace the `<li>` body with one that also renders the hand-off links (guard on
`parseCharacterPath` returning non-null, the same guard `LandingState.svelte` uses):

```svelte
            {#each me!.characters as character (character.key)}
              {@const path = parseCharacterPath(`/character/${character.key}`)}
              <li
                class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-2 text-[14px]"
              >
                <a href={characterHref(character.region, character.ruleset, character.name)}
                  >{character.name}</a
                >
                <span class="text-muted">
                  {rulesetLabel(character.ruleset)}
                  {character.region.toUpperCase()}
                </span>
                {#if path !== null}
                  <CharacterHandoffLinks {path} />
                {/if}
              </li>
            {/each}
```

- [ ] **Step 3: Wire `Character.svelte`**

Add to the import block:

```typescript
  import CharacterHandoffLinks from './CharacterHandoffLinks.svelte';
```

In the `<header>` block, after the `<p>` line ending in "ranked fights", add:

```svelte
      <CharacterHandoffLinks path={resolved} />
```

(`resolved` is already guaranteed non-null in this branch — the surrounding `{#if data !==
null && resolved !== null}` guard covers it.)

- [ ] **Step 4: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/components/Account.svelte src/components/Character.svelte`
Expected: no errors.

- [ ] **Step 5: Commit**

```bash
printf 'feat(handoffs): open in simulator/planner from /account and /character\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-5.txt
git add web/src/components/Account.svelte web/src/components/Character.svelte
git commit -F .superpowers/commit-msg-5.txt
```

---

### Task 6: Report pages gain a per-combatant Sim link

**Files:**
- Create: `web/src/lib/report/sim-link.ts`
- Test: `web/src/lib/report/sim-link.test.ts`
- Modify: `web/src/components/report/SummaryTab.svelte`
- Modify: `web/src/components/report/DeathsTab.svelte`
- Modify: `web/src/components/report/ReportView.svelte`

**Interfaces:**
- Consumes: `simFightHref` from Task 1's `handoff-links.ts`.
- Produces: `simLinkFor(reportId: string, fightIndex: number, guid: string): { href: string; label: string }`.
  `SummaryTab` and `DeathsTab` gain two new optional props, `reportId?: string` and
  `fightIndex?: number` — `ReportView.svelte` is the only caller of either and now passes
  both.

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/report/sim-link.test.ts
import { describe, expect, it } from 'vitest';
import { simLinkFor } from './sim-link';

describe('simLinkFor', () => {
  it('builds a fight-sourced sim link scoped to one combatant, labelled Sim', () => {
    const link = simLinkFor('abc2defg2hij', 3, 'Player-4184-000000A1');
    expect(link.href).toBe('/sim?source=fight&ref=abc2defg2hij%3A3%3APlayer-4184-000000A1');
    expect(link.label).toBe('Sim');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/report/sim-link.test.ts`
Expected: FAIL — `Cannot find module './sim-link'`.

- [ ] **Step 3: Write minimal implementation**

```typescript
// web/src/lib/report/sim-link.ts
// The per-combatant "Sim" link the report shows beside plannerLinkFor's own link
// (planner-link.ts): the same fight-sourced character, opened in the simulator instead of
// the planner. The ref extends the fight-level "Sim this fight" link's own ref shape
// (`<report_id>:<fight_index>`) with the combatant's GUID exactly as the report's roster
// carries it (docs/superpowers/specs/2026-09-21-one-product-design.md section 2) --
// Lane B's `fromLoggedFight` reads the optional third part.
import { simFightHref } from '../handoff-links';

export interface SimLink {
  href: string;
  label: string;
}

export function simLinkFor(reportId: string, fightIndex: number, guid: string): SimLink {
  return { href: simFightHref(reportId, fightIndex, guid), label: 'Sim' };
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/report/sim-link.test.ts`
Expected: PASS, 1 test.

- [ ] **Step 5: Wire `SummaryTab.svelte`**

Add the import:

```typescript
  import { simLinkFor } from '../../lib/report/sim-link';
```

Add two new optional props (alongside the existing `onTab` prop, in both the destructure
and the type):

```typescript
    reportId = undefined,
    fightIndex = undefined,
```

```typescript
    reportId?: string;
    fightIndex?: number;
```

In the combatants `{#each}` block, alongside the existing `{@const link = plannerLinkFor(...)}`,
add:

```svelte
          {@const simLink =
            reportId !== undefined && fightIndex !== undefined
              ? simLinkFor(reportId, fightIndex, combatant.guid)
              : null}
```

Immediately after the existing `{#if link}...{/if}` block (the `combatant-build-link`
anchor), add:

```svelte
            {#if simLink}
              <a
                class="text-gold inline-flex min-h-11 items-center text-[13px] md:min-h-0"
                href={simLink.href}
                data-testid="combatant-sim-link"
              >
                {simLink.label}
              </a>
            {/if}
```

(Keep `ml-auto` only on whichever of the two links renders first when the other is absent —
in practice `link` is `null` only for a class the planner has no data for, which is rare;
leaving `ml-auto` on the existing `link` anchor and adding the new one plain, immediately
after it, is enough: when `link` is null the new anchor loses the right-alignment, which is
acceptable and matches how `missing_buffs` already behaves when absent.)

- [ ] **Step 6: Wire `DeathsTab.svelte`**

Add the import:

```typescript
  import { simLinkFor } from '../../lib/report/sim-link';
```

Add the same two new optional props as Step 5 (destructure and type).

In the `{#each ordered as death}` block, alongside `{@const link = linkFor(death)}`, add:

```svelte
      {@const simLink =
        reportId !== undefined && fightIndex !== undefined
          ? simLinkFor(reportId, fightIndex, death.guid)
          : null}
```

Immediately after the existing `{#if link}...{/if}` block (the `death-build-link` anchor,
inside the `<div class="flex flex-wrap gap-2">`), add:

```svelte
            {#if simLink}
              <a
                class="border-line-warm-strong rounded-control text-strong inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
                href={simLink.href}
                data-testid="death-sim-link"
              >
                {simLink.label}
              </a>
            {/if}
```

- [ ] **Step 7: Wire `ReportView.svelte`**

Find the `<SummaryTab ... />` call (around the `state.tab === 'summary'` branch) and add
two props:

```svelte
          <SummaryTab
            summary={scoped}
            everyone={windowed ?? scoped}
            durationMs={scoped.duration_ms}
            {percentiles}
            {parseFallback}
            approximate={!windowIsWhole}
            dataBuild={activeBuild.build}
            {classOf}
            {treeSizesFor}
            onSelectPlayer={(guid) => patch({ source: guid })}
            players={playerSet}
            onTab={(tab) => patch({ tab })}
            {reportId}
            fightIndex={state.fight}
          />
```

Find the `<DeathsTab ... />` call (around line 1668, the `state.tab === 'deaths'` branch)
and add the same two props:

```svelte
          <DeathsTab
            deaths={scoped.deaths}
            casts={base?.casts ?? []}
            open={state.openDeaths}
            onPatch={patch}
            onSelectPlayer={(guid) => patch({ source: guid })}
            durationMs={summary?.duration_ms ?? scoped.duration_ms}
            pulls={scoped.pulls ?? []}
            combatants={scoped.combatants}
            {classOf}
            dataBuild={activeBuild.build}
            {treeSizesFor}
            onWindow={setWindow}
            {reportId}
            fightIndex={state.fight}
          />
```

- [ ] **Step 8: Scoped checks**

Run: `cd web && npx vitest run src/lib/report/sim-link.test.ts && npx astro check && npm run lint && npx prettier --check src/lib/report/sim-link.ts src/lib/report/sim-link.test.ts src/components/report/SummaryTab.svelte src/components/report/DeathsTab.svelte src/components/report/ReportView.svelte`
Expected: no errors.

- [ ] **Step 9: Commit**

```bash
printf 'feat(report): add a per-combatant Sim link beside the planner link\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-6.txt
git add web/src/lib/report/sim-link.ts web/src/lib/report/sim-link.test.ts web/src/components/report/SummaryTab.svelte web/src/components/report/DeathsTab.svelte web/src/components/report/ReportView.svelte
git commit -F .superpowers/commit-msg-6.txt
```

---

### Task 7: `/guides/<class>` links to the planner

**Files:**
- Modify: `web/src/pages/guides/[slug].astro`

**Interfaces:**
- Consumes: `plannerHref` from `web/src/lib/planner/reference.ts` (read-only import —
  already the exact pattern `web/src/pages/classes.astro` uses).

- [ ] **Step 1: Edit the page**

Replace the file's contents:

```astro
---
// web/src/pages/guides/[slug].astro
import { getCollection, render } from 'astro:content';
import Content from '../../layouts/Content.astro';
import { plannerHref } from '../../lib/planner/reference';
export async function getStaticPaths() {
  return (await getCollection('guides')).map((g) => ({ params: { slug: g.id }, props: { entry: g } }));
}
const { entry } = Astro.props;
const { Content: Body } = await render(entry);
---

<Content
  title={entry.data.title}
  description={entry.data.description}
  path={`/guides/${entry.id}`}
  updated={entry.data.updated}
  confidence={entry.data.confidence}
  sources={entry.data.sources}
>
  <p>
    <a href={plannerHref(entry.data.classSlug)}>Open the planner for this class</a>
  </p>
  <Body />
</Content>
```

- [ ] **Step 2: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/pages/guides/\[slug\].astro`
Expected: no errors.

- [ ] **Step 3: Commit**

```bash
printf 'feat(guides): link every class guide to the planner for that class\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-7.txt
git add "web/src/pages/guides/[slug].astro"
git commit -F .superpowers/commit-msg-7.txt
```

---

### Task 8: `/dungeons/<slug>` links its zone and (where loot exists) the drops tool

**Files:**
- Create: `web/src/lib/reference/dungeon-links.ts`
- Test: `web/src/lib/reference/dungeon-links.test.ts`
- Modify: `web/src/pages/dungeons/[slug].astro`

**Ruling (record in the ledger verbatim):** most dungeons' `zone` field names an old-world
zone (`Dun Morogh`, `Alcaz Island`, …) that this site's `zones` content collection has no
page for — only four new zones exist there. The zone name is linked only when it exactly
matches a `zones` collection entry's title; otherwise it stays plain text, which is honest
rather than a broken or misleading link. Likewise, today's `data/builds/<build>/loot.json`
carries old-world dungeon ids (`dungeon:blackrock-depths`, …), none of which match the new
Forever dungeon content slugs — so the "See drops for your character" link does not render
on any dungeon page yet under real data. It does render under the `FOREVER_DATA=fixture`
build this repo tests against (the fixture's `loot.json` deliberately carries
`dungeon:hall-of-thanes`, matching `src/content/dungeons/hall-of-thanes.md`), which is how
Task 10's e2e spec proves the wiring works, ready for when the data pipeline catches up.

**Interfaces:**
- Consumes: `simDropsHref` from Task 1's `handoff-links.ts`.
- Produces: `zoneEntryForDungeon(zones, dungeonZone): ZoneEntry | null`,
  `dungeonHasLoot(loot, slug): boolean` — both pure, tested directly; the page's own
  frontmatter does the (untested, thin) file read that feeds `dungeonHasLoot`.

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/lib/reference/dungeon-links.test.ts
import { describe, expect, it } from 'vitest';
import { dungeonHasLoot, zoneEntryForDungeon, type LootFile, type ZoneEntry } from './dungeon-links';

const ZONES: ZoneEntry[] = [
  { id: 'mount-hyjal', title: 'Mount Hyjal' },
  { id: 'riverglades', title: 'Riverglades' },
];

describe('zoneEntryForDungeon', () => {
  it('matches a zone by its exact title', () => {
    expect(zoneEntryForDungeon(ZONES, 'Riverglades')).toEqual({ id: 'riverglades', title: 'Riverglades' });
  });

  it('returns null for an old-world zone this site has no page for', () => {
    expect(zoneEntryForDungeon(ZONES, 'Dun Morogh')).toBeNull();
  });

  it('returns null for "Not yet known"', () => {
    expect(zoneEntryForDungeon(ZONES, 'Not yet known')).toBeNull();
  });
});

describe('dungeonHasLoot', () => {
  const loot: LootFile = {
    sources: [
      { id: 'dungeon:hall-of-thanes', kind: 'dungeon' },
      { id: 'raid:molten-core', kind: 'raid' },
    ],
  };

  it('is true when the loot file records a dungeon source at this slug', () => {
    expect(dungeonHasLoot(loot, 'hall-of-thanes')).toBe(true);
  });

  it('is false for a slug the loot file has no dungeon source for', () => {
    expect(dungeonHasLoot(loot, 'alcaz-prison')).toBe(false);
  });

  it('is false for a source of a different kind at the same-looking id', () => {
    expect(dungeonHasLoot(loot, 'molten-core')).toBe(false);
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd web && npx vitest run src/lib/reference/dungeon-links.test.ts`
Expected: FAIL — `Cannot find module './dungeon-links'`.

- [ ] **Step 3: Write minimal implementation**

```typescript
// web/src/lib/reference/dungeon-links.ts
// Two honest conditionals /dungeons/<slug> needs before it links anywhere: whether its
// zone name matches one of the (today, four) zones this site has an atlas page for, and
// whether the active build's loot data has caught up with this dungeon yet. Both pure —
// the astro page reads the zones collection and the build's loot.json itself and hands the
// rows in here, so this module needs no astro:content or filesystem access of its own.

export interface ZoneEntry {
  id: string;
  title: string;
}

/**
 * The zones collection entry whose title matches this dungeon's zone name exactly, or
 * null — most dungeons still name an old-world zone this site has no atlas page for, and a
 * few (`zone === 'Not yet known'`) name nothing at all yet.
 */
export function zoneEntryForDungeon(zones: readonly ZoneEntry[], dungeonZone: string): ZoneEntry | null {
  return zones.find((zone) => zone.title === dungeonZone) ?? null;
}

export interface LootSource {
  id: string;
  kind: string;
}

export interface LootFile {
  sources: LootSource[];
}

/**
 * Whether the loot file records any loot for the dungeon at this slug — `loot.json`'s own
 * `dungeon:<slug>` id (the pipeline's own vocabulary, `data/builds/<build>/loot.json`).
 */
export function dungeonHasLoot(loot: LootFile, slug: string): boolean {
  return loot.sources.some((source) => source.kind === 'dungeon' && source.id === `dungeon:${slug}`);
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd web && npx vitest run src/lib/reference/dungeon-links.test.ts`
Expected: PASS, 6 tests.

- [ ] **Step 5: Wire the page**

Replace the contents of `web/src/pages/dungeons/[slug].astro`:

```astro
---
// web/src/pages/dungeons/[slug].astro
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { getCollection, render } from 'astro:content';
import Content from '../../layouts/Content.astro';
import { levelRangeLong } from '../../lib/levels';
import activeBuild from '../../data/active-build.json';
import { simDropsHref } from '../../lib/handoff-links';
import { dungeonHasLoot, zoneEntryForDungeon, type LootFile } from '../../lib/reference/dungeon-links';

export async function getStaticPaths() {
  return (await getCollection('dungeons')).map((d) => ({ params: { slug: d.id }, props: { entry: d } }));
}
const { entry } = Astro.props;
const { Content: Body } = await render(entry);
const range = levelRangeLong(entry.data.levelMin, entry.data.levelMax);

const zones = (await getCollection('zones')).map((z) => ({ id: z.id, title: z.data.title }));
const zoneEntry = zoneEntryForDungeon(zones, entry.data.zone);

/**
 * public/data/<build>/loot.json is generated by scripts/sync-data.mjs (`npm run sync`)
 * before every build and test run; a build the data lane has not regenerated it for ships
 * none, the same optional-file contract sync-data.mjs documents for every entry in
 * SYNC_ENTRIES — so a missing or unparsable file reads as "no loot data yet", not an error.
 */
function readActiveLoot(): LootFile | null {
  try {
    const dataDir = fileURLToPath(new URL('../../../public/data', import.meta.url));
    const raw = readFileSync(`${dataDir}/${activeBuild.build}/loot.json`, 'utf8');
    return JSON.parse(raw) as LootFile;
  } catch {
    return null;
  }
}
const loot = readActiveLoot();
const hasLoot = loot !== null && dungeonHasLoot(loot, entry.id);
---

<Content
  title={entry.data.title}
  description={`${entry.data.zone} · ${range}`}
  path={`/dungeons/${entry.id}`}
  updated={entry.data.updated}
  confidence={entry.data.confidence}
  sources={entry.data.sources}
>
  <p>
    <strong>
      {zoneEntry !== null ? <a href={`/zones/${zoneEntry.id}`}>{entry.data.zone}</a> : entry.data.zone}
    </strong>{' '}
    · {range}
  </p>
  <Body />
  {hasLoot && (
    <p>
      <a href={simDropsHref(entry.id)} data-testid="dungeon-drops-link">
        See drops for your character
      </a>
    </p>
  )}
</Content>
```

- [ ] **Step 6: Scoped checks**

Run (sync first, so `public/data/<build>/loot.json` exists for `astro check`'s type pass
and for the manual smoke check below):

```bash
cd web && export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12 && FOREVER_DATA=fixture npm run sync
npx vitest run src/lib/reference/dungeon-links.test.ts && npx astro check && npm run lint && npx prettier --check src/lib/reference/dungeon-links.ts src/lib/reference/dungeon-links.test.ts "src/pages/dungeons/[slug].astro"
```

Expected: no errors.

- [ ] **Step 7: Commit**

```bash
printf 'feat(dungeons): link a dungeon'"'"'s zone and its drops when loot data exists\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-8.txt
git add web/src/lib/reference/dungeon-links.ts web/src/lib/reference/dungeon-links.test.ts "web/src/pages/dungeons/[slug].astro"
git commit -F .superpowers/commit-msg-8.txt
```

---

### Task 9: `/zones` says which zones are covered

**Files:**
- Modify: `web/src/pages/zones/index.astro`

- [ ] **Step 1: Add the coverage line**

In `web/src/pages/zones/index.astro`, after the existing `<p class="text-muted text-[14px]">`
(the "Level ranges are reported…" line), add a second paragraph built from the same `zones`
array already fetched, so it can never drift from what the page actually lists:

```astro
    <p class="text-muted text-[14px]">
      Level ranges are reported where a source gives them and not yet confirmed by Blizzard unless noted.
    </p>
    <p class="text-muted text-[14px]">
      {zones.length} {zones.length === 1 ? 'zone is' : 'zones are'} covered today:{' '}
      {zones.map((z) => z.data.title).join(', ')}. The rest of the world's zones are coming as this atlas
      grows.
    </p>
```

- [ ] **Step 2: Scoped checks**

Run: `cd web && npx astro check && npm run lint && npx prettier --check src/pages/zones/index.astro`
Expected: no errors.

- [ ] **Step 3: Commit**

```bash
printf 'feat(zones): say which zones are covered and that the rest are coming\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-9.txt
git add web/src/pages/zones/index.astro
git commit -F .superpowers/commit-msg-9.txt
```

---

### Task 10: E2E coverage for every hand-off (desktop and mobile)

**Files:**
- Create: `web/tests/e2e/handoffs-home.spec.ts`
- Create: `web/tests/e2e/handoffs-addon.spec.ts`
- Create: `web/tests/e2e/handoffs-account-character.spec.ts`
- Create: `web/tests/e2e/handoffs-report.spec.ts`
- Create: `web/tests/e2e/handoffs-reference.spec.ts`

**Interfaces:**
- Consumes: everything from Tasks 2–9. `ACTIVE_BUILD` from `tests/e2e/support/active-build.ts`.

- [ ] **Step 1: Homepage grid**

```typescript
// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from '@playwright/test';

test('the homepage tools grid carries Simulator and The addon, in spec order', async ({ page }) => {
  await page.goto('/');
  const grid = page.getByTestId('tools-grid');
  const cards = grid.getByRole('link');
  const titles = await cards.evaluateAll((links) =>
    links.map((a) => a.querySelector('.font-display')?.textContent?.trim() ?? ''),
  );
  expect(titles).toEqual([
    'Build planner',
    'Simulator',
    'Combat logs',
    'Rankings',
    'The addon',
    'Dungeons',
    'Zone atlas',
    'Class guides',
  ]);
  await expect(grid.getByRole('link', { name: /Simulator/ })).toHaveAttribute('href', '/sim');
  await expect(grid.getByRole('link', { name: /The addon/ })).toHaveAttribute('href', '/addon');
});
```

- [ ] **Step 2: `/addon` paste box**

```typescript
// web/tests/e2e/handoffs-addon.spec.ts
import { expect, test } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';

test('the /addon paste box offers the planner and the simulator once an export decodes', async ({ page }) => {
  await page.goto('/addon');
  await page
    .getByTestId('addon-paste-code')
    .fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`);
  await page.getByTestId('addon-paste-submit').click();

  await expect(page.getByTestId('addon-paste-error')).toHaveCount(0);
  const planner = page.getByTestId('addon-paste-planner');
  const sim = page.getByTestId('addon-paste-sim');
  await expect(planner).toHaveAttribute(
    'href',
    `/planner?code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
  await expect(sim).toHaveAttribute(
    'href',
    `/sim?code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
});

test('a code from another format is refused by name, not silently dropped', async ({ page }) => {
  await page.goto('/addon');
  await page.getByTestId('addon-paste-code').fill('FS2:nope');
  await page.getByTestId('addon-paste-submit').click();
  await expect(page.getByTestId('addon-paste-error')).toHaveText('That code is FS2; this site reads FS1.');
  await expect(page.getByTestId('addon-paste-planner')).toHaveCount(0);
});

test('the page says the in-game UI is in beta testing and ships no screenshot of it', async ({ page }) => {
  await page.goto('/addon');
  await expect(page.getByText('in beta testing in game')).toBeVisible();
  const images = await page.locator('main img').count();
  expect(images).toBe(0);
});
```

- [ ] **Step 3: `/account` and `/character/<key>`**

```typescript
// web/tests/e2e/handoffs-account-character.spec.ts
import { expect, test } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const ADDON_CODE = `FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`;

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [
    { key: 'us/normal/thrallgar', region: 'us', ruleset: 'normal', name: 'Thrallgar', class: 'Warrior' },
    { key: 'us/normal/roland', region: 'us', ruleset: 'normal', name: 'Roland', class: 'Mage' },
  ],
  guilds: [],
};

test('a signed-in member sees hand-off links only for a character with an addon export', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope({ devices: [] })));
  await page.route('**/v1/characters/us/normal/thrallgar/sim-input', (route) =>
    route.fulfill(
      envelope({
        spec: 'warrior-fury',
        gear: ADDON_CODE,
        talents: '',
        buffs: [],
        captured_at: new Date().toISOString(),
        source: 'addon',
      }),
    ),
  );
  await page.route('**/v1/characters/us/normal/roland/sim-input', (route) =>
    route.fulfill(
      envelope({
        spec: 'mage-fire',
        gear: { trinkets: [] },
        talents: '31/0/20',
        buffs: [],
        captured_at: new Date().toISOString(),
        source: 'fight',
      }),
    ),
  );

  await page.goto('/account');

  const thrallgarRow = page.getByRole('listitem').filter({ hasText: 'Thrallgar' });
  await expect(thrallgarRow.getByTestId('character-open-sim')).toHaveAttribute(
    'href',
    `/sim?code=${encodeURIComponent(ADDON_CODE)}`,
  );
  await expect(thrallgarRow.getByTestId('character-open-planner')).toHaveAttribute(
    'href',
    `/planner?code=${encodeURIComponent(ADDON_CODE)}`,
  );

  const rolandRow = page.getByRole('listitem').filter({ hasText: 'Roland' });
  await expect(rolandRow.getByTestId('character-needs-addon')).toHaveText(
    'Log in with the addon once to make this character simmable.',
  );
});
```

- [ ] **Step 4: Report per-combatant Sim link**

```typescript
// web/tests/e2e/handoffs-report.spec.ts
import { expect, test } from '@playwright/test';

test('each combatant in a report has its own Sim link beside the planner link', async ({ page }) => {
  // Reuses the same fixture report every other report-*.spec.ts file runs against.
  await page.goto('/reports/fixture2abcd?fight=3');
  const combatants = page.getByTestId('combatants').getByRole('listitem');
  await expect(combatants.first()).toBeVisible();
  const first = combatants.first();
  const simLink = first.getByTestId('combatant-sim-link');
  await expect(simLink).toBeVisible();
  const href = await simLink.getAttribute('href');
  expect(href).toMatch(/^\/sim\?source=fight&ref=fixture2abcd%3A3%3A/);
});
```

Confirmed against `tests/e2e/report-deaths.spec.ts` and `tests/e2e/report-tabs.spec.ts`,
which navigate to the same fixture report the same way
(`/reports/fixture2abcd?fight=3...`): `fixture2abcd` and fight `3` are correct as written
above.

- [ ] **Step 5: Reference pages**

```typescript
// web/tests/e2e/handoffs-reference.spec.ts
import { expect, test } from '@playwright/test';

test('a class guide links to the planner for that class', async ({ page }) => {
  await page.goto('/guides/warrior');
  const link = page.getByRole('link', { name: 'Open the planner for this class' });
  await expect(link).toHaveAttribute('href', '/planner?class=warrior');
});

test('a dungeon whose zone this site covers links to that zone', async ({ page }) => {
  await page.goto('/dungeons/kroldok-stronghold');
  await expect(page.getByRole('link', { name: 'Riverglades' })).toHaveAttribute(
    'href',
    '/zones/riverglades',
  );
});

test('a dungeon whose zone this site does not cover names it without a broken link', async ({ page }) => {
  await page.goto('/dungeons/hall-of-thanes');
  await expect(page.getByText('Dun Morogh')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Dun Morogh' })).toHaveCount(0);
});

test('a dungeon with fixture loot data offers drops for your character', async ({ page }) => {
  await page.goto('/dungeons/hall-of-thanes');
  await expect(page.getByTestId('dungeon-drops-link')).toHaveAttribute(
    'href',
    '/sim/drops?instance=hall-of-thanes',
  );
});

test('/zones names which zones are covered and that the rest are coming', async ({ page }) => {
  await page.goto('/zones');
  await expect(page.getByText('The rest of the world', { exact: false })).toBeVisible();
});
```

- [ ] **Step 6: Run every new spec, desktop and mobile**

```bash
cd web
export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12
npx astro preview stop || true
FOREVER_DATA=fixture npm run sync
E2E_PORT=4412 npx playwright test tests/e2e/handoffs-home.spec.ts tests/e2e/handoffs-addon.spec.ts tests/e2e/handoffs-account-character.spec.ts tests/e2e/handoffs-report.spec.ts tests/e2e/handoffs-reference.spec.ts
```

Expected: all pass on both the `desktop` and `mobile` projects. Fix any selector or href
mismatch against the real rendered markup (in particular Step 4's report id/fight number
and Step 1's `.font-display` title selector) before moving on — these are the kind of
detail that only the real DOM can confirm.

- [ ] **Step 7: Scoped checks and commit**

```bash
cd web && npx prettier --check tests/e2e/handoffs-*.spec.ts
```

```bash
printf 'test(e2e): cover every hand-off this lane adds, desktop and mobile\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\n' > .superpowers/commit-msg-10.txt
git add web/tests/e2e/handoffs-home.spec.ts web/tests/e2e/handoffs-addon.spec.ts web/tests/e2e/handoffs-account-character.spec.ts web/tests/e2e/handoffs-report.spec.ts web/tests/e2e/handoffs-reference.spec.ts
git commit -F .superpowers/commit-msg-10.txt
```

---

### Task 11: Whole-branch review and one fix wave

Not a code task — this is the `subagent-driven-development` skill's own final step. Run:

```bash
cd web
export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12
npx vitest run src/lib/handoff-links.test.ts src/lib/addon-export.test.ts src/lib/report/sim-link.test.ts src/lib/reference/dungeon-links.test.ts
npx astro check
npm run lint
npx prettier --check src/lib/handoff-links.ts src/lib/addon-export.ts src/components/AddonPasteBox.svelte src/components/CharacterHandoffLinks.svelte src/lib/report/sim-link.ts src/lib/reference/dungeon-links.ts src/components/Account.svelte src/components/Character.svelte src/components/report/SummaryTab.svelte src/components/report/DeathsTab.svelte src/components/report/ReportView.svelte "src/pages/guides/[slug].astro" "src/pages/dungeons/[slug].astro" src/pages/zones/index.astro src/pages/addon.astro src/pages/index.astro src/data/tools.json src/lib/addon/copy.ts
E2E_PORT=4412 npx playwright test tests/e2e/handoffs-home.spec.ts tests/e2e/handoffs-addon.spec.ts tests/e2e/handoffs-account-character.spec.ts tests/e2e/handoffs-report.spec.ts tests/e2e/handoffs-reference.spec.ts tests/e2e/addon.spec.ts tests/e2e/content.spec.ts tests/e2e/home.spec.ts
```

Dispatch one reviewer subagent against the whole branch diff (`git diff main...op-handoffs`
in the worktree) checking: every spec-2 bullet has a shipped file; no file under the
excluded directories was touched (`git diff --name-only main...op-handoffs` against the
file-ownership list in Global Constraints); no magic strings outside `addonCopy` for new
`/addon` copy; every new async component degrades honestly on failure (no thrown error
reaches the DOM); no layout shift was introduced (loading states reserve their space, per
Task 4/5's `invisible` placeholder). Apply one fix round for anything CRITICAL or HIGH per
`code-review.md`'s severity table, then re-run the block above.

Write the ledger at `.superpowers/sdd/2026-09-21-op-handoffs/progress.md` with one line per
task (`Ruling: <decision> — <why> — <cost if wrong>`) for at least the two rulings recorded
in Tasks 4 and 8, plus any the reviewer's fix round produced.
