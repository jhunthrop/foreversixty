# Sim-accuracy round 9: ranks lane (owner ruling 3)

Ruling: rotations may downrank. When a fight is mana-bound, the rotation casts the rank that does the most damage or healing over the fight. A deliberate lower rank in a curated rotation is kept and is never undone by a top-rank rule.

Integrator's rule for this lane: adopt any rank, higher or lower, that raises the published number beyond the error. That means seeds 1 and 50001 at 10,000 iterations both agree. There is no 1% or 2% bar. Two further checks apply:
- A healer row must not turn a fight that runs dry into one that ends with mana unspent.
- Defensive cooldowns are excluded (Divine Protection, Divine Shield, Ice Barrier, Mana Shield, the armors and the like). Threat per second gives them no credit for the immunity or shield they provide.

Where it was measured: site branch `wt/r9-ranks`, fork branch `lane/r9-ranks`, horde, published preset per band (bare 20 to 60, raid 60). Healers are scored with `-heal`, which is effective HPS on data/curated/heal-profile.json. paladin-protection is scored with `-tank` (threat/s); its out-of-mana seconds come from the plain run.
- **"int/r9"** means the integrator's merged engine, pinned here at 3b4ae7a9a (crit, casters, hybrids, healers, items, warlock, auras and physical), plus those lanes' site branches merged into a scratch site worktree. "Before" is the same tree without this lane.
- A first pass on the round-8 engine (40021374a) is superseded by the int/r9 numbers below.
- The int/r9 engine moved while I measured (the physical and auras lanes merged). The tables below were re-run on one pinned snapshot. Every adopted row was re-confirmed there.

## Mechanism

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| M1 | **The rank rewrite undid deliberate downranks below 60.** `sim/request/rotation_ranks.go` moved every ranked spell to the band's top learned rank. Below 60 that raised Holy and Discipline's Renew rank 3, Discipline's Power Word: Shield rank 4, and Holy paladin's party Flash of Light rank 4 and Holy Shock rank 1. | Failing first: `TestRewriteKeepsTheAuthoredDownranksBelowSixty`. At 50, Renew 6075 came out 10928 and Shield 3747 came out 10899. | CONFIRMED-BUG | S 5c63f6719 |
| M2 | **Default rule.** `spellranks.LearnedSpellID` (was HighestLearnedSpellID) casts the highest rank learned that is **not above the authored rank**. A rotation written at the top rank behaves as before. An id naming an Ahn'Qiraj book rank caps at the top trainer rank. | `TestLearnedUpToKeepsAnAuthoredDownrank`, `TestRankOfResolvesAChosenRank` | - | S 5c63f6719 |
| M3 | **Per-band rank choices.** `data/curated/rank-choices.json` holds rows per spec and authored id: `{from_level, to_level, rank, preset?}`. `make rank-choices-sync` copies it to `sim/request`, and it is parsed once per process. Everything resolves through `request.rankResolver`: the site's /sim, the ranker and the nightly (`request.rotation`), spec-breakdown `-rotation` and the ladder's zero-casts rule. | `TestRewriteHonoursAPerBandRankChoice`, `TestThePublishedRotationTakesTheCuratedRankChoices`, `TestCuratedRankChoicesAreCastableDownranks`, `TestEmbeddedRankChoicesMatchCurated` | - | S 5c63f6719, 5f1cbfacf |
| M4 | **Preset-aware rows.** A row may name `"preset": "bare"` or `"raid"`; a row without one applies to both. `api.SimRequest` gains an optional `preset`. When it is empty, the request takes the preset its level is published under (`request.PublishedPreset`), which is the /sim page's default. Unknown names are refused. At 60 the curated rotation still runs as written unless a row for that preset names one of its ids. A row at 60 for both presets is refused, because that rank belongs in the rotation itself. | Failing first: `TestRankChoicesFollowThePreset`, `TestRewriteAtSixtyTakesOnlyThePresetsRows`, `TestBuildRefusesAnUnknownPreset` | - | S 49356dbcf |
| M5 | **A probe can pin a rank.** spec-breakdown `-rotation` used to raise the top rank back before the probe ran. It now uses the same resolver and sends `-preset`. | `sim/cmd/spec-breakdown/main.go` loadRotation | - | S 5c63f6719, 49356dbcf |
| M6 | **The addon raised the same downranks.** `data/pipeline/addonrotation.py` now follows M2 and M3 below 60: it never raises a rank above the rotation's own, and a bare row wins. It is built on the healers lane's a65bd8111, cherry-picked unchanged as 547d3f60f. | `test_a_deliberate_lower_rank_is_never_lifted_below_sixty`, `test_a_band_casts_the_rank_choice_the_curated_table_names` | CONFIRMED-BUG | S 134d5065f, 49356dbcf |
| M7 | **Holy paladin's fork mana test priced Holy Light rank 8 at 660,** the book rank 9's cost. The client cost is 580 (spellconst 10329; `sim/paladin/holy_light.go` agrees). The tank's Flash of Light used to dominate the sum and hid the error. | `sim/paladin/holy/rotation_test.go` | CONFIRMED-BUG (test only) | E 0a9ef0d0c |

A flat rank per band was enough. A "behind the pace" second line was never better than the best flat rank: Elemental 50 measured +38.2% paced against +39.4% flat, and at raid 60 flat rank 9 beat the paced line.

## Band 60 raid: rotation ranks changed (curated rotation and fork copy)

These are the headline numbers. The rotations, the guides and the addon read these ranks.

| spec | change | raid 60 seed 1 / 50001 | raid 60, seed 1 | mana, s |
|---|---|---|---|---|
| shaman-elemental | Lightning Bolt rank 9 (15207) all fight, top rank 10 before | +0.58% / +0.56% | 464.2 to 466.9 DPS | OOM 5.5 to 0.3 |
| priest-holy | the tank's Flash Heal rank 2 (9472) | +2.27% / +2.18% | 678.0 to 693.4 HPS | lasts 137.2 to 461.5 |
| priest-discipline | Heal rank 1 (2054) | +1.02% / +1.02% | 618.6 to 624.9 | lasts 112.4 to 114.8 |
| paladin-holy | the tank's Flash of Light rank 1 (19750) | +1.44% / +1.55% | 632.3 to 641.4 | lasts 250.0 to 257.1 |
| druid-restoration | Rejuvenation rank 4 (2090) on every target | +0.28% / +0.30% | 777.0 to 779.2 | lasts 287.7 to 383.3 |

- **Elemental gain shrank on the pinned engine.** Before the auras and physical merges it measured +2.08% / +2.04%.
- **Restoration druid's party-only split did not hold on int/r9.** The healers lane measured +3.2% on its own engine. On int/r9, all casts at rank 4 did better.
- **Bare 60 keeps the stronger rank where these changes would lose there.** Through the bare-60 rows below:
  - Holy keeps Flash Heal rank 7 on the tank.
  - Discipline keeps Heal rank 4.
  - Restoration druid keeps Rejuvenation rank 10.

## Per-band rows (rank-choices.json), re-confirmed on int/r9

Each row is the published rotation (the raid-60 ranks above in place) with its band's rows, measured against the same rotation under the default rule alone. All rows are bare. Healer rows are **measured on the level-60 heal profile**. The healers lane notes that this profile is the level-60 Phase 1 fight at every band, so below 60 these rows are pure mana efficiency. The owner may want to rule on whether the profile should scale below 60.

| spec | levels | choice | gain seed 1 / 50001 | seed 1 | mana, s | note |
|---|---|---|---|---|---|---|
| druid-balance | 30-39 | Wrath rank 4 (5179); Moonfire rank 3 (8925); Insect Swarm rank 1 (5570) | +15.58% / +15.74% | 62.9 to 72.7 DPS | OOM 49.2 to 10.3 |  |
| druid-restoration | 20-29 | Rejuvenation rank 1 (774) | +2.08% / +2.08% | 38.5 to 39.3 HPS | lasts 253.2 to 252.4 |  |
| druid-restoration | 30-39 | Rejuvenation rank 2 (1058) | +6.19% / +6.20% | 100.1 to 106.3 HPS | lasts 211.2 to 222.5 |  |
| druid-restoration | 40-49 | Rejuvenation rank 7 (8910) | +1.94% / +1.94% | 169.8 to 173.1 HPS | lasts 255.0 to 213.1 |  |
| druid-restoration | 50-59 | Rejuvenation rank 8 (9839) | +1.22% / +1.18% | 262.5 to 265.7 HPS | lasts 246.8 to 213.9 |  |
| druid-restoration | 60, bare only | Rejuvenation rank 10 (9841) | +5.15% / +5.21% | 367.2 to 386.1 HPS | lasts 243.6 to 222.5 |  |
| hunter-beast-mastery | 40-49 | Aimed Shot rank 1 (19434) | +0.39% / +0.39% | 154.0 to 154.6 DPS | OOM 104.3 to 99.6 |  |
| hunter-beast-mastery | 60, bare only | Aimed Shot rank 1 (19434) | +0.38% / +0.34% | 291.8 to 292.9 DPS | OOM 91.6 to 86.7 |  |
| hunter-marksmanship | 60, bare only | Aimed Shot rank 1 (19434) | +0.32% / +0.25% | 283.3 to 284.2 DPS | OOM 46.3 to 46.4 |  |
| mage-arcane | 30-39 | Frostbolt rank 4 (7322) | +1.34% / +1.46% | 82.0 to 83.1 DPS | OOM 22.4 to 12.0 |  |
| mage-fire | 30-39 | Scorch rank 1 (2948) | +1.00% / +1.15% | 69.7 to 70.4 DPS | OOM 75.7 to 64.7 |  |
| mage-fire | 40-49 | Scorch rank 1 (2948); Fire Blast rank 4 (8412) | +6.30% / +6.50% | 144.5 to 153.6 DPS | OOM 55.7 to 13.4 |  |
| mage-fire | 50-59 | Scorch rank 1 (2948) | +6.26% / +6.31% | 179.0 to 190.2 DPS | OOM 58.2 to 19.4 |  |
| paladin-holy | 40-49 | Holy Light rank 3 (647) | +5.45% / +5.64% | 165.0 to 174.0 HPS | lasts 36.6 to 111.3 |  |
| paladin-holy | 50-59 | Holy Light rank 3 (647) | +10.07% / +10.08% | 206.5 to 227.3 HPS | lasts 44.8 to 191.9 |  |
| paladin-holy | 60, bare only | Holy Light rank 4 (1026) | +4.91% / +4.70% | 284.9 to 298.9 HPS | lasts 59.2 to 143.6 |  |
| paladin-protection | 40-49 | Consecration rank 1 (26573) | +1.21% / +1.32% | 189.9 to 192.2 threat/s | OOM 19.4 to 17.3 |  |
| paladin-protection | 50-59 | Consecration rank 1 (26573) | +1.80% / +1.73% | 282.9 to 288.0 threat/s | OOM 16.3 to 14.0 |  |
| paladin-protection | 60, bare only | Consecration rank 1 (26573) | +2.88% / +2.95% | 412.7 to 424.6 threat/s | OOM 5.3 to 3.9 |  |
| paladin-retribution | 30-39 | Consecration rank 1 (26573) | +0.57% / +0.72% | 70.0 to 70.4 DPS | OOM 0.0 to 0.0 | adopted: Consecration is rationed by a mana line |
| paladin-retribution | 40-49 | Consecration rank 1 (26573) | +0.99% / +1.09% | 100.8 to 101.8 DPS | OOM 0.0 to 0.0 | adopted: Consecration is rationed by a mana line |
| paladin-retribution | 50-59 | Consecration rank 1 (26573) | +1.65% / +1.65% | 163.2 to 165.9 DPS | OOM 1.3 to 0.1 |  |
| priest-discipline | 30-39 | Renew rank 1 (139); Power Word: Shield rank 5 (6065) | +2.35% / +2.35% | 89.3 to 91.4 HPS | lasts 44.5 to 46.2 |  |
| priest-discipline | 40-49 | Binding Heal rank 1 (401937); Power Word: Shield rank 2 (592) | +10.80% / +10.80% | 141.7 to 157.0 HPS | lasts 49.8 to 55.5 |  |
| priest-discipline | 50-59 | Binding Heal rank 1 (401937); Power Word: Shield rank 8 (10899) | +13.71% / +13.71% | 194.8 to 221.5 HPS | lasts 52.3 to 62.6 |  |
| priest-discipline | 60, bare only | Binding Heal rank 1 (401937); Heal rank 4 (6064) | +6.81% / +6.92% | 275.9 to 294.7 HPS | lasts 53.8 to 62.6 |  |
| priest-holy | 40-49 | Flash Heal rank 4 (9474) | +3.98% / +3.98% | 115.6 to 120.2 HPS | lasts 35.7 to 30.4 |  |
| priest-holy | 50-59 | Flash Heal rank 6 (10916) | +13.90% / +13.97% | 159.7 to 181.9 HPS | lasts 55.7 to 40.2 |  |
| priest-holy | 60, bare only | Flash Heal rank 7 (10917); Lightwell rank 1 (724); Prayer of Healing rank 1 (596) | +11.02% / +10.99% | 281.2 to 312.2 HPS | lasts 88.6 to 55.0 |  |
| priest-shadow | 30-39 | Shadow Word: Pain rank 3 (970); Devouring Plague rank 1 (2944) | +4.44% / +4.43% | 56.3 to 58.8 DPS | OOM 91.9 to 92.8 |  |
| priest-shadow | 40-49 | Devouring Plague rank 1 (2944); Mind Blast rank 4 (8104) | +4.69% / +4.77% | 123.6 to 129.4 DPS | OOM 70.8 to 58.7 |  |
| priest-shadow | 50-59 | Devouring Plague rank 1 (2944); Mind Blast rank 2 (8102); Shadow Word: Death rank 2 (1309633) | +7.42% / +7.35% | 186.1 to 199.9 DPS | OOM 67.4 to 40.6 |  |
| priest-shadow | 60, bare only | Devouring Plague rank 1 (2944); Shadow Word: Death rank 2 (1309633) | +3.47% / +3.49% | 343.4 to 355.3 DPS | OOM 19.0 to 7.6 |  |
| shaman-elemental | 30-39 | Lightning Bolt rank 4 (915) | +8.53% / +8.51% | 43.4 to 47.1 DPS | OOM 95.8 to 74.8 |  |
| shaman-elemental | 40-49 | Lightning Bolt rank 4 (915); Flame Shock rank 1 (8050); Lava Burst rank 1 (408490) | +21.66% / +21.95% | 77.1 to 93.8 DPS | OOM 101.7 to 42.4 |  |
| shaman-elemental | 50-59 | Lightning Bolt rank 4 (915); Lava Burst rank 1 (408490); Flame Shock rank 2 (8052) | +50.60% / +50.79% | 100.8 to 151.8 DPS | OOM 98.8 to 1.4 |  |
| shaman-elemental | 60, bare only | Lightning Bolt rank 4 (915); Lava Burst rank 2 (1238299) | +36.14% / +36.27% | 140.0 to 190.6 DPS | OOM 79.9 to 2.0 |  |
| shaman-enhancement | 60, bare only | Strength of Earth Totem rank 1 (8075) | +1.13% / +0.99% | 212.0 to 214.4 DPS | OOM 0.0 to 0.0 | NOT adopted: not mana-bound |
| shaman-restoration | 20-29 | Healing Wave rank 3 (547) | +8.50% / +8.50% | 45.9 to 49.8 HPS | lasts 31.9 to 50.1 |  |
| shaman-restoration | 30-39 | Healing Wave rank 3 (547) | +13.35% / +13.35% | 83.9 to 95.1 HPS | lasts 42.8 to 119.8 |  |
| shaman-restoration | 40-49 | Healing Wave rank 3 (547) | +19.16% / +19.16% | 142.5 to 169.8 HPS | lasts 39.1 to 243.9 |  |
| shaman-restoration | 50-59 | Healing Wave rank 3 (547) | +10.29% / +10.29% | 222.5 to 245.4 HPS | lasts 48.1 to 62.0 |  |
| shaman-restoration | 60, bare only | Healing Wave rank 3 (547); Riptide rank 2 (1239242) | +12.86% / +12.83% | 332.8 to 375.6 HPS | lasts 74.8 to 67.4 |  |
| warlock-destruction | 30-39 | Shadow Bolt rank 4 (1088) | +0.57% / +0.57% | 87.7 to 88.2 DPS | OOM 2.8 to 1.8 |  |

- **Holy paladin band 40:** the best single rank was Holy Light rank 2 (+12.9%), but its mana lasts 402 s of a 300 s fight, so it ends with mana unspent. Holy Light rank 3 is the best choice that still runs dry.
- **Not adopted:** Enhancement's Strength of Earth Totem rank 1 at bare 60 (+1.1% / +1.0%). Enhancement never runs dry, and a weaker Strength buff raising damage needs explaining first (see "For other lanes").
- **Swept with no gain:** Affliction and Demonology warlocks, Frost mage, Survival hunter, and Balance outside band 30.
- **Not swept:** rogues, warriors and feral druids, which use no mana. Their rotations name only top ranks and have no row, so they cast exactly what they did before.

Method: `scratchpad r9/ranks/int/`. `sweep.py` tries every learned rank of every ranked spell the rotation names, higher or lower, one at a time, at 3,000 iterations; defensives are excluded. `greedy2.py` combines ranks while each step adds at least 0.3%. `adopt.py` and `reconfirm.py` confirm at 10,000 iterations on both seeds. The pass mark is: DPS gain above twice the combined stated error at both seeds; healer and tank gain above 0.2% at both seeds, with the seeds agreeing. The engine applies the sub-20 coefficient penalty (`sim/core/spell.go` levelPenalty), so ranks learned below 20 were candidates.

## Every spec and band on int/r9 (published rotation, before to after this lane)

Columns: published metric at seed 1, the change at seed 1 / 50001, and mana seconds. Mana seconds are out-of-mana seconds for damage and "mana lasts" for healers; paladin-protection reports none in the tank fight.

| spec | band | metric, seed 1 | change | mana, s |
|---|---|---|---|---|
| druid-balance | bare 20 | 25.0 to 25.0 | +0.00% / +0.00% | 21.7 to 21.7 |
| druid-balance | bare 30 | 62.9 to 72.7 | +15.58% / +15.74% | 49.2 to 10.3 |
| druid-balance | bare 40 | 151.0 to 151.0 | +0.00% / +0.00% | 6.3 to 6.3 |
| druid-balance | bare 50 | 191.5 to 191.5 | +0.00% / +0.00% | 8.7 to 8.7 |
| druid-balance | bare 60 | 280.4 to 280.4 | +0.00% / +0.00% | 0.5 to 0.5 |
| druid-balance | raid 60 | 565.2 to 565.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| druid-restoration | bare 20 | 38.5 to 39.3 | +2.08% / +2.08% | 253.2 to 252.4 |
| druid-restoration | bare 30 | 96.6 to 106.3 | +10.04% / +10.05% | 199.3 to 222.5 |
| druid-restoration | bare 40 | 173.1 to 173.1 | +0.00% / +0.00% | 213.1 to 213.1 |
| druid-restoration | bare 50 | 265.7 to 265.7 | +0.00% / +0.00% | 213.9 to 213.9 |
| druid-restoration | bare 60 | 386.1 to 386.1 | +0.00% / +0.00% | 222.5 to 222.5 |
| druid-restoration | raid 60 | 777.0 to 779.2 | +0.28% / +0.30% | 287.7 to 383.3 |
| hunter-beast-mastery | bare 20 | 64.7 to 64.7 | +0.00% / +0.00% | 100.0 to 100.0 |
| hunter-beast-mastery | bare 30 | 99.2 to 99.2 | +0.00% / +0.00% | 115.4 to 115.4 |
| hunter-beast-mastery | bare 40 | 154.0 to 154.6 | +0.39% / +0.39% | 104.3 to 99.6 |
| hunter-beast-mastery | bare 50 | 215.5 to 215.5 | +0.00% / +0.00% | 106.6 to 106.6 |
| hunter-beast-mastery | bare 60 | 291.8 to 292.9 | +0.38% / +0.34% | 91.6 to 86.7 |
| hunter-beast-mastery | raid 60 | 633.4 to 633.4 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-marksmanship | bare 20 | 63.3 to 63.3 | +0.00% / +0.00% | 20.5 to 20.5 |
| hunter-marksmanship | bare 30 | 99.5 to 99.5 | +0.00% / +0.00% | 16.1 to 16.1 |
| hunter-marksmanship | bare 40 | 152.6 to 152.6 | +0.00% / +0.00% | 36.5 to 36.5 |
| hunter-marksmanship | bare 50 | 213.2 to 213.2 | +0.00% / +0.00% | 38.1 to 38.1 |
| hunter-marksmanship | bare 60 | 283.3 to 284.2 | +0.32% / +0.25% | 46.3 to 46.4 |
| hunter-marksmanship | raid 60 | 623.0 to 623.0 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | bare 20 | 47.2 to 47.2 | +0.00% / +0.00% | 0.1 to 0.1 |
| hunter-survival | bare 30 | 80.3 to 80.3 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | bare 40 | 152.8 to 152.8 | +0.00% / +0.00% | 19.1 to 19.1 |
| hunter-survival | bare 50 | 201.9 to 201.9 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | bare 60 | 295.2 to 295.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| hunter-survival | raid 60 | 734.9 to 734.9 | +0.00% / +0.00% | 0.0 to 0.0 |
| mage-arcane | bare 20 | 43.8 to 43.8 | +0.00% / +0.00% | 19.1 to 19.1 |
| mage-arcane | bare 30 | 82.0 to 83.1 | +1.34% / +1.46% | 22.4 to 12.0 |
| mage-arcane | bare 40 | 159.6 to 159.6 | +0.00% / +0.00% | 6.6 to 6.6 |
| mage-arcane | bare 50 | 206.0 to 206.0 | +0.00% / +0.00% | 10.8 to 10.8 |
| mage-arcane | bare 60 | 329.5 to 329.5 | +0.00% / +0.00% | 10.7 to 10.7 |
| mage-arcane | raid 60 | 600.2 to 600.2 | +0.00% / +0.00% | 0.7 to 0.7 |
| mage-fire | bare 20 | 34.5 to 34.5 | +0.00% / +0.00% | 68.4 to 68.4 |
| mage-fire | bare 30 | 69.7 to 70.4 | +1.00% / +1.15% | 75.7 to 64.7 |
| mage-fire | bare 40 | 144.5 to 153.6 | +6.30% / +6.50% | 55.7 to 13.4 |
| mage-fire | bare 50 | 179.0 to 190.2 | +6.26% / +6.31% | 58.2 to 19.4 |
| mage-fire | bare 60 | 329.6 to 329.6 | +0.00% / +0.00% | 24.2 to 24.2 |
| mage-fire | raid 60 | 660.8 to 660.8 | +0.00% / +0.00% | 0.7 to 0.7 |
| mage-frost | bare 20 | 44.0 to 44.0 | +0.00% / +0.00% | 52.9 to 52.9 |
| mage-frost | bare 30 | 75.7 to 75.7 | +0.00% / +0.00% | 21.9 to 21.9 |
| mage-frost | bare 40 | 183.1 to 183.1 | +0.00% / +0.00% | 6.7 to 6.7 |
| mage-frost | bare 50 | 244.5 to 244.5 | +0.00% / +0.00% | 11.5 to 11.5 |
| mage-frost | bare 60 | 343.0 to 343.0 | +0.00% / +0.00% | 4.9 to 4.9 |
| mage-frost | raid 60 | 594.1 to 594.1 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-holy | bare 20 | 29.5 to 29.5 | +0.00% / +0.00% | 43.0 to 43.0 |
| paladin-holy | bare 30 | 65.3 to 78.4 | +20.06% / +19.91% | 28.0 to 28.2 |
| paladin-holy | bare 40 | 116.5 to 174.0 | +49.36% / +49.57% | 35.4 to 111.3 |
| paladin-holy | bare 50 | 143.4 to 227.3 | +58.51% / +58.44% | 42.8 to 191.9 |
| paladin-holy | bare 60 | 206.5 to 298.9 | +44.75% / +44.39% | 57.1 to 143.6 |
| paladin-holy | raid 60 | 632.3 to 641.4 | +1.44% / +1.55% | 250.0 to 257.1 |
| paladin-protection | bare 20 | 40.0 to 40.0 | +0.00% / +0.00% | 59.2 to 59.2 |
| paladin-protection | bare 30 | 66.0 to 66.0 | +0.00% / +0.00% | 13.8 to 13.8 |
| paladin-protection | bare 40 | 189.9 to 192.2 | +1.21% / +1.32% | 19.4 to 17.3 |
| paladin-protection | bare 50 | 282.9 to 288.0 | +1.80% / +1.73% | 16.3 to 14.0 |
| paladin-protection | bare 60 | 412.7 to 424.6 | +2.88% / +2.95% | 5.3 to 3.9 |
| paladin-protection | raid 60 | 993.8 to 993.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-retribution | bare 20 | 39.0 to 39.0 | +0.00% / +0.00% | 0.1 to 0.1 |
| paladin-retribution | bare 30 | 70.0 to 70.4 | +0.57% / +0.72% | 0.0 to 0.0 |
| paladin-retribution | bare 40 | 100.8 to 101.8 | +0.99% / +1.09% | 0.0 to 0.0 |
| paladin-retribution | bare 50 | 163.2 to 165.9 | +1.65% / +1.65% | 1.3 to 0.1 |
| paladin-retribution | bare 60 | 241.4 to 241.4 | +0.00% / +0.00% | 0.0 to 0.0 |
| paladin-retribution | raid 60 | 621.3 to 621.3 | +0.00% / +0.00% | 0.0 to 0.0 |
| priest-discipline | bare 20 | 29.3 to 29.3 | +0.00% / +0.00% | 29.0 to 29.0 |
| priest-discipline | bare 30 | 88.5 to 91.4 | +3.28% / +3.16% | 43.5 to 46.2 |
| priest-discipline | bare 40 | 142.3 to 157.0 | +10.33% / +10.33% | 47.6 to 55.5 |
| priest-discipline | bare 50 | 204.1 to 221.5 | +8.53% / +8.53% | 49.5 to 62.6 |
| priest-discipline | bare 60 | 285.2 to 294.7 | +3.33% / +3.47% | 53.6 to 62.6 |
| priest-discipline | raid 60 | 618.6 to 624.9 | +1.02% / +1.02% | 112.4 to 114.8 |
| priest-holy | bare 20 | 30.1 to 30.1 | +0.00% / +0.00% | 24.3 to 24.3 |
| priest-holy | bare 30 | 64.0 to 64.4 | +0.63% / +0.63% | 26.1 to 26.3 |
| priest-holy | bare 40 | 118.1 to 120.2 | +1.78% / +1.69% | 29.3 to 30.4 |
| priest-holy | bare 50 | 171.9 to 181.9 | +5.82% / +5.76% | 38.3 to 40.2 |
| priest-holy | bare 60 | 307.1 to 312.2 | +1.66% / +1.63% | 50.8 to 55.0 |
| priest-holy | raid 60 | 678.0 to 693.4 | +2.27% / +2.18% | 137.2 to 461.5 |
| priest-shadow | bare 20 | 35.3 to 35.3 | +0.00% / +0.00% | 90.2 to 90.2 |
| priest-shadow | bare 30 | 56.3 to 58.8 | +4.44% / +4.43% | 91.9 to 92.8 |
| priest-shadow | bare 40 | 123.6 to 129.4 | +4.69% / +4.77% | 70.8 to 58.7 |
| priest-shadow | bare 50 | 186.1 to 199.9 | +7.42% / +7.35% | 67.4 to 40.6 |
| priest-shadow | bare 60 | 343.4 to 355.3 | +3.47% / +3.49% | 19.0 to 7.6 |
| priest-shadow | raid 60 | 712.2 to 712.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-elemental | bare 20 | 22.2 to 22.2 | +0.00% / +0.00% | 92.6 to 92.6 |
| shaman-elemental | bare 30 | 43.4 to 47.1 | +8.53% / +8.51% | 95.8 to 74.8 |
| shaman-elemental | bare 40 | 77.1 to 93.8 | +21.66% / +21.95% | 101.7 to 42.4 |
| shaman-elemental | bare 50 | 100.8 to 151.8 | +50.60% / +50.79% | 98.8 to 1.4 |
| shaman-elemental | bare 60 | 132.9 to 190.6 | +43.42% / +43.45% | 89.3 to 2.0 |
| shaman-elemental | raid 60 | 464.2 to 466.9 | +0.58% / +0.56% | 5.5 to 0.3 |
| shaman-enhancement | bare 20 | 32.0 to 32.0 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | bare 30 | 63.2 to 63.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | bare 40 | 115.7 to 115.7 | +0.00% / +0.00% | 0.5 to 0.5 |
| shaman-enhancement | bare 50 | 154.8 to 154.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | bare 60 | 212.0 to 212.0 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-enhancement | raid 60 | 575.8 to 575.8 | +0.00% / +0.00% | 0.0 to 0.0 |
| shaman-restoration | bare 20 | 45.9 to 49.8 | +8.50% / +8.50% | 31.9 to 50.1 |
| shaman-restoration | bare 30 | 83.9 to 95.1 | +13.35% / +13.35% | 42.8 to 119.8 |
| shaman-restoration | bare 40 | 142.5 to 169.8 | +19.16% / +19.16% | 39.1 to 243.9 |
| shaman-restoration | bare 50 | 222.5 to 245.4 | +10.29% / +10.29% | 48.1 to 62.0 |
| shaman-restoration | bare 60 | 332.8 to 375.6 | +12.86% / +12.83% | 74.8 to 67.4 |
| shaman-restoration | raid 60 | 838.0 to 838.0 | +0.00% / +0.00% | 266.2 to 266.2 |
| warlock-affliction | bare 20 | 53.7 to 53.7 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 30 | 87.1 to 87.1 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 40 | 132.9 to 132.9 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 50 | 193.3 to 193.3 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-affliction | bare 60 | 286.9 to 286.9 | +0.00% / +0.00% | 0.7 to 0.7 |
| warlock-affliction | raid 60 | 573.2 to 573.2 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 20 | 47.7 to 47.7 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 30 | 79.5 to 79.5 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 40 | 140.7 to 140.7 | +0.00% / +0.00% | 0.0 to 0.0 |
| warlock-demonology | bare 50 | 206.7 to 206.7 | +0.00% / +0.00% | 2.0 to 2.0 |
| warlock-demonology | bare 60 | 323.4 to 323.4 | +0.00% / +0.00% | 3.8 to 3.8 |
| warlock-demonology | raid 60 | 642.6 to 642.6 | +0.00% / +0.00% | 0.3 to 0.3 |
| warlock-destruction | bare 20 | 52.3 to 52.3 | +0.00% / +0.00% | 3.1 to 3.1 |
| warlock-destruction | bare 30 | 87.7 to 88.2 | +0.57% / +0.57% | 2.8 to 1.8 |
| warlock-destruction | bare 40 | 138.7 to 138.7 | +0.00% / +0.00% | 3.5 to 3.5 |
| warlock-destruction | bare 50 | 234.6 to 234.6 | +0.00% / +0.00% | 2.0 to 2.0 |
| warlock-destruction | bare 60 | 351.9 to 351.9 | +0.00% / +0.00% | 4.5 to 4.5 |
| warlock-destruction | raid 60 | 638.9 to 638.9 | +0.00% / +0.00% | 0.4 to 0.4 |

## Per-rank verification on int/r9 09d7ffddb (each adopted rank does its own amount and costs its own mana)

Asked after the hybrids lane found lower Strength of Earth and Grace of Air ranks giving the top rank's amount (E 09d7ffddb). For every rank in rank-choices.json and the five raid-60 ranks, I checked two things.

1. **Engine source.** The rank has its own registration and its own aura, labelled "(Rank N)", never one aura shared by every rank. Its mana cost and base amount match the client row in spellconst/<class>.json:
   - Moonfire, Insect Swarm, Wrath.
   - Rejuvenation, which goes through `registerHealingRanks`.
   - Aimed Shot, Frostbolt, Scorch, Fire Blast.
   - Holy Light, Flash of Light, Consecration (an aura per rank).
   - Heal, Flash Heal, Renew, Power Word: Shield, Prayer of Healing and Binding Heal (`registerHealRanks`).
   - Lightwell: the renew is per rank, and its ticks are 160/233/320 every 2 s, matching SpellEffect 7001/27873/27874.
   - Shadow Word: Pain, Devouring Plague, Mind Blast, Shadow Word: Death.
   - Lightning Bolt, Flame Shock, Lava Burst, Healing Wave, Riptide, Shadow Bolt.
2. **Simulation.** A character with no gear (only a hunter's ranged weapon) ran single-spell probes, chosen rank against reference rank. Per-cast mana matches the client ratio exactly; for example, Flame Shock rank 1 against rank 4 is 0.220 against 0.220. Per-cast amount follows the client ratio within per-level growth.

In the full-rotation probes, a few mana ratios looked off: Frostbolt, Flame Shock and Lava Burst for Elemental, and Power Word: Shield. The cause is the rotations: Clearcasting-gated lines make some casts free more often. Single-spell probes match the client.

Result: every rank is ok, with no engine fix. Tools are in `scratchpad r9/ranks/verify/`.

## For other lanes

- **Integrator:**
  - **Ladder goldens:** every TestRotationLadder golden moves. The rules header text changed (resolver renamed), and so did the bands and ranks above. Regenerate at the pin.
  - **Addon commit:** 547d3f60f is the healers lane's a65bd8111, unchanged. Merging wt/r9-healers with wt/r9-ranks conflicts in data/tests/test_addonrotation.py; take wt/r9-ranks's side.
  - **Fork copies:** the fork copies of the five changed rotations sync in E 0a9ef0d0c.
- **Platform:** the /sim page can send `preset` (`"bare"` or `"raid"`) from its preset selector. Without it, a request falls back to the published preset for its level, which is the page's default; a user who switches a 60 to bare gets the raid ranks. The TypeScript envelope type is hand-written (web/src/lib/sim/types.ts).
- **Ranker:** sim/cmd/leveling-bis should set `api.SimRequest.Preset` to the pass's preset. Today its bare-60 pass falls back to the raid preset (`PublishedPreset(60)`), so the bare-60 rows do not reach the published bare-60 BiS until it does.
- **Hybrids:** Enhancement casts Strength of Earth Totem rank 1 at bare 60 for +1% with no mana shortage. A weaker Strength buff should not raise damage, so check what the lower rank changes.
- **Healers:** the heal profile is the level-60 fight at every band (owner question above).

## Still open

- A row is keyed by authored id. To split one spell's tank casts from its party casts, the two lines must name different ids, as Holy priest's Flash Heal and Holy paladin's Flash of Light now do.
