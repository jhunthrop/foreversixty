# AI-slop review, rankings / tiers (round 2, captures in round2-captures/, 1440 + 390)

/ratings skipped (h-report-tail). Items in handoffs.md noted, not failed.

First impression: reads as a player's tool. Class-coloured names, real icons, dense rows, S/A/B/C/D serif letters. No gradients, glow, hero or eyebrows.

## Fixed since round 1
- "Percentile" repeated per row at 390: gone, bare band-coloured number.
- "On Onyxia, Normal ruleset" under "14th of 212 Fury Warriors": gone.

## Verdicts
- rankings: PASS WITH FIXES
- rankings-you: PASS WITH FIXES
- rankings-board: PASS WITH FIXES
- tiers-dps: PASS WITH FIXES
- tiers-dps-you: PASS WITH FIXES

## Remaining findings
- 8 | Ornate-corner frame on the Encounters / primary table, plain double-border plate on Leaders, sidebar plates ("Your rank", "Your spec", "Fury Warrior") with an inner second border; same plate stamped 2-3x per page | rankings, rankings-you, rankings-board, tiers-dps(-you), 1440+390 | web/src/components/rankings/*.svelte, rank-list.css, tiers/TierPage.astro | one frame (ornate) on the primary table only; Leaders and sidebar get a single hairline, no inner border.
- 19 | "Your spec / Sign in to mark your spec on the list." plate | tiers-dps 1440+390 | TierPage.astro | quiet "Sign in to mark your spec" link in the tab row.
- 22 | "Tier A · #2 · 1,004.2 DPS" repeats the highlighted row | tiers-dps-you 1440+390 | TierPage.astro | drop the line, keep spec name and Best in slot / Guide buttons.
- 13 | Three rows then 300px (1440) / 220px (390) of bare parchment | rankings-board | rankings/RankingsEmpty.svelte, Rankings.svelte | under the 3 rows show the same encounter's per-spec leaders. (390 dead space is f-selected's per handoffs; 1440 is not listed there.)
- density | "vs #1" column dropped at 390 | tiers-dps, tiers-dps-you 390 | tiers/TierRankRow.astro | -7.3% under the DPS value.
- 22 | Row subtitle wraps "Olympus XXVII · Fury Warrior" to 3 lines | rankings-board 1440 | RankRow.svelte | drop the spec when the filter already names it, or truncate.

## Handed off, noted only
Stranded "Comparing specs, not characters? Tier list" card (1440, rankings and board); Characters/Guilds collapsing to "Characters Change" at 390; Report column missing and Date on phone rows (rankings-you).

## Authentic, keep
Tier letters spanning row groups; real icons; class-coloured names; DPS with unit; "vs #1" at 1440; "14th of 212 Fury Warriors" as the point of the page; "You" chip on the one row; percentile colour legend; 31/20/0 talent split per row; compact title band.
