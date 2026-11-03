# Round 12 healers: client-row sweep of Restoration druid and Restoration shaman

Read only. Engine: `r12/healers/engine/sim` (paths below are relative to it). Client: `r11/c70753/builds/1.60.1.70753/raw` (push layered). Scripts: `sweep/ds/` (sp.py summarises a spell, tal2.py adds the trait curves, diff.py diffs 70707+push against 70753). No sims were run.

**What changed since round 10.** I diffed every SpellClassSet 7 and 11 row (SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellLevels, SpellAuraOptions, SpellClassOptions), plus every listed talent spell, between r9/casters/raw (70707 + push) and 70753. The only changes are feral ones: Poised 1328827 is new, and Ferocious Bite's variance changed in its 11th digit. TraitDefinition, TraitDefinitionEffectPoints, TraitNode, TraitEdge and ItemSetSpell did not change. CurvePoint changed in 5 rows, none of them a healer curve. The hotfix push carries no spell or trait rows. So every row below was re-read against 70753, and the gaps found are ones the earlier rounds missed. None is a regression.

Build strings were decoded by client (tier, column), which is the same order as the engine's proto field numbers (`core.FillTalentsProto`; checked against druid.proto and shaman.proto).
- Druid at 60, as published: Balance has Improved Wrath 4, Genesis 5, Moonglow 3, Improved Moonfire 2, Nature's Majesty 2 and Nature's Splendor 1. Restoration has Nature's Focus 5, Naturalist 5, Reflection 3, Gift of Nature 5, Bloom 1, Tranquil Spirit 5, Improved Rejuvenation 3, Swiftmend 1, Nature's Swiftness 1, Living Spirit 3 and Improved Tranquility 2.
- Shaman at 60: Thundering Strikes 5, Ancestral Knowledge 5, Improved Healing Wave 5, Tidal Focus 5, Mindfulness 3, Ancestral Healing 3, Healing Focus 3, Water Shield 1, Tidal Mastery 5, Restorative Totems 5, Mana Tide 1, Healing Way 3, Nature's Swiftness 1, Purification 5 and Riptide 1.

## 1. Rows checked

"=" means the engine value equals the client value. `gen` means the engine reads constants_auto_gen.go (70753).

### Restoration druid

| spell / talent | id(s) | field | engine (file:line) | client (table, row) | match |
|---|---|---|---|---|---|
| Healing Touch r1-r11 | 5185, 5186, 5187, 5188, 5189, 6778, 8903, 9758, 9888, 9889, 25297 | bp, variance, per-level, SpellLevel, MaxLevel, coef, mana, cast (11 x 8 fields) | druid/healing_tables.go:28-95 | SpellEffect eff 10 (e.g. 683467: 1920 / 0.16505 / 6.6 / 56 to 61 / 1.0), SpellPower, SpellMisc | yes, all 88 |
| Healing Touch | 9889 | GCD, no cooldown | healing.go:82 GCDDefault | SpellCooldowns gcd 1500 | yes |
| Regrowth r1-r9 | 8936-9858 | direct bp/var/ppl/levels/coef, tick bp/coef, mana, cast (9 x 10) | healing_tables.go:98-162 | SpellEffect eff 10 + eff 6/aura 8 (688380/688381) | yes, all |
| Regrowth | all ranks | period, duration, ticks | regrowth.go:10-11: 7 x 3 s | per 3000, dur 21000 | yes |
| Rejuvenation r1-r11 | 774-25299 | tick bp, coef 0.2, levels, mana (11 x 5) | healing_tables.go:165-221 | SpellEffect aura 8 (679789 bp 8 ... 681228 bp 32) | yes, all |
| Rejuvenation | all ranks | period, duration, ticks; no tick on apply | rejuvenation.go:10-11: 4 x 3 s | per 3000, dur 12000, no Attr5 tick-on-apply | yes |
| Rejuvenation r1 (bands 50-59, rank-choices.json) | 774 | learn level 4, bp 8, coef 0.2, mana 25 | healing_tables.go:167-169 | SpellLevels 4/9, SLA 573 | yes |
| Tranquility r1-r4 | 740, 8918, 9862, 9863 | tick bp/ppl/levels/coef, mana | healing_tables.go:224-245 | 687989: 285, +1 a level, 60 to 66, 0.067; mana 925 | yes |
| Tranquility | 9863 | ticks, period, cooldown, channel, target | tranquility.go:10-12 (5 x 2 s, 5 min), :26 channelled, :30 party | dur 10000, per 2000, CategoryRecovery 300000, A1 0x40, eff 35 party | yes |
| Wild Growth r1-r3 (not cast) | 408120, 1238214, 1238215 | bp/ppl/levels/coef, mana, cast, cooldown, ticks | healing_tables.go:248-266, wild_growth.go:10-22 | 3000 cast, 6000 cooldown, 7000/1000, coef 0.033 | yes |
| Swiftmend | 18562 | cost, cooldown, GCD, crit | swiftmend.go:10-12, 52-55 | PowerCostPct 20, Recovery 15000, gcd 1500, no CANT_CRIT | yes |
| Swiftmend | 18562 | reads Regrowth first, does not consume | swiftmend.go:77-88 | settled (round 10 D1) | settled |
| Swiftmend + Nature's Splendor | 18562 | ticks it pays out (5 Rejuvenation / 9 Regrowth) | swiftmend.go:68 `NumberOfTicks` | client silent; no logged druid has Splendor | OPEN (see 2.4) |
| Bloom | 1328621 / 1328622 | 100% / 50% of Spirit, no class mask, no procs | bloom.go:30-31, 43 | e0 aura 4 bp 100, e1 bp 50; 1328622 mask 0, A3 0x10000 | yes (crit is round 11 D3, still open) |
| Innervate | 29166 | cost, cooldown, duration, regen | innervate.go:53; core/buffs.go:1125-1126, 1180-1181 | PowerCostPct 5, Recovery 360000, dur 20000, aura 110 bp 400, aura 134 bp 100; learned at 40 | yes |
| Nature's Swiftness | 17116 | cooldown, GCD, charges, mask | talents.go:177-235; druid.go:112 | Recovery 180000, gcd 0, 1 charge, mask 0x1000261 (Wrath, HT, Regrowth; Wild Growth not in it) | yes |
| Omen of Clarity / Clearcasting | 16864 / 16870 | learn level 20, ICD, duration, charges, mask | talents.go:333-335, 370, 375-379 | 16864 ICD 10000, ptm 0x14014, level 20; 16870 dur 15000, 1 charge, mask includes heals and Swiftmend, not Innervate (m1 0x1000) | yes (the 4% rate is fitted, not a client value) |
| Gift of Nature | 17104 | 2%/rank, mask | talents_restoration.go:20, 44 | curves 85390/97152 2..10; e0 0xe0 + m1 0x2, e1 DOT 0x50 + Wild Growth | yes (Tranquility left out: settled) |
| Improved Rejuvenation | 17111 | 5%/rank on Rejuvenation | talents_restoration.go:21, 73 | curve 85403 5/10/15, mask 0x10 | yes (stacking: round 10 D4, open) |
| Tranquil Spirit | 24968 | -2%/rank cost, HT + Tranquility | talents_restoration.go:22, 59 | curve 85404, mask 0xa0 | yes |
| Improved Tranquility | 17123 | -30%/rank cooldown | talents_restoration.go:24, 90 | e1 misc 11, curve 85408 -30/-60, mask 0x80 | yes |
| Improved Regrowth (not taken) | 17074 | +10%/rank crit | talents_restoration.go:23 | curve 85410 10..50, mask 0x40 | yes |
| **Reflection** | 17106 | regen while casting | **talents_restoration.go:112, 118: rank/3 x 0.5 = 16.67 / 33.33 / 50%** | **curve 85365: 17 / 33 / 50** | **no at ranks 1-2 (see 2.2)** |
| Naturalist | 17069 | -0.1 s/rank on HT + Wild Growth; +2%/rank damage only | talents.go:760-761, 777-781 | curve 85364 -100..-500, mask 0x20 + m1 0x4000000; e1 aura 79 curve 85363 2..10 | yes |
| Nature's Focus | 17063 | 14%/rank pushback, mask | talents.go:800, 807-808 | curve 85362 14..70; mask 0x104002e5 (HT, Regrowth, Tranquility, Wrath, Starfire...) | yes |
| Living Spirit | 1309631 | 5%/rank Spirit | talents.go:792 | curve 85407 5/10/15, aura 137 misc_1 16 | yes |
| Heart of the Wild (guide build) | 17003 | 2%/rank Intellect | feral_talents.go:16, 26 | curve 85361 2..10, misc_1 8 | yes |
| Genesis | 1223081 | 2%/rank periodic; mask excludes Tranquility | talents.go:622, 636-641 | 1210794 DOT op, curve 85325; mask 0x220252/0x4000100 (no 0x80); e1 on Swiftmend | yes |
| Nature's Majesty | 1223082 | 2%/rank, every table | talents.go:599 | aura 290, curve 85326 2/4 | yes |
| Nature's Splendor | 1223083 | +1 tick on Rejuvenation, +2 on Regrowth | talents.go:107, 657-666 | e0 +3000 on 0x12, e1 +6000 on 0x40 (27 s and 15 s) | yes |
| Improved Wrath, Moonglow, Improved Moonfire | 16814, 16845, 16821 | no heal in mask | (damage only) | masks 0x1 / 0x700307 / 0x2 | yes |
| Nature's Grace (not taken) | 16880 / 16886 | heal-crit proc, GCD cut on heals | talents.go:136-168 (now procs on OnHealDealt) | ptm 0x14000; 16886 mask 0x7407f7 / 0x4000000 | yes (round 10 D5 is fixed) |
| Sub-20 penalty exemption | every heal path | levelPenalty returns 1 | core/spell.go:803-811, 341-345 | every druid heal sets ProcMaskSpellHealing (healing.go:72, swiftmend.go:47); Bloom is not a cast rank | yes |
| Heal crit and tick crit | all | 1.5x, ticks on the healing table | core/spell_outcome.go:78-87, 1083 | Attr8 0x200 on every Rejuvenation, Regrowth, Tranquility and Wild Growth rank (all in periodic_crit_auto_gen.go) | yes |
| Rank per band | HT, Regrowth, Rejuvenation, Tranquility | engine learn level = SpellLevel | healing.go:54 | trainables/druid.json levels = SpellLevels for all ranks | yes |
| **Idol of Swiftness** | item 279250 / spell 1306488 | Swiftmend cooldown -3 s | **druid/items.go:37-40 relicClassMasks has no Swiftmend family; the item is never registered (no NewEquipModItemEffect)** | **ItemXItemEffect -> 1306488: aura 107 misc 11 bp -3000 on m1 0x2 (Swiftmend); Enchanting-crafted (loot.json crafted:enchanting); requires level 60, BoE epic** | **no (see 2.1)** |
| Grovekeeper Raiment 5p (not in bis) | set 2112 / 1301250 | Swiftmend -3 s | item_sets_tier1.go:33, 38-43 | ItemSetSpell threshold 5 | yes (this table already maps Swiftmend) |

### Restoration shaman

| spell / talent | id(s) | field | engine (file:line) | client (table, row) | match |
|---|---|---|---|---|---|
| Healing Wave r1-r10 | 331, 332, 547, 913, 939, 959, 8005, 10395, 10396, 25357 | bp range, ppl, levels, coef, mana, cast | gen (healing_wave.go:11-25) | SpellEffect eff 10 (10396: 1344 / 0.13251 / 5.5 / 56 to 61 / 0.857; 560 mana; 3000 cast) | yes, all |
| Healing Wave r2 (bands 40-59 and bare 60, rank-choices.json) | 332 | 71 (64 to 78), +1/level to 11, coef 0.571, 45 mana, 2.0 s cast | gen | SpellLevels 6/11; SLA 374 acq 0 | yes (no sub-20 penalty: a heal) |
| Lesser Healing Wave r1-r6 | 8004, 8008, 8010, 10466, 10467, 10468 | bp, var, ppl, levels, coef, mana, cast | gen; r6 hand override lesser_healing_wave.go:15-19 | 10468: 820 / 0.10909 / 4.2 / 60 to 65; 27624 (880) is the BC row | yes |
| Chain Heal r1-r3 | 1064, 10622, 10623 | bp, ppl, levels, coef, mana, cast | gen | 10623: 490 / 0.1322 / 3.3 / 54 to 59, 0.714, 405 mana, 2500 cast | yes |
| Chain Heal | all | jumps and falloff | chain_heal.go:15-16 (2 jumps, x0.5) | ChainTargets 3, ChainAmplitude 0.5 | yes (jump targeting: round 10 R5, open) |
| Riptide r1-r3 | 408521, 1239242, 1239243 | direct bp/var/ppl/levels/coef, mana | gen | 841 / 0.1 / 4.9 / 60 to 68, 0.214; 385 mana | yes |
| Riptide | all | tick bp/ppl, coef, period, ticks, cooldown | riptide.go:15-35 | e1 aura 8: 89 / 115 / 161 (+3.7 a level from 60), coef 0.1, 3000/15000 = 5; CategoryRecovery 6000; gcd 1500 | yes |
| Riptide | all | Chain Heal bonus | riptide.go:24, chain_heal.go:51-54 | e2 aura 4 bp 25 | yes |
| Riptide r2 (bare 60, rank-choices.json) | 1239242 | 612 +4.2 a level to 58, tick 115 | gen + riptide.go:33 | SpellLevels 50/58 | yes |
| Healing Stream Totem r1-r5 | 5394, 6375, 6377, 10462, 10463 | mana, level, life, GCD | water_totems.go:13-15, 22, totems.go:15 | 40/50/60/70/80 mana; 20/30/40/50/60; dur 300000; gcd 1000 | yes |
| Healing Stream heal r1-r5 | 5672, 6371, 6372, 10460, 10461 | bp, coef, period, party | gen; water_totems.go:23, 114-128 | 5/6/7/9/11, 0.022, per 2000, eff 35 party | yes |
| Healing Stream heal | 10461 | crit; Restorative Totems | water_totems.go:81, 101-106 | Attr8 0x1001200; settled (every-table crit only); RT e1 DOT op 50% on 0x2000 | yes |
| Mana Tide Totem r1-r4 | 16190, 17354, 17359, 1325256 | mana, level, cooldown, life, GCD | gen; mana_tide_totem.go:16-19 | 10/20/30/60 mana; 1/38/48/58; CategoryRecovery 300000; dur 13000; gcd 1000 | yes |
| Mana Tide restore r1-r4 | 16191, 17355, 17360, 1325257 | per tick, period, ticks | gen; mana_tide_totem.go:17-18 | aura 24: 88/133/197/290 per 3000, 13 s = 4 ticks | yes |
| Mana Spring r1-r4 (not cast) | 5675, 10495, 10496, 10497 / 5677, 10491, 10493, 10494 | ids, mana, level, restore, life | water_totems.go:144-147, 167 | 40/60/80/100 mana; 26/36/46/56; 4/6/8/10 per 2000; 300000 | yes |
| Water Shield | 408510 / 408511 | charges, ICD, duration, GCD, 2% mana, proc mask | water_shield.go:36-41, 116 | 3 charges, ICD 3500, 600000, gcd 1500, eff 137 bp 2, ptm 0x362a8 (no periodic bit) | yes |
| Nature's Swiftness | 16188 | cooldown, GCD, mask, cooldown on consume | talents.go:374-430 | Recovery 180000, gcd 0, A0 0x2000000 COOLDOWN_ON_EVENT, mask 0x9c3 + m2 0x100000 | yes |
| Improved Healing Wave | 16182 | -0.1 s/rank | talents_restoration.go:15, 53-57 | curve 85307, mask 0x40 | yes |
| Tidal Focus | 16179 | -1%/rank cost (HW, LHW, CH, Riptide; not totems); +1 hit/rank | talents_restoration.go:18-19, 60-67 | curve 85321; mask 0x1c0 + m2 0x10; e1/e2 aura 54/55 curves 126021/126020 | yes |
| Mindfulness | 1223033 | 17/33/50% | talents_restoration.go:43, 101-103 | curve 85316 17/33/50 | yes |
| Healing Focus | 16181 | 23/47/70% pushback on HW, LHW, CH | talents.go:593, 601-604 | curve 85317; mask 0x1c0 + m2 0x100000 | yes |
| Tidal Mastery | 16194 | +1%/rank crit on heals incl. Riptide | talents_restoration.go:22, 69-75 | curve 85308; mask 0x5c0 + m2 0x10 | yes |
| Restorative Totems | 16187 | Mana Spring 5%/rank, Healing Stream 10%/rank | talents_restoration.go:32-33, 108-116 | curves 85320 (5..25, misc 8, 0x4000), 85319 (10..50, DOT op, 0x2000) | yes |
| Healing Way | 29206 | 8/17/25% on HW | talents_restoration.go:38, 85-91 | curve 85322, flat aura 108 on 0x40 | yes (stacking: R4, open) |
| Purification | 16178 | 2%/rank on HW, LHW, CH, Riptide (direct and ticks) | talents_restoration.go:25, 77-83 | curves 85323/87379; e0 0x1c0 + m2 0x10, e1 DOT op m2 0x10 | yes |
| Ancestral Knowledge | 17485 | 2%/rank Intellect | talents.go:28, 56-58 | curve 85266, misc_1 8 | yes |
| Thundering Strikes | 16255 | 1%/rank, every table | talents.go:64 | aura 290, curve 85275 | yes |
| Ancestral Healing | 16176 | armor only (unmodelled) | talents.go:934 | trig 16177 armor, curve 85318 | yes (no healing value) |
| Totem of Urgency | item 279249 / spell 1306448 | +4% crit on LHW | items.go:37-38 (mask 1<<7), :129 registered | aura 107 misc 7 bp 4 on 0x80 | yes (round 10 R1 fixed) |
| Tidal Totem | 272431 | -5% HW cost | not registered | vendor `opens: later` | n/a (not available yet) |
| Spiritcaller 5p (not in bis) | set 2109 / 1301713 | Riptide -1 s | shaman/item_sets_tier1.go:25, 33 | aura 107 misc 11 bp -1000 on m2 0x10 | yes |
| Sub-20 penalty exemption | every heal path | levelPenalty returns 1 | core/spell.go:803-811 | newHealingSpellConfig ProcMaskSpellHealing (healing_spell.go:49) for HW, LHW, CH and Riptide (direct and Hot); the Healing Stream tick spell has no cost or GCD, so it is not a cast rank | yes |
| Tick crit | Riptide, Healing Stream | PERIODIC_CAN_CRIT | core/periodic_crit_auto_gen.go | 408521, 1239242, 1239243, 5672, 6371, 6372, 10460, 10461 all listed | yes |
| Rank per band | HW, LHW, CH, Riptide, HST, MTT | engine learn level | gen *Level arrays | trainables/shaman.json levels = SpellLevels | yes |
| Talent string order | both classes | proto field numbers | core/character.go:759-776; talents_auto_gen.go (70753) | client tier/column order | yes |

## 2. Mismatches

### 2.1 Idol of Swiftness (279250) is never registered, so Restoration druid's relic slot stays empty at 60 (CONFIRMED gap)
- **Engine.** `druid/items.go:37-40`: `relicClassMasks` maps only Rip (m2 1<<21) and Shifting Power (m1 1<<31). The comment at :30-36 says an idol in the Swiftmend family "cannot be registered", and nothing calls `NewEquipModItemEffect(279250, ...)`. Its row is already generated (`druid/relic_mods_auto_gen.go`, 279250: aura 107, op 11, -3000, families {0, 2, 0, 0}). `druid/item_sets_tier1.go:33` already maps that same client family {0, 1<<1} to `DruidSpellMaskSwiftmend`.
- **Client.**
  - Item 279250 (ItemSparse): Idol of Swiftness, relic, item level 65, requires level 60, epic, BoE.
  - Its equip spell 1306488 "Swiftmend Cooldown Reduction": e0 1333954, aura 107, misc 11 (cooldown), bp -3000, mask m1 0x2. Swiftmend 18562 is SpellClassOptions m1 0x2 with RecoveryTime 15000.
  - Server tooltip (nether.wowhead.com/forever/tooltip/item/279250, read 2026-11-03): "Equip: Reduces the cooldown of your Swiftmend spell by 3 sec."
  - loot.json lists it under `crafted:enchanting`, next to Totem of Urgency, with no `opens`.
- **Reach.** Restoration druid at 60, raid and bare. It is the shaman's round 10 R1 (Totem of Urgency) gap for the druid, and R1 is fixed. Every published druid band has an empty ranged slot. The equippable item has zero value, so the gear search never picks it. Swiftmend is cast with a mana-pace gate, so the gain is a 12 s cooldown in place of 15 s (up to 25% more Swiftmends). Not measured.
- **Confidence.** High that the gap exists (client row, server text, generated row and engine table all agree). The size is unknown.

### 2.2 Reflection ranks 1 and 2 are 16.67% and 33.33%; the client says 17% and 33% (CONFIRMED, tiny)
- **Engine.** `druid/talents_restoration.go:112,118`: `SpiritRegenRateCasting += rank/3 * 0.5`.
- **Client.** 17106 e0 694418, aura 134. TraitDefinitionEffectPoints def 134381 uses curve 85365: (1, 17), (2, 33), (3, 50). The shaman's Mindfulness, which uses the same aura and curve shape, is written as 0.17 / 0.33 / 0.50 (`shaman/talents_restoration.go:43`) and matches.
- **Reach.** Only builds that hold Reflection 1 or 2. The published band-20 entry is `5051...` (Reflection 1), so it gets 16.67% where the client gives 17%: +0.33 points of regen while casting. Ladder-first now reserves Reflection 3 first, so after the nightly regen no band may reach it. Rank 3 is exact (50%).
- **Confidence.** High that it is a mismatch. Its effect is negligible (well under 0.1% at band 20).

### 2.3 The published bis entries trail the curated sources (data, not engine; the nightly owns it)
- `data/builds/1.60.1.70753/bis/druid-restoration.json` (generated 2026-11-01T20:25Z) still carries `4532200100000000-...-5053005153113200` at 60, with Improved Wrath 4, which heals nothing (mask 0x1). The guide build is `web/src/content/guides/druid/restoration.md:6` `05312001/0/55530051531132`: Heart of the Wild 5, Improved Wrath 0, Improved Moonfire 1.
- Band 20 is `5051` (Nature's Focus 5, Naturalist 5, Reflection 1). But ladder-first.json reserves Heart of the Wild 5, Reflection 3, Bloom and Nature's Swiftness first.
- The shaman 60 ranged slot is empty although Totem of Urgency is now registered. The druid ranged slot is empty as well (2.1).
- **Reach:** every druid band and shaman 60. Per the memory rule, nightly bis.yml regenerates these files, so this is not a hand fix. It is listed because the build strings in this lane's brief are the stale ones.

### 2.4 OPEN: Swiftmend's payout with Nature's Splendor
`druid/swiftmend.go:68` pays out the source HoT's `NumberOfTicks`. With Splendor taken at 60 that is 5 Rejuvenation ticks or 9 Regrowth ticks (15 s and 27 s). The client row (18562: one heal effect, bp 1) and the text ("the full duration of the periodic effect") do not say whether the extended duration counts. Round 11's logs had no druid with Splendor. This needs one logged level-60 Swiftmend with Splendor. If only the base duration counts, the engine overpays Swiftmend off Regrowth by 9/7 (+29%).

### Hygiene (no number moves)
- `shaman/talents_restoration.go:16` labels Tidal Focus as node 104735 (Totemic Focus's node). Tidal Focus is 104729.
- `shaman/talents.go:588` labels Healing Focus as node 104792. It is 104733.

No other mismatch. Every hand table (druid/healing_tables.go, the LHW r6 override, the Riptide ticks, the Healing Stream / Mana Spring / Mana Tide hand arrays) and every hand-written cooldown, duration, tick count, mask and talent value above equals the 70753 client row.
