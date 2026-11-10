# Density review: talent planner (round 4)

Job: pick and order talent points for a class. Captures: round4-captures planner, planner-you, planner-guide at 1440 and 390. Out of scope: handoffs.md items, respec cost, pick rates.

## Verdict: DENSE WITH ADDITIONS

## Fixed since round 3
- Default state (planner-1440) now opens on the Arms guide Point order (icons, levels 10 to 60) and a Guide builds list; no blank page.
- Guide builds rows carry points off, per-spec split delta (+26/+7/-33) and the first missing talent ("Missing Endurance Training"); the loaded build is tagged "Loaded".
- Level band card names the talents past the band ("Past it: Endurance Training 4/5, ...").
- In the game is one plate: Send to Felsnap, Copy, Share + Show DPS, Import.
- Phone (planner-guide-390): trees, Point order, Level 60 build, Guide builds, In the game; ~2990 px.
- Head carries "Sim this build"; tab heads carry the 5/13/33 split.

## Additions (ranked)
1. 1440 dead ground under Point order (planner-guide ~y1280 to 1580, planner-you ~y1310 to 1560) while the rail runs longer. Move Guide builds under Point order as a row, or let the rail sticky-pin; Planner.svelte layout.
2. Guide builds rows show no Load control or hover; add row click/Load and flash the differing talents (guideBuilds in planner.astro, lib/planner/guide-builds.ts).
3. Band card lists only what is past the band; add "you n, guide m" per differing talent and flash the cell on click (BandCompare.svelte).
4. Point order cells: tooltip with talent name, rank text and exact level (ranks[].description, /data/leveling-plans.json).
5. Character mode: crest plus Reset in the head in place of the text line.
