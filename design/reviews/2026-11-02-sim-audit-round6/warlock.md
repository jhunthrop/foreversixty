# Sim-accuracy round 6: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons. Client 1.60.1.70707. Pin 8baffe5fc (site 21d97955f, after the nightly). Lane engine `lane/r6-warlock` (fork), site branch `wt/r6-warlock`.

Method: spec-breakdown, raid preset, band 60, 10000 iterations at seeds 1 and 50001, horde troll and alliance gnome. Client rows are the 70707 tables from wago.tools (ItemSparse, Item, ItemEffect, ItemXItemEffect, SpellEffect, SpellClassOptions, SpellMisc, SpellPower, SpellCooldowns, SpellCastTimes, SpellDuration; fetched 2026-11-01). The server push is `r5/hotfix70707/diff-vs-base.tsv` plus a fresh dump of the 70707 ItemSparse push (868 rows, read 2026-11-01). Server text is Wowhead Forever (`nether.wowhead.com/forever/tooltip/...`, read 2026-11-01). Talent searches: `sim/cmd/talent-search -preset raid -faction horde -keep <raid-keeps.json>` on the lane engine.

Logs: still no level-60 Forever log (launch is 2026-11-04).

## Assigned item 1: Torch of Light

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Torch of Light (279246) is what the data says it is. ItemSparse: item level 65, required level 60, epic, any class, 1.7 s, DamageType 6 (Arcane), stat budget healing / spell damage / Spirit, no ItemXItemEffect row. Not in the ItemSparse push. Server tooltip: 91-170 Arcane, +11 Spirit, "healing done by up to 17 and damage done by up to 6". `items/warlock.json` and simdb carry 91-170, 1.7, 17 healing, 6 spell damage, 11 Spirit. | ItemSparse 279246; https://nether.wowhead.com/forever/tooltip/item/279246 | Aligned | - |
| T2 | Obtainable at launch: an Enchanting recipe (`loot.json` source `crafted:enchanting`, no phase gate), the same source as Brilliant Wand and Greater Eternal Wand. | loot.json; raw/items/item-sources.json | Aligned | - |
| T3 | **It is still in every band-60 warlock set after the nightly (all three specs, both factions, bare and raid), and it loses to every rival wand.** Raid, band 60, seed 1 / 50001, DPS: Affliction horde Torch 573.5 / 573.9, Brilliant Wand 575.4 / 575.7, Bonecreeper Stylus 576.1 / 576.1, Ritssyn's Wand 575.3 / 575.6. Every spec and faction gives the same order, Bonecreeper Stylus +0.3 to +0.45% and Brilliant Wand +0.25 to +0.35% (table in `scratch r6/warlock/ranged.txt`). The cause is the ranker: `score.go:184` adds a wand's whole DPS (`c.DPS / referenceDPSPerPoint`) whenever the rotation has a Shoot line. The warlock lists Shoot only when nothing else can be cast. On the band-60 raid sets Shoot is cast 0 times a fight, and 0.4 times on bare band 60. So Torch's 76.8 wand DPS, against Brilliant Wand's 59.6, buys a 155-point score for damage the sim never deals. The verify pass then kept it: alliance Brilliant Wand shows "dps_delta 0, verified", where 10,000 iterations give +1.9 ± 0.4. | spec-breakdown `-gear ranged:<id>`; Shoot rows in the breakdown; `sim/cmd/leveling-bis/score.go:184`; `bis/warlock-*.json` ranged slots | CONFIRMED-BUG (ranker). Verdict: drop. sim/cmd/leveling-bis is not this lane's, so the fix is with the integrator (see Still open). | - |
| T4 | **A wand shot was Physical.** `core/wand.go` registered Shoot as `SpellSchoolPhysical`, so armor reduced every wand hit (about 40% on a bare boss, about 5% under the raid's armor debuffs), and the wand's school never counted. A Shoot-only rotation with Torch, bare: 81 average hit against the client's 130.5 midpoint. Every playable client wand is magical (ItemSparse DamageType: 25 Fire, 19 Arcane, 13 Nature, 8 Shadow, 4 Frost; the 12 physical rows are "Monster - Wand" and deprecated items, plus Wicked Lich's Finger 281263). Shot cadence and range are right: 105.3 shots in 180 s at 1.7 s, 149.4 at Brilliant Wand's 1.2 s. Fix: SimItem gains `damage_school` (24), Item and the ranged Weapon carry it, and Shoot casts in that school with no bonus coefficient. An item without the field stays physical, so nothing moves until the data lane writes it. With the field patched into a scratch simdb for the 52 client wands, a bare Torch hit averages 127 (the midpoint less the level-63 partial resist). Tests `TestShootDealsTheWandsOwnSchool` (failed: school 2, physical index, coefficient 1; then passed) and `TestShootWithNoStatedSchoolStaysPhysical`. | ItemSparse DamageType; tooltips 279246 ("Arcane Damage") and 249385 ("Fire Damage"); `sim/core/wand.go` | CONFIRMED-BUG | E ee0393589 |

## Assigned item 2: the healing-hybrid pieces after the nightly

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Spire of Hakkar and the alliance Demonology Devout pieces are gone from every warlock set. Destruction's staff is Elemental Mage Staff at raid band 60 and horde bare 60. | `bis/warlock-*.json` at 21d97955f | Aligned | - |
| H2 | Kindling Stave stays on Affliction and Demonology (bands 50 and 60) and on Destruction alliance bare 60. It is not a phantom-third pick: its tooltip states "+60 damage and healing", so it carries 60 real spell damage. It beats Elemental Mage Staff's 100 Fire and Frost for the Shadow specs on damage they deal. | items/warlock.json (wowhead source); round 5 C2 | Aligned | - |
| H3 | Uther's Strength stays only in alliance bare bands (Affliction 50, Demonology 60, Destruction 50 and 60) and as a zero-delta alternative. It now gives 6 spell damage, not 12, and it still wins on that 6 against the confirmed candidates. Alliance bare, seed 1 / 50001: Demonology 60, 325.3 / 325.3 against Mark of the Chosen 322.9 / 323.0. Destruction 60, 358.9 / 359.1 against 356.2 / 356.5. Destruction 50, 238.2 against 235.8. Affliction 50, 197.8 against 195.7. Rune of the Guard Captain, which sims about equal, is a horde quest. | spec-breakdown `-preset bare -band <b> -gear trinket2:<id>` | Aligned | - |
| H4 | Torch of Light is the one hybrid piece left everywhere, for T3's reason, not round 5's. | T3 | see T3 | - |

## Owner ruling 2 (recorded, not reopened)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Horde race picks within 1% (Orc against the published Troll on all three warlocks): the guides stay as they are. Not re-measured. | owner ruling 2, 2026-11-01 | ruling | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Classic-db items lost their single-school spell damage.** `classicdb_items._equip_stats` mapped aura 13 (damage done) only when its mask was all magic. The comment dates from before round 4 gave the planner and the sim school-power keys. 41 class items whose only damage line is "Increases damage done by Shadow (Fire, Frost, Nature, Arcane, Holy) spells and effects by up to N" scored and simmed as nothing. That is 180 simdb rows in all, and every one of them is a classic-db item. Warlock examples open at launch: Skul's Ghastly Touch (+14 Shadow), Maleki's Footwraps (+27), Dark Advisor's Pendant (+20), Anastari Heirloom (+13), Deadwalker Mantle (+13), Flaming Band (+24 Fire), Fire Striders (+29 Fire), Eye of Flame (+43 Fire). Each school bit now maps to its school power, as the client path in `pipeline/simdb/equip.py` does. Test `test_equip_stats_maps_a_single_school_damage_aura_to_its_school_power` (failed, then passed). Sources: classic-db spell 9414 (aura 13, misc 32, 13 + 1), and Wowhead Classic item 13396 ("+14 Shadow"). On a scratch simdb with the fix, single swaps on the horde raid sets read (seed 1): Affliction Dark Advisor's Pendant 576.4 and Skul's 575.9 against 573.5; Destruction Eye of Flame 678.1 against 673.4. | `data/pipeline/classicdb_items.py:737`; classic-db extract; items/warlock.json | CONFIRMED-BUG | S 2df00eafd |
| C2 | C1 does not reach the published sets. The ranker drops every client-unconfirmed item (classic-db stats with no 70707 ItemSparse row) before ranking (`leveling-bis/data.go:264`, task L1.6). None of these items is in the ItemSparse push either, so the flag is right. The policy costs a lot: on the alliance Affliction raid set, Mark of the Chosen (no stats) holds trinket2. Unconfirmed Briarwood Reed (+29, UBRS) sims +2.85% over it (588.7 / 587.9 against 572.4 / 571.5), Draconic Infused Emblem +1.8%, and Eye of the Beast and Burst of Knowledge +1.6%. Talisman of Ascendance (22678, confirmed) gives +2.1% but is not a candidate either, probably behind a content gate. C1 still reaches /sim, the planner and anyone who wears these items. | `sim/cmd/leveling-bis/data.go:264-290`; ItemSparse push dump; spec-breakdown `-gear trinket2:<id>` | OPEN (policy, for the integrator) | - |
| C3 | The Attributes_14 bit 64 that the push adds to almost every class spell (Shadow Bolt, Immolate, Corruption, Life Tap, the banes, Soul Fire, Conflagrate, Shadowburn, Siphon Life, Amplify Curse, Demon Armor and more) changes no amount, mask, cost or timing. No warlock spell or talent row changed otherwise, except Improved Drains (round 5 H1). | diff-vs-base.tsv | Aligned | - |
| C4 | Breakdown cosmetics, not sim bugs. Master Demonologist (Succubus) reads 66.7% uptime with 2 procs because the breakdown averages the warlock's aura with the demon's. Zeroing the talent drops Shadow Bolt's average hit from 1412 to 1283 (-10.0%) and Lash of Pain's from 144 to 131, so both units hold it all fight. The Imp's -25% is the prepull aura before the sacrifice. "Thrash (21919)" at 100% is the raid Windfury buff's extra-attack bookkeeping and does nothing for a caster. | spec-breakdown, with and without the talent | Aligned | - |

## Affliction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Raid search: keeping every unmodeled guide talent, 1 Improved Drains to Improved Bane of Agony gives +4.2 ± 1.4 (+0.7%), below the bar. The +1.4 to +1.5% winners drop Grim Reach, an unmodeled guide talent. The build holds, as in rounds 4 and 5. | scratch `r6/warlock/ts-out/warlock-affliction-raid.md` | Aligned (held by the adoption rule) | - |
| A2 | Every cast spell against the client: Shadow Bolt r9 (251, 1.8 a level, variance 0.108, 0.857, 3.0 s, 370), Corruption r6 (57 a tick for 6 ticks, 0.20 a tick, 2.0 s, 290), Bane of Agony r6 (46 every 2 s for 12 ticks, 0.133 a tick, ramping per the server text "dealt slowly at first"), Siphon Life r4 (41 every 3 s, 0.05 a tick, 30 s), Immolate r7 (146 + 2.3 a level at 0.20, then 52 a tick for 5 ticks at 0.13), Life Tap, Amplify Curse (3 min). Engine constants, costs, casts and cooldowns match. | SpellEffect / SpellMisc / SpellPower / SpellCooldowns 70707; tooltips 11713, 18881, 11672; corruption.go, curses.go, siphon_life.go, immolate.go, shadowbolt.go | Aligned | - |
| A3 | Talent masks against the engine: Shadow Mastery (direct on Shadow Bolt, Shadowburn and Corruption; periodic on the banes, Corruption, the drains and Siphon Life), Malevolence (every warlock Shadow spell with damage, plus Lash of Pain), Pandemic (the DoTs only), Improved Corruption (-2 s and +10% periodic), Suppression (+5% spell and melee hit), Ruin (the Destruction set), Bane (-0.5 s on Shadow Bolt, Immolate and Incinerate; -2 s Soul Fire). | SpellEffect masks resolved over SpellClassOptions; talents.go | Aligned | - |
| A4 | Improved Corruption's +10% is a separate multiplier in the engine (corruption.go:64). If the server adds percent spell mods, as Classic did, Corruption reads 0.8% high, about 0.1% of DPS. Single question, no source either way. | corruption.go:64 | OPEN (minor) | - |

## Demonology

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | Raid search: the guide build is the best (646.3 ± 1.4). | scratch `ts-out/warlock-demonology-raid.md` | Aligned | - |
| D2 | Forever swaps Demonic Sacrifice: the Imp gives +15% Shadow (Burning Shadow 18789, aura 79 misc 32), the Succubus +15% Fire. The prepull sacrifices the Imp, and Burning Shadow is up 100% under Demonic Pact. Master Demonologist (Succubus) +10% Shadow, Soul Link +3% both, Demonic Knowledge +60 spell damage to warlock and demon at 3/3: all as the tooltips state. | client 18789 / 18791; tooltips 18788, 23785, 19028, 412732; talents.go | Aligned | - |
| D3 | Decimation: 90% shorter Soul Fire cooldown, and the 40% faster cast for 10 s after a Shadow Bolt below 35%. The 6% on that Shadow Bolt is server-scripted (client base 0), and the engine takes it from the text. Soul Fire r2: 431 + 1.9 a level, coefficient 1.0, 6 s, minus Bane's 2 s. | client 440870 / 440873, 17924; soul_fire.go | Aligned | - |
| D4 | Bane of Doom averages 6952 a hit against Shadow Bolt's 1412. That is consistent with Doom's 4.0 coefficient at about 700 spell power: the raid flask's 150, Greater Arcane Elixir, Wizard Oil, Demonic Knowledge's 60 and Elixir of Shadow Power. | client 603 (1742, coefficient 4) | Aligned | - |
| D5 | Succubus: Lash of Pain r6 is 50 base at 0.429, 12 s, 160 mana. Improved Sayaad 1/3 is an ALL_EFFECTS mod, +10% on the base points only, as the engine does. The demon's crit, haste, hit and level-60 sheet stay blocked on launch logs (round 5 P1 to P4). | client 11780, 18754; succubus.go:115-160 | Aligned (and open as before) | - |

## Destruction

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | Raid search: the guide build is the best (670.8 ± 1.1). | scratch `ts-out/warlock-destruction-raid.md` | Aligned | - |
| S2 | Incinerate r3 (217 + 1.4 a level, variance 0.15, 0.714, 2.5 s, 325, +25% on an Immolated target), Conflagrate r6 (282 + 1.3, variance 0.219, 0.429, 10 s, 255), Shadowburn r6 (266 + 1.8, variance 0.109, 0.429, 15 s, 365): all match the engine. | client 1293813, 18932, 18871; incinerate.go, conflagrate.go, shadowburn.go | Aligned | - |
| S3 | Agonizing Flames 3/3: +10% direct damage on the Destruction set (Shadow Bolt, Shadowburn, Soul Fire, Searing Pain, Immolate, Conflagrate, Incinerate, Rain of Fire) and +10% on Immolate's ticks, with +10% crit on Searing Pain. The engine's additive 10% on the Destruction-flagged spells covers both halves. Aftermath 4/5 is +40% on Immolate's initial hit only. Shadow and Flame 5/5 keeps Immolate after Conflagrate (100%), and gives +10% Shadow after Conflagrate and +10% Fire after Shadowburn for 20 s. Fire and Brimstone gives +25% crit on Conflagrate. Cataclysm is -10% cost. | client 17927, 18119, 426316, 412751, 17778; talents.go:806-1340 | Aligned | - |

## Before/after (relative, raid preset, band 60, 10000 iterations; pin 8baffe5fc against E ee0393589 + S 2df00eafd)

| Spec | Alliance seed 1 | Alliance seed 50001 | Horde seed 1 | Horde seed 50001 |
|---|---|---|---|---|
| Affliction | 572.4 to 572.4 (0.0%) | 0.0% | 573.5 to 573.5 (0.0%) | 0.0% |
| Demonology | 644.3 to 644.3 (0.0%) | 0.0% | 646.8 to 646.8 (0.0%) | 0.0% |
| Destruction | 670.2 to 670.2 (0.0%) | 0.0% | 673.4 to 673.4 (0.0%) | 0.0% |

The fork suite passes with `-tags=with_db`, except `sim/web`, which fails setup on the pin as well.

Neither fix moves today's published numbers. Each waits on data the nightly regenerates: E ee0393589 needs `damage_school` written for wands, and S 2df00eafd needs data.yml's simdb and items regeneration.

Measured on scratch simdbs carrying both fixes (horde, band 60, seed 1):
- Warlock raid: unchanged to the decimal. Bare: Affliction +0.03%, the others unchanged. Warlocks rarely Shoot.
- Outside this group, the wand school moves only bare bands: Frost 260.3 to 268.7 (+3.2%), Fire 341.4 to 344.8 (+1.0%), Arcane 313.1 to 316.3 (+1.0%), Shadow 378.2 to 379.6 (+0.4%). Their raid numbers are unchanged. Lower bands, where casters wand more, will move more.
- C1 moves no published set (C2), but it changes /sim for anyone wearing one of the 180 rows, in every caster and healer class.

## Still open

- **For the integrator (ranker, sim/cmd/leveling-bis):**
  - T3: weight a wand's DPS by the share of the fight the band's own reference run spends shooting, not 100%. Shoot casts 0 a fight on the warlock raid sets. Then re-rank; Bonecreeper Stylus is unconfirmed, so Brilliant Wand should take the slot.
  - The verify pass scored Brilliant Wand "0, verified" against a +0.33% gap.
  - C2: the client-unconfirmed exclusion and Talisman of Ascendance's gate, worth about 2 to 3% on the trinket slots.
- **For the integrator (data, after the pin moves the vendored proto):** write `SimItem.damage_school` for wands in `pipeline/simdb/items.py`.
  - Client rows: ItemSparse DamageType 1 Holy, 2 Fire, 3 Nature, 4 Frost, 5 Shadow, 6 Arcane, mapped to proto SpellSchool 4, 2, 5, 3, 6, 1.
  - Classic-db rows: `dmg_type1`.
  - Wands only. The engine reads the field for every weapon, and no melee row was checked.
- A4: whether percent spell mods add or multiply on the server (about 0.1%).
- Carried: demon crit, haste, hit and the level-60 sheets; DoT spell power read per tick (forever-bugs #398); Demonic Brand unmodeled; Siphon Life crit held for the owner; Affliction's +0.7% swap held.
- sim/request TestRotationLadder goldens are not regenerated, as the brief requires. Neither fix moves them today.

Sources (read 2026-11-01): https://wago.tools/db2/ItemSparse/csv?build=1.60.1.70707 (and Item, ItemEffect, ItemXItemEffect, SpellEffect, SpellClassOptions, SpellName, SpellMisc, SpellPower, SpellCooldowns, SpellCastTimes, SpellDuration); https://wago.tools/hotfixes (70707 ItemSparse push); https://nether.wowhead.com/forever/tooltip/item/279246 and /249385; https://nether.wowhead.com/forever/tooltip/spell/11713 (and 18881, 11672, 412732, 23785, 18788, 19028, 1293817, 426316, 18871); https://nether.wowhead.com/classic/tooltip/item/13396 and /22678.
