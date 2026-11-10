# ai-slop review round 3, logs and report pages (1440 and 390)

Verdicts: logs-you PASS; reports-list FAIL (thin row); report-summary 1440 PASS, 390 fail; report-damage-done 1440 PASS, 390 fail; report-healing 1440 PASS (fixture), 390 fail; report-damage-taken 1440 PASS, 390 fail; report-deaths fail (minor) at both widths.
No visual tell 1, 2, 5, 8, 12, 15. Round 2 fixes verified: companion paragraph gone, chart sub and legend DPS gone, "In your group"/"Deaths" rows gone, "All 1" and Group median hidden for one healer, overheal hidden, plate links arrows only, "Total" bare, DTPS rows number only.

## Remaining
1. 22 (390: summary, damage-done, healing, damage-taken, deaths) TimeChart.svelte closed disclosure "Show chart  Group 345 DPS / 42.9 HPS / 105 DTPS / 172 DTPS" repeats the band number and hides the chart. Deaths 390 also reads "Show chart" while open. Fix: drop the sub, label "Chart", render a short chart open on phone.
2. thin (reports-list 1440/390) ReportRow.svelte: row is only "Trash only / 42 pulls / 20 Oct", less than the public row on logs-you (zone, bosses, fights). Fix: zone, boss count, group DPS on the your-reports row.
3. 21 (report-healing 390) ActorTable.svelte: "1 hits". Fix: plural by count.
4. 22 (report-damage-taken 390) ActorTable.svelte: second figure (36.2, 18.0, 16.8) has no unit or label. Fix: "36.2/s".
5. 8 minor (report-deaths 1440, 390) DeathsTab.svelte: "The 20 s before it >" link plus Fold/Open per row; at 390 Fold/Open orphaned on its own line. Fix: one inline control at row right.

## Keep
Real ability icons, class-coloured names, hits/crit/avg/max, mitigated footnote, Healed and Sources lists, death recap with Killing blow tag and chart markers, solid buttons.
