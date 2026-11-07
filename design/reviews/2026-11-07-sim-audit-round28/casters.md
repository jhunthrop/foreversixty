# Sim-accuracy round 28: casters lane

Site `wt/r28-casters` from int/r28 eb22d9382. Fork: **no engine change** (no `lane/r28-casters` branch was made; every number here runs on the pinned engine 978e4c8aa). Client 1.60.1.70790. Ship mode: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Result: 11 confirmed bugs, all fixed, all `old` (3 are round 27's X1 leads; none is a guides-handoff).** Each of the three leveling walks spent points on talents that do nothing in a fight while a damage talent waited, and the sweep found the same pattern deeper in each walk than round 27's lead said. The rank rows on the moved levels were judged again on the new talents, and three specs' rows changed (W4, D3, F4).

Method (round 27 orders.md): `spec-breakdown -preset bare -band <20|30|40> -level N -talents <walk string> -bis-dir $R/bis-frozen`, old walk against new walk at every level the change touches, 10,000 iterations, seeds 1 and 50001, both factions' published races. Candidate orders were first explored at 3,000 iterations (two or three hand orders, then a level-by-level greedy pass over every legal next point of the guide build, `$R/casters/tools/greedy.py`), then the chosen order checked at 10,000. A dead point is one whose talent the engine does not read for the fight (cited per row) or whose spell the rotation never casts; where a 3,000-iteration run gave a dead point and its alternative identical DPS to the last digit, that is recorded. Rank rows on the moved levels were then re-judged on the new walk (`-level N -race R -rank id=rank`, every legal race, a gain above 0.15% at both seeds, as round 24). Outputs: `$R/casters/final-{destro,demo,frost}.txt`, `rank-*.txt`, `rank2.txt`, `rank3.txt`, `out-*-explore*.txt`, `out-*-greedy*.txt`, walks in `$R/casters/w/`.

Logs, last looked 01:16 UTC 7 Nov: `GET https://api.foreversixty.gg/v1/reports/recent` still lists only `z2opt2x2gs5u` (20 Oct, no kills). A leveling order is not settled by raid logs; nothing here waits on one.

Web research (owner rule, for the ties): the Forever guide pages that load (search summaries; wowfdb, Icy Veins and Mobalytics pages answered 403 to a fetch, 7 Nov) agree with the direction: Icy Veins' Forever Destruction leveling build takes Destructive Reach at 21 and 22, after Bane (https://www.icy-veins.com/wow-forever/destruction-warlock-leveling), and Warcraft Tavern's Forever Destruction guide takes Bane to 5/5 first, then Aftermath and Shadowburn (https://www.warcrafttavern.com/forever/guides/destruction-warlock/). No Forever Frost or Demonology leveling order could be read; the sims decide those, and no step in them is a tie between two live points except where noted.

## Warlock: Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Leveling order: Destructive Reach at 10 and 11 and Molten Skin at 17 to 19 before Bane, Aftermath and Cataclysm.** Destructive Reach is range only (client 17917: "Increases the range of your Destruction spells by 10%") and Molten Skin is damage taken (1225220: "Reduces all damage taken by 2%"); the engine reads neither (fork `sim/warlock/talents.go:134-141`), and the bare fight's target never swings at you. Bane (cast time on Shadow Bolt and Immolate, cast from 1) waited until 12 to 16. Now Bane 10 to 14, Aftermath 15 to 17 (Immolate's first hit), Cataclysm 18 and 19. Aftermath first beat Cataclysm first at 15, 16 and 18 (3,000 iterations, within 0.2%; Cataclysm first won 17 by 0.1%), so Aftermath leads. +1.0 to +1.6% at 10, +3.0 to +3.4% at 11 to 14, +1.0 to +2.4% at 15 to 19, both seeds, both factions. Red first: `TestDestructionTakesBaneBeforeDestructiveReach`. | `final-destro.txt`; `out-destro-explore.txt` | CONFIRMED-BUG | old (walk order since the 70755 trees; round 27 X1 flagged it) | S 80bffc9fd |
| W2 | **The same walk at 20 to 29 spent Molten Skin and Destructive Reach while Shadowburn, Ruin and Agonizing Flames waited.** With the dead points moved, the tree reaches Shadowburn at 20 (one point beat Ruin's first by 0.35% at 20), Ruin 21 to 24, Agonizing Flames 25 to 27 (beat Ruin's fifth point by 1.2% at 25), Ruin 28. Destructive Reach's first point comes at 29, where every other option is dead or worse (Conflagrate, W3). +2.0 to +2.5% at 20 to 24, +3.2 to +3.5% at 25, +6.1 to +7.8% at 26 to 29. **Moves band 20 (level 20) and the level 25 checkpoint.** | `out-destro-greedy2.txt`, `final-destro.txt` | CONFIRMED-BUG | old | S 80bffc9fd |
| W3 | **Conflagrate alone loses damage, and the old walk took it at 34, three levels before Fire and Brimstone and five before Shadow and Flame.** Before Shadow and Flame every Conflagrate eats your Immolate (client 426316 is what makes it keep it). At level 34 on the old build, casting Conflagrate under the card's rule (Immolate at 6 seconds or less) is -0.46% against not casting it (3,000 iterations, 109.20 against 109.71; Immolate at 3 seconds or less ties, longer windows lose more), with one point of Fire and Brimstone it is +0.27%, with three +1.4%. So the talent point at 34 lost damage, and Shadow and Flame (a damage talent from 35) waited until 39 to 44 behind Bane of Havoc and Fire and Brimstone. Now Bane of Havoc at 30 (the dead toll Incinerate needs), Conflagrate 31, Fire and Brimstone 32 to 34, Shadow and Flame 35 to 38, Destructive Reach 39, Incinerate 40 (unchanged), Molten Skin 41 to 43, Shadow and Flame 5 at 44 (unchanged). Conflagrate at 31 is 1.1% behind spending 31 on a dead point and taking Shadow and Flame first with Conflagrate at 36 (order N, `w/destro-N.txt`), but order N trails by 0.4 to 6% from 33 to 42 (3,000 iterations); the shipped order beats the old one at every level: +5.3% at 30, +2.5% at 31, +0.4 to +0.6% at 33, +4.9% at 35, +14.1 to +14.3% at 38, +11.5 to +11.6% at 40, +3.7% at 42. **Moves band 30 and band 40.** The card's Conflagrate rows (cast at 6 seconds or less until Shadow and Flame, then on cooldown) stay: with Fire and Brimstone the 6-second window is the better rule. | `rot-*.json` runs (in `$R/casters`), `out-destro-explore3.txt`, `final-destro.txt` | CONFIRMED-BUG | old | S 80bffc9fd |
| W4 | **Rank rows on the new talents.** Shadow Bolt rank 4 at 28 was Troll-only; on the new order an Orc (+0.32 / +0.22%) and a Human (+0.18 / +0.18%) also gain on rank 4, and a Gnome (+0.30 / +0.12%) or Undead (+0.21 / +0.10%) ties, so the row is now every race. The Immolate rank 4 row at 40 (round 19) no longer earns anything: with Shadow and Flame 4 rank 5 wins for an Orc (+0.20 / +0.29%), Human (+0.24 / +0.21%) and Undead (+0.20 / +0.21%), and a Gnome or Troll ties; the row is deleted. Held: rank 2 at 20 to 23 (every race, -0.6 to -2.1% on ranks 3 and 4), rank 3 at 24 for a Gnome or Human (an Orc ties, +0.08%), rank 4 at 30 (rank 5 -0.15 to -0.33%) and 31 (tie), Corruption rank 3 at 41 (rank 4 -0.12 to -0.40%). Red first: `TestRound28CasterRankRows`; `TestRound24CasterRankRows` (Orc at 28) and `TestWarlockDestructionCheaperCorruptionAt41` (was `...DoTsAt40And41`) updated. | `rank-destro.txt`, `rank2.txt` | CONFIRMED-BUG | old (rows judged on the old walk) | S 4ffcc8988 |

## Warlock: Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Leveling order: Improved Health Funnel at 10 and 11 and Demonic Embrace at 12 to 14 before Unholy Power.** The rotation never casts Health Funnel (the engine registers none, fork `talents.go:90-93`), and Demonic Embrace is 3% Stamina a point (client 18697), which deals no damage. Unholy Power (2% pet damage a point; the demon is out in every band) waited until 15 to 19, and Fel Vitality (5% of your mana a point) until 22 to 24. Now Unholy Power 10 to 14, Fel Vitality 15 to 17, Demonic Energies 18, Improved Health Funnel 19 and 20, Demonic Embrace 21 to 23, Demonic Aegis 24, Improved Voidwalker 25, where the old order rejoins. Demonic Energies is as dead as Improved Voidwalker here: the level 24 builds that trade one for the other give 60.658 at every seed and faction, to the last digit, so its place at 18 is real leveling play (a free pet heal while you DoT). +0.14% at 10, +0.6 to +0.8% at 13 and 14, +0.6 to +1.1% at 15 to 18, +0.3 to +0.9% at 19 to 22, +0.2% at 23, level at 24. Small, but beyond error at every level (error about 0.04%). **Moves band 20.** Round 13's `TestDemonologyKeepsDemonicAegisAtTwenty` (levels 10 to 24 equal the plain walk) is replaced by `TestDemonologyKeepsDemonicSacrificeOutBelowTwentyFive`, which keeps its point (no Demonic Sacrifice below 25). Red first: `TestDemonologyTakesUnholyPowerBeforeHealthFunnel`. | `final-demo.txt`, `out-demo-explore.txt` | CONFIRMED-BUG | old (round 27 X1 flagged it) | S 80bffc9fd |
| D2 | **Decimation's second point at 31 before Demonic Knowledge.** Decimation only lowers Soul Fire's cooldown and starts its Soul Fire window, and the leveling rotation learns Soul Fire at 48 (round 13 said the same of the point at 31); Demonic Knowledge (spell damage while a demon is out) waited until 32 to 34. Now Demonic Knowledge 31 to 33, Decimation 34; 35 on unchanged. +3.1 to +3.3% at 31 to 33, both seeds, both factions. Band 30 (level 30) does not move. | `final-demo.txt` | CONFIRMED-BUG | old | S 80bffc9fd |
| D3 | **Rank row on the new talents.** With Fel Vitality from 15, a Gnome gains on Shadow Bolt rank 3 at 22 (+0.53 / +0.52%) and 23 (+0.29 / +0.19%); Human 22 +0.14 / +0.18% (not beyond 0.15% at seed 1), Orc, Troll and Undead tie or lose. New Gnome row at 22 and 23. Held: rank 3 at 20 and 21 (rank 4 -0.16 to -0.70%, Gnome and Troll). Red first: `TestRound28CasterRankRows`. | `rank-demo.txt`, `rank2.txt` | CONFIRMED-BUG | old | S 4ffcc8988 |
| D4 | Improved Voidwalker 1, Demonic Sacrifice 1, Master Summoner 2 and Decimation 1 at 25 to 29 do nothing in a fight either (the Succubus is out, no Pact, the summon is before the pull, no Soul Fire), but they are the 20-point toll Soul Link needs at 30, and the only live point left in those rows, Improved Sayaad, is not in the leveling build. Order-only; a build change (Master Summoner 2 into Improved Sayaad 2) is a build question for the talent adoption rule, not an order, and was not measured here. | walk print, client rows | OPEN (build lane) | - | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Leveling order: Frost Warding at 15 and 16 before Precision.** Frost Warding is armor and resistance on Frost Armor and Ice Armor and a stronger Frost Ward and Ice Barrier shield (client 11189); none of it deals damage. Precision (2% spell hit a point) waited until 17 to 19. Now Precision 15 to 17. Ice Shards first instead loses 1.4 to 5.0% at 17 to 21, so Precision leads. +2.4% at 15, +4.8 to +4.9% at 16 and 17. Red first: `TestFrostTakesPrecisionBeforeFrostWarding`; round 15's `TestMageLadderTakesItsCastTalentsEarly` level 20 case moved. | `final-frost.txt`, `out-frost-explore.txt` | CONFIRMED-BUG | old (round 27 X1 flagged it) | S 80bffc9fd |
| F2 | **Frost Channeling waited until 27 to 29 behind Frost Warding, Frostbite and Permafrost.** Frost Channeling is 5% off every Frost spell's cost a point (client 11160), and below 30 the bare fight runs short of mana. Its row opens at 10 points, which Ice Shards 2 at 18 and 19 reach, so it comes at 20 to 22, then Ice Shards 3 to 5 at 23 to 25. Taking Precision 2 and Ice Shards 3 first to the same row lost 1.8 to 1.9% at 18 and 19. +0.7% at 19, +2.5 to +2.8% at 20, +7.4 to +8.9% at 22 to 25. **Moves band 20.** | `out-frost-explore2.txt`, `final-frost.txt` | CONFIRMED-BUG | old | S 80bffc9fd |
| F3 | **Fingers of Frost came at 36, behind Ice Lance at 30, Ice Block at 31 and Cold Snap at 35.** The rotation casts Ice Lance only on a Fingers of Frost charge, so Ice Lance alone does nothing; Ice Block is never cast; Permafrost (chill duration, not read by the engine) and Frost Warding deal nothing. Fingers of Frost's row opens at 20 points and its spell at 30 (10 + 5 x tier 4). Now Frostbite 26 (a freeze on non-boss targets, `chill.go`), Shatter 27 and 28 (crit on frozen, which Frostbite makes), Ice Lance 29, Fingers of Frost 30 and 31, Shatter 32, then Frost Warding 33 and 34, Permafrost 35, Ice Block 36, Cold Snap 37, Winter's Chill 38 and 39, Ice Barrier 40 (unchanged). Frostbite alone at 26 ties (57.944 against 57.941, 3,000 iterations); Shatter after it gains 1.3% a point. +7.0 to +7.8% at 26, +2.4 to +2.7% at 29, +6.2 to +6.4% at 30, +11.2 to +14.6% at 31 to 35, +6.1 to +6.4% at 36. **Moves band 30.** Two card notes name the level and moved with it (Frostbolt filler "Fingers of Frost at 30", Ice Lance "From 31 ... two charges"). | `out-frost-greedy.txt`, `final-frost.txt` | CONFIRMED-BUG | old | S 80bffc9fd |
| F4 | **Frostbolt rank rows on the new talents.** Round 24's rows (rank 4 at 26 to 28 for all; 29 for a Human or Troll; 30 and 31 for a Human; rank 5 at 32 and 33 for all) no longer hold from 29. Every race, seeds 1 / 50001, rank 4 against rank 5: 29 Skyborne +1.63 / +1.62% (Orc ties, +0.10 / +0.07%); 30 Troll +0.45 / +0.33%, Skyborne +1.75 / +1.64% (Orc -0.16 / -0.22%); 31 Orc +0.92 / +0.64%, Undead +0.50 / +0.46%, Troll +1.45 / +1.17% (Gnome ties); 32 Human +1.54 / +1.28%, Orc +0.74 / +0.43%, Undead +0.25 / +0.24%, Troll +1.40 / +1.10% (Gnome -0.17 / -0.33%); 33 Human +1.01 / +0.71%, Troll +0.89 / +0.60%, Skyborne +1.64 / +1.32% (Orc ties, Undead and Gnome lose). Rank 6 loses 4 to 5% at 32 and 33 for every race tried and wins from 34 (rank 5 -0.8 to -5.3% at 34 to 36). New rows: rank 4 at 29 and 30 for Human, Troll and Skyborne; at 31 and 32 for every race but a Gnome; at 33 for Human, Troll and Skyborne; rank 5 at 32 and 33 for a Gnome, and at 33 for an Orc or Undead. 26 to 28 hold (rank 5 -0.2 to -5.0%). Red first: the Frost cases of `TestRound28CasterRankRows`. | `rank-frost.txt`, `rank2.txt`, `rank3.txt` | CONFIRMED-BUG | old | S 4ffcc8988 |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The ladder reservation (`ladder-first.json`) bypasses the own-tree spine, so a reservation that does not list the capstone loses it: Destruction without `Incinerate` in its list dropped Incinerate at 40 (as round 27 saw with Adrenaline Rush). Destruction lists Incinerate and Frost lists Ice Barrier for that reason. Recorded, not a bug. | walk print | Aligned | - | - |
| X2 | The climb (`point_order.go` `climbLadder`) takes the lowest tree position first when a level's walk adds two talents, so a list that puts Shadow and Flame before Conflagrate still gave Conflagrate first at 35. An order has to be written so each level's walk adds the wanted point and no other (Bane of Havoc listed before Destructive Reach's first point, for example). Recorded for whoever writes the next order. | walk print | Aligned | - | - |
| X3 | Fresh re-audit: limited to the three walks 10 to 44 (every point against its client row and the engine's reader) and every rank row on the moved levels. The full class re-audit (every spell's amounts, costs and procs at raid 60) was not repeated this round; raid 60 does not move (the walks are spent by 44 and the raid builds are untouched). | - | not done (scope) | - | - |

## Before/after

Bare, the band's frozen gear (`$R/bis-frozen`, band 20 to 29, 30 to 39, 40 to 49), the guide build's walk at the level, 10,000 iterations. Before = the round 27 walk, after = this lane's walk, both on engine 978e4c8aa and the round 27 rank rows (the rank-row changes of W4, D3 and F4 add their own gains on top; their numbers are in the rows). Absolute values at seed 1, change at seeds 1 / 50001. Horde is Orc (warlocks) or Troll (mage), Alliance is Gnome (the published races). Levels not listed are identical builds (Demonology 25 to 30 and 34 on; Frost 10 to 14 and 37 on; Destruction 43 on).

### warlock-destruction

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 10 | 16.566 | 16.740 | +1.05% / +1.04% | 15.944 | 16.197 | +1.59% / +1.54% |
| 11 | 17.652 | 18.260 | +3.44% / +3.42% | 16.549 | 17.107 | +3.37% / +3.40% |
| 12 | 19.719 | 20.235 | +2.62% / +2.73% | 19.181 | 19.687 | +2.64% / +2.63% |
| 13 | 23.370 | 24.030 | +2.82% / +2.85% | 20.179 | 20.776 | +2.96% / +2.81% |
| 14 | 25.605 | 26.448 | +3.29% / +3.17% | 22.568 | 23.244 | +3.00% / +3.07% |
| 15 | 30.287 | 31.018 | +2.41% / +2.37% | 27.136 | 27.737 | +2.21% / +2.16% |
| 16 | 34.702 | 35.056 | +1.02% / +1.02% | 36.439 | 36.809 | +1.02% / +1.02% |
| 17 | 38.107 | 38.675 | +1.49% / +1.49% | 38.303 | 38.853 | +1.44% / +1.44% |
| 18 | 41.548 | 42.330 | +1.88% / +1.94% | 41.493 | 42.335 | +2.03% / +2.07% |
| 19 | 44.117 | 44.979 | +1.95% / +2.09% | 43.886 | 44.942 | +2.41% / +2.35% |
| 20 | 60.337 | 61.782 | +2.39% / +2.52% | 61.296 | 62.710 | +2.31% / +2.39% |
| 21 | 61.201 | 62.741 | +2.52% / +2.59% | 62.159 | 63.648 | +2.40% / +2.42% |
| 22 | 62.156 | 63.648 | +2.40% / +2.48% | 63.118 | 64.510 | +2.21% / +2.22% |
| 23 | 63.161 | 64.560 | +2.21% / +2.24% | 64.132 | 65.471 | +2.09% / +2.19% |
| 24 | 66.413 | 67.692 | +1.93% / +2.07% | 67.216 | 68.572 | +2.02% / +1.98% |
| 25 | 67.367 | 69.694 | +3.45% / +3.50% | 68.145 | 70.313 | +3.18% / +3.22% |
| 26 | 70.474 | 74.947 | +6.35% / +6.28% | 71.602 | 75.946 | +6.07% / +6.14% |
| 27 | 71.187 | 76.687 | +7.73% / +7.75% | 72.323 | 77.699 | +7.43% / +7.50% |
| 28 | 73.704 | 78.958 | +7.13% / +7.25% | 74.740 | 79.854 | +6.84% / +6.94% |
| 29 | 74.721 | 79.805 | +6.80% / +6.93% | 75.582 | 80.702 | +6.77% / +6.78% |
| 30 | 100.220 | 105.553 | +5.32% / +5.32% | 98.976 | 104.284 | +5.36% / +5.36% |
| 31 | 102.706 | 105.282 | +2.51% / +2.58% | 101.428 | 103.971 | +2.51% / +2.53% |
| 32 | 106.006 | 107.625 | +1.53% / +1.41% | 104.797 | 106.239 | +1.38% / +1.48% |
| 33 | 108.935 | 109.412 | +0.44% / +0.43% | 107.494 | 108.061 | +0.53% / +0.58% |
| 34 | 109.215 | 110.893 | +1.54% / +1.54% | 107.876 | 109.518 | +1.52% / +1.55% |
| 35 | 110.191 | 115.548 | +4.86% / +4.89% | 108.829 | 114.156 | +4.89% / +4.90% |
| 36 | 116.708 | 126.287 | +8.21% / +8.26% | 115.282 | 124.638 | +8.12% / +8.18% |
| 37 | 118.254 | 131.283 | +11.02% / +11.17% | 116.687 | 129.681 | +11.14% / +11.24% |
| 38 | 122.165 | 139.355 | +14.07% / +14.24% | 120.725 | 137.731 | +14.09% / +14.28% |
| 39 | 127.198 | 140.190 | +10.21% / +10.38% | 125.639 | 138.742 | +10.43% / +10.30% |
| 40 | 169.795 | 189.399 | +11.55% / +11.60% | 168.400 | 187.821 | +11.53% / +11.52% |
| 41 | 178.218 | 191.910 | +7.68% / +7.75% | 176.898 | 190.490 | +7.68% / +7.75% |
| 42 | 186.524 | 193.409 | +3.69% / +3.76% | 184.854 | 191.622 | +3.66% / +3.73% |

### warlock-demonology

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 10 | 16.284 | 16.307 | +0.14% / +0.14% | 15.896 | 15.919 | +0.14% / +0.14% |
| 11 | 17.342 | 17.393 | +0.29% / +0.29% | 16.472 | 16.522 | +0.30% / +0.30% |
| 12 | 19.456 | 19.540 | +0.43% / +0.44% | 18.775 | 18.859 | +0.45% / +0.45% |
| 13 | 22.558 | 22.689 | +0.58% / +0.58% | 19.468 | 19.593 | +0.64% / +0.64% |
| 14 | 23.929 | 24.108 | +0.75% / +0.75% | 21.480 | 21.654 | +0.81% / +0.81% |
| 15 | 27.985 | 28.250 | +0.95% / +0.90% | 25.385 | 25.658 | +1.08% / +0.98% |
| 16 | 31.173 | 31.468 | +0.95% / +0.86% | 33.249 | 33.463 | +0.64% / +0.80% |
| 17 | 34.472 | 34.733 | +0.76% / +0.83% | 34.875 | 35.173 | +0.85% / +0.89% |
| 18 | 37.620 | 37.900 | +0.74% / +0.85% | 38.026 | 38.252 | +0.59% / +0.56% |
| 19 | 39.554 | 39.669 | +0.29% / +0.34% | 40.150 | 40.482 | +0.83% / +0.87% |
| 20 | 54.952 | 55.208 | +0.47% / +0.43% | 56.556 | 56.938 | +0.68% / +0.63% |
| 21 | 55.561 | 55.800 | +0.43% / +0.35% | 57.202 | 57.659 | +0.80% / +0.70% |
| 22 | 56.162 | 56.431 | +0.48% / +0.35% | 57.803 | 57.984 | +0.31% / +0.40% |
| 23 | 57.115 | 57.214 | +0.17% / +0.18% | 58.664 | 58.777 | +0.19% / +0.17% |
| 24 | 60.658 | 60.658 | +0.00% / +0.00% | 62.282 | 62.282 | +0.00% / +0.00% |
| 25 | 61.411 | 61.411 | same / same | 62.996 | 62.996 | same / same |
| 26 | 64.184 | 64.184 | same / same | 65.669 | 65.669 | same / same |
| 27 | 64.808 | 64.808 | same / same | 66.360 | 66.360 | same / same |
| 28 | 67.126 | 67.126 | same / same | 68.628 | 68.628 | same / same |
| 29 | 67.879 | 67.879 | same / same | 69.455 | 69.455 | same / same |
| 30 | 91.750 | 91.750 | same / same | 91.529 | 91.529 | same / same |
| 31 | 92.677 | 95.643 | +3.20% / +3.20% | 92.289 | 95.232 | +3.19% / +3.19% |
| 32 | 96.958 | 100.122 | +3.26% / +3.26% | 96.649 | 99.795 | +3.26% / +3.25% |
| 33 | 101.641 | 104.796 | +3.10% / +3.10% | 101.288 | 104.424 | +3.10% / +3.09% |

### mage-frost

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 15 | 20.891 | 21.391 | +2.39% / +2.49% | 21.004 | 21.510 | +2.41% / +2.39% |
| 16 | 25.819 | 27.061 | +4.81% / +4.80% | 26.713 | 28.014 | +4.87% / +4.89% |
| 17 | 29.293 | 30.678 | +4.73% / +4.75% | 30.247 | 31.697 | +4.79% / +4.65% |
| 18 | 30.978 | 31.802 | +2.66% / +2.63% | 32.245 | 33.123 | +2.72% / +2.67% |
| 19 | 32.906 | 33.147 | +0.73% / +0.74% | 34.355 | 34.618 | +0.77% / +0.77% |
| 20 | 46.087 | 47.358 | +2.76% / +2.78% | 48.957 | 50.180 | +2.50% / +2.52% |
| 21 | 46.758 | 49.246 | +5.32% / +5.32% | 49.587 | 52.338 | +5.55% / +5.59% |
| 22 | 47.961 | 51.678 | +7.75% / +7.78% | 50.792 | 55.324 | +8.92% / +8.90% |
| 23 | 51.270 | 55.420 | +8.09% / +8.06% | 54.609 | 58.658 | +7.41% / +7.41% |
| 24 | 52.478 | 56.642 | +7.93% / +7.95% | 55.956 | 59.851 | +6.96% / +6.95% |
| 25 | 52.916 | 57.292 | +8.27% / +8.10% | 56.145 | 60.434 | +7.64% / +7.49% |
| 26 | 53.679 | 57.867 | +7.80% / +7.79% | 56.952 | 60.934 | +6.99% / +6.87% |
| 27 | 55.581 | 59.053 | +6.25% / +6.20% | 58.895 | 62.061 | +5.38% / +5.29% |
| 28 | 58.434 | 61.010 | +4.41% / +4.42% | 61.262 | 63.719 | +4.01% / +3.98% |
| 29 | 59.673 | 61.276 | +2.69% / +2.70% | 62.966 | 64.470 | +2.39% / +2.37% |
| 30 | 80.085 | 85.128 | +6.30% / +6.23% | 83.304 | 88.594 | +6.35% / +6.30% |
| 31 | 80.599 | 89.675 | +11.26% / +11.24% | 83.694 | 93.123 | +11.27% / +11.33% |
| 32 | 82.137 | 94.047 | +14.50% / +14.47% | 85.177 | 97.572 | +14.55% / +14.51% |
| 33 | 83.641 | 94.634 | +13.14% / +13.16% | 86.654 | 98.105 | +13.21% / +13.25% |
| 34 | 93.015 | 104.433 | +12.28% / +12.29% | 96.831 | 108.651 | +12.21% / +12.13% |
| 35 | 94.092 | 105.652 | +12.29% / +12.33% | 97.874 | 109.850 | +12.24% / +12.22% |
| 36 | 100.844 | 107.129 | +6.23% / +6.39% | 104.800 | 111.225 | +6.13% / +6.23% |

Raid 60 and bare 50 / 60: unchanged for all three specs (every walk rejoins by 44; the raid builds are untouched). No spec outside warlock and mage moves.

## Files to regenerate

- `sim/request/testdata/ladder/` goldens (`TestRotationLadder`): warlock-destruction rungs 10, 20, 25, 30, 35, 38, 40; warlock-demonology rungs 10, 20 (31 to 33 sit between rungs); mage-frost rungs 20, 25, 30, 35; plus the rank-row moves at Destruction 28 and 40, Demonology 22 and 23 and Frost 29 to 33 where a rung sits on them (30, 40). Until then `web/src/lib/guides/leveling-builds.test.ts` fails on those rungs.
- `addon-builds.json` (the three leveling orders), `addon-data.json` rank rows, then `Data.lua`.
- `dps-per-point` / stat weights for the moved bands (Destruction 20, 30, 40; Demonology 20; Frost 20, 30).
- Coach starters (`persona_starter.py --all`) after addon-builds, and see "For other lanes" for a hand line in `tools/coach/starter_overrides.py`.
- **Fork copy of the Frost APL** (`ui/mage/apls/forever_frost.apl.json`): two card notes changed in `data/curated/apl/mage-frost.json`; `sim/request/apl/mage-frost.apl.json` was updated by hand to match. No fork commit was made here (no engine worktree).
- Hand copies done, not regens: `sim/leveling/ladder-first.json` = `data/curated/ladder-first.json`; `sim/rankchoices/rank-choices.json` = `data/curated/rank-choices.json`.

## Re-rank needed

- **warlock-destruction bands 20, 30 and 40** (W2, W3: every one of those levels' talents changes; band 40 by +11.5%).
- **warlock-demonology band 20** (D1).
- **mage-frost bands 20 and 30** (F2, F3; band 30 also changes its Frostbolt rank for Troll and Skyborne).

## Measure at refresh

- The six re-ranked bands above, on their new gear: the published band numbers move by roughly +2.4% (Destruction 20), +5.3% (Destruction 30), +11.5% (Destruction 40), +0.5% (Demonology 20), +2.6% (Frost 20) and +6.3% (Frost 30) before the gear moves.
- Rank rows on the re-ranked gear: Destruction Shadow Bolt 20 to 31, Demonology 20 to 23, Frost 26 to 33 were judged on frozen gear.

## For other lanes

- **Coach (do not hand-edit; regenerate or the coach owner):** `tools/coach/starter_overrides.py:201` says "Conflagrate from thirty-four"; the leveling build now takes it at 31.
- **Build lane:** D4, Demonology's leveling build pays Master Summoner 2 and Improved Voidwalker 1 (both do nothing with the Succubus out and the summon before the pull) where Improved Sayaad 2 is legal in the same rows; a build change, not an order, so the talent adoption rule applies.
- **Integrator:** sync the fork's `ui/mage/apls/forever_frost.apl.json` notes.

## For the guide lane

Done here, `_prose.test.ts` and `_sections.test.ts` pass: `warlock/destruction.md` (Levels 1 to 19 talent paragraph; Levels 20 to 60: Shadowburn from 20, Bane of Havoc 30, Conflagrate 31 with Fire and Brimstone, Shadow and Flame 35 to 38, Shadow Bolt rank 4 at 28 for everyone, the Immolate rank 4 line at 40 removed); `warlock/demonology.md` (Levels 1 to 19 talent paragraph); `mage/frost.md` (Levels 1 to 19 talent line; Levels 20 to 60: the talent sentence and the Frostbolt rank sentence by race). Two Frost card notes (Frostbolt filler, Ice Lance). No wow-player SHIP was run (build mode).

## Guide text not done

- None.

## Still open

- D4 (build lane). X3 (full class re-audit not repeated). Logs: still only the 20 Oct report. Owner pending items untouched.

## Sources

- Client 1.60.1.70790: `talents/warlock.json`, `talents/mage.json` (rank texts quoted per row), `combos.json` (the legal races), `hotfix-push.json` (no row for any talent here). Fork 978e4c8aa: `sim/warlock/talents.go` (Destructive Reach, Molten Skin, Improved Health Funnel, Demonic Aegis and Improved Voidwalker unread), `sim/mage/chill.go` (Frostbite freezes non-boss targets).
- Icy Veins, WoW Forever Destruction Warlock leveling (search summary, page 403 to a fetch, 7 Nov): https://www.icy-veins.com/wow-forever/destruction-warlock-leveling
- Warcraft Tavern, Forever Destruction Warlock guide (search summary, 7 Nov): https://www.warcrafttavern.com/forever/guides/destruction-warlock/
