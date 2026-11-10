# UX review round 3: rankings, tiers (h-audit, 2026-11-10)
Captures: round3-captures/ at 1440 and 390. /ratings skipped (handed off). Handed-off f-selected items are noted, not failed.

## Verdicts
- rankings: SHIP
- rankings-you: SHIP
- rankings-board: SHIP
- tiers-dps: SHIP
- tiers-dps-you: SHIP

No non-handed-off must-fix remains on any of the five pages.

## Still open, all handed off (f-selected)
- rankings 390: tabs read "Characters  Change v", Guilds hidden. Rankings.svelte.
- rankings 1440: lone "Comparing specs, not characters? Tier list" card, right column empty below it. Rankings.svelte / rank-list.css.
- rankings-you 1440: no Report column; a row cannot reach its log. RankingsAnswer.svelte / rank-list.css.
- rankings-you 390: Executed, Length and Date dropped, no Date under the name.
- rankings-board 390: about 450 px bare parchment under the 3-row list. Board 1440: about 300 px bare parchment under the table.

## Closed since round 2
- tiers-dps 1440: S row now prints "best" in vs #1. Closed.
- rankings-board 1440: row 2 sub-line no longer wraps to three lines; it truncates to "Olympus XXVII . Fury Warr..." on one line. Wrap closed; see polish.

## Polish (not blocking)
- rankings-board 1440: the sub-line truncates mid-word ("Fury Warr..."). Drop the spec word when the guild shows, or widen the name column. rank-list.css, lib/rankings/row-view.ts.
- rankings-board 390: the third percentile "33" looks clipped at the right edge of the row; check the column width.
- tiers-dps 390: bar still dropped; "Marksmanship Hunter" wraps to two lines. Move vs #1 into the sub-line or shorten the name column gutter.
- tiers-dps-you 1440 and 390: panel has about 40 px of empty space under the Best in slot / Guide buttons. Trim bottom padding.
- tiers-dps-you: panel spec name is parchment tan, not the class colour of the row. Tint it.
- rankings-you: all bars equal length and class tan (fixture).

## Checklist
- [x] percentile word once, not per row
- [x] landing numerals neutral
- [x] tier list matches mock, "best" on the S row
- [x] Your rank panel first on phone, You pill, legend, pagination (rankings-you)
- [x] tiers-dps-you row highlighted, panel first on phone
- [ ] landing tabs at 390 (f-selected)
- [ ] Report column (f-selected)
- [ ] board dead space (f-selected)
- [ ] Date on phone rows (f-selected)
