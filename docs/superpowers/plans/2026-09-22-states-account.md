# The account and sign-in experience Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix the account/sign-in UX findings (F1–F8) and give `/account` the three-state
model (loading/failed/empty via `Skeleton`/`LoadError`/`EmptyState`) per spec section 2.

**Architecture:** No new fetches, no new routes. Reorder and regroup `Account.svelte`'s
existing `account` mode into three visual blocks (Characters standalone; Devices+You in one
panel; Guilds+Billing+Reports in a second panel), replace ad-hoc loading/error/empty text
with the shared primitives, reuse `SignInPrompt` for the signed-out account page, and shrink
the per-character hand-off line while the Characters section explains the export process
once. `CurrentCharacterBar` gains a `compact` prop so `/account` shows the pointer chip
(when one exists) instead of its body-prose empty sentence.

**Tech Stack:** Astro 5, Svelte 5 (runes: `$props`, `$state`, `$derived`, `$effect`,
`Snippet`/`{@render}`), Tailwind v4 utility classes, Vitest (`svelte/server` `render()`),
Playwright.

**Spec:** `docs/superpowers/specs/2026-09-22-account-and-island-states-design.md` (section 1
binds, section 2 is this lane's scope).

## Global Constraints

- Vocabulary: "addon" (in-game) vs "companion" (desktop app) — never blurred.
- Voice: reference, not pitch. Every visible string lives in a copy module.
- Motion budget: a skeleton shimmer while loading, one 160ms `.reveal` fade when data
  lands, nothing else. No Svelte `transition:`/`animate:` directives.
- Nothing moves: every island reserves its ready height while loading (`Skeleton` with
  `lines`/`minHeight`, or a fixed `min-h`). `.reveal` goes on the ready wrapper only.
  Lighthouse CLS budget is 0.05 on every bucket (`web/lighthouserc.json`); run `npm run
  lhci` before the final report and quote CLS per URL.
- Three states, three components: loading = `components/ui/Skeleton.svelte`, failed =
  `components/ui/LoadError.svelte` with an `onRetry` that re-fires the same request in
  place, empty = `components/ui/EmptyState.svelte` with at most one action.
- Lane ownership (this lane, `states-account`): `web/src/components/Account.svelte`,
  `web/src/components/account/**`, `web/src/components/CurrentCharacterBar.svelte`,
  `web/src/components/CurrentCharacterChip.svelte`, `web/src/components/SignInPrompt.svelte`,
  `web/src/components/CharacterHandoffLinks.svelte`, `web/src/lib/account/**`,
  `web/src/lib/handoff-copy.ts`, `web/src/lib/current-character-copy.ts`,
  `web/src/pages/login.astro`, `web/src/pages/account.astro`, e2e specs `auth.spec.ts`,
  `account*.spec.ts`, `handoffs-account-character.spec.ts`. Never touch `Header.astro`,
  `Footer.astro`, `components/ui/**`, `lib/ui/**` (coordinator's).
- Worktree: `/Users/jh/code/forever/.worktrees/states-account`, branch `states-account`.
  `web/node_modules` is a symlink — never `npm install`, never stage it.
- Every subagent runs on `sonnet`; `haiku` only for a purely mechanical single-file fix;
  never `opus`. One implementer at a time.
- Toolchain, from `web/`: `export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12`
  (already run for this lane; `FOREVER_DATA=fixture npm run sync` already run once).
  Scoped checks before every commit: `npx vitest run <paths>`, `npx astro check`,
  `npm run lint`, `npx prettier --check <paths>`. E2E: `E2E_PORT=4431 npx playwright test
  <spec files>` (`npx astro preview stop` first if a stale preview is running).
- Commits: message to a file under `.superpowers/` via `printf`, then `git commit -F
  <file>` alone in its own Bash call. Conventional subjects. End every message with
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` then `Claude-Session:
  https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`.
- Functions <50 lines, files <800 (already-large files: extract new copy into a new small
  module rather than growing them), no magic numbers, no dead code, immutable updates.
- Never save a public build or sim, never upload, never sign in on production. Tests use
  fixtures and route stubs.

---

## File Structure

- `web/src/lib/handoff-copy.ts` — shrink `needsExportLead`/`needsExportPasteLink`, drop
  `needsExportTail`; `NEEDS_EXPORT_TEXT` recomputed.
- `web/src/components/CharacterHandoffLinks.svelte` — one-line needs-export state.
- `web/src/lib/account/layout.ts` (new) — the reserved-height constants for `/account`'s
  loading skeletons and the signed-out/signed-in first-screen match.
- `web/src/components/account/AccountPanel.svelte` (new) — Svelte twin of the Astro-only
  `Panel.astro` (bordered/rounded/raised box), for the two grouped panels.
- `web/src/components/account/AccountPanel.test.ts` (new).
- `web/src/lib/account/character-list-copy.ts` — add the Characters section's intro line
  (once-only export explanation) and keep `empty`/`refreshFromBattlenet`.
- `web/src/components/account/CharacterList.svelte` — intro line, `EmptyState` swap
  (one action), `.reveal` on the ready list.
- `web/src/components/account/CharacterList.test.ts` — updated for the new empty state.
- `web/src/components/CurrentCharacterBar.svelte` — `compact` prop.
- `web/src/components/CurrentCharacterBar.test.ts` (new).
- `web/src/lib/account/signin-copy.ts` (new) — the one reason line shared by `/login` and
  `/account`'s signed-out state.
- `web/src/pages/login.astro` — reads the reason line from `signin-copy.ts`.
- `web/src/components/Account.svelte` — the account-mode reorder/regroup, loading
  skeletons, failed `LoadError`, signed-out `SignInPrompt` reuse, refreshed banner
  styling, email placeholder, compact chip, `.reveal`.
- `web/tests/e2e/account-load-error.spec.ts` (new) — failing `/v1/me`, retry, success.
- `web/lighthouserc.json` — add `http://localhost/account.html` to the URL list.

## Interfaces (cross-task contract)

- `layout.ts` exports `CHARACTERS_SKELETON_MIN_H: string`, `IDENTITY_SKELETON_MIN_H:
  string`, `MORE_SKELETON_MIN_H: string`, `SIGNED_OUT_MIN_H: string` — each a Tailwind
  `min-h-[...]` class string.
- `AccountPanel.svelte` props: `{ testid?: string; children: Snippet }`.
- `character-list-copy.ts` adds: `introLead: string`, `introPasteLink: string`,
  `introTail: string`.
- `CurrentCharacterBar.svelte` props: `{ hasOwnPasteBox?: boolean; compact?: boolean }`
  (both default `false`).
- `signin-copy.ts` exports `accountSignInCopy: { reason: string }`.

---

### Task 1: Shrink the per-character hand-off line (F5)

**Files:**
- Modify: `web/src/lib/handoff-copy.ts`
- Modify: `web/src/components/CharacterHandoffLinks.svelte:43-49`
- Test: `web/tests/e2e/handoffs-account-character.spec.ts` (already asserts
  `NEEDS_EXPORT_TEXT` — no edit needed, just re-run to confirm)

**Interfaces:**
- Produces: `handoffCopy.needsExportLead`, `handoffCopy.needsExportPasteLink`,
  `handoffCopy.pasteHref`, `NEEDS_EXPORT_TEXT` (no more `needsExportTail`).

- [ ] **Step 1: Rewrite the copy module**

```typescript
// web/src/lib/handoff-copy.ts
// Every visible string CharacterHandoffLinks.svelte uses. Two products are named here and
// they are never blurred: the addon runs in the game client; the companion is the desktop
// app that pairs with an account and sends what the addon wrote.

export const handoffCopy = {
  openInSimulator: 'Open in simulator',
  openInPlanner: 'Open in planner',
  /** Shown when the site holds no export for the character. Spec 2026-09-22 §2.3: the "how"
   *  (addon vs. companion) is explained once, in the Characters section's own intro line
   *  (character-list-copy.ts) -- this per-row line only points at the one action. */
  needsExportLead: 'No export yet ·',
  needsExportPasteLink: 'paste it here',
  pasteHref: '/addon#paste',
} as const;

/** The plain-text form of the needs-export line, for tests and accessible names. */
export const NEEDS_EXPORT_TEXT = `${handoffCopy.needsExportLead} ${handoffCopy.needsExportPasteLink}`;
```

- [ ] **Step 2: Update the component markup**

Replace lines 43-49 of `CharacterHandoffLinks.svelte`:

```svelte
  {:else}
    <span class="text-muted text-[13px]" data-testid="character-needs-addon">
      {handoffCopy.needsExportLead}
      <a class="text-text underline" href={handoffCopy.pasteHref}>{handoffCopy.needsExportPasteLink}</a>
    </span>
  {/if}
```

- [ ] **Step 3: Verify**

Run: `npx vitest run src/components/account/CharacterList.test.ts` (CharacterList renders
`CharacterHandoffLinks`; no crash) — full e2e re-run happens in Task 6/8, but confirm no
TypeScript error now: `npx astro check`.
Expected: no errors; `NEEDS_EXPORT_TEXT` is `"No export yet · paste it here"`.

- [ ] **Step 4: Commit**

```bash
git add web/src/lib/handoff-copy.ts web/src/components/CharacterHandoffLinks.svelte
printf 'fix(web): the per-character hand-off line shrinks to one clause\n\nThe how (addon vs. companion) is explained once in the Characters section intro (spec 2026-09-22 F5); each row now only points at the paste link.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-1.txt
git commit -F .superpowers/commit-msg-1.txt
```

---

### Task 2: `account/layout.ts` and `AccountPanel.svelte`

**Files:**
- Create: `web/src/lib/account/layout.ts`
- Create: `web/src/components/account/AccountPanel.svelte`
- Test: `web/src/components/account/AccountPanel.test.ts`

**Interfaces:**
- Produces: the five layout constants and `AccountPanel` listed above. Task 5 consumes
  both; Task 3 consumes `CHARACTERS_SKELETON_MIN_H`... (actually consumed directly by
  Account.svelte in Task 5, not by CharacterList — CharacterList has no loading state of
  its own).
- Note: the pixel values below are a first-pass estimate built from the same atomic units
  `current-character-layout.ts` already uses (44px rows via `min-h-11`, `gap-3`=12px,
  `gap-8`=32px, ~18px per `text-[13px]` line). Task 7 measures the real built page and
  corrects these constants before the final `lhci` run — do not treat this pass as final.

- [ ] **Step 1: Write `layout.ts`**

```typescript
// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading skeletons and its signed-out/signed-in
// first-screen match (spec 2026-09-22 §2.5, §2.6). First-pass values built from the same
// atomic units current-character-layout.ts uses (44px rows, 12px/32px gaps); Task 7 of
// docs/superpowers/plans/2026-09-22-states-account.md measures the built page and corrects
// them so the skeleton's height matches its ready view exactly (Lighthouse CLS budget,
// web/lighthouserc.json, is 0.05).

/** Characters section: heading + intro line + 3 placeholder rows at 360px. */
export const CHARACTERS_SKELETON_MIN_H = 'min-h-[360px]';

/** Devices + You panel: pairing block, device row, and the identity/sign-out block. */
export const IDENTITY_SKELETON_MIN_H = 'min-h-[280px]';

/** Guilds + Billing + Reports panel: the tallest panel once a guild row is present. */
export const MORE_SKELETON_MIN_H = 'min-h-[220px]';

/** The signed-out prompt's floor, matched to the signed-in first screen (title + toast
 *  slot + Characters) so the footer does not move when the session resolves either way. */
export const SIGNED_OUT_MIN_H = 'min-h-[420px]';
```

- [ ] **Step 2: Write `AccountPanel.svelte`**

```svelte
<!-- web/src/components/account/AccountPanel.svelte -->
<!-- Svelte twin of Panel.astro (bg-raised border border-line rounded-panel), for /account's
     two grouped panels (spec 2026-09-22 §2.1): Devices+You, and Guilds+Billing+Reports.
     Panel.astro is Astro-only (no <script>), so it cannot mount inside a Svelte island, and
     this lane does not own components/ui/ to add a shared one there -- this repeats the box
     treatment instead. Sub-sections inside are the caller's own <section> elements; the
     caller adds `border-line-soft border-t pt-4` to every one after the first for the
     divider the spec asks for. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let { testid, children }: { testid?: string; children: Snippet } = $props();
</script>

<div class="bg-raised border-line rounded-panel flex flex-col gap-4 border px-4 py-4 md:px-6 md:py-5" data-testid={testid}>
  {@render children()}
</div>
```

- [ ] **Step 3: Write the failing test, then the passing one (already passing since the
  component above exists — write test first per TDD, confirm it fails against a stub, then
  keep the real component)**

```typescript
// web/src/components/account/AccountPanel.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { createRawSnippet } from 'svelte';
import AccountPanel from './AccountPanel.svelte';

describe('AccountPanel', () => {
  it('wraps its children in the bordered, raised box and carries a testid', () => {
    const children = createRawSnippet(() => ({
      render: () => '<p data-testid="inner">Inside</p>',
    }));
    const { body } = render(AccountPanel, { props: { testid: 'account-more', children } });
    expect(body).toContain('data-testid="account-more"');
    expect(body).toContain('bg-raised');
    expect(body).toContain('data-testid="inner"');
  });
});
```

- [ ] **Step 4: Run the test**

Run: `npx vitest run src/components/account/AccountPanel.test.ts`
Expected: PASS (2 assertions-worth, 1 test).

- [ ] **Step 5: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/lib/account/layout.ts src/components/account/AccountPanel.svelte src/components/account/AccountPanel.test.ts
git add web/src/lib/account/layout.ts web/src/components/account/AccountPanel.svelte web/src/components/account/AccountPanel.test.ts
printf 'feat(web): account/layout.ts and AccountPanel, the grouped-section box for /account\n\nSpec 2026-09-22 §2.1/§2.6: reserved-height constants for the loading skeletons, and a Svelte twin of Panel.astro since that one is Astro-only and components/ui/ is not this lanes to extend.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-2.txt
git commit -F .superpowers/commit-msg-2.txt
```

---

### Task 3: `CharacterList.svelte` — intro line, `EmptyState`, `.reveal`

**Files:**
- Modify: `web/src/lib/account/character-list-copy.ts`
- Modify: `web/src/components/account/CharacterList.svelte`
- Modify: `web/src/components/account/CharacterList.test.ts`

**Interfaces:**
- Consumes: `components/ui/EmptyState.svelte` (`message`, `action: {label, href}`,
  `testid`), already on `main`.
- Produces: `characterListCopy.introLead`, `characterListCopy.introPasteLink`,
  `characterListCopy.introTail` for Task 5 (Account.svelte does not use these directly,
  but they must exist and read correctly as a sentence).

- [ ] **Step 1: Add the intro line to the copy module**

Add to `web/src/lib/account/character-list-copy.ts` (inside the existing `characterListCopy`
object, do not remove any existing key):

```typescript
  /** Spec 2026-09-22 §2.3: the export "how" explained once, here, instead of on every row
   *  (handoff-copy.ts's needsExportLead now only points at the paste link). */
  introLead: 'A character becomes simmable once the site has its export: type /fs export in game and',
  introPasteLink: 'paste it here',
  introTail: ', or run the desktop companion and it sends the export for you.',
```

- [ ] **Step 2: Update the failing test first**

Replace the "shows the empty state with both actions" test in `CharacterList.test.ts`:

```typescript
  it('shows the intro line once, with the paste link, above the list', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain(characterListCopy.introLead);
    expect(body).toContain('href="/addon#paste"');
  });

  it('shows EmptyState with one action (Refresh from Battle.net) when there are no characters', () => {
    const { body } = render(CharacterList, { props: { characters: [] } });
    expect(body).toContain(characterListCopy.empty);
    expect(body).toContain('data-testid="account-characters-empty"');
    expect(body).toContain(characterListCopy.refreshFromBattlenet);
    expect(body).not.toContain('data-testid="characters-paste"');
  });
```

- [ ] **Step 3: Run to see it fail**

Run: `npx vitest run src/components/account/CharacterList.test.ts`
Expected: FAIL — `introLead` undefined / `account-characters-empty` not found (the old
markup still renders `characters-refresh` + `characters-paste`).

- [ ] **Step 4: Update `CharacterList.svelte`**

Import `EmptyState` and replace the heading/empty blocks:

```svelte
  import EmptyState from '../ui/EmptyState.svelte';
```

```svelte
<section class="flex flex-col gap-3 reveal" data-testid="account-characters">
  <h2 class="section-title text-[18px]">{characterListCopy.heading}</h2>
  <p class="text-muted text-[13px]">
    {characterListCopy.introLead}
    <a class="text-text underline" href="/addon#paste">{characterListCopy.introPasteLink}</a
    >{characterListCopy.introTail}
  </p>

  {#if bnetImportedAt !== undefined}
    <p class="text-muted text-[13px]" data-testid="bnet-imported">
      {characterListCopy.importedFrom(relativeTime(new Date(bnetImportedAt)))}
      <a href={REFRESH_HREF}>{characterListCopy.refreshFromBattlenet}</a>
    </p>
  {/if}

  {#if characters.length === 0}
    <EmptyState
      message={characterListCopy.empty}
      action={{ label: characterListCopy.refreshFromBattlenet, href: REFRESH_HREF }}
      testid="account-characters-empty"
    />
  {:else}
    <ul class="flex flex-col">
      <!-- unchanged -->
    </ul>
  {/if}
</section>
```

(Everything inside the `{:else}` `<ul>` block is unchanged from the current file — do not
rewrite it, only the surrounding heading/empty markup.) Note `.reveal` moves onto this
section's own root: `CharacterList` is only ever mounted once `/v1/me` has already answered
(Account.svelte, Task 5, renders the Characters skeleton itself while loading and mounts
this component only in the ready branch), so `.reveal` here is exactly "the wrapper of the
ready state" the global constraint asks for.

- [ ] **Step 5: Run tests to see them pass**

Run: `npx vitest run src/components/account/CharacterList.test.ts`
Expected: PASS, all tests including the 3 unchanged ones (race/class/item-level, guild
line, unguilded-no-guild-line) plus the two updated ones and the unchanged
imported-from-line test.

- [ ] **Step 6: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/lib/account/character-list-copy.ts src/components/account/CharacterList.svelte src/components/account/CharacterList.test.ts
git add web/src/lib/account/character-list-copy.ts web/src/components/account/CharacterList.svelte web/src/components/account/CharacterList.test.ts
printf 'fix(web): Characters explains the export once, and its empty state uses EmptyState\n\nSpec 2026-09-22 F2/F5/§2.6: the intro line carries the addon-vs-companion explanation once; the empty state is the shared EmptyState primitive with one action instead of two ad-hoc links.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-3.txt
git commit -F .superpowers/commit-msg-3.txt
```

---

### Task 4: `CurrentCharacterBar.svelte` gets a `compact` prop (F2)

**Files:**
- Modify: `web/src/components/CurrentCharacterBar.svelte`
- Create: `web/src/components/CurrentCharacterBar.test.ts`

**Interfaces:**
- Produces: `compact?: boolean` prop, default `false`. When `true`, forwarded to
  `CurrentCharacterChip` as `hasOwnPasteBox={true}` regardless of the `hasOwnPasteBox` prop
  — `CurrentCharacterChip` already renders nothing when `current === null &&
  hasOwnPasteBox === true` (see its `{:else if !hasOwnPasteBox}` branch, unchanged, this
  lane owns that file too but no edit is needed there).
- Consumed by: Task 5 (`Account.svelte` passes `compact`).

- [ ] **Step 1: Write the failing test**

```typescript
// web/src/components/CurrentCharacterBar.test.ts
import { render } from 'svelte/server';
import { describe, expect, it, vi, beforeEach, afterEach } from 'vitest';
import CurrentCharacterBar from './CurrentCharacterBar.svelte';
import { currentCharacterCopy } from '../lib/current-character-copy';

// readCurrent() reads localStorage via `window`, which svelte/server's render() has no
// DOM for; CurrentCharacterBar's own $effect (where readCurrent runs) never fires during
// SSR (same reasoning Account.svelte's header comment gives for its own $effects), so the
// component always SSRs with current === null. That is exactly the case this test needs.

describe('CurrentCharacterBar', () => {
  it('shows the no-character line by default', () => {
    const { body } = render(CurrentCharacterBar, { props: {} });
    expect(body).toContain(currentCharacterCopy.noCharacterLine);
  });

  it('renders nothing when compact and there is no current character', () => {
    const { body } = render(CurrentCharacterBar, { props: { compact: true } });
    expect(body).not.toContain(currentCharacterCopy.noCharacterLine);
    expect(body).not.toContain('data-testid="current-character-chip"');
  });
});
```

- [ ] **Step 2: Run to see the second test fail**

Run: `npx vitest run src/components/CurrentCharacterBar.test.ts`
Expected: first test PASSes, second FAILs (the no-character line still renders — there is
no `compact` prop yet, so TypeScript itself may already flag the unknown prop; either way
it's red).

- [ ] **Step 3: Add the prop**

In `CurrentCharacterBar.svelte`, change the props line and the render:

```svelte
  let { hasOwnPasteBox = false, compact = false }: { hasOwnPasteBox?: boolean; compact?: boolean } =
    $props();
```

```svelte
<CurrentCharacterChip {current} hasOwnPasteBox={hasOwnPasteBox || compact} {guildLine} onforget={forget} />
```

Add a doc-comment line above the props (extending the file's existing header comment):
`compact` is set by Account.svelte's account mode (spec 2026-09-22 F2): the page already
has a full Characters list a few hundred px below, so the pointer chip beside the title
shows only when a pointer exists and renders nothing otherwise, never the "no character"
sentence.

- [ ] **Step 4: Run tests to see them pass**

Run: `npx vitest run src/components/CurrentCharacterBar.test.ts`
Expected: PASS (2/2).

- [ ] **Step 5: Scoped checks and commit**

```bash
npx astro check
npm run lint
npx prettier --check src/components/CurrentCharacterBar.svelte src/components/CurrentCharacterBar.test.ts
git add web/src/components/CurrentCharacterBar.svelte web/src/components/CurrentCharacterBar.test.ts
printf 'fix(web): CurrentCharacterBar gains a compact mode with no empty sentence\n\nSpec 2026-09-22 F2: /account renders the pointer as a chip beside the title, or nothing -- never the body-prose "No character loaded" line that read as an import failure next to a freshly imported Characters list.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-4.txt
git commit -F .superpowers/commit-msg-4.txt
```

---

### Task 5: `Account.svelte` reorder, panels, states, sign-in reuse (F1, F3, F4, F6, F7, F8)

This is the largest task: it lands spec §2.1 (order + panels), §2.2 (compact chip), §2.4
(refreshed banner), §2.5 (signed-out prompt + email placeholder), §2.6 (loading/failed
states). Read the full current file first: `web/src/components/Account.svelte` (451
lines) — only the `mode === 'account'` branch (currently lines 306-451) and its supporting
script state change; `nav`, `login` (except one placeholder attribute), `reports`, and
`pairing` modes are untouched.

**Files:**
- Modify: `web/src/components/Account.svelte`
- Create: `web/src/lib/account/signin-copy.ts`
- Modify: `web/src/pages/login.astro`

**Interfaces:**
- Consumes: `Skeleton` (`lines`, `minHeight`, `testid`), `LoadError` (`message`, `onRetry`,
  `testid`) from `components/ui/` (already on `main`); `CHARACTERS_SKELETON_MIN_H`,
  `IDENTITY_SKELETON_MIN_H`, `MORE_SKELETON_MIN_H`, `SIGNED_OUT_MIN_H` from
  `lib/account/layout.ts` (Task 2); `AccountPanel` (Task 2); `CurrentCharacterBar`'s
  `compact` prop (Task 4); `CharacterList` (Task 3, unchanged props).

- [ ] **Step 1: Write `signin-copy.ts`**

```typescript
// web/src/lib/account/signin-copy.ts
// The one "why sign in" reason line shared by /login (its own paragraph, outside the
// Account island) and /account's signed-out state (SignInPrompt's `line` prop) -- spec
// 2026-09-22 §2.5: the account page reuses the login page's own reason rather than
// inventing a second one.
export const accountSignInCopy = {
  reason:
    'An account is needed to upload logs, pair the companion and claim characters. Reading reports and rankings never needs one.',
} as const;
```

- [ ] **Step 2: Update `login.astro` to read it**

```astro
---
// web/src/pages/login.astro
import Base from '../layouts/Base.astro';
import Account from '../components/Account.svelte';
import { accountSignInCopy } from '../lib/account/signin-copy';
---

<Base
  title="Sign in"
  description="Sign in to Forever Sixty with Battle.net or an email link to upload combat logs and claim your characters."
  path="/login"
  session
>
  <main
    id="main"
    tabindex="-1"
    class="mx-auto flex w-full max-w-[640px] flex-col gap-[22px] px-[18px] pt-4 pb-10 md:gap-8 md:px-12 md:pt-7"
  >
    <h1 class="section-title text-[18px]">Sign in</h1>
    <p class="text-muted text-[14px]">{accountSignInCopy.reason}</p>
    <Account client:load mode="login" next="/logs" />
  </main>
</Base>
```

- [ ] **Step 3: Add the imports and the email placeholder in `Account.svelte`**

At the top of the `<script>` block, add three imports (alongside the existing ones):

```typescript
  import Skeleton from './ui/Skeleton.svelte';
  import LoadError from './ui/LoadError.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import AccountPanel from './account/AccountPanel.svelte';
  import {
    CHARACTERS_SKELETON_MIN_H,
    IDENTITY_SKELETON_MIN_H,
    MORE_SKELETON_MIN_H,
    SIGNED_OUT_MIN_H,
  } from '../lib/account/layout.ts';
  import { accountSignInCopy } from '../lib/account/signin-copy';
```

(`ui/Skeleton.svelte` etc. resolve from `components/Account.svelte`'s own directory,
matching the relative path `./ui/...` used by `CharacterList.svelte`'s sibling import in
Task 3 — confirm the exact relative depth against the real file tree before committing;
`Account.svelte` lives in `components/`, `ui/` is `components/ui/`, so `./ui/Skeleton.svelte`
is correct.)

In the `mode === 'login'` branch, the email `<input>` (existing lines ~239-246) gains one
attribute:

```svelte
          <input
            id="account-email"
            class="border-line-warm bg-raised rounded-control text-text h-11 flex-1 px-3 text-[15px]"
            type="email"
            autocomplete="email"
            placeholder="you@example.com"
            required
            bind:value={email}
          />
```

- [ ] **Step 4: Replace the whole `{:else}` (account mode) block**

Delete the current block (existing lines 306-451, from `{:else}` through the matching
`{/if}` just before the file's final `{/if}`) and replace it with:

```svelte
{:else}
  <div class="flex flex-col gap-8" data-testid="account">
    {#if status === 'loading'}
      <Skeleton lines={3} minHeight={CHARACTERS_SKELETON_MIN_H} testid="account-characters-skeleton" />
      <Skeleton lines={4} minHeight={IDENTITY_SKELETON_MIN_H} testid="account-identity-skeleton" />
      <Skeleton lines={3} minHeight={MORE_SKELETON_MIN_H} testid="account-more-skeleton" />
    {:else if status === 'failed'}
      <LoadError message={error} onRetry={() => void load()} testid="account-load-error" />
    {:else if !signedIn}
      <div class={SIGNED_OUT_MIN_H}>
        <SignInPrompt line={accountSignInCopy.reason} next="/account" testid="account-signin" />
      </div>
    {:else}
      <div class="flex flex-col gap-8 reveal">
        <CurrentCharacterBar compact />

        {#if toast !== ''}
          <div class="bg-raised border-l-2 border-gold flex flex-col gap-1 px-4 py-3" data-testid="account-toast">
            <p class="text-[14px]">{toast}</p>
            {#if me!.bnet_imported_at !== undefined}
              <p class="text-muted text-[13px]">
                {characterListCopy.importedFrom(relativeTime(new Date(me!.bnet_imported_at)))}
              </p>
            {/if}
          </div>
        {/if}

        <CharacterList characters={me!.characters} bnetImportedAt={me!.bnet_imported_at} />

        <AccountPanel testid="account-identity">
          <section class="flex flex-col gap-3">
            <h2 class="section-title text-[18px]">Devices</h2>
            {#if devices.length === 0}
              <p class="text-muted text-[14px]">No devices paired.</p>
            {:else}
              <ul class="flex flex-col">
                {#each devices as device (device.id)}
                  <li class="border-line-soft flex min-h-11 items-center justify-between gap-4 border-b py-2">
                    <span class="text-[14px]">
                      {device.name}
                      <span class="text-muted">· {device.platform}</span>
                    </span>
                    <button
                      class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                      onclick={() => onRevoke(device.id)}
                      disabled={busy}
                    >
                      Revoke
                    </button>
                  </li>
                {/each}
              </ul>
            {/if}
            {#if pairing === null}
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
                onclick={onPair}
                disabled={busy}
              >
                Pair a device
              </button>
            {:else}
              <p class="tabular text-strong font-mono text-[24px]" data-testid="pairing-code">{pairing.code}</p>
              <p class="text-muted text-[13px]">
                Type this into the companion within {Math.round(pairing.expires_in / 60)} minutes.
              </p>
            {/if}
          </section>

          <section class="border-line-soft flex flex-col gap-3 border-t pt-4">
            <h2 class="section-title text-[18px]">You</h2>
            <p class="text-[14px]">{displayName}</p>
            <button
              class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4"
              onclick={onSignOut}
              disabled={busy}
            >
              Sign out
            </button>
            <label class="flex min-h-11 items-center gap-3 text-[14px]">
              <input
                type="checkbox"
                checked={me!.user.anonymize}
                onchange={onAnonymize}
                disabled={busy}
                data-testid="anonymize"
              />
              Show a pseudonym instead of my character names
            </label>
            <p class="text-muted text-[13px]">
              Applies everywhere your characters appear, on reports and rankings alike. Reports themselves are
              never deleted or rewritten.
            </p>
          </section>
        </AccountPanel>

        <AccountPanel testid="account-more">
          {#if me!.guilds.length > 0}
            <section class="flex flex-col gap-3" data-testid="account-guilds">
              <h2 class="section-title text-[18px]">{guildConsentCopy.heading}</h2>
              <ul class="flex flex-col">
                {#each me!.guilds as guild (guild.id)}
                  <li
                    class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b py-2 text-[14px]"
                  >
                    <a href={guildHref(guild.region, guild.ruleset, guild.name)}>{guild.name}</a>
                    <select
                      class="border-line-warm bg-raised rounded-control text-text h-11 px-3 text-[13px] md:h-9"
                      value={guild.consent ?? 'gear'}
                      onchange={(event) => onConsentChange(guild.id, event)}
                      disabled={busy || guildBusy === guild.id}
                      data-testid="account-guild-consent"
                    >
                      <option value="roster">{guildConsentCopy.roster}</option>
                      <option value="gear">{guildConsentCopy.gear}</option>
                      <option value="gear_bags">{guildConsentCopy.gearBags}</option>
                    </select>
                    <button
                      class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
                      onclick={() => onLeaveGuild(guild.id)}
                      disabled={busy || guildBusy === guild.id}
                      data-testid="account-guild-leave"
                    >
                      {guildConsentCopy.leave}
                    </button>
                  </li>
                {/each}
              </ul>
            </section>
          {/if}

          <section class={me!.guilds.length > 0 ? 'border-line-soft flex flex-col gap-3 border-t pt-4' : 'flex flex-col gap-3'}>
            {#if billing === null}
              <p class="label text-muted">
                {billingBlockCopy.notSubscribed} <a class="text-text underline" href="/premium">{billingBlockCopy.seePlans}</a>.
              </p>
            {:else}
              <h2 class="section-title text-[18px]">Billing</h2>
              <p class="text-[14px]">
                {billing.plan} —
                {billing.cancel_at_period_end ? billingBlockCopy.ends : billingBlockCopy.renews}
                {billing.current_period_end ? new Date(billing.current_period_end).toLocaleDateString() : ''}
              </p>
              {#if billing.status === 'past_due'}
                <p class="text-strong text-[13px]" role="alert">{billingBlockCopy.pastDueBanner}</p>
              {/if}
              <button
                class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4"
                onclick={onManageBilling}
                disabled={busy}
              >
                {billingBlockCopy.manageBilling}
              </button>
            {/if}
          </section>

          <section class="border-line-soft border-t pt-4">
            <MyReports {signedIn} />
          </section>
        </AccountPanel>
      </div>
    {/if}
    <div class="min-h-[21px]">
      {#if status !== 'failed' && error !== ''}<p class="text-[14px]" role="alert" data-testid="account-error">{error}</p>{/if}
    </div>
  </div>
{/if}
```

Notes for the implementer:
- The bottom `min-h-[21px]` error line now guards `status !== 'failed'` because a failed
  `/v1/me` already shows its message via `LoadError` above — showing it twice would repeat
  the sentence. `run()`-triggered errors (pair/revoke/sign-out/anonymize/consent/leave
  failures) still land here exactly as before, since `run()` never sets `status`.
- `CurrentCharacterBar` no longer renders in the loading/failed/signed-out branches (F2's
  fix is specific to the signed-in view; the old code rendered it unconditionally at the
  very top, including while loading, which is also removed here — it never had a
  loading-aware behaviour of its own anyway, so nothing regresses).
- `guildConsentCopy`, `guildHref`, `updateConsent`, `leaveGuild`, `billingBlockCopy`,
  `openPortal`, `MyReports`, `SECONDARY_BUTTON_FIXED`, `characterListCopy`, `relativeTime`
  (needs a new import — see below) are all already imported or already used identically to
  the current file; only the JSX layout moved.
- Add one import the old file did not need at the top level (the toast block now formats a
  date itself): `import { relativeTime } from '../lib/dates';`

- [ ] **Step 5: Run the existing unit/vitest suite for this file's dependents**

Run: `npx vitest run src/components/account/CharacterList.test.ts src/components/CurrentCharacterBar.test.ts src/components/account/AccountPanel.test.ts`
Expected: PASS (Account.svelte itself has no dedicated vitest file — it is covered by e2e).

- [ ] **Step 6: `astro check`, lint, prettier**

```bash
npx astro check
npm run lint
npx prettier --check src/components/Account.svelte src/pages/login.astro src/lib/account/signin-copy.ts
```
Fix any type error before proceeding — a common one: `billing` is `$derived` and typed as
possibly `null`, so the `{:else}` branch of `{#if billing === null}` already narrows it to
non-null; if `astro check` disagrees, add `billing!` at each use inside that branch instead
of restructuring the derivation.

- [ ] **Step 7: Run the account/auth e2e specs already in this worktree**

```bash
npx astro preview stop || true
E2E_PORT=4431 npx playwright test tests/e2e/auth.spec.ts tests/e2e/account-billing.spec.ts tests/e2e/handoffs-account-character.spec.ts
```
Expected: all pass. If `auth.spec.ts`'s "every control on the account page clears 44px on
phone" fails because a new element (e.g. the `label`-utility Billing line, which is not a
button) got wrongly turned into a button, fix the markup — `label` per spec §2.1 is text,
never a control.

- [ ] **Step 8: Commit**

```bash
git add web/src/components/Account.svelte web/src/pages/login.astro web/src/lib/account/signin-copy.ts
printf 'fix(web): /account reorders to Characters first, groups Devices+You and Guilds+Billing+Reports, and gets loading/failed/signed-out states\n\nSpec 2026-09-22 F1/F3/F4/F6/F7/F8, §2.6: Characters leads (was buried under Account/Billing/Devices); Skeleton/LoadError replace the ad-hoc loading and error lines; the signed-out page reuses SignInPrompt instead of a bare link; the refreshed toast gets a gold border; Billing collapses to one label line when inactive; the email field gets a placeholder.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-5.txt
git commit -F .superpowers/commit-msg-5.txt
```

---

### Task 6: e2e — failing `/v1/me`, retry, success (spec §2.7)

**Files:**
- Create: `web/tests/e2e/account-load-error.spec.ts`

**Interfaces:**
- Consumes: `data-testid="account-load-error"` and `data-testid="account-load-error-retry"`
  (the `LoadError` primitive's own `{testid}-retry` convention, Task 5's
  `testid="account-load-error"`).

- [ ] **Step 1: Write the spec**

```typescript
// web/tests/e2e/account-load-error.spec.ts
import { expect, test } from '@playwright/test';

function fulfil(body: unknown, status = 200) {
  return { status, contentType: 'application/json', body: JSON.stringify(body) };
}

const ME_OK = {
  ok: true,
  data: {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [
      { key: 'us/hardcore/elyra-duskvale', region: 'us', ruleset: 'hardcore', name: 'Elyra Duskvale', class: 'Priest' },
    ],
    guilds: [],
  },
  error: null,
  request_id: 'r',
};

test('a failed /v1/me shows Try again, which re-fires the same request', async ({ page }) => {
  let call = 0;
  await page.route('**/v1/me', (route) => {
    call += 1;
    if (call === 1) return route.fulfill(fulfil({ ok: false, data: null, error: { message: 'oops' }, request_id: 'r' }, 500));
    return route.fulfill(fulfil(ME_OK));
  });
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })));

  await page.goto('/account');

  await expect(page.getByTestId('account-load-error')).toBeVisible();
  await page.getByTestId('account-load-error-retry').click();

  await expect(page.getByRole('link', { name: 'Elyra Duskvale' })).toBeVisible();
  await expect(page.getByTestId('account-load-error')).toHaveCount(0);
});
```

- [ ] **Step 2: Run it**

```bash
npx astro preview stop || true
E2E_PORT=4431 npx playwright test tests/e2e/account-load-error.spec.ts
```
Expected: PASS. If the retry button's testid differs, read `components/ui/LoadError.svelte`
again — it is `{testid}-retry` where `testid="account-load-error"`, so
`account-load-error-retry`, exactly as written above.

- [ ] **Step 3: Prettier and commit**

```bash
npx prettier --check tests/e2e/account-load-error.spec.ts
git add web/tests/e2e/account-load-error.spec.ts
printf 'test(web): a failed /v1/me on /account offers Try again and recovers in place\n\nSpec 2026-09-22 §2.7.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-6.txt
git commit -F .superpowers/commit-msg-6.txt
```

---

### Task 7: Measure the real page, correct `layout.ts`, add `/account.html` to `lhci`

**Files:**
- Modify: `web/src/lib/account/layout.ts` (correct the four constants)
- Modify: `web/lighthouserc.json` (add the URL)

**Interfaces:** none new — this task only tunes Task 2's constants against reality and
wires the Lighthouse budget spec §2.7 asks for.

- [ ] **Step 1: Build and start the preview server**

```bash
npx astro preview stop || true
npm run build
npm run preview -- --port 4431 &
sleep 2
```

- [ ] **Step 2: Measure the ready-state section heights with a throwaway script**

```bash
cat > /tmp/measure-account.mjs <<'EOF'
import { chromium } from 'playwright';

const ME = {
  ok: true,
  data: {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [
      { key: 'us/pvp/thoradin', region: 'us', ruleset: 'pvp', name: 'Thoradin', class: 'Warrior', realm: 'Whitemane', level: 60, faction: 'alliance', source: 'bnet', guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', rank_index: 1, verified: true } },
      { key: 'us/hardcore/elyra-duskvale', region: 'us', ruleset: 'hardcore', name: 'Elyra Duskvale', class: 'Priest' },
    ],
    guilds: [],
    bnet_imported_at: '2026-09-21T09:00:00Z',
  },
  error: null,
  request_id: 'r',
};

const browser = await chromium.launch();
const page = await browser.newPage({ viewport: { width: 360, height: 800 } });
await page.route('**/v1/me', (route) => route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ME) }));
await page.route('**/v1/devices', (route) => route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify({ ok: true, data: [], error: null, request_id: 'r' }) }));
await page.goto('http://localhost:4431/account');
await page.waitForSelector('[data-testid="account-characters"]');

for (const testid of ['account-characters', 'account-identity', 'account-more']) {
  const box = await page.getByTestId(testid).boundingBox();
  console.log(testid, box?.height);
}
await browser.close();
EOF
node /tmp/measure-account.mjs
```

- [ ] **Step 3: Write the measured constants**

For each of `CHARACTERS_SKELETON_MIN_H`, `IDENTITY_SKELETON_MIN_H`, `MORE_SKELETON_MIN_H`,
round the measured height **up** to the nearest 4px (Tailwind's spacing scale) and update
`web/src/lib/account/layout.ts`'s three constants to that exact `min-h-[Npx]` value, with a
comment recording the measured source: `/* measured against meBnetFixture at 360px,
2026-09-22 */`. For `SIGNED_OUT_MIN_H`, sum the measured `account-characters` height plus
the toast slot's reserved space (0, since no toast shows signed-out) plus any fixed gap
between title and Characters that the signed-in view has — read this off
`page.getByTestId('account').boundingBox()` at the point just before Characters starts, or
simplest: measure `account-characters`'s own top-relative-to-viewport offset via
`page.getByTestId('account-characters').boundingBox()` `y`, then `SIGNED_OUT_MIN_H` =
that `y` plus the Characters height (i.e., the whole first-screen block up to and including
Characters).

- [ ] **Step 4: Add `/account.html` to `lighthouserc.json`**

In `web/lighthouserc.json`, add `"http://localhost/account.html"` to the `"url"` array
(anywhere in the list; alphabetical-ish grouping already exists but is not strict — add it
after `"http://localhost/logs.html"`). It matches no exclusion in the default
`assertMatrix` bucket's `matchingUrlPattern`, so it is held to the default bucket's
budgets: performance/accessibility/seo ≥ 0.95, LCP ≤ 1700ms, TBT ≤ 100ms, CLS ≤ 0.05. The
static build's `account.html` SSRs the loading skeleton (the `$effect` that fetches
`/v1/me` never runs during the Astro build, only after hydration in the browser lhci
drives) — the skeleton's reserved height is therefore what CLS is actually measuring here,
which is exactly why Steps 1-3 must land first.

- [ ] **Step 5: Run `lhci` and confirm**

```bash
npm run lhci 2>&1 | tail -80
```
Expected: every URL's CLS ≤ 0.05, including `account.html`. If `account.html` fails CLS,
the skeleton height from Step 3 does not match what the static build's unhydrated markup
actually reserves before JS runs (the skeleton is present in the SSR'd HTML already, since
`status` starts `'loading'` — so this should not move at all unless the fetch's failure in
the sandboxed lhci run, which has no real API, causes a transition to the `'failed'`
`LoadError` state, which is a *shorter* element than the three skeletons and would itself
be a shift). If that happens, note it as a known lhci-environment limitation (no backend to
answer `/v1/me`) in the final report rather than fighting it further — record the CLS
number either way.

- [ ] **Step 6: Stop the preview server, scoped checks, commit**

```bash
npx astro preview stop
rm /tmp/measure-account.mjs
npx prettier --check src/lib/account/layout.ts lighthouserc.json
git add web/src/lib/account/layout.ts web/lighthouserc.json
printf 'fix(web): measured skeleton heights for /account, and its Lighthouse CLS budget\n\nSpec 2026-09-22 §2.6/§2.7: the three loading skeletons now reserve the exact height the ready panels render at 360px against the Battle.net fixture; account.html joins the default lhci bucket.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > .superpowers/commit-msg-7.txt
git commit -F .superpowers/commit-msg-7.txt
```

---

### Task 8: Whole-branch review, fix wave, final verification

- [ ] **Step 1:** Fresh reviewer subagent reads the full diff (`git diff main`) against
  spec section 2 (all of §2.1–§2.7) and section 1 (global constraints), F1–F8 from
  `.superpowers/ux-login/review.md`, and this plan. List every deviation as `Ruling:
  <decision> — <why> — <cost if wrong>` in `.superpowers/sdd/2026-09-22-states-account/progress.md`.
- [ ] **Step 2:** Fix wave for anything CRITICAL/HIGH the reviewer finds (fresh
  implementer per fix, sonnet).
- [ ] **Step 3:** Final full scoped checks:

```bash
cd web
npx vitest run src/components/account/CharacterList.test.ts src/components/account/AccountPanel.test.ts src/components/CurrentCharacterBar.test.ts
npx astro check
npm run lint
npx prettier --check src/components/Account.svelte src/components/account/*.svelte src/components/account/*.ts src/components/CurrentCharacterBar.svelte src/components/CurrentCharacterChip.svelte src/components/SignInPrompt.svelte src/components/CharacterHandoffLinks.svelte src/lib/account/*.ts src/lib/handoff-copy.ts src/lib/current-character-copy.ts src/pages/login.astro src/pages/account.astro
```

- [ ] **Step 4:** Full e2e suite (not just this lane's specs — per house rules, a failure
  elsewhere caused by this lane's changes is this lane's to fix):

```bash
npx astro preview stop || true
E2E_PORT=4431 npx playwright test
```

- [ ] **Step 5:** `npm run lhci`, quote CLS per URL in the final report.
- [ ] **Step 6:** Stop every server this lane started; confirm with
  `ps -axo command | grep "[a]stro.*states-account"` that none remains.
- [ ] **Step 7:** Write the final report per `.superpowers/journeys/lane-common-web.md`'s
  format.
