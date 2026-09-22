# The account and sign-in experience, and one state model for every island

**Date:** 2026-09-22
**Status:** APPROVED by the owner ("do a ux review and improvement for the profile and login
experience"; "implement proper loading states and transitions across islands"). Evidence:
the UX review with screenshots at the coordinator's scratchpad `ux-login/review.md`
(findings F1 to F9 below are its numbering) and the island audit summarised in section 3.
Two lanes: `states-account` (section 2) and `states-islands` (section 3). Section 1 binds
both. Both build on the primitives already on `main` (section 1.2).

## 1. Global constraints (both lanes)

1. **Vocabulary.** The **addon** runs in the game client. The **companion** is the desktop
   app that pairs with an account and sends what the addon wrote. Never blur them; never
   write "the addon uploads" or "the companion in game".
2. **Voice.** `design/DESIGN-SYSTEM.md`: reference, not pitch. State the thing and stop.
   Every visible string lives in a copy module.
3. **Motion budget** (`design/DESIGN-SYSTEM.md`, Motion): a skeleton shimmer while loading,
   one 160 ms `.reveal` fade when data lands, nothing else. `prefers-reduced-motion` turns
   both off (already in `global.css`). No Svelte `transition:`/`animate:` directives.
4. **Nothing moves.** Every island reserves its ready height while loading: a `Skeleton`
   with `lines`/`minHeight` sized to the ready view, or a fixed `min-h` the island already
   has. The `.reveal` class goes on the wrapper of the ready state only. Lighthouse budgets
   in `web/lighthouserc.json` (CLS 0.05 on every bucket, TBT 100/200 ms) must hold; run
   `npm run lhci` before the final report and quote the CLS for every URL it measures.
5. **Three states, three components.** Loading = `components/ui/Skeleton.svelte`. Failed =
   `components/ui/LoadError.svelte` with an `onRetry` that re-fires the same request in
   place (never "reload the page"). Empty = `components/ui/EmptyState.svelte` with at most
   one action. Ad-hoc `<p>Loading.</p>`, `invisible` placeholders and bare alert text are
   replaced, not kept beside the new ones.
6. **Lane ownership.** `states-account` owns `web/src/components/Account.svelte`,
   `web/src/components/account/**`, `web/src/components/CurrentCharacterBar.svelte`,
   `web/src/components/CurrentCharacterChip.svelte`, `web/src/components/SignInPrompt.svelte`,
   `web/src/components/CharacterHandoffLinks.svelte`, `web/src/lib/account/**`,
   `web/src/lib/handoff-copy.ts`, `web/src/lib/current-character-copy.ts`,
   `web/src/pages/login.astro`, `web/src/pages/account.astro`, and the e2e specs
   `auth.spec.ts`, `account*.spec.ts`, `handoffs-account-character.spec.ts`.
   `states-islands` owns everything else under `web/src/components` and `web/src/lib` that
   section 3 names, plus their tests. Neither lane touches `Header.astro`, `Footer.astro`,
   `components/ui/**` or `lib/ui/**` (the coordinator's; ask by leaving a note in the final
   report if a primitive needs a prop).
7. Everything in `.superpowers/journeys/lane-common-web.md` binds (worktrees, sonnet-only
   subagents, commit rules, scoped checks, stop servers before the final report).

### 1.2 The primitives (on `main`, read them)

- `web/src/components/ui/Skeleton.svelte`: `lines`, `rowHeight` (Tailwind `h-*`),
  `minHeight` (Tailwind `min-h-*`), `label`, `testid`. `role="status" aria-busy="true"`.
- `web/src/components/ui/LoadError.svelte`: `message`, `onRetry?`, `testid`. `role="alert"`;
  the button is `{testid}-retry`.
- `web/src/components/ui/EmptyState.svelte`: `message`, `action?: {label, href}`, `testid`.
- `web/src/lib/ui/copy.ts`: `uiCopy.loading`, `uiCopy.retry`.
- `global.css`: `.reveal`.

## 2. Lane `states-account`: the account and sign-in experience

### 2.1 Account page information architecture (F3, F6, F8)

Order, phone-first, with the first screen showing the title and Characters:

1. Title. Directly under it, only when `?refreshed=1`: a confirmation banner (2.4).
2. **Characters** (`account/CharacterList.svelte`). Its intro line explains the export once
   (2.3). The `CurrentCharacterBar` does not render on this page (2.2).
3. **Devices**: the pairing block and the device list merged into one section, one heading.
4. **You**: Battle.net tag or email, the pseudonym toggle with its explanation, sign out.
   (Account and Name merge; both are how the player is identified.)
5. **Guilds**: only when the list is non-empty.
6. **Billing**: one line when there is no plan ("Not on Premium. See plans.") under a small
   `text-muted` label, a full section only when a plan exists.
7. **Your reports**.

Sections are grouped visually: Characters stands alone; Devices and You share a panel with
a `border-line-soft` divider; Guilds, Billing and Reports share a second panel. Use the
design system's panel treatment that `pages/logs.astro`'s `Panel` already renders; extract
it into `components/ui/Panel.svelte` only if it is not already a component (if it is an
Astro-only component, build a Svelte twin under `components/account/AccountPanel.svelte`).
Headings keep `section-title`; Billing's inactive line uses the `label` utility instead.

### 2.2 The current-character bar on the account page (F2)

`Account.svelte` in account mode no longer renders `CurrentCharacterBar`'s empty sentence.
When a current character exists it renders as the compact chip beside the title (the
`CurrentCharacterChip`), not as body prose. When none exists, nothing. The bar keeps its
behaviour on the planner and simulator pages, which it exists for.

### 2.3 Per-row hand-off (F5) and the explanation, once

`CharacterHandoffLinks.svelte`'s needs-export state shrinks to `No export yet · paste it
here` (link to `/addon#paste`), one line, `text-muted`. The section intro under the
Characters heading carries the explanation once, in `character-list-copy.ts`:
"A character becomes simmable once the site has its export: type /fs export in game and
paste it here, or run the desktop companion and it sends the export for you." with "paste
it here" linking to `/addon#paste`. `handoff-copy.ts` keeps `NEEDS_EXPORT_TEXT` in step.

### 2.4 The refresh confirmation (F4)

`?refreshed=1` renders a banner above Characters: gold left border (`border-l-2
border-gold`), `bg-raised`, the existing copy `characterListCopy.refreshedToast`, and the
import summary line when `bnet_imported_at` is set. It stays until the next navigation; the
URL is cleaned with `history.replaceState` as today.

### 2.5 Signed-out account page (F1) and login (F7)

Signed-out `/account` renders `SignInPrompt` with the login page's own reason line and its
Battle.net button, and a link to `/login` for the email route, above a `min-h` that matches
the signed-in first screen so switching states does not move the footer. The login page's
email field gets `placeholder="you@example.com"` and `autocomplete="email"`.

### 2.6 States on the account page (section 1.5 applied)

- Loading: one `Skeleton` per section sized to the section's typical ready height
  (Characters: 3 rows, `min-h` equal to the signed-in fixture's rendered height at 360 px,
  measured once and written as a constant in `account/layout.ts`), each with `testid`
  `account-<section>-skeleton`.
- Failed `/v1/me`: one `LoadError` at the top with `onRetry` calling the same fetch; the
  sections below render nothing.
- Empty Characters: `EmptyState` with the message that exists and the action "Refresh from
  Battle.net" (the paste link stays in the intro line).
- Ready: the wrapper gets `.reveal`.

### 2.7 Tests

SSR render tests for each state of `CharacterList`, the banner, the signed-out prompt;
Playwright: the existing `auth.spec.ts` journey plus one that stubs a failing `/v1/me`,
clicks Try again with the stub then answering, and sees Characters. `npm run lhci` CLS for
`/account.html` (it is not in the Lighthouse URL list; add it to the default bucket with
the page's fixture data available, or record why it cannot be).

## 3. Lane `states-islands`: one state model for every other island

The audit found: Character, Rankings, Guild, GuildJoin, GuildClaim, GuildSettings,
RecentReports, MyReports show an unreserved "Loading." line and jump; MyReports says
"Reload the page"; RecentReports, SimView's saved-sim error and the guild family have no
retry; the report's and simulator's lazily imported panels render nothing while their chunk
loads; `aria-busy` exists on one button; AddonPasteBox's signed-in block pops in.

### 3.1 Per island

| Island (file) | Loading | Failed | Empty | Reveal |
|---|---|---|---|---|
| `Character.svelte` | `Skeleton` matching the profile header plus the first table (measure the fixture at 360 px; constant in `lib/character-layout.ts`) | `LoadError` with retry | keep copy in `EmptyState` | yes |
| `Rankings.svelte` | `Skeleton` `lines` = the page size of rows, `rowHeight` = a row | `LoadError` with retry | `EmptyState` | yes |
| `Guild.svelte`, `GuildJoin.svelte`, `GuildClaim.svelte`, `GuildSettings.svelte` (via `GuildShell.svelte`) | one shared `Skeleton` sized per view, from a small `lib/guild/layout.ts` | `LoadError` with retry (the four copies of the same loading/error markup collapse into `GuildShell`'s one) | `EmptyState` where a list can be empty | yes |
| `RecentReports.svelte`, `MyReports.svelte` | `Skeleton` with as many rows as the list shows | `LoadError` with retry; "Reload the page" is deleted | `EmptyState`, MyReports's action "Upload a log" → `/logs`, and the copy names the companion as the other way | yes |
| `ReportView.svelte` lazy modes and views, `SimView.svelte`'s `SimResults`, `ToolsView.svelte` (already has skeletons) | the `lazyFallback` snippet renders a `Skeleton` sized to the panel it replaces (a `min-h` per mode in `lib/report/layout.ts`; the existing `TOOL_SKELETONS` stay) | keep the existing retry, rendered through `LoadError` | as today | yes on the loaded panel |
| `SimView.svelte` saved-sim error | as today | `LoadError` with retry (it has none) | as today | as today |
| `AddonPasteBox.svelte` | reserve the signed-in block's height (`min-h`) while `fetchMeOnce` is pending, so the hint never pushes the box | as today | as today | yes |
| `Planner.svelte` | keep its tuned `min-h`; replace the "Loading talent data" line with a `Skeleton` inside that reservation | route the existing Retry through `LoadError` | n/a | yes |

### 3.2 Busy controls

Every button that starts a request (Run, Save, Submit, Upload, Pair, Claim) sets
`aria-busy="true"` and `disabled` while the request is in flight and shows its label
unchanged (no spinner glyph, no label swap). A shared `lib/ui/busy.ts` exports the one
class string for the busy look (`opacity-60 cursor-progress`), used everywhere.

### 3.3 Tests

SSR render tests for each island's three states through the primitives' `testid`s;
existing Playwright specs updated where they waited for "Loading." text (use
`tests/e2e/support/hydrated.ts` and the skeleton `testid`s); `npm run lhci` before the final
report with every URL's CLS quoted.

## 4. Out of scope

Page transitions between routes (Astro view transitions), the phone nav's horizontal scroll
(F9, a separate decision), light mode, any new fetches.
