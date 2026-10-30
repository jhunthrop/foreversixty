# Controls and charts: one shared dropdown, and charts a raider trusts (experience spec)

Owner, 2026-10-30, on the report's Threat tab:

- "we have a bug where the dropdown disappears over the chart" (fixed in its own lane; §C makes the class of bug impossible).
- "our charts are very basic - should we be using a more advanced chart library?"
- "also the dropdown boxes look like basic html pages"

This spec covers two site-wide upgrades: one shared dropdown (Part A) and one chart system (Part B). It builds on `design/specs/2026-10-29-site-alignment.md` (the Card, the compact head, `quiet-label`, the framed canvas of §3.7.3) and `design/specs/2026-10-28-polish.md` §4.2 (the `Select` recipe) and §4.5 (type: every number is Barlow tabular, field labels are `quiet-label`). Nothing here restates those rules; it cites them.

Boards: `design/mocks/2026-10-30-controls-charts/renders/*.png`. Generator: `gen_controls_charts.py` and `render.mjs` in the same folder. The fight on the boards is mock data shaped like a Ragnaros kill (4:12, submerge at 3:00). Every string on the boards is this spec's copy, or copy that ships today.

| Board | What it shows |
|---|---|
| `dropdowns-open-1440.png` | The list open on a fine pointer: plain with descriptions (sim Fight style), with icons and groups (report Source), searchable (sim item search Source), searchable with no match, and the flip above the field |
| `dropdown-states-1440.png` | The closed field in every state at the three sizes, the lead-icon field, and the option-row anatomy |
| `threat-chart-1440.png`, `threat-chart-2000.png` | The Threat tab: tooltip at 2:07, the taunt and death markers, zoomed to 1:30 to 2:40, the overview brush, the table under it |
| `damage-timeline-1440.png` | Damage done: one line per player, the cast and death lane, legend hover, a drag-to-zoom in progress |
| `chart-states-1440.png` | Loading, empty, error, and a menu open over the plot |
| `controls-charts-phone.png` | 390: the native picker's closed field, the threat chart while a finger scrubs, the collapsed main chart |

## 0. Decisions

| # | Decision | Ruling in this spec |
|---|---|---|
| D1 | Chart library | **uPlot** (MIT, canvas, about 21 KB gzipped). It is confirmed over ECharts. Reasons in §B3. |
| D2 | Dropdown engine | **Our own listbox** in `web/src/lib/ui/listbox.ts`, positioned by **`@floating-ui/dom`** (about 4 KB gzipped), rendered in the browser's **top layer** through the Popover API (`popover="manual"`). No component library: Bits UI and Melt UI bring their own styling model and a Svelte-only runtime, and the Astro pages need the same control without Svelte. |
| D3 | Phones and coarse pointers | **The native picker**, behind the same styled closed field. The test is `matchMedia('(pointer: fine) and (hover: hover)')`. When it is false, the field stays a native `<select>`. |
| D4 | CSS customizable select (`appearance: base-select`) | **Not now.** It ships in Chromium only. When Safari and Firefox ship it, the listbox enhancer is deleted and this spec's look moves onto the native element. |
| D5 | The report's main chart content | **One line per player** (top 8 by the window's total) when Source is All friendlies, the way Warcraft Logs draws it. The raid's total moves to the chart head as a figure. One player as the Source keeps today's gold area. |

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
3. **`web/src/lib/ui/listbox.ts`** is the enhancer. It is framework-free TypeScript. On a fine pointer it finds each field marked `data-listbox`, hides the native `<select>` (the `hidden` attribute, which keeps it in form submission), and inserts the trigger button and the popover list. Choosing an option sets `select.value` and dispatches `input` and `change` on the select, so every call site's existing handler runs unchanged.
4. The enhancer is loaded with a dynamic `import()` on `requestIdleCallback` after first paint, only on a fine pointer. Until it loads, the native `<select>` works. Islands with a 16 KB budget (Rankings, GuildShell, Character) carry no listbox code in their entry.
5. The closed field is pixel-identical in both modes: same box, same text position, same chevron. The swap is invisible and moves nothing (§A11).
6. `@floating-ui/dom` (`computePosition` with `offset(4)`, `flip({ fallbackPlacements: ['top-start'] })`, `shift({ padding: 8 })`, `size()`, and `autoUpdate`) places the list. The list element carries `popover="manual"` and is shown with `showPopover()`, so it renders in the top layer above every stacking context.

The `Select` props after this spec:

| Prop | Type | Note |
|---|---|---|
| `options` | `SelectOption[]` | `{ value, label, group?, description?, meta?, icon?, colour?, disabled? }` |
| `option.icon` | `{ kind: 'crest' \| 'spec' \| 'spell' \| 'item'; src: string; quality?: Quality }` | `crest` is the ringed HD crest (one crest language). `item` draws a 1 px quality border and the label in the quality's `-text` colour. No call site uses `item` today; the look is fixed here so nobody invents it later. |
| `option.colour` | a token name, for example `--color-class-mage` | The label's colour. Class names only. |
| `option.meta` | `string` | Right-aligned 13 px muted figure, for example `×8`. |
| `size` | `'default' \| 'compact' \| 'dense'` | Replaces the `compact` boolean. 52, 44 and 36 px (§A4). |
| `searchable` | `boolean` | Adds the filter field (§A4.4). Set by the call site (§A10). |
| `searchPlaceholder` | `string` | Required when `searchable`. |
| `loading` | `boolean` | §A5. |
| `emptyLabel` | `string` | Shown in the field when `options` is empty. Required at every call site whose list can be empty. |
| unchanged | `value`, `id`, `name`, `label`, `placeholder`, `disabled`, `error`, `describedBy`, `testid`, `onchange` | |

The native `<option>` carries the rich data as attributes the enhancer reads: `data-icon`, `data-icon-kind`, `data-quality`, `data-description`, `data-meta`, `data-colour`. Groups are real `<optgroup label>` elements, so the native picker groups them too.

### A4. Layout and tokens

#### A4.1 The closed field (`dropdown-states-1440.png`)

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
| Placement | `bottom-start`, 4 px below the field. It flips to `top-start` when less than 240 px fits below and more fits above. It stays 8 px inside every viewport edge. |
| Width | at least the field's width, at most 420 px; content decides between them |
| Height | at most `min(360px, available space − 16px)`. Past that it scrolls (`overscroll-behavior: contain`) with a 24 px fade to `#0d0f15` at the bottom, and at the top once scrolled. |
| Open position | the picked option is scrolled into the middle of the list, or to the top when it is near the start |

#### A4.3 Option rows

| Row | Value |
|---|---|
| Plain | min-height 36 px, padding 6 px 40 px 6 px 12 px, 14 px 600 `--color-strong` |
| With an icon | min-height 40 px, the 20 px icon, 10 px gap |
| With a description | label as above; description 12/16 px 500 `--color-muted`, wraps, at most 330 px wide, no clamp (the text is the help, and it is never cut) |
| Meta | 13 px 500 `--color-muted`, tabular, right-aligned, 16 px from the label |
| Selected | a 16 px `check` glyph in `--color-gold`, 12 px from the right edge. The label keeps its own colour. |
| Active (keyboard or pointer) | ground `rgba(229,185,85,.10)` and a 2 px inset `--color-gold` left edge. Hover moves the active row, so there is one treatment. |
| Disabled | 45% opacity, `aria-disabled="true"`, skipped by arrows and typeahead, a click does nothing |
| Group head | 11 px 700 `.14em` uppercase `--color-muted`, padding 10 px 12 px 4 px. Every group after the first has a 1 px `--color-line-soft` top rule with 4 px above it. A group head sticks to the list's top while its rows scroll. |
| Filter match | the matched run in `--color-gold` 700 |

#### A4.4 The searchable list

1. A filter row sits at the top of the list, outside the scrolling region: padding 6 px 8 px 8 px and a 1 px `--color-line-soft` bottom rule.
2. The filter input is 36 px tall, radius 6 px, ground `--color-bg`, 1 px `--color-line` border (`--color-gold` with a 2 px `rgba(229,185,85,.35)` ring while focused), a 16 px `search` glyph at the left, and 14 px text.
3. The count sits at the right of the row: 12 px `--color-muted`, `aria-live="polite"`.
4. Matching is a case- and accent-insensitive substring match on the label. A group keeps its head when any of its rows match.
5. No match shows one 14 px muted row: the copy in §A8.

### A5. States

| State | Closed field | List |
|---|---|---|
| Default | as §A4.1 | closed |
| Hover | border `rgba(229,185,85,.30)`, ground `rgba(255,255,255,.06)` | row hover is the active row |
| Focus-visible | 2 px `--color-gold` outline, 2 px offset | the active row (focus stays on the trigger, `aria-activedescendant` names the row) |
| Active (pressed) | ground `rgba(255,255,255,.08)` | the pressed row keeps the active look |
| Open | border `rgba(229,185,85,.55)`, ground `rgba(255,255,255,.06)`, chevron turned 180° in `--color-gold` | shown |
| Disabled | 45% opacity, `cursor: not-allowed`, not focusable, does not open | n/a |
| Error | 1 px `--color-wipe` border; the 13 px `--color-wipe` line under the field, `role="alert"`, linked by `aria-describedby`; `aria-invalid="true"` | unchanged |
| Loading | label `Loading…` in `--color-muted`, a 14 px gold `Spinner` replaces the chevron, `aria-busy="true"`, does not open, keeps its width | n/a |
| Empty (no options) | the call site's `emptyLabel` in `--color-muted`, no chevron, disabled | n/a |
| Placeholder (nothing picked) | the placeholder in `--color-muted` 500 | no row checked |
| Search, no match | n/a | the no-match row |

Touch and keyboard parity: every state above renders the same on the native field, except Open, which is the phone's own picker.

### A6. Keyboard and pointer (fine pointer; WAI-ARIA select-only combobox pattern)

Roles: the trigger is a `<button type="button" role="combobox" aria-haspopup="listbox" aria-expanded aria-controls aria-labelledby>`. The list is `role="listbox"`. Rows are `role="option"` with `aria-selected`. A group is `role="group"` labelled by its head. In a searchable list, the filter input is the `role="combobox"` with `aria-autocomplete="list"` and `aria-activedescendant`.

| Key or action | Closed | Open |
|---|---|---|
| Down, Up, Alt+Down, Enter, Space | open, active row = the picked row | Down and Up move the active row by one (no wrap) |
| Home, End | open, active = first or last | first or last enabled row |
| PageDown, PageUp | n/a | move by 10 rows |
| A printable character | open and jump to the first row whose label starts with the typed run (500 ms buffer); in a searchable list, open and type it into the filter | typeahead (plain) or the filter (searchable) |
| Enter | n/a | pick the active row and close |
| Space | n/a | pick and close (plain); types a space (searchable) |
| Alt+Up | n/a | pick the active row and close |
| Escape | n/a | searchable with text: clear the filter first; otherwise close with no change |
| Tab | moves on | pick the active row, close, and move focus on |
| Click the field | open | close with no change |
| Click a row | n/a | pick and close |
| Click outside, or the field scrolls out of view | n/a | close with no change |
| Another menu opens | n/a | this one closes (one open overlay at a time, §C) |

Focus returns to the trigger on every close. Picking fires `change` once. Picking the already-picked row fires nothing.

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
| 1 | `report/ModeBar.svelte` | Source | raw | I + S | 36 | Crest and class colour per player. Groups: none for the two scopes, then `Tanks`, `Healers`, `Damage`, `Enemies` (role from the roster). `searchable={roster.length > 15}`. |
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

After this lane, `grep -rn '<select' web/src --include=*.svelte --include=*.astro` returns only `Select.svelte` and `Select.astro`. A unit test enforces it.

### A11. Performance and polish

1. The swap from native to listbox moves nothing: the wrapper has a fixed height per size, and an e2e test compares the field's bounding box before and after the enhancer loads (0 px difference).
2. The enhancer and Floating UI load as one lazy chunk, under 9 KB gzipped (`check-island-size.mjs` gains the entry).
3. Icons in the list never flash. Icons of the first 12 rows are requested on `pointerenter` or `focus` of the field. Each icon slot is a fixed 20 px box with a `--color-raised` ground until it paints.
4. Reduced motion: the list appears with no transition and the chevron turns with no transition. Otherwise the list fades and rises 2 px over 100 ms, and the chevron turns over 120 ms.
5. Lighthouse budgets in `web/lighthouserc.json` are unchanged and must hold on every route that renders a `Select`.

## Part B. Charts

### B1. The question and the answer in view

The raider's question at a chart is: "What happened, when, and who did it?"

1. The plot answers "what" and "when". Lines read at a glance, and markers name the moments (a death, a taunt, a cooldown, a phase).
2. The hover tooltip answers "who": every visible series at that second, sorted, in class colours, with the events of that second on top.
3. **Desktop 1280 × 800, Threat tab after the tab change scrolls to the mode bar:** the chart head, the legend, the whole plot, the overview brush and the first table row are visible.
4. **Phone 360 × 640, same moment:** the threat chart head, the Enemy field, the legend, the plot and the brush fit (about 520 px).
5. The main chart on a phone stays collapsed by default (`fs.report.chartOpen`, today's rule), so a table's answer stays first.
6. One chart per tab. On the Threat tab the threat chart is the tab's chart, and the main chart (and its collapsed row) is not drawn there. Today both draw, and the main chart pushes the threat chart under the fold at 1280 × 800.

### B2. Reference

| Reference | What we borrow |
|---|---|
| Warcraft Logs' graph (Damage Done, Threat) | One line per player in class colours. A crosshair tooltip listing every player at that time, sorted. Drag across the graph to zoom, and the tables follow the range. Death and cast markers on the time axis. The threat graph's per-player cumulative lines. |
| Details! (the timeline and the death log) | Event icons on the time line, and the spell's own icon on every cast. |
| The game's own combat text colours and the class colours | Class colours on every player, the death red, the gold for the player's own selection |
| Our framed canvas (site alignment §3.7.3) | The chart sits in a `card-surface-quiet` frame with a compact head. The head has no IconTile. |

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

### B4. Components

| Component | Role |
|---|---|
| `web/src/components/chart/TimeSeriesChart.svelte` (new, shared) | Owns the uPlot instance, the marker lane, the crosshair, the drag selection, and the keyboard and touch handling. Props: `series`, `markers`, `phases`, `domain` (the whole fight in ms), `window`, `onWindow`, `plotHeight`, `scale` (`'linear' \| 'log'`), `syncKey`, `mode` (`'rate' \| 'cumulative'`), `label`, `emptyCopy`. |
| `web/src/components/chart/ChartLegend.svelte` (new) | The chip row (§B5.3) |
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
| Title (`section-title` 18 px uppercase) | the table's metric: `Damage done`, `Healing`, `Damage taken` | `Threat on {enemy}` |
| Meta (13 px `--color-muted`, tabular) | `Raid {n} DPS · 5 s average` (HPS for Healing, DTPS for Damage taken). With one player as the Source: `{name} · {n} DPS · 5 s average`. | `Cumulative · {start} to {end}` |
| Right slot, left to right | `Switch` `Log scale`; `quiet-label` `Window` + `Select` dense (presets) | `quiet-label` `Enemy` + `Select` dense; then, when zoomed, `button('ghost','sm')` with the `x` glyph: `Whole fight` |

The main chart shows `Whole fight` as the Window select's first option, so it has no separate ghost button.

#### B5.3 Legend (`ChartLegend`)

1. A row under the head with a 1 px `--color-line-soft` top rule, padding 8 px 12 px 6 px, gap 4 px 6 px, wrapping. It reserves its height before data arrives (§B11).
2. A chip is 28 px tall (44 px hit area on a coarse pointer), padding 0 8 px, `--radius-pill`. It holds a 12 × 3 px swatch in the series colour, the name in 13 px 600 in its class colour, and a value in 13 px 500 `--color-muted`, tabular, with a fixed 44 px minimum width so a changing value never reflows the row.
3. Chips are ordered by the window's total, highest first. The order is recomputed when the window changes, never on hover.
4. The value is the window's figure: DPS (or HPS, DTPS) over the window, or threat standing at the window's end. A player who died shows the value at death.
5. The top 8 are on by default. The rest are listed and off.
6. An off chip is at 40% opacity with a hollow swatch.
7. When any chip is off, a `button('ghost','sm')` at the end of the row reads `Show all`.

#### B5.4 Marker lane and plot

| Part | Value |
|---|---|
| Marker lane | 26 px above the plot, drawn only when the chart has markers |
| Plot height | main chart: 200 px from 1024, 160 px from 768, 140 px below. Threat chart: 260 px from 1024, 220 px from 768, 180 px below. Sample report strip (`compact`): 96 px, no legend, no lane, as today. |
| Y gutter | 52 px (40 px below 768); labels 11 px `--color-muted` tabular, right-aligned 8 px from the plot: `0`, `400`, `1.2k`, `60k`, `1.2M` |
| Gridlines | 4 or 5 horizontal, 1 px `--color-line-soft`; no vertical gridlines |
| X axis | 1 px `--color-line` baseline, 4 px ticks, labels 11 px `--color-muted` tabular `m:ss` (`h:mm:ss` from one hour). The step is the smallest of 5, 10, 15, 30, 60, 120 or 300 s that keeps labels 64 px apart. The first label is left-aligned and the last right-aligned. |
| Right padding | 12 px |
| Lines | 1.5 px, round joins. Rate charts draw a 5 s trailing average with straight segments. Cumulative charts draw each second. |
| Single series | a gold area: `--color-gold` at 20% fill under a 1.5 px `--color-gold` line (today's look). A charted ability adds one line in the palette's second colour. |
| Phase bands | every second phase is filled `rgba(229,185,85,.05)`. A phase start is a 1 px dashed (2/3) `--color-gold` line at 70%. The phase name sits 6 px right of its start, 14 px from the plot's top: 11 px 700 `.12em` uppercase `--color-muted`. |

#### B5.5 Series colours

1. A player's line is their class colour token (`--color-class-*`).
2. Two players of one class share the colour, as on Warcraft Logs. The tooltip and legend name them, and hover isolates them.
3. Enemies and abilities take the palette in order: `--color-gold`, `--color-school-frost`, `--color-school-nature`, `--color-school-arcane`, `--color-school-fire`, `--color-school-shadow`. A seventh series repeats the palette with a 4/3 dash.

#### B5.6 Markers

| Marker | Lane | Plot |
|---|---|---|
| Death | a 20 px circle in `--color-death-soft` with a 1 px `--color-death` ring at 60% and the 14 px `skull` glyph in `--color-death` | a 1.5 px dashed (4/3) `--color-death` line, full height |
| Taunt (threat chart) | the taunt spell's own 18 px icon in a 20 px tile: `--color-raised` ground, 1 px `--color-gold` border at 70%, radius 4 px | a 1 px `--color-gold` line at 35%, full height |
| Cooldown cast | the spell's own icon, same tile as a taunt. The casts are the ones `isRaidCooldown` (`lib/report/raid-cooldowns.ts`) already names; nothing new is classified. With one player as the Source, that player's casts from the same list. | the same 1 px gold line at 35% |
| Phase change | none (the band and its label carry it) | as §B5.4 |

Collisions: a marker within 22 px of another steps 22 px right, with a 1 px `--color-gold` leader at 60% back to its time. Three or more markers inside 44 px collapse into one 20 px gold-ringed circle with the count in 11 px 700. Its tooltip lists them.

The `skull` glyph is added to `lib/ui/glyphs.ts` (24 grid, 2 px stroke, round caps; path in the generator).

#### B5.7 Tooltip (`ChartTooltip`)

| Part | Value |
|---|---|
| Box | 268 px wide, ground `#0d0f15`, 1 px `--color-line-card`, radius 10 px, `inset 0 1px 0 var(--card-highlight)` plus `0 16px 38px rgba(0,0,0,.6)`, padding 8 px 0 6 px, `pointer-events: none` |
| Header | the time in 14 px 700 `--color-strong` tabular; at the right, 12 px `--color-muted`: `{phase} · {what}` (`Phase 1 · threat on Ragnaros`, `Phase 1 · 5 s average`). A 1 px `--color-line-soft` rule under it. |
| Events | one 12 px `--color-text` line per marker in that second, with its 16 px icon: deaths, taunts, casts, phase starts. Copy in §B8. |
| Rows | 22 px each: a 3 × 12 px swatch, the name in 13 px 600 in its colour, the value in 13 px 600 `--color-strong` tabular at the right. The threat chart adds a 38 px share column (12 px `--color-muted`): the share of the top value. |
| Order | highest value first. A legend-hovered series is always included and its row has a `rgba(229,185,85,.08)` ground. Players who died before this second come last, in `--color-muted`, with `Died {m:ss}` as their value. |
| Limit | 12 rows, then one 12 px muted line: `+{n} more` |
| Placement | 16 px right of the crosshair, top-aligned with the plot. It flips to 16 px left of the crosshair when it would pass the frame's right edge. It stays 8 px inside the viewport. It renders in the top layer (`popover="manual"`) so a short plot never clips it. |
| Shown | while a fine pointer is over the plot, or the plot has keyboard focus |
| Hidden | while dragging; while any overlay is open (§C); when the pointer leaves the plot; on a coarse pointer (the legend reads out instead, §B6.8) |

#### B5.8 Crosshair

A 1 px dashed (3/3) `--color-text` line at 70%, full plot height. A 7 px dot in the series colour with a 1.5 px `--color-bg` ring sits on every visible line. A time pill sits on the x axis: 40 × 17 px, radius 4 px, ground `--color-text`, 11 px 700 `--color-bg` tabular.

#### B5.9 Overview brush (`OverviewBrush`)

1. A 36 px strip under the plot (padding-top 6 px), across the plot's width, radius 6 px, ground `rgba(255,255,255,.025)`, 1 px `--color-line-soft` border.
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
| Click a chip | toggles that series. The y scale refits to the visible series. |
| Hover or focus a chip (fine pointer) | that line draws at 2.5 px; every other line drops to 25% opacity; its tooltip row is pinned. An `Only` text button (12 px 600 `--color-gold`) appears in the chip. |
| Click `Only`, or Alt-click a chip | shows that series alone |
| `Show all` | every series on |
| Keyboard | each chip is a `button` with `aria-pressed`; `Only` is a second button inside the chip's focus group |
| Touch | a tap toggles. Isolating is a fine-pointer accelerator; on touch, `Show all` and toggles do the same job. |

The legend's state lasts for the page view. It is not written to the URL. The window and the log scale are.

#### B6.5 Log scale

1. A `Switch` labelled `Log scale` in the head of Damage done, Healing and Damage taken. Threat is cumulative from zero and has none.
2. On: y ticks at powers of ten (`10`, `100`, `1k`, `10k`), a floor of 1, and a zero drawn at the floor.
3. The URL carries `scale=log`, so a shared link reproduces it.

#### B6.6 Markers

Hovering a marker in the lane shows the tooltip at the marker's second. A marker is never a link.

#### B6.7 Keyboard (the plot is focusable)

1. The plot is `tabindex="0"`, `role="group"`, `aria-roledescription="chart"`, with the `aria-label` in §B8. Focus draws a 2 px inset `--color-gold` ring.
2. Left and Right move the crosshair one second; with Shift, ten. The tooltip follows.
3. A polite live region reads the time and the top three series at that second, 300 ms after the last key.
4. Escape returns to the whole fight when zoomed.
5. The window is set from the keyboard with the brush's range inputs: arrows ±1 s, Shift+arrows ±10 s, PageUp and PageDown ±30 s, Home and End to the fight's bounds.

#### B6.8 Touch (`controls-charts-phone.png`)

1. The plot has `touch-action: pan-y`, so a swipe scrolls the page.
2. A touch held for 300 ms starts a scrub. The crosshair follows the finger.
3. While scrubbing, the legend becomes the readout: each chip's value shows the value at the finger, and a time pill (the crosshair's pill style, 24 px tall) leads the legend row. The chips do not reorder.
4. A 28 px event line under the plot (reserved at all times on a coarse pointer) names the marker at the finger's second: `Taunt · Thandril on Ragnaros`.
5. Lifting the finger keeps the last reading until the next tap outside the plot.
6. Zoom on touch is the brush handles and the Window select. The plot itself takes no drag-to-zoom on touch.
7. One 12 px `--color-muted` line under the chart, shown until the first scrub and then never again (`localStorage` `fs.chart.scrubHint`): `Touch and hold the chart to read a moment.`

#### B6.9 Reduced motion

There is no zoom animation at any setting. Under `prefers-reduced-motion: reduce`, the tooltip appears with no fade (otherwise an 80 ms fade) and the loading skeleton does not shimmer.

### B7. States (`chart-states-1440.png`)

| State | What renders |
|---|---|
| Loading (the uPlot chunk or the data) | The frame at its final height: the real head; `Skeleton` chips in the legend; one `Skeleton` block over the plot area inside the gutters; the brush strip's box. No spinner. The plot carries `aria-busy="true"` and the label `Loading the chart.` |
| Ready | as §B5 |
| Empty window | axes, gridlines and the brush stay. One 14 px `--color-muted` line centred in the plot (§B8). The legend keeps its height and shows no chips. |
| Error (the chunk failed) | the head; a 248 px body with `LoadError` centred: message `The chart didn't load.` and its `Try again`, which re-imports the chunk. The tables below are unaffected. |
| A menu open over the plot | the menu draws above the plot (§C); the tooltip and crosshair stand down until it closes |
| Night (a whole night, not one fight) | no chart, as today |
| A report parsed before threat was kept second by second | no threat chart, today's sentence, as today |

### B8. Copy (verbatim; new keys in `web/src/lib/chart/copy.ts`)

| Key | Copy |
|---|---|
| Head meta, rate | `Raid {n} DPS · 5 s average` / `Raid {n} HPS · 5 s average` / `Raid {n} DTPS · 5 s average` |
| Head meta, one player | `{name} · {n} DPS · 5 s average` (HPS, DTPS likewise) |
| Head meta, threat | `Cumulative · {start} to {end}` (`Cumulative · Whole fight` when not zoomed) |
| Threat title | `Threat on {enemy}` |
| Zoom reset button | `Whole fight` |
| Log toggle | `Log scale` |
| Legend | `Show all`; `Only`; chip aria-label `Hide {name}` when on, `Show {name}` when off; `Only` aria-label `Show only {name}` |
| Drag label | `{start} to {end} · {n} s` |
| Tooltip header, right | `{phase} · 5 s average` / `{phase} · threat on {enemy}` (no phase prefix when the fight has no phases) |
| Tooltip more | `+{n} more` |
| Tooltip, dead row value | `Died {m:ss}` |
| Event, death | `Death · {name}` |
| Event, taunt | `Taunt · {source} on {target}` (the spell's icon names the spell; Growl and Taunt differ) |
| Event, cast | `{spell} · {caster}` |
| Event, phase | `Phase · {name}` |
| Plot aria-label | `{title} chart, {start} to {end}. Arrow keys move through time.` |
| Live region | `{m:ss}. {name} {value}, {name} {value}, {name} {value}.` |
| Loading | `Loading the chart.` (screen readers only) |
| Error | `The chart didn't load.` (button: `Try again`, LoadError's own) |
| Empty | `No damage in this window.` / `No healing in this window.` / `No damage taken in this window.` / `No threat on {enemy} in this window.` |
| Touch hint | `Touch and hold the chart to read a moment.` |

The threat note under the chart (`threat-chart-note`) keeps today's copy and moves under the frame, 12 px `--color-muted`, at most 900 px wide.

### B9. Viewports

| Width | Behaviour |
|---|---|
| 2000 | The page column is centred and caps the content at 1248 px; the chart is identical to 1440 (`threat-chart-2000.png`). The legend stays one row for 10 series. |
| 1440 | As §B5. The threat plot is 260 px; the tooltip shows up to 12 rows without clipping because it is in the top layer. |
| 1024 to 1279 | Plot heights as §B5.4; the legend wraps to two rows. |
| 768 to 1023 | Main plot 160 px, threat plot 220 px. Fine-pointer behaviour if the pointer is fine. |
| 390 | Main chart collapsed by default (today's `Show chart` row, its button becomes `button('secondary','sm')`). Threat chart open: Enemy `Select` full width at 44 px under the head; legend top 6 plus `+{n} more`; plot 180 px with a 40 px gutter; the 28 px event line; the brush with 44 px handles. No floating tooltip. |
| What never collapses | the x axis labels, the crosshair time pill, death markers, the brush |

### B10. Every chart on the site, and its target

| # | Chart | File | Today | Target | Rank |
|---|---|---|---|---|---|
| 1 | Report main chart (every table tab but Coach and Threat) | `report/TimeChart.svelte` via `ReportView.svelte` | Canvas, hand-drawn: one gold series, an optional ability line, 96 px tall, death ticks, a brush rectangle | `TimeSeriesChart` (uPlot): §B5 and §B6 in full; per-player lines (D5) | Blocker |
| 2 | Threat chart | `report/ThreatTable.svelte` via `TimeChart` | Canvas, one line per player, taunt dots, no tooltip | `TimeSeriesChart`, `mode="cumulative"`: Enemy in the head, taunt icons, share column, 260 px | Blocker |
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
2. The frame reserves its exact height before the chunk or data arrives: head 48 px, legend 42 px, lane 26 px (when markers exist), plot, axis 22 px, brush 60 px. Cumulative Layout Shift on `/reports/*` stays under the 0.05 budget.
3. The chart chunk (uPlot plus the wrapper) is under 30 KB gzipped. `check-island-size.mjs` gains it as a named budget.
4. Report Total Blocking Time stays under 200 ms (`lighthouserc.json`); uPlot's first draw of 10 series over 252 points is under 10 ms on Lighthouse's throttled CPU.
5. Every colour is read from tokens at draw time; no hex appears in chart code (today's rule in `TimeChart`).
6. Icons in the lane are requested with the summary, at 18 px in a fixed 20 px tile with a `--color-raised` ground, so they never flash.
7. The canvas is drawn at `devicePixelRatio`, and it redraws on `ResizeObserver` with `setSize`.

## C. Menus always above charts (the overlay rule)

1. Every menu and list (the listbox, the More menu, the header account panel, the chart tooltip) renders in the browser's top layer through the Popover API. No overlay depends on `z-index` again. This removes the bug class the owner hit.
2. `lib/ui/overlay.svelte.ts` counts open overlays. One menu opening closes any other.
3. While the count is above zero, every chart hides its tooltip and crosshair and ignores hover. The chart never covers an open menu, and nothing under a menu reacts to the pointer.
4. The More menu (`ModeBar.svelte`) moves off `<details>` and `z-10` onto the same popover core in this lane.

## D. Acceptance screenshots

The implementer attaches each of these from the built fixture site (`FOREVER_DATA=fixture`), side by side with the board named, at the widths given. Every state is captured in a real browser, not composed.

| # | Capture | Widths | Compare with |
|---|---|---|---|
| 1 | Report Source list open, scrolled to the picked player | 1440 | `dropdowns-open-1440.png` (middle) |
| 2 | Sim Fight style list open with descriptions | 1440 | `dropdowns-open-1440.png` (left) |
| 3 | Sim item search Source, filter `black`, then filter `naxx` | 1440 | `dropdowns-open-1440.png` (right, lower left) |
| 4 | A list near the viewport bottom, flipped above | 1440 | `dropdowns-open-1440.png` (lower right) |
| 5 | The `UiSheet` gallery: every field state at 52, 44 and 36 px, and the option-row anatomy | 1440 | `dropdown-states-1440.png` |
| 6 | Keyboard: focus-visible on a field, the list open by Down, the active row moved by arrows, typeahead landing | 1440 | §A6 |
| 7 | The same field at 390 with the native picker open (iOS Simulator or Android emulator) | 390 | §A9 |
| 8 | Threat tab, hover at a taunt, zoomed by drag | 1440, 2000 | `threat-chart-1440.png`, `threat-chart-2000.png` |
| 9 | Damage done, legend hover with `Only`, a drag in progress, cooldown and death markers | 1440 | `damage-timeline-1440.png` |
| 10 | Log scale on | 1440 | §B6.5 |
| 11 | Cursor sync on the Buffs tab: pointer on the main chart, crosshair on the Raid cooldowns lanes too; then the reverse | 1440 | §B6.3 |
| 12 | Loading, empty, error, and the Source menu open over the plot | 1440 | `chart-states-1440.png` |
| 13 | Threat tab on a phone, mid-scrub, and the first-visit hint | 390 | `controls-charts-phone.png` |
| 14 | Reduced motion on: list open and chart hover | 1440 | §A11, §B6.9 |
| 15 | Lighthouse: `/reports/fixture2abcd.html`, `/sim.html`, `/rankings` route, `/planner.html`; and `check-island-size.mjs` output | n/a | §A11, §B11 |

## E. Checklist

- [ ] `grep '<select'` returns only `Select.svelte` and `Select.astro`; the unit test enforces it.
- [ ] `FilterSelect.svelte` is deleted; `compact` boolean replaced by `size`.
- [ ] Every row of §A10 uses the variant and size named.
- [ ] The listbox follows §A6 for every key; axe reports no violations with a list open.
- [ ] The native-to-listbox swap moves the field by 0 px (e2e).
- [ ] Coarse pointer renders the native picker; `dense` renders 44 px there.
- [ ] Every list renders in the top layer; the More menu is on the popover core.
- [ ] One open overlay at a time; charts stand down while one is open.
- [ ] uPlot is a lazy chunk under 30 KB gzipped; `report-island.js` does not grow.
- [ ] Both report charts use `TimeSeriesChart`; the main chart draws one line per player for All friendlies.
- [ ] Tooltip: sorted, class colours, events first, dead last, 12-row cap, flips at the right edge.
- [ ] Drag, brush, presets, double-click and Escape all set one window; the URL's `start` and `end` follow.
- [ ] Cursor sync across the charts, `TimelinesView`, `RaidCooldowns` and `ResourceGraphs`.
- [ ] Legend toggle, `Only`, Alt-click, `Show all`, hover highlight.
- [ ] Markers: deaths, taunts, cooldown casts, phases; the collision rule holds.
- [ ] Log scale on the three rate charts, in the URL.
- [ ] Keyboard on the plot and on the brush; live region text as §B8.
- [ ] Touch: long-press scrub, legend readout, event line, one-time hint.
- [ ] Every state in §A5 and §B7 captured; every copy string verbatim from §A8 and §B8.
- [ ] CLS, TBT and island budgets hold (§A11, §B11).

## F. Order of work

1. `lib/ui/overlay.svelte.ts` and the popover core; the More menu moves onto it (fixes the owner's bug class).
2. `listbox.ts` with Floating UI; `Select` and `Select.astro` gain the new props; `UiSheet` gains the gallery.
3. The call sites in §A10, report first (rows 1 to 13), then sim, rankings, guild, the rest.
4. uPlot, `TimeSeriesChart`, `ChartLegend`, `ChartTooltip`, `OverviewBrush`, `time-axis.ts`, `cursor.svelte.ts`.
5. The threat chart, then the main chart, then the sample strip.
6. Cursor sync into `TimelinesView`, `RaidCooldowns`, `ResourceGraphs`; then the §B10 polish rows.
