# Sim-accuracy round 17: platform lane

Lane: web, addon, data and tools (no class auditing). Site branches cut from int/r17 (d7fc51fe7):

- `wt/r17-gearscore` 3ad1bedf4: the gear-score fix alone (coordinator: lands first, ships in addon 2.0.24).
- `wt/r17-platform`: items 2 to 5 and this report.

No engine change, no fork branch.

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| 1 | **Gear score multiplied rating points by per-1% weights** (addon, API, site upgrades panel). Every weight the scorers read is per 1% of a rating stat: `stat-weights.json` (what Data.lua ships as `weights`), the companion's engine sweep, a BiS band's `weight_per_percent`. Item stat lines carry rating points (Forever: 10 hit, 14 crit, 10 expertise, 12 dodge, 15 parry, 5 block, 10 haste per 1%; defense 1:1). Hit scored 10x, crit 14x, dodge 12x, parry 15x, block 5x, expertise 10x. | `data/builds/1.60.1.70755/gametables/combatratings.txt` (flat at every level 1 to 123); `addon/ForeverSixty/Gear.lua:246` (`value * weight`); `api/internal/sims/gearscore.go:311` (same, feeds the Forever Score gear part via `score_v2.go:359/372`, the gear panel via `rating/gear_panel.go:154/198`, loot values via `lootvalues.go:272/297`); `web/src/lib/addon/score.ts:36` fed band weights by `home/upgrades.ts` (per rating for hit/crit/dodge/parry/block, so those were right there; expertise per 1% and on-equip literal percents were not). Fury hands under warrior-fury weights: Primal Batskin Gloves (10 agi, 20 hit) showed **647**, correct **71**, and ranked first; Stormshroud Gloves 572 vs 50; Mage-fire head Bloodvine Goggles 406 vs 37; MM hunter neck Onyxia Tooth Pendant 612 vs 64. The order flips in every slot that has rating gear. | CONFIRMED-BUG | old (the stat-weights scorer predates Forever's rating itemization; no round looked at the addon/API scorers; round 16 fixed only the BiS ranker's expertise) | S 3ad1bedf4 (`wt/r17-gearscore`) |
| 1a | Fix, one table: `docs/contracts/fixtures/rating-per-percent.json`. The site imports it; the Go map (`sims.ratingPerPercent`) and the Lua table (`Rules.RATING_PER_PERCENT`) are pinned to it by tests; a vitest pins the fixture to every row of the active build's combatratings.txt. Same columns as `sim/cmd/leveling-bis` `ratingStatColumns`. Rating is divided before the weight; an on-equip literal percent (`percent_stats`, e.g. Fists of the Unrelenting's 1% hit and 1% crit) stays percent (site and API). The upgrades panel now feeds `scoreItem` per-1% band weights (`weight_per_percent`). | tests: `addon/tests/gear_spec.lua` "rating stats"; `api/internal/sims/ratings_test.go`; `web/src/lib/addon/score.test.ts`, `web/src/lib/home/upgrades.test.ts` | - | - | S 3ad1bedf4 |
| 1b | **Addon: caster haste was never scored.** `ITEM_MOD_HASTE_RATING_SHORT` mapped to `melee_haste` only, while every caster weight table names `spell_haste`. Rating haste now maps to `haste` and counts at `melee_haste` + `spell_haste` (it speeds attacks and casts alike), divided by 10. Mage-fire: 10 haste rating + 20 SP = 26 (was 20). Items JSON carries no haste key, so the API and site have no such gap. | `addon/ForeverSixty/Rules.lua:254`; stat-weights.json caster rows | CONFIRMED-BUG | old | S 3ad1bedf4 |
| 2 | Fury hands now. API and site: Gauntlets of Annihilation 116.5 > Fists of the Unrelenting 99.4 > Gloves of Enforcement 99.2 > Primal Batskin Gloves 71. Addon: Annihilation 116.5 > Enforcement 99.2 > Batskin 71, Fists 49.4, because its 1%/1% is an on-equip aura that GetItemStats is not known to return and the addon has no `percent_stats`. | `/dump GetItemStats("item:23072")` in game settles whether the client returns it | OPEN | - | - |
| 3 | **Skyborne base stats: the import never read `/statistics`.** Added a guarded capture: enrichment reads `/profile/wow/character/{name}-{surname}/statistics` and stores the body verbatim in `characters.bnet_statistics` (migration 0069). A 403/404 is skipped silently; any other error is logged (`op=statistics_unavailable`) and skipped. It never fails an import. Equipment, Specializations and Statistics now share one reader. | `api/internal/bnetimport/capture.go`, `character.go`; tests `bnetapi/profile_test.go` TestStatisticsReturnsTheRawBody, `bnetimport/forever_test.go` (records when it answers; keeps going on a 500) | CONFIRMED gap | old | S 03159ddd2 |
| 3a | Live check (read-only): production has 5 Forever characters; 4 are Skyborne (High Order and Windshaper), all **level 1**, none with a public profile. Blizzard answers **404** for the summary and `/statistics` on all four (namespace profile-forever-us, app token). No level-60 Skyborne exists to read. The engine is unchanged. | prod DB `select ... where race ilike '%skyborne%'` through cloud-sql-proxy with a read-only session; Blizzard API 2026-11-04 08:04 | OPEN | - | - |
| 4 | **Rebirth's tooltip had no description.** `$q` (EffectMiscValue_0) was not a token the spelltext filler knew, so "with 2200 health and $q mana." kept the raw token, and the site's game tip drops a sentence with a raw token. Rebirth 20748 effect 0 (resurrect, 329): base points 2200, misc value 2800 (ranks 1 to 5: 700/1200/1700/2200/2800 mana). Revive's five rows ("$q1 mana") are fixed by the same change. | `data/pipeline/spelltext.py` TOKEN and `_EFFECT_FACTS`; SpellEffect.csv 1.60.1.70755 (cached r16/int raw); `web/src/lib/tooltips/game-tip.ts:170` UNRESOLVED | CONFIRMED-BUG | old | S 4f1fdf041 |
| 5 | **Item tooltips dropped expertise.** Round 16 kept stat 37 in items JSON, but `STAT_KEYS` had no `expertise`, so the tooltip filtered it out. It now prints a white "+20 Expertise" line after hit, like "+10 Hit"; planner stat totals carry it too. | `web/src/lib/planner/types.ts` STAT_KEYS; `web/src/lib/items/tooltip.ts:203` | CONFIRMED-BUG | R16 fa02e9c6d (key added to items JSON without the tooltip) | S f81ca7bdc |
| 6 | Full web vitest on `wt/r17-platform`: 737 files, 8730 tests, all pass. | log in lane scratchpad | Aligned | - | - |

## Before/after

No DPS moves: no engine, APL, build or rank change. Relative changes are to gear scores only.

- Addon and API gear scores for items with hit, crit, expertise, dodge, parry or block fall to their correct size (hit /10, crit /14, dodge /12, parry /15, block /5, expertise /10). Non-rating stats are unchanged.
- API full-gear mark (the Forever Score gear denominator, the best band-60 set) drops for every spec that has rating gear: warrior-arms on 1.60.1.70291 5344 -> 995 (`gearscore_test.go`). The gear component is a ratio of the character's score to this mark, so a character with hit/crit-heavy gear loses part of the gear points it was over-credited.

## Files to regenerate

- `data/builds/1.60.1.70755/spelltips/*.json` (and any talent text carrying `$q`): the data job, for the Rebirth/Revive mana.
- `data/builds/1.60.1.70755/items/<class>.json`: already pending from round 16 (expertise key); the expertise tooltip line appears after it.
- No Data.lua regen: its `weights` are already per 1%, and the addon now converts at score time.

## Measure at refresh

- None from this lane.

## For other lanes

- **Gear-score fix touches every spec's addon/API/site gear scores** (no DPS). Lands alone from `wt/r17-gearscore` 3ad1bedf4.
- Migration **0069** (`bnet_statistics`): if another lane adds a 0069, renumber at integration. Only the deployed API migrates.
- Apps (iOS/Android) mirror the site's item tooltip stat list ("The apps' ItemTooltipModel keeps the same list", `tooltip.ts:149`): they need an expertise line too.
- Pre-existing: `api/internal/sims` TestRotationListsEveryRankOfAListedAbility fails on untouched int/r17 ("band 20 lists Execute or Bloodthirst, which it does not carry yet"). Not this lane.
- The API's `statScore` ignores an item's armour column while the addon scores armour from `RESISTANCE0_NAME`; tank gear scores differ between the two by the armour term. Not changed here.

## For the guide lane

- Nothing.

## Still open

- Addon 2.0.24 version and CHANGELOG bump: the integrator's release step.
- Whether GetItemStats returns an on-equip aura's literal percent (row 2).
- A level-60 Skyborne `/statistics` read (row 3a): re-run when a Skyborne profile is public.
