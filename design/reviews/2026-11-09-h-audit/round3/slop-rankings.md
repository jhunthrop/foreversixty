# AI-slop review, rankings / tiers (round 3, round3-captures/, 1440 + 390)
/ratings skipped (handed off). Handed-off items in handoffs.md not failed.

## Fixed since round 2
- tiers-dps: "Your spec" plate is now a quiet "Sign in to mark your spec" link (19).
- tiers-dps-you: "Tier A . #2 . DPS" line gone (22).
- tiers-dps 390: "vs #1" now under the DPS value.
- rankings-board: row subtitle no longer wraps to 3 lines.

## Verdicts
- rankings: FAIL (tell 8)
- rankings-you: FAIL (tell 8)
- rankings-board: FAIL (tell 8, 13)
- tiers-dps: PASS
- tiers-dps-you: FAIL (tell 8)

## Remaining findings
- 8 | Ornate-corner frame on the Encounters / primary table, plain plate on Leaders, sidebar "Your rank" plate with an inner second border (card in card); three plate treatments | rankings, rankings-you, rankings-board 1440+390 | web/src/components/rankings/*.svelte, rank-list.css | one frame; Leaders and sidebar a single hairline, no inner border.
- 8 | "Fury Warrior" plate with inner border, two buttons and ~40px empty plate under them; at 390 it pushes the table down 160px | tiers-dps-you 1440+390 | tiers/TierPage.astro | drop the plate: Best in slot / Guide as links in the tab row or on the highlighted row.
- 13 | Three rows then ~300px bare parchment | rankings-board 1440 | rankings/RankingsEmpty.svelte, Rankings.svelte | show the encounter's per-spec leaders under the 3 rows (390 is f-selected's).
