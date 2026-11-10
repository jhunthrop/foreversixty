# Density review: character page (visitor-60, owner-60-premium), round 1

Verdict: THIN (desktop dead space, closed panels with nothing in them, numbers without context)

Page job: show one character's standing and gear, and what to do next. Sim numbers, BiS and tier detail stay on their pages; link only.

## What the page can reach (not rendered)
- CharacterFight (web/src/lib/rankings/api.ts): percentile, execution_score, kill, spec, report_id/fight_index per fight. Best per encounter renders only parse + value + date + spec.
- /v1/characters/{region}/{ruleset}/{name}/history (lib/character/history-api.ts): score series (tier, encounter, parse_percentile per point), per-encounter kill series, gaps, all fights. Visitor page shows none of it; owner page hides it behind a closed "History" row.
- /v1/characters/.../rating and home answer (fetchHomeAnswer: guild_sims, spec_week): rank in guild, tier grade. Owner page shows them only as the link strip; visitor page has no rank at all.
- /v1/rankings/percentile, /v1/rankings/encounters: spec rank per boss and spec median.
- Gear: Upgrade list has per-slot HPS; GearDoll tiles show icons only at 1440.

## Dead space
1. Desktop right rail ends at y~960 (visitor) / ~960 (owner) while the left column runs to 1680 / 2060. About 700-1100px of empty ground under Builds seen.
2. Owner: "History" and "Coach history" are closed one-line rows with a heading each; the page ships two panels with no content visible.
3. Owner: standing strip (4 gold links) is a row of text links with no figures; "#3 of 14 healers" is the only number.
4. Visitor 1440: Activity is three rows of feed with the character name repeated on every row and the same crest each time.

## MUST additions
1. Fill the right rail: move/stack a "Guild and spec" panel under Builds seen: guild sim rank (#3 of 14), spec tier grade, next raid; visitors get guild rank from guild_sims too. Source: fetchHomeAnswer guild_sims/spec_week, lib/character/standing.ts. Change: web/src/components/Character.svelte, CharacterStanding.svelte (render as rail rows, not a link strip).
2. Best per encounter: add spec median / best-of-spec next to each parse and a delta vs last kill, plus execution_score and a kill/wipe mark. Source: CharacterFight.percentile, execution_score, kill; /v1/rankings/encounters for the median. Change: components/character/CharacterFightTable.svelte, lib/character-profile.ts.
3. Forever Score: show the delta against the guild/spec median and per-part rank (Gear 410 has no context); sparkline is a bare line with no axis or dates. Source: HistoryScorePoint[] from /history, rating endpoint. Change: components/character/CharacterScore.svelte, lib/character/score-parts.ts.
4. Open History by default with the score sparkline plus last 5 parses as a spark per boss (visitor page too, first page of data). Source: history-api.ts score[] and encounters[].points. Change: components/character/CharacterHistory.svelte, ScoreChart.svelte, ParseChart.svelte.
5. Merge Activity into the rail or cut name repetition: show the object first (Lucifron first kill, +6 places to #4) with timestamp; drop "Elyra Duskvale" prefix and crest on a page about her. Source: existing feed rows. Change: components/character/CharacterActivity.svelte.
6. Gear: show item level and enchant/gem state on the doll tiles or in a 1440 side list (already shown on 390 as names); at 1440 the paperdoll is icons only and the Upgrades card floats in the teal panel. Source: gear payload in /v1/characters/... ; link slot to Best in slot page. Change: components/character/GearDoll.svelte, CharacterGear.svelte.
7. Coach history: either show the last review title + date in the closed row or remove the panel until a review exists. Source: coach reviews API (owner). Change: components/character/CharacterOwnerSections.svelte.

## Phone (390)
Order is Gear, Best per encounter, Activity, History, Coach, Forever Score, Builds. Score (the key number) is below the fold at y~2600 of 3600: move the Score summary (score, tier, rating) up under the gear/upgrades or into the title band; collapse Builds seen to one line. Keep Upgrades after the gear list (works).

## Needs data
- Spec median and spec rank per encounter on the character payload (may need one call per boss to rankings/encounters).
- Item level and enchant/gem flags per slot if the character payload does not carry them (confirm in the payload).
- Per-boss attempt count and kill time trend on the best-per-encounter rows.
