# F6 Talents (read only): build spec for iOS and Android

Author: ux-designer, 2026-10-31. Refines `design/specs/2026-10-30-mobile-apps.md` §5.6, §7.1, §8 and §13 (Revision 2 blocker 1 stands). Where this file and §5.6 differ, this file binds. Mock: `design/mocks/2026-10-30-mobile-apps/screens_talents.py`, boards `f6-talents`, `f6-rank`, `f6-states`, `f6-xl` (`-ios`, `-and`) in `renders/`.

## Revision 1 (2026-10-31), the fix round on the wow-player HOLD

Verdict: `design/reviews/2026-10-31-mobile-talents/mock-player-verdict.md`. The rules below bind and override anything later in this file that disagrees (the sections are patched to match).

1. **Closed trees.** On Your build, every tree tab whose tree holds a dashed cell carries a 8 pt `--color-wipe` dot after its points. Spoken value: `talents.a11yTabDiffers` (`28 points. Talents that differ from the guide: 2.`); a tab with none reads `talents.points`. In the stacked tabs the dot sits after the points.
2. **No gold on Your build.** "Available" is dropped from the compare. The planner's gold means "you can spend a point here", which on a read-only compare sends the player to talents the guide does not want while the guide's missing points sit on dashed cells. Cells are now: maxed and filled (green), **open** (tier reached and prerequisite held, rank 0: `--color-line` border, face at full brightness) and **locked** (tier or prerequisite not met: `--color-line` border, face 40%). Open versus locked is independent of unspent points. What the guide wants is shown by the direction chip (item 3).
3. **Dash direction.** Each dashed cell wears a corner chip with the signed difference, top left, hanging 8: `+2` when your rank is above the guide's (over: respec it), `−2` when below (the guide takes points here). The chip and the dashed ring are `--color-wipe` for over and `--color-gold` for under (gold here means "the guide wants this", a link to the missing points); the chip carries the meaning, so colour is never alone. When `diffCount > 0` and `missingCount > 0`, a `meta` line under the unspent line says `talents.missing`.
4. **Colour convention kept.** Maxed stays green and "has points" green, as the site planner (`CELL_BORDER`, `CELL_PILL`); the in-game frame is gold for maxed. Reason: one design system with the planner (tenet 10), and the badge states the rank in text on every cell. Not changed.
5. **Requirement lines in the sheet.** Adopted, in the game's words, in `--color-wipe`, only when unmet in the drawn build: `talents.requiresTree` (tier x 5 points in the tree) and `talents.requiresTalent` / `talents.requiresTalentOne` (the prerequisite's rank). The planner's `messages.tierLocked` and `prereqMissing` are refusal messages for an edit, not tooltip lines, so their wording is not reused.
6. **Arrowheads.** Adopted: a 7 pt filled triangle in the line's colour (1.5 `--color-bg` outline) at the dependent end of every prerequisite line. The planner draws none; a follow-up for the planner lane is flagged so the two stay one system.
7. **iOS title.** The title is inline (centred) on this screen, not large: it frees about 48 pt so the capstone tier clears the tab bar on iPhone 17, and removes the large title bleeding behind the bar. The bar ground is solid `--color-bg` from the status bar to the bottom of the back button (iOS `.toolbarBackground(.visible)`). This deviates from §3.2's large title on pushed screens, for this screen only.
8. **Capstone visibility.** With the inline title and Matches, tier 7 is fully above the tab bar on iPhone 17 and Pixel 9. In Differs the two extra lines push it half under the bar: it is the first scroll. Accepted; the pitch (70) is not shrunk.
9. **`pickSpec` reworded** (a spec comes from spending points, and the addon sends it): `We do not know this character's spec yet, so there is no guide build to compare. Open the addon once to send it.` The `Open the planner` button is removed from this state. Home's `home.answer.pickSpec` has the same flaw; flagged to the coordinator, not changed here.
10. **Copy changes.** `talents.differ` and `home.next.talentsDiffer` now end `guide build.` (`... differ from the {band} guide build.`), as does `talents.toTake`; `talents.legend` gains the chip sentence.
11. **Names in mixed case.** The sheet's talent name is `name` role (Barlow 700), not the uppercase condensed `sec`. The tree tabs were already mixed case.
12. **Matches with unspent points.** The unspent line shows (a level 55 character who spent exactly the band 50 build's 41 points reads `Matches` and `5 points unspent.`).
13. **Band strings are sized to the band's bottom level.** A band's string holds `band - 9` points (band 20: 11, 50: 41, 60: 51). **Decision: count as the planner does.** `BandCompare` runs `diffAgainstBand` on the band's whole string and the build's ranks with no level cut (`loadFromBandAtLevel` cuts only when loading). So a level 58 character (49 points) against the 41-point band 50 build reads `Differs` by at least their 8 extra points, on the site and in the app. Parity with the planner and "the count equals the drawn dashed cells" win; the app does not invent a second rule. To explain it, when `yourTotal > guideTotal` a `meta` line reads `talents.beyond`. A smarter rule (the points past the guide's length are free) is a change to `diffAgainstBand` plus both apps, to be ruled by the owner, and is flagged. No board for it; a unit test covers it.
14. **Elemental Mastery is 0 in the guide string.** That is the data's (the band's talent string), not a design choice, and is not ours to change; the accuracy reviewer owns it. The mock draws the data as shipped.

## 1. The one question and the first-screen answer

**Question:** "Is my build the guide build?"

**Answer, visible without scrolling** on iPhone 17 (402 x 874), Pixel 9 (412 x 915), iPhone SE 3 (375 x 667) and a 360 x 640 Android, at the default text size:
- the status line: tag `Matches` (green) with `Your build is the {spec} {band} guide build.`, or tag `Differs` (gold) with the count sentence;
- the two controls (Guide build / Your build, the three tree tabs);
- the top of the open tree. On 402 x 874 tiers 1 to 6 show; the last tier and the footer are the first scroll. On the small phones tiers 1 to 5 show.

The answer never depends on the tree. The tree is the proof.

## 2. Reference and what is borrowed

| Reference | Borrowed | Not borrowed |
|---|---|---|
| The in-game talent frame | A 4-column grid, one tier per row, rank `n/max` on each icon's corner, green for points, gold for "can take a point", grey and dimmed for locked, prerequisite arrows between icons, the tooltip's `Rank n/max`, current text, `Next rank` | Spending points; the paper background art; hover |
| The site planner (`web/src/lib/planner`, `components/planner`) | `cellState`, `CELL_BORDER`, `CELL_PILL`, `connectorPath` routing, `diffAgainstBand`, the tooltip's content order, `bandCompareCopy` wording, `selectBand` band choice | The rail, DPS figure, gear row, Load button (one job per page: the planner owns planning) |
| F2.6 item tooltip sheet (this app) | The sheet: header icon 56, name, tooltip panel `#0a0d14`, detents, close rules | |
| F4 sign-up segments (this app) | The two-segment control look (§4.5 Sign-up segments) | |

## 3. Reachability and who sees what

- **Entry:** one way only. Home's `Next for {name}` row `Talents` (F2.1) pushes F6 onto the Home tab's stack. There is no tab, no deep link (`apps/shared/links.json` has none) and no push notification for it.
- **Signed out:** unreachable. Home signed out has no Talents row, and `Pick your class` opens BiS. If the session dies while F6 is open (401), the standard F1 expiry flow runs and the stack pops to Home signed out. F6 never draws a sign-in prompt.
- **Character:** F6 shows Home's current character and has no selector (§3.4: pushed screens have none). Back returns to Home; switching is done there.
- **Home row and F6 share one function** per platform (`TalentCompare`, see §4.4). Home's row never reads the home endpoint's `talents {state, points_differ, band_points}` while lane A3 is held; when A3 ships, a unit test asserts it equals `TalentCompare` on the fixtures.

## 4. Data, exactly

Existing static data and `/v1/me` only. No new endpoint, no request F6 owns beyond the static files below.

### 4.1 Inputs

| Input | Source | Used for |
|---|---|---|
| Active build | `/data/active-build.json` (cached per §5.0) | URLs |
| The trees | `/data/{build}/talents/{class}.json` | Layout, names, max rank, prerequisites, `ranks[].description`, icons |
| The guide build | `/data/{build}/bis/{spec}.json`, one band (4.2) | `talents` string |
| Your build | `/v1/me` `characters[].build.talents.trees`, `build.data_build`, `characters[].level`, `.faction`, `.spec`, `.class`, `.race` (already cached by Home) | Your ranks |
| Icons | `/data/{build}/icons/{icon}.webp` | Cells, sheet header |

`{spec}` is the BiS file slug the BiS tab already uses for this character's spec (one mapping, in one place).

### 4.2 Choosing the guide band (the planner's own rule, not `bare`)

The brief proposed `bare`. **This spec uses the planner's rule instead**, because the planner is the source of truth and the count must match it. `BandCompare` calls `bandEntryFor(file, band, faction)` with no preset, and `selectBand` (`web/src/lib/bis/presets.ts`) then returns the **raid** entry when the band has one, else **bare**, else the first entry. At 60 the two presets can carry different talent strings (the Elemental fixture differs in the Restoration tree), so choosing `bare` would tell a player `Differs` here and `Matches` on the site. Rule:

1. **Band by level** (`characters[].level`): 60 or more gives band 60; 20 to 59 gives `floor(level / 10) * 10` (20, 30, 40, 50); under 20 there is no band.
2. **Entries** are the file's `bands` where `band` and `faction` (the character's) match.
3. **Preset:** the entry with `preset == "raid"` if present, else `preset == "bare"`, else the first entry.
4. No band, no file (404), or no entry gives the **no build** state (§6).
5. The `{band}` word in copy is the band's label: `60`, or `50 to 59` for bands 20 to 50 (the planner's `bandLabel`).

### 4.3 Drawing both trees

- Trees are the file's `trees` sorted by `position`. Talent cells are the tree's `talents` in list order. Both rank strings use that order.
- A rank string is one base-36 digit per talent, tree by tree. The guide string joins trees with `-`; Your build is `trees[]`, one string per tree. A short or missing string reads 0 for the rest (a tree stored as `"0"` is empty). Extra digits are ignored. Each digit is clamped to the talent's `max_rank`.
- **Totals** (tab numbers, `{m}`) are summed from the drawn ranks, never stored, never read from `build.talents.points`.
- **Your build is drawn on the active build's trees only when valid for them.** If `build.data_build` equals the active build, draw. If it differs, fetch `/data/{data_build}/talents/{class}.json` (static, cached) and draw only when every tree lists the same talent ids in the same order as the active file. Otherwise, or if that file is missing, treat the export as **stale** (§6).
- The guide build has no level cut: it is the band's whole string, as the planner compares it.

### 4.4 The compare, on the device (`TalentCompare`)

For every tree `i` and talent `j`, with `g` the guide rank and `y` your rank:
- `diffCount += max(y - g, 0)` (your points where the guide spends fewer; the planner's `diffCount`);
- `missingCount += max(g - y, 0)`;
- the cell is **dashed** when `y != g`.

Status:
- no dashed cell: `Matches`;
- else `diffCount > 0`: `Differs`, `{n}` is `diffCount`, `{m}` is your total (`talents.differ`);
- else (your build is inside the guide's and only short of it): `Differs`, `{n}` is `missingCount` (`talents.toTake`), as the planner's `toTakeSuffix`.

The count and the dashes come from the same pass over the same drawn ranks, so they cannot disagree (Revision 2 blocker 1). `Matches` and the dashes also agree: `Matches` is exactly "zero dashed cells". Note the count can be smaller than the number of dashed cells (the mock gives 2 points of count and 4 dashes); the legend says dashed means "rank differs", and the sentence says "points differ".

**Test vector** (Elemental Shaman, level 60, Horde, band 60 raid). Guide `5510331300001350-030000000000000000-0533520000000000`; yours `5512311100001350-030000000000000000-0433520000000000`. Expected: `diffCount 2`, `missingCount 5`, 4 dashed cells (Reverberation `+2`, Call of Flame `−2`, Elemental Warding `−2` in Elemental; Tidal Focus `−1` in Restoration, so the Restoration tab carries a dot), totals 28, 3, 17 (= 48), unspent 3. With yours equal to the guide: 0, 0, no dashed cell, `Matches`, totals 30, 3, 18 (= 51), unspent 0.

### 4.5 Unspent, open and locked

- `allowed = clamp(level - 9, 0, 51)`; `unspent = max(allowed - yourTotal, 0)`. The line `{n} points unspent.` shows when an export exists and `unspent > 0`, in every state with a drawable Your build and on both segments, including Matches.
- **Open** means the tree holds at least `5 x tier` points in the drawn build and the prerequisite holds its rank; **locked** is anything else. Both apply to rank-0 cells only; a cell with points is filled or maxed. There is no gold "available" state (Revision 1, item 2).

## 5. Layout, region by region

Tokens as §4.1 of the base spec. Sizes in pt (iOS) and dp (Android); the same number on both unless stated. All margins are 16.

### 5.1 The pushed screen

Top to bottom (iOS shown; Android has the title in the app bar and no large title):

1. **Bar.** iOS: glass back button `Home` (44 tall) and the inline centred title `Talents` (Revision 1, item 7), solid ground. Android: `TopAppBar` with back arrow and `Talents` (Barlow Condensed 800, 22). Reused as `pushed()` screens in §3.2.
2. **Offline strip** (only offline), 36 tall, directly under the bar. Reuse `OfflineStrip` (§5.0): `Offline. Showing what you had at {time}.` (`offline.strip`).
3. **Status block**, 8 above, min height 28, always present (no layout shift). A `Tag` (22 tall, `micro` role, Tag/TagView, kinds green, gold, muted) then the sentence in `body`, `--color-strong`, numbers in mono. They sit on one line and wrap together; from iOS AX1 / Android 1.5x the tag sits above the sentence. This is the screen's one primary element.
4. **Unspent line**, `meta` muted, 4 above, then the `talents.missing` line when it applies (Revision 1, item 3), then the `talents.beyond` line when it applies (item 13). Secondary.
5. **Compare bar**: two controls, 12 above the first:
   - **Segments** `Guide build` | `Your build`. A custom two-button control built from the sign-up segments' look (`SignupSegments` idiom): container `--color-raised`, 1 pt `--color-line` border, `--radius-control` (8) on iOS and a pill (24) on Android, buttons 44 tall (iOS) and 48 tall (Android, `SingleChoiceSegmentedButtonRow`). Checked: `--tile` fill, `--color-strong` label, plus a 16 check glyph on Android. Unchecked: `--color-muted` label. Hidden entirely when only one build can be drawn (a lone segment is noise).
   - **Tree tabs** `{Tree} {points}`, three equal-width tabs, 8 below the segments, 44 (iOS) and 48 (Android) tall, a 1 pt `--color-line` rule under the row. Selected: `--color-strong`, 700, a 2 pt (iOS) / 3 dp (Android, Material `PrimaryTabRow`) `--color-gold` underline. Unselected: `--color-nav`. `{points}` in mono, then the 8 pt wipe dot for a tree with a dashed cell (Revision 1, item 1).
   - **Pinned** below the bar once scrolled (`LazyColumn` `stickyHeader`, SwiftUI pinned section header), `--color-bg` ground, 1 pt `--color-line-soft` bottom rule, 8 padding below. It is not pinned at iOS AX1 and above or Android 1.5x and above (it would fill the screen); there it scrolls with the page.
6. **The tree** (5.2), 16 above.
7. **Legend**, `meta` muted, 8 below the tree: `talents.legend` when at least one cell is dashed, else `talents.tapHint`.
8. **Footer**, 16 above: `meta` muted `talents.footer`, then a text link (`Text link` §4.5, gold 15/600, 48 tall, 16 external glyph) `talents.planLink`. It opens `https://foreversixty.gg/planner?class={class slug}&race={race slug}` in the browser. The planner is at `/planner` (there is no `/talents` route); do not link `/talents`. The link text wraps and breaks inside the URL at large text; it is never cut.
9. 130 of bottom space above the tab bar so the footer clears it.

### 5.2 The tree

- **Grid.** 4 columns (0 to 3), tiers 0 to 6 (7 tiers; the region is always reserved at 7, so switching tabs never moves the legend or footer). Cell = 48 x 48, `--radius-control` (8), a 2 pt state border, icon 44 x 44 inside (radius 6) on `#131824`. Tier pitch 70 (22 between tiers).
- **Column pitch** `p = min(96, floor((W - 32 - 48) / 3))`, the grid centred. Check: 402 gives 96 (grid 336, 33 each side); 375 gives 96 (grid 336); 412 gives 96; 360 gives 93 (grid 327, 16 each side). Minimum clearance between neighbours at 360 is 45, and the 6 hang of the badge plus the 5 of the dashed ring leaves 34. The right-most badge reaches 6 past the icon, inside the 16 margin. It fits; nothing clips at 360 or 375.
- **State (the planner's `cellState`, without `available`).**

| State | Border | Face | Badge text |
|---|---|---|---|
| maxed (rank = max) | `--color-kill` | full | `--color-kill` |
| filled (0 < rank < max) | `--color-kill` | full | `--color-text` |
| open (rank 0, tier and prerequisite met) | `--color-line` | full | `--color-muted` |
| locked | `--color-line` | 40% opacity | `--color-muted` |

- **Rank badge.** Bottom right, hanging 6 past the icon corner. Mono, `{rank}/{max}`, 1 pt border in the state colour, `--color-bg` fill, radius 6, padding 1 x 5. Text 11 at default, scales with text size to a ceiling of 14 (iOS `.dynamicTypeSize(...DynamicTypeSize.xxxLarge)` on the badge; Android sp capped by `min(fontScale, 1.3)`).
- **Dashed (Your build only).** 2 pt dashed outline, radius 12, 3 outside the cell (5 from the icon edge), `--color-wipe` when over the guide and `--color-gold` when under, with the signed chip (Revision 1, item 3): mono 11 (14 at large text, same cap as the badge), `--color-bg` fill, 1 pt border in the same colour, top left, hanging 8. Never on Guide build.
- **Prerequisite lines (`connectorPath`).** Straight down a column (prerequisite and dependent share a column); straight across a row (same tier); or down the prerequisite's column to the dependent's row, then across into its side. Every line starts and ends at a cell edge (never over an icon). Stroke 2 over a 5 `--color-bg` halo, round caps, ending in a 7 pt arrowhead into the dependent. Gold once the prerequisite holds the rank the dependent needs (on the drawn build), `--color-line` before. Lines draw under the cells. Hidden from VoiceOver and TalkBack.
- **Icon states.** The tile and its state border draw at once on `#131824`; the image fades in over 120 ms. A 404, or a 1 x 1 placeholder file, keeps the tile and shows the first two letters of the name (Barlow Condensed 800, 13, `--color-muted`), as `TalentCell.svelte` does. No broken-image glyph, ever.
- **Selected** (sheet open): 2 pt solid `--color-gold` ring, 9 outside the cell (outside the dashed ring), kept while the sheet is open.

### 5.3 Which tree opens

The tabs are in the game's order (the planner's `TreeTabs` order, so a digit string and a tab always agree). The tab that **opens** is the tree with the most points in the build first shown (ties: the first). This is the spec's tree for every spec in practice and needs no spec-to-tree table. §5.6's "the spec's tree first" is satisfied as "opens first"; the order of the tabs is not changed. The choice survives toggling Guide/Your build and does not jump.

### 5.4 Segment default

Opens on `Your build` when Your build can be drawn (export present, not stale) and a guide exists; else there is no segment control.

## 6. States

| State | Trigger | What draws |
|---|---|---|
| **Loading** | First open, the tree file not yet cached | Status block: a 72 x 22 and a 220 x 14 skeleton. Segments skeleton full width at their real height. Tabs: three 100 x 14. Tree: 14 skeleton cells (48 x 48, radius 8) in a fixed sparse pattern inside the reserved 476-tall region. `meta` `talents.loading` below (also the announcement). Footer drawn. Never a spinner. A refresh never replaces content with skeletons. Home prefetches the current character's talent file, so this is rare. |
| **Matches** | No dashed cell | Status `Matches`; no legend dashes; `talents.tapHint`. |
| **Differs** | At least one dashed cell | Status `Differs` with `talents.differ` or `talents.toTake`; dashed cells on Your build; `talents.legend`. |
| **Guide segment** | Guide build selected | Guide totals, no dashes, no gold; `talents.tapHint`. Status unchanged (it is about your build). |
| **No export** | `build` or `build.talents` absent | Guide build only. No segments. Status: muted tag `Guide build` and `home.next.talentsNone`. If there is no guide either, show that line and `talents.noList` and no tree. |
| **Stale export** | `data_build` trees differ (4.3) | As No export, line `talents.staleExport`. |
| **No build** | Level under 20, no file, or no entry (4.2) | Your build alone, no segments, no dashes. Status: muted tag `Your build` and `talents.noList`. The unspent line still shows. |
| **No spec** | `characters[].spec` empty | Your build alone. Status: muted tag `Your build` and `talents.pickSpec`. No button. |
| **Offline** | Last fetch failed offline and a cache exists | Strip, then whatever is cached. No write exists on this screen. |
| **Error** | The talent file is neither cached nor fetchable (offline with no cache counts) | `talents.failed` in `--color-wipe` with role alert, secondary `common.tryAgain`; footer still draws. |
| **Guide failed** | The BiS file failed, the tree file is fine | Your build drawn without the compare; status line `talents.guideFailed` in `--color-wipe` with `common.tryAgain`; no segments. |

## 7. The rank sheet

A tap on a cell opens a sheet (same idiom as F2.6): iOS `.sheet`, Android `ModalBottomSheet`. Close: swipe down, tap the scrim, VoiceOver escape (two-finger Z), TalkBack back, Escape on a keyboard.

**Content** (20 side padding):
1. Header: the talent's icon 56 (tile, state border, face at full brightness), the name (`name` role, Barlow 700, `--color-strong`, wraps, breaks inside a word if it must), under it `Rank {n}/{max}` (`talents.rank`, mono `meta` muted). `Rank 0/5` is shown for an empty talent.
2. The tooltip panel (`#0a0d14`, 1 pt `#3a4256`, 8 radius, 12 / 14 padding), the planner's order:
   - the **current rank text**, `ranks[n - 1].description`, `body` `--color-text`, only when `n > 0`;
   - **`Next rank`** (`talents.nextRank`, `micro` label) then `ranks[n].description`, `body` `--color-muted`, only when `n < max`.
   - the **requirement lines**, only when unmet in the drawn build, `body` `--color-wipe`, 12 above each: `Requires {5 x tier} points in {Tree} Talents` (`talents.requiresTree`), `Requires {n} points in {Prerequisite}` / `Requires 1 point in {Prerequisite}` (`talents.requiresTalent`, `talents.requiresTalentOne`).
3. **The guide line**, only on `Your build`, only when your rank differs from the guide's: `Guide build: rank {g}/{max}` (`talents.guideRank`), `body` 600 `--color-strong`, in a box with a 1.5 pt dashed outline (wipe when you are over, gold when under, as the cell's ring), 8 radius, 10 / 12 padding, 16 above (it echoes the dashed cell). Nothing on the Guide segment.

The panel text is the talent file's `ranks[].description`, verbatim, never trimmed.

**How it never clips.** It is a system sheet, not a tooltip anchored to a cell, so there is no placement maths and the bottom row and the right column are no different from any other cell.
1. The sheet is always the full width of the screen. Text wraps inside it with no line limit, no ellipsis and no scale-down.
2. The body is one vertical scroll container with the header inside it. Bottom padding is 16 plus the system inset (34 on iOS, the gesture-bar inset on Android). The last line is always reachable: XL.4 shows the end of the content with its padding.
3. Detents. Default text: iOS `[.height(content, clamped to 220...medium), .large]`; Android partially expanded (content height, at most half the window) then expanded. From iOS AX1 and Android 1.5x: a single full-height detent (iOS `[.large]`, Android `skipPartiallyExpanded = true`), opening at the top below the status bar. Content longer than the detent scrolls inside it.
4. The tapped cell is not covered. When the sheet opens, the tree scrolls so the cell's bottom sits at least 28 above the sheet's top edge (F6.4 draws a bottom-row cell this way). The cell keeps the gold selected ring.
5. At AX5 / 2.0 the description wraps to about 14 lines; the sheet is full height and scrolls (XL.3, XL.4).

## 8. States of every interactive element (§8 of the base spec)

| Element | Default | Pressed | Focused (keyboard, VO, TalkBack) | Disabled | Loading | Error |
|---|---|---|---|---|---|---|
| Talent cell | per 5.2 | 0.96 scale, `rgba(229,185,85,.10)` wash over the tile | 2 pt gold ring, 9 outside (same look as selected) | none: a locked cell is a look, not disabled; it still opens the sheet | tile and border at once, image fades in | tile stays with two letters |
| Segment | 5.1 | checked ground at 60% | gold 2 pt ring inset 3 | not applicable (the control is hidden rather than disabled) | skeleton | not applicable |
| Tree tab | 5.1 | `rgba(229,185,85,.10)` | gold 2 pt ring inset 3 | not applicable | skeleton | not applicable |
| Plan link, `Try again`, `Open the planner` | §4.5 | §8 | gold ring | not applicable | not applicable | not applicable |
| Sheet | closed | grab handle drag | focus trapped, returns to the cell | not applicable | not applicable | not applicable |

Touch targets: cells 48, segments and tabs 44 / 48, links 48. Hardware keyboard: Tab enters the tree, arrows move between cells by the planner's `moveFocus` (`grid.ts`: left/right walk the occupied cells in reading order, up/down jump to the nearest column of the next occupied tier), Space or Return opens the sheet. Touch and keyboard open the same sheet. Haptics: `.selection` / `SEGMENT_TICK` on a segment or tab change (§7), none on a cell.

## 9. Hierarchy

| Region | The one primary | Visibly secondary |
|---|---|---|
| Status block | The sentence (`body`, strong) with its tag | Unspent line, legend, footer (`meta` muted) |
| Compare bar | The checked segment | The tabs (the selected tab strong, the rest nav) |
| Tree | Cells with points (green) | Locked cells (40% face), prerequisite lines |
| Sheet | The talent's name and its rank text | `Next rank` (muted), the guide line (boxed) |

One gold-filled button per screen at most: none here (`Open the planner` is secondary). No gear, DPS, weights or sim figure anywhere on this screen (one job per page: BiS owns gear, the planner owns talents, sims own numbers). The footer link is the only way out.

## 10. Large text

- Icons do not scale (48). Badges cap at 14. The grid, lines and ring offsets are fixed.
- From iOS AX1 and Android 1.5x: status tag sits above the sentence (which wraps; never an ellipsis); the two segments stack as two full-width rows; the three tabs become a stacked list of three rows (min 56 tall, name left, mono points right, selected row `--tile` fill with a 3 pt gold rail on its leading edge) inside one bordered card; the compare bar is not pinned.
- Below that, the three tabs stay in a row while every label fits on one line at the current size (SwiftUI `ViewThatFits`; Compose measure with `TextMeasurer`); otherwise they stack. Labels never wrap or truncate.
- The legend, footer and plan link wrap; the link text breaks inside the URL.
- First-screen answer at AX5: the title, the tag and the whole sentence (iPhone 17 shows them above the fold; XL.1).

## 11. VoiceOver and TalkBack

Icons beside their names and the prerequisite lines are hidden. Reading order is the visual order (status, unspent, segments, tabs, cells row by row left to right, legend, footer).

| Element | Label | Hint / state |
|---|---|---|
| Status | the tag then the sentence as one element: `Differs. 2 of your 49 points differ from the 60 build.` | announced politely on change |
| Segment | `Guide build`, `Your build` | selected trait / `selected` |
| Tab | label `{Tree}`, value `talents.points` (`28 points`), or with a dot `talents.a11yTabDiffers` (`28 points. Talents that differ from the guide: 2.`) | selected |
| Cell | `talents.a11yCell` (`Convection, rank 5 of 5`); on Your build with a different guide rank `talents.a11yCellGuide` (`Call of Flame, rank 1 of 3, guide rank 3`) | hint `talents.a11yHint` (`Shows the talent`) |
| Sheet | a heading with the name, then `Rank n/max`, the texts, the guide line; focus returns to the cell on close | |
| Loading | `talents.loading` announced once | |

## 12. Copy, verbatim

Reused (in `apps/shared/copy/en.json`): `talents.title` Talents; `talents.matches` Matches; `talents.differs` Differs; `talents.match` `Your build is the {spec} {band} guide build.`; `talents.differ` (changed, below); `talents.unspent` `1 point unspent.` / `{n} points unspent.`; `talents.guide` Guide build; `talents.yours` Your build; `talents.legend` (changed, below); `talents.rank` `Rank {n}/{max}`; `talents.nextRank` Next rank; `talents.footer` `Change talents in the game; the addon sends them on your next logout.`; `talents.noList` `No talent build published yet for this spec and level range.`; `home.next.talentsNone` `Not available yet: no talent export for this character. Open the addon once to send it.` (one string for Home and F6); `home.answer.openPlanner` Open the planner; `offline.strip`; `common.tryAgain`.

Changed values: `talents.differ` `{n} of your {m} points differ from the {band} guide build.` (and `home.next.talentsDiffer` to match); `talents.legend` `Dashed: your rank differs from the guide. +n is points over it, −n points under it. Tap a talent for both ranks and its text.`; `talents.pickSpec` as Revision 1, item 9.

New keys (add to `apps/shared/copy/en.json` under `talents.`):

| Key | Value |
|---|---|
| `talents.tapHint` | `Tap a talent for its rank text.` |
| `talents.toTake` | `{n} of the {band} guide build's points still to take.` (n int, band string) |
| `talents.missing` | plural, one `The guide build takes 1 point you have not.`, other `The guide build takes {n} points you have not.` |
| `talents.beyond` | `Level {l} is past the {band} guide build's {g} points, so points past it show as differing.` (l, g int notPlural; band string) |
| `talents.requiresTree` | `Requires {n} points in {tree} Talents` (n int notPlural) |
| `talents.requiresTalent` | `Requires {n} points in {talent}` (n int notPlural, at least 2) |
| `talents.requiresTalentOne` | `Requires 1 point in {talent}` |
| `talents.a11yTabDiffers` | `{n} points. Talents that differ from the guide: {k}.` (n, k int notPlural) |
| `talents.guideRank` | `Guide build: rank {g}/{max}` (g, max int, notPlural) |
| `talents.planLink` | `Plan a build on foreversixty.gg/planner` |
| `talents.staleExport` | `Your talent export is from an older game build. Open the addon once to send it again.` |
| `talents.pickSpec` | `We do not know this character's spec yet, so there is no guide build to compare. Open the addon once to send it.` |
| `talents.loading` | `Loading talents…` |
| `talents.failed` | `Your talents did not load.` |
| `talents.guideFailed` | `The guide build did not load.` |
| `talents.points` | plural, one `1 point`, other `{n} points` (the tab's accessibility value) |
| `talents.a11yCell` | `{name}, rank {n} of {max}` (notPlural) |
| `talents.a11yCellGuide` | `{name}, rank {n} of {max}, guide rank {g}` (notPlural) |
| `talents.a11yHint` | `Shows the talent` |

`talents.footer` no longer carries the planner sentence (the link replaces it). §5.6's "foreversixty.gg/talents" text is superseded: the page is `/planner`.

## 13. Performance and polish

- No layout shift: status block min height 28; tree region reserved at 7 tiers; segments and tabs reserved while loading; the offline strip enters at 36 and the content moves once, with no animation under Reduce Motion.
- No missing-icon flash: tile and border draw at once, the image fades in over 120 ms, icons cached on disk per build and prefetched with the talent file after launch. All 16 to 18 icons of the open tree load in the first frame set; the other two trees load when their tab is first opened (or in the background after the first draw).
- Toggling Guide / Your build and switching tabs draw from memory with no request. Crossfade of the border, dash and badge colours 120 ms.
- One `TalentCompare` per platform; unit tests use the §4.4 vector plus: a short string, a missing tree, `toTake`, level 59 vs 60 band choice, preset choice (raid present, raid missing), under 20, a stale `data_build`.

## 14. Acceptance captures

For `design/reviews/2026-10-31-mobile-talents/`, iPhone 17 simulator (iOS 26) and Pixel 9 emulator (API 35), fixture mode. Each is named `<screen>-<state>-ios.png` / `-and.png`.

| File stem | State | Size |
|---|---|---|
| `talents-matches` | F6.1 | default |
| `talents-differs` | F6.2 (the vector in 4.4) | default |
| `talents-guide` | Guide segment | default |
| `talents-differs-scrolled` | pinned compare bar, bottom tier visible | default |
| `talents-rank-bottom-row` | sheet on a tier 7 cell, the cell visible and ringed | default |
| `talents-rank-right-column` | sheet on a column 4 cell with the guide line | default |
| `talents-rank-maxed` | no Next rank | default |
| `talents-loading`, `-noexport`, `-nobuild`, `-nospec`, `-offline`, `-error` | states of §6 | default |
| `talents-differs`, `talents-differs-tree` | status wrap, stacked tabs, the tree | AX5 (`UIContentSizeCategory.accessibilityExtraExtraExtraLarge`) / font scale 2.0 |
| `talents-rank-large`, `talents-rank-large-end` | the sheet at the top and scrolled to the end | AX5 / 2.0 |
| `talents-differs-restoration-tab` | Your build with the dot on a closed tree | default |
| `talents-differs-360` (Android 360 x 640) and `-se` (iPhone SE 3) | the first screen | default |

Review compares them with the `f6-*` boards side by side.

## 15. Checklist

- [ ] Reached only from Home's Talents row; never signed out; the stack pops on a 401.
- [ ] Guide band chosen by `selectBand` (raid if present, else bare); under 20 is No build.
- [ ] Count is `diffCount` (or `missingCount` in the `toTake` case); `Matches` iff zero dashed cells; totals summed from the drawn ranks.
- [ ] 4 columns, 48 icons, pitch `min(96, ...)`, fits 360 and 375; 7 tiers reserved.
- [ ] Cells, badges, lines (with arrowheads) per 5.2; dashes and `+n`/`−n` chips on Your build only; no gold on Your build; a dot on every tab whose tree has a dashed cell.
- [ ] Sheet: name, `Rank n/max`, current text, `Next rank`, guide line only on difference; never clips (7, points 1 to 5); the tapped cell stays visible above it.
- [ ] Large text: icons stay 48, badges cap at 14, status, tabs, segments, legend, link wrap.
- [ ] Every state of §6 and every §8 state; VoiceOver and TalkBack per §11.
- [ ] No gear, DPS, weights or sim figure. Footer links to `/planner`.
- [ ] Captures per §14 attached; ux-designer and wow-player SHIP on them.
