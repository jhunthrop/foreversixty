# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-051500000000000000-000000000000000000)

Set DPS (verified): 35.1. Weights run: 5.7s. Verify run: 29.8s. 330 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.386 ± 0.036, crit=0.760 ± 0.038 per rating point (14 rating = 1%, 10.634 per %), hit=1.131 ± 0.086 per rating point (10 rating = 1%, 11.311 per %), melee_haste=6.829 ± 0.403

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.05 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.32 DPS) [crafted]; Brawler's Leather Hood (252504, -0.47 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.44 DPS) | yes | Erudite's Amulet (277204, -0.20 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.75 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.9 attack_power points (0.36 DPS) | yes | Rough Bronze Shoulders (3480, -0.05 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.05 DPS) [crafted] |
| back | Dark Leather Cloak (2316) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Lambent Scale Cloak (4706, -0.01 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.09 DPS, sim-verified) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.21 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.09 DPS, sim-verified) [quest]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Veteran's Gloves (250508, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.52 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Belt (252428, -0.23 DPS) [crafted]; Deviate Scale Belt (6468, -0.27 DPS) [crafted]; Cobrahn's Grasp (6460, -2.71 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.48 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Veteran's Boots (250503, -0.07 DPS) [crafted]; Feet of the Lynx (1121, -0.09 DPS, sim-verified) [world_drop]; Defender's Leather Boots (252441, -0.36 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.5 attack_power points (0.71 DPS) | yes | The 1 Ring (8350, -0.53 DPS) [world]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Ring of the Moon (12052, -0.61 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.42 DPS) | yes | The 1 Ring (8350, -0.24 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; Ring of the Moon (12052, -0.31 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.2 attack_power points (12.56 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Diamond Hammer (2194, -0.72 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.5 attack_power points (0.29 DPS) | yes | Dwarven Fishing Pole (3567, -0.08 DPS) [quest]; Fine Longbow (11304, -0.08 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.08 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Dark Leather Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 30 (human, 00000000000000000-051531023000100000-000000000000000000)

Set DPS (verified): 60.6. Weights run: 6.7s. Verify run: 11.6s. 568 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.262 ± 0.040, crit=0.939 ± 0.045 per rating point (14 rating = 1%, 13.142 per %), hit=1.182 ± 0.112 per rating point (10 rating = 1%, 11.821 per %), melee_haste=9.415 ± 0.815

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.4 attack_power points (1.82 DPS) | yes | Veteran's Chain Helm (250498, -0.33 DPS) [crafted]; Defender's Leather Helm (252455, -0.33 DPS) [crafted]; Tusken Helm (6686, -0.36 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.0 attack_power points (0.81 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Sentinel's Medallion (19541, -0.18 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.48 DPS) | yes | Barbaric Iron Shoulders (7913, -0.38 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.47 DPS) [crafted]; Golden Scale Shoulders (3841, -0.61 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (60.6 DPS) | yes | Sergeant Major's Cape (16315, -0.11 DPS) [pvp]; Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.30 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-verified (60.6 DPS) | yes | Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.15 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.37 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.99 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.36 DPS) | yes | Insignia Gloves (6408, -0.03 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.15 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.6 attack_power points (1.71 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Girdle of Golem Strength (9405, -0.22 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.8 attack_power points (1.91 DPS) | yes | Brawler's Leather Legguards (252516, -0.34 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.37 DPS) [world_drop]; Ferine Leggings (6690, -0.42 DPS, sim-verified) [dungeon] |
| feet | Alacritous Treads (277234) (or Brawler's Leather Boots (252439)) | WANTED: Incinerator Gar'im [quest] | 16.3 attack_power points (1.01 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (1.23 DPS) | yes | Silverlaine's Family Seal (6321, -0.61 DPS) [dungeon]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.6 attack_power points (1.21 DPS) | yes | Silverlaine's Family Seal (6321, -0.59 DPS) [dungeon]; Tiger Band (6749, -0.59 DPS) [quest]; Bonebinder's Signet (273032, -0.59 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-verified (60.6 DPS) | yes | Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop]; Thornspike (6681, -2.95 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (19.73 DPS) | yes | Zealot Blade (13033, -2.52 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -5.34 DPS) [quest] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.8 attack_power points (0.61 DPS) | yes | Gravespike Repeater (273033, -0.11 DPS) [world]; Golemsight Long Gun (273029, -0.14 DPS) [dungeon]; Long Battle Bow (15284, -0.23 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Alacritous Treads; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 40 (human, 00000000000000000-051531025120100221-000000000000000000)

Set DPS (verified): 143.2. Weights run: 6.7s. Verify run: 15.2s. 777 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.583 ± 0.060, crit=1.532 ± 0.077 per rating point (14 rating = 1%, 21.442 per %), hit=1.830 ± 0.128 per rating point (10 rating = 1%, 18.301 per %), melee_haste=9.575 ± 1.298

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.4 attack_power points (5.23 DPS) | yes | Chromite Barbute (8142, -0.49 DPS) [dungeon]; White Bandit Mask (10008, -0.88 DPS) [crafted]; Barbaric Iron Helm (7915, -1.67 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 17.4 attack_power points (1.92 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; River Pride Choker (13087, -1.04 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.4 attack_power points (3.02 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.60 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.65 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.49 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.66 DPS) [quest]; Parachute Cloak (10518, -0.97 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kolkar Marauder Chain (6773, -1.90 DPS) [quest]; Carapace of Tuten'kash (10775, -2.60 DPS) [dungeon]; Jouster's Chestplate (8157, -2.90 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Bracers (14590, -0.28 DPS) [world_drop]; Ravager's Armguards (14770, -0.33 DPS, sim-verified) [world_drop]; Yorgen Bracers (13012, -0.36 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.36 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.2 attack_power points (3.55 DPS) | yes | Boar Champion's Belt (10768, -0.25 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.25 DPS) [rep]; Highlander's Plate Girdle (20125, -0.25 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (4.63 DPS) | yes | Symbolic Legplates (14829, -0.28 DPS) [world_drop]; Firemane Leggings (13129, -0.44 DPS) [world_drop]; Triprunner Dungarees (9624, -0.83 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.1 attack_power points (4.09 DPS) | yes | Blackforge Greaves (6423, -0.36 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.44 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.7 attack_power points (3.16 DPS) | yes | Falcon's Hook (7552, -0.93 DPS) [dungeon]; Mark of Kern (2262, -0.95 DPS) [dungeon]; Assault Band (13095, -0.95 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.7 attack_power points (2.29 DPS) | yes | Falcon's Hook (7552, -0.06 DPS) [dungeon]; Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Ardent Custodian (868, -1.95 DPS, sim-verified) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (48.70 DPS) | yes | Southsea Lamp (9359, +0.00 DPS) [world_drop]; Shoni's Disarming Tool (9608, -23.12 DPS) [quest] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.27 DPS) [quest]; The Silencer (13138, -0.30 DPS) [world_drop]; Swiftwind (13038, -0.62 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Thunderbrow Ring; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 50 (human, 30321000000000000-051531025120100231-000000000000000000)

Set DPS (verified): 220.1. Weights run: 7.1s. Verify run: 28.9s. 982 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.434 ± 0.085, crit=1.691 ± 0.117 per rating point (14 rating = 1%, 23.679 per %), hit=2.018 ± 0.187 per rating point (10 rating = 1%, 20.177 per %), melee_haste=14.141 ± 1.722

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.7 attack_power points (5.32 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.64 DPS) [crafted]; Chromite Barbute (8142, -0.91 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.6 attack_power points (3.07 DPS) | yes | Sentinel's Medallion (19540, -1.38 DPS) [rep]; Kaleidoscope Chain (13084, -1.60 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 37.5 attack_power points (4.01 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.21 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.27 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16336, -0.10 DPS) [pvp]; Pridelord Cape (14673, -0.28 DPS) [world_drop]; Bloodlust Cape (14801, -0.38 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight's Plate Hauberk (220794, -0.10 DPS) [vendor]; Warforged Chestplate (11195, -0.71 DPS) [quest]; Warbear Harness (15064, -0.73 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.6 attack_power points (3.28 DPS) | yes | Deepfury Bracers (13120, -0.12 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.27 DPS) [crafted]; Arena Bands (18711, -0.28 DPS) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -4.08 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 45.7 attack_power points (4.89 DPS) | yes | Prowler's Leather Waistguard (252473, -0.05 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.17 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.21 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 62.0 attack_power points (6.64 DPS) | yes | Centurion Legplates (10740, -1.13 DPS) [quest]; Serpentskin Leggings (8262, -1.40 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.58 DPS, sim-verified) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 41.8 attack_power points (4.47 DPS) | yes | Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Officer's Sabatons (250561, -0.31 DPS) [crafted]; Officer's Boots (250546, -0.61 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.2 attack_power points (4.30 DPS) | yes | Thunderbrow Ring (13097, -2.13 DPS) [world_drop]; Masons Fraternity Ring (9533, -2.15 DPS) [quest]; Mark of Kern (2262, -2.16 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 32.9 attack_power points (3.52 DPS) | yes | Thunderbrow Ring (13097, -1.35 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.37 DPS) [quest]; Mark of Kern (2262, -1.38 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -2.69 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -0.65 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Flurry Axe (871, -4.03 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 533.7 attack_power points (57.17 DPS) | yes | Thorium Cestus (250614, -2.65 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.51 DPS) [dungeon]; Shoni's Disarming Tool (9608, -32.30 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.01 DPS) [world_drop]; Houndmaster's Bow (11628, -0.22 DPS) [dungeon]; Guttbuster (13139, -0.28 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Officer's Wristguards; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 982, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 60 (human, 30323013002000000-051531025120100231-200000000000000000)

Set DPS (verified): 329.0. Weights run: 7.4s. Verify run: 102.6s. 1833 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.192 ± 0.081, crit=1.702 ± 0.116 per rating point (14 rating = 1%, 23.833 per %), hit=2.296 ± 0.196 per rating point (10 rating = 1%, 22.957 per %), melee_haste=16.583 ± 1.993

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 129.6 attack_power points (19.61 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -6.17 DPS) [vendor]; Stalwart Helm (250599, -9.84 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (3.86 DPS) | yes | Stormpike Soldier's Pendant (19097, -1.08 DPS, sim-verified) [rep]; Sentinel's Medallion (19538, -1.15 DPS) [rep]; Sunprism Pendant (20649, -1.69 DPS) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 56.3 attack_power points (8.51 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.32 DPS) [dungeon]; Sergeant Major's Cape (16337, -2.18 DPS) [pvp]; Stormpike Soldier's Cloak (19084, -2.42 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -3.09 DPS) [crafted]; Savage Gladiator Chain (11726, -4.54 DPS, sim-verified) [dungeon]; Field Marshal's Plate Armor (231536, -5.87 DPS) [pvp] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 47.5 attack_power points (7.19 DPS) | yes | Forest Stalker's Bracers (19587, -0.44 DPS) [rep]; Marshal's Plate Bracers (16481, -0.79 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gauntlets (271940, -1.54 DPS, sim-verified) [rep] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 58.7 attack_power points (8.88 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Plate Girdle (20041, -0.13 DPS) [rep]; Highlander's Chain Girdle (20043, -0.13 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -1.86 DPS) [vendor]; Marshal's Plate Legguards (231540, -2.12 DPS) [pvp]; Sentinel's Lamellar Legguards (22753, -5.42 DPS) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 47.5 attack_power points (7.19 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Highlander's Plate Greaves (20048, -0.79 DPS) [rep]; Highlander's Lamellar Greaves (20049, -0.79 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 60.8 attack_power points (9.20 DPS) | yes | Protector's Band (19514, -3.58 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.0 attack_power points (6.50 DPS) | yes | Protector's Band (19516, -1.85 DPS) [rep] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -5.50 DPS, sim-verified) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -1.67 DPS, sim-verified) [crafted] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -0.82 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -30.96 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.99 DPS) [world_drop]; Houndmaster's Bow (11628, -1.29 DPS) [dungeon]; Bloodseeker (19107, -3.07 DPS, sim-verified) [quest] |

**New at 60:** head: Lionheart Helm; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; waist: Might of the Timbermaw; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: Ironwood Blade; off_hand: Persuader; ranged: Blackcrow

No-known-source sample (15 of 1833, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 60, raid preset (human, 30323013002000000-051531025120100231-200000000000000000)

Set DPS (verified): 875.9. Weights run: 8.2s. Verify run: 117.4s. 1833 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.108 ± 0.084, crit=1.417 ± 0.117 per rating point (14 rating = 1%, 19.836 per %), hit=2.594 ± 0.226 per rating point (10 rating = 1%, 25.943 per %), melee_haste=20.340 ± 2.446

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 131.2 attack_power points (37.69 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -11.26 DPS) [vendor]; Stalwart Helm (250599, -18.52 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.4 attack_power points (7.30 DPS) | yes | Sentinel's Medallion (19538, -2.52 DPS) [rep]; Sunprism Pendant (20649, -3.48 DPS) [quest]; Stormpike Soldier's Pendant (19097, -6.13 DPS, sim-verified) [rep] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 58.4 attack_power points (16.79 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -0.35 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.84 DPS) [dungeon]; Sergeant Major's Cape (16337, -3.76 DPS) [pvp]; Cloak of Flames (3475, -3.91 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -7.35 DPS) [crafted]; Savage Gladiator Chain (11726, -9.36 DPS, sim-verified) [dungeon]; Field Marshal's Plate Armor (231536, -12.02 DPS) [pvp] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 50.7 attack_power points (14.56 DPS) | yes | Forest Stalker's Bracers (19587, -1.55 DPS) [rep]; Marshal's Plate Bracers (16481, -1.58 DPS) [pvp] |
| hands | Theramore Gauntlets (271940) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Plate Gauntlets (231541, -2.23 DPS) [pvp]; Timbermaw Brawlers (19049, -2.84 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 61.7 attack_power points (17.73 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Plate Girdle (20041, -1.29 DPS) [rep]; Brigam Girdle (13142, -1.43 DPS) [dungeon] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (+10.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Plate Legguards (231540, -4.42 DPS) [pvp]; Sentinel's Chain Leggings (237819, -5.92 DPS) [vendor]; Sentinel's Lamellar Legguards (22753, -9.49 DPS) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 46.9 attack_power points (13.47 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Highlander's Plate Greaves (20048, -0.79 DPS) [rep]; Highlander's Lamellar Greaves (20049, -0.79 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 59.8 attack_power points (17.18 DPS) | yes | Protector's Band (19514, -6.09 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.9 attack_power points (13.20 DPS) | yes | Protector's Band (19514, -2.11 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -5.20 DPS, sim-verified) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -4.15 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -58.49 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -2.19 DPS, sim-verified) [quest]; Skull Splitting Crossbow (13039, -2.49 DPS) [world_drop]; Houndmaster's Bow (11628, -3.07 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Theramore Gauntlets; waist: Might of the Timbermaw; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ironwood Blade; off_hand: Persuader; ranged: Blackcrow

No-known-source sample (15 of 1833, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

## Horde

### Band 20 (troll, 00000000000000000-051500000000000000-000000000000000000)

Set DPS (verified): 35.3. Weights run: 5.7s. Verify run: 28.6s. 313 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.042 ± 0.024, crit=0.571 ± 0.026 per rating point (14 rating = 1%, 7.999 per %), hit=0.816 ± 0.052 per rating point (10 rating = 1%, 8.163 per %), melee_haste=5.698 ± 0.319

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.13 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.54 DPS) [crafted]; Brawler's Leather Hood (252504, -0.66 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.3 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.31 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.34 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.04 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.10 DPS) [dungeon]; Subterranean Cape (14149, -0.11 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Veteran's Chain Shirt (250488, -0.09 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.15 DPS) [crafted]; Defender's Leather Armor (252434, -0.16 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.56 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.16 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.23 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.3 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.47 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.01 DPS) | yes | Brawler's Leather Belt (252428, -0.33 DPS) [crafted]; Deviate Scale Belt (6468, -0.38 DPS) [crafted]; Cobrahn's Grasp (6460, -2.62 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.3 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.44 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (35.3 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -2.23 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.2 attack_power points (0.69 DPS) | yes | Loop of Sacrifice (281673, -0.35 DPS) [quest]; Bounty Hunter's Ring (5351, -0.51 DPS) [quest]; The 1 Ring (8350, -0.51 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.45 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Bounty Hunter's Ring (5351, -0.27 DPS) [quest]; The 1 Ring (8350, -0.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.6 attack_power points (13.44 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (13.12 DPS) | yes | Cruel Barb (5191, +0.00 DPS) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 attack_power points (0.23 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.01 DPS) [crafted]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (troll, 00000000000000000-051531023000100000-000000000000000000)

Set DPS (verified): 60.6. Weights run: 6.7s. Verify run: 12.5s. 553 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.262 ± 0.040, crit=0.939 ± 0.045 per rating point (14 rating = 1%, 13.142 per %), hit=1.182 ± 0.112 per rating point (10 rating = 1%, 11.821 per %), melee_haste=9.415 ± 0.815

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.4 attack_power points (1.82 DPS) | yes | Tusken Helm (6686, -0.21 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.33 DPS) [crafted]; Defender's Leather Helm (252455, -0.33 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.0 attack_power points (0.81 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Scout's Medallion (19537, -0.18 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.48 DPS) | yes | Barbaric Iron Shoulders (7913, -0.27 DPS) [crafted]; Barbaric Shoulders (5964, -0.47 DPS) [crafted]; Golden Scale Shoulders (3841, -0.61 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (60.6 DPS) | yes | Tigerstrike Mantle (13108, -0.27 DPS, sim-verified) [world_drop]; Wildhunter Cloak (16658, -0.30 DPS) [quest]; Construct Cloak (279848, -0.30 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-verified (60.6 DPS) | yes | Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.15 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.37 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.99 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.36 DPS) | yes | Insignia Gloves (6408, -0.03 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Warsong Gauntlets (16978, -0.12 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.6 attack_power points (1.71 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Girdle of Golem Strength (9405, -0.22 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.8 attack_power points (1.91 DPS) | yes | Ferine Leggings (6690, -0.30 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.34 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.37 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.3 attack_power points (1.01 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Hard Gold Boots (250534, -0.14 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (1.23 DPS) | yes | Band of the Fist (17694, -0.42 DPS) [quest]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.6 attack_power points (1.21 DPS) | yes | Band of the Fist (17694, -0.32 DPS, sim-verified) [quest]; Tiger Band (6749, -0.59 DPS) [quest]; Bonebinder's Signet (273032, -0.59 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-verified (60.6 DPS) | yes | Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop]; Thornspike (6681, -3.02 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (19.73 DPS) | yes | Zealot Blade (13033, -2.31 DPS, sim-verified) [world_drop] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.8 attack_power points (0.61 DPS) | yes | Gravespike Repeater (273033, -0.11 DPS) [world]; Golemsight Long Gun (273029, -0.14 DPS) [dungeon]; Long Battle Bow (15284, -0.23 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 553, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (troll, 00000000000000000-051531025120100221-000000000000000000)

Set DPS (verified): 141.9. Weights run: 6.7s. Verify run: 12.7s. 759 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.583 ± 0.060, crit=1.532 ± 0.077 per rating point (14 rating = 1%, 21.442 per %), hit=1.830 ± 0.128 per rating point (10 rating = 1%, 18.301 per %), melee_haste=9.575 ± 1.298

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.4 attack_power points (5.23 DPS) | yes | Chromite Barbute (8142, -0.49 DPS) [dungeon]; White Bandit Mask (10008, -0.88 DPS) [crafted]; Barbaric Iron Helm (7915, -1.67 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 17.4 attack_power points (1.92 DPS) | yes | Ethereal Talisman (4430, -0.12 DPS) [quest]; Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.60 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.4 attack_power points (3.02 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.60 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.65 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Imperial Cloak (6432, -0.49 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.49 DPS) [world_drop]; Parachute Cloak (10518, -0.75 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Kolkar Marauder Chain (6773, -1.32 DPS, sim-verified) [quest]; Carapace of Tuten'kash (10775, -2.60 DPS) [dungeon]; Jouster's Chestplate (8157, -2.90 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.3 attack_power points (2.24 DPS) | yes | Branded Leather Bracers (19508, -0.04 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.31 DPS) [world_drop]; Yorgen Bracers (13012, -0.39 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.18 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.2 attack_power points (3.55 DPS) | yes | Boar Champion's Belt (10768, -0.25 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.25 DPS) [rep]; Defiler's Leather Girdle (20192, -0.72 DPS, sim-verified) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (4.63 DPS) | yes | Firemane Leggings (13129, -0.44 DPS) [world_drop]; Triprunner Dungarees (9624, -0.83 DPS) [quest]; Symbolic Legplates (14829, -0.92 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.1 attack_power points (4.09 DPS) | yes | Blackforge Greaves (6423, -0.36 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.44 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.7 attack_power points (3.16 DPS) | yes | Falcon's Hook (7552, -0.93 DPS) [dungeon]; Mark of Kern (2262, -0.95 DPS) [dungeon]; Assault Band (13095, -0.95 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.7 attack_power points (2.29 DPS) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop]; Falcon's Hook (7552, -0.70 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Ardent Custodian (868, -1.81 DPS, sim-verified) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (48.70 DPS) | yes | Southsea Lamp (9359, -0.82 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (141.9 DPS) | yes | Master Hunter's Rifle (17687, -0.27 DPS) [quest]; The Silencer (13138, -0.30 DPS) [world_drop]; Swiftwind (13038, -0.62 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 759, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 30321000000000000-051531025120100231-000000000000000000)

Set DPS (verified): 220.6. Weights run: 7.1s. Verify run: 31.8s. 963 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.242 ± 0.079, crit=1.495 ± 0.107 per rating point (14 rating = 1%, 20.926 per %), hit=2.540 ± 0.178 per rating point (10 rating = 1%, 25.404 per %), melee_haste=12.929 ± 1.790

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.9 attack_power points (5.95 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.76 DPS) [crafted]; Sunscale Helmet (14849, -0.95 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.2 attack_power points (3.32 DPS) | yes | Woven Ivy Necklace (19159, -0.38 DPS) [quest]; Ethereal Talisman (4430, -1.42 DPS) [quest]; Scout's Medallion (19535, -1.43 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.9 attack_power points (4.56 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.25 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.35 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (220.6 DPS) | yes | Bloodlust Cape (14801, -0.08 DPS) [world_drop]; Pridelord Cape (14673, -0.18 DPS) [world_drop]; Blackflame Cape (13109, -0.47 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (220.6 DPS) | yes | Stone Guard's Plate Armor (220801, -0.54 DPS) [vendor]; Warforged Chestplate (11195, -0.92 DPS) [quest]; Warbear Harness (15064, -1.38 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.5 attack_power points (3.73 DPS) | yes | Bracers of the Stone Princess (17714, -0.18 DPS) [dungeon]; Arena Bands (18711, -0.18 DPS) [world]; Prowler's Leather Bracers (252539, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (220.6 DPS) | yes | Officer's Gloves (250551, +0.00 DPS) [crafted]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -4.25 DPS, sim-verified) [world_drop] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 42.9 attack_power points (5.44 DPS) | yes | Prowler's Leather Waistguard (252473, -0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.19 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.25 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 57.3 attack_power points (7.27 DPS) | yes | Golem Shard Leggings (13074, -1.69 DPS) [world_drop]; Scarlet Leggings (10330, -1.95 DPS) [dungeon]; Serpentskin Leggings (8262, -2.06 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.7 attack_power points (5.03 DPS) | yes | Skulker's Leather Boots (252469, -0.19 DPS) [crafted]; Officer's Sabatons (250561, -0.32 DPS) [crafted]; Officer's Boots (250546, -0.63 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.4 attack_power points (5.76 DPS) | yes | White Bone Band (11862, -2.71 DPS) [quest]; Mark of Kern (2262, -3.22 DPS) [dungeon]; Assault Band (13095, -3.22 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.2 attack_power points (3.95 DPS) | yes | White Bone Band (11862, -0.91 DPS) [quest]; Mark of Kern (2262, -1.42 DPS) [dungeon]; Assault Band (13095, -1.42 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (220.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.08 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (220.6 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Flurry Axe (871, -4.07 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 530.9 attack_power points (67.32 DPS) | yes | Thorium Cestus (250614, -2.56 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -6.18 DPS) [dungeon]; White Bone Shredder (11863, -10.05 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (220.6 DPS) | yes | Houndmaster's Bow (11628, -0.25 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.33 DPS) [world_drop]; Guttbuster (13139, -0.51 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Officer's Wristguards; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 963, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 30323013002000000-051531025120100231-200000000000000000)

Set DPS (verified): 327.5. Weights run: 7.4s. Verify run: 104.5s. 1824 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.881 ± 0.124, crit=2.687 ± 0.177 per rating point (14 rating = 1%, 37.617 per %), hit=3.291 ± 0.297 per rating point (10 rating = 1%, 32.913 per %), melee_haste=14.988 ± 2.778

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 177.1 attack_power points (19.01 DPS) | yes | Mask of the Unforgiven (13404, -6.82 DPS, sim-verified) [dungeon]; Champion's Plate Helm (227043, -6.93 DPS) [pvp]; Blood Guard's Plate Helm (220803, -8.65 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.5 attack_power points (3.70 DPS) | yes | Woven Ivy Necklace (19159, -0.59 DPS) [quest]; Scout's Medallion (19534, -0.67 DPS) [rep]; Sunprism Pendant (20649, -1.28 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 88.4 attack_power points (9.49 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -1.80 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.52 DPS) [rep]; Fel Cape (279269, -1.14 DPS) [crafted]; Cloak of Flames (3475, -1.64 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Gladiator Chain (11726, -0.42 DPS) [dungeon]; Dawn Armor (252483, -4.02 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -4.46 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 57.7 attack_power points (6.20 DPS) | yes | Berserker Bracers (19578, -0.50 DPS) [rep]; Vambraces of the Sadist (13400, -0.66 DPS) [dungeon]; General's Plate Armguards (16546, -1.13 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor]; Raider Gauntlets (272095, -1.02 DPS, sim-verified) [vendor] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190), Defiler's Plate Girdle (20204)) | The Defilers [rep] | 71.6 attack_power points (7.69 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; Defiler's Plate Girdle (20204, +0.00 DPS) [rep] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Chain Leggings (22673, -0.96 DPS, sim-verified) [rep]; Legionnaire's Plate Leggings (227048, -2.93 DPS) [pvp] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 53.0 attack_power points (5.70 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Defiler's Lamellar Greaves (20181, -0.27 DPS) [rep]; Defiler's Plate Greaves (20208, -0.27 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 84.5 attack_power points (9.08 DPS) | yes | Legionnaire's Band (19510, -4.28 DPS) [rep]; Band of the Penitent (13217, -5.04 DPS) [quest]; Don Rodrigo's Band (21563, -5.04 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 52.9 attack_power points (5.68 DPS) | yes | Legionnaire's Band (19510, -0.88 DPS) [rep]; Band of the Penitent (13217, -1.64 DPS) [quest]; Don Rodrigo's Band (21563, -1.64 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.80 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Annihilator (12798) | Blacksmithing [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Greenhammer (279261, -0.71 DPS, sim-verified) [crafted] |
| off_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | 711.0 attack_power points (76.34 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -24.56 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.01 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.31 DPS) [world_drop]; Guttbuster (13139, -1.88 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; waist: Defiler's Chain Girdle; legs: Outrider's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: Annihilator; off_hand: Ironwood Blade; ranged: Blackcrow

No-known-source sample (15 of 1824, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 30323013002000000-051531025120100231-200000000000000000)

Set DPS (verified): 870.5. Weights run: 8.2s. Verify run: 115.5s. 1824 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.108 ± 0.084, crit=1.417 ± 0.117 per rating point (14 rating = 1%, 19.836 per %), hit=2.594 ± 0.226 per rating point (10 rating = 1%, 25.943 per %), melee_haste=20.340 ± 2.446

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 131.2 attack_power points (37.69 DPS) | yes | Champion's Plate Helm (227043, -11.26 DPS) [pvp]; Warlord's Plate Headpiece (231535, -14.29 DPS) [pvp]; Stalwart Helm (250599, -23.19 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.4 attack_power points (7.30 DPS) | yes | Woven Ivy Necklace (19159, -0.64 DPS) [quest]; Frostwolf Legionnaire's Pendant (19095, -2.13 DPS) [rep]; Scout's Medallion (19534, -2.52 DPS) [rep] |
| shoulder | Defiler's Plate Spaulders (20212) (or Defiler's Lamellar Spaulders (20184)) | The Defilers [rep] | 58.4 attack_power points (16.79 DPS) | yes | Defiler's Lamellar Spaulders (20184, +0.00 DPS) [rep]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Champion's Plate Shoulders (227042, -0.35 DPS) [pvp] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.84 DPS) [dungeon]; Cloak of Flames (3475, -3.91 DPS) [world_drop]; Frostwolf Legionnaire's Cloak (19083, -4.47 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -7.35 DPS) [crafted]; Savage Gladiator Chain (11726, -8.97 DPS, sim-verified) [dungeon]; Warlord's Plate Armor (231530, -12.02 DPS) [pvp] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 50.7 attack_power points (14.56 DPS) | yes | Forest Stalker's Bracers (19587, -1.55 DPS) [rep]; General's Plate Armguards (16546, -1.58 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Gauntlets (231532, -2.23 DPS) [pvp]; Timbermaw Brawlers (19049, -2.84 DPS) [crafted]; Gloves of Holy Might (867, -4.54 DPS, sim-verified) [world_drop] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 61.7 attack_power points (17.73 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Defiler's Plate Girdle (20204, -1.29 DPS) [rep]; Brigam Girdle (13142, -1.43 DPS) [dungeon] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Chain Leggings (22673, -6.55 DPS) [rep]; Warbear Woolies (15065, -7.42 DPS) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 46.9 attack_power points (13.47 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Defiler's Lamellar Greaves (20181, -0.79 DPS) [rep]; Defiler's Plate Greaves (20208, -0.79 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 59.8 attack_power points (17.18 DPS) | yes | Legionnaire's Band (19510, -6.09 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.9 attack_power points (13.20 DPS) | yes | Legionnaire's Band (19510, -2.11 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.98 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -2.88 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -58.49 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.12 DPS) [quest]; Skull Splitting Crossbow (13039, -2.49 DPS) [world_drop]; Houndmaster's Bow (11628, -3.07 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Might of the Timbermaw; legs: Outrider's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Ironwood Blade; off_hand: Persuader; ranged: Blackcrow

No-known-source sample (15 of 1824, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

