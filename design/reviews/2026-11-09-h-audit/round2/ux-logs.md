# UX review, logs and reports (h-audit round 2)

Reviewed: round2-captures/{logs-you,reports-list,report-summary,report-damage-done,report-healing,report-damage-taken,report-deaths}-{1440,390}.png against round 1 and the current design system. Accepted per the brief: /logs title-only band (parity mock), ornate lead-panel frame on the Upload plate. Fixture artefacts ignored.

## Verdicts
- logs-you: SHIP WITH FIXES
- reports-list: SHIP
- report-summary: SHIP WITH FIXES
- report-damage-done: SHIP WITH FIXES
- report-healing: SHIP WITH FIXES
- report-damage-taken: SHIP WITH FIXES
- report-deaths: SHIP WITH FIXES

## Round 1 items closed
- Row titles gold (logs, reports-list): closed.
- Pair-it heading now strong, paragraph shortened to two lines: closed.
- Killing blow: Tag, red title, red damage bars, green heal row; four other deaths collapsed to header + "Open / The 20 s before it": closed.
- Overheal column on healing at 1440: closed.
- Damage taken filter row on one line at 1440: closed.
- Phone meta one line on summary, damage done, healing, damage taken: closed.

## Must fix
1. Phone chart strip, 390, summary / damage done / healing / damage taken / deaths (components/report/TimeChart.svelte, Disclosure summary): still "> Group 345 DPS" with a 6 px caret and no "Show". Spec: "Show chart" plus caret, figure after, 44 px target. Not fixed since round 1.
2. Phone report actions, 390, all report pages (components/report/ReportHead.svelte): no 44 px "···" Report actions button at the answer's top right; "Sim this fight / Copy link" still sit in the "This pull" plate at the page foot (y about 1900 on summary, 3100 on damage done). Fix: add the Menu button below 768 holding Sim this fight and Copy link.
3. Deaths, 390 (components/report/ReportHead.svelte meta): meta wraps to a second line and orphans "died" ("3 mobs / died"). Fix: drop "· {n} mobs died" from the phone meta, or truncate to one line.
4. /logs, 1440 and 390, Upload plate (components/logs/LogsPage.svelte, lead-panel frame): the top-left corner flourish collides with the "Upload a log" heading (flourish x 103 to 125, heading starts x 117 at 1440; same at 390). Fix: heading padding-left and padding-top at least 28 px inside the frame so no glyph touches an ornament.

## Polish
- Summary 1440: "What these mean ›" (y 1759) and "Rankings for Saltspine ›" (y 1721) are 38 px apart on different columns; place both on one footer line under the grid. The band is still about 250 px (description line "A real dungeon night, names changed." plus tabs plus controls); move the description into the Sample report Tag tooltip.
- Summary 390: "Damage done ›" / "Damage taken ›" / "Healing ›" links drop to their own row under each plate title (components/report/SummaryPanels.svelte); keep them on the title line at right, as the Players plate does.
- Damage done 1440 (AbilityTable / ActorTable): still no per-row "›" to the player view.
- Reports-list 1440 and 390: date shows "20 Oct" with no year; no 1 px rule under the band. Page is one row high (fixture).
- All report pages 390: tab strip clips "Damage ta" with no fade; add a right-edge fade on the strip.
- /logs 1440: the right rail paragraph is 13 px muted; acceptable, keep at two lines.

## Checklist
- [x] Compact title band, answer figure right at 1440, under title at 390
- [x] Class colours on plates, circular crests, spell icons present
- [x] Numbers carry units only; no em dashes in captured copy
- [x] No sideways scroll at 390
- [x] Killing-blow treatment
- [x] Collapsed death cards, recap header always visible
- [x] Gold row titles
- [ ] Phone chart Show control (must fix 1)
- [ ] Phone Report actions button (must fix 2)
- [ ] One-line phone meta on deaths (must fix 3)
- [ ] Upload plate heading clear of the frame (must fix 4)
