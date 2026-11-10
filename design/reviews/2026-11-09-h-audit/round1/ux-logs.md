# UX review, logs and reports (h-audit round 1)

Reviewed: first/{logs-you,reports-list,report-summary,report-damage-done,report-healing,report-damage-taken,report-deaths}-{1440,390}.png against design/specs/2026-11-09-directions.md, 2026-11-07-logs-reports.md and png/b1-*. Fixture artefacts (Fixture#1234 chip with no character, "Sign in" on the sample report, names) are not findings.

## Verdicts
- logs-you: SHIP WITH FIXES
- reports-list: SHIP
- report-summary: SHIP WITH FIXES
- report-damage-done: SHIP WITH FIXES
- report-healing: SHIP
- report-damage-taken: SHIP
- report-deaths: SHIP WITH FIXES

## Must fix
1. /logs, 1440 and 390, title band: shows only "Logs". Spec 11-07 section 11 wants the meta "Upload a combat log or log live, then read every pull." and the answer "Your latest report · {age}" + "{zone} · {n} bosses · {n} fights" + "Open it ›". The signed-in visitor's answer (their newest report) is a row in the middle of the page. Fix: add meta and answer to the band (web/src/pages/logs, components/logs/*).
2. /logs, 1440 right column: "Pair it with your account" is a 12 px muted caption and "Using the companion?..." a 3-line muted paragraph; mock b1-logs has a 15 px strong sub-heading and no paragraph. Fix: sub-heading at strong weight; shorten or drop the paragraph (pairingCopy.companionFirst).
3. /logs, 1440 and 390, "Upload a log" plate: ornate double-corner frame with heavy flourishes reads as a different language from the other plates; mock draws thin corner ticks only. Fix: thin the flourish asset to the mock's.
4. /logs, 1440, "Your reports" and "Recent public reports" row titles are white bold; mock and the link rule make a row title gold. Same on /reports row (reports-list 1440 and 390). Fix: ReportRow.svelte title uses the gold link colour; hover underline.
5. Deaths, 1440 and 390 (components/report/DeathsTab.svelte): the killing-blow row is not distinguished. Mock b1-report-deaths tints the row, shows a "Killing blow" Tag and uses red damage bars; shipped shows green health bars on every row and only "dead" as text. Fix: tint the lethal row, add the Tag, red Meter for damage rows, green for heals.
6. Deaths, 1440 (4,966 px) and 390 (5,353 px): all five death cards are fully open, 10 to 11 rows each, the Damage-taken chart sits at the very bottom. The "who died and to what" answer is five headers 400 px apart. Fix: open the first death, collapse the others to their header line (time, name, killer, "The 20 s before it ›"), or add a one-line index of the five deaths at the top. Spec 12 says the recap never collapses; the header line stays visible, so the rule holds.
7. Phone report head, all report pages at 390: meta wraps to two lines ("· 5 players" orphaned) and Sim this fight / Copy link sit in the rail at the page foot. Spec 11-07 6.1: one-line meta, a 44 px "···" Report actions button at the answer's top right. Fix: components/report/ReportHead.svelte add the Menu button below 768; drop "· {n} players" from the phone meta.
8. Phone chart strip, summary/damage done/damage taken/healing/deaths at 390: collapsed Disclosure reads "› Group 345 DPS over the whole fight" with a 6 px caret and no "Show". Spec: the strip reads "Show ▾" with a 44 px target. Fix: Disclosure summary text "Show chart ▾" with the figure after it, 44 px high.

## Polish
- reports-list 1440: row shows "Uploaded 20 Oct 2026" in the mock, "20 Oct" shipped; fine, but add the year when older than the current one. Band fade: shipped band has no bottom rule, mock has a 1 px line under the band.
- Summary 1440: the title band is about 250 px tall (description line plus tab row plus control bar) against the mock's 208 px. The "A real dungeon night, names changed." line pushes the tabs down on the sample only; set it to the Tag's tooltip or move it into the Sample report Tag row.
- Summary 1440: "What these mean ›" (x 96) and "Rankings for Saltspine ›" (x 1004) sit at different baselines, 6 px apart. Align.
- Damage done 1440: the Damage done player table has no per-row "›" to the player view that the mock draws; add it.
- Damage taken 1440: the chart header wraps the Players select and Log scale onto a second row (label overlap at 1440); keep on one row from 1280.
- Healing 1440 and 390: one healer row only is a fixture fact; the card still reads thin. Show "Overheal" as its own column at 1440 on the player table (it is in the sub line today).

## Checklist
- [x] Compact title band, no hero, answer figure at right (1440), under title (390)
- [x] Class colours on plates, crests circular, spell icons present, no blank squares
- [x] Numbers carry units only ("345 group DPS", "damage" under figures on phone)
- [x] First player row above y 500 at 390 (summary ~451, damage done ~451)
- [x] No sideways scroll at 390 (tab strip scrolls inside its line as specced)
- [x] No em dashes in the captured copy
- [ ] /logs answer and meta (item 1)
- [ ] Killing-blow treatment (item 5)
- [ ] Phone Report actions button (item 7)
- [ ] Phone chart Show control (item 8)
