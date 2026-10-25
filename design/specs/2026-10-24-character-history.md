# Character history charts: experience spec

Author: ux-designer. 2026-10-24. Lane 4 of the premium experience. Product spec: `docs/superpowers/specs/2026-10-24-premium-experience-design.md` section 4.4 (the promise, the rulings), section 5 (the account page's History row, the upsell sentence), section 2 (anti-goals), section 8. References: `docs/tenets.md` 1, 4, 9, 10, 11, 13, 14; `design/specs/2026-10-11-character-header.md`; `design/specs/2026-10-24-server-runs.md` (reserve rule, hint rule). Mock: `design/mocks/gen_character_history.py`, render `design/mocks/renders/character-history-states.png` (1440 and 390: the premium owner's section, every chart and list state, the free owner, the visitor, the account row).

Surfaces changed: `web/src/components/Character.svelte` (one mount point, 3.1; the free owner's sentence, 3.9), a new `web/src/components/character/CharacterHistory.svelte` (lazy, loaded only for owners with the flag) with a new pure model `web/src/lib/score/history-chart.ts`, `web/src/components/character/CharacterFightTable.svelte` (four optional props, 3.7), `web/src/lib/billing/copy.ts` (the upsell string, the account row, the /premium line), `web/src/components/account/AccountPremiumBlock.svelte` (one row). New copy goes in a new `web/src/lib/score/history-copy.ts`. `ForeverScoreCard.svelte` and the sparkline are not changed.

Two facts the brief got wrong, found in the code, that this spec is built on:

- The Forever Score runs 0 to **2100**, not 0 to 1000 (`web/src/lib/score/score.ts` `SCORE_MAX`, the five ceilings added). Tier lines are 0, 500, 1000, 1500, 2000 (`TIER_FROM`). Every axis and band below uses those.
- The score reads only **public, unflagged** reports (`api/internal/rating/score_v2.go` `ratedFights`: `r.visibility = 'public'` plus `presence.EvidenceReport`). A private or guild-visible fight is rated and stored, but it cannot move the score; its snapshot carries the score as it stood. Section 8 ruling 3 says how the chart tells the truth about that.

## 0. The experience in one paragraph

An owner on Premium opens their character and, directly under the Forever Score card, finds `History`. On the left, their score over time as a gold line on tier bands, the newest point selected and named (`1830 Epic`, `11 Dec 2026 · 2 fights`). On the right, their parse percentile over their kills of one boss, the boss chosen from a row of pills, the most recent kill first. Under each chart, the fights behind the points in the same row grammar as `Every ranked fight`, the selected point's rows marked. A gap in the record is drawn as a break with the rating module's reason inside it. A free owner sees the card as today plus one sentence. A visitor sees the card as today and nothing added.

## 1. The one question and the answer that must be visible

**"Am I getting better, and where?"**

| | 360 / 390 px phone | 1280 / 1440 px desktop |
|---|---|---|
| Premium owner | From the top of the `History` heading to the bottom of the score chart's date axis: at most 500 px at 360 (28 heading, 36 caption, 8 gaps, 20 sub-heading, 44 note, 56 readout, 213 chart at 320 px wide, 16 axis line). The readout (`1830 Epic`) and the line sit inside one 640 px screen once the heading is scrolled to the top. | Both charts side by side. Heading to the chart axes at most 540 px, inside a 1280 x 720 screen: the score chart and the parse chart are both whole on screen. |
| Free owner | One sentence under the score card (44 px). The card's sparkline is the answer, as today. | Same. |
| Visitor | The card's sparkline, as today. Nothing new. | Same. |

The primary answer is the **readout** (the selected point's number in its tier or parse colour) and the **line's direction**. `Up 1090 since 12 Oct` under each chart states the net. Acting on it is the fights list under the chart: every row links to its report.

## 2. The reference, and exactly what we borrow

**Warcraft Logs, a character's progress graph on a boss page.** Players trust it for this job: one dot per kill over time, the parse in its colour, the report one click away. Borrowed: one dot per kill; the parse number in the percentile ladder colour (`percentileToken`: grey under 25, green 25, blue 50, purple 75, orange 95, pink 99, gold 100); a readout instead of a floating tooltip; the report link on every point's row. Not borrowed: its metric toggle, legend, zoom and animation.

**Raider.IO, the score-over-time chart.** Borrowed: the score named with its tier and the sense that tiers are ground you cross. Ours draws that as bands behind the line in the game's quality colours (`docs/contracts/fixtures/score-labels.json`). Not borrowed: season resets on the x axis (our record never resets).

**The Forever Score card's own sparkline** (`ForeverScoreCard.svelte`, `Sparkline.svelte`, `lib/score/sparkline.ts`) is the site's chart grammar and the base: gold 2 px line, round joins, the newest point in its tier colour, `Up 330 since 10 Nov` at 12 px in `--color-text`, an image role with a screen-reader table beside it, one point written as a sentence and not a box. This spec adds axes, bands, selection and gaps and changes none of that grammar.

**The fights list** is `CharacterFightTable.svelte`'s row, unchanged in look: boss link, parse in mono with `percentileToken`, amount, date in mono muted, execution link.

## 3. Layout, region by region

Tokens (`web/src/styles/tokens.css`): `--color-raised #0d111a` panel ground, `--color-line #262e40` panel border and tier lines, `--color-line-soft #1c2230` grid and row dividers, `--color-gold #e5b955` the line, `--color-gold-hover #f5d27a`, `--color-strong #f2eee4`, `--color-text #e9e4d8`, `--color-muted #9a9484`, `--color-line-warm-strong #4a4030` hover borders. Type: Barlow for words; JetBrains Mono tabular for every number, date and axis label; sub-headings 15 px 600 `--color-strong`; axis and band labels 11 px; captions 12 px. Radius `rounded-panel` (6 px) for the panel, `rounded-control` (4 px) for pills and steps. No chart library: `web/package.json` ships none and none is added. Charts are inline SVG drawn by `history-chart.ts` in pixel units from the container's measured width, so text never scales.

### 3.1 Where it mounts

`Character.svelte` renders, in DOM order, `CharacterRatingPanel` (its first child is `ForeverScoreCard`, `-order-2`), then **`CharacterHistory` (new, `-order-2`)**, then `CharacterGearPanel` (`-order-2`). Equal `order` keeps DOM order, so the page reads: header, Forever Score, **History**, Gear, In game, the fight tables, Builds seen. Mount condition: the viewer owns the character, `me.entitlements.history` is true and `hasScore(foreverOf(data))`. Otherwise it is not mounted and reserves nothing.

### 3.2 The section

A `section` with `id="history"` and `data-testid="character-history"`, flex column, gap 8 px, as the Forever Score section. Rows:

1. `h2` `History`, `section-title text-[18px]` (the sibling headings).
2. Caption, 13 px `--color-muted`: `Includes your private and guild reports. Only you see this section.`
3. One panel, `bg-raised border-line rounded-panel border`, padding 16 px below 768 and 20 px from 768, **no gold top highlight** (the score card keeps the only one, so it stays primary). Inside, two columns from 1280 (`xl:grid xl:grid-cols-2 xl:gap-x-8`, the score card's own split); one column below, 28 px gap, score first.

Both columns share four rows so the chart tops line up: sub-heading, control row (min height 44 px), readout (min height 56 px), chart. Use `grid-template-rows: subgrid` on the two columns inside the `xl` grid.

### 3.3 The score column (Score over time)

| Row | Content | Style |
|---|---|---|
| Sub-heading | `Score over time` | 15 px 600 `--color-strong` |
| Control row | One sentence: `One point a day: the day's last rated fight.`, plus ` Rings are private or guild reports.` only when at least one ring is drawn | 12 px `--color-muted`, up to two lines, centred in 44 px |
| Readout | Left: the selected point's score, 28 px mono 600 in `tierTextColor(tier)` (no thousands comma, as the card writes `1830`), then the tier word 15 px 600 in the same colour; under it `11 Dec 2026 · 2 fights`, 12 px mono muted (`formatDayMonthYear`). Right: two 44 x 44 step buttons, earlier and later (a left and a right chevron). | Min height 56 px |
| Chart | 3.5 | |
| Under the chart | `Up 1090 since 12 Oct` (`scoreCopy.historyChange`), 12 px `--color-text`. Then one 12 px muted line per gap too narrow for an inside label: `Gap 20 Jul to 9 Aug: No rated fights` | |
| List | 3.7, label `Fights behind the score` | |

### 3.4 The parse column (Parse by boss)

| Row | Content | Style |
|---|---|---|
| Sub-heading | `Parse by boss` | as above |
| Control row | **The boss picker.** `role="tablist"` of `role="tab"` pills, one per boss the member has killed, newest kill first. Each pill: the boss's short name (the words before "the", so Golemagg for Golemagg the Incinerator) in 12 px 700 uppercase, tracking .06em, then the kill count in 11 px mono muted. Pattern: `MyReports.svelte`'s tab strip (44 px high, `rounded-control`, border `--color-line`, never wraps, scrolls sideways). Selected pill: `--color-gold` border, `--color-strong` text, ground `#e5b95514`. The row fades 28 px at its right edge when it overflows (`mask-image`), and the selected pill is scrolled into view on mount and on selection. | Min height 44 px |
| Readout | The selected kill's parse, 28 px mono 600 in `percentileToken`, then `parse` 15 px muted; under it `11 Dec 2026 · kill 5 of 5`. The same step buttons. | |
| Chart | 3.5 | |
| Under the chart | `Up 34 since 21 Nov` | |
| List | 3.7, label `Kills of Golemagg` | |

Default boss: the one with the newest kill. Parse is kills only (a wipe has no parse). Two difficulties of one boss are two pills, the second written `Boss · Hard` (the API's `difficulty` word); with one difficulty the word is never written. Every boss with at least one kill is in the picker. A boss with one kill shows the one-kill state (section 4), not a chart.

### 3.5 The chart (both)

Box: a `div` with `aspect-ratio` and `max-width: 100%`, width 100% of its column. Aspect: 3 / 2 below 640 px (320 x 213 at 390), 16 / 5 from 640 to 1279 (one column, up to about 1000 x 312), 2 / 1 from 1280 (588 x 294 at 1440). The SVG inside takes `width` and `height` in px from a `ResizeObserver`; before the first measure the empty box holds its aspect, so nothing shifts.

Plot margins: left 40 px (36 below 640) for y labels, right 14 px, top 12 px, bottom 26 px for dates.

Layers, bottom to top:

1. **Tier bands (score only).** For each tier whose range intersects the y domain: a rect across the plot in the tier's colour from `score-labels.json` at 8 % opacity (`common` white at 4 %). The tier name inside the band, 8 px from the left edge, 6 px above the band's lower boundary, 11 px Barlow 600, `tierTextColor` at 85 %, drawn only when the band is at least 22 px tall. Boundaries between tiers are 1 px `--color-line`.
2. **Grid.** 1 px `--color-line-soft` horizontal lines at each y label that is not a tier boundary. No vertical grid.
3. **Y labels.** 11 px mono `--color-muted`, right-aligned 8 px left of the plot. Only values the chart reaches. Score: the tier boundaries inside the domain when the domain spans more than 1200, else every 200 (every 100 up to a 600 span). Parse: 25, 50, 75, 100 inside the domain, and 0 when the domain starts there.
4. **Gap bands** (3.6).
5. **Selection guide.** 1 px vertical line, full plot height, `--color-line`; `--color-line-warm-strong` while a hover preview is active.
6. **The line.** 2 px `--color-gold`, round joins and caps, one path per unbroken run. No smoothing, no area fill, no shadow.
7. **Points.** Public points: 3 px radius gold dots. Private or guild points (score chart only): 3.5 px radius ring, 1.5 px gold stroke, `--color-raised` fill. Above 40 points only the selected and the newest are drawn as dots. The **selected** point: 5 px radius, 2 px `--color-strong` stroke, filled in its tier colour (score) or parse colour (parse); a ring point keeps the raised fill.
8. **X labels.** 11 px mono muted, `12 Oct` form (`shortDate`; with the year, `20 Dec 2026`, when the span crosses a year boundary, as the sparkline does). First label start-aligned to the plot's left edge, last end-aligned to its right edge, the rest centred; three labels below 640 px, up to six from 640, at least 96 px apart.

**Domains.** X is **proportional to time**, from the first point's date to the newest. Score y: `floor((min - 60) / 100) * 100` to `ceil((max + 60) / 100) * 100`, at least 400 wide (centred on the midpoint), clamped to 0 to 2100. Parse y: lower bound 0 when the lowest parse is under 30, 25 under 55, 50 under 80, else 75; upper bound 100. Parse bounds follow that rule, never the exact data, so a boss's chart does not rescale as points arrive within a band.

No animation, no draw-in, no transition: the resting state is the only state. No legend (one series). No axis titles: the sub-heading and the readout name the unit.

### 3.6 Gaps

A **gap** is a stretch the API lists in `gaps` (section 10), two consecutive dates more than **14 days** apart; the web draws the API gaps and derives none. The line is broken: no segment between them. The span is a full-height band, `--color-muted` at 8 % opacity, with 1 px dashed `--color-muted` edges at 40 %. The reason is written inside the band, 11 px Barlow `--color-muted`, centred, up to two lines, **when the band is at least 72 px wide**. A narrower band carries no label; its reason is the muted 12 px line under the chart (3.3). The reason strings are the API's, verbatim (section 8 ruling 2). Parse charts draw no gaps: a boss is not killed every week, and that is not a gap in the record.

### 3.7 The fights lists

Both lists use `CharacterFightTable.svelte` with four new optional props (a shared-component edit, section 11 decision 9):

- `heading: false`: no `h2` and no execution caption; the list's own label (12 px caps `label` style, `--color-muted`) replaces them.
- `columns`: omits the Spec column. The `md` grid becomes `minmax(0,1fr) 64px 88px 112px 72px`. Spec is the character's own spec in every row and the 588 px column has no room for it.
- `lead: 'date'` (the parse list): the first column is the kill's date (mono, `--color-muted`, linked to the report) instead of the boss, and the Date column is blank, because every row is the same boss.
- `selected`: a set of `report_id:fight_index` keys to mark. A marked row gets the `TARGET` treatment the table already has (2 px `--color-gold` left edge, `rgba(229,185,85,.08)` ground), without `id` or scroll.

Rows come from the API's `fights` (section 10), newest first, in the full `CharacterFight` shape, so the row, its parse colour, its amount link and its execution cell are the shipped ones. A fight in a private or guild report carries the word `Private` or `Guild` after the boss name, 12 px `--color-muted` (a word, not a pill). The boss link goes to `/reports/<id>?fight=<index>`.

Score list: the newest **8** fights. Parse list: the newest **8** kills of the selected boss. Under each, when more exist, a 44 px text button `Show all 31 fights` or `Show all 12 kills` (13 px 600 `--color-text`, underlined, 4 px offset). It expands the list in place. Above 100 rows the list stops at 100 and the button goes. While expanded the button reads `Show fewer`. Selecting a point whose rows are hidden expands the list and scrolls the first marked row into view (`scroll-margin-top: 96px`, the table's own).

### 3.8 Selection, hover, and what moves what

- **Select** (click, tap, Enter or Space on the focused plot, a step button, an arrow key) sets the selected point. The readout, the guide, the point and the list marks follow. A selected score point marks every fight of its UTC date; a selected parse point marks that kill.
- **Hover preview** (mouse only): moving over the plot previews the nearest point in the readout and moves the guide to `--color-line-warm-strong`. Leaving restores the selection. A preview never moves the list. There is no floating tooltip: the readout is the tooltip, so nothing overflows a phone screen.
- **Touch:** a tap selects the nearest point by x. A horizontal drag scrubs the selection live; the list updates on release. The plot has `touch-action: pan-y`, so vertical scroll is never captured.
- **Keyboard:** each plot is one tab stop (`tabindex="0"`, `role="group"`, `aria-label="Score chart. Use the arrow keys to move between points."` and `"Parse chart for Golemagg. Use the arrow keys to move between kills."`). Left and Right move one point; Home and End go to the first and the newest. The step buttons repeat this for pointer users. The picker is one tab stop; Left and Right move between pills (roving `tabindex`); Enter and Space select.
- Initial selection: the newest point. Changing the boss resets the parse selection to that boss's newest kill.
- A screen-reader table sits beside each chart (`Sparkline.svelte`'s pattern: an `sr-only` table with a caption, `Forever Score by date, oldest first` and `Parse by kill, oldest first`; each gap is a row reading `Gap from 26 Oct to 17 Nov: No rated fights`).

### 3.9 The free owner's sentence and the visitor

Free owner (signed in, owns the character, `me.entitlements.history` is not true, the character has a score): `Character.svelte` renders one `p` directly after the Forever Score section, 13 px `--color-muted`, min height 44 px, with the page's own 22 px (32 px from 768) gap above it: `History charts draw your score and parses over time. See plans.` `See plans` is the secondary link of `/premium` and the account block (`text-text underline underline-offset-2`, `href="/premium"`, `inline-flex min-h-11` on phone), no icon, no heading, no panel. The sentence is not shown for an owner with no score (the card says `No logs yet`), for a visitor, or for an anonymised character (404 page). The card's sparkline is unchanged for every viewer, including the Premium owner: it stays the card's glance and History is the detail.

Visitor: no History element, no sentence, no placeholder, no reserved space.

### 3.10 Account page and /premium

Account Premium block, History row (copy in `billing/copy.ts`, replacing `Ships November`): on, with a main character: `Open on your main character`, a link to that character's page at `#history`, underlined like the Server sims row; on, no character yet: `On`; off: `Sparkline only`. `/premium`: the line becomes `Character history charts over time.`; `shipsMonth.history` and its suffix go when the lane ships. A changelog entry in the past tense, dated, ships with the lane.

### 3.11 No layout shift

The section sits above Gear, so a late section would push the whole page. It reserves its box before data, as the server-runs list does: the panel's last rendered height in px, stored in `localStorage` as `fs.history.reserve` after every paint, with fallbacks of 1060 px at 1440 and 1670 px at 390 for the whole section (measured from the mock, re-measured by the implementer from the acceptance captures). The reserve is applied only under the existing `html[data-session='1']` hint plus a second pre-paint hint `data-history='1'`, set from the cached `/v1/me` snapshot when `entitlements.history` is true and the cached characters include this page's character. Anyone else reserves nothing and the static HTML of the character page is unchanged. The reserved box holds the skeleton (section 4). When the data resolves to the empty, one-point or error state the box collapses once, inside the paint that fills it; accepted, owner only. The chart's own box holds its aspect ratio before the first measure. A Playwright check asserts CLS at most 0.02 for the hinted case at 1440 and 390.

## 4. States

### 4.1 Interactive elements

| Element | Default | Hover | Focus-visible | Active | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|---|
| Plot (the chart) | gold line, dots, selected point | mouse: preview in the readout, warm guide; touch: none | 2 px `--color-gold` outline, 2 px offset, radius 4, around the whole plot box | drag scrubs (touch), click selects | n/a | skeleton box at the reserved aspect | see 4.2 | see 4.2 |
| Step button (earlier, later) | 44 x 44, 1 px `--color-line`, glyph `--color-text` | border `--color-gold`, glyph `--color-gold-hover` | 2 px gold outline, 2 px offset | ground `#ffffff0d` | at the first or newest point: glyph `#5d5a50`, border `--color-line-soft`, `aria-disabled="true"`, not removed from the tab order | n/a | n/a | n/a |
| Boss pill | 44 px, 1 px `--color-line`, text `--color-text` | border `--color-line-warm-strong`, text `--color-strong` | 2 px gold outline, 2 px offset | ground `#ffffff0d` | n/a | n/a | n/a | n/a |
| Boss pill, selected | gold border, `--color-strong` text, ground `#e5b95514` | unchanged | as above | as above | n/a | n/a | n/a | n/a |
| Fight row (boss, amount and execution links) | as `CharacterFightTable` | as shipped (underline on the link) | as shipped | as shipped | n/a | n/a | n/a | n/a |
| Fight row, marked | 2 px gold left edge, 8 % gold ground | unchanged | unchanged | unchanged | n/a | n/a | n/a | n/a |
| `Show all N fights` | 13 px 600 `--color-text`, underlined | `--color-gold-hover` | 2 px gold outline, 2 px offset | `--color-gold-hover` | n/a | n/a | absent with 8 rows or fewer | n/a |
| `Try again` (error) | `LoadError`'s own | `LoadError`'s own | `LoadError`'s own | `LoadError`'s own | n/a | while re-reading: skeleton | n/a | n/a |
| `See plans` (free owner) | `--color-text`, underlined | `--color-gold-hover` | 2 px gold outline, 2 px offset | `--color-gold-hover` | n/a | n/a | n/a | n/a |

Touch and keyboard parity: every control is at least 44 px high or tall on phone (pills 44, step buttons 44 x 44, `Show all` 44, `See plans` 44); the plot is operable by drag and tap and by the arrow keys. Tab order in the section: the score plot, its step buttons, `Show all`, the score list's links, the picker, the parse plot, its step buttons, the parse list's links.

### 4.2 Section and chart states

| State | What shows |
|---|---|
| Premium owner, data, several points | Section 3, both columns. |
| Premium owner, loading | The panel at its reserve with, per column, skeleton bars (sub-heading 140 x 15, readout 120 x 28, the chart box at its aspect, in `bg-line-soft`, as `Skeleton` draws them), `aria-busy="true"`, sr-only `Reading your history…`. The card above is already drawn. |
| Premium owner, error | The panel with one `LoadError`: `Your history did not load.` and `Try again` (re-reads in place). The score card and its sparkline above are untouched. |
| Premium owner, a score but no point carries a history row (every fight rated before snapshots, or re-rated by the backfill) | One `EmptyState`, no action: `No history yet. Each rated fight you upload adds a point.` Both columns are replaced by it. |
| Score, exactly one point | No chart box. The readout (score, tier, date), then `First score recorded 11 Dec. A line needs a second point.` at 13 px `--color-text`, then the list. The score column stays the width of its grid cell. |
| Score, with a gap | The line breaks; the band and its reason (3.6). |
| Score, a gap too narrow for its label | The band without a label; the reason as a muted line under the chart. |
| Score, rings | Private or guild points as rings; the control-row note names them. |
| Parse, a boss with two or more kills | Chart and list. |
| Parse, a boss with one kill | The selected boss's column shows `One kill so far. A line needs two.` (14 px `--color-muted`, 44 px), then the one-row list under the label `Kills of Gehennas`. No chart box, no readout. |
| Parse, no boss killed | No picker, no chart, no list: `No kills on record yet. Parse charts start with your first kill.` (14 px `--color-muted`). The score column is unaffected. |
| Free owner | The card, then the sentence (3.9). |
| Visitor | The card, nothing else. |
| Anonymised character | The 404 page; no card, no History. |
| Lapsed payer | A free owner here: the flag is false, so the sentence shows and the section is gone. History is not deleted; it is read again when the plan returns. |

## 5. Hierarchy

- **The section's primary element is the readout**: the 28 px mono number in its tier or parse colour, with the tier word. It is the largest type in the section and the only coloured text outside the bands and the parse numbers in the list rows.
- **One level of emphasis per column:** the gold line and the selected point. The other points are 3 px gold dots; the grid, the bands (8 %), the gap bands (8 %) and the axis labels are visibly quieter.
- **The score card stays the page's primary** for the score: 48 px number, gold top highlight, the only card with it. The History panel has none, and its numbers are 28 px.
- Secondary and visibly so: captions (12 px muted), axis labels (11 px muted), the picker's counts (11 px mono muted), the gap reasons (11 px muted), the `Private` and `Guild` words. Gold appears on the line and dots, the selected pill's border, the marked-row edge, focus rings and the list's links, as the page's other tables use it.
- Whitespace: 12 px between rows inside a column, 32 px between columns (28 px stacked), 6 px extra above each list's label.

## 6. Copy, verbatim

New strings live in `web/src/lib/score/history-copy.ts` (functions take the data and return the string; no number or date is typed into prose); the upsell, the account row and the /premium line live in `web/src/lib/billing/copy.ts`.

| Where | String |
|---|---|
| Section heading | `History` |
| Section caption | `Includes your private and guild reports. Only you see this section.` |
| Score sub-heading | `Score over time` |
| Score note | `One point a day: the day's last rated fight.` and, with rings, ` Rings are private or guild reports.` |
| Parse sub-heading | `Parse by boss` |
| Picker pill | `Golemagg` + the kill count (`5`); a pill's accessible name: `Golemagg, 5 kills` (`1 kill` for one) |
| Score readout | `1830`, `Epic`, `11 Dec 2026 · 2 fights` (`1 fight`) |
| Parse readout | `88`, `parse`, `11 Dec 2026 · kill 5 of 5` |
| Net change | `Up 1090 since 12 Oct`, `Down 40 since 12 Oct`, `No change since 12 Oct` (`scoreCopy.historyChange` for both charts) |
| Y axis | numbers only (`1000`, `1500`; `25`, `50`, `75`, `100`) |
| X axis | `12 Oct`, `11 Nov` (with the year, `20 Dec 2026`, across a year boundary) |
| Tier names in the bands | `Common`, `Uncommon`, `Rare`, `Epic`, `Legendary` (`tierLabel`) |
| Gap, inside the band | the reason: `No rated fights`, `Report flagged`, `Score not kept` |
| Gap, under the chart | `Gap 20 Jul to 9 Aug: No rated fights` |
| Step buttons, accessible names | `Earlier point`, `Later point` (`Earlier kill`, `Later kill` on the parse chart) |
| Score plot, accessible name | `Score chart. Use the arrow keys to move between points.` |
| Parse plot, accessible name | `Parse chart for Golemagg. Use the arrow keys to move between kills.` |
| Screen-reader table captions | `Forever Score by date, oldest first`; `Parse by kill, oldest first`; columns `Date`, `Score` / `Parse` |
| List labels | `Fights behind the score`, `Kills of Golemagg` |
| List columns | the table's own: `Boss` (`Date` on the parse list), `Parse %`, `DPS` or `HPS`, `Date`, `Execution` |
| Visibility words | `Private`, `Guild` |
| Show all / fewer | `Show all 31 fights`, `Show all 12 kills`, `Show fewer` |
| Loading (sr-only) | `Reading your history…` |
| Error | `Your history did not load.` + `Try again` |
| Empty, no points | `No history yet. Each rated fight you upload adds a point.` |
| Score, one point | `First score recorded 11 Dec. A line needs a second point.` |
| Parse, one kill | `One kill so far. A line needs two.` |
| Parse, no kills | `No kills on record yet. Parse charts start with your first kill.` |
| Free owner sentence | `History charts draw your score and parses over time. See plans.` |
| Account row | `Open on your main character`; `On`; `Sparkline only` |
| /premium line | `Character history charts over time.` |

Numbers: scores without a thousands comma (the card's form); amounts `toLocaleString('en-US')` (`1,402`); dates `11 Dec 2026` in readouts and lists (`formatDayMonthYear`), `12 Oct` on axes.

## 7. Phone and desktop

**Never collapses, at any width:** the readout (number, tier or parse word, date); the chart and its line; the selected point; the step buttons; the gap bands and their reasons (inside the band, or under the chart when narrow); the net-change line; the fights list's Boss (or Date), Parse % and DPS columns and its Execution column; the picker; the free owner's sentence.

**Collapses or thins:**

- Below 1280 the two columns stack, score first, each at the full column width.
- Below 640 the chart's aspect is 3 / 2 and x labels thin to three (first, middle, last); y labels keep their rule (only values reached). The Spec and Date columns of the fights list go, as on the shipped table (`CharacterFightTable`'s `max-md` grid `minmax(0,1fr) auto auto auto`); the date stays in the readout and on the row's report.
- Below 640 the picker scrolls sideways; its pills never wrap and each stays 44 px high.
- 360: the panel padding is 16 px, so the chart is 290 px wide (aspect 3 / 2, 193 px high); the control-row note may wrap to two lines inside its 44 px; the readout number stays 28 px.
- 1024 to 1279: one column, chart 16 / 5 at the full panel width (up to about 1000 x 312); the lists use the shipped `md` grid with Spec omitted.
- 1440 and **2000**: the section is as wide as its siblings in the centred 1344 column, never wider; at 2000 the container is centred and the panel is 1248 px, each chart 588 px. Nothing stretches past the grid.
- The chart is never replaced by a table or hidden behind a disclosure at any width (tenet 4); the screen-reader table is the only table form.

## 8. Rulings this spec makes

1. **The score's scale is 0 to 2100** and tier lines are the rating module's (section "Two facts"). The y axis labels only values the domain reaches.
2. **Gap reasons are the rating module's own, and today it has none.** The module has no gap vocabulary; the words in the brief (`no logs uploaded`, `spec changed`) exist nowhere in `api/internal/rating` (`spec` is derived from the newest fights, not a state that breaks a record; `report flagged` is real: `reports.flagged`). This spec fixes a closed vocabulary from facts the module does have, for the API lane to emit: `No rated fights` (the default: no fight was rated in the stretch), `Report flagged` (a report in the stretch is flagged, so its fights never count, `score_v2.go`), `Score not kept` (fights in the stretch were re-rated by the backfill, which writes no score snapshot, `score_v2.go` file comment). The web prints the string it is given and invents none.
3. **Private and guild fights are drawn as rings, because they do not move the score.** The product spec says the owner's chart includes private and guild reports. It does, as rings, and the control-row note says so. The score is computed from public reports only, so a ring sits at the score the rater recorded, to which a private report adds nothing, and the parse chart is where private kills change the picture. The alternative (public-only points on the score chart) is decision 1 in section 11.
4. **Flagged reports stay out of the owner's chart.** The owner-only query widens visibility (private, guild) and never un-flags: a flagged report cannot raise a score, so its fights are neither a point nor a parse point. The gap reason `Report flagged` tells the owner why the line stops.
5. **One point a day on the score chart,** the day's last rated fight (UTC date), exactly as the card's sparkline does, so the two never disagree about a day. The API returns every rated fight; the web collapses to the last of each date and uses the rest to mark list rows. The readout says `2 fights`.
6. **The x axis is proportional to time.** A raid week is clustered and a lull is empty, and the chart shows that honestly; it never spaces points evenly (the sparkline's evenly spaced x is a glance, not a chart).
7. **Parse is kills only,** one line per boss the member has killed at least twice, and a boss with one kill is a point list (the product spec). Every killed boss is in the picker so none is hidden; the one-kill state explains itself.
8. **The readout is the tooltip.** A hover preview, never a floating card.
9. **No chart library, no animation, no legend, no area fill.** The line, dots, bands and a shared readout.
10. **The History section is owner-only data and says so** in its caption; a visitor never receives the history payload.

## 9. Performance, polish, and acceptance

Limits:

- `CharacterHistory.svelte` is loaded with `createLazyComponent` only when the mount condition (3.1) holds, so the character page's bundle for visitors and free owners is unchanged by construction. No chart library; `history-chart.ts` is a pure module with unit tests.
- No new font, no image. The tier colours are the fixture's; the parse colours are `percentileToken`'s.
- No layout shift on hydration: section 3.11 (reserve, hints) and the chart's aspect box (3.5). CLS at most 0.02 for the hinted owner at 1440 and 390 in Playwright.
- Lighthouse (`web/lighthouserc.json`): no character page is in its URL list (the page is client-rendered from the API and gated), so the budgets (performance and accessibility at least 0.95, CLS at most 0.05) are not run against it; every static page in the list is unchanged by construction. The Playwright axe pass and CLS check above stand in for it.
- No missing-icon flash and no font swap: the section has no icon. The chart draws in one paint after the measure; the reserved box shows the skeleton until then.

Tests the lane adds: the product spec's e2e (seeded snapshots; the chart renders for the premium owner, the sparkline plus one sentence for the free owner, the sparkline alone for a visitor); unit tests for `history-chart.ts` (the score and parse domain rules, tick selection, gap breaking, the 72 px label rule, one-point and empty inputs, collapsing to one point a day); an axe pass on the section in both columns; a keyboard test (plot arrows, Home and End, picker roving).

**Acceptance screenshots the implementer attaches** (fixture build, API stubbed to section 10; 390, 1440 and 2000 unless stated):

1. Premium owner, the whole section: both charts, picker on the default boss, lists collapsed (390 stacked, 1440 side by side, 2000 centred).
2. The score chart with a gap (reason inside the band), and a narrow gap (reason under the chart), at 1440 and 390.
3. Score chart with rings (private and guild points) and the control-row note.
4. Selection: a middle point selected (readout, guide, marked rows), a hover preview at 1440, the step buttons' rest, hover, focus-visible and disabled (first and newest point), the plot's focus-visible ring.
5. Picker: default, hover, focus-visible, active, a boss with one kill (point list), another boss selected; the sideways fade at 390.
6. Lists: collapsed (8 rows), expanded (`Show fewer`), a marked row, a `Private` and a `Guild` row.
7. States: loading (skeleton at reserve), error (`Try again`), empty (`No history yet`), one score point, no kills, 1440 and 390.
8. Free owner (card and sentence) and visitor (card alone), 1440 and 390; the `See plans` focus-visible.
9. Account page History row: on with a character, on without, off, 1440 and 390.
10. The CLS number printed by the Playwright check, hinted case, 390 and 1440.
11. Same-scale side-by-side of each shipped surface with the mock's matching panel at 1440, with 1024 and 2000 captures.

## 10. API contract this spec assumes

The API lane builds exactly this; if it differs, change the contract here first.

**Route:** `GET /v1/characters/{region}/{ruleset}/{name}/history`. Owner-only: the caller must own the character (the same rule the own-character view uses) and `entitlements.Store.Can` must allow `history`. Not the owner: 404 (the route does not reveal the character). Owner without the entitlement: 402 (the web never calls it without `me.entitlements.history === true`). `Cache-Control: private, no-store`. The public character payload and its `score_history` (one point per UTC date, public unflagged reports, newest 30) are **unchanged** and still feed the card's sparkline for every viewer.

**Body:**

```json
{
  "as_of": "2026-12-11",
  "score": [
    {
      "at": "2026-12-11T21:14:03Z",
      "score": 1830,
      "tier": "epic",
      "report_id": "abcd1234efgh",
      "fight_index": 7,
      "encounter": "Golemagg the Incinerator",
      "difficulty": 3,
      "parse_percentile": 88,
      "kill": true,
      "visibility": "public"
    }
  ],
  "encounters": [
    {
      "encounter": "Golemagg the Incinerator",
      "encounter_id": 663,
      "difficulty": 3,
      "points": [
        { "at": "2026-12-11T21:14:03Z", "parse_percentile": 88, "report_id": "abcd1234efgh", "fight_index": 7 }
      ]
    }
  ],
  "gaps": [
    { "from": "2026-10-26", "to": "2026-11-17", "reason": "No rated fights" }
  ],
  "fights": []
}
```

Rules, each one tested by the API lane:

- **`score`:** one entry per rated fight whose `rating_scores` row carries a score snapshot, oldest first, at most the newest 1000. `score` is on the **0 to 2100** scale (`score_v2.go`); `tier` is one of `common, uncommon, rare, epic, legendary` (the rating module's, from the snapshot). `visibility` is the fight's report's visibility: `public`, `guild` or `private` (an unlisted report reads as `public` for this purpose, as it does for the score). Flagged reports are excluded (never a point). `parse_percentile` is omitted when none was computed; `kill` is false for a wipe.
- **`encounters`:** kills only (`kill` true), one entry per boss and difficulty, points oldest first, every boss with at least one kill included, ordered by newest kill first. Same visibility and flagged rules as `score`. A point with no `parse_percentile` is left out.
- **`gaps`:** every stretch of more than **14 days** between two consecutive UTC dates that have a score point, oldest first, `from` and `to` the two dates (`YYYY-MM-DD`, the points either side). `reason` is exactly one of `No rated fights`, `Report flagged`, `Score not kept`, chosen in that precedence: `Report flagged` if any fight in the stretch is in a flagged report; else `Score not kept` if any `rating_scores` row in the stretch has NULL snapshot columns; else `No rated fights`. The web draws these and derives none.
- **`fights`:** the rows behind every point in `score` and `encounters`, newest first, deduplicated by (`report_id`, `fight_index`), each in the shape the public page's `history` rows already use (`CharacterFight`: `encounter`, `encounter_id`, `difficulty`, `metric`, `value`, `percentile`, `spec`, `fought_at`, `report_id`, `fight_index`, `execution_score`) plus `visibility`. The brief's contract did not carry these rows; the lists need amount and execution, so they are added here.
- **Verify before building (the API lane reports back):** that a private or guild fight's snapshot is written at all (`RollUpScores` calls `ScoreV2`, which reads public reports only, so it is the public score as it stood, or NULL for a character with no public report yet); if rows are NULL the owner's score points for those fights are left out, and the ring treatment in 3.5 simply has fewer rings.
- **Keys:** the web keys a fight as `report_id:fight_index` everywhere (a score point, an encounter point, a list row).
- **Difficulty (amended 2026-10-25 by the controller):** `difficulty` is the integer id the public `history` rows already send, never a name; the web names it with the same helper the fight table uses. The API has no difficulty-name table and must not invent one.

## 11. Decisions the product spec left open (for the controller to rule)

1. **Rings for private and guild points on the score chart** (ruling 3). Because the score cannot move on private reports, the alternative is public-only points on the score chart and private kills only in the parse chart. Recommended: rings, because the product spec says the owner's chart includes them and the note makes the reason plain; if the API lane finds those fights have NULL snapshots, this resolves itself.
2. **The gap vocabulary** (ruling 2): three strings, none of which the brief listed except `Report flagged`. Needs the rating lane's agreement; the three words are one constant in the API and one test.
3. **Gap threshold 14 days,** computed by the API, between consecutive UTC dates with a point. A raid-schedule-aware rule (a missed raid week) is not possible from fights alone.
4. **Placement: after the Forever Score card and before Gear** (3.1). Alternative: after Gear, so the page's one action (Gear) stays directly under the score. Recommended: before Gear, because the product spec says directly below the score card and the chart answers the score's own question.
5. **Two columns from 1280, stacked below** (3.2), score left and parse right, with shared rows so the chart tops align.
6. **Picker as tab pills** (the `MyReports.svelte` strip), not a `select`: the member sees every boss and its kill count at once, and the newest-kill order puts the default first so a scrolling row never hides the selection. Cost: 44 px of height. Alternative: a `select` (shorter, hides the counts).
7. **Selection model** (3.8): selecting a point marks its fights in the list; hover only previews. Alternative: the list filters to the selected date. Chosen because a filtered list hides the rest of the record the member came to read.
8. **List caps:** 8 rows each, `Show all N` expands in place up to 100. Alternative: a link to `Every ranked fight` below. Chosen because that table is public-only and excludes private reports, so it cannot be the list behind an owner's chart.
9. **`CharacterFightTable.svelte` gets four optional props** (`heading`, `columns`, `lead`, `selected`), the smallest extension that reuses the shipped row. Alternative: a new row component; rejected as a second row grammar (tenet 10). The page's `Every ranked fight` partly repeats these rows for an owner; accepted, since that table is the public record and this one is the owner's.
10. **The free owner's sentence has no heading and no panel** (3.9), sits directly under the score card, and appears only when the owner has a score. Alternative: a `History` heading over the sentence; rejected as a heading over one line.
11. **The upsell sentence has no shared component yet.** The brief says lane 1 shipped one; the tree has only the string pattern (`billingBlockCopy.seePlans` and the underlined `/premium` link in `Account.svelte`, `RunControl.svelte` and `ServerRunRow.svelte`). This spec writes the sentence as copy plus that link. If lane 1's component exists on its branch, use it and keep the sentence verbatim.
12. **Reserve fallbacks** of 1060 px (1440) and 1670 px (390), and a `data-history` pre-paint hint that needs the cached `/v1/me` to list the character (3.11); a miss costs one shift for an owner whose cache is cold.
13. **The History account row links to `character page#history`** of the member's main character (product spec section 5), with the words `Open on your main character`; `Sparkline only` for free. Needs the account page to know the main character (it does: the header selector's current).
14. **Parse y bounds follow the colour ladder** (0, 25, 50, 75 to 100) rather than min and max of the data (3.5), so a boss does not rescale as its points arrive.
15. **Caps on the payload:** 1000 fights, so the owner of a long record sees the newest 1000 and a chart that starts where the payload does; no `Load more`. Revisit if a real record passes it.
