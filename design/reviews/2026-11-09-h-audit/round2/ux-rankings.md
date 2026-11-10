# UX review round 2: rankings, tiers (h-audit, 2026-11-10)
Captures: round2-captures/ at 1440 and 390. /ratings skipped (h-report-tail). Items in handoffs.md (f-selected) noted, not failed.

## Verdicts
- rankings: SHIP WITH FIXES (open items are all f-selected handoffs)
- rankings-you: SHIP WITH FIXES (open items are all f-selected handoffs)
- rankings-board: SHIP WITH FIXES (open items are f-selected handoffs plus one polish)
- tiers-dps: SHIP
- tiers-dps-you: SHIP

## Round-1 items closed
- "Percentile" word per row at 390: gone on all three rankings pages; the column is headed once and the percentile prints as a bare coloured number under the DPS. Closed.
- Landing rank numerals tinted with no legend: now neutral on the landing (1440 and 390). Board and you pages keep tint with the legend line under the title. Closed.

## rankings (landing)
- Open, f-selected: 390 tabs still read "Characters  Change v"; Guilds hidden. MUST FIX, Rankings.svelte / shared filter collapse.
- Open, f-selected: 1440 right column is the lone "Comparing specs, not characters? Tier list" card with 40 percent of width empty below it. Polish.
- Deviation from mock b1-rankings: top-3 per encounter plus Leaders by spec instead of a plain encounter list. Accepted.

## rankings-you
- Open, f-selected: no Report column at 1440; a rank row cannot reach its log. MUST FIX, RankingsAnswer.svelte / rank-list.css.
- Open, f-selected: 390 drops Executed, Length and Date with no Date under the name. Must fix, same owner.
- Fine: Your rank 14th panel first on phone, row 14 highlighted with You pill, legend present, pagination present.
- Polish: all bars equal length and class tan (fixture).

## rankings-board
- Open, f-selected: 390 has about 450 px of bare parchment between the 3-row list and the footer. MUST FIX, layout main min-height or Tier list link on phone.
- Polish, 1440: row 2 sub-line "Olympus XXVII . Fury Warrior" wraps to three lines in the name column. Widen the name column or drop the spec word when a guild shows. Source: rank-list.css name column, lib/rankings/row-view.ts sub-line.
- Polish, 1440: about 300 px bare parchment under the table before the footer.

## tiers-dps
- Unchanged from round 1 and matches g-tiers. SHIP.
- Polish, 390: bar and vs #1 dropped; Marksmanship Hunter wraps to two lines. Put vs #1 in the sub-line.
- Polish, 1440: S row has an empty vs #1 cell; print "best".

## tiers-dps-you
- Unchanged and matches spec: Fury Warrior / Tier A . #2 . 1,004.2 DPS / Best in slot / Guide, row highlighted, panel first on phone. SHIP.
- Polish: tint the panel's spec name in class colour.

## Checklist
- [x] percentile word once, not per row
- [x] landing numerals neutral
- [x] tier list matches mock
- [ ] landing tabs at 390 (f-selected)
- [ ] Report column (f-selected)
- [ ] board dead space at 390 (f-selected)
- [ ] Date on phone rows (f-selected)
