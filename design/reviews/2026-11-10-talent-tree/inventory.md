# Talent tree inventory (2026-11-10, lane/talent-tree)

Owner, 2026-11-10: "why don't the talent trees look the same everywhere? ... Some of them work on hover, some don't."
The screenshot (a read-only "Talents" panel, "Open in planner ›", gold-edged learned icons, green "3/3" badges, no arrows) is the sim's Talents section (`SimTalentsSection`).

Search: every `web/src` component, page and lib that names talents, plus `api/internal/site` (chrome.html, page.go).

## Seven renderings, four implementations

| # | Surface | Component (before) | Icon size | Border and states | Rank badge | Arrows | Tree head | Art | Tooltip | Click | Empty, loading |
|---|---|---|---|---|---|---|---|---|---|---|---|
| 1 | Planner `/planner` (edit) | `ui/TalentEditor` + `ui/TalentButton` + `ui/TalentTip` | 44 / 48 from 768 / 56 from 1760 | 2 px inset ring: avail `#4fae5c`, part green, max gold, lock line + greyscale 55%; gold "differs" dot; dashed pill when planned past level | pill, 12/14 700 on `--bg`, ring-coloured border, `-5 px` | yes, gold when met, line when not, `--bg` halo | spec icon 28, name 15/20, points 18/24 gold, per-tree Reset; tabs below 768 | client art at 50% | own `TalentTip` (300 px, no icon): name, `Rank n of m`, this rank, Next rank, requirement in `--wipe`, band line in gold; opens 150 ms after hover or focus, beside the cell (client style), Escape closes; **no tooltip on touch** (a tap adds) | add; right-click or 500 ms press removes; arrows, Backspace | Skeleton Plate with the opening tree art; LoadError |
| 2 | Shared build `/b/:id` (API-served, mounts the planner island) | same as 1 (the store refuses presses) | same | same | same | same | same | same | same | refused in the status line | same |
| 3 | Sim Talents section: `/sim` before a run (`SimPreRun`) and on a result (`RunReport`) | `sim/SimTalentsSection` → `planner/TreeGrid` + `planner/TalentCell` (`readOnly`, `plain`) + page CSS overrides | 36 px fixed | **outside** 2 px ring: max gold, part green; untaken greyscale 45%; cell border hidden; the old planner CELL_BORDER / CELL_PILL tokens underneath | the old planner pill (11/14, green border) hidden on untaken | **none** (`plain`) | name 15/20 + points gold, no icon | **none** | `data-game-tip` through the page host. On a result the RunReport host resolves talents from a module registry (`setResultTalentTips`); **before a run no host resolves talents: no tooltip at all** | none | Skeleton lines; LoadError; "does not fit" line. **Below 768 the grid is hidden** (name and points only) |
| 4 | `/sim/talents` loadouts (`TalentCandidates`) | `ui/TalentTrees` (old static primitive) + page CSS | container-fit 18 to 44 px; one tree per row below 768 up to 56 | outside 2 px ring, max gold, part green, untaken greyscale 45% | bare `3/3` text with a black text-shadow, no pill | none | name 13/18 + points gold | none | **none: `loadoutTrees` sets no tip key, so nothing opens on hover, focus or tap** | none | n/a (rendered from the loadout) |
| 5 | Guide At a glance (`GuideGlance.astro`, static) | `ui/TalentTrees` | container-fit 18 to 32 px | as 4 | as 4 | none | as 4 | none | `data-game-tip` through `GuideTooltips` (BisTooltipHost: hover, keyboard focus, tap toggles, Escape), below or above the cell | none | n/a (server rendered) |
| 6 | Guide Leveling lane (`GuideLevelingTree`, `client:visible`) | `ui/TalentTrees` | as 5 | as 5 | as 5 | none | as 5 | none | as 5 (the page host) | none | n/a |
| 7 | Planner Point order (`ui/PointOrder`) | a level list of talent icons, not a tree | 36 | n/a | rank digit | n/a | n/a | n/a | the planner `TalentTip` again (name, rank, this rank only) | none | dashed sockets |

Checked and holding **no** talent grid: the character page (`Character.svelte` and `character/*`; talents appear only in the "export not read" line), the report player and coach views (`report/*` only reads tree sizes for the planner link; `coach/CoachText` has no talent UI), the guild member profile and readiness board (points only), tier list and rankings rows (spec icons), the guide's Key talents and rotation rows (`GuideBuildsLane.astro`: icon rows, not trees), and the API site (`api/internal/site/page.go`, `chrome.html`: `/b` renders the planner island; the unavailable page links to the guides). No server-generated tree markup exists, so no generator changes.

## Differences found

1. **Three cell languages.** Inset ring with a pill (planner), outside ring with the old planner's green-bordered pill (sim), outside ring with a bare shadowed number (guides and /sim/talents). Available-to-spend exists only in the planner.
2. **Four icon sizes, no shared scale.** 44/48/56, 36, 18 to 32 fit, 18 to 44 fit (56 on a phone).
3. **Arrows and art only in the planner.** Every read-only tree drops the prerequisite lines and the client's tree art.
4. **Two tooltip panels.** The planner's `TalentTip` (no icon, "Rank n of m", 300 px, beside the cell) against the site's `GameTooltip` (icon, "Rank n/m", gold description, below or above).
5. **Hover works on three of six read-only surfaces.** `/sim` before a run and `/sim/talents` open nothing; the sim result works only because a module registry is filled after load.
6. **Touch.** The planner never shows a tooltip on touch; the guides toggle on tap; the sim before a run and /sim/talents show nothing.
7. **Keyboard.** The planner and the sim grid have one tab stop per tree with arrows; the guide trees make every talent a tab stop (up to 51 per build); /sim/talents has none.
8. **Phone.** The planner shows tabs and one tree; the sim hides the grid; /sim/talents stacks the trees; the guides keep three tiny trees side by side.
9. **A second, dead-end implementation.** `planner/TreeGrid` + `planner/TalentCell` + the `CELL_*` / `CONNECTOR_STROKE` tokens survive only for the sim section; `guides/tree-store` only feeds it.
