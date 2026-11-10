# Density review: talent planner (round 3)

Job: pick and order talent points for a class. Captures: round3-captures planner, planner-you, planner-guide at 1440 and 390 (planner-guide-390 viewed).
Out of scope: handoffs.md items, respec cost, pick rates.

## Verdict: THIN

## Fixed since round 2
- Guide builds rows now carry "n points off" against the loaded build.

## Still MUST
1. Dead space at 1440 (planner-you, planner-guide): rail ends ~y1000, main column runs to ~1800; ~800 px of empty ground right of Point order / Share / Import. Turn Guide builds into a "Build vs guide" table: first differing talent, per-spec split delta, Load button (source: guideBuilds in planner.astro, decoded codes lib/planner/guide-builds.ts). Files: components/planner/GuideBuilds.svelte, Planner.svelte.
2. Level band card still says only "Matches" or "21 of 25 points". List differing talents (name, you n, guide m), click flashes the cell (BandCompare diff on onband). File: components/planner/BandCompare.svelte.
3. Default state still blank (planner-1440): all 0/n, no Point order plate, head "0 of 51 points", rail of 3 small cards. Show the spec guide builds' first-10 point order icons (guideBuilds, /data/leveling-plans.json). Files: PlannerOrder.svelte, Planner.svelte.
4. In the game, Share, Import still three stacked plates. Merge into one "Take it with you" plate (Copy, Send to Felsnap, Share + Show DPS, Import row). Files: InGame.svelte, SharePanel.svelte, ImportBox.svelte.
5. Phone (planner-guide-390): order is still trees, Level build, In the game, Guide builds, then Point order (~y1400 of 3038). Put Point order right after trees, guide builds as a drawer, the merged plate last.

## Should
- Head: spec split 5/13/33 and loaded guide build name (SummaryBar.svelte).
- Point order tooltip: talent name, rank text, exact level (ranks[].description, leveling-plans.json).
- Character mode: crest plus Reset in place of text line (Planner.svelte).
