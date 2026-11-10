# Density review: talent planner (round 2)

Job: pick and order talent points for a class. Captures: round2-captures planner, planner-you, planner-guide at 1440 and 390 (planner-you-390 viewed; the others by layout).
Out of scope (handoffs.md, f-content / f-selected): tier-gate gutter and "n more to tier k", client build line, yellow dash marker. Needs data: respec cost, pick rates.

## Verdict: THIN

## Fixed since round 1
- Point order is now icon cells with rank stamps, a 0-9 column header and a level-band label per row (planner-you, planner-guide). Was blank squares.
- Guide builds rail rows carry a spec crest, kind and split; the loaded build is marked "Loaded" (planner-guide).
- Head shows "n of m points" and the loaded character.

## Still MUST
1. Dead space remains at 1440 (planner-you, planner-guide). The rail ends near y=1000 while the main column runs to y=1800; about 800 px of empty ground right of Point order / Share / Import. Fill with a "Build vs guide" table: one row per guide build of the class (Beast Mastery 31/20/0 etc.), columns points differing from the loaded build, first differing talent, Load button; replaces the link-only Guide builds list. Source: guideBuilds in planner.astro, decoded codes in lib/planner/guide-builds.ts. Files: components/planner/GuideBuilds.svelte, Planner.svelte. Phone: same rows with the diff count and Load.
2. Level band card says only "Matches" (planner-guide) or "21 of 25 points" (planner-you). List the differing talents inline (name, you n, guide m), click flashes the cell. Source: BandCompare diff computed on band change (Planner.svelte onband). File: components/planner/BandCompare.svelte.
3. Default state is blank (planner-1440): trees all 0/n, Point order says "No points spent yet", head reads Level 9 with 0 of 51. Show the class's guide builds in Point order as the first 10 icons per spec from the leveling plan, or pre-load the spec build faded. Source: guideBuilds, /data/leveling-plans.json. Files: PlannerOrder.svelte, Planner.svelte.
4. Rail and lower cards are still thin: In the game (1-2 buttons), Share (title, two buttons, toggle), Import (textarea) are three stacked plates. Merge into one "Take it with you" plate with Copy, Send to Felsnap, Share and an Import row; frees the rail. Files: InGame.svelte, SharePanel.svelte, ImportBox.svelte. At 390 this plate follows Point order, trees first, guide table as a drawer (today Level build, In the game, Guide builds all sit between trees and Point order, pushing the order 700 px down).

## Should
- Head: spec split 5/13/33 beside "51 of 51 points" and guide build name (SummaryBar.svelte).
- Point order: tooltip with talent name, rank text and exact level on hover (talents ranks[].description, leveling-plans.json order).
- Character mode: crest plus Reset in place of text line (Planner.svelte).

## Needs data
- Respec cost by level; talent pick rates; per-talent sim value (sim owns, link only).
