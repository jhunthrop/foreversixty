# Grouped header: four doors, each a hover panel

Author: ux-designer brief by the owner, 2026-10-15. Canvas (the approved mock, every state):
https://claude.ai/artifact/QMUea9sAGREktD7bPZhFmB. Supersedes section 2.2 (the eight doors) of
`design/specs/2026-10-09-nav-character-selector.md`; every other section of that spec stands
(the selector, the 56px phone bar, widths, the 1344px container, tenet 13's no-wrap rule).

## 1. Why

The addon v2 lane put the compact character lookup into the bar (`HeaderSearch.astro`) and the
eight-door row ran out of room at 1280 (fixed in 2ec19219 by shrinking gaps). Eight doors was
also the wrong shape: "Get set up" is onboarding, "Tier List" and "Rankings" are two views of
the same logs, and the guild centre (addon v2) had no door at all. The owner's call: fewer
doors, each opening an expanded list on hover, the pattern Wowhead and Warcraft Logs use.

## 2. The doors, in this order

| Door | href (click) | Panel rows, in order (title: one-line description) |
|---|---|---|
| Tools | `/planner` | Planner: "Talents, the build you follow, and what to train next" · BiS (`/bis`): "The gear to chase in your band, slot by slot" · Simulator (`/sim`): "What your gear does, simulated fight by fight" |
| Logs | `/logs` | Upload a log (`/logs`): "Drop a combat log, or let the companion stream it live" · Your reports (`/reports`): "Every fight you logged, newest first" · Rankings (`/rankings`): "Parses by boss, spec and band, from every log" · Tier list (`/tiers`): "DPS, healer and tank, ranked from the same logs" |
| Guild | `/guild` | Your guild (`/guild`): "Progress, the week's raids, who logged what" · Roster (`/guild#roster`): "Everyone the addon has seen, with gear score and last sync" · Loot council (`/guild#loot`): "Who gains most from each drop, by the sim, before the roll" · Guild logs (`/guild#raids`): "Reports from guild members, raid night by raid night" |
| Guides | `/guides` | The nine classes, `classes.json` order, in a three-column grid: 28px crest (`classCrestSrc`) ringed in the class colour, the class name in the class colour, under it the spec names in talent-tree order joined by " · ", each class linking to `/guides/{slug}`; the panel foot: "All guides" → `/guides` |

**Get set up stays on the bar** (owner, 2026-10-15): a fifth, plain door after Guides, `/setup`,
with no panel, no chevron and the same type as the others; on the phone it is its own row after
the four groups, in the Discord row's shape. It is not a Tools row.

A door is a real link: clicking it goes to its href. The panel is extra, never the only way in.
The door's label is uppercase Barlow 700 13px, tracking .12em, `--color-nav`, with a 12px
chevron after it (muted; gold and rotated 180° while its panel is open). A door is gold with a
2px gold rule on the bar's bottom edge while open, and marked `aria-current="page"` (gold, no
rule) when the page is one of its rows (`isNavGroupCurrent`: the group's own href or any row's
href by `isNavItemCurrent`).

When a Guild row's page is not resolvable (no session, no guild), the panel still shows the
four rows and `/guild` is the guild landing: a guild lookup (name, region, ruleset), the top
guilds, and a line for officers on claiming; a signed-in member of a verified guild is sent on
to their guild's page from it. The Guild door itself carries `data-guild-tab="overview"`, so
the session rewrite below points it at the member's guild too. When the session
resolves with a verified guild (`Me.guilds[0]`), the header fills the first row's title with the
guild's name and rewrites the four hrefs to `guildHref(region, ruleset, name)` plus the tab
hash. No other row carries live numbers in this round (section 7).

## 3. The panel

A `div[role="menu"]` inside the door's anchor box, absolutely positioned: top at the bar's
bottom edge (the bar is 81px; the panel's top edge sits on the 1px border line), left aligned
12px left of the door's text. 560px wide (Tools, Logs, Guild), 760px (Guides). `--color-raised`
background, 1px `--color-line`, radius 6, padding 10px, shadow `0 14px 36px rgb(0 0 0 / .6)`,
z-index above the page's own sticky elements. Rows are 44px minimum: a 20px gold stroke icon,
the title (Barlow 600 14px `--color-strong`), the description (12px `--color-muted`) under it;
hover and focus-visible background `#101624`. The Guild panel's foot: "Guild settings" →
`/guild#settings`; the Guides foot as in section 2. No other panel has a foot this round.

Below 1024 the panels do not exist: the phone Menu lists the same four groups (section 5).

## 4. Behaviour (tenet 9: every state, keyboard and touch parity)

- **Hover** opens a panel 120ms after the pointer enters its door and closes it 180ms after
  the pointer leaves both the door and its panel (a pointer crossing the bar never flickers
  panels; moving from the door down into the panel keeps it open). Only one panel is open.
- **Click or tap on the chevron** toggles the panel without navigating; click on the label
  navigates. On a touch device (`(hover: none)`) a first tap on the door opens the panel and
  a second tap navigates, so no door is hover-only.
- **Keyboard:** the door is a link in the tab order; `ArrowDown` or `Enter` on a focused door
  with the panel closed opens it and focuses its first row; `ArrowDown`/`ArrowUp` walk the
  rows; `Escape` closes and returns focus to the door; `Tab` out of the panel closes it.
  `aria-haspopup="true"`, `aria-expanded`, `aria-controls` on the door; rows are
  `role="menuitem"`.
- **Pointer leaving the header** closes any panel. A page click outside closes it.
- **Reduced motion:** no transitions.
- No layout shift: the bar's height and the doors' positions never change when a panel opens
  (the panel is absolutely positioned; the gold rule is a pseudo-element).

## 5. Widths

Four doors give the bar back the room the lookup took. The budget at 1280 (measured on the
rendered bar, signed in, "Obnoxious Yell"): wordmark 229, doors ≈ 330, lookup 65 (the "/" hint
returns from 1280, as HeaderSearch.astro first specified), selector 232, Discord 44, gaps and
padding as `nav.css` sets them; slack ≥ 60px. The 1024 to 1279 tier keeps its current values.
At 1440 and up Discord has its label and the bar sits in the 1344px container. The nav-selector
spec's measured rule stands: no sideways scroll at 1024, 1279, 1280, 1439, 1440 and 2000.

Phone Menu (below 1024): after the character row and the Search row, four groups in order,
each a muted uppercase 11px heading ("Tools", "Logs", "Guild", "Guides") over a two-column
grid of 44px rows (title 14px 600, the description 11px muted under it); Guides is a
three-column grid of crest-over-name tiles; Discord last. The Menu's first row and the Search
row are unchanged.

## 6. Copy (verbatim)

Doors: `Tools`, `Logs`, `Guild`, `Guides`. Row titles and descriptions: section 2's table.
Foots: `All guides`, `Guild settings`. The Guild panel's first row title before the session
resolves: `Your guild`. Screen-reader label of a chevron toggle: `Open the {Door} list`.

## 7. Not in this round

The live chips the mock shows beside rows ("11 upgrades · +3.4 DPS", "Arms · top 4%",
"38 · 11 raid-ready", "addon · 17 days ago", "no sim yet"): they need the home panel's upgrade
count, the character's rank and the roster summary available to the header on every page.
Shipped as a follow-up once the panel is in; the row markup leaves a trailing slot for them.

## 8. Tests

Unit: `lib/nav.test.ts` (groups, order, hrefs, `isNavGroupCurrent`), `Header.test.ts` (four
doors, their panels server-rendered and hidden, Guides grid from `classes.json`, the phone
groups). E2E: `nav.spec.ts` and `nav-selector.spec.ts` follow the four doors; `tiers.spec.ts`
asserts Tier list inside the Logs panel; a new `nav-panels.spec.ts` covers hover open and
close with the delays, click-toggle, keyboard (ArrowDown, Escape, Tab), touch (`hasTouch`
project or `(hover: none)` emulation), the phone groups, 44px targets, `aria-current` on the
group, and the no-sideways-scroll rule at the six widths. Lighthouse budgets unchanged.
