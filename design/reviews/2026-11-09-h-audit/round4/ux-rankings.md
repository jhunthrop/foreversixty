# UX review round 4: rankings, tiers (h-audit, 2026-11-10)
Captures: round4-captures/ at 1440 and 390. Handed-off f-selected items noted, not failed.

## Verdicts
- rankings: SHIP
- rankings-you: SHIP
- rankings-board: SHIP
- tiers-dps: SHIP
- tiers-dps-you: SHIP

No non-handed-off must-fix on any page.

## Still open, handed off (f-selected)
- rankings 390: tabs read "Characters  Change v", Guilds hidden. rankings 1440: lone "Comparing specs" card, right column empty.
- rankings-you 1440: no Report column. rankings-you 390: Executed/Length/Date dropped.
- rankings-board: bare parchment under the table (1440 about 300 px, 390 about 200 px).

## Changed since round 3
- tiers-dps-you: the Fury Warrior panel is now bare text ("Fury Warrior / Best in slot > Guide >") with no plate, right column at 1440, above the list at 390. The 40 px bottom padding is gone. Reads as secondary; acceptable.

## Polish (not blocking)
- tiers-dps-you: "Fury Warrior" label is parchment tan, not the class colour. Tint it.
- rankings-board 390: third percentile "33" still clipped at the right edge (rank-list.css column width).
- rankings-board 1440: sub-line truncates "Fury Warr..." (row-view.ts).
- tiers-dps 390: "Marksmanship Hunter" wraps to two lines.
- rankings-you: bars equal length and class tan (fixture).

## Checklist
- [x] percentile word once, not per row
- [x] landing numerals neutral
- [x] tier list matches mock, "best" on S row
- [x] Your rank panel first on phone, You pill, legend, pagination
- [x] tiers-dps-you row highlighted, panel first on phone
- [ ] landing tabs at 390, Report column, board dead space, Date on phone (all f-selected)
