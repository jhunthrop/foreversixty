# Sim-accuracy round 29: walks-heal lane

Site `wt/r29-walksheal` from int/r29 75d6264a5. Fork `lane/r29-walksheal` from int/r29 276ecc361, **unchanged** (no engine bug found; every number runs on the pinned engine). Client 1.60.1.70790. Ship mode: touched tests only, no regens. Every sim went through `$R/sim.sh`, one at a time.

**Result: 10 confirmed bugs, all fixed (9 `old`, 1 stale test `R28`).** All five healers' leveling orders spent points on talents that heal nothing in the band's fight, or on tolls they did not need yet, while live healing talents waited; the new orders heal more at every level they change (one Alliance cell at 41/42 for the Holy paladin and one at 20 for the Restoration shaman were closed by rank rows the new talents moved, H6 and S2). Survival's build spent Deterrence and Counterattack, which do nothing in the raid fight (seed item 2). The tank fight scaled a between-bands tank to its own level while it wears the lower band's gear (seed item 3). Demonology D4 (seed item 1) is aligned: Improved Sayaad does not beat Master Summoner under the adoption rule.

Method (round 27 orders.md, round 28 casters.md): `spec-breakdown -preset bare -band <band of the level> -level N -heal -talents <walk string> -bis-dir $R/bis-frozen`, judged on the ranker's heal score (effective HPS on the band's own heal profile, `sim/cmd/leveling-bis/score_heal.go`). The walks were explored with a level-by-level greedy over every legal next point (3,000 iterations, Horde seed 1, `tools/greedy.py`) and hand orders built from it, because the greedy alone defers tolls (Bloom, Holy Power, Prayer of Mending, Meditation, Nature's Splendor) that pay back a few levels later; every candidate was then compared with the old walk at every level 10 to 60 (3,000 iterations, Horde seed 1), and the chosen order at 10,000 iterations, seeds 1 and 50001, both factions' published races: every differing level for the druid, the Holy paladin and the Holy priest (10 to 44 plus checkpoints), checkpoints for Discipline and the shaman. Outputs: `$R/walksheal/final.tsv` (486 cells), `out-*-greedy*.txt`, `out-*-cmp*.txt`, walks in `$R/walksheal/w/`.

Healer sims are low-variance (error 0.01 to 0.3 HPS at 10,000 iterations), so "beyond error" means a gap above about 0.05%.

Logs, last looked 03:22 UTC 7 Nov: `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct, no kills). A leveling order is not settled by raid logs; nothing here waits on one.

## Druid: Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **Leveling order: Nature's Focus 3 to 5, five points of Naturalist and Tranquil Spirit took 21 to 39 while Gift of Nature (2% healing a point), Improved Rejuvenation, Swiftmend, Living Spirit and Genesis (10% on every tick, Balance tier 0) waited, and Nature's Splendor (3 s more Rejuvenation, 6 s more Regrowth) came last at 60.** Round 10's reservation of Nature's Swiftness "as soon as it can be" pulled all five Naturalist points (its prerequisite, client prereq rank 5) ahead of everything; measured now, Nature's Swiftness plus its toll is worth less than any of the waiting points at every level to 59. New order: Bloom 20, Nature's Focus 21, Genesis 22, Gift of Nature 23 to 25, Improved Rejuvenation 26, Swiftmend 27, Improved Rejuvenation 28 and 29, Gift of Nature 30, Living Spirit 31 to 33, Gift of Nature 34, Genesis 35 to 38, Nature's Majesty 39 and 40, Moonglow 41 to 43 (the toll), Nature's Splendor 44, Nature's Focus 45, Improved Regrowth 46, Nature's Focus 47, Improved Tranquility 48 and 49, Tranquil Spirit 50 to 54, Naturalist 55 to 59, Nature's Swiftness 60. 10 to 21 unchanged (Genesis at 10 to 14 beats Heart of the Wild by 0.5%, but delays Reflection and Bloom, which are worth 14% and 23%). Splendor at 44 needs Balance points ahead of the Restoration targets, which the ladder could not reserve (only first-row off-tree talents), so `reserveOffTree` now holds a deeper off-tree talent once the list has filled its rows (no other spec's walk changes; checked by printing all 56 walks before and after). Red first: `TestHealerWalksTakeTheirLiveTalentsFirst`, `TestReservationHoldsADeeperOffTreeTalentOnceItsRowsAre`. **Moves bands 30, 40, 50.** | `final.tsv`, `out-druid-greedy2.txt`, `out-druid-cmp1/2.txt`; client `talents/druid.json` | CONFIRMED-BUG | old (round 10's Nature's Swiftness reservation) | S 4690bfed2 |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Spiritual Focus 10 to 14: the only live first-row talent (Divine Strength is Strength); it beats Divine Intellect from Protection by 0.1% at 10. | `out-pal-greedy1` | Aligned | - | - |
| H2 | **Healing Light 15 to 17 before Reverence.** Reverence (30% regeneration while casting) first at 15 and 16, then Healing Light 17 to 19: +1.3 to +2.5% at 15 to 17, +10.4 to +10.8% at 18. | `final.tsv` | CONFIRMED-BUG | old | S 4690bfed2 |
| H3 | **Divine Favor at 21 before Illumination 2 to 4.** Illumination's refunds are worth more than a 2-minute sure crit at 21 to 24: Illumination 20 to 24, Divine Favor at 30. +1.5 to +2.7% at 21 to 25. | `final.tsv` | CONFIRMED-BUG | old (round 15 order) | S 4690bfed2 |
| H4 | **Purifying Power (Cleanse cost, Exorcism and Holy Wrath cooldowns) and Divine Precision (Holy spell hit; heals cannot miss) at 29 to 31, and Precision, Guardian's Favor and Sacred Duty (melee hit, Blessing of Protection and Divine Shield cooldowns) at 46 to 52, ahead of Divine Intellect (10% Intellect), Benediction and Conviction (6% crit, every Illumination refund).** None of the five does anything for a healer. New: Divine Intellect 26 to 29 and 31, Divine Favor 30, Benediction 32 to 36, Conviction 37 to 39, then the tolls (Purifying Power 40 and 41, Divine Precision 42, Infusion of Light 43 and 44), Holy Shock 45, Boundless Light 46 and 47, Holy Power 48 to 52, Light's Vigil 53, and Precision, Guardian's Favor, Sacred Duty 54 to 60. Holy Shock waits because, in a leveling fight, the paced Holy Shock line spends more mana than it returns: the talent alone loses 1.6% at 35 and 4 to 5% from 45 (`out-pal-greedy1`), and only pays back with Holy Power behind it. +0.2 to +4.8% at 26 to 40, +9.8 to +15.3% at 43 to 58, +4.3% at 59. The 7 dead Protection points stay in the build: every other legal point (Divine Strength, Improved Seals, Holy Conduit, Consecrated Ground, Deflection, the Judgement talents) is as dead, so there is no live replacement (H5). | `final.tsv`, `out-pal-cmp1..4.txt` | CONFIRMED-BUG | old | S 4690bfed2 |
| H5 | The level-60 build's 7 Protection points beyond Divine Intellect (Precision 3, Guardian's Favor 2, Sacred Duty 2) do nothing for a healer, but no legal point the build could take instead does anything either (the rotation casts no Judgement, Consecration, Exorcism or Holy Wrath; the Protection 10-point row is shield and armor). Not a bug under the adoption rule: a replacement must do something. | client rows | Aligned (no live replacement) | - | - |
| H6 | **Rank row on the new talents: the tank's Flash of Light rank 1 at 40 to 42.** With Conviction and Divine Intellect where Holy Shock and Holy Power were, rank 2 beats rank 1 by +1.6% (Undead 40) to +14.2% (Human 42), seeds 1 and 50001; Holy Light rank 6 instead of 4 ties and the group's rank 3 instead of 2 loses 4% on top of it (joint grid, `pal-joint.txt`), so only the tank row moves: rank 2 from 40 to 49. This also closes the one cell where the new order trailed (Human 42, -0.5% on the old row). Red first: `TestRound29HealerRankRows`; round 20's `TestHolyPaladinFlashOfLightAndHolyLightKeepTheirHealingBonus` updated. | `pal-tank2.txt`, `rank-paladin.txt` | CONFIRMED-BUG | old (row judged on the old order) | S 4690bfed2 |

## Priest: Holy

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Leveling order: Twilight Focus (pushback) 10 to 12 before Improved Renew and Holy Specialization; Divine Fury (Heal cast time) 21 to 25 before Holy Binding and Improved Healing; Discipline's Meditation last at 58 to 60.** New: Improved Renew 10 to 12, Holy Specialization 13 to 17, Divine Fury 18 and 19, Inspiration 20 to 22, Twilight Focus 23 and 24, Holy Binding 25 (the healer takes the raid pulse, so every Flash Heal on the group heals it too), Improved Healing 26 to 28, Twilight Focus 29, Litany of Light 30 and 31, Spiritual Guidance 32 and 33, Spirit of Redemption 34, Spiritual Healing 35 to 37, Spiritual Guidance 38 and 39, Prayer of Mending 40 (unchanged: all thirty Holy points are in, with Divine Fury's last three held back), Twin Disciplines 41 to 45, Mental Strength 46 to 50, Meditation 51 to 53, Inner Focus 54, Mental Agility 55 to 57, Divine Fury 58 to 60. +0.8 to +7.9% at 10 to 19, +0.05 to +0.7% at 20 to 24, +8 to +15.5% at 25 to 31, +1.1 to +4.6% at 32 to 52, +15 to +17% at 53, +9 to +10% at 58. One cell ties (22, Horde seed 1, +0.05%); every other cell gains beyond error. **Moves bands 20, 30, 40, 50.** | `final.tsv`, `out-ph-greedy1.txt`, `out-ph-cmp1/2.txt` | CONFIRMED-BUG | old | S 4690bfed2 |

## Priest: Discipline

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P2 | **Meditation's second and third points at 28 and 29, behind Mental Agility, Inner Focus and Improved Power Word: Shield.** Meditation 20 to 22: +6.4 to +15% at 21 to 28 (3,000 iterations), +12.7 to +13.3% at 22 (10,000, both seeds and factions). Penance stays at 30. | `final.tsv`, `out-pd-cmp1.txt` | CONFIRMED-BUG | old (round 13's order) | S 4690bfed2 |
| P3 | **After Penance: Soul Warding, Renewed Hope and Divine Aegis at 31 to 38 and Twilight Focus 39 to 41, ahead of Improved Renew and Holy Specialization; Divine Fury before Improved Healing.** New: Improved Renew 31 to 33, Soul Warding 34, Twilight Focus 35 to 37, Holy Specialization 38 to 42, Renewed Hope 43 to 45, Divine Aegis 46 to 48, Renewed Hope 49, Divine Fury 50 to 53, Improved Healing 54 to 56, Divine Fury 57, Inspiration 58 to 60. +0.1 to +0.6% at 31, +4.3 to +4.6% at 40, +5.7 to +6.1% at 56; 3,000-iteration sweep +0.2 to +4.9% at 31 to 48, +1.9 to +6.1% at 54 to 59. **Moves band 40.** | `final.tsv`, `out-pd-greedy2.txt` | CONFIRMED-BUG | old | S 4690bfed2 |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Leveling order: Improved Healing Wave 10 to 14 before Tidal Focus; Healing Focus (pushback) 28 to 30 and Tidal Mastery 31 to 35 before Restorative Totems and Healing Way; Riptide at 50.** A faster Healing Wave at 10 only spends the mana sooner: Improved Healing Wave's first point loses 15% to Tidal Focus's. New: Tidal Focus 10 to 12, Improved Healing Wave 13 and 14, Mindfulness 15 to 17, Tidal Focus 18 and 19, Water Shield 20, Improved Healing Wave 21 to 23, Ancestral Healing 24, Mana Tide Totem 25, Restorative Totems 26 to 29, Healing Way 30 to 32, Restorative Totems 33, Tidal Mastery 34, Purification 35 to 39, Riptide 40, Ancestral Knowledge 41 to 45, Thundering Strikes 46 to 50, Tidal Mastery 51 to 54, Nature's Swiftness 55, Ancestral Healing 56 and 57, Healing Focus 58 to 60. +15.9 to +17.8% at 10, +0.1 to +0.2% at 14, +14% at 30, +24% at 40, +3.5 to +3.7% at 50, +2.1 to +2.8% at 55, both seeds and factions; the 3,000-iteration sweep gains at every level but 23 to 25 (same build). The order Tidal Focus 10 to 14 first gains more at 15 to 18 but loses 0.06% at 14, so the split order was kept. **Moves bands 20, 30, 40, 50.** | `final.tsv`, `out-sh-cmp1/2.txt` | CONFIRMED-BUG | old | S 4690bfed2 |
| S2 | **Rank rows on the new talents: Healing Wave rank 3 at 18 and 19, rank 1 at 20 to 22.** With less Improved Healing Wave, rank 1 heals 10 to 12% more than rank 3 at 18 and 19 (Tauren and Dwarf), and rank 2 beats rank 1 by +0.5 to +1.1% at 20 to 22; at 22 a Dwarf (+0.12%) and a Windshaper (+0.08%) take rank 3, an Orc ties ranks 2 and 3, a Troll and a Tauren keep rank 2 (10,000 iterations, seeds 1 and 50001). This turns the one losing cell (Dwarf at 20, -0.2%) into +0.28%. Red first: `TestRound29HealerRankRows`; round 24's `TestRestorationShamanFillerHealingWaveByBand` and shaman r24 cases updated. | `shaman-hw.txt`, `shaman-hw-races.txt`, `shaman-old-18-22.txt` | CONFIRMED-BUG | old (rows judged on the old order) | S 4690bfed2 |

## Seed item 1: Demonology D4 (Master Summoner against Improved Sayaad)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | The leveling build's Master Summoner 2 into Improved Sayaad 2 (the plain walk then takes Improved Sayaad at 26 and 27): +0.14 to +0.42% at 27 and 28, a tie at 30 (+0.03 to +0.14%), and **-0.16 to -0.25% at 40 and 60**, -0.01 to -0.12% at 50, seeds 1 and 50001, both factions. Master Summoner is not dead in the sim: the Succubus is summoned before the pull and her mana (349 at 28, 554 at 40) comes out of the fight's bar, and Master Summoner cuts it 40%; Lash of Pain is about 3% of the damage, so Improved Sayaad's 20% on it is worth less from 40 on. Under the talent adoption rule (1% beyond error, both seeds) the change is refused. Leveling build unchanged. See W2 for the pre-pull mana. | `demo-isay.txt` | Aligned (refused under the rule) | - | - |
| W2 | The pre-pull summon's mana counts against the fight's mana bar. In play a warlock drinks after summoning, so whether that cost belongs in the fight is a model question for the warlock lane; it is the only thing Master Summoner does in the sim. The raid build also carries Master Summoner 2. | breakdown, `Summon Succubus (712) ResourceTypeMana -553.8` | OPEN (warlock lane) | - | - |

## Seed item 2: Survival's Deterrence and Counterattack

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| V1 | **The build (Survival's one build, raid and leveling) spent a point each on Deterrence and Counterattack, which do nothing in the raid fight.** The raid 60 build without either does the same damage to the last digit (868.16 Alliance seed 1, all four cells identical), so both are provably dead there: the boss swings at the tank and never at a hunter behind it, so no parry opens Counterattack and Deterrence is never pressed. Replacements tried (raid 60, frozen BiS, 10,000 iterations): Efficiency, Improved Stings and Improved Arcane Shot do nothing (identical); **Rapid Killing 2** (two minutes off Rapid Fire, which speeds the Survival hunter's melee swings in Forever, survival.md) gains +0.63 / +0.60% Alliance and +0.74 / +0.74% Horde, seeds 1 / 50001. Legal: Lethal Attacks and Careful Aim open Lone Wolf's row, and Survival keeps 32 points above Lacerating Strikes' 30. New build `5/005005201/500230230050222151`. The walk now takes Lightning Reflexes at 38 and 39 (+0.7% at 38, +1.3% at 40, +2.5 to +3.0% at 45 and 50, both seeds and factions), Rapid Killing at 58 and 59. The Counterattack rotation line goes (the build never has it), with its idle-line entry. Red first: `TestSurvivalBuildSpendsNothingOnDeterrenceOrCounterattack`; round 28's walk test moved at 38 to 40. **Moves bands 40, 50, bare 60 and raid 60.** | `survival-raid.txt`, `survival-walk.txt` | CONFIRMED-BUG | old (round 28 H6 left it open) | S 1e8cf0283 |
| V2 | `TestRound25...` remeasured Raptor Strike case "Orc at 27" still expected rank 3; round 28 H5 moved the Orc to rank 2 and did not update this case (red on int/r29 before this lane). | `go test -run Rank ./request/` | CONFIRMED (stale test) | R28 6cdbfb90d | S 1e8cf0283 |

## Seed item 3: the band 20 tank boss

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **A tank between two bands fought a boss scaled to its own level while wearing the lower band's gear.** `TankProfileForLevel` scaled the boss's damage by (level/60)^2 and the healing by (level/60)^1.5 at the character's level, so a level 28 tank in band 20 gear took (28/20)^2 = 1.96 times band 20's damage: bare at 25 and 28 the Protection paladin and the Feral bear died in 100% of fights (DTPS 147 and 188 against about 118 healing), while at 20 the paladin dies in 52%. The heal profile's own band rule says "a level between two bands fights the band at or below it" (data/curated/heal-profile.json bands.rule), because that is the gear it wears; the tank fight now follows it (below 20 a level keeps its own scaling: there is no band there). The boss's level (target +3) and attack power by level are unchanged, and band levels 20, 30, 40, 50, 60 fight exactly as before, so no published band number or BiS pick moves. Chance of death, before / after, 3,000 iterations, Horde seed 1: Protection paladin 25: 100 / 73.8%, 28: 100 / 74.3%, 35: 99.8 / 76.3%, 45: 94.0 / 50.1%; Feral bear 25: 100 / 65.1%, 28: 100 / 66.3%, 35: 100 / 99.8%, 45: 85.5 / 32.7%; Protection warrior 0% either way (Shield Block pushes crushing blows off the table). Threat between bands moves: paladin +4%, warrior -5 to -7%, bear -4 to -15% (less rage from damage taken). Red first: `TestTankProfileFightsTheBandAtOrBelowTheLevel`. | `tank-before.txt`, `tank-after.txt` | CONFIRMED-BUG | old | S ee02be261 |
| T2 | Below 60 the stand-in boss is still harder than the dungeon table the heal profile carries (cmangos last bosses, x1.2 client content modifier): its raw melee is 2.2x the band 20 bosses' median, 3.2x band 30's, 2.5x bands 40 and 50 and about 1.2x Onyxia's, and against the heal profile's sourced tank share (`tank_hit_damage` 183, 341, 600, 1289, 2645) it is 1.4 to 2x too hard from 20 to 50, most at 30 and 40. That is why a bare paladin or bear at 30 to 39 still dies in 76 to 100% of fights. The healing the tank fight assumes is about as much over a dungeon healer's tank healing (1.6 to 1.9x), so the damage-to-healing ratio is close and only spikes against health are too big. The fix is a model change (scale damage from the heal profile's band table and decide the healing to match), which reweights TMI for every tank band; left to the tank lane with these numbers. | heal-profile.json band rows; tank-encounter.json `level_scaling` reason ("A dungeon boss table would replace this") | OPEN (tank lane) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Rank rows on the new walks, screened at sample levels (3,000 iterations, Horde, one rank up and down, `rank-screen.txt`); flagged above 0.15% and not yet re-judged: druid party Rejuvenation rank 2 at 44 and 45 (+1.0 to +1.3%), tank Rejuvenation rank 3 at 25 (+0.3%), Healing Touch rank 2 at 30 and rank 3 at 35 (+0.3 to +0.4%), small cells at 49 and 57; Holy priest Renew rank 1 at 16 (+1.4%) and at 28 for an Undead (+3.7%, the race-keyed rows at 28 to 35 were judged on the old order); Discipline Renew rank 3 at 35 and rank 5 at 40 (+0.3 to +0.4%), Prayer of Healing rank 2 at 50 and 55 (+0.25 to +0.45%); shaman Healing Wave rank 5 at 40 (+0.3%). Each needs the level-by-level, race-by-race re-judge (round 24 method) on the new orders. | `rank-screen.txt` | OPEN (next round) | - | - |
| X2 | The fresh re-audit of every healer spell against the client (amounts, coefficients, costs) was not repeated this round; raid 60 does not move for any healer (the walks end on the unchanged guide builds). | - | not done (scope) | - | - |
| X3 | Guide-accuracy round 26 PALADIN S1 (row [3] at 18 to 24 "behind Divine Favor from 21"): Divine Favor now comes at 30, so the sim request changes shape. | guide-accuracy/round26/paladin.md | For the guide lane | - | - |

## Before/after

Bare, the band's frozen gear, the guide build's walk at the level, effective HPS on the band's heal fight, 10,000 iterations. Before = the round 28 walk and rows, after = this lane's walk and rows, both on engine 276ecc361. Absolute at seed 1, change at seeds 1 / 50001. Horde is Tauren (druid, shaman) or Undead (paladin, priests); Alliance is Night Elf, Human or Dwarf. Levels not listed are in `final.tsv`.

### druid-restoration

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 22 | 75.61 | 75.64 | +0.04% / +0.04% | 76.01 | 76.09 | +0.11% / +0.11% |
| 25 | 76.04 | 76.28 | +0.32% / +0.33% | 76.34 | 76.69 | +0.46% / +0.41% |
| 29 | 76.24 | 77.34 | +1.44% / +1.46% | 76.55 | 77.32 | +1.01% / +1.02% |
| 30 | 120.41 | 128.56 | +6.77% / +6.57% | 121.46 | 129.86 | +6.92% / +6.78% |
| 33 | 123.40 | 134.27 | +8.81% / +8.57% | 124.46 | 135.41 | +8.80% / +8.53% |
| 36 | 125.73 | 137.36 | +9.25% / +9.35% | 126.91 | 138.24 | +8.93% / +9.01% |
| 40 | 212.50 | 239.88 | +12.88% / +13.04% | 211.92 | 239.63 | +13.08% / +13.11% |
| 44 | 219.49 | 249.20 | +13.54% / +13.55% | 217.29 | 248.49 | +14.36% / +14.38% |
| 47 | 232.01 | 250.08 | +7.79% / +7.79% | 230.21 | 249.49 | +8.37% / +8.25% |
| 50 | 321.69 | 343.36 | +6.74% / +6.79% | 328.39 | 349.80 | +6.52% / +6.67% |
| 54 | 339.89 | 361.17 | +6.26% / +6.13% | 346.72 | 368.33 | +6.23% / +6.34% |
| 57 | 344.51 | 373.23 | +8.34% / +8.33% | 351.65 | 379.02 | +7.78% / +7.83% |
| 59 | 355.81 | 377.88 | +6.20% / +6.19% | 362.97 | 383.70 | +5.71% / +5.77% |

### paladin-holy (40 to 42 with H6's tank row)

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 15 | 21.00 | 21.30 | +1.43% / +1.33% | 21.26 | 21.63 | +1.74% / +1.69% |
| 18 | 37.95 | 42.01 | +10.70% / +10.76% | 37.96 | 41.91 | +10.41% / +10.35% |
| 21 | 50.47 | 51.65 | +2.34% / +2.38% | 50.67 | 52.03 | +2.68% / +2.72% |
| 25 | 74.79 | 76.23 | +1.93% / +1.82% | 75.75 | 77.10 | +1.78% / +1.74% |
| 28 | 77.84 | 78.20 | +0.46% / +0.46% | 78.24 | 78.46 | +0.28% / +0.31% |
| 30 | 125.35 | 128.31 | +2.36% / +2.29% | 131.12 | 133.77 | +2.02% / +2.07% |
| 32 | 127.18 | 132.12 | +3.88% / +3.90% | 131.83 | 137.35 | +4.19% / +4.23% |
| 36 | 139.13 | 141.71 | +1.85% / +1.83% | 141.93 | 143.11 | +0.83% / +0.83% |
| 39 | 135.78 | 142.05 | +4.62% / +4.63% | 139.30 | 142.06 | +1.98% / +1.98% |
| 40 | 191.09 | 200.29 | +4.81% / +4.79% | 195.95 | 208.25 | +6.28% / +6.21% |
| 42 | 182.54 | 204.87 | +12.23% / +12.12% | 187.19 | 212.73 | +13.64% / +13.50% |
| 43 | 179.74 | 207.29 | +15.33% / +15.34% | 189.03 | 211.64 | +11.96% / +11.90% |
| 47 | 173.02 | 194.02 | +12.14% / +12.08% | 181.55 | 199.34 | +9.80% / +9.84% |
| 50 | 197.02 | 223.12 | +13.25% / +13.37% | 206.32 | 234.11 | +13.47% / +13.45% |
| 53 | 188.35 | 216.96 | +15.19% / +15.03% | 196.69 | 226.71 | +15.26% / +15.15% |
| 56 | 194.64 | 222.66 | +14.40% / +14.44% | 203.45 | 232.84 | +14.45% / +14.50% |
| 59 | 219.82 | 229.83 | +4.55% / +4.42% | 230.15 | 240.15 | +4.34% / +4.36% |

### priest-holy

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 10 | 19.32 | 19.53 | +1.09% / +1.14% | 20.03 | 20.24 | +1.05% / +1.10% |
| 12 | 21.00 | 21.59 | +2.81% / +2.91% | 21.79 | 22.44 | +2.98% / +2.98% |
| 17 | 27.12 | 28.33 | +4.46% / +4.42% | 28.46 | 29.52 | +3.72% / +3.83% |
| 18 | 36.33 | 39.19 | +7.87% / +7.87% | 37.99 | 40.79 | +7.37% / +7.45% |
| 20 | 40.33 | 40.61 | +0.69% / +0.60% | 41.88 | 41.97 | +0.21% / +0.26% |
| 22 | 42.09 | 42.11 | +0.05% / +0.12% | 42.74 | 42.97 | +0.54% / +0.56% |
| 25 | 43.62 | 47.97 | +9.97% / +9.90% | 44.05 | 48.16 | +9.33% / +9.28% |
| 27 | 43.74 | 49.77 | +13.79% / +13.83% | 44.91 | 50.95 | +13.45% / +13.44% |
| 30 | 68.49 | 78.37 | +14.43% / +14.38% | 66.15 | 71.46 | +8.03% / +8.03% |
| 31 | 70.35 | 81.28 | +15.54% / +15.35% | 67.51 | 75.76 | +12.22% / +12.29% |
| 33 | 81.39 | 84.87 | +4.28% / +4.25% | 76.64 | 80.18 | +4.62% / +4.57% |
| 36 | 89.11 | 90.85 | +1.95% / +1.92% | 86.51 | 88.70 | +2.53% / +2.55% |
| 40 | 178.08 | 183.56 | +3.08% / +3.08% | 167.61 | 173.46 | +3.49% / +3.46% |
| 48 | 195.08 | 197.37 | +1.17% / +1.06% | 187.18 | 190.14 | +1.58% / +1.56% |
| 50 | 260.56 | 263.98 | +1.31% / +1.34% | 246.27 | 249.32 | +1.24% / +1.15% |
| 53 | 266.79 | 307.23 | +15.16% / +15.15% | 249.09 | 291.24 | +16.92% / +16.92% |
| 58 | 301.07 | 328.56 | +9.13% / +9.04% | 292.12 | 320.94 | +9.87% / +9.84% |

### priest-discipline

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 22 | 46.86 | 53.08 | +13.27% / +13.27% | 47.85 | 53.91 | +12.66% / +12.66% |
| 31 | 95.21 | 95.34 | +0.14% / +0.13% | 93.98 | 94.51 | +0.56% / +0.59% |
| 40 | 139.83 | 145.88 | +4.33% / +4.31% | 135.29 | 141.46 | +4.56% / +4.51% |
| 56 | 211.27 | 224.21 | +6.12% / +6.11% | 212.15 | 224.13 | +5.65% / +5.69% |

3,000-iteration sweep, Horde seed 1: +6.4 to +15.0% at 21 to 28, +0.2 to +4.9% at 31 to 48, +1.9 to +6.1% at 54 to 59, same build at 20, 29, 30 and 49 to 53.

### shaman-restoration (18 to 22 with S2's rows)

| level | Horde before (s1) | Horde after (s1) | Horde s1 / s50001 | Alliance before (s1) | Alliance after (s1) | Alliance s1 / s50001 |
|---|---|---|---|---|---|---|
| 10 | 11.16 | 13.15 | +17.83% / +17.83% | 12.85 | 14.89 | +15.88% / +15.88% |
| 14 | 17.33 | 17.36 | +0.17% / +0.23% | 18.49 | 18.51 | +0.11% / +0.11% |
| 18 | 35.24 | 43.46 | +23.3% / +23.4% | 35.21 | 43.51 | +23.6% / +23.6% |
| 20 | 71.57 | 72.63 | +1.48% / +1.51% | 72.60 | 72.80 | +0.28% / +0.30% |
| 22 | 71.26 | 72.69 | +2.01% / +2.02% | 72.57 | 72.95 | +0.52% / +0.50% |
| 30 | 115.37 | 131.63 | +14.09% / +14.11% | 116.06 | 132.24 | +13.94% / +13.98% |
| 40 | 162.91 | 201.95 | +23.96% / +23.99% | 163.72 | 202.84 | +23.89% / +23.96% |
| 50 | 249.70 | 258.53 | +3.54% / +3.59% | 250.25 | 259.51 | +3.70% / +3.67% |
| 55 | 270.13 | 276.05 | +2.19% / +2.13% | 269.42 | 276.89 | +2.77% / +2.81% |

### hunter-survival (DPS)

Raid 60, frozen BiS, 10,000 iterations: 868.16 to 873.62 Alliance seed 1 (+0.63 / +0.60%), 867.95 to 874.36 Horde seed 1 (+0.74 / +0.74%). Bare walk: 38 +0.68 to +0.69%, 40 +1.27 to +1.34%, 45 +2.82 to +3.00%, 50 +2.50 to +2.56%, both seeds and factions (`survival-walk.txt`).

### Tanks (T1)

Band levels unchanged. Between bands, chance of death and threat in T1 (3,000 iterations, Horde seed 1).

Raid 60 does not move for any healer. No spec outside the five healers, Survival and the three tanks (between bands only) moves.

## Files to regenerate

- `sim/request/testdata/ladder/` goldens (`TestRotationLadder`): druid-restoration rungs 30, 38, 40, 50; paladin-holy 30, 38, 40, 50 (40 also the tank Flash of Light row); priest-holy 10, 20, 30, 38, 40, 50; priest-discipline 38, 40; shaman-restoration 10, 20, 30, 38, 40, 50; hunter-survival 38, 40, 50, 60; the three tanks' rung 38 (T1, band 30 scaling). Until then `web/src/lib/guides/leveling-builds.test.ts` fails on the six walking specs (checked: exactly those six).
- `addon-builds.json` (six leveling orders), `addon-data.json` (card notes: Holy paladin Holy Light "at 30" and Holy Shock "From 46"; shaman Healing Wave ranks; Survival's Counterattack row gone; the paladin and shaman rank rows), then `Data.lua`.
- `dps-per-point` / stat weights: the healers' moved bands (druid 25 to 55, paladin 25 to 55, Holy priest 20 to 55, Discipline 25, 35 to 45, 55, shaman 20 to 55), Survival 40 to 60, and the tanks' mid-band levels 25, 35, 45, 55 (T1).
- Coach starters (`persona_starter.py --all`) after addon-builds.
- Fork APL copies (`apl-sync`): `ui/holy_paladin/apls/forever_holy.apl.json`, the Restoration shaman's and the Survival hunter's.
- Hand copies done, not regens: `sim/leveling/ladder-first.json`, `sim/rankchoices/rank-choices.json`, `sim/request/tank-encounter.json`, `sim/request/apl/{paladin-holy,shaman-restoration,hunter-survival}.apl.json`.

## Re-rank needed

- **hunter-survival bands 40, 50, bare 60 and raid 60 (LOUD: Survival is being re-ranked now on lane/bis-r28-rerank).** V1 changes the build itself (Rapid Killing 2 for Deterrence and Counterattack).
- druid-restoration bands 30, 40, 50.
- paladin-holy bands 30, 40, 50.
- priest-holy bands 20, 30, 40, 50.
- priest-discipline band 40.
- shaman-restoration bands 20, 30, 40, 50.
- Tanks: none (T1 leaves every band level's fight unchanged).
- Demonology: none (W1 refused).

## Measure at refresh

- The re-ranked healer bands: their published band HPS move before the gear does by about +6.8% (druid 30), +13% (druid 40), +6.7% (druid 50), +2.3% (paladin 30), +5% (paladin 40), +13.3% (paladin 50), +0.2 to +0.7% (Holy priest 20), +8 to +14% (Holy priest 30), +3.1 to +3.5% (Holy priest 40), +1.2 to +1.3% (Holy priest 50), +4.3 to +4.6% (Discipline 40), +0.3 to +1.5% (shaman 20), +14% (shaman 30), +24% (shaman 40), +3.6% (shaman 50).
- Survival raid 60 and bands 40, 50 on the re-ranked gear.
- X1's rank rows on the new orders, and every healer rank row between 10 and 59 that was judged on the old orders.
- Tank rank rows judged at between-band levels (the fight there changed, T1).

## For other lanes

- **Tank lane:** T2 (the stand-in boss is 1.4 to 2x the sourced dungeon table below 60; paladin and bear still die in 76 to 100% of bare fights at 30 to 39).
- **Warlock lane:** W2 (the pre-pull summon's mana comes out of the fight; the raid build's Master Summoner 2 rests on it).
- **Coach (generator lines, not touched):** every healer's and Survival's talent levels moved (e.g. Survival "Counterattack from thirty-nine" and "Deterrence from thirty-eight" no longer exist; Holy paladin Divine Favor now 30, Holy Shock 45).
- **Integrator:** the shared leveling change (off-tree reservations may hold deeper talents, Go and web) moves no other spec's walk (all 56 walks printed before and after).

## For the guide lane

Done here, `_prose.test.ts`, `_sections.test.ts` and `builds.test.ts` pass: druid/restoration.md (Levels 20 to 60, the talent order), paladin/holy.md (Levels 1 to 19 Reverence first; Levels 20 to 60 the order), priest/holy.md (Levels 1 to 19 and 20 to 60), priest/discipline.md (Levels 20 to 60), shaman/restoration.md (Levels 1 to 19 talents and Healing Wave rank at 18 and 19; Levels 20 to 60 talents, Healing Wave ranks at 20 to 22, Riptide at 40, Nature's Swiftness at 55), hunter/survival.md (Talents and builds: Rapid Killing; Rotation: Counterattack gone; Threat: Feign Death; Levels 20 to 60 order). Card notes: Holy paladin Holy Light ("Divine Favor once you have it at 30") and Holy Shock ("From 46 ... Boundless Light"); shaman Healing Wave ranks. No wow-player SHIP (build mode).

## Guide text not done

- None.

## Still open

- X1 (rank rows on the new orders), X2 (fresh re-audit), T2, W2. Owner pending items untouched.

## Sources

- Client 1.60.1.70790: `talents/{druid,paladin,priest,shaman,hunter,warlock}.json` (rank texts, prerequisites: Nature's Swiftness behind Naturalist 5, Nature's Splendor behind Nature's Majesty 2, Prayer of Mending behind Spirit of Redemption), `spellranks.json` (Riptide ranks 40, 50, 60; Healing Wave ranks).
- data/curated/heal-profile.json (band rule; cmangos classic-db last bosses; tank share and health per band) and data/curated/tank-encounter.json.
- hunter/survival.md: Rapid Fire speeds the Survival hunter's melee swings in Forever.
