# Forever Sixty site navigation and character selector: experience spec

Author: ux-designer. **Round 3 (2026-10-09): the four round-two deciders, section 14. Round 2 (2026-10-09): after the wow-player review `design/reviews/2026-10-09-nav-selector-player-review.md` (NOT YET) and the owner's instruction on the full name. Section 13 lists every change.** Mock: `design/mocks/gen_nav.py`; renders `design/mocks/renders/nav*.png`: `nav`, `nav-open`, `nav-open-many`, `nav-2000`, `nav-1439`, `nav-1280`, `nav-1024`, `nav-1023`, `nav-768`, `nav-phone`, `nav-phone-360`, `nav-phone-open`, `nav-phone-menu`, `nav-prehydration-phone`, `nav-signed-out`, `nav-signed-out-1024`, `nav-signed-out-phone`, `nav-signed-out-phone-open`, `nav-signed-out-pasted`, `nav-expired`, `nav-prehydration`, `nav-states` (the renders directory is gitignored; regenerate with `python3 design/mocks/gen_nav.py [name ...]`, run from anywhere, Playwright is found through `web/`). The current character in every render is "Obnoxious Yell", the owner's own, a two-word name. Every character in the mock is an example: invented names on real class and spec pairs.

References: `docs/tenets.md` (14 first: the front door is the player's character; 9 every state; 10 no invented controls; 13 no layout shift), `design/specs/2026-10-09-tier-list.md` (the 56 px phone bar, the Menu button, the wide-screen rule, the `panel` and `label` family), `web/src/components/Header.astro` (doors, 56 px phone bar, Discord), `web/src/components/AccountMenu.svelte` and `Account.svelte` mode `nav` (the session chip this replaces), `web/src/lib/current-character.ts` (the one pointer), `web/src/components/CurrentCharacterBar.svelte` (the spine bar's Switch popover, see open question 4), `web/src/lib/account/{api,character-descriptor,character-list-copy}.ts` (what `/v1/me` gives), `web/src/lib/nav.ts`.

Reference surface: the **in-game character frame and the Details! window title menu**. From the character frame: crest (portrait), name in class colour, "Level 60 Fury Warrior" under it. From Details!: the title bar control that names the current segment and opens a list of the others, current one first, a check on it. Not borrowed: ElvUI's profile switcher (a settings dialog) and any Battle.net launcher account drawer.

## 1. The one question and the one job

**"Which character is this site showing me, and can I change it?"** The answer is on screen on every page, in the bar, with no scroll: the crest, name and spec of the current character, or "Sign in".

| | 360 px phone | 1280 px desktop |
|---|---|---|
| Signed in with a current character | 44 px class crest with a chevron badge in the 56 px bar; the full name is the first row of the Menu (section 7) | Crest, full name in class colour, chevron; spec and level on a second line (level alone at 1024 to 1279) |
| Signed out, no pointer | 44 px dashed crest ring with a person glyph | Bordered button "Sign in" |

The job: one control that sets the site's current-character pointer (`fs.currentCharacter`). The planner, simulator, BiS and tier list already read it. Choosing a character here is the only way to change it from the chrome; nothing about gear, talents or numbers is shown or restated here.

## 2. Decisions, with the recommendation for each

1. **Bar order, left to right:** wordmark, the eight doors, the selector, Discord. Doors stay centred between wordmark and the right group (`justify-between`). The selector is the only bordered control in the bar, so it is the one primary element. **Discord loses its border** (today a 38 px bordered button): text link in `--color-nav` with the icon, **the real Discord mark** (not the bell glyph the site used) at 18 px; below 1440 the label goes and the mark stays, with a tooltip `Discord` (and `aria-label="Discord"`) on hover and focus.
2. **Doors, in this order: Planner, BiS, Simulator, Logs, Rankings, Tier List, Guides, Get set up.** BiS is the gear-gap door, so it sits next to the planner; Tier List sits beside Guides, away from Rankings, which the player reads as log rankings. "Leveling BiS" becomes "BiS" in `PRIMARY_NAV_ITEMS`, the phone menu grid, tests, `aria` strings and any doc that quotes the label (grep `Leveling BiS`; page-body headings that describe leveling bands are the owner's call, open question 5). The order change touches `nav.ts`, `Header.test.ts` and any e2e that clicks by position.
3. **The closed selector shows the character's full name at every width above the phone bar** (owner, round 2). Forever names can be two words ("Obnoxious Yell"). The slot is sized for the name, then the name only shortens with an ellipsis when it genuinely cannot fit. Widths are fixed per breakpoint (no shift): see 3.B for the box sizes and the ellipsis thresholds. Crest-only happens only on the phone bar, where the full name is the first row of the Menu.
4. **No wrap, and what gives.** The bar never wraps. Measured on the mock's real markup (Chromium, Barlow, signed in, "Obnoxious Yell"): single row, 81 px tall, no horizontal overflow at 1024, 1100, 1279, 1280, 1439, 1440 and 2000; 56 px at 1023 and below (768, 1023, 360 and 390 measured). In order of what gives as width falls:
   - **1440 and up:** everything, Discord with its label.
   - **1280 to 1439:** Discord goes icon-only (44 px); gaps shrink. The selector is unchanged (232 px, name and spec and level).
   - **1024 to 1279:** the selector's second line becomes `{Spec} · {level}` ("Fury · 60"; the class is dropped because the crest carries it, the name never is), box 176 px; door type 13 to 12 px, tracking .12em to .08em (.06em under 1100); door gap 22 to 14 px (13 under 1100); wordmark 20 to 17 px; Discord 40 px.
   - **Slack:** the smallest space between wordmark and doors, and between doors and selector, signed in: 35 px each at 1024 (14 px is the floor, so 21 px spare), 42 px at 1280, 73 px at 1440. Signed out: 80, 42 and 73 px.
   - **No "More" overflow**: eight short doors fit, and an overflow hides doors (tenet 4).
   - **At 1023 and below, the phone bar.** This moves the Menu breakpoint from `md` (768) to `lg` (1024). Today 768 to 1023 wraps the doors onto a second row. Tablets get the 56 px bar and the menu grid; at 560 to 1023 the bar has room for the full name beside the crest (up to 200 px) and shows it.
5. **Signed out: the sign-in door moves into the selector slot, and the panel leads with the addon.** One slot, two states, same width. The closed label is "Sign in" (with "or paste an export" under it at 1440). The panel is headed "Add your character" and leads with "Paste an export" and "Get the addon", because Blizzard returns no Forever characters today and the addon is the real way in. Sign-in is the quiet last row (email, or Battle.net as an identity, which keeps characters across devices). **No gold button anywhere in the panel.** Cost: sign-in is two taps instead of one; the review accepted that once the panel stopped leading with Battle.net, and the phone Menu carries a first-row Sign in so it is also one tap from there. `/login` stays as the page the row goes to.
6. **Battle.net is not in the selector until the import returns Forever characters.** No "Link Battle.net" row and no gold "Sign in with Battle.net" button. A Battle.net-sourced character that the account already has is listed like any other. When the import works, the row comes back as a quiet row shown only to an account with no Battle.net link (open question 3).
7. **Where sign out lives: `/account`, as today.** The list header's quiet "Manage" link goes there; there is no second "Your account" row. Sign out, anonymize, set-as-main, guilds, premium all stay there. The selector shows no battletag and has no account items.
8. **The pointer written on choosing:** `{ source: 'armory', ref: character.key, label: "{Name} · {Spec} {Class}", classSlug, savedAt }` through `writeCurrent`, then `window.dispatchEvent(CURRENT_CHARACTER_CHANGED)`. `'armory'` already means "the site's stored character, by key" (`sources.ts`). A pasted export that is not on the account stays `source: 'code'`.
9. **The page reacts without a reload** by the existing event: every pointer consumer (spine bar, planner, simulator, BiS "your spec", tier list callout) listens and re-renders. On `/planner` and `/sim`, whose state is in the URL, the selector also calls `history.replaceState` with `plannerHrefFor` / `simHrefFor` so a refresh keeps the choice. A page that cannot re-render in place (none should remain) falls back to one `location.reload()`; that page is a defect to list, not a design.
10. **Order of the list: the current character, then last chosen on this browser (most recent first), then newest sync for characters never chosen here.** A sync reorders only the never-chosen tail, so the list does not reshuffle under the player's hand. "Last chosen" is a small local list of keys (`fs.characterOrder`, at most 20, same try/catch rules as `current-character.ts`); it is a convenience, never the source of truth.
11. **Stale is more than 3 times the player's own median gap between syncs, clamped to 3 to 14 days.** A nightly raider who logs out every evening is stale after 3 days; a weekend player after 14. With fewer than three syncs on record the threshold is 7 days. The plain age is always shown on the row; ember colours only the words past the age, never the name and never the crest ring colour of the row. The closed bar shows a stale or failed **current** character as a small ring mark on the crest (12 px, 2 px ring, top-right, ember for stale, red for failed, ember for an expired session) with a tooltip on hover and keyboard focus. Failed is a Battle.net refresh whose last attempt errored. Both stay selectable: the data on the row is the last good one, and the row says so.
12. **Faction marks** are 14 px, in the faction colour (`--color-alliance #6fb1ff`, `--color-horde #ff6b5c`): the game emblem's silhouette filled with that colour (a CSS mask over the shipped emblem WebP). This is a variant of `FactionMark` for list rows only (the design system's own sizes are 16, 20 and 36 px and are unchanged); it needs the design-system owner's nod, open question 7.

## 3. Anatomy by region

Tokens (all existing): `--bg #07090d`, `--raised #0d111a`, `--line #262e40`, `--line-soft #1c2230`, `--line-warm #3a3326`, `--line-warm-strong #4a4030`, `--gold #e5b955`, `--muted #9a9484`, `--nav #b9b3a4`, `--strong #f2eee4`, body `#c9c2b2`, `--ember #d66e28` (stale), `--wipe #f0736b` (failed), class colours from `tokens.css` (`classColorVar`). Type: Cinzel 700 wordmark; Barlow for everything else; `label` 11 px uppercase .12em for the list heading. Components reused: the circular ringed `ClassCrest` (never a square class icon, letter square or Blizzard avatar), `FactionMark` at 14 px, `label`, the 44 px `rowLink` hit target, `Skeleton`, `StatePanel`-style message block, `rounded-control` radius 4, `rounded-panel` radius 6. **New shared component: `CharacterSelector.svelte`**, replacing the `session` slot's `AccountMenu` nav mode. Tenet 10: this spec is its design review. Nothing else is invented; the popover is the `panel` look with the existing shadow.

### 3.A The bar

Padding `18px 48px` desktop (shrinks per 2.3), inner width capped at 1344 px and centred above 1440 (the same container the page content uses, so wordmark and selector line up with page edges at 2000; border line stays full-bleed). Wordmark: seal 30 px and Cinzel 20 px, tracking .10em (28 px and 17 px at 1279 and below). Doors: Barlow 700 13 px, uppercase, tracking .12em, `--nav`; current door `--gold`; hover `--strong`; focus-visible 2 px gold outline, 2 px offset; hit height 44 px (the existing `min-h-11` on touch widths; desktop keeps the 18 px padding so the visual height is unchanged).

### 3.B The closed selector (a `button`, `aria-haspopup="dialog"`, `aria-expanded`)

| Width | Contents | Box | Name ellipsis begins at |
|---|---|---|---|
| 1280 and up | crest 32 px, full name (Barlow 600 14 px, class colour), second line `{Spec} {Class} · {level}` (12 px `--muted`), chevron 14 px; the ring mark top-right of the crest when stale, failed or expired | 232 x 44 fixed, 1 px `--line`, `--raised`, radius 4, padding `0 12px 0 6px`, gap 10 | name box 146 px: about **22 characters** ("Obnoxious Yell" is 14 characters and takes 89 px, 61% of the box). The spec line fits the longest real one, "Beast Mastery Hunter · 60" |
| 1024 to 1279 | crest, full name, second line `{Spec} · {level}` ("Fury · 60", the widest, "Beast Mastery · 60", is about 95 px of the 103), chevron, ring mark | 176 x 44 fixed | name box 103 px: about **16 characters** |
| 560 to 1023 (phone bar, tablet) | 36 px crest in a 44 px hit area with the ring mark only (no badge on the crest), a **10 px gap**, the name (Barlow 600 14 px, class colour), a 14 px chevron after the name | auto; the name box is `min(200px, 100vw - 460px)`, so it ellipsises before it can touch the wordmark or the 8 px gap to the Menu button | 200 px at 768 and up: about **30 characters** |
| below 560 (phone) | 36 px crest in a 44 px hit area, a 14 px chevron badge at its bottom-left (the only width that has a badge), ring mark top-right. No name in the bar | 44 x 44 | not applicable: the Menu's first row shows the full name in a box of about 250 px, about **38 characters**, which no real name reaches |

Character counts are for mixed-case names at 6.4 px per character, measured from a sample string; an all-capital or very wide name truncates sooner, a narrow one later. Widths are fixed per breakpoint, never content-sized, so a long name or the session resolving does not move the doors. When the name is cut, the `title` and the `aria-label` carry it whole. The `aria-label` is `{Name}, {Spec} {Class}, level {n}. Change character.`, with `Sync is stale.`, `Sync failed.` or `Signed out.` before "Change character." when a mark shows.

Signed-out closed: dashed 1.5 px ring `#5a5546` 32 px with a person glyph, `Sign in` (14 px 600 `--strong`) and `or paste an export` (12 px muted) at 1280 and up (232 px); `Sign in` in full, never cut, a 122 px box with a 24 px dashed ring and a 13 px glyph, at 1024 to 1279; the dashed ring alone below 1024, `aria-label="Sign in or choose a character"`.

Signed out with a pasted export in the pointer, and a signed-in session that has expired with a pointer still stored: the closed slot shows that character like any other. An expired session adds the ember ring mark with the tooltip `Signed out. Sign in to see your other characters.`; the list shows the one row and the line `Your session ended. Sign in to see your other characters.` (ember, 12 px) above the Paste and Sign in rows.

Signed out with a pasted export in the pointer: the closed slot shows that character like any other.

### 3.C The open list, desktop (a non-modal popover, `role="dialog"` named "Choose a character")

392 px wide, right edge flush with the selector's right edge, 10 px below the bar, `--raised`, 1 px `--line-warm`, radius 6, shadow `0 18px 40px` black at 70%. It overlays the page; no scrim. Regions, top to bottom:

1. **Header, 44 px:** `label` "Your characters" left; "Manage" right, a quiet link (12 px 600 `--muted`, underlined, 4 px offset; links `/account#characters`). **Gold marks the current row only.**
2. **List: six rows (a row with three lines is 72 px) fit in 440 px with no scroll.** With seven or more the spec picks **scroll under a fade, not a "Show all" row** (a second click to see your own characters is the disclosure tenet 4 forbids): the list is 476 px (six rows and the top half of the seventh), scrolls inside, and a 36 px fade at its foot sits over the seventh row. With **more than eight** a 36 px "Filter by name" input is added above the list. `nav-open-many.png` shows nine characters: the filter, six rows, the seventh fading out. Five or six characters never see a fade or a clipped row. More than 8 characters adds a 44 px filter input above the list ("Filter by name") that filters in place. 
3. **Action rows, 52 px each,** hairline above each: "Paste an export" always; "Get the addon" only when the account has no addon-synced character. Icon tile 36 px, 1 px `--line`, radius 4. Two setup rows at most, usually one. There is no footer and no "Link Battle.net" (decision 6).

**Order, checkable on screen.** Each row's third line carries its source and age, which is what the order is made from, so a reviewer can read the rule off the list. The rule: the current character; then characters chosen on this browser, most recent choice first; then the rest, newest sync first. In `nav-open.png` (six characters):

| # | Row | Why it is here | Line 3 |
|---|---|---|---|
| 1 | Obnoxious Yell | current | `Addon · 12 minutes ago` |
| 2 | Frostbyte | chosen yesterday (`fs.characterOrder[0]`) | `Battle.net · 3 days ago` |
| 3 | Shadowmend | chosen 3 days ago (`fs.characterOrder[1]`) | `Pasted export · 2 days ago` |
| 4 | Oakheart | never chosen here; newest sync of the rest | `Battle.net refresh failed · last good 3 days ago` |
| 5 | Quickshot | never chosen; next newest | `Addon · 5 days ago` |
| 6 | Treewalker | never chosen; oldest | `Addon · 19 days ago · Log in to the game to update` |

Rows 2 and 3 are older syncs than row 4 and still sit above it, which is the point: a sync never moves a character the player has chosen. Rows 4 to 6 read in descending recency, 3, 5, 19 days (a failed row sorts by its last good capture). In `nav-open-many.png` the never-chosen tail reads 3, 5, 6, 8, 12, 19 days.

**Row (64 px min, grid `36 | 1fr | auto`, gap 12, padding `8px 14px`):**
- Crest 36 px, ringed in class colour.
- Line 1: name, Barlow 600 15 px, class colour, then a 14 px faction mark in the faction colour (decision 12).
- Line 2: `{Spec} {Class} · {level} · {Realm}`, 13 px body colour; when the ruleset is not Normal the realm reads `{Realm} · {Ruleset}` (Hardcore, PvP, Roleplay: `rulesetLabel`). Omit any part the API did not send; never a placeholder.
- Line 3: `{Source} · {age}`, 12 px `--muted`, plain text, no pill. Source words: `Addon`, `Battle.net`, `Pasted export`. Age is relative ("12 minutes ago").
- Trailing: an 18 px gold check on the current row; a 44 x 56 "Retry" button on a failed row; nothing otherwise.
- Current row: `#e5b9551a` fill, 2 px gold left edge, `aria-current="true"`. One level of emphasis: gold marks the current row only; stale and failed use text colour, not fills.

**Stale (past the rule in decision 11):** line 3 keeps the plain age in `--muted` and adds ` · Log in to the game to update` in `--ember`. For a Battle.net row: ` · Refresh from your account` (links `/account`). The name stays in its class colour (an orange druid is never ember).
**Failed:** line 3 becomes `Battle.net refresh failed · last good {age}` in `--wipe`. The row is still selectable; "Retry" is a sibling button, not nested.

### 3.D The sheet, phone (a modal bottom sheet)

Opens from the crest. Scrim black at 65%, covering the page and the bar. Sheet: full width, radius 12 px top corners, 1 px `--line-warm`, max height 86% of the dynamic viewport, 20 px grab handle area (a 36 x 4 px bar), header 48 px (`label` "Your characters", Close 44 x 44 right), the same rows and actions as 3.C, bottom padding `env(safe-area-inset-bottom)` plus 14 px, sheet ends at the screen edge (no strip below it). **The sheet shows every character**: at 844 px high it holds six rows and the paste row with no scroll (`nav-phone-open.png`); more than that and the list scrolls inside with the 36 px fade, the header and the action rows pinned. Close: Close button, tap on the scrim, Escape, or a downward swipe on the handle. The page behind does not scroll.

### 3.E Signed-out popover and sheet

Heading `label` "Add your character". No intro paragraph and no gold button. Rows, top to bottom: **"Paste an export"** (title 14 px 600 `--strong`, sub `From the addon. Adds a character to this list.`), **"Get the addon"** (same style, sub `Export a character from the game.`), then **"Sign in"** as the quiet last row (title 13 px 600 `--nav`, sub `Email or Battle.net. Keeps your characters.`, links `/login`). No footer. The same three rows fill the phone sheet.

Signed out with a pasted export: the heading is "Your characters"; one row (source `Pasted export`); the line `Not saved to an account. Sign in to keep it.` (12 px muted); then Paste an export and Sign in.

### 3.F Paste an export, inline

"Paste an export" expands in place: the row is replaced by a 12 px `label` "Addon export", a 76 px mono textarea (placeholder "Paste the export here", focus moves into it on open), a 44 px "Use this character" and a 44 px "Cancel". The "Use this character" and "Cancel" buttons are neutral (1 px `#4a4030` border, `--strong` text); there is no gold fill. A valid export becomes the current character (`source: 'code'`, or `'armory'` if the account matches it), the popover closes, focus returns to the selector. An invalid export keeps the box and says "That is not a Forever Sixty export." in `--wipe` under it.

## 4. States

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Closed selector | as 3.B | border `#4a4030`, fill `#101624`, chevron to `--strong` | 2 px gold outline, 2 px offset | border gold, fill `#101624` | n/a | slot holds its fixed width; pre-paint from the cached `/v1/me` and the pointer (section 8); with neither, a 32 px skeleton ring and two skeleton lines | signed-out slot (3.B) | pointer shown if there is one, else signed-out slot; the list carries the error |
| Row | as 3.C | `#101624` fill, 2 px class-colour left edge; nothing moves | 2 px gold ring inside the row (`outline-offset -2px`) | fill `#0a0d15` | n/a | three 64 px skeleton rows (36 px ring, 120 and 210 px lines) | see "none" below | failed row, 3.C |
| Retry | 44 px, `--line-warm` border | text `--strong`, border `#4a4030` | gold outline | pressed `#0a0d15`; label becomes "Trying" with the button inert | n/a while a refresh runs | the row's line 3 reads "Refreshing from Battle.net" in `--muted` | n/a | stays on the row, line 3 as failed |
| Action row (Paste, Get the addon, Sign in) | as 3.C | `#101624` fill | gold ring inside the row | `#0a0d15` | "Get the addon" is absent, never disabled, when the account has an addon character | n/a | n/a | n/a |
| Ring mark on the crest | 12 px ring, ember or red | tooltip below the selector: `Last synced {age}. Log in to the game to update.` / `Battle.net refresh failed. Showing the last good data.` / `Signed out. Sign in to see your other characters.` | same tooltip on keyboard focus of the selector | n/a | n/a | n/a | n/a | n/a |
| Discord (below 1440) | real mark, `--nav` | mark to `--strong`, tooltip `Discord` | gold outline, 2 px offset, tooltip | `--strong` | n/a | n/a | n/a | n/a |
| Manage | 12 px `--muted`, underlined | text to `--strong` | gold outline | `--strong` | n/a | n/a | n/a | n/a |

**List states:** signed in with many (3.C); **one** (the single row, still opens: the paste row lives here); **none** (message "No characters yet." and "Paste an export from the addon to add one.", then Paste an export and Get the addon; the closed slot shows the signed-out button, not a crest); **stale**; **failed**; **loading**; **list did not load** ("Your characters did not load." "Check your connection and try again. You can still paste an export." Try again, then Paste an export); **signed out**, **signed out with a pasted export** and **session expired** (3.B and 3.E).

Keyboard: Enter, Space or ArrowDown on the selector opens and puts focus on the current row. Inside: ArrowUp and ArrowDown move between rows (roving `tabindex`, one tab stop for the list), Home and End jump, a letter jumps to the next name starting with it, Enter or Space chooses and closes, focus returns to the selector, and a polite live region says "Now using {Name}." Escape closes and returns focus. Tab from the list moves on through the action rows, then out of the popover, which closes. The phone sheet is modal: focus is trapped, the first stop is Close, Escape closes. Touch parity: every target is 44 px or more; no state depends on hover; the sheet closes by tap, button or swipe.

## 5. Hierarchy

Selector: the one bordered control in the bar, so it is the primary element. Doors are secondary text, current door gold. Discord is tertiary: no border, muted colour, mark-only below 1440. In the list the current row is the one emphasised row and the only gold in the panel (no gold button, no gold link); everything else is equal weight on purpose and is ordered by last chosen so the order is the player's own. Source and age are the quietest text in the row (12 px muted); the realm is body colour because two characters with one name on two realms need telling apart.

## 6. Copy (verbatim)

| Where | Text |
|---|---|
| Doors, in order | `Planner`, `BiS`, `Simulator`, `Logs`, `Rankings`, `Tier List`, `Guides`, `Get set up` |
| Discord | `Discord` (mark only below 1440, with the tooltip `Discord`; `aria-label="Discord"`) |
| Closed, signed out | `Sign in` and, at 1280 and up, `or paste an export` |
| Closed `aria-label` | `{Name}, {Spec} {Class}, level {n}. Change character.` (with `Sync is stale.`, `Sync failed.` or `Signed out.` before the last sentence when a mark shows) / `Sign in or choose a character` |
| Closed second line | `{Spec} {Class} · {level}` at 1280 and up; `{Spec} · {level}` at 1024 to 1279 |
| Ring mark tooltips | `Last synced {age}. Log in to the game to update.` / `Battle.net refresh failed. Showing the last good data.` / `Signed out. Sign in to see your other characters.` |
| Closed second line | `{Spec} {Class} · {level}` |
| List heading | `Your characters` ; manage link `Manage` |
| Row line 2 | `{Spec} {Class} · {level} · {Realm}` (`· {Ruleset}` unless Normal) |
| Row line 3 | `Addon · 12 minutes ago`, `Battle.net · 3 days ago`, `Pasted export · 2 days ago` |
| Stale (past the rule) | `Addon · 19 days ago · Log in to the game to update` ; Battle.net: `Battle.net · 19 days ago · Refresh from your account` (plain age always, ember on the last clause only) |
| Failed | `Battle.net refresh failed · last good 3 days ago` ; button `Retry` ; while running `Refreshing from Battle.net` |
| Paste row | title `Paste an export`, sub `From the addon. Adds a character to this list.` |
| Paste form | label `Addon export`, placeholder `Paste the export here`, buttons `Use this character`, `Cancel`, error `That is not a Forever Sixty export.` |
| Addon row | title `Get the addon`, sub `Export a character from the game.` |
| Filter | placeholder `Filter by name` |
| None | `No characters yet.` / `Paste an export from the addon to add one.` |
| Load error | `Your characters did not load.` / `Check your connection and try again. You can still paste an export.` / `Try again` |
| Signed-out heading | `Add your character` |
| Signed-out rows | `Paste an export` / `From the addon. Adds a character to this list.`; `Get the addon` / `Export a character from the game.`; `Sign in` / `Email or Battle.net. Keeps your characters.` |
| Session expired | `Your session ended. Sign in to see your other characters.` |
| Pasted, signed out | `Not saved to an account. Sign in to keep it.` |
| Live region | `Now using {Name}.` |
| Phone sheet | heading `Your characters` (`Add your character` signed out); Close `aria-label="Close"` |
| Phone Menu, first row | the full name (class colour), `{Spec} {Class} · {level}`, chevron; signed out: `Sign in` / `or paste an export` |

No source pills anywhere in the bar or the list: source is a word on a quiet line (owner 2026-09-30, provenance stays quiet).

## 7. Phone and wide screens

**Phone and tablet, 1023 and below:** the 56 px bar holds the wordmark (Cinzel 16 px, seal 28 px, the `minmax(0,1fr)` column), the 44 px selector crest and the 44 px Menu button; 8 px between. Measured, no overflow and no wordmark clipping at 360, 390, 768 and 1023 (at 360 the wordmark ends at 238 px and the selector starts at 246 px). From 560 up the full name sits beside the crest (up to 200 px); below 560 the bar is crest-only and the name lives in the Menu.
**The Menu (`nav-phone-menu.png`):** opens beneath the 56 px bar, as today. Its **first row is the current character**: a bordered `--raised` row, 56 px, crest 36 px, the full name in class colour, `{Spec} {Class} · {level}` under it, a chevron; tapping it opens the sheet. Signed out the first row is `Sign in` / `or paste an export` with the dashed ring, and opens the signed-out sheet (`nav-signed-out-phone.png`). Beneath: the eight doors in a three-column grid of 44 px cells in the new order (Planner, BiS, Simulator / Logs, Rankings, Tier List / Guides, Get set up), then Discord as a full-width 44 px bordered row with the real mark. With "BiS" the old longest label is gone; the widest is "Rankings" or "Get set up" (about 90 px in a 108 px cell at 360). At 640 and up the grid is four columns.
**The sheet** (3.D) opens from either the crest in the bar or the Menu's first row.
**Never collapses:** the current character's crest (the answer to which class), the full name in the Menu header (the answer to which character), the word Sign in at 1024 and up and in the Menu, every row's name, spec, level, source and age, the Retry button on a failed row.
**Wide screens:** the bar's contents sit in the 1344 px container, centred, from 1441 up; the selector keeps 232 px; the popover stays attached to the selector's right edge, so at 2000 it opens under the selector, not at the screen edge.

## 8. Performance and polish

- No layout shift: selector widths are fixed per breakpoint (3.B: 232, 176, auto on the phone bar); the Menu button keeps its grid column; the bar's height is 81 px desktop and 56 px phone in every state. CLS from the bar is 0.
- No missing-icon flash: Base.astro's pre-paint script, which already sets `data-pointer`, also sets `data-pointer-class="{slug}"` from the stored pointer and writes the name and `{Spec} {Class} · {level}` into the slot from the pointer's `label` (`Name · Fury Warrior`, so the full name is known before any fetch); the global CSS maps the attribute to the 128 px crest WebP as the selector's background image, so the crest and the name paint with the first frame (`nav-prehydration.png`). Only the matching class's file is requested. The island then mounts the same crest as an `img` with `width` and `height` set. Row crests are the same shipped 4 to 7 KB WebPs, and the list renders only after the user opens it, so none of them is on the page-load path.
- Hydration: `client:idle-after-load`, as `AccountMenu` today; the closed selector is real HTML before it (cached `/v1/me` entry, `session-cache.ts`), so the first interaction before hydration still opens nothing and loses nothing: the slot is an `<a href="/account#characters">` until the island swaps it for the button; the chevron is reserved (invisible) so nothing moves.
- No font swap: Barlow and Cinzel are the loaded faces; the bar introduces no new font.
- Lighthouse (`web/lighthouserc.json`): the bar is on every audited page; the budgets are unchanged (performance, accessibility, SEO at least 0.95, LCP at most 1700 ms, TBT at most 100 ms, CLS at most 0.05). Accessibility: muted `#9a9484` on `#0d111a` is used only at 12 px and up for secondary text; `--ember` and `--wipe` text sit on `#0d111a` and pass AA at 12 px.

## 9. Data contract

Read, never typed:
- **The session:** `/v1/me` through the existing cached query: `user.id` (signed in or not). A 401 with a pointer still stored is the expired-session state.
- **Per character (`MeCharacter`):** `key` (the pointer's `ref`), `name`, `class` (crest and colour), `spec` (the line-2 word; omitted when unknown), `level`, `realm`, `ruleset` (shown unless `normal`), `faction`, `source` (`'bnet'` or `'export'`), `build.source` (`'addon'` or `'blizzard'`) and `build.captured_at` (the source word and the age). A character with no `build` shows `Battle.net` or `Addon` per `source` and no age.
- **The pointer:** `readCurrent()`: `source`, `ref`, `label`, `classSlug`. The current row is the character whose `key` equals an `'armory'` pointer's `ref`; a `'code'` or `'addon'` pointer that matches no account character is shown as a `Pasted export` row built from `label` and `classSlug`.
- **Last chosen:** `fs.characterOrder`, a local list of character keys, most recent first (new, written by the selector; at most 20).
- **Stale:** `now - build.captured_at > clamp(3 x median sync gap, 3 days, 14 days)`; 7 days when fewer than three syncs are on record. The clamp bounds and the multiplier are constants in one place. The median needs sync history the API does not return today: requested `build.median_sync_gap_sec` (nullable) on `/v1/me`. Until it exists the 7-day fallback applies to everyone.
- **Not available today, needed:** a failed-sync signal. `MeCharacter` has no field for "the last Battle.net refresh failed". Requested: `build.sync_error` (nullable string or enum) on `/v1/me`. Until it exists the failed state is unbuilt and omitted (tenet 8: never shown as fact if we cannot verify it). The Battle.net-link row is not built (decision 6), so no link flag is needed now; when the import works it will need `me.battlenet: { linked, import_enabled }`.
- The doors: `PRIMARY_NAV_ITEMS` and `TRAILING_NAV_ITEMS` in `web/src/lib/nav.ts`; `isNavItemCurrent` unchanged.

## 10. What this is NOT

- **Not an account menu.** No battletag, no sign out, no anonymize, no premium, no settings. The quiet "Manage" link is the only way to those, and sign out lives on `/account`.
- **Not a guild switcher.** No guild names, ranks or roster. A character's guild is shown by the pages that need it (the spine bar's `guildLine`), not here.
- **Not a character manager.** No set-as-main, no remove, no rename: `/account`.
- **Not a place to read a character.** No gear, item level, talents or sim numbers; the page the player is on shows those.
- No hero, no feature pitch, no source pills, no counts of characters or syncs.

## 11. Acceptance screenshots (the implementer attaches, same viewports, same states)

1. Bar closed, signed in, current character "Obnoxious Yell": 2000, 1440, 1439, 1280, 1279, 1024, 1023, 768. At each, a stated measurement: single row, no horizontal overflow, bar height 81 px (56 px at 1023 and below), the full name visible without an ellipsis. At 2000 the wordmark and Discord sit on the 1344 px content edges.
2. A 22-character and a 16-character name at 1440 and 1024 showing where the ellipsis begins (3.B).
3. 390 and 360: bar closed (no wrap, crest and Menu both 44 px); the Menu open with "BiS" in the grid and the full name in its first row; the sheet with six characters, no strip below it.
4. 1440: the open list with six characters (no scroll, no fade) and with eight (scroll, fade at the foot).
5. Signed out at 1440, 1024 and 390: closed and open, including the Menu's Sign in row; signed out with a pasted export; session expired.
6. The states sheet: hover and focus-visible on the closed selector and on a row; the three ring marks with tooltips; Discord with its tooltip below 1440; stale; failed; one; none; loading; load error; paste expanded.
7. Before hydration (JavaScript disabled) at 1440 and 390: crest and name already painted, same bar height, no shift when the island mounts.
8. A chosen-character sequence: pick another character on `/planner`, show the page re-rendered without a reload, then on `/tiers` the "Your spec" callout following it.
9. Same-scale side by side with `nav.png`, `nav-open.png`, `nav-phone-menu.png` and `nav-phone-open.png` at 1440 and 390.

Tests: end-to-end on every page that reads the pointer (selecting a character changes the page without navigation); the no-wrap and no-ellipsis measurement for "Obnoxious Yell" at 1024, 1279 and 1280; keyboard path (open, arrow, choose, focus return); the 44 px target floor on phone; the door order.

## 12. Open questions for the owner and the wow-player

1. **Discord demoted.** Borderless, and mark-only below 1440 so the selector is the only bordered control. The review accepted this with the real mark and a tooltip.
2. **Stale needs sync history.** The 3x-median rule needs `build.median_sync_gap_sec`, which `/v1/me` does not carry. Until it does everyone gets the 7-day fallback, which is wrong for a nightly raider. Needs an API lane.
3. **Battle.net comes back when the import works.** Then as a quiet row for an account with no Battle.net link, never a gold button. Owner to say when.
4. **The spine bar's Switch popover** (`CurrentCharacterBar` / `barSwitch`) does the same job on the pages that mount it. Recommendation: its Switch becomes a link that opens this selector and its own list is deleted. Touches a shipped surface; needs the owner's go-ahead.
5. **BiS page headings.** The door is "BiS". The page's own h1 and the home door cards that say "Leveling BiS" describe leveling bands. Recommendation: nav, menu and aria say "BiS"; page copy about leveling keeps the word "leveling". Owner to confirm.
6. **Tablet gets the phone bar (768 to 1023).** Now shown in `nav-768.png` and `nav-1023.png`: the full name sits beside the crest and the Menu holds the doors. The alternative is an overflow, rejected.
7. **Faction mark variant.** A 14 px silhouette in the faction colour is not in the design system's FactionMark sizes (16, 20, 36) and recolours the emblem. Needs the design-system owner's nod; the fallback is the unrecoloured emblem at 16 px.
8. **Race.** The review listed it as optional. Not shown: line 2 already carries spec, level and realm, race is not a gear or talent decision, and the API sends it only after a Battle.net import.

## 13. Round 2 changes (what moved, from the review and the owner)

- **Full name at every width above the phone bar** (owner): selector 232 px at 1280 and up with spec and level, 176 px at 1024 to 1279 with level, full name beside the crest on tablets, crest-only only below 560 with the name as the Menu's first row. Ellipsis thresholds in 3.B. Current character is "Obnoxious Yell" in every render.
- Door order: Planner, BiS, Simulator, Logs, Rankings, Tier List, Guides, Get set up.
- Battle.net out of the selector: no "Link Battle.net", no gold sign-in. The signed-out panel leads with Paste and Get the addon; sign-in is the quiet last row.
- The list fits six rows (440 px), fade only on overflow; the phone sheet shows every character, ends at the screen edge.
- Setup rows cut: Paste stays, Get the addon only without an addon character, no footer row, Manage is a quiet link, gold only on the current row, neutral paste buttons.
- Stale rule: more than 3x the player's own median sync gap, 3 to 14 days, plain age always, ember only past the rule and never on the name; stale, failed and expired show on the closed bar as a ring mark with a tooltip.
- Copy: "Log in to the game to update".
- Real Discord mark with a tooltip below 1440. Faction marks 14 px in faction colour. List order by last chosen, then newest sync.
- New captures: 768, 1023, 1280, 1439, 360, the phone Menu with "BiS", signed out at 1024 and on the phone, signed out with a pasted export, expired session, before hydration, the eight-character list.
- Not done: race on the row (open question 8).

## 14. Round 3 changes (the four round-two deciders)

1. **Pre-hydration, phone.** The earlier cells drew the desktop door row because they sat inside a 1440 page. `nav-prehydration-phone.png` is now its own 390 px capture, "before" and "after" stacked, both the real 56 px phone bar. `nav-prehydration.png` is desktop only.
2. **Seven or more characters.** The spec picks scroll under a fade (3.C item 2), not "Show all". `nav-open-many.png` shows nine characters: the filter input, six full rows, the seventh fading under the 36 px fade. The earlier capture hid the seventh row entirely behind the fade, so it looked identical to `nav-open.png`.
3. **768 and 1023.** The chevron badge no longer sits on the crest; the crest carries only the ring mark, there is a 10 px gap to the name, a chevron follows the name, and the name box is capped so it ellipsises before it reaches the wordmark or the Menu button (3.B). The badge on the crest exists only below 560.
4. **1024.** Second line is `{Spec} · {level}` ("Fury · 60"): the player asked to keep class or spec, the owner asked for the level. Signed out reads "Sign in" in full in a 122 px box with a 24 px ring.
5. **List order visible** (3.C): a table naming why each row sits where it does, and the example data was reordered so the mock obeys it (the never-chosen tail is newest sync first).
