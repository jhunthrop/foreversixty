# Round 9 UX verdicts: logs pages (1440 and 390)

Handed-off, needs-data and fixture items excluded. Captures: round9-captures/.

## logs-you: SHIP
- Compact "Logs" title, drop zone first, own reports, public reports, companion in the rail at 1440 and last at 390. Answer above the fold at both widths.
- Polish: the 1440 page ends ~200 px above the footer on sand (short fixture list).

## reports-list: SHIP
- Own row shows name, "42 pulls", date. Zone and bosses render when the API sends them; the fixture row has neither (needs-data), so not a defect.
- "Upload a log" sits in the title band at both widths. Empty sand below the single row is fixture.

## report-summary: SHIP
- Players table with class-colour bars, chart (folded at 390), ability panels, Mechanics, Buffs, At pull rail, icons filled.
- Polish: tab strip clips "Damage ta" at 390; add a right-edge fade.

## report-damage-done: SHIP
- Meter bars in class colour, Total row, Targets rail, 8-column ability table at 1440, hits/crit/avg/max sub-line at 390. No defects.

## report-healing: SHIP
- Green bars, Healed rail, same chrome as damage done. No defects.

## report-damage-taken: SHIP
- Mitigated line inside the plate, Sources rail, grey bars. No defects.

## report-deaths: SHIP
- Round 8 polish items still open, none blocks the page:
  - 390: "Damage taken over the fight" chart is open by default, every other report page folds it; death markers collapse to a "4" badge and the Group median label crowds the lines.
  - Collapsed death rows end in a bare "Open" at both widths; use a chevron or "Last 20 s >" like the expanded row.
  - 390: killed-by line wraps to three lines beside "Last 20 s >" and "Open".
