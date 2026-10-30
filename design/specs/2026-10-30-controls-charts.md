# Controls and charts: one shared dropdown, and charts a raider trusts (experience spec)

Owner, 2026-10-30, on the report's Threat tab:

- "we have a bug where the dropdown disappears over the chart" (fixed in its own lane; §C makes the class of bug impossible).
- "our charts are very basic - should we be using a more advanced chart library?"
- "also the dropdown boxes look like basic html pages"

This spec covers two site-wide upgrades: one shared dropdown (Part A) and one chart system (Part B). It builds on `design/specs/2026-10-29-site-alignment.md` (the Card, the compact head, `quiet-label`, the framed canvas of §3.7.3) and `design/specs/2026-10-28-polish.md` §4.2 (the `Select` recipe) and §4.5 (type: every number is Barlow tabular, field labels are `quiet-label`). Nothing here restates those rules; it cites them.

This spec is also the design review that tenet 10 requires for each new shared component it introduces: the listbox enhancer and its popover core, `TimeSeriesChart`, `ChartLegend`, `ChartTooltip`, `OverviewBrush` and the Players control.

**Fix round 1 (2026-10-30).** Both reviewers held the first draft: `design/reviews/2026-10-30-controls-charts/ux-verdict.md` and `player-verdict.md`. This revision answers every blocker and every should, with the coordinator's rulings. The changes are marked "(fix round 1)" where they land. The main ones:

- Threat is titled and labelled as an estimate (§B12).
- Deaths carry their killing blow and open the Deaths tab.
- The damage chart draws "Top 8 of N", always with the Source player and you, plus a raid median, and the new Players control picks the rest.
- The tank is always on the threat chart.
- The phone layout fits 360 × 640, measured.
- Searchable focus is one rule.
- Every board's numbers come from one data model.

Boards: `design/mocks/2026-10-30-controls-charts/renders/*.png`. Generator: `gen_controls_charts.py` and `render.mjs` in the same folder. The fight on the boards is mock data shaped like a 40-player Horde Ragnaros kill: 5:40 long, with Ragnaros submerged and untargetable from 3:00 to 4:30. The generator builds per-second damage and healing once. Every chip, head figure, tooltip row, share and table row is computed from that model and never typed. Threat on the boards is the engine's own base-1 model. Every string on the boards is this spec's copy, or copy that ships today.

| Board | What it shows |
|---|---|
| `dropdowns-open-1440.png` | The list open on a fine pointer, in four cases. Plain with descriptions: sim Fight style. With icons, searchable and dense: report Source, opened scrolled to the picked player. Searchable: sim item search Source. The no-match row. It also shows the flip above the field, and a right-slot field opening bottom-end. |
| `dropdown-states-1440.png` | The closed field in every state at the three sizes, in the /account trigger's hover, open and pressed treatment; the lead-icon field; the option rows, including dense and multi-select |
| `threat-chart-1440.png`, `threat-chart-2000.png` | Estimated threat, zoomed to 1:30 to 2:40. Top 8 plus the pinned tanks. The tooltip at 2:07 lists the events within ±2 s: the death with its killing blow, the target change and the taunt. Below it are the caveat line and the table, which agrees with the legend. |
| `damage-timeline-1440.png` | Damage done on Ragnaros. Top 8 of 40 plus you, pinned, and the raid median. The submerge shows as flat lines. Also shows the cast and death lane, a legend hover with Only, and a drag in progress. The head figure reconciles with the sum of the 40 players. |
| `choose-players-1440.png` | The Players control open: searchable multi-select with the Compare group |
| `chart-states-1440.png` | Loading, empty, error, a menu over the plot, and a wipe collapsing into one counted marker. Every frame is the same 415 px. |
| `controls-charts-phone-360.png` | The 360 × 640 viewport after a tab change on the Threat tab, with every height in the fold stack (581 px) |
| `legend-sheet-phone-360.png` | The sheet behind the legend's `+N` on a phone |

**Fix round 2 (2026-10-30, build review).** Both reviewers held the built charts (`design/reviews/2026-10-30-controls-charts/build-charts-ux-verdict.md`, `build-charts-player-verdict.md`). The coordinator's rulings, which this spec now follows where they differ from the text below:

- **Who the damage chart draws (supersedes D5's "Top 8 of N").** The default is the **top 4 and the bottom 4** by the window's figure, plus you, the Source player and the raid median. The reason: a top-8 default named only players at or above the median, so the low performers a raid leader looks for after a wipe were never on the chart. The bottom is taken from the metric's own role (damage dealers on Damage done, healers on Healing, everyone on Damage taken), so a healer's small damage never reads as a low performer. The Players control gains a preset row: `Top 4 + bottom 4` (the default), `Top 8`, `Bottom 8`, and `Custom` (lit while the reader's own set is drawn). The url carries `players=top` and `players=bottom`, and a guid list for Custom. The threat chart keeps its top 8 by standing.
- **The Players field never names a limit above the roster.** With 8 or fewer players it reads `All {N}`; otherwise `Top 4, bottom 4 of {N}`, `Top 8 of {N}`, `Bottom 8 of {N}` or `{n} of {N}`.
- **The median stays readable.** It is 2 px at 85 %, and a hovered chip dims every other line but never a median.
- **Every figure says what it is.** The head reads `Raid {n} DPS over the whole fight` (or `over {start} to {end}`), each legend chip carries the unit (`444 DPS`), and the tooltip head reads `{phase} phase · DPS, 5 s average`. The head, the legend and the Players list use one definition, each player's DPS over the head's span, so the 40 figures add up to the head.
- **Threat's title and the Enemy field agree.** With `Every enemy` picked the chart sums each player's estimated threat over every enemy and is titled `Estimated threat, every enemy`; with an enemy picked it is `Estimated threat on {enemy}`. The phone keeps the name and wraps. The tooltip head reads `{phase} phase · estimated threat, {every enemy | enemy}`.
- **Pinned players lead the legend** at every width, after the medians and before `+{n}`. The phone caveat keeps the estimate wording: `An estimate: 1 per damage, 0.5 per heal.` and the link.
- **Phone chart heads wrap**: the title on its row, then each field on its own full-width row.
- **One chart per tab** includes Resources: its rows are the tab's lines, under a column head, with no damage chart above them.
- **Tooltip event order**: by second, then death, target change, taunt, cast.

## 0. Decisions

| # | Decision | Ruling in this spec |
|---|---|---|
| D1 | Chart library | **uPlot** (MIT, canvas, about 21 KB gzipped). It is confirmed over ECharts. Reasons in §B3. |
| D2 | Dropdown engine | **Our own listbox** in `web/src/lib/ui/listbox.ts`, positioned by **`@floating-ui/dom`** (about 4 KB gzipped), shown in the browser's **top layer** through the Popover API (`popover="manual"`). No component library: Bits UI and Melt UI bring their own styling model and a Svelte-only runtime, and the Astro pages need the same control without Svelte. |
| D3 | Phones and coarse pointers | **The native picker**, behind the same styled closed field. The test is `matchMedia('(pointer: fine) and (hover: hover)')`. When it is false, the field stays a native `<select>`. The one exception is the multi-select Players control, which has no native closed form. On a coarse pointer it opens the legend sheet (§B6.8). |
| D4 | CSS customizable select (`appearance: base-select`) | **Not now.** It ships in Chromium only. When Safari and Firefox ship it, the listbox enhancer is deleted and this spec's look moves onto the native element. |
| D5 | Who the damage chart draws (fix round 1) | **"Top 8 of N"** by the window's figure, with N shown. The Source player and your own character are always drawn when they are in the fight. A raid-median line is on by default. The **Players** control (§B5.3) picks anyone else and offers a same-class median. The raid total moves to the chart head as a figure. One player as the Source keeps today's gold area. |
| D6 | Threat (fix round 1) | **"Estimated threat"** everywhere it is a number (§B12). The tank is always drawn. Target changes are marked where the summary carries them. A one-line caveat links to the glossary. |

## Part A. The shared dropdown

### A1. The question and the answer in view

The visitor's question at any dropdown is: "What is picked, and what else can I pick?"

1. Closed, the field answers the first half. It shows the picked option's label, its icon when the option has one, and its class colour when it names a player.
2. Open, the list answers the second half. Every option is one glance away, the picked one carries a gold check, and the list is scrolled so the picked option is in view.
3. The list never needs a page scroll to be seen. It flips above the field when less than 240 px fits below. It is never clipped by a card and never sits under a chart (§C).
4. At 360 px and on any coarse pointer, the phone's own picker answers the second half (D3).

### A2. Reference

| Reference | What we borrow |
|---|---|
| Our header account menu (`.csel-panel` in `web/src/styles/nav.css`, owner-approved with /account) | The list's surface: `--surface-card`, the 1 px `--color-line-card` border, `--radius-panel` 12 px, the inset `--card-highlight` top edge and `--shadow-card`. The 11 px tracked muted group head. The filter field recipe. The bottom fade on a scrolling list. |
| The report's More menu (`ModeBar.svelte`) | A dense list anchored under its trigger in a tool bar, opened in place, one tap to any entry. |
| The WoW client's dropdown (UIDropDownMenu: the talent spec picker, the loot-method menu) | A gold check on the current entry, compact 36 px rows, a dark framed panel over the game. |
| Warcraft Logs' filter pickers (Source, Ability, Target) | Class-coloured player names with their class icon, grouped by role, filterable by typing when the list is long. |

### A3. One component, three layers

1. **`web/src/components/ui/Select.svelte`** stays the one Svelte entry point. Every raw `<select>` in §A10 moves to it. It renders the native `<select>` with every option and its data. That native element is the source of truth for value, form `name`, `bind:value` and `onchange` at every pointer.
2. **`web/src/components/ui/Select.astro`** is its twin for static pages. It renders the same markup with no Svelte runtime.
3. **`web/src/lib/ui/listbox.ts`** is the enhancer. It is framework-free TypeScript. On a fine pointer it finds each field marked `data-listbox`, hides the native `<select>` (the `hidden` attribute, which keeps it in form submission), and inserts the trigger button and the list. Choosing an option sets `select.value` and dispatches `input` and `change` on the select, so every call site's existing handler runs unchanged.
4. The enhancer is loaded with a dynamic `import()` on `requestIdleCallback` after first paint, only on a fine pointer. Until it loads, the native `<select>` works. Islands with a 16 KB budget (Rankings, GuildShell, Character) carry no listbox code in their entry.
5. The closed field is pixel-identical in both modes: same box, same text position, same chevron. The swap is invisible and moves nothing (§A11).
6. **Where the list mounts (fix round 1).**
   - The list element is a child of the field's own wrapper in the DOM, never of `document.body`.
   - It carries `popover="manual"` and is shown with `showPopover()`, so it renders in the top layer above every stacking context and every chart.
   - Inside a modal `<dialog>` or a focus trap (the raid form, the sim settings sheet, guild settings), the list stays inside the dialog's subtree. It is therefore interactive and drawn above the dialog.
   - Escape on an open list closes only the list: the handler calls `preventDefault()` and `stopPropagation()`, so the surrounding dialog or sheet stays open.
7. **Positioning.** `@floating-ui/dom` places the list: `computePosition` with `offset(4)`, `flip()`, `shift({ padding: 8 })`, `size()` and `autoUpdate`. The placement is `bottom-start`. It is `bottom-end` when the field sits in a right slot: a card head's `aside`, the end of a tool bar, a chart head.
8. **No Popover API (fix round 1).** The enhancer checks `typeof HTMLElement.prototype.showPopover === 'function'`. Where it is missing (Safari before 17), the enhancer does not run, and the native `<select>` stays, styled.
9. **Validation (fix round 1).** The hidden `<select>` keeps `required` and the form's own validity. On the select's `invalid` event, the enhancer calls `preventDefault()` (no hidden bubble). The trigger then takes the error state of §A5, with the call site's error copy, and receives focus.

The `Select` props after this spec:

| Prop | Type | Note |
|---|---|---|
| `options` | `SelectOption[]` | `{ value, label, group?, description?, meta?, icon?, colour?, disabled?, you? }` |
| `option.icon` | `{ kind: 'crest' \| 'spec' \| 'spell' \| 'item'; src: string; quality?: Quality }` | `crest` is the ringed HD crest (one crest language). `item` draws a 1 px quality border and the label in the quality's `-text` colour. No call site uses `item` today; the look is fixed here so nobody invents it later. |
| `option.colour` | a token name, for example `--color-class-mage` | The label's colour. Class names only. |
| `option.meta` | `string` | Right-aligned 13 px muted figure, for example `×8` or a player's DPS. |
| `option.you` | `boolean` | Adds the `You` tag after the label: the viewer's own character. |
| `size` | `'default' \| 'compact' \| 'dense'` | Replaces the `compact` boolean. 52, 44 and 36 px (§A4). |
| `searchable` | `boolean` | Adds the filter field (§A4.4). Set by the call site (§A10). |
| `searchPlaceholder` | `string` | Required when `searchable`. |
| `multiple` | `boolean` | The multi-select variant (§A4.5). Used only by the Players control (§B5.3). |
| `loading` | `boolean` | §A5. |
| `emptyLabel` | `string` | Shown in the field when `options` is empty. Required at every call site whose list can be empty. |
| unchanged | `value`, `id`, `name`, `label`, `placeholder`, `disabled`, `error`, `describedBy`, `testid`, `onchange` | |

The native `<option>` carries the rich data as attributes the enhancer reads: `data-icon`, `data-icon-kind`, `data-quality`, `data-description`, `data-meta`, `data-colour`, `data-you`. Groups are real `<optgroup label>` elements, so the native picker groups them too.

### A4. Layout and tokens

#### A4.1 The closed field (`dropdown-states-1440.png`)

The field takes the /account trigger's states (`.csel-btn` in `web/src/styles/nav.css`) so the two read as one family (fix round 1). It keeps an 8 px radius, not the trigger's 999 px pill: a field is an input, and every input on /account is `--radius-control`.

| Part | Value |
|---|---|
| Box | `--radius-control` 8 px, 1 px `--color-line-card` border, ground `rgba(255,255,255,.045)`, padding 0 40 px 0 12 px |
| Height | `default` 52 px (forms). `compact` 44 px (filter bars, card heads). `dense` 36 px (report and sim tool bars), at 768 px and up on a fine pointer only. Below 768 px or on a coarse pointer, `dense` renders at 44 px. |
| Text | 15 px 600 `--color-strong` (`dense`: 13 px 600, padding-left 10 px). One line, ellipsis, the full label in `title`. |
| Lead icon | 20 px, 8 px gap to the label. A crest is circular with its class ring. A spec, spell or item icon is square with `--radius-pill` 4 px and a 1 px `--color-line` border (quality colour for an item). |
| Label colour | the option's `colour` token when set (a class name), else `--color-strong` |
| Chevron | `SelectChevron`, 16 px (14 px in `dense`), `--color-muted`, 12 px from the right edge (10 px in `dense`) |
| Field label | outside the field, `quiet-label` (13 px 600 `--color-muted`, sentence case), above it in forms (6 px gap) and to its left in tool bars (8 px gap), as today |

#### A4.2 The list (`dropdowns-open-1440.png`)

| Part | Value |
|---|---|
| Surface | `--surface-card`, 1 px `--color-line-card`, `--radius-panel` 12 px, `box-shadow: inset 0 1px 0 var(--card-highlight), var(--shadow-card)`, padding 4 px 0 |
| Placement | 4 px from the field: `bottom-start`, or `bottom-end` in a right slot (§A3.7). It flips above when less than 240 px fits below and more fits above. It stays 8 px inside every viewport edge. |
| Width | at least the field's width, at most 420 px; content decides between them |
| Height | at most `min(360px, available space − 16px)`. Past that it scrolls (`overscroll-behavior: contain`) with a 24 px fade to `#0d0f15` at the bottom, and at the top once scrolled. |
| Open position | the picked option is scrolled into the middle of the list, or to the top when it is near the start |
| Dense list | under a `dense` field, rows are 13 px (§A4.3), so the text does not change size between the field and its list |

#### A4.3 Option rows

| Row | Value |
|---|---|
| Plain | min-height 36 px, padding 6 px 40 px 6 px 12 px, 14 px 600 `--color-strong`. Under a `dense` field: 13 px and min-height 32 px. |
| With an icon | min-height 40 px (36 px dense), the 20 px icon, 10 px gap |
| With a description | label as above; description 12/16 px 500 `--color-muted`, wraps, at most 330 px wide, no clamp (the text is the help, and it is never cut). The description is the option's `aria-describedby`, not part of its name (fix round 1). |
| Meta | 13 px 500 `--color-muted`, tabular, right-aligned, 16 px from the label |
| You | the `You` tag after the label: 11 px 700 `.08em` uppercase `--color-gold`, 1 px `rgba(229,185,85,.38)` border, `--radius-pill`, 18 px tall |
| Selected | a 16 px `check` glyph in `--color-gold`, 12 px from the right edge. The label keeps its own colour. |
| Active (keyboard) | ground `rgba(229,185,85,.10)` and a 2 px inset `--color-gold` left edge |
| Hover | the same look as Active. Hover moves the visual highlight only, and never commits (§A6). |
| Disabled | 45% opacity, `aria-disabled="true"`, skipped by arrows and typeahead, a click does nothing |
| Group head | 11 px 700 `.14em` uppercase `--color-muted`, padding 10 px 12 px 4 px. Every group after the first has a 1 px `--color-line-soft` top rule with 4 px above it. A group head sticks to the list's top while its rows scroll. |
| Filter match | the matched run in `--color-gold` 700 |
| Touch on a fine-pointer device | when the last `pointerdown` had `pointerType === 'touch'` (a touchscreen laptop), every row is at least 44 px until a mouse moves again (fix round 1) |

#### A4.4 The searchable list

1. A filter row sits at the top of the list, outside the scrolling region: padding 6 px 8 px 8 px and a 1 px `--color-line-soft` bottom rule.
2. The filter input is 36 px tall, radius 6 px, ground `--color-bg`, 1 px `--color-line` border (`--color-gold` with a 2 px `rgba(229,185,85,.35)` ring while focused), a 16 px `search` glyph at the left, and 14 px text.
3. The count sits at the right of the row: 12 px `--color-muted`, `aria-live="polite"`.
4. Matching is a case- and accent-insensitive substring match on the label. A group keeps its head when any of its rows match.
5. No match shows one 14 px muted row: the copy in §A8.

#### A4.5 The multi-select list (fix round 1; `choose-players-1440.png`)

1. The same list, with a 16 px box at the left of each row. A box is 1 px `--color-line-warm-strong` with radius 4 px when off. When on, it is filled `--color-gold` with a 12 px `--color-bg` check.
2. A row is `role="option"` with `aria-selected` for its on state, inside a listbox with `aria-multiselectable="true"`.
3. A foot row sits under the scrolling region: a 1 px `--color-line-soft` top rule, padding 6 px 8 px, `button('ghost','sm')` `Top 8` at the left and a 12 px muted count at the right (`{n} drawn`).
4. The closed field reads `Top 8 of {N}` while the default holds, and `{n} of {N}` after any change.
5. It is always searchable.

### A5. States

The closed field's hover, open and pressed states are the /account trigger's (fix round 1).

| State | Closed field | List |
|---|---|---|
| Default | as §A4.1 | closed |
| Hover | border `rgba(229,185,85,.45)`, ground `rgba(255,255,255,.06)`, chevron `--color-strong` | the hovered row takes the highlight; nothing commits |
| Focus-visible | 2 px `--color-gold` outline, 2 px offset | where focus lives is §A6.1 |
| Active (pressed) | border `--color-gold`, ground `#101624` | the pressed row keeps the highlight |
| Open | border `--color-gold`, ground `rgba(255,255,255,.06)`, `box-shadow: var(--glow-gold)`, chevron turned 180° in `--color-strong` | shown |
| Disabled | 45% opacity, `cursor: not-allowed`, not focusable, does not open | n/a |
| Error | 1 px `--color-wipe` border; the 13 px `--color-wipe` line under the field, `role="alert"`, linked by `aria-describedby`; `aria-invalid="true"` | unchanged |
| Loading | label `Loading…` in `--color-muted`, a 14 px gold `Spinner` replaces the chevron, `aria-busy="true"`, does not open, keeps its width | n/a |
| Empty (no options) | the call site's `emptyLabel` in `--color-muted`, no chevron, disabled | n/a |
| Placeholder (nothing picked) | the placeholder in `--color-muted` 500 | no row checked |
| Search, no match | n/a | the no-match row |

Touch and keyboard parity: every state above renders the same on the native field, except Open, which is the phone's own picker.

### A6. Keyboard and pointer (fine pointer; WAI-ARIA combobox patterns)

#### A6.1 Where focus lives (fix round 1: one rule)

| | Plain list | Searchable list (and multi-select) |
|---|---|---|
| Pattern | select-only combobox | combobox with a listbox popup and list autocomplete |
| The trigger | `<button type="button" role="combobox" aria-haspopup="listbox" aria-expanded aria-controls aria-labelledby="{field label id} {trigger id}">`. It owns `aria-activedescendant` while open. | `<button type="button" aria-haspopup="listbox" aria-expanded aria-controls aria-labelledby>`. No `role="combobox"`, and it never holds `aria-activedescendant`. |
| Focus while open | stays on the trigger | moves to the filter input on open |
| The filter input | n/a | `<input role="combobox" aria-autocomplete="list" aria-expanded="true" aria-controls="{list id}" aria-activedescendant="{active row id}">`, labelled by the field label |
| The list | `role="listbox"`, `aria-labelledby` the field label | the same, plus `aria-multiselectable="true"` for the Players control |
| On close | focus is on the trigger | focus returns to the trigger |

There is exactly one element with `role="combobox"` per open list. An axe check covers both variants (§D, capture 6).

#### A6.2 Two highlights, one commit rule (fix round 1)

1. The list tracks a **keyboard-active** row: the one the arrows, Home, End, PageUp, PageDown and typeahead moved to. On open it is the picked row.
2. A pointer over a row moves only the **visual highlight**. When the pointer moves, the keyboard-active row follows it; a pointer resting still never changes it.
3. Tab, Enter and Alt+Up commit the keyboard-active row only. A row that is highlighted only because a pointer is parked over it is never committed by Tab.
4. A click commits the clicked row.

#### A6.3 Keys and actions

| Key or action | Closed | Open, plain | Open, searchable |
|---|---|---|---|
| Down, Up, Alt+Down, Enter, Space | open | Down and Up move the keyboard-active row by one (no wrap) | the same, from the filter input |
| Home, End | open, active = first or last | first or last enabled row | move the caret in the filter (the input's own behaviour) |
| PageDown, PageUp | n/a | move by 10 rows | move by 10 rows |
| A printable character | open; plain jumps to the first row starting with the typed run (500 ms buffer); searchable types it into the filter | typeahead | types into the filter |
| Enter | n/a | commit the keyboard-active row and close | single: commit and close. Multi: toggle the active row; the list stays open |
| Space | n/a | commit and close | types a space. Multi: Ctrl+Space toggles the active row |
| Alt+Up | n/a | commit the keyboard-active row and close | the same (multi: close) |
| Escape | n/a | close with no change | with text in the filter: clear it first; otherwise close (multi keeps the toggles made). Escape never reaches a surrounding dialog (§A3.6). |
| Tab | moves on | commit the keyboard-active row (§A6.2), close, move focus on | single: the same. Multi: close, move focus on |
| Click the field | open | close with no change | close |
| Click a row | n/a | commit and close | single: commit and close. Multi: toggle; the list stays open |
| Click outside, or the field scrolls out of view | n/a | close with no change | close |
| Another menu opens | n/a | this one closes (one open overlay at a time, §C) | the same |

Committing fires `change` once. Committing the already-picked row fires nothing. In the multi-select, each toggle fires `change`, and the chart redraws behind the open list.

### A7. Hierarchy

1. In the closed field, the picked label is the one primary element. The chevron and icon are secondary.
2. In the list, the active row is the one emphasis. The selected check is a quieter second signal and never competes with it.
3. Group heads and descriptions are muted. A description never renders brighter than its label.

### A8. Copy (verbatim; new keys in `web/src/lib/ui/copy.ts`, group `select`)

| Key | Copy |
|---|---|
| `loading` | `Loading…` |
| `noMatch(query)` | `Nothing matches “{query}”.` |
| `matchCount(n)` | `1 match` / `{n} matches` |
| `clearFilter` (aria-label of the filter's clear button) | `Clear filter` |
| Report Source filter placeholder | `Filter players` |
| Report Ability filter placeholder | `Filter abilities` |
| Report Target filter placeholder | `Filter targets` |
| Sim item search Source placeholder | `Filter sources` |
| Raid form time zone placeholder | `Filter time zones` |
| Rankings Encounter placeholder | `Filter encounters` |
| Threat Enemy `emptyLabel` | `No enemies in this window` |
| Report Ability `emptyLabel` | `No abilities in this window` |
| Report Target `emptyLabel` | `No targets in this window` |
| Compare with `emptyLabel` | `No other fight to compare` |
| Window preset label for a dragged window | `{start} to {end}` (for example `1:30 to 2:40`) |

Every other label, placeholder and option text is unchanged.

### A9. Phone and desktop

1. **Fine pointer, any width:** the listbox. The field keeps its size class. `dense` is 36 px from 768 px.
2. **Coarse pointer or no hover, any width:** the native picker. The closed field is the same look. `dense` renders 44 px. Options lose their icons and descriptions inside the OS picker; the closed field still shows the picked option's icon and colour.
3. **What never collapses:** the field label, the picked value, the chevron, and the error line.
4. A hybrid laptop with a mouse is a fine pointer. The test is re-run on `change` of the media query, and an open list closes when it flips.

### A10. Every call site

Variant key: **P** plain, **I** with icons, **S** searchable, **D** with descriptions. Size: **52** default, **44** compact, **36** dense.

| # | File | Control (label) | Today | Variant | Size | Notes |
|---|---|---|---|---|---|---|
| 1 | `report/ModeBar.svelte` | Source | raw | I + S | 36 | Crest and class colour per player; the `You` tag on yours. Groups: none for the two scopes, then `Tanks`, `Healers`, `Damage`, `Enemies` (role from the roster). `searchable={roster.length > 15}`. |
| 2 | `report/ModeBar.svelte` | Table (below lg) | raw | P | 36 | |
| 3 | `report/ModeBar.svelte` | More | `<details>` | not a Select | n/a | Moves onto the same popover core (`lib/ui/popover.ts`), surface and 36 px row look, `role="menu"`; current tab carries the check. |
| 4 | `report/ThreatTable.svelte` | Enemy | raw | P | 36 | `meta` is the group count (`×8`) instead of a suffix in the label. Moves into the threat chart's head (§B5). |
| 5 | `report/TimeChart.svelte` | Window | raw | P | 36 | Moves into the main chart's head. A dragged window reads `{start} to {end}`. |
| 6 | `report/FilterBar.svelte` | Target | raw | P + S | 36 | `searchable={targets.length > 15}` |
| 7 | `report/FilterBar.svelte` | Ability | raw | I + S | 36 | Spell icon per ability. Always searchable. |
| 8 | `report/CompareMode.svelte` | Compare with | raw | D | 36 | Description `{Kill or Wipe} · {duration}` from the fight list already shown in FightSelector. |
| 9 | `report/CompareMode.svelte` | vs | raw | I | 36 | Crest and class colour |
| 10 | `report/CompareMode.svelte` | Metric | raw | P | 36 | |
| 11 | `report/CompareMode.svelte` | Phase | raw | P | 36 | |
| 12 | `report/RankingsMode.svelte` | Metric | raw | P | 36 | |
| 13 | `report/RankingsMode.svelte` | Spec | raw | I | 36 | Spec icon (`lib/bis/spec-icon.ts`) |
| 14 | `rankings/FilterSelect.svelte` (9 filters in `RankingsFilters.svelte`) | Class, Spec, Encounter, Ruleset, Region, Phase, Faction, Metric, Board | own select | Class I (crest), Spec I (spec icon), Encounter P + S when over 15, the rest P | 44 | `FilterSelect` is deleted; its callers use `Select size="compact"`. |
| 15 | `sim/SettingsBar.svelte` | Fight style | raw | D | 44 | Descriptions are `simCopy.fightStyleOptions`; the separate `<dl>` key under the bar is deleted, since the list now carries it. |
| 16 | `sim/SettingsBar.svelte` | Fight length, Targets, Buffs preset | raw | P | 44 | |
| 17 | `sim/SettingsSheet.svelte` | Variation, Target level, Target type | raw | P | 44 | |
| 18 | `sim/PrecisionSelect.svelte` | Precision | raw | P | 44 | The component stays as a thin wrapper over `Select`. |
| 19 | `sim/CharacterStrip.svelte` | Race | raw | P | 44 | |
| 20 | `sim/CooldownRows.svelte` | cooldown mode per row | raw | P | 36 | |
| 21 | `sim/BuffPanel.svelte` | grade per buff | raw | P | 36 | |
| 22 | `sim/tools/ItemSearch.svelte` | Slot | raw | P | 44 | |
| 23 | `sim/tools/ItemSearch.svelte` | Source | raw | P + S | 44 | Groups are `SOURCE_KIND_LABELS` (Raids, Dungeons, World bosses and the rest). Always searchable. |
| 24 | `sim/SimHistory.svelte` | Show | Select | P | 44 | |
| 25 | `planner/PlannerHeadControls.svelte` | Class | Select | I | 44 | Crest and class colour |
| 26 | `planner/PlannerHeadControls.svelte` | Race | Select | P | 44 | |
| 27 | `planner/SummaryBar.svelte` | Class | raw | I | 44 | Crest and class colour (mounted by `Planner.svelte`) |
| 28 | `ClassHeader.astro` | Race (BiS) | raw | P | 44 | `Select.astro` |
| 29 | `CharacterSearch.astro` | Region, Ruleset | Select | P | 52 | `Select.astro` |
| 30 | `CharacterNotFound.svelte` | Region, Ruleset | raw | P | 52 (44 below 768) | |
| 31 | `AddonPasteSave.svelte` | Region, Ruleset | raw | P | 44 | |
| 32 | `GuildLanding.svelte` | Region, Ruleset | Select | P | 44 | |
| 33 | `GuildSettings.svelte` | Default report visibility | Select | P | 44 | |
| 34 | `guild/RaidForm.svelte` | Raid | raw | P | 52 | |
| 35 | `guild/RaidForm.svelte` | Duration | raw | P | 52 | |
| 36 | `guild/RaidForm.svelte` | Guild time zone | raw | P + S | 52 | Groups are today's region optgroups. Always searchable. |
| 37 | `guild/SignupSpec.svelte` | Spec | Select | I | 44 | Spec icon |
| 38 | `guild/GuildLoot.svelte` | Boss | Select | P | 44 | |
| 39 | `guild/billing/GuildSeatPicker.svelte` | seat holder | raw | P | 52 | |
| 40 | `guild/billing/GuildWhoPays.svelte` | Transfer to | raw | P | 52 | |
| 41 | `account/AccountGuildBlock.svelte` | Share | Select | P | 44 | |
| 42 | `ui/UiSheet.svelte` | the gallery | Select | all | all | Gains one example of every variant and state in §A5, which the e2e and screenshot checks read. |
| 43 | `chart/ChartLegend.svelte` (new; in every report chart head) | Players | new | I + S, `multiple` | 36 | §A4.5 and §B5.3. Opens bottom-end. On a coarse pointer it opens the legend sheet instead (D3). |

After this lane, `grep -rn '<select' web/src --include=*.svelte --include=*.astro` returns only `Select.svelte` and `Select.astro`. A unit test enforces it.

### A11. Performance and polish

1. The swap from native to listbox moves nothing: the wrapper has a fixed height per size, and an e2e test compares the field's bounding box before and after the enhancer loads (0 px difference).
2. The enhancer and Floating UI load as one lazy chunk, under 9 KB gzipped (`check-island-size.mjs` gains the entry).
3. Icons in the list never flash. Icons of the first 12 rows are requested on `pointerenter` or `focus` of the field. Each icon slot is a fixed 20 px box with a `--color-raised` ground until it paints.
4. Reduced motion: the list appears with no transition and the chevron turns with no transition. Otherwise the list fades and rises 2 px over 100 ms, and the chevron turns over 120 ms.
5. Lighthouse budgets in `web/lighthouserc.json` are unchanged and must hold on every route that renders a `Select`.

## Part B. Charts

### B1. The question and the answer in view

A raider's questions at a chart are the three after a wipe: "Who died, and why?", "Who was low, and when?" and "Who held the boss?"

1. The plot answers "when". Lines read at a glance, and markers name the moments: a death, a taunt, a boss target change, a cooldown, a phase.
2. A death answers "why" in place. Its marker and the tooltip carry the killing blow (spell, source, amount, overkill), and the marker opens that death on the Deaths tab (§B6.6).
3. "Who was low" is answered against a reference. The damage chart draws the top 8, you and the Source player, against the raid median (§B5.3). The Players control adds anyone, and a same-class median.
4. "Who held the boss" is answered honestly. The threat chart is an estimate and says so. It always draws the tank and every player the boss turned to (§B12).
5. **Desktop, 1280 × 800, after the Threat tab change scrolls to the mode bar.** The chart head, the legend, the whole plot, the overview brush and the caveat line are visible.
6. **Phone, 360 × 640, after a tab change (fix round 1, measured).**
   - On a phone the tab change scrolls the **Source row** to the top, 8 px down. Today it scrolls to the mode bar.
   - Measured on the built fixture report at 360 × 640 (`/reports/fixture2abcd?fight=3&tab=threat`, 30 Oct), the mode bar is 252 px tall: the mode strip, two rows of view pills, Source and Table. The site header (56 px, `position: relative`) is not sticky and has scrolled away. Today the threat chart starts 494 px below the mode bar's top, and its plot is below the fold.
   - The new fold stack, every value fixed by this spec:

| Part | px |
|---|---|
| scroll margin | 8 |
| Source row | 44 |
| gap | 8 |
| Table row | 44 |
| gap to the chart | 12 |
| chart head (`Estimated threat` and the window) | 48 |
| Enemy field row (44 + 8) | 52 |
| legend, one row (36 + 1 rule) | 37 |
| marker lane | 26 |
| plot | 160 |
| x axis | 22 |
| event line | 28 |
| brush (6 + 32 strip + 16 labels + 8) | 62 |
| caveat line | 28 |
| frame borders | 2 |
| **total** | **581** |

   - The chart ends at 581 px, 59 px above the fold. The first table row starts in view.
   - What makes it fit:
     - The threat note above the chart moves under it, as the caveat line.
     - The legend is one row.
     - The plot is 160 px.
     - The scrub hint is not in the first view: it is the empty event line's text until the first scrub (§B6.8).
7. One chart per tab. On the Threat tab the threat chart is the tab's chart. The main chart, and its collapsed row, are not drawn there. Today both draw.
8. The main chart on a phone stays collapsed by default (`fs.report.chartOpen`, today's rule) on every other tab, so a table's answer stays first.

### B2. Reference

| Reference | What we borrow | Why ours is as good or better (tenet 12, fix round 1) |
|---|---|---|
| Warcraft Logs' graph (Damage Done, Threat) | One line per player in class colours. A crosshair tooltip listing every player at that time, sorted. Drag across the graph to zoom, and the tables follow the range. Death and cast markers on the time axis. | Ours adds three things WCL's graph lacks. The legend's `Only` isolates a class-coloured line in one click. The crosshair is synced across the chart and the timelines on the same tab. A death marker carries its killing blow and opens the death. We match WCL's defaults (top N, class colours) so nothing has to be relearned. |
| Details! (the timeline and the death log) | Event icons on the time line, the spell's own icon on every cast, the killing blow on every death | The same death line, on the chart where the raid's damage dropped, one click from the full death recap |
| The game's own combat text colours and the class colours | Class colours on every player, the death red, the gold for the player's own selection | n/a |
| Our framed canvas (site alignment §3.7.3) | The chart sits in a `card-surface-quiet` frame with a compact head. The head has no IconTile. | n/a |

### B3. The library: uPlot, confirmed

**Ruling: uPlot.** The reasons, against the report's facts:

1. **Fit for the data.** A report chart is up to 40 series of one point per second, about 250 to 900 points each. uPlot is built for exactly this: canvas, linear-time drawing, about 1 ms per redraw at this size. Hover stays at 60 fps with 40 lines.
2. **Weight.** uPlot is about 21 KB gzipped. It loads as a lazy chunk (§B11), so `report-island.js` (140 KB cap in `check-island-size.mjs`, about 129 KB today) does not grow at all.
3. **What it gives us.** Cursor sync between charts (`cursor.sync.key`), drag selection (`cursor.drag`, `hooks.setSelect`), a log scale (`distr: 3`), draw hooks for markers, and `setSize`, `setScale`, `setSeries`. That covers every interaction in §B6.
4. **What it does not give us, and we want to own anyway.** Its legend and tooltip are basic. We hide its legend (`legend.show: false`) and draw our own legend and tooltip in the DOM with our tokens. Its touch handling is minimal. We own the long-press scrub (§B6.8).
5. **ECharts rejected.** A tree-shaken ECharts with a line chart, tooltip, dataZoom and markLine is several times uPlot's weight (well over 100 KB gzipped). Its parse and init cost lands inside the report's 200 ms Total Blocking Time budget on Lighthouse's 4× CPU slowdown. Its tooltip, legend and dataZoom would all be restyled or replaced to match the site, so most of what the weight buys goes unused.
6. **Chart.js rejected.** It is about three times uPlot's weight with its zoom plugin. It is slower with dense multi-line series.
7. **SVG (hand-rolled or LayerCake) rejected for the dense charts.** 40 paths of 900 points each, re-rendered on every hover, is slow on a phone. SVG stays where a chart is sparse (§B10).

The lane measures the gzipped chunk on install. A chart chunk (uPlot plus our wrapper) over 30 KB gzipped stops the lane for a ruling.

**Ruling (2026-10-30, coordinator): the cap is 32 KB (32,768 B).** The chart lane measured the chunk at 31,697 B gzipped (level 9). uPlot 1.6.32 alone is about 22.5 KB, and the plot, the tooltip and the brush add about 9 KB. Moving the tooltip and the brush into the entry measured 28,957 B for the chunk, but Rollup then split a 19 KB shared chunk out of `report-island.js`, so the page would load no less. The budget is raised instead of moving files to pass the check.

### B4. Components

| Component | Role |
|---|---|
| `web/src/components/chart/TimeSeriesChart.svelte` (new, shared) | Owns the uPlot instance, the marker lane, the crosshair, the drag selection, and the keyboard and touch handling. Props: `series`, `markers`, `phases`, `domain` (the whole fight in ms), `window`, `onWindow`, `plotHeight`, `scale` (`'linear' \| 'log'`), `syncKey`, `mode` (`'rate' \| 'cumulative'`), `label`, `emptyCopy`. |
| `web/src/components/chart/ChartLegend.svelte` (new) | The chip row and the Players control (§B5.3) |
| `web/src/components/chart/LegendSheet.svelte` (new) | The phone sheet behind `+{n}` (§B6.8), on the account selector's sheet recipe |
| `web/src/lib/chart/drawn.ts` (new) | Who is drawn: top N, the always-drawn set, the medians (§B5.3). Pure and unit-tested. |
| `web/src/components/chart/ChartTooltip.svelte` (new) | The tooltip (§B5.6), rendered in the top layer |
| `web/src/components/chart/OverviewBrush.svelte` (new) | The whole-fight strip with the window and its two handles (§B5.7). It holds the two native range inputs that `TimeChart` draws today. |
| `web/src/lib/chart/time-axis.ts` (new) | Tick steps and the `m:ss` / `h:mm:ss` formatter, shared by every time axis on the site |
| `web/src/lib/chart/cursor.svelte.ts` (new) | The page's shared cursor: `{ syncKey, ms } \| null` |
| `web/src/lib/chart/palette.ts` (new) | Series colours (§B5.5), read from tokens through `getComputedStyle`, never a hex in code |
| `web/src/lib/chart/uplot-loader.ts` (new) | The dynamic `import('uplot')`, started when the report island mounts |
| `web/src/lib/ui/overlay.svelte.ts` (new, shared with Part A) | The count of open overlays (§C) |
| `report/TimeChart.svelte` (rewritten) | The report's adapter: summary to series, markers and presets. It keeps its name and its callers. |
| `ui/Card.svelte` (exists) | The frame: `quiet`, `compact`, `body="none"`, the head's `aside` slot for the chart's controls |

### B5. Layout, region by region (`threat-chart-1440.png`, `damage-timeline-1440.png`)

All values are at 1440. The chart fills the page column (1248 px content at 1440 and at 2000; the column is centred at 2000 and the chart never grows past it).

#### B5.1 Frame

`Card quiet compact body="none"`: `card-surface-quiet`, 1 px `--color-line-card`, `--radius-panel` 12 px, `--shadow-card`, no inset highlight (site alignment §3.7.3).

#### B5.2 Head (compact head, site alignment §3.3.1)

| Part | Main chart | Threat chart |
|---|---|---|
| Title (`section-title` 18 px uppercase) | the table's metric: `Damage done`, `Healing`, `Damage taken` | `Estimated threat on {enemy}` (`Estimated threat` on a phone, where the Enemy field under it names the enemy) |
| Meta (13 px `--color-muted`, tabular) | `Raid {n} DPS · 5 s average`, with ` · on {target}` when the Target filter is set (HPS for Healing, DTPS for Damage taken). With one player as the Source: `{name} · {n} DPS · 5 s average`. | `Cumulative · {start} to {end}` |
| Right slot, left to right | `quiet-label` `Players` + the Players control (dense); `quiet-label` `Window` + `Select` dense (presets); `Switch` `Log scale` last, so it stays out of the head's first read | `quiet-label` `Players` + the Players control; `quiet-label` `Enemy` + `Select` dense; then, when zoomed, `button('ghost','sm')` with the `x` glyph: `Whole fight` |

The raid figure in the head is the sum of every player's figure as the legend and the Players list show it, rounded per player first. A raider who adds the 40 players' figures gets the head's number (`damage-timeline-1440.png` prints the check).

The main chart shows `Whole fight` as the Window select's first option, so it has no separate ghost button.

#### B5.3 Legend (`ChartLegend`) and who is drawn (fix round 1)

**Who is drawn.**

1. By default, the chart draws the **top 8** players by the window's figure. For damage done, healing and damage taken that figure is DPS, HPS or DTPS over the window. For threat it is the standing at the window's end. On the Healing tab the top 8 are therefore the healers.
2. **Always drawn**, on top of the 8, whenever they are in the fight:
   - the Source player, when the Source is one player;
   - your own character, when you are signed in and it is in the fight;
   - on the threat chart, every player the enemy attacked during the window, including the one it was on at the window's start;
   - on the threat chart, every player who taunted it during the window.
3. **The raid median** is drawn on the rate charts by default. It is the per-second median of the 5 s averages of the players in the Damage role (Healing: the Healer role; Damage taken: every player). It is a 1.5 px dashed (5/4) `--color-text` line at 70%. It is drawn under the player lines.
4. Sized for 40: the default never draws more than 8 plus the always-drawn players (at most 12 lines in practice). The legend wraps to two rows at 1248 px.

**The legend row.**

1. A row under the head, with a 1 px `--color-line-soft` top rule, padding 8 px 12 px 6 px and gap 4 px 6 px, wrapping. It reserves its height before data arrives (§B11).
2. A chip is 28 px tall (44 px hit area on a coarse pointer), padding 0 8 px, `--radius-pill`. It holds:
   - a 12 × 3 px swatch in the series colour;
   - the name in 13 px 600 in its class colour;
   - a value in 13 px 500 `--color-muted`, tabular, at a fixed 44 px minimum width so a changing value never reflows the row.
3. An always-drawn chip carries a 12 px `pin` glyph in `--color-muted` before the name. Your chip carries `YOU` (10 px 700 `.08em` `--color-gold`) after the name.
4. The raid median's chip comes first, with a dashed swatch: `Raid median` and its window value.
5. Player chips follow, ordered by the window's figure, highest first. The order is recomputed when the window changes, never on hover.
6. A player who died shows the value at death.
7. A chip turned off is at 40% opacity with a hollow swatch.
8. The row ends with `button('ghost','sm')` `+{n} more`, where n is the players not drawn. On a fine pointer it opens the Players control. On a coarse pointer it opens the legend sheet (§B6.8).

**The Players control.** The multi-select of §A4.5 in the head's right slot (bottom-end), searchable (`Filter players`).

1. Its groups are, in order:
   - `Compare`: `Raid median` (on by default) and `{Class} median`. The class is the Source player's when the Source is one player, else yours. Its description is `Median DPS of the {n} damage-dealing {class plural}`. It draws a dashed line in that class's colour.
   - `Tanks`, `Healers` and `Damage`: each player with crest, class-coloured name, the window's figure as meta and the `You` tag on yours.
2. An always-drawn player's row is on and disabled, with the description `Always drawn`.
3. `Top 8` in the foot restores the default.

#### B5.4 Marker lane and plot

| Part | Value |
|---|---|
| Marker lane | 26 px above the plot, drawn only when the chart has markers |
| Plot height | Main chart: 220 px from 1440, 200 px from 1024, 160 px from 768, 140 px below. Threat chart: 300 px from 1440 (the owner works at 2000), 260 px from 1024, 220 px from 768, 160 px below (§B1.6). Sample report strip (`compact`): 96 px, no legend, no lane, as today. |
| Y gutter | 52 px (40 px below 768); labels 11 px `--color-muted` tabular, right-aligned 8 px from the plot: `0`, `400`, `1.2k`, `60k`, `1.2M`. The top tick is the smallest of 1, 1.5, 2, 2.5, 3, 4, 5 or 10 times a power of ten that holds the drawn series. |
| Gridlines | 4 or 5 horizontal, 1 px `--color-line-soft`; no vertical gridlines |
| X axis | 1 px `--color-line` baseline, 4 px ticks, labels 11 px `--color-muted` tabular `m:ss` (`h:mm:ss` from one hour). The step is the smallest of 5, 10, 15, 30, 60, 120 or 300 s that keeps labels 64 px apart. The first label is left-aligned and the last right-aligned. |
| Right padding | 12 px |
| Lines | 1.5 px, round joins. Rate charts draw a 5 s trailing average with straight segments. Cumulative charts draw each second. A player's line stops at their death with a 3 px hollow end dot. |
| Single series | a gold area: `--color-gold` at 20% fill under a 1.5 px `--color-gold` line (today's look). A charted ability adds one line in the palette's second colour. |
| Phase bands | every second phase is filled `rgba(229,185,85,.05)`. A phase start is a 1 px dashed (2/3) `--color-gold` line at 70%. The phase's name from the summary (`Ragnaros`, `Submerge`) sits 6 px right of its start, 4 px from the plot's top, on a `#0d0f15` 90% backing (radius 3, 16 px tall), so no marker line crosses it (fix round 1). Text is 11 px 700 `.12em` uppercase `--color-muted`. |
| A phase where the boss takes no damage | the lines fall to zero and stay flat there; nothing is interpolated across it (`damage-timeline-1440.png`, the submerge) |

#### B5.5 Series colours

1. A player's line is their class colour token (`--color-class-*`).
2. Two players of one class share the colour, as on Warcraft Logs. The tooltip and legend name them, and hover isolates them.
3. A median line is dashed (5/4) at 1.5 px: the raid median in `--color-text` at 70%, a class median in that class's colour at 70%.
4. Enemies and abilities take the palette in order: `--color-gold`, `--color-school-frost`, `--color-school-nature`, `--color-school-arcane`, `--color-school-fire`, `--color-school-shadow`. A seventh series repeats the palette with a 4/3 dash.

#### B5.6 Markers

| Marker | Lane | Plot | Link |
|---|---|---|---|
| Death | a 20 px circle in `--color-death-soft` with a 1 px `--color-death` ring at 60% and the 14 px `skull` glyph in `--color-death` | a 1.5 px dashed (4/3) `--color-death` line, full height, for every death | the lane marker is an `<a>` to that death on the Deaths tab (§B6.6) |
| Boss target change (threat chart; fix round 1) | a 20 px tile: `--color-raised` ground, 1 px `--color-text` border at 50%, radius 4 px, the 14 px `target` glyph in `--color-text` | a 1 px dotted (1/3) `--color-text` line at 35% | none |
| Taunt (threat chart) | the taunt spell's own 18 px icon in a 20 px tile: `--color-raised` ground, 1 px `--color-gold` border at 70%, radius 4 px | a 1 px `--color-gold` line at 35%, full height | none |
| Cooldown cast | the spell's own icon, same tile as a taunt | the same 1 px gold line at 35% | none |
| Phase change | none (the band and its label carry it) | as §B5.4 | none |

**Which casts are marked.**

1. The casts `isRaidCooldown` names (`lib/report/raid-cooldowns.ts`). With one player as the Source, that player's casts from the same list.
2. Fix round 1 adds a short list of life-saving casts a raider looks for beside a death: Rebirth, Soulstone Resurrection, Power Infusion, Mana Tide Totem, Major Healing Potion and Healthstone.
   - These are matched by name, the way `RAID_COOLDOWNS` already is.
   - Each name's spell is checked against the client's spell table before it ships, and the check is recorded in the lane report (tenet 8).
   - Classic names only. Nothing the fight never saw is drawn.

**Collisions.**

1. A marker within 22 px of another steps 22 px right, with a 1 px `--color-gold` leader at 60% back to its time.
2. Three or more deaths inside 44 px (a wipe) collapse into one 20 px `--color-death-soft` circle with a 1 px `--color-death` ring and the count in 11 px 700 `#ffb4b0`. Each death keeps its own dashed line in the plot.
3. The collapsed circle's tooltip lists every death in time order: 6 rows, then `+{n} more · open on the Deaths tab`.
4. Clicking the circle opens the Deaths tab with the window set to the cluster's span (`chart-states-1440.png`, the wipe).
5. Three or more of any other markers inside 44 px collapse the same way, with a gold ring.

The `skull`, `target` and `pin` glyphs are added to `lib/ui/glyphs.ts` (24 grid, 2 px stroke, round caps; paths in the generator).

#### B5.7 Tooltip (`ChartTooltip`)

| Part | Value |
|---|---|
| Box | 268 px wide (300 px on the threat chart), ground `#0d0f15`, 1 px `--color-line-card`, radius 10 px, `inset 0 1px 0 var(--card-highlight)` plus `0 16px 38px rgba(0,0,0,.6)`, padding 8 px 0 6 px, `pointer-events: none` |
| Header | the time in 14 px 700 `--color-strong` tabular; at the right, 12 px `--color-muted`: `{phase} phase · 5 s average`, or `{phase} phase · estimated threat`. A 1 px `--color-line-soft` rule under it. |
| Events (fix round 1) | every marker within **±2 s** of the hovered second, deaths included, in time order. Each is one 12 px line: its time in `--color-muted` (30 px column), its 16 px icon, its text (§B8). A death's line is `#ffb4b0` and carries the killing blow. |
| Column heads (threat only) | 11 px `--color-muted`, right-aligned over the value and share columns: `Estimated threat`, `Share` |
| Rows | 22 px each: a 3 × 12 px swatch, the name in 13 px 600 in its colour, the value in 13 px 600 `--color-strong` tabular at the right. The threat chart adds a 50 px share column (12 px `--color-muted`): the player's share of every player's estimated threat on that enemy at that second. It is the table's `Share of estimated threat`, never a share of the top value (fix round 1). |
| Order | highest value first. A legend-hovered series is always included and its row has a `rgba(229,185,85,.08)` ground. Players who died before this second come last, in `--color-muted`, with `Died {m:ss}` as their value. |
| Limit | 12 rows, then one 12 px muted line: `+{n} more` |
| Placement (fix round 1) | 16 px right of the crosshair while the cursor is in the left 45% of the plot. 16 px left of it once the cursor is in the right 55%, so the lines being read toward stay visible. It is top-aligned with the plot and stays 8 px inside the viewport. It renders in the top layer (`popover="manual"`), so a short plot never clips it. |
| Shown | while a fine pointer is over the plot or a lane marker, or the plot has keyboard focus |
| Hidden | while dragging; while any overlay is open (§C); when the pointer leaves the plot; on a coarse pointer (the legend and the event line read out instead, §B6.8) |

#### B5.8 Crosshair

A 1 px dashed (3/3) `--color-text` line at 70%, full plot height. A 7 px dot in the series colour with a 1.5 px `--color-bg` ring sits on every visible line. A time pill sits on the x axis: 40 × 17 px, radius 4 px, ground `--color-text`, 11 px 700 `--color-bg` tabular.

#### B5.9 Overview brush (`OverviewBrush`)

1. A 36 px strip (32 px below 768) under the plot (padding-top 6 px), across the plot's width, radius 6 px, ground `rgba(255,255,255,.025)`, 1 px `--color-line-soft` border.
2. It always shows the whole fight. Its shape is the raid's damage per second, a `--color-muted` area at 22% with a 50% line, on every chart, so the strip reads the same everywhere.
3. Phase starts are 1 px dashed gold lines at 50%. Deaths are 1.5 px `--color-death` lines.
4. Outside the window the strip is dimmed by `rgba(7,9,13,.6)`. The window is `rgba(229,185,85,.10)` with a 1 px `--color-gold` border at 80%.
5. Each window edge has a handle: 8 × 24 px (12 × 24 px on a coarse pointer), radius 3 px, `--color-gold`, two 1 px `#0b0e14` grip lines; the hit area is 24 × 36 px (44 × 44 px on a coarse pointer).
6. Labels under the strip, 11 px tabular: `0:00` at the left and the fight's end at the right in `--color-muted`, and the window's start and end in 700 `--color-gold` under their handles when zoomed.
7. The handles are the two native `<input type="range">` elements that `TimeChart` already uses, layered and styled. They are the keyboard and screen-reader path to the window.

### B6. Interaction model

#### B6.1 Hover

The pointer over the plot moves the crosshair to the nearest second and shows the tooltip. Moving off the plot hides both.

#### B6.2 Drag to zoom, kept in one window with the brush and the URL

1. A press and drag across the plot draws the selection: `rgba(229,185,85,.13)` with 1 px `--color-gold` edges, and a label centred above it inside the plot: `{start} to {end} · {n} s`, 12 px 600 `--color-strong` on `#0d0f15` with a 1 px gold border at 60%, radius 6 px.
2. On release, a selection of one second or more calls the same `setWindow` that the presets and the brush call. A shorter one is a click and changes nothing.
3. `setWindow` writes `start` and `end` (ms) to the URL as it does today, re-scopes every table, and sets the x scale of every chart on the page to the window.
4. The brush shows the new window. Dragging a handle, or dragging the window body to pan, calls `setWindow` on release.
5. Back to the whole fight: the Window select's `Whole fight`, the threat chart's `Whole fight` button, a double-click on the plot, or Escape while the plot has focus.
6. No zoom is ever animated.

#### B6.3 Cursor sync

1. Every chart and lane view on one page shares the cursor through `lib/chart/cursor.svelte.ts`, keyed by fight.
2. uPlot charts sync with `cursor.sync.key`. `TimelinesView`, `RaidCooldowns` and `ResourceGraphs` read the store and draw the same crosshair at the same ms.
3. Only the view under the pointer shows a tooltip. The others show the crosshair and their dots.

#### B6.4 Legend

| Action | Result |
|---|---|
| Click a chip | toggles that series. The y scale refits to the visible series. An always-drawn chip cannot be turned off; its click does nothing and its tooltip says `Always drawn`. |
| Hover or focus a chip (fine pointer) | that line draws at 2.5 px; every other line drops to 25% opacity; its tooltip row is pinned. An `Only` text button (12 px 600 `--color-gold`) appears in the chip. |
| Click `Only`, or Alt-click a chip | shows that series alone (the always-drawn players stay) |
| `+{n} more` | opens the Players control (fine pointer) or the legend sheet (coarse) |
| Keyboard | each chip is a `button` with `aria-pressed`. `Only` is a second button inside the chip: Tab reaches the chip, then its `Only`. `Only` takes the 2 px `--color-gold` focus-visible outline at 2 px offset, the chip's own outline hides while `Only` has focus, and `Only` stays visible while it has focus (fix round 1). |
| Touch | a tap toggles. The legend sheet carries isolation as Switches. |

The legend's toggles last for the page view and are not written to the URL. The window, the log scale and the Players choice are written to it (`players=` as a comma list of guids, absent for the default).

#### B6.5 Log scale

1. A `Switch` labelled `Log scale` in the head of Damage done, Healing and Damage taken. Threat is cumulative from zero and has none.
2. On: y ticks at powers of ten (`10`, `100`, `1k`, `10k`), a floor of 1, and a zero drawn at the floor.
3. The URL carries `scale=log`, so a shared link reproduces it.

#### B6.6 Markers (fix round 1)

1. Hovering a lane marker shows the tooltip at the marker's second, with its ±2 s events.
2. A death marker is a link to that death on the Deaths tab: `?fight={id}&tab=deaths&death={guid}-{at_ms}`, the key `DeathsTab.svelte` already opens and scrolls to (`state.openDeaths`). It keeps the fight and the window.
3. The marker's hit area is 24 × 26 px, or 44 × 44 px on a coarse pointer. Enter opens it.
4. Its accessible name is the death's event line (§B8).
5. A collapsed death cluster links the same way, with the window set to the cluster's span.
6. Taunt, target-change and cast markers are not links.

#### B6.7 Keyboard (the plot is focusable)

1. The plot is `tabindex="0"`, `role="group"`, `aria-roledescription="chart"`, with the `aria-label` in §B8. Focus draws a 2 px inset `--color-gold` ring.
2. Left and Right move the crosshair one second; with Shift, ten. The tooltip follows.
3. A polite live region reads the time and the top three series at that second, 300 ms after the last key.
4. Escape returns to the whole fight when zoomed.
5. The window is set from the keyboard with the brush's range inputs: arrows ±1 s, Shift+arrows ±10 s, PageUp and PageDown ±30 s, Home and End to the fight's bounds.

#### B6.8 Touch (`controls-charts-phone-360.png`, `legend-sheet-phone-360.png`)

1. The plot has `touch-action: pan-y`, so a swipe scrolls the page.
2. A touch held for 300 ms starts a scrub. The crosshair follows the finger. The crosshair's time pill on the x axis carries the time. There is no time pill in the legend, so no chip ever moves (fix round 1).
3. While scrubbing, the legend reads out. Each chip's value is the value at the finger. Chips keep their order.
4. The legend on a phone is one row:
   - the chips that fit, in the legend's order;
   - then `+{n}` (a `chip more` in `--color-gold`, right-aligned), where n counts the drawn players whose chips did not fit.
5. Tapping `+{n}`, or the Players control on a coarse pointer, opens the **legend sheet** (fix round 1). The sheet is a bottom sheet on the account selector's phone sheet recipe (`nav.css` popover and sheet), as a `<dialog>`:
   - It has a scrim, a grab handle and a max height of 80% of the viewport.
   - The title is `Players` with the meta `{n} of {N} drawn`. The head holds `button('ghost','sm')` `Top 8` and `Done`.
   - Then the `Filter players` field.
   - Then two groups, `Drawn` and `Not drawn`. Each row is 52 px: crest 24, the class-coloured name, the pin when always drawn, the window's value, and a `Switch`. An always-drawn row's Switch is on and disabled.
   - The `Compare` medians are the first two rows of `Drawn` when they are on.
   - Done, a tap on the scrim or a swipe down closes it. Focus returns to the chip that opened it.
6. The 28 px event line under the plot is reserved at all times on a coarse pointer.
   - While scrubbing, it names the first event within ±2 s of the finger, then `+{n}` for the rest.
   - A death reads `{name} died · {spell} · {amount}` on a phone. The source and overkill are on the Deaths tab, one tap away.
   - A line about a death is a link (with a 14 px chevron) to that death on the Deaths tab.
   - Before the first scrub, it holds the hint `Touch and hold the chart to read a moment.` in `--color-muted`. After the first scrub the hint never shows again (`localStorage` `fs.chart.scrubHint`).
7. Lifting the finger keeps the last reading until the next tap outside the plot.
8. Zoom on touch is the brush handles and the Window select. The plot itself takes no drag-to-zoom on touch.

#### B6.9 Reduced motion

There is no zoom animation at any setting. Under `prefers-reduced-motion: reduce`, the tooltip appears with no fade (otherwise an 80 ms fade) and the loading skeleton does not shimmer.

### B7. States (`chart-states-1440.png`)

Every state renders inside the same reserved frame height as Ready (fix round 1). At 612 px wide with a 200 px plot, every frame on the board is 415 px:

| Part | px |
|---|---|
| head | 48 |
| rule | 1 |
| legend | 42 |
| lane | 26 |
| plot | 200 |
| axis | 22 |
| brush | 74 |
| borders | 2 |

| State | What renders |
|---|---|
| Loading (the uPlot chunk or the data) | The real head with its controls (`Players`, `Window`, `Log scale`). `Skeleton` chips in the legend row. One `Skeleton` block over the plot area inside the gutters. A `Skeleton` strip where the brush goes. No spinner. The plot carries `aria-busy="true"` and the label `Loading the chart.` |
| Ready | as §B5 |
| Empty window | The axes, gridlines, phase band and **the brush** stay, with the window shown on the whole fight. The legend row keeps its height and shows no chips. One 14 px `--color-muted` line is centred in the plot (§B8). The head's meta says why: `{phase} · no damage` when the window sits inside a phase where the boss takes none, else `{start} to {end} · no damage`. |
| Error (the chunk failed) | The head, the empty legend row, `LoadError` centred over the plot area (message `The chart didn't load.` and its `Try again`, which re-imports the chunk), and the brush's reserved height. The tables below are unaffected. |
| A menu open over the plot | The menu draws above the plot (§C); the tooltip and crosshair stand down until it closes. |
| A wipe | Every death keeps its dashed line. The lane collapses them into one counted marker (§B5.6). |
| Night (a whole night, not one fight) | No chart, as today. |
| A report parsed before threat was kept second by second | No threat chart, and today's sentence, as today. |

### B8. Copy (verbatim; new keys in `web/src/lib/chart/copy.ts`)

| Key | Copy |
|---|---|
| Head meta, rate | `Raid {n} DPS · 5 s average` / `Raid {n} HPS · 5 s average` / `Raid {n} DTPS · 5 s average`; with a target filter, ` · on {target}` before ` · 5 s average` |
| Head meta, one player | `{name} · {n} DPS · 5 s average` (HPS, DTPS likewise) |
| Head meta, threat | `Cumulative · {start} to {end}` (`Cumulative · Whole fight` when not zoomed) |
| Head meta, empty | `{phase} · no damage` / `{start} to {end} · no damage` (healing, damage taken, threat likewise: `no healing`, `no damage taken`, `no threat`) |
| Threat title | `Estimated threat on {enemy}`; on a phone `Estimated threat` |
| Threat caveat (desktop) | `An estimate: 1 threat per point of damage and 0.5 per point of healing, split across the enemies in combat. No stance, form, talent or taunt modifiers yet.` then the link `What it leaves out` |
| Threat caveat (phone) | `Damage and healing only.` then the link `What it leaves out` |
| Threat table column heads | `Player`, `Estimated threat` (the bar), `Standing`, `Share of estimated threat`; the table's card title `Estimated threat`, meta `On {enemy} · standing at {m:ss} · {n} players` |
| Tooltip threat column heads | `Estimated threat`, `Share` |
| Players control label | `Players` |
| Players control field | `Top 8 of {N}` / `{n} of {N}` |
| Players control groups | `Compare`, `Tanks`, `Healers`, `Damage` |
| Players control compare rows | `Raid median` (description `Median DPS of the {n} damage dealers`; HPS of the healers on Healing); `{Class} median` (description `Median DPS of the {n} damage-dealing {class plural}`) |
| Players control foot | `Top 8`; `{n} drawn` |
| Players filter placeholder | `Filter players` |
| Always-drawn row description and chip tooltip | `Always drawn` |
| You tag | `You` (list), `YOU` (chip, uppercase by its recipe) |
| Legend | `Raid median`; `{Class} median`; `+{n} more` (phone: `+{n}`); `Only`; chip aria-label `Hide {name}` when on, `Show {name}` when off; `Only` aria-label `Show only {name}` |
| Legend sheet | title `Players`; meta `{n} of {N} drawn`; groups `Drawn`, `Not drawn`; buttons `Top 8`, `Done` |
| Zoom reset button | `Whole fight` |
| Log toggle | `Log scale` |
| Drag label | `{start} to {end} · {n} s` |
| Tooltip header, right | `{phase} phase · 5 s average` / `{phase} phase · estimated threat` (no `{phase} phase · ` when the fight has no phases) |
| Tooltip more | `+{n} more` |
| Tooltip, dead row value | `Died {m:ss}` |
| Event, death | `{name} died · {spell} from {source} · {amount}`, then ` ({overkill} overkill)` when the summary's `killing_blow.overkill` is above zero. Without a `killing_blow`: `{name} died`. |
| Event, death (phone event line) | `{name} died · {spell} · {amount}` |
| Event, death cluster | header `{start} to {end}`, meta `{n} deaths`; rows as Event, death, without the source; foot `+{n} more · open on the Deaths tab` |
| Event, target change | `{enemy} turned to {name}` |
| Event, taunt | `Taunt · {source} on {target}` (the spell's icon names the spell; Growl and Taunt differ) |
| Event, cast | `{spell} · {caster}` |
| Event, phase | `Phase · {name}` |
| Death marker aria-label | the death's event line, then `. Open on the Deaths tab.` |
| Plot aria-label | `{title} chart, {start} to {end}. Arrow keys move through time.` |
| Live region | `{m:ss}. {name} {value}, {name} {value}, {name} {value}.` |
| Loading | `Loading the chart.` (screen readers only) |
| Error | `The chart didn't load.` (button: `Try again`, LoadError's own) |
| Empty | `No damage in this window.` / `No healing in this window.` / `No damage taken in this window.` / `No threat on {enemy} in this window.` |
| Touch hint (the event line before the first scrub) | `Touch and hold the chart to read a moment.` |

The old threat note above the chart (`threat-chart-note`) and the `threat-incomplete` paragraph are replaced by the caveat line (§B12). Their facts move to the glossary.

### B9. Viewports

| Width | Behaviour |
|---|---|
| 2000 | The page column is centred and caps the content at 1248 px; the chart is identical to 1440 (`threat-chart-2000.png`). The threat plot is 300 px tall from 1440 to keep a readable aspect at the owner's width. |
| 1440 | As §B5. The tooltip shows up to 12 rows without clipping because it is in the top layer. |
| 1024 to 1439 | Plot heights as §B5.4; the legend wraps to two rows. |
| 768 to 1023 | Main plot 160 px, threat plot 220 px. Fine-pointer behaviour if the pointer is fine. |
| 360 to 767 | The layout of §B1.6, measured there.<br>Main chart: collapsed by default on every tab but Threat (today's `Show chart` row; its button becomes `button('secondary','sm')`).<br>Threat chart: open. Its head reads `Estimated threat` and the window. The Enemy `Select` sits full width at 44 px under the head. Then a one-row legend with `+{n}`, a 160 px plot with a 40 px gutter, the 28 px event line, the brush with 44 px handle hit areas, and the 28 px caveat line. There is no floating tooltip. |
| What never collapses | the x axis labels, the crosshair time pill, death markers, the brush, the caveat line |

The Source field on the Threat tab stays (fix round 1). It filters what the threat chart draws, the way it filters the table today. With one player as the Source, the chart draws that player and the always-drawn players. With All friendlies, it draws the default set.

### B10. Every chart on the site, and its target

| # | Chart | File | Today | Target | Rank |
|---|---|---|---|---|---|
| 1 | Report main chart (every table tab but Coach and Threat) | `report/TimeChart.svelte` via `ReportView.svelte` | Canvas, hand-drawn: one gold series, an optional ability line, 96 px tall, death ticks, a brush rectangle | `TimeSeriesChart` (uPlot): §B5 and §B6 in full. Top 8 of N, plus the Source player and you, the raid median and the Players control (D5). | Blocker |
| 2 | Threat chart | `report/ThreatTable.svelte` via `TimeChart` | Canvas, one line per player, taunt dots, no tooltip | `TimeSeriesChart`, `mode="cumulative"`. Titled Estimated threat, with the always-drawn tanks. Taunt and target-change markers, the share of estimated threat, the caveat line (§B12). 300 px from 1440. | Blocker |
| 3 | Sample report strip | `TimeChart` `compact` | 96 px strip, no legend | uPlot at 96 px: crosshair, tooltip and drag-to-zoom on; no legend, no lane, no brush | Must fix |
| 4 | Resource graphs | `report/ResourceGraphs.svelte` | SVG polylines, small multiples, not interactive | Stay SVG (small, many, static). Gain the synced crosshair and, while the cursor is set, the value at the cursor in the figure slot beside the line. X follows the window. Time labels from `time-axis.ts`. | Must fix |
| 5 | Timelines (casts and auras per player) | `report/TimelinesView.svelte` | DOM lanes | Stay DOM. Gain the synced crosshair, `time-axis.ts` labels, and a `ChartTooltip`-styled hover for a cast or aura. | Must fix |
| 6 | Raid cooldowns | `report/RaidCooldowns.svelte` | DOM lanes, ticks, deaths across | Stay DOM, as row 5 | Must fix |
| 7 | Character history: score and parse | `character/ChartPlot.svelte` via `ScoreChart.svelte`, `ParseChart.svelte` | SVG, sparse dated points, its own keyboard stepping (spec 2026-10-24) | Stay SVG: a few dozen points, a dated axis, and an owner-approved interaction. Gain the tooltip surface tokens of §B5.7 and the crosshair of §B5.8. | Polish |
| 8 | Sparklines | `Sparkline.svelte` in `ForeverScoreCard.svelte`, `RatingCard.svelte` | SVG | Unchanged | None |
| 9 | Sim distribution | `sim/DpsDistribution.svelte` | 20 divs | Stays divs (a shape, not a time series). Axis labels take the §B5.4 tick type; the band takes `rgba(229,185,85,.13)`; the mean is a 2 px `--color-gold` line. | Polish |
| 10 | In-row meters (report tables, guild attendance, deaths and raid rows, sim tool rows) | `ActorRow.svelte`, `guild/RaidAttendance.svelte`, `guild/RaidDeaths.svelte`, `GuildRaid.svelte`, sim tools | CSS bars | Unchanged: Details!-style meters are bars, not charts | None |
| 11 | Sim results time chart | `sim/SimResults.svelte` | Deliberately none (an average over iterations is not a fight) | Still none | None |

Guild pages and the history pages draw no other chart today.

### B11. Performance and polish

1. `uplot-loader.ts` starts the import when the report island mounts, in parallel with the summary fetch. The chart never waits on the chunk after its data arrives on a warm cache.
2. The frame reserves its exact height before the chunk or data arrives: head 48 px, legend 42 px (37 px on a phone), lane 26 px (when markers exist), plot, axis 22 px, brush 74 px (62 px on a phone), the caveat line 28 px on the threat chart. Every state of §B7 uses the same height. Cumulative Layout Shift on `/reports/*` stays under the 0.05 budget.
3. The chart chunk (uPlot plus the wrapper) is under 32 KB gzipped (raised from 30 KB by the §B3 ruling; measured 31,697 B). `check-island-size.mjs` gains it as a named budget. `report-island.js` stays within its existing 143,360 B budget.
4. Report Total Blocking Time stays under 200 ms (`lighthouserc.json`); uPlot's first draw of 10 series over 252 points is under 10 ms on Lighthouse's throttled CPU.
5. Every colour is read from tokens at draw time; no hex appears in chart code (today's rule in `TimeChart`).
6. Icons in the lane are requested with the summary, at 18 px in a fixed 20 px tile with a `--color-raised` ground, so they never flash.
7. The canvas is drawn at `devicePixelRatio`, and it redraws on `ResizeObserver` with `setSize`.

### B12. Threat, honestly (fix round 1)

**What the model is.**

1. The engine's threat model today is `base-1` (`logs/engine/summary/threat.go`, `BaseThreat`).
2. It models exactly two things:
   - one point of threat per point of effective damage to a hostile unit;
   - half a point per point of effective healing, spread across the enemies the healer is in combat with.
3. Its per-spell modifier table is empty, so `Complete()` is false. It does not model:
   - stance, form or Righteous Fury multipliers;
   - talent multipliers;
   - an ability's bonus threat (Sunder Armor, Heroic Strike, Revenge, Maul);
   - taunts;
   - threat drops and wipes (Feint, Fade, Wrath of Ragnaros' knockback);
   - the 110% and 130% rules for pulling aggro.
4. A tank therefore reads below the damage dealers they were holding threat over. The boards show this honestly: on the mock Ragnaros, the two tanks sit at about a third of the top damage dealers.

**What the page says.**

1. The chart title, the table's card title and every column that is a threat number say **Estimated threat**. The share column is **Share of estimated threat**: a player's part of every player's estimated threat on that enemy. That is the definition the table already uses (Glossary, "Threat share and bar"). The tooltip uses the same definition under the head `Share`. A share of the top value is never shown.
2. One caveat line sits in the threat chart's frame foot: 12 px `--color-muted`, a 1 px `--color-line-soft` top rule, padding 8 px 16 px, copy in §B8. Its link `What it leaves out` opens the report's glossary and scrolls to the Threat entry.
   - The glossary gains `id="glossary-threat"` on that entry.
   - The glossary `<details>` opens when the URL hash names an entry inside it.
   - The entry's text gains the list above: what the model does and does not include.
3. Players and markers carry the story the numbers cannot:
   - the tank is always drawn (§B5.3);
   - every taunt is marked;
   - every boss target change is marked, where the summary has them.
4. When the engine ships a complete model (`complete: true` on every row), the word `Estimated` and the caveat line go. The rest of this section stays.

**Boss target changes: the data contract.**

1. The summary does not carry them today. The engine adds a `target_changes` array to the summary: `{ at_ms, enemy_guid, enemy_name, to_guid, to_name, basis: 'melee' }`.
2. It is inferred from the enemy's melee swings. A change is recorded when a swing lands on a different unit from the previous swing's. The combat log carries no explicit "target changed" event, so this inference is the best the log allows. The glossary says so.
3. A summary without the field draws no target-change markers, and nothing else changes.
4. A known-good fixture (a log with a taunt and a tank death) pins the inferred changes before the markers ship (tenet 8). This is an engine-lane dependency. The chart lane ships everything else first.

## C. Menus always above charts (the overlay rule)

1. Every menu and list (the listbox, the More menu, the header account panel, the chart tooltip) renders in the browser's top layer through the Popover API. No overlay depends on `z-index` again. This removes the bug class the owner hit.
2. `lib/ui/overlay.svelte.ts` counts open overlays. One menu opening closes any other.
3. While the count is above zero, every chart hides its tooltip and crosshair and ignores hover. The chart never covers an open menu, and nothing under a menu reacts to the pointer.
4. The More menu (`ModeBar.svelte`) moves off `<details>` and `z-10` onto the same popover core in this lane.

## D. Acceptance screenshots

The implementer attaches each of these from the built fixture site (`FOREVER_DATA=fixture`), side by side with the board named, at the widths given. Every state is captured in a real browser, not composed.

| # | Capture | Widths | Compare with |
|---|---|---|---|
| 1 | Report Source list open, scrolled to the picked player, focus in its filter | 1440 | `dropdowns-open-1440.png` (middle) |
| 2 | Sim Fight style list open with descriptions | 1440 | `dropdowns-open-1440.png` (left) |
| 3 | Sim item search Source, filter `black`, then filter `naxx` | 1440 | `dropdowns-open-1440.png` (right, lower left) |
| 4 | A list near the viewport bottom, flipped above; a right-slot field opening bottom-end | 1440 | `dropdowns-open-1440.png` (lower right) |
| 5 | The `UiSheet` gallery: every field state at 52, 44 and 36 px, the option rows, dense and multi-select | 1440 | `dropdown-states-1440.png` |
| 6 | Keyboard and screen reader. A plain list: focus-visible, opened by Down, rows moved by the arrows, typeahead landing, focus on the trigger throughout. A searchable list: focus moves into the filter on open and returns to the trigger on close. A pointer parked on a row, then Tab: the keyboard-active row is committed, not the hovered one. **axe reports no violations with each variant open**, and there is exactly one `role="combobox"` per open list. | 1440 | §A6 |
| 7 | A list inside the raid form's dialog: drawn above the dialog; Escape closes the list and leaves the dialog open | 1440 | §A3.6 |
| 8 | The same field at 390 with the native picker open (iOS Simulator or Android emulator) | 390 | §A9 |
| 9 | Threat tab: hover at a death, taunt and target change within ±2 s; zoomed by drag; the caveat line; the table's leader equal to the legend's first chip | 1440, 2000 | `threat-chart-1440.png`, `threat-chart-2000.png` |
| 10 | Damage done on a 40-player fixture: top 8 of N, you pinned, raid median, the submerge flat, legend hover with `Only`, a drag in progress; the head figure equal to the sum of all players' figures (the test asserts it) | 1440 | `damage-timeline-1440.png` |
| 11 | The Players control open; a player toggled on draws at once; `Top 8` restores | 1440 | `choose-players-1440.png` |
| 12 | A death marker clicked: the Deaths tab opens with that death open and in view | 1440, 360 | §B6.6 |
| 13 | Log scale on | 1440 | §B6.5 |
| 14 | Cursor sync on the Buffs tab: pointer on the main chart, crosshair on the Raid cooldowns lanes too; then the reverse | 1440 | §B6.3 |
| 15 | Loading, empty (brush present), error, the Window menu open over the plot, a wipe fixture; all frames the same height (the test measures them) | 1440 | `chart-states-1440.png` |
| 16 | Threat tab on a phone after a tab change: the viewport with the measured fold stack (the e2e asserts the chart's bottom ≤ 640 at 360 × 640); mid-scrub; the event line linking a death | 360 | `controls-charts-phone-360.png` |
| 17 | The legend sheet open from `+{n}` | 360 | `legend-sheet-phone-360.png` |
| 18 | Reduced motion on: list open and chart hover | 1440 | §A11, §B6.9 |
| 19 | Lighthouse: `/reports/fixture2abcd.html`, `/sim.html`, the `/rankings` route, `/planner.html`; and `check-island-size.mjs` output | n/a | §A11, §B11 |

## E. Checklist

- [ ] `grep '<select'` returns only `Select.svelte` and `Select.astro`; the unit test enforces it.
- [ ] `FilterSelect.svelte` is deleted; `compact` boolean replaced by `size`.
- [ ] Every row of §A10 uses the variant and size named.
- [ ] Focus: plain lists keep it on the trigger; searchable lists move it to the filter input, the one combobox (§A6.1).
- [ ] Hover never commits; Tab commits only the keyboard-active row (§A6.2).
- [ ] The list is a child of the field's wrapper, shown in the top layer; inside a dialog Escape closes only the list.
- [ ] No Popover API: the native select stays. A hidden invalid select puts the error and focus on the trigger.
- [ ] The closed field's hover, open and pressed states are the /account trigger's.
- [ ] Dense lists use 13 px rows; rows grow to 44 px after a touch.
- [ ] Descriptions are `aria-describedby`, not part of the option's name.
- [ ] axe reports no violations with a plain, a searchable and the multi-select list open.
- [ ] The native-to-listbox swap moves the field by 0 px (e2e).
- [ ] Coarse pointer renders the native picker; `dense` renders 44 px there; the Players control opens the legend sheet.
- [ ] Every list renders in the top layer; the More menu is on the popover core.
- [ ] One open overlay at a time; charts stand down while one is open.
- [ ] uPlot is a lazy chunk under 32 KB gzipped (§B3 ruling); `report-island.js` stays within its 143,360 B budget.
- [ ] Both report charts use `TimeSeriesChart`.
- [ ] Damage draws top 8 of N, plus the Source player and you, plus the raid median; the Players control and the `players=` URL key work.
- [ ] The head's raid figure equals the sum of every player's figure (unit test on the adapter).
- [ ] Threat says `Estimated` in the title, the table and the tooltip; the share is `Share of estimated threat`; the caveat links to the glossary's Threat entry.
- [ ] Threat always draws the players the enemy attacked and the taunters in the window.
- [ ] Target-change markers draw when the summary carries `target_changes`, and nothing breaks when it does not.
- [ ] Death markers carry the killing blow and link to the Deaths tab; tooltips list events within ±2 s; a wipe collapses into a counted marker.
- [ ] Tooltip: sorted, class colours, events first, dead last, 12-row cap, flips left in the right 55%.
- [ ] Drag, brush, presets, double-click and Escape all set one window; the URL's `start` and `end` follow.
- [ ] Cursor sync across the charts, `TimelinesView`, `RaidCooldowns` and `ResourceGraphs`.
- [ ] Legend toggle, `Only` (with its focus-visible), Alt-click, `+{n} more`, hover highlight, pinned chips.
- [ ] Markers: deaths, target changes, taunts, cooldown and life-saving casts, phases with backed labels; the collision rule holds.
- [ ] Phases where the boss takes no damage draw flat at zero.
- [ ] Log scale on the three rate charts, in the URL, after the Window select.
- [ ] Keyboard on the plot and on the brush; live region text as §B8.
- [ ] Touch: long-press scrub, legend readout with no moving chip, event line with the death link, one-time hint in the event line, the legend sheet.
- [ ] The phone fold stack is 581 px at 360 × 640 (e2e).
- [ ] Every state in §A5 and §B7 captured at one frame height; every copy string verbatim from §A8 and §B8.
- [ ] CLS, TBT and island budgets hold (§A11, §B11).

## F. Order of work

0. Engine lane, in parallel: `target_changes` in the summary, inferred from melee swings, pinned by a fixture (§B12). The chart ships without it and draws the markers once it lands.
1. `lib/ui/overlay.svelte.ts` and the popover core; the More menu moves onto it (fixes the owner's bug class).
2. `listbox.ts` with Floating UI; `Select` and `Select.astro` gain the new props; `UiSheet` gains the gallery.
3. The call sites in §A10, report first (rows 1 to 13), then sim, rankings, guild, the rest.
4. uPlot, `TimeSeriesChart`, `ChartLegend`, `ChartTooltip`, `OverviewBrush`, `time-axis.ts`, `cursor.svelte.ts`.
5. The threat chart, then the main chart, then the sample strip.
6. Cursor sync into `TimelinesView`, `RaidCooldowns`, `ResourceGraphs`; then the §B10 polish rows.
