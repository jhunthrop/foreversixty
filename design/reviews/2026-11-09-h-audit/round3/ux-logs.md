# UX review, logs and reports (h-audit round 3)

Reviewed round3-captures/{logs-you,reports-list,report-summary,report-damage-done,report-healing,report-damage-taken,report-deaths}-{1440,390}.png against round 2. Handed-off items and fixture artefacts ignored.

## Verdicts
- logs-you: SHIP
- reports-list: SHIP
- report-summary: SHIP
- report-damage-done: SHIP WITH FIXES
- report-healing: SHIP WITH FIXES
- report-damage-taken: SHIP
- report-deaths: SHIP

## Round 2 must-fix, closed
1. Phone chart strip: now "Show chart" + caret + figure on summary, damage done, healing, damage taken, deaths.
2. Phone Report actions: "..." button at the answer's top right on all report pages; "This pull" plate no longer carries Sim/Copy at 390.
3. Deaths 390 meta: one line, "mobs died" dropped.
4. /logs Upload plate: heading sits clear of the frame flourish at 1440 and 390.

## Must fix
1. Pluralisation, 390, damage done ("Eviscerate 1 hits") and healing ("Healing Wave 1 hits"), ability row meta (components/report ability table): "1 hits" should read "1 hit". Fix: pluralise the count label.

## Polish
- All report pages 390: tab strip still clips "Damage ta" with no right-edge fade.
- Summary 1440: "What these mean" and "Rankings for Saltspine" sit on different columns; band still about 250 px.
- Reports-list and logs-you: date "20 Oct" without year; reports-list is one row (fixture).
- Damage taken 1440: Total row "Mitigated: ..." wraps to two lines inside the name column; give it a full-width row.
- Damage done 1440: ability rows have no per-row link (acceptable).

## Checklist
- [x] Compact title band, figure right at 1440, under title at 390
- [x] Class colours on plates, crests, spell icons
- [x] Units only on numbers, no em dashes
- [x] No sideways scroll at 390
- [x] Killing-blow treatment, collapsed death cards
- [x] Phone Show chart, Report actions, one-line deaths meta
- [x] Upload plate heading clear of frame
- [ ] "1 hit" singular (must fix 1)
