# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-050510000000000000-000000000000000000)

Set DPS (verified): 35.4. Weights run: 4.3s. Verify run: 15.6s. 388 eligible items had no known source.

Race sweep: high-order-skyborne 35.5 ± 0.0, night-elf 35.4 ± 0.0, human 35.4 ± 0.0, dwarf 35.0 ± 0.0, gnome 34.6 ± 0.0 (dps); high-order-skyborne beats the guide race night-elf by 0.2 (0.48%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.121 ± 0.025, crit=0.611 ± 0.027 per rating point (14 rating = 1%, 8.547 per %), hit=0.802 ± 0.056 per rating point (10 rating = 1%, 8.016 per %), melee_haste=6.207 ± 0.343

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.49 DPS) [crafted]; Brawler's Leather Hood (252504, -0.61 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.7 attack_power points (0.37 DPS) | yes | Erudite's Amulet (277204, -0.30 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.71 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.33 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.02 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.11 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Armor (252490, -0.12 DPS) [crafted]; Defender's Leather Armor (252434, -0.15 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.56 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.65 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.31 DPS) [crafted]; Deviate Scale Belt (6468, -0.36 DPS) [crafted]; Cobrahn's Grasp (6460, -2.91 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.61 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -2.39 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.5 attack_power points (0.69 DPS) | yes | Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Ring of the Moon (12052, -0.58 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.44 DPS) | yes | Ring of the Moon (12052, -0.33 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.7 attack_power points (13.26 DPS) | yes | Rockslicer (872, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (12.93 DPS) | yes | Cruel Barb (5191, +0.00 DPS) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.5 attack_power points (0.25 DPS) | yes | Dwarven Fishing Pole (3567, -0.03 DPS) [quest]; Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 30 (night-elf, 00000000000000000-050530025000100000-000000000000000000)

Set DPS (verified): 62.4. Weights run: 4.6s. Verify run: 9.2s. 637 eligible items had no known source.

Race sweep: high-order-skyborne 62.6 ± 0.1, night-elf 62.4 ± 0.1, human 62.3 ± 0.1, dwarf 61.7 ± 0.1, gnome 61.3 ± 0.1 (dps); high-order-skyborne beats the guide race night-elf by 0.3 (0.45%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.245 ± 0.040, crit=0.914 ± 0.046 per rating point (14 rating = 1%, 12.793 per %), hit=1.261 ± 0.105 per rating point (10 rating = 1%, 12.608 per %), melee_haste=7.558 ± 0.824

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.2 attack_power points (1.88 DPS) | yes | Veteran's Chain Helm (250498, -0.34 DPS) [crafted]; Defender's Leather Helm (252455, -0.34 DPS) [crafted]; Tusken Helm (6686, -0.47 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.0 attack_power points (0.84 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Sentinel's Medallion (19541, -0.19 DPS) [rep]; River Pride Choker (13087, -0.32 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.7 attack_power points (1.53 DPS) | yes | Barbaric Iron Shoulders (7913, -0.36 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.48 DPS) [crafted]; Golden Scale Shoulders (3841, -0.63 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, -0.11 DPS) [pvp]; Wolfmaster Cape (6314, -0.30 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.16 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (1.03 DPS) | yes | Yorgen Bracers (13012, -0.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Barbaric Bracers (18948, -0.19 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.42 DPS) | yes | Insignia Gloves (6408, -0.03 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.16 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.5 attack_power points (1.77 DPS) | yes | Prowler's Leather Belt (252459, -0.13 DPS) [crafted]; Girdle of Golem Strength (9405, -0.22 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.7 attack_power points (1.98 DPS) | yes | Brawler's Leather Legguards (252516, -0.36 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.39 DPS) [world_drop]; Ferine Leggings (6690, -0.52 DPS, sim-verified) [dungeon] |
| feet | Alacritous Treads (277234) (or Brawler's Leather Boots (252439)) | WANTED: Incinerator Gar'im [quest] | 16.2 attack_power points (1.05 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -0.02 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.7 attack_power points (1.27 DPS) | yes | Silverlaine's Family Seal (6321, -0.63 DPS) [dungeon]; Tiger Band (6749, -0.63 DPS) [quest]; Bonebinder's Signet (273032, -0.63 DPS) [world] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.5 attack_power points (1.26 DPS) | yes | Silverlaine's Family Seal (6321, -0.61 DPS) [dungeon]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop]; Thornspike (6681, -2.50 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (20.52 DPS) | yes | Zealot Blade (13033, -2.16 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -5.56 DPS) [quest] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.7 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.11 DPS) [world]; Golemsight Long Gun (273029, -0.15 DPS) [dungeon]; Long Battle Bow (15284, -0.24 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Alacritous Treads; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 637, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 00000000000000000-050530025121100231-000000000000000000)

Set DPS (verified): 149.0. Weights run: 4.6s. Verify run: 16.8s. 946 eligible items had no known source.

Race sweep: night-elf 149.0 ± 0.2, human 148.2 ± 0.2, high-order-skyborne 147.4 ± 0.2, gnome 146.0 ± 0.2, dwarf 145.9 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.646 ± 0.066, crit=1.550 ± 0.085 per rating point (14 rating = 1%, 21.698 per %), hit=1.949 ± 0.131 per rating point (10 rating = 1%, 19.487 per %), melee_haste=10.243 ± 1.330

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.7 attack_power points (5.49 DPS) | yes | Chromite Barbute (8142, -0.45 DPS) [dungeon]; White Bandit Mask (10008, -0.87 DPS) [crafted]; Barbaric Iron Helm (7915, -1.71 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 18.1 attack_power points (2.08 DPS) | yes | Kaleidoscope Chain (13084, -0.40 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.47 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.70 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.1 attack_power points (3.23 DPS) | yes | Flintrock Shoulders (7755, -0.19 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.70 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.72 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.50 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.69 DPS) [quest]; Parachute Cloak (10518, -1.00 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Kolkar Marauder Chain (6773, -0.89 DPS) [quest]; Icemail Jerkin (1981, -0.91 DPS, sim-verified) [world_drop]; Carapace of Tuten'kash (10775, -1.62 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (2.37 DPS) | yes | Branded Leather Bracers (19508, -0.07 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.31 DPS) [world_drop]; Yorgen Bracers (13012, -0.42 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.03 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.8 attack_power points (3.77 DPS) | yes | Boar Champion's Belt (10768, -0.32 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.32 DPS) [rep]; Highlander's Plate Girdle (20125, -0.32 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (4.83 DPS) | yes | Symbolic Legplates (14829, -0.24 DPS) [world_drop]; Firemane Leggings (13129, -0.46 DPS) [world_drop]; Triprunner Dungarees (9624, -0.73 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.5 attack_power points (4.32 DPS) | yes | Boots of Avoidance (14549, -0.29 DPS) [world_drop]; Blackforge Greaves (6423, -0.35 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.46 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 29.2 attack_power points (3.36 DPS) | yes | Falcon's Hook (7552, -0.96 DPS) [dungeon]; Ring of the Underwood (2951, -1.00 DPS) [world_drop]; Mark of Kern (2262, -1.05 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (2.41 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.11 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Ardent Custodian (868, -1.15 DPS, sim-verified) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (50.85 DPS) | yes | Southsea Lamp (9359, -0.69 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -24.15 DPS) [quest] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.27 DPS) [quest]; The Silencer (13138, -0.34 DPS) [world_drop]; Swiftwind (13038, -0.62 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Thunderbrow Ring; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 946, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 30312000000000000-050530025122100231-000000000000000000)

Set DPS (verified): 230.7. Weights run: 4.7s. Verify run: 14.6s. 1228 eligible items had no known source.

Race sweep: night-elf 230.7 ± 0.2, high-order-skyborne 230.1 ± 0.2, dwarf 229.5 ± 0.2, human 229.4 ± 0.2, gnome 227.7 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.467 ± 0.081, crit=1.694 ± 0.112 per rating point (14 rating = 1%, 23.715 per %), hit=2.481 ± 0.180 per rating point (10 rating = 1%, 24.812 per %), melee_haste=13.416 ± 1.810

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 84.5 attack_power points (10.61 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.25 DPS) [vendor]; Raging Berserker's Helm (7719, -4.00 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -4.58 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.1 attack_power points (3.65 DPS) | yes | Sentinel's Medallion (19539, -1.45 DPS, sim-verified) [rep]; Zealous Shadowshard Pendant (17772, -1.89 DPS) [quest] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 44.8 attack_power points (5.62 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.95 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -1.14 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Phantom Cape (13122, -1.20 DPS, sim-verified) [world_drop]; Blisterbane Wrap (12552, -1.32 DPS) [dungeon]; Sergeant Major's Cape (16336, -1.47 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight's Plate Hauberk (220794, -0.79 DPS) [vendor]; Mixologist's Tunic (12793, -1.17 DPS) [dungeon]; Warbear Harness (15064, -1.46 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (3.87 DPS) | yes | Deepfury Bracers (13120, -0.10 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.32 DPS) [crafted]; Runed Golem Shackles (12550, -0.35 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.67 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (5.77 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.04 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.05 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.18 DPS) [crafted] |
| legs | Knight's Plate Leggings (220797) | PvP rank 6 · Knight · Alliance [vendor] | 62.4 attack_power points (7.83 DPS) | yes | Infernal Trickster Leggings (17754, -1.29 DPS, sim-verified) [dungeon]; Gryphon Rider's Leggings (9652, -1.30 DPS) [quest]; Centurion Legplates (10740, -1.30 DPS) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 47.1 attack_power points (5.91 DPS) | yes | Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Skulker's Leather Boots (252469, -0.75 DPS) [crafted]; Officer's Sabatons (250561, -0.99 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.8 attack_power points (5.62 DPS) | yes | Masons Fraternity Ring (9533, -3.05 DPS) [quest]; Thunderbrow Ring (13097, -3.06 DPS) [world_drop]; Mark of Kern (2262, -3.11 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 33.2 attack_power points (4.17 DPS) | yes | Masons Fraternity Ring (9533, -1.59 DPS) [quest]; Thunderbrow Ring (13097, -1.61 DPS) [world_drop]; Mark of Kern (2262, -1.66 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths [dungeon] | 535.1 attack_power points (67.14 DPS) | yes | Dawn's Edge (12774, -1.19 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -6.63 DPS) [dungeon]; Shoni's Disarming Tool (9608, -38.02 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.27 DPS) [dungeon]; Houndmaster's Bow (11628, -0.52 DPS) [dungeon]; Skull Splitting Crossbow (13039, -1.11 DPS, sim-verified) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Icemail Jerkin; wrist: Officer's Wristguards; waist: Girdle of Beastial Fury; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Doomforged Straightedge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1228, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 361.8. Weights run: 5.0s. Verify run: 84.0s. 2728 eligible items had no known source.

Race sweep: night-elf 361.8 ± 0.4, high-order-skyborne 361.7 ± 0.4, dwarf 360.8 ± 0.4, human 360.2 ± 0.4, gnome 358.4 ± 0.4 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.263 ± 0.083, crit=1.805 ± 0.119 per rating point (14 rating = 1%, 25.269 per %), hit=1.403 ± 0.190 per rating point (10 rating = 1%, 14.034 per %), melee_haste=13.042 ± 2.014

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 114.6 attack_power points (17.85 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.19 DPS) [vendor]; Fury Visor (20521, -9.02 DPS, sim-verified) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 51.3 attack_power points (7.99 DPS) | yes | Will of the Martyr (17044, -3.31 DPS) [quest]; Imperial Jewel (11933, -3.58 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -3.87 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 59.4 attack_power points (9.25 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, -0.02 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.27 DPS) [rep]; Cape of the Black Baron (13340, -0.48 DPS) [dungeon]; Blackveil Cape (11626, -1.92 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -4.04 DPS, sim-verified) [dungeon]; Dawn Armor (252483, -4.36 DPS) [crafted]; Breastplate of Bloodthirst (12757, -5.84 DPS) [quest] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 48.1 attack_power points (7.49 DPS) | yes | Forest Stalker's Bracers (19587, -0.33 DPS) [rep]; Marshal's Plate Bracers (16481, -0.82 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gauntlets (271940, -0.98 DPS, sim-verified) [rep] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 59.7 attack_power points (9.30 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Plate Girdle (20041, -0.07 DPS) [rep]; Highlander's Chain Girdle (20043, -0.07 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Plate Legguards (231540, -2.18 DPS) [pvp]; Sentinel's Chain Leggings (237819, -2.41 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -5.28 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 48.1 attack_power points (7.49 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Battlechaser's Greaves (12555, -0.57 DPS) [dungeon]; Highlander's Plate Greaves (20048, -0.77 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -2.40 DPS) [rep]; Painweaver Band (13098, -2.50 DPS) [dungeon]; Naglering (11669, -11.02 DPS, sim-verified) [dungeon] |
| finger2 | Magni's Will (12548) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -0.21 DPS) [rep]; Painweaver Band (13098, -0.31 DPS) [dungeon]; Blackstone Ring (17713, -0.82 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Cannonball Runner (13382, -0.78 DPS) [dungeon] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironfoe (11684, +0.00 DPS) [dungeon]; Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (99.43 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -9.16 DPS) [dungeon]; Alcor's Sunrazor (14555, -15.29 DPS, sim-verified) [world_drop] |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Riphook (12653, -0.44 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.52 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.75 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Berserker Bracers; waist: Might of the Timbermaw; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Magni's Will; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Persuader; off_hand: Felstriker; ranged: Bloodseeker

No-known-source sample (15 of 2728, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 1021.8. Weights run: 5.6s. Verify run: 90.8s. 2728 eligible items had no known source.

Race sweep: high-order-skyborne 1025.0 ± 1.1, night-elf 1021.8 ± 1.0, human 1021.8 ± 1.0, gnome 1021.3 ± 1.0, dwarf 1017.7 ± 1.0 (dps); high-order-skyborne beats the guide race night-elf by 3.2 (0.31%) ± 1.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.331 ± 0.089, crit=1.728 ± 0.122 per rating point (14 rating = 1%, 24.197 per %), hit=3.047 ± 0.254 per rating point (10 rating = 1%, 30.472 per %), melee_haste=17.929 ± 2.953

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 148.9 attack_power points (44.60 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -14.40 DPS) [vendor]; Fury Visor (20521, -16.19 DPS, sim-verified) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 50.2 attack_power points (15.03 DPS) | yes | Imperial Jewel (11933, -5.45 DPS) [dungeon]; Nacreous Shell Necklace (22403, -5.91 DPS) [dungeon]; Star of Mystaria (12103, -17.50 DPS, sim-verified) [dungeon] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 69.7 attack_power points (20.88 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Highlander's Lamellar Spaulders (20058, -2.24 DPS) [rep]; Highlander's Plate Spaulders (20057, -5.47 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -5.34 DPS) [rep]; Cape of the Black Baron (13340, -5.54 DPS) [dungeon]; Cloak of Flames (3475, -8.35 DPS, sim-verified) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -9.17 DPS) [dungeon]; Dawn Armor (252483, -12.49 DPS, sim-verified) [crafted]; Field Marshal's Plate Armor (231536, -12.79 DPS) [pvp] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 54.7 attack_power points (16.37 DPS) | yes | Berserker Bracers (19578, -0.67 DPS) [rep]; Forest Stalker's Bracers (19587, -1.55 DPS) [rep]; Marshal's Plate Bracers (16481, -2.38 DPS) [pvp] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Gauntlets (271940, -5.20 DPS) [rep]; Marshal's Plate Gauntlets (231541, -6.95 DPS) [pvp]; Timbermaw Brawlers (19049, -7.96 DPS) [crafted] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Plate Girdle (20041, -0.34 DPS) [rep]; Might of the Timbermaw (19044, -2.65 DPS, sim-verified) [crafted] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -3.84 DPS) [vendor]; Marshal's Plate Legguards (231540, -4.61 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -2.53 DPS) [crafted]; Bloodmail Boots (14616, -2.92 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -5.45 DPS) [dungeon]; Protector's Band (19514, -8.28 DPS) [rep]; Naglering (11669, -20.30 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19514, -2.41 DPS) [rep]; Magni's Will (12548, -2.85 DPS) [quest]; Blackstone Ring (17713, -3.79 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+22.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -10.26 DPS, sim-verified) [crafted] |
| main_hand | Ironfoe (11684) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Persuader (22384, -19.93 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (191.13 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -17.61 DPS) [dungeon]; Alcor's Sunrazor (14555, -34.54 DPS, sim-verified) [world_drop] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.35 DPS) [quest]; Riphook (12653, -2.82 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -3.00 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Trueaim Gauntlets; waist: Valiant's Waistguard; legs: Sentinel's Plate Legguards; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Ironfoe; off_hand: Felstriker; ranged: Blackcrow

No-known-source sample (15 of 2728, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 00000000000000000-050510000000000000-000000000000000000)

Set DPS (verified): 36.5. Weights run: 4.3s. Verify run: 15.9s. 371 eligible items had no known source.

Race sweep: undead 36.5 ± 0.0, windshaper-skyborne 35.6 ± 0.0, troll 35.5 ± 0.0, orc 35.3 ± 0.0, tauren 35.1 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.121 ± 0.025, crit=0.611 ± 0.027 per rating point (14 rating = 1%, 8.547 per %), hit=0.802 ± 0.056 per rating point (10 rating = 1%, 8.016 per %), melee_haste=6.207 ± 0.343

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.49 DPS) [crafted]; Brawler's Leather Hood (252504, -0.61 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.7 attack_power points (0.37 DPS) | yes | Erudite's Amulet (277204, -0.19 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.33 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.02 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon]; Subterranean Cape (14149, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.27 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.56 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.22 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.33 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.31 DPS) [crafted]; Deviate Scale Belt (6468, -0.36 DPS) [crafted]; Cobrahn's Grasp (6460, -2.72 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.20 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.6 attack_power points (0.87 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Veteran's Boots (250503, -0.06 DPS) [crafted]; Defender's Leather Boots (252441, -0.31 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.5 attack_power points (0.69 DPS) | yes | Loop of Sacrifice (281673, -0.36 DPS) [quest]; Bounty Hunter's Ring (5351, -0.51 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Bounty Hunter's Ring (5351, -0.26 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.7 attack_power points (13.26 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 232.8 attack_power points (12.93 DPS) | yes | Cruel Barb (5191, -0.26 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.5 attack_power points (0.25 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.06 DPS, sim-verified) [crafted]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Butcher's Cleaver; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 371, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (undead, 00000000000000000-050530025000100000-000000000000000000)

Set DPS (verified): 63.6. Weights run: 4.6s. Verify run: 8.2s. 622 eligible items had no known source.

Race sweep: undead 63.6 ± 0.1, windshaper-skyborne 62.6 ± 0.1, troll 62.4 ± 0.1, orc 62.0 ± 0.1, tauren 61.6 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.245 ± 0.040, crit=0.914 ± 0.046 per rating point (14 rating = 1%, 12.793 per %), hit=1.261 ± 0.105 per rating point (10 rating = 1%, 12.608 per %), melee_haste=7.558 ± 0.824

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.2 attack_power points (1.88 DPS) | yes | Tusken Helm (6686, -0.21 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.34 DPS) [crafted]; Defender's Leather Helm (252455, -0.34 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.0 attack_power points (0.84 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Scout's Medallion (19537, -0.19 DPS) [rep]; River Pride Choker (13087, -0.32 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.7 attack_power points (1.53 DPS) | yes | Barbaric Iron Shoulders (7913, -0.27 DPS) [crafted]; Barbaric Shoulders (5964, -0.48 DPS) [crafted]; Golden Scale Shoulders (3841, -0.63 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wolfmaster Cape (6314, -0.30 DPS) [dungeon]; Wildhunter Cloak (16658, -0.30 DPS) [quest]; Construct Cloak (279848, -0.30 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.16 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (1.03 DPS) | yes | Yorgen Bracers (13012, -0.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Barbaric Bracers (18948, -0.19 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.42 DPS) | yes | Insignia Gloves (6408, -0.03 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Warsong Gauntlets (16978, -0.13 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.5 attack_power points (1.77 DPS) | yes | Prowler's Leather Belt (252459, -0.13 DPS) [crafted]; Girdle of Golem Strength (9405, -0.22 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.7 attack_power points (1.98 DPS) | yes | Brawler's Leather Legguards (252516, -0.36 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.39 DPS) [world_drop]; Ferine Leggings (6690, -0.39 DPS, sim-verified) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.2 attack_power points (1.05 DPS) | yes | Feet of the Lynx (1121, -0.02 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Hard Gold Boots (250534, -0.14 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.7 attack_power points (1.27 DPS) | yes | Band of the Fist (17694, -0.44 DPS) [quest]; Tiger Band (6749, -0.63 DPS) [quest]; Bonebinder's Signet (273032, -0.63 DPS) [world] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.5 attack_power points (1.26 DPS) | yes | Band of the Fist (17694, -0.45 DPS, sim-verified) [quest]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop]; Thornspike (6681, -2.12 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (20.52 DPS) | yes | Zealot Blade (13033, -1.96 DPS, sim-verified) [world_drop] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.7 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.11 DPS) [world]; Golemsight Long Gun (273029, -0.15 DPS) [dungeon]; Long Battle Bow (15284, -0.24 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 622, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 00000000000000000-050530025121100231-000000000000000000)

Set DPS (verified): 148.8. Weights run: 4.6s. Verify run: 15.7s. 928 eligible items had no known source.

Race sweep: undead 148.8 ± 0.2, troll 146.9 ± 0.2, windshaper-skyborne 146.9 ± 0.2, orc 145.9 ± 0.2, tauren 145.0 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.646 ± 0.066, crit=1.550 ± 0.085 per rating point (14 rating = 1%, 21.698 per %), hit=1.949 ± 0.131 per rating point (10 rating = 1%, 19.487 per %), melee_haste=10.243 ± 1.330

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 47.7 attack_power points (5.49 DPS) | yes | Chromite Barbute (8142, -0.45 DPS) [dungeon]; White Bandit Mask (10008, -0.87 DPS) [crafted]; Barbaric Iron Helm (7915, -1.71 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 18.1 attack_power points (2.08 DPS) | yes | Ethereal Talisman (4430, -0.17 DPS) [quest]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.1 attack_power points (3.23 DPS) | yes | Flintrock Shoulders (7755, -0.19 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.70 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.72 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.50 DPS) [dungeon]; Parachute Cloak (10518, -0.50 DPS) [crafted]; Tigerstrike Mantle (13108, -0.50 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Quillward Harness (10583, -0.67 DPS, sim-verified) [dungeon]; Kolkar Marauder Chain (6773, -2.24 DPS) [quest]; Carapace of Tuten'kash (10775, -2.97 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (2.37 DPS) | yes | Branded Leather Bracers (19508, -0.07 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.31 DPS) [world_drop]; Yorgen Bracers (13012, -0.42 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.19 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.8 attack_power points (3.77 DPS) | yes | Boar Champion's Belt (10768, -0.32 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.32 DPS) [rep]; Defiler's Plate Girdle (20206, -0.32 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (4.83 DPS) | yes | Symbolic Legplates (14829, -0.24 DPS) [world_drop]; Firemane Leggings (13129, -0.46 DPS) [world_drop]; Triprunner Dungarees (9624, -0.73 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.5 attack_power points (4.32 DPS) | yes | Boots of Avoidance (14549, -0.29 DPS) [world_drop]; Blackforge Greaves (6423, -0.35 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.46 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 29.2 attack_power points (3.36 DPS) | yes | Falcon's Hook (7552, -0.96 DPS) [dungeon]; Ring of the Underwood (2951, -1.00 DPS) [world_drop]; Mark of Kern (2262, -1.05 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (2.41 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.11 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Ardent Custodian (868, -0.80 DPS, sim-verified) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (50.85 DPS) | yes | Southsea Lamp (9359, +0.00 DPS) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.27 DPS) [quest]; The Silencer (13138, -0.34 DPS) [world_drop]; Swiftwind (13038, -0.62 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 928, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 30312000000000000-050530025122100231-000000000000000000)

Set DPS (verified): 234.2. Weights run: 4.7s. Verify run: 14.6s. 1209 eligible items had no known source.

Race sweep: undead 234.2 ± 0.2, windshaper-skyborne 231.4 ± 0.2, troll 230.0 ± 0.2, orc 229.5 ± 0.2, tauren 227.8 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.467 ± 0.081, crit=1.694 ± 0.112 per rating point (14 rating = 1%, 23.715 per %), hit=2.481 ± 0.180 per rating point (10 rating = 1%, 24.812 per %), melee_haste=13.416 ± 1.810

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 84.5 attack_power points (10.61 DPS) | yes | Blood Guard's Plate Helm (220803, -1.25 DPS) [vendor]; Embrace of the Lycan (9479, -4.58 DPS) [dungeon]; Raging Berserker's Helm (7719, -4.95 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.1 attack_power points (3.65 DPS) | yes | Woven Ivy Necklace (19159, -0.48 DPS, sim-verified) [quest]; Scout's Medallion (19535, -1.44 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 44.8 attack_power points (5.62 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.14 DPS) [world_drop]; Officer's Pauldrons (250576, -1.66 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.32 DPS) [dungeon]; Dark Phantom Cape (13122, -1.53 DPS, sim-verified) [world_drop]; Pridelord Cape (14673, -1.67 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stone Guard's Plate Armor (220801, -0.79 DPS) [vendor]; Mixologist's Tunic (12793, -1.17 DPS) [dungeon]; Warbear Harness (15064, -1.46 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (3.87 DPS) | yes | Deepfury Bracers (13120, -0.10 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.32 DPS) [crafted]; Runed Golem Shackles (12550, -0.35 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.94 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (5.77 DPS) | yes | Defiler's Lamellar Girdle (20179, -0.04 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.05 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.18 DPS) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 6 · Stone Guard · Horde [vendor] | 62.4 attack_power points (7.83 DPS) | yes | Serpentskin Leggings (8262, -1.62 DPS) [world_drop]; Infernal Trickster Leggings (17754, -1.71 DPS, sim-verified) [dungeon]; Stormshroud Pants (15057, -1.88 DPS) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 47.1 attack_power points (5.91 DPS) | yes | Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Skulker's Leather Boots (252469, -0.75 DPS) [crafted]; Officer's Sabatons (250561, -0.99 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.8 attack_power points (5.62 DPS) | yes | White Bone Band (11862, -2.61 DPS) [quest]; Masons Fraternity Ring (9533, -3.05 DPS) [quest]; Thunderbrow Ring (13097, -3.06 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 33.2 attack_power points (4.17 DPS) | yes | White Bone Band (11862, -1.16 DPS) [quest]; Masons Fraternity Ring (9533, -1.59 DPS) [quest]; Thunderbrow Ring (13097, -1.61 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.13 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.16 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Hammer of the Northern Wind (810, -1.38 DPS, sim-verified) [world_drop] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths [dungeon] | 535.1 attack_power points (67.14 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -6.63 DPS) [dungeon]; White Bone Shredder (11863, -10.26 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.27 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.27 DPS) [world_drop]; Houndmaster's Bow (11628, -0.52 DPS) [dungeon] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; wrist: Officer's Wristguards; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Doomforged Straightedge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1209, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 361.9. Weights run: 5.0s. Verify run: 91.2s. 2718 eligible items had no known source.

Race sweep: undead 361.9 ± 0.4, windshaper-skyborne 358.6 ± 0.4, troll 357.2 ± 0.4, orc 356.3 ± 0.4, tauren 354.1 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.996 ± 0.184, crit=4.280 ± 0.263 per rating point (14 rating = 1%, 59.921 per %), hit=3.980 ± 0.403 per rating point (10 rating = 1%, 39.801 per %), melee_haste=19.795 ± 3.875

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 235.4 attack_power points (25.27 DPS) | yes | Eye of Rend (12587, -7.24 DPS, sim-verified) [dungeon]; Champion's Plate Helm (227043, -10.06 DPS) [pvp]; Mask of the Unforgiven (13404, -10.30 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 85.9 attack_power points (9.22 DPS) | yes | Star of Mystaria (12103, -3.67 DPS) [dungeon]; Emberfury Talisman (12929, -3.75 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -3.97 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 140.8 attack_power points (15.11 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -1.83 DPS) [pvp]; Champion's Plate Shoulders (227042, -5.03 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.31 DPS) [dungeon]; Blackveil Cape (11626, -1.49 DPS) [dungeon]; Deathguard's Cloak (20068, -2.02 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Gladiator Chain (11726, -1.72 DPS) [dungeon]; Dawn Armor (252483, -3.41 DPS, sim-verified) [crafted]; Breastplate of Bloodthirst (12757, -3.65 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Vambraces of the Sadist (13400, -0.54 DPS) [dungeon]; Battleborn Armbraces (12936, -0.67 DPS, sim-verified) [dungeon]; Berserker Bracers (19578, -1.82 DPS) [rep] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Gloves (272099, -1.77 DPS) [vendor]; General's Plate Gauntlets (231532, -2.09 DPS) [pvp] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Chain Girdle (20150, -0.64 DPS) [rep]; Defiler's Leather Girdle (20190, -0.64 DPS) [rep]; Marksman's Girdle (22232, -2.24 DPS, sim-verified) [dungeon] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Chain Leggings (22673, -0.61 DPS, sim-verified) [rep]; Legionnaire's Plate Leggings (227048, -1.28 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 99.7 attack_power points (10.70 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.93 DPS) [dungeon]; Shadowcraft Boots (16711, -2.67 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -4.49 DPS) [dungeon]; Band of the Penitent (13217, -5.77 DPS) [quest]; Naglering (11669, -9.49 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.05 DPS) [dungeon]; Band of the Penitent (13217, -2.34 DPS) [quest]; Band of Rumination (18103, -2.34 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Molten Heart of the Mountain (249470, -2.22 DPS, sim-verified) [crafted] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+10.6 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -6.31 DPS) [dungeon]; Heartseeker (12783, -10.65 DPS, sim-verified) [crafted] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.66 DPS) [dungeon]; Bloodseeker (19107, -1.50 DPS) [quest]; Ancient Bone Bow (18680, -1.93 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Trueaim Gauntlets; waist: Valiant's Waistguard; legs: Outrider's Plate Legguards; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Persuader; off_hand: Felstriker

No-known-source sample (15 of 2718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30323013002000000-050530025122100231-200000000000000000)

Set DPS (verified): 1011.3. Weights run: 5.6s. Verify run: 89.4s. 2718 eligible items had no known source.

Race sweep: windshaper-skyborne 1012.5 ± 1.0, undead 1011.3 ± 1.0, orc 1009.3 ± 1.1, troll 1007.4 ± 1.0, tauren 998.7 ± 1.1 (dps); windshaper-skyborne beats the guide race undead by 1.2 (0.12%) ± 1.5, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.331 ± 0.089, crit=1.728 ± 0.122 per rating point (14 rating = 1%, 24.197 per %), hit=3.047 ± 0.254 per rating point (10 rating = 1%, 30.472 per %), melee_haste=17.929 ± 2.953

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 148.9 attack_power points (44.60 DPS) | yes | Champion's Plate Helm (227043, -14.40 DPS) [pvp]; Stalwart Helm (250599, -18.45 DPS) [crafted]; Fury Visor (20521, -24.53 DPS, sim-verified) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 50.2 attack_power points (15.03 DPS) | yes | Imperial Jewel (11933, -5.45 DPS) [dungeon]; Nacreous Shell Necklace (22403, -5.91 DPS) [dungeon]; Star of Mystaria (12103, -15.36 DPS, sim-verified) [dungeon] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 69.7 attack_power points (20.88 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Defiler's Lamellar Spaulders (20184, -2.24 DPS) [rep]; Defiler's Plate Spaulders (20212, -2.24 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -5.54 DPS) [dungeon]; Blackveil Cape (11626, -7.98 DPS) [dungeon]; Deathguard's Cloak (20068, -10.51 DPS, sim-verified) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -9.17 DPS) [dungeon]; Dawn Armor (252483, -10.43 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -12.79 DPS) [pvp] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 54.7 attack_power points (16.37 DPS) | yes | Forest Stalker's Bracers (19587, -1.55 DPS) [rep]; General's Plate Armguards (16546, -2.38 DPS) [pvp]; Berserker Bracers (19578, -9.77 DPS, sim-verified) [rep] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Gauntlets (272095, -5.20 DPS) [vendor]; General's Plate Gauntlets (231532, -6.95 DPS) [pvp]; Storm Gauntlets (12632, -11.04 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 64.8 attack_power points (19.41 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Valiant's Waistguard (272403, -0.63 DPS) [vendor]; Defiler's Plate Girdle (20204, -0.97 DPS) [rep] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Chain Leggings (22673, -5.37 DPS, sim-verified) [rep]; Warbear Woolies (15065, -9.20 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | 60.9 attack_power points (18.24 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Windreaver Greaves (13967, -1.15 DPS) [dungeon]; Scalegut Treaders (275618, -3.67 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -5.86 DPS) [vendor]; Legionnaire's Band (19510, -8.28 DPS) [rep]; Naglering (11669, -28.46 DPS, sim-verified) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.41 DPS) [vendor]; Legionnaire's Band (19510, -2.83 DPS) [rep]; Painweaver Band (13098, -4.27 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+19.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Rune of the Guard Captain (19120, -12.34 DPS, sim-verified) [quest] |
| main_hand | Ironfoe (11684) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Persuader (22384, -10.32 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (191.13 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -17.61 DPS) [dungeon]; Alcor's Sunrazor (14555, -35.29 DPS, sim-verified) [world_drop] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Riphook (12653, -2.82 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -3.00 DPS) [world_drop]; Bloodseeker (19107, -9.03 DPS, sim-verified) [quest] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Trueaim Gauntlets; waist: Might of the Timbermaw; legs: Outrider's Plate Legguards; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Ironfoe; off_hand: Felstriker; ranged: Blackcrow

No-known-source sample (15 of 2718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

