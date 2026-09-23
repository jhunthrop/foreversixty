# Character Identity Component + Simulator Cached Session — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development to
> implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the four hand-rolled copies of character avatar/name/descriptor markup
(`CharacterList`, `HomeAccountPanel`, `LandingState`, `Character`/`Account`) with three shared
components (`CharacterPortrait`, `CharacterIdentity`, `CharacterRow`), and move
`SimView.svelte`'s `/v1/me` read off a raw `fetchMe()` in `onMount` onto the shared
`createQueryState` cache every other island already uses.

**Architecture:** `CharacterPortrait` (avatar → class icon over letter square → letter square)
is the base primitive. `CharacterIdentity` wraps it with name + descriptor. `CharacterRow` wraps
`CharacterIdentity` with a build pill, an optional guild line, and a trailing `action` snippet —
the exact row shape `CharacterList` and `LandingState` both need. Every caller is rewritten to
use these instead of its own markup; `lib/account/character-descriptor.ts`'s `classSquare`/
`classIconUrl` are narrowed to the smallest shape they need so `CharacterPortrait`'s prop type
does not require a full `MeCharacter`. Separately, `SimView.svelte` swaps its `onMount(() =>
fetchMe()...)` for the same `createQueryState(...)` call `Account.svelte`/`SessionNav.svelte`/
`HomeAccountPanel.svelte` already make, with a `Skeleton` shown while the session is loading and
no stored character is up yet.

**Tech Stack:** Astro + Svelte 5 (runes) + TypeScript, Vitest (`svelte/server` SSR render tests),
Playwright e2e, Tailwind v4 utility classes.

**Spec:** `docs/superpowers/specs/2026-09-23-character-identity-design.md` (binding). Caching
layer background: `docs/superpowers/specs/2026-09-23-caching-layer-design.md` §3.

## Global Constraints

- `design/DESIGN-SYSTEM.md` in full; states model (`2026-09-22-account-and-island-states-design.md`
  §1): skeleton sized to the ready state, `LoadError` with retry, `EmptyState`, one `.reveal`,
  nothing else moves. Lighthouse budgets in `web/lighthouserc.json` hold; quote every URL's
  numbers before the final report (`/sim` and the home page in particular).
- Vocabulary: addon in game, companion on the desktop, "from Battle.net" for Blizzard-sourced
  data. Every visible string lives in a copy module. No new copy module is created unless a task
  below introduces a genuinely new visible string (none do — every string this plan renders
  already exists in `lib/account/character-list-copy.ts`, `lib/sim/copy.ts`,
  `lib/home-panel-copy.ts`, or is computed from `rulesetLabel`/region code, neither of which is
  copy).
- Caching layer spec §3.3: islands render the session through `createQueryState`
  (`lib/data/query.svelte.ts`); never `fetchMe()` from a component, and never call
  `fetchMeOnce()`/`query()` from an effect that can re-run on the answer.
- Not touched: `Header.astro`, `Footer.astro`, `components/ui/**`, `lib/ui/**`, `lib/data/**`,
  anything under `api/`.
- Lane rules (`.superpowers/journeys/lane-common-web.md`): sonnet only for every subagent, never
  opus (haiku only for a purely mechanical single-file fix — none of these tasks qualify); scoped
  checks before every commit (`npx vitest run <paths>`, `npx astro check`, `npm run lint`,
  `npx prettier --check <paths>`); no broad `pkill`; stop every server started; whole e2e suite
  once before the final report. Commit trailer for this lane:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then
  `Claude-Session: https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY` (this session, not the
  one the journey file names).
- Functions under 50 lines, files under 800, no magic numbers (size/box/font maps are named
  constants), immutable updates, no dead code. Delete every copy of the avatar/fallback markup
  the new components replace — no component keeps a private variant.

## Rulings recorded ahead of execution (see ledger for the full list as work proceeds)

1. **No new copy module.** Every visible string this plan renders already exists in an existing
   copy module or is computed (region code, `rulesetLabel`). `lib/character/identity-copy.ts` is
   not created. Cost if wrong: a two-line follow-up file; no behavior risk.
2. **`CharacterPortrait`'s prop type is `PortraitCharacter = { name: string; class?: string;
   avatar_url?: string }`**, not `MeCharacter`, per spec 2.1's "or the narrower shape it needs".
   `lib/account/character-descriptor.ts`'s `classSquare`/`classIconUrl` are narrowed to accept a
   `ClassIconSubject = { name: string; class?: string }` so both `MeCharacter` and
   `PortraitCharacter` satisfy them structurally — one function, every caller.
3. **`CharacterIdentity` gets an additional `testid?: string` prop** (default `'character'`),
   forwarded to its internal `CharacterPortrait` as `testid`. Spec 2.2 lists `nameTestid?`/
   `descriptorTestid?` but not a portrait-prefix prop; different callers need different avatar
   test-id prefixes (`character-avatar*` vs `home-hero-avatar*` vs `account-hero-avatar*`), and
   `CharacterPortrait` (spec 2.1) already established the prefix-prop pattern. Cost if wrong: a
   rename, no data risk.
4. **`CharacterIdentity` gets an additional `onNameClick?: (event: MouseEvent) => void` prop**,
   applied as the name anchor's `onclick` when `href` is set. Spec 2.3 says CharacterRow passes
   "an onclick on the identity's link via a `onNameClick?` prop" — CharacterIdentity must accept
   it to apply it to its own anchor. Cost if wrong: same, a rename.
5. **`CharacterRow`'s action snippet, for `CharacterList`, contains both the Main-pill/Set-as-main
   control and `<CharacterRowLink>`** (the existing `character-open-sim`/`character-needs-addon`
   element) — not just the Main control alone. Spec 2.4's CharacterList bullet names "the Main /
   Set as main action" but §4 requires the existing `CharacterList.test.ts` and
   `handoffs-account-character.spec.ts` (which assert `character-open-sim`/`character-needs-addon`
   inside each row) to keep passing, and those test ids are `CharacterRowLink`'s. `CharacterRowLink`
   is unchanged and unaffected by "delete every duplicated avatar/fallback copy" (it renders no
   avatar markup). Cost if wrong: reviewer flags an extra element in the snippet, one-line fix.
6. **`CharacterRow`'s build pill gets a `pillTestid?` prop** (`CharacterList` passes
   `'character-build-pill'`, its existing default; `LandingState` passes
   `` `sim-character-build-${character.key}` ``, its existing per-row id) since the two callers'
   existing ids differ in shape and neither can be derived from the row's own `testid` prop
   without breaking the other. Cost if wrong: a rename.
7. **The guild line's test ids (`character-guild-line`, `character-guild-verified`) are hardcoded
   inside `CharacterRow`**, gated by the `guildLine` boolean prop — `LandingState` never sets
   `guildLine`, so no second caller ever needs a different id for it. No extra prop.
8. **Layout normalization inside `CharacterRow`:** the row (`<li>`) carries `min-h-11` (today only
   on `LandingState`'s row, not `CharacterList`'s) and `last:border-b-0` (today only on
   `LandingState`'s), applied to both callers now that they share one component. Neither is
   asserted by any existing test (they check text/test-id presence, not CSS). Cost if wrong: a
   trivial visual nit, caught (if it moves CLS) by the final `npm run lhci` gate this plan's last
   task runs.
9. **Horizontal padding:** `CharacterRow`'s own `<li>` carries no horizontal padding (matches
   `CharacterList`'s current row, which sits inside `StatePanel`'s own `px-4 md:px-6`).
   `LandingState`'s `<ul>` — not nested in a padded panel — gains `px-3` itself (moved off the old
   per-`<li>` `px-3`) so its rows keep a left/right inset; the row divider (`border-b`) is now
   inset by that `px-3` from the panel's outer border, a few pixels short of edge-to-edge, versus
   today. Not asserted by any test. Cost if wrong: a CSS nit, same lhci safety net as #8.
10. **The build pill and guild line render inside a `flex flex-col` wrapper alongside
    `CharacterIdentity`** (not spliced into `CharacterIdentity`'s own name/descriptor column,
    which spec 2.2 defines as a fixed three-part block with no slot for extra rows). They stack
    below the avatar+name line, flush left under the whole block rather than indented to align
    under just the name text — a minor, untested visual simplification from today's
    `CharacterList` (which indents them under the name only). Cost if wrong: same lhci safety net.
11. **`CharacterStrip.svelte` (the loaded-character strip on `/sim`) is out of scope.** It is not
    named in spec §2.4's caller list; its `sim-character-name`/`sim-character-descriptor` test ids
    (asserted by `sim-sources.spec.ts`, `sim-tabs.spec.ts`) belong to it, not to the new
    components, and it is left untouched.
12. **The `/sim` landing skeleton's `min-h` is a first-pass estimate** (`SIM_LANDING_SKELETON_MIN_H
    = 'min-h-[340px]'`, `lib/sim/layout.ts`), not a devtools-measured figure like
    `lib/account/layout.ts`'s constants — this lane has no interactive-browser measurement step
    in its toolchain. The final task's `npm run lhci` run is the check: if `/sim`'s CLS moves
    versus the `main` baseline, that constant is tightened before the final report, not shipped
    on a guess with no verification.

---

## File Structure

**New files:**
- `web/src/components/character/CharacterPortrait.svelte` — the avatar/class-icon/letter-square
  primitive.
- `web/src/components/character/CharacterPortrait.test.ts` — SSR render tests.
- `web/src/components/character/CharacterIdentity.svelte` — portrait + name + descriptor.
- `web/src/components/character/CharacterIdentity.test.ts` — SSR render tests.
- `web/src/components/character/CharacterRow.svelte` — identity + pill + guild line + action.
- `web/src/components/character/CharacterRow.test.ts` — SSR render tests.
- `web/src/lib/sim/layout.ts` — `SIM_LANDING_SKELETON_MIN_H`, the `/sim` landing loading
  skeleton's reserved height (ruling #12).

**Modified files:**
- `web/src/lib/account/character-descriptor.ts` — narrow `classSquare`/`classIconUrl` to accept
  `ClassIconSubject` instead of `MeCharacter`.
- `web/src/components/account/CharacterList.svelte` — rows become `CharacterRow`.
- `web/src/components/sim/LandingState.svelte` — rows become `CharacterRow`.
- `web/src/components/HomeAccountPanel.svelte` — hero becomes `CharacterIdentity`, chips'
  avatars become `CharacterPortrait`.
- `web/src/components/Account.svelte` — hero band's name/descriptor/avatar become
  `CharacterIdentity`.
- `web/src/components/Character.svelte` — header avatar becomes `CharacterPortrait`.
- `web/src/components/sim/SimView.svelte` — `fetchMe()`/`onMount` replaced by
  `createQueryState`; new loading-skeleton branch.
- `web/tests/e2e/sim-landing.spec.ts` — one new returning-visitor test appended (task 8).

**Untouched (confirmed in scope-reading, ruling #11 and Global Constraints):**
`CharacterStrip.svelte`, `CharacterRowLink.svelte`, `CharacterHandoffLinks.svelte`,
`lib/account/build-pill.ts` (imported, not edited), `lib/characters.ts` (imported, not edited),
everything under `components/ui/**`, `lib/ui/**`, `lib/data/**`.

---

## Task 1: `CharacterPortrait.svelte`

**Files:**
- Modify: `web/src/lib/account/character-descriptor.ts`
- Create: `web/src/components/character/CharacterPortrait.svelte`
- Test: `web/src/components/character/CharacterPortrait.test.ts`

**Interfaces:**
- Produces: `PortraitCharacter { name: string; class?: string; avatar_url?: string }`, exported
  from `CharacterPortrait.svelte`. `ClassIconSubject { name: string; class?: string }`, exported
  from `character-descriptor.ts`.
- Produces: `CharacterPortrait` props `{ character: PortraitCharacter; size: 'sm' | 'md' | 'lg';
  testid: string }`. Renders `data-testid="${testid}-avatar"` (avatar `<img>`),
  `"${testid}-avatar-fallback"` (letter-square `<span>`), `"${testid}-class-icon"` (class icon
  `<img>` inside the fallback).
- Consumes: `classSquare`, `classIconUrl` from `lib/account/character-descriptor.ts` (this task
  narrows their parameter type); `classColorVar` from `lib/report/format.ts` (already used
  inside `classSquare`, no new import needed by this component).

- [ ] **Step 1: Narrow `classSquare`/`classIconUrl` in `character-descriptor.ts`**

Replace the two functions' signatures (keep every line of their bodies — only the parameter
type changes) with:

```ts
/** The smallest shape `classSquare`/`classIconUrl` need — every `MeCharacter` satisfies it, and
 *  so does `CharacterPortrait.svelte`'s own narrower `PortraitCharacter` prop type (spec
 *  2026-09-23 §2.1: "type it as the smallest shape and let MeCharacter satisfy it"). */
export interface ClassIconSubject {
  name: string;
  class?: string;
}

export function classSquare(character: ClassIconSubject): { letter: string; color: string } {
  const label = character.class === undefined ? character.name : classDisplayName(character.class);
  return { letter: label.charAt(0).toUpperCase(), color: classColorVar(character.class) };
}
```

(the `classIconUrl` function's body is unchanged — only its signature becomes
`export function classIconUrl(character: ClassIconSubject): string | undefined {`).

- [ ] **Step 2: Run the existing tests to confirm the narrowing is safe**

Run: `npx vitest run src/components/account/CharacterList.test.ts src/components/HomeAccountPanel.test.ts`
Expected: all existing tests still PASS (both callers pass full `MeCharacter` objects, which
still satisfy the narrower parameter type).

- [ ] **Step 3: Write the failing SSR tests for `CharacterPortrait`**

Create `web/src/components/character/CharacterPortrait.test.ts`:

```ts
// web/src/components/character/CharacterPortrait.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterPortrait from './CharacterPortrait.svelte';

describe('CharacterPortrait', () => {
  it('shows the avatar image when avatar_url is set', () => {
    const { body } = render(CharacterPortrait, {
      props: { character: { name: 'Thoradin', class: 'warrior', avatar_url: '/a.jpg' }, size: 'md', testid: 't' },
    });
    expect(body).toContain('data-testid="t-avatar"');
    expect(body).toContain('src="/a.jpg"');
    expect(body).toContain('alt=""');
    expect(body).toContain('loading="lazy"');
    expect(body).not.toContain('data-testid="t-avatar-fallback"');
  });

  it('shows the class icon over the letter square when there is no avatar', () => {
    const { body } = render(CharacterPortrait, {
      props: { character: { name: 'Thoradin', class: 'warrior' }, size: 'md', testid: 't' },
    });
    expect(body).toContain('data-testid="t-avatar-fallback"');
    expect(body).toContain('t-avatar-fallback">T');
    expect(body).toContain('data-testid="t-class-icon"');
    expect(body).toContain('classicon_warrior.jpg');
    expect(body).not.toContain('data-testid="t-avatar"');
  });

  it('shows the letter square alone for a class with no known icon', () => {
    const { body } = render(CharacterPortrait, {
      props: { character: { name: 'Elyra' }, size: 'md', testid: 't' },
    });
    expect(body).toContain('data-testid="t-avatar-fallback"');
    expect(body).toContain('t-avatar-fallback">E');
    expect(body).not.toContain('data-testid="t-class-icon"');
  });

  it('sizes the box for sm, md and lg', () => {
    const sm = render(CharacterPortrait, { props: { character: { name: 'E' }, size: 'sm', testid: 't' } });
    expect(sm.body).toContain('h-7 w-7');
    const md = render(CharacterPortrait, { props: { character: { name: 'E' }, size: 'md', testid: 't' } });
    expect(md.body).toContain('h-9 w-9');
    const lg = render(CharacterPortrait, { props: { character: { name: 'E' }, size: 'lg', testid: 't' } });
    expect(lg.body).toContain('h-11 w-11');
  });
});
```

- [ ] **Step 4: Run it to confirm it fails (the component does not exist yet)**

Run: `npx vitest run src/components/character/CharacterPortrait.test.ts`
Expected: FAIL — cannot find module `./CharacterPortrait.svelte`.

- [ ] **Step 5: Write `CharacterPortrait.svelte`**

```svelte
<!-- web/src/components/character/CharacterPortrait.svelte -->
<!-- The one place that draws a character's avatar/class-icon/letter-square (spec 2026-09-23
     §2.1), replacing four copies of the same markup: `CharacterList.svelte`,
     `HomeAccountPanel.svelte` (hero at 40px and chips at 24px), `Character.svelte`'s public
     header. Renders, in priority order: the avatar (`avatar_url`) -> the class icon
     (`classIconUrl`) over the class-coloured letter square (`classSquare`, so a blank/failed
     icon load still shows the letter) -> the letter square alone. -->
<script lang="ts">
  import { classSquare, classIconUrl } from '../../lib/account/character-descriptor';

  export interface PortraitCharacter {
    name: string;
    class?: string;
    avatar_url?: string;
  }

  /** 28 / 36 / 44px, spec 2026-09-23 §2.1 -- one named map, no magic numbers at the call site. */
  const SIZE_CLASS = {
    sm: { box: 'h-7 w-7', letter: 'text-[12px]' },
    md: { box: 'h-9 w-9', letter: 'text-[15px]' },
    lg: { box: 'h-11 w-11', letter: 'text-[18px]' },
  } as const;

  let {
    character,
    size,
    testid,
  }: { character: PortraitCharacter; size: 'sm' | 'md' | 'lg'; testid: string } = $props();

  const box = $derived(SIZE_CLASS[size].box);
  const letterClass = $derived(SIZE_CLASS[size].letter);
  const square = $derived(classSquare(character));
  const classIcon = $derived(classIconUrl(character));
</script>

{#if character.avatar_url !== undefined}
  <img
    class={`${box} shrink-0 rounded-[3px] object-cover`}
    src={character.avatar_url}
    alt=""
    loading="lazy"
    data-testid={`${testid}-avatar`}
  />
{:else}
  <span
    class={`relative flex ${box} shrink-0 items-center justify-center rounded-[3px] ${letterClass} font-bold`}
    style={`background-color: color-mix(in srgb, ${square.color} 22%, transparent); color: ${square.color}`}
    data-testid={`${testid}-avatar-fallback`}
  >
    {square.letter}
    {#if classIcon !== undefined}
      <img
        class={`absolute inset-0 ${box} rounded-[3px] object-cover`}
        src={classIcon}
        alt=""
        loading="lazy"
        data-testid={`${testid}-class-icon`}
      />
    {/if}
  </span>
{/if}
```

- [ ] **Step 6: Run the test to confirm it passes**

Run: `npx vitest run src/components/character/CharacterPortrait.test.ts`
Expected: PASS, 4/4.

- [ ] **Step 7: Scoped checks and commit**

Run, from `web/`:
```
npx vitest run src/components/character/CharacterPortrait.test.ts src/components/account/CharacterList.test.ts src/components/HomeAccountPanel.test.ts
npx astro check
npm run lint
npx prettier --check src/components/character/CharacterPortrait.svelte src/components/character/CharacterPortrait.test.ts src/lib/account/character-descriptor.ts
```
Expected: all green.

Write the commit message to a file under this worktree's `.superpowers/` with `printf`, then
`git add web/src/lib/account/character-descriptor.ts web/src/components/character/CharacterPortrait.svelte web/src/components/character/CharacterPortrait.test.ts && git commit -F <file>`
as its own Bash command. Subject: `feat(web): add the shared CharacterPortrait component`.

---

## Task 2: `CharacterIdentity.svelte`

**Files:**
- Create: `web/src/components/character/CharacterIdentity.svelte`
- Test: `web/src/components/character/CharacterIdentity.test.ts`

**Interfaces:**
- Consumes: `CharacterPortrait` (Task 1) with props `{character, size, testid}`.
  `characterDescriptor(character: MeCharacter): string` from
  `lib/account/character-descriptor.ts` (unchanged). `rulesetLabel(id: string): string` and
  `MeCharacter` type from `lib/characters.ts` / `lib/account/api.ts`. `classColorVar(cls?:
  string): string` from `lib/report/format.ts`.
- Produces: `CharacterIdentity` props `{ character: MeCharacter; size: 'sm' | 'md' | 'lg';
  descriptor: 'full' | 'realm' | 'none'; href?: string; onNameClick?: (event: MouseEvent) =>
  void; nameTestid?: string; descriptorTestid?: string; testid?: string }` (`testid` defaults to
  `'character'`, ruling #3; `onNameClick`, ruling #4). Renders the name as `<a>` when `href` is
  given, `<span>` otherwise; the descriptor line only when non-empty.

- [ ] **Step 1: Write the failing SSR tests**

Create `web/src/components/character/CharacterIdentity.test.ts`:

```ts
// web/src/components/character/CharacterIdentity.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterIdentity from './CharacterIdentity.svelte';
import type { MeCharacter } from '../../lib/account/api';

const CHAR: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'priest',
  race: 'Night Elf',
  realm: 'Whitemane',
  level: 60,
};

describe('CharacterIdentity', () => {
  it('renders the full descriptor', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'full' },
    });
    expect(body).toContain('Night Elf Priest · Level 60 · Whitemane (Hardcore US)');
  });

  it('renders the realm descriptor from ruleset and region alone', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'realm' },
    });
    expect(body).toContain('Hardcore · US');
    expect(body).not.toContain('Night Elf');
  });

  it('renders no descriptor line for "none"', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', descriptorTestid: 'd' },
    });
    expect(body).not.toContain('data-testid="d"');
  });

  it('renders the name as plain text with no href', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', nameTestid: 'n' },
    });
    expect(body).toContain('<span');
    expect(body).toContain('data-testid="n"');
    expect(body).not.toContain('<a');
  });

  it('renders the name as a link with href, in the display face only at lg', () => {
    const md = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', href: '/character/x', nameTestid: 'n' },
    });
    expect(md.body).toContain('<a');
    expect(md.body).toContain('href="/character/x"');
    expect(md.body).not.toContain('font-display');

    const lg = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', href: '/character/x' },
    });
    expect(lg.body).toContain('font-display');
  });

  it('uses a caller-chosen portrait test id prefix, defaulting to "character"', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none' },
    });
    expect(body).toContain('data-testid="character-avatar-fallback"');

    const custom = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', testid: 'home-hero' },
    });
    expect(custom.body).toContain('data-testid="home-hero-avatar-fallback"');
  });
});
```

- [ ] **Step 2: Run it to confirm it fails**

Run: `npx vitest run src/components/character/CharacterIdentity.test.ts`
Expected: FAIL — module not found.

- [ ] **Step 3: Write `CharacterIdentity.svelte`**

```svelte
<!-- web/src/components/character/CharacterIdentity.svelte -->
<!-- Portrait + name + descriptor as one block (spec 2026-09-23 §2.2): the account hero band,
     the home hub's hero, and (through CharacterRow) every character-list row draw their name
     and descriptor line through this, instead of four hand-rolled copies. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { characterDescriptor } from '../../lib/account/character-descriptor';
  import { rulesetLabel } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import CharacterPortrait from './CharacterPortrait.svelte';

  /** 14 / 15 / 18px, spec 2026-09-23 §2.2. */
  const NAME_SIZE = { sm: 'text-[14px]', md: 'text-[15px]', lg: 'text-[18px]' } as const;

  let {
    character,
    size,
    descriptor,
    href,
    onNameClick,
    nameTestid,
    descriptorTestid,
    testid = 'character',
  }: {
    character: MeCharacter;
    size: 'sm' | 'md' | 'lg';
    descriptor: 'full' | 'realm' | 'none';
    href?: string;
    onNameClick?: (event: MouseEvent) => void;
    nameTestid?: string;
    descriptorTestid?: string;
    testid?: string;
  } = $props();

  const line = $derived(
    descriptor === 'full'
      ? characterDescriptor(character)
      : descriptor === 'realm'
        ? `${rulesetLabel(character.ruleset)} · ${character.region.toUpperCase()}`
        : '',
  );
  // Display font only at lg (spec 2026-09-23 §2.2): the account hero band is the one place
  // CharacterIdentity's own name needs it -- every <h1>/<h2>/.section-title elsewhere already
  // carries it globally (global.css), and a row/chip name is never that prominent.
  const nameClass = $derived(
    `w-fit ${NAME_SIZE[size]} font-semibold${size === 'lg' ? ' [font-family:var(--font-display)]' : ''}`,
  );
</script>

<div class="flex min-w-0 items-center gap-3">
  <CharacterPortrait {character} {size} {testid} />
  <div class="flex min-w-0 flex-col gap-0.5">
    {#if href !== undefined}
      <a class={nameClass} style:color={classColorVar(character.class)} {href} onclick={onNameClick} data-testid={nameTestid}>
        {character.name}
      </a>
    {:else}
      <span class={nameClass} style:color={classColorVar(character.class)} data-testid={nameTestid}>
        {character.name}
      </span>
    {/if}
    {#if line !== ''}
      <span class="text-muted text-[13px]" data-testid={descriptorTestid}>{line}</span>
    {/if}
  </div>
</div>
```

- [ ] **Step 4: Run the test to confirm it passes**

Run: `npx vitest run src/components/character/CharacterIdentity.test.ts`
Expected: PASS, 6/6.

- [ ] **Step 5: Scoped checks and commit**

Run: `npx vitest run src/components/character/`, `npx astro check`, `npm run lint`,
`npx prettier --check src/components/character/CharacterIdentity.svelte src/components/character/CharacterIdentity.test.ts`.
Commit subject: `feat(web): add the shared CharacterIdentity component`.

---

## Task 3: `CharacterRow.svelte`

**Files:**
- Create: `web/src/components/character/CharacterRow.svelte`
- Test: `web/src/components/character/CharacterRow.test.ts`

**Interfaces:**
- Consumes: `CharacterIdentity` (Task 2). `buildSourcePill` from `lib/account/build-pill.ts`
  (unchanged). `guildRankLabel` from `lib/characters.ts`. `characterListCopy.verified` from
  `lib/account/character-list-copy.ts`.
- Produces: `CharacterRow` props `{ character: MeCharacter; descriptor?: 'full' | 'realm' |
  'none'; guildLine?: boolean; pillTestid?: string; href?: string; onNameClick?: (event:
  MouseEvent) => void; nameTestid?: string; descriptorTestid?: string; testid?: string; action?:
  Snippet }` (`descriptor` defaults `'realm'`, `guildLine` defaults `false`, `pillTestid` defaults
  `'character-build-pill'`, rulings #5-#10). Renders one `<li>`.

- [ ] **Step 1: Write the failing SSR tests**

Create `web/src/components/character/CharacterRow.test.ts`:

```ts
// web/src/components/character/CharacterRow.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterRow from './CharacterRow.svelte';
import type { MeCharacter } from '../../lib/account/api';

const GUILDED: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'priest',
  guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', verified: true },
};

const PLAIN: MeCharacter = {
  key: 'us/pvp/thoradin',
  region: 'us',
  ruleset: 'pvp',
  name: 'Thoradin',
  class: 'warrior',
  build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
};

describe('CharacterRow', () => {
  it('renders one <li> with the realm descriptor by default', () => {
    const { body } = render(CharacterRow, { props: { character: PLAIN } });
    expect(body).toContain('<li');
    expect(body).toContain('PvP · US');
  });

  it('renders the full descriptor when asked', () => {
    const { body } = render(CharacterRow, { props: { character: PLAIN, descriptor: 'full' } });
    expect(body).toContain('Warrior');
    expect(body).not.toContain('PvP · US');
  });

  it('renders the build pill with the given test id, addon vs no build', () => {
    const withBuild = render(CharacterRow, { props: { character: PLAIN, pillTestid: 'p' } });
    expect(withBuild.body).toContain('data-testid="p">Addon');

    const withoutBuild = render(CharacterRow, { props: { character: GUILDED, pillTestid: 'p' } });
    expect(withoutBuild.body).toContain('data-testid="p">No build yet');
  });

  it('renders the guild line and verified pill only when guildLine is true', () => {
    const on = render(CharacterRow, { props: { character: GUILDED, guildLine: true } });
    expect(on.body).toContain('data-testid="character-guild-line"');
    expect(on.body).toContain('data-testid="character-guild-verified"');

    const off = render(CharacterRow, { props: { character: GUILDED } });
    expect(off.body).not.toContain('data-testid="character-guild-line"');
  });

  it('omits the guild line for a character with no guild even when guildLine is true', () => {
    const { body } = render(CharacterRow, { props: { character: PLAIN, guildLine: true } });
    expect(body).not.toContain('data-testid="character-guild-line"');
  });

  it('renders the name with the given href and nameTestid', () => {
    const { body } = render(CharacterRow, {
      props: { character: PLAIN, href: '/character/x', nameTestid: 'n' },
    });
    expect(body).toContain('href="/character/x"');
    expect(body).toContain('data-testid="n"');
  });
});
```

- [ ] **Step 2: Run it to confirm it fails**

Run: `npx vitest run src/components/character/CharacterRow.test.ts`
Expected: FAIL — module not found.

- [ ] **Step 3: Write `CharacterRow.svelte`**

```svelte
<!-- web/src/components/character/CharacterRow.svelte -->
<!-- The list row every character list renders (spec 2026-09-23 §2.3): CharacterIdentity plus
     an optional build pill, an optional guild line, and a trailing action -- the exact row
     `CharacterList.svelte` and `LandingState.svelte` both draw, unified into one component. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { buildSourcePill } from '../../lib/account/build-pill';
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { guildRankLabel } from '../../lib/characters';
  import CharacterIdentity from './CharacterIdentity.svelte';

  let {
    character,
    descriptor = 'realm',
    guildLine = false,
    pillTestid = 'character-build-pill',
    href,
    onNameClick,
    nameTestid,
    descriptorTestid,
    testid,
    action,
  }: {
    character: MeCharacter;
    descriptor?: 'full' | 'realm' | 'none';
    guildLine?: boolean;
    pillTestid?: string;
    href?: string;
    onNameClick?: (event: MouseEvent) => void;
    nameTestid?: string;
    descriptorTestid?: string;
    testid?: string;
    action?: Snippet;
  } = $props();

  const pill = $derived(buildSourcePill(character.build));
  const showGuildLine = $derived(guildLine && character.guild !== undefined);
</script>

<li
  class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-3 text-[14px] last:border-b-0"
  data-testid={testid}
>
  <div class="flex min-w-0 flex-1 flex-col gap-0.5">
    <CharacterIdentity {character} size="md" {descriptor} {href} {onNameClick} {nameTestid} {descriptorTestid} />
    <span
      class={pill.pillClass === null ? 'text-muted text-[12px]' : `pill ${pill.pillClass} w-fit`}
      data-testid={pillTestid}
    >
      {pill.label}
    </span>
    {#if showGuildLine && character.guild !== undefined}
      <span class="text-muted text-[13px]" data-testid="character-guild-line">
        {character.guild.name}
        {#if character.guild.rank !== undefined}· {guildRankLabel(character.guild.rank)}{/if}
        {#if character.guild.verified}
          <span class="text-strong" data-testid="character-guild-verified">{characterListCopy.verified}</span>
        {/if}
      </span>
    {/if}
  </div>
  {#if action !== undefined}{@render action()}{/if}
</li>
```

- [ ] **Step 4: Run the test to confirm it passes**

Run: `npx vitest run src/components/character/CharacterRow.test.ts`
Expected: PASS, 6/6.

- [ ] **Step 5: Scoped checks and commit**

Run: `npx vitest run src/components/character/`, `npx astro check`, `npm run lint`,
`npx prettier --check src/components/character/CharacterRow.svelte src/components/character/CharacterRow.test.ts`.
Commit subject: `feat(web): add the shared CharacterRow component`.

---

## Task 4: `CharacterList.svelte` uses `CharacterRow`

**Files:**
- Modify: `web/src/components/account/CharacterList.svelte`
- Test: `web/src/components/account/CharacterList.test.ts` (existing — must keep passing unmodified)

**Interfaces:**
- Consumes: `CharacterRow` (Task 3) with `descriptor="full"`, `guildLine={true}`,
  `pillTestid="character-build-pill"`, `href`, `descriptorTestid="character-descriptor"`, and an
  `action` snippet containing the existing Main-pill/Set-as-main control plus
  `<CharacterRowLink>` (ruling #5).

- [ ] **Step 1: Rewrite the character list `<li>` loop**

In `CharacterList.svelte`, replace the imports block:

```ts
import type { MeCharacter } from '../../lib/account/api';
import { battlenetStartUrl } from '../../lib/account/api';
import { characterListCopy } from '../../lib/account/character-list-copy';
import { characterHref } from '../../lib/characters';
import { relativeTime } from '../../lib/dates';
import CharacterRow from '../character/CharacterRow.svelte';
import CharacterRowLink from './CharacterRowLink.svelte';
import EmptyState from '../ui/EmptyState.svelte';
import StatePanel from '../ui/StatePanel.svelte';
```

(drop `classSquare`, `classIconUrl`, `guildRankLabel`, `classColorVar`, `buildSourcePill` —
`CharacterRow` computes all of these itself now). Keep every other script line (`settingMain`,
`setMainError`, `setMain`, `REFRESH_HREF`, the `importedAside` snippet) unchanged.

Replace the `{#each characters as character (character.key)}...{/each}` block with:

```svelte
{#each characters as character (character.key)}
  <CharacterRow
    {character}
    descriptor="full"
    guildLine
    href={characterHref(character.region, character.ruleset, character.name)}
    descriptorTestid="character-descriptor"
  >
    {#snippet action()}
      {#if mainKey === character.key}
        <span class="pill pill-site" data-testid="character-main-pill">{characterListCopy.main}</span>
      {:else if onSetMain !== undefined}
        <button
          type="button"
          class="text-nav inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0"
          onclick={() => void setMain(character.key)}
          disabled={settingMain !== ''}
          aria-busy={settingMain === character.key}
          data-testid="character-set-main"
        >
          {characterListCopy.setAsMain}
        </button>
      {/if}
      <CharacterRowLink {character} />
    {/snippet}
  </CharacterRow>
{/each}
```

The surrounding `{#if characters.length === 0}...{:else}<ul class="flex flex-col">...</ul>{/if}`
wrapper is unchanged.

- [ ] **Step 2: Run the existing test suite (no changes to the test file)**

Run: `npx vitest run src/components/account/CharacterList.test.ts`
Expected: PASS, all cases unchanged, including `character-main-pill`, `character-set-main`,
`character-open-sim`, `character-needs-addon`, `character-guild-line`,
`character-guild-verified`, `character-descriptor`, `character-avatar*`, `character-class-icon`,
`character-build-pill`.

- [ ] **Step 3: Scoped checks and commit**

Run: `npx vitest run src/components/account/`, `npx astro check`, `npm run lint`,
`npx prettier --check src/components/account/CharacterList.svelte`.
Commit subject: `refactor(web): CharacterList rows draw through the shared CharacterRow`.

---

## Task 5: `LandingState.svelte` uses `CharacterRow`

**Files:**
- Modify: `web/src/components/sim/LandingState.svelte`
- Test: `web/tests/e2e/sim-landing.spec.ts` (existing — must keep passing unmodified in this task;
  task 8 adds one new test to this file)

**Interfaces:**
- Consumes: `CharacterRow` (Task 3), default `descriptor="realm"`, `guildLine` omitted (false),
  `pillTestid={`sim-character-build-${character.key}`}`, `testid={`sim-character-${character.key}`}`,
  `nameTestid={`sim-character-link-${character.key}`}`, `href={hrefFor(character)}`,
  `onNameClick` wired to the existing `follow` handler, and an `action` snippet containing the
  existing "Sim" button.

- [ ] **Step 1: Rewrite `LandingState.svelte`**

Replace the imports block:

```ts
import type { MeCharacter } from '../../lib/account/api';
import type { CharacterPath } from '../../lib/characters';
import { parseCharacterPath } from '../../lib/characters';
import { simCopy } from '../../lib/sim/copy';
import { BUSY_CLASS } from '../../lib/ui/busy';
import { defaultSimState, simSearch, withSimState } from '../../lib/sim/url';
import CharacterRow from '../character/CharacterRow.svelte';
```

(drop `buildSourcePill`, `rulesetLabel`, `classColorVar` — `CharacterRow` covers the pill and the
realm descriptor; `classColorVar` was only used for the swatch, which is gone). Keep `pathOf`,
`hrefFor`, and `follow` unchanged.

Replace the `<ul>` body:

```svelte
<ul class="border-line bg-raised rounded-panel flex flex-col border px-3">
  {#each characters as character (character.key)}
    {@const path = pathOf(character)}
    <CharacterRow
      {character}
      href={hrefFor(character)}
      onNameClick={(event) => path !== null && follow(event, path)}
      nameTestid={`sim-character-link-${character.key}`}
      pillTestid={`sim-character-build-${character.key}`}
      testid={`sim-character-${character.key}`}
    >
      {#snippet action()}
        <button
          type="button"
          class={`border-line-warm-strong rounded-control text-strong label ml-auto min-h-11 shrink-0 border px-4 disabled:opacity-50 md:min-h-9 ${busyKey === character.key ? BUSY_CLASS : ''}`}
          disabled={busyKey !== null || path === null}
          aria-busy={busyKey === character.key}
          onclick={() => path !== null && onpick(path)}
          data-testid={`sim-pick-${character.key}`}
        >
          {simCopy.simIt}
        </button>
      {/snippet}
    </CharacterRow>
  {/each}
</ul>
```

(ruling #9: `px-3` moves from each `<li>` onto the `<ul>` itself).

- [ ] **Step 2: Run the e2e spec for this component**

From `web/`: `E2E_PORT=4381 npx playwright test tests/e2e/sim-landing.spec.ts`
Expected: all tests PASS unmodified, including the `sim-character-link-...` href assertion and
the busy/disabled Sim-button behavior.

- [ ] **Step 3: Scoped checks and commit**

Run: `npx astro check`, `npm run lint`, `npx prettier --check src/components/sim/LandingState.svelte`.
Commit subject: `refactor(web): sim landing rows draw through the shared CharacterRow`.

---

## Task 6: `HomeAccountPanel.svelte` uses `CharacterIdentity`/`CharacterPortrait`

**Files:**
- Modify: `web/src/components/HomeAccountPanel.svelte`
- Test: `web/src/components/HomeAccountPanel.test.ts` (existing), `web/tests/e2e/home-panel.spec.ts`
  (existing — must keep passing unmodified)

**Interfaces:**
- Consumes: `CharacterIdentity` (Task 2) for the hero, at `size="md"`, `descriptor="full"`, no
  `href` (the hero name is not a link today — unchanged), `testid="home-hero"` (ruling #3, so the
  portrait renders `home-hero-avatar`/`home-hero-avatar-fallback`). `CharacterPortrait` (Task 1)
  for each chip, at `size="sm"`, `testid="home-chip"`.

- [ ] **Step 1: Rewrite the hero block**

Add the import: `import CharacterIdentity from './character/CharacterIdentity.svelte';` and
`import CharacterPortrait from './character/CharacterPortrait.svelte';`. Remove the now-unused
`classSquare, classIconUrl, characterDescriptor` import from `'../lib/account/character-descriptor'`
and the `classColorVar` import stays (still used for the chip name's colour). Remove the local
`descriptor`, `square`, `classIcon` `$derived`s (Task deletes lines 104-106 of the original file);
`heroPath` stays.

Replace the hero markup (the block from `{#if hero.avatar_url !== undefined}` through the closing
`</div>` that held the name/descriptor, i.e. everything between the opening
`data-testid="home-account-panel"` div and the `<a href="/planner">` link) with:

```svelte
<CharacterIdentity character={hero} size="md" descriptor="full" testid="home-hero" />
```

- [ ] **Step 2: Rewrite each chip's avatar**

Replace the chip's `{#if other.avatar_url !== undefined}...{/if}` block (the avatar/fallback
markup inside each `<button>`, before the name `<span>`) with:

```svelte
<CharacterPortrait character={other} size="sm" testid="home-chip" />
```

The name `<span>` and the level `<span>` after it are unchanged.

- [ ] **Step 3: Run the existing tests**

Run: `npx vitest run src/components/HomeAccountPanel.test.ts`
Expected: PASS (this test only checks the pre-session SSR shell, unaffected).

From `web/`: `E2E_PORT=4381 npx playwright test tests/e2e/home-panel.spec.ts`
Expected: PASS — every test in this file, including the rating figure, the chip click/switch,
and the returning-visitor snapshot test.

- [ ] **Step 4: Scoped checks and commit**

Run: `npx astro check`, `npm run lint`, `npx prettier --check src/components/HomeAccountPanel.svelte`.
Commit subject: `refactor(web): home hero and chips draw through the shared identity components`.

---

## Task 7: `Account.svelte` hero band and `Character.svelte` header use the shared components

**Files:**
- Modify: `web/src/components/Account.svelte`
- Modify: `web/src/components/Character.svelte`
- Test: `web/tests/e2e/auth.spec.ts`, `web/tests/e2e/bnet-hub.spec.ts`,
  `web/tests/e2e/character-render.spec.ts`, `web/tests/e2e/handoffs-account-character.spec.ts`
  (existing — must keep passing unmodified)

**Interfaces:**
- Consumes: `CharacterIdentity` (Task 2) in `Account.svelte`'s hero band, `size="lg"`,
  `descriptor="full"`, `href={characterHref(hero.region, hero.ruleset, hero.name)}`,
  `testid="account-hero"` (so the portrait renders `account-hero-avatar`/
  `account-hero-avatar-fallback` — new elements; the hero band shows no avatar today).
  `CharacterPortrait` (Task 1) in `Character.svelte`'s header, `size="lg"`, `testid="character"`.

- [ ] **Step 1: Rewrite `Account.svelte`'s hero band**

Add `import CharacterIdentity from './character/CharacterIdentity.svelte';`. Remove the
`characterDescriptor` import from `'../lib/account/character-descriptor'` (no longer called
directly) and the `characterHref` import stays (still used for `href`); `classColorVar` stays
(unused after this change only if nothing else in the file needs it — check before removing: it
is also used by `CharacterList`'s import path, not this file, so remove it from `Account.svelte`
only if no other usage remains in that file after this edit — grep the file first).

Replace:

```svelte
<div class="flex flex-col gap-1">
  <a
    class="w-fit [font-family:var(--font-display)] text-[15px] font-semibold"
    style:color={classColorVar(hero.class)}
    href={characterHref(hero.region, hero.ruleset, hero.name)}
  >
    {hero.name}
  </a>
  <span class="text-muted text-[13px]">{characterDescriptor(hero)}</span>
  {#if heroPath !== null}
```

with:

```svelte
<div class="flex flex-col gap-1">
  <CharacterIdentity
    character={hero}
    size="lg"
    descriptor="full"
    href={characterHref(hero.region, hero.ruleset, hero.name)}
    testid="account-hero"
  />
  {#if heroPath !== null}
```

(everything from `<CharacterHandoffLinks path={heroPath} />` onward, and the closing tags, is
unchanged).

- [ ] **Step 2: Rewrite `Character.svelte`'s header avatar**

Add `import CharacterPortrait from './character/CharacterPortrait.svelte';`. Replace:

```svelte
{#if data.character.render_url === undefined && data.character.avatar_url !== undefined}
  <img
    class="h-11 w-11 shrink-0 rounded-[3px] object-cover"
    src={data.character.avatar_url}
    alt=""
    loading="lazy"
    data-testid="character-avatar"
  />
{/if}
```

with:

```svelte
{#if data.character.render_url === undefined}
  <CharacterPortrait character={data.character} size="lg" testid="character" />
{/if}
```

(this now shows the letter-square fallback for a character with neither `render_url` nor
`avatar_url`, instead of nothing — an intentional improvement per spec §2.4; the existing
`character-render.spec.ts` "shows neither" test only asserts `character-render` and
`character-avatar` are both absent, which still holds: the new element is
`character-avatar-fallback`).

- [ ] **Step 3: Run the e2e specs**

From `web/`:
```
E2E_PORT=4381 npx playwright test tests/e2e/auth.spec.ts tests/e2e/bnet-hub.spec.ts tests/e2e/character-render.spec.ts tests/e2e/handoffs-account-character.spec.ts
```
Expected: all PASS.

- [ ] **Step 4: Scoped checks and commit**

Run: `npx astro check`, `npm run lint`, `npx prettier --check src/components/Account.svelte src/components/Character.svelte`.
Commit subject: `refactor(web): account hero and character header draw through the shared identity components`.

---

## Task 8: `SimView.svelte` reads the session from the cache

**Files:**
- Create: `web/src/lib/sim/layout.ts`
- Modify: `web/src/components/sim/SimView.svelte`
- Modify: `web/tests/e2e/sim-landing.spec.ts` (append one test)

**Interfaces:**
- Consumes: `createQueryState<Me | null>(key, load, options)` from `lib/data/query.svelte.ts`
  (unchanged, already used by `Account.svelte`/`SessionNav.svelte`/`HomeAccountPanel.svelte`).
  `fetchMeOnce` from `lib/account/api.ts` (unchanged). `effectiveServerSims` from
  `lib/account/api.ts` (unchanged).
- Produces: `SIM_LANDING_SKELETON_MIN_H` (`lib/sim/layout.ts`, ruling #12).

- [ ] **Step 1: Add `web/src/lib/sim/layout.ts`**

```ts
// web/src/lib/sim/layout.ts
// The /sim landing area's loading floor (spec 2026-09-23 §3): while the session read is in
// flight and no character is loaded yet, the landing area shows a Skeleton reserving roughly
// LandingState.svelte's own rendered height for four rows, so the reveal changes only opacity
// -- lib/account/layout.ts's own reason, same pattern. First-pass estimate (ruling #12 in this
// lane's plan): this toolchain has no interactive-browser measurement step, so this is sized by
// component-height arithmetic (the h2, the bordered row list at four rows, the scope note, the
// "Sim something else" button) rather than a devtools capture. Confirmed or tightened against
// `npm run lhci`'s own /sim CLS number before the branch's final report.
export const SIM_LANDING_SKELETON_MIN_H = 'min-h-[340px]';
```

- [ ] **Step 2: Replace `SimView.svelte`'s session read**

Change the import line:

```ts
import { battlenetStartUrl, effectiveServerSims, fetchMeOnce, type Me } from '../../lib/account/api';
```

(was `fetchMe`, now `fetchMeOnce`). Add:

```ts
import { createQueryState } from '../../lib/data/query.svelte';
import { API_BASE_URL } from '../../lib/planner/config';
import { SIM_LANDING_SKELETON_MIN_H } from '../../lib/sim/layout';
```

Replace:

```ts
  // Set from `onMount`'s own `fetchMe` below -- read by the history panel, the landing
  // state and the source switcher's signed-in card.
  let me = $state<Me | null>(null);
```

with:

```ts
  // One `/v1/me` read, shared with every other island through the client cache
  // (web/src/lib/data/query.ts) -- see Account.svelte, SessionNav.svelte and
  // HomeAccountPanel.svelte's own copies of this same call. Read by the history panel, the
  // landing state and the source switcher's signed-in card.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
```

Replace the `onMount`'s `fetchMe` call. Before:

```ts
  onMount(() => {
    // effectiveServerSims(me) on GET /v1/me -- the server lane renders only once this answers
    // true. A signed-out visitor and an unreachable API read the same way (fetchMe resolves
    // null, or the promise rejects and is swallowed): both mean "no premium control", matching
    // Account.svelte's own load() treating a failed fetchMe as "not signed in", not an error.
    //
    // The same answer also gates the history panel (Task 17), the landing state and the
    // source switcher's signed-in card (Task 18): `signedIn` above is `me !== null`, and
    // the history panel's own `GET /v1/sims?mine=1` fires only then -- a signed-out visitor
    // gets no second request for a list that would come back empty anyway.
    void fetchMe()
      .then((result) => {
        me = result;
        store.setPremium(effectiveServerSims(result));
        if (result !== null) void loadHistory();
      })
      .catch(() => {});
    // Never on /sim/<id>: a saved sim never restores the visitor's own current character.
    if (!hasSavedSimId) void restoreFromPointer();
    return () => store.dispose();
  });
```

After:

```ts
  // effectiveServerSims(me) on GET /v1/me -- the server lane renders only once this answers
  // true. A signed-out visitor and an unreachable API read the same way now (`me` stays null
  // whether the session read answered null or failed -- the sim never shows a session error,
  // spec 2026-09-23 §3), matching Account.svelte's own hero treating a failed read as "not
  // signed in", not an error.
  //
  // The same answer also gates the history panel (Task 17), the landing state and the source
  // switcher's signed-in card (Task 18): `signedIn` above is `me !== null`, and the history
  // panel's own `GET /v1/sims?mine=1` fires only then. Keyed on `me`'s own reference and
  // guarded against re-running for the same object (spec 2026-09-23 §3, the 2026-09-23
  // account-page loop this guards against) -- and this never calls a session read itself,
  // only reacts to one `createQueryState` already made above.
  let lastMe: Me | null = null;
  $effect(() => {
    if (me === lastMe) return;
    lastMe = me;
    store.setPremium(effectiveServerSims(me));
    if (me !== null) void loadHistory();
  });

  onMount(() => {
    // Never on /sim/<id>: a saved sim never restores the visitor's own current character.
    if (!hasSavedSimId) void restoreFromPointer();
    return () => store.dispose();
  });
```

- [ ] **Step 3: Add the loading-skeleton branch to the landing area**

In the template, the branch chain reads (today):

```svelte
      {#if store.character !== null && !switcherOpen}
        <CharacterStrip ... />
      {:else if me !== null && me.characters.length > 0 && !switcherOpen}
        <LandingState ... />
        ...
      {:else if me !== null && me.characters.length === 0}
        ...
      {:else}
        <SourceSwitcher ... />
      {/if}
```

Insert a new branch immediately after the `CharacterStrip` branch's condition, before the
`LandingState` branch:

```svelte
      {#if store.character !== null && !switcherOpen}
        <CharacterStrip
          character={store.character}
          items={store.items}
          races={store.races}
          onchange={() => (switcherOpen = true)}
          onrace={(slug) => store.setRace(slug)}
        />
      {:else if session.status === 'loading' && session.data === null && !switcherOpen}
        <!-- Spec 2026-09-23 §3: no live-round-trip flash while a returning signed-in visitor's
             own snapshot is still loading (createQueryState's ttlMs keeps it instant on a
             repeat visit, so this only ever shows on a cold cache) or a genuinely fresh session
             read is in flight. -->
        <Skeleton lines={4} minHeight={SIM_LANDING_SKELETON_MIN_H} testid="sim-landing-skeleton" />
      {:else if me !== null && me.characters.length > 0 && !switcherOpen}
```

(the rest of the chain is unchanged).

- [ ] **Step 4: Add the returning-visitor Playwright test**

Append to `web/tests/e2e/sim-landing.spec.ts` (same shape as `home-panel.spec.ts`'s own
returning-visitor test):

```ts
test('a returning signed-in visitor sees their characters from the session snapshot before /v1/me answers', async ({
  page,
  context,
}) => {
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/characters/**', (route) => route.fulfill(failure('none', 404)));
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.goto('/sim');
  await expect(page.getByTestId('sim-landing')).toBeVisible();
  const stored = await page.evaluate(() =>
    Object.keys(window.localStorage)
      .filter((k) => k.startsWith('fs.q.'))
      .map((k) => window.localStorage.getItem(k) ?? '')
      .join('\n'),
  );
  expect(stored).toContain('Thrallgar');

  // Second load: /v1/me is held for five seconds, yet the landing state renders at once from
  // the snapshot.
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 5000));
    await route.fulfill(envelope(ME));
  });
  await page.goto('/sim');
  await expect(page.getByTestId('sim-landing')).toBeVisible({ timeout: 3000 });
  await expect(page.getByTestId('sim-landing-skeleton')).toHaveCount(0);
});
```

Add this test inside `test.describe('a signed-in member with characters', ...)`'s block (it uses
`stubSignedIn`'s `beforeEach`? No — it needs its own `/v1/me` control, matching
`home-panel.spec.ts`'s own returning-visitor test living outside that file's per-suite
`beforeEach`). Place it as a standalone `test(...)` after the `describe` block closes (same
indentation level as the `'a signed-out visitor...'` test below it), and route
`**/v1/characters/us/normal/thrallgar/sim-input` and `.../roland/sim-input` are not needed since
this test never picks a character.

- [ ] **Step 5: Run the sim e2e specs**

From `web/`:
```
E2E_PORT=4381 npx playwright test tests/e2e/sim-landing.spec.ts tests/e2e/sim-sources.spec.ts tests/e2e/sim-history.spec.ts tests/e2e/sim-tabs.spec.ts
```
Expected: all PASS, including the new returning-visitor test.

- [ ] **Step 6: Scoped checks and commit**

Run: `npx astro check`, `npm run lint`,
`npx prettier --check src/components/sim/SimView.svelte src/lib/sim/layout.ts tests/e2e/sim-landing.spec.ts`.
Commit subject: `fix(web): the simulator reads the session from the cache, not a raw fetchMe`.

---

## Task 9: Whole-branch verification

This task is run by the controller, not a fresh implementer subagent — it is the "whole e2e
suite once before the final report" step every lane's house rules require, plus the plan's own
Lighthouse and unit-test sweep. No commit is expected from it unless it turns up a real,
in-scope regression, in which case fix it as its own task (fresh implementer, task review, same
loop as above) before writing the final report.

- [ ] **Step 1:** `npx vitest run` (whole suite, not just this lane's new files)
- [ ] **Step 2:** `npx astro check`
- [ ] **Step 3:** `npm run lint`
- [ ] **Step 4:** `npx prettier --check .`
- [ ] **Step 5:** `E2E_PORT=4381 npx playwright test` (the whole e2e suite)
- [ ] **Step 6:** `npm run build` then `npm run lhci`, quoting every URL's performance/
  accessibility/SEO scores, LCP, TBT and CLS in the final report, with particular attention to
  `/sim.html` and `/index.html` (CLS/TBT must not move versus `main`'s own numbers).
- [ ] **Step 7:** `npx astro preview stop`, confirm no `astro dev`/`astro preview` process for
  this worktree remains (`ps -axo command | grep "[a]stro.*character-identity"`).
