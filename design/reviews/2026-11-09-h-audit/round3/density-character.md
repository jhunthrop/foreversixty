# Density review: character page (visitor-60, owner-60-premium), round 3

Verdict: THIN. Round 2 MUSTs 1 and 2 are not fixed; the page looks unchanged from round 2 apart from the title band. Handoffs.md items excluded.

Page job: one character's standing and gear, and what to do next.

## MUST
1. History panel (owner, 1440 y~1500-1660; 390 y~1770-1890): still a heading over two closed one-line rows ("Score and parses 7 fights", "Coach reviews"). Open Score and parses by default with the score series and a per-boss spark. Source: /v1/characters/{region}/{ruleset}/{name}/history, lib/character/history-api.ts (score[], encounters[].points). Files: character/CharacterHistory.svelte, ScoreChart.svelte, ParseChart.svelte. Visitor has no History at all; render the same first page.
2. Best per encounter rows (both widths) are still parse + value + "date, spec, Execution n%". Add kill/wipe mark and delta against the previous kill of that boss (history payload), and make the boss name the drill-down to the report (report_id, fight_index exist). Files: character/CharacterFightTable.svelte, lib/character-profile.ts.
3. Visitor 1440: right rail ends y~1200, left ~1415, so ~200 px of empty ground under Activity. Fill it with the Standing rows if the public rating endpoint carries rank and tier (lib/character/standing.ts), or with the History spark from item 1.

## Should
- Forever Score parts: bars but no rank or delta per part (rating endpoint parts, HistoryScorePoint first vs last); sparkline has no hover value. File: CharacterScore.svelte.
- Builds seen: spec + date only; add fights per build from /history and link to the planner.
- Activity: the crest on every row is redundant on her own page (CharacterActivity.svelte).

## Phone (390)
- Forever Score still sits at y~1950 of 3587, after History. Move it directly after Upgrades or into the title band. Character.svelte section order.
- Gear list with item names is good. Keep.

## Needs data
- Attempt count and kill-time trend per boss; visitor guild rank if the public endpoint lacks it.
