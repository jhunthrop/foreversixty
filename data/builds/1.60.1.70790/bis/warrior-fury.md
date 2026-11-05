# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-050510000000000000-000000000000000000)

Set DPS (verified): 35.4. Weights run: 6.4s. Verify run: 24.0s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.006 ± 0.024, crit=0.567 ± 0.025 per rating point (14 rating = 1%, 7.939 per %), hit=0.869 ± 0.054 per rating point (10 rating = 1%, 8.686 per %), melee_haste=5.240 ± 0.341

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.15 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.57 DPS) [crafted]; Brawler's Leather Hood (252504, -0.69 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.0 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.30 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.69 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.34 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-verified (35.4 DPS) | yes | Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.11 DPS) [dungeon]; Grave Shroud (279865, -0.11 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.15 DPS) | yes | Brawler's Leather Armor (252490, -0.17 DPS) [crafted]; Defender's Leather Armor (252434, -0.17 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.57 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.06 DPS) [quest]; Bristlebark Bindings (14569, -0.17 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.4 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.65 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.34 DPS) [crafted]; Deviate Scale Belt (6468, -0.40 DPS) [crafted]; Cobrahn's Grasp (6460, -2.91 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (35.4 DPS) | yes | Chausses of Westfall (6087, +0.00 DPS) [quest]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.61 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (35.4 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -2.39 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.69 DPS) | yes | The 1 Ring (8350, -0.52 DPS) [world]; Lavishly Jeweled Ring (1156, -0.58 DPS) [dungeon]; Ring of the Moon (12052, -0.58 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.46 DPS) | yes | The 1 Ring (8350, -0.32 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Ring of the Moon (12052, -0.34 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.5 attack_power points (13.71 DPS) | yes | Rockslicer (872, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (13.38 DPS) | yes | Cruel Barb (5191, +0.00 DPS) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.23 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Fine Longbow (11304, -0.00 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 30 (night-elf, 00000000000000000-050530025000100000-000000000000000000)

Set DPS (verified): 62.5. Weights run: 7.5s. Verify run: 18.2s. 596 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.176 ± 0.038, crit=0.860 ± 0.042 per rating point (14 rating = 1%, 12.042 per %), hit=1.194 ± 0.107 per rating point (10 rating = 1%, 11.936 per %), melee_haste=7.591 ± 0.772

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 28.6 attack_power points (1.88 DPS) | yes | Veteran's Chain Helm (250498, -0.30 DPS) [crafted]; Defender's Leather Helm (252455, -0.30 DPS) [crafted]; Tusken Helm (6686, -0.47 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.7 attack_power points (0.83 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Sentinel's Medallion (19541, -0.22 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.9 attack_power points (1.51 DPS) | yes | Barbaric Iron Shoulders (7913, -0.36 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.46 DPS) [crafted]; Golden Scale Shoulders (3841, -0.59 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (62.5 DPS) | yes | Sergeant Major's Cape (16315, -0.10 DPS) [pvp]; Wolfmaster Cape (6314, -0.28 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.32 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-verified (62.5 DPS) | yes | Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (1.05 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.22 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.44 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.19 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.1 attack_power points (1.78 DPS) | yes | Prowler's Leather Belt (252459, -0.13 DPS) [crafted]; Girdle of Golem Strength (9405, -0.20 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.20 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.2 attack_power points (1.98 DPS) | yes | Brawler's Leather Legguards (252516, -0.37 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.39 DPS) [world_drop]; Ferine Leggings (6690, -0.52 DPS, sim-verified) [dungeon] |
| feet | Alacritous Treads (277234) (or Brawler's Leather Boots (252439)) | WANTED: Incinerator Gar'im [quest] | 15.9 attack_power points (1.04 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.5 attack_power points (1.28 DPS) | yes | Silverlaine's Family Seal (6321, -0.63 DPS) [dungeon]; Tiger Band (6749, -0.63 DPS) [quest]; Bonebinder's Signet (273032, -0.63 DPS) [world] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.1 attack_power points (1.25 DPS) | yes | Silverlaine's Family Seal (6321, -0.59 DPS) [dungeon]; Tiger Band (6749, -0.59 DPS) [quest]; Bonebinder's Signet (273032, -0.59 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-verified (62.5 DPS) | yes | Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop]; Thornspike (6681, -2.50 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (20.89 DPS) | yes | Zealot Blade (13033, -2.16 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -5.66 DPS) [quest] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.5 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.10 DPS) [world]; Golemsight Long Gun (273029, -0.16 DPS) [dungeon]; Long Battle Bow (15284, -0.23 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Alacritous Treads; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 596, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 00000000000000000-050530025122100221-000000000000000000)

Set DPS (verified): 144.0. Weights run: 7.4s. Verify run: 28.9s. 815 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.535 ± 0.051, crit=1.426 ± 0.067 per rating point (14 rating = 1%, 19.968 per %), hit=1.304 ± 0.110 per rating point (10 rating = 1%, 13.041 per %), melee_haste=6.739 ± 1.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.0 attack_power points (4.48 DPS) | yes | Chromite Barbute (8142, -0.35 DPS) [dungeon]; White Bandit Mask (10008, -0.69 DPS) [crafted]; Barbaric Iron Helm (7915, -1.38 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.95 DPS) | yes | Sentinel's Medallion (19540, -0.30 DPS) [rep]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.9 attack_power points (2.62 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.48 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.55 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (144.0 DPS) | yes | Hawkeye's Cloak (14593, -0.44 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.58 DPS) [quest]; Parachute Cloak (10518, -0.87 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, -0.46 DPS) [dungeon]; Kolkar Marauder Chain (6773, -1.09 DPS) [quest]; Carapace of Tuten'kash (10775, -1.72 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.96 DPS) | yes | Branded Leather Bracers (19508, -0.01 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.29 DPS) [world_drop]; Yorgen Bracers (13012, -0.34 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (144.0 DPS) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -2.66 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (3.10 DPS) | yes | Boar Champion's Belt (10768, -0.18 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.18 DPS) [rep]; Highlander's Plate Girdle (20125, -0.18 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (4.09 DPS) | yes | Symbolic Legplates (14829, -0.27 DPS) [world_drop]; Firemane Leggings (13129, -0.39 DPS) [world_drop]; Orcish War Leggings (7929, -0.78 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 36.7 attack_power points (3.58 DPS) | yes | Blackforge Greaves (6423, -0.35 DPS, sim-verified) [dungeon]; Prowler's Leather Shoes (252465, -0.39 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.57 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.3 attack_power points (2.76 DPS) | yes | Mark of Kern (2262, -0.81 DPS) [dungeon]; Assault Band (13095, -0.81 DPS) [world_drop]; Falcon's Hook (7552, -0.83 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.6 attack_power points (2.01 DPS) | yes | Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop]; Falcon's Hook (7552, -0.08 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (144.0 DPS) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (43.09 DPS) | yes | Southsea Lamp (9359, -0.87 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -20.46 DPS) [quest] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.24 DPS) [quest]; The Silencer (13138, -0.25 DPS) [world_drop]; Swiftwind (13038, -0.57 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Thunderbrow Ring; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 815, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 30321000000000000-050530025122100231-000000000000000000)

Set DPS (verified): 227.2. Weights run: 8.1s. Verify run: 32.6s. 1030 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.466 ± 0.086, crit=1.753 ± 0.117 per rating point (14 rating = 1%, 24.537 per %), hit=2.643 ± 0.192 per rating point (10 rating = 1%, 26.425 per %), melee_haste=14.513 ± 1.876

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 87.0 attack_power points (10.76 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.24 DPS) [vendor]; Raging Berserker's Helm (7719, -3.91 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -4.82 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.1 attack_power points (3.60 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.12 DPS) [quest]; Sentinel's Medallion (19539, -1.42 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 37.7 attack_power points (4.67 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.25 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.31 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.15 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight's Plate Hauberk (220794, -0.88 DPS) [vendor]; Mixologist's Tunic (12793, -1.17 DPS) [dungeon]; Warbear Harness (15064, -1.64 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (3.81 DPS) | yes | Deepfury Bracers (13120, -0.10 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.31 DPS) [crafted]; Runed Golem Shackles (12550, -0.35 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.99 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, -0.07 DPS) [dungeon]; Prowler's Leather Waistguard (252473, -0.12 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.25 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 63.2 attack_power points (7.82 DPS) | yes | Gryphon Rider's Leggings (9652, -1.39 DPS) [quest]; Centurion Legplates (10740, -1.39 DPS) [quest]; Serpentskin Leggings (8262, -1.70 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (227.2 DPS) | yes | Battlechaser's Greaves (12555, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.36 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.4 attack_power points (5.75 DPS) | yes | Masons Fraternity Ring (9533, -3.21 DPS) [quest]; Thunderbrow Ring (13097, -3.22 DPS) [world_drop]; Mark of Kern (2262, -3.27 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 33.2 attack_power points (4.11 DPS) | yes | Masons Fraternity Ring (9533, -1.57 DPS) [quest]; Thunderbrow Ring (13097, -1.58 DPS) [world_drop]; Mark of Kern (2262, -1.63 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -3.18 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -2.58 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Flurry Axe (871, -3.53 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadowblade (2163, -2.39 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -6.47 DPS) [dungeon]; Shoni's Disarming Tool (9608, -37.43 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.07 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.07 DPS) [world_drop]; Houndmaster's Bow (11628, -0.31 DPS) [dungeon] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Officer's Wristguards; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1030, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 364.9. Weights run: 7.8s. Verify run: 135.3s. 2366 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.474 ± 0.093, crit=2.106 ± 0.133 per rating point (14 rating = 1%, 29.482 per %), hit=1.312 ± 0.219 per rating point (10 rating = 1%, 13.122 per %), melee_haste=16.683 ± 2.329

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 121.2 attack_power points (19.11 DPS) | yes | Field Marshal's Plate Helm (231538, -5.63 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -5.77 DPS) [vendor]; Eye of Rend (12587, -11.15 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Jewel (11933, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Will of the Martyr (17044, +0.00 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 69.3 attack_power points (10.93 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, -0.91 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.12 DPS) [rep]; Howler's Furs (272414, -0.16 DPS) [vendor]; Blackveil Cape (11626, -1.49 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -3.53 DPS, sim-verified) [dungeon]; Dawn Armor (252483, -3.92 DPS) [crafted]; Cadaverous Armor (14637, -5.40 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 50.0 attack_power points (7.89 DPS) | yes | Berserker Bracers (19578, -0.03 DPS) [rep]; Marshal's Plate Bracers (16481, -0.90 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gauntlets (271940, -2.26 DPS, sim-verified) [rep] |
| waist | Highlander's Plate Girdle (20041) (or Highlander's Chain Girdle (20043), Highlander's Leather Girdle (20045)) | The League of Arathor [rep] | 63.5 attack_power points (10.01 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Chain Girdle (20043, +0.00 DPS) [rep]; Highlander's Leather Girdle (20045, +0.00 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -0.38 DPS) [vendor]; Marshal's Plate Legguards (231540, -2.21 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.31 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 49.8 attack_power points (7.85 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Battlechaser's Greaves (12555, -0.41 DPS) [dungeon]; Highlander's Plate Greaves (20048, -0.65 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.75 DPS) [dungeon]; Magni's Will (12548, -2.38 DPS) [quest]; Naglering (11669, -11.20 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (364.9 DPS) | yes | Magni's Will (12548, +0.00 DPS) [quest]; Painweaver Band (13098, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -1.02 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heart of Wyrmthalak (22321, +0.00 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -4.08 DPS, sim-verified) [crafted] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -1.66 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -16.27 DPS) [dungeon] |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Riphook (12653, -0.68 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.90 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.00 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; waist: Highlander's Plate Girdle; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Hand of Justice; main_hand: Felstriker; off_hand: Persuader; ranged: Bloodseeker

No-known-source sample (15 of 2366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 975.6. Weights run: 8.7s. Verify run: 135.2s. 2366 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.362 ± 0.099, crit=1.748 ± 0.134 per rating point (14 rating = 1%, 24.474 per %), hit=3.344 ± 0.293 per rating point (10 rating = 1%, 33.437 per %), melee_haste=20.691 ± 3.273

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 155.4 attack_power points (45.76 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -15.11 DPS) [vendor]; Helm of the Executioner (22411, -26.86 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Highlander's Plate Spaulders (20057) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -6.07 DPS) [rep]; Cape of the Black Baron (13340, -6.19 DPS) [dungeon]; Cloak of Flames (3475, -7.79 DPS, sim-verified) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -5.41 DPS) [dungeon]; Dawn Armor (252483, -6.52 DPS, sim-verified) [crafted]; Cadaverous Armor (14637, -9.74 DPS) [dungeon] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, -0.77 DPS) [rep]; Marshal's Plate Bracers (16481, -1.70 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gauntlets (271940, +0.00 DPS) [rep] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 65.3 attack_power points (19.22 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Valiant's Waistguard (272403, -0.52 DPS) [vendor]; Brigam Girdle (13142, -5.69 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Plate Legguards (231540, -1.08 DPS) [pvp]; Sentinel's Plate Legguards (22672, -3.10 DPS) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -5.72 DPS) [vendor]; Protector's Band (19514, -8.99 DPS) [rep]; Naglering (11669, -25.42 DPS, sim-verified) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.28 DPS) [vendor]; Protector's Band (19514, -3.55 DPS) [rep]; Painweaver Band (13098, -3.82 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -14.18 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (975.6 DPS) | yes | - |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -9.22 DPS, sim-verified) [crafted] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 707.8 attack_power points (208.41 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -37.80 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -3.06 DPS) [quest]; Riphook (12653, -4.57 DPS) [dungeon]; Skull Splitting Crossbow (13039, -4.98 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Highlander's Plate Spaulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Berserker Bracers; waist: Might of the Timbermaw; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Greenhammer; ranged: Blackcrow

No-known-source sample (15 of 2366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 00000000000000000-050510000000000000-000000000000000000)

Set DPS (verified): 36.4. Weights run: 6.4s. Verify run: 26.2s. 332 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.006 ± 0.024, crit=0.567 ± 0.025 per rating point (14 rating = 1%, 7.939 per %), hit=0.869 ± 0.054 per rating point (10 rating = 1%, 8.686 per %), melee_haste=5.240 ± 0.341

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.15 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.57 DPS) [crafted]; Brawler's Leather Hood (252504, -0.69 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.0 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.19 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.34 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.11 DPS) [dungeon]; Subterranean Cape (14149, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.4 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.27 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.57 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.17 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.23 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.4 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.33 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.34 DPS) [crafted]; Deviate Scale Belt (6468, -0.40 DPS) [crafted]; Cobrahn's Grasp (6460, -2.72 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.4 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.20 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.0 attack_power points (0.86 DPS) | yes | Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Veteran's Boots (250503, -0.06 DPS) [crafted]; Defender's Leather Boots (252441, -0.29 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.0 attack_power points (0.69 DPS) | yes | Loop of Sacrifice (281673, -0.35 DPS) [quest]; Bounty Hunter's Ring (5351, -0.52 DPS) [quest]; The 1 Ring (8350, -0.52 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.46 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Bounty Hunter's Ring (5351, -0.29 DPS) [quest]; The 1 Ring (8350, -0.29 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.5 attack_power points (13.71 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (13.38 DPS) | yes | Cruel Barb (5191, -0.26 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.23 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.06 DPS, sim-verified) [crafted]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 332, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (undead, 00000000000000000-050530025000100000-000000000000000000)

Set DPS (verified): 63.4. Weights run: 7.5s. Verify run: 15.8s. 581 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.176 ± 0.038, crit=0.860 ± 0.042 per rating point (14 rating = 1%, 12.042 per %), hit=1.194 ± 0.107 per rating point (10 rating = 1%, 11.936 per %), melee_haste=7.591 ± 0.772

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 28.6 attack_power points (1.88 DPS) | yes | Tusken Helm (6686, -0.17 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.30 DPS) [crafted]; Defender's Leather Helm (252455, -0.30 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.7 attack_power points (0.83 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Scout's Medallion (19537, -0.22 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.9 attack_power points (1.51 DPS) | yes | Barbaric Iron Shoulders (7913, -0.25 DPS) [crafted]; Barbaric Shoulders (5964, -0.46 DPS) [crafted]; Golden Scale Shoulders (3841, -0.59 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (63.4 DPS) | yes | Wolfmaster Cape (6314, -0.28 DPS) [dungeon]; Wildhunter Cloak (16658, -0.28 DPS) [quest]; Construct Cloak (279848, -0.28 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-verified (63.4 DPS) | yes | Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.19 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (1.05 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.22 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.44 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.09 DPS) [world_drop]; Warsong Gauntlets (16978, -0.13 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.1 attack_power points (1.78 DPS) | yes | Prowler's Leather Belt (252459, -0.13 DPS) [crafted]; Girdle of Golem Strength (9405, -0.20 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.20 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.2 attack_power points (1.98 DPS) | yes | Brawler's Leather Legguards (252516, -0.37 DPS) [crafted]; Ferine Leggings (6690, -0.39 DPS, sim-verified) [dungeon]; Glimmering Mail Legguards (6386, -0.39 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.9 attack_power points (1.04 DPS) | yes | Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Hard Gold Boots (250534, -0.12 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.5 attack_power points (1.28 DPS) | yes | Band of the Fist (17694, -0.45 DPS) [quest]; Tiger Band (6749, -0.63 DPS) [quest]; Bonebinder's Signet (273032, -0.63 DPS) [world] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.1 attack_power points (1.25 DPS) | yes | Band of the Fist (17694, -0.45 DPS, sim-verified) [quest]; Tiger Band (6749, -0.59 DPS) [quest]; Bonebinder's Signet (273032, -0.59 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-verified (63.4 DPS) | yes | Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop]; Thornspike (6681, -2.12 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (20.89 DPS) | yes | Zealot Blade (13033, -1.96 DPS, sim-verified) [world_drop] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.5 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.10 DPS) [world]; Golemsight Long Gun (273029, -0.16 DPS) [dungeon]; Long Battle Bow (15284, -0.23 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 581, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 00000000000000000-050530025122100221-000000000000000000)

Set DPS (verified): 144.6. Weights run: 7.4s. Verify run: 28.2s. 797 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.535 ± 0.051, crit=1.426 ± 0.067 per rating point (14 rating = 1%, 19.968 per %), hit=1.304 ± 0.110 per rating point (10 rating = 1%, 13.041 per %), melee_haste=6.739 ± 1.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.0 attack_power points (4.48 DPS) | yes | Chromite Barbute (8142, -0.35 DPS) [dungeon]; White Bandit Mask (10008, -0.69 DPS) [crafted]; Barbaric Iron Helm (7915, -1.38 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Ethereal Talisman (4430, -0.07 DPS) [quest]; Kaleidoscope Chain (13084, -0.27 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.9 attack_power points (2.62 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.48 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.55 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.44 DPS) [dungeon]; Parachute Cloak (10518, -0.44 DPS) [crafted]; Sporid Cape (6629, -1.99 DPS, sim-verified) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (144.6 DPS) | yes | Quillward Harness (10583, -0.46 DPS) [dungeon]; Kolkar Marauder Chain (6773, -1.09 DPS) [quest]; Carapace of Tuten'kash (10775, -1.72 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.1 attack_power points (1.96 DPS) | yes | Branded Leather Bracers (19508, -0.01 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.29 DPS) [world_drop]; Yorgen Bracers (13012, -0.34 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -2.36 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 31.8 attack_power points (3.10 DPS) | yes | Boar Champion's Belt (10768, -0.18 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.18 DPS) [rep]; Defiler's Plate Girdle (20206, -0.18 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (4.09 DPS) | yes | Symbolic Legplates (14829, -0.27 DPS) [world_drop]; Firemane Leggings (13129, -0.39 DPS) [world_drop]; Orcish War Leggings (7929, -0.78 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 36.7 attack_power points (3.58 DPS) | yes | Blackforge Greaves (6423, -0.33 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.39 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.57 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.3 attack_power points (2.76 DPS) | yes | Mark of Kern (2262, -0.81 DPS) [dungeon]; Assault Band (13095, -0.81 DPS) [world_drop]; Falcon's Hook (7552, -0.83 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.6 attack_power points (2.01 DPS) | yes | Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop]; Falcon's Hook (7552, -0.08 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ardent Custodian (868, +0.00 DPS) [world_drop]; Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (43.09 DPS) | yes | Southsea Lamp (9359, +0.00 DPS) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.24 DPS) [quest]; The Silencer (13138, -0.25 DPS) [world_drop]; Swiftwind (13038, -0.57 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 797, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 30321000000000000-050530025122100231-000000000000000000)

Set DPS (verified): 229.0. Weights run: 8.1s. Verify run: 32.0s. 1011 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.466 ± 0.086, crit=1.753 ± 0.117 per rating point (14 rating = 1%, 24.537 per %), hit=2.643 ± 0.192 per rating point (10 rating = 1%, 26.425 per %), melee_haste=14.513 ± 1.876

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 87.0 attack_power points (10.76 DPS) | yes | Blood Guard's Plate Helm (220803, -1.24 DPS) [vendor]; Embrace of the Lycan (9479, -4.82 DPS) [dungeon]; Raging Berserker's Helm (7719, -4.95 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.1 attack_power points (3.60 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.12 DPS) [quest]; Scout's Medallion (19535, -1.42 DPS) [rep]; Woven Ivy Necklace (19159, -1.54 DPS, sim-verified) [quest] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 37.7 attack_power points (4.67 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.25 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.31 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Pridelord Cape (14673, -0.35 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -1.17 DPS) [dungeon]; Warbear Harness (15064, -1.64 DPS) [crafted]; Stone Guard's Plate Armor (220801, -2.47 DPS, sim-verified) [vendor] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (3.81 DPS) | yes | Prowler's Leather Bracers (252539, -0.31 DPS) [crafted]; Runed Golem Shackles (12550, -0.35 DPS) [dungeon]; Deepfury Bracers (13120, -1.26 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.62 DPS, sim-verified) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 46.5 attack_power points (5.76 DPS) | yes | Prowler's Leather Waistguard (252473, -0.12 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.25 DPS) [rep]; Girdle of Beastial Fury (11686, -1.99 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 63.2 attack_power points (7.82 DPS) | yes | Stormshroud Pants (15057, -1.75 DPS) [crafted]; Golem Shard Leggings (13074, -2.38 DPS) [world_drop]; Serpentskin Leggings (8262, -3.14 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (229.0 DPS) | yes | Battlechaser's Greaves (12555, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.36 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.4 attack_power points (5.75 DPS) | yes | White Bone Band (11862, -2.78 DPS) [quest]; Masons Fraternity Ring (9533, -3.21 DPS) [quest]; Thunderbrow Ring (13097, -3.22 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 33.2 attack_power points (4.11 DPS) | yes | White Bone Band (11862, -1.14 DPS) [quest]; Masons Fraternity Ring (9533, -1.57 DPS) [quest]; Thunderbrow Ring (13097, -1.58 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.28 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Flurry Axe (871, -2.73 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadowblade (2163, -1.00 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -6.47 DPS) [dungeon]; White Bone Shredder (11863, -10.06 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.07 DPS) [dungeon]; Houndmaster's Bow (11628, -0.31 DPS) [dungeon]; Skull Splitting Crossbow (13039, -2.29 DPS, sim-verified) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Officer's Wristguards; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1011, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 360.2. Weights run: 7.8s. Verify run: 132.1s. 2356 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.870 ± 0.190, crit=4.099 ± 0.272 per rating point (14 rating = 1%, 57.392 per %), hit=4.276 ± 0.446 per rating point (10 rating = 1%, 42.763 per %), melee_haste=24.018 ± 4.329

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 236.3 attack_power points (25.38 DPS) | yes | Outlaw's Collar (279253, -6.13 DPS, sim-verified) [crafted]; Mask of the Unforgiven (13404, -10.03 DPS) [dungeon]; Champion's Plate Helm (227043, -10.11 DPS) [pvp] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 83.4 attack_power points (8.95 DPS) | yes | Emberfury Talisman (12929, -3.77 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -3.88 DPS) [world_drop]; Scout's Medallion (19534, -4.33 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 134.9 attack_power points (14.48 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -1.09 DPS) [pvp]; Truestrike Shoulders (12927, -2.72 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Fel Cape (279269, -0.61 DPS) [crafted]; Howler's Furs (272414, -0.99 DPS, sim-verified) [vendor]; Blackveil Cape (11626, -1.17 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -2.08 DPS) [dungeon]; Timbermaw Tunic (252484, -3.36 DPS) [crafted]; Breastplate of Bloodthirst (12757, -7.04 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Vambraces of the Sadist (13400, -0.55 DPS) [dungeon]; Battleborn Armbraces (12936, -0.95 DPS, sim-verified) [dungeon]; Blinkstrike Armguards (15860, -1.41 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, -3.40 DPS, sim-verified) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Chain Girdle (20150, -0.59 DPS) [rep]; Defiler's Leather Girdle (20190, -0.59 DPS) [rep]; Marksman's Girdle (22232, -3.01 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, -0.33 DPS) [pvp]; Outrider's Plate Legguards (22651, -4.77 DPS) [rep]; Legionnaire's Plate Leggings (227048, -6.64 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 100.2 attack_power points (10.75 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.46 DPS) [dungeon]; Shadowcraft Boots (16711, -2.91 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -4.38 DPS) [dungeon]; Blackstone Ring (17713, -5.52 DPS) [dungeon]; Naglering (11669, -10.50 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.03 DPS) [dungeon]; Blackstone Ring (17713, -2.17 DPS) [dungeon]; Band of Rumination (18103, -2.74 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.58 DPS) [crafted]; Hand of Justice (11815, -2.57 DPS) [dungeon] |
| trinket2 | Heart of Wyrmthalak (22321) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (360.2 DPS) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -1.01 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -15.18 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.20 DPS) [world_drop]; Bloodseeker (19107, -1.64 DPS) [quest]; Ancient Bone Bow (18680, -2.13 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Heart of Wyrmthalak; main_hand: Persuader; off_hand: Ironwood Blade; ranged: Blackcrow

No-known-source sample (15 of 2356, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 975.8. Weights run: 8.7s. Verify run: 124.4s. 2356 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.003, agility=3.329 ± 0.234, crit=4.306 ± 0.326 per rating point (14 rating = 1%, 60.278 per %), hit=3.932 ± 0.526 per rating point (10 rating = 1%, 39.321 per %), melee_haste=24.872 ± 5.367

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 238.8 attack_power points (50.16 DPS) | yes | Eye of Rend (12587, -18.83 DPS) [dungeon]; Champion's Plate Helm (227043, -19.53 DPS) [pvp]; Outlaw's Collar (279253, -30.24 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emberfury Talisman (12929, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Scout's Medallion (19534, -0.91 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 150.2 attack_power points (31.54 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -3.78 DPS) [pvp]; Truestrike Shoulders (12927, -9.98 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fel Cape (279269, -2.03 DPS) [crafted]; Blackveil Cape (11626, -2.13 DPS) [dungeon]; Howler's Furs (272414, -6.02 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -4.56 DPS, sim-verified) [dungeon]; Timbermaw Tunic (252484, -6.30 DPS) [crafted]; Breastplate of Bloodthirst (12757, -15.46 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Vambraces of the Sadist (13400, -2.47 DPS) [dungeon]; Berserker Bracers (19578, -3.99 DPS) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, -1.15 DPS) [quest]; General's Plate Gauntlets (231532, -2.06 DPS) [pvp]; Gloves of Holy Might (867, -6.06 DPS, sim-verified) [world_drop] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | 110.7 attack_power points (23.25 DPS) | yes | Defiler's Plate Girdle (20204, -2.73 DPS) [rep]; General's Plate Girdle (16547, -3.31 DPS) [pvp]; Marksman's Girdle (22232, -13.89 DPS, sim-verified) [dungeon] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -1.62 DPS) [rep]; Outrider's Chain Leggings (22673, -8.07 DPS, sim-verified) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 105.9 attack_power points (22.24 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -3.53 DPS) [dungeon]; Shadowcraft Boots (16711, -5.08 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -7.84 DPS) [dungeon]; Legionnaire's Band (19510, -10.62 DPS) [rep]; Naglering (11669, -23.49 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -4.81 DPS) [rep]; Painweaver Band (13098, -5.36 DPS, sim-verified) [dungeon]; Band of Rumination (18103, -5.39 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -5.52 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (975.8 DPS) | yes | Frozen Heart of the Mountain (249469, -1.29 DPS) [crafted]; Hand of Justice (11815, -4.52 DPS) [dungeon] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Felstriker (12590, -9.08 DPS, sim-verified) [dungeon] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 729.4 attack_power points (153.21 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -5.03 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -31.50 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.57 DPS) [world_drop]; Bloodseeker (19107, -1.77 DPS) [quest]; Ancient Bone Bow (18680, -2.67 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Valiant's Waistguard; legs: Outrider's Plate Legguards; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; main_hand: Persuader; off_hand: Greenhammer; ranged: Blackcrow

No-known-source sample (15 of 2356, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

