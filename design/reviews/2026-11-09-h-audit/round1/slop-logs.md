# ai-slop review, logs and report pages (first/ captures, 1440 and 390)

Verdicts: logs-you PASS WITH FIXES; reports-list PASS WITH FIXES (thin); report-summary PASS WITH FIXES; report-damage-done PASS WITH FIXES; report-healing PASS WITH FIXES (thin with one player); report-damage-taken PASS WITH FIXES (minor); report-deaths PASS WITH FIXES (mobile duplicate value bug).

First impression: the report pages read as a developer's tool (real ability icons, class-coloured names on plates, hits/crit/avg/max columns, a death recap with healing and a dodge/parry footnote). The slop is in the Logs landing copy and in repeated explanatory qualifiers.

## Findings
logs-you
- 18/21/22 companion card (1440 right col, 390 last card) copy "Log live while you raid." + "Pair it with your account" + "Using the companion? Press Sign in with Battle.net in it. The code below is only for the fallback." Explains the obvious, and the code is not below. web/src/lib/reports/copy.ts lines 20-39. Fix: "Companion" with Get it / Pair code buttons, drop the paragraph.
- 22 Your reports row "Report, 17 Sep 2026" / "20 Oct": generic title, two unlabelled dates. components/ReportRow.svelte. Fix: title by zone ("Trash, 42 pulls"), one date.
- 13/thin Upload plate: ~170px of empty dashed box at 1440, right column ends at 480px, 350px of bare parchment below. components/logs/LogsPage.svelte. Fix: shorter drop strip, put Recent public reports beside Your reports.
- 8 (mild) Your reports, Recent public reports, companion are three identical plates with a serif H2 each. Fix: plain lists under rule lines, plate only on the upload.
- 17 (mild) "Sample report" pill on a public row, and again in every report header band. Fix: keep once, in the row; header carries zone only.
reports-list
- thin 1440, 390: one 60px row and ~600px empty parchment. components/ReportRow.svelte. Fix: columns for zone, bosses, kills, group DPS, uploaded date; show Recent public reports under it when the list is short.
- 22 title "Your reports" with a row titled "Report, ...". Fix: as above.
report-summary
- 22 chart sub "Group 345 DPS over the whole fight" (1440, card header, TimeChart) when the header already shows 345 group DPS and the filter says "Whole fight"; legend repeats "92.7 DPS" per player that the Players table shows. components/report/TimeChart.svelte. Fix: no sub; legend names only.
- 22 This pull card: "In your group 5", "Deaths 0", "Started 23:49:25" repeat the header ("5 players") and the Deaths column; "Ratings come with boss kills in raids." explains a non-feature. lib/rating/copy.ts:63. Fix: keep Sim this fight / Copy link and Started, remove the rating line on trash.
- 22 players sub "Paladin · item level 25" on every row (components/report/SummaryTab.svelte:141, ReportView.svelte:2014). Fix: "ilvl 25" or drop on trash.
- 22 390: every row has "damage"/"healing" under the amount ("5,782 damage"). Fix: number only, column header carries the unit.
- 13 390: "Damage over the fight" is a closed disclosure holding one line of text; the chart is gone. Fix: render a short chart or remove the section.
- 8 (mild) 1440: eight same-border plates; Mechanics, Healing by ability, Damage taken by ability hold one or two rows each. Fix: merge damage taken and healing into one two-column plate.
- title band: three text lines plus tabs plus filters is ~340px before content. Fix: fold the sample note into the pill tooltip, drop the extra line.
report-damage-done
- 22 same chart sub and legend DPS repeats (1440). TimeChart.svelte. Fix as summary.
- 22 "item level 25" per row (ReportView.svelte:2014); 390 "damage" suffix on every ability. Fix as summary.
- layout 1440: right rail ends after Targets at ~800px while the ability table runs 1200px more. Fix: sticky rail.
report-healing
- thin 1440: one healer, so a one-row table, a one-line chart, 450px of empty parchment. Fix: hide chart and Players "All 1" selector with a single source, merge into ability table.
- 22 "Total, 1 player", "100.0%" share, sub "1 overheal (0%)" (ActorTable.svelte:130, 244). Fix: omit total when one row, omit overheal when 0%.
report-damage-taken
- 22 chart sub repeated; the "Mitigated: 283 blocked..." footnote is good, keep. Chart header wraps its controls onto a second row unlike the other pages. Fix: one header row.
report-deaths
- 22 "5 in this fight" beside the Deaths title, big "5 deaths" in the band, "Deaths 5" in This pull: three times. Fix: keep the band only.
- 22 "The 20 s before it ›" repeated on every death. Fix: one toggle.
- bug 390: each row shows the amount twice ("-110" then "-110"; "+266" twice) where 1440 shows health after (677 / 996). Source components/report/DeathsTab.svelte mobile row. Fix: second line = "677 / 996".
- 21 recap sentence "the highest they stood at in that span was 68%" is clumsy and "Already dead: ..." chains names. Fix: "Peaked 68% HP. 266 healed."
- health bars are one green even at 30/996. Fix: shift to red under 35%.
- 13 390: death-over-time chart collapsed behind a disclosure, event markers lost.

## Authentic, keep
Real ability icons; class-coloured player names on plates; hits, crit, avg, max columns; the mitigated footnote; Healed and Sources side lists; death recap lines with healing; "Sim this fight"; solid teal buttons (no glow); no emoji, no em dashes seen.
