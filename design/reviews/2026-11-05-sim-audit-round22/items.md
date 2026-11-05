# Round 22, items lane: weapon damage ranges against the 1.60.1.70790 client

Seed: round 21 compared weapon speeds only. This lane compares every simdb.bin weapon's damage range (and speed again) with what the 70790 client gives, audits the formula in data/pipeline, and spot-checks the Wowhead supplement.

Summary: **zero confirmed sim bugs.** The pipeline's damage formula matches the client and Wowhead Forever's tooltips on every weapon the sim uses. The committed simdb.bin is stale on 8 weapons: hotfixes landed after the 4 Nov 20:43 UTC data run. None of the 8 is in any frozen BiS list, so no spec's DPS moves. A new permanent test names them.

## How the client gives a weapon's damage

ItemSparse 70790 has no MinDamage/MaxDamage columns. Damage comes from a curve (data/pipeline/normalize/weapon_curves.py):

- `dps = ItemDamage{OneHand|TwoHand|Ranged|Wand|Thrown}[ItemLevel][OverallQualityID]`, times 2/3 (one-hand) or 29/39 (two-hand) when the item has the caster bit (Flags_4 & 0x100).
- `base = dps * ItemDelay/1000`
- `avg = (dps + QualityModifier) * ItemDelay/1000`
- `min = trunc(avg - base*DmgVariance/2)`
- `max = round(avg + base*DmgVariance/2)`

Inputs come from ItemSparse and Item, with the committed DBCache seed (raw/hotfixes, add-only) and wago.tools' layered hotfix pushes (raw/wago-hotfixes) merged on top (pipeline/hotfix_merge.py).

**The ItemDamage tables:** I fetched all five for 70790, plain and with `useHotfixes=1`. Each pair is byte-identical (100 item-level rows), so no hotfix touches the curves.

**Rounding:** `max` uses Python's `round()`, which rounds an exact .5 to the even number. The client rounds half up (gear.py `round_half_up`). I scanned every client weapon row: none has a high end within 0.001 of a .5 and none has a negative low end. On this build the two rules give the same answer every time. Recorded as latent, not a bug.

## Cross-cutting findings

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| I1 | The formula matches the client on every client-backed weapon in simdb.bin, except the 8 in I2. | I dumped all 1936 weapons in simdb.bin: 1541 from the client, 57 from Wowhead, 338 from classic-db. For every client weapon I recomputed damage from wago's 70790 ItemSparse and Item (plain, plus the committed seed, plus wago's layered pushes for 70245..70790) and the five ItemDamage tables. 1534 of 1542 are identical (min, max and speed). | Aligned | - | - |
| I2 | **simdb.bin is stale on 8 weapons:** the 70790 hotfix pushes wago recorded after the data run. (1) **Earthshaker 288900:** the bin has 40-61 at speed 3.5, which is the item-level-10 curve. The client hotfix row is item level 30, giving 87-132. Wowhead Forever's tooltip also shows item level 30, 87-132, 31.29 DPS. (2) **Seven thrown weapons (287059-287064, 287070):** the current hotfix rows carry DmgVariance 0.6, so Footman's Pocketknife resolves to 6-13. The bin has 9-10 (no spread). Wowhead's tooltips still show 9-10. | data/diffs/1.60.1.70755__1.60.1.70790.json has Earthshaker going from item level 30 to 10 in the data run. wago-hotfixes ItemSparse rows `_HotfixBuild=1.60.1.70790`: 2791 rows. Tooltips: https://nether.wowhead.com/forever/tooltip/item/288900 and /287059 (fetched 5 Nov ~07:00 UTC). Earthshaker has no loot.json source and appears in no BiS file. No thrown weapon's damage reaches any APL. | CONFIRMED (data, not code) | new-data | - (regen at refresh) |
| I3 | **The wowhead supplement (57 weapons, 21 of them in frozen BiS) follows the Forever curve.** Recomputing each from its Wowhead item level, quality, subclass, speed and dmgrange: 48 match exactly. The other 9 are caster weapons (Spellshifter Rod, Spire of Hakkar, Kindling Stave, Charstone Dirk, Inventor's Focal Sword, Zum'rah's Vexing Cane, Soul Harvester, Farraki Kris, Dune Hacker). These match once the 2/3 or 29/39 caster cut is applied. Wowhead's carry file has no Flags_4. | raw/wowhead/items.json; scratch supcurve.py. All 21 BiS supplement weapons are carried records, and their Forever tooltips now return 404 ("Entity not found"). Examples: 6830 Bonebiter, 7723 Mograine's Might, 10844 Spire of Hakkar, 12651 Blackcrow, 17766 Princess Theradras' Scepter. So the tooltip spot-check could not run on them. The curve check is the independent check. | Aligned | - | - |
| I4 | **Spot-check against live Wowhead Forever tooltips:** Fiery War Axe 870 116-175 @3.6 (sim the same). Thunderfury 44-115, Atiesh 130-243, Torturing Poker 22-45 (QualityModifier -2.5), Kris of Unspoken Names 39-86, Earthshaker 87-132 (the client value). All agree with the client formula. Wowhead's raw dmgmin/dmgmax are unrounded floats (Atiesh 130.17-243.17, Torturing Poker 21.75-44.75); the tooltip shows the sim's integers. | Comparing the bin with Wowhead's raw data: 1397 client weapons identical. 13 differ only by float noise or rounding, and the tooltip shows the sim's integers on every one checked. | Aligned | - | - |
| I5 | **Thrown "Satchel of ... Bombs" 285275-285278 (QualityModifier -5.3 to -20.3):** the client formula gives 0-11, 1-20, 1-32 and 2-39. Wowhead shows 2-10, 2-19, 2-31 and 2-39. The low end looks clamped at 2 and the high end is lower. The formula matches every other negative-QualityModifier weapon (Torturing Poker, Thunderfury, Atiesh, the AQ maces). | Tooltips /285275, /285278. No APL throws a weapon, so this changes no DPS. | OPEN (to settle: an in-game tooltip of Satchel of Copper Bombs) | - | - |
| I6 | **Keen Machete 18610:** the 70790 hotfix makes it item level 10, quality 2, 6-13. Wowhead shows item level 10, quality 1, 5-10. The bin follows the client. | seed row: item level 8, q1. wago 70790 row: item level 10, q2. | Aligned (Wowhead stale) | - | - |
| I7 | **classic-db weapons (338) carry Era 1.12 damage.** 51 of them match no Forever curve: mostly caster weapons and MC/BWL/ZG/Naxx items such as Azuresong Mageblade, Staff of Dominance, Claw of Chromaggus and Wraith Blade. None has a Forever client row or a Wowhead Forever tooltip (all 404), and none is in any frozen BiS list. | scratch cdbcurve.py. https://nether.wowhead.com/forever/tooltip/item/17103 returns 404. | OPEN, for the data lane: when Forever ships these items (raid content 9 Dec), their client or hotfix rows will replace the Era values | - | - |
| I8 | Speeds: 1541 of 1541 client weapons match ItemDelay, Fiery War Axe at 3.6 included. | Same dump as I1. | Aligned | - | - |

New permanent test: `data/tests/test_simdb_weapon_client_parity.py`. It checks every client-backed weapon in the committed simdb.bin (min, max, speed) against the client tables the pipeline reads: merged ItemSparse/Item plus the ItemDamage curves. It skips unless the raw client tables have been fetched; they are never committed. With 70790's tables fetched, it fails today and lists exactly the 8 weapons in I2. That is the stale-bin signal; the next data run clears it.

## Before/after

No engine or pipeline change. No spec moves: none of the 8 stale weapons is in any frozen BiS list ($R/bis-frozen), so the round-22 raid-60 numbers are unchanged.

## Files to regenerate

- `data/builds/1.60.1.70790/simdb.bin` and `sim/internal/simdb/simdb.bin` (plus simitems.json). Pick up the 70790 hotfix pushes in the next data run (`fetch-wago-hotfixes`, then `pipeline simdb`); the refresh hold applies. Not regenerated here: the stale rows are a data change, not a code fix.

## Measure at refresh

- Earthshaker 288900 (2H mace, item level 30, epic, 87-132) has no loot source. If one appears, it is a band-30 candidate for two-handed specs.
- Re-run the parity test after the refresh. It should pass.

## For other lanes

- Data or platform: the 70790 hotfix layer on wago now has 2791 ItemSparse rows. The bin may be stale on stats as well as weapon damage (this lane checked weapons only).
- Data: `weapon_curves.resolve_weapon_damage` uses `round()` for the high end, and the client rounds half up. No current row hits a .5. Switch to `round_half_up` when someone next touches the file (it changes no current output).

## For the guide lane

None.

## Still open

- I5 (bomb satchel ranges) and I7 (Era classic-db weapons): see the table.
- The thrown-knife spread (I2): the client hotfix rows say 0.6, Wowhead still shows no spread. An in-game tooltip of Footman's Pocketknife settles it. It changes no DPS.
- Logs: last looked 5 Nov 07:10 UTC. api.foreversixty.gg/v1/reports/recent lists only the 20 Oct test report. Logs cannot settle weapon ranges anyway; tooltips do.
