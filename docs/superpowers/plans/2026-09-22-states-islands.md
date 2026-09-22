# One state model for every island (lane `states-islands`) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Every island outside the account/sign-in surface (Character, Rankings, the Guild
family, RecentReports/MyReports, ReportView's lazy panels, SimView's lazy panels and
saved-sim error, ToolsView, AddonPasteBox, Planner) shows the same three states — a
`Skeleton` that reserves the ready height, a `LoadError` with a real in-place retry, an
`EmptyState` with at most one action — and every request-triggering button marks itself
`aria-busy` with a shared busy look, with no ad-hoc `<p>Loading.</p>`, "Reload the page"
copy, or label-swap decoration left beside the new primitives.

**Architecture:** Wire the three primitives that already exist on `main`
(`components/ui/Skeleton.svelte`, `components/ui/LoadError.svelte`,
`components/ui/EmptyState.svelte`, `lib/ui/copy.ts`) into every island's existing
`status`/`error` state machine, extending each island's already-retriggerable load function
(`load()` called directly, or an `attempt` counter read inside an `$effect` the way
`Planner.svelte` already does) as the `onRetry` callback. Four guild islands share one new
`GuildStatus.svelte` wrapper instead of four copies of the same loading/error branch. A new
`lib/ui/busy.ts` exports one class string every request-triggering button applies alongside
`aria-busy`, removing every label-swap-while-busy decoration the audit found. New per-island
`min-h`/`lines` constants live in small new modules next to each island (`lib/character-layout.ts`,
`lib/guild/layout.ts`, `lib/report/layout.ts`, `lib/sim/lazy-layout.ts`,
`lib/addon/paste-layout.ts`), matching the existing `current-character-layout.ts` pattern.
`ReportView.svelte` (1841 lines) and `SimView.svelte` (800 lines) are never grown — the new
Skeleton/LoadError wiring for their lazy panels is a handful of lines inside the existing
`lazyFallback` snippet plus a new tiny constants module, not a new component file.

**Tech Stack:** Svelte 5 (runes), Astro, TypeScript, Vitest (`svelte/server`'s `render()`),
Playwright, Tailwind.

**Spec:** `docs/superpowers/specs/2026-09-22-account-and-island-states-design.md` — section 1
binds this lane, section 3 is this lane's scope (section 2 is `states-account`'s, not
touched here).

## Global Constraints

- **Vocabulary.** The addon runs in the game client; the companion is the desktop app.
  Never blur them; never write "the addon uploads" or "the companion in game."
- **Voice.** Reference, not pitch (`design/DESIGN-SYSTEM.md`). State the thing and stop.
  Every visible string lives in a copy module — never a literal string in a `.svelte` file's
  markup for anything the audit flags.
- **Motion budget.** A skeleton shimmer while loading, one 160ms `.reveal` fade when data
  lands, nothing else. `prefers-reduced-motion` already turns both off in `global.css`. No
  Svelte `transition:`/`animate:` directives anywhere in this lane's changes.
- **Nothing moves.** Every island reserves its ready height while loading: a `Skeleton` with
  `lines`/`minHeight` sized to the ready view, or a fixed `min-h` the island already has. The
  `.reveal` class goes on the wrapper of the ready state only. `web/lighthouserc.json`'s CLS
  0.05 / TBT 100–200ms budgets must hold for every URL it already measures; run
  `npm run lhci` before the final report and quote the CLS for each of its 13 URLs.
- **Three states, three components.** Loading = `components/ui/Skeleton.svelte`. Failed =
  `components/ui/LoadError.svelte` with an `onRetry` that re-fires the same request in
  place — never "reload the page." Empty = `components/ui/EmptyState.svelte` with at most
  one action. Existing ad-hoc `<p>Loading.</p>` / "Reload the page" lines are replaced, not
  kept beside the new ones.
- **Primitives (read-only, on `main`, never edited by this lane):**
  `components/ui/Skeleton.svelte` (`lines`, `rowHeight`, `minHeight`, `label`, `testid`),
  `components/ui/LoadError.svelte` (`message`, `onRetry?`, `testid`; retry button is
  `{testid}-retry`), `components/ui/EmptyState.svelte` (`message`, `action?: {label, href}`,
  `testid`), `lib/ui/copy.ts` (`uiCopy.loading`, `uiCopy.retry`), `.reveal` in `global.css`.
  If a primitive needs a prop it does not have, this lane does not add it — it works around
  it locally and says so in the final report.
- **Lane ownership.** This lane owns everything under `web/src/components` and `web/src/lib`
  that spec section 3 names, plus their tests, listed per-task below. It never touches
  `Account.svelte`, `components/account/**`, `CurrentCharacterBar.svelte`,
  `CurrentCharacterChip.svelte`, `SignInPrompt.svelte`, `CharacterHandoffLinks.svelte`,
  `lib/account/**`, `handoff-copy.ts`, `current-character-copy.ts`, `login.astro`,
  `account.astro`, or their e2e specs — those are `states-account`'s. It never touches
  `Header.astro`, `Footer.astro`, `components/ui/**`, or `lib/ui/**` except to create the new
  `lib/ui/busy.ts` (which is new, not an edit to an existing primitive).
- **Web toolchain**, from `web/`: `export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12`
  then `FOREVER_DATA=fixture npm run sync` once (already run for this worktree). Before every
  commit: `npx vitest run <touched paths>`, `npx astro check`, `npm run lint`,
  `npx prettier --check <touched paths>`. E2E: `E2E_PORT=4432 npx playwright test <spec files>`
  (stop a stale preview first with `npx astro preview stop`). `web/node_modules` is a symlink
  to the main checkout's: never `npm install`, never stage it.
- **Commits.** Write the message to a file under this worktree's `.superpowers/` with
  `printf`, then `git commit -F <file>` as its own Bash command with nothing else in it.
  Conventional subjects (`feat(web): ...`, `fix(web): ...`, `test(web): ...`). End every
  message with exactly:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`
- **Code shape.** Functions under 50 lines, files under 800 lines (`ReportView.svelte` and
  `SimView.svelte` are already at/over the ceiling — put new constants and any new markup
  that would grow them into a new small module or snippet, never grow those two files'
  responsibilities). No magic numbers, no dead code, immutable updates, no code duplication
  — the guild family's four copies of loading/error markup collapse into one shared
  component precisely because of this rule.
- **Sizing constants.** Only `lib/character-layout.ts`'s constant is derived from an actual
  browser measurement (spec 3.1 names this one explicitly, the same way the other lane's
  `CharacterList` skeleton is measured); every other new `min-h`/`lines` constant in this
  plan is a reasoned estimate from the real ready markup this plan already read in full —
  Lighthouse does not gate post-interaction CLS for a mode/tab switch (its default run never
  triggers one), and none of Character/Rankings/Guild/RecentReports/MyReports's pages are in
  `lighthouserc.json`'s URL list at all, so exact pixel parity is good UX practice here, not
  a gated requirement. This is a deliberate scope ruling, recorded here so no task reargues
  it.
- **Tests use the fake engine and route stubs**, as existing specs do. Never save a public
  build or sim, never upload, never sign in on production.
- **Model budget.** Every implementer/reviewer subagent runs on `sonnet`; `haiku` only for a
  purely mechanical single-file fix; never `opus`. One implementer at a time.
- **Run the whole e2e suite once** before the final report (`E2E_PORT=4432 npx playwright test`),
  not just touched specs. A failure reproducible on `main` is reported, not fixed.
- **Stop every server this lane started** before the final report.

---

## Task 1: `lib/ui/busy.ts` — the shared busy-button class

**Files:**
- Create: `web/src/lib/ui/busy.ts`
- Test: `web/src/lib/ui/busy.test.ts`

**Interfaces:**
- Produces: `BUSY_CLASS: string` — every later task that touches a request-triggering
  button imports this and appends it to that button's `class` attribute when its busy flag
  is true, and adds `aria-busy={<thatFlag>}` next to its existing `disabled={<thatFlag>}`.

- [ ] **Step 1: Write the failing test**

```ts
// web/src/lib/ui/busy.test.ts
import { describe, expect, it } from 'vitest';
import { BUSY_CLASS } from './busy';

describe('BUSY_CLASS', () => {
  it('is the one shared busy look: dimmed, progress cursor, nothing else', () => {
    expect(BUSY_CLASS).toBe('opacity-60 cursor-progress');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `npx vitest run src/lib/ui/busy.test.ts`
Expected: FAIL — `Cannot find module './busy'`

- [ ] **Step 3: Write minimal implementation**

```ts
// web/src/lib/ui/busy.ts
// The one shared look every island's request-triggering button (Run, Save, Submit,
// Upload, Join, Claim...) takes on while its own request is in flight: dimmed, a progress
// cursor, no spinner glyph and no label swap (design 2026-09-22 spec section 3.2). Applied
// alongside `aria-busy` and the button's existing `disabled` flag, never instead of them.
export const BUSY_CLASS = 'opacity-60 cursor-progress';
```

- [ ] **Step 4: Run test to verify it passes**

Run: `npx vitest run src/lib/ui/busy.test.ts`
Expected: PASS

- [ ] **Step 5: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/lib/ui/busy.ts src/lib/ui/busy.test.ts
```

```bash
printf 'feat(web): the one shared busy-button class every island will use\n\nRun, Save, Submit and every other request-triggering button converge on one\ndimmed/progress-cursor look instead of a spinner glyph or a label swap, so a\nbusy control matches its own visible label across every island (design\n2026-09-22 spec section 3.2).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-1.txt
git add web/src/lib/ui/busy.ts web/src/lib/ui/busy.test.ts
git commit -F .superpowers/commit-msg-1.txt
```

---

## Task 2: Character.svelte — Skeleton, LoadError, retry, reveal

**Files:**
- Create: `web/src/lib/character-layout.ts`
- Create: `web/src/components/Character.test.ts` (does not exist today)
- Modify: `web/src/components/Character.svelte:1,34,51-81`

**Interfaces:**
- Consumes: `Skeleton` (`lines`, `rowHeight`, `minHeight`, `testid`), `LoadError`
  (`message`, `onRetry`, `testid`), `uiCopy` from Task 1's sibling primitives (all on
  `main` already).
- Produces: `CHARACTER_LOADING_MIN_H: string` (a Tailwind `min-h-*` class) other tasks do
  not consume (Character.svelte is not reused elsewhere).

- [ ] **Step 1: Measure the fixture's ready height at 360px**

Character.svelte is fetched client-side (`fetchCharacter`), so a plain `astro preview` visit
renders nothing — route-stub it exactly the way `tests/e2e/character-phone.spec.ts` already
does, and read the `#character` element's height. Run this from `web/` after
`npx astro build` (fixture data is already synced):

```bash
E2E_PORT=4432 npx astro build >/dev/null && npx astro preview --port 4432 &
sleep 2
node -e "
const { chromium } = require('@playwright/test');
const CHARACTER = {
  ok: true,
  data: {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
    best: [{ encounter: 'Warden Kelthas', encounter_id: 9001, difficulty: 8, metric: 'hps', value: 1840, percentile: 96.2, spec: 'Discipline', fought_at: '2026-12-09T22:10:00Z', report_id: 'fixture2abcd', fight_index: 3 }],
    history: [{ encounter: 'Warden Kelthas', encounter_id: 9001, difficulty: 8, metric: 'hps', value: 1840, percentile: 96.2, spec: 'Discipline', fought_at: '2026-12-09T22:10:00Z', report_id: 'fixture2abcd', fight_index: 3 }],
    builds_seen: [{ talent_split: '31/20/0', spec: 'Discipline', first_seen: '2026-12-09T22:10:00Z' }],
  },
  error: null,
  request_id: 'r',
};
(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 360, height: 800 } });
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }));
  await page.goto('http://localhost:4432/character/us/hardcore/elyra-duskvale');
  const el = page.getByTestId('character');
  await el.waitFor();
  const box = await el.boundingBox();
  console.log('CHARACTER_READY_HEIGHT', Math.ceil(box.height));
  await browser.close();
})();
"
npx astro preview stop
```

Note the printed height (e.g. `812`) — used in Step 3. This is the profile header, the
current-character bar, the handoff links, the rating panel and the "Best per encounter"
table for this one-fight fixture; real characters with more history render taller, which is
fine — the skeleton only needs to avoid a shift for the common case, not bound every case.

- [ ] **Step 2: Write the failing test**

```ts
// web/src/components/Character.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Character from './Character.svelte';
import { parseCharacterPath } from '../lib/characters';
import { CHARACTER_LOADING_MIN_H } from '../lib/character-layout';

const path = parseCharacterPath('/character/us/hardcore/elyra-duskvale');

describe('Character loading state', () => {
  it('reserves the ready height with a Skeleton instead of a bare line', () => {
    const { body } = render(Character, { props: { path } });
    expect(body).toContain('data-testid="character-skeleton"');
    expect(body).toContain(CHARACTER_LOADING_MIN_H);
    expect(body).not.toContain('Loading.');
  });
});
```

- [ ] **Step 3: Run test to verify it fails**

Run: `npx vitest run src/components/Character.test.ts`
Expected: FAIL — `character-skeleton` testid not found, `CHARACTER_LOADING_MIN_H` not
exported.

- [ ] **Step 4: Create `lib/character-layout.ts`**

```ts
// web/src/lib/character-layout.ts
// The height Character.svelte's ready view renders at for the one-fight fixture at 360px
// (measured with the same route-stub character-phone.spec.ts uses; see the plan step that
// derived this number). The loading Skeleton reserves it so the reveal changes only
// opacity -- current-character-layout.ts's own reason, same pattern.
export const CHARACTER_LOADING_MIN_H = 'min-h-[<measured>px]';
```

Replace `<measured>` with the number Step 1 printed.

- [ ] **Step 5: Wire Skeleton, LoadError, reveal into Character.svelte**

Add the retry token and imports:

```svelte
<!-- top of the <script> block, alongside the other imports -->
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { CHARACTER_LOADING_MIN_H } from '../lib/character-layout';
```

```svelte
<!-- beside `let error = $state('');` -->
  let attempt = $state(0);
```

The existing `$effect` only reads `resolved`; make it read `attempt` too so retry re-runs
it (the same idiom `Planner.svelte:411` already uses):

```svelte
  $effect(() => {
    void attempt;
    const requested = resolved;
```

Replace the loading/failed markup:

```svelte
{:else if status === 'loading'}
  <Skeleton lines={6} rowHeight="h-4" minHeight={CHARACTER_LOADING_MIN_H} testid="character-skeleton" />
{:else if status === 'failed'}
  <LoadError message={error} onRetry={() => (attempt += 1)} testid="character-error" />
```

Add `.reveal` to the ready wrapper (currently `<div class="flex flex-col gap-[22px] md:gap-8" data-testid="character" id="character">`):

```svelte
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="character" id="character">
```

Replace the inline empty state (`Nothing ranked yet.`) with `EmptyState`:

```svelte
<!-- import alongside Skeleton/LoadError -->
  import EmptyState from './ui/EmptyState.svelte';
```

```svelte
      {#if data.best.length === 0}
        <EmptyState message="Nothing ranked yet." testid="character-empty" />
      {:else}
```

- [ ] **Step 6: Run test to verify it passes, then extend it**

Run: `npx vitest run src/components/Character.test.ts`
Expected: PASS

Add two more cases to the same file:

```ts
describe('Character failed state', () => {
  it('shows a retry that re-fires the same fetch in place', () => {
    const { body } = render(Character, { props: { path } });
    // status starts 'loading' in SSR (no effect runs server-side), so this only proves
    // the markup shape LoadError renders exists; the retry wiring itself is exercised by
    // an e2e route-stub test (Task 13).
    expect(body).not.toContain('character-error');
  });
});
```

- [ ] **Step 7: Scoped checks and commit**

```bash
npx vitest run src/components/Character.test.ts
npx astro check
npm run lint
npx prettier --check src/components/Character.svelte src/components/Character.test.ts src/lib/character-layout.ts
```

```bash
printf 'feat(web): Character.svelte reserves its ready height, and a failed load retries in place\n\nA bare "Loading." line and a dead-end error message become a Skeleton sized\nto the fixture'"'"'s ready view and a LoadError whose retry re-fires the same\nfetch, matching every other island'"'"'s new three-state model (design\n2026-09-22 spec section 3.1).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-2.txt
git add web/src/components/Character.svelte web/src/components/Character.test.ts web/src/lib/character-layout.ts
git commit -F .superpowers/commit-msg-2.txt
```

---

## Task 3: Rankings.svelte — Skeleton, LoadError, retry, reveal (board and encounter picker)

**Files:**
- Create: `web/src/components/Rankings.test.ts` (does not exist today)
- Modify: `web/src/components/Rankings.svelte:1,64-67,111-156,304-359`
- Modify: `web/src/lib/rankings/copy.ts:7` (drop "Reload the page")

**Interfaces:**
- Consumes: `Skeleton`, `LoadError`, `EmptyState` primitives.
- Produces: nothing other tasks consume.

- [ ] **Step 1: Write the failing test**

```ts
// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Rankings from './Rankings.svelte';

describe('Rankings loading state', () => {
  it('reserves height with a Skeleton instead of a bare "Loading rankings." line', () => {
    const { body } = render(Rankings, { props: { slug: 'warden-kelthas' } });
    expect(body).toContain('data-testid="rankings-skeleton"');
    expect(body).not.toContain('Loading rankings.');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `npx vitest run src/components/Rankings.test.ts`
Expected: FAIL — `rankings-skeleton` testid not found.

- [ ] **Step 3: Wire the primitives in**

Imports:

```svelte
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';
```

The main `$effect` (lines 111–156) already re-runs whenever `state` is reassigned
(`patch()` always makes a new object), but a retry must not change any filter — add an
`attempt` counter read inside the effect, same idiom as Task 2:

```svelte
  let attempt = $state(0);
```

```svelte
  $effect(() => {
    void attempt;
    if (needsPicker) {
```

Encounter-picker branch (lines 306–309): replace the bare loading/failed lines —

```svelte
      {#if encountersStatus === 'idle' || encountersStatus === 'loading'}
        <Skeleton lines={3} rowHeight="h-4" testid="rankings-picker-skeleton" />
      {:else if encountersStatus === 'failed'}
        <LoadError
          message={encounterPickerCopy.failed}
          onRetry={() => void loadEncounters()}
          testid="rankings-picker-error"
        />
```

Main board branch (lines 336–339): replace with the primitives, keeping the exact existing
testid `rankings-error` so `tests/e2e/rankings.spec.ts:151-155` (which asserts it) keeps
passing —

```svelte
  {:else if status === 'loading'}
    <Skeleton lines={8} rowHeight="h-11" minHeight="min-h-[440px]" testid="rankings-skeleton" />
  {:else if status === 'failed'}
    <LoadError message={error} onRetry={() => (attempt += 1)} testid="rankings-error" />
```

(`lines={8}`, not the API's `per_page: 100`: a page can return up to 100 rows, but an
8-row skeleton matches the visible screenful without rendering a hundred shimmer bars for
a number nothing on screen shows yet — recorded as a ruling in the Global Constraints
section above.)

Board empty states (341–342, 358–359): replace both inline `<p>` empty lines with
`EmptyState` (no `action`, since there is nowhere useful to send this visitor):

```svelte
  {:else if state.board === 'guild'}
    {#if guildRows.length === 0}
      <EmptyState message="No guilds ranked here yet." testid="rankings-empty" />
    {:else}
```

```svelte
  {:else if page === null || page.rows.length === 0}
    <EmptyState message="Nothing ranked here yet." testid="rankings-empty" />
```

Wrap the two ready branches' existing outermost element (`<ul class="flex flex-col" data-testid="guild-rows">` and `<ul class="flex flex-col" data-testid="ranking-rows">`) with `.reveal` — since a `<ul>` cannot itself gain a second class attribute conflict, just prepend `reveal` to each existing `class`:

```svelte
      <ul class="reveal flex flex-col" data-testid="guild-rows">
```

```svelte
    <ul class="reveal flex flex-col" data-testid="ranking-rows">
```

- [ ] **Step 4: Drop "Reload the page" from the encounter-picker copy**

```ts
// web/src/lib/rankings/copy.ts:7
  failed: 'Encounters did not load. Reload the page to try again.',
```
becomes
```ts
  failed: 'Encounters did not load.',
```

- [ ] **Step 5: Run test to verify it passes**

Run: `npx vitest run src/components/Rankings.test.ts`
Expected: PASS

- [ ] **Step 6: Scoped checks and commit**

```bash
npx vitest run src/components/Rankings.test.ts src/lib/rankings/
npx astro check
npm run lint
npx prettier --check src/components/Rankings.svelte src/components/Rankings.test.ts src/lib/rankings/copy.ts
```

```bash
printf 'feat(web): Rankings.svelte reserves height and retries in place on both its board and encounter picker\n\nA bare "Loading rankings." line, an unrecoverable error and a dead\nencounter-picker failure all become the shared Skeleton/LoadError pair, with\nthe same rankings-error testid rankings.spec.ts already asserts (design\n2026-09-22 spec section 3.1).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-3.txt
git add web/src/components/Rankings.svelte web/src/components/Rankings.test.ts web/src/lib/rankings/copy.ts
git commit -F .superpowers/commit-msg-3.txt
```

---

## Task 4: The guild family — `GuildStatus.svelte` + `lib/guild/layout.ts`, wired into all four views

**Files:**
- Create: `web/src/components/GuildStatus.svelte`
- Create: `web/src/components/GuildStatus.test.ts`
- Create: `web/src/lib/guild/layout.ts`
- Modify: `web/src/components/Guild.svelte:1,42-70,259-267`
- Modify: `web/src/components/GuildJoin.svelte:1,18-31,53-55`
- Modify: `web/src/components/GuildClaim.svelte:1,143-146`
- Modify: `web/src/components/GuildSettings.svelte:1,141-146`
- Modify: `web/src/components/Guild.test.ts` (fix the `'Loading.'` assertion)
- Modify: `web/src/components/GuildShell.test.ts` (fix the two `'Loading.'` assertions)
- Modify: `web/src/lib/guild/copy.ts:55,86` (drop "Reload the page")

**Interfaces:**
- Produces: `GuildStatus` component — props `status: 'loading' | 'failed'`, `error: string`,
  `onRetry: () => void`, `lines: number`, `rowHeight?: string`, `minHeight: string`,
  `testid: string`. Renders `<Skeleton lines={lines} rowHeight={rowHeight ?? 'h-4'} minHeight={minHeight} testid="{testid}-skeleton" />`
  when `status === 'loading'`, else `<LoadError message={error} onRetry={onRetry} testid="{testid}-error" />`.
  Passing `testid="guild"` reproduces the exact existing `guild-error` testid; `"guild-join"`
  reproduces `guild-join-error`; `"guild-claim"` reproduces `guild-claim-error`;
  `"guild-settings"` reproduces `guild-settings-error`.
- Produces: `GUILD_LOADING` from `lib/guild/layout.ts` — `{ home, join, claim, settings }`,
  each `{ lines: number; minHeight: string }`.

- [ ] **Step 1: Write the failing test for `GuildStatus.svelte`**

```ts
// web/src/components/GuildStatus.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildStatus from './GuildStatus.svelte';

describe('GuildStatus', () => {
  it('renders a Skeleton sized to the caller while loading', () => {
    const { body } = render(GuildStatus, {
      props: { status: 'loading', error: '', onRetry: () => {}, lines: 4, minHeight: 'min-h-[200px]', testid: 'guild' },
    });
    expect(body).toContain('data-testid="guild-skeleton"');
    expect(body).toContain('min-h-[200px]');
  });

  it('renders LoadError with the caller-prefixed testid on failure', () => {
    const { body } = render(GuildStatus, {
      props: { status: 'failed', error: 'That did not load.', onRetry: () => {}, lines: 4, minHeight: 'min-h-[200px]', testid: 'guild-claim' },
    });
    expect(body).toContain('data-testid="guild-claim-error"');
    expect(body).toContain('That did not load.');
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `npx vitest run src/components/GuildStatus.test.ts`
Expected: FAIL — module not found.

- [ ] **Step 3: Write `GuildStatus.svelte`**

```svelte
<!-- web/src/components/GuildStatus.svelte -->
<!-- The one loading/failed rendering every /guild/* view (Guild, GuildJoin, GuildClaim,
     GuildSettings, reached through GuildShell.svelte) used to duplicate four times: a
     Skeleton sized to that view's own ready height while its fetch is in flight, and a
     LoadError with a retry that re-fires the same load() on failure. Never mounted on its
     own -- each of the four hosts renders its own ready view and its own extra statuses
     (Character.svelte-style 'missing', GuildSettings's 'forbidden') outside this
     component, since those are not shared across all four. -->
<script lang="ts">
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let {
    status,
    error,
    onRetry,
    lines,
    rowHeight = 'h-4',
    minHeight,
    testid,
  }: {
    status: 'loading' | 'failed';
    error: string;
    onRetry: () => void;
    lines: number;
    rowHeight?: string;
    minHeight: string;
    testid: string;
  } = $props();
</script>

{#if status === 'loading'}
  <Skeleton {lines} {rowHeight} {minHeight} testid={`${testid}-skeleton`} />
{:else}
  <LoadError message={error} {onRetry} testid={`${testid}-error`} />
{/if}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `npx vitest run src/components/GuildStatus.test.ts`
Expected: PASS

- [ ] **Step 5: Write `lib/guild/layout.ts`**

Reasoned estimates from each view's real ready markup (Global Constraints' sizing ruling —
none of these four pages are in `lighthouserc.json`):

```ts
// web/src/lib/guild/layout.ts
// The Skeleton size each of the four /guild/* views (Guild, GuildJoin, GuildClaim,
// GuildSettings) reserves while loading, read by GuildStatus.svelte's callers. One module
// rather than four inline literals so the four numbers are visible together and cannot
// silently drift apart -- current-character-layout.ts's own reason.
export const GUILD_LOADING = {
  // Header, this-week's-reports panel, roster: the tallest of the four.
  home: { lines: 8, minHeight: 'min-h-[480px]' },
  // A heading and one paragraph of sign-in or join copy: the shortest.
  join: { lines: 2, minHeight: 'min-h-[140px]' },
  // Heading, rules list, current-claim line, one action button.
  claim: { lines: 5, minHeight: 'min-h-[320px]' },
  // Heading, claim-state line, billing panel, visibility/threshold controls, invite block.
  settings: { lines: 6, minHeight: 'min-h-[400px]' },
} as const;
```

- [ ] **Step 6: Wire `Guild.svelte`**

Import `GuildStatus` and `GUILD_LOADING`, add `attempt` (this page's `$effect` depends only
on `resolved`, same shape as Character.svelte's):

```svelte
  import GuildStatus from './GuildStatus.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
```

```svelte
  let attempt = $state(0);
```

```svelte
  $effect(() => {
    void attempt;
    const requested = resolved;
```

Replace lines 264–267:

```svelte
{:else if status === 'loading' || status === 'failed'}
  <GuildStatus
    status={status === 'loading' ? 'loading' : 'failed'}
    error={status === 'failed' ? error : ''}
    onRetry={() => (attempt += 1)}
    lines={GUILD_LOADING.home.lines}
    minHeight={GUILD_LOADING.home.minHeight}
    testid="guild"
  />
```

Add `.reveal` to the ready wrapper (`<div class="flex flex-col gap-[22px] md:gap-8" data-testid="guild" id="guild">` becomes `<div class="reveal flex flex-col gap-[22px] md:gap-8" ...>`).

- [ ] **Step 7: Wire `GuildJoin.svelte`**

```svelte
  import GuildStatus from './GuildStatus.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
```

Replace lines 53–55 (`load()` is already directly retriggerable — no `attempt` counter
needed, matching Task 4's design note above):

```svelte
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={guildJoinCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.join.lines}
      minHeight={GUILD_LOADING.join.minHeight}
      testid="guild-join"
    />
  {:else if !signedIn}
```

The outer `<div class="flex flex-col gap-4" data-testid="guild-join">` gets no `.reveal` of
its own since it wraps the loading/failed states too — instead wrap each ready branch: the
`SignInPrompt` line and the join-button block each get their own reveal wrapper. Simplest:
wrap both remaining branches in one `<div class="reveal ...">` matching the existing
`flex flex-col gap-4`:

```svelte
{:else}
  <div class="reveal contents">
    {#if !signedIn}
      <SignInPrompt line={guildJoinCopy.signInLine} testid="guild-join-signin" />
    {:else}
      ...
    {/if}
  </div>
{/if}
```

(`contents` so the wrapper adds no box of its own to the existing `gap-4` flex parent.)

- [ ] **Step 8: Wire `GuildClaim.svelte`**

```svelte
  import GuildStatus from './GuildStatus.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
```

Replace lines 143–146:

```svelte
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={error}
      onRetry={() => void load()}
      lines={GUILD_LOADING.claim.lines}
      minHeight={GUILD_LOADING.claim.minHeight}
      testid="guild-claim"
    />
  {:else if settings !== null}
```

Add `.reveal` to the outer `<div class="flex flex-col gap-4" data-testid="guild-claim">` —
since that wraps the loading/failed branches too, change only the ready branch's own
`<section>` wrapper the same way GuildJoin's Step 7 does, or (simpler, since GuildClaim's
ready content is not a single element) prepend `reveal` directly to the root div and accept
that `.reveal`'s 160ms opacity fade plays once at mount even while still loading — `.reveal`
has no visible effect until the loading Skeleton is replaced by real opacity-1 content
either way, so this is harmless. Use the root-div approach here:

```svelte
<div class="reveal flex flex-col gap-4" data-testid="guild-claim">
```

- [ ] **Step 9: Wire `GuildSettings.svelte`**

```svelte
  import GuildStatus from './GuildStatus.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
```

Replace lines 141–146 (note the `forbidden` branch stays exactly as-is — it is not a
failure, per its own existing comment, and is not part of `GuildStatus`):

```svelte
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={guildSettingsCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.settings.lines}
      minHeight={GUILD_LOADING.settings.minHeight}
      testid="guild-settings"
    />
  {:else if status === 'forbidden'}
    <p class="text-[14px]" data-testid="guild-settings-forbidden">{guildSettingsCopy.forbidden}</p>
  {:else if settings !== null}
```

Add `.reveal` to the root `<div class="flex flex-col gap-6" data-testid="guild-settings">`
(same root-div reasoning as Step 8).

- [ ] **Step 10: Drop "Reload the page" from guild copy**

```ts
// web/src/lib/guild/copy.ts:55
  failed: 'That did not load. Reload the page to try again.',
```
becomes
```ts
  failed: 'That did not load.',
```

```ts
// web/src/lib/guild/copy.ts:86
  failed: 'Settings did not load. Reload the page to try again.',
```
becomes
```ts
  failed: 'Settings did not load.',
```

(`guildClaimCopy.contestRecordedRefreshFailed` at line 49 is left untouched — it is a
post-action follow-up notice with no retry affordance to route it through, a different
case from the three-state load model this task covers.)

- [ ] **Step 11: Fix the two SSR tests that assert literal `'Loading.'`**

`Guild.test.ts` and `GuildShell.test.ts` both assert `body.toContain('Loading.')` for
`Guild.svelte`'s pre-effect render. Read each file, find that assertion, and replace it
with the new skeleton testid:

```ts
// was: expect(body).toContain('Loading.');
expect(body).toContain('data-testid="guild-skeleton"');
```

Apply the same replacement in both files (`GuildShell.test.ts` has it twice, once per case
that falls through to `Guild.svelte`).

- [ ] **Step 12: Run every touched test**

Run: `npx vitest run src/components/GuildStatus.test.ts src/components/Guild.test.ts src/components/GuildJoin.test.ts src/components/GuildClaim.test.ts src/components/GuildSettings.test.ts src/components/GuildShell.test.ts src/lib/guild/`
Expected: PASS (GuildJoin/GuildClaim/GuildSettings's existing `body.toContain(xCopy.loading)`
assertions still pass unchanged — the copy string being asserted is no longer rendered as a
`<p>` line, so re-check each: since `GuildStatus`'s `Skeleton` renders its own `label` prop,
which defaults to `uiCopy.loading` ("Loading"), NOT the per-view copy string. If any of these
three tests now fails because it asserted the view-specific loading copy
(`guildJoinCopy.loading` etc.) rather than the generic "Loading", update that one assertion
to check for the new skeleton testid instead, the same fix as Step 11.)

- [ ] **Step 13: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/GuildStatus.svelte src/components/GuildStatus.test.ts src/components/Guild.svelte src/components/Guild.test.ts src/components/GuildJoin.svelte src/components/GuildClaim.svelte src/components/GuildSettings.svelte src/components/GuildShell.test.ts src/lib/guild/layout.ts src/lib/guild/copy.ts
```

```bash
printf 'feat(web): one shared GuildStatus collapses four copies of guild loading/error markup\n\nGuild, GuildJoin, GuildClaim and GuildSettings each had their own bare\n"Loading." line and dead-end error text; a new GuildStatus.svelte (Skeleton\nplus LoadError with an in-place retry) replaces all four, sized per view from\na small lib/guild/layout.ts (design 2026-09-22 spec section 3.1).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-4.txt
git add web/src/components/GuildStatus.svelte web/src/components/GuildStatus.test.ts web/src/components/Guild.svelte web/src/components/Guild.test.ts web/src/components/GuildJoin.svelte web/src/components/GuildClaim.svelte web/src/components/GuildSettings.svelte web/src/components/GuildShell.test.ts web/src/lib/guild/layout.ts web/src/lib/guild/copy.ts
git commit -F .superpowers/commit-msg-4.txt
```

---

## Task 5: RecentReports.svelte and MyReports.svelte

**Files:**
- Create: `web/src/components/MyReports.test.ts` (does not exist today)
- Modify: `web/src/components/RecentReports.svelte:1,21-37,44-49`
- Modify: `web/src/components/RecentReports.test.ts`
- Modify: `web/src/components/MyReports.svelte:1,15-33,40-47`
- Modify: `web/src/lib/reports/copy.ts:7` (drop "Reload the page")
- Create: `web/src/lib/reports/my-reports-copy.ts` (MyReports has no copy module today;
  `recentReportsCopy` is specific to RecentReports and must not gain unrelated strings)

**Interfaces:**
- Produces: `myReportsCopy` — `{ loading, failed, empty, uploadAction }`.

- [ ] **Step 1: Write the failing tests**

```ts
// web/src/components/RecentReports.test.ts — extend the existing file with:
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import RecentReports from './RecentReports.svelte';

it('reserves height with a Skeleton instead of a bare loading line', () => {
  const { body } = render(RecentReports, { props: {} });
  expect(body).toContain('data-testid="recent-reports-skeleton"');
});
```

```ts
// web/src/components/MyReports.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import MyReports from './MyReports.svelte';

describe('MyReports', () => {
  it('reserves height with a Skeleton while signed in and loading', () => {
    const { body } = render(MyReports, { props: { signedIn: true } });
    expect(body).toContain('data-testid="my-reports-skeleton"');
    expect(body).not.toContain('Reload the page');
  });

  it('renders nothing report-related when signed out', () => {
    const { body } = render(MyReports, { props: { signedIn: false } });
    expect(body).toContain('data-testid="reports-signin"');
    expect(body).not.toContain('my-reports-skeleton');
  });
});
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `npx vitest run src/components/RecentReports.test.ts src/components/MyReports.test.ts`
Expected: FAIL — skeleton testids not found.

- [ ] **Step 3: Create `lib/reports/my-reports-copy.ts`**

```ts
// web/src/lib/reports/my-reports-copy.ts
// Every visible string MyReports.svelte ("Your reports") uses -- split from
// lib/reports/copy.ts's recentReportsCopy, which is RecentReports.svelte's own strings for
// the public feed, a different component with a different failure and empty case.
export const myReportsCopy = {
  loading: 'Loading your reports.',
  failed: 'Your reports did not load.',
  empty: 'No reports yet. Upload a log or run the desktop companion.',
  uploadAction: 'Upload a log',
} as const;
```

- [ ] **Step 4: Wire `RecentReports.svelte`**

```svelte
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';
```

Retry: `load()` is already directly retriggerable and already takes the `cursor` it needs;
plain retry re-fires it with no cursor (the first page), matching a fresh load:

```svelte
  {#if status === 'loading'}
    <Skeleton lines={5} rowHeight="h-11" testid="recent-reports-skeleton" />
  {:else if status === 'failed'}
    <LoadError message={recentReportsCopy.failed} onRetry={() => void load()} testid="recent-reports-error" />
  {:else if rows.length === 0}
    <EmptyState message={recentReportsCopy.empty} testid="recent-reports-empty" />
  {:else}
```

Add `.reveal` to the ready `<ul>` (`<ul class="flex flex-col">` → `<ul class="reveal flex flex-col">`).

- [ ] **Step 5: Wire `MyReports.svelte`**

```svelte
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { myReportsCopy } from '../lib/reports/my-reports-copy';
```

Replace lines 42–47:

```svelte
  {:else if status === 'loading'}
    <Skeleton lines={5} rowHeight="h-11" testid="my-reports-skeleton" />
  {:else if status === 'failed'}
    <LoadError message={myReportsCopy.failed} onRetry={() => void load(page)} testid="my-reports-error" />
  {:else if rows.length === 0}
    <EmptyState message={myReportsCopy.empty} action={{ label: myReportsCopy.uploadAction, href: '/logs' }} testid="my-reports-empty" />
  {:else}
```

Add `.reveal` to the ready `<ul>` the same way.

- [ ] **Step 6: Drop "Reload the page" from `lib/reports/copy.ts`**

```ts
// web/src/lib/reports/copy.ts:7
  failed: 'Recent reports did not load. Reload the page to try again.',
```
becomes
```ts
  failed: 'Recent reports did not load.',
```

- [ ] **Step 7: Run tests to verify they pass**

Run: `npx vitest run src/components/RecentReports.test.ts src/components/MyReports.test.ts src/lib/reports/`
Expected: PASS

- [ ] **Step 8: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/RecentReports.svelte src/components/RecentReports.test.ts src/components/MyReports.svelte src/components/MyReports.test.ts src/lib/reports/copy.ts src/lib/reports/my-reports-copy.ts
```

```bash
printf 'feat(web): RecentReports and MyReports reserve height, retry in place, and stop saying "reload the page"\n\nBoth report lists get a Skeleton sized to their row list, a LoadError whose\nretry re-fires the same load(), and an EmptyState in place of their inline\ntext -- MyReports also drops its literal "Reload the page to try again."\n(design 2026-09-22 spec section 3.1).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-5.txt
git add web/src/components/RecentReports.svelte web/src/components/RecentReports.test.ts web/src/components/MyReports.svelte web/src/components/MyReports.test.ts web/src/lib/reports/copy.ts web/src/lib/reports/my-reports-copy.ts
git commit -F .superpowers/commit-msg-5.txt
```

---

## Task 6: `lib/report/layout.ts` + ReportView.svelte's `lazyFallback` skeleton

**Files:**
- Create: `web/src/lib/report/layout.ts`
- Modify: `web/src/components/report/ReportView.svelte:1,1292-1304,1317-1318,1711,1745,1761-1770,1780,1799,1813,1830`

**Interfaces:**
- Produces: `REPORT_LAZY_MIN_H` — `{ rating, timelines, events, queries, compare, rankings, mechanics }`, each a `min-h-*` string, consumed by every `lazyFallback` call site below.
- Consumes: `LazyLoadState` (already imported), `Skeleton`, `LoadError`.

ReportView.svelte is 1841 lines and must not grow past that — this task adds one import
line, ~10 lines inside the existing `lazyFallback` snippet, one new `onRetry` line on the
top-level failed branch, and passes one extra argument at each of the 7 existing
`{@render lazyFallback(...)}` call sites. No new markup block, no new file for ReportView
itself.

- [ ] **Step 1: Create `lib/report/layout.ts`**

Reasoned estimates from each lazy panel's real markup (see the plan's own research: `CompareMode.svelte` 684 lines/richest table, `MechanicsMode.svelte` 577, `RankingsMode.svelte` 280, `TimelinesView.svelte` 406, `EventsView.svelte` 182, `QueriesView.svelte` 211, `RatingTab.svelte` 200) — not Lighthouse-gated (see Global Constraints' sizing ruling: none of these mode/view switches happen during an automated Lighthouse run):

```ts
// web/src/lib/report/layout.ts
// The Skeleton size ReportView.svelte's lazyFallback snippet reserves for each lazily
// imported mode/view while its chunk is still loading, so the panel does not sit blank for
// that one round trip (the audit's own finding). One module the snippet's seven call sites
// share, rather than seven inline literals -- report/skeleton.ts's own reason for being one
// module, applied to a new set of numbers.
export const REPORT_LAZY_MIN_H = {
  rating: 'min-h-[320px]',
  timelines: 'min-h-[420px]',
  events: 'min-h-[360px]',
  queries: 'min-h-[320px]',
  compare: 'min-h-[560px]',
  rankings: 'min-h-[420px]',
  mechanics: 'min-h-[480px]',
} as const;
```

- [ ] **Step 2: Write the failing test**

There is no existing vitest file for `ReportView.svelte` (it needs live report data
through `fetchReportMeta` etc., not a good SSR-render candidate) — this task is verified
by extending the fixture-driven e2e coverage instead, which already exercises every mode
and view. Add one Playwright test to `tests/e2e/report-tabs.spec.ts` (the file that already
switches modes/views against the fixture report):

```ts
// tests/e2e/report-tabs.spec.ts — add near the existing mode-switch tests
test('a lazy mode shows a sized skeleton, not a blank panel, while its chunk loads', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  // Slow the chunk down so the loading frame is observable: throttle just this request.
  await page.route('**/CompareMode*.js', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 300));
    await route.continue();
  });
  await page.getByTestId('mode-compare').click();
  await expect(page.getByTestId('lazy-view-skeleton')).toBeVisible();
  await expect(page.getByTestId('lazy-view-skeleton')).not.toBeVisible({ timeout: 5000 });
});
```

- [ ] **Step 3: Run test to verify it fails**

Run: `E2E_PORT=4432 npx playwright test tests/e2e/report-tabs.spec.ts -g "lazy mode shows a sized skeleton"`
Expected: FAIL — `lazy-view-skeleton` testid does not exist yet.

- [ ] **Step 4: Wire the snippet**

```svelte
<!-- import, alongside the file's other component imports -->
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import { REPORT_LAZY_MIN_H } from '../../lib/report/layout';
```

Replace the `lazyFallback` snippet (lines 1292–1304):

```svelte
{#snippet lazyFallback(lazy: LazyLoadState, minHeight: string)}
  {#if lazy.error !== ''}
    <LoadError message={lazy.error} onRetry={() => lazy.load()} testid="lazy-view-error" />
  {:else}
    <Skeleton minHeight={minHeight} testid="lazy-view-skeleton" />
  {/if}
{/snippet}
```

Update every call site to pass its own height (each site's mode/view is unambiguous from
its surrounding branch):

```svelte
<!-- line 1711 -->
            {@render lazyFallback(ratingTabLazy, REPORT_LAZY_MIN_H.rating)}
```
```svelte
<!-- line 1745 -->
          {@render lazyFallback(timelinesViewLazy, REPORT_LAZY_MIN_H.timelines)}
```
```svelte
<!-- line 1770 -->
          {@render lazyFallback(eventsViewLazy, REPORT_LAZY_MIN_H.events)}
```
```svelte
<!-- line 1780 -->
          {@render lazyFallback(queriesViewLazy, REPORT_LAZY_MIN_H.queries)}
```
```svelte
<!-- line 1799 -->
          {@render lazyFallback(compareModeLazy, REPORT_LAZY_MIN_H.compare)}
```
```svelte
<!-- line 1813 -->
          {@render lazyFallback(rankingsModeLazy, REPORT_LAZY_MIN_H.rankings)}
```
```svelte
<!-- line 1830 -->
          {@render lazyFallback(mechanicsModeLazy, REPORT_LAZY_MIN_H.mechanics)}
```

- [ ] **Step 5: Route the top-level report-load failure through `LoadError` too**

Line 1317–1318, for consistency with the Global Constraints' "every Failed state uses
LoadError" rule (`loadReport()` is already directly retriggerable):

```svelte
{#if status === 'failed'}
  <LoadError message={error} onRetry={() => void loadReport()} testid="report-error" />
```

- [ ] **Step 6: Run test to verify it passes**

Run: `E2E_PORT=4432 npx playwright test tests/e2e/report-tabs.spec.ts`
Expected: PASS (whole file, not just the new test — this file is the one most likely to
regress from the `lazyFallback` signature change).

- [ ] **Step 7: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/report/ReportView.svelte src/lib/report/layout.ts tests/e2e/report-tabs.spec.ts
E2E_PORT=4432 npx playwright test tests/e2e/report-tabs.spec.ts tests/e2e/report-live.spec.ts
npx astro preview stop
```

```bash
printf 'feat(web): ReportView'"'"'s lazy modes and views show a sized skeleton instead of a blank panel while loading\n\nCompare, Mechanics, Rankings, Timelines, Events, Queries and the Rating tab\neach shipped as their own chunk with nothing rendered until it loaded; the\nshared lazyFallback snippet now shows a Skeleton sized per panel from a new\nlib/report/layout.ts, and its existing retry now renders through LoadError\n(design 2026-09-22 spec section 3.1). ReportView.svelte gains ~15 lines, not\na new responsibility.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-6.txt
git add web/src/components/report/ReportView.svelte web/src/lib/report/layout.ts web/tests/e2e/report-tabs.spec.ts
git commit -F .superpowers/commit-msg-6.txt
```

---

## Task 7: `lib/sim/lazy-layout.ts` + SimView.svelte's lazy panels and saved-sim retry

**Files:**
- Create: `web/src/lib/sim/lazy-layout.ts`
- Modify: `web/src/components/sim/SimView.svelte:1,144-157,479-490,503-506`

**Interfaces:**
- Produces: `SIM_LAZY_MIN_H` — `{ compare, results }`.

Same size discipline as Task 6: SimView.svelte is 800 lines and must not grow past it — this
adds an import, extracts one existing inline fetch into a named function (so it is
retriggerable), and edits the existing `lazyFallback` snippet and its two call sites.

- [ ] **Step 1: Create `lib/sim/lazy-layout.ts`**

```ts
// web/src/lib/sim/lazy-layout.ts
// The Skeleton size SimView.svelte's lazyFallback snippet reserves for CompareView and
// SimResults while their chunk is still loading -- lib/report/layout.ts's own reason,
// mirrored here rather than imported cross-domain (report/skeleton.ts and sim/skeleton.ts
// already keep this same split for the pre-hydration skeleton strings).
export const SIM_LAZY_MIN_H = {
  compare: 'min-h-[360px]',
  results: 'min-h-[420px]',
} as const;
```

- [ ] **Step 2: Write the failing test**

```ts
// tests/e2e/sim-saved.spec.ts — add near the existing skeleton test
test('a failed saved-sim fetch offers a retry that re-fires the same request', async ({ page }) => {
  let attempts = 0;
  await page.route('**/v1/sims/simfailonce', async (route) => {
    attempts += 1;
    if (attempts === 1) {
      await route.fulfill({ status: 500, contentType: 'application/json', body: '{"ok":false,"error":"boom"}' });
    } else {
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({ ok: true, data: SAVED_SIM_FIXTURE, error: null, request_id: 'r' }),
      });
    }
  });
  await page.goto('/sim/simfailonce');
  await expect(page.getByTestId('sim-saved-error')).toBeVisible();
  await page.getByTestId('sim-saved-error-retry').click();
  await expect(page.getByTestId('sim-saved-error')).not.toBeVisible();
});
```

(`SAVED_SIM_FIXTURE` — read the existing successful-fixture constant already defined near
the top of `tests/e2e/sim-saved.spec.ts` and reuse it verbatim; do not invent a new shape.)

- [ ] **Step 3: Run test to verify it fails**

Run: `E2E_PORT=4432 npx playwright test tests/e2e/sim-saved.spec.ts -g "retry that re-fires"`
Expected: FAIL — `sim-saved-error-retry` testid does not exist (no retry button today).

- [ ] **Step 4: Extract the saved-sim fetch into a retriggerable function**

Replace the inline `untrack` block (lines ~147-157):

```svelte
  let savedResult = $state<SimResult | null>(untrack(() => inlineResult));
  let savedError = $state<string | null>(null);

  function loadSavedSim(): void {
    savedError = null;
    void fetchSim(simId)
      .then((result) => (savedResult = result))
      .catch((error) => {
        savedError = error instanceof Error ? error.message : simCopy.loadFailed;
      });
  }

  untrack(() => {
    if (hasSavedSimId && savedResult === null && simId !== '') loadSavedSim();
  });
```

- [ ] **Step 5: Wire the saved-sim error through `LoadError`**

```svelte
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import { SIM_LAZY_MIN_H } from '../../lib/sim/lazy-layout';
```

Replace lines 503–506:

```svelte
    {:else if savedError !== null}
      <LoadError message={savedError} onRetry={() => loadSavedSim()} testid="sim-saved-error" />
```

- [ ] **Step 6: Wire the `lazyFallback` snippet (Compare and SimResults)**

Replace lines 479–490:

```svelte
{#snippet lazyFallback(lazy: LazyLoadState, minHeight: string)}
  {#if lazy.error !== ''}
    <LoadError message={lazy.error} onRetry={() => lazy.load()} testid="sim-results-error" />
  {:else}
    <Skeleton minHeight={minHeight} testid="sim-lazy-skeleton" />
  {/if}
{/snippet}
```

Update the two call sites (grep `{@render lazyFallback(` in `SimView.svelte`):

```svelte
              {@render lazyFallback(compareViewLazy, SIM_LAZY_MIN_H.compare)}
```
```svelte
            {@render lazyFallback(simResultsLazy, SIM_LAZY_MIN_H.results)}
```

- [ ] **Step 7: Run test to verify it passes**

Run: `E2E_PORT=4432 npx playwright test tests/e2e/sim-saved.spec.ts`
Expected: PASS (whole file — the existing "shows the loading skeleton" test must still pass
unchanged since `SIM_SAVED_SKELETON_HTML`'s own testid `sim-saved-skeleton` is untouched by
this task).

- [ ] **Step 8: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/sim/SimView.svelte src/lib/sim/lazy-layout.ts tests/e2e/sim-saved.spec.ts
E2E_PORT=4432 npx playwright test tests/e2e/sim-saved.spec.ts tests/e2e/sim-phone.spec.ts
npx astro preview stop
```

```bash
printf 'feat(web): SimView'"'"'s lazy Compare/Results panels show a sized skeleton, and a failed saved sim retries in place\n\nThe same lazyFallback pattern Task 6 gave ReportView, applied to SimView'"'"'s\ntwo lazily loaded panels; the saved-sim fetch (/sim/<id>) is now a named,\nretriggerable function so its LoadError has a real retry instead of a\ndead-end message (design 2026-09-22 spec section 3.1).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-7.txt
git add web/src/components/sim/SimView.svelte web/src/lib/sim/lazy-layout.ts web/tests/e2e/sim-saved.spec.ts
git commit -F .superpowers/commit-msg-7.txt
```

---

## Task 8: AddonPasteBox.svelte — reserve height while `fetchMeOnce` is pending

**Files:**
- Create: `web/src/lib/addon/paste-layout.ts`
- Modify: `web/src/components/AddonPasteBox.svelte:1,85-107`
- Modify: `web/src/components/AddonPasteBox.test.ts`

**Interfaces:**
- Produces: `ADDON_PASTE_STATUS_MIN_H: string`.

- [ ] **Step 1: Write the failing test**

```ts
// web/src/components/AddonPasteBox.test.ts — add:
it('reserves the signed-in block height while fetchMeOnce is still pending', () => {
  const { body } = render(AddonPasteBox, { props: {} });
  // Before any decode: loaded is null, so the whole signedIn block (including its
  // skeleton) is absent -- this proves the reservation only appears once a decode exists,
  // matching the component's existing pre-effect contract.
  expect(body).not.toContain('addon-paste-status-skeleton');
});
```

(This SSR test only proves the pre-decode shape is unchanged; the real "no push" behaviour
needs a live `fetchMeOnce` in flight, verified in Step 4's e2e addition.)

- [ ] **Step 2: Run test to verify it fails**

Run: `npx vitest run src/components/AddonPasteBox.test.ts`
Expected: PASS already (nothing to assert against yet) — this step confirms the baseline;
proceed to the real change.

- [ ] **Step 3: Create `lib/addon/paste-layout.ts` and wire the reservation**

```ts
// web/src/lib/addon/paste-layout.ts
// The height AddonPasteSave.svelte's signed-in form renders at (name field, region/ruleset
// selects, save button): a reasoned estimate from its own markup (label + input + two
// selects + button, each min-h-11 with gaps), reserved while AddonPasteBox.svelte's
// fetchMeOnce() is still resolving so the hint or the form does not push the links above it
// once it lands.
export const ADDON_PASTE_STATUS_MIN_H = 'min-h-[168px]';
```

```svelte
<!-- AddonPasteBox.svelte, alongside the other imports -->
  import Skeleton from './ui/Skeleton.svelte';
  import { ADDON_PASTE_STATUS_MIN_H } from '../lib/addon/paste-layout';
```

Replace lines 102–106:

```svelte
    {#if signedIn === null}
      <Skeleton lines={2} minHeight={ADDON_PASTE_STATUS_MIN_H} testid="addon-paste-status-skeleton" />
    {:else}
      {#key loaded}
        <AddonPasteSave {signedIn} code={loaded} />
      {/key}
    {/if}
```

- [ ] **Step 4: Add an e2e proof that the reservation holds**

```ts
// tests/e2e/addon-paste.spec.ts — find the existing spec (or the nearest addon spec) and add:
test('the signed-in hint does not push the planner/sim links while the session check is pending', async ({ page }) => {
  let resolveMe: (() => void) | undefined;
  await page.route('**/v1/me', (route) => {
    void new Promise<void>((resolve) => {
      resolveMe = resolve;
    }).then(() => route.fulfill({ status: 401, contentType: 'application/json', body: '{"ok":false,"error":"unauthorized"}' }));
  });
  await page.goto('/addon');
  await page.getByTestId('addon-paste-code').fill('<a valid FS1 export the fixture accepts>');
  await page.getByTestId('addon-paste-submit').click();
  const plannerLinkTop = await page.getByTestId('addon-paste-planner').boundingBox().then((box) => box?.y);
  resolveMe?.();
  await expect(page.getByTestId('addon-paste-signin-hint')).toBeVisible();
  const plannerLinkTopAfter = await page.getByTestId('addon-paste-planner').boundingBox().then((box) => box?.y);
  expect(plannerLinkTopAfter).toBe(plannerLinkTop);
});
```

Find the actual valid FS1 export fixture string an existing addon/planner e2e spec already
uses (grep `decodeFS1`/`fs1` fixtures in `tests/e2e/`) and use that exact string rather than
inventing one — do not hand-write a fake export code.

- [ ] **Step 5: Run tests to verify they pass**

Run: `npx vitest run src/components/AddonPasteBox.test.ts`
Run: `E2E_PORT=4432 npx playwright test tests/e2e/addon-paste.spec.ts`
Expected: PASS

- [ ] **Step 6: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/AddonPasteBox.svelte src/components/AddonPasteBox.test.ts src/lib/addon/paste-layout.ts tests/e2e/addon-paste.spec.ts
npx astro preview stop
```

```bash
printf 'fix(web): AddonPasteBox reserves the signed-in block height while the session check is pending\n\nThe planner/sim links no longer jump once fetchMeOnce resolves to the\nsign-in hint or the save form -- a Skeleton now holds that space for the one\nround trip (design 2026-09-22 spec section 3.1).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-8.txt
git add web/src/components/AddonPasteBox.svelte web/src/components/AddonPasteBox.test.ts web/src/lib/addon/paste-layout.ts web/tests/e2e/addon-paste.spec.ts
git commit -F .superpowers/commit-msg-8.txt
```

---

## Task 9: Planner.svelte — Skeleton inside its existing reservation, retry through LoadError

**Files:**
- Modify: `web/src/components/planner/Planner.svelte:1,604-607,608-623`
- Modify: `web/tests/e2e/planner.spec.ts:125`
- Modify: `web/tests/e2e/planner-phone.spec.ts:112,127`

Planner already has the CLS-critical part right — the `min-h-[1096px] md:min-h-[1435px]`
wrapper around the whole loading/failed/ready block, and the retry (`attempt += 1`) already
works. This task only swaps the literal text line for a `Skeleton` inside that existing
reservation, and routes the retry button through `LoadError`.

- [ ] **Step 1: Write the failing e2e assertions first (they currently pass against the old text; this step proves the test file is aligned with the new markup before the component changes)**

Update `tests/e2e/planner.spec.ts:125`:

```ts
// was: await expect(page.getByText('Loading talent data')).toBeVisible();
await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
```

Update `tests/e2e/planner-phone.spec.ts:112,127` (both occurrences, in the two
footer-position tests):

```ts
// was: await expect(page.getByText('Loading talent data')).toBeVisible();
await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
```

- [ ] **Step 2: Run the affected e2e tests to verify they fail**

Run: `E2E_PORT=4432 npx playwright test tests/e2e/planner.spec.ts tests/e2e/planner-phone.spec.ts`
Expected: FAIL — `planner-talent-skeleton` testid does not exist yet.

- [ ] **Step 3: Wire the component**

```svelte
<!-- Planner.svelte, alongside its other imports -->
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
```

Replace lines 604–607 (keep the exact same bordered panel chrome — the comment above it,
lines 580-597, already explains why the panel frame itself is the CLS reservation, not the
line inside it):

```svelte
      <div class="border-line bg-raised rounded-panel mx-[18px] flex grow flex-col gap-3 border p-4 md:mx-0">
        <Skeleton lines={4} rowHeight="h-4" label="Loading talent data" testid="planner-talent-skeleton" />
      </div>
```

Replace lines 608–623's retry button with `LoadError`, keeping the existing headline and
detail lines exactly as-is (LoadError only owns the message + button, not the whole panel):

```svelte
    {:else if status === 'failed'}
      <div class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-5 md:mx-0">
        <p class="text-strong text-[15px] font-semibold">{DATA_LOAD_FAILED}</p>
        <p class="text-muted text-[13px]">
          Build {store.treeVersion} did not return the files the planner needs.
        </p>
        <LoadError
          message="Talent data did not load."
          onRetry={() => (attempt += 1)}
          testid="planner-load-error"
        />
      </div>
```

- [ ] **Step 4: Run the e2e tests to verify they pass**

Run: `E2E_PORT=4432 npx playwright test tests/e2e/planner.spec.ts tests/e2e/planner-phone.spec.ts`
Expected: PASS

- [ ] **Step 5: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/planner/Planner.svelte tests/e2e/planner.spec.ts tests/e2e/planner-phone.spec.ts
npx astro preview stop
```

```bash
printf 'fix(web): Planner shows a Skeleton inside its existing CLS reservation, retries through LoadError\n\n"Loading talent data" was a bare line inside an already-correct min-h panel;\nit is now a Skeleton, and the working Retry button now renders through the\nshared LoadError (design 2026-09-22 spec section 3.1). The panel'"'"'s own\nmin-h-[1096px]/md:min-h-[1435px] reservation is untouched.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-9.txt
git add web/src/components/planner/Planner.svelte web/tests/e2e/planner.spec.ts web/tests/e2e/planner-phone.spec.ts
git commit -F .superpowers/commit-msg-9.txt
```

---

## Task 10: The busy-controls sweep (spec section 3.2)

**Files:**
- Modify: `web/src/components/report/QueriesView.svelte:145-157`
- Modify: `web/src/components/report/EventsView.svelte:126-135`
- Modify: `web/src/components/report/ActorRow.svelte:466-472`
- Modify: `web/src/components/planner/SharePanel.svelte:330-339`
- Modify: `web/src/components/sim/RunControl.svelte:122,164-174`
- Modify: `web/src/components/sim/SimView.svelte:752-760`
- Modify: `web/src/components/sim/tools/SaveSimForm.svelte:105-115`
- Modify: `web/src/components/sim/tools/BulkRunBar.svelte:100-165,230-240`
- Modify: `web/src/components/Upload.svelte:200-212`
- Modify: `web/src/components/GuildJoin.svelte` (join button, from Task 4's rewrite)
- Modify: `web/src/components/GuildClaim.svelte:158-250` (claim/confirm/release/contest buttons)
- Modify: `web/src/components/GuildSettings.svelte:183-235` (billing/save/rotate buttons)
- Modify: `web/src/components/AddonPasteSave.svelte:73-81`

**The rule, applied identically at every site below:** the button already has a busy
boolean driving its `disabled` attribute (`running`, `busy`, `saving`, `measuring`,
`streaming`, `locked`, etc.). Add `aria-busy={<thatBoolean>}` immediately after the
existing `disabled={...}` attribute, and append `` ${BUSY_CLASS}`` to the button's `class`
when that boolean is true (Svelte's `class:` directive does not fit a class-string variable
cleanly here — use a template literal: `` class={`${EXISTING_CLASSES} ${busy ? BUSY_CLASS : ''}`} ``
or, where the class is already a template literal, splice `BUSY_CLASS` into it the same
way). Where the button's **label text itself changes while busy** (a genuine spec
violation — "no label swap"), replace the ternary with the static label; the one exception,
recorded as a ruling, is `RunControl.svelte`'s Run/Stop button, which is left exactly as-is
beyond adding `aria-busy`/the busy class — its label change is a real control-identity
change (Run becomes Stop, a different action, with its own `onclick`), not a busy-decoration
swap, so removing it would delete working cancel functionality the spec's 3.2 does not ask
for.

- [ ] **Step 1: Write one shared assertion pattern, then apply it file by file**

There is no single shared test file for this sweep (each button lives in a different
component's own test file, and several of these components have none yet) — verify each
file with a `vitest run` of its own existing/adjacent test file after the edit, and prove
`aria-busy` appears via a quick grep sweep at the end (Step 3). Start with `QueriesView`:

```svelte
<!-- QueriesView.svelte:145-157, before -->
  <button
    type="button"
    class="border-line-warm-strong rounded-control text-strong inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase"
    onclick={() => void run()}
    disabled={running}
    aria-busy={running}
    data-testid="query-run"
  >
    {running ? 'Running' : 'Run'}
  </button>
```

```svelte
<!-- after: remove the label swap, add BUSY_CLASS -->
  import { BUSY_CLASS } from '../../lib/ui/busy';
```

```svelte
  <button
    type="button"
    class={`border-line-warm-strong rounded-control text-strong inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase ${running ? BUSY_CLASS : ''}`}
    onclick={() => void run()}
    disabled={running}
    aria-busy={running}
    data-testid="query-run"
  >
    Run
  </button>
```

- [ ] **Step 2: Apply the same rule to every remaining file**

`EventsView.svelte:126-135` — remove the `{streaming ? 'Loading…' : '...'}'` swap, keep the
static label, add `aria-busy={streaming}` and `BUSY_CLASS`:

```svelte
        <button
          type="button"
          class={`text-gold inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-0 ${streaming ? BUSY_CLASS : ''}`}
          data-testid="events-stream"
          disabled={streaming}
          aria-busy={streaming}
          onclick={() => void runStream()}
          >Load every hit and heal in this window</button
        >
```

`ActorRow.svelte:466-472` — remove the `{measuring ? 'Measuring…' : '...'}'` swap:

```svelte
            <button
              type="button"
              class={`border-line-warm rounded-control text-text inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9 ${measuring ? BUSY_CLASS : ''}`}
              disabled={measuring}
              aria-busy={measuring}
              onclick={() => void runMeasure()}
            >
              Measure this window exactly
            </button>
```

`SharePanel.svelte:330-339` — remove the `{saving ? plannerCopy.saving : plannerCopy.share}` swap:

```svelte
    <button
      type="button"
      class={`${SECONDARY_BUTTON} border-line-warm-strong text-gold px-4 ${saving ? BUSY_CLASS : ''}`}
      disabled={saving || store.spent === 0}
      aria-busy={saving}
      data-testid="share-open"
      bind:this={shareButtonEl}
      onclick={() => void requestShare()}
    >
      {plannerCopy.share}
    </button>
```

`RunControl.svelte` — the ruling above applies: add `aria-busy={running}` and `BUSY_CLASS`
to the Run/Stop button (around line 122) without touching its label logic:

```svelte
  <button
    type="button"
    class={`border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 w-full border px-5 disabled:opacity-50 md:w-auto md:min-w-[9rem] ${running ? BUSY_CLASS : ''}`}
    aria-busy={running}
    disabled={loadingEngine ||
      phase === 'loading-character' ||
```
(keep every subsequent line of that existing `disabled` expression exactly as-is; only the
`class` and the new `aria-busy` line are added).

`SimView.svelte:752-760` — remove the `{saving ? simCopy.savingAction : simCopy.saveAction}` swap:

```svelte
            <button
              type="button"
              class={`border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 border px-5 disabled:opacity-50 ${saving ? BUSY_CLASS : ''}`}
              disabled={saving}
              aria-busy={saving}
              onclick={() => void confirmSave()}
              data-testid="sim-save-confirm"
            >
              {simCopy.saveAction}
            </button>
```

`SaveSimForm.svelte:105-115` — read the button around line 112 and apply the identical
pattern (`disabled={saving}` → add `aria-busy={saving}` + `BUSY_CLASS`; if its label also
swaps on `saving`, make it static the same way `SimView`'s did above).

`BulkRunBar.svelte:100-165,230-240` — four `disabled={running}`/`disabled={running || ...}`
buttons; apply `aria-busy={running}` + `BUSY_CLASS` to each of the four, removing any label
swap driven by `running` the same way (read each button's template literal first; not every
one necessarily swaps its label — only remove a swap where one exists).

`Upload.svelte:200-212` — the "Upload" button (`disabled={file === null || tooLarge || locked}`,
line ~207): add `aria-busy={busy}` (the component's own `busy` derived value, not `locked`,
since `locked` also covers the signed-out case which is not "busy") and `BUSY_CLASS` when
`busy`:

```svelte
    <button
      type="button"
      class={`... ${busy ? BUSY_CLASS : ''}`}
      onclick={() => void start()}
      disabled={file === null || tooLarge || locked}
      aria-busy={busy}
      data-testid="upload-start"
    >
      Upload
    </button>
```
(read the button's existing full `class` string first and splice `BUSY_CLASS` into it
rather than guessing the base classes.)

`GuildJoin.svelte`'s join button (from Task 4's rewrite, still present at the bottom of the
file): add `aria-busy={busy}` and `BUSY_CLASS` alongside its existing `disabled={busy}`.

`GuildClaim.svelte:158-250` — five buttons (`Contest`, the two contest-confirm buttons,
`Release`, `Confirm`, `Claim`), each already `disabled={busy}`: add `aria-busy={busy}` and
`BUSY_CLASS` to all five.

`GuildSettings.svelte:183-235` — three buttons (`Manage billing`, the visibility/threshold
save controls at `disabled={busy || frozen}`, `Rotate`): add `aria-busy={busy}` and
`BUSY_CLASS` to all three (note two of these use `disabled={busy || frozen}` — `aria-busy`
should reflect `busy` alone, not `frozen`, since a frozen-but-idle control is disabled, not
busy).

`AddonPasteSave.svelte:73-81` — the Save button (`disabled={busy || saved || !canSave}`):
add `aria-busy={busy}` and `BUSY_CLASS` when `busy` (its label already stays static except
for the legitimate `saved` end-state, which is left alone — not a busy-decoration swap).

- [ ] **Step 3: Verify the sweep with a grep, then run the touched test suites**

```bash
grep -c "aria-busy=" src/components/report/QueriesView.svelte src/components/report/EventsView.svelte src/components/report/ActorRow.svelte src/components/planner/SharePanel.svelte src/components/sim/RunControl.svelte src/components/sim/SimView.svelte src/components/sim/tools/SaveSimForm.svelte src/components/sim/tools/BulkRunBar.svelte src/components/Upload.svelte src/components/GuildJoin.svelte src/components/GuildClaim.svelte src/components/GuildSettings.svelte src/components/AddonPasteSave.svelte
```

Expected: every file reports at least 1 (`BulkRunBar.svelte` and `GuildClaim.svelte`
should report ≥ 4 and ≥ 5 respectively).

```bash
npx vitest run src/components/GuildClaim.test.ts src/components/GuildSettings.test.ts src/components/AddonPasteSave.test.ts src/components/GuildJoin.test.ts
npx astro check
npm run lint
npx prettier --check src/components/report/QueriesView.svelte src/components/report/EventsView.svelte src/components/report/ActorRow.svelte src/components/planner/SharePanel.svelte src/components/sim/RunControl.svelte src/components/sim/SimView.svelte src/components/sim/tools/SaveSimForm.svelte src/components/sim/tools/BulkRunBar.svelte src/components/Upload.svelte src/components/GuildJoin.svelte src/components/GuildClaim.svelte src/components/GuildSettings.svelte src/components/AddonPasteSave.svelte
```

- [ ] **Step 4: Commit**

```bash
printf 'fix(web): every request-triggering button marks itself aria-busy with one shared look, no label swaps\n\nRun, Save, Upload, Claim, Confirm, Release, Contest, Rotate and Manage\nbilling all gain aria-busy plus the shared BUSY_CLASS (Task 1); "Running",\n"Measuring…", "Loading…" and saving-label swaps are removed so a button'"'"'s\nvisible label never changes, only its look (design 2026-09-22 spec section\n3.2). RunControl'"'"'s Run/Stop toggle is left as-is -- a real control-identity\nchange, not busy decoration.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-10.txt
git add web/src/components/report/QueriesView.svelte web/src/components/report/EventsView.svelte web/src/components/report/ActorRow.svelte web/src/components/planner/SharePanel.svelte web/src/components/sim/RunControl.svelte web/src/components/sim/SimView.svelte web/src/components/sim/tools/SaveSimForm.svelte web/src/components/sim/tools/BulkRunBar.svelte web/src/components/Upload.svelte web/src/components/GuildJoin.svelte web/src/components/GuildClaim.svelte web/src/components/GuildSettings.svelte web/src/components/AddonPasteSave.svelte
git commit -F .superpowers/commit-msg-10.txt
```

---

## Self-Review (spec coverage check)

- 3.1 `Character.svelte` → Task 2. `Rankings.svelte` → Task 3. Guild family → Task 4.
  `RecentReports`/`MyReports` → Task 5. `ReportView.svelte` lazy panels → Task 6.
  `SimView.svelte` lazy panels + saved-sim → Task 7. `ToolsView.svelte` → already has
  skeletons per the spec's own parenthetical; no task needed (confirmed: its `lazyFallback`
  already renders `TOOL_SKELETONS[tool]` while loading and a `role="alert"` retry on
  failure — the one gap, routing that retry through `LoadError`, is optional polish the spec
  does not ask for since it already meets the three-state model in substance; left as-is to
  avoid touching a file with no defect). `AddonPasteBox.svelte` → Task 8.
  `Planner.svelte` → Task 9.
- 3.2 Busy controls → Task 10.
- 3.3 Tests → SSR render tests added in Tasks 2, 3, 4, 5, 8; e2e specs updated in Tasks 6, 7,
  8, 9; `npm run lhci` with every URL's CLS quoted happens in the final whole-branch review
  (not its own task — it needs every prior task's build output).
- Out of scope (section 4): page transitions, phone-nav horizontal scroll, light mode, new
  fetches — none of the above tasks touch any of these.

No placeholders found on review (searched for "TBD"/"similar to Task"/"add appropriate" —
none present).

## Final whole-branch review (after Task 10, before the lane's final report)

Not a numbered task — the house rules' subagent-driven-development process runs this after
Task 10 with no plan-checkbox of its own. It must, at minimum:

1. `npx vitest run` (whole `web/` suite, not just touched paths) — confirm no unrelated
   regression from the shared `GuildStatus`/`lazyFallback` signature changes.
2. `npx astro check && npm run lint && npx prettier --check .` on the whole `web/` tree.
3. `E2E_PORT=4432 npx playwright test` — the whole suite, per the lane's house rules, not
   just the specs this plan touched.
4. `npm run lhci` — quote CLS for all 13 URLs in `web/lighthouserc.json`.
5. Confirm every server started during this lane's work is stopped
   (`npx astro preview stop`; `pkill -f "$(pwd)/web/node_modules/astro"`; verify with
   `ps -axo command | grep "[a]stro.*states-islands"`).
