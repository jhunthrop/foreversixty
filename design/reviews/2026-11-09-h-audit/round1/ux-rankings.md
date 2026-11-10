# UX review: rankings, ratings, tiers (h-audit round 1, 2026-11-09)
Captures: design/reviews/2026-11-09-h-audit/first/. Mocks: design/mocks/directions/png/b1-rankings*, g-tiers. Specs: 2026-11-09-directions.md, 2026-10-11-rankings.md, 2026-11-06-tiers.md.

## Verdicts
- rankings (visitor, landing): SHIP WITH FIXES
- rankings-you (character board, signed in): SHIP WITH FIXES
- rankings-board (every class, fixture 3 rows): SHIP WITH FIXES
- ratings: NOT SHIPPABLE
- tiers-dps (visitor): SHIP
- tiers-dps-you: SHIP WITH FIXES (polish only)

## rankings (landing)
- MUST FIX, 390: the Characters/Guilds tabs collapse to "Characters  Change v". On the landing there is no encounter/class filter, so "Change" names nothing and Guilds is hidden. Show both tabs (mock b1-rankings-390 shows "Characters Guilds" as tabs). Source: web/src/components/rankings/RankingsLanding.svelte and the shared filter-bar collapse.
- MUST FIX, 1440: the rank numeral is tinted by percentile on the landing with no legend (legend exists only on the board). Either neutral rank numerals on the landing or add the percentile legend line under each list head. Source: lib/rankings/row-view.ts.
- MUST FIX, 390: every row repeats the word "Percentile" (36 times). A number takes its unit and nothing else; head it once ("DPS / Pctl" column head) and drop the per-row word. Source: lib/rankings/copy.ts, rank-row-view.ts.
- POLISH, 1440: right column holds only "Comparing specs, not characters? Tier list" and 600 px of bare parchment. Move that link into the band or add Your rank when signed in (rankings-you does).
- POLISH: Warden Kelthas ranks at 66 and 33 percentile on a 3-entry board (fixture), fine.
- Deviation from mock: mock was a plain "Choose an encounter" list; shipped adds top-3 per encounter and Leaders by spec. Better, accepted.

## rankings-you
- MUST FIX, 1440: "Your rank 14th" is purple and rank list row 14 purple, consistent; but the mock shows the filter row as plain text links, shipped shows five boxed selects plus Filters button (~48 px each) as a second band. Accepted as a control upgrade; ensure focus-visible rings exist (not provable from captures).
- MUST FIX, 390: per-row "Percentile 99" repeated 15 times (same as landing). Executed, Length, Build and Date are dropped on phone with no way to see them. Show Date under the name on phone or document the drop.
- POLISH, 1440: "You" pill sits in the name line and shifts the bar column by ~0 px; fine. Bars all the same class tan and almost equal length (fixture). Row 14 highlight is correct.
- POLISH: the "Your rank" panel lists Killed 27 Oct but the table's Date for the same row says 27 Oct; the panel duplicates the row. Keep, but add "Report" link per row as in the mock (mock has a Report column; shipped has none). Must reach a log from a rank row: MUST FIX, 1440, add Report link column. Source: components/rankings/RankingsAnswer.svelte.

## rankings-board
- MUST FIX, 390: dead parchment 450 px between a 3-row table and the footer. Footer should not float mid-screen; give main min-height so the footer sits at the fold or add the Tier list link here (hidden on phone; it only appears on desktop). Source: layout main wrapper.
- POLISH, 1440: Guild and Build columns from the mock are folded into the sub-line only for characters with a guild; "Olympus XXVII . Fury Warrior" wraps to two lines on the rogue/warrior row at 1440 (row 2). Widen the name column or drop the spec word when the guild is present.
- POLISH: Executed is 90% on every row, Length steps 1 s (fixture).

## ratings
- BLOCKER, 1440 and 390: empty icon gutter. The Forever Score rows, the weights table and the component list reserve a 48 px left gutter with no icon (gear, experience, parses, attendance, preparation have no glyphs). Never an empty icon square: draw the six component glyphs or remove the gutter. Source: components/score/RatingsPage.svelte and the shared StatTable/StatRow.
- BLOCKER, 390: the weights table collapses to "35 / 30 / 10" with no head saying DPS / Healer / Tank. Meaning is lost. Render three labelled cells or "DPS 35 . Healer 30 . Tank 10". Source: RatingsPage.svelte, StatRow phone prop.
- BLOCKER, 390: the Forever Score bars disappear and each row prints the number twice (412, then "412 of 500"). Keep the bar, print the number once with its "of 500" unit.
- MUST FIX, 1440: the Performance Rating sparkline is a flat yellow hairline with "Up 25 since the first"; a static explainer shows fake data without a plate or axis. Replace with a drawn example or remove. "Sample . Forever Score 1,830 of 2,100 . Epic" sits in the title band at 24 px purple: the band is meant compact; the explainer leads with a fake number. Move the sample into the Tiers card.
- MUST FIX, 1440: "What each component measures" body copy is ~12 px dim grey on the dark plate (about 3.8:1). Raise to 14 px text-secondary. Prose sections ("Ranked against your spec", "The cap") are unplated, capped at about 55 percent width and leave the right 40 percent empty; plate them or set them in two columns.
- MUST FIX, 1440: the two tabs "Forever Score / Performance Rating" are anchors to sections of one long page, styled as route tabs. Either make them real panels or make them in-page jump links without the active underline.
- MUST FIX, 390: the page is 4500 px. The Tiers legend (which defines the colours used above) sits after all prose; move it under the Forever Score card. Collapse the five component paragraphs behind Row disclosures (open attribute is forced).
- POLISH: "Weights, by role" comma-inverted heading; use "Weights by role".
- Copy: no em dashes found.

## tiers-dps
- Matches g-tiers: S/A/B/C/D letter gutter, class-coloured names on rows with spec icons, DPS bar, vs #1, Your spec panel with Sign in. SHIP.
- POLISH, 390: bar and "vs #1" dropped on phone; at 390 there is 90 px free after the name column on most rows; show vs #1 in the sub-line. Row 1 "Marksmanship Hunter" wraps to two lines because the name column is narrow.
- POLISH, 1440: the S row has an empty vs #1 cell; print "best" or a dash.

## tiers-dps-you
- Spec panel "Fury Warrior / Tier A . #2 . 1,004.2 DPS / Best in slot / Guide" present, row highlighted, first on phone. Matches spec.
- POLISH: panel title is neutral serif; tint the spec name with its class colour as the list does. The A letter is not tinted either; fine.
- POLISH, 390: panel sits above the list at 160 px tall before the first row (spec says it sits first on a phone): accepted.

## Systemic
1. "Percentile" word repeated per row on phone, all rankings pages.
2. Collapsed filter summary "Change v" with no object, on landing.
3. Empty icon gutters in score components (ratings).
