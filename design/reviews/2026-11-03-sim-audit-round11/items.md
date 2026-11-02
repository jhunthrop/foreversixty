# Sim audit round 11: items lane

Scope:
- Owner ruling 4: the 165 item ids gone from round 10's local data run.
- A fresh re-audit of every item on the 336 published band sets (`data/builds/1.60.1.70707/bis/*.json`, 719 distinct items). It covers proc triggers and rates, use effects and shared cooldowns, set bonuses at every threshold a band reaches, the hotfix push, round 10's fixes, and the simdb stat mapping.

Method:
- Client tables: the 70707 raw tables with the push layered in (`push-check`: moved=false, 2026-11-02 22:10 UTC). SpellCategories, SpellProcsPerMinute, SpellClassOptions, SpellEquippedItems and SpellCategory were fetched fresh from wago.tools for build 70707.
- Item and ItemSparse: fetched fresh from wago.tools, 2026-11-02 22:09 UTC, for 70707 and 70753, both shipped and `useHotfixes=1`. The item layer was rebuilt with the pipeline's own `fetch_wago_hotfixes`.
- Server text: nether.wowhead.com/forever/tooltip, read 2026-11-02. A tooltip plus the client row counts as one source.
- Scratch: scripts and dumps are in `scratchpad/r11/items/` (`scripts/`, `census165.json`, `bisitems.json`, `meas.tsv`).

## 1. The 165 ids (owner ruling 4)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **The 165 are not gone from the client. Round 10's local run skipped a fetch.** All 165 are hotfix-only items: none is in the shipped 70707 ItemSparse, and every one is in the item layer that `fetch-wago-hotfixes` writes to `raw/wago-hotfixes/`. 126 come from 70532's recorded push and 39 from 70707's. Round 10's ranker lane ran `fetch`, then `normalize` (no `wago-hotfixes/` in `scratchpad/r10/ranker/root/.../raw/`, and no `useHotfixes` request in its `datarun*.log`). `fetch` never fetched the item layer, so `normalize` lost these 165 and the shrink gate fired. A second, quieter loss came with it: 404 more hotfix-only items fell back to wowhead or classic-db stats. With the layer present, today's merge loses none: 24,213 client ids against the committed 24,210, 0 gone, and 3 move from the supplement to the client. | `scripts/gone.py`: 0 gone with the layer, 569 without. The 165 are exactly round 10's committed-minus-run set. Shipped Item and ItemSparse are byte-identical to the 2026-10-30 and round-10 copies, so the export is not truncated. | CONFIRMED-BUG (fetch) | old (S 5d7a422cb, round 5, added `fetch-wago-hotfixes` as a step of its own) | S 3484405dd |
| G2 | **Fix:** `python -m pipeline fetch` now fetches the item layer right after it layers the push, so a fetched build is never left without the item pushes. The data workflow already runs `fetch-wago-hotfixes` (`data.yml:213`), so the held refresh would not have lost them. | `test_fetch_layers_the_push_and_the_item_pushes_after_the_download` failed first (no call), then passed. The 28 wago, hotfix and push tests pass. | fix | - | S 3484405dd |
| G3 | **The refresh must NOT pass `--allow-shrink` for these ids.** Round 10's checklist line rests on G1. Passing it would switch off the gate that caught G1. | as G1 | Ruling outcome | - | - |
| G4 | **Use.** No published BiS set, guide, curated file or the addon names any of the 165. loot.json names 2: Precision Arrow 9399 (Baelog) and Tabard of the Scarlet Crusade 23192 (Scarlet Trainee). Both have live Wowhead Forever pages. The nightly `simitems.json` lists 105. | appendix | Aligned | - | - |
| G5 | **70753 (owner-approved, listed on wago 2026-11-02 21:35 UTC).** The 39 rows from 70707's push are shipped in 70753's ItemSparse. The other 126 are still in no shipped table. A 70753 run would keep them through the same layer, because the layer reads 70532's push for every later build on the line. Verdict unchanged: no drop. | `wago-70753/` matches the integrator's copy in `c70753/` byte for byte | Aligned | - | - |
| G6 | **Possible stale content.** 98 of the 126 rows from the 70532 push have no Wowhead Forever page: 96 item-level-30, required-level-1 copies of Molten Core gear, plus Gnoll Greataxe 288358 and Quel'Serrar 289302. Wago's 70707 record is partial, so the client tables cannot prove whether the server still sends these rows. None has a loot source, so the ranker cannot pick one. | tooltip 404s, 2026-11-02 | OPEN (data, not sim) | - | - |

The full list, with every column, is in [items-165.md](items-165.md).

## 2. Set bonuses reached by published bands

The census counts every set worn with 2 or more pieces at each band (`sets.json` pieces × `bis/*.json`). The reached sets are:
- Gladiator 2
- Deathbone 2 and 3
- Stormshroud 3
- Devilsaur 2
- Ironfeather 2
- Defias Leather 4 (22 bands)
- Scarlet Crusade 2
- Magister's 3, 5 and 6
- Devout 5, 6 and 7
- Shadowcraft 2, 3 and 5
- Elements 3
- Lightforge 2
- Valor 2
- the PvP rank sets at 2 and 3

Round 10's T4 listed only Ironfeather and Devilsaur as reached; that row was about crit bonuses only. Every flat bonus is "declared, matches" in the fork's conformance golden.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Set procs rolled on every melee hit, whatever the client's flags said.** The client says which attacks a proc fires on (SpellAuraOptions.ProcTypeMask): 0x4 is a melee autoattack, 0x10 a melee special. The engine fired every hand-written melee set proc on any landed melee hit. That broke these four, which the client limits to melee autoattacks (0x4):<br>- Devious Strike (Defias Leather 4P, 1292028, 5%)<br>- Rogue Armor Energize (Shadowcraft 5P and Darkmantle 4P, 27787)<br>- Valor's and Heroism's rage (450587)<br>- Bloodmail (1299740)<br>So Claw, Shred, Sinister Strike, Crusader Strike and Heroic Strike each rolled a 5% bleed or an energize they never get in game. Lightning, Revitalize, Spider's Kiss and Firebolt are 0x14 (autoattacks and specials) and are unchanged. **Fix:** the set rows carry ProcTypeMask, and each melee bonus procs on the attacks its flags name. A melee bonus whose flags name no autoattack stops the engine at init. | SpellAuraOptions rows 245820, 136944, 236876, 246449 (mask 4) and 130048, 134496, 127963, 122953 (mask 20); single source, the client. `TestMeleeSetBonusesProcOnTheClientsAttacks`: a build failure first; with the old rule restored, 4 of its 8 cases fail. It passes now. The fork's core, common, conformance, rogue, druid and warrior suites pass with no golden moved. | CONFIRMED-BUG | old (set models E e0a523557, 2026-10-08; rounds 8 and 9 checked these bonuses' chances but not their triggers) | E c35ac4a2f, S 8d5e502f1 |
| S2 | Every other reached bonus reads its amount, chance and recovery from the client: Deathbone Surge (15%, 60 sec, when struck), Lightning (10%, 20 Nature), Revitalize (4%, 30 energy), the 5% 200-mana procs on spellcast (mask 0x14000), Roar of the Crowd (4P, not reached) and the PvP flat bonuses. The bonuses the sim leaves out (resist fear, freeze, the "when struck" shields and snare removals) are declared with a reason. | `sets.golden.md`; client rows | Aligned | - | - |
| S3 | Rogue Armor Energize states ProcChance 100 with no PPM row, so it keeps the unstated one-a-minute rate (round 10's P1, still open). S1 changes only the attacks it rolls on, not the rate. | rows above | OPEN (rate, P1) | - | - |

## 3. Item effects on published gear

These are all 44 published items with a client effect that is not a flat equip stat.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **Thunderbrew's Boot Flask (744) dealt its damage as damage-over-time ticks.** Its spell, 12257, fires a separate spell once a second (aura 23): 12278, a direct 75 Fire hit (DefenseType 1, SchoolMask 4) that rolls spell hit and crit on its own. The engine dealt the 75 as a dot tick, which rolls neither. Each second is now a 12278 cast. Published for Frost at 50 bare (Alliance). It changes nothing within error. | 12257 / 12278 SpellEffect, SpellCategories, SpellMisc; server text "Deals 75 Fire damage every 1 sec for 5 sec". `TestThunderbrewsBootFlaskTicksAreFireSpellsThatCrit` failed (no 12278), then passed. | CONFIRMED-BUG | old (E febc3c25a, round 8) | E b03f0bb57 |
| E2 | The engine-modelled procs and uses agree with the client's amounts, durations, cooldowns and categories:<br>- Manual Crowd Pummeler: 13494, +50% for 30 sec, 3 min.<br>- Hand of Edward the Odd: 18803, 4 sec, one charge.<br>- Arcanite Champion: +120 Strength for 30 sec. Its self-heal is 360 flat; the engine heals 270 to 450, which changes no damage number.<br>- Mark of the Chosen: 2%, 120 sec recovery, +21 to all stats for 60 sec.<br>- Spire of Hakkar: 28 every 2 sec for 8 sec.<br>- Aegis of the Scarlet Commander: 14, 5 sec.<br>- Toxic Revenger: 14 every 3 sec for 12 sec.<br>- Hearts of the Mountain: 5 min, category 1141 15 sec.<br>- Hauberks: 1 hour, category 1141.<br>- Ankh of Life: 720 ±20%, 5 min.<br>- Robe of the Archmage: 500 ±50%, 5 min.<br>- Talons of Wrath: 50%, 35 mana.<br>- Kindling Stave and the weapon procs in `weapon_proc_rolls.go`. | client rows; `scripts/sp.py` | Aligned | - | - |
| E3 | Fifteen published effects have no damage, threat or mana in the sim, as round 8 found:<br>- Self-heals and absorbs: Uther's Light 10368 (when struck), Lifeforce 1300356, Julie's Blessing 8348, Savior's Sacrifice 18826, Bone Armor 18828.<br>- Other effects: Guarded 1300364 (aggro radius), Slow Fall, Reckless Charge, Robe of the Void's pet heal, the stun on Hammer of the Titans, Arena Grand Master's absorb, and Archmagister's Wisdom 1292997 (aura 447, an experience buff).<br>- The Satchels of Bombs: they proc on ranged autoattacks (mask 0x40), and the warrior tanks that wear them never throw. | client rows | Aligned | - | - |
| E4 | **Hotfix push.** The 70707 push changes 4 ItemEffect rows, none on a published item. No SpellEffect, SpellAuraOptions, SpellMisc or SpellCooldowns row of a published item's spells is pushed. ItemSet, ItemSetSpell and ItemXItemEffect are unchanged. | shipped/ versus layered raw | Aligned | - | - |
| E5 | **Round 10's 30 fixes, rechecked against the client:** all match.<br>- Defense types: Gloves of Holy Might 1325229 and Thornspike 1293183 are 3.<br>- Damage: Blaze 20 every 3 sec, Gutwrencher 11 every 1 sec, Barman Shanker 14, Brain Damage 250 ±40% Shadow.<br>- Razor Gauntlets 7 Nature; the Emblem 35 for 10 sec.<br>- The fourteen R6 rows.<br>- Mindtap 125 every 1 sec; Sanctified Orb 340 ±12.5%. | client rows | Aligned | - | - |
| E6 | **Robe of the Archmage and the mana gems.** All share spell category 30, 60 sec (SpellCooldowns 18385 and 10058). The gems also carry item category 1153, 120 sec. The engine runs one "conjured" timer and sets it to 2 min on a gem. After the Robe it sets it to 1 min, but the Robe itself never waits on it. So a Robe can follow a gem inside 60 sec, which the client forbids. This touches only mana (Arcane and Fire at 60) and is small. The fix needs a second timer that `mage/mana_gem.go` also sets. | rows above | OPEN (routed: casters) | - | - |

## 4. Item stats in simdb

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | I compared every published item's simdb stats with Wowhead Forever's payload (`raw/wowhead/items.json`, 691 of 719 items). 142 differ, nearly all by naming. The engine counts attack power as both melee and ranged (aura 99 + 124). Forever's stat 45 counts as spell power, and the 18 + 6 split is the payload's own (Earthen Silk Slippers: "healing 18, damage 24"). The payload's "splheal" is healing minus damage (Holy Shroud: 33 healing, 11 damage, as the tooltip and the engine both say). The negative stats are right (Ring of Scorn −3 Spirit, Tarnished Locket −3 Strength). Rune of the Guard Captain's tripled terrain AP is ignored, per the ruling. | `scripts/statcmp.py`; tooltips | Aligned | - | - |
| D2 | **The committed simdb predates two round-8 data fixes.** S 946dde232 (2026-11-02 02:54 UTC) lands after the last simdb commit (c0bd0bd3d, 23:47 UTC the day before). So the committed simdb lacks stat-50 bonus armor on 337 items. Published tank picks among them: Sergeant's Cloak and Cape at 30, and the Warden's and Defender's Leather at bear 20 to 50. It also still rounds Wowhead percentages: Blackcrow 1.0 hit where it should be 0.9 (22 hunter bands), Savage Gladiator Chain 2.0 crit where it should be 1.5. Today's code gives 0.9, 1.5 and the bonus armor; the 95 statmap and simdb-item tests pass. This lands with the held refresh. No code change. | `wh.sim_percent_stats`; ItemSparse StatModifier_bonusStat 50 rows; tooltips (Sergeant's Cloak "+40 Armor") | Held output (lands at refresh) | - | - |

## 5. Follow-ups from the crit and casters lanes

This work builds on the crit lane's simdb commits (S eff02da4c, 3ff4491a4 and 7bac21279, cherry-picked as 1a37820a2, 51c012d18 and 17ed227eb) and on its engine scope commit (E e183994e8, cherry-picked as 8428600c5). The SpellEquippedItems rows are the 70707 table fetched for this lane. They agree with the crit lane's fetch on every row read here.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Classic-db items whose client links state Forever's own stats.** These items have no ItemSparse row, so simdb takes their stats from classic-db's 1.12 rows. The crit lane's rule (where the client links any crit, the client's crit replaces Era's) now covers every stat the links state, key by key; crit and hit each replace as a group. On 70707 it changes eight items besides the crit lane's two:<br>- Royal Seal of Eldre'Thalas 18467 and 18471: 18 spell damage and healing (1318479, 1318476), Era 23.<br>- 18469 and 18470: +16 Spirit (aura 29).<br>- 18472: 9 mp5.<br>- Counterattack Lodestone 18537: 30 attack power (1318938), Era 22. Era's 1% parry, which the client does not restate, stays.<br>- Briarwood Reed 12930: 30, Era 29.<br>- Trueaim Gauntlets 13255: Era's every-table 1% hit goes (F2).<br>Auras on these links that the sim does not read (procs, attack power against creature types) are skipped, not classified. Of the 1,351 supplement items, 80 have client links; the other 70 either state the same amounts or state nothing the sim reads (Hand of Justice links only its proc, Omokk's Girth Restrainer a threat reduction). Those keep Era's stats, because a hotfix item's client links can be partial. | ItemXItemEffect → ItemEffect → SpellEffect 70707. Wowhead Forever has no page for any of them, so the client row is the only source. `test_the_clients_own_stat_links_replace_classic_dbs_era_stats` failed first (no `item_client_stats`), then passed; 79 simdb tests pass. | CONFIRMED-BUG (no published set; all ten have a loot source, so they are ranker candidates) | old (S 8b5358dfe, the classic-db supplement, 2026-09-30) | S 34af35fb9 |
| F2 | **Ranged-only hit on four classic-db items.** The client links these to "Increased Hit Chance - Ranged" (aura 54), restricted to ranged weapons (SpellEquippedItems mask 262156):<br>- Trueaim Gauntlets 13255 → 1301127, 3%.<br>- Infernal Trickster Leggings 17754, Marksman Bands 18296 and Stoneshatter 18388 → 1294769, 1%.<br>simdb now carries none of it, and the engine gives ranged attacks the hit, as it does for the scopes. | `TestForeverRangedHitItemsReachRangedAttacksOnly` failed first (Auto Shot +0), then passed. The fork's hunter, common, core and conformance suites pass. | CONFIRMED-BUG (no published set) | old (S 8b5358dfe) | E 9a4835260, S 34af35fb9 |
| F3 | **Blackfathom Sharpening Stone (enchant 7098, item 211845) gave its hit to spells.** "Hit +2%" is 430391 "Hit Weapon +2" (aura 54, 2), restricted to melee weapons (SpellEquippedItems class 2, mask 42483). simdb gave it as the one hit stat, which reaches spells. Aura 54 is melee and ranged hit, as a stone's aura 52 is its crit (round 9: stones are melee). The row now carries none of it, and the engine adds 2% to every weapon attack and nothing to spells. Elemental Sharpening Stone 2506 (aura 52, same mask) keeps its melee crit in the row. Across 70707 this rule changes 7098 only. No preset, published set, `simconsumes.json` or `enchants.json` carries the stone, so no request reaches it today. The engine change is an enchant effect in `sim/common/enchant_effects.go`; it uses the existing per-spell `BonusHitRating` and does not touch the stat plumbing. | `TestBlackfathomSharpeningStoneHitReachesWeaponAttacksOnly` (E) failed first (swing +0), then passed. `test_a_melee_weapon_stones_hit_is_left_to_the_engine_and_its_crit_stays` (S) failed first (import), then passed. | CONFIRMED-BUG (unreachable) | old (S ae2accbec, enchant equip spells as stats, 2026-09-18) | E 897b1c910, S f66fd9684 |
| F4 | **Shard of the Defiler (17142)** is a real 70707 ItemSparse row: an item-level-70 legendary dagger with 15% crit, 15% melee haste and 15% spell haste. Wowhead Forever has a page for it. It has no loot source, so the ranker leaves it in NoSource (`leveling-bis/band.go:550`) and it can never be a pick, and no published set carries it. It does sit in `items/<class>.json` for eight classes, where /sim's gear list can offer it. That is a catalogue question, not a sim one. | ItemSparse; loot.json (0 hits); tooltip | Aligned (candidates) / routed (catalogue) | - | - |
| F5 | **Elemental Mage Staff (944), for the casters lane.** It has no shipped or committed-dump row; wago's recorded pushes carry it (`raw/wago-hotfixes/`). Its stats, from StatModifier_bonusStat 5, 47, 51, 52, 85 and 87, are 26 Intellect, 8 spell penetration, 10 Fire and 10 Frost resistance, and 100 Fire and 100 Frost damage. Weapon damage is 109 to 164 at 3.20. It has no ItemXItemEffect link. simdb has exactly these: Intellect 26, FirePower 100, FrostPower 100, SpellPenetration 8, the two resistances 10, and 109 to 164 at 3.2. The server text says the same. It is published for Fire and Frost at 60. | merged ItemSparse; simdb; nether.wowhead.com/forever/tooltip/item/944 (2026-11-02) | Aligned | - | - |

## Before/after

10,000 iterations. Before: `sb-before` (engine cfcb76e80, site 8f047a62f). After: `lane/r11-items` b03f0bb57 with `wt/r11-items`. Only S1 and E1 move a number.

| spec | band | factions | seed 1 / 50001 (alliance; horde) | absolute, seed 1 alliance |
|---|---|---|---|---|
| Feral cat | 20 bare | both | −1.72 / −1.93; −1.96 / −1.75 | 46.6 → 45.8 |
| Feral cat | 30 bare | both | −0.94 / −1.06; −0.83 / −1.07 | 85.2 → 84.4 |
| Feral cat (Shadowcraft 5P) | 60 bare | both | −1.54 / −1.50; −1.56 / −1.56 | 272.9 → 268.7 |
| Feral cat (Shadowcraft 5P) | 60 raid | both | −1.47 / −1.35; −1.54 / −1.42 | 674.7 → 664.8 |
| Survival | 20 bare | both | −0.86 / −1.08; −0.85 / −0.85 | 46.4 → 46.0 |
| Retribution | 20 bare | both | −3.34 / −3.11; −3.12 / −2.86 | 41.9 → 40.5 |
| Assassination | 20 / 30 bare | both | −1.57 / −1.57 and −1.22 / −1.22; horde the same within 0.2 | 38.1 → 37.5 (20) |
| Combat | 20 bare / 30 bare horde | 20 both; 30 horde | −1.22 / −1.46; 30 horde −0.77 / −1.07 | 40.9 → 40.4 (20) |
| Subtlety | 20 bare / 30 bare alliance | 20 both; 30 alliance | −2.10 / −2.10; 30 alliance −1.16 / −1.32 | 38.1 → 37.3 (20) |
| Arms | 20 bare | both | −1.13 / −1.50; −1.12 / −1.49 | 26.6 → 26.3 |
| Fury | 20 bare | both | −0.58 / −0.58; −0.57 / −0.29 | 34.7 → 34.5 |
| Frost | 50 bare | alliance | 0.00 / −0.04 | 249.2 → 249.2 |
| Combat 30 alliance, Subtlety 30 horde (no 4P worn) | | | 0.00 / 0.00 | identical |

Section 5 moves no published number. No published set, preset or consumable list carries F1 to F3, and simdb.bin is not regenerated. F1 and F2 change ranker candidates at the refresh.

## For other lanes

- **Integrator, section 5.** E 8428600c5 and S 1a37820a2, 51c012d18 and 17ed227eb are cherry-picks of the crit lane's commits. Merge the crit lane's branches as they are; the cherry-picks collapse. S 34af35fb9 renames `item_crit_bonuses` to `item_client_stats`, and the crit lane's two tests follow the new type. F1 and F2 also need a re-rank at the refresh: Royal Seals for healers and casters at 60, Counterattack Lodestone for tanks, Trueaim and Marksman Bands for hunters.
- **Platform or data lane (F4).** Consider leaving unsourced legendary placeholders (Shard of the Defiler 17142) out of `items/<class>.json`, or flagging them in /sim's gear list.

- **Integrator, pin order.** E c35ac4a2f needs S 8d5e502f1 (`tools/gen_set_rows.py` emits ProcTypeMask). If the engine follow regenerates `client_sets_gen.go` with the old generator, every ProcTypeMask reads 0, and `meleeKindOf` stops the engine at init. Land the site commit first or with the pin.
- **Integrator, refresh checklist.** Drop "the data run needs `--allow-shrink` (165 ids)". The full workflow keeps them (G1 to G3). After S 3484405dd, `fetch` fetches the item layer itself, so `data.yml`'s separate `fetch-wago-hotfixes` step fetches it a second time. That is harmless; the platform lane can drop the step once the hold lifts. The refresh also brings D2 (bonus armor, Blackcrow, Savage Gladiator Chain), which can move tank and hunter picks.
- **Ranker lane.** S1 changes what Defias Leather 4P, Shadowcraft 5P, Darkmantle 4P, Valor 5P, Heroism 4P and Bloodmail 3P are worth to every melee candidate. Re-rank the bands in the table above at the pin. G6's 98 unsourced items sit in the candidate list without a source.
- **Casters lane (E6).** Robe of the Archmage and Sanctified Orb need the 60 sec category 30 that the mana gems also set. Gems are 1153 (120 sec) plus 30 (60 sec); today's single conjured timer conflates the two.
- **Crit lane.** No crit-plumbing finding. E1 is an item file change: 12278 rolls spell crit through the normal spell table.
- **Held-output test (not mine, seen while running the simdb tests).** `data/tests/test_loot_build.py::test_simbuffs_names_every_id_the_engine_lets_a_request_send` fails on origin/main against the committed `simbuffs.json`, which waits for the refresh.

## Still open

1. P1 (round 10): the one-a-minute rate, now covering Rogue Armor Energize, Valor and Bloodmail (S3).
2. G6: whether the server still sends 70532's 126 pushed items. It needs a login or wago's next full push record.
3. E6: Robe of the Archmage and the mana gem category.
4. Round 9 D4: Toughness on bonus armor. D2 makes it matter once the refresh adds stat-50 armor.
5. Round 10 R7 and N1 (Pendulum of Doom, Fang of the Crystal Spider, Icy Weapon, Crusader's self-heal) are unchanged.

Sources: wago.tools db2 1.60.1.70707 and 1.60.1.70753 (Item, ItemSparse, with and without useHotfixes; SpellCategories, SpellProcsPerMinute, SpellClassOptions, SpellEquippedItems, SpellCategory) and /api/builds, fetched 2026-11-02 22:09 UTC. The build's layered raw tables. nether.wowhead.com/forever/tooltip/item and /spell, read 2026-11-02.
