# AI-slop review, rankings / ratings / tiers (round 1, captures in first/)

First impression: rankings and tiers read as a player's tool (class colours, real icons, dense rows). ratings reads as generated documentation: six identical plates and a grey explainer under every line.

Source files (mapped by grep, not line-verified): web/src/components/rankings/{RankingsLanding,RankingsEncounters,RankingsAnswer,RankingsEmpty}.svelte, web/src/components/Rankings.svelte, web/src/components/ui/RankRow.svelte, web/src/components/score/RatingsPage.svelte, web/src/components/tiers/{TierPage,TierRankRow}.astro, web/src/lib/tiers/tier-copy.ts.

## Verdicts
- rankings: PASS WITH FIXES
- rankings-you: PASS WITH FIXES
- rankings-board: PASS WITH FIXES (thin: 3 rows, 500px of empty parchment at 1440)
- ratings: FAIL (tells 8, 7, 22, 19/21)
- tiers-dps: PASS WITH FIXES
- tiers-dps-you: PASS WITH FIXES

## Findings
- 22 | "Percentile" label repeated on every row | rankings, rankings-you, rankings-board at 390 | RankRow.svelte | column header already says it; show the bare number coloured by band.
- 22 | "On Onyxia, Normal ruleset" under "14th of 212 Fury Warriors" | rankings-you 1440+390 | RankingsAnswer.svelte | filter bar says it; delete.
- 8 | Ornate-corner frame on Encounters, plain double-border plate on Leaders by spec; sidebar cards are plates with an inner second border | rankings, rankings-you, tiers sidebar | rankings/*.svelte | one frame treatment by role; frame only the primary table, sidebar cards get a hairline.
- 19 | "Comparing specs, not characters? Tier list ›" stranded card in an empty right column | rankings, rankings-board 1440 | RankingsLanding.svelte / Rankings.svelte | make it a link in the title band or tab row.
- 13 | Board shows 3 rows then blank page | rankings-board 1440+390 | RankingsEmpty.svelte | show next-best filter result or neighbouring spec leaders below.
- 22 | "Sample · Forever Score", "of 2,100 · Epic", "The sample" tag on Epic tier | ratings 1440 | RatingsPage.svelte | "1,830 / 2,100" with Epic coloured; drop the sample labels.
- 7 | Grey explainer under every Forever Score part (Gear, Experience, Parses, Attendance, Preparation) | ratings 1440+390 | RatingsPage.svelte | keep name, bar, points; rules go in a tooltip.
- 21 | "A rating comes only from a public report. There is no leaderboard of ratings." / "Topping the meter while dying to avoidable damage with no flask on does not rate well." / "The tiers take the colours of item quality." | ratings | RatingsPage.svelte | cut all three; state the cap as a number.
- 8 | Six same-style bordered plates stacked, Performance plate double-bordered | ratings 1440+390 | RatingsPage.svelte | one primary plate (score + parts), rest as plain table sections.
- 19 | "What each component measures" and "Ranked against your spec, or absolute" are small grey paragraph walls | ratings | RatingsPage.svelte | table: component, what it counts, weight; one line each.
- 16/13 | Performance Rating "Last 6 fights" is a flat one-line sparkline | ratings 1440 | RatingsPage.svelte | real points with 65 marked, or print "65, up 25 over 6 fights".
- icon | empty leading column in Forever Score and Weights tables | ratings 1440 | RatingsPage.svelte | remove the column or fill it.
- 22 | "412 of 500" repeated under the 412 / 500 columns | ratings 390 | RatingsPage.svelte | one "412 / 500".
- 22 | Subtitle "Forever Score: your whole record, 0 to 2100. Performance Rating: one fight, 0 to 100." duplicates tabs; "2100" vs "2,100" | ratings 1440 | RatingsPage.svelte | delete; one number format.
- 19 | "Your spec / Sign in to mark your spec on the list." card | tiers-dps 1440+390 | TierPage.astro | quiet text link in the tab row.
- 22 | "Tier A · #2 · 1,004.2 DPS" repeats the highlighted row | tiers-dps-you | TierPage.astro | keep spec name and buttons only.
- density | "vs #1" column dropped at 390 | tiers 390 | TierRankRow.astro | show -7.3% under the DPS value.

## Authentic, keep
- Tier list: S/A/B/C/D serif letters spanning row groups, real icons, class-coloured names, DPS with unit, "vs #1" deltas, header "Best in slot, raid buffs, level 63, 3 min".
- Rankings answer: big "14th of 212 Fury Warriors" is the point; "You" chip on your row only; percentile colour legend; 31/20/0 per row.
- Leaders by spec with class-coloured bars, real encounters (Onyxia, Warden Kelthas).
- Compact title band, no hero, no gradients, no glow.
