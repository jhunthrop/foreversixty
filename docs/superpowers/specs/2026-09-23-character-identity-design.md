# One character identity component, and the simulator reads the session from the cache

**Date:** 2026-09-23
**Status:** APPROVED by the owner ("Should we not just have a single character display
component to use everywhere with all of the enhancements and efficiencies?" — "do it").
One lane: `character-identity`, `web/**` only.

## 0. Why

Four places draw a character's identity today, each its own way:

- `components/account/CharacterList.svelte`: avatar, else the class icon over the letter
  square, name in class colour, descriptor, build pill, guild line, Main pill or "Set as
  main" (the complete treatment).
- `components/HomeAccountPanel.svelte`: the hero (same avatar markup copied at 40 px) and
  the alt chips (copied again at 28 px).
- `components/sim/LandingState.svelte`: a bare colour swatch, the name, "PvP · US", the
  build pill. No avatar, no class icon (the owner's screenshot of /sim).
- `components/Character.svelte` (the public page header) and `components/Account.svelte`
  (the account hero band): the avatar or render, name, descriptor, hand-off links.

And the simulator island (`components/sim/SimView.svelte`) still calls the raw `fetchMe()`
on mount, not the cached session read every other island uses, so "Your characters" on
/sim waits for a live `/v1/me` round trip on every visit while the home and account pages
paint from the snapshot instantly.

## 1. Binding constraints

1. `design/DESIGN-SYSTEM.md` in full; the states model (spec
   `2026-09-22-account-and-island-states-design.md` §1): skeleton sized to the ready
   state, `LoadError` with retry, `EmptyState`, one `.reveal`, nothing else moves.
   Lighthouse budgets in `web/lighthouserc.json` hold; quote every URL's numbers before the
   final report (`/sim` and the home page in particular).
2. Vocabulary: addon in game, companion on the desktop, "from Battle.net" for
   Blizzard-sourced data. Every visible string in a copy module (`lib/account/
   character-list-copy.ts`, `lib/sim/copy.ts`, `lib/home-panel-copy.ts` exist; add a small
   `lib/character/identity-copy.ts` only for strings that are new).
3. The caching layer spec (`2026-09-23-caching-layer-design.md`) §3.3: islands render the
   session through `createQueryState` (`lib/data/query.svelte.ts`); never `fetchMe()` from
   a component, and never call `fetchMeOnce()`/`query()` from an effect that can re-run on
   the answer (the 2026-09-23 account-page loop; see the status memory).
4. Not touched: `Header.astro`, `Footer.astro`, `components/ui/**`, `lib/ui/**`,
   `lib/data/**`, anything under `api/`.
5. Lane rules: `.superpowers/journeys/lane-common-web.md` (sonnet only, never opus; commit
   rules; scoped checks; no broad `pkill`; stop servers; whole e2e suite once before the
   final report). Commit trailer for this lane: `Co-Authored-By: Claude Fable 5.1
   <noreply@anthropic.com>` then `Claude-Session:
   https://claude.ai/code/session_01DAHb7A9Ytd6qpZu2UujXgY` (this session, not the one the
   journey file names).

## 2. The components

### 2.1 `components/character/CharacterPortrait.svelte`

The one portrait. Props: `character: MeCharacter` (or the narrower `{name, class,
avatar_url?}` it needs; type it as the smallest shape and let `MeCharacter` satisfy it),
`size: 'sm' | 'md' | 'lg'` (28, 36, 44 px; a `SIZE_CLASS` map, no magic numbers inline),
`testid` prefix. Renders, in priority order: the avatar (`avatar_url`) → the class icon
(`classIconUrl`) over the letter square (`classSquare`, so a blank icon still shows the
letter) → the letter square alone. `loading="lazy"`, `alt=""`. Test ids `${testid}-avatar`,
`${testid}-avatar-fallback`, `${testid}-class-icon`, matching the ones `CharacterList` uses
today (`character-avatar`, `character-avatar-fallback`, `character-class-icon`) so the
existing account specs keep passing with `testid="character"`.

### 2.2 `components/character/CharacterIdentity.svelte`

Portrait + name + descriptor as one block. Props: `character`, `size` (drives the portrait
size and the name size: sm 14 px, md 15 px, lg 18 px display), `descriptor: 'full' |
'realm' | 'none'` (`characterDescriptor(character)` / `rulesetLabel(ruleset) · REGION` /
nothing), `href?` (the name becomes a link when given), `nameTestid?`,
`descriptorTestid?`. Name in `classColorVar(character.class)`, display font at `lg` only
(the account list and public header use the display face; chips and rows do not).

### 2.3 `components/character/CharacterRow.svelte`

The list row every character list renders: `CharacterIdentity` at `md` with the realm
descriptor by default, an optional build pill (`buildSourcePill`, existing), an optional
guild line, and a trailing `action` snippet (Svelte 5 `Snippet`) for whatever the list
needs (the Sim button, the Main pill / Set as main, nothing). `min-h-11`, the
`border-line-soft` divider, `flex-wrap` on phone, exactly the layout `CharacterList` has
today. The whole row is not a link; the name is (as today), and a list that wants the row
to act (the simulator) passes its own `onpick` through `href` plus an `onclick` on the
identity's link via a `onNameClick?` prop. Keep it to props the two callers use; nothing
speculative.

### 2.4 Callers, rewritten to use them

- `account/CharacterList.svelte`: each row is `CharacterRow` with the guild line, the build
  pill, and the Main / Set as main action. Behaviour, copy and test ids unchanged
  (`character-descriptor`, `character-build-pill`, `character-guild-line`,
  `character-guild-verified`, `character-main-pill`, `character-set-main`,
  `character-avatar*`, `character-class-icon`).
- `sim/LandingState.svelte`: each row is `CharacterRow` with the build pill and the Sim
  button as the action; the swatch is gone, the portrait is there. Test ids unchanged
  (`sim-character-<key>`, `sim-character-link-<key>`, `sim-character-build-<key>`,
  `sim-pick-<key>`, `sim-character-name`, `sim-character-descriptor` where the specs use
  them). The in-place `follow` behaviour (left click picks in place, modifier clicks open
  the link) stays.
- `HomeAccountPanel.svelte`: the hero strip is a compact band, so it uses
  `CharacterIdentity` at `md` with the full descriptor and the existing action links
  beside it; the alt chips use `CharacterPortrait` at `sm` plus the name. Test ids `home-hero-avatar`, `home-hero-avatar-fallback`,
  `home-character-chip` stay.
- `Account.svelte`'s hero band: `CharacterIdentity` at `lg` with the full descriptor and
  the character link, the render image beside it as today (`account-hero-render`).
- `Character.svelte`'s public header: `CharacterPortrait` at `lg` when there is no render
  (as today) and the name as the `h1`; leave the ranked-fights line as it is. Test ids
  `character-avatar`, `character-render` stay.

`lib/account/character-descriptor.ts` and `lib/account/build-pill.ts` stay the single
sources for the descriptor, square, icon URL and pill; the components import them.
Delete every copy of the avatar/fallback markup the components replace. No component
keeps a private variant.

## 3. The simulator reads the session from the cache

`SimView.svelte` replaces `onMount(() => fetchMe()...)` with
`createQueryState<Me | null>(\`${API_BASE_URL}/v1/me\`, () => fetchMeOnce(), {scope:
'private', ttlMs: 10 * 60 * 1000})` (the exact call `Account.svelte`, `SessionNav.svelte`
and `HomeAccountPanel.svelte` make). `me` becomes `$derived(session.data)`; the premium
flag (`store.setPremium(effectiveServerSims(me))`) and the history load run from an
`$effect` keyed on `me` that guards against re-running for the same object (compare the
reference, and never call a session read from inside it). While `session.status ===
'loading'` with no data, the landing area shows a `Skeleton` sized to four rows
(`lib/sim/layout.ts` constant, measured once); a failed session read renders as signed out
exactly as today (the sim never shows a session error). A returning signed-in visitor sees
"Your characters" from the snapshot before `/v1/me` answers: one Playwright test proves it
(same shape as `home-panel.spec.ts`'s returning-visitor test, reading the `fs.q.*` keys).

## 4. Tests and evidence

- SSR render tests (`svelte/server` `render`, as `CharacterList.test.ts` does) for
  `CharacterPortrait` (avatar / icon over square / square alone; three sizes),
  `CharacterIdentity` (three descriptor modes, link or not), `CharacterRow` (with and
  without pill, guild line, action).
- The existing `CharacterList.test.ts`, `HomeAccountPanel` tests and every e2e spec that
  touches these lists (`auth.spec.ts`, `account*.spec.ts`, `home-panel.spec.ts`,
  `sim-landing.spec.ts`, `sim-sources.spec.ts`, `sim-history.spec.ts`,
  `handoffs-account-character.spec.ts`, `character*.spec.ts`) pass; the new returning-
  visitor test for /sim.
- `npm run lhci` before the final report with every URL's numbers quoted; CLS and TBT must
  not move on `/sim` or the home page.
- The final report lists every file whose avatar markup was deleted, so the owner can see
  the four copies became one.
