# Web layout system, and the shared paperdoll that leads Best in slot: experience spec

Author: ux-designer. 2026-11-05. Round 2 (same day): the type is decided (section 3.3), the shared primitives are defined (section 3A), and the guild landing and Overview are mocked from those primitives alone to prove the system carries over (section 6.1). Round 1 was the system and the BiS page. The other pages follow one at a time (section 9), on the owner's order: page by page, BiS first.

Owner directions this answers (all binding):
1. "It looks like a news site with the slim center. Use more of the screen." (about 2000 px wide)
2. "Does this feel cluttered? A whole new design layout for the web."
3. "Dense, but simplistically clean interfaces." Fewer borders and cards; hierarchy from type, spacing and alignment; the best WoW addons and Wowhead's data views, not a marketing site.
4. "Players are used to BiS lists on Wowhead like their paperdoll setup. We have a paperdoll but it's also not great." BiS leads with a paperdoll; clicking a slot opens its runners-up in a side panel; signed in, the doll shows the player's own gear against the pick; evaluate a "where to go" grouping; one shared paperdoll for BiS, character, sim and reports.

References: `docs/tenets.md` 1, 2, 4, 9, 10, 11, 12, 13, 14; `design/DESIGN-SYSTEM.md` principle 6 ("content areas are flat, calm, and dense enough to be useful") and 7 ("ornament stays out"), which the shipped pages have drifted from; `design/specs/2026-10-31-paperdoll.md` (the component this replaces); `design/specs/2026-10-10-bis-landing.md`; memory notes `feedback-wide-screen-check`, `feedback-one-way-to-switch` (the header selector owns character context), `feedback-provenance-quiet`, `feedback-one-job-per-page`.

## 0. Mocks and captures (all in the repo)

| What | Path |
|---|---|
| Mock generators | `design/mocks/ws_lib.py` (the system CSS and the chosen type), `ws_prims.py` (the primitives, section 3A), `ws_bis.py` (BiS + paperdoll), `ws_guild.py` (guild, built only from `ws_prims`), `ws_guild_before.py`, `ws_compose.py` (before/after), `ws_shoot.mjs` |
| BiS mocks (responsive HTML, one file each; rev 3) | `design/mocks/web/web-bis.html` (band 20, signed out, Where to go default, reputation off), `web-bis-slots.html`, `web-bis-you.html` and `web-bis-you-where.html` (signed in), `web-bis-dungeon.html` (Dungeon drops only), `web-bis-sources.html` (the Sources checklist open), `web-bis-60.html` and `web-bis-60-you.html` (band 60: enchants, trinket inspector, weapon note, hit cap), `web-bis-weights-open.html`, `web-bis-open.html` and `web-bis-bar.html` (phone), `web-bis-states.html` (every cell state). Generator `ws_bis.py` |
| Renders of those at 390, 1440, 1920, 2000 (full page and first screen `-fold`) | `design/mocks/renders/web/web-bis-<width>.png` (that folder is gitignored; regenerate with `python3 design/mocks/ws_bis.py && node design/mocks/ws_shoot.mjs web-bis 1440,1920,2000,390`). The key captures are copied, tracked, to `design/reviews/2026-11-05-wide-layout/after/` |
| Before/after at the same scale | `design/mocks/renders/web/compare-bis-1440.png`, `-1920.png`, `-2000.png`, `-390.png` |
| Live "before" captures | `design/reviews/2026-11-05-wide-layout/before/` (bis, guide, sim, tiers, guild, home at 1440, 1920, 2000; bis also at 390) |
| Guild mocks (section 6.1) | `design/mocks/web/web-guild.html` (Overview, officer), `web-guild-landing.html` (signed in), `web-guild-landing-out.html` (signed out), `web-guild-before.html` (the shipped Overview structure reproduced). Compare images: `compare-guild-overview-<w>.png`, `compare-guild-landing-<w>.png` at 390, 1440, 1920, 2000 |
| Type study | `design/mocks/web/web-bis-font-b2-serif-class.html` and its siblings (the owner's choice, 2026-11-05) |
| Width-only fallback (Option B, section 8) | `design/mocks/wide/wide-*.html`, renders in `design/mocks/renders/wide/` |

Data in the BiS mocks: the repo's active build (1.60.1.70532), band 20 to 29, Alliance Dwarf, the page's default. Every item, icon, quality, source, drop chance, item level, stat and DPS delta is read from `data/builds/`. The live page serves a newer build, so a few picks differ (the mock's Bronze Dory is the live page's Impaling Harpoon); layout is unaffected. The signed-in character "Felsnap" is a sample whose worn items are the page's own runner-ups, so every gain is a real sim number.

## 1. Diagnosis: confirmed, and what the captures add

Confirmed from the live captures (`before/bis-1920.png`, `bis-1440.png`, `bis-390.png`):

1. **Every region is the same card.** Gold-bordered 1 px box, gradient top, icon chip, caps title: This set, Stat weights, Play it, The list, Where to get it, New at this level range. Six boxes signed out, seven signed in (Your character), one weight. Nothing is primary.
2. **The page's job starts below the fold.** Measured on production 2026-11-05: at 1920 The list starts at y 1,145 of a 4,502 px page (the first screen is 1,000 px), at 390 it starts at y 2,328 of 8,153: three screens of phone scrolling before the first slot.
3. **Controls in three rows, three styles**: spec as icon tabs, level and faction as pills, race as a select. Six controls, three visual languages.
4. **The middle cards are mismatched.** At 1920 This set is 235 px tall holding one number, Stat weights 512 px, Play it 621 px, side by side; the shortest is mostly empty. At 390 they stack to 295, 520 and 665 px (1,480 px) before the list.
5. **Play it restates the guide's rotation.** One job per page: BiS owns gear. It becomes one link.
6. **Text is 11 to 12 px** inside large boxes (weights caption, source lines, DPS deltas, column heads), and the deltas that are the answer sit at the far right edge at that size.
7. **Width.** The column is 1,344 px (content 1,248). At 1920 the page uses 65% of the screen, at 2000 62%; the right 336 to 376 px on each side is empty, and the page head's right half is empty at every width.

Added by this review:

8. **A tall block per slot.** The list is 2,394 px at 1920 for 17 slots (about 141 px a slot) and 3,988 px at 390. Runners-up are rows of equal weight to the pick, up to three per slot: 54 icons in the list for a set of 14 items.
9. **The pick and its runners-up look alike.** Same icon size, same blue, same type size; only position tells them apart.
10. **"NEW" on every item** (it means "new at this level range" and at 20 to 29 everything is new), and a gold check with no label. Both are noise: at 20 to 29, "NEW" carries no information. The check means "sim-checked"; it is moved into the inspector (section 5.5).
11. **A source line is 11 px grey**, though "where do I get it" is half of why a player opens this page.
12. **The same data is shown twice**: a "Where to get it" card at the page's bottom already groups pieces by source (Wailing Caverns +4, The Deadmines +3, World drop +3, ...). It is the right idea in the wrong place: after 3,300 px of scroll. Section 4.7 promotes it.
13. **Bug: `?band=60` does not switch the band.** Verified on production 2026-11-05: `/bis/hunter/beast-mastery?band=60` renders "20 to 29" selected. The URL must be the source of truth for band, faction and race (section 5.8, requirement R1).
14. **The shipped Paperdoll (character page, sim, reports) has its own problems**, section 4.1.

## 2. Direction: one flat system, dense, quiet

Reference surfaces, and what is borrowed:
- **Wowhead gear tables and the Classic BiS lists.** One row per slot, icon in a quality border, name in quality colour, source visible on the row, one click for the rest. Borrowed: row grammar, density (58 px rows), source on the row. Not borrowed: ads, tabs that hide the list, five columns of text.
- **The in-game character panel.** Slots in two columns, weapons at the bottom, icon-first. Borrowed: slot order head to feet and the idea that the whole set is one object seen at once. Not borrowed: the mirrored right column with names pushed to the screen edge (section 4.1).
- **Details! and WeakAuras options.** One row style, numbers aligned in a column, one highlight for the selected row, no boxes inside boxes.

Principles (each replaces a defect above):
- **P1. Three surfaces only.** S0 page `#07090d`. S1 plate `#0d111a`, 1 px `--line` `#262e40`, 6 px radius: used for the page's one primary object (the doll, the table, the result) and for overlays (tooltips, menus). S2 hover and selected `#131824`. Nothing else has a box. Secondary content is open on S0, separated by 48 px of space and a 1 px `--line-soft` rule under a text heading.
- **P2. Gold is rare.** Gold means: the one primary button, the selected state, links, the answer number. The gold-warm border and the card gradient are retired.
- **P3. Hierarchy from type.** One Title, one answer number, section heads in text. No icon chips.
- **P4. Dense by default.** 58 px rows, 13 px minimum for any sentence, 12 px only for uppercase labels. Whitespace is spent between sections (48 px, 64 px at xl), not inside rows.
- **P5. One control bar.** Section 3.4.
- **P6. The player, not a card.** Character context lives in the header selector (section 3.6).

## 3. The system

### 3.1 Tokens (existing, from `web/src/styles/tokens.css`, unchanged)
`--color-bg #07090d`, `--color-raised #0d111a`, `--color-card-top #131824` (reused as S2), `--color-line #262e40`, `--color-line-soft #1c2230`, `--color-text #e9e4d8`, `--color-strong #f2eee4`, `--color-muted #9a9484`, `--color-nav #b9b3a4`, `--color-gold #e5b955`, `--color-gold-hover #f5d27a`, `--color-kill #7fd48a`, `--color-wipe #f0736b`, class and rarity colours, `--radius-panel 6px`, `--radius-control 4px`. Item names use the AA-safe `--color-rarity-rare-text` and `--color-rarity-epic-text`.

### 3.2 Proposed tokens (new)
| Token | Value | Use |
|---|---|---|
| `--page-max` | 1344 px below 1760 px; at 1760 px and up `clamp(1600px, 100vw - 160px, 1792px)` | The one site column (Base.astro). 1920 gives 1760, 2000 gives 1792, 2560 stays 1792. |
| `--page-gutter` | 48 px; 64 px at 1760 px and up | Content width: 1248 at 1440; 1632 at 1920; 1664 at 2000. |
| `--section-gap` | 48 px; 64 px at 1760 px and up | Space between page sections. |
| `--gap` | 24 px; 32 px at 1760 px and up | Grid gap. |
| `--rail-w` | 360 px; 400 px at 1760 px and up | A side rail. The paperdoll inspector is 320 at 1024 to 1759 and 360 at 1760 up (`--inspector-w`; 440 until 2026-11-06, narrowed so the doll's rows hold two lines at 1920: rows beat inspector width). |
| `--bar-h` | 56 px | The control bar. |
| `--row-h` | 58 px | A dense row (paperdoll cell). |
| `--toc-w` | 200 px | The guide's contents rail. |
| `--measure-reading` | 720 px (unchanged) | Prose. |

Breakpoints (Tailwind names): `md` 768, `lg` 1024, and one new `xl` at 1760. Below 1760 every page composes as today's 1344 column; at 1760 and up the column widens and the layouts below add a rail. `site-column.test.ts` still guards the one-container rule; it reads `--page-max`.

Amended 2026-11-10 (owner: "why don't we have consistent width for every experience?"): the column is the one page frame on every page, with no per-page opt-in (the `wide` prop and `data-column` are gone). `.page-column` holds the site header bar (Base.astro wraps it), every page head band, the control bar, the page body and the footer, so all of them share one left and one right edge at a given viewport, the API's /b pages too (chrome.html). Only `.page-column` reads `--page-max` and `--page-gutter`. A narrower measure inside the frame is named: the `.measure-*` classes and ui/Page's `reading` and `form` layouts. Review: `design/reviews/2026-11-10-header-width/`.

Why 1760 and 1792 and not full width: a 1,664 px content line is the longest a list row can run before the eye cannot return to its slot label. Above 2000 the margins grow; the page does not.

### 3.3 Type scale (decided by the owner 2026-11-05: option 2 with class colour; nothing under 12)
Mock: `design/mocks/web/web-bis-font-b2-serif-class.html` (the chosen face study); every mock in this spec now uses it through `ws_lib.py`.
| Role | Size / line | Face and weight | Use |
|---|---|---|---|
| Title | 40 / 46 (32 / 38 phone) | Marcellus 400, normal case, tracking .01em, **in the spec or class colour** (faction colour on a guild; `--color-strong` where there is none) | The H1. Once per page. The only use of Marcellus (plus the Cinzel wordmark in the header, which is the logo, not type) |
| Section head | 22 / 28 (20 / 26 phone) on reading pages, 18 / 24 elsewhere | Barlow 700, normal case, `--color-strong` | `Section` titles. No caps, no condensed face |
| Item name | 15 / 20 in rows, 18 / 22 in the inspector | Barlow **700** in the AA-safe quality colour | Every item name, everywhere including the inspector and tooltips' title |
| Lead | 17 / 26, max 62 ch | Barlow 400 | One sentence under a Title where needed |
| Prose | 16 / 26, max 720 px | Barlow 400 | Guides |
| Body | 15 / 24 | Barlow 400 and 600 | Default |
| Small | 13 / 20, `--muted` | Barlow 400 | Sources, captions, secondary |
| Label | 12 / 16, caps, tracking .12em, `--muted` | Barlow 700 | Group labels and column heads only |
| Figure | 15 to 56 | **Barlow 700 with tabular figures** (`font-variant-numeric: tabular-nums`) | Every number that is compared. **No mono face anywhere** (JetBrains Mono leaves the site's UI) |
Retired by this decision: Barlow Condensed 800 caps for headings and titles, JetBrains Mono for figures, the `section-title` and `heading-plan` caps recipes. Marcellus is loaded once, `font-display: swap`, with a size-adjusted fallback so the title does not shift (tenet 13); it is not preloaded on pages that have no title in view.

### 3.4 The one control bar
Pattern for every page that has choices (BiS spec, level, faction, race; tiers role and faction; sim fight; guild tabs keep the existing tab strip).
- **Desktop (768 and up):** one row, `--bar-h` 56 px, sticky under the header, full-bleed hairlines above and below (`--line-soft`), content in the page column. Items left to right: segmented controls, then selects, then `end` content (a link or a view toggle). Group labels (12 px caps) appear at 1280 and up.
- **Segmented control:** a 1 px `--line` outline around the group, 40 px high, items divided by hairlines, 14 px 600 text. Selected: gold text, `#e5b9551a` ground, 2 px gold underline inside the item. A count or icon may sit in an item (spec icons at 22 px). All options visible; no overflow menu.
- **Select:** same outline and height, 14 px 600 text, a caret.
- **Phone (below 768):** the bar is one 48 px line: the current selection in words ("Beast Mastery · 20-29 · Alliance · Dwarf") and "Change ▾". Activating it expands the controls in place, stacked, each segmented control full width at 44 px, labelled with its group label; the page below moves down (no sheet, no overlay). It collapses on a selection. `web-bis-bar-390.png`.
- **States:** default; hover (`#131824` ground, `--strong` text); focus-visible (2 px gold outline, 2 px offset); active (pressed: `#1a2030`); selected (above); disabled (a level with no data: `--muted` text at 50%, `aria-disabled`, no hover, a title "No data yet at this level"); loading (selection is applied at once and the plate below shows its skeleton; the bar never shifts); error (the bar is unaffected; the plate shows the LoadError row).
- **Keyboard:** each segmented control is a radio group (arrow keys move and select, Tab leaves the group). Touch: 40 px items on desktop, 44 px on phone.
- **URL is the truth.** Every control reads from and writes to the query string (`?band=60&faction=horde&race=orc`) with `history.replaceState`; a reload, a share and the back button restore it. Spec is the path segment, as today.
- Components: this is `PresetControl.astro` and `BisBandStrip.astro` merged into one `ControlBar` (shared with tiers and sim). The old pill-and-tab styles are removed from these pages.

### 3.5 Buttons and links
One primary (gold fill, `#17110a` text, 44 px, caps 13 px 700) per view, and on BiS there is none: its actions are links. Secondary: 1 px `--line` outline, same size. Link: gold, 600, no underline, `›` suffix where it navigates. `SECONDARY_BUTTON_FIXED` and the existing button classes stay the single source.

### 3.6 Where the character lives
The header selector (`CurrentCharacterChip`, already shipped) is the only place a character is chosen. Pages read it and annotate; no page draws a "Your character" card. On BiS: signed out, one sentence in the summary line ("Sign in to see them against your gear."); signed in, the same line reports the comparison and every slot shows the worn item (section 4.4). The old strip "Sign in or paste an export to see your character here" and the "Your character" card are removed. A pasted export (no account) arrives through the same selector as a temporary character.

### 3.7 Page layouts (names used by the page sections)
- **L1 Ledger:** a control bar, then one plate holding a list or table. BiS, tiers, rankings.
- **L2 Reading:** at 1760 and up `toc 200 | prose 720 | lane (440 to 640)`; the lane carries the section's own data open on S0, sticky inside its section. Below 1760: prose 720 and the toc on the right as today. Guides.
- **L3 Workbench:** `main 1fr | rail 400` at 1760 and up. Sim, guild roster.
- **L4 Overview:** 12 columns, spans 8 and 4, or 4, 4, 4. Home, guild Overview.
Phone: every layout is one column in the order: head, control bar, the primary object, then secondary sections.

## 3A. The shared primitives (the centralization, built with the BiS work)

Owner direction: the centralization is built as part of the BiS work, so later pages are mostly assembly. This section is the contract. Mock library: `design/mocks/ws_prims.py` (each function emits only classes; a page generator may not write a border, background, radius or shadow, and the generator fails if it does). The guild pages (section 6.1) are composed only from it.

### 3A.1 Where things stand (origin/main, measured 2026-11-05)
A shared layer already exists in `web/src/components/ui/`: `Card` (.astro and .svelte), `CardHead`, `IconTile`, `PageHead` (.astro and .svelte), `TabStrip`, `Select`/`SelectSheet`, `StatTile`, `Tag`, `Disclosure`, `ActionLink`, `SettingsRow`, `UiSheet`, `Skeleton`, `EmptyState`, `LoadError`, `Glyph`. It is the layer that draws the card-everything look: `Card` is the gold-edged gradient `card-surface`, `CardHead` is the icon chip and caps title. What leaks around it: 61 component, page and layout files draw `rounded-panel`, 15 use the `card-surface` utility directly, 213 use a `border-line*` class and 72 a `bg-raised` ground; 71 files import `ui/Card`. Named bespoke panels in the BiS and character area alone: `ThisSetPanel`, `StatWeightsPanel`, `PlayItPanel`, `BandPanelBody`, `HealerSetPanel`, `CharacterGearPanel`, `HeaderPanel`, `ArtPanel`, `StatePanel`, `RatingCard`, `ToolCard`, `ForeverScoreCard`, `SimGearCard`, `DetailsCard`, `RotationCard`.
The rule of this spec is therefore not "add a layer": it is **one layer, with a flat surface, and a guard that keeps everything else out of the surface business.** Existing names are kept where the job is the same (`TabStrip`, `Select`, `Tag`, `Disclosure`, `ActionLink`, `Skeleton`, `EmptyState`, `LoadError`, `Glyph`).

### 3A.2 The primitives
Location: `web/src/components/ui/` (existing directory), Astro and Svelte twins where a page needs both, props typed in `web/src/lib/ui/`. All read tokens from `web/src/styles/tokens.css`; none hard-codes a colour, a size or a radius.

| Primitive | Props (name: type) | Tokens read | Used by | Replaces |
|---|---|---|---|---|
| **Page** | `layout: 'ledger' \| 'split' \| 'reading' \| 'overview'`; slots `header`, `bar`, `main`, `rail`; `stickyRail?: boolean`; `railWidth?: 'default'` | `--page-max`, `--page-gutter`, `--page-inset`, `--gap`, `--section-gap`, `--rail-w`, `--toc-w`, `--measure-reading` | Every page. `ledger`: BiS, tiers, rankings. `split` (main + rail 340 / 400): sim, guild, report. `reading`: guides. `overview`: home | Per-page wrapper classes (`reveal flex flex-col gap-[22px] md:gap-8`, `GUILD_BANDLESS_TOP`, hand-written `lg:grid-cols-*`), `Content.astro`'s layout, the `.measure-*` classes. `Base.astro`'s one `.page-column` stays the single container |
| **PageHeader** | `title: string`; `titleColor?: 'class'\|'spec'\|'faction'\|token`; `crest?: Snippet`; `meta?: (string \| Tag)[]` (one `FactLine`); `answer?: Snippet` (a `Figure`); `actions?: Snippet` (Buttons); `tabs?: Snippet` (`TabStrip`) | Title type, `--color-class-*`, `--color-alliance`, `--color-horde`, header gradient | Every page, once | `ui/PageHead` (the eyebrow and `job` paragraph go), `ClassHeader`, `HeaderPanel`, `ArtPanel` (the sky band survives as PageHeader's backdrop and nowhere else), `BisSpecCharacterHeader`, `BisLandingHeader`, `PlannerHead`, `CharacterHeader`'s band |
| **ControlBar** | `groups: { label, kind: 'segmented' \| 'select', options, value, param }[]`; `summary: string` (phone); `end?: Snippet`; `sticky?: boolean` (default true) | `--bar-h`, `--color-line*`, `--color-gold`, control radius | BiS (spec, level, faction, race), tiers (role, faction), rankings filters, sim (fight, source), guild lists | `PresetControl`, `BisBandStrip`, `PresetCaption`, `RankingsFilters`, `ModeBar`, `SettingsBar`, `SourceSwitcher`, `IntervalControl`/`IntervalSegments`, the role and faction pills in `TierStrip` and `TierPage`. Internals reuse `Select` and `SelectSheet` |
| **TabStrip** (kept) | as shipped | as shipped | Route-level tabs: sim tools, guild tabs, report tabs | none; `ControlBar` is for choices on one view, `TabStrip` for views that are routes |
| **Section** | `title: string`; `level?: 2 \| 3`; `meta?: string`; `action?: ActionLink`; `id?: string` | Section head type, `--color-line-soft`, `--section-gap` | Every page: all content below the header and bar that is not the Plate | `ui/Card` + `CardHead` when a card only holds text or rows; `IconTile`; `section-title`; every `<h2 class="section-title">` followed by a bordered box |
| **Plate** | default slot only; `flush?: boolean` | `--color-raised`, `--color-line`, `--radius-panel` | **One per page**: the page's primary object (the doll, the tier list, the raid nights, the result, the guild ranking). The e2e `plates.spec.ts` fails a route with more than two | `ui/Card` (all variants), the `card-surface` and `card-surface-quiet` utilities, `StatePanel`, `ToolCard`'s frame, `BandPanelBody`, every `rounded-panel` wrapper |
| **Row** | `title`, `titleColor?`, `sub?`, `lead?: Snippet \| rank: string`, `trail?: Snippet`, `href?: string` **xor** `action?` , `selected?`, `density?: 'default' \| 'slim'` | `--row-h`, `--color-line-soft`, `--color-hover` (S2), gold selection | Guild lists, rankings, reports, characters, servers runs, signups, loot candidates | `SettingsRow`, `CharacterRow`, `ReportRow`, `RankingRow`, `FeedRow`, `ServerRunRow`, `SignupRow`, `PlannedRaidRow`, `GuildSeatRow`, `UpgradeRow`, `LootCandidateRow`, `CandidateRows`, `ToolCard` as a list |
| **ItemRow** (a Row preset) | `item`, `sourceKind`, `source`, `chance?`, `gain?: Figure props`, `worn?: item`, `badge?: 'upgrade' \| 'worn'`, `slotLabel?` | quality colours, `ItemHover` | Paperdoll cells, loot lists, inspector runner-ups, guide gear tables, consumables | `GearRow`, `ItemPill` (as a row), `RaidLootItem`, `ConsumableCandidates`, `EnchantList`, `GuideBisTable` rows, BiS "The list" |
| **Inspector** | `selection: string`; `head: Snippet`; `sections: { label, body }[]`; `placement: 'side' (>= 1024) \| 'inline' (below)` (automatic) | `--inspector-w` (320 at 1024, 360 at 1760), `--color-line` | Paperdoll first; then roster character detail, tiers spec detail, report fight detail | New; takes over from tall per-row blocks |
| **Paperdoll** | `mode: 'pick' \| 'compare' \| 'worn' \| 'report'`; `slots`; `selected`; `onselect`; `view: 'slot' \| 'where'`; `loading`, `failed` | cell tokens, quality colours | BiS, character page, sim character card, reports | `CharacterGearPanel`, `SimGearCard`, `CharacterStrip` gear, `GearList`, `SlotGrid`, `GuideBisTable`, BiS "The list" (section 4) |
| **Figure** | `value`, `unit?`, `label?`, `delta?`, `tone?: 'strong' \| 'gold' \| 'kill' \| 'wipe'`, `size?: 'sm' \| 'default' \| 'big'`; `FigureRow` lays three out | Figure type, outcome colours | PageHeader answer, roster counts, rating, progression, weights | `StatTile`, the number in `ThisSetPanel`, `ForeverScoreCard`, `RatingCard` figures, `SummaryBar` figures |
| **Meter** | `value: 0..1`, `tone?: 'gold' \| 'kill' \| 'muted'`, `segments?` | `--color-line-soft`, outcome colours | Weights, tier bars, progression, role mix, readiness | The ad-hoc bars in `StatWeights`, `StatWeightsPanel`, `BestPerSpecTable`, `TierRow`, readiness |
| **Field**, **Select**, **FormRow** | `Field`: `label`, `value`, `placeholder`, `error?`; `FormRow`: `fields[]`, `submit` | control radius, `--color-line` | Guild lookup, character lookup, sim paste, login, checkout steps | the hand-styled inputs in `AddonPasteBox`, `ImportBox`, `SaveSimForm`, `ItemSearch`, `GuildJoin`, `CharacterSearch`, `DiscordPasteForm` |
| **Notice** | `text`, `action?`, `tone?: 'gold' \| 'neutral'` | `--color-gold` | A standing line or state under the header | The guild "standing line" and "before Thursday" sentences, "Before raids open" lead, `HelpNote` |
| **Tag**, **Button**, **ActionLink**, **Disclosure**, **Skeleton**, **EmptyState**, **LoadError**, **Glyph** (kept) | as shipped | as shipped | everywhere | none |
| **Overlay** | `kind: 'tooltip' \| 'sheet' \| 'menu'` | `--color-raised`, `--color-line`, overlay shadow | `GameTooltip`, `ItemTooltip`, `UiSheet`, `SelectSheet`, `RequestDrawer` | the overlays' own borders |
| **ItemIcon**, **Crest** (kept: `ClassCrestRing`, `FactionCrest`) | `size`, `quality?`, `badge?` | quality and class colours | everywhere | per-file icon frames |
Retired outright: `ui/Card.astro`, `ui/Card.svelte`, `ui/CardHead.svelte`, `ui/IconTile.svelte` (as a head chip), `lib/ui/card.ts`'s `CARD*` strings, the `card-surface` utilities in `global.css`.

Rules inside the primitives:
- **Row is a link or has an action, never both** (nested anchors are invalid and the action steals the click). A link row can carry a `Tag` or `Figure` in `trail`; a row with a `Button` is not a link.
- **One Plate per page.** Everything else is an open `Section` with open `Row`s (no side padding, hairlines between). Pages that want a second framed object name a reason in review.
- **Section is never inside a Plate.** A Plate holds the primary object's rows, doll or table, and nothing with a heading of its own.
- **Quality, class and faction colours are applied by the primitive** from a prop, never by the page.

### 3A.3 The guard: no component outside the primitives draws a surface
A lint-style test, `web/src/components/surface-guard.test.ts`, in the manner of `site-column.test.ts`. It scans every `.astro`, `.svelte` and `styles/*.css` file. Outside `web/src/components/ui/` it fails on any of: `card-surface`, `rounded-panel`, `border-line`, `border-line-soft`, `border-line-warm`, `bg-raised`, `bg-card-top`, `bg-tile`, `shadow-card`, `border-[var(--color-line`, and `background` or `border` declarations that read `--color-raised`, `--color-card-top` or `--color-line*` in a `<style>` block. Allowed always: SVG strokes in charts (they read the colour by `var`, not by class), and the overlays through `Overlay`.
Because 213 files use `border-line*` today, the test ships as a **ratchet**: `surface-guard.allowlist.json` lists the files that still break the rule, the test fails if a file not on the list breaks it and if the list ever grows, and each page lane deletes its files from the list as it adopts the primitives. By the end of the page-by-page work the list is empty and the allowance becomes a hard fail (done 2026-11-07, the restyle-rest round: the list is deleted and the test fails on any surface outside the primitives). `site-column.test.ts` keeps guarding the one container.

### 3A.3a Parity regions
Every primitive above carries a `data-region` (`head`, `head-title`, `head-meta`, `head-answer`, `controls`, `summary`, `notes`, `plate`, `slots`, `inspector`, `inspector-inline`, `section-<slug>`, `footer-note`), the same names the mock library emits; the table, the tolerance and the gate are in `design/parity.md` (names in code: `web/src/lib/ui/regions.ts`). A page region lands only when its parity spec matches the mock (tenet 15).

### 3A.4 What each page is, in primitives
| Page | Page layout | PageHeader | ControlBar | Plate holds | Sections (open) |
|---|---|---|---|---|---|
| BiS | `ledger` | crest, class-colour Title, answer = set DPS | spec, level, faction, race, view toggle in `end` | Paperdoll + Inspector | Stat weights, guide link |
| Guild Overview | `split` | faction emblem, faction-colour Title, meta, officer actions, TabStrip | none | This week's raid nights | Progression, Loot; rail: Roster, Readiness |
| Guild landing | `split` | Title, one meta line | none | Top guilds ranking | Open a guild (FormRow); rail: Your guilds or Run a guild |
| Tiers | `ledger` | Title, one meta line | role, faction | The tier list | The three notes |
| Guide | `reading` | crest, spec Title | spec | "At a glance" band | every guide section |
| Sim result | `split` | Title, answer = the DPS | fight, source | Ability table | Rotation, Details, Your sims |
| Home | `overview` | Title | none | none (the class picker is Rows of Crests) | Look up, Get started |

## 4. The shared paperdoll

### 4.1 Critique of the shipped Paperdoll
Captures: `design/reviews/2026-10-31-paperdoll/build/character-beta-1440.png`, `sim-raid-1440.png` (the build review's own captures of the shipped component; the character page needs a real character, which production does not offer a public example of).
1. **Two frames around one object.** A card ("GEAR") holds a plate holding the doll. Three borders for one thing.
2. **Names live at the screen edge, 700 px from the crest.** Left column names left-aligned at the left edge, right column names right-aligned at the right edge, a 560 to 700 px empty centre. The eye travels 1,400 px to compare head and hands. Mirroring the game's panel was right for icons and wrong for names.
3. **The centre is the least useful 40% of the panel**: a crest, a name, a level and a pill, all of which the page header already says.
4. **A different unit from BiS.** Gains read "+60.0" and "+148.8 points across 9 slots": scale-factor points. BiS and the sim speak DPS. A player cannot connect "+60.0 points" to "+11 DPS".
5. **Two lines of different colour per slot** (worn name, then the pick's icon, name and gain) with no source, so it cannot answer "where do I get the upgrade".
6. **Not interactive beyond a tooltip.** Nothing to click for runners-up, sources or drop chances.
7. **Phone**: the doll folds to two mirrored columns plus full-width "upgrade rows" under each pair, the tall-block problem again.
8. **Hard-wired to the character page card**: the doll draws "the frame only" and each caller wraps its own Card, so the page already has a card around it.
9. **It stays correct.** The model (`lib/gear/paperdoll.ts`: slot layout, the two-hander covering the off hand, unknown-item handling, "Not scored", the item-level ruling) and `ItemHover` tooltips are right and are kept.

### 4.2 One component, four modes
`web/src/components/character/Paperdoll.svelte` keeps its name, model and tooltips and changes its frame. Props gain `mode` and the inspector. No second component.

| Mode | Used by | A cell shows | Right of the cell |
|---|---|---|---|
| `pick` | BiS, signed out | The BiS pick: icon, name, source line | `+1.8 over next` (lead over the best runner-up) or `Tie next best` |
| `compare` | BiS, signed in | The BiS pick as in `pick`, with the worn item's 28 px icon at the right | `+1.8 DPS` in gold for an upgrade; a green check and `Worn` when the player has the pick |
| `worn` | Character page, the sim's Gear section (the frame on both, owner ruling 2026-11-11; the sim's saved-runs strip keeps the list form) | The worn item: icon, name, enchant line | Upgrade `+1.8 DPS` with the pick's icon, or `Best in slot`; `Not scored` when unscored |
| `report` | Reports (GearList) | The worn item and its detail line ("Item level 66 · Enchanted") | none |

The frame, the cell, the inspector, the states and the keyboard model are identical in all four. A caller passes `PaperdollSlot[]` (existing type) plus, for BiS, a `pick` and `alternatives` per slot; `PaperdollUpgrade.points` becomes `dps` (section 4.6).

### 4.3 Frame and cell
- **No card around it.** The page owns the one plate; the doll is the plate's content. Callers stop wrapping it.
- **Two columns, icon-first, everything left-aligned.** Left: head, neck, shoulder, back, chest, wrist, main hand, off hand, ranged (9). Right: hands, waist, legs, feet, ring, ring, trinket, trinket (8). Shirt and tabard appear in the left column after chest only when the data carries them (as today). Reading order down the left then down the right matches the character panel's head-to-feet order. Weapons end the left column where the game puts them at the bottom. Rationale for dropping the mirror: names and numbers align in one vertical line per column.
- **Cell:** grid `60 px label | 40 px icon | 1fr body | auto gain`, min height 58 px, padding 8 x 16. Label: slot name, 13 px 600 `--muted` (above the name on columns narrower than 400 px, caps 12 px). Icon: the HD icon in a 1 px quality border (`--color-rarity-*`), 5 px radius. Body: line 1 item name 15 / 20 700 in the AA-safe quality colour; line 2 source, 13 / 18 `--muted`, with a source-kind glyph (`SourceKindIcon.astro`): `Boss · Instance · 40%` for a drop (the chance is shown when known), `Crafted · Leatherworking`, `Silverwing Sentinels reputation`, `World drop`, a quest's name. The source is one line. When a drop's line is too long for the cell, only the boss name gives way (an ellipsis, the full line in the cell's `title` and in the inspector); the instance and the drop chance never truncate. With a 320 px inspector and 12 px cell padding nothing truncates at 1440 on the 14 picks of the mock (`web-bis-1440-fold.png`); the rule is for longer names and for 1280. Gain: Figure 16 (Barlow 700 tabular), right-aligned, with a 12 px unit line under it.
- **Gain colouring:** gold only for an upgrade on a signed-in doll; `--strong` for a lead over the next pick; `--muted` for a tie; `--kill` for worn. Colour is never the only signal: the sign, the word under the figure (`over next`, `DPS`, `Worn`) and the icon badge carry it.
- **Badges (compare mode):** an 18 px disc on the icon's lower right: gold `▲` upgrade, green `✓` worn. The worn item's 28 px icon sits left of the gain with the real item tooltip on hover.
- **No "NEW", no sim-check glyph on the row.** "New at this level" moves into the inspector as a line ("New at levels 20 to 29") and is shown on the row only when the band's previous pick differed and the player can see the change ("Changed from Trapper's Leather Hood" in the inspector). Sim-check ("Sim-checked against Trapper's Leather Hood: +0.6 DPS") lives in the inspector's runners-up section.
- **Hierarchy:** the item name is the one coloured, 600-weight element; the gain is the second; the source is muted. One level of emphasis per cell.
- **Two-handed main hand:** the empty off hand reads "Two-hand in main hand" with the two-hander greyed in the socket, as shipped. In `pick` mode an empty slot reads the data's own reason (shipped copy: "No trinket you can get at 20 to 29 raises your damage. The first that does comes at 40.", "Impaling Harpoon is a two-hander, the off hand is taken."). The mock shows a generic line.

### 4.4 Cell states (all verified in `web-bis-states.html`)
| State | Treatment |
|---|---|
| Default | S1 ground, 1 px `--line-soft` divider under the cell |
| Hover | S2 ground `#131824`; over the icon or name, the shared `ItemHover` tooltip after 150 ms (the real item: stats, item level, required level, source) |
| Focus-visible | 2 px gold outline inset 2 px; `Tab` enters the doll once (roving tabindex), arrow keys move in reading order (down, then across to the next column), `Home` and `End` jump to first and last slot |
| Active (pressed) | `#1a2030` |
| Selected | `#e5b9551a` ground and a 3 px gold bar on the cell's left edge; exactly one cell selected; the inspector shows it. Selection is `aria-current="true"` |
| Disabled (empty slot) | Dashed 40 px socket, "Nothing better yet" in `--muted`; not focusable, not selectable |
| Unknown item (worn id missing from the data) | Dark red-tinted socket, "Item 288887", "Not in our item data yet" (shipped copy), selectable, the inspector says what the id is |
| Loading | 17 rows at the final 58 px height: 40 px icon block, two skeleton lines, a figure block (`Skeleton`); the inspector shows the same; no layout shift |
| Error | One row in the plate: "We could not load the gear list." and a Try again button (`LoadError`) |
| Touch | Tap selects (on phone, expands in place); a second tap collapses. Tooltip on long-press, as today |

### 4.5 The inspector (the side panel)
A column inside the same plate, divided from the doll by one `--line` hairline. Width `--inspector-w`: 360 px from 1760 only (440 until 2026-11-06; see 5.13). Below 1760 there is no side panel: the selected slot expands in place (rev 3, so the cells keep the full width they need for the persona fixes at 1440). Sticky at 60 px below the top, `max-height: calc(100vh - 72px)`. Always showing a slot: on load the first slot with an upgrade (signed in) or Head (signed out), so the panel is never empty at 1024 and up.
Content, top to bottom, hairlines between sections:
1. **Head:** 56 px icon, slot label (12 px caps), the item name 18 / 22 700 in quality colour, "Item level 25 · requires level 20".
2. **Stats:** one line, 14 / 22 ("+10 Stamina · +8 Agility"), plus the item's effect text if it has one.
3. **You wear** (compare mode only): 32 px icon, name, and "The pick" or "Upgrade to the pick for +0.8 DPS".
4. **Where to get it:** the pick's full source: "Wailing Caverns · Lady Anacondra · 40% drop"; for crafted "Leatherworking · Recipe from ..." when the data has it; for reputation the standing required; a quest's giver and zone (shipped copy rules, `lib/bis/`). Drop chance in `--strong`.
5. **Runners-up · DPS against the pick:** the pick pinned first ("Pick"), then up to six alternatives by DPS delta: 32 px icon, name (quality colour), source one line, delta in Figure on the right (`−1.8` in `--wipe`, `same` in `--muted`). Hover on a row shows the item tooltip. Source and chance appear on each row. This replaces the tall per-slot blocks: runners-up are loaded and drawn only for the selected slot.
6. **Link:** "Compare in the simulator ›".
Phone: no side panel. The same content expands in place under the tapped cell (the head section hidden because the cell above it already shows the item), a 3 px gold left bar joining it to the cell; `web-bis-open-390.png`.

### 4.6 Unit and data (decisions for engineering and the owner)
- **Gains in DPS everywhere.** `PaperdollUpgrade` carries `points` today (scale-factor points) and the doll says "+148.8 points". Convert with the spec's own `dps_per_point` per weight (already in `bis/*.json`), so BiS, character, sim and report all read "+1.8 DPS". Requires `lib/gear/paperdoll.ts` and the three mappers to change; the existing rules (a pick must gain at least the adoption margin, unscored slots) are unchanged. Decision D4.
- **Lead over next** (`pick` mode) is the smallest absolute `dps_delta` among the slot's runners-up, from the data as shipped.
- **Drop chance** from `loot.json` `item_chances` (percent already; show whole percent from 1, one decimal below 1).
- **Compare** uses the existing planner gain rule (`lib/home/upgrades.ts`), unchanged except for the unit.

### 4.7 "Where to go": evaluated, recommended
The shipped page already carries this data as a card at the bottom. Value: when levelling, the question after "what is best" is "which dungeon do I run for it". At 20 to 29 the grouping says Wailing Caverns holds 5 of the 14 picks and The Deadmines 2; that is a farming plan, not a list. **Recommendation: ship it as the second view of the same plate, not a tab and not a card.** A two-item segmented toggle at the right of the summary line: `By slot` (default) and `Where to go`. Same plate, same inspector, same selection.
- Groups, sorted: dungeons and raids first by number of picks (descending), then crafted (by profession), reputation (by faction), quest, vendor, world drop, PvP.
- Group head: kind glyph, name in H3 (18 px), count ("5 pieces"). Rows: slot label, 36 px icon, item name, and the boss and drop chance for a drop ("Lady Anacondra · 40%") or "Requires level 18" otherwise.
- Two columns (CSS columns, no group split across a break) from 1024 up; one on phone.
- The toggle's state is in the URL (`?view=where`).
- Not a default: the question "what do I wear" comes first, and a player on a PvP or crafting route reads By slot.
Mock: `web-bis-where-1440.png`, `-1920.png`.

## 5. Best in slot page

### 5.1 The one question, and the answer visible without scrolling
**"What do I wear?"** The whole set is the answer. Rev 3 note: the persona fixes (5.10) add stats, full source detail, next best and a level flag to each row, so a row is 3 to 4 lines (about 80 px). Measured on `web-bis-slots`: the doll runs y 407 to 1,127 at 1440 and 407 to 1,145 at 1920 and 2000, so the first 12 of 17 slots (head to the middle of the right column) are in the first screen of a 1080 monitor and all 17 need about 150 px of scroll. The earlier one-screen claim (58 px rows) held only for the stripped-down row. This is the density cost of the persona asks; the alternative is to put next best and the worn item behind the inspector, which they asked us not to do. Owner decision D10.

| | 360 / 390 phone | 1280 / 1440 desktop | 1920 / 2000 |
|---|---|---|---|
| Visible without scrolling | The set's DPS, the control summary line, the summary line and the first 5 slots (head to chest; the fifth ends at y 827 of an 844 px screen) | The set's DPS, the control bar and the doll. At 1440 the doll ends at y 844: on a 1440 x 900 monitor in a maximised browser (viewport about 790 px) the last one or two rows of the left column are below the fold and 15 of 17 slots are visible; at 1440 x 1000 and up all 17 are. The inspector shows Head | All 17 slots and the inspector end at y 829, inside the 1,000 px first screen; the weights' heading starts below the fold |

At 1920 the first 1,000 px of the shipped page show 3 cards and no slot; the mock shows every slot. `compare-bis-1920.png`.

### 5.2 Reference
Wowhead's class BiS tables (row per slot, source on the row) and the in-game character panel (one object, slots in columns). See section 2 for what is borrowed. In `web-bis-1920-fold.png` the page reads like the panel a player already trusts: icons in quality borders, names, where, and the gain.

### 5.3 Layout, region by region (components named)
1. **Header** (`Header.astro`, 64 px; the 80 px of today is reduced, `nav.css` bar height token). The header selector shows the current character (compare mode) or "Sign in".
2. **Page head** (`ClassHeader.astro`, no art band): 56 px class crest (`ClassCrest.astro`, the one crest language), the Title "Beast Mastery Hunter" (Marcellus, 40 / 46) in the class colour, one Small line under it. At the right, the answer: label "This set, simmed 5 Nov", Figure 36 px "110.9 DPS" (Barlow 700 tabular), two links. Padding 16 / 14. No eyebrow, no description paragraph (the shipped sentence "The gear that raises your damage most..." is removed; the page head says what the page is).
3. **Control bar** (`ControlBar`, section 3.4): spec segmented control with the HD spec icons at 22 px, level, faction, race select.
4. **Summary line** (15 / 24): the comparison sentence on the left, the `By slot | Where to go` toggle on the right. 10 px padding.
5. **The plate** (S1): `Paperdoll` (section 4) with its inspector.
6. **Stat weights** (open on S0, `StatWeightsPanel.astro` flattened): H3 "Stat weights", Small caption, five columns (one per stat: name 15, Figure 24 scale, a bar, "0.215 DPS per point"), then one line with the addon link. Hit-to-cap and the sim-error figure move into the stat's tooltip. At 1760 and up the five columns fill the 1632 px line; on phone they stack as rows.
7. **One line:** "Talents, rotation and stat priority for this spec: Beast Mastery Hunter guide ›". Replaces Play it. Link, never restate.
8. **Footer** (full-bleed, unchanged).
Retired from this page: This set card (its number is the answer in the head), Your character, Play it, The list's per-slot blocks, the Where to get it card at the bottom (replaced by the toggle), the "New at this level range" card (becomes a line in the summary line: "Changes from 20 to 29: see 30 to 39 ›" shown only when a next band exists).

### 5.4 Hierarchy
Primary: the doll (S1 plate). Next: the set's DPS in the head. Secondary and visibly so: the control bar (outlined segments, no fill but the selection), the weights (open, no plate), the guide line (Small). One gold block per region: the selected control, the selected cell, the upgrade figures.

### 5.5 Copy (verbatim)
- H1: `Beast Mastery Hunter`. Sub: `Best in slot · levels 20 to 29 · Alliance Dwarf`.
- Answer: label `This set, simmed 5 Nov`; `110.9 DPS`; links `Open in simulator`, `Talents in planner`.
- Control groups: `Spec`, `Level`, `Faction`, `Race`. Levels `20-29`, `30-39`, `40-49`, `50-59`, `60`. Factions `Alliance`, `Horde`. Phone summary: `Beast Mastery · 20-29 · Alliance · Dwarf`, `Change`.
- Summary, signed out: `14 of 17 slots have a pick at this level. Sign in to see them against your gear.` Signed in: `Felsnap wears 5 of 14 picks. 9 upgrades, +5.1 DPS available.` Signed in, all best: `Felsnap wears every pick for this level.` Toggle: `By slot`, `Where to go`.
- Cell gain captions: `over next`, `next best` (with `Tie`), `DPS`, `Worn`. Badge titles: `Upgrade`, `You wear the pick`.
- Inspector headings: `Where to get it`, `Runners-up · DPS against the pick`, `You wear`; row labels `Pick`, `same`; link `Compare in the simulator ›`.
- Where to go group counts: `1 piece`, `5 pieces`; non-drop row tail `Requires level 18`.
- Weights: `Stat weights`; `Normalized to Agility = 1.00`; `0.215 DPS per point`; `The addon rates every item in your bags and tooltips with these weights.`; link `Use these weights in the addon ›`.
- Guide line: `Talents, rotation and stat priority for this spec: Beast Mastery Hunter guide ›`.
- Loading: skeleton only, no text. Error: `We could not load the gear list.` with `Try again`. Empty slot and unknown item: section 4.3 and 4.4.
- No provenance pills (`feedback-provenance-quiet`); the one "simmed" stamp is in the answer label.

### 5.6 Phone and desktop behaviour
- **Never collapses:** the doll's 17 slots, the selection state, the set DPS.
- **Collapses:** the control bar (to the summary line), the inspector (to in-place expansion), the weights (five columns to five rows), the page head's answer (to one line: figure and links beside each other, label hidden).
- 768 to 1023: one column of cells, no inspector (expansion in place). 1024 to 1759: two columns, no side inspector (the selected slot expands in place under its cell, as on phone). 1760 and up: two columns plus a 360 px inspector, wider cells.
- At 2000 the column is 1792 wide, 104 px margins; the plate, the bar and the head share one left edge (`web-bis-2000-fold.png`).

### 5.7 Performance and polish limits
- No layout shift on hydration: the skeleton is the final geometry (17 x 58 px; inspector reserved at its width); icons carry width and height; the control bar and head are server-rendered from the URL.
- Fewer nodes: 17 cells and one inspector instead of 17 blocks with 1 to 3 runners-up each (54 list icons today). Runners-up icons load when a slot is selected. Measured on production: `main` holds 96 sections and 1,017 `img` elements because every level band is in the DOM, hidden; the redesign renders the selected band only (R1 makes the band a URL parameter, so the other bands need not ship).
- No missing-icon flash: icons are inlined as the shipped `dataUrl` icon set; an absent icon draws the dark socket, never a broken image.
- Lighthouse budgets in `web/lighthouserc.json` unchanged; the BiS URLs (`/bis/hunter/marksmanship.html`) are already in the run.
- Tests that must change with the build: `paperdoll.spec.ts` (selection, keyboard, the unit), `bis` e2e for the URL state (R1), the `site-column` test for the token.

### 5.8 Requirements the build must meet
R1. Band, faction, race and view are query parameters, read on load and written on change. `?band=60` selects 60. (Production today ignores it.)
R2. Selecting a slot updates the inspector without a network request: the alternatives are in the page's data.
R3. Compare mode needs no new request: the worn items come from the current character already loaded by the header.
R4. Every item in the page opens the shared tooltip on hover, focus and long-press, on this page and in the inspector (tenet 2, 9).

### 5.9 Acceptance screenshots the implementer attaches
Viewports 390, 1440, 1920, 2000 (and 1024 for the no-inspector breakpoint) for: signed out default; signed in compare; Where to go; a slot selected on phone; controls expanded on phone; loading; error; an unknown worn item; a two-hander with an empty off hand. Plus the full-page side-by-side with `compare-bis-<width>.png` at the same scale and the focus-visible state on a cell and a segmented control. The reviewer opens the 2000 capture first for alignment.

### 5.10 Persona-review fixes (rev 3, 2026-11-05): a level 24 leveler, a min-max raider at 60, a casual altoholic on phone
Mocks: `web-bis` (band 20, signed out), `web-bis-slots`, `web-bis-you` (signed in), `web-bis-you-where`, `web-bis-60`, `web-bis-60-you`, `web-bis-weights-open`; renders at 390, 1440, 1920, 2000. Row anatomy replaces 4.3's: label column 76 px (slot name and the level flag), 40 px icon with badge, body, gain column.
1. **Gains in plain units and basis.** The gain column reads `+1.8` over `DPS vs next best` (signed out); `+0.8` over `DPS vs your Slime-encrusted Pads` (signed in); a tie reads `Same as next best` (with `=` as the figure), never "Tie". The body's last line spells the basis: `Next best: Trapper's Leather Hood −1.5 DPS`.
2. **Required level on every row, flagged.** Under the slot label: signed out `Level 20`; signed in `Usable now` (green) or `At 25` (gold) against the header character's level. Sources carry level ranges: quests `requires level 14` (from `loot.json` quests), dungeons `(levels 17-24)`.
3. **Full source detail, wrapping, never truncated.** Drop: `Boss · Instance (levels a-b) · 40% drop` (chance from `loot.json` `item_chances`); world drop: `World drop · bind on equip · also on the auction house`; crafted: `Crafted · Leatherworking (skill 100)`; reputation: `Silverwing Sentinels reputation · Honored`; quest: `Quest: name · requires level 14`; vendor: `Vendor: name`. The line wraps to as many lines as it needs on every width; the 4.3 ellipsis rule is removed.
4. **Inline short stats** after the name: `+8 Agi · +10 Sta` (13 px muted, three stats; the abbreviation map is `ABBR` in `ws_bis.py`).
5. **Runners-up inline:** `Next best: <name> −0.4 DPS` on every row, in quality colour, and `(same DPS)` for ties.
6. **Trust line under the DPS**, 13 px: `Level 29 Dwarf · 11 talent points · training dummy, 180 s · no buffs or consumes · 1,000 runs · simmed 5 Nov · View ranking sim ›`; at 60: `... Raid-ready, Phase 1: 15 buffs, 9 debuffs, 6 consumes ...`. Buff counts come from `presets.raid`; `1,000 runs` and the 180 s fight length are data gaps (5.12).
7. **Band 60 frame** (`web-bis-60`): the enchant per slot in green after the stats (`Superior Agility`), a trinket section in the inspector (both picks, and the ranking note), a weapon note (`Legionite Glaive is a two-hander: the off hand stays empty`), the hit-cap state under the summary (`Hit: 7.9% of the 9% cap, 1.1% short. Hit rating is the first stat to stack.`, green when capped).
8. **Signed in** (`web-bis-you`): the header chip shows `Felsnap 24` at 390 (the search button hides below 480 and the wordmark shrinks to make room); a green tick badge and `You have this` on slots already worn; the summary line counts upgrades (`Felsnap (level 24) has 5 of 14 picks. 9 upgrades, +5.1 DPS available.`); an upgrade row's last line shows the worn item with item level and stats and the stat delta: `You wear Slime-encrusted Pads (item level 25, +3 Mp5). The pick is +0.8 DPS, −3 Mp5, +5 Agi.`
9. **Empty slots explained** with the shipped copy patterns, data-driven: `No trinket you can get at 20 to 29 raises your damage. The first that does comes at 50.` (found by scanning the later bands); `Impaling Harpoon is a two-hander, the off hand is taken.`
10. **Default view:** `Where to go` below 60, `By slot` at 60 (URL `?view=` still overrides).
11. **Stat weights fold.** Section "What to stack": one sentence (`Stack agility, then hit rating, then ranged attack power.`) and a disclosure "Show the weights and their error"; open it shows the table with a bar, the scale, DPS per point and the error column and a `Copy weights for the addon` button (no Pawn string). `web-bis-weights-open-1440`.
12. **Dev text removed.** No mock footer line; the page head's sub is the single word "Best in slot" (the level, faction and race live only in the control bar, so nothing is stated twice on desktop; on phone the bar's summary line states them once).

### 5.11 Source filters
- **Control:** a `Sources` group in the `ControlBar` (after Race): a multi-select button reading `All sources`, `7 of 8 sources` or `Dungeon only`, opening a checklist (`web-bis-sources-1440`): Dungeon, Raid, World / BoE, Quest, Crafted, Reputation, Vendor, PvP, each with a count (slots with a ranked option from that source) and a footer line: `Counts are slots with a ranked option from that source. Turning one off re-picks each slot; nothing is hidden.` On phone it is a row of toggle chips (each with its count) in the expanded control stack (`web-bis-bar-390`); the collapsed summary appends the selection.
- **Re-pick, not hide.** With a kind off, each slot re-picks from the kinds still on and shows its own gap to the best pick (`−1.8 DPS vs the best pick`, red) and its source; the line under it says what changed (`Re-picked (dungeon only). Best overall: Brawler's Leather Hood (crafted), +1.8 DPS.`). A slot with no option from the chosen sources stays visible and explained: `No option from the sources you chose. None of the ranked neck items comes from dungeon. Best overall: Sentinel's Medallion (reputation).`
- **The set's number** becomes `≈ 105.1 DPS` with `Set: dungeon only` and `5.8 DPS under the best set; the sim re-runs this set`. The mock sums slot gaps as a stand-in; the build shows a loading state and then the re-simmed set's DPS (data gap).
- **Leveling defaults:** below 60, Reputation is off (a standing above Friendly is out of reach for most levelers) and out-of-reach crafts are off; the summary shows `Reputation above Friendly is off at levels below 60, so the picks are re-picked without it. Turn Reputation on in Sources to include it.` At 60 everything is on.
- **State:** the selection is in the URL (`?sources=dungeon,quest`), default per band.
- **Mock:** `web-bis-dungeon` at 1440 and 390 ("Dungeon drops only": 11 of 14 slots have a dungeon option in the ranked items, three explain why they have none). The ranked options in the data are each slot's pick plus up to three runners-up, so "no option" can mean "none in the top four"; the build computes the best option per kind over the full candidate list (data gap).

### 5.12 Data gaps behind the mocks (named, not built)
| Shown | Needs | Where from |
|---|---|---|
| Dungeon level ranges (`levels 17-24`) | `instance_levels` per dungeon | classic-db / AreaTable into `loot.json` sources (mock uses sample ranges) |
| Crafted skill (`skill 100`) | recipe skill per crafted item | recipe tables (`SkillLineAbility` already used for trainables) (mock uses sample skills) |
| Bind on equip, auction house | bind type per item | the client item table (`Bonding`) |
| Reputation standing (`Honored`) | standing per reputation item | `loot.json` `factions` (shipped page already says "honored") |
| Enchant per slot at 60 | the sim's best enchant per slot | the enchant sim; the mock scores `enchants.json` on the band's own weights (indicative) |
| Trinket ranking | all sourced trinkets simmed per spec (144 eligible, 74 sourced) | the BiS sim |
| Trust line: `1,000 runs`, `180 s` fight | `iterations`, `encounter_seconds` in `bis/*.json` | the BiS sim header |
| Filtered set DPS and best option per kind | a per-kind best over the full candidate list, and a re-sim of the filtered set | the BiS sim, on demand |
| Worn item's stat delta, flags `Usable now` | the header character's gear and level | already loaded for the header selector (R3) |

### 5.13 Build decisions (2026-11-06, coordinator for the owner)
- **Row anatomy, option C (replaces 4.3's cell and 5.10's line 2).** Line 1: slot label and level flag, icon, item name (never truncated) and its short stats (two, which ellipsize from 768; all stats in the title and the slot detail), and the gain figure with its unit at the right. Line 2: the full source alone, never truncated, in short words that keep every fact (`Lady Anacondra · Wailing Caverns (15-24) · 40%`, `World drop (53-61) · BoE · on the AH`, a one-level range printed once). The gain column (144 px, 112 on a phone) carries what the figure is against on one line that may ellipsize, the full words in its title: `vs <next best>`; a tied pick reads `Same` over `next: −1.7 vs <next different>` or `No close alternative`; signed in `vs your <worn item>`, `Worn` over `You have this`; a re-picked slot `vs best: <pick>`. The enchant, the worn item's facts and the re-pick explanation live in the slot's detail. Where a source or an empty slot's reason still cannot fit, the row takes a third line (measured on the real data across 28 specs at bands 20, 40, 60).
- **Inspector 360 px from 1760** (was 440): rows beat inspector width.
- **Sources** is a compact `All sources` / `Sources (N)` button opening a popover (Overlay) at desktop; chips on phone. The BiS bar's group labels show from 1920.
- **Type:** no condensed face anywhere; every heading is Barlow 700; the wordmark is an SVG (Cinzel no longer loads); the site loads Barlow 400/600/700 and, on primitive pages, Marcellus 400.
- **Healer and tank heads kept (supersedes 5.3 item 2's one figure for those roles).** A healer's head shows its HPS figure and, under it, Overheal, Mana lasts and Healing per mana; a tank's head shows four figures (Effective health, Damage taken per second, Chance of death, Threat per second) and the TMI line in place of the one figure. Both shipped and `bis-healer` / `bis-tank` e2e assert them; the parity mocks (`ws_bis.py`, `web-bis-healer-60`, `web-bis-tank-60`) carry the same blocks, order and words, no disclosures.

## 6. The other pages: direction, mocks to follow one at a time
Each keeps the system above. None is mocked in this round.
- **Guide (L2).** The hero's three boxed cards become one band (S1, full-bleed hairlines) split by vertical rules: talents, rotation, stat priority, open inside. Below, `toc | prose 720 | lane` at 1760 and up; the lane carries the section's own build, cooldowns or weights, open and sticky in its section. BiS in the guide is the shared Paperdoll in `pick` mode (no second BiS table). A first width-only mock exists: `design/mocks/wide/wide-guide.html`.
- **Sim with a result (L3).** The run block compresses to one line (character from the header selector, fight, Run). The result leads: the DPS figure at 56 px with its error and the change from the last sim, then the ability table on S1 with bars at wide widths, with a rail (rotation, details, your sims) open on S0. A width-only mock exists: `wide-sim-report.html`.
- **Tiers (L1).** One control bar (role, faction); the three honesty paragraphs become a one-line caption and a three-column open note under the list; the list is one plate with the letter column; the bar column widens to 560 px at 1920. Width-only mock: `wide-tiers.html`.
- **Guild landing and Overview.** Mocked in section 6.1 from the primitives alone.
- **Home (L4).** The class picker is the hero (nine crests at 72 px under the H1); Look up a character is one input row; Get started is a four-item open row; the launch countdown is open text at the right. Signed in, the header selector and `HomeAccountPanel` lead with the character's next upgrade, flattened.

### 6.1 Guild landing and Overview: the proof that the primitives carry over
Mocks: `web-guild.html`, `web-guild-landing.html`, `web-guild-landing-out.html` at 390, 1440, 1920 and 2000, each with a before/after (`compare-guild-*.png`). Built only from `ws_prims.py`; the generator asserts that no page writes a border or background. Data: the shape of `api/cmd/seedguild` read from `roster.go`, `lifecycle.go` and `readiness.go` (24 characters: one tank, six healers, seventeen DPS, three awaiting approval; Emberfall, Pinefeather and Shiverknife are the seed's own worst-readiness raiders; Onyxia is the only named encounter; 22 drops). Not run against the database (it needs the API's migrations and Docker); the seed's data shapes are reproduced, and raid nights fall after the 9 December 2026 opening as the guild spec requires. Characters shown: the seed's owner main "Dunmaro", a Beast Mastery Hunter.
- **Overview, as shipped:** five Cards (Roster, This week's raid nights, Progression, Readiness, Loot), each a gold-edged box with an icon chip, a caps title and "See all"; the Roster and Progression cards hold one sentence each; the page head, an officer strip and a tab row are three more blocks. Production has no public guild to capture, so `web-guild-before.html` reproduces `GuildOverview.svelte`'s own cards, copy and order in the shipped look at the 1344 column.
- **Overview, after:** `PageHeader` (faction emblem, "Olympus XXVII" in the alliance colour, one meta line, the two officer actions, `TabStrip`), a `Notice` (the standing line, verbatim pattern from the guild spec), then `Page split`: main holds one `Plate` (This week's raid nights, three `Row`s with a day `Figure` and a kill or wipe `Tag`) and two open `Section`s (Progression: a `Figure`, a `Meter`, one `Row`; Loot: one `Row` with an Award `Button`); the rail holds two open `Section`s (Roster: three `Figure`s and a role-mix `Row` per role with a `Meter`; Readiness: three `Row`s with `Crest` leads, worst first). One framed object instead of five, the right 4 columns used at 1920 and 2000, the page's content ends at y 897 at 1440 and y 929 at 1920 and 2000, inside the first 1,000 px (1,631 at 390), where the shipped structure ends at y 1,293 at 1440, 1920 and 2000 and 1,392 at 390 (`web-guild-1440-fold.png`, `compare-guild-overview-<w>.png`).
- **Landing:** `PageHeader` (Title "Guilds", the shipped lead sentence), `Page split`: main holds an open `Section` "Open a guild" (a `FormRow`: guild name `Field`, region and ruleset `Select`, one primary `Button`) and "Top guilds" (one `Plate`, ranked `Row`s with a `Meter`); the rail holds "Your guilds" (signed in: an open `Row` with a Tag "Officer") or "Run a guild?" (signed out: the shipped sentence and the sign-in button). The two states differ by that rail section only, so they are one template.

**What did not translate cleanly, and what the system now says about it**
1. **The lookup form is not a control bar.** `ControlBar` is page-level, sticky and holds choices; the guild lookup is a text field, two selects and a submit. Added `Field` and `FormRow` (3A.2). `ControlBar` keeps its job.
2. **A Row that is a link cannot hold a Button.** The loot row's Award button nested inside a link row (invalid anchors, and the click went to the wrong target in the first render). Rule added to `Row`: `href` xor `action`.
3. **Three framed objects wanted to be Plates.** Raid nights, progression and loot each looked like "a list in a box". Resolved by the one-Plate rule: only the primary object is framed; the others are open `Section`s with open `Row`s. This tightened the rule and fixed the mock; the guild landing follows it too.
4. **The standing line and "before Thursday" sentence had no home.** Today they are paragraphs between the header and the tab strip. Added `Notice` (3A.2).
5. **The page Title colour has no class.** The Title rule is "in the spec or class colour"; a guild has neither. The mock uses the faction colour (alliance blue `#6fb1ff`, horde red `#ff6b5c`). Decision D8.
6. **A guild crest is not a class crest.** The mock uses the faction emblem in the one circular crest frame (`FactionCrest`), because the shipped guild crest settings (`GuildCrestSettings`) are an upload feature outside this round. `PageHeader.crest` takes whichever the page has.
7. **Officer actions live in the header.** The shipped officer strip (claim badge, approvals, invite) is a band of its own. It fits `PageHeader.actions` as two buttons; the claim badge (a state) would be a `Tag` in `meta`, which `meta: string[]` does not allow. `meta` therefore accepts `Tag`s (added to the contract).
8. **The seed was not run.** Seeded data is read from the seed's Go source rather than the database, as above.
9. **No new primitive was needed for the page body.** Guild needed `Row`, `Figure`, `Meter`, `Tag`, `Button` and `Section`, all defined for BiS or its generalisation (`ItemRow` became a preset of `Row`), plus the three additions above (`Field`/`FormRow`, `Notice`, `Emblem` via `Crest`).

## 7. Guild experience (persona round: GM, raid leader and loot officer, member)

Mocks (1440, 1920, 390; renders in `design/mocks/renders/web/`, key copies in `design/reviews/2026-11-05-wide-layout/after/`): `web-guild-member.html`, `web-guild-officer.html` (the landing as each role), `web-guild-raid.html` (planner and sign-ups), `web-guild-loot.html` (council), `web-guild-report.html`, `web-guild-recruit.html` (public card and application form), `web-guild-applications.html` (officer votes), `web-guild-profile.html` (member, with private officer notes). Compare images `compare-guild-member-<w>.png` and `compare-guild-officer-<w>.png` set the role-aware landing against the shipped five-card Overview. Generator: `design/mocks/ws_guild2.py`, composed only from `ws_prims.py`; the earlier `web-guild.html` (round 1) is superseded by the two role landings.

### 7.0 What already exists (read before inventing; `api/internal/guilds/` on origin/main and the 2026-10-26 to 2026-11-05 guild specs)
- **Raid posts and sign-ups exist,** on the site and as Discord buttons, through one function (`raids_signup_service.go`, `SignUp`): the states are **In, Late, Bench** (`raids_signups_store.go`), with a role and a raid spec per row (`raids_spec.go`); `raids_discord.go` posts and updates the channel message; `raids_roster.go` holds the snapshot (name, class, spec).
- **The comp vocabulary is shipped:** a class tally row, roster groups Tanks, Healers, Melee, Ranged, "23 of 40 signed up" (`design/specs/2026-11-05-guild-raid-comp.md`).
- **Loot exists:** Onyxia's 22 drops, candidates in two tiers (BiS-band match with `gain_dps`, then a class-can-use fallback with `ilvl_delta`), awards in `loot_awards` (`loot.go`, `design/specs/2026-11-05-guild-loot.md`).
- **Readiness exists:** `gear_gap`, `enchants.checked`, `consumables.state` (presence of any bagged item, not flask identification), professions; `needs_before_next_raid` (`readiness.go`, `me_standing.go`).
- **Attendance and the raid report exist:** auto attendance from logs, kill order, deaths, parses; the report's consumables column is null because the log carries no buff events (`raids_attendance*.go`, `raids_report.go`).
- **Does not exist:** applications, officer notes, loot votes, MS/OS flags, raid assignments, buff-cover logic, comp targets, resistance totals.
The brief's "Accept / Tentative / Decline" is not the shipped vocabulary. The mocks use **In / Late / Bench** (decision D11).

### 7.1 The landing is role-aware (`/guild`, Overview)
Same `Page split`, same `PageHeader` (faction emblem, faction-colour Title, one meta line that now carries the next raid, `TabStrip`), different sections by role. Everyone gets the next-raid Plate with the viewer's own row: character, spec and level, their state, a three-way `Segmented` (In / Late / Bench) that writes the same sign-up the Discord buttons write, the raid in the viewer's local time with the countdown (`2d 4h`).
- **Member** (`web-guild-member`): a `Notice` that states the verdict in words, `Ready to raid`, `Almost there` or `Getting started` (never "below floor"); the next-raid Plate; **Before the raid**, a private `Checklist` (level, attunement, gear gap, enchants, flask or food, fire resistance) where each failing item links to its fix (the BiS list, "What to bring"); **Your numbers** (raids attended, items this tier, item level, plain numbers); **Officers asked you to** (open `Row`s); a rail with **Your class in the guild** (a `Tally`: "Priests 2/3") and Last raid and Progression as one line each. No officer buttons and no list of other people's readiness.
- **Officer** (`web-guild-officer`): **Action needed** first, four open `Row`s with a button each (approvals, applications, who has not signed up with `Ping on Discord`, drops to award); the next-raid Plate with the In, Late, Bench, Not signed `Figure`s; **Comp against the target** (a `Tally` per role, targets editable, Onyxia default 40 players); a rail with **Buff and debuff cover** (a `Grid`: need against signed-up providers, red at 0), **Raid-ready** (Attuned 18 of 24; fire resistance on tanks and flask/food marked as not tracked), **Attendance, last 4 raids** (a `Grid` of the slipping members and the bench) and Last raid.
- **Demoted:** past nights collapse to **Last raid** (one `Row` linking the report); progression is one line.
- **Signed-in landing:** a member or officer of one guild lands on its Overview (`/guild` redirects); the directory ("Open a guild", "Top guilds") is for signed-out visitors and for people in no guild or several (then "Your guilds" leads). No dev captions on any state: the mocks carry none.
- **Tabs at 390:** the tab row wraps to two lines instead of scrolling off ("Readiness", "Loot" and "Settings" were cut off); `TabStrip` gains a wrap mode below 768 (`web-guild-member-390.png`).
- **Phone first:** the 390 landings keep the order head, notice, next raid (with the three-way control inside the first 700 px), then the role's sections; the rail sections follow in the same column.

### 7.2 Pages
| Page | The question | Composition (primitives) | Notes |
|---|---|---|---|
| **Raid planner** `web-guild-raid` | "Who is in, in what comp, and what do I swap?" | `ControlBar` (group by Status, Role, Class), Plate of sign-up `Row`s (crest, spec, ilvl, state `Tag`, a `Segmented` to change state for officers), rail: comp `Tally`, buff `Grid`, bench and swaps (`Row`s with Swap in), assignments (`Row`s: tanks, interrupts, heal groups) | The roster reads as the shipped Raid-Helper-style groups; assignments are new |
| **Loot council** `web-guild-loot` | "This dropped. Who gets it, and why?" | Plate of the night's drops (`ItemRow`s, grouped per raid night, Open or Awarded `Tag`), rail: candidates for the selected item (`Row`s: gain in DPS in the BiS page's unit, what it replaces, BiS for them, attendance, loot this tier, MS or OS `Tag`, `Votes`, Award `Button`), the decision record (`Row`s) with Export history | Candidate gain uses `gain_dps` (tier 0) and falls back to `ilvl_delta` (tier 1), as shipped |
| **Raid report** `web-guild-report` | "How did the night go and who needs a word?" | `PageHeader` with the answer `Figure`s (kills, wipes, attempts, deaths), Plate of pulls (`Row`s with result `Tag`s and the death causes), Sections: Parses, Consumables by pull (a `Grid`), rail: Attendance with a manual fix, Loot awarded, the Discord summary | The consumables grid is a marked gap |
| **Recruitment** `web-guild-recruit` | "Should I apply?" | Public guild card Plate (raid nights, progress, what we look for), `FormRow` application, officers listed | Public page |
| **Applications** `web-guild-applications` | "Do we take them?" | Plate of applicant `Row`s (class, spec, ilvl, best parse, message, `Votes`, Accept and Decline `Button`s), rail: the public card's fields and recent decisions | Accept grants the Discord role |
| **Member profile** `web-guild-profile` | "What do I know about this raider?" | `PageHeader` (class-colour Title, `Figure`s), Sections: gear gap (`Row`s, same units as BiS), attendance `Grid`, loot this tier, rail: **Officer notes** (`Row`s with an "Officers only" `Tag`, a `FormRow` to add one) | Notes are private to officers; members never see the section |

### 7.3 Primitives added by this round (all reusable outside guild)
`Segmented` (an inline one-of-few control, the same look as the ControlBar's; used for sign-up state and later for any in-row choice), `Grid` (a coverage matrix of ok, warn, miss and none cells: buff cover, attendance, readiness, consumables by pull, later raid-comp by spec), `Tally` (counts against a target with a `Meter`: comp, a class in the guild, later the guide's raid roles), `Checklist` (private checks as `Row`s with a state mark and a fix link: the member's pre-raid list, later "Before you raid" on the character page), `Votes` (a `Tag` pair and a count), `Notice` (round 1). `Gap` is a mock-only marker for an unconfirmed module and is not a shipped component.

### 7.4 Data honesty: which modules have a confirmed source
| Module | Status | Source it needs |
|---|---|---|
| Next raid, local time, countdown, the viewer's In / Late / Bench | **Exists** | `raids_*` schedule and sign-up |
| Class and spec on sign-ups, comp counts | **Exists** | `raids_roster.go`, `raids_spec.go` |
| Attendance (3 of 4, last 4 raids, slipping, bench) | **Exists** | `raids_attendance*.go` |
| Loot this tier, award, candidates with gain | **Exists** | `loot.go`, `loot_awards` |
| Gear gap, enchants | **Exists** | `readiness.go` (consumer: the member's checklist) |
| Level | **Exists** | addon export level section |
| Attunement | **Derivable, unconfirmed** | the key item (`raids_zones.go` `KeyItem`) read from the addon export's bags section; no readiness field today |
| Flask or food | **Partial** | `consumables.state` is presence of any bagged item; naming a flask needs a consumables-by-item catalogue |
| Fire (and other) resistance | **Gap** | resistance totals in the addon export (a new field in `addon/ForeverSixty/Export.lua` and the FS1 codec) |
| Buff and debuff cover | **Derivable, unconfirmed** | a static class-to-buff catalogue (new curated file) over signed-up class and spec; "Windfury groups" also needs group assignments |
| Comp targets per raid, editable | **Gap** | a target per raid zone with an edit (a column on the raid or a per-zone default) |
| Raid assignments (tanks, interrupts, heal groups) | **Gap** | a `raid_assignments` table |
| "Officers asked you to" | **Gap** | an officer-request table; the mock row is derivable from export staleness |
| Ping non-signers on Discord | **Unconfirmed** | a bot mention or DM action on the raid post (`raids_discord.go`) |
| Loot votes, MS/OS, decision reason | **Gap** | `loot_votes` and a spec flag and reason on `loot_awards` |
| Exportable loot history | **Partly** | `export.go` exports attendance; add loot |
| Deaths with cause | **Derivable, unconfirmed** | the log's damage-taken events per death |
| Parses | **Exists** | `best_parse` and rating |
| Missed consumables per pull | **Gap** | a buff-event parser (the log carries none; `raids_report.go` says so) |
| Manual attendance fix | **Unconfirmed** | an override on the attendance row |
| Applications, votes, Accept grants the Discord role | **Gap** | `guild_applications`, `application_votes`, and a bot permission to grant a role |
| Public recruiting card | **Partly** | `guild_search.go` for the card; recruiting fields are new |
| Private officer notes | **Gap** | a `guild_member_notes` table, officer-gated reads |
No migration is built from this round; the mocks only name the source.

### 7.5 Decisions the owner must make (guild)
D11. **Sign-up words.** The shipped states are In, Late and Bench; the brief says Accept, Tentative and Decline. Keep In / Late / Bench (recommended: it matches the Discord buttons and the data) or rename on both surfaces.
D12. **What a member sees about readiness.** A private checklist with Ready to raid / Almost there / Getting started, and no public worst-first list of others (recommended); officers keep the full table on the Readiness tab.
D13. **New stores.** Approve, in principle, the tables the gaps name: applications and votes, officer notes, loot votes and MS/OS, raid assignments, comp targets.
D14. **Addon change.** Approve adding resistance totals (and named consumables) to the addon export so attunement, resistance and consumes can be read; without it those rows stay marked "not tracked".
D15. **Buff-cover catalogue.** Approve a curated class-to-buff file for the cover grid, and whether "Windfury groups" waits for group assignments.
D16. **Discord actions.** Confirm the bot may mention or DM non-signers and grant the accepted applicant's role.
D17. **Signed-in `/guild`.** One guild: go straight to its Overview (recommended); several: "Your guilds" first, the directory below.

## 8. Recommendation
**Direction A, recommended: the flat system with the paperdoll-led BiS (this spec).** It fixes the clutter (six boxes become one plate), the below-the-fold answer, the control sprawl, the unit mismatch, and it uses the wide screen by composition (two doll columns plus an inspector) instead of by stretching cards. It shares one component across four pages.
**Direction B, a close second only if the redesign is too large to land at once: keep the cards, widen the column.** Page max 1344 to 1792 at 1760, three-column guide, rails on BiS, tiers and sim, as in the `design/mocks/wide/` mocks. Cheaper (CSS and layout only, no component rewrite) and it solves the "slim centre" complaint; it does not solve clutter, hierarchy or the paperdoll. A reasonable sequence is B's `--page-max` token first (it is also A's), then A page by page.

## 9. Order of work
0. The primitives of section 3A and the `surface-guard` ratchet, landed with the BiS lane (the BiS page is their first consumer, the guild landing and Overview their second). 1. BiS plus the shared Paperdoll and `ControlBar` (this spec). 2. Character page and sim card adopt `worn` mode (they already use the component). 3. Reports adopt `report` mode. 4. Guide, then tiers, sim, guild Overview, home, each with its own mock and review.

## 10. Decisions the owner must make
D1. **Direction.** A (flat system, paperdoll-led BiS) or B (widen, keep the cards). Recommendation: A, with B's width token first.
D2. **Page max.** 1792 at 1760 and up (this spec), or 1920 (fuller at 2000, margins of 40 px), or keep 1344. Recommendation: 1792.
D3. **Paperdoll shape.** Two left-aligned columns plus an inspector (this spec) versus the game's mirrored frame with names at the edges. Recommendation: left-aligned columns; the mirror wastes the screen it is meant to use.
D4. **Unit.** Convert every gain to DPS across BiS, character, sim and reports (recommended; touches `paperdoll.ts` and the three mappers), or keep "points" on the old pages and use DPS only on BiS (two units on one site).
D5. **Where to go.** Ship as a second view on the BiS plate (recommended) or leave it as a card at the page's bottom.
D6. **Phone controls.** One summary line that expands in place (recommended), or a bottom sheet.
D7. **The guard.** Land `surface-guard.test.ts` as a ratchet with the allowlist (recommended: it cannot block a lane on day one and it only gets stricter), or as a hard fail now (forces the 213-file migration into one lane).
D8. **Title colour for pages with no class.** Guild: faction colour (mocked). Tiers, sim, home: `--color-strong` (white) or gold? Recommendation: white, so colour stays a class, spec or faction signal and nothing else.
D10. **Row density.** With the persona fixes a BiS row is 3 to 4 lines and the full doll is about 740 px tall (first 12 slots above the fold on a 1080 monitor). Keep every fix inline (recommended: the personas asked for them), or move next best and the worn item behind the inspector and keep the earlier one-screen set.
D9. **Retire `Card` and `CardHead` outright** (recommended; the primitives replace them page by page and the last page deletes them), or keep `Card` alive as a deprecated alias of `Plate` with a head, which preserves the icon-chip look the owner is moving away from.
