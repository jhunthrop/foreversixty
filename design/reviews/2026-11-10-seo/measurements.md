# Measurements (real-data build, 2026-11-10)

## Server HTML words inside `<main>`

Scripts, styles, noscript and svg are not counted. The before figures are from the dcead254c build, the after figures from this lane's final build.

| Page | Before | After | Added |
|---|---|---|---|
| /sim | 95 | 120 | +25 |
| /sim/weights | 28 | 353 | +325 |
| /planner | 24 | 59 | +35 |
| /rankings | 17 | 46 | +29 |
| / (home) | 98, two h1s | 98, one h1 | 0 |

## BiS titles

Every BiS title is now 60 characters or fewer. The longest is "Beast Mastery Hunter BiS: Pre-Raid and Leveling, WoW Forever" at 60.

## Island budgets (scripts/check-island-size.mjs)

No budget changes.

| Island | Before (B) | After (B) | Limit (B) |
|---|---|---|---|
| sim-tools closure | 90,831 | 92,546 | 94,001 |
| planner closure | 104,795 | 105,096 | 108,347 |
| Rankings closure | 113,883 | 114,121 | within limit |

## CLS

Measured in Playwright with JavaScript on, at 1440, 390 and 2000 px (shots-*/cls.jsonl).

- **/sim, /sim/weights and home:** 0 before and after at every width.
- **/planner:** 0.0265 before and 0.0044 after at 1440; 0.0029 before and 0.0026 after at 2000.
- **/rankings at phone width (fixed in round 4):**
  - The source was the pre-existing "How to get your logs here" panel. With nothing ranked it sat in the rail, which stacks first below 1280 (railFirst), so it was inserted above the board after the board answered.
  - Below 1280 it now renders after the board; from 1280 it stays in the rail beside the board, as the mock has it.
  - lhci (rankings.html, mobile 360x640, API unreachable, 3 runs, real-data build): **0.146** at dcead254c, **0.093** after round 3, **0.0008** after round 4 (all three runs). Performance 0.99, LCP 1.8 s.
  - Playwright, 360, API unreachable: 0.0008. Playwright, 390, production API from localhost: 0.047 before the lane, 0.082 after round 3, 0.0073 after round 4.
  - ux-designer SHIP on /rankings at 390 (r4/verdict.md).
- **lhci on index, planner, rankings, sim and sim/weights (real data):** rankings failed on CLS until round 4; it now passes. (rankings/warden-kelthas.html exists only in the fixture build, so the real-data run cannot load it.) Planner has a best-practices warning of 0.93, which is pre-existing.
