# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 37.5. Weights run: 2.7s. Verify run: 4.0s. 287 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.338 ± 0.029, crit=0.741 ± 0.029 per rating point (14 rating = 1%, 10.371 per %), hit=0.934 ± 0.056 per rating point (10 rating = 1%, 9.338 per %), melee_haste=6.064 ± 0.595

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.21 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.40 DPS) [crafted]; Brawler's Leather Hood (252504, -0.56 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.48 DPS) | yes | Erudite's Amulet (277204, -0.35 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.7 attack_power points (0.40 DPS) | yes | Rough Bronze Shoulders (3480, -0.04 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.04 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.7 attack_power points (0.52 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.69 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.60 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.5 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -3.05 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (37.5 DPS) | yes | Brawler's Leather Belt (252428, -0.28 DPS) [crafted]; Deviate Scale Belt (6468, -0.32 DPS) [crafted]; Cobrahn's Grasp (6460, -3.21 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.85 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.7 attack_power points (1.01 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Defender's Leather Boots (252441, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.4 attack_power points (0.80 DPS) | yes | The 1 Ring (8350, -0.60 DPS) [world]; Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Ring of the Moon (12052, -0.68 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.48 DPS) | yes | The 1 Ring (8350, -0.25 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (15.00 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.2 attack_power points (14.42 DPS) | yes | Diamond Hammer (2194, -0.80 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.4 attack_power points (0.32 DPS) | yes | Dwarven Fishing Pole (3567, -0.08 DPS) [quest]; Fine Longbow (11304, -0.08 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.08 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 287, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 65.7. Weights run: 2.9s. Verify run: 2.3s. 516 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.449 ± 0.037, crit=1.050 ± 0.041 per rating point (14 rating = 1%, 14.704 per %), hit=1.235 ± 0.085 per rating point (10 rating = 1%, 12.355 per %), melee_haste=8.455 ± 0.935

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 31.0 attack_power points (1.93 DPS) | yes | Tusken Helm (6686, -0.26 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.44 DPS) [crafted]; Defender's Leather Helm (252455, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.8 attack_power points (0.86 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.11 DPS) [world]; Sentinel's Medallion (19541, -0.14 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 25.9 attack_power points (1.61 DPS) | yes | Barbaric Iron Shoulders (7913, -0.28 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.54 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.1 attack_power points (1.00 DPS) | yes | Sergeant Major's Cape (16315, -0.15 DPS) [pvp]; Tigerstrike Mantle (13108, -0.28 DPS) [world_drop]; Wolfmaster Cape (6314, -0.38 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.86 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.08 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.35 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 16.7 attack_power points (1.04 DPS) | yes | Yorgen Bracers (13012, -0.02 DPS) [world_drop]; Pugilist Bracers (4438, -0.04 DPS) [dungeon]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.7 attack_power points (1.41 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.04 DPS) [world]; Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 28.7 attack_power points (1.78 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.23 DPS) [crafted]; Girdle of Golem Strength (9405, -0.29 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.1 attack_power points (2.00 DPS) | yes | Brawler's Leather Legguards (252516, -0.24 DPS, sim-verified) [crafted]; Glimmering Mail Legguards (6386, -0.37 DPS) [world_drop]; Ferine Leggings (6690, -0.38 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.6 attack_power points (1.09 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Alacritous Treads (277234, -0.02 DPS) [quest]; Veteran's Boots (250503, -0.11 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 20.7 attack_power points (1.29 DPS) | yes | Seal of Wrynn (2933, -0.64 DPS) [quest]; Tiger Band (6749, -0.66 DPS) [quest]; Bonebinder's Signet (273032, -0.66 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.3 attack_power points (1.26 DPS) | yes | Seal of Wrynn (2933, -0.44 DPS, sim-verified) [quest]; Tiger Band (6749, -0.64 DPS) [quest]; Bonebinder's Signet (273032, -0.64 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (21.47 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (65.7 DPS) | yes | Shoni's Disarming Tool (9608, -5.82 DPS) [quest]; Swinetusk Shank (6691, -24.24 DPS, sim-verified) [dungeon] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.3 attack_power points (0.64 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.08 DPS) [vendor]; Golemsight Long Gun (273029, -0.10 DPS) [dungeon]; Gravespike Repeater (273033, -0.15 DPS) [world] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler

### Band 40 (human, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 124.3. Weights run: 3.1s. Verify run: 2.7s. 716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.640 ± 0.065, crit=1.533 ± 0.082 per rating point (14 rating = 1%, 21.466 per %), hit=1.157 ± 0.137 per rating point (10 rating = 1%, 11.573 per %), melee_haste=9.125 ± 1.138

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.7 attack_power points (3.76 DPS) | yes | White Bandit Mask (10008, -0.31 DPS) [crafted]; Barbaric Iron Helm (7915, -0.94 DPS) [crafted]; Hard Gold Coif (250537, -1.35 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 18.0 attack_power points (1.55 DPS) | yes | Kaleidoscope Chain (13084, -0.30 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.52 DPS) [world]; River Pride Choker (13087, -0.86 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.0 attack_power points (2.58 DPS) | yes | Forest Tracker Epaulets (2278, -0.17 DPS) [world_drop]; Flintrock Shoulders (7755, -0.31 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.69 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 21.8 attack_power points (1.88 DPS) | yes | Hawkeye's Cloak (14593, -0.37 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.52 DPS) [quest]; Parachute Cloak (10518, -0.75 DPS) [crafted] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 39.5 attack_power points (3.39 DPS) | yes | Icemail Jerkin (1981, -0.01 DPS) [world_drop]; Carapace of Tuten'kash (10775, -0.55 DPS) [dungeon]; Jouster's Chestplate (8157, -0.82 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (1.77 DPS) | yes | Branded Leather Bracers (19508, -0.05 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.5 attack_power points (3.57 DPS) | yes | Scarlet Gauntlets (10331, -0.09 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.44 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.58 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.8 attack_power points (2.82 DPS) | yes | Highlander's Leather Girdle (20116, -0.24 DPS) [rep]; Highlander's Plate Girdle (20125, -0.24 DPS) [rep]; Officer's Belt (250556, -0.25 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.61 DPS) | yes | Symbolic Legplates (14829, -0.19 DPS) [world_drop]; Firemane Leggings (13129, -0.34 DPS) [world_drop]; Triprunner Dungarees (9624, -0.56 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.5 attack_power points (3.22 DPS) | yes | Blackforge Greaves (6423, -0.26 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.47 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 29.1 attack_power points (2.50 DPS) | yes | Falcon's Hook (7552, -0.72 DPS) [dungeon]; Ring of the Underwood (2951, -0.75 DPS) [world_drop]; Mark of Kern (2262, -0.78 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (1.80 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.08 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (124.3 DPS) | yes | Shoni's Disarming Tool (9608, -19.41 DPS) [quest]; Ardent Custodian (868, -40.83 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.20 DPS) [quest]; The Silencer (13138, -0.25 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.23 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 177.9. Weights run: 3.2s. Verify run: 3.5s. 901 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.582 ± 0.064, crit=1.778 ± 0.088 per rating point (14 rating = 1%, 24.895 per %), hit=2.045 ± 0.154 per rating point (10 rating = 1%, 20.446 per %), melee_haste=10.857 ± 1.318

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 71.3 attack_power points (7.03 DPS) | yes | Chromite Barbute (8142, -2.79 DPS) [dungeon]; Sunscale Helmet (14849, -2.94 DPS) [world_drop]; Ornate Mithril Helm (7937, -3.35 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.6 attack_power points (3.01 DPS) | yes | Sentinel's Medallion (19539, -1.13 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -1.60 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 46.9 attack_power points (4.62 DPS) | yes | Wyrmslayer Spaulders (13066, -1.01 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -1.05 DPS) [crafted]; Officer's Pauldrons (250576, -2.12 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.7 attack_power points (2.34 DPS) | yes | Pridelord Cape (14673, -0.34 DPS) [world_drop]; Blackflame Cape (13109, -0.47 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.18 DPS, sim-verified) [pvp] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 54.9 attack_power points (5.41 DPS) | yes | Warbear Harness (15064, -0.44 DPS) [crafted]; Warforged Chestplate (11195, -0.68 DPS) [quest]; Blazewind Breastplate (11193, -1.23 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.7 attack_power points (3.13 DPS) | yes | Officer's Wristguards (250581, -0.02 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.26 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 59.8 attack_power points (5.89 DPS) | yes | Raider Gloves (272100, -0.49 DPS) [vendor]; Officer's Gloves (250551, -1.34 DPS) [crafted]; Gloves of Holy Might (867, -1.47 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 47.0 attack_power points (4.63 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.01 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.21 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 64.7 attack_power points (6.38 DPS) | yes | Centurion Legplates (10740, -1.08 DPS) [quest]; Serpentskin Leggings (8262, -1.32 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -2.52 DPS, sim-verified) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 43.4 attack_power points (4.28 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Officer's Sabatons (250561, -0.31 DPS) [crafted]; Officer's Boots (250546, -0.62 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 34.2 attack_power points (3.37 DPS) | yes | Thunderbrow Ring (13097, -1.33 DPS) [world_drop]; Falcon's Hook (7552, -1.38 DPS) [dungeon]; Blackstone Ring (17713, -1.40 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.1 attack_power points (2.18 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Blackstone Ring (17713, -0.21 DPS) [dungeon]; Thunderbrow Ring (13097, -1.42 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (177.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (177.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (177.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Flurry Axe (871, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 534.9 attack_power points (52.71 DPS) | yes | Thorium Cestus (250614, -3.46 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.20 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.57 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (177.9 DPS) | yes | Skull Splitting Crossbow (13039, -0.21 DPS) [world_drop]; Guttbuster (13139, -0.34 DPS) [world_drop]; Dark Iron Rifle (16004, -3.50 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Dark Phantom Cape; chest: Knight's Plate Hauberk; wrist: Deepfury Bracers; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 901, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 298.7. Weights run: 3.2s. Verify run: 10.4s. 1630 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.844 ± 0.089, crit=2.634 ± 0.127 per rating point (14 rating = 1%, 36.881 per %), hit=3.040 ± 0.233 per rating point (10 rating = 1%, 30.398 per %), melee_haste=16.844 ± 2.072

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 170.6 attack_power points (17.04 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -6.12 DPS) [vendor]; Outlaw's Collar (279253, -11.43 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (298.7 DPS) | yes | Amulet of the Darkmoon (19491, -0.58 DPS) [quest]; Beads of Ogre Might (22150, -0.65 DPS) [quest]; Rage of Mugamba (19577, -3.35 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 86.7 attack_power points (8.66 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -1.58 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 58.4 attack_power points (5.84 DPS) | yes | Cloak of the Honor Guard (20073, -1.52 DPS) [rep]; Fel Cape (279269, -2.15 DPS) [crafted]; Cape of the Black Baron (13340, -2.35 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 136.1 attack_power points (13.60 DPS) | yes | Dawn Armor (252483, -1.89 DPS, sim-verified) [crafted]; Obsidian Mail Tunic (22191, -2.32 DPS) [crafted]; Breastplate of Heroism (226862, -3.97 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 57.0 attack_power points (5.70 DPS) | yes | Berserker Bracers (19578, -0.43 DPS) [rep]; Bracers of Subterfuge (22668, -0.94 DPS) [quest]; Marshal's Plate Bracers (16481, -1.01 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 74.3 attack_power points (7.42 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gauntlets of Heroism (226861, -0.14 DPS) [quest]; Raider Gloves (272099, -0.25 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 90.9 attack_power points (9.08 DPS) | yes | Ferocity of the Timbermaw (227805, -0.75 DPS) [vendor]; Belt of Preserved Heads (20216, -1.50 DPS, sim-verified) [quest]; Valiant's Waistguard (272403, -1.86 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 168.7 attack_power points (16.86 DPS) | yes | Sentinel's Plate Legguards (237825, -1.05 DPS) [vendor]; Titanic Leggings (22385, -1.10 DPS) [crafted]; Marshal's Plate Legguards (231540, -2.45 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 70.4 attack_power points (7.03 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Drudge Boots (21532, -1.30 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 81.3 attack_power points (8.12 DPS) | yes | Cutthroat's Signet (272408, -2.50 DPS) [vendor]; Protector's Band (19514, -3.70 DPS) [rep]; Don Rodrigo's Band (21563, -4.44 DPS) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 78.4 attack_power points (7.83 DPS) | yes | Cutthroat's Signet (272408, -2.22 DPS) [vendor]; Protector's Band (19514, -3.41 DPS) [rep]; Don Rodrigo's Band (21563, -4.15 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (298.7 DPS) | yes | Frozen Heart of the Mountain (249469, -7.63 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (298.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (298.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -10.25 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 844.9 attack_power points (84.42 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -2.81 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -36.24 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (298.7 DPS) | yes | Bloodseeker (19107, -0.80 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.11 DPS) [world_drop]; Dark Iron Rifle (16004, -4.52 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1630, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 60, raid preset (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 877.5. Weights run: 3.6s. Verify run: 11.3s. 1630 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.003, agility=1.993 ± 0.117, crit=2.592 ± 0.163 per rating point (14 rating = 1%, 36.293 per %), hit=4.113 ± 0.299 per rating point (10 rating = 1%, 41.134 per %), melee_haste=22.331 ± 3.040

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 194.5 attack_power points (41.09 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -14.96 DPS) [vendor]; Stalwart Helm (250599, -18.20 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (877.5 DPS) | yes | Medallion of the Dawn (22659, -1.02 DPS) [quest]; Amulet of the Darkmoon (19491, -1.11 DPS) [quest]; Rage of Mugamba (19577, -7.32 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 90.1 attack_power points (19.04 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -3.47 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 69.1 attack_power points (14.61 DPS) | yes | Cape of the Black Baron (13340, -4.07 DPS) [dungeon]; Cloak of the Honor Guard (20073, -5.32 DPS) [rep]; Stalwart Cloak (272415, -5.92 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.2 attack_power points (31.94 DPS) | yes | Timbermaw Tunic (252484, -1.35 DPS) [crafted]; Breastplate of Heroism (226862, -8.01 DPS) [quest]; Obsidian Mail Tunic (22191, -8.21 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 62.1 attack_power points (13.11 DPS) | yes | Berserker Bracers (19578, -0.91 DPS) [rep]; Bracers of Subterfuge (22668, -2.15 DPS) [quest]; Marshal's Plate Bracers (16481, -2.26 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 81.3 attack_power points (17.18 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.22 DPS) [pvp]; Raider Gloves (272099, -0.70 DPS) [vendor]; Stormshroud Gloves (21278, -0.82 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 101.8 attack_power points (21.51 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.30 DPS) [vendor]; Ferocity of the Timbermaw (227805, -2.27 DPS) [vendor]; Marshal's Plate Girdle (16482, -4.84 DPS) [pvp] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 184.6 attack_power points (39.00 DPS) | yes | Sentinel's Chain Leggings (237819, -0.23 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.42 DPS) [vendor]; Marshal's Plate Legguards (231540, -5.67 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 85.1 attack_power points (17.99 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -4.67 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 93.9 attack_power points (19.85 DPS) | yes | Cutthroat's Signet (272408, -5.26 DPS) [vendor]; Protector's Band (19514, -9.64 DPS) [rep]; Channeler's Ring (272406, -11.16 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 91.4 attack_power points (19.32 DPS) | yes | Cutthroat's Signet (272408, -5.97 DPS, sim-verified) [vendor]; Protector's Band (19514, -9.11 DPS) [rep]; Channeler's Ring (272406, -10.63 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (877.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (877.5 DPS) | yes | Darkmoon Card: Heroism (19287, -5.65 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (877.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -17.12 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 848.6 attack_power points (179.30 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -8.56 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -77.43 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (877.5 DPS) | yes | Bloodseeker (19107, -1.00 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.77 DPS) [world_drop]; Dark Iron Rifle (16004, -10.56 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Belt of Preserved Heads; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1630, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

## Horde

### Band 20 (troll, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 37.7. Weights run: 2.7s. Verify run: 3.9s. 257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.338 ± 0.029, crit=0.741 ± 0.029 per rating point (14 rating = 1%, 10.371 per %), hit=0.934 ± 0.056 per rating point (10 rating = 1%, 9.338 per %), melee_haste=6.064 ± 0.595

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.21 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.40 DPS) [crafted]; Brawler's Leather Hood (252504, -0.56 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.0 attack_power points (0.48 DPS) | yes | Erudite's Amulet (277204, -0.24 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.7 attack_power points (0.40 DPS) | yes | Rough Bronze Shoulders (3480, -0.04 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.04 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.7 attack_power points (0.52 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.7 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.14 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.60 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.7 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.61 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (37.7 DPS) | yes | Brawler's Leather Belt (252428, -0.28 DPS) [crafted]; Deviate Scale Belt (6468, -0.32 DPS) [crafted]; Cobrahn's Grasp (6460, -2.76 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.7 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.40 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.7 attack_power points (1.01 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Defender's Leather Boots (252441, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.4 attack_power points (0.80 DPS) | yes | Loop of Sacrifice (281673, -0.44 DPS) [quest]; Bounty Hunter's Ring (5351, -0.56 DPS) [quest]; The 1 Ring (8350, -0.60 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Bounty Hunter's Ring (5351, -0.24 DPS) [quest]; The 1 Ring (8350, -0.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (15.00 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.2 attack_power points (14.42 DPS) | yes | Diamond Hammer (2194, -0.51 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.4 attack_power points (0.32 DPS) | yes | Fine Longbow (11304, -0.08 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.08 DPS) [crafted]; Deadly Blunderbuss (4369, -0.16 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 257, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 209616 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 66.8. Weights run: 2.9s. Verify run: 2.4s. 478 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.449 ± 0.037, crit=1.050 ± 0.041 per rating point (14 rating = 1%, 14.704 per %), hit=1.235 ± 0.085 per rating point (10 rating = 1%, 12.355 per %), melee_haste=8.455 ± 0.935

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 31.0 attack_power points (1.93 DPS) | yes | Tusken Helm (6686, -0.31 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.44 DPS) [crafted]; Defender's Leather Helm (252455, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.8 attack_power points (0.86 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.11 DPS) [world]; Scout's Medallion (19537, -0.14 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 25.9 attack_power points (1.61 DPS) | yes | Barbaric Iron Shoulders (7913, -0.33 DPS) [crafted]; Barbaric Shoulders (5964, -0.54 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.1 attack_power points (1.00 DPS) | yes | Tigerstrike Mantle (13108, -0.25 DPS, sim-verified) [world_drop]; Wildhunter Cloak (16658, -0.38 DPS) [quest]; Construct Cloak (279848, -0.38 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.86 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.08 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.35 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 16.7 attack_power points (1.04 DPS) | yes | Yorgen Bracers (13012, -0.02 DPS) [world_drop]; Pugilist Bracers (4438, -0.04 DPS) [dungeon]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.7 attack_power points (1.41 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.04 DPS) [world]; Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 28.7 attack_power points (1.78 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.23 DPS) [crafted]; Girdle of Golem Strength (9405, -0.29 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.1 attack_power points (2.00 DPS) | yes | Brawler's Leather Legguards (252516, -0.28 DPS, sim-verified) [crafted]; Glimmering Mail Legguards (6386, -0.37 DPS) [world_drop]; Ferine Leggings (6690, -0.38 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.6 attack_power points (1.09 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.11 DPS) [crafted]; Hard Gold Boots (250534, -0.22 DPS) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 20.7 attack_power points (1.29 DPS) | yes | Band of the Fist (17694, -0.43 DPS) [quest]; Tiger Band (6749, -0.66 DPS) [quest]; Bonebinder's Signet (273032, -0.66 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.3 attack_power points (1.26 DPS) | yes | Band of the Fist (17694, -0.34 DPS, sim-verified) [quest]; Tiger Band (6749, -0.64 DPS) [quest]; Bonebinder's Signet (273032, -0.64 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (21.47 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (66.8 DPS) | yes | Swinetusk Shank (6691, -25.04 DPS, sim-verified) [dungeon] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.3 attack_power points (0.64 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.08 DPS) [vendor]; Golemsight Long Gun (273029, -0.10 DPS) [dungeon]; Gravespike Repeater (273033, -0.15 DPS) [world] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 478, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (troll, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 125.8. Weights run: 3.1s. Verify run: 2.8s. 667 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.640 ± 0.065, crit=1.533 ± 0.082 per rating point (14 rating = 1%, 21.466 per %), hit=1.157 ± 0.137 per rating point (10 rating = 1%, 11.573 per %), melee_haste=9.125 ± 1.138

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.7 attack_power points (3.76 DPS) | yes | White Bandit Mask (10008, -0.31 DPS) [crafted]; Barbaric Iron Helm (7915, -0.94 DPS) [crafted]; Hard Gold Coif (250537, -1.35 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 18.0 attack_power points (1.55 DPS) | yes | Ethereal Talisman (4430, -0.13 DPS) [quest]; Kaleidoscope Chain (13084, -0.30 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.52 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.0 attack_power points (2.58 DPS) | yes | Forest Tracker Epaulets (2278, -0.17 DPS) [world_drop]; Flintrock Shoulders (7755, -0.31 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.69 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (1.50 DPS) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 39.5 attack_power points (3.39 DPS) | yes | Icemail Jerkin (1981, -0.01 DPS) [world_drop]; Carapace of Tuten'kash (10775, -0.55 DPS) [dungeon]; Jouster's Chestplate (8157, -0.82 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (1.77 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.5 attack_power points (3.57 DPS) | yes | Scarlet Gauntlets (10331, -0.09 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.44 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.58 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.8 attack_power points (2.82 DPS) | yes | Defiler's Leather Girdle (20192, -0.24 DPS) [rep]; Defiler's Plate Girdle (20206, -0.24 DPS) [rep]; Officer's Belt (250556, -0.25 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.61 DPS) | yes | Symbolic Legplates (14829, -0.19 DPS) [world_drop]; Firemane Leggings (13129, -0.34 DPS) [world_drop]; Triprunner Dungarees (9624, -0.56 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.5 attack_power points (3.22 DPS) | yes | Blackforge Greaves (6423, -0.26 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.47 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 29.1 attack_power points (2.50 DPS) | yes | Falcon's Hook (7552, -0.72 DPS) [dungeon]; Ring of the Underwood (2951, -0.75 DPS) [world_drop]; Mark of Kern (2262, -0.78 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (1.80 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.08 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (125.8 DPS) | yes | Ardent Custodian (868, -41.59 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.20 DPS) [quest]; The Silencer (13138, -0.25 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.20 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Kolkar Marauder Chain; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 667, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 182.0. Weights run: 3.2s. Verify run: 3.6s. 843 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.582 ± 0.064, crit=1.778 ± 0.088 per rating point (14 rating = 1%, 24.895 per %), hit=2.045 ± 0.154 per rating point (10 rating = 1%, 20.446 per %), melee_haste=10.857 ± 1.318

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 44.9 attack_power points (4.42 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.19 DPS) [dungeon]; Sunscale Helmet (14849, -0.34 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.6 attack_power points (3.01 DPS) | yes | Woven Ivy Necklace (19159, -0.43 DPS) [quest]; Scout's Medallion (19535, -1.14 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 38.7 attack_power points (3.81 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.20 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.7 attack_power points (2.34 DPS) | yes | Pridelord Cape (14673, -0.34 DPS) [world_drop]; Blackflame Cape (13109, -0.47 DPS) [world_drop]; Bloodlust Cape (14801, -0.56 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 54.9 attack_power points (5.41 DPS) | yes | Warbear Harness (15064, -0.44 DPS) [crafted]; Warforged Chestplate (11195, -0.68 DPS) [quest]; Blazewind Breastplate (11193, -1.23 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.7 attack_power points (3.13 DPS) | yes | Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Officer's Wristguards (250581, +0.00 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 59.8 attack_power points (5.89 DPS) | yes | Raider Gloves (272100, -0.49 DPS) [vendor]; Officer's Gloves (250551, -1.34 DPS) [crafted]; Gloves of Holy Might (867, -1.47 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 47.0 attack_power points (4.63 DPS) | yes | Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.21 DPS) [rep]; Defiler's Leather Girdle (20193, -0.21 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 64.7 attack_power points (6.38 DPS) | yes | Serpentskin Leggings (8262, -1.32 DPS) [world_drop]; Stormshroud Pants (15057, -1.47 DPS) [crafted]; Golem Shard Leggings (13074, -2.04 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 43.4 attack_power points (4.28 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Officer's Sabatons (250561, -0.31 DPS) [crafted]; Officer's Boots (250546, -0.62 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 34.2 attack_power points (3.37 DPS) | yes | Masons Fraternity Ring (9533, -1.19 DPS) [quest]; Thunderbrow Ring (13097, -1.33 DPS) [world_drop]; Falcon's Hook (7552, -1.38 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.36 DPS) | yes | Masons Fraternity Ring (9533, -0.18 DPS) [quest]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Falcon's Hook (7552, -0.37 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (182.0 DPS) | yes | Molten Heart of the Mountain (249470, -5.50 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (182.0 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (182.0 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Flurry Axe (871, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 534.9 attack_power points (52.71 DPS) | yes | Thorium Cestus (250614, -5.03 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.20 DPS) [dungeon]; White Bone Shredder (11863, -7.97 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (182.0 DPS) | yes | Skull Splitting Crossbow (13039, -0.21 DPS) [world_drop]; Guttbuster (13139, -0.34 DPS) [world_drop]; Dark Iron Rifle (16004, -2.00 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Deepfury Bracers; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 843, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 308.4. Weights run: 3.2s. Verify run: 10.5s. 1599 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.844 ± 0.089, crit=2.634 ± 0.127 per rating point (14 rating = 1%, 36.881 per %), hit=3.040 ± 0.233 per rating point (10 rating = 1%, 30.398 per %), melee_haste=16.844 ± 2.072

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 170.6 attack_power points (17.04 DPS) | yes | Champion's Plate Helm (227043, -6.12 DPS) [pvp]; Blood Guard's Plate Helm (220803, -7.72 DPS) [vendor]; Outlaw's Collar (279253, -12.04 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (308.4 DPS) | yes | Amulet of the Darkmoon (19491, -0.58 DPS) [quest]; Beads of Ogre Might (22150, -0.65 DPS) [quest]; Rage of Mugamba (19577, -2.83 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 86.7 attack_power points (8.66 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -1.58 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 58.4 attack_power points (5.84 DPS) | yes | Cape of the Black Baron (13340, -1.07 DPS) [dungeon]; Deathguard's Cloak (20068, -1.52 DPS) [rep]; Fel Cape (279269, -2.15 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 136.1 attack_power points (13.60 DPS) | yes | Dawn Armor (252483, -0.06 DPS) [crafted]; Obsidian Mail Tunic (22191, -2.32 DPS) [crafted]; Breastplate of Heroism (226862, -3.97 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 57.0 attack_power points (5.70 DPS) | yes | Berserker Bracers (19578, -0.43 DPS) [rep]; Bracers of Subterfuge (22668, -0.94 DPS) [quest]; General's Plate Armguards (16546, -1.01 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 74.3 attack_power points (7.42 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gauntlets of Heroism (226861, -0.14 DPS) [quest]; Raider Gloves (272099, -0.25 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 90.9 attack_power points (9.08 DPS) | yes | Belt of Preserved Heads (20216, -0.48 DPS) [quest]; Ferocity of the Timbermaw (227805, -0.75 DPS) [vendor]; Valiant's Waistguard (272403, -1.86 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 168.7 attack_power points (16.86 DPS) | yes | Sentinel's Plate Legguards (237825, -1.05 DPS) [vendor]; Titanic Leggings (22385, -1.10 DPS) [crafted]; General's Plate Leggings (231533, -2.45 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 70.4 attack_power points (7.03 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Drudge Boots (21532, -1.30 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 81.3 attack_power points (8.12 DPS) | yes | Cutthroat's Signet (272408, -2.50 DPS) [vendor]; Legionnaire's Band (19510, -3.70 DPS) [rep]; Don Rodrigo's Band (21563, -4.44 DPS) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 78.4 attack_power points (7.83 DPS) | yes | Cutthroat's Signet (272408, -2.22 DPS) [vendor]; Legionnaire's Band (19510, -3.41 DPS) [rep]; Don Rodrigo's Band (21563, -4.15 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (308.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (308.4 DPS) | yes | Frozen Heart of the Mountain (249469, -5.98 DPS, sim-verified) [crafted] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (308.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -11.50 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 844.9 attack_power points (84.42 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -2.17 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -36.24 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (308.4 DPS) | yes | Bloodseeker (19107, -0.80 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.11 DPS) [world_drop]; Dark Iron Rifle (16004, -3.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1599, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 894.1. Weights run: 3.6s. Verify run: 11.3s. 1599 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.003, agility=1.993 ± 0.117, crit=2.592 ± 0.163 per rating point (14 rating = 1%, 36.293 per %), hit=4.113 ± 0.299 per rating point (10 rating = 1%, 41.134 per %), melee_haste=22.331 ± 3.040

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 194.5 attack_power points (41.09 DPS) | yes | Champion's Plate Helm (227043, -14.96 DPS) [pvp]; Blood Guard's Plate Helm (220803, -18.68 DPS) [vendor]; Stalwart Helm (250599, -23.71 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (894.1 DPS) | yes | Medallion of the Dawn (22659, -1.02 DPS) [quest]; Amulet of the Darkmoon (19491, -1.11 DPS) [quest]; Rage of Mugamba (19577, -5.28 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 90.1 attack_power points (19.04 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -3.47 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 69.1 attack_power points (14.61 DPS) | yes | Cape of the Black Baron (13340, -4.07 DPS) [dungeon]; Deathguard's Cloak (20068, -5.32 DPS) [rep]; Stalwart Cloak (272415, -5.92 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.2 attack_power points (31.94 DPS) | yes | Timbermaw Tunic (252484, -1.35 DPS) [crafted]; Breastplate of Heroism (226862, -8.01 DPS) [quest]; Obsidian Mail Tunic (22191, -8.21 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 62.1 attack_power points (13.11 DPS) | yes | Berserker Bracers (19578, -0.91 DPS) [rep]; Bracers of Subterfuge (22668, -2.15 DPS) [quest]; General's Plate Armguards (16546, -2.26 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 81.3 attack_power points (17.18 DPS) | yes | General's Plate Gauntlets (231532, -0.22 DPS) [pvp]; Raider Gloves (272099, -0.70 DPS) [vendor]; Stormshroud Gloves (21278, -0.82 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 101.8 attack_power points (21.51 DPS) | yes | Radiant Girdle of the Dawn (227814, +0.00 DPS, sim-verified) [vendor]; Ferocity of the Timbermaw (227805, -2.27 DPS) [vendor]; General's Plate Girdle (16547, -4.84 DPS) [pvp] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 184.6 attack_power points (39.00 DPS) | yes | Sentinel's Chain Leggings (237819, -0.23 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.42 DPS) [vendor]; General's Plate Leggings (231533, -5.67 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 85.1 attack_power points (17.99 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -4.67 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 93.9 attack_power points (19.85 DPS) | yes | Cutthroat's Signet (272408, -5.26 DPS) [vendor]; Legionnaire's Band (19510, -9.64 DPS) [rep]; Channeler's Ring (272406, -11.16 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 91.4 attack_power points (19.32 DPS) | yes | Cutthroat's Signet (272408, -6.68 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -9.11 DPS) [rep]; Channeler's Ring (272406, -10.63 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (894.1 DPS) | yes | Frozen Heart of the Mountain (249469, -7.14 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (894.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -9.27 DPS, sim-verified) [crafted] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (894.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -12.95 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 848.6 attack_power points (179.30 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -8.27 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -77.43 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (894.1 DPS) | yes | Bloodseeker (19107, -1.00 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.77 DPS) [world_drop]; Dark Iron Rifle (16004, -9.19 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Belt of Preserved Heads; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Darkmoon Card: Maelstrom; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1599, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

