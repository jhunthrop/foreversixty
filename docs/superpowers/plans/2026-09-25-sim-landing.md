# Sim landing pass (owner-approved UX review, 2026-09-24) — plan

Lane: `sim-landing`, worktree `.worktrees/sim-landing`, branch `sim-landing`, e2e port 4384.

## Global constraints (from the lane dispatch)

- Files owned: `LandingState.svelte`, `SimView.svelte` (landing branch + message only),
  `sim/copy.ts` (or a new small `sim/landing-copy.ts`), `CharacterRow.svelte`,
  `CharacterIdentity.svelte` (one additive prop only), their tests, and e2e specs
  `sim-landing.spec.ts`, `sim-phone.spec.ts`, `sim-sources.spec.ts`,
  `auth.spec.ts`/`account*.spec.ts` where the row's markup change moves an assertion.
- Not owned: `HomeAccountPanel.svelte`, `lib/home/**`, `lib/account/home-hero.svelte.ts`,
  `index.astro`, `components/ui/**`, `lib/ui/**`, `lib/data/**`, `api/**`.
- No new fetches. Skeleton heights unchanged. Every visible string in a copy module.
  Reference voice. Lighthouse budgets hold on `/sim` (compare against a `main` build).

## Files touched outside the owned list (narrow, directed by Finding 2's own trace)

- `src/lib/sim/api.ts` — `fetchSimInput`'s 404 gets its own message instead of the saved-sim
  `simCopy.notFound`, via a `notFoundMessage` argument to the shared `call()` helper. No
  other call site passes one, so every other 404 (`fetchSim`, `dispatchServerSim`, …) is
  byte-for-byte unchanged.
- `src/components/sim/SettingsBar.svelte` — Finding 6 explicitly names this file as the new
  home for the engine-version link; the same anchor, same classes, same test id, added at
  the end of the section, right-aligned.

`sim/sources.ts` and `sim/store.svelte.ts` need **no** change: `fromStoredCharacter`'s catch
already forwards `error.message` verbatim, and the store's `message` getter already forwards
that unchanged — so api.ts's new 404 message flows straight through both with zero edits.

## Ruling log (recorded as each is made, mirrored in the ledger)

1. **404 detection without touching sources.ts/store.svelte.ts**: `fetchSimInput` gets a
   distinct `notFoundMessage` fallback (`landingCopy.buildMissingFallback`, a plain
   sentence with no name/links). `SimView.pickCharacter` compares `store.message` against
   that exact sentinel string once the pick settles, to learn whether *this* failure was the
   character-has-no-build 404 (vs. the race refusal, vs. anything else) — cheaper and no
   riskier than threading a `SimApiError.status` check through `fromStoredCharacter`, and it
   touches zero files outside api.ts. Cost if wrong: a 404 shows the old plain-text row
   instead of the new alert, or vice versa — cosmetic, caught by the e2e spec.
2. **Pill suppression is opt-in, not global**: `CharacterRow` gets a `hidePillWhenNoBuild`
   prop, default `false`. Only `LandingState` sets it. The account page's rows (not mine)
   keep today's muted "No build yet" text. Cost if wrong: account page pill disappears
   unexpectedly — caught by `CharacterList.test.ts` and `auth.spec.ts`, both run in full.
3. **`below` snippet, not a `class` prop**: `CharacterIdentity` gains one additive
   `below?: Snippet`, rendered in the text column after the descriptor line.
   `CharacterRow` passes its pill + guild line through it. Every other `CharacterIdentity`
   caller (`Account.svelte`, `HomeAccountPanel.svelte`) is unaffected — the prop is optional
   and they don't pass it.
4. **Guild line is not newly enabled on the landing rows.** Finding 3 repositions the pill
   and guild line *within* `CharacterRow`; it does not ask for a guild line landing did not
   already have. `LandingState`'s `<CharacterRow>` call does not pass `guildLine` (stays the
   existing default `false`) — no scope creep.
5. **New copy module**: `sim/copy.ts` is 1,286 lines already, over the 800-line ceiling, so
   every new string for this pass goes in a new `src/lib/sim/landing-copy.ts` per the
   dispatch's own suggestion.
6. **Engine-hash new home**: the settings bar (not the results card) — it's the first thing
   that renders once a character is loaded and is visible whether or not a run has happened
   yet, so the hash is visible in every state "the engine matters" without waiting for a
   result.

## Tasks

1. `sim/landing-copy.ts` (new): every new string — `pasteExport`/`pasteExportHref`, the
   build-missing alert's three text segments + two link labels/hrefs, the api-level 404
   fallback sentence.
2. `sim/build-pill.ts`: add `hasBuild(character)` — the one presence check the action branch
   and the pill visibility both make, so they can't read `character.build` two different
   ways. Small, additive, no behaviour change to `buildSourcePill`.
3. `sim/api.ts`: `call()` grows a `notFoundMessage?: string` 5th argument; `asSimError` uses
   it in place of `simCopy.notFound` when given. `fetchSimInput` passes
   `landingCopy.buildMissingFallback`. Every other call site unchanged (no 5th argument).
4. `CharacterIdentity.svelte`: additive `below?: Snippet`, rendered after the descriptor
   line, inside the text column. Test: renders content passed through it; every existing
   test untouched and still green.
5. `CharacterRow.svelte`: wrap `CharacterIdentity` (now the sole child of the flex-1/min-w-0
   wrapper) and pass pill + guild line through its `below` snippet instead of stacking them
   as CharacterRow's own siblings. New `hidePillWhenNoBuild` prop. Tests: pill still renders
   by default; suppressed only with the new prop and no build.
6. `LandingState.svelte`: `descriptor="full"` on the row; `hidePillWhenNoBuild`; the action
   snippet branches on `hasBuild(character)` — the existing Sim button, or a new
   `sim-paste-<key>` link to `/addon#paste`; the own `sim-landing-scope-note` paragraph is
   deleted; a new `failedKey` prop drives a `sim-landing-build-missing` alert, rendered
   between the list and the "Sim something else" link, with the character's name and two
   links. Tests: no-build row (paste link, no pill, no Sim button); build row unchanged;
   alert renders only for a matching `failedKey`, with both links; full descriptor renders.
7. `SimView.svelte` (landing branch + message only): delete the top, unconditional
   engine-version block (Finding 6) and its now-unused `ENGINE_VERSION`/`engineLabel`
   import; add `landingFailedKey` state, set from `pickCharacter`'s post-await sentinel
   check; pass `failedKey={landingFailedKey}` to `LandingState`; gate the existing
   `sim-landing-message` paragraph on `landingFailedKey === null` so the race hint never
   follows a 404.
8. `SettingsBar.svelte`: import `ENGINE_VERSION`/`engineLabel`; add the same anchor, same
   classes, same test id, right-aligned at the end of the section.
9. `sim-landing.spec.ts`: give `ME.characters` a `build` on both fixtures (today's tests
   assume a Sim button that the new gate would otherwise remove); add a no-build character
   + its `sim-input` 404 stub; new tests for the paste link and the build-missing alert
   (both links); update/remove the `sim-landing-scope-note` assertion.
10. `sim-phone.spec.ts`: one new test — a signed-in member with an unbuilt character clears
    44px on the `sim-paste-<key>` link (reuses the file's existing `noHorizontalScroll`/
    `targetsAreBigEnough` helpers).
11. `sim-sources.spec.ts`, `auth.spec.ts`: read only — confirmed neither assertion is
    position-dependent on the pill/guild-line move; re-run in full, no edit expected.
12. Whole-branch pass: `vitest run`, `astro check`, `lint`, `prettier --check` on touched
    files; full e2e suite once; `lhci` on `/sim` against a `main` baseline build under the
    same load.

## Method

Implemented directly (single lane controller), given the task's size and the amount of
shared context already loaded — not dispatched task-by-task to fresh subagents. Ledger at
`.superpowers/sdd/2026-09-25-sim-landing/progress.md` records each task and ruling as it
lands.
