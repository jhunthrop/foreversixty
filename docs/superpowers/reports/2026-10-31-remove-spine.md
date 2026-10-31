# Remove the "Jump to" spine bar (2026-10-31)

Branch `wt/remove-spine`, from origin/main (532a4fad7). Not pushed.

The owner asked why the "Jump to" bar showed only on the sim pages, then said "yeah get rid of it". This lane removes the bar everywhere it was still mounted. Its two remaining jobs, the "Restored your last character" note and Forget, move into the header character selector.

## What changed

1. **Spine mode is gone.** `CurrentCharacterBar.svelte` is now the /setup chip wrapper and nothing else: it reads the pointer, adds the guild line and handles Forget.
   - Removed with it:
     - the door links, and `spineDoorsFor` / `resolveSpineClassSlug` (`lib/current-character-bar.ts` and its test are deleted);
     - `STRIP_HEIGHT`;
     - the `.chip-slot` pre-paint rule in global.css;
     - the spine copy keys (`barDoorsLabel`, `barSignedOutLine`, `barSignIn`, `barPasteExport`, `restoredNote`);
     - the `spine`, `compact`, `currentDoor`, `restored`, `profile`, `quietSignedOut`, `cardName` and `inCard` props.
   - Mounts removed:
     - SimView (the standalone strip, and the strip as the first row of the Your characters card, along with LandingState's `lead` snippet);
     - ToolsView;
     - **the character page** (`Character.svelte`). It mounted spine mode too, not the compact chip; see Concerns.
   - The planner did not mount it.
   - Dead code that went with it:
     - SimView's and ToolsView's `restored` state;
     - `followCurrentCharacter`'s `onSettled`;
     - `CurrentCharacterChip`'s `restored` prop;
     - `characterPrompt().quietBar`.
   - The shells no longer reserve the bar's band: `sim.astro`, `sim/[id].astro`, `sim/specs.astro`, and every tool skeleton in `bulk-skeleton.ts`.
   - The compact chip on /setup stays. Its `inCard` attribute did nothing in chip mode and is gone.
2. **"Restored your last character" and Forget are in the selector panel** (`SelectorPanel.svelte`, below the list):
   - a muted `.csel-note` row;
   - Forget sits at its end in the Manage link's quiet style, as a 44 px button with the accessible name `Forget {name}`;
   - choosing Forget closes the panel, returns focus to the trigger and announces `Forgot {name}.` in the selector's live region.

   The rule lives in `lib/character-selector/restored.ts`:
   - On the first page of a tab session, Base.astro's pre-paint script records which character the stored pointer named, in `sessionStorage['fs.sessionStartPointer']` as `source:ref` ('' for none).
   - The line shows while the pointer still names that character. It never shows for the account's main, because forgetting the main would change nothing.
   - The rule compares which character the pointer names, not its `savedAt`. The simulator rewrites the pointer every time it loads the restored character, and a timestamp check broke on that (the e2e caught it).
   - Forget runs `forgetPointer`: `clearCurrent` plus `CURRENT_CHARACTER_CHANGED`. The header falls back to the main and every island follows. Today's Forget dispatched no event, so other islands on the page stayed on the old character until reload.
3. **Reserves re-measured.** Method: zero the mount's `min-height`, measure the fixture state each comment names, then add 24 px.

   | Page | Phone (measured) | Desktop (measured) | Was |
   |---|---|---|---|
   | /sim | 1244 (1220 at 390) | 859 (834.5) | 1202/834 |
   | /sim/[id] | 1555 (1531 at 360) | 1135 (1110.5) | 1406/990 |
   | /sim/gear | 3469 (3444.5) | 2018 (1994) | 3167/2075 |
   | /sim/talents | 1949 (1924.5) | 1060 (1036) | 1703/1057 |
   | /sim/drops | 2856 (2832) | 1822 (1797.5) | 2288/1549 |
   | /sim/weights | 2137 (2112.5) | 1194 (1170) | 2175/1201 |

   - Most old figures had drifted below the hydrated height before this lane. The measured states are signed out with no pointer, where the bar never rendered.
   - Pages that lose the bar (signed in, or with a pointer) are now shorter than their reserves, which is safe for CLS.
   - Each figure's reason is in its file's comment. The planner reserve was not affected.

## Tests

- New `lib/character-selector/restored.test.ts` covers the identity rule, the main exclusion, a pasted export and Forget's clear-and-announce. There are also copy tests for the new keys.
- New `tests/e2e/selector-restored.spec.ts`:
  - no bar or "Jump to" on /sim, /sim/gear, /planner and a character page;
  - a restored character shows the line, and Forget falls back to the main and returns focus;
  - Forget works from the keyboard;
  - a character chosen this session shows no line;
  - a restored main shows no line;
  - signed out, a pasted export is forgotten and Sign in is left.
- Updated or removed specs and tests that asserted the bar:
  - deleted `spine-bar.spec.ts`;
  - current-character: Forget is now proven in a new tab;
  - sim-landing, one-character-journey, handoffs-setup, character-identity, character-score, character-history-panel, sim-saved, planner;
  - unit tests: `_sim.test.ts`, `bulk-skeleton.test.ts`, `current-character-layout.test.ts`, `CurrentCharacterBar.test.ts`, `CurrentCharacterChip.test.ts`, `character-prompt.test.ts`, `follow-current-character.test.ts`.

## Gates

Every build and run used `FOREVER_DATA=fixture PUBLIC_API_BASE_URL=http://127.0.0.1:9`, served on 4709.

- Full web vitest: 7757 passed, exit 0. A first run hit one unhandled `document is not defined` in `sim-tools-island.test.ts`, a boot timer firing after teardown. The file is untouched here, it passed in isolation three times, and the full rerun was clean.
- `npm run check`: 0 errors, 0 warnings. `lint`: exit 0. `format:check`: exit 0.
- Island-size check: passed in postbuild (sim-island 77,967 B gzipped, sim-tools-island 23,064 B).
- Full Playwright, desktop and mobile: 3368 passed, 408 skipped, 28 failed.
  - **22 failures come from the base URL alone.** They assert the real API origin in hrefs or request lists: account-discord-link, account-discord connect/connected, guild-discord-bot, guild-settings-discord and report-queries, all saying "every byte comes from this origin" while 127.0.0.1:9 calls are listed. They are not this lane's, and the hermetic-suite lane owns them.
  - **The other 6 were load flakes.** header-account sign-out failure, home-panel late shifts at 390 and 1440, header-finished-sim CLS at 1440, and account-discord private-logs toggle all passed on rerun (71/71) with the same build.

## Captures

All in `design/reviews/2026-10-31-remove-spine/`, signed in with a restored alt (Frostbyte) as the current character:
- `sim-1440/390.png`
- `sim-gear-1440/390.png` (the fixture warrior loaded by `?code=`)
- `planner-1440/390.png`
- `selector-restored-1440/390.png` (the panel open with the restored line and Forget; at 390 it is the sheet)

## Concerns

- **The character page lost the bar too.** The brief said that page uses the compact chip, but it mounted spine mode. For an owner it showed "Jump to · Plan · Sim · Logs · Rankings"; for a signed-out visitor it showed the sentence "Sign in with Battle.net or paste an export…" with two links, below the score card. Both are gone. The visitor still has the header's Sign in. If the owner wants the visitor's sign-in line back on profiles, it needs its own small component.
- Forget shows only while the character was restored from an earlier visit, as the brief specifies. A character picked in this session can be switched in the list but has no Forget, whereas today's Forget showed whenever any pointer existed.
- Forget now tells the other islands (`CURRENT_CHARACTER_CHANGED`) where today's did not. Every listener already handles a null pointer, because a Forget in another tab sends exactly that through the `storage` event.
- `CurrentCharacterChip`'s `addonCode` prop has had no caller since the planner rebuild. I left it alone as out of scope.
- The tool pages' reserves reflect signed-out fixture heights. A signed-in visitor whose character fails to load (e.g. the API down) sees a tall empty area above the footer. That was already true before this lane.
