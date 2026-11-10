# ai-slop review round 2, logs and report pages (round2-captures/, 1440 and 390)

Verdicts: logs-you PASS WITH FIXES; reports-list PASS WITH FIXES (thin, fixture-bound, plus row shape); report-summary PASS WITH FIXES; report-damage-done PASS WITH FIXES; report-healing PASS WITH FIXES (thin, mostly fixture); report-damage-taken PASS WITH FIXES (minor); report-deaths PASS WITH FIXES (minor). No page trips visual tells 1, 2, 5, 8 (report plates are data, not stamped feature cards), 12, 13 or 15.

First impression: reads as a developer's tool. Real ability icons, class-coloured names, hits/crit/avg/max, a death recap with heals and a mitigated footnote.

## Fixed since round 1 (verified)
- logs-you companion copy cut to two lines; report rows now "date / 42 pulls · Trash only".
- summary and damage-done: "ilvl 25"; no "damage/healing" suffix on phone rows; rating line gone from This pull.
- deaths: phone row shows "677 / 996"; recap reads "Peaked at 68%. 266 healed."; "5 in this fight" gone from the title.
- damage-taken: chart header is one row.

## Remaining findings
logs-you
- 22 1440/390, Your reports row: title is a bare date ("17 Sep 2026") with a second unlabelled date ("20 Oct"). components/ReportRow.svelte. Fix: zone as title, one date.
- 22 1440/390, companion plate: "Pair it with your account" subhead plus grey "Press Sign in with Battle.net in the companion. A code is the fallback." still explains. lib/reports/copy.ts. Fix: drop the paragraph, keep the button.
- thin 1440: upload box ~170px of empty dashed area; right column ends at 455px, ~400px bare parchment below. components/logs/LogsPage.svelte. Fix: 80px drop strip, companion buttons inline.
- 8 mild: Your reports, Recent public reports, companion are three identical plates with a serif H2. Fix: rule-line lists, plate only on the upload.
- 17 mild: "Sample report" pill on the public row and in every report header band. Keep it once.
reports-list
- thin 1440/390: one 60px row, ~600px empty parchment. The emptiness is a one-report fixture artefact; the row shape is not: no zone, boss count or group DPS. components/ReportRow.svelte. Fix: columns (zone, bosses, kills, group DPS, uploaded) so one row still carries data.
- 22: same date-only title and unlabelled "20 Oct".
report-summary
- 22 1440 chart header: sub "Group 345 DPS" repeats the band's 345 group DPS; legend repeats "92.7 DPS" per player already in the Players table. components/report/TimeChart.svelte. Fix: no sub, names only.
- 22 1440/390 This pull: "In your group 5" and "Deaths 0" repeat the band and the Deaths column. SummaryTab.svelte. Fix: keep Started, Sim this fight, Copy link.
- 13 mild 390: "Damage over the fight" is a closed disclosure holding "Group 345 DPS"; chart hidden. TimeChart.svelte. Fix: render a short chart or drop it on phone.
- 21 mild: plate links repeat their own title ("Damage done by ability" / "Damage done ›", "Buffs" / "Buffs ›", "Mechanics" / "Mechanics ›"). SummaryTab.svelte. Fix: arrow only, or "All ›".
- band weight: meta line, "A real dungeon night, names changed." and the Sample pill before the tabs. Fold the note into the pill title.
report-damage-done
- 22 same chart sub and legend DPS repeats (1440). TimeChart.svelte.
- 22 "Total, 5 players" repeats the band. ActorTable.svelte. Fix: "Total".
- layout 1440: rail (This pull, Targets) ends ~800px, ability table runs ~1200px further. Fix: sticky rail.
- 13 mild 390: chart collapsed behind a disclosure.
report-healing
- thin 1440: one healer, so one-row table, one-line chart, "All 1" selector, a "Group median" line identical to the only player, 450px empty parchment. The one-healer fixture explains the emptiness; the "All 1" selector and duplicate median do not. TimeChart.svelte. Fix: hide both with a single source.
- 22 "100.0%" share and an Overheal column reading 0% on both rows; 390 "0% overheal" on each ability. ActorTable.svelte. Fix: hide overheal when all 0.
report-damage-taken
- 22 chart sub "Group 105 DTPS" repeats the band; 390 every row prints "36.2 Taken / s" after the amount. ActorTable.svelte. Fix: number only, column header carries the unit.
- keep: the Mitigated footnote.
report-deaths
- 22 "Deaths 5" in This pull repeats the band's "5 deaths". Fix: drop the row.
- 22 "The 20 s before it ›" on every death, next to Fold/Open. components/report/DeathsTab.svelte. Fix: one expand control per row.
- 13 mild 390: death-over-time chart collapsed behind a disclosure, event markers lost (1440 shows them well).

## Handoffs (other lanes, no fail)
/ratings (h-report-tail) and the f-selected, f-content, h-doll lists. None touch these seven pages.

## Authentic, keep
Real ability icons and class crests; class-coloured names on plates; hits/crit/avg/max; mitigated footnote with dodge/parry/miss; Healed and Sources side lists; death recap with heals and "Killing blow" tag; death markers on the 1440 chart; solid buttons, no glow; no emoji, no em dashes.
