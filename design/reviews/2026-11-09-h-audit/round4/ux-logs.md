# UX review, logs and reports (h-audit round 4)

Reviewed round4-captures/{logs-you,reports-list,report-summary,report-damage-done,report-healing,report-damage-taken,report-deaths}-{1440,390}.png against round 3. Handed-off items and fixture artefacts ignored.

## Verdicts
- logs-you: SHIP
- reports-list: SHIP
- report-summary: SHIP
- report-damage-done: SHIP
- report-healing: SHIP
- report-damage-taken: SHIP
- report-deaths: SHIP

## Round 3 items, closed
1. "1 hit" singular: damage done 390 (Eviscerate "1 hit") and healing 390 (Healing Wave "1 hit") now read correctly.
2. Damage taken Mitigated line: now its own full-width row under Total at 1440 and 390, no wrap inside the name column.

## Must fix
None.

## Polish (carry, none fails a page)
- All report pages 390: tab strip clips "Damage ta" with no right-edge fade (still).
- Summary 1440: "What these mean" and "Rankings for Saltspine" sit on different columns; band about 250 px.
- Reports-list and logs-you: date "20 Oct" has no year.
- Deaths 1440 meta still says "3 mobs died" while 390 drops it; make both the same.
- Deaths 390: opened chart legend stacks six rows before the plot; one wrapped row would put the chart higher.

## Checklist
- [x] Compact title band, figure right at 1440, under title at 390
- [x] Class colours on plates, crests, spell icons
- [x] Units only on numbers, no em dashes
- [x] No sideways scroll at 390
- [x] Killing-blow treatment, collapsed death cards
- [x] Phone Show chart, Report actions, one-line deaths meta
- [x] Upload plate heading clear of frame
- [x] "1 hit" singular
