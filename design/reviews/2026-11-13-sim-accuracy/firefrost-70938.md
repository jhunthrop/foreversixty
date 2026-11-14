# Sim accuracy 2026-11-13: Fire and Frost mage on 70938 and the 13 Nov pushes

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sa-firefrost-70938`. Site from origin/main 31ce477eb, engine pin 037798db1 (unchanged). Client: 1.60.1.70938 with the hotfix push as of the 14 Nov 01:05 UTC data run (b5bdd533b). Last lanes on these specs: 2026-11-12 `casters.md` (spell by spell, searches, rank rows), 2026-11-13 `crit-flags.md` (Ignite's ticks never crit) and `arcane-shadow-70938.md` (mana gems largest first); none of their closed items is redone here.

**Totals: 1 data bug fixed, new data (Blizzard's 13 Nov item hotfixes, pushes 114062 and 114063). The equip fill took every stat Wowhead Forever lists that a client row leaves out, and Wowhead still shows three items as they were before the hotfix, so the stats Blizzard removed came back in the planner and in simdb: Blessed Band of Light (272407) kept 1% hit, Channeler's Ring (272406) kept 8 Spirit next to its new spell power, the Bad Hat (260212) kept 23 shadow spell power. A stat type the shipped ItemSparse row states and the hotfixed row no longer does now stays off. Channeler's Ring is in every Fire and Frost level 60 set, Blessed Band of Light in Fire's bare 60 Alliance set. 0 engine bugs: no Fire or Frost spell, talent or trigger row moves between 70790, 70938 and the push, and the one mage line in Blizzard's 13 Nov notes (Frost Nova broken by proc and Blizzard damage) does not reach a boss fight. Engine pin unchanged.**

## Method

1. **Client and push.** `spellconst/mage.json`, `talents/mage.json` and `trainables/mage.json` of 70790 against 70938 (values compared at four decimals), then 70938 at 273623f0f (the Arcane lane's data) against b5bdd533b (the 14 Nov run that took the 13 Nov pushes). Item side: `items/mage.json` between the two 70938 runs, the hotfixed and the shipped ItemSparse from wago.tools (70938, `useHotfixes=1` and without), and `raw/wago-push/items.json`.
2. **Blizzard.** "WoW: Forever Hotfixes" (t/2403040), the whole thread as of 14 Nov: the 13 Nov section (posted 14 Nov).
3. **Engine.** No row moved, so the 12 Nov table of every Fire and Frost mechanic stands at the pin (the two 13 Nov changes, Ignite and the gems, are in it).
4. **Logs.** Our reports (Cloud SQL, read only): still the 8 the Arcane lane read, none newer than 13 Nov 18:38, no caster above 33. foreverlogs.gg answered again (it refused the API on 13 Nov 23:00): 29 public reports, four new since the 12 Nov lanes (2854 to 2857). Every mage in them was listed (`firefrost-70938/fl_page.py`).

## Client and push

- **Spells and talents.** 70790 to 70938: float formatting, two new NPC rows (Frost Armor 1283419, Blast Wave 1291176). 273623f0f to b5bdd533b: an effect 61 row on three NPC spells (Flame of Hakkar, Dark Energy, Atal'ai Poison). No Fire or Frost card, build or proc row changes.
- **Items (b5bdd533b).** Twelve mage-usable rows change. Darkspear Signet's mp5 halves at every rank and Militant's Band's rises by one (Blizzard: "now match Alliance-equivalent rings"), Souvenir Sea Shell's Stamina drops (Blizzard: "stats have been corrected"): the client rows carry these and the data has them. Three level 1 weapons at item level 30 (288906, 288910, 288915) gain 46 spell power through the Wowhead fill; none is on a published set and none has a source.
- **The bug.** Blizzard, 13 Nov: "Blessed Band of Light's Hit stat has been removed", "Channeler's Ring now has Spell Power (was Spirit)", "The Bad Hat no longer provides an incorrect amount of spell power." The pushed rows say the same (272406 stats 5, 31, 45; 272407 41, 42, 7, 6; 260212 drops 88). The data has `272406 {intellect 14, hit 10, spell_power 9, healing 9, spirit 8}` and `272407 {..., spirit 8, hit 1}`, and Bad Hat's 23 shadow power: each extra is Wowhead's pre-hotfix record (`spi 8`, `hitpct 1`, `shasplpwr 23`) put back by `fill_client_equip_stats` (planner) and `with_wowhead_equip_stats` (simdb), which fill any stat the client row leaves out because a hotfix item's on-equip link never reaches the client. Of the 30 items whose ItemSparse stat list the hotfixes restate on 70938, these three are the only ones the fill changes (`firefrost-70938/restated_fills.py`), and Blizzard's notes name all three.

## Real fights

| report | who | what it holds | verdict |
|---|---|---|---|
| 2856 (13 Nov, City of Dalaran) | one Frost-leaning mage, Frostbolt 8406 (rank 5, about level 26 to 30), Ice Lance, Cone of Cold, Blizzard with its chill | 5,390 events: 448 Frostbolts, 778 Blizzard ticks, 323 Flamestrike burn ticks, 162 Frostbite roots | below the scope's 50; nothing the earlier lanes have not read |
| 2855, 2857 (13 Nov, Shadowfang Keep) | four mages with 2 to 8 events, two with about 2,000 to 2,900 | level 20s | below scope |
| 2854 (12 Nov) | open world in the Barrens | no boss, level teens | below scope |

No log anywhere holds a Fire or Frost mage at 50 or above; the 12 Nov lane's three open log items stay open.

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| D1 | **The Wowhead equip fill put back stats a hotfix removed.** Blessed Band of Light 1% hit, Channeler's Ring 8 Spirit, Bad Hat 23 shadow spell power, in `items/<class>.json` and in simdb. New `pipeline/normalize/dropped_stats.py`: per item, the stat types the shipped ItemSparse row states and the merged row (play-session dump, wago export and item push) no longer does, minus anything the merged row still grants (45 also grants healing); both fills skip those keys for that item. Every other fill is unchanged: an item whose row a hotfix did not restate, or a stat the row never had, still takes Wowhead's on-equip amount (Staff of Jordan, Glowing Brightwood Staff). | Blizzard 13 Nov; push 114062, 114063; `raw/wago-push/items.json` | new data | site: `dropped_stats.py`, `normalize/wowhead.py`, `simdb/equip_fill.py`, the two callers; red first: `test_hotfix_dropped_stats.py` (module missing), `test_a_stat_a_hotfix_removed_is_not_filled_back` in `test_simdb_equip_fill.py` and `test_normalize_wowhead.py` |

## Before and after

The fix changes data, not the engine, and the brief allows no data run or simdb rebuild here, so the next data run (data.yml) applies it and the nightly re-ranks. Upper bounds from each published set's own weights (the phantom stat times its DPS a point, before any re-pick; hit at 10 points a percent):

| entry | ring(s) | Alliance | Horde |
|---|---|---|---|
| Fire bare 60 | Blessed Band (Alliance), Channeler's Ring | up to -1.25% | up to -0.57% |
| Fire raid 60 | Channeler's Ring | up to -0.16% | up to -0.16% |
| Frost bare 60 | Channeler's Ring | up to -0.46% | up to -0.18% |
| Frost raid 60 | Channeler's Ring | 0.00% | 0.00% |
| Arcane raid 60 | both | up to -1.13% | up to -1.14% |
| Arcane bare 60 | Channeler's Ring | up to -0.29% | up to -0.52% |
| Shadow bare 60 | Channeler's Ring | up to -0.23% | up to -0.21% |

Every band below 60 is untouched (both rings are level 60). The Bad Hat is on no published set.

## Measure at refresh (MR110)

Every band 60 entry that wears 272406 or 272407 (Fire, Frost, Arcane, Shadow, and the healers' sets that list them): the ring slots, weights, set_dps, after the next data run.

## Guide and card

No guide, card or coach text names either ring or the hat; nothing to change, no wow-player round.

## Gates

On the Mac (pipeline Python only, no sim or Go): ruff check on the touched files; pytest on `test_hotfix_dropped_stats.py`, `test_simdb_equip_fill.py`, `test_normalize_wowhead.py` and every `test_simdb_*` and `test_normalize_*` suite: 412 passed, 5 skipped. No engine, web or addon file changes.

## Shas

- Fork: none; the pin stays 037798db1.
- Site: the fix and this report (see the ledger row).

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (the 13 Nov section).
- wago.tools ItemSparse, ItemXItemEffect, ItemEffect for 1.60.1.70938, with and without hotfixes (13 Nov); `data/builds/1.60.1.70938/raw/wago-push/items.json`.
- foreverlogs.gg reports 2854 to 2857; api.foreversixty.gg reports table (read only).
