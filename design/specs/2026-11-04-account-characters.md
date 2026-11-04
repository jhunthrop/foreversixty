# /account Characters card, current-character card, Battle.net refresh (2026-11-04)

Owner verdict: "the character page here looks like trash" (2000 px, five level-1 Skyborne/Troll characters).
Build budget: about 90 minutes. Lands by 20:30 UTC. No open questions; every choice below is final.
Tokens come from design/DESIGN-SYSTEM.md. Reference surface: the in-game character-select list (crest, name, "Level 1 Troll Warrior" on one line, status at the right edge) and the Details! window row (identity left, one quiet figure right).

## 1. The question and the answer without scrolling

Question: "Which of my characters is set up, and what is the one thing to do next?"
Answer visible without scrolling at 360 px and 1280 px: the current-character card (name, one status sentence, one button) and the first three character rows. Day one: card says "No gear or talents yet", rows say "No gear yet". Populated: card shows "Open in simulator", rows show their build pill.

## 2. Files touched (nothing new except one small pure helper)

- web/src/components/character/CharacterRow.svelte: settingsRow layout (section 3).
- web/src/components/account/CharacterList.svelte: header aside, last-raid placement, refresh state (sections 3, 6).
- web/src/components/account/CharacterRowLink.svelte: drops both muted sentences; renders only a link or nothing (section 3).
- web/src/components/Account.svelte: hero card (section 4), `?refreshed=1` handling (section 6).
- web/src/lib/account/character-list-copy.ts, account-page-copy.ts, handoff-copy.ts: copy (sections 3, 4, 6).
- New web/src/lib/account/refresh-snapshot.ts: `saveSnapshot(me)`, `readSnapshot()`, `compareSnapshot(snapshot, me)`. Pure, unit tested.
- Reuse: CardHead, Tag, Spinner, button('secondary','sm'), CharacterIdentity, CharacterPortrait. Invent no control.

## 3. Characters card

### Layout of one row (the cause of the bug was one flex line holding identity, pills, a button and a sentence)

A row is a CSS grid, not a wrap-flex. Class on the `li`: `ROW ROW_HOVER` as today, with `grid grid-cols-[40px_minmax(0,1fr)_auto] items-center gap-x-[14px] gap-y-2 py-3`.

Desktop and tablet (md and up, 768 px+), three columns:
1. Crest, 40 px, existing `CharacterPortrait` (circular ringed HD crest). Never shrinks.
2. Identity column, `minmax(0,1fr)`. Line 1: name, 15 px semibold, class colour, `truncate` (one line, ellipsis, full name in `title`). Line 2: descriptor, 13 px, `text-muted`, `[text-wrap:balance]`, wraps only at the " · " separators. Each segment is `inline-block whitespace-nowrap`, so "High Order Skyborne" and "Warrior" never split inside a segment. Segments, in order: race, class, "Level N", ruleset. Example: `High Order Skyborne · Warrior · Level 1 · Forever`. If the guild line exists it is line 3, 13 px muted, truncated.
3. Actions column, `auto` width, `justify-self-end`, `flex items-center gap-2 whitespace-nowrap`. Left to right: status pill, then at most one button or link.

Phone (under 768 px), same grid with the actions column moved to a second grid row: `grid-cols-[40px_minmax(0,1fr)]`, actions `col-start-2 flex-wrap gap-2`, so they sit under the identity column, aligned with the name, not under the crest. Identity always owns the full width beside the crest. Minimum touch target 44 px on every action (`min-h-11`, `md:min-h-0`, as today).

### Row contents (one primary element per row: the name)

Pill (one, `Tag`, 12 px, never wraps):
- Current character: `Tag` gold, "Current". Replaces nothing else; it sits left of the status pill.
- Main: `Tag` muted, "Main".
- No build: muted 12 px text, "No gear yet" (replaces "No build yet" and both per-row sentences). Not a bordered pill; it is quiet text.
- Has build: the existing source pill, "Battle.net · 2 days ago" / "Addon · today".

Pill order in the actions cell: Current, Main, then build status. Maximum two shown on phone: if Current and Main and status all apply, the status goes on the identity column's line 3 on phone only.

Action (exactly one per row, secondary, 13 px semibold):
- Row with a build: link "Open in simulator" (existing `LINK` class).
- Row without a build, not the main: `button('secondary','sm')` "Set as main".
- Row without a build that is the main: nothing. The pill says Main; the hero card holds "Paste an export".
- Row with a build, not the main: both "Set as main" button and "Open in simulator". The link is a text link, the button secondary; on phone they share the actions line.
`CharacterRowLink` no longer renders `battlenetPendingLead`, `needsExportLead` or any paste link. Delete those strings from handoff-copy.ts (`battlenetPendingLead`, `battlenetPendingPasteLink`, `needsExportLead`, `needsExportPasteLink`, `NEEDS_EXPORT_TEXT`) and fix the tests that read them. `pasteHref` stays.

### States of row controls

- Row hover: `bg-row-hover`, 120 ms (existing). Name is a link: underline on hover, gold focus ring `outline-2 outline-gold` with `-outline-offset-2` on keyboard focus-visible.
- "Set as main": default secondary; hover per `button`; focus-visible gold ring; active pressed per `button`; disabled while any set-main call runs (`aria-busy` on the pressed one, label unchanged, 60% opacity); error: one `role="alert"` line under the list, text `Could not set your main just now.` (existing).
- "Open in simulator": link states per global link style; no disabled state.
- Touch and keyboard parity: Tab order is name, Set as main, Open in simulator, per row. Every action is reachable by keyboard and is at least 44 px tall on touch.

### Card header

`CardHead` glyph user, title "Characters". Aside (right-aligned on md+, wraps under the title on phone, 13 px muted):
- Idle with import time: `Imported from Battle.net 2 hours ago.` then the link `Refresh from Battle.net` (section 6 for its states).
- Empty list, Battle.net: link `Refresh from Battle.net` only.

### Day-one status (said once)

No strip in the card. The one explanation lives in the current-character card (section 4). The list says only "No gear yet" per row.

### Last raid line

Move `CoachLastRaidLine` out of the top of the card to the bottom of the rows, directly above the footer help paragraph. It keeps its own `border-t`, `px-5`, reserved height (77 px phone, 57 px desktop) and its existing copy. Rationale: it is a coach hint about raids, not about characters, and above the rows it read as a stray line. It stays inside the card, so nothing else moves. It still renders only when there are characters.

### Footer help paragraph

Keep, one paragraph, 13 px muted, with the existing links. Copy changes:
`Your characters come from Battle.net each time you sign in. Install the addon for bags, bank and live sync. In Korea, or Battle.net is down? Paste an export instead.`
(Assembled from the existing `intro*` strings; no change to the strings.)

### Long names

Name: one line, `truncate`, `title` shows full text. "Windshaper Skyborne" fits at 390 (identity column is 390 - 40 padding - 40 crest - 14 gap = about 296 px). Descriptor wraps at separators into at most two lines on phone: `High Order Skyborne · Warrior` / `Level 1 · Forever`. Never a single-word line, because segments do not break inside.

## 4. Current-character card (the hero)

Layout unchanged (CharacterIdentity xxl, label "Current character"). Only the `hero.build === undefined` block changes. Remove the `noBattlenetData` paragraph (`account-hero-no-bnet-data`) from the hero entirely; delete `accountPageCopy.noBattlenetDataForRealm` and its condition. Replace `noExportPrompt`.

Block: dashed-border box as today, `flex flex-wrap items-center justify-between gap-3`, two lines of copy in the left cell, one button right:
- Line 1, 14 px `text-strong` semibold: `No gear or talents yet.`
- Line 2, 13 px `text-muted`: `Battle.net sends them after you next log out of the game. To see them now, paste an export from the addon.`
- Button, `button('primary')`: `Paste an export` (href `/setup#paste`, existing testid `account-hero-paste`).

Phone: button goes under the text, full width of the box (`w-full` under 640 px), 44 px tall. Desktop: right-aligned, auto width.
Row with a build: unchanged (Open in simulator, Open in planner, Logs).
At 2000 px the hero card and list column keep the existing 1fr + 440 px rail grid; the left column is capped at `max-w-[1100px]` and centred with the page, as every container is.

## 5. Hierarchy

- Page level: the hero card is the primary element; its single primary button is the only gold-filled button on the page.
- Card level: the name is primary in each row; descriptor and status are secondary muted; actions are secondary buttons or text links.
- One emphasis per row; no coloured pills except Current (gold) and the source pill when a build exists.

## 6. Battle.net refresh

### Trigger A: "Refresh from Battle.net" (Characters header, OAuth redirect)

The link stays an anchor to `battlenetStartUrl('/account?refreshed=1')`. Add an `onclick` that, before navigation: saves the snapshot, sets busy, and lets the navigation proceed (do not `preventDefault`).

Snapshot, sessionStorage key `fs.refresh.snapshot`, JSON: `{ "at": <me.bnet_imported_at or null>, "chars": { "<character.key>": <level or 0> } }`. Read and deleted on return.

In-progress (until the page navigates, typically under 1 s, but specified because it is the click feedback):
- Link becomes `button('secondary','sm')`-looking text with `Spinner` (14 px) and the text `Refreshing…`. `aria-disabled="true"`, `pointer-events-none`, a second click is ignored in the handler (flag), so no double submit.
- Rows get `opacity-60` (120 ms transition) and a 2 px progress bar sits on the card's top edge: `absolute inset-x-0 top-0 h-0.5 bg-gold`, indeterminate, a 40% wide segment sliding 1.2 s linear infinite. Reduced motion: no sliding, the bar is static full width at 40% opacity.
- Live region (one, persistent, `role="status" aria-live="polite" class="sr-only"`) says: `Refreshing from Battle.net.`
- The card keeps its height; only opacity changes. Nothing shifts.

On return with `?refreshed=1` (Account.svelte; remove the old `toast` for this param; keep `replaceState` to drop the query):
Read snapshot, delete it, compare with the fresh `/v1/me`. Result line replaces the aside text, same slot, same 13 px, with a leading 14 px glyph, announced through the same live region:
- Error: snapshot `at` is not null and `me.bnet_imported_at` is missing or not later than `at`. Text, `text-wipe` colour: `Battle.net did not answer. Nothing changed.` plus a button `Try again` (secondary sm, same behaviour as the refresh link). Persists until the next click or reload.
- Changed: imported_at advanced and any key was added, removed or had a level change. Text, `text-text`: `Updated just now · 5 characters` (N is `me.characters.length`; singular `1 character`). If characters were added append ` · 2 new`.
- No change: imported_at advanced, keys and levels identical. Text: `No changes · 5 characters`.
- No snapshot (new tab, storage blocked) and `refreshed=1`: treat as Changed text without the ` · new` suffix.
- Success lines revert to `Imported from Battle.net just now.` after 8 s, plus the `Refresh from Battle.net` link. Errors never auto-revert.
- Rows that are new or whose level changed get the 2 s gold wash (`bg-gold/10` fading to transparent, `transition-colors duration-[2000ms]`). Reduced motion: no fade, no wash.

The `Refresh from Battle.net` link states: default gold link; hover underline; focus-visible gold outline 2 px; active as link; busy as above; disabled equals busy; error state replaced by `Try again`.

### Trigger B: "Check again" (home hero, in place)

Behaviour is already in HomeAccountPanel.svelte; the copy and states change to match.
- Default: `Check again`.
- Busy: `Spinner` 14 px plus `Checking…`, `disabled`, `aria-busy="true"`. The note element (`home-check-again-note`, already `role="status"`) is cleared. Minimum busy time 600 ms so the state is visible.
- Result, in the note, 13 px muted: found characters, the panel swaps to the character hero (no note); none: `Still nothing from Battle.net. Make a character in game, then check again.`; failure, `text-wipe`: `Could not reach Forever Sixty. Try again.` and the button label becomes `Try again`.
- Button width is fixed (`min-w-[116px]`) so label changes do not shift the row.

Reduced motion: `Spinner.svelte` already stops rotating (`motion-reduce:animate-none`) and shows the open ring; the label "Refreshing…" or "Checking…" carries the meaning. The progress bar is static (above).

## 7. Empty and error states of the card

- No characters, Battle.net: existing `characterListCopy.empty` with `Get set up` and `Paste an export` (unchanged).
- No characters, email: `characterListCopy.emptyByEmail` (unchanged).
- Loading: existing `Skeleton` rows, 3 rows at 69 px (crest 40 plus padding), fixed card min-height so hydration shifts nothing.

## 8. Copy (final strings)

character-list-copy.ts: `noBuildYet: 'No gear yet'`; add `refreshing: 'Refreshing…'`, `refreshingAnnounce: 'Refreshing from Battle.net.'`, `refreshError: 'Battle.net did not answer. Nothing changed.'`, `tryAgain: 'Try again'`, `updatedJustNow: (n) => 'Updated just now · ' + count(n)`, `noChanges: (n) => 'No changes · ' + count(n)`, `newSuffix: (k) => ' · ' + k + ' new'`, `importedJustNow: 'Imported from Battle.net just now.'`. `count(n)` is `1 character` or `N characters`. Delete `refreshedToast`.
account-page-copy.ts: `noExportPrompt` becomes `noExportTitle: 'No gear or talents yet.'` and `noExportBody: 'Battle.net sends them after you next log out of the game. To see them now, paste an export from the addon.'`. Delete `noBattlenetDataForRealm`.
home-no-character copy: `checking: 'Checking…'`, `stillNothing`, `checkFailed` as in section 6.

## 9. Responsive summary

- 2000 and 1440 px: row is three columns, actions right-aligned; descriptor on one line; no wrapping at any name length up to the column width. Card width is the left column (not stretched across the gutter); centred container.
- 1024 px: single column below the `lg` breakpoint; same row.
- 390 and 360 px: crest + identity, actions on a second line under the name, wrapped; hero button full width.
- Never collapses: crest, name, descriptor, one action. Collapses to a second line: pills and actions.

## 10. Performance and polish

No layout shift: result line occupies the same aside slot; the last-raid block keeps its reserved height; crests have fixed 40 px boxes with `width`/`height`; no icon fetch flash (existing crest sprites). No new dependency. Lighthouse budgets in web/lighthouserc.json unchanged; the progress bar is CSS only.

## 11. Acceptance screenshots (attach, all with the mocked /v1/me)

Viewports 2000, 1440, 390 for each:
1. Day-one: five level-1 characters, no builds, one current.
2. Populated: three builds (60/60/42), one in a guild, a main chosen.
3. Long names: "Windshaper Skyborne" and a 12-letter name at 390.
Plus at 1440 and 390:
4. Refresh busy (spinner, dimmed rows, progress bar).
5. Result lines: Updated, No changes, Error with Try again.
6. Row hover and keyboard focus on "Set as main".
7. Home "Check again" busy and "Still nothing" result at 390 and 1280.
Tests: unit for refresh-snapshot compare (changed, no change, error, no snapshot); component test that no row renders a sentence containing "Battle.net publishes"; a test that the descriptor segments render as separate nowrap spans.
