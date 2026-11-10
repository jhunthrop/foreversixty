# Density review: character page (visitor-60, owner-60-premium), round 2

Verdict: DENSE WITH ADDITIONS (round 1 was THIN). Items in handoffs.md (doll item level/enchant, floating Upgrades card, spec median/rank, ilvl/enchant payload) are excluded.

Page job: one character's standing and gear, and what to do next.

## Fixed since round 1
- Right rail now runs level with the left column (owner 1440: rail ends y~1610 vs left ~1640; visitor rail ends ~1240 vs left ~1415, no more 700px hole).
- Standing is a rail panel with rows (#3 of 14 healers in Cx66 Olympus, tier grade, next raid), not a link strip. Owner only.
- Forever Score has a delta row ("Since 20 Oct", sparkline, +310) plus tier and rating; parts carry bars.
- Activity no longer prefixes the character name; rows lead with the object and a timestamp. The crest repeated on each row is still redundant on a page about her.
- Title band carries +128.9 HPS from upgrades and Sim this character; owner adds the "Upgrade: Devout Skirt, Stratholme: Baron Rivendare" line.

## Remaining additions, ranked
1. MUST: History panel (owner) is still a heading over two closed one-line rows ("Score and parses 7 fights", "Coach reviews"), y~1480-1640 at 1440 and ~1240 (displayed) at 390. Open Score and parses by default with the score series and a per-boss spark of the last parses. Source: /v1/characters/{region}/{ruleset}/{name}/history, lib/character/history-api.ts (score[], encounters[].points). Files: web/src/components/character/CharacterHistory.svelte, ScoreChart.svelte, ParseChart.svelte. Visitor page has no history at all: render the same first page there.
2. MUST: Best per encounter rows are parse + value only. Add kill/wipe mark, execution_score and delta against the previous kill of that boss. Source: CharacterFight.percentile, execution_score, kill (web/src/lib/rankings/api.ts); previous fight from the history payload (no spec median needed, that is in handoffs). Files: components/character/CharacterFightTable.svelte, lib/character-profile.ts. Make the boss name the drill-down (report_id, fight_index exist).
3. Should: Visitor page has no Standing. guild_sims rank and tier grade come from fetchHomeAnswer/lib/character/standing.ts; if the public rating endpoint carries them, show the same rail rows. If owner-only by design, leave it.
4. Should: Forever Score parts (Gear 410 etc.) have bars but no rank or delta per part. Source: rating endpoint parts and HistoryScorePoint (first vs last). Files: CharacterScore.svelte, lib/character/score-parts.ts. Sparkline has no date/value on hover.
5. Should: Builds seen rows are spec + date only; add fight count or best parse per build from /history fights, link to the planner.
6. Optional: drop the Activity crest per row. File: CharacterActivity.svelte.

## Phone (390)
- Score is still low on the page (~2750 of 3632 px), after Best per encounter and History. Move Forever Score directly after Upgrades or into the title band; Standing and Builds seen fold to one row each. File: Character.svelte section order.
- Gear list with item names at 390 is good. Keep.
- Once History opens by default it collapses to the score spark plus a "7 fights" row at 390.

## Needs data
- Spec median and spec rank per boss (handoff).
- Item level and enchant/gem per slot (handoff).
- Attempt count and kill-time trend per boss.
- Visitor guild rank if the public rating endpoint does not expose it.
