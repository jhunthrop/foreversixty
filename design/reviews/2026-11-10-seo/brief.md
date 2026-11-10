# SEO lane review brief (2026-11-10)

Owner-approved lane: give Google real text on the tool pages without JavaScript, plus three cleanups. Background: report.md in this folder.

## What changed (visible)

1. **/sim**: the landing head gains a lead sentence (`simPageCopy.lead`, web/src/lib/sim/page-copy.ts). The landing was already server-rendered.
2. **/sim/weights**:
   - The head gains a lead (`bulkCopy.weightsLead`).
   - With no character loaded, the main column now shows **Published weights** for the default spec (Fury Warrior, level 60 headline preset) under the caution line. Before, it was empty.
   - The same table is server-rendered (web/src/components/sim/tools/WeightsShell.svelte) and handed to the island through `data-default-weights`, so the island shows the same figures.
   - The figures come from the nightly BiS file; no number is typed into prose.
3. **/planner**:
   - The head gains a lead (`plannerCopy.lead`).
   - The loading Plate's tree heads now show the opening class's tree names at 0 points (the editor's own head), where there were grey bars before.
   - The talent grid itself still arrives with the island. The talent names in it are aria-labels and tooltips, not visible text, so server-rendering it adds nothing a crawler reads and risks the planner's blocking-time budget.
4. **/rankings**:
   - The head gains a lead (`rankingsCopy.lead`).
   - The invisible "Updated 13 Dec" placeholder is gone; a no-break space holds the meta line's height until the board names a real date.
   - The board's rows are not server-rendered: they come only from the live API, and there is no build-time file to render them from.
5. **Home**: the signed-in skeleton's pending title is no longer a second, hidden h1 (`PageHeader` `secondHead`). This change is not visible.
6. **BiS spec page titles**: "<Spec> BiS: Pre-Raid and Leveling, WoW Forever" (the level range is dropped), so every title is 60 characters or fewer. This change is not visible on the page.

## Rules the copy follows

- Game words only.
- No em dashes.
- No meta talk about the site, sim, pipeline or "band".
- Never explain what veterans know.
- Prose never quotes a nightly number (memory: feedback-prose-never-quotes-nightly-numbers). "1.00" in the weights lead is the scale convention, not a nightly figure.

## Screenshots

Screenshots are committed as JPEG (quality 55) for shots-before, shots-r3 and shots-r4. The round 1 and round 2 PNGs were reviewed from the working tree and are not committed; their CLS files (cls.jsonl) are.


- Full page, at 1440, 390 and 2000 px, with JavaScript on (`-js`) and off (`-nojs`).
- **After:** design/reviews/2026-11-10-seo/shots-r1/
- **Before:** design/reviews/2026-11-10-seo/shots-before/
- Measured CLS for the JavaScript runs is in shots-r1/cls.jsonl.
- Rankings with JavaScript calls the production API from localhost, so its board may stay on the skeleton. That state is not part of this lane.

## Scope

- Judge only what this lane changed: the leads, the weights default table and shell, the planner's loading tree heads, the rankings stamp and the home heading.
- Report pre-existing issues you notice separately, as "outside the lane".
