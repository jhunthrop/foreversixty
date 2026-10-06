# Guides audit — 2026-10-05

Branch `guides-audit-27`. Build: real data (`FOREVER_DATA` unset / `real`), build `1.60.1.70009`,
served with `npx astro preview --port 4329`. Captures in `design/mocks/renders/audit/`
(git-ignored) at 1440 and 390 for all 37 guide pages (`/guides`, 9 class landings, 27 spec
guides), plus 2000 for `/guides` and `/guides/warrior/fury`.

## 1. Rail order fix

`railStatRows` (`web/src/lib/guides/rail-stats.ts`) now orders the Stat priority rail card by
the sim, not by the guide's own `statPriority` frontmatter order: per-point significant stats
sorted by scale factor descending, then haste last ("N per 1%"), then insignificant stats last
("Not significant"), with the guide's written order kept only as a stable tiebreak. Label
vocabulary (`Critical strike` → `Crit` alias) unchanged. Specs with no band-60 file keep
today's behaviour (card omitted, 1-up rail). Unit test updated
(`web/src/lib/guides/rail-stats.test.ts`) with a new case asserting the Fury warrior order and
a tiebreak case.

**Guides whose written order disagrees with the sim** (19 of 20 specs with band-60 data; only
Destruction Warlock agrees):

| Spec | Written | Sim (card now shows) |
|---|---|---|
| Balance Druid | SP, Int, Crit, Hit, SH, SPen | Int, SP, Hit, Crit, SH, SPen |
| Feral Druid | AP, Feral AP, Str, Agi, Crit, Hit | Str, Agi, AP, Crit, Hit, Feral AP |
| Beast Mastery Hunter | AP, RAP, Agi, Crit, Hit | Agi, Crit, RAP, Hit, AP |
| Marksmanship Hunter | AP, RAP, Agi, Crit, Hit | Agi, Crit, RAP, Hit, AP |
| Survival Hunter | AP, Agi, Str, Crit, Hit | Agi, AP, Str, Crit, Hit |
| Arcane Mage | SP, Int, Crit, Hit, SH, SPen | SP, Hit, Crit, Int, SH, SPen |
| Fire Mage | SP, Int, Crit, Hit, SH, SPen | Hit, SP, Int, Crit, SH, SPen |
| Frost Mage | SP, Int, Crit, Hit, SH, SPen | Hit, Int, SP, Crit, SH, SPen |
| Retribution Paladin | AP, Str, Agi, Crit, Hit | Str, AP, Crit, Hit, Agi |
| Shadow Priest | SP, Int, Crit, Hit, SH, SPen | Int, Hit, SP, Shadow power, Crit, SH |
| Assassination Rogue | AP, Agi, Crit, Hit | Agi, AP, Crit, Hit |
| Combat Rogue | AP, Agi, Crit, Hit | Agi, AP, Crit, Hit |
| Subtlety Rogue | AP, Agi, Crit, Hit | Agi, AP, Crit, Hit |
| Elemental Shaman | SP, Int, Crit, Hit, SH, SPen | Int, SP, Hit, Nature power, Crit, SH |
| Enhancement Shaman | AP, Str, Agi, Crit, Hit | Str, AP, Hit, Crit, Agi |
| Affliction Warlock | SP, Int, Crit, Hit, SH, SPen | SP, Int, Hit, Crit, SH, SPen |
| Demonology Warlock | SP, Int, Crit, Hit, SH, SPen | SP, Hit, Crit, Int, SH, SPen |
| Arms Warrior | AP, Str, Agi, Crit, Hit | Crit, AP, Hit, Haste, Str, Agi |
| Fury Warrior | AP, Str, Agi, Crit, Hit | AP, Crit, Hit, Haste, Str, Agi |

Agrees: Destruction Warlock. No band-60 file (card omitted, unaffected by this change):
Restoration Druid, Holy Paladin, Protection Paladin, Discipline Priest, Holy Priest,
Restoration Shaman, Protection Warrior.

## 2. Fixes made in `web/src`

1. **`GuideActionRail.astro`, `.guide-rail-tree-grid`** — the talent rank badge
   (`.guide-rail-tree-rank`, `position: absolute; right: -4px; bottom: -4px`) pushed the whole
   build-card tree grid 4px past its container whenever a maxed talent landed in the last
   column or row. Confirmed by the DOM check's `scrollWidth > clientWidth` probe on every spec
   guide at both 1440 and 390 (ancestor chain: `.guide-rail-tree-grid` → `.guide-rail-tree` →
   `.guide-rail-trees`, occasionally bubbling further). Fixed by giving the grid
   `box-sizing: border-box; padding: 0 4px 4px 0;`, reserving exactly the badge's own offset
   inside the grid's box. Re-audit: no more grid/tree/trees-level overflow on any page; the
   leaf `.guide-rail-tree-cell` itself still reads `scrollWidth` slightly over its own
   `clientWidth` where a badge sits, which is the intended decorative overlap and no longer
   propagates to any ancestor or the page.
2. **`SourcesList.astro`** — the Sources footer's `<li>` (`items-center`, no `min-w-0`) let a
   long, space-free source label segment (e.g.
   `data/builds/1.60.1.70009/talents/hunter.json`) refuse to wrap, overflowing the whole page
   horizontally at 390px. Confirmed as real `document.documentElement.scrollWidth >
   clientWidth` (404 vs 390, 14px) on Survival Hunter; the same label pattern recurs across
   most guides' sources, so this was a live risk everywhere, not just that one page. Fixed
   with `items-start` (so a wrapped label doesn't look mis-aligned against the pill) plus
   `min-w-0 break-words` on the anchor. Re-audit: `bodyOverflow` is false on every page at
   both widths.

Both fixes verified by rebuilding (`npm run build`, real data), re-serving, and re-running the
DOM check + full screenshot pass; no regressions found on the 19 other specs or the class
landings.

## 3. DOM check results (post-fix)

74 page × width combinations checked (37 pages × {1440, 390}), every spec guide's rail,
rotation card, stat priority card, and leveling strip also probed directly (37 pages ×
{build-tree-count, has-build-code, rail-card-count, leveling-strip-present}).

| Check | Result |
|---|---|
| Broken images (`naturalWidth === 0`) | 2 per page, every page site-wide (not guides-specific — see "needs data / out of scope" below) |
| Horizontal overflow (`scrollWidth > clientWidth`) at page/ancestor level | 0 after fix (was present on every spec guide's build card and on Survival Hunter's full page) |
| Empty rail cards | 0 |
| Rotation lines without icon or placeholder | 0 |
| Stat values that are NaN/undefined | 0 (undefined values correctly render "Not significant", never a raw `NaN`) |
| Build card shows all three trees when the guide has a build code | 27/27 |
| Leveling strip rows present | 20/27 (present wherever a band-60 file exists; correctly omitted on the 7 specs with no band-60 file, matching the rail's own Stat priority/Rotation card omission rule) |
| Console errors (excluding `api.foreversixty.gg/v1/me` CORS noise) | 0 |

**CORS noise**: every page's preview-mode console shows repeated `Access to fetch at
'https://api.foreversixty.gg/v1/me' ... blocked by CORS policy` — the account menu's
session probe hitting the real API from `localhost:4329`, which isn't an allowed origin. Not a
guides defect; happens on every page of the site in preview mode and is unrelated to this
lane's scope.

**Broken images**: `AccountMenu.svelte`'s `data-testid="account-menu-preflight-crest"` /
`-preflight-faction` `<img>`s render with no `src` attribute (client-side JS fills it in after
a successful `/v1/me`, which CORS blocks in preview) — not a real image-load failure, not
guides-specific (same header on every page site-wide), and out of this lane's scope
(`AccountMenu.svelte`, not a guides component).

## 4. Manual screenshot review

Read at least six captures directly: `guides-index-1440`, `guides-warrior-fury-1440`,
`guides-warrior-fury-390`, `guides-paladin-holy-1440` (1-up rail, no band-60 file),
`guides-rogue-subtlety-390`, `guides-druid-1440` (class landing). All clean: rail cards render
at their intended column count (3-up for specs with band-60 data, 1-up build-only otherwise),
stat priority card order matches the sim table above, build card shows all three trees with
readable icons at both widths, leveling strip renders with icons and "Load in planner" links
where present, no visible clipping or broken art.

## 5. Needs data (not fixable in `web/src`)

None found. All 27 spec guides have a build code and render all three trees; the 7 specs
without a band-60 file degrade exactly per the existing design rule (card/strip omitted, never
a stand-in for missing data) and were not expected to show the rail's Rotation/Stat priority
cards or the Leveling strip.

## Findings summary

- 2 real defects found and fixed in `web/src` (tree-grid badge overflow; Sources footer
  long-label overflow).
- 0 defects needing data changes.
- 19 spec guides' written `statPriority` order disagrees with the sim (table in §1); the rail
  card now follows the sim for all of them per the task's ruling.
