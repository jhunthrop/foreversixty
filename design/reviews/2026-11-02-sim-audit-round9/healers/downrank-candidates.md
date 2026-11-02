# Healer downrank candidates, round 9 (healers lane)

For the RANKS lane (owner ruling 3). Measured, not applied: this lane edits no rank.

How: every healer, every band, on its published entry (bare below 60, raid at 60), the heal profile's 300 s fight, effective HPS. 'Today' is what the site runs now: below 60 every ranked spell lifted to the highest rank learned (sim/request rotation_ranks.go); at 60 the list as written. Each candidate is a greedy search over every learned rank of every heal in the list, for all its casts and for the tank's and the party's casts apart, taking the best change while it gains at least 1% (3,000 iterations, seed 1), then confirmed at 10,000 iterations on seeds 1 and 50001 below. Ranks learned below 20 carry the sub-20 coefficient penalty. Tool: scratchpad r9/healers/downrank.py (a scratch spec-breakdown with the rewrite switched off).

Read this first: the heal profile is the level-60 Phase 1 fight at every band (2,128 tank hits on a 9,500-health tank, 555 pulses). Below 60 a bare healer covers 3 to 25% of it and nearly nothing overheals, so the best rank there is simply the most healing per mana; most bare healers run dry 30 to 50 s into the fight today (the druid is the exception). That is why some candidates are large.

Deliberate band-60 downranks the rule must keep (sim/request/healer_ranks_test.go `chosen`): Holy priest Renew rank 3 (6075); Discipline Renew rank 3 (6075) and Power Word: Shield rank 4 (3747); Holy paladin party Flash of Light rank 4 (19941) and Holy Shock rank 1 (1311606). Below 60 today's rewrite lifts every one of them to the band's top rank.

## Candidates (confirmed)

| spec | band | change | effective HPS today -> candidate, seed 1 | seed 50001 | mana lasts (s) today -> candidate, seed 1 |
|---|---|---|---|---|---|
| paladin-holy | 20 | none: no change beat today's ranks by 1% | - | - | - |
| paladin-holy | 30 | Flash of Light at rank 1 (19750, learned 20) on every cast | 73.1 -> 89.4 (+22.3%) | 73.1 -> 89.5 (+22.4%) | 29.3 -> 29.6 |
| paladin-holy | 40 | Holy Light at rank 2 (639, learned 6) on every cast | 114.7 -> 180.1 (+57.0%) | 114.7 -> 180.1 (+57.0%) | 35.3 -> 489.0 |
| paladin-holy | 50 | Holy Light at rank 3 (647, learned 14) on every cast; Flash of Light at rank 1 (19750, learned 20) on every cast | 145.0 -> 230.5 (+59.0%) | 144.9 -> 230.3 (+58.9%) | 42.6 -> 203.6 |
| paladin-holy | 60 | Flash of Light at rank 1 (19750, learned 20) on the tank's casts | 662.7 -> 669.7 (+1.1%) | 663.1 -> 669.4 (+1.0%) | 266.4 -> 281.5 |
| shaman-restoration | 20 | Healing Wave at rank 3 (547, learned 12) on every cast | 50.7 -> 55.4 (+9.3%) | 50.7 -> 55.4 (+9.3%) | 31.7 -> 49.4 |
| shaman-restoration | 30 | Healing Wave at rank 3 (547, learned 12) on every cast | 86.5 -> 98.3 (+13.6%) | 86.6 -> 98.3 (+13.5%) | 45.1 -> 129.0 |
| shaman-restoration | 40 | Healing Wave at rank 3 (547, learned 12) on every cast | 146.1 -> 174.8 (+19.6%) | 146.2 -> 174.8 (+19.6%) | 41.2 -> 259.5 |
| shaman-restoration | 50 | Healing Wave at rank 3 (547, learned 12) on every cast | 226.8 -> 250.0 (+10.2%) | 227.0 -> 249.9 (+10.1%) | 49.1 -> 63.0 |
| shaman-restoration | 60 | none: no change beat today's ranks by 1% | - | - | - |
| priest-discipline | 20 | none: no change beat today's ranks by 1% | - | - | - |
| priest-discipline | 30 | Renew at rank 1 (139, learned 8) on every cast | 82.7 -> 86.1 (+4.1%) | 82.7 -> 86.1 (+4.1%) | 33.3 -> 35.5 |
| priest-discipline | 40 | Renew at rank 1 (139, learned 8) on every cast; Binding Heal at rank 1 (401937, learned 25) on every cast | 135.9 -> 156.1 (+14.9%) | 135.9 -> 156.1 (+14.9%) | 37.9 -> 43.7 |
| priest-discipline | 50 | Binding Heal at rank 3 (1240771, learned 38) on every cast | 201.6 -> 213.8 (+6.1%) | 201.6 -> 213.7 (+6.0%) | 40.9 -> 42.5 |
| priest-discipline | 60 | Heal at rank 1 (2054, learned 16) on every cast | 619.6 -> 626.6 (+1.1%) | 619.5 -> 626.9 (+1.2%) | 109.7 -> 112.5 |
| priest-holy | 20 | none: no change beat today's ranks by 1% | - | - | - |
| priest-holy | 30 | Renew at rank 3 (6075, learned 20) on every cast | 69.4 -> 70.4 (+1.4%) | 69.4 -> 70.4 (+1.4%) | 28.5 -> 29.1 |
| priest-holy | 40 | Renew at rank 3 (6075, learned 20) on every cast; Heal at rank 2 (2055, learned 22) on the party's casts | 125.6 -> 129.3 (+2.9%) | 125.5 -> 129.4 (+3.1%) | 31.8 -> 33.1 |
| priest-holy | 50 | none: no change beat today's ranks by 1% | - | - | - |
| priest-holy | 60 | Flash Heal at rank 2 (9472, learned 26) on the tank's casts | 687.0 -> 696.6 (+1.4%) | 687.5 -> 696.7 (+1.3%) | 142.2 -> 446.3 |
| druid-restoration | 20 | Rejuvenation at rank 2 (1058, learned 10) on every cast | 42.3 -> 43.2 (+2.1%) | 42.3 -> 43.2 (+2.1%) | 256.2 -> 265.0 |
| druid-restoration | 30 | Rejuvenation at rank 2 (1058, learned 10) on every cast | 101.0 -> 110.9 (+9.8%) | 101.1 -> 110.8 (+9.6%) | 203.3 -> 231.5 |
| druid-restoration | 40 | none: no change beat today's ranks by 1% | - | - | - |
| druid-restoration | 50 | none: no change beat today's ranks by 1% | - | - | - |
| druid-restoration | 60 | Rejuvenation at rank 4 (2090, learned 22) on the party's casts | 779.6 -> 804.8 (+3.2%) | 779.2 -> 804.6 (+3.3%) | 287.4 -> 346.0 |

## Every single change (3,000 iterations, seed 1)

The best lower rank of each heal on its own, from today's ranks. A greedy step above is one of these.


### priest-holy, band 20 (bare): today 35.5 effective HPS, mana lasts 27.6 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Renew (rank 3 today) | rank 2 (6074, level 14) | all | 33.8 | -4.8% | 28.3 |

### priest-holy, band 30 (bare): today 69.4 effective HPS, mana lasts 28.5 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Renew (rank 4 today) | rank 3 (6075, level 20) | all | 70.4 | +1.4% | 29.1 |
| Flash Heal (rank 2 today) | rank 1 (2061, level 20) | all | 67.4 | -2.9% | 31.6 |
| Heal (rank 3 today) | rank 2 (2055, level 22) | party | 69.4 | +0.0% | 28.5 |

### priest-holy, band 40 (bare): today 125.5 effective HPS, mana lasts 31.8 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Renew (rank 6 today) | rank 3 (6075, level 20) | all | 127.4 | +1.5% | 33.1 |
| Flash Heal (rank 4 today) | rank 3 (9473, level 32) | party | 125.3 | -0.2% | 31.8 |
| Prayer of Healing (rank 2 today) | rank 1 (596, level 30) | all | 125.5 | +0.0% | 31.9 |
| Heal (rank 4 today) | rank 2 (2055, level 22) | party | 127.2 | +1.4% | 31.8 |

### priest-holy, band 50 (bare): today 187.4 effective HPS, mana lasts 43.7 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Prayer of Mending (rank 2 today) | rank 1 (401859, level 40) | all | 185.6 | -1.0% | 43.3 |
| Renew (rank 8 today) | rank 3 (6075, level 20) | all | 189.1 | +0.9% | 44.2 |
| Flash Heal (rank 6 today) | rank 5 (10915, level 44) | party | 183.6 | -2.0% | 43.7 |
| Prayer of Healing (rank 3 today) | rank 1 (596, level 30) | all | 187.5 | +0.1% | 44.3 |
| Heal (rank 4 today) | rank 3 (6063, level 28) | party | 185.6 | -1.0% | 43.7 |

### priest-holy, band 60 (raid): today 686.9 effective HPS, mana lasts 142.4 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Prayer of Mending (as the list writes it) | rank 2 (1240826, level 50) | all | 677.6 | -1.4% | 137.4 |
| Renew (as the list writes it) | rank 3 (6075, level 20) | all | 686.9 | +0.0% | 142.4 |
| Flash Heal (as the list writes it) | rank 2 (9472, level 26) | tank | 696.9 | +1.5% | 445.9 |
| Prayer of Healing (as the list writes it) | rank 1 (596, level 30) | all | 694.1 | +1.0% | 156.3 |
| Heal (as the list writes it) | rank 1 (2054, level 16) | party | 695.8 | +1.3% | 142.5 |

### priest-discipline, band 20 (bare): today 32.6 effective HPS, mana lasts 29.5 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Power Word: Shield (rank 3 today) | rank 1 (17, level 6) | all | 32.5 | -0.3% | 30.0 |
| Renew (rank 3 today) | rank 2 (6074, level 14) | all | 31.3 | -4.0% | 29.3 |

### priest-discipline, band 30 (bare): today 82.7 effective HPS, mana lasts 33.3 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Power Word: Shield (rank 5 today) | rank 1 (17, level 6) | all | 82.9 | +0.2% | 36.3 |
| Renew (rank 4 today) | rank 1 (139, level 8) | all | 86.1 | +4.1% | 35.5 |
| Flash Heal (rank 2 today) | rank 1 (2061, level 20) | party | 82.3 | -0.5% | 33.3 |
| Heal (rank 3 today) | rank 1 (2054, level 16) | all | 82.7 | +0.0% | 33.3 |

### priest-discipline, band 40 (bare): today 136.0 effective HPS, mana lasts 37.9 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Penance (rank 2 today) | rank 1 (402174, level 30) | all | 141.7 | +4.2% | 39.3 |
| Power Word: Shield (rank 6 today) | rank 2 (592, level 12) | all | 136.3 | +0.2% | 37.7 |
| Renew (rank 6 today) | rank 1 (139, level 8) | all | 147.3 | +8.3% | 40.3 |
| Binding Heal (rank 3 today) | rank 1 (401937, level 25) | all | 144.9 | +6.5% | 40.7 |
| Flash Heal (rank 4 today) | rank 3 (9473, level 32) | party | 135.8 | -0.1% | 37.9 |
| Prayer of Healing (rank 2 today) | rank 1 (596, level 30) | all | 136.1 | +0.1% | 38.2 |
| Heal (rank 4 today) | rank 3 (6063, level 28) | all | 136.0 | +0.0% | 37.9 |

### priest-discipline, band 50 (bare): today 201.5 effective HPS, mana lasts 40.9 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Penance (rank 3 today) | rank 1 (402174, level 30) | party | 201.2 | -0.1% | 40.9 |
| Power Word: Shield (rank 8 today) | rank 1 (17, level 6) | all | 201.5 | +0.0% | 41.1 |
| Renew (rank 8 today) | rank 3 (6075, level 20) | all | 207.7 | +3.1% | 42.1 |
| Binding Heal (rank 5 today) | rank 3 (1240771, level 38) | all | 213.7 | +6.1% | 42.5 |
| Flash Heal (rank 6 today) | rank 5 (10915, level 44) | party | 201.4 | -0.0% | 40.9 |
| Prayer of Healing (rank 3 today) | rank 1 (596, level 30) | all | 201.7 | +0.1% | 41.9 |
| Greater Heal (rank 2 today) | rank 1 (2060, level 40) | all | 201.0 | -0.2% | 41.4 |
| Heal (rank 4 today) | rank 3 (6063, level 28) | all | 191.5 | -5.0% | 40.9 |

### priest-discipline, band 60 (raid): today 619.3 effective HPS, mana lasts 109.8 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Penance (as the list writes it) | rank 1 (402174, level 30) | party | 620.9 | +0.3% | 109.6 |
| Power Word: Shield (as the list writes it) | rank 1 (17, level 6) | all | 620.8 | +0.2% | 110.8 |
| Renew (as the list writes it) | rank 3 (6075, level 20) | all | 619.3 | +0.0% | 109.8 |
| Binding Heal (as the list writes it) | rank 1 (401937, level 25) | party | 620.1 | +0.1% | 109.8 |
| Flash Heal (as the list writes it) | rank 1 (2061, level 20) | party | 619.5 | +0.0% | 109.8 |
| Prayer of Healing (as the list writes it) | rank 3 (10960, level 50) | all | 618.4 | -0.1% | 111.1 |
| Greater Heal (as the list writes it) | rank 3 (10964, level 52) | all | 616.3 | -0.5% | 110.3 |
| Heal (as the list writes it) | rank 1 (2054, level 16) | all | 626.3 | +1.1% | 112.4 |

### druid-restoration, band 20 (bare): today 42.3 effective HPS, mana lasts 256.2 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Touch (rank 4 today) | rank 3 (5187, level 14) | all | 42.0 | -0.7% | 266.8 |
| Regrowth (rank 2 today) | rank 1 (8936, level 12) | all | 40.4 | -4.5% | 272.5 |
| Rejuvenation (rank 3 today) | rank 2 (1058, level 10) | all | 43.1 | +1.9% | 265.2 |

### druid-restoration, band 30 (bare): today 101.0 effective HPS, mana lasts 203.5 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Touch (rank 5 today) | rank 4 (5188, level 20) | all | 101.0 | +0.0% | 223.9 |
| Regrowth (rank 4 today) | rank 3 (8939, level 24) | all | 99.9 | -1.1% | 215.2 |
| Rejuvenation (rank 5 today) | rank 2 (1058, level 10) | all | 110.9 | +9.8% | 233.7 |

### druid-restoration, band 40 (bare): today 177.4 effective HPS, mana lasts 214.4 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Touch (rank 7 today) | rank 6 (6778, level 32) | all | 176.6 | -0.5% | 226.2 |
| Tranquility (rank 2 today) | rank 1 (740, level 30) | all | 175.1 | -1.3% | 214.1 |
| Regrowth (rank 5 today) | rank 4 (8940, level 30) | all | 174.2 | -1.8% | 223.8 |
| Rejuvenation (rank 7 today) | rank 2 (1058, level 10) | party | 178.5 | +0.6% | 214.4 |

### druid-restoration, band 50 (bare): today 274.3 effective HPS, mana lasts 214.1 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Touch (rank 9 today) | rank 8 (9758, level 44) | all | 271.6 | -1.0% | 222.8 |
| Tranquility (rank 3 today) | rank 2 (8918, level 40) | all | 266.7 | -2.8% | 214.2 |
| Regrowth (rank 7 today) | rank 6 (9750, level 42) | all | 267.7 | -2.4% | 221.1 |
| Rejuvenation (rank 8 today) | rank 3 (1430, level 16) | party | 275.8 | +0.5% | 214.1 |

### druid-restoration, band 60 (raid): today 779.5 effective HPS, mana lasts 288.1 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Touch (as the list writes it) | rank 9 (9888, level 50) | all | 771.5 | -1.0% | 298.8 |
| Tranquility (as the list writes it) | rank 3 (9862, level 50) | all | 766.5 | -1.7% | 289.8 |
| Regrowth (as the list writes it) | rank 8 (9857, level 54) | all | 769.5 | -1.3% | 295.6 |
| Rejuvenation (as the list writes it) | rank 4 (2090, level 22) | party | 804.3 | +3.2% | 346.9 |

### paladin-holy, band 20 (bare): today 31.4 effective HPS, mana lasts 39.8 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Holy Light (rank 3 today) | rank 2 (639, level 6) | all | 31.1 | -1.0% | 79.6 |

### paladin-holy, band 30 (bare): today 73.1 effective HPS, mana lasts 29.2 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Holy Light (rank 5 today) | rank 3 (647, level 14) | all | 73.3 | +0.3% | 71.4 |
| Flash of Light (rank 2 today) | rank 1 (19750, level 20) | all | 89.4 | +22.3% | 29.6 |

### paladin-holy, band 40 (bare): today 114.6 effective HPS, mana lasts 35.3 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Holy Light (rank 6 today) | rank 2 (639, level 6) | all | 180.1 | +57.2% | 488.8 |
| Holy Shock (rank 2 today) | rank 1 (1311606, level 30) | all | 114.9 | +0.3% | 35.4 |
| Flash of Light (rank 3 today) | rank 1 (19750, level 20) | all | 162.7 | +42.0% | 36.4 |

### paladin-holy, band 50 (bare): today 144.9 effective HPS, mana lasts 42.7 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Holy Light (rank 7 today) | rank 3 (647, level 14) | all | 216.6 | +49.5% | 176.7 |
| Holy Shock (rank 3 today) | rank 2 (20473, level 40) | all | 144.9 | +0.0% | 42.7 |
| Flash of Light (rank 5 today) | rank 1 (19750, level 20) | all | 206.9 | +42.8% | 44.5 |

### paladin-holy, band 60 (raid): today 662.3 effective HPS, mana lasts 266.5 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Holy Light (as the list writes it) | rank 7 (10328, level 46) | all | 662.0 | -0.0% | 279.4 |
| Holy Shock (as the list writes it) | rank 1 (1311606, level 30) | all | 662.3 | +0.0% | 266.5 |
| Flash of Light (as the list writes it) | rank 1 (19750, level 20) | tank | 670.1 | +1.2% | 282.6 |

### shaman-restoration, band 20 (bare): today 50.7 effective HPS, mana lasts 31.7 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Wave (rank 4 today) | rank 3 (547, level 12) | all | 55.4 | +9.3% | 49.4 |

### shaman-restoration, band 30 (bare): today 86.6 effective HPS, mana lasts 45.0 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Wave (rank 5 today) | rank 3 (547, level 12) | all | 98.3 | +13.5% | 129.0 |
| Lesser Healing Wave (rank 2 today) | rank 1 (8004, level 20) | all | 86.7 | +0.1% | 45.4 |

### shaman-restoration, band 40 (bare): today 146.0 effective HPS, mana lasts 41.1 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Wave (rank 7 today) | rank 3 (547, level 12) | all | 174.8 | +19.7% | 259.8 |
| Lesser Healing Wave (rank 3 today) | rank 1 (8004, level 20) | all | 149.3 | +2.3% | 41.7 |

### shaman-restoration, band 50 (bare): today 226.9 effective HPS, mana lasts 49.1 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Wave (rank 8 today) | rank 3 (547, level 12) | all | 249.9 | +10.1% | 63.0 |
| Lesser Healing Wave (rank 4 today) | rank 1 (8004, level 20) | all | 228.2 | +0.6% | 54.6 |
| Riptide (rank 2 today) | rank 1 (408521, level 40) | party | 226.9 | +0.0% | 49.1 |
| Chain Heal (rank 2 today) | rank 1 (1064, level 40) | all | 226.5 | -0.2% | 49.7 |

### shaman-restoration, band 60 (raid): today 840.7 effective HPS, mana lasts 264.4 s of 300

| spell (site's rank today) | best other rank | casts | effective HPS | change | mana lasts (s) |
|---|---|---|---|---|---|
| Healing Wave (as the list writes it) | rank 8 (10395, level 48) | all | 837.8 | -0.3% | 266.5 |
| Lesser Healing Wave (as the list writes it) | rank 5 (10467, level 52) | party | 840.8 | +0.0% | 264.5 |
| Riptide (as the list writes it) | rank 2 (1239242, level 50) | party | 840.5 | -0.0% | 265.1 |
| Chain Heal (as the list writes it) | rank 2 (10622, level 46) | all | 803.9 | -4.4% | 264.2 |

